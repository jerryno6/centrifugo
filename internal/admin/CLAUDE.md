# Admin Package

Admin web UI handlers for the Centrifugo administration interface.

## Purpose
Provides REST API endpoints for the embedded admin web UI, including:
- Server metrics and statistics
- Channel management
- Connection monitoring
- Configuration inspection

## Endpoints

### API Endpoints
- `GET /admin/settings` - Get admin web interface settings (insecure mode, edition)
- `POST /admin/auth` - Authenticate admin user and get secure token
- `POST /admin/api` - Main admin API endpoint (requires secure token)

### Static Content
- `GET /*` - Serve static admin web UI assets

## Authentication
- **Token-based authentication** using secure cookies
- **Insecure mode** for development (bypasses auth)
- **Password-based login** for obtaining secure tokens

## Files
- `handlers.go` - HTTP handlers for admin UI endpoints
- `handlers_test.go` - Tests for admin handlers