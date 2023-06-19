package dbx

type Dbx struct {
	IsConnected bool   `json:"is_connected"`
	Error       error  `json:"error,omitempty"`
	Message     string `json:"message,omitempty"`
	Database    string `json:"database,omitempty"`
}

type DbxConfig struct {
}
