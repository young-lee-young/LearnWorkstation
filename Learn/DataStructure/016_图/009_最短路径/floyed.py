# -*- coding: utf-8 -*-
# @Time:    2021/5/6 14:29
# @Author:  leeyoung
# @File:    floyed.py
# @Content: 弗洛伊德算法


import sys


class Floyed:
    def __init__(self, graph):
        self.graph = graph
        # 点对之间的距离，为二维数组
        self.distance = None
        # 图是否有负权环
        self.negative_cycle = False

    def generate(self):
        # 初始化distance，值为正无穷
        self.distance = [[sys.maxsize for _ in range(self.graph.vertices)] for _ in range(self.graph.vertices)]
        # 修改部分distance值，顶点到自己为0，两个顶点之间有边，值设置为权值
        for v in range(self.graph.vertices):
            self.distance[v][v] = 0
            # 遍历v相邻的顶点
            for w in self.graph.adj(v):
                self.distance[v][w] = self.graph.get_weight(v, w)

        # floyed算法
        for t in range(self.graph.vertices):
            for v in range(self.graph.vertices):
                for w in range(self.graph.vertices):
                    if self.distance[v][t] != sys.maxsize and self.distance[t][w] != sys.maxsize and (self.distance[v][t] + self.distance[t][w] < self.distance[v][w]):
                        self.distance[v][w] = self.distance[v][t] + self.distance[t][w]

        # 检测是否有负权环
        for v in range(self.graph.vertices):
            # 顶点到自己的距离小于0，证明有负权环
            if self.distance[v][v] < 0:
                self.negative_cycle = True

    def has_negative_cycle(self):
        return self.negative_cycle

    def is_connected(self, v, w):
        self.graph.validate_vertex(v)
        self.graph.validate_vertex(w)
        return self.distance[v][w] != sys.maxsize

    def distance_to(self, v, w):
        self.graph.validate_vertex(v)
        self.graph.validate_vertex(w)
        if self.has_negative_cycle():
            raise Exception('exist negative cycle')
        return self.distance[v][w]


def main():
    matrix = [[5, 8], [0, 1, 4], [0, 2, 2], [1, 2, 1], [1, 3, 2], [1, 4, 3], [2, 3, 4], [2, 4, 5], [3, 4, 1]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    floyed = Floyed(weighted_graph)
    floyed.generate()
    print(floyed.distance_to(0, 2))


if __name__ == '__main__':
    main()
