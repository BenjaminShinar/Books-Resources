## Chapter 7 - The Concurrency API

<summary>
Using the correct asynchronous tool from the standard library, understating different behavior and dangers associated with each.
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

<details>
<summary>
The std::future object has different behavior at the destructor.
</summary>

A joinable _std::thread_ corresponds to an underlying system thread of execution
a future of a non-deferred task corresponds similarly to ta system thread, and therefore both *std::thread *and _std::future_ (and _std::shared_future_) objects can be conceptualized as _handlers_ to the system threads.

despite that, there are difference of behavior with the destructors of the two objects, _std::thread_ terminates the program as part of the destructor behavior if it wasn't previously joined or detached.
but for _std::future_ objects, the destructor behavior is sometimes similar to _.join()_, sometimes to _.detach()_ and sometimes something else. and it never causes program termination.

A _std::future_ is one end of a communication channel, we have a caller and a callee, and both usually run asynchronously.
the callee transmits a result to the caller, this is done by writing the result of a computation into the communication channel, usually in the form of an _std::promise_ object.

caller has _std::future_.
callee has _std::promise_ (most often).

but the result of the callee can't be actually stored inside the _std::promise_ itself, because that object might go out of scope when the computation is complete. it's also not practical to store the result in the _std::future_ of the caller, one problem is that a _std::shared_future_ can be created from it, and that the new object might be copied, but the result isn't always a copyable object. so the result must have both a lifetime that exceeds that of the _std::promise_ object and not be part of the _std::future_.

the result is actually stored in a location called _'shared state'_ which is in neither of the options. it's usually a heap-based object, although the standard leaves all the implementation details free for the library developers to decide.

this leads us to the behavior of the _std::future_ destructor, which depends on the _shared state_.

> - **The destructor for the last future referring to a shared state for a non-deferred task launched via _std::async_ blocks** until the task completes. In essence, the destructor for such a future does an implicit join on the thread on which the asynchronously executing task is running.
> - **The destructor for all other futures simply destroys the future object**. For asynchronously running tasks, this is akin to an implicit detach on the underlying thread. For deferred tasks for which this is the final future, it means that the deferred task will never run.

to make it simple, this is similar to reference counting objects, any _std::future_ except the last simply disassociates itself from the task. the last standing _std::future_ waits until the task is complete before it can be destroyed.

in different words, the special behavior of blocking happens only when:

> - It refers to a shared state that was **created due to a call to _std::async_**.
> - The task’s launch policy is **_std::launch::async_**, either because that was chosen by the runtime system or because it was specified in the call to _std::async_.
> - The future is the **last** future referring to the shared state. For _std::future_,
>   this will always be the case. For *std::shared_future*s, if other *std::shared_future*s refer to the same shared state as the future being destroyed, the future being destroyed follows the normal behavior (i.e., it simply destroys its data members).

the api for _std::future_ doesn't provide a way to determine if it the shared state it refers to is a result of a call to _std::async_, so there's no sure way to know if the destructor will cause blocking behavior. in this example, the vector might contain _std::future_ that were created with _std::launch::async_,so the destructor of the vector might block. the Widget class has a _std::shared_future_, so it's possible that any Widget object will cause a blocking as part of the destructor behavior.

```cpp
std::vector<std::future<void>> futs;
class Widget
{
public:
//...
private:
std::shared_future<double> fut;
};
```

although we can't know if a _std::future_ will block, we can know for sure that it doesn't block. if it doesn't satisfy the above conditions, it won't block. so if it wasn't created bt _std::async_, it won't block. that't the case for _std::packaged_task_ which can also produce a _std::future_ object.

```cpp
int calcValue();
std::packaged_task<int()> pt (calcValue);
auto fut = pt.get_future(); // won't block
std::thread t(std::move(pt));
//...
```

_std::packaged_task_ is an object that can be run on a thread,and is a move-only object, so we can cast it to rvalue reference to move it into a _std::thread_. the behavior now depends on what happens to the _std::thread_.

- if nothing happens with t, the _std::thread_ is still joinable at the end of the scope, and the program is terminated.
- if t was joined, there is no need for fut to block.
- if t was detached, no need to block.

the behavior of the _std::packaged_task_ and the _std::future_ it creates depends on how the containing _std::thread_ is handled

#### Things to Remember

> - Future destructors normally just destroy the future’s data members.
> - The final future referring to a shared state for a non-deferred task launched via _std::async_ blocks until the task completes.

</details>

### Item 39: Consider void Futures for One-Shot Event Communication

