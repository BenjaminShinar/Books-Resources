<!--
// cSpell:ignore fsanitize Fertig FTXUI NOLINT ssupported lstdc libuv Codecov fanalyzer pypy cppyy consteval emptycrate chrono constinit cppcheck INTERPROCEDURAL functools libbacktrace nodiscard valgrind csdint remove_cvref_t nlohmann catchorg Pico subspan Bloaty McBloatface rodata dynstr fullsymbols conan gitlab jenkins unreachables conway gtest codecov typesafe ftxui Hylo
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

## C++ Weekly - Ep 301 - C++ Homework: _constexpr_ All The Things

<details>
<summary>
an home exercise to make everything `constexpr` and see how it goes
</summary>

[C++ Homework: `constexpr` All The Things](https://youtu.be/cpdjQiRxEJ8)

another c++ homework assignment,after "auto everything" and "const everything". continuing with the smallpt file. now we try making everything _constexpr_. this includes member functions.

if we use compiler explorer, we can will see how the binary changes and more stuff becomes pre-calculated. it's theoretically possible to make everything at compile time, but it will require work (hint: the sqrt function). then only writing the file is at runtime.

</details>

## C++ Weekly - Ep 302 - It's Not Complicated, It's _std::complex_

 <details>
 <summary>
 a numeric type for complex number, with all the operators.
 </summary>

[It's Not Complicated, It's std::complex](https://youtu.be/s_1SymtU0BI)

inside the "complex" header of the standard library. been here since forever, but still being worked on. the equality operator was removed and replace with the spaceship operator.\
there's also a user defined literals, constexpr support for getting the parts and for operators.

```cpp
std::complex<double> z =1.0 +2i;
```

a side note: some math functions still don't have constexpr support, as those depend on the cmath header. this will probably change in future standards of C++.

 </details>
  
## C++ Weekly - Ep 303 - C++ Homework: Lambda All The Things

 <details>
 <summary>
 an home exercise to make everything with lambdas.
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

`if constexpr`, or `constexpr if`, was added in c++17, it's an conditional expression that must be evaluated in compile time, it must be part of a template.

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
constexpr int do_work_is_constant_evaluated()
{
    if (std::is_constant_evaluated())
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

we cannot combine the two, it's always true

```cpp
if constexpr(std::is_constant_evaluated()) //always true
```

In c++23, we will get `if consteval`. note that **we don't have parentheses after the `if consteval`**. we can also negate the value. of.
it's behaves the same as `std::is_constant_evaluated`, but clearer. there are still some uses for earlier version.

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

the ISO website says that we shouldn't have **using directives**, the most we can do is have a **using declaration**, which is taking only the things we really care about.

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

<!-- it's ok to use namespace directives inside function, but in that case, we should use namespace decelerations. we can also pull in string literals suffixes, or chrono literals. -->

```cp
using namespace std::literals;

auto my_string= "Hello World"sv; //string view
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

the result of subtracting the two uint8_t variables is an int.

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

shifting is also a huge mess, arithmetic shift right does sign extention.

```cpp
std::uint8_t result1 = (value1-value2) >>1 ; //still 255
std::uint8_t result2 = (value1-value2) >>3 ; //still 255
std::uint8_t result3 = static_cast<std::uint8_t>(value1-value2)) >>3 ; // now its 31
```

shifting logic.

```none
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

if we remove optimization, we can see the difference in the assembly code output. the difference is small because iterators are generally cheap to create.

</details>

## C++ Weekly - Ep 312 - Stop Using `constexpr` (And Use This Instead!)

<details>
<summary>
Using `static constexpr` variables and not `constexpr`.
</summary>

[Stop Using `constexpr` (And Use This Instead!)](https://youtu.be/4pKtPWcl1Go)

`constexpr` isn't what we (probably) think.

```cpp
// constexpr -probably doesn't do what you think it does

constexpr int get_value (int value)
{
    return value *2;
}

int main()
{
    int value = get_value(6); // when is this calculated? compile time or run time?
    return value;
}
```

is the value usable in constant expression? we check this with a static_assert. this has to do with **core constant expressions**.

```cpp
int value = get_value(6);
static_assert(value == 12); // fails
const int value2 = get_value(6);
static_assert(value2 == 12); // passes
constexpr int value3 = get_value(6);
static_assert(value3 == 12); // passes
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

in clang O3 the value is what we expect (985\*3), in gcc, we get an error for using an uninitialized value, if we add address sanitizer flag `--fsanitize=address` we see a warning about "stack-use-after-scope".

1. must run all test with address sanitizer enabled
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

## C++ Weekly - Ep 313 - The `constexpr` Problem That Took Me 5 Years To Fix!

<details>
<summary>
Getting compile time values to be usable in runtime.
</summary>

[The `constexpr` Problem That Took Me 5 Years To Fix!](https://youtu.be/ABg4_EV5L3w)

> (Compile-time views Into optimally sized compile-time data. I'ts awesome, no really, trust me!)

taking a standard string from compile time to runtime.

```cpp
#include <string>
#include <fmt>
constexpr std::string make_string(std::string_view base, const int repeat)
{
    std::string retval;
    for (int count =0; count<repeat; ++count)
    {
        retval += base;
    }
    return retval
}

int main()
{
    std::string result = make_string("Hello Jason, ",3);
    fmt::print("{}",result); //this works
    constexpr std::string result2 = make_string("Hello Jason, ",3);
    fmt::print("{}",result2); //this fails
}
```

we can't let the constexpr string escape into a non-constexpr context.

however, this does work, we get length at compile time and print it at runtime.

```cpp
constexpr auto get_length(std::string_view base, const int repeat)
{
    return make_string(base,repeat).size();
}
int main()
{
    constexpr static auto length = get_length("Hello world,",4);
    fmt::print("{}", length); //this works
}
```

we can get the size, but not the string itself.

(he does something with std::array, but it needs to call the make_string function twice)

```cpp
template <std::size_t Len>
constexpr auto get_array(const std::string& str)
{
    std::array<char,Len> result;
    std::copy(str.begin(),str.end(),result.begin());
    return result;
}

int main()
{
    constexpr static auto length = get_length(make_string("hello Jason, ",3));
    constexpr static auto str = get_array<length>(make_string("hello Jason, ",3));
    constexpr static auto sv = std::string_view(str.begin(), str.end());

}
```

it's impossible to do this

```cpp
constexpr std::string value; //doesn't compile!
```

it can be a bit nicer if we delegate the creation of the string to a lambda, but it's still the same issue.

lets try this, use some buffer data. it works, but the size of the binary increases!

```cpp
struct oversized_array
{
    std::array<char, 10*1024*1024> data{};
    std::size_t size;
};
constexpr auto to_oversized_array(const std::string & str)
{
    oversize_array result;
    std::copy(str.begin(),str.end(),result.begin());
    result.size=str.size();
    return result;
}

int main()
{
    constexpr auto make_data =[](){
        return make_string("hello Jason, ",3); // lambda
    };
    constexpr static auto str = to_oversized_array(make_data()));
    constexpr static auto sv = std::string_view(str.begin(), str.end());
    fmt::print("{}: {}", sv.size(), sv);
}
```

lets try to get the correct size: but this doesn't work. an input variable can't be an constant expression value.

```cpp
constexpr auto to_right_size_array(const std::string & str)
{
    constexpr static auto oversized = to_oversized_array(str));
    std::array<char, oversized.size> result;
    std::copy(oversized.begin(),oversized.end(),result.begin());
    return result;
}
```

but we can pass a lambda that creates a constant expression value. the function should actually be **consteval**, because we would never want to all it in runtime.

```cpp
template<typename Callable>
consteval auto to_right_size_array(Callable callable)
{
    constexpr static auto oversized = to_oversized_array(callable());
    std::array<char, oversized.size> result;
    std::copy(oversized.data.begin(),std::next(oversized.data.begin(),oversize.size),result.begin());
    return result;
}
int main()
{
    constexpr auto make_data =[](){
        return make_string("hello Jason, ",3); // lambda
    };
    constexpr static auto str = to_right_size_array(make_data));
    constexpr static auto sv = std::string_view(str.begin(), str.end());
    fmt::print("{}: {}", sv.size(), sv);
}
```

this still isn't good enough, we still create two oject, a _std::array_ and the _std::string_view_. there also a problem with having static variables in the _consteval_ function.

so now we try other crazy stuff, we have a function that returns a reference to the template argument. and now we got something that the compiler can optimize.

> Class non template type parameter

```cpp
template<auto Data>
consteval auto & make_static()
{
    //take a template parameter and return a reference to it.
    return Data;
}

consteval auto to_string_view(auto callable) -> std::string_view
{
    constexpr auto &static_data = make_static<to_right_sized_array(callable)>();
    return std::string_view{static_data.begin(), static_data.end()};
}

int main()
{
    constexpr static auto sv = to_string_view(make_data);
    fmt::print("{}: {}",sv.size(),sv);
}
```

- lambda that returns a string
- we create an oversized array (which should be big enough for any reason) which is constant time value
- then we use the oversized array as template argument to create a smaller array.
- which we use as static reference
- and then we use it to create the string_view.

</details>

## C++ Weekly - Ep 314 - Every Possible Way To Force The Compiler To Do Work At Compile-Time in C++

<details>
<summary>
Different ways to do compile-time calculations.
</summary>

[Every Possible Way To Force The Compiler To Do Work At Compile-Time in C++](https://youtu.be/UdwdJWQ5o78)

just making a value or function `constexpr` doesn't force the compiler to run it a compile time.

we can make the value _static_, which forces the compiler to compute the value at compile time, but also requires it to be const.

we can use `constinit`, but it also has to be static.

```cpp
constexpr int get_value(int value)
{
    return value *3;
}

int main()
{
    constexpr auto value1 = get_value(2); //up to the compiler to decide.
    constexpr static auto value2 = get_value(3); //will be calculated at runtime.
    constinit static auto value3 = get_value(4); //also must be static, but not const.
}
```

if we change the function to be `consteval`, then it must be done it compile time sense, but that's not always what we want.

```cpp
consteval int get_value_consteval(int value)
{
    return value *3;
}
```

in the previous episodes, we had some other tricks, like using a template parameter

```cpp
template<auto Value>
consteval const auto make_compile_time()
{
    return Value;
}
int main()
{
    auto value 5 = make_compile_time<get_value(7)>();//comp
}
```

there's also a [blog post](https://andreasfertig.info/) by Andreas Fertig, which wraps the normal function with a `consteval auto as_constant` function to force compile time calculations.

```cpp
consteval auto as_constant(auto value)
{
    return value;
}
int main()
{
    auto value7 = as_constant(get_value(15));
}
```

and we want to generalize it to moveable stuff as well

```cpp
template <typename ... Param>
consteval decltype(auto) consteval_invoke(Param && ... param)
{
    return std::invoke(std::forward<Param>(param)...);
}

int main()
{
    auto value8 = consteval_invoke(get_value, 9);
}
```

| type/keyword                  | compile-time calculation | const | static | example                                           | notes                                                                        |
| ----------------------------- | ------------------------ | ----- | ------ | ------------------------------------------------- | ---------------------------------------------------------------------------- |
| `constexpr`                   | up to the compiler       | yes   | no     | `constexpr auto value = get_value(1);`            |
| `constexpr static`            | yes                      | yes   | yes    | `constexpr static auto value = get_value(1);`     | must be static const                                                         |
| `constinit static`            | yes                      | no    | yes    | `constinit static auto value = get_value(1);`     | must be static                                                               |
| `consteval` function          | yes                      | no    | no     | `auto value = get_value_consteval(5)`             | argument must be compile time constants, function can't be used in run time. |
| template parameter            | yes                      | no    | no     | `auto value = make_compile_time<get_value(10)>()` | using templates                                                              |
| wrapping `consteval` function | yes                      | no    | no     | `auto value = as_constant(get_value(10))`         | inner function can be reused                                                 |
| `consteval invoke` wrapper    | yes                      | no    | no     | with moveable and callable                        |

</details>

## C++ Weekly - Ep 315 - `constexpr` vs `static constexpr`

<details>
<summary>
more comparison between constexpr and static constexpr. looking at benchmark numbers.
</summary>

[`constexpr` vs `static constexpr`](https://youtu.be/IDQ0ng8RIqs)

clarify: static at global scope isn't the same as static in the function scope. static at global scope is duplicated into each translation unit.

```cpp
//some header

static constexpr auto bigData = generate_bigData(); //duplicated
inline constexpr auto bigData2 = generate_bigData(); //probably what i meant
```

at the function level scope we use `static constexpr`, and we usually mean this scope in the previous videos.

benchmark examples, the version with the local constexpr array (dynamic initialization) is faster than the on with the static constexpr array. this is counter to what we said earlier.

```cpp
// in the current stack
std::uint32_t to_ascii_base36_digit_dynamic(std::uint32 digit)
{
    constexpr std::array<char, 32> base36_map = {'0','1',/*...*/, 'x','y','z'};
    return base36_map[digit];
}
// in the global storage
std::uint32_t to_ascii_base36_digit_static(std::uint32 digit)
{
    static constexpr std::array<char, 32> base36_map = {'0','1',/*...*/, 'x','y','z'};
    return base36_map[digit];
}
```

he plays with the numbers (data size) in the benchmark, and increases the map size to 72, then 144 and 2048. now the results are reversed, the static constexpr version is much faster. it's just a matter of copying data onto the stack vs accessing the global data. it also changes with the optimization level and the compiler (clang vs gcc vs visual studio).

</details>

## C++ Weekly - Ep 316 - What Are `const` Member Functions?

<details>
<summary>
The basics on `const` and non-`const` member functions.
</summary>

[What Are `const` Member Functions?](https://youtu.be/bqd9ILyQRxQ)

`const` and none `const` member functions.

the only difference between `struct` and `class` is the default access level.

we can use `const` member functions on non-`const` objects, just like we can make a `const` reference to a non-const variable, but not the opposite.

```cpp
#include <fmt/format.h>

