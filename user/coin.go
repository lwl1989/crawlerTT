package user

import (
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type CoinRes struct {
	Code int         `json:"code"`
	Data CoinResData `json:"data"`
}
type CoinResData struct {
	RCoin int64  `json:"remaining_coin"`
	Exp   string `json:"expire_date,omitempty"`
	TCoin int64  `json:"tot_coin"`
}

type CoinResL struct {
	Code int         `json:"code"`
	Data []*CoinResLogData `json:"data"`
}
type CoinResLogData struct {
	Date string `json:"his_date,omitempty"`
	Type string `json:"his_trans_type,omitempty"`
	Trans int `json:"his_trans_coin,omitempty"`
	Info string `json:"trans_info,omitempty"`
	Exp string `json:"expire_date,omitempty"`
	Img string `json:"his_img,omitempty"`
}

func getCoinReq(userId string) CoinResData {
	content := "account_id=" + userId

	req, _ := http.NewRequest(http.MethodPost, UserCoinUrl, strings.NewReader(content))

	c := getClient()
	req.ContentLength = int64(len(content))
	setHeader(req)
	response, err := c.Do(req)
	if err != nil {
		fmt.Println("coin:", err.Error())
		return CoinResData{}
	}

	body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body[:]))
	if err != nil {
		fmt.Println("io:", err.Error())
		return CoinResData{}
	}
	res := &CoinRes{}
	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("json:", err.Error())
		return CoinResData{}
	}
	return res.Data
}


func getCoinReqLog(userId string) []*CoinResLogData {
	content := "account_id=" + userId

	req, _ := http.NewRequest(http.MethodPost, UserCoinLogUrl, strings.NewReader(content))

	c := getClient()
	req.ContentLength = int64(len(content))
	setHeader(req)
	response, err := c.Do(req)
	if err != nil {
		fmt.Println("coin:", err.Error())
		return nil
	}
	body, err := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body[:]))
	if err != nil {
		fmt.Println("io:", err.Error())
		return nil
	}
	res := &CoinResL{}
	err = json.Unmarshal(body, res)
	if err != nil {
		fmt.Println("json:", err.Error())
		return nil
	}
	//fmt.Println(res.Data)
	return res.Data
}
