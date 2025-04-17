package logger

import (
	"github.com/Chengxufeng1994/water-ball-missions/chapter04/log_framework/internal/pkg"
)

type Logger struct {
	Name     string
	parent   pkg.Logger
	Level    pkg.Level
	Layout   pkg.Layout
	Exporter pkg.Exporter
}

var _ pkg.Logger = (*Logger)(nil)

func NewRoot(name string, level pkg.Level, layout pkg.Layout, exporter pkg.Exporter) *Logger {
	return &Logger{
		parent:   nil,
		Name:     name,
		Level:    level,
		Layout:   layout,
		Exporter: exporter,
	}
}

func New(parent pkg.Logger, name string, level pkg.Level, layout pkg.Layout, exporter pkg.Exporter) *Logger {
	if level == -1 {
		level = parent.GetLevel()
	}

	if layout == nil {
		layout = parent.GetLayout()
	}

	if exporter == nil {
		exporter = parent.GetExporter()
	}

	return &Logger{
		parent:   parent,
		Name:     name,
		Level:    level,
		Layout:   layout,
		Exporter: exporter,
	}
}

func (l *Logger) GetName() string {
	return l.Name
}

func (l *Logger) GetLevel() pkg.Level {
	return l.Level
}

func (l *Logger) GetLayout() pkg.Layout {
	return l.Layout
}

func (l *Logger) GetExporter() pkg.Exporter {
	return l.Exporter
}

func (l *Logger) log(level pkg.Level, msg string) {
	if l.Level > level {
		return
	}
	log := l.Layout.Format(level, l.Name, msg)
	_ = l.Exporter.Export(log)
}

// Trace implements domain.Logger.
func (l *Logger) Trace(msg string) {
	l.log(pkg.LevelTrace, msg)
}

// Info implements domain.Logger.
func (l *Logger) Info(msg string) {
	l.log(pkg.LevelInfo, msg)
}

// Debug implements domain.Logger.
func (l *Logger) Debug(msg string) {
	l.log(pkg.LevelDebug, msg)
}

// Warn implements domain.Logger.
func (l *Logger) Warn(msg string) {
	l.log(pkg.LevelWarn, msg)
}

// Error implements domain.Logger.
func (l *Logger) Error(msg string) {
	l.log(pkg.LevelError, msg)
}
