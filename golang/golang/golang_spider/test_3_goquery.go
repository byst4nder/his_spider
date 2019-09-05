package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

const html string  = `<body>
                <div id="div1" >DIV1</div>
                <div class="div2">DIV2</div>
                <span id="div1">SPAN</span>
            	</body>`

func main()  {
	//testElement()
	testId()
}
/*
标签选择器的用法
*/

func testElement()  {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil{
		fmt.Println("失败原因：", err)
	}
	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("i",i, "select text", selection.Text())
	})
}

/*
id 选择器的用法和class选择器的用法
#div1, .div1
*/

func testId(){
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil{
		fmt.Println("失败原因为：", err)
	}
	dom.Find("div#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println("i:", i, "select text:", selection.Text())
	})
}