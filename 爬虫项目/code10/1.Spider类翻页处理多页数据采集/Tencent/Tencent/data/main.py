#!/usr/bin/env python
# -*- coding:utf-8 -*-

from scrapy import cmdline

#cmdline.execute(["scrapy", "crawl", "tencent"])
cmdline.execute("scrapy crawl tencent3 -o tencent.json".split())
