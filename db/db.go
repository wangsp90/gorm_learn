package db

//package main

import (
	//"business"
	"config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"log"
	//"strconv"
)

func Init() *gorm.DB {
	cfg := config.Getcfg()
	db, err := gorm.Open("mysql", cfg.Db)
	if err != nil {
		panic(err)
	}
	return db
}
