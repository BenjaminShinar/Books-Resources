<!--
ignore these words in spell check for this file
// cSpell:ignore Doxsey IntelliJ cmplx Expresso Elven
-->

# The Complete Google Go Programming Course For Beginners

for the Udemy course [The Complete Google Go Programming Course For Beginners](https://www.udemy.com/course/draft/1020950/learn/lecture/6134450)

[golang](https://go.dev/learn/)

[dockerhub go image](https://hub.docker.com/_/golang)

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

    for var i:=0;i<10;++i{
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

    for d:=1; d<=12;++d{
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

    for d:=1; d<=12;++d{
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

    for d:=1; d<=12;++d{
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

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### Arrays and Slices - Part 1 - Arrays... and seeing multiple!

### Arrays and Slices - Part 2 - Hands on Arrays

### Arrays and Slices - Part 3 - Slices - A slice of nice!

### Arrays and Slices - Part 4 - Hands on slices..... and the power within!

### Advanced Topics - Simple Statements (that aren't quite so simple....)

### For Range Loops - Processing forloops in a blink of an eye...

### Variadic Functions - No function ever sounded "so cool". Variadic...functions.

</details>

## Bonus Lecture

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>
