package shared

type Context interface {
	GetPrefix() string
	GetValue(key string) any
	SetValue(key string, value any)
	Del(key string)
}
