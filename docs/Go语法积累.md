# Go 语法积累

## 基础语法

### 一、基本数据类型

#### 1、指针类型

- 优秀博客
  - [【入门】Go语言指针详解 - 乱七八糟博客备份 - 博客园 ](https://www.cnblogs.com/qinziteng/p/17280926.html)
  - [指针定义、指针特点、空指针、指针数组、指向指针的指针、指针作为函数入参_go的 指针数组实体定义- CSDN](https://blog.csdn.net/wohu1104/article/details/99694277)

#### 2、字符串类型

- 优秀博客
  - [Go语言字符串综合指南：函数、方法和最佳实践_go语言 字符串-CSDN博客](https://blog.csdn.net/walkskyer/article/details/135093920)

#### 3、数学类型

- 优秀博客
  - [如何在Go中使用操作符进行数学运算_go 浮点数和整数乘法-CSDN博客](https://blog.csdn.net/QIU176161650/article/details/133672895)

#### 4、`常用技巧`

##### ① 数据类型相互转化

- [go语言string、int、int64、float64、complex 互相转换](https://studygolang.com/articles/13139)





### 二、变量

#### 1、变量初始化

```go
在Go语言中，初始化变量有以下几种方式：

1、使用 var 关键字，显式声明变量类型，并初始化。

var name string = "John"
var age int = 30

2、使用 := 符号，Go会自动推断类型。

name := "John"
age := 30

3、使用 new 函数，返回指针，内置函数，用于分配内存。

var name = new(string)
*name = "John"

4、使用 make 函数，用于 map 和 chan 的初始化。

var nameMap = make(map[string]string)
var ch = make(chan int)

5、匿名结构体初始化。

person := struct {
    name string
    age  int
}{
    "John",
    30,
}

6、使用 struct 关键字，初始化结构体的所有字段。

type person struct {
    name string
    age  int
}

p := person{name: "John", age: 30}

7、使用 & 符号，获取变量的内存地址。

var name = "John"
var ptr = &name

以上就是Go语言中初始化变量的几种常见方式。
```





### 三、常用库

#### 1、fmt 标准输出库

① **`fmt.printf`**

- [golang fmt.printf() - 码农骆驼 - 博客园](https://www.cnblogs.com/rxbook/p/7085783.html)



#### 2、sort 库
