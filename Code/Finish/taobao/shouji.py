from selenium import webdriver
from bs4 import BeautifulSoup
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException
import time
import re
import pymongo

client = pymongo.MongoClient('localhost')
db = client['taobao']
service_args = ['--load-images=false','--disk-cache=true']
browser = webdriver.PhantomJS(service_args=service_args)
browser.set_window_size(1400,900)
wait = WebDriverWait(browser,10)

def search(keyword,url):
	try:
		browser.get(url)
		#判断已经加载
		inputbox = wait.until(EC.presence_of_element_located((By.ID,'q')))
		#判断已经可以点击
		submit = wait.until(EC.element_to_be_clickable((By.XPATH,'//*[@id="J_TSearchForm"]/div[1]/button')))
		inputbox.send_keys(keyword)
		submit.click()
		#得到一共多少页
		print('正在获取第 1 页')
		total_page = wait.until(EC.presence_of_element_located((By.XPATH,'//*[@id="spulist-pager"]/div/div/div/div[1]')))
		get_products(1)
		return total_page.text
	except TimeoutException:
		return search(url)

def next_page(page_number):
	try:
		print('正在获取第 %s 页' % (page_number))
		inputbox = wait.until(EC.presence_of_element_located((By.XPATH,'//*[@id="spulist-pager"]/div/div/div/div[2]/input')))
		submit = wait.until(EC.element_to_be_clickable((By.XPATH,'//*[@id="spulist-pager"]/div/div/div/div[2]/span[3]')))
		inputbox.clear()
		inputbox.send_keys(page_number)
		submit.click()
		wait.until(EC.text_to_be_present_in_element((By.XPATH,'//*[@id="spulist-pager"]/div/div/div/ul/li[@class="item active"]/span'),str(page_number)))
		get_products(page_number)
	except TimeoutException:
		next_page(page_number)

def get_products(i):
	print('正在解析第 %s 页' % (i))
	wait.until(EC.presence_of_element_located((By.XPATH,'//*[@id="J_SPUBlankRow11"]/div[4]')))
	for i in range(0,8):
		js = 'window.scrollTo(0,(document.body.scrollHeight/10)*{})'.format(i)
		browser.execute_script(js)
		time.sleep(1)
	html = browser.page_source
	soup = BeautifulSoup(html, 'html.parser')
	product_list = soup.find_all('div', attrs={'class': 'grid-item col'})
	print(len(product_list))
	for product in product_list:
		try:
			name = product.find('a', attrs={'class': 'product-title'}).get('title')
			screen = product.find('span', attrs={'class': 'important-key'}).get_text().strip()
			price = product.find('strong').get_text().strip()
			deal = product.find('span', attrs={'class': 'num'}).get_text().strip()
			info = {
				'name':name,
				'screen':screen,
				'price':price,
				'deal':deal
			}
			save_to_mongo(info)
		except:
			print('解析错误')

def save_to_mongo(info):
	try:
		db['shouji'].insert(info)
		print('存储成功',info['name'][:4])
	except:
		print('存储失败',info['name'][:4])

def main(keyword):
	print('关键字是',keyword)
	url = 'http://www.taobao.com'
	try:
		total_page = search(keyword,url)
		pattern = re.compile(r'(\d+)',re.S)
		pages = int(re.search(pattern,total_page).group(1))
		for i in range(2,pages+1):
			next_page(i)
	except Exception:
		print('抓取出错')
	finally:
		browser.close()

if __name__ == '__main__':
	keyword = '手机'
	main(keyword)