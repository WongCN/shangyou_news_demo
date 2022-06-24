package models

type User struct {
	// 展示新闻上传者,简单写两个字段意思一下
	ID       uint
	Username string `gorm:"column:username" json:"username"`
	Age      uint8  `gorm:"column:age" json:"age"`
}

func (User) TableName() string {
	return "n_user"
}
