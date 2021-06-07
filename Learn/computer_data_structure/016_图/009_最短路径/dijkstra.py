# -*- coding: utf-8 -*-
# @Time:    2021/5/3 20:16
# @Author:  leeyoung
# @File:    dijkstra.py
# @Content: 狄杰斯特拉算法


import sys


class Dijkstra:
    def __init__(self, graph, source):
        self.graph = graph
        self.source = source
        # 源点到各个顶点的最短距离
        self.distance = None
        # 判断顶点是否访问过
        self.visited = None

    def generate(self):
        # 检查源点是否合法
        self.graph.validate_vertex(self.source)
        # 将距离初始值设置为无穷大
        self.distance = [sys.maxsize for _ in range(self.graph.vertices)]
        # 将源点距离值设置为0
        self.distance[self.source] = 0
        self.visited = [False for _ in range(self.graph.vertices)]

        # djikstra算法
        while True:
            # 1
            # 在所有没有被访问过的顶点中寻找距离源点最近的顶点
            current_distance = sys.maxsize
            current = -1
            for i in range(self.graph.vertices):
                if not self.visited[i] and self.distance[i] < current_distance:
                    current_distance = self.distance[i]
                    current = i

            # 所有顶点都遍历过，退出循环
            if current == -1:
                break

            # 2
            # 当前点已经确定了最短路径
            self.visited[current] = True

            # 3
            # 遍历当前顶点相邻顶点
            for w in self.graph.adj(current):
                if not self.visited[w]:
                    # 更新当前顶点相邻顶点的距离
                    if self.distance[current] + self.graph.get_weight(current, w) < self.distance[w]:
                        self.distance[w] = self.distance[current] + self.graph.get_weight(current, w)

    def is_connected(self, v):
        self.graph.validate_vertex(v)
        return self.visited[v]

    def distance_to(self, v):
        self.graph.validate_vertex(v)
        return self.distance[v]


def main():
    matrix = [[5, 8], [0, 1, 4], [0, 2, 2], [1, 2, 1], [1, 3, 2], [1, 4, 3], [2, 3, 4], [2, 4, 5], [3, 4, 1]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    dijkstra = Dijkstra(weighted_graph, 0)
    dijkstra.generate()
    print(dijkstra.distance_to(1))


if __name__ == '__main__':
    main()
