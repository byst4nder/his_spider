#coding:utf-8
import scrapy

class RenrenSpider(scrapy.Spider):
    name = "renren_login"
    #allowed_domains = []

    #start_urls = []


    def start_requests(self):
        """ 手动构建并发送登录的post请求，如果登录成功，scrapy会记录cookie"""
        post_url = "http://www.renren.com/PLogin.do"
        form_data = {
            "email" : "mr_mao_hacker@163.com",
            "password" : "ALARMCHIME"
        }
        # 发送post请求，scrapy会保存cookie，同时程序会跳转到self.parse执行后续代码
        yield scrapy.FormRequest(post_url, formdata=form_data, callback=self.parse)


    def parse(self, response):
        """ 直接发送好友的页面请求，scrapy会自动传递cookie"""
        url_list = [
            "http://www.renren.com/410043129/profile",
            "http://www.renren.com/965999739/profile"
        ]
        # scrapy会附带之前保存的cookie，发送每个好友的url地址请求，并通过parse_page解析响应
        for url in url_list:
            yield scrapy.Request(url, callback=self.parse_page)


    def parse_page(self, response):
        """ 处理每个好友页面的response响应 """
        title = response.xpath("//title/text()").extract_first()

        with open(title + ".html", "w") as f:
            f.write(response.body)
