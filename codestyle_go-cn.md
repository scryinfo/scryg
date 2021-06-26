[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)
[한국어](./codestyle_go-ko.md)  
[日本語](./codestyle_go-ja.md)

# Code Style -- go

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
5. 处理每一个error，并记录到日志中
6. 处理所有分支，特别出现的异常情况的分支（如，不应该出现的数据等，写入error日志）
7. 重要调用都需要写入 info日志
8. 直接对外提供服务接口，必须稳定，不能因为一个错误就让整个服务停止工作
9. 对外提供的接口，统一错误编号及错误信息
10. 定义函数时要考虑两个方面，一实现函数代码是否合理，二使用是否方便，是否容易出错
11. 验证开发中的代码，使用单元测试；研究一项目技术实现等使用demo/单元测试
12. 如果要使用使用全局变量，给出足够的理由
13. 通用的小功能，经过讨论后写入 scryg 中
14. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由

## Name

常用单词[Words](./words_cn_en.md)

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不使用数据库的关键字或保留字命名，如不能使用from,index来命名一个字段
5. 命名就是最好的注释，当命名能清楚地表达含义时，不要加注释，这时的注释就是画蛇添足

## 目录文件

1. 单元测试与源代码文件放在同一目录下面，如代码文件为 “server.go”，单元测试文件为 “server_test.go”
2. 所有的demo放入“ 仓库名/demo ” 目录中
3. 如果是框架或基础库，需要“仓库名/sample”
4. 项目使用包管理(go mod)

## 代码

1. 代码规则 函数内部结构使用数据为主线，分为三大块：定义数据，生成数据，使用数据 例子：

```go
func funName(){
   //定义数据
   var data []int
   //生成数据
   {}
   //使用数据
   return data
}
```

