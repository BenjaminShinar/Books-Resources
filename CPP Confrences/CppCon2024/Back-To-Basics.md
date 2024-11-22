<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Back To Basics

<summary>
11 Talks about basics topics.
</summary>

- [x] Back to Basics: Almost Always Vector? - Kevin Carpenter
- [ ] Back to Basics: Concepts - Nicolai Josuttis
- [ ] Back to Basics: Debugging and Testing - Greg Law, Mike Shah
- [ ] Back to Basics: Function Call Resolution - Ben Saks
- [ ] Back to Basics: Functional Programming in C++ - Jonathan Müller
- [ ] Back to Basics: Generic Programming - David Olsen
- [ ] Back to Basics: Lifetime Management - Phil Nash
- [ ] Back to Basics: Object-Oriented Programming - Andreas Fertig
- [ ] Back to Basics: R-values and Move Semantics - Amir Kirsh
- [ ] Back to Basics: Unit Testing - Dave Steffen

### Back to Basics: Almost Always Vector - Kevin Carpenter

<details>
<summary>
Vector is usually the correct container to use.
</summary>

[Back to Basics: Almost Always Vector](https://youtu.be/VRGRTvfOxb4?si=EESDcTX28liWC3ZU), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Back_to_Basics_Almost_Always_Vector.pdf), [github](https://github.com/kevinbcarpenter/almost-always-vector).

```cpp
include <array>
#include <iostream>
#include <vector>

int a[] = {0, 1, 2, 3, 4};

std::vector<int> c = {0, 1, 2, 3, 4};

auto main() -> int {
  std::cout << "C style array: " << sizeof(a) / sizeof(a[0]) << std::endl;
  std::cout << "Vector size: " << c.size() << std::endl;
  return 0;
}
```

<cpp>std::vector</cpp> is very popular. we can compare it to a C-style array. C-style arrays don't support deletions or adding elements, and we can't copy from one another (must do element-by-element assignment). the original paper for the vector (dynamic array) came out in 1992.\
we can look at the original defintion, a templated class of dynamic array. it had only one templated parameter (no allocator parameter). over the years there were more member functions added, more ways to access element, iterators, allocators, modifiers...

#### Basics

different ways to create vectors, empty constructor, with a size, with a size and default value for all elements, or with the data elements we want. we can access elements with the index operator `[]` or the `at()` method, using index directly can go out of bounds, but using the methods checks the boundaries and throws an exception. we can always use `.data()` to retrieve the underlying C-style array from the vector.

#### Memory Management

a vector stores the data on the heap (actually using the provided allocator), C-style arrays are limited by the size of the stack (usually 8mb), unless we request the data from the heap with dynamic allocation.

(we can check the size of the stack on mac/linux with `ulimit -s` command)

| Stack                        | Heap                            |
| ---------------------------- | ------------------------------- |
| Fast - pointer adjustment    | Flexible - dynamic at runtime   |
| Automatic - easy clean up    | Large - bigger is better right? |
| Predictable - easy to debug  | Lifetime - !f(x) dependent      |
| Locality - cache performance | Sharing - between two threads   |
| Safety - see automatic       | Memory leaks                    |

> Use the stack when data is small doesn't need to persist beyond the function and you want speed.
> Use the heap when data is large, needs to persist, be flexible. Watch your toes.

Vector have a size and a capacity, the size is the number of elements currently in the container, while the capacity is the number of elements we can store before needing to allocate more memory from the heap. with old arrays, we have to manually manage the memory allocations and deletions, and manage copying the data and make sure to never access the memory address of the old array.\
We can use `.reserve()` to prevent allocations, if we have a guess how many elements we would want, we can allocate the data before hand, which will prevent allocating memory and prevent copies of the elements. we also have `.shrink_to_fit()`, which reduces the capacity. however, there is no guarantee the memory is really released, the behaviour is up to the C++ Standard Library implementation.\
The second template parameter is the allocator, we can use <cloud>std::pmr</cloud> allocator to allocate the memory on the stack, rather than the heap.

#### Iterators

Iterators are *usually* wrappers for pointers with some guardrails to minimize errors, the `*` operator will give us the underlying data (just like a pointer), but it's not always a pointer. it's an abstraction.\
There are different types of iterators, forward, backwards, bi-directional, random-access, some are const and some allow changes. The `.end()` iterator points to the element **after** the final element.

```cpp
#include <iostream>
#include <vector>

int main() {
  std::vector<int> co = {2019, 2020, 2021, 2022, 2023, 2024};

  std::cout << "Is range based for loop an iterator?s\n";
  for (auto yr : co) {
    std::cout << yr << " ";
  }

  std::cout << "\nIterator works as a pointer and not a copy...\n";
  for (auto it = co.begin(); it != co.end(); ++it) {
    std::cout << *it << " "; // Dereference the iterator to get the value
  }

  return 0;
}
```

Iterators can be invalidated, this depends on the container and the actions that happened since it was created. ([documentation](https://en.cppreference.com/w/cpp/container#Iterator_invalidation), [video](https://youtu.be/Fv8oj8EdssY?si=AIfpp3gOqXNXQkfI)), the <cpp>std::erase</cpp> used to be an issue, as it would have to be paired with <cpp>std::remove</cpp>.

#### Algorithms

we use a lot of algorithms with unary predicates and iterators. it can be a function pointer, a functor or a lambda.

#### Container Comparisons

comparing to other containers:

- <cpp>std::list</cpp> (double linked list) - for multiple insertions and deletions
- <cpp>std::deque</cpp> - allows random access, but also fast when working on the front and back of the data.
- <cpp>std::map</cpp> - not the same thing.

#### Why Almost Always Vector

- cache friendly
- efficient
- practical
- allows random access
- versatile

</details>
