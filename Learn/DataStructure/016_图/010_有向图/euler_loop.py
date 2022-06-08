# -*- coding: utf-8 -*-
# @Time:    2021/5/29 10:38
# @Author:  leeyoung
# @File:    euler_loop.py
# @Content: 有向图欧拉路径


from collections import deque


# 使用双端队列实现栈
class CustomizeStack:
    def __init__(self):
        self.stack = deque()

    def push(self, value):
        return self.stack.append(value)

    def pop(self):
        return self.stack.pop()

    def top(self):
        return self.stack[-1]

    def size(self):
        return len(self.stack)

    def is_empty(self):
        if self.stack:
            return False
        return True


class EulerLoop:
    def __init__(self, graph):
        self.graph = graph

    def has_euler_loop(self):
        # 有向图联通性问题后续介绍

        # 如果任意顶点入度不等于出度，没有欧拉回路
        for v in range(self.graph.vertices):
            if self.graph.in_degree(v) != self.graph.out_degree(v):
                return False
        return True

    def path(self):
        # 相当于loop栈
        result = list()

        # 不存在欧拉回路
        if not self.has_euler_loop():
            return result

        # 新建current_path栈
        stack = CustomizeStack()
        # 从0开始，将起始点入栈
        current_v = 0
        stack.push(current_v)
        while not stack.is_empty():
            # 从current_v有路可走
            if self.graph.out_degree(current_v) != 0:
                stack.push(current_v)
                # 随便取一个相邻顶点
                w = self.graph.adj(current_v)[0]
                # 将两个顶点之间的边删除
                self.graph.remove_edge(current_v, w)
                current_v = w
            # 无路可走
            else:
                # 将当前节点加入到栈
                result.append(current_v)
                # 并且回退一步
                current_v = stack.pop()

        result.reverse()
        return result


def main():
    # 没有欧拉回路
    matrix = [[5, 5], [0, 1], [1, 2], [1, 3], [2, 4], [3, 2]]
    # 有欧拉回路
    matrix = [[5, 8], [0, 1], [1, 2], [1, 3], [2, 0], [2, 4], [3, 1], [3, 2], [4, 3]]

    graph = Graph(matrix)
    graph.generate()

    euler_loop = EulerLoop(graph)
    print(euler_loop.has_euler_loop())
    print(euler_loop.path())


if __name__ == '__main__':
    main()
