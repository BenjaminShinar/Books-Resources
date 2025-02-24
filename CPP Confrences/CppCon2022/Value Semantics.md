<!--
ignore these words in spell check for this file
// cSpell:ignore Simula Hearn offsetoff interconvertibility
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Value Semantics

### Back to Basics: Cpp Value Semantics - Klaus Iglberger

<details>
<summary>
Moving to value semantics over reference semantics
</summary>

[Back to Basics: Cpp Value Semantics](https://youtu.be/G9MxNwUoSt0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Back-to-Basics-Value-Semantics-Klaus-Iglberger-CppCon-2022.pdf)

we start with an overview, from the late 70's and the Simula language, which was one of the first languages to introduces classes. then we look at the Gang of Four design pattern book, from the 23 patterns, most are based on inheritance.

basic, classic, visitor design pattern implementation

a shape visitor is a base class, a shape is a base class, all these base classes and virtual functions require us to allocations, and use the objects via pointers.

```cpp
class Circle;
class Square;
class ShapeVisitor
{
  public:
  virtual ~ShapeVisitor() = default;
  virtual void visit(Circle const&) const = 0;
  virtual void visit(Square const&) const = 0;
};

class Shape
{
public:
  Shape() = default;
  virtual ~Shape() = default;

  virtual void accept(ShapeVisitor const&) = 0;
};

class Circle : public Shape
{
  public:
  explicit Circle(double rad)
  : radius{rad}
  , // ... Remaining data members
  {}
  double getRadius() const noexcept;
  // ... getCenter(), getRotation(), ...
  void accept(ShapeVisitor const&) override;
  // ...
  private:
  double radius;
  // ... Remaining data members
};

class Draw : public ShapeVisitor
{
 public:
 void visit(Circle const&) const override;
 void visit(Square const&) const override;
};

void drawAllShapes(std::vector<std::unique_ptr<Shape>> const& shapes)
{
  for(auto const& s : shapes)
  {
    s->accept(Draw{})
  }
}

int main()
{
  using Shapes = std::vector<std::unique_ptr<Shape>>;
  // Creating some shapes
  Shapes shapes;
  shapes.emplace_back(std::make_unique<Circle>(2.0));
  shapes.emplace_back(std::make_unique<Square>(1.5));
  shapes.emplace_back(std::make_unique<Circle>(4.2));
  // Drawing all shapes
  drawAllShapes(shapes);
}
```

> This style of programming has many disadvantages:
>
> - We have a two inheritance hierarchies (intrusive)
> - Performance is reduced due to two virtual function calls per operation
> - Performance is affected due to many pointers (indirections)
> - Promotes dynamic memory allocation
> - Performance is reduced due to many small, manual allocations
> - We need to manage lifetimes explicitly (std::unique_ptr)
> - Danger of lifetime-related bug

the better solution is to use value semantics. in this case, std::variant. if we use this, we can start writing the code again, this time without base classes, virtual functions, and without knowledge about external functions. our objects are values, not pointers. an

```cpp
class Circle {
  public:
  explicit Circle(double rad): radius{rad}, // ... Remaining data members
  {}
  double getRadius() const noexcept;
  // ... getCenter(), getRotation(), ...
  private:
  double radius;
  // ... Remaining data members
};

class Square {
  public:
  explicit Square(double s): side{s}, // ... Remaining data members
  {}
  double getSide() const noexcept;
  // ... getCenter(), getRotation(), ...
  private:
  double side;
  // ... Remaining data members
};

class Draw {
  public:
  void operator()(Circle const&) const;
  void operator()(Square const&) const;
};

using Shape = std::variant<Circle,Square>;

void drawAllShapes( std::vector<Shape> const& shapes )
{
  for(auto const& s : shapes)
  {
    std::visit(Draw{}, s);
  }
}

int main()
{
  using Shapes = std::vector<Shape>;
  // Creating some shapes
  Shapes shapes;
  shapes.emplace_back(Circle{2.0});
  shapes.emplace_back(Square{1.5});
  shapes.emplace_back(Circle{4.2});
  // Drawing all shapes
  drawAllShapes(shapes);
}
```

> This style of programming has many advantages:
>
> - There is no inheritance hierarchy
> - The code is so much simpler (KISS)
> - There are no virtual functions
> - There are no pointers or indirections
> - There is no manual dynamic memory allocation
> - There is no need to manage lifetime
> - There is no lifetime-related issue (no need for smart pointers)
> - The performance is better

some performance benchmarking results

another example, std::span which uses reference semantics, and somehow allows us to go over const stuff, and allows for undefined behavior if we reallocate the original data.

```cpp
#include <vector>
#include <span>
void print(std::vector<int> const& vec);
int main()
{
 std::vector<int> v{1, 2, 3, 4};
 std::vector<int> const w{v};
 std::span<int> const s{v};
 w[2] = 99; // Compilation error!
 s[2] = 99; // Works!
 // Prints 1 2 99 4
 print(v);
 return EXIT_SUCCESS;
}
```

std::span is a dangerous thing to have around, it should be used as a function argument. another example is what happens when we use iterators and `std::erase`, which use references internally.

There are further examples for value semantics from the Standard Library:

- The design of the STL (C++98) - containers are values, copy is deep copy
- std::optional (C++17) - avoid throwing exceptions for common cases of 'failure'
- std::expected (C++23) - a more advanced case std::optional
- std::function (C++11) - the command pattern, a callable object
- std::any (C++17)

</details>

### Back to Basics: Master C++ Value Categories With Standard Tools - Inbal Levi

<details>
<summary>
Overview of value categories (lvalue, rvalue, etc...).
</summary>

[Back to Basics: Master C++ Value Categories With Standard Tools](https://youtu.be/tH0Z2OvHAd8), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Back-To-Basics-Value-Categories-Inbal-Levi-CppCon-2022.pptx)

#### Motivation and Introduction

in this example, we ask how many times will this object be created.

```cpp
struct Data {
   Data(size_t s);	 	// Constructor
   Data(const Data&);	// Copy Constructor
   Data(Data&&);	   	  	// Move Constructor

   size_t s;
   int* b;
};

const Data getData(size_t s) {
   return Data(s);
}

auto d1 = Data(42);
auto d2 = std::move(d1);

auto d3 = getData(42);
auto d4 = std::move(getData(42)); // 2 objects
```

the final line has two object, because the function returns a `const` object, so the copy constructor is called rather than a move constructor.

value categories began in C, what was once called 'lvalue' and 'rvalue' expressions. it defines:

- lifetime
- is it temporary
- is it observable after changes
- does it have an address

value categories affects performance and overload resolution.

```cpp
Data a = 42;
Data& lval_ref_a = &a; // lvalue ref
//Data&& rval_ref_a = &a; // rvalue ref (FAILS!)
Data&& rval_ref_a = 42; // rvalue ref (OK)
```

> "Value category is a quality of **an expression**"

```cpp
struct Data {
   Data(int x);
   int x_;
};
void foo(Data&& x) {
   x = 42;
}

Data&& a = 42;
foo(a);		   		// Fail! lvalue!
foo(Data(73));	   	// OK
```

in the above example, we see the confusion between the type (rvalue reference to Data) and the value category (lvalue).

the same object can have different value category depending on the context.

```plantuml
@startuml
digraph g {
"node0" [label = "expression"];
"node1" [label = "glvalue"];
"node2" [label = "rvalue"];
"node3" [label = "lvalue"];
"node4" [label = "xvalue"];
"node5" [label = "prvalue"];
"node0"-> "node1";
"node0"-> "node2";
"node1"-> "node3";
"node1"-> "node4";
"node2"-> "node4";
"node2"-> "node5";
}
@enduml
```

value categories evolved, starting with C: lvalue, non-lvalue expression, and function designators. in C++98, expressions could be lvalue (objects, functions, references) or rvalue (non-lvalue, can be bound by `const` lvalue reference). C++11 added rvalue references and move semantics.

| -                          | has identity (glvalue) | doesn't have identity |
| -------------------------- | ---------------------- | --------------------- |
| Can't be moved from        | lvalue                 | -                     |
| Can be moved from (rvalue) | xvalue                 | prvalue               |

c++17 added _guaranteed copy elision_, defining prvalue materialization
-| has identity (glvalue) | doesn't have identity
---|---|---
Can't be moved from | lvalue | -
Can be moved from (rvalue) | xvalue | prvalue materialization

c++20 added implicit move from rvalue references in return statements,C++23 will added `deducing this`, `like_t` and `forward_like`.

glvalue - has identity, rvalue - can be moved from.

> - glvalue - expression whose evaluation determines the identity of an object or function.
> - xvalue - glvalue that denotes an object whose resources can be reused (usually because it is near the end of its lifetime).
> - lvalue - glvalue that is not an xvalue.
> - prvalue - expression whose evaluation initializes an object, or computes the value of the operand of an operator, as specified by the context in which it appears, or an expression that type _cv void_.

string literals are lvalue (have address, can't be moved from). but literals (like integers) are prvalue (can be moved from, don't have an address).

#### Value Categories in Practice

expression with different value categories "bind" to different types of references.

- initialization or assignment
- function call
- return statements

| ---                        | binds lvalue? | binds rvalue? | function can modify data? | function can observe (old) data? |
| -------------------------- | ------------- | ------------- | ------------------------- | -------------------------------- |
| **lvalue** reference       | V             | X             | V                         | V                                |
| const **lvalue** reference | V             | V             | X                         | V                                |
| **rvalue** reference       | X             | V             | V                         | X                                |
| const **rvalue** reference | X             | V             | X                         | X                                |

the lifetime of an objects can extended by binding to references.

since c++17, there is guaranteed copy elision,

```cpp
Data d = Data(Data(42)); // one constructor call
Data = getData(42); // avoid copy constructor. uses move constructor instead.
```

a prvalue of type T can be converted to an xvalue of type T.

#### Value Categories in Generic code

reference collision: when there are multiple `&` symbols.

```cpp
typedef int&  lr;
typedef int&& rr;

int a;
lr& b = a; // int&& -> int&
lr&& c = a; // int&&& -> int&
rr& d = a; // int&&& -> int&
rr&& e = 73; // int&&&& -> int&&
```

forwarding references (universal references), inside a function template, the "rvalue reference" has a special meaning - it's a forwarding reference.

```cpp
Template <typename T>
void foo(T&& t) { // forwarding reference
   // Type of T here
}

int a = 42;
const int& cla = a;
int&& b = 73;
foo(a); // int &&, T = int&
foo(cla);	// const int &&, T = const int &
foo(std::move(a)); // int &&, T = int
```

T && keeps the value category of the type the instantiation is based on.

#### Tools

utility functions that operate on value categories:

- `std::move` - produces a _xvalue_ expression `T&&`. a casting operator. may not always do what we hope it does. same as doing:\
  `static_cast<typename std::remove_reference<T>::type&&>(t)`.
- `std::forward` - preserves the value category of the object passed to the template. used `std::remove_reference<T>` to get the value type
- `std::decay` - type trait, converts array to pointer, function to function pointer, lvalue to rvalue (removes cv qualifiers, references). the `auto` keyword performs decay.
- decltype specifier - returns the the type of the expression, while persevering the value category. can be used instead of the type
  - if expression is **xvalue**, yields `T&&`
  - if expression is **lvalue**, yields `T&`
  - if expression is **prvalue**, yields `T`
  - special behavior for `decltype(auto)`
  - another special expression for double parentheses `decltype(a)` can be different from `decltype((a))`.
- `std::declval` - (deceleration evaluation)
  - produces xvalue expression T&&
  - if T is void, returns T
  - can return non-constructable or incomplete type.
  - combined with `decltype`
- Deducing this (c++23) - specifying the value category of the object as part of the member function.`template <typename T> void Foo(this T&& t){}
`. we can then combine it with forwarding references

```cpp
// old way
struct Type{
  auto Foo() const &;
  auto Foo() &;
  auto Foo() &&;
};

// new way
struct Type{
  auto Foo(this const Type &):
  auto Foo(this Type &):
  auto Foo(this Type &&):
};

// best way
struct Type{
  template <typename Self>
  auto Foo(this Self && self):
};
```

#### Summary

> - Qualities:
>
>   - Value Categories are a quality of an expression
>   - What (misleadingly) looks like the value category, can in fact be the type
>   - What (misleadingly) looks like the same entity, is, in fact depend on context
>
> - Binding rules apply in the following “events”:
>
>   - Initialization or assignment
>   - Function call (including non-static class member function called on an object)
>   - Return statement
>
> - The behavior of an entity is defined by the binding object
>   - Initialization: limits are according to the reference which binds it
>   - Function call: limits inside the function are according to the overload which binds it
>   - Return statement: limits as in initialization, with additional rules due to optimizations and const

</details>

### Val: A Safe Language to Interoperate with C++ - Dimitri Racordon

<details>
<summary>
The design of a safe, fast, and simple programming language.
</summary>

[Val: A Safe Language to Interoperate with C++](https://youtu.be/ws-Z8xKbP4w), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Val-at-CppCon-2022.pdf), [Val-lang](https://www.val-lang.dev/).

> Val Wants To Be Your Friend:
> The design of a safe, fast, and simple programming language.

```val
func main(){
  print("Hello, World!")
}
```

Mutable Value Semantics (MVS)

> "To understand how a program works, it should
> be possible for reasoning and specification to
> be confined to the cells that the program
> actually accesses. The value of any other cell will
> automatically remain unchanged."\
> ~ Peter O'Hearn"

contrast between reference to value semantics, value semantics help with local reasoning. the c++ code will print "3", while the python code prints "4".

```cpp
int main() {
  std::vector v1 = { 1, 2, 3 };
  std::vector v2 = v1;
  v2[0] += 10;
  v2.push_back(4);
  std::cout << v1.size() << std::endl;
}
```

```py
def main():
  v1 = [1, 2, 3]
  v2 = v1
  v2[0] += 10
  v2.append(4)
  print(len(v1))
```

- Explicit Copies
- Passing conventions
  - `let` - default, const reference
  - `inout` - mutable reference
  - `sink` - relinquish ownership (move and destroy)
  - `set` - out parameter, must be initialized when function exits
- Method bundles - function overload sets
- Local bindings
  - `let` - immutable access
  - `inout` - mutable access
  - `var` - consuming the value and removing it from existence
- Projections - computed accessor properties without in memory store
  - `property` - accessor (mutable)
  - `subscript` - `[]` operator
- Unsafe operations - marking operations as not being protected by the compiler
- Generic programming
  - `traits` - like <cpp>concepts</cpp>, can have default implementations
  - `conformance` - make a type behave like it has an interface.
  - `extension` - add methods on traits, or act as an customization point
- Concurrency
  - `spawn` - run code asynchronously

#### Val in a NutShell

> Can a language enforce the guarantees of MVS as practiced in C++, without loss of efficiency?
>
> - Fast by definition
> - Safe by default
> - Simple
> - Interoperable with C++

tha language was influenced by the swift programming language. it tries to enforce c++ best practices through the compiler.

all copies in Val must be written explicitly, even copies of primitives.

```val
fun f(x: Int) -> Int {
  var y = x.copy()
  print(x)
  y += 1
  return y
}
```

the default behavior of passing parameters to function is by value, in c++ we get better performance by passing a reference and marking it `const`. val achieves the same result in the compiler. it knows to take references on it's own, and warns us against unnecessary copies. it also works for return value, if the value is no longer used, then there's no need to explicitly copy it.

```val
fun print_all(_ things: Array<String>) {
  print(things.joined(by: ", "))
}
fun main() {
  var fruits = Array(["durian", "mango", "apple"])
  print_all(fruits.copy())
}

fun unused_space(_ things: Array<String>) -> Int {
  let space = things.capacity() - things.count()
  return space.copy()
}
```

Passing Conventions are at the core of Val. this type conforms to the "copyable" concept, so it has `.copy()`

```val
type Vec2: Copyable {
  var x: Double
  var y: Double
}

fun offset (_ v: let Vec2, by d: Vec2) -> Vec2 {
  Vec2(x: v.x + d.x, y: v.y + d.y)
}
```

`let` is like c++ `const &`, immutable, not changing the parameter. we don't write the explicit `return`.

```val
fun offset_inout (_ v: inout Vec2, by d: Vec2) {
  &v.x += d.x
  &v.y += d.y
}
```

the `&` isn't taking the address, it means that we do an assignment. `inout` modifies the parameter, like a reference, we can mark if this is the only usage of the parameter or if it's shared in the same function (slicing), which allows us to avoid the hidden bug in this code:

```cpp
void offset_inout(Vec2& v, Vec2 const& d) {
  v.x += d.x;
  v.y += d.y;
}

// Offsets `v` by `2 * d`.
void double_offset_inout(Vec2& v, Vec2 const& d) {
  offset_inout(v, d);
  offset_inout(v, d);
}

void main() {
  Vec2 vec = {3, 4};
  double_offset_inout(vec, vec);
  std::cout << vec.x << std::endl; // Should print 9, but prints 12 instead.
}
```

since we passed the same argument as both parameters, the "d" parameter changes inside the function, even though we marked it as const reference. in Val we would see an error and the compiler would suggest we create a copy.

```val
fun offset_sink (_ v: sink Vec2, by d: Vec2) -> Vec2 {
  Vec2(x: v.x + d.x, y: v.y + d.y)
}
fun main() {
  defer { print("will exit main") }
  var vec = Vec2(x: 3, y: 4)
  vec = offset_sink(vec.copy(), by: vec)
  print(vec.x)
}
```

`sink` parameters 'consume' the argument, they take and destroy it, so it can't be accessed afterwards. it's like <cpp>std::move</cpp> that always destroys the object. there is no guarantee when destructors are called, so RAII-lie behavior is implemented with `defer` blocks that run on the end of the scope.

`inout` and `sink` are very similar and can achieve te same result. example of having three functions that do essentially the same thing, but with different memory models. we can around this by employing **method bundles**. we define one function and have three implementations, which the compilers choose between based on the usage.

```val
type Polygon: Copyable {
  var vertices: Array<Vec2>
  fun offset(by d: Vec2) -> Polygon {
  let { ... }
  inout {.... }
  sink { ... }
 }
}

fun main() {
  var shape = Polygon(vertices: ...)
  &shape.offset(by: Vec2.unit_x)
  print(shape)
}
```

usually, the compiler can synthesize the `inout` and `sink` version based on the `let`. a last passing convention is `set`, which is rarely used, this is an `out` parameter.

**Local bindings**

```val
fun main() {
  var velocity = Vec2(x: 3, y: 4)
  let x = velocity.x
  print(x)
  &velocity.x += 1
  print(velocity)
}
```

if we switch the order of the print and assignment statements, we get an error, because x was bounded and so we can't mutate it. we need to explicitly copy it. the "lifetime" of an object starts with the declaration and ends with it's last use. this protects us from changes and mutation that other "holders" of this value can cause (much like rust "borrow" mechanics).

Projections are kind of properties without in-memory representations.

```val
type Angle {
  var radians: Double
}

fun main() {
  var theta = Angle(radians: Double.pi)
  inout x = &theta.radians
  &x += Double.pi * 0.5
  print(theta)
}
```

we can use projections to have specific accessors. the `yield` command gives the control to the caller for the entire useful lifetime of it, and then it is returned the callee.

```val
type Angle {
  var radians: Double
  property degrees: Double {
    inout {
      var d = radians * 180.0 / Double.pi
      yield &d
      radians = d * Double.pi / 180.0
    }
  }
}
fun main() {
  var theta = Angle(radians: Double.pi)
  inout x = &theta.degrees
  &x += 90
  print(theta)
}
```

an example of 2x2 matrix, we use `subscript` to create a row representation from a column major matrix. we again use the method bundles to define different behavior based on the usage (`let` for reading, `inout` for reading and changing, `set` for writing only, and `sink` to extract and destroy the matrix).

some operation are **unsafe operation** by design, such as manipulations raw memory.\
in this code, we have two different behavior based on the number of bytes. either storing in place or allocating memory. when we operate on raw memory, we mark our operations with `unsafe`. we als have `precondition` for out of bound index check.

```val
public type Bytes {
  var rep: {count: Int, contents: Int8[7] | MutablePointer<Int8>}

  public init(bytes: Array<Int8>) {
    if bytes.count() <= 7 {
      rep = (count: bytes.count(), contents: Int8[7](contents_of: bytes, filling_with: 0))
    } else {
      let buffer = MutablePointer.allocate(count: bytes.count())
      for n in 0 ..< bytes.count() {
        unsafe buffer.advanced(by: n).initialize(to: bytes[n].copy())
      }
    rep = (count: bytes.count(), contents: buffer)
    }
  }

  public subscript(_ i: Int): Int8 {
    inout {
      precondition(i >= 0 && i < rep.count, "index out of bounds")
      match rep.contents {
        let buffer: Int8[7] { yield &buffer[i] }
        let buffer: MutablePointer<Int8> { yield unsafe &buffer[i] }
      }
    }
  }

  public deinit {
    if let buffer: MutablePointer<Int8> = rep.contents { buffer.deallocate() }
  }
}
```

#### Generic Programming

> The basic of generic code:
>
> - Concepts are interfaces with associated types and values
> - Generics are type-checked separately
> - Generics use static dispatch
> - But dynamic dispatch is useful too

val has type checking, stronger than C++ templates which have massive error messages, and even stronger than using concepts. concepts in val are called `traits`. we can force our types to behave as if they have a trait by defining a `conformance` (like the _adaptor_ design pattern), and we can add methods to an existing trait with `extension`, or use it to customize the behavior of types.\
an example of type erasure using shapes. the `any` keyword allows a container to hold elements which belong to types with certain traits.

#### Concurrency

`spawn` to evaluate arbitrary code asynchronously, which are later joined. we still have warnings on overlapping access, so it gives us some thread safety.

#### Memory Safety

Val is memory safe by default since construction, while Carbon adds safety post-hoc by improving on c++.

</details>

### A Tour of C++ Recognised User Type Categories - Nina Ranns

<details>
<summary>
Some definitions from the C++ Standard and how types are categorized.
</summary>

[A Tour of C++ Recognised User Type Categories](https://youtu.be/pdoUnvTwnr4), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CPPCON-2022-Tour-of-User-defined-types-in-C.pdf).

part of the team that wrote "embracing modern C++ safely" book.

types with special provisions,

> - treating user defined types as built in types in certain cases
> - new features require new definitions (constexpr specifier and literal types)
> - "blessing" certain common programming patterns inherited from C code.

types: scalar and user defined

- scalar types
  - arithmetic
  - enumerations
  - pointer
  - pointer to member
  - <cpp>std::nullptr_t</cpp>
- class types
- array types

types with special provisions beyond the usual:

> - aggregate types
> - standard-layout types
> - trivially copyable types
> - trivial types
> - POD types
> - literal types
> - structural types
> - implicit-lifetime types

#### Aggregate Types

since C++98, can be initialized as a collection of objects (aggregate initialisation), and since c++20 have designated initialization with the dot notation. over time, more and more types are considered aggregates.

library trait <cpp>std::is_aggregate</cpp> since C++17, which might be obsolete because C+20 added parens initialisation.

#### Standard-Layout Types

classes with "predictable" layout, exist for communicating with code from other languages. same access control for all non-static data members, and all data members are in one the classes. no base classes or sub objects with with the same address. c++20 has the `[[no_unique_address]]` attribute. no virtual functions, no reference members, all member are standard layout.

> Implications
>
> - layout of the object depends only on the non-static data members
> - the address of a standard layout class object and its first non-static data member is the same
> - the address of a standard layout class object and all of its base classes is the same.
> - <cpp>offsetof</cpp> is well defined with standard layout class types.

common inital sequence and unions.

library trait from c++11 <cpp>std::is_standard_layout</cpp>, which isn't widely used.

#### Trivially Copyable Types

trivial special member functions, trivial constructor and destructor, all member are also trivially copyable. has at least one non-deleted copy operation,

> Implications of Trivially Copyable
>
> - the value is contained in the underlying byte representation
> - can be memcpy-ed to and from another object of same type
> - can be memcpy-ed to and from an array of char, unsigned char,or <cpp>std::byte</cpp>

library trait from c++11 <cpp>std::is_trivially_copyable</cpp>, a prerequisite for bitwise copy.

#### Trivial Types

> - trivial class type = trivially copyable + non-deleted/eligible trivial default constructor
> - remnant from POD days - only used in the library to describe types which used to be required to be of POD type

library trait from c++11 <cpp>std::is_trivial</cpp>, also limited use.

#### POD Types

plain old data, a term defined in c++98 for C compatiblity, used to mean an object that behaves like a C type in terms of construction, destruction and copying. term is no longer relevant, mostly replaced with _trivial standard layout_.

library trait from c++11 <cpp>std::is_pod</cpp>, was deprecated in C++20.

#### Literal Types

> - A literal type is one for which it might be possible to create an object within a constant expression.
> - It is not a guarantee that it is possible to create such an object, nor is it a guarantee that any object of that type will usable in a constant expression.

new keywords:

- <cpp>constexpr</cpp> (c++11) - can be used in both compile and runtime.
- <cpp>consteval</cpp> (c++20) - functions only be called in compile time.
- <cpp>constinit</cpp> (c++20) - variables that only exists in compile time.

library trait from c++11 <cpp>std::is_literal_type</cpp>, deprecated in c++17, and removed in c++20.

#### Structural Types

> - prerequisite for non-type template parameter
> - scalar type
> - lvalue reference type
> - literal class type with all base classes and non static data members public, non mutable and of structural type

#### Implicit-Lifetime Types

The C++ object model didn't define some types (usually those who come from C), so the implicit lifetime defintion was created.

> - scalar types
> - array types
> - aggregates classes and classes with a trivial destructor and at least one trivial constructor

#### Summary

| type                   | when do you care                                                                               | trait                                                                            |
| ---------------------- | ---------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| aggregate type         | member-wise initialisation, ability to use designated initializers                             | <cpp>std::is_aggregate</cpp> (C++17)                                             |
| trivially copyable     | type copying using <cpp>memcpy</cpp>, <cpp>std::memmove</cpp>, and bit_cast                    | <cpp>std::is_trivially_copyable</cpp> (C++11)                                    |
| trivial type           | partial check for C code compatibility                                                         | <cpp>std::is_trivial</cpp> (C+11)                                                |
| standard layout type   | pointer-interconvertibility, union access through common initial sequence, <cpp>offsetof</cpp> | <cpp>std::is_standard_layout</cpp> (C++11)                                       |
| POD type               | check for C code compatibility                                                                 | <cpp>std::is_pod</cpp> (C++11, deprecated in C++20)                              |
| literal type           | requirement for compile time initialisation                                                    | <cpp>std::is_literal_type </cpp> (C++11, deprecated in C+ +17, removed in C++20) |
| structural type        | requirement for non-type template parameter                                                    | no trait                                                                         |
| implicit-lifetime type | makes common programming patterns well defined                                                 | no trait                                                                         |

</details>

### Value Semantics: Safety, Independence, Projection, & Future of Programming - Dave Abrahams

<details>
<summary>
Reference semantics are bad, value semantics are the future.
</summary>

[Value Semantics: Safety, Independence, Projection, & Future of Programming](https://youtu.be/QthAU-t3PQ4), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CPPCon-Sept-2022.pdf).

Reference semantics are problematic, just like having a global variable. a data can be changed by another user, so we sometimes must employ defensive copying to avoid being affected by something else.

design by contract:

- pre-conditions
- post-conditions
- invariants

> problems with references semantics:
>
> - technical debt
> - spooky action
> - incidental algorithms
> - visibly broken invariants
> - race conditions
> - surprise mutation
> - un-specifiable mutation

```cpp
// Offset by delta
template <class Numeric>
void offset(Numeric& target, Numeric& delta){
  target += delta;
}

// Offset by delta
template <class Numeric>
void offset2(Numeric& target, Numeric& delta){
  offset(target, delta);
  offset(target, delta);
}

void main() {
  auto x = 3;
  offset2(x,x);
  std::cout<< x << std::endl;
}
```

in the above example, the same object is both the target and the offset, so it's mutated internally in the function.

#### Safety

> A safe operation cannot cause undefined behavior.\
> A safe language has only safe operations.

lifetime safety solutions:

- dynamic allocation + garbage collection and reference counting
- borrow checker/named lifetimes (RUST)
- adding static analysis

Thread safety solutions:

- define the Undefined behavior and ignore the real problem,
- borrow checker/named lifetimes (RUST)
- detect and trap

we can try using immutability to avoid it, but giving it up costs performance.

reference semantics mess up our ability for local reasoning about the code in front of us.

#### Value Semantics

value semantics combines _regularity_ and _independence_. "behave like the ints"

- regularity: object is equal to itself, objects equality works both ways, copying an object makes it equal to the original.
- independence - only one access point, no side effects.

we never document the independence pre-condition, it's already part of our mental model. but it's not part of the standard, and we can't bake it into the compiler.

whole-part relationship:

- copying
- equality
- hashing
- comparison
- assignment
- serialization
- differentiation

"access to the parts is always done through the whole".

terms such as "deep" and "shallow" (copying) are red flags that tell us the boundaries of what we are working with are not well defined.

getting value semantics with move semantics for updates, or fixing the language to have "in-out" and "in" parameters, which require exclusivity.

#### Achieving Value Semantics Today

identify the "whole-part" relationship, decoupling a graph-object.

</details>

##

[Main](README.md)
