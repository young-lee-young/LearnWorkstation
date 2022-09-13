import requests
from bs4 import BeautifulSoup
import multiprocessing
from multiprocessing import cpu_count
from multiprocessing import Pool
import time
import pymongo
import re

client = pymongo.MongoClient('localhost')
db = client['anjvke']

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.108 Safari/537.36'
}


def get_proxy():
    url = 'http://api.xdaili.cn/xdaili-api//privateProxy/applyStaticProxy?spiderId=f8a45e416b4f47eca9b6aab050703e7e&returnType=1&count=1'
    try:
        response = requests.get(url)
        time.sleep(5)
        if response.status_code == 200:
            print(response.text)
            pattern = re.compile(r'(.*?)\r')
            result = re.findall(pattern, response.text)
            return result[0]
        else:
            return None
    except ConnectionError:
        return None


def get_html(i):
    global proxies
    url = 'https://bj.fang.anjuke.com/loupan/all/p{}/'.format(i)
    print(url)
    try:
        response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
        # response = requests.get(url,headers=headers)
        print(response.status_code)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            proxies = {
                'https': 'http://' + proxy[:-2]
            }
            print(proxies)
            return get_html(i)
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        print('获取页面出错', i)


def parse_html(i, html):
    try:
        soup = BeautifulSoup(html, 'html.parser')
        div_list = soup.find('div', attrs={'class': 'key-list'}).find_all('div', attrs={'class': 'item-mod'})
        for div in div_list:
            try:
                huxing_str = ''
                href = div.find('a', attrs={'class': 'pic'}).get('href')
                pic = div.find('img').get('src')
                name = div.find('span', attrs={'class': 'items-name'}).get_text().strip()
                comments = div.find('span', attrs={'class': 'list-dp'}).get_text().strip()
                address = div.find('span', attrs={'class': 'list-map'}).get_text().strip()
                huxing_list = div.find('a', attrs={'class': 'huxing'}).find_all('span')
                for huxing in huxing_list:
                    # huxing_str = ','.join(huxing.get_text().strip())
                    huxing_str += huxing.get_text().strip() + ','
                other = div.find('div', attrs={'class': 'tag-panel'}).get_text().strip()
                try:
                    price = div.find('p', attrs={'class': 'price'}).get_text().strip()
                    zhoubian = None
                except:
                    price = div.find('p', attrs={'class': 'price-txt'}).get_text().strip()
                    try:
                        zhoubian = div.find('p', attrs={'class': 'favor-tag around-price'}).span.get_text().strip()
                    except:
                        zhoubian = None
                tel = div.find('p', attrs={'class', 'tel'}).get_text().strip()
                yield {
                    'url': href,
                    'pic': pic,
                    'name': name,
                    'comments': comments,
                    'address': address,
                    'huxing': huxing_str,
                    'other': other,
                    'price': price,
                    'zhoubian': zhoubian,
                    'tel': tel
                }
            except:
                print('解析详情出错')
    except:
        print('解析页面出错', i)


def save_to_mongo(data):
    if db['anjvke2'].insert(data):
        print('存储成功', data['name'])
    else:
        print('存储失败')


def main(i):
    html = get_html(i)
    for data in parse_html(i, html):
        save_to_mongo(data)


if __name__ == '__main__':
    global proxies
    proxy = get_proxy()
    proxies = {
        'https': 'http://' + proxy
    }
    print(proxies)
    # pool = Pool(processes=cpu_count())
    # for i in range(0,10):
    #     pool.apply_async(main,(i,))
    # pool.close()
    # pool.join()
    for i in range(1, 10000):
        main(i)
