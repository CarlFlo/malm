package malm

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// Here are the various types of logs that can be used.
const (
	logFatal uint8 = 1 << iota
	logError
	logWarning
	logInfo
	logDebug
	logCustom
)

var (
	logBitmask     uint8
	verboseBitmask uint8
	defaultWriter  io.Writer
	timeFormat     string
)

// Initializes the package with default settings.
func init() {

	defaultWriter = os.Stderr
	timeFormat = "2006-01-02 15:04:05"
	TurnOnAllLogging()
	TurnOnAllVerbose()
}

// Fatal is used for logging a fatal problem that has occured and will thus be using the [FATAL] tag.
// Fatal works similarly to log.Fatal and is the only logging function that wont return a bool on completion.
// Please note that Fatal will run os.Exit(1).
func Fatal(format string, args ...interface{}) {
	formatter(defaultWriter, logFatal, "FATAL", format, args...)
	os.Exit(1)
}

// Error is used for logging an error and will thus be using the [ERROR] tag.
func Error(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logError, "ERROR", format, args...)
}

// Warn is used for logging a warning and will thus be using the [WARN] tag.
func Warn(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logWarning, "WARN", format, args...)
}

// Info is used for logging information and will thus be using the [INFO] tag.
func Info(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logInfo, "INFO", format, args...)
}

// Debug is used for logging debug messages and will thus be using the [DEBUG] tag.
func Debug(format string, args ...interface{}) bool {
	return formatter(defaultWriter, logDebug, "DEBUG", format, args...)
}

// Custom is used for logging a customized messages under the tag of the users choice.
// It allows the specification of a custom io.Writer, the default is os.Stderr.
// logTag is a string that specifices what the tag will say.
func Custom(writer io.Writer, logTag string, format string, args ...interface{}) bool {
	// The numerical log level for a custom log message is 1
	return formatter(writer, logCustom, logTag, format, args...)
}

// formatter formats and crafts the log message.
// It also makes sure if it is supposed to be printed.
// Returns true on success and false on block.
func formatter(writer io.Writer, numericalLogType uint8, logTag string, format string, args ...interface{}) bool {

	// Checks if the message shouldn't be printed
	if numericalLogType&logBitmask == 0 {
		return false
	}

	now := time.Now().Format(timeFormat)

	message := fmt.Sprintf(format, args...)

	// Checks if the message shouldn't be verbose
	if numericalLogType&verboseBitmask == 0 {

		// <date and time> [<log tag>] <formatted message>\n
		fmt.Fprintf(writer, "%s [%s] %s\n", now, logTag, message)
		return true
	}

	// Message is verbose
	fileName, caller, line := getDetails()

	// <date and time> [<log tag>] <filePath>:<line number>:<caller>() <formatted message>\n
	fmt.Fprintf(writer, "%s [%s] %s:%d:%s() %s\n", now, logTag, fileName, line, caller, message)
	return true
}

// This function retrieves the function which called the function,
// the file it is in and the line the function is on.
func getDetails() (string, string, int) {
	pc, path, line, ok := runtime.Caller(3)

	// Something went wrong
	if !ok {
		return "???", "???", -1
	}

	paths := strings.Split(path, "/")
	file := paths[len(paths)-1]

	caller := runtime.FuncForPC(pc).Name()
	stack := strings.Split(caller, ".")
	caller = stack[len(stack)-1]

	return file, caller, line
}

// SetDefaultWriter allows for changing of the default io.Writer that the logger uses for outputting the message.
func SetDefaultWriter(newWriter io.Writer) {
	defaultWriter = newWriter
}

// SetLogBitmask allows for changing the permission of what types of log messages gets outputted.
// A log gets printed if the bitmask 'allows it'.
// This function allows the user to quickly update the entire bitmask.
//
// The value '63' will turn on all logging
func SetLogBitmask(bitmask uint8) uint8 {
	logBitmask = bitmask
	return logBitmask
}

