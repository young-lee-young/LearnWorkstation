import requests
import re
import json
from bs4 import BeautifulSoup
import time
import random
import pymongo

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}

session = requests.Session()
client = pymongo.MongoClient('localhost')
db = client['douban']


def login():
    loginurl = 'https://accounts.douban.com/login'
    my_logindata = {
        'source': 'index_nav',
        'redir': 'https://www.douban.com',
        'form_email': '豆瓣用户名',
        'form_password': '密码',
        'login': u'登录'
    }
    response = session.post(loginurl, data=my_logindata, headers=headers)
    soup = BeautifulSoup(response.text, 'html.parser')
    captcha = soup.find('img', id='captcha_image')
    if captcha:
        captcha_url = captcha['src']
        pattern = r'<input type="hidden" name="captcha-id" value="(.*?)"/'
        captcha_id = re.findall(pattern, response.text)
        print(captcha_url)
        captcha_text = input('please input:')
        my_logindata['captcha-solution'] = captcha_text
        my_logindata['captcha-id'] = captcha_id
        print(my_logindata)
        session.post(loginurl, data=my_logindata, headers=headers)


def get_proxy():
    url = '代理的接口'
    try:
        response = requests.get(url)
        time.sleep(5)
        if response.status_code == 200:
            print(response.text)
            return response.text
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        return None


def get_html(i,page):
    global proxies
    url = 'https://movie.douban.com/j/new_search_subjects?sort=T&range=0,{}&tags=&start={}'.format(i,page*20)
    try:
        response = session.get(url, headers=headers, proxies=proxies, allow_redirects=False, timeout=10)
        response = session.get(url,headers=headers,allow_redirects=False)
        print(response.status_code,i,page*20)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            proxies = {
                'https':'http://' + proxy[:-2]
            }
            print(proxies)
            return get_html(i)
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        print('获取页面失败')
        return None


def parse_html(html):
    try:
        json_str = json.loads(html)
        if 'data' in json_str:
            return json_str['data']
    except:
        print('解析失败')


def save_to_mongo(data):
    try:
        if db['doubanmovie2'].insert(data):
            print('插入成功')
    except:
        print('插入失败')


def main():
    login()
    global proxies
    proxy = get_proxy()
    proxies = {
        'https': 'http://' + proxy[:-2]
    }
    print(proxies)
    try:
        for i in range(2, 11):
            page = 0
            html = get_html(i, page)
            data = parse_html(html)
            data_len = len(data)
            while not data_len == 0:
                save_to_mongo(data)
                page += 1
                html = get_html(i, page)
                data = parse_html(html)
                data_len = len(data)
    except Exception as e:
        print(e.args)
        print('爬取出错')
    finally:
        print('爬取完成')


if __name__ == '__main__':
    main()