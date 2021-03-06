# Yrel (from WoW, a guardian and great bishop of Draenei)

这算是 <两周自制脚本语言> 这本书的练习项目, 不过我从 Java 实现换成了 Golang 实现, 并且把语言的名字从平淡的 Stone 换成了我很喜欢 Yrel(伊瑞尔).

---

## Day 1: 起步

1. 不需要借助软件执行的设计语言, 也就是机器语言
2. 解释器和编译器:

    - 解释器(Interpreter): 根据程序中的算法执行运算, 这种软件也可以称为虚拟机.
    - 编译器(Compiler): 将某种语言携程的程序转换成另一种语言的程序, 通常目标程序为机器语言程序, 这个过程称为编译.

程序执行的典型流程:

1. 源代码, 通过```词法分析```到2
2. 单词排列( 简短字符串的排列 ), 通过```语法分析```到3
3. 抽象语法树, 到4具体执行
4. 分为两种

    - 编译器: 生成其他语言的程序
    - 解释器: 执行程序, 获得结果

---

## Day 2: 设计语言

1. 没有静态类型系统, 变量可以随时被赋给合法值
2. 句末没有分号
3. 必须使用花括号```{}```来包裹语句块
4. 没有```return```关键字, 语句块的最后一个表达式作为返回值
5. 使用双斜线```//```作为注释
6. ```if```和```while```中的条件表达式不需要括起来
7. ```else```关键字必须和前一个语句块的右括号以及下一个语句块的左括号一行

一段代码示例:

    sum = 0
    i = 1
    while i < 10 {
      sum = sum + i
      i = i + 1
    }

    if i % 2 == 0 {
      even = even + 1
    } else {
      odd = odd + 1
    }

---

## Day 3: 分割单词

这个章节开始词法分析的内容, 这里的主要作用是把代码分割成若干的```单词(token)```, 去掉没用的空格和注释.

Yrel 语言的单词只有三种类型: 标识符, 整型和字符串类型. 分别用```IdToken```, ```NumToken```和```StrToken```实现.

每次逐行从文件中读取内容, 并且通过正则匹配的方式来获得token:

1. IdToken: ```([A-Z_a-z][A-Z_a-z0-9]*|=|!=|==|<=|>=|&&|\|\|)```
2. numPattern: ```([0-9]+)```
3. strPattern: ```(\"[\S\s]*\")```
4. commentPattern: ```(//[\S\s]*)```

这三个结构体都实现了```Token```接口, 主要方法有:

1. ```IsNumber() bool```
2. ```IsIdentifier() bool```
3. ```IsString() bool```
4. ```GetText() string```

可以使用```GetTokenType()```方法来获得一个token的类型: ```Identifier```, ```Number```, ```String```或```Undefined```.

我们可以使用```NewLexer()```方法来创建一个词法分析器```Lexer```的实例, 这个实例的主要内容是一个```Queue[]```数组, 通过使用借口实现泛型, 用来保存分析得到的token序列.

```Lexer```有一个```Read()```方法, 可以用于从词法分析器中读取token, 当所有元素已经读取出来时, 将会返回```EOF```.

---

## Day4: 语法分析

进行语法分析的时候, 语法元素有终结符和非终结符的却别, 这里我使用```go tool```原生的```yacc```库编写一个语法分析器.

语法的定义使用BNF范式, 并且约定终结符为全大写, 非终结符为全小写.

1. 终结符

    - ```IDENTIFIER```: 对应词法分析中的```IdToken```类型
    - ```NUMBER```: 对应词法分析中的```NumToken```类型
    - ```STRING```: 对应词法分析中的```StrToken```类型
    - ```BOOL```: 对应词法中的```BoolToken```类型

2. 非终结符及其表达式

    - ```primary```: ```"(" expr ")" | NUMBER | IDENTIFIER | STRING | BOOL```
    - ```factor```: ```"-" primary | primary```
    - ```expr```: ```factor { OP factor }```
    - ```block```: ```"{" [statement] "}"```
    - ```simple```: expr
    - ```statement```:
        - ```"if" expr block [ "else" block ]```
        - ```"while" expr block```
        - ```simple```
    - ```program```: ```[statement]```