package rest

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LogHandler struct{}

func (l *LogHandler) Router(r *gin.Engine) {
	r.POST("/Log/Error", l.AddErrorLog)
}

func (l *LogHandler) AddErrorLog(c *gin.Context) {
	zap.L().Error(
		"test error log",
	)
}
