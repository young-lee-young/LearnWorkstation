# pytorch Module 模块


### Module

* 代码示例

```python
import torch
from torch import nn


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()

    # nn.Module 需要重写 forward 方法
    def forward(self, input):
        output = input + 1
        return output


mynn = MyNN()
x = torch.tensor(1.0)
out = mynn.forward(x)
print(out)
```
