#coding:utf-8


import scrapy


class TestSpider(scrapy.Spider):
    name = "test"
    allowed_domains = ["hr.tencent.com"]

    start_urls = ["http://www.baidu.com/", "http://www.baidu.com/", "http://www.baidu.com/", "http://www.baidu.com/",]

    def parse(self, response):
        print(len(response.body))
