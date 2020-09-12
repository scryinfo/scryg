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
1. 命名，使用小写，加下划线
2. 命名使用有明确函义的英文单词
## 代码
1. 使用like查寻时，分二类处理
   1. 直接查询
   2. 编码like的关键字