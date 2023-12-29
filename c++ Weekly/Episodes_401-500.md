<!--
// cSpell:ignore codecov cppcoro dogbolt decompiler
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

once we get it to a stable state, we can start modifying the code and use best practice. we can use the standard library and proper types (<cpp>std::string</cpp> rather than <cpp>char *</cpp> pointers, <cpp>bool</cpp> rather than <cpp>int</cpp>). we can move to using references instead of pointers, and make sure to use <cpp>const</cpp> when needed. templates existed back then, so we can use them instead of void pointers. it's ok to discover bugs, we just need to have tests that monitor them.\
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
