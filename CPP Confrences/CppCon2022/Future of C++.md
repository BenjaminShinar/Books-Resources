<!--
ignore these words in spell check for this file
// cSpell:ignore libav asio pckbxdata ospanstream stoi Simula coro Toolability composability __xmmcat_pckbxdata elifdef elifndef spanstream stackfull stackless monostate homoiconic
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Future of C++

### Contemporary C++ in Action - Daniela Engert

<details>
<summary>
making an video streaming application using modern c++ code.
</summary>

[Contemporary C++ in Action](https://youtu.be/yUIFdL3D0Vk), [source code](https://github.com/DanielaE/CppInAction).

people are vocal about the "facts", which is what they believe to be true regarding c++.

> - "C++ is dead!"
> - "The committee only care about nerds"
> - "The committee only care about library writers"
> - "The committee is too slow and doesn't deliver"
> - "The committee is too fast delivers half-baked junk"
> - "The committee is a bureaucratic bunch of troglodytes detached from reality"
> - "C++ is all you need, the rest is just syntactic sugar"

to respond to these "facts", she created a demo application that does stuff from scratch using non-trivial features:

- core language features
- standard library features
- 3rd-party library features
- concepts
- co-routine
- infinite ranges
- c++23 features

the goal is to make everything work together smoothly by using contemporary c++. this leads to the question of defying contemporary c++. there are different marks of "eras" in c++. the c++03, the c++11+14 Renascence, the c++17 enhancements, the c++20/23 "big bang". the answer is that all of these are "contemporary".

The specifications:

> The Server
>
> - _waits for clients to connect_ at any of a list of given endpoints.
> - when a client connects _observes a given directory_ for all the files in there, repeats endlessly.
> - filters all _GIF_ files which contain a _video_
> - *decode*s each video into individual _video frames_
> - _sends_ each frame _at the correct time_ to the client
> - sends _filler frames if_ there happen to be _no GIF files_ to process.
>
> The Client
>
> - tries to _connect to_ any of a list of given _server endpoints_
> - *receives video frames *from the network connection
> - _presents_ the video frames in a reasonable manner _in a gui window_
>
> The Application
>
> - performs a _clean shutdown_ from all inputs that the user can interact with
> - _handles timeout and errors properly_ and performs a clan shutdown

the code will use c++23 in many places, will relay on c++20, use compile time programming. It will use some C libraries (libav, SDL), so some C stuff to consider, like memory handling, and we have wrapper libraries, which create value types from these heap-allocated. there are many uses of generators, custom iterators, and co-routines. The code expects that failures and errors will happen, and it must be ready for them. tuples are used all around, and so are awaitable objects. the lifetime is clear. composition is also used, we get execution context from the asio library instead fo executors, and we can use the `stop_service` to control remote termination and shutdown. ownership of objects is transferred, and we use a watchdog at each level to abort when needed.

stuff which is used in the code

- `std::generator`
- `std::expected` (c++23)
- `std::println`
- `std::stop_source`, `std::stop_token`, `std::stop_callback`
- `std::span`
- ranges, views
- concepts
- requires expressions
- co-routines
- structured binding
- fold expressions
- compile-time decision, compile-time functions, compile-time reflection
- value types
- strong types
- compiler generated closure types
- compiler generated state-machines

and also **source code composition**, using c++20 modules. pre-compiled and cached modules, every type of module is used:

- Interface unit
- Implementation unit
- Interface partition
- Internal partition
- Private module fragment
- Header unit

</details>

### Can C++ be 10x Simpler & Safer? - Herb Sutter

<details>
<summary>
C++ Simplicity, Safety, and Toolability. “Simplifying C++” #9 of N.
</summary>

[Can C++ be 10x Simpler & Safer?](https://youtu.be/ELeZAKCN4tY), [github](https://github.com/hsutter/cppfront)

a continuation of the “Simplifying C++” series, finding the refined language hidden within the c++ language.

the metrics are safety - less CVE and bugs. simplicity - less guidance and text when teaching c++.

> What is C++?
>
> - Zero-overhead abstraction
> - Determinism and Control
> - Friction-free interop with C and previous C++.
>
> What is _not_ core C++:
>
> - Specific syntax
> - Unsafe code
> - Security exploits
> - Tedium
> - Vexing parsing
> - Obsolete features
> - Lack of good defaults
> - Difficulty writing tools
> - 1,000-page lists of guidelines
> - Sharp edges
> - (General: Not having nice things)

the reason we cannot fix everything what away is that c++ must be 100% **syntax backward compatibility**. but can we have compatibility and make things better? what if we could only "pay" for that backward compatibility if we really needed it?

> What if we could **"C++11 feels like a new language"** again, for the whole language?

this lecture is part of a series about "re-factoring" c++ into fewer, simpler, composable, general features. taking apart features into something that can standalone as proposal.\
in the past:

- Lifetime
- gc_arena
- `<=>` - spaceship operator
- Reflection and metaclass
- Value-based exceptions
- Parameter passing
- Pattern matching using `is` and `as`

in the past, C++ compilers created C code, the "C with Classes" era. So this time we are doing this again, rather than having a **CFront** compiler to create C code, we will have a **CppFront** compiler which will create c++ code from our new features/language.

goals:

> - be **consistent** - don't make similar things different, Make important differences visible.
> - be **Orthogonal** - Avoid arbitrary coupling. Let features be used freely in combination.
> - be **General** - don't restrict what is inherited. don't arbitrarily restrict a complete set of uses. Avoid special cases and partial features.

left-to-right syntax: "declare left to right", x is a thing.

now:

```cpp
class shape {/*Syntax 1 code since 1980, can't update semantics without backward compatiblity breakage concerns */};

auto f (int i) -> string {/*Syntax 1 code C++11, can't update semantics without backward compatiblity breakage concerns */}
```

but maybe we could have:

```cpp
shape: type = {/* Syntax 2 code doesn't exist today, can update semantics as desired without any breaking changes*/}
f: (i: int) -> string {/* Syntax 2 code doesn't exist today, can update semantics as desired without any breaking changes*/}
```

and in practice: this is a working code example, modules-first, strong ODR, fast-build times.

we also use in-out param passing. `_` is wildcard, implicit template, and optional `{return }` for single expression functions. we use the same syntax for lambda as we do normal functions, and we don't need to declare functions before using them.

```cpp
main: () -> int = {
  vec: std::vector<std::string> = ("Hello","2022");
  view: std::span = vec;

  for view do:(intput str:_){
  len := decorate(str);
  println(str, len);
  }
}

decorate: (inout thing: _ ) -> int = {
  thing = "[" + thing +"]";
  return thing.ssize();
}

println: (x: _, len: _) =
std::cout << ">>" << x << " - length" << len << "\n";
```

running this code through the cppFront compiler creates a cpp file, which does what we wanted. it makes use of good defaults and uses universal function calls. the code is also readable, it doesn't feel like something a machine generated, so it can still be used in a codebase. the defintions are forward declared. commas and comments maintain their position.

1. combining current style syntax with new syntax (mixed code).
2. pure cpp2 compiling.
3. using C code. allowing for tools to run.

using the proper defaults, avoid pre-processor. order independent, heap allocation is always managed using proper pointers (unique pointer as default).

just like cpp gave immediate advantages over c files, and typescript improved javascript without requiring additional changes to non-typescript files. this allows us to avoid the python 2/3 fiasco.

#### Safety

four types of safety:

- type safety
- bounds safety
- lifetime safety
- initialization safety

those are still the most common weakness, and they can be handled by the language.

in 2021, there was an executive order about cyber-security, and the office of commerce suggested to avoid using memory-unsafe languages (and it called out c and c++). each software is now a target for attacks.

for type safety, a lot of the guidelines are solved by introducing `as` and `is` into the language.

for bounds safety, we can't ban pointer arithmetic in current code, but we can reject any new cpp2 code that does. there are also issues about bound-checking containers, with compiler flags to add extra security. this done by using contracts (pre, post condition and assertions).

lifetime safety, avoiding null pointers, not having uninitialized pointers point to null.

differentiating between essential and accidental complexity. some problems are inherent to the issue, and some are just badly design.

for initialization safety, we have checks for first/last uses of a variable. composability of parameter passing.

#### Parameter passing

- in - const, uses good defaults for copy or reference
- copy - copy, move when needed
- inout - by reference, not const
- move
- forward - `auto &&` parameter with a require clause, perfect forwarding now for concrete types (not just templates)

#### Consistency

one way to declare functions, at each scope, named, un-named, lambda, range-for body. capture is done with a special symbol (string interpolation). also comes into play with post conditions, capture is done by value.

the point of cpp2 isn't to bring in weird concepts from other languages, we don't pull in terms which are foreign to how we think about our code. we only care about the simplified ideas.

#### Conclusion

support c++ evolution, be able to stop teaching so many stuff which are confusing and have exceptions.

</details>

### C++ in Constrained Environments - Bjarne Stroustrup

<details>
<summary>
understanding what C++ is and how it should be used in a low-resource environments.
</summary>

[C++ in Constrained Environments](https://youtu.be/2BuJjaGuInI)

> - some fundamentals
> - low level manipulation
> - raise the level of abstractions
> - core guidelines
> - exception and error codes
> - foundations of c++

it's easy to discuss individual features and rules in isolation, but it's also useless to think about them without considering the rest of the framework.

constrained environments - notable embedded systems

> - Limited space for code and run time support
> - Limited space for data
> - Limited processing power
> - Need to manipulate "unusual hardware"
> - No human operator
> - Extreme reliability requirements
> - Limited opportunities for hardware/software updates
> - Zero or very limited downtime
> - Strict limits on latency
> - Availability of simulators
> - Ability to run in "unusual" run-time environments or without run-0time support
> - Long service life

so the focus is:

> - Reliability - not easy to repair
> - Dependability - where failure can be disastrous
> - Resources - small relative to the tasks the handle
> - Messy Problems - the real world can be messy
> - Maintainability - "lives" for decades

the issues are closely related to one another, and they are core parts of how we look at new cod that we write use.\
we can't and shouldn't look for a single "silver bullet", there won't be anything that solves all the problems.

another note: the language is just a part of the toolbox, we have many other parts that help us, and some of them are better equipped to address some problems.

the complexity of the technological world is increasing, the demands for applications are increasing, the code is more complex, and it takes over more and more of the design.

there are some misconceptions which we need to be released from.

> - Learning C is a prerequisite for learning c++.
> - C is the best C++ sub-set for constrained systems
> - Good C++ relies on a lot of dynamic memory use and large class hierarchies with run-time resolution of all function calls.
> - To be reliable, code must be littered with run-time tests
> - Embedded systems must avoid exceptions
> - To write good code, you must use the latest features:
>   - only the latest features
>   - all the latest features
>   - also, the likely future features

the language isn't the key factor, the key factor is how we use it. it serves the programmer, not the other way around.

we use the _onion principle_, we start with a simple, safe and general interface, when needed, we give an alternative of improved, and eventually, we get to use the hardware directly (GPU, FPGA), but we shouldn't optimize every single line of code.

> "every time you peel a layer, you cry more."

we measure what important to us, and we estimate what we can't directly measure, but we need to be careful that we aren't missing the marks and measuring things which don't count.

for constrained systems, it can be done by either professional programmers, but they can lack the domain knowledge. on the other hand, engineers with extensive domain knowledge might not be the most amazing programmers.

two pillars of C++:

> 1. Direct map to hardware - even "unusual hardware" (e.g. signal processors, GPU, FPGAs).
> 2. Zero-overhead abstraction in production code - abstraction must be general and affordable.

foundations:

> 1. Static type system
> 1. Systematic and general resource management
> 1. Efficient object oriented programming
> 1. Flexible and efficient generic programming
> 1. Compile time programming
> 1. Direct use of machine and operating system rearouses

direct mapping to hardware, everything is memory, pointers are machine addresses, objects are location in memory. however, the machine model is still a model and an abstraction, when we write embedded systems, we need to know actually what is the machine we work on.

so we need to **express intent**, checking what we can check, the `std::variant` vs `union` question, a variant has a built-in check to make sure we use the right form.\
Immutable data lets us protect data, no race conditions. always initialize values. bit manipulation is sometimes necessary, but it's messy.

#### Raise the level of abstraction

abstractions - classes, templates, compile-time computation.

zero overhead principle, it's not the same as zero-cost, we use something, so we pay for it. User defined types are useful. resources - memory, handles, locks, sockets. we acquire them and then release. we have the scope-based resource management (RAII). we want to move work from run-time to compile time - compute values, remove indirections, inline function calls, copy elisions, avoid selecting virtual functions at runtime.

#### C++ Core Guidelines

> Stay out of "dark corners" and eliminate bad practice.

it's not just a collection of rules, it tries to be coherent philosophy. application code doesn't need to use all the features of the language. the guidelines tell us what to use and what to avoid. some examples (pointers again!). there is a static analyzer that matches common problems.

#### Exception and Error Codes

sometimes mistakes and errors are inevitable, different application domains have different requirements for error handling, there isn't a single solution that fits all problems and all domains.

- exceptions
- error codes
- crash/terminate

no error handling mechanism is without over-head, there is always prorogation (stack unwinding), we can get zero-over implementation for the green flow, if we design correctly.\
Error handling is occasionally the most error-prone part of the code. we want to get the code from the detection point to the handling point. when to throw exceptions and when to use error-code. avoid littering `try-catch` blocks everywhere, only when there's something to do in this level.

#### Dreams of the Future

> - The ISO standard
>   - Executors
>   - Pattern Matching
>   - Static reflection
>   - Better library support for coroutines
>   - Make the preprocessor redundant
>   - Integrated language support for static-analyzers
> - Stability/Compatibility
> - Performance

not everything has to be standardized in the language standard. a package mechanims, repositories, analyzers, support for code transformations.

</details>

### What’s New in C++23 - Sy Brand

<details>
<summary>
overview of some new C++23 features.
</summary>

[What’s New in C++23](https://youtu.be/vbHWDvY59SQ)

a short video (TikTok? reel?) with the new features:

- explicit object parameters
- _std::expected_
- monadic interface for _std::optional_
- multidimensional subscript operator
- _std::ranges::to_
- _std::generator_
- _sts::views::zip_
- constexpr for _\<cmath\>_ and _\<cstdlib\>_
- formatting ranges
- standard library modules
- _std::print_

We start with three elements, and we want to have each element duplicated and placed into an array of it's own. so from `[a,b]` we wish to have `[[a,a],[b,b]]`. we start by expanding with our own custom function and then using the _std::views::chunk_ function (from c++23) to create sub ranges. we could also use _std::views::chunk_by_ with a predicate. we then use the new range formatting and _std::print_ to print the data to the console.

```cpp
auto treats = std::vector<std::string_view>{"toy","chicken","catnip"};

auto chunked = treats
  | my::views::repeat_range_n(3) //  custom function which doesn't exists
  | std::views::chunk(3); // library range function from c++23.

std::print("{}",chunked);
```

we will want to combine the two actions, first repeat and then chunk. we could make it a closure (lambda). we compose the range adapter together, the closure doesn't take the range itself. [godbolt example](https://godbolt.org/z/EvTaoadrn). we could add the _[[always_inline]]_ attribute to force the compiler to inline it

```cpp
auto repeat_groups = [](std::size_t n)[[always_inline]] {
  return my::views::repeat_range_n(3) | std::views::chunk(3);
};
```

our next goal is to combine the elements of ranges into tuples (or pairs). we use _std::views::zip_, there is also a customizable function _std::views::zip_transform_. we can push the result into a container with _std::ranges::to_. we could use the new container _std::flat_map_ to avoid cache miss in favor of binary search. a third option is using _std::views::cartesian_product_ to create all the possibilities.

```cpp
auto treats = std::vector<std::string_view>{"toy", "chicken", "catnip"};
auto cats = std::vector<std::string_view>{"marshmallow", "milkshake", "lexical cat"};

auto map = std::views::zip(cats, treats) |
  std::ranges::to<std::unordered_map>();
auto flatMap = std::views::zip(cats, treats) |
  std::ranges::to<std::flat_map>();
auto cartesianMap = std::views::cartesian_product(cats, treats) |
  std::ranges::to<std::unordered_multimap>();
```

now in the story, we want to have the cats perform in order to get the treats. but the cats won't always play along. we could write the combination function ourself, which can get complicated and hard to read. we would want to have come composability, which we can get with the new _and_then_ function. this is the monadic interface for _std::optional_. there is also _std::optional_transform_ and _std::optional::or_else_.

```cpp
struct mouth {
  std::optional<mess> feed();
};

struct mouth {
  std::optional<mouth> shake();
};

struct cat {
  std::optional<paw> offer_food();
};

// explict function
std::optional<mess> feed_cat(cat c){
  auto maybe_paw = c.offer_food();
  if (!maybe_paw) {
    return std::nullopt;
  }

  auto maybe_mouth = c.maybe_paw();
  if (!maybe_mouth) {
    return std::nullopt;
  }

  return maybe_mouth->feed();
}

// still the same, but using initializers inside the if statement
std::optional<mess> feed_cat_2(cat c){
  if (auto maybe_paw = c.offer_food(); maybe_paw) {
    if (auto maybe_mouth = c.maybe_paw(); maybe_mouth) {
      return maybe_mouth->feed();
    }
  }
}

// using the new syntax
std::optional<mess> feed_cat_3(cat c){
  return c.offer_food()
    .and_then(&paw::shake)
    .and_then(&mouth::feed);
}
```

however, the _std::optional_ can only return a value or indicate a lack of value. if we wish to know why there is no value, we can use the new _std::expected_. (there still isn't a monadic interface for it).

```cpp
struct mouth {
  std::expected<mess,err> feed();
};

struct mouth {
  std::expected<mouth,err> shake();
};

struct cat {
  std::expected<paw,err> offer_food();
};
```

had we wanted to implement it ourselves, we would have to define overloads for each combination of const and rvalue and lvalue. this is where _explict object parameters_ can help us. we template the _this_ parameter, so now we deduce the type of the object, and we can use _std::forward_ to write the code once. this helps us with recursive lambda, CRTP, and other stuff.

```cpp

// multiple overloads
template <class T, class E>
class expected_old {
  //..
  constexpr const T& value() const&;
  constexpr T& value() &;
  constexpr const T&& value() const&&;
  constexpr T&& value() &&;
  //..
};

// explict object parameters
template <class T, class E>
class expected_new {
  //..
  template <class Self>
  constexpr auto && value(this Self&& self);
  //..
};
```

now we move on to generator (and combining them with the for-range loop). the generator can be treated as an infinite range. _std::generator_ uses coroutines. this allows us to avoid maintaining state outside of the coroutine.

```cpp
// ...
std::generator<treat> treat_generator = make_treat_generator();
for (auto treat : treat_generator | std::views::take(10))
{
  marshmallow.give_treat(treat);
}
// ...

std::generator<treat> make_treat_generator() {
  int n_treats = 50;
  while (true) {
    if (n_treats == 0)
    {
      n_treats = get_treats_from_shop(50);
    }
    n_treats--;
    co_yield treat{};
  }
}
```

in our story, there are some C functions which we use. we could wrap them in a smart pointer and pass a deleter struct. and in c++23, we could avoid creating the object by making the operator _static_. when we wish to get the object, we use _std::outptr_ to direct the output to the smart pointer and _std::to_underlying_ to safely cast the enum to the correct type. there are also _std::inout_ptr_, _std::is_scoped_enum_ type trait and _constexpr std::unique_ptr_ to allow usage in compile time context.

```cpp
void get_cat(cat** c, int color);
void delete_cat(cat* c);

struct cat_deleter {
  static void operator()(cat* c){delete_cat(c);}
};


void getCatCpp()
{
  std::unique_ptr<cat, cat_deleter> c;
  cat* raw_c;
  get_cat(&raw_c, static_cast<int>(color::calico));
  c.reset(raw_c);
}

void getCatCpp23()
{
  std::unique_ptr<cat, cat_deleter> c;
  get_cat(std::out_ptr(c), std::to_underlying(color::calico));
}
```

we have he _std::move_only_function_.

```cpp
struct cat_action {
  std::string name_;
  std::function<mess()> action_;
};

std::unique_ptr<cat> cat_p;
cat_action my_action;
// my_action.action_ = [cat_p = std::move(cat_p)]{/*...*/}; // error, doesn't compile.
```

the _std::unreachable_ which marks something as undefined behavior, and lets the optimizer work better.

```cpp
enum class color {black, tortoiseshell, ginger, calico};

cat cat_from_color (color c){
  using enum color;
  switch (c){
    case black: return lexical_cat;
    case tortoiseshell: return marshmallow;
    case ginger: return milkshake;
    default: std::unreachable();
  }
}
```

we have a none-owning multidimensional container _std::mdspan_, we can use it to wrap over a continues data in memory. and we can also use the new multidimensional bracket operator. and we can use the new literal _uz_ to denote the index as being of type unsigned size_t

```cpp
std::optional<cat> boxes[2*3*4] = {};
auto box_span = std::mdspan<std::optional<cat>,extents<2,3,4>>(boxes);

box_span[0uz,1uz,1uz]= cat_from_color(color::black);
```

_std::byteswap_ which transforms the byte order (network to host), and allows us portability over using compiler intrinsic. the _if consteval_ allows different behavior at compile time than at runtime. the _std::ospanstream_ creates an output string from an existing buffer. _std::stacktrace_ exposes the stack trace of the current frame.

```cpp
constexpr box_data pack_box_date(auto const& box_span) {
  if consteval {
    // costly constexpr-friendly algorithm
  }
  else {
    return __xmmcat_pckbxdata(box_span); //specialized compiler intrinsic that runs in runtime
  }
}

std::span<char> get_display_buffer();

void dispatch_box_data(auto const& box_span) {
  auto bd = pack_box_data(box_span);
  auto result = send_over_network(std::byteswap(bd));

  if (!result) {
    auto buf = get_display_buffer();
    std::ospanstream os(buf);
    os << std::stacktrace::current();
  }

}
```

_std::views::join_with_ finally allows us to concatenate strings together. another quality of life function is _.contains()_ on strings, which replaces calling _.find_ and comparing to _std::npos_.

```cpp
void send_commands(std::span<std::string_view> commands) {
  auto joined_commands = commands | std::views::join_with(';');
  std::string response = dispatch(joined_commands);
  if (response.contains("error")){
    //handle error
  }
}
```

c++23 has an eager ranges iota function _std::ranges::iota_ (unlike the lazy evaluated _std::ranges::views::iota_ from c++20). and _std::ranges::shift_left_ operates on each element in the range.

```cpp
void count_sheep() {
  std::vector<int> numbers_to_count(100);
  std::ranges::iota(counts,1);
  while (!numbers_to_count.empty()) {
    auto to_count = std::ranges::shift_left(numbers_to_count,1);
    count(to_count[0])'
  }
}
```

since modules were introduced, we no longer need to `#include` headers. instead, we can `import std` to get everything.

not in this talk:

- ranges:
  - _ranges::starts_with_,_ranges::ends_with_
  - _ranges::fold_
  - _ranges::slide_
  - _ranges::adjacent_, _ranges::adjacent_transform_
- _invoke_r_
- _reference_converts_from_temporary_, _reference_constructs_from_temporary_

</details>

### C++23 - What's in it for You? - Marc Gregoire

<details>
<summary>
going over the changes in C++23
</summary>

[C++23 - What's in it for You?](https://youtu.be/b0NkuoUkv0M),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Cpp23-Whats-In-It-For-You-Marc-Gregoire-CppCon-2022.pptx)

#### C++23 Core Language

##### Deducing _this_

current condition:

```cpp
class Person {
public:
   Person(std::string name) : m_name{std::move(name)} {}
   std::string& GetName() & { return m_name; }
   const std::string& GetName() const & { return m_name; }
   std::string&& GetName() && { return std::move(m_name); }
   const std::string&& GetName() const && { return std::move(m_name); }
private:
   std::string m_name;
};
```

using a template function that deduces the current const and lvalue/rvalue of the instance.

```cpp
class Person {
public:
   Person(std::string name) : m_name{std::move(name)} {}

  template <typename Self>
  auto&& GetName(this Self&& self) {
    return std::forward<Self>(self).m_name;
  }
private:
   std::string m_name;
};

```

we can also change how the reference qualifies are part of the function signature, rather than being stuck at the end, we move them up to the first `this` argument.

```cpp
struct DataBefore {
  // before
  void f() &;
  void g() const &;
  void h() &&;
}
struct Data
  // after
  void f(this Data&);
  void g(this const Data&);
  void h(this Data&&);
}
```

this also allows us to create recursive lambda expressions.

```cpp
auto fibonacci = [](this auto self, int n) {
  if (n < 2) { return n; }
  return self(n - 1) + self(n - 2);
};
```

##### if consteval

this is better than `std::is_constant_evaluated`, because it's part of the language (no header), and we can call other immediate (consteval) function.

##### multidimensional subscript operator

the square bracket `[]` operator can accept more than one parameter.

##### Attributes on Lambda Expressions

we can add attributes to lambda expressions results. before we could have attributes on the lambda object, but now we can also control what happens to the result generated by the lambda.

```c++
auto a =[]() [[deprecated]] {return 42}; //c++20
auto a =[][[nodiscard]]() [[deprecated]] {return 42}; //c++23
```

##### Literal Suffix for size_t

existing suffixes: U,L,UL, LL, ULL.\
adding UZ for unsigned std::size_t and Z for the signed corrsponding type.

```cpp
std::vector data {11,22,33};
for (auto i=0, count = data.size(); i <count ;++i)
{/*...*/} // won't compile, i is int

for (auto i=0uz, count = data.size(); i <count ;++i)
{/*...*/} // will compile, i is size_t like count.
```

##### auto(x): decay-copy in The Language

create a copy that is an rvalue, using `auto(x)` or `auto{x}`.

```cpp
void process(int&& value) {
  std::cout << value << '\n';
}

void process_all(std::vector<int>& data) {
   for (auto& i : data) { process(auto(i)); }
}

```

##### #elifdef, #elifndef, and #warning

new preprocessor directives, some shorthand for existing directives, and `#warning` is now part of the standard.

##### Marking Unreachable Code

marking part of the code as unreachable, evoking undefined behavior but getting better optimization from the compiler.

##### Assumptions

assume is now part of the language, `[[assume(expr)]]`, we can tell the compiler our assumptions and get better performance.

##### Named Universal Character Escapes

`\N` for name character escape - using standard Names of unicode characters.

```cpp
//Pre-C++23
// UTF-32 character literal {LATIN CAPITAL LETTER A WITH MACRON} (Ā)
auto a { U'\u0100' };
// UTF-8 string literal {LATIN CAPITAL LETTER A WITH MACRON}{COMBINING GRAVE ACCENT} (Ā̀)
auto b { u8"\u0100\u0300" };
//C++23
auto c { U'\N{LATIN CAPITAL LETTER A WITH MACRON}' };
auto c { u8"\N{LATIN CAPITAL LETTER A WITH MACRON}\N{COMBINING GRAVE ACCENT}" };
```

##### Trim Whitespace Before Line Splicing

anything whitespace after "\\" the end of a line is striped

#### C++23 Standard Library

##### String Formatting Improvements

`std::print` and `std::println` make using `std::format` easier to use. we can remove calls to `std::cout` (no more chevrons!) and we can stop writing `\n`.

printing ranges becomes easier.

##### Standard Library Modules

two named modules imports

- `import std` - all c++ headers and c wrappers.
- `import std.compact` - all c++ headers and bringing the c functions int the global namespace.

##### std::flat_map / std::flat_set

new containers / adapters

- std::flat_map
- std::flat_multimap
- std::flat_set
- std::flat_multiset

they provide associative retrieval, which we can define the underlying types for each container (keys and values).

##### std::mdspan

multidimensional expansion of `std::span`, non-owning wrapper to access continues data in multidimensional way. supports policies.

##### std::generator

a standard coroutine generator for coroutines.

```cpp
std::generator<int> getSequenceGenerator(int startValue, int numberOfValues) {
   for (int i { startValue }; i < startValue + numberOfValues; ++i) {
      // Yield a value to the caller, and suspend the coroutine.
      co_yield i;
   }
}
int main() {
   auto gen { getSequenceGenerator(10, 5) };
   for (const auto& value : gen) {
      std::cout << value << " (Press enter for next value)";
      std::cin.ignore();
   }
}
```

##### basic_string(\_view)::contains()

a `.contains()` method to search for a substring without comparing an iterator to an end iterator.

##### Construct string(\_view) From nullptr

prohibits creating an string from nullptr. it was undefined behavior previously, and now it's an error.

##### basic_string::resize_and_overwrite()

used when performance is critical,

bad:

```cpp
std::string GeneratePattern(const std::string& pattern, size_t count) {
   std::string result;
   result.reserve(pattern.size() * count);
   for (size_t i = 0; i < count; i++) {
      // Not optimal:
      // - Writes 'count' nulls
      // - Updates size and checks for potential resize 'count' times
      result.append(pattern);
   }
   return result;
}

```

now

```cpp
std::string GeneratePattern(const std::string& pattern, size_t count) {
   std::string result;
   const auto step = pattern.size();
   // GOOD: No initialization
   result.resize_and_overwrite(step * count, [&](char* buf, size_t n) {
      for (size_t i = 0; i < count; i++) {
         // GOOD: No bookkeeping
         memcpy(buf + i * step, pattern.data(), step);
      }
      return step * count;
   });
   return result;
}
```

##### Monadic Operations for std::optional

removing the need to check if there is an optional at each step, which was convoluted.

- `transform()`
- `and_then(foo)`
- `or_else(foo)`

```cpp
std::optional<int> Parse(const std::string& s) {
   try { return std::stoi(s); }
   catch (...) { return {}; }
}
int main() {
  while (true) {
    std::string s;
    std::getline(std::cin, s);
    auto result = Parse(s)
      .and_then ([](int value) -> std::optional<int> { return value * 2; }) // double the value
      .transform([](int value)    { return std::to_string(value); })
      .or_else  ([]               { return std::optional<std::string>{"No Integer"}; }); // if we didn't have a value
    std::cout << *result << '\n';
  }
}

```

##### Stacktrace Library

getting and working on a stacktrace. we can iterate over it if we want.

##### Changes to Ranges Library

- `starts_with()`, `ends_with()`
- `shift_left()`, `shift_right()`
- `to` - convert the range into a container.
- `find_last()`, `find_last_if()`, `find_last_if_not()` - searching for values, returns a subrange.
- `contains()`, `contains_subrange()`
- folding:
  - `fold_left()`, `fold_left_first()`
  - `fold_right()`, `fold_right_last()`
  - `fold_left_with_iter()`, `fold_left_first_with_iter()`

##### Changes to Views Library

the transform functions operate on the resulting tuple, it can reduce them into a single value.

- `zip`, `zip_transform` - make one tuple out of multiple tuples.
- `adjacent`, `adjacent_transform` - sliding window of elements.
  - `pairwise`, `pairwise_transform` - like adjacent<2>.
- `slide` - like adjacent, but a runtime parameter.
- `chunk`, `chunk_by` - batch into non-overlapping tuples.
- `join_with` - finally concatenate strings.
- `stride` - takes every nth element,
- `repeat` - repeat a value
- `cartesian_product` - match every element with every element
- `as_rvalue`

##### std::expected

a class template that can return the expected Value or an error type. like std::optional, but with extra data.

std::expected<T,E>.

- std::unexpected() - create the unexpected value
- `.has_value()`
- `.value()`
- `.error()`

##### std::move_only_function<>

a type of function that can be only moved, such as a lambda with a captured unique_ptr.

before:

```cpp
int Process(std::function<int()> f) {
   return f() * 2;
}
//This can be called as follows:
std::cout << Process([] { return 21; });   // 42
//But this fails:
//std::cout << Process([p = std::make_unique<int>(42)] { return *p; }); //error
```

now we can take both the normal functions or those with move-only parts.

```cpp
int Process(std::move_only_function<int()> f) {
   return f() * 2;
}
//This can now be called as follows:
std::cout << Process([] { return 21; });                              // 42
std::cout << Process([p = std::make_unique<int>(42)] { return *p; }); // 84
```

##### std::spanstream

use stream operators on external, existing buffers. either as input or output streams.

##### std::byteswap()

swap bytes in an integral type

##### std::to_underlying()

replaces `static_ cast<std::underlying_type_t<E>>(value)` with a simple `std::to_underlying(value)`.

##### Associative Containers Heterogeneous Erasure

avoid creating temporary key object when performing lookups. in c++ it also works with `erase` and `extract`.

#### Removed Features

##### Garbage Collection Support

removing the garbage collection support, it was added in c++11, was never used and no compiler vendor ever implemented it. so now it's gone.

</details>

### Understanding C++ coroutines by example: Generators - Pavel Novikov

<details>
<summary>
Building up a generator by ourselves.
</summary>

[Understanding C++ coroutines by example: Generators part 1](https://youtu.be/lm10Cj-HNKQ), [Understanding C++ coroutines by example: Generators part 2,](https://youtu.be/lz3F036_OvU) [slides](https://toughengineer.github.io/talks/2022/CppCon/Understanding%20C++%20coroutines%20by%20example%202-generators.pdf)

---

(part 1)

Goals:

> Develop intuition about how generators work:
>
> - coroutine generators in general
> - range generators
>   - how recursive generators work in principle
> - async generators

A function is a coroutine if it contains one of these:

- `co_return` (coroutine return statement)
- `co_await` (await expression)
- `co_yield` (yield expression)

an example of a coroutine.

```cpp
Generator<int> foo() {
  co_yield 42;
}
```

co-routines can stackfull(fiber) or **stackless**. c++ borrowed the idea fro the Simula language. when one of the predefined keywords exists in a function, it is transformed to have a promise type (initial_suspend, final_suspend, unhandled_exception).

`co_yield` -> `co_await promise.yield_value(expression)`.

`co_return` -> `{expression; promise.return_void(); goto final-suspend;}`

there is an implicit co_return at the end of coroutines, even if we don't write them. coroutines are lazy asynchronous tasks, they are suspended at the initial suspend point.

a naive generator example:

```cpp
Generator<std::string> foo() {
  co_yield "hello";
  const auto s = std::string{ "world" };
  co_yield s;
}

const auto f = foo(); // generator object
std::cout << f() << " " << f() <<'\n'; // each call corresponds to a co_yield
```

the coroutine starts at a suspended state, and when we call it, it starts running until the `co_yield`.

the inital suspends protects us from the coroutine throwing an exception when we create it. it can only happen when we run it.

```cpp
Generator<int> bar() {
  const auto values = getValues(); // may throw
  for (auto n : values)
    co_yield n;
}

{
  const auto g = bar(); // created but not used
}
```

it's safe to destroy a generator, even if it hasn't started and even if we didn't get to end of it. a coroutine can take arguments like normal functions, but it stores the parameters inside the stack frame. it's better to pass parameters by value, not by reference, to avoid issues.

```cpp
template<typename T>
struct Generator {
    struct promise_type;
    Generator(Generator &&other) noexcept;
    Generator &operator=(Generator &&other) noexcept;
    ~Generator();
    auto &operator()() const;
  private:

    explicit Generator(promise_type &promise) noexcept;
    std::coroutine_handle<promise_type> coro;
};
```

there are mandatory functions that our promise_type has to implement

```cpp
struct promise_type {
  auto get_return_object() noexcept;
  std::suspend_always initial_suspend() const noexcept;
  std::suspend_always final_suspend() const noexcept;
  std::suspend_always yield_value(const T &value) noexcept(std::is_nothrow_copy_constructible_v<T>);
  void return_void() const noexcept {}
  void unhandled_exception() noexcept(std::is_nothrow_copy_constructible_v<std::exception_ptr>);

  T& getValue(): // not mandatory

  private:
  std::variant<std::monostate,T, std::exception_ptr> result; // not mandatory
};
```

the final_suspend is std::suspend_always so we could retrieve the value, if it exists, before the promise_type object goes out of scope. the coroutine handle is a move-only object,

(a coroutine is a built-in state machine, it's synthetic sugar around it.)

futures and promises always use synchronization, while coroutines don't have to. so there is less overhead.

there is an extra copy inside `yield_value`, an object is constructed before the suspension, and is destroyed when it's resumed. we can probably get around this with function overloads and have a variant of the pointer type.

when we call a generator more times than the coroutine has `co_yield` statements, then the behavior is up to us to define.

---

what we want to have:

> required operations:
>
> - check if there are values
> - get a value
> - advance to the next value

iterators and generators

```cpp
auto it = r.begin();
if (it != r.end) { // check
  auto x = *it; // get
  ++it; // advance
}
```

use it like this:

```cpp
const auto g = bar(); //get generator
for (auto &i : g) {
  std::cout << i <<'\n';
}
```

to implement this interface, we need to have an Iterator struct, with `.begin()` and `.end()` methods.

we modify the promise_type struct, and we start working on the iterator type, we do the boiler plate iterator stuff, and we have it hold a coroutine handler to the promise_type.\
We need to change the increment operator, it needs to check if there is a coroutine, to check that it isn't done, to resume the coroutine and to throw the exception from the coroutine if it had one.\
The iterator is invalidated when the generator coroutine finishes.\
we redefine the dereference operator to access the coroutine value, it doesn't throw exceptions (noexcept).

#### Lazy Iterator

an iterator which advancing doesn't throw, but getting the value can throw, and we can check if there is an exception waiting.

#### Yielding from Nested Generators

we can have generators inside generators, so each time we resume and suspend at each level.

```cpp
Generator<int> bar(){
  const auto values = getValues();
  for (auto n: values)
    co_yield n;
}

Generator<int> baz(){
  co_yield 1;
  co_yield 2;
  co_yield 3;

  for (auto n: bar())
    co_yield n;
}
```

to avoid this overhead, we use **RecursiveGenerator** to yield the entire generator back to the caller.

```cpp
RecursiveGenerator<int> bar(){
  const auto values = getValues();
  for (auto n: values)
    co_yield n;
}

RecursiveGenerator<int> baz(){
  co_yield 1;
  co_yield 2;
  co_yield 3;

  co_yield bar(); // yield the whole thing.
}
```

for a recursive generator we need a field (root) to track the nested-ness, so we can set the root and continuation points of the generators to one another.

in c++23, we will have a std::generator class in STL. calling `.begin()` twice on the same iterator will be undefined behavior. this can effect how we write code. especially with ranges. the is probably because it was modeled after istream_iterator, but in that case it was predictable behavior.

#### Async Generator

asynchronous generators, a result that will be ready sometimes in the future. we suspend when we start the async call, and we will only yield valid values.

```cpp
Task<std::vector<int>> getValuesAsync();

AsyncGenerator<int> generateValuesAsync(){
  const auto values = co_await getValuesAsync();
  for (auto &v : values){
    if (isValueValid(v))
      co_yield v;
  }
}
```

</details>

### 10 Years of Meeting C++ - Historical Highlights and the Future of C++ - Jens Weller

<details>
<summary>
Looking at the past 10 years in terms of the C++ language and the conferences.
</summary>

[10 Years of Meeting C++ - Historical Highlights and the Future of C++](https://youtu.be/Rbu_YP2sydo),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/10-years-of-meeting-C.pdf)

past, present and future.

Ten years ago, C++11 was released, it was the second major update of the language, and it is considered the beginning of the "modern c++" era, since then, c++ was able to update every three years (c++14,c++17,c++20, and upcoming c++23).\
the C++Now2011 conference led to him starting a german local c++ group and starting the Meeting C++ conference. the Conference evolved into a platform on it's own, a community based in berlin, and is helping in the creation and advancement of user groups.

in 2014, CppCon started, stable committee of the standard, a lot of c++ content being created (blogs, videos, podcasts - cppCast). he names the era between 2013 and 2019 as "a golden age".

since 2017 and 2018, there are surveys of the conferences audiences, and there is data about c++ usage and habits.

in 2020 - everything changed. no more onsite meetings, limitations, cancellations, lockDowns. but this also boosted online meetings and online events, this became the current situation of hosting hybrid events: pre-recordings, liveStreams, some live talks. this has lots of challenges to organize, hard to get sponsors and funding. attendance is down.

the cycle begins again, for the conferences and the language. c++20 is the new c++11, c++23 is the new c++14 (bug fix and polishing touches for the "big new standard"), conferences will build back up again from the lowered attendance.

looking ahead: two fundamental issues - "running c++ in context", scheduling, parallelism, concurrency, coroutines (executors),this is a being change which needs to be handled correctly, but it is blocking development of many other features.\
The other big issue is ABI (break ABI to save c++, keep ABI to save the world). maybe this could be solved with interfaces, maybe not. there are other proposals, but nothing is known for now. there are also a lot of things we also want - pattern matching, reflection, networking, and more. and this will require maintainers.

but C++ is more than just the standard. libraries, frameworks.

</details>

### Reflection in C++ - Past, Present, and Hopeful Future - Andrei Alexandrescu

<details>
<summary>
Overview of reflection and how we would use it.
</summary>

[Reflection in C++ - Past, Present, and Hopeful Future](https://youtu.be/YXIVw6QFgAI)

#### Motivation - "Why do we care about adding reflection to C++?"

in human speak, reflection is the ability to reason about oneself and others. in programming, it's about observing code and shaping behavior. in C++ we focus on compile-time reflection and generic code.

> "Ability of a template to query details of its template parameters and shape its definition accordingly"

design by reflection has two sides to it: the reflection proper, and then insertion of code driven by that reflection.

three steps:

> 1. extract information about the code using reflection.
> 2. process information, derive and synthesize additional type information
> 3. use result to generate code.

MLOC - million lines of code. projects are becoming larger and larger, and at some point, it becomes a liability for maintainence, and the number of problems increase linearly with the number of lines.

we have tools that eliminate bugs but cost us in increased line of codes, types, constraints, checkers, linters, they all protects us from bad code, but force us to write more lines.

some of the bloat is duplicated boilerplate code, which is obvious, but there are other places which have duplicated that in interspersed with the functionality.

example: a "tainted string" - a string type that comes from outside source. our options:

- <cpp>typedef</cpp> / <cpp>using</cpp> declarations - just alternative names
- public inheritance - we should never inherit from a **value type**
- private inheritance - forwarding to a member

```cpp
struct tainted_string: private std::string {
  using std::string::empty;
  using std::string::size;
  /*  more boiler plate */
}
```

the problems:

- Constructors must be written by hand
- Iterators must be manually "copied"
- Methods that take and return <cpp>std::string</cpp> such as `append`, `insert`, `operator+=`

this type does nothing, so far, but takes hundreds of lines to write. we can do this manually and create the perfect "Strong Using" type, but we can't do this in code. this is absolute opposite of abstraction.

a second example is inheritance of implementation:

> - "I need a map/unordered_map that throws from `operator[]` if the key is not found"
> - "I need a type like <cpp>std::vector</cpp> that grows automatically in `operator[]` or `at()`"
> - "I need a url/filename type that is like <cpp>std::string</cpp> with a few tweaks"

inherit the implementation of a type, with a slight variations. this was what inheritance was supposed to solve.

there are containers adaptors, which are hardly used because of the rigid assumptions, and there are standard "container independet" methods, but they either face the problem of working against a rigid interface, or require writing a lot of code to function.

a classic example of inheritance and interface is the flying penguin.

```cpp
void Penguin::fly() override{
  throw NotImplementedException("Penguins can't fly outside an airplane.")
}
```

not implementing an interface also provides information. a way around this is to create a sub interface for each method, which results in a disjointed set of interfaces. A possible solution, which would be possible by reflection, is **PowerSet interface**. which allows us to define any number and any combination of methods in a given interface.

this will allow us to have large meaningful interfaces, without making them hard to work with. all interfaces are powerSets, and user can implement just a subset.

other possible uses:

> - Send a struct over the wire/marshal RPC arguments/bind to foreign languages.
> - Define a class just like <cpp>std::map/vector/etc</cpp>, but traces/times/counts/instruments/sends notification on calls to member functions.
> - Generate SQL CREATE TABLE/JSON/XML from a struct, or a struct from SQL SELECT.
> - Take a regex, generate C++ code for automaton (like CRTE, simpler and fast to compile).
> - Take a grammar, generate parser for it; run parser at compile-time or run-time.
> - Foreign language integration: reflection is (paradoxically) key to integration.
> - General high-leverage code (multitude of complex behaviors from terse defintion).

#### Brief History

- LISP has reflection - code is data, data is code, homoiconic.
- AOP/MOP - aspect oriented programming, focus on whole system policy rather than types and code generation.
- Macro systems: focus on code generation less than introspection.
- Meta Classes - circle, rust macros, the D language.
- Static Reflection is almost mainstream in modern c++.

there was attempts to do static reflection in the past, such as template based - which didn't scale well, inheritance based - which wasn't efficient and was hard to use.

the "sweet spot" is using value based and compile time expressions, we get to use more and more parts of the language, without using external tools.

#### P1240 Proposal

changes to primitives:

- alias to value `^T` means "reify T", yielding a constexpr value of type <cpp>meta::info</cpp>.
- convert value back to alias `[:e:]` means "splice e", yielding a template name or type.
- Reciprocity axiom: `[:^T:]` is T, and `^[:e:]` is e.
- primitives for manipulating <cpp>meta::info</cpp> values, in the form of regular consteval functions exposed to the user.

example: printing the name of the enum.

```cpp
#include <meta>
template <typename T> requires(std::ie_enum_v<T>)
constexpr std::string to_string(T value){
  template for (constexpr auto e: std::meta::members_of(^T)){
    if ([:e:] == value) {
      return std::string(std::meta::name_of(e));
    }
  }
  return "<unnamed>";
}
```

this "template for" unrolls the loop during compilation, but the if statement runs in runtime. we can implement this with a switch statement, and select which version is used based on the number of possibilities in the enum. like this example which uses loop unrolling for small enums, and a switch statement for larger enums.

```cpp
#include <meta>
template <typename T> requires(std::ie_enum_v<T>)
constexpr std::string to_string(T value) {
  if constexpr (std::meta::members_of(^T).size() <= 7>) {
    template for (constexpr auto e: std::meta::members_of(^T)) {
      if ([:e:] == value) {
        return std::string(std::meta::name_of(e));
      }
    }
  }
  else {
    switch (value) {
      template for (constexpr auto e: std::meta::members_of(^T)) {
        case [:e:] {
        return std::string(std::meta::name_of(e));
        }
      }
    }
  }
  return "<unnamed>";
}
```

the problem with the code above is that `if constexpr` introduces a scope, so we can't use the stuff inside the scope outside of it, which would effect some stuff.

operations on reified values

- add const
- instantiate template with arguments
- get the reference type
- the the pointer type
- conditional type

example of adding a "traced" class wrapper, not part of a proposal, but should be easy to create. we create all the signatures of the original class member functions, add our customizable code to each of them and forward to the original implementation. this again requires us to have a decision control flow without introducing a scope.

```cpp
template<Class T>
class traced {
private:
  T payload_;
  template for (constexpr auto e: non_private_member_functions_of(^T)) {
    [/ e /] {
      std:: cout << "Calling " << name_of(e) << '\n';
      return payload_.[:e:](std::forward<[:parameter_types_of(e):]>([:parameters_of(e):])...);
    }
  }
};
```

</details>

##

[Main](README.md)
