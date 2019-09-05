package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	urlBaiDu = "http://www.baidu.com"
)

func main() {
	//test_get()
	//test_post()
	//testPostForm()
	testClient()
}

func testClient()  {
	url := "http://www.baidu.com"
	client := &http.Client{
		Timeout:3*time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	if err != nil {
		fmt.Println("失败原因为：", err)
	}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	defer resp.Body.Close()
	rlt, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	fmt.Println(string(rlt))

}

func testPostForm()  {
	url1 :="https://accounts.douban.com/j/mobile/login/basic"
	resp, err := http.PostForm(url1, url.Values{"name": {"zhangsan@xx.com"}, "password":{"123456"}})
	fmt.Println(resp.Request.URL)
	defer resp.Body.Close()
	if err != nil{
		fmt.Println("直白原因为：", err)
	}
	rlt, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	fmt.Println(string(rlt))
}

func testGet() {
	resp, err := http.Get(urlBaiDu)
	if err != nil {
		fmt.Printf("失败原因为：%v\n", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("失败原因为：%v\n", err)
	}
	fmt.Println(string(body))
}

func testPost() {
	body := " "
	resp, err := http.Post(urlBaiDu, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println("net http post method err, ", err)
	}
	defer resp.Body.Close()
	rlt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取失败原因为：", err)
	}
	fmt.Println(string(rlt))
}
