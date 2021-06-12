# gcflogger
This package is a simple wrapper of the standard Go log [package](https://golang.org/pkg/log/) that adds [structured logging](https://cloud.google.com/logging/docs/structured-logging) suitable for delivering Google Cloud Function logs to Google Cloud Logging.

## Install
```
go get -u github.com/emanuelem/gcflogger
```

## Sample usage
```
package main

import (
	"context"

	"github.com/emanuelem/gcflogger"
)

func main() {
	ctx := context.Background()
	logger := gcflogger.New(ctx)

	logger.Notice("I am a notice!")
	code := 300
	logger.Noticef("I am also a notice with code %v!", code)
}
```

The available methods correspond to the severity levels:

| Enum      | Code | Description |
| ----------|------| ------------|
| DEFAULT   | 0    | The log entry has no assigned severity level. |
| DEBUG     | 100  | Debug or trace information. |
| INFO      | 200  | Routine information, such as ongoing status or performance. |
| NOTICE    | 300  | Normal but significant events, such as start up, shut down, or a configuration change. |
| WARNING   | 400  | Warning events might cause problems. |
| ERROR     | 500  | Error events are likely to cause problems. |
| CRITICAL  | 600  | Critical events cause more severe problems or outages. |
| ALERT     | 700  | A person must take an action immediately. |
| EMERGENCY | 800  | One or more systems are unusable. |

## License
GPLv3 - See [LICENSE][license] file

[license]: https://github.com/emanuelem/gcflogger/blob/master/LICENSE