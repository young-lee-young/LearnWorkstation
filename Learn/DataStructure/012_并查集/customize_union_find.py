# -*- coding: utf-8 -*-
# @Time:    2021/1/18 下午11:28
# @Author:  leeyoung
# @File:    customize_union_find.py
# @Content: 并查集

import random


# 第一版并查集
class CustomizeUnionFindOne:
    def __init__(self, size):
        self.data = [None] * size

    # 构造函数，初始化数组，数据存储的是集合的id
    def constructor(self):
        for i in range(len(self.data)):
            self.data[i] = i

    def size(self):
        return len(self.data)

    # 查找num对应的集合编号，O(1)
    def find(self, num):
        print(num)
        return self.data[num]

    # 合并元素num1和num2所属集合，Quick Find下的union复杂度为O(n)
    def union(self, num1, num2):
        set1 = self.find(num1)
        set2 = self.find(num2)
        if set1 == set2:
            return

        # 将集合编号为set1的修改为set2
        for index in range(len(self.data)):
            if self.data[index] == set1:
                self.data[index] = set2

    # Quick Find，查看元素num1和num2是否属于一个集合，O(1)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


# 第二版并查集
class CustomizeUnionFindTwo:
    def __init__(self, size):
        self.parent = [None] * size

    # 构造函数，数据存储的是父节点id
    def constructor(self):
        for i in range(len(self.parent)):
            self.parent[i] = i

    def size(self):
        return len(self.parent)

    # 寻找一个节点的根节点
    def find(self, num):
        while num != self.parent[num]:
            num = self.parent[num]
        return num

    # 将节点2的根节点设置为节点1的根节点，O(h)
    def union(self, num1, num2):
        root1 = self.find(num1)
        root2 = self.find(num2)
        if root1 == root2:
            return
        self.parent[root2] = root1

    # 如果两个节点的根节点相同，说明两个节点连接，O(h)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


# 第三版并查集，基于每个集合元素个数进行优化，将元素少的集合合并到元素多的集合
class CustomizeUnionFindThree:
    def __init__(self, size):
        self.parent = [None] * size
        # 集合包含节点个数
        self.set_size = [None] * size

    # 构造函数，数据存储的是父节点id
    def constructor(self):
        for i in range(len(self.parent)):
            self.parent[i] = i
            # 开始时每个集合节点数都为1
            self.set_size[i] = 1

    def size(self):
        return len(self.parent)

    # 寻找一个节点的根节点
    def find(self, num):
        while num != self.parent[num]:
            num = self.parent[num]
        return num

    # 将节点2的根节点设置为节点1的根节点，O(h)
    def union(self, num1, num2):
        root1 = self.find(num1)
        root2 = self.find(num2)
        if root1 == root2:
            return
        # num1所在集合节点数大于num2所在集合节点数，将集合2合并到集合1
        if self.set_size[root1] > self.set_size[root2]:
            self.parent[root2] = root1
            # 将集合1的节点树更新
            self.set_size[root1] += self.set_size[root2]
        # 将集合1合并到集合2
        else:
            self.parent[root1] = root2
            self.set_size[root2] += self.set_size[root1]

    # 如果两个节点的根节点相同，说明两个节点连接，O(h)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


# 第四版并查集，基于集合树高度进行优化，将高度小的和合并到高度大的
class CustomizeUnionFindFour:
    def __init__(self, size):
        self.parent = [None] * size
        # 集合树高度
        self.rank = [None] * size

    # 构造函数，数据存储的是父节点id
    def constructor(self):
        for i in range(len(self.parent)):
            self.parent[i] = i
            # 开始时每个集合节点数都为1
            self.rank[i] = 1

    def size(self):
        return len(self.parent)

    # 寻找一个节点的根节点
    def find(self, num):
        while num != self.parent[num]:
            num = self.parent[num]
        return num

    # 将节点2的根节点设置为节点1的根节点，O(h)
    def union(self, num1, num2):
        root1 = self.find(num1)
        root2 = self.find(num2)
        if root1 == root2:
            return
        # num1所在集合树高度小于num2所在集合树高度，将集合1合并到集合2
        if self.rank[root1] < self.rank[root2]:
            self.parent[root1] = root2
        # num1所在集合树高度大于num2所在集合树高度，将集合2合并到集合1
        elif self.rank[root1] > self.rank[root2]:
            self.parent[root2] = root1
        # 相等，将集合1合并到集合2，并且集合2的高度加1
        else:
            self.parent[root1] = root2
            self.rank[root2] += 1

    # 如果两个节点的根节点相同，说明两个节点连接，O(h)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


