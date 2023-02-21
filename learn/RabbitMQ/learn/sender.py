# -*- coding: utf-8 -*-
# @Time    : 2018-10-11 16:52
# @Author  : yao.li
# @Content : 消息发送者

import logging
import pika
import sys

logging.basicConfig(level=logging.INFO, format='%(asctime)s %(filename)s line:%(lineno)d %(levelname)s content: %(message)s', filename='sender.log')
logger = logging.getLogger(__name__)


class Sender(object):
    def __init__(self, message):
        self.connection = None
        self.channel = None
        self.message = message

    def main(self):
        self.connect_mq()
        self.send_message()
        self.close_mq()

    def send_message(self):
        # 发送消息，使用默认的exchange
        logger.info('>>>>>>>>>>>>>> start send message, the message is %s' % self.message)
        try:
            # properties将消息设置为持久化
            self.channel.basic_publish(exchange='', routing_key='hello', body=self.message, properties=pika.BasicProperties(delivery_mode=2,))
            logger.info('>>>>>>>>>>>>>>> send message success')
        except Exception as e:
            logger.info('>>>>>>>>>>>>>>> send message fail, the reason is %s', e.message, exc_info=True)

    def connect_mq(self):
        logger.info('>>>>>>>>>>>>>>> start connect rabbitmq')
        try:
            # 连接
            credentials = pika.PlainCredentials('leeyoung', 'leeyoung')
            self.connection = pika.BlockingConnection(pika.ConnectionParameters('60.205.177.93', credentials=credentials))
            self.channel = self.connection.channel()

            # 声明队列，queue队列名，durable将队列声明为持久化
            # 如果已经用durable=False生成了一个队列，再使用True会报错
            self.channel.queue_declare(queue='hello', durable=False)
            logger.info('>>>>>>>>>>>>>>>> connect rabbitmq success')
        except Exception as e:
            logger.error('>>>>>>>>>>>>>>> connect rabbitmq failed, the reason is %s', e.message, exc_info=True)

    def close_mq(self):
        self.connection.close()


if __name__ == '__main__':
    message = ','.join(sys.argv)
    sender = Sender(message)
    sender.main()
