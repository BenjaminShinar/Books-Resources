<!--
// cSpell:ignore NTTP
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
- [x] Back to Basics: Generic Programming - David Olsen
- [ ] Back to Basics: Lifetime Management - Phil Nash
- [ ] Back to Basics: Object-Oriented Programming - Andreas Fertig
- [ ] Back to Basics: R-values and Move Semantics - Amir Kirsh
- [ ] Back to Basics: Unit Testing - Dave Steffen

---

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

### Back to Basics: Generic Programming in C++ - David Olsen

<details>
<summary>
Overview of writing and using templates.
</summary>

[Back to Basics: Generic Programming in C++](https://youtu.be/0_0HsEBsgPc?si=IqN1Kk4OWX8RpwPo), [event](https://cppcon2024.sched.com/event/1gZdo/back-to-basics-generic-programming), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Back_to_Basics_Generic_Programming.pdf).

generic - not specified or specialized, fits a wide variety of needs.\
generic programming - writing code that works across different types.

guidelines

> - Define templates in header files
> - Substitution checks the declaration and template arguments
> - Instantiation checks the entire definition
> - SFINAE: Substitution Failure Is Not An Error
> - Let the compiler deduce arguments for a function template
> - Constrain your template parameters
> - Keep it simple

#### Templateing basics

the basic example of computing a sum from container elements.

```cpp
template <class C, class T>
T sum(C container) {
  T result = 0;
  for (T value : container) {
    result += value;
  }
  return result;
  }
```

this is static polymorphism, the template is resolved at compiled time, and the appropiate functions are generated. in other languages, we call this "Generics", but in C++ it's called templates.

defining a template:

> template \<template-parameters\> declaration;
>
> declaration can be:
>
> - class / struct
> - function
> - type alias
> - variable
> - <cpp>concept</cpp>
>
> template-parameter is\
> `class | typename identifier [= default-value]`
>
> Template definition should be in a header file.

example of templates

```cpp

// class
template <typename T, typename U>
class pair {
  T m0;
  U m1;
public:
  pair() { }
  pair(T v0, U v1) : m0(v0), m1(v1) { }
  T first() const { return m0; }
  U second() const { return m1; }
};

// method
template <class T>
void swap_pointed_to(T* a, T* b) {
  T temp = *a;
  *a = *b;
  *b = temp;
}

// type alias declaration
template <class T> using ptr = T*;
// type alias declaration - almost meta programming
template <class Iter1, class Iter2>
using result_type = typename std::common_type<typename std::iterator_traits<Iter1>::value_type, typename std::iterator_traits<Iter2>::value_type>::type;

// variable template defintion
template <class T>
constexpr bool is_big_and_trivial = sizeof(T) > 16 &&std::is_trivially_copyable<T>::value && std::is_trivially_destructible<T>::value;
```

three kinds of template parameters (what goes inside the diamond brackets), the value must be known at compile time. the name (identifier) is optional, in case it's not used. we can provide a default value, if the user doesn't provide one. we can define the template parameter as variadic (not default values), we just add `...` before the parameter name.

> - type template parameter `class|typename identifier`
> - non-type template parameter (NTTP) `type|auto identifier`
> - template template parameter `template<template-parameter> class|typename identifier`

as an example, <cpp>std::array</cpp> has a a non-type template parameter - the array size.\
we can have a class template with templated methods, following the same syntax. they can be defined in a cpp file, but it's easier to write them in the class header.

when we use a template, sometimes the compiler can deduce the the types based on the arguments.

#### Substituion & Instantiations

first substituion happens, and the instantiations, the first checks the template parameters and how the types fit it, while the second checks the template body.

> Substitution:
>
> - Substitute template arguments for template parameters
> - Results in class declaration or function declaration
> - Checks the correctness of the template arguments
>
> Instantiation
>
> - Full definition of the class or function or type alias or variable
> - Happens after substitution, only when full definition is needed
> - Checks the correctness of the definition

there are cases when we don't need an instantiation. such as in class templates.

> - Substitution without instantiation in two contexts
>   - Incomplete type is sufficient
>   - Class template partial specialization resolution
> - Results in an incomplete class type
>   - Contents of the class are not checked
>   - Only the template arguments are checked

for function templates, the substituion happens during overload resolution, which is how we get SFINAE.

> - Substitution happens during overload resolution
>   - Unselected overloads are not instantiated
> - Results in function declaration
>   - Only function signature is checked
>     - Including parameters, return type, noexcept clause
> - Function body is not checked

class template instantiation - creating a "real" class. uses mangled names.

> - Replace template parameters with template arguments in the class definition
> - Results in a complete class definition
> - Member functions are not instantiated until they are used

the class instantiation can fail, if the resulting class definition is not valid, like defining the a c-style array as the template parameters when instantiating a pair.

function template instantiation:

> - Replace template parameters with template arguments in the function definition
> - Results in a complete function definition

SFINAE - Substitution Failure Is Not An Error:

> - A failure during substitution does not fail compilation
>   - Instead, the candidate is discarded
> - A function overload that fails substitution is not a viable candidate
> - This feature is necessary for function templates and class
> - template partial specializations to be useful

if there is no matching substitute, then we get an error, but as long as one match exists, we can move forward.

#### Using Class Templates

`class-template-name <template-arguments>`

this results in a regular type, each instantiation is a type of it's own.

```cpp
// regular distinct classes
struct A {};
struct B {};
A a;
B b = a; // error
B* bp = &a; // error

// template distinct types

template<class T> struct D {};

D<int> di;
D<double> dd = di; // error, cant convert
D<double>* ddp = & di; // error
D<const int> dci = di; // error
D<const int>* dci_p = &di; // error
```

##### Class Template Argument Kinds

Matching the type to the way the template was declared (class, typename, auto). the <cpp>std::array</cpp> requires a constant <cpp>std::size_t</cpp> as the template parameter, so it can't accept a type. templated templates must match as well.

#### Using Function Templates

> - Use function template like a regular function
> - Let the compiler deduce the arguments
>   - Unless the function's API requires explicit template arguments

#### Constraints

listing the requirements of the template, like requiring them to be copyable, default constructable, etc... they can be checked in the substituion phase, so we can use the to remove methods from the overload set.

> - Requirements on a template argument
> - Checked during substitution, not instantiation
> - Often make use of concepts and requires clauses

C++20 added concepts (the <cpp>requires</cpp> clause), but we have ways to do the same even without it. there are many ways to write it.

- <cpp>std::enable_if</cpp>
- <cpp>std::is_integral\<T>::value</cpp>
- <cpp>std::is_integral_v\<T></cpp>

### Writing Class Templates

Keep it simple

> - Keep It Simple and Straightforward
> - No fancy template metaprogramming or type-based  metafunctions
> - Make it easy for your users

document the requirements

> - Document expectations for the template parameters
>   - In code if possible
>     - via constraints
>   - In documentation otherwise
> - Member functions can have additional requirements

specialized templates

> - Sometimes one instantiation of a class template should behave differently than the others
> - we could tell the user to define a class to be used in place of the normal instantiation
> - Specialization can have a completely different interface
>   - But that is usually a bad idea

specialization example - a `sizeof` operator that won't fail for incomplete types.

```cpp
template <class T>
struct safe_sizeof {
  static constexpr std::size_t value = sizeof(T);
};

template <>
struct safe_sizeof<void> {
  static constexpr std::size_t value = 0;
};
```

it's also possible to have partial specializations

> Sometimes you want to specialize for one template parameter,
but not for all of them, Or specialize when one template parameter fits a pattern.\
> Similar to full specialization, but template parameter list is not empty.

```cpp
template <class T>
struct safe_sizeof {
  static constexpr std::size_t value = sizeof(T);
};

template <>
struct safe_sizeof<void> {
  static constexpr std::size_t value = 0;
};

template <class T>
struct safe_sizeof<T[]> {
  // Matches any array with unspecified bound
  static constexpr std::size_t value = 0;
};

template <class R, class... Args>
struct safe_sizeof<R(Args...)> {
  // Matches any function type
  static constexpr std::size_t value = 0;
};
```

but there is an easier way to do this, rather than defining more and more specializations, we can flip things around, and define the substitution to require the <cpp>sizeof</cpp> operator to be valid.

```cpp
template <class T>
constexpr std::size_t safe_sizeof = 0;

template <class T> requires (sizeof(T) > 0)
constexpr std::size_t safe_sizeof<T> = sizeof(T);
```

> Specialization Allowed?
>
> - Class template: Yes
> - Variable template: Yes
> - Type alias template: No
> - Concept: No
> - Function template: see next section

Type Alias Specialization Workaround, actually used in the standard library.

```cpp
template <class T>
struct remove_pointer {
  using type = T;
};

template <class T>
struct remove_pointer<T*> {
  using type = T;
};

template <class T>
using remove_pointer_t = typename remove_pointer<T>::type;
```

sometimes we need to help the compiler so the `typename` keyword is added.

> Compiler needs help parsing template definition.
> Keyword typename must precede any qualified type name that
depends on a template parameter.

as an example, the outcome of the code below depends on the types.

```cpp
A * B(C(D)); // what is this?
```

- if A not a typename (is a value), this is Expression statement: multiply A and `B(C(D))`.
- if A is a type name, but either B or D are not, then this is variable definition: B is a variable of type pointer-to-A with the initial value of `C(D)`.
- if A, C and D are all type names, then this is a Function declaration: B is a function with parameter pointer-to-(function with parameter D returning C) returning pointer-to-A

> Keyword typename must precede any qualified type name that
**depends on a template parameter**

if the keyword doesn't exists, then the compiler won't consider this a type.

### Writing Function Templates

> Make all template parameters deducible (Except when you can't).

if the type is only in the result, then it can't be deduced. we can also not deduce the parent type of a function parameter .

```cpp
template <class Result, class Source>
Result my_fancy_cast(const Source& src) {
// ...
}

template <class T>
void f(typename T::type arg) { }

template <class T> struct A {
  using type = T;
};

template <class T>
void g(typename A<T>::type arg) { }
```

we should avoid complicated overload sets, they are already complex, and adding templates to the mix just makes things harder.

side example, two overload of <cpp>std::vector</cpp> constructor which use either parentheses or curly braces, and behave differently.

we should avoid functions that accept anything and have a simple common name, otherwise we might get something like namespace pollution.

> Function Template Specialization
>
> - What is allowed:
>   - Full specialization of non-member function templates
> - What is not allowed:
>   - Partial specialization of non-member function templates
>   - Any specialization of member function templates

but we shouldn't do this at all. either use a template overload or a non-template overload.
</details>
