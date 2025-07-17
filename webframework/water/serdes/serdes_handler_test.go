package serdes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// DummyXmlSerDes is a dummy SerDes for testing purposes.
type DummyXmlSerDes struct{}

func (x *DummyXmlSerDes) Deserialize(d []byte, v any) error { return nil }
func (x *DummyXmlSerDes) Serialize(v any) ([]byte, error)   { return nil, nil }

func TestSerDesProviderChain(t *testing.T) {
	jsonNode := NewJsonSerDesHandler()
	xmlNode := NewBaseSerDesHandler("application/xml", &DummyXmlSerDes{})

	chain, err := NewSerDesProviderChain(jsonNode, xmlNode)
	assert.NoError(t, err)

	t.Run("Find JSON serdes", func(t *testing.T) {
		s, err := chain.Handle("application/json")
		assert.NoError(t, err)
		assert.IsType(t, &JsonSerDes{}, s)
	})

	t.Run("Find XML serdes", func(t *testing.T) {
		s, err := chain.Handle("application/xml")
		assert.NoError(t, err)
		assert.IsType(t, &DummyXmlSerDes{}, s)
	})

	t.Run("Find unsupported serdes", func(t *testing.T) {
		_, err := chain.Handle("application/yaml")
		assert.Error(t, err)
	})

	t.Run("Create chain with no nodes", func(t *testing.T) {
		_, err := NewSerDesProviderChain()
		assert.Error(t, err)
	})
}
