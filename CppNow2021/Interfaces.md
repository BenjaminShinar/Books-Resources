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
        if (dataCache_) //if optinal value exists
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
