package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/api"
)

func main() {
	r := gin.Default()

	api.InitRouter(r)

	r.Run(":3333")
}
