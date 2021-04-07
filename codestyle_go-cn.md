[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)  
[한국어](./codestyle_go-ko.md)  
[日本語](./codestyle_go-ja.md)  
# Code Style -- go
[Uber Go Style](https://github.com/uber-go/guide)
[Effective Go](https://golang.org/doc/effective_go.html)
[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
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
11. 验证开发中的代码，使用单元测试；研究一项目技术实现等使用demo
12. 如果要使用使用全局变量，给出足够的理由
13. 通用的小功能，经过讨论后写入 scryg 中
14. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由
15. 参考Effective Go中的建议  https://golang.org/doc/effective_go.html
16. 代码规则    
    *. 函数内部结构使用数据为主线，分为三大块：定义数据，生成数据，使用数据 
例子： 
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
## Name 
1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不使用数据库的关键字或保留字命名，如不能使用for来命名一个字段名
## 目录文件
1. 单元测试与源代码文件放在同一目录下面，如代码文件为 “server.go”，单元测试文件为 “server_test.go”
2. 所有的demo放入“ 仓库名/demo ” 目录中
3. 如果是框架或基础库，需要“仓库名/sample”
4. 所有项目使用包管理(go mod)
## 代码
1. 不要定义interface的指针，它本身就是一个胖指针 
2. for i, v := range str { // code block } 中v是复制，所以避免不必要的复制，可以只使用i来遍历；值得注意的是，在code block中对i的修改会在下一轮循环前被重置。
```go
str := "abc.def"
for i := range str {
	if str[i] == '.' {
    	i += 2
	}
	fmt.Println(i, string(str[i]))
}
```
3. 如果匿名函数（也称闭包）有使用到循环变量时，有两种方式解决  
	一通过传参数的方式，不使用循环变量  
	二定义一个新的变量  
4. channel如果为空，使用它时，不是panic，而是直接卡死
5. 读取已关闭的channel，可以正确读取到channel中的剩余值；如果channel为空，则会读取到该channel类型的空值，且v, ok := <- c中，ok为false
6. 判断一个channel关闭的方法是 _, ok := <-c ，注释它还读取了数据，如果channel中没有数它会wait（当然close的除外）。在1.10的版本之前一直没有提供直接判断channel已关闭的方法
7. select：如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
8. type T int  与 type T = int是不一样的， 前一个定义一个新的类型，后一个定义int的一个别名
9. 实现接口时，加上如下代码以确保实现接口的所有函数  
```go
var (
	_ interfaceName = (*interfaceImpl)(nil)
)
```
10. recover:   
    * 使用recover来捕获panic时，只能捕获当前 goroutine的panic。
    * 只有在defer函数的内部，调用recover才有用。
11. return 和 defer 的执行顺序，see https://github.com/googege/blog/blob/master/go/go/important/README.md
运行到return处，给返回值赋值，运行defer（defer之间是堆栈顺序，后进先出）。注意对返回值是否为同一变量（没有产生副本，是同一个），如果是那么在defer中的修改会影响到最后的返回值，下面是两个特殊的例子（更具体的内容参见网页）
``` go
//返回值为 1，不是2
func tt3() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return i
}
//返回的函数运行结果为 2， 不是1
func tt4() func() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return func() int {
		return i // 引用变量/同一变量
	}
}
//返回值为 13，不是12
func tt5() (num int) {
	defer func() {
		num++
	}()
	return 12 // 给返回值 num 赋值为12
}
//返回值为 1
func tt6() (num int) {
	defer func() {
		num++
	}()
	return
}
//返回值指向 1
func tt7() (*int) {
    num := 0
	defer func() {
		num++
	}()
	return &num
}
```
注意：如果defer后面只有一条语句，则其中的变量会立刻被赋值；如果defer后面是一个函数，则其中的变量会在执行时才被赋值。
``` go
func main() {
	var a int
	defer fmt.Println("Print a in defer : ", a)
	defer func() {
		fmt.Println("Print a in deferf: ", a)
	}()
	a++
	fmt.Println("Print a in main  : ", a)
}
```
14. go的参数传递，全部是值传递（不支持引用传递，少数语言如C++，C#支持）进入函数的参数都是副本
16. new 出来是指针类型，所以不能使用new来初始引用类型，一般不使用new，而使用“ &TypeName{...}”，如下  
``` go
var a = new([]int) //a 是 *[]int 类型
var a2 = make([]int, 0) //a2 是 []int 类型
```
17. error的建议处理方式
    see github.com/pkg/errors
    func New(message string) error //如果有一个现成的error，这时候有三个函数可以选择。
    func WithMessage(err error, message string) error //只附加新的信息
    func WithStack(err error) error //只附加调用堆栈信息
    func Wrap(err error, message string) error //同时附加堆栈和信息
18. 检查接口最终对象是否为空
```go
func IsNil(any interface{}) bool {
	fmt.Println()
	re := false
	if any != nil {
		v := reflect.ValueOf(any)

		if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			re = v.IsNil()
			if !re {
				for {
					fmt.Println(v.Type())
					v2 := v.Elem()
					if v2.Kind() != reflect.Ptr && v2.Kind() != reflect.Interface {
						break
					}
					re = v2.IsNil()
					if re {
						break
					}
					v = v2
				}
			}

		}
	}
	return re
}
```
19. interface{}与nil
interface{}为nill时它的类型与指向的对象都为nil
```go
var inter1 interface{} = nil  // == nil
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
err :  <nil>
err2:  *errors.errorString
false

//err 与 err2的类型都为 error类型，一个为nil值一个不为空， 这时他们的类型是不相同的
```
怎么取一个interface的Type
```go
type TypeInterface interface{}
t := reflect.TypeOf((*TypeInterface)(nil)).Elem()
//注：因为interface不能实例化，所以只有先拿到它的指针类型
```
20. slice copy, 如果size太小（不是容量），那么最多只复制size的内容，且不会出错
21. 在使用append向slice增加内容时，如果size没有超出容量，不会重新分配sclice，也就是说原slice的地址不变
22. slice中的两个冒号：对于v :=data\[ a : b : c\],a,b分别为上下界，c为容量  
```go
//clone 一个slice的最高效方法
clone := append(data[:0:0], data...)
//不要使用下面的方法
clone2 := append(data[:0:len(data)], data...)// 没有clone
//先 make后再 copy， 也慢，且语句很多

//合并两slice的最高效方法

//合并clone的建议方法（可以依据len 与 cap的不同值进行clone，如果有特别的性能要求才需要这么处理） , see the scryg/sutils/skit/MergeClone
clone := make([]type,0, len(a) + len(b))
clone = append(clone, a...)
clone = append(clone, b...)

//最高效的合并clone
alen, blen := len(a), len(b)
switch alen + blen {
case blen :
    clone = append(b[:0:0], b...)
case alen:
    clone = append(a[:0:0], a...)
default:
    clone = append(a[:0:alen], a...)//不会产生clone
    clone = append(clone, b...)//一定产生clone,因为 alen 与 blen 都不为零
}
```
特别注意：  
    *. 如果没有超出容量append不会新分配内存，slice常因为这个而出错  
    *. slice的第三个参数不能超过 cap的值， 不然运行时panic  (slice bounds out of range)
下面是错误的示例  
```go
a := make([]int,0,4)
b := a[:0:cap(a)]  //运行正确
b2:= a[:0:10] //运行 panic

data := []int{1}
errClone := append(data[:0:len(data)], data...)
errClone[0] = 6
// data[0] == errClone[0] 这时值相等，因为并没有clone， data与errClone指向同一内存

```
23. 判断两个函数签名相同 ConvertibleTo AssignableTo
24. mod管理依赖包时，要指定依赖的版本，如果直接依赖于master请说明充分的理由
25. 已经声明的变量v可以出现在”:=”声明中的条件：
    * 本次声明的v与已经声明的v处于同一个作用域中（如果v已经在外层作用域中声明过，则此次声明会创建一个新的变量）。
    * 初始化中与v的值的类型相同的值才能赋予v。
    * 此次声明中至少有一个变量时新声明得。
26. iota枚举器：
    * iota常量自动生成器，每隔一行，自动累加1。
    * iota遇到const，重置为0。
    * 可以只写一个iota，常量声明省略值时，默认和之前一个字面得值相同。
    * 如果在同一行，值都一样。
    * iota被中断之后必须显式恢复。
27. 常量表达式：除了移位运算符之外，如果二元运算符是不同类型的无类型常量，结果类型是靠后的一个。比如一个无类型的整数常量除以一个无类型的复数常量，结果是一个无类型的复数常量。
28. fallthrough：强制执行switch匹配之后的case，但是它不会判断下一条case的表达式的结果是true或者false。并且fallthrough不能再type switch中使用。
29. 不要在struct中定义没有名字的接口(embedding interface)
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
30. 编译约束（go build constraint / tag）：一行以```// +build```开头的注释，后面需要跟一个空行；可能出现在任何类型的文件中，
    但需要放在文件的开头，上面只能有空行和其他注释（具体到.go文件，则表示该部分代码需要位于“包声明（即package行）”之前）
    
    规则：  
    1. 符号—— ","：and； " "：or； "!"：not。  
       举例来说：```// +build linux,386 darwin,!cgo```表示```(linux AND 386) OR (darwin AND (NOT cgo))```
    1. 如果约束分为多行，则每行之间是“and”关系：  
       ```go
       // +build windows
       // +build cgo
       //     ↓
       // +build windows,cgo
       ```
    1. 如果文件名符合以下特殊匹配规则，视为该文件具有对应的隐式约束：
        1. *_GOOS
        1. *_GOARCH
        1. *_GOOS_GOARCH  
        
        举例来说：source_windows_amd64.go，在编译时，会应用隐式约束：GOOS=windows, GOARCH=amd64
    1. 可以通过```// +build ignore```忽略该文件的build约束
### 多线程（goroutines）
1. 退出/取消设计  
    * 如果线程会长时间运行，必须有退出/取消
    * 长时间运行的代码中，必须有退出/取消检查
    * 在任何形式的wait中加入退出/取消机制（参见定时任务实现）
    * 退出/取消功能的channel数据类型使用 struct{}，因为它的size为零
```go
   //一般的限出/取消
    cancel := make(chan struct{})//使用channel 实现
    go func() {
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
    go func() {
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
   go func() {//方式一使用channel
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
   
   cancelValue := int32(0)//方式二 使用变量
    go func() {
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
2. 线程安全的条件
    * 有多线程
    * 运行过程中数据会变化  
    注： 只有一个线程写，多个线程读，是有线程安全问题的
3. write copy，在函数中都使用临时变量，反写回去时，必须整体替换
4. 尽量不使用time.Sleep函数，因为在sleep的过程中，不能正常退出
5. 最多等待运行10秒  
带context的实现
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
go func() {
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
* 特别注意任务本身运行的时间，对定时时间的影响
* 使用 ticker时，一定要想好任务本身运行的时间是否确定，是否会大于定时的时间，这些情况下定义器是否还有意义，  
比如：定时时间为1秒，而任务本身运行需要1秒钟，那么相当于这个任务一直连继在工作，与直接使用一个for循环效果是一样的，这此就没有必要使用定时器了。
```go
//不计算定时任务本身运行时间
timer := time.NewTimer(time.Second * 10)
for {

    select {
    case <-cancel:
        //退出清理
        timer.Stop() //尽快清理timer
        return
    case <-timer.C:
        //do something
    }
    timer.Reset(time.Second * 10)
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
8. go中有Mutex不可重入（不能在同一线程中调用两次lock，这样会死锁），没有获取锁超时功能（要么获取到，要么wait），  
不能检查其它线程是否已经获取锁(只能调用lock方法获取)，这些功能在<=1.14.6时。

31. 其它

```