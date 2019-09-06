# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

import pymongo
#import requests

class ItcastMongoDBPipeline(object):
    #def __init__(self):
    def open_spider(self, spider):
        """ 爬虫启动时候，会执行一次"""
        self.client = pymongo.MongoClient()
        self.collection = self.client.itcast.teacher

    def process_item(self, item, spider):
        """ 每次引擎传入一个item，就会调用一次process_item"""
        # 将item对象转为python的字典，并写入mongodb
        self.collection.insert(dict(item))
        return item

    #def __del__(self)
    def close_spider(self, spider):
        """ 爬虫结束时，会执行一次 """
        pass


# pipeline = ItcastPipeline()

# pipeline.open_spider(spider)

# pipeline.process_item(item, spider)
# pipeline.process_item(item, spider)
# pipeline.process_item(item, spider)
# pipeline.process_item(item, spider)
# pipeline.process_item(item, spider)
# pipeline.process_item(item, spider)
# ...

# pipeline.close_spider(spider)
