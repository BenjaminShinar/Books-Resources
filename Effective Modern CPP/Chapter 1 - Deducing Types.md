## Chapter 1 - Deducing Types

<summary>
Type deduction in different contexts, like templates, auto and decltype.
</summary>

* c++98 had a single set of rules for type deduction, they one for function templates.
* c++11 modified that rule set and added a rules set for *auto* and for *decltype*.
* c++14 extended the scope of auto and decltype. it even adds a cryptic *decltype(auto)*.

these features make code easier to write, as we no longer have to explicitly type and change types across multiple source files, but aso harder to write, as the classes deduced by the compiler might not be what we expect them to be.

### Item 1: Understand Template Type Deduction

<details>
<summary>
who is T in template < typename T>? what is the type of the parameter?
</summary>

> most users of modern C++ use templates without having to know about how the type deduction works, which is a badge of honor for the language, that it simply works.

however, if we wish to understand how *auto* ([item2]("Item 2: Understand auto Type Deduction")) works, we need a deeper understanding of template type deductions.

imagine the following pseudoscope
```
template<typename T>
void f(ParamType param);
f(expr);
```
in the compilation stage, the compiler uses *expr* to figure out both ParamType and T. ParamType can be different from T, as it can contain *adornments*, such as const or reference classifiers. in a real code example:
``` cpp
template<typename T>
void f(const T& param);
int x;
f(x);
```
T is a typename, int in this case, but ParamType is const int reference. T is dependent on ParamType, but it's not always that simple. there are 3 distinct cases:
1. ParamType is a pointer or reference type, but not a *universal reference* (will be explained in item 24).
1. ParamType is a universal reference.
1. ParamType is neither a pointer nor a reference.
  
#### Case 1: ParamType is a pointer or reference type, but not a *universal reference*.

the simplest case, ParamType is a reference type or a pointer type,but not a universal reference. the type deduction works like this:
1. If expr's type is a reference, ignore the reference type.
1. Then pattern-matcher expr's type against ParamType to determine T.

``` cpp
template<typename T>
void f(T& param); // reference
int x = 27;
const int cx = x;
const int & rx = x;
f(x); // T is int, ParamType is int&
f(cx); // T is const int, ParamType is const int&
f(rx); // T is const int, ParamType is const int&, reference is ignored.
```

although the template wasn't defined as taking a const value, it's still possible to pass one to it. the users of the template don't need to worry about const casting to use it.  
in the third example, the reference type of the expression is ignored, and it's treated as const int;
type deduction works the same way for lvalue and rvalue types.

lets change the template to accept const T& argument;

``` cpp
template<typename T>
void f(const T& param); // reference
int x = 27;
const int cx = x;
const int & rx = x;
f(x); // T is int, ParamType is const int&
f(cx); // T is int, ParamType is const int&
f(rx); // T is int, ParamType is const int&, reference is ignored.
```
this time, the const of the arguments is matched against the signature, leaving T as int.

and now with pointers
``` cpp
template<typename T>
void f(T* param); // pointer
int x = 27;
const int *px = &x;
f(&x); // T is int, ParamType is int *
f(px); // T is const int, ParamType is const int *
```
same as before.

#### Case 2: ParamType is a universal reference

now things are less clear. universal templates are declared as T&&, like an rvalue reference, but they behave differently. the full story is in item 24, but a short version follows:
* if expr is lvalue, both T and ParamType are lvalue reference.
* if expr is rvalue, the same rules as case 1 apply.


``` cpp
template<typename T>
void f(T&& param); // universal reference
int x = 27;
const int cx = x;
const int & rx = x;
f(x); // x is lvalue,so T is int&, ParamType is also int&
f(cx); // cx is lvalue,so T is const int&, ParamType is also const int&
f(rx); // rx is lvalue,so T is const int&, ParamType is also const int&
f(27); // 27 is rValue, so T is int, ParamType is int&&
```

>Item 24 explains exactly why these examples play out the way they do. The key point here is that the type deduction rules for universal reference parameters are different from those for parameters that are lvalue references or rvalue references. In particular, when universal references are in use, type deduction distinguishes between lvalue arguments and rvalue arguments. That never happens for non-universal references.


