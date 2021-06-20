package declaration

const (
	Off = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
	All
)

var PrecencyName = []string{
	Off:   "OFF",
	Fatal: "FATAL",
	Error: "ERROR",
	Warn:  "WARN",
	Info:  "INFO",
	Debug: "DEBUG",
	Trace: "TRACE",
	All:   "ALL",
}
