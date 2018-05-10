package main

import (
	"github.com/lwl1989/crawlerTT/user"
	"fmt"
)

func main()  {
	path := ""
	fmt.Println("請輸入配置文件(默認當前目錄下config.json)")
	fmt.Scanln(&path)
	//"/www/go_path/src/github.com/lwl1989/crawlerTT/config.json"
	user.GetConfig(path)
	user.GetList()
	//user.GetUserInfo("MkigJdpKFSkE6ZC")
}