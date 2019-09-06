# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# http://doc.scrapy.org/en/latest/topics/items.html

import scrapy


# class Tencent2Item(scrapy.Item):

#     ### 列表页的 6个字段
#     # 职位名
#     position_name = scrapy.Field()
#     # 详情链接
#     position_link = scrapy.Field()
#     # 职位类别
#     position_type = scrapy.Field()
#     # 招聘人数
#     people_number = scrapy.Field()
#     # 工作地点
#     work_location = scrapy.Field()
#     # 发布时间
#     publish_times = scrapy.Field()


#     ### 详情页的 2个字段
#     # 工作职责：
#     position_zhize = scrapy.Field()
#     # 工作要求
#     position_yaoqiu = scrapy.Field()


#     # 爬虫名（数据源）
#     spider = scrapy.Field()
#     # 抓取时间
#     time = scrapy.Field()



class TencentItem(scrapy.Item):

    ### 列表页的 6个字段
    # 职位名
    position_name = scrapy.Field()
    # 详情链接
    position_link = scrapy.Field()
    # 职位类别
    position_type = scrapy.Field()
    # 招聘人数
    people_number = scrapy.Field()
    # 工作地点
    work_location = scrapy.Field()
    # 发布时间
    publish_times = scrapy.Field()
    # 爬虫名（数据源）
    spider = scrapy.Field()
    # 抓取时间
    time = scrapy.Field()



class PositionItem(scrapy.Item):
    ### 详情页的 2个字段
    # 工作职责：
    position_zhize = scrapy.Field()
    # 工作要求
    position_yaoqiu = scrapy.Field()


    # 爬虫名（数据源）
    spider = scrapy.Field()
    # 抓取时间
    time = scrapy.Field()
