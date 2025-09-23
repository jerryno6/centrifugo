package client

type WSMessage struct {
	Type string `json:"type,intern"`
	Data Score  `json:"data"`
}

// Score represents a score submission from client
type Score struct {
	GameID     string `json:"gameId,intern"`
	UserID     string `json:"userId,intern"`
	Score      int    `json:"score"`
	TotalScore int    `json:"totalScore"`
	Timestamp  int64  `json:"timestamp"` // epoch time in UnixMicro()
}
