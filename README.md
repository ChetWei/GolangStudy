# 语言基础

## 1.语言结构

```go
package main //声明包名，报名和文件没有特定的联系

import "fmt"

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")  //可以加；或不加，建议不加
}
```



1. 第一行代码 `package main` 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. 下一行 `import "fmt"` 告诉 Go 编译器这个程序需要使用 fmt 包（的函数，或其他元素），fmt 包实现了格式化 IO（输入/输出）的函数。
3. 下一行 `func main()` 是程序开始执行的函数。main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4. 下一行 /*...*/ 是注释，在程序执行时将被忽略。单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾，且不可以嵌套使用，多行注释一般用于包的文档描述或注释成块的代码片段。
5. 下一行 *fmt.Println(...)* 可以将字符串输出到控制台，并在最后自动增加换行字符 \n。
   使用 fmt.Print("hello, world\n") 可以得到相同的结果。
   Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台。
6. 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。



编译并执行程序

``` sh
go run hello.go  
```

先编译成可执行文件

```sh
go build hello.go
这时候会生成一个可执行文件hello
```



## 2.数据类型

>  数据类型用于声明函数和变量

### 布尔型

```go
var b bool = true;  //值可以为 true 或 false
```

### 数字类型

- 整型
    - 无符号 uint8、uint16、uint32、uint64
    - 有符号 int8、int16、int32、int64
- 浮点型
    - float32、float64
    - complex64（32 位实数和虚数）、complex128(64 位实数和虚数)
- 其他
    - byte 类似 uint8
    - rune 类似 int32
    - uintprt 无符号类型，用于存放一个指针

### 字符串类型

> 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

### 派生类型

- 指针类型（Pointer）
- 数组类型
- 结构化类型(struct)
- Channel 类型
- 函数类型
- 切片类型
- 接口类型（interface）
- Map 类型



## 3.变量

### **一、指定变量类型、定义变量**

```go
var varname type
var varname1, varname2 type
```

#### 实例

```go
package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)
}
```

#### 没有初始化变量的默认值

- 数值类型（包括complex64/128）为 **0**
- 布尔类型为 **false**
- 字符串为 **""**（空字符串）

### 二、根据值自行判定变量类型

```go
var v_name = value //省略的变量的类型
```

实例

```go
package main
import "fmt"
func main() {
    var d = true //省略了变量的类型，自动根据值推断类型
    fmt.Println(d)
}
```

### 三、使用 `:=` 省略 `var关键字`和`变量的类型`

```go
intVal :=1 // 等价于 var intVal int = 1
```

> - 使用 := 不能对已经声明过的变量进行再次声明
>
> - := 只能用在函数体里面，不能用在全局变量



### 多变量声明

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

//只能用在函数体里面
vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

```go
package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}
```

### 值类型和引用类型

所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：

