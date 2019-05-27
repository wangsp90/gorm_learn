package db

import (
	"time"
)

type Biz struct {
	Bk_biz_id           int       `json:"bk_biz_id" gorm:"int;not null;unique_index;primary key"`
	Default             int       `json:"default" gorm:"int"`
	Language            string    `json:"language" gorm:"VARCHAR(32)"`
	Life_cycle          string    `json:"life_cycle" gorm:"VARCHAR(32)"`
	Bk_biz_developer    string    `json:"bk_biz_developer" gorm:"LONGTEXT"`
	Bk_biz_maintainer   string    `json:"bk_biz_maintainer" gorm:"LONGTEXT"`
	Bk_biz_tester       string    `json:"bk_biz_tester" gorm:"LONGTEXT"`
	Time_zone           string    `json:"time_zone" gorm:"VARCHAR(255)"`
	App_manager_org     string    `json:"app_manager_org" gorm:"VARCHAR(255)"`
	Biz_type            string    `json:"biz_type" gorm:"VARCHAR(255)"`
	Create_time         time.Time `json:"create_time"`
	Biz_manager         string    `json:"biz_manager" gorm:"LONGTEXT"`
	Bk_supplier_account string    `json:"bk_supplier_account" gorm:"VARCHAR(32)"`
	Operator            string    `json:"operator" gorm:"VARCHAR(255)"`
	Bk_biz_productor    string    `json:"bk_biz_productor" gorm:"LONGTEXT"`
	Bk_biz_name         string    `json:"bk_biz_name" gorm:"VARCHAR(255)"`
	Last_time           time.Time `json:"last_time"`
	Bk_supplier_id      int       `json:"bk_supplier_id" gorm:"int"`
}

var s string = `{
	"bk_biz_id": 26, "default": 0, "language": "1", "life_cycle": "2", "bk_biz_developer": "", 
	"bk_biz_maintainer": "admin,010500571,030200531,122466310,082922381,021127100,100651110,102220740", "bk_biz_tester": "", 
	"time_zone": "Asia/Shanghai", "app_manager_org": "", "biz_type": "prod", "create_time": "2019-04-23T08:51:27.52+08:00", 
	"biz_manager": "", "bk_supplier_account": "0", "operator": "", "bk_biz_productor": "", "bk_biz_name": "KK测试", 
	"last_time": "2019-05-13T09:18:48.661+08:00", "bk_supplier_id": 0}, {"bk_biz_id": 37, "default": 0, "language": "1", "life_cycle": "2", 
	"bk_biz_developer": "", "bk_biz_maintainer": "admin,010500571,122466310,082922381,021127100,100651110,102220740", "bk_biz_tester": "", 
	"time_zone": "Asia/Shanghai", "app_manager_org": "", "biz_type": "prod", "create_time": "2019-04-28T09:11:22.908+08:00", "biz_manager": "", 
	"bk_supplier_account": "0", "operator": "", "bk_biz_productor": "", "bk_biz_name": "KK", "last_time": "2019-05-13T09:20:54.273+08:00", "bk_supplier_id": 0
}`
