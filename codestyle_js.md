[中文](./codestyle_js-cn.md)  
[EN](./codestyle_js.md)  
# Code Style -- js

## Rule 

1. Follow the six principles of software design 

   1. Open-Closed Principle (OCP) 
   2. Liskov Substitution Principle (LSP) 
   3. Dependence Inversion Principle 
   4. Interface Segregation Principle (ISP)
   5. Law of Demeter LoD is also known as the Least Knowledge Principle (LKP)
   6. Simple responsibility pinciple SRP

1. Complete the function which is necessary to run the entire project rather than your own computer
2. Take the problem out and don't forget it in the development process; add the todo description in the code, add issues in github
3. Think firstly and then write the code starting with naming 
4. Handle each error and log it to the log 
5. Handle all branches, especially the branch in the abnormal situation (for example, data that should not appear and write it to the error log) 
6. Directly provide the service interface and it must be stable; the whole service process should not shut down for one error
7. Unify error number and error information for the external provided interface
8. while defining the function, there are two aspects to consider: 1. whether the function code is reasonable for achieving the function code; 2. whether it is convenient for utilization or easy to make mistakes
9. Verify the code in development process, use the unit testing, research the demo for the project technology implementation, etc
10. Give the sufficient reason if you want to use global variables 
11. Write the general small function in scryg after discussion
12 Requirements for code submit: Format compile passed. There should be a special reason to submit the code that fails to compile


## Name

1. Use the lowercase and underline in all source code file
2. Use the lowercase and underline in all directory file names
3. Use the clear meaning English words while naming
