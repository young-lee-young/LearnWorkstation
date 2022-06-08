# -*- coding: utf-8 -*-
# @Time:    2021/3/21 13:42
# @Author:  leeyoung
# @File:    single_source_path.py
# @Content: 单源路径


from collections import deque


class SingleSourcePath:
    def __init__(self, graph, source):
        self.graph = graph
        # 源
        self.source = source
        # 遍历时判断是否被访问过
        self.visited = list()
        # 遍历时的队列
        self.queue = deque()
        # 每个顶点前面的顶点
        self.pre = list()
        # 从源点到各个顶点的距离
        self.dis = list()

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.pre = [-1 for _ in range(self.graph.vertices)]
        self.dis = [-1 for _ in range(self.graph.vertices)]

        # 对source进行广度优先遍历
        self.bfs(self.source, self.source)

    # 深度优先遍历
    def bfs(self, s, parent):
        self.visited[s] = True
        self.queue.append(s)
        self.pre[s] = parent
        # 源点到源点的距离为0
        self.dis[s] = 0

        while len(self.queue) != 0:
            v = self.queue.popleft()
            adjacent = self.graph.adj(v)
            for w in adjacent:
                if not self.visited[w]:
                    self.visited[w] = True
                    self.queue.append(w)
                    self.pre[w] = v
                    self.dis[w] = self.dis[v] + 1

    # 判断是否和源点联通
    def is_connected(self, v):
        return self.visited[v]

    # 源点到v的路径
    def path(self, v):
        result = list()
        # 判断是否和源点联通，不联通直接返回
        if not self.is_connected(v):
            return result

        current = v
        while current != self.source:
            result.append(current)
            # 获取当前点的上一个点
            current = self.pre[current]
        result.append(self.source)
        result.reverse()
        return result

    def distance(self, v):
        return self.dis[v]


def main():
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

    graph = Graph(matrix)
    graph.generate()

    sspath = SingleSourcePath(graph, 0)
    sspath.generate()

    print("0 -> 6 : ", sspath.path(6))
    print("0 -> 6 distance ", sspath.distance(6))


if __name__ == '__main__':
    main()
