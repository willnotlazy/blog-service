package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, nil
	}

	return w.ResponseWriter.Write(p)
}


func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"url" : c.Request.URL.String(),
			"request" : c.Request.PostForm.Encode(),
			"response" : bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d", c.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}