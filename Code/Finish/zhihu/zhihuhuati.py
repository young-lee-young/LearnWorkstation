import requests
import json
from requests.exceptions import RequestException
from lxml import etree
import random
import time
import re
from bs4 import BeautifulSoup
import chardet
from multiprocessing import Process
from multiprocessing import Pool
import os

headers = {
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
    'X-Requested-With':'XMLHttpRequest',
    'X-Xsrftoken':'9a79427f-9a1e-43c3-a40d-a87d733afde8'
}

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

def get_index_html():
    global proxies
    url = 'https://www.zhihu.com/topics'
    try:
        # response = requests.get(url,headers=headers,proxies=proxies,allow_redirects=False)
        response = requests.get(url,headers=headers)
        print(response.status_code)
        time.sleep(random.random()*3)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            proxies = {
                'https': 'http://' + proxy[:-2]
            }
            print(proxies)
            return get_index_html()
        else:
            return None
    except RequestException:
        print('获取话题索引页出错',url)

def parse_index(html):
    try:
        selector = etree.HTML(html)
        topic_list = selector.xpath('/html/body/div[3]/div[1]/div/div/ul/li')
        for topic in topic_list:
            topic_name = topic.xpath('a/text()')[0]
            topic_id = topic.xpath('@data-id')[0]
            yield [topic_name,topic_id]
    except:
        print('解析话题索引页出错')

def get_topic_html(topic_info,off_set):
    topic_id = int(topic_info[1])
    post_data = {
        'method':'next',
        'params':'{"topic_id":253,"offset":0,"hash_id":"1bb0946678fa20cbe611af63a62cc798"}'
    }
    id = eval(post_data['params'])
    id['topic_id'] = topic_id
    id['offset'] = off_set*20
    id = str(id).replace('\'','\"')
    post_data['params'] = id
    url = 'https://www.zhihu.com/node/TopicsPlazzaListV2'
    global proxies
    try:
        # response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
        response = requests.post(url,data=post_data,headers=headers)
        time.sleep(random.random()*10)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            proxies = {
                'https':'http://' + proxy[:-2]
            }
            print(proxies)
            return get_topic_html(topic_info,off_set)
        else:
            return None
    except RequestException:
        print('获取二级话题索引出错',topic_info[0])

def is_end(html):
    try:
        json_str = json.loads(html, encoding='unicode')
        msg_list = json_str['msg']
        if len(msg_list) == 0:
            return True
        else:
            return False
    except:
        print('判断话题是否结束出错')

def parse_topic_name_href(html):
    try:
        json_str = json.loads(html, encoding='unicode')
        pattern = re.compile('href="/topic/(.*?)">.*?alt="(.*?)"')
        topic_href_names = re.findall(pattern, str(json_str))
        for topic_href_name in topic_href_names:
            yield topic_href_name
    except:
        print('解析二级话题索引失败')

def num_1(topic_id,topic_name):
    print('获取动态问题线程ID为:',os.getpid())
    hot_first_html = get_hot_first_html(topic_id)
    for question_id in parse_question_id(hot_first_html):
        print('%s二级话题动态问题ID为:%s' % (topic_name,question_id))
    data_score = parse_data_score(hot_first_html,i = 1)
    while data_score:
        hot_other_html = get_hot_other_html(topic_id, data_score)
        for question_id in parse_question_id(hot_other_html):
            print('%s二级话题动态问题ID为:%s' % (topic_name,question_id))
        data_score = parse_data_score(hot_other_html)
    print('%s二级话题动态问题ID搜索完成' % (topic_name))

def get_hot_first_html(topic_id):
    global proxies
    url = 'https://www.zhihu.com/topic/{}/hot'.format(topic_id)
    try:
        # response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
        response = requests.get(url,headers=headers)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            return response.text
        if response.status_code == 302:
            proxy = get_proxy()
            proxies = {
                'https':'http://' + proxy[:-2]
            }
            print(proxies)
            return get_hot_first_html(topic_id)
        else:
            return None
    except RequestException:
        print('获取动态第一页出错',url)

def get_hot_other_html(topic_id,data_score):
    global proxies
    url = 'https://www.zhihu.com/topic/{}/hot'.format(topic_id)
    post_data = {
        'start': '0',
        'offset': '1'
    }
    post_data['offset'] = data_score
    try:
        # response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
        response = requests.post(url,data=post_data,headers=headers)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            json_str = json.loads(response.text)
            return json_str
        if response.status_code ==302:
            proxy = get_proxy()
            proxies = {
                'https':'http://' + proxy[:-2]
            }
            print(proxies)
            return get_hot_other_html(topic_id,data_score)
        else:
            return None
    except RequestException:
        print('获取动态其他页出错',url)

def parse_data_score(html,i = 0):
    if i == 0:
        if html.get('msg')[0] == 0:
            return False
    try:
        pattern = re.compile(r'data-score="(.*?)"',re.S)
        result = re.findall(pattern,str(html))[-1]
        return result
    except:
        print('解析动态页出错')

