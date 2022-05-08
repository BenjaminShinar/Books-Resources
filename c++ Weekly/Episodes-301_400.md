<!--
// cSpell:ignore fsanitize Fertig FTXUI NOLINT
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

## C++ Weekly - Ep 313 - The `constexpr` Problem That Took Me 5 Years To Fix!
<details>
<summary>
Getting compile time values to be usable in runtime.
</summary>

[The `constexpr` Problem That Took Me 5 Years To Fix!](https://youtu.be/ABg4_EV5L3w)

> (Compile-time views Into optimally sized comppile-time data. I'ts awesome, no really, trust me!)


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
    constexpr static auto length = get_legnth(make_string("hello jason, ",3));
    constexpr static auto str = get_array<legnth>(make_string("hello jason, ",3));
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
        return make_string("hello jason, ",3); // lambda
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
        return make_string("hello jason, ",3); // lambda
    };
    constexpr static auto str = to_right_size_array(make_data));
    constexpr static auto sv = std::string_view(str.begin(), str.end());
    fmt::print("{}: {}", sv.size(), sv);
}
```

this still isn't good enough, we still create two oject, a *std::array* and the *std::string_view*. there also a problem with having static variables in the *consteval* function.

so now we try other crazy stuff, we have a function that returns a refernce to the template argument. and now we got something that the compiler can optimize.

> Class non template type parameter
>
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
- which we use as static refernce 
- and then we use it to create the string_view.




</details>

## C++ Weekly - Ep 314 - Every Possible Way To Force The Compiler To Do Work At Compile-Time in C++
<details>
<summary>
Different ways to do compile-time calculations.
</summary>

[Every Possible Way To Force The Compiler To Do Work At Compile-Time in C++](https://youtu.be/UdwdJWQ5o78)

just making a value or function `constexpr` doesn't force the compiler to run it a compile time.

we can make the value *static*, which forecs the compiler to compute the value at compile time, but also requires it to be const. 

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
type/keyword | compile-time calculation| const |static | example |notes
----|----|----|---|---|---
`constexpr` | up to the compiler | yes | no | `constexpr auto value = get_value(1);` 
`constexpr static` | yes |yes | yes| `constexpr static auto value = get_value(1);` | must be static const
`constinit static` | yes |no |yes | `constinit static auto value = get_value(1);` | must be static
`consteval` function | yes |no |no | `auto value = get_value_consteval(5)` | argument must be compile time constants, function can't be used in run time.
template parameter | yes |no |no | `auto value = make_compile_time<get_value(10)>()` | using templates
wraping `consteval` function | yes | no | no | `auto value = as_constant(get_value(10))` |inner function can be reused
`consteval invoke` wrapper | yes | no |no | with moveable and callable 
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

continuing our string example, now supposedly we look at the *iterator*. we again need a const and non const version, and this is important if we want **for loops**.

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

if we play with compiler explorer and the optimizations, we can see that the compiler passes the *this* pointer as the first argument to the function, and if the member function is `const`, then the pointer is const.\
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

1. something that takes a callable and varidatic parameters.
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

the paramaters are copied each time, which might be a problem, and more than that, the function doesn't work for the basic case.
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

the problem is the copying, we don't handle forwarding. if we take references, we run into object lifetime issues. there might be a way to parametrize it (take copy of rvalue, refernce of lvalue), but it would probably quickly become a monsteroues code.
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
    return lefticus::calclate_things(some_data);
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
    return lefticus::calclate_things(some_data);
}
```

now we have two defnintios, so we either get a compile time error if we try to use them, or a linkage error. this protects us from undefined behavior.

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
    auto x = lefticus::calclate_things(some_data); // overload resolution
    auto old_x = lefticus::calclate_things(some_data); // overload resolution
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
this behavior stops us from performing move operatons, as we can't move from const, so we must perform a copy/assignment operator, which is a performance issue.

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

we have a techincally true but actually pointless warning about a move constructor.

*std::optional* has an implicit conversion, because it's a value type, rather than a pointer type.

```cpp
inline std::optional<S> make_value_5()
{
    //const S s;
    return std::optional<S>{std:in_place_t{}};
}
```

> if you have multiple different objects that might be returned, then you are also relying on implicit move-on-return (aka automatic move).


in the following case we have two constructors and a copy, because both options are initiliazed, if we would move the objects into the inner scopes, we could create just one and get move operations and return value optimization.
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

if we have an invarient data member which we can't change without breaking other stuff, then we should simply write an accessor/mutator.


</details>
