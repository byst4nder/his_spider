import random
import re

import redis


class BaiDuoCi:
    def __init__(self):
        self.tag = True
        self.rdp = redis.ConnectionPool(host='47.104.96.186', port=6379, db=1)
        self.redis_client = redis.StrictRedis(connection_pool=self.rdp)
        self.key = 'baidu1'

    def run(self):
        self.redis_client.set(self.key, "111")
        # a = self.redis_client.get(self.key)
        # print(a)

    def test(self):
        html_str = """<link rel="stylesheet" href="https://csdnimg.cn/release/phoenix/template/css/ck_htmledit_views-3019150162.css">"""
        regex = re.compile('href="([\w\W]+?)"')
        data_list = regex.findall(html_str)


if __name__ == '__main__':
    baidu = BaiDuoCi()
    baidu.run()