package news

import (
	"github.com/gin-gonic/gin"
	"shangyou_news_demo/news/service"
)

var news = service.News{}

func GetAllNewsById(ctx *gin.Context) {
	news.GetAllNewsById(ctx)
}
func GetNewsDetailById(ctx *gin.Context) {
	news.GetNewsDetailById(ctx)
}
