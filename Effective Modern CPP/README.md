# Effective Modern C++

Coming to grips with c++11 and c++14.

## Introduction

it's not just "C++, but more".


c++11 most prominent change is the move semantics with rvalue, lvalue expressions. in a general sense, if you can take the address of it, then it's an lvalue. all parameters are lvalues.
``` cpp
class Widget
{
    public:
    Widget(Widget&& rhs); //rhs is an lvalue, though it has an rvalue reference type.
};
```
the code above declares the move constructor. in the function call, the expressions passed inside the parentheses are *arguments*, they are used to initialize the expressions inside function, where they are called *parameters*. parameters are lvalue, while arguments can be rvalue or lvalue. passing a parameter as an argument while retaining it's type is called *perfect forwarding*.

Well designed functions should be *exception safe*, meaning that they offer the basic guarantee of safety: even if an exception is thrown, no data is corrupted and no resources are leaked. the *strong exception guarantee* assures that if an exception is thrown, the state of the program remains as it was prior to the call.

when the book refers to a *function object*, it usually means an object with the *()* operator. it can also mean something that can be invoked with the syntax of a non-member function call, such as function pointers (from C) or functions. we can also add member function pointers and call the entire thing *callable objects*.  
function objects created through a lambda expression are known as *closures*, but the distinction is not terrible important for the current book. also, there are *function templates* (templates that generate functions),*template functions*(the generated functions), as well as *class templates* and *template classes* (one is the 'generator', the other is the generated class).

things in c++ can be *declared* and *defined*. declarations introduce names and types without giving details. Decelerations are missing storage and implementations.
``` cpp
extern int x; // object deceleration
class Widget; // class deceleration
bool func(const Widget &w); // function deceleration
enum class Color; // scoped enum deceleration.
```
in contrast, definitions provide storage location and implementations details
``` cpp
int x; // object definition
class Widget
{
    //...
}; // class definition
bool func(const Widget &w)
{
    return w.Size() < 10;
    
}    // function definition
enum class Color
{
    Yellow,
    Red,
    Blue
}; // scoped enum definition.
```
a definition can act as both definition and declaration. it's mostly interchangeable, until it's not.

a function signature in this book specifies parameters and return type, it doesn't specify the function or the parameter names, or the exception modifiers or the const values (const, volatile, constexpr). this is not the common definition of a function signature.

modern c++ preserves old code validity (backwards compatibly), so old code will run. however, it does routinely *deprecate* features, which means that they should not be used and be removed shortly, and that they might be completely removed in the future.  
for some operations, the standard defines the result as *Undefined Behavior* (or UB). this means that there is no accepted behavior. this should be avoided like the plague.
* using [] to go beyond the limits of a vector.
* dereferencing an uninitialized iterator
* data races of different threads to the same memory location.

in this book, built in pointers (like C pointers) are called *raw pointers*, while classes that wrap around a pointer are *smart pointers*. they usually overload the dereferencing operators(*operator-> and operator\**), but *std::weak_ptr* is an exception.