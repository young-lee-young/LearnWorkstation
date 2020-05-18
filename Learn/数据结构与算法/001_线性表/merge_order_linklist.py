# -*- coding: utf-8 -*-
# @Time:    2020/2/27 9:59 AM
# @Author:  leeyoung
# @File:    merge_order_listnode.py
# @Content: 合并两个有序链表

from customize_linklist import CustomizeLinkList


class Solution:
    def __init__(self, node_1, node_2):
        self.node_1 = node_1
        self.node_2 = node_2
        self.new_link_list = CustomizeLinkList()

    def merge_order(self):
        # 两个链表均未循环到结尾
        while self.node_1 and self.node_2:
            if self.node_1.value < self.node_2.value:
                self.new_link_list.add(self.node_1.value)
                self.node_1 = self.node_1.next
            else:
                self.new_link_list.add(self.node_2.value)
                self.node_2 = self.node_2.next

        # 链表1不为空
        while self.node_1:
            self.new_link_list.add(self.node_1.value)
            self.node_1 = self.node_1.next
        # 链表2不为空
        while self.node_2:
            self.new_link_list.add(self.node_2.value)
            self.node_2 = self.node_2.next

        return self.new_link_list.root


link_list_one = CustomizeLinkList()
for i in range(1, 10, 2):
    link_list_one.add(i)

link_list_two = CustomizeLinkList()
for i in range(2, 11, 2):
    link_list_two.add(i)

solution = Solution(link_list_one.root, link_list_two.root)
head = solution.merge_order()
print(head)
