import itchat
import pymongo
import numpy as np
import pandas as pd
from pandas import DataFrame
import matplotlib.pyplot as plt
import re
import jieba
from wordcloud import WordCloud
from PIL import Image


plt.rcParams['font.sans-serif'] = ['SimHei']    # 正常显示中文
plt.rcParams['axes.unicode_minus'] = False  # 正常显示负号

client = pymongo.MongoClient(host='localhost', port=27017)
db_auth = client.admin
db_auth.authenticate('mongdb_username', 'password')
db = client.wechat


def login():
    itchat.login()  # 会自动跳转出一个二维码,扫描后登陆


def get_friends():
    all_friends = []
    print('获取好友列表中...')
    for friend in itchat.get_friends(update=True):
        save_txt(friend)
        friend_popdone = pop_keys(friend)
        all_friends.append(friend_popdone)
    print('%s.txt成功保存' % name)
    return all_friends


def save_txt(friend):
    with open('D:/spider/' + name + '.txt', 'a', encoding='utf-8') as f:
        f.writelines(str(friend) + '\n')


""""
获取到的朋友信息有很多事用不到的，直接扔掉，以方便数据入库
"""
def pop_keys(friend):
    new_friend_dict = {}
    friend_dict = dict(friend)
    key_list = ['UserName', 'City', 'Province', 'RemarkName', 'Signature', 'NickName',
                'RemarkPYQuanPin', 'HeadImgUrl', 'Sex', 'AttrStatus', 'SnsFlag', 'ContactFlag']
    for key in friend_dict:
        if key in key_list:
            new_friend_dict[key] = friend_dict[key]
    return new_friend_dict


def count_save(friends):
    data = count(friends)
    save_to_csv(data)
    save_to_mongo(friends)


def count(friends):
    key_list = ['NickName', 'Province', 'City', 'Sex', 'Signature']
    data = {}
    for key in key_list:
        variable = []
        for friend in friends:
            variable.append(friend[key])
        data[key] = variable
    return data


def save_to_csv(data):
    df = pd.DataFrame(data)
    df.to_csv('D:/spider/' + name + '.csv', index=True)
    print('%s.csv成功保存' % name)


def save_to_mongo(friends):
    if db[name].insert(friends):
        print('插入数据库成功')


def analysis_data():
    data_dict = read_mongo()
    dataframe = DataFrame(data_dict)
    sex(dataframe)
    address(dataframe)
    wordcloud(dataframe)
    plt.show()


def read_mongo():
    NickName_list = []
    Province_list = []
    City_list = []
    Sex_list = []
    Signature_list = []
    datas = db[name].find({}, {'NickName': 1, 'Province': 1, 'City': 1, 'Sex': 1, 'Signature':1})
    for data in datas:
        NickName_list.append(data['NickName'])
        Province_list.append(data['Province'])
        City_list.append(data['City'])
        Sex_list.append(data['Sex'])
        Signature_list.append(data['Signature'])
    return {
        'NickName': NickName_list,
        'Province': Province_list,
        'City': City_list,
        'Sex': Sex_list,
        'Signature': Signature_list
    }


def create_figure(num):
    figure = plt.figure(num=num, figsize=(16, 9))
    return figure


def sex(dataframe):
    figure = create_figure(num='好友性别')
    sex_list = count_sex(dataframe)
    labels = ['male', 'female', 'unkown']
    sex_bar(figure, sex_list, labels)
    sex_pie(figure, sex_list, labels)


def count_sex(dataframe):
    male = (dataframe['Sex'] == 1).sum()
    female = (dataframe['Sex'] == 2).sum()
    unkown = (dataframe['Sex'] == 0).sum()
    return [male, female, unkown]


def sex_bar(figure, sex_list, labels):
    ax = figure.add_subplot(1, 2, 1)
    ax.set_title('性别人数')
    x = np.arange(0,3)
    colors = ['b', 'pink', 'g']
    bars = ax.bar(x, sex_list, color=colors, width=0.4)
    for bar in bars:
        x_bar = bar.get_x()
        height = bar.get_height()
        ax.text(x_bar + 0.15, height + 3, str(int(height)), color='r', fontsize='10')
    plt.xticks(x, labels)


def sex_pie(figure, sex_list, labels):
    ax = figure.add_subplot(1,2,2)
    ax.set_title('男女比例')
    ax.pie(sex_list, labels=labels, autopct='%1.2f%%')
    ax.axis('equal')
    plt.legend(loc='upper right')


def address(dataframe):
    address_dict = count_address(dataframe)
    address_dict['province'][0] = '未知'
    address_bar(address_dict)
    address_pie(address_dict)


def count_address(dataframe):
    province_list = []
    num = []
    groupby = dataframe.groupby('Province')
    for name, group in groupby:
        province_list.append(name)
        num.append(len(group))
    return {
        'province': province_list,
        'num': num
    }


def address_bar(address_dict):
    figure = create_figure(num='好友分布')
    ax = figure.add_subplot(1,1,1)
    ax.set_title('好友分布柱状图')
    x = np.arange(0,len(address_dict['province']))
    for i in x:
        colors = np.random.rand(1, 3)
        ax.bar(x[i], address_dict['num'][i], width=0.6, color=colors, label=address_dict['province'][i])
    plt.legend(loc='upper right')
    ax.grid(axis='y')
    plt.xticks(x, address_dict['province'])


def address_pie(address_dict):
    figure = create_figure(num='好友分布2')
    ax = figure.add_subplot(1,1,1)
    ax.set_title('好友分布饼图')
    ax.pie(address_dict['num'], labels=address_dict['province'], autopct='%1.2f%%')
    ax.axis('equal')
    plt.legend(loc='upper right')


def wordcloud(dataframe):
    stopword = get_signature(dataframe)
    generate_wc(','.join(stopword))
    fitword_wc(stopword)


def get_signature(dataframe):
    signature = dataframe['Signature'].sum()
    return pop_punctuaion(signature)


def pop_punctuaion(signature):
    pattern = re.compile(r'[\u4e00-\u9fa5]+')
    filterword = re.findall(pattern, signature)
    chinese_string = ''.join(filterword)
    return cut_word(chinese_string)


def cut_word(chinese_string):
    segment = jieba.lcut(chinese_string)
    return stop_words(segment)


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
    mask = np.array(Image.open('D:/spider/background.png'))
    wordcloud_class = WordCloud(font_path='D:/spider/simhei.ttf', background_color='yellow', max_words=2000, max_font_size=80, width=1000, height=700, mask=mask)
    return wordcloud_class


def make_wordcloud(wordcloud_obj, figure_name):
    figure = create_figure(num=figure_name)
    plt.imshow(wordcloud_obj, cmap=plt.cm.gray)
    plt.axis('off')


def main():
    global name
    login()
    name = str(input('please input filename and mongodbname:'))
    all_friends = get_friends()
    count_save(all_friends)
    analysis_data()


if __name__ == '__main__':
    main()