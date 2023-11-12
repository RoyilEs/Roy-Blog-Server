package flag

import (
	"Goblog/global"
	"Goblog/models"
	"Goblog/models/ctype"
	"Goblog/utils/pwd"
	"fmt"
)

var (
	userName   string
	nickName   string
	password   string
	rePassword string
	email      string
)

func CreateUser(permissions string) {
	//创建用户逻辑结构
	// 用户名 昵称 密码 确认密码 邮箱

	fmt.Printf("请输入用户名:")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称:")
	fmt.Scan(&nickName)
	fmt.Printf("请输入邮箱:")
	fmt.Scan(&email)
	fmt.Printf("请输入密码:")
	fmt.Scan(&password)
	fmt.Printf("请确定密码:")
	fmt.Scan(&rePassword)

	fmt.Println(toString())

	//判断逻辑结构
	//判断用户名是否存在
	var userModel models.User
	err := global.DB.Take(&userModel, "username = ?", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在,请重新输入")
		return
	}
	//校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致,请重新输入")
		return
	}
	//TODO 正则密码强度
	//加密密码 hash
	hashPassword := pwd.HashPassword(password)
	//判断角色
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	//头像问题 1.默认 2.随机选择
	avatar := "/uploads/avatar/头像.png"

	//入库
	err = global.DB.Create(&models.User{
		Nickname:   nickName,
		Username:   userName,
		Email:      email,
		Password:   hashPassword,
		Permission: role,
		Avatar:     avatar,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功!", userName)
}

func toString() string {
	return fmt.Sprintf("用户名:%s \\昵称:%s \\密码:%s \\确定密码:%s \\邮箱:%s", userName, nickName, password, rePassword, email)
}
