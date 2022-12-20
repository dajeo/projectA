package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"projectA/docs"
	"projectA/gateway"
	"projectA/models"
	"projectA/routes/auth"
	"projectA/routes/customers"
	"projectA/utils"
	"time"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		return
	}

	// Initializing databases
	utils.InitDb()
	utils.InitRedis()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	// See this before deploy to production
	// https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	proxyErr := r.SetTrustedProxies(nil)
	if proxyErr != nil {
		return
	}

	// todo: use database
	authMiddleware, authErr := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "zone",
		Key:         []byte("secret"), // todo: generate secret key
		Timeout:     time.Hour * 24 * 30,
		MaxRefresh:  time.Hour,
		IdentityKey: "tel",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.UserRes); ok {
				return jwt.MapClaims{
					"tel": v.Tel,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.UserRes{
				Tel: claims["tel"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			userTel := loginVals.Tel
			userPass := loginVals.Password

			if userTel == "908" && userPass == "admin" {
				return &models.UserRes{
					Tel: userTel,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.UserRes); ok && v.Tel == "908" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if authErr != nil {
		return
	}

	authErrInit := authMiddleware.MiddlewareInit()
	if authErrInit != nil {
		return
	}

	// Initializing routes
	auth.Router(r)
	customers.Router(r, authMiddleware)

	// Initialize gateway
	gateway.Router(r)

	r.POST("/login", authMiddleware.LoginHandler)

	r.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
	r.Use(cors.Default()) // Configure before production
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
