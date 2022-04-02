package middlewares

import "github.com/gin-gonic/gin"

func JsonResponseHeader(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	context.Next()
}
