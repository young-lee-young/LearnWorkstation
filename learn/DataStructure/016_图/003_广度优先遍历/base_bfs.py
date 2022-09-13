# -*- coding: utf-8 -*-
# @Time:    2021/3/21 12:36
# @Author:  leeyoung
# @File:    base_bfs.py
# @Content: 广度优先遍历


from collections import deque


# 深度优先遍历
class BFS:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否被访问过
        self.visited = list()
        # 遍历时的队列
        self.queue = deque()

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]

    # 如果是联通图，时间复杂度为O(e)，如果不联通，时间复杂度为O(v + e)
    def graph_bfs(self):
        for v in range(self.graph.vertices):
            if not self.visited[v]:
                self.bfs(v)

    def bfs(self, s):
        self.visited[s] = True
        self.queue.append(s)

        while len(self.queue) != 0:
            v = self.queue.popleft()
            print(v)
            # 获取当前顶点所有相邻顶点
            adjacent = self.graph.adj(v)
            for w in adjacent:
                # 如果w没有访问过
                if not self.visited[w]:
                    self.visited[w] = True
                    self.queue.append(w)


def main():
    matrix = [[7, 7], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [5, 6]]

    graph = Graph(matrix)
    graph.generate()

    bfs = BFS(graph)
    bfs.generate()

    # 测试广度优先遍历函数
    bfs.graph_bfs()


if __name__ == '__main__':
    main()
