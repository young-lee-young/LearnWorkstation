# 字符串拼接,注意函数的调用顺序

def combine(source, maxsize):
    parts = []
    size = 0
    for part in source: # 此时函数sample函数才开始执行
        parts.append(part)
        size += len(part)
        if size > maxsize:
            yield ' '.join(parts)
            parts = []
            size = 0
    yield ' '.join(parts)


def sample():
    yield 'IS'
    yield 'Chicago'
    yield 'Not'
    yield 'Chicago?'


for part in combine(sample(), 3):   # 把一个函数作为变量传过去,此时函数并没有执行
    print(part)
print('-----------------')