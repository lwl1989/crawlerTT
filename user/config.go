package user

import (
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"os"
	"strings"
)


type Config struct {
	Mgo string	`json:"mgo"`
	DB  string	`json:"db"`
	ShowRecord int	`json:"show_record"`
	SessionId	string `json:"session_id"`
}

var OneConfig *Config

func getConfig() *Config{
	return &Config{
	}
}

func GetConfig(path string) *Config {
	if path == "" {
		path = getCurrentDirectory() + "/config.json"
	}
	data, err := ioutil.ReadFile(path)
	if err != nil{
		panic("load config err "+path)
	}
	b := []byte(data)
	config := getConfig()
	err = json.Unmarshal(b, config)
	if err != nil{
		panic("load config err "+path)
	}
	OneConfig = config
	return config
}


func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1)
}
