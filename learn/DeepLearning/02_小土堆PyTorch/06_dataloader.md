# DataLoader


### 基础

* 作用

从 DataSet 加载数据


### 使用

* 代码示例

```python
import torchvision
from torch.utils.data import DataLoader
from torch.utils.tensorboard import SummaryWriter


dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
# 数据集
test_data = torchvision.datasets.CIFAR10(dataset_dir, train=False, transform=torchvision.transforms.ToTensor())


# 加载数据
# num_workers：加载数据的线程数
# drop_last：如果为 True，如果最后的数据不够 batch_size，会丢弃最后的数据
test_loader = DataLoader(dataset=test_data, batch_size=4, shuffle=True, num_workers=0, drop_last=False)


writer = SummaryWriter("logs")
step = 0

# 循环每个 batch
for data in test_loader:
    imgs, targets = data
    print(imgs.shape)
    # 输出：
    # torch.Size([4, 3, 32, 32])
    # 表示 4 张图片，每张图片 3 通道，长 32，宽 32
    print(targets)
    # 输出：
    # tensor([2, 7, 7, 1])
    # 表示 4 张图片的 label
    writer.add_images("dataloader", imgs, step)
    step = step + 1

writer.close()
```
