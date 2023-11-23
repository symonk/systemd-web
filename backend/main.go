package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var uptime time.Time

func init() {
	uptime = time.Now()
}

func main() {
	engine := gin.Default()
	engine.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"uptime": fmt.Sprintf("%d seconds", int(time.Since(uptime).Seconds())),
		})
	})
	engine.Run()
}
