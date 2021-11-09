<!--
// cSpell:ignore goto gotos fmin lefticus xoroshiro structs nodiscard cerr rotr lippincott spdlog gdbgui kcov tandy jongg pscii endl constexpr decltype mutex noexcept struct smallpt
 -->

# C++ Weekly takeaways

what I learned from each episode.

[c++ Weekly Playlist](https//www.youtube.com/playlist?list=PLs3KjaCtOwSZ2tbuV1hx8Xz-rFZTan2J1)

[Best Practices Book](https//github.com/lefticus/cppBestPractices)

## C++ Weekly - Ep 1 - ChaiScript_Parser Initialization Refactor

## C++ Weekly - Ep 2 - Cost of Using Statics

## C++ Weekly - Ep 3 - Intro to clang-tidy

## C++ Weekly - Ep 4 - Variadic Template Refactor

## C++ Weekly - Ep 5 - Intro To ChaiScript

## C++ Weekly - Ep 6 - Intro To Variadic Templates

## C++ Weekly - Ep 7 - Stop Using std::endl

_std::endl_ also does flush.
This is bad. Use ‘\n’ instead.

## C++ Weekly - Ep 8 - C++ Name De-Mangling

## C++ Weekly - Ep 9 - std::future Quick-Start

## C++ Weekly - Ep 10 - Variadic Expansion Wrap-Up

## C++ Weekly - Ep 11 - std::future Part 2

## C++ Weekly - Ep 12 - C++17's std::any

## C++ Weekly - Ep 13 - Fibonacci You're Doing It Wrong

## C++ Weekly - Ep 14 - Standard Library Gems _next_ and _exchange_

## C++ Weekly - Ep 15 - Using _std::bind_

## C++ Weekly - Ep 16 - Avoiding _std::bind_

## C++ Weekly - Ep 17 - C++17's _std::invoke_

## C++ Weekly - Ep 18 - C++17's constexpr if

## C++ Weekly Special Edition - Using C++17's constexpr if

## C++ Weekly - Ep 19 - C++14 For The Commodore 64

## C++ Weekly - Ep 20 - C++17's Fold Expressions - Introduction

## C++ Weekly - Ep 21 - C++17's _if_ and _switch_ Init Statements

## C++ Weekly - Ep 22 - C++17's Nested Namespaces and _std::clamp_

## C++ Weekly - Ep 23 - C++17's \_\_has_include

## C++ Weekly - Ep 24 - C++17's Structured Bindings

## C++ Weekly - Ep 25 - C++17's Aggregate Initializations

## C++ Weekly - Ep 26 - Language Features Removed in C++17

## C++ Weekly - Ep 27 - C++17 Changes To Sequence Containers

## C++ Weekly - Ep 28 - C++17's \[\[fallthrough]] Attribute

[[fallthrough]]

Tells the compiler we know we have a fallthrough in our switch case.

## C++ Weekly - Ep 29 - C++17's \[\[maybe_unused]] Attribute

<details>
<summary>
[[maybe_unused]] attribute
</summary>
Means that this variable might not be used in all version (build vs debug), this will prevent the compiler from warning us.

</details>

## C++ Weekly - Ep 30 - C++17's \[\[nodiscard]] Attribute

## C++ Weekly - Ep 31 - IncludeOS

## C++ Weekly - Ep 32 - Lambdas For Free

## C++ Weekly - Ep 33 - Start Using Default Member Initializers

## C++ Weekly - Ep 34 - Reading Assembly Language - Part 1

## C++ Weekly - Ep 35 - Reading Assembly Language - Part 2

## C++ Weekly - Ep 36 - Reading Assembly Language - Part 3

## C++ Weekly - Ep 37 - Stateful Lambdas

## C++ Weekly - Ep 38 - C++17's Class Template Argument Type Deduction

## C++ Weekly - Ep 39 - C++17's Deduction Guides

## C++ Weekly - Ep 40 - Inheriting From Lambdas

## C++ Weekly - Ep 41 - C++17's constexpr Lambda Support

## C++ Weekly - Ep 42 - Clang's Heap Elision

## C++ Weekly - Ep 43 - Stack Elision?

## C++ Weekly - Ep 44 - constexpr Compile Time Random

## C++ Weekly - Ep 45 - Compile Time Maze Generator (and Solver)

## C++ Weekly - Ep 46 - Start Using Visual C++

## C++ Weekly - Ep 47 - My Visual C++ Pet Peeve

## C++ Weekly - Ep 48 - C++17's Variadic _using_

## C++ Weekly SE - Why -0xFFFFFFFF == 1

Numeric literals are unsigned by default;  
Minus prefix is not part of the literal, it’s an operator.  
Two complement - take every negate each byte, and then add 1.  
Unsigned will always be unsigned

## C++ Weekly - Ep 49 - Why Inherit From Lambdas?

## C++ Weekly - Ep 50 - Inheriting Lambdas vs Generic Lambdas

Variadic template <typename ..T>
Inheriting from lambdas with std::forward and std::decay and using
std::variant<double, int>
std::common_type<double,int>
Decltype(v)
Lambda capturing []
std::is_same<double, decltype(v)>
Constexpr
Std::visit

Std::common_type<decltype(somename)> varname = 0;

If (std::is_same<double,decltype(v)>{})
{
}
Else
{
}

## C++ Weekly - Ep 51 - Advanced Stateful Lambdas

Mutable
std::exchange(a,b);
Local class inside lambda, lambda capture list

Return a struct with references to the capture list, and have it be modifiable from outside

## C++ Weekly - Ep 52 - C++ To C Compilation

## C++ Weekly - Ep 53 - Gotos Are Everywhere

Continues are gotos
Loops are whiles
Whiles are gotos

## C++ Weekly - Ep 54 - Zero Cost Embedded C++ - Part 1 Interrupt vector

## C++ Weekly - Ep 55 - Zero Cost Embedded C++ - Part 2

## C++ Weekly - Ep 56 - Zero Cost Embedded C++ - Part 3

## C++ Weekly - Ep 57 - Dissecting An Optimization

## C++ Weekly - Ep 58 - Negative Cost Embedded C++ - Part 1

## C++ Weekly - Ep 59 - Negative Cost Embedded C++ - Part 2

## C++ Weekly - Ep 60 - std::quoted

## C++ Weekly - Ep 61 - Storage Duration with Lambdas

## C++ Weekly - Ep 62 - std::regex

## C++ Weekly - Ep 63 - What is Negative Zero?

## C++ Weekly - Ep 64 - C++11's std::min (and my version)

## C++ Weekly - Ep 65 - C++11's std::fmin

## C++ Weekly - Ep 66 - Variadic fmin for C++11

## C++ Weekly - Ep 67 - C++17's std::gcd and std::lcm

## C++ Weekly - Ep 68 - std::iota

## C++ Weekly - Ep 69 - C++17's Searchers

## C++ Weekly - Ep 70 - C++ IIFE in quick-bench.com

http://quick-bench.com/
A site that allows quick benchmarking of c++ code.

## C++ Weekly - Ep 71 - Hidden C++ 17 - Part 1

New stuff in c++17
Some regex features, ‘multiline’ option.
Shared_ptr with index
reinterpet_pointer_cast<> and shared_ptr/
Std::as_const - takes a reference and returns a const reference.
Std::to_chars - fills a string with a number at the locale format
Std::from_chars - takes a char\* (begin and end) and returns a numeric result.

## C++ Weekly - Ep 72 - Hidden C++17 - Part 2

Std::shared_mutex, Std::shared_timed_mutex (for reader / writer)
Math stuff: specialized stuff with overloads

## C++ Weekly - Ep 73 - std::string_view

Use std::string_view instead of const std::string &s.
If we don’t need any ‘std::string’ things and we just need to look at a string.
If we don’t change it, and it’s a string literal, a string view should be enough.

## C++ Weekly - Ep 74 - std::regex optimize

A flag that’s supposed to make regex faster.
Make sure to put the stuff where it needs to be scope.
You should limit the scope of the variables, but be aware of repeated creations which have serious costs.

## C++ Weekly - Ep 75 - Why You Cannot Move From Const

Move operations ‘create’ an rValue reference (&&)
A move from const is a const rValue reference

it might silently use a constructor rather than a move if the try to move from const objects

## C++ Weekly - Ep 76 - static_print

## C++ Weekly - Ep 77 - G++ 7.1 for DOS

## C++ Weekly - Ep 78 - Intro to CMake

## C++ Weekly - Ep 79 - Intro To Travis CI

## C++ Weekly - Ep 80 - Intro to AppVeyor

## C++ Weekly - Ep 81 - Basic Computer Architecture

## C++ Weekly - Ep 82 - Intro To CTest

## C++ Weekly - Ep 83 - Installing Compiler Explorer

## C++ Weekly - Ep 84 - C++ Sanitizers

## C++ Weekly - Ep 85 - Fuzz Testing

## C++ Weekly - Ep 87 - std::optional

## C++ Weekly - Ep 88 - Don't Forget About puts

## C++ Weekly - Ep 89 - Overusing Lambdas

## C++ Weekly - Ep 91 - Using Lippincott Functions

## C++ Weekly - Ep 92 - function-try-blocks

## C++ Weekly - Ep 93 - Custom Comparators for Containers

## C++ Weekly - Ep 94 - Lambdas as Comparators

## C++ Weekly - Ep 95 - Transparent Comparators

## C++ Weekly - Ep 96 - Transparent Lambda Comparators

## C++ Weekly - Ep 97 - Lambda To Function Pointer Conversion

## C++ Weekly - Ep 98 - Precision Loss with Accumulate

## C++ Weekly - Ep 99 - C++ 20's Default Bit-field Member Initializers

## C++ Weekly - Ep 100 - All The Assignment Operators

## C++ Weekly - Ep 101 - Learning "Modern" C++ - The Tools

## C++ Weekly - Ep 102 - Learning "Modern C++" - Hello World

## C++ Weekly - Ep 103 - Learning "Modern" C++ - Inheritance

## C++ Weekly - Ep 104 - Learning "Modern" C++ - 4 const and constexpr

## C++ Weekly - Ep 105 - Learning "Modern" C++ - Looping and Algorithms

## C++ Weekly - Ep 106 - Disabling Move From const

## C++ Weekly - Ep 107 - The Power of =delete

## C++ Weekly - Ep 108 - Understanding emplace_back

## C++ Weekly - Ep 109 - When noexcept Really Matters

## C++ Weekly - Ep 110 - gdbgui

## C++ Weekly - Ep 111 - kcov

## C++ Weekly - Ep 112 - GCC's Leaky Abstractions

## C++ Weekly - Ep 113 - Will It C++? Atari Touch Me From 1978

## C++ Weekly - Ep 114 - cpp_starter_project GUI Additions

## C++ Weekly - Ep 115 - Compile Time ARM Emulator

## C++ Weekly - Ep 116 - Trying Out The Conan Package Manager

## C++ Weekly - Ep 117 - Trying Out The Hunter Package Manager

## C++ Weekly - Ep 118 - Trying Out The vcpkg Package Manager

## C++ Weekly - Ep 119 - Negative Cost Structs

## C++ Weekly - Ep 120 - Will It C++? The Tandy 1000 From 1984

## C++ Weekly - Ep 121 - Strict Aliasing In The Real world

## C++ Weekly - Ep 122 - _constexpr_ With _optional_ And _variant_

## C++ Weekly - Ep 123 - Using in_place_t

## C++ Weekly - Ep 124 - ABM and BMI Instruction Sets

## C++ Weekly - Ep 125 - The Optimal Way To Return From A ## Function

## C++ Weekly - Ep 126 - Lambdas With Destructors

## C++ Weekly - Ep 127 - C++20's Designated Initializers

## C++ Weekly - Ep 128 - C++20's Template Syntax For Lambdas

## C++ Weekly - Ep 129 - The One Feature I'd Remove From C++

## C++ Weekly - Ep 130 - C++20's for init-statements

## C++ Weekly - Ep 131 - Literals in ARM Assembly

## C++ Weekly - Ep 132 - Lambdas In Fold Expressions

## C++ Weekly - Ep 133 - What Exactly IS A Lambda Anyhow?

## C++ Weekly - Ep 134 - The Best Possible Way To Create A Visitor?

## C++ Weekly - Ep 135 - {fmt} is Addictive! Using {fmt} and spdlog

## C++ Weekly - Ep 137 - C++ Is Not An Object Oriented Language

## C++ Weekly - Ep 138 - Will It C++? MIPS Architecture (1985)

## C++ Weekly - Ep 139 - References To Pointers

## C++ Weekly - Ep 140 - Use _cout_, _cerr_, and _clog_ Correctly

## C++ Weekly - Ep 141 - C++20's Designated Initializers And Lambdas

## C++ Weekly - Ep 142 - Short Circuiting With Logical Operators

## C++ Weekly - Ep 143 - GNU Function Attributes

## C++ Weekly - Ep 144 - Pure Functions in C++

## C++ Weekly - Ep 145 - Semi-Automatic _constexpr_ and _noexcept_

## C++ Weekly - Ep 146 - C++20's std::to_pointer

## C++ Weekly - Ep 147 - C++ And Python Tooling

## C++ Weekly - Ep 148 - clang-tidy Checks To Avoid

## C++ Weekly - Ep 149 - C++20's Lambda Usability Changes

## C++ Weekly - Ep 150 - C++20's Lambdas For Resource Management

## C++ Weekly - Ep 151 - C++20's Lambdas As Custom Comparators

## C++ Weekly - Ep 152 - Lambdas The Key To Understanding C++

## C++ Weekly - Ep 153 - 24-Core C++ Builds Using Spare Computers!

## C++ Weekly - Ep 154 - One Simple Trick For Reducing Code Bloat

## C++ Weekly - Ep 155 - Misuse of pure Function Attribute

## C++ Weekly - Ep 156 - A C++ Conference Near You

## C++ Weekly - Ep 157 - Never Overload Operator && or ||

## C++ Weekly - Ep 158 - Getting The Most Out Of Your CPU

## C++ Weekly - Ep 159 - _constexpr_ _virtual_ Members In C++20

## C++ Weekly - Ep 160 - Argument Dependent Lookup (ADL)

## C++ Weekly - Ep 161 - The C++ Box Project

## C++ Weekly - Ep 162 - Recursive Lambdas

## C++ Weekly - Ep 163 - Practicing ARM Assembly

## C++ Weekly - Ep 164 - Adding a Random Device To The C++ Box

## C++ Weekly - Ep 165 - C++20's is_constant_evaluated()

## C++ Weekly - Ep 166 - C++20's Uniform Container Erasure

## C++ Weekly - Ep 167 - What Is Variable Shadowing?

## C++ Weekly - Ep 168 - Discovering Warnings You Should Be Using

## C++ Weekly - Ep 169 - C++20 Aggregates With User Defined Constructors

## C++ Weekly - Ep 170 - C++17's _inline_ Variables

## C++ Weekly - Ep 171 - C++20's Parameter Packs In Captures

## C++ Weekly - Ep 172 - Execution Support in Compiler Explorer

## C++ Weekly - Ep 173 - The Important Parts of C++98 in 13 Minutes

## C++ Weekly - Ep 174 - C++20's _std::bind_front_

## C++ Weekly - Ep 175 - Spaceships in C++ operator 〈=〉

## C++ Weekly - Ep 176 - Important Parts of C++11 in 12 Minutes

## C++ Weekly - Ep 177 - _std::bind_front_ Implemented With Lambdas

## C++ Weekly - Ep 178 - The Important Parts of C++14 In 9 Minutes

## C++ Weekly - Ep 179 - Power C - A Native C Compiler for the Commodore 64

## C++ Weekly - Ep 180 - Whitespace Is Meaningless

## C++ Weekly - Ep 181 - Fixing Our bind_front with std::invoke

## C++ Weekly - Ep 182 - Overloading In C and C++

## C++ Weekly - Ep 183 - Start Using Raw String Literals

## C++ Weekly - Ep 184 - What Are Higher Order Functions?

## C++ Weekly - Ep 185 - Stop Using reinterpret_cast!

## C++ Weekly - Ep 186 - What Are Callables?

## C++ Weekly - Ep 187 - C++20's _constexpr_ Algorithms

## C++ Weekly - Ep 285 - Experiments With Generating Stably Random Game Assets

<details>
<summary>
Creating random objects that are independent from one another but still consistent and stable.
</summary>
[Experiments With Generating Stably Random Game Assets](https://youtu.be/xMdwK9p5qOw)

looking at randomly generated procedurals games contents (like **No Man's Sky**). taking a RNG (random number generator) called [Xoroshiro](https://en.wikipedia.org/wiki/Xoroshiro128%2B), Xoroshiro stands for the operators: XOR, rotate, shift,rotate.
the rng version is changed to be constexpr.the generator needs to be passed by reference.  
there is a problem that if we add an element inside a planet, all the other plants are changed. we want each planet to be random, but not having one change in a planet to change everything. the solution used in the video is to 'fork' the generator for each planet. the original generator creates a new generator with a new seed based on the next random value, and this makes nested generation of elements to still be random, but independent and consistent.  
in the comments someone says that the 'fork' part is called 'RNG splitting'.

_std::rotl_, _std::rotr_ - are shorthands for rotate left and which does some number checks and deals with positive and negative rotation shifts

</details>

## C++ Weekly - Ep 286 - How Command and Conquer's Dual Screen DOS Support Worked

<details>
<summary>
Old games using different memory address ranges to display debugging data to a second monitor.
</summary>

[How Command and Conquer's Dual Screen DOS Support Worked](https://youtu.be/wDvEzmEurlQ)
in the past there was a code review of Command and Conquer engine, there was a use of a mono-chrome screen as debugging. dual monitor support was added to windows 98, but there were hardware configurations that supported dual monitor even before that. so it seems that it was possible to use a mono-chrome adapter as one screen, and have an additional monitor. it wasn't plug-and-play.
memory ports of data didn't overlap between the devices, and those memory addresses were fixed. we could write data into the EGA and VGA ports,and different data into the monochrome ports. some games used this feature to display debugging data. we can do this with dos-box and setting secondary display. there is an example of using this with the game _mag jongg - VBA_ which displayed additional data.
this is a memory layout thing, the displays have different address range.
jason does a simple example with assembly code that tries to write to monochrome display.

</details>

## C++ Weekly - Ep 287 - Understanding _'auto'_

<details>
<summary>
Clearing up misunderstandings about the 'auto' keyword.
</summary>

[Understanding _'auto'_](https://youtu.be/tn69TCMdYbQ)

the old meaning of auto was 'explicit local variable', this was until C99 and c++98.

questions about _auto_ are

> - is there a hidden copy?
> - why isn't it a reference?
> - how does _auto_ even work?

the simple answer is that auto uses the exact same rules as template type parameters.

```cpp
template <typename T>
void func1(T param);

void func2(auto param);
```

this is the same, a copy by value, auto will never deduce a reference type for us. if we return a reference,we still get a copy, unless we declared our variable to be an reference itself.

```cpp
std::string getValue();
std::string &getReference();
void func3(const auto & param);

auto a = getValue(); //copy by value
auto b = getReference(); // another copy by value
auto & c = getReference(); //actually a reference
```

we can use _decltype(auto)_ to deduce a reference. **but shouldn't**.

in regards to pointers, it can deduce pointers, not reference. const-ness can be deduced.

```cpp
int *p = nullptr;
int i{};
auto p_copy = p; //p_copy is int*
auto i_copy = i; // i_copy is i*
//auto * i_copy2=i; //error. can't deduce auto* from i
//p_copy = i; //error = can't assign int to int*
p_copy = &i; //this is cool.
const int x{}; //const int
const int * px = &x;
//p_copy = &x; //error, conversion between const int* to int*
auto px_copy =  px; //px_copy is const int *
auto & x_ref =x; //x_ref is a reference to const int;
auto x_copy = x;  // not const, copy by value
x_copy = 55;  // fine
decltype(auto) x_decl = x; // the exact type of x, so const int
//x_decl = 9; error! x_decl is const int
```

> - _'auto'_ uses the same rules as template type parameters.
> - _'auto'_ will never deduce a reference.
> - 'const'-ness will be deduced only for reference and pointer types. when copying by value it won't maintain const.

</details>

## C++ Weekly - Ep 288 - Quick Perf Tip: Prefer _'auto'_

<details>
<summary>
Using 'auto' is good for performance by protecting us from accidental type conversions.
</summary>

[Quick Perf Tip: Prefer _'auto'_](https://youtu.be/PJ-byW33-Hs)

continuing the last video for understanding the _'auto'_ keyword.
the compiler always had to do the work of 'auto', it had to know what's returning from functions in order to do implicit casts and return type errors.

_'auto'_ will never do implicit cats. it will deduce const-ness (for reference and pointers), but will never perform a conversion.
if we don't have the _'-Wconversion'_ compiler turned on, we won't know that we might be doing something silly.
it really can come into play when iterating over a map or set. the key are const, if we forget to specify that we might cause a constructor call each iteration to construct the pair, just because we forgot to write const somewhere.
_std::pair_ has an implicit conversion to change const-ness for the members.

```cpp
#include <map>
const std::string &get_ref();
const std::string &get_ref2();
const char* get_str(){return "Hello World! or some other really long string to avoid optimizing";};
const std::map<std::string,std::string> & get_date();
int main()
{
    //int i = get_ref() + get_ref2();// can't do this, can't convert string + string to int
    auto s; = get_ref() + get_ref2(); //what the difference for the compiler if it's auto or std::string?
    std::string s1 = get_str(); //legal, compiles, but causes an implicit cast, a std::string was constructed here.
    auto s2 = get_str(); //no type conversion.
    std::string someString;
    const int stringLen = someString.size(); //another conversion! from std::size_t (unsigned 64bit) to int (signed 32bit)

    for (const std::pair<std::string, std::string> & elem : get_data())
    {
        //do something/
        // OOPS! we create a temporary pair! 100 lines of instructions just to create the copy of the pair!
        // binds to temporary rvalue
    }

    for (const std::pair<const std::string, std::string> & elem : get_data())
    {
        //do something/
        // now we are ok! we actually use a reference, back to no extra work
    }

    for (const auto & elem : get_data())
    {
        //do something/
        // now we are ok! we can't have an accidental type conversion. we might get an accidental copy if we don't specify a reference.
    }
    for (const auto & [key,value] : get_data())
    {
        //do something/
        // c++17 structured binding! very easy with auto!
    }
}
```

> - _'auto'_ uses the same rules as template type parameters.
> - _'auto'_ will never deduce a reference.
> - 'const'-ness will be deduced only for reference and pointer types. when copying by value it won't maintain const.
> - _'auto'_ will never never perform a conversion.

</details>

## C++ Weekly - Ep 289 - Returning From The 'void'

<details>
<summary>
We can return the result of a void() function.
</summary>

[Returning From The 'void'](https://youtu.be/26PtAmYk12M)

we can have function that don't return anything (void return type). and we can nest those function calls and return their results. this is legal.

```cpp
void do_something();
void do_something_else();

void func(int i)
{
    switch(i)
    {
    case 0: return do_something();
    case 1: return do_something_else();
    }
}
```

we can do auto type deduction. and decltype(auto) for "perfect returning"

```cpp
void fooVoid();
int fooInt();
int & fooIntReference1();
const int & fooIntReference2();

auto call(auto f) //maybe should be decltype(auto)
{
    return f();
}

call(fooVoid);
call(fooInt);
call(fooIntReference1); //what is the return type?
call(fooIntReference2); //what is the return type?
```

</details>

## C++ Weekly - Ep 290 - C++14's Digit Separators and Binary Literals

<details>
<summary>
Binary literals, digit separators
</summary>

[C++14's Digit Separators and Binary Literals](https://youtu.be/Yop9D3V2KBk)

binary literals to denote binary numbers and digit separators to make code more readable.

```cpp
auto b = 0b11101;
//auto b = 0b11121;
auto o = 0x1A7'D4;
auto b2=0b1110'1110;
auto d=100'000;
```

</details>

## C++ Weekly - Ep 291 - Start Using 'as_const'

<details>
<summary>
Calling a const member function from a non const object.
</summary>

[Start Using 'as_const'](https://youtu.be/w996YXhkpkE), [std::as_const](https://en.cppreference.com/w/cpp/utility/as_const)

objects can have const and non const member functions. if we want to use the const version function on a non-const variable

1. use _static_cast_ (which can invoke implicit conversions)
2. use _const_cast_ - can remove const, but won't invoke conversions
3. take a const reference
4. write as 'as_const' function (before c++17)
5. use _std::as_const_ (from the utilities header in c++17 and above)

```cpp
struct S
{
    S() =default;
    explicit S(int){};
    std::string_view get_value()
    {
        return "non-const";
    }
    std::string_view get_value() const
    {
        return "const";
    }
};

const S& as_const(const S &s) //take by const reference,always allowed
{
    return s;
}

//templated version
template <typename T>
const T& as_const(const T& t)
{
    return t;
}

int main()
{
    const S sc;
    fmt::print("A string: '{}'\n",sc.get_value()); //"const"

    S snc;
    fmt::print("A string: '{}'\n",snc.get_value()); //"non-const"
    fmt::print("A string: '{}'\n",(static_cast<const s>(snc)).get_value()); //"const" - static_cast
    fmt::print("A string: '{}'\n",(const_cast<const s>(snc)).get_value()); //"const" - const_cast

    //bad
    fmt::print("A string: '{}'\n",(const_cast<const s>(1)).get_value()); //"const" oops, we called on the constructor

    const S &sRef {snc}; //reference to const
    fmt::print("A string: '{}'\n",sRef.get_value()); //"const"

    fmt::print("A string: '{}'\n",as_const(snc).get_value()); //"const" -  using an 'as_const'
}
```

</details>

## C++ Weekly - Ep 292 - Safely Using 'goto' In C++

<details>
<summary>
What the compiler stops us from doing with 'goto'.
</summary>

[Safely Using 'goto' In C++](https://youtu.be/ELCc7JYW49k)

it's a common saying that we shouldn't use 'goto' in our programmes.

examples of some usages of goto that fail.

- can't jump and skip initialization of a variable.
- can't jump between functions (labels are scoped).
- can't have blank labels.
</details>

## C++ Weekly - Ep 293 - RPG in C++20 Project: Major Updates!

<details>
<summary>
Update about the game progress.
</summary>

[RPG in C++20 Project: Major Updates!](https://youtu.be/IWEdaC9evBc)

status update for the rpg game with new feature.\
running on a commador emulator.

optimization passes, we see the differences between different levels of optimization, in execution and size. ascii vs pscii(?spelling?) \
AVR architecture processor? \
looking at object assembly code.

</details>

## C++ Weekly - Ep 294 - Hello Commander X16

<details>
<summary>
Video controller addressing.
</summary>

[Hello Commander X16](https://youtu.be/JVoBZA2u2eM)

commander 16 is a modern retro computer design. getting the RPG project running on the new system. getting the video correctly, difference between x560 in how the addressing works

</details>

## C++ Weekly - Ep 295 - API Design: Principle of Least Surprise

<details>
<summary>
getting the api names and behaviors right.
</summary>

[API Design: Principle of Least Surprise](https://youtu.be/Qs4nje3KaFw)

continuing from _xoroshiro_ episode 285, changing the name from _'fork'_ to _'split'_ because thats the name of that functionality in the domain of generating random numbers.\
Also looking at the copy constructor and whether it should do the same. that would require a copy constructor which is mutable, so that would be two different methods, which behave differently depending on the const qualifier. and that's completely confusing which we should never do.

finally, thinking about the 'jump' functionality for random number generators (skip a large number of iterations), which in a naive way would take years to finish.

</details>

## C++ Weekly - Ep 296 - Constraining `auto` in C++20

<details>
<summary>
constraining the deduction of *auto* by using concepts.
</summary>

[Constraining `auto` in C++20](https://youtu.be/A8nNjpaiP5M)


```cpp
std::integral auto some_function()
{
    return 1;
    //return 2.5; // this will fail!
}
int main()
{
    const auto p1 = some_function();
    const std::integral auto p2 = some_function(); 
    
    //const std::floating_point auto p3 = some_function(); // this will fail
}
```
some confusion about pointers: the auto can hide the ptr type, we should be clear with the names and the types.
```cpp

int * get_object_ptr();
int get_object();
int main()
{
const auto obj_ptr1 = get_object_ptr(); //obj_ptr_1 is a pointer
const auto * obj_ptr2 = get_object(); //obj_ptr_2 wants to be a pointer, but receives a value, so we will get an error
}
```
we can define our own concept that "is not a pointer" so we can require the variable to fit our expectations. this all happens in compile time.
```cpp
template <typename T>
concept not_pointer = !std::is_pointer_v<T>;

int * get_some_pointer();
int main()
{
    const auto val1 = get_some_pointer(); //legal code, but the pointer is hidden!
    const not_pointer auto val2 = get_some_pointer(); // error!
}
```
</details>

## C++ Weekly - Ep 297 - C++ Homework: `auto` All The Things

<details>
<summary>
an home excerisice to make everything `auto` and see how it goes.
</summary>

not a best practice episode, rather an exercise for us to follow. we should find a project and make everything auto. the variables, the return type of function, and the parameters.

the example given is a ray tracer [smallpt](https://www.kevinbeason.com/smallpt/), which is 99 lines of code. we should try and make everything auto, use concepts if it makes sense. and try to replace everything with auto. see where it works, where it doesn't, where it makes sense and why.

[MyCode](ep297_smallpt.cpp)
</details>