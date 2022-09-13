# -*- coding: utf-8 -*-
# @Time:    2019/11/26 4:49 PM
# @Author:  leeyoung
# @File:    response.py
# @Content:


from constant.response_code import SUCCESS


class CommonReturn(object):
    def __init__(self, code_msg, message, data=None):
        self.code_msg = code_msg
        self.message = message
        self.data = dict() if data is None else data

    @property
    def code(self):
        return self.code_msg.get('code')

    @property
    def msg(self):
        return self.code_msg.get('msg')

    @property
    def is_success(self):
        if self.code == SUCCESS.get('code'):
            return True
        return False
