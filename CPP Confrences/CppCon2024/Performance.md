<!--
// cSpell:ignore uffer NRVO URVO
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Performance

<summary>
6 Talks about Performance.
</summary>

- [x] Can You RVO? Using Return Value Optimization for Performance in Bloomberg C++ Codebases - Michelle Fae D'Souza
- [x] Designing C++ code generator guardrails: A collaboration among outreach and development teams and users - Sherry Sontag, CB Bailey
- [x] Fast and small C++ - When efficiency matters - Andreas Fertig
- [ ] Limitations and Problems in std::function and Similar Constructs: Mitigations and Alternatives - Amandeep Chawla
- [x] Session Types in C++: A Programmer's Journey - Miodrag Misha Djukic
- [x] When Nanoseconds Matter: Ultrafast Trading Systems in C++ - David Gross

### When Nanoseconds Matter: UltraFast Trading Systems in C++ - David Gross

<details>
<summary>
Engineering low-level systems and trying to get good performance.
</summary>

[When Nanoseconds Matter: UltraFast Trading Systems in C++](https://youtu.be/sX2nF1fW7kI?si=baUk2_c9e6ZOifLe), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/When_Nanoseconds_Matter.pdf)

Market making, make a lot of small profits and avoid big losses. being fast enough to react and being smart enough to react correctly.

order book: "bids" and "asks", the prices in which others are willing to buy something (a stock), and the prices in which others are selling the same thing. there are hundreds of thousands of price updates per second.

#### Data Structures

our C++ data will look something like:

```cpp
enum class Side {Bid, Ask};

using OrderId = uint64_t;
using Price = uint64_t;
using Volume = uint64_t;

void AddOrder(OrderId orderId, Side side, Price price, Volume volume);
void ModifyOrder(OrderId orderId, Volume newVolume);
void DeleteOrder(OrderId orderId);
```

we definitely would want to use a hashmap, the most obvious choice would be to use <cpp>std::map</cpp> to store the orders, but benchmarking is a bit of a lie, because it has dynamic allocations. so we can make that our first principle - "No node containers".\
So what about <cpp>std::vector</cpp> as the backing container and <cpp>std::lower_bound</cpp>? the complexity is much worse than that of the map, there is a problem that when we modify the top (start) of the vector, we need to copy all the elements afterwards, which is a performance hit. we can get around this by reversing the vector and focusing our actions at the end of it, which reduces the number of copy operations. This is our second principle - "Understanding your problem (by looking at data)".\
the third principle will be "Hand tailored (specialized) algorithms are key to achieve performance".

> Running perf on a benchmark - Can't run perf on the entire binary if our initialization function is "big".

```cpp
void RunPerf()
{
  pid_t pid = fork();
  if (pid == 0)
  {
    const auto parentPid = std::to_string(getppid());
    std::cout << "Running perf on parent process " << parentPid << std::endl;
    execlp("perf", "perf", ..., parentPid.c_str(), (char*)nullptr);
    throw std::runtime_error("execlp failed");
  }
}

void InitAndRunBenchmark()
{
  InitBenchmark(); // might take a long time!
  RunPerf();
  RunBenchmark();
}
```

and then we run `perf stat –I 10000 –M Frontend_Bound,Backend_Bound,Bad_Speculation,Retiring –p pid`, with the `-M` flag being a list of metrics we are interested in, the `-I` is the interval in milliseconds for each report.

some categories (by intel)

- not stalled pipelines
  - retiring
    - base
    - ms-rom
  - bad speculation
    - branch miss-predict
    - machine clear
- stalled pipelines
  - frontend bound
    - fetch latency
    - fetch bandwidth
  - backend bound
    - core bound
    - memory bound

we can also run `perf record -g -p <pid>` to check which instructions take the most time. in our example, it's obvious that the binary search is the problem, and we get a lot of branch miss-predictions. so we try creating a branchless binary-search instead.

```cpp
template <class ForwardIt, class T, class Compare>
ForwardIt branchless_lower_bound(ForwardIt first, ForwardIt last, const T& value, Compare comp)
{
  auto length = last - first;
  while (length > 0)
  {
    auto half = length / 2;
    // multiplication (by 1) is needed for GCC to generate CMOV
    first += comp(first[half], value) * (length - half);
    length = half;
  }
  return first;
}
```

(note: we don't have early exit conditions in the code above).

we can use another tool to check hardware counters, IPCs, branch misses, cycles and instructions. this will show us that our custom code has less prediction misses and more instructions,

we next think about the memory access and hardware, and actually using liner search gives us the most uniform latency. so the fourth principle is "Simplicity is the ultimate sophistication" and the fifth principle is "Mechanical sympathy".

there are also stuff that we could try, but it's hard to measure in vacuum, such as the <cpp>[[likely]]</cpp> and <cpp>[[unlikely]]</cpp> attributes or inlining expressions using lambda functions (IIFE - immediately invoked functions expressions). we should avoid type erasure (<cpp>std::function</cpp>) as it gets in the way of compiler optimizations.

#### Transport: Networking and Concurrency

> General pattern
>
> - Kernel bypass when receiving data from the exchange (or other low-latency signals)
> - Dispatch / fan-out to processes on the same server

in the userspace networking, we can use "SolarFlare" as the industry standard and "OpenOnload" for BSD sockets. we can also have custom TCP/UDP stacks or move even lower to layer2 networking (which will require writing it ourselves). this is our sixth principle - "True efficiency is found not in the layers of complexity we add, but in the unnecessary layers we remove".

we can sometimes replace sockets with shared memory, have as much operations happen away from the kernel.

> Shared Memory\
>
> - If you don't need sockets, no need to pay for their complexity - "As fast as it gets"
> - Kernel isn't involved in any operations
> - Multi processes requires it – which is good for minimizing operational risk
>
> What works well in shared memory
>
> - Contiguous blocks of data: arrays!
> - One writer, one or multiple readers -> stay away from multiple writers

| Metric              | Concurrent Queues                    |
| ------------------- | ------------------------------------ |
| Bounded             | Yes – simpler & faster               |
| Blocking            | No – readers don't affect the writer |
| Number of Consumers | Many                                 |
| Message Size        | Variable length                      |
| Dispatch            | Fan-out                              |
| Type Support        | PODs                                 |

the seventh principle is "Choose the right tool for the right task", so lets design our queue: "FastQueue".

two counters (pointers) - read and write counters. when we write data, we first move the write counter, copy the data, and move the read counter to the same location. when there are no writes in place, the two counters are at the same place.

```cpp
struct QProducer
{
  void Write(std::span<std::byte> buffer);
};

struct QConsumer
{
  int32_t TryRead(std::span<std::byte> buffer); // returns #bytes read, 0 if nothing to read
};

// simplified code!
void QProducer::Write(std::span<std::byte> buffer)
{
  const int32_t payloadSize = sizeof(int32_t) + buffer.size(); 
  mLocalCounter += payloadSize; // advance the write counter
  mQ->mWriteCounter.store(mLocalCounter, std::memory_order_release);
  std::memcpy(mNextElement, &size, sizeof(int32_t));
  std::memcpy(mNextElement + sizeof(int32_t), buffer.data(), buffer.size());
  mQ->mReadCounter.store(mLocalCounter, std::memory_order_release);
  mNextElement += payloadSize; // advance the read counter
}

int32_t QConsumer::TryRead(std::span<std::byte> buffer)
{
  if (mLocalCounter == mQ->mReadCounter.load(std::memory_order_acquire))
  {
    return 0; // nothing to read
  }

  int32_t size;
  std::memcpy(&size, mNextElement, sizeof(int32_t)); // data race
  int32_t writeCounter = mQ->mWriteCounter.load(std::memory_order_acquire);
  EXPECT(writeCounter – mLocalCounter <= QUEUE_SIZE, "queue overflow");
  EXPECT(size <= buffer.size(), "buffer space isn’t large enough");
  std::memcpy(buffer.data(), mNextElement + sizeof(size), size); // data race
  const int32_t payloadSize = sizeof(size) + size;
  mLocalCounter += payloadSize;
  mNextElement += payloadSize;
  writeCounter = mQ->mWriteCounter.load(std::memory_order_acquire);
  EXPECT(writeCounter– mLocalCounter <= QUEUE_SIZE, "queue overflow");
}
```

we have a data race situation, <cpp>std::memcpy</cpp> is not atomic. the performance is ok, but not great, so we try to make it better.

first, we try to avoid moving the write-counter every time.

```cpp
void QProducer::Write(std::span<std::byte> buffer)
{
  const int32_t payloadSize = sizeof(int32_t) + buffer.size(); 
  mLocalCounter += payloadSize;
  // we "reserve" more space (X% of the total queue) 
  // to avoid touching this cache line on every message written
  if (mCachedWriteCounter < mLocalCounter)
  {
    mCachedWriteCounter = Align<Q_WRITE_COUNTER_BLOCK_BYTES>(mLocalCounter);
    mQ->mWriterCounter.store(mCachedWriteCounter, std::memory_order_release);
  }
  std::memcpy(mNextElement, &size, sizeof(int32_t));
  std::memcpy(mNextElement + sizeof(int32_t), buffer.data(), buffer.size());
  mQ->mReadCounter.store(mLocalCounter, std::memory_order_release);
  mNextElement += payloadSize;
}
```

next, we try and optimize the data alignment

```cpp
void QProducer::Write(std::span<std::byte> buffer)
{
  const int32_t payloadSize = sizeof(int32_t) + Align<Q_BLOCK_ALIGNMENT>(buffer.size()); 
  mLocalCounter += payloadSize;
  // ...
}
```

and caching the read counter

```cpp
int32_t QConsumer::TryRead(std::span<std::byte> buffer)
{
  // we might already know from the previous read counter that more data is available, and
  // in this case we avoid reading this cache line for no reason
  if (mLocalReadCounter == mCachedReadCounter)
  {
    mCachedReadCounter = mQ->mReadCounter.load(std::memory_order_acquire);
  }

  if (mLocalReadCounter == mCachedReadCounter)
  {
    return 0;
  }
  // ...
}
```

another option is to avoid copies, serialize the data directly into the queue. this is an API change.

#### Measurements

measurements are intrusive, and add overhead to performance. we don't know in advance where the bottleneck will be. we can use simple scoped measurements, or have some instrumentation framework. but we also need audits and alerts on the data we gather. this is the eighth principle - "Being fast is good - staying fast is better".

"Thinking about the system as a whole" is the ninth principle.

#### Summary - Principles

1. No node containers.
2. Understanding your problem (by looking at data).
3. Hand tailored (specialized) algorithms are key to achieve performance.
4. Simplicity is the ultimate sophistication.
5. Mechanical sympathy.
6. True efficiency is found not in the layers of complexity we add, but in the unnecessary layers we remove.
7. Choose the right tool for the right task.
8. Being fast is good - staying fast is better.
9. Thinking about the system as a whole.

</details>

### Fast and Small C++ - When Efficiency Matters - Andreas Fertig

<details>
<summary>
Some ways that code can be made smaller and more performant.
</summary>

[Fast and Small C++ - When Efficiency Matters](https://youtu.be/rNl591__9zY?si=9nmg1MK_9S9pwtvU), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Fast_and_small_cpp.pdf), [event](https://cppcon2024.sched.com/event/1gZfc/fast-and-small-c-when-efficiency-matters)

#### Unique Pointer optimization

starting with a <cpp>std::unique_ptr</cpp> that defines a custom deleter, and checking that it's size is the same as two pointers.

```cpp
auto f = std::unique_ptr<FILE, decltype(&fclose)>{fopen("SomeFile.txt", "r"), &fclose};
static_assert(sizeof(f) == (2 * sizeof(void*)));
```

can we get this behavior with less memory? why do we need to pay for the extra pointer? there is an optimization for empty base classes, if the class is empty, it has the size of zero, but since it must have an address, the size becomes 1. however, if we derive from such an empty base class, the size of the base class becomes zero again, and we only pay for the data in the derived class.

```cpp
class Base {
public:
  void Fun() { puts("Hello, EBO!"); }
};

class Derived : public Base {
  int32_t mData{};
public:
};

void Use()
{
  Derived d{};
  static_assert(sizeof(d) == sizeof(int32_t));
  d.Fun();
}
```

we can use this optimization for our custom case. this already happens in the standard library for the default deleter. here is a simplified implementation.

```cpp
template<class T>
struct default_delete {
  default_delete() = default;
  constexpr void operator()(T* ptr) const noexcept
  {
    static_assert(0 < sizeof(T), "can't delete an incomplete type");
    delete ptr;
  }
};

template<typename T, typename U>
struct CompressedPair {
  [[no_unique_address]] T first; // special flag for C++20
  [[no_unique_address]] U second; // special flag for C++20
  CompressedPair(U s) : second{s} {}
  CompressedPair(T f, U s) : first{f}, second{s} {}
};

template<class T, class Del = default_delete<T>>
class unique_ptr {
  CompressedPair <Del, T*> mPair; // internal data, pointer and deleter
public:
  unique_ptr(T* ptr) : mPair{ptr} {}
  unique_ptr(T* ptr, Del deleter) : mPair{deleter, ptr} {}
  unique_ptr(const unique_ptr&) = delete;
  unique_ptr operator=(const unique_ptr&) = delete;
  unique_ptr(unique_ptr&& src) : mPair{std::exchange(src.mPair.second, nullptr)} {}
  unique_ptr& operator=(unique_ptr&& src)
  {
    mPair.second = std::exchange(src.mPair.second, mPair.second);
    mPair.first = std::exchange(src.mPair.first, mPair.first);
    return *this;
  }

  ~unique_ptr()
  {
    if(mPair.second) { mPair.first(mPair.second); } // if the pointer isn't nullptr, call the deleter on the pointer
  }

  T* operator->() {return mPair.second;}
};
```

so if we modify the original example, we pass a captureless lambda using <cpp>decltype</cpp> with the function pointer.

```cpp
template<typename T, auto DeleteFn>
using unique_ptr_deleter = 
  std::unique_ptr<T, decltype([](T*obj) { DeleteFn(obj); })>;

auto f = unique_ptr_deleter <FILE, fclose>{fopen("SomeFile.txt", "r")};
static_assert(sizeof(f) == sizeof(void*));
```

we still want something better, so we need to move to C++23. we define a static call operator on the lambda, which removes the implicit "this" parameters, and this should save us some assembly operations.

```cpp
template<typename T, auto DeleteFn>
using unique_ptr_deleter = 
  std::unique_ptr<T, decltype([](T*obj) static { DeleteFn(obj); })>;
auto f = unique_ptr_deleter <FILE, fclose>{fopen("SomeFile.txt", "r")};
static_assert(sizeof(f) == sizeof(void*));
```

#### Implementing the small string optimization

we can also look at a naive implementation small string optimization, which allows us to store a bit of data (up to 15 characters plus the null characters) without going to the heap. however, the single boolean value that denotes if the string is optimized ends up costing us additional 7 bytes for padding along the alignment.

```cpp
struct string {
  size_t mSize{};
  size_t mCapacity{};
  char* mData{};
  char mSSOBuffer[16]{};
  bool mIsSSO{true};
};
static_assert(sizeof(string) == 48);
```

can we do the same thing without exceeding 24 bytes of data? the standard library manages it.\
for libstd++ it combines the capcity and the buffer data, since if we are optimizing the string, the capacity is known. we only get 7 bytes for small string (rather than 15), but we require only half the memory. Ms-STL does it a bit different, and libc++ has another approach and employ bit fiddeling which gives us the same 15 bytes to store data without using the heap.\
each implementation does things in a different way, which focuses on different things, which means we have branches in different operations, and the implementation optimizes for different use cases.

```cpp
// libstdc++
struct string {
  char* mPtr;
  size_t mSize{};
  union {
    size_t mCapacity;
    char mBuf[8];
  };
  /* more code */
};

// MS STL
struct string {
  union {
    char* mPtr;
    char mBuf[8];
  };
  size_t mSize{};
  size_t mCapacity{};
  /* more code */
};

// libc++
struct string {
  static constexpr unsigned BIT_FOR_CAP{sizeof(size_t) * 8 − 1};
  struct normal {
    size_t large : 1;
    size_t capacity : BIT_FOR_CAP; // MSB for large bit
    size_t size;
    char* data;
  };

  struct sso {
    uint8_t large : 1;
    uint8_t size : (sizeof(uint8_t) *8) − 1; // large+size == sizeof(uint8_t)
    uint8_t padding[sizeof(size_t) − sizeof(uint8_t)]; // Padding large + size + padding == sizeof(size_t)
    char data[sizeof(normal) − sizeof(size_t)];
  };

  union {
    sso small;
    normal large;
  } packed;
/* more code */
};
```

we could inspect another optimization by facebook, this one is designed for long text, so the optimization is for cases when the heap is used and tries to allow as much space before going to the heap (23 bytes). there is some playing with the most significant bit as well.

```cpp
// fb-string
struct string {
  struct normal {
  char* data;
  size_t size;
  size_t capacity; // virtually reduced by one byte
  };

  struct sso {
    char data[sizeof(normal)]; // MSB for long string mode indicator
  };

  union {
    sso small;
    normal large;
  } packed;
  /* more code */
};
```

#### The power of `constexpr` and initializer list

changing a `const static` function to `constexpr` can improve performance, both in debug mode and with optimization flags. we can see it both from the number of assembly instructions and the instructions themselves, in the example it even fixes the layout in memory.

<cpp>std::initializer_list</cpp> is transformed into a backing vector, which means we can avoid paying for reading the data during runtime. something about recursions and backing arrays.

```cpp

void Receiver(const int list[4]) noexcept; // forward declaration
void Fun1() noexcept
{
  const int list[4]{3, 4, 5, 6};
  Receiver(list);
}

void Fun2() noexcept
{
  static const int list[4]{3, 4, 5, 6}; // better optimization
  Receiver(list);
}

void Receiver(std::initializer_list<int> list) noexcept; // forward declaration

void Fun3()
{
  std::initializer_list <int> list{3, 4, 5, 6}; // behaves differently in c++26 compliant compilers
  Receiver(list);
}
```

</details>

### Can You RVO? Using Return Value Optimization for Performance in Bloomberg C++ Codebases - Michelle Fae D'Souza

<details>
<summary>
Return Value optimizations - when they happen and when to use.
</summary>

[Can You RVO? Using Return Value Optimization for Performance in Bloomberg C++ Codebases](https://youtu.be/WyxUilrR6fU?si=N-84m2JJ6Lfs627T), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Can_You_RVO.pdf), [event](https://cppcon2024.sched.com/event/1gZgE/can-you-rvo-using-return-value-optimization-for-performance-in-bloomberg-c-codebases), [compiler explorer playground](https://tinyurl.com/ICanRVO123).

RVO - return value optimizations

#### What is RVO

Some history: in 1997, copy ellison was added to the standard. it defines when copy-ellison can be used.however, this still isn't used as much as it should, and this can cost in performance and in clean code.

in short, copy elision is removing the construction of temporary objects and constructing them directly on the caller location, this saves us from copying and destroying the temporary object. there are two kinds of RVO - regular (unnamed) and named. NRVO returns a l-value object.

```cpp
MyObj example1()
{
  // NRVO
  MyObj a = MyObj(3);
  return a;
}

MyObj example2()
{
  // RVO
  return MyObj(3);
}

int main()
{
  MyObj e1 = example1();
  MyObj e2 = example2();
  return 0;
}
```

> Does my compiler have to support RVO?
>
> - Compilers are allowed to perform URVO since C++98
> - Compilers are required to provide URVO support since C++17
> - NRVO support is optional, but recommended

We can turn down RVO with a compilation flag `-fno-elide-constructors`. it turns down Named RVO, but not Unnamed RVO. for msvc we need to turn it on with `/Zc:nrvo`, or use `/O2` optimization, and some other cases which apply this behavior.

#### Anti-Patterns That Prevent RVO

RVO doesn't happen if we disable it, if we return something that was constructed from outside the scope of the current function (returning a global, a parameter), or if the return type isn't the same as the thing being returned, such as returning a derived class from a function that declares to return a base class, if there are multiple return statements NRVO won't happen. NRVO won't happen when retruning a complex expression.

```cpp
MyObj example3()
{
  // two possible return values of the same type
  MyObj x1 = MyObj(3);
  MyObj x2 = MyObj(3);
  int a = rand() % 100;
  if (a > 50) {
    return x1;
  }
  return x2;
}

MyObj example4()
{
  // complex expression
  MyObj x1 = MyObj(3);
  MyObj x2 = MyObj(3);
  int a = rand() % 100;
  return std::move(a>50 ? x1 : x2);
}
```

for NRVO - either the copy or move constructors must exist (not deleted), URVO can work even if they are deleted.

some examples: seeing if there's RVO. branches, expressions, CV qualifiers, return l-value or p-value variables, throwing expressions. returning parameters, returning a member from a local object, structured bindings, and other cases...

> WARNING: If you are writing a copy / move constructor, never make the constructor do anything else other than a copy/move, because it can get elided!

#### Using RVO for Performance Gains

we shouldn't rewrite the codebase for RVO, it's usually a micro-optimization, so unless profiling tells us it's a bottleneck, existing code can remain as it is.\
tooling can help us detect when we have a possibility for RVO gains, like using <cpp>std::move</cpp> incorrectly.

</details>

### Session Types in C++ - A Programmer's Journey - Miodrag Misha Djukic

<details>
<summary>
Defining session types (interactions) and expressing them in C++.
</summary>

[Session Types in C++ - A Programmer's Journey](https://youtu.be/2uzYFhJjbNM?si=iPzH4z0vdVpQjH6A), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Session_Types_in_Cpp.pdf), [event](https://cppcon2024.sched.com/event/1kXEp/session-types-in-c-a-programmers-journey)

session types is a mathematical term. this talk will try to implement the concept in C++.

#### What are Types

we start by looking at the "type" part of the session types, in mathmatical terms, types are connected to "sets". so we need to define what a set is. this is a rabbit hole, so we ignore it. we focus on types in a programming sense. types have different roles, and for most cases, a single type encompasses all the roles, but there are exceptions.

> What are types used for in programming?\
> Types are used for:
>
> - Abstraction
>   - Knowing a type of something is enough to "work" with it. (We do not need to know all
the details.)
> - Documentation
>   - Type informs us how we should interact with something.
> - Efficiency
>   - Carefully chosen type can lead to more efficient code.
> - Expressivity
>   - Meaning is encoded in both operation and operands' type.
> - Detecting errors
>   - Doing something (by accident) that does not "fit" the type will be detected as error.
> - Safety
>   - Varying explanations.
>   - A guarantee that there are no certain kinds of errors.
>   - Being unable to do a "wrong thing".

types aren't the only ways to achieve everything, but they are part of them. so, a defintion that encompasses the above usage might be something like:

> "A type defines a set of possible values and a set of operations (for an object)."\
> ~ Bjarne Stroustrup, Programming Principles and Practices Using C++.

#### What is a Session

> - Interaction of two or more entities.
> - It has a beginning and (usually) an end.
> - In between, a sequence of interactions is happening.

if we have an example of a client providing the server two numbers, and then an operation (addition or division) and getting a response back.
we could describe it with an interaction diagram, which is also doing the things we wanted the type to do. we could also describe sessions as a valid sequence of interactions, or a formula:

- side A: "?int; ?int; &(ADD: !int, DIV: !double); end".
- side B: "!int; !int; &CirclePlus;(ADD: ?int, DIV: ?double); end"

sessions can involve more than two processes, but we focus on binary session types for now. so session types describe behavior: what is the message, what is allowed, what is the order. (also: no deadlocks, eventual termination).

(this is like a state machine).

an example of a "scribble" protocol defintion

```scribble
module scribble.example.Purchasing;

type <xsd> "QuoteRequest" from "Purchasing.xsd" as QuoteRequest;

global protocol BuyGoods (role Buyer, role Seller) {
  quote(QuoteRequest) from Buyer to Seller;

  choice at Seller {
    quote(Quote) from Seller to Buyer;
    buy(Order) from Buyer to Seller;
    buy(OrderAck) from Seller to Buyer;
  } or {
    quote(OutOfStock) from Seller to Buyer;
  }
}
```

in practice, we use session types in three different ways.

1. write and verify protocol, generate code, use type.
2. write code, extract protocol and verify
3. use language facilities to describe the protocol, and then write a type, and extract protocol defintion.

we will use the third way, starting from the middle, not abstract protocol defintion, but also not direct code.

we could take a functional approach, monads, creating new objects, continuation passes. we need some way to have communication (channels between actors), either as using queues or networking sockets. we use a simplified abstraction.

```cpp
struct Comm {
  Comm(SQ& forSending, SQ& forReceiving);
  //..
  template<typename T>
  void send(T x) { /**/ }
  template<typename T>
  T recv() { /**/ }
};
Comm(Q1, Q2); // One end-point
Comm(Q2, Q1); // The other end-point

void serverFunc(Comm chan) {
  auto v1 = chan.recv<int>();
  auto v2 = chan.recv<int>();
  chan.send(v1 + v2); // addition
  chan.close();
}

void clientFunc(Comm chan) {
  cout << "First num: ";
  int x;
  cin >> x;
  chan.send(x);
  cout << "Second num: ";
  cin >> x;
  chan.send(x);
  auto r = chan.recv<int>();
  cout << "Result: " << r << endl;
  chan.close();
}
```

we could introduce more types, make things more explicit, write for one side, and express the other side in those terms. using templates, channels, loops and stuff to express our interaction.\
using <cpp>std::move</cpp> on the "this" object,

</details>

### Designing C++ Code Generator Guardrails: A Collaboration Among Outreach And Development Teams And Users - Sherry Sontag, Cb Bailey

<details>
<summary>
The process of modernizing the code generator tool.
</summary>

[Designing C++ code generator guardrails: A collaboration among outreach and development teams and users](https://youtu.be/sAfUQUs_GbI?si=VSOmj_qO-ZWK_TSc), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Designing_Cpp_Code_Generator_Guardrails.pdf), [event](https://cppcon2024.sched.com/event/1gZhD/designing-c-code-generator-guardrails-a-collaboration-among-outreach-and-development-teams-and-users).

Code generation tool in Bloomberg internal infrastructure, a talk about an RFC (request for comments) and how the modernization process went.

cycles of shrinking and growing executables, each library adds more code and serialization, so the client-server architecture must keep working, and they need internal tooling to make sure things run, and this is achieved by having a code generator.
the code generator handles the serialization, sending and receiving requests, and providing the basic framework for networking, threading, logging and metrics. all for allowing the users (internal developers inside the company) to focus on the business logic.\
over time, those types escaped outside the bloomberg services eco-system, and things become hectic. a lot of libraries, sometime duplicating behavior, creating bloat and making things confusing.

their goal:\
reducing bloat, creating guidelines for new libraries, reviewing existing libraries. outreach to teams about how to use new tooling, how to name libraries, and what should be in it. stripping debug information from all generated code libraries, writing validtores for the new policies.

</details>
