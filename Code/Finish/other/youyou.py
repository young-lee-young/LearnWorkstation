# 爬取优优网的图片和信息

import requests
from multiprocessing import Pool
from lxml import etree
import re
import os
from bs4 import BeautifulSoup

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}


def get_index_html(i):
    url = 'http://m.233mr.com/nvyou/{}/'.format(i)
    try:
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            for name_href in parse_index_html(response.text):
                html = get_detail_html(name_href[1])
                av_name_href_notsplit = parse_detail_html(html, name_href[0])
                av_name_hrefs = av_name_href_notsplit[2:]
                for i in range(len(av_name_hrefs)):
                    av_name_href = av_name_hrefs[i]
                    html = get_detail_html(av_name_href[1])
                    parse_av(html, name_href[0], i)
        else:
            return None
    except:
        print('获取主页错误', url)


def parse_index_html(html):
    try:
        selector = etree.HTML(html)
        a_list = selector.xpath('/html/body/div/section[3]/nav/a')
        for a in a_list:
            href = a.get('href')
            name = a.get('title')
            yield [name, href]
    except:
        print('解析主页错误')


def get_detail_html(url):
    try:
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            response.encoding = 'utf-8'
            return response.text
        else:
            return None
    except:
        print('获取详情页错误', url)


def parse_detail_html(html, name):
    try:
        pattern = re.compile(r'img src="(.*?)" alt=', re.S)
        image_url_list = re.findall(pattern, html)
        for i in range(len(image_url_list)):
            content = get_image_html(image_url_list[i], name)
            download_image(name, content, i)
        soup = BeautifulSoup(html, 'html.parser')
        li_list = soup.find('div', attrs={'class': 'news-list'}).find_all('li')
        av_name_href = []
        for li in li_list:
            av_name = li.find('h3').get_text()
            href = li.find('a').get('href')
            av_name_href.append((av_name, href))
        return av_name_href
    except:
        print('解析详情页出错')


def parse_av(html, nvyou_name, i):
    soup = BeautifulSoup(html, 'html.parser')
    div = soup.find('div', attrs={'class': 'letter-box'})
    ps = div.find_all('p')
    for p in ps:
        text = p.get_text()
        write_av(nvyou_name, i, text=text)
    image_href = ps[-1].find('img').get('src')
    image_content = get_image_html(image_href, nvyou_name)
    write_av(nvyou_name, i, image_content=image_content)


def write_av(nvyou_name, i, text=None, image_content=None):
    base_path = 'D:/youyou/' + nvyou_name
    file_path = base_path + '/' + str(i) + '.txt'
    try:
        if text:
            if not os.path.exists(base_path):
                os.makedirs(base_path)
            with open(file_path, 'a') as f:
                f.writelines(text + '\n')
        else:
            image_path = base_path + '/' + str(i) + '.jpg'
            if os.path.exists(base_path):
                with open(image_path, 'wb') as f:
                    f.write(image_content)
                    print('保存图片成功', nvyou_name)
    except:
        print('保存文件失败')


def get_image_html(url, name):
    try:
        response = requests.get(url, headers=headers)
        if response.status_code == 200:
            return response.content
        else:
            return None
    except:
        print('获取图片出错', url, name)


def download_image(name, content, i):
    base_path = 'D:/youyou/' + name + '/main'
    path = base_path + '/' + str(i) + '.jpg'
    try:
        if not os.path.exists(base_path):
            os.makedirs(base_path)
        if not os.path.exists(path):
            with open(path, 'wb') as f:
                f.write(content)
            print('保存图片成功', name)
        else:
            print('图片已存在')
    except:
        print('保存图片失败', name)


def main():
    pool = Pool(processes=4)
    for i in range(1, 70):
        pool.apply_async(get_index_html, (i,))
    pool.close()
    pool.join()
    print('------------------完成-------------------')


if __name__ == '__main__':
    main()
