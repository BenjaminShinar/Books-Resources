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

<!-- <details> -->
<summary>

</summary>

<!-- </details> -->
