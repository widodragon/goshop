package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s %s %s %d [%s]",
			params.ClientIP,
			params.Method,
			params.Path,
			params.StatusCode,
			params.TimeStamp.Format(time.RFC3339),
		)
	})
}