#### Case 3: ParamType is neither a pointer nor a reference.

in this case, we are dealing with a pass-by-value call. param will be a copy of whatever is passed into it. it will be a completely new object.
1. as before, if expr type is reference, ignore the reference type.
2. ignore any const and volatile qualifiers (hereby abbreviated as CV qualifiers)

``` cpp
template<typename T>
void f(T param);
int x = 27;
const int cx = x;
const int & rx = x;
f(x); // T is int, ParamType is also int
f(cx); // T is int, ParamType is also int, const qualifier is ignored
f(rx); // T is int, ParamType is also int, reference part and const qualifier is ignored
```
CV qualifiers (const and volatile are ignored only for) pass by value parameters, other cases might retain this data.

``` cpp
template<typename T>
void f(T param);
const char * const ptr = "Fun with Pointers";
f(ptr); // pass arg of type const char * const. T is char * const; param is also char * const.
```
the parameter deduction will results in a modifiable pointer to a const char. we can change where the pointer points to, but we cannot change the data through the pointer.

#### Array Arguments

Usually, C-arrays decay into pointers when passed into functions, but this case is unique.

``` cpp
template<typename T>
void f(T param);
const char name[] = "j. P. Briggs"; // const char[13]
const char * ptrToName = name; // const char *;
f(name); 
```
we usually treat C-array and pointers the same, but what about the case when we pass it by reference?

``` cpp
template<typename T>
void f(T& param); // pass by reference
const char name[] = "j. P. Briggs"; // const char[13]
const char * ptrToName = name; // const char *;
f(name);  // T is const char[13], so paramType is const char&[13]
```
this means we can do something like this
``` cpp
template<typename T, std::size_t N>
constexpr std::size_t arraySize(T (&)[N]) noexcept
{
   return N; 
}
int main()
{
    int keyVals[] = {1,3,7,8,11,22,35};
    std::size_t site = arraySize(keyVals); // will return the number of elements.
}
```

#### Function Arguments

other stuff can decay into pointers, like function types into function pointers
``` cpp
void someFunc(int, double); // type is void(int,double)
template<typename T>
void f1(T param); // pass by value
template<typename T>
void f2(T & param); // pass by reference
f1(someFunc); // type is ptr-to-func. void(*)(int, double)
f2(someFunc); // type is ref-to-func. void(&)(int, double)
```

we won't see this in practice, but it also exists.

#### Things to Remember

>
* During template type deduction, arguments that are references are treated as non-references, i.e., their reference-ness is ignored.
* When deducing types for universal reference parameters, lvalue arguments get special treatment.
*  When deducing types for by-value parameters, const and/or volatile arguments are treated as non-const and non-volatile.
* During template type deduction, arguments that are array or function names decay to pointers, unless they’re used to initialize references.

</details>

### Item 2: Understand Auto Type Deduction

<details>
<summary>
What is the type of Auto?
</summary>
Auto type deduction follows similar rules to template type deduction, in fact, it's literally an algorithmic transformation to get from one to the other.

``` 
template<typename T>
void f(ParamType param);
f(expr);
```

recall that the types (T, ParamType) are deduced by the type of expr. when we use the auto keyword to declare a variable, auto substitutes the T, and the type specifiers/qualifiers act as ParamType.
``` cpp
auto x = 27; // T is int, ParamType is int
const auto cx = x; // T is int, ParamType is const int
const auto & rx =x; // T is int, ParamType is const int &
```
this is equivalent to:
``` cpp
template<typename T>
void func_for_x(T param);
func_for_x(27); // T is int, ParamType is int

template<typename T>
void func_for_cx(const T param);
func_for_cx(x); // T is int, ParamType is const int

template<typename T>
void func_for_rx(const T & param);
func_for_rx(x); // T is int, ParamType is const int &
```
we see that deducing type for auto is similar to deducing template types, but there is one exception. as before, there are three cases:
1. The Type specifier is a pointer or a reference, but not a universal reference.
1. The type specifier is a universal reference.
1. The type specifier is neither a pointer nor a reference.

