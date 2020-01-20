package main

import (
	util "./utils"
	"fmt"
)

func main() {
	defer util.From_Db_close()
	err := util.From_Db_iniit()
	if err != nil {
		fmt.Printf("init Db failed err:%v\n", err)
	}
	fmt.Println("初始化数据库成功")
	util.From_Db_Select(0)
	//util.Login()
	util.From_Db_Insert("王五", 22)
	util.From_Db_Select(0)
	util.From_Db_Deleat(1)
}
