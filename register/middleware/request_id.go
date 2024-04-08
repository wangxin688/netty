package middleware

import (
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const XRequestIDKey string = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.Request.Header.Get(XRequestIDKey)
		if reqID == "" {
			reqID = uuid.NewString()
		}
		c.Set(XRequestIDKey, reqID)
		c.Request.Header.Add(XRequestIDKey, reqID)
		c.Next()
	}
}

func GetLoggerConfig(formatter gin.LogFormatter, output io.Writer, skipPaths []string) gin.LoggerConfig {
	if formatter == nil {
		formatter = XReqIdLogFormatter()
	}

	return gin.LoggerConfig{
		Formatter: formatter,
		Output:    output,
		SkipPaths: skipPaths,
	}
}

func XReqIdLogFormatter() gin.LogFormatter {
	return func(p gin.LogFormatterParams) string {
		return fmt.Sprintf("%s [%s] - \"%s %s %s %d %s\" %s\n",
			p.TimeStamp.Format(time.RFC3339),
			p.Request.Header.Get(XRequestIDKey),
			p.Method,
			p.Path,
			p.Request.Proto,
			p.StatusCode,
			p.Latency,
			p.Request.UserAgent(),
		)
	}
}

func GetRequestID(c *gin.Context) string {
	if reqID, exists := c.Get(XRequestIDKey); exists {
		return reqID.(string)
	}
	return ""
}
