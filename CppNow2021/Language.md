Language

## The C++ Rvalue Lifetime Disaster - Arno Schödl

<details>
<summary>
Using rvalues and lifetime extension can fail us. member accessor for rvalue objects shouldn't return a reference.
const & should never have bound to rvalues.
</summary>

[The C++ Rvalue Lifetime Disaster](https://youtu.be/sb7cj-3l1Kc)

the use of rvalue references and move semantics. replace copying with moves when possible tp avoid memory operations.

also used to manage lifetime, as well as for c++20 ranges

```cpp
auto rng=std::vector<int>{1,2,3} | std::views::filter([](int i){return i%2==0;}); //doesn't compile
```

this doesn't compile for rvalue

### Pitfalls

can't move from a const value, and moving will mess with NRVO (names return value optimization) and make it harder for the compiler to elide the construction.

```cpp
A foo()
{
    const A a;
    return std::move(a); //error!
}
A foo2()
{
    A a;
    return std::move(a); //works, but we are messing with RVO
}
A foo3()
{
    const A a; //doesn't matter if we're const or not, elision works
    return a;
}
```

but if we have two possible values, we can't do NRVO, and we also can't do move (because of const).

```cpp
A foo4()
{
    if ()
    {
        const A a;
        return a;
    }
    else
    {
        const A a;
        return a;
    }
}
```

and here we can't do copy/move ellison, because it's member variable. we also can't do a move, members don't automatically become rvalues.

```cpp
struct B {
    A m_a;
};
A foo()
{
    B b;
    return b.m_a;
    //return std::move(b).m_a; //this would work.
}
```

recommendations:

> - make return variables non-const
> - use clang -Wmove flag

### Temporary Lifetime Extension

```cpp
struct A;
struct B {
private:
A m_a;
public:
const A& GetA() const &
{
    return m_a; //return by reference
}
};
B b;
const auto & a = b.getA();
struct C{
    A getA() const &; // return by value
};
C c;
const auto & a1 = c.getA();
const auto & a2 = B.getA();
```

if we capture something with const reference, it can extended the lifetime of the object it's capturing.

_std::min_ doesn't take rvalue-ness into consideration, it returns a lvalue reference. a will dangle.

```cpp
bool operator<(const A&, const A&);
struct C
{
    A getA() const&;
} ;
C c1,c2;
//...
const auto & a = std::min(c1.getA(),c2.getA()); //a will dangle
```

lets' have a min function that keeps rvalue references using perfect forwarding. but it still doesn't work

```cpp
namespace out
{
    template<typename Lhs,typename Rhs>
    decltype(auto) min(Lhs && lhs,Rhs && rhs)
    {
        return rhs<lhs ? std::forward<Rhs>(rhs)? std::forward<Lhs>(lhs);
    }
}
```

lifetime extension only works where there an object.

an example with forwarding a return and _'decltype(auto)'_

the advice is to stop using temporary life time extension,
what we want is :

> automatically declare variable
>
> - _auto_ if constructed from value or rvalue reference
> - _const auto &_ if constructed from lvalue reference

he suggest this macro code instead of lifetime extension.

```cpp
template<typename T>
struct decay_rvalues
{
    using type = std::decay_t<T>;
};
template <typename T>
struct decay_rvalue<T&>
{
    using type=T&;
};

#define auto_cref(var,...) \
typename decay_rvalue<decltype((__VA_ARGS__))>::type var = ( __VA_ARGS__)'

```

if we add parentheses it's bad, it will always return a reference.

```cpp
decltype(auto) foo()
{
    auto_cref (a, some_a()); // a = some(); with type deduced
    return a; //if we have parentheses, things will be different, it will be a reference.
}
```

theres a debate about whether the macro should return const or not (if not, it can get optimized in NRVO).

```cpp
struct A;
struct B {
    A m_a;
    A const & GetA() const
    {
        return m_a;
    }
}
auto_cref(a1, B().m_a); // B() is rvalue, so it's members are also rvalues;
auto_cref(a2, B().GetA()); // we have a const reference as the return type, so we get a dangling reference const A &;
```

now the problem is that our 'auto_cref' binds to everything, but should rvalues be converted to values?

```cpp
struct A;
A const & L(); //lvalue
A const && R(); //rvalue

decltype(false? L(): L()); // A const &
decltype(false? R(): R());// A const &&
decltype(false? R(): L());// A const, not reference. forces a copy.
```

c++20 has a new trait _common_reference_t_. which was invented for c++20 ranges,

```cpp
std::common_reference_t<A const &, A const &>; //A const &
std::common_reference_t<A const &&, A const &&>; // A const &&
std::common_reference_t<A const &, A const &&>; //A const &. lvalue reference
std::common_reference_t<A const, A const &>; //A. a value
```

so, std::common_reference embraces rvalue amnesia.

### Promises of defences

| Mutability | short Lifetime               | long Lifetime |
| ---------- | ---------------------------- | ------------- |
| immutable  | const &&                     | const &       |
| mutable    | && (can scavenge, move from) | &             |

currently, c++20 reference binding strengths lifetime promise.
from short to long, and from mutable to immutable.

what if we could go the reverse?

> - Allow binding only if promises get weaker
>   - less lifetime
>   - less mutability
>   - less 'scavenge-ability'
>
> * only lvalues should bind to _const &_
> * anything may bind to _const &&_

but we can't allow going from lvalue to rvalue.

### Ideas to Fix the Issue

some things that must hold true before any changes.

where are references used?

> - local/global variable declarations
> - structured binding
> - function/lambda parameter lists
> - members (initialized in PODs)
> - members (initialized in constructors)
> - lambda captures

how it would look with a pragma change. we would need a feature test macro, replace const & parameters with const &&. we will need to change std::common_reference.

</details>

## C++11/14 at Scale: What Have We Learned? - John Lakos & Vittorio Romeo

<details>
<summary>
How we teach C++ versions, which features are suitable for which skill level? what should be taught and when.
</summary>

[C++11/14 at Scale: What Have We Learned?](https://youtu.be/E3JG2Ijjei4)

They're are publishing a book later this year: **Embracing Modern C++ Safety**

> - Why are we talking about C++11/14 in 2021?
> - How C++11/14 an surprise you today
> - C++ at scale
> - "safety" of a feature
> - case study: extended _friend_ declrations

### Why are we talking about C++11/14 in 2021?

adoption rates of C++ standards, some projects are still lagging behind and haven't adapted the newer standads yet, even in 2021 there are places where C++11/14 are just being adopted.

> - There are great learning resources
>   - But most teach "the features" rather than "the experience"
>   - What looks good on paper might not work in the "real world"

not just what the new features are, learn when and how to use them, how they operate inside the bigger context.

### How C++11/14 an Surprise You Today

> Q: "What is the smallest change to the core language you can think of in C++11?"

c++11 changed how double ">" behaved. before c++11 ">>" was parsed as a right shift, so a space was needed to make this recognizable as closing a nested template. in c++11 things were changed and ">>" was now somthing else. this means that a valid c++03 code is invalid in c++11.

in this example, c++03 will see `256 >> 4`, but c++11 will reject this code.

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<256 >> 4> smallBuffer;
```

to fix this issue, we simply wrap the right shift expression in parentheses

```cpp
template <int Power_Of_Two>
struct PaddedBuffer {

};
PaddedBuffer<(256 >> 4)> smallBuffer;
```

in this example, c++03 returns 100, while c++11 returns 0; the compiler gives a warning.

I think that c++03 treats this as a sequence of enums and nested enums (we can change the final 'a' to 'c' and get 102, but not to 'b'). and c++11 treats this as some comparions thing.

```cpp
enum Outer{a=1,b=2,c=3};
template <typename>
struct S {enum Inner {a=100, c=102};};
template <int>
struct G{typedef int b;};
int main()
{

    return S<G< 0 >> ::c>::b>::a;
}
```

other stuff: every one of those can have a dark side.

> - Attributes that can make you code ill-formed NDR.
> - 'extern templates' not improving compilations time or code size at all?
> - Destruction order UB with meyers singletons.
> - Encoding of white space withing raw string literals.

> "NDR - No Diagnostic Required"

### Modern C++ at Scale

how do we teach modern c++? what to prioritize, what kind of approach? how do we integrate the new features into the company style guide? what if we have a tool chain for the style guide? how do we communicate these changes to other teams?

### "Safety" of a Feature

> - Every features of c++ is "safe" when used correctly.
> - But what is the likelihood that it is used correctly?
> - Does the feature have any "attractive nuisance"? (does it invite misuse?)
> - What are the advantages of using a feature compated to its risks?
> - Is it worth teaching to a new hire? to an expereinced hire?

from the book:

> "The degree of safety of a given feature is the the relative likelihood that the widespread use of that feature will have positive impact and no adverse effect on a large software company's codebase."

how likely is teaching and implementing a feature is to go smoothly, be used correctly and give good results, as opposed to being hard to teach/understand, prone to create oppertunities for bugs, hard to maintain by inexperienced workers, or have small scale effects on performance.

three categories of safety:

> - **Safe**
>   - Adds considerable value, easy to use, hard to misuse.
>   - Ubiquitous adoption of such features is productive.
> - **Conditionally Safe**
>   - Adds considerable value, but prone to misuse.
>   - Require in-depth training and additional care.
> - **Unsafe**
>   - Provide value only in the hands of an 'expert', prone to misuse.
>   - Wouldn't teach these as part of genereal c++11/14 course.
>   - Require explicit training on their use cases and pitfalls.

the 'override' keyword is a 'safe' feature. it prevents bugs, makes code self-explantory, and has no real technical downsides.
the one problem that can happen is that people overrely on it, and people except this feature as the norm, and forget that this is just a bonus, you can still have overriding methods without this keyword.
(we can use compiler flags '-Winconsistent-missing-override' and '-Wsuggest-override', but they aren't perfect).

```cpp
class MockConnection : public Connection
{
    void connect(IPV4Address ip) override;
};
```

the 'auto' keyword is _conditionally safe_, it can be great, but can also cause readability problems, and can introduce bugs. _Range based for loops_ are great, but they also have the possibility of bugs, and therefore are marked _conditionally safe_.

in this example, we are actually ok, because we return the vector by value and get lifetime extentsion. this is not true once we decide to be smart and return the vector by reference. now we don't have lifetime extension.

```cpp
class TriggerGetter
{
std::vector<Combo> getCombos() const; //no Issues
const std::vector<Combo> & getCombosRef() const; //oops
};

for (Combo& c : keyBoardTriggerGetters[bindID]().getCombos()) //return by value
{
    //..
}
for (Combo& c : keyBoardTriggerGetters[bindID]().getCombosRef()) //return by const reference
{
    //..
}
```

this is something we overcome in c++20 with init statements, but we must be aware of this issue. it's not an entirely safe action.

_decltype(auto)_ is a strong feature, but it's often misunderstood, misused, a requires training to use correctly, it should be defined as _unsafe_, and only be used with carefull consideration in the codebase. it allows us to deduce the return type from the expression, doesn't strip away qualifiers, returns value or reference objects,and doesn't change anything.

example: higher-order functions.

```cpp
template <typename F>
decltype(auto) logAndCall(F&& f)
{
    log("invoking function ", nameOf<F()>);
    return std::forward<F>(f)();
}
```

but if we teach '_auto_', '_decltype_' and '_decltype(auto)_' together, we push people towards overusing '_decltype(auto)_'.

> - some misconceptions:
> - "If _decltype(auto)_ does everything _auto_ does and more, why not use it all the time?"
> - "If _decltype(auto)_ is more flexible, wht no use it when I'm not sure when to choose between _auto_ and _auto&_?"

> understanding _decltype(auto)_ requires:
>
> - Having a solid grasp on type inference and value categories.
> - Being somewhat experienced with using _auto_ and _decltype_.
> - Having some metaprogrammin expereice.

if you have just learned about _auto_ and _decltype_, you probably aren't in the right level to use _decltype(auto)_ yet.

plus, _decltype(auto)_ has some issues with parentheses surronding it, it's not an easy thing to understand, it can effect SFINAE behavior, so it's not the allways the best tool for the job.

safe features: attributes (most of them), _nullptr_,_static_assert_, digit seperators.
conditionally safe: _auto_, _constexpr_, _rvalue_ references
unsafe: _\[\[carries_dependency]]_, _final_, inline namepaces.

when we teach a new version of c++, we should:

> - Teach _safe_ features early and quickly
>   - Most of them are quality-of-life improvements or hard to misuse.
>   - Trust the student
> - Teach _conditionally safe_ features by building on top of _safe_ knowledge
>   - They require more time and examples.
>   - Show how the can backfire.
>   - Have exercises that make student question whether to use a feature or not.
> - Leave a subset of of _unsafe_ features for self-contained CE courses
>   - E.g. "Library API and ABI version with the 'inline namespaces'"

### Case study: Extended _friend_ Declrations

> - Prior to C++11, _friend_ declareations require an 'elaborated type specifier'.
>   - _elaborated type specifier_: Syntitical element having the form of \<class|struct|union> \<identifier>
> - This restriction prevents other entities to be designated as friends.
>   - E.g. type aliases, template parameters.
> - A surprsing behavior with namespaces.
>   - it wasn't possible to refer to a entity in the global namespace, a new entity was being declared instead.

```cpp
//C++03 friend
struct S;
struct Example
{
    friend class S; //ok
    friend class NonExistent; //ok, even it this class doesn't exist.
};

using WindowManger = UnixWindowManager;

template <typename T>
struct Example2
{
    friend class WindowManger; //error! type alias
    friend class T; //error! template parameter
};

struct SA; //this SA is in the global namespace
namespace ns
{
    class X3
    {
        friend struct SA; // ok, declares a new ns::SA class instead of refereing to the global ::SA
    };
}
```

> C++ 11 extended 'friend' declarations lift all the aforementioned limitations. and fixes the weird behavior of creating types. we don't need the class|struct|union specifier anymore.

```cpp
//C++11 friend

Struct S;
typedef S Salias;
using Salias2 = S;

namespace ns
{
    template <typename T>
    struct X4
    {
        friend T; //ok
        friend S; //ok, refers to global ::S
        friend SAlias; //ok, also refers to global ::S
        friend Salias2; //ok, also refers to global ::S
        friend decltype(0); //ok, same as 'friend int'
        friend C; //error! 'C' does not name a type.
    }
}
```

> so why is this feature categorized as _unsafe_?
>
> - It is rarely useful in practice, like c++03 _friend_
> - Promotes _long-distance friendship_!
>
> When a type 'X' befriends type 'Y' which lives in a separate component...
>
> - 'X' and 'Y' cannot be thoroughly tested independently anymore.
> - Physical coupling occurs between 'X' and 'Y' components,
> - Possible physical design cycles can happen.

if we have too many friends, it might be a symptom of a design problem, having friends from diffrent namespaces means more coupleing and less modularity. but even if it's _unsafe_, it does have it's benefits, like helping us spot typos when declaring friends.

other intersting points: type alias customization points, PassKey idiom...

and we will focus on CRTP - curiously recursive template pattern.

base knows who it derives from, thanks to T. usefull to implement _mixins_ and factor out copy-pasted code.

```cpp
template <typename T>
class Base
{

};

class Derived : public Base<Derived>
{

};
```

example use case, having a counter for classes creations.

```cpp
//header file
class A {
    static int s_count; //decleration
    public:
    static int count() {return s_count;}
    A(){++_count;}
    A(const A&) {++s_count;}
    A(A&&) {++s_count;}
    ~A() {--s_count;}
};

//defintion file
int A::s_count;
```

we can factor out the counter behavior, using the protected access modifier. (it's a mixin, whatever that means).

```cpp
template <typename T>
class InstanceCounter
{
    protected:
    static int s_count; //declaration
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)
```

we can then use in other classes

```cpp
struct A :InstanceCounter<A>
{
    A() {++s_count};
    //also add this for the destructor
};

struct B : InstanceCounter<A> //oops, made a typo! we will use the same counter.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //oops, we messed with the entire tree!
}
```

actually, this is something we could use the friend declerations! we move from 'protected' to 'private', and make T a friend class of the counter. now only the class that declared the counter can access it, not others classes and not derived.

```cpp
template <typename T>
class InstanceCounter
{
    private:
    static int s_count; //declaration, private
    friend T; //only T can access us.
    public:
    static int count(){return s_count;}
}

template <typename T>
int InstanceCounter<T>::s_count;  //definition (in the same file)

struct B : InstanceCounter<A> //error, s_count is private within this context.
{
    B() {++s_count};
};

struct AA : A
{
    AA() {s_count =-1;} //error, s_count is private within this context
}

```

the crtp allows us avoid boiler plate code, this is also an example of using inheritance without virtual functions. this is also a case where we don't want to use 'final'.

### when to use 'final'

do we really know a class shouldn't be inherited from? are we sure.

it's okay if we have a class that is supposed to behave like a primitive.
but EBO (empty base optimization) doesn't play nice with 'final'.

### Conclusion

> - The "Human cost" of a feature is not easy to quantify.
> - Categorizing features by "safety" helps with devising learning paths.
> - All features have good use cases and nasty pitfalls.

the book will be out in the future, check [this page](https://vittorioromeo.info/emcpps.html)

</details>
