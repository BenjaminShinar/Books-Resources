## Chapter 5 - Rvalue References, Move Semantics, and Perfect Forwarding

<summary>
The usage conventions of std::move and std::forward, the ambiguous nature of "type&&" parameters.
</summary>

Move semantics and and perfect forwarding as they seem at first glance:

- _Move semantics_ allow the compilers to replace expensive copy operations with move operations, the same we can define what copying means for our class (deep vs shallow copy), we can define what move means for our class. move semantics are also what powers move-only types: _std::unique_ptr_, _std::thread_ and _std::future_.
- _Perfect forwarding_ allow function templates that receive arbitrary arguments pass them forward to other functions without effecting their type.

the whole concept is glued together by the idea of _rvalue references_, which make both move semantics and perfect forwarding possible. **But things aren't as simple as they appear**. everything is nuanced, _std::move_ doesn't really move anything, perfect forwarding is imperfect, move operations aren't always cheaper and aren't always used, and _type&&_ doesn't always represent an _rvalue reference_.

a quick note to remember, **even if the parameter type is a rvalue, the parameter itself is an lvalue**.

```cpp
void f(Widget && w);
```

while the type of the parameter is a rvalue, the parameter itself is an lvalue, there is a memory address for it.

### Item 23: Understand _std::move_ and _std::forward_

<details>
<summary>
std::move and std::forward are casts to an rvalue. that's all.
</summary>

despite their names, neither _std::move_ or _std::forward_ actually do what they say. _std::move_ doesn't move and _std::forward_ doesn't forward. they actually don't do anything, they don't generate executable code. what they actually do is perform _casting_. _std::move_ **always** casts the value into an rvalue form, while _std::forward_ **might** cast the value into an rvalue, depending on some conditions.

here is a simple implementation of _std::move_, it's not according to standard, but it's enough for us now.

```cpp
//assume we are inside namespace std
template <typename T>
typename remove_reference<T>::type&& //return type
move(T&& param)
{
    using ReturnType = typename remove_reference<T>::type&&; //alias declaration
    return static_cast<ReturnType>(param);
}
```

the parameter type is a universal reference (see [Item 24]()), and the return type is a reference.
if the argument type (T) was an lvalue reference, then casting would return an lvalue reference, for that reason, we strip the reference from T(_std::remove_reference\<T>_), take the type(_::type_), and then say it's a rvalue reference (_&&_).

in c++14, things are easier to write, thanks to auto

```cpp
template <typename T>
decltype<auto> move(T&& param)
{
    using ReturnType=remove_reference_t<T>&&;
    return static_cast<ReturnType>(param);
}
```

to reiterate. _std::move_ **doesn't move, it casts**. maybe they could have chosen a better name. but we have _std::move_ as the name. rvalues are candidates for moving, so applying _std::move_ tells the compiler that it should move from this object, and that's the reason for the name. _std::move_ tells the compiler that it can move from this object. simple, right?
except that rvalues are only _usually_ candidates for moving.

here is an example. we have the Annotation class, which takes a text in it's constructor.
we go thorough several iterations to get the best results.
we first use copy by value, and then we decide that actually we should put const on the argument, because c++ loves const. and then we try moving from the parameter. and things work properly, but we actually keep copying, not moving.

```cpp
class Annotation{
public:
//explicit Annotation(std::string text): value(text){} // copy by value
//explicit Annotation(const std::string text): value(text){}; // copy by value, without modifying
explicit Annotation(const std::string text): value(std::move(text)){} // doesn't do what we think!
private:
std::string value;
};
```

the reason is that move constructors don't accept const. after all, moving from an object can (and should) modify it. and we can't modify const values.

```cpp
class Widget
{
    Widget(const Widget& other); //copy constructor
    Widget(Widget&& other); //move constructor, no const
};
```

there are two lessons to be learned

1. don't declare objects const if you plan to move from them. move operations are silently transformed into copy operations.
2. _std::move_ doesn't guarantee move operations, it just makes the object an rvalue.

#### _std::forward_ as a Conditional Cast to Rvalue

the story is similar for _std::forward_, only that _std::forward_ casts to rvalue only under some conditions. the common use case for _std::forward_ is passing a parameter that was taken as a universal reference to a different function.

```cpp
void process(const Widget& lvalueArg);
void process(Widget && rvalueArg);
template <typename T>
void logAndProcess(T&& param) //universal reference
{
    auto now = std::chrono::system_clock::now();
    makeLogEntry("calling `process`",now);
    process(std::forward<T>(param));
}

Widget w;
logAndProcess(w); //call with lvalue
logAndProcess(std::move(w)); //call with rvalue
```

inside the logAndProcess function, even when the function is called with an rvalue argument, the parameter itself is an lvalue. all parameters are lvalues. _std::forward_ is a conditional cast that checks if the parameter was initialized with a rvalue, and if it was, it performs the cast. this information is part of the template T parameter.

in theory, _std::forward_ could be used wherever _std::move_ is used, but it would require more typing (we need to pass the non reference type), but more important, both functions have different usages, so it's nice that we can disguise between them.

#### Things to Remember

> - _std::move_ performs an unconditional cast to an rvalue. In and of itself, it
>   doesn’t move anything.
> - _std::forward_ casts its argument to an rvalue only if that argument is bound
>   to an rvalue.
> - Neither _std::move_ nor _std::forward_ do anything at runtime.

</details>

### Item 24: Distinguish Universal References from Rvalue References

<details>
<summary>
The && symbol can mean either a rvalue reference or universal reference. The two constraints to determine if it's an universal reference are the presence of type deduction and the form (syntax). otherwise, it's a rvalue reference.
</summary>

the symbol **&&** has two meanings, one is rvalue reference, something that binds only to rvalues, and indicates a value that can be _moved_ from. the other meaning is that something is either a lvalue or a rvalue reference. under this meaning, this can be an lvalue,rvalue, const or non const,volatile and any combination of those. this is the _universal reference_.

