# Python有三种类型方法，实例方法，类方法，静态方法

name = 'lee'


class Testclass(object):
    def foo(self, param):
        pass

    @classmethod
    def class_foo(cls, param):
        pass

    @staticmethod
    def static_foo(param):
        pass


testclass = Testclass()


# testclass.foo(name)相当于foo(testclass, name)
# 类方法调用, testclass.class_foo或者TestClass.class_foo(name)
# 静态方法调用, testclass.static_foo(name)或者Testclass.static_class(name)



# 类变量和实例变量
# 类变量是实例之间共享的变量
# 实例变量是实例自己独有的变量
class Test(object):
    num_of_instance = 0   # 类变量

    def __init__(self, name):
        self.name = name    # 实例变量
        Test.num_of_instance += 1
