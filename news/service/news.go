package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"shangyou_news_demo/initialize"
	"shangyou_news_demo/models"
	"shangyou_news_demo/models/vo"
	"strconv"
	"time"
)

var DBB = initialize.DB.Begin()
var DB = initialize.DB
var Rdb = initialize.Rdb

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

func (news News) GetNewsDetailById(ctx *gin.Context) {
	query := ctx.Query("id")
	detail := vo.VoNewsTypeDetail{}
	DB.Where("id = ?", query).Preload("Type").Preload("UploaderDetail").Find(&detail)
	detail.UploaderName = detail.UploaderDetail.Username
	detail.TypeName = detail.Type.NewsType
	ctx.JSON(200, detail)
}

func (news News) GetAllNewsById(ctx *gin.Context) {
	// 获取url中的id参数
	query := ctx.Query("id")
	// 在redis中查询对应的id是否有缓存数据
	get := Rdb.Get("News::ID::" + query)
	// 有缓存在缓存中拿
	if get.Val() != "" {
		fmt.Println("读取的redis")
		n2 := models.News{}
		jsoniter.Unmarshal([]byte(get.Val()), &n2)
		ctx.JSON(200, n2)
	} else {
		// 没有缓存就查询数据库
		n := models.News{}
		// 连接查询
		find := DB.Where("id = ?", query).Preload("Type").Preload("Picture").Preload("Comment").Find(&n)
		// 如果查询数据库没有对应的数据,就往redis中存入空信息.防止穿透
		if find.RowsAffected < int64(1) {
			n.Error = "查询的id不存在,请重新输入"
			marshal, _ := jsoniter.Marshal(n)
			Rdb.Set("News::ID::"+query, marshal, time.Second*3)
			ctx.JSON(400, n)
			return
		}
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
				}
			}
			comment[i].TestMap = m2
		}
		marshal, _ := jsoniter.Marshal(n)
		i2, i3 := getRand(20, 60)
		// 取一个随机时间,作为缓存过期时间
		var cacheTime = i2*60 + i3
		Rdb.Set("News::ID::"+strconv.Itoa(int(n.ID)), marshal, time.Duration(cacheTime)*time.Second)
		fmt.Println("读取的数据库")
		ctx.JSON(200, n)
		return
	}
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

/**
获取随机值
*/
func getRand(m int, s int) (int64, int64) {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(m)), int64(rand.Intn(s))
}
