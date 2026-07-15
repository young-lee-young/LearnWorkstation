# transforms


### ToTensor

* 作用

将 PIL Image 或 ndarray 转成 tensor

* 代码示例

```python
from PIL import Image
from torchvision import transforms
from torch.utils.tensorboard import SummaryWriter


img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"
img = Image.open(img_path)


trans_totensor = transforms.ToTensor()
img_totensor = trans_totensor(img)
# 输出：<class 'torch.Tensor'>
print(type(img_totensor))


writer = SummaryWriter("logs")
# 第二个参数是 torch.Tensor 类型
writer.add_image("totensor", img_totensor)
writer.close()
```


### Normalize

* 作用

归一化数据

* 代码示例

```python
from PIL import Image
from torchvision import transforms
from torch.utils.tensorboard import SummaryWriter


img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"
img = Image.open(img_path)


trans_totensor = transforms.ToTensor()
img_totensor = trans_totensor(img)
print(img_totensor.shape)
print(img_totensor[0][0][0])


# 参数为均值和标准差
# 图片是 3 个维度的数据（通道、长、宽），所以是均值和标准差是 3 维
# 比如 [2][10][20] 表示 [10][20] 这个像素的绿色通道的值
trans_norm = transforms.Normalize([0.5, 0.5, 0.5], [0.5, 0.5, 0.5])
img_norm = trans_norm(img_totensor)
print(img_norm[0][0][0])


writer = SummaryWriter("logs")
writer.add_image("normalize", img_totensor)
writer.close()
```


### Resize

* 作用

将 PIL Image 转换形状

* 代码示例

```python
from PIL import Image
from torchvision import transforms
from torch.utils.tensorboard import SummaryWriter


img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"
img = Image.open(img_path)
print(img.size)


trans_resize = transforms.Resize((512, 512))
# resize 后还是一个 PIL Image
img_resize = trans_resize(img)


# 将 PIL Image 转换成 tensor
trans_totensor = transforms.ToTensor()
img_resize_totensor = trans_totensor(img_resize)


writer = SummaryWriter("logs")
writer.add_image("resize", img_resize_totensor)
writer.close()
```


### Compose

* 作用

将多个 transforms 操作组合起来

* 代码示例

```python
from PIL import Image
from torchvision import transforms
from torch.utils.tensorboard import SummaryWriter


img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"
img = Image.open(img_path)


trans_resize = transforms.Resize(512)
trans_totensor = transforms.ToTensor()
# 组合 Resize 和 ToTensor
trans_compose = transforms.Compose([trans_resize, trans_totensor])


img_resize_totensor = trans_compose(img)


writer = SummaryWriter("logs")
writer.add_image("compose", img_resize_totensor)
writer.close()
```


### RandomCrop

* 作用

随机裁剪 PIL Image

* 代码示例

```python
from PIL import Image
from torchvision import transforms
from torch.utils.tensorboard import SummaryWriter


img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"
img = Image.open(img_path)
print(img.size)


trans_randomcrop = transforms.RandomCrop(512)
trans_totensor = transforms.ToTensor()
trans_compose = transforms.Compose([trans_randomcrop, trans_totensor])


img_randomcrop = trans_compose(img)
print(img_randomcrop.size)


writer = SummaryWriter("logs")
writer.add_image("randomcrop", img_randomcrop)
writer.close()
```
