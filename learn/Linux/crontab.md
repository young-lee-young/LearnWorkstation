# Crontab定时任务的使用


### Crontab安装及启动

* 检查服务是否启动

```
# Ubuntu中crontab的服务cron
service cron status

# 重启服务
service cron restart
```

* 列举crontab

```
crontab -l
crontab -l -u username
```

* 编辑crontab

```
crontab -e
crontab -e -u username
```

* 删除crontab

```
crontab -r
```


### 配置文件的编写

```
* * * * * command
1. 分钟，0-59
2. 小时，0-23
3. 日期，1-31
4. 月份，1-12
5. 星期，0-7（0和7表示星期日）
```

```
# 每天21:30
  30 21 * * * service httpd restart

# 每月1, 10, 30日的4:45
  45 4 1,10,30 * * service httpd restart

# 每月的1到10号4:45
  45 4 1-10 * * service httpd restart

# 每隔2分钟
  */2 * * * * service httpd restart
  1-59/2 * * * * service httpd restart

# 晚上11点到7点，每隔一小时
  0 23-7/1 * * * service httpd restart

# 每天18:00到23:00之间每隔30分钟
  0,30 18-23 * * * service httpd restart
  0-59/30 18-23 * * * service httpd restart
```


### crontab的配置文件

* crontab文件及log地址

文件在/var/spool/cron/crontabs/root
/etc/crontab（PATH可以添加第三方的程序）
/etc/cron.d下（系统配置文件的补充）


日志文件需要设置

```
sudo vim /etc/rsyslog.d/50-default.conf
cron.*		/var/log/cron.log（将前面的注释去掉）
sudo service rsyslog restart

# 还有
var/spool/mail
```


### 常见错误

1. 环境变量
2. 命令行双引号中使用%时，为加反斜线
3. 第三个和第五个域之间执行的是"或"操作
```
# 四月的第一个星期日早晨1时59分运行a.sh
59 1 1-7 4 0 /root/a.sh    # 第2个和第5个是或运算，这种写法就是1-7号加上4月的所有周日
59 1 1-7 4 * test 'data +\%w' -eq 0 && /root/a.sh
```
4. 分钟设置误用
每2小时执行一次
```
* */2 * * * service httpd restart	# 这种写法是错误的
0 */2 * * * service httpd restart	# 这种写法正确的
```


### 注意事项

* crontab是每次启动一个新进程
* 实现半分钟或者几十秒的操作使用sleep 30s;
