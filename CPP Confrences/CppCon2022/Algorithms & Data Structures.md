<!--
ignore these words in spell check for this file
// cSpell:ignore Berne Nonius eastl absl actionize cmov Ahnentafel eytzinger cmpgt popcount ized
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Algorithms & Data Structures

### The Imperatives Must Go! [Functional Programming in Modern C++] - Victor Ciura

<details>
<summary>
Functional programming in C++.
</summary>

[The Imperatives Must Go! [Functional Programming in Modern C++]](https://youtu.be/M5HuOZ4sgJE), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/The-Imperatives-Must-Go-Victor-Ciura-CppCon-2022.pdf)

Functional Programming has a lot of stuff going on, and it has been slowly integrated into C++ and other languages.

Humans and computers are different from one another, in terms of perspective, priorities, orientation, and many other topics. the art of programming is to translate human ideas into something that the machine can follow and execute. over the history, a formal, rigor and clear language was created, this being **mathematics**.

> What is Functional Programming?
>
> - Functional programming is a style of programming in which the basic method of
>   computation is the application of functions to arguments.
> - A functional language is one that supports and encourages the functional style.

summation in c (or Java, C#, C++) uses variable assignment as the main operation.

```c
int total = 0;
for (int i = 1; i ≤ 10; i++)
 total = total + i;
```

summation in Haskell:

```Haskell
sum[1..10]
```

Functional style is about **What** we want, while imperative style is **How** to get the results.

> "OO makes code understandable by encapsulating moving parts. FP makes code understandable by minimizing moving parts."\
> ~Micheal Feather

functional programming has a long history, basing on the ideas of **lambda calculus**, with **lisp** being the first functional programming language, and **ISWIM**, a pure functional language with no assignment, **FP** a functional language with higher-order functions, **ML** with type interface and polymorphic types. and the **Miranda** system with lazy evaluation, and eventually **Haskell**, a standardized lazy functional language. in the 90's type classes and monads were incorporated.

and while haskell didn't become a leading language, it did effect many other programming languages and ideas from functional programming are being developed into them.

- lambdas, closures, _std::function_
- value types
- constants
- comptability of algorithms
- lazy ranges
- folding
- mapping
- partial application(bind)
- higher-order function
- monads (optional, future)

some haskell code

```haskell
f[]=[]
f(x:xs) = f ys ++ [x] ++ f zs
  where
    yz = [a | a <- xs, a <= x]
    zs = [b | b <- xs, b > x]
```

which is actually quicksort

```haskell
#qsort :: Ord a -> [a] -> [a]
qsort[]=[]
qsort (x:xs) = qsort smaller ++ [x] ++ larger zs
  where
    smaller = [a | a <- xs, a <= x]
    larger = [b | b <- xs, b > x]
```

which is much simpler than even the pseudo code in other language.

> The task:\
> Read a file of text, determine the n most frequently used words, and print out a sorted list of those words along with their frequencies.

in pascal the solution was 10 pages long. while in a shell script this can be achieved in six lines

```sh
tr -cs A-za-z '\n' | # transform anything not in the A-Za-z set into new line
  tr A-Z a-z | # to lower
  sort | # sort alphabetically before unique
  uniq -c | # remove duplicates and count
  sort -rn | # sort by numbers descending
  sed ${1}q # read stream and exit
```

a lot of this is explored in the book [Category Theory for Programmers](https://github.com/hmemcpy/milewski-ctfp-pdf) or ["Functional Programming in C++: How to improve your C++ programs using functional techniques".](https://www.amazon.com/Functional-Programming-programs-functional-techniques/dp/1617293814).

#### Higher Order function

libraries for higher order function: **boost::hof** and [lift](https://github.com/rollbear/lift), both provide helper functionality for writing functional code.

```cpp
struct Employee {
 std::string name;
 unsigned number;
};

const std::string& select_name(const Employee& e) { return e.name; }
unsigned select_number(const Employee& e) { return e.number; }
std::vector<Employee> staff;

// sort employees by name
std::sort(staff.begin(), staff.end(), lift::compose(std::less<>{}, select_name);

// retire employee number 5
auto i = std::find_if(staff.begin(), staff.end(),
  lift::compose(lift::equal(5), select_number));
if (i != staff.end()) staff.erase(i);
```

Lift uses function and makes them into an overloaded set of functions, so there can be used with other higher order functions.

#### Boxes

boxes hide the value

| box                   | getting the value    |
| --------------------- | -------------------- |
| `unique_ptr<T> p;`    | `*p`, `p.get()`      |
| `shared_ptr<T> p;`    | `*p`, `p.get()`      |
| `vector<T> v;`        | `v[0]`, `*v.begin()` |
| `optional<T> o;`      | `*o`, `o.value()`    |
| `function<T(int)> f;` | `f(5)`               |

we want to perform operations on the hidden value inside the box, without breaking it. ideally, we would:

1. unwrap value from context
2. apply function (modify value)
3. re-warp modified value in context

(the value might not be in the box, an _std::optional_ can be without a value.) we apply the transformation on the box itself without unpacking it. this will be denoted as _fmap_.

```cpp
string capitalize(string str);
//...
std::optional<string> str = ...; // an operation that could fail

optional<string> cap;
if (str)
  cap = capitalize(str.value()); // capitalize (*str)
```

we want to _lift_ the "capitalize" function to make it work on boxes, from an optional to optional.

```cpp
std::optional<string> liftedCapitalize(const std::optional<string> & str)
{
  std::optional<string> result;
  if (str)
    result = capitalize(*str)
  return result;
}
```

and as a template, _fmap_ that takes a function and runs is.

```cpp
template<class A, class B>
optional<B> fmap(function<B(A)> f, const optional<A> & o)
{
  optional<B> result;
  if (o)
    result = f(*o); // wrap a <B>
  return result;
}
```

and with this, we can combine and compose functions.

```cpp
optional<string> str{" Some text "};
auto len = fmap<string, int>(&length,fmap<string, string>(&trim, str);
```

and making it more generic

```cpp
template<typename T, typename F>
auto fmap(const optional<T> & o, F f) -> decltype( f(o.value()) )
{
  if (o)
    return f(o.value());
  else
    return {}; // std::nullopt
}
```

in C++23 there are monadic extensions to _std::optional_: _.and_then()_, _.transform()_, _.or_else()_, so we compose things sequentially rather than mess around with callbacks.

#### Values

> Expression yield values, statements do not.

focusing on value semantics, but in a pragmatic way.

[C++ Weekly - Ep 322 - Top 4 Places To Never Use `const`
s](https://www.youtube.com/watch?v=dGCxMmGvocE)

> - don't `const` non-reference return types
> - don't `const` local values that need take advantage of implicit move-on-return operations (even if you have multiple different objects that might be returned)
> - don't `const` non-trivial value parameters that you might need to return directly from the function
> - don't `const` any member data
>   - it breaks implicit and explicit moves
>   - it breaks common use cases of assignment

#### C++ Ranges

> the beginning of the end of _[begin,end)_.

a. print the **even** elements of a range in a **reverse order**.
b. **Skip** the first 2 elements of the range and print only the **even** numbers of the **next** 3 in the range.
c. Modify an **unsorted range** so that it retains only the **unique** values but in **reverse order**.
d. Create a range of **strings** containing the **last 3** numbers **divisible** to 7 in the range [101, 200], in **reverse order**.

without ranges:

```cpp
std::for_each(
   std::crbegin(v), std::crend(v),
   [](auto const i) {
      if(is_even(i))
         cout << i;
   });

auto it = std::cbegin(v);
std::advance(it, 2);
auto ix = 0;
while (it != cend(v) && ix++ < 3)
{
   if (is_even(*it))
      cout << (*it);
   it++;
}

vector<int> v{ 21, 1, 3, 8, 13, 1, 5, 2 };
std::sort(std::begin(v), std::end(v));
v.erase(
   std::unique(std::begin(v), std::end(v)),
   std::end(v));
std::reverse(std::begin(v), std::end(v));

vector<std::string> v;
for (int n = 200, count = 0;
     n >= 101 && count < 3; --n)
{
   if (n % 7 == 0)
   {
      v.push_back(to_string(n));
      count++;
   }
}

```

with ranges:

```cpp
for (auto const i : v
                  | rv::reverse
                  | rv::filter(is_even))
{
   cout << i;
}

for (auto const i : v
                  | rv::drop(2)
                  | rv::take(3)
                  | rv::filter(is_even))
{
   cout << i;
}

vector<int> v{ 21, 1, 3, 8, 13,
 1, 5, 2 };
v = std::move(v) |
    ra::sort |
    ra::unique |
    ra::reverse;

auto v = rs::iota_view(101, 201)
       | rv::reverse
       | rv::filter([](auto v) { return v%7==0; })
       | rv::transform(to_string)
       | rv::take(3)
       | rs::to_vector;
```

</details>

### Functional Composable Operations with Unix-Style Pipes in C++ - Ankur Satle

<details>
<summary>
Implementing a pipe operator in C++ for functional programming.
</summary>

[Functional Composable Operations with Unix-Style Pipes in C++](https://youtu.be/L_bomNazb8M), [slides](https://github.com/sankurm/talks-and-presentations/blob/master/cppcon%202022/Functional%20Composable%20Operations%20with%20Unix-Style%20Pipes%20in%20C%2B%2B%20-%20Ankur%20Satle%20-%20CppCon%202022%20-%2012-Sep-22.pdf), [generic pipe line github](https://github.com/sankurm/generic-pipeline).

The coding world is moving toward functional programming, and unix pipelines are a way of doing just that.

uniq shell pipe lines are really simple

```sh
sort colors.txt | uniq -c | sort -r | head -3 > favColors.txt
grep "warning: " compile.log | sed 's/^.*warning: //1' | uniq | sort | uniq
git log --since=1.day | grep '^Author: ' | aws '{print $3}' | xargs  -n1 mail -s "Build failure Alert: Urgent action needed"
```

but the simple command to sort the colors is messy when using regular c++ code. as seen in the [example code at compiler explorer](https://ankursatle.godbolt.org/z/W4sMYd3hG).

with ranges, lets say we want to filter the numbers which are dividable by three and square them.

before ranges:

```cpp
std::vector<int> input = {0,1,2,3,4,5,6,7,8,9,10};
std::vector<int>, intermediate, output;

std::copy_if(input.begin(), input.end(), std::back_inserter(intermediate), [](const int i) {return i % 3 == 0;});

std::transform(intermediate.begin(), intermediate.end(), std::back_inserter(output), [](const int i) {return i * i;});
```

in c++20, with the ranges syntax, we can get a much clearer code. they are also efficient because of lazy evaluation.

```cpp
std::vector<int> input = {0,1,2,3,4,5,6,7,8,9,10};
auto divisible_by_three= [](const int i) {return i % 3 == 0;};
auto square = [](const int i) {return i * i;};

auto output = input
  | std::views::filter(divisible_by_three)
  | std::views::transform(square);
```

#### Kafka Queue Example

can we use the same principles even without ranges? extend the ideas and apply them for cases which aren't using ranges. for example, an Kafka/MQ queue.

the usual steps are:

1. read enviornment variables
2. open configuration file
3. pare configuration file
4. create kafka consumer
5. connect
6. subscribe to topic

at the consumer initialization, we need to check many things, and if any of them fails, we throw an error.

we would like a code that looks like this:

```cpp
kafka_consumer init_kafka() {
  return get_env("kafka-config-filename")
  | get_file_contents
  | pares_kafka_config
  | create_kafka_consumer
  | connect
  | subscribe;
}
```

we can actually implement this in c++11. we create an overload for the `|` operator

```cpp
using Callable = std::string(std::string&&);
auto operator|(std::string&& val, Callable fn) -> std::string {
    return fn(std::move(val));
}

// usage
auto contents = get_env("kafka-config-filename")
  | get_file_contents;
if (contents.empty()) { throw file_error{}; };

```

with this, we can start working on the code

the callable doesn't have to be function, it can be a lambda, functor, or any invokable, so we modify the code and make it generic.

```cpp
template<typename T, typename Callable>
auto operator|(T&& val, Callable&& fn) -> typename std::result_of<Callable(T)>::type {
    return std::forward<Callable>(fn)(std::forward<T>(val));
}

// usage
auto config = get_env("kafka-config-filename")
  | get_file_contents;
  | parse_kafka_config;
if (!config)) { throw json_error{}; };

```

points of notice:

- type conversion cannot take place when invoking templates.
- The type T and the passed parameter of the Callable **MUST** match.

to chain a member function, we can either write a new function, a wrapper or a lambda, or we can use _std::mem_fn_ and pass it the member function, this will work if the return types are matching.

but for the current implementation, we aren't checking the returned values at each step, and we don't want each step to assume what the previous step was. each step in the pipeline should check itself and raise the error for itself. **throw where the error occurs!**

he details an issue about having the operator in the global namespace. there is also an issue for templated functions.\
another issue is that some function should take more than one parameter. for this we use lambda captures, and in c++14, lambda initializations. we could also use _std::bind_.

```cpp
kafka_consumer init_kafka() {
  auto cert = get_certificate(); //c++11
  return get_env("kafka-config-filename")
    | get_file_contents
    | parse_kafka_config<xml>
    | create_kafka_consumer
    //| connect
    // std::bind option
    //| std::bind(connect, std::placeholders::_1, get_certificate())
    // c++11 code
    | [&cert](kafka_consumer&& consumer) { return connect(std::move(consumer), cert); }
    //This needs C++14 :(
    //| [cert = get_certificate()](kafka_consumer&& consumer) { return connect(std::move(consumer), cert); }
    | [](kafka_consumer&& consumer) {
        if (!consumer.subscribe()) { throw connect_error{}; }
        return consumer;
    };
}
```

we can also work with higher-order functions.

#### Functional programming:

principles of functional programming:

- Operations are pure functions
  - depend only on arguments
  - have no side-effects, just return or throw
- Work with values
  - work with objects with ownership
- Return a new object create or the original object passed
  - `buffer<byte> serialize(Message)`
  - `T add_timestamp(T)`
  - enable composeAbility - keep the chain going

possible errors?

what about memory leaks? this is solved when we use smart pointer (std::unique_ptr) to pass ownership to and from steps.

we can now use pipes everywhere:

```
rx_msg | parse | validate | extract | store
rx_msg | parse | validate | extract | enrich | send
rx_http_msg | extract_body | validate | actionize | make_responses | send
trigger | make_msg | encode<json> | encrypt(certificate) | send
```

#### _std::optional_

in c++17, we got the _std::optional_ type, which is a value type that may hold an object.

```cpp
std::optional<kafka_consumer> create_kafka_consumer(kafka_config&& config) {
  return std::make_optional<kafka_consumer>(std::move(config));
}
```

with this, we don't need to use smart pointers, and we can check the value at each step. we can return std::nullopt to signify a "bad" result.

so we need to modify our pipe operator. and also to deal with const ref std::optional

```cpp
template<typename T, typename Callable>
auto operator|(std::optional<T>&& opt, Callable&& fn)->typename std::invoke_result_t<Callable, T>{
  return opt
    ? std::invoke(std::forward<Callable>(fn),*std::move(opt))
    : std::nullopt;
}

template<typename T, typename Callable>
auto operator|(const std::optional<T>& opt, Callable&& fn)->typename std::invoke_result_t<Callable, T> {
  return opt
    ? std::invoke(std::forward<Callable>(fn),*opt)
    : std::nullopt;
}
```

#### Refinement in C++20

in c++20 we can constrain the pipe operator with concepts. this will give use helpful error messages.

```cpp
template<typename T, typename Callable>
requires std::invocable<Callable, T>
auto operator|(T&& val, Callable&& fn)->typename std::invoke_result_t<Callable, T> {
  return std::invoke(std::forward<Callable>(fn),std::forward<T>(val));
}
```

but what about ranges? does the overload interfere with the range pipeline?

```cpp
std::vector<int> numbers {1,3,5,7,9};
auto pipeline = numbers | std::views::transform([](int n){return n*n});
for (auto n : pipeline){std::cout << n <<'\n';}
```

it does. we need better constraints, we state the value isn't a from the ranges library.

```cpp
template<typename T, typename Callable>
requires (not std::ranges::range<T> && std::invocable<Callable, T>)
auto operator|(T&& val, Callable&& fn)-> typename std::invoke_result_t<Callable, T> {
  return std::invoke(std::forward<Callable>(fn),std::forward<T>(val));
}
```

#### Stepping into the future in C++23

c++23 has monadic operations on std::optional.

- `.and_then()`
- `.transform()`
- `.or_else()`

this makes the code similar to what we wrote earlier.

but our code so far doesn't require the return type to optional itself, so we need another constraint.

`and_then` is implemented this way.

```cpp
template<typename T>
concept basic_optional = requires(T t) {
  typename T::value_type;
  std::convertible_to<T, bool>;
  std::same_as< std::remove_cvref<decltype(*t)>,typename T::value_type>;
  std::constructible_from<T, std::nullopt_t>;
};
static_assert(basic_optional<std::optional<int>>);

template<typename T, typename Callable>
requires (not std::ranges::range<T> && std::invocable<Callable, T> && basic_optional<typename std::invoke_result_t<Callable,T>>)
auto operator|(T&& val, Callable&& fn)-> typename std::invoke_result_t<Callable, T> {
  return std::invoke(std::forward<Callable>(fn),std::forward<T>(val));
}
```

with te optional and transform versions, we no longer need to return std::optional from our functions, as it is wrapped but the pipe.

to implement or_else we can hack the pipe operator to create an overload.

</details>

### Understanding Allocator Impact on Runtime Performance in C++ - Parsa Amini

<details>
<summary>
Showing how performance is measured for allocators.
</summary>

[Understanding Allocator Impact on Runtime Performance in C++](https://youtu.be/Ctfbs6UVJ9Y), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/cppcon-UAIoRP.pdf)

talking about an upcoming book “C++ Allocators for the Working Programmer” by Joshua Berne and John Lakos.

the talk will be about:

- Local memory allocator performance
- simulating allocator allocations
- existing tools

**not** about allocator design, embedded design, alternate memory types, concurrency, instrumented allocator or alternate global allocators.

we think about how to measure the effect of allocators on performance

normal allocators: `new`/`delete` (`malloc`,`free`) and their internal implementations. special purpose (local) allocators (`std::pmr`) and the special purpose allocator resources.

special-purpose allocators give us performance boost:

> - Faster allocation and deallocations
>   - Fewer calls to new/delete
>   - Simpler allocation algorithm
>   - Monotonic incrementing
>   - Delay freeing until the end of algorithm
>   - No-op deallocation
> - Better locality
>   - Better spatial locality
>     - Access fewer pages (4KiB or 2MiB)
>     - Access fewer cache lines (64 bytes)
>     - Invalidate fewer cache lines
>     - Fewer page faults
>     - e.g., less diffusion
>   - Better temporal locality
>     - Invalidate fewer cache lines (again)
>     - Fewer page faults (again)

diffusion, where are our allocations, how close they are to each other. in this example, the internal strings can be stored next to one another (compact) or fat away (diffused).

```cpp
std::vector<std::string> strCol;
strCol.reserve(4);

for (std::size_t i = 0; i < 4; ++i) {
  strCol.emplace_back("some large string");
}

for (std::size_t i = 0; i < 4; ++i) {
  std::cout << (void*)(&(strCol[i].data()[0])) << '\n';
}
```

this problem becomes greater the longer a program runs, we get more heap fragmentation. to overcome this, we can use have the global allocator do a better job, but as programmers, we can use local allocators.

#### Allocator Performance Analysis

metrics: execution time, page faults (from the OS), cache statistics, and others. there is a two parts talk from 2017 ([part 1](https://youtu.be/ko6uyw0C8r0), [part 2](https://youtu.be/fN7nVzbRiEk)) John Lakos.

> Experiment Requirements
>
> - Allocator implementations
> - Programs that showcases allocator usage
> - A benchmarking framework to measure those programs performance
> - Interface with hardware performance counter
> - Machine(s)

they chose to create a benchmarking framework rather use an existing one (such a Google Benchmark, nanoBench, Nonius). not all performance counters are supported, and getting those directly requires cpu access privileges. system calls have overheads.

there are some existing access performance counters. but the way they act, what they can do, where and when, it all varies depending on the machine running it and the mode.

the first case study is counting unique characters in a string (std::string_view).

> 1. Use a std::set
> 1. Use a std::pmr::set, monotonic_buffer_resource
> 1. Use a std::pmr::set, monotonic_buffer_resource with a local buffer

```cpp
// baseline
std::size_t countUniqueChars1(const std::string_view& s)
{
  std::set<char> uniq;
  uniq.insert(s.begin(), s.end());
  return uniq.size();
}

// special purpose memory allocator
std::size_t countUniqueChars2(const std::string_view& s)
{
  std::pmr::monotonic_buffer_resource mr;
  std::pmr::set<char> uniq(&mr);
  uniq.insert(s.begin(), s.end());
  return uniq.size();
}

// storing the data on the stack
std::size_t countUniqueChars3(const std::string_view& s)
{
  std::array<std::byte, 10'240> buffer;
  std::pmr::monotonic_buffer_resource mr(
  buffer.data(), buffer.size(), std::pmr::null_memory_resource());
  std::pmr::set<char> uniq(&mr);
  uniq.insert(s.begin(), s.end());
  return uniq.size();
}
```

as expected, the performance gains increase as the input size increases, with small input, the overhead is larger than the gains, but soon after the specialized versions become better. using a local memory resource for this example is always better than the baseline.

#### Simulate Allocation Diffusion: Heap littering

the first case is down from a clean slate of memory, but for long running programs, the conditions are different, and we need to simulate the diffusion.

> 1. Instrument new and delete
>    - Run our algorithm some number of times with delete disabled
>    - Keep allocations in a static list
> 2. Randomly free a configurable fraction of the actual allocations
>    - Emphasis: Random
> 3. Start benchmarking as normal
>    - The Global heap is now fragmented.

this of course effects the performance, and we need to separate the data from the littering stage and the actual check. we start and stop the performance profiler
via Intra Process Communication.

the results indicate that the littering effected the global allocator, but not the special purpose allocators.

</details>

### Refresher on Containers, Algorithms and Performance in C++ - Vladimir Vishnevskii

<details>
<summary>
Understanding performance, not just big O notation.
</summary>

[Refresher on Containers, Algorithms and Performance in C++](https://youtu.be/F4n3ModsWHI),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Refresher-on-Containers-Algorithms-and-Performance.pdf).

> Motivation and agenda
>
> - Performance of containers can be a crucial component of an application performance
> - Performance is exciting topic as it is where multiple theoretical disciplines meet practice
> - The purpose of the talk is to revisit the basic factors defining efficiency of C++ containers and algorithms and elaborate recommendations on effective usage considering individual characteristics of containers.

**time complexity and big O notation**. how the number of operation increase as the input grows, we have a known function of growth, but it ignores the constants and the smaller term, so it's not a perfect comparison for real time comparison.

it does help when we compare the same operations, and we want to say that one way is faster than another, we say that linear is better than quadratic, because we use the same operations, and just compare number of them.\
this is true for comparing traversal, sorting, and searching algorithms. so for different data structures, we can say how some operations are comparable to a different data structure.

#### Traversal

bench marking iteration through _std::vector_ and _std::list_, both are linear, but list is much slower. this is because of the memory access. a vector has contigious data store, so better cache locality, less cache misses, as opposed to the many indirections of the link list which can be at different places in memory, and we aren't expecting them to be next to one another.

code comparison, both have four instructions, but not all instructions were born equal, accessing memory is expensive, when moving to the next element in the vector, we move the pointer 4 places, but for the list, we need the memory access to find the next element.

```mips
# std::vector
.L23:
  add r12d, DWORD PTR [rax]
  add rax, 4
  cmp rcx, rax
  jne .L23

# std::list
L8:
  add r12d, DWORD PTR [rax+16]
  mov rax, QWORD PTR [rax]
  cmp rax, rbp
  jne .L8
```

#### Insertion

next we benchmark insertion at reverse order, we know that worst case for std::vector is O(n^2), and for a list it's O(1) for each inserted element.

```cpp
std::list<uint32_t>l_container;
for (auto & it: date_to_insert){
  l_container.insert(l_container.begin(), it);
}
```

if we look at the graph from a distance, then using vector has a worse performance as expected, but zooming in, for the first ~1000 of elements, it actually outperforms the std::list.\
this shows how the input size matters.

if we change the way we insert data, then now vector outperforms the list for all sizes:

```cpp
std::vector<uint32_t> v_container;
for (auto const& it : data_to_insert) {
 v_container.push_back(it);
}
std::reverse(v_container.begin(), v_container.end());
```

> - std::vector::push_back doesn’t lead to copying/moving of existent items if no reallocation is required
> - std::reverse performs n/2 swaps

but trying add parallelism degrades performance.

```cpp
std::vector<uint32_t> container;
for (auto const& it : data_to_insert) {
  container.push_back(it);
}

std::reverse(std::execution::par, container.begin(), container.end());
```

> - Not all operations can be effectively parallelized
> - Multi-core parallelism improves performance if parallel operations are CPU intensive and access to shared resources is minimized
> - Overhead associated with parallelism orchestration can exceed benefits of parallel execution
> - Memory access to adjacent locations mapped to the same cache line from multiple cores can lead to issues such as “false sharing"

if we use large (non moveable) objects, then the operations std::vector performs are much more costly (copying when resizing the vector, swapping to reverse the order), and it again performs worse than std::list.

> - Operations with higher algorithmic complexity can outperform seemingly faster operations for particular data set size.
> - If use case allows, population of the container according to more performant pattern with subsequent transformation (`std::vector::push_back` + `std::reverse` in the example) can provide significant speedup.
> - Parallel algorithms should be applied with caution as they can increase execution time. They can’t be considered as simple drop-in replacements and their applicability should be evaluated.
>
> Factors of performance. Summary
>
> - Time complexity of algorithms for data organization and processing
> - Memory data access patterns (cash efficiency for systems where it is relevant)
> - Generated code complexity
> - Memory allocation patterns
> - Nature of stored elements (cheap copy/movable, static footprint)
> - Potential for parallelization
>
> Effective design should consider the individual container properties and usage scenarios to find proper application patterns.

#### Sorted Sequence

`std::lower_bound` and `std::upper_bound` can find a window of data in a sorted range.

```cpp
std::vector<uint32_t> sorted_vector{/* data*/};
auto const it_from = std::lower_bound(sorted_vector.begin(), sorted_vector.end(), from);
auto const it_to = std::upper_bound(sorted_vector.begin(), sorted_vector.end(), to);
auto const result = std::accumulate(it_from, it_to, 0u);
```

comparing inserting into std::set and maintaining the sorted order of a std::vector

```cpp
std::set<uint32_t> s_container;
for (auto const& it : random_data) {
 s_container.insert(it);
}
std::vector<uint32_t> v_container; // will contain unique sorted values
for (auto const& it : random_data) {
 auto const position
 = std::lower_bound(v_container.begin(), v_container.end(), it);
 if (position == v_container.end() || *v_container != it)
 v_container.insert(position, it);
}
```

as before, the std::set outperforms the vector for large input, but there is a number of elements for which the vector works better.

we can try to preallocate the vector to avoid reallocation and copies. we can first store everything and then sort once, and even sort with parallel.\
benchmarking shows that sorting afterwards really helps, and makes it better than using std::set.

however, we can also try to improve the std::set, such as custom allocation. if we take the EASTL (Electronic Arts standard template library) which has "fixed_set" version which is pre allocated. this is better than the normal std::set, but doesn't outperforms the post-processing std::vector (but the items are always in order).

the boost::container::flat_set data structure (and c++23) uses a continues data storage and allows random access.

#### Searching

if we compare _std::set_, _boost::flat_set_ and _eastl::fixed_set_, we again see that for small inputs, one data structure can be better than for large input.

#### Deleting

if we can, rather than delete the elements each time, we can hold them for a while and remove them in a batching process eventually. there is a slight overhead and the logic gets complicated, but eventually, there is performance gain.

#### Unsorted container

if we don't care about the ordering, then we can use _std::unordered_set_ and get better performance. typically O(1) for insertion, deletion and search.

#### Container Combination

we start with an example on N records, each record has a label. the labels are not unique, and we now that the number of possible labels is much smaller than n. we wish to count the number of records with a certain label within a specified range.

The trivial approach uses `std::count_if`

```cpp
struct record {
  std::string label;
  size_t value;
};
std::vector<record> container{/*data*/};
auto const result = std::count_if(
container.begin(), container.end(), [](auto const& it) {
return it.label >= range_from && it.label < range_to;
});
```

if we want to avoid indirection the string label, we can use a fixed memory size to avoid the indirection, we pay more in memory for each record, but we get inlining. this requires us to know the maximal size of the label.

```cpp
using label_string = eastl::fixed_string<char, 40, false>;

struct record_with_fixed_string {
  label_string label;
  size_t value;
};
std::vector<record_with_fixed_string> container{/*data*/};
auto const result = std::count_if(
container.begin(), container.end(), [](auto const& it) {
return it.label >= range_from && it.label < range_to;
});
```

the performance is better, we managed to "inline" the string.

the next thing is to precompute the label into an id, and then store the id instead of the string. we also operate on a different container.

```cpp
struct record_with_label_id {
  uint32_t label_id;
  size_t value;
};

std::unordered_map<std::string, uint32_t> label_id_mapping;
//absl::flat_hash_map<std::string, uint32_t> label_id_mapping;

auto const id_from = label_id_mapping.at(range_from);
auto const id_to = label_id_mapping.at(range_to);
auto const result = std::count_if(container_label_id.begin(), container_label_id.end(),
 [id_from, id_to](const auto& it) {
 return it.label_id >= id_from && it.label_id < id_to;
 })
```

this way we drop the string comparison and have a faster comparison.

we can presort the values and then use distance between the bounds.

```cpp
boost::container::flat_multimap<std::string, record> label_record_map;
// using label_string = eastl::fixed_string<char, 40, false>;
// boost::container::flat_multimap<label_string, record> label_record_map;
auto const it_from = label_record_map.lower_bound(range_from);
auto const it_to = label_record_map.upper_bound(range_to);
auto const result = std::distance(it_from, it_to);
```

we can combine all the techniques above and get the best performance so far.

> Container combination. Summary
>
> - Minimization of memory indirection by flattening data structures can demonstrate substantial speedup
> - Reduction of complex types into simpler ones by mapping reduces code complexity leading to reduction of execution time
> - Using precomputed data structures tailored for specific access pattern allows to reduce algorithmic complexity or minimize number of required operations
> - For insert/search/delete operations unordered hash-based containers can be preferable, but their performance can vary depending on data set size

#### Indirection

we want to avoid indirection as much as possible, even if we have some more memory to extract the key, as long as we can reduce the memory indirection, we can get better performance.

it's better to store a vector of a `pair <key, pointer to <key, value>>` and check the key in the top level than store a vector of pointers and dereference them to check the key. the redundancy in memory costs gives us better performance

#### Summary

> - Although some reasoning about performance can be done based on knowledge about existing containers, only benchmarking and profiling can validate the hypothesis about a performance for particular settings.
> - Apart from standard (STL) containers, third party alternatives can provide drop-in replacements often exhibiting better performance.
> - Combination of containers can complement functionality and mitigate downsides.
> - Separation of data preparation and data access can allow to pick best suitable patterns and containers.
> - C++17 parallel algorithms should be considered as they can provide speed up, but their contribution should be evaluated.

</details>

### Optimizing A String Class for Computer Graphics in Cpp - Zander Majercik, Morgan McGuire

<details>
<summary>
A string class that's designed to work better for some cases.
</summary>

[Optimizing A String Class for Computer Graphics in Cpp](https://youtu.be/fglXeSWGVDc), [github](https://github.com/RobloxResearch/SIMDString).

there are many strings used in computer graphics, both in the UI and in the backend (networking, localization, code generation, argument binding). There are also shaders, which work with a GPU. matching resources with dynamic string names, like a hash table. the cpu and the gpu might have different memory representations.

strings in graphics:

- often small, even empty
- constructed from smaller string
- constructed from const char\*
- lots of const segment strings.

so if we can have a drop-in replacement for std::string that works better for the common use cases, we can get better performance. so they have **SIMD string** optimized for graphics.

detecting if a string is located in the readonly segment of the memory. (avoiding creating copies whenever possible)

</details>

### Optimizing Binary Search - Sergey Slotin

<details>
<summary>
Different ways of speeding up the binary search algorithm.
</summary>

[Optimizing Binary Search](https://youtu.be/1RIPMQQRBWk), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/binary-search-cppcon.pdf), [optimized algorithms](http://github.com/sslotin/amh-code)

```cpp
// perform any O(n) pre computation with a sorted array
void prepare(int *a, int n);
// return the first element y ≥ x in O(log n)
int lower_bound(int x);
```

(not the same as _std::lower_bound_)

the basic, non optimized form which everybody stats with is simple: check the middle element,and repeat until match.

```cpp
int lower_bound(int x) {
  int l = 0, r = n - 1;
  while (l < r) {
    int m = (l + r) / 2;
    if (t[m] >= x)
      r = m;
    else
      l = m + 1;
  }
  return t[l];
}
```

when we run it in a profiler, we find that the bottleneck is getting the data elements to use in the comparison. so the problem is cache misses. but even before the the drop in performance, we see that it should be faster, so maybe the first problem to solve is the CPU pipeline, and there might be pipeline hazards which slow down the processing.

> - Structural hazard: two or more instructions need the same part of the CPU.
> - Data hazard: need to wait for operands to be computed from a previous step.
> - Control hazard: the CPU can’t tell which instructions to execute next.

if we have a miss-prediction, we get a pipeline flush, this works better than chance, but it would still be better if we didn't have it at all.

#### Branchless Programming

the outer branch prediction is fairly easy (checking the size of the remaining sub section of the data), but the inner comparison is hard to predict. so if we
reduce the number of variables that depend on the result of the comparison, we can remove branches.\
following some iterations of rewriting the code, we get this code.

- the length of the next iteration is always correct
- the start of the search section uses a different instruction.

```cpp
int lower_bound(int x) {
  int* base = t, len = n;
  while (len > 1) {
    int half = len / 2;
    base += (base[half - 1] < x) * half; // will be replaced with a "cmov"
    len -= half;
  }
  return *base;
}
```

this version works better for smaller arrays, but worse when we get into larger data sizes (4 Million). we would still like to get the benefits from branch prediction and prefetching, so we use some special instincts to explicitly pre-fetch the data that we wil use in the next iteration.

```cpp
int lower_bound(int x) {
  int* base = t, len = n;
  while (len > 1) {
    int half = len / 2;
    len -= half;
    int next_half = len/2; // BENJAMINS's addition to make the code easier to understand
    __builtin_prefetch(&base[next_half - 1]); // middle of the left half
    __builtin_prefetch(&base[half + next_half - 1]); // middle of the right half
    base += (base[half - 1] < x) * half;
  }
  return *base;
}
```

compared to the complete branchless approach, we are only slightly worse in the small arrays, and we are still better than the normal version for a while (until 16 million elements).

#### Cache Locality

> - Temporal locality: only okay for the first dozen or so requests.
> - Spatial locality: only okay for the last 4-5 requests (Memory is fetched in 64 bytes blocks called cache lines).

so we want to store the data in a different way, which is better suited. for this we look at a 16th century genealogist Ahnentafel who worked on ancestry tables.

> - the subject is listed as number 1.
> - the father of person k is listed as 2k.
> - the mother of a person k is listed as 2k+1.

this is also used in heaps, segment trees and other binary tree structures. so if we start with all the numbers between 0 and 9.\
`[0,1,2,3,4,5,6,7,8,9]`\
we store them as:\
`[6,3,7,1,5,8,9,0,2,4`]\
or in tree form:\
`[6]`\
`[3,7]`\
`[1,5][8,9]`\
`[0,2],[4]`

```plantuml
@startsalt
{
    {T
     + 6
     ++ 3
     +++ 1
     ++++ 0
     ++++ 2
     +++ 5
     ++++ 4
     ++ 7
     +++ 8
     +++ 9
    }
}
@endsalt
```

we build this permutation before starting, this is an O(n) time, and we need one more element because we work with 1-based indexing.

```cpp
int a[n], t[n + 1]; // the original sorted array and the eytzinger array we build
// ^ we need one element more because of one-based indexing
// Building
void eytzinger(int k = 1) {
  static int i = 0;
  if (k <= n) {
    eytzinger(2 * k);
    t[k] = a[i++];
    eytzinger(2 * k + 1);
  }
}
```

searching doesn't work as we want it to, as it doesn't point to the lower bound, but it can be used to get the index of the lower bound.

```cpp
// Searching:
int k = 1;
while (k <= n)
  k = 2 * k + (t[k] < x);
```

we use some binary operations to get the lower bounds, it's either we shift left once and set the last bit as 0 or 1 (effectively, $k*2$ or $k*2 +1$). and we then cancel the turn out at the end with a built-in intrinsic (ffs: find first set bit)

```cpp
int lower_bound(int x) {
  int k = 1;
  while (k <= n)
    k = 2 * k + (t[k] < x);
  k >>= __builtin_ffs(~k);
  return t[k];
}
```

this verion doesn't do so good by itself, and the number of iterations is no longer constant. but if we combine it with pre-fetching (which allows us to exploit cache locality).

```cpp
alignas(64) int t[n + 1];
// ^ allocate the array on the beginning of a cache line

int lower_bound(int x) {
  int k = 1;
  while (k <= n){
    __builtin_prefetch(&t[k * 16]);
    k = 2 * k + (t[k] < x);
  }
  k >>= __builtin_ffs(~k);
  return t[k];
}
```

now we get much better performance, as we can prefetch the data we need ahead of time (4 iterations before we use it). so we outperform the normal algorithm even for large arrays. the downside is that since we traded bandwidth (the prefetch) to reduce latency, it won't work well in bandwidth-constrained environments (multi-threaded).

#### B-Trees

to get around this issue, we can switch from binary tree to its' generalization, the B-Tree. rather than having just 2 keys in each node, we can have more, making the tree have lower height, but more comparison at each node. this allows us to minimize the number of catch line fetches.\
we set the value of B to be 16, the same as the cache line read. the tree is about 1/4 as high.

$$
\frac{\log 2 n}{\log 17 n} = \frac{\log 17}{\log 2} = \log 17 \approx 4.09
$$

we generalize the eytzinger numeration, the root is numbered zero, and node K has (B+1) child nodes numbered $\{k*(b+1)+i+1\} for i \in[0,B]$.

```cpp
const int B = 16, n_blocks = (n + B - 1) / B;
int btree[n_blocks][B];

int go(int k, int i) {
  return k * (B + 1) + i + 1; // get child node
}

void build(int k = 0) {
  static int t = 0;
  while (k < n_blocks) {
    for (int i = 0; i < B; i++) {
      build(go(k, i));
      btree[k][i] = (t < n ? a[t++] : INT_MAX);
    }
    build(go(k, B));
  }
}
```

we can find the lower bound in a simple loop, or use branchless programming to remove prediction (and do more comparison).

```cpp
// compute the "local" lower bound in a node
int rank(int x, int *node) {
  for (int i = 0; i < B; i++)
    if (node[i] >= x)
      return i;
  return B;
}

// branchless, more comparison
int rank_branchless(int x, int *node) {
  int mask = (1 << B);
  for (int i = 0; i < B; i++)
    mask |= (btree[k][i] >= x) << i;
  return __builtin_ffs(mask) - 1;
}
```

#### SIMD - Single Instruction, Multiple Data

we use SIMD instructions to compress the branchless version into less comparisons.

```cpp
typedef __m256i reg;
// compute a 8-bit mask corresponding to "<" elements
int cmp(reg x_vec, int* y_ptr) {
  reg y_vec = _mm256_load_si256((reg*) y_ptr); // load 8 sorted elements
  reg mask = _mm256_cmpgt_epi32(x_vec, y_vec); // compare against the key
  return _mm256_movemask_ps((__m256) mask); // extract the 8-bit mask
}

int rank(reg x_vec, int *node) {
  int mask = ~(
    cmp(x, node) +
    (cmp(x, node + 8) << 8)
  );

  return __builtin_ffs(mask) - 1; // alternative: popcount
}
```

so it's combined into this code:

```cpp
int lower_bound(int _x) {
  int k = 0, res = INT_MAX;
  reg x = _mm256_set1_epi32(_x);
  while (k < n_blocks) {
    int i = rank(x, btree[k]);
    if ( i < B) { // a local lower bound may not exist in the leaf node
      res = btree[k][i];
    }
    k = go(k, i);
  }
  return res;
}
```

> "S-Tree: Static, Succinct, Small, Speed, SIMDized"

this performs even better the eytzinger and prefetched version, but there is still room for improvements: we still have non constant tree height, and the update is usually unnecessary.

#### B+ Trees

we get around those shortfalls by employing B+ Trees, where the internal Nodes store Up to B keys and up to (B+1) pointers, and the leaves store a pointer to the the next leaf, and keys are copied across levels (so the tree is not succinct). (the tree grows upwards). this also allows us to reuse index from the lower bound, (no permutation on the data).

```cpp
int lower_bound(int _x) {
  unsigned k = 0;
  reg x = _mm256_set1_epi32(_x - 1);
  for (int h = H - 1; h > 0; h--) {
    unsigned i = rank(x, btree + offset(h) + k);
    k = k * (B + 1) + i * B;
  }
  unsigned i = rank(x, btree + k);
  return btree[k + i]
}
```

this is much faster (up to 15 times) than the basic algorithm, and there is still room to improvement. and there is still need to customize it to non-numerical data types, especially when we couldn't use SIMD vectorization.

</details>

### Back to Basics: Standard Library Containers in Cpp - Rainer Grimm

<details>
<summary>
Overview of the basic STL containers
</summary>

[Back to Basics: Standard Library Containers in Cpp](https://youtu.be/ZMUKa2kWtTk), [slides](https://www.modernescpp.org/wp-content/uploads/2022/09/STLContainers.pdf)

The STL defined containers, iterators and algorithms.

- Sequence Containers
  - std::array
  - std::vector
  - std::deque
  - std::list
  - std::forward_list
- Ordered Associative Containers
  - std::map
  - std::set
  - std::multimap
  - std::multiset
- Unordered Associative Containers
  - std::unordered_map
  - std::unordered_set
  - std::unordered_multimap
  - std::unordered_multiset

(std::string is also a container)

#### Common Interface

there is a common interface for all containers, they provide value semantics (not reference semantics).

```cpp
template<typename T, typename Allocator= std::allocator<T>>
```

when we push element into a container, the container becomes the owner of the data.
std::array has a fixed size, you cannot change the size during runtime. std::forward_list has no information about the length. std::string is similar to std::vector\<char\>, this is by design.

Constructors:

- default -`std::vector<int> first;`
- count - `std::vector<int> second(5);`
- range - `std::vector<int> third(second.begin(), second.end());`
- copy - `std::vector<int> fourth(third);`
- move - `std::vector<int> fifth(std::move(fourth));`
- sequence (std::initializer_list) - `std::vector<int> sixth{1,2,3,4,5,6};`

before c++20, we had to use the _erase-remove_ idiom, but now we have a function that handles this for us.

- `vec.clear();` - remove all elements
- c++20: `std::erase(vec, 5);` - remove **all elements** by value
- c++20: `std::erase_if(vec, [](auto i){return i >=3;});`

- `vec.empty();`
- `vec.size();`
- `vec.max_size();`

assignment and swaps, equality operators (equal `==` and not equal `!=`). sequence and ordered associate containers also have comparison operators (&lt;,&le;,&ge;,&gt;).

> - The containers must have the same type
> - Two containers are equal, if they have the same elements in the same sequence (applies to sequence containers and ordered associative containers)
> - The containers are compared lexicographically

range based for loop, if we which to modify the elements they must be taken by reference

```cpp
std::vector<int> vec{1,2,3,4,5};
for(auto v: vec) std::cout << v <<","; // 1,2,3,4,5
for(auto v&: vec) v*=2; // modify
for(auto v: vec) std::cout << v <<","; //2,4,6,8,10
```

#### Sequence containers

| Characteristic        | std::array                                        | std::::vector                                      | std::deque                                         | std::list                          | std::forward_list                                        |
| --------------------- | ------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | ---------------------------------- | -------------------------------------------------------- |
| Size                  | static                                            | dynamic                                            | dynamic                                            | dynamic                            | dynamic                                                  |
| Implementation        | static array                                      | dynamic array                                      | sequence of arrays                                 | doubly linked list                 | singly linked list                                       |
| Access                | random access                                     | random access                                      | random access                                      | forward and backward               | forward                                                  |
| Optimized for         |                                                   | end O(1)                                           | begin and end O(1)                                 | begin and end O(1)                 | begin O(1)                                               |
| Memory reservation    |                                                   | yes                                                | no                                                 | no                                 | no                                                       |
| Memory release        |                                                   | `shrink_to_fit()`                                  | `shrink_to_fit()`                                  | always                             | always                                                   |
| Iterator invalidation |                                                   | yes                                                | yes                                                | no                                 | no                                                       |
| Strength              | no memory allocation, minimal memory requirements | 95% solution                                       | insert and delete at the begin and end             | insert and delete at each position | fast insertion and deletion, minimal memory requirements |
| Weakness              | no dynamic memory allocation                      | insertion and deletion at arbitrary positions O(n) | insertion and deletion at arbitrary positions 0(n) | no random access                   | no random access                                         |

when we use a vector, we should pre-allocate memory (`.reserve()`) to avoid copying the elements when we reach the capacity. we can also call `shrink_to_fit()` to ask the compiler to reduce the reserved memory, but it's not guaranteed to happen. iterator invalidation can happen when we change the vector, this means that any existing iterators cannot be used anymore.

std::array has aggregate initialization syntax

```cpp
std::array<int,10> arr; // elements are not initialized
std::array<int,10> arr{}; // elements are default initialized
std::array<int,10> arr{1,2,3}; // remaining elements are default initialized
```

std::vector, using the brackets operator doesn't check boundaries, but the `.at()` operation does. we can insert (push element), or emplace (create in place).

std::deque is a double end queue, std::list (doubly linked list) has optimized for some pointer manipulation, std::forward_list is optimized for minimal memory requirements.

#### Ordered Associative Containers

(like a phone book)

| Ordered Associative Containers | Value Associated | More Identical Keys | Header  |
| ------------------------------ | ---------------- | ------------------- | ------- |
| std::set                       | no               | no                  | \<set\> |
| std::multiset                  | no               | yes                 | \<set\> |
| std::map                       | yes              | no                  | \<map\> |
| std::multimap                  | yes              | yes                 | \<map\> |

the implementation of the std::map is a balanced binary tree. we can't modify the key inside the map, (`std::pair<const Key, Value>`), std::map supports the index operator, but when we use it with a key that is not available, it default constructs the value (this is a pitfall), so reading from a map is mutation operation, so we can use the `.at()` operator to read without creating a value.

#### Unordered Associative Containers

c++11 addition, a digital phone book, sequence of key-value pairs, but the order is not visible. knowns as dictionaries, associative arrays or hash-tables.

| Unordered Associative Containers | Value Associated | More Identical Keys | Header            |
| -------------------------------- | ---------------- | ------------------- | ----------------- |
| std::unordered_set               | no               | no                  | \<unordered_set\> |
| std::unordered_multiset          | no               | yes                 | \<unordered_set\> |
| std::unordered_map               | yes              | no                  | \<unordered_map\> |
| std::unordered_multimap          | yes              | yes                 | \<unordered_map\> |

they extended the interface of the ordered associative containers, we need a hash function and equality function,the has function maps the key in a constant time to the index.

- collisions - different keys with the same hash value can be stored in the same bucket
- capacity - number of buckets
- load factor - average number of elements at each bucket
- rehashing - new buckets are created, if the load factor is bigger than 1(default value)

#### "Best" Containers to Use

for most cases, we can use std::array if it the size is known at compile time and is small (because the stack is smaller than the heap), and std::vector if the container size is not known or is large. the continues memory layout is usually the best for performance.

for associative containers, the unordered versions are faster, and unless we want the keys to be sorted, we should use those versions.

</details>

### Implementing Understandable World Class Hash Tables in C++ - Eduardo Madrid, Scott Bruce

<details>
<summary>
A fast hash table implementation - Type erased, dynamically configured, auto balancing hash table that is suitable for a wide range of workloads.
</summary>

[Implementing Understandable World Class Hash Tables in C++](https://youtu.be/IMnbytvHCjM), [slides](https://docs.google.com/presentation/d/1Du-0QLeslqqVrDp-m0Qm6_ATLA-eqNeDW9Db3zGnPrA/edit#slide=id.p)

> What is required for fast hash tables?
>
> 1. Locality friendly algorithm
> 2. Parallel computation
> 3. Efficient encoding of metadata

#### Hashtable briefing

two types of hash tables, closed addressing with chaining and open addressing with probing.

Chaining means that we have a linked list at each "slot", while probing means that once a slot is filled, we look for the next slot afterwards. chaining hash tables have very bad locality.

keys aren't always integers, so the equality test can be expensive, so we need a way to reduce the complexity and get faster results. we want to reduce _false positive_ rates on parallel key checks.

**Robin Hood** Hash Tables are a type of linear probed hash table. using two types of indexes

1. home index - the optimal slot for an element to map into
2. actual index - where the element actually ends at.

with those two indexes, we create a metric _Probe Sequence Length_, which is the difference between them, we use this PSL on _insertions_, evicting (swapping) elements with low PSL values (richer, closer to home) in favor of poorer values (elements which are further than the home index).\
When we do a _find_ operation, we can either find the required element, or do an early exit, if the calculated PSL of the searched element is lower (richer) than the one we compare against, because if the element would exist, then that would be it's location.\
_Delete_ operations don't use tombstones. instead, all elements from that position until an empty slot or an element it's home index are shifted left (becoming closer to home).

this Robin Hood table has good locality, with the maximum find operation complexity being O(log(k)) (size of the table) on non-adversarial hash functions. because PSL can only grow by 1 per slot (or become zero), it also helps in other aspects.

there are some other implementations of Robin Hood tables and other kinds of hash tables.

#### Design and Goals

SWAR - SIMD within a registerer, an abstraction that creates software lanes on a register that cannot interact with one another. like how SIMD has hardware lanes.

Generic Programming allows us to have zero-cost abstractions, which are also optimizer friendly and works well with composition.

if we want to write better hash tables, we need to change how we think about "costs" of operations. traditional costs model either focus on the number of comparisons or the number of cacheline fetches. counting comparisons is a good way to get moderate performance, counting cacheline fetches treats fetching from memory as the only important operation, which provides good performance at most cases, but fails when the memory hierarchy is complex, and ignores other slow-down issues.

the tiered instruction cost model assigns a value to each operation based on tiers - some operations are nearly free, while others are very expensive (division, modulo), with branch and cache operations being the most expensive (branch and cache misses, unexpected indirect jumps).

the hash table metadata can be customized, including data layout.

#### Implementation

metadata bit packing, placing the PSL and least/most significant bit of the key. we use both to quickly compare many possible positions.

they break the "Hash" into "Scatter" and "Hoist", instead of using the hash as both a way to find an index and the equality checks. with the two separated, different functions can be used to get better performance - more entropy, better spread of elements on the hashSet.

another goals is to make the code more deterministic, which would reduce the branch predictor workload. sometimes its easier to do both sides and merge them together than use the branch predictor.

(more stuff about SWAR and alignments)

insertions cause evictions, which cause insertions, this needs to be handled somehow. some ways about resizing, different ways of growing the tables.

#### Benchmarking Results

</details>

<cpp>std::unordered_map</cpp> is faster than <cpp>std::map</cpp> for finds , and the robin hood map is even better that that. for misses, the normal map is better the unordered version, but both still fall short of the robin hood variant.

showing the code on compiler explorer.

##

[Main](README.md)
