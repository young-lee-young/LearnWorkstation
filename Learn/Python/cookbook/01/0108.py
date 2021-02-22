# 与字典相关的计算

d = {}
for i in range(10):
    d[str(i)] = i


zipped = zip(d.values(), d.keys())
print(sorted(zipped))
print('------------------------')