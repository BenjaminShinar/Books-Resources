<!--
// cSpell:ignore codecov cppcoro dogbolt decompiler Lippincott bfloat stdfloat -Wnrvo nrvo
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

## C++ Weekly - Ep 401 - C++23's `chunk_view` and `stride_view`

<details>
<summary>
range methods that return a block of elements.
</summary>

[C++23's chunk_view and stride_view](https://youtu.be/3ZeV-F1Rbaw?si=QpoyEmENmaIKExh9)

1. <cpp>chunk</cpp> view - a range of ranges with a runtime argument. a "chunk" of n or less elements from the range.
2. <cpp>stride</cpp> view - a range of ranges with a runtime argument. takes the first element from each chunk of n.

```cpp
int main()
{
  std::array a{1,2,3,4,5,6,7,8};

  // will chunk into [1,2,3][4,5,6][7,8]
  for (const auto &chunk : std::ranges::views::chunk(a,3))
  {
    for (const auto &elem : chunk)
    {
      std::cout << elem << ',';
    }
    std::cout << '\n';
  }

  // will print [1,4,7]
  for (const auto &elem : std::ranges::views::stride(a,3))
  {
    std::cout << elem << '\n';
  }
}
```

</details>

## C++ Weekly - Ep 402 - Reviewing My 25 Year Old C++ Code (IT'S BAD!)

<details>
<summary>
Looking at code from Jason's university days (1997/8).
</summary>

[Reviewing My 25 Year Old C++ Code (IT'S BAD!)](https://youtu.be/7kqxYZKm64A?si=eCcrRuru5lwaud-h)

lots of formatting stuff, using the correct types and not making things member functions. being careful of using an int when<cloud>std::size_t</cloud> should be used. making sure the memory isn't leaked and the ownership is clear.

</details>

## C++ Weekly - Ep 403 - Easier Coroutines with CppCoro

<details>
<summary>
A helper library for coroutines.
</summary>

[Easier Coroutines with CppCoro](https://youtu.be/TWoZ9SGIE9o?si=p5RHxYcwurWiiH22)

Coroutines were added in C++20, but without library support. C++23 provided <cpp>std::generator</cpp> as a helper template. CppCoro is a library with helper and tools to make Coroutines more usable. it is also available on compiler explorer, so it's easier to experiment with.

an example of Fibonacci numbers:

```cpp
cppcoro::generator<const std::uint64_t> Fibonacci()
{
  std::uint64_t a = 0, b = 1;
  while (true)
  {
    co_yield b;
    a = std::exchange(b, a + b); // instead of temporary variables
  }
}

int main()
{
  for (auto i : Fibonacci())
  {
    if (i > 1'000'000) break;
    std::println("{}", i);
  }
}
```

However, coroutines can't be used in constexpr. They exist (by defintion) on the heap, and we cannot choose how it's done, there is no custom memory allocation (<cpp>std::pmr</cpp>).

other features include:

- asynchronous generator for yielding lazily created values
- Tasks
- Networking support
- File support

</details>

## C++ Weekly - Ep 404 - How (and Why) To Write Code That Avoids `std::move`

<details>
<summary>
Avoiding work by understanding how the compiler works.
</summary>

[How (and Why) To Write Code That Avoids std::move](https://youtu.be/6SaUwqw4ueE?si=7Imdff_WXt_96IRy)

Avoiding <cloud>std::move</cloud> calls. making use of named return value optimizations.

```cpp
struct Lifetime{
  Lifetime() {std::puts("Lifetime() // default ctor")}
  Lifetime(const Lifetime & other) {std::puts("Lifetime(const Lifetime & other) // copy ctor")}
  ~Lifetime() {std::puts("~Lifetime() // dtor")}
  // more
  int member_data;
};

int main()
{

  auto make_lifetime = [](const int value) {
    Lifetime l;
    l.member_data = value;
    return l;
  };

  {
    std::array<Lifetime, 2> a{}; // default ctor
    // 2 object destroyed
  }

  {
    Lifetime l1;
    l1.member_data = 41;
    Lifetime l2;
    l2.member_data = 42;
    std::array<Lifetime, 2> b{l1,l2}; // copy ctor
    // 4 objects destroyed
  }

  {
    Lifetime l1;
    l1.member_data = 41;
    Lifetime l2;
    l2.member_data = 42;
    std::array<Lifetime, 2> c{std::move(l1),std::move(l2)}; // move ctor
    // 4 objects destroyed
  }

  {
    std::array<Lifetime, 2> d{make_lifetime(42),make_lifetime(43)}; // copy elision
    // 2 object destroyed
  }
}
```

it's always better to rely on simple composable function instead of <cpp>std::move</cpp>.

</details>

## C++ Weekly - Ep 405 - Dogbolt: The Decompiler Explorer

<details>
<summary>
De-compile binary files into source code.
</summary>

[Dogbolt: The Decompiler Explorer](https://youtu.be/h3F0Fw0R7ME?si=-WkvV_SK_zzXIlRO)

a [website](https://dogbolt.org/?) similar to compiler explorer, it can take an executable binary file and then runs it through engines and de-compiles it back into source code.

</details>

## C++ Weekly - Ep 406 - Why Avoid Pointer Arithmetic?

<details>
<summary>
Examples of pointer arithmetics going wrong.
</summary>

[Why Avoid Pointer Arithmetic?](https://youtu.be/MsujPM2wDmk?si=B540jHpG-vXF9uzv)

adding to a pointer means adding the size of the pointed object. so moving a character pointer moves one byte, int pointer moves 4 bytes at a time.

```cpp
struct S;
void use_ptr(S* ptr)
{
  ptr += 1;
}
```

pointer arithmetics is ripe for bugs of accessing memory outside the actual variable, can happen when passing arrays, when parsing strings. it's better to use things like <cpp>std::string_view</cpp> or <cpp>std::span</cpp>. and it's always important to test using "fuzzy testing".

</details>

## C++ Weekly - Ep 407 - C++98 Code Restoration

<details>
<summary>
Continue from episode 402.
</summary>

[C++98 Code Restoration](https://youtu.be/A5haG_UCbRI?si=RsUhgejnlWsM7j3P)

Working on C++98 code, using tools that were available at the time. we first add some tests (in a new project), needing to resolve dependencies across them.

- adding code to source control.
- upgrade build tools as much as possible (what was available then).
- capture current state with tests before making changes.
- split long files to headers.
- keep changes to minimal at start.
- add header guards

bugs are ok for now, we want to have something that builds, links and is testable.

we can set the ".gitattributes" file to change to way git adds line endings on legacy files.

once we get it to a stable state, we can start modifying the code and use best practice. we can use the standard library and proper types (<cpp>std::string</cpp> rather than <cpp>char \*</cpp> pointers, <cpp>bool</cpp> rather than <cpp>int</cpp>). we can move to using references instead of pointers, and make sure to use <cpp>const</cpp> when needed. templates existed back then, so we can use them instead of void pointers. it's ok to discover bugs, we just need to have tests that monitor them.\
Since this is an arithmetical project, we use operator overloading instead of calling named functions like `divide` and `add`.

</details>

## C++ Weekly - Ep 408 - Implementing C++23's constexpr unique_ptr in C++20

<details>
<summary>
<cpp>std::unique_ptr</cpp> is now supported a compile time.
</summary>

[Implementing C++23's constexpr unique_ptr in C++20](https://youtu.be/p8Q-bapMShs?si=FiHoZUqe2fPW7njJ)

<cpp>constexpr</cpp> and <cpp>std::unique_ptr</cpp> finally work together in C++ 23. until now it didn't have constant expression constructor and destructor. generally, if the code exists in a header file, it usually easy to make it <cpp>constexpr</cpp>.

</details>

## C++ Weekly - Ep 409 - How To 2x or 3x Your Developer Salary in 2024 (Plus Some Interviewing Tips!)

<details>
<summary>
Tips for job hunting!
</summary>

[How To 2x or 3x Your Developer Salary in 2024 (Plus Some Interviewing Tips!)](https://youtu.be/jWhFuK7J5HY?si=KL9QUqvMBWem6uzo)

Networking (the human kind) is important, reputation, making your name known. become involved in the community. find out what you are passionate about, and then find opportunities to discuss in front of your crowed. this is **brand building**.\
Be explicit about your goals - "I am looking for a job".

Interviewing - be aware of the Dunning-Kruger (it goes both ways!), be honest, don't undersell and don't oversell (avoid terms like "expert", "guru"). don't be afraid to admit you don't know the answer, but demonstrate curiosity. Ask questions back and be engaged.\

> - what is your testing culture?
> - what is the CI setup?
> - what are the training budget and learning opportunities?
> - can I keep speaking/contributing to the community?

Practice speaking about the things you are passionate about.

</details>

## C++ Weekly - Ep 410 - What Are Padding and Alignment? (And Why You Might Care)

<details>
<summary>
Alignment and Padding in member layouts.
</summary>

[What Are Padding and Alignment? (And Why You Might Care)](https://youtu.be/E0QhZ6tNoRg?si=rww5ZQvWvl8lmFwv)

alignment shows where in memory a variable can start.

```cpp
#include <type_traits>

struct S {
  char a; // 1
  int i; // 4
  char b; // 1
  int j; // 4
  long long l; // 8
  char c; // 1
  // total of 3*1 + 4*2 + 8 = 19
};
int main()
{
  //return std::alignment_of_v<int>; // return 4
  return sizeof(S); // returns 32
}
```

| Alignment | char | int | long long |
| --------- | ---- | --- | --------- |
| 00        | V    | V   | V         |
| 01        | V    | X   | X         |
| 02        | V    | X   | X         |
| 03        | V    | X   | X         |
| 04        | V    | V   | X         |
| 05        | V    | X   | X         |
| 06        | V    | X   | X         |
| 07        | V    | X   | X         |
| 08        | V    | V   | V         |

we can create a struct and based on the order of the members, we get different sizes. some compiler have an option to warn about padding (`-Wpadded` in clang). **changing layout will break the ABI.**

</details>

## C++ Weekly - Ep 411 - Intro to C++ Exceptions

<details>
<summary>
General Suggestions for using exceptions.
</summary>

[Intro to C++ Exceptions](https://youtu.be/uE0h79vB-rw?si=tNPZ5HvW-OlYctOJ)

exceptions have a bad reputation, but they aren't that bad.

1. don't overuse them
2. don't "return by exception" - don't make it the normal control flow
3. aim for "exceptional cases"
4. compilers are very bad at optimizing around exceptions
5. we usually catch with <cpp>const</cpp> and as a reference

we can throw any kind of value, not just <cpp>std::exceptions</cpp>. we can have catch cases based on types, or use `catch(...)` as a default case.

Lippincott function for centralized exception handling

```cpp
void handler()
{
  try
  {
    throw; // rethrow currently in-flight exception
  }
  catch(const std::runtime_error &e)
  {

  }
  catch(const std::exception &e)
  {

  }
  catch(...)
  {

  }
}

int main()
{
  try{

  }
  catch(...)
  {
    handler(); // handle all exceptions the same way!
  }
}
```

</details>

## C++ Weekly - Ep 412 - Possible Uses of C++23's [[assume]] Attribute

<details>
<summary>
trying to find a potential use-case for the [[assume]] attribute.
</summary>

[Possible Uses of C++23's [[assume]] Attribute](https://youtu.be/Frl8XKhvA4Q?si=0sWAyEj4GUonnq83)

give the compiler a hint about the nature of the code. we can tell the compiler ahead of time that we know something which we don't have as part of the type system, and then the compiler can optimize with it in mind.

```cpp
int do_work(int x)
{
  if (x<5) {
    return x +10;
  }
  else {
    return x - 10;
  }
}

int main(int argc, const char*[])
{
  [[assume(argc == 5)]];
  return do_work(argc);
}
```

a better example uses the assumption attribute to skip a null check.

```cpp
int get_value(int * ptr)
{
  if (ptr == nullptr)
  {
    throw "oops!";
  }
  return *ptr;
}

int * get_ptr();

int main()
{
  int *ptr = get_ptr();
  [[assume(ptr != nullptr)]];
  return get_value(ptr);
}
```

</details>

## C++ Weekly - Ep 413 - (2x Faster!) What are Unity Builds (And How They Help)

<details>
<summary>
Using Unity Build to increase build speed.
</summary>

[(2x Faster!) What are Unity Builds (And How They Help)](https://youtu.be/POYVF6urMwg?si=FYPHn6wPXH-J-c76)

we usually compile source files one by one, we can maybe compile them in parallel (depending on the number of cores), but we always need to parse the same headers again and again.\
A unity Build takes the files and concatenates their content together, which can potentially speed up the process. it can also act as a kind of LTO (link time optimization), since the compiler has more information about the code. another upside is that it detects ODR violations much quicker.

this is a built in option in CMake, we can set it globally and then make exceptions for specific target projects.

we might get a warning about redefinition of macro across files (but we shouldn't be using macro anyway).

</details>

## C++ Weekly - Ep 414 - C++26's Placeholder Variables With No Name

<details>
<summary>
special way to re-use the underscore as a variable name.
</summary>

[C++26's Placeholder Variables With No Name](https://youtu.be/OZ1gNuF60BU?si=yTfYAzQ0bRV32AIb)

allows usage of placeholders for unused variables.

```cpp
std::tuple<int, double, float> get_values();

int main()
{
  const auto &[count, volume, rate] = get_values();

  return count;
}
```

we could one instance of `_` underscore in c++23, but starting in c++26, multiple uses are allowed.

```cpp
int main()
{
    const auto &[count, _, _] = get_values();
    int _ = 2;
    float _ = 3.0;

  //return _; // not allowed, ambiguous
  return count;
}
```

</details>

## C++ Weekly - Ep 415 - Moving From C++98 to C++11

<details>
<summary>
Continue from episode 407.
</summary>

[Moving From C++98 to C++11](https://youtu.be/84Zy1D8MWaI?si=4TJ_OnRadt4mVa9j)

first we move the project to the toolchain, and then we can start upgrading the code to c++11. removing casts, using standard containers. using <cpp>auto</cpp> sometimes when we don't care about the types and we don't want conversions (we get the compiler to warn us about them).

</details>

## C++ Weekly - Ep 416 - Moving From C++11 to C++14

<details>
<summary>
Continuing from previous episode.
</summary>

[Moving From C++11 to C++14](https://youtu.be/_Rq8gWimRcA?si=VvdY9F_N2komXyg1)

next we move from C++11 to C++14. some of the stuff could be done using the automatic tools.

we can add Cpp attributes, such as c++11 <cpp>[[noreturn]]</cpp> and <cpp>[[deprecated]]</cpp> (<cpp>[[nodiscard]]</cpp> is from cpp++17).\
Each time we move in standards, we can use more external libraries, and in our case, we can use the testing framework of <cpp>Catch2</cpp>. we add it to the CMake file (some stuff to uncomment).

in the test.cpp file we previously had custom unit tests code.

```cpp
int main()
{
  bool all_passed = true;
  runt_test("3 + 2", RationalNumber(5,1), all_passed);
  runt_test("(3 + 2)", RationalNumber(5,1), all_passed);
  // more tests
  if (all_passed) {
    return EXIT_SUCCESS;
  }
  else {
    return EXIT_FAILURE;
  }
}
```

which we can replace with the test framework, which better integrates with the IDE. we also can create constexpr tests, which will prevent the code from compiling.

```cpp
TEST_CASE("addition")
{
  CHECK(evaluate("3 + 2") == RationalNumber(5,1));
  CHECK(evaluate("(3 + 2)") == RationalNumber(5,1));
}

TEST_CASE("constexpr constructor")
{
  STATIC_REQUIRE(RationalNumber(5,1) == RationalNumber(5,1));
  STATIC_REQUIRE(RationalNumber(1,1) + RationalNumber(2,1) == RationalNumber(3,1));
  //STATIC_REQUIRE(RationalNumber(5,1) == RationalNumber(4,1)); // will fail compilation
}
```

we can make some more code to be <cpp>constexpr</cpp>, but for now it's mostly preparation for the future.

</details>

## C++ Weekly - Ep 417 - Turbocharge Your Build With Mold?

<details>
<summary>
A linker that utilizes mutiple cpus and parallel processing
</summary>

[Turbocharge Your Build With Mold?](https://youtu.be/gOBbu2dL_R8?si=C84ki0_Uy2Z3ARp3), [github](https://github.com/rui314/mold)

getting faster linking, especially when multiple object files rely on the same libraries. we switch linkers with by passing the `-fuse-ld=<linker>` flag (`-fuse-ld=mold` in our case). it also reduces memory usage.

it can also help when building debug versions, which are usually slower.

</details>

## C++ Weekly - Ep 418 - Moving From C++14 to C++17

<details>
<summary>
Another upgrade
</summary>

[Moving From C++14 to C++17](https://youtu.be/yL0DWa2LxNU?si=XX8iAq79FGWh8pyQ)

we can start by adding the attributes, mostly the <cpp>[[nodiscard]]</cpp> one. when we move to using <cpp>std::string_view</cpp> (pass by value) we start getting into issues of lifetime management. we also replace <cpp>atoi</cpp> with <cpp>std::from_chars</cpp>, there is also an added bonus of better <cpp>constexpr</cpp> support. we can also add fuzz-testing to check our string tokenizer.

</details>

## C++ Weekly - Ep 419 - The Important Parts of C++23

<details>
<summary>
overview of the major changes in C++23.
</summary>

[The Important Parts of C++23](https://youtu.be/N2HG___9QFI?si=0GVPU_ZMFLrpPPit)

- <cpp>std::print</cpp> and <cpp>std::println</cpp> - easily format argument
- <cpp>import std;</cpp> - replace headers with modules. still not supported in compiler explorer.
- <cpp>std::stacktrace</cpp> - access the stacktrace.
- <cpp>std::flat_map</cpp>, <cpp>std::flat_set</cpp> - container adapters for contiguously allocated map and set behavior. somehow not constexpr yet. default uses vectors.
- multidimensional subscript. use commas inside the `[]` operator.
- ranges upgrades - many stuff, includes <cpp>zip</cpp> to iterate over multiple containers at a time.
- constexpr "cmath" - except for trigonometry functions (c++26).
- <cpp>std::expected</cpp> - an alternate error handling mechanism. like an optional pair. the downside is that it interferes with return value optimization and would require explicit move semantics.
- generators for coroutines
- <cpp>md_span</cpp> - multiple dimension span (not owning the data itself)
- explicit `this` (deducing `this`)
- size literal suffix - signed and and unsigned size_t suffixes.
- <cpp>std::start_lifetime_as</cpp> - allows us to avoid <cpp>std::reinterpret_cast</cpp> in another case. tells the compiler there is an object at that memory point, which clears up some undefined behavior cases.

```cpp
#include <print>
// import std;
int main(int argc, const char[]*)
{
  std::println("{} argument passed to main", argc);
}

void foo()
{
  std::println("{}", std::stacktrace::current());
}
```

</details>

## C++ Weekly - Ep 420 - Moving From C++17 to C++20 (More constexpr!)

<details>
<summary>
continuing from Episode 418.
</summary>

[Moving From C++17 to C++20 (More constexpr!)](https://youtu.be/s2XWAxbxk9M?si=SOta0gZ_qCS8QgYZ)

we start by improving the error handling, this is noticeable since we have fuzzy testing to generate test cases that fail.

replacing runtime asserts with <cpp>std::throw</cpp>, sometimes we see division by zero, sometimes we get inputs that overflow the limits of the integer type.

in the cmake file, we set ourselves to to C++20.

now we can continue making things <cpp>constexpr</cpp>. sometimes all it takes is moving code into header file. we also add constexpr test. we just have a problem with <cpp>std::from_chars()</cpp> which wasn't compile time until C++23. so we create a similarly named function that is compile time compatible.

```cpp
template<std::integral Type>
[[nodiscard]] constexpr Type from_chars(std::string_view input) -> auto
{
  Type result = 0;
  for (const char digit: input) {
    result *= 10;
    if (digit >= '0' && digit <= '9')
    {
      result += static_cast<Type>(digit - '0');
    }
  }
  return result;
}
```

at this point, it means that we have only header files, so our project is now a header only library. which means some changes to the CMake file.

C++20 also has <cpp>consteval</cpp>, which forces a function to be evaluated at compile time. we can use it to create a user defined suffix and force the compiler to evaluate our test cases during compilation.

```cpp
consteval auto operator""_rn(const char* str, std::size_t len) noexcept
{
  return evaluate(std::string_view(str, len));
}
```

we can take advantage the three way comparison (spaceship) operator to get some default implementation. **which is wrong!**.

</details>

## C++ Weekly - Ep 421 - You're Using optional, variant, pair, tuple, any, and expected Wrong!

<details>
<summary>
THings that break return value optimizations
</summary>

[You're Using optional, variant, pair, tuple, any, and expected Wrong!](https://youtu.be/0yJk5yfdih0?si=5ps0Zl5n30DMKAba)

using the lifetime struct which prints when it's created or destroyed to show some interesting behavior.

in each of these cases, we create and destroy two objects, which isn't what we want. we want return value optimization, but we don't get it!

```cpp
std::optional<Lifetime> get_value_move()
{
  std::optional<Lifetime> retval;
  retval = Lifetime{42}; // move ctor

  return retval;
}

std::optional<Lifetime> get_value_copy()
{
  std::optional<Lifetime> Lifetime myLifeTime{42};
  retval = myLifeTime; // copy ctor

  return retval;
}

std::optional<Lifetime> get_value_move2()
{
  return Lifetime{42}; // also move
}

std::optional<Lifetime> get_value_move3()
{
  Lifetime l{42}
  return l; // also move
}

int main()
{
  get_value();
}
```

if we use <cpp>.emplace()</cpp> on our value, then it's fine. but it's horrible to write.

```cpp
std::optional<Lifetime> get_value_emplace()
{
  std::optional<Lifetime> l;
  l.emplace(42);
  return l; // here we get optimization
}
```

a one liner return statement works. but only if we get the type correct.

```cpp
std::optional<Lifetime> get_value()
{
  //return Lifetime{42}; // not RVO
  return std::optional<Lifetime>{42}; // RVO
}
```

we could try `return {42};`, which would work, but if we mark our single parameter constructor as <cpp>explicit</cpp> (which we should), then it doesn't work anymore.

this issue get more serious when using c++23 <cpp>std::expected</cpp>.

</details>

## C++ Weekly - Ep 422 - Moving from C++20 to C++23

<details>
<summary>
continuing from episode 420.
</summary>

[Moving from C++20 to C++23](https://youtu.be/dvxj39gZ22I?si=fOigTKMKfPVAwCSu)

we can put attribute on lambda, so we use <cpp>[[noreturn]]</cpp> on a lambda that throws exception.

we have a new tool for error handling instead of exceptions, <cpp>std::expected</cpp> allows us to return an object of a different type, which also makes it possible to work in compile time. **but we don't use it for now**.

</details>

## C++ Weekly - Ep 423 - Complete Guide to Attributes Through C++23

<details>
<summary>
Going over the existing attributes.
</summary>

[Complete Guide to Attributes Through C++23](https://youtu.be/BpulWncdn9Y?si=vltQs2GgLWxpoTSK),
[cppreference](https://en.cppreference.com/w/cpp/language/attributes).

- <cpp>[[noreturn]]</cpp> - tell the compiler a function will not return (abort, terminate, throws).
- <cpp>[[carries_dependency]]</cpp> - very weird and not commonly used. something about memory order. also <cpp>std::kill_dependency</cpp> which removes it.
- <cpp>[[deprecated("reason")]]</cpp> - warning on compile.
- <cpp>[[fall_through]]</cpp> - inside switch statements, silence warning on cases that do something and fall through.
- <cpp>[[no_discard]]</cpp> - warning about functions or types which must be captured.
- <cpp>[[maybe_unused]]</cpp> - avoid the warning about unused variable.
- <cpp>[[likely]]</cpp>, <cpp>[[unlikely]]</cpp> - hints to the compiler which case to optimize to.
- <cpp>[[no_unique_address]]</cpp> - optimizing for empty objects inside our struct.
- <cpp>[[assume]]</cpp> - place pre-conditions which aren't enforced by the type system, allows for optimizations.

</details>

## C++ Weekly - Ep 424 - `.reset()` vs `->reset()`

<details>
<summary>
A thing that can confuse us.
</summary>

[.reset() vs ->reset()](https://youtu.be/HgPfbYfV9eE?si=5lO3tmaeIx9wXyTL)

anything that is a pointer-like thing to a pointer-like thing. are we changing the internal value or the holder object.

there's no static analysis for this.

</details>

## C++ Weekly - Ep 425 - Using string_view, span, and Pointers Safely!

<details>
<summary>
Undefined Behavior with non-owning views
</summary>

[Using string_view, span, and Pointers Safely!](https://youtu.be/cUvdtLTJeec?si=dQBUdW5MkVuQoigO)

non-owning "views" into stuff.

```cpp
std::string_view get_data() {
  std::string result {"Check out this data string"};
  return result;
}
```

we have a lifetime issue, and only very new compilers know to identify it as a warning (we can run the address sanitizer, and then we'll catch the problem). our code returns a string view to an object that no longer exists.

It's very easy to create this bug and have undefined behavior.

we can also fix this issue with <cpp>constexpr</cpp>, and we need to have tests for it (when possible), since it doesn't allow for un-defined behavior.

</details>

## C++ Weekly - Ep 426 - Lambdas As State Machines

<details>
<summary>
Thinking of Lambdas in a different way.
</summary>

[Lambdas As State Machines](https://youtu.be/fZe7gNgjV4A?si=sB6D2fjcWG9U07D_)

a lambda with a capture and mutable carries it's own state.

```cpp
int main()
{
  auto fib = [i = 0, j = 1](){
    i = std::exchange (j, i + j);
    return i;
  };

  fib(); // 1
  fib(); // 1
  fib(); // 2
  fib(); // 3
  fib(); // 5
}
```

we can create a more complicated example, something that parses text into numbers. we mark our lambda as mutable.

```cpp
constexpr auto make_int_parser()
{
  enum State {Startup, Numbers, Invalid};
  return [value = 0, state = Startup, is_negative = false](const char input) mutable -> std::expected<int, char> {
    switch (state) {
      case Startup:
      if (input == '-') {
        is_negative = true;
        state = Numbers;
        return value;
      }
      [[fallthrough]]
      case Numbers:
      if (input >= '0' && input <= '9') {
        value *= 10;
        value += static_cast<int>(input -'0');
        return (is_negative ? (value * -1) : value);
      } else {
        state = Invalid;
        return std::unexpected(input);
      }
      case Invalid:
      return std::unexpected(input);
    };
    return std::unexpected(input);
  };
}
```

</details>

## C++ Weekly - Ep 427 - Simple Generators Without Coroutines

<details>
<summary>
making a generator from a lambda.
</summary>

[Simple Generators Without Coroutines](https://youtu.be/F37h3FuA8kM?si=GMwg-Dp6_RCgQ92M)

generators with lambdas.

```cpp
int main()
{
  auto fib = [i=0, j=1]() mutable{return i = std::exchange(j,i+j);};

  for (const auto val : fib | std::views::take(20)) // doesn't work.
  {
    std::cout<< val '\n';
  }
  return fib();
}
```

we would want to be able to call range operators (like <cpp>std::views::take</cpp>) and have the generator return the value. but we can't do it just yet.

we need a helper utility. we take advantage of <cpp>std::iota</cpp>

```cpp
auto generator(auto func) {
  return std::views::iota(0) | std::views::transform(func);
}

int main()
{
  auto fib = [i=0, j=1](auto) mutable{return i = std::exchange(j,i+j);};

  for (const auto val : generator(fib) | std::views::take(20)) // this does work
  {
    fmt::print("{}\n");
  }
  return fib();
}

```

</details>

## C++ Weekly - Ep 428 - C++23's Coroutine Support: `std::generator`

<details>
<summary>
Now an example with actual coroutine
</summary>

[C++23's Coroutine Support: std::generator](https://youtu.be/7ZazVQB-RKc?si=O7WzWFHkxaSkSLkx)

finally we have compiler and library support for generators - resumable functions. so we can write the same example from last week.

```cpp
#include <generator>
#include <ranges>
std::generator<int> fib()
{
  int i = 0;
  int j = 1;
  while (true) {
    co_yield i = std::exchange(j, i + j);
  }
}

int main()
{

  for (const auto val : fib() | std::views::take(20))
  {
    fmt::print("{}\n");
  }
  return fib();
}
```

</details>

## C++ Weekly - Ep 429 - C++26's Parameter Pack Indexing

<details>
<summary>
Upcoming feature that simplifies paramater pack expansion and allows direct indexing.
</summary>

[C++26's Parameter Pack Indexing](https://youtu.be/wl7uWes7Sys?si=sxOIT2tf7Vhcgz8Y).

[CppReference](https://en.cppreference.com/w/cpp/language/pack_indexing), [compiler explorer](https://compiler-explorer.com/z/3vc3Yvf4o).

Until now, if we wanted to do something with a paramater at a specific location (like the 3rd parameter), we had some workaround options:

- create specific overload
- do some recursive programming and parameter expansion
- (from video comments: use <cpp>std::tie</cpp> and <cpp>std::get</cpp> to create a tuple.)

```cpp
template<typename ... Param>
void func(Param ... param)
{
  // how to do something with the 3rd parameter?
}

// explicit overload
template<typename Param0, typename Param1, typename Param2, typename ... Param>
void func(Param0 param0, Param1 param1, Param2 param2, Param ... param)
{
  // doing something with third parameter
  param2 = 42;
}

// recursive style
template<std::size_t count, typename Param0, typename ... Param> requires (count > 0)
auto &get(Param0 param0, Param ... param)
{
  return get<count-1>(param...);
}

template<std::size_t count, typename Param0, typename ... Param> requires (count == 0)
auto &get(Param0 param0, Param ... param)
{
  return param0;
}

template<typename ... Param>
void funcR(Param ... param)
{
  get<2>(param...) = 42; // modify the third parameter
}

// from the video comments
template<typename... Param>
void funcTuple(Param... param) {
    std::get<2>(std::tie(param...)) = 42;
}
```

In the upcoming C++26 standard, we can directly index to the parameter pack - we use the `[]` brackets operator, just with the `...` before it.

```cpp
template<typename ... Param>
void func26(Param ... param)
{
  // simple
  param...[2] = 42;
}
```

</details>

## C++ Weekly - Ep 430 - How Short String Optimizations Work

<details>
<summary>
A coding example of small string optimizations
</summary>

[How Short String Optimizations Work](https://youtu.be/CIB_khrNPSU?si=ihFR8bT_4PvuTSK2)

things that usually allocate, but might be able to avoid this allocation if the object is small enough.

- <cpp>std::string</cpp>
- <cpp>std::function</cpp>
- <cpp>std::any</cpp>

a small example of how it can be done. we create an internal data object which can either point to the data somewhere else (pointer and size) or be the data itself. we can use the same memory address via a union.

```cpp
struct string
{
  struct allocated_storage
  {
    char *m_data_ptr;
    std::size_t m_capacity; // allocated_size
  }; // internal things we can use.

  using small_storage = char[sizeof(allocated_storage)];


  union storage {
    allocated_storage m_alloc;
    small_storage m_small;
  };

  storage m_storage{.m_small = small_storageP{}}; // default initialize

  std::size_t m_size = 0; // current size - must be known at all times
  bool m_is_small_storage = true;

  [[nodiscard]] constexpr const char *date() const noexcept const {
    if (m_is_small_storage) {
      return m_storage.m_small;
    }
    else {
      return m_storage.m_alloc.m_data_ptr;
    }
  }

  // constructors

  constexpr allocated_storage alloc(std::size_t len, const char* data)
  {
    auto *allocated = new char[len+1]; // include null termination
    std::copy(data, data+len+1, allocated);
    return {allocated, len};
  }

  constexpr string() = default;
  constexpr string(const char* data): m_size{std::strlen(data)}
  {
    if (m_size < sizeof(small_storage))
    {
      std::copy(data, data+m_size + 1, m_storage.m_small); // small string optimization
    }
    else
    {
      m_storage.m_alloc = alloc(m_size, data);
      m_is_small_storage = false;
    }
  }
};

constexpr ~string()
{
  // C++20 constexpr destructor
  if (!m_is_small_storage)
  {
    delete[] m_storage.m_alloc.m_data_ptr;
  }
}


consteval void test_string(const char str*)
{
  // using consteval makes sure we don't have undefined behavior or memory leaks!
  string my_str{str};
}
int main()
{
  tes_string("");
  tes_string("Hello World");
  tes_string("Hello World long log string!");
}
```

</details>

## C++ Weekly - Ep 431 - CTAD for NTTP

<details>
<summary>
Non-type template parameters auto deduction.
</summary>

[CTAD for NTTP](https://youtu.be/yPB_btV8epo?si=HuHsozVRlR2jpXBK)

- CTAD - class template argument deduction.
- NTTP - non type template parameters.

```cpp

template<int i> auto get_value() {return i;} // non-type template parameter

// template<auto y> auto get_value() {return y;} // non-type template parameter - templated
// template<std::integral auto x> auto get_value() {return x;} // non-type template parameter - constrained with concept

int main()
{
  std:vector v1{1,2,3}; // vector of int
  std:vector v2{1.1,2.1,3.1}; // vector of double
  std:vector<int> v3{1.1,2.1,3.1}; // narrowing!
  v1 = v2; // ERROR

  return get_value<1>()

}
```

we can combine the two things together - CNTTP - class non-type template parameters.

```cpp
template<std::array a> auto get_value() {return a[1];}
int main()
{
  [[maybe_unused]] std::array a{1,2,3};
  //return get_value<a>();
  return get_value<{5,4,3}>(); // this also works!
}
```

the name of the function actually contains the values of the arguments in the mangled name.

</details>

## C++ Weekly - Ep 432 - Why `constexpr` Matters

<details>
<summary>
How `constexpr` helps us.
</summary>

[Why `constexpr` Matters](https://youtu.be/QZxfyGmpanM?si=1lclZpUOwcieF2Lt)

There is some pushback against `constexpr`, with many people saying that they don't have use for it since they don't have the required information in compile time.

looking back at the string parser example, if we can move some stuff to the `consteval` function, we can get some optimizations to happen, that wouldn't happen unless we made our program `constexpr`-friendly.
</details>

## C++ Weekly - Ep 433 - C++'s First New Floating Point Types in 40 Years!

<details>
<summary>
New fixed width floating point types.
</summary>

[C++'s First New Floating Point Types in 40 Years!](https://youtu.be/YM1nbexgGYw?si=deSK4di9RzCTsWNQ)

fixed width (size) floating point types

- <cpp>std::float16_t</cpp>
- <cpp>std::float32_t</cpp>
- <cpp>std::float64_t</cpp>
- <cpp>std::float128_t</cpp>
- <cpp>std::bfloat16_t</cpp> - (brain floating-point)

the new 16 bits type helps when using small devices, bfloat is used in machine learning algorithms, it optimizes the mantissa and exponents location in the layout to work better on GPU hardware (8 bits for each).

in this example we see how they differ from one another.

```cpp
#include <stdfloat>

void explore_float16(std::uint16_t val)
{
  const auto f16 = std::bit_cast<std::float16_t>(val);
  const auto bf16 = std::bit_cast<std::bfloat16_t>(val);

  std::cout << std::format("{:#016b} {:#04x} {} {}f16 (0b{:01b}'{:010b}) {}bf16 (0b{:01b}'{:08b}'{:07b})\n", val, val, val, f16, (val >> 15) & 0b1, (val >> 10) & 0b11111, (val & 0b11111_11111_1), bf16, (val >> 15) & 0b1, (val >> 7) & 0b11111_111, (val & 0b11111_11));
}

int main()
{
  explore_float16(0b0_01101_01010_10101);
  explore_float16(0b0_11110_11111_11111);
  explore_float16(0b1_11100_00111_11111);
}
```

</details>

## C++ Weekly - Ep 434 - GCC's Amazing NEW (2024) `-Wnrvo`

<details>
<summary>
New warning flag - named return value optimization.
</summary>

[GCC's Amazing NEW (2024) -Wnrvo](https://youtu.be/PTCFddZfnXc?si=-6PnRMFCSwumyf2X)

a new warning flag on GCC - warns on named return value optimization - when we thought it would happen but it doesn't.

```cpp
// compiler is allowed - named return value optimization
std::string make_sting_nrvo() {
  std::string retval = "Hello World";
  return retval;
}

// compiler is required to do RVO
std::string make_sting_rvo() {
 return "Hello World";
}

// compiler has to decide at compile time which one will be returned, but it can't
std::string make_string_opt1(const bool input) {
  std::string val1 = "Hello";
  std::string val2 = "World";
  if (input)
  {
    return val1;
  }
  else
  {
    return val2;
  }
}

int main()
{
  const auto val = make_string_rvo(); // exactly 1 string existed,
}
```

return value optimization constructs the data on the callsite, so it elides (removes) the calls to copy or move the object, but if there are two options, it isn't able to choose which one to create on the call stack. the warning will tell us when it can't optimize properly.

</details>

## C++ Weekly - Ep 435 - Easy GPU Programming With AdaptiveCpp (68x Faster!)

<details>
<summary>
GPU and parallel algorithms
</summary>

[Easy GPU Programming With AdaptiveCpp (68x Faster!)](https://youtu.be/ImM7f5IQOaw?si=gxCoJ1KXaKocDMEi)

parallel programming with GPU, using conway's game of life.

</details>
