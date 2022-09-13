# -*- coding: utf-8 -*-
# @Time:    2021/4/14 21:51
# @Author:  leeyoung
# @File:    hamilton_loop_v2.py
# @Content:


class HamiltonLoopV2:
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

        adjacent = self.graph.adj(v)
        for w in adjacent:
            if not self.visited[w]:
                if self.dfs(w, v, left):
                    return True
            # w 是起点并且所有的顶点都访问过，证明存在哈密尔顿回路
            elif w == 0 and left == 0:
                self.end = v
                return True
        # 进行回溯，表示从v不能回到起点
        self.visited[v] = False
        # 回溯将left还原，这里其实是不用写的，因为递归返回上一层，这一层的left是没有用到的
        left += 1
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

    hamilton = HamiltonLoopV2(graph)
    hamilton.generate()
    print(hamilton.path())


if __name__ == '__main__':
    main()
