# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html


# 设计多种item数据存储方案，前一个管道的输出，将做为下一个管道的输入，来出item

class BaiduMongoDBPipeline(object):
    def process_item(self, item, spider):
        collection.insert(dict(item))

        return item


class BaiduJsonPipeline(object):
    def process_item(self, item, spider):

        json_strjson.dumps(dict(item))
        return item


class BaiduMySQLPipeline(object):
    def process_item(self, item, spider):

        json_strjson.dumps(dict(item))
        return item
