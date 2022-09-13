import types


def get_grades_out(self):
    if self.score >= 90:
        return 'A'
    elif self.score >= 60:
        return 'B'
    else:
        return 'C'


class Person(object):
    # 类属性，实例属性优先级高于类属性
    count = 0
    # 外部是无法访问的
    __counter = 0

    def __init__(self, name, job, score):
        Person.count += 1
        self.name = name
        # 双下划线开头的属性是无法被外部访问的
        self.__job = job
        # __属性名__这种是特殊属性，可以被外部访问
        self.score = score


lee = Person('leeyoung', 'student', 90)
# print(lee.__job)    # 这种访问是错误的


# ma = Person('ma', 'student', 80)
# # 将函数动态的添加到实例上
# ma.get_grades_in = types.MethodType(get_grades_out, ma, Person)
# print(ma.get_grades_in)


# 类方法
class Dog(object):
    count = 0
    __counter = 0

    def __init__(self, name):
        self.name = name
        Dog.count += 1
        Dog.__counter += 1

    # 定义类方法
    @classmethod
    def how_many(cls):
        return cls.count


# 调用类方法
print(Dog.how_many())
dog_small = Dog('dahuang')
print(Dog.how_many())
