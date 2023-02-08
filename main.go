package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/api"
)

type User struct {
	Name       string
	Age        int
	UserExtend string
}

type Friend struct {
	Name         string
	Age          int
	FriendExtend string
}

func main() {
	r := gin.Default()

	api.InitRouter(r)

	r.Run(":3333")
}