struct string
{
    std::size_t size(){ return m_size;}
    std::size_t const_size const (){ return m_size;}
    private:
    std::size_t m_size{};
};

int main()
{
    const string my_const_str;

    // fmt::print("string size: {}",my_const_str.size()); //fails
    fmt::print("string size: {}",my_const_str.const_size()); //ok

    string my_str;
    [[maybe_unused]] const &str_ref_const = my_str; // no problem
    [[maybe_unused]] &my_const_str = my_str; // error!
}
```

continuing our string example, now supposedly we look at the _iterator_. we again need a const and non const version, and this is important if we want **for loops**.

luckily, const and non const functions acts as overloads, so we have both version and the correct one is chose as needed.

```cpp
#include <fmt/format.h>

class string
{
    public:
    std::size_t const size(){ return m_size;}
    char * begin();
    char * end();
    const char * begin() const;
    const char * end() const;
    private:
    std::size_t m_size{};
};

int main()
{
    const string my_str;

    for (const auto character : my_str)
    {
        fmt::print("character: {}\n",character);
    }
}
```

</details>

## C++ Best Practices Game Jam Info, Rules and Quick-Start

<details>
<summary>
C++ game jam. starts April 1,2022.
</summary>
[Info, Rules and Quick-Start](https://youtu.be/4V4ZrB3o6g4)

- must use FTXUI
- must follow c++ best practices
- must start from the provided template and compile all the actions.
- run with no errors or address sanitizer warning.
- try not to disable warnings.

<kbd>Use this template</kbd>, then <kbd>Create Repository from template</kbd>.

we need a build enviornment,visual studio and some other stuff.

```sh
sudo apt install python3-pip g++ clang-tidy clang-format git cppcheck
pip3 install cmake ninja conan
# add folders to path
```

c/c++ extension pack (from microsoft)

configure to run with debug. launch target "intro" to compile the ftxui dependencies

there are two demo

```sh
./intro turn_based
./intro loop_based
```

we document disabling warning with `NOLINT`, for debugging we need a debug configuration.

</details>

## C++ Weekly - Ep 317 - How Member Functions Work

<details>
<summary>
continuing the previous video, this time understanding what the compiler does.
</summary>

[How Member Functions Work](https://youtu.be/4etjb2_KAaE)

first of all, a function overload is happening at compile time, unlike virtual functions, which happen at runtime.

if we play with compiler explorer and the optimizations, we can see that the compiler passes the _this_ pointer as the first argument to the function, and if the member function is `const`, then the pointer is const.\
this parameter is sometimes passed in the registers.

there is actually even another thing which is passed, the return type, according to the caller conventions.

</details>

## C++ Weekly - Ep 318 - My Meetup Got Nerd Sniped! A C++ Curry Function

<details>
<summary>
Creating a `curry` function, which can either execute a function or return a new one.
</summary>

[My Meetup Got Nerd Sniped! A C++ Curry Function](https://youtu.be/15U4qutsPGk)

someone said that it's hard to create a currying function in c++.

requirements:

- be like `bind`
- take the first N parameters
- either return a function or execute it

```cpp
int add (int x, int y, int z){return x+y+z;}

int main()
{
    auto new_func = curry(add,1,2);
    auto result = new_func(3);
    return result; //should return 6
}
```

so here is a solution that was suggested.

1. something that takes a callable and variadic parameters.
2. create a lambda which captures the callable and the parameters, and can take in another set of variadic parameters. when the lambda is called, it executes the callable with both sets of parameters.
3. we check (during compile time, `if constexpr`) with the `requires` clause if we can immediately execute the function with the current set of parameters,if it's possible, then call the function without returning the lambda.

```cpp
template<typename Callable, typename ... Param>
auto bind(Callable callable, Param ... param)
{
    auto bound = [callable, param ...](auto ... inner_param)
    {
        return callable(param..., inner_param...);
    }

    if constexpr (requires {callable(param...);})
    {
        return callable(param...);
    }
    else
    {
        return bound;
    }
}

int main()
{
    auto bound1 = bind(add,1,2);
    auto result1 = bound1(3);
    auto result2 = bind(add,1,4,5); // int value of 10, not a function.
    return result1; //should return 6
}
```

however, this will fail for trying to cascade the calls.

```cpp
int main()
{
    const auto bound = bind(add,1)(2)(3);
    return bound;
}
```

so we move to the next version, which works recursively and instantiates additional templates. it doesn't package them one inside another.

```cpp
template<typename Callable, typename ... Param>
auto curry(Callable f, Param ... ps)
{
    return [f,ps...](auto...qs)
    {
        if constexpr (requires {f(ps...,qs...);})
        {
            return f(ps...,qs...);
        }
        else
        {
            return curry(f,ps...,qs...);
        }
    };
}

