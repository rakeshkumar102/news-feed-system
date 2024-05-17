package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranay999000/apiGateway/configs"
	"github.com/pranay999000/apiGateway/middleware"
	"github.com/pranay999000/apiGateway/proxies"
	bearertoken "github.com/vence722/gin-middleware-bearer-token"
)

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "APIGateway@v1.0.0.0",
	})
}

func InitRouters(r *gin.Engine) {
	r.GET("/api/v1", middleware.RateLimit, rootHandler)
	r.Any("/api/v1/:service/*proxyPath", middleware.RateLimit, proxies.Services)

	r.Use(bearertoken.Middleware(func (token string, c *gin.Context) bool {
		claims, ok := configs.ValidateToken(token)

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
			})
			return false
		}

		c.Set("email", claims["email"])
		c.Set("name", claims["name"])
		c.Set("id", claims["id"])

		return true
	}))

	r.Any("/api/v2/:service/*proxyPath", middleware.RateLimit, proxies.Services)
}

func main() {
	r := gin.Default()

	InitRouters(r)

	log.Fatal(r.Run(":8000"))
}