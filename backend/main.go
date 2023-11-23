package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	V1 = "/v1"
)

var uptime time.Time

func init() {
	uptime = time.Now()
}

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uptime": fmt.Sprintf("%d seconds", int(time.Since(uptime).Seconds())),
		})
	})
	registerVersionOne(router)
	router.Run()
}

// Implement API version 1 capabilities.
func registerVersionOne(router *gin.Engine) *gin.Engine {
	apiV1 := router.Group(V1)
	apiV1.GET("/envs", func(c *gin.Context) {
		c.JSON(http.StatusOK, "return env list later")
	})
	return router
}
