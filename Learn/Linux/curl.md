# curl命令使用


### 请求

1. -X 请求方法
2. -H 请求头
3. -d 请求参数

```sh
curl -X POST 'http://10.13.2.133:6101/v1/resop/app/list/' -H "Content-Type:application/json" -d '{"customer_id":"CDS00916","user_id":"yi.pan","api_flag":1}'
```
