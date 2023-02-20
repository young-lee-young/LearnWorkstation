# 下载、解析及sleep用的包
import requests
from bs4 import BeautifulSoup
import time
import random
# 写入文件用的包
import csv
# 分词等用的包
import re
import pandas as pd
import jieba
import numpy
import matplotlib.pyplot as plt
from IPython import get_ipython
# get_ipython().magic('matplotlib inline')
# 绘制词云用的包
import matplotlib

matplotlib.rcParams['figure.figsize'] = (10.0, 5.0)
from wordcloud import WordCloud, ImageColorGenerator

session = requests.Session()
headers = {
    'User-Agent': 'Moilla/5.0(Macintosh;Intel Mac OS X 10_12_4)AppleWebKit/537.36(KHTML,Like Gecko) Chrome/57.0.2987.133 Safari/537.36',
}


def main(movie_id):
    # login()
    # for html in get_html(movie_id):
    #     for comment_info in parsecomments(html):
    #         writecomment(comment_info)
    comment_list = read_comment()
    comments = get_string(comment_list)
    world_cloud(comments)


def get_html(movie_id):
    for i in range(25):
        url = 'https://movie.douban.com/subject/' + movie_id + '/comments?start={}&limit=20'.format(i * 20)
        response = session.get(url, headers=headers)
        print(response.status_code)
        time.sleep(random.random() * 5)
        if response.status_code == 200:
            if i == 0:
                parseinfo(response.text)
                yield response.text
            else:
                yield response.text


def parseinfo(html):
    soup = BeautifulSoup(html, 'html.parser')
    ps = soup.body.find('span', attrs={'class': 'attrs'}).find_all('p')
    for p in ps:
        keys = p.find('span', attrs={'class': 'pl'}).get_text()
        names = p.find_all('a')
        for name in names:
            name_text = name.get_text()
            print(keys, name_text)


def parsecomments(html):
    soup = BeautifulSoup(html, 'html.parser')
    comment_list = soup.find_all('div', attrs={'class': 'comment-item'})
    for comment_item in comment_list:
        user_name = comment_item.find('span', attrs={'class': 'comment-info'}).a.get_text().strip()
        commnet = comment_item.find('p').get_text().strip()
        yield (user_name, commnet)


def writecomment(comment):
    try:
        with open('D:/zhanlang2.csv', 'a', newline='') as f:
            writer = csv.writer(f, dialect='excel')
            writer.writerow(comment)
            print('存储成功', comment[0])
    except:
        print('存储失败', comment[0])


def read_comment():
    try:
        with open('D:/spider/zhanlang2.csv', 'r') as f:
            reader = csv.reader(f)
            comment = [row[1] for row in reader]
            return comment
    except:
        print('读取失败')


def get_string(comment_list):
    comments = ''
    for i in range(len(comment_list)):
        comments = comments + comment_list[i]
    return comments


def world_cloud(comments):
    # 去除标点
    pattern = re.compile(r'[\u4e00-\u9fa5]+')
    filterdata = re.findall(pattern, comments)
    cleaned_comments = ''.join(filterdata)
    # print(cleaned_comments)

    # 分词
    segment = jieba.lcut(cleaned_comments)
    words_df = pd.DataFrame({'segment': segment})
    # print(words_df.head())      #head()默认5个

    # 停用词
    #     stopwords = [line.strip() for line in open('D:/stopwords.txt','r',encoding = 'gbk').readlines()]
    stopwords = pd.read_csv('D:/spider/stopwords.txt', index_col=False, quoting=3, sep="\t", names=['stopword'],
                            encoding='gbk')
    new_words_df = words_df[~words_df.segment.isin(stopwords.stopword)]
    # print(new_words_df.head())

    # 统计词频
    words_stat = new_words_df.groupby(by=['segment'])['segment'].agg({'count': numpy.size}).reset_index().sort_values(
        by=['count'], ascending=False)
    # print(words_stat)

    # 词云
    wordcloud = WordCloud(font_path='D:/spider/simhei.ttf', background_color='white', width=800, height=600, max_font_size=80)
    #    wordcloud = wordcloud.fit_words(words_stat.itertuples(index = False))
    word_frequence = {x[0]: x[1] for x in words_stat.head(1000).values}
    #     word_frequence_list = []
    #     for key in word_frequence:
    #         temp = (key,word_frequence[key])
    #         word_frequence_list.append(temp)
    wordcloud = wordcloud.fit_words(word_frequence)
    wordcloud.to_file('D:/cloud.jpg')
    plt.figure('战狼词云')
    plt.imshow(wordcloud)
    plt.axis('off')
    plt.show()


def login():
    loginurl = 'https://accounts.douban.com/login'
    logindata = {
        'source': 'index_nav',
        'redir': 'https://www.douban.com',
        'form_email': '18400000000',
        'form_password': 'helloworld',
        'login': u'登录'
    }
    response = session.post(loginurl, data=logindata, headers=headers)
    soup = BeautifulSoup(response.text, 'html.parser')
    captcha = soup.find('img', id='captcha_image')
    if captcha:
        captcha_url = captcha['src']
        pattern = r'<input type="hidden" name="captcha-id" value="(.*?)"/'
        captcha_id = re.findall(pattern, response.text)
        print(captcha_url)
        captcha_text = input('please input:')
        logindata['captcha-solution'] = captcha_text
        logindata['captcha-id'] = captcha_id
        session.post(loginurl, data=logindata, headers=headers)


if __name__ == '__main__':
    movie_id = '26363254'
    main(movie_id)