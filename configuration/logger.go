package configuration

// LoggerConfig holds the configuration for the logger
type LoggerConfig struct {
	FileOutput bool   // True to log to a file, False to log to console
	FilePath   string // File path to log if FileOutput is true
	JSONFormat bool   // True for JSON logs, False for plain text
	LogLevel   string // Log level: "debug", "info", "warn", "error"
}
