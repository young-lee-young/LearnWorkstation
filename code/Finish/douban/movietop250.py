# coding: utf-8

# 本例使用requests库进行网页请求，得到HTML，用BeautifulSoup进行网页解析，写入Excel文件
import requests
from lxml import etree
import xlwt
import os

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36 OPR/48.0.2685.52'
}


def get_html(url, flag):
    response = requests.get(url, headers=headers)
    if response.status_code == 200:
        if flag == 1:
            return response.text
        else:
            return response.content


def parse_html(html):
    page = etree.HTML(html)
    lis = page.xpath('//*[@id="content"]/div/div[1]/ol/li')
    for li in lis:
        chinese_name = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="hd"]/a/span[1]/text()')[0].strip()
        other_name = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="hd"]/a/span[2]/text()')[0].strip()
        director = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="bd"]/p[1]/text()')[0].strip()
        comment_num = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="bd"]/div[@class="star"]/span[last()]/text()')[0].strip()
        rating_num = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="bd"]/div[@class="star"]/span[@class="rating_num"]/text()')[0].strip()
        quote = li.xpath('div[@class="item"]/div[@class="info"]/div[@class="bd"]/p[@class="quote"]/span/text()')[0]
        if len(quote) == 0:
            quote == '暂无简介'
        pic_url = li.xpath('div[@class="item"]/div[@class="pic"]/a/img/@src')[0]
        yield [chinese_name, other_name, director, comment_num, rating_num, quote, pic_url]


def save_file(info_list):
    for info in info_list:
        print('保存图片 :', info[0])
        save_pic(info[0], info[-1])
    save_excel(info_list)


def save_pic(name, url):
    content = get_html(url, 2)
    path = 'movietop250/'
    pic_path = path + name + '.jpg'
    if not os.path.exists(path):
        os.mkdir(path)
    if not os.path.exists(pic_path):
        with open(pic_path, 'wb') as f:
            f.write(content)


def save_excel(info_list):
    file = xlwt.Workbook(encoding='utf-8', style_compression=0)
    sheet = file.add_sheet('豆瓣电影TOP250', cell_overwrite_ok=True)
    col = ('中文名', '其它名', '导演及演员', '评论数', '评分', '简介')
    for i in range(0, 6):
        sheet.write(0, i, col[i])
    for i, info in enumerate(info_list, 1):
        for j in range(0, 6):
            sheet.write(i, j, info[j])
        print('写入 :', info[0])
    file.save('movietop250/top250.xls')


def main():
    info_list = []
    for i in range(0, 10):
        url = 'https://movie.douban.com/top250?start={}&filter='.format(i*25)
        html = get_html(url, 1)
        for infos in parse_html(html):
            info_list.append(infos)
    save_file(info_list)


if __name__ == '__main__':
    main()
