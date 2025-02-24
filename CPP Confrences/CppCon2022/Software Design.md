<!--
ignore these words in spell check for this file
// cSpell:ignore rror  שלום עולם libletized configurability TFoos cppcoro liblets templating Coplien irange Gitter absl
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Software Design

### How C++23 Changes the Way We Write Code - Timur Doumler

<details>
<summary>
four big changes in c++23: stronger deduction of this, std::expected (a powerhouse std::optional), std::mdspan (non-owning multiple dimension span), and std::print (a cleaner way to print to console).
</summary>

[How C++23 Changes the Way We Write Code](https://youtu.be/eD-ceG-oByA),

C++23 is feature complete, so we now know the scope of C++23 and what will be included in it.

the big features of c++20:

- coroutines
- concepts
- ranges
- modules

c++23 doesn't have features at that magnitude, it is a smaller release. the changes are divided into two parts: core language features and library features. we can see them up on [cpp reference](https://en.cppreference.com/w/cpp/23). everything marked with **DR** is "defect report", basically a bug fix. there are some deprecations and simplifications, and a lot of improvements and extensions of the ranges and views, and other changes which tackle problems with text, unicode and character sets. some other changes are there to remove boiler plate code and make writing easier. (he goes on to remove some more features from the list, one by one, each for a different reason).

we end up with four changes.

- Deducing `this`
- _std::expected_:
- _std::mdspan_: a non-owning multidimensional array reference
- _std::print_: formatted output library

> "It's a good feature proposal if it solves at least two unrelated problems simultaneously."\
> ~ Bjarne Stroustrup

#### Deducing `this`

A core language feature, this one solves at least three problems, and maybe even more.

overloads are possible on arguments, on const, on whether the object is an LValue or RValue reference, and on the combinations. it means that we change the function based on the `this` argument. in c++23 there is a clearer way to express this, by explicitly writing the first _this_ parameter.

```cpp
// c++20
struct X {
  T& value();
  const T& value() const;
};

struct Y {
  T& value() &;
  T&& value() &&;
};

struct Z {
  T& value() &;
  const T& value() const &;
  T&& value() &&;
  const T&& value() const &&;
};

// c++23
struct XYZ {
  T& value(this X& self);
  const T& value(this const X& self);
  T&& value(this X&& self);
  const T&& value(this const X&& self);
};
```

we can also declare it as a template. using the "normal" rules of function template parameters deduction, with or without declaring the template parameter.

```cpp
struct X {
  template <typename Self>
  auto && value(this Self&& self);
};

// same as
struct Y {
  auto && value(this auto&& self);
};
```

it solves a problem of duplications. such as for the _optional_ struct, which requires the same body in all overloads. and trying to remove the duplicated code requires writing static_casts, writing weird private implementation or omitting some of them. in c++23 it's easier

```cpp
template <typename T>
struct optional {
  template <typename Self>
  constexpr auto&& value(this Self&& self) {
    if (!self.has(value))
      throw bad_optional_access();

    return std::forward<Self>(self).m_value;
  }
};
```

the second problem that is solved is simplifying CRTP (curiously recurring template pattern). classic CRTP allows us to create "mixins", a templated class that is based on the derived class that is passed to it. (he shows an example of a bug by passing the wrong template parameter to the base). in c++23, the `this` keyword will be deduced into the most derived class. (there are some cases which still requires templating when we use friend fuctions, and we can still have bugs if we try to have runtime polymorphism), so there are still stuff to consider. we can combine the new `this` with concepts.

```cpp
struct Incrementable
{
  //prefix
  auto & operator++(this auto& self)
  {
    self.setValue(self.getValue() + 1);
    return self;
  }
  // postfix
  auto operator++(this auto& self, int)
  {
    auto tmp = self;
    self.setValue(self.getValue() + 1);
    return tmp;
  }
};
```

a third problem solved is recursive lambdas. the naive approach fails, some workaround are possible using _std::function_ or passing the function to itself, but c++23 makes it easier to do.

```cpp
// naive approach

void foo()
{
  auto f =[](int i){
    if (i == 0)
      return 1;

    return i *f(i-1); // can't do this.
  };
  std::cout << f(5) << '\n';
}

// workaround using std::function
void foo2()
{
  std::function<int(int)> f =[&](int i){
    if (i == 0)
      return 1;

    return i *f(i-1); // this works,
  };
  std::cout << f(5) << '\n';
}

// workaround with self parameter
void foo3()
{
  auto f =[](auto&& self,int i){
    if (i == 0)
      return 1;

    return i *self(self,i-1); // this works, but weirdly
  };
  std::cout << f(f,5) << '\n'; // we pass function to itself
}

// c++23 - explicating `this`
void foo3()
{
  auto f =[](this auto&& self,int i){
    if (i == 0)
      return 1;

    return i *self(i-1); // this works,
  };
  std::cout << f(5) << '\n';
}
```

an actual example is for counting leafs in a tree. we use a variant, operator overload, and the visitor pattern. (this is great for code interviews, if you have a cutting edge compiler)

```cpp
struct Leaf{};
struct Node;
using Tree = std::variant<Leaf, Node*>;
struct Node {
  Tree left, right;
};

template <typename... Ts>
struct overload : Ts ...{using Ts::operator()...; }

int countLeaves(const Tree& tree) {
  return std::visit(overload {
    [] (const Leaf&) {return 1;},
    [] (this const auto& self, const Node* node) -> int {
          return visit(self, node->left) + visit(self, node->right);
      }
  }, tree);
}
```

#### _std::expected_

an example of a text parser which uses _std::strtod_ to parse doubles, but it is designed to return 0 for non numeric data, we can use exceptions to catch the possible issues, but exceptions have their costs, and should be used when we have fatal errors, the overhead is large because of the stack-unwinding. and they aren't suited for deterministic or memory-limited environments. if we use error codes rather than exceptions, we then have to change the signature of the code, or we can use _std::optional_, but it doesn't provide all the information. the _std::expected_ allows for more possibilities. the error type can be whatever we want (none void types).

```cpp
std::expected<double, parse_error> parse_number(std::string_view& str) {
  const char* begin = str.data();
  char * end;
  double retval = std::strtod(begin, &end);

  if (begin == end)
    return std::unexpected(invalid_char);
  if (retval == HUGE_VAL)
    return std::unexpected(overflow);

  str.remove_prefix(end-begin);
  return retval;
}

//call site

void foo() {
  std::string_view src("meow");

  auto num = parse_number(src);

  if (num) {
    std::cout << *num << '\n';
  }
  else if (num.error() == invalid_char) {
    // handle ...
  }
  else if (num.error() == overflow) {
    // handle ...
  }
}
```

because it is a vocabulary type integrated into the language, we can use it across API boundaries.

in this example we want to make sure we never get an overflow, so we have to write into the accumulate operator itself, we check for overflow and short-circuit it. but if we want to avoid costly exceptions, we use the _std::expected_ type instead. there is some weird parts of the declaration.

```cpp
// library code - c++20
template <typename InputIt, class T, class BinaryOperation>
T accumulate(InputIt first, InputIt last, T init, BinaryOperation op) {
  for ( ; first != last ; ++first) {
    init = op(std::move(init),*first);
  }
  return init;
}

//user code - c++20
void callerCpp20() {
  std::vector<int> vec = {/*...*/};
  auto sum = accumulate(vec.begin(), vec.end(),0, [](int a, int b) {
    if (b > 0 && a > INT_MAX - b) throw signed_integer_overflow();
    if (b < 0 && a < INT_MIN - b) throw signed_integer_underflow();
    return a+b;
  });
}

/******/
// library code - c++23
template <typename InputIt, class T, class BinaryOperation>
auto accumulate(InputIt first, InputIt last, T init, BinaryOperation op) -> std::expected<T, typename std::remove_cvref_t<std::invoke_result_t<BinaryOp, T, typename std::iterator_traits<InputIt>::value_type>>::error_type> {
  for ( ; first != last ; ++first) {
    auto result = op(std::move(init),*first);
    if (!result) return std::unexpected(result.error());
    init = *result;
  }
  return init;
}


// user code - c++23
void callerCpp23() {
  std::vector<int> vec = {/*...*/};
  auto sum = accumulate(vec.begin(), vec.end(),std::expected<int, numeric_error>, [](int a, int b) {
    if (b > 0 && a > INT_MAX - b) return std::unexpected(signed_integer_overflow);
    if (b < 0 && a < INT_MIN - b) return std::unexpected(signed_integer_underflow);
    return a+b;
  });
}
```

it's also possible to use the _std::variant_ as a type of the _std:;expected_ struct. this allows us to collate different errors from different layers, and then we can use the visitor pattern as before. it looks like a try-catch block, but without the overhead from exception handling. if we use this pattern, _std::visit_ will have a compiler error if we add a type to the variant type.

```cpp
auto getWidget() -> std::expected<Widget, std::variant<ParseError,IOError>>;

void foo() {
  auto widget = getWidget();
  if (widget) {
    process(*widget);
  }
  else {
    std::visit(overload{
      [] (ParseError& error) {
        // handle parse error...
      },
      [] (IOError& error) {
        // handle I/O error...
      },
    }, widget.error());
  }
}
```

in c++26, we might get more power to _std::expected_ and make it more composable.

#### _std::mdspan_

a non-owning multidimensional array reference,

multidimensional array in fortran, where they are built-in into the language, and easy to use. and even dynamic arrays are easy.

```fortran
INTEGER, PARAMETER :: NX = 4, NY = 4, NZ = 2 ! ARRAY EXTENTS
REAL, DIMENSTION(:,:,:), ALLOCATABLE :: DATA ! THE ACTUAL ARRAY
INTEGER :: I, J, K

ALLOCATE(DATA(NX,NY,NZ))

DO K = 1, NZ
  DO J = 1, NY
    DO I = 1, NX
      DATA(I, J, K) = I + J + K
    END DO0
  END DO0
END DO0

DATA = DATA * 0.1
WRITE(*,*) DATA
```

in c the code looks like this if we have known dimensions,

```c
const int nx = 4, ny = 4, nz =2;
double data[nx][ny][nz];

for (int i = 0; i < nx; i++)
  for (int j = 0; j < ny; j++)
    for (int k = 0; k < nz; k++)
      data[i][j][k] = i+j+k;

```

but for dynamic sizes, we need to allocate each dimension correctly, and de-allocate it.

```c
const int nx = 4, ny = 4, nz =2;

// allocation
double*** data = (double***)malloc(nx * sizeof(double**));
for (int i = 0; i < nx; i++) {
  data[i] = (double**)malloc(ny * sizeof(double*));
  for (int j = 0; j < ny; j++)
  {
    data[i][j]=(double*)malloc(nz * sizeof(double));
  }
}

for (int i = 0; i < nx; i++)
  for (int j = 0; j < ny; j++)
    for (int k = 0; k < nz; k++)
      data[i][j][k] = i+j+k;

// deallocation
for (int i = 0; i < nx; i++) {
  for (int j = 0; j < ny; j++)
  {
    free(data[i][j])
  }
  free(data[i]);
}
free(data);
```

although the better way was to have continuos data in memory and avoid fragmentation and multiple dereferencing. then we just need to be smart with pointer arithmetics.

```c
const int nx = 4, ny = 4, nz =2;

// allocation
double* data = (double*)malloc(nx * ny * nz * sizeof(double));

for (int i = 0; i < nx; i++)
  for (int j = 0; j < ny; j++)
    for (int k = 0; k < nz; k++)
      *(data + (i * ny * nz) + (j * nz) + k) = i + j + k;

// deallocation
free(data)
```

in c++ we can put everything into a class and hide some overhead, this involves templating for each dimension, and then templating to allow any number of dimensions, and then trying to get some optimizations for known dimensions, which devolves into a mess of a code. _std::mdspan_ hides this behavior as part of the standard library.\
we also have the multiple argument subscript operator (square brackets `[]`). it's a wrapper over the data, so it doesn't own it, and we can have multiple wrapper over the same data with different dimensions.

```cpp
int nx = 4, ny =4, nz =2;
double* data = /*...*/; //memory block

std::mdspan data3d(data, nx, ny, nz);
for (int i = 0; i < nx; i++)
  for (int j = 0; j < ny; j++)
    for (int k = 0; k < nz; k++)
      data3d[i,j,k] = i + j + k;

std::mdspan data2d(data, nx * ny, nz);
for (int i = 0; i < nx * ny; i++)
    for (int k = 0; k < nz; k++)
      data2d[i,k] = i + k;

data3d.rank(); // returns 3
data2d.rank(); // returns 2
```

we can have **extents** to mix static and dynamic dimensions, and some helpers. we can also determine the layout policy, such as row-major (_std::layout_right_, usual for c++ arrays) and column-major(_std::layout_right_, usual for fortran arrays), and this policy is a customization point, so we can have tiled layout for images. there is an accessor policy, which defaults to returning a reference, but can be customizable.\
in c++26, we might get _std::mdarray_ as multidimensional array.

#### _std::print_

changing how we write the **Hello World** example.

```cpp
#include <iostream>

int main() {
  std::cout << "Hello World\n";
}
```

console out doesn't work perfectly with non utf8 characters, it just forwards the data to the output, it also requires extra work to make boolean values print as "true" and "false", otherwise it prints 1 or 0. and it support common format specifiers. it checks for the number of arguments and helps us avoid errors.\
The down side is that _std::format_ doesn't actually print, it just creates a string object, which we need to pass to the console output somehow.

```cpp
#include <format>

int main() {
  bool b = true;
  std::format("שלום עולם! b == {}\n", b); // "true"
  std::format("b == {:d}\n", b); // 1
  std::format("b == {:d}\n"); // error, missing argument
  std::cout << std::format("actual print {b}\n",b); // print to screen
}
```

c++23 introduces _std::print_, which does the work for us, and calls the correct output stream api to deal with unicode. and it allows us to determine which stream we want to print to.

```cpp
#include <print>

int main() {
  bool b = true;
  std::print("שלום עולם! b == {}\n", b); // "true"
}
```

so now the "hello world" looks like this, we can even use _std::println_ to remove the '\n'.

```cpp
import std;

int main(){
  std::println("Hello, World!")
}
```

</details>

### How Microsoft Uses C++ to Deliver Office - Huge Size, Small Components - Zachary Henkel

<details>
<summary>
Exploring Microsoft Office and how it solved dependencies
</summary>

[How Microsoft Uses C++ to Deliver Office - Huge Size, Small Components](https://youtu.be/0QtX-nMlz0Q), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/How-Microsoft-Uses-C-to-Deliver-Office-Zachary-Henkel-CppCon-2022.pdf).

#### Background

some history:

> - 1983: Word for DOS released. Written in C
> - 1985: Word for Classic Mac OS released
> - 1990: First release of Office suite
> - 1992: PowerPoint moves to C++
> - 1995: Common functionality moved to shared library (mso.dll)
> - 1996: Office moves to C++
> - 2001: First release for macOS (OS X)
> - 2010: Office for the Web
> - 2013: Office on iOS & Android
> - 2019: Clang analysis for Windows code

operation system (windows, apple, android), architectures(arm, arm64, x64, x86).

#### Huge Size

office is a huge codebase, 350 million lines code, native code. the common metric for codebase size is number of lines of code, which is skewed, based on the preprocessor.

the ideal measure would be the number of unique translation units that are used in the binaries. which is a hard thing to measure.\
instead, we can count compilations:

- platform
- architecture
- debug/ship

costs of size:\
it is a really big workload to build, it can cost time and hardware. but if we want to run static analysis, then we need to build for that as well. it's hard to migrate between compilers and language versions. everything needs to be tested, on each platform, and hope there are no regressions.\
there is also a problem of 'decommissioning' code, it's hard to get rid of code that is ingrained in the codebase, and there might be residuals.

#### Small Components

liblets - the small pieces of common code. like **MSO.dll** in the past (dialogs, help, asserts). but of course, the dll grew and exploded in complexity, so it was hard to work with, everything depended on everything, and encapsulation was hard. no way to differentiate between private and public code.

in 2010, office started using liblets, with two goals:

1. develop of philosophy for how code should be structured.
2. break up the monolith.

Layers: it's preferred to add code at a higher level. each layers is a group of shared libraries. low level layers can't use high level layers.

- mso - Non-libletized code
- mso98 - Clean liblets
- mso40ui - UI Frameworks, Graphics
- mso30 - Document synchronization: Authentication, File I/O, Identity
- mso20 - Core functionality: Diagnostics, Experimentation, Telemetry

pillars of liblets:

1. modern c++
2. distinct public api
3. self contained
4. clean dependencies
5. unit tested

modern c++ started at 2011, with the new c++11 standard, exception safe code, using the STL and being modern. no more writing C code in C++ compiler.

headers are marked for public consumption, this is enforced by the build system. each translation unit should work in isolation.

symbol visibility. exporting, mangled names.

when using liblets, the consumer decides on which endpoints it uses. **dependencies validation**, stricter than the linker, and making sure that only public API are used.

liblets are unit tested, mocks are automatically generated.

#### The Future

migrating to header units - a step towards modules, more flexible than pre compiled headers.

header units follow similar ideas to liblets:

- self contained headers
- well defined, acyclic dependencies
- no conditional compilation

</details>

### The Hidden Performance Price of C++ Virtual Functions - Ivica Bogosavljevic

<details>
<summary>
Measuring and demonstrating the effects of virtual functions, and how they can be addressed.
</summary>

[The Hidden Performance Price of C++ Virtual Functions](https://youtu.be/n6PvvE_tEPk), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-The-Hidden-Performance-Price-of-Virtual-Functions.pdf).

#### Introduction

> **Virtual functions in C++**
>
> - Enable flexibility
> - The basic component of OOP
> - Virtual functions are slower than regular functions
> - The performance price of virtual functions depends on several factors

the way virtual functions are implemented is compiler dependant, but most choose to use virtual tables. which is a table with pointers to the actual function. this means that each object has pointer to the the table, and in runtime, a virtual function call must first go to the virtual table, and then find the correct function to use. so there are two layers of indirection.

so we can see that virtual functions are more expensive than non-virtual. because the addresses are unknown at compile time.\
in a simple benchmarking, the overhead is relevant for short and fast functions, but for complex functions, the overhead is negligible.

however, if we have vector of pointers, we need to access the object itself, and if we have the objects laid out in a non-optimal way (objects are neighbors, like the pointers), then we might get cache misses for accessing the object for the pointer dereferencing to access the virtual table.\
A vector of objects is contiguous, and accessing each object is sequential and has better memory locality. however, a vector of objects doesn't allow for polymorphism.

we can experiment with this by starting with a perfectly ordered vector of pointers, and for each iteration, we can swap the pointers and shuffle them. we see degradation, but it's not because of the virtual functions, the same slow down occurs with non-virtual functions. the problem is the memory layout.

#### Solutions and Optimizations

> Alternatives to vector of pointers:
>
> - Use `std::variant` with `std::visitor`
> - Use polymorphic_vector - uses virtual dispatching, but doesn’t uses pointers. Downside is increased memory consumption → google `polymorphic_vector`
> - Use per type vector (e.g. `boost::base_collection`), a very useful if you don’t need a specific ordering in the vector

The compiler can help us with optimizations, it can sometimes inline the function call, and then skip the function call, and even perform further optimizations (unrolling loops, loop un-switching, loop invariants).

**Type Based Processing**

> - Don’t mix the types, each type has its own container and its own loop
> - The compiler can inline small functions and perform the compiler optimizations
> - Already implemented in `boost::base_collection`
> - This approach is applicable if objects in the vector don’t have to be sorted

some code benefits more than other from inlining.

**Jump Destination Guessing** (_speculative execution_), one virtual function is guessed to be the correct one, and the execution starts. if the function is determined to be another one, then there a flush and this has performance costs.\
we can see this in action by changing the order of the objects. this is what type-based processing does.

**Instruction Cache Evictions**

> Modern CPUs rely on “getting to know” the instructions they are executing
>
> - The code that has already been executed is hot code
>   - Its instructions are in the instruction cache
>   - Its branch predictors know the outcome of the branch (true/false)
>   - Its jump predictors know the target of the jump
> - The CPU is faster when executing hot code compared to executing cold code
> - The CPU’s memory is limited
>   - The code that is currently hot will eventually become cold unless executed frequently
> - Virtual functions, especially large virtual functions where each object has a different virtual function, mean that we are switching from one implementation to another
>   - The CPU is constantly switching between different implementations and is always running cold code

it's hard to demonstrate this effect, it depends on many factors, and it's not related to virtual functions themselves. but when we have a collection of virtual objects, it is more likely to occur.

#### Conclusion

> - Virtual functions do not incur too much additional cost by themselves
> - It is the environment where they run which determines their speed
> - The hardware craves predictability: same type, same function, neighboring virtual address
>   - When this is true, the hardware run at its fastest
>   - It’s difficult to achieve this with casual usage of virtual functions
> - In game development, they use another paradigm instead of OOP called: data-oriented design
>   - One of its major parts is type based processing: each vector holds one type only
>     - This eliminates all the problems related to virtual functions
>     - However, this approach is not applicable everywhere
> - If you need to use virtual functions, bear in mind:
>   - The number one factor that is responsible for bad performance are data cache misses
>     - Avoiding vector of pointers on a hot path is a must!
>   - Other factors also play their role, but to a lesser extend
>   - With careful design, you can reap most benefit of virtual functions without incurring too much additional cost
> - Here are a few ideas to fix your code with virtual functions:
>   - Arrangement of objects in memory is very important!
>   - Try to make small functions non-virtual!
>     - Most overhead of virtual functions comes from small functions, they cost more to call than to execute
>   - Try to keep objects in the vector sorted by type

</details>

### Using Modern C++ to Eliminate Virtual Functions - Jonathan Gopel

<details>
<summary>
Replacing virtual functions in modern c++, Using concepts to avoid interfaces and inheritance. 
</summary>

[Using Modern C++ to Eliminate Virtual Functions](https://youtu.be/gTNJXVmuRRA), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Using-Modern-C-to-Eliminate-Virtual-Functions-Jonathan-Gopel-CppCon-2022.pdf)

> When is virtual useful
>
> - Requiring a specific interface
> - Adding configurability to objects
> - Holding multiple different derived types with a shared
>   base class in a single container

reasons to replace virtual functions:

- less indirection
- move errors to compile time and get static behavior
- greater flexibility in design (sometimes)
- possible performance improvement

we don't want to re-create the vtable(std::any, std::variant) or focus on type erasure.

#### Binding interfaces

starting with an example, a base class with a virtual function. however, if we need the virtual destructor, then now we need to have ugly boilerplate code in both of the structs(rule of five).

```cpp
struct FooInterface {
  [[nodiscard]] virtual auto func() const -> int = 0;

  FooInterface() = default;
  FooInterface(const FooInterface&) = default;
  FooInterface(FooInterface&&) = default;
  FooInterface& operator=(const FooInterface&) = default;
  FooInterface& operator=(FooInterface&&) = default;
  virtual ~FooInterface() = default;
};

struct Foo final : public FooInterface {
  [[nodiscard]] auto func() const -> int override {
  return 42;
  }

  Foo() = default;
  Foo(const Foo&) = default;
  Foo(Foo&&) = default;
  Foo& operator=(const Foo&) = default;
  Foo& operator=(Foo&&) = default;
  virtual ~Foo() = default;
};
```

we can use concepts instead (and SFINAE in pre c++20 code), we define the function in the concept, and then assert that our object adheres to it. another bonus of concepts is that we can use more loose return types, such as std::integer rather than int.

```cpp
concept CFoo = requires(T foo) {
  {foo.func()} -> std::same_as<int>;
  //{foo.func()} -> std::integer;
};

struct Foo {
  [[nodiscard]] auto func() const -> int {
  return 42;
  }
};
static_assert(CFoo<Foo>);
```

with virtual functions, we need to create a unique pointer, and pass it to the function, while using concepts only requires us to specify the concept in the function signature.

```cpp
// with virtual
std::unique_ptr<FooInterface> foo = std::make_unique<Foo>();
auto func(std::unique_ptr<FooInterface> foo2) {
  // Implementation here
}
// without virtual
Foo foo{};
auto func(CFoo auto& foo2) {
  // Implementation here
}
```

#### Owning a Polymorphic Type

a type that holds/owns the polymorphic type (at compile type), with virtual functions, we just store the pointer.

```cpp
class Bar {
  public:
  constexpr Bar(std::unique_ptr<FooInterface> input_foo)
  : foo{std::move(input_foo)} {}
  private:
  std::unique_ptr<FooInterface> foo{};
};
```

without virtual functions, we use templates, the downside is that we can't change the internal type at runtime.

```cpp
template <CFoo TFoo> // TFoo must adhere to CFoo concept
class Bar {
  public:
  constexpr Bar(TFoo input_foo): foo{input_foo} {}
  private:
  TFoo foo{};
};
```

if we want to manipulate this at runtime, then the virtual example gets a 'set' function that moves the input inside the object.

```cpp
class Bar {
  public:
  constexpr Bar(std::unique_ptr<FooInterface> input_foo) : foo{std::move(input_foo)} {}
  constexpr auto set_foo(std::unique_ptr<FooInterface> input_foo) {
    foo = std::move(input_foo);
  }
  private:
  std::unique_ptr<FooInterface> foo{};
};
```

for the non virtual implementation, we need to use std::variant, and we need to start using the auto keyword, we also create a new concept "same_as_any" to further constrain the input types.\
this isn't a perfect design, std::variant has performance overhead, and we would prefer to always know the types.

```cpp
template<typename T, typename... Ts>
concept same_as_any = (... or std::same_as<T, Ts>);

template <CFoo... TFoos>
class Bar {
  public:
  constexpr Bar(same_as_any<TFoos...> auto input_foo): foo{input_foo} {}
  constexpr auto set_foo(same_as_any<TFoos...> auto input_foo) -> void {
    foo = input_foo;
  }
  private:
  std::variant<TFoos...> foo{};
};
```

when we use the non-virtual option, we just provide the allowed types if we decide that the object might need to use a different type in a runtime.

```cpp
// with virtual
Bar bar{std::make_unique<Foo>()};
// without virtual
Bar bar{Foo{}}; // single type allowed
Bar<Foo1, Foo2> bar{Foo1{}}; // multiple types
```

#### Storing Multiple Types

having a container store multiple types, with virtual functions, we just have a base class and use vector of pointers.

```cpp
class Baz {
  public:
  auto store(std::unique_ptr<FooInterface> value) -> void {
    data.push_back(std::move(value));
  }
  private:
  std::vector<std::unique_ptr<FooInterface>> data{};
};
```

> Desired properties
>
> - List of all the types that might be stored
> - Container that can hold many different types simultaneously
> - Container that can hold multiple objects of a single type

for the non-virtual option, we use a tuple that stores vectors. (a two layers storage - we have a different vector for each type). we again use the "same_as_any" concept.we pay for the empty vectors, but the order is not stable between the types (only inside the vectors). this should give us static locality (which is good for performance), but we might be in trouble if wanted the insertion order.

```cpp
template <typename T, typename... Ts>
concept same_as_any = (... or std::same_as<T, Ts>);

template <same_as_any<TFoos...> T>
template <CFoo... TFoos>
class Baz {
  public:
  auto store(T value) {
    return std::get<std::vector<T>>(data).push_back(value);
  }

  private:
  std::tuple<std::vector<TFoos>...> data{};
};
```

the usage is pretty simple, we define the expected types when we create the storage class, this does mean that ww must know them ahead of time, and it can make things harder to extend.

```cpp
// with virtual
Baz baz{};
baz.store(std::make_unique<Foo1>());
baz.store(std::make_unique<Foo2>());
// without virtual
using foo_storage_t = Baz<Foo1, Foo2>;
foo_storage_t baz{};
baz.store(Foo1{});
baz.store(Foo2{})
```

---

> Review
>
> - Concepts bind interfaces
> - Deduced class templates provide compile-time configurability of contained objects
>   - Runtime configurability can be achieved with std::variant if absolutely needed
> - Clever use of type lists and containers allows for statically typed storage of multiple types simultaneously – design will vary by use case
>
> Downsides
>
> - Increased translation unit size
> - Potential increase to binary size
> - May increase compile time
> - May add complexity

#### Example - Monitor Network Devices

another example, this time we want to

> Task:
>
> - We want to monitor some set of devices on the same network that we are on
> - Each device type is unique in how we must interact with it
> - It is not possible to know the device's connection information before we join the network - we must find it in-situ
>
> Design considerations:
>
> - Device detection
>   - Easiest to find all devices of a single type at once
>   - One scan per device type
> - Device state monitoring
>   - Need to allow each device type to have different communication mechanisms
>   - Want to update state only on-command to avoid network overhead

with virtual functions, we define an interface, it has a instance member function 'update'. unfortunately, we can't have a virtual static function (class function), so it must be statically and we can't override it. the device manager (the monitor) gets all the devices and can call the update function on them. it uses the specific device type to call the specific function to find all the devices. which means we have repeated code (we could use a macro), and we need to be careful from making copy-paste errors.

```cpp
class DeviceInterface;
using device_list_t = std::vector<std::unique_ptr<DeviceInterface>>;
class DeviceInterface {
public:
  //[[nodiscard]] static virtual auto find_in_env() -> device_list_t = 0; // can't have a virtual static function! too bad!
  virtual auto update() -> void = 0;
};

class Switch final : DeviceInterface {
public:
  [[nodiscard]] static auto find_in_env() -> device_list_t {
  // Some device finding logic
  }
  auto update() -> void override { /* Update is_on */ }
  private:
  bool is_on{false};
};

class Dimmer final : DeviceInterface {
  public:
  [[nodiscard]] static auto find_in_env() -> device_list_t {
  // Some device finding logic
  }
  auto update() -> void override { /* Update brightness */ }
  private:
  uint_fast8_t brightness{0};
};

class DeviceManager {
  public:
  DeviceManager(device_list_t devices_) : devices{std::move(devices_)} {}
  auto update() -> void {
    for (auto &device : devices) {
      device->update();
    }
  }

  [[nodiscard]] static auto get_devices() -> device_list_t {
    device_list_t output{};

    { // Switch
      auto device_list = Switch::find_in_env();
      output.insert(std::end(output),
      std::make_move_iterator(std::begin(device_list)),
      std::make_move_iterator(std::end(device_list)));
    }

    { // Dimmer
      auto device_list = Dimmer::find_in_env();
      output.insert(std::end(output),
      std::make_move_iterator(std::begin(device_list)),
      std::make_move_iterator(std::end(device_list)));
    }
    return output;
 }
private:
  device_list_t devices{};
};

//usage
auto main() -> int {
  DeviceManager manager(DeviceManager::get_devices());
  manager.update();
}
```

the non-virtual implementation, we write a concept for a device and not an interface, but here we can define the static function for the type. we can also work directly with the types, rather than pointers. when we update, we call a function on each vector in the tuple (so we don't have strict ordering), the function to retrieve all the devices can't be called from outside the class, it must be inside.

```cpp
template <typename T>
concept CDevice = requires(T device) {
  {T::find_in_env()} -> std::same_as<std::vector<T>>;
  {device.update()} -> std::same_as<void>;
};

class Switch {
  public:
  [[nodiscard]] static auto find_in_env() -> std::vector<Switch> {
    // Some device finding logic
  }
  auto update() -> void { /* Update is_on */ }
  private:
  bool is_on{false};
};

class Dimmer {
public:
  [[nodiscard]] static auto find_in_env() -> std::vector<Dimmer> {
  // Some device finding logic
  }
  auto update() -> void { /* Update the brightness */ }
  private:
  uint_fast8_t brightness{0};
};

template <CDevice... TDevices>
class DeviceManager {
  public:
  DeviceManager() : devices(get_devices()){} // get all devices
  auto update() -> void {
    std::apply([this](auto &... device_lists) {
    (update_device(device_lists), ...);
    }, devices);
  }
  private:
  using device_list_t = std::tuple<std::vector<TDevices>...>;
  auto update_device(auto &device_list) -> void {
    for (auto &device : device_list) {
      device.update();
    }
  }

  [[nodiscard]] static auto get_devices() -> device_list_t {
     return std::tuple{TDevices::find_in_env()...};
  }

  device_list_t devices{};
};

//usage
auto main() -> int {
  DeviceManager<Switch, Dimmer> manager{}; // the function of getting the device is called inside the constructor
  manager.update();
}
```

</details>

### 10 Tips for Cleaner C++ 20 Code - David Sackstein

<details>
<summary>
Creating an application that follows clean code principles.
</summary>

[10 Tips for Cleaner C++ 20 Code](https://youtu.be/9ch7tZN4jeI), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Your-Compiler-Understands-It-But-Does-Anyone-Else.pdf), [code sample github](https://github.com/david-sackstein/clean-it).

> Your Compiler Understands it, But Does Anyone Else?

the STL is great, but it's hard to read and understand, the classes and functions are long, many templates, many paths inside the functions, long header files. this isn't what we define as "Clean Code".

#### Introduction - Clean Code in Context and Best Practices for C++ Programming

there is a trade-off that affect clean code, depending on the layer we are in, the more stable, backward compatible and the reviewers are more expert when we think about core library c++ code, but in higher layers (library code, component code), the readability is more important, changes are more frequent.

clean code has different requirements at those different layers.

The SOLID Principles are general for OOP code.

> - Single Responsibility Principle - small pieces that do one thing.
> - Open-Closed Principle - define abstractions for extensibility. (_"Open for extension, closed for modification"_)
> - Liskov Substituion Principle - be careful with what you inherit.
> - Interface Segregation Principle - like SRP but with interfaces
> - Dependency Inversion Principle - depend on abstractions.

and in C++, we have the core guidelines.

#### VOD Application

> Develop an Application and Clean It.

The sample application will be a VOD application (video on demand), it accepts connections from clients, provides a list of available movies and streams the media. the client can stop the video or can disconnect. we will build the client and server, as well as the MovieReader to read movies from disk.

we use the built in compiler warning and tools such as a clang-tidy and resharper. we want the highest level of warnings and inspection data.

> Managing headers
>
> - A header should include all the headers it depends on.
> - A header should NOT include headers it does not depend on.
> - Use pimpl or an interface to represent dependencies to avoid needing to include implementation headers.
> - Separate internal headers from API headers.
> - See the implementation at git tag: v2-header-ordering

in C++20 we can use modules as an alternative to headers. module explicitly export data, but guidelines will still remain.

Comments that tell us what the code is doing are probably bad, and there is the downside of error handling with exceptions. they make the code harder read, they hide the actual business data, and they aren't fit for use in embedded environment. exceptions should only be used for exceptional problems, bad input, bad files and most other "errors" aren't exceptional, they are just problems which can happen. In c++23 we will have _std::expected<T,E>_ that will help us.

> Improving error handling
>
> - Advantages of exceptions
>   - Allow programmers to focus on the business logic.
>   - Exceptions require immediate and exclusive attention.
> - Disadvantages
>   - The exceptional path is slow (not necessarily a problem)
>   - Local handling litters the code with try/catch clauses.
>   - Not always supported in free-standing environments.
> - Proposed guideline:
>   - Use exceptions for the exceptional
>   - Use expected for the expected.

We should also extract code into functions. we can use functional syntax ("and_then") to compose "railways" or "paths" of code execution. each section does one thing, just as the SRP Principle. we can also use lazy iteration (generator and coroutine), which saves us a bit more memory.

> Use generators for lazy iteration
>
> - Lazy evaluation avoids storing all elements in memory.
> - Generators can be used to invert dependencies
> - Use a library rather then reinvent
> - The famous cppcoro library by Lewis Baker is not maintained. (there are some forks which are)

if we have an exception, we simply catch and convert it immediately.

```cpp
template <typename F>
auto expect(F&& func) noexcept -> std::expected<std::invoke_result_t<F>>{
  try {
    return std::invoke(std::forward<F>(func));
  }
  catch (std::exception& ex) {
    return std::unexpected(ex.what());
  }
}
```

there is a bias toward using primitives, which isn't always a good thing to do. types help us validate ranges and give us non-ambiguous defintions, this way we can ensure that all object are valid.

> Avoid Primitive Obsession
>
> - Primitive Obsession is a code smell in which primitive data types are used excessively to represent data models.
> - Movie uses an int to represent seconds.
> - Movie itself does not validate its invariants (the duration must be within a certain range)
> - We can fix this by hiding the constructor and use expect() on static factory method.

next we build the VodServer and VodClient. we will use a weak pointer to refer to the client and we will use the observer pattern. we have a state machine using a flag. we want to avoid inheriting from implementation if we don't have to. so we start refactoring:

- minimizing the api by defining interfaces.
- moving the defintions to another file.
- using the factory pattern to return the interface.
- extracting methods to have them do one thing alone.
- separate test code and production code.

Dependency injection - taking interfaces as arguments to the constructor, so we can mock our code for testing.

> Dependency Injection
>
> - The objective is to decrease coupling between components and their implementations.
> - Components should depend on abstractions not concrete types.
> - C++ provides two abstraction models:
>   - Compile time using templates.
>   - Run time using virtual functions.
>
> Constructor Injection
>
> - A component declares the implementations it requires as interfaces which are arguments to its constructor.
> - The component stores a pointer or reference to the interface.
> - The lifetime of the interface must cover the lifetime of the component.
> - The benefits:
>   - The caller can specify which implementation will be used.
>   - Dependencies are explicit and easy to find.
> - The challenge:
>   - Composing objects is complex and … breaks encapsulation.
>   - When constructors change, the wiring up code needs to change.

other languages use runtime reflection to create dependency injection frameworks, but c++ doesn't have runtime reflection yet. IOC containers

> Inversion of Control Container
>
> - IOC containers resolve the challenges of DI.
> - An IOC container is a factory with two aspects:
>   - **Register methods**: Specify which objects should be instantiated for which interface.
>   - **Resolve methods**: Build objects specified by interface based on the specifications.
> - How does this solve the complexities of composition?
>   - The number of specifications is proportional to the number of abstractions – not to the number of types that need to be resolved.

distinguishing between implementation inheritance and interface inheritance.for implementation inheritance, we can use composition, for interfaces we can use runtime polymorphism or CRTP.

#### Lessons Learnt

> 1. Use SOLID principles to guide your design.
> 2. Use the Core CPP Guidelines and tools that help you implement them.
> 3. Organize headers to reduce coupling.
> 4. Use modules for new projects.
> 5. Consider using _std::expected_ for expected errors.
> 6. Use generators for lazy iteration and for inversion of dependencies.
> 7. Avoid Primitive Obsession.
> 8. Use interfaces on component boundaries.
> 9. Inject implementation dependencies.
> 10. Prefer composition over implementation inheritance.
> 11. Consider the use of Inversion of Control Containers to construct objects.

</details>

### Using Modern C++ to Revive an Old Design - Jody Hagins

<details>
<summary>
designing a stream based processing queue to handle coupling, cohesion, and modularity.
</summary>

[Using Modern C++ to Revive an Old Design](https://youtu.be/Taf5eqUZAA0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/cppcon2022_streams.pdf)

"AKA: Coupling and Cohesion are Guiding Lights"

> Pick All of Them
>
> - Low Coupling
> - High Cohesion
> - Composable
> - Testable
> - Reusable
> - Functional
> - Modular
> - Easy to Use
> - Easy to Change
> - High Throughput
> - Low Latency
> - Optimal Code Generation
>
> "One goal of design is to minimize coupling between parts and to maximize cohesion within them."\
> ~ Multi-Paradigm Design for C++ James Coplien

high/loose coupling, cohesion, modularity.,

> Cohesion is how much one part of a code base forms an atomic program unit.\
> Coupling is how much a single program unit depends upon other program units.

program units can be objects (classes), but also can be functions, types, even code blocks.

example of a flow that checks if packet needs decompressing, moving into a modern c++ design.

talking about user space and kernel space. streams, message queues. a story about making a distributed system back in the 90s.

#### Implementation

this very complicated decltype thing... using dependency injection.

```cpp
auto check_timestamp = [](auto & fw, Packet const & pkt) -> decltype(
  add_tag<HasCompressionFlag>(fw, pkt),
  bool{supports_compression(pkt)},
  void())
{
  if (supports_compression(pkt)) {
    put_next(fw, add_tag<HasCompressionFlag>(fw,pkt));
  } else {
    put_next(fw, pkt);
  }
};

auto check_flag = [](auto & fw, auto const & ev)
-> decltype(
  check_tagged<HasCompressionFlag>(fw, ev),
  bool{should_compress(event_for(fw, ev))},
  void())
{
  if (should_compress(event_for(fw, ev))) {
    put_next(fw, add_tag<Compressed>(fw,  remove_tag<HasCompressionFlag>(fw, ev)));
  } else {
    put_next(fw, ev);
  }
};

auto uncompress = [](auto & fw, auto const & ev)
-> decltype(
 check_tagged<Compressed>(fw, ev),
 uncompress(event_for(fw, ev)),
 void())
{
  put_next(fw,
  remove_tag<Compressed>(
    fw,tag_as(fw, ev)
    (uncompress(event_for(fw, ev))))
  );
};

auto process_packet = [](auto & fw, Packet const & pkt)
{
 dependency<ExchangeFooSession>(fw).process_packet(pkt);
};

// the stream
auto strm = StreamHead
 | check_timestamp
 | check_flag
 | uncompress
 | process_packet
 ;
```

we could also make modules that are more complicated.

</details>

### Back to Basics: Object-Oriented Programming in C++ - Amir Kirsh

<details>
<summary>
over view of OOP in C++. syntax and design considerations.
</summary>

[Back to Basics: Object-Oriented Programming in C++](https://youtu.be/_go74QpFPAw)

#### Object Oriented Programming

focus on the data, how is the data represented, what operations are allowed and what context.

class as the code that describes the data, and object as an instantiation of the data.

```cpp
class Point {
  int x, y;
public:
  Point(int x1=0, int y1=0): x(x1), y(y1){}
  void set(int x1, int y1){
    x=x1;
    y=y1;
  }
  void move(int diffX, int diffY) ;
  void print() const { std::cout<< "x = " << x <<", y = " << y <<'\n';}
};
```

access modifiers:

- public
- protected
- private

default privileges between classes and structs, also effects inheritance access modifier.

data members, no default initialization for primitive types. member functions (sometimes called methods), can have access modifiers, be const or not const, don't effect the size of the object. the size is effected by the members, the base class, the vtable pointer and additional padding.

functions that are in the header are more likely to be inlined. the _this_ pointer always exists.

Constructor:

- no constructor - default empty
- constructor delegation
- default parameters
- initializing lists.

when we reach the body of constructor, the members are already constructed, so if we assign data to it, then we done the same work twice.

we some times must use the constructor initializing list

> - Contained object with no default constructor and no initialization on declaration.
> - Contained `const` data member.
> - Contained reference data member.
> - Base class with no default constructor.

constructor delegation and inheritance.

- copy constructor. the default uses member wise copy (which might result in a shallow copy for pointers), which calls the copy constructor of each member. if we try an invalid signature (passing by value) then the compiler stops us.
- assignment operator, not the same as copy constructor. there is also a default assignment operator.
- constructors are used for casting, when we didn't define the constructor as _explicit_, and it works only for "const ref" or "by value", not for "by ref".

mutable data members can be changed even during `const` member functions, good for memoization, something that isn't part of the object's defintion. usually mutex are mutable if they are used in `const` context.

```cpp
class Array{
  int arr[SIZE]{}
  mutable int sum=0;
  mutable bool isSumUpdated = true;
  void calcSum() const;
public:
  Array(){}
  //...
};
```

Destructor, takes no arguments, happens at the end of the lifetime of the object.

Rule of zero - best if we don't need any manual resource management. it's better to keep management classes away from the business logic class.\
Rule of three - if we need a destructor, implement or block the other methods. don't wait.

```cpp
class A {
  ~A(){/*...*/}
  A(const A&) = delete;
  A& operator=(const A&) = delete;
};
```

Rule of five - if we implement or block any of the five operations, we won't get the defaulted move operations, so we should either default or block them.

```cpp
class A {
  ~A(){/*...*/}
  A(A&&) = default;
  A& operator=(A&&) = default;
};
```

Inheritance for code reuse and polymorphism.

calling the base constructor, destructor are also called automatically, but if we are using inheritance for polymorphism, then we need to define the destructor as a virtual method. dispatching is only based on the calling object. runtime dispatching is based on the calling object. the const-ness of the overriding functions must be maintained. using ` = 0;` makes the function pure virtual, meaning the class can't be instantiated.

#### Beyond the "Classic" Model

when and why to not use the inheritance model.

- array of struct and struct of arrays. memory locality, performance.

the square and rectangle example. Liskov substitution principle. composition over inheritance (the state pattern). the strategy pattern to model different behavior (rather than state).\
Using the factory\abstract factory pattern to avoid exposing internal decisions making process.

polymorphism vs Templates: don't use hierarchies just to force functionality (such as ISortable). virtual functions have performance costs in runtime, while templates have compile time costs (can cause code bloat).

</details>

### Breaking Dependencies - C++ Type Erasure - The Implementation Details - Klaus Iglberger

<details>
<summary>
Moving from inheritance to type erasure
</summary>

[Breaking Dependencies - C++ Type Erasure - The Implementation Details](https://youtu.be/qn6OqefuH08), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Type-Erasure-The-Implementation-Details-Klaus-Iglberger-CppCon-2022.pdf)

continuing the ideas from a talk last year, showing how type erasure improves performance.

#### The Motivation for Type Erasure

example of the command pattern, we could use inheritance, but today it's better to use _std::function_, which offers value-semantics, and avoid memory allocation. the example of the "shape" base class hierarchy, where we replace the base class with a type erasure.

#### A Basic Type Erasure Implementation

Basic terminology for C++

> Type Erasure is not
>
> - a void\*;
> - a pointer-to-base;
> - a std::variant.
>
> Type Erasure is
>
> - a **templated constructor** plus
> - a completely non-virtual interface;
> - **External Polymorphism** + **Bridge** + **Prototype**.

the toy problem: shapes. no base class, no connection between the two, and the classes don't know about any operations, they are just data,

```cpp
class Circle {
  public:
  explicit Circle (double rad): radius{rad}, // ... Remaining data members
  {}

  double getRadius() const noexcept;
  // ... getCenter(), getRotation(), ...
  private:
  double radius;
   // ... Remaining data members
};

class Square {
public:
 explicit Square (double s):side {s}, // ... Remaining data members
 {}
 double getSide() const noexcept;
 // ... getCenter(), getRotation(), ...
private:
 double side;
 // ... Remaining data members
};

struct ShapeConcept
{
  virtual ~ShapeConcept() = default;
  virtual void do_serialize(/*...*/) const = 0;
  virtual void do_draw(/*...*/) const = 0;
  //...
};

template<typename ShapeT>
struct ShapeModel : public ShapeConcept
{
  ShapeModel (ShapeT shape): shape_{ std::move(shape)} {}
  // ...
  ShapeT shape_;
  void do_serialize(/*...*/) const override {
    serialize(shape_, /*...*/ );
  }
  void do_draw(/*...*/) const override {
    draw(shape_, /*...*/ );
  }
};

// free functions
void serialize(const Circle &,/*...*/);
void draw(const Circle &,/*...*/);

void serialize(const Square &,/*...*/);
void draw(const Square &,/*...*/);
```

now, any shape to fit into a shapeModel must have a draw and serialize functions that accept it as an argument. the shapeModel and shapeConcept are the **external Polymorphism** Design pattern (which is not part of the GOF design patterns), it extracts the implementation details and removes dependencies, and allows for multiple implementations.

in this implementation, we use a vector of std::unique pointers that take a ShapeModel. but this still has allocation.

we clean this by wrapping the unique_pointer into a class, and this class constructor deduces the type of the shape and stores it. the templated constructor is the the **bridge** design pattern. if we want to be able to copy a shape, we need to add clone function, thereby using the **prototype** design pattern. also some bits about the move operations.

This improves testability - the shape wrapper is a high level abstraction, with the low level individual shapes and the functions are separated. we can add another model that allows us to inject the functions, so we use the **strategy** design pattern and we can test behaviors as we wish.

The performance is still the same as the inheritance model, so we need to optimize this.

#### Type Erasure with Small Buffer Optimization (SBO)

the std::make_unique creates a dynamic allocation via the new allocator. we can replace this with a local buffer, we create an array of bytes and use the global `placement new`, we use reinterpret cast in the `operator*` overloads. we also need to modify the clone function from before and the copy operations, and the move operation also changes, and we need to call the destructor explicitly.

(we could inject the storage policy, but that would be too much for this talk).

#### Type Erasure with Manual Virtual Dispatch

another abstraction, rather than owning a copy of the shape, we have a non-owning abstraction.

```cpp
class Circle { /*...*/ };
class Square { /*...*/ };
void draw( ShapeConstRef shape )
{
 /* Drawing the given shape */
}

int main() {
  Circle circle( 2.3 );
  Square square( 1.2 );

  draw( circle );
  draw( square );
  // ...
}
```

which looks like this:

```cpp
class ShapeConstRef {
public:
 template< typename ShapeT >
  ShapeConstRef(ShapeT const& shape):
   shape_{std::addressof(shape)}, // somebody might have overloaded the `&` operator
   draw_{ []( void const* shape) {draw(*static_cast<ShapeT const*>(shape));}} // stateless lambda
   {}

private:
 friend void draw(ShapeConstRef const& shape) {
  shape.draw_( shape.shape_ ); // hidden friend
 }
 using DrawOperation = void(void const*);
 void const* shape_{nullptr}; // this was created in the constructor
 DrawOperation* draw_{nullptr};
};
```

</details>

### Pragmatic Simplicity - Actionable Guidelines To Tame Cpp Complexity - Vittorio Romeo

<details>
<summary>
guidelines towards making code less complex and easier to reason about.
</summary>

[Pragmatic Simplicity - Actionable Guidelines To Tame Cpp Complexity](https://youtu.be/3eH7JRgLnG8), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/vittorio_romeo_pragmaticsimplicity.pdf).

complexity at high level (system design, software architecture) and low level (abstraction, coding style, language features), this talk will focus on low-level complexity.

> Goals:
>
> - Derive pragmatic and actionable guidelines from various examples
> - When facing a doubt/conflict, solve it using the knowledge gained in this talk
> - Spark some interesting discussion!

equivelent code examples, these snippets

```cpp
int hardcoded_ids[] = {997, 998, 999};
int age = 27;
int main()
{
std::cout << age << '\n';
}

for (int i = 0 ; i < 100; ++i)
{
std::cout << i << ", ";
}
```

are the same as these:

```cpp
std::array hardcoded_ids{997, 998, 999};
auto age = int{27};
auto main() -> int
{
std::cout << age << '\n';
}

for (const int i : irange(0, 100))
{
std::cout << i << ", ";
}
```

simplicity and complexity can be confused with familiarity.

simple code:

- easy to understand, maintain, change, debug, test
- protects us from mistakes at compile-time
- has a limited number of moving parts

but there are always trade-offs. we need a way to derive general principals from specific examples:

> Let’s compromise:
>
> 1. Somewhat agree on what simplicity means by comparing examples
> 2. Derive more general precepts from those examples
> 3. Discuss where such general precepts fall short in the real world

#### Guideline 1 - Limited power

example of choosing between `(int) x` and `static_cast<int>(x)`. they are both the same at primitive values, and the static cast is longer to write and read. but we should choose static_cast. side rant about the proposal for adding `as` (pattern matching) to the language.

`.push_back()` and `.emplace_back()` onto a vector. `std::lock_guard` vs `std::scoped_lock`. all three examples show one operations using a more powerful tool, and one using something more specific. sometimes the more general tool has places to make a mistake (such as scoped lock which can take zero mutex objects). choosing one form over the other also conveys the intent.

> “Local reasoning is the idea that the reader can make sense of the code directly in front of them,
> without going on a journey discovering how the code works.” – Nathan Gitter

| requirement        | specific case         | general case (powerful)              |
| ------------------ | --------------------- | ------------------------------------ |
| casting            | `static_cast<int>(x)` | c-style `(int) x`                    |
| inserting elements | `.push_back()`        | `.emplace_back()`                    |
| locking            | `std::lock_guard`     | `std::scoped_lock`                   |
| arrays             | `std::array<int>`     | c-style `int[]`                      |
| polymorphism       | `std::variant`        | `virtual` inheritance                |
| byte manipulation  | `sts::byte`           | `char`                               |
| enum               | `enum class`          | `enum` (allows implicit conversions) |
| references         | `const auto &`        | `auto&&`                             |

> Use the ~~right~~ _limited_ tool for the job.

not always great, such as choosing between `std::map`, `std::unordered_map` and `absl::hash_map` (is it worth bringing in an external library just to have a more limited map?), or preferring aggregate types over non-aggregates. it's nice having less code and structs without constructors, but we run the risk of having non-initialized fields.

list initialization can invoke `std::initializer_list` constructor.

> Use the most _limited_ tool for the job (<sub>within reason</sub>)

#### Guideline 2 - Signal and Noise Ratio

adding the `[[nodiscard]]` attribute to all class member functions, constantly using `auto` and `noexcept`.

repetition can become noise. some features are technically correct to use everywhere, but using them sparingly calls attention to them when they are used:

- `[[nodiscard]]`
- `auto`, `noexcept`, `final`, `constexpr`
- trailing return types
- `const` variables

if we have them everywhere, we don't notice when they are important. using te features when they have benefits. and not using it if there is a chance someone will have a use case for it.

some things are better to be used liberally. such as `override`.

> Value is a function of rarity(<sub>most of the time</sub>).

if we always have everything everywhere, we stop noticing it. the signal stops being useful and becomes part of the background noise.

it's not always consistent, but consistency isn't the only value. simplicity and correctness are more important than consistency.

#### Conclusion

> How to use these precepts
>
> 1. “Use the most limited tool for the job. ”
> 2. “Value is a function of rarity.> ”
>
> Scenarios:
>
> - Excitement when using a new C++ (or library) feature
> - Resolving conflict during code reviews or debates
> - Migrating a legacy project to more modern standards
> - Preventive damage control for new C++ developers
> - Teaching and mentorship, reducing decision-making surface area
>
> Shortcomings:
>
> - Additional verbosity
> - Loss of style consistency
> - More mental focus required
> - Sometimes subjective

</details>

### The Observer Design Pattern in Cpp - Mike Shah

<details>
<summary>
Overview of the Observer Design Pattern.
</summary>

[The Observer Design Pattern in Cpp](https://youtu.be/4GU2YNsHrwg), [slides](https://mshah.io/conf/22/CPPCON2022_%20SoftwareDesignTrack_TheObserverPattern.pdf), [github](https://github.com/MikeShah/Talks/tree/main/2022_cppcon_observer).

looking at a clip from the AngryBirds video game and thinking about "events" and subsystems which are involved.

one action that effects many subsystems and events. some pseudo code with deep nesting and highly coupled behavior.

> design patterns: 'templates' or 'flexible blueprints' for developing software.
>
> should be:
>
> - flexible
> - maintainable
> - extensible

#### The Observer Pattern Implementations

the observer pattern is about communication between objects. this behavior also comes up in the Model-View-Controller (MVC) architecture.

> "The observer pattern is a software design pattern in which an object, named the subject, maintains a list of its dependents, called observers, and notifies them automatically of any state changes, usually by calling one of their methods." ~ wikipedia

Starting with a basic implementation, multiple observers and one subject. the observers subscribe to the subject (which sometimes is $$called publisher) and get notified when things change.\
The publisher has a `std::forward_list<Observer*>` list of observers and it notifies all of them. when an event happens, the subscribers each respond to it.

A better way to do this is by utilizing interfaces using virtual inheritance: **ISubject**, **IObserver**, the resulting behavior is similar. it's a bit more stable and less vulnerable, and more extendible, as each observers responds based on its' own implementation.

Next we make the observers part of sub systems, and we allow the "subject" to notify only some of the observers. we have types of events, and we store the observers based on those events.

#### Going further

Java really likes this pattern, so does QT library, Godot, Blender3D, Maya3D. we could change the code a bit, ownership, better memory management, type safety, etc...

we still have the problem of scaling, we don't have just one subject, we have multiple things publishing, so we don't want to hang the program until everything has been notified and handled. we can move to an event queue or the actor model.

</details>

### Breaking Dependencies - The Visitor Design Pattern in Cpp - Klaus Iglberger

<details>
<summary>
Using the visitor design pattern to break dependencies and allow for easy addition of operations.
</summary>

[Breaking Dependencies - The Visitor Design Pattern in Cpp](https://youtu.be/PEcy1vYHb8A), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Breaking-Dependencies-The-Visitor-Design-Pattern-Klaus-Iglberger-CppCon-2022.pdf)

starting with the usual example of shapes. a base class with a virtual destructor, and member of type enum which keeps track of the shaped (circle, square). in the first example we use static casting based on the enum type. this type based programming, which makes things hard to maintain and change overtime.

simply changing the enum and adding a value requires re-compiling each of the shapes, and every switch-case statement must handle the new value of the enum. the compiler will help us if we have an switch statement, but not for `if-else` statements.

the next step is using object oriented programming, making the base class a virtual base class and implementations for each function in the interface. this means that every operation that we add to the base class must be virtual. and adding a virtual function forces everyone down the line to change the code to accommodate.

> In dynamic polymorphism you have to make a choice:
>
> - Design for the addition of types
> - Design for the addition of operations

in this talk we will go for the 2nd option, adding operations and not types. if we focus on adding operations, then OOP is probably not the right solution.

#### The Visitor Design Pattern

- Client - > visitor interface -> concrete visitor
- object -> interface - > concrete objects

in our example, we have a shape interface and a shape visitor interface with a `visit` function for each type of shape. then we have concrete visitors (such as "rotate" and "draw") with the implementations of the functions. this makes changing operations non-intrusive,there can be many variants of visitors. the shapes themselves need to have the `accept` operation that calls the `visit` from the provided visitor.\
This design fulfills the single-responsibility principles and the open-closed principle. But it also triggers a massive code change when we add types, as the shape visitor interface must know about the new shape, so all of the specific implementations of the visitor will also have to be updated.\
We simply moved the problem from "adding operations changes all shapes" to "adding shapes changes all visitors". this also means that the visitors can only use the public interface of the shapes, and we always have two virtual function calls, which effect performance.

#### A More Modern Solution

we can use <cpp>std::variant</cpp> as an abstraction. it replaces the base class abstraction, and we refactor the code accordingly. our concrete shapes no longer depend on a visitor interface or a base class, and the visitors don't have dependencies either.

```cpp
class Draw{
  public:
    void operator()(Circle const&) const;
    void operator()(Square const&) const;
};

using Shape = std::variant<Circle, Square>;

void drawAllShapes(std::vector<Shape> const& shapes) {
  for (auto const& s:shapes) {
    std::visit(Draw{}, s);
  }
}
```

using the call operator allows us to pass lambdas as needed. and using value types and not pointers (inside the variant) makes things have better local reasoning and we don't care about memory allocation.\
with this change, we don't have any dependencies cycles, and we can add shapes as we wish and a shapes variants for them if we need. we are never forced to change code because someone else made an unrelated change.

we get a performance increase, it's as fast as the enum approach, but protects us against making mistakes.

| ---         | Classic Visitor                                        | Modern Visitor                                       |
| ----------- | ------------------------------------------------------ | ---------------------------------------------------- |
| Intrusive   | intrusive (base class)                                 | Non-intrusive ( ca) (can be added on-the-fly)        |
| Semantics   | Reference-semantics (based on references/pointers)     | Value-semantics (based on values)                    |
| Style       | OOP                                                    | Procedural                                           |
| Performance | Slow (many virtual functions, scattered memory access) | Fast(no virtual functions, contiguous memory access) |

some potential disadvantages of <cpp>std::variant</cpp>

- not great when the classes have different sizes
- it can reveal dependencies, the variant must know the full defintion of the objects it uses.

#### The Acyclic Visitor Design Pattern

we have an empty "abstract visitor", and we also have specific shapes visitors interfaces. the concrete visitor always inherit from the abstract visitor, and from any specific shape interfaces they care about. this splits the implementation effort, as concrete visitors aren't required to immediately support all new shapes. but it does introduce a dynamic cast into the design inside the `accept` function, and it has a worse performance, which is worse in magnitudes than the other options. this is because of the dynamic cast which is a cross-casting operation (between things which aren't at the same chain).

#### Summary

> - The Visitor design pattern is the right choice if you want to add operations.
> - The Visitor design pattern is the wrong choice if you want to add types.
> - Prefer the value-semantics based implementation based on <cpp>std::variant</cpp>.
> - Beware the performance of Acyclic Visitors.

</details>

##

[Main](README.md)