<details>
<summary>
a void std::future can be used a one time synchronization mechanism with a surprisingly easy syntax and good performance.
</summary>

#### The _conditional variable_ Approach

an traditional way of inter-thread communication involves using the _conditional variable_ as a way to check if a condition was qualified before continuing. one thread waits on the _conditional variable_ and other notifies it.

```cpp
std::condition_variable cv;
std::mutex m;
//...
// detecting thread A
{
    cv.notify_one();
}

// waiting/reacting thread B
{
    {
        std::unique_lock<std::mutex> lk(m); //lock mutex
        cv.wait(lk);
        // react to event while holding m
    }
    //m is unlocked
}
```

the code above works, but the use of the _std::mutex_ might be too much if the only thing it does is limit access until a object was initialized and that data is no longer shared. furthermore, the code ignores two possible cases:

- if the detecting thread A acts before the waiting thread B starts waiting, the notification to the _conditional variable_ is wasted, and thread B will never be released. this can be solved by having a condition checked prior to waiting, but it means the detecting thread will also need to control the mutex.
- the waiting thread ignores the possibility of spurious wake-ups, a mythical occurrence that can happen for no reason, but tends to happen a lot more with multi-core machines. the solution is to check the whether the condition is satisfied, which was traditionally done with a while loop, but can be done with a lambda in the _conditional variable_ call to _.wait()_.

```cpp
cv.wait(lk,[] {/*condition*/})
```

the problem in our current scenario is that the _conditional variable_ is waiting for an event to happen, that's the whole point of it's existence. we would need another shared data to use for this condition, like the _std::atomic\<bool>_ flag.

we can do the easy thing and just use the atomic flag on it's own, which is easy and short. no _std::mutex_, no spurious wake-ups, no _std::condition_variable_. But the waiting thread B is spending all of it's allocated time slice checking on the flag. it's **busy waiting** and wasting cpu resources.

```cpp
std::atomic<bool> flag(false;)
//...
// detecting thread A
{
    flag= true;
}

//waiting thread B
{
   while(!flag)
   {
     //do nothing!
   }; //wait for event
   //react to event.
}
```

the advantage of the _std::condition_variable_ is that it truly causes the thread to wait, it yields away the time-slot it receives from the scheduler so the program runs faster. it's possible to combine the two approaches together. because the flag is now guarded by the _std::mutex_, we can use a normal variable and not an atomic one.

```cpp
std::condition_variable cv;
std::mutex m;
bool flag{false};
//...
// detecting thread A
{
    {
        std::lock_guard<std::mutex>g(m);
        flag =true;
    }
    cv.notify_one(); // no need to hold the mutex right now
}

// waiting/reacting thread B
{
    {
        std::unique_lock<std::mutex> lk(m); //lock mutex
        cv.wait(lk,[]{return flag;}); // check before entering and at wake-ups, if true, continue, otherwise wait.
        // react to event while holding m
    }
    //m is unlocked
}
```

#### The Task and Future Approach

the alternative way to handle this case is by using a task and waiting on a _std::future_ object. rather than using _std::future_ and _std::promise_ as a two way communication channel, they can be used as a way to inform that an event has taken place.

the detecting task has and _std::promise_ object, while the waiting/reacting task holds the corresponding _std::future_ object. when the detecting task is ready, it sets the _std::promise_ (writes a value into the communication channel), and the waiting thread waits on the _std::future_ that it holds, and when the value is set, the task is free to continue running.
both _std::future_ and _std::promise_ are templated and require a type parameter, which specifies what type of data is passed through the channel. in this case, we have no type, only the existence of having something written to the channel, so the _void_ type is good for us.
the detecting task will use _std::promise\<void>_ and the reacting task will use a _std::future\<void>_ object.

```cpp
std::promise<void> p;
// detecting thread A
{
    p.set_value();
}

// waiting/reacting thread B
{
    p.get_future().wait(); //get the future from the promise and wait on it.
}
```

this design is simple, uses only one shared object (the _std::promise_), immune to spurious wake-ups and works even if the detecting task A finished before task B started waiting.

but it's not perfect. a _std::promise_ and _std::future_ have a _shared state_ between them, so there are some heap allocation costs. but more than that, this approach is a one-time only mechanism. once set, we can't unset the value of the _std::promise_, unlike conditional variables or flags which can be reused.

the limit isn't as bad as it seems, in this example, we want to first create the thread in a suspended state, so that when we want it to start, it's already allocated and configured (maybe we change the thread priority or cache affinity with the native handles).

