## Chapter 6 - Lambda Expressions

<summary>
Like everything in c++, lambdas are great, but they still have some pitfalls associated with them, we should know how to use them and what to be mindful of.
</summary>

lambda expression aren't able to do something that can't be done with regular c++, but they allow things to be comfortable and connivent, this is especially for stl 'if' algorithms (_std::find_if_, _std::remove_if_, _std::count_if_, etc...) and algorithms with a comparison functionality (_std::sort_, _std::nth_element_, _std::lower_bound_, etc..), lambdas are useful for custom deleters for smart pointers and for conditional variables in the multi threaded world. and of course, callback functions. lambda expressions make coding a more pleasant activity (even more than it already is!).

> - A lambda expression it just that: an expression.
>
> ```cpp
> std::find_id(container.begin(),container.end(),
>   [](int val){return 0 < val && val <10;}); // this is the lambda
> ```
>
> - A _closure_ is the runtime object created by a lambda. depending on the capture mode, closure hold copies of or references to the captured data. in the call to _std::find_if_ above, the closure is the objects that's passed at the runtime as the third argument to _std::find_if_.
> - A _closure class_ is a class from which a closure is instantiated. Each lambda causes compilers to generate a unique closure class. The statements inside a lambda become executable instructions in the member functions of its closure class.

we often use lambda to create a closure that's only used as an argument to function, but closures can also be copied, so we can have multiple closures of a closure type corresponding to a single lambda.

```cpp
int x;
auto c1 = [x](int y){return x *y > 55;}; //c1 is the copy of the closure produced by the lambda.
auto c2 = c1; //copy of c1
auto c3 = c2; //copy of c2
```

we usually say lambdas and mean lambdas, closures and closure classes all together, but here we need to be precise. lambdas and closure classes exist during compilation, while closures exists in runtime.

### Item 31: Avoid Default Capture Modes

<details>
<summary>
lambdas can capture all needed variables by either value [=] or reference [&], both should be avoided, as they can lead to dangling data and mutated values.
</summary>

There are two default capture modes in c++11, by reference and by value. default capture by reference can lead to dangling reference, default capture by value makes lures us thinking we don't have the problem of dangling references, which still exists.

#### Default Capture by Reference

capturing by reference causes a closure to contain a reference to the local variable or parameters as it was in the scope where the lambda is defined. if the lambda exceeds the lifetime of that variable, we get a dangling reference. let's pretend we have a container of filters (a predicate that takes an int and returns a boolean value).
we start with a hard coded version, which checks if the number is divisible by 5. we then want to pass the value in runtime.

```cpp
using filterContainer = std::vector<std::function<bool(int)>>; //alias statement
filterContainer filters;

filterContainer.emplace_back([](int value){return value % 5 ==0;}); //adding a filter
int getDivisor(int a,int b);// function,
auto divisor = getDivisor(x,y);
filters.emplace_back([&](int value){return value % divisor ==0;}); //add a filter, capture by reference, implicit
filters.emplace_back([&divisor](int value){return value % (divisor+2) ==0;}); //add a filter, capture by reference, explicit
```

when we leave the current scope, the divisor variable will be destroyed, and we are stuck with a dangling reference, it doesn't matter if we use implicit or explicit capture (although, using explicit capture makes us consider what we capture, and might remind us to consider their lifetime).

this won't be a problem if the lambda is used immediately, and only in the current context (like when called as part of an algorithm), but code has tendency to move around. so it's always a cause for concern.

```cpp
template <typename C>
void workWithContainer (const C& container)
{
    auto calc1 = computeSomeValue1();
    auto calc2 = computeSomeValue2();
    auto divisor = computeDivisor(calc1,calc2);
    using ContElemT = typename C::value_type; //the type of element inside the container

    if (std::all_of(std::begin(container),std::end(container),[&](const ContElemT & value){return value % divisor ==0;}))
    {

    }
    else
    {

    }
}
```

in c++14 we can drop the type and use _auto_ instead of the awful value_type syntax

```cpp
    if (std::all_of(std::begin(container),std::end(container),[&](const auto& value){return value % divisor ==0;}))
    {

    }
    else
    {

    }
```

#### Default Capture by Value

