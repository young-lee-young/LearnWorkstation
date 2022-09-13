

#爬取西刺代理IP，保存到数据库

import requests
from lxml import etree
import time
import random
import pymongo
from multiprocessing import Pool

headers = {
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}

client = pymongo.MongoClient('localhost')
db = client['proxy']

def main():
    page = get_page()
    print(page)
    pool = Pool(processes=4)
    for i in range(1,page):
        pool.apply_async(get_html,(i,))
    pool.close()
    pool.join()

def get_page():
    base_url = 'http://www.xicidaili.com/nn'
    response = requests.get(base_url, headers=headers)
    print(response.status_code)
    time.sleep(random.random()*3)
    if response.status_code == 200:
        page_source = etree.HTML(response.text)
        i = page_source.xpath('//*[@id="body"]/div[2]/a[last()-1]/text()')[0]
        return int(i)
    else:
        print('获取页数失败')
        return None

def get_html(i):
    url = 'http://www.xicidaili.com/nn/{}'.format(i)
    print('正在解析', url)
    response = requests.get(url, headers=headers)
    print(response.status_code)
    print(response.text)
    sleep_time = random.randint(4, 8)
    time.sleep(sleep_time)
    if response.status_code == 200:
        parse_html(response.text)
    else:
        print('获取HTML失败')
        return None

def parse_html(html):
    page_source = etree.HTML(html)
    trs = page_source.xpath('//*[@id="ip_list"]/tr[@class]')
    for tr in trs:
        base_ip = tr.xpath('td[2]/text()')[0]
        port = tr.xpath('td[3]/text()')[0]
        ip = {
            'ip':base_ip + ':' + port
        }
        save_to_mongo(ip)

def save_to_mongo(ip):
    if db['xicidaili2'].insert(ip):
        print('存储成功',ip)
        return True
    else:
        print('存储失败')
        return False

if __name__ == '__main__':
    main()