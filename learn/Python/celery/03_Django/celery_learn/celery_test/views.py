# -*- coding: utf-8 -*-

from django.http import JsonResponse
from .tasks import TestTask


def do_task(request):
    print('>>>>>>>>>>>>>>> start request')
    TestTask.delay()
    print('>>>>>>>>>>>>>>> end request')
    result = {
        'status': 2000,
        'msg': 'success'
    }
    return JsonResponse(result)
