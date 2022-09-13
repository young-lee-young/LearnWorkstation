# MongoDB基本命令


### mongodb命令行

```sh
# 进入命令行
mongo

# 退出命令行
exit
```


### 基本命令

* 查看版本

```sh
mongo --version
```


### 远程操作

```sh
# 远程连接
mongo 10.2.7.61:27017/库名 -u 用户名 -p 密码

# 远程导出
mongoexport -h 地址 --port 端口 -u 用户名 -p 密码 -c 库 -o 保存文件

# 远程导入
mongoimport -h 地址 --port 端口 -u 用户名 -p 密码 --upsert 保存文件
```


### 插入数据

```bash
client = pymongo.MongoClient('localhost')
db = client['databasename']
db.authenticate('用户名', '密码')
db['tablename'].insert(data)
```

### 查询数据
```
names = db.tablename.find()  # 查询所有数据，返回的是cursor对象，可以使用for循环获得数据
# <pymongo.cursor.Cursor object at 0x000001AAC1A55D68>
for name in names:
    print(name)  # name是字典，包括数据库自动生成的id
# {'_id': ObjectId('5a4863452ffd6f95280541a8'), 'name': 'liyao'}
# {'_id': ObjectId('5a4863452ffd6f95280541a9'), 'name': 'macheng'}
# {'_id': ObjectId('5a4863452ffd6f95280541aa'), 'name': 'liwei'}
# {'_id': ObjectId('5a4863452ffd6f95280541ab'), 'name': 'liuyafei'}
```
```
db.tablename.find_one()  # 查询第一条数据，返回的是dict对象
# {'_id': ObjectId('5a4863452ffd6f95280541a8'), 'name': 'liyao'}
```
```
name = db.tablename.find({'name': 'liyao'})  # 返回的是cursor对象，使用for循环获得数据
print(name)
# <pymongo.cursor.Cursor object at 0x0000016BE60E6DD8>
for info in name:
    print(info)
# {'_id': ObjectId('5a4863452ffd6f95280541a8'), 'name': 'liyao'}
```


### cmd端操作

* 用户类:

创建用户:

