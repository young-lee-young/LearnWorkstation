import numpy as np

"""
这个方法的使用举个例子

假如有输入数据
[[ 1,  2,  3,  4],
 [ 5,  6,  7,  8],
 [ 9, 10, 11, 12],
 [13, 14, 15, 16]]

# 输出
[[ 1,  2,  5,  6],
 [ 2,  3,  6,  7],
 [ 3,  4,  7,  8],
 [ 5,  6,  9, 10],
 [ 6,  7, 10, 11],
 [ 7,  8, 11, 12],
 [ 9, 10, 13, 14],
 [10, 11, 14, 15],
 [11, 12, 15, 16]]

然后再和卷积核进行乘积运算
"""


def im2col(input_data, filter_h, filter_w, stride=1, pad=0):
    # N 表示数据的个数
    # C 通道数
    # H 图片的高度
    # W 图片的宽度
    N, C, H, W = input_data.shape

    # 算出输出的长和宽
    out_h = (H + 2 * pad - filter_h) // stride + 1
    out_w = (W + 2 * pad - filter_w) // stride + 1

    # 表示每个维度怎么扩，第 1 维和第 2 维不扩展，第 3 维和第 4 维扩展 pad 大小
    # 比如从 (1, 1, 4, 4)，假如 pad 是 1，那么扩展后 (1, 1, 6, 6)
    # mode='constant' 表示用常数填充，默认值是 0
    img = np.pad(input_data, [(0, 0), (0, 0), (pad, pad), (pad, pad)], mode='constant')

    # TODO 下面的代码我还没完全理解
    col = np.zeros((N, C, filter_h, filter_w, out_h, out_w))

    for y in range(filter_h):
        y_max = y + stride * out_h
        for x in range(filter_w):
            x_max = x + stride * out_w
            col[:, :, y, x, :, :] = img[:, :, y:y_max:stride, x:x_max:stride]

    col = col.transpose(0, 4, 5, 1, 2, 3).reshape(N * out_h * out_w, -1)
    return col