let's start with some examples:

```cpp
void f1(Widget&& param); //rvalue reference

Widget && var1 = Widget(); //rvalue reference

auto && var2= var1; // not an rvalue reference, var is a universal reference

template <typename T>
void f2(std::vector<T>&&param); //rvalue reference

template <typename T>
void f3(T&& param); // no an rvalue reference, param is a universal reference
```

an important thing to note is that universal references often go together with *std::forward\*\*, so they are sometimes called *forwarding reference\* instead..

we see universal references in two contexts, one in inside function template parameters, and the other is _auto_ decelerations. the common core of those two situations is _type deduction_. in the template example, the type of param is being deducted, and the _auto_ keyword is all about type deduction. so if we see **&&** without type deduction it's an rvalue reference. if there is type deduction, the declaration must also follow the correct form of **T&&**.

```cpp
void f1(Widget&& param); //rvalue reference, no Type deduction needed.

Widget && var1 = Widget(); //rvalue reference, no Type deduction needed.

template <typename T>
void f2(std::vector<T>&&param); //rvalue reference, type deduction but un proper

auto && var2= var1; // not an rvalue reference, var is a universal reference, auto is deducing type.

template <typename T>
void f3(T&& param); // no an rvalue reference, param is a universal reference
```

universal reference are references, and therefore, must be initialized, the initializer determines if the universal reference represents an rvalue reference or an lvalue reference. when the universal reference are function parameters, the initializer is provided at the call site.

```cpp
template <typename T>
void f(T&& param); // param is universal reference.

Widget w;
f(w); //passing an lvalue to the function, param type is an lvalue reference Widget &
f(std::move(w)); // casting w into rvalue reference, param type is rvalue reference Widget &&.
```

other than type deduction, there is another condition necessary for **&&** to represent an universal reference. that is the _form_ of the declaration, it must be **T&&**, nothing else is enough. in the following example we have type deduction but improper form. even a _const_ is enough to prevent something from being a universal reference.

```cpp
template <typename T>
void f(std::vector<T>&& param); // param is rvalue reference.

std::vector<int> v;
f(v); //error! can't bind lvalue to rvalue reference!

template <typename T>
void g(const T&& param); //param is n rvalue reference
g(v); // same!
```

even the presence of the parameter type **T&&** isn't enough to ensure something is a universal reference, that's because templates by themselves don't guarantee type deduction.

in the std::vector class, there is the _push_back_ function, which seems to qualify for universal reference in form, but actually doesn't do any type deduction. in contrast, the _emplace_back_ function does employ type deduction.

```cpp
template <class T, class Allocator = allocator<T>>
class vector{
    public:
    void push_back(T&& x); //rvalue reference actually, despite the form.
    template <class Args>
    void emplace_back(Args&&...args); // universal reference.
    //...
};
```

the reason is although the std::vector uses type deduction, the _push_back_ function doesn't. by the time the compiler generates the function, the required type is fully known. contrary to that, the _emplace_back_ function requires type deduction at the call site, so it's a universal reference

```cpp
class vector<Widget, allocator<Widget>>
{
    public:
    void push_back(Widget && x); //no type deduction!

    template <class Args>
    void emplace_back(Args&&...args); // universal reference. we know that args are required to create a Widget, but not what they are.
    //...
};
```

auto variables can be universal references, or to be precise, variables declared with _auto_ and _&&_ are universal references. those cases aren't as common as universal references with template parameters in c++11, but they are more common in c++14, with lambda expressions.

```cpp
auto timeFuncInvocation =[](auto && func, auto &&...params)
{
    //start timer;
    std::forward<decltype(func)>
    (func)(std::forward<decltype(params)>(params)...)); //invoke func on params
    //stop timer and record elapsed time;
}
```

the code will become more clear in [Item 33](). the important factor for now is that the _auto &&_ parameters (both func and params) are universal references. also note the the entire concept of universal references is an abstraction of [Reference Collapsing](). the benefits of the distinction between rvalue references and universal references are understanding exactly the parameters and being able to properly communicate with others.

#### Things to Remember

> - If a function template parameter has type T&& for a deduced type T, or if an
>   object is declared using auto&&, the parameter or object is a universal reference.
> - If the form of the type declaration isn’t precisely type&&, or if type deduction
>   does not occur, type&& denotes an rvalue reference.
> - Universal references correspond to rvalue references if they’re initialized with
>   rvalues. They correspond to lvalue references if they’re initialized with lvalues.

</details>

### Item 25: Use _std::move_ on Rvalue References, _std::forward_ on Universal References

<details>
<summary>
Use the correct function depending on rvalue and universal reference objects.
</summary>

rvalue references bind only to objects that are candidates for moving, therefore, if we have a rvalue reference, we know it's a candidate for moving. so Once we have this kind of object, we want to permit other functions to use this object's rvalue-ness. this is what _std::move_ was created for.

```cpp
class Widget {
    public:
    Widget(Widget && rhs) //move constructor, definably an object eligible for moving
    : name(std::move(rhs.name)),p(std::move(rhs.p)) //use moves
    {

    }
    private:
    std::string name;
    std::shared_ptr<SomeData> p;
};
```

on the other hand, an _universal reference_ might be bound to an object that's eligible for moving, and might not be. we should cast it to an rvalue only if it was initialized as such, so we use _std::forward_.

```cpp
class Widget {
    public:
    template <typename T>
    void setName(T && newName) // type deduction + T&& form = universal reference
    {
        name =std::forward<T>(newName);
    }
};
```

rvalue references should be unconditionally cast rvalues with _std::move_, because they are always bound to rvalues. universal references should be conditionally cast into rvalues with _std::forward_, because they are only sometimes bound to rvalues.
It's a bad idea to try and mix those two. mixing _std::forward_ for rvalue requires extracting the types,which is possible, but the code is wordy and error prone, using _std::move_ when not bound the a rvalue is worse.

