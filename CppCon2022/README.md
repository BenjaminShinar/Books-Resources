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
- [Concurrency](Concurrency.md)
  - C++20’s Coroutines for Beginners - Andreas Fertig
  - Deciphering C++ Coroutines - A Diagrammatic Coroutine Cheat Sheet - Andreas Weis
- [Debugging & Logging & Testing](Debugging%20&%20Logging%20&%20Testing.md)
  - Compilation Speedup Using C++ Modules: A Case Study - Chuanqi Xu
  - Back to Basics: Debugging in C++ - Mike Shah
  - C++ Dependencies Don’t Have To Be Painful - Why You Should Use a Package Manager - Augustin Popa
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
- [General C++](General%20C++.md)
- [Idioms & Techniques](Idioms%20&%20Techniques.md)
- [Interface Design & Portability](Interface%20Design%20&%20Portability.md)
- [Math & Numbers](Math%20&%20Numbers.md)
  - Principia Mathematica - The Foundations of Arithmetic in C++
- [Networking & Web](Networking%20&%20Web.md)
- [Resource Management](Resource%20Management.md)
- [Scientific Computing](Scientific%20Computing.md)
  - C++ Performance Portablity - A Decade of Lessons Learned - Christian Trott
- [Social Registration](Social%20Registration.md)
- [Software Design](Software%20Design.md)
  - How C++23 Changes the Way We Write Code - Timur Doumler
  - How Microsoft Uses C++ to Deliver Office - Huge Size, Small Components - Zachary Henkel
- [Templates & Metaprogramming](Templates%20&%20Metaprogramming.md)
  - Back to Basics: Templates in C++ - Nicolai Josuttis
- [Tooling](Tooling.md)
- [Value Semantics](Value%20Semantics.md)
- [Lighting Talks](#lightning-talks)

## Lightning Talks


### Help! My Codebase has 5 JSON Libraries - How Generic Programming Rescued Me - Christopher McArthur

<details>
<summary>

</summary>

[Help! My Codebase has 5 JSON Libraries - How Generic Programming Rescued Me](https://youtu.be/Oq4NW5idmiI), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/CppCon-2022-How-Generic-Programming-came-to-the-rescue.pdf)

> Focus of the talk - “**implementing traits with functions**”.\
> Explanation of template metaprogramming implementation to abstraction JSON libraries.
> - Detecting if a function or method are implement for a type
> - Checking if an ADL implementation exists
> - Compile time requirements and SFINAE


once they added a package manager, it was easy to get more packages, and different libraries which do the same things are added to the stack.

> “Why don’t you template out the logic and metaprogram a traits
implementation?”

```cpp
template<typename json_traits>
class basic_claim {
  static_assert(details::is_valid_traits<json_traits>::value,
  "traits must satisfy requirements");

  static_assert(
  details::is_valid_json_types<typename json_traits::value_type,
  typename json_traits::string_type,
  typename json_traits::integer_type,
  typename json_traits::object_type,
  typename json_traits::array_type>::value,
  "must satisfy json container requirements");
}
```

JWT -JSON Web Token, some predefined key. we need to both **create** and **verify** tokens.

the values can only be a limited set of types.

#### Verifying Claims
> Q: How can we check a type implements a function?

combine: using `std::experimental::is_detected`, which employs SFINAE to detect a named entity, but we need to add `std::is_function`, and eventually, we want to create a `is_signature` trait.

#### Creating Tokens

decltype doesn't work easily with overload functions, we need to resolve the function at compile time. more assertion errors over member functions.

variance with methods (array `.at` vs `[]`).


> Review
> - Check static function signatures with `is_detected`, `is_function`, and
`is_same`
> - We can resolve overloaded functions with the help of `declval`
> - To overcome `declval`’s lack of substitution we can add template helpers to return `true_type` of `false_type` is it does not resolve.
> - More indirection is usually the answer with SFINAE

</details>

### HPX - A C++ Library for Parallelism and Concurrency - Hartmut Kaiser
<!-- <details> -->
<summary>

</summary>



[HPX - A C++ Library for Parallelism and Concurrency](https://youtu.be/npufmMlGOoM), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/HPX-A-C-Standard-Library-for-Parallelism-and-Concurrency-CppCon-2022-1.pdf)


> HPX – An Asynchronous Many-task
Runtime System

differences in performance between compute-bound and memory-bound parallel tasks. also difference performance patterns depending on the number of cores.

HPX: a threading implementation, more efficient than naive threads (jthread, std::thread, pthread), with several functional layers, conforming to the standard, and offers some externsions. allows for distributes execution.

this talk will focus on parallel loop and algorithms

> - Simple iterative algorithms
>   - One pass over the input sequence.
>   - for_each, copy, fill, generate, reverse, etc. 
> - Iterative algorithms ‘with a twist’
>   - One pass over the input sequence.
>   - Parallel execution requires additional operation after first pass, most of
the time this is a reduction step
>   - min_element, all_of, find, count, equal, etc.
> - Scan based algorithms
>   - At least three algorithmic steps.
>   - inclusive_scan, exclusive_scan, etc. 
> - Auxillary algorithms
>   - Sorting, heap operations, set operations, rotate


parallelization can work on CPU (threads and cores), and on GPUs.


#### Parallelize Loops

execution policies:

> Convey guarantees/requirements imposed by loop body
> -  seq: execute in-order (sequenced) on current thread
> -  unseq: allow out-of-order execution (un-sequenced) on current thread - vectorization
> -  par: allow parallel execution on different threads
> - par_unseq: allow parallel out-of-order (vectorized) execution on different threads

in the future the standard might include *std::simd*, which will require explicit vectorization. but that's still in the experimental stage.

HPX adds more policies, explicit parallelized vectorization and execution.

the first example uses std::future to launch threads and parallelize on them, it cuts the input data into chunks.

#### Background

Amdahl's Law (Strong scaling), the speed up of parallelizion is capped by the number of processors and how much of the code can be parallelized

$
S = \frac 1{(1-p) + \frac {P}{N}}
$

SLOW - problems with parallelization
> - Starvation
>   - Insufficient concurrent work to maintain high utilization of resources.
> - Latencies
>   - Time-distance delay of remote resource
access and services
> - Overheads
>   - Work for management of parallel actions and resources on critical path which are not necessary in sequential variant.
> - Waiting for Contention resolution.
>   - Delays due to lack of availability of
oversubscribed shared resources.


there is a U-shaped curve of gains from spliting chunks. the overhead can be greater than the gains, but if the there are too many chunks (with smaller data), the performance will go down. this goes together with the number of cores.

#### Executors

> Executors abstract different task launching infrastructures
> - Synchronization using futures
>   - HPX historically uses futures as main means of coordinating.
> -  Synchronization using sender/receivers (C++26?)
>   - C++ standardization focusses on developing an infrastructure for anything related to asynchrony and parallelism.
>   - P2300: std::execution (senders & receivers)
>   - Computational basis for asynchronous programming
>   - Current discussions focus on integrating parallel algorithms.


some examples: execution policies, attaching to an executor, parallel task execution (asynchronous and futures), senders & receivers. eager and lazy executions.

**Explicit vectorization**: 
simd - single instruction, multiple data.

**Linear algebra**, a proposal for standardization, might also work with execution policies. `std::linalg::scale`

</details>