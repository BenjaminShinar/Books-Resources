<!--
ignore these words in spell check for this file
// cSpell:ignore mCoroHdl coro Fibo Goldblatt stackless Amdahl ftor cartoonify Tempesta softirq Htrie
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Concurrency

### C++20’s Coroutines for Beginners - Andreas Fertig

<details>
<summary>
Simple examples of using coroutines.
</summary>

[C++20’s Coroutines for Beginners](https://youtu.be/8sEe-4tig_A), [slides](https://andreasfertig.com/talks/dl/afertig-2022-cppcon-cpp20--coroutines-for-beginners.pdf), [github](https://github.com/CppCon/CppCon2022).

function vs coroutine comparison in control flow. coroutines can suspend and resume operation.

- function: call &rarr; return.
- coroutine: call &rarr; co_yield (suspend) &rarr; co_return (finish).

coroutines decide by themselves when to suspend. in c++ the coroutine is _stackless_, which means that the data is stored on the heap. coroutines allow us to remove state maintainence code and to replace function pointers (callbacks).

| Keyword     | Action | State     |
| ----------- | ------ | --------- |
| `co_yield`  | output | suspended |
| `co_return` | output | ended     |
| `co_await`  | input  | suspended |

the coroutine has a wrapper type which wraps around the coroutine return type, together with the _promise_type_, it becomes a **finite state machine**.

#### Examples

a chat example. `Fun` returns _Chat_ object, and uses `co_yield`, `co_await` and `co_return`. `Use` calls the function and uses it with the `listen` and `answer` functions defined as part of the _chat_ object.

```cpp
Chat Fun() // Wrapper type Chat containing the promise type
{
  co_yield "Hello!\n"s;// Calls promise_type.yield_value
  std::cout << co_await std::string{}; // Calls promise_type.await_transform
  co_return "Here!\n"s; // Calls promise_type.return_value
}

void Use()
{
  Chat chat = Fun(); // Creation of the coroutine
  std::cout << chat.listen(); // Trigger the machine
  chat.answer("Where are you?\n"s); // Send data into the coroutine
  std::cout << chat.listen(); // Wait for more data from the coroutine
}

```

going over the promise type of the coroutine, we have some customization points:

- `void unhandled_exception() noexcept`
- `std::suspend_always initial _suspend() noexcept`
- `std::suspend_always yield_value(std::string msg) noexcept`
- `constexpr bool await_ready() const noexcept`
- `std::string await_resume() const noexcept`
- `void await_suspend(std::coroutine handle<>) const noexcept`
- `void return_value(std::string) noexcept`
- `std::suspend_always final_suspend() noexcept`

```cpp
struct promise_type {
  std::string _msgOut{};
  std::string _msgIn{}; // Storing a value from or for the coroutine

  void unhandled_exception() noexcept {} // What to do in case of an exception
  Chat get_return_object() { return Chat{this}; } // Coroutine creation
  std::suspend_always initial_suspend() noexcept { return {}; } // Startup
  std::suspend_always yield_value(std::string msg) noexcept // Value from co_yield
  {
    _msgOut = std::move(msg);
    return {};
  }

  auto await_transform(std::string) noexcept // Value from co_await
  {
    struct awaiter { // Customized version instead of using suspend_always or suspend_never
      promise_type& pt;
      constexpr bool await_ready() const noexcept { return true; }
      std::string await_resume() const noexcept {
        return std::move(pt._msgIn);
      }
      void await_suspend(std::coroutine_handle <>) const noexcept {}
    };

    return awaiter{*this};
  }

  void return_value(std::string msg) noexcept { _msgOut = std::move(msg); } // Value from co_return
  std::suspend_always final_suspend() noexcept { return {}; } // Ending
};
```

and some other functions:

```cpp
struct Chat {
  #include "promise−type.h"

  using Handle = std::coroutine_handle <promise_type >; // Shortcut for the handle type
  Handle mCoroHdl{}; //

  explicit Chat(promise_type* p) : mCoroHdl{Handle::from_promise(*
  p)} {} // Get the handle form the promise
  Chat(Chat&& rhs) : mCoroHdl{std::exchange(rhs.mCoroHdl, nullptr)} {} // Move only!

  ~Chat() // Care taking, destroying the handle if needed
  {
    if(mCoroHdl)
    {
      mCoroHdl.destroy();
    }
  }

  std::string listen() // Active the coroutine and wait for data.
  {
    if(not mCoroHdl.done())
    {
      mCoroHdl.resume();
    }
    return std::move(mCoroHdl.promise()._msgOut);
  }

  void answer(std::string msg) // Send data to the coroutine and activate it.
  {
    mCoroHdl.promise()._msgIn = msg;
    if (not mCoroHdl.done())
    {
      mCoroHdl.resume();
    }
  }
};
```

the compiler creates a coroutine block on the heap. a _Task_ is a coroutine that does a job without returning a value, while a _Generator_ returns a value (either by `co_return` or `co_yield`).

there are helper types in the library

| type                  | result of `await_ready` | notes                    |
| --------------------- | ----------------------- | ------------------------ |
| `std::suspend_always` | false                   | always waits for a value |
| `std::suspend_never`  | true                    | never suspend            |

interleaving two vectors: taking one element of each vector and making a new vector.

usage:

```cpp
void Use()
{
  std::vector a{2, 4, 6, 8};
  std::vector b{3, 5, 7, 9};
  Generator g{interleaved(std::move(a), std::move(b))};

  while(not g.finished())
  {
    std::cout << g.value() << '\n';
    g.resume();
  }
}
```

using a generator inside a a generator. calling `co_yield` to return tha value from the nested coroutines.

```cpp
Generator interleaved(std::vector<int> a, std::vector<int> b)
{
  auto lamb = [](std::vector<int>& v) −> Generator {
      for(const auto& e : v) { co_yield e; }
  };

  auto x = lamb(a);
  auto y = lamb(b);

  while(not x.finished() or not y.finished())
  {
    if(not x.finished()) {
      co_yield x.value();
      x.resume();
    }

    if(not y.finished()) {
      co_yield y.value();
      y.resume();
    }
  }
}
```

with the promise_type, which is mostly boiler plate, and we customize the `yield_value` to store the value received from the `yield_value` call. we also modify the `initial_suspend` and `final_suspend` using the helper types.

```cpp
struct promise_type
{
  int _val{};
  Generator get_return_object() { return Generator{this}; }
  std::suspend_never initial_suspend() noexcept { return {}; }
  std::suspend_always final_suspend() noexcept { return {}; }
  std::suspend_always yield_value(int v)
  {
    _val = v;
    return {};
  }
  void unhandled_exception() {}
};
```

and the generator, which uses the promise_type.

```cpp
//struct Generator {
using Handle = std::coroutine_handle <promise_type >;
Handle mCoroHdl{};
explicit Generator(promise_type* p) : mCoroHdl{Handle::from_promise(*
p)} {}
Generator(Generator&& rhs) : mCoroHdl{std::exchange(rhs.mCoroHdl, nullptr)} {}

~Generator()
{
  if(mCoroHdl) {
    mCoroHdl.destroy();
    }
}

int value() const { return mCoroHdl.promise()._val; }
bool finished() const { return mCoroHdl.done(); }
void resume()
{
  if(not finished())
  {
    mCoroHdl.resume();
  }
}
```

now we want to use range-based for loops, and not the while loop. for this we need an iterator (and a sentinel), with equality operator, increment operator and dereference operator. and now we can hide the fact that there are coroutines.

```cpp
struct sentinel {};

struct iterator {
  Handle mCoroHdl{};

  bool operator==(sentinel) const
  {
    return mCoroHdl.done();
  }

  iterator& operator++()
  {
    mCoroHdl.resume();
    return * this;
  }

  const int operator* () const
  {
    return mCoroHdl.promise()._val;
  }
};
```

Scheduling multiple tasks: letting something else manage how tasks are run.

```cpp
void Use()
{
  Scheduler scheduler{};
  taskA(scheduler);
  taskB(scheduler);
  while(scheduler.schedule()) {}
}
```

we define the simple task, which give up control/

```cpp
 Task taskA(Scheduler& sched)
{
  std::cout << "Hello, from task A\n";
  co_await sched.suspend();
  std::cout << "a is back doing work\n";
  co_await sched.suspend();
  std::cout << "a is back doing more work\n";
}

Task taskB(Scheduler& sched)
{
  std::cout << "Hello, from task B\n";
  co_await sched.suspend();
  std::cout << "b is back doing work\n";
  co_await sched.suspend();
  std::cout << "b is back doing more work\n";
}
```

and the scheduler, which starts working on a task, but when a task calls the suspend function, it's returned to the list of tasks.

```cpp
struct Scheduler {
  std::list<std::coroutine_handle<>> _tasks{};
  bool schedule()
  {
    auto task = _tasks.front();
    _tasks.pop_front();
    if(not task.done())
    {
      task.resume();
    }
    return not _tasks.empty();
  }

  auto suspend()
  {
      struct awaiter : std::suspend_always {
      Scheduler& _sched;
      explicit awaiter(Scheduler& sched) : _sched{sched} {}
      void await_suspend(std::coroutine_handle <> coro) const noexcept {
        _sched._tasks.push_back(coro);
      }
    };

  return awaiter{*this};
  }
};
```

the promise_type isn't complicated, it just defines the behaviors (never suspends or waits for data).

```cpp
struct Task {
  struct promise_type
  {
    Task get_return_object() { return {}; }
    std::suspend_never initial_suspend() noexcept { return {}; }
    std::suspend_never final_suspend() noexcept { return {}; }
    void unhandled_exception() {}
  };
};
```

we could also use a global scheduler (static variable). with some slight changes.

</details>

### Deciphering C++ Coroutines - A Diagrammatic Coroutine Cheat Sheet - Andreas Weis

<details>
<summary>
Understanding the stages and steps of coroutines.
</summary>

[Deciphering C++ Coroutines - A Diagrammatic Coroutine Cheat Sheet](https://youtu.be/J7fYddslH0Q), [slides](https://github.com/ComicSansMS/presentations/releases/download/cppcon2022/deciphering_coroutines.pdf)

> "Coroutines are hard!"

coroutines can conceptualized a functions that can be existed and resumed in the middle. coroutines are always **stateful**, not just the location in the coroutine, but also the context. c++20 coroutines are stackless.\
it might be better to think of coroutines not as regular functions, but as some sort of context factory?

#### Essential use cases for coroutines

asynchronous computation, with callbacks

```cpp
auto [ec, bytes_read] = read(socket, buffer);

async_read(socket, buffer,[](std::error_code ec, std::size_t bytes_read){...});
```

but with coroutines, we no longer need callback and inverted control, the code looks much simpler, and more like blocking code, but without blocking. there are some considerations, like who wakes up the function, and on which thread.

```cpp
auto [ec, bytes_read] = read(socket, buffer);

auto [ec, bytes_read] = co_await async_read(socket, buffer);
```

another use case is with lazy evaluations:

```cpp
MyCoroutine co = startComputation(initial_data);
auto some_results = co.provide(some_data);
auto more_results = co.provide(more_data);
auto final_results = co.results;
```

here is a a simple example, computing the fibonacci sequence. a naive approach is to return all the numbers up to N, but this storage intensive, and requires to know upfront what's needed. we can create a "generator" class. but this class doesn't have to be a coroutine, the interface is all that matters.

```cpp
// here is a naive approach
std::vector<int> Fibo(int n);

// approach 2, creating a stateful functor.
class FiboGenerator {
// Successive calls to next () return the
// numbers from the Fibonacci series
// 1 , 1 , 2 , 3 , 5 , 8 , 13 , 21 , ...
int next ();
};

// Returns a new FiboGenerator object that
// will start from the first Fibonacci number
FiboGenerator makeFiboGenerator ();
```

#### Coroutines from the caller’s perspective

as said before: coroutines can be conceptualized as factory functions which return a function that computes the result.

the **Return Type** is what matters.

the signature doesn't tell us that a function is a coroutine, coroutines use the following keywords:

- `co_await`
- `co_return`
- `co_yield`

if one of those appears, the function is a coroutine. in this example, the body returns an int, but the return type is something else.

```cpp
FiboGenerator makeFiboGenerator() {
  int i1 = 1;
  int i2 = 1;
  while (;;) {
    co_yield i1; // return the number
    i1 = std::exchange(i2 ,i1 + i2);
  }
}
```

The simplest example: we must define the return type, which must have a nested type called _promise_type_. it's similar, but not the same as the _promise_ type from the threading libraries (futures). but they work differently.

the promise*type handles the lifecycle of the coroutine. the \_promise_type* is nested, but it's not a member.\
the return type doesn't have to be nested type, it could be specialized as the _std::coroutine_traits_.

#### Steps involved in starting a coroutine

- `ReturnType get_return_object()`
- `void return_void()`
- `void unhandled_exception()`
- `std::suspend_always initial_suspend()`
- `std::suspend_always final_suspend()`

we can use _std::suspend_always_ or _std::suspend_never_, or any awaitable which defines the following member functions

```cpp
ReturnType hello_coroutine () {
  co_return;
}

struct ReturnType {
  struct promise_type {
    ReturnType get_return_object() {
      return {};
    }
    std::suspend_always initial_suspend() {
      return {};
    }
    void return_void() { };
    void unhandled_exception() { };
    std::suspend_always final_suspend() noexcept {
      return {};
    }
  };
};
```

now we can try running the code, at this point, we don't see nothing printed yet, the function is still suspended.

```cpp
// running the code
std::suspend_always initial_suspend() {
  return { };
}

ReturnType hello_coroutine () {
  std::println ("Hello from coroutine!");
  co_return;
}
int main () {
  hello_coroutine (); // prints nothing
}
```

if we want to see this in action, we need to tell the function to start when it's being called, without requiring us to resume it.

```cpp
// not suspending at startup
std::suspend_never initial_suspend() {
  return { };
}

ReturnType hello_coroutine () {
  std::println ("Hello from coroutine!");
  co_return;
}
int main () {
  hello_coroutine (); // prints "Hello from coroutine!"
}
```

#### Suspend and resume

The **Awaitable** is a type that we can call `co_await` on, with or without an argument. these are _opportunities_ to suspend the execution. the awaitable controls what happens in these case, it can be suspended, or not. it's a customization point.

key components:

- Return Type
- Promise Type
- Awaitable Type
  - `await_ready()`
  - `await_suspend(std::coroutine_handle<>)`
  - `await_resume()`
- _std::coroutine handle<>_
  - `void resume() const`
  - `void destroy() const`
  - `promise_type& promise() const` - convert
  - `static coroutine_handle from_promise -(promise_type &)` - convert

```cpp
AsyncRead awaitable = async_read(socket, buffer);
auto [ec,bytes_read] = co_await awaitable;

struct Awaitable {
  bool await_ready();
  void await_suspend(std::coroutine_handle promise_type>);
  void await_resume();
};

struct std::coroutine_handle<promise_type> {
// ...
  void resume() const;
  void destroy() const;
  promise_type& promise() const;
  static coroutine_handle from_promise(promise_type &);
};
```

a cheat sheet:

1. the _ReturnType_ - this determines the promise type. could be nested or specialized in the _std::coroutine_traits_.
2. the _promise_type_ - can have a constructor.
3. suspension point
   1. initial_suspend
   2. final_suspend
4. either return a value or void
5. handle exceptions somehow

```cpp
struct ReturnType
// std::coroutine_traits<ReturnType, ...>
{
  struct promise_type {
    promise_type(T...); // optional
    ReturnType get_return_object();
    std::suspend_always initial_suspend();
    // ---- ⇑ Start / ⇓ Shutdown ----
    void return_value(T);
    // void return_void();
    void unhandled_exception();
    std::suspend_always final_suspend() noexcept;
  };
};
```

in the earlier example, our coroutine never fired because it was suspended at the start, but it could be resumed.

```cpp
// running the code
std::suspend_always initial_suspend() {
  return { };
}

ReturnType hello_coroutine () {
  std::println("Hello from coroutine!");
  co_return;
}

int main() {
  ReturnType c = hello_coroutine (); // prints nothing
  std::println("resuming the coroutine");
  c.resume(); // prints "Hello from coroutine"
}
```

#### Drawing a map of coroutine land

- caller &rarr; ReturnType
- promise &#8644; coroutine_handle
- coroutine &rarr; Awaitable

we connect the types, we define the return type to accept and hold a coroutine_handle<promise_type> as member, and then we can operate on it through the ReturnType object as a handler.

```cpp
struct promise_type {
ReturnType get_return_object()
return ReturnType{  std::coroutine_handle<promise>::from_promise(* this)};
}
struct ReturnType {
ReturnType (std::coroutine_handle<promise_type> h) : handle(h) {} // constructor
std::coroutine_handle<promise_type> handle; // member function
void resume() {
  handle.resume();
  }
};
```

#### Interacting with coroutines

we want to coroutine to produce some data that the user can grab, the coroutine has connection to the awaitable, and it constructs it we the `co_await` is called. so we pass it as an argument. we want to pass it into the promise_type, which is the connection point between the coroutine and the ReturnType that the user has.\
if we want to pass data into a coroutine, we can also do that, this can be done when the coroutine goes to sleep (is suspended). this is also done via the returnType and the promise_type. the `co_await` can also return a value inside the function.

- await_suspend - takes the handler
- await_resume - uses the handler to continue executing.

when we want to use _co_yield_, we need to implement `yield_value` function. it also allows us to jump over the Awaitable.

symmetric transfer - handing control back to a different coroutine, rather than to the caller.

</details>

### C++ Concurrency TS-2 Use Cases and Future Direction - Michael Wong, Maged Michael, Paul McKenney

<details>
<summary>
Overview of the Concurrency technical specification second proposal (TS-2) and what we might see coming from it.
</summary>

[Concurrency TS-2 Use Cases and Future Direction](https://youtu.be/3sO4IrWQPnc),[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CPPCON-2022_Concurrency_TS_2_Use_Cases.pdf)

TS - technical specification, Concurrency TS1 had std::future, latches, barriers, and plans for atomic smart pointers.

two features will probably be pushed forward into the standard. maybe even c++26.

- Hazard Pointers
- Read Copy Updates

both of those a part of the **Deferred Reclamation** idea. they are low-level APIs for it. it's a lock-free algorithm.

#### Hazard Pointers,Read Copy Updates, Snapshots, Fences

> - Readers access data while holding reader locks or data is protected
>   - Guarantee data will remain live while lock is held or data is protected
> - One or more updaters update data by replacing it with newly allocated data
>   - All subsequent readers will see new value
>   - Old values is not destroyed until all readers access it have released their locks
>   - Here is where you can have 2 views of Schrödinger’s Cat: one alive and one dead
> - Benefits
>   - readers never block the updater or other readers
>   - Updaters never block readers
> - What you pay: Updates have extra cost, could be very small
>   - They need allocation and new values construction
>   - OK if updates are rare

we trade certainty for performance and scalability.

approaches:

1. Reader Writer Locks
2. Reference counts
3. RCU
4. Hazard Pointers

> Deferred Removal via Reference Counting:\
> Combines waiting for readers and multiple versions:
>
> - Writer removes the cat's element from the list (Unlink cat)
> - Writer waits for all readers to finish
> - Writer can then free the cat's element

new readers can't get the element because it was removed from the chain, but this doesn't account for multiple versions.

> Beyond performance, you also need to choose from other
> properties of lock-free programming

(some of this table is only theoretically interesting)

> | Metric                                            | Reader Writer Locks  | Reference Counting                                    | RCU Hazard                             | Pointers          |
> | ------------------------------------------------- | -------------------- | ----------------------------------------------------- | -------------------------------------- | ----------------- |
> | **Readers**                                       | Slow and un-scalable | Slow and un-scalable                                  | Fast and scalable                      | Fast and Scalable |
> | **Unreclaimed objects**                           | None                 | None                                                  | Unbounded                              | Bounded           |
> | **Traversal speed**                               | No or low overhead   | Atomic RMW updates                                    | No or low overhead                     | Low overhead      |
> | **Reference acquisition**                         | Unconditional        | Depends on use case                                   | Unconditional                          | Conditional       |
> | **Contention among readers**                      | Can be very high     | Can be very high                                      | No contention                          | No contention     |
> | **Automatic reclamation**                         | No                   | Yes                                                   | No                                     | No                |
> | **Reclamation timing**                            | Immediate            | Immediate                                             | Deferred                               | Deferred          |
> | **Non-blocking traversal**                        | Blocking             | Either blocking or lock free with limited reclamation | Bounded population oblivious wait free | Lock free.        |
> | **Non-blocking reclamation** (nomemory allocator) | Blocking             | Either blocking or lock free with limited reclamation | Blocking                               | Bounded wait free |

**Snapshot**:\
a high-level interface for deferred reclamation. applies to a single object, built on top of RCU or hazard pointers.

**Asymmetric Fences**:\
speed up the common cases by sacrificing the uncommon ones.

there are some tips of using TS-2 in a safe way.

#### Hazard Pointers

> Protect access to objects that may be concurrently removed.
> A hazard pointer is a single-writer multi-reader pointer.\
> If a hazard pointer points to an object before its removal, then the object will not be reclaimed as long as the hazard pointer remains unchanged.\
> features:
>
> - Fast and scalable protection
> - supports arbitrarily long protection.

there are protectors who set hazard pointers, and there are processors which remove and reclaim them.

there is an interface proposal, but there is subset of essential apis.

```cpp
template <typename T> class hazard_pointer_obj_base {
  void retire() noexcept; // Object must be already removed
};

class hazard_pointer {
  hazard_pointer() noexcept; // Construct an empty hazard pointer
  hazard_pointer(hazard_pointer&&) noexcept;
  hazard_pointer& operator=(hazard_pointer&&) noexcept;
  ~hazard_pointer();
  template <typename T> bool try_protect(T*& ptr, const atomic<T*>& src)
  noexcept;
  template <typename T> T* protect(const atomic<T*>& src) noexcept;
  template <typename T> void reset_protection(const T* ptr) noexcept;
};

hazard_pointer make_hazard_pointer(); // Construct a non-empty hazard pointer
void swap(hazard_pointer&, hazard_pointer&) noexcept;
```

> - _hazard_pointer_obj_base_: base type of objects protect-able by hazard pointers.
>   - retire: removed object is to be reclaimed when no longer protected.
> - _hazard_pointer_: hazard pointer object, may be empty, a nonempty hazard pointer object owns a hazard pointer.
>   - `hazard_pointer()`: constructs an empty hazard pointer object.
>   - `operator=(hazard_pointer&&)`: moves hazard pointer objects,
>     ends moved to and continues moved from protection if any,
>     moved from becomes empty.
>   - `~hazard_pointer()`: destroys the hazard pointer object, ends protection by the owned hazard pointer if any.
>   - `try_protect(ptr, src)`: protects ptr only if src equals ptr. -` protect(src)`: protects a pointer from src.
>   - `reset_protection(ptr)`: ends current protection if any, starts protecting ptr if not null and not removed.
> - `make_hazard_pointer`: constructs a nonempty hazard pointer object.
> - `swap`: swaps two hazard pointer objects Hazard Pointers TS2 Interface.

examples:

1. protecting arbitrarily long access
1. hand-over-hand traversal
1. iteration

#### RCU Trick

(technical difficulties)

order of possible situations, state changes, the common case of going fast and the uncommon case of working carefully and slowly (being synchronized).

</details>

### An Introduction to Multithreading in C++20 - Anthony Williams

<details>
<summary>
Multithreading strategies and synchronization methods.
</summary>

[An Introduction to Multithreading in C++20](https://youtu.be/A7sVFJLJM-A), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/introduction_to_multithreading_cpp20-1.pdf)

#### Choosing your Concurrency Model

there are two fundamental reasons for using multiple threads:

- Scalability - run on any kind of processor
- Separation of concerns - processes which are distinct from one another, shouldn't interfere with other

choosing one reason informs our concurrency model.\
if we want scalability, we follow Amdahl's law to calculate the speedup :\
$
SpeedUp = \frac{1}{1-p+\frac{p}{n}}
$

standard library algorithms have parallel versions, using the execution strategy. it is usually a good idea to stick to them if possible.

```cpp
std::vector<MyData> data;
std::sort(std::execution::par, data.begin(), data.end(), myComparator{});
```

the problem is that using standard algorithms are hard to combine together into a single parallel algorithms.

another option is to implement a thread pool (still not standard) and split independent workloads between threads. the thread pool controls the execution order, and how the hardware parallel is used.

```cpp
thread_pool tp;

void foo(){
  execute(tp, []{do_work();});
  execute(tp, []{do_other_work();});
}
```

if we aren't concerned with throughput and performance, and we wish to have parallel task running in the background, each doing a different work. in this case, we might be more concerned about responsiveness. this might call for dedicated threads.

```cpp
std::jthread gui{[]{run_gui();}};
std::jthread printing{[]{do_printing();}};
```

#### Starting and Managing Thread

cooperative cancellations: ability to cancel an operation, but it's undesirable to stop a running thread with a system call. in c++20 there are cooperative cancellation options: _std::stop_source_ and _std::stop_token_. they are cooperative because the running task must check them, and if they aren't checked, nothing happens.

1. create a _std::stop_source_
2. create a _std::stop_token_ from the _std::stop_source_
3. pass the _std::stop_token_ to a new thread of task
4. when you want to stop the operation, call `source.request_stop()`
5. periodically call `token.stop_requested()` to check and respond. what the thread does it completely up to it.

```cpp
void stoppable_func(std::stop_token st){
  while(!st.stop_requested()){
    do_stuff();
  }
}

void stopper(std::stop_source source){
  while(!done()){
    do_something();
  }
  source.request_stop();
}
```

we can also provide a custom cancellation, by using _std::stop_callback_. this time it's not cooperative, the function runs in the calling thread.

```cpp
Data read_file(std::stop_token st, std::filesystem::path filename) {
  auto handle=open_file(filename);
  std::stop_callback cb(st,[&]{ cancel_io(handle);});
  return read_data(handle); // blocking
}
```

stop tokens integrate with _std::jthread_, _std::async_ can be used when we want a result from the thread (which doesn't fit the idea of a background thread), _std::thread_ should remain th last option.\
When we use _std::jthread_ it passes a cancellation token. the destructor calls the stop token to stop the operation, _std::jthread_ is a moveable value type, non copyable. the callable and the arguments are copied to the storage local, if we really references, we can use _std::ref_.

```cpp
void thread_func(std::stop_token st, std::string arg1,int arg2){
  while(!st.stop_requested()){
    do_stuff(arg1,arg2);
  }
}
void foo(std::string s){
  std::jthread t(thread_func,s,42);
  do_stuff();
} // destructor requests stop and joins
```

we can get the stop token and stop source from the thread, and we can request the stop directly, without using the token.

#### Synchronizing Data

Shared state - watchout for data races - un-synchronized access to memory location, where one of the thread performs a write. this is undefined behavior, so we need some method to synchronize the data:

- Latches - _std::latch_ - a single use counter that allows threads to wait for the count to reach zero.
- Barriers - _std::barrier_ - a reusable barrier that stops an amount of threads, and when it releases them, it is reset.
- Futures - _std::future_ - one shot transfer of data between threads (_std::async_, _std::promise_, _std::packaged_task_). can also store an exception (which can terminate the process if not captured).
- Mutexes - mutual exclusion - preventing concurrent execution, forces a serial section. mutex are directly operated on, while locks implement RAII behavior over mutexes. most forms of mutex and locks should be avoided.
  - _std::mutex_ - (**use this**)
  - _std::timed_mutex_
  - _std::recursive_mutex_
  - _std::recursive_timed_mutex_
  - _std::shared_mutex_
  - _std::shared_timed_mutex_
  - _std::scoped_lock_ - (**use this**) multiple mutex
  - _std::unique_lock_ - work with _std::conditional_variable_
  - _std::lock_guard_ - one mutex, backwards compatibility
  - _std::shared_lock_ - for _std::shared_mutex_
- Semaphores - _std::counting_semaphore<max_count>_ - limited available "slots" for execution, there is no need to acquire a slot before releasing or not.
  - _std::binary_semaphore_ - an alias for _std::counting_semaphore<1>_
- Atomics - _std::atomic\<T>_ - lock-free(maybe), using intrinsic locking instructions.
  - _std::atomic_flag_ - guaranteed to be lock-free
  - _std::atomic_signed_lock_free_ - guaranteed to be lock-free
  - _std::atomic_unsigned_lock_free_ - guaranteed to be lock-free

latches are great for coordinating threads in tests to run at the same time. barriers are great for loops, and we can define a _completion function_ to do something between the loops. futures are useful for polling and waiting for data to be ready, _std::async_ takes the _std::launch_ argument(_async_, _deferred_), we can't call `.get()` more than once, except for _std::shared_future_. mutex objects force a serial section inside parallel code, and can be used with locks for RAII behavior. _std::scoped_lock_ can be used with multiple mutexes and prevent deadlocks (theoretically, even if they are passed in different order).

we sometimes want to pass data between threads (and we can't use futures), this can be achieved with **busy wait**, but we should avoid it, as it consumes CPU time waiting, wastes electricity running, and can actually delay the other process from running triggering the notification. instead, we can use _std::conditional_variable_. _std::unique_lock_ is used by it and locks and unlocks itself. when the _std::conditional_variable_ waits, it releases the lock, and when it wakes up, it acquires it again. a different thread needs to notify the _std::conditional_variable_ in order to wake it up for a check _std::conditional_variable_any_ allows for cancelling waits with a _std::stop_token_.\
_(note: we can use code blocks or fences to limit the scope of the lock, this makes code simpler and might decrease the delay)_\
Semaphores can be used to limit how many concurrent threads execute, a thread can operate on a semaphore (acquire,release) without a lock. a binary semaphore can be used as a mutex. acquiring calls can block or be timed. for _std::counting_semaphore_, the template parameter defines the maximal amount of of slots, and the parameter in the constructor is the inital number of slots. for atomics, _std::atomic<T>_ - the type must be trivially copyable and bitwise comparable, except for _std::shared_ptr_ or _std::weak_ptr_. not actually guaranteed to be lock free in all cases, might use a mutex internally.

> summary:
>
> - Avoid managing your own threads if you can.
> - Use _std::jthread_ for threads.
> - Use _std::stop_token_ for cancellation.
> - Use _std::future_, _std::latch_ and _std::barrier_ where you can.
> - Use _std::mutex_ almost everywhere else.
> - Use _std::atomic_ in rare cases.

</details>

### A Lock-Free Atomic Shared Pointer in Modern Cpp - Timur Doumler

<details>
<summary>
Looking at what we can do to have a lock-free atomic shared pointer that is protected from undefined behavior.
</summary>

[A Lock-Free Atomic Shared Pointer in Modern Cpp](https://youtu.be/gTpubZ8N0no),

```cpp
std::atomic<std::shared_ptr<const T>>
```

a shared global state object,

something to do with RCU (read, copy, update) and hazard pointers. There is a proposal by Herb Sutter, the problems are managing lifetime and ABA.

a **shared pointer** a smart pointer that manages an object and shares ownership of it, usually by reference count. we usually have the shared pointer point to a control block, which has the count and points to the data. **atomic** is used for thread safety, but the shared pointer only guarantees the reference count is thread safe, not the data.

here is an example of a data race:

```cpp
auto ptr = std::make_shared<widget>();
//thread 1
auto ptr1 = ptr;

//thread 2
auto ptr = ptr2;

//this is undefined behavior, a data race
```

since c++11, we have functions that can do stuff on shared pointers in atomic way, the free function with the `atomic_` prefix. however, the implementation was messy, and it's prone to errors, because it's a free function that if we don't use them in one place, the whole program becomes undefined behavior.

```cpp
auto ptr = std::make_shared<widget>();
//thread 1
auto ptr1 = std::atomic_load(&ptr);

//thread 2
std::atomic_store(&ptr, ptr2);

// Ok (deprecated in C++20)
```

since c++20, we got `std::atomic<std::shared_ptr<T>>` as a type, so we don't need to use special free functions, or even the atomic function if we want to keep things simple and use the same interface as any other shared pointer.

```cpp
std::atomic<std::shared<widget>> ptr = std::make_shared<widget>();
//thread 1
//auto ptr1 = ptr.load();
auto ptr1 = ptr; // same as above

//thread 2
//ptr.store(ptr2)
ptr=ptr2; // same as above

// Ok (since c++20)
```

lock-free means that we can't acquire mutexes, unfortunately atomic shared pointer is not lock-free in the current standard.

```cpp
std::atomic<std::shared<widget>> ptr = std::make_shared<widget>();
assert(ptr.is_lock_free()); // false
static_assert(decltype(ptr)::is_always_lock_free); // false
```

however, there are some people who talked about having implementations of lock-free atomic shared pointers, and there are some open source versions on github.

a good, simple and efficient way of doing this will require us to have one CPU instruction that performs two memory modifications (the control block pointer from the object on the stack and the reference count value), called _double compare-and-swap_. modern CPUs don't have that kind of instruction. but there is _double-width compare-and-swap_ (DWCAS), which does exist on modern CPUs, but it can only swap two "words" which are one next to another, which isn't helpful to us, because we need to modify the shared\*ptr object and the control block at the same time, and they can't be located next to one another (one is on the stack, one on the heap).

however, the _split refcount_ algorithm can help us, but it's not a common algorithm

> "It's sort of a folk algorithm"\
> ~ David Goldblatt

the basic idea is that the atomic shared pointer object on the stack also has a reference count (local refcount), and the control block has the global refcount. so now the pointer to the control block and the local refcount are located next to one another.

1. increment the local refcount.
2. increment the global refcount
3. decrement the local refcount.

so in the end, the global reference count is increased, but we can always roll back the changes.

```cpp
template <typename T>
class atomic<shared_ptr<T>>{
  struct counted_ptr {
    control_block* cb;
    size_t local_refcount;
  };

  atomic<counted_ptr> cptr;
  static_assert(decltype(cptr)::is_always_lock_free);
};
```

> atomic copy: basic idea
>
> - atomically increment local_refcount
>   - remember resulting [cb, local_refcount] pair
> - try to sync refcounts:
>   - atomically increment cb->global_refcount
>   - atomically decrement local_refcount
>   - if another thread modified [cb, local_refcount] in the meantime:
>     - try again
>   - else:
>     - success!
>     - return new shared_ptr instance with value [cb, local_refcount]

example of the `load` function

```cpp
template <typename T>
class atomic<shared_ptr<T>>{
  struct counted_ptr {
    control_block* cb;
    size_t local_refcount;
  };

  atomic<counted_ptr> cptr;
  static_assert(decltype(cptr)::is_always_lock_free);

  shared_ptr<T> load(){
    // 1. increment local refcount
    auto cptr_copy = cptr.load();
    while (true){
      auto cptr_newval = cptr_copy;
      ++cptr_newval.local_refcount;
      if (cptr.compare_exchange_weak(cptr_copy,cptr_newval)) // make sure no one else modified this in the meantime
        break;
    }
    ++cptr_copy.local_refcount;

    // 2. get copy
    auto cb = cptr_copy.cb;

    // 3. increment global refcount
    cb->global_refcount.fetch_add(1);

    // 4. decrement local refcount
    auto cptr_expected = cptr_copy;
    while(true){
      auto cptr_newval = cptr_expected;
      --cptr_newval.local_refcount;

      if(cptr.compare_exchange_weak(cptr_expected,cptr_newval)){
        break; // good case!
      }

      // if value changed, we were not supposed to modify global_refcount!
      if (cptr_expected->cb != cptr_copy->cb){
        cb->global_refcount.fetch_sub(1);
        break;
      }
    }

    return shared_ptr<T>(cb);
  }
};
```

Performance:

> each atomic<shared_ptr> access:
>
> - at least 3 atomic operations, more under contention
> - slower than mutex when no contention
> - but much better than mutex under high contention
>
> folly optimization: "refcount batching"
>
> - increment global and local refcount to some number K > numThreads
> - next K/2 loads do not have to touch control block (=1 atomic operation)
> - tradeoff: `use_count()` member function becomes meaningless

however, that's not enough to make a lock free atomic shared pointer. the above relies on a very simplified model of shared pointer, because the shared pointer also has a pointer to the object, not just to the control block. we need this pointer for aliasing constructor. the shared pointer manages the lifetime of one object, but the data is in another object (which should have the same life time)

```cpp
struct widget{
int i=0;
int j=0;
};

int main(){
  auto pw = std::make_shared<widget>(2,5);
  std::shared_ptr<int> pj(pw, &pw->j);
}
```

or in a more common example, using inheritance.

```cpp
std::shared_ptr<base> ptr(new derived);
```

so now we need to somehow modify three objects in memory in an atomic way. and there is no cpu instruction to modify three memory locations.

there is one version by Anthony Williams that uses additional extra memory allocation, but for most memory allocators, allocation is not lock-free.

> suggested actions:
>
> - move allocation from atomic<shared_ptr> to shared_ptr constructor.
>   - but that's overhead for every aliased shared_ptr.
> - compute ptrdiff between object ptr and aliased ptr.
>   - if it fits into 32 bits (on x64):
>     - just store that directly
>   - otherwise:
>     - use extension list (will happen very rarely)
>   - add 1 bit flag to indicate method (use an alignment bit)

still not enough, we have another problem, this time with the Double-Width Compare-And-Swap instruction. the CPU supports it, but the standard doesn't support it, because of ABI. there are some ways to try and get around it. one suggested way is to make use of the unused bits in the address (in 64bit machines, the default is to use 48 bits out of the 64, with the last three being alignment bits). but it becomes more problematic, and we might have to re-implement the shared pointer to use DWCAS.

there is no current open source implementation that is portable, lock-free for aliasing, standard conforming and production ready.

</details>

### A Pattern Language for Expressing Concurrency in Cpp - Lucian Radu Teodorescu

<details>
<summary>
composing concurrent work and using patterns as building blocks
</summary>

[A Pattern Language for Expressing Concurrency in Cpp](https://youtu.be/0i2MnO2_uic), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon2022-pres.pdf), [github example of server application](https://github.com/lucteo/structured_concurrency_example).

> What we want:
>
> - safety
> - performance
> - structured approach

the `std::execution` (senders and receivers) proposal wasn't accepted into c++23.

#### Structured Concurrency

extending the ideas of **Structured Programming**:

> 1. abstractions as building blocks
> 2. recursive decomposition
> 3. local reasoning
> 4. single entry, single exit point
> 5. soundness and completeness

functions are core part of the idea, using functions, avoiding *GO TO*s, and using "sequence, selection, repetition" as a guideline.

concurrency - with primitives (threads and mutexes). barely pass the ideas.
concurrency - with primitives (tasks) - a bit better.
the upcoming (maybe C++26) senders and receivers executioners does fit the requirements.

| Concurrency Model            | thread primitives           | task primitives | std::execution | coroutine |
| ---------------------------- | --------------------------- | --------------- | -------------- | --------- |
| abstractions building blocks | no                          | yes             | yes            | yes       |
| recursive decomposition      | no                          | no              | yes            | yes       |
| local reasoning              | no                          | yes             | yes            | yes       |
| single entry and exit points | -                           | no              | yes            | yes       |
| soundness and completeness   | &half; (maybe completeness) | yes             | yes            | yes       |

senders describe computations - **any chunk of work with one entry and one exit point**. a computation can be a task, mutiple tasks over multiple threads, a group of computations, or even the entire application. function operate on the same thread, while computations can have a different entry and exit threads.

coroutine tasks function similar to senders.

examples:

hello world with composition, each sender is a building block

```cpp
namespace ex = std::execution;

auto say_hello() {
  return ex::just() // just a signal
        | ex::then([] {
          std::printf("Hello, concurrent world!\n");
          return 0;
          }
        );
}

int main() {
  auto [r] = std::this_thread::sync_wait(say_hello()).value();
  return r;
}
```

- `just` -> _()-> sender_
- `then` -> _(sender, ftor)-> sender_
- `snyc_wait` -> _(sender)->optional<tuple<vals...>>_

a sender describes work, and eventually produces a result (similar to a future).

example 2, using coroutines.

```cpp
task<int> say_hello() {
  std::printf("Hello, concurrent world!\n");
  co_return 0;
}
int main() {
  auto [r] = std::this_thread::sync_wait(say_hello()).value();
  return r;
}
```

a rock-paper-scissors example, work is done in parallel

```cpp
enum class shape { rock, paper, scissors };
shape player_choose() { return static_cast<shape>(rand() % 3); }
void print_result(shape r1, shape r2); // not defined here

void play_rock_paper_scissors() {
  static_thread_pool pool{8};
  ex::scheduler auto sched = pool.get_scheduler();
  ex::sender auto game_work = ex::when_all( //
    ex::schedule(sched) | ex::then(player_choose), //
    ex::schedule(sched) | ex::then(player_choose) //
  );
  auto [r1, r2] = sync_wait(std::move(game_work)).value();
  print_result(r1, r2);
}
```

> Concurrency analysis:
>
> - threads are hidden
> - expressed concurrency as a graph
> - no explicit synchronization
> - framework ensures efficiency

#### Patterns

The problem is how to describe concurrency, functions and senders:

**exit points:**

| functions                   | senders              |
| --------------------------- | -------------------- |
| return value (or void)      | `set_value(vals...)` |
| exceptions                  | `set_error()`        |
| stopped `exit()` was called | `set_stopped()`      |

patterns:

1. Pattern 1: creating value
2. Pattern 2: synchronous wait
3. Pattern 3: transforming values
4. Pattern 4: joining
5. Pattern 5: scheduling
6. Pattern 6: composition
7. Pattern 7: starting sender in other contexts
8. Pattern 8: transfer between contexts

Pattern 1: creating value - zero or more values, starting a a flow

- `just()`
- `just_error(e)`
- `just_stopped()`

```cpp
void create_value_example() {
  ex::sender auto s = ex::just(13);
  auto [r] = sync_wait(s).value();
  assert(r == 13);
  std::printf("%d\n", r);
}

void create_value_example2() {
  ex::sender auto s = ex::just(17, 19);
  auto [a, b] = sync_wait(s).value();
  assert(a == 17);
  assert(b == 19);
  std::printf("%d, %d\n", a, b);
}

void create_value_example3() {
  ex::sender auto s = ex::just();
  sync_wait(s); // finishes immediately
}
```

Pattern 2: synchronous wait - waiting for a sender to complete (blocking)

```cpp
void create_value_example() {
  ex::sender auto s = ex::just(13); // sender
  auto [r] = sync_wait(s).value(); // waiting
  assert(r == 13);
  std::printf("%d\n", r);
}
```

Pattern 3: transforming values - from input sender

- `then()`
- `upon_error()`
- `upon_stopped()`

```cpp
int process_value(int x) {
  std::printf("I've got value: %d\n", x);
  return x*x;
}

void then_example() {
  ex::sender auto s = ex::just(13); // creating
  ex::sender auto s2 = ex::then(s, process_value); // transform
  auto [r] = sync_wait(std::move(s2)).value(); // wait
  assert(r == 169);
  std::printf("%d\n", r);
}

void then_example2() {
  ex::sender auto s = ex::just(13)
    | ex::then(process_value); // pipeline syntax
  auto [r] = sync_wait(std::move(s)).value();
  assert(r == 169);
  std::printf("%d\n", r);
}
```

Pattern 4: joining - combine parallel work, detecting when multiple execution paths have finished

- `when_all()`

```cpp
void join_example() {
  ex::sender auto s = ex ::when_all( //
    ex::just(7), //
    ex::just(3.14), //
    ex::just("hello!") //
  );
  auto [i, d, str] = sync_wait(std::move(s)).value();
  std::printf("%d, %g, %s\n", i, d, str);
}
```

Pattern 5: scheduling - parallel work chain, start work on a different execution context, and transform a scheduler into a sender.

- `schedule(scheduler)`
- `transfer_just(vals...)`

```cpp
void schedule_example() {
  static_thread_pool pool{8};
  ex::scheduler auto sch = pool.get_scheduler();
  ex::sender auto s = ex::schedule(sch) // from scheduler to sender
    | ex::then([]{std::printf("Hello from another thread!");}); // this happens on a different thread

  sync_wait(std::move(s)); //wait for s to happen
}
```

next we can move to higher order patterns, which are created by combining the basic patterns together.

Pattern 6: composition - compose senders together, ensure data is alive for the entire lifetime. _(monadic bind use case)_

- `let_value()`
- `let_error()`
- `let_stopped()`

(the difference between `let_value` and `then` is that composition returns a sender, rather then the value)

```cpp
ex::sender auto schedule_request_start(read_requests_ctx ctx) { /*...*/ }
ex::sender auto validate_request(const http_request& req) { /*...*/ }
ex::sender auto handle_request(const http_request& req) { /*...*/ }
ex::sender auto send_response(const http_response& resp) { /*...*/ }

ex::sender auto request_pipeline(read_requests_ctx ctx) {
  return schedule_request_start(ctx) // normal creation
    | ex::let_value(validate_request)
    | ex::let_value(handle_request)
    | ex::let_value(send_response)
  ;
}
```

Pattern 7: starting sender in other contexts - work needs to start at a specific sender

- `on()`

```cpp
ex::sender auto do_read_from_socket() { /*...*/ }
io_context io_threads;
ex::scheduler auto sch = io_threads.get_scheduler();
ex::sender auto snd = ex::on(sch, do_read_from_socket());
sync_wait(std::move(snd));
```

it is equivelent to using `ex::schedule` and the `ex::let_value` together.

```cpp
ex::on(sch, snd);
ex::schedule(sch) | ex::let_value([] {return snd;})
```

Pattern 8: transfer between contexts - work needs to change context, or new work on the same context

- `transfer(scheduler)`

the `ex::on` starts work in a context, while `ex::transfer` moves the already concurrent work into a different context.

```cpp
ex::sender auto read_from_socket() { /*...*/ }
ex::sender auto process(in_data) { /*...*/ }
ex::sender auto write_output(out_data) { /*...*/ }

io_context io_threads;
static_thread_pool work_pool{8};
ex::scheduler auto sch_io = io_threads.get_scheduler();
ex::scheduler auto sch_cpu = work_pool.get_scheduler();

ex::sender auto snd = ex::on(sch_io, read_from_socket()) // start work here
 | ex::transfer(sch_cpu) // move to other context (thread)
 | ex::let_value(process) //
 | ex::transfer(sch_io) // return to the first context (thread)
 | ex::let_value(write_output)
 ;
sync_wait(std::move(snd));
```

there are other stuff which will come once the P2300 proposal is accepted:

- `bulk`
- `split`,`ensure_started`
- `start_detected`
- `into_variant`, `stopped_as_optional`, `stopped_as_error`
- `read`
- coroutine support

#### Compositions

now that we have a vocabulary of patterns, we can start building applications, such the http server application for image processing (see github). the entire application is a sender (since senders are a generalization of functions, it's not something radical).

the basic structure of the application:

```cpp
auto get_main_sender() {
 return ex::just() | ex::then([] {
 //... the entire application logic
 return 0;
 });
}
auto main() -> int {
  auto [r] = std::this_thread::sync_wait(get_main_sender()).value();
  return r;
}
```

```cpp
auto listener(int port, io::io_context& ctx, static_thread_pool& pool)-> task<bool> {
 // ... some coroutine
 co_return true;
}

auto get_main_sender() {
  return ex::just() | ex::then([] { // start work without any input value, and then perform some actions in the lambda
    int port = 8080;
    static_thread_pool pool{8}; // thread pool
    io::io_context ctx;
    set_sig_handler(ctx, SIGTERM);
    ex::sender auto snd = ex::on(ctx.get_scheduler(), listener(port, ctx, pool));
    ex::start_detached(std::move(snd));
    ctx.run();
    return 0;
  });
}
```

a closer look into the _listener_ coroutine (return type is task, and has the `co_return` keyword) ,

```cpp
auto listener(int port, io::io_context& ctx, static_thread_pool& pool)-> task<bool> {
  io::listening_socket listen_sock;
  listen_sock.bind(port);
  listen_sock.listen();
  while (!ctx.is_stopped()) {
    io::connection conn = co_await io::async_accept(ctx, listen_sock); // another coroutine
    conn_data data{std::move(conn), ctx, pool};
    ex::sender auto snd = ex::just() // sender
      | ex::let_value([data = std::move(data)]() { //
        return handle_connection(data);
      });
    ex::start_detached(std::move(snd)); // perform on a different thread
  }
  co_return true;
}
```

the high level logic for each connection

```cpp
auto handle_connection(const conn_data& cdata) {
  return read_http_request(cdata.io_ctx_, cdata.conn_)
    | ex::transfer(cdata.pool_.get_scheduler())
    | ex::let_value([&cdata](http_server::http_request req) {
      return handle_request(cdata, std::move(req));
    })
    | ex::let_error([](std::exception_ptr) { return just_500_response(); }) // when error or stop signal
    | ex::let_stopped([]() { return just_500_response(); }) // when error or stop signal
    | ex::let_value([&cdata](http_server::http_response r) {
      return write_http_response(cdata.io_ctx_, cdata.conn_, std::move(r));
    });
}
```

a sender that creates 500 responses

```cpp
auto just_500_response() {
 auto resp = http_server::create_response(http_server::status_code::s_500_internal_server_error);
 return ex::just(std::move(resp));
}
```

the reading data coroutine

```cpp
auto read_http_request(io::io_context& ctx, const io::connection& conn) -> task<http_server::http_request> {
  http_server::request_parser parser;
  std::string buf;
  buf.reserve(1024 * 1024);
  io::out_buffer out_buf{buf};
  while (true) {
    std::size_t n = co_await io::async_read(ctx, conn, out_buf);
    auto data = std::string_view{buf.data(), n};
    auto r = parser.parse_next_packet(data);
    if (r) {
      co_return {std::move(r.value())};
    }
  }
}
```

the writing response coroutine

```cpp
auto write_http_response(io::io_context& ctx, const io::connection& conn, http_server::http_response resp) -> task<std::size_t> {
  std::vector<std::string_view> out_buffers;
  http_server::to_buffers(resp, out_buffers);
  std::size_t bytes_written{0};
  for (auto buf : out_buffers) {
    while (!buf.empty()) {
      auto n = co_await io::async_write(ctx, conn, buf);
      bytes_written += n;
      buf = buf.substr(n);
    }
  }
  co_return bytes_written;
}
```

directing the request to a proper handler (image filter).

```cpp
auto handle_request(const conn_data& cdata, http_server::http_request req) -> task<http_server::http_response> {
  auto puri = parse_uri(req.uri_);
  if (puri.path_ == "/transform/blur")
    co_return handle_blur(cdata, std::move(req), puri);
  else if (puri.path_ == "/transform/adapt_thresh")
    co_return handle_adapt_thresh(cdata, std::move(req), puri);
  else if (puri.path_ == "/transform/reduce_colors")
    co_return handle_reduce_colors(cdata, std::move(req), puri);
  else if (puri.path_ == "/transform/cartoonify")
    co_return co_await handle_cartoonify(cdata, std::move(req), puri);
  else if (puri.path_ == "/transform/oil_painting")
    co_return handle_oil_painting(cdata, std::move(req), puri);
  else if (puri.path_ == "/transform/contour_paint")
    co_return co_await handle_contour_paint(cdata, std::move(req), puri);
  co_return http_server::create_response(http_server::status_code::s_404_not_found);
}
```

using `transfer_just` which ensures the work switches threads, but doesn't move into a different execution process.

```cpp
auto handle_cartoonify(const conn_data& cdata, http_server::http_request&& req, parsed_uri puri) -> task<http_server::http_response> {
  int blur_size = get_param_int(puri, "blur_size", 3);
  int num_colors = get_param_int(puri, "num_colors", 5);
  int block_size = get_param_int(puri, "block_size", 5);
  int diff = get_param_int(puri, "diff", 5);
  auto src = to_cv(req.body_);
  ex::sender auto snd = ex::when_all( // parallel work
    ex::transfer_just(cdata.pool_.get_scheduler(), src)
    | ex::then([=](const cv::Mat& src) {
      auto gray = tr_to_grayscale(tr_blur(src, blur_size));
      return tr_adapt_thresh(gray, block_size, diff);
    }),
    ex::transfer_just(cdata.pool_.get_scheduler(), src)
    | ex::then([=](const cv::Mat& src) {
      return tr_reduce_colors(src, num_colors);
    })
    )
  | ex::then([](const cv::Mat& edges, const cv::Mat& reduced_colors) {
    return tr_apply_mask(reduced_colors, edges);
  })
  | ex::then(img_to_response);
 co_return co_await std::move(snd);
}
```

#### Conclusions

this approach allows for a visual representation of concurrent work.

> senders
>
> - good abstractions for concurrency
> - no need for synchronization
> - highly composable

</details>

### C++ Coroutines, from Scratch - Phil Nash

<details>
<summary>
Building coroutines backwards from a complicated example/
</summary>

[C++ Coroutines, from Scratch](https://youtu.be/EGqz7vmoKco)

starting with a real world example of coroutines.

Financial data with a complex hierarchy of types which hold pointers to polymorphic types, we want to get the data from json format. deserializing data with dependencies (recursively calling the deserialization on nested types). the main bottleneck will be calling the database to get the data, so we will want to batch the calls to the database and request data by batches according to the layers (levels).

the early versions of the example use callbacks and continuations, which start at complex, and grow at complexity at a staggering rate as we add customization points and management state. the more we move the continuation code into a framework, the more it turns out to look like coroutines.

coroutines:

- User
  - promise_type
  - Task/Generator
  - Awaiter
- Framework
  - _std::coroutine_handle_
  - _std::coroutine_traits_
  - _std::noop_coroutine_promise_
  - _std::suspend_always_
  - _std::suspend_never_

coroutines are always allocated on the heap, but they have a stable address. we don't need to move objects as much, because the life time is managed and remains throughout suspension points.

(going over code)

</details>

### Sockets - Applying the Unix Readiness Model When Composing Concurrent Operations in C++ - Filipp Gelman

<details>
<summary>
Understanding blocking code, *select* and pipes.
</summary>

[Sockets - Applying the Unix Readiness Model When Composing Concurrent Operations in C++](https://youtu.be/YmjZ_052pyY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/cppcon2022.pdf)

#### What I learned From Sockets

concurrent operation involve waiting

1. setup
2. wait
3. react

so the question is what to do while waiting, suspend? yield control? busy wait?

we can either be waiting for a specific operation, all operations, or any one operation.

breaking code into regular code and code that waits, it's better to wait for several things, and react to the thing that has happened.

> "_Don't communicate by sharing memory. Share memory by communicating_"

> - How do i `.get()` the first of several futures?\
> - How can I `co_await` the first of several awaitables
> - How can i `select` several senders?

#### Introduction to Sockets

- reading and writing
- sockets
- blocking vs non-blocking

readers return an integer, which has different dimensions of meaning.

```cpp
char buff[1024];
int result = read(fd, buffer,1024);
if (result > 0){
  // read this many bytes
} else if (result == 0){
  // end of file
} else {
  // error, check errno
}
```

writers are similar, with 0 bytes written meaning some form of being out of memory.

```cpp
int sock = socket(AF_INET, SOCK_STREAM, 0); // create socket

sockaddr_in addr{
  .sin_family = AF_INET,
  .sin_port = htons(80),
  .sin_addr = {s.addr = /* cidr */},
  .sin_zero = {}
}; // create address

connect(sock, &addr, sizeof(adr)); // connect socket to address
```

we create a socket, link it to an address, and now we can treat the socket like a file. but for this "file", operations can be blocking. a socket can be conceptualized as having 'read' and 'write' buffers, and if we try reading an empty socket, or write to a full socket, then the operation will block.

we could also have the socket as non blocking,

```cpp
int sock = socket(AF_INET, SOCK_STREAM | SOCK_NONBLOCK, 0); // create non-blocking socket

//.. connect to address

char buff[1024];
int result = read(socket, buffer,1024);
if (result > 0){
  // read this many bytes
} else if (result == 0){
  // end of "file " - other side done writing
} else if (errno == EAGAIN){
  // no data yet
} else {
  // error, check errno
}
```

now we don't block, and we can have several sockets, and read from either of them. this isn't blocking, but we're **busy waiting** and burning cpu.

#### Select

Instead of that busy waiting, we can use `select`, we have a set of file descriptors (sockets) that we want to listen to. and we sleep on the operation, and we will wake up when something happens, and the set will be filled with the descriptors that something happened on.

```cpp
while (true){
  fd_set fds;
  FD_ZERO(&fds);
  FD_SET(sock, &fds);
  FD_SET(sock2, &fds);

  // wait
  select(FD_SETSIZE, &fds, nullptr, nullptr, nullptr);

  // react

  if (FD_ISSET(sock, &fds)){
    int result = read(sock, buffer, 1024);
    // handle result
  }

  if (FD_ISSET(sock2, &fds)){
    int result = read(sock2, buffer, 1024);
    // handle result
  }
}
```

there are some limitations to this basic implementations, the number of descriptors is limited (compile time constant), and transferring them in after each operation can become costly.

this is achieved with Epoll

```cpp
int ep_fd = epoll_create1(0);
epoll_event ep_events[2] = {
  epoll_event{
    .events = EPOLLIN,
    .data = epoll_data_t{.fd = sock}},
  epoll_event{
    .events = EPOLLIN,
    .data = epoll_data_t{.fd = sock2}}
  };

epoll_ctl(ep_fd, EPOLL_CTL_ADD, sock, ep_events + 0);
epoll_ctl(ep_fd, EPOLL_CTL_ADD, sock2, ep_events + 1);

while (true) {
  epoll_event evt;
  // wait
  epoll_wait(ep_fd, &evt, 1, -1);

  // react
  if (evt.data.fd == sock) {
  int result = read(sock, buffer, 1024);
  // handle result
  } else if (evt.data.fd == sock2) {
  int result = read(sock2, buffer, 1024);
  // handle result
  }
}
```

unix readiness model

> - Perform initial setup
> - `while (true)`
> - **Wait** for events (blocking).
> - **React** to events (non-blocking).
> - On completion or error: `break;`
> - `close()`

when we establish a connection, we convert our address to an IP address using the DNS. then we have the handshake stage, and then we have the connection. but in the modern world, we get more than one address from the DNS, so we have several ip address to try from, and we only need one.

- use them in sequence
- connect to all of them
- "Happy eyeballs"
  - start connecting to one
  - wait for a short while, if no response. try another.
  - when either of them comes back, use that connection and send a RESET to the discarded connections.

```cpp
int connect(vector<sockaddr_in> addrs){
  // establish connection to an address in addrs
  int ep_fd = epoll_create1(0);
  vector<epoll_event> events(addrs.size());

  for (sockaddr_in addr : addrs) {
    int sock = socket(AF_INET, SOCK_STREAM | SOCK_NONBLOCK, 0);
    connect(sock, &addr, sizeof(addr));

    epoll_event evt{
      .events = EPOLLOUT | EPOLLHUP,
      .data = epoll_data_t{.fd = sock}};

    epoll_ctl(ep_fd, EPOLL_CTL_ADD, sock, &evt);

    // wait
    int result = epoll_wait(ep_fd, events.data(), events.size(), 250);

    // react - to any event that has happened, even if it was in previous iteration of the outer loop.
    for (int i = 0; i < result; ++i) {
      if (events[i].revents == EPOLLOUT) {
        // connection established
        return events[i].data.fd;
      } else {
        // connection failed
        epoll_ctl(ep_fd, EPOLL_CTL_DEL, events[i].data.fd, nullptr);
      }
    }
  }

  return -1;
}
```

this has the problem of that the function must complete all the addresses before finishing, and can't be stopped from outside. we get around this by passing an socket to the function.

we add the out socket, and register it to wake up when the socket is closed. so now we can use this "event" as the trigger for canceling the work, or when a different connection is established, we push the file descriptor into the socket and handle the read from it elsewhere.

#### Implementations in C++

so far, we worked on C code (more or less), now we move to C++ code. we have a list of addresses, we try them one after another, and if one fails, we remove it from the set and try the next. we can be cancelled via the socket, and we return the result back through it. if all connections failed, we close the socket without writing to it, thus signaling a problem.

```cpp
void connect(std::queue<sockaddr_in> addrs, WriteHandle<Socket> out) {
  Select select;
  select.insert(out, Events::hup);

  Timer next_connection;
  next_connection.set(clock::now());
  select.insert(next_connection, Events::io);

  std:::set<Socket> connections;

  while (true) {
    if (result.handle == out) {
      return; // cancellation
    }
    else if (auto iter = connections.find(result.handle); iter != connections.end()) {
      // check connection status
      if (result.event & Events::io) {
        out.write(*iter);  // write to socket
        return;
      } else {
        select.erase(result.handle); // connection failed
        connections.erase(iter);
        goto next_connection; // using goto to create a new connection
      }
    } else {
      // start next connection
      next_connection:
      if (!addrs.empty()) {
        Socket sock = connect_to(addrs.front());
        addrs.pop_front();
        select.insert(sock, Events::io | Events::hup);
        delay.set(clock::now() + 250ms);
      } else if (connections.empty()) {
        return; // we are out of connections to try
      }
    }
  }
}
```

we can make this into a coroutine.

pipe are the fundamental type of sockets, it's a buffer we can read and write from.

| result                           | C read | C write | C++ Read        | C++ write       |
| -------------------------------- | ------ | ------- | --------------- | --------------- |
| Type                             | int    | int     | variant         | Variant         |
| Success - read one byte          | 1      | 1       | `Success<char>` | `Success<void>` |
| EOF - end of file                | 0      | 0       | `EndOfFile`     | `EndOfFile`     |
| Blocking - operation would block | −1     | -1      | `WouldBlock`    | `WouldBlock`    |

```cpp
template <typename T>
struct Success { T value; };

template <>
struct Success<void> {};

struct EndOfFile {};
struct WouldBlock {};

template <typename T>
using Result= variant<Success<T>, EndOfFile, WouldBlock>;
```

we have a shared state, and we can employ many readers and writers.

this is an example with busy waiting

```cpp
void capitalize_busy_wait(ReadHandle in, WriteHandle out) {
  while (true) {
    Result<char> input = in.read();

    if (input == EndOfFile{}) return; // socket closed
    if (input == WouldBlock{}) continue; // nothing to read now

    char capital = toupper(get<Success<char>>(input).value); // transform value

    while (true) {
      Result<void> output = out.write(capital); // attempt to write
      if (output == EndOfFile{}) return; // socket is closed
      if (output == WouldBlock{}) continue; // socket is full, try again
      break; // go back to reading
    }
  }
}
```

we have handles (with a shared pointer to the shared state). now we want to have no blocking, and no spinning versions of the operations. the handles call the relevant function from the shared state object.

```cpp
struct SharedState {
  optional<char> buffer;
  bool closed{false};
  mutex mutex;
  atomic<unsigned> num_readers{0}, num_writers{0};
  Result<char> read();
  Result<void> write(char value);
  void close();
};


Result<char> SharedState::read() {
  scoped_lock lock(mutex); // RAII
  if (buffer) return Success{*exchange(buffer, nullopt)};
  if (closed) return EndOfFile{};
  return WouldBlock{};
}

Result<char> SharedState::write(char value) {
  scoped_lock lock(mutex); // RAII
  if (closed) return EndOfFile{};
  if (buffer) return WouldBlock{};
  buffer.emplace(value);
  return Success<void>{};
  }

void SharedState::close() {
  scoped_lock lock(mutex);
  closed = true;
}
```

the `select` function:

> - **One** caller waits for events from **many** handles.
> - **One** shared state notifies **many** callers when events occur.
> - Many-to-Many relationship

the _select_ object contains all the relationship pairs, and handles notifying the correct handles for each events they're subscribed for. we can use an intrusive (non owing) list to simplify the implementation by having only a single copy of the connections.

(going over all the function needed and how the objects relate to one another)

#### Senders, Receivers and Coroutines

combining senders, receivers, sockets and coroutines.

- Caller
- Select
- Operation State
  - `subscribe()`
  - `unsubscribe()`
  - `read()`
  - `write()`
  - `close()`
  - /\_/\_OP_State
    - `request_stop()`
  - Receiver
    - `set_value()`
    - `set_stopped()`

this isn't like _any_ - we don't throw away the rest of the connections.\
We can make the _select_ a sender by itself, as it already hasa conditional variable object inside it. we can also make it into a coroutine.

</details>

### Architecting Multithreaded Robotics Applications in C++ - Arian Ajdari

<details>
<summary>
LLAMBA - multi-threading in robotics
</summary>

[Architecting Multithreaded Robotics Applications in C++](https://youtu.be/Wzkl_FugMc0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/08/Presentation_Arian_Ajdari_C_21.08.2022.pptx)

> "Developing impactful software requires _more than_ writing fast, concise, and correct code"

> Software Architecture:\
> Software Architecture in a nutshell presents:
>
> - A blueprint
> - Organisation of the software systems
> - A process which leads from idea, to realization of the software, and its maintenance
> - providers a structure to ensure quality properties have been fulfilled

stages of software architecting

- idea
- system requirements
- architecture drivers
- architecture solutions
- implementations
- maintainence

the Software architect interacts with many internal and external stakeholders in the company, developers, customers, marketing, operations, support teams, etc..

#### Context and the problem

in terms of robotics, the common c++ frameworks is ROS2 (robot operating system), an open source solution that uses the publisher/subscriber pattern.

robotics engineers understand hardware, and want to have multi-threading for performance benefits. on the other hand, software developers know that multi-threading is hard and can have many errors. it's up to the architect to design the software architecture that can support it.

pthread is too low-level and the standard threading library is too abstract for our case, so we need a middleware - LLAMBA.

#### LLAMBA

a tool for robotics engineers for multi-threaded applications. schedulers, affinities, parallelization approach.

python bindings.

</details>

### Smarter Cpp Atomic Smart Pointers - Efficient Concurrent Memory Management - Daniel Anderson

<details>
<summary>
trying to get a performant shared pointer for concurrent memory management.
</summary>

[Smarter Cpp Atomic Smart Pointers - Efficient Concurrent Memory Management](https://youtu.be/OS7Asaa6zmY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/cppcon2022-smarter-atomic-smart-pointers.pptx), [github](https://github.com/cmuparlay/concurrent_deferred_rc).

#### Introduction

smart pointers are c++ answer for garbage collection, we avoid using raw pointers.

| pointer type               | description                                  |
| -------------------------- | -------------------------------------------- |
| <cpp>std::unique_ptr</cpp> | exclusive ownership, move-only, not copyable |
| <cpp>std::shared_ptr</cpp> | shared ownership, reference counting         |
| <cpp>std::weak_ptr</cpp>   | a non-owning reference to a shared resource  |

we are mostly concerned with shared ownership.

**The concurrent memory management problem**

a linked list, thread 1 has an iterator at node B, but while thread 1 is asleep, thread 2 deletes the node.

> A thread might want to delete something that another thread is still reading.

#### Existing Management techniques

##### Atomic smart pointers - easy to use

since c++20, there are atomic smart pointers. in shared pointers, the reference counting is managed atomically, but it's not thread safe to copy a them between threads. <cpp>atomic\<shared_ptr></cpp> hold a shared pointer that can be manipulated in an atomic matter.

- store
- load
- compare_exchange_weak

this makes memory safe easier, but it's not lock-free. and it's really slow. it uses locks and invalidate cache.

##### Hazard pointers and Read-Copy-Update(RCU) - high performance

> **Deferred reclamation**: if someone is reading something that you want to delete, defer it until later.
>
> Deferred reclamation schemes differ in:
>
> - How to determine whether something might currently be being read
> - How and when they perform deferred deletion (synchronously or asynchronously, same thread, different thread)
>
> Replace delete with retire
>
> - retire: delete this object after any existing readers are finished
> - Usage is extremely subtle, and bugs are common

it's very easy to get this wrong.

> **Hazard Pointers**
>
> - Threads announce particular pointers that they plan to read (“protect”)
> - Periodically clean up retired objects, but defer those with an active hazard pointer
>
> Limitations/pitfalls:
>
> - Often, for searches, the operation must restart if an update touches it
> - This means that hazard pointers do not work with every data structure.
> - Traversals that require multiple hazard pointers (e.g., hand-over-hand) require the user to keep the correct ones active

because we don't change the elements of the data structure, we don't mess up caches of other threads, so it should be better the atomic shared pointers.

but it's hard to get the code right.

> **RCU**:
>
> - Threads announce their intention to **read the data structure**
> - Retired objects cleaned up synchronously or asynchronously
>
> Comparison to hazard pointers
>
> - Protection is coarse grained
> - Throughput is higher
> - Memory overhead is higher
> - Easier to use

Hazard pointers protect a memory location, while RCU protects the entire data structure. the performance of RCU is nearly the same as the baseline, with slight overhead and larger memory garbage.

but it's still hard to use. we want something which is as easy to use as atomic smart pointers, but still high performant.

#### Possible Solutions

combining RAII and deferred reclamation.

there is a proposal for <cpp>std::cell, std::latest, std::snapshot_src</cpp> which should support this, it has a guaranteed lifetime without taking ownership. it will only work with unique ownership.

but maybe we could combine it with an atomic shared pointer?

hazard pointer:

```cpp
struct Node : public hazard_pointer_obj_base<Node> {
  T value; atomic<Node*> next;
};
atomic<Node*> head;
// Single (or synchronized) writer
void remove(Node* prev, Node* target) {
  prev->next.store(target->next.load());
  target->next.store(nullptr); // this has to be here!
  target->retire(); // this has to be after the previous line
};

// May be called by multiple concurrent readers
bool find(const T& val) {
   hazard_pointer hPrev = make_hazard_pointer();
   hazard_pointer hCurr = make_hazard_pointer(); while (true) {
    atomic<Node*>* prev = &head_;
    atomic<Node*> curr = prev->load();
    while (true) {
      if (!curr) return false;
      if (!hCurr->try_protect(cur, *prev)) break;auto next = cur->next.load();
      if (prev->load() != cur) break;
      if (cur->value == val) return true;
      swap(hCurr, hPrev);
      prev = &(cur->next);
      cur = next;
    }
  }
}
```

deferred smart shared pointer

```cpp
struct Node {
  T value;
  deferred_shared_ptr<Node> next;
};

deferred_shared_ptr<Node> head;
// Single (or synchronized) writer
void remove(Node* prev, Node* target) {
  prev->next.store(target->next.load());
};

// May be called by multiple concurrent readers
bool find(const T& val) {
  snapshot_ptr<Node> curr = head.get_snapshot();
  while (true) {
    if (!curr) return false;
    if (curr->value == val) return true;
    curr = curr->next.get_snapshot();
  }
}
```

for remove: it's hard to explain why those extra lines must exist, and why the order matters so much. and for the find algorithm, there is a retry loop (a loop inside a loop), in case that someone protects the node, or if someone removed the next node already. it's very easy to write bugged code.

the solution uses the deferred reclamation to manage the reference count, this protects both the reference count and the object itself.

implementation

```cpp
std::atomic<control_block*> ptr;

shared_ptr<T> load() {
  // No atomic op can decrement the ref count
  auto protected_ptr = protect(&ptr);
  protected_ptr->increment_ref_count();
  return shared_ptr<T>(protected_ptr);
} // protection is released

void store(shared_ptr<T> desired) {
  auto new_ptr = desired.release();
  auto old_ptr = ptr.exchange(new_ptr);
  if (old_ptr != nullptr) retire(old_ptr); // decrement the reference count
}

snapshot_ptr<T> get_snapshot() {
  auto protected_ptr = protect(&ptr);
  return snapshot_ptr<T>(move(protected_ptr));
}

bool compare_exchange_weak(shared_ptr<T>& expected, shared_ptr<T>&& desired) {
  control_block* exp = expected.ctrl_block();
  control_block* des = desired.ctrl_block();
  if (ptr.compare_exchange_strong(exp,des)) {
    if (exp != nullptr) retire(exp);
      desired.release();
  }
  else {
    expected = load();
    return false;
  }
}
```

we can implement it with either Hazard pointers or RCU, and we get minimal overhead but it's trivial to use.

> Fun facts
>
> - Only single-word CAS
> - No bit-stealing / pointer packing
> - We also support weak pointers!

</details>

### Scalable and Low Latency Lock-Free Data Structures - Alexander Krizhanovsky

<details>
<summary>
Challenges for low latency data structures in a scalable environment.
</summary>

[Scalable and Low Latency Lock-Free Data Structures](https://youtu.be/j_FCgQmgp_M), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/scalable_and_low_latency_lock-free_data_structures.pdf), [Tempesta github](https://github.com/tempesta-tech/tempesta), [TempestaDB](https://github.com/tempesta-tech/blog/tree/master/htrie).

designing a data structure to handle requests with the Tempesta web content delivery system CDN.

this data structure acts as a database and provides CRUD operations, and it should be lock-free with low latency

lock-free isn't the same as wait-free. lock-free is a "weaker" requirement, it guarantees system progress and finite step for a task, but not the system-wide throughput and deterministic steps for all operations.

lock-free deletions are hard: we can have other nodes accessing the deleted node via helpers, and we need to protect against concurrent deletions. we also need to think about memory fragmentation and garbage collections.

some solutions:

- upper layer responsibility (e.g. reference counting)
- RCU - read copy update
- Dummy nodes (split-ordered lost, skip trees)

#### Tempesta DB

> - Is part of Tempesta FW (a hybrid of a firewall and web-accelerator)
> - Linux kernel space (softirq – deferred interrupt context)
> - Can be concurrently accessed by many CPUs
> - In-memory database
> - Simple persistence by dumping `mmap()`’ed areas => offsets instead of pointers
> - Duplicate key entries (stale web responses)
> - Multiple indexes (e.g. URL or Vary for web cache)

storing data with keys of type string, can be large keys. has large records, with variable size, can be large or very large. also has some fixed sized small records, such as client accounts, session cookies, filter rules and network masks.

trie trees have ordering and range queries, hash tables use fast point queries and they need rehashing, which is bad for tail latency.

memory allocation:

- storing index blocks and memory blocks apart.
- index blocks should be compressed and kept together (spatial locality).

#### Data Structures

_binary trees_ (<cpp>std::map</cpp>) require rebalancing - rotations involving many nodes, hard to implement lock free or even fine-grained locking.

_hash tables_ (<cpp>std::unordered_map</cpp>) have buckets chains that may grow infinitely, rehashing takes time and requires a global lock - this effects tail latency. it is easy to implement granular locks per bucket.

_split-ordered lists_ (<cpp>tbb::concurrent_unordered_map</cpp>, part of **MariaDB**) use persistent dummy nodes, which cause memory degradation after removing a node or requires a lock for removing.

_Radix/patricia tree_ (trie)

> - not cache conscious
> - hard to make concurrent
> - Height depends on the key length (fixed height)
> - No reconstruction (e.g. rebalancing or rehashing)
> - Memory greedy on uniformly distributed keys in a large space
> - Easy to make lock-free

path compression - instead of having a "node" for each character, we can compress paths when there is only a single child.

_Burst tries_ - combine trie with buckets

> - Collapses trie chains into buckets – better memory usage
>   1. use only string prefixes in a trie
>   2. resolve collisions by suffixes in a small bucket
>   3. once the bucket inefficient (several heuristics) – burst it
> - Adaptive: e.g. buckets with rare hits do not burst
> - Poor performance on the start

(some words about caches and memory ordering in x86-64 machines) cache conscious data structures access L1-L3 data cache.

#### Implementation

starting a Burst Hash Trie as a root. each node holds an array of shift (child nodes)

```cpp
const size_t HTRIE_BITS = 4;
const size_t HTRIE_FANOUT = 1 << HTRIE_BITS;
// 2 ^ 31 * 64bytes = 128GB
const size_t HTRIE_DBIT = 1 << (sizeof(int) * 8 – 1);
struct HtrieNode {
 // 16 * 4 = 64 = 1 cache line on x86-64
 unsigned int shifts[HTRIE_FANOUT];
};
```

buckets contain the collision chain as a bitmap, and are a fixed size (contain only two members: integer and long). we create them outside of chain, so there is no need to lock it, inserting the bucket is done via an atomic operations. buckets can burst into chains there are multiple records at the node location.

(more stuff I don't follow)

</details>

##

[Main](README.md)
