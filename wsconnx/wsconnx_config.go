package wsconnx

// WebSocket connection statuses
const (
	StatusConnected     = "connected"     // The client has successfully connected to the WebSocket server.
	StatusDisconnected  = "disconnected"  // The client has disconnected from the WebSocket server.
	StatusReconnecting  = "reconnecting"  // The client is attempting to reconnect to the WebSocket server after a disconnection.
	StatusClosed        = "closed"        // The WebSocket connection has been closed intentionally.
	StatusConnecting    = "connecting"    // The client is in the process of establishing a connection to the WebSocket server.
	StatusAuthenticated = "authenticated" // The client has successfully authenticated with the WebSocket server, if authentication is required.
	StatusFailed        = "failed"        // The WebSocket connection attempt has failed due to an error or connection timeout.
	StatusTerminated    = "terminated"    // The WebSocket connection has been terminated either by the server or the client.
	StatusIdle          = "idle"          // The WebSocket connection is open, but there is no active communication or data exchange.
	StatusBusy          = "busy"          // The WebSocket connection is actively handling data exchanges and processing messages.
	StatusError         = "error"         // An error has occurred in the WebSocket connection, and it needs attention or troubleshooting.
	StatusReconnected   = "reconnected"   // The client has successfully reconnected to the WebSocket server after a previous disconnection.
	StatusStale         = "stale"         // The WebSocket connection has been idle for a prolonged period, and it might need to be refreshed.
)

// WebSocket connection scopes
const (
	ScopePublic    = "public"     // A public scope where data is accessible to all users.
	ScopePrivate   = "private"    // A private scope where data is accessible only to authorized users.
	ScopeGroup     = "group"      // A group scope where data is accessible to users within a specific group or channel.
	ScopeRead      = "read"       // A read-only scope where users can receive data but cannot send data.
	ScopeWrite     = "write"      // A write-only scope where users can send data but cannot receive data.
	ScopeReadWrite = "read-write" // A read-write scope where users can both send and receive data.
	ScopeAdmin     = "admin"      // An administrative scope with elevated privileges for server-side operations and management.
)
