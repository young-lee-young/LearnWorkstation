import requests
from urllib.parse import urlencode
from bs4 import BeautifulSoup
from requests.exceptions import ConnectionError
import time
from multiprocessing import Pool
import pymongo

client = pymongo.MongoClient('localhost')
db = client['weixin']
base_url = 'http://weixin.sogou.com/weixin?'
headers = {
    'Host': 'weixin.sogou.com',
    'Cookie': 'ad=P0osaZllll2BM14QlllllVXEyNYlllllNYkldyllll9llllllOdZb@@@@@@@@@@@; CXID=986832796DDB844BC82E3F7383217E98; SUID=F50088753965860A59FB41D60003ED7A; SUV=00E317847945528459FC5EBABA697567; usid=Gbz1iQMQjPqciMPe; IPLOC=CN1100; ABTEST=6|1511875812|v1; weixinIndexVisited=1; sct=3; ppinf=5|1511931416|1513141016|dHJ1c3Q6MToxfGNsaWVudGlkOjQ6MjAxN3x1bmlxbmFtZTo1NDolRTUlODMlOEYlRTklQTMlOEUlRTQlQjglODAlRTYlQTAlQjclRTglODclQUElRTclOTQlQjF8Y3J0OjEwOjE1MTE5MzE0MTZ8cmVmbmljazo1NDolRTUlODMlOEYlRTklQTMlOEUlRTQlQjglODAlRTYlQTAlQjclRTglODclQUElRTclOTQlQjF8dXNlcmlkOjQ0Om85dDJsdUVXR2dKekxoNk9MZnIxdWRtY2xfcTBAd2VpeGluLnNvaHUuY29tfA; pprdig=LXFwbOUhqN_85h5KL_36EvPA5UPa9Nr_4BeYKc0M-yFfaHRvqZMSaO3YV4zBuZbaWIbXmQkFwfcH8asVyd81u1P_yn0wJr3fNSxVtV2c2gxwrAFxlzgRZu6c9AnlHW0xicwoeFOTTVIJ4fN_G6a4VEt7t0gv_m-9zv1HhuD_E7Y; sgid=31-32178215-AVoePhiarasicEx6CMb976M9Y; SUIR=3DEBFCC1B8BCE6B2BD78C092B998E7EB; SNUID=30E9FEC2BABEE5A27C4A9EFCBBC2F5FE; JSESSIONID=aaanqRAJr8JKq5bs6Kv8v; ppmdig=15119522680000007125508e96fbeeeb05e73970f187968b',
    'User_Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
    'Upgrade-Insecure-Requests': '1'
}
proxy = None
max_count = 5


def get_proxy():
    url = 'http://127.0.0.1:5000/get'
    try:
        response = requests.get(url)
        if response.status_code == 200:
            return response.text
        return None
    except ConnectionError:
        return None


def get_html(url, count=1):
    print('crawing url', url)
    print('trying count', count)
    global proxy
    if count >= max_count:
        return None
    try:
        if proxy:
            proxies = {
                'http': 'http://' + proxy
            }
            response = requests.get(url, allow_redirects=False, headers=headers, proxies=proxies)
            print(response.status_code)
        else:
            response = requests.get(url, allow_redirects=False, headers=headers)
            print(response.status_code)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            if proxy:
                print('Using proxy', proxy)
                return get_html(url)
            else:
                print('get proxy failed')
                return None
            return None
    except ConnectionError as e:
        print('Error Occured', e.args)
        proxy = get_proxy()
        count += 1
        return get_html(url, count)


def parse_html(html):
    soup = BeautifulSoup(html, 'html.parser')
    div = soup.find('ul', attrs={'class': 'news-list'})
    li_list = div.find_all('li')
    for li in li_list:
        href = li.find('div', attrs={'class': 'txt-box'}).h3.a.get('href')
        yield href


def get_detail(url):
    try:
        response = requests.get(url)
        # time.sleep(5)
        if response.status_code == 200:
            return response.text
        return None
    except ConnectionError:
        return None


def parse_detail(html):
    try:
        soup = BeautifulSoup(html, 'html.parser')
        title = soup.find('h2', attrs={'class': 'rich_media_title'}).get_text()
        content = soup.find('div', attrs={'class': 'rich_media_content '}).get_text()
        date = soup.find('em', attrs={'id': 'post-date'}).get_text()
        nickname = soup.find('span', attrs={
            'class': 'rich_media_meta rich_media_meta_text rich_media_meta_nickname'}).get_text()
        return {
            'title': title,
            'content': content,
            'date': date,
            'nickname': nickname
        }
    except AttributeError:
        print('没有找到标签')


def save_to_mongo(data):
    if db['articles'].update({'title': data['title']}, {'$set': data}, True):
        print('save to mongo', data['title'])
    else:
        print('save to mongo failed', data['title'])


def main(keyword, page):
    urldata = {
        'query': keyword,
        'type': 2,
        'page': page
    }
    url = base_url + urlencode(urldata)
    html = get_html(url)
    if html:
        article_urls = parse_html(html)
        for article_url in article_urls:
            article_html = get_detail(article_url)
            article_data = parse_detail(article_html)
            print(article_data)
            if article_data:
                save_to_mongo(article_data)


if __name__ == '__main__':
    pool = Pool(processes=4)
    for i in range(1, 100):
        pool.apply_async(main, ('风景', i))
        # main(u'风景',i)
        # pool.map(main,[i for i in])
    pool.close()
    pool.join()
