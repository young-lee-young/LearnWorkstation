import requests
import time
import datetime
from lxml import etree
from bs4 import BeautifulSoup
import re
import random
import pymysql.cursors
import pymongo

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}

# connection = pymysql.connect(
#     host='localhost',
#     port=3306,
#     user='MySQL用户名',
#     password='密码',
#     use_unicode=True,
#     charset='utf8',
#     database='doubanmusic'
# )

client = pymongo.MongoClient('localhost')
db = client.douban
db.authenticate('MongoDB数据库用户名', '密码')
session = requests.Session()

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
        session.post(loginurl, data=my_logindata, headers=headers)

def create_database():
    cursor = connection.cursor()
    cursor.execute('drop table if exists album2')
    creat_table_sql = 'create table album2(' \
                      'id int not null auto_increment primary key,' \
                      'albumname varchar(100),' \
                      'albumhref varchar(100),' \
                      'albuminfo varchar(200),' \
                      'albumstar varchar(100),' \
                      'ratingnums float,' \
                      'commentsnums int(10))'
    cursor.execute(creat_table_sql)
    cursor.close()

def get_proxy():
    global start_time
    start_time = datetime.datetime.now()
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

def get_index():
    global proxies
    url = 'https://music.douban.com/tag/'
    try:
        response = requests.get(url,headers=headers,proxies=proxies,allow_redirects=False,timeout=10)
        # response = session.get(url,headers=headers,allow_redirects=False)
        # time.sleep(random.random()*3)
        if response.status_code == 200:
            return response.text
        if response.status_code ==302:
            proxy = get_proxy()
            if proxy:
                proxies = {
                    'https': 'http://' + proxy
                    # 'https': 'http://' + proxy[:-1]
                    # 'https': 'http://' + proxy[:-2]
                }
                print(proxies)
                html = get_index()
                return html
    except Exception as e:
        print(e.args)


def parse_index(html):
    name_href_list = []
    try:
        soup = BeautifulSoup(html,'html.parser')
        div = soup.find('div',attrs={'id':'wrapper'})
        tds = div.find_all('td')
        for td in tds:
            name = td.find('a').get_text().strip()
            href = td.find('a').get('href')
            name_href_list.append((name,href))
        return name_href_list
    except Exception as e:
        print(e.args)


def get_detail(i,href):
    global proxies
    base_url = 'https://music.douban.com' + href
    url = base_url + '?start={}&type=T'.format(i*20)
    print('正在解析',url)
    try:
        response = requests.get(url,headers=headers,proxies=proxies,allow_redirects=False,timeout=10)
        # response = session.get(url,headers=headers,allow_redirects=False)
        # time.sleep(random.random()*5)
        print(response.status_code)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            if proxy:
                proxies = {
                    'https': 'http://' + proxy
                    # 'https': 'http://' + proxy[:-1]
                    # 'https': 'http://' + proxy[:-2]
                }
                print(proxies)
                html = get_detail(i,href)
                return html
    except Exception as e:
        print(e.args)
        return None

def get_is_end(html):
    try:
        soup = BeautifulSoup(html,'html.parser')
        subject_list = soup.find('div',{'id':'subject_list'})
        div_list = subject_list.find('div')

        if div_list is None:
            return True
        else:
            return False
    except:
        print('判断是否结束出错')

def parse_detail(html):
    try:
        selector = etree.HTML(html)
        tables = selector.xpath('//*[@id="subject_list"]/table')
        for table in tables:
            main_album_name = table.xpath('tr/td[@valign="top"]/div/a/text()')[0].strip()
            sub_album_name = table.xpath('tr/td[@valign="top"]/div/a/span/text()')
            if len(sub_album_name) == 0:
                album_name = main_album_name
            else:
                album_name = main_album_name + sub_album_name[0].strip()
            album_href = table.xpath('tr/td[@valign="top"]/div/a/@href')[0].strip()
            album_info = table.xpath('tr/td[@valign="top"]/div/p/text()')[0].strip()
            try:
                star = table.xpath('tr/td[@valign="top"]/div/div/span[1]/@class')[0].strip()
                rating_nums = table.xpath('tr/td[@valign="top"]/div/div/span[2]/text()')[0].strip()
                nums = table.xpath('tr/td[@valign="top"]/div/div/span[3]/text()')[0].strip()
                pattern = re.compile(r'(\d.*?)人评价')
                comments_nums = re.findall(pattern,nums)[0]
            except:
                star = None
                rating_nums = None
                comments_nums = None
            yield [album_name,album_href,album_info,star,rating_nums,comments_nums]
    except Exception as e:
        print(e.args)

def save_to_sql(info):
    album_name = info[0]
    album_href = info[1]
    album_info = info[2]
    star = info[3]
    rating_nums = info[4]
    comments_nums = info[5]
    try:
        cursor = connection.cursor()
        cursor.execute(
            'insert into album2(albumname,albumhref,albuminfo,albumstar,ratingnums,commentsnums) values (%s,%s,%s,%s,%s,%s)',
            [album_name, album_href, album_info, star, rating_nums, comments_nums])
        connection.commit()
        cursor.close()
        print('存储成功', album_name)
    except:
        print('存储失败')

def save_to_mongo(tag,info):
    try:
        data = {
            'tag':tag,
            'name':info[0],
            'href':info[1],
            'info':info[2],
            'star':info[3],
            'ratingnums':info[4],
            'commentsnums':info[5]
        }
        if db['doubanmusic3'].insert(data):
            print('插入成功',info[0])
    except:
        print('插入失败')

def main():
    login()
    # create_database()
    global start_time
    global end_time
    global proxies
    proxy = get_proxy()
    if proxy:
        proxies = {
            'https': 'http://' + proxy
            # 'https': 'http://' + proxy[:-1]
            # 'https': 'http://' + proxy[:-2]
        }
    print(proxies)
    try:
        index_html = get_index()
        for tag_name_href in parse_index(index_html)[117:]:
            is_end = False
            i = 0
            detail_html = get_detail(i, tag_name_href[1])
            while not is_end:
                for info in parse_detail(detail_html):
                    # save_to_sql(info)
                    save_to_mongo(tag_name_href[0], info)
                i += 1
                detail_html = get_detail(i, tag_name_href[1])
                end_time = datetime.datetime.now()
                time_diff = (end_time - start_time).seconds
                print(time_diff)
                if time_diff > 900:
                    proxy = get_proxy()
                    if proxy:
                        proxies = {
                            'https': 'http://' + proxy
                            # 'https': 'http://' + proxy[:-1]
                            # 'https': 'http://' + proxy[:-2]
                        }
                        print(proxies)
                is_end = get_is_end(detail_html)
    except Exception as e:
        print(e.args)
    finally:
        # connection.close()
        print('爬取完成')

if __name__ == '__main__':
    main()