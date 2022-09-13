# 反向迭代

num_list = [i for i in range(1, 10)]
print(num_list)


for i in reversed(num_list):    # 把列表反转再迭代
    print(i)
print('------------------')


# 在类里面自实现__reversed__方法,就可以在类上实现反向迭代
class Countup:
    def __init__(self, start):
        self.start = start


    def __iter__(self):
        n = 1
        while n <= self.start:
            yield n
            n += 1


    def __reversed__(self):
        n = self.start
        while n > 0:
            yield n
            n -= 1


for i in Countup(5):
    print(i)
print('------------------')


for i in reversed(Countup(5)):
    print(i)