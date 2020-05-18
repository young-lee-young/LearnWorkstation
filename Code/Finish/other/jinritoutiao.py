
#多进程爬取今日头条图片，存储到本地

import requests
from urllib.parse import urlencode
from requests.exceptions import RequestException
import json
from bs4 import BeautifulSoup
import re
import os
from hashlib import md5
import time
import random
from multiprocessing import Pool

def get_page_index(offset,keyword):
    data = {
        'offset':offset,
        'format':'json',
        'keyword':keyword,
        'autoload':'true',
        'count':'20',
        'cur_tab':3,
        'from':'gallery'
    }
    url = 'https://www.toutiao.com/search_content/?' + urlencode(data)
    try:
        response = requests.get(url)
        print('请求索引页状态码为:',response.status_code,url)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        print('请求索引页出错',url)
        return None

def parse_page_index(html):
    try:
        data = json.loads(html)
        #data.keys()得到json键名
        if data and 'data' in data.keys():
            for item in data.get('data'):
                yield item.get('article_url')
    except:
        print('解析索引页出错')

def get_page_detail(url):
    try:
        response = requests.get(url)
        print('请求详情页状态码为:',response.status_code,url)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        print('请求详情页出错',url)
        return None

def parse_page_detail(html):
    images_info = []
    try:
        soup = BeautifulSoup(html,'html.parser')
        title = soup.find('title').get_text()
        images_info.append(title)
        pattern = re.compile(r'gallery: JSON.parse\(\"(.*?)\"\),',re.S)
        result = re.search(pattern,html).group(1)
        jsonstr = result.replace("\\","")
        if jsonstr:
            data = json.loads(jsonstr)
            if data and 'sub_images' in data.keys():
                sub_images = data.get('sub_images')
                image_urls = [item.get('url') for item in sub_images]
                for image_url in image_urls:
                    images_info.append(image_url)
                return images_info
    except Exception:
        print('解析详情页出错')

def get_images_html(image_info):
    try :
        for i in range(1,len(image_info)):
            url = image_info[i]
            response = requests.get(url)
            # time.sleep(random.random()*2)
            if response.status_code == 200:
                yield [response.content,url]
            else:
                return None
    except RequestException:
        print('请求图片出错')

def save_image(title,content):
    try:
        base_file_path = '{0}/{1}'.format('D:/toutiao',title)
        if not os.path.exists(base_file_path):
            os.mkdir(base_file_path)
        file_path = '{0}/{1}.jpg'.format(base_file_path,md5(content[0]).hexdigest())
        if not os.path.exists(file_path):
            with open(file_path,'wb') as f:
                f.write(content[0])
                print('存储图片成功',content[1],time.ctime())
    except:
        print('存储失败',content[1])

def main(i):
    html = get_page_index(i*20,'街拍')
    for url in parse_page_index(html):
        html2 = get_page_detail(url)
        if html2:
            image_info = parse_page_detail(html2)
            title = image_info[0]
            for content in get_images_html(image_info):
                save_image(title,content)

if __name__ == '__main__':
    print('主进程ID为',os.getpid())
    pool = Pool(processes=4)
    for i in range(0,100):
        pool.apply_async(main,(i,))
    pool.close()
    pool.join()
