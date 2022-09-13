# -*- coding: utf-8 -*-
# @Time:    2019/11/26 4:38 PM
# @Author:  leeyoung
# @File:    customer.py
# @Content:


from django.db.models import Manager


class CustomerManager(Manager):
    def check_user_exists(self, username):
        if self.filter(username=username, is_valid=1).exists():
            return True
        return False
