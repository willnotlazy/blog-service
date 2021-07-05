package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct {

}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {

}
func (a Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError.WithDetails([]string{"这个错了", "那个错了", "所有的都错了"}...))
	return
}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}