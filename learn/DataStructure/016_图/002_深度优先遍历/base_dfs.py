# -*- coding: utf-8 -*-
# @Time:    2021/3/21 12:34
# @Author:  leeyoung
# @File:    base_dfs.py
# @Content: 深度优先遍历


# 深度优先遍历，depth first search
# 空间复杂度：O(v + e)，v为顶点个数，e为边数
class DFS:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否被访问过
        self.visited = list()

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]

    # 如果是联通图，时间复杂度为O(e)，如果不联通，时间复杂度为O(v + e)
    def graph_dfs(self):
        # 遍历节点，为了解决遍历多个联通分量
        for i in range(self.graph.vertices):
            if not self.visited[i]:
                self.dfs(i)

    def dfs(self, v):
        self.visited[v] = True
        # 在这个位置为图深度遍历先序遍历
        print(v)
        # 获取当前顶点所有相邻顶点
        adjacent = self.graph.adj(v)
        # 遍历相邻顶点
        for w in adjacent:
            if not self.visited[w]:
                self.dfs(w)
        # 在这个位置为图深度遍历后序遍历
        # print(v)


def main():
    # 一个联通分量
    # matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
    # 两个联通分量
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

    graph = Graph(matrix)
    graph.generate()

    dfs = DFS(graph)
    dfs.generate()
    # 测试深度优先遍历函数
    dfs.graph_dfs()


if __name__ == '__main__':
    main()
