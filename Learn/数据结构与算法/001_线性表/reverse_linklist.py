# -*- coding: utf-8 -*-
# @Time:    2020/2/26 4:31 PM
# @Author:  leeyoung
# @File:    reverse.py
# @Content: 反转链表

from customize_linklist import CustomizeLinkList


def reverse_list(head):
    pre = None
    cur = head

    while cur:
        next_node = cur.next
        cur.next = pre
        pre = cur
        cur = next_node
    return pre


# 原链表 1 -> 2 -> 3 -> None
link_list = CustomizeLinkList()
for i in range(10):
    link_list.add(i)

node = reverse_list(link_list.root)
# 现链表 3 -> 2 -> 1 -> None
