[中文](./codestyle_ts-cn.md)  
# Code Style -- Typescript
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
5. 处理每一个error
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
1. 定义一个数组，进行push操作之前，一定要有值或者初始化。  
2. 会变化的值最好不要直接当做，for循环的判断条件，因为for循环有可能改变它。  
3. 通过控制光标，可让移动端软键盘收回。
4. 一个域对应一组localStorage cookie。
5. document.referrer 只会是进入这个页面的url。
6. removeEventListener的时候永远不需写 passive 和 once。
7. 数字减一后再取模可以保持数据模运算后的顺序性。
8. 其他