one way to be safer is to capture by value, which works for direct values, but doesn't protect us from capturing pointers (by value), which point to something that can go out of scope before the closure does. this can happen even when the code looks safe, like the following example with the Widget class. just because we aren't writing pointers in modern c++, doesn't mean we don't use them all time.

```cpp
filters.emplace_back([=](int value){return value % divisor;}); //capture divisor by value, copy into the closure. fine.
class Widget{
    public:
    void addFilter() const;
    private:
    int divisor;
};
//cpp file
Widget::addFilter() const
{
    filters.emplace_back([=](int value){return value % divisor;});  // this is actually bad.
}
```

to understand the problem, we need to note that captures apply only to non-static local variables (including parameters) visible in the scope where the lambda is created.
which should mean that divisor should be captured. after all, it's not a local variable. all the following ways fail to compile:

```cpp
    filters.emplace_back([](int value){return value % divisor;});  // divisor not available
    filters.emplace_back([=divisor](int value){return value % divisor;});  // no local divisor to capture
```

so why does the above way compile?

```cpp
    filters.emplace_back([=](int value){return value % divisor;});  // this is actually bad.
```

the secret is that we aren't capturing the divisor, we capture the _this_ pointer that hides inside each member function. so what we actually access is the _this->divisor_, and when the _this_ pointer goes out of scope, we lose the divisor.

```cpp
void Widget::addFilter()const
{
    auto currentObjectPtr = this;
    filters.emplace_back([currentObjectPtr](int value){return value % currentObjectPtr->divisor;});
}
```

now it's much clearer, once the Widget goes out of scope, the ptr becomes dangling. this also happens with smart pointers. in this example, the _std::unique_ptr_ goes out of scope at the end of the function, but the pointer still points to somewhere, but that somewhere is already long gone from the meaning it had originally.

```cpp
void doSomeWork()
{
    auto p =std::make_unique<Widget>();
    pw->addFilter();
}
doSomeWork();
//use filters - oops! dangling!
```

we can solve this by making a local copy inside the scope and then copying from it.

```cpp
void Widget::addFilter()const
{
    auto divisorCopy = divisor;
    filters.emplace_back([divisorCopy](int value){return value % divisorCopy;}); //explicit capture by copy
    //filters.emplace_back([=](int value){return value % divisorCopy;}); default capture by copy, but we already said we should be careful, lets not make this mistake again
}
```

in c++14 we have a better way to do this, with generalized lambda captures

```cpp
void Widget::addFilter()const
{
    filters.emplace_back([divisorCopy = divisor](int value){return value % divisorCopy;}); // now it's clear, we copy the divisor into our lambda captures.
}
```

there is no default capture mode for generalized lambda capture, but even in c++14, we should avoid default capture modes.

even with capture by copy, we still aren't completely insulated from changes that can mutate data that is used inside our lambda. objects with _static storage duration_ can't be captured, but they can be used inside our lambdas. default capture by value makes us think we are safe, but we aren't.

```cpp
static auto divisor = computeDivisor(a,b);
filters.emplace_back([=](int value){return value & divisor == 0;});
//.. many lines later

++divisor;
```

even though we captured by default pass by copy, the divisor has a static storage duration, so we are actually accessing it directly, and eventually, some one changes it and our lambda behaves differently.

#### Things to Remember

> - Default by-reference capture can lead to dangling references.
> - Default by-value capture is susceptible to dangling pointers (especially _this_),
>   and it misleadingly suggests that lambdas are self-contained

</details>

### Item 32: Use Init Capture to Move Objects Into Closures

<details>
<summary>
Init capture, also known as generalized lambda capture, allows us to be more explicit about data members of the closure class created from the lambda.
</summary>

when we have the capture list, we might have an object that we need to capture directly (not by reference), but is costly to copy, such as the standard containers. c++11 doesn't have a solution for us, but c++14 provides direct support for moving int the capture list (rather than just copying).

the new capabilities of c++14 aren't just capturing by move, it's much more, and they are called **init capture**, with it, we can specify:

> 1. The name of a data member in closure class generated from the lambda.
> 2. An expression initializing that data member

and to clarify, the lambda generates a closure class, which like any class, can have data members, in regular captures, the data members have the same name as the variables they are initialized from, but now we can give them different names, and give them value based on some expression.

