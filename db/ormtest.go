package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func CreateBizTable(db *gorm.DB) {
	if i := db.HasTable(&Biz{}); i == true {
		log.Println("数据库表已经存在！")
	} else {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Biz{}).Error
		if err != nil {
			log.Println(err)
		} else {
			log.Println("数据库表创建成功")
		}
	}
}

func CreateHomeTable(db *gorm.DB) {
	if i := db.HasTable(&Home{}); i == true {
		log.Println("数据库表已经存在！")
	} else {
		err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Home{}).Error
		if err != nil {
			log.Println(err)
		} else {
			log.Println("数据库表创建成功")
		}
	}

}
