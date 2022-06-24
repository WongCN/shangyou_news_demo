package vo

import "shangyou_news_demo/models"

type VoNewsTypeDetail struct {
	//gorm.Model
	ID uint
	// 新闻标题
	Title string `gorm:"column:title" json:"title"`
	// 新闻正文
	Text string `gorm:"column:text" json:"text"`
	// 新闻上传者id
	UploaderId     uint        `gorm:"column:uploader_id" json:"-"`
	UploaderDetail models.User `json:"-" gorm:"references:UploaderId;foreignkey:ID"`
	UploaderName   string      `gorm:"-" json:"uploader_name"`
	// 新闻类型id
	TypeId uint8 `gorm:"column:type_id" json:"type_id"`
	// 点赞数
	Likes uint64 `gorm:"column:likes" json:"likes"`
	// 浏览量
	PageViews uint64 `gorm:"column:page_views" json:"page_views"`
	// 关联新闻类型表
	Type     models.NewsType `json:"-" gorm:"references:TypeId;foreignkey:ID"`
	TypeName string          `gorm:"-" json:"type_name"`
}

func (VoNewsTypeDetail) TableName() string {
	return "n_news"
}
