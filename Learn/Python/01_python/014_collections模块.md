# collections模块用法


### namedtuple

```python
from collections import namedtuple


People = namedtuple('People', ['name', 'age', 'address'])

lee = People('lee', 20, 'jilin')

zhao = People(name='zhao', age=30, address='henan')

print(lee.name)
```


### deque（双端队列）

```python
from collections import deque


d = deque()

d.append(1)
d.appendleft(2)

for item in d:
    print(item)
```


### Counter

```python
from collections import Counter


string = 'leeyoung love zhaomeng'

count = Counter(string)
print(count)
```


### orderDict（有序字典）

```python
from collections import OrderedDict


order_dict = OrderedDict()

order_dict['c'] = 'c'
order_dict['b'] = 'b'

print(order_dict.keys())
```


### defaultdict（默认字典）

```python
from collections import defaultdict


default_dict = defaultdict(int)

print(default_dict['a'])
```
