package config

// Config 实体化 yaml 信息
type Config struct {
	MySql    MySql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	Jwt      Jwt      `yaml:"jwt"`
	QiNiu    QiNiu    `yaml:"qi_niu"`
	Upload   Upload   `yaml:"upload"`
}
