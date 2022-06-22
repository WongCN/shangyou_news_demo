package news

import (
	"github.com/gin-gonic/gin"
	"shangyou_news_demo/news/service"
)

var news = service.News{}

func GetAllNews(ctx *gin.Context) {
	news.GetAllNews(ctx)
}
