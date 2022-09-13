# -*- coding: utf-8 -*-
# @Time:    2021/2/26 17:20
# @Author:  leeyoung
# @File:    customize_graph.py
# @Content: 无向图 - 基于邻接矩阵实现


# 空间复杂度：O(v^2)，v为顶点个数
class AdjacencyMatrix:
    def __init__(self, matrix):
        # 顶点个数
        self.vertices = None
        # 边个数
        self.edge = None
        # 邻接矩阵
        self.adjacency_matrix = None
        # 原矩阵
        self.matrix = matrix

    # 生成临街矩阵，时间复杂度：O(e)，e为边数
    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edge = self.matrix[0][1]
        # 初始化邻接矩阵
        self.adjacency_matrix = [[0 for col in range(self.vertices)] for row in range(self.vertices)]

        for i, j in self.matrix[1:]:
            # 检查顶点合法性
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception('有自环边，非简单图')
            # 检查是否有平行边
            if self.adjacency_matrix[i][j] == 1:
                raise Exception('有平行边，非简单图')

            self.adjacency_matrix[i][j] = 1
            self.adjacency_matrix[j][i] = 1

    # 检查顶点是否合法
    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception('顶点不合法')

    # 获取顶点个数
    def get_vertices(self):
        return self.vertices

    # 获取边个数
    def get_edge(self):
        return self.edge

    # 判断顶点m和n是否相邻，时间复杂度：O(1)
    def has_edge(self, m, n):
        self.validate_vertex(m)
        self.validate_vertex(n)
        return self.adjacency_matrix[m][n] == 1

    # 获取当前顶点相临的所有顶点列表，时间复杂度：O(v)，v为顶点个数
    def adj(self, v):
        self.validate_vertex(v)
        # 顶点列表
        vertices_list = []
        for i in range(self.vertices):
            # 判断是否相邻
            if self.adjacency_matrix[v][i] == 1:
                vertices_list.append(i)
        return vertices_list

    # 获取顶点的度
    def degree(self, v):
        return len(self.adj(v))

    def print_matrix(self):
        for index, item in enumerate(self.adjacency_matrix):
            print(index, ': ', item)


matrix = [[7, 9], [0, 1], [0, 3], [1, 2], [1, 6], [2, 3], [2, 5], [3, 4], [4, 5], [5, 6]]
adjacency_matrix = AdjacencyMatrix(matrix)
adjacency_matrix.generate()
adjacency_matrix.print_matrix()
