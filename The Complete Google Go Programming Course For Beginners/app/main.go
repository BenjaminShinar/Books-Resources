package main

import "fmt"

func main(){
    fmt.Println("Hello world!")
    var a[8] int // array of 8 integers
    x:= a[2:4] // slice of the array,
    fmt.Println(len(x))// legnth is 2(a[2],a[3]),
    fmt.Println(cap(x)) //but capacity is 6(a[2],a[3],a[4],a[5],a[6],a[7])

    x[0]=7
    fmt.Println(a[2]) //7
    z:=5
    z++
    fmt.Println("prefix",z)
    //++z not a thing!
    //fmt.Println("suffix",z)
}