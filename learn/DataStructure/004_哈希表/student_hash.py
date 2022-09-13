# -*- coding: utf-8 -*-
# @Time:    2020/3/17 9:25 AM
# @Author:  leeyoung
# @File:    student_hash.py
# @Content: 模拟Java的hashCode


class StudentHash:
    def __init__(self, grade, cls, first_name, last_name):
        self.grade = grade
        self.cls = cls
        self.first_name = first_name
        self.last_name = last_name

    def hash_code(self):
        B = 31

        hash_code = 0
        hash_code = hash_code * B + self.grade
        hash_code = hash_code * B + self.cls
        hash_code = hash_code * B + hash(self.first_name.lower())
        hash_code = hash_code * B + hash(self.last_name.lower())

        return hash_code


student_hash = StudentHash(6, 3, 'lee', 'young')
print(student_hash.hash_code())