```cpp
class Widget {
    public:
    //...
    bool isValidated() const;
    bool isProcessed() const;
    bool isArchived() const;
    private:
    //...
};

auto pw = std::make_unique<Widget>();
// do something with pw,
auto func = [pw =std::move(pw)]{return pw->isValidated() && pw->isArchived()}; // pw is init captured to move construct pw inside func
```

the expression inside the square brackets is the init capture, we initialize the data member pw with the move constructor of pw. despite both having the same names, the exist in different scopes. the left hand side is in the scope of the closure class,the right hand side is in the scope where the lambda is defined.

if we don't need to do anything with pw before passing it to the lambda, we can create it directly inside the init capture list.

```cpp
auto func = [pw =std::make_unique<Widget>()]{return pw->isValidated() && pw->isArchived()}; // pw is created in the capture list.
```

#### Init Capture Behavior in c++11

even if the init capture isn't part of c++11, we can still get the same results if we are willing to write code by hand.

```cpp
class isValAndArch
{
    public:
    using DataType = std::unique_ptr<Widget>;
    explicit isValAndArch(DataType&& ptr) : pw(std::move(ptr)){}
    bool operator()()const
    {
        return pw->isValidated() && pw->isArchived();
    }
    private:
    DataType pw;
};

auto func = isValAndArch(std::make_unique<Widget>());
```

we can also

> 1. move the object to be capture into a function object produced by _std::bind_.
> 2. give the the lambda the reference to the "captured" object.

we want to create a vector with some values, and then move it into the a closure, in c++14:

```cpp
std::vector<double> data;
//.. add data
auto func = [data= std::move(data)]{/*do something */};
```

in c++11, we create a lambda that takes the data as the parameter (by reference), and from that lambda we create a 'wrap' function object with _std::bind_. this object has the data itself, which was passed to it by moving. when we call func(), we actually call the lambda with the vector stored inside the func object.

```cpp
std::vector<double> data;
//.. add data
auto func =std::bind([](const std::vector<double>&data){/*do something */}, std::move(data));
```

#### Mutable State

the _operator()_ of the the closure is const by default, so all the data members inside
the capture init are effectively const as well. we can use the mutable keyword to mark our lambda as capable of mutating it's data members.
the lifetime of the closure is now tied to bound object.

```cpp
CounterType cnt1{}; //assume this is a movable counter type
CounterType cnt2{}; //assume this is a movable counter type
auto func14 = [cnt = std::move(i1)] () mutable{return ++cnt;};
auto func11 = std::bind([](CounterType& cnt)mutable{return ++cnt;}, std::move(cnt2));
```

some fundamental points about _std::bind_ and closures

> - It’s not possible to move-construct an object into a C++11 closure, but it is possible to move-construct an object into a C++11 bind object.
> - Emulating move-capture in C++11 consists of move-constructing an object into
>   a bind object, then passing the move-constructed object to the lambda by reference.
> - Because the lifetime of the bind object is the same as that of the closure, it’s possible to treat objects in the bind object as if they were in the closure.

[Item 34]() suggests that _std::bind_ should be replaced by lambdas, but if we can only use c++11, the form above is a good enough emulation.

#### Things to Remember

> - Use C++14’s init capture to move objects into closures.
> - In C++11, emulate init capture via hand-written classes or _std::bind_.

</details>

### Item 33: Use _decltype_ on _auto&&_ Parameters to _std::forward_ Them

<details>
<summary>
We can perfect-forward inside a generic lambda.
</summary>

c++14 added _generic lambdas_, lambdas that have _auto_ in their parameter specification. this is done by templates, so the following lambda is roughly equivalent to the class beneath it.

```cpp
auto f = [](auto x){return func(normalize(x));};

class CompiledF {
public:
template <typename T>
auto operator()(T x) const
{
    return func(normalize(x));
}
//..
};
```

in this example, all that we do is pass the argument x into the normalize function, this means that even if we had a rvalue, we treat it as if it was an lvalue, so we might have some unnecessary work done. what we should do is use perfect forwarding. for this, x needs to be a universal reference(type and form), and be passed with _std::forward_.

so, our first guess is something like this. we just need to figure out what the correct type is for the std::forward.

