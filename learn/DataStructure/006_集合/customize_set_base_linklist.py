# -*- coding: utf-8 -*-
# @Time:    2020/3/17 10:59 AM
# @Author:  leeyoung
# @File:    customize_set_base_listnode.py
# @Content: 自定义集合 - 基于链表实现

from customize_linklist import CustomizeLinkList


# 性能低
class CustomizeSetBaseLinkList:
    def __init__(self):
        self.link_list = CustomizeLinkList()

    def size(self):
        return self.link_list.size()

    def is_empty(self):
        return self.link_list.is_empty()

    def contains(self, num):
        # O(n)，n为元素个数
        return self.link_list.contains(num)

    def add(self, num):
        # O(n)
        if not self.link_list.contains(num):
            self.link_list.add_head(num)

    def remove(self, num):
        # O(n)
        self.link_list.remove_element(num)


def main():
    customize_set = CustomizeSetBaseLinkList()
    customize_set.add(1)
    customize_set.add(1)
    customize_set.add(2)
    print(customize_set.size())
    print(customize_set.contains(1))
    customize_set.remove(1)
    print(customize_set.size())


if __name__ == '__main__':
    main()
