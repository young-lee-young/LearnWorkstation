# -*- coding: utf-8 -*-
# @Time:    2020/2/27 11:44 PM
# @Author:  leeyoung
# @File:    0002_两数相加.py
# @Content:

from customize_listnode import CustomizeListNode


def add_two_number(listnode_1, listnode_2):
    root = CustomizeListNode(None)
    cur = root

    add_one = False
    while listnode_1 and listnode_2:
        value = listnode_1.val + listnode_2.val
        value = value + 1 if add_one else value
        if value >= 10:
            value = value % 10
            add_one = True
        else:
            add_one = False

        listnode_1, listnode_2 = listnode_1.next, listnode_2.next
        node = CustomizeListNode(value)
        cur.next = node
        cur = node

    if not listnode_1 and not listnode_2 and add_one:
        node = CustomizeListNode(1)
        cur.next = node
    elif listnode_1:
        while listnode_1:
            value = listnode_1.val
            value = value + 1 if add_one else value
            if value >= 10:
                value = value % 10
                add_one = True
            else:
                add_one = False
            listnode_1 = listnode_1.next

            node = CustomizeListNode(value)
            cur.next = node
            cur = node
    else:
        while listnode_2:
            value = listnode_2.val
            value = value + 1 if add_one else value
            if value >= 10:
                value = value % 10
                add_one = True
            else:
                add_one = False
            listnode_2 = listnode_2.next

            node = CustomizeListNode(value)
            cur.next = node
            cur = node

    if add_one:
        node = CustomizeListNode(1)
        cur.next = node
    return root.next


listnode_1 = CustomizeListNode(1)
listnode_9 = CustomizeListNode(9)
listnode_2 = CustomizeListNode(9, listnode_9)

node = add_two_number(listnode_1, listnode_2)
print(node)
