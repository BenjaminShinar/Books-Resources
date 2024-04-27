<!--
// cSpell:ignore objdump Browsable Guttag nsenter setcap getpcaps fsanitize Nlohmann httplib Dennard Metaparse Lexy ctre idents crend crbegin truncatable awslabs composability toolset ftest fprofile fcoverage Doxygen Fuzzers ixin nonstatic
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Core C++ 2023

<!-- <details> -->
<summary>
Israel C++ Convention.
</summary>

[playlist](https://www.youtube.com/playlist?list=PLn4wYlDYx4bs0p9S6aFvKaASoCLFVwt_E), [Schedule](https://corecpp.org/schedule/).

- Designing and Implementing Safe C++ Applications
- Concurrency in Modern C++
- Linux System Programming Essentials
- ~~Keynote :: Approaching C++ safety~~
- ~~Let's talk about C++'s Abstraction Layers~~
- ~~Instruction Level Parallelism in Your C++ Program~~
- ~~Being RESTful with billions of dollars in transactions, thanks to C++, JSON and HTTP~~
- C++23 ranges: conceptual changes and useful practicalities
- ~~Running Away From Computation - An Introduction~~
- ~~Accelerated C++ with OpenMP~~
- ~~Virtual templated methods~~
- ~~UB effects on real world projects~~
- ~~Scope Sensitive Programming~~
- ~~Does the C++ compiler work too hard?~~
- ~~Linker & Loader: The build process after-party~~
- ~~Easy to use, hard to misuse: Practical guidelines to tame your API~~
- ~~Expressive Compile Time Parsers~~
- ~~Lazy and Proud: How I Failed to Standardize lazy_counted_iterator~~
- Better Code: Exploring Validity
- ~~Understanding Linux user namespaces~~
- ~~From a modern to an unbelievably modern C++~~
- ~~Multi-Paradigm Programming and Beyond~~
- ~~The Imperatives Must Go!~~
- Nobody Can Program Correctly. A Practical and Interactive Guide to Debugging C++ Code
- ~~MDSPAN: A Deep Dive Spanning C++, Kokkos & SYCL~~
- ~~C++ Incidental Explorations~~
- ~~More Ranges Please~~
- ~~To Int or to Uint, This is the Question~~
- ~~Improving Compilation Times: Tools & Techniques~~
- ~~Development Strategies - The stuff around the code~~
- ~~Standard C++ toolset~~
- ~~Performance-related coding guidelines~~
- ~~Building low latency, network intense applications with C++ (in Hebrew)~~
- ~~Lessons I learn from improving legacy product and doubling its performance (in Hebrew)~~
- Compile time polymorphism: the optimization that (sometimes) isn't
- The Concept of Templates
- ~~C++ Horizons~~
- ~~Concurrency Improvements in C++20: A Deep Dive~~
- ~~C++ for the cloud~~
- Exceptionally Bad : The story on the misuse of exceptions and how to do better
- ~~Keynote :: Expressing Implementation Sameness and Similarity in Modern C++~~
- Ready your C++ code for the multi-device, multi-vendor world
- CUDA and the Latest Innovations in GPU Technology

## David Sankel :: The Semicolon is a Lie

<details>

<summary>
Historical tour of how programming evolves and moves away from directly translating source code into machine code.
</summary>

[The Semicolon is a Lie](https://youtu.be/ICf_6L1kJcE)

programming history, David's history and how they intersect.

> 1. Computers are fast
> 2. Programming is an illusion

TI 99/4A computer:\
Texas instruments old "computers" came with manuals of the code to type into it and then you could run the "games".

| Metric           | TI 99/4a  | 386 sx      | Pentium       | Z600                                     |
| ---------------- | --------- | ----------- | ------------- | ---------------------------------------- |
| Ram              | 16 Kb     | 4 Mb        | 16 Mb         | 24 Gb                                    |
| Registers        | 16 bit    | 32 bit      | 32 bit        | 64 bit                                   |
| CPU              | 3 Mhz     | -           | 133Mgz        | 2.64Ghz                                  |
| Memory           | -         | 40Mb        | 500Mb         | 1Tb                                      |
| Speed Over human | 3,829,787 | 101,333,333 | 2,537,000,000 | 72,090,000,000 (without multi-threading) |

> **semicolon** - A mark (`;`) of punctuation, indicating a greater degree of separation than the comma.

**1947** - Assembly language, by _Kathleen Booth_ and her husband _Andrew Booth_, created for the A.R.C machine, in preparation that the same instructions could carry over to more modern machines as they become available. **1951** - abstraction from a machine that creates instructions to a language that creates a the instructions. later we got _Grace Hopper_ and the A-0 system (arithmetic Language version 0), which laid the foundation for the first compiler. next we meet _John Backus_ and **Fortran**, which introduced the optimizing compiler, and since then we no longer directly translate source code into machine code, we have something that changes it. _Frances E.Allen_ introduces graph theory in the sixties.\
_Dennis Richie_ and _Ken Thompson_ creating the C language to work on the Unix operating system, later on _Bjarne Stroustrup_ which borrowed from **Simula** and BCPL and created **C++**.

When Pentium 4 were released, the "NetBurst" architecture, instructions execute at the same time, using something called "shadow registers". so it's not only that the compiler modifies the source code into something else, the CPU also modifies the assembly code.

</details>

## Bryce Adelstein Lelbach :: AI-Assisted Software Engineering

<details>
<summary>
Using AI to change how we write code.
</summary>

[AI-Assisted Software Engineering](https://youtu.be/9P0PN29VrfY), [C-Why github](https://github.com/plasma-umass/cwhy).

Large Language Model (LLM), using text context (widnow of text), which can be dropped after a while. so we need to know how to interact with the model.

Neural networks are the building blocks of AI, not the end-all type. we are moving towards more advanced model, such as neural network agents, information retrieval and execution environments.

### What Can We Do With AI

Creation and Analysis. Creation is the process of answering open-questions, creating new code, there is no "right" answer. Analysis is the next stage, reviewing code, fixing errors, re-factoring. these are questions that have an answer, this is a task that is easier for AI to do.

one example is the **C-why** tool which explains why code compilation fails. it takes a diagnostics output and tries to understand it.

> - Classification - what tools are involved?
> - Source Identification - what files or resources do we need to access?

the cycle is:

1. Analyze what we have.
2. Determine what else we need.
3. Collect what we need.

these are series of tasks, so it's suitable for AI tools. data extraction can be text only, but it's better to have code entities (such as function + comments).

We would want the tool to be integrated with the automation CI-CD pipeline, and to run when it fails. it will analyze the diagnostic check, gather the data, and create a suggestion for a fix and re-run tests on the patch.

</details>

## Noam Weiss :: Virtual Templated Methods

<details>
<summary>
Virtual Templated Function don't actual exist, but we can get around it.
</summary>

[Virtual Templated Methods](https://youtu.be/Z-WzYbTm8k0)

> the use case:
>
> - I want to decouple my debugging/logging by using dependency injection.
> - but I also want to support types that I don't know in advance.
> - Templated virtual functions would have been great.

our other options would be:

1. using inheritance instead of templates.
2. Break into two functions:
   1. use template to reduce to a common type.
   2. use virtual function on the common type.
3. Combine both approaches:
   1. use template to create the derived class.
   2. use the virtual function with the base class.
4. Reflection (if we implement it ourselves).

</details>

## Tomer Vromen :: Linker & Loader: The Build Process After-Party

<details>
<summary>
Some Stuff About the Linker and the Linking Process.
</summary>

[Linker & Loader: The Build Process After-Party](https://youtu.be/xc23weUlZ0A)

Linkage errors aren't the same as compiling errors, the compiler turns the source code into a machine code (object file), the linker takes all the object files and system libraries and creates the executable.

The Linker's Responsibilities are:

- Layout Code
- Layout Data
- Resolve Symbols

The gcc `-c` flag makes it so only the first step of compilation is performed, and it outputs an object file, we can then call `objdump -d` and look at the code disassembly (`objdump -t` will show the symbols).\
This includes the mangled names and the machine code instructions (we can pass it through `| c++fill` to get de-mangled names for better readability).\
the assembly code has function calls in assembly, the instruction is call, and the op-code is zero. this is because the compiler doesn't know where the actual code is, and it needs the linker to fill it in. we could also add the `--reloc` flag to the object dump and see how the code expects the re-locations should work.\
we can run the disassembler on the executable file and see how it looks after linking. now the addresses are filled out with actual locations.\
if we want to link with a library (static archive) we pass the library with `-L` path argument, we can use the `-###` flag to tell gcc to print what it actually will run, and it will show the entire command it would use, including linking the standard libraries. the linking order is sometimes important.

C++ bring some complexity to the table, function overloading and templates create name mangling. there are also C function that we need to define as <cpp>extern</cpp>, and there are inline class method defintions, which relate to ODR. there is a special memory location for <cpp>thread_local</cpp> data. we can even use unicode identifiers (🦆).

> Common errors:
>
> - "undefined reference"
>   - missing library object file in linkage command
>   - missing <cpp>extern "C"</cpp>
>   - wrong linkage order (it's the opposite of how `# include` works)
>   - "abi::cxx11" or "\_\_cxx11" - libstdc++ dual ABI mismatch
>   - missing destructor, in virtual classes, must be defined, even if <cpp>~Class() =0;</cpp>
> - "multiple definitons" - probably a function defined in the header
> - "linker out of memory" - are you creating to many types?

### LTO - Link Time Optimization

all modern compilers support LTO, it requires a special flag in both compilation and linkage (so the object file keeps some information), and it's not always worth doing it.

### Share Libraries

instead of packaging the same common libraries, we can have one shared version of it in the memory and use it for all programs, but it can lead to "DLL hell". we can also have dynamic linking loader, or use <cpp>dlopen</cpp>. there is also the issue of **Wrapping/Hijacking**, we can tell the linker to call a wrapper object instead of calling the function directly,and then we can use the wrapper to redirect the calls.

> Caveats:
>
> - Hijacking inside a library doesn't always work
> - Hijacking non-function symbols is not officially supported
> - Hijacking class methods is complicated

we do this by passing two flags `--Wl` which instructs gcc to pass a command to the linker, and `--wrap=<mangled name>` which replaces the symbol with a symbol that is defined with the same name. (the demo didn't work so great).

</details>

## Bjarne Stroustrup :: Approaching C++ Safety

<details>
<summary>
Different Notions of Safety and How to Achieve it.
</summary>

[Approaching C++ Safety](https://youtu.be/eo-4ZSLn3jc)

### The Challenge of Safety

The NSA guide says that software should be written in memory safe languages, and it doesn't mention C++ as a memory safe language.

the C++ language is "strongly typed, weakly checked language", which is nice and well, but it doesn't scale up. but we don't want to limit what kind of applications can be written, and without adding run-time overhead.

> Type and Resource Safety
>
> - Every object is accessed according to the type with which is was defined (type safety).
> - Every object is properly constructed and destroyed (resource safety).
> - Every pointer either points to a valid object or is the <cpp>nullptr</cpp> (memory safety).
> - Every reference through a pointer is not through teh <cpp>nullptr</cpp> (often a run-time check).
> - Every access through t a subscripted pointer is in-range (often a run-time check).

The solution must serve a wide variety of user/areas, it can't break existing code, it can't defer to another language, and it can't rely on all the developers "magically" improving. the challenge is to have a type-safe c++ language and to convince developers to use C++ in a safe way.

### C++ Evolution

C++ stared with two goals - efficient use of hardware (like C), and managing complexity (based on Simula). it also meant enforcing argument type checking. a different jey idea is to "represent concepts in code". <cpp>RAII</cpp> - resource acquisition is initialization, not only memory resources, also file handles, locks, sockets, shaders.\
In the early 80's, Object oriented programming was emerging, encapsulation, abstraction, overloading. then we have templates, containers, algorithms, smart pointers and exceptions.

### C++ Core Guidelines

> - no implicit violations of he static type system.
> - provide as good support for user-defined types as for built-in types
> - say what you mean - emphasizes declarative styles and abstractions.
> - syntax matters (often in perverse ways) - in general, verbosity is to be avoided.
> - leave no room for a lower-level language (except assembler).
> - preprocessor usage should be eliminated.
> -
> - make simple tasks simple.
> - make error handling regular.

(module <cpp>std</cpp> is better than `#include`).

the core guidelines are designed to be an answer to the question "what is good modern C++?". a useful answer that many people can use, and not just language experts. this is something that can be sometimes achieved with static analyzers.

but people don't like coding rules, and those coding rules usually don't provide good advice. it should be:

> - Good
>   - Comprehensive
>   - Browsable
>   - Supported by tools
>   - Suitable for gradual adoption
> - Modern
>   - "Compatibility and legacy code be dammed! (initially)"
> - Prescriptive
>   - Not punitive
> - Teachable
>   - Rationales and examples
> - Flexible
>   - Adaptable to many communities and tasks
> - Non-proprietary
>   - But assembled with taste and responsiveness

In the guidelines, the first rules are high-level conceptual ideas to defined the mental framework, these rules can't be "checked" by machines. the rest of the rules are "lower-level" rules, which can be automated and checked statically. if we can't remove un-safe stuff from the language (such as pointers), we can still hide it behind a zero cost abstraction (a span) and enforce that those unsafe operations are never used directly.

dangling pointers - example of unsafe code that usually works, until it doesn't (when the memory was recycled for some reason).

```cpp
void (X* p)
{
   // ...
   delete p;
}

void g()
{
   X* q = new X;
   f(q); // delete is called here
   // do stuff
   q->use(); // will crash, or read random memory,
}
```

> Owners and Pointers:
>
> - Every object has one owner.
> - An object can have many pointers to it.
> - No pointer can outlive the scope of he owner it points to.
> - An owner is responsible for owners in its object.

dangling pointers, pointers to local data, invalidations when re-allocation happens.

there are problems that require run-time checking.

### C++ Profiles

how to guarantee safety? making everybody follow the best guidelines without having them magically follow all the rules.

> Different notions of safety:
>
> - Logic errors
> - Resource Leaks
> - Concurrency Issues
> - Memory Corruption
> - type Errors
> - Overflows and Unanticipated Conversions
> - Timing Errors
> - Allocation Unpredictability
> - Termination Errors

these things can't be done by the compiler alone, and not everything could be achieved from static analysis. A safety profile is a set of rules that gauntness a safety result, such as bounds safe, type safe or memory safe, we want to be sure that unsafe code is never executed.

There is a problem of mixing profiles, between libraries and between languages.

</details>

## Inbal Levi :: Let's talk about C++'s Abstraction Layers

<details>
<summary>
A mental model of abstraction layers and how they interact together.
</summary>

[Let's talk about C++'s Abstraction Layers](https://youtu.be/wODpT8HJn-E)

### What Are Abstraction Layers?

software development is all about communicating logic to the computer, to achieve that, we need some level of abstraction (rather than writing assembly and machine code).

examples of abstractions: iterating, messaging. we can have under abstraction (not using enough) and over abstraction (not having enough data).

> The essences of Abstraction is **preserving information that is relevant** in a given context, and **forgetting information** that is irrelevant **in that context**.\
> ~ John V.Guttag

Types and pointer arithmetic also implement abstractions, advancing a pointer "moves" the pointer to a different location based on the types.

### Abstraction Layers Model For C++

analyzing keywords, concepts and elements in the language and identify layers and borders between them, and find which are dangerous.

```cpp
int main()
{
   int i = 0;
   std::cout << & i; // 0x7ffc8584005c
   *(*int)0x7ffc8584005c = 1; // undefined behavior
   return i;
}
```

in this example, we have three topic:

- the invalidity of the address.
- the duality of int and memory address
- the UB created by using the address.

we can say that have problem with the memory layout, the type system and memory contorl. lets add to it the "program and source code" topic, and we eventually have an hierarchy of concepts and how the relate to one another. with this classification in tact, we can say which statement relates to which layer.

```cpp
#include <sstream>

int main()
{
   auto iss = std::istringstream("0 1 2");
   auto j = 0;
   while (iss >> j){
      std::cout << "j: " << j << '\n';
   }
}
```

this print zero, one, two, as we expect. but let's add ranges.

```cpp
#include <sstream>
#include <ranges>

int main()
{
   auto iss = std::istringstream("0 1 2");
   for (auto i : rn::istream_view<int>(iss) || rv::take(1)) {
      std:::cout << "j inside loop: " << i << '\n';
   }

   auto j = 0;
   iss >> j;
   std:::cout << "j after loop: " << j << '\n';
}
```

In this example we see zero and then 2. this is contrary to our expectations (zero and one). the problem is that ranges take ownership.

### Existing Solutions

we need to be wary of the boundaries and be careful at spots where the interact with.

1. solution 1 - write better code, use better guidelines, enforce with tooling.
2. solution 2 - use a "different language" for new features - always write at the modern langrage style.

### Future Solutions - How Can We Do Better?

apply the layers model to our tools and give better error messages. classifying tokens according to layers, and warning when we combine layers that don't fit together. in the problematic example, we can warn that we move from the I.O abstraction layer to the rangers layer, and then we try moving back.

coroutines example:

```cpp
Task doWork(); // Coroutine

struct Task {
   struct promise_type {
      HandleWrap get_return_object() {return HandleWrap(this);}
      std::suspend_always initial_suspend()
      {
         //..
      }
      struct HandleWrap {
         void resume() {
            std::cout << "work\n";
            mHandle->resume();
         }
      };
   };
};

int main()
{
   auto work_handle = doWork();
   work_handle.resume();
}
```

this is similar to <cpp>std::execution</cpp>> that is planned for c++26.

```cpp
scheduler auto sch = thread_pool.scheduler();
sender auto begin = schedule(sch);
sender auto doWork = then(schedule(sch),[](){
   std::cout << "Work\n";
});

int main()
{
   this_thread::sync_wait(dorWork);
}
```

Because the implementers knew how similar the two ideas are, they designed the scheduler so it will fit with coroutines. but they still run into issues and limitations. other proposal should also follow and consider how their features interact with existing and other future features. tools can help us identify those interaction points.

</details>

## Hana Dusíková :: Lightning Updates

<details>
<summary>
Updating user applications fast and safely.
</summary>

[Lightning Updates](https://youtu.be/8zyTovAXXkQ?si=munNqdzIVNkmiQIT)

the basic requirement is:

> "I need to update an object on 100's of millions of clients, quickly and whenever I want"

the thing we want to update can be:

- an executable
- resources (database, model, textures)
- the state of the application or part of it

the state should be

- immutable
- consistent and secure
- representable with a data structure

the update mechanism can be replacement of everything, additional overlays, or differential.

$$
state_{n+1} = state_{n} + difference_{(n,n+1)}
$$

We can represent this a as a matrix or as a graph. but not all clients update everything in the same order, we don't want to have to go through all of the small updates each time, we would rather have points of major updates. we can do a search to find how to go from one point (version) to another. we represent the link between states as either a filename with a version or name of the release, or we identify each release as with the hash of the contents themselves. this hash value can act as a pointer, a unique value for the content, which makes the data immutable and easy to cache (can be stored on an edge location). each update includes snapshots of deltas of previous updates. this makes the search easier.

### Model of the Graph with Vocabulary Types

we want to mark the objects we use, the "nouns". a hash is just a bunch of bytes.

```cpp
template <size_t N> using hash = std::array<std::byte, N>;
template <size_t N> using hash_view = std::span<const std::byte, N>; // non-owning
```

or we can a have a strong type

```cpp
template <size_t N> struct hash {
   std::array<std::byte, N> value{};

   // constructors
   hash() = default;
   hash(const hash &) = default;
   hash(hash &&) = default;
   explicit hash(std::array<std::byte, N> in) noexcept: value{in} {}

   // comparisons
   friend auto operator <=>(hsh, hash) = default;
   friend book operator ==(hsh, hash) = default;

   // iterable
   auto begin() const noexcept {
      return value.begin();
   }

   auto end() const noexcept {
      return value.end();
   }

   auto begin() noexcept {
      return value.begin();
   }

   auto end() noexcept {
      return value.end();
   }
};

// same with hash_view
```

The above can be simplified by using inheritance, and we add the tagged hash over it, with sha256 options as well. we need a metadata type, it contains the hash of the subject, timestamps, links to previous state and snapshots, we also have some other objects like metadata, delta links, snapshots, etc..\
We need a way to serialize and deserialize the objects.

after we created the objects, we need a way to use them, these are the "verbs" we use, such as `unwrap_and_validate` which act on raw bye data and check if the object is what we expect it to be. there are unique methods to validating each of the inner types (tags, identifier, snapshots, metadata).

### State

we represent the state as a struct with metadata and a shared pointer to the subject. we can find the path between two links (for update) by using the `select_next` method to find it.

$$
\begin{align*}
state_m = state_n + path_{(n,m)} \\\
state_m = state_n + delta_{(n,n+1)} + ... + delta_{(m-1, m)} + metadata_m
\end{align*}
$$

this gives us a user api for updating any kind of object.

</details>

## Michael Kerrisk :: Understanding Linux User Namespaces

<details>
<summary>
Overview of Linux namespaces and capabilites.
</summary>

[Understanding Linux User Namespaces](https://youtu.be/XgThPoL9mPE?si=hDZEVQLJFEafIw63)

user namespace are important for building unprivileged containers in linux.

Namespace "wrap" around some global system resource to provide isolation, there are currently eight types of linux namespaces (the most resent one is from 2020).

> - UTS: isolate system identifiers (e.g., `hostname`, `domainname`)
> - Mount: isolate mount point list
> - IPC: isolate interprocess communication resources
> - PID: isolate PID number space
> - Network: isolate network resources such as firewall and routing rules, socket port numbers (`/proc/net`, `/sys/class/net`)
> - (and others: cgroup, time, user)

each namespace type can have multiple instance, but at system boot, there is one of each, this is the **inital namespace**. a process resides in one namespace instance (of each of the types).

for example, the UTS (comes from the ancient "unix time sharing") namespace isolates the two system identifiers returned by `uname(2)`: the node name the NIS domain name. all processes inside the same UTS namespace see the same hostname and domain name, but cannot effect and see what going on in other namespaces.

each process has symlink files (symbolic link) in `/proc/PID/ns` that link it to the correspondng namespace, the value of the links has the form of `<namespace type>:[magic inode number]`. the number is from an internally mounted namespace filesystem.

```sh
readlink /proc/$$/ns/uts
#uts:[4026531838]
```

if two processes have the same inode number for symlink, they are in the same namespace of that type.

- `unshare(1)` - create new namespaces and execute a command in them. default command in `sh`
- `nsenter(1)` - enter an existing namespace and execute a command in them.

### Demonstration of creating a UTS namespace

running in two shells at the same time, starting at the default namespace. in one shell we will create a new uts namespace, and then we'll enter it from the second shell.

```sh
hostname
readlink /proc/$$/ns/uts
# shell 1 only
sudo unshare -u bash
hostname # inherits the namespace from above
hostname changedName
echo $$ # get pid number
# continue in both shells
hostname
readlink /proc/$$/ns/uts
# shell 2
sudo nsenter -t <pid number from shell-1> -u
hostname
readlink /proc/$$/ns/uts # verify we are at the same namespace as shell 1
#
```

### Namespace Capabilites

Traditional linux has normal users and roo user. with the root user being able to skip many checks. normally, if we wish to have program run with root privileges, we need to make it capable of assuming the root role. so when it runs, it takes the UID of the file owner.

```sh
sudo -i
chown root prog
chmod u+s prog
```

this is powerful, but dangerous. if the program gets comprised, it can do anything the root user can. we don't have a way to limit the blast radius of the power. if we want the program to be able to change system time, then we must give it complete root user powers.

the concept of **Capabilites** is meant to remedy this by breaking the power of the super user into small pieces. at linux 6.4 there are 41 capabilites (see `capabilites(7)`). instead of setting programs to assume root user, we can have the attached with capabilites (using `setcap(8)`) to only do what it has to do. this is following the principle of least privilege.

### User Namespaces

we can have per-namespace mapping of user and group ids. for example, a process can have a non-zero UID (normal user) outside a certain namespace, and a UID 0 (super user) inside it.

user namespace are inside hierarchical relationship, each one has a parent (which created it), those relationship effect how the capabilites are moved. when a namespace is created, the first process in it has the super user privileges, but only for the namespace, this is done by having UID and GID mappings (writing to two files: `/proc/PID/uid_map` and `/proc/PID/gid_map`). such as mapping the zero uid inside the namespace to uid 1000 outside it.

shell 1:

```sh
id
unshare -U -r bash
id
cat /proc/$$/uid_map
cat /proc/$$/gid_map
grep -E 'Cap' /proc/$$/status # see capabilites
getpcaps $$ # same as above
hostname newName #fails, we don't have root for uts namespace
```

shell 2:

```sh
ps -o 'uid, gid,pid' 5356
```

the first process in the namespace has full privileges, but only for objects owned by that namespace. (something about non-user namespaces). if we want to discover the namespace relationships, we can check the `ioctl_ns(2)` manual page.

### Use Case and Applications

permit the application to do things without root privileges, such as docker containers and LXC or chrome-style sand-boxing.

</details>

## Elazar Leibovich :: UB Effects On Real World Projects

<details>
<summary>
some examples of undefined behavior.
</summary>

[UB Effects On Real World Projects](https://youtu.be/SEhNmLqrVxc?si=GkYupHTbu0SYGZFX)

real undefined behavior and examples of it.

undefined behavior is code that violates the language contract. but another way to put it is by saying that it is a problem of culture and values.

> "The language shall be designed to avoid error prone features and _maximize automatic detection_ of programming errors"\
> ~ the ADA language programming guide

but C++ isn't like that, the focus of C++ is on performance.

the first example is with excessive shifts, if we shift more than 32 bits, we have undefined behavior. the compiler knows it's undefined behavior, so it can optimize away the check against zero.

```cpp
groups_per_flex = 1 << sbi->s_log_groups_per_flex;
if (groups_per_flex == 0)
   return 1;
flex_group_count = v / groups_per_flex;
```

another example, the compiler is allowed to pointers passed to <cpp>strncpy</cpp> are not null, so it can omit any checks for null on them, and if any variable has been set to that pointer, all null checks on it are omitted as well. Many times undefined behavior is discovered when compilers are updated, since new compilers are better at optimizing, and can expose them.

in this example, we copy wide characters, but in windows it sometimes failed to copy all the bytes. it turns out that there different alignments for wide characters in linux and windows.

```cpp
void foo(char *src)
{
   wchar_t dst[100]={};
   wcsncpy(std, sec, 5);
   dst[5] = '\0';
}
```

undefined behavior of boolean evaluating to both true and false. the uninitialized value was first tested for non-zero, but the second test was optimized to just taking the first bit, and the results were different.

```cpp
bool b;
if (b) puts("B");
//...
if (!b) puts("!B");
```

an example with a macro, using the <cpp>this</cpp> pointer in the initialization list is undefined behavior.

```cpp
#define IDX_INIT(req) this->init((req, (Compile*) this__out))

Node::Node(uint req): _idx(IDX_INIT(req))
{}
```

Strict Aliasing is a common example of undefined behavior bugs (accessing an object through a pointer of a different type). it leads to a lot of compiler re-ordering.

```cpp
uint32_t a;
uint16_t *a_half = std::reinterpret_cast<uint16_t*>(&a);
std::cout << *a_half;
```

invalid pointers cannot be accessed or compared. and <cpp>realloc</cpp> can free the memory from the source pointer.

```cpp
int p* = malloc(sizeof(int));
int q* = realloc(p, sizeof(int));

if (p == q)
   printf("%d %d\n", *p, *q);
```

accessing a union in-active member is undefined behavior, adding pointers past the containers limit is undefined behavior.

we can't always use <cpp>-fsanitize=undefined</cpp>, but we should try it. we can add compiler flags to avoid some optimizations, and we should try with more than one compiler and interpreter to make sure we don't break because of it.

</details>

## Yossi Moalem :: Easy to use, Hard to Misuse: Practical Guidelines to Tame Your API

<details>
<summary>
Some Thoughts about creating APIs.
</summary>

[Easy to use, Hard to Misuse: Practical Guidelines to Tame Your API](https://youtu.be/wP9C36DM8K4?si=ookNiBxbZhEXKO9E).

APIs that have assumptions and preconditions that aren't properly conveyed, an example is misleading or unclear argument types, like days and months for dates, or when creating points by either coordinates or polar calculations. strong types are one solution, but they force boiler plate.

levels of ease for incorrect use:

- incorrect code will not compile
- incorrect code will crash
- Need to look at te prototype to get it right
- Need to read comments, documentations and examples to get it right
- requiring a non-trivial workaround

The more flexibility and power the users have, the more likely they are to use it wrong. the common use should be safe and easy. uncommon (and potentially unsafe) use can (and should be) harder.

Examples of constructors that can fail, and how to possible handle them. like c++23 <cpp>std::expected</cpp>.
Unclear names, should names indicate domain or software and program state in the stack? avoid having misleading names.

</details>

## Ivica Bogosavljevic :: Instruction Level Parallelism in Your C++ Program

<details>
<summary>
Tricks to increase Instruction Level Parallelism, and hiding dependency chains.
</summary>

[Instruction Level Parallelism in Your C++ Program](https://youtu.be/jfE8FqQIYko?si=mNy50AQyzwkMzffu)

ILP - Instruction Level Parallelism.

> - Modern CPU can:
>   - Execute more than one instruction in a single cycle
>   - Execute instructions our of order
>   - But there is a limit to how much work a CPU can do regardless of all tht tricks hardware uses to speed up computation.
> - Instruction Level Parallelism
>   - How much code can profit from the available HW resources

the main limiting factor on instruction parallelism is dependencies, an instruction can't be executed until the input data variables are ready. so even if we had a magical endless chip, it would still have to wait.

quiz: endless machine, can do infinite parallelism, exception, memory load and store operation take 3 cycles, other operation take one cycle.

loop1: equivelent of <cpp>std::transform</cpp> on an array or a vector.

```cpp
for (int i = 0; i < n; i++) {
   c[i] = a[i] + b[i];
}

// "semi-assembly" code equivelent
for (int i = 0; i < n; i++) {
   register a_v = load(a + i);
   register b_v = load(b + i);
   register c_v = a_v + b_v;
   store(c + i, c_v);
}
```

~~the `a+i` is one cycle, loading is 3 cycles. we can do `b+i` at the same time (we can probably calculate `c+i` as well). but we have to wait for both to finish before we can run `a_v+b_v`, and only then can we store the value. so the total is $1+3+1+3=8$ cycles. regardless of the size of the vector.~~

(actually, we don't have dependencies for `a+i` as an instruction, so it's seven cycles.)

loop2 example: equivelent of <cpp>std::reduce</cpp> on a vector. with a bit of unrolling.

```cpp
auto sum = 0;
for (int i = 0; i < n; i++) {
   sum += a[i];
}

// "semi-assembly" code equivelent
register sum_v = 0;
for (int i = 0; i < n; i+=2) {
   register a_v_1 = load(a + i);
   sum_v += a_v_1;
   register a_v_2 = load(a + i +1);
   sum_v += a_v_2;
}
```

we can load all data at the same time, but we have dependencies for the summing operations. so the total operations are $3 + n*1$ (without adding any tricks).

loop3 example: summing elements of linked list, equivelent of <cpp>std::reduce</cpp> on a linked list

```cpp
sum = 0;
while (current != null) {
   sum += curent->val;
   current = current->next;
}

// "semi-assembly" code equivelent
register sum_v = 0;
while (current != 0) {
   register current_val_1 = load(current + offsetof(val));
   sum_v += current_val_1;
   current = load(current + offsetof(next));
   if (current == 0) break;
   register current_val_2 = load(current + offsetof(val));
   sum_v += current_val_2;
   current = load(current + offsetof(next));
}
```

we can't do all loads at the same time, so we have dependencies on both the summations and the loads.

so we see that instruction dependencies force a speed limit even on the most powerful hardware imaginable. even the cpu can skip instructions and run them out of order, it still needs to wait for the input to be ready. so the instruction level parallelism is not a property of the machine, it arises from the source code. in the three examples:

> - Loop 1: no loop carried dependencies, all decencies are within a single iteration of the loop - high ILP.
> - Loop 2: no loop carried dependencies in data loads - medium ILP.
> - Loop 3: loop carried dependencies in data loads - low ILP.

the same constraints also apply for vectorization, multi threading, and other forms of parallelism.\
Each instruction in the cpu has two numbers that we need to consider:

> - **Latency**: Number of cycles that pass between the time the instruction is issued and it is finished.
> - **Throughput**: How many cycles does the CPU need to wait to issue the same instruction again.
> - Latency is always smaller than throughput.
> - Latency Limits software with low ILP, throughput limit software with high (ILP).

### Increasing ILP

code that usually has medium and low ILP:

> - reductions such as summing over arrays or vectors
> - "pointer chasing code" (linked lists, trees, hash maps with separate chaining)
> - long sequences of auto generated code
> - loops with large bodies (even without loop carried dependencies)

techniques:

> - Interleaving dependency chains - instead of processing only one dependency chain at a the time, we process two or more of them simultaneously
> - Shorting dependency chains - we decrease the length of the dependency chain
> - Decreasing the number of times we need to iterate a dependency chain
> - Break dependency chains - we completely remove the the dependency chain

interleaving example: two dependency chains at the same time.

```cpp
double cosine(double x)
{
   constexpr double tp = 1.0 / (2.0 * M_PI);
   x = x * tp;
   x = x - (double(0.25) + std::floor(x + double(0.25)));
   x = x * (double(16.0) * (std::abs(x) - double(0.5)));
   x = x * (double(0.225) * x * (std::abs(x) * double(1.0)));
   return x;
}

// interleaving

std::pair<double, double> cosine(std::pair<double, double> x)
{
   constexpr double tp = 1.0 / (2.0 * M_PI);
   double x1 = x.first * tp;
   double x2 = x.second * tp;
   x1 = x1 - (double(0.25) + std::floor(x1 + double(0.25)));
   x2 = x2 - (double(0.25) + std::floor(x2 + double(0.25)));
   x1 = x1 * (double(16.0) * (std::abs(x1) - double(0.5)));
   x2 = x2 * (double(16.0) * (std::abs(x2) - double(0.5)));
   x1 = x1 * (double(0.225) * x1 * (std::abs(x1) * double(1.0)));
   x2 = x2 * (double(0.225) * x2 * (std::abs(x2) * double(1.0)));
   return {x1,x2};
}
```

recursive calls to `cosine` create a loop carried dependency (each loop created the input of the next iteration). so this is an opportunity for us to employ interleaving, also an example of parallel lookups in a single tree (independet searches). this is the simplest thing to do, but it only works when the dependency chain is long enough, and it increases register pressure (less efficient assembly). the other options it keep the dependency chain, but make it shorter.

shortening dependency chains

```cpp
auto sum = 0;
for (int i = 0; i < n; i++) {
   sum += a[i];
}

// shorter chain + interleaving

auto sum_0 = 0;
auto sum_1 = 0;
auto sum_2 = 0;
auto sum_3 = 0;
for (int i = 0; i < n; i+=4) {
   sum_0 += a[i];
   sum_1 += a[i + 1];
   sum_2 += a[i + 2];
   sum_3 += a[i + 3];
}

auto sum = sum_0 + sum_1 + sum_2 + sum_3;
```

this is something that compilers can do automatically with integers on `-O3` flag, for floating point numbers, `-associative-math` or `-ffast-math` flags are needed (because of precision differences). when doing assembly intrinsics or vectorization this should be done manually.

vectorization shortening example

```cpp
int count(int x) {
   int cnt = 0;
   for (int i = 0; i < N; i++) {
      cnt += (a[i] == x);
   }
   return cnt;
}
```

for Linked Lists and Trees, we can still try and shorten the dependency chain. since lists are bad for memory locality, there are some alternatives, such as `Colony` that combine vectors with lists. we can decrease the number of passes over the chain. a simple example is inverting data accesses. if we have a vector and list, it's easier to iterate many times over the vector than over the list, so rather than search the list for each element of the vector, we search the vector for each element of the list.

the best case is breaking the dependency entirely, but this is harder to do and requires redesign. one option it to store the data from a list inside a temporary array (for repeated searches), or using n-array trees instead of pointer trees.

### Compilers, In-Order Processors and ILP

Smaller low-level processor in embedded world don't support instructions skipping. they are much more sensitive to even shorter dependency chains. when using compiler intrinsics, we can do loop unrolling and interleaving, and the more complicated technique of "loop pipelining". compiler explorer has an analytical tool `llvm-mca` to check for dependency pressure.

</details>

## Mike Shah :: Running Away From Computation - An Introduction

<details>
<summary>
Strategies to reduce computation.
</summary>

[Running Away From Computation - An Introduction](https://youtu.be/wbnzNWmZ-kU?si=aUoYAZs7YaCan_kR)

we have a lot of trade-offs:

- memory and CPU
- abstraction and performance
- readability
- **time and space**

we can usually trade space to get faster performance, or the other way around. such as the big O notation.

an example of linked list implementations, one that iterates over all the nodes, and one that stores the tail node and has a quicker `append` operation.

we have another trade off in C++, between compile-time and runtime. (also the link time). maybe there are runtime optimizations we can also apply to compile-time?

examples of runtime optimizations:

- using better algorithms - like the linked list with quicker `append` to reduce redundant computations
- do less computations - using short circuiting to avoid expensive operations

we can divide them into micro-optimizations - hand tunning the code, removing dead code, extracting common code, using quicker instructions, etc`. there are also Macro optimizations - re thinking our design and the data structures.

the difference between <cpp>std::map</cpp> and <cpp>std::unordered_map</cpp>: tree based and sorted vs map based and unsorted. we can use the better data structure for our use-case if we know the differences. another option is to delay the computation until we're sure we need it, in C++ we use <cpp>std::promise</cpp> and <cpp>std::future</cpp> and the higher level of <cpp>std::launch::deferred</cpp> execution policy. when (and if) we need the value, we can call the `.get()` method and then block to wait for the computation to be completed. a similar concept is "copy on write" (sometimes called lazy-initialization), which doesn't make actual copies until something is changed.

### Compile-Time Optimizations

There are things we can control at compile time, if we allow the compiler to optimize, it can remove dead code itself, and extract common sub expressions itself.

```cpp
int global;

void passByPointer(const int* p){
   global += *p;
}

void passByReference(const int& p){
   global += p;
}
```

The instructions are the same, but reference must point to a value, while a pointer can be null. so even if the assembly is the same, we usually prefer references, and we don't want to manually check for null pointers. this is part of the core guidelines: **Use References**.

we want to discover bugs and errors as early as possible, like at compile time. we can use <cpp>static_assert</cpp> to check things at compile time, there is no cost at runtime (unlike runtime <cpp>assert</cpp>).\
Moving forward from that, we can have compile time expressions <cpp>constexpr</cpp>, which must be evaluated at compile time. the classic example is factorial function, we can run it at compile time (if we know what value we want) and then we don't need to calculate it at runtime. if we have data that we want to use, we can embed it into the executable directly.

the last example is template meta-programming. we can reduce runtime complexity by paying more at compile time and binary size.

</details>

## Kevin Carpenter :: Being RESTful with Billions of Dollars in Transactions

<details>
<summary>
Case study for making http requests in C++.
</summary>

[Being RESTful with Billions of Dollars in Transactions](https://youtu.be/KIpUrDUa-vw?si=LEOJcVukxofuva0T), [restful-with-billions github](https://github.com/kevinbcarpenter/restful-with-billions).

using header only libraries for http requests and RESTful APIs.

> What is Rest?
>
> - Dissertation of _Roy Fielding_
> - Representational State Transfer
> - High level rules only - lower level implementation is not specified
> - Constrains:
>   - Client Server Architecture
>   - Uniform Interface - consistent, well-defined, endpoints.
>   - Stateless - each operation should conclude and be done (avoiding session management), state management should be done by the client, not the server.
>   - Cacheable - can be done by an intermediate layer.
>   - Layers System - system should be layered (end points, backend server, database).
>   - Code On Demand (optional) - could return binary data (maybe to update a terminal)

in the Electronic payments world, the clients can be smartphones and browsers for online shopping, but there are also traditional client such s registers and terminals.

| HTTP Verb | Crud Operation | URI                   | Payload | Result          |
| --------- | -------------- | --------------------- | ------- | --------------- |
| GET       | Read           | /batch/{batchId}      | empty   | returns Json    |
| POST      | Create         | /sale                 | Json    | Create record   |
| PUT       | Update         | /void/{transactionId} | Json    | Updates record  |
| DELETE    | Delete         | /sale/{transactionId} | Empty   | 405 not allowed |

> - upgrading existing XML and legacy systems, adding modern JSON/REST API
> - Previously using 0MQ - did we change? why?
> - Header only please! why it matters in our environment
> - Performance Considerations
> - <cpp>Nlohmann::json</cpp> - pros and cons
> - <cpp>Cpp-httplib</cpp> - pros and cons

Using Json over XML - json is humanly readable, and saves a bit in size (around 20%), but in large volumes, the difference adds up and is better for older terminals with limited bandwidth options. there are no comments in json, not error handling, no date type, and it's not as robust as XML.

Headers only libraries are used because they are easier to follow and have minimal decencies, and if the library stops getting updates, then the team must be able to keep marinating it locally (at least for a while) and make small customizations.

### C++ REST

- basic HTTP server
- REST Practices!
- creating HTTP client
- Lessons Learned

detaching a thread to run the server, passing a configuration file (json), listening on a host and port, and setting up routing (also pre-routing for security,post-routing and error handling). authentication and authorization. json serialization, test example. (live demo).

choosing between singleton and injection. using concrete types to bridge between json files and writing typed code.

</details>

## Coral Kashri and Daisy Hollman :: From a Modern to an Unbelievably Modern C++

<details>
<summary>
Showing why it's advised to move toward the newer standards.
</summary>

[From a Modern to an Unbelievably Modern C++](https://youtu.be/3ZWYrlmA5g4?si=YE-z1dd8ZucNPt8Z)

reasons to move from a "modern" standard (11/14/17) to a "more modern" one (17/20/23). many new features, less bugs, better optimizations, shorter development time. this talk will show code comparison, and give a roadmap for migrating and moving forward to a newer version, and also introduce some nice c++23 features.

### Code Comparisons

showing how the new standard makes writing code easier and safer.

#### Example 1: Extracting Values from Pair/Tuple.

```cpp
std::map<std::string, std::string> my_map;
// C++11/14
for (std::pair<const std::string, std::string>& key_val : my_map) {
   auto& key = key_val.second;
   auto& val = key_val.first; // oops! this is a bug
   // some magic with key & val
}
// C++17
for (auto& [key, val]: my_map) {
   // some magic with key & val
}
```

C++17 added the structured binding concept, which we can use for any <cloud>std::tuple</cloud> return type.

```cpp
std::tuple<int, double, std::string> func() { return {42, 4.2, "*"}; }
auto [i, d, s] = func();
```

#### Example 2: if statements

```cpp
// C++11/14
template <typename ContT>
void my_func(ContT &container, const typename ContT::value_type &value) {
   auto it = std::find(container.cbegin(), container.cend(), value);
   if (it != container.cend()) {
      std::cout << "The value " << value << " exists in container\n";
      // func_when_value_exist(container, value);
   } else {
      std::cout << "The value " << value << " doesn't exists in container\n";
      // func_when_value_does_not_exist(container, value);
   }
   // `it` continues to exists in the scope
   container.emplace_back(value + 1);
   // now te iterator might be invalidated, depending on the container type
}

// C++17
template <typename ContT>
void my_func(ContT &container, const typename ContT::value_type &value) {
   if (auto it = std::find(container.cbegin(), container.cend(), value); it != container.cend())
   {
      std::cout << "The value " << value << " exists in container\n";
      // `it` exists here
   } else {
      std::cout << "The value " << value << " doesn't exists in container\n";
      // `it` exists here
   }
   // `it` doesn't exist anymore
}
```

in C++11/14, the iterator still exists, so if the container is changed, it might be invalidated. we could use an inner scope to make sure the iterator is no longer accessible, but in C++17 we got initializers inside `if` and `switch` statement.

#### Example 3: If Statement on Compile Time Information

```cpp
struct Number {virtual void inc() = 0;};

// C++11/14
// bad code!
template<typename T>
void func(T &t) {
   // Runtime if-else condition on compile time information
   // both branches should be able to perform the same commands
   // which means the following code wo't compile for arithmetic types
   if (std::is_arithmetic<T>::value) {
      ++t;
   }
   else if (std::is_base_of<Number, T>::value) {
      t.inc();
   }
   std::cout << "I am here\n";
}

// working code, SFINAE
template<typename T, std:: enable_if_t<std::is_arithmetic<T>::value>>
void func(T &t) {
   ++t;
   std::cout << "I am here\n";
}

template<typename T, std:: enable_if_t<std::is_base_of<Number,T>::value>>
void func(T &t) {
   t.inc();
   std::cout << "I am here\n";
}

// C++17
// this simply works now, like wanted before.
template<typename T>
void func(T &t) {
   if constexpr (std::is_arithmetic<T>::value) {
      ++t;
   }
   else if constexpr (std::is_base_of<Number, T>::value) {
      t.inc();
   }
   std::cout << "I am here\n";
}
```

compile time <cpp>if constexpr</cpp> allow us to make decisions on compile-time information and write simpler code without abusing SFINAE for some cases.

#### Example 4: Unions

using a <cpp>union</cpp> can be undefined behavior if used inside a type with a constructor or destructor.

```cpp
struct MyStructure {
   int a;
   double b;
};

// C++11/14
union myUnion {
   int a;
   double b;
   MyStructure ms;
};

class MyUnionHolder {
   enum Types {
      a, b, ms, none
   };
   Types current_type;
   MyUnion m;

   public:
   void set_a(int val) { m.a=val; current_type = Types::a; }
   void get_a(int val) {
      if (current_type == Types::a) {
         return m.a;
      } else {
         // what to do here? throw? do nothing? crash?
      }
   }
   // more getters and setters
};

// C++17
std::Variant<int, double, MyStructure> my_variant;
my_variant = 5;
int res = std::get<int>(my_variant);
try {
   double d = std::get<double>(my_variant); // throws
} catch (std::bad_variant_access const &ex) {
   std::cout << ex.what() << : " my_variant contained int, nou double \n"; // ex.what() -> "Unexpected index"
}
```

C++17 added <cpp>std::variant</cpp> as an alternative to `union`, with clear defintions and the ability to use <cpp>std::visit</cpp>.

```cpp
template<class... Ts> struct overloaded: Ts... { using Ts::operator()...; };
// explicit deduction guide (until C++20):
template<class... Ts> overloaded(Ts...) -> overloaded<Ts...>;

std::Variant<int, double, MyStructure> my_variant;
std::visit(overloaded {
   [] (auto& val) { std::cout << val << '\n'; },
   [] (MyStructure ms) { std::cout << ms.a << " " << ms.b << '\n'; }
}, my_variant)
```

#### Example 5: There is a Return Value

Forcing the user to user the return value with the <cpp>[[nodiscard]]</cpp> attribute.

```cpp
enum Status {
   SUCCESS,
   FAILURE,
   FORCE_EXIT_OR_SOMETHING_TERRIBLE_WOULD_HAPPEN
};

// C++11/14
class MyClass {
   public:
   Status do_something() { Status s; /*...*/; return s; }
};

void my_func(MyClass& mc)
{
   mc.do_something(); // no waring
   // more code
}

// C++17
class MyClass {
   public:
   [[nodiscard]] Status do_something() { Status s; /*...*/; return s; }
};

void my_func(MyClass& mc)
{
   mc.do_something(); // warning, and if -Werror is used then an error
   // more code
}
```

#### Example 6: Sub String

```cpp
// C++11/14
std::string remove_prefix(const std::string& str, const std::string& prefix) {
   return str.substr(str.find(prefix) + prefix.size()); // creates a copy
}

// C++17
std::string_view remove_prefix(std::string_view str, std::string_view prefix) {
   return str.substr(str.find(prefix) + prefix.size()); // a non owning span
}

```

<cpp>std::string_view</cpp> is a non-owning span object that doesn't create a new copy of the original string. we don't need to pass a `const` objects anymore.

#### Example 7: Variadic Templates

```cpp
// C++11/14
// this will look weird if someone passes a string instead
template <typename T>
auto sum14(T value)
{
   return value;
}

template <typename T, typename... Args>
auto sum14(T value, Args... args)
{
   return value + sum14(args...);
}

// other code which is too long to copy to limit this to only integrals
```

unpacking variadic templates in C++11/14 required a templated function with one argument, and a function with variadic arguments. then at compile time there would be multiple functions created which would call one another at runtime. this is massive code bloat. we would need to more code to do conjunction on integral variables. C++17 have fold expressions which are easier to write and don't create as many functions.

```cpp
// C++17
template <typename T>
auto sum17(Args... args)
{
   return (args + ...); // fold expression
}

// only for arithmetics
template <typename T, typename = std::enable_if_t<std::conjunction_v<std::is_arithmetic_v<Args>...>>>
auto sum17(Args... args)
{
   return (args + ...); // fold expression
}

// even easier!
template <typename T, typename = std::enable_if_t<(std::is_arithmetic_v<Args> &&...)>>
auto sum17(Args... args)
{
   return (args + ...); // fold expression
}
```

#### C++20 examples

<cpp>std::span</cpp>

```cpp
std::vector<int> vec = {1,2,3,4,5,6,7,8,9};
// C++11/14/17
std::vector<int> sub_vec(vec.begin() + 2, vec.end() - 2); // copy
auto start = vec.begin() +2 ; // not copying, but here are two more object to keep track of!
auto end = vec.end() - 2;
// C++20
std::span<int> sub_vec(vec.begin() + 2, vec.end() - 2); // no copy, only pointers
```

in C++20, we got <cpp>std::span</cpp> - which do owning spans, and we can also send the span to a function without defining the type. which means we can use <cpp>std::array</cpp> without specifying the size (no creating a new function instance for each size)

```cpp
void func(std::span<int> cont) {/*...*/}
```

<cpp>std::ssize</cpp>

```cpp
// C++11/14/17
for (auto i = vec.size() - 1; i >= 0; --i) { // oops, underflow
   std::cout << vec[i] << ', ';
}

for (auto i = vec.size(); i > 0; --i) { // no underflow
   std::cout << vec[i - 1] << ', '; // good luck remembering why this was used
}

// C++20
for (auto i = std::ssize(vec) - 1; i >= 0; --i) { // oops, underflow
   std::cout << vec[i] << ', ';
}
```

we can get the size of containers with <cpp>std::ssize</cpp> and it will be a signed value, so we won't get underflow when it's zero.

concatenating strings:

```cpp
// C++11/14/17
std::ofstream file("FileID_" + file_id + "." file_version + "." + system_version + "." file_ext);

// using the format library
std::ofstream file(fmt::format("FileID_{}.{}.{}", file_id , file_version, + system_version, file_ext));

// C++20
std::ofstream file(std::format("FileID_{}.{}.{}", file_id , file_version, + system_version, file_ext));
```

C++20 adopted the formatting library and made it into part of the standard.

nodiscard with description:

```cpp

// C++20
enum Status {
   SUCCESS,
   FAILURE,
   FORCE_EXIT_OR_SOMETHING_TERRIBLE_WOULD_HAPPEN
};

class MyClass {
   public:
   [[nodiscard("Possible status might indicate to shutdown the program/pc/office")]]
   Status do_something() { Status s; /*...*/; return s; }
};

void my_func(MyClass& mc)
{
   auto s = mc.do_something();
   if (s == Status::FORCE_EXIT_OR_SOMETHING_TERRIBLE_WOULD_HAPPEN) {
      exit(1);
   }
   // more code
}
```

explain in the warning/error why the return must be used.

<cpp>concepts</cpp> and `<cpp>requires</cpp>`

```cpp
template <typename T>
concept Arithmetic = std::is_arithmetic_v<T>;

template <Arithmetic... Args>
auto sum20(Arithmetic&& ... args)
{
   return (args + ...);
}
```

making the constraints visible and clear. we can replace virtual interfaces with concepts, which makes the inheritance chains shorter, and don't require declaring the interfaces from the start.

### Moving Forward

migrating to a newer version. we want to maintain backward compatiblity. the standard doesn't remove capabilities easily, the list is short and anything that was removed is already a problem if somebody uses it.

- trigraphs - did you even hear about this?
- <cpp>std::random_shuffle</cpp> - was buggy
- <cpp>std::auto_ptr</cpp> - didn't work
- comma operator within subscript operator
- iterator class - deprecated, not removed.

they don't remove stuff that will break code, unless the code was already broken.

using new compilers allows to have new optimizations, which means faster code. so unless there are reasons to use an older compiler (like a custom made one, or with a dedicated hardware), then we should move forward.

```cpp
// C++17
std::vector<int> vec = {1,2,3,4};
std::cout << vec[1]; // prints 2
std::cout << vec[1, 2]; // this prints 3\
// C++20
std::cout << vec[(1,2)]; // C++20 onwards
```

this was deprecated to allow overloading the subscript operator for multiple parameters (like for matrix)

### Other Features

the above examples used C++17 features:

- structured bindings
- `if` statement wit initializer
- <cpp>if constexpr</cpp>
- <cpp>std::variant</cpp> and <cpp>std::visit</cpp>
- <cpp>[[nodiscard]]</cpp>
- <cpp>std::string_view</cpp>
- fold expressions

but there are also other features, such as:

- Guaranteed copy elision
- Class template argument deduction
- Non-type template parameters declared with auto
- Simplified nested namespaces

other c++20 features:

- <cpp>consteval</cpp>
- Designated initializers
- <cpp>[[likely]]</cpp> and <cpp>[[unlikely]]</cpp>
- <cpp>[[no_unique_address]]</cpp>
- Modules
- Coroutines
- Three way comparison operator (spaceship `<=>`)

C++23 features:

- <cpp>std::mdspan</cpp>
- deducing <cpp>this</cpp>
- <cpp>std::flat_map</cpp> and <cpp>std::flat_set</cpp>

</details>

## Alex Cohn :: Does The C++ Compiler Work Too Hard?

<details>
<summary>
Case study of analyzing a complete build of the Unreal Game Engine.
</summary>

[Does The C++ Compiler Work Too Hard?](https://youtu.be/IjstIoRM7MI?si=CFVCbwDYNQVN5HJu)

showing a case study about the **Unreal Engine** and how long it takes to build, with an optimized build taking 5 hours on a normal machine. build speed can increase by reducing the number of compilation units - unity build (this has massive effects), using precompiled headers and there is a slight (very small) effect of using different compilation flags (`-Oz`, `-Os`).\
Of course, compilation time isn't the only consideration, we want the program to run efficiently.\
there is a correlation between the time it takes to optimize each unit and the size of the resulting object. The effect of unity build (combining multiple files into one) seems to be that it reduces the number of times the compiler initializes. The "sqlite" library seems to be an outlier, it takes a long time to build for a relatively small resulting object size. it's actually 100 C files, so if they are combined into a single compilation unit, there can be some speed gains.

</details>

## Eran Talmor :: Scope Sensitive Programming

<details>
<summary>
A way of programming that relies on scope based thread specific objects.
</summary>

[Scope Sensitive Programming](https://youtu.be/UBTvUY9IEsA?si=iIXCU8qUz36c18fG), [scoped github](https://github.com/erangithub/scoped).

The motivating problem

> - A computational SW
> - Graph in the infrastructure layer
> - Multiple layers of business logic
> - Big API layers
> - Script thread: heavy, lengthy computations
> - GUI thread: light, short computations
>
> How do we set different threshold from the Scripting and GUI threads?

both the Scripting thread and the GUI thread want to use the graph infrastructure.

there are commonly prescribed solutions: either setting the threshold with an direct API, passing the threshold argument as part of the call, or changing the architecture entirely to support 'per-thread graph views'. each of the solutions has it's own issues.

The lecture suggests using a `scoped<>` template ("hacking the stack"). wrapping a value with a tag to differentiate them between instances. they must be put on the stack.

```cpp
template<class T, class... tags>
class scoped {
/* ... */
};

using ScopedThreshold = scoped<int, struct ScopedThresholdTag>;
```

> - **Doubly-Linked list** embedded in the call-stack.
> - Each template instance of `<data, tags...>` is separate.
> - **Thread-local** static pointer to the current top or bottom maintained by the constructor and destructor.

```cpp
void f1() {
   ScopedThreshold thresh(30);
   f2();
}

void f2() {
   f3();
}

void f3() {
   ScopedThreshold thresh(60);
}
```

each thread sees a unique doubly linked list.

```cpp
using ScopedThreshold = scoped::scoped<int, struct ScopedThresholdTag>;

void print_number(int x) {
   std::cout << "The number is ";
   if (auto thresh = ScopedThreshold::top())
   {
      if (x >= thrash->value())
      {
         std::cout << "BIG" << std::endl;
         return
      }
   }
   std::cout << x << std::endl;
}

int main()
{
   {
      ScopedThreshold scoped_threshold{4};
      print_number(3); // Expected: "The number is 3"
      print_number(10); // Expected: "The number is BIG"
   }
   print_number(10); // Expected: "The number is 10"
   return 0;
}
```

we use a static (thread specific) function to check if function is inside a scoped stack, and then we can use it.

Caching

- Many computations modules enjoy caching
- Caching best determined by the consumer of a function/module
- `scoped<>` can help set caching at different levels/scopes
- in our example `is_prime()` prefers the outer most scope.

```cpp
// Create a new scoped cache for caching prime numbers
using ScopedPrimeCache = scoped::scoped::<std::unordered_map<int, bool>, struct ScopedPrimeCacheTag>;

// check if a given number is a prime number
bool is_prime(int n)
{
   if (n<2) return false;
   // Retrieve the prime number cache, taking the bottom (outer) most scope
   auto pScopedCache = ScopedPrimeCache::bottom();

   // if previously cached, return the result
   if (pScopedCache) {
      auto & cache = pScopedCache->value();
      if (auto it = cache.find(n); it != cache.end()) {
         std::cout << "Cache hit for " << n << std::endl;
         return it->second();
      }
   }

   // calculate wether the number is a prime or not, and cache the result
}

// Find the next prime number greater than a given number
int next_prime(int n) {
   int k = n + 1;
   for (; is_prime(k); k++;);
   return k;
}

// Find the first n prime numbers
std::vector<int> first_n_primes(int n) {
   std::vector<int> primes;
   int p = 0;
   for (int i = 0; i < n; ++i) {
      p = next_prime(p);
      primes.push_back(p);
   }
   return primes;
}

int main()
{
   // Create a new scoped cache for caching prime numbers
   ScopedPrimeCache primeCache;

   first_n_primes(5);
   for (auto p : first_n_primes(10)) // 5 cache hits
   {
      std::cout << p << std::endl;
   }
}
```

we can use this for event Counting, such as collecting usage statistics, internal code, external users. we can count events inside a function or inside a thread.

(example of a calculator class)

another use case is with decorators (and dependency injection), such as decorating log messages. each scope can add decorations to the message, and they are applied sequentially. and since this is thread specific, there is no problem of interference.

There are of course a few drawbacks to the solution, mostly that it introduces invisible affects to functions. for that reason, there is the option to use a _scoped::shield_ object. it acts the same as any other scope, but it sets the top and bottom scopes to null at creation, and resets them back to what they were at the destructor.

```cpp
int main()
{
   {
      ScopedThreshold scoped_threshold{4};
      print_number(3); // Expected: "The number is 3"
      print_number(10); // Expected: "The number is BIG"
      {
         ScopedThreshold::shield shield;
         print_number(10); // Expected: "The number is 10"
      }
      print_number(10); // Expected: "The number is BIG"
   }
   print_number(10); // Expected: "The number is 10"
   return 0;
}
```

another option is adding a "manifest", so the scopes gain more visibility. this is done via macros. this allows attaching scopes to specific objects, classes and functions.

> **Pros**:
>
> - flexible, supporting multiple use-cases
> - Thread Safe
> - Fast - no locks
> - No need to set and reset configurations
> - Hide low-level details from header files
> - Polymorphic
> - Easy to apply to existing code
> - Changes the code, not architecture
> - Intuitive?
> - Easy to maintain
>
> **Cons**:
>
> - Provide hidden knobs / side channel
>   - "Manifests" may help
> - Code prone to External affects
>   - "Shields" may help
> - If not used carefully, may lead to bad coding
> - Don't use instead of arguments

</details>

## Gal Oren :: Accelerated C++ with OpenMP

<details>
<summary>
Many-Core programming with OpenMP.
</summary>

[Accelerated C++ with OpenMP](https://youtu.be/kxN2JOxrwzs?si=9udHZ67ifINThDZ7)

Parallel computing helps us increase our computation power by scaling across multiple instances and not by just building stronger machines.

"Dennard Scaling" is an observation (similar to "Moore's law") that it's possible to increase the number of transistors in a given space and still have the same power efficiency, thanks to technological advancements. this was the case until the early 2000's, now the power consumption increases more than the computing capabilites. this gives the raise for multi-core programming.

Many-Core computing is the idea of reducing the power of each invidual core, but spreading the work across many of them.

Graphical Processing Units (GPU):

- The strength of a GPU lies in the massive parallelism it offers.
- Each compute unit in the GPU is not very powerful, but there are many of such units.

this works for simple and single instructions, which are repeated across large data sets. in recent years, more than half of the super computers in the world rely on GPUs to dome large part of their computing.

we need to write programs that can take advantage of those capabilites.

> programming models:
>
> - OpenMP
> - OpenCl
> - OpenACC
> - CUDA
> - SYCL

### OpenMP High Performance Computing

moving from multi-core to many-core. OpenMP is hardware agnostic, vendor agnostic, support incremental "upgrading" to parallelism. it exposes a runtime library on the system layer with support for shared memory access and threading on the hardware.

a simple model is "fork-join" parallelism, in which some sections of the code are parallelized, with the number of threads increasing as long as there are performance gains.

Calculating Pi

$$
\begin{align*}
\int_0^1 \frac{4.0}{1+x^2}dx = \pi \\
\sum\limits_{i=0}^{n}F(X_i)\Delta \approx \pi
\end{align*}
$$

serial code to calculate &pi;

```c
static long num_steps = 100000;
double step;

int main()
{
   int i;
   double x;
   double pi;
   double sum = 0.0;
   step = 1.0 / (double) num_steps;
   for (i = 0; num_steps < n; i++) {
      x = (i + 0.5) * step;
      sum = sum + 4.0/(1.0 + x * x);
   }
   pi = step * sum;
}
```

parallel code with OpenMP loop & reduction

```c
#include <omp.h>
static long num_steps = 100000;
double step;

int main()
{
   int i;
   double pi;
   double sum = 0.0;
   step = 1.0 / (double) num_steps;
   #pragma omp parallel // create a team of threads
   {
      double x; // create a scalar local to head thread to hold the value of x at the center of each interval
      #pragma omp for reduction(+:sum)
      for (i = 0; num_steps < n; i++) {
         x = (i + 0.5) * step;
         sum = sum + 4.0/(1.0 + x * x);
      }
   }
   pi = step * sum;
}
```

later, OpenMP started working with the concept of "Tasks", attaching threads to a workload. then it started supporting offloading workloads to other devices. this meant that the memory isn't necessary shared across all devices, some can be in the CPU and can some can be GPU with separate memory space.

default data sharing example, the stack arrays are copied to the device and the calculation is executed there, and then they are copied back.

```c
int main(void)
{
   int N = 1024;
   double A[N];
   double B[N];

   #pragma omp target
   {
      // ii is private on the device since it's being declared within the target region
      for (int ii = 0; ii < n; ii++) {
         A[ii] = A[ii] + B[ii]
      }
   } // end of target region
}
```

we can choose to run asynchronously and launch the target without calculation without waiting for it to complete, and explicitly wait for it at a later point. for dynamic memory we need to explicitly map memory for it to correctly copy the data (we can specify which data we copy from the host and back from the device to get better performance). we can combine device parallelism and thread parallelism (and then use SIMD). there are many additional `#pragma omp` options to get granular. but in recent versions of OpenMP we can get really good defaults from the simple `#pragma omp loop`.

(demo of running code on two different GPUs)

</details>

## Alon Wolf :: Expressive Compile Time Parsers

<details>
<summary>
Creating compile time parsers for custom text
</summary>

[Expressive Compile Time Parsers](https://youtu.be/g_GJ_a_lxLg?si=IY5rBvxucSkzGS3i)

> "Expressive Code" refers to the style of writing code in a way that is easy to read, write, understand and communicate its purpose. Relies on both the syntax of the programming language and the quality of naming conventions.

the definition of expressive code changes as the language evolves, with each version allowing for more concise code and less boiler-plate. operator overloading can be a part of expressive code.

DSL - domain specific language provides specialized vocabulary and syntax that are aligned with the domain. the dsl syntax has to be legal C++ code, which limits how the DSL can change.

(example of Boost Spirit DSL code for writing runtime compilers)

example of using a filter transformation in c++11 vs c++23.

```cpp
std::vector<Cat> cats = {/* ... */};
std::vector<std::tuple<int, std::string>> result;

for (auto itr = cats.cbegin(); itr = cats.cend(); ++itr)
{
   if (itr-> age < 42)
   {
      results.emplace_back(std::tuple{itr->id, itr->name});
   }
}

// modern approach
auto view =  cats |
   std::views::filter([] (const Cat& cat) {return cat.age > 42;}) |
   std::views::transform([] (const Cat& cat) {return std::tuple{cat.id, cat.name};});
auto result2 = std::vector(view.begin(), view.end());
```

but what if we could write this in a single line? this could be done by using compile time parsers.

```cpp
std::vector<Cat> cat = {/*...*/};
auto results = "Select id, name WHERE age > 42"_FROM(cats);
```

in this case, we have a compile time value with a string literal that we transform into a lambda, in the general case, we transform any arbitrary syntax to any compile time value.

### Existing Parsing Libraries

- LL parser - left left, left most deviations, top-down, 'predictive parsers'
- LR parser - left right, right most deviations, bottom-up, 'shift reduce parsers'

turns out there was a boost library for compile time parsing: **Boost Metaparse**, it was created for C++98, and allowed to write and create meta functions from raw string using template parsers.

a modern library is [Lexy](https://lexy.foonathan.net/), a parser combinator for C++17 and onwards. it supports unicode strings and compile-time parsing. a parser has a rule for constructing and value. Lexy also has the option for multiple composable parsers.

Another C++17 library is **Compile Time Regular Expressions** (CTRE). it allows for regular expressions at either compile or runtime. it constructs the regex during compile time (error when regex is illegal), it also performs much better than the standard <cpp>std::regex</cpp>, and similarly to other runtime libraries such as boost.

```cpp
auto [v1, a1] = "REGEX"_ctre.match(s); // C++17 with N3599
auto [v2, a2] = ctre::match<"REGEX">(s); // C++20
```

another libary is **Compile time Parser Generator** (CTPG), which generates parsers from a grammar, it itself uses a self-generated regex parsers.

**Macro Rules** was originally created for C++23, but was converted to C++20, it describes DSL with Rust macro style rules. it is more experimental than other libraries (proof of concept rather than usable library). it uses macros other than string literals, and it has hashed identifies for faster lookups.

### Reflection

C++ still doesn't have reflection, only several introspection features. we could use compile time parsers to parse the source code itself and generate the metadata for reflection. the generated string is valid C++ source code.

```cpp
// macro style reflection
struct Obj {
   int a = 42;
   int b = 1337;
   REFLECT(MyObj, a, b);
};

// parsing style

template<StaticString src>
constexpr auto Reflect()
{
   constexpr auto attributes = attributes_parser(src);
   constexpr auto base_classes = bases_parser(src);
   return std::tuple(attributes, base_classes);
}
```

example [compiler explorer](https://godbolt.org/z/xzeT7rM7Y): we want to match a member inside a value with dot separated string. each level is an identifier (member value) of the object we want to "reflect" onto.

```cpp
// desired syntax
auto email = "manager.details.email"_in(company);

// parser
constexpr auto identifier = ... >>= [] (auto sv) {return Hash(sv);};
constexpr auto path_parser = separate_by(identifier, "."_lit);

// reflect member with identifier
static constexpr auto ResolveIdentifier(ValueWrapper<Hash("x")>) {
   return &S::x;
}

struct S{
   int x;
   int y;
   REFLECT_MEMBERS(S, (x)(y));
};

// stirng literal _in operator
template<StaticString path_str>
constexpr auto operator""_in(){
   return [] (auto& value)->decltype(auto) {
      constexpr auto path = path_parser(path_str); // parse at compile time
      return GetPath<path>(value);
   };
}

// recursive iteration of path and get member by identifier
template<auto path>
constexpr auto GetPath(auto& value){
   return RecursiveFor([] (auto& value, auto idx, auto next)->decltype(auto) {
      if constexpr (idx.Value() == path.m_Size) {
         return (decltype(value)&)value;
      } else {
         constexpr auto ident = path.m_Data[idx.Value()];
         return next(GetMember<ident>(value));
      }
   }, value);
}
```

if we look at the generated assembly, we see that both ways of accessing the data have the same assembly instructions, so this becomes a zero-cost abstraction (zero cost at runtime, of course) we the `-O1` optimization flag

```cpp
auto& F1(const Line& line){
    return line.p2.x;
}

auto& F2(const Line& line){
    return "p2.x"_in(line);
}
```

we might wish to extended this syntax, and allow something such as: `"players:character.items:textures:height|sqs|print"_of(game);`. with the "`:`" being a range based for loop and "`|func`" being a function call with the current value. so the line above:

1. takes the "players" field from the "game" object
1. iterates over all the players
1. takes the "character" field
1. takes "items" field for the character and iterates over all the items
1. takes the "textures" field for the item and iterates over it
1. takes the "height" field
1. squares the value
1. prints the values.

we create parsers for the "iterate" and "pipe" operators, and now we have a "scope". the scope maps an identifier to a value (usually a lambda), it can be the global scope or a local scope.

```cpp
constexpr auto iterate = ":"_lit >>= [] (auto) {
   return IterateTag{};
};
constexpr auto pipe = ("|"_lit + identifier) >>= [] (auto ident) {
   return Pipe{ident};
};
constexpr auto custom_parser = *(path_parser | iterate | pipe);

// extended implementation for path
template<auto path>
constexpr auto CustomPath(auto& value, Scope& scope = GlobalScope()) {
   return RecursiveFor([] (auto& value, auto idx, auto next)->decltype(auto) {
      if constexpr (iterate) {
         for (auto& v : value) next(v);
      } else if constexpr (path) {
         return next(GetPath<*path>()(FWD(value)));
      } else {
         constexpr auto ident = ValueWrapper<pipe->m_Action>();
         return next(scope(ident)(FWD(value)));
      }
   }, value);
}

// scope
constexpr auto sqr = [] (auto&& x) {return x*x;};
constexpr auto ResolveIdentifier(Vw<Hash("sqs")>) {
   return sqr;
}
#define SCOPE [] (auto v) {return ResolveIdentifier(v);}
using GlobalScope = decltype(SCOPE);

// local scope example

```

so the string of `"players:character.items:textures:height|sqs|print"_of(game, SCOPE);` is the same as:

```cpp
for (auto& player : game.players)
   for (auto& item : player.character.items)
      for (auto& texture : item.textures)
         print(sqr(texture.height));
```

but now it's being read left to right, has abstracted away loops and the boiler plate code. and we can see the same assembly code in [compiler explorer](https://godbolt.org/z/fP6cose6q), meaning that this is a zero-cost abstraction again.

next we want to create a struct, rather than a lambda. our parser should take a syntax such as:

```text
STRUCT(Point, {
   x: float
   y: float
})
```

and make it into a valid C++ struct.

the parser will look like this

```cpp
constexpr auto struct_parser = "{"_lit + *(identifier + ":"_lit + identifier) + "}"_lit;
template<auto Id, class T>
struct StructMember {
private:
   T m_value;
public:
   static constexpr auto ResolveIdentifier(Vw<Id>) {
      return &StructMember::m_value;
   }
   constexpr auto& operator[](Vw<Id>) {return m_Value;}
   constexpr auto& operator[](Vw<Id>) const {return m_Value;}
};

// more code examples, multiple members through folding composition, parsing into the actual struct
```

the limits on using compile time parsing is the compilation speed and the limits on `constexpr`, in theory, we could write parsers for every programming language and have them integrated into the C++ code.

</details>

## Victor Ciura :: The Imperatives Must Go!

<details>
<summary>
Functional Programming in C++
</summary>

[The Imperatives Must Go!](https://youtu.be/WdWwbDJWx78?si=OGQeF19zw8zlHO4n), [slides](https://ciura.ro/presentations/2023/Conferences/The%20Imperatives%20Must%20Go%20-%20Victor%20Ciura%20-%20Core%20C++%202023.pdf), [lift github](https://github.com/rollbear/lift)

Humans and machines operate differently, and when we write software, we try to "think like a machine" and create a mental model that replicates what the computer does.

> What is Functional Programming ?
>
> - Functional programming is a _style_ of programming in which the basic method of computation is the _application of functions_ to arguments.
> - A functional language is one that supports and encourages the functional style

the elephant (&#x1F418;) in the room is the **Haskell** programming language.

in imperative style, we tell the machine how want the computation to be performed, and our main "tool" is _variable assignment_.

```cpp
int total = 0;
for (int i = 1; i <= 10; i++) {
   total = total + i;
}
```

in contrast, for Haskell, we use **function application**.

```Haskell
sum [1..10]
```

in other words, imperative style is about the "how", while functional style is about the "what". if we compare to OOP, then OOP minimizes complexity by encapsulating the moving parts into classes, and functional programming minimizes complexity by reducing the amount of the moving parts.

### Historical Background

we can connect functional programming to lambda calculus from the 1930's, and the first functional programming language (**Lisp**) was developed in the 1950's. it took inspiration from lambda calculus, but retain variable assignments. then we get **ISWIM**, which is the first pure functional programming language. in the seventies we start looking at _higher order functions_, _type inference_ and _polymorphic types_. then we add _lazy computation_ and evaluations. all of these reach the **Haskell** language, which then get the innovation of _type classes_ and _monads_.

despite that, Haskell didn't take over the world, at least not directly. however, other languages are moving towards functional programming, taking inspiration from Haskell and incorporating elements from it.

- lambdas and closures
- <cpp>std::function</cpp>
- values
- ADT
- composable algorithms
- lazy ranges
- folding
- mapping
- partial applications
- higher order functions
- monads
- <cpp>std::optional</cpp>, <cpp>std::future</cpp>, <cpp>std::expected</cpp> (c++23)

here is an example of haskell code.

```Haskell
f [] = []
f (x: xs) = f ys ++ [x] ++ f zs
            where
               ys = [a | a <- xs, a <= x]
               xs = [b | b <- xs, b > x]
```

this is actually a quick sort algorithm, it recursively applies the algorithm and combines the result. starting with the first element as the pivot.

```Haskell
qsort [] = []
qsort (x: xs) = qsort smaller ++ [x] ++ qsort larger
            where
               smaller = [a | a <- xs, a <= x]
               larger  = [b | b <- xs, b > x]
```

the computations is done like this:

```na
q[3,2,4,1,5]
q[2,1] ++ [3] ++ q[4,5]
q[1] ++ [2] + q[] ++ [3] ++ q[] ++ [4] ++ q[5]
[1] ++ [2] + [] ++ [3] ++ [] ++ [4] ++ [5]

[1,2,3,4,5]
```

once we learn how to read this code, it might be easier to understand than C language quickSort source code.

an historical example is a task of reading a text file, and outputting a sorted list of words and their frequencies. a solution submitted by Donald Knuth in **Pascal** was ten pages long, and a rebuttal by Douglas Mcllroy done in a shell script was only six lines.

```sh
tr -cs A-Za-z '\n' |
   tr A-Z a-z |
   sort |
   uniq -c |
   sort -rn |
   set ${1}q
```

### The Journey for Functional Programming

one option to learn functional programming is learning category theory. but there are also more practical approaches that focus on software rather than mathmatical theory.

we now move to more examples: higher-order function, "lifting" and "boxing".

<cpp>boost::hof</cpp> is a library that uses higher order functions, there is also the _lift_ library which is simpler.

- equal
- not_equal
- less_than
- less_equal
- greater_than
- greater_equal
- negate
- compose
- when_all
- when_any
- when_none
- if_then
- if_then_else
- do_all

a simple code example for using the library, a composition of function wraps them over one another. it also allows us to "compose" an overloaded set of functions using a macro.

```cpp
struct Employee {
   std::string name;
   unsigned number;
};

const std::string& select_name(const Employee& e) { return e.name; }
unsigned select_number(const Employee& e) { return e.number; }

std::vector<Employee> staff;

// sort employees by name
std::sort(staff.begin(), staff.end(), lift::compose(std::less<>{}, select_name));

// retire employee number 5
auto i = std::find_if(staff.begin(), staff.end(), lift::compose(lift::equal(5), select_number));
if (i != staff.end()) {
   staff.erase(i);
}
```

C++20 gave us ranges, which could be used in a similar way.

boxing is a way of encapsulation values. we use the terms "Functor", "Applicative" and "Monad".

| method                     | storage            | accesses             |
| -------------------------- | ------------------ | -------------------- |
| <cpp>std::unique_ptr</cpp> | `unique_ptr<T> p;` | `*p`, `p.get()`      |
| <cpp>std::shared_ptr</cpp> | `shared_ptr<T> p;` | `*p`, `p.get()`      |
| <cpp>std::vector</cpp>     | `vector<T> v;`     | `v[0]`, `*v.begin()` |
| <cpp>std::optional</cpp>   | `optional<T> o;`   | `*o`, `o.value()`    |
| <cpp>std::function</cpp>   | `function<T> f;`   | `f(5)`               |

but this might be an anti-pattern, retrieving and storing the data again and again. we want to preserve the context of a value across computations.

we can learn from a common un-wrapping mistake in Rust, and avoid constantly peeking into the box.

this code unwraps the box (the anti-pattern)

```cpp
string capitalize(string str);

std::optional<string> str = f(); // from an operation that could fail

std::string cap;
if (str)
   cap = capitalize(str.value()); // capitalize(*str);
```

a functional way of doing this would be by lifting the capitalizing function, which now takes the box itself, if the value is valid, then it performs the unwrapping ,does the capitalization and boxes it again, otherwise it returns the input without changing it.

```cpp
std::optional<std::string> liftedCapitalize(const std::optional<std::string> & s)
{
   std::optional<std::string> result;
   if (s)
      result = capitalize(*s);

   return result;
}
```

we can extend the idea further, into an `fmap` function.

```cpp
template<class A, class B>
std::optional<B> fmap(std::function<B(A)> f, const std::optional<A> & o)
{
   std::optional<B> result;
   if (o)
      result = f(*o); // wrap a <B>
   return result;
}
// or
template<typename T, typename F>
auto fmap(const std::optional<T> & o, F f) -> decltype(f(o.value()))
{
   if (o)
      return f(o.value());
   else
      return {}; // std::nullopt
}
```

another example about containers holding different types, and <cpp>std::transform</cpp>. if any operation fails (because the input is optional), we can return the default (empty) value. C++23 added monadic operations to <cpp>std::optional</cpp>:

- <cpp>o.and_then()</cpp>
- <cpp>o.transform()</cpp>
- <cpp>o.or_else()</cpp>

| Concept | C++                  | Haskell      |
| ------- | -------------------- | ------------ |
| functor | <cpp>transform</cpp> | `fmap`       |
| monad   | <cpp>and_then</cpp>  | `>>= (bind)` |

we shouldn't use <cpp>std::optional</cpp> for error handling, instead, we can use <cpp>std::expected</cpp> from C++23, which can return another value explaining why we don't have a value.

> expressions yield values, statement do not;

we want to program with expressions and not statements.

as mentioned, C++20 ranges promote functional programming. a simple task of printing only the even elements from a range in revere order.

```cpp
// imperative style
std::for_each(std::crbegin(v), std::crend(v),[](auto const i) {
   if(is_even(i))
      cout << i;
});

// functional style
for (auto const i : v
                  | reverse
                  | filter(is_even))
{
   cout << i;
}
```

> Gotchas with ranges / views:
>
> C++20 ranges library is fantastic tool, but watch out for gotchas
>
> - views have reference semantics => all the reference gotchas apply
> - as always with C++, const is shallow and doesn't propagate (as you might expect)
> - some functions do caching, eg. <cpp>begin()</cpp>, <cpp>empty()</cpp>, <cpp>| filter</cpp>, <cpp>| drop</cpp>
> - don't hold on to views or try to reuse them
> - safest to use them ad-hoc, as temporaries
> - if needed, better "copy" them (cheap) for reuse

</details>

## Yehezkel Bernat :: Lazy and Proud: How I Failed to Standardize lazy_counted_iterator

<details>
<summary>
How do we change the standard?
</summary>

[Lazy and Proud: How I Failed to Standardize lazy_counted_iterator](https://youtu.be/JSM2aKKAH1s?si=6EMn9vCwn4Mw66Zp)

the motivation for this talks come from project [Euler problem #37](https://projecteuler.net/problem=37), reading:

> The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97 and 7. Similarly we can work from right to left: 3797,379,37,3.
>
> Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
>
> NOTE: 2,3,5 and 7 are not considered to be truncatable primes.

this problem can be probably be solved using ranges. let's try it. the solution will probably start with something such as

```cpp
views::iota(10) | // generate numbers starting at 10
    views::filter(/*something here*/) | // filter somehow
    views::take(11); // take first eleven number
```

but when running this in compiler explorer, it seems to continue running even after 11 results, and it time out! the problem is the <cloud>views::take</cloud>
uses <cpp>std::counted_iterator</cpp> internally. which continues to work even if there are no more object to take (we only generated 10 numbers)

simple example:

```cpp
#include <ranges>
#include <iostream>

bool f(int n){
    return n < 5;
}

int main()
{
    auto vec = std::views::iota(0) | // endless generator
               std::views::filter(f) | // numbers smaller than 5
               std::views::take(6); // take six of them
    for (const auto v : vec){
        std::cout << v <<',';
    }
}
```

the bug is also present with input_iterators. an extra incrimination "eats" an element, which is now lost.

```cpp
auto iss = std::istringstream("0 1 2");
for (auto i : std::ranges::istream_view<int>(iss) | std::views::take(1)) {
   std::cout << i << '\n' // print 0
}

auto i = 0;
iss >> i;
assert(i==1); // FAILS,  i == 2
```

fixing a bug in the standard requires changing the implementation and text that defines the standard itself. there are all sorts of functions in the committee, including domain specific study groups(SG9 for ranges), a proposal needs to go through all the stages before it's accepted and becomes part of the standard.

<cpp>lazy_counted_iterator</cpp> was suggested as a new class to avoid breaking the ABI. after writing the inital proposals, some people commented and pointed issues, so the process repeats.

eventually, it was decided to reject the proposal, and deciding that the problem should be addressed differently.

</details>

## Chris Ryan :: Multi-Paradigm Programming and Beyond

<details>
<summary>
What Does Being "Multi-Paradigm" Mean?
</summary>

[Multi-Paradigm Programming and Beyond](https://youtu.be/i1cliYm9qQY?si=Tvfa-I8SqtJWJrxE)

- Imperative
- Procedural
- Functional
- Declarative
- Object Oriented

Paradigm can be the building block, but also be something bigger or more conceptual -

> - a way of doing or thinking about something
> - a way of organizing or managing a process or data
> - a way of showing your intent

- Imperative - Sequencing of statements to describe how the program should execute
- Procedural - Organizes code into functions to perform tasks and manipulate step-by-step
- Functional - Programming techniques, such as high-order functions, and immutability
- Declarative - Describing what needs to be done, rather than how to achieve it
- Object Oriented - Organizing code around objects that encapsulate data and behavior

other than Object Oriented, we could do this in C.

we could put them on a scale, getting more structured as we move forward.\
Imperative -> Procedural -> Functional -> Declarative

we also have generic programming - coding for a wide range of data types, without specialization. Templatization - generate types at compile time, then also execute code at compile time.

there are also other kinds of paradigms: super-loops, state machines (state dispatcher, windows event-driven frames), I/O loops, MFC(microsoft foundation class), MVC(model view controller).

Data models:

- Conceptual - An abstract representation or high level description of as system, process, or domain.
- Logical - Represents the structure and organization of data at a higher-level of abstraction.
- Physical - Represents the implementation-specific details of a system or database.

Entity-Relationship(ER): Conceptual modeling technique used in database design to represent the relationships between entities in a system or domain.

we can say C++ is more of a mixed-paradigm language than multi-paradigm. we can use whatever we want, and combine them together.

</details>

## Mike Spertus :: C++ For The Cloud

<details>
<summary>
Writing C++ code that plays with the cloud
</summary>

[C++ For The Cloud](https://youtu.be/y8D_jMGYxJE?si=ShC9TYrM0c5AjeeE), [AWS High-level Enhancements for C++](https://github.com/awslabs/high-level-enhancements-for-aws-in-cpp)

### The Cloud

- Storage and Compute hosted on the internet, not the computer.
- Sharable with whoever you chooses wherever in the world you choose.
- Infinitely scalable.
- Pat for what you use on demand.
- Managed - no worries about patching operating system, replacing failed disks, power management, etc...
- It's possible to "lift and shift" existing application from the data center to the cloud.
- It also enables new paradigms for structuring applications.
  - assemble serverless application from managed microservices.

> Microservices
>
> - Run separately
> - Can be build and deployed independently by small teams
> - Owns its resources
> - Independent SLA (service level agreement)
> - Optimizes around the cloud elasticity

Microservices have advantages over large monolith executables, they are less tightly coupled, easier to modify and test. we can deploy each microservice on it's own, without deploying the entire application. **Independet Loosely Coupled Microservices**. the "two pizzas" rule - each feature is developed by a team that's small enough that 2 pizzas are sufficient for them.\
there are existing services to act as building blocks for the application, some of them are very high-level and specific (such as <cloud>Amazon Transcribe</cloud> for text-to-speech), while other services are lower-level (compute, storage). we focus on <cloud>S3 bucket</cloud> object storage and <cloud>Lambda</cloud> compute functions (serverless, the user doesn't manage the machine running the code). we can invoke lambda functions either directly (SDK, url, web console) or based on events (many aws services). arguments to lambda function are language-agnostic, they are json objects, which can be schemaless for direct invocations, events triggered by other services have a schema.

### Traditional Approach For Using C++ And The Cloud

AWS has a SDK for C++. All AWS SDKs are generated by [Smithy](https://smithy.io/2.0/index.html). we can see an example for a lambda application written in C++ that takes a value and increments it, and we look at the sample code for invoking the lambda from a program. this example is very complicated with parsing, invoking, and validating, it also uses a very specific interfaces, as opposed to generalized concepts such as streams.

> "In C++, We don't subset, we create abstractions"\
> ~ Bjarne Stroustrup

### High Level Enactments For C++ And AWS

The traditional approach is powerful, but harder to use and creates mental overhead. this means that we will probably only use them when absolutely necessary. but what if we want to move all our code to the cloud - use microservices for all cases, except for when there a reason _not_ to do so? we could get all the benefits of cloud computing for free!

could we turn the classic quote about templates

> "What if generic programming was 'just' normal programming?"
> ~ Bjarne Stroustrup

into a mission statement for cloud computing (in aws)

> "What if cloud programming was 'just' normal programming?"

- High level abstraction over common microservice use cases
- Consistent with the style of the C++ Standard Library

A new library for C++ that does all that.

```cpp
AwsApi api; // RAII class to initialize the AWS API

void write_to_s3()
{
   // write to a bucket just like writing into a file
   os3stream os3s("us-east-1","bucket-name", "object_name"); // output stream to S3
   os3s << test_content;
}
```

writing a lambda that adds two integers.

```cpp
#include "awslabs/enhanced/aws_lambda.h"

int add(int i, int j) { return i + j; }

Handler handle(&add); // set it as the handler request
```

calling the lambda from the executable

```cpp
AwsApi api; // RAII class to initialize the AWS API
EnhancedLambdaClient client;

auto add = client.bind_lambda<int(int,int)>("add"); // mapping function in code to something in the cloud. "add" is the name of the lambda
int i = add(1,3);
```

Each time we call the lambda, a new instance might be created (default maximal concurrency is 1000 instances), this gives us the option for scalability. this can become a distinct approach for shared memory concurrency (like threads and GPUs). The library has a new execution policy for STL algorithms.\
Errors can be handled with c++23 <cpp>std::expected</cpp> (or <cpp>tl::expected</cpp>), if it's not declared as such, the library automatically convert the return value into that and throw the exception as a string in the client code. we can also specify the return type as <cpp>invocation_response</cpp> to get the complete json back. serialization is a customizable plugin, and uses the <cloud>alpaca</cloud> library by default. this is still not perfect, and won't work for all cases, and there will be problems until real reflection is introduced into the language. only pass-by-value will be supported in the forseeable future. there is also some wonky-ness with concurrency.

### What's Next?

> - New services and feature coverage (short term).
> - Leverage more modern C++ language features (short term).
> - Tool support (medium term).
> - Language changes (long term).

an example for a new service is using NoSQL databases (such as <cloud>DynamoDB</cloud>) as "maps in the cloud" and making them as easy to use as <cpp>std::map</cpp>. another future feature is making lambda triggers in C++.\
The tool chain is designed to make an executable, but we want it to assemble an application out of managed services. could we have a linker that links cloud functions? we need better tooling to automate deployment, roles, etc...

Some problems must be handled by the language itself, like execution policies for ranges.

</details>

## Gili Kamma :: Lessons I Learn From Improving Legacy Product And Doubling Its Performance

<details>
<summary>
Case study of improving performance in an old product.
</summary>

[Lessons I Learn From Improving Legacy Product And Doubling Its Performance](https://youtu.be/qIss19hhfTA?si=BrB1aDA8BaVD4mkL)

(Hebrew lecture)

case study, a utility company, reading meters (water, electricity) online. the original design was to work with 1000 end-units, and it started failing when communicating with 5000 devices, and in some cases, it had to work with more than that. so something had to be done.

The stack was very old school, running on a linux kernel with a specific distribution, and used <cpp>QT</cpp> application wrapper.

The problems were never observed in the lab, and could happen when there many connected devices or there were network errors. the goal was to reduce loss of data (which means loss of revenue) and make the product stable.

> Steps
>
> 1. Analyze/Be familiar with the code
> 2. Reproduce the errors in the lab
> 3. Solve

In the first impression of the code, she saw many memory allocations (<cpp>new</cpp> and <cpp>delete</cpp> statements), which could lead to memory leaks and memory fragmentation. in theory, this could be the reason. But eventually it turned out not to be the case this time. the memory leaks and fragmentation didn't cause the resets.\
To tests the assumption, they developed a load simulator to run in the lab, with some workarounds to make the load get there faster. they combined it with a memory profiler to understand the memory allocations (size and quantities, typical allocation, frequency) and even put a breakpoint when a memory allocation fails and get a memory dump when the system is about to fail.\
The unit crashed on race condition, two threads trying to create large messages. this wasn't about the memory size.

the fix itself involved several changes:

> 1. Changed from asynchronous to synchronous work.
> 2. Reduce message size and shorten time by replacing Qt with C++. (char in qt is two bytes, in C++ it's one byte).
> 3. Split large messages into smaller ones.
> 4. Use a static array to prepare the message instead of dynamic allocation.

This caused a problem, data waiting for transmission remained in the RAM until being sent. this takes up a lot of space, and when there's a reset or communication lost, we end up losing data again. a solution was to separate the business logic from the networking logic, and keep a queue in a persistent database in the SD card. this allows for holding memory even when there are resets, and removes the sensitivity to network errors.

the take away is to measure everything, test the assumptions, check if your theories are correct. have a simulator (injecting events), the earlier the better, don't wait for the problems to come from the field. another suggestion is to separate the business logic from the networking, keep a queue for input and output messages. an application that is simple is harder to write, but easier to understand. multiple threads are a source of problems.

</details>

## Amir Kirsh :: C++ Incidental Explorations

<details>
<summary>
Learning C++ from questions and answers on StackOverflow.
</summary>

[C++ Incidental Explorations](https://youtu.be/Vibfa1GqP18?si=QtAEFCUa3Jo6tSZI)

Learning from stackOverflow, either by reading question, asking them or trying to answer. it teaches us how to convey our knowledge, and other people can comment and request fixes to the answer, like a code review session.

the questions act as a knowledge database which we can return to in a later time or send to other people.

- does <cpp>auto</cpp> increase performance?
- how to implement a custom iterator?

some questions also teach us about nuances of the language, what is possible, what is legal, and what works. even bad questions can teach us something, since they come from a misunderstanding.

- why can you access discarded local variables through the address?
- can an object hold a member data of it self?

we can also run queries on SO and see questions by metrics, like the most answered questions and most down-voted, and which tags are most popular.

other stuff

- why isn't <cpp>std::unique_ptr::release</cpp> marked with <cpp>[[nodiscard]]</cpp>?

questions that the answer changes for over the year.

- best practices
- changes of the standards
- old questions that didn't have an answer back when they wre created, but now do

_can Chat-GPT do better than a human as a Q&A site?_

</details>

## Roi Barkan :: More Ranges Please

<details>
<summary>
Discussing Ranges, Views and algorithms.
</summary>

[More Ranges Please](https://youtu.be/rfPq6hYCvWI?si=yXwJwp2cLfxc_fvq), [slides](https://docs.google.com/presentation/d/115iU6rA6gAUNErenmBIPqdl_HHlcYp9_YoYLOFSUNho/present?slide=id.p)

### Libraries

starting with the defintion of libraries. what they do and why they exist. layers of abstraction, the reason we should write our own. what can future libraries give us? composing and stacking library code together. the steps that we go through when writing library (any of our code can become one).

- API layer
  - easy to use, hard to misuse
  - specifying pre and post conditions
- data types - regular
  - value semantics
  - algebraic - do they behave according to mathmatical rules (groups, monads, monoid)?

### Ranges

we got ranges as one of C++20 big four features (ranges, coroutines, concepts, modules). they are extension of the iterator based algorithms, and draw inspiration from functional languages. they also focus on composability, the algorithms take ranges and return ranges back. and we can use range adaptors (algorithms as objects), views (lazy initialization) and projections (unary transformations).

```cpp
// chaining algorithms
std::ranges::reverse(std::ranges::search(str, "abc"sv));
// views as composable lazy ranges
str |
   std::views::split(' ') |
   std::views::take(2);
// views have a value/algorithm duality
auto square_evens =
   std::views::filter([] (auto x) { return (int(x) % 2 == 0;)}) |
   std::views::transform([] (auto x) { return x * x; });
```

example of assigning seats in parliament based on proportional voting. using ranges and views to make the calculation simpler and faster.

### The Future

suggestions for new views and ranges, we look at what we use in our codebase and that's how we discover what kinds of algorithms we need. even if we don't create new views, we can still extract common behaviors into composable objects.

```cpp
auto histogram = std::ranges::views::chunk_by(std::equals{}) |
   std::ranges::views::transform([] (const auto& rng) {return std::make_pair(std::begin(rng), std::size(rng));});
```

we can potentially break the sorting algorithm apart, inspired by the R standard library. a sorting algorithm makes a number of comparisons (O(NlogN) complexity) and swaps elements with <cpp>iter_swap</cpp>. if the cost of swapping elements is high, maybe we can avoid that action. like getting the index order without performing the swap (will require allocating memory) and then we just need to perform the final swaps (N operations), and then we can use the same order to reverse the ordering. also applies to permutations of ranges.

</details>

## Anastasiia Kazakova :: Standard C++ toolset

<details>
<summary>
C++ doesn't have a complete toolset, but it has the tools.
</summary>

[Standard C++ toolset](https://youtu.be/9WYhDD3YY-k?si=ZynJoQLJdvW3YfF4)

C++ doesn't have a standard toolset, and we can look at the frustration point from developer surveys and see that many of them involve tool (managing libraries, building projects, creating a new project). we have essential tools such as the compiler, build-system, libraries and dependency management and debugging tools. we also have complementary tools, which are nice to have, but not required, analyzers, profilers, unit-testing framework and code-coverage reports.

other languages have a standardized (or nearly standard) toolset.

having a toolset has benefits:

1. help new developers
2. onboarding via dev environments
3. code unification
4. easy to adopt best practices
5. code sharing via libraries

In c++ we have four big compilers

- GCC
- Clang
- MSVC
- Intel

They generate different assemblies, have different error reporting, each with it's own predefined macros.

```sh
gcc -E -dM -O3 c++ /dev/null # gcc, also clang
cl /EP /Zc:preprocessor /PD empty.cpp 2>nul # msvc
```

different compilation flags.

```cmake
if (MSVC)
   add_compile_options(/W4 /WX)
else()
   add_compile_options(-Wall -Wextra -Werror)
endif()
```

Clang by itself isn't just the compiler, it as analysis, formatting and tidy tools, AST (abstract syntax tree) parsing tool. so it becomes the basis for other tools.

CMake is the most popular build system, with msBuild and makefile behind. the adoption starts with the IDEs, package managers and libraries. there are even cmake debuggers and profilers in the CLION IDE. there was a big effort to have support for c++ modules. Dependency management is still far wide spread adoption, even with some tools gaining popularity (conan, vcpkg).

- formatting tools - clangFormat
- code analysis - most developers don't use at all
- data flow analysis
- domain specific analyzers
- doing work in the ci pipeline
- unit tests - google tests, catch(no longer header only)
- test code coverage - llvm-com (`-fprofile-instr-generate -fcoverage-mapping`), gcov (`-profile-arcs -ftest-coverage`).

</details>

## Dor Rozen, Igor Pora-Leonovich :: Performance-related coding guidelines

<details>
<summary>
benchmarking performance for containers and guidelines to improve it.
</summary>

[Performance-related coding guidelines](https://youtu.be/PGd7vG1CMD0?si=r2vDVmvVTqDjo003)

This talk will compare data structures, go into the implementation, and look at the differences between theoretical and practical considerations.

Starting with benchmarking <cpp>std::vector</cpp>, <cpp>std::list</cpp>, <cpp>std::deque</cpp>. going over the usual operations, and seeing that the main contributor to high speed is spacial locality (line cacheing), and that increasing the element size makes pointer based data structures perform better.

Move semantics in containers, benchmarking performance. how move semantics affect the containers common actions. optimizing by shifting calculations to compile-time.

</details>

## Alex Dathskovsky :: To Int or to Uint, This is the Question

<details>
<summary>
Looking into signed and unsigned types, how they differ and what should we do with them.
</summary>

[To Int or to Uint, This is the Question](https://youtu.be/Iz2UOgLMj58?si=T0UwTLCir9QUqsqw)

is there a difference between these two pieces of code?

```cpp
int64_t add_and_divide_s(int64_t a, int64_t b) {
   return (a+b)/2;
}

uint64_t add_and_divide_u(uint64_t a, uint64_t b) {
   return (a+b)/2;
}
```

turns out there is, if we look at the assembly code: the signed code is much longer and more complex. they both use shifts and now division instructions (which is cheaper). but the signed version has additional instructions.

```mips
# add_and_divide_s(long, long)
lea rcx, [rdi + rsi]
mov rac, rcx
shr rax, 63
add rax, rcx
sar rax # arithmetical shift right
ret

# add_and_divide_u(unsigned long, unsigned long)
lea rax, [rdi + rsi]
shr rax # logical shift right
ret
```

unsigned integers are represented in a very simple way, with a consistent overflow defintion. in contrast, signed integers can be stored in different ways (sign and magnitude, Ones complement, Twos complement), and overflow is undefined behavior.\
To represent a negative number with Ones complement, we simply invert the bits. for negative numbers with Twos complement, we start with the positive number, invert the bits and add 1 (ignoring overflow)

| Bits | Unsigned Value | Ones' Complement | Twos' Complement |
| ---- | -------------- | ---------------- | ---------------- |
| 000  | 0              | 0                | 0                |
| 001  | 1              | 1                | 1                |
| 010  | 2              | 2                | 2                |
| 011  | 3              | 3                | 3                |
| 100  | 4              | -3               | -4               |
| 101  | 5              | -2               | -3               |
| 110  | 6              | -1               | -2               |
| 111  | 7              | -0               | -1               |

Since C++20, negative numbers must be represented with Twos Complement.

there are two kinds of Shift instructions: logical and arithmetical

- `SHR` (logical shift right) - shifting the bits to the right and the most significant bit (MSB) becomes zero. `shr 10110111` -> `01011011`.
- `SAR` (arithmetical shift right) - maintains the MSB as it was in the original number. `sar 10110111` -> `11011011`.

but because of some rounding mode, the compiler must do some work with the signed numbers so it could use shifts and not division.

the difference between the performance isn't great, but unsigned numbers are faster.

### Pitfalls

```cpp
auto add_uint8(uint8_t a, uint8_t b)
{
   return a+b;
}

// compiler sees
int add_uint8(uint8_t a, uint8_t b)
{
   return static_cast<int>(a)+static_cast<int>(b);
}
```

if we call the above function `add_uint8(255u, 1u)`, what would we get?
we actually get integer promotions. the compiler prefers to do arithmetic operations only on types which are at least 32-bit.\
And because we used <cpp>auto</cpp> as the return type, the result is an int with the value 256. had we used <cpp>uint8_6</cpp> as the return type, we would get overflow, and the result would have been narrowed to zero.

```cpp
uint8_t add_uint8(uint8_t a, uint8_t b)
{
   return a+b;
}

// compiler sees
uint8_t add_uint8(uint8_t a, uint8_t b)
{
   return static_cast<unsigned_char>(static_cast<int>(a)+static_cast<int>(b));
}
```

another possible surprise, what if everything was automatically deduced?

```cpp
auto my_add(auto a, auto b)
{
   return a+b;
}
```

if we call `my_add(uint64_t(1), int64_t(-2))`, what would be the result? the actually does an integer promotion, turning the signed number to unsigned, and the result is some monster number.

```cpp
// compiler sees
unsigned long my_add(unsigned long a, long b)
{
   return a + static_cast<unsigned long>(b);
}
```

another example of mixing can happen with loop counters. in this case, we get infinite loop.

```cpp
uint64_t count(uint64_t size)
{
   uint64_t count;
   for (int i = 0; size - i >= 0; i++) {
      count++;
   }
   return count;
}
```

a different example, a promotion causes buffer overflow.

```cpp
void decode(std::byte* bytes, int size)
{
   if (size == 0) return;
   std::byte decoded[255];
   for (uint64_t i = 0; i < size, i++) {
      decoded[i] = static_cast<std::byte>(static_cast<uint8_t>(bytes[i])^0xc);
   }
}
```

this is a pattern to be aware of, using <cpp>auto</cpp> is good, but can lead to problems. using it with a number causes it to become a concrete type.

```cpp
auto a1 = 0; // int
auto a2 = 0u; // unsigned int
auto a3 = 0l; // long
auto a4 = 0ul; // unsigned long
auto a5 = 0ll; // long long
auto a6 = 0ull; // unsigned long long

// C++23 additions
auto a7 = 0z; // signed size_t
auto a8 = 0uz; // unsigned size_t
```

there are special integer types:

> - <cpp>size_t</cpp> - unsigned integer
>   - used for size operations
>   - defined in cstddef
>   - limit is `SIZE_MAX`
>   - introduced in C89 to eliminate portability problems.
> - <cpp>ssize_t</cpp> - signed integer
>   - defined by POSIX.1-2017
>   - represent at least the range of [-1, {SSIZE_MAX}]
>   - limit is `SSIZE_MAX`

the option for -1 size is good for vectors.

this program tries to shift a signed integer 1, but it can overflow.

```cpp
uint64_t do_it(uint64_t count) {
   return 1 << (count % 64);
}
```

### Arithmetic Series

> Series of numbers where the difference between any two sequential number is constant.\
> For example, _1,2,3,4,5,6,7,8,9,10,...,n_ is a an Arithmetic Series where the difference between any two sequential numbers is _1_.

we know the quick formula to get the result without doing the entire computation.

$$
\sum\limits_{k=1}^{n}a_k = \frac{n(a_1 + a_n)}{2}
$$

we can write the same code with signed and unsigned numbers,

```cpp
uint64_t arc_unsigned(uint64_t n)
{
   uint64_t sum = 0;
   for (uint64_t i = 0; i < n; i++) {
      sum += i;
   }
   return sum;
}

int64_t arc_signed(int64_t n)
{
   int64_t sum = 0;
   for (int64_t i = 0; i < n; i++) {
      sum += i;
   }
   return sum;
}
```

if we look at the assembly, the compiler creates the loop for the unsigned number, but for the singed version it uses the formula and makes things faster. the performance difference is noticeable. the reason is that overflow is defined behavior in unsigned numbers, so the compilers must give the 'correct response' of doing the loop calculation. with signed numbers, overflow is undefined, so the compiler can do what it wants and optimize away the loop.\
Newer compilers optimize better, and we can use sanitizers to force the same behavior `-fsanitize=signed-integer-overflow` and `-fsanitize=unsigned-integer-overflow`.\
We can also use special types that are defined for better performance <cpp>int_fastN_t</cpp> or <cpp>uint_fastN_t</cpp> to ask to compiler to choose the correct type for us. \
There are also helpers from the standard: <cpp>std::make_signed_t</cpp> and <cpp>std::make_unsigned_t</cpp> which can be used to created signed and unsigned number safely.\
since C++20 we have safe comparisons for signed and unsigned numbers:

- <cpp>std::cmp_equal</cpp> `==`
- <cpp>std::cmp_not_equal</cpp> `!=`
- <cpp>std::cmp_less</cpp> `<`
- <cpp>std::cmp_less_equal</cpp> `<=`
- <cpp>std::cmp_greater</cpp> `>`
- <cpp>std::cmp_greater_equal</cpp> `>=`

in general, we should avoid using <cpp>auto</cpp> when we aren't sure about the correct integer type, and prefer using concrete types when possible. a better alternative is using explicit strong types. modern loops are also a good way to avoid bugs.

</details>

## Sebastian Theophil :: A Practical and Interactive Guide to Debugging C++ Code

<details>
<summary>
Ways to debug code.
</summary>

[A Practical and Interactive Guide to Debugging C++ Code](https://youtu.be/ogV0olQJax4?si=u7fVLO888bstrVm0)

A bug is behavior in the software that doesn't conform to the specification. we usually see a "bug report", a symptom of the bug, something that happened, an event that shouldn't have. it can be a crash, core dump, error message, log line. we don't always see them, especially if they happen on the client's machine.

we have a few ways of detecting bugs before shipping. some of them even at compile time.

- type checking
- <cpp>static_assert</cpp>
- <cpp>constexpr</cpp> evaluation

at build time and QA we use unit testing and other automated testing systems. but that still leaves undiscovered bugs. even at runtime, we should:

> - Strict Error Checking
>   - Check all API return values and report unexpected values
>   - Assert pre-conditions and post-conditions
>   - Report if they fail
> - Enforce invariants, notice unexpected behavior sooner

we need the bugs from the customers, even if we can't reproduce them in the lab. sometimes the fix is obvious, sometimes we can just try something to get better error reporting and understand the problem better if it happens again.

here is an example of a bug.

```cpp
enum EState { WAITING, DATA, ERROR };
struct http_delegate {
   EState m_estate = WAITING;
   std::mutex m_mtx;
   std::condition_variable m_cv;

   void on_new_data()
   {
      {
         std::lock_guard lock(m_mtx);
         m_estate = is_error() ? ERROR : DATA;
         // copy data to buffer
      }
      m_cv.notify_all();
   }
};

struct sync_http_request {
   http_delegate m_delegate;

   sync_http_request(/*... */)
   {
      // set up everything

      // wait for response
      std::unique_lock lock{m_delegate.m_mtx};
      m_delegate.m_cv.wait(lock, [&](){
         return m_delegate.m_estate != WAITING;
      });

      if (m_delegate.m_estate == ERROR) {
         throw http_exception();
      }

      // do something with the data
   }
};
```

the reason for the bug is spurious wake-ups, there is a tiny chance that the conditional variable will wake up between the time the state was changed, but before it was supposed to wake up by the call `m_cv.notify_all()`. and if that one in a million thing happened, the delegate might be destroyed while the member is being accessed. for this reason, the standard says the action on the conditional variable should also be inside the lock.

### Debugging Process

An iterative process, it's important to do right, not fast. we form an hypothesis, test it, repeat, and eventually fix it.

the first step is to be able to reproduce it. there are different levels of "reproducibility" - the best case is that it always occurs, on any machine, with the interactive debugger attached. the worst case are bugs which sometime occur, only on release build, only on specific machines. part of debugging is moving up the chain, even if we don't know the underlying issue.\
we can use tools to make bugs appear:

- address sanitizer
- thread sanitizer
- undefined behavior sanitizer

interactive debuggers make some bugs disappear if they are based on timing issues, and won't appear if the program is too slow. we want to be able to force the bug to appear and diagnose the system when the problem happens.\
if the issue happens only on a specific machine, it could be the result of other software interfering, such as virus scanners, system tools, drm software. we want to be able to reproduce the environments as a VM.\
If everything fails, there are still things we can do, such as writing more logs, or more detailed reporting, and even trying out solutions in production, if possible.

an example of another bug, this time with <cpp>std::bad_alloc</cpp> error. it started with a bad configuration file. part of the bug was the difference between the value of <cpp>std::size_t</cpp> in 32 and 64bit machines.

we can use logging statements to debug the issue, but a better way is to take advantage of **tracing breakpoints** (available in gdb, lldb, visual studio), which don't require re-compilation of the source code, can be added to OS functions and binary code. if we suspect the reason is interaction between the software and the OS, we can use specific tools

- Process Monitor for windows
- dtrace for macOS
- strace for linux

if we know a bug is new, we can look at the changes that introduced the bug, and find a solution for both the bug and the issue the code was meant to address.

```cpp
std::array<int, 4> an ={1,5,7,8};
auto rng = GetItemFromIndices(&an[0],4); // external code
assert(rng.size() == 4); // fails with rng.size == 2?
```

we can set a watch point `watch set expression -s 4 -- &an[0]` and see where the input variable is touched.

before we fix the bug, we take a moment and think about it. some bugs are because we didn't write what we meant - we made a mistake such as using uninitialized data, using memory that was freed (other memory management problems). sometimes we wrote what we meant, but our mental model was wrong. we didn't understand the specification (ours, or external tools), and we need to think about our code base again.

when we fix the bug, we can have two approaches - one is "smallest fix possible", but that might not fix the root cause. the other option is to continue working on the root cause of the bug. we can combine the two - small fix now to production, and a more through fix at the development branch.

the next step is reviewing how the bug happened in the first place, and what can we do to avoid it in the future. we also need to document that changes.

</details>

## Rainer Grimm :: Concurrency Improvements in C++20: A Deep Dive

<details>
<summary>
Discussion of Concurrency mechanisms and the changes in C++20.
</summary>

[Concurrency Improvements in C++20: A Deep Dive](https://youtu.be/J-s5jNq3VA0?si=9zK2cgxwgy_wmx3D)

Things we got in C++20.

- Atomics
- Semaphores
- Latches and Barriers
- Cooperative interruption
- Joinable Threads

### Atomics

> "Atomic operation on atomic define the synchronization and ordering constraints"

we had atomics since C++11, they are the foundation for the C++ memory model, and are used by higher-level threading interfaces such as:

- threads and tasks
- mutexes and locks
- condition variables

C++ has the <cpp>std::atomic_flag</cpp> inteface, which is guaranteed to be lock-free. the interface is:

- `test_and_set` - set the value and return previous one
- `clear` - clear the value (make false)

can't be read without modification.

C++20 added support for floating point and smart pointers to also be atomic with <cpp>std::atomic<></cpp>, which has a more complex interface

- `is_lock_free` - check if lock free
- `load` - return the value of the atomic
- `store` - set the value of the atomic with non-atomic
- `exchange` - replace with new value and return the old
- `compare_exchange_weak`, `compare_exchange_strong` - check condition, return boolean and either replace the atomic value or the other one (input value).
- `fetch_add`, `fetch_sub` - adds the vales and return the pervious value
- `++`, `--` - increment or decrement and return the new value

the "compare-exchange" member functions are the core functionalities, and can be used to create the other operations (multiplication, division). we check the existing value and only do the operation if the state didn't changed.

[compiler explorer](https://godbolt.org/z/YbMarGvoh)

```cpp
template <std::integral T>
T fetch_mult(std::atomic<T>& shared, T mult){
  T oldValue = shared.load();
  while (!shared.compare_exchange_strong(oldValue, oldValue * mult));
  return oldValue;
}
```

there are also new operations in C++20.

- `notify_one()`
- `notify_all()`
- `wait(val)` - wait for notification, blocking if the value of the atomic is `val`.

C++11 gave us <cpp>std::shared_ptr</cpp>, the handling of the control block is thread-safe, but the access to the resource isn't. there are new specializations of <cpp>std::atomic</cpp> for <cpp>std::shared_ptr</cpp> and <cpp>std::weak_ptr</cpp>. there are currently no lock-free implementations for them.\
The benefits are for consistency, correctness, and speed.

another option is <cpp>std::atomic_ref</cpp> to apply atomic operations to referenced objects, (the lifetime of the object must exceed the lifetime of the reference).

(we should always use `-fsanitize=thread` to when dealing with concurrency).

### Semaphores

> Semaphores are synchronization mechanisms to control access to a shared variable.\
> A semaphore is initialzed with a counter (limit) greater than zero,
>
> 1. requesting a semaphore increases the counter
> 2. releasing a semaphore decreases the counter
> 3. a requesting thread is blocked if the counter is 0

C++20 support two semaphores:

- <cpp>std::counting_semaphore</cpp>
- <cpp>std::binary_semaphore</cpp> (same as counting semaphore initialzed with 1)

we have operations to get the limit, to release (post), to acquire (get, take, wait), and variation on acquire (for time duration or until an absolute time). the semaphore is used for producer-consumer workflows.

There is also <cpp>std::conditional_variables</cpp>, for sender-receiver workflows. it can be activated by spurious wakeups,and there are also "lost wakeups", so we need a predicate protection in both cases.

we can see a comparison table between the options, the slowest is the conditional variables, the other options (semaphores, atomic boolean, atomic flag) are faster.

### Latches and Barriers

<cpp>std::latch</cpp> - a thread waits at a synchronization until a counter reaches zero. useful for managing one task by multiple threads.

- `count_down` - don't block, atomically reduce the counter
- `wait` - wait (block) for counter to be zero
- `try_wait` - check if counter is zero, but don't block
- `arrive_and_wait` - reduce the counter and wait.

can only be used once, if we want to reuse it, we need a <cpp>std::barrier</cpp>.

- `arrive` - decrements counter
- `wait`
- `arrive_and_wait` - decrement counter and wait
- `arrive_and_drop` - decrement counter for the current and the next phase

### Cooperative interruption

each running entity can be interrupted, using <cpp>std::stop_token</cpp> which is passed to either a joinable thread or a conditional_variable_any. this doesn't force the thread to stop, it tells it that someone asked it to stop.

- `token.stop_possible` - check if the token has a associated stop state
- `token.stop_requested` - check if `stop_source` was called
- `source.get_token` - create token if possible,
- `source.stop_possible` - check if the source has an associated stop state
- `source.stop_requested` - check if a stop was requested
- `source.request_stop` - request a stop

there is also a <cpp>std::stop_callback</cpp> which we can register on the token.

### Joinable Threads

<cpp>std::jthread</cpp> automatically joins in the destructor. so we can't have dangling threads which calls terminate when it's destroyed and aborts the entire program. it's an RAII mechanism.

</details>

## Nevin Liber :: MDSPAN: A Deep Dive Spanning C++, Kokkos & SYCL

<details>
<summary>
How MSpan got created and accepted to the standard
</summary>

[MDSPAN: A Deep Dive Spanning C++, Kokkos & SYCL](https://youtu.be/AqibvB5Kl8I?si=KV-q7mVyAfTTmwML)

> MDspan is a non-owning multidimensional array view for C++23

```cpp
template<
   class ElementType,
   class Extents,
   class LayoutPolicy = layout_right,
   class AccessorPolicy = default_accessor<ElementType>
   >
struct mdspan {
   template<class... OtherIndexTypes>
   explicit constexpr mdspan(data_handle_type p, OtherIndexTypes... ext);
   // ...
   template<class... OtherIndexTypes>
   constexpr reference operator[](OtherIndexTypes... indices) const;
};
```

- Element type - self explanitory, what kind of elements is stored
- Extents - dimensions of the multidimensional array, static or dynamic extents
- LayoutPolicy - row major, column major, stride, or user defined (tiled, symmetric, sparse, compressed, etc...)
- AccessorPolicy - way to access the data, if we want to customize it to return something special.

### The road to get here

2014 - proposal for "array_view", static extents only, got positive reviews and pushed into Arrays technical Specification. but there was also a suggestion for runtime-sized arrays, the whole suggestion died out in 2016.\
array view got another revision which was sent to another committee as part of the library fundamentals discussions, which also got some issues raised. multiple revision cycles over the years.

Kokkos had a similar thing, <cpp>Kokkos:View</cpp>, multidimensional view for zero or more dimensions, which sometimes owned the data and sometimes didn't and had some weird syntax to define dynamic and static dimensions.

```cpp
template <
   class DataType,
   [, class LayoutType]
   [, class MemorySpace]
   [, class MemoryTraits]
   >
class View;
```

other suggestion were for "shared_array" and "weak_array", fairly complicated stuff. eventually the single dimension `span` got rolling. more talks about multidimensional arrays, some more issues to work through. in 2017 the name was changed to "mdspan", and the proposal was moved to library fundamentals v3 comity. some changes with the comma operator and the brackets operator, the class template deduction guide was also added and made things easier for classes with many template arguments.

eventually it was accepted into C++23.
hopefully c++26 will have the "mdarray" - an owning version of the multidimensional data. also a proposal for sub-mdspan and atomic references and atomic accessors.

</details>

## Daniel Babitsky :: Building low latency, network intense applications with C++

<details>
<summary>
Short Hebrew talk about low latency applications.
</summary>

[Daniel Babitsky :: Building low latency, network intense applications with C++](https://youtu.be/XLoL5vJLvV4?si=ca0ZpU74JHF1ClJy)

Algo-trading companies need low latency and high performance. they want algorithms to run on live data and be as fast as possible.

options to use for low latency:

- Kafka
- RabbitMQ
- Redis

but they aren't fast enough, they need a faster way to send messages. that's why the company went with the publisher/subscriber messages.

other stuff - using templates to avoid runtime branches. minimizing memory allocations by using managed memory pools. using lock-free data structures. using kernel bypasses with "openOnload" library (reading from userSpace buffer without copying). using EFVI Api to access the network adapter data path faster than the normal pathway.

</details>

## Marshall Clow :: Development Strategies - The Stuff Around The Code

<details>
<summary>
Things That we do which aren't just the application code.
</summary>

[Development Strategies - The Stuff Around The Code](https://youtu.be/qNdp4C1AS8c?si=ovKZS6yepXzVRqdU)

the stuff that we do that isn't code:

- Documentation
- Tests
- Tools
- Dealing with Users
- Releases
- Managing Change

### Documentation

Documentation is the first impression, even a readme file is a documentation. this is how many users evaluate the library, if there aren't any, or if it's bad, people will move on to the next one. and it also helps prevent answering the same question again and again.

parts of the documentation that we might have.

1. Overview - what does the application do.
2. Getting Started - getting the library, installing it, building, linking.
3. Examples - Using the library at the most simple way.
4. Tutorials - How to do more than just the basic.
5. Reference - What do the calls do, what are the parameters for them.

other things can be a roadmap for the library (future stuff), known bugs, how to contribute if it's an open source, etc...\
We can use tools like Doxygen to generate documentation from the code, reducing the risk of it going out of date.

### Tests

> Having a good set of test can help with:
>
> 1. Making sure you've fixed a bug
> 2. Catch Regressions
> 3. Pinpoint problems with someones installation or configuration
> 4. Porting to a new platform

large projects can go into the realm of "fear-based programming", where changes are dangerous because they might break something, a good test suite (which we trust) can allow us to be more confident about what we write. we add tests to validate the things the user except work as excepted.

> a good test suite is:
>
> 1. easy to run
> 2. always green
> 3. comprehensive
> 4. runs often - preferably on every change
> 5. easy to automate

same principles with compiler warning, if there are always hundreds of warning, then no one will notice a new one. we want the metric to be at zero problems, so if there's a new issue (test fail, warning), it stands out immediately.

There should be documentation for running the test suite, even on different platforms and machines. this especially helps detect hidden assumptions that we made.

performance tests are harder, we can instrument the tests to check for number of operations, but it's much harder to test for timing, as it is dependant on many things.

### Tools

we have tools that can help us:

1. Compilers - first line of defense, should have all flags and warning on.
2. Static Analyzers - runs on source code, detects cases which are likely to be errors or will find bad code. we can extend it with custom checks that are specific to our codebase.
3. Dynamic Analyzers - perform their work on running programs, run assertions and debuggers, sanitizers (address, thread, undefined-behavior, etc...) to detect issues - they are very good at finding bugs, and have a very low false-positive rate.
4. Code Coverage Tools - detect which code is run, which sections aren't tested.
5. Fuzzers - creating random input and checking how the program handles it - if it misbehaves or crashes. they can combine with code coverage tools to find corner cases. very important when we have parsers.

### Dealing with Users

> Users can:
>
> 1. ask questions
> 2. make comments
> 3. file bug reports
> 4. make feature requests
> 5. offer contributions
> 6. port to new systems

we need to respond to them and address what the say. answer questions (add to documentation), check if things are bug or not, decide if a feature is really needed or not.

users use the library in ways we didn't expect.

### Releases

do we have release versions, milestones, etc...?

they should be announced, detailed with notes (what changed, what was fixed, what was added) and a way to obtain the specific release.

### Managing Change

When the usage grows, we might run across the need to have breaking changes, this needs to be handled carefully. the benefits are clear, but the costs are obscured since they are experienced by the users and not by you. there needs to be a clear path to upgrade that users can follow and make sure things will keep working.

</details>

## Vittorio Romeo :: Improving Compilation Times: Tools & Techniques

<details>
<summary>
understanding what effects build time and what actions can improve them.
</summary>

[Improving Compilation Times: Tools & Techniques](https://youtu.be/wCajYrd1PIk?si=V_vjWWM0hTCRZ-Uw), [github repo](https://github.com/vittorioromeo/corecpp2023)

SFML - Simple and Fast Multimedia Library

### Why compilation times are important

C++ has reputation for being slow to compile (compared to C), we talk about "zero-cost abstractions", but we mean that they don't have cost at runtime, as we usually pay it in compile time. the costs are often shifted to the CI pipeline, and even in modern hardware, building from scratch can a long time.

However, for projects that never considered build times, there is a lot of room to improve.

> Why can C++ compilation times be poor?
>
> - Build model and textual `#include` system is archaic
> - The language itself is complicated
>   Overload resolution, template instantiation, SFINAE, etc...
> - Highly generic and abstracted libraries tend to be bulky
>   - Think about Boost or the Standard Library
>   - Many reasons: backwards compatibility, build time not a priority, etc...
> - Poor "physical design"
>   - E.g., `#include` when forward declaration is enough
>   - E.g., templates unnecessarily defined in a header file
> - Compilation times are often not a priority for low-level libraries
>   - This includes the Standard Library
>   - Needs a "cultural" change

C++20 modules will help, but we are far from that point. the compiler support is limited, and many projects haven't migrated yet.

### Low hanging fruit

the first thing we ask is if our compilation time is fast enough - in terms of costs, developer time, frustration, reputation or other stuff. defining "fast enough" is subjective.

the low hanging fruits are small changes, which are fairly easy to introduce and effect the system as a whole.

### Build System

the first option is to switch the build system, moving from _make_ to alternatives such as _ninja_ can speed up the process and provide more advantages. we can use them as drop-in replacements for cmake.

```sh
cmake -GNinja
ninja
```

#### Linker

another option is to replace the default linker `ld` with something else, even switching to llvm linker `lld` can give an order of magnitude change in link time.

```sh
cmake -GNinja -DCMAKE_CXX_FLAGS="-fuse-ld=lld"
```

there are even faster options, such as `mold` - "modern open source linker", which focus on high parallelization. the commercial version is called `sold` which also works for macOs and eventually Windows.

#### Compiler Cache

another easy way is to avoid re-compiling unchanged source files, we can have a compilation cache that maps compiler, flags and the file hash to an object file. so even with fresh build, we can re-use existing ".obj" files. there are all sorts of tools (open source and commercial) to assist with this.

#### Precompiled headers

Precompiled headers allow us to avoid re-processing the same header file again and again each time it is used. we can define some headers as commonly used and prepare them ahead of time. there are some pitfalls, but the gains can be massive if we choose the headers correctly.

```cmake
target_precompile_header(<target> PRIVAE "PCH.hpp")
target_precompile_header(<base_target> PRIVAE "PCH.hpp")
target_precompile_header(<other_target> REUSE_FROM <base_target>)
```

the fle itself will look something like this:

```cpp
// PCH.hpp
#pragma once
// Commonly-used first-party headers (e.g., logging, assertions, basic components)
#include <SFML/System/Err.hpp>
#include <SFML/System/String.hpp>
#include <SFML/System/Time.hpp>
#include <SFML/System/Vector2.hpp>
// Expensive headers, like `windows.h`
#ifdef SFML_SYSTEM_WINDOWS
#include <SFML/System/Win32/WindowsHeader.hpp>
#endif
// Commonly used Standard Library or third-party headers
#include <algorithm>
#include <filesystem>
#include <iostream>
#include <memory>
#include <string>
#include <unordered_map>
#include <vector>
/* ... */
```

precompiled headers can cause problems if they grow too large, are unnecessary included, or change too often. there are ways to find which headers are mostly used or which are expensive to compile.

#### Unity Builds

coalesce mutiple source files together, reduce the number of times the same work is done, less files are created, which acts as a substitute for link time optimizations. it also helps with ODR violations and detecting other possible mistakes.

```sh
cmake -GNinja -DCMAKE_UNITY_BUILD=ON -DCMAKE_CXX_FLAGS="-fuse-ld=lld"
```

or

```cmake
set_target_properties(<target> PROPERTIES UNITY_BUIILD ON)
```

there are drawbacks, and many places which can fail, but it might be worth fixing the code to gain the speedup.

### Profiling and dealing with bottlenecks

There are other ways to improve compilation times, but they are harder. the first step is to profile the compilation and find the bottlenecks. we can add the `-ftime-trace` flag to the clang compiler and then run `./ClangBuildAnalyzer` on the resulting files and see which files are the heaviest and take longest to compile.

#### Physical Design

when we find the bottlenecks in our first-party code, we can start looking into it. simplifying dependency chains, separating the declaration and defintions, using forward declarations when possible (better for internal libraries, not externally facing code). isolating implementation from the declaration (PIMPL idiom) if it's expensive.

we can check the impact of each header from the standard library on the compilation time, so if we include the entire header for just one function, we might be better off writing the code directly. every edition of the standard becomes larger and larger, which also effects the time. if we know what we are doing, we can look at ways to work around it.

</details>

## Daisy Hollman :: Expressing Implementation Sameness and Similarity in Modern C++

<details>
<summary>
Different ways to express things are the same.
</summary>

[Expressing Implementation Sameness and Similarity in Modern C++](https://youtu.be/wp9cf0u2iss?si=iphVXWZ_V4s5NiiC), [slides](https://daisyh.dev/talks/corecpp-2023/expressing-implementation-sameness.html)

"sameness" isn't just polymorphism, and it isn't just solved by code re-use.

> - code is written and read by humans
>   - we can ask the computer to do the same things in many ways
>   - we should think about code in terms of "how will the reader experience it"
> - cognitive load is the amount of mental effort required to and understand and use code
>   - when we write code, we want to minimize the amount of cognitive load for the reader/user
> - information loss: when the reader of the code doesn't have an easy way to retrieve a piece of information that is obvious to the writer
>   - copy-pasting code causes information loss
>   - imagine how our code would be if we weren't able to copy-paste
>   - copy-paste code is selfish - it saves time for to the writer at the expense of the reader.
>
> things that don't matter
>
> - how long it takes to type
>   - if the function name is long, it's ok. the problem ins't the number of characters, it's the cognitive load that matters.
> - how many layers of abstraction there are
> - compilation costs are ok to add, it is cheaper than development time.
>   - but not exponentially scaling compilation costs
>   - avoid using arcane (weird language features)
> - how fast it is to write code.
>   - this talk is about reading code.

Two kinds of "sameness": Interface Sameness and Implementation Sameness

> Interface Sameness:
>
> - Allows user to treat things the same way in their code (create their own sameness)
> - Hard to remove later, you can't change user code, once you allow users to treat things the same way, it's very hard to change that later
> - When done well: low code coupling (low interferences between software modules)
>
> Implementation Sameness:
>
> - Enables readers to understand and use existing sameness or similarity
> - Can be changed or removed at any point in the future when it stop being helpful.
> - When done well: high code cohesion (high degree of elements belonging together in the same module)

This goes hand in hand with the DRY principle (don't-repeat-yourself)

> _"Every piece of knowledge must have a single, unambiguous, authoritative representation within a system"_
>
> ~ Andy Hunt and Dave Thomas, The Pragmatic Programmer

### Expressing "Sameness" in C++

traditional polymorphism, inheritance, virtual function, using base class pointers, this is interface sameness (with some shared implementation).

we also have templates, which are mostly implementation sameness, but also have interface sameness. we can also use templates to add "interfaces" specific intrusiveness. we sometimes create sameness that we have a problem abstracting away. we have an example with class that contains things and expose more functionalities. we handle it with "mixins" (using inheritance), other languages have other ways to handle this.\
Other mixins are comparison operators, when we write operator overloads for `==` and `!=`, we usually do the same thing across many classes - **elementwise comparison**. we can fix it with a curiously recurring template pattern (CRTP), we have the `elements()` member function that acts as a customization point.

```cpp
template <class Derived>
struct ComparableElementwise {
   friend bool operator==(Derived const& a, Derived const& b) {
      return a.elements() == b.elements();
   }

   friend bool operator!=(Derived const& a, Derived const& b) {
      return a.elements() != b.elements();
   }
   /* ... */
};

// CRTP
struct Foo : ComparableElementwise<Foo> {
   int x;
   double y;
   std::string z;
   auto elements() const { // customization point
      return std::forward_as_tuple(x,y,z);
   }
};
```

now that we have this customization point, we can re-use it across other operations. like for printing all elements.

```cpp
template <class Derived>
struct PrintableElementwise {
   void print() const
   {
      auto const& e = static_cast<Derived const&>(*this).elements();
      std::apply([](auto const&... el) {
         ([&]{ cout << el << '\n';}(), ...);
      }, e);
   }
};
```

we should be careful over coupling, if we have the same name for the customization point, then we run the risk of creating coupling. when in doubt, we should use different names and forward to the common function, making things easier to change later on.

we can also have mixin mixins.

```cpp
auto const& e = static_cast<Derived const&>(*this).elements();
```

the code above is a downcast that we commonly do in CRTP, but is still weird arcane code that is repeated, so we should express the sameness in code.

```cpp
template <class> struct CRTPMixin;
template <template<class> class Mixin, class Derived>
struct CRTPMixin<Mixin<Derived>>{
   consteval auto& self() { return static_cast<Derived&>(*this); }
   consteval auto& self() const { return static_cast<Derived const&>(*this); }
};

template <class Derived>
struct PrintableElementwise : CRTPMixin<PrintableElementwise <Derived>> {
   void print() const
   {
      auto const& e = self().elements();
      std::apply([](auto const&... el) {
         ([&]{ cout << el << '\n';}(), ...);
      }, e);
   }
};
```

and now in C++23, we can be better with the new Deducing `this` syntax.

```cpp

struct PrintableElementwise{
   void print(this auto const& self)
   {
      auto const& e = self.elements();
      std::apply([](auto const&... el) {
         ([&]{ cout << el << '\n';}(), ...);
      }, e);
   }
};

struct Foo : PrintableElementwise {
   int x;
   double y;
   std::string z;
   auto elements() const {
      return std::forward_as_tuple(x,y,z);
   }
};
```

the problem is that we introduced inheritance, we have interface sameness when we didn't want it.

we still have other forms of "sameness" to look at. when we forward as tuple, we re-write the same members as the struct, but without reflection, we can't easily express the same elements.

### Mixin Language Support

in C++20, the spaceship operator is a mixin, but it's a single piece, a small fix that was introduced as a special thing that is hard to re-use. if we had reflection, it wouldn't be a special case.

### Qualifier Forwarding

we had "perfect forwarding" (universal reference) since c++11. it has a very strict syntax that needs to be maintained. it becomes worse with member functions, which need to created for each combination of const and volatile qualifiers, this is solved with "deducing `this`" in c++23.

```cpp
template <class T>
void foo(T&& t) {
   bar(std::forward<T>(t));
}

template <class Thing>
class OwningCollection {
   private:
      vector<unique_ptr<Thing>> things_;
   public:
      // pre c++23
      void for_each(invocable<Thing const&> auto && f) const
      {
         /*...*/
      }
};

//c++23
template <class Self, class Func>
void for_each(this Self&& self, Func&& f)
{
   for (auto&& ptr: self.things_)
   {
      std::forward<Func>(f)(std::forward_like<Self>(*ptr));
   }
}
```

### Name Forwarding

this doesn't exist in C++, but does exist in languages such as ruby, it allows exposing some functionality to the user, but without exposing the data itself. we can re-use names, but not limiting ourself.

### Back to DRY

in our abstractions, we have few pieces of knowledge, we can separate the ownership mechanism from the collection mechanism. it doesn't have to be a unique pointer, it can ba a shared pointer, small buffer optimization, or something else entirely.

for example, in the <cpp>std::vector</cpp> class, we have a customization for the allocator. the expand container "knowledge" is separable from how the storage is created, same with <cpp>std::queue</cpp>, which can use a <cpp>std::deque</cpp> as the underlying data point, but also further customize it itself. same with the smart pointers <cpp>std::unique_ptr</cpp>, which can take a unique deleter function which is separated from the creation of the data (which happens at the construction point).\
we can follow the rabbit hole of the unique pointer. are allocation and destruction really separate? why are shared pointer and unique pointer different? turns out shared pointers have an allocator for the control block. there also <cpp>std::allocate_shared</cpp> which puts the two together. something weird about type erasure.

### Concepts

C++20 concepts allow users to extract interface sameness without permission, they can force coupling that we didn't intend to. it's basing interface on names. they don't check for namespace, and the same name can mean different things across classes, even if the signatures are the same.

other stuff:

- normal functions
- macros and code generation
- customization point objects
- type erasure
- constexpr function - "sameness" for compile-time and runtime
- dependency injection - missing in C++, requires reflection
- aspect oriented programming - missing in C++, requires reflection
- decoration - missing in C++, requires reflection

</details>

## Bryce Adelstein Lelbach :: C++ Horizons

<details>
<summary>
Features we might see in the future.
</summary>

[C++ Horizons](https://youtu.be/gQ0n7taWtv4?si=irUiYfI2In-z_wnr)

C++20 is already in the field, with the big four changes:

- Modules
- Coroutines
- Concepts
- Ranges

C++23 is also fresh, with it's own set of changes

- More ranges
- Formatted output
- <cpp>std::mdspan</cpp>
- <cpp>std::expected</cpp>
- `import std`
- deducting `this`

there are still long term goals that aren't scheduled yet, but will hopefully change how we write C++ code

### Reflection and Meta-Programming

meta-programming creates source code as the result of code, reflective meta-programming creates code inside the same application that has he source code.

this is a possible example of reflection, we use the symbol `^` as the "reify" operator. it takes an entity and produces a reflected representation. we can then manipulate the date, using the `template for` compile time expansion format to iterate over each of the members. we then use another new operator to splice from the reflection back into an entity `:<reflection>:`.

```cpp
template <typename T> requires is_enum_v<T>
constexpr string to_string(T value)
{
   template for (constexpr meta::info e : meta::members_of(^T))
   {
      if ([:e:] == value)
         return meta:name_of(e);
   }
   return "<unamed>";
}
```

moving to another example, we use reflection to get the non-static data members, and then we do another `template for` to iterate over all the members and add them to the hasher.

```cpp
template <typename Hasher, typename T>
constexpr void hash_append(Hasher& hasher, T const& t)
{
   constexpr meta::info data_members = meta::members_of(^T, meta::is_nonstatic_data_member);
   template for (constexpr meta::info member : data_members)
   {
      hash_append(hasher, t.[:member:]);
   }
}
```

another example of a wrapping class that traces member calls via special splice operator `[#<reflection>#]`, with pack expansion, it generates a new class with the same interface as another one.

```cpp
template <typename T> requires is_class<T>
struct traced {
private:
   T Payload_;
   template for (constexpr auto e : member_functions_of(^T))
   [:protection_of(e):]: [:attributes_of(e):] [:return_of(e):] [#e#]
   (...[:parameter_types_of(e):] ... [#parameters_of(e)#]...)
   [:qualifiers_of(e):]
   {
      print("Calling {}::{}\n", name_of(^T), name_of(e));
      return payload_.[:e:](forward<[:parameter_types_of(e):]>{...[#parameter_names_of(e)#]});
   }
};
```

example of how we can use reflection to move between array of struct to struct of arrays.

### Pattern Matching (Selection)

we have the `switch` statement, which only acts on single integral value, and we have the `if` statement, which is general and arbitrary. pattern matching will give us `inspect`, a middle ground expression that matches values against patterns, and bind variables on success.\
it stops on the first match, not the best one. the matching has to be done against constant values, but they don't have to be integral values.

```cpp
switch(i)
{
   case 0: print("got zero"); break;
   case 1: print("got one"); break;
   default: print("don't care");
}

inspect(i)
{
   0 => {print("got zero");}
   1 => {print("got one");}
   _ => {print("don't care")};
};
```

an example of the Fibonacci function with pattern matching.

```cpp
unsigned Fibonacci(unsigned n) {
   return inspect(n){
      0 => 0;
      1 => 1;
      e if ( e > 47 ) => {throw overflow_error("too large");};
      a => Fibonacci(a-1) + Fibonacci(a-2);
   };
}
```

we can use more complex patterns, using decomposition

```cpp
inspect (p) {
   [0, 0, 0] => {print("on origin");}
   [x, 0, 0] => {print("on x-axis");}
   [0, y, 0] => {print("on y-axis");}
   [0, 0, z] => {print("on z-axis");}
   [x, y, z] => {print("{}, {}, {}", x, y, z);}
};
```

we could use pattern matching on types and variant, much like <cpp>std::visitor</cpp>, it can then be combined with de-composition.\
We could create our own extractor pattern, to match against all sorts of formats. it can be combiner with types and de-composition, and even with reflection. this can create a json serializer.

### Sending (Senders and Receivers)

C++ doesn't have a standard way to express where things should execute, and no standard model for a-synchrony.

_Schedulers_ are handles to execution contexts, _Senders_ are asynchronous work, and _Receivers_ process asynchronous signals.

we get the scheduler from somewhere (thread pool, distributed system), and we schedule work onto it. we can expand it with the pipe syntax to make it composable. the code is the same for a single CPU and a multiple distributed system with hundreds of GPUs.

```cpp
ex::scheduler auto sch = thread_pool.scheduler();

ex::sender auto begin = ex::schedule(sch);
ex::sender auto hi = ex::then(begin, [] {return 13;});
ex::sender auto add = ex::then(hi, [] (int a) {return a + 42;});

auto [i] = this_thread::sync_wait(add).value();
```

</details>

## Separator

</details>