int main()
{
    const auto curried = curry(add,1)(2)(3);
    return curried;
}
```

the parameters are copied each time, which might be a problem, and more than that, the function doesn't work for the basic case.

```cpp
int main()
{
    const auto curried = curry(add,1,2,3);
    return curried;
}
```

so the updated form is similar, but with a `if constexpr` check at the start.

```cpp
template<typename Callable, typename ... Param>
auto curry(Callable f, Param ... ps)
{
    if constexpr( requires {f(ps...)})
    {
        return f(ps...);
    }
    else
    {
        return [f,ps...](auto...qs)
        {
            if constexpr (requires {f(ps...,qs...);})
            {
                return f(ps...,qs...);
            }
            else
            {
                return curry(f,ps...,qs...);
            }
        };
    }
}
```

not quite there, we have duplicated code of checking if the current form is callable. so we move to the next form. which works for all cases so far.

```cpp
template<typename Callable, typename ... Param>
auto curry(Callable f, Param ... ps)
{
    if constexpr( requires {f(ps...)})
    {
        return f(ps...);
    }
    else
    {
        return [f,ps...](auto...qs)
        {
            return curry(f,ps...,qs...);
        };
    }
}
```

it even works for weird cases, like passing it no parameters.

```cpp
int main()
{
    const auto curried = curry(add,1,2)()()(()(3);
    return curried;
}
```

the problem is the copying, we don't handle forwarding. if we take references, we run into object lifetime issues. there might be a way to parametrize it (take copy of rvalue, reference of lvalue), but it would probably quickly become a monstrous code.

</details>

## C++ Weekly - Ep 319 - A JSON To C++ Converter

<details>
<summary>
A zero runtime library that allows using Json resources at compile time.
</summary>

[A JSON To C++ Converter](https://youtu.be/HROQPE59q_w)

introducing json2cpp compiler library, the goal is to have no runtime overhead, to use a statically compiled json resource that can be used in compile time. compatible with what it needs to be, and with an adaptor.

everything is statically known at compile time, it creates a cpp class that is directly mapped to the properties of the json file. it's a custom data structure that we can include as part of the compile process. this can be used as a configuration file that makes compile time decisions. he suggests that it's great for embedded devices.

</details>

## C++ Weekly - Ep 320 - Using `inline namespace` To Save Your ABI

<details>
<summary>
Another attempt to mend the ABI problem.
</summary>

[Using `inline namespace` To Save Your ABI](https://youtu.be/rUESOjhvLw0).

avoid problems with ABI (application binary interface) breaking.

imagine that we start with this code:

```cpp
namespace lefticus{
    struct Data{
        char c;
        int i;
        char c2'

    };
    int calculate_things(const Data& data);
}

int main()
{
    const lefticus::Data some_data{};
    return lefticus::calculate_things(some_data);
}
```

but now we want to change the order of the arguments in the struct. but this is breaking ABI. the layout changed, the size changed. We want to be able to safely change the ABI, in c++11 there was a new feature called **inline namespaces**

```cpp
namespace lefticus{
inline namespace v2_0_0 {
    struct Data{
        int i;
        char c;
        char c2'

    };
    int calculate_things(const Data& data);
}
}


int main()
{
    const lefticus::Data some_data{};
    return lefticus::calculate_things(some_data);
}
```

now we have two definitions, so we either get a compile time error if we try to use them, or a linkage error. this protects us from undefined behavior.

the downside is that we need to manually change the namespace. the inline namespace means that the name never shows up in the code. we can have multiple ABIs maintained at the same time.

```cpp
namespace lefticus{
namespace v1_0_0 { //explicit namespace
    struct Data{
        char c;
        int i;
        char c2'

    };
    int calculate_things(const Data& data);
}

inline namespace v2_0_0 { //implicit inline namespace
    struct Data{
        int i;
        char c;
        char c2'

    };
    int calculate_things(const Data& data);
}
}


int main()
{
    const lefticus::Data some_data{}; // uses implicit namespace
    const lefticus::v2_0_0::Data some_old_data{}; // uses explict namespace
    auto x = lefticus::calculate_things(some_data); // overload resolution
    auto old_x = lefticus::calculate_things(some_data); // overload resolution
}
```

</details>

## C++ Weekly - Ep 321 - April 2022 Best Practices Game Jam Results

<details>
<summary>
Game Jam conclusion
</summary>

[April 2022 Best Practices Game Jam Results](https://youtu.be/TQTb6ewowtk).

the topic of the gameJam was "round", some problems were encountered, etc...

(going over some games - not much of an episode)

</details>

## C++ Weekly - Ep 322 - Top 4 Places To Never Use `const`

<details>
<summary>
Cases where declaring const is not the preferred behavior.
</summary>

[Top 4 Places To Never Use `const`](https://youtu.be/dGCxMmGvocE)

a list episode!

> On a non-reference return type

```cpp
std::string make_value();
const std::string make_value_const(); //bad

int main()
{

    std::string s;
    s= make_value_const();
}
```

this behavior stops us from performing move operations, as we can't move from const, so we must perform a copy/assignment operator, which is a performance issue.

> Don't `const` local values that need to take advantage of implicit moe-on-return operations

```cpp
inline const S make_value_3()
{
    const S s;
    return s;
}

inline const std::optional<S> make_value_4()
{
    const S s;
    return s;
}

int main()
{
    S s = make_value_3(); // again, no move
    auto s2 = make_value_4(); // use move
}
```

we have a technically true but actually pointless warning about a move constructor.

_std::optional_ has an implicit conversion, because it's a value type, rather than a pointer type.

```cpp
inline std::optional<S> make_value_5()
{
    //const S s;
    return std::optional<S>{std:in_place_t{}};
}
```

> if you have multiple different objects that might be returned, then you are also relying on implicit move-on-return (aka automatic move).

in the following case we have two constructors and a copy, because both options are initialized, if we would move the objects into the inner scopes, we could create just one and get move operations and return value optimization.

```cpp
inline S make_value_multiple(bool option)
{
    S s1;
    S s2;
    if (option)
    {
        return s1;
    }
    else
    {
        return s2;
    }

}

int main(int argc, const char*[])
{
    auto s = make_value_multiple(argc==1); // can't optimize return value

}
```

> don't `const' non-trivial value parameters that you might need to return directly from the function.

```cpp

inline S make_value_from_arg_const(const S s)
{
    return s; //because we return it, const is bad in function defintion
}

inline S make_value_from_arg_move(S s)
{
    return s;
}

int main([[maybe_unused]] int argc, const char*[])
{
    auto s1 = make_value_from_arg_const(s{}); // no move
    auto s2 = make_value_from_arg_move(s{}); // move
}
```

> Don't `const` any **member** data!\
> It breaks implicit and explicit moves\
> It breaks common use cases

```cpp
struct Data
{
    const S s;
}

int main()
{
    Data d;
    //d = Data{}; // doesn't work, default assignment operator for D can't assign const;
    Data d2 = std::move(d); // also a copy, not a move. can't move from const

}
```

this behavior is seen when we use data containers, this prevents us from efficiently resizing containers.

```cpp
struct StringData
{
    const std::string s;
}

int main()
{
    std::vector<StringData> data;
    data.emplace_back();
    data.emplace_back(); //resizing requires copying because we don't have move operations
}
```

if we have an invariant data member which we can't change without breaking other stuff, then we should simply write an accessor/mutator.

</details>

## C++ Weekly - Ep 323 - C++23's `auto{}` and `auto()`

<details>
<summary>
explicitly copy a value.

</summary>

[C++23's `auto{}` and `auto()`](https://youtu.be/5zVQ50LEnuQ)

a C++23 feature, we can use `auto{}` to make a copy of something.

```cpp
int main()
{
    int i =4;
    return auto{4}; //explicitly make a copy
}
```

this comes into use in templates and when we use _auto_ type parameters.

```cpp
void use (const auto &);
void function(const auto &something)
{

    //auto copy = something; //can't be done
    use(std::decay_t<decltype(something)>{something});
}
```

this is the motivating example. we want to erase all the elements which are like the first one.\
but the output of the code also removes all additional instances of the second unique element.
this is because we use swapping internally in (`std::erase_if`).

```cpp
void erase_all_of_first(auto & container)
{
    //c++20 std::erase standard form
    std::erase(container, container.front());
}

int main()
{
    std::vector<std::string> values {"test3","test3","hello there world","bod", "test","hello there world"};

    erase_all_of_first(values);
    for (const auto &str : values)
    {
        std::cout<< str<< '\n';
    }
    // "hello there world","bob", "test"
}
```

to fix this, we take a copy.

```cpp
void erase_all_of_first(auto & container)
{
    //c++20 std::erase standard form
    std::erase(container, {container.front()});
}
```

note: the same functionality can be achieved with a short function

```cpp
auto copy (const auto & value)
{
    return value;
}
```

</details>

## C++ Weekly - Ep 324 - C++20's Feature Test Macros

<details>
<summary>
Macros that allow us to check if we can use a feature in our current standard library implementation.
</summary>

[C++20's Feature Test Macros](https://youtu.be/4Bf8TmbibXw)

c++20 standardized compile time behavior, it allows us to check at compile time if a feature exists in the standard. the value of the macro is the year and the month the feature was accepted, so if something was added in c++20, the value might be "201707L" - designating that it was accepted early on to the standard, back in July 2017.

having this macros allow us to check if the library which we are using supports a feature

```cpp
#if __cpp_lib_constexpr_string >= 201907L
constexpr std::string make_string()
{
    std::string result;
    result = "Hello ";
    result += "World";
    result += " Test Long String";
    return result;
}

TEST_CASE("to_string_view produces a std::string_view from std::string")
{
    constexpr static auto result = lefticus::tools::to_string_view([](){return make_string();});
    static_assert(std::is_sam
    e_v<decltype(result), const std::string_view>);
    STATIC_REQUIRE(result == "Hello World Test Long String");
}

#endif
```

this allows us to check if we can use a specific version of implementation, in cases that the feature had changes over time, or if we are using a truncated compiler version and we want to make sure a feature from the next standard is supported.

</details>

## C++ Weekly - Ep 325 - Why vector of bool is Weird

<details>
<summary>
The special case of the vector of booleans.
</summary>

[Why vector of bool is Weird](https://youtu.be/OP9IDIeicZE)

Vector of bool isn't a straight forward as a vector of other types. even thought a boolean can be represented in a single bit, it usually is stored as a single byte.

```cpp
int main()
{
    bool data[15]{};
    return sizeof(data); //15
}
```

vector of bool has a space optimized version

```cpp
#include <vector>
#include <iostream>
int main()
{
    std::vector<bool> data;
    data.push_back(true);
    std::cout << data.size() << '\n'; // 1
    std::cout << data.capacity() << '\n'; // 64
}
```

there is a problem, what if we want a reference to one of the elements? this doesn't make sense, we cant reference bits in the memory, only bytes.

```cpp
#include <vector>
#include <iostream>
int main()
{
    std::vector<bool> data;
    data.push_back(true);
    auto &front = data.front();

}
```

we might try to create a bit field instead, but this is also not possible.

```cpp
struct Data{

    bool b1:1;
    bool b1:1;
    bool b3:1;
};

void getRef()
{
    Data d{};
    auto &b = d.b3; // error!
}
```

in the case of _std::vector\<bool\>_, **a proxy object** is used instead, this object knows how to interact with the correct bit.

```cpp
#include <vector>
#include <iostream>
int main()
{
    std::vector<bool> data;
    data.push_back(true);
    auto front = data.front(); // proxy
    front = false; // modify data through proxy
    std::cout << data.front() << '\n'; // false
}
```

this does get annoying, as we can't bind directly.

```cpp
#include <vector>
int main()
{
    std::vector<bool> data;
    data.push_back(true);
    data.push_back(false);

    for (auto & bit : data) //can't be done.
    {
        bit = 0;
    }
}
```

the form of `for (const auto & bit : data)` works, but it can't modify the data. and the form `for (auto bit : data)` can modify the data, but doesn't look right. in any other case we wouldn't expect to work.\
one way to avoid this is to use forwarding references. `auto &&`, which works for both proxies and regular behavior, so if we see it, we should know that it's a proxy object and be careful

```cpp
#include <vector>
int main()
{
    std::vector<bool> data;
    data.push_back(true);
    data.push_back(false);

    for (auto && bit : data) //forwarding
    {
        bit = 0;
    }
}
```

</details>

## C++ Weekly - Ep 326 - C++23's Deducing `this`

<details>
<summary>
Matrix use case
</summary>

[C++23's Deducing `this`](https://youtu.be/5EGw4_NKZlY)

one use case is when we have const and non-const member functions, like the `at` function.

```cpp
#include <array>
#include <cstddef>

template <typename Contained,std::size_t Width, std::size_t Height>
struct Matrix{
    std::array<Contained,Width * Height> data;

    Contained &d at(std::size_t X, std::size_t Y)
    {
        return data(Y * Width + X);
    }
    const Contained &d at(std::size_t X, std::size_t Y) const
    {
        return data(Y * Width + X);
    }

};

int main()
{
    Matrix<int, 5,5> data{};
    data.at(2,3) = 15;
    return data.at(2,3);
}
```

this is code duplication, can we get around this? we can have a shared function that makes use of the static deduction.

```cpp

template <typename Contained,std::size_t Width, std::size_t Height>
struct Matrix{
    std::array<Contained,Width * Height> data;

    Contained &d at(std::size_t X, std::size_t Y)
    {
        return at (*this.X,Y):
    }
    const Contained &d at(std::size_t X, std::size_t Y) const
    {
        return at (*this.X,Y):
    }

    private:
    static template <typename This>
    auto& at(This &obj,std::size_t X, std::size_t Y)
    {
        return obj.data.at(Y* Width + X);
    }

};
```

and a simplified form will end up like this

```cpp
#include <array>
#include <cstddef>

template <typename Contained,std::size_t Width, std::size_t Height>
struct Matrix{

    template<typename Self>
    auto &at(this Self &&self, std::size_t X, std::size_t Y)
    {
        //data[0]; // can't access
        return std::forward<Self>(self).data.at(Y * Width + X);
    }

};

int main()
{
    const Matrix<int, 5,5> data{};
    // data.at(2,3) = 15;
    return data.at(2,3);
}
```

however, inside this modified function, we can't access member never, even if it's not explicitly a static function. and it's also weird to get it. we cant do some work to recover it in a different way.

```cpp
int main()
{
    const Matrix<int, 5,5> data{};
    //return Matrix<int,5,5>::at(data,2,3); // doesn't work;

    //  form 2
    using const_at = const int&(*)(const Matrix<int,5,5>&, std::size_t X, std::size_t Y);
    auto func = Matrix<int,5,5>::at<const Matrix<int,5,5>&>;
    return func(data,2,3);
}
```

</details>

## C++ Weekly - Ep 327 - C++23's Multidimensional Subscript Operator Support

<details>
<summary>
multiple parameters in subscript operators.
</summary>

[C++23's Multidimensional Subscript Operator Support](https://youtu.be/g4aNGgLzVqw)

in the previous video, we used tha the `at` operator, which should throw an exception if we try to access an element out of range. so we should be clear about the issue.

```cpp
#include <array>
#include <cstddef>

template <typename Contained,std::size_t Width, std::size_t Height>
struct Matrix
{

    template<typename Self>
    auto &at(this Self && self, std::size_t X, std::size_t Y)
    {
        if(X>= Width) throw std::range_error("X out of range");
        if(Y>= Height) throw std::range_error("Y out of range");
        return std::forward<Self>(self).data[Y * Width + X];
        //return std::forward<Self>(self).data.at(Y * Width + X); //not relevant anymore
    }

};

int main()
{
    const Matrix<int, 5,5> data{};
    return data.at(9,3);
}
```

C++23 introduced multi-dimensional subscript operators, meaning we can use more than one index inside the brackets. we can't combine both forms until the compilers support both options, but for a simple example, we can use a const version of this.

```cpp
#include <array>
#include <cstddef>

template <typename Contained,std::size_t Width, std::size_t Height>
struct Matrix
{

    const auto & operator[](std::size_t X, std::size_t Y)
    {
        return data [Y * Width + X];
    }

};

int main()
{
    const Matrix<int, 5,5> data{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16};
    return data[2,3]; // 13
}
```

</details>

## C++ Weekly - Ep 328 - Recursive Lambdas in C++23

<details>
<summary>

</summary>

[Recursive Lambdas in C++23](https://youtu.be/hwD06FNXndI)

</details>

## C++ Weekly - Ep 329 - How LTO Easily Makes Your Program Faster

<details>
<summary>

</summary>

[How LTO Easily Makes Your Program Faster](https://youtu.be/9nzT1AFprYM)

</details>

## C++ Weekly - Ep 330 - Faster Builds with `extern template` (And How It Relates to LTO)

<details>
<summary>
extern templates are a way to instantiate templates in one file rather than recreate it each time.
</summary>

[Faster Builds with `extern template` (And How It Relates to LTO)](https://youtu.be/pyiKhRmvMF4)

like the earlier video, we have an 'add' function in a different complication unit. this time we make it a template.

since c++ there was a feature called `extern template`, which stops the compiler from creating the same template again and again. we need to declare the template type as `extern`, and then have one place file that instantiates it explicitly.

```cpp
#ifndef DECELERATIONS
#define DECELERATIONS

template<typename Type>
Type add(Type lhs, Type rhs)
{
    return lhs+rhs;
}

extern template int add<int>(int, int);

#endif
```

and in a separate cpp file

```cpp
#include "decelerations.hpp"

template int add<int>(int,int);
```

however, this prevents us from the optimizing, so we need the LTO again.

```cmake
project(test CXX)
cmake_minimum_required(VERSION 3.18)
include(CheckIPOSsupported)

add_executable(tst file1.cpp file2.cpp impl.cpp)
set_property(TARGET test PROPERTY INTERPROCEDURAL_OPTIMIZATION)
```

> "we get to have our cpp template cake and eat it too."
>
> extern template saves us on build time in each cpp file of our large project. (assuming the function is expensive to compile).

if the function is expensive, LTO probably wouldn't be able to inline it anyways.

</details>

### C++ Weekly - Ep 331 - This Game Teaches C++!

<details>
<summary>
Making C++ Fun and accessible.
</summary>

[This Game Teaches C++!](https://youtu.be/snQhhWE1xR4), [best practices github](https://github.com/cpp-best-practices).

a shell script to install all sorts of stuff to run on a new machine. we clone the game repo. let it run conan and cmake, compile whatever it needs. we can start the game.

We need to modify the code, to make the lesson start rather than the game. the 'game' tells us to change the source code, and we launch the game again. each time we get another lesson we are told what to look at in the code.

the game can run on computers, and even on raspberry pie (even if vscode struggles a bit)

</details>

## C++ Weekly - Ep 332 - C++ Lambda vs `std::function` vs Function Pointer

<details>
<summary>
The differences between lambda, function pointers and `std::function`.
</summary>

[C++ Lambda vs `std::function` vs Function Pointer](https://youtu.be/aC-aAiS5Wuc)

the three are similar but not the same. a lambda is not a `std::function`.

```cpp
#include <functional>

int add(int x, int y)
{
    return x+y;
}

int call (int (f*)(int,int))
{
    return f(2,3);
}
int call (const std::function<int(int,int)> &f)
{
    return f(10,11);
}

int main ()
{
    auto add_lambda = [](const int x,const int y){return x+yl};
    std::function<int(const int, const int)> func{add};

    call(add);

}
```

a lambda is an anonymous struct. std::function is a function holder, a type-erased wrapper around a callable. it's an abstraction

note: a capture-less lambda is implicitly convertible to function pointer. but a std::function can be constructed from either a lambda (with or without a capture), a function pointer or anything else with the same format.

</details>

## C++ Weekly - Ep 333 - A Simplified `std::function` Implementation

<details>
<summary>
trying to make an implementation of std::function in c++20.
</summary>

[A Simplified `std::function` Implementation](https://youtu.be/xJSKk_q25oQ)

remember, _std::function_ is not a lambda or function pointer.

we start with a forward template deceleration, and then we specialize on it. we want to make sure the signature is correct, is that the template arguments are the same. we want the compiler to throw an error if we pass something which isn't a valid function signature (and return type).

```cpp
template <typename T>
class function;

template <typename T, typename ... Param>
class function<Ret (Param....)>
{

};

int main()
{
    function<int> func1; // shouldn't compile
    function<int(int, int)> func2; // should compile
}
```

next, we want to overload the operator(), to make sure this is a callable. we need a constructor, and to somehow store the thing which we got, in a type erasure format. don't forget the rule of five.

```cpp
#include <memory>

template <typename T, typename ... Param>
class function<Ret (Param....)>
{
    public:
    function(Ret (*f)(Param...)) : callable{std::make_unique<callable_impl<Ret (*)(Param....)>>(f)}; //constructor

    Ret operator()(Param... param) {
        return callable->call(param...); // unpack parameters
    }
    private:
    struct callable_interface {
        virtual Ret(Param....) = 0; // pure virtual
        virtual ~callable_interface = default;
    };
    std:::unique_ptr<callable_interface> callable;

    template<typename Callable>
    struct callable_impl : callable_interface {
        callable_impl(Callable callable_): callable{std::move(callable_)}

        Callable callable;
        Ret call (Param... param)
        {
            return callable(param...)
        }

    };

};

int f(int x, int y){
    return x+y;
}

int main()
{
    function<int(int, int)> func{f}; // should compile
    auto x = func(1); // should fail
    auto y = func(1,2); // should work
}
```

we might need to add many more ctors, like one for function objects.

```cpp
template <typename T, typename ... Param>
class function<Ret (Param....)>
{
    public:
    function(Ret (*f)(Param...)) : callable{std::make_unique<callable_impl<Ret (*)(Param....)>>(f)}; //constructor for function pointers

    template<typename FunctionObject>
    function(FunctionObject fo) : callable{std::make_unique<callable_impl<FunctionObject>>(std::move(fo))}; //constructor for function object
};
```

we could also use `std::invoke` instead. we might need a _Clone_ method, also forwarding and unwrapping references.

</details>

## C++ Weekly - Ep 334 - How to Put a Lambda in a Container

<details>
<summary>
Three Ways to put lambdas in a container
</summary>

[How to Put a Lambda in a Container](https://youtu.be/qmd_yxSOsAE)

it's actually possible to put a lambda in a container, but it's not straight forward.

pushing only a single lambda, even if another lambda has a similar structure

```cpp
#include <vector>

int main()
{
    auto l = [](){return 42;};
    std::vector<decltype(l)> data;
    data.push_back(l);
    //data.push_back([](){return 43;}); fails
}
```

converting to a std::function, has some massive overhead.

```cpp
#include <vector>
#include <functional>

int main()
{
    auto l = [](auto j){return j+ 42;};
    std::vector<std::function<int(int)>> data;
    data.push_back(l);
    data.push_back([](auto k){return k;});

}
```

converting to a function pointer, doesn't allow a capture.

```cpp
#include <vector>

int main()
{
    auto l = [](auto j){return 42+j;};
    std::vector<int(*)(int)> data;
    data.push_back(l);
}
```

back to the first case, we can move the lambda creation to a function, so all the lambdas have the same type.

```cpp
#include <vector>

auto make_lambda(int value)
{
    return [value](int i){ return i+value;};
}


int main()
{
    std::vector<decltype(make_lambda(42))> data;
    data.push_back(make_lambda(2));
    data.push_back(make_lambda(5));
}
```

there is a danger of ODR (one defintion rule) violation, so it's probably better to create a callable directly.

</details>

## C++ Weekly - Ep 335 - Projects That Need Your Help! - Channel News and Updates

<details>
<summary>
Some updates and general information.
</summary>

[Projects That Need Your Help! - Channel News and Updates](https://youtu.be/Xu1O-44ikso)

Describing what Microsoft MVP is, what on-site training is. the [second channel](https://www.youtube.com/channel/UCADySP7Hy8TxgfDEe2GZQyw) was re-branded from "The Retro Programmer" to "The [Fill in the Blank] Programmer".
Going over the Repositories and the Best Practice book. There will be a workshop before CppCon 2022.

1. The CPP tutorial game was renamed into "Travels", requires ideas about which lessons to create, and help with trying to create those lessons.
2. Cpp weekly project - all future ideas, what topics we would like to see in he future?
3. he's still looking for sponsorships.

</details>

## C++ Weekly - Ep 336 - C++23's Awesome std::stacktrace Library

<details>
<summary>
A new library in c++23, able to show the stacktrace in runtime.
</summary>

[C++23's Awesome std::stacktrace Library](https://youtu.be/9IcxniCxKlQ)

in compiler explorer with `-std=c++2b --lstdc++_libbacktrace` and all debug symbols enabled

```cpp
#include <stacktrace>
#include <string>
#include <iostream>

int main()
{
    auto trace = std::stacktrace::current();
    std::cout << std::to_string(trace) <<'\n';
}
```

the output isn't perfect, some levels are are empty, the function name isn't shown properly, etc...

if we add a function call, it seems that there is no inlining (the library might prevent them in some way). we can put the stack trace print into a strcut

```cpp
struct S{
    S() {
        auto trace = std::stacktrace::current();
        std::cout << std::to_string(trace) <<'\n';
    }
};
```

we can't currently do a stacktrace from a compile time context. but if we make it part of the function default value, then it's called outside of the function initiation

```cpp
void func3 ([[maybe_unused]] S s =S{})
{

}

int main()
{
    func3();
}
```

next is playing with lambda.

if we look at the API, the class allocator aware (we can use polymorphic allocators), it has many of the characteristics of a container. it uses a class called 'stacktrace_entry' which is similar, but not exactly like 'source_location' class.

this class will probably change until c++23 is released.

</details>

## C++ Weekly - Ep 337 - C23 Features That Affect C++ Programmers

<details>
<summary>
C23 changes which can be useful to c++ programmers.
</summary>

[C23 Features That Affect C++ Programmers](https://youtu.be/jOFrKN54M5g)

a new version of **C**, which will have changes which can be useful for c++ programmers who need to combine code between c and c++.

- `#embed` - pull a text file, compile time (using `#include` doesn't work well).
- `constexpr` - values, but not functions.
- attributes - such as `[[nodiscard]]`, `[[maybe_unused]]`.
- unnamed parameters in function defintion.
- typed enumeration - `enum myEnum: int {ONE, TWO};`.
- `__has_include` - compile time check if a file can be included.

</details>

## C++ Weekly - Ep 338 - Analyzing and Improving Build Times

<details>
<summary>
Tools and tips to optimize build time.
</summary>

[Analyzing and Improving Build Times](https://youtu.be/Iybb9wnpF00)

Cmake options, `-ftime-trace` clang compile time trace flag, a json file is created. we go to "chrome://tracing" and load the file. we see a graph.

1. Expensive header includes to move to PCH (pre compiled header)
2. Template instantiations (reduce or eliminate)
3. Function templates to `extern` for our users

the example in the video is chai-script, so not all of those options are viable.

</details>

## C++ Weekly - Ep 339 - `static constexpr` vs `inline constexpr`

<details>
<summary>
inline static data at file (global) scope.
</summary>

[`static constexpr` vs `inline constexpr`](https://youtu.be/QVHwOOrSh3w)

follow up on episode 312, which argued against using constexpr in favor of `static constexpr`.

```cpp
void func()
{
    constexpr auto value = some_function();
    static constexpr auto value2 = some_function2(); // preferable
}
```

but there are some more considerations, such as static constexpr at a file header.

```cpp
//header.hpp
static constexpr int data[10000000];
//main.cpp
#include "header.hpp"
const int *get_data();

int main()
{
    return get_date()[100];
}

//file1.cpp
#include "header.hpp"
const int *get_data()
{
    return data;
}
```

if we look at the file sizes, the both files now have the large sizes. and if we link them together, the size doubles.

if we change the header file to `inline constexpr`.

```cpp
//header.hpp
inline constexpr int data[10000000];
```

each compiled file is still very large, but the linker merges them together, so the output file only contains one 'copy' of this data.

</details>

## C++ Weekly - Ep 340 - Finally! A Simple String Split in C++!

<details>
<summary>
A language standard split string capability
</summary>

[Finally! A Simple String Split in C++!](https://youtu.be/V14xGZAyVKI)

since c++20, there is finally a simple way to split strings

```cpp
#include <fmt/format.h>
#include <ranges>
#include <string_view>
int main()
{
    auto split_strings = std::string_view{"Hello world C++20!"} | std::ranges::views::split(' ');
    for (const auto &string: split_strings)
    {
        std::cout << std::string_view{string.begin(), string.end()}; // still needed
        fmt::print("{}\n",string);
    }
}
```

in c++23, the std::string_view will be able to take the range and construct a string out of it.

the ranges views are lazily evaluated, so there might be some issues with `constexpr` functions.

</details>

## C++ Weekly - Ep 341 - `std format` vs lib `{fmt}`

<details>
<summary>
comparing between the standard library format library and the {fmt} package.
</summary>

[std format vs lib {fmt}](https://youtu.be/zc6B-j0S9Iw)

`std <format>` is what we get with the compiler, while `lib {fmt}` is what it was based on.

1. `{fmt}` is available today
2. `<format>` is currently only available only on visual studio,

```cpp
#include <fmt/format.h> // lib format
#include <format> // std library

int main()
{
    std::cout << std::format("Hello {}",42);
    std::puts(std::format("Hello {}",42).c_str());
    fmt::print("Hello {}",42);
}
```

`std::format` doesn't have a way to print to the console, so we need to jump through some hoops. maybe in c++23. it is guaranteed to have ABI stability. `{fmt}` has more utility built in into it, with helpers and so on. it also has constexpr capability (with `FMT_COMPILE`)

```cpp
#include <fmt/format.h>
#include <fmt/compile.h>
#include <array>

constexpr auto get_string(int value)
{
    std::array<char, 1000> result{};
    fmt::format_to(result.begin(), FMT_COMPILE("Hello {}"),value);
    return result;
}

int main()
{
    static constexpr auto str = get_string(42);
    std::puts(str.date());
}
```

</details>

## C++ Weekly - Ep 342 - C++20's Ranges: A Quick Start

<details>
<summary>
Ranges help us solve common problems and avoid bugs.
</summary>

[C++20's Ranges: A Quick Start](https://youtu.be/sZy9XcGHmI4)

std::ranges are wrappers which help us write better code for our common algorithms.

> 1. handy adapters for common algorithms
> 1. pipe-able range views
> 1. simple solutions to annoying problems
> 1. lazy transform has interesting implications
>
> all of these are constexpr capable

we start with a buggy implementation, in this case the `get_data()` function returns a temporary object each time, so we should get a iterator mismatch runtime error.

```cpp
#include <algorithm>
#include <vector>
#include <ranges>

std::vector<int> get_data() { return std::vector<int>{1,2,3,4,5,6}; }

bool test_data()
{
    auto result = std::all_of(get_date().begin(),get_date().end(), [](const int i){return i<5;});
    return result;
}
int main()
{
    test_data();
}
```

we could fix it by being explicit

```cpp
bool test_data()
{
    auto data = get_data();
    auto result = std::all_of(data.begin(),data.end(), [](const int i){return i<5;});
    return result;
}
```

but ranges help us avoid this.

```cpp
bool test_data()
{
    auto result = std::ranges::all_of(get_date(), [](const int i){return i<5;});
    return result;
}
```

ranges also allow us pipes and views, and use lazy evaluation.

```cpp
//lazy evaluated
bool test_data()
{
    auto result = std::ranges::all_of(
        get_date() |std::ranges::views::drop(1) | std::ranges::views::take(2),
        [](const int i){return i<5;});
    return result;
}
```

we can also use it for simple data wrangling, like skipping the first element.

```cpp
#include <format>
void iterate_data()
{
    for (const auto &elem : get_data() | std::ranges::views::drop(1)
    {
        fmt::print("{}\n",elem);
    }
}
```

lazy evaluation also helps us with transformations. we use explicit template parameters for lambda (introduced in c++20).

```cpp
void iterate_with_index()
{
    auto make_index = [idx = std::size_t{0}]<typename T>(const T &elem) mutable{
        return std::pair<std::size_t, const T &>{idx++, elem};
    };
    for (const auto &[index, elem]: get_date() | std::ranges::views::transform(make_index))
    {
        fmt::print("{}: {}\n",index, elem);
    }
}
```

</details>

## C++ Weekly - Ep 343 - Digging Into Type Erasure

<details>
<summary>
Hiding away the concrete type.
</summary>

[Digging Into Type Erasure](https://youtu.be/iMzEUdacznQ)

type erasure allows us to "hide" types in runtime, inheritance is one such way to do so in c++. we could run `dynamic_cast<>` and check if it fits to one specific type, but that's all.

the `std::function` is type-erased, it can take a function, a function pointer, a lambda, or anything which is callable, even combining the `std::bind_front`.

```cpp
int use_function(const std::function<int(int,int)> &f)
{
    return f(2,3);
}
```

> Type Erasure: hide the exact type of an object wen you work with it.
>
> - simpler and faster to compile interfaces.
> - can work with any type that might be declared in the future.
> - compilation firewall to prevent recompiling the entire library for adding a new type.

we create an `animal_view` type, which has constructor that takes a reference, and it works with anything that has a `speak` function. defined to it. it works like an interface.

```cpp
class animal_view{
    public:
        template <typename Speakable>
        explicit animal_view (const Speakable *speakable) : object{&speakable}, speak_impl{
            [](const void *obj) {return static_cast<const Speakable *>(obj)->speak();}
            }
        {
        }

        void speak() const
        {
            speak_impl(object);
        }
    private:
        const void *object;
        void (*speak_impl)(const void *);
};

void do_animal_things(animal_view animal){animal.speak();}

int main()
{
    struct Cow {
        void speak() const {fmt::print("Moo\n");}
    };
    struct Sheep {
        void speak() const {fmt::print("Baa\n");}
    };

    do_animal_things(animal_view{Cow{}});
    do_animal_things(animal_view{Sheep{}});
}
```

</details>

## C++ Weekly - Ep 344 - `decltype(auto)`: An Overview of How, Why and Where

<details>
<summary>
"Perfect returning" for types.
</summary>

[`decltype(auto)`: An Overview of How, Why and Where](https://youtu.be/E5L66fkNlpE)

`decltype(auto)` deduces the **exact** type of an expresion.

in this example, what is the type of _i_? it changes depending on how we set the callsite variable.

```cpp
#include <type_traits>

const int &get_value();

int main()
{
    auto i = get_value();
    static_assert(std::is_same_v<decltype(i),int>);

    const auto i2 = get_value();
    static_assert(std::is_same_v<decltype(i2),const int>);

    const auto &i3 = get_value();
    static_assert(std::is_same_v<decltype(i3),const int &>);

    auto &i4 = get_value();
    static_assert(std::is_same_v<decltype(i4),const int &>); // must be const
}
```

auto will never deduce a reference on it's own. but `decltype(auto)` does. it comes into use in generic code,

```cpp
int main()
{
    decltype(auto) x = get_value():
    static_assert(std::is_same_v<decltype(x),const int &>);
}
```

there is also perfect returning, which seems silly, but is needed when working with templated generic code.

```cpp

const int &get_value();

auto get_value_wrapped_error()
{
    return get_value(); // this is now an int
}

decltype(auto) get_value_wrapped()
{
    return get_value(); // this maintains the exact type
}

int main()
{
    decltype(auto) y = get_value_wrapped_error():
    static_assert(std::is_same_v<decltype(y),const int &>); // this fails

    decltype(auto) z = get_value_wrapped():
    static_assert(std::is_same_v<decltype(z),const int &>);
}
```

we can get in trouble if we use parentheses in the return statement.

```cpp
auto get_value() -> decltype(auto){
    int i = 42;
    return i; //no problem
    // return (i); // error! this is now an expression, and therefore an rvalue - reference to local value
}
```

so it's suggested to not use parentheses when returning values from functions.

</details>

## C++ Weekly - Ep 345 - No Networking in C++20 or C++23! Now What?

<details>
<summary>
Some recommended libraries for networking.
</summary>

[No Networking in C++20 or C++23! Now What?](https://youtu.be/v6m70HyI0XE)

The networking features weren't part of the c++ standard library in either c++20 or c++23.

we finally have package managers for c++, and we can use them go get networking libraries!

- conan
- hunter
- vcpkg

There are c++ libraries for networking:

- ASIO - which is the closest thing to what we are expecting to get from the standard
- QT - a framework\utility library with a lot of stuff, including networking
- Poco - utility library - networking focused
- ACE - low level
- uvw - wrapper around C libuv library

there are libraries which wrap around libcurl, such as **cpr**. which is meant to be C++ version of python **Request** library.

for message passing, there is **ZeroMQ**, and there's **asio-grpc** for asynchronous interface, and there are web-sockets libraries, and REST libraries.

</details>

## C++ Weekly - Ep 346 - C++23's `bind_back`

<details>
<summary>
binding parameters to an object.
</summary>

[C++23's `bind_back`](https://youtu.be/pDiP2frdMnI)

a new c++23 feature, _std::bind_back_

```cpp

struct Point {
    int x;
    int y;

    void displace(int x_displacement, int y_displacement) {
        x += x_displacement;
        y += y_displacement;
    }

    Point operator+(Point displacement) const {
        return Point{x + displacement.x, y+ y.displacement};
    }

    void print() const {
        fmt::print("{{{},{}}}\n",x,y);
    }
};

int main()
{
    Point p{42,24};
    p.print();
    auto displace_by_1_2 = std::bind_back(&point:displace,1,2);
    displace_by_1_2(p1);
    p.print();
}
```

we can use it in algorithms or in a range, it can sometimes replace a lambda.

</details>

## C++ Weekly - Ep 347 - This PlayStation Jailbreak NEVER SHOULD HAVE HAPPENED

<details>
<summary>
Implicit type conversion created a bug.
</summary>

[This PlayStation Jailbreak NEVER SHOULD HAVE HAPPENED](https://youtu.be/rWCvk4KZuV4)

there was ps5 jail break, which was a C vulnerability, a heap buffer size over flow.

[CVE report](https://hackerone.com/reports/1340942)

the size is of type size_t, but its being used in a function that gets an integer. so there is some conversion going on and the allocated buffer is much smaller than expected, and then we get possible overflows.

```c++
std::size_t getSize();
void * doAlloc(int size);
int main()
{
    std::size_t size = getSize();
    void * memory = doAlloc(size); // this is now truncated into integer.
}
```

this might have been avoided if the correct warning flags: `-Werror -Wall -Wextra -Wconversion`.

</details>

## C++ Weekly - Ep 348 - A Modern C++ Quick Start Tutorial - 90 Topics in 20 Minutes

<details>
<summary>
c++ topics and examples of how they are used in lambdas.
</summary>

[A Modern C++ Quick Start Tutorial - 90 Topics in 20 Minutes](https://youtu.be/VpqwCDSfgz0), [notes](https://github.com/lefticus/cpp_weekly/issues/168)

topics in c++

- lambdas
- struct
- constexpr
- operator overloading
- call operator
- const member functions
- braced initialization
- `auto` return type deduction
- "compiler magic"function parameters
- pass-by-value
- attributes on parameters
- pass-by-reference
- pass-by-value vs pass-by-reference
- pre-increment and post increment
- trailing return types
- class vs struct
- private vs public
- implicit conversions
- function pointer
- static member function- using alias
- efficiency when changing functions
- templates
- template argument type deduction
- alias templates
- template instantiations
- `noexcept`
- `noexcept` in the type system
- variadic templates
- variadic lambdas
- fold expressions
- function attributes
- concepts
- non-type template parameter
- integer sequences
- template parameter pattern matching
- explicit lambda templates
- tuples
- unpacking of tuples
- variadic `sizeof...()` operator
- direct initialization of members
- `mutable` keyword
- non-const member function
- reference members
- member copies
- object layout
- member padding
- order of construction/destruction
- generalized lambda capture
- immediately invoked lambda
- return value optimization
- guaranteed return value optimization
- initializer_list
- recursive lambdas
- deducing `this`
- recursive functions
- trivially copyable types
- higher order functions
- dangling references
- undefined behavior
- inheritance
- mutiple inheritance
- function hiding
- variadic `using` deceleration
- scoping / lookup rules
- class template argument deduction
- deduction guides
- algorithms
- ranges
- \<functional> header
- virtual member function
- member function pointers
- special member functions
- member function call syntax
- type erasure
- dynamic vs automatic storage
- type_traits
- operator `<=>`
- protected
- virtual inheritance
- compilation model
- ODR violations
- preprocessor
- project structure and layout
- the breadth of the standard layout
- variable templates
- coroutines
- modules

</details>

## C++ Weekly - Ep 349 - C++23's `move_only_function`

<details>
<summary>
Passing a type erased function which owns non-copyable data.
</summary>

[C++23's `move_only_function`](https://youtu.be/OJtGOJI0JEw)

a new c++23 feature.

when we try to use std::function, it can't be used with a lambda that owns a non-copyable capture, such as a unique pointer, or any move-only element.

```cpp
#include <functional>

void register_callback(std::function<int (int)> callback);

int main()
{

    std::function<int (int)> cb{
        [i= std::make_unique<int(42)>](const int val) {return val + *i;}
    };
    //this fails
    register_callback(cb);
}
```

_std::move_only_function_ was created to over come this limitation.

```cpp
#include <functional>

void register_callback(std::move_only_function<int (int)> callback);

int main()
{
    //this works
    std::move_only_function<int (int)> cb{
        [i= std::make_unique<int(42)>](const int val) {return val + *i;}
    };

    // this works
    register_callback(std::move(cb));
}
```

we can't a move of const object, so if we want, we can create the lambda as a temporary object (rvalue) by writing the code directly in the function call.

</details>

## C++ Weekly - Ep 350 - The Right Way to Write C++ Code in 2022

<details>
<summary>
Best Practices for working with C++ code today.
</summary>

[The Right Way to Write C++ Code in 2022](https://youtu.be/q7Gv4J3FyYE)

> "We are in a golden age of C++ tooling"

many of these examples are up on the best practices repo.

Continuos Build Enviornment:\
github, gitlab, jenkins, there are many

use as many compilers as you can, compilers can have different warnings

- gcc
- clang
- MSVC (cl.exe)
- clang.cl

testing enviornment:

- docTest
- Catch2
- gtest
- boost

test coverage analysis and reporting, preferably integrated with the build enviornment. it should have historical tracking to catch regressions.

- coveralls
- codecov.io

static analysis, as many as possible.

- at least all the warnings treated as errors (`-Werror`, `-Wconversion`)
- `gcc -fanalyzer`
- `cl.exe /analyze`
- cppcheck
- clang-tidy
- sonar
- PVS studio

runtime analysis during testing

- Address sanitizer
- Undefined behavior sanitizer
- Memory sanitizer
- Thread sanitizer
- valgrind
- **DrMemory** (windows compatible)
- Debug Checked Iterators from the compiler

Fuzz Testing - random input to the API

- used together with runtime analysis
- trying to feed the API weird data

Ship with Hardening enabled

- control flow guard (visual studio)
- `_FORTIFY_SOURCE` (GCC)
- stack protector
- **UBSan** `-fsanitize-minimal-runtime`

</details>

## C++ Weekly - Ep 351 - Your 5 Step Plan For Deeper C++ Knowledge

<details>
<summary>
5 things to do in order to get better understanding of c++.
</summary>

[Your 5 Step Plan For Deeper C++ Knowledge](https://youtu.be/287_oG4CNMc)

object life time: a class that writes to the console when it's created or destroyed, and tells us what happened, this allows us to see return value optimization, copy elision, and so on.

study the lambda: when are things created, what happens if it's inside the body of the lambda, inside the capture list, what happens when we copy a lambda.

create a `std::function implementation`, get an understanding of type erasure, lifetime. make it work with lambda, free functions, static functions, etc... . also make it `constexpr`, and to try and impalement small function optimization.

</details>

## C++ Weekly - Ep 352 - Not Doing This Should Be Illegal! (Always Fuzz Your C++!)

<details>
<summary>
Fuzz testing
</summary>

[Not Doing This Should Be Illegal! (Always Fuzz Your C++!)](https://youtu.be/Is1MurHeZvg)

fuzz testing is a way of testing software for vulnerabilities, we use a tool to feed random input into the program, and seeing if it causes problems. we can use sanitizers to see if the input has effect. this isn't a new concept, the term "fuzzing" exists since the late 80's.

example of a bug, not checking for string ending, and how using fuzz testing finds the bug. another example with rotating bits in an integer. another example of how fuzzing would have found earlier bugs.

</details>

## C++ Weekly - Ep 353 - Implicit Conversions Are Evil

<details>
<summary>
Implicit conversions when we don't expect them
</summary>

[Implicit Conversions Are Evil](https://youtu.be/T97QJ0KBaBU)

> "This C++ Feature Must GO!"

writing safer C++. starting with some examples.

the key value of a map (dictionary) is always const. so if we specify the std::pair, the first element must be const, even if we didn't specify the map as such.

converting a std::string to string_view and converting back.

slicing when converting from derived class to base class.

implicit conversion from std::shared_ptr to `const std::shared_ptr`.

there are many sharp edges, and many places where we can fall into. it can create non-trivial objects and cost us in performance.

</details>

## C++ Weekly - Ep 354 - Can AI And ChatGPT Replace C++ Programmers?

<details>
<summary>
Chatting with GPT and asking him to create c++ code.
</summary>

[Can AI And ChatGPT Replace C++ Programmers?](https://youtu.be/TIDA6pvjEE0)

asking the chatbot to write c++ code, modernize it, update it to use different features, etc...\
we even ask it to write clang-tidy stuff.

</details>

## C++ Weekly - Ep 355 - 3 Steps For Safer C++

<details>
<summary>
Static Analysis, Sanitizers and Fuzzing
</summary>

[3 Steps For Safer C++](https://youtu.be/dSYFm65KcYo)

these steps depend on us having tests and ci-cd integration

- static analysis
  - warnings as errors
  - fix all existing warning
  - build both `release` and `debug`
  - add `clang-tidy`(without project-specific options)
  - increase warning level `-Weverything`, `-Wall`
- Satirizers
  - address sanitizers
  - undefined behavior
  - thread sanitizer
  - `valgrind`, `Dr memory` if we can't add sanitizers to the tests.
- Add fuzzing
  - finds all the thing we don't think about
  - first with all user-facing apis
  </details>

## C++ Weekly - Ep 356 - The Python Enabled Calculators of 2022

<details>
<summary>
The Future of Programming Education - Calculators with Python
</summary>

[The Python Enabled Calculators of 2022](https://youtu.be/82v0jYGh0p0)

python is available out of the box in many devices, while c++ isn't that easy to get started with. there are calculators with python built-in into them. the power of the devices varies (from 256kb to 8mb in the charts presented), they have special versions to python (python3 compatible).\
it's not great to type on them, but they have some common shortcuts to make it easier.

showing different calculators, going over how they behave and how they support python.

</details>

## C++ Weekly - Ep 357 - `typename` VS `class` In Templates

<details>
<summary>
Which should be used?
</summary>

[`typename` VS `class` In Templates](https://youtu.be/86Pa973BW4Y)

when defining a template, our template parameters could be defined as either "Typename" or "Class". in most cases it doesn't matter.

```cpp
template<typename Type1, class Type2>
void func(Type1 t1, Type2 t2)
{
    /*...*/
}
```

but there are some cases which require using one or the other. this has been true for nested template declarations, but since c++17 things have changed.

</details>

## C++ Weekly - Ep 358 - C23's `#embed` and C++23's `#warning`

<details>
<summary>
New pre-processor directives in upcoming C and C++ standards.
</summary>

[C23's `#embed` and C++23's `#warning`](https://youtu.be/ibKnNRAq5UY).

in C++23 and C23 standards, there are new pre-processor directive (commands/defintions/macros).

C23 now has `#embed` - which allows us to directly pull in data from other files into the binary. we can limit the numer of the bytes taken. theoretically, we could use "/dev/random" to get random data from the local machine.

```cpp
#include <cstdio>

static constexpr char data[] =
{
#embed "input.txt" limit(5)
};

int main()
{
    std::puts(data);
}
```

C++23 provides the `#warning` directive, which we can use to write warnings at compile time directly.

```cpp
#include <climits>
#include <csdint>
#if UINT_MAX == UINT16_MAX
#warning "Project not tested on 8bit platforms"
#endif
```

</details>

## C++ Weekly - Ep 359 - _std::array_'s Implementation Secret (That You Can't Use!)

<details>
<summary>
Simple Implementation of the standard array.
</summary>

[_std::array_'s Implementation Secret (That You Can't Use!)](https://youtu.be/uLbv2u536G0)

the array type is meant to allow a typesafe way to view arrays, unlike C-arrays, which devolve into pointers.

this video is a basic implementation an array.

- template for type and size
- indexing operator (square brackets)
- initialization (without constructors)
- making things constexpr
- allowing for range-based for-loops - tons of accessors
- structured binding support (tuple_size, get)

the standard array has publicly accessible internal C-array, but using it directly is undefined behavior.

```cpp
#include <cstdint>
#include <fmt/format.h>

template<typename Contained, std::size_t Size>
struct array{
    Contained _values[Size];

    constexpr Contained &operator[](std::size_t idx){
        return _values[idx];
    }
    constexpr const Contained &operator[](std::size_t idx) const {
        return _values[idx];
    }

    constexpr Contained *begin() {return _values;}
    constexpr const Contained *begin() const {return _values;}
    constexpr const Contained *end() const {return _values + Size;}
    constexpr const Contained *cbegin() const {return _values;}
    constexpr const Contained *cend() const {return _values + Size;}
    constexpr Contained &front() {return _values;}
    constexpr const Contained &front() const {return _values;}
};

//tuple support
template <typename class T, std::size_t N>
struct std::tuple_size<array<T, N>>: std::integral_constant<std::size_t, N>
{ };

template<std::size_t I, class T, std::size_t N>
struct std::tuple_element<I, array<T ,N>>
{
    using type =T;
};

template<std::size_t I, class T, std::size_t N>
[[nodiscard]] const T &get(const array<T,N> & a){
    return a[I];
}

int main()
{
    array<int, 5> data{2,3,4,5,6};
    for (auto value:data)
    {
        fmt::print("{}", value);
    }
    const [a,b,c,d,e] = data;
    fmt::print("{}",a+b+c+d+e);
    return data[0];
}
```

</details>

## C++ Weekly - Ep 360 - Scripting C++ Inside Python With cppyy

<details>
<summary>
Integrating C++ code inside python
</summary>

[Scripting C++ Inside Python With cppyy](https://youtu.be/TL83P77vZ1k)

a way to run c++ code from python

```python
import cppyy
cppyy.include("cppyy-test.hpp")
cppyy.gbl.go()
```

we can run the code as a script, or in an interactive way - we access our code, structs, templates... all from python.

</details>

## C++ Weekly - Ep 361 - Is A Better `main` Possible?

<details>
<summary>
Creating a `main` function with input arguments that we can operate on properly.
</summary>

[Is A Better `main` Possible?](https://youtu.be/zCzD9uSDI8c)

could we make the **main** function more modern? get some type safety, be able to use range operations, etc...

```cpp
int better_main(std::span<const std::string_view> args)
{
    return 0;
}
```

we could forward declare the better_main and use a vector as a contiguous container.

```cpp
int main(const int argc, char** argv)
{
    [[nodiscard]] int better_main(std::span<const std::string_view>) noexcept;

    std::vector<std::string_view> args (argv, std::next(argv, static_cast<std::ptrdiff_t>(argc)));

    return better_main(args);
}

[[nodiscard]] int better_main([[maybe_unused]] std::span<const std::string_view> args) noexcept
{
    return 0;
}
```

</details>

## C++ Weekly - Ep 362 - C++ vs Python vs Python (jit) vs Python With C++!

<details>
<summary>
Comparing performance between C++ and python (with C++).
</summary>

[C++ vs Python vs Python (jit) vs Python With C++!](https://youtu.be/lhqP50YVT-I)

coding conway's game of life. both in python and C++. checking normal python, compiled python (pypy), and cppyy (python which uses C++ objects). the real cost is the startup, but even just compiled python makes a huge difference.

</details>

## C++ Weekly - Ep 363 - A (Complete?) Guide To C++ Unions

<details>
<summary>
Overview of unions.
</summary>

[A (Complete?) Guide To C++ Unions](https://youtu.be/Lu1WsdQOi0E)

unions can be anonymous (nameless)

```cpp
int main() {
    union {
        int i;
        float f;
    };
    i = 42; // unnamed union
}
```

we can get some warnings from the compiler if we run it in compile time expression. a union is not aggregated, it has only a single active member. the size of the union is the largest member. there is no default active member. unions can have

```cpp
consteval auto use_union()
{
    union U{
        constexpr U(){}
        constexpr ~U(){}
        std::string s;
    };
    U u;
    u.s = std::string() // assignment, not constructor
    return u.s.size();
}

int main() {
    [[maybe_unused]] constexpr v =use_union();
    return v;
}
```

we can't do placement new inside a constexpr context, but in c++23 we could use `std::construct_at`. but it's not a good solution.

```cpp
 union U{
        constexpr U(){}
        constexpr ~U(){
            // s.std::string::~string(); //only if this was initialized
        }
        std::string s;
    };
    U u;
    //u.s = std::string() // assignment, not constructor
    //new (&u.s) std::string;// (placement new)
    std::construct_at(&u.s); // c++23
    return u.s.size();
```

at short, it's better to not use the union destructor.

we can use operator overloads (such as overloading the assignment operator for specific types), but that doesn't tell us which member is active. there is a proposal that will allow to check which is the active type at compile time, but that's still in the future.

it's better to use std::variant and std::optional in most cases. we don't want to handle the manual bookkeeping.

</details>

## C++ Weekly - Ep 364 - Python-Inspired Function Cache for C++

<details>
<summary>
Creating a cache wrapper for C++
</summary>

[Python-Inspired Function Cache for C++](https://youtu.be/lHnYSkZ7Cis)

single function cache, similar to the `@cache` decorator from the _functools_ module.

it is quite complicated, we need somewhere to store everything (function, parameters)

```cpp
#include <tuple>
#include <type_traits>
#include <map>

// intentionally copies the params
template<typename Func, typename ... Params>
auto cache(Func func, Params && ... params)
{
    using param_set= std::tuple<std::remove_cvref_t<Params>...>;

    param_set key {params...};

    using result_type = std::remove_cvref_t<std::invoke_result_t<Func, decltype(params)...>>;

    // this is not thread safe, basically a global
    static std::map<param_set, result_type> cached_values;

    using value_type = decltype(cached_values)::value_type; // std::pair<const param_set, result_type>

    auto iter = cached_values.find(key);

    if (iter != cached_values.end())
        return iter->second;
    }

    return cached_values.insert(value_type{std::move(key),func(std::forward<Params>(params)...)}).first->second;
}

int calculate(int i){
    return 42+i;
}

int main(){
    auto x = cache(calculate, 13);
    return x;
}
```

we can attempt this with the Fibonacci sequence, and in compiler explorer, we can check that we no longer get execution timeout.

</details>

## C++ Weekly - Ep 365 - Modulo (%): More Complicated Than You Think

<details>
<summary>
Unexpected behavior for the modulo operator when used with negative numbers.
</summary>

[Modulo (%): More Complicated Than You Think](https://youtu.be/xVNYurap-lk)

wat happens when using the modulo operator on negative numbers?\
we have `%` `fmod`, `remainder` (`std::fmod`, `std::reminder`), on uses flooring, other using truncating.c and c++ uses the truncated version by default, the Dart programming language uses the "Euclidean" version, which always returns a positive number. python uses "floored" modulo, which works great for wrapping around in positional indexing.

</details>

## C++ Weekly - Ep 366 - C++ vs Compiled Python (Codon)

<details>
<summary>
more timing comparisons between c++ and compiled python
</summary>

[C++ vs Compiled Python (Codon)](https://youtu.be/vXahGgWfzcA)

continuing an early episode about conway's game of life, Codon is a python compiler that creates native machine code.

codon doesn't use floor modulo.

</details>

## C++ Weekly - Ep 367 - Forgotten C++: `std::valarray`

<details>
<summary>

</summary>

[Forgotten C++: `std::valarray`](https://youtu.be/hxcrOwfPhkE)

a vector-like container that provides easy vectorization. operations on the vector occur to all elements.

```cpp
#include <valarray>
#include <vector>
std::valarray<int> get_data();
std::vector<int> get_vector_data();

int use_val_array(){
    auto data = get_data();
    data += 4;  // add 4 to all elements
}

void use_vector(){
    auto data = get_vector_data();
    for (auto & item: data) {
        item += 4; // add 4 to each element
    }
}
```

it's also possible to use `std::val_array` as the other operand, such as multiplying one by another. the standard allows operations to return other types, and it can also do lazy evaluation with them.

it is "forgotten", as there weren't many updates for it over the years, it didn't even get `constexpr` support.

</details>

## C++ Weekly - Ep 368 - The Power of template-template Parameters: A Basic Guide

<details>
<summary>
Passing a template type to a function, to be used inside a template.
</summary>

[The Power of template-template Parameters: A Basic Guide](https://youtu.be/s6Cub7EFLXo)

```cpp
#include <vector>
#include <list>

template<typename ResultType>
auto get_data(){
    // do something
    ResultType result;
    result.push_back(1);
    result.push_back(2);

    return result;
}

int main(){
    auto data_vec = get_data<std::vector<int>>();
}
```

but in this case, we can still have out function push in floating point numbers, and the conversion will be silent. we just want to pass the container type, but not the data type. std::vector or list, but we choose the data itself.

```cpp
template<template<typename Contained,
    typename Alloc=std::allocator<Contained>,
    ResultType>
auto get_data(){
    // do something
    ResultType<double> result;
    result.push_back(1.0);
    result.push_back(2.0);

    return result;
}

int main(){
    auto data_vec = get_data<std::vector>(); //no need to pass in the data type itself,
}
```

we can now use `constexpr` to be more precise.

```cpp
template<template<typename Contained,
    typename Alloc=std::allocator<Contained>,
    ResultType>
auto get_data(){
    ResultType<double> result;
    if constexpr(requires{result.reserve(1);}) { // is it possible to call reserve?
        result.reserve(3);
    }
    result.push_back(1.0);
    result.push_back(2.0);
    result.push_back(3.0);

    return result;
}
```

for non primitives types, we should use `emplace_back`, rather than `push_back`.

</details>

## C++ Weekly - Ep 369 - llvm-mos: Bringing C++23 To Your Favorite 80's Computers

<details>
<summary>
A project that compiles modern code into other languages assembly code.
</summary>

[llvm-mos: Bringing C++23 To Your Favorite 80's Computers](https://youtu.be/R30EQGjxoAc)

a project that generates 6502 (and other forms) assembly code from modern C++ code.

we can create a short example

```cpp
#include <cstdint>
#include <cstdio>

int main() {
    volatile std::uint8_t *border= reinterpret_cast<volatile std::uint8_t *>(53280); // take a pointer to somewhere in memory

    *border = 10;
    std::puts("Hello World!");
    for (int i=0; i<16000; ++i){
        ++(*border);
    }
}
```

compile it, and load it into an emulator and see that the message is printed and the screen changes colors.

</details>

## C++ Weekly - Ep 370 - Do Constructors Exist?

<details>
<summary>
Constructors don't actually exist!
</summary>

[Do Constructors Exist?](https://youtu.be/afDB4kpYnzY)

we can create a pointer to member function with the `using` syntax and explicitly writing the function address.

```cpp
struct S{
    S();
    ~S();

    int get();
    int other_get();
};

int main(){
    using mem = int(S::*)();
    mem foo = &S::get;

    S s;
    (s.*foo)(); // call function

    foo = &S::get_other; // reassign
    (s.*foo)(); // call function, this time a different one!

    s.~S(); // explicitly calling destructor, actually undefined behavior.
}
```

but we couldn't do it with constructors or destructors.

```cpp
int main(){
    using ctor_destructor = S (S::*)();
    ctor_destructor ctor = &S::S; // Error! can't take address of constructor.
    using mem_destructor = void (S::*)();
    mem_destructor des = &S::~S; // Error! can't take address of destructor.
}
```

we can't make a pointer to a constructor or a destructor. some example with templates

```cpp
struct S{
    template<int>
    S();
    ~S();
};

int main(){
    S s<int>; //can't be done.
}

```

example with cast operator syntax that look like constructors.

</details>

## C++ Weekly - Ep 371 - Best Practices for Using AI Code Generators (ChatGPT and GitHub Copilot)

<details>
<summary>
some best practices for using AI-code generator.
</summary>

[Best Practices for Using AI Code Generators (ChatGPT and GitHub Copilot)](https://youtu.be/I2c969I-KmM)

1. Assume the answer is wrong - even when it appears correct and is confident
2. Assume the answer is stolen from somewhere - be concerned about copyrights
3. Be explicit about the required guidelines - "prefer libfmt over iostream"
4. The answers represent the "average" quality of code on the internet - garbage in, garbage out.

The tools (chatGPT) can be used as a "rubber duck" and help with clarifying the problem, and as jumping point for inspiration and starting. it can also be used as a "second set of eyes", like figuring out a compiler error.

A very good use case is using the tools to generate a quick start for a project. like driver code, basic tests, cmake, etc... it's a good way to get started.

</details>

## C++ Weekly - Special Episode - Make Your Own Godbolt (Compiler-Explorer) With GPT-4 in 5 Minutes!

<details>
<summary>
asking chatGPT to create a python program that compiles c++ code and generates the assembly.
</summary>

[Make Your Own Godbolt (Compiler-Explorer) With GPT-4 in 5 Minutes!](https://youtu.be/XD3b2HA_7BQ)

more playing with chat-GPT.

the first version is functional, it creates python code that runs and gives decent results in a very fast manner.

</details>

## C++ Weekly - Ep 372 - CPM For Trivially Easy Dependency Management With CMake?

<details>
<summary>
Using a dependency manager in cmake.
</summary>

[CPM For Trivially Easy Dependency Management With CMake?](https://youtu.be/dZMU3iAPhtI)

CPM - a dependency manager for cmake.

the example from the [CPM.cmake github](https://github.com/cpm-cmake/CPM.cmake)

```cmake
cmake_minimum_required(VERSION 3.14 FATAL_ERROR)

# create project
project(MyProject)

# add executable
add_executable(main main.cpp)

# add dependencies
include(cmake/CPM.cmake)

CPMAddPackage("gh:fmtlib/fmt#7.1.3")
CPMAddPackage("gh:nlohmann/json@3.10.5")
CPMAddPackage("gh:catchorg/Catch2@3.2.1")

# link dependencies
target_link_libraries(main fmt::fmt nlohmann_json::nlohmann_json Catch2::Catch2WithMain)
```

before that we need to add the cpm.cmake to the project cmake, so there's a script for that.

</details>

## C++ Weekly - Special Edition - Getting Started with Embedded Python

<details>
<summary>
Showing off some devices which run embedded python.
</summary>

[Getting Started with Embedded Python](https://youtu.be/MIM_PTv_VjU)

devices that run python:

- MacroPad - acts as a keyboard or a mouse,
- BBC Micro:bit and Pico:ed - larger LED grid.
- (external power source) addons
- Calculators that run micro python and allow for programming.

Microsoft has an online tool that uses these devices to teach code.

</details>

## C++ Weekly - Ep 373 - Design Patterns in "Modern" C++ (2023)

<details>
<summary>
Giving some pointers towards implementing the observer pattern.
</summary>

[Design Patterns in "Modern" C++ (2023)](https://youtu.be/A_MsXney3EU)

starting with the observer pattern, which can easily be done poorly. there are some libraries which do this already. most implementations rely on reference counting, which might not be the best idea.

the observer pattern is meant to reduce coupling between components.

- signal and slot
- published subscriber
- observer and observable(subject)

registering a callback on an 'event'.

things to consider:

1. are the relationships known at compile time?
2. are the relationships dynamic or fixed? do we connect or remove observers after construction?
3. can the connections outlive one another? are the calls lifetime dependant?

</details>

## C++ Weekly - Ep 374 - C++23's `out_ptr` and `inout_ptr`

<details>
<summary>
Proxy object over pointer that we get from external sources
</summary>

[C++23's `out_ptr` and `inout_ptr`](https://youtu.be/DHKoN6ZBrkA)

new features that will come out in c++23. they exist to make interaction with C apis easier.

```cpp
extern "C"{
    void get_data(int **ptr) {
        int* result = (int *)malloc*(sizeof(int));
        *result=42;
        *ptr=result;
    }
}

int main() {
    std::unique_ptr<int, decltype([](int *ptr) {free(ptr);})> something = std::make_unique<int>(1);
    get_data(std::out_ptr(something));

    std::cout<< *something << '\n';
}
```

<cpp>std::out_ptr</cpp> creates a wrapper over the smart pointer and eventually calls the destructor and resets the pointer. a light-weight proxy object. because we allocate memory with malloc and remove it with delete we might get undefined behavior, so we need out smart pointer to call <cpp>free</cpp> as it's deleter. we do it by using lambda and decltype.

</details>

## C++ Weekly - Ep 375 - Using IPO and LTO to Catch UB, ODR, and ABI Issues

<details>
<summary>
Tools to prevent ABI violations
</summary>

[Using IPO and LTO to Catch UB, ODR, and ABI Issues](https://youtu.be/Ii-zuK1cd90)

having two versions of the same API, resulting in ABI violation.

IPO is "Inter-procedural optimization", LTO is "link time optimization", so if we enable them we can get warnings about those issues, so it's not only about better performance, it becomes something that helps us write better code.

</details>

## C++ Weekly - Ep 376 - Ultimate CMake C++ Starter Template (2023 Updates)

<details>
<summary>
2023 update for cmake starter template with the current best practice.
</summary>

[Ultimate CMake C++ Starter Template (2023 Updates)](https://youtu.be/ucl0cw9X3e8), [cmake_template](https://github.com/cpp-best-practices/cmake_template) repository.

the previous templates (gui_stater_template, cmake_conan_boilerplate_template) are archived, and replaced with the new cmake_template repository.

it has tooling support by default for address sanitizer, undefined behavior, IPO (link-time optimizer) and some others. it's also changed from conan package management to CPM. it also has "hardening" options, such as control flow guards, and using checks on vector index operations to avoid undefined behavior. these options make it harder for other to exploit weaknesses in the binary.

integrations with github actions for packaging, code quality checks, etc...

</details>

## C++ Weekly - Ep 377 - Looking Forward to C++26: What C++ Needs Next

<details>
<summary>
Jason's opinions on what C++ should have.
</summary>

[Looking Forward to C++26: What C++ Needs Next](https://youtu.be/L4PqCIMmc-A)

- More <cpp>constexpr</cpp> - more containers, more adapters, we get some of it in C++23, but there still gaps.
- Pattern Matching - enumerations, strings, tuples
- Contracts - another attempt, telling the compiler about pre and post conditions
- Reflection - compile time reflection, even for compile time `enum` to `string` conversion.
- removal of implicit conversions from the standard libary - this will never happen because it will break a lot of existing code.

</details>

## C++ Weekly - Ep 378 - Should You Ever `std::move` An `std::array`?

<details>
<summary>
calling move on a std::array to move the sub objects.
</summary>

[Should You Ever `std::move` An `std::array`?](https://youtu.be/56DMwqKffi0)

it's preferable to avoid moving. it's not that <cpp>std::move</cpp> is bad, it's just that not using it is better.

it's ok in this case:

```cpp
struct Holder {
    Holder(std::string data): data{std::move(data_);}

    private:
    std::string data;
};
```

if the object is and <cpp>std::array</cpp> of integers, then there's nothing happening. but if it's an array of strings, the move does help! because it calls the move operation on each element. there is even a clang tidy check for the "copy and move" idiom.

</details>

## C++ Weekly - Ep 379 - clang-tidy's "Easily Swappable Parameters" Warning - And How to Fix It!

<details>
<summary>
a new clang-tidy check that warns about parameters that are easy to confuse between on another.
</summary>

[clang-tidy's "Easily Swappable Parameters" Warning - And How to Fix It!](https://youtu.be/Zq4yYPG7Erc)

a new clang-tidy check that warns about parameters that are easy to confuse between on another.

```cpp
#include <algorithm>
auto someClamp(int min, int max, int value);

int main(int argc)
{
    return someClamp(1,100, argc);
}
```

but how do we handle this warning? we are really working on three integers. we can use strongly typed types, or use enum conversion which would force us to initialize our variables with the correct type.

```cpp
enum Min: int;
enum Max: int;
```

another option is to create a object for the function.

```cpp
struct Range {
    int min;
    int max;
};

auto myClamp(Range range, int value)->int;
int main(int argc)
{
    return myClamp(Range{1,100}, argc);
}
```

this might be more efficient, as the compiler is allowed more leeway of how to pass the struct into function.

</details>

## C++ Weekly - Ep 380 - What Are `std::ref` and `std::cref` and When Should You Use Them?

<details>
<summary>
//TODO: add Summary
</summary>

[What Are `std::ref` and `std::cref` and When Should You Use Them?](https://youtu.be/YxSg_Gzm-VQ)

shortcuts for creating reference wrappers, <cpp>std::reference_wrapper\<Type></cpp> and <cpp>std::reference_wrapper\<const Type></cpp>.

```cpp
#include <functional>

void func(int &i){
    ++i;
}

int main()
{
    int local_i = 0;
    auto bound = std::bind(func, local_i);
    bound();
    bound();
    return local_i;
}
```

this returns zero, not 2, because <cpp>std::bind</cpp> takes a copy, if we want to make it pass a reference, we need to create a reference wrapper.

```cpp
int main()
{
    int local_i = 0;
    //auto bound = std::bind(func, std::reference_wrapper<int>(local_i)); // long form
    auto bound = std::bind(func, std::ref(local_i)); // long form
    bound();
    bound();
    return local_i;
}
```

it also allows us to make member variables a reference. which is better than using a normal reference member declaration in how it handles moves and other operations.

```cpp
struct Data{
    std::reference_wrapper<int> value;
};
```

a third case is when using the standard library algorithms. they normally take the callable object by copy, so our original object isn't modifed.

```cpp
int main() {
    const std::array<int, 5> data{1,2,3,4,5};
    auto accumulator = [sum=0] (int value) mutable {
        sum += value;
        return sum;
    };
    std::for_each(data.begin(), data.end(), accumulator);
    return accumulator(0); // returns zero?
}
```

had we used <cpp>std::ref</cpp> when passing the lambda, then we would get the intended value.

</details>

## C++ Weekly - Ep 381 - C++23's `basic_string::resize_and_overwrite`

<details>
<summary>
filling the value of a string with a function to avoid unnecessary initialization of it.
</summary>

[C++23's `basic_string::resize_and_overwrite`](https://youtu.be/Ymm0yN_QUQA)

another feature meant for interacting with external data, such as files or C functions. if we have something like this:

```cpp
#include <string>

std::string make_data() {
    std::string result(10000,0);
    unsigned char current_c = 0;
    /* fill the data*/
    for (auto &c : result) {
        c = current_c;
        ++current_c;
    }
    return result;
}
```

we can replace it with a callable object, which avoids the <cpp>memset</cpp> step, and we communicate to the reader that we don't care about the inital value of the string.

```cpp
std::string make_data() {
    std::string result;
    result.resize_and_overwrite(10000,[](char *ptr, std::size_t size){
        unsigned char current_c = 0;

        for(std::size_t idx =0; idx < size; ++idx)
        {
            ptr[idx]= current_c;
            ++current_c;
        }

        return size;
    });

    return result;
}
```

</details>

## C++ Weekly - Ep 382 - The Static Initialization Order Fiasco and C++20's `constinit`

<details>
<summary>
A strange case of initialization order
</summary>

[The Static Initialization Order Fiasco and C++20's `constinit`](https://youtu.be/rEwijXgC_Kg)

using compiler explorer project mode, we set up a provider/consumer classes. assuming a global provider and consumer (it's not a good idea), because the static globals rely on one another, the order matters and depending on it, the global consumer might not have the global provider available. if we play with the cmake file, then we can manipulate the behavior.

`constinit` comes into play if we want to guarantee that the initialization happens when the program starts, it requires a `constexpr` constructor.

</details>

## C++ Weekly - Ep 383 - C++ Cross Training

<details>
<summary>
Becoming a better C++ programmer through learning other programmes
</summary>

[C++ Cross Training](https://youtu.be/9RxPRr-fk7Q)

"becoming a better runner by practicing a different sport"

it's also one of the item in best practices book, and there will a python book.

</details>

## C++ Weekly - Ep 384 - Lambda-Only Programming

<details>
<summary>
(not a best practices video), making everything a lambda.
</summary>

[Lambda-Only Programming](https://youtu.be/z5ndvveb2qM)

creating a flat map (dictionary) as a lambda expression.

```cpp
auto flat_map = [data= std::vector<std::pair<int, int>>()] (int input_key) mutable -> auto & {
    for (auto &[key, value]: data) {
        if (key == input_key) {
            return value.second;
        }
    }

    return data.emplace_back(input_key, 0).second;
};
```

and now we try making it into a variadic set, we need to expand and use the fold operation. we can also add `constexpr` when appropiate.

```cpp
auto flat_map_v2 = [data= std::vector<std::pair<int, int>>()] (auto ...&&input_key) mutable -> auto & {
    if constexpr (sizeof...(input_key) == 1) {
        for (auto &[key, value]: data) {
            if (key == (input_key, ...)) {
                return value.second;
            }
        }
        return data.emplace_back((input_key, ...), 0).second;
    } else {
        return data;
    }
};
```

we can also declare types in the lambda capture expression as an immediately invoked expression and use it in the lambda itself.

</details>

## C++ Weekly - Ep 385 - The Important Parts of C++20

<details>
<summary>
A quick run down of what c++20 features are viable to use.
</summary>

[The Important Parts of C++20](https://youtu.be/N1gOSgZy7h4)

**Designated initializers** - specify values to struct initialization by key using the dot notation. can't leave an uninitialized value or initialize them out of order.

Three way comparison (**spaceship operator**). requires the <cpp>\<compare></cpp> header.

**Concept** and **auto concepts** allow us to constrain our function to use only certain types.

The **Format** header allows for nice string formatting.

The **Source Location** header gives us library support (rather than macro support) for knowing where we are inside the source code itself, such as file and function.

There were also **Calender** updates to <cpp>std::chrono</cpp> to make working with time and dates easier.

<cpp>Ranges</cpp> were added and they have simplified working with lists in a standardized way without worrying about edge cases. we also got <cpp>std::span</cpp> to work with continues data.

Many more containers and algorithms were made compatible with <cpp>constexpr</cpp>, and there were additions to multithreading programming, such as <cpp>std::latch</cpp>, <cpp>std::barrier</cpp> and built-in semaphores.

<cpp>Modules</cpp> are new and important, but the tooling support is still lacking. <cpp>Coroutines</cpp> still require much boilerplate code to get working.

```cpp
struct S {
    int i;
    int j;
    float f;
    constexpr auto operator<=>(const S &) const = default;
};

void designated_initializers() {
    S s{.i=5,.j=42 .f=3.14f};
}

auto auto_concepts(std::integral auto lhs, std::integral auto rhs) {
    return lhs + rhs;
}

std::string make_log(int id, std::string_view description) {
    return std::format("Event {}: {}", id,description);
}

void use_ranges(const std::vector<int> &vec) {
    for (const auto &val : vec | std::ranges::views::drop(2))
    {
        // do things
    }
}
```

</details>

## C++ Weekly - Ep 386 - C++23's Lambda Attributes

<details>
<summary>
putting attributes on lambdas
</summary>

[C++23's Lambda Attributes](https://youtu.be/YlmxNJnone0)

in c++23, we can stick attributes onto lambda, like the <cpp> [[nodiscard]] </cpp> on the return value, we can also have attribute that apply to the _type of the function_. there is also the attributes we can stick on the parameters

```cpp
int main()
{
    auto l = [] () {return 42;};
    auto l_23 = [] [[nodiscard]] () [[gnu::deprecated]] {return 42;}; //c++23
}
```

</details>

## C++ Weekly - Ep 387 - My Customized C++ Programming Keyboard!

<details>
<summary>
Programming a device into a a CPP macro keypad.
</summary>

[My Customized C++ Programming Keyboard!](https://youtu.be/LwxBLG8aGlo)

A macro pad device that's been programmed to handle repetitive c++ code. a python code that handles key presses events.

(in my opinion, code snippets are easier)

</details>

## C++ Weekly - Ep 388 - My `constexpr` Revenge Against Lisp

<details>
<summary>
A scripting language based on modern C++.
</summary>

[My constexpr Revenge Against Lisp](https://youtu.be/NQEE0k9i7FA), [github](https://github.com/lefticus/cons_expr).

"Concepts of Programming Languages" book. Lisp (recursive only programming), looking at chai-script again, and re-inventing it into **cons_expr** - a modern scripting language which works with `constexpr` behavior.\
The tests run in compile time and runtime. has a gui thing that shows the internal, and this supports lisp-like expressions.

</details>

## C++ Weekly - Ep 389 - Avoiding Pointer Arithmetic

<details>
<summary>
Alternatives to Pointer Arithmetic.
</summary>

[Avoiding Pointer Arithmetic](https://youtu.be/YahYVRS1Ktg)

don't write code like this

```cpp
void func(const char* argv, const std::size_t argc) {
    const std::vector<char> args(argv, argv + argc); // pointer arithmetic
}
```

one option is using <cpp>std::next</cpp>. which hides the arithmetics away.

```cpp
void func(const char* argv, const std::size_t argc) {
    const std::vector<char> args(argv,
        std::next(argv, static_cast<std::ptrdiff_t>(argc)));
}
```

a different option is by using <cpp>std::span</cpp>. which has the advantage that it can work with ranges and pipings, as well as having <cpp>.subspan()</cpp> method.

```cpp
for (const auto &arg : std::span<const char>(argv, argc)) {
    std::cout << arg << '\n';
}
```

</details>

## C++ Weekly - Ep 390 - `constexpr` + `mutable` ?!

<details>
<summary>
Combining compile time objects with mutability.
</summary>

[`constexpr` + `mutable` ?!](https://youtu.be/67DenIV45xY?si=a3QV45xM67VQgMrG)

combining <cpp>constexpr</cpp> and <cpp>mutable</cpp> expressions. we might want to do it with a lambda.

```cpp
struct Lambda {
    int i = 0;

    constexpr auto ()() {
        return ++i;
    }
};

int main()
{
    auto l = [i = 0]() mutable {
        return ++i;
    };
    l();
    return l();
}
```

the two are equivelent, we could remove the <cpp>mutable</cpp> mark from the lambda and add <cpp>const</cpp> to the member function and then we couldn't change the return value.

(more about compile time members, constructors and mutability).

</details>

## C++ Weekly - Ep 391 - Finally! C++23's std::views::enumerate

<details>
<summary>
Enumerate over collections.
</summary>

[Finally! C++23's std::views::enumerate](https://youtu.be/HuRbLPRh-Nk?si=6HHyZcAkCmCrb7Cl)

C++ 23 will finally have an _enumerate_ capability: <cpp>std::ranges::views::enumerate</cpp> returns both the index (as long) and the value.

```cpp
int main()
{
    std::array<std::string_view, 5> data {"Jason", "Was", "Here", "To", "Enumerate!"};
    for (const auto [index, str] : data | std::ranges::views::enumerate)
    {
        std::cout << index << " " << str << '\n';
    }
}
```

</details>

## C++ Weekly - Ep 392 - Google's Bloaty McBloatface

<details>
<summary>
A binary size profiler by google.
</summary>

[Google's Bloaty McBloatface](https://youtu.be/MY5DTDc3e-I?si=9EuZ6lQM7fg4Ye8w)

a tool for profiling binary files, like identifying which parts of the binary contribute most to the its size (or virtual memory). such as the ".text" section for functions, ".rodata" for readonly data, ".strtab" and ".dynstr" for strings, and also how much of the binary is debug data. it's also possible to check which function takes up the most space by passing the `-d fullsymbols` flag

</details>

## C++ Weekly - Ep 393 - C++23's std::unreachables

<details>
<summary>
Marking code as unreachable to allow optimizations.
</summary>

[C++23's std::unreachable](https://youtu.be/ohMyb4jPIAQ?si=zcgnCav9kWUmgnY3)

another C++23 feature, <cpp>std::unreachable</cpp>, this raises undefined behavior and is used to optimize away branches.

```cpp
enum struct Options {Opt1, Opt2, Opt3};

std::string make_string_1();
std::string make_string_3();
std::string make_string_3();

std::string process(Options option)
{
    switch(option) {
        case Options::Opt1:
            return make_string_1();
        case Options::Opt2:
            return make_string_2();
        case Options::Opt3:
            return make_string_3();
    };
    std::unreachable();
}
```

but since an enum is an int, the function will also accept any integer value, marking this with <cpp>std::unreachable</cpp> allows the compiler to skip one of the check (optimizing) at release build, or have an undefined behavior.

</details>

## C++ Weekly - Ep 394 - C++11's Most Overlooked Feature: Delegating Constructors

<details>
<summary>
Constructors calling other constructors.
</summary>

[C++11's Most Overlooked Feature: Delegating Constructors](https://youtu.be/G5ewfxJ0KMU?si=leOMGvBQxCps3IJ4)

```cpp
struct S
{
    S() : S(42) {}
    S(int _x) : x{_x} {}

    int x;
};

int main()
{
    S s;
    return s.x;
}
```

> 1. If your find yourself duplicating code in a constructor: use a delegating constructor.
> 2. If _Any_ constructor completes, the object had it's lifetime begin and it's destructor is called.
> 3. Delegating constructors can be used for advanced lifetime management.
> 4. If you delegate, you cannot initialize any other members.

another option is to have a public constructor call the private one.

</details>

## C++ Weekly - Ep 395 - How Much is 100,000 Subscribers Worth?

<details>
<summary>
General YouTube stuff.
</summary>

[How Much is 100,000 Subscribers Worth?](https://youtu.be/wLXnH0Z08EU?si=EEMqcZPbS_S0jXem)

the growth of the channel, revenue streams.

the channel gains around 700 subscribers each month, had a bump at April 2020 with the doom port eight hours live stream and the PANDEMIC.

- adds - about 200$ dollars a month
- amazon affiliate - around 10$
- sponsorships - costs 600$ an episode, but this doesn't happen much
- Patreon - not specified, but not much.
- book sales and contracts - not going into details.

views mean money, not subscribers, clickbait titles work and bring in more views than other videos.

</details>

## C++ Weekly - Ep 396 - `emplace` vs `emplace_hint`! What's The Difference?

<details>
<summary>
Understanding the different forms of the operation.
</summary>

[emplace vs emplace_hint! What's the difference?](https://youtu.be/hW4NJF4RLnE?si=AoH4ER09ItnKxwf_), [github_issue](https://github.com/lefticus/cpp_weekly/issues/16), [benchmark](https://quick-bench.com/q/I9Wmi2dnRwmCeNo-2aLvP6aUTUw).

<cpp>emplace</cpp> creates an object in a specific place, avoiding copying and moving objects.

<cpp>emplace_back</cpp> in containers. <cpp>emplace_front</cpp> when supported (lists, deque), <cpp>emplace(iterator, args)</cpp> to create objects at a specific location, or use <cpp>emplace</cpp> in associative containers (<cpp>std::set</cpp>, <cpp>std::map</cpp>,
<cpp>std::multiset</cpp>, <cpp>std::multimap</cpp>, <cpp>std::unordered_set</cpp>, <cpp>std::unordered_map</cpp>, <cpp>std::unordered_multiset</cpp>, <cpp>std::unordered_multimap</cpp>) to have the object created at the appropiate location. the return value is a pair object with iterator and a boolean. <cpp>try_emplace</cpp> first checks if the key exists, and then passes the arguments to create the object. <cpp>emplace_hint</cpp> and <cpp>try_emplace_hint</cpp> take a location argument that acts as a hint where to construct the object. <cpp>emplace</cpp> in gcc standard library is defined with <cpp>emplace_hint</cpp>, but not in clang standard library.

```cpp
std::vector<std::string> vec;
vec.emplace_back(42, 'a'); // call the string constructor with the arguments to create a string of 42 times the 'a' character
std::list<std::string> list;
list.emplace_front("77");
list.emplace(list.begin(), "here");
std::set<std::string> set;
set.emplace(42, 'a');
std::map<std::string> map;
map.try_emplace("a", 42, 'a'); // insert object
map.try_emplace("a", 42, 'a'); // doesn't create the string object
map.emplace_hint(map.back(), "b", 42, 'a'); // doesn't create the string object
```

also used in <cpp>std::optional</cpp>, <cpp>std::variant</cpp>, <cpp>std::expected</cpp> and <cpp>std::any</cpp>.

```cpp
std::optional<std::string> opt;
opt.emplace(42, 'a');
```

</details>

## C++ Weekly - Ep 397 - std::chrono QuickStart With C++20 Calendars!

<details>
<summary>
Different clocks and durations in the standard library.
</summary>

[std::chrono QuickStart With C++20 Calendars!](https://youtu.be/I53iT3gPXrk?si=YemLxOaIRDkxN5v9)

inside the <cpp>chrono</cpp> header there are all sorts of clocks (system, utc, atomic and steady clock). the header is strongly typed, so we can't subtract between two clocks directly.

</details>

## C++ Weekly - Ep 398 - C++23's `zip_view`

<details>
<summary>
Iterate over multiple ranges.
</summary>

[C++23's zip_view](https://youtu.be/MVXGdwREo_E?si=3MG6-8Sb9bPOXXFv)

The <cpp>std::ranges::views::zip</cpp> from C++23 groups ranges together (one element from each range) and returns a **range of tuples of references to elements**.

```cpp
int main()
{
    std::array a{1,2,3,4,5,6};
    std::array b{7,8,9,10};
    std::array c{'a','b','c'};

    for (const auto &[p1, p2, p3] : std::ranges::views::zip(a,b,c))
    {
        std::cout << std::format("{}, {}, {}\n", p1, p2, p3);
    }
}
```

there is also <cpp>zip_transform</cpp> for manipulating the result.

</details>

## C++ Weekly - Ep 399 - C++23's `slide_view` vs `adjacent_view`

<details>
<summary>
Two options to return a sliding window view from ranges.
</summary>

[C++23's slide_view vs adjacent_view](https://youtu.be/czmGjH16Hb0?si=ah8I89cru-D2nSa6)

there are two views in c++23 which can be confused with one another:

1. <cpp>adjacent</cpp> view - returns a _range of tuples_ of references to elements in a sliding window. Window size is known at **compile time**.
2. <cpp>slide</cpp> view - returns a _range of ranges_ of references to elements in a sliding window. Window size is provided at runtime.

```cpp
int main()
{
    std::array a{1,2,3,4,5,6};

    // adjacent, range or tuples, compile time
    for (const auto &elem : std::ranges::views::adjacent<3>(a))
    {
        std::cout << std::format("{}, {}, {}\n", std::get<1>(elem), std::get<2>(elem), std::get<3>(elem));
    }

    // slide, range of ranges, runtime argument
    for (const auto &window : std::ranges::views::slide(a,3))
    {
        for (const auto &elem : windows)
        {
            std::cout << elem << ',';
        }
        std::cout << '\n';
    }
}
```

</details>

## C++ Weekly - Ep 400 - C++ is 40... Is C++ DYING?

<details>
<summary>
The state of C++ Language and the future of it.
</summary>

[C++ is 40... Is C++ DYING?](https://youtu.be/hxjSpasg3gk?si=m_OzMKQbEaUWUID6)

another podcast on the same topic by ADSP: [Is C++ Dying?](https://adspthepodcast.com/2023/10/06/Episode-150.html)

in 2022, Nasa suggested moving to "Memory Safe Programming Languages", which means away from C++ such as Golang (or C#, Ruby, Swift). There is also an article that documents the rise of C++ successor languages - such as [CppFront](https://github.com/hsutter/cppfront) by Herb Sutter, [Carbon](https://github.com/carbon-language/carbon-lang) or [Hylo](https://www.hylo-lang.org/).

however, seeing that many of the largest software suits in the world use C++ prominently:

- Microsoft Office
- Adobe PhotoShop
- Firefox
- Chrome
- Nearly all Databases
- Most compilers (including LLVM for Rust)
- TensorFlow for Machine Learning
- Open JDK
- V8 JavaScript engine
- Game Development Engines (Unreal, Godot, Unity, 3DE)

So C++ will continue to be around for a long time.

C++ itself begins in 1979 as "C with Classes", and got its' name in 1982/3. with the first standard in 1998. C++03 was a first revision, and then C++11 begins the modern era of C++. and since then there was a three year release cycle (c++14, c++17, c++20, c++23).

Programming Language ranking shows a decline in popularity until 2018, and a rise since then, with usage being in a steady state.

as for the language not evolving anymore, the Hello World example has it's first major change in C++23 with the formatting library and <cpp>std::println</cpp>.

```cpp
int main()
{
    std::cout << "Hello World\n"; // before c++23
    std::println("Hello World"); // since C++23
}
```

</details>
