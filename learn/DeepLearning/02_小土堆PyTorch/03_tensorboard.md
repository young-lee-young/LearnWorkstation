# tensorboard 


### 基础

* 作用

将训练过程可视化

* 安装

```bash
python -m pip install tensorboard
```


### add_scalar 使用

* 生成事件文件

```python
from torch.utils.tensorboard import SummaryWriter

# 会在 logs 文件夹下生成一个文件
writer = SummaryWriter("logs")

for x in range(100):
    # 曲线名称
    # 竖坐标值
    # 横坐标值
    y = 2 * x
    writer.add_scalar("y=2x", y, x)

writer.close()
```

* 启动 web 服务

```bash
# 会生成一个 web，监听 6007 端口，默认是 6006 端口
tensorboard --logdir=logs --port=6007
```


### add_image 使用

```python
import numpy as np
from PIL import Image
from torch.utils.tensorboard import SummaryWriter


# 会在 logs 文件夹下生成一个文件
writer = SummaryWriter("logs")


ant_img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/ants/0013035.jpg"

ant_img = Image.open(ant_img_path)
print(type(ant_img))

ant_img_array = np.array(ant_img)
print(type(ant_img_array))
print(ant_img_array.shape)

writer.add_image("ant", ant_img_array, global_step=1, dataformats='HWC')


bee_img_path = "/Users/lee/Workstation/DeepLearning/dataset/hymenoptera_data/train/bees/1092977343_cb42b38d62.jpg"

bee_img = Image.open(bee_img_path)
print(type(ant_img))

bee_img_array = np.array(bee_img)
print(type(bee_img_array))
print(bee_img_array.shape)

writer.add_image("bee", bee_img_array, global_step=1, dataformats='HWC')


writer.close()
```
