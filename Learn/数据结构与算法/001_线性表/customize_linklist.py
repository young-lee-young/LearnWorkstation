# -*- coding: utf-8 -*-
# @Time:    2020/3/17 11:03 AM
# @Author:  leeyoung
# @File:    customize_linklist.py
# @Content: 自定义链表

from customize_listnode import CustomizeListNode


class CustomizeLinkList:
    def __init__(self):
        self.root = None
        self.link_size = 0

    def size(self):
        return self.link_size

    def is_empty(self):
        return True if self.link_size == 0 else False

    # 包含元素
    def contains(self, num):
        cur = self.root
        while cur:
            if cur.value == num:
                return True
            cur = cur.next
        return False

    # 在链表头添加元素
    def add_head(self, num):
        node = CustomizeListNode(num)
        node.next = self.root
        self.root = node
        self.link_size += 1

    # 在index处添加元素
    def add(self, index, num):
        if index != 0:
            prev = self.root
            for i in range(index - 1):
                prev = prev.next
            node = CustomizeListNode(num)
            node.next = prev.next
            prev.next = node
            self.link_size += 1
        else:
            self.add_head(num)

    # 在链表尾部添加元素
    def add_tail(self, num):
        self.add(self.link_size, num)

    # 更新元素
    def update(self, index, num):
        cur = self.root
        for i in range(index):
            cur = cur.next
        cur.value = num

    # 获取元素
    def get(self, index):
        cur = self.root
        for i in range(index):
            cur = cur.next
        return cur.value

    # 根据索引删除
    def remove(self, index):
        prev = self.root
        for i in range(index - 1):
            prev = prev.next
        index_node = prev.next
        prev.next = index_node.next
        self.link_size -= 1

    # 根据元素删除
    def remove_element(self, num):
        prev = self.root
        while prev.next:
            if prev.next.value == num:
                break
            prev = prev.next
        if prev.next:
            delete_node = prev.next
            prev.next = delete_node.next
            self.link_size -= 1


def main():
    link_list = CustomizeLinkList()
    for i in range(5):
        link_list.add(i, i)
    link_list.add_tail(10)
    link_list.remove(2)

    print(link_list.get(2))
    print(link_list.contains(4))
    link_list.update(2, 15)
    print(link_list)


if __name__ == '__main__':
    main()
