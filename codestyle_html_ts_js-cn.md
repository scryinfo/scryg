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

### ts

[Google Ts Style](https://google.github.io/styleguide/tsguide.html)  
[TypeScript style guide -- ts dev](https://ts.dev/style/)  

1. 统一使用“;”结尾
2. 定义变量
    * 不使用var
    * 在类型明确的情况下可以不给出类型
    * 尽量不使用any类型
    * 不使用 Object.create(null)

    ```ts
    // 不可修改，首选。 
    const count =  0;
    // 可修改
    let amount = 0;
    // 明确类型
    let name: string| null = null;
    // 不使用，因为一般的ts语法不能给object给null值，当使用Object.create(null)时就绕过空值检查
    let obj = Object.create(null)
    ```

3. string类型  
    * 没有特别原因不使用大写开头的String，它是一个object类型，而小写的string是基本类型
    * 字符串使用单引号。原因是html的属性，使用的双引，这样方便在里面表示字符串
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

    ```

4. 定义Array

    ```ts
    let keys = [''];//类型明确，可以不指定类型
    let keys2 = new Array<string>(); //明确指出类型
    let keys3 = new Array(); //不允许
    let keys4 = []; //不允许
    let keys5 = {}; //不允许
    ```

5. 定义Map
6. "可空类型"
7. 通过控制光标，可让移动端软键盘收回。
8. 一个域对应一组localStorage cookie。
9. document.referrer 只会是进入这个页面的url。
10. removeEventListener的时候永远不需写 passive 和 once。
11. 数字减一后再取模可以保持数据模运算后的顺序性。
12. 不使用null返回值
    * null是一种程序异常，不是返回值
    * 字符串时，如果没有值返回“”零长度字符
    * Object时可以提供一个全局的空对象
    * 如果有特殊情况，需要特别处理null的，给出足够的理由
13. 取数组的一部分时，slice更快

    ```ts
    let a = [0,1,2,3,4,5];
    a.length = 2;//第一种
    //或者
    a = a.slice(0,2);//第二种更快
    ```

14. 循环

* 不要循环中改变判断条件，如果业务实现需要更改，请给出足够的理由
* 在使用for循环时，不要在第二个参数上调用函数，因为每一次循环都会运行对应的函数，浪费cpu。如果判断条件在变化，需要运行函数时，请给出足够的理由。  
  如下是错误做法：

```ts
//let array = [];
//for(let i = 0; i < array.length; i++){}
```

正确的做法是：

```ts
let array = [];
for(let i = 0, len = array.length; i < len; i++){}
```

* for ... in 与 for ... of的区别

```ts
let a = ["one","two"];
for(let it in a){
    console.log(it); //,"0,1"
}
//for ... in
//输出属性，当为数组时，输出数组的下标
//类型为string

for(let it of a){
    console.log(it); //"one, two"
}
//for ... of
//输出集合的元素
//类型为集合元素的类型
```

* 最快的循环是for，如果特别需要性能，就不使用 foreach for .. in for .. of等这些循环

### vue

1. 在v-for中为item添加key 当列表有变化时，方便Vue精准找到该条列表数据，进行新旧状态对比，更新变化。
2. 尽量不要在v-for中使用v-if来过虑集合中的元素 可以增加一个计算属性，在计算属性中增加条件来过虑集合，因为计算属性是有缓存的
3. v-show与v-if的区别 v-show是修改display:none来达到不可见的，dom是一直存在的。v-if只有条件成立才生成dom
4. 如果列表是不可修改的，使用Object.freeze来告诉vue，以提高长列表的性能 Vue通过Object.defineProperty对数据进行劫持，实现视图响应数据的变化，直接告诉vue数据不可改，减少vue做无用的事情
5. 如果在Promise中修改邦定的数据，不能正常刷新到界面时，可以使用 this.$nextTick方法
6. 在像tree这样的递归组件使用时，记得在所有使用tree标签时，都要邦定事情，不然只能有第一层有效
7. 使用@Prop的字段时，记得处理@Watch它，以保证产生相应的变化
