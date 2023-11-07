package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// InList 是否存在列表里面
func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

// Md5 加密
func Md5(str []byte) string {
	m := md5.New()
	m.Write(str)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// GetValidMsg 返回结构体中中的msg参数
func GetValidMsg(err error, obj any) string {
	//使用的时候 需要传Obj的指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名 获得结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
