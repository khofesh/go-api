package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

// SecureFunc : custom middleware
func SecureFunc() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedHostsAreRegex:  false,
		SSLRedirect:           false,
		SSLTemporaryRedirect:  false,
		SSLHost:               "",
		SSLProxyHeaders:       map[string]string{},
		STSSeconds:            0,
		STSIncludeSubdomains:  false,
		STSPreload:            false,
		ForceSTSHeader:        false,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'",
		ReferrerPolicy:        "same-origin",
		FeaturePolicy:         "vibrate 'none';",
		IsDevelopment:         true,
	})
	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			c.Abort()
			return
		}

		// Avoid header rewrite if response is a redirection.
		if status := c.Writer.Status(); status > 300 && status < 399 {
			c.Abort()
		}

		c.Next()
	}
}
