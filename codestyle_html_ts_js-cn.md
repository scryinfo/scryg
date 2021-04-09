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
14. 使用css不使用style
15. 在布局时能不使用具体的数字就不使用
16. 使用flex进行布局
17. 单位使用rem

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
4. 相对于另一个div2的对齐
	增加一个div1，把需要的内容放在里面，且设置float，这样增加的div1的高度为零就不会占用布局的位置，让增加的div1与div2靠在一起，内容部分使用position: relative。这样就可以实现比较完美的对齐
5. !important 语法 可以覆盖 element.style, 以及 JS 中控制的样式。
6. z-index 会影响触发事件的元素，并不会影响事件机制本身
7. z-index表示在层叠上下文中的显示顺序，不能超过层叠上下文
8. bootstrap navbar里面的下拉菜单 safari 点击空白 不会回收。
9. safari 中 元素是默认不可点击的。

### ts/js
1. 定义一个数组，进行push操作之前，一定要有值或者初始化。
2. 会变化的值最好不要直接当做，for循环的判断条件，因为for循环有可能改变它。
3. 通过控制光标，可让移动端软键盘收回。
4. 一个域对应一组localStorage cookie。
5. document.referrer 只会是进入这个页面的url。
6. removeEventListener的时候永远不需写 passive 和 once。
7. 数字减一后再取模可以保持数据模运算后的顺序性。

### vue
1. 在v-for中为item添加key
   当列表有变化时，方便Vue精准找到该条列表数据，进行新旧状态对比，更新变化。
2. 尽量不要在v-for中使用v-if来过虑集合中的元素
   可以增加一个计算属性，在计算属性中增加条件来过虑集合，因为计算属性是有缓存的
3. v-show与v-if的区别
	v-show是修改display:none来达到不可见的，dom是一直存在的。v-if只有条件成立才生成dom
4. 如果列表是不可修改的，使用Object.freeze来告诉vue，以提高长列表的性能
   Vue通过Object.defineProperty对数据进行劫持，实现视图响应数据的变化，直接告诉vue数据不可改，减少vue做无用的事情


