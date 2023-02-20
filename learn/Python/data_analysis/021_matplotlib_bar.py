# matplotlib画条形图bar和直方图hist

from numpy.random import randn
import numpy as np
import matplotlib.pyplot as plt

plt.rcParams['font.sans-serif'] = ['SimHei']    # 正常显示中文
plt.rcParams['axes.unicode_minus'] = False  # 正常显示负号


fig1 = plt.figure('条形图1',figsize=(16,9))

# 竖向条形图
ax1_1 = fig1.add_subplot(2,2,1)
x = np.arange(10)
y = np.random.randint(1,10,10)
ax1_1.bar(left=x, height=y, width=0.5, color='r', orientation='vertical')
ax1_1.set_title('竖向条形图',bbox={'facecolor':'lightgreen','pad':5})
# left:条形图下面的横坐标, height:条形图的高度, width:每个条形图的宽度, color:条形图颜色, orientation:垂直还是水平


# 横向条形图
ax1_2 = fig1.add_subplot(2,2,2)
label = ['a','b','c','d','e','f']
x = [100,300,900,300,500,400]
y = np.arange(6)
colors = np.random.rand(6,3)
# colors = np.random.rand(6,4)  # 加上alpha
ax1_2.barh(y, x, color=colors)
plt.yticks(y,label)
ax1_2.grid(axis='x')
ax1_2.set_title('横向条形图')


# 横向双条形图
ax1_3 = fig1.add_subplot(2,2,3)
label2 = ['A','B','C','D','E']
x = np.arange(5)
y1 = np.random.randint(1,10,5)
y2 = np.random.randint(1,10,5)
ax1_3.bar(x-0.2, y1, width=0.4, color='b', label='2017年')
ax1_3.bar(x+0.2, y2, width=0.4, color='y', label='2018年')
ax1_3.grid(axis='y')
plt.xticks(x,label2)
plt.legend(loc='best')
# 图例的位置,best(0,最好的位置)、upper right(1,右上角)、upper left(2,左上角)、lower left(3,左下角)、lower right(4,右下角)、right(5,右面)、
# cennter left(6,左中)、center right(7,右中)、lower center(8,下中)、upper center(9,上中)、center(10,正中)
ax1_3.set_title('横向双条形图')


# 竖向双条形图
ax1_4 = fig1.add_subplot(2,2,4)
label3 =['','星期日','星期一','星期二','星期三','星期四','星期五','星期六']
x = np.arange(7)
y1 = np.random.randint(1,10,7)
y2 = np.random.randint(1,10,7)
ax1_4.bar(left=x, height=y1, width=0.4, color='b', label='李二狗')
ax1_4.bar(left=x, height=y2, bottom=y1, width=0.4, color='y', label='赵三猫')
ax1_4.set_xticklabels(label3)
ax1_4.grid(True)    # 显示网格
plt.legend(loc='upper left')
ax1_4.set_title('竖向双条形图')

# fig1.savefig('D:/spider/barfig1.png', dpi=400)


fig2 = plt.figure('条形图2',figsize=(16,9))

# 添加注释
ax2_1 = fig2.add_subplot(2,2,1)
x = np.arange(10)
y = np.random.randint(1,10,10)
bars1 = ax2_1.bar(left=x, height=y, width=0.4, color='coral')
for bar1 in bars1:
    x_bar = bar1.get_x()
    height = bar1.get_height()
    ax2_1.text(x_bar, 1.02*height, str(height), family='monospace', fontsize='10')
ax2_1.set_title('添加注释')


# 图例在表格外面
'''
暂时还没有用到,等用过,学习会了后再来整理
'''
ax2_2 = fig2.add_subplot(2,2,2)
x = np.arange(8)
y1 = np.random.randint(1,10,8)
y2 = np.random.randint(1,10,8)
ax2_2.bar(x-0.2, y1, 0.4, color='#FF69B4', label='中国')
ax2_2.bar(x+0.2, y2, 0.4, color='#000080', label='美国')
plt.legend(loc='best')
ax2_2.set_title('图例在表格外面')


# 向两边延伸的横向条形图
ax2_3 = fig2.add_subplot(2,2,3)
x1 = np.random.randint(10,20,8)
x2 = np.random.randint(10,20,8)
y = np.arange(1,9)
bars2 = ax2_3.barh(y, x1, height=0.8, color='r')
bars3 = ax2_3.barh(y, -x2, height=0.8, color='y')
frame = plt.gca()
frame.axes.get_yaxis().set_visible(False)   # 设置y轴不可见
# frame.axes.get_xaxis().set_visible(False)   # 设置x轴不可见
for bar2 in bars2:
    y_bar2 = bar2.get_y()
    width_bar2 = bar2.get_width()
    ax2_3.text(width_bar2+0.5, y_bar2+0.3, str(width_bar2),color='g')
for bar3 in bars3:
    y_bar3 = bar3.get_y()
    width_bar3 = bar3.get_width()
    ax2_3.text(-width_bar3-2.5, y_bar3+0.3, str(width_bar3),color='g')
ax2_3.set_title('向两边延伸的横向条形图')

# fig2.savefig('D:/spider/barfig2.png', dpi=400)


# 概率密度分布图
fig3, axes3 = plt.subplots(2,2,sharex=True,sharey=False)    # 分享轴

for i in range(2):
    if i == 0:
        data1 = randn(5000)
        axes3[i,0].hist(data1, bins=50, color='c', alpha=0.5)  # 概率密度分布
        axes3[i,1].hist(data1, bins=50, color='c', alpha=0.5, cumulative=True)    # 累积概率密度分布
    else:
        mean = 0
        sigma = 1
        data2 = mean + sigma*randn(10000)   # 和上面的是一样的算法,只是更容易理解
        axes3[i,0].hist(data2, bins=50, color='c', alpha=0.5)
        axes3[i,1].hist(data2, bins=50, color='c', alpha=0.5, cumulative=True)
plt.subplots_adjust(wspace=0, hspace=0) # 设置子图间的距离,0为没有距离

# fig3.savefig('D:/spider/barfig3.png', dpi=400)

plt.show()