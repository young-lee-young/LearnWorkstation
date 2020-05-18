import time
import requests
from bs4 import BeautifulSoup
import json
import os
import re
import jieba
from pandas import DataFrame
from PIL import Image
from wordcloud import WordCloud
import matplotlib.pyplot as plt
import numpy as np
from snownlp import SnowNLP


plt.rcParams['font.sans-serif'] = ['SimHei']    # 正常显示中文
plt.rcParams['axes.unicode_minus'] = False  # 正常显示负号

headers = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36',
    'Referer': 'http://music.163.com/',
    'Host': 'music.163.com',
}


def lyric_spider(artist_list):
    for artist_name, artist_id in artist_list:
        url = 'http://music.163.com/artist?id=' + artist_id
        html = get_html(url)
        song_infos = parse_song_id(html)
        for song_name, lyric in get_lyrics(song_infos):
            save_file(artist_name + '-' + song_name, lyric)


def get_html(url):
    response = requests.get(url, headers=headers)
    print(response.status_code)
    return response.text


def parse_song_id(html):
    song_infos = []
    soup = BeautifulSoup(html, 'html.parser')
    lis = soup.find('ul', attrs={'class': 'f-hide'}).find_all('li')
    for li in lis:
        title = li.find('a').get_text()
        song_id = li.find('a').get('href').split('=')[-1]
        song_infos.append((title,song_id))
    return song_infos


def get_lyrics(song_infos):
    for song_info in song_infos:
        url = 'http://music.163.com/api/song/lyric?id=' + str(song_info[1]) + '&lv=1&kv=1&tv=-1'
        html = get_html(url)
        lyric_json = json.loads(html)
        if 'lrc' in lyric_json.keys():
            lyric = lyric_json['lrc']['lyric']
            yield (song_info[0], lyric)


def save_file(name, content):
    with open('D:/spider/163cloudmusic/' + name + '.txt', 'w', encoding='utf-8', errors='ignore') as f:
        f.write(content)
    print('写入:' + name)


def analysis(artist_list):
    # sentiment_analysis(artist_list)
    wordcloud()
    plt.show()


def sentiment_analysis(artist_list):
    sentiment_dict = {}
    for artist in artist_list:
        score = read_singer_lyric(artist[0])
        sentiment_dict[artist[0]] = score
        print(sentiment_dict)
    sentiment_plot(sentiment_dict)


def read_singer_lyric(artist):
    lyrics = []
    files = os.listdir('D:/spider/163cloudmusic')
    for file in files:
        if artist in file:
            with open('D:/spider/163cloudmusic/' + file, encoding='utf-8') as f:
                for lyric in f.readlines():
                    if '作曲' not in lyric and '作词' not in lyric:
                        lyrics.append(lyric)
    return calculate_score(lyrics)


def calculate_score(lyrics):
    lyric_num = 0
    score = 0.0
    for lyric in lyrics:
        lyric_string = pop_punctuaion(lyric)
        if len(lyric_string) > 3:
            lyric_num += 1
            snownlp = SnowNLP(lyric_string)
            # sen = snownlp.sentences
            score += snownlp.sentiments
    return score/lyric_num


def sentiment_plot(sentiment_dict):
    figure = plt.figure(num='歌词情感分析', figsize=(16, 9))
    ax = figure.add_subplot(1, 1, 1)
    sort = input('need sort? if yes please input 1, no please input other')
    x, scores, labels = sort_score_label(sentiment_dict, int(sort))
    for i in x:
        color = np.random.rand(1, 3)
        ax.bar(x[i], scores[i], width=0.5, label=labels[i], color=color)
    plt.xticks(x, labels)
    plt.ylim([0.5, 0.8])
    plt.legend(loc='upper right')
    plt.grid(axis='y')
    figure.savefig('D:/spider/163cloudmusic/lyric_sentiment.png', api=800)


