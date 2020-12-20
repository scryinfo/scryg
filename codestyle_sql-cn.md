[中文](./codestyle_sql-cn.md)  
# Code Style -- SQL
SCRYINFO
## 规则
1. 遵守设计原则
	1. 不使用业务数据作为主键
2. 更新数据时使用乐观锁
3. 禁止使用参数直接生成sql语句，使用sql的？或命名参数
4. 钱或token等类型，使用整数或字符串存储， 对于Postgressql数据库使用decimal或numeric类型支持差不多任意精度
## Name 
1. 命名，使用小写，加下划线，因为数据库关键字不区分大小写的
2. 命名使用有明确函义的英文单词
3. 不使用数据库的关键字或保留字命名，如不能使用from命名列名等。 
    [Postgres Keywords](https://www.postgresql.org/docs/12/sql-keywords-appendix.html)  
    [SQLite Keywords](https://www.sqlite.org/lang_keywords.html)    
    [MySQL Keywords](https://dev.mysql.com/doc/refman/8.0/en/keywords.html)     
    注：这三个文档中数据库的关键字都使用大写字母
4. 在使用ORM或与数据库表对应关系的对象时，不能使用数据库的关键字来命名
5. sql有大小写，有资料说，sql语句会被转换成大写再运行（引号内的字符值不会转换），所以为了减少转换的工作量，就建议sql语句写成大写

## 代码
1. 使用like查寻时，分二类处理
   1. 直接查询
   2. 编码like的关键字(pg:  https://www.postgresql.org/docs/12/functions-matching.html)
## 表设计
1. 每一个表都有一个主键字段，使用uuid varchar(36)类型
2. 数据库中不设计外键关联，这部分在代码中实现，主要是当有外键关联时可能会使用升级或备份失败（如果必须要外键关联，必须要讨论）
3. 数据库中基础信息类数据规则（如员工，分类等）
    1. 只能软件删除
    2. 保存enum类型的值在数据库中保存的，不允许改名，建议使用字符串类型，方便在数据库中查看
    3. 
4. 数据库开启慢查询日志   
    把超过10s的查询记录下来，分析并处理这些查询
5. 对程序中常用或并发的sql语句做 explain analyze分析，以优化sql语句
