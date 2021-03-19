# -*- coding: utf-8 -*-
# @Time:    2021/3/3 23:01
# @Author:  leeyoung
# @File:    customize_graph_base_adjacency_list.py
# @Content: 无向图 - 基于临接表


# 空间复杂度：O(v + e)，v为顶点个数，e为边数
class AdjacencyList:
    def __init__(self, matrix):
        # 顶点个数
        self.vertices = None
        # 边个数
        self.edge = None
        # 临接表，用链表实现，优化的话使用红黑树
        self.adjacency_list = None
        # 原矩阵
        self.matrix = matrix
        # 遍历时判断是否被访问过
        self.visited = list()
        # 联通分量个数
        self.connected_component_count = 0

    # 生成临接表，时间复杂度：O(e * v)，主要耗费在判断是否存在平行边
    def generate(self):
        self.vertices = self.matrix[0][0]
        self.edge = self.matrix[0][1]
        self.visited = [-1 for row in range(self.vertices)]

        # 初始化临接表
        self.adjacency_list = [[] for row in range(self.vertices)]
        for i, j in self.matrix[1:]:
            # 检查顶点合法性
            self.validate_vertex(i)
            self.validate_vertex(j)

            # 检查是否有自环边
            if i == j:
                raise Exception('有自环边，非简单图')
            # 检查是否有平行边
            if i in self.adjacency_list[j]:
                raise Exception('有平行边，非简单图')

            self.adjacency_list[i].append(j)
            self.adjacency_list[j].append(i)

    # 检查顶点是否合法
    def validate_vertex(self, v):
        if v < 0 or v >= self.vertices:
            raise Exception('顶点不合法')

    # 获取顶点个数
    def get_vertices(self):
        return self.vertices

    # 获取边个数
    def get_edge(self):
        return self.edge

    # 判读顶点m和n是否相邻，时间复杂度：O(degree(v))
    def has_edge(self, m, n):
        self.validate_vertex(m)
        self.validate_vertex(n)
        return n in self.adjacency_list[m]

    # 获取当前顶点相邻的所有顶点列表，时间复杂度：O(degree(v))
    def adj(self, v):
        self.validate_vertex(v)
        return self.adjacency_list[v]

    # 获取顶点的度
    def degree(self, v):
        return len(self.adj(v))

    # 求图的联通分量
    def connected_component(self):
        # 返回联通分量
        if self.connected_component_count == 0:
            self.graph_dfs()

        # 查看联通分量的分组情况
        print(self.visited)
        return self.connected_component_count

    # 每个联通分量里包含的数据
    def components(self):
        result = [[] for i in range(self.connected_component_count)]
        for v in range(self.vertices):
            result[self.visited[v]].append(v)
        return result

    # 深度优先遍历，depth first search
    # 如果是联通图，时间复杂度为O(e)，如果不联通，时间复杂度为O(v + e)
    def graph_dfs(self):
        # 遍历节点，为了解决遍历多个联通分量
        for i in range(self.vertices):
            # 等于-1表示没有被遍历过
            if self.visited[i] == -1:
                self.dfs(i, self.connected_component_count)
                self.connected_component_count += 1

    def dfs(self, v, connected_component_id):
        self.visited[v] = connected_component_id
        # 在这个位置为图深度遍历先序遍历
        print(v)
        # 获取当前顶点所有相邻顶点
        adjacent = self.adj(v)
        # 遍历相邻顶点
        for w in adjacent:
            # 判断是否已经遍历过
            if self.visited[w] == -1:
                self.dfs(w, connected_component_id)
        # 在这个位置为图深度遍历后序遍历

    # 判断两个顶点是否在同一个联通分量中
    def is_connected(self, v, w):
        self.validate_vertex(v)
        self.validate_vertex(w)
        return self.visited[v] == self.visited[w]

    def print_matrix(self):
        for index, item in enumerate(self.adjacency_list):
            print(index, ': ', item)


# 一个联通分量
matrix = [[7, 8], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6], [3, 5], [5, 6]]
# 两个联通分量
matrix = [[7, 6], [0, 1], [0, 2], [1, 3], [1, 4], [2, 3], [2, 6]]

adjacency_list = AdjacencyList(matrix)
# 测试生成函数
adjacency_list.generate()
adjacency_list.print_matrix()

# 测试深度优先遍历函数
adjacency_list.graph_dfs()

# 测试联通分量函数
print(adjacency_list.connected_component())

# 测试每个联通分量里包含元素函数
print(adjacency_list.components())

# 测试联通函数
print(adjacency_list.is_connected(1, 2))
print(adjacency_list.is_connected(1, 5))
