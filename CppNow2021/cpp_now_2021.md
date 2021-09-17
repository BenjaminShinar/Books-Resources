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
