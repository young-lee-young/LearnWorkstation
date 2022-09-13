#本代码用requests进行网页请求，用正则表达式进行解析,开四个进程

import requests
from requests.exceptions import RequestException
# from bs4 import BeautifulSoup
import re
import time
import json
from multiprocessing import Pool

headers = {
    'User-Agent':'Mozilla/5.0(Macintosh;Intel Mac OS X 10_12_4)'
                 'AppleWebKit/537.36(KHTML,Like Gecko) Chrome/57.0.2987.133 Safari/537.36',
}

def get_html(url):
	try:
		response = requests.get(url,headers = headers)
		if(response.status_code == 200):
			return response.text
		return None
	except RequestException:
		return None

def parse_html(html):
	pattern = re.compile('<dd>.*?board-index.*?>(\d+)</i>.*?title="(.*?)".*?data-src="(.*?)".*?class="star">(.*?)</p>.*?time">(.*?)</p>.*?integer">(.*?)</i>.*?fraction">(.*?)</i>.*?</dd>',re.S)
	items = re.findall(pattern,html)
	for item in items:
		yield{
			'index':item[0],
			'title':item[1],
			'image':item[2],
			'actor':item[3].strip()[3:],
			'time':item[4].strip()[5:],
			'comment':item[5] + item[6],
		}

def writefile(content):
	#这样写进文件是Unicode编码，不是汉字
	# with open('D:/maoyan.txt','a') as f:
	# 	f.write(json.dumps(content) + '\n')
	with open('D:/maoyan.txt','a',encoding = 'utf-8') as f:
		f.write(json.dumps(content,ensure_ascii = False) + '\n')
		f.close()

def main(i):
	url = 'http://maoyan.com/board/4?offset={}'.format(i*10)
	html = get_html(url)
	for item in parse_html(html):
		writefile(item)

if __name__ == '__main__':
	pool = Pool()
	pool.map(main,[i for i in range(0,10)])
	pool.close()
	pool.join()