## Chapter 3 - Moving to Modern C++

<summary>
The small parts to remember when using modern C++.
</summary>

the modern versions of c++ define all sorts of high level ideas and features like move semantics, lambdas, concurrency and smart pointers. but also many smaller features, that don't necessary mean big swiping changes, but are important none the less. this chapter deals with them, what they are, and how to build a modern code style.

### Item 7: Distinguish Between _()_ And _{}_ When Creating Objects

<details>
<summary>
Braced Initialization might be uniform, but it's not without issues.
</summary>
there are many ways to create objects in cpp, and there are subtle differences between using the normal parentheses, the equal sign and the curly braces.

```cpp
int x(0); // initializer in parentheses
int y = 0; // initializer follows "="
int z{0}; // initializer is in braces
int z2 = {0}; // initializer uses both "=" and curly braces.
```

we will ignore that last option, because we already saw that it can lead to trouble if used together with the auto keyword. the equal sign in a problem, it is used both as an assignment operator and as a constructor call.

```cpp
Widget w1; //default ctor
Widget w2 = w2; //copy constructor.
w1 = w2; //assignment, copy operator=
```

#### The Advantages

<details>
<summary>
Curly braces initialization can work when other cases fail
</summary>
cpp11 introduces the concept of **uniform initialization**, or **braced initialization**. the syntax is the curly braces. we should be able to use it for everything that we could use equal sign or parentheses: from basic variables, initial contents of containers,inside classes and to initialize un-copyable objects.  
in some of those situations, only one of the two classic methods work to initialize a value, but in all cases, the braced initialization does the job.

```cpp
std::vector<int> v1 = {1,2,3}; //vector with initial contents. works
std::vector<int> v2 {1,2,3}; //vector with initial contents. fine.
//std::vector<int> v3 (1,2,3); //doesn't work!

class Widget
{
    private:
    int x{0}; // works
    int y =0; // no problem
    //int z(0); //error! can't be done
};

std::atomic<int> ai1(0); //works
std::atomic<int> ai2{0}; //also works!
//std::atomic<int> ai3 = 0; // doesn't work
```

another feature is that it doesn't allow narrowing conversions between built in types. if it's not 100% assured to work, it won't compile! the standard had to allow this for the older methods to avoid breaking legacy code, but with the new braced initialization, we can get better type safety!

```cpp
double x,y,z;
int sum1(x+y+z); //value truncated to an int.
int sum2 =x+y+z; //value truncated to an int.
//int sum3{x+y+z}; //rejected!
```

another upside of braced initialization is that it avoids another one of those parsing pitfalls. assume Widget has a default constructor and a single argument constructor. if we forget to pass the argument but still have the parentheses, we aren't calling the default constructor, we are actually declaring a function. good luck seeing that in a glance! braced initialization can't be in a function parameter lists declaration, so this is not an issue.

```cpp
Widget w0; // default constructor
Widget w1(0); //great, constructor call;
Widget w2(); // this isn't a compiler error! this is declaration of a function called w2 which takes no value and returns a Widget.
Widget w3{}; // this wil also work and call the default constructor.
```

</details>

#### The Disadvantages

<details>
<summary>
Compilers really, and a I mean really, prefer constructors with std::initializer_lists over other constructors, and the curly braces really loves to turn into std::initializer_lists.
</summary>

Things aren't entirely perfect in the realm of braced initialization. this syntax sometimes carries it's own bag of surprising behavior. we see this a lot with the relationship between the curly braces and the **std::initializer_list class and constructor overload resolution.** we think the code does one thing, but it actually does something else. add the auto keyword to the mix and we have a party.

```cpp
int a = {1}; //weird, but works
auto b = {1}; // surprise! an initializer_list<int> remember item 2?
```

and now for the fun bits, here is a normal, class two two constructors, none of them use a std::initializer_list as arguments.

```cpp
class Widget
{
    Widget(int i, bool b); // constructor 1 - no std::initializer_list involved
    Widget(int i, double d); // constructor 2 - no std::initializer_list involved
};

Widget w1(10,true); //constructor 1 is called
Widget w2{10,true}; //constructor 1 is called
Widget w3(10,5.0); //constructor 2 is called
Widget w4{10,5.0}; //constructor 2 is called
```

but here is a version that does use std::initializer_list. enjoy the mess of implicit conversations.

```cpp
class Widget
{
    Widget(int i, bool b); // constructor 1
    Widget(int i, double d); // constructor 2
    Widget(std::initializer_list<long double> li);  // constructor 3 - std::initializer_list involved
};

Widget w1(10,true); //constructor 1 is called
Widget w2{10,true}; //bad! constructor 3 is called! 10 and true are converted to long double!
Widget w3(10,5.0); //constructor 2 is called
Widget w4{10,5.0}; //bad! constructor 3 is called! 5.0 and true are converted to long double!
```

and here is another issue, which messes us copy and move constructors.

```cpp
class Widget
{
    Widget(int i, bool b); // constructor 1
    Widget(int i, double d); // constructor 2
    Widget(std::initializer_list<long double> li);  // constructor 3 - std::initializer_list involved
    Widget(const Widget & w); // copy constructor;
    Widget(Widget && w); // copy move constructor;
    operator float() const; //conversion to float;
};

Widget w5(w4); // parentheses, copy constructor
Widget w6{w4}; // nope. w4 is converted into float, float into long double and then initializer list, so constructor 3 is called.
Widget w7(std::move(w4)); // parentheses, move constructor
Widget w8{std::move(w4)}; // curly braces, same as above.
```

compilers really love to match constructors with std::initializer_lists, even if it prevents matches that should be better.

```cpp
class Widget
{
    Widget(int i, bool b); // constructor 1
    Widget(int i, double d); // constructor2
    Widget(std::initializer_list<bool> li);  // constructor 3 - std::initializer_list involved
}
Widget w9(10,5.0);// parentheses, work fine
Widget w10{10,5.0}// braces, error! tries to convert 10,5 into std::initializer_list<bool>, which is narrowing and isn't allowed! ignoring the better option!
```

the std::initializer_list constructors are preferred in nearly all cases, only if there is no possible conversion the compiler falls back to other constructors.

```cpp
class Widget
{
    Widget(int i, bool b); // constructor 1
    Widget(int i, double d); // constructor2
    Widget(std::initializer_list<std::string> li);  // constructor 3 - std::initializer_list involved. but this time a none-numeric type so there is not implicit conversion.
};
Widget w11(10,true); //parentheses, works
Widget w12{10,true}; //braces, works because the compiler doesn't try to force values into strings
Widget w13(10,5.0); //parentheses, works
Widget w14{10,5.0}; ///braces, works because the compiler doesn't try to force values into strings.
```

