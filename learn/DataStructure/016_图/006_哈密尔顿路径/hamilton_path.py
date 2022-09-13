# -*- coding: utf-8 -*-
# @Time:    2021/4/14 22:51
# @Author:  leeyoung
# @File:    hamilton_path.py
# @Content: 哈密尔顿路径


class HamiltonPath:
    def __init__(self, graph):
        self.graph = graph
        # 保存具体回路，保存当前顶点前一个顶点
        self.pre = list()
        # 哈密尔顿路径最后一个顶点
        self.end = -1
        # 起点
        self.s = None

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self, s):
        self.s = s
        self.pre = [-1 for _ in range(self.graph.vertices)]

        # 遍历时判断是否被访问过
        visited = 0

        self.dfs(visited, s, s, self.graph.vertices)

    def dfs(self, visited, v, parent, left):
        visited += (1 << v)
        self.pre[v] = parent
        left -= 1

        # 顶点全部访问
        if left == 0:
            self.end = v
            return True

        # 遍历相邻顶点
        adjacent = self.graph.adj(v)
        for w in adjacent:
            if (visited & (1 << w)) == 0:
                if self.dfs(visited, w, v, left):
                    return True

        # 进行回溯，表示从v不能回到起点
        visited -= (1 << v)
        return False

    # 哈密尔顿路径
    def path(self):
        result = list()
        # 没有哈密尔顿路径
        if self.end == -1:
            return result
        # 循环pre中的数组
        current = self.end
        while current != self.s:
            result.append(current)
            current = self.pre[current]
        result.append(self.s)
        result.reverse()
        return result


def main():
    matrix = [[4, 5], [0, 1], [0, 2], [0, 3], [1, 2], [1, 3]]
    graph = Graph(matrix)
    graph.generate()

    hamilton = HamiltonPath(graph)
    # hamilton.generate(0)
    hamilton.generate(1)
    print(hamilton.path())


if __name__ == '__main__':
    main()
