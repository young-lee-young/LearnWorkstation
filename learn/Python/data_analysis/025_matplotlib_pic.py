# matplotlib处理图片

import matplotlib.pyplot as plt

# 显示肉片
img = plt.imread('D:/spider/me.jpg')    # 读取图片文件
fig2,axes2 = plt.subplots(2,2,figsize=(16,9))
ax2_1, ax2_2, ax2_3, ax2_4 = axes2.ravel()  # 把四个实例分开

ax2_1.set_title('original')
ax2_1.axis('off')
ax2_1.imshow(img)

ax2_2.set_title('R channel')
ax2_2.axis('off')
ax2_2.imshow(img[:,:,0],plt.cm.gray)

ax2_3.set_title('G channel')
ax2_3.axis('off')
ax2_3.imshow(img[:,:,1],plt.cm.gray)

ax2_4.set_title('B channel')
ax2_4.axis('off')
ax2_4.imshow(img[:,:,2],plt.cm.gray)

plt.show()