a final edge case is the case of empty braces for a class that supports default constructor and initialization with std::initializer_list. what do they empty braces mean in this situation? is it an empty list or no arguments? **the rule is that you should get default constructor**. if we want the empty list, we should create it.

```cpp
class Widget
{
    Widget();//default constructor
    Widget(std::initializer_list<std::string> li); //initialization with std::initializer_list.
};

Widget w15; //default;
Widget w16{}; //default;
Widget w17(); //actually a function declaration, sorry.
Widget w18({});// empty initializer_list
Widget w19{{}};// empty initializer_list
```

this weird behavior hits the c++ community in the place that hurts us the most, std::vector. the vector has a constructor taking an std::initializer_list for initial values and a constructor for the number of elements and the initial value of each element.

```cpp
std::vector<int> v1(10,20); //vector with 10 elements of value 20;
std::vector<int> v2{10,20}; //vector with two elements, 10 and 20.
```

as class creators, we should be careful when adding std::initializer_list to our constructors, and we should design our objects so that they behave the same for regular parentheses and curly braces.  
as class consumers, we need to choose deliberately what we call, and what is our default style of constructor call.

if we have a template class, we get another layer of fun, because now we also must decide how to forward calls. in this example, a variadic template argument needs to be constructed into T, and we have to choose how. and there is no correct choice.

```cpp
template<typename T, typename... Ts>
void doSomeWork(Ts &&... params)
{
    T localObject1(std::forward<Ts>(params)...); //using parentheses
    T localObject2{std::forward<Ts>(params)...}; //using curly braces
}
```

in the standard, std::make_unique and std::make_shared decided on the parentheses call, which is included in the documentation.

</details>

#### Things to Remember

> - Braced initialization is the most widely usable initialization syntax, it prevents narrowing conversions, and it’s immune to C++’s most vexing parse.
> - During constructor overload resolution, braced initializers are matched to
>   std::initializer_list parameters if at all possible, even if other constructors offer seemingly better matches.
> - An example of where the choice between parentheses and braces can make a
>   significant difference is creating a std::vector<numeric type> with two
>   arguments.
> - Choosing between parentheses and braces for object creation inside templates
>   can be challenging

</details>

### Item 8: Prefer _nullptr_ To 0 And NULL

<details>
<summary>
NULL and 0 are numbers, not pointers. this can cause surprises, which we prefer not to see.
</summary>
the value 0 is an int, we can use it as a pointer address, but that's a fallback. we can define NULL to be 0 or as long type with value zero, the problem is that neither zero or NULL are pointers types.
in classic C++, this meant that we could get some surprises if we had overloads for int and pointer types. and type conversions might results in weird behavior again.

```cpp
#define NULL 0;
// regular null, int
#define LongNULL 0L;
//long null, long;
void f(int);
void f(bool);
void f(void*);
f(0); // calls f(int)
f(NULL); // might not compile, or worse, will call f(int), but never f(void*)
f(LongNULL); //ambiguous call, long can be converted to either int, bool or void*, the same level of priority
```

there is a guideline to avoid overloading on pointer and integral types, exactly because of this reason. in contrast, nullptr isn't an integral type, its' actual type is std::null_pointer_t. a type that can acts as a pointer to any other type, but can't be converted into other types. it also allows us more expressiveness, if something is compared to nullptr, it must be a pointer;

```cpp
f(nullptr); //calls f(void*)
auto res = foo(); //some function with some return type;
if (res ==0) //this is what res ==NULL means
{
    //does that mean res was a number?
}
if (res == nullptr)
{
    //now it's clear res is a pointer.
}
```

this also has advantages with templates. in this example we have functions and a mutex for each function. we can somehow convert zero and NULL to smart pointer types, but nullptr won't budge into something that it's not.

```cpp
int f1(std::shared_ptr<Widget> spw);
double f2(std::unique_ptr<Widget> upw);
bool f3(Widget* pw);
std::mutex f1m,f2m,f3m;
using MuxGuard = std::lock_guard<std::mutex>; // using statement c++
{
    MuxGuard(f1m);
    auto result = f1(0); // no problem int is int; nullptr won't work here
}
{
    MuxGuard(f2m);
    auto result = f2(NULL); //NULL is int, and int can be boolean. nullptr won't work here
}
{
    MuxGuard(f3m);
    auto result = f3(nullptr); //pointer is pointer,
}
```

and here is a templated version

```cpp
template<typename FuncType,typename MuxType, typename PtrType>
auto lockAndCall(FuncType func, MuxType & mutex,PtrType ptr)-> decltype(func(ptr))
{
MuxGuard g(mutex);
return func(ptr);
}

auto r1 =lockAndCall(f1,f1m,0); //error! f1 isn't expecting int, it wants a unique_ptr;
auto r1 =lockAndCall(f2,f2m,NULL); //error! f2 isn't expecting int, it wants a shared_ptr;
auto r1 =lockAndCall(f3,f3m,nullptr); // this is fine, f3 wants a pointer type, which is what it gets.
```

the function's return type is whatever the return type of calling func on ptr is. that's what decltype does. in c++14 we won't even need that decltype.
when we use nullptr. we get type safety and less surprises.

#### Things to Remember

> - Prefer nullptr to 0 and NULL.
> - Avoid overloading on integral and pointer types.

</details>

### Item 9: Prefer Alias Decelerations to _typedef_

<details>
<summary>
In a general case, alias decelerations are easier to read, write, and use.
</summary>
in classic c/c++98 fashion, if we wanted to shorten the name of a type, we could use typedef. in modern c++,we have alias decelerations. at first glance, they seem to be the same thing, just with the positions switched.

```cpp
typedef std::unique_ptr<std::unordered_map<std::string, std::string>> UPtrMapSS; //typedef, a unique_pointer to unordered map of key string and value string.
using APtrMapSS = std::unique_ptr<std::unordered_map<std::string, std::string>> ;// alias deceleration;
```

but when we get to function types, the difference is clearer. the alias deceleration is consistent.

```cpp
typedef void (*FP_TD)(int, const std::string &); //typedef, the type name is FP_DD. better not forget the *.
using FP_AD = void(*)(int, const std::string &); //alias declaration, they name is always in the left side,
```

