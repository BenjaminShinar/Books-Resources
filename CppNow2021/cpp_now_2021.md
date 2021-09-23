# CPPNow 2021

[CppNow 2021 conference lectures](https://youtube.com/playlist?list=PL_AKIMJc4roXvFWuYzTL7Xe7j4qukOXPq)

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

## The C++ Rvalue Lifetime Disaster - Arno Schödl

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

## A Crash Course in Unicode for C++ Developers - Steve Downey

<details>
<summary>
Different types of encodings, encoder/decoder (compose / decompose) algorithms. unicode in c++.
</summary>

[A Crash Course in Unicode for C++ Developers](https://youtu.be/iQWtiYNK3kQ)
[unicode](http://unicode.org)
[utf8 encoding](https://en.wikipedia.org/wiki/UTF-8)

std::u8string

code units, code point, graphemes, abstract characters.

> code units
>
> - char
> - wchar_t
> - octet
> - Word

code points and scalar values
grapheme clusters, extended grapheme clusters.

> utf-8 is good
>
> - C string safe
> - No aliasing
> - Self syncing
> - Single errors lost one character
> - ASCII compatible
> - Start is easy to find

a table about how we encode different ranges of values into different bytes.
some 'well formed' byte sequences.

utf-16 and utf-32. if the value fits inside 16 bit, then put it in one, otherwise, split it into two code points (surrogate pairs).
ucs-2, ucs-4. wtf-8 (wobbly transformation format), wtf-16

### Encoding and Decoding

> Encoders take text and output octets.
> Decoders take octets and output text.
> Text is this context is scalar values.

in utf8 the order is set.
in utf-16 there are byte order marks, for big and little endian.

legacy encoding from before unicode

> - Windows 1252, 125x
> - ISO-8859-x
>   ...others

multi-byte encodings

transcoding, from one character set to another.

### Normalization

combined text might have more than one representation, like special forms, canonical equivalence and compatible equivalence

canonical equivalence:

three difference ways to produce the same symbol.  
latin capital with a rings &#x00c5; is like angstrom sigh &#x212b; or combining the letter with the symbol &#x0041;&#x030A;

compatible equivalence:
not the same symbol, losing some data, but meaning is preserved (mostly)

decomposed and composed text. there are some forms of charters that are already predefined, but can also be created by composing different code points together,we have

NFD,NFD,NFKD,NFKC are forms for different usages, like search (human), identifiers in linkage (strong equality), decomposing ignored diacritics.

nfc is the least risky in terms of information lost.

quick_check of code points to test if it's normalized: yes, no and maybe.

stream safe text format, a way to avoid some problems that can occur in full normalization, so there's a stream-safe format.

the unicode character database. UCD files (txt files and xml files) - all sorts of data.

theres an issue with emojis.

### Algorithms

there are still many problems in the standard solutions, different text directions. word wrapping (line breaks): positions where is's possible to break between lines, there are many ways to get this wrong. text segmentation: from data into user perceived characters, words and sentences.

unicode and regular expressions. matching words on word boundaries. sentences embedded within sentences.

### The Future For C++

what might be in the future, c++23 and c++26 are what they hope to achieve.

| version | features                   |
| ------- | -------------------------- |
| C++20   | char8_t                    |
| C++23   | literal Encoding           |
|         | Portable Source Code       |
|         | Encoding / Decoding Ranges |
| C++26   | Algorithms with ranges     |

</details>

## How to: Colony - Matthew Bentley

<details>
<summary>
Explaining how the current plf::colony container works. it might become part of the standard library.
</summary>

[How to: Colony](https://youtu.be/V6ZVUBhls38)

this talk is about a data structure called **Colony** that is int the process of being added to the standard.

[PlfLib - Some Header-only libraries](https://plflib.org/)

the main thing about it is that it maintains pointers/iterators/reference integrity.

> use scenarios
>
> - you have a lot of unordered data and you're erasing/inserting on the fly.
> - you have multiple collections of interrelated data.
> - preserving the pointer/iterator validity of non-erased elements is important

started from the design of a game, entities that have bidirectional references and many interrelationships, with a lot of inserting and creating entities which link to one another.

an existing solution is a 'bucket array' and entries can be active or inactive to determine if they're processed or not.

### Core Aspects

> Three Core Aspects
>
> 1. A collection of element memory blocks + metadata, to prevent reallocation during insertions (as opposed to a single memory block).
> 2. A method of skipping erased elements in O(1) time during iterations (as opposed to reallocating subsequent elements during erasure, or doing individual allocation of elements)
> 3. An erased-element location recording mechanism, to enable the re-use of memory from erased element in subsequent insertions, which in turn increases cache locality and reduces the number of block allocations/de-allocations.

Avoid reallocations to avoid invalidating iterators/pointers/references. Reuse memory locations that have been erased, keep the data together for cache locality.

A collection of element memory blocks + metadata

> - Can do linked list of blocks, or vector of pointers to blocks.
> - Blocks have growth factor, so can not do vector of blocks.
> - Minimum/Maximum block capacities can be user-defined.
> - Can house block metadata separately, or together in a struct.
> - Metadata includes skip-field of other erasure-skipping mechanism, and any data related to erased-location recording.
> - Necessary metadata:
>   - size - to remove blocks once empty.
>   - capacity - to ascertain end-of-block.

we remove empty blocks to maintain iteration at O(1). but we can retain them as reserved blocks for later use.

considerations about which block to retain.
A method of skipping erased elements in O(1)

> Definitions:
>
> - LCJC:'low complexity jump-counting'.
> - HCJC:'high complexity jump-counting'.
> - Block: a colony's element memory block.
> - Skipfield: an array of intgers of bits used to skip over certain object in an accompanying data structure during iteration. Separate from the elements.
> - Skipblock: a run of skipped nodes within a skipfield.
>   - Start node: the first node in any skipblock.
>   - End node: the last node in any skipblock.
>   - Middle node: any non start/end node in skipblock.

booleans SkipFields are good because they simple, and might be usefull for multi-threaded environments (atomics, etc). they are prolemeatic because they aren't constant time (not O(1)), cause branching and latency.

low and high complexity jump counting:
time-complexity of algorithms differs for modification of fields, but both have O(1) for iteration.
HCJC allows recording of middle nodes, LCJC doesn't allow.

> acronym "Theyaton" - Traversing Homogenus Elements Yielding Amortised Time O(n).

> boolean skipfield example
> 0 0 0 1 1 1 1 1 1 0 0 0
> equivalent HCJC
> 0 0 0 6 2 3 4 5 6 0 0 0
> the 6's are the start and end nodes. the other numbers are Middle nodes.Start and End record the length of runs of erased elements. Middle node record left distance to first non erased element.

```
// skipping
++i;
i +=S[i];
```

> equivalent LCJC. multiple forms.
> 0 0 0 6 2 6 7 3 6 0 0 0
> 0 0 0 6 0 0 0 0 6 0 0 0
> 0 0 0 6 0 5 2 1 6 0 0 0

The middle nodes are ignored and have no meaning. The start and end nodes record the skip length.
good for recording and re-using skipblocks rather than individual skipped nodes. lower time complexity O(1) in all operations, where as HJCJ can have undefined time complexity, fewer instructions overall.

```
// skipping
do {
    ++i;
} while(S[i] == 1);
```

we can have per-memory-block skip fields or global skipfields. global skipfields can create problems of reallocation (invalidating iterators and causing latency). the bit-width of the skipfield must be able to describe the memory block, so it must be about the same size (capacity -1),

possible ideas with skipfields:

- using two 8 skipblocks instead of 16 bit skipblocks. forces some computations.
- using a boolean bitfield and storing the skipdata instrad of the elements. forces more memory reads.

colony once used a stack of pointers, which was problematic because it meant creating memory allocation during erasure, which is not up to the standarts.

**"Free List"**

> Linked list of erased elements (typically singly-linked) using erased element memory space reinterpert_cast'd via pointers as linked list nodes.
> requires over-alignment of the type to the width of a free list node.
> per-block (not global) free-lists reduce bit depth.
> a global free-list must use pointers (not indexes),also causes O(N) when erasure.

effect of removing a block requires something with the skip-block. something about doubly-linked free list.

summary of the current implementation, the container structure and the iterator structure.

example with a blackboard.

extra operations:

- advance/next/prev/distance/range-erase optimization.
- range/fill/initializer_list insert/assign optimization.
- splicing.
- sorting.

</details>

## When Should You Give Two Things the Same Name? - Arthur O'Dwyer

<details>
<summary>
Do we need to call methods by the same name? when and why? 
</summary>

[When Should You Give Two Things the Same Name?](https://youtu.be/OQgFEkgKx2s)

> - When do ue us classical inheritance.
> - Idiosyncratic philosophical digressions.
> - Copious anecdotes from the STL.
> - Kind of a major rabbit-hole about constructors.
> - Mental templates, macros and polyfills
> - Bonus mantras and takeaways.

### The role of (OO) Inheritance

What do we expects from inheritance?/
We expect virtual functions and somewhere that they're used: polymorphic methods, deletion of pointers through the virtual destructor...

```cpp
struct Animal{
    virtual void feed();
    virtual ~Animal()=default
};
struct Cat: public Animal{
    // what do we expect here?
    //probably an override of feed()
};
```

but if we see code with value-semantic and none polymorphic code, we will be confused.

```cpp
Cat acquirePet();
void foo(Cat & current)
{
    auto newPet = acquirePet();
    std::swap(current, newPet);
}
```

the two approches can be combined (public inheritance without polymorphism)

> - EBO - [Empty Base Optimization](https://en.cppreference.com/w/cpp/language/ebo)
> - CRTP - [Curiously Recurring Template Pattern]https://en.cppreference.com/w/cpp/language/crtp
> - TagDispatch

but they are more of a corner cases, not the intended usage of inheritance.

```cpp
template <class Allocator>
struct CatEBO : Allocator{

};

struct CatCRTP : CanFightWith<CatCRTP>
{

};

struct CatTagDispatch: AnyAnimalTag
{

};
```

according to Liskov's substition principle:

> "**If** for each object o1 of type S **there is** an object o2 of type T **such that** all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2, **then** S is a subtype of T."

and adding Occam's Razor

> "Make class S a chile of class T **if and only if** you intended to pass an objet o1 of types S as the argument to some function P defined in terms of T"
> if you don't intend to do that, there is no reason for that public inheritance relationship to exists,... and therefore that relationship should **not** exists.

Chesterton speaking against unnecessary changes and the mindset of 'modern reformers' (someone who does reforms for the sake of refroms).

> "The modern reformer says "I don't see the use of this fence; let us clear it away". The more intelligent type answers, "When you can tell me that you **do** see why it is here, **then** maybe i will allow you to destroy it"". \
> --G.K. Chesterton(1929), lightly abridged
>
> Since fences generally have reasons, tearing down fences should not be done lightly.

so if we see classical inheritance, we shouldn't change it (in a refactor) until we see why it was done this way in the first place.

Robert Frost

> "Before I build a wall I'd ask To know\
> What I was walling in, or walling out"\
> --Robert Frost, "Mending Wall" (1914)

before we put up a fence, we should know what we're doing, the reason for it, and we should document it clearly, so if we come across it in the future, we can rationally consider if it's safe to remove in the current situation.

otherwise, we might run into 'The paradox of the useless fence'./

- before we tear down a fence, we must understand why it's there.
- if there was no reason to build the fence, it will be hard to understand why it was build.
- therefore: it's harder to remove a fence that was build for no good reason than a fence that was built for good reasons with a sound rationale.

and this is a thing that we can see in many codebases. somebody writes a code that uses a technique without a good reason, and then we can't remove the code because we can't understand what they were trying to do.

in c++, when we see inheritance, we expect to see a reason why it was designed this way, and specifically, we expect to see someone using a polymorphic method. if we aren't "forced" into inheritance, we should avoid it. **Prefer composition over inheritance (Has-A is better than Is-A)**.

### Naming and STL Examples

A single name for a single entity:

> - We should use different words to refer to different ideas.
> - When refering to the same idea, we should use the same word.
> - Any given single identifier should refer unambiguously to a single entity.

two codebase, which is easier? A uses the same name (diffrent signature) for two functions. B uses different names.

```cpp
//A
bool feed(Snake& snake);
bool feed(Bear& bear);
```

or

```cpp
//B
bool feed_snake(Snake& snake);
bool feed_bear(Bear& bear);
```

using the specialized name helps us detect and trace, we can always find all the usages, jump to it, rename it, and we can always tell if which function is used. it help the computer with overload resolution, and makes it easier for the IDE.

so if we see the version A (with the overloads of the identical name), we expect that there was a reason for this, and we should actually expect a specific reason - polymorphism.

polymorphism isn't just virtual functions, there's also static polymorphism of templates.

```cpp
template <typename T>
void solve_puzzle(Animal& a)
{
    feed(a); //calling a specific overload.
}
```

both std::vector and std::list (and many other containers) use the identifier "_.push_back()_" as a method name. this same name allows us to create a template function. like the _std::back_inserter_ iterator, _std::swap_.

```cpp
template <typename T>
struct back_insert_iterator{
    //...
    // container is T*
    back_insert_iterator& operator=(const T::value_type& x)
    {
        container->push_back(x);
    }
};
```

if there was no use of polymorphism, a unique identifier would be easier to read, understand and maintain, but we get so much functionality from the STL,which makes the overloaded versions preferabl.

a counter example from the STL, _erase_ has two overloads. one identifier with two entities. Arthur says that this code doesn't facilitate any polymorphism.

```cpp
class vector
{
    using CI = const_iterator;
    iterator erase(CI first,CI last);
    iterator erase(CI where){
        return erase(where,where+1);
    }
};
```

here is an example where we trip over ourselves, we have a vector of numbers, we want to keep only the even numbers. we use the erase-remove idiom but we forget to pass the second argument to _.erase()_, so we erase only one element.

```cpp
bool isOdd(int);
std::vector<int> v= {1,2,3,4,5,6,7};
v.erase(std::remove_if(v.begin(),v.end(),isOdd)); // erase remove idiom, erase with one arguments
static_assert(std::none_of(v.begin(),v.end(), isOdd)); // this fails!
```

what we should have done is

```cpp
v.erase(std::remove_if(v.begin(),v.end(),isOdd),v.end()); // erase remove idiom, erase with two arguments
static_assert(std::none_of(v.begin(),v.end(), isOdd)); // now it's ok
```

why was the overload created? arthur says there isn't a good reason.

### An Issue with the Constructor

STL classes have too many overloads, especially std::string,

```cpp
class string {
    string(size_t n,char); // string with n times of char
    string(const char * ,size_t n); // first n chars of char*
    string(const string &,size_t pos); // copy of other string, starting at some position
    template<InputIterator It>
    string(It,it); //take two iterators
}

size_t zero =0;
auto a =std::string(zero,0); //what is called here? zero instance of character 0
auto b =std::string(0,zero); // calls the overload with the const char*, undefined behavior probably
auto c = std::string("abcd",2);  // "ab" constructor first n chars,
auto d = std::string("abcd"s,2); // "cd" constructor copy of other string from position, just because we added the string literal
```

could all these constructor be replaced with factories?

```cpp
class stringRevised {
    static stringRevised fromCopiesOfN(size_t n,char); // string with n times of char
    static stringRevised fromPtrAndLength(const char * ,size_t n); // first n chars of char*
    static stringRevised fromSuffixStartingAt(const stringRevised &,size_t pos); // copy of other string, starting at some position
    template<InputIterator It>
    static stringRevised fromRange(It,it); //take two iterators
};
size_t zero =0;
auto a =stringRevised::fromCopiesOfN(zero,0);
auto b =stringRevised::fromPtrAndLength(0,zero);
```

we couldn't do this, constructors are special.
factory functions are self documenting and easy to understand, but they don't work with the perfect-forwarding wrappers.

- _std::make_shared_, _std::make_unique_
- _emplace_back_, _optional::emplace_, _variant::emplace_

```cpp
auto a1 = std::make_shared<std::string>(zero,0);
auto a2 = std::make_shared<stringRevised>(stringRevised::fromCopiesOfN(zero,0)); //extra move operation in the good case, copy also possible.
```

constructor syntax allows us to create objects not on the stack in a comfortable way. we can actually 'new auto' (c++17) to heap allocate a factoy function p-rvalue result, gurantess heap ellision,actually good

```cpp
T t1 =T(1,2);
T* p1 =new T(2,3);
T t2 = T::fromTwoInt(3,4);
T* p2 = new auto(T::fromTwoInt(4,5)); //this works!
```

could we make a generic perfect forwarding function with factory functions?
something like this? this would work, but now instead of having a single identifier for many entities as the constructor, we simply have to choose a different name that all the classes are going to use and it won't be informative

```cpp
template <typename T, typename... Args>
auto build_shared(Args...args)
{
    T* p= new auto(T::createGenerically(args...));
    return std::make_shared<T>(p);
}
```

our fantasy: could we pass the creation format itself? pass in the factory function itself? in today's c++, this must be a concrete set (not overload set). there is one proposal for "lifting" an overload set into a concrete lambda object. a different proposal for an object that deduces types from an overload set(std::overload_set, like std::initializer_list), some sort of compiler magic.

```cpp
template <typename How, typename... Args> //the class 'How' is the problem
auto build_shared_How(How how,Args...args)
{
    auto *p= new auto(how(std::forward<Args>(args)...));
    return std::shared_ptr(p);
}
std::shared_ptr<stringRevised> sp1 = build_shared_How(stringRevised::fromCopiesOfN,0,0);
std::shared_ptr<stringRevised> sp2 = build_shared_How(stringRevised::fromPtrAndLength,0,0);

//proposal 1,
//auto sp3 = build_shared_How([]stringRevised::fromCopiesOfN,0,0);
```

### Mental Models, Macros, Polyfills

to recap, sometimes we give two entities (in different classes) the same name with the same signature, because we are going to template on the class type. this is what _std::make_shared_ does (with perfect forwarding)

```cpp
template<class Animal>
void foo(Animal & a)
{
    a.feed();
}
```

sometimes we give the same name but different signatures, because we're going to template on the argument types.

```cpp
template <class Animal, class... Foods>
void bar(Animal &a, Foods... foods)
{
    a.feed(foods...);
}
```

all STL containters provide _c.insert(pos,value)_, associative containers (like std::set) ignore the positional value. this allows us to create an _std::inserter_ with the same arguments for all containers.

```cpp
std::vector<int> data ={1,2,3}

std::vector<int> c1;
std::copy(data.begin(),data.end(),std::inserter(c1,c1.begin()));
std::set<int> c2;
std::copy(data.begin(),data.end(),std::inserter(c2,c2.begin()));
```

inserting into a set doesn't always make it bigger, if the set contains the element, it just returns it. the mental model of inserting into a set is different.
should all insert functions have the same name? why not _insertAt(pos,x)_ ,_insertNodeHandle(nh)_, _insertRange(it1,it2)_.

STL provides uniformity of containers, all containers share the same API, we can switch from _std::vector_ to _std::deque_, _std::list_ or even _std::multiset_, but does it work work the same?

no. the behavior is different. _.push_back()_ on _std::deque_ maintains the iterators, but not on a _std::vector_, _.push_back()_ invalidates the iterators (the vector might have be reallocated).

```cpp
//std::deque<int> data = {3,1,4,1,5,9,2,6,5}; //replace deque for vector
std::vector<int> data = {3,1,4,1,5,9,2,6,5};
std::sort(data.begin(), data.end());
auto [first,last]= std::equal_range(data,begin(),data.end());
data.push_back(100); // invalidates iterators in vector
data.erase(first, last); // undefined behavior in vector
for (int i: data)
{
    std::cout << i << '\n';
}
```

Can templates be mental?

> "Software engineering is programming integrated over time"\
> -- Titus Winters.

Sharing names as upgrade paths?
std::string*view and std::string share the same names for many functionalities, it was done in purpose. this was done so we could upgrade the std::string to std::string_view without issues. this was done with \_std::optional*, it has the same operators as the smart pointer classes. the reasoning was that we could replace _std::unique_ptr_ with _std::optional_, this way we reduce heap allocation and still get the 'not created' option.

reusing names can still lead to bugs. in this example both _std::optional_ and the inner type have _.reset()_ method, if we use it with the dot notation, we call the _std::optional_ method, the arrow notations is for the inner type. this would happen also with a smart pointer.

the code compiles and runs, but it doesn't do what we think it does!

```cpp
struct DataCache{
    void update(key,value);
    void reset();
};
struct Connection
{
    std::optional<DataCache> dataCache_;
    void resetCache()
    {
        if (dataCache_) //if optinal value exists
        {
            dataCache_.reset(); //oops! bug! not we don't have a cache at all.
            //dataCache_->reset(); //this is what we wanted!
        }
    }
};
```

the STL and boost libraries also try to have the same names for the sake of upgrade paths.it's not a template metaprogramming, more of a **macro based static polymorphism**. the API was designed to allow this behavior. it's also called **polyfill**. the boost version is a _polyfill_ for the std version.

```cpp
#if __cplusplus >= 201703L
#include <optional>
using std::optional;
#else
#include <boost/optional/hpp>
using boost::optional;
#endif
```

we can also use this from compiler flags as platform specific polymorphism.

```cpp
namespace curses
{
    void clearScreen();
    void drawAt(int x, int y, char ch);
}

namespace conio
{
    void clearScreen();
    void drawAt(int x, int y, char ch);
}

using namespace TERMLIB; // -DTERMLIB=curses or -DTERMLIB=conio
void drawTitleBar()
{
    for (int x =0; x< 100; ++x)
    {
        drawAt(x,1,'#');// calls different function according to the TERMLIB flag.
    }
}
```

### Takeaways and Mantra

if the default parameters isn't used, don't use it. it's like an overload set, check if it's justified to use.

concepts are constrains on types, but we define them based on the algorithms, we define things based on usage.

std::enumerators - template specialization on enums that have the same name.

> - Inheritance is for sharing an interface.
>   - and so is overloading
> - Use a single names for a single entity
> - When you see two things with the same name, assume there is a reason for it.
> - When you have option to give two things the same name, **don't, unless** there is a reason for it.
> - To find concepts, don't study what your callees provide in common; study what your callers require
> - Default function arguments are the devil.

</details>

## C++11/14 at Scale: What Have We Learned? - John Lakos & Vittorio Romeo

<details>
<summary>
How we teach C++ versions, which features are suitable for which skill level? what should be taught and when.
</summary>

[C++11/14 at Scale: What Have We Learned?](https://youtu.be/E3JG2Ijjei4)

They're are publishing a book later this year: **Embracing Modern C++ Safety**

> - Why are we talking about C++11/14 in 2021?
> - How C++11/14 an surprise you today
> - C++ at scale
> - "safety" of a feature
> - case study: extended _friend_ declrations

### Why are we talking about C++11/14 in 2021?

adoption rates of C++ standards, some projects are still lagging behind and haven't adapted the newer standads yet, even in 2021 there are places where C++11/14 are just being adopted.

> - There are great learning resources
>   - But most teach "the features" rather than "the experience"
>   - What looks good on paper might not work in the "real world"

not just what the new features are, learn when and how to use them, how they operate inside the bigger context.

### How C++11/14 an Surprise You Today

> Q: "What is the smallest change to the core language you can think of in C++11?"

c++11 changed how double ">" behaved. before c++11 ">>" was parsed as a right shift, so a space was needed to make this recognizable as closing a nested template. in c++11 things were changed and ">>" was now somthing else. this means that a valid c++03 code is invalid in c++11.

in this example, c++03 will see `256 >> 4`, but c++11 will reject this code.

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<256 >> 4> smallBuffer;
```

to fix this issue, we simply wrap the right shift expression in parentheses

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<(256 >> 4)> smallBuffer;
```

in this example, c++03 returns 100, while c++11 returns 0; the compiler gives a warning.

I think that c++03 treats this as a sequence of enums and nested enums (we can change the final 'a' to 'c' and get 102, but not to 'b'). and c++11 treats this as some comparions thing.

```cpp
enum Outer{a=1,b=2,c=3};
template <typename>
struct S {enum Inner {a=100, c=102};};
template <int>
struct G{typedef int b;};
int main()
{

    return S<G< 0 >> ::c>::b>::a;
}
```

other stuff: every one of those can have a dark side.

> - Attributes that can make you code ill-formed NDR.
> - 'extern templates' not improving compilations time or code size at all?
> - Destruction order UB with meyers singletons.
> - Encoding of white space withing raw string literals.

> "NDR - No Diagnostic Required"

### Modern C++ at Scale

how do we teach modern c++? what to prioritize, what kind of approach? how do we integrate the new features into the company style guide? what if we have a tool chain for the style guide? how do we communicate these changes to other teams?

### "Safety" of a Feature

> - Every features of c++ is "safe" when used correctly.
> - But what is the likelihood that it is used correctly?
> - Does the feature have any "attractive nuisance"? (does it invite misuse?)
> - What are the advantages of using a feature compated to its risks?
> - Is it worth teaching to a new hire? to an expereinced hire?

from the book:

> "The degree of safety of a given feature is the the relative likelihood that the widespread use of that feature will have positive impact and no adverse effect on a large software company's codebase."

how likely is teaching and implementing a feature is to go smoothly, be used correctly and give good results, as opposed to being hard to teach/understand, prone to create oppertunities for bugs, hard to maintain by inexperienced workers, or have small scale effects on performance.

three categories of safety:

> - **Safe**
>   - Adds considerable value, easy to use, hard to misuse.
>   - Ubiquitous adoption of such features is productive.
> - **Conditionally Safe**
>   - Adds considerable value, but prone to misuse.
>   - Require in-depth training and additional care.
> - **Unsafe**
>   - Provide value only in the hands of an 'expert', prone to misuse.
>   - Wouldn't teach these as part of genereal c++11/14 course.
>   - Require explicit training on their use cases and pitfalls.

the 'override' keyword is a 'safe' feature. it prevents bugs, makes code self-explantory, and has no real technical downsides.
the one problem that can happen is that people overrely on it, and people except this feature as the norm, and forget that this is just a bonus, you can still have overriding methods without this keyword.
(we can use compiler flags '-Winconsistent-missing-override' and '-Wsuggest-override', but they aren't perfect).

```cpp
class MockConnection : public Connection
{
    void connect(IPV4Address ip) override;
};
```

the 'auto' keyword is _conditionally safe_, it can be great, but can also cause readability problems, and can introduce bugs. _Range based for loops_ are great, but they also have the possibility of bugs, and therefore are marked _conditionally safe_.

in this example, we are actually ok, because we return the vector by value and get lifetime extentsion. this is not true once we decide to be smart and return the vector by reference. now we don't have lifetime extension.

```cpp
class TriggerGetter
{
std::vector<Combo> getCombos() const; //no Issues
const std::vector<Combo> & getCombosRef() const; //oops
};

for (Combo& c : keyBoardTriggerGetters[bindID]().getCombos()) //return by value
{
    //..
}
for (Combo& c : keyBoardTriggerGetters[bindID]().getCombosRef()) //return by const reference
{
    //..
}
```

this is something we overcome in c++20 with init statements, but we must be aware of this issue. it's not an entirely safe action.

_decltype(auto)_ is a strong feature, but it's often misunderstood, misused, a requires training to use correctly, it should be defined as _unsafe_, and only be used with carefull consideration in the codebase. it allows us to deduce the return type from the expression, doesn't strip away qualifiers, returns value or reference objects,and doesn't change anything.

example: higher-order functions.

```cpp
template <typename F>
decltype(auto) logAndCall(F&& f)
{
    log("invoking function ", nameOf<F()>);
    return std::forward<F>(f)();
}
```

but if we teach '_auto_', '_decltype_' and '_decltype(auto)_' together, we push people towards overusing '_decltype(auto)_'.

> - some misconceptions:
> - "If _decltype(auto)_ does everything _auto_ does and more, why not use it all the time?"
> - "If _decltype(auto)_ is more flexible, wht no use it when I'm not sure when to choose between _auto_ and _auto&_?"

> understanding _decltype(auto)_ requires:
>
> - Having a solid grasp on type inference and value categories.
> - Being somewhat experienced with using _auto_ and _decltype_.
> - Having some metaprogrammin expereice.

if you have just learned about _auto_ and _decltype_, you probably aren't in the right level to use _decltype(auto)_ yet.

plus, _decltype(auto)_ has some issues with parentheses surronding it, it's not an easy thing to understand, it can effect SFINAE behavior, so it's not the allways the best tool for the job.

safe features: attributes (most of them), _nullptr_,_static_assert_, digit seperators.
conditionally safe: _auto_, _constexpr_, _rvalue_ references
unsafe: _\[\[carries_dependency]]_, _final_, inline namepaces.

when we teach a new version of c++, we should:

> - Teach _safe_ features early and quickly
>   - Most of them are quality-of-life improvements or hard to misuse.
>   - Trust the student
> - Teach _conditionally safe_ features by building on top of _safe_ knowledge
>   - They require more time and examples.
>   - Show how the can backfire.
>   - Have exercises that make student question whether to use a feature or not.
> - Leave a subset of of _unsafe_ features for self-contained CE courses
>   - E.g. "Library API and ABI version with the 'inline namespaces'"

### Case study: Extended _friend_ Declrations

> - Prior to C++11, _friend_ declareations require an 'elaborated type specifier'.
>   - _elaborated type specifier_: Syntitical element having the form of \<class|struct|union> \<identifier>
> - This restriction prevents other entities to be designated as friends.
>   - E.g. type aliases, template parameters.
> - A surprsing behavior with namespaces.
>   - it wasn't possible to refer to a entity in the global namespace, a new entity was being declared instead.

```cpp
//C++03 friend
struct S;
struct Example
{
    friend class S; //ok
    friend class NonExistent; //ok, even it this class doesn't exist.
};

using WindowManger = UnixWindowManager;

template <typename T>
struct Example2
{
    friend class WindowManger; //error! type alias
    friend class T; //error! template parameter
};

struct SA; //this SA is in the global namespace
namespace ns
{
    class X3
    {
        friend struct SA; // ok, declares a new ns::SA class instead of refereing to the global ::SA
    };
}
```

> C++ 11 extended 'friend' declarations lift all the aforementioned limitations. and fixes the weird behavior of creating types. we don't need the class|struct|union specifier anymore.

```cpp
//C++11 friend

Struct S;
typedef S Salias;
using Salias2 = S;

namespace ns
{
    template <typename T>
    struct X4
    {
        friend T; //ok
        friend S; //ok, refers to global ::S
        friend SAlias; //ok, also refers to global ::S
        friend Salias2; //ok, also refers to global ::S
        friend decltype(0); //ok, same as 'friend int'
        friend C; //error! 'C' does not name a type.
    }
}
```

> so why is this feature categorized as _unsafe_?
>
> - It is rarely useful in practice, like c++03 _friend_
> - Promotes _long-distance friendship_!
>
> When a type 'X' befriends type 'Y' which lives in a separate component...
>
> - 'X' and 'Y' cannot be thoroughly tested independently anymore.
> - Physical coupling occurs between 'X' and 'Y' components,
> - Possible physical design cycles can happen.

if we have too many friends, it might be a symptom of a design problem, having friends from diffrent namespaces means more coupleing and less modularity. but even if it's _unsafe_, it does have it's benefits, like helping us spot typos when declaring friends.

other intersting points: type alias customization points, PassKey idiom...

and we will focus on CRTP - curiously recursive template pattern.

base knows who it derives from, thanks to T. usefull to implement _mixins_ and factor out copy-pasted code.

```cpp
template <typename T>
class Base
{

};

class Derived : public Base<Derived>
{

};
```

example use case, having a counter for classes creations.

```cpp
//header file
class A {
    static int s_count; //decleration
    public:
    static int count() {return s_count;}
    A(){++_count;}
    A(const A&) {++s_count;}
    A(A&&) {++s_count;}
    ~A() {--s_count;}
};

//defintion file
int A::s_count;
```

we can factor out the counter behavior, using the protected access modifier. (it's a mixin, whatever that means).

```cpp
template <typename T>
class InstanceCounter
{
    protected:
    static int s_count; //declaration
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)
```

we can then use in other classes

```cpp
struct A :InstanceCounter<A>
{
    A() {++s_count};
    //also add this for the destructor
};

struct B : InstanceCounter<A> //oops, made a typo! we will use the same counter.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //oops, we messed with the entire tree!
}
```

actually, this is something we could use the friend declerations! we move from 'protected' to 'private', and make T a friend class of the counter. now only the class that declared the counter can access it, not others classes and not derived.

```cpp
template <typename T>
class InstanceCounter
{
    private:
    static int s_count; //declaration, private
    friend T; //only T can access us.
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)

struct B : InstanceCounter<A> //error, s_count is private within this context.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //error, s_count is private within this context
}

```

the crtp allows us avoid boiler plate code, this is also an example of using inheritance without virtual functions. this is also a case where we don't want to use 'final'.

### when to use 'final'

do we really know a class shouldn't be inherited from? are we sure.

it's okay if we have a class that is supposed to behave like a primitive.
but EBO (empty base optimization) doesn't play nice with 'final'.

### Conclusion

> - The "Human cost" of a feature is not easy to quantify.
> - Categorizing features by "safety" helps with devising learning paths.
> - All features have good use cases and nasty pitfalls.

the book will be out in the future, check [this page](https://vittorioromeo.info/emcpps.html)

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

## Variations on Variants - Roi Barkan

<details>
<summary>
Variants, Unions, proposals for new kind of variants, standard-layout considerations.
</summary>

[Variations on Variants](https://youtu.be/YBXRiPKa_bc), [slides](https://docs.google.com/presentation/d/1W0QBblWpJ-AXsdo70kvtxg5G2RZLiqb-mewQKySTf6s/preview?slide=id.p)

> - What are variants
> - Variant vs unions
> - Intrusive variants
> - Streams of variants
> - Variants for de-virtualization

### Variants Introduction

std::variants, a typesafe union

> - CppReference.com: "the class template std::variants represent a type-safe union",
> - boost.org: "the variants class template is a safe, generic, stack-bases discriminated union container"
> - plain english: "a union that knows (holds) its type".

we can use std::get\<type> accessor to get the memberm or use std::visit.\
sum type, as opposed to product types
memory layout, the std::variant has some extra memory requirements for the tag.

> usages
>
> - State machines
> - Value Semantics for Dynamic Types
>   - commands
> - Success / Fail
>   - expected\<T>
> - Exist / void
>   - optional\<T>
> - Runtime Dispatch (polymorphism)
> - Pattern Matching

(roi looks at some old lectures from different people)

std::optional is a specialzed form of std::variant. \
runtime dispatch can allow us to avoid virtual function call.\
pattern matching is another way for dynamic, trying to do something similar to a switch case, perhaps using std::visit.

extra reading:\
[std::visit is everything wrong with modern C++](https://bitbashing.io/std-visit.html)

pattern matching vs concepts/contracts:\
concepts - compile time sfinae.\
contracts - run time assertions (not existing yet, but soon)\
pattern matching - run time advanced switch case (not existing yet, low priority), might have 'inspect' keyword with some new mechanics, combining values, lambda, dynamic values, different types, etc...

_inspect() in a nutshell_
switch + std::visit() + [structured, bindings]

example of using pattern matching for balancing a red/black tree.

### Variants vs Unions

the tag in the variant is private. only the constructor and assignemt operations can change the flag. the compiler knows in advance all the possible alternative forms.

this is a buggy code:\
missing break statements (fall through), missing cases and no default, calling the wrong function.

```cpp
union IdentityCard
{
    IDNumber nationalID;
    PassportNumber passport;
    UUID factoryCertificate;
};
enum IDType {CITIZEN,TOURIST,ROBOT};
void checkID(IdentityCard id, IDType type)
{
    switch(type)
    {
        case TOURIST: checkPassport(id.passport);
        case CITIZEN: checkPassport(id.passport);
    }
}
```

if we used variants, it would look like this. the compiler must account for all the cases, it's not just a warning if one is missing. we aren't handling the tag.

```cpp
using IdentityCard = std::variant<IDNumber,PassportNumber, UUID>;
void checkID(IdentityCard id)
{
    std::visit([](auto& x){x.check();},id);
}
```

the example for the C style code was a strawman, real union code looks more like a variant, with a struct that holds the tag (explicit header). or as a union with structs that each have the same header part (implicit).
the explicit type looks more organized and less crowded, but the implicit type allows direct access to the header, if we need data from it and we got a pointer/reference to the inner member data.

```cpp
struct IdentityCardExplicit
{
    IDType type;
    union value {
        IDNumber nationalID;
        PassportNumber passport;
    } value;
};

union IdentityCardImplicit
{
    struct Header {
        IDType id;
        };
    struct Citizen{
        IDType id;
        IDNumber nationalID;
        };
    struct Tourist{
        IDType id;
        PassportNumber passport;
        };
}
```

the header field can contain more than the tag itself, and we use a macro to define them.\
_(I don't like this, it feels like some semi-colons are missing)_

```cpp
#define HEADER_FIELDS \
IDType type; \
Date expiration;

struct Header { HEADER_FIELDS};
struct Citizen{
    HEADER_FIELDS
    IDNumber nationalID;
    };
struct Tourist{
    HEADER_FIELDS
    PassportNumber passport;
    };

union IdentityCardImplicit
{
    Header header;
    Citizen citizen;
    Tourist tourist;
}
```

> Keeping the C Layout
>
> - Header with type is common.
>   - Network protocols- TCP/IP, Finance, etc...
>   - File formats - ELF, etc...
>   - Serialization - [Cap'n Proto](https://capnproto.org/), [Apache Avro](https://avro.apache.org/)
> - C layout is important.
>   - Compatibility with existing code.
> - Goal - Be safer than C, keep the layout.
>   - Sacrifice some safety.

some protocols use a specific layout that we can't change. we must remain compatible we legacy code.\
the c++ standard says that tier are cases when unions are good, it makes standard layour classes possible.

> - "Standard-Layout classes are usefull for communicating with code written in other programming languages"
> - various constraints [StandardLayoutType](https://en.cppreference.com/w/cpp/named_req/StandardLayoutType)
>   - no virtual functions or virtual base classes
>   - single access control - all (non static member are the same, no mixing between public/protected/private.
>   - all non-statics in the same class - all none static members live in the same class (this one or a base class).
>   - and more.
> - [std::is_standard_layout](https://en.cppreference.com/w/cpp/types/is_standard_layout)
>   ```cpp
>   static_assert(std::is_standard_layout<Citizen>::value,"not standard layout"); //c++11
>   static_assert(std::is_standard_layout_v<Citizen>); //c++11
>   ```
> - Layout-compatibility allows accessing members without knowing the type
>   - "In a standard-layout union with an active member of struct type T1, it is permitted to read a non-static data member m of another union member of struct T2 provided m is part of the common initial sequence of T1 and T2"
>   - [std::is_corresponding_member](https://en.cppreference.com/w/cpp/types/is_corresponding_member)
>   ```cpp
>   static_assert(std::is_corresponding_member(&Header::type, &Citizen::type)); //c++20
>   ```

### Intrusive variants

the a variant, but making the tag visible, and therefore layout compatible. a suggestion, using the [offsetof macro](https://en.cppreference.com/w/cpp/types/offsetof) to either create something similar to the union of struct (implicit headers)

```cpp
using ID_intrusive = intrusive_variant<
IDType, offsetoff(IDHeader, type), //where is the tag
intrusive_variant_tag_type<IDType::CITIZEN,Citizen>, //add options per tag
intrusive_variant_tag_type<IDType::TOURIST,Tourist>>; //

//will become something like this
union IdentityCard
{
    struct IDHeader {
        const IDType id;
        };
    struct Citizen{
        const IDType id;
        IDNumber nationalID;
        };
    struct Tourist{
        const IDType id;
        PassportNumber passport;
        };
};
```

a different approach will be using static constexpr tags.

```cpp
using ID_intrusive = intrusive_variant<
IDType, offsetoff(IDHeader, type), //where is the tag
intrusive_variant_type<IDType::Citizen>, //add types
intrusive_variant_type<IDType::Tourist>); //

//will become something like this
union IdentityCard
{
    struct IDHeader {
        const IDType id;
        };
    struct Citizen{
        const IDType id;
        static constexpr IDType TAG = IDType::CITIZEN;
        IDNumber id;
        } citizen;
    struct Tourist{
        const IDType id;
        static constexpr IDType TAG = IDType::TOURIST;
        PassportNumber passport;
        } tourist;
};
```

> - the user dictates the type and location of the tag.
> - visit() will still be O(1).
>   - potentially larger lookup table.
> - Customization Point for tag deduction.

we can have Different Approaches to get the Tag, even if we don't know what the type is, standard layout can have all private members, or the tag might be calculated somehow

> - Offset of the field in the object
>   - getTag<IDType>(hdr, offsetof(Hdr, tag));
>   - getTag<IDType>(hdr, std::integral_constant<size_t, offsetof(Hdr, tag)>());
> - Pointer to the field
>   - getTag<IDType>(hdr, &Hdr::tag);
> - Call a member function
> - getTag<IDType>(hdr, &Hdr::getTag);
>   - Useful when the tag is private.
>   - Useful when the tag is calculated.
> - Call a free-function / lambda.
>   - getTag<IDType>(hdr, [](const Hdr& h) { return h.tag; });
>   - Less intrusive

possible implementations, using c++20 concepts (either with reinterpret_cast or with std::invoke).

c++20 can give us extra type safety, with possible validations and without explicitly stating the offset.

```cpp
using ID_intrusive = decltype(decl_safe_intrusive_variant(
    &Citizen::type, IDType::CITIZEN,
    &Tourist::type, IDType::TOURIST));
```

intrusive variant will have a safe visit() function, but it still has place for bugs and still requires boilerplate code. class hierarchies could help us, and we will need some help from the utilities functions, but adding base classes with data breaks the standard-layout specification.

(clip from sean Parent)

trying to show a example of 'variant_of_base',base class with data and then a variant that knows to use the derived classes with better safety, but still not standard conforming.

### Streams of variants

sending arrays of variants,comparing std::variant, intrusive variant, and how it's done it the real world (send tag, then struct, so only use the ram we need, trying to minimize waste).

helper with arrays of variants, like a forward iterator, a container that like a special queue for the variants. jumping between elements.

Summary so far:

> - variants are different than unions.
> - real-world unions already have tags (and headers).
> - Intrusive_variant - C++ safety with high C compatibility.
> - Variant_of_base - add classes to your code.
>   - Not standard-layout, undefined behavior.
>   - Perhaps we should widen the rules, add \[\[standard_layout]] attribute?
> - condensed_variant - real world streams of binary data.

### Variants for de-virtualization

virtual dispatch- polymorphism.
de-virtualization - trying to get virtual function to run just as fast as non virtual functions. a talk from 2013 about creating our own virtual table because virtual functions are pretty wastefull, especially for small hireachy.
the problem with virtual functions is **Branch MisPredictions**, but compilers and processors get better with branch prediction over the years, the compiler only sees the code, but the processor sees the data so it can re-arrange the date by itself.
compilers can use PGO (profilers guided optimizations) or LTO (link time optimization), or the \[\[likely]] attribute. this allows the compiler to create code that is better suited for performance.
we want to inline, rearrange and inspect functions code, and virtual functions aren't great for that.

we might be able to use variant and std::visit() to get better de-virtualization. if we have a different implementaition of visit (without a jump table), we could get a much better performace.

</details>

## Windows, MacOS and Web: Lessons from Cross-Platform Development @ Think-Cell - Sebastian Theophil

<details>
<summary>
Challenges for cross platform code.
</summary>

[Windows, MacOS and Web: Lessons from Cross-Platform Development @ think-cell](https://youtu.be/Cmud1jO__VA)

they started with a library that was developed in windows environment,it was a plug-in, and therefore, dynamically loaded and not in control of the entire process, many shared resources.
they

> "need a cross-platform toolkit that hides platforms specifics and **behaves identically** on different platforms"

(if such things can exists)

> Agenda
>
> 1. Levels of Abstraction: Hiding Platform Specifics
> 2. Kernel Object Lifetimes: Interprocess Shared Memory
> 3. Common Tooling I: Text Internationalization
> 4. Common Tooling II: Error Reporting
> 5. Moving to WebAssembly

### Levels of Abstraction: Hiding Platform Specifics

platform independent c++?
there are easy cases, like rendering, http requests (with the system API), child process and setting IO pipes. theses cases can be

> "Clearly defind as '**data In, data Out**'"

but even these cases can be difficult to make true platfrom indpendent, like direct call to rename/move files, which has different behavior flag for windows and macOs.

consider what the function really does and what it needs, what is the purpose of the function? if we know the "Why" - the reasoning for the function (what the user tries to achieve), we can tailor the "How" - what do we call in each platform. we don't simply route the arguments to the OS system call.

creating a file that is automatically deleted by the OS when the system closes (even at crush) but while it's alive it can be used by other processes. this behavior can be easily down on windows, but not on Mac, so maybe we need to rethink the 'how', and use a sqlite database for this in macOS, rather than file.

> - cross platform interfaces need to have well-defined, strong semantics.
> - weak semantics lead to subtle errors.
>   - Warning Sign: Having to look at the implementation.
> - Strong semantics increase DoF (degrees of freedom) for the implementor.
> - Too high-level.
>   - missed chance to unify code. Rare, we are lazy.
> - Too low-level.
>   - You'll force identical interfaces on very different things.
>   - semantics don't match operating system (_QFile::setPermissions_).
>   - or you'll loose a lot of expressiveness (_rename_).

### Kernel Object Lifetimes: Interprocess Shared Memory

boost and other libraries solve some of the problems for us, but sometimes we can to better.

boost offere interprocess communication tools, different shared memory behavior for windows and mac, windows cleans up, Posix can keep files alive for a long while. there are Robust Mutexes, file locking.

### Common Tooling I: Text Internationalization

a tool for text internationalization: translating, numbers formats.
text, context, plurality forms, what we wants.

[Boost.Locale] (https://www.boost.org/doc/libs/1_51_0/libs/locale/doc/html/main.html) was added in 2018 (boost 1.67), which supports tranlation by creating a catalog of transaltion, in boost it's runtime, in their implementation they try to make it constexpr. we don't want to read a file from disk, it's dangerous, we rather link the translations as part of the program.

> reminder about constexpr

strong and identical semantics can also refer to external tools in the build process.

### Common Tooling II: Error Reporting

dumping stack data to file, different for windows and Unix. they have an error report system that sends error to the backend and tries to identify the error. but file formats for dumps are different, and it needs to be standardized.
macOS allows to send access permissions to other processes.

### Moving to WebAssembly

the products ships with chrome extensions and webapp. they tried to use TypeScript (not JavaScript). but they weren't able to share data with c++. using C++ was typeunsafe because it lacked wrappers for JavaScript. so they built something of their own.
it's called 'Defiantly typed", and they have 'typescriptem'. which creates type safe c++ that does JavaScript.

in typescript, decleration order doesn't matter. so there needs to be some dependency list. typescript has non-integer enums, so they created a marshal enum template, and they had to create function callbacks.

</details>

## Algorithms from a Compiler Developer's Toolbox - Gábor Horváth

<details>
<summary>
A bit of compiler algorithms for optimization, using identities and algebra.
</summary>

[Algorithms from a Compiler Developer's Toolbox](https://youtu.be/eeS1WP7FK-A),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/CompilerAlgorithmsTalk.pdf)

### Why Study Compiler?

there are many algorithms and data structure that are used in compilers, compilers are everywhere, like web browsers (html+css, svg, JavaScript of course), GPUs also have compilers, databases have compilers and optimizers, even some configuration format files have something like a compiler inside them. python compiles down to something for machine learning, and routers and modem have something running on them.

A lot of opportunities to improve code, any small improvement is multiplied because it effects every program compiled by it. if we improve a low level compiler (like c++), then we also effect any compiler that uses it (like python or JavaScript).

example: loop strength reduction.

video of a talk by matt godbolt. replacing sum by loop with sum by formula

$
\sum x \equiv \\\frac{x (x+1)}{2}
$

playing with loops kinds and looking at the assembly, we see that the compiler manages to remove the loop and figures out a closed-form formula.

but floating point messes up the optimizations, floating point arithemetic.

### What's Inside the Compiler?

Math.😅\
Chains of recurrences - recursive function when the increment is also a recursive function.

two kinds of recursive formula notations. making functions at incrementoars

$
f(i*i(n)) = {initial,+,incrememt }
$

algebra, operations identities. making loops into recursive notations with those identities, sub expressions combine together. turning this

```cpp
for(i=0;i<m;++i)
v[i] = (i+1)*(i*1) -i*i -2*i;
```

into a constant expression with identities.

$(x+y)(x+y) = (x+y)^2 = x^2 +2xy + y^2$

so we open up the identity, and we can then cancel out stuff and reach a constant.

```cpp
for(i=0;i<m;++i)
v[i]=1;
```

arithemetic series which are supposed to be loops can be made into closed formulas or at least have much less operations per loop

```cpp
int t[20];
for (int i=0;i< 20; i+=1)
{
    t[i]=(i+1)*(i+1) + 3*i - 5; // four additions, two multiplications
}
```

is transformed into this compact form with only two additions.

$
f_{(i+1)^2+3i-5}(n) = \\
\{-4,+6,+2\}
$

which is equivalent to writing this c++ code.

```cpp
int t[20];
int a = -4;
int b = 6;
for (int i=0;i< 20; i+=1)
{
    t[i]=a;
    a+=b;
    b+=2;
}
```

an example of how clang does it. we take this code into a file.

```cpp
int f(int num)
{
    int result =0;
    for (int i =0;i<num;++i)
    {
        result += i;
    }
    return result;
}
```

and then run the following command on it (replace $1 with file name)

```bash
clang++ $1.cpp -c -02 -Xclang -disable-llvm-passes -emit-llvm -S
opt $1.ll -mem2reg -S > {$1}2.ll
opt ${1}2.ll --analyze --scalar-evolution
(other)
```

- -Xclang \<arg> Pass \<arg> to the clang compiler.
- -disable-llvm-passes
- -emit-llvm Use the LLVM representation for assembler and object files
- -S preprocessor only

we can see in the slides how loops are eliminated.

> Recapping Chains of recurrences
>
> - Great to model some loop varian values.
> - Algebra of simple recursive function
> - Algebraic simplifications
> - Strength reduction
> - Closed forms
> - and many more...

### Value Numbering

eliminating some forms of redundancy.

this code has redundancy.

```cpp
int calculate(int a, int b)
{
    int result = (a * b) +2;
    if (a %2 ==0)
    {
        result +=a*b;
    }
    return result;
}
```

the compiler can do the common expression optimization in some cases. but most of the redundancy isn't from the programmer. this code had redundancy in terms of memory access;

```cpp
int matrix[5][5];
//...
matrix[1][2]=bar();
matrix[1][3]=baz();
```

is actually memory dereferencing with a common sub expression.

```cpp
int matrix[5][5];
//...
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 2)=bar();
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 3)=baz();
```

we can also have dead_code and unused code that passes around (constant propagation).
compilers work in phases, and at each pass the complier cleans up the code to make it optimize. each pass does a small change.

[BRIL - big red intermediate language](https://github.com/sampsyo/bril) is a compiler IR (Intermediate representation) that is used in some courses to teach about compilers.

optimizations can work across different scopes (function, loop body, and even higher!);

local value numbering optimization. algebraic identities, dead code elimination, constant folding,

### where to learn more

some sources to learn mode about compilers.
audience questions

</details>
