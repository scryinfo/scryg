[中文](./codestyle_js-cn.md)  
[EN](./codestyle_js.md)  
# Code Style -- js

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

## Name 

1. 所有源代码文件名，使用小写，加下划线
2. 所有目录文件名，使用小写，加下划线
3. 命名使用有明确函义的英文单词
4. 不使用数据库的关键字或保留字命名，如不能使用for来命名一个字段名
## 目录文件
1. 

## 代码
1. 
