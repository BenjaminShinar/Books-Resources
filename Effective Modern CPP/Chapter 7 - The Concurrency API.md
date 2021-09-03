## Chapter 7 - The Concurrency API

<summary>
</summary>

C++11 is the first time that the language has integrated support for threads and concurrency, until then, multi-threading was achieved with libraries like _pthread_ or windows threads. but now, there is finally a common set of building blocks in the standard library: tasks, futures and shared_futures, threads, mutex, conditional variables, atomic objects and others.

### Item 35: Prefer Task-Based Programming to Thread-Based

<details>
<summary>
The higher level of abstraction for std::async means we are less likely to run into surprise issues and we have better ways to deal with them.
</summary>

assume we have a function we want to run asynchronously, we can either run it as a thread or as a task.

```cpp
int doAsyncWork();
std::thread t(doAsyncWork); // thread approach, join it somewhere
auto fut = std::async(doAsyncWork); //task approach, return future;
```

the task based approach is better. for start,it's clear that we have a return value. with the thread approach, we need to remember to retrieve the return value from the thread eventually. another advantage is that if the function throws an exception, we can handle it with the task approach, but we might be terminated completely with the thread approach.

tasks represent a higher level of abstraction than threads,it means we aren't concerned with the details of thread management.
there are actually three meanings to 'threads'. hardware, software and c++ objects.

> 1.  _Hardware threads_ are the threads that actually perform computation. Contemporary machine architectures offer one or more hardware threads per CPU core.
> 2.  _Software threads_ (also known as OS threads or system threads) are the threads
>     that the operating system manages across all processes and schedules for execution on hardware threads. It’s typically possible to create more software threads than hardware threads, because when a software thread is blocked (e.g., on I/O or waiting for a mutex or condition variable), throughput can be improved by executing other, unblocked, threads.
> 3.  _std::threads_ are objects in a C++ process that act as handles to underlying
>     software threads. Some _std::thread_ objects represent “null” handles, i.e., correspond to no software thread, because they’re in a default-constructed state(hence have no function to execute), have been moved from (the moved-to _std::thread_ then acts as the handle to the underlying software thread), have been joined (the function they were to run has finished), or have been detached (the connection between them and their underlying software thread has been severed).

software threads are a limited results, if we try to create too many of them we get a _std::system_error_ exception. trying to program around the possibility of that issue is a hassle by itself. even if we can create software threads, we can encounter the issue of _oversubscription_, which means there are more software threads ready to run than hardware threads, so we start the process of time slicing, thread-scheduler and context switches. this all means overhead on the performance of the system. there is also the possibility of losing cache locality if a software thread is scheduled to run on a different core.

avoiding the oversubscription issue is difficult, there is no constant optimal ratio of how many software threads we should use, as it depends on the kind of work we do, on the hardware, on the memory caches and machine architecture.

it's better to avoid thinking about all that, and leave the decision to something else, which is what _std::async_ does. all the issues with oversubscription are handled by the standard library implementation. _std::async_ with the default launch policy is able to decide not to create a new thread and run it on the current thread. there is still a possibility of running into load-balancing issues, but the odds of that happening are lower.

```cpp
auto fut = std::async(doAsyncWork); //task approach, return future;
```

if we have a GUI thread, _std::async_ might be a problematic, but in this case we can decide to explicitly pass the _std::launch::async_ policy to request the task to run on a different thread. programming with _std::async_ takes advantage of what the OS can give us in terms of load balancing across cores and work stealing. managing threads ourselves means we use less of what the machine can do, and we have to program around all the issues ourselves.

we might still want to use *std:::thread*s directly

- using the underlying API of the thread implementation. if we want some platform specific behavior, like windows threads or *pthread*s functions.
- Optimizing thread usage for the application. if we know what machine we will run on and we know the execution profile.
- Abilities that the current standard library doesn't provide. c++ still doesn't have thread pools.

#### Things to Remember

