#coding:utf-8

import random
import requests

from fake_useragent import UserAgent
from settings import USER_AGENT_LIST

class RandomUserAgentMiddleware(object):
    def __init__(self):
        self.ua_obj = UserAgent()

    def process_request(self, request, spider):
        #user_agent = random.choice(USER_AGENT_LIST)
        user_agent = self.ua_obj.random
        request.headers["User-Agent"] = user_agent
        print('---' * 10)
        print(request.headers)
        # 在中间件里不需要写return操作
        # return request

class RandomProxyMiddleware(object):
    def __init__(self):
        self.proxy_url = "http://kps.kdlapi.com/api/getkps/?orderid=914194268627142&num=1&pt=1&sep=1"
        # 获取代理服务器里提供的proxy
        self.proxy_list = [requests.get(self.proxy_url).content]
        self.count = 0

    def process_request(self, request, spider):
        if self.count < 20:
            proxy = random.choice(self.proxy_list)
            #http://47.99.65.91:16818
            # http://maozhaojun:ntkn0npx@47.99.65.91:16818
            request.meta['proxy'] = "http://maozhaojun:ntkn0npx@" + proxy

            self.count += 1
        else:
            self.proxy_list = [requests.get(self.proxy_url).content]
            self.count = 0
