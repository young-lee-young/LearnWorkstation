# -*- coding: utf-8 -*-
# @Time:    2021/3/20 21:34
# @Author:  leeyoung
# @File:    customize_cycle_detection.py
# @Content: 无向图环检测


class CycleDetection:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否访问过
        self.visited = list()
        # 是否有环标志
        self.has_cycle_flag = False

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        for v in range(self.graph.vertices):
            if not self.visited[v]:
                if self.dfs(v, v):
                    self.has_cycle_flag = True
                    break

    def dfs(self, v, parent):
        self.visited[v] = True
        # 遍历与v相邻顶点
        for w in self.graph.adj(v):
            if not self.visited[w]:
                # 如果w里面有环，直接返回
                if self.dfs(w, v):
                    return True
            # w访问过，并且不是v的父亲节点，证明在w这里存在环
            elif w != parent:
                return True
        return False

    def has_cycle(self):
        return self.has_cycle_flag


def main():
    # 有环
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]
    # 无环
    matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 6]]

    graph = Graph(matrix)
    graph.generate()

    cycle_detection = CycleDetection(graph)
    cycle_detection.generate()
    print("has cycle", cycle_detection.has_cycle())


if __name__ == '__main__':
    main()
