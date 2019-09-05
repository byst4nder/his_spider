import requests
import random
import time


class Bilibili:
    def __init__(self):
        self.headers = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; WOW64)AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.87 UBrowser/6.2.3964.2 Safari/537.36",
        }
        self.url = 'https://www.bilibili.com/video/av36210854/?p=7'

    def run(self):
        html = requests.get(url=self.url, headers=self.headers)
        print(html.json())
        # html_json = html.json()
        # for info in html_json['data']["items"]:
        #     print(info["item"]["description"])


if __name__ == '__main__':
    bilibili = Bilibili()
    bilibili.run()