package water

import (
	"strings"
)

type (
	Router interface {
		addRoute(method Method, path string, handler HttpHandlerFunc)
		handle(ctx *Context)
	}

	router struct {
		// roots key is HTTP method
		roots map[string]*node
	}

	node struct {
		pattern  string // The route pattern e.g. /p/:lang
		part     string // The part of the URL that this node represents
		children map[string]*node
		isWild   bool // isWild is true if part contains a ':' or '*'
		handler  HttpHandlerFunc
	}
)

var _ Router = (*router)(nil)

func newRouter() *router {
	return &router{
		roots: make(map[string]*node),
	}
}

var _ Router = (*router)(nil)

// parsePattern splits the URL pattern into parts.
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method Method, pattern string, handler HttpHandlerFunc) {
	m := string(method)
	parts := parsePattern(pattern)

	if _, ok := r.roots[m]; !ok {
		r.roots[m] = &node{}
	}
	root := r.roots[m]
	root.insert(pattern, parts, 0, handler)
}

func (n *node) insert(pattern string, parts []string, height int, handler HttpHandlerFunc) {
	if len(parts) == height {
		n.pattern = pattern
		n.handler = handler
		return
	}

	part := parts[height]
	child, ok := n.children[part]
	if !ok {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		if n.children == nil {
			n.children = make(map[string]*node)
		}
		n.children[part] = child
	}
	child.insert(pattern, parts, height+1, handler)
}

func (r *router) getRoute(method Method, path string) (*node, map[string]string) {
	m := string(method)
	searchParts := parsePattern(path)
	params := make(map[string]string)

	root, ok := r.roots[m]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.handler == nil {
			return nil
		}
		return n
	}

	part := parts[height]

	// Prioritize static match
	if child, ok := n.children[part]; ok {
		if result := child.search(parts, height+1); result != nil {
			return result
		}
	}

	// Then wildcard match
	for _, child := range n.children {
		if child.isWild {
			if result := child.search(parts, height+1); result != nil {
				return result
			}
		}
	}

	return nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(Method(c.HttpRequest.Req.Method), c.HttpRequest.Req.URL.Path)

	if n != nil {
		c.HttpRequest.Params = params
		c.handlers = append(c.handlers, n.handler)
	} else {
		// Not found
	}

	c.Next()
}
