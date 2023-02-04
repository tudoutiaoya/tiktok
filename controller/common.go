package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func HandlerErr(context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{})
}
