package models

type AnosuAll struct {
	ID  uint   `json:"id"`
	Url string `json:"url"`
}

func (AnosuAll) TableName() string {
	return "anosu_all"
}
