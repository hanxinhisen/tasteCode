// Created by Hisen at 2019-07-15.
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

var limiter = rate.NewLimiter(1, 5)

func keyPool() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println(limiter.Allow())
		if limiter.Allow() {
			context.Next()
		} else {
			context.JSON(http.StatusForbidden, gin.H{"status": "1"})
			context.Abort()
		}
	}
}

func main() {
	app := gin.Default()
	app.Use(keyPool())
	app.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status_code": 0})
	})
	app.Run(":8888")
}
