# -*- coding: utf-8 -*-
# @Time:    2021/5/3 23:55
# @Author:  leeyoung
# @File:    dijkstra_optimize.py
# @Content: 狄杰斯特拉算法优化


import sys
from queue import PriorityQueue


# 时间复杂度：O(ElogE)
class DijkstraOptimize:
    def __init__(self, graph, source):
        self.graph = graph
        self.source = source
        # 源点到各个顶点的最短距离
        self.distance = None
        # 判断顶点是否访问过
        self.visited = None

    class Node:
        def __init__(self, v, dis):
            self.v = v
            self.dis = dis

        def __lt__(self, other):
            return self.dis if self.dis < other.dis else other.dis

    def generate(self):
        # 检查源点是否合法
        self.graph.validate_vertex(self.source)
        # 将初始值设置为无穷大
        self.distance = [sys.maxsize for _ in range(self.graph.vertices)]
        # 将源点距离值设置为0
        self.distance[self.source] = 0
        self.visited = [False for _ in range(self.graph.vertices)]

        priority_queue = PriorityQueue()
        # 将源点加入到优先队列
        priority_queue.put(self.Node(self.source, 0))

        while not priority_queue.empty():
            # 1
            # 获取没有访问过的顶点中，距离源点最近的顶点
            current = priority_queue.get().v

            # 由于队列中会添加多次顶点值相同、距离值不同的顶点，如果访问过，就继续循环
            if self.visited[current]:
                continue

            # 2
            # 当前顶点已经确定了最短路径
            self.visited[current] = True

            # 3
            # 遍历当前顶点的所有相邻节点
            for w in self.graph.adj(current):
                if not self.visited[w]:
                    if self.distance[current] + self.graph.get_weight(current, w) < self.distance[w]:
                        # 这里更新distance值，但是已经加入到优先队列中的值并没有改变
                        self.distance[w] = self.distance[current] + self.graph.get_weight(current, w)
                        # 所以这里需要更新有点队列中的值，方法就是再添加一份，并且根据距离小的排序，但是这样就有重复的顶点值，所以上面判断访问过就continue
                        priority_queue.put(self.Node(w, self.distance[w]))

    def is_connected(self, v):
        self.graph.validate_vertex(v)
        return self.visited[v]

    def distance_to(self, v):
        self.graph.validate_vertex(v)
        return self.distance[v]


def main():
    matrix = [[5, 8], [0, 1, 4], [0, 2, 2], [1, 2, 1], [1, 3, 2], [1, 4, 3], [2, 3, 4], [2, 4, 5], [3, 4, 1]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    dijkstra_optimize = DijkstraOptimize(weighted_graph, 0)
    dijkstra_optimize.generate()
    print(dijkstra_optimize.distance_to(3))


if __name__ == '__main__':
    main()
