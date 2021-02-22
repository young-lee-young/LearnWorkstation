# 不同的分隔符或行结尾符完成打印

print('liyao', 22, 'boy')

print('liyao', 22,'boy', sep=',')   # 指定分隔符

print('liyao', 22, 'boy', end='!!!')    # 指定结尾符


for i in range(5):
    print(i)

for i in range(5):
    print(i, end=' ')   # 使用end可以禁止打印换行符