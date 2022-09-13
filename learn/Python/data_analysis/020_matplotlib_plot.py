# matplotlib画直线plot

import matplotlib.pyplot as plt # 导入的常规写法
from numpy.random import randn
import numpy as np

# 创建绘画窗口和子图
fig1 = plt.figure(num='一个子图', figsize=(16,9), dpi=80, facecolor=(1,0,1,1), edgecolor='r')
# num可以是字符,也可以是数字, figsize设置绘图窗口的大小, dpi设置分辨率, facecolor窗口背景的颜色, edgecolor边框颜色
fig1.add_subplot(1,1,1) # 添加一个子图
fig1.subplots_adjust(left=None,right=None,top=None,bottom=None,wspace=0,hspace=0)
# 有两个要求left cannot be >= right,bottom cannot be >= top,也就是要求往左上偏
print(fig1)
# Figure(1280x720)
print(type(fig1))
# <class 'matplotlib.figure.Figure'>


fig2 = plt.figure(num='九个子图',figsize=(16,9))
ax1_1 = fig2.add_subplot(3,3,1)   # 添加子图,子图是3*3格式的,一共九个单元格,取第一个
print(ax1_1)
# Axes(0.125,0.653529;0.227941x0.226471)
print(type(ax1_1))
# <class 'matplotlib.axes._subplots.AxesSubplot'>


# 画线
ax1_1.plot([1,2,3,4],[1,2,3,4], color='#ADFF2F', linestyle='--', linewidth=1, marker='>')
# color设置线的颜色,b(blue,蓝色)、c(cyan,青色)、g(green,绿色)、k(black,黑色)、m(magenta,品红)、r(red,红色)、w(white,白色)、y(yellow,黄色)
# 第二种方法:使用RGB值表示,如:'#ADFF2F'
# 第三中方法:灰度值,1.0代表是白色,0.0代表是黑色
# linestyle设置线的类型,linestyle可选的有'-'(solid,实线)、'--'(dashed,虚线)、'-.'(dash dot,杠点线)、':'(dotted点线), linewidth设置线的宽度,
# marker设置类型,标记实际点形状,可选见附录
plt.title('plot1')   # 设置子图的标题
plt.axis('on')  # 是否启用坐标轴,on为启用,off为不启用


# 设置横纵坐标轴的范围
ax1_2 = fig2.add_subplot(3,3,2)
ax1_2.plot([50,60,70,80,90,100], color='r', linestyle='-', marker='*')
ax1_2.set_xlim([0,5])
ax1_2.set_ylim([0,100])
ax1_2.set_title('plot2')


# 设置横纵坐标轴的说明
ax1_3 = fig2.add_subplot(3,3,3)
ax1_3.set_xticklabels(['','zhangya','liyao','liwei','liuyafei','macheng'])
ax1_3.set_yticklabels(['','cha','zhong','liang','you'])
ax1_3.set_xlabel('name')
ax1_3.set_ylabel('score')
ax1_3.set_title('plot3')


# 在一个图里画三条线,并且添加图例
ax1_4 = fig2.add_subplot(3,3,4)
ax1_4.plot(randn(1000).cumsum(), color='r', linestyle='-', label='one')
ax1_4.plot(randn(1000).cumsum(), color='b', linestyle='--', label='two')
ax1_4.plot(randn(1000).cumsum(), color='y', linestyle=':', label='three')
ax1_4.set_xlim([0,1000])
# ax1_4.set_ylim([-40,40])
ax1_4.legend(loc='best')        # 设置图例,最好的位置
ax1_4.set_title('plot4')


# 横纵坐标确定特定的点
ax1_5 = fig2.add_subplot(3,3,5)
ax1_5.plot([1,2,4,8,16,32,64,128], color='g', linestyle='-.', marker='s')
ax1_5.set_xticks([0,1,2,3,4,5,6,7])
ax1_5.set_yticks([1,2,4,8,16,32,64,128])
ax1_5.set_title('plot5')


