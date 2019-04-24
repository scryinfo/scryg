# Code Style

## 规则

1. 功能完成，不是在自己电脑上能运行，是要整个项目能正常运行部署
2. 把问题拿出，不要把它遗忘在开发的过程中，在代码中加入 todo 说明
3. 先思考后写代码，从命名开始
4. 处理每一个error，并记录到日志中
5. 处理所有分支，特别出现的异常分支（如，不应该出现的数据等，写入error日志）
6. 直接对外提供服务接口，必须稳定，不能因为一个错误就让整个服务停止工作
7. 在对外提供的接口中，统一错误编号及提示
8. 定义函数时要考虑两个方面，一实现函数代码是否合理，二使用是否方便，是否容易出错
9. 重要调用都需要写入 info日志
10. 验证开发中的代码时，使用单元测试；在研究一项目技术实现等使用demo
11. 如果要使用使用全局变量，给出足够的理由，因为它很难测试
12. 参考Effective Go中的建议  https://golang.org/doc/effective_go.html

## Name 

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词

## 目录文件
1. 单元测试与源代码文件放在同一目录下面，如代码文件为 “server.go”，单元测试文件为 “server_test.go”
2. 所有的demo放入“ 仓库名/demo ” 目录中

## 代码
1. 不要定义interface的指针，它本身就是一个胖指针  
引用类型一般实现 value method，值类型一般实现pointer method

2. for range 中是复制，所以避免不必要的复制，可以使用索引来遍历
3. 在循环中，使用匿名函数（也称闭包）时，如果使用到循环变量，一定要注意，循环变量只有一份实例，他会一直接变化。如果需要记录下来当前的值或索引等，请另外定义变量
4. defer 是在函数退出前运行的
5. 使用recover来捕获panic时，只能捕获当前 goroutine的panic
6. 在使用append向slice增加内容时，如果容量没有超出，返回的地址还有原来的
7. 在使用 "x, err := ..."时，如果与err在同一个｛｝内有同一变量，err不会新定义一个变量
8. channel如果为空，使用它时，不是panic，而是直接卡死
9. channel关闭后，读取出来的数据为 “0”且立刻返回
10. 判断一个channel关闭的方法是 ok,_ := <-c ，这个方法实际上是一个读取，如果channel中没有数它会wait，如果有数据会把它读取出来。在1.10的版本之前一直没有提供直接判断channel已关闭的方法
11. type T int  与 type T = int是不一样的， 前一个定义一个新的类型，后一个定义int的一个别名
12. 实现接口时，加上如下代码以确保实现接口的所有函数  
var (
	_ interfaceName     = (*interfaceImpl)(nil)
)
13. return 和 defer 的执行顺序，see https://github.com/googege/blog/blob/master/go/go/important/README.md

运行到return处，给返回值赋值，运行defer（defer之间是堆栈顺序，后进先出）。注意对返回值是否为同一变量（就是没有产生副本，是同一个），如果是那么在defer中的修改会影响到最后的返回值，下面是两个特殊的例子（更具体的内容参见网页）
``` go
package main
import (
	"fmt"
)
func main() {
	fmt.Println("tt3 return :", tt3())
	fmt.Println("tt4 return :", tt4()())
	fmt.Println("tt5 return :", tt5())
}
func tt3() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt3: ", i)
		i++
	}()
	i++
	return i
}
func tt4() func() int {
	var i = 0
	defer func() {
		fmt.Println("defer tt4:", i)
		i++
	}()

	i++
	return func() int {
		return i // 引用变量。
	}
}
func tt5() (num int) {
	defer func() {
		fmt.Println("defer tt5", num)
		num++
	}()
	return 12 // 给返回值 num 赋值为12
}

//////
func tt6() (num int) {
	defer func() {
		fmt.Println("defer tt5", num)
		num++
	}()
	return
}

func tt6() (*int) {
    num := 0
	defer func() {
		fmt.Println("defer tt5", num)
		num++
	}()
	return &num
}


```
下面是输出结果
```
defer tt3:  1
tt3 return : 1
defer tt4: 1
tt4 return : 2
defer tt5 12
tt5 return : 13
```
14. 无法取map的value的地址，原因是它在变化
15. go的参数传递，全部分都是值传递（不支持引用传递的，少数语言如C++，C#是支持的）   
    进入函数的参数都是一个副本，对于指针，是使用一副本来存放指针的地址，指针所指向的对象 并没有产生副本，对于引用类型（go中的每一种引用类型，都有各自的实现，引用类型其实是指针），也与指针类似，引用的对象不会产生副本，副本只是这个引用（具体是引用实现的struct的副本，还是只是一个指针的副本或其它，没有研究过）

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
```
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

19. interface为nil与不为nil时的typeof是不相同的
   typeof(nilInterface) != typeof(notNilInterface)
20. slice copy, 如果size太小（不是容量），那么最多只复制size的内容，不会出错
21. 如果需要slice 的append不改变指针，那么可以可以slice以足够大的容量，让他不重新分配新的slice
22. 判断两个函数签名相同 ConvertibleTo AssignableTo
23. channel select  
    
24. mod管理在依赖第三方包时，要指定依赖的版本，如果直接依赖于master请说明充分的理由

    