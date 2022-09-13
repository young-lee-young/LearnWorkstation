import requests
import json
import time
import random
from urllib.parse import urlencode
from requests.exceptions import RequestException
import pymongo

client = pymongo.MongoClient('localhost')
db = client['zhihu']

headers = {
    'authorization':'Bearer Mi4xS1hZRUFnQUFBQUFBRUFLQTlaR19EQmNBQUFCaEFsVk5LYjhPV3dCd0x6Vld4OUVQYnc2a2YtZXg3bC1XVmMwV2h3|1512141097|cceaa8e6be9f280aeaf0cf2303b4480f0fa0ef82',
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}


def get_answer_first_html(question_id):
    data = {
        'include': 'data[*].is_normal,admin_closed_comment,reward_info,is_collapsed,annotation_action,annotation_detail,collapse_reason,is_sticky,collapsed_by,suggest_edit,comment_count,can_comment,content,editable_content,voteup_count,reshipment_settings,comment_permission,created_time,updated_time,review_info,question,excerpt,relationship.is_authorized,is_author,voting,is_thanked,is_nothelp,upvoted_followees;data[*].mark_infos[*].url;data[*].author.follower_count,badge[?(type=best_answerer)].topics',
        'offset': '0',
        'limit': '20',
        'sort_by': 'default'
    }
    url = 'https://www.zhihu.com/api/v4/questions/{}/answers?'.format(question_id) + urlencode(data)
    try:
        response = requests.get(url, headers=headers)
        print(response.status_code)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        print('获取第一页出错')


def get_answer_other_html(url):
    try:
        response = requests.get(url,headers=headers)
        print(response.status_code)
        time.sleep(random.random()*5)
        if response.status_code == 200:
            return response.text
        return None
    except RequestException:
        print('获取页面出错',url)


def parse_answer_html(html):
    result = json.loads(html)
    if result and 'data' in result.keys():
        datas = result.get('data')
        is_end = result.get('paging').get('is_end')
        next_url = result.get('paging').get('next')
        for data in datas:
            author_name = data.get('author').get('name')
            content = data.get('content').replace('<br>','').replace('<p>','').replace('</p>','')
            print(content)
            # save_to_mongo(author_name,content)
        if not is_end:
            return next_url
        else:
            return None


def save_to_mongo(author,content):
    data = {
        'author':author,
        'answer':content
    }
    try:
        db['xuwei'].insert(data)
        print('%s的答案插入成功' % (author))
    except:
        print('插入数据失败',author)


def main(question_id):
    i = 1
    print('第 1 页')
    first_html = get_answer_first_html(question_id)
    next_url = parse_answer_html(first_html)
    while next_url:
        i += 1
        print('第 %s 页' % (i))
        other_html = get_answer_other_html(next_url)
        next_url = parse_answer_html(other_html)
    print('-----------------------答案抓取完成------------------------')


if __name__ == '__main__':
    question_id = '21543621'
    main(question_id)