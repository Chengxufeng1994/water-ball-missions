package pkg

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type FileExporter struct {
	ch     chan string
	writer *bufio.Writer
	mu     sync.Mutex
}

func NewFileExporter(filepath string) *FileExporter {
	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	bufWriter := bufio.NewWriterSize(f, 4096)

	exporter := &FileExporter{
		ch:     make(chan string, 1024),
		writer: bufWriter,
	}

	go exporter.tick()

	return exporter
}

func (f *FileExporter) tick() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			f.mu.Lock()
			f.writer.Flush()
			f.mu.Unlock()
		case msg := <-f.ch:
			f.mu.Lock()
			f.writer.WriteString(msg)
			f.writer.WriteString("\n")
			if f.writer.Buffered() >= f.writer.Size() {
				f.writer.Flush()
			}
			f.mu.Unlock()
		}
	}
}

func (f *FileExporter) Export(log Log) error {
	select {
	case f.ch <- log.Message:
		return nil
	default:
		return fmt.Errorf("channel is full")
	}
}
