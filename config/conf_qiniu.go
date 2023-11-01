package config

type QiNiu struct {
	Enable    string  `json:"enable" yaml:"enable"` //是否启用七牛存储
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` //存储桶的名字
	CDN       string  `json:"cdn" yaml:"cdn"`       //访问图片的地址前缀
	Zone      string  `json:"zone" yaml:"zone"`     //存储的地区
	Size      float64 `json:"size" yaml:"size"`     //存储大小限制 单位MB
}
