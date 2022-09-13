# -*- coding: utf-8 -*-
# @Time:    2021/5/23 18:36
# @Author:  leeyoung
# @File:    graph.py
# @Content: 有向无权图


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
        # 入度
        self.in_degrees = None
        # 出度
        self.out_degrees = None

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edges = self.matrix[0][1]

        # 初始化临接表
        self.adjacency_list = [[] for _ in range(self.vertices)]
        self.in_degrees = [0 for _ in range(self.vertices)]
        self.out_degrees = [0 for _ in range(self.vertices)]

        for i, j in self.matrix[1:]:
            # 检查顶点合法性
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception('有自环边，非简单图')

            # 记录每个顶点的相邻顶点
            self.adjacency_list[i].append(j)

            # i出度+1，j入度+1
            self.out_degrees[i] += 1
            self.in_degrees[j] += 1

    def get_vertices(self):
        return self.vertices

    def get_edges(self):
        return self.edges

    # 判断v和w是否相邻，时间复杂度O(degree(v))
    def has_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)
        return v in self.adjacency_list[w]

    # 获取当前顶点的相邻顶点列表，时间复杂度O(degree(v))
    def adj(self, v):
        self.validate_vertex(v)
        return self.adjacency_list[v]

    # 获取顶点v的入度
    def in_degree(self, v):
        return self.in_degrees[v]

    # 获取顶点出度
    def out_degree(self, v):
        return self.out_degrees[v]

    # 从图中删除一条边
    def remove_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)

        self.adjacency_list[v].remove(w)
        self.out_degrees[v] -= 1
        self.in_degrees[w] -= 1
        self.edges -= 1

    def print_matrix(self):
        for index, item in enumerate(self.adjacency_list):
            print(index, ': ', item)

    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception('顶点不合法')


def main():
    matrix = [[5, 5], [0, 1], [1, 2], [1, 3], [2, 4], [3, 2]]

    graph = Graph(matrix)
    # 测试生成函数
    graph.generate()
    graph.print_matrix()

    print(graph.in_degree(2))
    print(graph.out_degree(1))


if __name__ == '__main__':
    main()