# 第五版并查集，路径压缩，将当前节点的父节点设置为爷爷节点
class CustomizeUnionFindFive:
    def __init__(self, size):
        self.parent = [None] * size
        # 集合树高度
        self.rank = [None] * size

    # 构造函数，数据存储的是父节点id
    def constructor(self):
        for i in range(len(self.parent)):
            self.parent[i] = i
            # 开始时每个集合节点数都为1
            self.rank[i] = 1

    def size(self):
        return len(self.parent)

    # 寻找一个节点的根节点
    def find(self, num):
        while num != self.parent[num]:
            # 路径压缩，将当前节点的父节点设置爷爷节点
            self.parent[num] = self.parent[self.parent[num]]
            num = self.parent[num]
        return num

    # 将节点2的根节点设置为节点1的根节点，O(h)
    def union(self, num1, num2):
        root1 = self.find(num1)
        root2 = self.find(num2)
        if root1 == root2:
            return
        # num1所在集合树高度小于num2所在集合树高度，将集合1合并到集合2
        if self.rank[root1] < self.rank[root2]:
            self.parent[root1] = root2
        # num1所在集合树高度大于num2所在集合树高度，将集合2合并到集合1
        elif self.rank[root1] > self.rank[root2]:
            self.parent[root2] = root1
        # 相等，将集合1合并到集合2，并且集合2的高度加1
        else:
            self.parent[root1] = root2
            self.rank[root2] += 1

    # 如果两个节点的根节点相同，说明两个节点连接，O(h)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


# 第六版并查集，路径压缩，将当前节点所有祖先节点都设置为根节点
class CustomizeUnionFindSix:
    def __init__(self, size):
        self.parent = [None] * size
        # 集合树高度
        self.rank = [None] * size

    # 构造函数，数据存储的是父节点id
    def constructor(self):
        for i in range(len(self.parent)):
            self.parent[i] = i
            # 开始时每个集合节点数都为1
            self.rank[i] = 1

    def size(self):
        return len(self.parent)

    # 寻找一个节点的根节点
    def find(self, num):
        # 如果当前节点不是根节点
        if num != self.parent[num]:
            # 将递归父节点
            self.parent[num] = self.find(self.parent[num])
        return self.parent[num]

    # 将节点2的根节点设置为节点1的根节点，O(h)
    def union(self, num1, num2):
        root1 = self.find(num1)
        root2 = self.find(num2)
        if root1 == root2:
            return
        # num1所在集合树高度小于num2所在集合树高度，将集合1合并到集合2
        if self.rank[root1] < self.rank[root2]:
            self.parent[root1] = root2
        # num1所在集合树高度大于num2所在集合树高度，将集合2合并到集合1
        elif self.rank[root1] > self.rank[root2]:
            self.parent[root2] = root1
        # 相等，将集合1合并到集合2，并且集合2的高度加1
        else:
            self.parent[root1] = root2
            self.rank[root2] += 1

    # 如果两个节点的根节点相同，说明两个节点连接，O(h)
    def is_connected(self, num1, num2):
        return self.find(num1) == self.find(num2)


version = 6
# 测试第一版并查集
size = 10000
if version == 1:
    union_find = CustomizeUnionFindOne(size)
elif version == 2:
    union_find = CustomizeUnionFindTwo(size)
elif version == 3:
    union_find = CustomizeUnionFindThree(size)
elif version == 4:
    union_find = CustomizeUnionFindFour(size)
elif version ==  5:
    union_find = CustomizeUnionFindFive(size)
else:
    union_find = CustomizeUnionFindSix(size)

union_find.constructor()
for i in range(size):
    num1 = random.randint(0, size - 1)
    num2 = random.randint(0, size - 1)
    print('----- union -----', num1, num2)
    union_find.union(num1, num2)

for i in range(size):
    num1 = random.randint(0, size - 1)
    num2 = random.randint(0, size - 1)
    print('----- is connected -----', num1, num2)
    union_find.is_connected(num1, num2)
print(union_find)
