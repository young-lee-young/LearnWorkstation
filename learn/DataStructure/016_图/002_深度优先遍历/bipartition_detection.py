# -*- coding: utf-8 -*-
# @Time:    2021/3/20 22:49
# @Author:  leeyoung
# @File:    customize_bipartition_detection.py
# @Content: 二分图检测


class BipartitionDetection:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否访问过
        self.visited = list()
        # 各个顶点的颜色
        self.colors = list()
        # 是否二分图
        self.is_bipartite_flag = True

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.colors = [-1 for _ in range(self.graph.vertices)]

        # 检查所有联通分量是否时二分图
        for v in range(self.graph.vertices):
            if not self.visited[v]:
                if not self.dfs(v, 0):
                    self.is_bipartite_flag = False
                    # 不是二分图直接退出
                    break

    def dfs(self, v, color):
        self.visited[v] = True
        self.colors[v] = color
        adjacent = self.graph.adj(v)
        for w in adjacent:
            if not self.visited[w]:
                # 深度遍历不是二分图
                if not self.dfs(w, 1 - color):
                    return False
            # w已经访问过，w已经被染过颜色
            if self.colors[w] == self.colors[v]:
                return False
        return True

    def is_bipartite(self):
        return self.is_bipartite_flag


def main():
    # 是二分图
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]
    # 不是二分图
    matrix = [[4, 6], [0, 1], [0, 2], [0, 3], [1, 2], [1, 3], [2, 3]]

    graph = Graph(matrix)
    graph.generate()

    bipartition_detection = BipartitionDetection(graph)
    bipartition_detection.generate()
    print("is bipartite", bipartition_detection.is_bipartite())


if __name__ == '__main__':
    main()
