#coding:utf-8

import scrapy

from scrapy_redis.spiders import RedisSpider

from ..items import AqiItem

try:
    # python2
    from urllib import unquote
except:
    # python3
    from urllib.parse import unquote


#class AqiSpider(scrapy.Spider):
class AqiRedisSpider(RedisSpider):

    name = "aqi_redisspider"

    allowed_domains = ["aqistudy.cn"]

    base_url = "https://www.aqistudy.cn/historydata/"

    #start_urls = [base_url]
    redis_key = "aqiredisspider:start_urls"

    # def start_requests(self):
    #     pass

    def parse(self, response):
        """ 解析城市列表页的响应，提取所有城市的链接并发送请求, 返回的响应交给parse_month提取月份链接"""

        # 所有城市的链接： 384个
        city_link_list = response.xpath("//div[@class='all']//a/@href").extract()
        # 所有的城市名：384
        city_name_list = response.xpath("//div[@class='all']//a/text()").extract()
        # print("--" * 30)
        # print(city_link_list)
        # 发送每个城市链接的请求，并通过meta传递城市名
        for city_link, city_name in zip(city_link_list, city_name_list)[10:15]:
            yield scrapy.Request(url=self.base_url + city_link, meta={"city" : city_name}, callback=self.parse_month)


    def parse_month(self, response):
        """ 处理每个城市的响应，提取所有月份的链接并发送请求，返回的响应交给parse_day提取每一天的数据"""

        month_link_list = response.xpath("//tbody//td[1]/a/@href").extract()
        # print("--" * 30)
        # self.logger.info(month_link_list)
        # 发送每个月的链接的请求，并通过meta传递城市名
        for month_link in month_link_list:
            yield scrapy.Request(url=self.base_url + month_link, meta=response.meta, callback=self.parse_day)


    def parse_day(self, response):

        # 通过url地址取出城市名
        url = response.url
        city = unquote(url[url.find("=")+1:url.find("&")]).decode("utf-8")

        # 通过meta取出城市名
        # city = response.meta["city"]

        tr_list = response.xpath("//tbody/tr")
        tr_list.pop(0)

        for tr in tr_list:
            item = AqiItem()

            item["city"] = city
            # 空气数据的日期
            item['date'] = tr.xpath("./td[1]/text()").extract_first()
            # 空气质量指数
            item['aqi'] = tr.xpath("./td[2]/text()").extract_first()
            # 空气质量等级
            item['level'] = tr.xpath("./td[3]/span/text()").extract_first()
            # pm2.5
            item['pm2_5'] = tr.xpath("./td[4]/text()").extract_first()
            # pm10
            item['pm10'] = tr.xpath("./td[5]/text()").extract_first()
            # 二氧化硫
            item['so2'] = tr.xpath("./td[6]/text()").extract_first()
            # 一氧化碳
            item['co'] = tr.xpath("./td[7]/text()").extract_first()
            # 二氧化氮
            item['no2'] = tr.xpath("./td[8]/text()").extract_first()
            # 臭氧
            item['o3'] = tr.xpath("./td[9]/text()").extract_first()


            yield item

#https://www.aqistudy.cn/historydata/daydata.php?city=%E8%9A%8C%E5%9F%A0&month=2015-01
