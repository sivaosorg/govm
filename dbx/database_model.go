package dbx

type Dbx struct {
	IsConnected   bool   `json:"is_connected"`
	DebugMode     bool   `json:"debug_mode"`
	IsNewInstance bool   `json:"is_new_instance"`
	Pid           int    `json:"instance_pid,omitempty"`
	Error         error  `json:"error,omitempty"`
	Message       string `json:"message,omitempty"`
	Database      string `json:"database,omitempty"`
}

type DbxConfig struct {
}