# 线形图中,默认是以线性方式插值的,可以通过drawstyle选项修改
'''
还不知道这几个drawstyle的区别
'''
ax1_6 = fig2.add_subplot(3,3,6)
ax1_6.plot(randn(10).cumsum(), linestyle='--', label='default')
ax1_6.plot(randn(10).cumsum(), linestyle='--', drawstyle='steps', label='steps')
ax1_6.plot(randn(10).cumsum(), linestyle='--', drawstyle='steps-pre', label='steps-pre')
ax1_6.plot(randn(10).cumsum(), linestyle='--', drawstyle='steps-mid', label='steps-mid')
ax1_6.plot(randn(10).cumsum(), linestyle='--', drawstyle='steps-post', label='steps-post')
ax1_6.legend(loc='best')
ax1_6.set_title('plot6')


# 绘制sin和cos曲线
ax1_7 = fig2.add_subplot(3,3,7)
x = np.linspace(-np.pi, np.pi, num=256, endpoint=False)
sin = np.sin(x)
cos = np.cos(x)
ax1_7.plot(x, sin, color='r', linestyle='-', label='sin')
ax1_7.plot(x, cos, color='b', linestyle='--', label='cos')
plt.legend(loc='best')
ax1_7.set_title('plot7')


ax1_8 = fig2.add_subplot(3,3,8)
ax1_9 = fig2.add_subplot(3,3,9)

# 保存图片
# fig2.savefig('D:/spider/plotfig2.png', dpi=400, bbox_inches='tight')
plt.show()


