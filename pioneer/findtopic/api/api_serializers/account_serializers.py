# -*- coding: utf-8 -*-
# @Time:    2019/11/26 7:04 PM
# @Author:  leeyoung
# @File:    account_serializers.py
# @Content: 关于客户请求的参数序列化


from rest_framework import serializers


class CustomerSerializers(serializers.Serializer):
    user_id = serializers.CharField(max_length=64)
    username = serializers.CharField(max_length=64, required=True, allow_null=False, allow_blank=False)
    password = serializers.CharField(max_length=256, required=True)
    gender = serializers.CharField(max_length=8, required=False)
    address = serializers.CharField(max_length=256, required=False)
    phone = serializers.CharField(max_length=64, required=False)
    email = serializers.EmailField()
    is_valid = serializers.BooleanField(default=1)
    is_block = serializers.BooleanField(default=0)
