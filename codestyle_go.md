[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)  
[한국어](./codestyle_go-ko.md)  
# Code Style -- go
SCRYINFO

## Rule 

1. Follow the six principles of software design 

	1. Open-Closed Principle (OCP) 
	2. Liskov Substitution Principle (LSP) 
	3. Dependence Inversion Principle 
	4. Interface Segregation Principle (ISP)
	5. Law of Demeter LoD is also known as the Least Knowledge Principle (LKP). 
	6. Simple responsibility pinciple SRP

2. Complete the function which is necessary to run the entire project rather than your own computer. 
3. Take the problem out and don't forget it in the development process; add the todo description in the code, add issues in github
4. Think firstly and then write the code starting with naming 
5 Handle each error and log it to the log 
6. Handle all branches, especially the branch in the abnormal situation (for example, data that should not appear and write it to the error log) 
7. Write the important call in the info log
8. Directly provide the service interface and it must be stable; the whole service process should not shut down for one error
9. Unify error number and error information for the external provided interface.
10. While defining the function, there are two aspects to consider: 1. whether the function code is reasonable for achieving the function code; 2. whether it is convenient for utilization or it is easy to make mistakes.
11. Verify the code in development process, use the unit testing, research the demo for the project technology implementation, etc
12. Give the sufficient reason if you want to use global variables 
13. Write the general small function in scryg after discussion
14. Requirements for code submit: Format compile passed. There should be a special reason to submit the code that fails to compile
15. Reference the recommendations in Effective Go https://golang.org/doc/effective_go.html

## Name

1. Use the lowercase and underline in all source code file
2. Use the lowercase and underline in all directory file names, 
3. Use the clear meaning English words while naming

## Catalog file
1. The unit test is placed in the same directory as the source code file. For example, the code file is “server.go” and the unit test file is “server_test.go”.
2. Place all demos in the "warehouse name /demo" directory
3. If it is a framework or base library, there is the "warehouse name / sample" to be required
4. All projects use package management (go mod)

## code 

1. Do not define the pointer to the interface, itself is the fat pointer

