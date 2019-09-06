# 先安装 ： pip install selenium
# 还要安装PhantomJS

# 导入webdriver
from selenium import webdriver

# 创建一个PhantomJS浏览器对象
driver = webdriver.PhantomJS()

# 通过浏览器打开指定页面
driver.get("http://www.baidu.com/")
# 保存当前页面截图
driver.save_screenshot("baidu.png")
# 获取渲染后的完整源码（Unicode编码）
html = driver.page_source
print(html)


## 通常情况下，在爬虫里使用Selenium用来获取网页渲染后的源码，并使用xpath解析数据
from lxml import etree
html_obj = etree.HTML(html)


## 也可以通过Selenium自带的页面元素查找方法，进行数据提取

# 通过浏览器 打开baidu 首页
driver.get("http://www.baidu.com/")

# 查找当前标签页里的 id=u1 结点
driver.find_element_by_id("u1")
# 如果找不到则抛出NoSuchElementException 异常
#driver.find_element_by_id("usss1")

# 查找当前标签页里的 id=u1 结点，并获取该标签页的文本内容
driver.find_element_by_id("u1").text
driver.find_element_by_id("u1").get_attribute("href")

# 查找当前标签页里的 name=tj_trnews 结点，并点击 （通过点击 a 标签）
driver.find_element_by_name("tj_trnews").click() # 找到指定的a标签，模拟鼠标点击事件
driver.save_screenshot("baidu_news.png")

# 查找当前标签页里的 id=ww 结点，并输入文本 （通过输入文本 为input 标签）
driver.find_element_by_id("ww").send_keys(u"华为")
driver.save_screenshot("baidu_news_huawei.png")

# 查找当前标签页里的 class=btn 结点，并点击 （通过点击 a 标签）
driver.find_element_by_class_name("btn").click()
driver.save_screenshot("baidu_news_huawei_click.png")

# 通过xpath查找指定的a标签，并点击
driver.find_element_by_xpath("//div[@id='1']/h3/a").click()
driver.save_screenshot("baidu_news_huawei_click_a.png")

# 获取当前浏览器所有的标签页句柄
driver.window_handles
# 根据标签页窗口句柄，切换到指定的标签页
driver.switch_to_window(driver.window_handles[1])

driver.save_screenshot("baidu_news_huawei_click_a2.png")

# 获取当前标签页的网页源码
print(driver.page_source)

# 获取当前标签页的 url地址
driver.current_url
# 获取当前标签页的 网页标题
driver.title
print(driver.title)

# 关闭当前标签页（窗口句柄会减少）
driver.close() # 关闭标签页
driver.window_handles

# 关闭浏览器（不能再使用PhantomJS浏览器）
driver.quit()
driver.window_handles

%hist -f webdriver.py


