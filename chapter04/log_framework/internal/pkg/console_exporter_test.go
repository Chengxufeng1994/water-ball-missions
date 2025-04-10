package pkg

import (
	"testing"
)

func TestConsoleExporter(t *testing.T) {
	NewConsoleExporter().Export(Log{Message: "Hello World!"})
}
