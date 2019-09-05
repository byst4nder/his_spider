#!/usr/bin/env python
# -*- coding: utf-8 -*
import itertools as its
import os

rootPath = os.path.dirname(os.path.realpath(__file__))
words = "1234568790"
r = its.product(words, repeat=8)
j = 0
next_num = ''
file_name = rootPath + "/pass.txt"
dic = open(file_name, "a")  # 50
for i in r:
    now_num = int(i[1])
    if now_num != next_num:
        dic.close()
        j += 1
        file_name = rootPath + "pass_v" + str(j) + ".txt"
        dic = open(file_name, "a")
    content = "".join(i) + "\r\n"
    dic.write(content)
    next_num = now_num
