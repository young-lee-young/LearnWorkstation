# matplotlib画饼图pie

import matplotlib.pyplot as plt

plt.rcParams['font.sans-serif'] = ['SimHei']
plt.rcParams['axes.unicode_minus'] = False

fig1 = plt.figure(num='饼图1', figsize=(16,9))
labels = ['中国', '美国', '日本', '欧盟', '印度'] # 每个部分的名字
sizes = [30, 40, 20, 30, 10]    # 每个部分所占的大小
# color = ['yellowgreen', 'gold', 'lightskyblue', 'r', 'pink']
explode = (0, 0, 0, 0, 0)   # 每个部分到圆心的距离


# 0度开始,设置字体大小
ax1_1 = fig1.add_subplot(2,2,1)
patches, label_texts, percent_texts = ax1_1.pie(sizes, explode=explode, labels=labels, autopct='%1.2f%%', shadow=False, startangle=0)
for label_text in label_texts:  # 设置字体label大小
    label_text.set_size(20)
for percent_text in percent_texts:  # 设置比例字体大小
    percent_text.set_size(15)
ax1_1.axis('equal') # 横纵坐标轴统一刻度,园才会是正圆
plt.legend(loc='upper left')
ax1_1.set_title('0度开始,设置字体大小')


# 90度开始
ax1_2 = fig1.add_subplot(2,2,2)
ax1_2.pie(sizes, explode=explode, labels=labels, autopct='%1.2f%%', shadow=False, startangle=90)
ax1_2.axis('equal')
plt.legend(loc='upper right')
ax1_2.set_title('90度开始')


# 有阴影,不自动生成比例
ax1_3 = fig1.add_subplot(2,2,3)
ax1_3.pie(sizes, explode=explode, labels=labels, shadow=True, startangle=0)
ax1_3.axis('equal')
plt.legend(loc='lower left')
ax1_3.set_title('有阴影,不自动生成比例')


# 凸起一部分
ax1_4 = fig1.add_subplot(2,2,4)
explode2 = (0.1, 0, 0, 0, 0)
ax1_4.pie(sizes, explode=explode2, labels=labels, autopct='%1.2f%%', shadow=False, startangle=0)
ax1_4.axis('equal')
plt.legend(loc='lower right')
ax1_4.set_title('凸起一部分')

fig1.savefig('D:/spider/piefig1.png', api=300)
plt.show()