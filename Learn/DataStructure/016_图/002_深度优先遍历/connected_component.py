# -*- coding: utf-8 -*-
# @Time:    2021/4/20 21:40
# @Author:  leeyoung
# @File:    connected_component.py
# @Content: 联通分量


class ConnectedComponent:
    def __init__(self, graph):
        # 无向带权图
        self.graph = graph
        # 遍历时判断是否被访问过
        self.visited = list()
        # 联通分量个数
        self.connected_component_count = 0

    def generate(self):
        self.visited = [-1 for _ in range(self.graph.vertices)]

        # 遍历节点，为了解决遍历多个联通分量
        for i in range(self.graph.vertices):
            if self.visited[i] == -1:
                self.dfs(i, self.connected_component_count)
                self.connected_component_count += 1

    def dfs(self, v, connected_component_id):
        self.visited[v] = connected_component_id
        # 获取当前顶点所有相邻顶点
        adjacent = self.graph.adj(v)
        # 遍历相邻顶点
        for w in adjacent:
            # 判断是否已经遍历过
            if self.visited[w] == -1:
                self.dfs(w, connected_component_id)

    # 求图的联通分量
    def connected_component(self):
        return self.connected_component_count

    # 求联通分量里包含的数据
    def components(self):
        result = [[] for _ in range(self.connected_component_count)]
        for v in range(self.graph.vertices):
            result[self.visited[v]].append(v)
        return result

    # 判断两个顶点是否在同一个联通分量中
    def is_connected(self, v, w):
        return self.visited[v] == self.visited[w]


def main():
    # 一个联通分量
    matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
    # 两个联通分量
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

    graph = Graph(matrix)
    graph.generate()

    connected_component = ConnectedComponent(graph)
    connected_component.generate()

    # 测试联通分量函数
    print(connected_component.connected_component())

    # 测试每个联通分量里包含元素函数
    print(connected_component.components())

    # 测试联通函数
    print(connected_component.is_connected(1, 2))
    print(connected_component.is_connected(1, 5))


if __name__ == '__main__':
    main()
