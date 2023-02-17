# MVCC（Multiversion Concurrency Control，多版本并发控制）


### 实现

MVCC实现主要由3部分实现，分别为：隐藏字段、UndoLog、ReadView


### 隐藏字段

DB_TRX_ID（6字节）：创建这条记录或最后一次修改该记录的事务ID

DB_ROLL_PTR（7字节）：回滚指针，指向这条记录的上一个版本，用于配合 UndoLog，指向上一个旧版本

DB_ROW_ID（6字节）：隐藏主键，如果数据表没有主键，InnoDB会自动生成一个6字节的ROW_ID


### ReadView

* 作用

在版本链中选择哪一个版本的记录


* ReadView 生成时机

RC（读已提交）：每个 SQL 开始执行时生成的

RR（可重复读）：事务开始时生成


* 三个全局属性值

m_ids：用来维护 ReadView 生成时刻的正活跃（没有commit）的事务ID列表
min_trx_id：m_ids 列表中最小的事务ID
max_trx_id：ReadView 即将分配的下一个事务ID


* 可见性判断：选择可见版本

1. DB_TRX_ID < min_trx_id，能看到 TRX_ID 这个事务的修改，否则下一个判断
2. DB_TRX_ID >= max_trx_id，不能看到 TRX_ID 这个事务的修改，否则下一个判断
3. DB_TRX_ID 是否在后于事务中，如果不在，说明在生成ReadView时已经提交，可以看到；否则看不到这个事务的修改
4. 如果都看不到，则选择下一个版本链中的 TRX_ID 继续判断

总结就三点：

1. 不可以访问在我自己之后创建的事务ID
2. 不可以访问在我创建时活跃的事务ID
3. 版本链中最近的一个版本