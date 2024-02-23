package main

import (
	lib "github.com/TylerJGabb/my-go-module/lib"
	"github.com/TylerJGabb/new-go-priv-mod/priv"
	gin "github.com/gin-gonic/gin"
)

func main() {
	lib.Hello()
	priv.PrivateFunc()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
