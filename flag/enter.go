package flag

import FLAG "flag"

type Option struct {
	DB bool
}

// Parse 解析命令行参数
func Parse() Option {
	db := FLAG.Bool("db", false, "初始化数据库")
	//解析命令写入注册的flag中
	FLAG.Parse()
	return Option{DB: *db}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
