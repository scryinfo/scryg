[中文](./codestyle_rust-cn.md)  
# Code Style -- rust
SCRYINFO
## 规则
1. 遵守软件设计六大原则
	1. 开-闭原则(Open-Closed Principle, OCP)
	2. 里氏代换原则(Liskov Substitution Principle,常缩写为.LSP)
	3. 依赖倒置原则(Dependence Inversion Principle)
	4. 口隔离原则(Interface Segregation Principle, ISP)
	5. 迪米特法则(Law of Demeter LoD)又叫做最少知识原则(Least Knowledge Principle,LKP)
	6. 单一职责原则(Simple responsibility pinciple SRP)
2. 功能完成，不是在自己电脑上能运行，是要整个项目能正常运行部署
3. 把问题拿出，不要把它遗忘在开发的过程中，在代码中加入 todo 说明，添加到github的issues
4. 先思考后写代码，从命名开始
5. 处理每一个error，并记录到日志中（最终写入日志中）
6. 处理所有分支，特别出现的异常情况的分支（如，不应该出现的数据等，写入error日志）
7. 重要调用都需要写入 info日志
8. 直接对外提供服务接口，必须稳定，不能因为一个错误就让整个服务停止工作
9. 对外提供的接口，统一错误编号及错误信息
10. 定义函数时要考虑两个方面，一实现函数代码是否合理，二使用是否方便，是否容易出错
11. 验证开发中的代码，使用单元测试；研究技术实现等使用demo；给库提供例子使用sample或文档
12. 如果要使用使用全局变量，给出足够的理由
13. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由
## Name 
[遵守Rust的命名](https://rust-lang.github.io/api-guidelines/naming.html)
1. 所有源代码文件名，使用小写，加下划线
2. toml中的package或bin 命名使用下划线，不使用减号， 保持package name 与crate或lib 的名字一至
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
## 目录文件
1. 单元测试？
2. 所有的demo放入“ 仓库名/demo ” 目录中
3. 如果是框架或基础库，需要“仓库名/sample”
## 代码
### 代码提交前准备
1. fmt --> clippy --> cargo test --no-run。这三样通过后，才提交代码
2. 清除编译警告

### 单元测试
1. 在assert语句中如果使用了对象的字段或方法，建议打印整个对象。    
    * 检测Option为None    
    ```rust
    //let data = Option ....
    assert_eq!(true,v.is_none(),"{:?}", v)
    ```
    * 检测Result为Err 
    ```rust
    //let data = Result ....
    assert_eq!(true,v.is_err(),"{:?}", v)
    ```
### 其它
1. 使用确定大小的类型，如i32而不使用int类型
2. 函数入参优先使用&str代替String, 使用&[T]代替Vec
3. 尽量不使用panic!，如果需要使用，给出详细说明
4. 尽量不使用unswap 与 expect，正常情况下let与match，？ 就能很好的解决
5. 
## 库使用说明
### Rbatis
1.  [rbatis::tx::TxGuard] 是异步事务，如果需要不同连接之间数据的及时通知，请使用如下方式
```rust
//tx 只处理异常情况下，事务的rollback，所以会在事务提交成功后，调用 tx.manager = None; 阻止 [rbatis::tx::TxGuard]再管理事务
let mut tx = rb.begin_tx_defer(false).await?;
// .... 
rb.commit(&tx.tx_id).await?;
tx.manager = None;
```
