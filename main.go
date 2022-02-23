package doggo

import (
	"log"
	"os"
)

var (
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

const (
	errorFlags = log.Ldate | log.Ltime | log.Lshortfile
	infoFlags  = log.LstdFlags
)

// Must be called before using any of the logging functions.
func Init() {
	errorLogger = log.New(os.Stderr, "", errorFlags)
	infoLogger = log.New(os.Stdout, "", infoFlags)
}

// Logs to stdout which is interpreted as level 6 priority
// in `journald` when the `journald` logging driver is used.
// Appears in white in `journalctl`.
func Print(v ...interface{}) {
	infoLogger.Print(v...)
}

// Logs to stdout which is interpreted as level 6 priority
// in `journald` when the `journald` logging driver is used.
// Appears in white in `journalctl`.
func Printf(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

// Logs to stderr which is interpreted as level 3 priority
// in `journald` when the `journald` logging driver is used.
// Appears in red in `journalctl`.
func Error(v ...interface{}) {
	errorLogger.Print(v...)
}

// Logs to stderr which is interpreted as level 3 priority
// in `journald` when the `journald` logging driver is used.
// Appears in red in `journalctl`.
func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

// Logs to stderr which is interpreted as level 3 priority
// in `journald` when the `journald` logging driver is used.
// Appears in red in `journalctl`.
// Included for backward compatibility.
func Fatal(v ...interface{}) {
	log.Fatal(v...)
}

// Logs to stderr which is interpreted as level 3 priority
// in `journald` when the `journald` logging driver is used.
// Appears in red in `journalctl`.
// Included for backward compatibility.
func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
