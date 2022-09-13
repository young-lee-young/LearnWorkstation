# -*- coding: utf-8 -*-
# @Time:    2021/4/15 23:00
# @Author:  leeyoung
# @File:    euler_loop.py
# @Content: 欧拉路径 - Hierholzer算法

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

    # 判读存在欧拉路径
    def has_euler_loop(self):
        # 联通分量大于1，不存在欧拉回路
        connected_component = ConnectedComponent(self.graph)
        connected_component.generate()
        connected_component_count = connected_component.connected_component()
        if connected_component_count > 1:
            return False

        for v in range(self.graph.vertices):
            # 如果一个点相邻顶点是奇数，不存在欧拉回路
            if self.graph.degree(v) % 2 == 1:
                return False
        return True

    # 欧拉回路
    def path(self):
        # 相当于loop栈
        result = list()

        # 不存在欧拉路径
        if not self.has_euler_loop():
            return result

        # 新建current_path栈
        stack = CustomizeStack()
        # 从0开始，将起始点入栈
        current_v = 0
        stack.push(current_v)
        while not stack.is_empty():
            # 从current_v有路可走
            if self.graph.degree(current_v) != 0:
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

        return result


def main():
    matrix = [[5, 6], [0, 1], [0, 2], [1, 2], [2, 3], [2, 4], [3, 4]]
    matrix = [[11, 15], [0, 1], [0, 3], [1, 2], [1, 4], [1, 5], [2, 5], [3, 4], [4, 5], [4, 6], [5, 7], [6, 7], [7, 8], [7, 9], [8, 10], [9, 10]]
    graph = Graph(matrix)
    graph.generate()

    euler_loop = EulerLoop(graph)
    print(euler_loop.has_euler_loop())
    print(euler_loop.path())


if __name__ == '__main__':
    main()
