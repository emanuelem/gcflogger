package gcflogger

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

// Logger defines the Cloud Functions logger object.
type Logger struct{}

// Entry defines a log entry.
type Entry struct {
	Message  string `json:"message"`
	Severity string `json:"severity,omitempty"`
}

// String renders an entry structure to the JSON format expected by Cloud Logging.
func (e Entry) String() string {
	out, err := json.Marshal(e)
	if err != nil {
		log.Printf("json.Marshal error: %v", err)
	}
	return string(out)
}

// New creates a new logger designed to work with Google Cloud Functions.
// ctx is a context.
func New(ctx context.Context) *Logger {
	log.SetFlags(0)
	return &Logger{}
}

// Default prints a log entry with no assigned severity level.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Default(v ...interface{}) {
	log.Println(Entry{
		Severity: "DEFAULT",
		Message:  fmt.Sprint(v...),
	})
}

// Defaultf prints a log entry with no assigned severity level.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Defaultf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "DEFAULT",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Debug prints a log entry with severity level 100.
// It should be used for debug or trace information.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Debug(v ...interface{}) {
	log.Println(Entry{
		Severity: "DEBUG",
		Message:  fmt.Sprint(v...),
	})
}

// Debugf prints a log entry with severity level 100.
// It should be used for debug or trace information.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "DEBUG",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Info prints a log entry with severity level 200.
// It should be used for routine information, such as ongoing status or
// performance.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Info(v ...interface{}) {
	log.Println(Entry{
		Severity: "INFO",
		Message:  fmt.Sprint(v...),
	})
}

// Infof prints a log entry with severity level 200.
// It should be used for routine information, such as ongoing status or
// performance.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Infof(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "INFO",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Notice prints a log entry with severity level 300.
// It should be used for normal but significant events, such as start up, shut
// down, or a configuration change.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Notice(v ...interface{}) {
	log.Println(Entry{
		Severity: "NOTICE",
		Message:  fmt.Sprint(v...),
	})
}

// Noticef prints a log entry with severity level 300.
// It should be used for normal but significant events, such as start up, shut
// down, or a configuration change.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Noticef(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "NOTICE",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Warning prints a log entry with severity level 400.
// It should be used for warning events that might cause problems.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Warning(v ...interface{}) {
	log.Println(Entry{
		Severity: "WARNING",
		Message:  fmt.Sprint(v...),
	})
}

// Warningf prints a log entry with severity level 400.
// It should be used for warning events that might cause problems.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Warningf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "WARNING",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Error prints a log entry with severity level 500.
// It should be used for error events that are likely to cause problems.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Error(v ...interface{}) {
	log.Println(Entry{
		Severity: "ERROR",
		Message:  fmt.Sprint(v...),
	})
}

// Errorf prints a log entry with severity level 500.
// It should be used for error events that are likely to cause problems.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "ERROR",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Critical prints a log entry with severity level 600.
// It should be used for critical events that cause more severe problems or
// outages.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Critical(v ...interface{}) {
	log.Println(Entry{
		Severity: "CRITICAL",
		Message:  fmt.Sprint(v...),
	})
}

// Criticalf prints a log entry with severity level 600.
// It should be used for critical events that cause more severe problems or
// outages.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "CRITICAL",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Alert prints a log entry with severity level 700.
// It should be used when a person must take an action immediately.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Alert(v ...interface{}) {
	log.Println(Entry{
		Severity: "ALERT",
		Message:  fmt.Sprint(v...),
	})
}

// Alertf prints a log entry with severity level 700.
// It should be used when a person must take an action immediately.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Alertf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "ALERT",
		Message:  fmt.Sprintf(format, v...),
	})
}

// Emergency prints a log entry with severity level 800.
// It should be used when one or more systems are unusable.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Emergency(v ...interface{}) {
	log.Println(Entry{
		Severity: "EMERGENCY",
		Message:  fmt.Sprint(v...),
	})
}

// Emergencyf prints a log entry with severity level 800.
// It should be used when one or more systems are unusable.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Emergencyf(format string, v ...interface{}) {
	log.Println(Entry{
		Severity: "EMERGENCY",
		Message:  fmt.Sprintf(format, v...),
	})
}
