package main

import (
	"io"
	"os"
    "fmt"
    "time"
    "strings"
    "net/http"
    "io/ioutil"
)

func get_img_url() string {

    resp, err := http.Get("https://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=0")
    if err != nil {
        fmt.Println(err)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()

    ret_json := string(data)
    url_start := strings.Index(ret_json, `/az/hprichbg/rb/`)
    url_end := strings.Index(ret_json, `","urlbase"`)
    img_url := string("")

    for i := url_start;i < url_end;i++ {
        img_url += string(ret_json[i])
    }

    return ("http://www.bing.com" + img_url)
}

func dowload_img(url string) {

    res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
    }
    
    f, err := os.Create(time.Now().Format("2006-01-02 15:04:05") + ".jpg")
	if err != nil {
		fmt.Println(err)
    }
    
    io.Copy(f, res.Body)
    
}

func main(){

    dowload_img(get_img_url())

}

