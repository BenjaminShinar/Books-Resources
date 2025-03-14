<!--
// cSpell:ignore Qlibs fftw rtsan noundef dispatchv resultv Wfunction Wperf nonallocating Wunknown
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Tooling

<summary>
15 Talks about tooling and supports for C++.
</summary>

- [ ] Beyond Compilation Databases to Support C++ Modules: Build Databases - Ben Boeckel
- [ ] Building Cppcheck - What We Learned from 17 Years of Development - Daniel Marjamäki
- [ ] C++/Rust Interop: Using Bridges in Practice - Tyler Weaver
- [ ] Common Package Specification (CPS) in practice: A full round trip implementation in Conan C++ package manager - Diego Rodriguez-Losada Gonzalez
- [x] Compile-time Validation - Alon Wolf
- [ ] Implementing Reflection using the new C++20 Tooling Opportunity: Modules - Maiko Steeman
- [ ] import CMake; // Mastering C++ Modules - Bill Hoffman
- [x] LLVM's Realtime Safety Revolution: Tools for Modern Mission Critical Systems - Christopher Apple, David Trevelyan
- [ ] Mix Assertion, Logging, Unit Testing and Fuzzing: Build Safer Modern C++ Application - Xiaofan Sun
- [ ] Secrets of C++ Scripting Bindings: Bridging Compile Time and Run Time - Jason Turner
- [ ] Shared Libraries and Where To Find Them - Luis Caro Campos
- [ ] What's eating my RAM? - Jianfei Pan
- [ ] What's new for Visual Studio Code: Performance, GitHub Copilot, and CMake Enhancements - Alexandra Kemper, Sinem Akinci
- [x] What's New in Visual Studio for C++ Developers - Michael Price & Mryam Girmay
- [ ] Why is my build so slow? Compilation Profiling and Visualization - Samuel Privett

---

### What's New in Visual Studio for C++ Developers - Michael Price & Mryam Girmay

<details>
<summary>
What's new and what's coming up in Visual Studio IDE.
</summary>

[What's New in Visual Studio for C++ Developers](https://youtu.be/Ulq3yUANeCA?si=voZfhAjwwzOx_544), [event](https://cppcon2024.sched.com/event/1gZgR/whats-new-in-visual-studio-for-c-developers), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/What's_New_in_Visual_Studio_For_Cpp_Developers.pdf)

The yearly talk by the Visual Studio team in Microsoft

#### Productivity

Github Copilot will be bundled with visual studio, mainly as a chatbot and editor suggestion. requires license from github. combines context from the project and opened files. we can also add context to the project with the ".github/copilot-instructions.md" file.\
We can ask copilot to improve memory layout for our classes, or to ask it to reduce the build time (with build insight).

> Build Insights - Analyze and optimize your build
>
> - Detailed analytics about your C++ builds
> - Integrated into Visual Studio
> - Visualize your include tree
> - Identify "expensive" included files
> - Find inlined functions that bloat your binaries

#### Game Development

direct support for unreal engine 5 projects, better integration and a dedicated toolbar.

#### MSVC Toolchain

improved support for C++23 features, work towards C++26 features. improvements to the address sanitizer, also integrated with copilot.\
better integration with the vcpkg package manager.

#### Debugging, Cross-Platform & Source Control

> - CMake Debugger in Visual Studio
> - Remote File Explorer for Linux
> - Target View Improvements
> - Automatically Install WSL from Visual Studio
> - Debug Linux Console Apps in Integrated Terminal

Better source control integration with popular repository hosting platforms. more copilot stuff to create commit messages.\
Better experience for connecting to remote server systems. also running tests on remote machines and modify files over there.

</details>

### What's New for Visual Studio Code: Performance, GitHub Copilot, and CMake Enhancements - Alexandra Kemper, Sinem Akinci

<details>
<summary>
What's new and what's coming up in Visual Studio Code and the C++ extention.
</summary>

[What's New for Visual Studio Code: Performance, GitHub Copilot, and CMake Enhancements](https://youtu.be/pjarNT2YgSQ?si=Q5n85mH93Q3Ppxzu), [event](https://cppcon2024.sched.com/event/1gZgQ/whats-new-for-visual-studio-code-performance-github-copilot-and-cmake-enhancements), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/What's_New_For_Visual_Studio_Code.pdf).