``` cpp
auto x =27; // case 3
const auto cx =x; // case 3
const auto &rx =x; // case 1. non-universal reference


auto && uRef1= x; // case 2. x is int, lvalue. uRef1 type is int &;
auto && uRef2= cx; // case 2. cx is const int, lvalue. uRef2 type is const int &;
auto && uRef3 = 27; // case 2. 27 is rvalue int. uRef3 is int &&;
```
the same issue of array and function decay from the previous items continues here.
```cpp
const char name[] = "R. N. Briggs"; // name's type is const char[13]
auto arr1 = name; // arr1 type is const char *, decayed.
auto & arr2 = name; // arr2 type is const char &[13], no decay

void someFunc(int,double);
auto func1 = someFunc; // func1 type is void(*)(int, double)
auto & func2 = someFunc; // func2 type is void(&)(int, double)
```
so far, there was no difference between template type deduction and auto type deduction. but there is one difference.

#### The Only Difference

if we use the four styles of intimidating values with the int type, we get the same result: a variable of type int. However, if we use auto, there is a difference
``` cpp
int x1 = 27; // classic c
int x2(27); // classic c++
int x3 = {27}; // modern c++,uniform initialization
int x4{27}; // modern c++,uniform initialization

auto x1 = 27; // type is int.
auto x2(27); // type is int;
auto x3 = {27}; // type is std::initializer_list<int> with value { 27}
auto x4{27}; // type is int;
```

this is a special rule for auto type deduction. an auto declared variable enclosed in curly braces is of type std::initializer_list and the templated type is deduced from the arguments.

``` cpp
//auto x5 = {1,2,3.0}; // error! cant deduce T for std::initializer_list<>
auto x6 = {1,2,3}; // std::initializer_list<int>
```
apparently, there are two kinds of type deduction here. the first is for the variable x5, which follows the regular rules and is deduced to be what's on the right hand side (std::initializer_list\<T\>), but the template type deduction fails for T.
``` cpp
auto x7 = {11,23,9}; // std::initializer_list<int>
template<typename T>
void f(T param);
f({11,23,9}); // this will fail! can't deduce type T
template<typename T>
void f2(std::initializer_list<T> param);
f2({11,23,9}); // will work, T is int, ParamType is std::initializer_list<int>
```
> So the only real difference between auto and template type deduction is that auto assumes that a braced initializer represents a std::initializer_list, but template type deduction doesn’t.

the Issue continues in c++14. where auto can indicate a function return type, and lambdas may also use auto in parameter declarations. in those cases, the **template type deduction** rules are used, rather than the **auto type deduction** rules. so we can't return an initializer list directly, or use auto it in lambda parameter type specification.
``` cpp
auto createInitList()
{
    return {1,2,3}; // won't work! can't deduce type
}
std::vector<int> v;

auto resetV = [&v](const auto & newValue){v=newValue;}; // c++14.
resetV({1,2,3}); // won't work! can't deduce type
```

#### Things to Remember

>
* auto type deduction is usually the same as template type deduction, but auto type deduction assumes that a braced initializer represents a std::initializer_list, and template type deduction doesn’t.
* auto in a function return type or a lambda parameter implies template type
deduction, not auto type deduction.
</details>

### Item 3: Understand Decltype
<details>
<summary>
How does decltype determine the type? and what is decltype(auto)?
</summary>
Decltype returns the type of a named variable or an expression. but it's not always clear how this was decided. unlike template and auto deduction rules, decltype parrots back the exact type.

``` cpp
const int i =0;
// decltype(i) is const int;

bool f(const Widget & w);
// decltype(w) is const Widget &;
// decltype(f) is bool(const Widget&)

struct Point {
    int x;
    int y;
};
// decltype(Point::x) is int;
// decltype(Point::y) is int;

Widget w;
if (f(w))
{
    //...
}
// decltype(w) is Widget;
// decltype(f(w)) is bool;

template<typename T>
class simpleVector {
    public:
    T& operator[](std::size_t index);
};
simpleVector<int> v;
if (v[0]==0)
{
    // ...
}
// decltype(v) is simpleVector<int>;
// decltype(v[0]) is int&;
```

