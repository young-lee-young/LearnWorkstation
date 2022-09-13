# -*- coding: utf-8 -*-
# @Time:    2020/3/17 9:06 AM
# @Author:  leeyoung
# @File:    customize_hash.py
# @Content: 自定义哈希表


"""
没有扩容机制的哈希表时间复杂度

总共元素为n，哈希表长度为m，平均每个地址开辟的链表长度为n/m，时间复杂度为O(log(n/m))

有扩容机制的哈希表时间复杂度

每n次操作触发一次扩容，每次扩容时间复杂度为O(n)，将扩容操作平均到前n次操作，均摊时间复杂度为O(1)
"""


class CustomizeHash:
    capacity_list = [53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613, 393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 100663319, 201326611, 402653189, 805306457, 1610612741]

    def __init__(self):
        # 扩容列表索引
        self.capacity_index = 0
        # 哈希表的容量
        self.capacity = self.capacity_list[self.capacity_index]
        # 存放元素数组
        self.hash_table = [{} for i in range(self.capacity)]
        # 哈希表中数据个数
        self.hash_size = 0
        # 最高容忍每个地址链表上存放10个元素
        self.upper_tolerance = 10
        # 最低容忍每个地址链表上存放2个元素
        self.lower_tolerance = 2

    def size(self):
        return self.hash_size

    # 哈希函数
    def hash(self, key):
        return (key & 0x7fffffff) % self.capacity

    # 总共有capacity个地址，存入元素为hash_size个元素
    def resize(self, new_capacity):
        print("----------------- resize from %d to %d -----------------" % (self.capacity, new_capacity))
        # 初始化新哈希表
        new_hash_table = [{} for i in range(new_capacity)]

        # 更新capacity
        self.capacity = new_capacity

        # 将原哈希表中数据放入新哈希表
        for item in self.hash_table:
            for key, value in item.items():
                new_hash_table[self.hash(key)].update({key: value})

        self.hash_table = new_hash_table

    # 添加操作
    def add(self, key, value):
        element = self.hash_table[self.hash(key)]
        # 更新元素
        if key in element.keys():
            self.hash_table[self.hash(key)][key] = value
        # 添加元素
        else:
            self.hash_table[self.hash(key)].update({key: value})
            self.hash_size += 1

        # 扩容操作，如果元素总数比最大容忍度乘哈希表容量大，并且索引没有越界
        if self.hash_size >= self.upper_tolerance * self.capacity and self.capacity_index + 1 < len(self.capacity_list):
            # 确保扩容、所容后的值为素数
            self.capacity_index += 1
            self.resize(self.capacity_list[self.capacity_index])

    # 删除操作
    def remove(self, key):
        element = self.hash_table[self.hash(key)]
        # 删除元素，更新size
        if key in element.keys():
            del self.hash_table[self.hash(key)][key]
            self.hash_size -= 1

        # 缩容操作
        if self.hash_size < self.lower_tolerance * self.capacity and self.capacity_index - 1 >= 0:
            self.capacity_index -= 1
            self.resize(self.capacity_list[self.capacity_index])

    # 更新操作
    def set(self, key, value):
        element = self.hash_table[self.hash(key)]
        # 如果存在进行更新
        if key not in element.keys():
            raise KeyError("key not exist error")
        self.hash_table[self.hash(key)][key] = value

    # 是否包含key
    def contains(self, key):
        element = self.hash_table[self.hash(key)]
        if key in element.keys():
            return True
        return False

    # 获取元素
    def get(self, key):
        return self.hash_table[self.hash(key)].get(key)


hash_table = CustomizeHash()

# 测试add
for i in range(100):
    hash_table.add(i, i)

# 测试get
print(hash_table.get(1))

# 测试set
hash_table.set(1, "li")
print(hash_table.get(1))

# 测试contains
print(hash_table.contains(15))

# 测试remove
hash_table.remove(15)
print(hash_table.contains(15))

print(hash_table)
