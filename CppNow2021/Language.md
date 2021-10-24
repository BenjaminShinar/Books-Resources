<!--
ignore these words in spell check for this file
// cSpell:ignore Schödl Lakos Vittorio Ivica Bogosavljevic mmap strided Emde Dwyer prvalue


-->

Language

## The C++ Rvalue Lifetime Disaster - Arno Schödl

<details>
<summary>
Using rvalues and lifetime extension can fail us. member accessor for rvalue objects shouldn't return a reference.
const & should never have bound to rvalues.
</summary>

[The C++ Rvalue Lifetime Disaster](https://youtu.be/sb7cj-3l1Kc)

the use of rvalue references and move semantics. replace copying with moves when possible tp avoid memory operations.

also used to manage lifetime, as well as for c++20 ranges

```cpp
auto rng=std::vector<int>{1,2,3} | std::views::filter([](int i){return i%2==0;}); //doesn't compile
```

this doesn't compile for rvalue

### Pitfalls

can't move from a const value, and moving will mess with NRVO (names return value optimization) and make it harder for the compiler to elide the construction.

```cpp
A foo()
{
    const A a;
    return std::move(a); //error!
}
A foo2()
{
    A a;
    return std::move(a); //works, but we are messing with RVO
}
A foo3()
{
    const A a; //doesn't matter if we're const or not, elision works
    return a;
}
```

but if we have two possible values, we can't do NRVO, and we also can't do move (because of const).

```cpp
A foo4()
{
    if ()
    {
        const A a;
        return a;
    }
    else
    {
        const A a;
        return a;
    }
}
```

and here we can't do copy/move ellison, because it's member variable. we also can't do a move, members don't automatically become rvalues.

```cpp
struct B {
    A m_a;
};
A foo()
{
    B b;
    return b.m_a;
    //return std::move(b).m_a; //this would work.
}
```

recommendations:

> - make return variables non-const
> - use clang -Wmove flag

### Temporary Lifetime Extension

```cpp
struct A;
struct B {
private:
A m_a;
public:
const A& GetA() const &
{
    return m_a; //return by reference
}
};
B b;
const auto & a = b.getA();
struct C{
    A getA() const &; // return by value
};
C c;
const auto & a1 = c.getA();
const auto & a2 = B.getA();
```

if we capture something with const reference, it can extended the lifetime of the object it's capturing.

_std::min_ doesn't take rvalue-ness into consideration, it returns a lvalue reference. a will dangle.

```cpp
bool operator<(const A&, const A&);
struct C
{
    A getA() const&;
} ;
C c1,c2;
//...
const auto & a = std::min(c1.getA(),c2.getA()); //a will dangle
```

lets' have a min function that keeps rvalue references using perfect forwarding. but it still doesn't work

```cpp
namespace out
{
    template<typename Lhs,typename Rhs>
    decltype(auto) min(Lhs && lhs,Rhs && rhs)
    {
        return rhs<lhs ? std::forward<Rhs>(rhs)? std::forward<Lhs>(lhs);
    }
}
```

lifetime extension only works where there an object.

an example with forwarding a return and _'decltype(auto)'_

the advice is to stop using temporary life time extension,
what we want is :

> automatically declare variable
>
> - _auto_ if constructed from value or rvalue reference
> - _const auto &_ if constructed from lvalue reference

he suggest this macro code instead of lifetime extension.

```cpp
template<typename T>
struct decay_rvalues
{
    using type = std::decay_t<T>;
};
template <typename T>
struct decay_rvalue<T&>
{
    using type=T&;
};

#define auto_cref(var,...) \
typename decay_rvalue<decltype((__VA_ARGS__))>::type var = ( __VA_ARGS__)'

```

if we add parentheses it's bad, it will always return a reference.

```cpp
decltype(auto) foo()
{
    auto_cref (a, some_a()); // a = some(); with type deduced
    return a; //if we have parentheses, things will be different, it will be a reference.
}
```

theres a debate about whether the macro should return const or not (if not, it can get optimized in NRVO).

```cpp
struct A;
struct B {
    A m_a;
    A const & GetA() const
    {
        return m_a;
    }
}
auto_cref(a1, B().m_a); // B() is rvalue, so it's members are also rvalues;
auto_cref(a2, B().GetA()); // we have a const reference as the return type, so we get a dangling reference const A &;
```

now the problem is that our 'auto_cref' binds to everything, but should rvalues be converted to values?

```cpp
struct A;
A const & L(); //lvalue
A const && R(); //rvalue

decltype(false? L(): L()); // A const &
decltype(false? R(): R());// A const &&
decltype(false? R(): L());// A const, not reference. forces a copy.
```

c++20 has a new trait _common_reference_t_. which was invented for c++20 ranges,

```cpp
std::common_reference_t<A const &, A const &>; //A const &
std::common_reference_t<A const &&, A const &&>; // A const &&
std::common_reference_t<A const &, A const &&>; //A const &. lvalue reference
std::common_reference_t<A const, A const &>; //A. a value
```

so, std::common_reference embraces rvalue amnesia.

### Promises of defences

| Mutability | short Lifetime               | long Lifetime |
| ---------- | ---------------------------- | ------------- |
| immutable  | const &&                     | const &       |
| mutable    | && (can scavenge, move from) | &             |

currently, c++20 reference binding strengths lifetime promise.
from short to long, and from mutable to immutable.

what if we could go the reverse?

> - Allow binding only if promises get weaker
>   - less lifetime
>   - less mutability
>   - less 'scavenge-ability'
>
> * only lvalues should bind to _const &_
> * anything may bind to _const &&_

but we can't allow going from lvalue to rvalue.

### Ideas to Fix the Issue

some things that must hold true before any changes.

where are references used?

> - local/global variable declarations
> - structured binding
> - function/lambda parameter lists
> - members (initialized in PODs)
> - members (initialized in constructors)
> - lambda captures

how it would look with a pragma change. we would need a feature test macro, replace const & parameters with const &&. we will need to change std::common_reference.

</details>

## C++11/14 at Scale: What Have We Learned? - John Lakos & Vittorio Romeo

<details>
<summary>
How we teach C++ versions, which features are suitable for which skill level? what should be taught and when.
</summary>

[C++11/14 at Scale: What Have We Learned?](https://youtu.be/E3JG2Ijjei4)

They're are publishing a book later this year: **Embracing Modern C++ Safety**

> - Why are we talking about C++11/14 in 2021?
> - How C++11/14 an surprise you today
> - C++ at scale
> - "safety" of a feature
> - case study: extended _friend_ declrations

### Why are we talking about C++11/14 in 2021?

adoption rates of C++ standards, some projects are still lagging behind and haven't adapted the newer standads yet, even in 2021 there are places where C++11/14 are just being adopted.

> - There are great learning resources
>   - But most teach "the features" rather than "the experience"
>   - What looks good on paper might not work in the "real world"

not just what the new features are, learn when and how to use them, how they operate inside the bigger context.

### How C++11/14 an Surprise You Today

> Q: "What is the smallest change to the core language you can think of in C++11?"

c++11 changed how double ">" behaved. before c++11 ">>" was parsed as a right shift, so a space was needed to make this recognizable as closing a nested template. in c++11 things were changed and ">>" was now somthing else. this means that a valid c++03 code is invalid in c++11.

in this example, c++03 will see `256 >> 4`, but c++11 will reject this code.

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<256 >> 4> smallBuffer;
```

to fix this issue, we simply wrap the right shift expression in parentheses

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<(256 >> 4)> smallBuffer;
```

in this example, c++03 returns 100, while c++11 returns 0; the compiler gives a warning.

I think that c++03 treats this as a sequence of enums and nested enums (we can change the final 'a' to 'c' and get 102, but not to 'b'). and c++11 treats this as some comparions thing.

```cpp
enum Outer{a=1,b=2,c=3};
template <typename>
struct S {enum Inner {a=100, c=102};};
template <int>
struct G{typedef int b;};
int main()
{

    return S<G< 0 >> ::c>::b>::a;
}
```

other stuff: every one of those can have a dark side.

> - Attributes that can make you code ill-formed NDR.
> - 'extern templates' not improving compilations time or code size at all?
> - Destruction order UB with meyers singletons.
> - Encoding of white space withing raw string literals.

> "NDR - No Diagnostic Required"

### Modern C++ at Scale

how do we teach modern c++? what to prioritize, what kind of approach? how do we integrate the new features into the company style guide? what if we have a tool chain for the style guide? how do we communicate these changes to other teams?

### "Safety" of a Feature

> - Every features of c++ is "safe" when used correctly.
> - But what is the likelihood that it is used correctly?
> - Does the feature have any "attractive nuisance"? (does it invite misuse?)
> - What are the advantages of using a feature compated to its risks?
> - Is it worth teaching to a new hire? to an expereinced hire?

from the book:

> "The degree of safety of a given feature is the the relative likelihood that the widespread use of that feature will have positive impact and no adverse effect on a large software company's codebase."

how likely is teaching and implementing a feature is to go smoothly, be used correctly and give good results, as opposed to being hard to teach/understand, prone to create oppertunities for bugs, hard to maintain by inexperienced workers, or have small scale effects on performance.

three categories of safety:

> - **Safe**
>   - Adds considerable value, easy to use, hard to misuse.
>   - Ubiquitous adoption of such features is productive.
> - **Conditionally Safe**
>   - Adds considerable value, but prone to misuse.
>   - Require in-depth training and additional care.
> - **Unsafe**
>   - Provide value only in the hands of an 'expert', prone to misuse.
>   - Wouldn't teach these as part of genereal c++11/14 course.
>   - Require explicit training on their use cases and pitfalls.

the 'override' keyword is a 'safe' feature. it prevents bugs, makes code self-explantory, and has no real technical downsides.
the one problem that can happen is that people overrely on it, and people except this feature as the norm, and forget that this is just a bonus, you can still have overriding methods without this keyword.
(we can use compiler flags '-Winconsistent-missing-override' and '-Wsuggest-override', but they aren't perfect).

```cpp
class MockConnection : public Connection
{
    void connect(IPV4Address ip) override;
};
```

the 'auto' keyword is _conditionally safe_, it can be great, but can also cause readability problems, and can introduce bugs. _Range based for loops_ are great, but they also have the possibility of bugs, and therefore are marked _conditionally safe_.

in this example, we are actually ok, because we return the vector by value and get lifetime extentsion. this is not true once we decide to be smart and return the vector by reference. now we don't have lifetime extension.

```cpp
class TriggerGetter
{
std::vector<Combo> getCombos() const; //no Issues
const std::vector<Combo> & getCombosRef() const; //oops
};

for (Combo& c : keyBoardTriggerGetters[bindID]().getCombos()) //return by value
{
    //..
}
for (Combo& c : keyBoardTriggerGetters[bindID]().getCombosRef()) //return by const reference
{
    //..
}
```

this is something we overcome in c++20 with init statements, but we must be aware of this issue. it's not an entirely safe action.

_decltype(auto)_ is a strong feature, but it's often misunderstood, misused, a requires training to use correctly, it should be defined as _unsafe_, and only be used with carefull consideration in the codebase. it allows us to deduce the return type from the expression, doesn't strip away qualifiers, returns value or reference objects,and doesn't change anything.

example: higher-order functions.

```cpp
template <typename F>
decltype(auto) logAndCall(F&& f)
{
    log("invoking function ", nameOf<F()>);
    return std::forward<F>(f)();
}
```

but if we teach '_auto_', '_decltype_' and '_decltype(auto)_' together, we push people towards overusing '_decltype(auto)_'.

> - some misconceptions:
> - "If _decltype(auto)_ does everything _auto_ does and more, why not use it all the time?"
> - "If _decltype(auto)_ is more flexible, wht no use it when I'm not sure when to choose between _auto_ and _auto&_?"

> understanding _decltype(auto)_ requires:
>
> - Having a solid grasp on type inference and value categories.
> - Being somewhat experienced with using _auto_ and _decltype_.
> - Having some metaprogrammin expereice.

if you have just learned about _auto_ and _decltype_, you probably aren't in the right level to use _decltype(auto)_ yet.

plus, _decltype(auto)_ has some issues with parentheses surronding it, it's not an easy thing to understand, it can effect SFINAE behavior, so it's not the allways the best tool for the job.

safe features: attributes (most of them), _nullptr_,_static_assert_, digit seperators.
conditionally safe: _auto_, _constexpr_, _rvalue_ references
unsafe: _\[\[carries_dependency]]_, _final_, inline namepaces.

when we teach a new version of c++, we should:

> - Teach _safe_ features early and quickly
>   - Most of them are quality-of-life improvements or hard to misuse.
>   - Trust the student
> - Teach _conditionally safe_ features by building on top of _safe_ knowledge
>   - They require more time and examples.
>   - Show how the can backfire.
>   - Have exercises that make student question whether to use a feature or not.
> - Leave a subset of of _unsafe_ features for self-contained CE courses
>   - E.g. "Library API and ABI version with the 'inline namespaces'"

### Case study: Extended _friend_ Declrations

> - Prior to C++11, _friend_ declareations require an 'elaborated type specifier'.
>   - _elaborated type specifier_: Syntitical element having the form of \<class|struct|union> \<identifier>
> - This restriction prevents other entities to be designated as friends.
>   - E.g. type aliases, template parameters.
> - A surprsing behavior with namespaces.
>   - it wasn't possible to refer to a entity in the global namespace, a new entity was being declared instead.

```cpp
//C++03 friend
struct S;
struct Example
{
    friend class S; //ok
    friend class NonExistent; //ok, even it this class doesn't exist.
};

using WindowManger = UnixWindowManager;

template <typename T>
struct Example2
{
    friend class WindowManger; //error! type alias
    friend class T; //error! template parameter
};

struct SA; //this SA is in the global namespace
namespace ns
{
    class X3
    {
        friend struct SA; // ok, declares a new ns::SA class instead of refereing to the global ::SA
    };
}
```

> C++ 11 extended 'friend' declarations lift all the aforementioned limitations. and fixes the weird behavior of creating types. we don't need the class|struct|union specifier anymore.

```cpp
//C++11 friend

Struct S;
typedef S Salias;
using Salias2 = S;

namespace ns
{
    template <typename T>
    struct X4
    {
        friend T; //ok
        friend S; //ok, refers to global ::S
        friend SAlias; //ok, also refers to global ::S
        friend Salias2; //ok, also refers to global ::S
        friend decltype(0); //ok, same as 'friend int'
        friend C; //error! 'C' does not name a type.
    }
}
```

> so why is this feature categorized as _unsafe_?
>
> - It is rarely useful in practice, like c++03 _friend_
> - Promotes _long-distance friendship_!
>
> When a type 'X' befriends type 'Y' which lives in a separate component...
>
> - 'X' and 'Y' cannot be thoroughly tested independently anymore.
> - Physical coupling occurs between 'X' and 'Y' components,
> - Possible physical design cycles can happen.

if we have too many friends, it might be a symptom of a design problem, having friends from diffrent namespaces means more coupleing and less modularity. but even if it's _unsafe_, it does have it's benefits, like helping us spot typos when declaring friends.

other intersting points: type alias customization points, PassKey idiom...

and we will focus on CRTP - curiously recursive template pattern.

base knows who it derives from, thanks to T. usefull to implement _mixins_ and factor out copy-pasted code.

```cpp
template <typename T>
class Base
{

};

class Derived : public Base<Derived>
{

};
```

example use case, having a counter for classes creations.

```cpp
//header file
class A {
    static int s_count; //decleration
    public:
    static int count() {return s_count;}
    A(){++_count;}
    A(const A&) {++s_count;}
    A(A&&) {++s_count;}
    ~A() {--s_count;}
};

//defintion file
int A::s_count;
```

we can factor out the counter behavior, using the protected access modifier. (it's a mixin, whatever that means).

```cpp
template <typename T>
class InstanceCounter
{
    protected:
    static int s_count; //declaration
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)
```

we can then use in other classes

```cpp
struct A :InstanceCounter<A>
{
    A() {++s_count};
    //also add this for the destructor
};

struct B : InstanceCounter<A> //oops, made a typo! we will use the same counter.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //oops, we messed with the entire tree!
}
```

actually, this is something we could use the friend declerations! we move from 'protected' to 'private', and make T a friend class of the counter. now only the class that declared the counter can access it, not others classes and not derived.

```cpp
template <typename T>
class InstanceCounter
{
    private:
    static int s_count; //declaration, private
    friend T; //only T can access us.
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)

struct B : InstanceCounter<A> //error, s_count is private within this context.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //error, s_count is private within this context
}

```

the crtp allows us avoid boiler plate code, this is also an example of using inheritance without virtual functions. this is also a case where we don't want to use 'final'.

### when to use 'final'

do we really know a class shouldn't be inherited from? are we sure.

it's okay if we have a class that is supposed to behave like a primitive.
but EBO (empty base optimization) doesn't play nice with 'final'.

### Conclusion

> - The "Human cost" of a feature is not easy to quantify.
> - Categorizing features by "safety" helps with devising learning paths.
> - All features have good use cases and nasty pitfalls.

the book will be out in the future, check [this page](https://vittorioromeo.info/emcpps.html)

</details>

## The Performance Price of Dynamic Memory in C++ - Ivica Bogosavljevic

<details>
<summary>
The performance costs of using dynamic memory, in terms of allocating and deallocating memory from the system and in terms of accessing the memory and data locality. suggestions to improve our software desgin to make it cache aware.
</summary>

[The Performance Price of Dynamic Memory in C++](https://youtu.be/LC4jOs6z-ZI), [slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Price-of-Dynamic-Memory-CNow2021.pdf)

### Introduction

two types of programs when it comes to memory allocations:

- allocate all memory in few large blocks of memory, like how arrays, vector and matrices have continuous memory blocks. this is used in image, audio and video processing, where the algorithm requires large buffers.
- allocate many blocks of data on demand during runtime. fast random access, pointer types.

there is a performance cost to allocating and deallocting memory during runtime. malloc(new) and free(delete) can be the bottleneck, this will show up in the profiler. but this can also lead to increased cache misses, which require specialized tools to see this.

### Performance of Memory Allocation

the system allocator, not the same as the stl allocator. they allocate large memory to the program, and then give the application parts of the block when a _malloc_ is perfromed. allocation algorithm find the free chunk inside the block. the more allocations are done, the more time it takes to find a correct sized chunk of memory. so performance degrades over time.

there are three possible reasons to why memory allocation and deallocation are slow:

> - Your program is allocating and deallocating a lot of memory, especially small memory chunks.
> - Memory fragmentation.
> - Your program is using an inefficient implementation of malloc (new) and free (delete).

if we can fix one of those, we can get better performance. there are some guidelines:

The worst offenders are vector of pointers, they require lots of allocations/deallocations, and they have poor cache locality. an alternative is to use a separate vector per type (struct of arrays rather than array of struct approach), or use _std::variant_ instead of pointers , which will force locality and reduce memory fragmentation. both suggestions reduce the number of calls to the allocator.

Another solution is to use the [STL allocator](https://en.cppreference.com/w/cpp/memory/allocator) for our data structures, each data structure gets it's own dedicated block of memory, so fragmentation is reduced.

in this custom allocator example, we allocate the memory using mmap (so it's only POSIX compatible), which is deallocated in the destructor. the system memory allocation is done only when the allocator is created or destroyed, not for each time.

```cpp
template <typename _Tp>
class zone_allocator
{
private:
    _Tp* my_memory;
    int free_block_index;
    static constexpr int mem_size = 1000*1024*1024; //1GB?
public:
    zone_allocator()
    {
        my_memory = reinterpret_cast<_Tp*>(mmap(0, mem_size, PROT_READ|PROT_WRITE,
        MAP_PRIVATE|MAP_ANONYMOUS, -1, 0));
        free_block_index = 0;
    }
    ~zone_allocator()
    {
        munmap(my_memory, mem_size);
    }
    //...
    pointer allocate(size_type __n, const void* = 0)
    {
        pointer result = &my_memory[free_block_index];
        free_block_index += __n;
        return result;
    }
    void deallocate(pointer __p, size_type __n)
    {
        // We deallocate everything when destroyed, not for each deallocation
    }
    //...
};

std::map<int, my_class, std::less<int>, zone_allocator<std::pair<const int, my_class>>>  some_map;
```

something similar is part of the standard since c++17 _std::pmr::polymorphic_allocator_.

Another issue is allocating objects for communication:

a common pattern is sending messages between threads as objects, the sends allocates the object and sends it, while the receiver thread reads and deallocates them.\
Obviously, this causes a lot of allocations and deallocations, and additionally, system allocators don't play nice when they need do allocate memory in one thread and deallocate it in another.\
a solution for this issue is to avoid releasing memory back to the allocator, and to cache it instead:

```cpp
T* memory = allocator.get_memory_chunk(); //uninitlazed memory chunk
new (memory) T (args...); // in-place constructor
//...
memory->~T(); //destructor
allocator.release_memory_chunk();
```

other solutions:

- Preallocate all the needed memory upfront (like in embedded systems), if data structure has a limit, request that limit upfront.
- Restart the program (when possible), don't forget to save the state!
- Use _small buffer optimizations_ for small enough data structure, such as std::string.
- Use special system allocators that promise low fragmentation.

there are off-the-shelf allocators. we need to consider:

- allocation speed
- memory consumption
- memory fragmentation
- data locality

we can get allocators other than the standard one, most of the big companies have an allocator they use (microsot, google, facebook). we can change allcators without recompiling by switching the environment variable.

### Performance of Memory Access

if we allocate and deallocate a lot of memory, we need to think about the layout of the memory, how is it stored in memory? We also need to consider how we access the memory itself, sequential access is better than random access in terms of speed.\
the allocator operates under the abstraction of the underlying hardware, but if we break this abstraction and make it aware of the hardware, the allocator algorithm can give us better results. Memory speed is a bottleneck on modern systems. Waiting for memory fetch (load into register) is about 200-300 cycles of cpu. this is the _cache memory_ (located on the cpu) cost is much faster (3-15 cycles). also called _dataset_.

stages of memory access:

1. check if the data is already present in the cache memory.
2. if so, load it into the registers, otherwise, fetch it from the main memory.
3. when the data in the dataset isn't accessed for some time, the modified results are written back to the main memory and removed to make space for new data, this is called _eviction_.

example with hash maps, the larger the data, the smaller the chance we use the same data again before it's evicted.

There is a way to pre-fetch data, this is done if our program has a predictable access pattern (usually means iterating linearly over a vector of objects) then the cache memory preFetcher can figure that out an load the data from the memory before it's needed by the cpu. this means major improvements in speed.\
the prefetcher works with sequential and strided memory access, but is powerless with random memory access.

Cache memories are divided into _cache lines_ (usually 64 bytes), each line corresponds to a block of the same size in memory. Taking one byte from that memory means the whole line will be fetched. So access to other data on that line will also be fast.\
We take advantage of this by organizing our data so that data that is close in usage is close together in layout. we also should use as much data that we can once we load a cache line.

> "It is a sin to load data into the cache line and then not use it."

this explains why the stride example is worse than sequential access, and why large strides are worse than small ones.

- CPU always works with simple types: char, int, double, float, etc.
- From the performance point of view, dependence between the memory
  access pattern and the access speed looks like this:

  - **Sequential access**: you are accessing neighboring simple types - best performance.
    ```cpp
    vector<int> a;
    int sum = 0;
    for (i = 0; i < a.size(); i++)
    {
        sum += a[i];
    }
    ```
  - **Strided access**: you are accessing simple type in a vector of class instances - bad for performance, the bigger the class, the worse the performance.
    ```cpp
    vector<rectangle> a;
    int sum = 0;
    for (i = 0; i < a.size(); i++)
    {
       sum += a[i].visible; //each rectangle is located at a difference from one another, so we don't use all of our cacheline
    }
    ```
  - **Random access**: you are randomly accessing objects in memory (std::set, std::map, std::list), or accessing an object through a pointer

    ```cpp
    set<rectangle> a;
    int sum = 0;
    for (auto& r: a) //set is not continuous in memory
    {
        sum += r.visible;
    }

    class car
    {
        driver* m_driver;
    };

    if (my_car.m_driver->experience() > 5) //access through pointer
    {
        //..
    }
    ```

experimenting with class size and member layout:

the two methods differ by how many members of the objects they access. when we change the class size (by increasing the padding) the performance becomes worse, but for a fixed class size of 248, changing the padding between the _m_visible_ and the points members effects the version that checks for visibility (calculate visible), but not the method which doesn't check for visibility.

```cpp
template <int pad1_size, int pad2_size>
class rectangle
{
    bool isVisible();
    int surface();
private:
    bool m_visible;
    int m_padding1[pad1_size]; //padding, unused memory
    point m_p1;
    point m_p2;
    int m_padding2[pad2_size]; //padding, unused memory
};

template <typename R>
int calculate_surface_visible(std::vector<R>&rectangles)
{
    int sum = 0;
    for (int i = 0; i < rectangles.size(); i++)
    {
        if (rectangles[i].is_visible())
        {
            sum += rectangles[i].surface();
        }
    }
    return sum;
}

template <typename R>
int calculate_surface(std::vector<R>&rectangles)
{
    int sum = 0;
    for (int i = 0; i < rectangles.size(); i++)
    {
        sum += rectangles[i].surface();
    }
    return sum;
}
```

### Principles of Cache-Aware Software

this means that making data access predictable helps us.\
Removing branches makes access predictable. We also get better performance with vectors of objects rather than vector of pointers, random access memory containers aren't as good for cache locality.
Pointers are bad because of the dereferencing (to other heap allocated classes). instead of linked list we can use an array based alternative (colony, bucket array, gap-buffer).

the optional layout of memory for array pointers is having the pointers point to object ordered in the same order, pointer1 to obj1, p2 to o2, etc... but this order isn't likely at all.

#### arrays of values are better than arrays of pointers

- All memory allocated in a single block.
- Sequential access to objects translates into sequential access to memory addresses.
- No calls to malloc/free.
- No virtual dispatching.
- Enables small function inlining because typ is known at compile time.
- **Downside** - no polymorphism

the whole block of 64 bytes is loaded into memory, we shouldn't let data into cache line go wasted. we would like classes to be small, we would prefer them separate from the large class in the vector.

because of how memory is mapped, we should have members that are used together packed together.

#### Paradigms

> object oriented paradigm can be inefficient from the performance perspective.
>
> - Containers of pointers.
> - Contianers of objects of different types.
> - Large classes
> - Classes that have member that point to other heap allocated classes.

in game development, they use Entity-Component-System (ECS) paradigm.

> - Get rid of large classes.
> - Get rid of inheritance.
>   - Instead of inheritance An entity consists of components.
> - Components are processed independent of the entity they belong to.
> - Entity can change its "type" at runtime.

more principles of cache-aware software design.

- If we need to use trees, we should prefer n-ary trees over binary trees, as they get better cache line effectiveness.
- We should store pointers in a hash map, we should store the whole objects (to reduce cache misses).
- Hash maps with **open addressing** perfrom better than separate chainning maps (all the data is stored in the table itself) in cases of collisions, but they have downsides when the load factor is high.

#### Binary tree example

how to optimally represent the tree in memory? there are three options for the layout:

- Breadth First Search (BFS)
- Depth First Search (DFS)
- Random order

the DFS is more efficient than the others, but the most efficient is called _Van Emde Boas layout_. which lays out nodes in triplets, the node and it's children together.

we gain performance if we allocate a dedicated block of memory for the data structure (using a custom allocator). if we keep the block of memory as compact as possible, and if we take advantage of cache line organization by storing adjacent nodes in the tree in adjacent memory locations.\
we can also play with the struct to get better performance, like decreasing the size of the struct or playing with the pointers data.\
Modifying the tree hurts the memory layout, we can recreate the tree after some time to make it optimal again, we can perform defragmentaiton on it or keep the 'removed nodes' without deleting them in case they need to be reused.

#### Hash map example

elements are number (count) and pointer to array. we usually store zero, one or two elements, so it's better to optimize the struct to either contain the value itself (for one value, the common case) or the pointer to array (for the rare case of many values)

don't let data that we will reuse get evicted from the cache.
don't reiterate the same collection two times (finding min and max)

### When Not to Optimize

Most optimizations only make sense for large data sets. we can classify the sizes by how they fit into cache levels. the small size fits into L1 cache, about 16Kb - 32Kb in size. large datasets are those which are larger than the last layer (LL), which is usually a few megabytes.\
Data structures with short life cycles don't benefit much from optimizations, as the creation costs of an optimal layout are usually larger than the savings for an object that won't be used in the future.

</details>

## The Complete Guide to 'return x;' - Arthur O'Dwyer

<details>
<summary>
optimizations for returning values - getting moves rather than copies and dropping those calls alltogether. there are still probelms to solve!
</summary>

[The Complete Guide to `return x;`](https://youtu.be/OGKAJD7bmr8),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/2021-05-04_-The-Complete-Guide-to-return-x-1.pdf)

### The "return slot", **NRVO**, C++17 **"deferred materialization"**

in x86 machines, the return value usually goes to the %eax register.

copy semantics, depending on size, if the return type size is too large, the caller should also pass an empty space (return slot).

```cpp
struct S {int m[3]}; //too large to copy in register

S f()
{
    S i = S {{1,2,3}};
    printf("%p\n",&i);
    return i;
}

S test()
{
    S j =f();
    printf("%p\n",&j);
    return j;
}
```

the test can allocate the local variable j on the return slot, the optimization is transparent and invisible.\
C++98 even had special cases where we could do this "copy ellision" on types with user created explicit constructors.\
In c++14 f() was "eagerly evaluated" into a temporary object and then moved/copied.\
In c++17, this was expanded into "deferred prvalue materialization", a result object.

function f controls the allocation of element i, so f can allocate the struct immediately inside the return slot.

NRVO - named return value optimization.

conditions to NRVO of variable _i_ from function _f_.

> - There must be a return slot. Trivial types can just be returned in registers, But types with non-trivial SMFs (special member function) will always be returned via return slot.
> - The allocation of “return variable” _i_ must be under _f_’s control, Otherwise _f_ can’t allocate _i_ into the return slot!
> - _i_ must have the exact same type (modulo cv-qualification) as _f_’s return slot, Otherwise _i_ won’t fit into the return slot!
> - One mental — not physical — caveat: The return’s operand must be exactly
>   (possibly parenthesized) id-expression, such as _i_. Nothing more complicated. Otherwise things could get very confusing for the human programmer!

examples of cases where NRVO doesn't happen

```cpp
struct Trivial { int m; };
struct S { S(); ~S(); };
struct D : public S { int n; };
Trivial f() { Trivial x; return x; } // no return slot. trivial type. fits in a register.
S x;
S g1() { return x; } // g doesn’t control allocation of x
S g2() { static S x; return x; } // same deal
S g3(S x) { return x; } // same deal (params are caller-allocated!)
S h() { D x; return x; } // D is too big for the return slot
```

in c++11 we added move semantics to the mix, this doesn't play well with NRVVO.

```cpp
template <typename T>
unique_ptr<T> f() {
 unique_ptr<T> x =T(1,2,3); //... assume this works
 return x;
}
```

x is an lvalue, not an rvalue, _std::unique_ptr_ is a move only type, we can't construct it from an lvalue, only an rvalue. if we wrote _std::move(x)_, then our NRVO won't work, because it only works on simple _return x_ statements!

the text it self is:

> \[...] overload resolution to select the constructor for the copy is first performed as if the object were designated by an rvalue. If overload resolution fails, or if the type of the first parameter of the selected constructor is not an rvalue reference to the object’s type (possibly cv-qualified), overload resolution is performed again, considering the object as an lvalue...

### C++11 Implicit Move

in c++11, the solution is implicit move, we have an overload resultion, if we see a valid move constructor, we use it,even if it's an lvalue. if this is an well formed expression, this will still work with RNVO. if overload resultion fails, then regular rules apply.

"first try as if it was an rvalue"

(example with the deprecated _std::auto_ptr_)

> \[...] overload resolution to select the constructor for the copy is first performed as if the object were designated by an rvalue. **If overload resolution fails, or if the type of the first parameter of the selected constructor is not an rvalue reference** to the object’s type (possibly cv-qualified), overload resolution is performed again, considering the object as an lvalue...

after c++11 was released, it was modified to allow implicit moves even to cases where _copy elision_ wasn't considered. such as when trying to return a move only function parameter, or when the return type isn't the same, but an move constructor exists.

```cpp
unique_ptr<Base> g3(unique_ptr<Base> x) {
 return x; // OK, implicit move! return a move only parameter
}
unique_ptr<Base> h2() {
 unique_ptr<Derived> x = Derived();
 return x; // OK, implicit move! there is a well defined move constructor
}
```

### Problems in C++11, Solutions in C++20

things didn't change that much in c++14 and c++17. but there were still problematic cases.

this example is well formed, the return type is constructed with the copy constructor `(Const base &)`, but we would prefer it to use the move constructor `(Base &&)`.

```cpp
Base h3() {
 Derived x = Derived();
 return x; // Ugh, copy! slicing!
}
```

this is because the standard require the move constructor to take the object type, rather than something that can be considered as such.

> \[...] overload resolution to select the constructor for the copy is first performed as if the object were designated by an rvalue. If overload resolution fails, or if the type of the first parameter of the selected **constructor** is not an rvalue reference **to the object’s type (possibly cv-qualified)**, overload resolution is performed again, considering the object as an lvalue...

other examples of failing to call implicit move.

```cpp
struct Source {
    Source(Source&&);
    Source(const Source&);
    };
struct Sink {
    Sink(Source);
    Sink(unique_ptr<int>);
    };

Sink f() { Source x; return x; } // C++17 calls Source(const Source&), then Sink(Source)
Sink g() { unique_ptr<int> p; return p; } // C++17: ill-formed, not a moveable type
```

we don't want to write _std::move_, as it hampers optimizations.

we also fail to move when we have a conversion operator and not a constructor. we don't consider them as part of the overload resultion.

```cpp
struct To {};
struct From {
    operator To() &&;
    operator To() const&;
    };

To f()
{
   From x;
   return x; // C++17 calls From::operator To() const&
}
```

because of those, the c++20 standard was expanded to include "More implicit move".

> \[...] overload resolution to select the constructor for the copy is first performed as if the object were designated by an rvalue. If overload resolution fails, ~~or if the type of the first parameter of the selected constructor is not an rvalue reference to the object’s type (possibly cv-qualified),~~ overload resolution is performed again, considering the object as an lvalue...

and now things can work better, overload resolution is expanded. this helped with other cases like throw and co_return (coroutine), and function parameters.

actually, even before c++20 changed the standard, compiler vendors had optimizations for implicit move, which were more efficient. even if not 'up to standard'. updating the standard closed the gap between formal specification and performance.

c++20 had other big changes, paper P0527 "Implicitly move from rvalue references in return statement" expanded the possibilities of implicit moving, allowing for returnning named entities (lvalues) as implicit moves.

"perfect backwarding", maintain the exact type. allow for more moves and less copying. things aren't perfect yet. returning decaltype(auto) required a _std::move_, the wording didn't allow for implicit move of returnning references.

so function that returns an object can be implicitly moved constructed from a rvalue (universal?) reference, but we couldn't properly implicitly return that reference.

```cpp
MoveOnly one(MoveOnly&& rr)
{
 return rr; // OK, move-constructs from rr (in C++20)
}
MoveOnly&& two(MoveOnly&& rr)
{
 return rr; // ill-formed, rr is an lvalue
}
```

we can add _std::forward_, but it will stop NRVO from happenning.

### The _reference_wrapper_ Saga, Pretty Tables of Vendor Divergence

the reference wrapper from c++98. this code compiles but creates a dangaling reference to a garbage memory.

```cpp
reference_wrapper<int> f() {
    int x = 42;
    return x;
}
template<class T>
struct reference_wrapper {
reference_wrapper(T&);
reference_wrapper(T&&) = delete; //c++11
};
```

in c++11 the rules for implicit moves were introduced, so the candidate of the move constructor was found, but it was also deleted, which made the entire thing ill-formed (which was good). but in some context we want to ignore the deleted functions in overload resolutions.

this was mended by introducting a SFINAE test rather than deleting the overload, which is still the current situation.

the question still remains, what does matching a deleted function mean in this context. how is ambiguity in overload resolution considered?

```cpp
struct RefWrap { RefWrap(T&); RefWrap(T&&) = delete; };

RefWrap f() { T x; return x; } // ill-formed since CWG1579 (C++11)

struct Left {};
struct Right {};
struct Both: Left, Right {};
struct Ambiguous {
    Ambiguous(Left&&);
    Ambiguous(Right&&);
    Ambiguous(Both&);
};

Ambiguous f() { Both x; return x; } // ill-formed since P1155 (C++20)
```

the implicit move rules only apply for objects, not references, not pointers, etc..
vendors still have differences between them and the standard.

### Quick Sidebar on Coroutines and Related Topics

c++20 expanded implicit moves to handle co-routines returns, but the wording doesn't necessarily limit the return type to be objects, so technically, behavior is different again.

co_yield also might seem to require moves, but it doesn't.

```cpp
template<class T>
struct generator {
    struct promise_type {
        std::suspend_always yield_value(const T&);
        std::suspend_always yield_value(T&&);
    };
};

generator<std::string> g() {
    for (int i=0; i < 100; ++i)
    {
        std::string x = std::to_string(i);
        co_yield x; // Hmm... Couldn’t we move-from x here?
    }
}

```

this doesn't work, because the frame of the coroutine doesn't go away, we might want to use the value again after we 'yielded' it.

returning a member variable is like that, we shouldn't move it from return statement in member functions, lambdas are also like that, if we can't control the lifetime/allocation of a variable, we can't move from them. we would have to explicitly move with stdmove.

strutted bindings are also not playing well with implicit moves, and nor are references.

### P2266 proposed for C++23

this is all confusing, here is what arthur suggest - paper P2266 "Simpler Implicit move".

remove the fallbacks that were created to allow using auto_ptr- like types, find a unifrom definition, do one overload resolution.

might break compatibility with auto_ptr, might change some bad code and allow other bad code.

some stuff with decaltype(auto)

**P2266 and decltype(expr):**

| Return type                                     | C++14,17, 20     | P2266(C++23)    |
| ----------------------------------------------- | ---------------- | --------------- |
| auto a(T x) -> decltype(x) { return x; }        | T                | T               |
| auto b(T x) -> decltype((x)) { return (x); }    | T&               | T& (ill-formed) |
| auto c(T x) -> decltype(auto) { return x; }     | T                | T               |
| auto d(T x) -> decltype(auto) { return (x); }   | T&               | T&&             |
| auto e(T&& x) -> decltype(x) { return x; }      | T&& (ill-formed) | T&&             |
| auto f(T&& x) -> decltype((x)) { return (x); }  | T&               | T& (ill-formed) |
| auto g(T&& x) -> decltype(auto) { return x; }   | T&& (ill-formed) | T&&             |
| auto h(T&& x) -> decltype(auto) { return (x); } | T&               | T&&             |

</details>

## The Worst Best Practices - Jason Turner

<details>
<summary>
Thoughts about being a C++ content creator, as told by jason after doing this for the past 15 years.
</summary>

[The Worst Best Practices](https://youtu.be/KeI03tv9EKE),
[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Jason-Turner-CNow-2021-The-Worst-Best-Practices.pdf), [github](https://github.com/cpp-best-practices/cppbestpractices)

jason retelling some old blog posts of his and talking about comments that he got, talking about resources he read and stuff that he wrote, talks he gave over the years.

talking about weird comments in youtube he got and saw others getting. how he got into making his own c++ content. what he thinks about publishing in leanpub (pros and cons). how the pricing and sales work for the book. setting the expectation of what the readers will get from the book.

critisms he got for the book, what people disagreed with in the book.

</details>
