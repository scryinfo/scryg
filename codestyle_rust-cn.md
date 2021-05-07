[中文](./codestyle_rust-cn.md)  
# Code Style -- rust
SCRYINFO
## 说明
函数(function)：由fn定义的函数
方法(method)：是一种特殊的函数，第一个参数含self（与一个struct或trait关联的）
关联函数(Associated functions)：Associated functions are functions associated with a type
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
14. 代码规则    
    *. 函数内部结构使用数据主线，分为三大块：定义数据，生成数据，使用数据 
       例子： 
   ```rust
   fn fun_name() -> Vec<i32>{
       //定义数据
       let mut data = Vec::new();
       //生成数据
       {}
       //使用数据
       return data
   }
   ```
## Name 
[遵守Rust的命名](https://rust-lang.github.io/api-guidelines/naming.html)
1. 所有源代码文件名，使用小写，加下划线
2. toml中的package或bin 命名使用下划线，不使用减号， 保持package name 与crate或lib 的名字一至
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不要翻译lifetime这个单词，因为它的通常翻译是“生命周期”，这个不能很准备的表达英文的意思，“生命周期”更适合与lifecycle对应
## 目录文件
1. 单元测试？
2. 所有的demo放入“ 仓库名/demo ” 目录中
3. 如果是框架或基础库，需要“仓库名/sample”
## 代码
1. 使用确定大小的类型，如i32而不使用int类型
2. 函数入参优先使用&str代替String, 使用&[T]代替Vec
3. 尽量不使用panic!，如果需要使用，给出详细说明
4. 尽量不使用unswap 与 expect，正常情况下let与match
5. 个别字符看起是一个char，可能它是两个char。在使用时，容易产生混淆的字符，加上注释以示区别
```rust
let chars = "é".chars().collect::<Vec<_>>();
// U+00e9: 'latin small letter e with acute'
assert_eq!(vec!['\u{00e9}'], chars);
let chars2 = "é".chars().collect::<Vec<_>>();
// U+0065: 'latin small letter e',U+0301: 'combining acute accent'
assert_eq!(vec!['\u{0065}','\u{0301}'], chars2);
```
[see](https://doc.rust-lang.org/stable/std/primitive.char.html)
6. print!或println!每次调用时会锁定stdout，如果连续使用且有性能要求时，可以手动锁定，也可以增加buffer，下面是手动锁的例子
```rust
use std::io::Write;
let mut stdout = std::io::stdout();
let mut lock = stdout.lock();
for line in lines {
    writeln!(lock, "{}", line)?;
}
```
[see](https://poly000.github.io/perf-book-zh/io_zh.html)
### Pointer/Reference
7. Reference是一种特殊的指针。
   * 它由编译器来保证reference的有效性  
   * 在Rust的安全代码中，所有的引用都是有效的  
   * 在c++中的引用，是一个变量的别名，他们就是同一个对象，没有产生新的内存或对象
     rust的reference就是个特别的指针，与C++中的引用不是一个概念。
     C++与C#支持C++的引用，
     java go dart rust等都不支持C++的引用，而这些语言中的所谓“引用”，只是类似指针
     rust函数参数不支持C++的引用传参，因为rust在传参时使用的是指针，存放指针值的对象产生了副本
     java go dart rust等语言也不支持引用传参，都是传值
     ？不确定&self/&mut self是否也产生传数时的副本，从调用工具上看没有产生。
```rust

```
```C++

```
8. Box::new 使用了一个神奇的单词“box”，它的作用是把对象安全的移动到heap上面，它在编译时可能会作优化以减少内存上的不必须的复制移动
   ？分配两次
```rust
let d = Box::new(Data::default());
// Data这个对被分配了两次，一次是Data::default()，在stack上, 一次是Box::new，在heap
// 有没有方法直接分配到heap上面，且是default的？
// 网上参考，有人说编译可能会优化，让它只会分配一次，[see](https://github.com/rust-lang/rust/issues/53827)，[see2](https://stackoverflow.com/questions/31502597/do-values-in-return-position-always-get-allocated-in-the-parents-stack-frame-or/31506225#31506225)
// 这个问题与C++中的placement new情况是一样的，但是没有找到在指定内存上调用default的方法，分析default的返回值是一个对象，没有转入参数的地方，它只产生一个对象，而不能在一个对象上运行。
// 所以下面是解决方法的思路（这不是一个可靠的方法，只是为了说明思路，在实际的代码中不要这样使用）
struct Data{
    name: String,
}
impl Data{
    pub fn _init(&mut self){
        self.name = "test".to_owned();
    }
}
impl Default for Data{
    fn default() -> Self {
        let mut d = Data{name:String::default(),};
        d._init();
        d
    }
}
//这不是一个可靠的方法，只是为了说明思路，在实际的代码中不要这样使用
let mut d = unsafe {
    let ptr:*mut Data = alloc(Layout::new::<Data>()) as _;
    (*ptr)._init();
    Box::from_raw(ptr)
};
println!("{}", d.name);
```
    
9. 如果使用Pin<T>，给出足够的原因。Pin的作用是阻止使用可修改指针（&mut T）.也就是对象不可以被移动，在std中使用在在future的pull函数，解决自己引的问题。
10. trait Drop是一个trait，在对象的lifetime结束时（释放内存之前），会自动调用它的drop函数，这是编译器完成的。  
    * 对于drop来说它并不释放内存，它只是在释放内存前调用的一个函数。
    * 在drop函数实现中一般会释放内存或关闭文件等清理工作，但不要弄混了，drop只是一个函数，怎么实现都是可以的。
    * drop函数一般是自动调用的，也可以手动调用，也可以阻止自动调用
```rust

```
11. std::Vec实现说明  
    * set_len是不安全函数，要小心使用
    * 
```rust

```

### 代码提交前准备
1. fmt --> clippy --> cargo test --no-run。这三样通过后，才提交代码
2. 清除编译警告

### 单元测试
1. 在assert语句中如果使用了对象的字段或函数，建议打印整个对象。    
    * 检测Option为None    
    ```rust
    //let data = Option ....
    assert_eq!(true,v.is_none(),"{:?}", v);
    ```
    * 检测Result为Err 
    ```rust
    //let data = Result ....
    assert_eq!(true,v.is_err(),"{:?}", v);
    ```

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