github copilot integration, cmake, etc...

adding fuzzy search support, making intelliSense faster, faster project start up.

chat copilot context using participants in the chat, commands with the `/` prefix like "/fix", "/explain" and "/tests".

better cmake presets, workflows. multi window support.

support for LLMs for vsCode extensions - language model API. chat participants and API changes.

</details>

### Compile-Time Validation in C++ Programming - Alon Wolf

<details>
<summary>
Trying some compile-time code validity checks.
</summary>

[Compile-Time Validation in C++ Programming](https://youtu.be/jDn0rxWr0RY?si=h6p5wxMOovSG-iDh),[slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Compile-Time_Validation.pdf), [event](https://cppcon2024.sched.com/event/1gZgF/compile-time-validation).

starting with some defintions by ISO or wikipedia:

> - Software Validation: - "Confirmation, through the provision of objective evidence, that the requirements for a specific intended use or application have been fulfilled"
> - Security "Resistance to intentional, unauthorized act(s) designed to cause harm or damage to a system"
> - Memory Safety - "Memory safety is the state of being protected from various software bugs and security vulnerabilities when dealing with memory access, such as buffer overflows and dangling pointers"
> - Software Safety - Ability of software to be free from unacceptable risk. [...] It is the ability of software to resist failure and malfunctions that can lead to death or serious injury to people, loss or severe damage to property, or severe environmental harm."

we can have memory safety issues, such as memory invalidation (through pointers, references, iterators) or going out of bounds. we can also have options for arbitrary code injections from user input.

```cpp
void foo() {
    vector<int> vec = { 0, 1, 2, /*...*/ };
    auto& ref = vec[0];
    vec.push_back(42);
    cout << ref; // ref may be invalid
}

void foo() {
    int index, value;
    cin >> index >> value;
    vector<int> vec = { 0, 1, /* ... */ };
    vec[index] = value;
}

void foo() {
    string str;
    cin >> str;
    db.run("SELECT * FROM Users WHERE name = " + str);
}
```

different kinds of application have different performance needs and focuses, some require low latency, some must have high throughput and scalability, while others focus on battery usage or electricity consumption.

validations can be done before executing the code (static, compile time) or during (runtime). compile time checks are a subset of static checks.

for example, we can have this code, a potential out-of-bounds access, we can check this statically, but there's also a runtime bounds check with the <cpp>at()</cpp> method of the container.

```cpp
void foo1(){
    std::vector<int> vec = get_vec();
    size_t index = get_index();
    vec[index] = 42;
}

void foo2(){
    std::vector<int> vec = get_vec();
    size_t index = get_index();
    vec.at(index) = 42;
}
```

for performance, we can run checks during the program execution or analyze the source code to detect possible bottlenecks.

```cpp
void must_be_fast_runtime() {
    using namespace std::chrono;
    auto start = high_resolution_clock::now();
    /*...*/
    auto end = high_resolution_clock::now();
    validate_performance(start, end);
}

void must_be_fast() {
    /*...*/
    can_slowly_read_huge_file();
}
```

#### Compile-time Validation

> Detecting errors early in the development pipeline reduces costs, saves time, minimizes risk, and improves efficiency.

we want to detect the error as soon as possible, and have the error be clear, informative and accurate.\
we sometimes need to choose between having high performance and flexible programs and having safe programs. C++ is usually used for cases that need the high-performance of a low-level language. we can have error detection at compile time with error reporting at runtime.

```cpp
void foo_1() {
    auto error = detect_error();
    if (error) {
        report_error(error);
    }
}

void foo_2() {
    constexpr auto error = detect_error();
    if constexpr (error) {
        report_error(error);
    }
}
```

one limitation of <cpp>static_assert</cpp> is that it must use string literals, which limits the information that can appear in the error message.

```cpp
void foo() {
    constexpr auto error = detect_error();
    static_assert(!error, "error message");
}
```

we can try working around this by using templates, this prints out the custom error object, but the error itself is the <cpp>sizeof</cpp> comparison.

```cpp
struct custom_error {};
void foo() {
    constexpr auto error = std::optional(custom_error{});
    if constexpr (error) {
    report_error<*error>();
    }
}

template<auto error>
constexpr auto report_error() {
    static_assert(sizeof(error) == 0);
}
```

our next attempt moves the error struct into the error line, and we can add custom compile time fields to our error.

```cpp
inline constexpr auto always_false = sizeof(error) == 0;
template<auto error>
constexpr auto report_error() {
    static_assert(always_false<error>);
}

struct invalid_index {
    int index;
};
report_error<invalid_index{42}>();
```

we can also use a fixed width string, but this has different behavior in clang/gcc and MSVC.

```cpp
template<int N>
struct fixed_str {
    constexpr fixed_str(const char(&str)[N]) {
        std::copy(str, str + N, data);
    }

    char data[N] = {};
};

report_error<fixed_str("Hello Cppcon :)")>();
```

for C++26, we can use user generated errors in static asserts, this requires an object with the `.size()` and `.data()` members.

#### Compile-Time Unit Tests

> Unit tests are automated tests written to validate that individual components of a program function as expected. Some C++ computations run at compile-time by using constexpr, consteval, or template metaprogramming.\
> These compile-time components can also be tested at compile-time.

we can simply write tests as <cpp>static_assert</cpp> statements, inside immediately invoked lambdas.

```cpp
static_assert(([]{
    static_assert(foo() == 42, "test failed");
    // more unit tests
}(), true));
```

or we can use a library such as <cpp>Qlibs++</cpp> to do components and fetuses validation.

#### Consistency Validation

we can make sure we always update the switch statement for enums, so we don't add a default case. we can use the <cpp>magic_enum</cpp> library to count the number of actions.

```cpp
enum class action {
    jump,
    fly
};
void on_action(action user_action) {
    switch(user_action) {
        static_assert(magic_enum::enum_count<action>() == 2);
        case action::jump:
            jump(); break;
        case action::fly:
            fly(); break;
    }
}
```

we could also convert the switch statement to a `magic_enum::enum_switch` which uses compile-time switch statement and makes sure all cases are handled.

the C++26 reflection proposal would allow use to do the same thing natively.

```cpp
template<class E>
constexpr auto enum_switch(auto callback, E arg) {
    return [: expand(enumerators_of(^E)) :] >> [&]<auto value>() {
        if (arg == [:value:]) {
            return callback.template operator()<[:value:]>();
        }
    };
}
```

#### Functional Programming and Metaprogramming

> - Immutability: Data is immutable, meaning once created, it cannot be changed. Instead of modifying existing data, you create new data structures with the desired changes.
> - Function Composition: Combining simple functions to build more complex functions. This is often done using function composition operators.
> - Monads: Encapsulates computations with context, allowing for the chaining of operations while managing side effects or state through a standardized interface.

function composition can allow us to validate the composed function properties and of the arguments.\
(something about <cpp>std::expected</cpp>).

using chain of context to detect dangling pointers.

we can combine functional programming with Stateful Metaprogramming.\
doing some stuff to detect issues during compile time - code branches, state changes, etc..\
tracking state, actions, changes in a meta-programming compile time operations and applying rules on them. for example, implementing Rust reference-borrowing checks. at any point in time, an object can have multiple read reference, or a single mutating reference. having a two mutating references is invalid, and so is having a mutating reference and a reading reference.\
control flow validation, pointer validation, performance validation (using a slow allocator when a faster one could be used instead).

C++26 reflection will make recording actions much easier and more generic, and allow us better rules, and we could create the proxy types through it.

#### Circle - Lifetime Safety

<cpp>Circle</cpp> is a C++ compiler extention by *Sean Baxter* with language extensions, we can mark functions and blocks as safe and unsafe, which would allow/disallow some operations, there are also rules for lifetime annotations.

there are other proposals for safety features in the C++ language.
</details>

### LLVM's Realtime Safety Revolution: Tools for Modern Mission Critical Systems - Christopher Apple, David Trevelyan

<details>
<summary>
A compile-time and runtime-sanitizer tool for realtime code.
</summary>

[LLVM's Realtime Safety Revolution: Tools for Modern Mission Critical Systems](https://youtu.be/KvhgNdxX6Uw?si=dPZCqvjyq11Rq3kq), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/LLVMs_Realtime_Safety_Revolution.pdf), [event](https://cppcon2024.sched.com/event/1gZgL/llvms-realtime-safety-revolution-tools-for-modern-mission-critical-systems), [realtime sanitizer repository](https://github.com/realtime-sanitizer/rtsan).

#### Realtime Programming

defintion - provide the right answer in the right time. consequences of dropping a deadline can be missing data (input or output), and can be life threatening in critical applications.

> Real-time programs must guarantee response within specified time constraints, often referred to as "deadlines".
>
> Worst case execution time must be..
>
> 1. deterministic
> 2. known in advance
> 3. independent of input data
> 4. shorter than the given deadline

for example, <cpp>malloc</cpp> is usually 10 nanoseconds (ns), but can also take up to 1,000,000 ns in some cases. this means it can't be used when deadlines are in the range of milliseconds.\
there are other knows calls with non-determistic execution time, they also block the progress of the processing. therefore, they are prohibited for real time programming.

1. system calls
2. allocations
3. <cpp>mutex</cpp> locks/unlocks
4. thrown exceptions
5. Indefinite waits (CAS loops, infinite loops)
6. others...

CAS - compare and swap.

it's sometimes easy to find the violations, but for other cases, the violation might be inside 3rd party code, or depend on the use case. for example, operating on a container might mean allocating a new node or triggering a resize, or calling a destructor. even a lambda might be creating <cpp>std::function</cpp> object under the hood. and even when a piece of code is safe when you use it, it might be changed by someone else and become non-deterministic, and we never know what happens in 3rd party libraries.

```cpp
void process_audio()
{
    numbers[1] = 2; // what kind of container is this?
}

void dispatch()
{
    auto const x = input_array(); // what if this changes?
    auto const y = output_array();

    post_report([x,y](auto & data) { 
        data.input = x;
        data.output = y;
    }); // does this allocate?
}

void process_audio()
{
    fftw_execute(plan); // third party library
}
```

#### Existing Strategies

we currently need to rely on

> - Shared experience
> - Code review
> - Profilers and debuggers
> - <cpp>static_assert</cpp>
> - Documentation

which have the problems of...

> - Getting experience takes a long time
> - Code review is prone to human error
> - Profiling/debugging is a manual process
> - Static assertions are limited
> - Documentation goes out of date
> - What about pre-built dependencies?

what if we could have a tool? something that can detect violations, even from code we bring from elsewhere? we would want it to be part of the pipeline.\
there are two tools coming soon:

1. RealtimeSanitizer - runtime
2. Performance Constraints - compiletime

#### RealtimeSanitizer

using sanitizers is important, in this example, an address sanitizer (`clang++ -fsanitize=address main.cpp`) would detect that we use the wrong index.

```cpp
#include <vector>
int main()
{
    auto v = std::vector<int> (16);
    return v[16];
}
```

so, we would like a sanitizer to detect code that isn't realtime safe in the manner (`clang++ --fsanitize=realtime main.cpp`), we declare the realtime functions with an attribute <cpp>[[clang::nonblocking]]</cpp>.

```cpp
float process(float x) [[clang::nonblocking]]
{
 auto const y = std::vector<float> (16); // this allocates!
 //...
}
```

Under the hood, the sanitizer instruments and intercepts the calls, and replaces any known blocking calls with errors. this a two step process, compiling and tracking with a runtime library.

(something like this)

```cpp
void __rtsan_realtime_enter() { /**/ }
void __rtsan_realtime_exit() { /**/ }

INTERCEPTOR (void *, malloc, size_t size) {
 if (is_in_realtime_context()):
 print_stack_and_die("malloc");
 return REAL(malloc)(size);
}

/* original code*/
int dispatch() [[clang::nonblocking]]
{
    return calculate_result();
}

/* compiled with sanitizer */
define noundef i32 @_Z8dispatchv() #1 
{
    call void @__rtsan_realtime_enter()
    %1 = call noundef i32 @_Z16calculate_resultv()
    call void @__rtsan_realtime_exit()
    ret i32 %1
}
```

in most cases, these calls are defined in the <cpp>libc</cpp> library, and are calling the kernel.

- allocations
- threads and sleep
- fileSystems and streams
- sockets

#### Performance Constraints

we also have compile time approach, we define the code we want to constrain, and compile with special flags.

attributes:

- <cpp>[[clang::nonallocating]]</cpp>
- <cpp>[[clang::nonblocking]]</cpp>

compilation flags:

- `-Wfunction-effects`
- `-Wperf-constraint-implies-noexcept`

this is an hierarchy, not blocking implies not allocating, which requires the call to not throw exceptions (<cpp>noexcept</cpp>).\
the "perf-constraint-implies-noexcept" flag warns that we didn't mark the functions as not throwing, while the "function-effects" flag checks that all our marked functions only call functions with the same contraint or stricter.\
To make things easier, if the called function is defined in the same translation unit, the compiler can infer if the function satisfies the requirements, even without explicitly marking it. for 34d party libraries, we can override this be declaring the function again and adding the attribute. this is a risky behavior, and can cause compilation to pass, even if the called function uses non realtime calls inside it.

```cpp
// third_party.h
void defined_elsewhere();

// main.cpp
void defined_elsewhere() [[clang::nonblocking]] // manual declaration, after careful review of the source code

void process() [[clang::nonblocking]] {
 defined_elsewhere();
}
```

note: neither the compilation nor the sanitizer can guarantee realtime performance. there are known blind spots:

> - No guarantee of processor time.
> - No guarantee your code runs faster than allotted time.
> - No detection of hand-written assembly system calls.
> - Not all libc wrapper functions implemented.
> - No detection of nondeterministic loops.
>   - Infinite loops
>   - Nondeterministic loops (CAS - compare and swap)
> - MisDeclared functions. (for performance constraints)

#### Comparing and Contrasting

the two tools are designed to work together, and to solve the same problem, but there are some differences.\
one is compile time, one is runtime, for the runtime sanitizer, it only detects code that runs, so it needs to hit every path (code coverage). the sanitizer is more prone to false negatives (misses) - if the blocking path isn't hit, or if the system call isn't yet intercepted, we might miss it. the performance constraints is stricter, and can lead to false positives (false alarm), for example, clearing and then pushing to a <cpp>std::vector</cpp> might be non-blocking, if the container is known to have enough capacity reserved. the two tools would produce different results.

```cpp
int main()
{
    std::vector<int> v;
    v.reserve(512);
    dispatch(v);
}

void dispatch(vector<int>& v)
// noexcept[[clang::nonblocking]]
{
    v.clear();
    v.push_back(3);
}
```

Both options have costs, the runtime sanitizer has run costs, intercepting calls, adding some operations, and can interfere with inlining optimizations. there are additional checks for determining if we are inside a nonblocking context.  the compile-time flags require converting the code, adding the attributes all across the codebase. adding the attribute on one function can cascade to require adding it on many more functions, since they are called internally.\
both options can be disabled locally, either with a macro to disable the compilation flags, or an special object to disable runtime checks (uses RAII).

```cpp
// macro to allow non blocking
#define NONBLOCKING_UNSAFE(...) \
 _Pragma("clang diagnostic push") \
 _Pragma("clang diagnostic ignored \"-Wunknown-warning-option\"") \
 _Pragma("clang diagnostic ignored \"-Wfunction-effects\"") \
 __VA_ARGS__ \
 _Pragma("clang diagnostic pop")

void process() noexcept [[clang::nonblocking]]
{
 NONBLOCKING_UNSAFE(foo()); // use macro
}

#include <sanitizer/rtsan_interface.h>

void lock_error_mutex(std::mutex& m)
{
    __rtsan::ScopedDisabler disabler{}; // disable sanitizer
    m.lock();
}

void process() noexcept [[clang::nonblocking]]
{
    if (buffer_overflow) 
    {
        lock_error_mutex(m);
    }
}
```

for now, the compiletime performance constraints is only used with llvm (clang), so other compilers like gcc can't use it yet. for the runtime sanitizer, it can be used as a standalone by enabling a flag and inserting a special object (also RAII) in the code anywhere the attribute would gone into. we need to build and link the code with the library.
</details>
