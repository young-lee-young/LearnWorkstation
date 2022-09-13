
# 类的继承（复用代码）
# 1、新类不必从头编写
# 2、新类从现有的类继承，就自动拥有了现有类的所有功能
# 3、新类只需要编写现有类缺少的新功能


class Person(object):
    def __init__(self, name):
        self.name = name

    def whoAmI(self):
        return 'I am a Person, my name is %s' % self.name

class Student(Person):
    def __init__(self, name, score):
        # 初始化父类
        super(Student, self).__init__(name)
        self.score = score

    def whoAmI(self):
        return 'I am a Student, my name is %s' % self.name


guo = Person('guo')
lee = Student('lee', 90)


# 判断类型
print(isinstance(guo, Person))
print(isinstance(lee, Person))
print(isinstance(guo, Student))


# 多态：方法调用将作用在person的实际类型上，先查找自身的定义，如果没有定义，则顺着继承链向上查找，直到在某个父类中找到为止
def who_am_i(person):
    print(person.whoAmI())


print(who_am_i(guo))
print(who_am_i(lee))


# 多重继承


# 获取属性
print(getattr(lee, 'name'))
setattr(guo, 'name', 'guowei')
print(getattr(guo, 'name'))
# 获取age属性，如果不存在，默认是20
print(getattr(lee, 'age', 20))