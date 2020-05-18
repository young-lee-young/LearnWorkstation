# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: http://doc.scrapy.org/en/latest/topics/item-pipeline.html

import pymongo


class MongoPipeline(object):
    def __init__(self, mongo_uri, mongo_auth, mongo_db):
        self.mongo_uri = mongo_uri
        self.mongo_auth = mongo_auth
        self.mongo_db = mongo_db

    @classmethod
    def from_crawler(cls, crawler):
        return cls(
            mongo_uri=crawler.settings.get('MONGO_URI'),
            mongo_auth=crawler.settings.get('MONGO_AUTH'),
            mongo_db=crawler.settings.get('MONGO_DATABASE')
        )

    def open_spider(self, spider):
        self.client = pymongo.MongoClient(self.mongo_uri)
        self.db_auth = self.client[self.mongo_auth]
        self.db_auth.authenticate('liyao', '123456')
        self.db = self.client[self.mongo_db]

    def close_spider(self, spider):
        self.client.close()

    def process_item(self, item, spider):
        self.db['user'].update({'url_token': item['url_token']}, {'$set': item}, True)
        return item