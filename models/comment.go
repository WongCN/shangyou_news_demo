package models

type Comment struct {
	ID uint `gorm:"column:id" json:"id"`
	// 新闻id
	NewsId uint `gorm:"column:news_id" json:"news_id"`
	// 评论内容
	Content string `gorm:"column:content" json:"content"`
	// 评论者id
	UserId uint `gorm:"column:user_id" json:"user_id"`
	// 评论下面的回复
	Reply   []*Reply            `json:"reply" gorm:"foreignkey:CommentId;references:NewsId"`
	TestMap map[int]interface{} `json:"TestMap" gorm:"-"`
}

func (Comment) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "n_comment"
}
