# enumerate函数使用

枚举，列举

### 

```python
name_list = ['liyao', 'macheng', 'liwei', 'liangbo', 'liuyafei', 'pengzeyu']

for index, name in enumerate(name_list):
    print(index, name)
# 0 liyao
# 1 macheng
# 2 liwei
# 3 liangbo
# 4 liuyafei
# 5 pengzeyu
```

### 指定索引初始值

```python
name_list = ['liyao', 'macheng', 'liwei', 'liangbo', 'liuyafei', 'pengzeyu']

for index, name in enumerate(name_list, 1):
    print(index, name)
# 1 liyao
# 2 macheng
# 3 liwei
# 4 liangbo
# 5 liuyafei
# 6 pengzeyu
```
