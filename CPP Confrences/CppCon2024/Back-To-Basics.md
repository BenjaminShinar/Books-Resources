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
- [x] Back to Basics: Function Call Resolution - Ben Saks
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
we can look at the original definition, a templated class of dynamic array. it had only one templated parameter (no allocator parameter). over the years there were more member functions added, more ways to access element, iterators, allocators, modifiers...

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

// variable template definition
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

#### Writing Class Templates

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

#### Writing Function Templates

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

### Back to Basics: Function Call Resolution - Ben Saks

<details>
<summary>
How does the compiler choose the correct function?
</summary>

[Back to Basics: Function Call Resolution](https://youtu.be/ab_RzvGAS1Q?si=CBGqcANi8BVBoJLz), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Back_to_Basics_Function_Call_Resolution.pdf), [event](https://cppcon2024.sched.com/event/1gZdm/back-to-basics-function-call-resolution).

C++ allows multiple functions with the same name.

- name hiding
- function overloading
- function templates

when the compiler sees a function call, it needs to choose which function to use.

- Function overloading and overload resolution
- Name lookup
- Default function arguments
- Function templates

function overloading is declaring functions with the same name in the same scope.

```cpp
int put(int c);
int put(int c, FILE *f);
int put(char const *s);
int put(char const *s, FILE *f);
```

the compiler matches the number and types of argument in the function call and compares to function parameters in the definitions. the compiler is allowed to perform some conversions to match the types.\
when there are multiple candidates for a function, there is a ranking process.

- best match - "exact"
  - no conversion
  - array-to-pointer
  - qualification conversion (like <cpp>const</cpp>)
- Promotion
  - integral promotion
  - floating point promotion
- Conversion - "expensive"
  - integral conversion
  - floating point conversion
  - pointer conversion
  - boolean conversion

if there is a tie (multiple candidates in the same category), the call is ambiguous and we get an error.

the literal zero (`0`) is the only integer with an implicit conversion to a pointer type (`NULL` is defined as 0). if we had a function call with any other integer (or just a integer variable), it won't consider pointer conversion.\
if we have multiple arguments that differ, there is raking process that matches the options:

```cpp
void f(double x, double y, double z);
void f(double x, int y, double z);

f(1.1, 2, 3); // could work with either function by itself
```

> For a given call, function $F_i$ is a better match than $F_j$ if:
>
> - for every argument $A_k$ in the call, $F_i$'s conversion for $A_k$ is no worse than $F_j$'s conversion for $A_k$, and
> - for at least one argument $A$ in the call, $F_i$'s conversion for A is better than $F_j$'s conversion for $A$

so in our example, the second argument is an exact match, which makes the second option a strictly better match.

the function signature is the function name and parameter type list (also the enclosing namespace). with a member function, it includes the class, cv-qualifiers (<cpp>const</cpp> and <cpp>volatile</cpp>) and ref-qualifiers (reference `&` or temporary `&&`).\
declarations are stored in the compiler symbol table, which is why we sometimes need to forward declare types.

```cpp
class widget { // stores widget
public:
  string name() const; // looks up string; stores name
};
```

#### Name Lookup

> In C++, declarations can appear at:
>
> - local scope: inside a function declaration, including that function's parameter list or a block nested inside a function definition
>   - A name declared at local scope is in scope to the end of the function declaration or block containing that name.
> - class scope: in the brace-enclosed body of a class definition.
>   - A name declared at class scope is in scope to the end of its class definition .and within the parameter list and body of a member definition of the same class.
> - namespace scope: outside of any function, class, structure, or union, whether global or in some other namespace.
>   - A name declared at namespace scope is in scope to the end of its namespace definition, which for the global scope is the end of its translation unit.

Qualified names appear after (to the right) of `::`, `.` or `->`, non-qualified names aren't. the process of looking a qualified name is different. with a qualified name, we first determine what appears before the qualifying symbol:

- namespace - search in the namespace scope
- class name, object or pointer to object - search in the object or class, and if missing, search in the base class (direct first, and then to indirect).

if the lookup doesn't find anything, the compilation fails.\
for unqualified names, the process is different. it behaves like name lookup in classic C

```cpp
namespace S {
  int m, n;
  void f(int n) {
    m = n; // m refers to S::m; n refers to parameter
  }
}
```

the look up starts in the local scope, first in the code block, and then working outward. next the compiler searches in the namespace scope, starting with the current namespace and working outward to the global namespace scope. if the name is inside a function member of a class, it also looks at the class scope and the base class scopes. this happens after the local scope and before the namespaces.\
name lookup takes precedence over overload resolution. the candidates come from the lookup, so even if one candidate would be better than another based on the ranking process, it won't be considered if the name lookup process was matched before reaching that scope.

```cpp
int put(int c); // (1) not found
int put(int c, FILE *f); // (2) not found

namespace N {
  int put(char const *s); // (3) found
  int put(char const *s, FILE *f); // (4) found
  void f(int c) {
    put(c); // (5) error: no valid match
  }
}
```

in this example, the call `put(c)` only considers the functions in the N namespace, and since neither are valid matches (it's only one argument, and a non-zero integer can't be converted to a pointer), it doesn't look at the other options.

```cpp
class Base {
public:
  void f(int n);
};

class Derived: public Base {
public:
  void f(double d); // hides Base::f(int)
};

Derived dx;
dx.f(3);
```

the name lookup concludes when seeing the function in the derived class, even if a 'better' match exists in the base class.

> Argument-Dependent Lookup
>
> - Actually, it's not quite true that overloaded functions "must be in the same scope".
> - There's one more facet of unqualified name lookup to consider: argument dependent lookup (ADL).
> - ADL is specifically for unqualified function names in function calls.
> - ADL adds this name lookup rule:
>   - For each argument in the function call whose type is declared in a namespace, look in that namespace for the function name, as well as in other scopes searched by the usual name lookup.

if there was no argument dependent lookup

```cpp
namespace N {
  class T;
  void f(T &r);
}

N::T x;
f(x); // compile error w/o ADL: never looks in N
N::f(x); // OK w/o ADL: finds f in N
```

because ADL exits, and since type `T` is in the `N` namespace, we also search for the `f()` function in that namespace and we can match the function there. this is important mostly for overloaded operators. lets take the same example and replace the values.

```cpp
// <string>
#include <iosfwd>
namespace std {
  class string;
  ostream &operator<<(ostream &, string const &);
/* ...*/
}

// other file

#include <iostream>
#include <string>

std::string s;
std::operator<<(std::cout, s); // OK w/o ADL
std::cout << s; // compile error w/o ADL
operator<<(std::cout, s); // compile error w/o ADL
```

> Sutter's Interface Principle\
> Sutter (2000) offers some advice for grouping code into namespaces based on what he calls the Interface Principle:
>
> - "For a class X, all functions, including free functions, that both "mention" X and are "supplied with" X are logically part of X, because they form part of the interface of X."
> If you put a class into a namespace, be sure to put of its interface functions into the same namespace.

function declarations found via ADL are considered the same way as functions found via the unqualified name lookup.

#### Default Arguments

functions can have default arguments. the default argument is not part of the signature. when considering the cost of an overloaded function, 'filling in' the default is considered free.

```cpp
void g(double d);
void g(int x, int y = 1);

g(0); // calls g(int,int)
g(0.0); // calls g(double)
```

so in the code above, even if we only have one argument, adding the default doesn't effect the ranking, and it is considered an exact match.

#### Function Templates

a function template isn't a function by itself, it's something that is instantiated into a function during compilation.

```cpp
template <typename T>
constexpr T const &max(T const &a, T const &b) {
  return (a > b) ? a : b;
}
```

we can often omit the template type and rely on template argument deduction.

```cpp
int i, j;
float f, g;

int k = max(i, j); // calls max<int>(i, j)
float h = max(f, g); // calls max<float>(f, g)
```

the compiler doesn't allow converting on deduced template calls, it will only do converting if we explicitly specify the template argument.

> - when performing template argument deduction, the compiler doesn't consider most type conversions.
> - For example, although there's normally a standard conversion from int to
double, the compiler won't convert x into a doublein this call:
>
> ```cpp
> int x = 1;
> double y = 2.5;
> double z = max(x, y); // compile error: can't deduce T
> ```
>
> Instead, the compiler rejects the call.

if the compiler sees a tie between an overloaded function and a function template, it will prefer the non-templated type. we can always force use the templated version if we specify the template and the type argument

```cpp
template <typename T>
constexpr T const &max(T const &a, T const &b);

constexpr char const *max(char const *a, char const *b) {
  return strcmp(a, b) > 0 ? a : b; // not templated
}

char const N[] = "Nancy";
char const D[] = "Dan";
char const *p;
p = max(D, N); // calls non-template max
t = max<char const *>(D, N); // calls template max
```

#### The Two Steps Swap

looking at the <cpp>std::swap</cpp> and customization points, the c++ standard has a generic swap function using move mechanics. but it might be better to have type specific swaps.

```cpp
template <typename T>
void std::swap(T &a, T &b) {
  T temp {std::move(a)};
  a = std::move(b);
  b = std::move(temp);
}

namespace Saks {
  class string {
    friend void swap(string &a, string &b);
  private:
    std::size_t stored_length;
    char *actual_str;
  };

  // probably better than default swap
  void Saks::swap(string &a, string &b) {
    std::swap(a.stored_length, b.stored_length);
    std::swap(a.actual_str, b.actual_str);
  }
}
```

however, it's still possible for a user to accidentally use the regular swap rather than the specialized one.

```cpp
class Person {
public:
  void swap(Person &other);
private:
  Saks::string name;
  unsigned idnum;
};

// explicitly using the std::swap
void Person::swap(Person &other) {
  std::swap(name, other.name); // oops, uses std::swap
  std::swap(idnum, other.idnum);
}

// better version
void Person::swap(Person &other) {
  using std::swap;
  swap(name, other.name); // OK, uses Saks::swap
  swap(idnum, other.idnum);
};
```

the mistaken version uses qualified names <cpp>std::swap</cpp>, the other version brings in the function block with the `using` statement, and then we allow the function resolution to work. the argument dependant lookup will bring the SAKS namespace into the mix, and the non-template function would be preferred.

</details>
