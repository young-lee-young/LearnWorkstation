# -*- coding: utf-8 -*-
# @Time:    2020/2/27 9:24 PM
# @Author:  leeyoung
# @File:    reverse_tree.py
# @Content: 反转二叉树


class ReverseBinTree:
    def reverse_bin_tree(self, node):
        if node:
            node.left, node.right = node.right, node.left
            self.reverse_bin_tree(node.left)
            self.reverse_bin_tree(node.right)
        return node
