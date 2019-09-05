package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

func main() {
	keyWords := [...]string{"体育", "娱乐", "登录"}
	startPage := 0
	endPage := 5
	for _, value := range keyWords {
		for i := startPage; i <= endPage; i++ {
			SpiderPage(i, value)
		}
	}
}

func HttpGet(url string) *http.Response {
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("request建立失败原因：", err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	request.Header.Set("Host", "www.sogou.com")
	request.Header.Set("Referer", "https://www.sogou.com/sogou?query=%E4%BD%93%E8%82%B2&ie=utf8&_ast=1562381568&_asf=null&w=01029901&pid=sogou-wsse-a9e18cb5dd9d3ab4&duppid=1&cid=&s_from=result_up&insite=wenwen.sogou.com")
	request.Header.Set("Cookie", "sw_uuid=3273955461; ABTEST=0|1562912310|v17; SNUID=80723F2A4E48C33BAF58BBBF4EE8EBF6; IPLOC=CN8100; SUID=CE3C71672313940A000000005D282636; ld=Nyllllllll2N3jYLlllllV1oJ0YlllllBT9D8kllll9lllllRllll5@@@@@@@@@@; SUV=1562912311779681; browerV=3; osV=1; LSTMV=1682%2C173; LCLKINT=4909; taspeed=taspeedexist; pgv_pvi=7106629632; pgv_si=s2438505472")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request请求失败原因为：", err)
	}
	return response
}

func SpiderPage(page int, keyWord string) {
	url := "https://www.sogou.com/sogou?query=" + keyWord + "&pid=sogou-wsse-a9e18cb5dd9d3ab4&insite=wenwen.sogou.com&duppid=1&rcer=&page=" + strconv.Itoa(page) + "&ie=utf8"
	rsl := HttpGet(url)
	dom, err := goquery.NewDocumentFromReader(rsl.Body)
	if err != nil{
		fmt.Println("失败原因：", err)
	}
	fmt.Println("1", dom)
	dom.Find("#t").Each(func(i int, selection *goquery.Selection) {
		ContentUrl, err:= selection.Attr("href")
		fmt.Println("2", ContentUrl, err)
	})
}
