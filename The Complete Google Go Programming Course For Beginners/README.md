<!--
ignore these words in spell check for this file
// cSpell:ignore Doxsey IntelliJ cmplx Expresso
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

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Advanced Beginers

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>

## Bonus Lecture

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

</details>
