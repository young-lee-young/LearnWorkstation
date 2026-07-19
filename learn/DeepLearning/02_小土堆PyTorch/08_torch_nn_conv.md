# pytorch 卷积操作


### 使用 

* 代码示例

```python
import torch
import torch.nn.functional as F


input = torch.tensor([[1, 2, 0, 3, 1],
                      [0, 1, 2, 3, 1],
                      [1, 2, 1, 0, 0],
                      [5, 2, 3, 1, 1],
                      [2, 1, 0, 1, 1]])

kernel = torch.tensor([[1, 2, 1],
                       [0, 1, 0],
                       [2, 1, 0]])


# conv2d input 需要 (batch, 输入通道, 高, 宽) 这样的数据，所以需要 reshape
input = torch.reshape(input, (1, 1, 5, 5))
# conv2d 卷积核需要 (输出通道, 对输入通道分组, 高，宽) 这样的数据
kernel = torch.reshape(kernel, (1, 1, 3, 3))


output = F.conv2d(input, kernel)
print(output)
# 输出：
# tensor([[[[10, 12, 12],
#           [18, 16, 16],
#           [13,  9,  3]]]])


output = F.conv2d(input, kernel, stride=2)
print(output)
# 输出：
# tensor([[[[10, 12],
#           [13,  3]]]])


output = F.conv2d(input, kernel, padding=1)
print(output)
# 输出：
# tensor([[[[ 1,  3,  4, 10,  8],
#           [ 5, 10, 12, 12,  6],
#           [ 7, 18, 16, 16,  8],
#           [11, 13,  9,  3,  4],
#           [14, 13,  9,  7,  4]]]])
```
