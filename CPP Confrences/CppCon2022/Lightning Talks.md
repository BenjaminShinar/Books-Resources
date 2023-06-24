<!--
ignore these words in spell check for this file
// cSpell:ignore JupyText
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Lightning Talks

<!-- <details> -->
<summary>
Lightening Talks - short, not necessary to the point, talks about whatever.
</summary>

### How to Win at Coding Interviews - David Stone

<details>

[How to Win at Coding Interviews](https://youtu.be/y872bCqQ_P0)

1. repeat the question, ask for clarifications
2. write the interface
3. use hashmap or the \<algorithm> header

example:

> given an integer arrays of size n, find all elements that appear more the $n/3$ times.

we write an interface that takes a vector (or span) and returns a vector. then we ask ourselves how it could be solved with the hashmap.

we also have the advantage that the time complexity of hash map is constants.

if we are told the values are ordered or need to be returned in a specific order, then that's a sign we shouldn't use a hashmap.

</details>

### Best Practices Every C++ Programmer Needs to Follow - Oz Syed

<details>

[Best Practices Every C++ Programmer Needs to Follow](https://youtu.be/xhTINjoihrk)

1. memory management is crucial, freeing memory, properly writing destructors, using smart pointers.
2. don't take shortcuts, write all the operators that you needs (copy, move, assignment)
3. use multiple compilers to find different issues.
4. test and repeat.
5. white box tests by the developer, black box testing by qa. good tests are also those which fail.
6. some tasks are best done by tools

</details>

### C++20 - A New Way of Meta-Programming? - Kris Jusiak

<details>

[C++20 - A New Way of Meta-Programming?](https://youtu.be/zRYlQGMdISI).

five game changing features of meta programming

1. Design by introspection - `if constexpr (requires{ t.foo; })`
2. Immediately invoked function expression inside compile-time expressions.
3. Fixed String - passing the type of string as template argument.
4. reflection `to_tuple` -
5. `constexpr std::vector`

</details>

### MP: Template Meta-Programming in C++ - Kris Jusiak

<details>

[MP: Template Meta-Programming in C++](https://youtu.be/-4MSlna4gKE)

circle -using the `@meta` annotation keyword to write compile time code with the same C++ Standard Template Library.

example of rotate a range in compile time, using _circle_ and using c++20.

there is a proposal of "generalized pack declaration and usage" that will bring more compile time capabilities into the standard.

</details>

### The Future of C++ - Neil Henderson

<details>

[The Future of C++](https://youtu.be/QqQA7_8QuwY) - Neil Henderson

(many jokes about Australians)

the only real part of the talk is that RAII is renamed into "Scope-bound resource management"

</details>

### Dependency Injection for Modern C++ - Tyler Weaver

<details>

[Dependency Injection for Modern C++](https://youtu.be/Yr0w62Gjrlw)

- separation of concerns
- make appendices minimal and explict
- mocking for testing

```cpp
TEST(OptimizerTest, MockClient){
  struct ClientMock: public ClientInterface {
    double cost(State) override {return 0.3;}
  };
  auto optimizer = Optimizer(std::make_unique<ClientMock()>);
  // test!
};
```

we can replace this interface with <cpp>std::function</cpp> and make te tests trivial to write. so if we have functions that take functions themselves, things get simpler. this also simplifies objects that have complicated states.

</details>

### Cute Approach for Polymorphism in C++ - Liad Aben Sour Asayag

<details>

[Cute Approach for Polymorphism in C++](https://youtu.be/Yr0w62Gjrlw)

a container with many elements, many handlers can handle the elements, but we only know which one is used only at runtime. the basic way is to use virtual inheritance. this is good for encapsulation, but not for performance. we could use concepts and then drop the inheritance in favour of <cpp>std::variant</cpp>, but that means the calling code needs to know about all the types, and be recompiled again and again.

Virtual on Aggregation - a different approach is to have the virtual function take the span of elements, and internally call the actual function, this way we have less virtual calls and we can optimize the actual working function, because it is known at compile time. we can add an abstraction layer by using templates to reduce boiler plate and make the code elegant.

</details>

### Finding the Average of 2 Integers - Tomer Vromen

<details>

[Finding the Average of 2 Integers](https://youtu.be/rUt9xcPyKEY)

the infamous midpoint problem, for big integer numbers we can go into overflow, so we need to use <cpp>std::midpoint</cpp>. but in nearly half the cases, the average of two integers is a floating point number, but even converting to a double has problems, if we use a number that can be represented as a double we get a rounding error.

</details>

### The Lambda Calculus in C++ Lambdas - David Stone

<details>

[The Lambda Calculus in C++ Lambdas](https://youtu.be/1hwRxW99lg0)

lambdas calculations are at the same level of foundation importance as turing machines.

everything is defined in terms of Functions.

```cpp
auto main(int argc, char** argv)-> int {
  return -(-![]{});
}
```

now that we defined zero, it's time to define the rest of the numbers, which we do with the successor function.

```cpp
auto zero = (-(-![]{}));
auto  one = (-(-![]{})) - (-!(![]{}));
auto  two = ((-(-![]{})) - (-!(![]{}))) - (-!(![]{}));
```

a demo of a code that reads input and squares it using only that weird lambda

</details>

### `find-move-candidates` in Cpp - Chris Cotter

<details>

[find-move-candidates in Cpp](https://youtu.be/F8wbpi2kTmY), [github](https://github.com/bloomberg/clangmetatool)

this code does unnecessary copies, when it could be using <cpp>std::move</cpp>.

```cpp
std::vector<std::string> fields;
for (int i: views::iota(0,1000))
  fields.push_back(/*...*/);

Data data;
data.set_fields(fields);

Request request;
request.set_data(data);
```

so there's a tool that helps find candidates for that and will suggest moving data when it's safe to do so.

</details>

### Modernizing SFML in Cpp - Chris Thrasher

<details>

[Modernizing SFML in Cpp](https://youtu.be/JJPL17sDxUs)

modernizing the SFML - simple and fast multimedia library from c++03 to C++17. a cross library platform with Cmake.

c++17 migration had many stuff to take advantage of, such as the filesystem api, using optional and not output parameters, attributes.

</details>

### Finding Whether a Number is a Power of 2 - Ankur Satle

<details>

[Finding Whether a Number is a Power of 2](https://youtu.be/Df-qEsWjzQw)

a common interview question, counting "on" bits in number.

```cpp
constexpr bool is_pow_knr(std::unsigned_integral auto n){
  return (n & (n-1)) == 0;
}
```

bitset version isn't constexpr. in c++20 there is the <cpp>bit</cpp> header with built-in functions.

</details>

### Const Mayhem in C++ - Ofek Shilon

<details>

[Const Mayhem in C++](https://youtu.be/cuQKiXZUmyA)

can't create a `const` object when there is no default constructor. but that's only if the constructor is defined inside the class deceleration.\
an example of a const method that modifies values without being mutable - abusing pointers.

```cpp
struct C {
  int m_i;
  int* m_p = &m_i;
  void const_method() const { ++(*m_p); } // m_i is modified
};
```

</details>

### Who is Looking for a C++ Job? - Jens Weller

<details>

[Who is Looking for a C++ Job? - Jens Weller](https://youtu.be/0xI0H7djNOc)

an online C++ job fair that happens a few time each year. showing up some data about applicants.

</details>

### Effective APIs in Practice in C++ - Thamara Andrade

<details>
[Effective APIs in Practice in C++](https://youtu.be/YdZLsSDZ_Qc)

tips to improve API and make he

- better naming - for the give the unit name in the api
- use strong types - explicit creation of the required object
  - adding our literals
- avoid easily swappable parameters - there is a clang tidy tool to help finding those cases
- think about intent
- keep learning!
</details>

### Standard Standards for C++ - Ezra Chung

<details>

[Standard Standards for C++](https://youtu.be/vds3uT9dRCc)

What is the "standard" C++?

- dictionary defintion of "standard", many dictionaries, many defintions, many domains.

so there are many definitons of "standard C++", such as versions specifications (past, present and future), writing standard c++ means different code for each standard. this also takes into consideration the target and compiler flag.

this doesn't include formatting, design, tools or architecture, which can also be referred to as "standard C++", we need to be explicit and clear.

</details>

### Ref, C++ const ref, immutable ref? - Francesco Zoffoli

<details>

[Ref, C++ const ref, immutable ref?](https://youtu.be/OupN6FMZbmA)

- `ref` - "I want to access some data, and I can modify it"
- `const ref` - "I want to access some data, and I promise to not modify it"
- `immutable ref` - "I want to access some data, and no one can modify it" - **doesn't exist**

we can take a const reference to a global variable, so then it could be changed. we don't have an easy way to have an immutable reference

</details>

### C++ Debug Performance is Improving: What Now? - Vittorio Romeo

<details>

[C++ Debug Performance is Improving: What Now?](https://youtu.be/CfbuJAAwA8Y)

- debugging with optimization enabled is a nightmare
- we want to retain the performance and still be able to debug.
- zero cost abstractions rely on compiler optimizations
- people sometimes write less abstracted code so they would have debug builds they can actually use.

```cpp
#include <cstddef>

using byte_type = char;
//using byte_type = std::byte;

byte_type example()
{
  byte_type b{123};
  b <<= 1 ; // shift left
  return b;
}
```

in this example, the debug build of using the <cpp>std::byte</cpp> abstraction produces twice as many assembly lines. Which makes performance worse and it's harder to understand. there is an example of how using <cpp>std::accumulate</cpp> in C++20 without optimization is slower than C++17. this is because the C++20 version tries to use move semantics, which the compiler will optimize away for primitive types.

</details>

### Embrace Leaky Abstractions in C++ - Phil Nash

<details>

[Embrace Leaky Abstractions in C++](https://youtu.be/uh15LjpBIP0)

A leaky abstractions exposes both the complexities of the abstraction and of the thing that is abstracted away. we say that all non trivial abstracting are leaking away, even math is an abstractions, that's how we get negative numbers and division by zero, and square roots of negative numbers.

if all abstractions are leaky, we need to consider this when creating abstractions, we would want them to be small and simple.\
some abstractions global - like libraries, and types, we want them to be complete and water-tight.\
other abstractions are local - they come from our own code base, we want them to be shallow, so when the leak happens, it's obvious where it comes from.

</details>

### 10 Things an Entry-Level Software Engineer Asks You to Do - Katherine Rocha

<details>

[10 Things an Entry-Level Software Engineer Asks You to Do](https://youtu.be/RkH8P1RgYIs)

what new software developers want.

1. Let us set up our environment - which guides are lacking?
2. walkthrough the design process
3. walkthrough the code strucure
4. explain the weird code things
5. explain the "why" - how the code got to be this way
6. show the end result
7. assign a task that spans the code base - get the programmer familiar with it
8. assign feature work
9. invite us to offsite stuff - lunch, meeting etc...

</details>

### C++ on Fly - C++ on Jupyter Notebook - Nipun Jindal

<details>

[C++ on Fly - C++ on Jupyter Notebook](https://youtu.be/MtKdza3RJNM)

A scratch pad in jupyter notebook - code and rich text. **Xeus** is an implementation of the jupyter Kernel protocol, and together with **cling** interferer, we can have a working C++ notebook.

```sh
conda create -n cling
conda install xeus-cling -c conda-forge
conda install xeus -c conda-forge
conda activate cling

jupyter notebook
jupyter lab
```

there is also a docker setup.

versioning with **JupyText**, quick sharing with **Binder**.

</details>

### Using This Correctly it's `[[unlikely]]` at Best - Staffan Tjernstrom

<details>

[Using This Correctly it's `[[unlikely]]` at Best](https://youtu.be/_1A1eSriCV4)

new attributes in C++20. <cpp>[[likely]], [[unlikely]]</cpp>. it's tricky to use them, and using micro-benchmark will give us skewed results. using PGO (profiling guided optimizations) can bring us worse results, there is even an example, if we have an "emergency" break function on a self-driving car, then it will go on the cold path, but this is disastrous for actual behavior of the car.

the placement of the attribute in the code also matters.

</details>

### Programming is Fun in Cpp! - Pier-Antoine Giguère

<details>

[Programming is Fun in Cpp!](https://youtu.be/F9c1ZuSRdsM)

playing with ray tracing and generating images.

</details>

### Adventures in Benchmarking Timestamp Taking in Cpp - Nataly Rasovsky

<details>

[Adventures in Benchmarking Timestamp Taking in Cpp](https://youtu.be/-XU2silGr6g)

a story about benchmarking - are system calls really so much slower? maybe, but in this story the system call isn't really called, it doesn't move from userspace.

</details>

### `-std=c++20` -- Will This C++ Code Compile? - Tulio Leao

<details>

[`-std=c++20` -- Will This C++ Code Compile?](https://youtu.be/87_Ld6CMHAw)

- using reserved keywords as variable names ("requires","concept")
- identifiers for the standard library (such as "\_TR")
- incompatibility between <cpp>std::string</cpp> and <cpp>std::fs::path u8string()</cpp>. was ok in c++17, but not anymore.
- redundant template-id on constructors
- aggregate initialization of structs with deleted default constructors
- removed <cpp>std::allocator</cpp> members and functions, previously deprecated, then obsolete, now removed, replace with <cpp>std::allocator_traits</cpp>
- <cpp>std::accumulate</cpp> attempts to move the first element, so it won't work with non-const references.
</details>

### `majsdown`: Metaprogramming? In my Slides? - Vittorio Romeo

<details>

[`majsdown`: Metaprogramming? In my Slides?](https://youtu.be/vbhaZHpomg0)

markdown + node.js - writing markdown slides with code, such as generating compiler explorer links, dynamic expressions in markdown, operating on code blocks etc...

</details>

### The Decade Long Rewind: Lambdas in C++ - Pranay Kumar

<details>

[The Decade Long Rewind: Lambdas in C++](https://youtu.be/xBkDkCgQsAM).

- C++11 introduced lambdas, unique typed closure which the compiler produces. it had capture, mutable, throws.
- c++14 added default parameters, template parameters with `auto&&`,generalized capture inside the parentheses, and being able to return lambdas from function with the `auto` return type.
- C++17 allowed lambdas to be constexpr, and fave a simple capture for with `[*this](){}`.
- C++20 made lambda mre aligned with templates, following a similar syntax and allowing to capture variadic parameter pack.
- C++23 allows for omitted empty parameters list, and makes recursive lambdas easier to use.
</details>

### Cpp Change Detector Tests - Guy Bensky

<details>

[Cpp Change Detector Tests](https://youtu.be/rv-zHn_Afko)

a term for "tests" that break on changes, they are usually considered bad tests, because they don't fail when the code behavior fails, rather they break on stuff randomly.

he argues that sometimes this is ok, such as when testing performance, which can have tradeoffs between components, there is also "testing" for metrics.

we can use "change detector" tests to find un-intended consequences that our code makes. there are always "soft decisions" that a developer makes, so this moves them to the code-review stage.

</details>

### Developer Grief - Thamara Andrade

<details>

[Developer Grief](https://youtu.be/xI9YEp0G_JQ)

bugs and griefs, five stages of grief:

1. denial - the user did a mistake - analysis stage
2. anger - how did this happen? - debugging
3. bargaining - what if i did things different? is there another solution - negotiation
4. depression - dealing with the bug - development
5. acceptance - committing the fix, and knowing we did our best - release
</details>

### Now in Our Lifetimes Cpp - Staffan Tjernstrom

<details>

[Now in Our Lifetimes Cpp](https://youtu.be/pLVg3c6bljE)

implicit object creation in C++20, except for all the the parts where it doesn't work.
In C++23 we will get <cpp>std::start_lifetime_as</cpp> and <cpp>std::start_lifetime_as_array</cpp> which will make things easier.\
there is also <cpp>std::launder</cpp>, but ut doesn't create objects or create lifetimes.

</details>