```cpp
std::promise<void> p;
void react();
void detect()
{
    std::thread t([]
    {
        p.get_future().wait();
        react();
    });
    // do something before launching thread.
    p.set_value(); // the thread can not start;
    //... do more work;

    t.join(); //never forget to join
}
```

as we know, we have a problem with letting _std::thread_ run wild, so we can use our RAII class from before, unfortunately, this isn't as secured as we think it is. if there is an exception before the value is set, we have effectively forced ourselves into an hung situation. we are waiting for a thread that is blocked with no one to unset it.

```cpp
void detect()
{
    ThreadRAII (std::thread t([]
    {
        p.get_future().wait();
        react();
    }),ThreadRAII::DtorAction::join); //risk involved
    //... do something before launching thread. what if there's an exception here?
    p.set_value(); // the thread can not start;
    //... do more work;
}
```

this version doesn't fix the above issues, but it shows that even a one way communication channel can be effective for a large number of tasks, if they all depend on the same event. we do this by using _std::shared_future_ instead to suspend and un-suspend man reacting tasks.

```cpp
std::promise<void> p;
auto threadToRun = 5;
void react();
void detect()
{
   auto sf = p.get_future().share(); //get std::shared_future
   std::vector<std::thread> vt;
   for (int i = 0; i < threadToRun;++i)
   {
       vt.emplace_back(
           [sv]
           {
               sf.wait();
               react();
            };) // lambda, local copy of shared_future sf.
   }
   //... detect hangs somehow and do something,
   p.set_value(); // unsuspend all thread
   for (auto & t :vt)
   {
       t.join(); //join all threads
   }
}
```

#### Things to Remember

> - For simple event communication, _conditional variable_-based designs require a superfluous mutex, impose constraints on the relative progress of detecting and reacting tasks, and require reacting tasks to verify that the event has taken place.
> - Designs employing a flag avoid those problems, but are based on polling, not blocking.
> - A _conditional variable_ and flag can be used together, but the resulting communications mechanism is somewhat stilted.
> - Using _std::promises_ and futures dodges these issues, but the approach uses heap memory for shared states, and it’s limited to one-shot communication.

</details>

### Item 40: Use _std::atomic_ for Concurrency, _volatile_ for Special Memory

<details>
<summary>
Atomic variables and volatile variables are different, and have different use cases. they aren't interchangeable and there isn't a superior choice between them.
</summary>

The _volatile_ keyword has nothing to do with concurrency in c++. it's used in some other languages to work together with concurrent code, but it's supposed to be used for it.

#### _volatile_ Variables Shortcomings for Concurrency

The _std::atomic_ template is what should be used, all operations on it are guaranteed to be atomic, not only as if they were guarded by a mutex, but actually using special machine instructions.

```cpp
std::atomic<int> ai(0);
at = 10; //atomic set
std::cout << ai; //atomic read - 10
++ai; //atomic increment - 11
--ai; //atomic decrement - 10

volatile int vi(0);
vi = 10;
std::cout << vi;
++vi;
--vi;
```

no matter what other thread does, the value of the _std::atomic_ is set in an atomic matter. a read will always produce a valid value, as will a write or any other operations. even the increment and decrement operations (_++_,_--_), which are usually considered to be read-modify-write (_RMW_) operations are atomic for this class (unlike `ai = ai+1;`, which is atomic in each sub expression, but not as a whole). even the comparison operators are atomic.

in contrast, _volatile_ variables have no guarantees in a multi-threaded context. there are no assurances for the values in the variable.
assume we have two counters, one atomic and one _volatile_, and two threads the increment them. we can use the example code from the previous chapter for this (_std::promise_ and _std::shared_future_).

```cpp
std::atomic<int> atomic_counter(0);
volatile int volatile_counter(0);

auto func = [&atomic_counter,&volatile_counter](){++atomic_counter;++volatile_counter;};
// do run the code in two different threads.

```

we can be absolutely sure that the value of the atomic_counter is 2. but for the volatile_counter, we have no assurances.
after all. this is a possible scenario

> 1. Thread 1 reads vc’s value, which is 0.
> 2. Thread 2 reads vc’s value, which is still 0.
> 3. Thread 1 increments the 0 it read to 1, then writes that value into vc.
> 4. Thread 2 increments the 0 it read to 1, then writes that value into vc.

which in pseudo code will be like

```cpp
int temp_vc1 = vc; //0
int temp_vc2 =vc; //0
vc =temp_vc1 +1; //1
vc =temp_vc2 +1; //1
```

