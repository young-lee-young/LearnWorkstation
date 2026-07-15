# torchvision 数据集


### CIFAR-10 数据集

* 地址

https://cave.cs.toronto.edu/kriz/cifar.html


### 数据集使用

```python
import torchvision

dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"

train_set = torchvision.datasets.CIFAR10(root=dataset_dir, train=True, download=True)
test_set = torchvision.datasets.CIFAR10(root=dataset_dir, train=False, download=True)


print(train_set)
# 输出：
# Dataset CIFAR10
#     Number of datapoints: 50000
#     Root location: /Users/lee/Workstation/DeepLearning/dataset
#     Split: Train
print(test_set)
# 输出：
# Dataset CIFAR10
#     Number of datapoints: 10000
#     Root location: /Users/lee/Workstation/DeepLearning/dataset
#     Split: Test


test_0 = test_set[0]
print(test_0)
# 输出：(<PIL.Image.Image image mode=RGB size=32x32 at 0x140515950>, 3)
# 是一个元组，第一个元素表示 PIL Image 的图片，第二个 3 表示数据的 tag


img, tag = test_0
# test_set 的第 3 个标签是 cat
print(test_set.classes[tag])

img.show()
```


### 数据集和 torchvision 结合

```python
import torchvision
from torch.utils.tensorboard import SummaryWriter


# 将 dataset 中的 PIL Image 转成 Tensor
dataset_transform = torchvision.transforms.Compose(
    [torchvision.transforms.ToTensor()]
)


dataset_dir = "/Users/lee/Workstation/DeepLearning/dataset"
train_set = torchvision.datasets.CIFAR10(root=dataset_dir, train=True, transform=dataset_transform, download=True)
test_set = torchvision.datasets.CIFAR10(root=dataset_dir, train=False, transform=dataset_transform, download=True)


writer = SummaryWriter("logs")
for i in range(10):
    img, tag = test_set[i]
    writer.add_image("cifar10", img, i)
writer.close()
```