def sort_score_label(sentiment_dict, sort):
    x = np.arange(0, len(sentiment_dict))
    if sort == 1:
        scores = []
        labels = []
        zipinfo = zip(sentiment_dict.values(), sentiment_dict.keys())
        sort_info = sorted(zipinfo, reverse=True)
        for score, label in sort_info:
            scores.append(score)
            labels.append(label)
    else:
        scores = list(sentiment_dict.values())
        labels = list(sentiment_dict.keys())
    return x, scores, labels


def wordcloud():
    lyrics = read_lyrics()
    chinese_string = pop_punctuaion(lyrics)
    print('歌词长度:' + str(len(chinese_string)))
    segment = cut_word(chinese_string)
    stop_word = stop_words(segment)
    generate_wc(','.join(stop_word))
    fitword_wc(stop_word)


def read_lyrics():
    lyrics = ''
    files = os.listdir('D:/spider/163cloudmusic')
    for file in files:
        file_path = 'D:/spider/163cloudmusic/' + file
        with open(file_path, 'r', encoding='utf-8') as f:
            lyrics += f.read()
    return lyrics


def pop_punctuaion(lyrics):
    pattern = re.compile(r'[\u4e00-\u9fa5]+')
    filterword = re.findall(pattern, lyrics)
    chinese_string = ''.join(filterword)
    return chinese_string


def cut_word(chinese_string):
    segment = jieba.lcut(chinese_string)
    return segment


def stop_words(segment):
    stopword = []
    with open('D:/spider/stopwords.txt') as f:
        stopwords = f.read()
    stopword_list = stopwords.split('\n')
    for word in segment:
        if word not in stopword_list:
            stopword.append(word)
    return stopword


def generate_wc(word_string):
    wordcloud_obj = create_wc_class().generate(word_string)
    make_wordcloud(wordcloud_obj, 'generate词云')


def fitword_wc(stopword):
    word_df = DataFrame({'word_df': stopword})
    sort_word_df = count_word(word_df)
    words_dict = {key: num for key, num in sort_word_df.values}
    wordcloud_obj = create_wc_class().fit_words(words_dict)
    make_wordcloud(wordcloud_obj, 'fitword词云')


def count_word(word_df):
    new_word_df = word_df.groupby('word_df')['word_df'].agg({'count':np.size}).reset_index()
    sort_word_df = new_word_df.sort_values(by='count', ascending=False)
    return sort_word_df


def create_wc_class():
    # mask = np.array(Image.open('D:/spider/background.png'))
    mask = None
    wordcloud_class = WordCloud(font_path='D:/spider/simhei.ttf', background_color='white', max_words=2000, max_font_size=100, width=1000, height=700, mask=mask)
    return wordcloud_class


def make_wordcloud(wordcloud_obj, figure_name):
    figure = plt.figure(figure_name)
    plt.imshow(wordcloud_obj, cmap=plt.cm.gray)
    plt.axis('off')


def main():
    artist_list = [('许巍', '5770'), ('汪峰', '5347'), ('崔健', '2111'), ('郑钧', '6458'), ('伍佰', '5354'), ('窦唯', '2515'),
                   ('何勇', '3055'), ('张楚', '6455'), ('石磊', '1158182'),('黑豹乐队', '11759'), ('唐朝乐队', '12972'),
                   ('飘乐队', '12537'), ('Beyond', '11127'), ('零点乐队', '12115'), ('呼吸乐队', '11777'), ('指南针乐队', '13587'),
                   ('痛仰乐队', '12971'), ('逃跑计划', '12977'), ('鲍家街43号', '11094'), ('梁博', '166010'), ('天堂乐队', '12984'),
                   ('谢天笑', '5767'), ('臧天朔', '6512'), ('超载乐队', '11704'), ('左小祖咒', '6467'), ('轮回乐队', '12116'),
                   ('子曰', '13581')]
    # lyric_spider(artist_list)
    analysis(artist_list)


if __name__ == '__main__':
    main()