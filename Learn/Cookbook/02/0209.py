# 将Unicode文本统一表示为规范形式

import unicodedata

s1 = 'Spicy Jalape\u00f1o'  # 全组成的
s2 = 'Spicy Jalapen\u0303o' # 组合而成的
print(s1)
print(s2)
print(s1 == s2) # 两者是不相等的

t1 = unicodedata.normalize('NFC', s1)   # NFC统一成全组成的,NFD统一为组合使用的字符
t2 = unicodedata.normalize('NFC', s2)
print(t1 == t2) # 两者是相等的