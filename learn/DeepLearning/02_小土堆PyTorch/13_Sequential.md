# Sequential


### 使用

* 代码示例

```python
import torch
from torch import nn
from torch.nn import Conv2d, Flatten, Linear, MaxPool2d, Sequential
from torch.utils.tensorboard import SummaryWriter


# 生成输入数据
input = torch.ones((64, 3, 32, 32))


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        self.model1 = Sequential(
            Conv2d(3, 32, 5, padding=2),
            MaxPool2d(2),
            Conv2d(32, 32, 5, padding=2),
            MaxPool2d(2),
            Conv2d(32, 64, 5, padding=2),
            MaxPool2d(2),
            Flatten(),
            Linear(1024, 64),
            Linear(64, 10),
        )

    def forward(self, input):
        output = self.model1(input)
        return output


mynn = MyNN()
output = mynn(input)
print(output.shape)


# 可视化
writer = SummaryWriter("logs")
writer.add_graph(mynn, input)
writer.close()
```
