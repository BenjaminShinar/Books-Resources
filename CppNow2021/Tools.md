<!--
ignore these words in spell check for this file
// cSpell:ignore Bjarne Strostrup Bazel libcxx libstdc libc cppstd soname ccmake spack cstdio ipython cppm fimplicit fmodules fmodule clangmi pybind numpy pyplot asarray vcrt getptd tpxcptinfoptrs pxcptinfoptrs mtdll vcruntime
 -->

[Main](README.md)

Tools

## Keynote: CMake: One Tool To Build Them All - Bill Hoffman

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

## Code Analysis++ - Anastasia Kazakova

<details>
<summary>
Static analyzers, Tools to make our code better.
</summary>

[Code Analysis++](https://youtu.be/qUmG61aQyQE)

### Software Quality

not having bugs, readability, maintainability, extendability, scaleability

> - a trade-off between quality and cost of development.
> - external vs internal quality
>   - external - features, performance.
>   - internal - architecture.

Developers frustration points: \
What developes care about and worry about.
Style

look at this code, what does it do? it just constucrts an int from the number 42.

```cpp
template <class T, int ...X>
T pi (T(X...));
int main
{
    return pi<int,42>;
}
```

if we have 10 ways to do one thing in the language, then our code base might use all ten ways.
certification process.

Undefined Behavior

> - Data races.
> - Memory accesses outside of array bounds.
> - Signed integer overflow.
> - Null pointer dereference.
> - Access to an object through a pointer of a different type.
> - etc...

NDR - no diagnostic required - some code is illformed, but no warnings or errors.

> **"Compilers are not required to diagnose undefined behavior"**

### Code Analysis suggestions

improve software quality, lower develop frustration, avoid undefined behavior. \
getting help from the language, the lifetime safety suggestions for diagnostics with or without annotations. contracts, assertions (pre and post conditions),parameter passing semantics (in/in-out/out/move/forward). we do something in the code to help an external tool know what to look for.

| Language & Compiler                               | Stand-alone analyzer                       |
| ------------------------------------------------- | ------------------------------------------ |
| core tool - hard to update                        | side tool, any adopted by tht team is ok   |
| code base might require specific compiler version | no strong requirement for analyzer version |
| set of checks is defined by compiler vendor       | custom checks are possible                 |
| standard to everyone                              | depends on the tool                        |

### Tooling

software quality: how to

> pre-compilation stage
>
> - Refactoring
> - Pair programming
> - Static analysis
>
> post-compilation state
>
> - Static analysis
> - Unit testing
> - Dynamic analysis
> - Code review
> - Other Testing

static analysis can happen before compilation and after it.\
we can get some help from the compiler with flags

> - -Wall
> - -Wextra
> - -Wsign-compare
> - -Wsizeof-pointer-memeacess
> - -Wmisleading-indentation

comparision between using the compiler and an external tool

| Compiler checks                                | Stand-alone analyzer                 |
| ---------------------------------------------- | ------------------------------------ |
| Checks the code **after it's written**         | Check code **while writing** it      |
| Analysing the code with the proper fags / vars | Should use compilation flags & env   |
| Using specific compiler                        | Can get checks from other compilers  |
| Different compiler flags                       | Checks are independent from compiler |

lifetime safety

```cpp
std::string get_string();
void dangaling_string_view()
{
    std::string_view sv =get_string();
    auto c = sv.at(0);
}

void dangling_iterator()
{
    std::vector<int> v = {1,2,3};
    auto it = v.begin();
    *it = 0;
    v.push_back(4);
    *it = 0;
}
```

gsl suggest annotations for owner, pointers, etc...

> **gsl: guideline support library**

### Data Flow Analysis (DFA)

static analyzers can catch incoherent data flow, like in this example: \
this example uses multiple assignment with the comma operator, but the important thing is that the second if statemt is always true. static analyzers can find things like this

```cpp
enum class Color {Red, Blue, Green, Yellow};
void do_shadow_color(int shadow)
{
    Color cl1,cl2;
    if (shadow)
    cl1= Color::Red, cl2= Color::Blue;
    else
    cl1= Color::Green, cl2= Color::Yellow;
    if (cl1 == Color::Red || cl2 == Color::Yellow)
    {
        //... always executed
    }

}
```

and it can also detect code like this, where we dereference a deleted pointer

```cpp
void linked_list::process()
{
    for (node *pt = head; pt!= nullptr; pt= pt->next)
    {
        delete pt;
    }
}
```

we can also do global data flow analysis, rather than just in the scope of a function or a code block. like seeing that we deallocate inside a function but then dereference the pointer.

```cpp
static void delete_ptr(int* p)
{
    delete p;
}

int handle_pointer()
{
    int *ptr = new int;
    delete_ptr(ptr);
    *ptr = 1; // local variable may point to deallocated memory
    return 0;
}
```

it's quite hard to do global static analysis on the entire program, so it's mostly contained into translation unit. we distinguish between private entities (entire operations happen in the translation unit, only called from this unit), and 'unsafe entities', which involve multiple translation units.

we can use data flow analysis to identify

> Local issues:
>
> - Constant conditions.
> - Dead code.
> - Null dereference.
> - Dangling pointers.
> - Endless loops.
> - Infinite recursion.
> - Unused values.
> - Escape analysis (local memory being returned).
>
> Global issues (limited to translation unit):
>
> - Constant function result.
> - Constant function parameters.
> - Unreachable calls of function.

some parts of this have been included in CLion.

in the future there might be cross translation unit (CTU) analysis.

### Core Guidelines Issues

> "Within C++ is a smaller, simpler, safer language struggling to get out."
> --Bjarne Strostrup

we want the tools to enforce us to follow the guidelines, if it's possible. some guidelines are toolable, some aren't worth the work, some require changes to the language, and some are completely not toolable.

for example, the following two guidelines are fairly easy to identify and write enforcements for.

> Toolable guidelines:
>
> - F.16: "For "in" parameters, pass cheaply copied types by value and others by reference to const"
>   - E1: Parameter being _passed by value_ has a _size > 2\*sizeof(void\*)_ -> suggest reference to const.
>   - E2: Parameter being _passed by reference to const_ has a _size < 2\*sizeof(void\*)_ -> suggest passing by value.
>   - E3: Warn when a parameter _passed by reference to const_ is _moved_.
> - F.43: "Never (directly or indirectly) return a pointer or a reference to a local object"

however, other guidelines aren't so easy. even if we can identify them somehow, it's harder to decide what to do with them.

> Less toolable guidelines
>
> - F.1: ""Package" meaningfull operations as carefully names functions"
>   - Detect identical and _similar_ lambdas used in diffrent places.
> - F.2: "A function should preform a single logical operation"
>   - More than one 'out' parameter or more than six parameters are suspicious.
>   - Rule of one screen - 60 lines by 140 characters.
> - F.3: "Keep functions short and simple"
>   - Rule of one screen?
>   - Cyclomatic complexity of more than 10 logical paths.

it's hard to find duplicate code, there are some tools, but again, there are many ways to do the same things, and we would want the tool to identify them.

there a guidelines that might be possible to enforce, but it isn't necessarily a smart idea, maybe the compiler can do this better, and maybe these decisions should be left to the programmer. the tool shouldn't decide for us, even if we didn't think about it. changing API shouldn't be done by a tool.

> Guidelines that might not be worth the effort to make toolable
>
> F.4: "If a function might have to be evaluated at compile time, declare it constexpr"
> F.5: "If a function is very small and time-critical, declare it inline"
> F.6: "If your function may not throw, delcare it noexcept"

core guidelines tools and static analyzers tool are available and some are open sourced. there might even be _too many_ options for normal projects. using too many tools and checks create noise.
we can opt in or opt out for checks, in **Clang-Tidy** its either take all the checks except some, or take only some checks.

```clang
*, <disabled-checks>
-*, <enabled-checks>
```

we can have additional checks, like LLVM coding standard, embedded programming checks, MISRA/AUTOSAR for security, and others

MISRA

we can have a diffrent set of operations for development stage and when we release the process.

| Development stage                         | Certification stage               |
| ----------------------------------------- | --------------------------------- |
| Good to have                              | Must have                         |
| Low costs                                 | High costs                        |
| Flexible set of checks, detailed messages | Defined checks and error messages |
| Checks + Quick-fixes                      | Rule violations messages only     |

several standards and sets of guidelines exist (core, MISRA,CERT), and most of them have similar items and recomendations.

### Style and Naming

we also have tools for naming and styles, some of them can live on the build tool chain.
clang format, for example, there are cases when it breaks compatibility, and it has a fuzzy parsing.

Naming is hard

naming conventions require a proper 'renaming' tool.

- camelCase, PascalCase, SCREAMING_SNAKE_CASE
- google style, llvm style, unreal engine conversions.

syntax style, can the tool enforce this?

- east const, west const.
- when is auto used.
- trailing return type, when to use.

an idea for the future:
how to reduce the noise generated by the tools? we can use "game-ify" the tool to motivate us, like create levels of required actions (beginner level, advanced level) to decompose the problem. we can added motivation units (points, score, whatever). it's better to show issues as call to action points than as a list of problems. Team collaborative work is always helpful

### Questions and comments

> "Code analysis only works when it's enforced by tools" - people don't like using external tools that just make the work harder. if we aren't enforcing the checks, they won't be used.
> "Why are there so many standards?" - because different industries

</details>

## CMake + Conan: 3 Years Later - Mateusz Pusz

<details>
<summary>
How CMake changed in the past 3 years. Combining it together with conan package manager
</summary>

[CMake + Conan: 3 Years Later](https://youtu.be/mrSwJBJ-0z8),[Slides](https://github.com/train-it-eu/conf-slides/blob/master/2021.05%20-%20C%2B%2BNow/CMake%20%2B%20Conan%20-%203%20years%20later.pdf)

changes over time with cmake and conan (since his previous lecture three years ago)

### CMAKE isn't a Scriping Language

> **CMAKE** - cross platform c++ build generator (not a general purpose scripting langauge)

cmake moved from variables to _targets_ and _properties_. cmake variables aren't as easy as we think.

what will this code print during this configuration phase?

```Cmake
set (foo 0)

message(foo)
message("foo")
message(${foo})

if(foo)
    message("#1")
endif()
if("foo")
    message("#2")
endif()
if(${foo})
    message("#3")
endif()
```

the correct answer is

> foo \
> foo \
> 0 \

what will this code print during this configuration phase?

```Cmake
set (foo ON)

message(foo)
message("foo")
message(${foo})

if(foo)
    message("#1")
endif()
if("foo")
    message("#2")
endif()
if(${foo})
    message("#3")
endif()
```

the correct answer is

> foo \
> foo \
> ON \
> #1 \
> #3

and if we set foo to abc? (without quotes)

```Cmake
set (foo abc)

message(foo)
message("foo")
message(${foo})

if(foo)
    message("#1")
endif()
if("foo")
    message("#2")
endif()
if(${foo})
    message("#3")
endif()
```

the correct answer this time is

> foo \
> foo \
> abc \
> #1

as we can see, the cmake variables are missleading. heres' another one

```Cmake
set (abc abc)
set (foo abc)

message(foo)
message("foo")
message(${foo})

if(foo)
    message("#1")
endif()
if("foo")
    message("#2")
endif()
if(${foo})
    message("#3")
endif()
```

the correct answer this time is

> foo \
> foo \
> abc \
> #1 \
> #3

different behavior, magic constants, quoted strings, different default behavior, fall-through cases.

and here is an example using cache variables

```Cmake
cmake_minimum_required(3.19)
project(variables NONE)

message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})

set (BUILD_DOCS OFF CACHE BOOL "Docs generation")
message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})
```

on the first run:

> \<empty>\
> ON\
> OFF\
> ON

on the next runs

> OFF\
> ON\
> ON\
> ON\

and if decide to use 'option' instead of 'set'

```Cmake
cmake_minimum_required(3.19)
project(variables NONE)

message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})

option (BUILD_DOCS "Docs generation" OFF)
message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})
```

on the first run (and on the other runs) we get the same output,

> \<empty>\
> ON\
> ON\
> ON

but that depends on the version,, if we use an earlier version

```cmake
cmake_minimum_required(3.12)
project(variables NONE)

message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})

option (BUILD_DOCS "Docs generation" OFF)
message (${BUILD_DOCS})

set (BUILD_DOCS ON)
message (${BUILD_DOCS})
```

we get different output
on the first run:

> \<empty>\
> ON\
> OFF\
> ON

on the next runs

> OFF\
> ON\
> ON\
> ON\

so, we can see variables are a mess.

> Normal and cache variables are two separate things. It is possible to have a normal variable and a cache variable with the same name but holding different values.

- normal variables take precedence over cache variables.
- setting a cache variables value remove the normal variables from the scope
- until Cmake3.13 _option_ was the same as _set_, but it was then changed.

which leads us to the quoute above:

> **CMAKE** - cross platform c++ build generator (not a general purpose scripting langauge)

the less cmake is better, only use cmake for a build system, we should use a dedicated language for scripts (python, etc...), and consider using a package-manager for packages (conan, vcpkg, etc..).

### Good Featres in CMake

#### C++20 Supports (cmake 3.12)

```Cmake
cmake_minimum_required(version 3.12)
add_library(mp-units-core INTERFACE)
target_compile_features(mp-units-core INTERFACE cxx_std_20)
```

#### Simplified Install Destination Handling (cmake 3.14)

before,

> - Every project defind all the destinations by itself
> - Poor consistency among projects
> - hard to make it correct for every platform (lib, lib64)

```Cmake
install(TARGETS myLib Export MyLibTargets
    LIBRARY DESTINATION lib
    ARCHIVE DESTINATION lib
    RUNTIME DESTINATION bin
    INCLUDES DESTINATION include
)
```

and now we can use defaults

```Cmake
include(GNUInstallDirs)
install(TARGETS MyLib EXPORT MyLibTargets)
```

#### MSVC Compilation Warning Handling (cmake 3.15)

before MSVC had deafult warnings flags (gcc and clang didn't have)

```Cmake
function(set_warnings)
string(REGEX REPLACE "/W[0-4]" "" CMAKE_CXX_FLAGS "${CMAE_CXX_FLAGS}")
set (CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}" PARENT_SCOPE)
add_compile_options(
    /W4
    #...
)
endfunction()
```

and now we don't have those special default warning flags for MSVC anymore

```Cmake
function(set_warnings)
add_compile_options(
    /W4
    #...
)
endfunction()
```

#### Ninja Build

single configuration

```bash
cmake -G Ninja ...
```

multi-configuration (cmake 3.17)

```bash
cmake -G "Ninja Multi-Config" ...
```

#### Executing An Install (cmake 3.15)

a quick way to install a project without invoking the whole build tool

before (gcc)

```bash
cmake -DCMAKE_BUILD_TYPE=Release _DCMAKE_INSTALL_PREFIX=~/.local ..
cmake --build .
ctest -VV
cmake --build . --target install
```

now we have something much faster! invoke the install directly!

```bash
cmake --install <bin_dir>
```

we can add some flags for the install

so now the workflow (for gcc or any other compilers) works the same with NINJA

single configuration

```bash
cmake .. -G Ninja -DCMAKE_INSTALL_PREFIX=~/.local -DCMAKE_BUILD_TYPE=Release
cmake --build .
ctest -VV
cmake --install --strip
```

multi-config

```bash
cmake .. -G "Ninja Multi-Config" -DCMAKE_INSTALL_PREFIX=~/.local
cmake --build . --config Release
ctest -VV -C Release
cmake --install . --config Release --strip
```

and a generic script "run.sh"

```bash
cmake .. -G $1 -DCMAKE_INSTALL_PREFIX=~/.local -DCMAKE_BUILD_TYPE=$2
cmake --build . --config $2
ctest -VV -C $2
cmake --install . --config $2 --strip
```

and to run it

```bash
./run.sh "Ninja" Release
./run.sh "Ninja Multi-Config" Release
```

> - single-configuration generators ignore any build-time specification
> - multi-configuration generators ingore the **CMAKE_BUILD_TYPE** variable

#### Setting a Default Generator (cmake 3.15)

adding this to the _.bashrc_ file so we don't need to specify the generator type.
the default build type only works for a ninja multi-config

```bash
#...
# set a default cmake generator
export CMAKE_GENERATOR="Ninja Multi-Config"
export CMAKE_DEFAULT_BUILD_TYPE=Release
```

and now the workflow looks like this

for relese

```bash
cmake .. -G -DCMAKE_INSTALL_PREFIX=~/.local
cmake --build .
ctest -VV -C Release
cmake --install . --strip
```

for debug (not the deafult) we need to specify the config,

```bash
cmake .. -G -DCMAKE_INSTALL_PREFIX=~/.local
cmake --build . --config Debug
ctest -VV -C Debug
cmake --install . --config Debug
```

#### Verbose Builds (cmake 3.14)

before:
switch from ninja to makefile
enable verbosity for makefile generator

```cmake
set(CMAKE_VERBOSE_MAKEFILE ON)
```

now
enable verbosity with a build step command line flag (-v | --verbose)

```bash
cmake --build. -v
```

#### File-Based API (cmake 3.14)

integration with IDEs, better performance, allow configuration of the cmake generator from the ide.

#### Preferring User-Provided Packages (cmake 3.15)

before

find_package preferred system install packages or user provided packaged based on the overload (with CONFIG - user, without - system)

```cmake
find_package(GTest CONFIG REQUIRED)

add_executable(unit-test
#...
)
target_link_libraries(unit-tests PRIVATE
    Gtest::gtest_main
)
```

now there is a flag to set the default behavior. this means we can look for the package in the package configuration file first, which simplifies the usage of package managers.

```cmake
set(CMAKE_FIND_PACKAGE_PREFER_CONFIG ON)
```

and then

```cmake
find_package(GTest REQUIRED)

add_executable(unit-test
#...
)
target_link_libraries(unit-tests PRIVATE
    Gtest::gtest_main
)
```

### Modern Project Structure

the problem with 'add_subdirectory()' for dependencies, different libraries expose different dependencies (headers versions), so libA can use boost v1.66, libB use boost v1.57, and libC can use libA and LibB, so now it has two different versions of boost.

> "Handling dependencies as subdirectories does not scale!"

part of the problem with mono-repos.

not **IMPORTED** CMake targets have global scopes.

> "evn if there are no version conflicts, 'add_subdirectory' still doesn't scale"

name collisions, duplicated targets, multiple names for the same target. same names with different targets.

the more projects we have, the more likely we are to get collisions.

one good practice is to always prefix the target name with the name of the project and alias the name for the linking. it makes using 'add_subdirectory()' less awful, but still bad.

```cmake
add_library(myProject-core
    source_1.cpp
    source_2.cpp
)
add_library(myProject::core ALIAS myProject-core)
```

additionally, we can change the EXPORT_NAME property of a target, so we don't repeat the prefix. so we fixed a problem of none-unique names, and now we need to patch the fix.

```cmake
add_library(myProject-core
    source_1.cpp
    source_2.cpp
)
set_target_properties(myProject-core EXPORT MyProjectTargets)
add_library(myProject::core ALIAS myProject-core)

install(TARGETS myProject-core EXPORT MyProjectTargets)
install(EXPORT MyProjectTargets
    DESTINATION ${CMAKE_INSTALL_LIBDIR}/cmake/myProject
    NAMESPACE myProject::
)
```

projects also have private targets, which aren't exported by the library, we don't wand our dependents to be forced to add them.

> the modern project structure is "Designed to help separate project development workflow from it's usage by dependers"

dependers are not forced to include not-exported targets and aren't affected by our development environment.

example in the slides.

Separating what the development workflow uses and what the end user uses. differencing between public headers (which are exposed outside) and private (used internally in the library).

developers use the top level ./CMakeLists.txt, users user ./src/CMakeLists.txt

for developers, use IDE

```bash
mkDir build && cd build
cmake .. -DCMAKE_INSTALL_PREFIX=~./local
cmake --build .
ctest -VV -C Release
cmake --install . --strip
```

for users, they don't care about compiling and running the unit tests.

```bash
mkDir build && cd build
cmake ../src -DCMAKE_INSTALL_PREFIX=~./local
cmake --build .
cmake --install . --strip
```

If we store dependencies in the subsirctories, we might need download the each time for every project. this takes up space, compile time, and causes the problems we saw above.

It's better to install them once, but this still means many version (because of ABI differences) - build types, compiliers, standard library, versions, preprocessor flags. so this also stops scaling.

we should use package managers for big-scaled projects.

### Conan - The Path Towards 2.0

conan is a package management tool

has a high-quality documentation, jFrog-academy has free courses

```bash
conan [command] -h
```

conan abstracts away the build system. we don't need to learn how to work with each project dependency (how it's being built), the package manager encapsulates all those interfaces to build tools, we can use our build system together with conan for all the external projects.

> 1. Install conan
> 2. Set a development profile
> 3. \[Optional] Add custom remotes
> 4. Create a conanfile
> 5. Provide dependencies with conan

the conan configuration is stored in a default directory, with different profiles which can be changed
~/.conan
~/.conan/profiles/deafult

```bash
pip3 install -U conan
conan profile show default
```

we currently need to manually change the compiler.libcxx setting to use modern standards (default will change in conan 2.0)

```bash
#gcc
conan profile update settings.compiler.libcxx=libstdc++11 default
#clang
conan profile update settings.compiler.libcxx=libc++ default
#c++ 20
conan profile update setting.compiler.cppstd=20 default
#or specify for each installation
conan install .. -s compiler.cppstd=20
#see profiles list
conan profile list
#create debug  based on existing profiles in installation
conan install .. -pr <profile_name> -s build_type=Debug
#use specific profile in installation
conan install .. -pr ../cross_compile
# profile compositian
conan install . -pr=windows -pr=vs2017
conan install . -pr=windows -pr=vs2017 -s build_type=Debug
conan create . -pr=windows -pr=vs2017
#shared profile - obtain and install
conan config install http://githun/com/user/conan_config/.git
```

we can tell the profile where to find the compiler. \
we can include other profile inside the profile.\
we can use specific settings for some packages, we can combine profiles, share proflies, and have company wide configuration.

[ConanCenter](https://conan.io/center/) is a public conan repository for open source packages. moderated and maintained by the conan team.

we can also use a custom remote repository, like a company repo, or a repo that has non standard configuration (which aren't supported by conanCenter yet). jFrog cloud for managing the repo

```bash
#view remote sources
conan remote list
```

the conanfile.txt says what it needs (\[requires]), \[options], \[generators], the dependence's are added automatically.
we call the conan stuff 'recipes', one recipe can represent any number of binaries.

```conan
package_name/package_version@owner/channel
```

conan doesn't build dependencies by deafult, but we can force it to build them.

```bash
conan install -b <none|never|missing|outdate|cascade|patter> ..
```

conan is the one responsible for ousekeeping of different ABI versions,

Conan Generators

- CMakeToolChain - for cmake
- CMakeDeps - multi-configuration generator
- deploy - copies folders, deploy binaries
  \*json - create a json file for packages

we can use conanfile.py to define stronger recipes using python, and then we can import cmake

```bash
conan install .. -pr <your_conan_profile>
conan build ..
```

starting a new library

```bash
conan new -m v2_cmake my_project/0.1.0
```

conan creates a unique identifier for each configuration, which is used to store ABI information. we can also use package_id to override the creation of different versions.

(many more slides)

</details>

## C++ Insights: How Stuff Works, Lambdas and More! - Andreas Fertig

<details>
<summary>
How C++ Insight does some stuff. seeing how common problems are actually created and how the fixes effect them.
</summary>

[C++ Insights: How Stuff Works, Lambdas and More!](https://youtu.be/p-8wndrTaTs),[slides](https://andreasfertig.info/talks/dl/afertig-2021-cppnow-cpp-insights.pdf), [C++ Insights](https://cppinsights.io/)

c++ Insights shows us what is going on, how the compiler handles the source code

> - Show what is going on.
> - Make invisible things visible to assist in teaching.
> - Create valid code.
> - Create code that compiles.
> - Of course, it is open-source

example of showing implicit conversions, this code prints **1**, why is that? cpp insight shows us all the implicit conversions.

```cpp
short int max(short int a, short int b)
{
    return (a > b) ? a : b;
}

void Use()
{
    short int a = 1;
    unsigned short int b = 65'530;
    printf("max: %d\n", max(a, b));
}
```

the above code turns into the code below: numeric comparisons are only possible on integers, not shorts. and the unsigned short is converted to signed int (overflowing into negative), and we can view all the implicit conversions happening.

```cpp
short max(short a, short b)
{
  return (static_cast<int>(a) > static_cast<int>(b)) ? a : b;
}

void Use()
{
  short a = 1;
  unsigned short b = 65530;
  printf("max: %d\n", static_cast<int>(max(a, static_cast<short>(b))));
}
```

c++ insights takes c++ code and returns c++ code, it uses Clang AST (abstract syntax tree), so it's more than just a preprocessor stage. six lines of code turn into thousends of AST code.

there are limitations with templates. what is instantiated and what is not?
default parameters and default member initializer. constexpr functions, differences in c++11 and 14 (implicit const in 11, not in 14).

captures and lambdas: this code captures _a_ by value(copy),but prints 4,5, rather than the expected 4,4. cpp insights show use what was really captured (not the value actually)

```cpp
class Test
{
    int a;

    public:
    Test(int x) : a{x}
    {
        auto l1 = [=] { return a + 2; };
        printf(”l1: %d\n”, l1());
        ++a;
        printf(”l1: %d\n”, l1());
    }
};

int main()
{
    Test t{2};
}
```

the reason being that the capture is actually a class, and it captures the _\*this_ pointer by value, so it reflects the change.

```cpp
class Test
{
    int a;

    public:
    Test(int x) : a{x}
    {
        auto l1 = [*this] { return a + 2; }; // capture the dereferenced object by value
        printf(”l1: %d\n”, l1());
        ++a;
        printf(”l1: %d\n”, l1());
    }
};
```

the better way is to do an init capture, which means we specify exactly what we want.

```cpp
class Test
{
    int a;
    int b{}

    public:
    Test(int x) : a{x}
    {
        auto l1 = [y=a] { return y + 2; }; //init capture of this->a
        printf(”l1: %d\n”, l1());
        ++a;
        printf(”l1: %d\n”, l1());
    }
};

int main()
{
    Test t{2};
}
```

c++20 brought us templated lambdas, but we first need to look at c++14 generic lambdas. cpp insight shows what lambdas were created.

```cpp
int main()
{
    auto max =[](auto x, auto y){
        return (x>y) ? x :y;
    }
    max(2,3); //ok
    max(2,3.0); //works, but not what we wanted, mixed types, integer promotion
}
```

the compiler creates one function for int and int, and one for int and double which returns double.
in c++20, the lambda can be templated and we get better control of the types.

```cpp
int main()
{
    auto max =[]<typename T>(T x, T y){
        return (x>y) ? x :y;
    }
    max(2,3); //ok
    max(2,3.0); //no longer compiles, we decided that both must be the same type
}
```

range statements with temporary objects, this code is Undefined behavior, the Keeper object is a temporary object, and we don't get lifetime extention. cpp insights shows that we actually have a non const reference.

```cpp
struct Keeper
{
    std::vector<int> data{2, 3, 4};

    auto& items()
    {
        return data;
    }
};

Keeper get()
{
    return {};
}

int main()
{
    for(auto& item : get().items())
    {
        std::cout << item << ’\n’;
    }
}
```

c++20 range based for statement with initializer. the life time extension might be part of the standard in c++23, but it's still a long way to go. if we look at this code in c++ insights, we can see that the object is alive for entirety of the loop.

```cpp
int main()
{
    for(auto && obj = get();
        auto & item : obj.items())
    {
        std::cout << item << ’\n’;
    }
}
```

this also allows us to get an index, just like python or JavaScript.

```cpp
 #include <cstdio>
 #include <vector>

int main()
{
    std::vector<int> v{2, 3, 4, 5, 6};

    for(size_t idx{0}; const auto& e : v)
    {
        printf(”[%ld] %d\n”, idx++, e);
    }
}
```

c++20 gave us the spaceship operator. which allows us to replace six operators with one (lower than, greater then, lower equal, greater equal, equal, not equal), we can even default it! .c++ insights lets us see what is instatinated and how it looks. unfortunately, this time we we make use of library internals so the code isn't very readable.

```cpp
struct Spaceship
{
    int x;
    std::weak_ordering operator<=>(const Spaceship& value) const = default;
};

bool Use()
{
    Spaceship enterprise{2};
    Spaceship millenniumFalcon{2};

    return enterprise <= millenniumFalcon;
}
```

but if we change the example and introduce an equality operator for a integer value. switching the order isn't allowed in c++17 (we would need to write a friend function and stuff). but c++20 has **operator reordering** for cases like this. c++ insight shows us this reordering in action. we also see how instatianted member function are `const noexcept` while the spaceship operator isn't (this is done to save some compiler checks, apparently).

```cpp
struct Spaceship
{
    int x;
    auto operator<=>(const Spaceship& value) const = default;
    bool operator==(const int & rhs) const {return rhs ==x;}
};

bool Use()
{
    constexpr Spaceship enterprise{2};
    constexpr Spaceship millenniumFalcon{2};

    //return enterprise ==2; this will work
    return 2 == enterprise; //won't work in c++17
}
```

we can see how the default initialization is happenning. how private and public members effect us. how copy destructors, constructors and assignment operators are created depending on the members in the type. we can see NRVO optimizations in c++ insights. this can help us stop writing _std::move_ when we don't need to.

summary: how can c++ insights help us?

> - Seeing is a very valuable thing. Even if you know something in general, C++ Insights may put your attention on it.
> - Classes I taught using C++ Insights (as well as Matt Godbolt’s Compiler Explorer) tend to be more interactive. Attendees start asking
>   broader questions about certain constructs.
> - C++ Insights can help to settle two different opinions by visualizing what the compiler (at least Clang) does.
> - Like Integrated Development Environments (IDEs), C++ Insights visualizes template instantiations. Seeing them often helps, but seeing the absence of a specific instantiation may lead you to the issue you’re looking for.

</details>

## Interactive C++ in a Jupyter Notebook Using Modules for Incremental Compilation - Steven R. Brandt

<details>
<summary>
Writing and executing C++ code in a jupyter notebook
</summary>

[Interactive C++ in a Jupyter Notebook Using Modules for Incremental Compilation](https://youtu.be/9XWCm9iV-wk)

tools to make teaching c++ easier:

- Cling (an interpreted version of Clang)
- Jupyter
- Docker

which led to the creation of [CXX Explorer](https://github.com/stevenrbrandt/CxxExplorer) but that's an aside.

> notebooks are a tool for experimenting with code:
>
> - each cell is a distinct evaluation with distinct results that build on each other
> - they persist the output of each cell action.
> - cells can contain markdown, not just code.
> - usually based on python, but not necessarily.
> - we can use `%%` cells to execute non-python code.
> - this makes them great as teaching tools.

notebooks contain documentation, code and the output of that code in one executable!

python cell

```py
from IPython.core.magic import register_cell_magic
@register_cell_magic
def bash(line, cell):get_ipython().system(cell)
```

bash magic cell

```sh
%%bash
echo Hello, world!
```

Docker is a lightweight container that uses the linux kernel. it encapsulates the build/installation process.
[docker hub image](https://hub.docker.com/r/stevenrbrandt/clangmi), [repository with compose files](https://github.com/stevenrbrandt/module-interactive).

Cling is based on Clang, it's an interpreted version of clang. there are some problems, when encountering a bug, it crashes entirely, which makes teaching harder.

Modules can help us overcome those problems if we use them in notebooks. modules provide incremental compilation and chainning. we have defintion cells and run code cells.

jupyter notebook cell to create a cpp module

```cpp
//ipython magic
%%writefile aloha.cppm

export module aloha;
#include <iostream>
export {
    void aloha_world()
    {
        std::cout<<"Aloha, world!\n";
    }
}
```

jupyter notebook cell to compile the module, we need to compile it twice

```sh
%%bash
rm -f aloha.pcm aloha.o
clang++ -std=c++20 -fmodules-ts \
--precompile -x c++-module -c aloha.cppm \ -fimplicit-modules -fimplicit-modules-maps \
-stdlib=libc++ # create .pcm file

clang++ -std=c++20 -fmodules-ts \
-c aloha.cppm \
-fimplicit-modules -fimplicit-modules-maps \
-stdlib=libc++ # create .o file
```

another cell to write the driver code

```cpp
//ipython magic
&&writefile aloha.cpp

import aloha;
int main(){
    aloha_world();
}
```

and a cell to execute code

```sh
%%bash
clang++ -std=c++20 -fmodules-ts -o aloha aloha.cpp \
aloha.o -fimplicit-modules -fimplicit-module-maps \
-stdlib=lib++ -fmodule-file=aloha.pcm

./aloha
```

this is a lot of typing for each file, so there's a python package that hides it away. this is where the **def_code** and **run_code** stuff comes into

this simple cell

```cpp
%%def_code
std::string hello= "Hello";
```

is evaluated into a complete c++ module

```cpp
export module tmp1;
export import clangmi; //initial loads
export {
    std::string hello = "Hello";
}
```

including the compilation step (as a module and a .o file) and archiving into a library.

the cells that run the code are actually importing the module and create a simple program that is compiled and uses the code from the cell.

```cpp
%%run_code
std::cout<< hello << '\n';
```

now we have cells that can build objects and cells that can run simple programs. with each cell we can change the verbosity level to display more or less details about the cell.

however, in this version, each time we run a cell, we do all the computations again and again. this is a problem.

we would want to use constexpr functions, so that the value will be computed once.

```cpp
%%def_code
constexpr int fib(int n)
{
    if (n<2)
    {
        return n;
    }
    else
    {
        return fib(n-1)+fib(n-2);
    }
}
```

but this doesn't work, it pushes the the data into the _.o_ file, but not the _module_.

there are some hacks with using external variables and storing the results in a library. maybe in the future it'll be easier. the problem still remains that we store those objects in a library, we will constantly use disk space.

the other way is to use shared memory, lets create a counter class.

```cpp
%%def_code

struct Counter{
    int n;
    Counter():n(0){}
    ~Counter(){std::cout<<"reset Counter\n";}
    void count(){
        std::cout << "n="<<(n++)<<'\n';
    }
};
```

creating the shared memory, if we run the code again and again the Counter persists and changes the value!

```cpp
%%run_code
Seg seg("mem");
Counter *c = seg.allocate<Counter>("counter");
c->count();
if (c->n ==5)
{
    seg.remove(c);
}
```

with some special code for arrays

```cpp
%%run_code
Seg seg("mem");
Array<double>& arr = *seg.allocate_array<double>("date",100);
f (arr.init())
{
    //if first invocation
    std::cout<<"init\n";
}
//remove array
seg.remove(&arr);

```

and now lets use this array in a semi-real example, a sinusoidal wave

```cpp
%%run_code
#include <math.h>

Seg seg("mem");
const int N=100;
Array<double>& a = *seg.allocate_array<double>("date1",N);
Array<double>& b = *seg.allocate_array<double>("date2",N);

double dx = 15.0/a.size();
for (int i = 0; i<a.size();i++)
{
    double x =i *dx;
    a[i]=x;
    b[i]=sin(x);
}
```

it would be nice to use the cpp data directly in python code, so we use the python library of **pybind11** to intgrate python and cpp code.

```py
import clangmi
import numpy as np
import matplotlib.pyplot as plt

a_buf = clangami.allocate_array("mem","data1",100)
a = np.asarray(a_buf)
b_buf = clangami.allocate_array("mem","data2",100)
b = np.asarray(b_buf)

plt.plot(a,b)

```

theres also parallel processing, even if clang actually has a bug!

```cpp
%%run_code
import <future>

auto a= std::async(std::launch::async,[](){return 42;});
std::cout<< "a="<<a.get() <<'\n';
```

bringing in the hpx package requires some ugly work in python. \
using the hpx code takes much longer to compile.

```cpp
%%run_code

#include <hpx/hpx.main>
#include <hpx/hpx_main.hpp>

auto a = hpx::async([](){return 42});
std::cout << a.get() <<'\n;
```

we can use hpx to actually run on multiply threads

```
runcode.flags=["-t","4"]
```

and then run the code, even if the output get jumbled.

```cpp
%%run_code
#include <hpx/hpx_main.hpp>
#include <hpx/algorithm.hpp>
#include <hpx/execution.hpp>

std::vector<std::size_t> v{1,2,3,4,5,6};
hpx::for_loop(hpx::execution::par,0 v.size(),
[](std::size_t n){std::cout<< "n = " << n << '\n';});
```

theres also the `.then()` to avoid blocking and make composable code.

in conclusion, the c++ jupyter notebook is a a prototype, with some hurdles to overcome.

</details>

## The Quest For A Better Crash - Victor Ciura

<details>
<summary>
Introducing a framework to get and analyse crash reports.
</summary>

[The Quest For A Better Crash](https://youtu.be/pJPRdNTxL-E),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/The-Quest-For-A-Better-Crash-Victor-Ciura-C-Now-2021.pdf).

> - _CRT_ - C Runtime Library.
> - _SEH_ - structured exception handling.
> - _ISA_ - Instruction Set Architecture.

getting the crash report, reproducing it, investigating it, and having the appropiate infrastructure for this.

> Vignette in 3 parts
>
> 1. Remember the crash
> 2. Roll your own
> 3. The Future: post-pandemic crashes

note: there are some differences between how windows and linux machines handle exceptions and stack traces.

### Remember The Crash

we have crash reports (**windows error reporting**), which we can sometimes send a report, but how do we use the same capabilities as developers?

- can we register to receive crash dumps?
- how does the crash report look like?
- where is it stored?

in the past it was custom regiteration which required a **Microsoft symbol server** deployed on premise, where each version had to be separately registered, and the whole thing was less in the developer level, and more of a configuration nightmare.

### Roll Your Own

if we want a crash reporting infrastructure, what should it be like?

- **quick to develop.** this isn't the main product, it's an add-on. it shouldn't take ages to develop and get right.
- **easy to integrate into the CI/CD (no special service, no symbol server).** no additional custom work to configure.
- **zero footprint on client side (not shipping symbols).** this shouldn't bloat our binaries.
- **zero performance impact on Release Binaries (on the happy path).**
- **east to use standalone tool (non-dev machine) for processing crash reports.**

it needs to be integrated with the ci/cd tools (jenkins, gitlab, build artifacts), so we need to remember the symbols for each release and have a tool that works for each version, some build systems are non-determinstic, so we need to be prepared. each build can contain different symbols, which is how we process the crash, so we need to have the symbols stored.

the symbols are stored in the **.pdb** extentions files ([pdb - program database](https://en.wikipedia.org/wiki/Program_database)). there are some variations of how the pdb files are build (incremental, full build):

[Debug Information options](https://docs.microsoft.com/en-us/cpp/build/reference/debug-generate-debug-info?view=msvc-170)

> - Generate Debug Information (`/DEBUG`)
> - Generate Debug Information optimized for faster links (`/DEBUG:FASTLINK`)
> - Generate Debug Information optimized for sharing and publishing (`/DEBUG`)

a stack trace can be "dry" or "full", where full stack traces contain enviornment, context and names, just like when we crash in the development cycle. so when we collect crash reports, we need to collect that data (and anonymise it), add back the symbols (which we remove in the release build) and be able to recreate the situation.

### How is works

a lot of windows programs use SEH, and even _asynchronous_ exception (`/EHa`), so our workflow needs to support those. this is one of the settings for the build process.

`/EHa /DEBUG:FULL /Zi`

[Debug Information Formats](https://docs.microsoft.com/en-us/cpp/build/reference/z7-zi-zi-debug-information-format?view=msvc-170)

- `/Z7` - object file with full symbolic debugging information built in.
- `/Zi` - separate pdb file with debuging symbols.
- `/ZI` - like `/Zi`, but a pdb file that supports _edit and continue_.

we also need to set ourselves to handle C structured exceptions (primitive exceptions) as C++ typed exceptions.

`_set_se_translator(ExceptionHandling:TransFunc);`

we also want to intercept exceptions if they happen in a process which we don't debug. so we intercept them.

```cpp
static bool installedFilter = false;
if (!installedFilter)
{
    ::SetUnhandledExceptionFilter(ExceptionHandling:UnhandledException);
    installedFilter = true;
}
```

> If an exception occurs in a process that is not being debugged, and the exception makes it to the Unhandled exception filter => we intercept it.
>
> This replaces the existing top-level exception filter for ALL existing and ALL future threads in the calling process.

this is like having a massive try-catch block that encompasses everything. we can use this opportunity to log the issue.

below is the logging function example and the translation between C (primitive) exception into C++ typed exceptions.

_(I'm not super clear about this code)_

```cpp
LONG ExceptionHandling::UnhandledException(EXCEPTION_POINTERS * aExceptionInfo)
{
    wstring message(L"[EXCEPTION_UNHANDLED] ");
    wchar_t buf[MSG_BUFFER_LEN];
    swprintf_s(buf, MSG_BUFFER_LEN, L"(0x%.8x) at address " ADDRESS_FORMAT SW_EOL,
    aExceptionInfo->ExceptionRecord->ExceptionCode,
    aExceptionInfo->ExceptionRecord->ExceptionAddress);
    message += buf;

    StackWalker::TraceFromContext(message, aExceptionInfo->ContextRecord);
    ErrMsgPresenter::Message(message);
    return EXCEPTION_EXECUTE_HANDLER;
}

void ExceptionHandling::TransFunc(unsigned int aSECode, EXCEPTION_POINTERS * aExInfo)
{
    // write the exception prolog (type, code, address, etc.)
    switch (aSECode) // decode SEH exception type
    {
        case EXCEPTION_ACCESS_VIOLATION:
            swprintf_s(buf, MSG_BUFFER_LEN, L"%hs (0x%.8x) at address " ADDRESS_FORMAT SW_EOL,
            "ACCESS_VIOLATION", EXCEPTION_ACCESS_VIOLATION,
            aExInfo->ExceptionRecord->ExceptionAddress);
            break;
        case EXCEPTION_DATATYPE_MISALIGNMENT:
            break;
        case EXCEPTION_INT_DIVIDE_BY_ZERO:
            break;
        case EXCEPTION_INT_OVERFLOW:
            break;
        case EXCEPTION_ILLEGAL_INSTRUCTION:
            break;
        case EXCEPTION_STACK_OVERFLOW:
            break;
        /*...*/
    }

    SehException::SEType seType = SehException::SEH_GENERIC;
    // for AV exception, we can determine the type of operation that caused it
    if (aSECode == EXCEPTION_ACCESS_VIOLATION)
    {
        // the first element of the array contains a read-write flag
        // that indicates the type of operation that caused the access violation
        ULONG_PTR operationType = aExInfo->ExceptionRecord->ExceptionInformation[0];
        // the second array element specifies the virtual address of the inaccessible data
        ULONG_PTR virtualAddress = aExInfo->ExceptionRecord->ExceptionInformation[1];
        if (operationType == 0)
            seType = virtualAddress ? SehException::SEH_AV_READ_BADPTR : SehException::SEH_AV_READ_NULLPTR;
        else if (operationType == 1)
            seType = virtualAddress ? SehException::SEH_AV_WRITE_BADPTR : SehException::SEH_AV_WRITE_NULLPTR;
        else if (operationType == 8)
            seType = virtualAddress ? SehException::SEH_AV_DEP_BADPTR : SehException::SEH_AV_DEP_NULLPTR;
    }
    // record SEH type info in exception message
    exceptionMsg.insert(0, L"[" + SehException::SeTypeToString(seType) + L"] ");

    // write the exception prolog (type, code, address, etc.)
    // decode SEH exception type
    /*...*/

    // walk the function call stack and gather information about each frame
    StackWalker::TraceFromContext(exceptionMsg, aExInfo->ContextRecord);
    // for AV exception, we can determine the type of operation that caused it
    /* ... => seType */

    // extract SEH exception origin from StackTrace
    SymbolUtil::SrcPos exOrigin = GetExceptionOrigin(aExInfo->ContextRecord);

    // throw a C++ typed exception with the necessary fault information (attached)
    throw SehException(exOrigin.mFile, exOrigin.mLine, seType, exceptionMsg);
}
```

we also want the stack trace for the exception in the current thread, this gives us the context, which is a struct called **PCONTEXT** in windows. this is alo held as part of the _EXCEPTION_POINTERS_ structures

```cpp
/*utility function*/
std::wstring ExceptionHandling::GetStackTraceForCurrentException()
{
    std::wstring stackTrace;
    StackWalker::TraceFromContext(stackTrace,ExceptionHandling::GetCurrentExceptionContext());
    return stackTrace;
}
```

(aside, we can give the same struct different names, and special names for a pointer. [stackoverflow](https://stackoverflow.com/questions/30370036/how-a-struct-being-typedef-ed-to-multiple-names))

```cpp
typedef struct _EXCEPTION_POINTERS
{
    PEXCEPTION_RECORD ExceptionRecord;
    PCONTEXT ContextRecord;
} EXCEPTION_POINTERS, *PEXCEPTION_POINTERS;
```

and also the _EXCEPTION_RECORD_ stucture, which can link to other exceptions and has different meanings based on the architecture (32/64 bit).

```cpp
typedef struct _EXCEPTION_RECORD
{
    DWORD ExceptionCode;
    DWORD ExceptionFlags;
    struct _EXCEPTION_RECORD * ExceptionRecord;
    PVOID ExceptionAddress;
    DWORD NumberParameters;
    ULONG_PTR ExceptionInformation[EXCEPTION_MAXIMUM_PARAMETERS];
} EXCEPTION_RECORD;

typedef struct _EXCEPTION_RECORD32 {
    DWORD ExceptionCode;
    DWORD ExceptionFlags;
    DWORD ExceptionRecord;
    DWORD ExceptionAddress;
    DWORD NumberParameters;
    DWORD ExceptionInformation[EXCEPTION_MAXIMUM_PARAMETERS];
} EXCEPTION_RECORD32, *PEXCEPTION_RECORD32;

typedef struct _EXCEPTION_RECORD64 {
    DWORD ExceptionCode;
    DWORD ExceptionFlags;
    DWORD64 ExceptionRecord;
    DWORD64 ExceptionAddress;
    DWORD NumberParameters;
    DWORD __unusedAlignment;
    DWORD64 ExceptionInformation[EXCEPTION_MAXIMUM_PARAMETERS];
} EXCEPTION_RECORD64, *PEXCEPTION_RECORD64;
```

if we want this **PCONTEXT**, we can write a function, but it will be different based on the version of visual-studio

`#if _MSC_VER >= 1900` visual studio 2015-19

```cpp
PCONTEXT ExceptionHandling::GetCurrentExceptionContext()
{
    __vcrt_ptd * pTid = nullptr;
    #ifdef _DLL // Multi-Threaded DLL /MD or /MDd

    pTid = (__vcrt_ptd *) (((BYTE *)__current_exception_context())
    - offsetof(__vcrt_ptd, _curcontext));

    #else // Multi-Threaded /MT or /MTd

    pTid = __vcrt_getptd();
    #endif
    return (CONTEXT *)pTid->_curcontext;
}
```

`#if _MSC_VER < 1900` visual studio 2013

```cpp
PCONTEXT ExceptionHandling::GetCurrentExceptionContext()
{
    _tiddata * pTid = nullptr;
    #ifdef _DLL // Multi-Threaded DLL /MD or /MDd

    pTid = (_tiddata *) (((BYTE *)__pxcptinfoptrs())
    - offsetof(_tiddata, _tpxcptinfoptrs));

    #else // Multi-Threaded /MT or /MTd

    pTid = _getptd();
    #endif
    return (CONTEXT *)pTid->_curcontext;
}
```

those functions are hard to get, it's a lot of undocumentaed API.

```cpp
#include <eh.h>
#include <signal.h> // for use of API void ** __pxcptinfoptrs()
#if _MSC_VER >= 1900
 #include <../CRT/src/vcruntime/vcruntime_internal.h>
 extern "C" __vcrt_ptd * __cdecl __vcrt_getptd();
 extern "C" void ** __cdecl __current_exception_context();
#else
 // for use of (private) API _tiddata * _getptd()
 #include <../CRT/src/mtdll.h>
#endif
```

the data itself looks like this, this allows us to get the context for the stack trace.

```cpp
// per-thread data
typedef struct __vcrt_ptd // #include <../CRT/src/vcruntime/vcruntime_internal.h>
{
    // C++ Exception Handling (EH) state
    unsigned long _NLG_dwCode; // Required by NLG routines
    unexpected_handler _unexpected; // unexpected() routine
    void * _translator; // S.E. translator
    void * _purecall; // called when pure virtual happens
    void * _curexception; // current exception
    void * _curcontext; // current exception context
    int _ProcessingThrow; // for uncaught_exception
    void * _curexcspec; // for handling exceptions thrown from std::unexpected
    int _cxxReThrow; // true if it's a rethrown C++ exception
#if defined _M_X64 || defined _M_ARM || defined _M_ARM64
    void * _pExitContext;
    void * _pUnwindContext;
    void * _pFrameInfoChain;
    uintptr_t _ImageBase;
    uintptr_t _ThrowImageBase;
    void * _pForeignException;
#elif defined _M_IX86
    void * _pFrameInfoChain;
#endif
} __vcrt_ptd;
```

we also want to the stackTrace from the context of the caller, how the user called the function. even if there was no exception, maybe we want to identify how our code is being used.

```cpp
void StackWalker::TraceFromCaller(wstring & aStackMsg)
{
    using PF_RtlCaptureContext = void(WINAPI *)(PCONTEXT aContextRecord);

    // dynamically load the RtlCaptureContext() kernel API
    static auto CaptureCtx = (PF_RtlCaptureContext)::GetProcAddress(
    ::LoadLibraryA("Kernel32.dll"), "RtlCaptureContext");
    CONTEXT context;
    ::ZeroMemory(&context, sizeof(context));

    // retrieve the context record of the caller function
    CaptureCtx(&context);
    StackWalker::TraceFromContext(aStackMsg, &context);
}
```

and now that we have the context, we need to do something with, we want to _walk the stack_, we see this in many of the code snippets so far.

```cpp
std::wstring stackTrace;
StackWalker::TraceFromContext(stackTrace, GetCurrentExceptionContext());
```

so lets look at it

```cpp
void StackWalker::TraceFromContext(wstring & aStackMsg, PCONTEXT aContext, int MaxFrameCount)
{
    // All <DbgHelp> functions, such as StackWalk(), are single threaded.
    // (calls from more than one thread to this function will likely result
    // in unexpected behavior or memory corruption)
    // => we must synchronize all concurrent calls to this function
    SyncGuard guard(sEHSyncSupport);

    // Copy the given machine CONTEXT structure because the StackWalk() API
    // might modify it and subsequent calls needing the CONTEXT will fail
    CONTEXT context;
    ::CopyMemory(&context, aContext, sizeof(context));

    HANDLE hProcess = ::GetCurrentProcess();
    HANDLE hThread = ::GetCurrentThread();
    // create a symbol explorer
    SymbolUtil symMgr;
    if (!symMgr.Init(hProcess))
    return;

    /*...*/

    // initialize the STACKFRAME according to the platform we are working on (PE type)
    STACKFRAME sf;
    DWORD imageType = InitStackFrameFromContext(&sf, &context);
    for (int frmIndex = 0; frmIndex < MaxFrameCount; frmIndex++)
    {
        // get the current frame info
        BOOL result = ::StackWalk(imageType, hProcess, hThread, &sf, &context, nullptr,
        SymFunctionTableAccess, SymGetModuleBase, nullptr);
        if (!result)
            break;
        aStackMsg += symMgr.ComposeStackFrame(sf.AddrPC.Offset);
    }

    // write the module load address - needed because of ASLR (Address Space Layout Randomization)
    aStackMsg += symMgr.ComposeModuleBaseAddress();
}
```

this introduces another utility function `InitStackFrameFromContext` and the `STACKFRAME` struct.

```cpp
DWORD InitStackFrameFromContext(LPSTACKFRAME aStackFrame, PCONTEXT aContext)
{
    ::ZeroMemory(aStackFrame, sizeof(STACKFRAME));
#if defined _M_IX86
    DWORD imageType = IMAGE_FILE_MACHINE_I386;
    aStackFrame->AddrStack.Offset = aContext->Esp;
    aStackFrame->AddrStack.Mode = AddrModeFlat;
    aStackFrame->AddrFrame.Offset = aContext->Ebp;
    aStackFrame->AddrFrame.Mode = AddrModeFlat;
    aStackFrame->AddrPC.Offset = aContext->Eip;
    aStackFrame->AddrPC.Mode = AddrModeFlat;
#elif defined _M_X64
    DWORD imageType = IMAGE_FILE_MACHINE_AMD64;
    aStackFrame->AddrStack.Offset = aContext->Rsp;
    aStackFrame->AddrStack.Mode = AddrModeFlat;
    aStackFrame->AddrFrame.Offset = aContext->Rbp;
    aStackFrame->AddrFrame.Mode = AddrModeFlat;
    aStackFrame->AddrPC.Offset = aContext->Rip;
    aStackFrame->AddrPC.Mode = AddrModeFlat
#endif
    return imageType;
}
```

> _ASLR_ - Address Space Layout Randomization

we need the baseModuleBaseAdderess to counter the ASLR and find the correct real module base address.

```cpp
// Serialize module base address - needed because of ASLR
wstring SymbolUtil::ComposeModuleBaseAddress()
{
    std::wstring stackFrame;
    HMODULE moduleLoadAddress = ::GetModuleHandle(nullptr);
    wchar_t buf[MAX_PATH];
    swprintf_s(buf, MAX_PATH, ADDRESS_FORMAT L" ", (size_t) moduleLoadAddress);
    stackFrame += buf;
    stackFrame += SW_MODULE_LOAD_ADDRESS;
    stackFrame += SW_EOL;
    return stackFrame;
}
```

### _Symbolicate_ tool to handle crash reports

in the tool:

- select build
- select crash report
- get a full stack trace

the tool comes in two instruction set architecture flavours: X86 and X64. the tool must match the architecture of the debugged (crashed) process.

functions (_not copying the code here_):

they use a lot of helper functions from the debug helper library.

- header: `#include <dbghelp.h>`
- linker: `/LINK Dbghelp.lib`
- dynamic dependency on **Dbghelp.dll**

```cpp
bool SymbolUtil::Init(HANDLE aProcess, const std::wstring & aSymbolsPath);
bool SymbolUtil::SetSymbolSearchPath(std::wstring symbolSearchPath);
std::wstring StackTraceAnalyzer::Symbolicate(const std::wstring & aRawCrashReport);
std::wstring StackTraceAnalyzer::ProcessStackFrame(const StackFrame & aStackFrame);
std::wstring SymbolUtil::ComposeStackFrame(DWORD_PTR aAddress);
std::wstring SymbolUtil::SymbolNameFromAddress(DWORD_PTR aAddress) const;
std::wstring SymbolUtil::SymbolSourceFromAddress(DWORD_PTR aAddress) const;
```

### Out of the box Alternative - Minidump

in the recent years, minidump got some major enhancement.

minidump has the _.dmp_ extension, it works with the Windows snapshot process, visual studio can parse and open the file, and point at the location the error occurred. it's a very nice experience. it even supports _life share_, as shared debugging session/

### Post pandemic crashes

where are we going from here? how will things look like in future versions of c++?

there is a proposal that is based on **Boost.Stacktrace**, `#include <stackrace>`, which is in the works for many years,but didn't make it into _c++20_.

> C++23 \<stracktrace> Key features (desired):
>
> - all functions are **lazy**: do not query the stacktrace entry info without explicit request
> - **dynamic size** for trace. all the available invokers must be stored in a stacktrace.
> - implementations. allow to **disable/enable** gathering stacktraces by a linker switch.
> - stacktracing shouldn't prevent any **optimizations**.
> - stacktrace should be **usable** in contract violation handler, coroutines, handler functions, parallel algorithms.
> - `stacktrace_entry::description()` should return a **demangled** function signature.
> - `to_string(stacktrace)` should query information from **debug symbols**, symbol export tables and any other sources, returning **demangled signatures**.
> - information about **inlined functions** that have no separate stacktrace entries is welcomed `-> to_string(stacktrace)`.
> - **avoid doing heavy** operations in `basic_stacktrace` constructors or
>   `stacktrace_entry::current()`..

the datatype will have a native handle (just like the standard thread has),will have the source code locations (function name, file, line), it will have a dynamic memory and ways to iterate over them. a way to get the current stack frame.

</details>

##

[Main](README.md)
