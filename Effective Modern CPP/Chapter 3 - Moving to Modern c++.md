## Chapter 3 - Moving to Modern C++

<summary>
The small parts to remember when using modern C++.
</summary>

the modern versions of c++ define all sorts of high level ideas and features like move semantics, lambdas, concurrency and smart pointers. but also many smaller features, that don't necessary mean big swiping changes, but are important none the less. this chapter deals with them, what they are, and how to build a modern code style.

### Item 7: Distinguish Between () And {} When Creating Objects

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

### Item 8: Prefer nullptr To 0 And NULL

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

### Item 9: Prefer Alias Decelerations to Typedef

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

### Item 10: Prefer Scoped enums to Un-scoped enums

<!-- <details> -->
<summary>

</summary>

<!-- </details> -->
