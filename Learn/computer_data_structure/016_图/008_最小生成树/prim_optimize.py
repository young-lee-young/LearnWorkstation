# -*- coding: utf-8 -*-
# @Time:    2021/4/24 22:18
# @Author:  leeyoung
# @File:    prim_optimize.py
# @Content:


from queue import PriorityQueue


# 时间复杂度：O(ElogE)
class PrimOptimize:
    def __init__(self, graph):
        self.graph = graph
        self.minimum_spanninng_tree = list()

    def generate(self):
        connected_component = ConnectedComponent(self.graph)
        connected_component.generate()
        # 联通分量大于0，没有最小生成树
        if connected_component.connected_component() > 1:
            return

        # 判断顶点是否访问过
        visited = [False for _ in range(self.graph.vertices)]
        visited[0] = True

        priority_queue = PriorityQueue()
        for w in self.graph.adj(0):
            priority_queue.put(WeightedEdge(0, w, self.graph.get_weight(0, w)))

        while not priority_queue.empty():
            min_edge = priority_queue.get()
            # 如果两个顶点均访问过，这条边不是横切边
            if visited[min_edge.get_v()] and visited[min_edge.get_w()]:
                continue
            self.minimum_spanninng_tree.append(min_edge)

            # 将新点加入到切分中
            new_v = min_edge.get_w() if visited[min_edge.get_v()] else min_edge.get_v()
            visited[new_v] = True
            for w in self.graph.adj(new_v):
                if not visited[w]:
                    priority_queue.put(WeightedEdge(new_v, w, self.graph.get_weight(new_v, w)))

    def result(self):
        return self.minimum_spanninng_tree


def main():
    matrix = [[7, 12], [0, 1, 2], [0, 3, 7], [0, 5, 2], [1, 2, 1], [1, 3, 4], [1, 4, 3], [1, 5, 5], [2, 4, 4],
              [2, 5, 4], [3, 4, 1], [3, 6, 5], [4, 6, 7]]
    weighted_graph = WeightedGraph(matrix)
    weighted_graph.generate()

    prim_optimize = PrimOptimize(weighted_graph)
    prim_optimize.generate()
    print(prim_optimize.result())


if __name__ == '__main__':
    main()
