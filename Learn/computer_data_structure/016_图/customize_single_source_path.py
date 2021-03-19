# -*- coding: utf-8 -*-
# @Time:    2021/3/16 22:38
# @Author:  leeyoung
# @File:    customize_single_source_path.py
# @Content: 单源路径


class SingleSourcePath:
    def __init__(self, matrix, source):
        # 顶点个数
        self.vertices = None
        # 边个数
        self.edge = None
        # 临接表，用链表实现
        self.adjacency_list = None
        # 原矩阵
        self.matrix = matrix
        # 遍历时判断是否被访问过
        self.visited = list()
        # 源
        self.source = source
        # 每个顶点前面的顶点
        self.pre = list()

    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edge = self.matrix[0][1]
        self.visited = [False for i in range(self.vertices)]
        self.pre = [-1 for i in range(self.vertices)]
        self.adjacency_list = [[] for i in range(self.vertices)]

        for i, j in self.matrix[1:]:
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception("有自环表，非简单图")
            if i in self.adjacency_list[j]:
                raise Exception("有平行边，非简单图")

            self.adjacency_list[i].append(j)
            self.adjacency_list[j].append(i)

        # 对source进行深度优先遍历
        self.validate_vertex(self.source)
        self.dfs(self.source, self.source)

    # 获取当前顶点相邻的所有顶点列表
    def adj(self, v):
        self.validate_vertex(v)
        return self.adjacency_list[v]

    # 深度优先遍历
    def dfs(self, v, parent):
        # 标记已经访问过
        self.visited[v] = True
        self.pre[v] = parent
        adjacent = self.adj(v)
        for w in adjacent:
            if not self.visited[w]:
                self.dfs(w, v)

    # 判断是否和源点联通
    def is_connected(self, v):
        self.validate_vertex(v)
        return self.visited[v]

    # 源点到v的路径
    def path(self, v):
        result = list()
        if not self.is_connected(v):
            return result
        current = v
        while current != self.source:
            result.append(current)
            current = self.pre[current]
        result.append(self.source)
        result.reverse()
        return result

    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception("顶点不合法")


# 一个联通分量
matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
# 两个联通分量
matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

sspath = SingleSourcePath(matrix, 0)
# 测试生成函数
sspath.generate()

print("0 -> 6 : ", sspath.path(6))
print("0 -> 5 : ", sspath.path(5))
