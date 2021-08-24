## Chapter 4 - Smart Pointers

<summary>
aaa
</summary>

raw pointers aren't nice in a variety of ways:

1. the deceleration doesn't indicate whether it's pointing to a single object or an array.
2. the deceleration doesn't tell you if you should delete it after using it.
3. there's no way to know how to delete the pointer, maybe there's a dedicated call that you should use on pointers of that type before deleting them? was this pointer created with _malloc_ (and should be *free*d) or with _new_ and (_delete_).
4. even if you know that you should delete, we're back at point 1: _delete_ or _delete[]_? getting it wrong is bad.
5. it's difficult to be sure that we delete the pointer exactly once for each path along the code, and that we haven't double deleted it or missed one.
6. there's usually no good way to tell if a pointer is dangling (referring to a place in memory that was destroyed).

raw pointers are powerful, but we have decades of experience with them, and we know that they tend to be a source of problems,confusions and bugs. smart pointers are a way to address this issue. a smart pointer is a a wrapper around a raw pointer that has the same powers, but protects the user from the silent pitfalls like resource leaking or untimely destructions.

modern C++ has several smart pointer types.

1. _std::auto_ptr_ was **deprecated**. it was an attempt to standardize the pointer type at c++98 before we had move semantics, so it used work-around on it's copy operations. it had all sorts of issues.
   1. copying an _std::auto_ptr_ set it's value to null.
   2. it couldn't be stored in containers.
2. _std::unique_ptr_ is the functional version of the _std::auto_ptr_ idea. it does everything we wanted from it without performing weird twists around the idea of copying.
3. _std::shared_ptr_ - reference counting pointer
4. _std::weak_ptr_ - non owning reference

### Item 18: Use _std::unique_ptr_ for Exclusive-Ownership Resource Management

<details>
<summary>
std::unique_ptr is the go-to smart pointer to use in terms of performance and simplicity
</summary>
the default smart pointer we should use is *std::unique_ptr*, it's usually the same size as raw pointer, it behaves the same in nearly all cases, so they perform good even when memory is tight and performance is critical.

_std::unique_ptr_ embodies _exclusive ownership_ semantics. any non-null _std::unique_ptr_ owns what it points to, moving an _std::unique_ptr_ transfers ownership (the current pointer is set to _nullptr_), and copying is simply not allowed(that would violate the exclusive ownership idea), so it's a _move only type_.when the _std::unique_ptr_ reaches the end of it's lifetime, it deletes the held objects it contains (with _delete_, by default).

#### Factory Function

a common use case is as a factory function return type for objects in an hierarchy. in this example, Investment is a base class, while Stock,Bond and RealEstate are derived classes.

```cpp
class Investment{
    //some class
};
class Stock: public Investment{};
class Bond: public Investment{};
class RealEstate: public Investment{};
```

a factory function for this hierarchy would allocate an object on the heap and return a base class pointer to it, counting on the user to delete it at the end of the usage. this is perfect for _std::unique_ptr_. the ownership is transferred from the factory function to the user, and once the user is done with the objects, it's deleted immediately by the _std::unique_ptr_. we can use it either in single block or pass it to and from containers.

```cpp
template<typename... Ts> //variadic template arguments
std::unique_ptr<Investment> makeInvestment(Ts&&... params) //perfect forwarding?
{
    //..function body
}
{//block
    auto pInvestment = makeInvestment(arguments);
    //do stuff
    //pInvestment goes out of scope and is destroyed
}
```

the default destructor is the _delete_ operator, but we can also pass _custom deleters_ to the _std::unique_ptr_ constructor. this could be a different function, a pre-deletion logger, a lambda, a functor on something like that.

```cpp
auto delInvestment = [](Investment * pInvestment)
{
    makeLogEntry(pInvestment);
    delete pInvestment;
}; //custom deleter
template<typename... Ts>
std::unique_ptr<Investment, decltype(delInvestment)> makeInvestment(Ts&&... params)
{
    std::unique_ptr<Investment, decltype(delInvestment)> pInv(nullptr,delInvestment);
    if (/*make Stock object */)
    {
        pInv.reset(new Stock(std::forward<Ts>(params)...));
    }
    else if (/*make Bond object */)
    {
        pInv.reset(new Bond(std::forward<Ts>(params)...));
    }
    else if (/*make RealEstate object */)
    {
        pInv.reset(new RealEstate(std::forward<Ts>(params)...));
    }
    return pInv;
}
```

