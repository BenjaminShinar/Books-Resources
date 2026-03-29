<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

## C++ Weekly - Ep 501 - Does C++26 Solve the constexpr Problem?

<details>
<summary>
P3491 - C++26 <cpp>define_static_</cpp> constexpr variables to be used in runtime.
</summary>

[Does C++26 Solve the constexpr Problem?](https://youtu.be/x3Z-k34u3Q8?si=Hi8SaQB0GOoi0G9A)

using [wg21.link](https://wg21.link/) as a redirection server, just add the Paper number ([p/n]xxxx), with or without the revision (PxxxxRx) to see papers. add "github" to see the github issue.

a paper accepted into C++26: <cpp>define_static_string, define_static_object, define_static_array</cpp>. this feature will rely on C++26 reflection, and we now have a clang reflection branch over on compiler explorer.

writing plain, simple, efficient C++ code that runs in compile time, and generates a string that persists to runtime. we wanted to declare a <cpp>constexpr std::string</cpp> in the main function and use it, but it wasn't allowed since it required data allocation, the workaround was messy and cumbersome, and it involved copying it into an array.\
the new functionality does it for us for strings and arrays, and it relies on ranges, currently it seems to mess up the symbol table and makes it much larger.

```cpp
#include <cstdio>
#include <meta>
#include <string>
#include <array>

// old way to do this.
constexpr auto make_static_string(std::string& dynamic_string)
{
    std::array<char, 256> retval{}; // 0 initialize
    std::copy(dynamic_string.begin(), dynamic_string.end(), retval.begin());
    return retval;
}

constexpr std::string make_string(const std::string& p1, const std::string& p2, const std::string& p3)
{
    return p1 + ' ' + p2 + ' ' + p3;
}

int main() {
    // old way
    static constexpr auto static_str = make_static_string(make_string("Jason", "Was", "Here"));
    std::puts(static_str);
    // new way
    static constexpr auto str = std::define_static_string{make_string("Jason", "Was", "Here")};
    std::puts(str);
}
```

</details>

## C++ Weekly - Ep 502 - Simple Reflection For C++20

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Simple Reflection For C++20](https://youtu.be/voljWhjl0bA?si=XoEocGHmjDr4MQvh)

native reflection will come in C++26, but we can work it into C++20 code. we need to hard code for every number of fields, because we don't have partial destructring, but this can work for different types. we just need to avoid copy-paste typos!

```cpp
#include <string_view>
#include <array>
#include <utility>
#include <format>

template<typename First, typename Second>
struct Pair {
    static constexpr std::array<std::string_view, 2> names {
        "first", "second"
    };
    First first;
    Second second;
};

template<typename Object>
void visit(auto Visitor, Object &&object) {
    if constexpr(object.names.size() == 0) {

    }
    else if constexpr(object.names.size() == 1) {
        auto &&[m0] = std::forward<Object>(object);
        Visitor(object.names[0], std::forward<decltype(m0)>(m0));
    }
    else if constexpr(object.names.size() == 2) {
        auto &&[m0, m1] = std::forward<Object>(object);
        Visitor(object.names[0], std::forward<decltype(m0)>(m0));
        Visitor(object.names[1], std::forward<decltype(m1)>(m1));
    }
}

int main()
{
    Pair<int, float> p1{2, 2.3f};
    visit([](std::string_view name, const auto &value) {
        std::puts(std::format("{}: {}", name, value).c_str());
    }, p1);
}
```

next episode will show how we do it for real in C++26.

</details>

## C++ Weekly - Ep 503 - The Amazing Power of C++26's Template For Statements
<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[The Amazing Power of C++26's Template For Statements](https://youtu.be/yaWiGLSDc64?si=oz-F4JdFQntp69Wj)

expansion templates, <cpp>template for</cpp>, a C++26 feature.\
will act as either ranged-for-loop on templated objects, or as a _destructing expression statement_. 

we can't run a for-loop over a tuple, since the elemnts aren't the same type.

```c++
#include <tuple>

std::tuple<int, char, double> get_tuple();

void use(double);
void use(int);
void use(char);

int main()
{
    for (const auto &val : get_tuple()) {
        // this is an error
    }
    template for (const auto &val : get_tuple()) {
        // this should work
        use(val); // conversions might happen here
    }
}
```

a <cpp>std::tuple</cpp> is something that can be destructred, so we can use the same `template for` on data types and 'visit' each of the elements.

this is somewhat equivelent.
```cpp
void use_data(const auto &data) {
    const auto&[...elem] = data;
    (use(elem),...);
}
```

the real strengh will be working over constant expressions, we could make the iteration item a constant expression, we need to do some workarounds, but we could do stuff during compile time.

</details>

## C++ Weekly - Ep 504 - Practical Reflection in C++26
<details>
<summary>
//TODO: add Summary
</summary>

[Practical Reflection in C++26](https://youtu.be/Mg_TBYppQwU?si=xqCT-25XjYv3Yy1K)

combining the power of reflection with the scripting langugage.

```cpp
namespace lefticus::interface {
    // defining functions
}

int main(){
    lefticus::cons_expr evaluator;

    // bind members
    bind<^^lefticus::interface>(evaluator);

    // test calling functions defind in the interface namespace
    evaluator.evaluate(R"(
        (q)
        (u 2 3)
        (print (myfloor (+ 3.2 13.9)))
    )");
}
```

the `^^` is the reflection operator, it creates a metadata object, the magic itself happens in the *bind* function. the `[:<>:]` is a splicing syntax for reflection, there are many functions in the <cpp>std::meta</cpp> library for reflection.

```cpp
template <auto Member>
constexpr auto bind_member(auto &engine) {
    if constexpr (!std::meta::is_special_member_function(Member)) {
        engine.template add<&[:Member:]>(std::meta::display_string_of(Member));
    }
}

template <auto Type>
constexpr auto bind(auto &engine) {
    static constexpr auto ctx = std::meta::access_context::unprivileged();

    template for (constexpr auto mem : std:define_static_array(members_of(Type, ctx)) {
        bind_member<mem>(engine);
    }
}
```
</details>

## C++ Weekly - Ep 505 - C++26's CNTTP bind Functions

<details>
<summary>
using `bind` for non-type template parameters.
</summary>

[C++26's CNTTP bind Functions](https://youtu.be/gIyuvqJnhi0?si=6o0loIodc_QPc88E).

the <cpp>std::bind_front</cpp> and <cpp>std::bind_back</cpp> functions.

```cpp
#include <functional>

int add(int x, int y) {
    return x + y;
}

void use(const auto f&);

int main()
{
    const auto old_bound = std::bind_front(&add, 2);
    //return old_bound(3);
    const auto bound = std::bind_front<add>(1);
    use(bound);
}
```

we no longer need to pass the pointer to the binding function, which might reduce the binary size.

</details>

## C++ Weekly - Ep 506 - Zero Cost Function Binding
<details>
<summary>
improving on last week.
</summary>

[Zero Cost Function Binding](https://youtu.be/9laCL5GixNk?si=7rFQI8J1juiIMhur)


continuing from last week example of CNNTP function binding. trying to implement it ourselves. we pass the callable, but we don't need to capture it. we use a forwarding reference, which can capture the parameters, and we eventually return a lambda closer object.

```cpp
#include <functional>

int add(int x, int y) {
    return x + y;
}

void use(const auto f&);

template<auto Func, typename... Param>
constexpr auto bind_front(Param && ... param) {
    return [...param = std::forward<Param>(param)]<typename ... Inner>(Inner && ... inner) {
        return Func(param..., std::forward<Inner>(inner));
    }
}

int main()
{

    const auto bound = bind_front<add>(1);
    use(bound);
}
```

we want to exapnd the functionality to also take additional parameters at compile time, and not just the callable object. and if there are no runtime bounded parameters, we can make the returned callable static, saving even more stack space. however, this isn't currently part of the standard bounding functions.


```cpp
void use(const auto f&);

template<auto Func, auto ... Constexpers, typename... Param>
constexpr auto bind_front(Param && ... param) {
    if constexpr (sizeof...(param) == 0) {
        return []<typename ... Inner>(Inner && ... inner) static {
            return Func(Constexpers..., std::forward<Inner>(inner));
    } else {
        return [...param = std::forward<Param>(param)]<typename ... Inner>(Inner && ... inner) {
            return Func(Constexpers..., param..., std::forward<Inner>(inner));
        }
    }
}

int main()
{

    const auto bound = bind_front<add,1>);
    use(bound);
}
```

</details>

## C++ Weekly - Ep 507 - Insidious Accidental Lambda Conversion

<details>
<summary>
Another example of implict conversion.
</summary>

[Insidious Accidental Lambda Conversion](https://youtu.be/b3fFxneoHso?si=r1tGhK1uuW7Dv-Zj)

(talking about the new book)

an accidental lambda conversion, submitted by a viewer.

```cpp
#include <cstdio>

int main()
{
    auto succeeded = [](){ return false; };

    if (succeeded) {
        std::puts("success");
    } else {
        std::puts("failure!");
    }
}
```

the bug is that we didn't call the lambda, we only defined it and checked the truthiness of the objcet, so since the address isn't null, it will always be considered True. getting warnings depends on the compiler.

we can set up some compiler flags to avoid boolean convesions (`-Wbool-conversion`, `-Wpointer-bool-conversion`, `-Wundefined-bool-conversion`) to warn us against implict conversions to booleans, but not all developers want these warnings.

it can also mess us with strings, where it prints 1 rather than the address of the object.
</details>

## C++ Weekly - Ep 508 - What if You're Windows Only?

<details>
<summary>
Building a project for Windows and older targets.
</summary>

[What if You're Windows Only?](https://youtu.be/aEuRTzj1qrM?si=iwBkkM6nFVkNV5Kx)

some project aren't designed to be multi-platform, and are strictly for windows.

the steps:


> 1. Port my build system to CMake
> 1. upgrade visual studio to the laters
>     a. if you need the older compiler, configure it in the cmake
>     a. also add new build system with the latest compiler
>     a. add the `clang-cl`
>     a. turning compatability flags with old featurse from MSVC.
> 1. start turning off compatability flags
> 1. isolate the code that uses windows extentsions libraries
>     a. move those clases into a separate static library.
> 1. isolate the code that doesn't need anything windows specific.
>     a. movet those classes into a different static library as well
> 1. add a linux/gcc/clang build that compiles and tests just those isolated components

wider testing with different compilers helps catching bugs, we can get better compilers for the CI, even if we don't use the most modern tools for the release (since we need to use the older tool chain).

</details>

