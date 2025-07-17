package water

import (
	"log"
	"time"
)

func Logger() HttpHandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.HttpResponse.StatusCode, c.HttpRequest.Req.RequestURI, time.Since(t))
	}
}
