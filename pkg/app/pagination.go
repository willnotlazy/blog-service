package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	page := convert.Strto(c.Query("page")).MustInt()
	if page < 0 {
		page = 1
	}

	return page
}

func GetPageSize(c *gin.Context) int {
	pagesize := convert.Strto(c.Query("page_size")).MustInt()
	if pagesize < 0 {
		return global.AppSetting.DefaultPageSize
	}

	if pagesize > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}

	return pagesize
}

func GetPageOffset(page, pagesize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pagesize
	}

	return result
}