but another case for them to shine is templates: alias decelerations can be templated and nested, typedef can at best be hacked into submission using a templated struct.
in this case, we try to define a list for any type using a custom allocator.

```cpp
template <typename T>
using MyAllocatedAliasList = std::list<T,MyAllocator<T>>;
MyAllocatedAliasList<Widget> aliasLw;

template<typename T>
struct MyAllocatedList{
    typedef std::list<T, MyAllocator<T>> type;
};
MyAllocatedList<Widget>::type lw;
```

and now typedef inside a template to create a linked list, we must precede the typedef name with _typename_

```cpp
//alias declaration.
template <typename T>
class AliasWidget
{
private:
    MyAllocatedAliasList<T> list;// this is all;
};
//typedef
template<typename T>
class Widget{
    private:
    typename MyAllocatedList<T>::type list; // typename preceding and ::type
}
```

here is an example of bad code that will cause issus with typedef, we are dependent on T and MyAllocatedList, so if someone changed the specialization and defined type differently, things will be bad.

```cpp
class Wine{
    //..
};
template<>
class MyAllocatedList<Wine> //temperate specialization on MyAllocatedList<Wine>
{
    private:
    enum class WineType
    {White,Red,Rose};
    WineType type;//oops, now type is a data member

};
```

in the case of templates, c++11 gives us _type traits_ inside the header<type_traits>. notice how they all use the ::type to match the type, this is actually a usage of the typedef version for historical reasons.

- std::remove_const\<T>::type - return T from const T;
- std::remove_reference\<T>::type - return T from & and T&&;
- std::add_lvalue_reference\<T>::type - return T& from T;

c++14 decided that alias decelerations are better, so they introduced a shorthand for this.

- std::remove_const_t\<T>- return T from const T;
- std::remove_reference_t\<T> - return T from & and T&&;
- std::add_lvalue_reference_t\<T> - return T& from T;

#### Things to Remember

> - typedefs don’t support templatization, but alias declarations do.
> - Alias templates avoid the “::type” suffix and, in templates, the “typename”
>   prefix often required to refer to typedefs.
> - C++14 offers alias templates for all the C++11 type traits transformations

</details>

### Item 10: Prefer Scoped _enums_ to unscoped _enums_

<details>
<summary>
Scoped enums prevent namespace pollution, unintended conversions and promote type safety.
</summary>
C style enum are unscoped,they belong in the same scope as other variables. scoped enums are limited inside their own scope. this scope behaves like a mini-namespace. this means we can reduce namespace pollution.

```cpp
enum Color{black,white,red}; //unscoped enum, black, white,red are same scope as Color;
//auto white = false; //error! white is already declared!

enum class ScopedColor{green, blue, orange}; //scoped enum, the values exist only int the ScopedColor scope.
auto green = 9; //no problem, green is in this scope, and the ScopedColor::green is a different scope.
auto blueC = ScopedColor::blue; //fine;
auto orangeC = orange; //error! no such thing as orange, only ScopedColor::orange
```

another advantage of scoped enums is that they provide strong typing. and they don't implicitly convert into numerics. un scoped enums can participate in any operation involving numbers, even when it doesn't make sense!

```cpp
enum Color{black,white,red}; //unscoped enum
std::vector<std::size_t> primeFactors(std::size_t x); //function
Color c = red; //enum
if (c < 14.5) //compare color to double!
{
    auto factors = primeFactors(c); //what are the prime factors of red?
}
```

if we use a scoped enum, we no longer have implicit conversions, and if we want to do something weird, that's our choice.

```cpp
enum class Color{black,white,red}; //unscoped enum
std::vector<std::size_t> primeFactors(std::size_t x); //function
Color c = Color::red; //enum
//if (c < 14.5) //error!
if (static_cast<double>(c)<14.5) //explicit casting, will work
{
    //auto factors = primeFactors(c); //also an error!
    auto factors = primeFactors(static_cast<std::size_t>(c)); //again, will work
}
```

a third advantage for scoped enums is that they can be forward declared. truth to be told, unscoped enums can also be forward declared, but there's an issue. the forward declaration of enums requires the size to be known. the underlying default type of scoped-enums is int. but we can set it to a different type.

```cpp
enum class Status1{A=-1,B=9}; //underlying type is int;
enum class Status2: std::uint32_t{A,B,C}; //underlying type is unsigned int for a 32 bit machine;
```

anyway, we can forward declare enums, if we want to.

there is one case where uns-coped enums might be preferable, this is for getting elements out of a tuple from the get<> template. there are ways to overcome the issue,though.

```cpp
using UserInfo = std::tuple<std::string, std::string, std::size_t>; //type alias, fields are name, email, reputation
userInfo uInfo;
auto emailValue = std::get<1>(uInfo); // 1 is the index of the email.
enum UserInfoFields {Name, Email,Reputation}; // unscoped enum
auto repValue = std::get<Reputation>(uInfo); // this is good for readability.
enum class ScopedInfoFields{NameScoped, EmailScoped,ReputationScoped}; //scoped enum
//auto nameValue = std::get<ScopedInfoFields::NameScoped>(uInfo); // error! we need std::size_t
auto nameValue = std::get<static_cast<std::size_t>(ScopedInfoFields::NameScoped)>(uInfo); // this works, but so much typing.
```

if we want to avoid the long typing, we can write a function, but it must be known in compile-time, so a constexpr, and while we're here, let's template it for any type of enum and return type. we can even use auto in c++14 to reduce the weird parts of the code!

```cpp
// most basic form
constexpr std::size_t GetUserField(ScopedInfoFields field)
{
    return static_cast<std::size_t>(field);
}
// c++11 form
template<typename E>
constexpr typename std::underlying_type<E>::type // return type
ToUType(E enumerator) noexcept
{
        return static_cast<std::underlying_type<E>::type>(enumerator);

}
// c++14 form v1
template<typename E>
constexpr std::underlying_type_t<E> ToUType141(E enumerator) noexcept
{
    return static_cast<std::underlying_type_t<E>>(enumerator);
}

// c++14 form v2
template<typename E>
constexpr auto ToUType142(E enumerator) noexcept
{
    return static_cast<std::underlying_type_t<E>>(enumerator);
}

//usage
auto nameValue = std::get<ToUType142(ScopedInfoFields::NameScoped)>(uInfo);
```

the extended form still requires more typing, but it's worth it.

#### Things to Remember