in this case, we pass a lvalue local object into the method, the object is then moved from, leaving us with a string object at an unknown state.

```cpp
class Widget {
    public:
    template <typename T>
    void setName(T&& newName)
    {
        name = std::move(newName); //compiles, but is a bad idea!
    }
    private:
    std::string name;
};

std::string getWidgetName(); //factory function
Widget w;
auto n = getWidgetName(); // n is a local lvalue object
w.setName(n); // we move n into w, what is tha value of n now?
```

we could have decided that there should be to functions for setName, one for lvalue which doesn't modify the parameter, and one for rvalue which does,

```cpp
class Widget {
    public:
    void setName(const std::string& newName)
    {
        name = newName; //lvalue, copy assignment
    }
    void setName(std::string&& newName)
    {
        name = std::move(newName); //rvalue, move assignment
    }
    private:
    std::string name;
};

std::string getWidgetName(); //factory function
Widget w;
auto n = getWidgetName(); // n is a local lvalue object
w.setName(n); // using the lvalue version
w.setName(getWidgetName()); // using the rvalue version
```

but this is a both more code to write, and can cause performance drawbacks:

```cpp
w.setName("Alpha Charlie");
```

in the universal reference version

1. string literal is passed into the setName function.
2. the data member is assigned from the string literal.

in the overloaded functions version

1. the string literal is used to create a temporary std::string object
2. this string is passed to the setName function
3. the assignment operator moves from the temporary string.
4. a destructor destroys the temporary string object.

we could create more functions for the \*const char\*\* case, but that leads us to the problem of scale. if we had two parameters, we would double the amount of functions needed, not two functions instead of one, but four, and then 8, and so on...

```cpp
class WidgetUniversal
{
    template <typename T>
    void twoParams(T&&a, T&&b)
    {
        f(std::forward<T,T>(a,b)); // do something with a,b
    }
};

class WidgetRvalue
{
    void twoParams(std::string && a, std::string &&b)
    {
        f(std::move(a),std::move(b)); // do something with a,b
    }

    void twoParams(const std::string & a, std::string &&b)
    {
        f(a,std::move(b)); // do something with a,b
    }

    void twoParams(std::string && a, const std::string &b)
    {
        f(std::move(a),b); // do something with a,b
    }
    void twoParams(const std::string & a,const std::string &b)
    {
        f(a,b); // do something with a,b
    }
};
```

this comes into play with the utility make functions, _std::make_shared_ (c++11) and _std::make_unique_ (c++14). they both take an unlimited amount of arguments, and they both use _std::forward_ internally.

```cpp
template <class T, class... Args>
std::shared_ptr<T> make_shared(Args&&... args);

template <class T, class... Args>
std::unique<T> make_unique(Args&&... args);
```

#### Using the Same Object Multiple Times

in some cases, we want to use our rvalue / universal reference objects more than once in the function, and we have to make sure that we don't move from it until the final operation. we therefore don't use _std::move_ or _std::forward_ until the final usage.
this is an example with _std::forward_, but _std::move_ is similar, although sometimes we will want to call _std::move_if_noexcept_ instead.

```cpp
template <typename T>
void setSignText(T&& text) //universal reference
{
    sign.setText(text); //use text without modifying it
    auto now = std::chrono::system_clock::now();
    signHistory.add(now, std::forward<T>(text)); //now we are done with text,
}
```

#### Returning from Functions

if we return an object from a function _by value_ and what we return is bound to an rvalue or universal reference, we would apply _std::forward_ or _std::move_ on the result.

```cpp
Matrix operator+(Matrix &&lhs, const matrix& rhs) //return by value
{
    lhs += rhs;
    return std::move(lhs); //move lhs into return value
}
auto Matrix m1;
//...
auto m2 = Matrix() + m2; //copy constructor
```

this ensures that the return value is treated as an rvalue when returning from the function.

```cpp
Matrix operator+(Matrix &&lhs, const matrix& rhs) //return by value
{
    lhs += rhs;
    return lhs; //copy lhs into return value;
}

auto Matrix m1;
//...
auto m2 = Matrix() + m2; //copy constructor
```

it the class supports move construction, we will get better performance, if not, nothing happens.
the same idea applies for universal references. if we call _std::forward_, then we could use the move constructor when appropriate, rather than always using the copy operations.

```cpp
template <typename T>
Fraction reduceAndCopy(T&& frac) //return by value
{
    frac.reduce();
    return std::forward<T>(frac);
}

Faction f1;
auto f2 = reduceAndCopy(f1); // f1 is lvalue, so copy constructor
auto f3 = reduceAndCopy(Fraction()); // rvalue, so move constructor
```

we shouldn't extend this logic to situations where it doesn't belong. like trying to optimize return values from copies into moves.

```cpp
Widget makeWidget1()
{
    Widget w;
    //...
    return w; //copy w into return value;
}

Widget makeWidget2()
{
    Widget w;
    //...
    return std::move(w); //bad! useless! don't!
}
```

the reason this is a bad idea is because the standards committee is way ahead of us. this behavior falls under the **return value optimizations (RVO)** idea, rather than copy the local object, the standards allows the object to be constructed on the caller site, rather than inside the local function scope. this is called **copy elision** - compilers may elide the copying or moving of objects from functions, and it can happen if

> 1. the type of the local object is the same as that returned by the function
> 2. the local object is what's being returned

in the normal version of the function, the return type is the same as the local object, and that's what's being returned. in the 2nd version, the conditions aren't being fulfilled. the return objects isn't the local object,it's an rvalue reference to it. so RVO won't happen here.

even in cases that we suspect RVO won't happen (because of multiple branches, different local variables, whatever), it's still a bad idea, as the the same rules for RVO require that even if there is no copy ellison, was long as the conditions are met, the returned object is treated as rvalue.

