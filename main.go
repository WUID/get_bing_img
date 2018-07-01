package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main(){
    fmt.Println("hello")
    resp, err := http.Get("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1")
    if err != nil {
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
    }
    fmt.Println(string(body))
}

