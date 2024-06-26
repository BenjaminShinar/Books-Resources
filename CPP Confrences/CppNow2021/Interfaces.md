<!--
ignore these words in spell check for this file
// cSpell:ignore Niebloids Hollman Niebler libnv Hyrum Hohfeld
-->

[Main](README.md)

Interfaces

## When Should You Give Two Things the Same Name? - Arthur O'Dwyer

<details>
<summary>
Do we need to call methods by the same name? when and why? 
</summary>

[When Should You Give Two Things the Same Name?](https://youtu.be/OQgFEkgKx2s)

> - When do ue us classical inheritance.
> - Idiosyncratic philosophical digressions.
> - Copious anecdotes from the STL.
> - Kind of a major rabbit-hole about constructors.
> - Mental templates, macros and polyfills
> - Bonus mantras and takeaways.

### The role of (OO) Inheritance

What do we expects from inheritance?/
We expect virtual functions and somewhere that they're used: polymorphic methods, deletion of pointers through the virtual destructor...

```cpp
struct Animal{
    virtual void feed();
    virtual ~Animal()=default
};
struct Cat: public Animal{
    // what do we expect here?
    //probably an override of feed()
};
```

but if we see code with value-semantic and none polymorphic code, we will be confused.

```cpp
Cat acquirePet();
void foo(Cat & current)
{
    auto newPet = acquirePet();
    std::swap(current, newPet);
}
```

the two approches can be combined (public inheritance without polymorphism)

> - EBO - [Empty Base Optimization](https://en.cppreference.com/w/cpp/language/ebo)
> - CRTP - [Curiously Recurring Template Pattern]https://en.cppreference.com/w/cpp/language/crtp
> - TagDispatch

but they are more of a corner cases, not the intended usage of inheritance.

```cpp
template <class Allocator>
struct CatEBO : Allocator{

};

struct CatCRTP : CanFightWith<CatCRTP>
{

};

struct CatTagDispatch: AnyAnimalTag
{

};
```

according to Liskov's substition principle:

> "**If** for each object o1 of type S **there is** an object o2 of type T **such that** all programs P defined in terms of T, the behavior of P is unchanged when o1 is substituted for o2, **then** S is a subtype of T."

and adding Occam's Razor

> "Make class S a chile of class T **if and only if** you intended to pass an objet o1 of types S as the argument to some function P defined in terms of T"
> if you don't intend to do that, there is no reason for that public inheritance relationship to exists,... and therefore that relationship should **not** exists.

Chesterton speaking against unnecessary changes and the mindset of 'modern reformers' (someone who does reforms for the sake of refroms).

> "The modern reformer says "I don't see the use of this fence; let us clear it away". The more intelligent type answers, "When you can tell me that you **do** see why it is here, **then** maybe i will allow you to destroy it"". \
> --G.K. Chesterton(1929), lightly abridged
>
> Since fences generally have reasons, tearing down fences should not be done lightly.

so if we see classical inheritance, we shouldn't change it (in a refactor) until we see why it was done this way in the first place.

Robert Frost

> "Before I build a wall I'd ask To know\
> What I was walling in, or walling out"\
> --Robert Frost, "Mending Wall" (1914)

before we put up a fence, we should know what we're doing, the reason for it, and we should document it clearly, so if we come across it in the future, we can rationally consider if it's safe to remove in the current situation.

otherwise, we might run into 'The paradox of the useless fence'./

- before we tear down a fence, we must understand why it's there.
- if there was no reason to build the fence, it will be hard to understand why it was build.
- therefore: it's harder to remove a fence that was build for no good reason than a fence that was built for good reasons with a sound rationale.

and this is a thing that we can see in many codebases. somebody writes a code that uses a technique without a good reason, and then we can't remove the code because we can't understand what they were trying to do.

in c++, when we see inheritance, we expect to see a reason why it was designed this way, and specifically, we expect to see someone using a polymorphic method. if we aren't "forced" into inheritance, we should avoid it. **Prefer composition over inheritance (Has-A is better than Is-A)**.

### Naming and STL Examples

A single name for a single entity:

> - We should use different words to refer to different ideas.
> - When refering to the same idea, we should use the same word.
> - Any given single identifier should refer unambiguously to a single entity.

two codebase, which is easier? A uses the same name (diffrent signature) for two functions. B uses different names.

```cpp
//A
bool feed(Snake& snake);
bool feed(Bear& bear);
```

or

```cpp
//B
bool feed_snake(Snake& snake);
bool feed_bear(Bear& bear);
```

using the specialized name helps us detect and trace, we can always find all the usages, jump to it, rename it, and we can always tell if which function is used. it help the computer with overload resolution, and makes it easier for the IDE.

so if we see the version A (with the overloads of the identical name), we expect that there was a reason for this, and we should actually expect a specific reason - polymorphism.

polymorphism isn't just virtual functions, there's also static polymorphism of templates.

```cpp
template <typename T>
void solve_puzzle(Animal& a)
{
    feed(a); //calling a specific overload.
}
```

both std::vector and std::list (and many other containers) use the identifier "_.push_back()_" as a method name. this same name allows us to create a template function. like the _std::back_inserter_ iterator, _std::swap_.

```cpp
template <typename T>
struct back_insert_iterator{
    //...
    // container is T*
    back_insert_iterator& operator=(const T::value_type& x)
    {
        container->push_back(x);
    }
};
```

if there was no use of polymorphism, a unique identifier would be easier to read, understand and maintain, but we get so much functionality from the STL,which makes the overloaded versions preferabl.

a counter example from the STL, _erase_ has two overloads. one identifier with two entities. Arthur says that this code doesn't facilitate any polymorphism.

```cpp
class vector
{
    using CI = const_iterator;
    iterator erase(CI first,CI last);
    iterator erase(CI where){
        return erase(where,where+1);
    }
};
```

here is an example where we trip over ourselves, we have a vector of numbers, we want to keep only the even numbers. we use the erase-remove idiom but we forget to pass the second argument to _.erase()_, so we erase only one element.

```cpp
bool isOdd(int);
std::vector<int> v= {1,2,3,4,5,6,7};
v.erase(std::remove_if(v.begin(),v.end(),isOdd)); // erase remove idiom, erase with one arguments
static_assert(std::none_of(v.begin(),v.end(), isOdd)); // this fails!
```

what we should have done is

```cpp
v.erase(std::remove_if(v.begin(),v.end(),isOdd),v.end()); // erase remove idiom, erase with two arguments
static_assert(std::none_of(v.begin(),v.end(), isOdd)); // now it's ok
```

why was the overload created? arthur says there isn't a good reason.

### An Issue with the Constructor

STL classes have too many overloads, especially std::string,

```cpp
class string {
    string(size_t n,char); // string with n times of char
    string(const char * ,size_t n); // first n chars of char*
    string(const string &,size_t pos); // copy of other string, starting at some position
    template<InputIterator It>
    string(It,it); //take two iterators
}

size_t zero =0;
auto a =std::string(zero,0); //what is called here? zero instance of character 0
auto b =std::string(0,zero); // calls the overload with the const char*, undefined behavior probably
auto c = std::string("abcd",2);  // "ab" constructor first n chars,
auto d = std::string("abcd"s,2); // "cd" constructor copy of other string from position, just because we added the string literal
```

could all these constructor be replaced with factories?

```cpp
class stringRevised {
    static stringRevised fromCopiesOfN(size_t n,char); // string with n times of char
    static stringRevised fromPtrAndLength(const char * ,size_t n); // first n chars of char*
    static stringRevised fromSuffixStartingAt(const stringRevised &,size_t pos); // copy of other string, starting at some position
    template<InputIterator It>
    static stringRevised fromRange(It,it); //take two iterators
};
size_t zero =0;
auto a =stringRevised::fromCopiesOfN(zero,0);
auto b =stringRevised::fromPtrAndLength(0,zero);
```

we couldn't do this, constructors are special.
factory functions are self documenting and easy to understand, but they don't work with the perfect-forwarding wrappers.

- _std::make_shared_, _std::make_unique_
- _emplace_back_, _optional::emplace_, _variant::emplace_

```cpp
auto a1 = std::make_shared<std::string>(zero,0);
auto a2 = std::make_shared<stringRevised>(stringRevised::fromCopiesOfN(zero,0)); //extra move operation in the good case, copy also possible.
```

constructor syntax allows us to create objects not on the stack in a comfortable way. we can actually 'new auto' (c++17) to heap allocate a factoy function p-rvalue result, gurantess heap ellision,actually good

```cpp
T t1 =T(1,2);
T* p1 =new T(2,3);
T t2 = T::fromTwoInt(3,4);
T* p2 = new auto(T::fromTwoInt(4,5)); //this works!
```

could we make a generic perfect forwarding function with factory functions?
something like this? this would work, but now instead of having a single identifier for many entities as the constructor, we simply have to choose a different name that all the classes are going to use and it won't be informative

```cpp
template <typename T, typename... Args>
auto build_shared(Args...args)
{
    T* p= new auto(T::createGenerically(args...));
    return std::make_shared<T>(p);
}
```

our fantasy: could we pass the creation format itself? pass in the factory function itself? in today's c++, this must be a concrete set (not overload set). there is one proposal for "lifting" an overload set into a concrete lambda object. a different proposal for an object that deduces types from an overload set(std::overload_set, like std::initializer_list), some sort of compiler magic.

```cpp
template <typename How, typename... Args> //the class 'How' is the problem
auto build_shared_How(How how,Args...args)
{
    auto *p= new auto(how(std::forward<Args>(args)...));
    return std::shared_ptr(p);
}
std::shared_ptr<stringRevised> sp1 = build_shared_How(stringRevised::fromCopiesOfN,0,0);
std::shared_ptr<stringRevised> sp2 = build_shared_How(stringRevised::fromPtrAndLength,0,0);

//proposal 1,
//auto sp3 = build_shared_How([]stringRevised::fromCopiesOfN,0,0);
```

### Mental Models, Macros, Polyfills

to recap, sometimes we give two entities (in different classes) the same name with the same signature, because we are going to template on the class type. this is what _std::make_shared_ does (with perfect forwarding)

```cpp
template<class Animal>
void foo(Animal & a)
{
    a.feed();
}
```

sometimes we give the same name but different signatures, because we're going to template on the argument types.

```cpp
template <class Animal, class... Foods>
void bar(Animal &a, Foods... foods)
{
    a.feed(foods...);
}
```

all STL containters provide _c.insert(pos,value)_, associative containers (like std::set) ignore the positional value. this allows us to create an _std::inserter_ with the same arguments for all containers.

```cpp
std::vector<int> data ={1,2,3}

std::vector<int> c1;
std::copy(data.begin(),data.end(),std::inserter(c1,c1.begin()));
std::set<int> c2;
std::copy(data.begin(),data.end(),std::inserter(c2,c2.begin()));
```

inserting into a set doesn't always make it bigger, if the set contains the element, it just returns it. the mental model of inserting into a set is different.
should all insert functions have the same name? why not _insertAt(pos,x)_ ,_insertNodeHandle(nh)_, _insertRange(it1,it2)_.

STL provides uniformity of containers, all containers share the same API, we can switch from _std::vector_ to _std::deque_, _std::list_ or even _std::multiset_, but does it work work the same?

no. the behavior is different. _.push_back()_ on _std::deque_ maintains the iterators, but not on a _std::vector_, _.push_back()_ invalidates the iterators (the vector might have be reallocated).

```cpp
//std::deque<int> data = {3,1,4,1,5,9,2,6,5}; //replace deque for vector
std::vector<int> data = {3,1,4,1,5,9,2,6,5};
std::sort(data.begin(), data.end());
auto [first,last]= std::equal_range(data,begin(),data.end());
data.push_back(100); // invalidates iterators in vector
data.erase(first, last); // undefined behavior in vector
for (int i: data)
{
    std::cout << i << '\n';
}
```

Can templates be mental?

> "Software engineering is programming integrated over time"\
> -- Titus Winters.

Sharing names as upgrade paths?
std::string*view and std::string share the same names for many functionalities, it was done in purpose. this was done so we could upgrade the std::string to std::string_view without issues. this was done with \_std::optional*, it has the same operators as the smart pointer classes. the reasoning was that we could replace _std::unique_ptr_ with _std::optional_, this way we reduce heap allocation and still get the 'not created' option.

reusing names can still lead to bugs. in this example both _std::optional_ and the inner type have _.reset()_ method, if we use it with the dot notation, we call the _std::optional_ method, the arrow notations is for the inner type. this would happen also with a smart pointer.

the code compiles and runs, but it doesn't do what we think it does!

```cpp
struct DataCache{
    void update(key,value);
    void reset();
};
struct Connection
{
    std::optional<DataCache> dataCache_;
    void resetCache()
    {
        if (dataCache_) //if optional value exists
        {
            dataCache_.reset(); //oops! bug! not we don't have a cache at all.
            //dataCache_->reset(); //this is what we wanted!
        }
    }
};
```

the STL and boost libraries also try to have the same names for the sake of upgrade paths.it's not a template metaprogramming, more of a **macro based static polymorphism**. the API was designed to allow this behavior. it's also called **polyfill**. the boost version is a _polyfill_ for the std version.

```cpp
#if __cplusplus >= 201703L
#include <optional>
using std::optional;
#else
#include <boost/optional/hpp>
using boost::optional;
#endif
```

we can also use this from compiler flags as platform specific polymorphism.

```cpp
namespace curses
{
    void clearScreen();
    void drawAt(int x, int y, char ch);
}

namespace conio
{
    void clearScreen();
    void drawAt(int x, int y, char ch);
}

using namespace TERMLIB; // -DTERMLIB=curses or -DTERMLIB=conio
void drawTitleBar()
{
    for (int x =0; x< 100; ++x)
    {
        drawAt(x,1,'#');// calls different function according to the TERMLIB flag.
    }
}
```

### Takeaways and Mantra

if the default parameters isn't used, don't use it. it's like an overload set, check if it's justified to use.

concepts are constrains on types, but we define them based on the algorithms, we define things based on usage.

std::enumerators - template specialization on enums that have the same name.

> - Inheritance is for sharing an interface.
>   - and so is overloading
> - Use a single names for a single entity
> - When you see two things with the same name, assume there is a reason for it.
> - When you have option to give two things the same name, **don't, unless** there is a reason for it.
> - To find concepts, don't study what your callees provide in common; study what your callers require
> - Default function arguments are the devil.

</details>

## Windows, MacOS and Web: Lessons from Cross-Platform Development @ Think-Cell - Sebastian Theophil

<details>
<summary>
Challenges for cross platform code.
</summary>

[Windows, MacOS and Web: Lessons from Cross-Platform Development @ think-cell](https://youtu.be/Cmud1jO__VA)

they started with a library that was developed in windows environment,it was a plug-in, and therefore, dynamically loaded and not in control of the entire process, many shared resources.
they

> "need a cross-platform toolkit that hides platforms specifics and **behaves identically** on different platforms"

(if such things can exists)

> Agenda
>
> 1. Levels of Abstraction: Hiding Platform Specifics
> 2. Kernel Object Lifetimes: Interprocess Shared Memory
> 3. Common Tooling I: Text Internationalization
> 4. Common Tooling II: Error Reporting
> 5. Moving to WebAssembly

### Levels of Abstraction: Hiding Platform Specifics

platform independent c++?
there are easy cases, like rendering, http requests (with the system API), child process and setting IO pipes. theses cases can be

> "Clearly defind as '**data In, data Out**'"

but even these cases can be difficult to make true platfrom indpendent, like direct call to rename/move files, which has different behavior flag for windows and macOs.

consider what the function really does and what it needs, what is the purpose of the function? if we know the "Why" - the reasoning for the function (what the user tries to achieve), we can tailor the "How" - what do we call in each platform. we don't simply route the arguments to the OS system call.

creating a file that is automatically deleted by the OS when the system closes (even at crush) but while it's alive it can be used by other processes. this behavior can be easily down on windows, but not on Mac, so maybe we need to rethink the 'how', and use a sqlite database for this in macOS, rather than file.

> - cross platform interfaces need to have well-defined, strong semantics.
> - weak semantics lead to subtle errors.
>   - Warning Sign: Having to look at the implementation.
> - Strong semantics increase DoF (degrees of freedom) for the implementor.
> - Too high-level.
>   - missed chance to unify code. Rare, we are lazy.
> - Too low-level.
>   - You'll force identical interfaces on very different things.
>   - semantics don't match operating system (_QFile::setPermissions_).
>   - or you'll loose a lot of expressiveness (_rename_).

### Kernel Object Lifetimes: Interprocess Shared Memory

boost and other libraries solve some of the problems for us, but sometimes we can to better.

boost offere interprocess communication tools, different shared memory behavior for windows and mac, windows cleans up, Posix can keep files alive for a long while. there are Robust Mutexes, file locking.

### Common Tooling I: Text Internationalization

a tool for text internationalization: translating, numbers formats.
text, context, plurality forms, what we wants.

[Boost.Locale] (https://www.boost.org/doc/libs/1_51_0/libs/locale/doc/html/main.html) was added in 2018 (boost 1.67), which supports tranlation by creating a catalog of transaltion, in boost it's runtime, in their implementation they try to make it constexpr. we don't want to read a file from disk, it's dangerous, we rather link the translations as part of the program.

> reminder about constexpr

strong and identical semantics can also refer to external tools in the build process.

### Common Tooling II: Error Reporting

dumping stack data to file, different for windows and Unix. they have an error report system that sends error to the backend and tries to identify the error. but file formats for dumps are different, and it needs to be standardized.
macOS allows to send access permissions to other processes.

### Moving to WebAssembly

the products ships with chrome extensions and webapp. they tried to use TypeScript (not JavaScript). but they weren't able to share data with c++. using C++ was typeunsafe because it lacked wrappers for JavaScript. so they built something of their own.
it's called 'Defiantly typed", and they have 'typescriptem'. which creates type safe c++ that does JavaScript.

in typescript, decleration order doesn't matter. so there needs to be some dependency list. typescript has non-integer enums, so they created a marshal enum template, and they had to create function callbacks.

</details>

## Iterators and Ranges: Comparing C++ to D to Rust - Barry Revzin

<details>
<summary>
How ranges are represented in langauges, different models of iterators.
</summary>

[Iterators and Ranges: Comparing C++ to D to Rust](https://youtu.be/d3qY4dZ2r4w), [Slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Iteration-Models-slides.pdf)

a sequence (list, vector, map, generator,etc..) of data, we need a unifrom set of operations (read, advance, check if done?).
C++ has an iterator pair model, one for the first elements (which we move), and one iterator for one-past-the last (end iterator), we can advance the iterator (forward, backward for bi-directional, or even jump if it's a random access), read from it and check if it's the same as the ending iterator.

### C++ Ranges

in c++ ranges, we have an iterator as one type and the sentinel as the end point.

in this example we have the pair (begin, end), andvance, read, and check if done. it's deliberately a classic for loop and not a range for loop.

```cpp
template <range R>
void print_all(R&& r)
{
    auto it = ranges::begin(r);
    auto it = ranges::end(r);

    for (;it != last;++it)
    {
        fmt::print("{}\n",*it);
    }
}
```

_transform_ (map in other langauges) and _filter_\
a slide of how to implement c++ _transform_, a bit about the _end()_ method that returns iterator rather than the sentinel in case of _common_range_, all sorts of const overloads. transform is a wrapper around the iterator behavior, the 'read' behavior is a function that uses the value returned from the underlying 'read'. with filter, we have a problem that the we can't tell before hand what is the first element (it's not a one-to-one relationship), so there are some issues about constant times and stashing, but in the end it's wrapper over find-if

### The D Ranages Model - Iterators Must Go

D supports popping (shrinking, slicing), we work directly on the range object, we read the front element and pop it to advance.

a slide about implementing D behavior in C++. a slide about map (c++ transform) and filter in D, a lot less boiler plate code, D has more pronounced problem with the filter behavior (it's harder to find the first element), it fixes it with something called 'prime', which ensures the first element is correct.

in csharp, we use an enumerator that starts before the first element.

```csharp
var l = new List<int>{1,2,3};
var e = l.GetEnumerator();
while (e.MoveNext())
{
    Console.WriteLine(e.Current);
}
```

he describes c++,D and c# as 'reading languages'

> - _read_ is a distinct, idempotent function (can call it as many times as we want and get the same result)
> - it has an intresting downside..
>   - in the example below, how many times is the 'some_operation' invoked?

```cpp
auto some_operation(int)->int;
void impl()
{
    std::vector<int> v = {1,2,3,4,5,6};
    auto r = v
        | map(some_operation) //c++ transform
        | filter([](int i){return i % 2 === 0;});
    for (int i:r)
    {
        fmt::print("{}\n",i);
    }
}
```

we actually have more than expected invocations! each element that satisfies the condition has an extra invocation, one to check if it's a match, and another to actually return the value. this happend in D and C# as well.

if the D model is simpler than c++ iterator pair model, then why wasn't c++ ranges implemented like D?\
looking at find-if example, and then trying to make a subrange of all the elements until that element. it's easy in c++, find the first match and use it as a cutoff point for the subrange.

```cpp
template <forward_iterator I, sentinel_for<I> S, indirect_unary_predicate<I> Pred>
auto until(I first, S last, Pred pred)-> subrange<I>
{
    I firstMatch = find_if(first,last, pred);
    return{first,firstMatch};
}
```

in D, things aren't as easy. D works with ranges, and there is perfect way to do it, we can do it lazy, like _take-while_, which is a range that lazy evalutes and continues until the predicate is matched. alternatively, we can implement it as like _take-range_, we use the predicate to count the number of elements before the first element that fulfills the predicate, and then take exactly that many elements
(position). in D we have a bit more algorithms, and we rely more on indices.

splitting ranges, taking a subrange. or even breaking the range according to some idea. we build everything on top of the search functionality.

```cpp
template <forward_iterator I, sentinel_for<I> S,forward_iterator I2, sentinel_for<I2> S2>
auto find_split(I first, S last, I2 first2, S2 last2)
{
    auto mid = ranges::search(first,last, first2,last2);
    auto pre = ranges::subrange(first,mid.begin());
    auto post = ranges::subrange(mid.end(),last);

    return tuple{pre,mid,post};
}
```

D has different methods for findSplit, findSplitAfter,findSplitBefore, and it doesn't have many chances for code reuse, it again needs to rely on indices.

### The Rust Iterator Model

Rust is very different from C++, it has one method, _next_, we can use it once to get the data, and that's it. it gives the data, moves to the next element, the checking is done by returning an OptionalReference.

the map operation is rust is simply taking the next element, if it's nullopt, return the nullopt, otherwise, return the result of invoking the mapping function on it. filter moves forward until it reaches the end or an element that satisfies the condition.
actually, python and java are similar.\
we can call them 'iterator languages', read isn't a separate operation. java is a bit different than rust and python, it has the _.hasNext()_, which uses a cached element.

rust also has cached element mechanics, which is called _peek_.

### Iterator Languages

if the reading languages, calling map and then filter had extra operations of the map function, in iterator languages, there is no extra overhead, the mapping is called once per element. but if we have drop operation, we still pay for the mapping, (because the read and advance operations are linker), while in reading langauges we don't need to read the data to advance over it.

in c++, we can take the iterator returned from find_if, and simply delete it.

in iterator languages, how do we remove the matched value? we need a different algorithm, instead of find_if to match the value, we search for the position and then delete the element in that position.

for group_by (group_on, partitaion_by,chunkBy, chunkOn) there are two approaches: binary and unary. we can implement the unary approach in terms of the binary approach. they return a range of ranges.

slides about rust, group*by and \_getlines*,

> Functional gaps in iterator languages
>
> - No container/iterator cohesion
> - What about algorithms?
>   - no binary _group_by_
>   - no _adjacent_find_ or _adjacent_difference_
>   - no _sort_
>   - no _slide_
>   - no _search_, _mismatch_ of _find_end_
>   - no _lower_bound_, _upper_bound_, _equal range_ or _binary_search_
>   - no _next_permutation_ or _prev_permutation_
>   - no _stable_partition_ or _rotate_ (_partition_ returns index)
>   - no _min_element_, _max_element_ or _minmax_element_

| language | type     | element       | read              | advance           | done?             |
| -------- | -------- | ------------- | ----------------- | ----------------- | ----------------- |
| C++      | reading  | iterator it   | \*it              | ++it              | it == last        |
| D        | reading  | range r       | r.front()         | r.popFront()      | r.empty()         |
| C#       | reading  | IEnumerator e | e.Current         | e.moveNext()      | e.moveNext()      |
| Rust     | iterator | iterator it   | it.next()         | it.next()         | i.next()          |
| Python   | iterator | iterator it   | it.\_\_next\_\_() | it.\_\_next\_\_() | it.\_\_next\_\_() |
| Java     | iterator | iterator it   | it.next()         | it.next()         | it.hasNext()      |

</details>

## Semantic Sugar: Tips for Effective Template Library APIs - Roi Barkan

<details>
<summary>
different way that templates could be used and specialized
</summary>

[Semantic Sugar: Tips for Effective Template Library API](https://youtu.be/u0rvEMV8Qq4), [slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Semantic-Sugar_-Tips-for-Effective-Template-Library-APIs-1.pdf)

template libraries. concepts were conceptualized even back in 2013,2014, before the language was able to provide them.

templates and overload resultion, writing the same algorithm for multiple types, metaprogramming for implementing different overloaded algorithms.

```cpp
template <class T>
constexpr const T & min(const T &a, const T &b)
{
    return (b<a) ? b :a;
}

template <class T>
constexpr void swap(T &a, T &b) noexcept;

template <class T2, std::size_t N>
constexpr void swap(T2 (&a)[N], T2 (&b)[N]) noexcept;
```

somtimes multiple overloads are legitmate, but one is preferable, so we can use _std::enable_if_ and SFINAE.

things that we will see this lecture

- Putting constrains on our templates
- C++20 Concepts- alternatives and ancestors
- Many opinions, some facts
- Tips and ideas, when should use various mechanisms
- Suggestions for changes to the language (opinions, not facts)
- Snippets from the STL
- Clips from Youtube

concepts are:

### A Bunch of Boolean Expressions

defining concepts with boolean expressions, and with c++20 'requires' keyword.

```cpp
template <class T>
concept integral = std::is_integral_v<T>;

template <class T>
concept signed_integral = std::is_integral_v<T> && std::is_signed_v<T>;

template <class T>
concept swappable = requires(T& a,T& b)
    {
        ranges::swap(a,b;)
    };
```

we could do this before c++20, with type traits (classes that have `::value` member), variable templates and function templates

[integral_constant](https://en.cppreference.com/w/cpp/types/integral_constant), \
[std::enable_borrowed_range](https://en.cppreference.com/w/cpp/ranges/borrowed_range), \
[std::is_pointer_interconvertible_with_class](https://en.cppreference.com/w/cpp/types/is_pointer_interconvertible_with_class)

```cpp
template <bool B>
using bool_constant = std::integral_constant<bool,B>;
typedef bool_constant<true> true_type;

template <class R>
inline constexpr bool std::enable_borrowed_range= false;

template <class S, class M>
constexpr bool std::is_pointer_interconvertible_with_class(M S::* mp) noexcept;
```

full expressiveness is possible [std::is_scalar](https://en.cppreference.com/w/cpp/types/is_scalar), is defined with boolean OR expressions

```cpp
template <class T>
struct is_scalar : integral_constant<bool,
        is_arithemetic<T>::value ||
        is_enum<T>::value ||
        is_pointer<T>::value ||
        is_member_pointer<T>::value ||
        is_null_pointer<T>::value
    >
```

SFINAE, void_t, the detection idiom a way, to use something like 'required' in pre c++20 standards (the new syntax makes things easier to read a and write). the default is false, but we specialize on the true types.

```cpp
template <typename T,typename =void>
struct has_meow : std::false_type{};
template <typename T>
struct has_meow<T, void_t<decltype(std::declval<T>().meow())>>
    : std::true_type();
```

**concepts still don't allow specialization**

```cpp
template <class T> struct is_const : std::false_type{};
template <class T> struct is_const<const T> : std::true_type{};
```

out in/opt out specialization, the std::enable_borrowed_range can specialized to true and opt-in to get some functionality.

```cpp
template <class R>
inline constexpr bool std::enable_borrowed_range= false;
```

predicates on traits (not type traits), here the temperature class is specialized to have predicate

```cpp
namespace std {
    template<>
    struct numeric_limits<Temperature>{
        static constexpr bool has_infinity = false;
    };
}
```

### Take the Overload that Meets the Largest Number of Predicates

controlling library-application interation

> - When applications use libraries there's a risk of error due to incorrect expections.
> - Overload-resolution is a way to try and verify expectations are matched.
> - This can be an 'on/off' constraining to allow/disallow certain interactions, or more advance mechanism to choose or tailer specifs of an interactions.
> - some resolution mechanisms can easily be bypassed, while other are less negotiable.

overload resolution with concepts:

> - 'requires' clause
>   - Two more syntax alternatives for good measure
> - The most specialized version wins (see standard for details)
> - SFINAE friendly
> - Clear error messages
> - Faster compilation speed
> - Library defines requirements - Application must conform.

a 'requires clause' is not a 'requires expression'

we can impose restriction from the library side - the library dictates the constraints

> - std::enable_if
>   - library guided, the requirements are defind by the library
>   - no ranking, error on multiple matchs
> - "partial specialization" - choose the function more specialized than others (be carefull of universal forwarding functions)
> - ranking down by the compiler
> - Tag dispatch
>   - this is what the STL uses
>   - iterators opt-in to their category/concept
>   - in the STL this dispatch is hiedden as an implementation detail,
>   - libraries could technically allow call-site override.

```cpp
//from the stl
template <class _InputIter>
inline void advance (_InputIter & __i,typename iterator_traits<_InputIter>::differene_type __n)
{
   __advance(__i,__n,typename iterator_traits<_InputIter>::iterator_category() );
}
```

alternatively, we can have constraints coming from the application,

```cpp
template <class ForwardIt, class Compare= std::less<>>
constexpr ForwardIt max_element(ForwardIt first, ForwardIt last, Compare comp =Compare{});
```

> - Policy-Based Desgin
>   - this is what we use in many stl algorithms.
>   - the callers can overide at the call-site.
>   - (isn't this the strategy design pattern?)
> - Customization Points (and CPOs, tag_invoke)
>   - Algorithms have a default.
>   - Algorithms that can be specialized by the library, but for the entire type, not per call.
>   - Examples: std::swap, ranges::ssize, ranges::empty,
>   - CPOs are objects with operator() that deal with overload resolution intricacies.
>   - 'Niebloids' - similar mechanism for ADL avoidance.
> - Behavioral Properties (P1393, C++23 executors, Hollman & Niebler)
>   - _std::require(executor, execution::blocking.always);_
>   - Library defines properties and Application can use them.

CPO (customization points objects) are callable objects (have the `()` method) that help with overload resolution,tag invokes is an attempt to standardize the CPO idea. Niebloids are a similar but different mechanism.

maybe in c++23 we can have behavioral properties,

Overload Resolution / Customizations

| type                   | On/Off (compiler error?) | Choose from Few | User Code   | Simplicity                                             |
| ---------------------- | ------------------------ | --------------- | ----------- | ------------------------------------------------------ |
| _requires_             | Library                  | Library         | No          | Yes                                                    |
| _std::enable_if_       | Library                  | Library         | No          | No                                                     |
| 'Specialization'       | Library                  | Library         | No          | Yes                                                    |
| 'Tag Dispatch'         | Application              | Application     | No          | Medium                                                 |
| Policy Based Design    | N/A                      | No              | Caller      | (simple for algorithms, less so for classes and types) |
| CPOs                   | Application              | No              | Application | Medium                                                 |
| _std::require_ (c++23) | No                       | Caller          | No          | Yes                                                    |

Advanced Overload Resolution Schemes

| type                   | On/Off      | Choose from Few       | User Code        | How?        |
| ---------------------- | ----------- | --------------------- | ---------------- | ----------- |
| _requires_             | Library     | Library \ Application | No               | Warrents    |
| _std::enable_if_       | Library     | Library \ Application | No               | Warrents    |
| 'Specialization'       | Library     | Library               | No               | N/A         |
| 'Tag Dispatch'         | Application | Application \ Caller  | No \ Application | Expose/Add  |
| Policy Based Design    | N/A         | No \ Caller           | Caller           | Policy tags |
| CPOs                   | Application | No (runtime)          | Application      | N/A         |
| _std::require_ (c++23) | No          | Caller                | No               | N/A         |

### Syntactic and Semantic

semantics can be tricky. like

> - _std::is_trivially_copyable_v\<std::pair\<int, int>>_ -> **false**.
> - the complexicy of _std::list::size()_ - was constant or linear until c++11, but required to be constant in c++11.

trivially copyable means we can do memcopy rather than call constructors, but despite that, a standard pair of int is syntactically not trivially copyable (same as tuples), because it would constitute as an ABI break because of past reasons. std::pair has none-trivial assignment operators (to work with rvalue references).

std::list::size() was implementation dependant (linear or constant) for a while, but this was changed (which required an ABI break) for c++11.

there are escape hatches that allow specialization to opt out from behaviors in order to implement things differently. a positive escape hatch is a 'warrant', a way to opt-in to behaviors that are default disabled. this is dangerous, 'footguns' (a way to shot yourself in the foot). we saw this earlier with _std::ranges::enable_borrowed_range_, which is default false.

'cheaply_copyable_t' - from Herb Sutter's lecture in CppCon2020.

special cases:

_std::equivalence_relation_ - a relation that is reflexive (f(x,x) is true), symmetric(f(a,b) == f(b,a)) and transitive (if f(a,b) is true and f(b,c) is true, then f(a,c) is true). there is an issue that the compiler can't differentiate between the general relation and specific equivalence_relation.

example of semantic sugars to attach semantics to lambdas.

### Take Away

> - Concepts are great
> - 'requires' doesn’t require concepts
> - Library writers - give your users power
>   - Build escape hatches / warrants
>   - Consider call-site customizations
> - C++ Standard
>   - Consider concept specialization
>   - Consider type-trait specialization

</details>

## What is an ABI and Why is Breaking it a Problem? - Marshall Clow

<details>
<summary>
What are ABI changes, and what would happen if we break the ABI?
</summary>

[What is an ABI and Why is Breaking it a Problem?](https://youtu.be/-XjUiLgJE2Y), [slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Slides.pdf)

in 2020, there was a formal request to the standard committee to commit to breaking ABI in the future, people wanted to know that the committee was ready to do so if needed and in order to improve the language. the committee didn't fully respond.

ABI - application binary interface.

changes to the standard library that would entail an ABI break.

> the ABI includes stuff such as:
>
> - Structure layout.
> - Vtable layout.
> - Parameter passing convetions.
> - Name mangling.
> - Exception handling methods.
>
> But also - Library changes, which are different from compiler changes.

### The One Defintion Rule (ODR)

> "If there is more than one (non-identical) defintion of an entity visible in a program than the behavior of the program is undefined"

the actual term for the standard is "Ill-formed, no diagnostic required" (IF-NDR). this means that the toolchain is allowed to produce a program that the can do anything, and doesn't have to tell you that it has done so.

examples of ODR violations:

> 1. Two diffrent defintions.
> 2. Change the layout of the struct.
> 3. Add a virtual method.
> 4. More subtle things.

in this example, code with header1.h belives that Foo's size is 8 and b is at offset 4, and code with header2.h believes the size is 12 with b at offset 8. if we go around passing objects, the defintions are different and things will go bad.
cpp isn't python, we use offset to determine members, not member names.

when this effects class inheritance or composotion, it is called "**the fragile base class problem**", if we change the base, everything must be re-built.

```cpp
//header1.h
struct Foo
{
    int32_t a;
    int32_t b;
};

//header2.h
struct Foo
{
    int32_t a;
    int32_t added;
    int32_t b;
};
```

adding a virtual method is the same is changing the layout of the vtable.

> Variations on a theme:
>
> 1. Removing a member.
> 2. Reordering members.
> 3. #pragma pack.

removing a problem is the same as adding, reordering changes the offset, changing the #pragma pack can both change the size and the offsets.

```cpp
//header 1.h
struct Bar{
    virtual int One(int);
    virtual ~Bar();
};

//header 2.h
struct Bar{
    virtual int One(int);
    virtual double Two(std::string);
    virtual ~Bar();
};
```

the vtable is a static struct with function pointers, we can't say how the members are ordered, it is up to the implementation of the compiler, it can be in lexical order or the order of declaration.

in this example we have a pair with two members and a copy constructor with memberwise copy. the other pair uses default copy constructor. but with trivially copyable pairs, we can get better performance specialization.

```cpp
//header 1.h
template<Typename T1, typename T2>
struct pair {
    T1 first;
    T2 second;
    //...
    pair(const pair &p): first(p.first),second(p.second) //copy constructor
    {}
};

//header 2.h
template<Typename T1, typename T2>
struct pair {
    T1 first;
    T2 second;
    //...
    pair(const pair &p)=default
};
```

> on some platforms, parameters of trivially-copyable types which can fit into a register are passed in a register instead of on the stack.

if one piece of code expects the data to be on the stack and the other puts it on the register, this is a serious debugging challenge. this will be solved with a full re-compile of the code.

users can re-compile all of their code to make it fit the new defintions, but this can't work with the standard library. this was an actual problem that required the committee to do some special trickery to ensure this issue won't happen.

| issue                    | header1.h | header2.h        |
| ------------------------ | --------- | ---------------- |
| size of Foo              | 8         | 12               |
| offset of Foo.b          | 4         | 8                |
| Bar vtable size          | $2*8=16$  | $3 * 8 = 24$     |
| trivially copyable pairs | no        | yes (supposedly) |

so if IF-NDR is so scary, why can't the compiler diagnose this?

three cases to consider

> - Two different defintions in the same translation unit.
> - Two different defintions in different translation units, statically linked.
> - Two different defintions in different translation units, dynamically linked.

the first case is covered, we get a warning from the compiler. the second case is theoretically possible to diagnose, if we make object files bigger and the linker does more work. the third case doesn't involve the toolchain, it's getting done by the program loader, it happens after the compiler and the linker did the work, the objects might have been compiled by different compilers at different times. the chain of dynamic linking can include many object files.

### From ODR to ABI Break

> - "ODR violation between the environment in which the program was built and the environment in which it is run."
> - "An ABI break is just an ODR violation in time."
>
> we can also consider two different versions of the same file as two files. say we install a new version of a shared library with an updated header file, we have program A that uses this library, but we don't recompile it. when we run A, it will load the updated shared library, but will use it as if it was the old version.

but what can we do with this?

> - Don't change things that effect ABI.
> - Don't have 'stale binaries'.
> - Have only one defintion for everything.

we don't always know that something is an ABI change. avoiding stale binaries mean building everything, everytime, from scratch, which is not only time consuming, but what about external libraries which we don't have the source code for?

### Examples of ABI Breaks Happenning in the Past

in c++17, _emplace_back_ was changed to return a reference to the newly created emplaced object (rather than void), this didn't cause problems for existing code. (the return type is not part of the mangling).

for c++11, _libstdc++_ changed the layout and behavior of _std::basic_string_ . the old string class implemented '_copy-on-write_' semantics, while the new one did not. this change was mitigated by providing a flag that retained the old behavior \_GLIBCXX\*USE_CXX11_ABI\*. so it was more of transition than a break, but it was still a big pain in the ass, which occasionally pops up in these days. there are actually two basic_strings implementations still lurking around, and each vendoer chooses which layout to use.

in c++20, there was rejected proposal for a 'half float' type (16 bit, like short), which would play very well with GPUs. it was rejected (in part) because that would involve adding virtual functions on the iostream to support _num_put_ and _num_get_, which would break code using iostream. (there were also other reasons).

### And Back Into Politics

the standard doesn't say anything about ABI, compatability, versions, and so on. in the past, ABI breaks were stopped by implementors, they would warn about this or speak up about changes. they don't want the flack from users, and they are the ones who interact with users.

[P2028: What is ABI, and What Should WG21 Do about it?](http://www.open-std.org/jtc1/sc22/wg21/docs/papers/2020/p2028r0.pdf) - the paper by Titus which called for ABI changes (breaks). this would improve performance, but the suggestion generated a lot of discussion (which this talk is part of).

These changes would have to be detectable, we can't have wide-spread ABI changes running around.

One suggestion would be to change name mangaling scheme and prevent linkage of new files with old files. effectively breaking the c++ language in half. new programs could only interact with new object files.

Another suggestion is to have 'fat binaries', which contain defintions for both the old versions and the new versions, and then stuff would be determined at run time / linkage. the problem is that until the program is run, we don't know what is used (which shared library, plugin's, executables, etc...), and by that time, information has already been stripped away and it's impossible to tell which version/vendor was responsible for them.

for users:

> - If you have source to every bit of software that you use, and are willing to rebuild it, then this is not a problem for you.
> - If you never use any third-party software, then this is not a problem for you - your OS vendor will resolve any issues.
> - Otherwise, if you have binaries that use C++ internally, then this would affect you.

a case from a photoshop user, a story of how ABI breaks would effect it. the user has 3rd party-plugins (shared libraries), which assuming an ABI change, all break on update (in the best case, we know that the problem is the abi change). some plug-in creators will have the update, some will charge for it, some will take short time, some long time, and some will never update. most users will avoid the update to keep their current plug-ins operational.

### Summary

> - There’s a real problem here.
> - Historically, the committee has prized stability.
> - ”We” would like a solution that will allow us to make changes.
> - We do not have such a solution today.

We assume stability and backward compatibility, this was historically prized in c++. but we don't stability to mean stagnation, if changes can be made to make stuff better, we want those changes to happen. can we make them happen? what about closed-source software - how can users be protected? and if we change the ABI, how do we make this safe to change more than once?

</details>

## What Belongs In The C++ Standard Library? - Bryce Adelstein Lelbach

<details>
<summary>
How the standard is changed, what are the considerations, and what can be done.
</summary>

[What Belongs In The C++ Standard Library?](https://youtu.be/OgM0MYb4DqE),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/04/what_belongs_in_the_cpp_standard_library__r2__2021_05_07_cppnow.pdf)

(starts the talk by introducing **#include <C++>**, an initiative to increase divercity in the c++ community).

the Standard C++ committee and the language evolution sub committee.

### What has made C++ successful?

is it performance (zero cost abstraction - don't pay for what you don't need)? portability? stability? there are other langueges that do each of those things better, and C++ sometimes falls short in them.

But she argues that c++ is universal, solve any type of problem (HPC, Finance, gGaming, Apps, Robotics, Vehicles) using any paradigm (Imperative, OOP, Functional, Generic, Parallel, Reactive) for any platform (windows, linux, mac, embedded).

> A multi purpose and multi paradigm language.

it's both the advantage and the weakpoint of the language, it's hard to move forward and set priorites that are acceptable for everyone. each domain has different, sometimes conflicting, needs and desires.

> "Use Case Sympathy" (noun)
>
> Accepting the importance and validity of use cases that you are not personally familiar with or believe in.

the committee needs to find solutions that are relevent and usefull for a large number of the language users. but the more universal a feature is, the harder it is to get the standardization right. but we also don't want to prioritise one subset of users in the community and cause fragmentation, if the changes negatively effect other users

incrementalism - features are relased in limited form, and are expended later on.

an example topic:

**Should everything in std:: support allocators?**

for some domains, controlling allocations is critical, like determinstic allocations, performance critical. and while c++ treats all memories the same way, this isn't true, there is shared memory, io memory, and gpu memory.

can most container support allocators? even for the basic _std::vector_ it's not so simple.

(something about statefull allocators )

- `propagate_on_container_copy_assignment`
- `propagate_on_container_move_assignment`
- `propagate_on_container_swap`

and also the other containers, _std::tuple_, _std::pair_ and _std::optional_, they don't allocate memory themselves, but they might hold elements that do. there are also type-erasing facilities such as _std::function_ and _std::generator_ (in c++23), how do we recover the type-erased allocator?

in the past, _std::function_ supported allocator, but it was removed in c++17 [cpp-reference page](https://en.cppreference.com/w/cpp/utility/functional/function). are we going to keep delaying the release of _std::generator_ until we manage to figure out allocator support

**Should everything in std:: have a type erased form?**

type erasure is important for some, such as dynamic typing, plugins, reducing build times, hiding away types...

but should everything support it? ranges, iterators?

```cpp
//currently
template <typename R>
requires std::ranges::random_access_range<R>
void my_algorithm(R && r)
{
    //...
}

//should this also exist?

void my_algorithm(std::any_random_access_range<int> r)
{
    //...
}
```

it's hard to decide what's in scope, and what should get priority, it's hard to be universal, it's a source of strength and a weakness. if the committee doesn't support all it's users, then it risks the entire community.

### What is the C++ Standard Library

we shouldn't conflate the **standard library** with a **standard library implementation**.

> - GCC's libstdc++ is not **The** C++ Standard Library.
> - MSVC's STL is not **The** C++ Standard Library.
> - LLVM's libc++ is not **The** C++ Standard Library.
> - NVIDIA's libnv++ is not **The** C++ Standard Library.
>
> The C++ Standard library is a **specification**.

however, this means that it's not an efficient way of delivering features, each implementation does duplicated work when they all implement the same features separately.

> The C++ Standard library is a **descriptive**, not **prescriptive**.

it's the structure of the code, the semantics of the abstract machine and requirements. it's a principle called **Implementation Freedom**. Enough to be portable and consistent, but not so much as to dictate the design.

> "implemenation-defined and undefined behavior are often a feature, not a bug."

example of _std::mutex_, different implemations have different advantages.

| Implementation     | Supported On                      | Pros                                                | Cons                         |
| ------------------ | --------------------------------- | --------------------------------------------------- | ---------------------------- |
| OS kernel mutexes  | Older and newer operating systems | Fair, Good perf under contention                    | Higher latency               |
| Futexes            | Newer operating systems           | Fair, Lower latency                                 |
| Spinlocks          | Bare metal                        | Much lower latency, Never yields,Doesn’t need an OS | Unfair,Less energy efficient |
| No synchronization | Single core platforms             | No overhead,Doesn’t need an OS                      |

Also the index operator _[]_ on containers: `std::vector::operator[]` and `std::string::operator[]`. some people said that they should perform out of bounds checking, and some say that they shouldn't, as it harms performance. therefore, the standard doesn't **require** or **forbid** out of bounds checking, it **permits** each implementation to decide. so the users can decide and choose different implementions for different scenarios (like debug and production).

**Standardization takes time**\
Is the committee too slow? too fast? unlike other libraries, the standard library can't choose to drop support for one platform, so the time to deployment is much longer. this is why we have such a large gap in standrads releases, and why even today, the industry still hasn't fully integrated the changes from c++11. the committee plans and operates in spans of decades,not months or even years.
Because of that, the scope of what the standard can focus on is much more limited.

**The _std::_ implemnators aren't domain experts**\
They don't have both the necessary knowledge and skills to maintain C++ Standard Library implementations and the specialized domain expertise for each domain for which c++ is used. so it's sometimes better to grab a domain specific library which fits the platform you are using and has optimized performance than wait for the C++ Standard Library to figure out all the kinks, which might result in a less optimized solution.

a prime example is _std::regex_, which is notoriously slow. boost's implementation was the foundation for it, but is still much faster than most standard implementations.

> implementers are experts at:
>
> - Their specific platform.
> - Balancing tradeoffs.
> - Handling corner cases.
>
> They are less good at:
>
> - Domain specific work:
>   - Math special functions
>   - _std::regex_
>   - `<charconv>`

### Stability vs Velocity

it's not so clear as to say "one or the other".

> Hyrum's Law:\
> "With a sufficent number of users, it doesn't matter what you promise in the contract: all observable behaviors of your system will be depended on by somebody."

every observable part of an implementation is implicitly part of the interface.

there is an example that in one case, the implementation of _std::string_ had to be delayed for one company because a customer of theirs was relying on the ability to destroy the same string twice.

- API - Syntax & semantics. source code, in the C++ standard.
- ABI - Binary representation & conventions. compiled code, platform specific.

> C++ Language ABI:\
> Binary represtation & conventions for language facilities.
>
> - function calling conventions
> - name mangling
> - layour and size of types
> - layour of virtual tables
> - exception handling
> - floating point mathematics

it's hareder to change, and can have cascading effects

> C++ Standard Library ABI:\
> Binary represtation & conventions for C++ Standard Library facilities.
>
> - linkage of _std::_ functions
> - _std::_ name mangling (of types and function)
> - layouts and size of _std::_ types
> - _std::_ virtual tables
> - _std:: constexpr_ values and functions
> - `<type_traits>` and _std:: concepts_

the focus is on the C++ Standard Library ABI, not the core language facilities.

> **API Stability**: existing syntax and semantics should rarely change.

it's important, but we would like to sometimes change them.

> **ABI Stability**: binary representations of existing facilities should rarely change.

there are voices who question the case for ABI stability.

- backward compatiblity: older code, new builds.
- forward compatiblity: newer code, old builds.

also having objects which were compiled as different versions depending on one another. users don't always have the option to compile everything from scratch. how can address this problem today?
What if we have multiple components that depend on different versions of other components?

one suggested solution is "don't upgrade if you can't get newer dependencies", but that makes it harder to upgrade, and stops adoption from happening as everyone waits for something else to finish upgrade. this is what happened with python, which took years to transistion from python2 to python3.

it's also important to consider how th breaking change manifests:

- at build time: compile or runtime?
- run time: can we detect it? is a graceful or catastrophic?
- are these breaks consistent? will they always happen?

if the layout of _std::string_ changes by requireing small string optimization, we can run into this problem, this won't be caught at compile time, and the bug will be hard to track. the receiving function from the old ABI would treat the string as one without the SSO, and will missinterpet it.

```cpp
void f(std::string &s); // compiled as c+11 code. no small string optimization

void g() //this is now c++23 imaginary code with the abi change
{
    f(std::string("BAL")); //small string
}
```

or in a return value from a function.

```cpp

std::string f(); // compiled as c+11 code. no small string optimization

std::string g() //this is now c++23 imaginary code with the abi change
{
    std::string s = f(); //small string
    return s + "bryce";
}
```

we can also run into problem with our own data types, if they have members from the standard library

```cpp
//this was compiled as c++11
struct X {
    std::string s;
}
X make_x()


// and this was compiled after the ABI change
void g()
{
    X x = make_x();
    x.s = "hello world\n";
}
```

also inlining of data members and functions can cause problems (as they use the old implementation rather than the shared implemention), ODR (one defintion rule) mayhem. there is also problem with constexpr stuff. also concepts.

also polymorphism, removing virtual tables changes the layout, adding virtual functions change the address.

> C++ Standard Library polymorphism, type erasure and named concepts are fixed forever.

this is a problem for incremental development, this makes the deployment of features harder, as there is less room for course correction.

#### examples

> In c++03, _std::list::size_ can have linear complexity, no size data member needed.\
> In c++11 _std::list::size_ must have constant time complexity, a size data member is required.

this was an ABI breaking change.

```cpp
template <class T,class A =std::allocator<T>>
class std::list03
{
  __list_node<T> root;
};

sizeof(std::list03<int>) == 16;

template <class T,class A =std::allocator<T>>
class std::list11
{
  __list_node<T> root;
  std::size_t size;
};

sizeof(std::list11<int>) == 24;
```

also, in c++03, _copy on write_ was allowed for _std::string_, but it was prohibited in c++11. (Copy on write - delay the actual copy of the string until it needs to change). in c++11, the specification changed, and it was also an ABI break.

those two changes were massive, and caused a lot of issues, maybe even delayed the adoption of c++11.

in c++11, _std::lock_guard_ had a single template parameter, so using multiple locks was problematic and could cause deadlocks. in c++17, it was supposed to become a variadic template, and could take multiple arguments. which worked great with class template argument deduction, and made the process of taking multiple locks much easier to handle.\
While the API (the source) itself was backward compatible, it did change the name mangeling of the function. this was an ABI change the committee didn't want to push forward, so the name was changed to _std::scoped_lock_.\

```cpp
template <class MutexType>
struct std::lock_guard;
//...
{
    std::lock_guard<std::mutex> l0(mtx0);
    std::lock_guard<std::mutex> l1(mtx1);
}


template <class... MutexTypes>
struct std::scoped_lock;
//...
{
    std::scoped_lock l(mtx0,mtx1);
}
```

while introducing a new feature allows us to avoid breaking changes to the ABI, it also clutters the implementations and bloats it. as both functionalities must now be maintained.

again in c++17, a new virtual function overload was suggested for _std::system_error:message_. one that would not allocate memory and would not throw.

> "Proposed addition was non-pure, existing derived classes would continue to compile"

```cpp
virtual char const * message(int, char*, size_t) const noexcept;
```

but because it would change the virtual table layout, it would be an ABI break, and it was rejected.

these restrictions also effect performance

in c++17, _std::shared_mutex_ was added in addition to _std::mutex_, it was a more complex type, so the simpler type should be faster and more performant. however, the MSVC implementation is faster, because the older version must continue to support the old abi which uses windows _CRITICAL SECTION_ api, rather than the newer _SRW lock interface_. the older version is also ten times the size of the new version, but it can't be changed.

> "The C++ Standard Library is good at stability, but bad at fixing mistakes"

so, do we have to choose between stability and velocity?

there are some users who must have stability, but other don't. we can't fix this via policy, a technical solution is needed.

#### Proposed Fixes

c++11 introduced inline namespaces, members of the inline namespace can be used as part of the exterior namespace, but have a mangled name with the internal namespace.

```cpp
namespace std {
    inline namespace __cxxNN {

    template <class C, class T = std::char_traits<C>, class A = std::allocator<C>>
    class basic_string;

    using string = basic_string<char>;
    }
}
```

so we our previous example,the source code remains the same, but the mangled code now directs to different implementions of _std::string_.

```cpp
void f(std::__cxx11::string &s); // compiled as c+11 code. no small string optimization

void g() //this is now c++23 imaginary code with the abi change
{
    f(std::__cxx23::string("BAL")); //small string
}
```

unfortunately. this isn't fully supported today, the implementations don't support multiple versions of the standard in the same translation unit.

also, what if we have a function that returns a object which has different implementations? the mangled name doesn't contain data about the return type, so we can't use this to detect calls to functions which were compiled differently.\
we also can't detect breaks for data types with _std::_ members inside of them. the data members of a type effect the size and layout, but not the mangling, so we can't use this to catch ABI changes.

inline namespace can help with ABI breaks, they can help us detect them, but not solve the problem.

another solution is the `abi_tag`, which can be applied as an attribute to namespaces, functions of variables. the tag is viral and recursive, if the function signature contains something with the tag, then so does the function. the tag is added to the mangled name.

```cpp
namespace std {
inline namespace __cxxNN __attribute__((abi_tag)) {
// ...
}}
```

like before, it helps us find problems, but it doesn't solve the issue.

another possible fix is `std2::`, a new version of existing `std::` features which are not compatbile, but might be interoperable. should everything have a new version? this came down to an understanding about this approach:

> **std2::approach** - Any solution that is equivalent to "duplicate and maintian multiple generations of the same facilities".

this is what happened with _std::scoped_lock_, they cause maintainence burden, and causes confusion, as many ideas now have different names.

but what if the type system had the interface implemented in it?

```cpp
struct point {
    interface(std::cxx23) { //interface tag
    int x, y, z;
    interface(std::cxx26) int w; //interface block, only for std::cxx26
    int get_x() const { return x; }
    int get_y() const { return y; }
    int get_z() const { return z; }
    int get_w() const interface(std::cxx26) { return w; } //only for std::cxx26
    }
};

sizeof(interface(std::cxx23) point) == 12
sizeof(interface(std::cxx26) point) == 16

```

they support forward compatiblity, with resiliency overloads.

```cpp
void f(interface(std::cxx23) std::string & s); //only support std::23
void f(interface(std::cxx23+) std::string & s); //support std::23 and above
```

there will be drawbacks, like losing type erasure and inlining capabilites in resilient functions, but they would only come into effect when we decide to use them.

| ABI Problem         | Internal Name space | `abi_tag`     | Interfaces         |
| ------------------- | ------------------- | ------------- | ------------------ |
| Parameters          | Diagnose only       | Diagnose only | Diagnose and solve |
| Return Types        | No effect           | Diagnose only | Diagnose and solve |
| Non-Local Variables | Diagnose only       | Diagnose only | Diagnose and solve |
| Data Members        | No effect           | No effect     | Diagnose and solve |
| Inlining            | No effect           | No effect     | Diagnose and solve |
| Constant Evaluation | Diagnose only       | Diagnose only | Diagnose and solve |
| Polymorphism        | Diagnose only       | Diagnose only | Diagnose and solve |

### The Stability Thesis

> Until we learn to change things after we ship them, the C++ Standard Library should only contain things that are unlikely to need many changes.

the standard library shouldn't innovate, the c++ community should innovate. the changes should come from the field before being standardized, we need to wait until the research and discussion is done before putting things into the standard. **Avoid premature standardization!**

as an example, the _std::_ unordered containers (like _std::unordered_set_) were node based, but today, we know that flat containers can be better in terms of performance, but we can't change the standard to fit that.

Field experience is needed before standardization

- Implementaiton experience: was this implemented before? tested? on enough platforms?
- Usage experience: how did users accept this? is this up to their needs?
- Deployment experience: do we understand the maintence burdens?

| Good Implementation Experience       | Better Implementation Experience           |
| ------------------------------------ | ------------------------------------------ |
| Prototype                            | Production                                 |
| Preceding or similar to the standard | Written from or conforming to the standard |
| For one platform                     | For multiple platforms                     |
| In any publicly available codebase   | In a C++ Standard Library codebase.        |

> Incrementalism is key to C++ Standard Library Evolution.\
> We’re bad at changing things, but we’re good at extending things.

a flowchart to determine if something should go in the inital release.

#### The Necessity Thesis

> The C++ Standard Library should only contain facilities that can’t live elsewhere.

> **Language Support:** Facilities that require language support for correct or optimal implementation.
>
> - `<type_traits>`
> - _std::stacktrace_
> - _std::tuple_element_
> - _std::memcpy_
>
> **Portability**: Facilities that provide portable abstractions of platform-specific behavior and interfaces.
>
> - _std::chrono_
> - _std::atomic_
> - _std::sort_
> - _std::numeric_limits_
>
> **Vocabulary**: Facilities that need a common definition for interoperability across the C++ ecosystem.
>
> - **Interface Vocabulary**: Concepts, types, and operations that commonly appear in C++ interfaces. Common definitions means different codebases can interoperate.
>   - Concepts
>   - Containers and Views (ranges and iterators, string_view)
>   - `<algorithm>`
>   - _std::format_
> - **Tooling Vocabulary**: Facilities that tools want to recognize.
>   - MSVC iterator debugging.
>   - GDB container pretty printing.
>   - Clang thread safety analysis.

#### the Priorites

- Asynchronous and Parallelism - hopefully executers will help.
- Input and output
- Text processing
- Metaprogramming & Reflections
- Compile time expression support.

An alternative to the necessity thesis is the _usefullness thesis_

> The Usefulness Thesis:\
> "The C++ Standard Library should expand in scope to contain anything that is useful to C++ programmers."

this isn't good, this is a recipe for losing focus, there are many burdens in being in the standard library. but features want to be part of the base package, to be default availability?

(another flow chart)

the C++ Standard Library is not a package manager. it shouldn't be. but currently people aren't using package manegers enough. but maybe it should standardize how package managers operate? this could be nice. what about a standard build system (why not stick with cmake)? this is a better path to pursue rather than pushing everything into the standard. this will help facilitate innovation.

another suggestion is to intoduce an intermediate scoped library, between the default standard library and individual external libraries.
middle ground - solutions with a single implementations, like how boost used to operate. they wont need to be re-implemented for each C++ Standard Library implementation.

| Library                     | Availability                                             | Stability                       | Implementation                           |     |
| --------------------------- | -------------------------------------------------------- | ------------------------------- | ---------------------------------------- | --- |
| C++ Standard Library        | Comes with toolchain by default                          | 10-20 year stability guarantees | Implemented separately for each provider |
| C++ Collections (suggested) | Optionally comes with toolchain                          | Flexible stability guarantees   | Single code base                         |
| External Libraries          | Must be acquired separately; aren’t available by default | Flexible stability guarantees   | Single source                            |

### Audience Questions

- why do companies use c++?
- why do companies support and get involved in the c++ committee?
- is cmake really the only future? is standardization needed
- should we keep the same 3 year cycle of standard releases
- what happend to library technical specifications? why did the committee stop releasing them?
- did we over-standardize parallelism?
- again, who should take care of the package manager standardization / deployment?

</details>

## What Neighborhoods Owe Each Other: Rights, Duties, Immunities & Disabilities of Functions & Objects - Lisa Lippincott

<details>
<summary>
Defining code behavior in terms of contracts. 
</summary>

[What Neighborhoods Owe Each Other: Rights, Duties, Immunities & Disabilities of Functions & Objects](https://youtu.be/pDpdb6T1c1k)

a vision of programming, having neighborhoods

> Neighborgood: a portion we can reason about independently.

we already do this in a way, we have functions, objects, etc...

the function interface is the "area where neighborhoods interact". everything that goes inside our outside the function passes through those interfaces.

we also have "claims", or pre- and post- conditions. those conditions are connected to one another across different levels.

a simple example:

```
void foo() implementation
{
    bool b = bar();

    if (b)
        return;

    if (b)
        fail;
}
```

reaching the fail path is outside the scope of the neighborhood, getting to it can only happen from outside interference.

another sample where reaching the fail statement means that something outside the neighborhood went wrong. the _false_ value acts as an interface to a world of boolean expressions.

```
void foo() implementation
{
    bool b = false;

    if (b)
        fail;
}
```

a harder example, the true and false now have phantom parentheses to indicate that they are akin to function calls, and they represent a gateway to a different neighborhood.

```
void foo()
// implementation
{
    bool b = false();

    if (b)
        fail;

    b = true();

    if (b)
    {}
    else
        fail;
}
```

the questions this example raises are

- Repetition
  - Q: "Why do branches sometimes repeat?"
  - A: "Branched repeat when neighborhood has a right to the stability of an object."
- Not repeating
  - Q: "And why is that not the case here?"
  - A: "To perfrom the assignment, this neighborhood must relinquish its right to that stability."

### Rights And Duties

in a previous lecture "Locally atomic capabilites: and how to count them" (may 2017), there were some missing explaintations.

Wesley Newcomb Hohfeld (1879-1918) wrote a paper "Some fundamental legal conceptions as applied in judical reasoning" (1913), which helped guide the current talk. the paper pushes for the words of the law to be more precise, and describe contracts in unambiguous ways. he provides two terms:

> - Right: The ability to call upon another to act.
> - Duty: The responsability to act when called upon.

a Right and a Duty are a pair, two parts of a whole, but the purpose of this relationship derives meaning when two different entities hold the two parts. there is no point in having both the right and the duty.

> - No right: The absence of right, the lack of ability to call upon another.
> - Privilege: The absence of duty. the lack of responsability to act when called upon.

for programming we focus on Right and Duty, the "No right" and "Privilege" are the background, the common, unspoken case (they can come up when discussing threads)

the stability of a boolean objects is the stability of it's bytes.

```
inline claimable bool::stability() const
{
    entail_right byte0.stability();
    entail_right byte1.stability();
    entail_right byte2.stability();
    entail_right byte3.stability();
}
```

we can't define the stability of bytes in cpp, but it would look like this:

```
claimable byte::stability() const
//implementaiton
{
    if (bit0) {} else {};
    if (bit1) {} else {};
    if (bit2) {} else {};
    if (bit3) {} else {};
    if (bit4) {} else {};
    if (bit5) {} else {};
    if (bit6) {} else {};
    if (bit7) {} else {};
}
```

the branches just demonstarate that a state of a byte determine the path(flow), and a stable byte would always behave the same.

we can also imagine a _parity_ stability, _initialzed_ stability, or other theoretical properties.

### Lifetime of a Right

right and duty are created together. but the right leaves the neighborhood, and is **returned** from it's creator, and is given to a higher (more abstract) neighborhood. the right can then be passed around, but the duty remains with it's creator. and eventually the right is returned to the creator, and then the right and duty are cancelled together. there is no expectation to stability.

```
void foo()
// implementation
{
    bool b = false();

    if (b)
        fail;

    b = true();

    if (b)
    {}
    else
        fail;

    end_lifetime(b);
}
```

when an object lifetime ends, the right of stability is returned to it, so there is no longer a duty for it to behave a certain way. this is also true for assignment operations, the right is reliqueshed in the proluge of the operation, and a different right (value) is created and returned, the state is changed, so the path inside the object is different,

```
bool & bool::operator (const bool r)
//interface
{
    claim_right this->stability();
    // implementation
    claim_right this->stability();
    // ...

    claim aliased(result, *this);
    claim result == r
}
```

but what is this "alias" part?

lets look at swapping, are there two rights? are we claiming the same right twice?

```
template <class T>
void swap_with_other(T& a, T & b)
//interface
{
    claim_right a->stability();
    claim_right b->stability();
    // implementation
    claim_right a->stability();
    claim_right b->stability();
}
```

a simple answer is saying the each claim is linked to a distinct right, so there are two rights being claimed. so this doesn't fit the situation of passing the same object. for that case we need something else:

```
template <class T>
void swap_with_other(T& a, T & b)
//interface
{
    claim_right a->stability();
    if (&a != &b)
        claim_right b->stability();
    else
        claim aliased(a,b); // see below

    // implementation

    claim_right a->stability();
    if (&a != &b)
        claim_right b->stability();
}

template <class T>
inline claimable aliased(T & a, T & b)
{
    entail &a == &b;
    entail substitutable(&a,&b);
}

template <class T>
inline claimable may_be_aliased(T & a, T & b)
{
    if (&a == &b)
        entail substitutable(&a,&b);
}
```

(there is a point about equability of pointers and refernecing objects)

back to our earlier example, there are three assignments in it when we assign the boolean value `true` to b.

```
void foo()
// implementation
{
    bool b = false();

    if (b)
        fail;

    //b = true();
    {
        auto r = true();
        b = r;
        end_lifetime(r);
    }

    if (b)
    {}
    else
        fail;

    end_lifetime(b);
}
```

so there is another right, the temporary value of 'true' which is created in the assignment operation.

### Immunity and Disability

so we now get two more terms

> - Disability: The responsability to refrain from chaning a realtionship.
> - Immunity: The ability to rely upon an unchanging relationship.

and their negations:

> - Power: The absence of disability, the lack of responsability to refrain from change.
> - Liability: The absence of immunity. the lack of an excpectation of constancy.

a police search warrent is a "power", it dispels a previously held right (the property right in the house), getting an injunction provides immunity from the warrent, so the right remains as it was. "Power" changes a relationship. "Immunity" and "Disability" prevent a relationship from changing.

> a disability begins when immunity is extended downward from a right.\
> the disability remains with the right holder - the immunity is extended and retracted.\
> the disability ends when the immunity is retraced upwards to its right.

the immunity is what moves around.

```
bool & bool::operator (const bool r)
//interface
{
    claim_right this->stability();
    discern_input r->value();

    // implementation

    claim_right this->stability();
    discern_output this->value();
    // ...

    claim aliased(result, *this);
    claim result == r
}

discernible bool::value() const
//interface
{
    claim_immunity stability();
    //implementation
}
```

### Claims as Assertions

and now for something real, two assertions (claims), but one has side effects, and the value is changed.

```
claim p = nulllptr; // bad assertion - assignment
claim p != nulllptr; // good assertion
```

but side effects is a very broad term, even wrting a log message is a side effect, not to mention that allocating memory is a side effect.

but the correct way to tell apart a good or bad assertions is to say that a good assertion does not effect the logic of the surronding code. this can be translated into the more precise terms of rights.

> - `claim p = nullptr;`
>   - consumens an external right.
>   - produces an unconsumed internal right.
> - `claim p != nullptr`;
>   - consumes no external rights.
>   - consumes all internally produced rights.

more examples

```
// bad assessments
claim (delete p), true; // consumes an external right
claim new int != nullptr; // produces an unconsumed internal right
claim p = nullptr; // consumes an external right, produces unconsumed internal right

// good assessments
claim (delete new int), true; // consumes no external rights and consumes all created internal rights
```

another example:

1. This right passes to the impelmentation.
   - The prologue may not consume it.
   - The prologue has immunity to it.
1. This right passes to the epilogue.
   - The implementation has no right to it.
   - The implementation has liability to it.
   - The epilogue must consume it
1. This right passes to the caller.
   - The epilogue may not consume it.
   - The epilogue has immunity to it.

```
int & int::operator++()
//interface
{

    claim_right this->stability(); //1
    discent_input this->value();

    int expected_result = *this+1; //2

    //implementation

    claim_right this->stability(); //3
    discent_output this->value();

    claim aliased(result, *this);
    claim result == expected_result;
}
```

> interfaced need not mention aliasing between immunities.

```
int & int::operator+(const int& a, const int &b)
//interface
{

    discern_input a.value();
    discern_input b.value();
    claim can_add(a,b);

    //implementation

    claim_right result.stability();
    discern_output result.value();
}
```

counter case,

> interfaces must mention aliasing between immunities and rights.

```
int & int::operator+=(int& a, const int &b)
//interface
{
    claim may_be_aliased(a,b);
    claim right a.stability();

    discern_input a.value();
    discern_input b.value();
    claim can_add(a,b);

    int expected_result = a+b;

    //implementation

    claim_right result.stability();
    discern_output result.value();

    claim aliased(result,a);
    claim result == expected_result;
}
```

we can remove the aliasing part if we decide that b is passed by value.

```
int & int::operator+=(int& a, const int b)
//interface
{
    claim right a.stability();

    discern_input a.value();
    discern_input b.value();

    int expected_result = a+b;

    //implementation

    claim_right result.stability();
    discern_output result.value();

    claim aliased(result,a);
    claim result == expected_result;
}
```

### Duties in Objects

- Constructor
- Mutator
- Accessor
- Destructor

duties flow, either from lowlevel code, or from class types.

the constructor produces a right of stability. it flows to the caller, but a mutator take the right a produce a different right. an accessor doesn't need the right of stability, as it has _immunity_. a destructor consumes that right of stability. the duty of stability is also passed around inside the objects. the constructor and destructor are connected by "the right of lifetime", this right is connected with the right of stability. there is also an immunity of lifetime.

> - the promise of stability depends on lifetime.
> - A neighborhood holding the _right of **stability**_ holds _immunity_ to **lifetime**.
> - A neighborhood holding the _duty of **stability**_ holds _immunity_ to **lifetime**.
> - A neighborhood holding the _immunity of **stability**_ holds _immunity_ to **lifetime**.
> - For the entire duration of a promise of **stability**, the promise of **lifetime** must not be cancelled.

```
claimable some_class::stability() const
//interface
{
    entail_immunity lifetime();

    //implementation
}
```

dependant rights, what happens in a function call, how rights,duties and immunities pass in function call. the outer right must always be in a higher level code from the inner right.

one case where this is clear is with threads.

> Launching a thread:
>
> - ➡ Provides _rights_ for the thread to consume.
> - ➡ Extend _immunities_ for the thread to rely upon.
> - ⬅ Receive the _right_ for the thread lifetime which depends on the _immunities_.
>
>   Thread runtime:
>
> - ▶ While the Thread is running, the disabilities are prolonged
> - ◼ After the thread completes, the disabilities are prolonged
>
> Joining a thread
>
> - ➡ Provide the _right_ of the thread lifetime, which is consumed.
> - ⬅ Receive any _rights_ produced by the thread.

Change is Ordinary, the normal way of things. Stability is extra-ordinary.

</details>

##

[Main](README.md)
