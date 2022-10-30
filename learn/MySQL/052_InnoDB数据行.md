### 变长列（数据在磁盘中存储的空间有可能长有可能短）

* 长度不固定的数据类型

VARCHAR、VARBINARY、BLOB、TEXT

* 占用空间大于 768字节 的 CHAR

* 变长编码（如 UTF-8）下的 CHAR


### 行溢出数据

InnoDB 每个数据页容量有限（16KB），导致字段长度也有限，当字段过大时，InnoDB 会使用行溢出机制，行溢出机制会把超长字段放入单独开辟的数据页


### 行记录格式（Raw Format）

InnoDB 行记录格式主要分为两个时代

* Redundant/Compact（Antelope 文件格式）

MySQL 5.0 之前的行记录格式，现在已经被淘汰

![行记录格式](052_Redundant.png)

1. 字段偏移列表：记录每个字段的相对位置，按照列的顺序逆序放置
2. Header：列数量、字段偏移列表的单位、下一行记录的指针等信息
3. RowID：没有可用主键时，使用 RowID 作为隐藏主键
4. TxID：Transaction ID，事务ID
5. Roll Pointer：回滚指针
6. 数据


* Dynamic/Compressed（Barracuda 文件格式）

![行记录格式](052_Compact.png)

1. 变长字段长度表：记录每个变长字段的长度，按照列的顺序逆序放置
2. NULL 标志为：指示行记录中的 NULL 值，每个 bit 代表一个字段
