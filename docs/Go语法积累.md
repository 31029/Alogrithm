# Go 语法积累

## 基础语法

### 一、基本数据类型

[The Go Programming Language Specification - The Go Programming Language](https://go.dev/ref/spec#Types)

[go 基本数据类型、关键字、内建函数、内建常量   --  博客园](https://www.cnblogs.com/dev-shi/p/12421077.html)

#### 1、指针类型

- **`优秀博客`**
  - [【入门】Go语言指针详解 - 乱七八糟博客备份 - 博客园 ](https://www.cnblogs.com/qinziteng/p/17280926.html)

#### 2、字符串类型

- **`优秀博客`**
  - [Go语言字符串综合指南：函数、方法和最佳实践_go语言 字符串-CSDN博客](https://blog.csdn.net/walkskyer/article/details/135093920)

#### 3、数学类型

- **`优秀博客`**
  - [如何在Go中使用操作符进行数学运算_go 浮点数和整数乘法-CSDN博客](https://blog.csdn.net/QIU176161650/article/details/133672895)

#### 4、`常用技巧`

##### ① 数据类型相互转化

- [go语言string、int、int64、float64、complex 互相转换](https://studygolang.com/articles/13139)





### 二、复合类型

#### 1、数组、切片

- [Go 切片常用操作与使用技巧_go 切片操作-CSDN博客](https://blog.csdn.net/u014082714/article/details/139151999)

#### 2、结构体

- [golang interface{} 和 struct{} - 多课网，360度全方位IT技术服务站！ (duoke360.com)](https://www.duoke360.com/post/5883)



### 三、函数

#### 1、形参、实参

- **Go语言中slice，map这三种类型的实现机制类似指针**，所以可以直接传递，而不用取地址后传递指针。（注：若函数需改变slice的长度，则仍需要取地址传递指针）
  - [Golang语言基础教程：函数的参数 - 知乎 (zhihu.com)](https://zhuanlan.zhihu.com/p/68049307)



#### 2、函数调用栈



### Ⅰ、其他

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





### Ⅱ、常用库

#### 1、标准库：fmt

① **`fmt.printf`**

- [golang fmt.printf() - 码农骆驼 - 博客园](https://www.cnblogs.com/rxbook/p/7085783.html)



#### 2、标准库：sort



#### 3、标准库：container

- [精通Go：container库的全面解析与实战应用_go container-CSDN博客](https://blog.csdn.net/walkskyer/article/details/135670678)
