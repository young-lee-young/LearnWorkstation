# -*- coding: utf-8 -*-
# @Time:    2021/3/16 22:38
# @Author:  leeyoung
# @File:    customize_single_source_path.py
# @Content: 单源路径


class SingleSourcePath:
    def __init__(self, graph, source):
        self.graph = graph
        # 源
        self.source = source
        # 遍历时判断是否被访问过
        self.visited = list()
        # 每个顶点前面的顶点
        self.pre = list()

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.pre = [-1 for _ in range(self.graph.vertices)]

        # 对source进行深度优先遍历
        self.dfs(self.source, self.source)

    # 深度优先遍历
    def dfs(self, v, parent):
        # 标记已经访问过
        self.visited[v] = True
        self.pre[v] = parent
        adjacent = self.graph.adj(v)
        for w in adjacent:
            if not self.visited[w]:
                self.dfs(w, v)

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
        # 进行反转
        result.reverse()
        return result


def main():
    # 一个联通分量
    matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
    # 两个联通分量
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

    graph = Graph(matrix)
    graph.generate()

    sspath = SingleSourcePath(graph, 0)
    # 测试生成函数
    sspath.generate()

    print("0 -> 6 : ", sspath.path(6))
    print("0 -> 5 : ", sspath.path(5))


if __name__ == '__main__':
    main()
