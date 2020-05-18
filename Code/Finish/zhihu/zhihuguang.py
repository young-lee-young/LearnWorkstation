import requests
from requests.exceptions import RequestException
from urllib.parse import urlencode
import json
import re
import time
import random
from multiprocessing import Process
import threading
import os

headers = {
    'authorization':'Bearer Mi4xS1hZRUFnQUFBQUFBZ01LT3JGcmVEQmNBQUFCaEFsVk5jX3NvV3dCT3BIRC0zZmkxbzRLZDZ6djUxWl9LTnRUdXV3|1513860467|e85818df2fec674c16b69e644caba09110c0f553',
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}

lock = threading.Lock()

def get_html(question_id=None,url=None,i=0):
    if i == -1:
        url = 'https://zhuanlan.zhihu.com/p/22804420'
    if i == 1:
        data = {
            'include':'data[*].is_normal,admin_closed_comment,reward_info,is_collapsed,annotation_action,annotation_detail,collapse_reason,is_sticky,collapsed_by,suggest_edit,comment_count,can_comment,content,editable_content,voteup_count,reshipment_settings,comment_permission,created_time,updated_time,review_info,question,excerpt,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp,upvoted_followees;data[*].mark_infos[*].url;data[*].author.follower_count,badge[?(type=best_answerer)].topics',
            'offset':'0',
            'limit':'20',
            'sort_by':'default'
        }
        url = 'https://www.zhihu.com/api/v4/questions/{}/answers?'.format(question_id) + urlencode(data)
    try:
        response = requests.get(url,headers= headers)
        print(response.status_code)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException as e:
        print(e.args)
        print('获取页面出错')


def parse_index(html):
    global question_name
    try:
        pattern = re.compile(r'href="https://www.zhihu.com/question/(\d*?)" class="internal">(.*?)</a>')
        question_id_list = re.findall(pattern,html)
        for question_info in question_id_list[9:]:
            question_name = question_info[1]
            print(question_name)
            yield question_info[0]
    except:
        print('解析索引页错误')
        return None


def parse_next(html):
    try:
        json_str = json.loads(html)
        is_end = json_str.get('paging').get('is_end')
        next_url = json_str.get('paging').get('next')
        if is_end == 'false':
            return next_url
        else:
            return None
    except:
        print('解析下一页出错')


def parse_data(html,question_name):
    global questions_name
    questions_name = question_name
    try:
        json_str = json.loads(html)
        for data in json_str['data']:

            thread = threading.Thread(target=parse_answer,args=(data,))
            thread.start()
    except:
        print('解析数据出错')

def parse_answer(data):
    try:
        lock.acquire()
        pattern_name = re.compile(r"\'name\': \'(.*?)\'",re.S)
        pattern_id = re.compile(r"'id': (.*?),",re.S)
        pattern_href = re.compile(r'data-actualsrc="(.*?)"', re.S)
        name = re.search(pattern_name,str(data)).group(1)
        answer_id = re.search(pattern_id,str(data)).group(1)
        href_list = re.findall(pattern_href,str(data))
        if len(href_list) != 0:
            save_file(name,answer_id,href_list)
    except:
        print('解析答案失败')
    finally:
        lock.release()


def save_file(name,answer_id,href_list):
    i = 0
    for href in href_list:
        content = get_content(href)
        if content:
            try:
                i += 1
                file_path = 'D:/zhihu/daiguang2/' + questions_name
                path = file_path + '/' + answer_id + name + str(i) + '.jpg'
                if not os.path.exists(file_path):
                    os.makedirs(file_path)
                if not os.path.exists(path):
                    with open(path,'wb') as f:
                        f.write(content)
                    print('保存图片成功',answer_id + name + str(i))
            except:
                print('存储图片失败')


def get_content(href):
    try:
        response = requests.get(href,headers=headers)
        print(response.status_code)
        time.sleep(random.random())
        if response.status_code == 200:
            return response.content
        else:
            return None
    except ConnectionError as e:
        print(e.args)
        print('获取图片失败')


def main():
    try:
        html = get_html(i=-1)
        for question_id in parse_index(html):
            i = 1
            next_url = True
            while next_url:
                print('第 %s 页' % (i))
                html = get_html(question_id=question_id,url=next_url,i=i)
                process = Process(target=parse_data, args=(html,question_name))
                process.start()
                next_url = parse_next(html)
                i += 1
                print('--------------------------------------------')
    except:
        print('爬取出错')
    finally:
        print('答案抓取完成')


if __name__ == '__main__':
    main()