so in the original version of the function, the compiler can perform copy ellison, or return the local object as a rvalue (just like what we tried to optimize). a similar thing happens with returning pass by value parameters.

code written like this:

```cpp
Widget makeWidget(Widget w)
{
    return w;
}
```

is treated as if:

```cpp
Widget makeWidget(Widget w)
{
    return std::move(w);
}
```

trying to force the move to happen doesn't benefit us at any case. at most, we hinder the compiler by removing the RVO and forcing it to do what it was already perfectly willing to do.

#### Things To Remember

> - Apply std::move to rvalue references and std::forward to universal references the last time each is used.
> - Do the same thing for rvalue references and universal references being
>   returned from functions that return by value.
> - Never apply std::move or std::forward to local objects if they would other‐
>   wise be eligible for the return value optimization.

</details>

### Item 26: Avoid Overloading on Universal References

<details>
<summary>
A universal reference template function is greedy, and will be invoked in many cases, even at the expense of copy and move constructors.
</summary>

this is an possible function implementation. with some calls to it to show it's shortcomings.

```cpp
std::multiset<std::string> names;

void logAndAdd(const std::string& name)
{
    auto now =std::chrono::system_clock::now();
    log(now, "logAndAdd");
    names.emplace(name);
}
std::string petName("Darla");
logAndAdd(petName);
logAndAdd(std::string("Persephone"));
logAndAdd("Patty");
```

