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
    "time"
)

func main() {
    home := &db.Home{
        Name:      "Baolan",
        Sex:       "M",
        Age:       18,
        Creattime: time.Now(),
    }

    mydb := db.Init()
    //db.CreateBizTable(mydb)
    //db.CreateHomeTable(mydb)
    defer mydb.Close()
    if err := mydb.Create(home).Error; err != nil {
        panic(err)
    }
}
