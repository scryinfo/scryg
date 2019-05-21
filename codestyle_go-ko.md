[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)  
[한국어](./codestyle_go-ko.md)  
# Code Style -- go
SCRYINFO

## 규칙

1. 소프트웨어 설계의 6가지 원칙을 준수
	1. 오픈-클로즈 원칙(Open-Closed Principle, OCP)
	2. 리스코프 치환 원칙(Liskov Substitution Principle, 약자. LSP)
	3. 의존성 뒤집기 원칙(Dependence Inversion Principle)
	4. 인터페이스 분리 원칙(Interface Segregation Principle, ISP)
	6. 데미테르의 법칙(Law of Demeter LoD) 그리고 최소자승법이라고 함(Least Knowledge Principle,LKP)
	7. 단일 책임 원칙(Simple responsibility pinciple SRP)
2. 본인만의 컴퓨터가 아니라 전적인 프로젝트에서 정상적으로 실행해야만 기능 완료라고 볼 수 있다
3. 문제를 제기해야 한다. 즉 개발 과정 중 문제점을 쉽게 잊지 말고 코드에 todo 설명을 추가하여 github의 issues에 저장해야 한다
4. 사고한 후 명령 코드부터 작성해야 한다
5. 매 하나의 error를 처리하고 일지에 기록해야 한다
6. 모든 분기를 처리해야 한다. 특히 이상 문제가 발생한 분기를 처리(예: 나오지 말아야 하는 데이터가 나오는 등 문제를 error 일지에 기록해야 함) 
7. 주요한 호출은 info의 일지에 기록해야 한다
8. 외부에 직접 서비스 인터페이스를 지원하려면 반드시 안정적인 상태를 유지해야 하며 단 하나의 오류로 인하여 전체 서비스를 중단할 수 없다
9. 외부에 인터페이스를 지원하려면 에러 코드 및 에러 정보를 통일해야 한다
10. 함수를 정의할 시 두 가지 측면을 고려해야 한다. 첫째, 함수의 코드를 실현하는데 합리적인지 확인. 둘째, 사용하기 편리하고 쉽게 오류가 발생하는 여부에 대해 확인
11. 개발 중인 코드를 점검하려면 단위 테스트를 사용해야 한다. 즉 한 프로젝트 기술을 연구하는데 데모(Demo)를 사용해야 함
12. 전역 변수를 사용하려면 반드시 합리적인 이유를 제출해야 한다
13. 일반적인 작은 기능은 토론 후 scryg에서 작성해야 한다
14. 코드의 요구를 제출하려면 포맷 컴파일을 걸쳐야 한다. 만약 제출한 코드가 컴파일에서 통과하지 못한다면 특별한 이유가 있어야 한다
15. Effective Go 중의 건의를 참고 https://golang.org/doc/effective_go.html

## Name 

1. 모든 소스 코드의 파일 네임은 소문자 및 밑줄 치기를 해야 한다
2. 모든 디렉토리의 파일 네임은 소문자 및 밑줄 치기를 해야 한다
3. 명명할 시 의미가 있는 명확한 영어 단어를 사용해야 한다

## 디렉토리 파일
1. 단위 테스트 및 소스 코드 파일은 하나의 디렉토리에 저장해야 한다. 예를 들어, 코드 파일은 “server.go”이고 단위 테스트 파일은 “server_test.go” 이다
2. 모든 데모는 “창고명/Demo” 디렉토리에 저장해야 한다
3. 만약 프레임 워크 혹은 기본 라이브러리이면 “창고명/ sample”라고 해야 한다
4. 모든 프로젝트는 패키지 관리를 사용해야 한다(go mod)

## 코드
1. 인터페이스에 대한 포인터를 정하지 말아야 한다. 그 자체가 하나의 굵은 포인터이기 때문이다

