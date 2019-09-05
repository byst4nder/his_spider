
"""
百度资讯的爬取
"""
import json
import os
import random
import time
import chardet
import redis
import requests
from bs4 import BeautifulSoup

USER_AGENT_LIST = [
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/22.0.1207.1 Safari/537.1",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
            "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1090.0 Safari/536.6",
            "Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/19.77.34.5 Safari/537.1",
            "Mozilla/5.0 (Windows NT 6.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.36 Safari/536.5",
            "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
            "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
        ]

rootPath = os.path.dirname(os.path.realpath(__file__))
stopwords = {}.fromkeys(['a8', '2018', '2018', '201820190605A', '99', 'mzz', '10', '...', '..._',
                         '365', 'PK', 'mg', '888', 'mod', 'Realm', 'of', 'Swords'])


class BaiDuoCi:
    def __init__(self):
        self.headers = {
            "User-Agent": random.choice(USER_AGENT_LIST)
        }
        self.url = 'https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=baidu&wd={}&oq=%25E4%25BF%25A1%25E5%25BD%25A9%25E5%25A8%25B1%25E4%25B9%2590%25E8%25AE%25A1%25E5%2588%2592&rsv_pq=f09e43ea00007117&rsv_t=20e6r4KiP2DalRpXsk7rmobX4atwzmmX5QXndXbmvpcK52WEA39eWqfNGho&rqlang=cn&rsv_enter=1&rsv_n=2&rsv_sug3=1&bs=%E4%BF%A1%E5%BD%A9%E5%A8%B1%E4%B9%90%E8%AE%A1%E5%88%92'
        self.tag = True
        self.rdp = redis.ConnectionPool(host='127.0.0.1', port=6379)
        self.redis_client = redis.StrictRedis(connection_pool=self.rdp)
        self.key = 'baiduciv5'
        self.key_temp = 'baiduciv5_temp'
        self.key_queue = 'baiduciv5_queue'

    def request(self, wd):
        url = self.url.format(wd)
        resp = requests.get(url=url, headers=self.headers)
        html_encode = resp.content
        items = []
        try:
            detect_result = chardet.detect(html_encode)
            encoding = detect_result["encoding"]
            html = str(html_encode, encoding)
            soup = BeautifulSoup(html, 'html.parser')
            items = soup.select('th a')
        except Exception as e:
            print('失败原因：{}'.format(e))
            return items
        else:
            return items

    def is_pass(self, title):
        """ 判断当前标题是否符合要求"""
        list_filter = ['金城', '黑娱', '娱乐', '人娱', '庆彩', '经纬', '经纬', 'cnc', 'sky', 't6', '亿博', 'k彩', '捕鱼', '帝一', '新火', '吉祥',
                       '如意', '宝盈', '宝马娱乐', '云顶', '娱乐', '平台']
        result = False
        for filter_str in list_filter:
            if title.__contains__(filter_str):
                result = True
        return result

    def save_title(self, data):
        now_time = str(time.strftime('%Y-%m-%d-%H'))
        file_name = rootPath + '/' + now_time + "百度相关搜索词5.text"
        with open(file_name, 'a+', encoding="utf-8") as f:
            data_list = data.keys()
            print("输出关键字{}条".format(len(data_list)))
            f.write('总计{}条关键词'.format(len(data_list)) + '\r\n')
            for i in data_list:
                f.write(i.decode() + '\r\n')

    def run(self):
        # keywords_list = ['金城娱乐平台', '黑娱乐平台', '娱乐平台大全', '人娱乐平台', '庆彩娱乐平台', '经纬娱乐平台', 'cnc娱乐平台', 'sky娱乐平台',
        #                  '宝马娱乐平台', 't6娱乐平台', '亿博娱乐平台', 'k彩娱乐平台', '捕鱼娱乐平台', '帝一娱乐平台', '新火娱乐平台', '吉祥娱乐平台', '如意娱乐平台',
        #                  '宝盈娱乐平台',
        #                  '云顶娱乐平台']
        while True:
            keywords = self.redis_client.lpop(self.key_queue)
            if not keywords:
                break
            print("当前搜索的词汇为：{}".format(keywords.decode()))
            list_title = self.request(keywords.decode())
            for i in list_title:
                title = i.text.strip()
                if self.is_pass(title):
                    self.redis_client.hset(self.key, title, title)
                    self.redis_client.rpush(self.key_queue, title)

        # result = self.redis_client.hgetall(self.key)
        # self.save_title(result)


if __name__ == '__main__':
    baidu = BaiDuoCi()
    baidu.run()


