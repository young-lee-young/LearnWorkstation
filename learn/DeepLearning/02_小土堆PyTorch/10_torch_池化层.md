# pytorch 池化层


### 基础

* 最大池化的目的是什么

保留特征，减少数据


### 使用

* 代码示例

```python
import torch
from torch import nn
from torch.nn import MaxPool2d

input = torch.tensor([[1, 2, 0, 3, 1],
                      [0, 1, 2, 3, 1],
                      [1, 2, 1, 0, 0],
                      [5, 2, 3, 1, 1],
                      [2, 1, 0, 1, 1]])

# MaxPool2d 需要 (batch, 通道，高，宽) 这样的数据，所以需要进行 reshape
input = torch.reshape(input, (-1, 1, 5, 5))


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        self.maxpool1 = MaxPool2d(kernel_size=3, ceil_mode=False)

    def forward(self, input):
        output = self.maxpool1(input)
        return output


mynn = MyNN()
output = mynn(input)
print(output)
```

* 可视化

```python
import torchvision
from torch import nn
from torch.nn import MaxPool2d
from torch.utils.data import DataLoader
from torch.utils.tensorboard import SummaryWriter


# 数据集
dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
test_data = torchvision.datasets.CIFAR10(dataset_dir, train=False, transform=torchvision.transforms.ToTensor(), download=True)


# 加载数据集
data_loader = DataLoader(dataset=test_data, batch_size=64)


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        self.maxpool1 = MaxPool2d(kernel_size=3, ceil_mode=True)

    def forward(self, input):
        output = self.maxpool1(input)
        return output


step = 0
writer = SummaryWriter("logs")
mynn = MyNN()

for data in data_loader:
    imgs, target = data
    output = mynn(imgs)

    writer.add_images("input", imgs, step)

    writer.add_images("output", output, step)

    step = step + 1

writer.close()
```
