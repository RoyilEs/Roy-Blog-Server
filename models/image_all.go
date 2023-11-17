package models

type ImageAll struct {
	ID  uint   `json:"id" gorm:"primarykey"`
	Pid uint   `json:"pid"`
	Url string `json:"url"`
}

// TableName 解决gorm表明映射
func (ImageAll) TableName() string {
	return "image_all"
}