> - C++98-style enums are now known as unscoped enums.
> - Enumerators of scoped enums are visible only within the enum. They convert
>   to other types only with a cast.
> - Both scoped and unscoped enums support specification of the underlying type.
>   The default underlying type for scoped enums is int. Unscoped enums have no
>   default underlying type.
> - Scoped enums may always be forward-declared. Unscoped enums may be
>   forward-declared only if their declaration specifies an underlying type.

</details>

### Item 11: Prefer Deleted Functions to Private Undefined Ones

<details>
<summary>
Explicitly deleting functions improves readability by conveying intent, moves failure form link-time to compile-time, gives out better error messages, and can be done on non member functions and template function specializations
</summary>

Usually, no function declaration means that there is no function to call,but sometimes things aren't so simple. c++ has some _'special member functions'_, which are automatically generated when they are needed. [Item 17]() introduces the concept in greater details, but for now we will focus on the copy constructor and copy assignment operator.

in classic C (c++98) the way to prevent those functions from being called was to declare them as private and not define them. this was done for classes and objects in the library where it was not clear what copying them means.

here is code for the basic io stream:

```cpp
template<class charT, class traits= char_traits<CharT>>
class basic_ios:public ios_base{
    public:
    //...
    private:
    basic_ios(const basic_ios&); //not defined
    basic_ios& operator=(const basic_ios&); //not defined
}
```

making the function private means they can't be called from outside, and not defining them means that even if a member function tries to call them, it'll cause an linking error.

in modern c++, we can to something better, mark them both as _'deleted functions'_

```cpp
template<class charT, class traits= char_traits<CharT>>
class basic_ios:public ios_base{
    public:
    //...

    basic_ios(const basic_ios&) = delete; // deleted
    basic_ios & operator=(const basic_ios&) = delete; // deleted
}
```

Deleting functions isn't just a stylistic choice. any file trying to call the functions will fail to compile, so we moved our error detection closer. by convention, deleted functions should be public, this is done because compilers might check for accessability before and report the function is private (which is an important detail, but not informative) rather than that it was deleted (which is the critical reason for the failure). **in general** public functions provide better error message from compilers.

#### Not Just Member Functions

An additional advantage of deleting functions is that while only member functions can be made private, any function can be deleted. this means we can restrict type conversions by providing overloads and deleting them.

```cpp
bool isLucky(int number);

if (isLucky('a')) // char can become int
{

}
if (isLucky(true)) // bool can become int
{

}
if (isLucky(3.5)) // should we truncate to 3? who makes this choice>
{

})
```

all of the calls above are possible, but if we want to block them, we can do so with deleted functions

```cpp
bool isLucky(char) = delete; //no more char
bool isLucky(bool) = delete; // no more boo;
bool isLucky(double) = delete; //no more double, or float. float prefers to become double
```

