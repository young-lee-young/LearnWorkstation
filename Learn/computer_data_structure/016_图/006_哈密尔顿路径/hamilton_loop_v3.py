# -*- coding: utf-8 -*-
# @Time:    2021/4/5 17:55
# @Author:  leeyoung
# @File:    hamilton_loop.py
# @Content: 哈密尔顿回路优化


class HamiltonLoopV3:
    def __init__(self, graph):
        self.graph = graph
        # 遍历时判断是否被访问过
        self.visited = list()
        # 保存具体回路，保存当前顶点前一个顶点
        self.pre = list()
        # 哈密尔顿回路最后一个顶点
        self.end = None

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.visited = [False for _ in range(self.graph.vertices)]
        self.pre = [-1 for _ in range(self.graph.vertices)]
        self.end = -1

        self.dfs(0, 0, self.graph.vertices)

    def dfs(self, v, parent, left):
        self.visited[v] = True
        self.pre[v] = parent
        left -= 1

        # 全部顶点都已经访问，并且v和0之间相连
        if left == 0 and self.graph.has_edge(v, 0):
            self.end = v
            return True

        adjacent = self.graph.adj(v)
        for w in adjacent:
            if not self.visited[w]:
                if self.dfs(w, v, left):
                    return True

        # 进行回溯，表示从v不能回到起点
        self.visited[v] = False
        return False

    # 哈密尔顿回路路径
    def path(self):
        result = list()
        # 没有哈密尔顿回路
        if self.end == -1:
            return result
        # 循环pre中的数组
        current = self.end
        while current != 0:
            result.append(current)
            current = self.pre[current]
        result.append(0)
        result.reverse()
        return result


def main():
    matrix = [[4, 5], [0, 1], [0, 2], [0, 3], [1, 2], [1, 3]]
    graph = Graph(matrix)
    graph.generate()

    hamilton = HamiltonLoopV3(graph)
    hamilton.generate()
    print(hamilton.path())


if __name__ == '__main__':
    main()
