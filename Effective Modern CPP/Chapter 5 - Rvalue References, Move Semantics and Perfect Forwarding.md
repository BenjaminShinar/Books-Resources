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
