package config

// LogConfig defines the logging configuration
type LogConfig struct {
	// ErrorLogFilePath specifies the error log file path.
	// If it's empty then the log is redirected to std.err
	ErrorLogFilePath string `json:"error"`

	// DebugLogFilePath specifies the error log file path.
	// If it's empty then the log is redirected to std.out
	DebugLogFilePath string `json:"debug"`
}

// TLSConfig defines the TLS configuration
type TLSConfig struct {
	// FullChainCertFilePath specifies the path
	// to the full-chain certificate file
	FullChainCertFilePath string `json:"cert"`

	// KeyFilePath specifies the path to the certificate's private key file
	KeyFilePath string `json:"key"`
}

// Config defines the configuration of the API server instance
type Config struct {
	// Address must specifies the address:[port] combination
	// for the API server to listen on
	ServerAddress string `json:"server-addr"`

	// MetricsServerAddress specifies the address:[port] combination
	// for the metrics server to listen on
	MetricsServerAddress string `json:"metrics-server-addr"`

	// TLS specifies the TLS encryption configuration,
	// which is disabled if TLS is nil
	TLS *TLSConfig `json:"tls"`

	// Log specifies the logging configuration.
	// Logs will be redirected to std.out and std.err respectively
	// if Log is nil
	Log *LogConfig `json:"log"`
}
