#coding:utf-8


from ..items import TencentItem, PositionItem
#from ..settings import BOT_NAME
from datetime import datetime

from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule


class TencentCrawlSpider(CrawlSpider):
    name = "tencent_crawl"
    allowed_domains = ["hr.tencent.com"]

    # start_urls里的响应没有解析方法，但是会经过每一个Rule进行链接提取
    start_urls = ["https://hr.tencent.com/position.php?&start=0"]


    rules = [
        # 1. LinkExtractor用来提取所有的 列表页的链接 并发送链接请求
        # 2. callback 表示链接请求的响应,交给callback解析(注意不能写parse)
        # 3. follow = True，表示响应会继续提取新的链接并发送；follow=False表示响应不需要提取链接

        # 提取列表页
        Rule(LinkExtractor(allow=r"start=\d+"), callback = "parse_page", follow=True),
        # 提取详情页
        Rule(LinkExtractor(allow=r"position_detail\.php\?id=\d+"), callback="parse_detail", follow=False)
    ]

    def parse_page(self, response):

        """ 默认列表页的解析方法"""
        node_list = response.xpath("//tr[@class='even'] | //tr[@class='odd']")

        for node in node_list:
            item = TencentItem()
            # 职位名
            item['position_name'] = node.xpath("./td[1]/a/text()").extract_first()
            # 详情链接
            item['position_link'] = "https://hr.tencent.com/" + node.xpath("./td[1]/a/@href").extract_first()
            # 职位类别
            item['position_type'] = node.xpath("./td[2]/text()").extract_first()
            # 招聘人数
            item['people_number'] = node.xpath("./td[3]/text()").extract_first()
            # 工作地点
            item['work_location'] = node.xpath("./td[4]/text()").extract_first()
            # 发布时间
            item['publish_times'] = node.xpath("./td[5]/text()").extract_first()
            # 爬虫名（数据源）
            item['spider'] = self.name
            # 抓取时间
            item['time'] = str(datetime.now())

            yield item



    def parse_detail(self, response):
        """ 解析详情页的响应内容"""
        #item = response.meta['item']

        item = PositionItem()
        item['position_zhize'] = response.xpath("//ul[@class='squareli']")[0].xpath("./li/text()").extract()
        item['position_yaoqiu'] = response.xpath("//ul[@class='squareli']")[1].xpath("./li/text()").extract()

        yield item
