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

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不使用数据库的关键字或保留字命名，如不能使用for来命名一个字段名

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

* 类型使用 UUpperCamelCase命名
* 变量使用 lowerCamelCase
* 全局常理使用 CONSTANT_CASE
* 不使用"_"下划线来作为前缀或后缀
* 

### TS

1. ECMA, ECMAScript,JavaScript, TypeScript  
    ECMA: 是一个组织，全称 European Computer Manufacturers Association，它制定了很多标准规范，其中就有ECMAScript  
    ECMAScript: 简称ES，在ECMA中的编号为 [ECMA-262](https://www.ecma-international.org/publications-and-standards/standards/ecma-262/) ，  
        JSON也是由ECMA制定的，JSON的代号为 [ECMA-404](https://www.ecma-international.org/publications-and-standards/standards/ecma-404/)  
    TS39: Technical Committee 39, 是ECMA的一个技术委员会，主要制定ECMAScript，所以也会看到 TS39-ECMAScript的字样  
    ES6： ECMA-262 6th Edition, [ECMA-262的第六版](https://262.ecma-international.org/6.0/)  
    ES7： [ES2016](https://262.ecma-international.org/7.0/)，  
        从这里开始，更多的以年来表示版本，现在最新的是[ES2022](https://262.ecma-international.org/13.0/)  
    JavaScript: 简称JS，是（符合）ECMAScript的一个语言实现，它们大体相容，有不相的。[JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)  
    TypeScript: 简称TS，TypeScript is a strongly typed programming language that builds on JavaScript, giving you better tooling at any scale. [TypeScript](https://www.typescriptlang.org/docs/handbook/typescript-from-scratch.html)  

    汇总： 有一个组织叫ECMA，有代号为ECMA-262的规范，名叫ECMAScript，TC39是ECMA其中的一个技术委员会，由它来做ECMAScript，  
        Javascript是ECMAScript大体相容的语言实现，Typescript是基于Js的强类型语言。  

2. 统一使用“;”结尾，方便于在html等插入代码  
3. 优先使用小写类型，如string,number; 尽量不使用String,Number等类型
4. 定义变量  
    * 不使用var  
    * 在类型明确的情况下可以不给出类型  
    * 尽量不使用any类型  
    * const 与 readonly  
        const是不能给变量重新赋值，可以修改内部的数据  
        readonly在修辞变量时，是不可以修改内部数据，与是否重新赋值无关
        readonly在定义字段时是只赋值一次，且在构造时  
        readonly不能使用于所有类型（'readonly' type modifier is only permitted on array and tuple literal types）  
    * as const 用在 literal values定义时，他们变成readonly

    ```ts
    // 不可修改，首选。 
    const count =  0;
    // 可修改
    const amount = 0;
    // 明确类型
    const name: string| null = null;
    // 使用 as const
    const obj = {name:"test", h: 10} as const;
    ```

5. string类型  
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

6. Array
    优先使用type[]这种类型的数组

    ```ts
    const keys = [''];//可以
    const key: string[] = ['']; //明确给出类型
    const keys2 = new Array<string>(); //不建议使用Array类型，直接使用string[]
    const keys3 = new Array(); //不允许
    const keys4 = []; //不允许
    const keys5 = {}; //不允许

    const keys6 = Array.from<number>({length: 5}).fill(0);//这个可以 Array.from的返回类型为 number[]
    //合并数组
    const a1: number[] = [1,2,3];
    const a2: number[] = [4,5];
    const a12 = [...a1,...a2];
    //调整数组大小

    ```

7. Map
    没有小写的map类型或关键字

    ```ts
    const m = new Map<string, number>();//类型明确
    const m2 : Map<number,number> = new Map([[1,2]]);//可以
    const m3 = new Map(); //不允许
    const m4 = {}; //不允许
    const m5 = {"1":2}; //不允许
    
    //合并Map
    const data1: Map<number,number> = new Map([[1,2],[3,4]]);
    const data2: Map<number,number> = new Map([[1,2],[3,4]]);
    const data = new Map([...data1,...data2]);
    console.log(data.size);// 2, 已去掉重复
    ```

8. 类型特点
    * primitives: string,number,boolean  
        number包含整数与浮点数，实际上它是浮点数，所以不要使用number 来存储特别大的int64的整数。  
        它是f64的浮点数，最大存储的整数是2^53 - 1，如果超过这个值就会有问题。[js number](https://en.wikipedia.org/wiki/IEEE_754)  
        Number.MAX_VALUE = 1.7976931348623157e+308  
        Number.MIN_VALUE = 5e-324  
        在使用时，一定注意它的最大值不是int64的最大值  
        在ES2020 中有 [BigInt](https://v8.dev/features/bigint) 类型，这个类型不用担心int64的问题  
        如果想要表示准备的整数，请使用BigInt等类型，而不要使用number类型
    * null，它是一个特殊的object类型，有唯一值 null  
        不要定义一个null类型的变量  
        null使用在可空类型上  
        typeof 返回的类型是 object  

        ```ts
        const v: null = null; //不允许，也不可以定义class的字符，函数参数
        const v2: string|null = null; //可空的string
        ```

    * undefined，是一个特殊类型，有唯一值 undefined  
        不要定义undefined类型的变量，它是为了解决一个变量或字段是否定义而引入的关键字，并不是用来定义变量的  
        typeof 返回的类型是 undefined  

    * Symbol 类型: 不可变唯一,不建议使用,[see](https://www.typescriptlang.org/docs/handbook/symbols.html)

        ```ts
        const id2 = Symbol('key');
        const id3 = Symbol('key');
        id2 === id3; // false, 因为 symbol是唯一
        ```

    * object类型  
        除了“基础类型”外，都是object类型，如Array,Map,DateTime等 build-in object，null也是一种特殊的object类型
    * “可空类型”

        ```ts
        const v3: string|null = null; //可空的string，正常会修改它的值所以使用let
        const v4: string|null|undefined; //不允许，undefined是表示没有定义，而这里在定义一个变量。
        ```

    * any类型，尽量不要使用，使用它相当于把ts退回来js

    * 没有byte类型
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

9. 函数  

    * => 函数，只使用在变量或参数上  
        不使用 => 函数定义成员函数  
        不使用 => 定义全局函数  
        可以使用 => 定义局部变量  
        可以使用 => 调用函数时  
        可以使用 => 实现接口
    * 使用默认值参数代替可选参数  
        默认值在写代码时，更不容易出。而在使用上是一样的  

    * 明确 Destructuring 参数类型  
    * Destructuring 参数默认值，不建议整个{}给默认值  
    * 如果语法可以，尽量给方法参数加上readonly说明，表明方法不修改参数值
    * 明确函数的返回类型，如果没有返回值，那为void
    * 如果函数不返回任务类型，那么使用never

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
    ```

10. this，是明确，不像js中的this
    * 成员函数this就是对像自己  
    * => 函数，在类中时，this就是对象自己  
    * => 类函数外，不要使用this  

11. if(v)

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
    empty.toString; //可以正常运行，因为toString是一个不在在的字段，返回值为 undefined
    const v: any;
    if (v){} //false， 此时v的值为 undefined
    const v2: Object;
    if (v2){} //编译不通过 “Variable 'v2' is used before being assigned”
    // 

    ```

12. === and !==  
    尽量使用三等或不等，这个比较结果是明确的，而“==”不明确。
    当与null 或undefined比较时，可以使用“==”或 “！=”，这个要看比较的结果

    ```ts
    const v : string | null = null;
    if (v === null) {} //false

    //注意，在“== null”时如果变量为undefined, 结果为true
    const v2: string | undefined = undefined;
    if(v2 == null) {} // true, 不要这样使用
    ```
  
13. for循环

    * 不要循环中改变判断条件，如果业务实现需要更改，请给出足够的理由
    * 在使用for循环时，不要在第二个参数上调用函数，因为每一次循环都会运行对应的函数，浪费cpu。如果判断条件在变化，需要运行函数时，请给出足够的理由。  
        如下是错误做法：
    * 不使用 for in循环，它容易误解，实际上是遍历对象的属性
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

14. Destructuring/解构

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

    // 函数参数


    ```

15. class
    1. 尽量减少使用this
    2. 在单个函中，不能使用this,可以明确的传参数
    3. 不要要构造函数中使用this,这时的this并不明确或没有构造出来
    4. 不能使用 => 字义成员函数
    5. 可以使用 => 定义事件触发函数，这时它相当于一个成员变量
    6. 不要在成员函数上使用bind, 它会让人误解，且可能会有内容问题

    7. object to interface时，不要使用as  

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


16. 不要使用’// @ts-ignore’

17. 当使用“a as Type”或"a!" 一定要加上明确的说明，为什么类型一定是对的。不要使用 (<Type>a)这种语法

18. 使用model, 不使用 namespace  

19. 不使用default export,它的含义不明确（除非代码不能正常编译）

[Google Ts Style](https://google.github.io/styleguide/tsguide.html)  
[TypeScript style guide -- ts dev](https://ts.dev/style/)  
[ECMAScript 2020](https://262.ecma-international.org/11.0/#sec-intro)  
[TypeScript 5.2](https://www.typescriptlang.org/docs/handbook/release-notes/typescript-5-2.html)  
[v8](https://github.com/v8/v8)  
[TypeScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)  

### vue

1. 在v-for中为item添加key 当列表有变化时，方便Vue精准找到该条列表数据，进行新旧状态对比，更新变化。
2. 尽量不要在v-for中使用v-if来过虑集合中的元素 可以增加一个计算属性，在计算属性中增加条件来过虑集合，因为计算属性是有缓存的
3. v-show与v-if的区别 v-show是修改display:none来达到不可见的，dom是一直存在的。v-if只有条件成立才生成dom
4. 如果列表是不可修改的，使用Object.freeze来告诉vue，以提高长列表的性能 Vue通过Object.defineProperty对数据进行劫持，实现视图响应数据的变化，直接告诉vue数据不可改，减少vue做无用的事情
5. 如果在Promise中修改邦定的数据，不能正常刷新到界面时，可以使用 this.$nextTick方法
6. 在像tree这样的递归组件使用时，记得在所有使用tree标签时，都要邦定事情，不然只能有第一层有效
7. 使用@Prop的字段时，记得处理@Watch它，以保证产生相应的变化
