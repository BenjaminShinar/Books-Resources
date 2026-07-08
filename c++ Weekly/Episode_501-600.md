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

<details>
<summary>
Working back reflection into C++20 code.
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
<details>
<summary>
Template expansion.
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
Combining the power of reflection with the scripting langugage.
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

## C++ Weekly - Ep 509 - Can Lambdas Inherit Interfaces?

<details>
<summary>
Playing with lambdas.
</summary>

[Can Lambdas Inherit Interfaces?](https://youtu.be/f0heIju3udc)

creating a lambda that implements some abstract base class, so it functionally inherits from a class even though it itself is nameless. like how Java has.

something like
```cpp
struct BaseClass {
    virtual int do_work() = 0;
};

int main()
{
    auto lambda = []() : BaseClass {
        // body of the lambda
    };
}
```

This isn't like inheriting from a lambda using <cpp>decltype(lambda)</cpp>, we want the other way around. so a differenet question is to ask "can we inject behavior into a lambda".

starting with C++23, we can use 'deducing this' (like how we create recursive lambdas).

```cpp
template<typename ... Base>
struct overload : Base {
    using Base::operator()...;

    void update_values() {
        ++value;
    }

    int value =0;
};

template<typename ... Base>
overload(Base && ...) -> overload<Base...>;

int main()
{
    auto lambda = 

    overload d{
        [](this auto & self, int i) {
            self.update_values();
            std::cout << "i: " << i << " value: " << self.value << '\n'
            return self.value;
        },
        [](this auto & self, double d) {
            self.update_values();
            std::cout << "d: " << d << " value: " << self.value << '\n'
            return self.value;
        },
    };
    d(42);
    d(4.2);
}
```

so the lambdas share a state, inherit behavior, but don't know about other things.

</details>

## C++ Weekly - Ep 510 - The AMAZING Performance of array (and span)!

<details>
<summary>
getting better performance using compile time containers.
</summary>

[The AMAZING Performance of array (and span)!](https://youtu.be/u0mVnuUh46w)

<cpp>std::array</cpp> is the zero cost abstraction of a continues data.

```cpp
#include <array>
#include <vector>
#include <numeric>

int sum(const std::vector<int> & data) {
    return std::accumalte(data.begin(), data.end(), 0);
}

int sum(const std::array<int, 10> & data) {
    return std::accumalte(data.begin(), data.end(), 0);
}
```
we can compare the generated code between the two options, and see that the version with the array is much simpler than the vectror version, it does some loop unrolling (using architecture native, if possible). the compiler knows at compile time how much work it would end up doing, so it can optimize the generated code. this also works with other operations, such as <cpp>std::ranges::for_each</cpp> (though not always).

there's also a non-owning continous data, which is <cpp>std::span</cpp>, and the compiler operates on both.

</details>

## C++ Weekly - Ep 511 - `move(obj).fun()` vs `move(obj.fun())`

<details>
<summary>
move an object and call a member, or call a member and move the result.
</summary>

[`move(obj).fun()` vs `move(obj.fun())`](https://youtu.be/nLjrMcjsa0Y)

this can depend on the object we have, and what members it has.

note: a moved from object is required to be in a valid state, but not a specified one. for most implementations, a moved-from string ends up empty, but that's not a requirement.

```cpp
#include <optional>
#include <string>
#include <iostream>

int main()
{
    std::string data = "Hello World";
    std::cout << "Before: " << data << '\n';
    std::string other = std::move(data);
    std::cout << "After: " << data << '\n'; // empty

    std::optional<std::string> os{"optional"};
    std::cout << "Optional Before: " << os.value() << '\n';
    auto os2 = std::move(os).value();
    // we pick the && qualified reference
    std::cout << "Optional After: " << os.value() << '\n'; // empty
}
```

(he uses an custom object as an example), we see what operator is called depending on which thing is moved, if the class has reference qualified members (move operators), the result is the same at all cases. but if we don't have reference qualified members (only copy operators), then we might want to take ownership of the data.
(in c++23 we can use the 'deducing this' and <cpp>std::forward_like</cpp>).

</details>

## C++ Weekly - Ep 512 - `reinterpret_cast` is Finally Fixed!

<details>
<summary>
lifetime managment functions.
</summary>

[`reinterpret_cast` is Finally Fixed!](https://youtu.be/JtFVyXQ00PQ)

C++23 added <cpp>std::start_lifetime_as</cpp> and <cpp>std::start_lifetime_as_array</cpp>


this is undefined behavior, we create an integer and modify it as a float.
```cpp
#include <memory>

int main()
{
    int *i = new int(42);
    float *f = std::reinterpet_cast<float *>(i);

    return static_cast<int>(*f);
}
```


in C++26 we can cast in and out of void pointers (since <cpp>reintepert_cast</cpp> is not `constexpr`), and undefined behavior is not allowed during compliation time. so we can see this in action.

```cpp
consteval void bad_things() {
    int *i = new int(42);
    [[maybe_unused]] float *f = static_cast<float f*>(static_cast<void *>(i));
    *f = 4.5; // undefined behavior
    delete i;
}

int main()
{
    bad_things();
}
```

the new lifetime methods have defined behavior and are not undefined, we can treat it as if it was the new type.

</details>

## C++ Weekly - Ep 513 - How Many Ways Can You End a Program?

<details>
<summary>
Seven ways to end a program.
</summary>

[How Many Ways Can You End a Program?](https://youtu.be/ki9omnMeYS8)

1. return statement from main.
1. `exit`.
1. <cpp>std::abort</cpp>.
1. <cpp>std::exit</cpp>, <cpp>std::quick_exit</cpp>.
1. <cpp>std::terminate</cpp>.
1. crash for unefined behavior.
1. get killed for taking too many resources or by watchdog.

the first way, the book correct <cpp>EXIT_SUCCESS</cpp> return statement. this is the correct way, since it's taking the value from the platform. the second way is more commonly used, with `return 0`, which could potentially be wrong depending on the platform. and finally, we can can drop the return statement entirely and keep it empty, which is like returning the success code.

```cpp
#include <cstdlib>

int main()
{
    // return EXIT_SUCCESS; // explicit
    // return 0; // usually correct
    // do nothing!
}
```

we can also return a value, like `return 42`, which would be interpetted as failure, but that's still a way to end the program.

after the "boring" options to return from the `main`, we can move on to more intresting ways. such as calling `exit(EXIT_SUCCESS)` somewhere in the middle of the code. this is mean we don't have stack unwinding and we might lose out on destructor calls - static and thread locals object are cleaed up, but automatic objects are not.

```cpp
struct Lifetime {
    ~Lifetime() { std::puts("Destroyed!");}
};

void helper_function(bool condition) {
    if (condition) {
        exit(EXIT_SUCCESS);
    }
}

int main()
{
    lifetime l1; // not cleaned up
    static Lifetime l2; // cleaned up
    helper_function(true);
}
```

similiar to <cpp>exit</cpp>, there's <cpp>std::abort</cpp>, which raises a signal, this is also what the assert stament does. no cleanup happens at all. we can still interupt the signal handler if we want.


```cpp
int main()
{
    lifetime l1; // not cleaned up
    static Lifetime l2; // not cleaned up either in abort.
    assert(false);
    std::abort();
}
```

<cpp>std::quick_exit</cpp> only calls the handlers, which we can define with <cpp>std::atexit</cpp> and <cpp>std::at_quick_exit</cpp> (the underscore are differnt). the quick exit doesn't trigger the buffer flushing.

```cpp
Lifetime *l = new Lifetime();
void do_quick_exit() {
    std::atexit([](){ delete l ;});  // not triggered for quick exit
    //std::at_quick_exit([](){ delete l ;});  // triggered for quick exit, but not printing
    std::at_quick_exit([](){ delete l ; std::fflush(stdout)});  // triggered for quick exit, and now we see the printed output
    std::quick_exit(EXIT_FAILURE);
}

int main()
{
    do_quick_exit();
}
```

next we have <cpp>std::terminate</cpp>, which is implictly called with unhandled exceptions.


```cpp
void directly_terminate() {
    Lifetime local;
    sd::terminate();
}

void indirectly_terminate_maybe() {
    throw 42; // terminates if not caught
}

void very_indirectly_terminate() noexcept {
    // guarnteed to call terminate
    throw 42; // exception in noexcept function
}

int main()
{
    directly_terminate();
}
```

our next option it to cause undefined behavior and hope for a crash. this depends on the compiler, it can crash immediatly by generating an invalid opcode (which is what gcc does) or it can ignore the bad code (which is what clang does)

```cpp
void ub_crash_maybe() {
    int i* = nullptr;
    *i = 42;
}


int main()
{
    ub_crash_maybe();
}
```

the final option is to loop or consume all resources until the process is killed.

```cpp
void end_program() {
    while (true) {}
}

void end_program_linux() {
    std::vector<int> vec(1'000'000'000'000); // take all the resources
}
```

actual suggestions to handle undefined state if we want to protect against unknown state or exit really qucikyl

> 1. register the `at_quick_exit` handler.
> 2. in the handler, save the stareas much as possible to a new file, flush output.
> 3. register the sigabrt handler.
> 4. in the handler, call `quick_exit()`.
> 5. make the terminate handler call the quick exit.

</details>

## C++ Weekly - Ep 514 - C++26 on 1990 DOS?

<details>
<summary>
Compiling modern code on dos. doesn't work properly out of the box.
</summary>

[C++26 on 1990 DOS?](https://youtu.be/dtO94ifh7Ac)

C++ is portable, we can run a dosbox to emulate dos, and run `gpp --version` to see which gcc version we are running there, and we can check what features are avaialbe there. djgpp is a port of gcc to dos.

for some reason, we don't have all the stuff out of the box to run modern C++ code, so there's some problem about supporting long file names, and there's some hacking needed. even iostream doesn't work.

the support is limited.

check out the [cpp evolution over time](cppevo.dev) for some stuff.

</details>

## C++ Weekly - Ep 515 - Revolutionize Your Templates with `static_assert` of non-value-dependent Exprs

<details>
<summary>
Backport of static assertions changes.
</summary>

[Revolutionize Your Templates with static_assert of non-value-dependent Exprs](https://youtu.be/pwf45vaXm3Q)

until C++23, all static assertions had to valid at compilation time, regardless of whether the code block they were in was valid (the unused branch in `if constexpr` shouldn't exist).\
this is actually not a new lanugage feature, the change was backported so new compilers will have this behavior even in older C++ code.

```cpp
#include <type_traits>

template<typename Output>
Output Func() {
    if constexpr(std::is_same_v<Output, int>) {
        static_assert(false);
    } else {
        return 42;
    }
}

int main()
{
    func<float>(); // fails in C++20, even though the line shouldn't even be compiled
    
}
```

but this is usually the wrong tool for the task, we can delete function specializations or use template contratints to delete.

```cpp
auto func(int) = delete;
auto func(std::integral auto) = delete;
int main()
{
    func(10); // not allowed
}
```

</details>

## C++ Weekly - Ep 516 - C++26's User Generated static_assert Messages

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[C++26's User Generated static_assert Messages](https://youtu.be/CmfgZa-bcTg)

</details>

## C++ Weekly - Ep 517 - Tool Spotlight: ClangBuildAnalyzer

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Tool Spotlight: ClangBuildAnalyzer](https://youtu.be/gEQ5_FjCihA)

</details>

## C++ Weekly - Ep 518 - Online C++ Tools You Must See! (2026)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Online C++ Tools You Must See! (2026)](https://youtu.be/VAgC2bCwOQo)

</details>

## C++ Weekly - Ep 519 - initializer_list vs Initializer List

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[initializer_list vs Initializer List](https://youtu.be/8OlG6ya3kIY)

</details>

## C++ Weekly - Ep 520 - Yes, UB Really is THAT BAD

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Yes, UB Really is THAT BAD](https://youtu.be/atEP9wbuaL0)

</details>

## C++ Weekly - Ep 521 - Job Hunting and Optimizing Compilers with Jamie Pendergast (CppCast Ep404)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Job Hunting and Optimizing Compilers with Jamie Pendergast (CppCast Ep404)](https://youtu.be/vzar4IDKTys)

</details>

## C++ Weekly - Ep 522 - Don't Remove Code. =delete it!

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Don't Remove Code. =delete it!](https://youtu.be/gwwxD_l9T28)

</details>

## C++ Weekly - Ep 523 - Why I'm Still Using std::cout (on this channel)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Why I'm Still Using std::cout (on this channel)](https://youtu.be/TreruByxQWE)

</details>

## C++ Weekly - Ep 524 - Line Coverage vs Branch Coverage vs Path Coverage

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Line Coverage vs Branch Coverage vs Path Coverage](https://youtu.be/Gr0aI-TPRiQ)

</details>

## C++ Weekly - Ep 525 - Compiling at Compile Time with Daniel Nikpayuk (CppCast Ep405)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Compiling at Compile Time with Daniel Nikpayuk (CppCast Ep405)](https://youtu.be/JJjBQ95e28s)

</details>

