# -*- coding: utf-8 -*-
# @Time:    2019/11/26 2:35 PM
# @Author:  leeyoung
# @File:    account.py
# @Content:

from rest_framework import viewsets
from rest_framework.decorators import list_route
from rest_framework.response import Response

from api.api_serializers.account_serializers import CustomerSerializers
from service.account_operation import AccountOperation


class AccountViewSet(viewsets.GenericViewSet):
    @list_route(methods=['POST'], url_path='register', serializer_class=CustomerSerializers)
    def register(self, request, serializer_data=None):
        json = {
            'name': 'lee'
        }
        return Response(json)

    def delete(self, request):
        pass

    def login(self, request):
        pass

    def logout(self, request):
        pass
