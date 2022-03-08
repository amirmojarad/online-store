package log

import (
	"github.com/gin-gonic/gin"
)

func CustomLog() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		defer ctx.Next()

	}
}
