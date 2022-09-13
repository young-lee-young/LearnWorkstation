# numpy下random模块的使用

from numpy import random


seed = random.seed(10)  # 随机数种子,有随机数种子后,以后随机生成的数字是以种子为基础的,seed方法不返回任何值
print(seed)
# None


# 产生指定范围的整数随机数,并指定产生的个数
array1 = random.randint(0,20,10)   # 产生10个
print(type(array1))
# <class 'numpy.ndarray'>
print('10个整数随机数:\n',array1)
array2 = random.randint(0,20,[2,4])    # 产生2行4列数组
print('2行4列整数随机数:\n',array2)


# 产生0-1的随机数,等概率
array3 = random.rand(2,4)  # 2行4列的数组
print('2行4列0-1随机数:\n',array3)


# 产生正态分布随机数
array4 = random.randn(2,4)
print('2行4列正态分布随机数:\n',array4)
arr_sum = array4.sum()
print('正态分布和为:',arr_sum)




# 附录:
# 正态分布特点:
# 1、对称,对称轴是平均数
# 2、中央点最高,两侧下降
# 3、曲线下的面积是1
# 标准正态分布是正态分布的特例,平均值是0,方差是1