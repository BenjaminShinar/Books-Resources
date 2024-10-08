<!--
ignore these words in spell check for this file
// cSpell:ignore Rpass vectorize fopt, allt vectorizable Withdrawl prefetcher periph rodata configurators Misra
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Embedded

### Using C++14 in an Embedded **SuperLoop** Firmware - Erik Rainey

<details>
<summary>
Challenges of writing embedded code in amazons PrimeAir drone.
</summary>

[Using C++14 in an Embedded “SuperLoop” Firmware](https://youtu.be/QwDUJYLCpxo).

talking about Amazon _prime air_ drone for package delivery. it runs on c++14 software.

#### Introduction

_what is firmware?_ everybody has a slightly different defintion. there is a spectrum of complexity, a toaster firmware isn't the same as drone firmware, and there are different constraints for each case.

firmware:

> - installed into the internal memory of a processor
> - usually has 1 execution context (plus interrupts)
> - runs as highly privileged process
> - intentionally difficult to update
> - update comes as a single large "image" instead of piecemeal updates of individual parts
> - limited memory space environments
> - limited memory protection mechanisms

types of of firmware, from least to most complexity.

> - Bare metal
> - SuperLoop firmware
> - Real time operating system
>   - free RTOS
> - Reduced functionality operating systems
>   - ucLinux (linux for micro controllers)
> - Full operating systems, but with limited flexibility
>   - Symbian OS
>   - Some distributions of linux for routers (dd-wrt, tomato)

Firmware usually runs on micro controllers, mostly ARM. there isn't an MMU (so no virtual memory), but they have very fast SRAM memory. depending on the chip used, there may or may not be floating point operations. the computational power is much lower than what we have on modern computers, and they use a hardware watchdog.\
the firmware interacts with low level, sometimes slower busses, which is what we find in the industry. the device must have fast boot time, and be able to run for long extended periods with no issues. sometimes it must support a low-power more.\
might need to support ECC (_error correcting code_) memory and error injections.

#### What is a "SuperLoop" Firmware?

> - Is conceptually just a big `for` loop.
> - There is a single execution context
> - All the executable "tasks" are visited once in the SuperLoop then it repeats:
>
>   ```cpp
>   class Task {
>     public:
>       virtual void Execute(void)=0;
>   };
>
>   void TaskManager::RunAllTasks(){
>     for (auto & task : GetTaskList()){
>       task.Execute();
>     }
>   }
>   ```
>
> - No _software_ limit execution time (there is a hardware limit)
> - No guarantee of when the next start time is
> - Blocking constructs effectively halts the system
> - No idle task (this is a choice)
> - Consequently, always at 100% processor usage
> - Metrics (usually trailing window)
>   - Min/Avg/Max SuperLoop time
>   - Min/Avg/Max Task time

in the case of the _primeAir_ drone, the device doesn't have a heap memory, so all memory allocations are done upfront. There isn't a virtual memory on the microchip , so no threads, no processes and no scheduler. There isn't a user-model, and no file system. they also decided to remove exception handling, and to avoid using blocking system calls.\
The goal was to eliminate runtime errors by moving them into compile time.

> - no memory fragmentation error.
> - no runtime out of memory conditions.
> - no thread starvation, priority inversion, priority arrangements issues by the scheduler.

Less functionality means less things can go wrong, and a simple system is easier to reason about. it's more deterministic and repeatable (removing variability caused by the scheduler or background tasks of the operating system). this reduces the testing burden.

> "Allows you to focus on solving the hight level functionality _including_ account for failure modes **upfront**."

we also removes some undefined behavior, such as a full heap or an overly fragmented heap, or a kernel in a deadlock state, or an exception inside the kernel (and an exception within the exception handler). undefined behavior is not an option in a critical enviornment. this is part of the certification burden: all paths must be deterministic, and if there is a path that leads to undefined behavior, then it should be removed.\
Even after removing all that, we can still make use of c++ features.The stronger type system allows us to be sure about the behavior, c++ code is used in the entire chain all the way down to the bare metal, and the last C parts (assembler, builtins) are wrapped in good C++ interfaces.

The system still has a stack, both a normal mode stack and an optional priviledged mode stack. System calls are still available, with some "dangerous" operations behind a barrier (such as rebooting, logging fatal errors, data cache operations, fetching keys). There still is Vector Table and Handlers, for hard faults, bus faults, etc...\
They use a special "libc" version which is intended for embedded processors with further restrictions, such as removing floating point from `printf` (no `%f` option), removing 64bit options, and stubbing out functions which might effect memory. an example is `_sbrk`, which grows the top of the memory (`_end`).\
A special cross compiler is used, which makes only a few defaults assumptions, linker scripts must be customized and have strict limits on binary size which are asserted during build time. the goal is to move _runtime allocation_ problems into compile time.

they decided to avoid **ISR** (interservice routines) as much as possible, and for some cases there must still be code running during those ISRs. this means that peripherals are polled.\
In most cases, using the global memory is a bad practice, but not in this case. the boot path is also modified. the boot sequence order is important, we need to check what is available to us in terms of global objects, system calls, peripherals and which core system features are running.

The build process takes all the source code into a single binary, all types are known at compile time, and any unknown types in the communication channels are ignored. all protocols has a predefined enumerated type.

software patters and data structures

> - Ring buffers / FIFO / Circular buffer.
> - Queues (which are really ring buffers).
> - Containers which control _when_ objects are constructed.
> - Custom error code tuple-like objects (result, cause, location) with mechanism to capture the location of problem.
> - Bus drivers use Queues of Transactions.
> - A system for "orthogonal" errors collection in normally un-reportable locations (constructors, operators, void functions).

because of these constraints, the code ends up looking like a _finite state machine_ or a _state chart_, and there are many libraries and 3rd party tools which support this. the state chart is always written as a non-blocking and concurrent. it never blocks the system from runnings, and it best to process data in chunks. Function recursion can end up running the small stack, so they are to be avoided.

#### Embedded C++14 Language restrictions

No "forever" `for` or `while` loops. only one is allowed, and that is the SuperLoop in `main()`. loops must have definite bounds, and use timeouts or counters to enforce the exit condition from the loop as a fallback.\
Pointers should be avoided, because there is no MMU in a micro-controller, there might be real memory at address zero. so we can't count on `nullptr_t` dereferences as protection. obviously there isn't any code using exceptions (space contains, recursive issues), no `throw/noexcept` and not `try-catch`. Turning off RunTime TypeInfo, no `typeid()` and no `dynamic_cast<>`, and no dynamic memory at all, we can use compiler flags to wrap those calls into wrappers, and leave them unimplemented, so that all calls to memory allocation cause an error.\
The side effect of removing dynamic allocations is that many STL types are unusable, such as _std::vector_ which is resizable, _std::unique_ptr_ which allocates, and less obvious stuff such as _std::bind_.

The idea is to remove the _notion_ of allocators, not just this or that implementation of it. new programmers might run into **"Heap Withdrawl"**:

> stages of **"Heap Withdrawl"**
>
> - "I don't know how much memory I need at runtime, so I **need** a heap."
> - "I know host much 'worst case` memory I need at runtime but I still **would like** to use a heap for convenience, even though my inital allocation could still technically 'fail'."
>   - (somebody else can use the heap in runtime and run out of memory)
> - "Well I could use a statically allocated pool of memory fixed sizes and mange usage bits with an allocator interface."
>   - (this is reinventing the heap)
> - "Oh, that can still run out, so make an enum of preassigned pools!"
>   - (this is reinventing static memory allocation)
> - Store static memory with the objects that need it. Size it appropriately, protect against failure modes which could cause overruns.

luckily for us, the compiler and analyzer can help us, this requires knowing the compiler and the use case, and this might be a problem with 3rd party library, which might not support these modes. Static analyzers can check with even stricter rules to check the code.

another problem is trying to reinvent basic types/

```cpp
using my_uint8 = unsigned char; // is that true on all of your platforms?
using my_uint16 = unsigned short;
using my_int32 = long; // is that really 32bits on your platform?

uint32_t PointerToAddress(void* ptr){
  return reinterpret_cast<uint32_t>(ptr); // error on 64bit platforms
}

uintptr_t PointerToAddress(void* ptr){
  return reinterpret_cast<uintptr_t>(ptr); // works anywhere
}
```

they use references wherever they can, pointers are rare, and should be isolated to low level code, and with effort, they can be wrapped into interfaces or avoided.

```cpp
// declare as reference
Foo &foo = context.getFoo();
foo.bar();

// pass as a reference
void SomeFunc(Foo& foo);

// reference to peripheral via placement new into a reference
Peripheral& periph = *new(PERIPH_ADDRESS) Peripheral{};

// Drivers should take references to other Drivers or buffers
Driver(Timer& timer, std::array<uint8_t,1024>& dma_buffer);
```

in embedded learning, most people are still using the "C" mindset.

```cpp

// Don't half-way commit to C++ Naming
namespace Chip {
  enum class ComponentMode{ComponentModeFreeRunning};
  class ComponentPeripheral{
    static void InitializePeripheral(ComponentMode mode) = 0;
  };
}

// DON'T
Chip::ComponentPeripheral::InitializePeripheral(Chip::ComponentMode::ComponentModeFreeRunning);
```

> - Keep names DRY- don't repeat yourself.
> - Keep enums and classes properly scoped to the level they are used.
> - Keep functionality related to the peripheral hardware in the peripheral class.

so instead,

```cpp
// Do
namespace Chip {
  namespace Component {
    class Peripheral {
      enum class Mode: ... {FreeRunning};
      static void Initialize(Mode mode) = 0;
    };
  }
}

// Do
Chip::Component::Peripheral::Initialize(Chip::Component::Peripheral::Mode::FreeRunning);
```

Layers interact thorough pure virtual interfaces. so there is a good separation between implementation details and common set of objects. virtual function call overhead exists and we should keep an eye on it, but utility and design are often more important.

```cpp
// Interface
class Foo {
  public:
    virtual uint64_t GetCount(void) const = 0;
};

// Implementation
class Bar: public Foo {
  public:
  virtual uint64_t GetCount(void) const override {/*...*/};
};
```

but never have virtual diamond inheritance, it's a bad practice to have.

```cpp
class Foo {
  public:
    virtual uint64_t GetCount(void) const = 0;
};

class Bar: public Foo {
  public:
  virtual uint64_t GetCount(void) const override {/*...*/};
};

class Baz: public Foo {
  public:
  virtual uint64_t GetCount(void) const override {/*...*/};
};

class Gaf: public Bar, public Baz {};

Gaf g{};
g.GetCount(); // where is this from?
```

the `final` keyword is useful, it controls what can and what cannot be subclassed.

```cpp
// Interface to a controller
class IController {
  /*...*/
};

// Implementation, can't be subclassed
class ControllerImpl final: public IController {
  /*...*/
};

// attempt to subclass a `final`
class MyController: public ControllerImpl { // Error!
  /*...*/
};
```

use RAII, it's always important in C++, but doubly so in Embedded. when we use placement `new`, objects must be destructed manually!

```cpp
class DriverImpl {
  DriverImpl(volatile Peripheral& peripheral, ...) :
  periph(peripheral) {
    periph.configure(...);
    periph.enable();
  }

~DriverImpl() { //must be called explicitly
  periph.disable();
  }
};
```

inlining for only simple methods, `const` always helps, we need to know if there are side-effects to calling the functions. hardware behaves differently than normal c++, and it's up to us to know whether the method truly has no side effects. a function that has "read and clear" effect on registers should not be marked `const`.

(c++ example with mutable member values in a `const` method).

the `volatile` keyword is used to describe memory mapped registers and hardware, there might be a cascading effects.

(example of marking a member value as `volatile`, both in the constructor and as the member variable).

`mutable` is to e avoided, but sometimes might be necessary. the _std::atomic_ types are helpful, but there are possible pitfalls and gotchas!. _static storage_ defintions allows control of constructors, but sequence and lifetime must be thought through.

```cpp
// Somewhere in global memory
static std::aligned_storage<sizeof(Foo), alignof(Foo)>::type storage;
// ...

void main(void){
  // object initialization happens in preallocated global memory
  Foo& foo1 = *::new(&storage) Foo{/*...*/};

  // object initialization happens on the stack
  Foo foo2{/*...*/};
}
```

avoiding macros, mostly `#define` or `#ifdef`, replace with templates, `constexpr`. and try to use as many `constexpr` code as possible. have the code check itself with the compiler, but also have the linker remove the unreferenced code.

```cpp
// header
constexpr uint32_t foo(uint_t){/*...*/}

//source
namespace {
  constexpr uint32_t foo_test[] = {foo(X),foo(Y)};
  static_assert(foo_test[0]== 0x10, "Must be true or something is wrong");
} // anonymous namespace

// on constructors
class Foo Final {
  constexpr Foo(): /*...*/ {}
};
```

use `static_assert` as much as possible, make sure that your mental model fits the hardware reality in terms of `sizeof()`, `offsetof()`, `alignof()`.\
`alignas(x)` to make sure types that are relevant Data Cache or Bus are properly aligned to word boundaries or cache lines.

strict enumerations, using c++ `enum class`, with underlying type.

when defining constructors, be clear about what is possible, use the _rule of 5+1_, default, move, copy constructors, as well as move and copy assignment should be defined if existing, set to default behavior or deleted. expressing the intent to other developers. classes that shouldn't be moved (such as peripherals) shouldn't have those operations deleted. there are many edge cases about `new` and `delete`, especially with containers.\
**Operator overloads** are also a point to look at, the `operator""` allows defining literals,

```cpp
class furlong {
  public:
    constexpr furlong (long double v): value{v} {}
  protected:
    long double value;
};

namespace literals {
  constexpr furlong operator"" _fur(long double x) {return furlong {x};}
}

using namespace literals; //allows access to quote operator
constexpr std::array<furlong, 4> track_lengths = {2.0_fur, 5.0_fur, 6.5_fur, 9.0_fur};
```

be wary of narrowing warnings, prefer `{}` brace initialization over `()` to get a warning about narrowing. `explicit` constructors are always nice.

C++17 and C++20 aren't covered yet by the certifications, but when they do become, some other parts of the STL will become usable: `std::optional`, `std::variant`, new attributes `[[deprecated]]`, `[[deprecated]]`, `if constexpr` and `consteval` will help move code to compile time, and c++20 concepts will help with templates errors.

</details>

### Taking a Byte Out of C++ - Avoiding Punning by Starting Lifetimes - Robert Leahy

<details>
<summary>
Type Model, Object Model, implicit lifetime types ,std::bit_cast, std::start_lifetime_as. ensuring there are no invalid states of objects.
</summary>

[Taking a Byte Out of C++ - Avoiding Punning by Starting Lifetimes](https://youtu.be/pbkQG09grFw), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Slides-Final.pdf).

having a mental model of the memory model in c++. a struct with two members is model of the memory, instead of using offsets, we have fields. but in the end, it's only data in memory.

all of these have the same memory representation. the type system is a way to look at memory.

```cpp
std::byte[8];
std::uint16_t[4];
std::uint32_t[2];
std::uint64_t;
struct foo {
  std::uint32_t a;
  std::uint32_t b;
};
```

let's write some code, we check the size is the same, and we write a function that sends the same data in the two parameters.

```cpp
struct foo {
 std::uint32_t a;
 std::uint32_t b;
};

static_assert(sizeof(foo) == sizeof(std::uint64_t));

std::uint32_t bar(std::uint64_t& i, const foo& f) noexcept {
  if (f.a == 2) {
    i = 4;
  }
  if (f.a == 2) {
    return f.a;
  }
  return f.b;
}

int main() {
 foo f{2, 3};
 return bar((std::uint64_t&)f, f); // cast the variable into a reference to another type
}
```

when we run the function in an actual compiler, the code is executed as if there were two different objects. this it **type punning**

> No Type Punning\
> An object within its lifetime may only be accessed in certain ways:
>
> - Through a reference to its type (addition of cv qualification allowed)
> - Through a reference to its signed or unsigned equivalent
> - Through a reference to char, unsigned char, or std::byte
>
> Any other access modality is undefined behavior.

if we want to see this in action, we can use a different example, where the arguments are of related types, now the compiler must check this and we will see a different results.

```cpp
std::uint32_t bar(std::uint32_t& i, const foo& f) noexcept {
  if (f.a == 2) {
    i = 4;
  }
  if (f.a == 2) {
    return f.a;
  }
  return f.b;
}

int main() {
 foo f{2, 3};
 return bar((std::uint32_t&)f, f); // cast the variable into a reference to another type
}
```

> C++ Has an Object Model\
> Bytes supply storage for objects.\
> Objects have lifetimes:
>
> - Duration of storage is not necessarily the same as object lifetime.
>
> Accessing object outside lifetime is undefined behavior.

in c++17. this code is about storage and not lifetime.

```cpp
const auto ptr = (int*)std::malloc(sizeof(int) * 4);

if (!ptr) {
  throw std::bad_alloc();
}

for (int i = 0; i < 4; ++i) {
  ptr[i] = i;
}
```

but a similar looking good is very uncomfortable. they are no strings in this code.

```cpp
const auto ptr = (std::string*)std::malloc(sizeof(std::string) * 4);

if (!ptr) {
  throw std::bad_alloc();
}

for (int i = 0; i < 4; ++i) {
  ptr[i] = std::to_string(i);
}
```

> C++ Types May Have Invariants\
> One of the core value propositions of
>
> - C++ Invariants are established by constructors.
> - Invariants are maintained by members.
>
> Some types don’t have such strict requirements:
>
> - Contain basic values
> - Don’t maintain complicated (or any) invariants
>
> Such types are trivial types.\
> Trivial types still have lifetime

for integers, there are no pre-post conditions, there is no constructor or destructor for a double.

c++20 has **implicit-lifetime types**, which are either aggregate types (such as an array) or types that have at least on trivial constructor or destructor.

and now some operations can create those types without undefined behavior.

in the examples above, integers are implicit life time type, but std::string isn't, so the undefined behavior is only for the "problematic" example, and we can address it with placement `new`.

```cpp
const auto ptr = (std::string*)std::malloc(sizeof(std::string) * 4);

if (!ptr) {
  throw std::bad_alloc();
}

for (int i = 0; i < 4; ++i) {
  // ptr[i] = std::to_string(i); // UB
  new(ptr + i) std::string(std::to_string(i));
}
```

with this change, we can remove some previously undefined behaviors from the world.

getting an integer from a buffer of bytes.

```cpp
int baz(const void* ptr) noexcept {
  return *static_cast<const int*>(ptr);
}

int baz2(const void* ptr) noexcept {
  int retr;
  std::memcpy(&retr, ptr, sizeof(int));
  return retr;
}

int baz3(const void* ptr) noexcept {
  alignas(int) std::byte buffer[sizeof(int)];
  const auto retr = std::memcpy(buffer, ptr, sizeof(int));
  return *reinterpret_cast<int*>(retr);
}
```

the code above is an over greedy version of `std::bit_cast`, but our version allow for pointers to types of any size to be casted into any other object, rather than restrict it to both types having the same size in memory.

(a part about something the compiler isn't allowed to modify)

in c++23, there are now new operations, which explicitly starts the lifetime of an implicit-lifetime type.

- `start_lifetime_as`
- `start_lifetime_as_array`

> Ending an Object’s Lifetime\
> Lifetime can end in the usual ways:
>
> - Object with automatic storage duration goes out of scope
> - Object with dynamic storage duration is deleted
>
> Can also end in other ways:
>
> - Pseudo-destructor call:
>   - t.~T()
>   - ptr->~T()
> - Reuse of backing storage

this a special operation in the standard, that has almost no usage.

> std::launder\
> Reusing storage invalidates pointers and references to the old object:
>
> - Unless the old and new objects are “transparently replaceable”
> - Pointers point to the storage but no longer to the object
>
> std::launder obtains a pointer to the object from a pointer to the storage.

(and now there are more complications)

> Summary:
>
> - Bytes which constitute an object reside in storage.\
> - Meaningfulness of the concept of an object not necessarily related to storage.
> - All objects have lifetimes regardless of how trivial they are.
> - Implicit lifetime rules enable zero copy techniques with well-defined behavior
> - As with all low level techniques care must be taken.
> - Potentially-dangerous operations can and should be factored out and isolated.
> - Remainder of code is clean, correct, efficient, and well-defined.

</details>
 
### Introduction to Hardware Efficiency in Cpp - Ivica Bogosavljevic

<details>
<summary>
Ways to get better performance by writing hardware friendly code
</summary>

[Introduction to Hardware Efficiency in Cpp](https://youtu.be/Fs_T070H9C8)

making software fast, that's one of the reasons C++ is used. there are some ways to get performance through software:

- software architecture
- avoiding unneeded work
- good programming language practices.

we also need to use the hardware efficiently to get peak performance.\
HW efficiency has two major bottlenecks

- computationally intensive (core bound) - **CI**
- memory intensive (memory bound) - **MI**

we won't look a disk bound or network bound bottlenecks. there are some tools to determine if the program is cpu bound or memory bound. those are moe tools are mostly profilers. the tools can look at data cache hits and misses to determine if it's memory intensive, or by check code prediction and branches to determine if its core bound.

a core-bound loop:

- processing simple data types
- the data is stored in an array or vector
- sequential access to values

if we access classes (structs), it's more likely that the loop is memory intensive, or if it's using data structures like trees and maps.

we find core-intensive code in domains such as:

- image, audio and video processing
- machine learning
- telecommunications
- scientific computing

#### Switching Memory intensive to Computationally Intensive and Optimizing it

in general, it's better to be **CI** than **MI**. so if we can, we should convert code.

##### Struct Of Arrays

one way to do it is by changing memory layout. if we have an array of classes, we can can convert it to hold the different data members of each object as elements in a array. moving from a ArrayOfStructs to a StructOfArrays.

```cpp
class student{
  std::string name;
  double average_mark;
};

std::vector<student> students_a; // array of structs

class students{
  std::vector<std::string> names;
  std::vector<double> average_marks;
};

students students_b; // struct of arrays
```

this conversion need to be done manually, but for calculating the average_marks it's better to be core intensive for performance reasons, and benchmarking shows that StructOfArrays is faster.

##### Loop Interchange

if we have a matrix (two dimensional data), accessing it column-wise is memory bound, but row-wise is core-bound. so if we change our access pattern, we will get better performance. this way, arrays are always accessed sequentially, which means we are no longer bound by memory (cache misses), so performance is faster.\
_(data is stored sequentially in rows, so every time we need a new line we need to move in memory)_

```cpp
// column wise
for (int i = 0; i < n; i++) {
  for (int j = 0; j < n; j++) {
    for (int k = 0; k < n; k++) { // switch row in b[]
      c[i][j] = c[i][j] + a[i][k] * b[k][j];
    }
  }
}
// row wise
for (int i = 0; i < n; i++) {
  for (int k = 0; k < n; k++) {
    for (int j = 0; j < n; j++) { // still using the same rows in all loops
      c[i][j] = c[i][j] + a[i][k] * b[k][j];
    }
  }
}
```

#### Vectorization

with CPU bound code, we can get performance boosts by employing vectorization. modern cpu can process more than one piece of data in a single instruction. this is called **SIMD** - _Single Instruction, Multiple Data_.

sometimes the compiler can create the optimized vectorized code itself - auto vectorization, this requires the compiler to have the correct flags and architecture known.

```cpp

for (int i = 0; i < n; i++) {
  A[i] = B[i] + C[i];
}

for (int i = 0; i < n; i += 4) {
  // Loads 4 Integers into a SIMD register B_tmp and C_tmp
  vec4 B_tmp = load4(B + i);
  vec4 C_tmp = load4(C + i);
  // Adds together vectors of four integers
  vec4 A_tmp = add4(B_tmp, C_tmp);
  // Stores four intergers
  store(A_tmp, A + i);
}
```

flags:

- clang: `-Rpass=loop-vectorize`,`-Rpass-analysis=loop-vectorize`
- GCC: `-fopt-info-vec-allt`

not all loops are easy to vectorize, some pre-requisites:

> - Simple data types.
> - Good memory access pattern - accessing memory sequentially.
> - Independent iterations (loop can run in any direction).
> - Number of iterations can be calculated before the loop starts.
> - Doesn't have a lot of conditionals inside the loop.
> - No pointer aliasing.

search loops can't be vectorized automatically, as the numer of iterations isn't known in advance. an example of pointer aliasing is when the two ranges overlap in memory. we can use the compiler report flags to see how much of our code is vectorized.

> - Too many conditionals or complex loop:
>   - Split loop into vectorizable and non-vectorizable parts - **loop fission**
>   - Move loop invariant conditions outside of the loop - **loop un-switching**
> - leaving the loop earlier: happens with search loops
>   - The solution is to apply a technique called **loop sectioning** to enable vectorization.
> - Bad memory access pattern
>   - Moving to structure of arrays for class/struct processing loops.
>   - **Loop Interchange** for matrix processing loops.
> - Pointer aliasing
>   - Use `__restrict__` keyword to mark the pointers as independet of one another.

#### Optimizing memory intensive code

cpu are faster than memory, so they end up being _data-starved_. many OOP programs use memory inefficiently. a solution is to use **data cache memory**, data which is much faster to access.

data cache misses common causes:

> - Dereferencing a pointer for the first time: a very common operation, using `operator*` and `operator ->`.
>   - Pointer data members
>   - Heap Allocated objects
>   - Vector of pointers
>   - Chasing pointers (like in a linked list)
> - Accessing a member of array (or another data structure) in a random fashion.
>   - `for (int i = 0; i < n; i++) {histogram[a[i]] ++;}`
>   - Binary search
> - Lookup in a hash map or a binary tree (_std::set_, _std::unordered_set_, _std::map_, _std::hash_map_).

in contrast, data cache hits typically happen:

> - Iterating through an array/vector sequentially.
>   - Smaller classes -> Higher data cache hit rate.
>   - Peak performance is achievable only when working with arrays.
> - Keeping the data set small.
> - What is accessed together should be close neighbors in memory.
> - What is accessed one after another should be laid out in memory one after another.
>   - linked lists (nodes are allocated next to one another).
>   - binary tree (at least have one child as neighbor).

example of finding minimum and maximum in an array of 100 million elements:

- finding both in the same loop is faster than running two loops. (STL has `std::minmax` to do this).
- it's better to iterate over the data just once, the bottleneck is memory bandwidth.

linked lists memory layout

- random layout - nodes are scattered across the memory layout.
- compact memory layout - all nodes are in the same block of memory
- perfect layout - nodes are ordered sequentially in memory, each node is a the neighbor or `->next()`. this is basically a contiguous memory block like vector or array.

we can control the memory layout using a custom allocator. the performance boost is achieved thanks to the CPU _data prefetcher_, which is a component that tries to get the data before it's needed.

the next example is a vector of pointers (which we use for polymorphism), this again can have optimal layout (neighboring pointers point to neighboring data). a way to fix this is by using a vector of values:

> - All memory allocated in a single block.
> - Sequential access to objects -> sequential access to memory addresses.
> - No calls to `malloc`/`free`.
> - No virtual dispatching mechanism to slow things down.
> - Enables small function inlining because type is known at compile time.
> - Downside: no polymorphism

(note: we can use std::variant to have polymorphism by employing the visitor pattern )

example class size and members layout. we have a class that we can control the total size of the class itself, and we can control the distance between the members. the larger the class, the slower the code, but we also see the effect in the function that calculates the visible size, which is effected by the padding inside the class itself. this padding doesn't effect the function that doesn't use the boolean flag.

```cpp
template <int pad1_size, int pad2_size>
class rectangle {
  private:
  bool m_visible;
  int m_padding1[pad1_size];
  point m_p1;
  point m_p2;
  int m_padding2[pad2_size];
};

template <typename R>
int calculate_surface_all(std::vector<R>& rectangles) {
  int sum = 0;
  for (int i = 0; i < rectangles.size(); i++) {
      sum += rectangles[i].surface();
  }
  return sum;
  // affected only by the class size (pad1_size + pad2_size)
}

template <typename R>
int calculate_surface_visible(std::vector<R>& rectangles) {
  int sum = 0;
  for (int i = 0; i < rectangles.size(); i++) {
    if (rectangles[i].is_visible()){
      sum += rectangles[i].surface();
    }
  }
  return sum;
  // affected the class size (pad1_size + pad2_size) and by the padding inside the class (pad1_size)
}
```

data is brought to the cpu in blocks (usually 64 bytes), so large classes need more bandwidth. data members which are accessed together should be packed together, so they'll be on the same data cache.

there are other ways to optimize memory access, but in some cases it won't help: small data sets that fit L1 cache, large data set that don't fit the largest cache (Last Level Cache), short lived data (memory optimized classes usually have more complicated constructors).

</details>

### Killing C++ Serialization Overhead & Complexity - Eyal Zedaka

<details>
<summary>
Presenting the zpp::bits library for serializing c++ objects.
</summary>

[Killing C++ Serialization Overhead & Complexity](https://youtu.be/G7-GQhCw8eE), [slides](https://docs.google.com/presentation/d/1sYwEH63mtvXpcHFH6y_97yVFe9HTs-Btuel9fZvL34o), [github](https://github.com/eyalz800/zpp_bits).

this talk is about the **zpp::bits** serialization libary. it will be simple to use, won't have overhead if it's not used, support RPC (remote procedure call) to have it be able to be called from a different machine, and it won't use exceptions.

Object serialization is turning c++ objects into a sequence of bytes, this is used to send them over the network, to save them to disk, etc...

one myth about C++ serialization is that they have zero runtime overhead, if we look at popular libraries, then there is a difference in performance, so it can't be so simple. Another statement is that C++ serialization doesn't fit embedded systems, this can be because serialization rely on exceptions for error checking, and the size of the library might not fit the embedded machine.

the _zpp::bits_ library uses an input and out archives, anything that isn't fixed size (like lists, strings, etc...) has the size (count) before the data itself.\
making everything available for inline, header only, avoiding virtual function calls. using concepts and templates to be fully generic (std::span, std::array, c array). customizable code by using `if constexpr` and using return code to avoid exceptions. we get around some overhead by forcing inlining (with non standard attributes). we use `constexpr` functions with _std::bit_cast_ to get some stuff.

Next part is reflection. we don't want each class to have a serializer/deserializer functions, we take could advantage of structured bindings (by using a feature that might get standardized in the future - _variadic structured bindings_), but instead we use a visitor pattern. we first need to get the number of members in a class, there is some stuff to play with the _requires_ clause. so instead we use a type erasure with _std::any_.

for Remote Procedure Call: as the basic part, we use _std::variant_ and the visitor pattern. we can also make things better with a special _zpp::bits::rpc_ object that takes bindings, this uses some wrapper classes, different for the client and the server.

the next part is cross language serialization. the library supports serializing into other types as well.

</details>

### Modern C++: C++ Patterns to Make Embedded Programming More Productive - Steve Bush

<details>
<summary>
Making Embedded code easier (and more pleasant) to write and maintain.
</summary>

[Modern C++: C++ Patterns to Make Embedded Programming More Productive](https://youtu.be/6pXhQ28FVlU), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Cppcon2022-slides-Bush.pdf), [github](https://github.com/sgbush/cppcon2022).

(Modern C++ to Impress Your Embedded Dev Friends)

> Making Embedded code easier (and more pleasant) to write and maintain.

#### GPIO Configuration

GPIO - the IO pins on the embedded board - the pin number, mode, speed, alternate fuctions and pulldown. in the old way it's procedural, and it uses constant, which aren't typed.\
The declarative way to configure GPIO creates a struct (using enums), push all of them into an array. everything is done in one place. they use weak enums because they are implicitly converted to the underlying type. It's better to do everything in one place.\
We might need to pass the GPIO reference around, so we also have stuff for that.

#### Compile-Time Lookup Tables Generation

in embedded system, in many cases we have constant arrays which we need to bring in, it could be done with spreadsheets (and then copied to code), or created in a script and then to a header file, it's all very error-prone.\
with compile time expressions, we can have the code create the data as part of the build process, so it's better to have the lookup tables as part of the code which uses them. that way the lookup table ends up in the nonvolatile "flash" part of the memory (".rodata" - 'read-only data"), rather than the RAM.

#### Code with Human-friendly Numeric Structures

address like structures - frequencies, ip addresses, MAC addresses, usb string descriptors... we store the data as bytes, but we think of them in a specific format. we can have custom literals that convert the canonical human text into data.

#### Lean stream-based IO

> "Using stream based IO, but skipping the library"

avoiding overhead of the IOstream library, we create a wrapper over a FileStream and we write the `<<` operator for our types. this way we reduce the size of the binary that we push into the embedded board.

#### Heap memory management for "allocate-once"

> "Dare to use the heap"

we usually don't use the heap in embedded applications, runtime behavior can cause heap exhaustion, and long runtime can cause fragmentation, and in general, heap errors have no graceful resolution. so we end up losing some nice data structures such as vector, list or deque.

but if the size of the data is known at the start of the program, and is never de-allocated, we can use arena-allocators: (suited for monotonic, allocate-once applications). this can be done by overriding the global new operator or overriding the operator on a "per-class" basis.

#### Use _std::chrono_ like a boss

unlock std::chrono features, configuring timers from the hardware

#### A _std::random_ topic

> "implement _std::random_device_ and get free library code"

creating a multi-modal generator

</details>

### Overcoming C++ Embedded Development Tooling Challenges - Marc Goodner

<details>
<summary>
Tools for embedded development.
</summary>

[Overcoming C++ Embedded Development Tooling Challenges](https://youtu.be/8XRI_pWqvWg), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Overcoming-embedded-tooling-challenges.pptx)

#### Embedded Tooling Today

many silicone vendors, most of which have their own IDE and tools (propriety compiler, flash/debug tool) and others have "configurators" for the specific hardware, and occasionally they support direct support for makefile/cmake exports. having a custom IDE is a good way to package everything together, but it makes integrating with the tool chain more difficult.

vendor tools are designed to work with embedded projects, so they have specialized compilers, hardware aware debugging, safety certification.

CMSIS - reusable components for embedded.

#### Using your IDE for Embedded (CLion, Visual Studio, Visual Studio Code)

tips for using Cmake, using vcpkg artifacts (sets of packages for a dev env: compilers, linkers, build system,platform SDK) in a manifest.

demo of a cmake template in github. cmake presets, manifests, boot strapping from source:

```ps
. ~/.vcpkg/vcpkg-init.ps1
vcpkg activate

cmake --list-presets
cmake --preset <preset name>
cmake --build --preset <preset name>
```

Peripheral view - mapping of external registers.

demo of NXP SDK - rtos example in visual studio. extensions for visual studio code.

cloud based IDEs and other tools.

#### Continuous Integration

ci-cd wizards - build image containers, static analyzers, build tools and chains.

using multi stage docker files, getting the tools into the containers. github workflows, github actions.

simulation enviornment ReNode. runs the same binary as devices.

</details>

### Simulating Low-Level Hardware Devices in Cpp - Ben Saks

<details>
<summary>
Creating a simulator for embedded devices to simulate side effects of registries.
</summary>

[Simulating Low-Level Hardware Devices in Cpp](https://youtu.be/zqHvN8xpuKY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Simulating-Low-Level-Hardware-Devices-Ben-Saks-CppCon-2022.pptx)

embedded software is designed to run on a target hardware, but developing for the hardware has difficulties.

- takes time to upload code to the hardware
- limited capabilites to test and debug the code

we want to have a simulator to run as much of our code as possible.

in this talk we will have a simulated code, and we'll have some operator overloads to change the registries (the side effects of the program).

> Our example program uses two objects representing hardware devices:
>
> - _uart0_, a pointer to a UART object that represents a serial port, and
> - _display_, a pointer to an object of type _seg7_display_.
>
> We’ll focus on simulating the hardware for UART 0.
> However, the simulated 7-segment display uses the same techniques.

```cpp
int main() {
  // object definitions for uart0 and display
  for (;;) {
      char c;
      if (uart0->get(c)) {
          display->put(c);
      }
  }
}
```

#### Classes and Memory-Mapped hardware

> - These slides use a variation of the techniques described in CppCon 2020’s Memory-Mapped Devices as Objects by Dan Saks. [CppCon 2020]
> - This session focuses on using operator overloading to create simulated hardware registers that behave much like real hardware registers.
> - CPUs typically communicate with external devices via device registers.
> - **A device register** is circuitry that provides an interface to a device.
> - Like most modern hardware, the E7T uses memory-mapped device registers.
> - That is, the hardware assigns each device register an address in the memory space.

each register has an address in the memory space, so if we know the address, we can create a pointer to it and write code as if it was just another c++ object.

```cpp
using device_register = std::uint32_t volatile;

device_register *const ULCON {
    reinterpret_cast<device_register *>(0x03FFD000)
};

// or as a reference

device_register_ref &ULCON {
    *reinterpret_cast<device_register *>(0x03FFD000)
};
```

each serial port (UART) has the same layout

| Offset    | Register | Description                 |
| --------- | -------- | --------------------------- |
| 0x00 (0)  | ULCON    | line control                |
| 0x04 (4)  | UCON     | control                     |
| 0x08 (8)  | USTAT    | status                      |
| 0x0c (12) | UTXBUF   | transmit buffer             |
| 0x10 (16) | URXBUF   | receive buffer              |
| 0x14 (20) | UBRDIV   | baud rate divisor (control) |

we can make a class out of them, which controls interacting with the device.

reading from UART, one character at a time, each time setting the RDR ("Receive Data Ready") bit in the status register.

```cpp
constexpr std::uint32_t UART::RDR {0x20};

bool UART::get(char &c) {
  if ((USTAT & RDR) == 0) {
      return false;
  }
  c = static_cast<char>(URXBUF);
  return true;
}
```

if we want the same code to run both on the target device and on the simulated device, we need to make sure the mappings are set to the correct location. to do that we use override the `operator new` operator so we have a custom `new`. we don't want to allocate memory, so we tell the operator to always point to the same memory address.

we must override the `operator delete`, because we don't want to call the global operator delete

```cpp
class UART {
public:
  void *operator new(std::size_t) {
      return reinterpret_cast<void *>(address);
  }
  void operator delete(void *){}
  /* ... */
private:
  static constexpr uintptr_t address = 0x03FFD000;
  /*...*/
};
```

so now our code is:

```cpp
int main() {
  UART *uart0 = new UART;
  seg7_display *display = new display;
  for (;;) {
      char c;
      if (uart0->get(c)) {
          display->put(c);
      }
  }
}
```

and it should run properly on our target platform, but we want to have a simulated environment. it's easier to run development cycles, and we want to trigger the error flows.

#### Creating the Simulator

we could have simulated UART and DISPLAY classes, which read from a file or from a different source. in the target hardware, when we access the registers, there are side effects. in the simulation, we need to produce similar effects.

- on-read side effect - such as altering the status bits
- on-write side effect - such as enabling or disabling UART when writing to UCON

so we wrap the device class and provide a conversion operator and over ride the equality and bitwise operations. (we don't allow writing to return the expression, because then it could be used to read the value without triggering the side effects)

```cpp
class device_register {
public:
  operator std::uint32_t(){
    // trigger on-read side effect
    return value;
  }
  void operator=(uint32_t v){
    value = v;
    // trigger on-write side effect
  }
  void operator&(uint32_t v){
    // trigger on-read side effect
    value &= v;
    // trigger on-write side effect
  }
private:
  uint32_t value;
};
```

we need to support different side effects. but we can't use virtual functions or templates, as our code requires all of our device registers to be of the same type inside UART class. we don't want any members so that we won't change the constructor for the real code, so we can't inject code.

but we can work with the locations. we could also have `std::function<void(std::uint32_t)>` for greater flexibility.

```cpp
class device_register {
public:
  using effect_handler = void (*)(std::uint32_t);
  struct effect_handlers {
    effect_handler on_read;
    effect_handler on_write;
  };
};
```

so, for example, clearing the RDR when a character is read will look like this:

```cpp
void on_URBXUF_read(std::uint32_t v) {
  USTAT &= ~UART::RDR;
}

class device_register {
public:
  operator std::uint32_t() {
    effect_handler on_read = get_read_handler(this);
    if (on_read) {
      on_read(value);
    }
    return value;
  }
};
```

we associate the registers and the side effects using a dictionary to map between the address and the side effect. this map will never run on the target hardware, so we can use dynamic memory, and it's not subject to platform specific problems. we again have to change the `operator new` to associate the intended side effects. we change a bit of the code to use inheritance (but without virtual functions) and compile time pre-processors to change the base class for simulation builds and the real build.

some issue with recursive calls, because of all the operator overloading, so we create another device_register base class to expose back-doors. we can now simulate the code as we want.

</details>

### Parallelism Safety-Critical Guidelines for C++ - Michael Wong, Andreas Weis, Ilya Burylov

<details>
<summary>
Trying to add guidelines for Parallel code.
</summary>

[Parallelism Safety-Critical Guidelines for C++](https://youtu.be/OD2huQx0Gco), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/cppcon-2022-safety-guidelines-for-C-parallel-and-concurrency-1.pdf).

> Agenda
>
> - Adding safety to parallelism for both MISRA and C++ CG
>   - This year: focus on what we intend to do for C++CG by hazards
> - Deep dive to C++CG rules
>   - Deadlocks and rejected rules
>   - Lifetime violation and new/modified rules
> - CG+MISRA: the close to ideal safety partners
>   - Ongoing analysis of future C++ parallelism for safety

#### Adding Safety to Parallelism for Both MISRA and C++ CG

we need safe code for embedded machines, such as drones, medical devices and even nuclear reactors. the safe code should parallel, as it's expected to run on critical machines that need to be fast.

there are a few existing coding standards, and this talk will be about integrating Misra and the C++ core guidelines. they started by reviewing existing rules and deciding which is worth keeping. rules can be reviewed by humans and machines.

they ended up with 24 rules, which we should consider adding to the core guidelines (which doesn't have enough rules for parallelism). some rules might not fit into the core guidelines "philosophy" but can be kept as Misra rules.

#### Deep Dive to C++CG Rules

The parallelism rules are categorized into "deadlocks" and "lifetime" guidelines.

##### Deadlocks

> Current CG focused on deadlocks prevention:
>
> - CP.20: Use RAII, never plain lock()/unlock()
> - CP.21: Use <cpp>std::lock()</cpp> or <cpp>std::scoped_lock</cpp> to acquire multiple mutexes
> - CP.22: Never call unknown code while holding a lock (e.g., a callback)
> - CP.50: Define a mutex together with the data it guards. Use <cpp>synchronized_value\<T></cpp> where possible
> - CP.52: Do not hold locks or other synchronization primitives across suspension points

they suggest a modification for the CP.20 rule to make it clearer and also encompasses timed locks. for the rule about multiple mutex, the suggest adding <cpp>std::try_lock</cpp> to the wording of the guideline.

There is a Misra rules against destroying mutex object while being locked, but this rule is not meant for humans, it's not something that will ever be done on purpose. this is a bug no matter what, and should be handled by code analysis, rather than becoming part of the guidelines.\
there is also a rule about the order of locking and unlocking - saying the order of nested unlocks should form a DAG (directed acyclic graph), which is also not something a human thinks about, and even tools might have a problem to detect this.

##### Lifetime Violations

> Misra: A thread shall not access objects whose lifetime has expired.

but it's not fit for the core guidelines, because it's nearly impossible to have local reasoning about it, and for tooling it requires a system-wide check, which is considered too costly to be a reliable part of the development process.

> Current CG rules for lifetime:
>
> - CP.23: Think of a joining thread as a scoped container
> - CP.24: Think of a thread as a global container
> - CP.26: Don’t `detach()` a thread
>
> They promote scope-based reasoning about lifetimes

In this example, the "obj" object is accessed by a thread, so it must remain "alive" until the end of the thread. this example is safe, because it uses a joinable thread (<cpp>std::jthread</cpp>) which is bounded inside the function. this will not be safe if the thread was detached.

```cpp
void f() {
 MyClass obj;
 std::jthread t([&obj]() {
 do_work(obj);
 });
}
```

the rule for classic threads (<cpp>std::thread</cpp> - not automatically joined on destructor) disallows capturing local variables, so it's conceptualized as global container. this ties into a "lifetime" profile idea for formalize object lifetime safety. it argues that if threads are scoped, then they are checked the same way containers are checked (and the objects they capture) and follow the same lifetime models.

#### CG+MISRA: the Close to Ideal Safety Partners

going over c++20 features relating to parallelism and concurrency, marking which are safe and ready to be used.

</details>

##

[Main](README.md)
