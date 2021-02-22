# -*- coding: utf-8 -*-
# @Time:    2021/1/17 下午4:29
# @Author:  leeyoung
# @File:    trie_node.py
# @Content:


class TrieNode:
    def __init__(self, is_word):
        self.is_word = is_word
        self.next = dict()
