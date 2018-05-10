package main

import "github.com/lwl1989/crawlerTT/user"

func main()  {
	user.GetConfig("/www/go_path/src/github.com/lwl1989/crawlerTT/config.json")
	user.GetList()
	//user.GetUserInfo("MkigJdpKFSkE6ZC")
}