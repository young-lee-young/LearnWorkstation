# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

import json


# 写入json文件
class QuotesPipeline(object):
    def open_spider(self, spider):
        self.file = open('D:/test.txt', 'a')

    def process_item(self, item, spider):
        line = json.dump(dict(item)) + '\n'
        self.file.write(line)
        return item

    def close_spider(self, spider):
        self.file.close()