# 颜色附录
# cnames = {
# 'aliceblue':            '#F0F8FF',
# 'antiquewhite':         '#FAEBD7',
# 'aqua':                 '#00FFFF',
# 'aquamarine':           '#7FFFD4',
# 'azure':                '#F0FFFF',
# 'beige':                '#F5F5DC',
# 'bisque':               '#FFE4C4',
# 'black':                '#000000',
# 'blanchedalmond':       '#FFEBCD',
# 'blue':                 '#0000FF',
# 'blueviolet':           '#8A2BE2',
# 'brown':                '#A52A2A',
# 'burlywood':            '#DEB887',
# 'cadetblue':            '#5F9EA0',
# 'chartreuse':           '#7FFF00',
# 'chocolate':            '#D2691E',
# 'coral':                '#FF7F50',
# 'cornflowerblue':       '#6495ED',
# 'cornsilk':             '#FFF8DC',
# 'crimson':              '#DC143C',
# 'cyan':                 '#00FFFF',
# 'darkblue':             '#00008B',
# 'darkcyan':             '#008B8B',
# 'darkgoldenrod':        '#B8860B',
# 'darkgray':             '#A9A9A9',
# 'darkgreen':            '#006400',
# 'darkkhaki':            '#BDB76B',
# 'darkmagenta':          '#8B008B',
# 'darkolivegreen':       '#556B2F',
# 'darkorange':           '#FF8C00',
# 'darkorchid':           '#9932CC',
# 'darkred':              '#8B0000',
# 'darksalmon':           '#E9967A',
# 'darkseagreen':         '#8FBC8F',
# 'darkslateblue':        '#483D8B',
# 'darkslategray':        '#2F4F4F',
# 'darkturquoise':        '#00CED1',
# 'darkviolet':           '#9400D3',
# 'deeppink':             '#FF1493',
# 'deepskyblue':          '#00BFFF',
# 'dimgray':              '#696969',
# 'dodgerblue':           '#1E90FF',
# 'firebrick':            '#B22222',
# 'floralwhite':          '#FFFAF0',
# 'forestgreen':          '#228B22',
# 'fuchsia':              '#FF00FF',
# 'gainsboro':            '#DCDCDC',
# 'ghostwhite':           '#F8F8FF',
# 'gold':                 '#FFD700',
# 'goldenrod':            '#DAA520',
# 'gray':                 '#808080',
# 'green':                '#008000',
# 'greenyellow':          '#ADFF2F',
# 'honeydew':             '#F0FFF0',
# 'hotpink':              '#FF69B4',
# 'indianred':            '#CD5C5C',
# 'indigo':               '#4B0082',
# 'ivory':                '#FFFFF0',
# 'khaki':                '#F0E68C',
# 'lavender':             '#E6E6FA',
# 'lavenderblush':        '#FFF0F5',
# 'lawngreen':            '#7CFC00',
# 'lemonchiffon':         '#FFFACD',
# 'lightblue':            '#ADD8E6',
# 'lightcoral':           '#F08080',
# 'lightcyan':            '#E0FFFF',
# 'lightgoldenrodyellow': '#FAFAD2',
# 'lightgreen':           '#90EE90',
# 'lightgray':            '#D3D3D3',
# 'lightpink':            '#FFB6C1',
# 'lightsalmon':          '#FFA07A',
# 'lightseagreen':        '#20B2AA',
# 'lightskyblue':         '#87CEFA',
# 'lightslategray':       '#778899',
# 'lightsteelblue':       '#B0C4DE',
# 'lightyellow':          '#FFFFE0',
# 'lime':                 '#00FF00',
# 'limegreen':            '#32CD32',
# 'linen':                '#FAF0E6',
# 'magenta':              '#FF00FF',
# 'maroon':               '#800000',
# 'mediumaquamarine':     '#66CDAA',
# 'mediumblue':           '#0000CD',
# 'mediumorchid':         '#BA55D3',
# 'mediumpurple':         '#9370DB',
# 'mediumseagreen':       '#3CB371',
# 'mediumslateblue':      '#7B68EE',
# 'mediumspringgreen':    '#00FA9A',
# 'mediumturquoise':      '#48D1CC',
# 'mediumvioletred':      '#C71585',
# 'midnightblue':         '#191970',
# 'mintcream':            '#F5FFFA',
# 'mistyrose':            '#FFE4E1',
# 'moccasin':             '#FFE4B5',
# 'navajowhite':          '#FFDEAD',
# 'navy':                 '#000080',
# 'oldlace':              '#FDF5E6',
# 'olive':                '#808000',
# 'olivedrab':            '#6B8E23',
# 'orange':               '#FFA500',
# 'orangered':            '#FF4500',
# 'orchid':               '#DA70D6',
# 'palegoldenrod':        '#EEE8AA',
# 'palegreen':            '#98FB98',
# 'paleturquoise':        '#AFEEEE',
# 'palevioletred':        '#DB7093',
# 'papayawhip':           '#FFEFD5',
# 'peachpuff':            '#FFDAB9',
# 'peru':                 '#CD853F',
# 'pink':                 '#FFC0CB',
# 'plum':                 '#DDA0DD',
# 'powderblue':           '#B0E0E6',
# 'purple':               '#800080',
# 'red':                  '#FF0000',
# 'rosybrown':            '#BC8F8F',
# 'royalblue':            '#4169E1',
# 'saddlebrown':          '#8B4513',
# 'salmon':               '#FA8072',
# 'sandybrown':           '#FAA460',
# 'seagreen':             '#2E8B57',
# 'seashell':             '#FFF5EE',
# 'sienna':               '#A0522D',
# 'silver':               '#C0C0C0',
# 'skyblue':              '#87CEEB',
# 'slateblue':            '#6A5ACD',
# 'slategray':            '#708090',
# 'snow':                 '#FFFAFA',
# 'springgreen':          '#00FF7F',
# 'steelblue':            '#4682B4',
# 'tan':                  '#D2B48C',
# 'teal':                 '#008080',
# 'thistle':              '#D8BFD8',
# 'tomato':               '#FF6347',
# 'turquoise':            '#40E0D0',
# 'violet':               '#EE82EE',
# 'wheat':                '#F5DEB3',
# 'white':                '#FFFFFF',
# 'whitesmoke':           '#F5F5F5',
# 'yellow':               '#FFFF00',
# 'yellowgreen':          '#9ACD32'}

# marker附录:
# '.'       point marker
# ','       pixel marker
# 'o'       circle marker
# 'v'       triangle_down marker
# '^'       triangle_up marker
# '<'       triangle_left marker
# '>'       triangle_right marker
# '1'       tri_down marker
# '2'       tri_up marker
# '3'       tri_left marker
# '4'       tri_right marker
# 's'       square marker
# 'p'       pentagon marker
# '*'       star marker
# 'h'       hexagon1 marker
# 'H'       hexagon2 marker
# '+'       plus marker
# 'x'       x marker
# 'D'       diamond marker
# 'd'       thin_diamond marker
# '|'       vline marker
# '_'       hline marker