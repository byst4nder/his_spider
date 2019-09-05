#!/usr/bin/env python
# -*- coding: utf-8 -*
import itertools as its

words = "1234568790"
r = its.product(words, repeat=8)
file_name = "pass.txt"
dic = open(file_name, "a")  # 50
j = 0
end_count = 2
next_num = ''
for i in r:
    now_num = str(i)[0:end_count]
    if now_num != next_num:
        j += 1
        file_name = "pass_v" + str(j) + ".txt"
        dic = open(file_name, "a")
    dic.write("".join(i))
    dic.write("".join("\n"))
    next_num = now_num
dic.close()


