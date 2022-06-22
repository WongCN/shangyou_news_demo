package service

import (
	"github.com/gin-gonic/gin"
	"shangyou_news_demo/initialize"
	"shangyou_news_demo/models"
)

var DB = initialize.DB.Begin()

type News struct {
}

//func (news News) AddNews(ctx *gin.Context) {
//	n := models.News{}
//	ctx.ShouldBindJSON(n)
//	err := DB.Create(&n).Error
//	if err != nil {
//		DB.Rollback()
//	}
//	DB.Commit()
//}

func (news News) GetAllNews(ctx *gin.Context) {
	n := models.News{}
	// 三表连接查询
	DB.Preload("Type").Preload("Picture").Preload("Comment").Find(&n)
	// 获取评论字段
	comment := n.Comment
	// 获取回复信息
	son := getSon(0)
	// 将回复信息填充到评论中
	for i := 0; i < len(comment); i++ {
		m2 := map[int]interface{}{}
		for j := 0; j < len(son); j++ {
			if comment[i].ID == son[j].CommentId {
				// 回复者的id作为key
				m2[int(son[j].FromUserId)] = son[j]
				continue
			}
		}
		comment[i].TestMap = m2
	}
	ctx.JSON(200, n)
}

/**
递归获取回复信息
*/
func getSon(pid uint) []*models.Reply {
	var reply []models.Reply
	DB.Where("parent_id = ?", pid).Find(&reply)
	ReplyList := []*models.Reply{}
	for _, v := range reply {
		child := getSon(v.ID)
		node := &models.Reply{
			ID:           v.ID,
			CommentId:    v.CommentId,
			ParentId:     v.ParentId,
			FromUserId:   v.FromUserId,
			ToUserId:     v.ToUserId,
			ReplyContent: v.ReplyContent,
		}
		node.Children = child
		ReplyList = append(ReplyList, node)
	}
	return ReplyList
}
