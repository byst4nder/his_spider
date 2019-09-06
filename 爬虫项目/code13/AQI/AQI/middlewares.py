#coding:utf-8

import time

from retrying import retry

from selenium import webdriver

# scrapy所有响应对象的父类
from scrapy.http import HtmlResponse


class AqiSeleniumMiddleare(object):

    def __init__(self):
        # 默认有界面
        #self.driver = webdriver.Chrome()

        # 修改chrome的配置为无界面模式
        self.options = webdriver.ChromeOptions()
        self.options.add_argument("--headless")

        # 创建chrome浏览器，使用该配置
        self.driver = webdriver.Chrome(options=self.options)



    @retry(stop_max_attempt_number=20, wait_fixed=200)
    def retry_load_page(self, request, spider):
        try:
            # 在页面查找指定数据，如果数据出现，则没有异常，程序正常执行
            # 在页面查找指定数据，如果数据没出现，会出现异常，异常会被retry捕获
            #   retry总共捕获20次，如果20次之后还有异常，则异常抛出给函数调用的地方
            self.driver.find_element_by_xpath("//tbody/tr[2]/td[1]")
        except:
            spider.logger.debug("{} retry {} times".format(request.url, self.count))
            self.count += 1
            # 手动抛出异常，让retry工作
            raise Exception(" {} page load failed".format(request))



    def process_request(self, request, spider):
        # 判断当前请求是否是动态页面，
        # 如果if成立表示是动态页面，则自定义构建响应；如果if不成立则表示是静态页面，则让下载器处理请求返回响应
        if "monthdata" in request.url or "daydata" in request.url:
            url = request.url

            # 通过chrome发送请求
            self.driver.get(request.url)

            # 隐式等待：固定等待时间让浏览器渲染（但是有效率问题）
            #time.sleep(2)

            # 显示等待：判断页面需要的数据是否渲染成功，如果渲染成功则构建响应返回；没有渲染成功则一直等待。
            try:
                self.count = 1
                self.retry_load_page(request, spider)

                # 获取网页源码字符串（Unicode编码）
                html = self.driver.page_source

                # 构建响应对象，必须提供下面四个属性：
                # 响应的url地址：response.url
                # 响应网页原始编码字符串：response.body
                # 响应body字符串的编码response.encoding
                # 响应对应的请求对象：response.request
                response = HtmlResponse(
                    url=self.driver.current_url,
                    body=html.encode("utf-8"),
                    encoding="utf-8",
                    request=request
                )

                # 在下载中间件里自行处理请求，不通过下载器处理
                return response
            except Exception as e:
                spider.logger.error(e)
                # 如果retry 20次后还有异常，则接收异常并输出到日志里，通过返回request交给引擎-调度器再次重新入队列
                return request

    def __del__(self):
        self.driver.quit()
