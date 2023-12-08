[中文](codestyle_html_ts_js-cn.md)

# Code Style -- html

[See](https://google.github.io/styleguide/htmlcssguide.html)

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
6. 处理所有分支，特别出现的异常分支（如，不应该出现的数据等，写入error日志）
7. 直接对外提供服务接口，必须稳定，不能因为一个错误就让整个服务停止工作
8. 在对外提供的接口中，统一错误编号及提示
9. 定义函数时要考虑两个方面，一实现函数代码是否合理，二使用是否方便，是否容易出错
10. 验证开发中的代码时，使用单元测试；在研究一项目技术实现等使用demo
11. 如果要使用使用全局变量，给出足够的理由，因为它很难测试
12. 通用的小功能，经过讨论后写入 scryjs 中
13. 提交代码的要求， 说明 格式化 编译通过，如果提交编译不通过的代码需要有特别的理由
14. 使用css不使用style，如果要使用style给出理由
15. 在布局时能不使用具体的数字就不使用
16. 使用flex进行布局
17. 单位使用rem或em或vh/vw

## Name

不建议使用： 可以使用但不推荐，如果要使用需要有足够的理由  
禁止使用： 不可以使用，理由再多都不行  
允许使用： 可以使用  
建议使用： 推荐使用  
必须使用： 一定这样使用

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不使用数据库的关键字或保留字命名，如不能使用for来命名一个property名

## 目录文件

1. 所有的demo放入“ 仓库名/demo ” 目录中
2. 如果是框架或基础库，需要“仓库名/sample”

## 代码

### CSS/HTML

1. 左右居中
2. 上下居中
    1. 文字上下居中
       *. line-height == height
    2. div上下居中
3. 图片居中
4. 相对于另一个div2的对齐 增加一个div1，把需要的内容放在里面，且设置float，这样增加的div1的高度为零就不会占用布局的位置，让增加的div1与div2靠在一起，内容部分使用position:
   relative。这样就可以实现比较完美的对齐
5. !important 语法 可以覆盖 element.style, 以及 JS 中控制的样式。
6. z-index 会影响触发事件的元素，并不会影响事件机制本身
7. z-index表示在层叠上下文中的显示顺序，不能超过层叠上下文
8. bootstrap navbar里面的下拉菜单 safari 点击空白 不会回收。
9. safari 中 元素是默认不可点击的。
10. metia中的max-width与min-width中的值都不包含“=”
11. 图片在适应不同屏幕宽度时，可以使用metia来设置不同的图片

```html
<picture>
    <source srcset="/media/cc0-images/surfer-240-200.jpg"
            media="(min-width: 800px)">
    <img src="/media/cc0-images/painted-hand-298-332.jpg" alt="" />
</picture>
<img
 srcset="example-320w.jpg 320w,example-480w.jpg 480w,example-800w.jpg 800w"
 sizes="(max-width: 320px) 280px,(max-width: 480px) 440px,800px"
 src="example-800w.jpg" alt="An example image">
```

12. 单位  
    width:10% 相对于元素的10%  
    width:10vw 相对于viewport宽度的10%  
    width:10em 相对于元素的font-size的，如果没有设置会继承元素的  
    width:10rem 相对于html的font-size的，如果没有设置为16px

### TS name

* 函数Function, 是广义的，所有函数的统称，有全局函数，局部函数，=> 函数，成员函数/方法，等都是函数
* 全局变量，定义在函数外的变量
* 局部变量，定义在函数内的变量
* 函数变量，函数类型的变量
* 变量，是各种变量的统称，当单独使用变量时，不包含Field
* const变量，由const定义的变量
* let变量，由let定义的变量
* 常量，指不会被修改的变量，可以由const readonly定义与修辞。注：const变量，并不一定是常量，如使用const定义一个数组，是可以在定义后修改数组中元素
* Class函数，与Class相关的函数，包含有this与无this的
* 成员函数Memeber Function，就是class函数，在本文档中不使用成员函数一说（成员函数，一般C++中使用的比较多）
* static Function, 没有this的class函数
* 方法Method，有this的class函数
* 字段Field, 与Class的this相关，可以使用this使用的“变量”，单独使用时，不包括static Field
* static Field, 与Class相关，直接使用Class可以使用的变量
* Class变量/字段，包含Field与static Field
* Property, js的Propery在ts中对应关系，会增加使用ts的难度，在本文档中不单独使用property
* get set Property, 是一种特殊的Method，可以当作字段来使用
* =>函数(箭头函数)，使用 =>实现的函数
* 自定义类型使用 UpperCamelCase命名
* 变量与Field使用 lowerCamelCase
* 全局常理使用 CONSTANT_CASE
* 不使用"_"下划线来作为前缀或后缀

### TS

1. ECMA, ECMAScript,JavaScript, TypeScript  
    ECMA: 是一个组织，全称 European Computer Manufacturers Association，它制定了很多标准规范，其中就有ECMAScript  
    ECMAScript: 简称ES，在ECMA中的编号为 [ECMA-262](https://www.ecma-international.org/publications-and-standards/standards/ecma-262/) ，  
        JSON也是由ECMA制定的，JSON的代号为 [ECMA-404](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)  
    TS39: Technical Committee 39, 是ECMA的一个技术委员会，主要制定ECMAScript，所以也会看到 TS39-ECMAScript的字样  
    ES6： ECMA-262 6th Edition, [ECMA-262的第六版](https://262.ecma-international.org/6.0/)  
    ES7： [ES2016](https://262.ecma-international.org/7.0/)，  
        从这里开始，更多的以年来表示版本，现在最新的是[ES2022](https://262.ecma-international.org/13.0/)  
    JavaScript: 简称JS，是（符合）ECMAScript的一个语言实现，它们大体相容，有不相容的。[JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)  
    TypeScript: 简称TS，TypeScript is a strongly typed programming language that builds on JavaScript, giving you better tooling at any scale. [TypeScript](https://www.typescriptlang.org/docs/handbook/typescript-from-scratch.html)  

    汇总： 有一个组织叫ECMA，有代号为ECMA-262的规范，名叫ECMAScript，TC39是ECMA其中的一个技术委员会，由它来做ECMAScript，  
        Javascript是ECMAScript大体相容的语言实现，Typescript是基于Js的强类型语言。  

2. 变量与class字段  
    1. 禁止使用var  
    2. 在类型明确的情况下可以不给出类型  
    3. 不建议使用any类型  
    4. const 与 readonly  
        const是不能给变量重新赋值，可以修改内部的数据  
        readonly修辞变量时，是不可以修改内部数据，与是否重新赋值无关
        readonly在定义字段时是只赋值一次，且在构造时  
        readonly在语法上，不能使用于所有类型（'readonly' type modifier is only permitted on array and tuple literal types）  
    5. as const 用在 literal values定义时，他们变成readonly

        ```ts
        // 不可修改，首选。 
        const count =  0;
        // 可修改
        const amount = 0;
        // 明确类型
        const name: string| null = null;
        // 使用 as const
        const obj = {name:'test', h: 10} as const;
        ```

    6. Destructuring/解构

        ```ts
        // 数组
        const data = [0,6];
        const [f1,l1] = data;// f = 0, l = 6
        const [f2,,l2] = data;// f = 0, l = undefined
        const [f3,l3, ...t] = data;// f = 0, l = 6, t = [] 不是undefined
        const [f4,...t2,l4] = data; //语法错误，并不支持这种方式
        const [f5 = 2, l5 = 4] = data; // f = 0, l = 6， 有对应的值，不会使用默认值
        const [f6,,l6 = 8] = data;// f = 0, l = 8， 

        // 对象
        {
            const obj = { a: 1, b: { c: 2 } };
            const {
                a,
                b: { c: d },
            } = obj;
            //上面的代码，只定义了 a,d两个变量，其中b与c只是字段名而已
        }

        {
            // undefined在解构时认为不存在，而null是存在的
            const {a = 1} = {a:undefined}; //a为1
            const {b = 1} = {b:null}; //b为null
        }

        ```

    7. 展开

        ```ts
        //数组
        const a1 = [1,2];
        const a2 = [3,4];
        const a3 = [...a1, ...a2,5];
        //map
        const m1: Map<number,number> = new Map([[1,2]]); // key: 1, value: 2
        const m2: Map<number,number> = new Map([[3,4]]);
        const array3 = [...m1, ...m2]; // 不建议使用，小心，array3是一个数组，不是map
        const m4 = new Map([...m1, ...m2]); // 
        const m5 = {...m1, ...m2}; // 结果为 {}， 不允许使用
        // map在初始化与展开，都使用的是是kv的数组，特别小心

        //object 
        const o1 = {name:'x'};
        const o2 = {point: 10};
        const o3 = {...o1,...o2,other:'data'}

        ```

    8. {} 是什么类型，它是empty class

        ```ts
        const v1 = {}; //v是什么
        const v2 = Object.create(null);// console.log输出的结果是一样
        if (v1 === v2) {} // false
        if (v1 == v2){} // false
        // 不建议使用 Object.create()来创建对象

        ```

3. 函数  

    * => 函数，使用在变量或参数上  
        禁止使用 => 定义class函数  
        禁止使用 => 定义全局函数  
        可以使用 => 定义局部变量  
        建议使用 => 调用函数传参时  
        可以使用 => 实现接口  
        可以使用 => 事件响应，在这时，可以把事件响应定义为字段（只有这时可以使用=> 定义字段）  
    * 可以使用，用默认值参数代替可选参数  
        默认值在写代码时，更不容易出。而在使用上是一样的  

    * 明确 Destructuring 参数类型  
    * Destructuring 参数默认值，不建议整个{}给默认值，除非默认值是导入或定义的外部分变量  
    * 如果语法可以，尽量给函数参数加上readonly说明，表明函数不修改参数值  
    * 明确函数的返回类型，如果没有返回值，那为void  
    * 如果函数不返回任务类型，那么使用never  
    * 不建议使用 overload,可以使用默认参数替代

    ```ts
    function f1(name?:string) {} //可选参，不建议使用，是一种语法糖， string|undefined
    function f2(name: string|null){} //建议使用
    function f3(name: string|undefined){} //不建议使用
    function f4(name: string|null|undefined){} //不建议使用
    function f5(name = '10') {} //建议使用
    //建议使用 string | null方式的原因是
    //函数在使用时，会给出明确的参数，减少调用者发生错的可能性
    //带undefined时，与默认值不好区分，所以不建议使用
    function f6({a,b}) {} //不建议使用，给出明确的类型
    // readonly
    function f7(data: readonly number[]) {
        data.pop();// 编译错误，没有方法pop
        data[0] = 0;//编译错误，不能赋值
    }

    // 必须使用，检查是否为可空参数，一定使用 三个“=”
    if (p === null) { }

    // 必须使用，检查是否为可选参数，使用 三个“=”
    if (p === undefined) {} 

    {//解构
        //给出默认值及参数类型，在有默认值且类型明确时，可以不指定类型
        function f8({a = 1,b:string = ''}) {}

        //不允许在两个地方给默认， 没有什么特别原因，不要把默认值给在整个对象上
        //在给默认值时，要么给在前面，要么给在后面

        function f9({a:number,b:string = ''} = {a: 10}) {}
        //小心（不建议）使用可选参数, 它相当于一个控制变量，也就是说它大部分情况下，可以分为两个函数或者使用默认值解决
        function f10({a, b = ''}:{a?: number, b: string}) {}
        function f11({a, b = ''}:{a: number|undefined, b: string}) {}
    }

    ```

4. 类型
    1. primitives type : string,number,boolean，优先使用 primitives type，少使用内部Wrap类型  
        number包含整数与浮点数，实际上它是浮点数，所以不要使用number 来存储特别大的int64的整数。  
        它是f64的浮点数，最大存储的整数是2^53 - 1，如果超过这个值就会有问题。[js number](https://en.wikipedia.org/wiki/IEEE_754)  
        Number.MAX_VALUE = 1.7976931348623157e+308  
        Number.MIN_VALUE = 5e-324  
        在使用时，一定注意它的最大值不是int64的最大值  
        在ES2020 中有 [BigInt](https://v8.dev/features/bigint) 类型，这个类型不用担心int64的问题  
        如果想要表示准确的整数，请使用BigInt等类型，而不要使用number类型
    2. null，它是一个特殊的object类型，有唯一值 null  
        不要定义一个null类型的变量或字段  
        null使用在可空类型上  
        typeof 返回的类型是 object  

        ```ts
        const v: null = null; //不允许定义null类型的变量或字段
        const v2: string|null = null; //可空的string
        ```

    3. undefined，是一个特殊类型，有唯一值 undefined  
        不允许定义undefined类型的变量或字段，它是为了解决一个变量或字段是否定义而引入的关键字，并不是用来定义变量的  
        typeof 返回的类型是 undefined  

    4. Symbol 类型: 不可变且唯一,不建议使用,[see](https://www.typescriptlang.org/docs/handbook/symbols.html)

        ```ts
        const id2 = Symbol('key');
        const id3 = Symbol('key');
        id2 === id3; // false, 因为 symbol是唯一
        ```

    5. object类型  
        除了“基础类型”外，都是object类型，如Array,Map,DateTime等 build-in object，null也是一种特殊的object类型
    6. “可空类型”

        ```ts
        const v3: string|null = null; //可空的string，正常会修改它的值所以使用let
        const v4: string|null|undefined; //不允许，undefined是表示没有定义，而这里在定义一个变量。
        ```

    7. any类型，尽量不要使用，使用它相当于把ts退回到js

    8. 没有byte类型
        一般使用Uint8Array来处理bytes，如果是stream或很大的bytes可以使用ArrayBuffer类型

        ```ts
        // bytes to hex string
        const uint8Array = new Uint8Array([72, 101, 108, 108, 111]);
        const hexString = Array.from(uint8Array, (byte) => byte.toString(16).padStart(2, '0')).join('');
        //上一行代码中 (byte) 的括号可以省略

        // number to hex string
        const number = 26;
        const hexString = number.toString(16);
        
        // hex string to bytes
        const hexArray = hexString.match(/.{1,2}/g).map(byte => parseInt(byte, 16));
        return new Uint8Array(hexArray);
        
        // 可以使用库buffer（在node中有Buffer而在浏览器中没有，所以可以使用库）
        cosnt buf = Buffer.from('0236','hex'); //hex string to buffer
        const bytes = Uint8Array.from(buf); //buffer to Uint8Array
        const buf2 = Buffer.from(bytes); //Uint8Array to buffer
        const hexString = buf.toString('hex');// buffer to hex string
        ```

    9. string类型  
        * string的编码是utf-16
        * 没有特别原因不使用大写开头的String，它是一个object类型，而小写的string是基本类型
        * 字符串使用单引号。原因是html的属性，使用双引，这样方便在里面表示字符串
        * 当字符串中有相互包含时，不受“字符串使用单引号的聘限制”
        * 有格式字符时，建议使用 "Template literal types "/字符模板

        ```ts
        // 字符串
        const name = '';
        // 如果包含单引用时
        const name2 = "'";
        // 字符拼接
        // Template literal types
        const name3 = `key: ${name}`;
        const name4 = name + name2; //简单连接
        const name5 = [name,name2,'none'].join();
        const name6 = name.concat(name2);
        ```

    10. array
        优先使用type[]这种类型的数组

        ```ts
        const keys = [''];//可以
        const key: string[] = ['']; //明确给出类型
        const keys2 = new Array<string>(); //不建议使用Array类型，直接使用string[]
        const keys3 = new Array(); //禁止使用, 没有明确的类型
        const keys4 = []; //禁止使用
        const keys5 = {}; //禁止使用

        const keys6 = Array.from<number>({length: 5}).fill(0);//可以使用 Array.from的返回类型为 number[]
        //合并数组
        const a1: number[] = [1,2,3];
        const a2: number[] = [4,5];
        const a12 = [...a1,...a2];
        //调整数组大小

        ```

    11. Map
        没有小写的map类型或关键字  

        ```ts
        const m = new Map<string, number>();//类型明确
        const m2 : Map<number,number> = new Map([[1,2]]);//可以使用
        const m3 = new Map(); //禁止使用
        const m4 = {}; //禁止使用
        const m5 = {"1":2}; //禁止使用
        
        //合并Map
        const data1: Map<number,number> = new Map([[1,2],[3,4]]);
        const data2: Map<number,number> = new Map([[1,2],[3,4]]);
        const data = new Map([...data1,...data2]);
        console.log(data.size);// 2, 已去掉重复
        ```

5. Class/Type/Interface
    1. 在非方法中，不建议使用this,可以明确的传参数  
    2. 特别小心在构造函数中使用this,这时的this并不明确或没有构造出来
    3. this，是明确，不像js中的this  
        * 成员函数this就是对像自己  
        * => 函数，在类中时，this就是对象自己  
        * => 类函数外，不建议使用this  

    4. 禁止使用 => 字义class函数
    5. 可以使用 => 定义事件触发函数，这时它相当于一个函数类型的字段
    6. 禁止作用，在class函数上使用bind, 它会让人误解，且可能会有内容问题
    7. object to interface时，不建议使用as  

        ```ts
            interface Data {
                index: number;
                nam?: string;
            }

            // 正确做法
            const foo: Data = {
                index: 123,
                name: "abc", 
            };
            // 错误做法
            const foo = {
                index: 123,
                name: "abc", 
            } as Data;

        ```

    8. 可以使用 parameter propertied(在构造函数上定义属性)  
    9. 可以使用getters与setters  
    10. 类型转换及相关
        * （不建议使用）判断类型 util.types.“ function isDate(object: unknown): object is Date” 方法实现说明  
            Object.prototype.toString.call(data) === '[object Date]'  
            调用toString方法来比较字符串，这种方法给人的感觉是很不可靠的
        * （不建议使用）instanceof :  checks whether the prototype chain of x contains Foo.prototype,  
            也就是说它是检查类型是否相容，并不是确定类型。  
        * （不建议使用）x as ClassName
            在ts中编译会类型检查，但类型转换，最终是在js运行时完成的，而class类型主要由ts来处理，这种判断很可能在运行时出问题。  

    11. Intersection Types (交叉类型)

        ```ts
        interface Colorful {
            color: string;
            radius: number;
        }
        interface Circle {
            radius: number;
        }
        type ColorfulCircle = Colorful & Circle;
        interface ColorfulCircleEx extends Colorful{
            radius: number;
        }

        const v : ColorfulCircle = {color:'red', radius: 10};
        const v2: Circle = v; //ok

        const v3 : ColorfulCircleEx = {color:'red', radius: 10};
        const v4: ColorfulCircle = v3; //ok

        // & 与 extends 的效果是一样的，它们的区别是，当有冲突（同名但类型不同）时的处理方式不一样
        // &： 变成never类型
        // extends: 两个都有

        ```

    12. Generic Type  
        Think of Box as a template for a real type, where Type is a placeholder that will get replaced with some other type  
        理解，泛型是一种type的template，在运行时类型参数替换，它与C#的实现很像。
        一个最直接的效果时Generic中static field只有一份，因为generic type只有“一份”

        ```ts
        class Box<Type> {
            static onlyOne = 1;
        }

        const b1 = Box<number>();
        const b2 = Box<string>();
        Box.onlyOne = 2;
        Box<number>.onlyOne = 10; //语法错，因为只有一份，所以不能通过Box<number>访问
        b1.onlyOne = 3; //不建议这样使用，它会误以为"onlyOne"是在实例上的
        
        ```

    13. 在constructor在开始处调用supper()
    14. (建议)在derived class，如果要override方法时，参数相同，不然容易让人误解
    15. 禁止继承 Built-in Types。因为它们是在js-engine中实现的，继承的效果与一般的class不一样，为为避免这种不一至的结果。
    16. 不建议，在derived class中定义与base class同名字段
    17. static Blocks in Classes， 这是比较好的初始方式，注它只会运行一次，是在运行import时
    18. 禁止直接把方法作为参数传递。

        ```ts
        class MyClass {
            name = "MyClass";
            getName() {
                return this.name;
            }
        }
        const c = new MyClass();
        const obj = {
            name: "obj",
            getName: c.getName,
        };

        // Prints "obj", not "MyClass"
        console.log(obj.getName());
        
        ```

    19. class is a collection of propeties.

        ```ts
        class Point1 {
            x = 0;
        }
        class Point2 {
            x = 0;
        }
        
        // OK,  这个赋值是可以的，因为他们有相容的field
        const p: Point1 = new Point2();

        class Person {
            name: string;
            age: number;
        }
        
        class Employee {
            name: string;
            age: number;
            salary: number;
        }
        
        // OK, 这个赋值是可以的，因为他们有相容的field
        const p: Person = new Employee();
        
        ```

    20. Utility Types(转换现在类型)

    ```ts
    // Partial<Type>, 把字段转换为可选字段
    interface Todo {
        title: string;
        description: string;
    }
    Partial<Todo> == {
        title?: string;
        description?: string;
    }
    
    //Required<Type> 把字段转换为非可选
    //Readonly<Type> 把类型转换为readonly, 不可改
    //Record<Keys, Type> 把类型转换为 Keys与value，有点像Map
    interface CatInfo {
        age: number;
        breed: string;
    }
    type CatName = "miffy" | "boris" | "mordred";
    const cats: Record<CatName, CatInfo> = {
        miffy: { age: 10, breed: "Persian" },
        boris: { age: 5, breed: "Maine Coon" },
        mordred: { age: 16, breed: "British Shorthair" },
    };
    //Pick<Type, Keys> 包含类型中的指定字段
    //Omit<Type, Keys> 排除类型中的指定字段，Exclude<UnionType, ExcludedMembers>对于UnionType
    


    ```

6. tsconfig.json

    [tsconfig](https://www.typescriptlang.org/tsconfig)  

    ```json
    {
        "compilerOptions": {
            "strict": true,  //使用 strict
            "strictNullChecks": true,
            "strictPropertyInitialization": true,
            "noImplicitAny": true,
            "noImplicitReturns": true,
            "noImplicitThis": true,
            "strictBindCallApply": true,
            "strictFunctionTypes": true,
            "useDefineForClassFields": false, // 使用[Define]不使用[Set]，两条路二选一
            "skipLibCheck": false,
            "noEmit": true, //如果相查看输出文件，可以改为false

            "moduleResolution": "node16",  // nodenext == node16
            "module": "node16", 
            "target": "ESNext", // 使用最新版，最后再经过vite后，会生成都可以使用的版
            "lib": ["ESNext", "DOM","DOM.Iterable", "ScriptHost"],

            "jsx": "preserve",
            "resolveJsonModule": true,
            "isolatedModules": true,
            "esModuleInterop": true,

            "baseUrl": ".",
            "paths": {
                "@/*": ["src/*"]
            },

        },
        "include": [
            "src/**/*.ts",
            "src/**/*.d.ts",
            "src/**/*.tsx",
            "src/**/*.vue"
        ],
        "references": [
            {
            "path": "./tsconfig.node.json"
            }
        ]
    }
    ```

7. 统一使用“;”结尾，明确表明代码结束，也方便于在html等插入代码  

8. if(v)

    ```ts
    if (undefined || null ){} //false
    if ("" || '' || false || 0 || 0.0 ){
        console.log("any true");// not print
    }else{
        console.log("false");// 输出 false
    }
    if ({} && [] && Object.create(null)){} // true
    
    // 特别说明Object.create(null)值为null，但是if(Object.create(null)) 这里为true
    const empty = Object.create(null);// empty 的值是 {}， 所以它在if中是true
    empty.toString(); //toString 不是一个方法
    empty.toString; //可以正常运行，因为toString是一个undefined的property，返回值为 undefined
    
    const v: any;
    if (v){} //false
    const v2: Object;
    if (v2){} //编译不通过 “Variable 'v2' is used before being assigned”
    
    ```

    [see](https://www.typescriptlang.org/docs/handbook/2/narrowing.html#truthiness-narrowing)  
    0,NaN,"" (the empty string),0n (the bigint version of zero),null,undefined => false

9. === and !==  
    必须使用三等或不等，这个比较结果是明确的，而“==”不明确。

    ```ts
    const v : string | null = null;
    if (v === null) {} //true，建议使用

    //注意，在“== null”时如果变量为undefined, 结果为true
    const v2: string | undefined = undefined;
    if(v2 == null) {} // true, 不建议使用，除非明确目标
    if(v2 === null) {} // false, 可以使用
    ```
  
10. for循环

    * 不建议，在循环中改变判断条件，如果业务实现需要更改，请给出足够的理由
    * 在使用for循环时，不建议在第二个参数上调用函数，因为每一次循环都会运行对应的函数，浪费cpu。如果判断条件在变化，需要运行函数时，请给出足够的理由。  
        如下是错误做法：
    * 不建议使用 for in循环，它容易误解，实际上是遍历对象的属性
    * 使用 for of循环或明确要遍历的内容，如使用 for(const entry of a.entries()){}

    ```ts
    // 不建议做法
    let array = [];
    for(const i = 0; i < array.length; i++){}
    // 建议做法 
    let array = [];
    for(const i = 0, len = array.length; i < len; i++){}
    // * for ... in 与 for ... of的区别
    // for in 遍历的是属性，for of遍历可iterable的集合中的元素
    let a = ["one","two"];
    for(const it in a){ //不建议使用
        console.log(it); //,"0,1"
    }
    //for ... in
    //输出属性，当为数组时，输出数组的下标
    //类型为string
    for(const it of a){
        console.log(it); //"one, two"
    }
    //for ... of
    //输出集合的元素
    //类型为集合元素的类型

    // 也可以使用使用，明确要遍历的内容
    for(const it of a.entries()){}

    // 明确遍历属性
    for(const it of Object.keys(a)){}

    ```

11. 不建议使用’// @ts-ignore’

12. 当使用“a as Type”或"a!" 一定要加上明确的说明，为什么类型一定是对的。不要使用 (<Type>a)这种语法

13. 建议使用model, 不建议使用 namespace  

14. 不建议使用default export,它的含义不明确（除非代码不能正常编译）

15. import

    ```ts
    import * from 'x'; //禁止使用
    import * as name from 'x'; // 可以使用, 定义别名
    import {name} from 'x'; // 可以使用

    ``` 

[Google Ts Style](https://google.github.io/styleguide/tsguide.html)  
[TypeScript style guide -- ts dev](https://ts.dev/style/)  
[ECMAScript 2020](https://262.ecma-international.org/11.0/#sec-intro)  
[TypeScript 5.2](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-5-2.html)  
[v8](https://github.com/v8/v8)  
[JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)  

### vue

1. 建议使用setup来定义组件

    ```ts
    <template>
    </template>

    <script setup lang="ts">

    </script>
    ```

2. 建议使用ref而不使用reactive，因为reactive看起来像是使用普通变化，更容易发生错误的理解

    ```ts
    // ref是深层响应性的，也就是深层嵌套对像
    import {ref} from 'vue';
    const data = ref({inner: {count:0}});
    // 以下代码会触发响应
    data.value.inner.count++;

    // ref的对象的层次或properties数量不要太多，如果实在太多，为了提高性能可以使用 shallowRef/triggerRef
    const data2 = shallowRef({count:0});
    data2.value.count++; //不触发响应
    data2.value = 10;//触发响应
    triggerRef(data2); //手动触发响应

    // 整个对象替换后，修改单个property会触发响应
    const data3 = ref({inner:{level2:{count:0}}});
    data3.value.inner.level2 = {count:1};
    //...
    data3.value.inner.level2.count = 10;//会触发响应

    ```

3. 

11. 在v-for中为item添加key 当列表有变化时，方便Vue精准找到该条列表数据，进行新旧状态对比，更新变化。
12. 尽量不要在v-for中使用v-if来过虑集合中的元素 可以增加一个计算属性，在计算属性中增加条件来过虑集合，因为计算属性是有缓存的
13. v-show与v-if的区别 v-show是修改display:none来达到不可见的，dom是一直存在的。v-if只有条件成立才生成dom
14. 如果列表是不可修改的，使用Object.freeze来告诉vue，以提高长列表的性能 Vue通过Object.defineProperty对数据进行劫持，实现视图响应数据的变化，直接告诉vue数据不可改，减少vue做无用的事情
15. 如果在Promise中修改邦定的数据，不能正常刷新到界面时，可以使用 this.$nextTick方法
16. 在像tree这样的递归组件使用时，记得在所有使用tree标签时，都要邦定事情，不然只能有第一层有效
17. 使用@Prop的property时，记得处理@Watch它，以保证产生相应的变化
