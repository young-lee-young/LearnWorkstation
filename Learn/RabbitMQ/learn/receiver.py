# -*- coding: utf-8 -*-
# @Time    : 2018-10-11 19:01
# @Author  : yao.li
# @Content : 消息接受者

import logging
import pika
import time

logging.basicConfig(level=logging.INFO, format='%(asctime)s %(filename)s line:%(lineno)d %(levelname)s content: %(message)s')
logger = logging.getLogger(__name__)


class Receiver(object):
    def __init__(self):
        self.connection = None
        self.channel = None

    def main(self):
        self.connect_mq()
        self.receive_message()
        # 阻塞进程，一直等待新消息，接受到新消息后调用回调函数
        self.channel.start_consuming()
        self.close_mq()

    def receive_message(self):
        # 公平调度
        self.channel.basic_qos(prefetch_count=1)
        # callback参数为回调函数，no_ack为是否开启响应，默认是开启
        self.channel.basic_consume(self.callback, queue='', no_ack=False)

    def callback(self, ch, method, properties, body):
        # 模拟耗时任务
        logger.info('>>>>>>>>>>>>>>> receive message is %s' % body)
        time.sleep(10)

        # 处理完消息后发送一个相应，队列收到处理成功的响应后再删除队列中响应的消息
        self.channel.basic_ack(delivery_tag=method.delivery_tag)
        logger.info('>>>>>>>>>>>>>>> handle message success')
        time.sleep(5)

    def connect_mq(self):
        logger.info('>>>>>>>>>>>>>>> start connect rabbitmq')
        try:
            # 连接
            credentials = pika.PlainCredentials('leeyoung', 'leeyoung')
            self.connection = pika.BlockingConnection(pika.ConnectionParameters('60.205.177.93', credentials=credentials))
            self.channel = self.connection.channel()

            # 声明队列
            self.channel.queue_declare(queue='hello', durable=False)
            logger.info('>>>>>>>>>>>>>>>> connect rabbitmq success')
        except Exception as e:
            logger.error('>>>>>>>>>>>>>>> connect rabbitmq failed, the reason is %s', e.message, exc_info=True)

    def close_mq(self):
        self.connection.close()


if __name__ == '__main__':
    receiver = Receiver()
    receiver.main()
