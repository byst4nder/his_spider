#coding:utf-8

# python的测试模块
import unittest
import time

from selenium import webdriver
from lxml import etree

import pymongo

class DouyuSpider(unittest.TestCase):
    # __init__(self) 构造方法（创建对象时调用）
    def setUp(self):
        #self.s = "hello world"
        # 创建一个PhantomJS浏览器
        self.driver = webdriver.PhantomJS()
        # 计数器，统计总共主播人数
        self.count = 0

        # 创建MongoDB数据库链接和集合，用来保存数据
        self.client = pymongo.MongoClient()
        self.collection = self.client.douyu.directory

    def testDouyu(self):
        # 先通过浏览器 打开页面
        self.driver.get("https://www.douyu.com/directory/all")

        while True:
            # 获取当前页面的Unicode源码字符串
            html = self.driver.page_source
            # 通过xpath提取所有主播结点
            html_obj = etree.HTML(html)
            node_list = html_obj.xpath('//div[@id="live-list-content"]//div[@class="mes"]')

            # 迭代每个主播结点，获取主播信息，并保存到MongoDB中
            for node in node_list:
                item = {}
                # 房间名：
                item["room"] = node.xpath(".//h3[@class='ellipsis']/text()")[0].strip()
                # 分类：
                item["tag"] = node.xpath(".//span[@class='tag ellipsis']/text()")[0].strip()
                # 主播名：s
                item['name'] = node.xpath(".//span[@class='dy-name ellipsis fl']/text()")[0].strip()
                try:
                    # 观众人数：
                    item['people'] = node.xpath(".//span[@class='dy-num fr']/text()")[0].strip()
                except:
                    item['people'] = 0

                print(item)
                # 写入每个主播的item数据到MongoDB中
                self.collection.insert(item)
                self.count += 1

            #if 判断当前页面是否是最后一页，如果是:
            # 在页面里查找指定字符串，如果没找到，返回-1，表示没到最后一页
            # 如果找到了，返回是下标值（不是-1）,表示到了最后一页，可以退出循环
            if html.find("shark-pager-disable-next") != -1:
                break
            else:
                # 如果没到最后一页，则点击下一页链接，获取新的页面数据
                self.driver.find_element_by_class_name("shark-pager-next").click()
                time.sleep(1)


    #def __del__(self) 析构方法（销毁对象时调用）
    def tearDown(self):
        print("over")
        print(self.count)
        # 程序结束，关闭浏览器
        self.driver.quit()


if __name__ == '__main__':
    unittest.main()
    #spider = DouyuSpider()
