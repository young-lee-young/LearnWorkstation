# -*- coding: utf-8 -*-
# @Time:    2021/1/11 下午11:01
# @Author:  leeyoung
# @File:    customize_map_base_linklist.py
# @Content: 自定义映射 - 基于链表实现

from customize_map_node import CustomizeMapLinkListNode


class CustomizeMapBaseLinkList:
    def __init__(self):
        self.head = CustomizeMapLinkListNode(None, None)
        self.map_size = 0

    def size(self):
        return self.map_size

    def is_empty(self):
        return self.map_size == 0

    #   根据键获取值
    def get_node(self, key):
        current_node = self.head.next
        while current_node is not None:
            if current_node.key == key:
                return current_node
            current_node = current_node.next
        return None

    # 是否存在，O(n)
    def contains(self, key):
        return self.get_node(key) is not None

    # 获取值，O(n)
    def get(self, key):
        node = self.get_node(key)
        if node is None:
            return None
        return node.value

    # 添加新节点，O(n)
    def add(self, key, value):
        node = self.get_node(key)
        if node is None:
            self.head.next = CustomizeMapLinkListNode(key, value, self.head.next)
            self.map_size += 1
        else:
            node.value = value

    # 设置新值，O(n)
    def set(self, key, value):
        node = self.get_node(key)
        if node is None:
            return
        node.value = value

    # 删除节点，O(n)
    def remove(self, key):
        prev = self.head
        while prev.next:
            if prev.next.key == key:
                break
            prev = prev.next
        if prev.next:
            delete_node = prev.next
            prev.next = delete_node.next
            self.map_size -= 1


customize_map = CustomizeMapBaseLinkList()
customize_map.add("lee", 1)
customize_map.add("zhao", 2)
print(customize_map.size())
print(customize_map.is_empty())
print(customize_map.contains("lee"))
print(customize_map.get("lee"))
customize_map.set("lee", 3)
print(customize_map.get("lee"))
customize_map.remove("zhao")
print(customize_map.size())
