# -*- coding: utf-8 -*-
# @Time:    2019/7/30 下午2:14
# @Author:  leeyoung
# @File:    main.py
# @Content:


import time
from connect import Connect
from kombu import Producer


BROKER_TYPE = {
    'rabbitmq': 'amqp',
    'redis': 'redis'
}


class KombuOperation(object):
    def __init__(self, broker_type, username, password, host, port):
        self.broker_type = broker_type
        self.username = username
        self.password = password
        self.host = host
        self.port = port
        self._connection = None

    @property
    def broker_url(self):
        url = '{}://{}:{}@{}:{}//'.format(self.broker_type, self.username, self.password, self.host, self.port)
        return url

    @property
    def connection(self):
        if not self._connection:
            connect = Connect(self.broker_url)
            self._connection = connect.connction
        return self._connection

    def main(self):
        self.produce()

    def produce(self):
        with self.connection.channel() as channel:
            producer = Producer(channel)
            publish_msg = {'username': 'leeyoung', 'password': 'leeyoung'}
            producer.publish(publish_msg)


if __name__ == '__main__':
    broker_type = BROKER_TYPE.get('rabbitmq')
    kombu_operation = KombuOperation(broker_type, 'guest', 'guest', 'localhost', 5672)
    kombu_operation.main()