now all of the calls will fail during compilation. the functions don't exist, but they are part of the overload resolution process. we can do something similar with templated functions. let's say we have a function that process pointers of a generic type, and we never want to call it with a void* pointer (which can't be dereferenced) or a char* pointer (which should be handled by sting operations). we can simply provide deleted template specializations.

```cpp
template<typename T>
void processPointer(T* ptr);

template<>
void processPointer<void>(void* ptr) = delete;

template<>
void processPointer<char>(char* ptr) = delete;
```

we can go a step further and delete const volatile void* and const volatile char* overloads, or other types of character types (std::wchar_t, std::char16_t, std::char32_t)

in the case of function templates inside a class, we couldn't disable them by using the private method, because all function template specializations have the same access modifier. however, we can delete templates functions specializations, although the process requires deletion from outside the class scope.

```cpp
class Widget
{
    public:
    //...
    template<typename T>
    void processPointer(T* ptr)
    {
        //...
    }
    private:
    /*
    template<>
    void processPointer<void>(void* ptr) // this can't be done
    {
        //...
    }
    */
};
template<>
    void Widget::processPointer<void>(void* ptr) = delete; // this can be done.
```

#### Things to Remember

> - Prefer deleted functions to private undefined ones.
> - Any function may be deleted, including non-member functions and template
>   instantiations.

</details>

### Item 12: Declare Overriding Functions _override_

<details>
<summary>
Derived class overrides are important to get right, but easy to get
wrong. explicit override means we make sure we get them right.
</summary>
Object oriented programming in c++ revolves around classes, inheritance and virtual functions. but the number of ways we can fail to properly override a virtual function in a derived class is surprisingly large.

_overriding_ and _overloading_ sound similar and are parts of polymorphism, but aren't the same.
below is some code example

```cpp
class Base{
    public:
    virtual void doWork();
    //...
};
class Derived : public Base
{
    public:
    virtual void doWork(); //override, "virtual" is optional.
};
std::unique_ptr<Base> upb = std::make_unique<Derived>(); //unique ptr
upb->doWork(); //function call through the virtual function,
```

for an override to occur, the following must happen

- the base class function must be virtual
- the based and derived classes function names must be identical(this doesn't apply to destructors)
- the parameter types of the base and the derived functions must be identical.
- the _const-ness_ of the base and the derived functions must be identical
- the return types and exception specifications of the base and derived functions must be compatible.

and in modern c++, we also get an additional requirement

- the function _reference qualifiers_ must be identical.

```cpp
class Widget{
    public:
    //...
    void Foo() &; //this version only works when *this Widget is lvalue
    void Foo() &&; //this version only works when *this Widget is rvalue
}
Widget makeWidget(); //some factory function
Widget w; //lvalue widget
w.Foo(); //lvalue version
make.Widget().Foo(); //rvalue version
```

if we fail on one of those conditions, the 'overriding' function will still exists in the code, but it won't be used when called through a base class pointer. these problems can be hard to trace, and there is no compilation or runtime error, all that happens is that we call the wrong code.

here is an example of legal code, but still not what we wanted.

```cpp
class Base
{
    public:
    virtual void mf1() const;
    virtual void mf2(int x);
    virtual void mf3() &;
    void mf4() const;
};
class Derived
{
    public:
    virtual void mf1(); // oops, no const! this is an entirely new virtual function.
    virtual void mf2(unsigned int x); //oops, int and unsigned int, another new virtual function.
    virtual void mf3() &&; //this function is only for rvalue *this*, sorry, new virtual function again
    void mf4() const; // it wasn't a virtual function then, and it isn't one now.
};
```

the compiler might warn you about the issues above, and it might not. it might catch all, some, or none of them, and it might be lost inside a long list of other ignored warnings.

the override keyword will make sure an function that is declared to override another function does so. if it's doesn't its a compile-time error. in the example above, adding _override_ to the derived class will reveal that they aren't overriding any function (even mf4!). and if we have the override specifier, we can change the base class signature and then the compiler will tell us where all of the overriding functions in the derived classes are.

the _override_ and _final_ keywords are **contextual keywords**, they are reserved only in certain cases. they don't mess with old C legacy code that happens to use those names.override specifies a function overrides a function in base class, final declares no further overrides of the function are allowed in derived classes, or that the class cannot be derived from.

#### Member Function Reference Qualifiers

if we want a function to only accept lvalues, we declare it to take non-const references, if we want an rvalue only, we use the rvalue reference parameter.

```cpp
void Foo(Widget & w); //only lvalue
void Foo(Widget && w); //only rvalue
```

the reference qualifiers relate to the calling object,they aren't as popular as 'const' qualifiers, but they can be useful, imagine an object with a vector as a member variable, and a method that returns a reference to that vector.

```cpp
class Widget{
    public:
    using DataType = std::vector<double>; //alias deceleration
    //...
    DataType & data(){return values;} // returns a reference to the vector.
    private:
    DataType values;
};
Widget w;
auto vals1 = w.data(); // copy w.values into val1;
Widget makeWidget(); //factory function
auto vals2 = makeWidget().data(); // again, copy values, even though we could have used move semantics instep
```

by using reference qualifiers, we can have a correct behavior for this case.

```cpp
class Widget{
    public:
    using DataType = std::vector<double>; //alias deceleration
    //...
    DataType & data() & {return values;} // returns a lvalue reference to the vector.
    DataType data() && {return std::move(values);} // returns a rvalue.
    private:
    DataType values;
};
```

#### Things to Remember

> - Declare overriding functions override.
> - Member function reference qualifiers make it possible to treat lvalue and
>   rvalue objects (\*this) differently.

</details>

### Item 13: Prefer _const_iterators_ to _iterators_

<details>
<summary>
Using const_iterators is simple, easy, expressive, and correct.
</summary>
*const_iterators* are the STL Equivalent to pointer-to-const. they point to values that may not be modified. we should use const whenever possible, so if we have an iterator that shouldn't modify a value, it should be a const_iterator.

in classic c++, const iterator weren't well supported, if we wanted to add an element to a vector at some position, it was easy using iterators

```cpp
std::vector<int> values;
//...
std::vector<int>::iterator it = std::find(values.begin(),values.end(),1983);
values.insert(it,1998);
```

but if we wanted to be clear, it should be a const iterator, as it's not modifying the values. this is how c++98 programmers might write it. **and it might not even work**.

```cpp
typedef std::vector<int>::iterator IterT; // we didn't have auto
typedef std::vector<int>::const_iterator ConstIterT; // so typedefs, we also didn't have aliases.
std::vector<int> values;
//...
ConstIterT = std::find(static_cast<ConstIterT>(values.begin()),static_cast<ConstIterT>(values.end()),1983);
values.insert(static_cast<IterT>(it),1998); //may not compile
```

we need to cast the end and begin iterators, because in c++98, we couldn't get a const*iterator for a non-const container. there were other ways besides casting. we then need to cast back to iterator before calling the insert function, and even that's not guaranteed to work, there was no portable conversion from const_iterator to an iterator. it didn't make sense to go through all of the trouble, so const was used \_when practical* (we would want to use it _when possible_).

c++11 made things easier, _cbegin()_ and _cend()_ produce const*iterators even from non-const containers, the \_insert()* and _erase()_ functions can work with const iterators, and we even got _auto_ to reduce typing. now using const_iterators is just as easy as using regular iterators.

```cpp
std::vector<int> values;
//...
auto it = std::find(values.cbegin(),values.cend(),1983);
values.insert(it, 1998);
```

things get a bit muddier for generic code, if we want to use this approach on container-like data-structures. some classes offers those functions as non-member functions, so our generic code will look like this in **c++14**.

```cpp
template<typename T,typename V>
void findAndInsert(C& container,const V& targetVal,const V& insertVal)
{
    auto it = std::find(std::cbegin(container),std::cend(container),targetVal); // using non member functions
    container.insert(it,insertVal);
}
```

in **c++11 this will not work**, only _end()_ and _begin()_ were added as free function _cbegin(),cend(),rbegin(),rend(),crbegin(),crend()_ weren't added at that time. we could hide the fact with some templates.

```cpp
template<class C>
auto cbegin(const C& container)->decltype(std::begin(container))
{
    return std::begin(container);
}
```

because of how calls are resolved, a templated function will only be called if there is no explicit function or specialized template, so containers with _cbegin()_ will use their own function. this template takes advantage of the fact that c++11 ensures that calling _begin()_ on a const container will return a const_iterator, so this will work for const containers as well. c++11 also returns the first element pointer when the function is called on an array, and the same const rules apply.

#### Things to Remember

> - Prefer const_iterators to iterators.
> - In maximally generic code, prefer non-member versions of begin, end,
>   rbegin, etc., over their member function counterparts

</details>

### Item 14: Declare Functions _noexcept_ If They Won’t Emit Exceptions

<details>
<summary>
If we know that a function doesn't throw exceptions, mark it as such, it can generate benefits.
</summary>

in the age of c++98, exceptions were a challenging issue, libraries were expected to detail which exceptions are thrown, and client code might break if something changed. the compilers offered no help in maintaining consistency and things were chaotic.

When c++11 rolled around, the general opinion changed and it was decided that the important thing about functions and and exceptions was whether or not they had them. A function can guarantee that it doesn't throw exceptions, or it can't guarantee it and might throw exceptions. the c++98 exception specifiers are still legal code, but are deprecated. only the _noexcept_ specifier is relevant.

specifying functions _noexcept_ helps users write better code, just as much as making a function _const_ does. not only that, it also helps compilers optimize code and generate better programs.

```cpp
int f1(int x) throw(); //no exceptions, c++98 style
int f2(int x) noexcept; // no exception, c++11 style
```

in the c++98 style, an exception thrown from f1 will result in call stack un-wounding, and eventually termination. in c++11 style, an exception from f2 _might_ un-wound the call stack, and it _might not_ do so before termination. because there is no guarantee of un-wounding the call stack, compilers aren't required to keep it in a state that it is possible to un-wound from. there is no requirement to maintain the objects and destroy them in an inverse order of construction.

c++98 _throw()_ doesn't offer that kind of flexibility.

```cpp
void foo1() noexcept; //optimizable code.
void foo2() noexcept(true); //optimizable code. explicit usage

void foo3() throw(); //less optimizable code - change to modern style
void foo4(); //less optimizable code - decide if there are exceptions possible
void foo5() noexcept(false); //less optimizable code - we know we might throw exceptions
```

The case for _noexcept_ is even stronger in some cases, like move semantics. imagine this c++98 code. declare a vector, and push another element to it.

```cpp
std::vector<Widget> vw;
//... code happens here
Widget w;
vw.push_back(w);
```

When we add an element to a vector, we first check if there is enough space for it (size == capacity), and if not, there is in internal request for more memory, and than transferring the elements to the new chunk of memory. in the age of c++98, this was done with all copying elements first and then destroying them in the old data. this was was done to ensure that **push_back()** offered strong exception safety guarantee. if something failed in the copying process, the original data was unchanged.  
in c++11, we would want to make use of move semantics rather than copying, but this runs the risk of violating that guarantee, if we failed to move one of the elements, the original state of the data was already changed. because of this reason, c++11 couldn't change all of the resize operations, but if we ensure it that there are no exceptions possible, the vector can use the move semantics instead, with the driving idea be _"move if you can, but copy if you must"_. this holds true for other function with strong exception guarantee (std::vector::reserve, std::deque::insert, etc...). If the compiler is assured that there is no possibility of throwing an exception, it can use the faster move operations. if the required function is _noexcept_, we get better results.

**swap** functions also benefit from _noexcept_. swaps are used in many STL algorithms, so optimizing them would have cascading benefits. many types in the STL have the swap function _noexcept_ depending on the user type.

```cpp
//array swap
template<class T, size_t N>
void swap(T (&a)[N], T (&b)[N]) noexcept(noexcept(swap(*a,*b)));
//pair swap
template<class T1, class T2>
struct pair{
void swap(pair & p) noexcept(noexcept(swap(first,p.first)) &&
                             noexcept(swap(second,p.second)));
}
```

these functions are _conditionally noexcept_, if the expressions inside the clauses are noexcept, then they noexcept. _noexcept_ takes a boolean argument, default to false, so if we check the noexcept of some other function, we can 'inherit' that value from it.

#### The Danger

it's true that _noexcept_ functions can be pretty beneficial, but if we make them such, it's a commitment to keep them this way. client code might depend on classes and functions being _noexcept_ and changing them might break something in the long run.
most functions are, and should be, exception neutral. they might not throw exceptions by themselves, but they call other functions, and those functions might throw. going out of the way to force a function to be _noexcept_ (trying to catch all exceptions internally, returning error codes) can make the code less readable, heavier and have more code branches, this might effect the negatively performance more than any gains from the _noexcept_ optimizations.

some functions should be _noexcept_, the memory de-allocator functions (operator delete and operator delete[]) and destructors are all implicitly _noexcept_. the only way for a destructor not be implicitly _noexcept_ is if some data member destructor is _noexcept(false)_, and throwing an exception from a destructor is undefined behavior.

#### Wide and Narrow Contracts

functions with wide contracts have no pre-conditions, they can be called regardless of the state of a program, and impose no constraints on the argument passed to it. functions with wide contracts never exhibit undefined behavior.
functions without wide contracts are 'narrow contracts', if a precondition was violated, results are undefined.

when we write functions, if the function has a wide contract, and we know it doesn't emit exceptions (or call functions that emit them), we can safely declare it _noexcept_.
functions we narrow contracts are a tricker issue. say we have a precondition, we don't check it in the functions, as it is assumed to caller checked before hand. what do we do if the preconditions wasn't satisfied?

regarding the issue of compilers helping us identify issues, this is legal code.

```cpp
void setup();
void cleanup();
void doWork() noexcept
{
    setup();
    //..
    cleanup();
}
```

the doWork is declared noexcept even though setup and cleanup aren't. it might be that the documentation for them ensures no exceptions, maybe they were written in C or in C++98 style. the compiler will not warn us about this.

#### Things to Remember

> - noexcept is part of a function’s interface, and that means that callers may
>   depend on it.
> - noexcept functions are more optimizable than non-noexcept functions.
> - noexcept is particularly valuable for the move operations, swap, memory
>   de-allocation functions, and destructors.
> - Most functions are exception-neutral rather than noexcept.

</details>

### Item 15: Use _constexpr_ Whenever Possible

<details>
<summary>
Using constexpr objects and functions maximize the range of situations in which we can use our objects, this can lead to better performance and cleaner code.
</summary>
Constexpr is a confusing keyword, when applied to a objects, it's an extreme version of const. but when applied to a function, it has a different meaning. knowing what is possible with constexpr and how to use it can yield benefits, so it's advised to make the effort to use it properly.

Constexpr indicates something that's not only constant and not changing, but also something that's known during compilation time. in terms of functions, things are nuanced. the results of a constexpr function don't have to be const, or even known during compilation. and those are good features of constexpr.

#### Constexpr Objects

Constexpr values are const which are const, and their value is known during compilation (actually, translation, which is part compilation part linking). this makes them privileged values, they can be placed in the read-only memory and they can be used when c++ requires _integral constant expressions_, like specifying array sizes, std::arrays, integral template specifiers, enumerator values, memory alignment specifiers, and more. Declaring a variable constexpr ensures that the value is known during compile-time.

```cpp
int sz; // not a constexpr value
//constexpr auto arraySize =s1; //error! sz isn't know at compile time!
//std::array<int, sz>; //error! same!
constexpr auto arraySize2 =10
std::array<int, arraySize2>;
</details>
```

even using the const keywords doesn't save us

```cpp
int sz; // still not constexpr;
const auto arraySize =sz; //no problem, arraySize is a const copy of sz;
//std::array<int, arraySize>; //error! must be known during compile time!
```

All constexpr objects are const, but not all const objects are constexpr. if we want to ensure the compiler can use a value during compile time, we use the constexpr keyword.

#### Constexpr Functions

Constexpr functions produce compile-time constants when they are called with compile-time constraints. if we call them with constexpr values, we get compile time constants, if we call them we values that depend on runtime, we won't get the value until runtime.

> - constexpr functions can be used in contexts that demand compile-time constants. If the values of the arguments you pass to a constexpr function in such a context are known during compilation, the result will be computed during compilation. If any of the arguments’ values is not known during compilation your code will be rejected.
> - When a constexpr function is called with one or more values that are not known during compilation, it acts like a normal function, computing its result at runtime. This means you don’t need two functions to perform the same operation, one for compile-time constants and one for all other values. The constexpr function does it all.

lets try writing a constexpr version of std::pow

```cpp
constexpr int pow(int base, int exp) noexcept //constexpr don't throw
{
    //... implementation below
}
constexpr auto numConditions =5;
std::array<int,pow(3, numConditions)> results;
```

constexpr function can work in both compile time and runtime.

the constexpr function was expanded between c++11 and c++14. in c++11,we were limited to one statement, the return statement. so our code would have looked like this.

```cpp
constexpr int pow(int base, int exp) noexcept
{
    return (exp == 0 ? 1: base *pow(base,exp-1));
}
```

in c++14, restrictions were looser, and we could write something like this. we can declare variables, and even use loops, but the function is still limited to taking and returning **literal** types.

```cpp
constexpr int pow(int base, int exp) noexcept
{
   auto result = 1;
   for (int i=0; i<exp; ++i)
   {
       result *=base;
    }
   return result;
}
```

literal types are objects whose values can be determined during compilation, so we can have constexpr constructors and then use those objects in constexpr context.

```cpp
class Point
{
    public:
    constexpr Point(double xVal=0,double yVal=0) noexcept: x(xVal),y(yVal)
    {

    }
    constexpr double xValue() const noexcept {return x;}
    constexpr double yValue() const noexcept {return y;}

    void setX(double newX) noexcept {x= newX;}
    void setY(double newY) noexcept {y= newY;}

    private:
    double x,y;
};
```

the Point constructor can be made constexpr because the arguments are known at compile time. so we can use Point as a constexpr value

```cpp
    constexpr Point p1(9.4,27.7)
    constexpr Point p2(28.8,5.3)
```

because the getters are also constexpr values, they can be used in constexpr context.

```cpp
constexpr Point midPoint (const Point &p1, const Point &p2) noexcept
{
    return{(p1.xValue()+p2.xValue()/2),(p1.yValue()+p2.yValue()/2)};
}
constexpr auto mid = midPoint(p1,p2);
```

now we have mid as a read only value, so we can use it in all sorts of places which a known value must be used!like getting template or specifying values of enumerators. things which we previously runtime only are now in the realm of compile time. this means faster runtime (but unfortunately, slower compilation).

in c++11,we cant make the setters constexpr functions, as it was required for constexpr member function to be const, and they must return a literal value (which void isn't). in c++14, we can make the setters constexpr functions.

```cpp
class Point{
//...
constexpr void setX(double newX) noexcept {x=newX;}
constexpr void setY(double newY) noexcept {y=newY;}
};
constexpr Point reflection(const Point & p) noexcept
{
    Point result;
    result.setX(-p.yValue());
    result.setY(-p.xValue());
    return result;
}
```

as with noexcept, making a function constexpr means a commitment, we can't easily scale back without possible making a mess of client code that decided to use the function in constexpr context. even adding debugging IO statements are not usually permitted in compile time context.

#### Things to Remember

> - constexpr objects are const and are initialized with values known during compilation.
> - constexpr functions can produce compile-time results when called with arguments whose values are known during compilation.
> - constexpr objects and functions may be used in a wider range of contexts than non-constexpr objects and functions.
> - constexpr is part of an object’s or function’s interface.

</details>

### Item 16: Make _const_ Member Functions Thread Safe

<details>
<summary>
If a const member function might be accessed in a multi-threaded environment, and it's possible to get a data race inside it, we should ensure that the function is safe to call.
</summary>
an example of caching polynomial roots. we don't change the values, so it can be a const function, but we don't want to repeat this action, so we cache the results for later requests by using mutable members.

```cpp
class Polynomial{
    public:
    using RootsType= std::vector<double>;

    RootsType roots() const
    {
        if (!rootAreValid) // check if valid
        {
            //.. do the work, cache the values
            rootsAreValid=true; //declare valid for next time
        }
        return rootVals;
    }


    private:
    mutable bool rootAreValid{false}; //mutable members
    mutable RootsType rootsVals{};
};

```

the code above is safe and correct, but what will happen in a multi-threaded environment? we can reach a situation of a data-race without synchronization! this can be solved with a mutex (also declared mutable) and using _std::lock_guard\<std::mutex>_ but because mutex is a move only object, our entire object loses the capability to be copied.

mutex isn't always the correct way to go, even we simple want to count something, an atomic counter should fit our needs.

```cpp
class Point {
public:
double distanceFromOrigin() const noexcept
{
    ++callCount; // increment atomic counter
    return std::sqrt((x*x)+(y*y));
}
private:
mutable std::atomic<unsigned> callCount{0};
double x,y;
};
```

however, _std::atomic<>_ is also a move-only type, and one might overuse them when it's not appropriate.

```cpp
class Widget
{
public:
int magicValue() const
{
    if (cacheValid) return cachedValue;
    else
    {
        auto val1 = expensiveComputation1();
        auto val2 = expensiveComputation2();
        cachedValue = val1+val2; //bad!
        cacheValid = true; //bad!
        return cachedValue;
    }
}
private:
mutable std::atomic<bool> cacheValid{false};
mutable std::atomic<int> cachedValue;
};
```

we still get a data race,if two (or more!) threads reach this position before the atomic bool is changed, we will get repeated computations. we get even worse behavior if we switch the order of assignments. if a context switch occurs after we changed to bool value to true, but before the values were stored, we get not just bad performance, but actually bad results.

```cpp
class Widget
{
public:
int magicValue() const
{
    if (cacheValid) return cachedValue;
    else
    {
        auto val1 = expensiveComputation1();
        auto val2 = expensiveComputation2();
        cacheValid = true; //now it's worse
        // a context switch might happen here!
        cachedValue = val1+val2; //now it's worse
        return cachedValue;
    }
}
};
```

if we want to use a single value, we can make with _std::atomic<>_, but two or more variables require the use of a mutex.

```cpp
class WidgetMutex
{
public:
int magicValue() const
{
    std::guard<std::mutex> g(m); //lock m
    if (cacheValid) return cachedValue;
    else
    {
        auto val1 = expensiveComputation1();
        auto val2 = expensiveComputation2();
        cachedValue = val1+val2; //order doesn't matter!
        cacheValid = true; //order doesn't matter
        return cachedValue;
    }
}
private:
mutable std::mutex m;
mutable bool cacheValid{false};
mutable int cachedValue;
};
```

#### Things to Remember

> - Make const member functions thread safe unless you’re certain they’ll never
>   be used in a concurrent context.
> - Use of std::atomic variables may offer better performance than a mutex, but
>   they’re suited for manipulation of only a single variable or memory location.

</details>

### Item 17: Understand Special Member Function Generation

<details>
<summary>
Know When the member functions are generated and when not. difference between move and copy operations, using the =default keyword,
</summary>

In c++, there are several member functions that the compiler will generate for us. in c++98, there four functions:

- the default constructor
- the destructor
- the copy constructor
- the copy assignment operator

those functions are generated only if they are needed (called in the code), the default constructor is generated when there is no other constructor declared. all generated functions are implicitly _public_ and _inline_, and non-virtual, except for when a destructor is generated for a derived class inheriting a base class with a virtual destructor.

in c++11, there are two mre generated member functions:

- the move constructor
- the move assignment operator

```cpp
class Widget
{
    public:
    Widget(Widget && rhs); //move constructor
    Widget & operator=(Widget&& rhs); // move assignment operator
};
```

the same rules as before apply, the functions are generated only if needed, and the perform _member-wise moves_ on the non static-data of the object and it's base classes. the move constructor and move assignment operator, don't necessarily perform actual moving, they perform move when it's enabled,or copy operations when not.  
for the copy constructor and the copy assignment, the two functions are independent of one another, if only one is declared, the other can be generated.  
this isn't true for the move constructor and move assignment, if we declare one of them, the compiler won't generate the other. the reason behind this is that if we say the the default implementation of one those functions is not suitable, then the default implementation of the other one must also be incorrect. on the same rationale, if we declare a copy constructor or a copy assignment, this also disables automatic move operators from being generated. for the same reason, if there is something wrong with piece-wise copy, there must be something wrong with piece-wise moving. this applies in other direction, declaring a move operator disables generation of copt operators.

if we define a destructor, it means that we need to manage some resources, which should mean that member-wise copying shouldn't work. this reasoning is solid, but wasn't enforced in code by c++98.
this leads us to the classic \*_Rule of Three_ from c++98. if we declared one of the following: copy constructor, copy assignment operator or destructor, we should declare all of them.

because of how c++98 operated, the language doesn't enforce the _rule of three_ even today, but does enforce limitations on the newer features. so move operations are generated when

- they are needed.
- no copy operations are declared in the class
- no mover operations are declared
- no destructor is declared

perhaps in the future, some c++ version will extend the rules for copy operations, therefore invalidation older legacy code. the fix for this is fairly easy,all we need is to add a default operations for any class with user defined operations.

```cpp
class WidgetClassic
{
    public:
    ~WidgetClassic(); //user defined destructor
    WidgetClassic(const WidgetClassic&) = default; // use default copy constructor.
    WidgetClassic& operator=(const WidgetClassic &) = default; // use default copy assignments
    //...
};
```

this is useful in polymorphic base classes (which have virtual destructors). a user-defined destructor disables move operations,so we might need additional _=default_ declarations.

```cpp
class Base
{
    public:
    virtual ~Base() = default; // default virtual destructor.
    Base(const Base &) = default;
    Base& operator=(const Base&) = default; //copy operations

    Base(Base &&) = default;
    Base& operator=(Base&&) = default; //move operations
};
```

even if we don't have to declare the functions (in cases the compiler is willing to generate them for us), it might be the smart idea to _=default_ explicitly. this conveys intent and allows to avoid some bugs that might appear down the way.

in the following example, the rule of three is followed, we don't need any special functionality, so we don't define them and leave it to the compiler.

```cpp
class StringTable
{
    public:
    StringTable(){} // user defined default constructor
    //... other functions, but no copy/move/destructor.
    private:
    std::map<int, std::string> values;
};
```

but if in the future we want to add some functionality to the destructor, we have the **side effect of losing the move operations!** (not the copy operations).

```cpp
class StringTable
{
    public:
    StringTable(){
        makeLogEntry("Creating StringTable object");
    } // user defined default constructor
    ~StringTable()
    {
        makeLogEntry("Destroying StringTable object");
    } // now we have a user defined destructor

    //... other functions, but no copy/move
    private:
    std::map<int, std::string> values;
};
```

the issue is that this change is silent. the class is now not move-enabled, so the copy operations will be used instead. this is not likely to fail any tests or cause a difference in behavior, but it is a great hit to performance. before, move operations were possible so the contained members were moved as needed, but now we need to copy them, which means string copying! this is slower in a magnitude of times. and could have been avoided had we declared the operators (moves and copies) to be default.

the rules for the default constructor and the destructor have only slightly changed from c++98. the default constructor is the same, the destructor is now _noexcept_ by default.

|         member function         | C++98                                                             | C++11                                                                                                                          |
| :-----------------------------: | ----------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------ |
|       Default Constructor       | Generated if there are no user defined constructors               | Same as C++98                                                                                                                  |
|           Destructor            | Generated if needed, virtual if base class destructor was virtual | Same as c++98, now _noexcept_ by default                                                                                       |
|        Copy Constructor         | Generated if needed. Member-wise copy of non-static data          | Same as c++98, prevents generations of mover operations, if the copy assignment was defined, this generations is _deprecated_  |
|         Copy Assignment         | Generated if needed. Member-wise assignment of non-static data    | Same as c++98, prevents generations of mover operations, if the copy constructor was defined, this generations is _deprecated_ |
| Move Constructor and Assignment | Not available                                                     | Piece-wise move, generated only if there are no used defined copy operations, mover operations or destructor                   |

a final edge case is that member function Templates **don't** prevent compilers from generating special member functions.

```cpp
class Widget
{
    template<typename T>
    Widget(const T& rhs); //construct Widget from any T

    template<typename T>
    Widget& operator=(const T& rhs); //assign Widget from any T
};
```

the compiler will still generate copy and move operations for Widget (if needed and other conditions are satisfied), despite the template which should count as those functions (when T is Widget) and block that generation. this edge case is further detailed in [Item 26]()

#### Things to Remember

> - The special member functions are those compilers may generate on their own:
>   default constructor, destructor, copy operations, and move operations.
> - Move operations are generated only for classes lacking explicitly declared
>   move operations, copy operations, and a destructor.
> - The copy constructor is generated only for classes lacking an explicitly
>   declared copy constructor, and it’s deleted if a move operation is declared.
>   The copy assignment operator is generated only for classes lacking an explicitly declared copy assignment operator, and it’s deleted if a move operation is declared. Generation of the copy operations in classes with an explicitly declared destructor is deprecated.
> - Member function templates never suppress generation of special member
>   functions

</details>
