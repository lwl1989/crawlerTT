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
	//user.GetConfig(path)
	//path = "config.json"
	user.GetConfig(path)
	user.GetList()  //獲取用戶
	user.GetAccount() //獲取管理員
	user.GetMessageList()  //獲取後臺消息
	user.GetSendMessageList()  //獲取已發送消息
	user.GetShopAccount()  //獲取商店用戶
	//user.GetUserInfo("MkigJdpKFSkE6ZC")
}