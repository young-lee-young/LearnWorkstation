
# 系统常量配置

HEADERS = {
    'Host': 'www.zhihu.com',
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/'
                  '61.0.3163.100 Safari/537.36 OPR/48.0.2685.52',
}


LOGIN_URL = 'https://www.zhihu.com/signup'

LOGIN_API = 'https://www.zhihu.com/api/v3/oauth/sign_in'

CAPTCHA_API_EN = 'https://www.zhihu.com/api/v3/oauth/captcha?lang=en'
CAPTCHA_API_CN = 'https://www.zhihu.com/api/v3/oauth/captcha?lang=cn'

FORM_DATA = {
    'client_id': 'c3cef7c66a1843f8b3a9e6a1e3160e20',
    'grant_type': 'password',
    # 时间戳, 这个不需要
    # 'timestamp': '1522828703820',
    'source': 'com.zhihu.web',
    'username': '',
    'password': '',
    # en是英文验证码, cn是中文验证码
    'lang': 'en',
    'ref_source': '_other',
}


