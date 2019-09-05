# coding:utf-8
"""
360问答数据的爬取
"""
import json
import os
import random
import re
import time

import redis
import requests


# '娱乐', '棋牌',  '体育', '登陆', '注册', '国际', '在线', '平台'
# 违禁词语 彩票 CP   博彩BC  时时彩 SSC

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

class Request360:
    """ 360问答数据的爬取 """
    spider_name = '360问问'

    def __init__(self):

        self.headers = {
            "User-Agent": random.choice(USER_AGENT_LIST),
            'Host': 'wenda.so.com',
            'Referer': 'https://wenda.so.com',
        }
        self.url_code ='https://wenda.so.com/search/?q={}&pn={}'
        self.url_article = 'https://wenda.so.com/q/{}'
        self.rdp = redis.ConnectionPool(host='127.0.0.1', port=6379)
        self.redis_client = redis.StrictRedis(connection_pool=self.rdp)
        self.key = 'request_answer_360'
        self.key_hash = '360_article'
        self.not_get = '奇虎360旗下最大互动问答社区'
        self.not_get_data = '未获取到当前文章内容'
        self.no_exise = '微知识，大帮助——360问答是一个互动知识分享社区，网友们可以把自己工作、生活中遇到的问题提交给360问答，360问答会匹配到最适合的回答者来解答问题。解决后的问题可以被使用360搜索的其他用户搜索到，帮助更多人解决类似的问题。'
        self.no_exise_content = '当前页面内容已经不存在!!!!'

    def request(self, q, pn):
        url = self.url_code.format(q, pn)
        try:
            time_list = [5, 6, 7, 8]
            time.sleep(random.choice(time_list))
            resp = requests.get(url=url, headers=self.headers)
            result = resp.content.decode('utf-8')
        except Exception as e:
            print('失败原因为：{}'.format(e))
            return None
        else:
            return result

    def receive_code(self, list_code):
        for i in list_code:
            self.redis_client.hset(self.key_hash, i, i)
            self.redis_client.rpush(self.key, i)

    def parse_code(self, html_str):
        regex = re.compile("""<a index="" aId="\d+" target="_blank" """)
        result = regex.findall(html_str)
        result_list = list()
        for i in result:
            code = re.search(r'\d+', i)
            if code:
                result_list.append(code.group())
        return result_list

    def request_content(self):
        while True:
            code = self.redis_client.lpop(self.key)
            if not code:
                break
            else:
                return self.get_article(code.decode())

    def filter_article(self, data):
        """ 过滤文章 """
        filer_list = ['www', 'http', 'com']
        for i in filer_list:
            if data.__contains__(i):
                return False
        return True

    def get_article(self, i):
        url = self.url_article.format(int(i))
        # print('请求的url：', url)
        self.headers['Referer'] = url
        try:
            resp = requests.get(url=url, headers=self.headers)
            time_list = [5, 6, 7, 8]
            time.sleep(random.choice(time_list))
            html_str = resp.content.decode()
            regex = re.compile("""meta name="description" content="([\d\D]+?)" />""")
            data_list = regex.findall(html_str)
            print('data:', data_list)
        except Exception as e:
            print("失败原因为：{}".format(e))
            return None
        else:
            if data_list:
                data = data_list[0]
                if self.no_exise in data:
                    return None
                if self.not_get in data:
                    return None
                data.replace('彩票', 'CP')
                data.replace('博彩', 'BC')
                data.replace('时时彩', 'SSC')
                if self.filter_article(data):
                    return re.sub('\s\s+', ';   ', data)
                else:
                    return None
            else:
                return None

    def save_content(self, data, name):
        date = time.strftime('%Y-%m-%d')
        file_path = rootPath + '/spider_data'
        if not os.path.exists(file_path):
            os.makedirs(file_path)
        file_name = file_path + '/' + self.spider_name + '_' + name + '_' + date + ".text"
        with open(file_name, 'a+', encoding='utf-8') as f:
            f.write(data + "\r\n")

    def break_rank(self, list_data):
        """ 打乱有序数组"""
        list_index = [i for i in range(len(list_data))]
        random.shuffle(list_index)
        list_new = list()
        for index in list_index:
            list_new.append(list_data[index])
        return list_new

    def carve_up(self, sentence):
        """ 切割句子"""
        sentence_len = len(sentence)
        sentence_list = list()
        if sentence_len > 150:
            cut_list = [i + '。 ' for i in sentence.split("。") if i]
            a = ''
            b = ''
            total_str = ''
            for i in cut_list:
                a += i
                total_str += i
                if len(a) > 150:
                    b += a
                    sentence_list.append(a)
                    a = ''
            end_str = total_str.replace(b, '')
            if end_str:
                sentence_list.append(end_str)
        else:
            sentence_list.append(sentence)
        return sentence_list

    def run(self):  # '娱乐', '棋牌', '体育', '登陆', '注册', '国际', '在线', '平台'
        cate_list = ['娱乐', '棋牌', '体育', '登陆', '注册', '国际', '在线', '平台']
        start_page = 151
        total_pages = 201
        while True:
            code = self.redis_client.lpop(self.key)
            if not code:
                break
        for cate in cate_list:
            for i in range(start_page, total_pages):
                article_list = list()
                result = self.request(cate, i)
                if result:
                    list_code = self.parse_code(result)
                    self.receive_code(list_code)
                    article = self.request_content()
                    if article:
                        article_list_temp = self.carve_up(article)
                        article_list.extend(article_list_temp)
                    for article in self.break_rank(article_list):
                        print("当前文章长度为：{}".format(len(article)), article)
                        self.save_content(article, cate)


if __name__ == '__main__':
    res_obj = Request360()
    res_obj.run()