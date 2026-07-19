# pytorch 卷积层


### 使用

* 代码示例

```python
import torch
import torchvision
from torch import nn
from torch.nn import Conv2d
from torch.utils.data import DataLoader
from torch.utils.tensorboard import SummaryWriter


# 数据集
dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
test_data = torchvision.datasets.CIFAR10(dataset_dir, train=False, transform=torchvision.transforms.ToTensor())


# 加载数据集
data_loader = DataLoader(dataset=test_data, batch_size=64)


class MyNN(nn.Module):
    def __init__(self):
        super(MyNN, self).__init__()
        # in_channel：3 个通道，彩色图片有 3 个通道
        # out_channel：
        # kernel size：3，会生成 3 * 3 的卷积核
        self.conv1 = Conv2d(3, 6, 3, stride=1, padding=0)

    def forward(self, input):
        output = self.conv1(input)
        return output


step = 0
writer = SummaryWriter("logs")
mynn = MyNN()

for data in data_loader:
    imgs, target = data
    output = mynn(imgs)

    print(imgs.shape)
    # 输出：
    # torch.Size([64, 3, 32, 32])
    writer.add_images("input", imgs, step)

    print(output.shape)
    # 输出：
    # torch.Size([64, 6, 30, 30])
    output = torch.reshape(output, (-1, 3, 30, 30))
    writer.add_images("output", output, step)

    step = step + 1

writer.close()
```