![4.4.2_fig4.1](https://gitee.com/ChetWei/img/raw/master/img/202205142012203.jpgrawtrue)

当使用等号 `=` 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了拷贝：

```go
func main() {
	var a int = 9;
	var b = a;

	fmt.Println(&a)
	fmt.Println(&b)

}
```

```
a addr: 0x140000b4008
b addr: 0x140000b4010
```

**引用类型**

![4.4.2_fig4.3](https://gitee.com/ChetWei/img/raw/master/img/202205142019383.jpgrawtrue)

## 4.常量

> 常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

#### 普通常量

```go
const identifier type = value
```

可以省略类型说明符type，因为编译器可以根据变量的值来推断其类型。

```go
//显式类型定义
const b string = "abc"

//隐式类型定义
const b = "abc"

//多个相同类型的声明可以简写为：
const c_name1, c_name2 = value1, value2
```

```go
package main

import (
	"fmt"
	"unsafe"
)

//常量枚举
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = LENGTH * WIDTH

	sizea := unsafe.Sizeof(a)
	fmt.Println("sizea: ", sizea)

	fmt.Println("面积为 : %d", area)
	println(a, b, c)
	fmt.Println("unknown=", Unknown)
}

```

#### iota 特殊常量

> iota，特殊常量，可以认为是一个可以被编译器修改的常量。
>
>
>
> iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

```go
const (
	a = iota //0
	b //iota +=1  1
	c //iota +=1  2
	d = "ha" //独立值，iota += 1   3
	e        //"ha"   iota += 1  4
	f = 100  //iota +=1 	5 
	g        //100  iota +=1 6
	h = iota //7, iota +=1 恢复计数
	i        //8, iota +=1
)

func main() {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h)
	fmt.Println("i", i)
}
```

```
a 0
b 1
c 2
d ha
e ha
f 100
g 100
h 7
i 8
```

**位移操作**

```go
const (
	k = 1 << iota //左移 0 位，不变仍为 1
	l = 3 << iota //3 (011)  左移 1 位，变为二进制 110，即 6。
	m             //6 (110)左移2位，变为 12(1100)
	n             //12(1100)左移3位，变为 24(11000)
)

func main() {
	fmt.Println("k", k)
	fmt.Println("l", l)
	fmt.Println("m", m)
	fmt.Println("n", n)
}
```

## 5.数组

### 声明数组

```go
var blance [10]float32
```

### 初始化数组

```go
	//声明并初始化数组
	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//通过字面量在声明数组的同时快速初始化
	balance3 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//如果数组的长度不确定，可以使用...代替数组的长度，编译器会根据元素个数自行推断数组的长度
	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
```

### 访问数组元素

```go
func main() {
	//声明数组
	var blance [10]float32

	var i int
	for i = 0; i < 10; i++ {
		blance[i] = float32(i * 2.0)
	}

	for i = 0; i < 10; i++ {
		fmt.Println(blance[i])
	}
	//声明并初始化数组
	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//通过字面量在声明数组的同时快速初始化
	balance3 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	//如果数组的长度不确定，可以使用...代替数组的长度，编译器会根据元素个数自行推断数组的长度
	balance4 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	
	fmt.Println(balance2)
	fmt.Println(balance3)
	fmt.Println(balance4)
}
```

## 切片(动态数组)

### 普通数组的传参

```go
func printArray(myArray [3]int) {
   //有指定长度数组时，使用的是值拷贝
   fmt.Println("----------------printArray----------")
   for index, value := range myArray {
      fmt.Println(" index= ", index, " value= ", value)
   }
   myArray[0] = 1000
}

func main() {

   var myArray = [3]int{10, 20, 30}

   printArray(myArray)

   printArray(myArray)
}
```

### 动态数组(切片)的传参

```go
func printArray2(myArray []int) {
   //使用切片动态数组参数的时候，使用的是引用传递
   fmt.Println("----------------printArray----------")
   for _, value := range myArray {
      fmt.Println(" value= ", value)
   }
   myArray[0] = 1000
}

func main() {

   //动态数组，切片slice，根据元素个数自动分配空间
   var myArray = []int{10, 20, 30}

   //切片引用传递
   printArray2(myArray)

   printArray2(myArray)
}
```

### 判断空切片

```go
func main() {
   var slice1 []int

   //空切片
   if slice1 == nil {
      fmt.Println("len slice ", len(slice1)) //没有分配slice空间 长度为0
   }

   slice1 = make([]int, 3) //开辟3个空间，int默认值为0
   fmt.Println("len slice ", len(slice1))
}
```

### 切片的扩容机制

```go
func main() {
   //创建一个切片，最大容量 cap=5，3个默认值0
   var numbers = make([]int, 3, 5)
   fmt.Println("numbers", numbers)
   fmt.Printf("numbers len = %d, cap=%d,slice=%v \n", len(numbers), cap(numbers), numbers)

   numbers = append(numbers, 1)
   numbers = append(numbers, 2)

   //切片扩容机制
   //此时cap的容量已经满了，会动态开辟新的容量，开辟的容量为 +cap
   numbers = append(numbers, 3)
   numbers = append(numbers, 4)
   fmt.Printf("numbers len = %d, cap=%d,slice=%v \n", len(numbers), cap(numbers), numbers)
}
```

### 切片的截取与复制

```go
func main() {
   s := []int{1, 2, 3}

   //截取操作，s1 指向的内存值和 s是一致的
   s1 := s[0:2]

   //如果s1修改了，那么s的值也会被修改
   s1[0] = 55

   fmt.Println(s)
   fmt.Println(s1)

   //copy可以将底层数组的slice一起进行拷贝
   s2 := make([]int, 3)
   copy(s2, s)
   fmt.Println(s2)
}
```







## 6.结构体

```go
//给类型取别名
type myint int

//定义结构体
type Book struct {
   title string
   price float32
}

func changeBook(book Book) {
   //传递的是book的一个副本
   book.price *= 100
}

func changeBook2(book *Book) {
   //传递的是book的一个指针
   book.price *= 100
}

func main() {

   var book1 Book //声明一个Book类型
   book1.title = "Golang"
   book1.price = 12.34

   fmt.Println(book1)
   changeBook(book1)
   fmt.Println(book1)
   changeBook2(&book1)
   fmt.Println(book1)
}
```





## 7.运算符

>  Go 语言内置的运算符有：
>
>  - 算术运算符
>  - 关系运算符
>  - 逻辑运算符
>  - 位运算符
>  - 赋值运算符
>  - 其他运算符

### 算术运算符

| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

### 关系运算符

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |

### 逻辑运算符

| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True  |

```go
func main() {
   var a bool = true
   var b bool = false
   if ( a && b ) {
      fmt.Printf("第一行 - 条件为 true\n" )
   }
   if ( a || b ) {
      fmt.Printf("第二行 - 条件为 true\n" )
   }
   /* 修改 a 和 b 的值 */
   a = false
   b = true
   if ( a && b ) {
      fmt.Printf("第三行 - 条件为 true\n" )
   } else {
      fmt.Printf("第三行 - 条件为 false\n" )
   }
   if ( !(a && b) ) {
      fmt.Printf("第四行 - 条件为 true\n" )
   }
}
```

### 位运算符

> 位运算符对整数在内存中的二进制位进行操作

```go
A = 0011 1100 //60

B = 0000 1101 //13
```

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | **按位与运算符**"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | **按位或运算符**"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | **按位异或运算符**"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，**当两对应的二进位相异时，结果为1**。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | **左移运算符**"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | **右移运算符**">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2 结果为 15 ，二进制为 0000 1111  |

### 赋值运算符

| =    | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| ---- | ---------------------------------------------- | ------------------------------------- |
| +=   | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=   | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=   | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=   | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=   | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=  | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=  | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=   | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=   | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=  | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |

### 其他运算符

| 运算符 | 描述             | 实例                       |
| :----- | :--------------- | :------------------------- |
| &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| *      | 指针变量。       | *a; 是一个指针变量         |

```go
package main

import "fmt"

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int //声明一个int类型的指针

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("*ptr 为 %d\n", *ptr);
}
```

### 运算符优先级

> 有些运算符拥有较高的优先级，二元运算符的运算方向均是从左至右。
>
> 下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

## 8.条件语句

| [if 语句](https://www.runoob.com/go/go-if-statement.html)    | **if 语句** 由一个布尔表达式后紧跟一个或多个语句组成。       |
| ------------------------------------------------------------ | :----------------------------------------------------------- |
| [if...else 语句](https://www.runoob.com/go/go-if-else-statement.html) | **if 语句** 后可以使用可选的 **else 语句**, else 语句中的表达式在布尔表达式为 false 时执行。 |
| [if 嵌套语句](https://www.runoob.com/go/go-nested-if-statements.html) | 你可以在 **if** 或 **else if** 语句中嵌入一个或多个 **if** 或 **else if** 语句。 |
| [switch 语句](https://www.runoob.com/go/go-switch-statement.html) | **switch** 语句用于基于不同条件执行不同动作。                |
| [select 语句](https://www.runoob.com/go/go-select-statement.html) | **select** 语句类似于 **switch** 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

> 注意：Go 没有三目运算符，所以不支持 **?:** 形式的条件判断。



## 9. 循环语句

| 循环类型                                                   | 描述                                 |
| :--------------------------------------------------------- | :----------------------------------- |
| [for 循环](https://www.runoob.com/go/go-for-loop.html)     | 重复执行语句块                       |
| [循环嵌套](https://www.runoob.com/go/go-nested-loops.html) | 在 for 循环中嵌套一个或多个 for 循环 |

#### 循环控制语句

| 控制语句                                                     | 描述                                             |
| :----------------------------------------------------------- | :----------------------------------------------- |
| [break 语句](https://www.runoob.com/go/go-break-statement.html) | 经常用于中断当前 for 循环或跳出 switch 语句      |
| [continue 语句](https://www.runoob.com/go/go-continue-statement.html) | 跳过当前循环的剩余语句，然后继续进行下一轮循环。 |
| [goto 语句](https://www.runoob.com/go/go-goto-statement.html) | 将控制转移到被标记的语句。                       |



## 10.函数

> Go 语言最少有个 main() 函数
>
> 函数声明告诉了编译器函数的名称，返回类型，和参数
>
> Go 语言标准库提供了多种可动用的内置的函数。例如，len() 函数可以接受不同类型参数并返回该类型的长度。如果我们传入的是字符串则返回字符串的长度，如果传入的是数组，则返回数组中包含的元素个数。

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

```go
//返回类型为int
func fool(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

func main() {
	c := fool("abc", 600)

	fmt.Println("c=", c)
}

```

### 1）多返回值函数

```go
//返回类型为int
func foo1(a string, b int) int {
   fmt.Println("a=", a)
   fmt.Println("b=", b)
   c := 100
   return c
}

//多返回值,匿名返回两个int值
func foo2(a string) (int, int) {
   return 1, 2
}

//多返回值,有形参名称
func foo3(a string) (r1 int, r2 int) {
   r1 = 3 //对要返回的形参进行赋值,没有赋值之前是类型默认值
   r2 = 4
   // return 3,4
   return //返回形参
}

//多返回值,有形参名称,多个返回类型相同的情况
func foo4(a string) (r1, r2 int) {
   return 3, 4
}

func main() {
   c := foo1("abc", 600)

   fmt.Println("c=", c)

   r1, r2 := foo2("hello")
   fmt.Println(r1, r2)

   r3, r4 := foo3("hi")
   fmt.Println(r3, r4)
}
```

### 2）包管理 go mod

> - pakage不局限于一个文件，可以由多个文件组成
> - 不要求package的名称和所在目录名称相同，但是最好保持相同，否则容易引起歧义
> - 每个子目录只能存在一个package，否则编译时会报错
> - package是以绝对路径GOPATH来寻址，不要用相对路径来import

```sh
#在项目根目录下生成go.mod文件
god mod init 
```

#### 包的可见性

标识符的`首字母大写`就可以让标识符对外可见

```go
package main

import (
	"04func/lib1"
	mylib2 "04func/lib2" //给包起别名，可以使用.别名，这样调用的时候就可以不使用名称来调用，而是直接调用方法
)

func main() {
	lib1.A()
	lib1.B()

	//使用包别名调用方法
	mylib2.C()
}
```

## 11.指针

> 一个指针变量指向了一个值的内存地址

### 声明指针变量

```go
var var_name *var-type
```

`var-type` 为指针类型，`var_name` 为指针变量名，`*` 号用于指定变量是作为一个指针。以下是有效的指针声明：

```go
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

### 使用指针变量

```go
func main() {

	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

}
```

### 空指针 nil

> 当一个指针被定义后没有分配到任何变量时，它的值为 `nil`。
>
> nil 指针也称为空指针。
>
> nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

```go
var ip *int    /* 声明指针变量 */

	//判断空指针
	if ip == nil {
		fmt.Println("ip是空指针!")
	}
```

### 数组指针

数组指针，一个指针，数据类型为数组（指向数组

```go
/*1.数组指针，一个指针，数据类型为数组（指向数组）*/
//定义一个数组
var arr = [3]int{10, 200, 3000}
//定义数组指针
var arrPt *[3]int
//将数组的地址赋值给指针变量
arrPt = &arr
fmt.Println("arrpt=", arrPt) //arrpt= &[10 200 3000]
fmt.Println("&arr=", &arr)   //输出内容等同于上面
//输出地址的形式
fmt.Printf("arr 数组的地址为：%p\n", &arr)
fmt.Printf("arrPtr 存储的地址为：%p\n", arrPt)
fmt.Println("*arrpt=", *arrPt) //输出的是数组的值

//通过指针访问数组
fmt.Println((*arrPt)[0]) //为了保证寻址运算符号* 优先于 [] ,使用括号
```

### 指针数组

指针数组，一个数组，里面的使用元素都为地址值

```go
/*2.指针数组，一个数组，里面的使用元素都为地址值*/
a, b, c, d := 1, 2, 3, 4
var ptrArr [4]*int               //声明一个指针数组，里面存放的都是指针
ptrArr = [4]*int{&a, &b, &c, &d} //初始化
fmt.Println(ptrArr)              //输出的是数组里面每个元素的地址
fmt.Println(*ptrArr[0]) //取值
```

### 函数的指针参数

```go
//传递int类型指针变量
func changeValue(p *int) {
	*p = 10
}

func main() {
	var a int = 20
  //传递地址
	changeValue(&a)
	fmt.Println(a)
}
```

### 指向指针的指针

```go
func main() {

	var a int = 10

	//指针变量，指向a的地址
	var pt *int = &a

	//二级指针
	var pp **int
	pp = &pt //存放pt指针变量的地址

	fmt.Println(&pt) //pt指针变量的地址
	fmt.Println(pp)

	fmt.Println(*pt)  //取pt存放地址指向的值
	fmt.Println(*pp)  // *pp 得到的是 pt指针变量
	fmt.Println(**pp) // **pp 得到的是 pt指针变量指向的值

}
```

## defer 语句

> defer关键字类似于c++的析构，在方法执行结束的时候执行
>
> defer与return谁先执行呢?

```go
package main

import "fmt"

//defer 和 return 谁先执行

func deferFunc() int {
   fmt.Println("defer func call")
   return 0
}

func returnFunc() int {
   fmt.Println("rerutn func call")
   return 0
}

func returnAndDefer() int {
   defer deferFunc()
   return returnFunc()
}

func main() {
   //当有return和defer的时候，先执行return 再执行 defer
   returnAndDefer()
}
```









## 13.范围 Range









## 14.集合 Map

### map的声明方式

```go
func main() {

   //=====> 第一种声明方式

   //声明myMap1是一个map类型，key是string，value是string
   var myMap1 map[string]string

   if myMap1 == nil {
      fmt.Println("myMap1 是一个空map!")
   }

   //使用map前，需要使用make给map分配空间
   myMap1 = make(map[string]string, 10)

   //给map变量赋值
   myMap1["one"] = "java"
   myMap1["two"] = "c++"
   myMap1["three"] = "python"

   fmt.Println(myMap1)

   // ==> 第二种声明方式
   myMap2 := make(map[string]int)
   myMap2["java"] = 1
   myMap2["c++"] = 2
   myMap2["python"] = 3
   fmt.Println(myMap2)

   // ==> 第三种声明方式
   myMap3 := map[string]int{
      "one":   1,
      "two":   2,
      "three": 3,
   }
   fmt.Println(myMap3)
}
```

### map的操作

```go
//map传参使用的是引用传递
func modifyMap(cityMap map[string]string) {
	cityMap["England"] = "London"
}

func main() {
	cityMap := make(map[string]string)

	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历 无序的
	for key, value := range cityMap {
		fmt.Println("key =", key)
		fmt.Println("value = ", value)
	}

	//修改
	cityMap["USA"] = "DC"

	fmt.Println("cityMap: ", cityMap)

	//删除
	delete(cityMap, "China")

	//访问
	fmt.Println("cityMap[\"China\"]: ", cityMap["China"])

	modifyMap(cityMap)

	fmt.Println(cityMap)

}
```

## interface

### 万能类型

### 类型断言

## 反射

通过变量得到变量的具体数据类型

![image-20220516114038792](https://gitee.com/ChetWei/img/raw/master/img/202205161140850.png)

### pair 的概念

```go
//每个类型都有一个pair的概念，用来标记类型和值

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

//具体类型
type Book struct {
}

//实现Reader接口
func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

//实现Write接口
func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {
	//b: pair<type:Book,value:book{}的地址>
	b := &Book{}

	//声明一个 Reader接口变量
	//r: pair<type,value> 暂时为空
	var r Reader
	//r: pair<type:Book,value:book{}的地址>
	r = b //指向Book具体类，因为Book实现了Reader的接口
	r.ReadBook()

	//声明一个 Writer接口变量
	var w Writer
	//r: pair<type:Book,value:book{}的地址>
	w = r.(Writer) //这里对r 进行断言为Writer，因为r和w具体的type是一致的都是Book
	w.WriteBook()

}
```

### 通过反射获取类型的信息

#### reflect包的使用

```go
import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//类里面的方法
func (this *User) Call() {

	fmt.Println("user is called ..")
	fmt.Printf("%v \n", this)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is:", inputType)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is:", inputValue)

	inpuNumField := inputType.NumField() //类里的字段数量
	fmt.Println("inpuNumField is:", inpuNumField)

	//通过type获取里面的字段
	for i := 0; i < inpuNumField; i++ {
		field := inputType.Field(i)              //获取这个字段的信息
		value := inputValue.Field(i).Interface() //获取这个字段的具体实现值
		//字段的命名 field.Name
		//字段类型 field.Type
		fmt.Printf("%s : %v = %v \n", field.Name, field.Type, value)
	}
}

func main() {
	user := User{1, "Tom", 18}

	DoFiledAndMethod(user)
}
```

### 结构体标签

```go
import (
	"fmt"
	"reflect"
)

//结构体标签
type resume struct {
	//添加字段标签
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(arg interface{}) {
	t := reflect.TypeOf(arg).Elem()

	for i := 0; i < t.NumField(); i++ {
		//获取tag里面的 info对应的值
		taginfo := t.Field(i).Tag.Get("info")
		tagdoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info: ", taginfo, "| doc: ", tagdoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
```

### 结构体标签在json中的应用

```go
import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangbozhi"}}

	//结构体 -> json  编码过程
	//将结构体里面的字段名称与对应的json标签进行转换
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr= %s\n", jsonStr)

	//jsonstr -> 结构体 解码过程
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json ummashal", err)
		return
	}

	fmt.Printf("%v\n", my_movie)
}

```



## 面向对象特征

> golang中并没有明确的面向对象的说法，实在要扯上的话，可以将struct比作其它语言中的class

### 1）封装

>类名、属性名、方法名
>
>首字母大写表示对外（其他包可以访问，否则只能够在本包内访问）

```go
//定义一个Hero类
type Hero struct {
	Name  string
	Ad    int
	Level int
}

//给结构体绑定方法
func (this Hero) showName() {
	//this 是调用该方法的对象的一个副本（拷贝）
	fmt.Println("Name =", this.Name)
}

func (this *Hero) GetName() string {
	//this 是调用该方法的对象的一个引用
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	//创建一个对象
	hero := Hero{"zhangsan", 100, 1}

	hero.showName()

	hero.SetName("lisi")

	hero.showName()

	name := hero.GetName()
	fmt.Println(name)

}
```

### 2）继承

```go
//结构体类的继承
type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

type SuperMan struct {
	Human  //SupeMan类继承了Human类的成员和方法
	height float32
}

//重新定义父类的方法Eat
func (this *SuperMan) Eat() {
	fmt.Println("Superman.Eat()...")
}

//子类的新方法
func (this *SuperMan) Walk() {
	fmt.Println("Superman.Walk")
}

func main() {
	human := Human{"zhangsan", "female"}
	human.Eat()

	superMan := SuperMan{human, 160}
	superMan2 := SuperMan{Human{"lisi", "male"}, 180}

	superMan.Eat()
	superMan.Walk()

	superMan2.Eat()
	superMan2.Walk()
}
```



### 3）多态

```go
//定义一个接口，本质上是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色

}

//具体类1 实现所有的接口
type Cat struct {
	color string //猫的颜色
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleep...")
}

func (this *Cat) GetColor() string {
	return this.color
}

//具体类2
type Dog struct {
	color string //猫的颜色
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is sleep...")
}

func (this *Dog) GetColor() string {
	return this.color
}


//多态
func showAnimal(animal AnimalIF) {
	//传递进来的是引用
	animal.Sleep() //多态，会根据类的具体类型，调用具体的实现方法
	fmt.Println("color=", animal.GetColor())
}
```

```go

func main() {
  // >>> 直接通过接口指针进行具体类的调用
	/*	var animal AnimalIF //接口数据类型，父类指针

		animal = &Cat{"Green"}
		animal.Sleep()

		animal = &Dog{"yellow"}
		animal.Sleep()*/

  //使用多态
	cat := Cat{"green"}
	dog := Dog{"yellow"}
	showAnimal(&cat) //传递地址进去
	showAnimal(&dog)
}
```

#### 万能的数据类型 interface{}

```go
//万能的 数据类型

//interface{}是万能的数据类型 所有的基本数据类型都实现了interface
func myFunc(arg interface{}) {
   fmt.Println("myFunc is called..")
   fmt.Println(arg)

   //使用 interface{} 如何区分 引用的底层数据类型是什么？

   // 给interface{} 提供"类型断言"的机制
   value, ok := arg.(string) //如果是string ok=true
   if ok {
      fmt.Println("arg is string type,value= ", value)
   } else {
      fmt.Println("arg is not string type")
      fmt.Printf("value type is %T \n", value)
   }
}

type Book struct {
   title string
}

func main() {
   book := Book{"Golang"}

   myFunc(book)
   myFunc(100)
   myFunc("abc")
   myFunc(3.14)
  
  a := "abcd"
	var allType interface{} //万能数据类型
	allType = a
	fmt.Println(allType)
}
```



## 协程

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161856260.png" alt="image-20220516185631170" style="zoom:30%;" />



### Golang对协程的处理

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161857186.png" alt="image-20220516185734096" style="zoom:30%;" />



#### 早期的golang对协程的处理

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161858704.png" style="zoom:33%;" />

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161859960.png" alt="image-20220516185941872" style="zoom:33%;" />

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161901028.png" alt="image-20220516190100987" style="zoom:33%;" />

#### GMP

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161902026.png" alt="image-20220516190234945" style="zoom:33%;" />

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161907840.png" alt="image-20220516190747759" style="zoom:33%;" />

在执行的协程数量一Processor的数量为准

#### 调度器的设计策略

- 复用线程

  - work stealing 机制	

  - hand off机制

- 利用并行

  - GOMAXPROCS 限定P的个数 = CPU核数/2

- 抢占

- 全局G队列

##### 复用线程

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161912595.png" alt="image-20220516191206514" style="zoom:33%;" />





>当process的队列空闲的时候，从别的process的队列中偷取goroutine任务

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161912324.png" alt="image-20220516191240250" style="zoom:33%;" />

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161914406.png" alt="image-20220516191420314" style="zoom:33%;" />

> 当检测到process 出现阻塞，数量不够用，会创建/唤醒一个新的线程

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161915324.png" alt="image-20220516191545226" style="zoom:33%;" />

##### 利用并行



##### 抢占

> 每个任务有限定的时间片，超过时间可以被抢占

<img src="https://gitee.com/ChetWei/img/raw/master/img/202205161921971.png" alt="image-20220516192107893" style="zoom:30%;" />

##### 全局队列

当process队列中的goroutine任务都为空的时候，可以从全局队列中进行任务偷取



### 使用go协程

```go
func task1() {
	i := 0
	for {
		i++
		fmt.Println("task1 : i=", i)
		time.Sleep(1 * time.Second)
	}

}

func main() {
	//创建一个go协程，去执行task1的方法
	go task1()

	//主进程
	i := 0
	for {
		i++
		fmt.Println("main task : i=", i)
		time.Sleep(1 * time.Second)
	}
}
```

匿名协程函数，如何退出协程

```go
func main() {

	//go协程之间的执行是无序的

	//使用go协程 定义并调用匿名函数 形参为空，返回值为空
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//退出当前的go协程
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//带有形参的go匿名函数
	go func(a int, b int) bool {
		fmt.Println("a=", a, ",b=", b)
		return true
	}(10, 20)

	//死循环
	for {
		time.Sleep(1 * time.Second)
	}
}
```



### go协程之间的通信 channel

```go
func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine函数")
		fmt.Println("goroutine 运行")
		//发送数据到channel
		c <- 666
	}()
	// c 进行阻塞，等待取数据
	//从c中接受数据，赋值给num
	num := <-c
	fmt.Println("num=", num)

}
```

#### 无缓冲的channel

![image-20220516205700068](https://gitee.com/ChetWei/img/raw/master/img/202205162057122.png)

第一步：两个goroutine都到达通道，但都没有开始执行发送或接受

第二步：左侧的goroutine将它的手伸进通道，这模拟了向通道发送数据的行为。这时，这两个goroutine会在通道中被锁住，直到交换完成。

第三步：右侧的goroutine将它的手放入通道，模拟了从通道接受数据的行为。这个goroutine一样会在通道中被锁住，直到交换完成。

第四步，第五步，进行交换，最终，在第6步，两个goroutine都将他们的手从通道里拿出来，这模拟了两个被锁住的goroutine得到了释放。

#### 有缓存的channel

![image-20220516210454787](https://gitee.com/ChetWei/img/raw/master/img/202205162104862.png)

> 有缓冲的通道，goroutine将数据放入之后，就可以去做别的事情，通道相当于队列，是有空间的

> - 当channel为空，从里面取数据也会阻塞
>
> - 当缓冲队列中的容量满了之后，再放入数据，会进入阻塞，等待被缓冲区内容被取走，再放入

```go

func main() {
	//创建一个缓冲为3的通道
	c := make(chan int, 3)
	//放入的元素数量,开辟大小
	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("子go程运行")
		//当缓冲队列中的容量满了之后，会进入阻塞，等待被缓冲区内容被取走，再放入
		for i := 0; i < 4; i++ {
			c <- i + 1
			fmt.Println("发送元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))
		}
	}()

	//先睡两秒，让子go程发送完成
	time.Sleep(time.Second * 2)

	//当channel为空，从里面取数据也会阻塞
	for i := 0; i < 4; i++ {
		//从c中接受并赋值
		num := <-c
		fmt.Println("接受 num=", num)
	}

	fmt.Println("main 结束")
}
```

#### 关闭channel

> - channel 不用向文件那样去关闭，只有当没有数据发送才去关闭
>
> - 关闭channel之后，无法向channe发送数据，引发panic错误后导致立即返回零值
> - 关闭channel后，可以继续接受数据
> - 对于nil channel（`没有通过make来定义的`），无论收发都会被阻塞 

```go
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭channel
		close(c)
	}()

	for {
		//ok如果为true表示channel没有关闭,否则关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			//通道关闭退出循环
			break
		}
	}
	fmt.Println("main finished ..")
}
```

#### channel与range

> 使用range`循环`读取channel中的数据
>
> 当channel中没有数据时候等待，当channel关闭的时候退出循环

```go
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//关闭channel
		close(c)
	}()

	//使用range迭代读取channel，当channel中没有数据时候等待，当channel关闭的时候退出
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("main finished")
}
```

#### channel与select

> 单流程下一个go只能监控一个channel的状态。
>
> select可以完成监控多个channel的状态

```go
//将channel变量作为参数
func fibonacii(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		//使用select
		select {
		case c <- x:
			//如果检测到c可写，则这个case就会进来
			t := x
			x = y
			y = t + y

		case data := <-quit:
			//检测到quit可读
			fmt.Println("quit", data)
			return
		}
	}
}

func main() {
	//创建两个channel
	c := make(chan int)
	quit := make(chan int)

	//子程
	go func() {
		for i := 0; i < 10; i++ {
			//读取
			fmt.Println(<-c)
		}
		quit <- 0 //写入
	}()

	//main
	fibonacii(c, quit)
}
```



## GoModule管理项目模块、依赖管理

> GoModule的出现是为了解决GOPATH的不足之处

```
查看GOPATH的信息
go env
```

GOPATH是用户的工作目录，包含三个目录

- pkg 第三方的一些库
- src 源码
- bin 包含一些编译好的可执行程序



GOPATH的弊端

- 无版本控制的概念

  ```
  go get 安装库  无法指定版本
  ```

  

- 无法与第三方版本号同步一致



### GO Modules

#### go mod 命令

```
go mod init  生成go.mod文件
go mod download 下载go.mod 文件中指明的依赖
go mod tidy 整理现有的依赖
go mod graph  查看现有的依赖结构
go mod edit	编辑go.mod文件
go mod vendor 导出项目所有的依赖到vendor目录
go mod verfiy 校验一个模块是否被篡改过
go mod why 查看问什么需要依赖某模块
```

> GO 语言提供了 GO111MODULE这个环境变量来控制Go Module的开关,可以设置下面的参数：
>
> - auto 只要项目包含了go.mod文件的话就启用Go Modlue 
> - on
> - off

```sh
go env -w GO1111MODULE=on
```



#### GOPROXY
