# -*- coding: utf-8 -*-
# @Time:    2020/2/27 11:29 PM
# @Author:  leeyoung
# @File:    0023.py
# @Content: 合并K个排序链表

import heapq
from customize_listnode import CustomizeListNode


def merge_more_listnode(node_list):
    value_list = list()

    # 获取链表所有值
    for node in node_list:
        while node:
            value_list.append(node.value)
            node = node.next

    # 构造一个最小堆
    heapq.heapify(value_list)

    # 构造链表
    root_node = CustomizeListNode(None)
    cur = root_node
    while value_list:
        next_node = CustomizeListNode(heapq.heappop(value_list))
        cur.next = next_node
        cur = next_node

    return root_node
