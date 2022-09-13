import requests
import json
import time
import random
import os
import datetime
import multiprocessing
import threading

headers = {
    'Accept':'application/json',
    'Host':'huaban.com',
    'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
    'X-Request':'JSON',
    'X-Requested-With':'XMLHttpRequest'
}

params = {
    'max':'',
    'limit':'20',
}

params2 = {
    'jamjlgol': '',
    'max':'',
    'limit':'100',
}


def get_board_id(url):
    set_board_id = set()
    try:
        response = requests.get(url, params=params, headers=headers)
        # sleep_time = random.randint(20,30)
        # time.sleep(sleep_time)
        print(response.status_code)
        if response.status_code == 200:
            pins = response.json()['pins']
            if len(pins) == 0:
                params['max'] = ''
                return None
            else:
                for pin in pins:
                    board_id = pin['board']['board_id']
                    if board_id not in set_board_id:
                        set_board_id.add(board_id)
                        print(set_board_id)
                        get_key(board_id)
                    else:
                        print('已下载过此画板')
                last_pin_id = pins[-1]['pin_id']
                print(last_pin_id)
                params['max'] = last_pin_id
                get_board_id(url)
        else:
            print(response.status_code)
            return None
    except ConnectionError:
        print('连接错误')
        return None

def get_key(board_id):
    url = 'http://huaban.com/boards/' + str(board_id) + '/'
    print(url)
    try:
        response = requests.get(url,params=params2,headers = headers)
        # sleep_time = random.randint(10,15)
        # time.sleep(sleep_time)
        # print(response.status_code)
        # print('---------------------------------------------------------------------------------------------------------')
        # print(response.text)
        if response.status_code == 200:
            pins = response.json()['board']['pins']
            print(type(pins))
            if len(pins) == 0:
                params2['max'] = ''
                return None
            else:
                for pin in pins:
                    key = pin['file']['key']
                    get_pic(key)
                last_pin_id = pins[-1]['pin_id']
                print(last_pin_id)
                params2['max'] = last_pin_id
                get_key(board_id)
        else:
            print(response.status_code)
            return None
    except ConnectionError:
        print('连接错误')
        return None

def get_pic(key):
    url = 'http://img.hb.aicdn.com/' + key
    headers2 = {
        'User-Agent':'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
    }
    try:
        response = requests.get(url,headers = headers2)
        sleep_time = random.randint(1,2)
        time.sleep(sleep_time)
        # time.sleep(0.5)
        if response.status_code == 200:
            write_to_file(response.content,key)
        else:
            return None
    except ConnectionError:
        print('连接错误')
        return None

def write_to_file(content,key):
    start_time = time.time()
    root_path = 'D:/huaban/'
    pic_path = root_path + key + '.jpg'
    if not os.path.exists(root_path):
        os.mkdir(root_path)
    if not os.path.exists(pic_path):
        with open(pic_path,'wb') as f:
            f.write(content)
            f.close()
            nowtime = datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S')
            print(pic_path,'保存成功',nowtime)
            end_time = time.time()
            print('time is :',end_time - start_time)

def main(key):
    url = 'http://huaban.com/favorite/{}'.format(key)
    get_board_id(url)

if __name__ == '__main__':
    main('beauty')