# coding:utf-8

# 可以通过命令创建爬虫
# #scrapy genspider itcast itcast.cn


import scrapy

from ..items import ItcastItem

class ItcastSpider(scrapy.Spider):

    name = "itcast"

    allowed_domains = ["itcast.cn"]

    start_urls = ["http://www.itcast.cn/channel/teacher.shtml"]


    def parse(self, response):

        node_list = response.xpath("//div[@class='li_txt']")

        # 迭代取出每个老师信息，并保存在item中
        for node in node_list:
            item = ItcastItem()

            item['name'] = node.xpath("./h3/text()").extract_first()
            item['title'] = node.xpath("./h4/text()").extract_first()
            item['info'] = node.xpath("./p/text()").extract_first()

            yield item



# 1. scrapy crawl itcast -o itcast.json  (csv、xml、jl)
# 2. 如果需要将数据存储到scrpay不支持的格式里，比如数据库等，就必须通过管道实现

#engine.py

# Engine里的每次for迭代 parse() 方法，用来处理一个response响应提取的数据（请求、item)
# for result in spider.parse(response):

#     if isinstance(result, scrapy.Item):
#         pipeline.process_item(resutl, spider)

#     elif isinstance(result, scrapy.Request):
#         scheduler.add_request(result)




