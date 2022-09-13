# -*- coding: utf-8 -*-
# @Time:    2021/1/11 下午11:17
# @Author:  leeyoung
# @File:    customize_mapnode.py
# @Content:


class CustomizeMapLinkListNode:
    def __init__(self, key, value, next=None):
        self.key = key
        self.value = value
        self.next = next


class CustomizeMapTreeNode:
    def __init__(self, key, value, left=None, right=None):
        self.key = key
        self.value = value
        self.left = left
        self.right = right