```cpp
auto f = [](auto&& x){return func(normalize(std::forward<??>(x));};
```

according to [Item 28]() we can inspect the type of the universal reference with decltype.we will get an lvalue reference for lvalue arguments, and rvalue reference for rvalue arguments. also, when using _std::forward_, the type argument will be a lvalue reference for lvalues, and a non-reference for rvalue. if we call _decltype(x)_ on a lvalue, we get a reference (as we should), but if we call it on a rvalue, we get an rvalue-reference, instead of the expected non reference.

```cpp
auto lvalueRef = decltype(lvalueX); //T&
auto rvalueRef = decltype(rvalueX); //T&&, but should be T
```

if we look again at a possible c++14 _std::forward_. we start with the 'normal' template and then look at possible instantiations of it for the Widget class. the normal Type of T is a non reference type.but if it was a rvalue reference, type reference collapsing would resort in the same code

```cpp
template <typename T>
T&& forward(remove_reference_t<T>& param)
{
    return static_cast<T&&>(param);
}

Widget&& forward(Widget& param)
{
    return static_cast<Widget&&>(param);
}

//before reference collapsing.
Widget&& && forward(Widget& param)
{
    return static_cast<Widget&& &&>(param);
}
```

so, even if we get an rvalue form _decltype_, it still collapses into the type which we need for perfect forwarding, hooray! so the correct form is really:

```cpp
auto f = [](auto&&... params){return func(normalize(std::forward<decltype(params)>(params)...);};
```

#### Things to Remember

> - Use decltype on auto&& parameters to std::forward them.

</details>

### Item 34: Prefer Lambdas to _std::bind_

<details>
<summary>
Lambdas add benefits over std::bind in terms of ease to write, readability, clarity and even performance, as of c++14, there is no reason to use std::bind.
</summary>

in c++98, we had _std::bind1st_ and _std::bind2nd_, in 2005 we had the TR1 _std::tr1::bind_, and in c++11 we had a proper _std::bind_. and now, it's perhaps time to abandon this convection in favor of using lambdas.

#### Readability

one case for lambdas over _bind objects_ (the return of _std::bind_) is simply readability. lets start by comparing the two with the usage of a function tha sets an alarm with some sound. we then want to make the alarm go off after an hour and stay for 30 seconds, leaving only the sound up for decision.

```cpp
using Time = std::chrono::steady_clock::time_point;
enum class Sound{Beep, Siren, Whistle};
using Duration = std::chrono::steady_clock::duration;

void setAlarm(Time t,Sound s, Duration d);
//...

auto setSoundL11 =[](Sound s){ // c++11
    using namespace std::chrono; //bring the std::chrono namespace into the function space without qualifiers
    setAlarm(steady_clock::now()+ hours(1),s,seconds(30));
};

auto setSoundL14 =[](Sound s){ // c++14
    using namespace std::chrono; //bring the std::chrono namespace into the function space without qualifiers
    using namespace std::literals; //lets get some literal suffixes
    setAlarm(steady_clock::now()+ 1h,s,30s);
};
```

with lambdas, it's fairly easy to see what's happening, even without being an expert. now, we try the _std::bind_ syntax, along with it's easy to miss pitfalls.

```cpp
    using namespace std::chrono; //bring the std::chrono namespace into the function space without qualifiers
    using namespace std::literals; //lets get some literal suffixes
    using namespace std::placeholders; // for _1
    auto setSoundB = std::bind(setAlarm,steady_clock::now()+ 1h, // incorrect!,
    _1,
    30s
    );
```

the function setAlarm is 'called' in the code, there are no parentheses call to indicate that it's a function. and there is no mention here of the sound parameter, so readers will have to look at what '\_1' is referring to (first argument to SetSoundB). but most importantly, we have a mistake. the time argument is computed once, at the definition time, and won't work as we thought it would. the alarm will always go off 1 hour after we defined the bind object.

we can try to fix this with an additional nested _std::bind_ call, which makes thing more confusing.

```cpp
    using namespace std::chrono; //bring the std::chrono namespace into the function space without qualifiers
    using namespace std::literals; //lets get some literal suffixes
    using namespace std::placeholders; // for _1

    auto setSoundB11 = std::bind(setAlarm,
    std::bind(std::plus<steady_clock::time_point>,steady_clock::now(), 1h)
    _1,
    30s
    );

    auto setSoundB14 = std::bind(setAlarm,
    std::bind(std::plus<>,steady_clock::now(), 1h) // no need to set the type of std::plus in c++14
    _1,
    30s
    );
```

