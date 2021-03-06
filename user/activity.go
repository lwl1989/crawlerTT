package user


import (
	"strings"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"h12.io/html-query/expr"
	"gopkg.in/mgo.v2"
	"h12.io/html-query"
)


func GetActivity()  {
	page := &pageForm{currentPage:1,orderType:"",orderCol:"",pageFlag:true,showRecord:OneConfig.ShowRecord}

	req,err := http.NewRequest(http.MethodPost,ActivityUrl,strings.NewReader(page.toString()))

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

	root, err := query.Parse(r.Body)
	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial(OneConfig.Mgo)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err)
	}

	collection := session.DB(OneConfig.DB).C("activity_" + strconv.FormatInt(int64(time.Now().Day()), 10))

	j:=0
	root.Table(expr.Id("table")).Children(expr.Tbody).For(func(n *query.Node) {

		n.Children(expr.Tr).For(func(n *query.Node) {
			i:=0
			u:=make([]string,0)
			n.Children(expr.Td).For(func(n *query.Node) {

				if i > 1 {
					if i==2 {
						id := string(*n.PlainText())
						u=getActivityContent(id)
						//fmt.Println(u)
					}
				}
				i++
			})
			j++
			collection.Insert(&Data{u})
		})
	})

	defer r.Body.Close()
	defer session.Close()
}

func getActivityContent(accountId string) []string {
	content := "editAccountId=" + accountId
	u:=make([]string,0)
	u=append(u,accountId)

	req, err := http.NewRequest(http.MethodPost, ActivityUrlContent, strings.NewReader(content))
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

	root.Children(expr.Table).For(func(root *query.Node) {
		root.Table().Children(expr.Tbody).For(func(n *query.Node) {
			n.Children(expr.Tr).For(func(n *query.Node) {

				n.Children(expr.Td).For(func(n *query.Node) {
					n.Children().For(func(n *query.Node) {
						u = append(u, *n.Value())
						u = append(u, *n.PlainText())
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
	})
	return u
}
