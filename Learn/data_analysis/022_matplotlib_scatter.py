# matplotlib画散点图scatter

import matplotlib.pyplot as plt
import numpy as np

plt.rcParams['font.sans-serif'] = ['SimHei']
plt.rcParams['axes.unicode_minus'] = False

fig1 = plt.figure(num='散点图1', figsize=(16,9))


ax1_1 = fig1.add_subplot(2,2,1)
x = np.random.randint(1,100,1000)
y = np.random.randint(1,100,1000)
ax1_1.scatter(x, y, s=30, c='c', marker='o')
ax1_1.set_xlabel('x_label',fontsize=20)
ax1_1.set_ylabel('y_label',fontsize=20)
# s是点的面积,默认是20, c是颜色,默认是blue


ax1_2 = fig1.add_subplot(2,2,2)
x = np.random.rand(1,1024)
y = np.random.rand(1,1024)
color = np.arctan2(y, x)
ax1_2.scatter(x, y, s=75, c=color, alpha=0.5, marker='o', edgecolor='g')
plt.xticks([])
plt.yticks([])

# fig1.savefig('D:/spider/scatterfig1.png', dpi=400)
plt.show()