2. for i, v := range str { // code block } 中v是复制出的一个副本，为了避免不必要的复制，可以只使用i来遍历；值得注意的是，在code block中对i的修改会在下一轮循环前被重置。

```go
str := "abc.def"
for i := range str {
   if str[i] == '.' {
    i += 2
   }
   fmt.Println(i, string(str[i]))
}
```

3. 如果匿名函数（也称闭包）中使用到循环变量时，这时不能直接使用循环变量，因为在多次循环过程中，循环变量只有一个，每次修改它的值。有两种方式解决

* 通过传参数，不在闭包中直接使用循环变量
* 定义一个新的变量

```go
{//下面是错误的使用方法
   wg := sync.WaitGroup{}
   wg.Add(3)
   for i := 0; i < 3; i++ {
      go func () {
         fmt.Print(i)
         wg.Done()
      }()
   }
   wg.Wait()
   //“0123”中的三个数的重复组合，而不是三个数字"012"的组合
}
{//下面是正确做法，会输出三个数字"012"的组合
   wg := sync.WaitGroup{}
   wg.Add(3)
   for i:= 0; i< 3; i++ {
      go func (d int){
         fmt.Print(d)
         wg.Done()
      }(i)
   }
   wg.Wait()
}
```
4. type T int 与 type T2 = int是不一样的， 前面一个定义一个新类型T，后一个定义T2为int的别名
5. go的参数传递，全部是值传递（不支持引用传递，少数语言如C++，C#支持）进入函数的参数都是副本。 go中的引用类型实际上是一种特殊的指针，当引用类型在传值时，是浅copy的副本（把指针地址复制到，而指向的值并没有复制），
   与传指针达到一样的效果，实质上还是传值。
6. new 出来是指针类型，一般不使用new，而使用“ &TypeName{...}”，如下
```go
var a = new([]int) //a 是 *[]int 类型
var a2 = make([]int, 0) //a2 是 []int 类型
```
7. 在使用map或slice时，尽量给出容量，这样减少不必要的内存分配及复制
8. slice copy, 如果size太小（不是容量），那么最多只复制size的内容，且不会出错
9. 在使用append向slice增加内容时，如果size没有超出容量，不会重新分配slice，也就是说原slice的地址不变
10. slice中的两个冒号：对于v :=data\[ a : b : c\],a,b分别为上下界，c为容量
11. clone slice的建议方法
```go
func TestSliceClone(t *testing.T){
   data := []int{1}
   //clone 一个slice的最高效方法， data为nil时也可以正确工作
   clone1 := append(data[:0:0], data...)
   clone2 := append([]int(nil), data...)
   //不要使用下面的方法，他们都不会复制
   clone3Not := append(data[:0:len(data)], data...)
   clone4Not := append(data[:0], data...)
   //先 make后再 copy， 且语句更多，且需要特别处理nil情况，所以不建议使用
   
   assert.NotSame(t, &data[0], &clone1[0])
   assert.Equal(t,data, clone1)
   assert.NotSame(t, &data[0], &clone2[0])
   assert.Equal(t,data, clone2)
   
   assert.Same(t, &data[0], &clone3Not[0])
   assert.Same(t,&data[0], &clone4Not[0])
}
```
12. 合并两个slice的方法
```go
func TestSliceMergeClone(t *testing.T){
   data1 := []int{1}
   data2 := []int{4}
   
   //合并两个slice，建议方法
   merge := make([]int, 0, len(data1) + len(data2))
   merge = append(merge, data1...)
   merge = append(merge, data2...)
   assert.Equal(t, data1, merge[0:len(data1)])
   assert.Equal(t, data2, merge[len(data1):])
   
   //合并两个slice，高效方法，由于代码比较多所以不建议使用，如果把MergeClone做成库，那么应该使用下面的方法
   merge = nil
   data1Len, data2Len := len(data1), len(data2)
   switch data1Len + data2Len {
   case data2Len:
   	merge = append(data2[:0:0], data2...)
   case data1Len:
   	merge = append(data1[:0:0], data1...)
   default:
   	merge = append(data1, data2...)
   }
   assert.Equal(t, data1, merge[0:len(data1)])
   assert.Equal(t, data2, merge[len(data1):])
}
```
MergeClone两个slice的高效方法，已经实现在库 scryg/sutils/skit/MergeClone 中
13. nil是一个有效的slice 方法append的第一个参数可以为nil len(nil) == 0， 优先使用var a1 []int，而不是a2 := []int{}。 a2是零长度的slice，len(a2) == 0，
    而a1是nil的slice
14. 优先使用 strconv 而不是 fmt，strconv性能更好
15. mod管理依赖包时，要指定依赖的版本，如果直接依赖于master请说明充分的理由
16. 已经声明的变量v可以出现在”:=”声明中的条件：  
   * 本次声明的v与已经声明的v处于同一个作用域中（如果v已经在外层作用域中声明过，则此次声明会创建一个新的变量）。
   * 初始化中与v的值的类型相同的值才能赋予v。
   * 此次声明中至少有一个变量时新声明得。
17. iota枚举器：  
   * iota常量自动生成器，每隔一行，自动累加1。
   * iota遇到const，重置为0。
   * 可以只写一个iota，常量声明省略值时，默认和之前一个字面得值相同。
   * 如果在同一行，值都一样。
   * iota被中断之后必须显式恢复。
18. 常量表达式：除了移位运算符之外，如果二元运算符是不同类型的无类型常量，结果类型是靠后的一个。比如一个无类型的整数常量除以一个无类型的复数常量，结果是一个无类型的复数常量。
19. fallthrough：强制执行switch匹配之后的case，但是它不会判断下一条case的表达式的结果是true或者false。并且fallthrough不能再type switch中使用。
20. 减少导出方法或变量。如果函数或变量在外部不使用，就不要导出（不使用大写打头）
21. 使用字段名初始化结构体，不要省略字段名
22. 内嵌字段/embedded field  

| " " | T       | *T      |
|:----|:--------|:--------|
| S   | v       | v and p |
| *S  | v and p | v and p |

v methods/v: receiver为value的方法集合  
p methods/p: receiver为pointer的方法集合  
S是结构体 T是内嵌字段  
总结为：两个都是Value时，只包含T中接收者为Value的方法；其中一个为指针时，包含T的接收者为value与pointer的方法（所有方法） 当T中的方法或字段与S中的名字（或多个不同类型T中的字段或方法名）有冲突时，要用全名使用

23. 少使用init()函数 不要依赖init()运行的先后顺序。在修改代码之后，init函数之间的运行先后顺序可能会变化 不要在init()函数中运行“重”的工作，如果文件或网络操作等耗时操作，它们会因环境变化而出错，且很难恢复
24. go代码如果没有通过main入口引入，是不会被编译进来。 代码必须要有引用才会被编译进可执行文件

### 反射

1. 判断两个函数签名相同 ConvertibleTo AssignableTo

### interface

1. 不要定义interface的指针，它本身就是一个胖指针
2. 实现接口时，加上如下代码以确保实现接口的所有函数

```go
var (
    _ interfaceName = (*interfaceImpl)(nil)
)
```
go没有实现接口的语法，只要接口中所有的方法都有实现就认为是实现的接口，换句话说，检查是否实现接口是检查接口，是检查是否接口中的所有方法签名都能找到。如果两个接口有相同的方法，他们只会有一份实现

3. receiver参数默认使用 pointer，如果要使用value给出足够的理由 因为value方式会产生副本，而go语言很容易产生一个副本，所以receiver的类型一般为pointer。 下面是effective go中的说明  
   [Receiver pointer vs value](https://golang.org/doc/effective_go#pointers_vs_values)  
   规则是：  
   * pointer可以调用 receiver pointer/value
   * value 只能调用 receiver value 原因是，如果value调用receiver pointer，那么value会产生一个副本 value copy来调用，这里receiver
     pointer中修改了值，但是修改的是value copy中的，不会对原来的value产生影响，这个错误很因隐藏，很难发现，所以go语言不允许这样的调用发生
   * “接口实现”是不区分receiver是pointer还是value类型
   * T：包含receiver 为pointer&value的方法集 T：包含receiver为value的方法集

5. interface{}的实现  
   [code see](https://github.com/golang/go/blob/master/src/runtime/runtime2.go)  

```go

type iface struct {
   tab  *itab
   data unsafe.Pointer
}
type eface struct {
   _type *_type
   data  unsafe.Pointer
}
// layout of Itab known to compilers
// allocated in non-garbage-collected memory
// Needs to be in sync with
// ../cmd/compile/internal/gc/reflect.go:/^func.dumptabs.
type itab struct {
   inter *interfacetype
   _type *_type
   hash  uint32 // copy of _type.hash. Used for type switches.
   _     [4]byte
   fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
// Needs to be in sync with ../cmd/link/internal/ld/decodesym.go:/^func.commonsize,
// ../cmd/compile/internal/gc/reflect.go:/^func.dcommontype and
// ../reflect/type.go:/^type.rtype.
// ../internal/reflectlite/type.go:/^type.rtype.
type _type struct {
   size       uintptr
   ptrdata    uintptr // size of memory prefix holding all pointers
   hash       uint32
   tflag      tflag
   align      uint8
   fieldAlign uint8
   kind       uint8
   // function for comparing objects of this type
   // (ptr to object A, ptr to object B) -> ==?
   equal func (unsafe.Pointer, unsafe.Pointer) bool
   // gcdata stores the GC type data for the garbage collector.
   // If the KindGCProg bit is set in kind, gcdata is a GC program.
   // Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
   gcdata    *byte
   str       nameOff
   ptrToThis typeOff
}
```

interface{}分为两种类型有方法的iface与无方法的eface，有两个字段一个是data，另一个是type/tab。 一看实现代码就很容易理解，两个interface{}相等的条件，就要结构体中的两个字段分别相等才可以
一些可能使用的方法： func convI2I(inter *interfacetype, i iface) (r iface)，从a interface{}到b interface{}转换

6. 赋值给interface{}类型（包含传参数时）

```go
var fat interface{}
fat = nil //interface{} nil
var i int = 0
fat = i  //interface{}|int 0
fat = &i //interface{}|*int 0

var fat2 interface{} = i
fat = fat2
```

当赋值给interface{}类型时  
   * 如果右边是非interface{},会自动把右边的值转换为interface{}类型
   * 如果右边是interface{}, 直接复制fat指针的值（非深复制）

7. 检查接口最终对象是否为空

```go
//check if the final object pointed by interface is empty
func IsNil(any interface{}) bool {
   re := true
   if any != nil {
      v := reflect.ValueOf(any)
      switch v.Kind() {
      case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
         re = v.IsNil()
      default:
         re = false
         return re
      }
      if !re {
         for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface { //如果为指针或interface类型，要检查指向的值
            v = v.Elem() //Ptr或Interface时，返回内部的值
            switch v.Kind() {
            case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
                re = v.IsNil()
            default:
               re = false
               return re
            }
            if re {
               break
            }
         }
      }
   }
   return re
}
```

8. interface{}与nil interface{}为nil时它的类型与指向的对象都为nil

```go
var inter1 interface{} = nil // == nil
var inter2 interface{} = (*int)(nil) // != nil 因为类型值不为nil
fmt.Println(inter1 == nil)
fmt.Println(inter2 == nil)
//输出的结果为：
true
false
```

interface{}为nil与不为nil时的typeof是不相同的

```go
var err error = nil
var err2 error = errors.New("")
fmt.Println("err : ", reflect.TypeOf(err))
fmt.Println("err2: ", reflect.TypeOf(err2))
fmt.Println(reflect.TypeOf(err) == reflect.TypeOf(err2))
//输出结果为：
//err:<nil>
//err2: *errors.errorString
//false

//err 与 err2的类型都为 error类型，一个为nil值一个不为空， 这时他们的类型是不相同的
```
怎么取一个interface的Type
```go
type TypeInterface interface{}
t := reflect.TypeOf((*TypeInterface)(nil)).Elem()
//注：因为interface不能实例化，所以只有先拿到它的指针类型
```

9. 不要在struct中定义没有名字的接口(embedding interface)

```go
package main

func main() {
    Call()
}

type Hi interface {
    HiName() string
}
type Hello struct {
    Hi
}

func Call() {
   hello := &Hello{}
   hello.HiName()
}
```

这段代码是可以编译通过的， 运行时panic。为什么没有实现接口 Hi 就编译通过了，因为嵌入struct中的Hi只是一个字段而已，且是没有名字的，完整调用这样的： hello.Hi.HiName()，hello.Hi的值为nil，所以运行时panic  
   * 增加出错的机会，编译通过而运行出错
   * 如果Hello真的实现了接口Hi，那么 hello.HiName调用的是自己的方法，而不是 hello.Hi.HiName，容易让人误解
   * struct中嵌入的struct与inerface都是一个字段， 而interface中嵌入的interface，是要求实现对应方法的

10. go "=="总结 go语言不支持运算符重载

类型 | == | ？
----|----|----
bool | |
int | |
float | |
复数||
string | 字符内容 |
pointer | 比较地址，而不是指向的值|
struct |比较所有字段|由字段确定
array | 比较所有的元素|
slice | | 不支持
map | | 不支持
channel |同类型比较|
interface|比较两个字段的值|
interface{}| 特殊的interface|

```go
var i int = 1
var i2 int = 1
var pi *int = &i
var pi2 *int = &i2
var pi3 *int = &i
fmt.Printf("pi == pi3: %v, pi == pi2: %v\n", pi == pi3, pi == pi2)
//pi == pi3: true, pi == pi2: false

ch1 := make(chan int)
ch2 := make(chan int)
ch3 := ch1
fmt.Printf("ch1 == ch2: %v, ch1 == ch3: %v\n", ch1 == ch2, ch1 == ch3)
//ch1 == ch2: false, ch1 == ch3: true
```

引用或指针类型可以与nil进行比较。其它的不能与nil比较 两个channel比较时，同一个channel才相等（一个make出来的才是相等的,或者同一个channel才相等）
"assert.Equal"中使用的是“deepValueEqual”与"=="并不等价

[see](https://golang.org/ref/spec#Comparison_operators)

### channel

op | nil | closed | normal
----|----|----|----
close|panic|panic|closed
send| block|panic|block or sent
recv|block|return now|block or received

send： “chan <- 0”  
recv: "<- chan"  
在go的实现代码中使用的是send/recv不是write/read，这可能为了说明channel是用于通过而不是读写 block：就是卡死，不动了

1. close nil或已经closed的channel，都会panic
2. channel如果为nil时，send/recv，不是panic，而是block
3. send 关闭的channel会panic
4. recv 闭关的channel会立刻返回，如果buffer中有值，v, ok := <-c，v为正常的值，特别注意ok为true。当buffer为空时，v类型的默认值（int为0，引用为nil）,特别注意ok为false

```go
//recv的关闭chan
func TestRecvClosedChan(t *testing.T){
   c := make(chan int, 1)
   c <- 1
   close(c)
   v, ok := <- c
   assert.Equal(t, 1, v)
   assert.Equal(t, true, ok) //特别注意,buffer非空且closed时 true。如果这时用ok来判断channel是否关闭那真是个空难
   v,ok = <- c
   assert.Equal(t, 0, v)
   assert.Equal(t, false, ok) //特别注意，buffer为空且closed时 false
}
```

5. 没有直接判断一个channel是否关闭的方法，这个方法_, ok := <-c 有副作且不准确。 副作用是会recv chnnel中的数据，如果没有数据时且非closed时，还会block
不准准确，如上面的代码，如果buffer非空且closed时，ok的值为true  
[Receive_operator](https://golang.org/ref/spec#Receive_operator) 中说的很清楚，ok只是判断通过是否成功，并不是channel是否关闭
下面说一下并发的情况
6. close时，receiving立即返回

```go
//关闭 receiving的channel
func TestCloseReceivingChan(t *testing.T){
   c := make(chan int)
   var wg sync.WaitGroup
   wg.Add(1)
   go func () {
      defer func () {
         if r := recover(); r != nil {
            t.Fatalf("CloseReceivingChan is panic")
            fmt.Println(r)
         }
      }()
      wg.Done()
      t := <- c
      fmt.Println(t)
   }()
   wg.Wait() //确定goroutine 已经运行，这里不要使用 channel实现，这不是channel的正常功能，性能也不如WaitGroup
   time.Sleep(1) //确定goroutine 已经运行
   close(c)
}

```

7. close时，send会panic

```go
//关闭 sending的channel
func TestCloseSendingChan(t *testing.T){
   c := make(chan int)
   var wg sync.WaitGroup
   wg.Add(1)
   go func () {
      defer func () {
         if r := recover(); r != nil {
            fmt.Println(r)
         }
      }()
      wg.Done()
      c <- 1
      t.Fatalf("CloseSendingChan is not panic")
   }()
   wg.Wait() //确定goroutine 已经运行，这里不要使用 channel实现，这不是channel的正常功能，性能也不如WaitGroup
   time.Sleep(1) //确定goroutine 已经运行
   close(c)
}
```

8. Safe and high performance to close the channel  
[see](https://go101.org/article/channel-closing.html)  
channel的三个操作send/recv/close，只要有一个是多线程的，就要对channel进行保护
9. receiving queue,sending queue, buffering queue are all fifo（先进先出）

注：在sending或receiving的过程，close
channel会发生什么？[建议查看代码与测试代码结合](https://github.com/golang/go/blob/master/src/runtime/chan.go)
来确定结果

### error

1. go主要有如下错误 errors.New fmt.Errorf 实现Error的interface github.com/pkg/errors:
如果没有什么特别的就使用 github.com/pkg/errors包。下面是它的主要函数   
   func New(message string) error //如果有一个现成的error，这时候有三个函数可以选择  
   func WithMessage(err error, message string) error //只附加新的信息  
   func WithStack(err error) error //只附加调用堆栈信息  
   func Wrap(err error, message string) error //同时附加堆栈和信息  
2. 不要使用panic做正常的错误处理  
panic仅当发生不可恢复的事情（例如：nil 引用）时，程序才必须 panic。当错误无法正常处理或再运行会产生更严重的问题时，才panic时。
3. recover:  
   * 使用recover来捕获panic时，只能捕获当前 goroutine的panic。
   * 只有在defer函数的内部，调用recover才有用。
4. return 和 defer 的执行顺序,   
   [see](https://github.com/googege/blog/blob/master/go/go/important/README.md)  
运行到return处，给返回值赋值，运行defer（defer之间是堆栈顺序，后进先出）。注意对返回值是否为同一变量（没有产生副本，是同一个），如果是那么在defer中的修改会影响到最后的返回值，下面是两个特殊的例子（更具体的内容参见网页）

```go
//返回值为 1，不是2
func tt3() int {
   var i = 0
   defer func () {
    i++
   }()
   i++
   return i
}
//返回的函数运行结果为 2， 不是1
func tt4() func () int {
   var i = 0
   defer func () {
    i++
   }()
   i++
   return func () int {
    return i // 引用变量/同一变量
   }
}
//返回值为 13，不是12
func tt5() (num int) {
   defer func () {
    num++
   }()
   return 12 // 给返回值 num 赋值为12
}
//返回值为 1
func tt6() (num int) {
   defer func () {
    num++
   }()
   return
}
//返回值指向 1
func tt7() (*int) {
   num := 0
   defer func () {
    num++
   }()
   return &num
}
```

注意：如果defer后面只有一条语句，则其中的变量会立刻被赋值；如果defer后面是一个函数，则其中的变量会在执行时才被赋值。

```go
func main() {
   var a int
   defer fmt.Println("Print a in defer : ", a)
   defer func () {
    fmt.Println("Print a in deferf: ", a)
   }()
   a++
   fmt.Println("Print a in main : ", a)
}
//output
//Print a in main : 1
//Print a in deferf: 1
//Print a in defer : 0
```

### 多线程（goroutines）

1. 退出/取消线程设计，线程退出是不确定的，不能保证百分百退出，线程退出没有完美方案
* 如果线程会长时间运行，必须有退出/取消
* 长时间运行的代码中，必须有退出/取消检查
* 在任何形式的wait中加入退出/取消机制（参见定时任务实现）
* 退出/取消功能的channel数据类型使用 struct{}，因为它的size为零
* go中没有强制退出goroutine的方法，只能正常退出，有一个例外是当main退出时，不会等待goroutine运行完成就会直接退出。  
  
[see](https://github.com/golang/go/issues/32610)

```go
//一般的限出/取消
cancel := make(<-chan struct{}) //使用channel 实现
go func () {
   for {
      //do something
      select {
      case <-cancel:
        //清理
      return
        // 其它的case
      }
   }
}()

ctx := context.Background()//使用 context 实现
go func () {
   for {
      select {
      case <-ctx.Done():
        //清理
      return
        //其它case
      }
   }
}()
//长时间运行，比如在一个for循环中，每10次检查一次是否需要退出（做到1秒内检查一次，可以依据实际的代码来提高或降低检查的频率）
go func () {//方式一使用channel
   const maxCheck = 10
   for count := 0; true; count++ {
      if count > maxCheck {
         count = 0
         select {
         case <-cancel:
            //清理
            return
         default:
            //do nothing
         }
      }
      //do something
   }
}()

cancelValue := int32(0) //方式二 使用变量
go func () {
   const maxCheck = 10
   for count := 0; true; count++ {
      if count > maxCheck {
         if 1 == atomic.LoadInt32(&cancelValue) {
            //清理
            return
         }
      }
      //do something
   }
}()
```

2. 数据竟争/线程安全的条件
* 有多线程/超过一个线程
* 运行过程中数据会变化
注： 只有一个线程写，多个线程读，是有线程安全问题的 使用-race参数检查是否有数据竟争，如“go run xxx -race”
3. write copy，在函数中都使用临时变量，反写回去时，必须整体替换
4. 尽量不使用time.Sleep函数，因为在sleep的过程中，不能正常退出
5. 最多等待运行10秒
带context的实现

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
go func () {
   _, err := bind.WaitMined(ctx, client, ts)
   if err != nil {
    log.Println(err)
   }
   cancel() //没有到时间就运行完成，主动调用
}()
select {
case <-ctx.Done():
    //增加退出线程处理
}
```

6. 定时任务实现
* 定时器程序属于多线程，会有数据竟争问题
* 特别注意任务本身运行的时间，对定时时间的影响
* 使用 ticker时，一定要想好任务本身运行的时间是否确定，是否会大于定时的时间，这些情况下定时器是否还有意义，
比如：定时时间为1秒，而任务本身运行需要1秒钟，那么相当于这个任务一直连继在工作，与直接使用一个for循环效果是一样的，就没有必要使用定时器了。

```go
//不计算定时任务本身运行时间
cancel := make(<-chan struct{})
timer := time.NewTimer(time.Second * 10) //如果interval = 0，那么立刻运行
for {
   select {
   case <-cancel:
      //退出清理
      timer.Stop() //尽快清理timer
      return
   case <-timer.C: //这里也可以使用 “time.After(interval)”简化代码，但是不过After函数中，每次都会 new 一个新的timer出来，所以不建议使用
    //do something
   }
   timer.Reset(time.Second * 10)
}

//计算任务本身运行的时间
cancel := make(<-chan struct{})
const seconds = 10
interval := time.Second * seconds
timer := time.NewTimer(interval) //如果interval = 0，那么立刻运行
for {
   select {
   case <-cancel:
      //退出清理
      timer.Stop() //尽快清理timer
      return
   case <-timer.C: //这里也可以使用 “time.After(interval)”简化代码，但是不过After函数中，每次都会 new 一个新的timer出来，所以不建议使用
      start := time.Now()
      //do something
      diff := seconds - int(time.Now().Sub(start).Seconds())
      if diff <= 0 { //任务运行已经超过定时器了，依据业务选择怎么处理，这里只是简单的把定时改为一秒
        diff = 1
      }
      interval = time.Second * time.Duration(diff)
   }
   timer.Reset(interval)
}

```

7. 线程安全性能先后顺序
   1. 没有线程安全问题的算法(看业务实际情况)
   2. atomic
   3. writecopy等解决(用内存换时间)
   4. 标准库实现的线程安全的类型或容器
   5. mutex(Once, RWMutex, Mutex, Cond)
   6. channel
   7. timer或ticker
8. go中的Mutex是不可重入，没有获取锁超时功能（要么获取到，要么wait）， 不能检查其它线程是否已经获取锁(只能调用lock方法获取)，这些功能在<=1.14.6时。
可重入锁/reentrant lock：同一线程中可以多次调用lock，而不发生死锁的。它还有一个名字叫“递归锁”，一般使用可重入锁，因为“可重入”名字的含义更容易理解。
从好的设计上来讲可重入锁是一种糟糕的设计，原理上是一个线程锁一次就好了，但由于实际函数调用层次等原因， 很难做到一个线程只锁一次，所以可以在代码中使用可重入锁，不过go的sdk中没有提供。
自旋锁/spin lock：使用循环的方式获取锁，cpu使用较多，在并发比较少时性能更好。同一个lock可以是自旋且可重入的，也可以是不可以重入的。  
[Discuss the reentrant in Experimenting with GO](https://groups.google.com/g/golang-nuts/c/XqW1qcuZgKg/m/Ui3nQkeLV80J)  

9. 不要依赖goroutine id或使用"thread local"
go sdk并不直接提供方法返回当前的goroutine id，所以不要尝试去获得它，并做一些依赖 go sdk并不支持thread local功能，不尝试使用这样的功能
10. 单例的线程安全

```go
//这个代码可能产生多个实例，当多个线程同时调用 DBConnection时，dbInstance的值都为nil，
//它们都会运行dbInstance = &sql.Open(...) 新建一个连接
var dbInstance *DB
func DBConnection() *DB {
   if dbInstance != nil {
    return dbInstance
   }
   dbInstance = &sql.Open(...)
   return dbInstance
}

//正确的实现
var dbOnce sync.Once
var dbInstance *DB
func DBConnection() *DB {
   dbOnce.Do(func () {
    dbInstance = &sql.Open(...)
   }
   return dbInstance
}
```

11. sync.Once是只运行一次，且等待运行完成。
等待运行完成，这个丰常容易被忽略，当多个线程同时调用 “once.Do(....) ”方法时，所有方法都会阻塞，等once的第一次调用完成之后再运行。下面是运行一次任务，但不等待的实现

```go
//只关闭一次channel，且不等待关闭完成。
stopChannel := make(chan bool)
closing := int32(0) //atomic没有操作bool的
   closeByAtomic := func () {
   if atomic.CompareAndSwapInt32(&closing, 0, 1) {
    close(stopChannel)
   }
}

go closeByAtomic() //多次并发关闭，是安全的
go closeByAtomic()
closeByAtomic()

//sync.Once 实现
stopChannel := make(chan bool)
once := sync.Once{} //atomic没有操作bool的
closeByOnce := func () {
   once.Do(func (){
    close(stopChannel)
   })
}

go closeByOnce() //多次并发关闭，是安全的
go closeByOnce()
closeByOnce()
//两次实现的区别是，sync.Once是确定完成的，而atomic是不确定的。sync.Once时，closeByOnce()方法返回，那么close函数一定运行完成了。
//而使用atomic实现时，closeByAtomic()方法返回，close函数并不确定是否运行或完成，可能运行完成，可能还没有运行，可能运行到一半
```

12. Mutex,Cond,Once,WaitGroup,Channel选择

   * 在不线程之间传递数据时，一个生产一个消费，使用channel
   * 在多个线程中共享数据时（反复多次读写），使用Mutex，这时数据会被反复使用
   * 只做一次，且等运行完成时，使用Once
   * 多线程之间相互通知时，使用Cond，可以反复多次使用。这里有点像事件，当发生什么事件时，然后做什么事情。
   比如当上课时间时，触发一个事件“打上课铃”，同学们收到事件通过“听到上课铃”（注这里是多人），然后就进入教室。
   * 一个任务需要多部分或分步骤完成时，使用WaitGroup。如果数量只有一时，使用Cond与WaitGroup效果类似，但还是建议使用WaitGroup。
   比如做一个实验要分5步完成，每一步完成时，都做一次记录，直接所有的步骤都完成时，整个实验完成

## 文档

1. 注释与代码同步
2. 如果命名不能完整表达含义，使用简洁的语言说明

## 编译选项

1. 编译约束（go build constraint / tag）：一行以``` // +build```开头的注释，后面需要跟一个空行；可能出现在任何类型的文件中，
但需要放在文件的开头，上面只能有空行和其他注释（具体到.go文件，则表示该部分代码需要位于“包声明（即package行）”之前）

规则：
   * 符号—— ","：and； " "：or； "!"：not。
   举例来说：```// +build linux,386 darwin,!cgo```表示```(linux AND 386) OR (darwin AND (NOT cgo))```
   * 如果约束分为多行，则每行之间是“and”关系：
```go
    // +build windows
    // +build cgo
    //
    // +build windows,cgo
```
   * 如果文件名符合以下特殊匹配规则，视为该文件具有对应的隐式约束：
      * *_GOOS  
      * *_GOARCH  
      * *_GOOS_GOARCH  

举例来说：source_windows_amd64.go ，在编译时，会应用隐式约束：GOOS = windows, GOARCH =amd64
    * 可以通过
```go
// +build ignore
```
忽略该文件的build约束

## 优化

1. BCE: Bounds Check Elimination   
[see1](https://go101.org/article/bounds-check-elimination.html)  
[see2](https://docs.google.com/document/d/1vdAEAjYdzjnPA9WDOQ1e4e05cYVMpqSxJYZT33Cqw2g/edit)  
2. 逃逸分析/escape-analysis  
[see1](https://github.com/golang/go/wiki/CompilerOptimizations#escape-analysis)  

## 安全

1. 在进行数字运算操作时，需要做好长度限制，防止外部输入运算导致异常
   * 确保无符号整数运算时不会反转
   * 确保有符号整数运算时不会出现溢出
   * 确保整型转换时不会出现截断错误
   * 确保整型转换时不会出现符号错误
   //todo 给出反例
2. 防止出现指针的循环引用  
   //todo 给出例子
3. 小心使用unsafe包
4. SQL
   * 使用参数邦定
   * 禁止使用参数拼接SQL语句
5. Web
   * 跨域资源共享CORS限制请求来源 //todo 使用例子
   * CSRF防护

### 库使用记要

#### github.com/stretchr/testify

1. assert.Equal 中使用的是“deepValueEqual”，比较的两个值是否相等。与"=="并不等价
2. assert.Same比较两个指针地址是否相等，

```go
func TestSameEqual(t *testing.T) {
   a1 := 10
   a2 := 10
   p1 := &a1
   p2 := &a2
   assert.Equal(t, p1, p2) //地址不同，但值相等，所以成立
   assert.NotSame(t, p1, p2) //比较指针地址是否相等，所以成立
}
```

#### json

1. 自定义json的MarshalJSON函数时，小心receiver的类型。先看代码

```go
type JsonData struct {
    name string
}

func (c *JsonData) MarshalJSON() (bytes []byte, err error) {
    return json.Marshal(c.name)
}

func (c *JsonData) UnmarshalJSON(bytes []byte) error {
    return json.Unmarshal(bytes, &c.name)
}

func TestJsonData(t *testing.T) {
   d := JsonData{name: "test"}
   bytes, err := json.Marshal(d)
   assert.Equal(t, nil, err)
   var d2 JsonData
   err = json.Unmarshal(bytes, &d2)
   assert.Equal(t, nil, err) //has error，为什么？
   assert.Equal(t, d, d2) //not eq， 有两种方法，可以更正结果，一种是增加一个符号“&”,一种是删除一个符号“*”
}
```

以上代码中Unmarshal的结构不是预期的结果，原因是“func (c *JsonData) MarshalJSON() (bytes []byte, err error)”中的receiver是pointer。
这样如果在使用时如果没有使用pointer类型那么这个方法就不会生效。终上建议如下：
   * MarshalJSON方法的receiver不要使用pointer，如果要使用给出足够的理由及详细的文档说明
   * 自定义实现MarshalJSON与UnmarshalJSON的，在test中需要包含指针及非指针的测试用例

### 参考

[Effective Go](https://golang.org/doc/effective_go.html)  
[The Go Programming Language Specification](https://golang.org/ref/spec)  
[The Go Memory Model](https://golang.org/ref/mem)  
[Go Keywords](https://golang.org/ref/spec#Keywords)  
[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)  
[Go安全指南](https://github.com/Tencent/secguide/blob/main/Go%E5%AE%89%E5%85%A8%E6%8C%87%E5%8D%97.md)  
[Uber Go Style](https://github.com/uber-go/guide)  
[Postgres Keywords](https://www.postgresql.org/docs/12/sql-keywords-appendix.html)  
[SQLite Keywords](https://www.sqlite.org/lang_keywords.html)  
[MySQL Keywords](https://dev.mysql.com/doc/refman/8.0/en/keywords.html)  

