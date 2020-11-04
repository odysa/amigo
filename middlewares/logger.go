package middlewares

import (
	"github.com/odysa/amigo/lib"
	"log"
	"time"
)

func Logger() lib.HandlerFunc {
	return func(c *lib.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}
