package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

)

func AddLoggerToContext(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("logger", logger)
		ctx.Next()
	}
}
