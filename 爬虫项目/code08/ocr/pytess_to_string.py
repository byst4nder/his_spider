#!/usr/bin/env python
# -*- coding:utf-8 -*-

from pytesseract import image_to_string
from PIL import Image


def png_to_str():
    image_data = Image.open("排序算法.png")
    text = image_to_string(image_data, lang='chi_sim')
    print(text)

png_to_str()
