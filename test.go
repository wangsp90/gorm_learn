package main

import (
    //"encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

type Test struct {
    Id          string
    Bk_app_code string
    Bk_username string
}

func main() {
    str := `{
            "bk_app_code":"hdzjz",
            "bk_app_secret":"9d6b4db4-de11-4047-b60f-891caef02b8e",
            "bk_username":"021127100",
            "condition": {
                "bk_biz_name": "KK"
            }
        }`
    //s1 := make(map[string]interface{})
    // var s1 Test
    // err := json.Unmarshal([]byte(str), &s1)
    // if err != nil {
    //     fmt.Println(err)
    // }
    // fmt.Println(s1.Id)
    body := strings.NewReader(str)
    rq, err := http.Post("https://paas.evergrande.com/api/c/compapi/v2/cc/search_business/", "application/json; charset=utf-8", body)
    if err != nil {
        panic(err)
    }
    result, err := ioutil.ReadAll(rq.Body)
    defer rq.Body.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(result))
}
