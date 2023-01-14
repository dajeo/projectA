package server

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
)

func Init(authMiddleware *jwt.GinJWTMiddleware) {
	r := NewRouter(authMiddleware)

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	r.Use(cors.Default()) // Configure before production

	err := r.Run(":8000")
	if err != nil {
		return
	}
}
