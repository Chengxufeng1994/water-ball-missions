package serdes

import "fmt"

// SerDesHandler finds a SerDes for a given content type.
type SerDesHandler interface {
	Handle(contentType string) (SerDes, error)
	SetNext(next SerDesHandler)
}

type BaseSerDesHandler struct {
	next        SerDesHandler
	contentType string
	serdes      SerDes
}

var _ SerDesHandler = (*BaseSerDesHandler)(nil)

func NewBaseSerDesHandler(contentType string, serdes SerDes) *BaseSerDesHandler {
	return &BaseSerDesHandler{
		contentType: contentType,
		serdes:      serdes,
	}
}

func (p *BaseSerDesHandler) Handle(contentType string) (SerDes, error) {
	if p.contentType == contentType {
		return p.serdes, nil
	}

	if p.next != nil {
		return p.next.Handle(contentType)
	}

	return nil, fmt.Errorf("no serdes found for content type: %s", contentType)
}

func (p *BaseSerDesHandler) SetNext(next SerDesHandler) {
	p.next = next
}

type JsonSerDesHandler struct {
	contentType string
	serdes      SerDes
	next        SerDesHandler
}

var _ SerDesHandler = (*JsonSerDesHandler)(nil)

// NewJsonSerDesHandler creates a new node for the SerDesProvider chain.
func NewJsonSerDesHandler() *JsonSerDesHandler {
	return &JsonSerDesHandler{
		contentType: "application/json",
		serdes:      NewJsonSerDes(),
	}
}

// Handle searches the chain for a SerDes that can handle the given content type.
func (p *JsonSerDesHandler) Handle(contentType string) (SerDes, error) {
	if p.contentType == contentType {
		return p.serdes, nil
	}

	if p.next != nil {
		return p.next.Handle(contentType)
	}

	return nil, fmt.Errorf("no serdes found for content type: %s", contentType)
}

// SetNext sets the next provider in the chain.
func (p *JsonSerDesHandler) SetNext(next SerDesHandler) {
	p.next = next
}

type XmlSerDesHandler struct {
	contentType string
	serdes      SerDes
	next        SerDesHandler
}

var _ SerDesHandler = (*XmlSerDesHandler)(nil)

// NewXmlSerDesHandler creates a new node for the SerDesProvider chain.
func NewXmlSerDesHandler() *XmlSerDesHandler {
	return &XmlSerDesHandler{
		contentType: "application/xml",
		serdes:      NewXmlSerDes(),
	}
}

// Handle searches the chain for a SerDes that can handle the given content type.
func (p *XmlSerDesHandler) Handle(contentType string) (SerDes, error) {
	if p.contentType == contentType {
		return p.serdes, nil
	}

	if p.next != nil {
		return p.next.Handle(contentType)
	}

	return nil, fmt.Errorf("no serdes found for content type: %s", contentType)
}

// SetNext sets the next provider in the chain.
func (p *XmlSerDesHandler) SetNext(next SerDesHandler) {
	p.next = next
}

// NewSerDesProviderChain creates a chain of SerDesProviders.
func NewSerDesProviderChain(nodes ...SerDesHandler) (SerDesHandler, error) {
	if len(nodes) == 0 {
		return nil, fmt.Errorf("at least one serdes provider node is required")
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].SetNext(nodes[i+1])
	}
	return nodes[0], nil
}
