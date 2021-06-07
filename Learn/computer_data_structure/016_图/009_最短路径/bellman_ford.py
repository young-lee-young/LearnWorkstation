# -*- coding: utf-8 -*-
# @Time:    2021/5/5 23:12
# @Author:  leeyoung
# @File:    bellman_ford.py
# @Content: 贝尔曼-福德算法


import sys


# 暂时只支持无向带权图
class BellmanFord:
    def __init__(self, graph, source):
        self.graph = graph
        self.source = source
        # 源点到各个顶点的最短距离
        self.distance = None
        # 图是否有负权环
        self.negative_cycle = False

    def generate(self):
        # 检查源点是否合法
        self.graph.validate_vertex(self.source)
        # 将距离初始值设置为无穷大
        self.distance = [sys.maxsize for _ in range(self.graph.vertices)]
        # 将源点距离值设置为0
        self.distance[self.source] = 0

        # 对所有边进行 v - 1 次松弛操作，需要循环 v - 1 次
        for i in range(1, self.graph.vertices):
            # 两层for循环用于遍历所有边
            for v in range(self.graph.vertices):
                for w in self.graph.adj(v):
                    # 松弛操作，对w的距离值进行更新
                    if self.distance[v] != sys.maxsize and (self.distance[v] + self.graph.get_weight(v, w) < self.distance[w]):
                        self.distance[w] = self.distance[v] + self.graph.get_weight(v, w)

        # 再进行一次松弛操作，用于检测是否有负权环
        for v in range(self.graph.vertices):
            for w in self.graph.adj(v):
                # 松弛操作，对w的距离值进行更新
                if self.distance[v] != sys.maxsize and (self.distance[v] + self.graph.get_weight(v, w) < self.distance[w]):
                    self.negative_cycle = True

    # 判断是否有负权环
    def has_negative_cycle(self):
        return self.negative_cycle

    def is_connected(self, v):
        self.graph.validate_vertex(v)
        return self.distance[v] != sys.maxsize

    def distance_to(self, v):
        self.graph.validate_vertex(v)
        if self.has_negative_cycle():
            raise Exception('exist negative cycle')
        return self.distance[v]


def main():
    matrix = [[5, 8], [0, 1, 4], [0, 2, 2], [1, 2, 1], [1, 3, 2], [1, 4, 3], [2, 3, 4], [2, 4, 5], [3, 4, 1]]
    matrix = [[5, 8], [0, 1, -1], [0, 2, 2], [1, 2, 1], [1, 3, 2], [1, 4, 3], [2, 3, 4], [2, 4, 5], [3, 4, 1]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    bellman_ford = BellmanFord(weighted_graph, 0)
    bellman_ford.generate()
    print(bellman_ford.distance_to(3))


if __name__ == '__main__':
    main()
