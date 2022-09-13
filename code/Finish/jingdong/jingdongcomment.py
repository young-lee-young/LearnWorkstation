# 本代码用requests进行网页请求，返回的是json,直接解析数据

import requests
from requests.exceptions import RequestException
from urllib.parse import urlencode
import json
import time
import random
import pymongo

client = pymongo.MongoClient('localhost')
db = client['jingdong']
headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}

def get_html(i):
    # global proxies
    # proxy = get_proxy()
    # if proxy:
    #     proxies = {
    #         'https':'http://' + proxy[:-2]
    #     }
    data = {
        'productId': productid,
        'score': '0',
        'sortType': '5',
        'page': i,
        'pageSize': 10,
    }
    url = 'https://sclub.jd.com/comment/productPageComments.action?' + urlencode(data)
    try:
        # response = requests.get(url,headers=headers,proxies=proxies,allow_redirects=False)
        response = requests.get(url,headers=headers)
        print(response.status_code)
        time.sleep(random.random()*3)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            return get_html(i)
        else:
            return None
    except RequestException:
        print('获取页面出错',i)
        return None

def get_proxy():
    url = 'http://127.0.0.1:5000/get'
    try:
        response = requests.get(url)
        time.sleep(random.random() * 5)
        if response.status_code == 200:
            return response.text
        return None
    except ConnectionError:
        return None

def judge_end(html):
    try:
        dict = json.loads(html)
        comments = dict['comments']
        if len(comments) == 0:
            return True
        else:
            return False
    except:
        print('判断是否结束出错')

def get_product(html):
    for data in parse(html):
        save_to_mongo(data)

def parse(html):
    try:
        dict = json.loads(html)
        infos = dict['comments']
        for info in infos:
            try:
                name = info['nickname']
                comment = info['content']
                c_time = info['creationTime']
                color = info['productColor']
                yield {
                    'name':name,
                    'comment':comment,
                    'time':c_time,
                    'color':color
                }
            except:
                print('解析商品出错')
    except:
        print('解析页面出错')

def save_to_mongo(data):
    global productid
    try:
        db[productid].insert(data)
        print('存储成功',data['name'])
    except:
        print('存储失败',data['name'])

def main():
    i = 0
    is_end = False
    html = get_html(i)
    while not is_end:
        get_product(html)
        i += 1
        html = get_html(i)
        is_end = judge_end(html)

if __name__ == '__main__':
    global productid
    productid = '3133847'
    main()