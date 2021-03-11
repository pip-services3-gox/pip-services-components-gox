package log

/*
Standard log levels.

Logs at debug and trace levels are usually captured only locally for troubleshooting and never sent to consolidated log services.

None  = 0 Nothing to log
Fatal = 1 Log only fatal errors that cause processes to crash
Error = 2 Log all errors.
Warn  = 3 Log errors and warnings
Info  = 4 Log errors and important information messages
Debug = 5 Log everything except traces
Trace = 6 Log everything.
*/
const (
	None  = 0
	Fatal = 1
	Error = 2
	Warn  = 3
	Info  = 4
	Debug = 5
	Trace = 6
)
