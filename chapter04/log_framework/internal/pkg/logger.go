package pkg

type Logger interface {
	Trace(msg string)
	Info(msg string)
	Debug(msg string)
	Warn(msg string)
	Error(msg string)

	GetLevel() Level
	GetLayout() Layout
	GetExporter() Exporter
}
