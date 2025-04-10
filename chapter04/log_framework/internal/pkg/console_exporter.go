package pkg

import (
	"fmt"
	"os"
)

type ConsoleExporter struct {
}

func NewConsoleExporter() *ConsoleExporter {
	return &ConsoleExporter{}
}

func (c *ConsoleExporter) Export(log Log) error {
	_, err := fmt.Fprintln(os.Stdout, log.Message)
	return err
}
