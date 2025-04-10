package pkg

type Layout interface {
	Format(level Level, name string, msg string) Log
}
