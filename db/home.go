package db

import (
	"time"
)

type Home struct {
	Id        int       `json:"id" gorm:"int; not null; unique; primary key; auto increment"`
	Name      string    `json:"name" gorm:"VARCHAR(32)"`
	Sex       string    `json:"sex" gorm:"VARCHAR(32)"`
	Age       int       `json:"age" gorm:"int"`
	Creattime time.Time `json:"creattime"`
}
