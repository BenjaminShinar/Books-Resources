<!--
// cSpell:ignore Qlibs fftw rtsan noundef dispatchv resultv Wfunction Wperf nonallocating Wunknown perfetto IWYU Wirth valgrind dhat jemalloc
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Tooling

<summary>
15 Talks about tooling and supports for C++.
</summary>

- [x] Beyond Compilation Databases To Support C++ Modules: Build Databases - Ben Boeckel
- [X] Building Cppcheck - What We Learned From 17 Years Of Development - Daniel Marjamäki
- [ ] C++/Rust Interop: Using Bridges In Practice - Tyler Weaver
- [ ] Common Package Specification (Cps) In Practice: A Full Round Trip Implementation In Conan C++ Package Manager - Diego Rodriguez-Losada Gonzalez
- [X] Compile-Time Validation - Alon Wolf
- [ ] Implementing Reflection Using The New C++20 Tooling Opportunity: Modules - Maiko Steeman
- [ ] Import Cmake; // Mastering C++ Modules - Bill Hoffman
- [X] Llvm's Realtime Safety Revolution: Tools For Modern Mission Critical Systems - Christopher Apple, David Trevelyan
- [ ] Mix Assertion, Logging, Unit Testing And Fuzzing: Build Safer Modern C++ Application - Xiaofan Sun
- [ ] Secrets Of C++ Scripting Bindings: Bridging Compile Time And Run Time - Jason Turner
- [ ] Shared Libraries And Where To Find Them - Luis Caro Campos
- [x] What's Eating My Ram? - Jianfei Pan
- [X] What's New For Visual Studio Code: Performance, Github Copilot, And Cmake Enhancements - Alexandra Kemper, Sinem Akinci
- [X] What's New In Visual Studio For C++ Developers - Michael Price & Mryam Girmay
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

problems with the Fortran modules defintions...

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

we aren't protected by `#pragma once` or header guards, they exist to prevent multiple definitons, but they happen after the expansion.\
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
