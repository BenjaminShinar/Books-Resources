<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Performance

<summary>
6 Talks about Performance.
</summary>

- [ ] Can You RVO? Using Return Value Optimization for Performance in Bloomberg C++ Codebases - Michelle Fae D'Souza
- [ ] Designing C++ code generator guardrails: A collaboration among outreach and development teams and users - Sherry Sontag, CB Bailey
- [ ] Fast and small C++ - When efficiency matters - Andreas Fertig
- [ ] Limitations and Problems in std::function and Similar Constructs: Mitigations and Alternatives - Amandeep Chawla
- [ ] Session Types in C++: A Programmer's Journey - Miodrag Misha Djukic
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
