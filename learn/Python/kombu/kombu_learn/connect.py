# -*- coding: utf-8 -*-
# @Time:    2019/7/30 下午2:07
# @Author:  leeyoung
# @File:    connect.py
# @Content: Kombu连接RabbitMQ


from kombu import Connection


class Connect(object):
    def __init__(self, url):
        self.url = url

    @property
    def connction(self):
        return Connection(self.url)
