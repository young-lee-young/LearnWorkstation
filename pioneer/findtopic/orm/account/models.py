# -*- coding: utf-8 -*-
# @Time:    2019/11/26 3:20 PM
# @Author:  leeyoung
# @File:    models.py
# @Content:


from django.db import models

from orm.account.managers.customer import CustomerManager
from utils import create_uuid


class Customer(models.Model):
    GENDER = (
        ('male', '男'),
        ('female', '女')
    )

    id = models.CharField(max_length=64, primary_key=True, default=create_uuid(), help_text='客户id')
    username = models.CharField(max_length=64, unique=True, help_text='用户名')
    password = models.CharField(max_length=256, help_text='密码')
    gender = models.CharField(max_length=8, choices=GENDER, blank=True, null=True, help_text='性别')
    address = models.CharField(max_length=256, blank=True, null=True, help_text='地址')
    phone_num = models.CharField(max_length=16, blank=True, null=True, help_text='手机号')
    email = models.EmailField(max_length=64, blank=True, null=True, help_text='密码')
    is_valid = models.BooleanField(default=1, help_text='有效性')
    is_block = models.BooleanField(default=0, help_text='是否锁定')
    create_time = models.DateTimeField(auto_now_add=True, help_text='创建时间，暨用户注册时间')
    update_time = models.DateTimeField(auto_now=True, help_text='更新时间')

    objects = CustomerManager()

    def __str__(self):
        return self.username

    class Meta:
        db_table = 'account_customer'