2. for i, v := range str { // code block } v is a copy; in order to avoid unnecessary copying, you can use only i to traverse; it is worth noting that the modification of i in the code block would be reset before the next round of cycling.
  ```go
Str := "abc.def"
For i := range str {
If str[i] == '.' {
    i += 2
}
fmt.Println(i, string(str[i]))
}
```
3. If the anonymous function (also called the closure) has a loop variable, there are two ways to solve it.

	One way is to deliver parameters; do not use loop variables
	Other way is to define the new variable

4. If the channel is empty, it will directly stuck while using it rather than panic

5. Read the closed channel which can correctly read the remaining value in the channel; if the channel is empty, it would read the null value of the channel type, and v, ok := <- c, ok is False 

6. The way to determine whether the channel shutdown is _, ok := <-c , and it also reads the data. If there is no number in the channel, it will wait. There is no way to directly determine that the channel is closed before the 1.10 version. 

7. select: If there is no case statement that can be run and no default statement, select would be blocked until the certain case communication can be run 

8. type T int is not the same as type T = int, the previous one defines the new type, the latter one defines the alias of int 

9. When implementing the interface, add the following code to ensure that all functions of the interface are implemented       
```go
 var(     
   _ interfaceName = (*interfaceImpl)(nil)
   ) 
```
10. recover:   
  (1) When using recover to capture panic, only current goroutine panic can be captured
 (2) Calling recover is useful only inside the defer function. 


11. The order of execution of return and defer, see https://github.com/googege/blog/blob/master/go/go/important/README.md
while running to the return, it would assign the value to the return value and run defer (the stack order is between defers, last in, first out). There should be notifying that if the return value is the same variable (no copy, the same one), if so the modification in the defer would affect the final return value, here are two special examples (see the webpage for more details) 
``` go
 / / return value is 1 rather than 2
func tt3() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return i
}
// the operation of the return function is 2 rather than 1
func tt4() func() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return func() int {
		return i // reference variable / the same variable
	}
}
//return value is 13 rather than 12
func tt5() (num int) {
	defer func() {
		num++
	}()
	return 12 // Assign the return value num to 12
}

// the return value is 1 
func tt6() (num int) {
	defer func() {
		num++
	}()
	return
}
// the return value points to 1
func tt7() (*int) {
    num := 0
	defer func() {
		num++
	}()
	return &num
}
```

Note: If there is only one statement after defer, then the variable will be assigned immediately; if defer is followed by the function, the variable will be assigned at execution time.


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

14. pass the Go parameter, all of which are value passing (does not support reference passing, a few languages ​​such as C++, C# support). The parameters entering the function are all copies.
16. the coming out of new is the pointer type, so you can't use new to initialize the type. Generally, you don't use new, but use "&TypeName{...}" as follows.
``` go
Var a = new([]int) //a yes *[]int type
Var a2 = make([]int, 0) //a2 is a []int type
```
17. The suggestion way to handle error
    See github.com/pkg/errors

    Func New(message string) error //If there is a ready-made error, there are three functions to choose.
    Func WithMessage(err error, message string) error //only add the new information
    Func WithStack(err error) error //only add the call stack information
    Func Wrap(err error, message string) error //Add stack and information at the same time

18. Check if the final object of the interface is empty
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
19.When interface and nil interface is nil，the type and pointed object is nil
```go
      Typeof(nilInterface) != tyvar inter1 interface{} = nil  // == nil
var inter2 interface{} = (*int)(nil) // != nil 因为类型值不为nil
fmt.Println(inter1 == nil)
fmt.Println(inter2 == nil)
//Output result is：
true
false
```
When interface is nil, typeof is different with when it is not nil
```go
var err error = nilvar err2 error = errors.New("")
fmt.Println("err : ", reflect.TypeOf(err))
fmt.Println("err2: ", reflect.TypeOf(err2))
fmt.Println(reflect.TypeOf(err) == reflect.TypeOf(err2))//result is：
err :  <nil>
err2:  *errors.errorStringfalse

//err and err2 type are error type， one is nil value and one is not empty. The type is different
```
20. slice copy, if the size is too small (not capacity), then only copy the content of the size and will not go wrong

21. When using append to add content to the slice, if the size does not exceed the capacity, the sclice will not be reassigned that means the address of the original slice will not change.

22. Two colons in the slice: for v :=data\[ a : b : c\], a, b are the upper and lower bounds respectively, and c is the capacity
The correct way to generate the copy of the slice is: c := v\[:0:0\]

23. Determine that the two function signatures are the same. ConvertibleTo AssignableTo

24. When mod manages the library, you must specify the version of the dependency. If you rely directly on the master, please explain the reason.

25. The declared variable v can appear in the condition of the ":=" statement:
    (1) The v of this declaration is in the same scope as the already declared v (if v has already been declared in the outer scope, this declaration will create a new variable).
    (2) The same value as the value of v in initialization can be assigned to v.
    (3) There is at least one variable in this statement when it is newly declared.

26. iota enumerator:
    (1) The iota constant auto generator automatically increments 1 every other line.
    (2) iota encounters meets const and resets to 0.
    (3) You can write only one iota. When the constant declaration omits the value, the default value is the same as the previous one.
    (4) If they are in the same line, the values are the same.
    (5) The iota must be explicitly restored after being interrupted.

27. Constant expression: In addition to the shift operator, if the binary operator is a different type of untyped constant, the result type is the latter one. For example, an untyped integer constant divided by an untyped complex constant results in an untyped complex constant.

28. fallthrough: Force run switch paired case,but it does not judge whether the result of the next case expression is true or false. Fallthrough can no longer be used in type switch.
29. Don’t define noname API in struct(embedding interface)
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
The code above can be compiled, The reason why compile passed without realizing API hi while running panic is that the Hi in the struct is only one field and don’t have name, the full call is like:
hello.Hi.HiName()，hello.Hi value is nil，so while running panic

* Add possibilities for error, compile passed but running error
* If Hello realize API Hi，then hello.HiName call is its own method rather than hello.Hi.HiName，easy to be misunderstanding.
* Struct embedded struct and inerface is one field, but interface embedded in interface should be required for realization methods 
30.Others
