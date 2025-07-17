package serdes

import (
	"fmt"
	"reflect"
)

type (
	// Registerable defines the interface for types that can be registered with the registry.
	Registerable interface {
		Key() string
	}

	// Registry provides a mechanism for building types and serializing/deserializing them.
	Registry interface {
		Serialize(contentType string, v any) ([]byte, error)
		MustSerialize(contentType string, v any) []byte
		Build(key string) (any, error)
		MustBuild(key string) any
		Deserialize(contentType, key string, data []byte) (any, error)
		MustDeserialize(contentType, key string, data []byte) any
		Register(r Registerable) error
	}

	registry struct {
		factories   map[string]func() any
		serdesChain SerDesHandler
	}
)

// NewRegistry creates a new Registry with the given SerDesProvider.
func NewRegistry(serdesChain SerDesHandler) Registry {
	return &registry{
		factories:   make(map[string]func() any),
		serdesChain: serdesChain,
	}
}

// Register adds a new type to the registry.
// The registerable item must be a pointer.
func (r *registry) Register(registerable Registerable) error {
	key := registerable.Key()
	if _, ok := r.factories[key]; ok {
		return fmt.Errorf("key already registered: %s", key)
	}
	t := reflect.TypeOf(registerable)
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("registerable must be a pointer, but got %s", t.Kind())
	}
	elem := t.Elem()
	r.factories[key] = func() any {
		return reflect.New(elem).Interface()
	}
	return nil
}

// Serialize converts a value to bytes using the appropriate SerDes.
func (r *registry) Serialize(contentType string, v any) ([]byte, error) {
	serdes, err := r.serdesChain.Handle(contentType)
	if err != nil {
		return nil, err
	}
	return serdes.Serialize(v)
}

// MustSerialize is like Serialize but panics on error.
func (r *registry) MustSerialize(contentType string, v any) []byte {
	data, err := r.Serialize(contentType, v)
	if err != nil {
		panic(err)
	}
	return data
}

// Build creates an instance of a registered type.
func (r *registry) Build(key string) (any, error) {
	factory, ok := r.factories[key]
	if !ok {
		return nil, fmt.Errorf("key not found: %s", key)
	}
	return factory(), nil
}

// MustBuild is like Build but panics on error.
func (r *registry) MustBuild(key string) any {
	v, err := r.Build(key)
	if err != nil {
		panic(err)
	}
	return v
}

// Deserialize converts bytes to a new instance of a registered type.
func (r *registry) Deserialize(contentType, key string, data []byte) (any, error) {
	serdes, err := r.serdesChain.Handle(contentType)
	if err != nil {
		return nil, err
	}
	v, err := r.Build(key)
	if err != nil {
		return nil, err
	}
	err = serdes.Deserialize(data, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// MustDeserialize is like Deserialize but panics on error.
func (r *registry) MustDeserialize(contentType, key string, data []byte) any {
	v, err := r.Deserialize(contentType, key, data)
	if err != nil {
		panic(err)
	}
	return v
}
