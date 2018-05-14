

package user

import (
"strings"
"fmt"
"net/http"
"io"
"strconv"
"time"
"h12.io/html-query/expr"
"gopkg.in/mgo.v2"
"h12.io/html-query"
)

type Data struct {
	Data []string
}

func GetShopAccount()  {
	page := &pageForm{currentPage:1,orderType:"",orderCol:"",pageFlag:true,showRecord:OneConfig.ShowRecord}

	req,err := http.NewRequest(http.MethodPost,ShopUrl,strings.NewReader(page.toString()))

	if err != nil{
		fmt.Println(err.Error())
		return
	}
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

	SaveAccount(r.Body,false)
}

func GetAccount()  {

	page := &pageForm{currentPage:1,orderType:"",orderCol:"",pageFlag:true,showRecord:OneConfig.ShowRecord}

	req,err := http.NewRequest(http.MethodPost,AdminUrl,strings.NewReader(page.toString()))

	if err != nil{
		fmt.Println(err.Error())
		return
	}
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

	SaveAccount(r.Body,true)
}

func SaveAccount(body io.ReadCloser,admin bool) {

	root, err := query.Parse(body)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(OneConfig.Mgo)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}
	collection := &mgo.Collection{}
	if admin {
		collection = session.DB(OneConfig.DB).C("account_" + strconv.FormatInt(int64(time.Now().Day()), 10))
	}else{
		collection = session.DB(OneConfig.DB).C("shop_" + strconv.FormatInt(int64(time.Now().Day()), 10))
	}
	NowStep = 0
	j:=0
	root.Table(expr.Id("table")).Children(expr.Tbody).For(func(n *query.Node) {

		n.Children(expr.Tr).For(func(n *query.Node) {
			i:=0
			u:=make([]string,0)
			n.Children(expr.Td).For(func(n *query.Node) {

				if i > 1 {
					if i==2 {
						uid := string(*n.PlainText())
						u=getUsers(uid,admin)
						//fmt.Println(u)
					}
				}
				i++
			})
			j++
			collection.Insert(&Data{u})
		})
	})

	defer body.Close()
	defer session.Close()
}

func getUsers(accountId string,admin bool) []string {
	content := "editAccountId=" + accountId
	u:=make([]string,0)
	u=append(u,accountId)
	url := ""
	if admin {
		url = AdminContentUrl
	}else{
		url = ShopContentUrl
	}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(content))
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


	root.Table().Children(expr.Tbody).For(func(n *query.Node) {
		n.Children(expr.Tr).For(func(n *query.Node) {

			n.Children(expr.Td).For(func(n *query.Node) {
				n.Children().For(func(n *query.Node) {
					u=append(u,*n.Value())
					u=append(u,*n.PlainText())
				})
				//switch i {
				//case 1:
				//	n.Div(expr.Class("form-inline")).Children(expr.Input).For(func(n *query.Node) {
				//		u=append(u,*n.Value() )
				//	})
				//case 2:
				//	n.Div(expr.Class("form-inline")).Children(expr.Input).For(func(n *query.Node) {
				//		u=append(u,*n.Value() )
				//	})
				//
				//case 10:
				//	fmt.Println(*n.Input(expr.Id("canUseFunctions")).Value())
				//default:
				//	n.Div(expr.Class("form-inline")).Children(expr.Input).For(func(n *query.Node) {
				//		u=append(u,*n.Value() )
				//	})
				//}
				//i+=1

			})
		})
	})
	return u
}
