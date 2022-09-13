
"""
有问题,总是显示403错误
"""
from urllib.request import urlopen

class Urlopen:
    def __init__(self, url):
        self.url = url

    def open(self, **kwargs):
        return urlopen(self.url.format_map(kwargs))

yahoo = Urlopen('http://finance.yahoo.com/d/quotes.csv?s={names}&f={fields}')
for line in yahoo.open(names='IBM,AAPL,FB', fields='sllclv'):
    print(line.decode('utf-8'))


# 上面的类可以用一个函数来替换
def urlopen(url):
    def opener(**kwargs):
        return open(url.format_map(kwargs))
    return opener()

yahoo = urlopen('http://finance.yahoo.com/d/quotes.csv?s={names}&f={fields}')
for line in yahoo(names='IBM,AAPL,FB', fields='sllclv'):
    print(line.decode('utf-8'))