> - The std::thread API offers no direct way to get return values from asynchronously run functions, and if those functions throw, the program is terminated.
> - Thread-based programming calls for manual management of thread exhaustion, oversubscription, load balancing, and adaptation to new platforms.
> - Task-based programming via std::async with the default launch policy handles most of these issues for you.

</details>

### Item 36: Specify _std::launch::async_ if Asynchronicity is Essential

<details>
<summary>
The default behavior or std::async can be surprising. There are cases where it should be avoided.
</summary>

calling _std::async_ isn't just calling for running the command asynchronously, it calling the function to run according to some launch policy from the _std::launch_ enum.

- _std::launch::async_ - the command should run asynchronously.
- _std::launch::deferred_ - the command should only run when _get_ or _wait_ is called on the resulting _std::future_ (or _std::shared_future_), only then we run the command (blocking the other thread). nothing happens until then, and if the result is never needed, the function won't run.
- _std::launch::async_ | _std::launch::deferred_ - **the default behavior**. this gives the standard library implementation and the compiler the most freedom to control and optimize threading behavior.

```cpp
auto fut1 = std::async(f); //default
auto fut2 = std::async(std::launch::async | std::launch::deferred,f); //explicit default
```

this default behavior means:

> - **It’s not possible to predict whether f will run concurrently with** t, because f
>   might be scheduled to run deferred.
> - **It’s not possible to predict whether f runs on a thread different from the
>   thread invoking get or wait on fut**. If that thread is t, the implication is that
>   it’s not possible to predict whether f runs on a thread different from t.
> - **It may not be possible to predict whether f runs at all**, because it may not be
>   possible to guarantee that get or wait will be called on fut along every path
>   through the program.

this weird behavior and unpredictably doesn't work well the _thread-local storage_ (_thread_local_ variables), as it's not possible to say which thread local variables will be accessed.

it also means wait-based loops with timeout don't work well, because calling _.wait_for()_ and _.wait_until()_ on tasks that were deferred return the value _std::launch::deferred_. so the following code with a loop doesn't terminate.

```cpp
using namespace std::literals; // for duration suffixes
void f()
{
    std::this_thread::sleep_for(1s);
}
auto fut = std::async(f); //default policy
while (fut.wait_for(100ms) != std::future_status::ready) //loop until task is done
{
    //...
}
```

if the f function runs concurrently (as if _std::launch::async_ was used), then the code will execute properly and the while loop will terminate. but if the function is deferred (_std::launch::deferred_) then at no point the function is called, so the loop will never terminate. this behavior might be tricky to find and will occur only under certain loads conditions.

unfortunately, there is no simple way to tell upfront what was the policy the task was launched with. so we need a roundabout way to do this.as we saw with the _.wait_for()_ method.

```cpp
auto fut = std::async(f);

if (fut.wait_for(0s) == std::future::deferred)
{
    //either call wait or get to call f synchronously
}
else
{
    while (fut.wait_for(100ms) != std::future_status::ready) //loop until task is done
    {
    //... do some work until task is done.
    }
}
//fut is ready
```

using _std::async_ with the default policy is safe under the following conditions:

> - The task need not run concurrently with the thread calling get or wait.
> - It doesn’t matter which thread’s thread_local variables are read or written.
> - Either there’s a guarantee that get or wait will be called on the future returned by _std::async_ or it’s acceptable that the task may never execute.
>   Code using _wait_for_ or _wait_until_ takes the possibility of deferred status into
>   account.

if those conditions aren't guaranteed, it's probably better to call the task with true asynchronicity, but being explicit with the launch policy. this code shows a perfect forwarding utility function in c++11 and c++14 syntax.

```cpp
auto fut = std::async(std::launch::async); // must run concurrently

template <typename F, typename... Ts>
inline
std::future<typename std::result_of<F(TS...)>::type> //get return type
reallyAsync11(F&& f,Ts&&... params)
{
    return std::async(std::launch::async, std::forward<F>(f),std::forward<Ts>(params)...);
}

template <typename F, typename... Ts>
inline
auto reallyAsync14(F&& f,Ts&&... params) // auto return type
{
    return std::async(std::launch::async, std::forward<F>(f),std::forward<Ts>(params)...);
}
```

