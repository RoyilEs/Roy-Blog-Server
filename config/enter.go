package config

// 实体化
type Config struct {
	MySql  MySql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
}