// SetLogVerboseBitmask allows for changing the permission of what types of log messages gets outputted with more information.
// A log gets printed with more verbose information if the bitmask 'allows it'.
// This function allows the user to quickly update the entire bitmask.
//
// The value '63' will turn on all verbose logging
func SetLogVerboseBitmask(bitmask uint8) uint8 {
	verboseBitmask = bitmask
	return verboseBitmask
}

// SetTimeFormat allows for changing how the time is printed when a message is logged.
// Default: 2006-01-02 15:04:05
func SetTimeFormat(format string) {
	timeFormat = format
}

// TurnOnAllLogging enables all types of logging messages to go though.
func TurnOnAllLogging() uint8 {
	logBitmask = logFatal | logError | logWarning | logInfo | logDebug | logCustom
	return logBitmask
}

// TurnOnAllLogging enables verbosity for all types of logging messages.
func TurnOnAllVerbose() uint8 {
	verboseBitmask = logFatal | logError | logWarning | logInfo | logDebug | logCustom
	return verboseBitmask
}

// SetLogFatal sets if the Fatal log message will be printed.
// Returns the bitmask after the change.
// Please note that Fatal will run os.Exit(1) regardless of this setting.
func SetLogFatal(b bool) uint8 {
	logToggle(b, logFatal, &logBitmask)
	return logBitmask
}

// SetLogError sets if the Error log message will be printed.
// Returns the bitmask after the change.
func SetLogError(b bool) uint8 {
	logToggle(b, logError, &logBitmask)
	return logBitmask
}

// SetLogWarning sets if the warning log message will be printed.
// Returns the bitmask after the change.
func SetLogWarning(b bool) uint8 {
	logToggle(b, logWarning, &logBitmask)
	return logBitmask
}

// SetLogInfo sets if the information log message will be printed.
// Returns the bitmask after the change.
func SetLogInfo(b bool) uint8 {
	logToggle(b, logInfo, &logBitmask)
	return logBitmask
}

// SetLogDebug sets if the debug log message will be printed.
// Returns the bitmask after the change.
func SetLogDebug(b bool) uint8 {
	logToggle(b, logDebug, &logBitmask)
	return logBitmask
}

// SetLogCustom sets if custom log messages will be printed.
// Returns the bitmask after the change.
func SetLogCustom(b bool) uint8 {
	logToggle(b, logCustom, &logBitmask)
	return logBitmask
}

/* ### Verbose section ### */

// SetLogVerboseFatal sets if the Fatal log message should be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseFatal(b bool) uint8 {
	logToggle(b, logFatal, &verboseBitmask)
	return verboseBitmask
}

// SetLogVerboseError sets if the Error log message will be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseError(b bool) uint8 {
	logToggle(b, logError, &verboseBitmask)
	return verboseBitmask
}

// SetLogVerboseWarning sets if the warning log message will be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseWarning(b bool) uint8 {
	logToggle(b, logWarning, &verboseBitmask)
	return verboseBitmask
}

// SetLogVerboseInfo sets if the information log message will be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseInfo(b bool) uint8 {
	logToggle(b, logInfo, &verboseBitmask)
	return verboseBitmask
}

// SetLogVerboseDebug sets if the debug log message will be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseDebug(b bool) uint8 {
	logToggle(b, logDebug, &verboseBitmask)
	return verboseBitmask
}

// SetLogVerboseCustom sets if custom log messages will be verbose or not
// Returns the bitmask after the change.
func SetLogVerboseCustom(b bool) uint8 {
	logToggle(b, logCustom, &verboseBitmask)
	return verboseBitmask
}

// Toggles the bit in the bitmask depending on if it should be on or off.
func logToggle(b bool, logType uint8, bitmask *uint8) {
	if b {
		*bitmask |= logType
	} else {
		// Clear the logType bit from the LogBitmask
		*bitmask &^= logType
	}
}
