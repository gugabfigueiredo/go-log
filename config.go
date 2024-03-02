package log

import "github.com/rs/zerolog"

// Config is the configuration for the logger, supported by envconfig: https://github.com/kelseyhightower/envconfig
type Config struct {
	// Level is the log level to use
	Level zerolog.Level `default:"1"`
	// Context is the application name
	Context string `required:"true"`
	// Enable console logging
	ConsoleLoggingEnabled bool `default:"true"`
	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool `default:"true"`
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool `default:"false"`
	// Directory to log to when file logging is enabled
	Directory string `default:"/var/log"`
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string `default:"my-app"`
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int `default:"10"`
	// MaxBackups the max number of rolled files to keep
	MaxBackups int `default:"10"`
	// MaxAge the max age in days to keep a logfile
	MaxAge int `default:"2"`
}