we use _decltype_ to grab the type of the custom deleter.a modern example would use auto, and maybe _std::make_unique_. from the users point of view, he doesn't need to concern himself with managing resources.

- _delInvestment_ is the custom deleter, it takes a raw pointer,does logging, and then calls _operator delete_.
- the type of _delInvestment_ must be declared, so we use _decltype_ when declaring the return type.
- we first create an empty _std::unique_ptr_, and then we assign to it (via reset) the appropriate class. we can't directly assign a pointer value to the _std::unique_ptr_ with the assignment operator.
- we use perfect forwarding to maintain the form of the arguments needed for the construction of the objects without copying them if we don't have do. this is the meaning of _std::forward_.
- the custom deleter takes a raw pointer of the base object, and calls _delete_ on it, so the base class must have a virtual destructor.

in c++14, we could have used the _auto_ keyword to avoid manually declaring the return type.

```cpp
template<typename... Ts>
auto makeInvestment(Ts&&... params)
{
    //.. same as before
}
```

when using the default deleter (_delete_), a _std::unique_ptr_ will usually have the same size as a raw pointer, but when a custom deleter is used, the size of the object grows. if we use a function object, the size increases, but using a sateless function object(lambda without captures) might avoid all that size penalty. so if possible, prefer a lambda as a custom deleter

```cpp
auto delInvestment1 = [](Investment * pInvestment)
{
    makeLogEntry(pInvestment);
    delete pInvestment;
}; //custom deleter stateless Lambda.
template<typename... Ts>
std::unique_ptr<Investment, decltype(delInvestment1)> makeInvestment(Ts&&... params);  // probably same size as raw pointer

void delInvestment1(Investment * pInvestment)
{
    makeLogEntry(pInvestment);
    delete pInvestment;
} // custom deleter as function
template<typename... Ts>
std::unique_ptr<Investment, void(*)(Investment *)> makeInvestment(Ts&&... params);  // return types is probably larger now, because it stores a function pointer!
```

we can also use _std::unique_ptr_ to implement the PIMPL idiom(_pointer to implementation_), but that's covered in [Item 22]().

we can use _std::unique_ptr_ for both single objects _std::unique_ptr\<T>_ and for arrays _std::unique_ptr\<T[]>_ and we will know which kind of pointer we have (and the correct deleter will be be called), so if we have an array, we could use indexing (_operator[]_), but not dereferencing (_operator\* and operator->_), and vice versa, although, we should have any use _std::unique_ptr_ of arrays, as better data structures exists in modern c++.

besides expressing exclusive ownership, we can also easily convert _std::unique_ptr_ into a shared ownership policy _std::shared_ptr_. that's why a factory function should return the strongest smart*pointer type, and the user could then convert it to a \_std::shared_ptr* if needed

```cpp
std::shared_ptr<Investment> sp = makeInvestment(arguments); //convert unique_ptr into shared_ptr.
```

#### Things to Remember

> - std::unique_ptr is a small, fast, move-only smart pointer for managing
>   resources with exclusive-ownership semantics
> - By default, resource destruction takes place via delete, but custom deleters can be specified. Stateful deleters and function pointers as deleters increase the size of std::unique_ptr objects.
> - Converting a std::unique_ptr to a std::shared_ptr is easy.

</details>

### Item 19: Use _std::shared_ptr_ for Shared-Ownership Resource Management

<details>
<summary>
std::shared_ptr allows for several objects to use the same object in memory, without having to manually coordinate between them.
</summary>
Garbage collection programming languages exists since Lisp back in the 1960's. however, C++ can't settle for that version of garbage collection - the timing of resource reclamation is non deterministic. in other languages, garbage collection happens when it happens, without prior control, and sometimes, it's just not good enough.