#### Things to Remember

> - The default launch policy for _std::async_ permits both asynchronous and
>   synchronous task execution.
> - This flexibility leads to uncertainty when accessing thread_locals, implies
>   that the task may never execute, and affects program logic for timeout-based
>   wait calls.
> - Specify _std::launch::async_ if asynchronous task execution is essential.

</details>

### Item 37: Make *std::thread*s Un-Joinable on all Paths

<details>
<summary>
The destructor of std::thread causes program termination when called on a joinable object, we should not let this happen.
</summary>

a _std::thread_ can either be _joinable_ or _un-joinable_. a _joinable_ _std::thread_ can be running or in a state that could be changed to running. this means that a _std::thread_ in a blocking or waiting state is _joinable_. in contrast, _un-joinable_ threads include

- _default constructed std::threads_. as they have no function to execute and they don't correspond to an underlying thread of execution.
- _std::thread_ objects that have been move from.
- _std::thread_ objects that have been joined.
- _std::thread_ objects that have been detached.

**if the destructor for a _joinable_ thread is invoked, this leads to a program termination.**

in this example, we have a function that performs a computation (a callback function) on a subset of some values based on a filtering condition. we might want to do both action concurrently. in this example we will use _std::thread_ rather than a task-based approach.
this code has problems.

```cpp
constexpr auto tenMillion = 10000000;
constexpr auto tenMillionC14 = 10'000'000; //c++14 ives as digit separators
bool doWork(std::function<bool(int)> filter, int maxVal = tenMillion) //default parameter
{
std::vector<int> goodValues; // container for values that match the criteria.

std::thread t( //create a thread with the lambda function
    [&filter, maxVal,&goodValues]
    {
    for (auto i = 0; i<= maxVal;++i) //probably better to use std::copy_if
        {
            if (filter(i))
            {
                goodValues.push_back(i)
            }
        }
    });
auto nativeHandler  = t.native_handle(); // we want to set the thread priority
//...

if (conditionsAreSatisfied())
{
    t.join();
    performComputation()
    return true;
}
return false;
}
```

a first issue is that we change code priority after it started running, but a more serious issue is the _'conditionsAreSatisfied()'_ call, if it returns false or throws an exception, the _std::thread_ destructor will be called without it having been joined, and the entire program will be terminated.

while this behavior (terminate the process because of an unjoined thread) seems harsh, the alternatives were considered worse.

- **Implicit join** - call _.join()_ inside the destructor. this could lead to performance anomalies and weird behavior at a very critical spot.
- **Implicit detach** - call the _.detach()_ method inside the destructor. this can lead to situations where the detached function tries to change some local variable that has been captured and is now no longer used. a dangling reference issue that is horrible to detect and deal with.

because of how utterly annoying and confusing programs could be on either of those behavior, the committee decided that both options are bad, and that it's better to entirely terminate the program in this case.