but we can make it better, after all, the copy into the names multiset is required only in the first case (where it's bound to a lvalue). in the other two cases, it seems like we're doing some unneeded work. in case 2, we have a rvalue string, but we still treat it inside the function as a lvalue so we must copy it, wouldn't it be better if it was a move? in the third case, we have a temporary string literal that we construct a string with, and then we copy from that string into the names multiset before destroying it. would have been nice to avoid this temporary object..

following what we know from before, let's re-write this function to avoid all that hassle by using universal references.

```cpp
template <typename T>
void logAndAdd(T&& name)
{
    auto now =std::chrono::system_clock::now();
    log(now, "logAndAdd");
    names.emplace(std::forward<T>(name));
}
std::string petName("Darla");
logAndAdd(petName); //same as before
logAndAdd(std::string("Persephone")); //move rvalue rather than copy
logAndAdd("Patty"); // create the std::string in the multiset rather the copy a temporary object
```

if that was the case, great, but now we say that users don't have the names directly, and the use an overload that takes an index and retrieves the name. but now the user tries calling it with some type other than int, like short or char. this is an error!

```cpp
std::string nameFromIndex(int idx);
void logAndAdd(int idx)
{
    auto now =std::chrono::system_clock::now();
    log(now, "logAndAdd");
    names.emplace(nameFromIndex(idx));
}
logAndAdd(22); //fine call function with int overload
short nameIndex;
//...

logAndAdd(nameIndex); //Error!
```

the reason is that of the two overloads, the highest priority is the int one, because non-templated have precedence over templated ones, however, the template overload can yield an exact match by creating a function that takes short and forwards it, which is considered better than promotion of the short into integer. therefore,the templated overload is invoked! the _std::forward_ chain begins, and now we try construct a std::string with a short inside the _.emplace()_ method, and there is no such overload.

functions with universal references are the **greediest** of them all, and they will instantiate exact matches for nearly all arguments type (exceptions are detailed in [Item 30]()). for this reason, combining universal references and overloads is a bad idea, the templated overload is applied to nearly every case.

#### The Constructor Problem

let's see this again by trying to write a perfect forwarding constructor for a class that behaves similar to above.

```cpp
class Person{
    public:
    template <typename T>
    explicit Person(T&& n):name(std::forward(n));
    {}
    explicit Person(int idx):name(nameFromIdx(idx))
    {}

    private:
    std::string name;
};
```

like before, attempting to call the person constructor with anything but int will call the universal reference overload. but because this is a class, there are generated copy and move operations. so our class actually looks like this, and here we get a new problem.

```cpp
class Person{
    public:
    template <typename T>
    explicit Person(T&& n):name(std::forward(n));
    {}
    explicit Person(int idx):name(nameFromIdx(idx))
    {}
    Person (const Person & rhs); //copy constructor
    Person (Person && rhs); // move operator

    private:
    std::string name;
};

Person p("Nancy");
auto cloneOfP(p); //error!
```

in this case, we aren't calling the copy constructor, we are calling the perfect forwarding constructor. this is another case of the overload resolution order.
the copy constructor is declared to accept a _const Person &_, but we pass it a non const person, which is a better fit for the templated overload. so in this case, our class actually looks like this:

```cpp
class Person{
    public:
    template <typename T>
    explicit Person(T&& n):name(std::forward(n));
    {}
    explicit Person(int idx):name(nameFromIdx(idx))
    {}
    explicit Person(Person& n): name(std::forward(n)) //instantiated method
    {}
    Person (const Person & rhs); //copy constructor
    Person (Person && rhs); // move operator

    private:
    std::string name;
};

Person p("Nancy");
auto cloneOfP(p); //error!
```

calling the copy constructor requires adding a const to match the signature,while the templated overload does not. had we declared our copied object const, things would work, because even if the templated version can be generated, the normal version is preferred.

```cpp
const Person p2("Drew");
auto cloneOfP2(p2); //fine
```

Things get worse with when inheritance comes into play. even if we managed to call the proper function on the SpecialPerson object, we still get the forwarding version of the base class constructor.

```cpp
class SpecialPerson: public Person
{
    public:
    SpecialPerson (const SpecialPerson & rhs):Person(rhs)
    {
    //copy constructor calls base class forwarding constructor
    }
    SpecialPerson (SpecialPerson && rhs): Person(std::move(rhs))
    {
    //move constructor calls base class forwarding constructor
    }
};
```

the reason for this weird behavior is that although copy and move constructors are supposed to take derived class and use polymorphism to treat them as if they were the base class, the templated version is still created and is a better match.

#### Things to Remember

> - Overloading on universal references almost always leads to the universal reference overload being called more frequently than expected.
> - Perfect-forwarding constructors are especially problematic, because they’re
>   typically better matches than copy constructors for non-const lvalues, and
>   they can hijack derived class calls to base class copy and move constructors.

</details>

### Item 27: Familiarize Yourself with Alternatives to Overloading on Universal References

<details>
<summary>
Trying to solve the issue of overloading with different tactics and the things to consider.
</summary>
Following on the problems form the item above, there are some suggested alternatives.

#### Abandon Overloading

for some simple cases, it might be enough to give up on the idea of overloading and having different names for the two functions. it would work for the 'logAndAdd' functions, but not for the constructors.

#### Pass by const T&

in this solution, we give up on using universal references and stick with passing by reference to const lvalue. it's less efficient, but it does avoid the problem.

#### Pass by Value

Counterintuitively, sometimes passing by value leads to better performance. this is further detailed in [Item 41](). but the general motive is to pass by value arguments that we know we wish to copy.

```cpp
class Person
{
    public:
    explicit Person(std::string n) : name(std::move(n)){}
    explicit Person(int idx) : name(nameFromIdx(idx)){}
    private:
    std::string name
};
```

in this case, all int and int-like objects will go to the int constructor, and all arguments of type std::string (and it's relatives) will use the string constructor.

#### Use Tag Dispatch

if we don't want to abandon the perfect forwarding and we don't want to abandon overloading, we can use tag dispatch.
recall that the universal reference overload is preferred because it can create perfect matches based on the arguments and the call site. but if we have a parameter that isn't part of the universal reference list, there is no way for the perfect overload to be created, and the regular overload are now used.

we do this by using calling internal implementation functions who perform the overloading and have an extra parameter which isn't part of the universal reference.

we start with the old code:

```cpp
std::multiset<std::string> names;
template <typename T>
void logAndAdd(T&& name)
{
    auto now =std::chrono::system_clock::now();
    log(now, "logAndAdd");
    names.emplace(std::forward<T>(name));
}
```

and for now, we start with this incorrect version, which fails for lvalue reference of integrals values, in which T is int&, which is not an integral type

```cpp
template <typename T>
void logAndAdd(T&& name)
{
    logAndAddImpl(std::forward<T>(name),std::is_integral<T>());
}
logAndAdd(5L); //works
int x =5;
logAnddAdd(x); //doesn't work
```

to fix this case, we use the standard library to remove the reference qualifiers from T, so now reference types are also considered integrals

```cpp
template <typename T>
void logAndAdd(T&& name)
{
    logAndAddImpl(std::forward<T>(name),std::is_integral<typename std::remove_reference<T>::type>)();
}
logAndAdd(5L); //works
int x =5;
logAnddAdd(x); //also
```

a c++14 version would look like to remove some keystrokes

```cpp
template <typename T>
void logAndAdd(T&& name)
{
    logAndAddImpl(std::forward<T>(name),std::is_integral<typename std::remove_reference_t<T>>());
}
```

the implementations functions would have the second argument outside of the universal reference, the argument will be differentiated by type, not by value, to get the correct function resolved in compile time.

```cpp
template <typename T>
void logAndAddImpl(T&& name,std::false_type)
{
    auto now =std::chrono::system_clock::now();
    log(now, "logAndAdd");
    names.emplace(std::forward<T>(name));
}

void logAndAddImpl(int idx,std::true_type)
{
    logAndAdd(nameFromIdx(idx)); //call first function again.
}
```

both the _std::false_type_ and _std::true_type_ aren't used a parameters, they don't have names, and the compiler should ignore then after resolving the function calls. they act as tags. this is the reason behind the name of **tag dispatch**. we don't overload the function that has only universal references, we overload the implementation functions, which we can better control.

#### Constraining Templates that Take Universal References

in tag dispatch, a key point is the existence of one front facing function the client calls, and this approach solves the case for regular functions, but we still have the issue with perfect forwarding constructors and generated copy and move operations. the problem is we don't always bypass the tag dispatch version, which is what we want to happen (always use the copy constructor).

we want our copy constructor to be called even for non-const values, and for the correct base constructor to be called from inheritance class. for these situations we do away with tag-dispatch and welcome _std::enable_if_ and the concept of **SFINAE: Substitution Failure is not an Error**.

_std::enable_if_ forces the compiler to consider or ignore a template based on a condition, if the condition isn't met, the template won't be part of the overload resolution. in our case, we want to exclude the forwarding constructor from taking effect if the type is Person, or alternately, to allow if only if it is not.

the abstract versions looks like this

```cpp
class Person{
    public:
    template <typename T,typename std::enable_if<condition>::type>
    explicit Person(T &&);
};
```

the most simple versions looks like this

```cpp
class Person{
    public:
    template <typename T,typename std::enable_if<!std::is_same<Person,T>>::value>::type>
    explicit Person(T &&);
};

Person p("Nancy");
auto copyOfP(p); //still an error! Person and Person&
```

which will fail us for lvalue references of Person. as we saw above, we need to remove the reference qualifiers so that _Person_, _Person &_ and _Person&&_ are all matched to
_Person_, additionally, we would want to ignore any _const_ or _volatile_ qualifiers. for this we use _std::decay<>_, which performs both. _std::decay\<T>::type_ strips away all qualifiers from T .

```cpp
class Person{
    public:
    template <typename T,typename std::enable_if<!std::is_same<Person,typename std::decay<T>::type>>::value>::type>
    explicit Person(T &&);
};

Person p("Nancy");
auto copyOfP(p); //now this works
```

and even tough we are so close, we still have the issue of inheritance, no matter how much qualifiers we remove from a SpecialPerson class, it won't be a Person, so we need to handle this case and stop derived copy and move operations. so we turn our attention to _std::is_base_of<T1,T2>_, and just to be calm, a type is considered to be the base_of itself.

```cpp
class Person{
    public:
    template <typename T,typename std::enable_if<!std::is_base_of<Person,typename std::decay<T>::type>>::value>::type>
    explicit Person(T &&);
};

Person p("Nancy");
auto copyOfP(p); //now this works
SpecialPerson sp("aa");
SpecialPerson sp2(sp); //also works
```

if we happen to use c++14, we get some utility functions to reduce the weird parts of the code and maintain the functionality, the _\_t_ suffix implies the _::type_ part in the end.

```cpp
class Person{
    public:
    template <typename T,typename std::enable_if_t<!std::is_base_of<Person,typename std::decay_t<T>>::value>>
    explicit Person(T &&);
};
```

and now we need to combine it together with the two constructors forms (string, index) to get the benefits we had with tag dispatch.

```cpp
class Person{
public:
template <typename T,std::enable_if_t<
    !std::is_base_of<Person, typename std::decay_t<T>::value>>
    &&
    !std::is_integral<std::remove_reference_t<T>>::value
    >
>
explicit Person(T&& n): name(std::forward(n))
{

}
explicit Person(int idx): name(nameFromIdx(idx))
{

}
private:
std::string name;
};
```

this form is suitable for constructor, where we cannot create overloads or use different function names.

#### Trade Offs

we had five options in this chapter, the first three (_abandoning overloading_, _passing by const &T_ and _passing by value_) specify a type for each parameter to be called, the other options (_tag dispatch_,_constraining template eligibility_) use perfect forwarding, and don't specify types for parameters. this decision of whether or not to specify the type has consequences.

in general perfect forwarding is more efficient, we avoid creating temporary objects just to conform to a the function signature, like forwarding a string literal rather than construct and destruct a string. but one problem is that there are some **kinds of arguments which cannot be perfect-forwarded**. another problem is the comprehensibility of error messages for invalid arguments.

```cpp
Person p(u"Konard Zuse"); //char16_t
```

if we use a _char16_t_ string literal (an type of characters that represent 16 bit characters) and we pass it into one of the functions using the first set of functions, the error message would say that there is no available constructor for this class, and that we cannot create an std::string or an int from _char16_t_ objects, readable, more or less.
with the perfect forwarding set, the argument will be passed into the std::string constructor, and we will see a huge error message detailing how _char16_t_ cannot be converted into any of the type that can construct a string. the more levels of indirection and abstraction, and the more parameters are forwarded, the longer the message will be and it will make less sense.

in some simple cases, we can attempt to catch this case before hand with a _static_assert_ statement and our custom message.

```cpp
static_assert(std::is_constructible<std::string,T>::value,"Parameter n can't be used to construct a std::string");
```

but this won't work for us, because the _std::forward_ part happens in the members initiation list, not the the constructor body.in some compilers we might see both the long message and our custom one.

#### Things To Remember

> - Alternatives to the combination of universal references and overloading
>   include the use of distinct function names, passing parameters by lvalue reference-to-const, passing parameters by value, and using tag dispatch.
> - Constraining templates via std::enable_if permits the use of universal references and overloading together, but it controls the conditions under which compilers may use the universal reference overloads.
> - Universal reference parameters often have efficiency advantages, but they typically have usability disadvantages.

</details>

### Item 28: Understand Reference Collapsing

<details>
<summary>
We can not have references-to-references. the compiler can have them, and resolves them down to either an lvalue or an rvalue reference. this is what drives std::future.
</summary>

There are four contexts in which reference collapsing occurs

1. template instantiation.
2. auto variables.
3. typedefs and alias decelerations.
4. decltype.

#### Template Instantiation

when we pass an argument into a template function, the type deduction depends on whether the argument is lvalue or rvalue (which controls the behavior of _std::forward_). if we pass a lvalue, then T is an lvalue reference. if we pass a rvalue, then T is a non-reference.

```cpp
template <typename T>
void func(T&& param);
Widget widgetFactory(); //return rvalue
Widget w;
func(w); //T is Widget&
func(widgetFactory())// T is Widget
```

even though we pass a Widget object in both calls, the type of T is different.
and just to note, we can't have a reference to a reference (unlike pointer to pointer).

```cpp
int x;
auto & rx = x; //rx in reference to x;
auto & & rrx = x; // error! can't declare a reference to a reference!
```

but what about cases when an lvalue is passed to a function template taking a universal reference? if T is an lvalue reference, then the resulting function should be a reference to a reference, which is not allowed!

```cpp
void func(Widget& && param); // this shouldn't be allowed! so why is ok?
```

the answer is **reference collapsing**, the user may not declare references to references, but compilers can create them in some cases, and when they do, the process of references collapsing determines the outcome.

there are two kinds of references (lvalue, rvalue), so in total there are 4 cases.

> If either reference is an lvalue reference, the result is an lvalue reference.
> Otherwise (i.e., if both are rvalue references) the result is an rvalue reference.

| reference |      lvalue      |        rvalue        |
| --------- | :--------------: | :------------------: |
| lvalue    | lvalue reference |   lvalue reference   |
| rvalue    | lvalue reference | **rvalue reference** |

with _std::forward_, the type parameter T will determine if a cast to rvalue is needed or not. here is a basic implementation of _std::forward_.

```cpp
template <typename T>
T&& forward(typename remove_reference<T>::type & param)
{
    return static_cast<T&&>(param);
}
Widget w;
forward(w);
```

which will be

```cpp
Widget& && forward(typename remove_reference<Widget&>::type& param)
{
    return static_cast<Widget& &&>(param);
}
```

this will cause reference collapsing on the return type and statement

```cpp
Widget& forward(Widget& param)
{
    return static_cast<Widget&>(param);
}
```

so in the end, this does nothing for lvalue references, as we already know for _std::future_. now lets do the same with an rvalue reference.this time, T isn't a reference!

```cpp
Widget&& forward(typename remove_reference<Widget>::type& param)
{
    return static_cast<Widget&&>(param);
}
//which is
Widget&& forward(Widget & param)
{
    return static_cast<Widget&&>(param);
}
```

inside the function, the parameters are always lvalue, but we know they can be safely moved from, so wer return them as rvalue references.

in c++14 we can use _std::remove_reference_t_ to clear to code a bit.

```cpp
template <typename T>
T&& forward(remove_reference_t<T>&param)
{
    return static_cast<T&&>(param);
}
```

#### Auto Variables

the context of auto is similar to type deduction for templates.

```cpp
Widget w;
auto && w1 = w;
auto && w2 = widgetFactory();
```

w1 is initialized with an lvalue(w), so the type of is _Widget&_. w2 is initialized with an rvalue, so the type is _Widget&&_.

```cpp
Widget& && w1 =w;
Widget&& && w2 = widgetFactory();
```

and reference-to-reference implies reference collapsing.

```cpp
Widget& w1 = w;
Widget&& w2 = widgetFactory();
```

w1 is an lvalue reference. but w2 is an rvalue reference.

so back to universal references. it isn't a new kind of reference, it's an rvalue reference in the context of two conditions:

> - Type deduction distinguishes lvalues from rvalues. Lvalues of type T are
>   deduced to have type T&, while rvalues of type T yield T as their deduced type.
> - Reference collapsing occurs

#### Typedefs and Alias Decelerations, Decltype

if, in some point during evaluation of a typedef, there happens to be reference-to-reference situation, reference collapsing will resolve the confusion.

```cpp
template <typename T>
class Widget{
    public:
    typedef T&& RvalueRefToT;
};
Widget<int &> w;
```

so now we go through the motions of reference expanding and collapsing

```cpp
typedef int & && RvalueRefToT;
typedef int & RvalueRefToT;
```

which makes it clear that the name is not correct, we don't have an rvalue reference, it's a plain old lvalue reference.

in cases of decltype, the reference collapsing also happens when needed.

#### Things to Remember

> - Reference collapsing occurs in four contexts: template instantiation, auto type
>   generation, creation and use of typedefs and alias declarations, and
>   decltype.
> - When compilers generate a reference to a reference in a reference collapsing
>   context, the result becomes a single reference. If either of the original references is an lvalue reference, the result is an lvalue reference. Otherwise it’s an rvalue reference.
> - Universal references are rvalue references in contexts where type deduction
>   distinguishes lvalues from rvalues and where reference collapsing occurs.

</details>

### Item 29: Assume That Move Operations are Not Present, Not Cheap, and Not Used

<details>
<summary>
When we write code, we should not assume move operations will always be invoked, and if we don't know they will, we should write code that isn't reliant on them
</summary>
move semantics are perhaps the most noticeable part of c++11 and modern c++. because of that, some statements arise, ranging from untrue, oversimplified, and inaccurate, such as:

> "Moving containers is now as cheap as copying pointers"
>
> "Copying temporary objects is now so efficient, coding to avoid it as tantamount to premature optimizations"

and while move semantics are great, and they actually require compilers to replace some copy operations with move operations, those sentiments are exaggerated.

> - No move operations: The object to be moved from fails to offer move operations. The move request therefore becomes a copy request.
> - Move not faster: The object to be moved from has move operations that are no
>   faster than its copy operations.
> - Move not usable: The context in which the moving would take place requires a
>   move operation that emits no exceptions, but that operation isn’t declared noexcept.
> - Source object is lvalue: With very few exceptions (see e.g., [Item 25]()) only rvalues may be used as the source of a move operation.

#### Many Types Still Don't Support Move Semantics

the standard library has been updated to support move semantics, but chances are that most of the user code classes were not. Maybe even some of the external libraries are still lagging behind. the compiler will generate move operations on it's own, but only for classes that fit the criteria in [Item 11](), those which didn't declare copy operations,move operation or a destructor.

#### Moving isn't Always Cheap

even classes that have those operations might not benefit from them, in some cases, there might not be a cheap way to move, like _std::array_. unlike other data containers, which store the data on the heap, the _std::array_ is stored on the stack, so it's not a simple matter of moving pointers around. moving from a _std::vector_ is really mostly moving a single pointer (approximately) between on container and the other, which should happen in constant time. But _std::array_ will need to move all elements, one by one, assuming that the type supports move operations. so the amount of work to move a _std::array_ is proportional to it's size.

for the _std::string_ object, even though most string objects are storing data on the heap, they can also employ _small string optimizations (SSO)_, in which they store all the data on the stack inside the _std::string_ object. in theses cases, moving is not better than copying (for larger strings, it probably is).

#### Moving isn't Always Used

even if the move operation is implemented and is cheaper than copying, it's not always what the compiler chooses to use. as descried in [Item 14](), even if copying can sometimes be replaced by moving, this might require the strong exception safety guarantee, so the compiler might still resort to copying over moving.

#### Things to Remember

> - Assume that move operations are not present, not cheap, and not used.
> - In code with known types or support for move semantics, there is no need for
>   assumptions.

</details>

### Item 30: Familiarize Yourself With Perfect Forwarding Failure Cases

<details>
<summary>
"Perfect forwarding" isn't perfect. despite the name, prefect forwarding has it's caveats and pitfalls. they can be overcome if we know what they are.
</summary>

Perfect forwarding is meant to ensure that forwarding a parameter from a function to a different function maintains the exact form o the original. this means by-value parameters are ignored, because copy by value means the original isn't part of the game any more, Pointer parameters are also out, we don't use them any more. so we are left with reference parameters.

when we forward something, we don't only forward the object, we also push the salient characteristics: the type, whether they are lvalue or rvalue references, their const or volatile qualifiers. this means we are dealing with _std::forward_ and universal references.

```cpp
template <typename T>
void fwd(T&& param)
{
    f(std::forward<T>(param));
}
```

forwarding function are generic by nature, and rather then just being function templates, they should be variadic function templates, able to take any number and type of arguments. this is the form for the _std::make_shared_, _std::make_unique_ functions, as well as the _emplacement_ functions for containers, therefore:

```cpp
template <typename... Ts>
vod fws(Ts&&... params)
{
    f(std::forward<Ts>(param)...);
}
```

our criteria is that perfect forwarding should do exactly the same as directly calling the function.

```cpp
f(expression);
fwd(expression);
```

there are some cases in which this isn't the case.

#### Braced Initializer

suppose f is declared to take a const lvalue reference to _std::vector_, we can use braced initialization in the direct call, but not in the forwarded call!

```cpp
void f(const std::vector<int> & v );
f({1,2,3,}); // "{1,2,3}" is implicitly used to construct a vector
fwd({1,2,3}); // this doesn't compile!
```

in the direct call, the compiler can compare the argument to the required type, and perform the implicit conversion if needed. from `{1,2,3}` to _std::vector_, no problem. but the forwarded call can't do this. they use type deductions:

perfect forwarding fails:

> - Compilers are unable to deduce a type for one or more of fwd’s parameters. In
>   this case, the code fails to compile.
> - Compilers deduce the “wrong” type for one or more of fwd’s parameters. Here,
>   “wrong” could mean that fwd’s instantiation won’t compile with the types that
>   were deduced, but it could also mean that the call to f using fwd’s deduced types
>   behaves differently from a direct call to f with the arguments that were passed to
>   fwd. One source of such divergent behavior would be if f were an overloaded
>   function name, and, due to “incorrect” type deduction, the overload of f called
>   inside fwd were different from the overload that would be invoked if f were
>   called directly.

in our case, the type of braced initializer is, by the language standards, uneducable in this context. the compiler is prevented from deducing the type, so it must reject that call.
there is a simple work around because _auto_ can deduce the type of _std::initializer_list_, so we could do this

```cpp
auto il = {1,2,3};
fwd(il);
```

#### 0 or Null as null Pointers

As we saw earlier in [Item 8](), passing a NULL or zero to a template function treats the value as an integral type. the solution is to use **_null_ptr_**.

#### Declaration-Only Integral _static const_ Data Members

we don't need to define integral _static const_ data members, the declaration is enough. there is also no need to set aside memory for them, the compiler will replace any use of the member data with the value. but perfect forwarding is all about references, so what now?

this is another case of perfect forwarding failure.

```cpp
class Widget{
public:
static const std::size_t MinVals = 28; // integral static const
};

std::vector<int> widgetData;
widgetData.reserve(Widget::MinVals); //no problem
void f(std::size_t val);
f(Widget::MinVals); // works
fwd(Widget::MinVals); // error! shouldn't link!
```

under the hood, references are pointers, so we are passing the address of something without an address (rvalue variables still have addresses, even if they are on the stack for a temporary time).

This is how it should be. the standard dictates that passing something requires it to be defined. but not all implementations enforce this.
if we want to make our code compliant with the rules, we need to separate the deceleration and the definition

```cpp
//header
class Widget{
public:
static const std::size_t MinVals=28; // integral static const
};
//cpp file

const std::size_t Widget::MinVals; // no need to add anything
```

#### Overloaded Function Names and Template Names

passing a function name that has been overloaded confuses the perfect forwarding.

```cpp
void f(int (*pf)(int)); //f takes a function that takes int and returns int
//void f(int pf(int)); //simpler declaration
int processVal(int value);
int processVal(int value, int priority);
f(processVal); //works
fwd(processVal); // doesn't work
```

the name `processVal` doesn't carry by itself any information about it's type, it's simply a pointer, there is no type deduction and the compiler can't choose the correct version. same goes with passing a _function template_. which isn't a single value, but can represent many functions.

```cpp
template <typename T>
T workOnVal (T param)
{
    ///
}
fwd(workOnVal); //error! which template instantiation?
```

we can overcome this issue by being specific about what we are passing.

```cpp
using ProcessFuncType = int (*)(int); //alias declration;
ProcessFuncType processValPtr = processVal;
fwd(processValPtr); //will work
fws(static_cast<ProcessFuncType>(workOnVal)); //also works.
```

this will only work if we know what type of function is really expected.

#### Bitfields

bitfields can also cause problems when used in perfect forwarding. this is because:

> "A non-const reference shall not be bound to a bit-field"

clear as day,can't take a non const reference of a bit field. and just as you can't take the address of a bitfield (it might be mid word!), we cant take a non-const reference to it.

the IPv4Header uses a number to as it's data member, with different bits to denote different properties (4+4+6+2+16 =32).

```cpp
struct IPv4Header{
    std::uint32_t version:4,
                IHL:4,
                DSCP:6,
                ECN:2,
                totalLength:16;
    //..
};
IPv4Header h;
f(h.totalLength); //fine
fwd(h.totalLength); //error!
```

we simply aren't allowed to pass a non-const reference to a bit field. we can pass them by value (make a copy of) or pass by a const reference, which makes the copy under the hood.

we overcome this issue by copying it.

```cpp
auto length = static_cast<std::uint16_t>(h.totalLength);
fwd(length);
```

#### Things to Remember

> - Perfect forwarding fails when template type deduction fails or when it deduces
>   the wrong type.
> - The kinds of arguments that lead to perfect forwarding failure are braced initializers, null pointers expressed as 0 or NULL, declaration-only integral const
>   static data members, template and overloaded function names, and bitfields.

</details>