the final value of the volatile_counter is 1, even if it was incremented twice. compilers are built to assume no data races, and they re-organize code to make it optimized and run faster. having a data race together with the compiler optimizations can lead to unexpected behavior.

it isn't only the case of RMW that _std::atomic_ work and _volatile_ variable fail. this can also happen if we tried using a _volatile_ variable as a flag in inter-thread communication

```cpp
volatile bool valueAvailable(false);
auto importantValue =computeImportantValue();
valueAvailable = true;
```

while the programmer assumes this code is fine because the flag is set only after the important value is computed, but the compiler only sees two assignments which are independent from one another, so it can reorder them, or the hardware can reorder them. had we used an _std::atomic_ variable, things would have been different.

```cpp
std::atomic<bool> valueAvailable(false);
auto importantValue =computeImportantValue();
valueAvailable = true;
```

in this case, the compiler must create code in which anything before the assignment to the atomic variable must happen before, and anything afterwards must happen afterwards, the code is parted around the assignment. the compiler must also generate the machine code in a way that the hardware can't reorder the commands. (this is called _sequential consistency_, there are other models of consistency that are used for other cases).

#### _volatile_ Variables are for Special Memory Addresses

so, _volatile_ isn't good for operation atomicity or for code reordering, so what are they used for?

> "In a nutshell, it’s for telling compilers that they’re dealing with memory that doesn’t behave normally"

normal memory has values, and those values remain there unless the program changes them. so as long as the compiler can see who changes the values and who reads from them, it can optimize and reorder them as it pleases. so if we write to a memory address, never use the value, and then write to it again, a compiler might eliminate the first write.

```cpp
auto a = 5;
auto b = 9
auto c = b;

//... do stuff without a and without changing b or c
c=b; //c
a= 10; // change a's value
foo(a); // use a
```

a smart compiler would see that the first value of a is never used, so it will drop the instruction, and also that if b and c aren't changed, the second assignment is meaningless. these pointless read and writes are called _'redundant loads'_ and _'dead stores'_. and although human programmers don't write code like this directly, it can be generated from templates and inlining and other compiler work. after all that work, the compiler can remove all the pointless instructions it generated.
but this is only for normal code and normal memory. there is **special** memory that behaves differently.

the most common kind of special memory is _memory-mapped I/O_, or peripherals, like a screen, a controller, a sensor, network ports, etc.... if we mapped a memory address to an i/o device, those _redundant loads_ and _dead stores_ now make sense. we read from sensors multiple times, because the sensor can change that data, and when we write to a memory location that's mapped to a peripheral, we might be issuing commands to a radio or to the screen.

the _volatile_ keyword tells the compiler that we are dealing with special memory, and that it shouldn't perform optimizations on this memory address.

```cpp
volatile int x;
auto y =x; //y is int, auto deduces away cv qualifiers. can't be optimized because volatile x is involved
y= 9;
y=8; //can be optimized away
y = x; //can't be optimized because volatile x is involve
x =10; //not optimized away
x=20; //not optimized away
```

this is something that _std::atomic_ variables can't do, actually, even though they are guaranteed to be atomic, the compiler can remove redundant operations on them. and also, atomic variables can't be copied or moved even. if we want to initialize a new _std::atomic_ variable, we need to get value explicitly. there is also an issue with performing read and write on atomic variables that requires using the _.load()_ and _.store()_ methods.

```cpp
std::atomic<int> ai1(5);
ai1++;
//auto ai2 = ai1; //error!
auto i = ai1.load(); //i is int;
std::atomic<int> ai2(ai1.load()); //this works
ai2.store(ai1.load()); //also works, each part is atomic, but not atomic together
```

in the above case, the compiler might decide to store the results of the load in a register rather than read them again, which wouldn't work for io mapped memory.

The situation should thus be clear:

> - _std::atomic_ is useful for concurrent programming, but not for accessing special memory.
> - _volatile_ is useful for accessing special memory, but not for concurrent programming.

of course, we can actually use them together, maybe a special mapped memory is used by several threads.

```cpp
volatile std::atomic<int> val;
```

some programmers suggest to always use the _.load()_ and _.store()_ methods of the _std::atomic_ variables as a way to remind other programmers that this operations are costly, and seeing too many of them is a sign of a code that going to be hard to scale up.

#### Things to Remember

> - _std::atomic_ is for data accessed from multiple threads without using mutexe. It’s a tool for writing concurrent software.
> - _volatile_ is for memory where reads and writes should not be optimized away. It’s a tool for working with special memory.

</details>
