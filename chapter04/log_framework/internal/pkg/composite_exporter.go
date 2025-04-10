package pkg

import "fmt"

type CompositeExporter struct {
	exporters []Exporter
}

var _ Exporter = (*CompositeExporter)(nil)

func NewCompositeExporter(exporters ...Exporter) *CompositeExporter {
	return &CompositeExporter{exporters: exporters}
}

func (c *CompositeExporter) Export(log Log) error {
	for _, exp := range c.exporters {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println(r)
			}
		}()
		_ = exp.Export(log)
	}
	return nil
}
