# -*- coding: utf-8 -*-
# @Time:    2021/4/20 23:32
# @Author:  leeyoung
# @File:    prim.py
# @Content: 最小生成树 - prim算法

import sys


# 时间复杂度：O((V - 1) * (V + E)) = O(VE)
class Prim:
    def __init__(self, graph):
        self.graph = graph
        self.minimum_spanning_tree = list()

    def generate(self):
        connected_component = ConnectedComponent(self.graph)
        connected_component.generate()
        # 联通分量大于0，没有最小生成树
        if connected_component.connected_component() > 1:
            return

        # 判断顶点是否访问过
        visited = [False for _ in range(self.graph.vertices)]
        visited[0] = True
        # 循环 v - 1 次，为了寻找到 v - 1 条边
        for i in range(1, self.graph.vertices):
            # 当前最下横切边
            min_edge = WeightedEdge(-1, -1, sys.maxsize)
            # 扫描所有边
            for v in range(self.graph.vertices):
                # v访问过，w没有访问过，证明v和w之间的边是横切边
                if visited[v]:
                    for w in self.graph.adj(v):
                        if not visited[w] and self.graph.get_weight(v, w) < min_edge.get_wight():
                            min_edge = WeightedEdge(v, w, self.graph.get_weight(v, w))
            self.minimum_spanning_tree.append(min_edge)
            # 将当前横切边的两个顶点设置为访问过
            visited[min_edge.get_v()] = True
            visited[min_edge.get_w()] = True

    def result(self):
        return self.minimum_spanning_tree


def main():
    matrix = [[7, 12], [0, 1, 2], [0, 3, 7], [0, 5, 2], [1, 2, 1], [1, 3, 4], [1, 4, 3], [1, 5, 5], [2, 4, 4],
              [2, 5, 4], [3, 4, 1], [3, 6, 5], [4, 6, 7]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    prim = Prim(weighted_graph)
    prim.generate()
    print(prim.result())


if __name__ == '__main__':
    main()
