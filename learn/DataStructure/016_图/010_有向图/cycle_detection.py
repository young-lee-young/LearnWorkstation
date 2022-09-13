# -*- coding: utf-8 -*-
# @Time:    2021/5/24 23:15
# @Author:  leeyoung
# @File:    cycle_detection.py
# @Content: 有向图环检测，有向无环图(DAG，Directed Acyclic Graph)


# 有向无权图
class CycleDetection:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否访问过
        self.visited = None
        # 判断是否在路径上
        self.on_path = None
        # 是否有环标志
        self.has_cycle_flag = False

    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.on_path = [False for _ in range(self.graph.vertices)]
        for v in range(self.graph.vertices):
            if not self.visited[v]:
                if self.dfs(v, v):
                    self.has_cycle_flag = True
                    break

    def dfs(self, v, parent):
        self.visited[v] = True
        self.on_path[v] = True
        # 遍历与v相邻顶点
        for w in self.graph.adj(v):
            if not self.visited[w]:
                # 如果w里面有环，直接返回
                if self.dfs(w, v):
                    return True
            # 如果w被访问过，并且在路径上
            elif self.on_path[w]:
                return True
        self.on_path[v] = False
        return False

    def has_cycle(self):
        return self.has_cycle_flag


def main():
    # 无环
    matrix = [[5, 5], [0, 1], [1, 2], [1, 3], [2, 4], [3, 2]]

    # 有环
    matrix = [[5, 5], [0, 1], [1, 2], [1, 3], [2, 4], [3, 2], [3, 0]]

    graph = Graph(matrix)
    graph.generate()

    cycle_detection = CycleDetection(graph)
    cycle_detection.generate()
    print("has cycle", cycle_detection.has_cycle())


if __name__ == '__main__':
    main()
