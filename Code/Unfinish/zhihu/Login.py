import re
import time
import json
import requests
import hmac
import hashlib
import base64
from http import cookiejar
from PIL import Image
import matplotlib.pyplot as plt
import settings


class Zhihulogin(object):
    def __init__(self):
        self.login_url = settings.LOGIN_URL
        self.login_api = settings.LOGIN_API
        self.login_data = settings.FORM_DATA
        self.session = requests.session()
        self.session.headers = settings.HEADERS.copy()
        self.session.cookies = cookiejar.LWPCookieJar(filename='cookies.txt')

    def login(self, username=None, password=None, load_cookies=True):
        if load_cookies and self.load_cookies():
            if self.check_login():
                return True
        headers = self.session.headers.copy()
        headers.update({
            'authorization': 'oauth c3cef7c66a1843f8b3a9e6a1e3160e20',
            'x-xsrftoken': self.get_token(),
        })
        timestamp = str(int(time.time()*1000))
        self.login_data.update({
            'username': username,
            'password': password,
            'timestamp': timestamp,
            'signature': self.get_signature(timestamp),
            'captcha': self.get_captcha(self.login_data, headers),
        })
        response = self.session.post(self.login_api, data=self.login_data, headers=headers)
        if 'error' in response.text:
            print(re.findall(r'"message":"(.+?)"', response.text)[0])
            print('登陆失败')
        elif self.check_login():
            return True
        return False

    # 加载cookies
    def load_cookies(self):
        try:
            self.session.cookies.load(ignore_discard=True)
            return True
        except FileNotFoundError:
            print('没找到文件')
            return False

    def check_login(self):
        response = self.session.get(self.login_url, allow_redirects=False)
        if response.status_code == 302:
            self.session.cookies.save()
            return True
        return False

    def get_token(self):
        response = self.session.get(self.login_url)
        pattern = re.compile(r'_xsrf=([\w|-]+)')
        token = re.findall(pattern, response.headers.get('Set-Cookie'))[0]
        return token

    def get_signature(self, timestamp):
        ha = hmac.new(b'd1b964811afb40118a12068ff74a12f4', digestmod=hashlib.sha1)
        grant_type = self.login_data['grant_type']
        client_id = self.login_data['client_id']
        source = self.login_data['source']
        ha.update(bytes((grant_type + client_id + source + timestamp), 'utf-8'))
        return ha.hexdigest()

    def get_captcha(self, login_data, headers):
        lang = login_data.get('lang')
        if lang == 'cn':
            api = settings.CAPTCHA_API_CN
        else:
            api = settings.CAPTCHA_API_EN
        response = self.session.get(api, headers=headers)
        show_captcha = re.search(r'true', response.text)
        if show_captcha:
            put_response = self.session.put(api, headers=headers)
            img_base64 = re.findall(r'"img_base64":"(>+)"', put_response.text, re.S)[0].replace(r'\n', '')
            with open('captcha.jpg', 'wb') as f:
                f.write(base64.b64encode(img_base64))
            img = Image.open('captcha.jpg')
            if lang == 'cn':
                plt.imshow(img)
                print('请点击倒立的汉字，按回车提交')
                points = plt.ginput(7)
                captcha = json.dumps({'img_size': [200, 44], 'input_points': [[i[0]/2, i[1]/2] for i in points]})
            else:
                img.show()
                captcha = input('请输入验证码:')
            self.session.post(api, data={'input_text': captcha}, headers=headers)
            return captcha
        return ''


zhihulogin = Zhihulogin()
zhihulogin.login(username='18401681943', password='lyzy131', load_cookies=True)
