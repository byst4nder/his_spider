#!/usr/bin/env python
# -*- coding: utf-8 -*
import itertools as its
import os

rootPath = os.path.dirname(os.path.realpath(__file__))

words = "1234568790"
r = its.product(words, repeat=8)

j = 0
end_count = 2
next_num = 1
file_name = rootPath + "/pass.txt"
file_name = file_name.replace("\\", '/')
dic = open(file_name, 'a')
file_name_new = None
for i in r:
    now_num = int(i[1])
    if now_num != next_num:
        j += 1
        file_name = "pass_v" + str(j) + ".txt"
        file_name_new = file_name.replace("\\", '/')
        dic = open(file_name_new, 'a')
    content = "".join(i) + "\n"
    dic.write(content)
    next_num = now_num





