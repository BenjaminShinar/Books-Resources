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

moderns C++ has several smart pointer types.

1. _std::auto_ptr<>_ was **deprecated**. it was an attempt to standardize the pointer type at c++98 before we had move semantics, so it used work-around on it's copy operations. it had all sorts of issues.
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