#### Overloading Functions

if the bound function is overloaded, we have a new problem. we add a new overloaded function with an additional argument, and while the lambda works properly, the bind object no longer compiles. we need to cast it into the proper function pointer type.

```cpp
enum class Volume {Normal, Loud, LoudPlusPlus};
void setAlarm(Time t, Sound s, Duration d, Volume v);

using SetAlarm3ParamType = void(*)(Time t,Sound s, Duration d);

auto setSoundB11 = std::bind(static_cast<SetAlarm3ParamType>(setAlarm), //casting to proper form
    std::bind(std::plus<steady_clock::time_point>,steady_clock::now(), 1h)
    _1,30s);
```

#### Inlining

for lambdas, the inner call is easy to inline by the compiler,for the bind object, the function is accessed through a function pointer, which is less likely to be inlined.

```cpp
setSoundL(Sound::Siren); //using lambda, might be inlined.
setSoundB(Sound::Siren); //using bind object, less likely to be inlined.
```

all these issues are worse the more we scale the complexity, in this example, we ask if a value is between a min and max values, now we nest even more calls to _std::bind_, and we still have only a single placeholder.

```cpp
//local variables
const int lowVal = 0;
const int highVal = 100;
auto betweenL14 = [lowVal,highVal](const auto & val){return lowVal <= val && val <= highVal;}; // using auto

auto betweenL11 = [lowVal,highVal](int val){return lowVal <= val && val <= highVal;}; // no auto in c++11
using namespace std::placeholders; // for _1

auto betweenB14 = std::bind(std::logical_and<>(),
std::bind(std::less_equal<>(),lowVal,_1),
std::bind(std::less_equal<>(),_1,highVal)
);
auto betweenB11 = std::bind(std::logical_and<bool>(),
std::bind(std::less_equal<int>(),lowVal,_1),
std::bind(std::less_equal<int>(),_1,highVal)
);
```

#### Placeholders

the placeholders themselves aren't only confusing to read, they also obscure away the type. in this example we a function that creates compressed copies of our widget. we bind our function and the Widget. how are the object and the widget bound? by value or by reference? which is the corresponding lambda? the one that captures by value and copies it into the closure scope, or the one that captures by reference?

```cpp
enum class CompressionLevel{Low, Normal, High};
Widget compress(const Widget &w, CompressionLevel level);

Widget w;
using namespace std::placeholders;
auto compressRateB = std::bind(compress, w, _1); // how is w passed?
auto CompressRateValueL = [w](CompressionLevel lv){return compress(w,lv);}; //by value
auto CompressRateRefL = [&w](CompressionLevel lv){return compress(w,lv);}; // by ref
```

**the answer is that _std::bind_ stores by value** (unless _std::ref_ is used). unlike lambda captures which require the programmer to specify how the the variables are captured. but what about the argument passed to the bind objects itself(not the bound function)? **this time, the answer is by reference**, using perfect forwarding.

#### Cases for _std::bind_ in c++11

despite the points made above, there are cases where _std::bind_ should be used in c++11 (if c++14 isn't available).

one case is for move capture, as c++11 doesn't have init capture, so as detailed in [Item 32](), _std::bind_ can solve the issue.

the other case is **Polymorphic function objects**. the argument passed into bind object call are using perfect forwarding, so this can be useful when we have templated function pointers we wish to bind to, this ability was only added to lambda in c++14 with generic lambdas.

```cpp
class PolyWidget {
    public:
    template <typename T>
    void operator()(const T& param);
    //...
};
PolyWidget pw;
auto boundPW =std::bind(pw,_1);
boundPW(1930);
boundPW(nullptr);
boundPW("RoseBug");

auto PWL14 = [pw](const auto & param){pw(param);}; //c++14 only
```

#### Things to Remember

> - Lambdas are more readable, more expressive, and may be more efficient than
>   using std::bind.
> - In C++11 only, std::bind may be useful for implementing move capture or
>   for binding objects with templatized function call operators.

</details>
