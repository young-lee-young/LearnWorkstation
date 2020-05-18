# -*- coding: utf-8 -*-
# @Time:    2020/2/10 4:45 PM
# @Author:  leeyoung
# @File:    redis_learn.py
# @Content:

import json
import pymysql
import redis
import tornado.ioloop
import tornado.web
from tornado.escape import json_decode


class DatabaseOperation:
    def __init__(self, host, username, password, database, port=3306):
        self.host = host
        self.port = port
        self.username = username
        self.password = password
        self.database = database

    def get_connect(self):
        return pymysql.connect(host=self.host, port=self.port, user=self.username, password=self.password, database=self.database)

    def execute_sql(self, sql):
        conn = self.get_connect()
        cursor = conn.cursor()

        cursor.execute(sql)
        result = cursor.fetchone()

        cursor.close()
        conn.close()
        return result


class RedisOperation:
    def __init__(self, host, port=6379, database=0):
        self.host = host
        self.port = port
        self.database = database

    def get_connect(self):
        return redis.Redis(host=self.host, port=self.port, db=self.database)

    def get_connect_pool(self):
        pool = redis.ConnectionPool(host=self.host, port=self.port, db=self.database)
        return redis.Redis(connection_pool=pool)


class MainHandler(tornado.web.RequestHandler):
    def post(self, *args, **kwargs):
        host = '127.0.0.1'
        username = 'root'
        password = 'leeyoung'
        database = 'testdb'
        operation = DatabaseOperation(host, username, password, database)

        data = json_decode(self.request.body)
        user_name = data.get('username')

        redis_operation = RedisOperation(host, database=2)
        redis_connect = redis_operation.get_connect_pool()

        result = redis_connect.get(user_name)

        if result:
            result = result.decode('utf-8')
        else:
            query_sql = "SELECT `password` FROM `users` WHERE username = '{}'".format(user_name)
            result = operation.execute_sql(query_sql)

        self.set_header('Content-Type', 'application/json; charset=UTF-8')

        info = {
            'username': user_name,
            'password': result
        }
        self.write(json.dumps(info))
        self.finish()


def make_app():
    return tornado.web.Application([(r"/test", MainHandler),])


if __name__ == '__main__':
    app = make_app()
    app.listen(9000)
    tornado.ioloop.IOLoop.current().start()
