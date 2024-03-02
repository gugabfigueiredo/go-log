# tiny-go-log
A tiny conveniently setup go rolling logger powered by [zerolog](https://github.com/rs/zerolog) and [lumberjack](https://github.com/natefinch/lumberjack) based on [panta's gist](https://gist.github.com/panta/2530672ca641d953ae452ecb5ef79d7d)

## Usage

```go
package main

import (
    "github.com/gugabfigueiredo/tiny-go-log"
)

func main() {
    
    // create new logger from local config.go struct
    logger := log.New(&log.Config{
        // Context is the application name. 
        // It is also used to compose the file path if file logging is enabled
        Context: "my-app",
        // Level is the log level the logger should log at
        Level: "info",
        // Enable console logging, defaults to true
        ConsoleLoggingEnabled: true,
        // EncodeLogsAsJson makes the log framework log JSON
        EncodeLogsAsJson: true,
        // FileLoggingEnabled makes the framework log to a file
        // the fields below can be skipped if this value is false!
        FileLoggingEnabled: true,
        // Directory to log to when file logging is enabled
        Directory: "/var/log",
        // Filename is the name of the logfile which will be placed inside the directory
        Filename: "my-app",
        // MaxSize the max size in MB of the logfile before it's rolled
        MaxSize: 10,
        // MaxBackups the max number of rolled files to keep
        MaxBackups: 3,
        // MaxAge the max age in days to keep a logfile
        MaxAge: 2,
    })
    
    logger.Info("Hello World")
}
```

expected output:

```
{"level":"info","time":"2020-05-01T00:00:00Z","message":"Hello World","context":"my-app"}
```