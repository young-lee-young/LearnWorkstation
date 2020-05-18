# time模块使用
# datetime模块使用
# random模块使用

import time

# time.time()返回的是当前时间的时间戳,(1970年以后经过的浮点秒数)
print(time.time())
# 1515316873.005469


# 将一个时间戳转换为当前时区的struct_time,参数是时间戳,无参的话显示的是当前时间
time1 = time.localtime(1)
print(time1)
# time.struct_time(tm_year=1970, tm_mon=1, tm_mday=1, tm_hour=8, tm_min=0, tm_sec=1, tm_wday=3, tm_yday=1, tm_isdst=0)
# 解释上面打印的内容,年、月、日、时、分、秒、星期、一年中的第几天、是否是夏令时


# 将一个时间戳转换为UTC(格林威治天文时间)时区的struct_time,参数是时间戳,无参的话是当前的格林威治时间
time2 = time.gmtime(1)
print(time2)
# time.struct_time(tm_year=1970, tm_mon=1, tm_mday=1, tm_hour=0, tm_min=0, tm_sec=1, tm_wday=3, tm_yday=1, tm_isdst=0)
# 解释上面打印的内容,年、月、日、时、分、秒、星期、一年中的第几天、是否是夏令时


# 将一个struct_time转换为时间戳,参数是struct_time,必须有参数
time3 = time.mktime(time.localtime())
print(time3)
# 1515318966.0


# 将一个struct_time转换为格式化时间,参数为struct_time,无参的话传入的是time.localtime()
time4 = time.asctime()
print(time4)
# Sun Jan  7 17:56:06 2018


# 将一个时间戳转换为格式化时间,参数是时间戳,无参的话传入的是time.time()
time5 = time.ctime(1)
print(time5)
# Thu Jan  1 08:00:01 1970

# 按指定格式输出时间,第一个参数是输出格式，第二个参数是要输出的时间（struct_time格式），如果没有第二个参数，传入time.localtime()
time6 = time.strftime('%Y-%m-%d %H:%M:%S',time.localtime())
print(time6)
# 2018-01-07 18:47:09


# 将一个指定格式的时间输出为struct_time格式,第一个参数是格式化的时间,第二个参数要和第一个参数的格式一致
time7 = time.strptime('2018-01-07 18:47:09','%Y-%m-%d %H:%M:%S')
print(time7)
# time.struct_time(tm_year=2018, tm_mon=1, tm_mday=7, tm_hour=18, tm_min=47, tm_sec=9, tm_wday=6, tm_yday=7, tm_isdst=-1)


# 睡眠3秒钟
time.sleep(1)
print('------------------')


from datetime import datetime

datetime1 = datetime.now()
print(datetime1)
print('------------------')