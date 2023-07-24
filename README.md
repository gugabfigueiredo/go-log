# tiny-go-log
A tiny conveniently setup go rolling logger powered by [zerolog](https://github.com/rs/zerolog) and [lumberjack](https://github.com/natefinch/lumberjack)

## Usage

```go
package main

import (
    "github.com/gugabfigueiredo/tiny-go-log"
)

func main() {
    
	// create new logger from local config.go struct
	logger := log.New(&log.Config{
		Context: "my-app",
		// Enable console logging
		ConsoleLoggingEnabled: true,
		// EncodeLogsAsJson makes the log framework log JSON
		EncodeLogsAsJson: true,
		// FileLoggingEnabled makes the framework log to a file
		// the fields below can be skipped if this value is false!
		FileLoggingEnabled: true,
		// Directory to log to when file logging is enabled
		Directory: "/var/log/my-app",
		// Filename is the name of the logfile which will be placed inside the directory
		Filename: "my-app",
		// MaxSize the max size in MB of the logfile before it's rolled
		MaxSize: 10,
		// MaxBackups the max number of rolled files to keep
		MaxBackups: 3,
		// MaxAge the max age in days to keep a logfile
		MaxAge: 2,
    })
	
	logger.I("Hello World")
}
```

expected output:

```bash
{"level":"info","time":"2020-05-01T00:00:00Z","message":"Hello World","context":"my-app"}
```