# -*- coding: utf-8 -*-
# @Time:    2021/4/18 22:27
# @Author:  leeyoung
# @File:    kruskal.py
# @Content: 最小生成树 - Kruskal算法


class Kruskal:
    def __init__(self, graph):
        self.graph = graph
        # 最小生成树结果，存放边的列表
        self.minimum_spanning_tree = list()

    def generate(self):
        connected_component = ConnectedComponent(self.graph)
        connected_component.generate()
        # 联通分量大于0，没有最小生成树
        if connected_component.connected_component() > 1:
            return

        # 获取所有的边
        edges = list()
        for v in range(self.graph.vertices):
            for w in self.graph.adj(v):
                # 去除重复的边
                if v < w:
                    edges.append(WeightedEdge(v, w, self.graph.get_weight(v, w)))
        # 将所有的边根据权重从小到大排序
        edges.sort(key=lambda x: x.weight)

        # 使用并查集判断是否有环
        union_find = CustomizeUnionFind(self.graph.vertices)
        union_find.constructor()
        # 从小到大遍历所有的边
        for edge in edges:
            v = edge.get_v()
            w = edge.get_w()
            # 如果两个顶点不属于一个集合，他们之间不相连，此时边就是最小生成树的一部分，并且将两个顶点加入到一个集合中
            if not union_find.is_connected(v, w):
                self.minimum_spanning_tree.append(edge)
                union_find.union(v, w)

    # 返回最小生成树的结果
    def result(self):
        return self.minimum_spanning_tree


def main():
    matrix = [[7, 12], [0, 1, 2], [0, 3, 7], [0, 5, 2], [1, 2, 1], [1, 3, 4], [1, 4, 3], [1, 5, 5], [2, 4, 4],
              [2, 5, 4], [3, 4, 1], [3, 6, 5], [4, 6, 7]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    kruskal = Kruskal(weighted_graph)
    kruskal.generate()
    print(kruskal.result())


if __name__ == '__main__':
    main()
