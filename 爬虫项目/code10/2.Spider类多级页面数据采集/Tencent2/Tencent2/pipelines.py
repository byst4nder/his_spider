# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html


from items import TencentItem, PositionItem
import json


class TencentPipeline(object):
    def open_spider(self, spider):
        self.f  = open("tencent.json", "w")

    def process_item(self, item, spider):
        if isinstance(item, TencentItem):
            json_str = json.dumps(dict(item))
            self.f.write(json_str)

        return item

    def close_spider(self, spider):
        self.f.close()


class PositionPipeline(object):
    def open_spider(self, spider):
        self.f  = open("position.json", "w")

    def process_item(self, item, spider):
        if isinstance(item, PositionItem):
            json_str = json.dumps(dict(item))
            self.f.write(json_str)
        return item

    def close_spider(self, spider):
        self.f.close()

