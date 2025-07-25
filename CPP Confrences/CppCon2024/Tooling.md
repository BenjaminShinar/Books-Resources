<!--
// cSpell:ignore Qlibs fftw rtsan noundef dispatchv resultv Wfunction Wperf nonallocating Wunknown perfetto IWYU Wirth valgrind dhat jemalloc lfoo Xlinker lopencv Luabind chaiscript monostate libfuzzer lclang fuzztest cxxmodules RTTR repr rustc
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Tooling

<summary>
15 Talks about tooling and support for C++.
</summary>

- [x] Beyond Compilation Databases To Support C++ Modules: Build Databases - Ben Boeckel
- [x] Building Cppcheck - What We Learned From 17 Years Of Development - Daniel Marjamäki
- [x] C++/Rust Interop: Using Bridges In Practice - Tyler Weaver
- [x] Common Package Specification (Cps) In Practice: A Full Round Trip Implementation In Conan C++ Package Manager - Diego Rodriguez-Losada Gonzalez
- [x] Compile-Time Validation - Alon Wolf
- [x] Implementing Reflection Using The New C++20 Tooling Opportunity: Modules - Maiko Steeman
- [x] Import Cmake; // Mastering C++ Modules - Bill Hoffman
- [x] Llvm's Realtime Safety Revolution: Tools For Modern Mission Critical Systems - Christopher Apple, David Trevelyan
- [x] Mix Assertion, Logging, Unit Testing And Fuzzing: Build Safer Modern C++ Application - Xiaofan Sun
- [x] Secrets Of C++ Scripting Bindings: Bridging Compile Time And Run Time - Jason Turner
- [x] Shared Libraries And Where To Find Them - Luis Caro Campos
- [x] What's Eating My Ram? - Jianfei Pan
- [x] What's New For Visual Studio Code: Performance, Github Copilot, And Cmake Enhancements - Alexandra Kemper, Sinem Akinci
- [x] What's New In Visual Studio For C++ Developers - Michael Price & Mryam Girmay
- [x] Why Is My Build So Slow? Compilation Profiling And Visualization - Samuel Privett

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

starting with some definitions by ISO or wikipedia:

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

definition - provide the right answer in the right time. consequences of dropping a deadline can be missing data (input or output), and can be life threatening in critical applications.

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

### Building CppCheck - What We Learned from 17 Years of Development - Daniel Marjamäki

<details>
<summary>
Lessons learned from developing a static analysis tool and how the design changed over time.
</summary>

