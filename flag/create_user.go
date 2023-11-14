package flag

import (
	"Goblog/global"
	"Goblog/models/ctype"
	"Goblog/service/user_ser"
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

	//判断角色 命令行创建 只有 user 和 admin
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	//校验两次密码
	if password != rePassword {
		global.Log.Error("两次密码不一致,请重新输入")
		return
	}
	//判断逻辑结构
	//判断用户名是否存在
	err := user_ser.UserService{}.CreateUser(userName, nickName, password, role, email, "")
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功!", userName)
}

func toString() string {
	return fmt.Sprintf("用户名:%s \\昵称:%s \\密码:%s \\确定密码:%s \\邮箱:%s", userName, nickName, password, rePassword, email)
}
