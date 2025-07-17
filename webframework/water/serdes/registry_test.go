package serdes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jsonContentType = "application/json"
)

type A struct {
	Name string `json:"name"`
}

func (a *A) Key() string {
	return "A"
}

type B struct {
	Name string `json:"name"`
}

func (b *B) Key() string {
	return "B"
}

func TestRegistry(t *testing.T) {
	// Setup CoR chain
	jsonNode := NewJsonSerDesHandler()
	chain, err := NewSerDesProviderChain(jsonNode)
	assert.NoError(t, err)

	registry := NewRegistry(chain)

	err = registry.Register(&A{})
	assert.NoError(t, err)

	t.Run("Given a registered type with a content type", func(t *testing.T) {
		t.Run("When serialize an object", func(t *testing.T) {
			a := &A{Name: "test"}
			data, err := registry.Serialize(jsonContentType, a)
			assert.NoError(t, err)

			t.Run("Then deserialize it back to an object", func(t *testing.T) {
				v, err := registry.Deserialize(jsonContentType, "A", data)
				assert.NoError(t, err)
				assert.Equal(t, a, v.(*A))
			})
		})

		t.Run("When must serialize an object", func(t *testing.T) {
			a := &A{Name: "test"}
			data := registry.MustSerialize(jsonContentType, a)

			t.Run("Then must deserialize it back to an object", func(t *testing.T) {
				v := registry.MustDeserialize(jsonContentType, "A", data)
				assert.Equal(t, a, v.(*A))
			})
		})

		t.Run("When build an object", func(t *testing.T) {
			v, err := registry.Build("A")
			assert.NoError(t, err)
			assert.IsType(t, &A{}, v)
		})

		t.Run("When must build an object", func(t *testing.T) {
			v := registry.MustBuild("A")
			assert.IsType(t, &A{}, v)
		})
	})

	t.Run("Given an unregistered key", func(t *testing.T) {
		t.Run("When deserialize an object", func(t *testing.T) {
			a := &A{Name: "test"}
			data, _ := registry.Serialize(jsonContentType, a)
			_, err := registry.Deserialize(jsonContentType, "B", data)
			assert.Error(t, err)
		})
	})

	t.Run("Given an unregistered content type", func(t *testing.T) {
		t.Run("When serialize an object", func(t *testing.T) {
			_, err := registry.Serialize("application/xml", &A{})
			assert.Error(t, err)
		})
	})

	t.Run("Given a registered key", func(t *testing.T) {
		t.Run("When register it again", func(t *testing.T) {
			err := registry.Register(&A{})
			assert.Error(t, err)
		})
	})
}

func TestRegistry_Register(t *testing.T) {
	chain, _ := NewSerDesProviderChain(NewJsonSerDesHandler())
	registry := NewRegistry(chain)

	t.Run("Given a registerable type", func(t *testing.T) {
		t.Run("When register it", func(t *testing.T) {
			registerable := &B{Name: "test"}
			err := registry.Register(registerable)
			assert.NoError(t, err)

			t.Run("Then it should be registered", func(t *testing.T) {
				v, err := registry.Build(registerable.Key())
				assert.NoError(t, err)
				assert.IsType(t, &B{}, v)
			})
		})
	})

	t.Run("Given a non-pointer registerable type", func(t *testing.T) {
		t.Run("When register it", func(t *testing.T) {
			var registerable NonPointerRegisterable
			err := registry.Register(registerable)
			assert.Error(t, err)
		})
	})
}

type NonPointerRegisterable struct{}

func (np NonPointerRegisterable) Key() string {
	return "NonPointerRegisterable"
}
