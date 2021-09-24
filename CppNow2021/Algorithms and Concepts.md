Algorithms

## Algorithms from a Compiler Developer's Toolbox - Gábor Horváth

<details>
<summary>
A bit of compiler algorithms for optimization, using identities and algebra.
</summary>

[Algorithms from a Compiler Developer's Toolbox](https://youtu.be/eeS1WP7FK-A),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/CompilerAlgorithmsTalk.pdf)

### Why Study Compiler?

there are many algorithms and data structure that are used in compilers, compilers are everywhere, like web browsers (html+css, svg, JavaScript of course), GPUs also have compilers, databases have compilers and optimizers, even some configuration format files have something like a compiler inside them. python compiles down to something for machine learning, and routers and modem have something running on them.

A lot of opportunities to improve code, any small improvement is multiplied because it effects every program compiled by it. if we improve a low level compiler (like c++), then we also effect any compiler that uses it (like python or JavaScript).

example: loop strength reduction.

video of a talk by matt godbolt. replacing sum by loop with sum by formula

$
\sum x \equiv \\\frac{x (x+1)}{2}
$

playing with loops kinds and looking at the assembly, we see that the compiler manages to remove the loop and figures out a closed-form formula.

but floating point messes up the optimizations, floating point arithemetic.

### What's Inside the Compiler?

Math.😅\
Chains of recurrences - recursive function when the increment is also a recursive function.

two kinds of recursive formula notations. making functions at incrementoars

$
f(i*i(n)) = {initial,+,incrememt }
$

algebra, operations identities. making loops into recursive notations with those identities, sub expressions combine together. turning this

```cpp
for(i=0;i<m;++i)
v[i] = (i+1)*(i*1) -i*i -2*i;
```

into a constant expression with identities.

$(x+y)(x+y) = (x+y)^2 = x^2 +2xy + y^2$

so we open up the identity, and we can then cancel out stuff and reach a constant.

```cpp
for(i=0;i<m;++i)
v[i]=1;
```

arithemetic series which are supposed to be loops can be made into closed formulas or at least have much less operations per loop

```cpp
int t[20];
for (int i=0;i< 20; i+=1)
{
    t[i]=(i+1)*(i+1) + 3*i - 5; // four additions, two multiplications
}
```

is transformed into this compact form with only two additions.

$
f_{(i+1)^2+3i-5}(n) = \\
\{-4,+6,+2\}
$

which is equivalent to writing this c++ code.

```cpp
int t[20];
int a = -4;
int b = 6;
for (int i=0;i< 20; i+=1)
{
    t[i]=a;
    a+=b;
    b+=2;
}
```

an example of how clang does it. we take this code into a file.

```cpp
int f(int num)
{
    int result =0;
    for (int i =0;i<num;++i)
    {
        result += i;
    }
    return result;
}
```

and then run the following command on it (replace $1 with file name)

```bash
clang++ $1.cpp -c -02 -Xclang -disable-llvm-passes -emit-llvm -S
opt $1.ll -mem2reg -S > {$1}2.ll
opt ${1}2.ll --analyze --scalar-evolution
(other)
```

- -Xclang \<arg> Pass \<arg> to the clang compiler.
- -disable-llvm-passes
- -emit-llvm Use the LLVM representation for assembler and object files
- -S preprocessor only

we can see in the slides how loops are eliminated.

> Recapping Chains of recurrences
>
> - Great to model some loop varian values.
> - Algebra of simple recursive function
> - Algebraic simplifications
> - Strength reduction
> - Closed forms
> - and many more...

### Value Numbering

eliminating some forms of redundancy.

this code has redundancy.

```cpp
int calculate(int a, int b)
{
    int result = (a * b) +2;
    if (a %2 ==0)
    {
        result +=a*b;
    }
    return result;
}
```

the compiler can do the common expression optimization in some cases. but most of the redundancy isn't from the programmer. this code had redundancy in terms of memory access;

```cpp
int matrix[5][5];
//...
matrix[1][2]=bar();
matrix[1][3]=baz();
```

is actually memory dereferencing with a common sub expression.

```cpp
int matrix[5][5];
//...
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 2)=bar();
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 3)=baz();
```

we can also have dead_code and unused code that passes around (constant propagation).
compilers work in phases, and at each pass the complier cleans up the code to make it optimize. each pass does a small change.

[BRIL - big red intermediate language](https://github.com/sampsyo/bril) is a compiler IR (Intermediate representation) that is used in some courses to teach about compilers.

optimizations can work across different scopes (function, loop body, and even higher!);

local value numbering optimization. algebraic identities, dead code elimination, constant folding,

### where to learn more

some sources to learn mode about compilers.
audience questions

</details>
