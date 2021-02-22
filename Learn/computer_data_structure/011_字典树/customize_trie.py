# -*- coding: utf-8 -*-
# @Time:    2021/1/17 下午3:46
# @Author:  leeyoung
# @File:    customize_dict_tree.py
# @Content: 字典树（专门为处理字符串设计）

# from trie_node import TrieNode


class TrieNode:
    def __init__(self):
        self.is_word = False
        self.next = dict()


class CustomizeTrie:
    def __init__(self):
        self.root = TrieNode()
        self.trie_size = 0

    # 获取trie中存储的单词数量
    def size(self):
        return self.trie_size

    # 向trie中添加一个新单词
    def add(self, word):
        current = self.root
        # 循环单词每个字符
        for item in word:
            # 如果字符不存在，添加新的字符
            if item not in current.next.keys():
                current.next[item] = TrieNode()
            # 更新当前节点
            current = current.next[item]

        # 如果以前字典树里不存在这个单词，将单词标志设置为True
        if not current.is_word:
            current.is_word = True
            self.trie_size += 1

    # 查询trie中是否包含单词
    def contains(self, word):
        current = self.root
        # 循环单词每个字符，并判断当前字符是否存在
        for item in word:
            if item not in current.next.keys():
                return False
            current = current.next[item]
        # 如果当前不是单词，返回False
        if not current.is_word:
            return False
        return True

    # 查询是否存在前缀为prefix的单词
    def prefix(self, prefix):
        current = self.root
        for item in prefix:
            if item not in current.next.keys():
                return False
            current = current.next[item]
        return True


trie = CustomizeTrie()

# 测试add方法
print('---------- add ----------')
trie.add("lee")
trie.add("zhao")

# 测试size方法
print('---------- trie size ----------', trie.size())

# 测试contains方法
print('---------- contains lee ----------', trie.contains('lee'))
print('---------- contains young ----------', trie.contains('young'))

# 测试prefix方法
print('---------- prefix zh ----------', trie.prefix('zh'))
print('---------- prefix li ----------', trie.prefix('li'))
