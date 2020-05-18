import requests
from lxml import etree
from bs4 import BeautifulSoup
import time
import datetime
import random
import re
import pymysql.cursors
import pymongo

class Doubanread(object):

    headers = {
        'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
    }

    # connection = pymysql.connect(
    #     host='localhost',
    #     port=3306,
    #     user='MySQL用户名',
    #     password='密码',
    #     use_unicode=True,
    #     charset='utf8',
    #     database='doubanread'
    # )

    client = pymongo.MongoClient('localhost')
    db = client['douban']
    session = requests.Session()

    def login(self):
        loginurl = 'https://accounts.douban.com/login'
        my_logindata = {
            'source': 'None',
            'redir': 'https://www.douban.com',
            'form_email': '豆瓣用户名',
            'form_password': '密码',
            'login': u'登录'
        }
        response = self.session.post(loginurl, data=my_logindata, headers=self.headers)
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
            self.session.post(loginurl, data=my_logindata, headers=self.headers)

    def create_database(self):
        cursor = self.connection.cursor()
        cursor.execute('drop table if exists bookinfo3')
        creat_table_sql = 'create table bookinfo3(' \
                          'id int not null auto_increment primary key,' \
                          'bookname varchar(100),' \
                          'bookhref varchar(100),' \
                          'bookinfos varchar(200),' \
                          'bookstar varchar(100),' \
                          'ratingnums float,' \
                          'commentsnums varchar(20),' \
                          'xiangqing text)'
        cursor.execute(creat_table_sql)
        cursor.close()

    def get_proxy(self):
        global start_time
        start_time = datetime.datetime.now()
        url = '代理的接口'
        try:
            response = requests.get(url)
            time.sleep(5)
            if response.status_code == 200:
                print(response.text)
                return response.text
            return None
        except ConnectionError:
            return None

    def get_index_html(self):
        global proxies
        url = 'https://book.douban.com/tag/'
        try:
            response = requests.get(url, headers=self.headers, proxies=proxies, allow_redirects=False,timeout=10)
            # response = self.session.get(url,headers=self.headers,allow_redirects=False)
            # time.sleep(random.random()*3)
            print(url,response.status_code)
            if response.status_code == 200:
                return response.text
            if response.status_code == 302:
                proxy = self.get_proxy()
                if proxy:
                    proxies = {
                        'https': 'http://' + proxy[:-2]
                    }
                    print(proxies)
                    html = self.get_index_html()
                    return html
        except Exception as e:
            print(e.args)

    def parse_index(self,html):
        name_href_list = []
        try:
            soup = BeautifulSoup(html,'html.parser')
            tds = soup.find('div',attrs = {'class':'article'}).find_all('td')
            for td in tds:
                name = td.find('a').get_text().strip()
                href = td.find('a').get('href')
                # yield (name,href)
                name_href_list.append((name,href))
            return name_href_list
        except Exception as e:
            print(e.args)

    def get_pages(self,tag_href):
        global proxies
        url = 'https://book.douban.com' + tag_href
        try:
            response = requests.get(url, headers=self.headers, proxies=proxies, allow_redirects=False,timeout=10)
            # response = self.session.get(url,headers=self.headers,allow_redirects=False)
            # time.sleep(random.random()*3)
            print(url,response.status_code)
            if response.status_code == 200:
                page_source = etree.HTML(response.text)
                max_page = int(page_source.xpath('//*[@id="subject_list"]/div[@class="paginator"]/a[last()]/text()')[0])
                if max_page > 50:
                    page = 50
                else:
                    page = max_page
                return page
            if response == 302:
                proxy = self.get_proxy()
                if proxy:
                    proxies = {
                        'https': 'http://' + proxy[:-2]
                    }
                    print(proxies)
                    page = self.get_pages(tag_href)
                    return page
        except Exception as e:
            print(e.args)
            return 1

    def get_detail_html(self,i,href):
        global proxies
        base_url = 'https://book.douban.com' + href
        url = base_url + '?start={}&type=T'.format(i * 20)
        try:
            response = requests.get(url, headers=self.headers, proxies=proxies, allow_redirects=False,timeout=10)
            # response = self.session.get(url,headers=self.headers,allow_redirects=False)
            # time.sleep(random.random()*5)
            print(url,response.status_code)
            if response.status_code == 200:
                return response.text
            if response.status_code == 302:
                proxy = self.get_proxy()
                if proxy:
                    proxies = {
                        'https': 'http://' + proxy[:-2]
                    }
                    print(proxies)
                    html = self.get_detail_html(i, href)
                    return html
        except Exception as e:
            print(e.args)
            return None

    def parse_detail(self,html):
        try:
            page_source = etree.HTML(html)
            li_lists = page_source.xpath('//*[@id="subject_list"]/ul/li')
            for li_list in li_lists:
                main_book_name = li_list.xpath('div[@class="info"]/h2/a/text()')[0].strip()
                sub_book_name = li_list.xpath('div[@class="info"]/h2/a/span/text()')
                if len(sub_book_name) == 0:
                    book_name = main_book_name
                else:
                    book_name = main_book_name + sub_book_name[0].strip()
                href = li_list.xpath('div[@class="info"]/h2/a/@href')[0].strip()
                infos = li_list.xpath('div[@class="info"]/div/text()')[0].strip()
                try:
                    star = li_list.xpath('div[@class="info"]/div[@class="star clearfix"]/span[1]/@class')[0].strip()
                    rating_nums = li_list.xpath('div[@class="info"]/div[@class="star clearfix"]/span[2]/text()')[0].strip()
                    comments_nums = li_list.xpath('div[@class="info"]/div[@class="star clearfix"]/span[3]/text()')[0].strip()
                except:
                    star = None
                    rating_nums = None
                    comments_nums = None
                try:
                    xiangqing = li_list.xpath('div[@class="info"]/p/text()')[0].strip()
                except:
                    xiangqing = None
                yield [book_name,href,infos,star,rating_nums,comments_nums,xiangqing]
        except:
            print('解析错误',book_name)

    def save_to_sql(self,book_info):
        book_name = book_info[0]
        href = book_info[1]
        infos = book_info[2]
        star = book_info[3]
        rating_nums = book_info[4]
        comments_nums = book_info[5]
        xiangqing = book_info[6]
        try:
            cursor = self.connection.cursor()
            cursor.execute('insert into bookinfo3(bookname,bookhref,bookinfos,bookstar,ratingnums,commentsnums,xiangqing) values (%s,%s,%s,%s,%s,%s,%s)',[book_name,href,infos,star,rating_nums,comments_nums,xiangqing])
            self.connection.commit()
            cursor.close()
            print('存储成功',book_name)
        except:
            print('存储失败')

    def save_to_mongo(self,tag,bookinfo):
        data = {
            'tag':tag,
            'bookname':bookinfo[0],
            'href':bookinfo[1],
            'infos':bookinfo[2],
            'star':bookinfo[3],
            'ratingnums':bookinfo[4],
            'commentsnums':bookinfo[5],
            'xiangqing':bookinfo[6]
        }
        try:
            if self.db['doubanread4'].insert(data):
                print('插入成功',bookinfo[0])
        except:
            print('插入失败')

    def main(self):
        self.login()
        # self.create_database()
        global start_time
        global end_time
        global proxies
        proxy = self.get_proxy()
        if proxy:
            proxies = {
                'https':'http://' + proxy[:-2]
            }
        print(proxies)
        try:
            index_html = self.get_index_html()
            for tag_name_href in self.parse_index(index_html):
                page = self.get_pages(tag_name_href[1])
                for i in range(0,page):
                    detail_html = self.get_detail_html(i,tag_name_href[1])
                    end_time = datetime.datetime.now()
                    time_diff = (end_time - start_time).seconds
                    print(time_diff)
                    if time_diff > 60:
                        proxy = self.get_proxy()
                        if proxy:
                            proxies = {
                                'https': 'http://' + proxy[:-2]
                            }
                            print(proxies)
                    for book_info in self.parse_detail(detail_html):
                        # self.save_to_sql(book_info)
                        self.save_to_mongo(tag_name_href[0],book_info)
        except Exception as e:
            print(e.args)
        finally:
            # self.connection.close()
            print('爬取完成')

if __name__ == '__main__':
    doubanread = Doubanread()
    doubanread.main()