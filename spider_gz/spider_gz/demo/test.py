# -*- coding: utf-8 -*-
import time, os
from multiprocessing import Pool


# def run(n):
#     time.sleep(1)
#     print('Run child process %s (%s)...' % (n * n, os.getpid()))
#
#
# if __name__ == "__main__":
#     testFL = [1, 2, 3, 4, 5, 6]
#     print('顺序执行:')  # 顺序执行(也就是串行执行，单进程)
#     starttime = time.time()
#     for n in testFL:
#         run(n)
#     pool = Pool(10)  # 创建拥有10个进程数量的进程池
#     pool.map(run, testFL)
#     pool.close()  # 关闭进程池，不再接受新的进程
#     pool.join()  # 主进程阻塞等待子进程的退出



