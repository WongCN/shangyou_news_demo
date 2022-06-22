package models

type NewsType struct {
	ID uint8 `gorm:"column:id" json:"id"`
	// 新闻类型
	NewsType string `gorm:"column:news_type" json:"news_type"`
}

func (NewsType) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "n_news_type"
}
