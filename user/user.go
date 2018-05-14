package user

import (
	"strings"
	"net/http"
	"h12.io/html-query"
	//"h12.io/html-query/expr"
	"h12.io/html-query/expr"
)

type User struct {
	UserId	string
	UserName string
	Mobile 	string
	CardNum string
	Device 	string
	DeviceToken string
	AddTime string
	UpTime string
	RemainingCoin string
	ExpireDate	string
	TotalCoin	string

	Info	*UserInfo

	CoinLog []*CoinResLogData

}

type UserInfo struct {
	Email 	string
	Address string
	Nationality string  //國籍
	LastModify  string
	ModifyRecord []*UserModifyRecord
}

type UserModifyRecord struct {
	ModifyMessage string
	ModifyType	  string
	ModifyReason	string
	ModifyTime	string
}

func GetUserInfo(accountId string) *UserInfo {
	content := "editAccountId=" + accountId

	req, err := http.NewRequest(http.MethodPost, UserInfoUrl, strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	c := getClient()
	req.ContentLength = int64(len(content))
	setHeader(req)
	SetJarCook(c)
	response, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	root, err := query.Parse(response.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(*root.PlainText())

	//fmt.Println(*root.Table(expr.Id("userDataEidt")).Render())
	u:=&UserInfo{
		ModifyRecord: make([]*UserModifyRecord,0),
	}
	root.Table(expr.Id("userDataEidt")).Children(expr.Tbody).For(func(n *query.Node) {
		i:=0

		n.Children(expr.Tr).For(func(n *query.Node) {
				//fmt.Println(*n.Render())
				n.Children(expr.Td).For(func(n *query.Node) {
					switch i {
					case 5:
						u.Email = *n.PlainText()
					case 6:
						u.Address = *n.PlainText()
					case 7:
						u.Nationality = *n.PlainText()
					case 8:
						u.LastModify = *n.PlainText()
					}
				})
			i++
		})
		//fmt.Println(u)
	})
	root.Table(expr.Id("userLogTable")).Children(expr.Tbody).For(func(n *query.Node) {
		n.Children(expr.Tr).For(func(n *query.Node) {
			i:=0
			log := &UserModifyRecord{}
			n.Children(expr.Td).For(func(n *query.Node) {
				switch i {
				case 1:
					log.ModifyMessage = *n.PlainText()
				case 2:
					log.ModifyType = *n.PlainText()
				case 3:
					log.ModifyReason = *n.PlainText()
				case 4:
					log.ModifyTime = *n.PlainText()
				}
			})
			if log.ModifyTime != "" {
				u.ModifyRecord = append(u.ModifyRecord, log)
			}
		})
	})

	return u
}