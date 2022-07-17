<!--
ignore these words in spell check for this file
// cSpell:ignore adjacens plecto maat jfif fourcc gemi blarg bippity
-->

[Main](README.md)

Lightning Talks

## Complecting Made Easy - Tony Van Eerd

<details>
<summary>
'simple made easy' is hard.

</summary>

[Complecting Made Easy](https://youtu.be/jSIMzirLJKE)

the word "easy" comes from the latin "adjacens", which means "close at hand", things that are easy are "close to us".

> complect
>
> - "plecto" (latin) = fold/ braid
> - simple - one fold
> - complex - many folds
> - Complecting - interweaving

complecting is easy, make code intertwined:(bad code!)

- add global variable
  - one could change the global variables
- shared pointer (raw pointer as well)
  - "a shared pointer is as good as a global variable"
- a class that has access to both of them. (MODEL)
- base classes
  - static variable in the base class!
- grab code that you know and copy it!
- members functions are easier than free function,
  - "classes are made of Velcro"
  - naming files is hard... (utils is a bad name)

</details>

## Dashboards to the Rescue - Matthias Bilger

<details>
<summary>
Can Dashboards solve our problems? probably not
</summary>


[Dashboards to the Rescue](https://youtu.be/eOMqO0OKsCw)

everybody loves cats and dashboards. nice looking, no knowledge required, easy access to KPI (key performance index)

if we remove the assert statements, we can make the test pass! we can also count the assert statements, or add pointless tests to ramp up the KPI of line coverage or code coverage.

> - "easy measurable does not mean useful"
> - making a metric a target is bad
> - dashboards will not save us.

</details>

## Universal Function Call Syntax in C++20 - Devon Richards
<details>
<summary>
New syntax for calling functions.
</summary>

[Universal Function Call Syntax in C++20](https://youtu.be/uT1ZJHM8DkE)

calling a free function and a member function with the same syntax, like `std::size` on containers, which uses the member function `size` from the container.

but it also lets us write function from left to right, like the pipe operator for ranges. he has a macro that allows it in compiler explorer, and the performance hit is really small, the macro does a hack to allow templates somehow.
</details>

## When Paradigms Collide - Ben Deane

<details>
<summary>
The Pains of mixing Generic and object oriented technicques
</summary>

[When Paradigms Collide](https://youtu.be/VoIesQvx4_4)

we have a class, and we write an output stream operator for them, it works fine. then we have a vector of them, but if we want to use the operator as part of the algorithm, things get ugly and we end up using the _initialize-then-modify_ anti-pattern.

```cpp
std::stringsteam ss;
for_each(cbegin(array_of_s), cend(array_of_s),
  [&](const auto & elem){
    ss << elem;
  });
std::cout << ss.str() <<'\n';
```

> - I created a stringstream
> - Then I modify it.
> - I can't say it's const.
> - In the general case, I don't get RVO.

we can try using the accumulate algorithm and being more declarative, but this doesn't work. because the return type of `<<` is output stream, no stringstream.

```cpp
const auto ss = accumlate(
  cbegin(array_of_s),
  cend(array_of_s),
  std::stringstream{}, // initialize
  [&](std::stringstream ss, const auto & elem){
    return std::move(ss << elem); // this fails!
  });
```

there is a conflict between paradigms, streams use the OOP, while algorithms use generic approach.

> - a _std::stringstream_ is-a ostream
> - operator `<<` takes and returns reference-to-base (ostream).
> - algorithms work on values
> - algorithms preseve types ass they move thorough the calculation.
>
> You can't preserve the type you're given if the function returns the base!

we can try playing around with the operator, and then we try to constrain it with concepts

```cpp
template <typename T>
concept outputtable = requires (std::ostream& s, T t)
{
  { s << t} -> std::same_as<std::ostream&>;
};

template <typename TDerived, typename TBase>
concept subclass_of =
  std::derived_from<TDerived, TBase>  // remember, a class counts as being derived from it self
  and not std::same_as<TDerived, TBase>;

// wrapper
auto & operator<<(subclass_of<std::ostream> auto & os, const outputtable auto &t)
{
  static_cast<std::ostream&>(os) << t;
  return t;
}
```

> Conclusion
>
> - Mixing Paradigms is sometimes painful.
> - we're stuck with the stream interface.
> - careful use of concepts can help bridge the gap.
> - `std::derived_from` is misnamed.
> - when dealing with class hierarchies, reflexive concepts may not be best.

</details>

## A Few Gentle Rules (\*) but One in Particular - Chris Uzdavinis

<details>
<summary>
Gentle rules for happy c++ programmining -  Tag types can be problematic if declared inside the type declaration
</summary>

[A Few Gentle Rules (\*) but One in Particular](https://youtu.be/26-7V8M1TmQ)

> - Avoid raw pointers
> - Don't use objects after they have been moved
> - Avoid leading underscores in identifiers
> - Beware of potentially catastrophic results of impromptu type declarationsfrom embedded elaborated type specifiers.
>
> In other words, don't dclare types inside other declarations.

```cpp
namespace B
{
  std::vector<struct Foo*> vec; // new types
  void f(struct Arg*); // new type
}
```

it's better to declare them outside

```cpp
namespace B
{
  struct Foo;
  struct Arg;
  std::vector<Foo*> vec; // new types
  void f(Arg*); // new type
}
```

if we declare the type inside something else, a look-up is done, and if the type isn't found, then we create one. so if the type is added somewhere outside, then we suddenly change the meaning!

```cpp
struct Foo;
namespace B
{
  std::vector<struct Foo*> vec; // doesn't create a new type
}
```

this is relvent because of Tag types (for strong types). we see code like this. which is actually a minefield.

```cpp
template <typename T, typename Tag>
struct StrongType
{
  T value;
};
using Price = StrongType<std::uint64_t, struct PriceTag>;
using Qty = StrongType<std::uint64_t, struct QtyTag>;

static_assert(not std::is_same_v<Price, Qty>);
```

instead, it should be done like this:

```cpp
struct PriceTag;
using Price = StrongType<std::uint64_t, PriceTag>;
struct QtyTag;
using Qty = StrongType<std::uint64_t, QtyTag>;
```

```cpp
struct PriceTag;
using Price = StrongType<std::int, PriceTag>;
namespace B
{
  using Price = StrongType<std::int, struct PriceTag>;
}

static_assert(not std::is_same_v<Price, B::Price>); // this fails! same type

```

here's another failure! elaborated-type-specifies might introduce into the ecnlosing namespace,

```cpp
namespace A
{
  using Price = StrongType<std::int, struct PriceTag>;
}

namespace B
{
  using Price = StrongType<std::int, struct PriceTag>;
}
static_assert(not std::is_same_v<A::Price, B::Price>); // this fails! same type
```

there also probelm of how types affect ADL, so if the tag type is introduced from another scope, then we might not use the specialized version. this can also lead into ODR violations in linking.

> Summary:\
> Elaborated type specifier declarations can result in
>
> - Changed types and signatures
> - Unexpected Namespace injections
> - Difficult debugging
> - Linker errors
> - ADL flippancy
> - ODR violations
>
> Just Say No! prefer strong types that don't use tags. But if you must use tags, declare them separately

</details>

## A Simple GUI Programming Setup for Beginners - Jussi Pakkanen

<details>
<summary>
Meson GUI - graphical apps and games.
</summary>

[A Simple GUI Programming Setup for Beginners](https://youtu.be/WwzO7TjgBOg)

is c++ really that bad and hard for GUI applications?

looking at "one lone coder" header only pixel engine. there is an issue with loading PNG images.
Messon is a tool to wrapover and simplify the process.

</details>

## Answering a Question From My Talk… - Jens Weller

<details>
<summary>
Follow up on a question about boost usage.
</summary>

[Answering a Question From My Talk…](https://youtu.be/qqGl03QisI4)

> "Are you saying there is a correlation between Boost usage and C++17 usage, rather than with C++98 usage?"

part of understaning the c++ community. trying to figure out the correlation

</details>

## Classes With Many Fields - Stanisław J. Dobrowolski

<details>
<summary>
Exploring complexity of classes.
</summary>

[Classes With Many Fields](https://youtu.be/35y3OrsKq8c)

testing how many classes in big projects have more than 4 fields. all sorts of visualizations.

</details>

## Cyclomatic Complexity pmccabe as a Refactor Aid - Michael Wells

<details>
<summary>
A command line tool to help with factorting.
</summary>

[Cyclomatic Complexity pmccabe as a Refactor Aid](https://youtu.be/TmWc4mdLmvQ)

measureing and managing complexity. cyclomatic complexity of control flow. **pmccabe** is a tool that can help us measure complexity and determine how much effort will be required to test or change something. also guessting the `code maat` tool to find hot-spots in the code based on git data.

</details>

## Designated Initializers: Remembering Every Struct Member in Declaration Order Is Hard - Brian Ruth

<details>
<summary>
Aggregate initializing: C99 compared to C++20
</summary>

[Designated Initializers: Remembering Every Struct Member in Declaration Order Is Hard](https://youtu.be/dKgTet55Pn8)

c99 had desinated initlizars for structs, it also had array designators. for some values, for ranges, etc... . but i'ts posable to mix stuff, some default, some ordered, some named, some self referential.

c++20 added to standard designated initlizaers, but more restricted, the order matters much more.

</details>

## Exhuming "Castlequest" - Arthur O'Dwyer

<details>
<summary>
The Tale of finding lost adventure games
</summary>

[Exhuming "Castlequest"](https://youtu.be/-4LN7Uy46-0)

Arthur tried to find who wrote the code for the text adventure "Castlequest", which was part of the GEnie network thing. it's part of some convoluted web of simillar games, there's a tale of getting the copyright documents from the library of congress. and it has the source code!

it's now up on github! [castleQuest](https://github.com/Quuxplusone/Castlequest) with instructions on how to compile and run it.

</details>

## Finding Nemo, or Evolution of a for Loop - Arseniy Zaostrovnykh

<details>
<summary>
Evolution of a for Loop
</summary>

[Finding Nemo, or Evolution of a for Loop](https://youtu.be/jXWz0kwtxrQ)

predicates and searching for existing in different versions of c++.

metrics are lines of codes, number of tokens, words, and cognitive load/complexity.

we start with c++98, regular index based loop. in c++11, we have `for auto`, and also library algorithm like `std::any_of`. in c++20 we have ranges, which makes things simpler.

```cpp
bool isNemo(int fish)
{
  return fish == 34;
}

namespace cpp98
{
  bool findNemo(std::vector<int> const& fish)
  {
    for (int i = 0;  i < fish.size(); ++i)
    {
      if (isNemo(fish[i]))
      {
        return true;
      }
    }
    return false;
  }
}

namespace cpp11_lang
{
  bool findNemo(std::vector<int> const& fish)
  {
    for (auto f : fish)
    {
      if (isNemo(f))
      {
        return true;
      }
    }
    return false;
  }
}

namespace cpp11_lib
{
  bool findNemo(std::vector<int> const& fish)
  {
    return std::any_of(fish.begin(), find.end(), isNemo);
  }
}

namespace cpp20
{
  bool findNemo(std::vector<int> const& fish)
  {
    return std::ranges::any_of(fish, isNemo);
  }
}
```

but because all the code is still valid, we might have old style code hiding in our code base, so we should update it, luckily, there is static analysis tools to help us.

</details>

## FourCCs Done Right - Ben Deane

<details>
<summary>
Enum structs
</summary>

[FourCCs Done Right](https://youtu.be/ytu_yMT0mbo)

four cc is a way to identify file formats. like h264, jfif. commonly used for av codecs. supposed to be somewhat human readble. but we need someway to put this into code.

this is valid in C (and c++), but problematic. don't forget about endianess!

```cpp
unsigned int pig = 'Gemi';
```

we could use class in c++98, but that requires us to tell the compiler stuff.

now we can use enum structs and use string literals,

```cpp
enum struct FourCC : std::uint32_t{};

consteval auto operator""_4cc(const char* std::size_t);

constexpr auto to_fourcc(std::uint32_t);
constexpr auto to_fourcc(const std::string&);
constexpr auto to_fourcc(std::string_view);
```

and this allows us to use switch cases, and we don't have to write so many boiler plate code.

```cpp
switch(code){
  case "Gemi"_4cc:
  case "iota"_4cc:
  case "min"_4cc:
  case "Lulu"_4cc:
}
```

but sometimes we need to convert, like when we work with legacy code, we can overload the unary `+` operator to cast into the numeric value.

</details>

## Homogenous Variadic Functions - A Lightning-Library Approach in ~11.54 sec/LOC - Tobias Loew

<details>
<summary>
A workaround to allow for Homogenous Variadic functions.
</summary>

[Homogenous Variadic Functions - A Lightning-Library Approach in ~11.54 sec/LOC](https://youtu.be/PI-E2T1FBdw)

Homogenous variadic function (HFV):

function (overload-set) with an arbitrary number of arguments, all of the same type.

```cpp
Foo(T t1);
Foo(T t1,T t2);
Foo(T t1,T t2, T t3);
Foo(T t1,T t2, T t3,T t4);
//...
```

and the propsal was buried,

```cpp
Foo(T... ts);
```

a walk around is to combiner parameters pack and `requires` clause.

```cpp
template <class ... Args>
requires (std::is_convertible_v<Args, T> && ...) // all args are convertible to T
void Foo (Args&&... args)
{
  // ...
}
```

the question of which overload is called when we have conflicting cases? how is this resolved?

```cpp
requires (std::is_convertible_v<Args, int> && ...) // all args are convertible to int
void Foo (Args&&... args)
{}

requires (std::is_convertible_v<Args, float> && ...) // all args are convertible to float
void Foo (Args&&... args)
{}

Foo(1,2,3); //(a)
Foo(0.5f, -2.4f); //(b)
Foo(1.5f,3); //(c)
```

**turns out that all the cases are ambiguous.**

if we use a `mp_list` as a way to hold the types, we can make the first two cases work, and the last case fail as ambiguous, as we want.

</details>

## Numerical Differentiation ++ - Ian Bell

<details>
<summary>
The Agony of Generic Programming
</summary>

[Numerical Differentiation ++](https://youtu.be/irEtRmb4OME)

getting derivatives, differentiations. it's important for machine learning.

```cpp
#include <complex>
#include <iostream>
#include <iomanip>

int main()
{
  double x = 3.0;
  double h = 1e-100; //step
  std::cout
  << std::setprecision(18)
  << std::cox(x)
  << std::endl
  << std::sin(std::complex<double>(x,h)).imag()/h
  << std::endl;
}
```

both are the same value, more or less.

some way to connect the c+++ values to a C API. the solution was to combine variants, visitors and concepts to make something.

</details>

## One Friend Ain't Enough - Jody Hagins

<details>
<summary>
The quest for friend `Ts...` a proposal to allow variadic friends decleration.
</summary>

[One Friend Ain't Enough](https://youtu.be/zvWCgiVvpPU)

traditional friendship, grants access to data inside the class.

```cpp
class Blarg;
class Blip
{
  private:
  int data_;
  friend Blarg;
  public:
  //...
};
```

the passkey Idiom allows a limited way, which can also be templated.

```cpp
class Blarg;
class Blip
{
  private:
  int data_;
  public:
  int data(Passkey) const {return date_;}
};
```

Discrete "friendship"

```cpp
class Blip
{
  private:
  int data_;
  public:
  void bippity(Passkey<Foo>);
  void boppity(Passkey<Bar>);
  void boo(Passkey<Baz>) const;
};
```

Passkey implementation

```cpp
class PasskeyBase // aggregate initialization madness - to avoid the {} initialize
{
  protected:
  PasskeyBase() = default;
};

template <typename T>
class Passkey: PasskeyBase
{
  private:
  friend T;
  constexpr Passkey() = default;
  constexpr Passkey() & operator=(Passkey&&) =delete; // no move constructor
}
```

granting access to more than one class

```cpp
class Blarg;
class Baz;

class Blip
{
  private:
  int data_;
  public:
  int data(Passkey<Blarg, Baz>) const {return x;}

};
```

so we need to modify the key again. but this doesn't work.

```cpp
template <typename ... Ts>
class Passkey: PasskeyBase
{
  private:
  friend Ts...; // ERROR!
  constexpr Passkey() = default;
  constexpr Passkey() & operator=(Passkey&&) =delete; // no move constructor
}
```

another example with CRTP. trying to get variadic CRTP. but we can't get variadic friends.

</details>

## Stdfwd - Forward Declarations for C++ Standard Library - Oleh Fedorenko

<details>
<summary>
Forward Declarations for C++ Standard Library
</summary>

[Stdfwd - Forward Declarations for C++ Standard Library](https://youtu.be/lv6Q5P3v-tg)

still mostly undefined behavior, not sure why this is good? maybe it helps with keeping the size down instead of precompiling them?

</details>

## Surveying the Community – What Could Possibly Go Wrong - Anastasia Kazakova

<details>
<summary>
Explaining some methodology of the community survey.
</summary>

[Surveying the Community – What Could Possibly Go Wrong](https://youtu.be/Hg1gIxRr4eI)

understaning the biases. changes between survies, understanding how the question wording effects the answers.

</details>

##

[Main](README.md)
