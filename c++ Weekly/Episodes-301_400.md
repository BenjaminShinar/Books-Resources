<!--
// cSpell:ignore fsanitize
 -->

## C++ Weekly - Ep 301 - C++ Homework: _constexpr_ All The Things

<details>
<summary>
an home excerisice to make everything `constexpr` and see how it goes
</summary>

[C++ Homework: `constexpr` All The Things](https://youtu.be/cpdjQiRxEJ8)

another c++ homework assignment,after "auto everything" and "const everything". continuing with the smallpt file. now we try making everything _constexpr_. this includes member functions.

if we use compiler explorer, we can will see how the binary changes and more stuff becomes pre-calculated. it's theoretically possible to make everything at compile time, but it will require work (hint: the sqrt function). then only writing the file is at runtime.

</details>
 
## C++ Weekly - Ep 302 - It's Not Complicated, It's *std::complex*

 <details>
 <summary>
 a numeric type for complex number, with all the operators.
 </summary>
 
 [It's Not Complicated, It's std::complex](https://youtu.be/s_1SymtU0BI)

inside the "complex" header of the standard library. been here since foreaver, but still being worked on. the equality operator was removed and replace with the spaceship operator.\
there's also a user defined literals, constexpr support for getting the parts and for operators.

```cpp
std::complex<double> z =1.0 +2i;
```

a side note: some math functions still don't have constexpr support, as those depend on the cmath header. this will probably change in future standards of C++.

 </details>
  
## C++ Weekly - Ep 303 - C++ Homework: Lambda All The Things

 <details>
 <summary>
 an home excerisice to make everything with lambdas.
 </summary>
 
 [C++ Homework: Lambda All The Things](https://youtu.be/_xvAmEbK1vE)

continuing the homework series. now we want to make everything a lambda expression.\
Lambdas are by default const (unless stated to be mutable), and are implicitly constexpr. we use the same code sample. we need to decide what should and what shouldn't be a lambda expression. probably not the member functions. lambdas allow us to make code const. we can make free functions lambda.

is it possible to go too far with lambdas?

note: don't forget to have warnings on, use -std=c++20, and clear up the formatting.

</details>

## C++ Weekly - Ep 304 - C++23's `if consteval`

 <details>
 <summary>
 different ways and forms of checking a condition during compile time.
 </summary>
 
 [C++23's `if consteval`](https://youtu.be/AtdlMB_n2pI)

- C++17: `if constexpr`
- C++23: `is_constant_evaluated`
- C++23: `if consteval`

`if constexpr`, or `constexpr if`, was added in c++17, it's an conditional expressionthat must be evaluated in compile time, it must be part of a template.

for example, this will fail because of the two different return types:

```cpp
template<typename Param>
auto do_work(Param p)
{
    if (std::is_integral_v<Param>)
    {
        return 42+p;
    }
    else
    {
        return 4.2+p;
    }
}
```

but once we add the `if constexpr` to it, then it will be known at compile time and it will behave properly.

```cpp
template<typename Param>
constexpr auto do_work(Param p)
{
    if constexpr (std::is_integral_v<Param>)
    {
        return 42+p;
    }
    else
    {
        return 4.2+p;
    }
}
```

later, in c++20, we got `is_constant_evaluated`, this is different. this allows us to behave differently depending on whether the function was called in compile time or not.

only `if constexpr` allows to change types.

```cpp
constexpr int do_work_is_constant_evaluted()
{
    if (std::is_constant_evaluted())
    {
        //use compile time stuff,
        return 42;
    }
    else
    {
        return 43;
    }
}
```

example

```cpp
int main()
{
    [[maybe_unused]] constexpr auto a =do_work_is_constant_evaluated(); //42
    [[maybe_unused]] const auto b =do_work_is_constant_evaluated(); //42
    [[maybe_unused]] auto c =do_work_is_constant_evaluated(); //43 - not evaluated at compile time

    return a;
}
```

we cannot combine the two, it's allways true

```cpp
if constexpr(std::is_constant_evaluated()) //allways true
```

In c++23, we will get `if consteval`. note that **we don't have parentheses after the `if consteval`**. we can also negate the value. of.
it's behaves the same as `std::is_contant_evaluated`, but clearer. there are still some uses for earlier version.

```cpp
constexpr int do_work_23()
{
    if consteval
    {
        return 22;
    }
    else
    {
        return 11;
    }
}
```

</details>
  
## C++ Weekly - Ep 305 - Stop Using `using namespace`

<details>
<summary>
the case against using namespace directives.
</summary>

[Stop Using `using namespace`](https://youtu.be/MZqjl9HEPZ8E)

we all know we shouldn't write `using namespace std;` in header files, but what about inside implementation files?

```cpp
#include <iostream>

int main()
{
    std::string name="Benjamin";
    std::cout <<"hello " << name << '\n';
}
```

the ISO website says that we shouldn't have **using directives**, the most we can do is have a **using declration**, which is taking only the things we really care about.

```cpp
//using namespace std; //bad
using namespace std::cout; //ok
```

here is some bad code, which we don't have warnings for. we have different function overloads that we are unaware of, and changes to the namespaces can determine which version is being called.

```cpp
#include<fmt/format.h>

namespace emptycrate
{
    double calculate(double value)
    {
        return 4.23 * value;
    {
};

namespace company2
{
    int calculate(int value)
    {
        return 4*value;
    }
};

using namespace emptycrate; //using directive
using namespace company2; //using directive


int main()
{
    fmt::print("{}", emptycrate::calculate(2));
}
```

it's ok to use namespace directives inside function, but in that case, we should use namespace declerations. we can also pull in string literals suffixes, or chrono literals.

```cp
using namespace std::literals;

auto mystring= "Hello World"sv; //string view
```

</details>
 
## C++ Weekly - Ep 306 - What Are Local Functions, and Do They Exist in C++?

<details>
<summary>
C++ alternatives to local functions
</summary>

[What Are Local Functions, and Do They Exist in C++?](https://youtu.be/-EDx6fC6mkQ)

local functions aren't normally possible in c++. we can declare functions inside other function, but we can't define it.

//this doesn't work

```cpp
int main()
{
    int get_value() //can't do this
    {
        return 42;
    };
}
```

we can create a local class with a function in it

```cpp
int main()
{
    struct myStruct
    {
        static int get_value()
        {
            return 42;
        }
    }

    auto x= myStrcut::get_value();
// using myStruct::get_value; //not allowed!
// auto y= get_value(); //not allowed

}
```

we might try to name the class the name of the function and then define the `()` operator. then we get weird syntax like this:

```cpp
int main()
{
    struct get_value
    {
        int operator()()
        {
            return 55;
        }
    }
    auto x = get_value()();
}
```

but since c++11, we have lambdas. which is a struct behind the scenes, but one that the compiler creates for us.

</details>

## C++ Weekly - Ep 307 - Making C++ Fun and Accessible

<details>
<summary>
showing off the book of lifetime puzzles.
</summary>

[Making C++ Fun and Accessible](https://youtu.be/3RskKe7I6T4)

he made a nice book with puzzles about lifetime.

</details>

## C++ Weekly - Ep 308 - 'if consteval' - There's More To This Story

<details>
<summary>
Some extra information about C++23 `if consteval`
</summary>

['if consteval' - There's More To This Story](https://youtu.be/y3r9l3LZiJ8)

some parts that were left out.

- `if constexpr (conditional) {}`
- `if (std::is_constant_evaluated()){}`
- `if consteval {}`

```cpp
#include <type_traits>

constexpr int func()
{
   if (std::is_constant_evaluated())
   {
       return 42;
   }
   else
   {
       return 24;
   }
}
int main()
{
    auto value1 =func(); //24
    constexpr auto value2 = func(); 42
    return value1;
}
```

lets add a function

```cpp
//immediate function, must be called in compile-time
consteval int get_eval_value(int i)
{
return 42+i;
}
```

even though it seems like this is a compile time thing, it's not. \
 this doesn't work

```cpp
if (std::is_constant_evaluated())
{
    return get_eval_value(5);
}
```

but this does:

```cpp
if consteval
{
    return get_eval_value(5);
}
```

this has something to do with the difference between something that is truly compile time construct and things which are simply optimized away by the compiler.

 </details>

## C++ Weekly - Ep 309 - Are Code Comments a Code Smell?

<details>
<summary>
Discussion episode.
</summary>

[Are Code Comments a Code Smell?](https://youtu.be/8V6Ry5eTTcc)

the defintion of code smells.\
are all comments simply signs that we didn't try hard enough to make the code clear?

</details>

## C++ Weekly - Ep 310 - Your Small Integer Operations Are Broken!

<details>
<summary>
types that are promoted to integers are prone to weird conversion errors.
</summary>


[Your Small Integer Operations Are Broken!](https://youtu.be/R6_PFqOSa_c)

this code return zero, not -1. but also not some weird overflow max uint8_t thing. why?

```cpp
#include <cstdint>

int main()
{
    std::uint8_t value1 = 0;
    std::uint8_t value2 = 1;

    std::uint8_t result = value1-value2; //255 underflow
    auto result2 = value1-value2; //-1
    return value1 - value2; //why zero and not 
}
```

the result of substracting the two uint8_t variables is an int.


```cpp
#include <cstdint>
#include <typeinfo>
#include <type_traits>
#include <iostream>
int main()
{
    std::uint8_t value1 = 0;
    std::uint8_t value2 = 1;

    std::uint8_t result = value1-value2; //255 underflow
    auto result2 = value1-value2; //-1
    std::cout<<typeid(result).name() <<'\n'; //h for uint8_t
    std::cout<<typeid(result2).name() <<'\n'; //i for int
    return 0;
}
```
shifting is also a huge mess, arithmetic shift right does sign extentsion.
```cpp
std::uint8_t result1 = (value1-value2) >>1 ; //still 255
std::uint8_t result2 = (value1-value2) >>3 ; //still 255
std::uint8_t result3 = static_cast<std::uint8_t>(value1-value2)) >>3 ; // now its 31
```

shifting logic.
```
//signed
// 11000000 >> 1
// 11100000

//unsigned
// 11000000 >> 1
// 01100000
```
at other cases we might need casting over casting. we might decide to create a non_promoting type.

</details>

## C++ Weekly - Ep 311 - `++i` vs `i++`

<details>
<summary>
Prefix Increment and Postfix Increment
</summary>

[`++i` vs `i++`](https://youtu.be/ObVRSNvGitE)

the difference in semantics between the two versions.

```cpp
int main()
{
    int i = 0;

    //return ++i; // return 1
    //return i++; // return 0

    return 0;
}
```

if we want to define them for our own struct, we need to differentiate between the two versions, one with a dummy value. the postfix increment needs to return a copy.
```cpp
struct my_int
{
    //prefix increment
    constexpr my_int& operator++()
    {
        ++value;
        return *this
    }
    //postfix increment
    constexpr my_int operator(int)
    {
        const auto previous = *this; //make a copy
        value++; //doesn't matter if post or pre
        return previous;
    }

    int value;
};

int main(){
    my_int v{2};
    ++v; // prefix
    v++; // postfix
    return v.value;
}
```
the postfix version creates a copy, which is usually not what we wanted to do. if we have a complex object, this can cost us in performance.

```cpp
void sum_values(std::map<int,int>::const_iterator begin,std::map<int,int>::const_iterator end)
{
    int result = 0;
    while (begin != end)
    {
        result+=begin->second;
        ++begin;
        //++begin
    }
}
```

if we remove optimization, we can see the difference in the assembly code output. the difference is small because iterators are genrally cheap to create. 
</details>

## C++ Weekly - Ep 312 - Stop Using `constexpr` (And Use This Instead!)
<details>
<summary>
Using `static constexpr` variables and not `constexpr`.
</summary>

[Stop Using `constexpr` (And Use This Instead!)](https://youtu.be/4pKtPWcl1Go)

`constexpr` isn't what we (probably) think.

```cpp
// constrexpr -probably doesn't do what you think it does

constexpr int get_value (int value)
{
    return value *2;
}

int main()
{
    int value = get_value(6); // when is this calculated? complie time or run time?
    return value;
}
```

is the value usable in constant expression? we check this with a static_assert. this has to do with **core constant expressions**.

```cpp
int value = get_value(6);
static_assert(value == 12); // fails
const int value2 = get_value(6);
static_assert(value2 == 12); // passess
constexpr int value3 = get_value(6);
static_assert(value3 == 12); // passess
```

even with const, the value can still be calculated at compile time or at runtime, it's up to the the compiler.

but even if we declare it `constexpr`, it still isn't necessary calculated at compile time, as long as we don't use it in a compile time expression, then it's up to the compiler.

```cpp
constexpr std::array<int, 1000> get_values()
{
    std::array<int,1000> retval{};
    int count = 0;
    for (auto & val: retval){
        val = count*3;
        ++count;
    }
    return ret_val;
}

int main()
{
    constexpr auto values = get_values();
    return values[879];
}
```

this calculation can happen at compile time or at runtime, if we play with the optimization, things can behave differently.

**`constexpr` values are stack values**

in this example, what is going to be returned?

```cpp
int main()
{
    const int p* = nullptr;
    {
        constexpr auto values = get_values();
        p = &values[985];
    }
    return *p;
}
```

in clang O3 the value is what we expect (985*3), in gcc, we get an error for using an uninitialized value, if we add address sanitizer flag `--fsanitize=address` we see a warning about "stack-use-after-scope".


1. must run all test with address sanitizer enabed
2. must run both release and debug builds with address sanitizer

we actually only rarely want constexpr variables, we should use `static constexpr` instead. we want to force a static storage and initialization of those variables.

```cpp
int main()
{
    const int p* = nullptr;
    {
        static constexpr auto values = get_values();
        p = &values[985];
    }
    return *p;
}
```

this is part of the object life time puzzlers book!

the storage duration types are:
- static
- thread
- automatic
- dynamic

</details>
