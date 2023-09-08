<!--
// cSpell:ignore objdump Browsable Guttag nsenter setcap getpcaps fsanitize
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Core C++ 2023

<!-- <details> -->
<summary>
Israel C++ Convention.
</summary>

[playlist](https://www.youtube.com/playlist?list=PLn4wYlDYx4bs0p9S6aFvKaASoCLFVwt_E), [Schedule](https://corecpp.org/schedule/).

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
the assembly code has function calls in assembly, the instuction is call, and the op-code is zero. this is because the compiler doesn't know where the actual code is, and it needs the linker to fill it in. we could also add the `--reloc` flag to the object dump and see how the code expects the re-locations should work.\
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

C++ stared with two goals - efficient use of hardware (like C), and managing complexity (based on simula). it also meant enforcing argument type checking. a different jey idea is to "represent concepts in code". <cpp>RAII</cpp> - resource acquisition is initialization, not only memory resources, also file handles, locks, sockets, shaders.\
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

## Separator

</details>
