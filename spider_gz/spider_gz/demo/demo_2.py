# coding=utf-8
from selenium import webdriver
import time

browser = webdriver.Firefox()  # 启动Firefox浏览器
url1 = 'http://www.baidu.com/s?wd=马尔代夫'  # 马尔代夫的百度搜索页
url2 = 'http://scrapy.org/'  # 测试页面1
url3 = 'http://news.yahoo.com/air-strike-libyan-city-misrata-clashes-near-oil-135039996.html'  # 测试页面2
browser.get(url1)  # 打开页面
browser.maximize_window()  # 浏览器窗口最大化

res = browser.find_elements_by_xpath("//h3[@class='t c-gap-bottom-small']")  # 使用xpath查找页面中的h3元素

for r in res:  # 页面中h3元素有多个，所以需要遍历
    t = r.find_element_by_xpath("a")  # h3元素下的a标签
    print('%s - %s' % (t.text, type(t.text)))  # 打印a标签的标题以及文本格式
    if u'途牛' in t.text:  # 判断标题中是否有途牛，如果有则点击
        print("yes")  # 判断结果
        t.click()  # 点击这个a链接

print(len(res)) # 打印h3标签的总量