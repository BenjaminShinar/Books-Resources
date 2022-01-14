<!--
ignore these words in spell check for this file
// cSpell:ignore rtime conceptify Parnas
-->

[Main](README.md)

## Algorithms from a Compiler Developer's Toolbox - Gábor Horváth

<details>
<summary>
A bit of compiler algorithms for optimization, using identities and algebra.
</summary>

[Algorithms from a Compiler Developer's Toolbox](https://youtu.be/eeS1WP7FK-A),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/CompilerAlgorithmsTalk.pdf)

### Why Study Compiler?

there are many algorithms and data structure that are used in compilers, compilers are everywhere, like web browsers (html+css, svg, JavaScript of course), GPUs also have compilers, databases have compilers and optimizers, even some configuration format files have something like a compiler inside them. python compiles down to something for machine learning, and routers and modem have something running on them.

A lot of opportunities to improve code, any small improvement is multiplied because it effects every program compiled by it. if we improve a low level compiler (like c++), then we also effect any compiler that uses it (like python or JavaScript).

example: loop strength reduction.

video of a talk by matt godbolt. replacing sum by loop with sum by formula

$
\sum x \equiv \\\frac{x (x+1)}{2}
$

playing with loops kinds and looking at the assembly, we see that the compiler manages to remove the loop and figures out a closed-form formula.

but floating point messes up the optimizations, floating point arithemetic.

### What's Inside the Compiler?

Math.😅\
Chains of recurrences - recursive function when the increment is also a recursive function.

two kinds of recursive formula notations. making functions at incrementoars

$
f(i*i(n)) = {initial,+,incrememt }
$

algebra, operations identities. making loops into recursive notations with those identities, sub expressions combine together. turning this

```cpp
for(i=0;i<m;++i)
v[i] = (i+1)*(i*1) -i*i -2*i;
```

into a constant expression with identities.

$(x+y)(x+y) = (x+y)^2 = x^2 +2xy + y^2$

so we open up the identity, and we can then cancel out stuff and reach a constant.

```cpp
for(i=0;i<m;++i)
v[i]=1;
```

arithemetic series which are supposed to be loops can be made into closed formulas or at least have much less operations per loop

```cpp
int t[20];
for (int i=0;i< 20; i+=1)
{
    t[i]=(i+1)*(i+1) + 3*i - 5; // four additions, two multiplications
}
```

is transformed into this compact form with only two additions.

$
f_{(i+1)^2+3i-5}(n) = \\
\{-4,+6,+2\}
$

which is equivalent to writing this c++ code.

```cpp
int t[20];
int a = -4;
int b = 6;
for (int i=0;i< 20; i+=1)
{
    t[i]=a;
    a+=b;
    b+=2;
}
```

an example of how clang does it. we take this code into a file.

```cpp
int f(int num)
{
    int result =0;
    for (int i =0;i<num;++i)
    {
        result += i;
    }
    return result;
}
```

and then run the following command on it (replace $1 with file name)

```bash
clang++ $1.cpp -c -02 -Xclang -disable-llvm-passes -emit-llvm -S
opt $1.ll -mem2reg -S > {$1}2.ll
opt ${1}2.ll --analyze --scalar-evolution
(other)
```

- -Xclang \<arg> Pass \<arg> to the clang compiler.
- -disable-llvm-passes
- -emit-llvm Use the LLVM representation for assembler and object files
- -S preprocessor only

we can see in the slides how loops are eliminated.

> Recapping Chains of recurrences
>
> - Great to model some loop varian values.
> - Algebra of simple recursive function
> - Algebraic simplifications
> - Strength reduction
> - Closed forms
> - and many more...

### Value Numbering

eliminating some forms of redundancy.

this code has redundancy.

```cpp
int calculate(int a, int b)
{
    int result = (a * b) +2;
    if (a %2 ==0)
    {
        result +=a*b;
    }
    return result;
}
```

the compiler can do the common expression optimization in some cases. but most of the redundancy isn't from the programmer. this code had redundancy in terms of memory access;

```cpp
int matrix[5][5];
//...
matrix[1][2]=bar();
matrix[1][3]=baz();
```

is actually memory dereferencing with a common sub expression.

```cpp
int matrix[5][5];
//...
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 2)=bar();
*((int*)matrix + ROW * sizeof(int) *1 + sizeof(int) * 3)=baz();
```

we can also have dead_code and unused code that passes around (constant propagation).
compilers work in phases, and at each pass the complier cleans up the code to make it optimize. each pass does a small change.

[BRIL - big red intermediate language](https://github.com/sampsyo/bril) is a compiler IR (Intermediate representation) that is used in some courses to teach about compilers.

optimizations can work across different scopes (function, loop body, and even higher!);

local value numbering optimization. algebraic identities, dead code elimination, constant folding,

### where to learn more

some sources to learn mode about compilers.
audience questions

</details>

## So You Think You Know How to Work With Concepts? - Andrzej Krzemieński

<details>
<summary>
a different perspective on concepts and some issues
</summary>

[So You Think You Know How to Work With Concepts?](https://youtu.be/IUPaAcIk1Us), [Slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/andrzej_concepts.pdf)

a different perspective on concepts.

defintions:\
two meanings of "concept" in c++:

- Interfaces in _generic programming_.
- A c++20 language feature.

exploring the limitations of concepts.
he says that the term generic library should be reserved to those who use concepts.

not all 'templated libraries' use concepts, templates are generic, but aren't always based on concepts.

```cpp
boost::optional<std::mutex> om; //Only requires Destructible
auto om2=om; //won't compile, conditional interface
```

as opposed to a different library -[markable](https://github.com/akrzemi1/markable), which uses concepts

```cpp
template <typename MP>
concept mark_policy = requires
{
    // set marked value
    // check for marked value
};


markable<mark_int<int,-1>> oi;
```

a concept is a template (like class template, function template, variable template and alias templates), it's parameterized, it's can be composed of other concepts. the && in the concept world is a conjunction, not logical and. we have some predicates.

the requires statement has new syntax
live coding example

```cpp
#include <iostream>
#include <cassert>
#include <concepts>

namespace lib
{
    template <typename T>
    concept Addable = std::regular<T> //concept
    && requires (T & o,T const & a,T const & b) //conjunction
    {
        {a+b} -> std::convertible_to<T>; //concept
        {o+=b} -> std::same_as<T&>; //concpet
    };

    template <typename T>
    T sum(T const & a, T const & b)
    {
        return a +b;
    }
}

int main()
{
    auto r = lib::sum(1,2);
    std::cout << r <<'\n';
}
```

here is one potential pitfall, instead of requiring the T to be regular, we simply require it to be a valid argument for the std::regular concept, we check if this check is allowed, not the result of it.

```cpp
    template <typename T>
    concept Addable =
    //std::regular<T> &&
    requires (T & o,T const & a,T const & b) //conjunction
    {
        {a+b} -> std::convertible_to<T>; //concept
        {o+=b} -> std::same_as<T&>; //concpet
        std::regular<T>; //wrong!
        //requires std::regular<T>; // this will work
    };
```

the next step is to constrain the library on the concept, this will help the user understand the library when something goes wrong. concept is a predicate, so we can check if a type satisfies it directly.

```cpp
int main()
{
    static_assert(lib::Addable<int>); //check if int fits the concept of Addable
}
```

lets see the difference in the error messages, when we use the templated function, we get a fairly readable compiling error

```cpp
struct X{};
int main()
{
    auto r = lib::sum(X{},X{});
}
```

lets constrain the template itself to use the concept

```cpp
//form 1
template <typename T>
    requires Addable<T>
T sum(T const & a, T const & b)
{
    return a +b;
}

//form 2
template <Addable T>
T sum(T const & a, T const & b)
{
    return a +b;
}
```

now when we try the same code we see a huge error message, just because we used the concept. it tells us the problem is with the interface.

```cpp
struct X{};
int main()
{
    auto r = lib::sum(X{},X{});
}
```

but what if we had some class that has better 'sum' performance if we use the '+=' operator, when the bigger value is on the left side (maybe the function needs to duplicate nodes or something). we stick the implementation details in a nested namespace.

```cpp
namespace lib
{
    template <typename T>
    concept Addable = std::regular<T> //concept
    && requires (T & o,T const & a,T const & b) //conjunction
    {
        {a+b} -> std::convertible_to<T>; //concept
        {o+=b} -> std::same_as<T&>; //concpet
    };
    namespace details
    {
        template <typename T>
        T sum_(T a, T b)
        {
            assert(A > = b);
            return a +=b;
        }
    }
    template <Addable T>
    T sum(T const & a, T const & b)
    {
        if (a < b)
        return details::sum_(b,a);
        else
        return details::sum_(a,b);
    }

}
```

we can use concepts together with the 'auto' keyword

```cpp
int main()
{
    lib::Addable auto r= lib::sum(1,2);
}
```

lets test some more types and see if they work.

```cpp
#include <boost/rational.hpp>
#include <boost/multiprecision/cpp_int.hpp>

int main()
{
    {
    using Rational = boost::rational<int>;
    static_assert(lib::Addable<Rational>);
    Rational a(1,2), b(1,3);
    lib::Addable auto rt = lib::sum(a,b);
    }

    {
    using BigInt = boost::multiprecision::cpp_int;
    static_assert(lib::Addable<BigInt>);
    BigInt a("1122"), b("2233");
    lib::Addable auto rt = lib::sum(a,b);
    }
}
```

and now the user decides to use type complex. which fails. we ask to use the performance version, and we fail, because even though the type is conforming to our concept, it doesn't have the correct operators.

```cpp
#include <complex>
int main()
{
    using ComplexD = std::complex<Double>;
    static_assert(lib::Addable<ComplexD>);
    ComplexD a(0,1), b(1,0);
    lib::Addable auto rt = lib::sumPerformace(a,b);
}
```

> - No guarantee the the function uses only the concept interface.
> - `static_assert(LibConcept<UserType>);`
>   - on failure, a guarantee that the type will _not_ work.
>   - on pass, _no_ guarantee that the type will work.

this negative guarantee protects us from user types accidentally satisfying the function constrains but not the concept, which gives the library more implementation flexability.

```cpp
template <typename T>
bool differ(T const & a, T const & b)
{
    //return !(a==b); maybe this was too expensiva
    return (a!=b);
}
```

our problem is that we tested it on entire, rich, classes. we want to have narrow defintions to work with. concept archetype, we want the minimal interface. we call this "Concept Archetype".

this example is still too big, all we did was make the class conform to the addable concept, but it's still to big (and too specific in the return type of the plus operators)

```cpp
namespace lib
{
    ///template <typename T> concept Addable

    namespace details
    {
        class A{
            //A(A&&)=delete; //declare this as deleted, will also remove all other constructor (copy and move) and destructor. but we actually want them for the std::regular
            void operator&()= delete; // delete the 'address of' operator;
            friend void operator,(A,A) = delete; //delete the comma operator;
            public:
            A& operator+=(A const &); //satisfying the Addable concept
            friend A operator+(A const &,A const &); //satisfying the Addable concept
            friend bool operator==(A const &,A const &) = default; //for the std::regular which requires equality operators

        };
        using AddableArchetyp = A;
        static_assert(Addable<AddableArchetyp>);
    }
}
```

lets try making the concept smaller and narrower by using an inner class.

```cpp
namespace lib
{
    ///template <typename T> concept Addable

    namespace details
    {
        class A{
            //A(A&&)=delete; //declare this as deleted, will also remove all other constructor (copy and move) and destructor. but we actually want them for the std::regular
            void operator&()= delete; // delete the 'address of' operator;
            friend void operator,(A,A) = delete; //delete the comma operator;
            struct Result
            {
                operator A(); //casting
                Result(Result&&)=delete; //declare this as deleted, will also remove all other constructor (copy and move) and destructor.
                void operator&()= delete; // delete the 'address of' operator;
                friend void operator,(Result,Result) = delete; //delete the comma operator;
            };
            public:
            A& operator+=(A const &); //satisfying the Addable concept
            friend Result operator+(A const &,A const &); //satisfying the Addable concept
            friend bool operator==(A const &,A const &) = default; //for the std::regular which requires equality operators

        };
        using AddableArchetype = A;
        static_assert(Addable<AddableArchetype>);

        inline void test_sum()
        {
            sumPerformace(AddableArchetype{},AddableArchetype{})
        }
    }
}
```

this will work because this concept uses default constructable behavior in the test. so we need to change it.we simply assume we get those objects from outside, this function is only really used to ensure the behavior is possible, it's never truly called. it's just to make sure it complies. we can also make sure the result type is appropriate

```cpp
namespace lib
{
    ///template <typename T> concept Addable

    namespace details
    {
        class A{
            //..
        };
        using AddableArchetype = A;
        static_assert(Addable<AddableArchetype>);

        inline AddableArchetype test_sum(AddableArchetype const & a,AddableArchetype const & b)
        {
            return sumPerformace(a,b);
        }
    }
}
```

now, if we run the same example with teh complex number type, we detect the bug properly. it says that we can't instantate this function because of the class A.

```cpp
#include <complex>
int main()
{
    using ComplexD = std::complex<Double>;
    static_assert(lib::Addable<ComplexD>);
    ComplexD a(0,1), b(1,0);
    lib::Addable auto rt = lib::sumPerformace(a,b);
}
```

this bug tells us the we use more operations than what we specify in the interface, now we need to decide how we deal with this./
we can simply use 'if-constexpr' to determine this in a compile time.

```cpp
template <Addable T>
T sumDepending (const T & a, const T & b)
{
    if constexpr(std::totally_ordered<T>)
    {
        retrun sumPreformace(a,b);
    }
    else
    {
        return a+b;
    }
}
```

or we can use two overloads. concepts allows us two overloads that differ only by constraints.

```cpp
template <Addable T>
T sumOverload (const T & a, const T & b)
{
    return a+b;
}

template <Addable T>
    requires std::totally_ordered<T>
T sumOverload (const T & a, const T & b)
{
            //implementation
        if (a < b)
        return details::sum_(b,a);
        else
        return details::sum_(a,b);
}
```

we can also introduce this constraint as a concept by itself, but it has some issues that will be detailed later.

```cpp
template <typename T>
concept OrderedAddable = Addable<T> &&  std::totally_ordered<T>;
```

lets add archtypes to make sure the overloads apply and are ok.

```cpp
namespace lib
{
    namespace ordered
    {
        class A
        {
        //A(A&&)=delete; //declare this as deleted, will also remove all other constructor (copy and move) and destructor. but we actually want them for the std::regular
        void operator&()= delete; // delete the 'address of' operator;
        friend void operator,(A,A) = delete; //delete the comma operator;
        struct Result
        {
            operator A(); //casting
            Result(Result&&)=delete; //declare this as deleted, will also remove all other constructor (copy and move) and destructor.
            void operator&()= delete; // delete the 'address of' operator;
            friend void operator,(Result,Result) = delete; //delete the comma operator;
        };
        public:
        A& operator+=(A const &); //satisfying the Addable concept
        friend Result operator+(A const &,A const &); //satisfying the Addable concept
        friend auto operator<=>(A const &,A const &) = default; //spaceshipt!
        };
    }
    using OrderedArchetype = ordered::A;
    static_assert(Addable<OrderedArchetype>);
    namespace details
    {
        inline OrderedArchetype test_sum(OrderedArchetype const & a,OrderedArchetype const & b)
        {
            return sumPerformace(a,b);
        }
    }
}
```

there can be multitype-concepts. onc concept with three constrained types, this will require three archtypes to test the concept.

```cpp
template <class Iter,class Sentinel, class Pred>
concept PredicatedIteration = requires(Iter i,Sentinel s,Pred p)
{
    {i !=s }=> std::convertible_to<bool>;
    {p(*i)}=> std::convertible_to<bool>;
    ++i;
}
```

archtypes are a cartesian product of the number of constraints and concepts.

now lets try the _std::string_ class. we get a different result than expected, we come across another hidden requirement. we can't specify semantics assumptions requires easily.

```cpp
int main()
{
    static_assert(lib::addable<std::string>);
    std::string a("air"),b("bus");
    lib::Addable auto r = lib::sumDependant(a,b);
    std::cout<< r <<'\n'; //oops! we get "busair" instead of "airbus", the string type is ordered!
}
```

we could try all sorts of things,

```cpp
    template <typename T>
    concept Addable = std::regular<T> //concept
    && requires (T & o,T const & a,T const & b) //conjunction
    {
        {a+b} -> std::convertible_to<T>; //concept
        {o+=b} -> std::same_as<T&>; //concpet
        requires requires
        {
            a+b == b+a;
        }
    };
```

why doesn't this happen in OOP? in objects the type says it implements some base class, and then it's declaring itself to comply with the baseclass interface. concepts are more like 'duck typing'. we can violate syntactic requirements and get a compiler error (fine), but we can also violate semantic requirements and get bugs (undefined behavior for std concepts)

we can add more test to ensure compile time correctness, and we can try to emulate OOP by forcing the used type to specify they confrom to us. we add a template parameter (false) and specialize for it (opt in for what we want). or we can do the inverse, and specialize to optout the non confroming types.
the standard library has a similar trick wth _std::view_ and _std::range_

</details>

## The Concepts of Concepts - Sandor Dargo

<details>
<summary>
More about concepts, constraints and combining them.
</summary>

[The Concepts of concepts](https://youtu.be/weJD_ZCr6S8), [slides-pptx](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/The-Concepts-of-Concepts-CNow.pptx)

### Why Do We Need Concepts?

part of the big four changes of c++20(concepts,ranges, coroutines, modules).

concepts are extensions of templates, help us validate template arguments at compile time.

overloads don't scale

```cpp
long double add(long double a, long double b)
{
    return a+b;
}
int add(int a, int b)
{
    return a+b;
}
```

templates are better, but there are no constraints on them, and they might cause unexpected behavior.

```cpp
template <typename T>
T add(T a,T,b)
{
    return a+b;
}

int main()
{
    add(42,66);
    add(42.42L,66.6L);
    add('a','b'); //is this ok for us?, characters
}
```

we can use templates and forbid some specialization, which works, but we are back to the problem that it doesn't scale well.

```cpp
template <>
std::string add(std::string, std::string) = delete;

int main()
{
    add(std::string{"a"},std::string{"b"}); //deleted function
}
```

type traits and asserts are one step up, with better (sometimes) error messages, all the code is in one place. the problem is that the assert statement becomes a code smell once we try to use it in many places.

```cpp
template<typename T>
T add(T a, T b)
{
    static_assert(std::is_integral_v<T> || std::is_floating_point_v<T>,"add can be called only with numbers")
    return a+b;
}

int main()
{
    add(std::string{"a"},std::string{"b"}); //assertion fails
}
```

Concepts are a way to scale that idea up, we can express the earlier assertions as a 'type'.

```cpp
template <typename T>
concept Number = std::integral<T> || std::floating_point<T>;

template <Number T>
auto add (T a,T b)
{
    return a+b;
}

int main()
{
    add(1,2);
}
```

this makes templates more safe to use, generic programming becomes scalable, and we can put more of our domain knowledge into code

### 4 Ways to Use Concepts With Functions

for now, we will use the 'Number' concept from before.

```cpp
#include <concepts>
template <typename T>
concept Number = std::integral<T> || std::floating_point<T>;
```

#### Using the _requires_ clause

> - _requires_ following the template parameter list.
> - After _requires_ write your concept(s) to be satisfied.

```cpp
template <typename T>
requires Number<T>
auto add(T a, T b)
{
    return a+b;
}
```

we can also write the constraints directly, without creating the concept.

```cpp
template <typename T>
requires std::integral<T> || std::floating_point<T>
auto add(T a, T b)
{
    return a+b;
}
```

we constrain multiple template parameters type

```cpp
template <typename T,typename U>
requires Number<T> && Number<U>
auto add(T a, U b)
{
    return a+b;
}
```

function calls are written as usual, but the error messages are better and mention which concept fail

#### Trailing _requires_ clause

the _requires_ comes after the parameters and the CV qualifiers. other than that, same as before, supports combinations and multiple parameter types.

```cpp
template <typename T>
auto add(T a, T b)
requires Number<T>
{
    return a+b;
}
```

#### Constrained template parameter

this is currently the suggested form by the guidelines.

not writing _requires_ any more, the _typename_ is replaced by the concept. supports multiple parameters, but doesn't support combinations of concepts.

```cpp
template <Number T>
auto add(T a, T b)
{
    return a+b;
}
```

#### Abbreviated function templates

no _requires_, no template parameter list. use _conceptName auto_ in the parameter list. parameter types can be implicitly different. no support for combinations.

```cpp
auto add(Number auto a, Number auto b)
{
    return a+b;
}
```

we have to write both the concept and auto to make sure this is a concept and not a type which happens to have the same name...

#### How to choose?

| style                          | example                                                        | combinations | multiple types |
| ------------------------------ | -------------------------------------------------------------- | ------------ | -------------- |
| requires clause                | `template <typename T> requires Number<T> auto add(T a, T b) ` | possible     | possible       |
| trailing clause                | `template <typename T> auto add(T a, T b) requires Number<T>`  | possible     | possible       |
| constrained template parameter | `template <Number T> auto add(T a, T b)`                       | impossible   | possible       |
| abbreviated function temples   | `auto add(Number auto a, Number auto b)`                       | impossible   | implicit       |

If we have complex requirement that isn't expressed in a concept, we need to use _requires_.\
For a simple constraints we should use the abbreviated function template.\
If it's a simple constraint and we want to control the types, we can use the constrained template parameter style.

they are all the same thing underneath (if we have multiple parameters)

### Concepts with Classes

For classes there are fewer styles available. Abbreviated function templates won't make sense, and trailing _requires_ clause only fit certain circumstances.

#### The _requires clause_

same as with functions, we can use concepts and complex expression combinations.

```cpp
template <typename T>
requires Number<T>
class WrappedNumber {
public:
    WrappedNumber(T num) : m_num(num){}
private:
    T m_num;
};
```

#### Constrained template parameters

replace the `typename` with the concept name. no extra constraints and complex expressions.

```cpp
template <Number T>
class WrappedNumber {
public:
    WrappedNumber(T num) : m_num(num){}
private:
    T m_num;
};
```

#### Trailing _requires_ clause

class level templates with concepts on the functions. provide different 'overloads' for different parameter types. this is what we would do with _std::enable_if_.

```cpp
template <typename T>
class MyNumber {
public:
    T divide(const T& divisor)
    requires std::integral<T>
    {
        //...
    }
    T divide(const T& divisor)
    requires std::floating_point<T>
    {
        //...
    }
};
```

### What is in the STL?

there are about 50 or so concepts in the STL in 3 headers.

#### \<concepts>

> - Core language concepts (_same_as_, _integral_, _constructible_from_,...)
> - Comparison concepts (_totally_ordered_,...)
> - Object concepts (_copyable_, _regular_,...)
> - Callable concepts (_invokable_, _predicate_,..)

concepts are also combined together to create more complex concepts.

_std::constructible_from_ uses _destructible_ concept and _std::is_constructable_v_ and from the type traits.

> ```cpp
> template < class T, class... Args >
> concept constructible_from =
>   std::destructible<T> && std::is_constructible_v<T, Args...>
> ```

_std::default_initializable_ uses the _constructible_from_ concept and combines it with expressions that dictate it has a parameterless constructor and can be constructed on the heap with the default allocator.

> ```cpp
> template<class T>
> concept default_initializable =
>   std::constructible_from<T> &&
>    requires { T{}; } &&
>    requires { ::new (static_cast<void*>(nullptr)) T; };
> ```

#### \<iterator>

> - Iterator concepts (_incrementable_, _input_iterator_,...)
> - Indirect callable concepts (_indirectly_unary_invocable_,...)
> - Common algorithm requirements (_mergeable_, _sortable_,...)

the _std::output_iterator_ concept build on the input/output iterator concept and requires it to be writeable and others.

_std::indirect_unary_predicate\<F,I>_ combines iterator concepts and predicate concepts.

#### \<ranges>

concepts from ranges (not in the lecture)

> - ranges::range
> - ranges::borrowed_range
> - ranges::sized_range
> - ranges::view
> - ranges::input_range
> - ranges::output_range
> - ranges::forward_range
> - ranges::bidirectional_range
> - ranges::random_access_range
> - ranges::contiguous_range
> - ranges::common_range
> - ranges::viewable_range

### How to Write Concepts?

to write a concepts we first list all the template parameters, then the word concept and the concept name, and finally all the requirements.

the simplest concept looks like this, with the name 'Any'.

```cpp
template <typename T>
concept Any = true;
```

we already had the Number concept, we used predefined concepts and combined them together.

#### What does combining concept mean?

we can use conjunctions (_and &&_) and disjunctions (_or ||_).\

> - concepts
> - bool literals
> - bool expressions
> - type traits (_::value_, _\_v_)
> - _requires_ expressions

**We should be careful with the negation operator (_not ,!_)**

the negation means diffrent things: \
for **boolean expressions**, a negation means that the all subexpressions are well-formed, compile, but return _false_.\

for **concepts**, a subexpression can be ill-formed, might return false, and the rest can be still satisfied

example:

> - It doesn’t have to be compilable.
> - It can return false.
> - Expecting false is possible With a cast to bool or with a more explicit way.

```cpp
template <typename T, typename U>
requires
    std::unsigned_integral<typename T::Blah> ||
    std::unsigned_integral<typename U::Blah>
void foo(T bar, U baz) { /*...*/ }
class MyType
{
public:
    using Blah = unsigned int;
// ...
};
```

if just one type (T or U) has 'blah' that is an unsigned integer, it's fine. even if the other doesn't even have 'blah', this should not compile, but the concept is ok.

if we want both of them to have the nested type, and one of those types should be unsigned integer, we can write it differently

if we cast it to a boolen expression, it must be a well formed expression, so it won't compile if one of them doesn't have the nested type.

```cpp
template <typename T, typename U>
requires (bool(
    std::unsigned_integral<typename T::Blah> ||
    std::unsigned_integral<typename U::Blah>))
void foo(T bar, U baz)
{
/*...*/
}
```

the other option is a nested require expression. we first require both of them to exists and then require that one of them to be an unsigned integer. this is more verbose, and doesn't require understanding the small prints of boolean expressions as opposed to concepts, but it also seems messy

```cpp
template <typename T, typename U>
requires (
    //one constraint
    requires {typename T::Blah;} &&
    requires {typename U::Blah;})
    &&
    (
    // second constraint
    std::unsigned_integral< typename T::Blah> ||
    std::unsigned_integral<typename U::Blah>)
void foo(T bar, U baz)
{
/*...*/
}
```

#### How to find the most constrained constraint

the most constrained one will be chosen, based on the call

```cpp
template <typename Key>
class Ignition
{
public:
    void Start(Key key)
    requires(!Smart<Key>)
    {
        //... no concept
    }
    void Start(Key key)
    requires Smart<Key>
    {
        //... concept
    }
};
```

this is a bad design, we should have one overload with constrains and the other without.

```cpp
template <typename Key>
class Ignition
{
public:
    void Start(Key key)
    {
        //... no concept,
    }
    void Start(Key key)
    requires Smart<Key>
    {
        //... concept, the most constrained.
    }
};
```

if we had more overloads, the most constrained one will be chosen in compile time. this is called 'concepts subsumption'.

```cpp
template <typename Key>
class Ignition
{
public:
    void Start(Key key)
    {
        //... no concept,
    }
    void Start(Key key)
    requires Smart<Key>
    {
        //... concept
    }
    void Start(Key key)
    requires Smart<Key> && Personal<Key>
    {
        //... concept, most constrained
    }
};
```

in this case, negation brings ambiguity into the picture. the negation must use parentheses, and relies on source location of each constraints. the negation creates a new 'temporary concept', so if the overloads use the same syntax of negation, it's still parsed as two different concepts, and we can't choose either.

```cpp
template <typename Key>
class Ignition
{
public:
    void Start(Key key)
    {}
    void Start(Key key)
    requires (!Smart<Key>)
    {}
    void Start(Key key)
    requires (!Smart<Key>) && Personal<Key>
    {}
};
```

instead, we should have a named negative concept, even though the core guidelines caution against unnecessary named concepts.

```cpp
template <typename Key>
class Ignition
{
public:
    void Start(Key key)
    {}
    void Start(Key key)
    requires NotSmart<Key>
    {}
    void Start(Key key)
    requires NotSmart<Key> && Personal<Key>
    {}
};
```

#### So how to write?

when we write concepts we should list all the variables used in the requirement and their operations and function calls

```cpp
#include <concepts>

template <typename T>
concept Addable =
requires (T a, T b)
{
    a+b;
};

template <typename T>
concept HasSquare =
requires (T t)
{
    t.square(); //type T must have a function called square
};

template <typename T>
concept HasPower =
requires (T t, int exponenet)
{
    t.power(exponenet); //type T must have a function called power that can take an integer
};
```

if we want to define the return type, we can use _compound requirements_ with _std::convertable_to\<T>_ and _std::same_as\<T>_.\
We should pay attention to the **curly braces**.\
no bare types for future generalizations.

```cpp
template <typename T>
concept HasSquare =
requires (T t)
{
    {t.square()} -> std::convertable_to<int>;
};
```

we can also have 'Type requirements', such as:

> - A certain nested type exists.

```cpp
template<typename T>
concept TypeRequirement =
requires
{
    typename T::value_type;
};

int main()
{
    TypeRequirement auto myVec =  std::vector<int>{1, 2, 3};  // has value_type
    TypeRequirement auto myInt {3}; //error, deduced type int doesn't satisfy constraint
}
```

> - A class tempate specialization names a type.

the type can be used as type template parameter for a different type

```cpp
template <typename T>
requires (!std::same_as<T, std::vector<int>>)
struct Other {};

template<typename T>
concept TypeRequirement = requires
{
    typename Other<T>;
};

int main()
{
    TypeRequirement auto myVec =  std::vector<char>{'a', 'b', 'c'}; //works, Other can take vector<char>
    TypeRequirement auto myVec2 = std::vector<int>{1, 2, 3}; // error, Other can't be used with vector<int>
}

```

> - An alias template specialization names a type.

```cpp
template<typename T>
using Reference = T&;

template<typename T>
concept TypeRequirement = requires
{
  typename Reference<T>;
};
```

concepts can be nested, we can have new constains without new named constraints

```cpp
template<typename C>
concept Car = requires (C car)
{
    car.startEngine();
};

template<typename C>
concept Convertible = Car<C> && requires (C car)
{
    car.openRoof();
};

template<typename C>
concept Coupe = Car<C> && requires (C car)
{
    requires !Convertible<C>; //must be a car, but not must be convertable,
};
```

the better way is to do this, the parentheses and negation rules apply for _requires_ clauses.

```cpp
template <typename C>
concept Coupe = Car<C> && !Convertible<C>;
```

clones example:

```cpp

struct Droid {
    Droid clone(){return Droid{};}
};

struct DroidV2 {
    Droid clone(){return Droid{};} //oops! we made a copy and paste error!
};
template <typename C>
concept Clonable = requires (C cloneable)
{
    cloneable.clone(); // has function clone
    requires std::same_as<C,std::decaltype(<cloneable.clone())>>; //return type of clone must be the same as the calling object
}

int main()
{
    Clonable auto c1 =Droid{}; //fine
    Clonable auto c2 =DroidV2{}; //error! DroidV2.clone() doesn't satisfy the cloneable requirement!
}
```

and the easier syntax is

```cpp
template <typename C>
concept Clonable = requires (C cloneable)
{
    {cloneable.clone()} ->std::same_as<C>; // has function clone that returns the same type
}
```

we can use nested requirement to simulate boolean expressions, like we saw above. we should consider readability.

### Real-life Examples

some integral types are not numbers, such as bool and char, so we add this to our concept.

```cpp
template <typename T>
concept Number = (std::integral<T> || std::floating_point<T>)
&& !std::same_as<T,bool>
&& !std::same_as<T,char>
&& !std::same_as<T,unsigned char>
&& !std::same_as<T,char8_t>
&& !std::same_as<T,char16_t>
&& !std::same_as<T,char32_t>
&& !std::same_as<T,wchar_t>;

auto add(Number a,Number b)
{
    return a+b
}
```

we can turn utility template functions (sometimes without documentation, and with bad naming conventions) into self documenting code. we go into the template implementations and extract a concept outside and use it as a type. now the functions is constrained, explains the constraint and all the information is combined. we put the domain knowledge into the code.\
We make bad code better, its now more maintainable and easier to understand. even if it's not bad code, it's better code.

if we don't want a named concept, we can have '_requires requires_' to get us able to use parameters in an unnamed context.

```cpp
template <typename BOWithEncodeableStuff_t>
requires requires (BOWithEncodeableStuff_t bo) //get the ability to use the parameter Type
{
  bo.interfaceA();
  bo.interfaceB();
  { bo.interfaceC() } -> std::same_as<int>;
}

void encodeSomeStuff(BOWithEncodeableStuff_t iBusinessObject)
{ /*...*/ }
```

### Conclusion

> TakeAways:
>
> - Concepts help validate template arguments at compile-time.
> - Concepts provide a reusable and scalable way to constrain templates.
> - The standard library gives dozens of generic concepts.
> - There are plenty of ways to define our concepts.
>
> How to Start:
>
> - Start using concepts as soon as you switch to C++20.
> - Use them for your applications.
> - No more naked Ts and typenames.

</details>

## Using Concepts: C++ Design in a Concept World - Jeff Garland

<details>
<summary>
A two parts session about concepts. we now have tools and first class language support for concepts.
</summary>

[Using Concepts: C++ Design in a Concept World - part1](https://youtu.be/Ffu9C1BZ4-c),[part2](https://youtu.be/IXbf5lxGtr0),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/2021cppnow_learning_concepts.pdf)

goals:

> climb up the concept ladder - What's a concept - Using concepts in code - Reading concepts: `requires` expressions and clauses - Writing concepts: its hard - Designing with concepts

we want to get the community to use concepts, even people who don't write concepts will still see them in the documentation.

### Concept Basics

> - overview of concepts
> - concepts vs types

The reasons for concepts: we want good generic libraries, which are fast, and that fellow programmers can read, understand and maintain, which means they have reasonable error messages.

when we have unconstrained templates, we might see hundreds of lines for error messages. direct language support can help with tooling, debugging, and even better compiling time.
it will also make template code more readable, more descriptive, and give us better interface to use.

we can see the history of concepts starting in the 80's, but the modern idea began in the STL in the 90's, but without language specifications. in those days the world was dominated by object oriented programming.
Boost had a library called "concept check", which was mostly macro based, in 2011 concepts were supposed to become part of the language, but the committee decided it wasn't ready.
eventually, in c++20 we got a language definition of concepts, and even some libraries (ranges library) are specified in terms of concepts. but there are still mistakes that were made.

in the last minute, it was decided that concepts use snake*case (not PascalCase), and the \_concept bool* was replaced with _concept_.

> Concepts
>
> - boolean predicate on types and values
> - type requirement examples
>   - required methods
>   - required semantics\*
>   - required subtypes or base types
> - c++ realization includes
>   - new keywords `concept` and `requires`
>   - `\<concept_name> auto` for describing a set of types
>   - new rules for function overloading

Concepts are **predicates** that can check against types and against values. we mostly focus on types, but the size can also help with stuff. like running SIMD on small objects if possible.
we can use concepts to require a type to have specif methods, we can also require certain semantics (but this might not be feasible in terms of what the compiler can do) or require them to have some subtypes or base types (like inheritance).

> boolean predicate composition
>
> - support complex compile time logic composition
> - conjunction and disjunction (and/or) logic
> - used to classify types
>   - in or out of a set
>   - that share syntax/semantic
>   - (although semantics is the desire only syntax is checked)

we can have complex logical composition, based on compile time logic,

> types versus concepts
>
> - type
>   - describes a set of operation that can perform
>   - relationships with other types
>     - example: base class
>     - example: declared dependent type
>   - describes a memory layout
>   - for built in types this can be implicit
> - concept
>   - describes how a type can be used
>   - operations it can perform
>   - relationships with other types

a type (primitive, built-in or something that we built),has operations, relationship with other types, dependencies. types have memory layout, with members and inner variables.
the concepts has similar functionalities, but it doesn't describe a memory layout. concepts are more abstracted than types.

in this code, we have a concept that allows "printing" into an ostream. we declare the concept and the type which we want to check, and then we check if the type satisfies the concept.
concepts should go into a namespace, we don't want naming conflicts.
because templates aren't evaluated until they are called, we need to check manually, so a static asset is a good idea.

```cpp
namespace io
{
	// Type T has print ( std::ostream& )const member function
	template<class T>
	concept printable = requires //...more later
}
class my_type
{
	std::string s = "foo\n";
	public:
	void print( std::ostream& os ) const
	{
		os << "s: " << s;
	}
};
static_assert(io::printable<my_type> ); //good
```

concepts are entirely compile time, no runtime footprint. the dependency on a concept means that if we change it, then we probably need to recompile everything.
the core use of concept is to constrain the usage of templates.

### Using Concepts in Code

> - overloading, variables, pointers
> - std library concepts

what we use concepts for:

> concepts can:
>
> - Constrain an overload set
> - Initialize a variable with _\<concept_name> auto_
> - Conditional compilation with constexpr if
> - Can use a pointer or _unique_ptr_ of concept
> - Partially specialize a template with concept
> - Make template code into 'regular code'
>
> concepts cannot:
>
> - Cannot inherit from concept
> - Cannot constrain a concrete type using requires
> - Cannot 'allocate' via new
> - Cannot apply requires to virtual function

overload set - using a different function. we can use concepts in constant expressions.
we can't inherit directly from a concept, we can inherit from a class that is constrained by one.
we can't constrain a concrete type directly, but we can constrain it's usage.
concepts have nothing to do with storage, so no allocations, and they are entirely incompatible with virtual functions.

_/<consept_name> auto_ can be used as a type. either in templates or as placeholder for a type.

```cpp
// Type T has print ( std::ostream& )const member function
template<class T>
concept printable = requires //...

template<typename T>
concept associative_container = //

// take a parameter of type printable auto, a template that looks like regular code.
void f(printable auto s)
{
//...
}

// constraining a template, classic example of concepts.
template<associative_container T>
class MyType
{
	T map_;
}

int main()
{
	//these won't compile, we need a concrete type here, can't allocate storage for it
	//printable auto s;
	//associative_container auto myMap;

	//these will compile
	printable auto s = init_some_thing();
	associative_container auto myMap = MyType<std::map<int,int>>{{1,1}};
}
```

we can also use concepts as function parameters or return value.

[godbolt example](https://godbolt.org/z/r3dY3dsvd), both are the same, one is written in a traditional way, and one hides the template keyword. but both use deduced return type.

```cpp
template<printable T>
printable auto print( const T& s )
{
	//...
	return s;
};

printable auto print2( const printable auto& s )
{
	//...
	return s;
}
```

as always, we can write the same thing in many ways.

overload resolution - constrain function parameter
write different functions _print_line_, and choose the overloaded function based on the concept.
[godbolt example](https://godbolt.org/z/GKq8ns). this will fail, because there isn't a way to print the type. this is an unconstrained template function.

```cpp
#include <iostream>

//auto parameter -- this is a template function!
// template<typename T_
// void print_ln( T p )
void print_ln( auto p )
{
	std::cout << p << "\n";
}

class my_type {};

int main()
{
	print_ln("foo");
	print_ln(100);

	//compile error of course
	//my_type m;
	//print_ln ( m );
}
```

the traditional way to do these was to overload the function for the concrete type, have a better match than the template function.

```cpp
//selected ahead of print_ln (auto) because better match
void print_ln( my_type p )
{
	p.print( std::cout );
	std::cout << "\n";
}
```

but this is tedious, we want to use the printable concepts
[godbolt example 1](https://godbolt.org/z/36cdsGzzo), [godbolt example 1](https://godbolt.org/z/dYdhW7). we use the requires clause, it behaves like a parameter list in this case.
we decide that the expression `v.print(os)` must compile, in the _output_streamable_ case, we say that the expression must be a valid `os <<v` expression.
the public access modifiers are important, we need be able to actually call the function.

```cpp
#include <iostream>
#include <memory>

// Type T has print ( std::ostream& ) member function
template<typename T>
concept printable = requires(std::ostream& os, T v)
{
	v.print( os ) ; //<--an expression that if compiles yields true
};

template<class T>
concept output_streamable = requires (std::ostream& os, T v)
{
	os << v;
};


void print_ln( auto p )
{
	std::cout << p << "\n";
}

//auto parameter -- this is a template function!
void print_ln( printable auto p ) //<-- constrained resolution
{
	p.print(std::cout);
	std::cout << "\n";
}

int main()
{
	print_ln( "foo" );
	print_ln( 100 );
	my_type m;
	print_ln ( m );
}
```

we can use both concepts, have two overloaded functions, one for each concept.

```cpp
// example of overload resolution
void print_ln( output_streamable auto p)
{
	std::cout << p << "\n";
}

void print_ln( printable auto p)
{
	p.print(std::cout);
	std::cout << "\n";
}

class my_type2 {};

int main()
{
	print_ln( "foo" );
	my_type m;
	print_ln ( m );
	//compile error of course
	//my_type2 m2;
	//print_ln ( m2 );
}
```

Pointers and concepts

this is usefull fo things like factory functions

[godbolt](https://godbolt.org/z/d7bGhn) example with unique pointers.

```cpp
int main()
{
    const printable auto *m = new my_type();
    m->print(std::cout);
    const std::unique_ptr<printable auto> upm = std::make_unique<my_type>();
    upm->print(std::cout);
}
```

if we try to assign the address of something that doesn't satisfy the concept to a pointer to the concept, we can get a better error message.

```cpp
class whatever{}; //no print
int main()
{
    printable auto * m = new whatever{}; //error!
}
```

we can use concepts inside `if constexpr`, like in this [godbolt example](https://godbolt.org/z/nsojqK5e1). now we don't have an overload resolution, we use compile time decision making.

```cpp
template <class T>
std::ostream & print_ln (std::ostream & os, const T& v)
{
    if constexpr (requires {printable<T>;})
    // if constexpr (printable<T>) // short form
    {
        v.print(os);
    }
    else
    {
        os <<v;
    }
    os <<'\n';
    return os;
}

int main()
{
    my_type m;
    print_ln(std::cout,m);
    int i =100;
    print_ln(std::cout,i);
}
```

we can also skip the concept entirely, simply use check if this compiles and then use it. this works for simple, trivial stuff, not for complex logic.

```cpp
template <class T>
std::ostream & print_ln (std::ostream & os, const T& v)
{
    if constexpr (requires {v.print(os);})
    {
        v.print(os);
    }
    else
    {
        os <<v;
    }
    os <<'\n';
    return os;
}

int main()
{
    my_type m;
    print_ln(std::cout,m);
    int i =100;
    print_ln(std::cout,i);
}
```

theres a wa to constrain a type based on an internal type, in this example we have a wrapper type that allows dereferencing operator only if the wrapped type is a pointer.

```cpp
template <class T>
class wrapper
{
    T val_;
    public:
    wrapper (T val):val_(val){}
    T operator *() requires is_pointer_v<T>
    {
        return val_;
    }
};

int main()
{
    int i = 1;
    wrapper<int *> wi{&i};
    std::cout << *wi << '\n';
    wrapper<int> wi2{i};
    std::cout << *wi2 << '\n'; //error!
}
```

now we want a vector of objects that belong to the same concept. we can get this by using template alias. we can't instansite the alias with a type that doesn't satisfy the concept.

```cpp
#include <iostream>
#include <vector>
#include <string>

//template alias using concepts
template<printable T>
using vec_of_printable = std::vector<T>;

int main()
{
    vec_of_printable<my_type> vp {{},{},{}};
    for (const auto & e : vp)
    {
        e.print(std::cout);
    }
    vec_of_printable<int> vi; //won't compile
}
```

the relevent headers are \<concepts>,\<type_traits>,\<iterator> and \<ranges>, the concepts are grouped into

- Core language concepts
- Comparison concepts
- Object concepts
- Callable concepts
- Ranges concepts

about type*traits: we already have some stuff that seems like concepts and does compile time stuff, such as \_std::is_arithmetic*. some of them might be replaced by concepts, but more likely we will get much more.

| Concept                        | Description                                       |
| ------------------------------ | ------------------------------------------------- |
| floating_point\<T>             | float, double, long double                        |
| integral\<T>                   | char, int, unsigned int, bool                     |
| signed_integral\<T>            | char, int,                                        |
| unsigned_integral\<T>          | unsigned char, unsigned int                       |
| equality_comparable\<T>        |
| equality_comparable_with\<T,U> | `operator==` is an equivalence, between two types |
| totally_ordered\<T>            | `==`,`!=`,`<`,`>`,`<=`,`>=` are total ordering    |
| totally_ordered_with\<T,U>     | ordering between two types                        |
| same_as\<T,U>                  | types are same                                    |
| derived_from\<T,U>             | T is subclass of U                                |
| convertable_to\<T,U>           | T converts to U                                   |
| assignable_from\<T,U>          | T can assign from U                               |
| default_initializable\<T>      | T has a default ctor                              |
| constructable_from\<T,...>     | T can be constructed from variable pack           |
| move_constructable\<T>         | T has move ctor                                   |
| copy_constructable\<T>         | T has copy ctor                                   |
| semiregular\<T>                | T has deafult, copy and move ctor, and stor       |
| regular\<T>                    | T is semiregular and equality comparable          |

in this example we try to check if our type is regular, it fails because we don't have and equality operations defined

```cpp
class my_type
{
	std::string s = "foo\n";
	public:
	void print( std::ostream& os ) const
	{
		os << "s: " << s;
	}
};
static_assert(std::regular<my_type> ); //fails
```

we simply add the deafult operator to fix this

```cpp
class my_type
{
	std::string s = "foo\n";
	public:
	void print( std::ostream& os ) const
	{
		os << "s: " << s;
	}
    bool operator==(const my_type & ) const = default;
};
static_assert(std::regular<my_type> ); //passes
```

now let's look at the range concepts. we have a concept for something that we can iterate over. a vector, an array, a span and other stuff.

```cpp
void print_integers(const std::ranges::range auto & R)
{
    for (auto i : R)
    {
        std::cout << i << '\n';
    }
}
int main()
{
    std::vector<int> vi = {1,2,3,4,5};
    print_integers(vi);

    std::array<int,5> ai = {1,2,3,4,5};
    print_integers(ai);

    std::span<int> si2 (ai);
    print_integers(si2);

    int cai[] ={1,2,3,4,5};
    std::span<int> si3 (cai);
    print_integers(si3);

    ranges::iota_view iv{1,6};
    print_integers(iv);

}
```

### Reading Concepts

> - requires expressions
> - clause: a boolean expression
>   - used after template and method declarations
>   - clauses can contain an expression
> - expression: syntax for describing type constrains

a requires expression has a sequence of requirements, it can have a parameter sequence, but it it's not necessary.

heres a more realistic version of the printable concept, we declare required member functions and required free functions that must exists.

```cpp
template <typename T>
concept printable = requires(std::ostream & os, T v)
{
    //all of these must be true, they must compile
    v.print(os) ; //member function
    format(v); // free function
    std::movable<T>;
    typename T::format; //declare a type called format
};

template <class T>
concept output_streamable = requires(std::ostream & os, T v)
{
    //this compiles and gives the result we expect, trailing return type syntax.
    {os << v} -> std::same_as<std::ostream&>;
}
```

#### Constraint Composition

- atomic constraints
- conjunction constraints (and)
- disjunction constraints (or)

```cpp
//disjunction with '||' operator
template <typename T>
concept printable_or_streamable = requires printable<T> || output_streamable<T>;

//disjunction with 'or'
template <typename T>
concept printable_or_streamable = requires printable<T> or output_streamable<T>;

//conjunction
template <typename T>
concept fully_outputable = requires printable<T> and output_streamable<T>;
```

we can reactor requirements around to make concepts more readable.

```cpp
template <typename T>
concept printable =
    std::movable<T> and //bring this outside
    requires(std::ostream & os, T v)
    {
        v.print(os) ; //member function
        format(v); // free function
    };
```

another concept example, _std::derived_from_.

```cpp
template <class Derived, class Base>
concept derived_from =
    std::is_base_of_v<Base,Derived> and
    std::is_convertible_v<const volatile Derived *, const volatile Base*>
```

the _std::os_arithmetic_ is a concept that is a type trait that checks if the type has math operations (plus, minus,etc...), but unfortunately, also includes bool and char.

ranges and concepts example. how does the following code work?

```cpp
std::vector<int> vi {0,1,2,3,4,5,6};

auto is_even = [](int i){return 0 == i %2;};
for (int i : ranges::filter_view(vi, is_even))
{
    std::cout << i  << " ";
}
```

lets look at _std::ranges::filter_view_,it uses some other concepts as part of the defintion. it also has a requires clause it's derived from _view_interface_ and follows the CRTP pattern. but what is that really? we can't inherit from a concept!

```cpp
template<input_range V, indirect_unary_predicate<Iterator_t<V>> Pred>
    requires view<V> && is_object_v<Pred>
    class filter_view: public view_interface<filter_view<v,Pred>>
    {
        //...
    };
```

lets looks at it more. does it get better? we have way to bring out the derived class. there are functions that are specialized depending on the derived class.

```cpp
template <class D>
    requires is_class_v<D> && same<D, remove_cv_t<D>>
    class view_interface
    {
        constexpr const D& derived () const noexcept
        {
            return static_cast<const D&>(*this);
        }

        //concept based specialization of operator[]
        //only applies if subclass is a random_access_range
        template<random_access_range R = const D>
        constexpr decltype(auto) operator[](range_difference_t<R> n) const
        {
            return ranges::begin(derived())[n];
        }
        //... more...
    };
```

there are some cases when we can't use CRTP. like when the whole class must be known for some reason.

### Writing Concepts

> - concept details 102
> - writing _sleep_for_ with concepts
> - good concepts, bad concepts

here's a nice quote:

> "Everything Should be made as simple as possible, but not simpler" - Albert Einstein (maybe)

how are constraints evaluated? set of steps, first normalizing,then subsumption,which is some ordering of the concepts,

> - concepts subsume, arbitrary expression do not
> - general principle is "'more constrained' is better match"

[godbolt example](https://godbolt.org/z/z7bT1aas6), the _std::signed_integral_ is "stronger" than _std::integral_ in terms of specialization, because it contains the lesser concept (it's defined in terms of that).

if we pull up the type_trait, things become ambiguous, because type traits don't subsume.

```cpp
template <class T>
requires std::is_integral_v<T>
struct wrapper<T>{};
```

we want concepts to be more than one operation, they should express **more** than just what an algorithm needs, and should be based on some domain. operations come in groups, numbers have mathematical operations (plus, minus, multiply, etc...), and containers have diffrent operatins (insert, erase, iteration, etc...).

getting concepts right is hard to do, it's done in iterations, and requires a lot of compiling.

we will try to make a concept for _std::chrono::sleep_for_, which takes either a _time_duration_ or a _time_point_. maybe we want to use boost data types instead.

```cpp
// sleep for duration
void sleep_for(time_duration d);

// sleep until time
void sleep_for(time_point t);
```

[godbolt example](https://godbolt.org/z/3vreqf)

under the hood there are _sleep_for_ and _sleep_until_ with different signatures.
so we go and try to reverse engineer the requirements for the function and see how the types are really used.

requirements:

> - a constant _zero_ member function to return th zero value
> - a comparison operator (less equal)
> - ability to cast/retrive the seconds and milliseconds of the duration
> - a constant _count_ function that cast to _long_ and _std::time_t_

lets try to make a concept:
[godbolt example](https://godbolt.org/z/1Tvefcasb)

```cpp
#include <concepts>
#include <chrono>

template <class T>
concept time_duration = std::totally_ordered<T> and requires(const T& v)
{
   v.count();
   v.zero();
};

static_assert( time_duration<std::chrono::seconds> );
```

we get some problems with floating point.

another draft [godbolt example](https://godbolt.org/z/W5odTv)

a draft with boost [godbolt example](https://godbolt.org/z/W53PGP) - this doesn't work.

should we refactor the concept? boost library? we can create a converter between boost and chrono so it satisfies our concept.

we choose to split appart the concepts into different parts: a time duration access and the time duration.
[godbolt example](https://godbolt.org/z/WW1dGM)

```cpp
template<class T>
concept time_duration_access = requires(const T& v)
{
   v.count();
   v.zero();
};

template <class T>
concept time_duration =
    std::totally_ordered<T> && time_duration_access<T>;

static_assert( time_duration<std::chrono::seconds> );
```

the next step is to further split it apart and have different duration accesses concepts (one for the standard library, one for boosts) and to use `if constexpr` to decide between them and either use them directly or use built in constructor of the standard chrono library [godbolt example](https://godbolt.org/z/7eqexh).

so we have a parital concepts, a concept that maps for some stuff, but not for all.

```cpp
#include <iostream>
#include <chrono>
#include <thread>
#include <concepts>
#include <boost/date_time.hpp>

template<class T>
concept std_time_duration_access = requires(const T& v)
{
   v.count();
   v.zero();
};

template<class T>
concept boost_time_duration_access = requires(const T& v)
{
   v.total_milliseconds();
};


template <class T>
concept time_duration =
    std::totally_ordered<T> &&
    (std_time_duration_access<T> || boost_time_duration_access<T>);

static_assert( time_duration<std::chrono::seconds> );

void sleep_for( time_duration auto td)
{
  std::cout << "hello ";
  if constexpr (std_time_duration_access<decltype(td)>) {
    std::this_thread::sleep_for( td );
  }
  if constexpr (boost_time_duration_access<decltype(td)>) {
    auto d = std::chrono::milliseconds(td.total_milliseconds());
    std::this_thread::sleep_for( d );
  }
  std::cout << "there\n";
}

namespace bpt = boost::posix_time;

int main()
{
  sleep_for( std::chrono::seconds(2) );
  sleep_for( bpt::seconds(2) );
}
```

there are still problems, if we want to add other 'cases', we need to modify the code again. but we didn't have to modify any library.

### Designing with Concepts

[timestamp](https://youtu.be/IXbf5lxGtr0?t=2580)

> - What is design?
> - Review of some 'design principles'
>   - Dry, Wet "don't repeat yourself" and "write everything twice"
>   - Solid
>     - single responsibility
>     - open-closed
>     - Liskov substitution
>     - interface segregation
>     - dependency inversion
>   - KISS "keep it simple stupid" (aka Occam's razor)
> - Concepts and dependencies
> - Impact on multi-paradigm design in c++
>   - structures
>   - functional
>   - generic
>   - object oriented
> - Concept serialization

we have design principles for divide and conquer, we want decomposition to make code manageable

- break programs into manageable parts
- parts that can be tested
- parts that can be reasoned about
- seprate concerns

functions with too many lines won't be changed once they're written, people will be afraid to touch them.

but once we decompose programs, the parts become dependant on one another, so we need to de-couple them somehow. so maybe the dependency is the point?

#### concepts and dependencies

> - move dependency to an abstraction from a type
> - simple to test a type that modesl a concept
> - however, problems are just shifted:
>   - type may evolve to longer model concepts, working code might fail.
>   - concept may evolve so the type no longer models it, working code might fail.

the standard library warns that the concepts might change in the c++23 standard.

which is more likely to evolve? the concept or the type? probably the type.

#### code readability and evolution

```cpp
auto result = some_function(); //return type unknown, flexible
int result = some_function(); //return type known, not flexible at all. silent conversions might cause a bug.
time_duration auto result = some_function(); //return type unknown, but intent is clear, still flexible
```

this ties into the Liskov substition principle: concepts vs inheritance. however, concepts don't model pre and post conditions, so we can't prove substitutability in compile time (maybe _contracts_ in the future could help).

Information Hiding:\
An idea from the 70's that we kind of messed up on following through. concepts might bring us closer to this goal of abstraction.

> "The sequence of instructions necessary to call a given routine and the routine itself are part of the same module"
> Parnas - information hiding (1972)

Multi-paradigm design:

> "Most designs in fact have a nontrivial componenet that is not object-oriented"\
> James O. Coplien - Multi-paradigm design (1999)

how do we discover patterns (or concepts today)? we look at how things work (functionality) and their domain. naming and behavior commonality.

> - "abstractions that will remain stable over time"
> - a name often defines a common behavior.
> - behavior vs meaning:
>   - each overloaded or specialized function _should_ have different behavior.
>   - but the meaning should remain the same from the calling client POV.

implemting commonality and vairability:

> Commonality Techniques:
>
> - factor commonalities into a base class
> - factor policy into trains (policy based design)
> - value oriented programming - vocabulary types
> - factor commonalites to **concepts**
>
> Variation Tools:
>
> - pre-processor (build time)
> - inheritance (build or run time)
> - templates (build time)
> - overloading (build time)
> - **concepts** (build time)

positive variability - adding behavior.\
negative variations - removing behavior,hiding away. the `requires` clause, removing non supported class member based on type.

#### concept serialization

the serialization patterns was once OO based, where each class knew how to serialize itself, using archive-types. today boost it templated, but still very similar.

an archive-type is json, xml, etc.
we use some double dispatch and stuff.

- separation of concerns
- type data is nicely encapsulated - method to serialize is in the type.
- archive type is nicely encapsulated - only knows about fundamental types.

```cpp
class myType{
    int foo;
    std::string bar;
    std::vector<int> baz;

    //one method for both input and output
    template<class Archive>
    serialize(Archive& archive)
    {
        ar("foo",foo);
        ar("bar",bar);
        ar("baz",baz);
    }
};

//archive type example
class OutputArchive
{
    put(std::string name, const std::string s);
    put(std::string name, int i);
    put(std::string name, double d);
    /*
    and so one
    */
};
```

the problem is that we do don't have basetypes, so we can't get directions to write new ones, it creates ugly compile errors, and we still need external extentsions for collection types (vectors, maps, etc...)

but can we _conceptify_ this? what would it give us?

it would look something like this

```cpp
template<class T, class A>
concept serializable = requires(T val, a archive)
{
    val.archive<A>(); //this must compile
}
template<class A>
concept Archive = // some code

template<class A>
concept OutputArchive = Archive<A> and requires(A archive)
{
    put(string,int);
    put(string,double);
    put(string,string);
    /*
    and others
    */
}
```

> Advantages:
>
> - looks quite doable
> - fixes the docs/compile issues
>   - static assert your archive type
> - Allows refactoring of other subtle policies:
>   - archvie ordering
>   - devices like files or database also become policies
> - template aliases can help with collections

#### conclustion and resources

concepts still aren't the end-all solution

modern generic designs depend on customization point objects and tagged invokes, both of which express the desired variabilities, but aren't as clear as virtual methods.

P2279R0 paper:\
in rust there are traits, which are a language mechanism for customization points. we had some good ideas in the c++0x rejected concepts.

question from the chat:

</details>

[Main](README.md)
