package proxy

import (
	"encoding/base64"
	"time"

	"github.com/centrifugal/centrifugo/v6/internal/configtypes"
	"github.com/centrifugal/centrifugo/v6/internal/proxyproto"

	"github.com/centrifugal/centrifuge"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog/log"
)

// MessageHandlerConfig ...
type MessageHandlerConfig struct {
	Proxies map[string]MessageProxy
}

// MessageHandler ...
type MessageHandler struct {
	config    MessageHandlerConfig
	summary   map[string]prometheus.Observer
	histogram map[string]prometheus.Observer
	errors    map[string]prometheus.Counter
	inflight  map[string]prometheus.Gauge
}

// NewMessageHandler ...
func NewMessageHandler(c MessageHandlerConfig) *MessageHandler {
	h := &MessageHandler{
		config: c,
	}
	summary := map[string]prometheus.Observer{}
	histogram := map[string]prometheus.Observer{}
	errors := map[string]prometheus.Counter{}
	inflight := map[string]prometheus.Gauge{}
	for name, p := range c.Proxies {
		summary[name] = proxyCallDurationSummary.WithLabelValues(p.Protocol(), "message", name)
		histogram[name] = proxyCallDurationHistogram.WithLabelValues(p.Protocol(), "message", name)
		errors[name] = proxyCallErrorCount.WithLabelValues(p.Protocol(), "message", name)
		inflight[name] = proxyCallInflightRequests.WithLabelValues(p.Protocol(), "message", name)
	}
	h.summary = summary
	h.histogram = histogram
	h.errors = errors
	h.inflight = inflight
	return h
}

// MessageHandlerFunc ...
type MessageHandlerFunc func(Client, centrifuge.MessageEvent, configtypes.ChannelOptions, PerCallData) (centrifuge.PublishReply, error)

// Handle Message.
func (h *MessageHandler) Handle(node *centrifuge.Node) MessageHandlerFunc {
	return func(client Client, e centrifuge.MessageEvent, chOpts configtypes.ChannelOptions, pcd PerCallData) (centrifuge.PublishReply, error) {
		started := time.Now()

		var p MessageProxy
		var summary prometheus.Observer
		var histogram prometheus.Observer
		var errors prometheus.Counter

		proxyEnabled := chOpts.MessageProxyEnabled
		proxyName := chOpts.MessageProxyName
		if !proxyEnabled {
			// log.Info().Str("channel", e.Data.Channel).Msg("message proxy not enabled for a channel")
			log.Info().Str("channel", "LEON log this message_handler.line68").Msg("message proxy not enabled for a channel")
			return centrifuge.PublishReply{}, centrifuge.ErrorNotAvailable
		}
		p = h.config.Proxies[proxyName]
		summary = h.summary[proxyName]
		histogram = h.histogram[proxyName]
		errors = h.errors[proxyName]
		inflight := h.inflight[proxyName]
		inflight.Inc()
		defer inflight.Dec()

		req := &proxyproto.MessageRequest{
			Client:    client.ID(),
			Protocol:  string(client.Transport().Protocol()),
			Transport: client.Transport().Name(),
			Encoding:  getEncoding(p.UseBase64()),

			User: client.UserID(),
			// Channel: e.Channel,
		}
		if p.IncludeMeta() && pcd.Meta != nil {
			req.Meta = proxyproto.Raw(pcd.Meta)
		}
		if !p.UseBase64() {
			req.Data = e.Data
		} else {
			req.B64Data = base64.StdEncoding.EncodeToString(e.Data)
		}

		messageRep, err := p.ProxyMessage(client.Context(), req)
		duration := time.Since(started).Seconds()
		if err != nil {
			select {
			case <-client.Context().Done():
				// Client connection already closed.
				return centrifuge.PublishReply{}, centrifuge.DisconnectConnectionClosed
			default:
			}
			summary.Observe(duration)
			histogram.Observe(duration)
			errors.Inc()
			// log.Error().Err(err).Str("client", client.ID()).Str("channel", e.Channel).Msg("error proxying message")
			return centrifuge.PublishReply{}, err
		}
		summary.Observe(duration)
		histogram.Observe(duration)

		if messageRep.Disconnect != nil {
			return centrifuge.PublishReply{}, proxyproto.DisconnectFromProto(messageRep.Disconnect)
		}
		if messageRep.Error != nil {
			return centrifuge.PublishReply{}, proxyproto.ErrorFromProto(messageRep.Error)
		}

		// historySize := chOpts.HistorySize
		// historyTTL := chOpts.HistoryTTL
		// historyMetaTTL := chOpts.HistoryMetaTTL

		data := e.Data
		if messageRep.Result != nil {
			if messageRep.Result.Data != nil {
				data = messageRep.Result.Data
			} else if messageRep.Result.B64Data != "" {
				decodedData, err := base64.StdEncoding.DecodeString(messageRep.Result.B64Data)
				if err != nil {
					log.Error().Err(err).Str("client", client.ID()).Msg("error decoding base64 data")
					return centrifuge.PublishReply{}, centrifuge.ErrorInternal
				}
				data = decodedData
			}

			// if messageRep.Result.SkipHistory {
			// 	historySize = 0
			// 	historyTTL = 0
			// }
		}
		_ = data
		// result, err := node.Publish(
		// 	e.Channel, data,
		// 	centrifuge.WithClientInfo(e.ClientInfo),
		// 	centrifuge.WithHistory(historySize, historyTTL.ToDuration(), historyMetaTTL.ToDuration()),
		// )
		// messageResult := centrifuge.PublishResult{
		// 	StreamPosition: {},
		// 	FromCache:     false
		// }
		return centrifuge.PublishReply{}, err
	}
}
