
# 特殊方法的特点
# 1、不需要定义在class中
# 2、不需要直接调用
# Python的某些函数或操作符会调用对应的特殊方法


num_list = [1, 2, 3]
print(num_list) # 实际调用的是num_list.__str__()方法


class Person(object):
    def __init__(self, name):
        self.name = name

    def __call__(self, friend):
        print('My name is %s' %self.name)
        print('My friends is %s' %friend)



lee = Person('lee') # lee是一个类的实例
print(lee)  # 实际调用的是lee.__str__()方法

# 使用__call__函数使一个类变成了可调用对象
# lee是一个实例，但是下面你无法确定是一个函数还是一个实例，在Python中函数也是对象，对象和函数的区别并不大
print(lee('ma'))


# __str__方法
class Student(object):
    def __init__(self, name, gender, score):
        self.name = name
        self.gender = gender
        self.score = score

    # 定制返回的字符串
    def __str__(self):
        return 'Person: %s, %s' % (self.name, self.gender)

    # len函数
    def __len__(self):
        return len(self.name)

    __repr__ = __str__

    # cmp函数Python3中被放弃了
    # def __cmp__(self, s):
    #     if self.name < s.name:
    #         return -1
    #     elif self.name > s.name:
    #         return 1
    #     else:
    #         return 0


leeyoung = Student('leeyoung', 'male', 21)
print(leeyoung)

L = [Student('lee', 'male', 21), Student('ma', 'male', 22), Student('liu', 'male', 23)]
# print(sorted(L))

print(len(L))


# 数学运算
class Rational(object):
    def __init__(self, num):
        self.num = num

    # 相加函数
    def __add__(self, other):
        return Rational(self.num + other.num)

    # 相减函数
    def __sub__(self, other):
        return Rational(self.num - other.num)

    # 相乘函数
    def __mul__(self, other):
        return Rational(self.num * other.num)

    # 相除函数
    def __truediv__(self, other):
        return Rational(self.num / other.num)

    def __int__(self):
        return int(self.num)

    def __float__(self):
        return float(self.num)

    def __str__(self):
        return '%s' % self.num

    __repr__ = __str__


num1 = Rational(10)
num2 = Rational(2)
print(num1 + num2)
print(num1 - num2)
print(num1 * num2)
print(num1 / num2)

print(int(num1 / num2))
print(float(num1 + num2))



# property
class Teacher(object):
    # 限制添加的属性
    __slots__ = ('name', 'age', 'classnum')
    def __init__(self, name, classnum):
        self.name = name
        self.__classnum = classnum

    @property   # 把getter方法变成属性，property本身又创建了一个装饰器，@classnum.setter，负责把一个setter方法变成属性赋值
    def classnum(self):
        return self.__classnum

    @classnum.setter
    def classnum(self, classnum):
        if classnum < 0 or classnum > 10:
            # 报错信息写法
            raise ValueError('invalid classnum')
        else:
            self.__classnum = classnum

    @property
    def grade(self):
        if self.classnum >= 8:
            return 'A'
        elif self.classnum >= 5:
            return 'B'
        else:
            return 'C'

teacher1 = Teacher('guo', 5)
print(teacher1.classnum)
# teacher1.classnum = 20  # 写成这种会报错
# print(teacher1.classnum)


print(teacher1.grade)

# Python是动态语言，允许在运行过程中添加属性，
teacher1.age = 21
print(teacher1.age)
# 使用__slots__限制添加属性，如果一个子类继承父类，父类有一个__slots__，子类中只需要加入新的__slots__就可以了
