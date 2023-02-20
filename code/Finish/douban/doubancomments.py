from multiprocessing import Pool
import requests
from lxml import etree
from bs4 import BeautifulSoup
import re
import json
import time
import random
import pymongo
from xpinyin import Pinyin
headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}
session = requests.Session()


def login():
    loginurl = 'https://accounts.douban.com/login'
    logindata = {
        'source': 'index_nav',
        'redir': 'https://www.douban.com',
        'form_email': '18400000000',
        'form_password': 'helloworld',
        'login': u'登录'
    }
    response = session.post(loginurl, data=logindata, headers=headers)
    soup = BeautifulSoup(response.text, 'html.parser')
    captcha = soup.find('img', id='captcha_image')
    if captcha:
        captcha_url = captcha['src']
        pattern = r'<input type="hidden" name="captcha-id" value="(.*?)"/'
        captcha_id = re.findall(pattern, response.text)
        print(captcha_url)
        captcha_text = input('please input:')
        logindata['captcha-solution'] = captcha_text
        logindata['captcha-id'] = captcha_id
        session.post(loginurl, data=logindata, headers=headers)


def get_proxy():
    url = 'http://127.0.0.1:5000/get'
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        return None


def get_name_pages(movie_id):
    global proxies
    url = 'https://movie.douban.com/subject/{}/reviews'.format(movie_id)
    try:
        # response = session.get(url,headers=headers,proxies=proxies,allow_redirects=False)
        response = session.get(url, headers=headers)
        print(response.status_code, url)
        if response.status_code == 200:
            return parse_name_pages(response.text)
        if response.status_code == 302:
            proxy = get_proxy()
            if proxy:
                proxies = {
                    'https': 'http://' + proxy[:-2]
                }
                print(proxies)
                info = get_name_pages(movie_id)
                return info
    except ConnectionError as e:
        print('获取最大页出错')
        print(e.args)


def parse_name_pages(html):
    try:
        soup = BeautifulSoup(html, 'html.parser')
        result = soup.find('h1').get_text()
        pattern = re.compile(r'(.*?)的影评', re.S)
        name = re.search(pattern, result).group(1).strip()
        page = soup.find('span', attrs={'class': 'thispage'}).get('data-total-page')
        return (name, int(page))
    except:
        print('解析最大页出错')


def get_html(movie_id, i):
    global proxies
    url = 'https://movie.douban.com/subject/{}/reviews?start={}'.format(movie_id, i*20)
    try:
        # response = session.get(url,headers = headers,proxies = proxies,allow_redirects=False)
        response = session.get(url, headers=headers)
        time.sleep(random.random()*5)
        print(response.status_code, url)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            if proxy:
                proxies = {
                    'https': 'http://' + proxy[:-2]
                }
                print(proxies)
                html = get_html(movie_id, i)
                return html
    except ConnectionError as e:
        print('获取评论页出错', url)
        print(e.args)


def parse_html(html):
    try:
        soup = BeautifulSoup(html, 'html.parser')
        div_list = soup.find_all('div', attrs={'typeof': 'v:Review'})
        for div in div_list:
            try:
                image = div.find('a', attrs={'class': 'avator'}).img.get('src')
                name = div.find('a', attrs={'class': 'name'}).get_text().strip()
                try:
                    star = div.find('span', attrs={'property': 'v:rating'}).get('class')[0]
                except:
                    star = None
                time_c = div.find('span', attrs={'class': 'main-meta'}).get_text().strip()
                title = div.find('h2').a.get_text().strip()
                href = div.find('h2').a.get('href')
                up = div.find('a', attrs={'class': 'action-btn up'}).span.get_text().strip()
                down = div.find('a', attrs={'class': 'action-btn down'}).span.get_text().strip()
                reply = div.find('a', attrs={'class': 'reply'}).get_text().strip()
                yield {
                    'image': image,
                    'name': name,
                    'star': star,
                    'time': time_c,
                    'title': title,
                    'href': href,
                    'up': up,
                    'down': down,
                    'reply': reply
                }
            except Exception as e:
                print(e.args)
                print('解析详情出错')
    except Exception as e:
        print('解析页面出错')
        print(e.args)


def save_to_mongo(name, data):
    pinyin = Pinyin()
    title = pinyin.get_pinyin(name, '')
    try:
        client = pymongo.MongoClient('localhost')
        db = client['doubanmoviecomments']
        if db[title].insert(data):
            print('存储成功', data['name'])
    except:
        print('存储失败', data['name'])


def main(movie_id):
    # login()
    # global proxies
    # proxy = get_proxy()
    # if proxy:
    #     proxies = {
    #         'https:': 'http://' + proxy[:-2]
    #     }
    # print(proxies)
    name_pages = get_name_pages(movie_id)
    for i in range(0, name_pages[1] + 1):
        html = get_html(movie_id, i)
        for data in parse_html(html):
            save_to_mongo(name_pages[0], data)


if __name__ == '__main__':
    movie_id = '1292052'
    main(movie_id)