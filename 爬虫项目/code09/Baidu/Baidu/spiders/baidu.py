# -*- coding: utf-8 -*-

# 导入scrapy
import scrapy

from ..items import BaiduItem

# 自定义一个爬虫类，继承于scrapy.Spider
class BaiduSpider(scrapy.Spider):

    # (必须）爬虫名（用来区分其他爬虫）
    name = 'baidu'
    # （可选）允许抓取的网页域名范围
    allowed_domains = ['baidu.com']

    # （必须）程序启动，发送的 第一批请求的url地址
    # start_urls只需要提供url地址即可，不需要构建请求（引擎会读取start_urls并构建每个url请求再交给调度器保存 - 下载器发送请求 返回响应 - spider的parse方法解析响应
    start_urls = ['http://www.baidu.com/', "http://news.baidu.com/"]

    def parse(self, response):
        # 构建一个类字典的item对象
        item = BaiduItem()
        #item = {}

        # 从响应里提取数据，并保存到item中
        item['title'] = response.xpath("//title/text()").extract_first()
        item['url'] = response.url

        print(item)



# spider = BaiduSpider()
# for url in spider.start_urls:
#     response = Request(url)
#     spider.parse(response)
