package middlewares

import (
	"github.com/odysa/amigo"
	"log"
	"time"
)

func Logger() amigo.HandlerFunc {
	return func(c *amigo.Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}
