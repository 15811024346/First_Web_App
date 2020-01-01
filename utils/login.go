package utils

import "fmt"

var UserInfo struct {
	name        string
	pwd         string
	loginChange int8
}

//登陆方法
func Login() {

	UserInfo.loginChange = 3
	for i := 0; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&UserInfo.name)
		fmt.Println("请输入密码")
		fmt.Scanln(&UserInfo.pwd)
		if UserInfo.name == "王帅" && UserInfo.pwd == "123456" {
			fmt.Println("登陆成功")
			break
		} else {
			UserInfo.loginChange--
			if UserInfo.loginChange == 0 {
				fmt.Println("机会用完了，登陆失败。")
				return
			} else {
				fmt.Printf("用户名/密码错误。\n请重新输入。您还有%d次机会\n", UserInfo.loginChange)
			}
		}
	}
}
