#coding:utf-8

try:
    # python2
    from urllib import unquote
except:
    # python3
    from urllib.parse import unquote

from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import Rule, CrawlSpider

class AqiCrawlSpider(CrawlSpider):

    name = "aqi_crawlspider"

    allowed_domains = ["aqistudy.cn"]

    base_url = "https://www.aqistudy.cn/historydata/"
    start_urls = [base_url]


    rules = [
        # 提取每个城市的链接，并发送请求，返回的响应会继续提取新的链接
        Rule(LinkExtractor(allow=r"monthdata\.php\?city="), follow=True),
        # 提取每个月份的链接，并发送请求，返回的响应会通过parse-day提取该月的所有天数据，并不在提取新的链接
        Rule(LinkExtractor(allow=r"daydata\.php\?city="), callback="parse_day")
    ]

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
