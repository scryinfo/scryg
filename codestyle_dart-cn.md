[中文](./codestyle_dart-cn.md)  
# Code Style -- dart
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
11. 验证开发中的代码，使用单元测试；研究一项目技术实现等使用demo
12. 如果要使用使用全局变量，给出足够的理由
13. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由
## Name 
1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
## 目录文件
1. 单元测试?
2. 所有的demo放入“ 仓库名/demo ” 目录中
3. 如果是框架或基础库，需要“仓库名/sample”
## 代码
1.数组 可以看成列表一样。
var list=[] 和 List list = List();
2.函数 
dart中，所有类型都是对象，函数的对象类型是：Function。  可作为参数传递。
3.方法和变量
dart中并没有public、protected、private等关键字，声明变量与方法时，前面加上 "_" 即可作为private方法使用。不加，默认为public。 
注意： "_" 的限制范围并不是类级别，而是库访问级别
4.mixin混入
一般是 单继承、多实现，混入是多继承。
通过混入，一个类可以以非继承的方式，使用其他类中的变量和方法。
ps：mixin 在flutter源码中使用较多。
5.基本操作符
    除法与整除
        /       除号
        ~/      除号，但返回值是整数
    相等
        操作符是 ==            // 其中两个对象代表 同样内容 的时候返回true。
        如果要判两对象是否为 同一个对象，用identical（）方法。
    类型判定
        is     是指定类型，true
        is!    不是指定类型，true
    类型转换
        as     类型转换     eg：num x=666; x as int;
    赋值操作符
        ??= 
        a??=value, 如果a为null，则赋值value给a;如果不为null，则a不变 
    条件表达式
        常见表达式 term ? expr1 : expr2 
        另一种     expr1 ?? expr2     (如果expr1是non-null，返回其值；否则执行expr2并返回其接口)。
6.对象级联操作符
   ..  一个对象上，多次调用该对象的多个方法或成员。
    new Person()
        ..name = "not6"
        ..age = "110"
        ..saySomething();
7.条件成员访问操作符  
    区分操作符  ?. 和操作符 . 之间的区别
    ?.  和常规的成员访问操作符 . 相似， 但左边对象不能为null。
     如果左边操作对象为null，则返回 null，否则，返回右边的成员。
8.流程控制
    assert断言      ，只会在debug模式下生效。
    assert(x < 10);
    不符合条件，会抛出一个异常AssertionError。 程序中断
9.异常
    Exception和Error两个类型。    
    写法差异：
    try { 
        //...
    } on Exception catch (e) {    // catch 捕获所有异常
        print('Unknown exception: $e');
    }
    try { 
        //...
    } catch (e,s) {               // catch 可以带有一个或两个参数，第二个为堆栈信息
        print('Unknown exception: $e');
    }
10.异步dart-future和Microtask执行顺序
    Dart 中事件的执行顺序：Main > MicroTask > EventQueue
    示例代码：
    void testSX(){
      new Future(() => print('s_1'));
      scheduleMicrotask(() => print('s_2'));
      print('s_3');
    }
    结果：
    I/flutter (32415): s_3
    I/flutter (32415): s_2
    I/flutter (32415): s_1
    
        
10. 其它
