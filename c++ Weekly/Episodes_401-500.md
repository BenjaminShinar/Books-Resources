<!--
// cSpell:ignore codecov cppcoro dogbolt decompiler Lippincott 
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
- <cpp>std::expected</cpp> - an alternate error handling mechanism. like an optional pair. the downside is  that it interferes with return value optimization and would require explicit move semantics.
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
