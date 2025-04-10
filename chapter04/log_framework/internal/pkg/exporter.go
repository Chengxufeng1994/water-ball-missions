package pkg

type Exporter interface {
	Export(log Log) error
}
