# 以不区分大小写的方式对文本进行查找和替换

import re

text = 'UPPER PYTHON, lower python, Mixed Python'
result = re.findall('python', text, flags=re.IGNORECASE)    # 忽略大小写模式匹配
print(result)