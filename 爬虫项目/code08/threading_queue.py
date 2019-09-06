
import requests
import threading
from queue import Queue

url_queue = Queue()

def get_url_queue():

    for page in range(0, 226, 25):
        url = "https://movie.douban.com/top250?start=" + str(page)
        # 每次put，队列的计数器会自增1
        url_queue.put(url)


def func():
    while True:
        # 只get取值，计数器不减1
        url = url_queue.get()

        print("[INFO] : send request {}".format(url))
        response = requests.get(url)

        # 同时通过 task_done可以实现计数器减1
        url_queue.task_done()


def main():
    # 先获取所有的url地址到队列中
    get_url_queue()

    for _ in range(3):
        thread = threading.Thread(target=func)
        thread.setDaemon(True) # 将子线程设置为守护线程（主线程结束时，子线程会立刻被回收）
        thread.start()
    # 如果这个队列的计数器不为0，主线程会阻塞；
    # 当队列的计数器为0时，主线程会继续执行。
    url_queue.join()


if __name__ == '__main__':
    main()