[Building Cppcheck - What We Learned from 17 Years of Development](https://youtu.be/ztyhiMhvrqA?si=oZ57Lnm5h9zqvrgQ), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Building_Cppcheck.pdf), [event](https://cppcon2024.sched.com/event/1gZdy/building-cppcheck-what-we-learned-from-17-years-of-development), [cppCheck repository](https://github.com/danmar/simplecpp).

the creator of CppCheck, a static analyzer tool. There is a focus on keeping the library portable, so it shouldn't require the latest cpp version.

there was a problem with static analysis tools, they didn't play well with compiler extensions, and they would have a high rate of false positives.\
the development of the tool started with perl and regular expressions (regex), but quickly moved to be written in C++ itself. started with simple checkers to find common problems:

> - Look for `X >= '0' && X <= '9'` and recommend to use <cpp>isdigit</cpp> instead.
> - Warn if <cpp>memset</cpp> is used on class
> - Warn about includes that are not needed
> - Redundant condition: `if (ptr) delete ptr;`
> - Member variable that is not initialized in constructor

this was accomplished with token list, part of which was to keep the tool compiler agnostic, so it doesn't matter if there are compiler extensions.\
there was a spike in downloads of the tool when a wikipedia page about static analyzers was created (and included cppCheck), this also brought more bug reports and fixes.

over time, there was more power in the infrastructure, rather than in the individual checkers, this includes a symbol database, also instantiating templates and using inline code to reduce cpu. next they add an AST (abstract syntax tree), and a generic dataflow analysis. adding valueTypes analysis.

in 2014, they found a bug in libXFont library that could allow unprivileged user to attain root privileges in some cases. the bug was 22 years old, and this also caused a rise in attention to the tool. the bug is that using sscanf can allow for overwriting data if the buffer isn't large enough.

```cpp
char charName[100];

if (sscanf((char *) line, "STARTCHAR %s", charName) != 1) {
    bdfError("bad character name in BDF file\n");
    goto BAILOUT; /* bottom of function, free and return error */
}
```

The philosophy of the tool changed over time, it started with the goal of not having false positives, but this changed to also warn on "potential" problems (like portability issues) and depend on configurations.\
The tool is continuously checked on real code - the debian linux distribution. they checks the packages from the source code using both the current and previous versions of cppCheck (regression testing). there are probably many false positives in the results, but there are also real bugs.

the tool started as a side project, but eventually became a commercial product under a new company. there are now two versions of the tool, the open source cppCheck and the commercial cppCheckPremium, with additional coding standards, extra checkers and customer support. the commercial version is certified by TUV.

one question from the audience is about the challenges, how to keep it compiler agnostic, how to handle template and macros (preprocessors).
</details>

### Beyond Compilation Databases to Support C++ Modules: Build Databases - Ben Boeckel

<details>
<summary>
Supporting C++20 Modules in build systems by using build databases.
</summary>

[Beyond Compilation Databases to Support C++ Modules: Build Databases](https://youtu.be/GUqs_CM7K_0?si=sl94_m0I0DyEcxnb), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Beyond_Compilation_Databases_to_Support_Cpp_Modules.pdf), [event](https://cppcon2024.sched.com/event/1gZg8/beyond-compilation-databases-to-support-c-modules-build-databases).

part of CMake development

#### What Are Compilation Databases?

compilation databases are simple json files with multiple objects in an array.

[clang specification](https://clang.llvm.org/docs/JSONCompilationDatabase.html)

> Each object describes a single command
>
> - Working directory
> - Input file
> - Output file (optional)
> - Arguments (list of strings) or command (single string, shell escaped)

they are generated by the build system, and usually work, but can have problems when using multiple configurations (target, release/debug) and can have some additional problems, like how the shell commands are escaped (difference for linux, windows cmd and powershell), and some other stuff.

#### How Do Modules Change The Status Quo?

> Modules complicate C++ compilation.\
> Basically inherit the Fortran 90 modules compilation model.
>
> - Importing a module requires files generated during compilation of another TU (the "BMI")
>   - BMI: built module interface, binary module interface
>   - Also "CMI" for "compiled module interface"
> - Similarities
>   - BMIs are compiler-specific
>   - Lookup based on in-source identification (filenames are meaningless)
> - Differences in details
>   - Fortran supports "submodules" and exporting multiple modules per TU
>   - C++ has "partitions" and flags need to agree between the BMI and importer

problems with the Fortran modules definitions...

working with the example, define library A, which use a file that provide a module, and some files which dont, and we compile the library as c++20.

```CMake
add_library(A)
target_sources(A
    PRIVATE
        a2.cpp a3.cpp
    PRIVATE
        FILE_SET CXX_MODULES
        FILES
        a1.cpp)
target_compile_features(A PUBLIC cxx_std_20)
```

CMake scans the files to build the dependencies graph, seeing which files provide a module and which files import it (and therefore must delay compilation).\
in the next example, we have three libraries: A, B and C, B links with A, and C links with B.

```CMake
add_library(A)
# add sources to A

add_library(B)
# add sources to B
target_link_libraries(B PRIVATE A)

add_library(C)
# add sources to C
target_link_libraries(C PRIVATE B)
```

we have some issues which prevent us from incorporating modules.

> What is missing? (for modules)
>
> - Ordering between commands
> - Information about module usage
>   - Currently CMake "smuggles" through module mapper files (basically response files)
>   - These files are referenced by but not necessarily present with the compilation database
> - Visibility of modules
>   - Just because we have A.mod doesn't mean anything can use it
>     - Might be private to its target
>     - Might not be linked by the target owning the source in question
> - Flag compatibility questions
>   - -std=c++26 in importer P and -std=c++23 in importer Q
>   - Different BMIs for different flags!

#### Build Databases

The idea is to cover the gaps, rather than just a list of objects, we have more data in the file and group it diffenfly.

> - Group commands into "sets"
> - Ordering and module usage
>   - Includes information on modules provided and required by the TU in question
> - Visibility of modules
>   - TUs are tagged with a flag to indicate whether it can be used outside of its target
> - Flag compatibility
>   - Sets belong to "families"
>   - Each instance of a family's set is a different flag compatibility view of the set (e.g., CMake configuration or importer-influenced flags)

the schema of the new file include the version and revision at the top level, and the sets. a set has a name and family name, a list of visible sets and the translation units.\
The translation unit objects has the required data about the object, source, work directory, build arguments and whether it provides or requires modules.\
The file is versioned, a major version change indicates adding fields in a way that changes the semantic meaning of the contents. a minor version indicates a change that doesn't affect a correct interpretation of the content.\
Sets are globally unique, they are mapped into build targets, Translation units define the way to build something.

This is still a work in progress, CMake 3.31 will have experimental support, and there is a suggesting for the ISO standardization.

```CMake
set(CMAKE_EXPERIMENTAL_EXPORT_BUILD_DATABASE 4bd552e2-b7fb-429a-ab23-c83ef53f3f13)

# Initialize the EXPORT_BUILD_DATABASE property on targets
set(CMAKE_EXPORT_BUILD_DATABASE 1)

find_package(WithModules)

add_library(A)
# add sources to A
target_link_libraries(A PRIVATE WithModules::WithModules)
```

describing the future work - tooling, IDEs, adding header unit support, better argument representation.

(audience questions)

</details>

### Why Is My Build So Slow? Compilation Profiling And Visualization - Samuel Privett

<details>
<summary>
Steps to identify build time and how to make them faster.
</summary>

[Why Is My Build So Slow? Compilation Profiling And Visualization](https://youtu.be/Oih3K-3eZ4Y?si=6OaOAxoW1Rp9UALu), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Why_is_my_Build_so_Slow.pdf), [event](https://cppcon2024.sched.com/event/1gZf4/why-is-my-build-so-slow-compilation-profiling-and-visualization).

#### Visualization

Visualizing compilation with Ninja and Clang and writing code the compiles faster, but not how to speed up compilers themselves.\
In general, compilation increases with the lines of code, and as programs are developed, it usually gets more complicates and has more code. in the real world, compilation times wane, they go up and go down, but the general trend is upwards. eventually it becomes a problem that build times break the productivity flows, and then there's a period of reducing build times, and the cycle repeats.

When Ninja builds the program, it outputs a "ninja_log" file. this file tracks start/stop times for each output files (with a hash of the command used to build it). we can take this file and use "perfetto" as an interactive trace viewer tool. it uses the chrome event tracing format (json). we can dump the file into the visualizer to see basic view of the build. if we want more information, we can use the clang `-ftime-trace` flag when building the software, and then we will get a detailed traces file for each source file.\
newer versions of the ninja_log files require conversion using "ninjaTracing" python tool, and can combine with the time traces files for a more complete view.

#### Single File Compilation

we start with a simple example.

```cpp
#include <iostream>

int main()
{
    std::cout << "Hello, world" << std::endl;
}
```

the preprocessor `#include <iostream>` copies around 69,000 lines of code into the file. we can also see this if we run only the preprocessor (`-E` flag in clang and gcc) and look at the processed file and check the size. we can then compile the file normally and see where the time was spent.

```sh
clang main.cpp -stdlib=libc++ -E &> prep_main.cpp
du -sh main.cpp # disk usage 4.0Kb
du -sh prep_main.cpp # disk usage 2.9Mb
clang main.cpp -stdlib=libc++ -ftime-trace
```

we see that most of the time is spent processing the included header file, so this is where we can gain speed ups.

> - Refactor massive header files
>   - Having smaller header files gives consumers a better chance at including only what's strictly necessary for them
> - Forward declarations
>   - Requires no external tooling
>   - Frowned upon for entities defined in another project
>   - Can obfuscate dependencies
> - Include What You Use
>   - [IWYU Project](https://github.com/include-what-you-use/include-what-you-use)

the next example is about using templates, in this example, we have recursive template and a specialization.

```cpp
#include <iostream>

template <int N>
struct Sum {
    static const int value = N + Sum<N - 1>::value;
};

template <>
struct Sum<0> {
    static const int value = 0;
};

int main()
{
    std::cout << Sum<5>::value << std::endl;
}
```

We can use the [cpp insights](https://cppinsights.io/) tool and see that we actually instantiate a template for each value of N from 5 to 0. if we had a massive number in the template, the compiler would have to create those lines, and then parse them. we want faster run-time, so we pay for it with compilation time.

> Templates - Recommendations
>
> - Consider whether you need to use Template Meta Programming
>   - <cpp>constexpr</cpp> and <cpp>consteval</cpp> can go a long way
>   - Is there a way to more directly express intent to the compiler?
> - Re-evaluate your API
>   - Do you need to be generic over that extra type?
>   - Can you eliminate recursion?

with layered template, one instantiation can cause many many templates to be created and bloat the file.

#### Project Level Compilation

we can use the "ClangBuildAnalyzer" tool together with trace files for a high level overview of hot spots:

> - Files that took longest to parse
> - Templates that took longest to instantiate
> - Functions that took longest to compile
> - Expensive headers (with include chains!)

(it also supports incremental builds)

translation units are the "final representation" of the source code before the compiler creates the abstract syntax tree. this happens after including other files, templates instantiations and inline functions.

```cpp
// templates.hpp
#include <iostream>
template <int N>
struct Sum {
    static const int value = N + Sum<N - 1>::value;
};

template <>
struct Sum<0> {
    static const int value = 0;
};
```

if we have three files that include this file and use the template, each translation unit needs to do the same thing, we pay for the actions again and again for each file.

```cpp
// some other file: a.cpp, b.cpp, c.pp
#include "templates.hpp"

Sum<8192>::value;
```

we aren't protected by `#pragma once` or header guards, they exist to prevent multiple definitions, but they happen after the expansion.\
C++20 modules can help eliminate redundant parsing. this also applies for pre-compiled headers. the headers are compiled once and reused for the next file.

If we look at this in insulation, it's sometimes faster to do the duplicate work in parallel threads than wait for the pre-compiled header to complete. but in many cases, the free threads can be used for other work. so it depends on our project structure.

> Poor Dependency Management
>
> - Builds should be purely functional
>   - Single-core and parallel builds should just work
> - Avoid large dependency bottlenecks
>   - Can force the build to be synchronous
>   - Especially important when generating code (i.e., Protobuf)
> - Prefer smaller targets
> - Explicitly expressed dependencies enables efficient hardware utilization while maintaining build correctness

There is also an effect for passing Include headers to the compiler, if we have too many files in the `-I` flag, the effect can be non-linear, interacting with the filesystem isn't free, and caches can max out and that causes performance issues.

#### Project Level Analysis

Perfetto allows us to query the build times using SQL.

> "How much of the build is spent including this specific header?"

```SQL
WITH our_headers AS (
SELECT DISTINCT arg_set_id, display_value
FROM args
WHERE KEY = 'args.detail'
AND (display_value LIKE 'my_header.h')
) SELECT SUM (slice.dur) / 1e+9 AS duration_sec, COUNT (*) AS occurrence_count
FROM our_headers
JOIN slice
ON our_headers.arg_set_id = slice.arg_set_id
```

the proposed flow is to first identify the expensive parts of the code, and then make a decision. it might be possible to remove the code entirely, and if not, then to refactor it, and if that's not possible, perhaps move it to a module. it's also entirely possible that the code is simply expensive to build.\
optimizing build times is a lot of work, so the focus should be on low-hanging fruit with high value. after the easy stuff is done, there are higher-order solutions.

> - Do less work
>   - Incremental build
>   - Use a package manager
> - Compiler caching
> - Throw hardware at it
> - Distributed builds

there are sayings from the 90's about the issue:

> - What Andy Giveth, Bill Taketh Away
>   - Andy and Bill's Law - "New software will always consume any increase in computing power that new hardware can provide"
> - Wirth's law - "Software is getting slower more rapidly than hardware is becoming faster"

audience question about unity builds, optimizing for compilation resources (memory usage instead of time), other suggestions like using the PIMPL idiom (pointer to implementation), abusing forward declarations.

</details>

### What's Eating My Ram? - Jianfei Pan

<details>
<summary>
Memory Usage.
</summary>

[What's Eating My Ram?](https://youtu.be/y6AN0ks2q0A?si=8T9qdRDSh__LtwrC), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/What%E2%80%99s_eating_my_RAM.pdf), [event](https://cppcon2024.sched.com/event/1gZdx/whats-eating-my-ram).

a story about running out of RAM, a stateless application that suffered from high memory usage. high memory usage requires more swaps (reducing performance), can lead to OOM error and the service going down, and the impact is also felt with other processes using the same resource.

the first thing to do is to find the top memory consumers with the shell command `top -o RES`, it's better to have a dashboard that shows the consumption over time. since the service is supposed to be stateless, we only expect the memory usage to increase with the scale of requests, but since the number of requests didn't change, it must mean the problem is in the code. we look at the memory allocator and the operating system, and we should recall how the virtual memory is divided.

- Kernel
- stack
- heap
- data and text

> - Arena: a structure that is shared among one or
more threads.
> - Heap: a contiguous region of memory subdivided into chunks.
> - Chunk: a range of memory of various sizes allocated to the application.

the two function calls are <cpp>mmap</cpp> and <cpp>sbrk</cpp>. we have overview of allocated and free chunks. freeing chunk of memory doesn't return it to the operating system, it remains allocated to the application. chunks are managed by bins according to their sizes.

> <cpp>malloc</cpp> algorithm:
>
> - If the appropriate bin has a chunk in it, use that
> - If no chunk is available, create a new chunk(<cpp>sbrk()</cpp>: extends the heap)
> - If the request is large enough (*M_MMAP_THRESHOLD*): <cpp>mmap</cpp> request memory directly from OS
>
> <cpp>free</cpp> algorithm:
>
> - Place the free chunk in the appropriate bin
> - If this chunk is adjacent to another free chunk, combine
> - If this chunk is mapped: munmap

there are two problems: memory leak and memory fragmentation.

#### Memory Leaks

A memory leak means that memory which is no longer needed isn't released, it can happen if we forget to free heap memory, or if we just continue adding data to a container. this can also happen if we forget to mark a virtual class with a virtual destructor or when there are circular references.\
there are tools for detecting leaked memory, we can inject a test allocator to our standard containers. this approach can be done selectively, but that means it requires manually changing the code and re-compiling. we could also use an address sanitizer (`-fsanitize=address`), which still requires re-compiling the code, but it has some memory overhead. there are also the valgrind tools, which don't require compilation, but are much slower than the other tools and even more memory intensive. each tool provides different results, some provide a final report, while other have more detailed data over time (snapshots).

> Memory Leaks Tips:
>
> - "Static" leaks may hide the real issue – we need enough traffic for profiling.
> - They are all good tools, but for different cases.
> - Catch the problem in earlier stages – Integrate AddressSanitizer in CI.
> - Install the tools so we can start profiling easily.
> - Care about the lifecycle and ownership of what we allocate

#### Fragmentation

> "You try to allocate a big block and you can't, even
though you appear to have enough memory free"

chunks are placed in a way that prevents reusing them after they are freed, so the heap needs more and more memory from the operating system. there is also internal fragmentation, in which we allocate large chunks but only use a small part of the data. we can estimate the ratio of external and internal fragmentation.

$$
External Fragmentation = 1 - LargestAlloctableBlock / TotalFreeMemory \\
Internal Fragmentation = 1 - AccessedBytes / TotalAllocatedBytes
$$

For External Fragmentation, we check what is the largest block we can allocate, we can use <cpp>mallinfo</cpp> (malloc info) to see the data. for Internal Fragmentation, we check how much of the allocated data is actually is used. we get the data from the valgrind tool with the `--tool=dhat` flag.

we want to reduce the fragmentation, or de-fragment the memory. there is no magic solution, at the hardware level, there is page-meshing, at the OS level there is the linux buddy system, and at the memory allocator level, we might pass some parameters to tune the performance, or use a different malloc implementation such as "jemalloc" and replace the allocator in our code. the best thing to do is to change the memory usage pattern in the application. local allocators can separate where long lived data is stored from the short lived data.

</details>

### Shared Libraries And Where To Find Them - Luis Caro Campos

<details>
<summary>
Shared Libraries and how are they used.
</summary>

[Shared Libraries And Where To Find Them](https://youtu.be/Ik3gR65oVsM?si=ago4Lx7jBGUA11nB), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Shared_Libraries_and_Where_To_Find_Them.pdf), [event](https://cppcon2024.sched.com/event/1gZez/shared-libraries-and-where-to-find-them)

> From the point of view of a C++ developer, compiled code typically ends up in either:
>
> - The executables themselves
> - Static libraries
> - Shared libraries
>
> Libraries are a vehicle for "reusable" code code that can be  invoked by other libraries or applications (even from other languages)

the lecture covers:

- WHEN shared libraries are needed
- WHICH programs needs to locate shared libraries
- TOOLS to query, troubleshoot
- WHERE to find shared libraries

we can separate the "when" into phases, first into "build" and "run" phases, and then further division of configuring the build, static linking, dynamic linking and packaging application. the linker matches the library during the build phase, and the runtime linker matches it when running the application.

```sh
g++ my_program.cpp -o my_program # fails
g++ my_program.cpp -o my_program -lfoo # link with library
g++ my_program.cpp -o my_program -lfoo -Xlinker --verbose # verbose to show the search path
```

the linker has a search procedure, determining where the linker searches for the libray. the linker path is stored into the executable. it also stores the libraries that the executable needs

```sh
file my_program
readelf -l my_program | grep interpreter
readelf -d my_program
objdump -p my_program | grep NEEDED

ldd my_program # will run the libraries
```

a common occurrence is that our library is not in a default location, this can happen when the library is local to the project, when we use a different version from the one installed in the file-system, or when it's vendored in.

```sh
g++ my_program.cpp -o my_program -lfoo # library isn't at the default search path
g++ my_program.cpp -o my_program -L/opt/foo/lib -lfoo # add library path to the linker
./my_program # can't find foo library
LD_LIBRARY_PATH=/opt/foo/lib  # change environment variable
LD_DEBUG=libs ./my_program 
ld.so --library_path /opt/foo/lib ./my_program # invoke the application with linker and library path
g++ my_program.cpp -o my_program -L/opt/foo/lib -lfoo -Wl,-rpath,/opt/foo/lib # embed the library location and the run path into the executable
readelf -d my_program
```

there is also the case of transitive dependencies. creating a shared library and using it for an executable, there are indirect dependencies to consider. `RUNPATH` is a newer linux variable, but it only affects direct dependencies, unlike `RPATH` (which was deprecated). there is also `LD_LIBRARY_PATH`, but it has lower priority, some confusion.

Apple OS platforms are similar to linux, but not identical. the linker is still called `ld`, but there are other tools for file inspection, and there is a different search path behavior.\
Windows are also different, there are two files, one for the linktime (.LIB extention, not the same as static libraries) and one for runtime (.DLL), some other differences from linux and MAC-OS.

#### Build Systems

> Build systems: Why locate libraries?\
> The linker already tells us if it can't find a library.
>
> - We could just do `-lopencv_core` and let the linker fail if the library can't be found.
> - But we may see things like the following in our build scripts:
>   - `find_package(OpenCV)`
>   - `PKG_CHECK_MODULES([OPENCV], [opencv >= 2.0])`

package descriptor files allow us to work with build systems, do additional checks (versions, architectures, platforms, etc...), and fail early in the build process rather than wait for the entire build and link to finish.

CMake has a "build tree" and an "install tree", but again, it's different for windows. package managers like conan and vcpkg can help. something about bundling.
</details>

### Secrets of C++ Scripting Bindings: Bridging Compile Time and Run Time - Jason Turner

<details>
<summary>
Creating A simple scripting engine.
</summary>

[Secrets of C++ Scripting Bindings: Bridging Compile Time and Run Time](https://youtu.be/Ny9-516Gh28?si=hCmAgmtOiqwBW76k), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Secrets_of_Cpp_Scripting_Bindings.pdf), [event](https://cppcon2024.sched.com/event/1gZf2/secrets-of-c++-scripting-bindings-bridging-compile-time-and-run-time).

SWIG - Simplified Wrapper Interface Generator, generates wrapper for multiple languages to integrate with C++. later expanded into ChaiScript.

other alternatives existed:

- <cpp>boost::python</cpp>
- Sol2
- Luabind

> Talk Goals:
>
> 1. You'll understand the challenges of bridging compile time <-> run time
> worlds
> 2. You'll understand one possible solution
> 3. You'll be able to pick up from the simple example and build your own
> simple scripting tool

#### ChaiScript

> ChaiScript:
>
> - Header-only scripting engine designed for embedding in C++
> - Automatic function / type deduction
> - Native script function <-> C++ Function interaction
> - Full support for exceptions
> - Just Works (TM)

[minimal example](https://godbolt.org/z/z33z3qqo)

```cpp
#include <chaiscript/chaiscript.hpp>

std::string helloWorld(const std::string &t_name) {
    return "Hello " + t_name + "!";
}
int main() {
    chaiscript::ChaiScript chai; // declare engine
    chai.add(chaiscript::fun(&helloWorld), "helloWorld"); // add function to script
    chai.eval(R"(puts(helloWorld("Bob"));)"); // evaluate script
} 
```

Since it was a header-only library, it was compiled with whatever flags the users had, which requires a lot of work to make it compatible with any configuration. it was also a project that led to learning a lot about the language and got the career started.

- Template Meta Programming
- Lambdas
- <cpp>constexpr</cpp>
- Static Analysis
- Compiler Warning Options
- Runtime Analysis
- Sanitizers
- Fuzz Testing

*cons_expr* is a later incarnation of the same idea, a scheme-inspired embedded scripting engine without dynamic allocations.

#### Building Our Own Scripting Engine

> Code Goals:
>
> 1. Simple
> 2. Succinct
> 3. Readable
> 4. Leverages the standard library
>
> Code Non-Goals
>
> 1. We are not worried about performance
> 2. We are not worried about compile-time
> 3. We are not worried about binary size

we have two parts that we need, the first is to register the function with the engine, the second is to call the registered function with runtime values.

##### Registering Functions

we want to be able to "add" various functions onto the engine, it needs to be able to handle different kinds of functions with different arguments and return types.

[starting example](https://godbolt.org/z/W8M7Wed7j)

```cpp
int add(int, int);
int abs(int);
void print(int);
void print(std::string_view);

struct ScriptEngine {
    void add(/**/); /// what here?
};
int main() {
    ScriptEngine se;
    se.add(&add, "add");
    se.add(&abs, "abs");
    se.add(&print, "print_int");
    se.add(&print, "print_string");
}
```

our function *that can match functions* must be a template, and it will use template pattern matching.\
we need to handle:

- free functions
  - including class static member functions
  - both <cpp>noexcept(true)</cpp> and <cpp>noexcept(false)</cpp>
- member function
  - cv qualified <cpp>const</cpp>, <cpp>volatile</cpp>, <cpp>const volatile</cpp> and nothing.
  - reference qualified `&`, `&&` (and nothing)
  - both <cpp>noexcept(true)</cpp> and <cpp>noexcept(false)</cpp>
- object with the call operator (`operator()`) overload

(note: during the talk it came up that we can template the `noexcept` parameter)

```cpp
template<typename Ret, typename ... Param, bool NoExcept>
void add(Ret (*func)(Param...), noexcept(NoExcept) ); // can match any free function
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...));
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...) const);
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...) volatile);
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...) const volatile);
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...) &);
template<typename Ret, typename Class, typename ... Param>
void add(Ret (Class::*f)(Param...) const &);
/// and so many more
template<typename Func> void add(Func &&func) {
    add(&Func::operator()); /// fallback that then registers overloaded call operator
}
```

static member functions are just free functions, same as <cpp>explict this</cpp> functions.
before C+11 variadic templates, we could either declare a template for every number of parameters or use macro (such as `BOOST_PP` for preprocessing).

```cpp
template<typename Ret>
void add(Ret (*f)());
template<typename Ret, typename P1>
void add(Ret (*f)(P1));
template<typename Ret, typename P1, typename P2>
void add(Ret (*f)(P1, P2));
template<typename Ret, typename P1, typename P2, typename P3>
void add(Ret (*f)(P1, P2, P3));
```

after we define the function signature, we need to store them somehow on the scripting engine. a generic way to store any possible type of callable. the first thing that comes to mind is templates, but we can't store a function template, and we can't get a pointer to a function template. for this talk, we will simplify the mental model and say that generic function templates don't exist, only the instantiations of such, and because of that, we can't take the address of the generic function, only of a concrete type.

```cpp
SomeType generic_function(std::span<SomeType>);
```

the "SomeType" should be able to hold any type, including void (which can't be stored), we can either use <cpp>std::any</cpp>, or a variant <cpp>std::variant\<std::monostate, int, std::string_view\></cpp>.

> - `std::any` - The approach taken by **ChaiScript**. It's magic and trivially easy for the user
> - `std::variant` - The approach taken by **cons_expr**. Requires more foreknowledge of the types wanted

for this talk, we will use the `std::any`, even if it can cause heap allocations (and therefore can't be compile-time complete).

```cpp
#include <any>
#include <functional>
#include <span>
#include <vector>

struct ScriptEngine {
std::vector<std::function<std::any(std::span<std::any>)>> functions;

template<typename Ret, typename ... Param>
    void add(Ret (*f)(Param...));
};
```

lets actually write the function, we have some assumptions about typing, no conversion support, no checking for number of arguments. this uses a lot of parameter expansions (the ellipses `...` syntax).

```cpp
template<typename Ret, typename ... Param>
void add(Ret (*f)(Param...)) {
    functions.push_back(
        [f](std::span<std::any> params) -> std::any { // stored lambda
            // a lambda that knows how to take the span of anys
            // and cast them to the desired types
            const auto invoker = // helper lambda
                [&]<std::size_t... Index>(std::index_sequence<Index...>) {
                    /// we need to unpack the parameter types and the indices together
                    /// this works because they have the same pack size
                    /// replace `any_cast` with your own helper that does
                    /// any conversions you want.
                    return func(std::any_cast<Param>(params[Index])...);
                };
            return invoker(std::make_index_sequence<sizeof...(Param)>());
        }
    );
}
```

we also need to store the name of the function somehow, so we use a map - [compiler explorer example](https://godbolt.org/z/frnrT7h76)

##### Invoking a Function

now that we 'stored' the function, we want to call it.

```cpp
#include <any>
#include <functional>
#include <map>
#include <span>
#include <string>
#include <array>

struct Function {
    std::function<std::any(std::span<std::any>)> callable;
    std::size_t arity;
};
std::map<std::string, Function> functions;
template <typename Ret, typename... Param>
void add(std::string name, Ret (*func)(Param...)) {
    // code from before
}

int main() {
    add("+", +[](int x, int y) { return x + y; }); // register a lambda
    std::array<std::any, 2> values{1, 2}; /// turn the arguments into an array
    return std::any_cast<int>(functions.at("+").callable(values)); /// call the stored function
}
```

#### Why Do This At All?

scripting funnel,

this code block

```cpp
int main() {
    add("-", +[](int x, int y) { return x - y; }); // register substraction
    add("*", +[](int x, int y) { return x * y; }); // register multiplication
    add("to_string", +[](int x) { return std::to_string(x); }); // to_string
    add("print", +[](int x) { return std::to_string(x); }); // printing - probably a bug
    std::vector<std::any> stack;
    stack.push_back(1);
    stack.push_back(3);
    stack.push_back(6);
    eval("*", stack); // evaluate 3 * 6 and push back to stack
    eval("-", stack); // evaluate 1 - 18 and push to stack
    eval("to_string", stack); // transform -17 to string
    eval("print", stack); // ???
    return std::any_cast<int>();
}
```

is the same as `print(to_string(1 - (3 * 6)))`, which we could have in a input file. this is our new goal now. there is some road block with function that return `void`, and there is problem with overload functions (name mangling!).\
function execution takes the function name, grabs the function arity (number of arguments), and returns them as a span from the stack and sends them to the function and removes them from the stack. if the result has a value, we push it back onto the stack.

[the final result on compiler explorer](https://compiler-explorer.com/z/dzofzoros) - scripting engine, parsing input and evaluation, [example with some enrichments and homework](https://compiler-explorer.com/z/86558dW56).

##### Overloading

if we want to handle overloads, we can either use 'arity-based' overloading, which is usually easier, but doesn't work well with our stack-based approach. the other option is using type-based overloading, which usually only happens in languages with static polymorphism (c++, D, java, C#). so the best approach is to avoid using overloading entirely.
</details>

### Mix Assertion, Logging, Unit Testing And Fuzzing: Build Safer Modern C++ Application - Xiaofan Sun

<details>
<summary>
The ZeroErr framework for testing and how it was developed.
</summary>

[Mix Assertion, Logging, Unit Testing And Fuzzing: Build Safer Modern C++ Application](https://youtu.be/otSPZyXqY_M?si=S_Qlk9kaUoMYRRyq), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Mix_Assertion_Logging_Unit_Testing_and_Fuzzing.pdf), [event](https://cppcon2024.sched.com/event/1gZe3/mix-assertion-logging-unit-testing-and-fuzzing-build-safer-modern-c++-application), [repository](https://github.com/sunxfancy/zeroerr).

[online demo](https://replit.com/@sunxfancy/ZeroErr-Demo)

#### The ZeroErr Framework

started as a tool for printing data, instead of overloading the <cpp>operator<<</cpp> for each type.

```cpp
// LOG(INFO) << Data;
// ASSERT(a > b, “A > B is not true”, a, b);
std::ostream& operator<<(std::ostream& out, const myStruct& data);
std::ostream& operator<<(std::ostream& out, my_ordered_map<std::string, myStruct> data);
std::ostream& operator<<(std::ostream& out, std::unique_ptr<myStruct> ptr);
std::ostream& operator<<(std::ostream& out, llvm::Value* data);
```

Having a framework for logging avoids polluting the namespace, helps with templating and extendability, and has customization points. both logging and assertions have similar behavior, an assertion usually has a logging built in to print the error statement.

a story about checking a cache functionality in a unit test, leads to us wanting to check the log itself for a specific output - we would want to check a side effect of the code (writing to a log). we could have also used an error code as a return value for each type of known error, but there are advantages for checking the log.

> - No need to change the API
> - No need to maintain the Error Code
> - Can check detailed information for a log message
> - Can capture additional context information if needed
> - Make sure specific path is taken

Fuzzing tests generate input for a function, usually string based arguments, this is used to detect bugs. *libfuzzer* is an example of a fuzzer.

```cpp
extern "C" int LLVMFuzzerTestOneInput(const uint8_t *Data, size_t Size) {
    DoSomethingInterestingWithMyAPI(Data, Size);
    return 0;
}
```

we can integrate the fuzzer directly into the framework to avoid additional CI-CD steps, and have the fuzzer take advantage of the other features in the framework.

> Motivation of ZeroErr
>
> - Providing a way to make logged data can be accessed in unit testing
> - No need to write print function for a compositional type (e.g. <cpp>std::map\<std::string, int></cpp>)
> - Allow user to write assertion for both in source code and unit testing code
> - Failure assertion can be logged
> - All features provided could be used in fuzzing

a header only library.

```cpp
#define ZEROERR_IMPLEMENTATION
#include "zeroerr.hpp"
int fib(int n) {
    REQUIRE(n >= 0, "n must be non-negative");
    REQUIRE(n < 20, "n must be less than 20");
    if (n <= 1) {
        return 1;
    }
    return fib(n - 1) + fib(n - 2);
}
TEST_CASE("fib function test") {
    CHECK(fib(0) == 1);
    CHECK(fib(1) == 1);
    CHECK(fib(2) == 2);
    CHECK_THROWS(fib(20));
}

TEST_CASE("log test") {
    // logging macros
    LOG("Basic log");
    WARN("Warning log");
    ERR("Error log");
    FATAL("Fatal log");
    LOG("log with basic type {} {} {} {}", 1, true, 1.0, "string");
    std::vector<std::tuple<int, float, std::string>> data = {
        {1, 1.0, "string"},
        {2, 2.0, "string"}
    };
    LOG("log with complex type: {data}", data);
    LOG_IF(1==1, "log if condition is true");
    LOG_FIRST(1==1, "log only at the first time condition is true");
    WARN_EVERY_(2, "log every 2 times");
    WARN_IF_EVERY_(2, 1==1, "log if condition is true every 2 times");
    DLOG(WARN_IF, 1==1, "debug log for WARN_IF");
}

// check inside logs
Expr* parseExpr(std::string input) {
    static std::map<std::string, Expr*> cache;
    if (cache.count(input) == 0) {
        Expr* expr = parse_the_input(input);
        cache[input] = expr;
        return expr;
    } else {
        LOG("CacheHit: input = {input}", input);
        return cache[input]->Clone();
    }
}

TEST_CASE("parsing test") {
    zeroerr::suspendLog();
    std::string log;
    Expr* e1 = parseExpr("1 + 2");
    log = LOG_GET(parseExpr, "CacheHit", input, std::string);
    CHECK(log == std::string{});
    Expr* e2 = parseExpr("1 + 2");
    log = zeroerr::LogStream::getDefault()
        .getLog<std::string>("parseExpr", "CacheHit", "input");
    CHECK(log == "1 + 2");
    zeroerr::resumeLog();
}

//  fuzzing
unsigned find_the_biggest(const std::vector<unsigned>& vec) {
    if (vec.empty()) {
        WARN("Empty vector, vec.size() = {size}", vec.size());
        return 0;
    }
    // implementation
}

FUZZ_TEST_CASE("fuzz_test") {
    FUZZ_FUNC([=](const std::vector<unsigned>& vec) {
        zeroerr::suspendLog();
        unsigned ans = find_the_biggest(vec);
        // verify the result
        for (unsigned i = 0; i < vec.size(); ++i) CHECK(ans >= vec[i]);
        if (vec.size() == 0) {
            CHECK(ans == 0);
            // verify WARN message to make sure the path is correct
            CHECK(LOG_GET(find_the_biggest,
            "Empty vector, vec.size() = {size}", size, size_t) == 0);
        }
        zeroerr::resumeLog();
        })
    .WithDomains(ContainerOf<std::vector<unsigned>>(InRange<unsigned>(0, 100)))
    .WithSeeds({{{0, 1, 2, 3, 4, 5}}, {{1, 8, 4, 2, 3}}})
    .Run(100);
}
```

Using Clang and libFuzzer with the sanitizer.

```sh
clang++ -std=c++11 -fsanitize=fuzzer-no-link \
-L='clang++ -print-runtime-dir' \
-lclang_rt.fuzzer_no_main-x86_64 \
-o test_fuzz test_fuzz.cpp
```

> What makes ZeroErr Different
>
> - Provide a cohesive solution for mixing assertion, logging, unit testing and fuzzing.
> - Logged data is structural and accessible
> - A structure-aware fuzzing API for quickly create fuzzing test cases as easy as
writing unit tests.

#### Pretty Printer

The printing component of the framework, it uses a template and decompositions to print complex types.

starting with the basic implementation for a <cpp>std::map</cpp> object, assuming the Key and Value types themselves are streamable.

```cpp
template <typename K, typename V>
std::ostream& operator<<(std::ostream& os, const std::map<K, V>& map) {
    os << "{";
    for (auto it = map.begin(); it != map.end(); ++it) {
        os << it->first << ": " << it->second;
        if (std::next(it) != map.end()) {
            os << ", ";
        }
    }
    os << "}";
    return os;
}
```

instead of having a specified template for each container, we can decompose the type

> - <cpp>is_integral\<T></cpp>
> - <cpp>is_streamable\<T></cpp>
> - <cpp>is_autoptr\<T></cpp> - `print(ptr.get());`
> - <cpp>is_container\<T></cpp> - `for(auto ele : container) print(ele);`
> - <cpp>is_tuple\<T></cpp> - `print(std::get<I>(tup));`

matching integral types with <cpp>std::enable_if_t</cpp>, matching containers because they have `.begin()` and `.end()` functions, creating custom type traits, using an explicitly declared to_string if the type already has it, some conflicts in rules.

```cpp
template <typename... Ts>
using void_t = void;

// containers
template <typename T, typename = void>
struct iterable : std::false_type {};

template <typename T>
struct iterable<T,
    void_t<decltype(std::declval<T>().begin()),
    decltype(std::declval<T>().end())>
> : std::true_type {};

// has explicit to_string method
template <typename T, typename = void>
struct contain_to_string : std::false_type {};

template <typename T>
struct contain_to_string<T,
    void_t<decltype(std::declval<T>().to_string())>
> : std::true_type {};

// what happens if both apply?
template <typename T>
typename std::enable_if<iterable<T>::value, std::ostream&>::type
operator<<(std::ostream& os, const T& ctn);

template <typename T>
typename std::enable_if<contain_to_string<T>::value, std::ostream&>::type
operator<<(std::ostream& os, const T& obj);
```

if two template overloads match, then there's a compilation error, they don't have a priority ranking. this can be done either with overloading or partial specializations.

#### Assertion Engine

Assertion has two behaviors, in user source code, it logs a failure and throws exception, but in testing code, it prints when fails, counts failure and can throw exception to stop the test case.

#### Logging API

structured logs - log level, timestamp, file path, and then the data. has a stringify method when objects aren't copyable. support for concurrent queues (not messing up the output file)

#### Fuzzing API

> - Domain is a set of all possible inputs for a data structure.
> - Corpus is the internal representation of a domain.
>
> Those two concepts are coming from google/fuzztest and autotest.

for example, text is a string, the domain is any possible string, but the corpus is an array of characters - including the non-printable characters, escape characters and null terminators.

</details>

### Import Cmake; // Mastering C++ Modules - Bill Hoffman

<details>
<summary>
Showing some issues with CMake and Modules.
</summary>

[Import Cmake; // Mastering C++ Modules](https://youtu.be/7WK42YSfE9s?si=ZLXXz4IWHYmR_ZJc), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/import_CMake_Mastering_Cpp_Modules.pdf), [event](https://cppcon2024.sched.com/event/1gZe6/import-cmake-mastering-c++-modules).

integrating CMake with C++20 Modules.

ChatGPT doesn't have good answers, because there isn't enough code that uses modules.

basic module file

```cpp
// B.cpp
export module B;
export void b() { }

// A.cpp
export module A;
import B;
export void a() {
    b();
}
```

when a compiler sees this code, it creates a BMI (build module interface) file, it has some different names and extensions depending on the compiler

- MSVC - `.ifc`
- GCC - `.gcm`
- Clang - `.pcm`

simply running the command `cl -std:c++20 -interface -c A.cpp` doesn't work, as it cannot find the B module. we first need to run a compile step on file B, this is not what we're used to having with `#include` statements.

> Chicken and the Egg
>
> - Modules require the build system to know which files produce which BMI files and which files consume them
> - Need to parse/compile file to find that out
> - So... We need to compile the code before we can compile the code...

CMake has support with fortran modules since 2005, but fortran modules aren't the same as c++ modules.

DAG - directed acyclic graph.

currently (2024), the GNinja and MSBuild generators have some support for building modules with cmake.

`cl -std:c++20 -scanDependencies A.json -c A.cpp`

```json
{
    "version": 1,
    "revision": 0,
    "rules": [
        {
            "primary-output": "A.obj",
            "outputs": [
                "A.json",
                "A.ifc"
            ],
            "provides": [
                {
                    "logical-name": "A",
                    "source-path": "c:\\users\\hoffman\\work\\cxxmodules\\cxx-modules-examples\\simple\\a.cpp"
                }
            ],
            "requires": [
                {
                    "logical-name": "B"
                }
            ]
        }
    ]
}
```

CMake `FILE_SET`, either `HEADERS` or `CXX_MODULES`. also `CXX_MODULES_BMI` when installing targets, and `CXX_SCAN_FOR_MODULES` to make cmake look for modules per file or not.

```cmake
cmake_minimum_required(VERSION 3.23)
project(simple CXX)
set(CMAKE_CXX_STANDARD 20)
add_library(simple)
target_sources(simple
    PRIVATE
    FILE_SET cxx_modules TYPE CXX_MODULES FILES
    A.cpp B.cpp
)
```

[are we modules yet?](https://arewemodulesyet.org/) - a tracker for how much of the popular C++ libraries support modules. this still less than 100 libraries out of a few thousands.

attempting to add a module for CTRE(compile time regular expression) library.

the source code.

```cpp
import std;
import ctre;

std::optional<std::string_view> extract_number(std::string_view s) noexcept {
    if (auto m = ctre::match<"[a-z]+([0-9]+)">(s)) {
        return m.get<1>().to_view();
    } else {
        return std::nullopt;
    }
}

int main() {
    auto opt = extract_number("hello123");
    if (opt) {
        std::string s(*opt);
        std::cout << s << "\n";
    }
    return 0;
}
```

the cmake file

```cmake
cmake_minimum_required(VERSION 3.29)
set(CMAKE_EXPERIMENTAL_CXX_IMPORT_STD "0e5b6991-d74f-4b3d-a41c-cf096e0b2508")
project(import_ctre)
find_package(ctre REQUIRED)
add_executable(ctre_hello ctre_hello.cpp)
target_link_libraries(ctre_hello PRIVATE ctre::ctre)
```

all sorts of weird issues, looking at how <cpp>import std;</cpp> will be supported. looking at how clang and MSVC do it, looking as cmake 3.30 and how it will support importing the standard library.

</details>

### Common Package Specification (Cps) In Practice: A Full Round Trip Implementation In Conan C++ Package Manager - Diego Rodriguez-Losada Gonzalez

<details>
<summary>
Using Common Package Specification files with Conan.
</summary>

[Common Package Specification (Cps) In Practice: A Full Round Trip Implementation In Conan C++ Package Manager](https://youtu.be/pFQHQEm98Ho?si=k8-Loecfyc0XKmzy), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Common_Package_Specification_In_Practice.pdf), [event](https://cppcon2024.sched.com/event/1gZew/common-package-specification-cps-in-practice-a-full-round-trip-implementation-in-conan-c++-package-manager).

Outline:

> - Introduction to Common Package Specification (CPS)
> - Creation of CPS files from existing Conan packages
> - Loading CPS files generated by build systems
> - Generating build system native files from CPS
> - Location of CPS files
> - Lessons learned and conclusions

combine a package with the included requirements, giving more information to the build tools and package management systems. CPS is intended to be standardized eventually, and they will look something like this:

```json
{
    "cps_version": "0.12.0",
    "name": "zlib",
    "version": "1.3.1",
    "configurations": ["release"],
    "default_components": ["zlib"],
    "components": {
        "zlib": {
            "type": "archive",
            "includes": ["@prefix@/include"],
            "location": "@prefix@/lib/libz.a"
        }
    }
}
```

demo of generating cps files, using conan.\
loading cps files into build systems, still experimental in CMake (requires setting flag).\
more demo.

</details>

### Implementing Reflection Using The New C++20 Tooling Opportunity: Modules - Maiko Steeman

<details>
<summary>
Reflection before C++26 by taking advantage of modules.
</summary>

[Implementing Reflection Using The New C++20 Tooling Opportunity: Modules](https://youtu.be/AAKA5ozAIiA?si=rlSdpqHh7ryi2xWq), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Implementing_Reflection_using_the_new_Cpp20_Tooling_Opportunity.pdf), [event](https://cppcon2024.sched.com/event/1gZf7/implementing-reflection-using-the-new-c++20-tooling-opportunity-modules), [repository](https://github.com/FireFlyForLife/NeatReflection).

reflection will come as a language feature in C++26, but we can still implement the same behavior in C++20 with external tools.

reflection - the ability to ask about the code, getting 'metadata' about it, such as knowing which members and methods are available.\
this can be usefull in many cases, such as serialization:

```cpp
json value_to_json(any any_value) {
    // Handle builtins
    if (any_value.type() == the_type<int>())
        return json{ any_value.value<int>() };
    if (any_value.type() == the_type<double>())
        return json{ any_value.value<double>() };
    if (any_value.type() == the_type<std::string>())
        return json{ any_value.value<std::string>() };
    // Recurse for classes
    if (any_value.is_class()) {
        return serialize_struct(any_value);
    }
}

json serialize_struct(any any_value) {
    json json;
    for (Field f : reflect_fields(any_value)) {
        json[f.name]["value"] = value_to_json(f.value());
        json[f.name]["type"] = f.type;
    }
    return json;
}
```

other usages:

> - Extension to the type system
> - WPF, Automatic Bindings
> - Language bindings: Python
> - Content editors
> - Automatic change detection

RTTI - Real Time Type Information.

what we need to do

- query a type for the fields
- query an object for each of the fields

the client API will look something like this:

```cpp
struct AnyRef {
    void* value;
    const Type* type;
};

json serialize_struct(AnyRef any_value) {
    json json;
    for (const Field& f : any_value.type->fields) {
        json[f.name]["value"] = f.value(any_value);
        json[f.name]["type"] = f.type;
    }
    return json;
}
```

we will also have a type registry, an in-memory storage for all types, each type will have the metadata relating to it (fields and methods). we will use virtual inheritance.

```cpp
extern std::unordered_map<std::string, Type> type_registry;
template<typename T> void register_type(Field[], Method[]);

struct Field;
struct Method;
struct Type {
    std::string name;
    std::vector<Field*> fields;
    std::vector<Method*> methods;
};

class Field {
public:
    virtual ~Field() = default;
    virtual std::string_view name() = 0;
    virtual const Type* type() = 0;
    virtual AnyRef value(void* object) = 0;
};

class Method {
public:
    virtual ~Method() = default;
    virtual std::string_view name() = 0;
    virtual const Type* return_type() = 0;
    virtual std::span<const Type> parameter_types() = 0;
    virtual AnyRef invoke(void* object, std::span<void*> args) = 0;
};
```

today, without reflection, if we wish to store a type into the registry, we have an option to do this manually with a macro.

```cpp
struct MyStruct {
    MyStruct() {};
    void func(double) {};
    int data;
};

RTTR_REGISTRATION {
    registration::class_<MyStruct>("MyStruct")
    .constructor<>()
    .property("data", &MyStruct::data)
    .method("func", &MyStruct::func);
}
```

there are also other options, which aren't as flexible.

- using <cpp>boost</cpp> for aggregated types - only fields (not member methods). doesn't handle type invariants.
- code parsers - external tools to read the source code, very complex and hard to maintain, requires a build system integration. the Unreal game engine and the <cpp>QT</cpp> library use this.
- existing compilers frontend - <cpp>LibClang</cpp>, <cpp>ClangTooling</cpp> - exiting tools, but also require integrating with the build system.

we don't want those, the compiler already knows everything, so why do we need to jump through all those hoops?\
we can look at the c++ compilation process - header files contain type info, which is already parsed and inserted into the source files. the types are now known, but are in the form of object files. we would like to take the type information before it's compiled and use it for writing source code.\
C++20 modules replace header file with modules, which are stored as module interfaces, which are then compiled into binary module interfaces. there is no way to 'plug-in' commands in the middle of the process (which is where we would like to put our generator), but we can create an artificial boundary with a new 'project' that depends on the original data. it uses the original project as a dependency.

(NOTE: we can actually override the compiler and have it direct output to a generator program in the middle of compilation)

we can look at the structure of a ".ifc" file - the MSVC format for modules.

- file signature
- header
  - offset to the table of contents
- partitions 1.. n
  - type indices and qualifiers
- string table
- table of contents
  - mapping to the partitions

limitations to this model of reflection.

- modules only
- no user attributes
- BMI filename query
- templates aren't instantiated - additional parsing required
- compiler specific

#### How Can This Library Be Used

the library is meant to be used in different stages of the work, such as field abstraction in a type registry. using a "pointer to member", a relative pointer, somehow de-virtualizing functions?\
*(I don't understand this)*

replacing <cpp>typeid()</cpp> with a faster alternative, using <cpp>std::atomic</cpp> and local static variables to generate a consistent unique identifier for each type.\
having an array of base class objects without slicing.

</details>

### C++/Rust Interop: Using Bridges In Practice - Tyler Weaver

<details>
<summary>
Connect C++ and Rust with "bridging" libraries.
</summary>

[C++/Rust Interop: Using Bridges In Practice](https://youtu.be/RccCeMsXW0Q?si=WM3uaXdaRzLO614o), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Cpp_Rust_Interop.pdf), [event](https://cppcon2024.sched.com/event/1gZfi/c++rust-interop-using-bridges-in-practice).

#### Connecting Rust

connecting existing C++ codebase to work with new Rust components, for example, a library with multiple plugin, most of the algorithms were written in C++, but wanting to also connect with Rust plugins.

- c++ headers
- c++ source (used `extern "C"` from rust)
- unsafe Rust (C abi)
- safe rust

starting a "Joint" type, which exists in both c++ and Rust.\
it is a move-only type, since we explicitly declared the constructors for move but not for copying.

```cpp
class Joint {
public:
    Joint() noexcept;
    ~Joint() noexcept = default;
    Joint(Joint&& other) noexcept = default;
    Joint& operator=(Joint&& other) noexcept = default;
};
```

```rs
pub struct Joint {
    name: String,
    parent_link_to_joint_origin: Isometry3<f64>,
}

impl Joint {
    pub fn new() -> Self;
}
```

The unsafe rust, allocating and deallocating, declaring functions that call the rust "constructor" and "destructor". in rust, there is no UB for creating objects, so the allocating isn't `unsafe`. we also tell the rust compiler to not mangle the names, and to use C calling conventions.

```rs
#[no_mangle]
extern "C" fn robot_joint_new() -> *mut Joint {
    Box::into_raw(Box::new(Joint::new()))
}

#[no_mangle]
extern "C" fn robot_joint_free(joint: *mut Joint) {
    unsafe {
        drop(Box::from_raw(joint));
    }
}
```

back to the C++ code, we add a private member of a unique pointer to store the address of the object. we pull the destructor function from the rust code.

```cpp
namespace robot_joint {

class Joint {
public:
    Joint() noexcept;
    ~Joint() noexcept = default;
    Joint(Joint&& other) noexcept = default;
    Joint& operator=(Joint&& other) noexcept = default;
    private:
    std::unique_ptr<rust::Joint, deleter_from_fn<robot_joint_free>> robot_joint_;
};

} // namespace robot_joint

namespace robot_joint::rust {

struct Joint;

} // namespace robot_joint::rust

extern "C" {
    extern void robot_joint_free(robot_joint::rust::Joint*);
}

template<auto fn>
struct deleter_from_fn {
    template<typename T>
    constexpr void operator()(T* arg) const {
        fn(arg);
    }
};

extern "C" {
    extern robot_joint::rust::Joint* robot_joint_new();
}
namespace robot_joint {

Joint::Joint() noexcept : robot_joint_{robot_joint_new()} {}

} // namespace robot_joint
```

all of that stuff so far can be done by code generators, but there still isn't support for library code. for example, robotics code uses matrix calculation from the *Eigen* library. we want to glue a library type from C++ to a library type in Rust.

```cpp
class Joint {
public:
    Eigen::Isometry3d calculate_transform(const Eigen::VectorXd& variables);
};
```

our c++ code should call the rust code.

```rs
impl Joint {
    pub fn calculate_transform(&self, variables: &[f64]) -> Isometry3<f64>;
}
```

(luckily for us, the default is column-wise matrix), something about passing C arrays as parameters between C++ and Rust - they naturally decay into pointers, but if we store them inside a type, then it maintains the size.

```rs
use std::ffi::{c_double, c_uint};
#[repr(C)]
struct Mat4d {
    data: [c_double; 16], // array doubles
}

#[no_mangle]
extern "C" fn robot_joint_calculate_transform(
    joint: *const Joint,
    variables: *const c_double,
    size: c_uint,
    ) -> Mat4d {
    unsafe {
        let joint = joint.as_ref().expect("Invalid pointer to Joint"); // convert to reference, "borrow"
        let variables = std::slice::from_raw_parts(variables, size as usize); // create slice
        let transform = joint.calculate_transform(variables); // actual work using the rust joint object function
        Mat4d {
            data: transform.to_matrix().as_slice().try_into().unwrap(), // type cast into array of c_doubles
        }
    }
}
```

and the C++ calling side, we define a type with a c-array member to avoid pointer decay. we pass the raw pointer to the matrix data and the number of elements, and we get back a struct with an array of doubles, which we then use to construct back the eigen type.

```c++
struct Mat4d {
    double data[16];
};

extern "C" {
    extern struct Mat4d robot_joint_calculate_transform(const robot_joint::rust::Joint*, const double*, unsigned int);
}

namespace robot_joint {

Eigen::Isometry3d Joint::calculate_transform(const Eigen::VectorXd& variables) {
    const auto rust_isometry = robot_joint_calculate_transform(robot_joint_.get(), variables.data(), variables.size()); // call rust code through the pointer
    Eigen::Isometry3d transform;
    transform.matrix() = Eigen::Map<Eigen::Matrix4d>(rust_isometry.data); // transform raw data into eigen type
    return transform;
}

} // namespace robot_joint
```

> Manual Interop
>
> - Create unsafe Rust functions for creating and deleting Rust objects
> - Store the Rust object in a <cpp>std::unique_ptr</cpp>
> - Move only C++ type containing <cpp>std::unique_ptr</cpp> of Rust type with methods
> - Fixed sized arrays wrapped in structs can be used in FFI
> - Use method implementations to bridge library types

#### Exposing a CMake Target

making the library available for others to use by integrating with the build system.

using [corrosion](https://github.com/corrosion-rs/corrosion) (formerly "cmake-cargo") to integrate rust into existing cmake.

```cmake
include(FetchContent)

FetchContent_Declare(
    bridge
    GIT_REPOSITORY https://github.com/you/bridge
    GIT_TAG main
    SOURCE_SUBDIR "crates/bridge-cpp")

FetchContent_MakeAvailable(bridge)
target_link_libraries(mytarget PRIVATE bridge::bridge)
```

in rust the file folder layout is standardized so the build system simply works out of the box. the rust package is called a "crate".

- Cargo.toml
- README.md
- crates
  - bridge
    - Cargo.toml
    - src
      - lib.rs
  - bridge-cpp
    - Cargo.toml
    - CMakeLists.txt
    - cmake
      - bridgeConfig.cmake.in
    - include
      - bridge.hpp
    - src
      - lib.cpp
      - lib.rs
    - tests
      - CMakeLists.txt
      - tests.cpp

the "Cargo.toml" file in the root folder.

```toml
[workspace]
members = ["crates/bridge", "crates/bridge-cpp"]

[workspace.package]
version = "0.1.0"
edition = "2021
```

the "crates/bridge-cpp/Cargo.toml" file. we manually name the library and build it as a static library, and we define the dependency to a local path.

```toml
[package]
name = "bridge-cpp"
version.workspace = true
edition.workspace = true

[lib]
name = "bridge_unsafe"
crate-type = ["staticlib"]

[dependencies]
bridge = { path = "../bridge" }
```

the "crates/bridge-cpp/CMakeLists.txt" file. building it a static library and making the `install` work for end users that want to use the cpp library in their project.

```cmake
cmake_minimum_required(VERSION 3.16)
project(bridge VERSION 0.1.0)
find_package(Eigen3 REQUIRED)
include(FetchContent)

FetchContent_Declare(
    Corrosion
    GIT_REPOSITORY https://github.com/corrosion-rs/corrosion.git
    GIT_TAG v0.4
)

FetchContent_MakeAvailable(Corrosion)

corrosion_import_crate(
    MANIFEST_PATH Cargo.toml
    CRATES bridge-cpp
)

add_library(bridge STATIC src/lib.cpp)

target_include_directories(
    bridge PUBLIC $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>$<INSTALL_INTERFACE:include>
)

target_link_libraries(bridge PUBLIC Eigen3::Eigen)
target_link_libraries(bridge PRIVATE bridge_unsafe)
set_property(
    TARGET bridge 
    PROPERTY CXX_STANDARD 20
)
set_property(
    TARGET bridge
    PROPERTY POSITION_INDEPENDENT_CODE ON
)

include(CMakePackageConfigHelpers)
include(GNUInstallDirs)

install(
    TARGETS bridge bridge_unsafe
    EXPORT ${PROJECT_NAME}Targets
    RUNTIME DESTINATION ${CMAKE_INSTALL_BINDIR}
    LIBRARY DESTINATION ${CMAKE_INSTALL_LIBDIR}
    ARCHIVE DESTINATION ${CMAKE_INSTALL_LIBDIR}
)

install(
    EXPORT ${PROJECT_NAME}Targets
    NAMESPACE bridge::
    DESTINATION "${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}"
)

configure_package_config_file(
    cmake/bridgeConfig.cmake.in "${PROJECT_BINARY_DIR}/${PROJECT_NAME}Config.cmake"
    INSTALL_DESTINATION "${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}"
)

install(
    FILES "${PROJECT_BINARY_DIR}/${PROJECT_NAME}Config.cmake"
    DESTINATION "${CMAKE_INSTALL_LIBDIR}/cmake/${PROJECT_NAME}"
)

install(
    FILES include/bridge.hpp
    DESTINATION ${CMAKE_INSTALL_INCLUDEDIR}
)
```

and the "crates/bridge-cpp/cmake/bridgeConfig.cmake.in" file

```CMake
@PACKAGE_INIT@

include(CMakeFindDependencyMacro)
find_dependency(Eigen3)

include("${CMAKE_CURRENT_LIST_DIR}/@PROJECT_NAME@Targets.cmake")
```

#### Using C++ Libraries in Rust

calling C++ code from rust. doing the reverse of what we did earlier.

"cargo.toml" file.

```toml
[dependencies]
cxx = "1.0"

[build-dependencies]
cxx-build = "1.0"
anyhow = "1.0.79"
git2 = "0.18.2"
conan2 = "0.1"
cmake = "0.1"
```

the "build.rs" file is part of the build, we can put anything non-standard here, it integrates with conan, github, etc...\
the script runs only if the "src" folder or the script itself has changed. otherwise, it doesn't do anything. it expects to have some environment variables set for it. it creates a tool chain file and passing it to cmake.

```rs
use conan2::ConanInstall;
use std::path::{Path, PathBuf};

fn main() -> anyhow::Result<()> {
    println!("cargo:rerun-if-changed=build.rs");
    println!("cargo:rerun-if-changed=src");

    let out_dir: PathBuf = std::env::var_os("OUT_DIR")
        .expect("OUT_DIR environment variable must be set")
        .into();

    let data_tamer_url = "https://github.com/PickNikRobotics/data_tamer";
    let data_tamer_source = out_dir.join(Path::new("data_tamer"));
    if !data_tamer_source.exists() {
        git2::Repository::clone(data_tamer_url, data_tamer_source.as_path())?;
    }

    let data_tamer_cpp = out_dir.join(Path::new("data_tamer/data_tamer_cpp"));

    let conan_instructions = ConanInstall::with_recipe(&data_tamer_cpp)
        .build("missing")
        .run()
        .parse();

    let conan_includes = conan_instructions.include_paths();
    let toolchain_file = out_dir.join(Path::new("build/Debug/generators/conan_toolchain.cmake"));
    let data_tamer_install = cmake::Config::new(&data_tamer_cpp)
        .define("CMAKE_TOOLCHAIN_FILE", toolchain_file)
        .build();

    let data_tamer_lib_path = data_tamer_install.join(Path::new("lib"));
    let data_tamer_include_path = data_tamer_install.join(Path::new("include"));
    
    cxx_build::bridge("src/main.rs")
        .includes(conan_includes)
        .include(data_tamer_include_path)
        .include("src")
        .std("c++17")
        .compile("demo");
    
    println!("cargo:rustc-link-search=native={}", data_tamer_lib_path.display());
    println!("cargo:rustc-link-lib=static=data_tamer");

    conan_instructions.emit();
    Ok(())
}
```

the c++ "shim" header file with a single free function that calls <cpp>std::make_unique</cpp>.

```cpp
#pragma once
#include <memory>

namespace DataTamer {

template <typename T, typename... Args>
std::unique_ptr<T> construct_unique(Args... args) {
    return std::make_unique<T>(args...);
}

}
```

using the c++ library from the rust code, telling the compiler to bridge with a c++ namespace, what to include and matching c++ functions with their rust equivalents.

```rs
#[cxx::bridge(namespace = "DataTamer")]
mod data_tamer {
    unsafe extern "C++" {
        include!("shim.hpp");
        include!("data_tamer/data_tamer.hpp");

        type ChannelsRegistry;

        #[rust_name = "channels_registry_new"]
        fn construct_unique() -> UniquePtr<ChannelsRegistry>;
    }
}

fn main() {
    let mut registry = data_tamer::channels_registry_new();
}
```

</details>
