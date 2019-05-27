package main

import (
	//"business"
	// "bytes"
	// "config"
	"db"
	// "encoding/json"
	// "io/ioutil"
	// "log"
	// "net/http"
)

func main() {
	mydb := db.Init()
	//db.CreateBizTable(mydb)
	defer mydb.Close()

}

var parammap string = `{
    "bk_app_code": "hdzjz",
    "bk_app_secret": "9d6b4db4-de11-4047-b60f-891caef02b8e",
    "bk_username":"021127100",
    "ip": {
        "data": ["192.168.7.11"],
        "exact": 1,
        "flag": "bk_host_innerip"
    },
    "condition": [
        {
            "bk_obj_id":"biz",
            "fields":[],
            "condition":[]
        }
    ]
}`

//全新插入业务
// func main01() {
// 	//获取配置文件信息
// 	mapcfg := config.Getcfg()
// 	log.Println(mapcfg)
// 	apiend := "search_business/"
// 	//获取API完成链接
// 	var api string
// 	api = mapcfg.API + apiend
// 	log.Println(api)
// 	//设置传递的参数
// 	param, _ := json.MarshalIndent(mapcfg.Parameter, "", "  ")
// 	log.Println(string(param))
// 	//请求数据
// 	body := bytes.NewBuffer(param)
// 	rq, _ := http.Post(api, "application/json; charset=utf-8", body)
// 	result, err := ioutil.ReadAll(rq.Body)
// 	defer rq.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	//log.Println(string(result))
// 	// 数据处理
// 	var businesslist business.Business
// 	err1 := json.Unmarshal(result, &businesslist)
// 	if err1 != nil {
// 		log.Println(err1)
// 	}

// 	cmdb := db.ConnectDatabase()
// 	defer cmdb.Close()
// 	// done := db.Insertdata(cmdb, businesslist)
// 	// log.Println(done)

// }
