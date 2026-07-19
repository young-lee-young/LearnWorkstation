# GPU 训练


### 基础

需要修改 模型、数据（输入和标注）、损失函数 3 个地方


### 方式 1

通过 .cuda() 的方式

* 代码示例

```python
import torchvision
from torch import nn
from torch.utils.data import DataLoader


dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
train_data = torchvision.datasets.CIFAR10(root=dataset_dir, train=True, transform=torchvision.transforms.ToTensor(), download=True)
train_dataloader = DataLoader(train_data, batch_size=64)


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        pass

    def forward(self, input):
        pass


mynn = MyNN()
# 修改模型
mynn = mynn.cuda()


loss_fn = nn.CrossEntropyLoss()
# 修改损失函数
loss_fn = loss_fn.cuda()


for data in train_dataloader:
    imgs, targets = data
    # 修改输入和标注
    imgs = imgs.cuda()
    targets = targets.cuda()
    
    outputs = mynn(imgs)
```


### 方式 2

通过 device 的方式

* 代码示例

```python
import torch
import torchvision
from torch import nn
from torch.utils.data import DataLoader


device = torch.device("cuda:0" if torch.cuda.is_available() else "cpu")


dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
train_data = torchvision.datasets.CIFAR10(root=dataset_dir, train=True, transform=torchvision.transforms.ToTensor(), download=True)
train_dataloader = DataLoader(train_data, batch_size=64)


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        pass

    def forward(self, input):
        pass


mynn = MyNN()
# 修改模型
mynn.to(device)


loss_fn = nn.CrossEntropyLoss()
# 修改损失函数
loss_fn.to(device)


for data in train_dataloader:
    imgs, targets = data
    # 修改输入和标注
    imgs.to(device)
    targets.to(device)
    
    outputs = mynn(imgs)
```
