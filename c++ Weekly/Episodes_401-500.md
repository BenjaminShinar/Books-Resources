<!--
// cSpell:ignore codecov cppcoro dogbolt decompiler Lippincott bfloat stdfloat  nrvo fimplicit Decrypter gcem vgdb cvref debian chrono -Wpadded -Wsystem -Wnrvo
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

However, coroutines can't be used in constexpr. They exist (by definition) on the heap, and we cannot choose how it's done, there is no custom memory allocation (<cpp>std::pmr</cpp>).

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

parallel programming with GPU, using Conway's game of life.

</details>

## C++ Weekly - Ep 436 - Transforming Lambda Captures

<details>
<summary>
Transforming lambda capture variables as part of the capture
</summary>

[Transforming Lambda Captures](https://youtu.be/t6hFPKiOS-Q?si=C64xS3yot4gQvlpN)

in this example, we capture a variable and transform it as part of the capture/initialization, and we do this on a parameter pack. we expand the parameter into the closing scope. we usually use <cpp>std::move</cpp> or <cpp>std::forward</cpp>, but we can also do other stuff, like static cast.

```cpp
#include <string>

template<typename ... StringLike>
void work_with_string_like_things(const StringLike & ... stringLike)
{
  auto l = [&]() {
    ((std::cout << stringLike << '\n'), ...);
  };
  l(); // this works

  auto l2 = [&]() {
    ((std::cout << stringLike.size() << ' ' << stringLike << '\n'), ...);
  };
  l2();  // this doesn't work, const char* doesn't have size()

  auto l3 = [...transformedStringLike = std::string_view{stringLike}...]() {
    ((std::cout << transformedStringLike.size() << ' ' << transformedStringLike << '\n'), ...);
  };
  l3(); // this also works!
}

int main() 
{
  work_with_string_like_things("Hello World",std::string{"Jason was here"}, std::string_view{"Doing some C++ weekly stuff"});
}
```

</details>

## C++ Weekly - Ep 437 - Pointers To Overloaded Functions

<details>
<summary>
passing an overload function pointer to a generic function and other issues
</summary>

[Pointers To Overloaded Functions](https://youtu.be/NMWv2vQQjXE?si=0WMDWW5FHc2To7AV)

we can get the address of a function with the `&` operator, but when we try getting the address of an overloaded function, there is a problem to infer the correct one.

```cpp
void use_callable(auto);

void callable(int);
void callable(double);

int main()
{
  use_callable(&callable); // this works when there is one callable function, but not when we add the other
  use_callable(static_cast<void (*)(double)>>(&callable)); // we force this to pass
}
```

we can use static casts to force the behavior to pass, but it will interact in weird ways when `auto` is used as the function parameter. there is a different option with wrapping a lambda.

```cpp
int main()
{
  use_callable([](auto v) {return callable(v);});
}
```

</details>

## C++ Weekly - Ep 438 - C++23's ranged-for Fixes

<details>
<summary>
Object Life time problem fix for the temporary objects in the range initializer.
</summary>

[C++23's ranged-for Fixes](https://youtu.be/G6FTtZCtFXU?si=hEFyj00IaOV8eAcM)

there was an object life time issue when using ranged for loops and temporary objects and references.

```cpp
struct Thing {
  std::vector<int> data {1,2,3};

  const auto & get_data() {
    return data; 
  }
};

Thing make_thing() {return Thing{};}
int main()
{
  for (const auto &val: make_thing().get_data())
  {
    std::cout << val << ' ';
  }
  std::cout << '\n'
}
```

the temporary object of Thing has it's lifetime end, so either nothing is printed or we get an error. C++23 changed the lifetime rules to extended it until the end of the for-loop.
</details>

## C++ Weekly - Ep 439 - `mutable` (And Why To Avoid It)

<details>
<summary>
Thread safety doesn't work with `mutable` members.
</summary>

[`mutable` (And Why To Avoid It)](https://youtu.be/CagZYOdxYcA?si=9GymM9419RdLZYfv). references ["You don't know const and mutable"](https://web.archive.org/web/20130105015101/http://channel9.msdn.com/posts/C-and-Beyond-2012-Herb-Sutter-You-dont-know-blank-and-blank) talk by Herb Sutter.

a value that is declared mutable can be changed by `const` functions.

```cpp
struct S {
  int x;
  int y;
  mutable int cache = 0;
  mutable bool cache_set = false;

  int get_stuff() const {
    return x+y;
  }

  int change_stuff() const {
    if (!cache_set)
    {
      cache = x+y;
      cache_set = true;
    }
    return cache;
  }
};

int main()
{
  const S s{3,4};
  std::jthread t1(&::do_stuff, s);
  std::jthread t2(&::do_stuff, s);
}
```

we can trigger a race condition by running multiple threads accessing the same object, and we can add the thread sanitizer with the `-fsanitize=thread` flag.

</details>

## C++ Weekly - Ep 440 - Revisiting Visitors for std::visit

<details>
<summary>
Different ways of using the visitor pattern.
</summary>

[Revisiting Visitors for std::visit](https://youtu.be/et1fjd8X1ho?si=yIy8P2MrM29v_eJA)

we can use a generic lambda, a concrete object which defines the necessary overloads, or use <cpp>variadic template</cpp> inheritance and <cpp>variadic using</cpp> operations, and class argument deduction.

```cpp
#include <variant>
#include <print>
#include <string_view>

std::variant<int, float, std::string_view> get_variant();

template<typename ... Callable>
struct visitor2 : Callable
{ 
  using Callable::operator()....;
};
int main()
{
  const auto value = get_variant();
  std::visit([](const auto &param) {std::println("{}", parm);}, value); // generic lambda

  struct Visitor {
    void operator(float param) const {
      std::println("Float {}", param);
    }
    void operator(int param) const {
      std::println("Int {}", param);
    }
    void operator(std::string_view param) const {
      std::println("String_view {}", param);
    }
  };
  std::visit(Visitor{}, value); // visitor object
  
  std::visit(visitor2{
    [](int param){std::println("Int: {}", param);},
    [](float param){std::println("Float: {}", param);},
    [](std::string_view param){std::println("String_view: {}", param);}}
    , value); // variadic lambda
}
```

</details>

## C++ Weekly - Ep 441 - What is Multiple Dispatch (and how does it apply to C++?)

<details>
<summary>
Playing with double dispatch
</summary>

[What is Multiple Dispatch (and how does it apply to C++?)](https://youtu.be/l2VemFmfkG4?si=-m87G3p8EwkvS9wB)

also known as multi-method, multi-method dispatch. runtime dynamic type polymorphism on one or more of the arguments. there is the classic example from a fictional game about spaceships.

we have an example of a single dispatch (normal virtual polymorphism).

```cpp
#include <variant>
#include <print>
#include <string_view>

struct SpaceObject {
  constexpr virtual ~SpaceObject() = default;
  [[nodiscard]] virtual constexpr std::string_view get_name() const noexcept = 0;
  int x;
  int y;
};

struct Craft : SpaceObject {
  [[nodiscard]] virtual constexpr std::string_view get_name() override const noexcept {
    return "Craft";
    }
};

struct Asteroid : SpaceObject {
  [[nodiscard]] virtual constexpr std::string_view get_name() override const noexcept {
    return "Asteroid";
    }
};

std::unique_ptr<SpaceObject> factor();

std::variant<int, float, std::string_view> get_variant();

template<typename ... Callable>
struct visitor : Callable
{ 
  using Callable::operator()....;
};

int main()
{
  auto spaceObject1 = factory();
  std::println("object is {}", spaceObject1->get_name()); // single dispatch
  const auto value = get_variant();
  
  std::visit(visitor{
    [](int param){std::println("Int: {}", param);},
    [](float param){std::println("Float: {}", param);},
    [](std::string_view param){std::println("String_view: {}", param);}}
    , value); // variadic lambda
}
```

for multiple (or double) dispatch, we add some stuff, and the manual way would take forever and explode

```cpp

void collide(const Craft &, const Craft &) {std::puts("C/C");}
void collide(const Asteroid &, const Asteroid &) {std::puts("A/A");}
void collide(const Craft &, const Asteroid &) {std::puts("C/A");}
void collide(const Asteroid &, const Craft &) {std::puts("A/A");}

std::vector<std::unique_ptr<SpaceObject>> get_objects();
void process_collisions(const SpaceObject& obj)
{
  for (const auto &other : get_objects()) {
    if (other->x == obj.x && other->y == obj.y)
    {
      //collide(obj, *other); // this doesn't work
      // we need to map everything out
      const auto *obj_craft = dynamic_cast<const Craft*>(&obj);
      const auto *obj_asteroid = dynamic_cast<Asteroid Craft*>(&obj);
      const auto *other_craft = dynamic_cast<const Craft*>(other.get());
      const auto *other_asteroid = dynamic_cast<Asteroid Craft*>(other.get());

      // this is tedious and error-prone
      if (obj_craft && other_craft) { collide(obj_craft, other_craft);}
      if (obj_asteroid && other_asteroid) { collide(obj_asteroid, other_asteroid);}
      if (obj_craft && other_asteroid) { collide(obj_craft, other_asteroid);}
      if (obj_asteroid && other_craft) { collide(obj_asteroid, other_craft);}
    }
  }
}
```

if we work with <cpp>std::variant</cpp>, thing can be simpler

```cpp
std::vector<std::variant<Craft,Asteroid>> get_objects();
void process_collisions(const std::variant<Craft, Astroid>& obj)
{
  for (const auto &other : get_objects()) {
    if (other.x == obj.x && other.y == obj.y)
    {
      // if we forget one of the options, we get an error!
      std::visit(visitor{
        [](const Craft &, const Craft &) {;},
        [](const Asteroid &, const Asteroid &) {;},
        [](const Asteroid &, const Craft &) {;}
        [](const Craft &, const Asteroid &) {;}
      }, obj, other);
      // we still don't know what is the 
    }
  }
}
```

the better way is to do this, allow the compiler to do the pattern matching (not true pattern matching yet), this will only work when all derived types are known in compile time.

```cpp
void process_collisions(const std::variant<Craft, Astroid>& obj)
{
  for (const auto &other : get_objects()) {
      std::visit([] (const auto& lhs, const auto& rhs){
            if (lhs.x == rhs.x && lhs.y == rhs.y) 
            {
              collide(lhs, rhs);
            }
          }, obj, other);
    }
}
```

</details>

## C++ Weekly - Ep 442 - Stop Using .h For C++ Header Files!

<details>
<summary>
Reasons to use .hpp files instead.
</summary>

[Stop Using .h For C++ Header Files!](https://youtu.be/mr3sOT-Delg?si=KQPUhr6pByWk6vTk)

only use ".h" files for libraries that are intended for C library, this is what the tools assume. if we have a "c" header that's actually a "c++" file, things can go haywire.

we should actually wrap all our ".h" files with `extern C++ {}` (but maybe not).

this a problem of mixed language projects, with some C libraries and other C++ libraries.

</details>

## C++ Weekly - Ep 443 - Stupid `constexpr` Tricks

<details>
<summary>
Some things that compile even if they look like they aren't supposed to.
</summary>

[Stupid `constexpr` Tricks](https://youtu.be/HNn-PmrL5X8?si=cWu6vUj3HlOd_UMy)

things that compile even if they look like they aren't supposed to. <cpp>static_assert</cpp> statements. referencing constant expression in lambdas without capturing them, lambdas are implicitly <cpp>constexpr</cpp> if they can be since C++17.

```cpp
constexpr auto make_value() { return 42;}

int main()
{
  const auto value = make_value();
  static_assert(value == 42);
  const x = 10;
  const auto result = [](){ return x + value + 20; }();
  static_assert(result == 72);
}
```

</details>

## C++ Weekly - Ep 444 - GCC's Implicit constexpr

<details>
<summary>
A flag to make inline functions constexpr if possible.
</summary>

[GCC's Implicit `constexpr`](https://youtu.be/CHc39_qCgMU?si=b690Hueha66gIw--)

a GCC flag `-fimplicit-constexpr`, which tries to make inline functions <cpp>const</cpp> if possible, so the example below won't compile normally, but if we enable the flag, the function is considered <cpp>constexpr</cpp> and the code will compile.

```cpp
inline auto make_value() { return 42;} 

int main()
{
  const auto value = make_value();
  static_assert(value == 42);
}
```

</details>

## C++ Weekly - Ep 445 - C++11's thread_local

<details>
<summary>
a 'static' variable that is private to the current thread.
</summary>

[C++11's `thread_local`](https://youtu.be/q9_vljSaBDg?si=7RS9-rraccr1HByx)

`thread_local` are kind of like `static` variables, but are different for each thread. rather than have one instance for all the program, it has an instance for each of the threads. if we combine `static thread_local` together, it's still the same as `thread_local`. it is only destroyed when thread execution ends, not when it's destroyed (joined)

```cpp
#include <fmt.format.h>
#include <thread>

int func()
{
  thread_local int count = 0;
  ++count;
  auto thread_id = std::this_thread::get_id();
  fmt::print("Value: {} from thread\n", count, thread_id);
  return count;
}

int main()
{
  std::jthread t1(&func);
  func()
  return func();
}
```

</details>

## C++ Weekly - Ep 446 - ImHex: An Awesome Hex Editor

<details>
<summary>
Short look at an Hex Editor
</summary>

[ImHex: An Awesome Hex Editor](https://youtu.be/zJZhyTvL8UM?si=dFDLJ8dq4oZAVyFU)

a full featured hex editor, it can parse files and identify known patterns, it can try to identify what data might be depending on types, AES Decrypter, disassembler and all sorts of stuff.

</details>

## C++ Weekly - Ep 447 - What Are Reference Qualified Members?

<details>
<summary>
Functions overloads based on the reference qualified members.
</summary>

[What Are Reference Qualified Members?](https://youtu.be/5ELoDcqqyX4?si=2b1NwvfXZHUO73gl)

this is needed before we can talk about C++23 <cpp>std::forward_like</cpp>. we know functions are const qualified (this is how we tell apart const and non-const member functions).\
but function members are also reference qualified, we can add the lvalue `&` or the rvalue `&&` reference qualifies. the reference qualified and the non-qualified overloads can't exists together.

```cpp
#include <cstdio>
#include <string>
#include <utility>
struct S {
private:
  std::string value;

public:
  std::string & get_value() {
    std::puts("get_value()");
    return value;
  }

  const std::string & get_value() const {
    std::puts("get_value() const");
    return value;
  }

  std::string & get_value_qualified() & {
    std::puts("get_value_qualified() &");
    return value;
  }

  std::string & get_value_qualified() && {
    std::puts("get_value_qualified() &&");
    return value;
  }

};

void call(const S &obj)
{
  obj.get_value();
}

S get_obj() {
  return {};
}

int main()
{
  const S s0;
  s0.get_value();

  S s1;
  s1.get_value();
  std::as_const(s1).get_value(); // const
  call(s1); // const
  s1.get_value_qualified(); // lvalue &
  std::move(s1).get_value_qualified(); // rvalue &&
  static_cast<S&&>(s1).get_value_qualified(); // rvalue &&
  get_obj().get_value_qualified(); // rvalue && // temporary object

}
```

the use case is pretty rare, but it can allow for changing implementations, which might be good for return value optimizations, or when we are worried about lifetime issues.

</details>

## C++ Weekly - Ep 448 - C++23's `forward_like`

<details>
<summary>
Forward a parameter with the same reference and const type as another value.
</summary>

[C++23's forward_like](https://youtu.be/AFcfRf5KWe0?si=SmCoUEzgtnxZ0geb)

reference qualified members, in the previous video, we had different "getters" depending on if the object is lvalue (exists in memory) or temporary.

```cpp
#include <cstdio>
#include <string>
#include <utility>
struct S {
private:
  std::string value;

public:
  std::string & get_value() & {
    std::puts("get_value() &");
    return value;
  }

  const std::string & get_value() & {
    std::puts("get_value() &&");
    return value;
  }

  std::string & get_value() const & {
    std::puts("get_value() const &");
    return value;
  }

  std::string & get_value() const && {
    std::puts("get_value() const &&");
    return value;
  }
};
```

we can use the new "deducing this" from c++23, so we call /<cpp>std::forward_like</cpp> function. the return type should be `auto &&` or `decltype(auto)`.

```cpp
struct S{
  std::string value;
public:
  auto get_value(this auto&& self) -> decltype(auto) {
    return std::forward_like<decltype(self)>(self.value);
  }
};

void call(const S &obj) { obj.get_value(); }
S get_obj() {return {}; }
int main()
{
  const S obj;
  std::cout << std::is_rvalue_reference_v<decltype(std::move(obj).get_value())> << '\n'; // 0
  std::cout << std::is_rvalue_reference_v<decltype(obj.get_value())> << '\n'; // 1
}
```

</details>

## C++ Weekly - Ep 449 - More constexpr Math!

<details>
<summary>
getting compile time math with external library
</summary>

[More constexpr Math!](https://youtu.be/reWnel5uLS4?si=YrWP2QA-BWdOafkg)

The cmath header has a lot of functions, such as fmax and the overloads from c, the template version, and some of them are constexpr (not all are).  there are also functions for linear interpolation which rely on the underlying functions in the header. however, the trigonometry function aren't constexpr.\
There is an open source project called [gcem](https://github.com/kthohr/gcem) which provides compile time math expressions, which supports even C++11. back then, we couldn't have loops, only a single expression. the gcem math expressions are all implemented with ternary expressions, so they can work even back then.

```cpp
constexpr int calculateToZeroCpp11(int input)
{
  return input == 0 ? 0 : input + calculateToZeroCpp11(input -1);
}

constexpr int calculateToZeroCpp14(int input)
{
  int res = 0;
  while (input > 0)
  {
    res += input;
    --input;
  }
  return res
}

int main()
{
  constexpr auto cpp11sum = calculateToZeroCpp11(10);
  constexpr auto cpp14sum = calculateToZeroCpp14(10);
}
```

</details>

## C++ Weekly - Ep 450 - C++ is a Functional Programming Language

<details>
<summary>
C++ always had functional principles.
</summary>

[C++ is a Functional Programming Language](https://youtu.be/jf_OxE3j4AQ?si=OzCrqNaCsoNubbY1)

C++ discourse has some focus on using C++ as a functional programming language, we have the <cpp>functional</cpp> header since about forever, including functors that are overloaded for math operator (<cpp>std::plus</cpp>), which since C++14 was also templated, we can pass them to algorithms like <cpp>std::accumulate</cpp> (which takes a binary operator). in C++11 we got lambdas to act as operators and `auto` return type and `decltype` to get the return type.

</details>

## C++ Weekly - Ep 451 - Debunking bad_alloc Memory Errors (They're actually useful!)

<details>
<summary>
Getting the bad_alloc memory error.
</summary>

[Debunking bad_alloc Memory Errors (They're actually useful!)](https://youtu.be/-f5HmDR0GGY?si=umLgq-gF5IBHOhQZ)

if we request memory blocks without writing to them, we can get a ridiculous number before getting the <cpp>std::bad_alloc</cpp> exception.
but if we try writing to the memory, we don't get the error at all, the process is terminated.
so running out of memory terminates the process, but running out of addresses leads to the exception, this can happen for memory fragmentation issues.
there are all kinds of stuff related to memory space and how many addresses are available based on the machine we use and the operating system.

it's also possible to try allocating on an external memory (GPU, shared memory), and then we get the exception and we would need to deal with it.
</details>

## C++ Weekly - Ep 452 - The Confusing Way Moves Can Be Broken in C++

<details>
<summary>
Weird things about moves and the move constructor.
</summary>

[The Confusing Way Moves Can Be Broken in C++](https://youtu.be/ZuTJAP4oMwg?si=mHZ-vdyrBLPo1NOT)

if we define a destructor for a a type, we won't get the move operators defined for us.
however, the type will still pass the check for `std::is_move_constructible_v`. because it actually checks whether we can create an object from the r-value reference, which is true, since it can call the copy constructor.

```cpp
struct S {
  ~S(){}
};

static_assert(std::is_move_constructible_v<S>);
static_assert(std::is_constructible_from<S, S &&>);

int main()
{
  S obj1;
  auto obj2 = std::move(obj1);
}
```

there's a difference between removing the move constructor and deleting it.

</details>

## C++ Weekly - Ep 453 - Valgrind + GDB Together!

<details>
<summary>
Using Valgrind inside a debugging session.
</summary>

[Valgrind + GDB Together!](https://youtu.be/pvOYwxsDIJI?si=kmGPMo7l2GYmTRhq)

here is an example of use-after-free.

```cpp
int main()
{
  int *p = new int();
  delete p;
  *p = 42;
}
```

we can run it through Valgrind and see the issue, we could also run it with the address sanitizer.
we can also use Valgrind as a remote server when debugging.

```sh
gdb ./a.out
set remote exec-file ./a.out
set sysroot /
target extended-remote | vgdb --multi -vargs -q
start
```

</details>

## C++ Weekly - Ep 454 - std::apply vs std::invoke (and how they work!)

<details>
<summary>
Trying to create our own implementations.
</summary>

[std::apply vs std::invoke (and how they work!)](https://youtu.be/DxFpQa1PyaA?si=oM2W5WH08k3e-IqT)

the two functions <cpp>std::invoke</cpp> and <cpp>std::apply</cpp> have some similarities. we can use invoke to call function with an argument list, it can be normal function, a member function, or a lambda. for the apply function, we use it on "tuple-like" arguments. it's mostly two ways to do the same thing, based on convience and preference.

```cpp
#include <functional>
#include <tuple>

struct S {
  int lhs;
  int add(int rhs);
};

int add(int lhs, int rhs);
int main()
{
  std::invoke(add, 1, 3);
  S s{42};
  std::invoke(&S::add, s, 1);

  std::tuple params{&s, 2};
  std::apply(&S::add, params);
}
```

had we wanted to write it ourselves, it would look something like this. `decltype(auto)` tells the compiler to preserve the type. while re-creating "invoke" is fairly simple, re-creating "apply" requires some template magic.\<cpp>std::make_index_sequence</cpp> - returns a <cpp>std::index_sequence</cpp><0,1,2,...,N> for the number of elements in the tuple.\
we also do pack unfolding to use the parameters as arguments for the function.

```cpp
template<typename ... Param>
auto invoke(auto Func, Param && ... param) -> decltype(auto)
{
  // only free functions and lambda, not member functions
  return Func(std::forward<Param>(param...));
}

template<typename TupleType>
auto apply(auto Func, TupleType && tuple ) -> decltype(auto)
{
  // only free functions and lambda, not member functions
  // require c++ explicit lambda template parameters
  return [&]<std::size_t... Index>(std::index_sequence<Index...>)-> decltype(auto){
    ::invoke(std::move(Func), std::get<Index>(std::forward<TupleType>(tuple))..);
  }(std::make_index_sequence<std::tuple_size_v<std::remove_cvref_t<TupleType>>>()); // immediately invoke
  
}
```

</details>

## C++ Weekly - Ep 455 - Tool Spotlight: Mutation Testing with Mull

<details>
<summary>
Using testing the mull tool for mutation testing.
</summary>

[Tool Spotlight: Mutation Testing with Mull](https://youtu.be/lhcXAnNgzlo?si=Qku9k6ERWByOX8TZ)

an open-source plugin for clang and LLVM, for debian based distros there is an easy package to get, but running it requires passing some funky arguments.

mutation test changes tests (replaces operators), and if the test still pass, there's a problem with the test and it doesn't really test what we think it does. we can use it to find corner cases.

```cpp
bool valid_age(int age) {
  if (age >= 21) {
    return true;
  }
  return false;
}

TEST_CASE("Test Valid Age")
{
  CHECK(valid_age(22));
  CHECK(valid_age(21));
  CHECK(!valid_age(10));
}
```

</details>

## C++ Weekly - Ep 456 - RVO + Trivial Types = Faster Code

<details>
<summary>
Benchmarking RVO.
</summary>

[RVO + Trivial Types = Faster Code](https://youtu.be/DzUAqXMUjtc?si=WpW-T86TKAVSmt-C)

RVO - return value optimization.

RVO usually means constructing an object on the call site directly, rather than first creating it in the function and then copying it using copy constructor. we see this a lot on branches, and we can get better RVO when we use compositions and small functions and lambdas.

for gcc, we have the `-Wnrvo` flag to detect cases where we are missing out on optimizations.
</details>

## C++ Weekly - Ep 457 - I Read C++ Magazines (So you don't have to!)

<details>
<summary>
Reading some books!
</summary>

[I Read C++ Magazines (So you don't have to!)](https://youtu.be/TnsPaEmUSL0?si=wCH2h7g61Z_wYQ0P)

looking at some C++ book from and reviewing the advice given in them. **they aren't amazing**. and the also seem to be copies of one another.

</details>

## C++ Weekly - Ep 458 - array of bool? No! constexpr `std::bitset`!

<details>
<summary>
<cpp>std::bitset</cpp> will become constant expression.
</summary>

[array of bool? No! constexpr std::bitset!](https://youtu.be/E3sfXAaR1E4?si=PCB4tlF80-HQZo18)

there is no optimization of <cpp>std::Array</cpp> for the boolean type.
This is related to how <cpp>std::vector</cpp> of bool was developed...

instead of having an optimization of an array, we already have a dedicated data structure, the <cpp>std::bitset</cpp>. it provides some extra methods (flipping, seting and reseting all the bits, counting bits).\
in the upcoming C++26 release it will also have constexpr support. which will make it even easier to use.
</details>

## C++ Weekly - Ep 459 - C++26's Saturating Math Operations

<details>
<summary>
Math Operations which protect against unwanted conversions and overflow
</summary>

[C++26's Saturating Math Operations](https://youtu.be/XNMnQOFrEIY?si=UdnHZioKWxDwe-Eb)

a feature already approved by the committee for C++26.

Saturating math protects against type changes and integer promotions, it uses some special registry tricks.

```cpp
#include <numeric>
#include <cstdint>

int main()
{
  constexpr std::uint8_t val = 140;
  static_assert(std::add_sat(val, val) == 255); // protect against overflow
  
  constexpr std::uint8_16t val2 = 33`000;
  static_assert(std::saturate_cast<std::int8_t>(val2) == 127);
  static_assert(std::saturate_cast<std::uint8_t>(val2) == 255);
  static_assert(std::saturate_cast<std::int16_t>(val2) == 32767);
}
```

</details>

## C++ Weekly - Ep 460 - Why is GCC Better Than Clang?

<details>
<summary>
Understanding differences in performance between compilers.
</summary>

[Why is GCC Better Than Clang?](https://youtu.be/4P32EFClwuo?si=l_SlSXXkm99Q-5lt)

in episode 435, we saw a 68 times performance increase by using the GPU, we also saw the GCC had consistently better output than Clang. the code is for conway's game of life, it's a single file with less than 300 LOC. and for nearly all cases, GCC was faster. there are some other findings about which configuration is faster.

the odd case is function which checks the minimal size that we need to hold the index, rather than always using <cpp>std::size_t</cpp>. when clang uses different sized indexes, it produces very different code.
</details>

## C++ Weekly - Ep 461 - C++26's `std::views::concat`

<details>
<summary>
View on combined ranges.
</summary>

[C++26's `std::views::concat`](https://youtu.be/QeWdhHyBBv0?si=DWB3xGTHVCpfCQ-c), [cpp-reference](https://en.cppreference.com/w/cpp/ranges/concat_view).

a view over multiple ranges, chained together in the desired order.

```cpp
int main()
{
  std::array a1 {1,2,3,4};
  std::array a2 {5,6,7,8,9};
  std::vector v3 {10,11,12};

  for (const auto &val : std::views::concat(a2,v3,a1))
  {
    std::cout<< val << '\n';
  }
}
```

the iterators maintain the properties of the underlying container iterators, so if we have random access iterators for the vectors, we could modify them through the combined view.

```cpp
void update(std::vector<int> &v1, std::vector<int> &v2)
{
  std::views::concat(v1,v2)[10] = 10;
}
```

we can still concatenate views over ranges of things which don't support random access, but then we'll get compiler errors.
</details>

## C++ Weekly - Ep 462 - C++23's Amazing New Range Formatters

<details>
<summary>
Standard format specializations.
</summary>

[C++23's Amazing New Range Formatters](https://youtu.be/G6hhZGUE9S4?si=ReGrpeVkdSqd6C6p)

standard format specializations were added to the C++23 specifications, and there finally is support for them! we can also use formatting specifiers for better control (remove parentheses, print as key-pair), even at the internal level of nested containers.

```cpp
int main()
{
  std::println("pair: {}", std::pair{42, 4.2}); // (42, 4.2)
  std::println("pair: {:n}", std::pair{42, 4.2}); // 42, 4.2
  std::println("pair: {:m}", std::pair{42, 4.2}); // 42: 4.2
  std::println("tuple: {}", std::tuple{1,2,3, "hello world"}); // (1, 2, 3, "hello world")
  std::println("vector: {}", std::vector{1,2,3}); // [1, 2, 3])
  std::println("array: {}", std::array{1,2,3,4} ); // [1, 2, 3, 4])
  std::println("range of array: {}", std::array{1,2,3,4} | std::views::drop(2)); // [3, 4])
  std::vector<std::pair<char, int>> vp{{'1',1},{'2',2}};
  std::println("vector of pairs: {}", vp); // [('1', 1), ('2', 2)]
  std::println("vector of pairs: {::m}", vp); // ['1': 1, '2': 2]
  
  std::map<std::string,
    std::pair<int, 
      std::map<std::chrono::time_point<std::chrono::system_clock>, 
        std::vector<int>>>> data; // very nested container
  data["Hello"] = {1, {{std::chrono::system_clock::now(), {1,2,3,4}}}};
  std::println("very nested: {}", data); // {"Hello": (1, {2024-11-04 19:47:33:641445: [1, 2, 3, 4]})}
}
```

in the future, we might be able to pair this with c++26 reflection and get something that looks like json serialization.
</details>

## C++ Weekly - Ep 463 - C++26's Safer Returns

<details>
<summary>
Making C++ Safer to use.
</summary>

[C++26's Safer Returns](https://youtu.be/T4g92jtGkXM?si=G2JCBcF3Xsh4WcQf)

this example is undefined behavior, we try to return a reference to a temporay value.

```cpp
int get_value();
const int & call_get_value()
{
  return get_value();
}
```

in C++23, this is a compiler warning, and we need a flag to make it into an error. In C++26, this will be an error. no matter what flags we pass.

</details>

## C++ Weekly - Ep 464 - Easily Printing `std::variant`

<details>
<summary>
Printing variant using a lambda
</summary>

[Easily Printing `std::variant`](https://youtu.be/-oSuITjrzgU?si=7-i-9uP7BNZE-RHa)

we can't print something with a runtime only known type (variant) through a compile known print function. we need to use a visitor somehow, but it becomes a bit silly looking.

```cpp
std::variant<int, double, std::string> get_value();

int main()
{
  std::println("{}", get_value()); // fails

  std::visit([](const auto &value){
    std::print("{}", value);
  }, get_value()); // this does work
}
```

this works because under the hood, the visitor creates different versions of the callable generic thing with the different types of the variant
</details>

## C++ Weekly Ep 465 - C++26's `std::span Over` initializer_list

<details>
<summary>
New option in C++26.
</summary>

[C++26's `std::span` Over initializer_list](https://youtu.be/hWw_P6FUN_E?si=KHn4BN21tSAUZc8L)

passing <cpp>std::initializer_list</cpp> to something which expects a <cpp>std::span</cpp> fails in C++23, but should work in C++26.

```cpp
void use_data(std::span<const int> data);

int main()
{
  use_data(std::array{1,2,3,4}); // works
  use_data({1,2,3,4}); // fails in C++23, works in C++26.
}
```

however, the generated code isn't the same, which is weird.

</details>

## C++ Weekly Ep 466 - `invoke_r` Should Not Exist

<details>
<summary>
Weird function that maybe shouldn't exist.
</summary>

[invoke_r Should Not Exist](https://youtu.be/GyAZg1LfPZo?si=3w_NJxwON_-ApKHN), [cppreference](https://en.cppreference.com/w/cpp/utility/functional/invoke)

calling (invoking) something callable with parameters, like a lambda, function object, a free function or even a member function.

```cpp
int get_value(int input)
{
  return input + 23;
}
std::string_view get_sv()
{
  return "Hello World";
}

int main()
{
  std::cout << std::invoke(get_value, 10) << '\n';
  std::cout << std::invoke_r<float>(get_value, 10) << '\n'; // implicit conversion
  std::cout << std::invoke_r<std::string>(get_sv) << '\n'; // doesn't work
  std::cout << std::string(std::invoke(get_sv)) << '\n'; // explicit conversion
}
```

<cpp>std::invoke_r</cpp> exists to modify the return type, it is implicitly converted, if that conversion exists (must be explicit), this is weird, since we want to avoid implicit conversions in our code! it even messes with our compiler warning, since the compilers will never show us warnings from the standard library, unless we specify the `-Wsystem-headers` flag to see those warning. this depends on the compiler (gcc and clang exclude, msvc doesn't).

</details>

## C++ Weekly Ep 467 - enum struct vs enum class

<details>
<summary>
two ways to do the same thing
</summary>

[enum struct vs enum class](https://youtu.be/1CjVTCiY4fc?si=b-3V9TR3lCD1nCQm)

strongly typed scoped enumerations. they are functionally the same, but one can argue that since all the members of the enum are public, then it's a struct.

```cpp
enum struct Settings {
  True = 1,
  False = 0
};

enum class OtherSettings {
  True = 1,
  False = 0
};

int main()
{
  return static_cast<int>(Settings::True);
}
```

</details>

## C++ Weekly - Ep 468 - Use `-fvisibility=hidden`!

<details>
<summary>
hiding internal functions from library consumers.
</summary>

[Use -fvisibility=hidden!](https://youtu.be/vtz8S10hGuc?si=6q90XXUzxWQFiuxy)

when we build with `-fPIC` (position independent code) flag, we tell the compiler we are building a shared library (dll, so). this effects how the code is built and what can get inlined.

```cpp
auto value(const int val)
{
  return val;
}
auto get_value(const int val)
{
  return value(v);
}

int main(int argc, char** argv)
{
  return value(42);
}
```

if we add `-fno-semantic-interposition`, this can restore the inlining behavior. the `-fvisibility=hidden` option makes the default behavior of functions to be only usable inside the library, so they aren't exposed to the consumers of the client. only things which are explicitly exported will be visible to outside consumers. there is some way to set it with CMAKE.
</details>

## C++ Weekly - Ep 469 - How to Print in C++23

<details>
<summary>
the right way to print output in C++23
</summary>

[How to Print in C++23](https://youtu.be/s6CZmNOCoQU?si=cSC4ErGZbnEv938E)

```cpp
#include <cstdio>
#include <fmt/core.h>
#include <print>

int main()
{
  std::puts("hello world");
  // if we can use external libraries
  fmt::print("Hello {}, {}\n", "Jason", 42);
  // if we can't use external libraries
  std::println("Hello {}, {}", "Jason", 42);
  // last resort
  std::cout << "Hello " << "Jason, " << 42 << '\n';
}
```

</details>

## C++ Weekly - Ep 470 - `requires` Clause vs `requires` Expression?

<details>
<summary>
When to use which one?
</summary>

[requires Clause vs requires Expression?](https://youtu.be/k1I4ZoB50_I?si=f9auEtOOdxqGj7o8)

<cpp>requires</cpp> expressions check if a code is valid.

```cpp
int main()
{
  return requires (int i) {++i;};
  //return requires (const int i) {++i;}; // fails to compile
}
```

it should be inside a template. the code isn't executed, just checking the validity.

```cpp
template<typename Type>
bool test_type() {
  return requires (const Type i) { ++i; };
}

template<typename T>
void myFunc(T t) requires false; // uncallble code

int main()
{
  myFunc(42); // function doesn't exist
  return test_type<int>();
}
```

this is used for concepts, like <cpp>std::integral</cpp>.

requires clause expect a boolean constant value, a requires expression (`requires requires`) checks if an expression can compile, it's usually is a code smell, and should be a concept by itself.

</details>

## C++ Weekly - Ep 471 - C++26's `=delete` With a Reason

<details>
<summary>
Adding a reason in the error.
</summary>

[C++26's `=delete` With a Reason](https://youtu.be/9iFJRCHwSok?si=TVoh7X4xIfSShI_F)

when we delete a function we make it a compilation error to call it, in C++26 we can add a text to explain the reason for the deletion.

```cpp
void myfunc(float);
void myfunc(int) = delete("only float are allowed, don't using implicit conversions");

int main()
{
  myFunc(4.0f); // ok
  myFunc(4); // will fail compilation
  myFunc(4.6); // will fail compilation, ambiguous call
}
```

make it hard to use the library in the wrong way.

</details>

## C++ Weekly - Ep 472 - C++23's static lambdas?!

<details>
<summary>
Using lambdas without an object.
</summary>

[C++23's static lambdas?!](https://youtu.be/M_AUMiSbAwQ?si=pzCxnIaDpTvnjYTz)

making a call and index operator static.

we could do this in the past:

```cpp
struct Callable
{
  static auto operator()(int val) {return val + 42;}
};
int main()
{
  auto x = Callable(42); // won't work
  auto x2 = Callable::operator()(42); // will work
}
```

but we can also create static lambdas directly, maybe this should have the default behavior, since it would allow the compiler to skip creating the object.

```cpp
int main()
{
  auto l = [](int val) static {return val + 42;};
  return l(10);
}
```

(playing with the code and compiler options to show this in action).

</details>

## C++ Weekly - Ep 473 - `continue` and `break` Statements

<details>
<summary>
Flow Control statements.
</summary>

[continue and break Statements](https://youtu.be/CiB1ex0hi3w?si=NqJX-V_dmb_jKley)

we can use the `continue` and `break` flow control statements, even inside a for loop.

```cpp
#include <initializer_list>

int main()
{
  for (const auto x: {2,3,4,5,6,7,8,9})
  {
    if (x==4) {continue;}
    std::cout << x << '\n';
  }
}
```

the two should be considered a code-smell, and if we catch ourselves writing them, we should consider if there's an alternative.
</details>

## C++ Weekly - Ep 474 - What Can LEA Do For You?

<details>
<summary>
LEA - load effective address instruction.
</summary>

[What Can LEA Do For You?](https://youtu.be/Ac7cdsIbjE0?si=zdPxInixN_cunFb7)

LEA instruction - load effective address, comparing the instruction and how it operates in X86 and ARM architectures.
the compiler tries to use a few instructions as possible, and avoiding multiplication instructions.
</details>

## C++ Weekly - Ep 475 - Lambdas On The Heap?

<details>
<summary>
constructing lambdas on the heap - how and why.
</summary>

[Lambdas On The Heap?](https://youtu.be/W4G2xJX9Gnw?si=Hs2GMULTt2us8Tsz)

```cpp
int main()
{
  auto l = new decltype([]{return 42;});
  return (*l)();
}
```

this might be relevant for a lambda with a large capture, we can't create a lambda directly because of the capture, but we can create a lambda and then move it to the heap.

</details>

## C++ Weekly - Ep 476 - C++23's Fold Algorithms

<details>
<summary>
Fold operations
</summary>

[C++23's Fold Algorithms](https://youtu.be/s8UInkalVgs?si=06N-57W5zcabrRHy)

fold algorithms under the ranges library, and some other under the numeric header

- <cpp>ranges::fold_left</cpp>
- <cpp>ranges::fold_left_first</cpp>
- <cpp>ranges::fold_right</cpp>
- <cpp>ranges::fold_right_last</cpp>
- <cpp>ranges::fold_left_with_iter</cpp>
- <cpp>ranges::fold_left_first_with_iter</cpp>
- <cpp>accumulate</cpp>
- <cpp>reduce</cpp> - out of order, not determistic

for example, for input of `[1,2,3,4]` with `op`:

- left fold -> `op(op(op(op(init,1), 2), 3), 4)`
- right fold -> `op(1, op(2, op(3, op(4, init))))`

if the order of operations matters (like division), then this matters.

```cpp
#include <vector>
#include <print>
#include <numeric>
#include <functional>

int main() {

  std::vector data {1.1, 2.2, 3.3, 4.4, 5.5, 6.6};
  std::print("{}\n", std::accumulate(data.begin(), data.end(), 0)); // bug!
  std::print("{}\n", std::accumulate(data.begin(), data.end(), 0.0)); // use this!
}
```

we expect the answer to be 23.1, but we get 21 instead. this is because we used 0 as the initial value, and it got converted into integer operations.\
we could also have used the <cpp>fold_left_first</cpp> operation. which actually returns an <cpp>std::optional</cpp> object.

```cpp
int main() {
  std::vector data {1.1, 2.2, 3.3, 4.4, 5.5, 6.6};
  std::print("{}\n", std::ranges::fold_left_first(data, std::plus<>{}).value()); // UB if there was no data.
}
```

</details>

## C++ Weekly - Ep 477 - What is Empty Base Optimization? (EBO)?

<details>
<summary>
Optimizing Away empty base classes.
</summary>

[What is Empty Base Optimization? (EBO)?](https://youtu.be/KU-Ahb86izo?si=tat4rxrHe1qXHnCl)

objects must have unique address, even if they don't have data in them.

```cpp
struct E {};
struct B {
  int i;
};
struct C : B {
  int j;
}
struct S : E {
  int k;
}
int main()
{
  std::print("{}\n", sizeof(E)); // 1 - must have non-zero size to have a unique address
  std::print("{}\n", sizeof(B)); // 4 - size of int
  std::print("{}\n", sizeof(C)); // 8 - base + size of int
  std::print("{}\n", sizeof(S)); // 4 - only size of int
}
```

an empty base optimization removes generated non-zero sizes, base classes are collapsed down if they are empty in order to optimize object layouts, so if our class inherits from several empty base classes, those classes might be removed entirely.

</details>

## C++ Weekly - Ep 478 - Lambdas on the Heap (2)

<details>
<summary>
Following on a previous video.
</summary>

[Lambdas on the Heap (2)](https://youtu.be/qMGi_tdKrrk?si=J0kfvY6gpNYaDzD4)

the previous video had some mistakes - "can we move a lambda"?

```cpp
auto make_lambda() {
  struct S {
    S(S &&) = delete;
    S(const S &) = delete;
    S() = default;
    int x = 42;
  };
  auto lambda = [s=S{}]() {

  };
  //return lambda; // can't be moved or copied

 auto *lambda2 = new auto([s = S{}]() {
    return s.x;
  });
  (*lambda)(); // call lambda
  std::unique_ptr<std::remove_cvref_t<decltype(*lambda2)>> ptr(lambda2); // create a unique pointer.
  return ptr;
}
```

we should use return value optimization anyway, rather than create a local  object and rely on NRVO to optimize it. we use `new auto` to create something, but in our case we need some extra fodder.

</details>
