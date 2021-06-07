# -*- coding: utf-8 -*-
# @Time:    2021/3/3 23:01
# @Author:  leeyoung
# @File:    customize_graph_base_adjacency_list.py
# @Content: 无向图 - 基于临接表


# 空间复杂度：O(v + e)，v为顶点个数，e为边数
class Graph:
    def __init__(self, matrix):
        # 顶点个数
        self.vertices = None
        # 边个数
        self.edges = None
        # 临接表，用链表实现，优化的话使用红黑树
        self.adjacency_list = None
        # 原矩阵
        self.matrix = matrix
        # 遍历时判断是否被访问过
        self.visited = list()

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edges = self.matrix[0][1]
        self.visited = [False for _ in range(self.vertices)]

        # 初始化临接表
        self.adjacency_list = [[] for _ in range(self.vertices)]
        for i, j in self.matrix[1:]:
            # 检查顶点合法性
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception('有自环边，非简单图')
            # 检查是否有平行边
            if i in self.adjacency_list[j]:
                raise Exception('有平行边，非简单图')

            # 记录每个顶点相邻的顶点
            self.adjacency_list[i].append(j)
            self.adjacency_list[j].append(i)

    # 获取顶点个数
    def get_vertices(self):
        return self.vertices

    # 获取边个数
    def get_edge(self):
        return self.edges

    # 判读顶点m和n是否相邻，时间复杂度：O(degree(v))
    def has_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)
        return v in self.adjacency_list[w]

    # 获取当前顶点相邻的所有顶点列表，时间复杂度：O(degree(v))
    def adj(self, v):
        self.validate_vertex(v)
        return self.adjacency_list[v]

    # 获取顶点的度
    def degree(self, v):
        return len(self.adj(v))

    # 从图中删除一条边
    def remove_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)

        self.adjacency_list[v].remove(w)
        self.adjacency_list[w].remove(v)

        self.edges -= 1

    def print_matrix(self):
        for index, item in enumerate(self.adjacency_list):
            print(index, ': ', item)

    # 检查顶点是否合法
    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception('顶点不合法')


# 一个联通分量
matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
# 两个联通分量
matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

graph = Graph(matrix)
# 测试生成函数
graph.generate()
graph.print_matrix()
