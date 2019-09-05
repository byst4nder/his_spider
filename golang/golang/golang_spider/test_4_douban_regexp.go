package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main()  {
	t1 :=time.Now()
	totalPage :=9
	parse1(totalPage)
	elapsed :=time.Since(t1)
	fmt.Println("爬虫总共耗时：", elapsed)
}

type Spider1 struct {
	url string
	header map[string]string
}

func(keyword Spider1) get_html_header() string{
	client :=http.Client{}
	req, err:=http.NewRequest("GET", keyword.url, nil)
	if err !=nil{
		fmt.Println("失败原因为：", err)
	}
	for key, value := range keyword.header{
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	return string(body)
}

func parse1(totals int)  {
	header :=map[string]string{
		"host":"movie.douban.com",
		"Connection": "keep-alive",
		"Cache-Control": "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent": "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer": "https://movie.douban.com/top250",
	}

	f, err :=os.Create("D:/doubandianying.txt")
	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	defer f.Close()
	f.WriteString("电影名称"+"\t"+"评分"+"\t"+"评价人数"+"\r\n")
	for i:=0; i<=totals;i++{
		fmt.Println("正在抓取第"+strconv.Itoa(i)+"页数据")
		url :="https://movie.douban.com/top250?start="+ strconv.Itoa(i*25)+"&filter="
		spider := Spider1{url:url, header:header}
		html :=spider.get_html_header()

		pattern12 :=`<span>(.*?)评价</span>`
		resp12 :=regexp.MustCompile(pattern12)
		find_txt12 := resp12.FindAllStringSubmatch(html, -1)
		fmt.Println(find_txt12)

		pattern13 :=`property="v:average">(.*?)</span>`
		resp13 :=regexp.MustCompile(pattern13)
		find_txt13 := resp13.FindAllStringSubmatch(html, -1)
		fmt.Println(find_txt13)

		pattern14 :=`alt="(.*?)" src="`
		resp := regexp.MustCompile(pattern14)
		find_txt14 := resp.FindAllStringSubmatch(html, -1)
		fmt.Println(find_txt14)
		f.WriteString("\xEF\xBB\xBF")
		for i:=0; i<len(find_txt12);i++{
			f.WriteString(find_txt14[i][1]+"\t"+ find_txt13[i][1]+"\t"+find_txt12[i][1]+"\t"+"\r\n")
		}
	}
}