in c++11, decltype is primarily used when the functions return type depends on the parameter types. in most containers, the operator[] returns a &T (reference to T), but in std::vector\<bool\> (a specialized form a std::vector), it returns a different object. so decltype allows us to capture and express that type.  
this is a crude example, which we will refine later
```cpp
template<typename Container, typename Index>
auto authAndAccess(Container & c, Index i) ->
decltype(c[i])
{
    authenticateUser();
    return c[i];
}
```
the use of auto in this case isn't for type deduction, it's to indicate that we are using trailing return type syntax (the -> after the parameters). in this example, we say that we return whatever using the square brackets operator on the container returns.

in c++11 the return type for single statement lambdas can be deduced and dropped. and in c++14, this is extended to all functions and lambdas, regardless of the number of statements.
```cpp
// c++14, dropping trailing return type, won't work!
template<typename Container, typename Index>
auto authAndAccess(Container & c, Index i) ->
{
    authenticateUser();
    return c[i];
}
```

but, according to item 1, the reference-ness of the expression is ignored, so this code is actually problematic.

```cpp
std::deque<int> d;
authAndAccess(d,5) = 10;// won't compile!
```
the [] operator should return &int, but auto type deduction strips away the reference, which means we try to assign 10 to an rvalue. which is impossible. in order to fix this issue,  c++14 added *decltype(auto)* as a specifier.
```cpp
// c++14, dropping trailing return type but still getting the right value.
template<typename Container, typename Index>
decltype(auto) authAndAccess(Container & c, Index i)
{
    authenticateUser();
    return c[i];
}
```

there are other uses for decltype(auto), not just function return type.
```cpp
Widget w;
const Widget &cw=w;
auto myWidget1 = cw; //auto type deduction. myWidget1 type is Widget.
decltype(auto) myWidget2 = cw; //decltype type deduction. myWidget2 type is const Widget&.
```

now we will refine the function from above.
```cpp
// c++14, dropping trailing return type but still getting the right value.
template<typename Container, typename Index>
decltype(auto) authAndAccess(Container & c, Index i);
```
the container is passed by an lvalue reference to non const, because we want to be able to pass and change the values of the container. but this also means we can't pass an rvalue reference, because we cannot take a changing reference to an rvalue. if the function signature had declared it as const, we would have no problems. but for now, we cant bind rvalue to lvalue references.  
admittedly,this is an edge case, but it's still possible that someone will want to do this:
```cpp
//it won't work
std::deque<std::string> MakeStringDeque(); // some factory function;
auto s = authAndAccess(MakeStringDeque(),5); // make a copy of the 5th element
```
if we wish to allow this behavior, we need to have our *authAndAccess* function support both rvalue and lvalue reference parameters. which means using an **universal reference**. we also need to update the implementation of the function.
```cpp
// c++14, now with a universal reference.
template<typename Container, typename Index>
decltype(auto) authAndAccess(Container && c, Index i)
{
    authenticateUser();
    return std::forward<Container>(c)[i];
}

// c++11, now with a universal reference and training return type
template<typename Container, typename Index>
auto authAndAccess11(Container && c, Index i) ->
decltype(std::forward<Container>c[i])
{
    authenticateUser();
    return std::forward<Container>(c)[i];
}
```
note:
we are using pass-by-value for the Index type, which can potentially create unnecessary copying, but since the standard library does this with index values, we can slack off here.

there are other cases where decltype surprises us, while taking the type of an lvalue expressions such as variables, it's hardly a challenge to reason about the results, some expressions are more complicated.  
decltype should always return a lvalue reference, if any expression other than a name has type T, then decltype will return T&. also, while a variable of Type T is of type T, putting it into parentheses creates and expression of type T&.
this means that if we use decltype(auto), dropping the the parentheses from a return statement can change the return type.there is a [c++ Weekly] episode about this somewhere.

