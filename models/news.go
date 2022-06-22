package models

type News struct {
	//gorm.Model
	ID uint
	// gorm框架自带的 CreatedAt UpdatedAt DeletedAt   关于时间的字段一般存储时间戳,然后解析成 yyyy-MM-DD HH:mm:ss 的格式
	//CreatedAt time.Time      `json:"createdAt"`
	//UpdatedAt time.Time      `json:"-"`
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	// 新闻标题
	Title string `gorm:"column:title" json:"title"`
	// 新闻正文
	Text string `gorm:"column:text" json:"text"`
	// 新闻上传者id
	UploaderId uint `gorm:"column:uploader_id" json:"uploader_id"`
	// 新闻类型id
	TypeId uint8 `gorm:"column:type_id" json:"type_id"`
	// 点赞数
	Likes uint64 `gorm:"column:likes" json:"likes"`
	// 浏览量
	PageViews uint64 `gorm:"column:page_views" json:"page_views"`
	// 状态 0待审核 1可见 2不可见
	Status uint8 `json:"status" gorm:"column:status"`
	// 关联新闻类型表
	Type NewsType `json:"type" gorm:"references:TypeId;foreignkey:ID"`
	// 关联新闻图片表
	Picture []Picture `json:"picture" gorm:"foreignkey:NewsId;references:ID"`
	// 关联评论表
	Comment []Comment `json:"comment" gorm:"foreignkey:NewsId;references:ID"`
}

func (News) TableName() string {
	return "n_news"
}
