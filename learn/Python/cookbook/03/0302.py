# 执行精确的小数计算

from decimal import Decimal, localcontext

a = 4.2
b = 2.1
print(a + b)    # 6.300000000000001


a = Decimal('4.2')
b = Decimal('2.1')
print(a + b)


a = Decimal('1.3')
b = Decimal('1.7')
with localcontext() as ctx:
    ctx.prec = 3
    print(a/b)

with localcontext() as ctx:
    ctx.prec = 50
    print(a/b)
print('------------------')