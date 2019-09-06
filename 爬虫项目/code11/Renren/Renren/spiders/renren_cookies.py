#coding:utf-8


import scrapy

class RenrenSpider(scrapy.Spider):
    name = "renren_cookies"
    #allowed_domains = []
    start_urls = [
        "http://www.renren.com/410043129/profile",
        "http://www.renren.com/965999739/profile"
    ]

    headers = {
        "Accept" : "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
        # scrapy和requests都可以处理gzip解压缩
        "Accept-Encoding" : "gzip, deflate",
        "Accept-Language" : "zh-CN,zh;q=0.9,en;q=0.8",
        "Connection" : "keep-alive",
        "Host" : "www.renren.com",
        "Referer" : "http://zhibo.renren.com/top",
        "Upgrade-Insecure-Requests" : "1",
        "User-Agent" : "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
    }

    # scrapy里的cookie必须保存在字典里，传递给请求的cookies参数
    cookies = {
        "anonymid" : "j7wsz80ibwp8x3",
        "__utma" : "151146938.965194798.1536382787.1536382787.1536382787.1",
        "__utmz" : "151146938.1536382787.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none)",
        "_r01_" : "1",
        "ln_uact" : "mr_mao_hacker@163.com",
        "ln_hurl" : "http://hdn.xnimg.cn/photos/hdn321/20180921/1130/main_Affp_c9750000af421986.jpg",
        "_de" : "BF09EE3A28DED52E6B65F6A4705D973F1383380866D39FF5",
        "depovince" : "GUZ",
        "jebecookies" : "960ebd97-113f-4ef0-980e-46b998a974f7|||||",
        "ick_login" : "fc840a38-dec1-47f8-b531-b46f23865f82",
        "p" : "cfe8f3bcab9859d9425acc90022b30ca9",
        "first_login_flag" : "1",
        "t" : "159215a27330bc6509b0d3f5be5760109",
        "societyguester" : "159215a27330bc6509b0d3f5be5760109",
        "id" : "327550029",
        "xnsid" : "21cb9643",
        "loginfrom" : "syshome",
        "ch_id" : "10016"
    }

    def start_requests(self):
        for url in self.start_urls:
            yield scrapy.Request(url, cookies=self.cookies, headers=self.headers, callback=self.parse)

    def parse(self, response):
        title = response.xpath("//title/text()").extract_first()

        with open(title + ".html", "w") as f:
            f.write(response.body)
