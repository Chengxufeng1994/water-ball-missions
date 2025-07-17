package water

import (
	"net"
	"net/http"

	"github.com/Chengxufeng1994/water-ball-missions/webframework/water/serdes"
)

type Engine struct {
	listener         net.Listener
	server           *http.Server
	router           *router
	middlewares      []HttpHandlerFunc
	registry         serdes.Registry
	serdesChain      serdes.SerDesHandler
	exceptionHandler HttpHandlerFunc
}

var _ http.Handler = (*Engine)(nil)

func New(port string) (*Engine, error) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	jsonNode := serdes.NewJsonSerDesHandler()
	serdesChain, err := serdes.NewSerDesProviderChain(jsonNode)
	if err != nil {
		ln.Close()
		return nil, err
	}

	registry := serdes.NewRegistry(serdesChain)

	e := &Engine{
		listener:    ln,
		router:      newRouter(),
		registry:    registry,
		serdesChain: serdesChain,
	}

	e.middlewares = []HttpHandlerFunc{
		Logger(),
		NewDefaultExceptionHandler(e.router).Handle,
	}

	e = e.init()

	return e, nil
}

func (e *Engine) init() *Engine {
	e.server = &http.Server{
		Handler:      e,
		ReadTimeout:  10 * 1000000000, // 10 seconds
		WriteTimeout: 10 * 1000000000, // 10 seconds
		IdleTimeout:  60 * 1000000000, // 60 seconds
	}
	return e
}

func (e *Engine) GetRouter() *router {
	return e.router
}

func (e *Engine) GetRegistry() serdes.Registry {
	return e.registry
}

func (e *Engine) AddRoute(method Method, path string, handler HttpHandlerFunc) {
	e.router.addRoute(method, path, handler)
}

func (e *Engine) AddSerDesPlugin(serdesPlugin serdes.SerDesHandler) {
	// This will replace the next plugin, be careful if called multiple times
	e.serdesChain.SetNext(serdesPlugin)
}

func (e *Engine) AddMiddleware(mw HttpHandlerFunc) {
	e.middlewares = append(e.middlewares, mw)
}

func (e *Engine) SetExceptionHandler(handler HttpHandlerFunc) {
	e.exceptionHandler = handler
	e.middlewares = []HttpHandlerFunc{
		Logger(),
		e.exceptionHandler,
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := &HttpRequest{Req: r}
	response := &HttpResponse{Writer: w}
	ctx := newContext(e.registry, request, response)

	ctx.handlers = append([]HttpHandlerFunc{}, e.middlewares...)

	e.router.handle(ctx)
}

func (e *Engine) Launch() error {
	return e.server.Serve(e.listener)
}