``` cpp
int x; //decltype(x) is int, decltype((x)) is int&.
decltype(auto) f1()
{
    int x =0;
    return x; //f1 returns int;
}
decltype(auto) f2()
{
    int x =0;
    return (x); //f2 returns int &; 
    //this is undefined behavior.
}
```
however, those situations are rare, and decltype generally does what it's supposed to do.

#### Things to Remember


> * decltype almost always yields the type of a variable or expression without any modifications.
> * For lvalue expressions of type T other than names, decltype always reports a type of T&.
> * C++14 supports decltype(auto), which, like auto, deduces a type from its initializer, but it performs the type deduction using the decltype rules.
</details>

### Item 4: Know How To View Deduced types

<details>
<summary>
Three ways to view deduced types
</summary>

the way to view a deduced type depends on the stage where the information is required.
* during the writing of the code.
* during compilation
* during runtime

#### While Writing Code

Some IDEs provide information about types when hovering above them, this should work fine for most simple types, but complicated types might not be revealed.

#### During Compilation

we can use and abuse the compiler #warnings and #error messages to show type information. this can be done by using class templates without defining them.

```cpp
template<typename T>
class TD; // TD stands for Type Display
auto x = foo(); //some function that returns a complicated type
TD<decltype(x)> xType; //now we will get an error with x's type
```

this should result in a compilation error with the type.

#### Runtime

if we reached a situation where we must know the typename during runtime, we need to do some work to get a readable output. we might try to use [typeid](https://en.cppreference.com/w/cpp/language/typeid) and [std::type_info::name](https://en.cppreference.com/w/cpp/types/type_info) to create something like this:
```cpp
std::cout << "type of x is" << typeid(x).name() << '\n';
std::cout << "type of y is" << typeid(y).name() << '\n';
```
but these call results in a mangled typename, mostly i instead of int, P instead of pointer, and K for const. this is entirely dependant on the compiler.  
note: c++filt tool can decode these "mangled" types.  
in more complex types we get even worse names:

```cpp
template<typename T>
void f(const T& param);
{
    std::cout << "type of T is" << typeid(T).name() << '\n';
    std::cout << "type of param is" << typeid(param).name() << '\n';
}
std::vector<Widget> createVec(); //factory function.
const auto vw = createVec();
if !(vw.empty())
{
    f(&vw[0]);
}
```
in the gnu compile T is *PK6Widget* and param is *PK6Widget* as well. if we try to reason about this, then we know PK is pointer to const, and that Widget is somehow involved. but what is 6? and besides, we know that T and param shouldn't be the same. 
``` cpp
int x =7;
f(x);
```
based on template type deduction rules, T should be int, and param should be const int &. but we get different results.  
**this mistake is by design**, the standard requires *type_info::name()* to produce the names as if they were passed into a template by value. which means tha reference is removed, and so are the cv qualifiers (const and volatile) so the following is more accurate.
```cpp
template<typename T>
void f(T param);
{
    std::cout << "type of T is" << typeid(T).name() << '\n';
    std::cout << "type of param is" << typeid(param).name() << '\n';
}
```

actually many IDEs also remove the type parameters data from the output, or shows too much information. if we want a better runtime name, we can use the [TypeIndex library from boost](https://www.boost.org/doc/libs/1_76_0/doc/html/boost_typeindex.html). and choose the function which retains the cv qualifiers and returns a human friendly string representation of the type.
```cpp
#include <boost/type_index.hpp>
template<typename T>
void f(T param);
{
    using boost::typeindex::type_id_with_cvr;
    std::cout << "type of T is" << type_id_with_cvr<T>().pretty_name() << '\n';
    std::cout << "type of param is" << type_id_with_cvr<param>().pretty_name() << '\n';
}
```

#### Things to Remember

> * Deduced types can often be seen using IDE editors, compiler error messages, and the Boost TypeIndex library.
> * The results of some tools may be neither helpful nor accurate, so an understanding of C++’s type deduction rules remains essential.
</details>