def num_2(topic_id,topic_name):
    print('获取精华问题线程ID为:',os.getpid())
    max_page = get_max_page(topic_id, 2)
    for html in get_question_list_html(topic_id, max_page, 2):
        for question_id in parse_question_id(html):
            print('%s二级话题精华问题ID为:%s' % (topic_name,question_id))
    print('%s二级话题精华问题ID搜索完成' % (topic_name))

def num_3(topic_id,topic_name):
    print('获取未回答问题线程ID为:',os.getpid())
    max_page = get_max_page(topic_id, 3)
    for html in get_question_list_html(topic_id, max_page, 3):
        for question_id in parse_unanswered_html(html):
            print('%s二级话题未回答问题ID为:%s' % (topic_name,question_id))
    print('%s二级话题未回答问题ID搜索完成' % (topic_name))

def get_max_page(topic_id,num):
    global proxies
    if num == 2:
        url = 'https://www.zhihu.com/topic/{}/top-answers'.format(topic_id)
    if num == 3:
        url = 'https://www.zhihu.com/topic/{}/unanswered'.format(topic_id)
    try:
        # response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
        response = requests.get(url,headers=headers)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            page = parse_max_page(response.text)
            return page
        if response.status_code ==302:
            proxy = get_proxy()
            proxies = {
                'https':'http://' + proxy[:-2]
            }
            print(proxies)
            return get_max_page(topic_id,num)
        else:
            return None
    except RequestException:
        print('获取最大页数出错')
        return None

def parse_max_page(html):
    try:
        soup = BeautifulSoup(html,'html.parser')
        page = soup.find('div',attrs={'class':'zm-invite-pager'}).find_all('span')[-2].get_text()
        return page
    except:
        return 1

def get_question_list_html(topic_id,max_page,num):
    global proxies
    page = int(max_page)
    if num == 2:
        base_url = 'https://www.zhihu.com/topic/{}/top-answers?'.format(topic_id)
    if num == 3:
        base_url = 'https://www.zhihu.com/topic/{}/unanswered?'.format(topic_id)
    for i in range(1,page+1):
        url = base_url + 'page={}'.format(i)
        print(url)
        try:
            # response = requests.get(url, headers=headers, proxies=proxies, allow_redirects=False)
            response = requests.get(url,headers=headers)
            time.sleep(random.random()*10)
            if response.status_code == 200:
                yield response.text
            if response.status_code ==302:
                proxy = get_proxy()
                proxies = {
                    'https':'http://' + proxy[:-2]
                }
                print(proxies)
            else:
                return None
        except RequestException:
            print('获取问题列表页出错',url)
            return None

def parse_question_id(html):
    try:
        pattern = re.compile(r'href="/question/(\d*?)"[^\\]target="_blank',re.S)
        question_ids = re.findall(pattern,str(html))
        for question_id in question_ids:
            yield question_id
    except:
        print('解析动态或精华问题ID出错')

def parse_unanswered_html(html):
    try:
        selector = etree.HTML(html)
        question_list = selector.xpath('//*[@id="zh-topic-questions-list"]/div')
        for question in question_list:
            question_id = question.xpath('h2/a/@href')[0]
            pattern = re.compile(r'/question/(\d*)',re.S)
            result = re.findall(pattern,question_id)
            yield result[0]
    except:
        print('解析等待回答ID问题出错')

def main():
    global proxies
    proxy = get_proxy()
    proxies = {
        'https':'http://' + proxy[:-2]
    }
    print(proxies)
    print('主进程ID为:',os.getpid())
    index_html = get_index_html()
    for index_topic_info in parse_index(index_html):
        print('---------------%s一级话题开始解析-----------------' % (index_topic_info[0]))
        off_set = 1
        while off_set:
            topic_html = get_topic_html(index_topic_info,off_set - 1)
            is_ends = is_end(topic_html)
            if not is_ends:
                off_set += 1
                for topic_href_name in parse_topic_name_href(topic_html):
                #     print(topic_href_name)
                # topic_href_names = [('19769568','儿童网游'),('19856249','乐果游戏'),('19826295','Pianoid'),('19734061','游戏前台开发')]
                # for topic_href_name in topic_href_names:

                    topic_id = topic_href_name[0]
                    topic_name = topic_href_name[1]
                    print('--------------%s二级话题开始解析-------------' % (topic_name))
                    process1 = Process(target=num_1,args=(topic_id,topic_name,))
                    process1.start()
                    process2 = Process(target=num_2,args=(topic_id,topic_name,))
                    process2.start()
                    process3 = Process(target=num_3,args=(topic_id,topic_name,))
                    process3.start()
                    process1.join()
                    process2.join()
                    process3.join()
                    print('--------------%s二级话题解析完毕-------------' % (topic_name))
            else:
                off_set = None
                print('---------------%s一级话题解析完毕-----------------' % (index_topic_info[0]))

if __name__ == '__main__':
    main()
    print('爬取完成')