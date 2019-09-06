#coding:utf-8

import json

import scrapy

from ..items import DouyuItem


class DouyuSpider(scrapy.Spider):
    name = "douyu"
    allowed_domains = ["douyucdn.cn"]

    # global base_url
    base_url = "http://capi.douyucdn.cn/api/v1/getVerticalRoom?limit=100&offset={}"
    page = 0

    start_urls = [base_url.format(page)]

    def parse(self, response):
        data_list = json.loads(response.body)["data"]

        if data_list:

            for data in data_list:
                item = DouyuItem()
                # 直播间的url地址
                item["room_url"] = "http://www.douyu.com/" + data["room_id"]
                # 图片url地址
                item["vertical_src"] = data["vertical_src"]
                # 昵称
                item["nickname"] = data["nickname"]
                # 城市
                item["anchor_city"] = data["anchor_city"]

                yield scrapy.Request(item['vertical_src'], meta={"item" : item}, callback=self.parse_image)

            self.page += 100
            yield scrapy.Request(self.base_url.format(self.page), callback=self.parse)

    def parse_image(self, response):
        item = response.meta["item"]

        filename = item['nickname'] + ".jpg"
        # 图片的路径
        item["image_path"] = "/Users/Power/lesson_python/_20_0829/day11/Douyu/Douyu/Image/" + filename

        with open(item["image_path"], "wb") as f:
            f.write(response.body)

        yield item


