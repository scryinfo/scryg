[中文](./codestyle_rust-cn.md)

# Code Style -- rust

SCRYINFO

## 名词

函数(function)：由fn定义的函数  
方法(method)：是一种特殊的函数，第一个参数含self（与一个struct或trait关联的）  
关联函数(Associated functions)：Associated functions are functions associated with a type  
孤儿原则：trait与实现trait，有这样的要求。 impl的代码要么在trait所在的crate，要么在struct所在的crate。


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
    *. 函数内部结构使用数据主线，分为三大块：定义数据，生成数据，使用数据 例子：
   ```rust
   fn fun_name() -> Vec<i32>{
       //定义数据
       let mut data = Vec::new();
       //生成数据
       {}
       //使用数据
       return data;
   }
   ```

## Name

[遵守Rust的命名](https://rust-lang.github.io/api-guidelines/naming.html)  
[RFC0430 finalizing naming conventions](https://github.com/rust-lang/rfcs/blob/master/text/0430-finalizing-naming-conventions.md)  
常用单词[Words](./words_cn_en.md)

1. 所有源代码文件名，使用小写，加下划线
2. toml中的package或bin 命名使用下划线，不使用减号， 保持package name 与crate或lib 的名字一至
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不要翻译lifetime这个单词，因为它的通常翻译是“生命周期”，这个不能很准备的表达英文的意思，“生命周期”更适合与lifecycle对应
5. 优先选用as_与into_打头的方法，然后再是to_打头的方法。  
    std在实现时以as_/into_打头的方法，一般是没有代价的。而to_一般会做较多工作如内存分配等，如方法str::to_lowercase()转换为小写，一般会新分配内存。  
    str::as_bytes查看str的uft-8的字节，没有内存分配。String::into_bytes进入内部的Vec<u8>数据。  
    as 与 into很像， as转换类型查看，into是进入里面的类型查看。查看源代码String::to_bytes/into_bytes体会它们的区别
6. iter/iter_mut/into_iter，其中into_iter是生成独立的iter，它会消耗容器，也就是owened。
7. 使用一至的单词顺命名。  
    一般是动词-名词-错误。这里有一种特殊的情况是，是某一分类下内容，可以在动词前面增加一个分类的前缀如 Eth---，或Btc---等

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
let chars = "é".chars().collect::<Vec<_ > > ();
// U+00e9: 'latin small letter e with acute'
assert_eq!(vec!['\u{00e9}'], chars);
let chars2 = "é".chars().collect::<Vec<_ > > ();
// U+0065: 'latin small letter e',U+0301: 'combining acute accent'
assert_eq!(vec!['\u{0065}', '\u{0301}'], chars2);
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
    * 在c++中的引用，是一个变量的别名，他们就是同一个对象，没有产生新的内存或对象 rust的reference就是个特别的指针，与C++中的引用不是一个概念。 C++与C#支持C++的引用， java go dart
      rust等都不支持C++的引用，而这些语言中的所谓“引用”，只是类似指针 rust函数参数不支持C++的引用传参，因为rust在传参时使用的是指针，存放指针值的对象产生了副本 java go dart
      rust等语言也不支持引用传参，都是传值 ？不确定&self/&mut self是否也产生传数时的副本，从调用工具上看没有产生。

```rust

```

```C++

```

8. Box::new 使用了一个神奇的单词“box”，它的作用是把对象安全的移动到heap上面，它在编译时可能会作优化以减少内存上的不必须的复制移动 ？分配两次

```rust
let d = Box::new(Data::default ());
// Data这个对被分配了两次，一次是Data::default()，在stack上, 一次是Box::new，在heap
// 有没有方法直接分配到heap上面，且是default的？
// 网上参考，有人说编译可能会优化，让它只会分配一次，[see](https://github.com/rust-lang/rust/issues/53827)，[see2](https://stackoverflow.com/questions/31502597/do-values-in-return-position-always-get-allocated-in-the-parents-stack-frame-or/31506225#31506225)
// 这个问题与C++中的placement new情况是一样的，但是没有找到在指定内存上调用default的方法，分析default的返回值是一个对象，没有转入参数的地方，它只产生一个对象，而不能在一个对象上运行。
// 所以下面是解决方法的思路（这不是一个可靠的方法，只是为了说明思路，在实际的代码中不要这样使用）
struct Data {
    name: String,
}

impl Data {
    pub fn _init(&mut self) {
        self.name = "test".to_owned();
    }
}

impl Default for Data {
    fn default() -> Self {
        let mut d = Data { name: String::default(), };
        d._init();
        d
    }
}
//这不是一个可靠的方法，只是为了说明思路，在实际的代码中不要这样使用
let mut d = unsafe {
let ptr: * mut Data = alloc(Layout::new::< Data > ()) as _;
( * ptr)._init();
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
12. 析构函数简单可靠，不要执行不确定或费时费力或阻塞操作  
    析构函数主要是指像drop，close等方法，
13. 如果trait不希望被crate外实现，使用private::Sealed
```rust
pub trait InnerTrait: private::Sealed {
    //...
}
```    

14. derive与trait bound都可以实现时，优先使用derive，它使用更简单
```rust
// 优先使用这个
#[derive(Debug, PartialEq)]
struct Good<T> { /* ... */ }

// 不建议使用
#[derive(Debug, PartialEq)]
struct Bad<T: Debug + PartialEq> { /* ... */ }
```

### 多线程

1. Sync或Send只是告诉编译器是安全的，并不会做什么动着，保证安全。注意在实现时使用“unsafe”，说明需要我们自己写来保证线程的安全。  
   下面是一个反例

```rust
//这是错误代码
struct Data {
    c: Cell<i32>,
}

unsafe impl Sync for Data {}

unsafe impl Send for Data {}

impl Data {
    pub fn add(&self) {
        let t = self.c.get() + 1;
        self.c.set(t);
    }
}

let mut d = Arc::new(Data {
c: Cell::new(0),
});

let mut d2 = d.clone();
let t2 = spawn(move | | {
for i in 0..101 {
d2.add();
}
});
let mut d3 = d.clone();
let t3 = spawn(move | | {
for i in 0..101 {
d3.add();
}
});
for i in 0..101 {
d.add();
}
t2.join();
t3.join();
println!("len: {}", d.c.get());//输出的结果大部分情况都不是303

```
2. Send与Sync trait编译器会自动推断为struct实现这两个trait。  
    如果struct中的所有字段都是Send或Sync的，那么这个struct也是Send或Sync的。只有编译器推断为不是Send或Sync时，我们才会手动实现它。一定注释，手机实现时是unsafe的。

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

## 文档

1. 在文档中使用 Errors、Panics、Safety

```rust
/// # Errors
/// ....
/// # Panics
/// ....
/// # Safety
/// ....
/// # Examples
/// ....
```

[see](https://rust-lang.github.io/api-guidelines/documentation.html)

2. 包含指向相关内容的超链接，下面是例子

```rust
//!Just doc sample for the mod mylib::doc_
//!
//!
//! # Samples
//! ```
//! use mylib::doc_::Data;
//!
//! let modData = Data::new();
//! ```
//!

/// Just doc sample for the mylib::doc_::Data
///
///
/// # Samples
///
/// ```
/// use mylib::doc_::Data;
/// let data = Data::new();
/// ```
/// use the [`Data::new`]
///
/// use 2 the [crate::doc_::Data]
///
/// see the [`Vec`]
///
/// see1 the [Vec]
///
/// see2 the [std::vec::Vec]
///
pub struct Data {
    /// data's name
    /// see [Data::new]
    pub name: String,
    list: Vec<i32>,
}

impl Data {
    /// new struct [Data]
    /// see the field [Data::name]
    pub fn new() -> Self {
        Data {
            name: "".to_owned(),
            list: Vec::new(),
        }
    }
}
``` 

说明：

* [`Vec`]与[Vec]这两种连接方式都是可以的
* 可以使用[std::vec::Vec]/[crate::doc_::Data]全路径增加link
* 在lib.rs开头增加如下代码，会在Samples的右上角显示“Run”按钮

```rust
#![doc(html_playground_url = "https://play.rust-lang.org/")]
```

## 库使用说明

### Rbatis

1. [rbatis::tx::TxGuard] 是异步事务，如果需要不同连接之间数据的及时通知，请使用如下方式

```rust
//tx 只处理异常情况下，事务的rollback，所以会在事务提交成功后，调用 tx.manager = None; 阻止 [rbatis::tx::TxGuard]再管理事务
let mut tx = rb.begin_tx_defer(false).await?;
// .... 
rb.commit( & tx.tx_id).await?;
tx.manager = None;
```

## 参考

[The Rust Programming Language](https://doc.rust-lang.org/book/)
[The Rust Reference](https://doc.rust-lang.org/reference/index.html)
[The Rustonomicon/Rust 秘典](https://doc.rust-lang.org/nomicon/index.html)
[Rust API Guidelines](https://rust-lang.github.io/api-guidelines/)
[Rust by Example](https://doc.rust-lang.org/stable/rust-by-example/)
[The Cargo Book](https://doc.rust-lang.org/cargo/index.html)
[The rustdoc book](https://doc.rust-lang.org/rustdoc/index.html)
[Elegant Library APIs in Rust](https://deterministic.space/elegant-apis-in-rust.html)

