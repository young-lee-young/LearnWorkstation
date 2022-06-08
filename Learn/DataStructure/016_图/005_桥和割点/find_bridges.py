# -*- coding: utf-8 -*-
# @Time:    2021/4/1 22:10
# @Author:  leeyoung
# @File:    find_bridges.py
# @Content: 寻找桥


# 使用DFS寻找桥，不能使用BFS
class FindBridges:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否被访问过
        self.visited = list()
        # order[v]表示顶点v在DFS中的访问顺序
        self.order = list()
        # low[v]表示DFS中，顶点v能达到的最小order值
        self.low = list()
        # 遍历个数
        self.count = 0
        # 保存找到的桥
        self.bridges = list()

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.order = [-1 for _ in range(self.graph.vertices)]
        self.low = [-1 for _ in range(self.graph.vertices)]

        # 遍历节点，解决多联通分量
        for i in range(self.graph.vertices):
            if not self.visited[i]:
                self.dfs(i, i)

    def dfs(self, v, parent):
        self.visited[v] = True
        self.order[v] = self.count
        self.low[v] = self.order[v]
        self.count += 1

        # 获取当前顶点所有相邻节点
        adjacent = self.graph.adj(v)
        # 遍历相邻顶点，其实是遍历边 v - w
        for w in adjacent:
            if not self.visited[w]:
                self.dfs(w, v)
                # 更新当前节点的low值
                self.low[v] = min(self.low[v], self.low[w])
                # v的下一个顶点w可以到达的最小值比v要大，说明w没有从另一条路回到v，所以 v - w 是桥
                if self.low[w] > self.order[v]:
                    self.bridges.append((v, w))
            elif w != parent:
                self.low[v] = min(self.low[v], self.low[w])

    def find_bridges(self):
        return self.bridges


def main():
    matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [2, 3], [3, 5], [4, 5], [4, 6], [5, 6]]

    graph = Graph(matrix)
    graph.generate()

    find_bridges = FindBridges(graph)
    find_bridges.generate()
    print(find_bridges.find_bridges())


if __name__ == '__main__':
    main()
