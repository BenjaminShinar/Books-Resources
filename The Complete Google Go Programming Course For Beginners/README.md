<!--
ignore these words in spell check for this file
// cSpell:ignore Doxsey IntelliJ cmplx Expresso Elven
-->

# The Complete Google Go Programming Course For Beginners

for the Udemy course [The Complete Google Go Programming Course For Beginners](https://www.udemy.com/course/draft/1020950/learn/lecture/6134450)

- [golang](https://go.dev/learn/)
- [dockerhub go image](https://hub.docker.com/_/golang)
- [Joe Prays academy](https://www.joeparys.com/)

## Introduction

<details>
<summary>
Course introduction
</summary>

```go
package main

import "fmt"
func main(){
    fmt.Println("Hello world!")
    }
```

developed by google, it was designed to work in a multicore environment. it started as system based language, but it's general purpose, it's a compiled strongly typed language.

Projects using GO

- Docker and Kubernetes.
- Dropbox
- Twitter
- [and many more!](https://github.com/golang/go/wiki/GoUsers)

</details>

## Getting Started

<details>
<summary>
First steps in Go
</summary>

### Go Parachute- helpful resources

- [Cheat Sheet](https://github.com/a8m/golang-cheat-sheet)
- [Official Syntax Specifications](https://go.dev/ref/spec)
- ["How to write go code" starting point](https://go.dev/doc/code)
- [An introduction to programming in go (Caleb Doxsey)](https://www.golang-book.com/books/intro)
- [Example annotated code](https://gobyexample.com/)

### installing Go

`docker container run --rm -v ${pwd}/app:/usr/src/app -w /usr/src/app golang:1.17 go run main.go`

going to the website and installing.

### Hello World

we write our first program, and then compile and run it.

```sh
go run helloworld.go
```

### Installing an IDE

IDE - integrated development environment.
this course recommends IntelliJ IDEA, we need GIT and to configure them. it later moved on the use VSCODE.

### Installing Visual Studio

nothing serious, just getting the ide and the extensions. basic debugging tools, etc.

we can't have two `func main()` functions from the same package in the same folder.

- </details>

## Fundmental Google GO

<details>
<summary>
The basic of programming.
</summary>

### What is programming

> - A computer is a tool that can store, retrive, process and transmit data.
> - Programming is what people do to "teach" a computer how to do this/
> - Programming languages are the tool to do some.

binary numbers, base two. hexadecimal numbers (0-9a-f),8 bit = 1 byte, values between 0-255.

basic parts of the computer.

- Cpu - Cental processing unit, all binary, all numbers, only machine langues. modern computes have multi-core cpu.
- RAM - random acesss memory, volatile.
- Storage - persistent.

assembly languages, instructions. high level programming languages as opposed to machine language. compiled langauges vs run-time interpreted languages.

defintion and assignment.

### Understanding the Hello World Program - "No chicken or egg for us!"

```go
package main

import "fmt"

func main(){
    fmt.Println("Hello world!")
}
```

go is case sensitive, and extremely petty with where the curly braces are. but not about white spacelines.

the `package` is a way to organize code, `import` using other libraries/projects. function declaration and function signature, optional return type.

### Variables and Constants

Decelerations, assignments, initialization, data types.

var keyword, identifier, type.

```go
import "math/cmplx"

func main() {
var a int32
a = 15
var b bool
b=false
var f float32
f=15.0
var s string
s="string"
var c complex128
c=cmplx.Acos(-2+12i)
}
```

- bool
- string
- integer: int8, int16, int32, int64
- unsigned integer: uint8, uint16, uint32, uint64
- floating point: float32, float64
- complexnumber: complex64,complex64

string literals.

1. we get an error if we declare a varable and don't use it.
2. variables are initialized to zero, false, and empty string.
3. there are no implicit conversions between data types.

```go
var a int32
var b float64
a = 15
b=a //error!
```

```go
var a int32 //decleration without initialization
var b int32 = 10
var c,d int35 =15,16 //legal, but weird.
var e = 20 // inferred type.
z:=true // declare and initialize
const r int32 = 9 //const, unchangeable
```

### Expressions - Expresso? No... expressions - super important core concept here!

go is statically typed, no implicit conversions. expressions result in data, there is an order of evaluation based on the kind of the operators.

- general arithmetic operators: addition, subtraction, multiplication, division, reminder, and also string concatenation.

explicit casting with the type name.

```go
func main(){
    var i int32
    i=int32(math.Sqrt(15.0))
    fmt.Println("Value of i:",i)
}
```

comparison operator, logical boolean operators.

</details>

## Intermediate Beginers Google Go

<details>
<summary>
Really basic stuff
</summary>

### Loops

```go
package main

import "fmt"
import "time"

func main(){

    var i:=10
    for i>0{
        fmt.Println(i)
        time.Sleep(time.Second)
        i=i-1
    }
    fmt.Println("Happy new year!")
}
```

for loops have an initialization stage, the conditional expression, and the post statement, and of course, the repeated code itself.

types of loops:

```go
for {
    //infinite loop
}

for a>b {
    //while loop
}

for i:=0;i<10;i++{
    //for loop
}
```

`break`,`continue` are also parts of loops.

```go
package main

import "fmt"
import "time"

func main(){

    for var i:=10;i>0;--i{
        fmt.Println(i)
        time.Sleep(time.Second)
    }
    fmt.Println("Happy new year!")
}
```

the code block between the curly braces defines a scope.

```go
package main

import "fmt"

func main(){

    for var i:=0;i<10;i++{
        j:=15
        j++
    }
    fmt.Println(j) //error!
}
```

### Control Flow

do something based on a condition

```go
package main

import "fmt"
func main(){
    age:=10
    if age <12 {
        fmt.Println("you are a child!")
    }
    if age <20 && age >=12 {
        fmt.Println("you are a teen!")
    }
}
```

we also have `if - else if - else`. the curly braces location matters!

```go
package main

import "fmt"
func main(){
    age:=10
    if age < 12 {
        fmt.Println("you are a child!")
    } else if age < 20  {
        fmt.Println("you are a teen!")
    } else {
        fmt.Println("you are an adult!")
    }
}
```

### Switch Statement

lets take this if-else program

```go
package main

import "fmt"
func main(){

    for d:=1; d<=12;d++{
        fmt.Println("On the %d day of Christmax, my true love sent to me",d)
        if day == 12 {
            fmt.Println("Twelve curly braces")
        } else if day == 11 {
            fmt.Println("Eleven Elven wenches")
        }


    }
}
```

and make it into a switch statement, there is no fall through behavior.

```go
package main

import "fmt"
func main(){

    for d:=1; d<=12;d++{
        fmt.Println("On the %d day of Christmax, my true love sent to me",d)

        switch d{
            case 12:
                fmt.Println("Twelve curly braces")
            case 11:
                fmt.Println("Eleven Elven wenches")
            default:
        }
    }
}
```

```go
package main

import "fmt"
func main(){

    for d:=1; d<=12;d++{
        fmt.Println("On the %d day of Christmax, my true love sent to me",d)

        switch d{
            case 12:
                fmt.Println("Twelve curly braces")
                fallthrough
            case 11:
                fmt.Println("Eleven Elven wenches")
                fallthrough
            default:
        }
    }
}
```

actually, the default case doesn't have to the last one, if we decide for some reason that we want to fall through it.

### Functions

passing by value (on the stack), reusable code.

```go
func name(value1 type1,value2 type2) return_type{
    //code
}
```

there is something called named return.

references and pointers.

```go
func max(i int, j int)int{
    if i>j{
        fmt.Println(i)
        return i
    }else{
        fmt.Println(j)
        return j
    }
}
```

modifying data in function

```go
func doubleNumber(number *int){
    *number = *number*2
}

func main(){
    c:= 25
    doubleNumber(&c)
    fmt.Println(c)
}
```

### Understanding Scope

where variables are accessable. go is lexically scoped using blocks.

global scope, go doesn't raise an error about global unused variables.

```go
package main
import "fmt"
var exit bool=false

func testexit(){
    exit=true
}
```

we can't use the implicit assignment operator in global scope `:=`

</details>

## Advanced Beginers

<details>
<summary>
Some Go stuff.
</summary>

### Arrays

a container for data of the same type, arrays are index (starting at zero) and have a fixed length.

```go
var a[5] int; //array of size 5

a[0]=5;
b :=[5]int{0,3,6,9,12}

fmt.Println(b[0])
fmt.Println(b) //go prints all the elements of the array.
```

now an example of a grocery list

```go
package main

import "fmt"

const g_cap int = 5 //capacity of grocery list
var g_groceries[g_cap] string;
var g_len int=0 //items already in list
func add_grocery(item string){
    if g_len < g_cap{
        g_groceries[g_len]=item
        g_len++
    } else{
        fmt.Println("Array full!")
    }
}

func list_groceries(){
    fmt.Println("item in grocery list:");
    for i:=0;i<g_len;i++{
        fmt.Println(g_groceries[i])
    }
}

func main(){
    list_groceries()
    add_grocery("bread")
    list_groceries()
}
```

but actually, we almost never use raw arrays, we use slices instead.

### Slices

a slice is a segment of an array, or a window to array, it can be represented as a [low : high] slice of indexes, the low value is included, but not the high, so it's actually **[low:high)**. we can omit either the start or the end of the slice, which would default into the first (zero) element or the last. we can use `[:]` to choose all the elements of the array. Slices can't be outside the bounds of the array.

we can use the `make` command to create an array:\
`x:=make([]datatype, length, capacity)`

```go
var a[8] int // array of 8 integers
x:= a[2:4] // slice of the array,
fmt.Println(len(x))// legnth is 2(a[2],a[3]),
fmt.Println(cap(x))//but capacity is 6(a[2],a[3],a[4],a[5],a[6],a[7])

x[0]=7
fmt.Println(a[2]) //7
```

slices are references to arrays, and they are passed as references themselves. Golang simply does all the work behind the scenes.

```go
package main

import "fmt"

var g_groceries[] string;

func add_grocery(item string){
    fmt.Println("Capacity is %d",cap(g_groceries))
    g_groceries.append(g_groceries,a)
}

func list_groceries(){
    fmt.Println("item in grocery list:");
    for i:=0;i<len(g_groceries);i++{
        fmt.Println(g_groceries[i])
    }
}

func main(){
    list_groceries()
    add_grocery("bread")
    list_groceries()
}
```

the `append` function returns a new slice and the value must be assigned.

### Simple Statements

[Simple Statement in go](https://medium.com/golangspec/simple-statement-notion-in-go-b8afddfc7916)

Simple statements can be placed in some key places to belong to the coming block, such as for loops, if and switch blocks.
this is like initilazig a variable only for the current scope.

only a subset of go expressions are simple statements.

- Empty statement
- increments/decrement statements
- Assignment
- Short variable decleration
- Expression statements
- Send statements (through a channel)

we can use simple statements to extend the if statement. creating a variable the only exists for the scope of the if block.

```go
if <simple statement>; <boolean expression> {
    <code block>
}
```

from the documentation

```go
if err:=file.Chmod(0664); err !=nil{
    log.Print(err)
    return err
}
```

we can also use simple statements as part of switch statement.

```go
switch <simple statement>;<expression>{
    case <expression>:
        <code>
    case <expression>:
        <code>
    default:
        <code>
}
```

### For Range Loops

advanced for loops

```go
//basic loop
for i:=0; i<len(a);i++{
    fmt.Println("element at index ",i, "is value ",a[i])
}
//For Range Loop
for index, data:=range a{
    fmt.Println("element at index ",index, "is value ",data)
}
```

we can use `_` for the variable name if we aren't planning to use use

### Variadic Functions - No function ever sounded "so cool". Variadic...functions.

having functions that take any number of variables, rather than a separate function for each number.

this is tha basic way of writing non-variadic functions:

```go
func MaxOne(a1 int)int{
    return a1
}
func MaxTwo(a1 int,a2 int)int{
    if a1>a2 {
        return a1
    }
    return a2
}
func MaxThree(a1 int,a2 int,a3 int)int{
    //do something
}
```

but that is cumbersome, we can get all the elements as a slice by using a variadic function instead.

```go
func MaxMany(numbers ...int)int{
    //numbers is a slice!
}
```

lets modify our earlier function

```go
func add_grocery(items ...string){
    fmt.Println("Capacity",cap(g_groceries))
    for _,item:=range items{
        g_groceries.append(item)
    }
}
```

### Modifying the IDEA plugin

updating the golang plugin for IntellJ IDEA, the work around is changing the _plugin.xml_ file. we find the idea-version elements and modify the _until-build_ attribute.

```xml
<idea-plugin version="2">
<idea-version since-build="171.1834" until-build="172.*"/>
</idea-plugin>
```

### Next steps

- [Effective Go](https://golang.org/doc/effective_go.html) - writing good Go code.
- [Going Go Programming](https://www.goinggo.net/)
- [Awesome Go](http://awesome-go.com/) - examples of good GO frameworks, libraries, etc..
- [Writing Web Applications](https://golang.org/doc/articles/wiki/) - using the `net/http` package

</details>

## Takeaways

<details>
<summary>
things to remember
</summary>

There is no prefix increment operators! `++i` doesn't exist!

slices:

- `make([]datatype,length, capacity)` - create a slice
- `len(slice_name)` - length of slice.
- `cap(slice_name)` - capcity (depends on underlying array) from the start of the slice to the end of the array.
- `append(slice_name, item)` - add to slices, modify and reallocate array if needed. returns the new slice.

- `range <slice_name>` - no parentheses, enumerates over data with index,data.

</details>
