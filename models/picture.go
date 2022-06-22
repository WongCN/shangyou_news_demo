package models

type Picture struct {
	ID             uint8  `gorm:"column:id" json:"id"`
	NewsId         uint   `gorm:"column:news_id" json:"news_id"`
	PictureAddress string `gorm:"column:picture_address" json:"picture_address"`
}

func (Picture) TableName() string {
	return "n_picture"
}
