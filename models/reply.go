package models

type Reply struct {
	//gorm.Model
	ID uint `gorm:"column:id" json:"id"`
	// 评论id
	CommentId uint `gorm:"column:comment_id" json:"comment_id"`
	// 父级回复id
	ParentId uint `gorm:"column:parent_id" json:"parent_id"`
	// 来自谁的回复
	FromUserId uint `gorm:"column:from_user_id" json:"from_user_id"`
	// 给谁回复
	ToUserId uint `gorm:"column:to_user_id" json:"to_user_id"`
	// 回复内容
	ReplyContent string `gorm:"column:reply_content" json:"reply_content"`
	// 子级回复
	Children []*Reply `json:"children" gorm:"-"`
}

func (Reply) TableName() string {
	return "n_reply"
}
