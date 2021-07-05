package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/email"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defailMailer := email.NewEmail(&email.SMTPInfo{
			Host: global.EmailSetting.Host,
			Post: global.EmailSetting.Port,
			IsSSL: global.EmailSetting.IsSSL,
			UserName: global.EmailSetting.UserName,
			Password: global.EmailSetting.Password,
			From: global.EmailSetting.From,
		})
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Errorf("panic recover err %v", err)
				err := defailMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出,发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
					)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
