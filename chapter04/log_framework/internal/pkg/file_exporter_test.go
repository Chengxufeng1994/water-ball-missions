package pkg

import (
	"os"
	"testing"
	"time"
)

func TestFileExporter(t *testing.T) {
	exporter := NewFileExporter("test.log")

	exporter.Export(Log{Message: "Hello World!"})
	exporter.Export(Log{Message: "Hello World!"})
	exporter.Export(Log{Message: "Hello World!"})

	time.Sleep(1 * time.Second)

	os.Remove("test.log")
}