_std::shared_ptr_ is a way of trying to get the best of both approaches: have an easy and simple way to reclaim resources like a garbage collection language, but have it happened at known times, like manual resources allocation. _std::shared_ptr_ does this by the concept of _shared ownership_, no single _std::shared_ptr_ owns the resource, each time a _std::shared_ptr_ is destroyed, it checks if it's the last user of the resource, and if it is, then he deletes the resource. the process is automatic (no need to manually check and delete), but can happen only in deterministic times (a _std::shared_ptr_ goes out of scope).

the _std::shared_ptr_ employ reference counting, constructing a _std::shared_ptr_ (usually) increases the count, destructing decreases it, copy and assignment operations can decrease one and increase another. this does have performance costs, both in memory (storing the extra data) and in runtime (checking).

- _std::shared_ptr_ are usually twice the size of the raw pointer. there is extra data to store.
- the data for the memory count is dynamically allocated. the _std::shared_ptr_ can point to different objects, and there's no way to control this before runtime, sometimes using _std::make_shared_ utility function can help us avoid parts of this dynamic allocation, but it's can't always be used.
- increments and decrements of the reference count must be atomic, the same resource can be used by multiple readers/writers, and it's critical that the count be exact (to avoid an issue of resource leaking), so the costly check must be performed.

but wait, why does constructing a _std::shared_ptr_ only _usually_ increase the reference count? because of move constructors, that's why. moving from one _std::shared_ptr_ to another doesn't require all those manipulations of the reference count, so it's usually faster to move from one than to copy. same goes for move constructors over copy constructors.

