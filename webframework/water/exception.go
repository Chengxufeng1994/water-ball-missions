package water

import (
	"slices"
)

type ExceptionHandler interface {
	SetNext(handler ExceptionHandler) ExceptionHandler
	Handle(c *Context)
}

// NewDefaultExceptionHandler creates and chains the default exception handlers.
// The chain is: pathNotFoundHandler -> methodNotAllowedHandler -> internalServerErrorHandler
func NewDefaultExceptionHandler(router *router) ExceptionHandler {
	// The handlers that form the chain
	pathNotFoundH := &pathNotFoundHandler{router: router}
	methodNotAllowedH := &methodNotAllowedHandler{router: router}
	internalServerErrorH := &internalServerErrorHandler{}

	// Link the handlers in the correct order
	pathNotFoundH.SetNext(methodNotAllowedH)
	methodNotAllowedH.SetNext(internalServerErrorH)

	// The root of the chain is the first handler
	return pathNotFoundH
}

// internalServerErrorHandler handles uncaught exceptions and returns a 500 Internal Server Error.
// It is the last handler in the chain.
type internalServerErrorHandler struct {
	next ExceptionHandler
}

func (h *internalServerErrorHandler) SetNext(handler ExceptionHandler) ExceptionHandler {
	h.next = handler
	return handler
}

func (h *internalServerErrorHandler) Handle(c *Context) {
	// If the context has reached this point, it means an unhandled error occurred.
	// We default to an internal server error.
	c.Next()

	if c.Err != nil {
		c.String(int(StatusCodeInternalServerError), "%v", c.Err.Error())
		return
	}
}

// pathNotFoundHandler handles requests for which no route can be found.
// It checks if the path exists for any HTTP method. If not, it returns a 404.
// If the path exists for other methods, it passes the request to the next handler.
type pathNotFoundHandler struct {
	next   ExceptionHandler
	router *router
}

func (h *pathNotFoundHandler) SetNext(handler ExceptionHandler) ExceptionHandler {
	h.next = handler
	return handler
}

func (h *pathNotFoundHandler) Handle(c *Context) {
	// Check if a route exists for the given path under any method.
	for method := range h.router.roots {
		if n, _ := h.router.getRoute(Method(method), c.HttpRequest.Req.URL.Path); n != nil {
			// The path is valid, but the method is not. Pass to the next handler.
			h.next.Handle(c)
			return
		}
	}
	// If we've checked all methods and found no matching path, it's a 404 Not Found.
	c.String(
		int(StatusCodeBadRequest),
		`The method "%s" is not allowed for path "<%s>"`, c.HttpRequest.Req.Method, c.HttpRequest.Req.URL.Path,
	)
}

// methodNotAllowedHandler handles requests where a route exists for the path, but not for the requested HTTP method.
// It is called only when pathNotFoundHandler has confirmed the path's existence.
type methodNotAllowedHandler struct {
	next   ExceptionHandler
	router *router
}

func (h *methodNotAllowedHandler) SetNext(handler ExceptionHandler) ExceptionHandler {
	h.next = handler
	return handler
}

func (h *methodNotAllowedHandler) Handle(c *Context) {
	// If the path exists, but the method is not allowed, return a 405 Method Not Allowed.
	path := c.HttpRequest.Req.URL.Path
	allowedMethods := []string{}
	for method := range h.router.roots {
		if n, _ := h.router.getRoute(Method(method), path); n != nil {
			allowedMethods = append(allowedMethods, method)
		}
	}

	if slices.Contains(allowedMethods, string(c.HttpRequest.Req.Method)) {
		h.next.Handle(c)
		return
	}

	c.String(
		int(StatusCodeMethodNot),
		`The method "%s" is not allowed on the path "<%s>"`, c.HttpRequest.Req.Method, c.HttpRequest.Req.URL.Path,
	)
}
