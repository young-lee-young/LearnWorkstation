# 固定的列数格式化文本

import textwrap

text = 'Reformat the single paragraph in text so it fits in lines of no more than width columns,' \
       ' and return a list of wrapped lines. By default, tabs in text are expanded with string expandtabs, ' \
       'and all other whitespace characters including newline are converted to space.' \
       'See TextWrapper class for available keyword args to customize wrapping behaviour'


print(textwrap.fill(text, 30))