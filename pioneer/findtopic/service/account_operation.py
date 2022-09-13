# -*- coding: utf-8 -*-
# @Time:    2019/11/26 4:33 PM
# @Author:  leeyoung
# @File:    account_operation.py
# @Content:


from constant.response_code import USER_EXISTS
from orm.account.models import Customer
from utils.response import CommonReturn


class AccountOperation(object):
    def __init__(self, user_id, username, password):
        self.user_id = user_id
        self.username = username
        self.password = password

    @property
    def user_exists(self):
        return Customer.objects.check_user_exists(self.username)

    def register_user(self):
        if self.user_exists:
            return CommonReturn(USER_EXISTS, '用户已经存在')

    def login(self):
        pass

    def logout(self):
        pass
