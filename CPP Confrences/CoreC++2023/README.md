<!--
// cSpell:ignore objdump Browsable Guttag nsenter setcap getpcaps fsanitize Nlohmann httplib Dennard alon Metaparse Lexy ctre idents
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

## Separator

</details>
