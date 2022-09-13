# ElasticSearch Index操作


### 增删改查

* 查看所有索引

```sh
curl -X GET /_cat/indices?v
```

* 创建索引

```sh
# pretty表示结果以json格式返回
curl -X PUT /索引名?pretty
```

* 删除索引

```sh
curl -X DELETE /索引名
```
