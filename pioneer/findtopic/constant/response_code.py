# -*- coding: utf-8 -*-
# @Time:    2019/11/26 4:55 PM
# @Author:  leeyoung
# @File:    response_code.py
# @Content: 返回对象常用量

"""
所有请求返回code由六位数字组成，前三位表示大模块，后三位表示具体问题
"""

SUCCESS = {'code': '000000', 'msg': '请求成功'}

# 请求类错误，前三位码是001
METHOD_ERROR = {'code': '001001', 'msg': '请求方法错误'}
PARAM_ERROR = {'code': '001002', 'msg': '参数错误'}

# 用户类错误，前三位码为002
USER_EXISTS = {'code': '002001', 'msg': '用户名已存在'}
