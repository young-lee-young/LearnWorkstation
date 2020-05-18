# -*- coding: utf-8 -*-
# @Time:    2019/11/26 2:19 PM
# @Author:  leeyoung
# @File:    urls.py
# @Content:


from rest_framework.routers import SimpleRouter
from .views import account


router = SimpleRouter()
router.register('account', account.AccountViewSet, base_name='customer')

urlpatterns = router.urls
