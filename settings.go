package logger

type LoggingLevel int

const (
	ErrorLevel LoggingLevel = iota
	WarningLevel
	InfoLevel
	DebugLevel
)
