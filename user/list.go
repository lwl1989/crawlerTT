package user

import (
	"strings"
	"net/http"
	"fmt"
	"gopkg.in/mgo.v2"
	"io"
	"strconv"
	"h12.io/html-query/expr"
	"h12.io/html-query"
	"time"
)


var Got chan *User
var Stop chan bool
var Sum int
var NowStep int

func GetList()  {

	page := &pageForm{currentPage:1,orderType:"",orderCol:"",pageFlag:true,showRecord:OneConfig.ShowRecord}

	req,_ :=http.NewRequest(http.MethodPost,UserListUrl,strings.NewReader(page.toString()))

	c := getClient()

	SetJarCook(c)

	content := page.toString()
	req.ContentLength = int64(len(content))
	setHeader(req)


	r, err := c.Do(req)

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	foreachSave(r.Body)
}


func foreachSave(body io.ReadCloser) {

	root, err := query.Parse(body)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(OneConfig.Mgo)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}
	collection := session.DB(OneConfig.DB).C("user_"+strconv.FormatInt(int64(time.Now().Day()),10))


	Got = make(chan *User)
	Stop = make(chan bool)
	Sum = OneConfig.ShowRecord
	NowStep = 0
	j:=0
	root.Table().Children(expr.Tbody).For(func(n *query.Node) {

		n.Children(expr.Tr).For(func(n *query.Node) {
			i:=0
			u:=&User{

			}
			n.Children(expr.Td).For(func(n *query.Node) {
				if i > 1 {
					switch i {
					case 2:
						u.UserId=*n.PlainText()
					case 3:
						u.UserName=*n.PlainText()
					case 4:
						u.Mobile=*n.PlainText()
					case 5:
						u.CardNum=*n.PlainText()
					case 6:
						u.Device=*n.PlainText()
					case 7:
						u.DeviceToken=*n.PlainText()
					case 8:
						u.AddTime=*n.PlainText()
					case 9:
						u.UpTime=*n.PlainText()
					}
				}
				i++
			})
			j++
			go u.GetOther()


			//fmt.Println(u)
		})
	})
	Sum = j
	for {
		select{
			case u:=<-Got:
				collection.Insert(u)

			case <-Stop:
				fmt.Println("now exec over")
				break
		}

	}
	defer body.Close()
	defer session.Close()
}

func (u *User) GetOther() {
	//fmt.Println(u)
	u.GetCoin()
	u.GetUserInfo()
	u.GetUserCoinLog()

	fmt.Println(u)

	Got <- u
	NowStep += 1

	if NowStep == Sum {
		Stop <- true
	}
}
func (u *User) GetCoin()  {
	coin := getCoinReq(u.UserId)
	u.TotalCoin =  strconv.FormatInt(coin.TCoin,10)
	u.RemainingCoin = strconv.FormatInt(coin.RCoin,10)
	u.ExpireDate = coin.Exp
}
func (u *User) GetUserInfo()  {
	u.Info = GetUserInfo(u.UserId)
}

func (u *User) GetUserCoinLog()  {
	u.CoinLog = getCoinReqLog(u.UserId)
}


