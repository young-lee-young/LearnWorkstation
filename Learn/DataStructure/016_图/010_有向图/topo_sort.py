# -*- coding: utf-8 -*-
# @Time:    2021/6/3 22:53
# @Author:  leeyoung
# @File:    topo_sort.py
# @Content: 拓扑排序

from collections import deque


class TopoSort:
    def __init__(self, graph):
        self.graph = graph
        self.result = list()
        self.has_cycle_flag = False

    def generate(self):
        # 顶点入度列表
        in_degrees = [0 for _ in range(self.graph.vertices)]

        # 初始化队列，队列中的元素均入度为0
        queue = deque()
        for v in range(self.graph.vertices):
            # 初始化入度列表
            in_degrees[v] = self.graph.in_degree(v)
            # 入度为0，放入队列中
            if in_degrees[v] == 0:
                queue.append(v)

        # 起始点
        while len(queue) != 0:
            # 出队
            current = queue.popleft()
            self.result.append(current)

            # 循环节点相邻节点
            for w in self.graph.adj(current):
                # 入度减1
                in_degrees[w] -= 1
                if in_degrees[w] == 0:
                    queue.append(w)

        if len(self.result) != self.graph.vertices:
            self.has_cycle_flag = True
            self.result = []

    def has_cycle(self):
        return self.has_cycle_flag

    def path(self):
        return self.result


def main():
    # 有环
    matrix = [[5, 5], [0, 1], [2, 1], [1, 3], [2, 4], [3, 2]]
    # 无环
    # matrix = [[5, 5], [0, 1], [1, 2], [1, 3], [2, 4], [3, 2]]
    graph = Graph(matrix)
    graph.generate()

    topo_sort = TopoSort(graph)
    topo_sort.generate()
    print(topo_sort.has_cycle())
    print(topo_sort.path())


if __name__ == '__main__':
    main()
