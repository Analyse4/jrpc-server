package jlog

import (
	"io"
	"log"
	"runtime"
	"strconv"
	"time"
)

// Log level.
const (
	INFO  = 200
	DEBUG = 201
)

// Log flag
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

// A JLogger represents an active logging object that generates lines of
// output to an io.Writer. Each logging operation makes a single call to
// the Writer's Write method. A JLogger can be used simultaneously from
// multiple goroutines; it guarantees to serialize access to the Writer.
type JLogger struct {
	stdlog *log.Logger
	level  int
	flag   int
}

// New creates a new JLogger. The out variable sets the
// destination to which log data will be written.
// The prefix appears at the beginning of each generated log line, or
// after the log header if the Lmsgprefix flag is provided.
// The flag argument defines the logging properties.
func New(out io.Writer, prefix string, flag int) *JLogger {
	jl := new(JLogger)
	jl.stdlog = log.New(out, prefix, 0)
	jl.flag = flag
	return jl
}

// SetLevel set log level.
func (jl *JLogger) SetLevel(level int) {
	jl.level = level
}

// Info only print logs prefixed by INFO.
// Info calls jl.stdlog.Println to print the logger.
// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (jl *JLogger) Info(v ...interface{}) {
	var s string
	var sf string

	jl.stdlog.SetPrefix("[INFO] ")

	if jl.flag == LstdFlags|Lshortfile {

		_, fn, ln, ok := runtime.Caller(1)
		if ok {
			sf = fn + ":" + strconv.Itoa(ln)
		}
		s = time.Now().String() + " " + sf + ":"
	}

	jl.stdlog.Println(s, v)

	// _, fn, ln, ok := runtime.Caller(1)
	// fmt.Println(fn, ln, ok)
}

// Debug only print logs prefixed by INFO and higher.
// Debug calls jl.stdlog.Println to print the logger.
// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (jl *JLogger) Debug(v ...interface{}) {
	if jl.level != INFO {
		jl.stdlog.SetPrefix("[DEBUG] ")
		jl.stdlog.Println(v...)
	}
}
