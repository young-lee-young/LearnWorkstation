# -*- coding: utf-8 -*-
# @Time:    2021/5/3 20:17
# @Author:  leeyoung
# @File:    weighted_graph.py
# @Content: 带权图


class WeightedGraph:
    def __init__(self, matrix):
        # 顶点个数
        self.vertices = None
        # 边个数
        self.edges = None
        # 临接表，使用map实现
        self.adjacency_dict = None
        # 原矩阵
        self.matrix = matrix

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edges = self.matrix[0][1]

        # 初始化临接表
        self.adjacency_dict = {i: {} for i in range(self.vertices)}
        # i, j是顶点，k是权重
        for i, j, k in self.matrix[1:]:
            # 检查顶点合法性
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception('有自环边，非简单图')
            # 检查是否有平行边
            if i in self.adjacency_dict[j].keys():
                raise Exception('有平行边，非简单图')

            # 记录每个顶点相邻的顶点
            self.adjacency_dict[i][j] = k
            self.adjacency_dict[j][i] = k

    # 获取顶点个数
    def get_vertices(self):
        return self.vertices

    # 获取边个数
    def get_edges(self):
        return self.edges

    # 判读顶点m和n是否相邻，时间复杂度：O(degree(v))
    def has_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)
        return v in self.adjacency_dict[w].keys()

    # 获取顶点的度
    def degree(self, v):
        return len(self.adj(v))

    # 获取当前顶点相邻的所有顶点列表，时间复杂度：O(degree(v))
    def adj(self, v):
        self.validate_vertex(v)
        return self.adjacency_dict[v].keys()

    # 获取两个边之间的权重
    def get_weight(self, v, w):
        if not self.has_edge(v, w):
            raise Exception('v 和 w 之间没有边')
        return self.adjacency_dict[v].get(w)

    # 从图中删除一条边
    def remove_edge(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)
        del (self.adjacency_dict[v][w])
        del (self.adjacency_dict[w][v])
        self.edges -= 1

    def print_matrix(self):
        for key, value in self.adjacency_dict.items():
            print(key, ': ', value)

    # 检查顶点是否合法
    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception('顶点不合法')


def main():
    matrix = [[7, 12], [0, 1, 2], [0, 3, 7], [0, 5, 2], [1, 2, 1], [1, 3, 4], [1, 4, 3], [1, 5, 5], [2, 4, 4], [2, 5, 4], [3, 4, 1], [3, 6, 5], [4, 6, 7]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()
    weighted_graph.print_matrix()


if __name__ == '__main__':
    main()
