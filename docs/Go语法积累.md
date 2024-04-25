# Go 语法积累

## 基础语法

### 一、基本数据类型

#### 1、 指针类型

- 优秀博客
  - [【入门】Go语言指针详解 - 乱七八糟博客备份 - 博客园 ](https://www.cnblogs.com/qinziteng/p/17280926.html)



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

