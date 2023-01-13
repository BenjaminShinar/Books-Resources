<!--
ignore these words in spell check for this file
// cSpell:ignore
-->

# CppCon 2022

[website](https://cppcon.org/), [youtube playlist](https://www.youtube.com/playlist?list=PLHTh1InhhwT6c2JNtUiJkaH8YRqzhU7Ag), [materials index](https://github.com/CppCon/CppCon2022).

## Topics

- [Algorithms & Data Structures](Algorithms%20&%20Data%20Structures.md)
  - The Imperatives Must Go! [Functional Programming in Modern C++] - Victor Ciura
  - Functional Composable Operations with Unix-Style Pipes in C++ - Ankur Satle
  - Understanding Allocator Impact on Runtime Performance in C++ - Parsa Amini
  - Refresher on Containers, Algorithms and Performance in C++ - Vladimir Vishnevskii
- [Concurrency](Concurrency.md)
  - C++20’s Coroutines for Beginners - Andreas Fertig
  - Deciphering C++ Coroutines - A Diagrammatic Coroutine Cheat Sheet - Andreas Weis
  - C++ Concurrency TS-2 Use Cases and Future Direction - Michael Wong, Maged Michael, Paul McKenney
  - An Introduction to Multithreading in C++20 - Anthony Williams
  - A Lock-Free Atomic Shared Pointer in Modern Cpp - Timur Doumler
- [Debugging & Logging & Testing](Debugging%20&%20Logging%20&%20Testing.md)
  - Compilation Speedup Using C++ Modules: A Case Study - Chuanqi Xu
  - Back to Basics: Debugging in C++ - Mike Shah
  - C++ Dependencies Don’t Have To Be Painful - Why You Should Use a Package Manager - Augustin Popa
  - Going Beyond Build Distribution - Using Incredibuild to Accelerate Static Code Analysis and Builds - Jonathan "Beau" Peck
  - Case For a Standardized Package Description Format for External C++ Libraries - Luis Caro Campos
  - C++ Coding with Neovim - Prateek Raman
- [Education Coaching](Education%20Coaching.md)
- [Embedded](Embedded.md)
  - Using C++14 in an Embedded “SuperLoop” Firmware - Erik Rainey
  - Taking a Byte Out of C++ - Avoiding Punning by Starting Lifetimes - Robert Leahy
- [Future of C++](Future%20of%20C++.md)
  - Contemporary C++ in Action - Daniela Engert
  - Can C++ be 10x Simpler & Safer? - Herb Sutter
  - C++ in Constrained Environments - Bjarne Stroustrup
  - What’s New in C++23 - Sy Brand
  - C++23 - What's In It For You? - Marc Gregoire
  - Understanding C++ coroutines by example: Generators - Pavel Novikov
  - 10 Years of Meeting C++ - Historical Highlights and the Future of C++ - Jens Weller
- [General C++](General%20C++.md)
- [Idioms & Techniques](Idioms%20&%20Techniques.md)
  - C++ Lambda Idioms - Timur Doumler
  - Undefined Behavior in the STL - Sandor Dargo
  - C++ MythBusters - Victor Ciura
  - C++ Function Multiversioning in Windows - Joe Bialek and Pranav Kant
  - Embracing Trailing Return Types and `auto` Return SAFELY in Modern C++ - Pablo Halpern
  - The Most Important Optimizations to Apply in Your C++ Programs - Jan Bielak
  - Back to Basics: RAII in C++ - Andre Kostur
- [Interface Design & Portability](Interface%20Design%20&%20Portability.md)
  - Purging Undefined Behavior & Intel Assumptions in a Legacy C++ Codebase - Roth Michaels
  - Managing External API's in Enterprise Systems - Pete Muldoon
- [Math & Numbers](Math%20&%20Numbers.md)
  - Principia Mathematica - The Foundations of Arithmetic in C++
- [Networking & Web](Networking%20&%20Web.md)
  - WebAssembly: Taking Your C++ and Going Places - Nipun Jindal & Pranay Kumar
- [Resource Management](Resource%20Management.md)
  - Back to Basics: C++ Move Semantics - Andreas Fertig
  - Back to Basics: C++ Smart Pointers - David Olsen
- [Scientific Computing](Scientific%20Computing.md)
  - C++ Performance Portablity - A Decade of Lessons Learned - Christian Trott
  - HPX - A C++ Library for Parallelism and Concurrency - Hartmut Kaiser
  - Graph Algorithms and Data Structures in C++20 - Phil Ratzloff & Andrew Lumsdaine
  - Breaking Enigma With the Power of Modern C++ - Mathieu Ropert
  - MDSPAN - A Deep Dive Spanning C++, Kokkos & SYCL - Nevin Liber
- [Social Registration](Social%20Registration.md)
- [Software Design](Software%20Design.md)
  - How C++23 Changes the Way We Write Code - Timur Doumler
  - How Microsoft Uses C++ to Deliver Office - Huge Size, Small Components - Zachary Henkel
  - The Hidden Performance Price of C++ Virtual Functions - Ivica Bogosavljevic
  - Using Modern C++ to Eliminate Virtual Functions - Jonathan Gopel
- [Templates & Metaprogramming](Templates%20&%20Metaprogramming.md)
  - Back to Basics: Templates in C++ - Nicolai Josuttis
  - Help! My Codebase has 5 JSON Libraries - How Generic Programming Rescued Me - Christopher McArthur
  - High Speed Query Execution with Accelerators and C++ - Alex Dathskovsky
  - Taking Static Type-Safety to the Next Level - Physical Units for Matrices - Daniel Withopf
  - C++ for Enterprise Applications - Vincent Lextrait
  - From C++ Templates to C++ Concepts - Metaprogramming: an Amazing Journey - Alex Dathskovsky
- [Tooling](Tooling.md)
  - import CMake, CMake and C++20 Modules - Bill Hoffman
- [Value Semantics](Value%20Semantics.md)
  - Back to Basics: Cpp Value Semantics - Klaus Iglberger
- [Lighting Talks](#lightning-talks)
  - The Dark Corner of STL in Cpp: MinMax Algorithms - Simon Toth

## Lightning Talks

### The Dark Corner of STL in Cpp: MinMax Algorithms - Simon Toth

<details>
<summary>
Some edge cases and problems with the min max algorithms.
</summary>

[The Dark Corner of STL in Cpp: MinMax Algorithms](https://youtu.be/jBeTvNgW25M), [slides](https://github.com/HappyCerberus/cppcon22-talk).

Free book about standard C++ algorithms available on github.

why are the min/max algorithms so hard? aren't they simple?

```cpp
auto min = std::min(1,2); // 1
auto max = std::max(1,2); // 2
auto clamped = std::clamp(0,1,2); //1, value, min, max
auto minmax = std::minmax(1,2);
```

but if we look at the templates, we see that minmax returns a pair.

```cpp
template<class T>
const T& min(const T& a, const T& b);

template<class T>
const T& max(const T& a, const T& b);

template<class T>
const T& clamp(const T& v,const T& lo, const T& hi);

template<class T>
std::pair<const T&, const T&> min(const T& a, const T& b);
```

so if we write the code we get references to temporary elements, auto type deduction doesn't deduce reference type.

```cpp
std::pair<const int&, const int&> minMax =std::minmax(1,2);
const int& min = std::min(1,2); // min is a dangling reference

auto [x,y] = std::minmax(1,2); // still dangling
std::pair<int,int> a = std::minmax(1,2); // this is ok.
```

to find this behavior, we need to run an address sanitizer.

there are some variants of the min max algorithms, the c++20 range versions behave the same.

c++14 has variant that take an initializer list, which return by value and not by reference.

```cpp
auto x = std::min({1,2});
// ok, decltype(x) => int

auto pair = std::minmax({1,2});
// ok, decltype(pair) => std::pair<int, int>

const int &z = std::max({1,2});
// ok, lifetime extension
```

but there are problem, because it's impossible to move from an initializer list. so it fails for move only types, and can incur heavy costs for multiple copies.

```cpp
auto x = std::min({MoveOnly{}, MoveOnly{}});
// wouldn't compile

ExpensiveToCopy a,b;
auto y = std::min({a,b});
// 3 copies

auto z = std::min(ExpensiveToCopy{},ExpensiveToCopy{});
// 1 copy since c++17, copy-initialization from prvalue
```

next is the problem of **const correctness**. we have this code

```cpp
MyType a,b;
if (b<a){
  b.do_something();
} else {
  a.do_something();
}
```

but we want it to be more simple, like this:

```cpp
MyType a,b;
std::min(a,b).do_something();
```

this will only work if the method we call is const, because the return value of the algorithm is a const reference. we could use `const_cast<>`, but that isn't very readable, and we might have undefined behavior if the method mutates state.

we would like to fix this:

1. remove the need for `const_cast`
2. remove the potential for dangling reference
3. avoid excessive copies

we can fix this by adding more overloads for the algorithms. then we get things better by using `auto` templates. there's something about the terinary operator here, so we need to call `std::common_reference_t` if we don't use it. we add _requires_ clause.

</details>