```
db.createUser(
    {
        user: 'username',
        pwd: 'userpwd',
        roles: [{role: '角色', db: '对其他库的作用'}]
    }
)
```
role可选项:
Read:允许用户读取指定数据库
readWrite:允许用户读写指定数据库
dbAdmin:允许用户在指定数据库中执行管理函数,如索引创建、删除,查看统计或访问system.profile
userAdmin:允许用户向system.users集合写入,可以找指定数据库里创建、删除和管理用户
clusterAdmin:只在admin数据库中可用,赋予用户所有分片和复制集相关函数的管理权限
readAnyDatabase:只在admin数据库中可用,赋予用户所有数据库的读权限
readWriteAnyDatabase:只在admin数据库中可用,赋予用户所有数据库的读写权限
userAdminAnyDatabase:只在admin数据库中可用,赋予用户所有数据库的userAdmin权限
dbAdminAnyDatabase:只在admin数据库中可用,赋予用户所有数据库的dbAdmin权限
root:只在admin数据库中可用,超级账号,超级权限
认证用户:
db.auth('username', 'userpwd')
查看用户:
show users
删除用户:
db.dropUser('username')
MongoDB中用户的角色说明
read角色数据库的只读权限,包括:collStats、dbHash、dbStats、find、killCursors、listIndexes、listCollections
readWrite角色数据库的读写权限,包括:collStats、convertToCapped、createCollection、dbHash、dbStats、dropCollection、createIndex、dropIndex、emptycapped、find、insert、killCursors、listIndexes、listCollections、remove、renameCollectionSameDB、update
dbAdmin角色数据库的管理权限,包括:collStats、dbHash、dbStats、find、killCursors、listIndexes、listCollections、dropCollection and createCollection on system.profile only
userAdmin角色数据库的用户管理权限,包括:changeCustomData、changePassword、createRole、createUser、dropRole、dropUser、grantRole、revokeRole、viewRole、viewUser
readAnyDatabase角色任何数据库的只读权限(和read相似)
readWriteAnyDatabase角色任何数据库的读写权限(和readWrite相似)
userAdminAnyDatabase角色任何数据库用户的管理权限(和userAdmin相似)
dbAdminAnyDatabase角色任何数据库的管理权限(dbAdmin相似


显示当前数据库:
db或者db.getName()
显示数据库的状态:
db.stats()
更改数据库的名字:
没有直接的方法,先复制,再删除旧的
db.copyDatabase('oldname', 'newname')
use oldname
db.dropDatabase()  # 删除数据库，删除数据库应该是root权限才可以（不确定）
修复当前数据库:
db.repairDatabase()
显示连接的地址:
db.getMongo()


* 数据表类:
显示所有集合:
show tables或者show collections
得到数据库里所有集合:
和上面最大的区别是返回的是一个列表
```
[ "test", "useinfo", "username" ]
```
db.getCollectionNames()
创建集合:
db.createCollection('表名')
db.createCollection('表名', {capped: true, autoIndexId: true, size: 10000, max: 100})
capped,Boolean类型,(可选)如果为true,则启用封闭的集合.上限集合是固定大小的集合,它在达到其最大大小时自动覆盖其最旧的条目.如果指定tru,则还需要指定size参数
autoIndexId,Boolean类型,(可选)如果为true,则在_id字段上自动创建索引.默认值为false
size,数字类型,(可选)指定上限集合的最大大小(以字节为单位).如果capped为true,那么还需要指定此字段的值
max,数字类型,(可选)指定上限集合中允许的最大文档数
显示所有集合的状态:
db.printCollectionStats()
显示一个集合的状态:
db.集合名字.stats()
引用集合:
db.集合名字或者db.getCollection('集合名字')
重命名集合:
db.集合旧名字.renameCollection('新名字')
删除集合:
db.集合名字.drop()


* 列类:
添加列:

删除列:
db.表名.update({}, {$unset: {'列名': ''}}, false, true)

更改列的名字:
db.表名.update({}, {$rename: {'旧列名字', '新列名字'}}, false, true)


* 查询:

查询所有数据:
db.表名.find()
查询一条数据:
db.表名.findOne()
查询不同数据:
db.表名.distinct('列名')  # 以此列为基准查询这列的不同数据
根据条件查询:
db.表名.find({'列名': '要查询的数据'})  # 一个条件查询
db.表名.find({'列名': '要查询的数据', '列名2': '要查询的数据'})  # 多条件查询
db.表名.find({'列名': {$gt: 10}})  # 大于
db.表名.find({'列名': {$lt: 10}})  # 小于
db.表名.find({'列名': {$gte: 10}})  # 大于等于
db.表名.find({'列名': {$lte: 10}})  # 小于等于
db.表名.find({'列名': {$gt: 10, $lt: 20}})  # 大于多少，小于多少
db.表名.find({'列名': / 要查询的数据 /})  # 查询列名里包含要查询的数据
db.表名.find({'列名': / ^ 要查询的数据 /})  # 查询列名里以要查询数据开头的数据
db.表名.find({'查询条件'}, {'列名': true})  # 根据条件只查询列名的数据，true代表只查这个列，false代表查除了这个列的数据
db.表名.find().sort({'列名': 1})  # 1代表是升序排序，-1代表降序排序
db.表名.find().skip(10)  # 查询10条以后的数据
db.表名.find().limit(10)  # 限制查询的条数
db.表名.find().count()  # 显示查询数据的条数
db.表名.find({$ or:[{'列名': '查询的条件'}, {'列名': '查询条件'}]})  # or查找


* 删除数据:

db.表名.remove({'列名': '删除的条件'})  # 删除指定条件数据
db.表名.remove({})  # 删除所有数据
