package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Cfginfo struct {
	Parameter Parameter
	API       string
	Db        string
}
type Parameter struct {
	Bk_app_code   string `json:"bk_app_code"`
	Bk_app_secret string `json:"bk_app_secret"`
	Bk_username   string `json:"bk_username"`
}

func Getcfg() (cfginfo Cfginfo) {
	buf, errOpen := ioutil.ReadFile("./config/cfg.json")
	if errOpen != nil {
		log.Println("配置文件加载失败...")
		log.Println(errOpen.Error())
		return
	}
	errjson := json.Unmarshal(buf, &cfginfo)
	if errjson != nil {
		log.Println("error:", errjson)
		return
	}
	return
}
