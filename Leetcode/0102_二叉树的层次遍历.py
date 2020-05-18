# -*- coding: utf-8 -*-
# @Time:    2020/2/28 12:46 PM
# @Author:  leeyoung
# @File:    0102.py
# @Content:


def level_order(node):
    result = list()

    node_list = [node] if node else node
    while node_list:
        children_node_list = list()
        value_list = list()
        for node in node_list:
            if not node:
                continue
            if node.left:
                children_node_list.append(node.left)
            if node.right:
                children_node_list.append(node.right)
            value_list.append(node.value)

        result.append(value_list)
        node_list = children_node_list

    return result
