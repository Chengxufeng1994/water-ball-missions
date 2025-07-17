package water

type HttpHandlerFunc func(*Context)

func (f HttpHandlerFunc) Handle(ctx *Context) {
	f(ctx)
}

var _ HttpHandler = HttpHandlerFunc(nil)

type HttpHandler interface {
	Handle(ctx *Context)
}
