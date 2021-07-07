package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
	"time"

	//"github.com/go-programming-tour-book/blog-service/global"
	"net/http"

	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	"github.com/go-programming-tour-book/blog-service/internal/routers/api"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter()


func NewRouter() *gin.Engine {
	r := gin.New()

	methodLimiters.AddBuckets(limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use()
	}

	r.Use(middleware.Tracing())
	r.Use(middleware.AccessLog())
	r.Use(middleware.Translations())
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.ServerSetting.DefaultContextTimeout))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()

	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())

	apiv1.POST("/tags", tag.Create)
	apiv1.DELETE("/tags/:id", tag.Delete)
	apiv1.PUT("/tags/:id", tag.Update)
	apiv1.GET("/tags/:id", tag.Get)
	apiv1.GET("/tags", tag.List)
	apiv1.PATCH("/tags/:id", tag.Update)

	apiv1.POST("/articles", article.Create)
	apiv1.DELETE("/articles/:id", article.Delete)
	apiv1.PUT("/articles/:id", article.Update)
	apiv1.GET("/articles/:id", article.Get)
	apiv1.GET("/articles", article.List)
	apiv1.PATCH("/articles/:id", article.Update)

	return r
}
