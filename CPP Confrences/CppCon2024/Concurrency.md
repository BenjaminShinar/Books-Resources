<!--
// cSpell:ignore movabs dsll daddiu addi slli movk sethi sllx
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Concurrency

<summary>
2 Talks about concurrency.
</summary>

- [x] Coroutines And Structured Concurrency In Practice - Dmitry Prokoptsev
- [x] Performance Engineering - Being Friendly To Your Hardware - Ignas Bagdonas

---

### C++ Coroutines and Structured Concurrency in Practice - Dmitry Prokoptsev

<details>
<summary>
A story about bringing coroutines into an existing codebase.
</summary>

[C++ Coroutines and Structured Concurrency in Practice](https://youtu.be/aPqMQ7SXSiw?si=SsDO3NjVTLBAIMwC), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Coroutines_and_Structured_Concurrency_in_Practice.pdf), [event](https://cppcon2024.sched.com/event/1gZeP/coroutines-and-structured-concurrency-in-practice)

> C++20 coroutines present some exciting opportunities for organizing and simplifying concurrent logic, but they have proven challenging to integrate with production systems. Object lifetime and execution order issues can be difficult to manage, especially without guidance from well-established best practices. Practical use requires interoperability with existing event loops and with non-coroutine-aware components. While some experimental libraries have started to be released, the C++ "async ecosystem" is still young compared to many other languages.\
> We will describe our path through this thicket and introduce Corral, a new open-source concurrency framework for C++ that attempts to tame it. Corral is built around structured concurrency principles, so lifetimes and control flow are easier to reason about. It does not provide any opinionated I/O layer, so it can work with the one you're already using, or with a standard such as Boost.Asio. We have ported a number of our I/O-bound production processes to use this system, and have found the modernized versions to be substantially simpler with fewer weird bugs.\
> This talk is geared towards C++ developers who have some familiarity with coroutines but have found them challenging to work with. We will delve into the architectural design of our framework, offering insights that may be useful to others with similar needs. Special emphasis will be placed on practical aspects such as task cancellation, timeouts, and integration with legacy code. We will also share a few design patterns we've discovered after two years of production use.

replacing legacy code using callbacks with new coroutines. this might help in simplifying things, but can also create problems.

a typical async framework has classes to represent work (tasks), a way to detach those tasks to work alongside the rest of the program, and a way to explicitly suspend the current work until a task is complete, and then to handle the result or any errors.

with detached work, there is a problem with lifetime management,

in this example, pass the reference to a local buffer, which will be removed before the coroutine runs.

```cpp
// don’t do this
void bad(tcp::socket& s) {
    std::array<char> buf(1024);
    asio::co_spawn(
    ex,
    [&s, &buf]() -> asio::awaitable {
        co_await s.read_some(s,
            asio::buffer(buf),
            asio::use_awaitable);
    },
    asio::detached);
}
```

there is also the problem of handling errors in detached tasks, we lose the ability to handle them.

even if we decided to never use detached tasks and always join them, a task might be written in a way that never completes. we need to establish a rule that tasks can only run if there is someone waiting for them.

> A task can only run when it’s being
> awaited by another task.
>
> 1. That awaiting task is a caller
> 2. Once the caller resumes, the callee is done any resources it may have used can be freed
> 3. Any unhandled exception will get re-raised

this is how we write traditional code, and we want to keep those principals in the concurrent world. hence the term "structured concurrency".

we add some combiners - they can take several tasks and combine the results together. such as `allOf()`, `anyOf()`, using this allows us to add a way to handle signals and graceful shutdown. this forms a tree of suspended tasks, so we can propagate errors.

an example of a simple tcp echo server. the first example creates a recursive program, which isn't great.

```cpp
Task<void> serve(tcp::socket s) {
    std::array<char> buf(1024);
    try {
        for (;;) {
            size_t len = co_await s.async_read_some(
                asio::buffer(buf), use_awaitable);
            co_await async_write(s, asio::buffer(buf, len),
            asio::use_awaitable);
        }
    } catch (std::exception&) { /*connection closed or I/O error*/ }
}

Task<void> listen(asio::io_context& io_context, tcp::acceptor& acc) {
    tcp::socket s = co_await acc.async_accept(io_context,
    asio::use_awaitable);
    co_await anyOf(serve(std::move(s)), accept(acc));
}
```

we need a dynamic `allOf` combiner, which is actually a common concurrent concept, which is sometimes called a *nursery*.

> Properties of nurseries and combiners
>
> - Act as nodes in the task tree
> - Wait until all children complete
> - Propagate any exceptions from any children back to the parent
> - Cancel any children still running before completing
> - No task is ever left behind

we are missing Asynchronous destructors in C++. this causes us some problems.

> - <cpp>void await_suspend(std::coroutine_handle<> h)</cpp> - Initiate the asynchronous operation, and arrange h.resume() to be called when it completes
> - <cpp>auto await_resume()</cpp> - Fetch the result of the operation, or (re)throw an exception.
> - <cpp>bool await_ready() const noexcept</cpp> - Performance optimization

example of DNS resolution in a concurrent world.

#### Task Cancellation

in our design, we can always cancel, since we have the chain of parent-child tasks. our cancellations should be implicit (not written separately), fast (not using exceptions), and asynchronous themselves.

C++ has a <cpp>std::coroutine_handle\<void>::destroy()</cpp>, but it needs the awaitable object to co-operate with it. so we need to extended the coroutine interface with a new function `await_cancel`. we still need to know the awaitable and it's type. there are more and more stuff we need, including stuff about the coroutine frame.

#### Resource Management

RAII is even more important than ever, but the cleanup function might also be asynchronous, so we need a combiner that kicks off on cancellation.

#### Bridging into Legacy Code

legacy code probably has callbacks, and we would need our coroutines to interact with them, either call the callbacks or replace them. we might need to break our rules about structured concurrency.

</details>

### Performance Optimization in Software Development - Being Friendly to Your Hardware - Ignas Bagdonas

<details>
<summary>
Understanding The Memory, Processor, and what happens under the hood.
</summary>

[Performance Optimization in Software Development - Being Friendly to Your Hardware](https://youtu.be/kv6yqNjCjMM?si=0bMHaNXfGZFk7JVF), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Performance_Engineering.pdf), [event](https://cppcon2024.sched.com/event/1gZee/performance-engineering-being-friendly-to-your-hardware)

C++ code runs on an abstract machine, but in actual, it runs on hardware - processor, memory and I/O, with some interconnect between them.

#### Memory

memory is a an array of capacitor, which we can address into. we read the charge of the capacitor, and refresh it many times, this is the terms on nano-seconds. there are tons of physics and stuff involved, this puts a cap on the performance we can get.

> DDR architecture – it is only a fraction of bus transaction that is in fact
>
> - Row addressing
> - Column addressing
> - Generation to generation capacitor array operates at mostly the same frequency, bus transfer speed in fact increases
> - Multiple memory modules (sockets)
> - Page size – CPU vs memory
> - Address mapping

in our timeline, we have a lot of time where we're not doing any work, so if we have more memory bands, we could increase throughput. since one memory band can only have one row open at the same time, we could add more bands on the same command (?bus?). so we combine the bands into band groups. however, the latency doesn't increase, and in many cases it actually decreases.

there are some marketing terms about memory, which might help in some cases, but aren't an outright improvement for all cases. different memory compositions result in different performance profile.\
if we really care about getting the last bit of performance, we need to write the code in a way that utilizes the specific memory configuration, and restrict our application to only run on those machines.

#### Processor Core

steps:

1. get next block of partial instruction
2. linear fetch
3. incoming branch
4. instruction alignment
5. instruction fusing

points of interest, each of these is an entire complex thing by itself.

- instruction fetch
- branch prediction
- instruction decoding
- instruction input
- scheduling
- execution

not all instructions are equal, especially with X86 and variable length instructions, but in the hardware we could decode them in parallel.

code density also matters

```cpp
uint64_t v = 0x123456789abcdef0;
```

for x86, it one instruction

```x86
movabs r10, 0x123456789abcdef0
# 49 ba f0 de bc 9a 78 56 34 12
```

for other architectures, the number of instructions might be larger.

```mips
# MIPS
li $2, 38141952
ori $2, $2, 0x8acf
dsll $2, $2, 17
daddiu $2, $2, 9903
dsll $2, $2, 18
ori $2, $2, 0xdef0
# 46 02 02 3c cf 8a 42 34 78 14 02 00 af 26 42 64 b8 14 02 00 f0 de 42 34

# RISC-V
li a5, 305418240
addi a5, a5, 1657
li a0, -1698897920
slli a5, a5, 32
addi a0, a0, -272
add a0, a5, a0
# 12 34 57 b7 67 97 87 93 9a bc e5 37 02 07 97 93 ef 05 05 13 00 a7 85 33

# ARM
mov x0, 0xdef0
movk x0, 0x9abc, lsl 16
movk x0, 0x5678, lsl 32
movk x0, 0x1234, lsl 48
# 00 de 9b d2 80 57 b3 f2 00 cf ca f2 80 46 e2 f2

# SPARC
sethi %hi(0x12345400), %g1
sethi %hi(0x9ABCDC00), %o0
or %g1, 0x278, %g1
or %o0, 0x2F0, %o0
sllx %g1, 32, %g1
add %g1, %o0, %o0
# 15 8d 04 03 37 af 26 11 78 62 10 82 f0 22 12 90 20 70 28 83 08 40 00 90
```

could we somehow identify those instructions and make a single operation out of them? yes. this is an implementation behavior, it is not specified in the ISA standard.

everything that happens before the execution is speculative, only the execution stage does the actual work for external observable behavior, everything else could be discarded and flushed in some occasions.

address spaces - translate logical (virtual) and physical addresses.

```cpp
uint32_t fn(uint32_t x, uint32_t y, size_t cond) {
    if (cond == 27)
        return (x << 4) + y;
    else
        return (x >> 1) - y;
}

uint32_t res = fn1(x, y, 42);
uint32_t res = fn1(x, y, 27);
```

in machine code, this becomes the following code, we have a branch miss-prediction, and we discard work, but at the next time we again have a branch-miss, so we do again the same flush. a smarter machine could do both branches in parallel, and only decide at the end what to return.

```mips
cmp rdx, 27
jne L_0F
mov eax, edi
shl eax, 4
add eax, esi
jmp L_15

L_0F:
    mov eax, edi
    shr eax, 1
    sub eax, esi
L_15:
    ret
    nop 10
```

#### Some examples

the <cpp>memcpy</cpp> example, we could implement it in many ways, getting different machine code, which would internally be handled differently.

> - Problem space
> - Performance requirements
> - Scalar, various vectors, specialty instructions, on-core and offcore accelerators
> - Data layout: both software and hardware characteristics
>   - alignment (source and destination)
>   - size
>   - direction
>   - linearity

an example of hashing, how it could be performed at machine level.

measuring performance is hard, and micro-benchmarking isn't always a good indication for actual real-world performance.

</details>
