# 文本过滤和清理

import sys
import unicodedata

string = 'pýtĥöñ\fis\tawesome\r\n'


remap = {               # 转换表
    ord('\t'): ' ',
    ord('\f'): ' ',
    ord('\r'): None
}
string2 = string.translate(remap)   # 先去除各种转义等字符
print(string2)


remap2 = dict.fromkeys(c for c in range(sys.maxunicode) if unicodedata.combining(chr(c)))
string3 = unicodedata.normalize('NFD', string2) # 转换为使用组合字符形式
string4 = string3.translate(remap2)
print(string4)