2. for i, v := range str { // code block } 중 v는 복사이기 때문에 불필요한 복사를 피해야 하며 단 i로만 트래버스할 수 있다. 또한 code block에서 i에 대한 수정이 다음 순환 전에 초기화 된다는 것을 주의해야 한다
  ```go
str := "abc.def"
for i := range str {
	if str[i] == '.' {
    	i += 2
	}
	fmt.Println(i, string(str[i]))
}
  ```

3. 익명 함수(클로저라고도 부름)에 루프 변수가 있을 경우에는 두 가지 방법으로 해결할 수 있다

	첫째, 파라미터를 이전하려면 루프 변수를 사용하지 말아야 한다
	둘째, 새로운 변수를 정의해야 한다

8. channel이 비여 있을 경우 이것을 사용하면 panic가 아니라 바로 다운된다

9. 닫힌 channel을 읽으면 channel의 나머지 값도 정상적으로 읽을 수 있다. 만약 channel이 비여 있을 경우 해당 channel 타입의 빈값도 읽게 된다. 게다가 v, ok := <- c에서，ok는 false이다

10. 하나의 닫힌 channel를 판단하는 방법은 _, ok := <-c 이며 게다가 데이터도 읽을 수 있다. 만약 channel에서 데이터가 없을 경우 wait로 나온다. 1.10의 버전전까지 직접 판단하여 channel 닫히는 방법은 지원하지 않았다

10. Select: 운행 가능한 case 구절과 default 구절이 없을 경우 select는 막히게 되고 어느 하나의 case가 통해야만 운행될 수 있다

8. type T int 와 type T = int의 정의는 서로 다르다. 앞에 있는 정의는 새로운 타입을 표시하고 뒤에 있는 정의는 int의 일명이다

9. 인터페이스를 실현할 시 아래와 같은 코드를 추가하면 인터페이스의 모든 함수를 실현할 수 있다
       var (
       _ interfaceName     = (*interfaceImpl)(nil)
       )

10. recover:
    (1)Recover로 panic를 캡처할 시 현재 goroutine의 panic만이 캡처할 수 있다
    (2)Defer함수의 내부에서만 유용한 recover를 호출할 수 있다

11. return 및 defer의 실행 순서，see https://github.com/googege/blog/blob/master/go/go/important/README.md

Return까지 운행되면 반환 값에 대입되여 defer(defer 사이의 스택 순서는 후진 선출)으로 운행하게 된다. 반환 값이 동일한 변수(복사본이 없으면 동일함을 의미)인지 주의해야 한다. defer에서 수정할 경우 최종의 반환 값에 영향을 주게 된다. 아래는 두가지 특수 예시이다. ( 더 자세한 내용은 웹사이트를 참조)
``` go
//반환 값은 1이지 2가 아니다
func tt3() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return i
}
//리턴한 함수의 운행 결과는 2이고 1은 아니다
func tt4() func() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return func() int {
		return i // 참조 변수/동일 변수
	}
}
//반환 값은 13이지 12가 아니다
func tt5() (num int) {
	defer func() {
		num++
	}()
	return 12 // 반환 값 num의 대입은 12이다
}
//반환 값은 1
func tt6() (num int) {
	defer func() {
		num++
	}()
	return
}
//반환 값은 1에 향한다
func tt7() (*int) {
    num := 0
	defer func() {
		num++
	}()
	return &num
}
```
주의：Defer뒷면에 단 하나의 명령문이 있을 경우 변수는 즉시 대입되고 역시 defer 뒷면에 하나의 함수가 있을 경우에는 변수를 실행돼야만 대입될 수 있다

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

14. go의 변수 호출은 전부 다 값에 의해 호출(참조 호출을 지원하지 않음. C++，C#와 같은 몇 가지 언어를 지원함)되고 함수에 진입한 변수는 사본으로 되여 있다
16. new는 포인터 타입에서 나온 것이므로 new로 초기화 참조 타입에 사용하지 말아야 한다. 일반적으로 new를 사용하지 않고 “&TypeName{...}”를 사용해야 한다  
``` go
var a = new([]int) //a 是 *[]int 타입
var a2 = make([]int, 0) //a2 是 []int 타입
```
17. error에 대한 처리 방법
    see github.com/pkg/errors
    
    func New(message string) error //하나의 error가 발생할 경우 3개의 함수를 선택할 수 있다
    func WithMessage(err error, message string) error //새로운 정보만 첨부한다
    func WithStack(err error) error //호출할 스택 정보만 첨부된다
    func Wrap(err error, message string) error //스택 및 정보를 동시에 첨부한다

18. 인터페이스의 최종 객체가 비여 있는지 확인해야 한다
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

19. Interface가 nil거나 nil아닐 경우 typeof는 동일하지 않다
      typeof(nilInterface) != typeof(notNilInterface)
    
20. slice copy, size가 너무 작(용량은 아님)을 경우 많아서 size의 내용만 복사되고 오류는 발생되지 않는다

21. Append로 slice에 내용을 추가할 경우 마침 size가 최고 용량을 초과하지 않는다면 sclice는 다시 배정되지 않을 것이다. 즉 말 그대로 slice의 주소가 변하지 않는다

22. Slice에서의 두 콜론: v :=data\[ a : b : c\],a,b에 대해 a와b는 각각 상한과 하한이고 c는 용량이다
	Slice 사본을 생성하는 정확한 방법은 바로 c := v\[:0:0\] 이다

23. 두 함수의 사인이 동일한지 확인 ConvertibleTo AssignableTo

24. Mod가 디펜던스 패키지를 관리할 경우 디펜던스의 버전을 정해야 하고 직접 master를 의존하게 되면 충분한 해당 이유를 설명해야 한다

27. 스테이트먼트 된 변수 v는 ": ="성명문의 조건에서 나올 수 있다：
    (1)이번 스테이트먼트의 v는 이미 스테이트먼트 된 v와 동일한 영역(만약 v가 이미 외부의 영역에 스테이트먼트 될 경우 이번 스테이트먼트는 새로운 변수를 생성할 것임)에 있다
    (2)초기화에서 v의 값 타입이 같아야 v에 대입할 수 있다
    (3)이번 스테이트먼트에서 적어도 새로 스테이트먼트 된 하나의 변수가 있어야 한다

29. iota열거기：
    (1)Iota 상수 오토 제너레이터는 매 라인마다 자동으로 1을 누적한다
    (2)Lota가 const를 만나면 0으로 리셋된다
    (3)하나의 iota만 프로그래밍할 수 있으며 상수 스테이트먼트에서 값을 생략할 시 기본 값과 이전 값이 일치하다
    (4)한 라인에 있으면 그에 대한 값도 동일하다
    (5)Iota가 인터럽트 된 후 반드시 복구해야 한다

33. 상수 표현식: 시프트 연산자 외에 이항 연산이 다른 타입의 상수일 경우 결과는 후자이다. 예를 들어, 타입리스의 정수 상수를 타입리스의 복합 상수로 나눈다면 결과는 하나의 타입리스의 복합 상수가 된다

34. Fallthrough: switch 매칭한 case를 강제로 실행했지만 다음 case의 표현식 결과가 true 혹은 false라는 것에 대해 판단하지 않는다. 또한 fallthrough는 더 이상 type switch에서 사용할 수 없다