just like _std::unique_ptr_, _std::shared_ptr_ also supports custom deleters, however, in this case, the custom delete is part of the pointed object, rather than of the _std::shared_ptr_ objects.
this design is more flexible, we can place _std::shared_ptr_ of the for the same type with different custom deleters (it's not part of the type).

```cpp
auto loggingDeleter = [](Widget *pw)
{
    makeLogEntry(pw);
    delete(pw);
}; //custom stateless deleter

std::unique_ptr<Widget,decltype(loggingDeleter)> upw(new Widget, loggingDeleter); //deleter is part of the pointer
std::shared_ptr<Widget> spw(new Widget,loggingDeleter); //deleter is not part of the pointer

auto customDeleter1 =[](Widget *pw){/*...*/};
auto customDeleter2 =[](Widget *pw){/*...*/};

std::shared_ptr<Widget> pw1(new Widget,customDeleter1);
std::shared_ptr<Widget> pw2(new Widget,customDeleter2);

std::vector<std::shared_ptr<Widget>> vpw {pw1,pw2}; // vector contains shared_ptr with different deleters
```

another difference is that adding a custom deleter doesn't change the size of _std::shared_ptr_, it's always the size of two pointers. where is the rest of the memory stored? it's on the heap, but doesn't belong to the _std::shared_ptr_ object. the _std::shared_ptr_ object contains a pointer to the data, and a pointer to the _control block_, that's where we have the reference count and where we store the custom deleter, custom allocator, and maybe even a weak reference count ([Item 21]()).

#### The Control Block

this control block is created by the following rules:

- the function **std::make_shared** always created a control block. it manufactures a new object to point to, so it must also create the control block.
- if the _std::shared_ptr_ is created from a unique ownership smart pointer (_std::unique_ptr_ or the deprecated _std::auto_ptr_), then a control block is created, and ownership is transferred to the _std::shared_ptr_ object (the other object is set to _nullptr_)..
- if the _std::shared_ptr_ was created with a raw pointer, it creates a control block. if we had a shared-ownership object already, we would have used it to create our new _std::shared_ptr_.

as a result of these rules, if we decide to take the raw pointer of a _std::shared_ptr_ and use it to construct a different _std::shared_ptr_, then we are setting ourself into a party of undefined behavior. after all, we have to independent control blocks who think they manage the same data.

```cpp
auto pw = new Widget;
std::shared_ptr<Widget> spw1 (pw, loggingDeleter); //this creates a control block
std::shared_ptr<Widget> spw2 (pw, loggingDeleter); // and so does this.
```

in a general matter, we shouldn't create raw pointers on the stack anymore , that's what smart pointers are for. but in the above code, the problem is that we have two control blocks that manage this same data. this means two reference counts, two calls to the loggingDeleter, and eventually, undefined behavior.

in this case we must use the naked creation (std::make*shared doesn't play nice with custom deleters yet). but we should construct the object as part of the \_std::shared_ptr* creation.

```cpp
std::shared_ptr<Widget> spw1 (new Widget, loggingDeleter); //this creates a control block
std::shared_ptr<Widget> spw2 (spw1); // control block already exists.
```

unfortunately, there are other ways to get to this bad behavior. in the following example, we use emplace*back to create a \_shared_ptr* inside the vector, but this means calling the _shared_ptr_ constructor with a raw pointer, so we create an additional control block.

```cpp
std::vector<std::shard_ptr<Widget>> processedWidgets;
class Widget
{
    public:
    void process()
    {
        //.. do something
        processedWidgets.emplace_back(this); // add this to the processedWidgets vector. bad! wrong! Error!
    }
};

shared_ptr<Widget> spw = std::make_shared(new Widget); //one control block
spw->process(); //oops this means another control block.
```

if we wish to allow this behavior, we can use a special base class that allows this behavior **std::enable_shared_from_this\<>**, which uses CRTP(_curiously recurring template pattern_).this allows us to refer to a shared control block, but only if one already exists. if there isn't a control block, we are back to undefined behavior.

```cpp
class Widget: public std::enable_shared_from_this<Widget>
{
    public:
    void process()
    {
        //.. do something
        processedWidgets.emplace_back(shared_from_this()); // only if there is a shared_ptr outside for this objects
    }
};
shared_ptr<Widget> spw = std::make_shared(new Widget); //one control block
spw->process(); //cool, this works.
```

because of the issue with undefined behavior, if we use this base class, we should probably keep our constructors private and use factory functions to construct our objects

```cpp
class Widget: public std::enable_shared_from_this<Widget>
{
    public:
    template<typename... Ts>
    static std::shared_ptr<Widget> create(Ts&&... params); //factory function
    void process()
    {
        //.. do something
        processedWidgets.emplace_back(shared_from_this()); // only if there is a shared_ptr outside for this objects
    }
    private:
    //ctor
};
```

regardless of when it's created, the control block is stored on the heap, and it's usually several words in size (custom allocators and deleters might make it larger), it uses inheritance, virtual functions and other bits of complicated programming. in most cases, when there is no custom allocator or deleter, the size is mostly three words in size,dereferencing it doesn't cost more than dereferencing a regular pointer, and the reference manipulating operations are usually blazing fast and mapped to specific machine instructions.
it's easy to move from single-ownership models (_std::unique_ptr_) to shared-ownership, but there's no way to move back once an object is under shared-ownership.

another current issue with _std::share_ptr_ is their relationship with arrays. there is no _std::share_ptr\<T[]>_. even if we have a custom deleter with _delete[]_ operation, the _std::share_ptr_ doesn't offer the _operator[]_ for indexing, and besides, we have std::vector, std::array for that.

#### Things to Remember

> - *std::shared_ptr*s offer convenience approaching that of garbage collection for the shared lifetime management of arbitrary resources.
> - Compared to _std::unique_ptr_, _std::shared_ptr_ objects are typically twice as big, incur overhead for control blocks, and require atomic reference count manipulations.
> - Default resource destruction is via delete, but custom deleters are supported. The type of the deleter has no effect on the type of the _std::shared_ptr_.
> - Avoid creating *std::shared_ptr*s from variables of raw pointer type.

</details>

### Item 20: Use _std::weak_ptr_ for _std::shared_ptr_-like Pointers That Can Dangle

<details>
<summary>
Non-Owning smart pointers that can detect dangling data.
</summary>
_std::weak_ptr_ behaves like *std::shared_ptr*, but doesn't count towards the ownership count. this is intended to tackle the possibility that an object might be destroyed. the _std::weak_ptr_ is for these situations. it can't be dereferenced or checked for null (but can be checked for expiry), as it's actually an augmentation of the *std::shared_ptr* class.   
_std::weak_ptr_  are created from *std::shared_ptr*s, they point to the same control block as the *std::shared_ptr* that created them, but they don't affect the reference count.

```cpp
auto spw = std::make_shared<Widget>(); //reference count is 1
std::weak_ptr<Widget> wpw(spw); // reference count is still 1
spw=nulltpr; // or reset(nullptr), reference count is 0

if (!wpw.expired()) //false. it has expired
{
    //do something
}
```

even if we can check our _std::weak_ptr_ isn't dangling with _.expired()_, we still can't do anything with it. there's no dereferencing from _std::weak_ptr_. and even if there were, they could be a data race, what if the last owning _std::shared_ptr_ goes out of scope and destroys the data? what we need (and want, and have) is an atomic operation that checks the expiry status, and if its available, gives us an owning smart pointer to use the data with.

this operation has two forms in c++, we can either get a nullptr if the _std::weak_ptr_ is dangling by using the _lock()_ method, or an exception if we try to construct a *std::shared_ptr *from a _std::weak_ptr_ .

```cpp
auto spw = std::make_shared<Widget>();
std::weak_ptr<Widget> wpw(spw);
spw = nullptr;
auto spw1 = wpw.lock(); //creates a shared_ptr
if (spw1 != nullptr)
{
    //do something
}
try
{
auto spw2 = std::shared_ptr<Widget>(wpw);
//do something
}
catch(std::bad_weak_ptr & e) //exception
{

}
```

in terms of efficiency and size, _std::weak_ptr_ are about the same as _std::shared_ptr_, and they point to the same control block, it's just that _std::weak_ptr_ don't participate in the **shared-ownership count**, but they do participate in the other reference count in the control block.

#### Cache

a possible use case for _std::weak_ptr_ is for caching results. if we expect something to be used repeatedly and creation is costly (IO, database access), maybe it's better to keep the results in memory and use them again if possible. maybe we already have a shared_ptr to it somewhere.

```cpp
std::unique_ptr<const Widget> loadWidget(WidgetId id); //factory for unique widgets, but this is step 1

std::shared_ptr<const Widget> fastLoadWidget(WidgetId id)
{
    static std::unordered_map<WidgetId, std::weak_ptr<const Widget>> cache;
    //check if cache contains key at all...
    auto objPtr = cache[id].lock(); //shared_ptr, null if expired
    if (!objPtr)
    {
        objPtr = LoadWidget(id); // do the costly thin
        cache[id]=objPtr; //store a weak_ptr, created from shared_ptr.
    }
    return objPtr;

}
```

#### Observer Design Pattern

a different use case is the _Observer design pattern_. we have a _subject_, whose state can change, and _observers_, who wish to be informed when the subject change. we usually store all the observers in a list inside the subject memory, but we actually have no intention of controlling the lifetime of the observers via the the subject, but we don't want to access an observer that has been already destroyed. _std::weak_ptr_ allows us to do so, we can simply check if the object is dangling before accessing it.

```cpp
class Subject
{
    public:
    void stateChange()
    {
        for (auto & weakPtr  : observers)
        {
            auto sharedPtr = weakPtr.lock();
            if (sharedPtr)
            {
                //do something with the observer, now that it's a shared_ptr
            }
        }
    }
    private:
    std::list<std::weak_ptr<Observer>> observers;
}
```

#### Circular Reference

a final use case is for Circular references. if there are three elements A,B,C: A and C have shared ownership over B.

- A holds a shared_ptr of B
- C holds a shared_ptr of B
- B wants to access A, but how should he do it?

there are three options

- raw pointer - but if A is destroyed, B won't know about it. not good.
- _std::shared_ptr_ - an ownership cycle. if A goes out of scope, it still lives as a member of B. if B goes out of scope, it's still kept alive by A. even if both are out of scope their reference count is 1, they keep each other alive by the virtue of circular reference, their resources won't be reclaimed.
- _std::weak_ptr_ - this is the preferred solution, B can detect if A goes out of scope, and it doesn't extend A's lifetime.

in most hierarchical data structures, we don't expect children elements to outlive their parents, so there is usually no need for _std::weak_ptr_, but it doesn't hurt.

#### Things to Remember

> - Use _std::weak_ptr_ for _std::shared_ptr_-like pointers that can dangle.
> - Potential use cases for _std::weak_ptr_ include caching, observer lists, and the
>   prevention of _std::shared_ptr_ cycles.

</details>

### Item 21: Prefer _std::make_unique_ and _std::make_shared_ to Direct Use of _new_

<details>
<summary>
The utility functions usually offer benefits in performance and safety, use them unless there is specific reason not to.
</summary>

_std::make_shared_ exists in c++11, _std::make_unique_ exists in c++14, but even in c++11, it's simple to write, we just need to perfect forward the parameters to the constructor (recall that
we use the parentheses constructor, see [item 7]()).

```cpp
template<typename T, typename... Ts>
std::unique_ptr<T> make_unique(Ts&&... params)
{
    return std::unique_ptr<T>(new T(std::forward<Ts>(params)...));
}
```

as we can see, this form doesn't support custom deleters or creating unique\*ptr to arrays, but it will work for most cases.
_std::make_shared_ ,_std::make_unique_ are two of three _make_ functions that take an arbitrary set of arguments (the third function is _std::allocate_shared_, which acts like _std::make_shared_ but takes a allocator object as well).

in most cases, the difference in typing between the two forms is negligible, or even easier with the custom functions. we can still use auto, and we only write the typename once.

```cpp
auto upw1(std::make_unique<Widget>);
std::unique_ptr<Widget> upw2(new Widget);


auto spw1(std::make_shared<Widget>);
std::shared_ptr<Widget> spw2(new Widget);
```

#### Avoiding Resource Leaks

a bigger reason to use the make functions is to control exception safety. we have no problem with passing by value, the copy constructor will make a copy, which is perfectly fine fo _std::shared_ptr_. lets assume a function takes a _std::shared_ptr_ and another value, if we construct the _std::shared_ptr_ on spot, the compiler can play tricks on us and cause resource leaks because of the order of arguments evaluation (which is undefined).

```cpp
void processWidget(std::shared_ptr<Widget> spw,int priority);
int computePriority(); // return priority
processWidget(std::shared_ptr<Widget>(new Widget),computePriority()); // this is potential resources leak.
```

we must create the new Widget before calling the _std::shared_ptr_ constructor, but it's possible that between those two operations, the computePriority function threw an exception,

> possible order that causes leak
>
> 1. perform _"new Widget"_
> 2. execute _computePriority()_ - **but what if we had an exception here?**
> 3. ~~run the _std::shared_ptr_ constructor~~. **didn't happen, leaked resource**

had we used the utility functions, the new resource would have definitely been stored inside the smart pointer, and we wouldn't see the resource leak.

```cpp
processWidget(std::make_shared<Widget>(),computePriority()); // this is safe.
```

#### Better Performance

another bonus of using the utility functions is about the number and location of memory allocations. calling the constructor of a _std::shared_ptr_ creates the control block on the heap, and calling the new operator creates the data, also on the heap. by combining both calls into _std::make_shared_ the compiler can request one continues block of data of the size of the data and the control block combined, and place both of them together, which means better data locality. this is also true for _std::make_allocate_.

#### Edge Cases

despite the advantages of the make functions, there are still cases when its impossible to use them,there are two cases that happen with both _std::unique_ptr_ and _std::shared_ptr_, and two more cases that apply only to _std::shared_ptr_

for both kinds, we cannot use the utility function for passing a custom deleter.

```cpp
auto widgetDeleter=[](Widget* pw){...}; //deleter
std::unique_ptr<Widget,decltype(widgetDeleter)> upw(new Widget,widgetDeleter);
std::shared_ptr<Widget> spw(new Widget, widgetDeleter);
```

another case is the usage of the braced initialization (curly braces) in the constructor call. as we saw earlier, depending on how the make function ins implemented, we can get very different results for the following:

```cpp
auto upv = std::make_unique<std::vector<int>>(10,20);
```

one option uses the parentheses constructor, and will result in a vector with 10 elements with the value 20.
the other option uses uniform initialization (with std::initializer_list) and will result in a vector of two elements, 10 and 20.

the result is unambiguous, as part of the documentation, the decision was made to use parentheses over braces. but that means there is no way to use curly braces in make functions, and we must resort to calling _new_. in [item 30]() there is a work around for this.

```cpp
auto initList = {10,20}; //std::initializer_list.
auto spv = std::make_shared<std::vector<int>>(initList);
```

if our class defines it's own versions of _operator new_ and _operator delete_,we shouldn't use the utility functions to create it (we can still use the _std::shared_ptr_, just not the _std::make_shared_ function). also, there is a possibility that if we allocated the control block together with the data, and the control block contains weak reference count, the memory won't be released until all *std::weak_ptr*s go out of scope. those _std::weak_ptr_ are able to extend the lifetime of the object, even if it's no longer accessible.
if we allocate the memory in two different calls, then we can reclaim the object memory separately from that of the control block.

if we want to avoid the resource leaking issue, we need to make sure we construct the _std::shared_ptr_ immediately with the new object, without any other statements in the same time. now we start tweaking the call for better performance, first by ensuring a move is done rather than a copy (avoiding usage of atomics)

```cpp
void processWidget(std::shared_ptr<Widget> spw, int priority);
void deleter(Widget * ptr); //custom deleter
processWidget(std::shared_ptr<Widget>(new Widget,deleter),computePriority()); //unsafe
std::shared_ptr<Widget> spw(new Widget,deleter); //lvalue
processWidget(spw,computePriority()); //safe, but not optimal, we ensure copying, it would be better to allow moving, won't it?,
processWidget(std::move(spw), computePriority()) //safe, correct, and also better.
```

#### Things To Remember

> - Compared to direct use of new, make functions eliminate source code duplication, improve exception safety, and, for _std::make_shared_ and _std::allocate_shared_, generate code that’s smaller and faster.
> - Situations where use of make functions is inappropriate include the need to
>   specify custom deleters and a desire to pass braced initializers.
> - For *std::shared_ptr*s, additional situations where make functions may be ill-advised include (1) classes with custom memory management and (2) systems with memory concerns, very large objects, and *std::weak_ptr*s that outlive the corresponding *std::shared_ptr*s.

</details>

### Item 22: When Using the Pimpl Idiom, Define Special Member Functions in the Implementation File

<details>
<summary>
We can use the smart pointers to use the pimpl idiom, but it does require some work to ensure that the classes are incomplete in the header file and still have all the required member function such as a destructor or move/copy operations.
</summary>

The pimpl (_pointer to implementation_) idiom is a technique where rather than define data members as concrete objects, we declare them as pointer to a class/struct, and leave the definition for later.

first, the non idiom way. widget is defined in the header of widget.h, and because it has Gadget members, it should include the gadget header, and should be re-compiled whenever the Gadget class changes. any class that uses Widget is also required to re-compile when Gadget changes.

```cpp
class Widget{
public:
    Widget();
    //...
private:
std::string name;
std::vector<double> data;
Gadget g1,g2,g3;
};
```

in c++98, the PIMPL idiom would look like this. we use _incomplete types_ we no longer have headers for any other classes in the header, so changing one of those class doesn't require users of the Widget class to re-compile. all the data is encapsulated in the cpp file.

```cpp
// the widget.h header
class Widget{
public:
Widget();
~Widget();  //we need the destructor, probably also the other functions to complete the rule of three
//...
private:
struct Impl; //declare an implementation struct and a pointer to it
Impl *pImpl;
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(new Impl) //constructor, allocate data members for the implementation class
{

}
Widget::~Widget() //destructor, release data.
{
    delete pImpl;
}
```

in a more modern world, we would use smart pointers, this way we don't need to define the destructor. our code compiles, but somehow **simple client code doesn't**? we see an issue of incomplete type, the delete operator and sizeof, why is that?

```cpp
// the widget.h header
class Widget{
public:
Widget();

//...
private:
struct Impl; //declare an implementation struct
std::unique_ptr<Impl> pImpl; // use a smart pointer
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(std::make_unique<Impl>())) //constructor, use the utility function
{

}

//client code.
#include <widget.h>

Widget w; //error!
```

lets understand the issue.
we have a _std::unique_ptr_ to an incomplete type, we didn't define a custom deleter, so the default is used. we didn't define a destructor so we got one for free from the compiler, and as all member generated functions, it's assumed to be _implicitly inline_. so we have an inline destructor without any knowledge about the incomplete type it should destroy.
we can fix it by declaring our own, non inlined destructor in the .cpp file, where the Impl type is known and no longer incomplete. we can also use the _=default_ in the definition.

```cpp
// the widget.h header
class Widget{
public:
Widget();
~Widget(); //turns out we do need the destructor
//...
private:
struct Impl; //declare an implementation struct
std::unique_ptr<Impl> pImpl; // use a smart pointer
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(std::make_unique<Impl>())) //constructor, use the utility function
{

}
Widget::~Widget() //destructor definition
{
}
//Widget::~Widget() = default; also possible

//client code.
#include <widget.h>

Widget w; //now it's fine!
```

if we want move support (and we probably do, as we already have unique_ptr which is great for move semantics), we need to do something similar, we declare the operations in the .h file, but define them in .cpp file. their definitions can be defaulted.

```cpp
// the widget.h header
class Widget{
public:
Widget();
~Widget(); //turns out we do need the destructor
Widget(Widget && rhs); //declaration of move constructor
Widget & operator=(Widget&& rhs); //declaration of move assignment

//...
private:
struct Impl; //declare an implementation struct
std::unique_ptr<Impl> pImpl; // use a smart pointer
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(std::make_unique<Impl>())) //constructor, use the utility function
{

}
Widget::~Widget() = default;
WWidget::Widget(Widget && rhs)= default;
Widget::Widget & operator=(Widget&& rhs) = default;
```

what's missing? oh, right. Copy operations, once we defined move operations we no longer have copy operations generated for us, and even if we did, we can't copy _std::unique_ptr_, so we need to do it ourselves, luckily, we can use the default generated operations for the Impl data class.

```cpp
// the widget.h header
class Widget{
public:
Widget();
~Widget(); //turns out we do need the destructor
Widget(Widget && rhs); //declaration of move constructor
Widget & operator=(Widget&& rhs); //declaration of move assignment
Widget(const Widget & rhs);
Widget& operator=(const Widget & rhs)
//...
private:
struct Impl; //declare an implementation struct
std::unique_ptr<Impl> pImpl; // use a smart pointer
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(std::make_unique<Impl>())) //constructor, use the utility function
{}
Widget::~Widget() = default;
WWidget::Widget(Widget && rhs)= default;
Widget::Widget & operator=(Widget&& rhs) = default;
Widget::Widget(const Widget & rhs):pImpl(std::make_unique<Impl>(*rhs.pImpl))
{}
// create a new unique_ptr with an object that is the copy of the one in the other objects

Widget::Widget& operator=(const Widget & rhs)
{
    *pImpl = *rhs.Impl; //copy assignment of the data.
    return *this;
}
```

if we had decided to use a _std::shared_ptr_ to store our implementation data, things would be different, we wouldn't need as much code. we wouldn't need to declare any member functions in the header and then default it the cpp file. this is because _std::unique_ptr_ contain the custom deleter within them `std::unique_pre<T, decltype(customDeleter)>`, so everything must be known and we must have complete typing when created. on the other hand _std::shared_ptr_ stores the custom deleter on the heap in the control block, so it can be much more lenient with incomplete type.

```cpp
class Widget {
public:
Widget();
//... more code
private:
struct Impl;
std::shared_ptr<Impl> pImpl
};

// the widget.cpp file
#include <widget.h>
#include <gadget.h>
struct Widget::Impl
{
    std::string name;
    std::vector<double> data;
    Gadget g1,g2,g3;
}

Widget::Widget():pImpl(std::make_unique<Impl>())) //constructor, use the utility function
{}

//client code

Widget W1;
auto w2(std::move(w1));
w1 = std::move(w2);
```

in most cases, the std::unique_ptr is the correct choice for the Pimpl Idiom, as it most closely resembles the base form.

#### Things to Remember

> - The Pimpl Idiom decreases build times by reducing compilation dependencies
>   between class clients and class implementations.
> - For _std::unique_ptr_ pImpl pointers, declare special member functions in
>   the class header, but implement them in the implementation file. Do this even
>   if the default function implementations are acceptable.
> - The above advice applies to _std::unique_ptr_, but not to _std::shared_ptr_.

</details>
