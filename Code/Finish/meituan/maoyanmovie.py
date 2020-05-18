#猫眼电影在5000个电影信息之后就一直是重复的，爬取价值不大
#爬取猫眼电影适合用云打码平台，不适合用代理

import requests
from bs4 import BeautifulSoup
import time
import random
import re
from urllib.parse import urlencode
import pymongo

headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}
session = requests.Session()
client = pymongo.MongoClient('localhost')
db = client['meituan']

def login(username,password):
    html_pre_login = pre_login()
    param = parse_param(html_pre_login)
    html_login = formal_login(username,password,param)
    token = parse_token(html_login)
    redirect_login(token)

def pre_login():
    try:
        param = {
            'service':'maoyan',
            'continue':'http://maoyan.com/passport/login?redirect=%2F'
        }
        url = 'https://passport.meituan.com/account/unitivelogin?' + urlencode(param)
        response = session.get(url,headers=headers)
        if response.status_code == 200:
            return response.text
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        print('预登陆出错')

def parse_param(html):
    try:
        pattern1 = re.compile(r'name="csrf" value="(.*?)"',re.S)
        pattern2 = re.compile(r'class="form-uuid" style="display:none">(.*?)</i>',re.S)
        pattern3 = re.compile(r'class="form-field J-form-field-captcha form-field--captcha" style="display:(.*?)"')
        csrf = re.search(pattern1,html).group(1)
        uuid = re.search(pattern2,html).group(1)
        need_captcha = re.search(pattern3,html).group(1)
        return (csrf,uuid,need_captcha)
    except:
        print('解析csrf,uuid,need_captcha出错')

def formal_login(username,password,param):
    csrf = param[0]
    uuid = param[1]
    if param[2] is 'none':
        captcha_param = {
            'uuid':uuid,
        }
        url = 'https://passport.meituan.com/account/captcha?' + urlencode(captcha_param)
        print(url)
        captcha = input('需要验证码:')
    else:
        captcha = ''
    url_param = {
        'uuid':uuid,
        'service':'maoyan',
        'continue': 'http://maoyan.com/passport/login?redirect=%2F',
    }
    postdata ={
        'email': username,
        'password': password,
        'captcha': captcha,
        'origin': 'account - login',
        'fingerprint': '',
        'csrf': csrf
    }
    url = 'https://passport.meituan.com/account/unitivelogin?' + urlencode(url_param)
    try:
        response = session.post(url,data=postdata,headers=headers)
        if response.status_code == 200:
            return response.text
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        print('登录出错')

def parse_token(html):
    try:
        pattern = re.compile(r'name="token" type="hidden" value="(.*?)"',re.S)
        token = re.search(pattern,html).group(1)
        return token
    except:
        print('解析token出错')

def redirect_login(token):
    postdata = {
        'token': token,
        'expire': 0,
        'isdialog': 0,
        'autologin': 0,
        'logintype': 'normal'
    }
    try:
        url = 'http://maoyan.com/passport/login?redirect=%2F'
        session.post(url,data=postdata,headers=headers)
        print('登录成功')
    except ConnectionError as e:
        print(e.args)
        print('重定向出错')

def test():
    try:
        time.sleep(5)
        url = 'http://maoyan.com/profile'
        response = session.get(url,headers=headers)
        print(response.status_code)
        print(response.text)
    except ConnectionError as e:
        print(e.args)
        print('测试是否完成登录出错')

def get_html(i):
    try:
        url = 'http://maoyan.com/films?offset={}'.format(i*30)
        response = session.get(url,headers=headers,allow_redirects=False)
        print(url,response.status_code)
        soup = BeautifulSoup(response.text,'html.parser')
        title = soup.find('head').find('title').get_text()
        if title == '猫眼访问控制':
            print('需要验证码')
            param = get_param(response.text)
            print(param[0])
            captcha_code = input('please input:')
            post_data = {
                'captcha_code': captcha_code,
                'ip': param[1],
                'url': param[2],
                'ticket': param[3],
                'request_code': param[4],
                'contact': ''
            }
            post_url = 'http://maoyan.com/films?__oceanus_captcha=1'
            session.post(post_url,data=post_data)
            return get_html(i)
        else:
            return response.text
    except ConnectionError as e:
        print(e.args)
        print('获取页面出错')

def get_param(html):
    soup = BeautifulSoup(html,'html.parser')
    captcha_url = soup.find('div',{'class':'row'}).find('img').get('src')
    ip = soup.find('input',{'name':'ip'}).get('value')
    url = soup.find('input',{'name':'url'}).get('value')
    ticket = soup.find('input',{'name':'ticket'}).get('value')
    request_code = soup.find('input',{'name':'request_code'}).get('value')
    return (captcha_url,ip,url,ticket,request_code)

def get_max_page():
    try:
        html = get_html(0)
        soup = BeautifulSoup(html,'html.parser')
        max_page = soup.find('ul',attrs={'class':'list-pager'}).find_all('li')[-2].get_text()
        return int(max_page)
    except:
        print('解析最大页出错')

def parse_html(html):
    try:
        soup = BeautifulSoup(html,'html.parser')
        dd_list = soup.find('div',attrs={'class':'movies-list'}).find_all('dd')
        for dd in dd_list:
            try:
                pic = dd.find_all('img')[1].get('data-src')
                name = dd.find('a',{'data-act':'movies-click'}).get_text().strip()
                try:
                    integer = dd.find('i',{'class':'integer'}).get_text().strip()
                    fraction = dd.find('i',{'class':'fraction'}).get_text().strip()
                    score = integer + fraction
                except:
                    score = '暂无评价'
                yield {
                    'pic':pic,
                    'name':name,
                    'score':score
                }
            except:
                print('解析详情错误')
    except:
        print('解析页面错误')

def save_to_mongo(data):
    if db['maoyanmovie1'].insert(data):
        print('插入成功',data['name'])
    else:
        print('插入失败')

def main(username,password):
    login(username,password)
    max_page = get_max_page()
    for i in range(0,max_page):
        html = get_html(i)
        for data in parse_html(html):
            save_to_mongo(data)

if __name__ == '__main__':
    username = '18401681943'
    password = 'lyzy1314'
    main(username,password)