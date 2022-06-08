# -*- coding: utf-8 -*-
# @Time:    2021/1/11 下午10:58
# @Author:  leeyoung
# @File:    customize_map.py
# @ontent: 自定义映射 - 基于二分搜索树实现

from customize_map_node import CustomizeMapTreeNode


class CustomizeMap:
    def __init__(self):
        self.root = None
        self.map_size = 0

    def size(self):
        return self.map_size

    def is_empty(self):
        return self.map_size == 0

    # O(logn)
    def add(self, key, value):
        # 没有根节点，新建根节点
        if not self.root:
            self.root = CustomizeMapTreeNode(key, value)
            self.map_size += 1
        # 添加节点
        else:
            self._add_node(self.root, key, value)

    def _add_node(self, node, key, value):
        # 如果存在，更新节点值
        if node.key == key:
            node.value = value
            return
        # 直接添加到左子树节点
        if key < node.key and not node.left:
            node.left = CustomizeMapTreeNode(key, value)
            self.map_size += 1
            return
        # 直接添加到右子树节点
        if key > node.key and not node.right:
            node.right = CustomizeMapTreeNode(key, value)
            self.map_size += 1
            return

        # 往左子树递归添加
        if key < node.key:
            self._add_node(node.left, key, value)
        # 右子树递归添加
        else:
            self._add_node(node.right, key, value)

    def _get_node(self, node, key):
        if node is None:
            return None
        # 左子树里查找
        if key < node.key:
            return self._get_node(node.left, key)
        # 右子树里查找
        elif key > node.key:
            return self._get_node(node.right, key)
        # 等于直接返回
        else:
            return node

    # O(logn)
    def contains(self, key):
        # 查找是否存在
        return self._get_node(self.root, key) is not None

    # 获取值，O(logn)
    def get(self, key):
        node = self._get_node(self.root, key)
        if node is None:
            return None
        return node.value

    # 设置值，O(logn)
    def set(self, key, value):
        node = self._get_node(self.root, key)
        if node is None:
            return None
        node.value = value

    # 获取最小元素
    def minimum(self):
        if self.map_size == 0:
            return None
        return self._minimum(self.root)

    # 最小元素在最左边
    def _minimum(self, node):
        while node.left:
            node = node.left
        return node

    # 获取最大元素
    def maximum(self):
        if self.map_size == 0:
            return None
        return self._maximum(self.root)

    # 最大元素在最右边
    def _maximum(self, node):
        while node.right:
            node = node.right
        return node

    # 删除最小节点
    def remove_min(self):
        node = self.minimum()
        self._remove_min(self.root)
        return node

    def _remove_min(self, node):
        # 如果没有左侧节点，说明当前节点是最小节点，将当前节点右节点返回
        if not node.left:
            right_node = node.right
            node.right = None
            self.map_size -= 1
            return right_node
        # 最小节点的父节点左子树为最小节点的右子树
        node.left = self._remove_min(node.left)
        return node

    # 删除最大节点
    def remove_max(self):
        node = self.maximum()
        self._remove_max(self.root)
        return node

    def _remove_max(self, node):
        # 如果没有右侧节点，说明当前节点是最大节点，将当前节点左节点返回
        if not node.right:
            left_node = node.left
            node.left = None
            self.map_size -= 1
            return left_node
        # 最大节点的父节点右子树为最大节点的左子树
        node.right = self._remove_max(node.right)
        return node

    # 删除任意节点，O(logn)
    def remove(self, key):
        self.root = self._remove(self.root, key)

    def _remove(self, node, key):
        if not node:
            return
        # 在左子树里删除
        if key < node.key:
            node.left = self._remove(node.left, key)
            return node
        # 在右子树里删除
        elif key > node.key:
            node.right = self._remove(node.right, key)
        # 删除当前节点，重建子树
        else:
            # 只有右子树，直接删除右子树节点
            if not node.left:
                right_node = node.right
                node.right = None
                self.map_size -= 1
                return right_node
            # 只有左子树，直接删除左子树节点
            if not node.right:
                left_node = node.left
                node.left = None
                self.map_size -= 1
                return left_node
            # 即有左子树，又有右子树
            # 找到右子树最小节点，将最小节点从右子树删除，并当作当前节点
            # 当前节点左子树为原节点左子树，右子树为删除了最小节点的右子树
            successor = self.minimum(node.right)
            successor.right = self.remove_min(node.right)
            successor.left = node.left
            # 删除node节点
            node.left = node.right = None
            return successor


customize_map = CustomizeMap()
print(customize_map.is_empty())
customize_map.add(3, 3)
print(customize_map.size())
customize_map.add(1, 1)
customize_map.add(6, 6)
customize_map.add(4, 4)
customize_map.add(9, 9)
customize_map.add(2, 2)
customize_map.add(8, 8)
customize_map.add(7, 7)
customize_map.add(5, 5)
print(customize_map.is_empty())
print(customize_map.size())

# 测试contains方法
print("---------- contains ----------")
print(customize_map.contains(5))
print(customize_map.contains(50))

# 测试get方法
print("---------- get ----------")
print(customize_map.get(5))
print(customize_map.get(50))

# 测试set方法
print("---------- set ----------")
customize_map.set(5, 50)
print(customize_map.get(5))

# 获取最小元素
print("---------- minimum ----------")
print(customize_map.minimum().value)

# 获取最大元素
print("---------- maximum ----------")
print(customize_map.maximum().value)

# 删除最小节点
print("---------- remove min ----------")
customize_map.remove_min()
print(customize_map.minimum().value)

# 删除最大节点
print("---------- remove max ----------")
customize_map.remove_max()
print(customize_map.maximum().value)

# 删除任意节点
print("---------- remove ----------")
print(customize_map.get(5))
customize_map.remove(5)
print(customize_map.get(5))
