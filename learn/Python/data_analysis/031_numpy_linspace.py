# linspace用法(等差数列)

import numpy as np


# 两个参数时,默认是50个等差数列,endpoint为False,表示最后一个数不一定是第二个参数,默认为True,retstep为True,显示等差,默认为False
linspace1 = np.linspace(1,10,endpoint=False,retstep=True)
print(linspace1)


# 第三个参数是等差数列的个数
linspace2 = np.linspace(1,10,num=10)
print(linspace2)
