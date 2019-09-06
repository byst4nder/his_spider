# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/en/latest/topics/items.html

# 预先定义保存数据的字典

import scrapy

class BaiduItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    #标题
    title = scrapy.Field()
    #链接
    url = scrapy.Field()


#### baiduspider.py

# from ..items import BaiduItem

# # item是一个类字典的对象
# item = BaiduItem()
# #item = {}

# item['ttile'] = response.xpath("//title/text()").extract_first()
# item['url'] = response.url


