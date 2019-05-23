[中文](./codestyle_go-cn.md)  
[EN](./codestyle_go.md)  
[한국어](./codestyle_go-ko.md)  
[日本語](./codestyle_go-ja.md)  
# Code Style -- go
SCRYINFO
## 規則
1. ソフトウェア設計の七つ原則に守る
	1. 開―閉の原則(Open-Closed Principle, OCP)
	2. リスコフ置換の原則(Liskov Substitution Principle,略称LSP)
	3. 依存性逆転の原則(Dependence Inversion Principle)
	4. インタフェース分離の原則(Interface Segregation Principle, ISP)
	5. デメテルの法則(Law of Demeter LoD)最少知識の原則(Least Knowledge Principle,LKP)とも言う
	6. 単一責任の原則(Simple responsibility pinciple SRP)
2. 機能を完成した。自分のパソコンで運営することではなく、全プログラムが正常に運営し、配置することです。
3. 問題を忘れずに開発過程に選出します。コードの中にtodo説明を加え、Githubのissuesに追加します。
4. 名づけからコードのプログラミングを考えます。
5. 全てのerrorを処理して日誌に記録します。
6. 全ての分かれを処理します。特に現れた異常状況の分かれです。（出てきてはいけないデータなどをerror日誌に記録します）。
7. 重要なコールをinfo日誌に記録する必要があります。
8. 直接外界へ提供するサービスインタフェースを必ず安定させます。一つのエラーで全てのサービスを停止させてはいけません。
9. 外部に提供するインタフェース、エラー番号と錯誤情報を統一します。
10. 関数を定義する際に、二つの問題を考える必要があります。一つ目は関数コードの合理性です。二つ目は使用の利便性です。エラーが出やすいか
11. 開発しているコードを単体テストで検証します。プログラムの技術実現などのために使用するdemoを研究します。
12. 全局の変数を使用する場合には、十分な理由が必要です。
13. 通用の小機能は検討を通してからscrygに書き込みます。
14.コードを提出する要求、説明とフォーマッティングはコンパイルします。コンパイルしないコードを提出する場合には、特別な理由が必要です。
15. Effective Goの中のアドバイスを参考にします。  https://golang.org/doc/effective_go.html
## Name 
1. 全てのソースコードのファイル名は小文字を使って、アンダースコアも加えます。
2. 全てのコンテントのファイル名は小文字を使って、アンダースコアも加えます。
3. 命名は明確な意味がある英語を使います
## コンテントファイル
1. 単体テストとソースコードファイルを同じコンテンツに収録します。例えば、コードファイルは“server.go”の場合には単体テストのファイルは“server_test.go”です。
2. 全てのdemoを「倉庫名／demo」というコンテンツに収録します。
3. フレームワークや基礎ライブラリの場合は「倉庫名／sample」が必要です。
4. 全てのプログラムはゴーモード（go mod）で管理します。
## コード
1. interfaceのニードルを定義してはいけません。元々はファトゥニードルです。
2. for i, v := range str { // code block } の中のｖはコピーの意味です。従って、必要がないコピーをしないように、iだけでエルゴードできます。注意するべきのはcode blockの中にiに対する修正は次の循環の前にリセットします。
```go
str := "abc.def"
for i := range str {
	if str[i] == '.' {
    	i += 2
	}
	fmt.Println(i, string(str[i]))
}
```
3. もし匿名関数（クロージャ）が循環変数を使う際に、二つの解決方式があります。
	一つ目は循環変数ではなくパラメータをまわります。
	二つ目は新しい変数を定義します。
4. もしchannelがブランクの場合、使ったらpanicではなく直接的にシステムダウンになります。
5. 閉めたchannelを読み込んで、channelの残りの値を正確的に読み込まれます。Channelが空の場合に、このchannel類の空値が読み込まれます。v, ok := <- cの中に、okはfalseです。
6. Channelが閉めたかどうかを判断する方法は、_, ok := <-c ，註釈してデータを読み込みます。もしchannelの中に数がなかったら、プログラムがwaitになります。1.10バージョンの前に、channelが閉めたかどうかを直接的に判断できる方法を提供しません。
7. select：アクティブできるcaseとdefaultコードがない場合に、selectがあるcase通信がアクティブできるまでに閉塞します。
8. type T int  とtype T = intは違います。type T intは新しいタイプを定義します。type T = intはintのある別名を定義します。
9. インタフェースを実現する際に、以下のコードを加えてインタフェースの全ての関数の実現を確保します。
```go
var (
    _ interfaceName     = (*interfaceImpl)(nil)
)
```
10. recover:
* recoverでpanicを捕獲する際に、目下の goroutineのpanicしか捕獲できません。
* defer関数の内部しかrecoverがコールできません。
11. returnとdefer の実行順番は、see https://github.com/googege/blog/blob/master/go/go/important/README.md
がreturnまで実行して戻り値に値呼びます。defer（deferの間はスタック順番で、出るが前にして入るが後ろにします）を実行します。戻り値が同じ変数かどうかを注意するべきです（副本を生成することではなく、同じものです）。同じ変数の場合に、最後の戻り値がdeferの中の修正に影響されます。以下のは二つの特別な例です。（詳しい内容はホームページで見られます）
``` go
//戻り値が２ではなく１です
func tt3() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return i
}
//戻った関数の実行結果は１ではなく２です
func tt4() func() int {
	var i = 0
	defer func() {
		i++
	}()
	i++
	return func() int {
		return i // 変数引用／同一変数
	}
}
//戻り値は12でななく13です
func tt5() (num int) {
	defer func() {
		num++
	}()
	return 12 // 戻ったnumは12に値呼ばされます
}
//戻り値が１です
func tt6() (num int) {
	defer func() {
		num++
	}()
	return
}
//戻り値が１に指します
func tt7() (*int) {
    num := 0
	defer func() {
		num++
	}()
	return &num
}
```
注意：もしdeferの後ろのコードが一つだけの場合に、その中の変数がすぐに値呼ばされます。もしdeferの後ろには関数である場合に、その中の変数が実行される時に値呼ばされます。
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
14. goのパラメータの伝送は全部が値の伝送です。（伝送の引用はできません、C++，C#など少数のプログラミング言語を支持する）関数に入ったパラメータは全部副本です。
16. newはニードルタイプなので、newで引用タイプを初期化することができません。一般的には、以下のようにnewではなく“ &TypeName{...}”を使います。
``` go
var a = new([]int) //a は *[]int タイプ
var a2 = make([]int, 0) //a2 は []int タイプ
```
17. errorの処理方式は下記をアドバイスする：
    see github.com/pkg/errors
    func New(message string) error //既存のerrorがあったら、三つの関数が選ばれます。
    func WithMessage(err error, message string) error //新しい情報のみ付加します
    func WithStack(err error) error //コールしたスタック情報のみ付加します
    func Wrap(err error, message string) error //スタックと情報を同時に付加します
18. インタフェース最終の対象がブランクかをチェックします
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
19. Interface and nil
When interface and nil interface is nil，the type and pointed object is nil
```go
Typeof(nilInterface) != tyvar inter1 interface{} = nil  // == nil
var inter2 interface{} = (*int)(nil) // != nil,because the type is not nil 
fmt.Println(inter1 == nil)
fmt.Println(inter2 == nil)
//Output result is：
true
false
```
When interface is nil, typeof is different with when it is not nil
```go
var err error = nilvar 
var err2 error = errors.New("")
fmt.Println("err : ", reflect.TypeOf(err))
fmt.Println("err2: ", reflect.TypeOf(err2))
fmt.Println(reflect.TypeOf(err) == reflect.TypeOf(err2))
//result is：
err :  <nil>
err2:  *errors.errorString
false

//err and err2 type are error type， one is nil value and one is not empty. The type is different
```
20. slice copy,もしsizeが小さい過ぎる（容量ではなく）時に、せいぜいsizeの内容しかコピーしません。且つミスが出ません。
21. appendでsliceに内容を加える際に、sizeが容量を超えない場合、scliceを再分配しません。つまり、sliceの元アドレスが変わりません。
22. slice中の二つのコロン：v :=data\[ a : b : c\]に対して、a,bそれぞれが上限と下限で、cは容量です。
	slice副本を生成する正しい方法： c := v\[:0:0\]
23. 二つの関数サインの一致制を判断する ConvertibleTo AssignableTo
24. modで依存パッケージを管理する際に、依存バージョンを指定する必要があります。直接にmasterに依存する場合、十分な理由を説明する必要があります。
25. 声明した変数vは”:=”声明に現れる条件：
* 今回声明するvは声明したvと同じスコープにあります（もしvが外層スコープに声明されたら、今回の声明は新しい変数を作成します）
* 初期化中にv値と同じタイプの値しかvに値呼ばされます。
* 今回の声明の中に少なくとも一つの変数が得られます。
26. iotaイテレータ：
* iota定数自動発生器は一行ごとに自動的に１を累計します。
* iotaが constと出会うと、0にリセットします。
* iotaは一つだけ書いてもよいです。定数が省略値を声明するときに、前の値と同じと見做します。
* 同じ行にいれば、値は同じです。
* iotaが中断されたら必ず表示を回復します。
27. 定数エクスプレッション：シフト演算子以外に、もし二元演算子は異なるタイプの無タイプの定数であったら、結果タイプが後ろの一つになります。例えば、ある無タイプの整数定数を無タイプの複数定数で割り算すると、結果は無タイプの複数定数になります。
28. fallthrough：switchとマッチングしたcaseを強制に実行しますが、次のcaseのエクスプレッションの結果がtrueかfalseかを判断しません。そしてfallthroughはtype switchに使用できません。
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
The code above can be compiled, The reason why compile passed without realizing API hi while running panic is that the Hi in the struct is only one field and don’t have name, the full call is like:hello.Hi.HiName()，hello.Hi value is nil，so while running panic
* Add possibilities for error, compile passed but running error
* If Hello realize API Hi，then hello.HiName call is its own method rather than hello.Hi.HiName，easy to be misunderstanding.
* Struct embedded struct and inerface is one field, but interface embedded in interface should be required for realization methods 
30. Others
