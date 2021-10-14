<!--
ignore these words in spell check for this file
// cSpell:ignore Gábor Horváth Andrzej Krzemieński cassert Dargo invokable default_initializable Clonable
-->

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
