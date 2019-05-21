
[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)  
[한국어](./codestyle_go-ko.md)  
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
11. 验证开发中的代码，使用单元测试；研究一项目技术实现等使用demo
12. 如果要使用使用全局变量，给出足够的理由
13. 通用的小功能，经过讨论后写入 scryg 中
14. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由
15. 参考Effective Go中的建议  https://golang.org/doc/effective_go.html

## Name 

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词

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
(1)使用recover来捕获panic时，只能捕获当前 goroutine的panic。
(2)只有在defer函数的内部，调用recover才有用。

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

19. interface与nil
interface为nill时它的类型与指向的对象都为nil

```go
var inter1 interface{} = nil  // == nil
var inter2 interface{} = (*int)(nil) // != nil 因为类型值不为nil
fmt.Println(inter1 == nil)
fmt.Println(inter2 == nil)
//输出的结果为：
true
false
```
interface为nil与不为nil时的typeof是不相同的
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

20. slice copy, 如果size太小（不是容量），那么最多只复制size的内容，且不会出错

21. 在使用append向slice增加内容时，如果size没有超出容量，不会重新分配sclice，也就是说原slice的地址不变

22. slice中的两个冒号：对于v :=data\[ a : b : c\],a,b分别为上下界，c为容量  
    产生slice副本的正确方法是： c := v\[:0:0\]
    
23. 判断两个函数签名相同 ConvertibleTo AssignableTo

24. mod管理依赖包时，要指定依赖的版本，如果直接依赖于master请说明充分的理由

25. 已经声明的变量v可以出现在”:=”声明中的条件：
    (1)本次声明的v与已经声明的v处于同一个作用域中（如果v已经在外层作用域中声明过，则此次声明会创建一个新的变量）。
    (2)	初始化中与v的值的类型相同的值才能赋予v。
    (3)此次声明中至少有一个变量时新声明得。
    
26. iota枚举器：
    (1)iota常量自动生成器，每隔一行，自动累加1。
    (2)iota遇到const，重置为0。
    (3)可以只写一个iota，常量声明省略值时，默认和之前一个字面得值相同。
    (4)如果在同一行，值都一样。
    (5)iota被中断之后必须显式恢复。
    
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

30. 其它

