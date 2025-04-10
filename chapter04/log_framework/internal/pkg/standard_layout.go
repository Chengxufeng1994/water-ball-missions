package pkg

import (
	"fmt"
	"time"
)

type StandardLayout struct {
}

var _ Layout = (*StandardLayout)(nil)

func NewStandardLayout() *StandardLayout {
	return &StandardLayout{}
}

func (s *StandardLayout) Format(level Level, name string, msg string) Log {
	format := "%s |-%s %s - %s"

	return Log{
		Message: fmt.Sprintf(format, time.Now().Format("2006-01-02 15:04:05"), level, name, msg),
	}
}
