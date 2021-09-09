# CPPNow 2021

[CppNow 2021 conference lectures](https://youtube.com/playlist?list=PL_AKIMJc4roXvFWuYzTL7Xe7j4qukOXPq)

## Keynote: CMake: One Tool To Build Them All - Bill Hoffman [ CppNow 2021 ]

<details>
<summary>
Where CMake came from and where is it going.
</summary>

[CMake: One Tool To Build Them All](https://youtu.be/wULu83jQmIQ)

> Overview
>
> - Kitware,open Source and how CMake came to be
> - A high-level tour of what CMake has to offer
> - C++ Modules
> - how to Learn CMake
> - Packaging C++

introducing the company he works for,Kitware. they work with the private sector, academia and governments. they do visualizations, high-performance, computer vision, video analysis, etc... they also do a lot of open source. they have courses for CMake, features, developing an auditing build systems, they worked with the **MineCraft** team, and now even visual studio supports CMake.

CMake was started in 2001, as an offshoot from a project of the national library of medicine which had tons of images. it began as a toolkit for cross platform building. Like how boost aims to give c++ a set of useful libraries, CMake aims to give c++ a solution for compatibility and portability.

> - Same build too and files for all platforms
> - Easy to mix both large and small libraries
> - Only depend on a C++ compiler
> - Let developers use th IDE and the tools the are most familiar with

[Professional CMake - book](https://crascit.com/professional-cmake/)

CUDA is now first class language in CMake, with native support.

### Quick CMake Tour

> "make complicated things easy, so you don't have to have an expert on the team"

simple commands for complicated work!

> - add_library()
> - add_executable()
> - add_test()

CMake workflow

> - run cmake
> - run cmake --build
> - run ctest

there is a gui-version, interactive command line interface, and a non interactive command line.

//a diagram

CMakeCache.txt holds all sort of stuff.

ninja is a command line tool by google, that's also supported with CMake.

modern CMake is target-centric. not difference between internal and external targets. the whole point of CMake is that we describe the dependencies and it is then build accordingly.

#### Usage Requirements:

PRIVATE: means only the target use  
INTERFACE: means only consuming target will use  
PUBLIC: private + Interface  
\$\<BUILD_INTERFACE>:  
\$\<INSTALL_INTERFACE>:

this changes how what the call to the compiler uses as arguments

```CMake
target_link_libraries(trunk PUBLIC root)
target_link_libraries(leaf PUBLIC trunk)
```

will result in

```bash
/usr/bin/c++ -fPIC -shared -Wl, -soname, libleaf.so
-o libleaf.so leaf.cxx.o libtrunk.so libroot.so
```

and making root private for trunk

```CMake
target_link_libraries(trunk PRIVATE root)
target_link_libraries(leaf PUBLIC trunk)
```

will result in it not being part of the compile command for the 'leaf'

```bash
/usr/bin/c++ -fPIC -shared -Wl, -soname, libleaf.so
-o libleaf.so leaf.cxx.o libtrunk.so
```

we can propagate dependencies with as TLL (_target link libraries_)

```cmake
target_include_directories
target_compile_definitions
target_compile_options
target_sources
target_link_options
```

there is jumbo build/ Unity which does grouping

#### Presets

> Allow common configuration flags(variables, build directory, generator, etc...) for a project to be stored in a JSON file for reuse
>
> - CMakePresets.json - version controlled, for sharing between users
> - CMakeUserPresets.json - not version controlled, for local machine-specific or user-specific use

example of a preset.

#### Pre-Compiled Headers

CMake natively support pre-compiled headers for compilation speed up instead of repeatedly parsing header files

```cmake
add_library(leaf SHARED leaf.cxx)

target_precompile_headers(leaf
    PRIVATE
        <iostream>
        <vector>
        <unordered_map>
    INTERFACE
        "leaf.h")
```

support for multi config, build both release and debug.
ccmake has colors now.
integrate runtime dependencies with install target.

#### Full Cross Platform Install

> Specify rules to run at install time
> Can install target, files or directories
> Provides default install locations

```cmake
add_library(leaf SHARED leaf.cxx)
install(TARGETS root trunk leaf parasite)
```

#### CPack - Packaging Software

packaging the cmake installer with CPack, which creates installers for all sorts of platforms. once we get 'make install' to work, should be easier.

#### Testing

CMake supports testing, we need to either call '_include(CTest)_' or '_enable_testing()_' to get it running.

```cmake
add_test(NAME testName Command exeName arg1 arg2)
```

executable that returns 0 for success.

we then have an executable '_ctest_' that runs the tests. we can run it from the build directory.
options:

> - -j - parallel mode
> - -R - choose test
> - -vv - verbose
> - --rerun-failed - repeat failed tests
> - --help - get help

now googleTest is also integrated, with _gtest_discover_test_ that finds new test without re-running cmake.

```cmake
include(GoogleTest)
add_executable(tests tests.cpp)
target_link_libraries(test GTest::GTest)
gtest_discover_test(tests)
```

multi core test with processor Affinity

```cmake
set_test_properties(myTest PROPERTIES
    PROCESSOR_AFFINITY ON
    PROCESSORS 4)
```

#### CDash

a web based tool that is a dashboard for the build system, dynamic analysis, works with sanitizers. part of the integration cycle, with source code control, databases.

### C++20 Modules

c++20 now has modules.

```cpp
//helloworld.cpp
export module helloworld;
import <iostream>;
export void hello()
{
    std::cout<<"Hello World\n";
}

//main.cpp
import helloworld;
int main()
{
    hello();
}
```

and if we compile it out of order we get an error.

```bash
CC -o main.cpp
CC -o helloworld.cpp
```

CMake already knows how to deal with Modules, both internally and for the Fortran language. now ninja also works with modules. a huge diagram of how the build graph for fortran looks. a different graph for c++ modules build flow.

there is the issue of scanning and collating the sources, all together, one-by-one, a combination of scanning and collating.

the compilers for c++ don't yet support modules perfectly, so we need to wait and see.

### Learning CMake

- don't copy old CMake code - the syntax changes
- look at 'modern cmake' talks.
- read the "professional cmake" book
- look at tutorials at cmake.org
- check the documentation
- 'Mastering cmake' book is now open source - is constantly updated with modern examples

#### CMake Find Modules

we can find modules on the system, without having to always build it as part of our project

```cmake
<!-- find the png library -->
find_package(PNG REQUIRED)
add_library(trunk SHARED trunk.cxx)
target_link_libraries(trunk PRIVATE PNG::PNG)
```

#### Exporting targets

install rules can generate imported targets. install the library and sets the target import rules.

```cmake
add library(parasite STATIC eat_leaf.cxx)
install(TARGETS parasite root trunk leaf EXPORT tree-targets)
install(EXPORT tree-targets DESTINATION lib/cmake/tree)
```

the conan package manager can create cmake config.cmake files.

support for external projects to reference projects at build time. cloning the project, build the project, and use it as dependency for the current project.

```cmake
ExternalProject_Add(foo
GIT_REPOSITORY git@github.com:FooCo/Foo.git
GIT_TAG origin/release/1.2.3
)

ExternalProject_Add(foo
GIT_REPOSITORY git@github.com:BarCo/Bar.git
GIT_TAG origin/release/2.3.4
DEPENDS foo
)
```

build time and not configure time.

if we want configure time, we can do fetch Content.

```cmake
FetchContent_Declare(catch
    GIT_REPOSITORY https://github.com/catchorg/Catch2.git
    GIT_TAG v2.2.1
)

FetchContent_GetProperties(catch)
if (NOT catch_POPULATED)
    FetchContent_Populate(catch)
    add_subdirectory(${catch_SOURCE_DIR} ${catch_BINARY_DIR})
endif()
```

#### Package Managers

we still need package managers like conan,vcpkg or Spack, this is epically true for multi-language code and very big projects.

a photo showing the clusters of dependencies in some package manager. a page about how spack helped with building a project that combines c++ and python.

### The CMake Future

wishlist

> - All C++ compilers provide build system interfaces to collect c++20 modules dependencies information.
> - A cross platform standard for the information found in cmake config files.

### Questions From the Chat

- integration with cmake and conan.
- when to move from FetchContent to a package manager.
- The easiest way to extract CMake properties for use in other places.
- recommendations for large code base with wrappers for other stuff.
- do and don'ts for the tree structure.
- are there plans to support 'Bazel'.
- plans to support libraries with other meta-build system.
- a converter for vcproj to cmake.
- ninja over make and why?
- add support for multi-builds in parallels

</details>

## The C++ Rvalue Lifetime Disaster - Arno Schödl - [CppNow 2021]

<details>
<summary>
Using rvalues and lifetime extension can fail us. member accessor for rvalue objects shouldn't return a reference.
const & should never have bound to rvalues.
</summary>

[The C++ Rvalue Lifetime Disaster](https://youtu.be/sb7cj-3l1Kc)

the use of rvalue references and move semantics. replace copying with moves when possible tp avoid memory operations.

also used to manage lifetime, as well as for c++20 ranges

```cpp
auto rng=std::vector<int>{1,2,3} | std::views::filter([](int i){return i%2==0;}); //doesn't compile
```

this doesn't compile for rvalue

### Pitfalls

can't move from a const value, and moving will mess with NRVO (names return value optimization) and make it harder for the compiler to elide the construction.

```cpp
A foo()
{
    const A a;
    return std::move(a); //error!
}
A foo2()
{
    A a;
    return std::move(a); //works, but we are messing with RVO
}
A foo3()
{
    const A a; //doesn't matter if we're const or not, elision works
    return a;
}
```

but if we have two possible values, we can't do NRVO, and we also can't do move (because of const).

```cpp
A foo4()
{
    if ()
    {
        const A a;
        return a;
    }
    else
    {
        const A a;
        return a;
    }
}
```

and here we can't do copy/move ellison, because it's member variable. we also can't do a move, members don't automatically become rvalues.

```cpp
struct B {
    A m_a;
};
A foo()
{
    B b;
    return b.m_a;
    //return std::move(b).m_a; //this would work.
}
```

recommendations:

> - make return variables non-const
> - use clang -Wmove flag

### Temporary Lifetime Extension

```cpp
struct A;
struct B {
private:
A m_a;
public:
const A& GetA() const &
{
    return m_a; //return by reference
}
};
B b;
const auto & a = b.getA();
struct C{
    A getA() const &; // return by value
};
C c;
const auto & a1 = c.getA();
const auto & a2 = B.getA();
```

if we capture something with const reference, it can extended the lifetime of the object it's capturing.

_std::min_ doesn't take rvalue-ness into consideration, it returns a lvalue reference. a will dangle.

```cpp
bool operator<(const A&, const A&);
struct C
{
    A getA() const&;
} ;
C c1,c2;
//...
const auto & a = std::min(c1.getA(),c2.getA()); //a will dangle
```

lets' have a min function that keeps rvalue references using perfect forwarding. but it still doesn't work

```cpp
namespace out
{
    template<typename Lhs,typename Rhs>
    decltype(auto) min(Lhs && lhs,Rhs && rhs)
    {
        return rhs<lhs ? std::forward<Rhs>(rhs)? std::forward<Lhs>(lhs);
    }
}
```

lifetime extension only works where there an object.

an example with forwarding a return and _'decltype(auto)'_

the advice is to stop using temporary life time extension,
what we want is :

> automatically declare variable
>
> - _auto_ if constructed from value or rvalue reference
> - _const auto &_ if constructed from lvalue reference

he suggest this macro code instead of lifetime extension.

```cpp
template<typename T>
struct decay_rvalues
{
    using type = std::decay_t<T>;
};
template <typename T>
struct decay_rvalue<T&>
{
    using type=T&;
};

#define auto_cref(var,...) \
typename decay_rvalue<decltype((__VA_ARGS__))>::type var = ( __VA_ARGS__)'

```

if we add parentheses it's bad, it will always return a reference.

```cpp
decltype(auto) foo()
{
    auto_cref (a, some_a()); // a = some(); with type deduced
    return a; //if we have parentheses, things will be different, it will be a reference.
}
```

theres a debate about whether the macro should return const or not (if not, it can get optimized in NRVO).

```cpp
struct A;
struct B {
    A m_a;
    A const & GetA() const
    {
        return m_a;
    }
}
auto_cref(a1, B().m_a); // B() is rvalue, so it's members are also rvalues;
auto_cref(a2, B().GetA()); // we have a const reference as the return type, so we get a dangling reference const A &;
```

now the problem is that our 'auto_cref' binds to everything, but should rvalues be converted to values?

```cpp
struct A;
A const & L(); //lvalue
A const && R(); //rvalue

decltype(false? L(): L()); // A const &
decltype(false? R(): R());// A const &&
decltype(false? R(): L());// A const, not reference. forces a copy.
```

c++20 has a new trait _common_reference_t_. which was invented for c++20 ranges,

```cpp
std::common_reference_t<A const &, A const &>; //A const &
std::common_reference_t<A const &&, A const &&>; // A const &&
std::common_reference_t<A const &, A const &&>; //A const &. lvalue reference
std::common_reference_t<A const, A const &>; //A. a value
```

so, std::common_reference embraces rvalue amnesia.

### Promises of defences

| Mutability | short Lifetime               | long Lifetime |
| ---------- | ---------------------------- | ------------- |
| immutable  | const &&                     | const &       |
| mutable    | && (can scavenge, move from) | &             |

currently, c++20 reference binding strengths lifetime promise.
from short to long, and from mutable to immutable.

what if we could go the reverse?

> - Allow binding only if promises get weaker
>   - less lifetime
>   - less mutability
>   - less 'scavenge-ability'
>
> * only lvalues should bind to _const &_
> * anything may bind to _const &&_

but we can't allow going from lvalue to rvalue.

### Ideas to Fix the Issue

some things that must hold true before any changes.

where are references used?

> - local/global variable declarations
> - structured binding
> - function/lambda parameter lists
> - members (initialized in PODs)
> - members (initialized in constructors)
> - lambda captures

how it would look with a pragma change. we would need a feature test macro, replace const & parameters with const &&. we will need to change std::common_reference.

</details>
