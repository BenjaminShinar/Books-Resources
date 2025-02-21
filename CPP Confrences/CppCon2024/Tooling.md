<!--
// cSpell:ignore Qlibs
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
- [ ] LLVM's Realtime Safety Revolution: Tools for Modern Mission Critical Systems - Christopher Apple, David Trevelyan
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