this means that the programmer should ensure that no matter the path the code takes, no _std::thread_ object should reach the end of the scope while in a _joinable_ condition. this means all return paths on all branches, and also the case of an exception thrown from the scope (even if it's handled somewhere else). this behavior is usually delegated into an _RAII_ objet (_RAII_ stands for _'resource acquisition is initialization'_, despite the focus being destruction), but that's the exact same problem the regular object has. there is no good solution (implicit join or detach) that can be used.

As programmers, of course, we can write such a class himself and use it as we see fit.

```cpp
class ThreadRAII
{
    public:
    enum class DtorAction{join,detach};
    explicit ThreadRAII(std::thread&&t, DtorAction a):action(a),t(std::move(t))
    {}
    ~ThreadRAII()
    {
        if (t.joinable())
        {
            if (action == DtorAction::join)
            {
                t.join();
            }
            else
            {
                t.detach();
            }
        }
    }
    std::thread& get() {return t;}
    private:
    DtorAction action;
    std::thread t;
};
```

- The constructor accepts only rvalues, because we take the _std::thread_ object and move from it, it's an un-copyable class.
- The parameter order is switched from that of the initialization order (which mimics the class structure). the switch is for the sake of the users, as it would be more intuitive to first specify the _std::thread_ object and then the parameter. data member initialization can depend on one another, so in some other class structure, this could matter.
- We provide an option to get access the underlying thread object (the _.get()_ method), like how smart pointer classes allow access to the underlying pointer. this means we can still use our RAII in an interface that requires an _std::thread_ object.
- Before the destructor invokes the behavior detailed by the action, we check that the thread is joinable. calling a _.join()_ or _.detach()_ on a un-joinable _std::thread_ is undefined behavior. it's possible that the object was already detached or joined before the destructor.

there is a possible, theoretical, nearly impossible chance that a race condition exists between the checking of the whether the _std::thread_ is joinable and the action on it, but this code is called only during the destructor stage, so if there's a race condition, it's because the user tried really hard to somehow have the other method call on the std::thread object happen in the same time as the destructor.

here is how this class will be used with our earlier example.

```cpp
bool doWork(std::function<bool(int)> filter, int maxVal = tenMillion)
{
    std::vector<int> goodValues;
    ThreadRAII t(std::thread(
    [&filter, maxVal,&goodValues]
    {
    for (auto i = 0; i<= maxVal;++i) //probably better to use std::copy_if
        {
            if (filter(i))
            {
                goodValues.push_back(i)
            }
        }
    };) ,ThreadRAII::DtorAction::join);

    auto nativeHandler  = t.get().native_handle(); // we want to set the thread priority
    //...

    if (conditionsAreSatisfied())
    {
        t.get().join();
        performComputation()
        return true;
    }
    return false;
}
```

we choose the join as an destructor action because the thread uses local variables (the vector), and detaching the thread will mean that the thread will continue to change that memory address, even after it's out of scope. in this case, we will pay in performance (the destructor must wait for the thread action to complete), but that was a conscious decision by us.

there are cases that this won't lead just to performance hits, but to actual hung problems, but this will be detailed in [Item 39](). there is no standard 'interruptible thread' in c++11, but one we be theoretically written by hand for some cases.

because we declared a custom destructor, we won't get a compiler generated move operations (copy operations are disabled because _std::thread_ has no copy operations), but we can ask the compiler to generate them for us.

```cpp
class ThreadRAII
{
    public:
    enum class DtorAction{join,detach};
    explicit ThreadRAII(std::thread&&t, DtorAction a):action(a),t(std::move(t))
    {}
    ~ThreadRAII()
    {
        if (t.joinable())
        {
            if (action == DtorAction::join)
            {
                t.join();
            }
            else
            {
                t.detach();
            }
        }
    }
    std::thread& get() {return t;}
    ThreadRAII (ThreadRAII &&) = default; //move ctor
    ThreadRAII& operator=(ThreadRAII &&) = default; //move assignment
    private:
    DtorAction action;
    std::thread t;
};
```

#### Things to Remember

> - Make _std::threads_ un-joinable on all paths.
> - join-on-destruction can lead to difficult-to-debug performance anomalies.
> - detach-on-destruction can lead to difficult-to-debug undefined behavior.
> - Declare _std::thread_ objects last in lists of data members.

</details>

### Item 38: Be Aware of Varying Thread Handle Destructor Behavior

<!-- <details> -->
<summary>

</summary>

A joinable _std::thread_ corresponds to an underlying system thread of execution
a future of a non-deferred task corresponds similarly to ta system thread, and therefore both *std::thread *and _std::future_ (and _std::shared_future_) objects can be conceptualized as _handlers_ to the system threads.

despite that, there are difference of behavior with the destructors of the two objects, _std::thread_ terminates the program as part of the destructor behavior if it wasn't previously joined or detached.
but for _std::future_ objects, the destructor behavior is sometimes similar to _.join()_, sometimes to _.detach()_ and sometimes something else. and it never causes program termination.

<!-- </details> -->
