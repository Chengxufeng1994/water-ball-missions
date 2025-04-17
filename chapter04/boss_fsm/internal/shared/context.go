package shared

type Context interface {
	GetValue(key string) any
	SetValue(key string, value any)
	Del(key string)
}
