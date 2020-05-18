
#测试爬取到的代理是否有效

import requests
from multiprocessing import Process
from multiprocessing import Queue
from multiprocessing import Pool
from multiprocessing import Manager
import pymongo
import os
import time
import random

client = pymongo.MongoClient()
db = client['proxy']

def get_ip(queue):
    print('get process id is :',os.getpid())
    i = 0
    for ip in db.xicidaili.find():
        if ip:
            queue.put(ip['ip'])
            i += 1
        # print('put',ip['ip'],queue.qsize())
        print(i)

def test_ip(queue):
    print('test process id is :',os.getpid())
    i = 0
    if not queue.empty():
        while True:
            ip = queue.get(True)        #读取进程在阻塞
            i += 1
            # print('get',ip,i)
            test_douban(i,ip)

def test_douban(i,ip):
    # print('test douban process :',os.getpid(),ip,i)
    headers = {
        'Accept':'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
        'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
    }

    proxy = {
        'http':'http://' + ip
    }
    # url = 'https://www.douban.com'
    # print(proxy)
    url = 'http://www.tsinghua.edu.cn/publish/newthu/index.html'
    try:
        response = requests.get(url,headers = headers,proxies = proxy,allow_redirects=False)
        save_to_mongo(ip)
    except Exception as e:
        print(e)

def save_to_mongo(base_ip):
    ip = {
        'ip':base_ip
    }
    if db['proxycanuse'].insert(ip):
        print('存储成功',ip)
        return True
    else:
        print('存储失败')
        return False


def main():
    queue = Manager().Queue(maxsize=100)
    get_ip_process = Process(target=get_ip,args=(queue,))
    get_ip_process.start()
    time.sleep(2)       #让put_ip进程先把queue填充满再启动test进程
    pool = Pool(3)
    for i in range(5):
        pool.apply_async(test_ip,(queue,))
    pool.close()
    pool.join()

    print(get_ip_process.is_alive())
    print('main process id is :',os.getpid())
    print('_____________')

if __name__ == '__main__':
    main()



# import os
# import multiprocessing
#
#
# def getpid(name):
#     print('the %s process id is : %s' % (name, os.getpid()))
#
# def main():
#     pid = os.getpid()
#     p = multiprocessing.Process(target=getpid, args=('name',))
#     p.start()
#     p.join()
#     print('main pid is :', pid)
#
#
# if __name__ == '__main__':
#     main()