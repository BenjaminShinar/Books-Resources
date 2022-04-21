<!--
ignore these words in spell check for this file
// cSpell:ignore rtime conceptify Parnas Permutatic Distributic regtr Eigen padd Treap
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

## Practical TMP: A C++17 Compile Time Register Machine - Daniel Nikpayuk

<details>
<summary>
Trying to make a Register machine at runtime using template meta programming.
</summary>

[Practical TMP: A C++17 Compile Time Register Machine](https://youtu.be/HLFz7rWRlyk)

**TMP - template meta programming**

there will be demonstration at the end.

TMP is Turing Complete. it can do loop unrolling, it has paradigms such as CRTP and SFINAE,and even compile time Regular expression, this lecture says that a TMP Turing machine can be practical, it argues that there are seven bottlenecks that prevent this, the first two bottlenecks are theoretical roadblocks, while the others are practical roadblocks.

> 1. A Stack Machine
> 2. Continuation Passing
> 3. The Nesting Depth Problem
> 4. Interoperability
> 5. Organization Design
> 6. Debugging
> 7. Performance

### Stack Machine

There different ways to make a function or machine to be turing complete, Daniel uses variadic packs. the relation between the two concepts (stack, variadic packs) is the template parameter resolution have the ability to pattern match from the front of a variadic pack.

> "This means we can interpret a pack as a stack, and push and pop from the front:"
>
> ```cpp
> template<auto v0,auto... Vs>
> constexpr auto function()
> {
>   //...
> }
> ```
>
> when we call the above function, we can pass some parameter pack `Ws` which will match `function<Ws...>()`

however, template scope allows for only a single direct parameter pack, while we need at least two.

we need at least two two stacks for a turing machine because having tow stacks allows for random access for reading a writing or their memory states (similar two a "tape machine")
the solution is something like this:

```cpp
template <auto.... Stack1, auto... Stack2>
constexpr auto function(auto_pack<Stack2...>)
{
    //...
}
```

which can look like this eventually? if we take the first element from stack1 and put it on top of stack2, we can look at stack1 all the way down, which means that we have a random access into it. we can later just push the same data back from stack2 onto stack1.

```cpp
template <auto.... Stack1_Rest,auto.... Stack1_Front, auto... Stack2>
constexpr auto machine(auto_pack<Stack1_Front...,Stack2...>)
{
    //...
}
```

we've solved the memory issue, but we now require our constexpr function to be a finite state machine. in automata theory this is called a "transaction function". in terms of Register Machine Theory this is called a controller, we will add to the templated function at the top.

```cpp
template<MACHINE_CONTROLLER, auto... Stack1, auto... Stack2>
constexpr auto machine(auto_pack<Stack2...>)
{
    //...
}
```

> - A Register machine is made from labels, which in turn are made from instructions, this is the basis of real world assembly languages.
> - therefore, our stack machine also required a _controller language_
> - Daniel decided to model his language on a language from the book "Structure and Interpretation of Computer Programs" chapter 5.
>
> ```
> (assign <register-name> (reg <register-name>))
> (assign <register-name> (const <constant-value>))
> (assign <register-name> (op <operation-name>) <input_1>...<input_n>)
> (perform (op <operation-name>) <input_1>...<input_n>)
> (test (op <operation-name>) <input_1>...<input_n>)
> (branch (label <label-name>))
> (goto (label <label-name>))
> (assign <register-name> (label <label-name>))
> (goto (reg <register-name>))
> (save (reg <register-name>))
> (restore (reg <register-name>))
> ```

### Continuation Passing

going from one state to the next. according to Category Therory we need Monads.

we will use composition to achieve this.

```
f(x,cl(y)) compose g(y,c2(z)) := f(x, \y.g(y)(c2(z)))
```

each controller instruction behave the regular way, but rather than return the result, the pass it forward to next machine.

```cpp
template<MACHINE_CONTROLLER, auto... Stack1, auto... Stack2>
constexpr auto machine(auto_pack<Stack2...> Heap)
{
    return next_machine<MACHINE_CONTROLLER,Stack1>(Heap);
}
```

and now with a struct, we only need the name parameter, but the note parameter makes things easier to handle

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<MACHINE_CONTROLLER, auto... Stack1, auto... Stack2>
    static constexpr auto result(auto_pack<Stack2...>)
    {
        //...
    }
};
```

the information about the next machine is held in the controller. so we need an index telling us where we currently are within the controller, and some way to know what is the next instruction from the current index.

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<auto contoller, auto index, auto... Stack1, auto... Stack2>
    static constexpr auto result(auto_pack<Stack2...> Heap)
    {
        return machine<
            next_name(controller, index),
            next_note(controller, index)
            >::template result<
                controller, next_index(controller, index),
                Stack1...>
            (Heap);
    }
};
```

for continuation passing, we need a dispatch and an index, so in theory, we are done. in practice, it's better to have several dispatches and two indices.

```cpp
template<typename dispatches, auto controller, auto index1 auto index2, auto... Stack1, auto... Stack2>
static constexpr auto result(auto_pack<Stack2...>)
{
    //...
}
```

so the whole things ends up looking like this:

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<typename dispatches, auto controller, auto index1 auto index2, auto... Stack1, auto... Stack2>
    static constexpr auto result(auto_pack<Stack2...> Heap)
    {
        return machine<
            dispatches::next_name(controller, index1,index2),
            dispatches::next_note(controller, index1,index2)
            >::template result<
                dispatches, controller,
                dispatches::next_index1(controller, index1,index2),
                dispatches::next_index2(controller, index1,index2),
                Stack1...>
            (Heap);
    }
};
```

before adding more things, we simplify the namings

- dispatches - `n`
- controller - `c`
- index1 - `i`
- index2 - `j`

and now more to the practical parts

### Nesting Depth Problem

with continuation passing,each machine calls the next machine, and we quickly run out, this is enforced by compilers which set a 'total allowable depth'.

this mitgated by using something called _Trampolining_, which is returning from the continuation passing with an intermediate result. in effect, we reset our depth again and again so we never reach the maximum nesting depth. we get this by adding another index.

this index "Depth" controlles the execution. normally, we simply move to the next machine, but when we reach a certian depth, we return an intermediate results and go back to the topmost machine.

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<typename n, auto c, auto depth, auto i,auto j, auto... Stack1, auto... Stack2>
    static constexpr auto result(auto_pack<Stack2...> Heap)
    {
        return machine<
            n::next_name(c, depth, i, j),
            n::next_note(c, depth, i, j)
            >::template result<
                n, c,
                n::next_depth(depth),
                n::next_index1(c, depth, i, j),
                n::next_index2(c, depth, i, j),
                Stack1...>
            (Heap);
    }
};
```

for some reason this trampolining decrease the depth each time, so this is still a finite number of possbile depth. however, it pretty east to create an additional trampolling index that resets the first one, and so one.

### Interoperability

we need some way to perfrom the current instructions. but if we calls a non-machine function as a helper, we might run out of the allowed depth, so we must contstrait the types of functions we can call to only those with a known depth. the proposed solution is to extend the Heap into it's own stack.

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<typename n, auto c, auto depth, auto i,auto j, auto... Stack, typename... Heaps>
    static constexpr auto result(Heaps... Hs)
    {
        return machine<
            n::next_name(c, depth, i, j),
            n::next_note(c, depth, i, j)
            >::template result<
                n, c,
                n::next_depth(depth),
                n::next_index1(c, depth, i, j),
                n::next_index2(c, depth, i, j),
                Stack...>
            (Hs...);
    }
};
```

we effectively allow only fixed depth helper functions and other machines to be called, and when we call a machine, we cache the current controller as a heap and pass the current depth to the next machine.

there is another problem for interoperability, which is the probel of _typename vs auto_.

```cpp
template<MACHINE_CONTROLLER,typename... Stack, typename... Heaps>
static constexpr auto results(Heaps... Hs)
{
    //...
}
```

luckily for us, we don't have to decide betwen the two. we can do both!
we hide the typename as the inputput type of a void function, which is later retrived via pattern matching. this means we only auto packs for out machine.

```cpp
template<auto value>
struct value_cached_as_type{};
using x = value_cached_as_type<int(5)>;

template<typename type>
constexpr void type_cached_as_value(TYPE){};
constexpr y = type_cached_as_value<int>;
```

### Organization Desgin

we actually have the final form of our machine.

```cpp
template<>
struct machine<MN::(((name)),(((note))))>
{
    template<typename n, auto c, auto d, auto i,auto j, auto... Stack, typename... Heaps>
    static constexpr auto result(Heaps... Hs)
    {
        return machine<
            n::next_name(c, d, i, j),
            n::next_note(c, d, i, j)
            >::template result<
                n, c,
                n::next_depth(d),
                n::next_index1(c, d, i, j),
                n::next_index2(c, d, i, j),
                Stack...>
            (Hs...);
    }
};
```

- where do we go from here?
- how to build a register machine library?
- which specific machine fo we choose to make?
- how do we organize our machines in this library?

this becomes a bottleneck because of how it affect performance, maintenance and debugging.

> the suggested hierarchy is:
>
> 1. Block machines: (atomics 1: pop, push, fold, etc). patterns of two.
> 2. Variadic machines: (atomics 2: pop, push, fold, etc). generalized forms of blocks.
> 3. Permutatic machines: (linear 1: stack/heap operators). linear controllers.
> 4. Distributic machines: (linear 2: erase, insert, replace)
> 5. Near linear machines: (1-cycle loops: lift, stem, cycle). aren't strictly needed, but are helpful
> 6. Register machines (branch, goto, save,restore).

### Debugging

a bottleneck for any TMP project. the error messages are usually horrible, no matter the compiler used.

for now, we use basic tools like `static_assert`, and in c++20 we can use concepts. at the controller level things are actually more managable,

in this example, we have a filter controller, it has three labels.

```cpp
constexpr auto r_filer_contr=r_controller
<
    r_label // is loop end:
    <
    test < eq,,c_0>,
    branch<return_pack>,
    apply<n,sub,n,c_1>,
    check<condition, val>,
    branch<pop_value>,
    rotate_sn<val>,
    goto_contr<is_loop_end>
    >,
    r_label // pop value:
    <
        erase<val>,
        goto_contr<is_loop_end>
    >,
    r_label // return pack:
    <
        pop<six>,
        pack<>
    >
>;
```

at this state, we can use old time tested debugging techniques, like adding a print statement, which doesn't really print, but it is a way to mark where the function got to, so we simply compile it again and again with moving the position of the statement.

```cpp
//...
    r_label // is loop end:
    <
    test < eq,,c_0>,
    branch<return_pack>,
    apply<n,sub,n,c_1>,
    stop <val>, //halts, returns val instead.
    check<condition, val>,
    branch<pop_value>,
    rotate_sn<val>,
    goto_contr<is_loop_end>
    >,
//...
```

if we don't like the naive approach, compile time debuggers are possbile, altought the don't currently exists.

### Performance

the final bottleneck, how good is this the design? there are currently no benchmarks.

**Block optimization** paradigm. _fast tracking_, performing calculations on variadic packs in powers of two (blocks) rather than one at a time. we match on more than just one element of the pattern.\
this turns a linera algorithm into a logarithmetic one, the savings aren't from the blocking itself, it's because we copy less information, which also reduced depth.

```cpp
//variadic form
auto foo(V0, Vs...);
//adding block forms
auto foo(v0,V1,Vs...);
auto foo(v0,V1,V2,V3,Vs...);
```

**Mutator optimization**. some instructions are more common than others (erase, insert, replace), so we create both the general purpose versions of them and optimized versions for the first eight registers.

**Machine Call optimization**. finding the next machine to call is done by using overloaded functions (for registers), class specializations (for names and notes) and combinning them together.

**Dispatch optimization**. not only used for moving from one machine to the next, but also optimize, therefore having less stack copying, and less depth issues.

### Demonstartion

- Factorial function
- Fibonacci function
- Filter and function compostion

Naive factorial example, only goes up to factorial(20) before overflowing.

```cpp
template<
    // registers:
    index_type val = 0,
    index_type n = 1,
    index_type eq = 2,
    index_type sub = 3,
    index_type multi = 4,
    index_type c_1= 5,
    index_type cont = 6
    // labels:
    index_type fact_loop = 1,
    index_type after_fact = 2,
    index_type base_case = 3,
    index_type fact_done = 4,
>
constexpr auto fact_contr = r_controller
<
    r_label // fact loop:
    <
        test<eq,n,c_1>,
        branch<base_case>,
        save<cont>
        save<n>,
        apply<n,sub,n_c_1>,
        assign<cont,after_fact>,
        goto_contr<fact_loop>
    >,
    r_label // after fact:
    <
        restore<n>,
        restore<cont>,
        apply<val, multi, n, val>,
        goto_regtr<cont>

    >,
    r_label // base case:
    <
        replace<val,c_1>,
        goto_regtr<cont>
    >,
    r_label // fact done:
    <
        stop<val>,
        reg_size<seven>
    >
>;
```

naive fibonacci, this one goes up to 13 before reaching the maximum nested depth.

```cpp
template<
    // registers:
    index_type val = 0,
    index_type n = 1,
    index_type less_than = 2,
    index_type add = 3,
    index_type sub = 4,
    index_type c_1= 5,
    index_type c_2= 6,
    index_type cont = 7
    // labels:
    index_type fib_loop = 1,
    index_type after_fib_n_1 = 2,
    index_type after_fib_n_2 = 3,
    index_type immediate_answer = 4,
    index_type fib_done = 5,
>
constexpr auto fact_contr = r_controller
<
    r_label // fib_loop:
    <
        test<less_than,n,c_2>,
        branch<immediate_answer>,
        save<cont>
        assgin<cont,after_fib_n_1>,
        save<n>,
        apply<b,sub, n,c_1>,
        goto_contr<fib_loop>
    >,
    r_label // after_fib_n_1:
    <
        restore<n>,
        restore<cont>,
        apply<n,sub, n,c_2>,
        save<n>,
        assign<cont, after_fib_n_2>,
        save<val>,
        goto_contr<cont>
    >,
    r_label // after_fib_n_2:
    <
        replace<n,val>,
        restor<val>
        restore<cont>,
        apply<val,add,val,n>,
        goto_regtr<cont>
    >,
    r_label // immediate_answer:
    <
        replace<val,c_1n>,
        goto_regtr<cont>
    >,
    r_label // fib_done:
    <
        stop<val>,
        reg_size<eight>
    >
>;
```

compose functions - still runtime?

```cpp
template<typename T>
constexpr T square(T x){return x*x;}

constexpr void _id_();// defined elsewhere
//2x^4+1
constexpr auto func_1 = do_compose
<
    square<int>,
    square<int>,
    multiply_by<int,2>,
    add_by<int,1>
>;
//(3(x+1))^2
constexpr auto func_2 = safe_do_compose
<
    add_by<int,1>
    multiply_by<int,3>,
    _id_,
    square<int>
>;
```

</details>

## Better C++ Ranges - Arno Schödl

<details>
<summary>
Trying for more efficient ranges.
</summary>

[Better C++ Ranges](https://youtu.be/P8VdPsLLcaE),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/04/Arno-Scho%CC%88dl-Better-C-Ranges.pdf), [ranges on github](https://github.com/think-cell/range)

another ranges library implementation, this time by **think-cell**, with the _tc_ namespace.

vector is probably the most used data struture, and until c++20, it wsa used with iterators.

this example show how to remove duplicates from a vector using the erase-remove idiom. **actually, theres a possible bug here!**

```cpp
std::vector<T> vec= ...;
std::sort(vec.begin(), vec.end());
vec.erase(std::unique(vec.begin(), vec.end()),vec.end());
```

a simpler way would look like this: it still uses the end iterator, but it also knows how to get them from the object itself.

```cpp
std::vector<T> vec= ...;
std::sort(vec);
vec.erase(std::unique(vec),vec.end());
```

the bug is between the implementations of `std::sort` (uses the less than operator) and `std::unique` (uses the equality operator), this might cause a problem if our type is not using the same behavior for the two operators.

a wrapper function:

```cpp
tc::sort_unique_inplace(vec);
tc::sort_unique_inplace(vec, less); //use the less than operator as predicate
```

ranges can be containers, they own the elements, copying them is deep copying (O(n)) and they have deep constness, if the range is const, than so are it's elements.

views are built on top of ranges, it references elements, doesn't copy elements and it's constness is indepent of the elements.

```cpp
template<typename It>
struct subrange {
    It m_itBegin;
    It m_itEnd;
    It begin() const
    {
        return m_itBegin;
    }
    It end() const
    {
        return m_itEnd;
    }
};
```

two different ways to get the same results:

```cpp
std::vector<int> v {1,2,4};
auto it = ranges::find(v,4); //first element of value 4
```

but if we want to use a predicate.

```cpp
std::vector<int> v {1,2,4};
struct A{
    int id;
    double data;
};
auto it = ranges::find_if(v,[](A const& a){return a.id==4;}); // first element of value 4 in id
```

we can use a transform adaptor from the ranges module. we use the piping operator _|_ which does the work lazily.

```cpp
std::vector<A> v {1,2,4};
struct A{
    int id;
    double data;
};
auto element = ranges::find(
    v| views::transform(std::mem_fn(&A::id)),
    4
); // first element of value 4 in id

auto it = ranges::find(
    v| views::transform(std::mem_fn(&A::id)),
    4
).base(); // iterator pointing to an element of type A
```

an example of implementing a Transform Adapater. do we really need to carry around the function in each iterator?

```cpp
template<typename Base, typename Func>
struct transform_view {
    struct iterator {
        private:
        Func m_func; // in every iterator, hmmm...
        decltype(ranges::begin(std::declval<Base&>())) m_it;
        public:
        decltype(auto) operator*() const {
            return m_func(*m_it);
        }
        decltype(auto) base() const {
            return (m_it);
        }
        ///...
    };
}
```

another adaptor is the filter adaptor, this is also a lazy evaluator, it only acts when it's being iterated on.

```cpp
std::vector<A> v{/*... */};
auto rng = v |views::filter([](A const& a){return a.id ==4;});
```

an implementation can look like this, which again carries the function in each function. the static cast to bool protects us from an overloaded version of the not operator.

```cpp
template<typename Base, typename Func>
struct filter_view {
    struct iterator {
        private:
        Func m_func; // functor and TWO iterators!
        decltype( ranges::begin(std::declval<Base&>())) m_it;
        decltype( ranges::begin(std::declval<Base&>())) m_itEnd;

        public:
        iterator& operator++() {
            ++m_it;
            while (m_it!=m_itEnd && !static_cast<bool>(m_func(*m_it)) )
            {
                ++m_it;
            }
            // why static_cast<bool> ?
            return *this;
        }
    //...
    };
};
```

we can stack filter on top of one another

```cpp
views::filter(m_func3)(views::filter(m_func2)(views::filter(m_func1, ...)))
```

this might cause an expositional expolision, each iterator contains a function, and it might carry other iterators, which also contain their own functions...\
so the more efficient way is to have the adaptor objects itself carry everything that is common for all iterators, which is usually the function and the end iterator. this is how c++20 ranges operate. and iterator cannot outlive their range (unless they are declared to be `std::ranges::borrowed_range`)

this doesn't compile, the range is an R-value, and it might go out of scope. there is no actual problem, because we use `.base()`, but the compiler still won't allow it.

```cpp
auto it = ranges::find(
    v | views::transform(std::mem_fn(&A::id)),4).base(); //doesn't compile
```

a fix would be to transform the value into an lvalue.

```cpp
auto it = ranges::find(
    tc::as_lvalue(v | views::transform(std::mem_fn(&A::id)))
    ,4).base(); //now it compiles
```

### Index Concept

this is something that they (think cell) created, an index is like an iterator, but it keeps the iterator size small by always using the range object. and also a way to use indexers like iterator.

```cpp
template<typename Base, typename Func>
struct index_range {
    //...
    using Index=...;
    Index begin_index() const;
    Index end_index() const;
    void increment_index( Index& idx ) const;
    void decrement_index( Index& idx ) const;
    reference dereference( Index const& idx ) const;
    //...
};

template<typename IndexRng>
struct iterator_for_index {
    IndexRng* m_rng
    typename IndexRng::Index m_idx;
    iterator& operator++() {
        m_rng.increment_index(m_idx);
        return *this;
    }
//...
};
```

a filter view based on index. either iterate over everything or until a match.

```cpp
template<typename Base, typename Func>
struct filter_view {
    Func m_func;
    Base& m_base;
    using Index=typename Base::Index;
    void increment_index( Index& idx ) const {
        do {
            m_base.increment_index(idx);
        }
        while (idx!=m_base.end_index() && !static_cast<bool>(m_func(m_base.dereference_index(idx))));
    }
};
```

views use references, and are shallow. it works fine for lvalue ranges, but not for rvalues. and we can't always overcome this.

```cpp
auto v = create_vector();
auto rng = v | views::filter(pred1); //this works

auto rng2 = create_vector | views::filter(pred1); //this doesn't work

auto foo(){
    auto vec = create_vector();
    return std::make_tuple(vec, views::filter(pred)(vec)); // DANGELING REFERNCE!
}
```

their range implementation tries to solve this.

this also an issue with algorithm, if we don't find a match, we return the end iterator, and we have to check against it. the lecture uses a customization point. so we can choose the returning element as part of the template which we provide, so we can create a default, a nullable iterator, or fail the program when there is no match.

generator ranges, traversing. currently the standard only supports external. top of stack and bottom of stack behavior for consumer and producer. it is theocratically possible to use co-routines to get both to the bottom of the stack, but it's a mess.

(stackfull and stackless co-routines)

internal iteration is often good enough, even if it's on the top of the stack. adaptors can also help with this.

| Algorithm     | Internal Iteration           |
| ------------- | ---------------------------- |
| find          | no (single pass iterators)   |
| binary search | no (random access iterators) |
| for_each      | yes                          |
| accumulate    | yes                          |
| all_of        | yes                          |
| any_of        | yes                          |
| none_of       | yes                          |
| tc::filter    | yes                          |
| tc::transform | yes                          |

using an enum in `any_of` implementation.

```cpp
template <typename Rng1, typename Rng2>
struct concat_range{
    private:
        using Index1= typename_range_index<rng1>::type;
        using Index2= typename_range_index<rng2>::type;

        Rng1& m_rng1;
        Rng2& m_rng2;
        using index::std::variant<Index1,Index2>;
    public:
    //..

```

index iterators branches out for each increment in the implementation of concat (merging two ranges)

```cpp
void increment_index(index& idx){
    std::visit(tc::make_overload(
        [&](Index1& idx1){
            m_rng1.increment_index(idx1);
            if (m_rng1.at_end_index(idx1)) {
                idx = m_rng2.begin_index();
            }
        }, [&](Index2& idx2)){
            m_rng2.increment_index(idx2);
        }
    ),idx);
}
```

also branching out in the derefence of the index.

```cpp
auto derefence_index(Index const& idx) const{
    std::visit(tc::make_overload(
        [&](Index1 const& idx1){
            return m_rng1.derefence(idx1);
            }
        },
        [&](Index2 const& idx2)){
            return m_rng2.derefence(idx2);
        }
    ),idx);
}
```

if if turn this into a generator, things are a bit better.

```cpp
template <typename Rng1, typename Rng2>
struct concat_range{
    private:
        Rng1& m_rng1;
        Rng2& m_rng2;
    public:
    //..

    // version for non-breaking func, without the break/continue enum
    template <typename Func>
    void operator()(Func func){
        tc::for_each(m_rng1,func);
        tc::for_each(m_rng2,func);
    }
}
```

### Formatting

we can use ranges instead of `std::format`, a single unifying concept.

```cpp
double f=3.14;
tc::concat("You won ", tc::as_dec(f,2), " dollars.");
```

it's easy to extend this and create custom formatters. which are all lazily evaluated as ranges.

```cpp
auto dollars(double f){
    return tc::concat("$",tc::as_dec(f,2));
}
double f=3.14;
tc::concat("You won ", dollars(f), " dollars.");

tc::concat(
    "<body>",
    html_escape(tc::placeholders("You won {0} dollars.", tc::as_dec_f,2)),
    "</body>"
);
```

even support for names (rather than location)

```cpp
double f= 3.14;
tc::concat(
    "<body>",
    html_escape(tc::placeholders("You won {amount} dollars on {date}.",
    tc::named_arg("amount",tc::as_dec(f,2)),
    tc::named_arg("date",tc::as_ISO8601(std::chrono::system_clock::now()))),
    "</body>"
);
```

> - Formatting parameters (#decimal digits, etc) not part of format string.
>   - Internationalization: translator can rearrange placeholders, but not change parameters.

formatting into containrs:

with `std::string` we have many constructors (actually, an absurd amount), so they suggest adding range constructors

```cpp
//existing:
std::string s1; // empty construction
std::string s2("Hello"); // construction from a string literal
std::string s3(s2);// copy constructor
//suggested:
std::string s4(tc::as_dec(3.14,2)); //one range
std::string s5(tc::concat("You Won ",tc::as_dec(3.14,2)," dollars.")); // concatenated range
std::string s6("Hello", " World"); // 2 Ranges
std::string s7("You Won ",tc::as_dec(3.14,2)," dollars."); // N Ranges
```

because of std::string has so many constructors, these might run into conflicts.

_the favorite game: guess the string constructor output!_

```cpp
std::string sc1("A",3); //UB , tries to take the substring of size 3 from the string "A",
std::string sc2('A',3); // probably a bug 65 times Ctrl-c, 65 times of the char 3
std::string sc3(3,'A'); // probably what we wanted, 3 times the char 'A', "AAA"
```

maybe it's better to deprcatee them and be unambiguous? or if that's impossible(because of backward compatibility), to use pseudo constructor (explicit casting).

```cpp
//suggested
std::string s(tc::repeat_n('A',3));

// in the tc library
auto stc1 = tc::explicit_cast<std::string>("Hello", " World");
auto stc2 = tc::explicit_cast<std::string>("You won ", tc::as_dec(f,2), " dollars.");
```

also a wrapper for `emplace_back` and `push_back`: `tc::cont_emplace_back`. this uses the explicit casting as needed. rather than the usual bracket initialization syntax. also append

```cpp
std::vector<std::string> vec;
tc::cont_emplace_back(vec, tc::as_dec(3.14,2));

std::string s;
tc::append(s, tc::concat("You won", tc::as_dec(3.14,2), " dollars."));
tc::append(s, "You won", tc::as_dec(3.14,2), " dollars.");
```

### Fast Ranges Append

how to do this fast?

- determine string length
- allocate memory for whole string at once
- fill in characters

simple implementation for casting into a container. the problem is that for non-random-access ranges, the string constructor runs twice over the range, once to determine the size, and once to copy the characters. this is fine when iterators are cheap to copy and iterate over, but for adaptors it isn't.

```cpp
template <typename Container, typename Rng>
auto explicit_cast(Rng const& rng){
    return Container(ranges::begin(rng),ranges::end(rng));
}
```

lets try this differently, _rng_ implements a `size()` member, and we use an explicit loop taking advantage of `std::size`.

```cpp
template <typename Container, typename Rng, std::enable_if
</*Rng has size member and is not random-access */>>
auto explicit_cast(Rng const& rng){
    Container cont;
    cont.reserve(std::size(rng));
    for (auto it = ranges::begin(rng); it != ranges::end(rng),++it){
        tc::cont_emplace_back(cont, *it);
    }
    return cont;
}

template <typename Container, typename Rng, std::enable_if
</*Rng has size member and is not random-access */>>
void append(Container& cont,Rng const& rng){
    cont.reserve(cont.size() + std::size(rng));
    for (auto it = ranges::begin(rng); it != ranges::end(rng),++it){
        tc::cont_emplace_back(cont, *it);
    }
}
```

**oops! there is a problem. `.reserve` has issues!**\
reserve takes the exact neccesary size, so multiple calls will generate multiple allocations, even if it's byte by byte.

a better thing it to create an alternative `reserve` function, one that that

> "When adding N elements, guarantee **O(N)** moves and **O(log(N))** memory allocations!"

this is already what most containers do anyways, they increase in size in a logarithetc scale.

```cpp
template <typename Container>
void cont_reserve(Container & cont, typename Container::size_type n)
{
    if (cont.capacity() < n){
        cont.reserve(max(n, cont.capacity()*(8/5)));
    }
}

template <typename Container, typename Rng, std::enable_if
</*Rng has size member and is not random-access */>>
void append(Container& cont,Rng const& rng){
    tc::cont_reserve(cont, cont.size() + std::size(rng)); //use the better reserve
    for (auto it = ranges::begin(rng); it != ranges::end(rng),++it){
        tc::cont_emplace_back(cont, *it);
    }
}
```

but what happen when we don't have random access? how do we get the size? what about generator ranges (not iterator based ranges)?

for this, we introduce an appender, sink for `explicit_cast` and `append` to use.

```cpp
template <typename Container, typename Rng>
void append(Container & cont, Rng && rng)
{
    tc::for_each(std::forward<Rng>(rng), tc::appender(cont));
}
```

> - appender customization point
>   - returned by `container::appender()` member function.
>   - default for std containers.

```cpp
template <typename Container>
struct appender{
    Container & m_cont;

    template <typename T>
    void operator()(T&& t)
    {
        tc::cont_emplace_back(m_cont,std::forward<T>(t));
    }
}
```

in the basic case, this is the same as `std::back_inserter`.

**Chunk customization point**
we can pass in elements piece by piece or by chunks,

```cpp
template <typename Container, enable_if</*Container has reserve()*/>
struct reserving_appender:appender<Container>{
    template <typename Rng, enable_if</*Rng has size()*/>
    void chunk(Rng && rng) const
    {
        tc::cont_reserve(m_cont, m_cont.size() + std::size(rng));
        tc::for_each(std::forward<Rng>(rng), static_cast<appender<Cont> const&>(*this));
    }
}
```

file sink advertises interst in contiguous memory chunks, using chunks and not memory buffer. uniform treatment of string and files.

```cpp
struct file_appender
{
    void chunk(std::span<unsigned char const> rng) const
    {
        std::fwrite(rng.begin(),1, rng.size(),m_file);
    }
    void operator()(unsigned char ch) const
    {
        chunk(tc:single(ch));
    }
};
```

### Performance

trivial formatting task: writing 'A' 10 times, then 'B' 10 times and 'C' 10 times.

```cpp
struct Buffer
{
    char achBuffer[1024];
    char* pchEnd = &achBuffer[0];
} buffer;

void repeat_handwritten(char chA, int cchA,char chB, int cchB,char chC, int cchC)
{
    for (auto i = cchA; 0 < i; --i)
    {
        *buffer.pchEnd=chA; //write to buffer
        ++buffer.pchEnd; // move ptr in buffer
    }
    // repeat loop for chB and cchB
    // repeat loop for chC and cchC
}


struct AppenderBuffer
{
    //...
    auto appender() & {
        struct appender_t{
            AppenderBuffer* m_buffer;
            void operator()(char ch) noexcept{
                *m_buffer->pchEnd=ch;
                ++m_buffer->pchEnd;
            }
        };
        return appender_t{this};
    }
} appender_buffer;

void repeat_with ranges(char chA, int cchA,char chB, int cchB,char chC, int cchC)
{
    tc::append(appender_buffer, tc::repeat_n(chA,cchA), tc::repeat_n(chB,cchB), tc::repeat_n(chC,cchC));
}
```

the results show that the overhead cost is minimal, about 50% when `repeat_n`is iterator based, and about 15% when it supports internal iteraton. the overhead gets even smaller when the task isn't trivial, like when we need to convert numbers to strings.

and example of re-implemanting the basic string using ranges, getting some improvements over the standard. there are some other stuff to work on.

</details>

## STL Algorithms as Expressions - Oleksandr Bacherikov

<details>
<summary>
Using Expression templates to implements algorithms and customizing them with policies.
</summary>

[STL Algorithms as Expressions](https://youtu.be/ehtBUKHNJlw),[slides](https://docs.google.com/presentation/d/1Jr20Xa6dK8wU7W1SWlPScCQfnIx4Ng-mqPMGn0zYqLQ/edit?usp=sharing), [github](https://github.com/AlCash07/ACTL/tree/master/include/actl/operation)

challenge: comparing floats (used in computational geometry), there is not single way to compare floats that always yields a correct result. different thresholds are used in different places, and computations might require casting to integers (and then there's a danger of overflowing).

> "Basic operations such as comparison, multiplication and division are all customization points. Most STL algorithms have just 1 or 2 two customization points."

one example of customization is by using **traits**.

### Expression Templates

we can look at the example of vectors addition to understand template expressions.

given three vectors: a,b, and c, we want an operation to add them together and produce the sum.

```cpp
//not std::vector
Vector operator + (const Vector &lhs, const Vector &rhs)
{
    assert(lhs.size() == rhs.size());
    Vector sum(lhs.size());
    for (size_t i = 0; i< lhs.size() ;++i)
    {
        sum[i]= lhs[i]+rhs[i];
    }

    return sum;
}
```

the problem is that we create an additional temproray Vector for each sub expression `sum = a+b+c`.

in linear algebra libraries, it was solved in the past as an expression template. there is no computation overhead to create the class itself.

```cpp
template <class T, class U>
struct VectorSum
{
    // use references
    const T& lhs;
    const U& rhs;

    size_t size() const
    {
        assert(lhs.size() == rhs.size());
        return lhs.size();
    }

    auto operator[](size_t i) const
    {
        return lhs[i]+rhs[i];
    }
};

template <class Vector1,class Vector2>
VectorSum<Vector1,Vector2> operator + (const Vector1 &lhs,const Vector2 &rhs)
{
    return VectorSum<Vector1,Vector2>{lhs,rhs};
}
```

with this, adding the Vectors doesn't construct a new temporary vector, only the actual left hand side assignment does real work. this is also called _lazy/deferred evaluation_.

in the C++ Standard Library there are expressions: the adaptors and views from the _std::ranges_ library.

```cpp
std::vector<int> const vi{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
auto rng = vi |
        views::filter([](int i) { return i % 2 == 0; }) |
        views::transform([](int i) { return std::to_string(i)});

// prints: [2,4,6,8,10]
std::cout << rng << '\n'; // expression isn't evaluated until this line

//we never do the work on the unreachable number
int sum = accumulate(views::ints(1, unreachable)
        | views::transform([](int i) {return i * i;}) //square
        | views::take(10), // take the first 10
        0); // initial value
// prints: 385
std::cout << sum << '\n';
```

### Designing Operations to support Operations.

the C++ Standard Library has arithemetic operations and comparison operators defined (e.g `std::plus`, `std::greater_equal`), which can be passed to algorithms, but they are defined in terms of the operators themselves, and don't help with the actual implementation.

Scalar operations in the Eigen Linear Algebra library

- wrappers to support vector operations.
- function object wrap all the functions ins _\<math>_.

```cpp
template <typename LhsScalar, typename RhsScalar>
struct scalar_sum_op : binary_op_base<LhsScalar, RhsScalar>
{
    const result_type operator()(const LhsScalar& a, const RhsScalar& b) const
    {
        return a + b;
    }

    template <typename Packet>
    const Packet packetOp(const Packet& a, const Packet& b) const
    {
        return internal::padd(a, b);
    }
};
```

extensible plus expression support

```cpp
namespace scalar
{
    struct add_f : operation<add_f>
    {
        template class<T, class U>
        static auto evaluate(T lhs, U rhs)
        {
            return lhs + rhs;
        }
    }
    inline constexpr add_f add;
}

struct add_f: operation<add_f>
{
    static constexpr auto formula = scalar::add;
};
inline constexpr add_f add;
```

and also

```cpp
template <class T, class U, enable_operators<T,U>=true>
auto operator + (T&& lhs, U&& rhs)
{
    return add(std::forward<T>(lhs), std::forward<U>(rhs));
}

template <class T, class U, enable_operators<T,U>=true>
decltype(auto) operator += (T&& lhs, U&& rhs)
{
    return add(inout{std::forward<T>(lhs), std::forward<U>(rhs)});
}
```

(this required traits and sfinae).

operation based class, using the Curiosly Recurring Template Patterns

```cpp
template <class Derived> // Curiously Recurring Template Pattern
struct operation
{
    template <class... Ts>
    constexpr decltype(auto) operator()(Ts&&... xs) const
    {
        if constexpr (!is_any_inout_v<Ts...>)
        {
            return expression{derived(), std::forward<Ts>(xs)...};
        }
        else
        {
            //extention to arbitrary types
            auto&& op = resolve_overload<Ts...>(derived());
            static_assert(1 == (... + is_inout_v<Ts>));
            auto& dst = find_dst(xs...);
            op.evaluate_to(dst, remove_inout(xs)...);
            return dst;
        }
    }
};

const Derived& derived() const
{
   return static_cast<const Derived&>(*this);
}
```

> In order to support operations like this for a Vector class (including expressions):

```
    Vector a, b, c;
    Vector sum = a + b + c;
    sum += a;
```

> 1. Implement a Vector.
> 1. Make sure the operators are found via ADL (Argument dependant lookup), for example using an empty base class from the namespace with operators (like in _boost::operators_).
> 1. Implement an operation that applies a given scalar operation to the vectors element-wise (single one to handle all the scalar operations).
> 1. Write a rule that resolves universal operations (that are scalar by default) to this element-wise operation if its arguments are vectors.

input-output parameters wrapper - main purpose is to distinguish between pure operators and assignment operators (`+` and `+=`).

```cpp
template<class T>
struct inout
{
    T x;
};

template<class T>
inout(T&&) -> inout<T>;
```

overload resolution - distinguish scalares from ranges and tuples, currenlty template class specialization. something about broadcasting.

expression implementation

```cpp
template <class Op, class... Ts>
struct expression : expression_base<expression<Op, Ts...>, result_tag_t<Op, Ts...>>
{
    std::tuple<Op, Ts...> args;

    template <class... Us> // temporaries are moved, not stroed by refernce
    constexpr expression(Us&&... xs) : args{std::forward<Us>(xs)...}
    {}


    constexpr auto& operation() const
    {
        return std::get<0>(args);
    }
};
```

simplified code examples:

sum of squares

```cpp

int sum_before = accumulate(views::ints(1, unreachable) |
    views::transform([](int i) {return i * i;}) |
    views::take(10),
    0);

int sum_after = accumulate(sqr(views::ints(1, unreachable)) |
    views::take(10), 0);
```

range projections vs expressions, fold operations as expressions(including short-circuting folds)

### Alternatives to Lambda Expressions

revisiting how we define functions

```cpp
template <class T>
auto cos_derivative(const T& x)
{
    return -sin(x);
}
constexpr auto cos_derivative = -sin;
constexpr auto cos_derivative = -sin(x_); //placeholder-like syntax

double x = cos_derivative(1.0);
```

and also

```cpp
constexpr auto f = add * sub; // define expression operator
constexpr auto f = (lhs_ + rhs_) * (lhs_ - rhs_); // same
constexpr auto f = sqr(lhs_) - sqr(rhs_); // simplified
static assert(f(4,1)== 15); // (4+1)*(4-1)
```

(how to implement the placeholder like syntax, with some aliases).

more operations that can be defined in terms of basic template expressions.

```cpp
struct sqr_f : operation<sqr_f> {
    static constexpr auto formula = x_ * x_;
};
inline constexpr sqr_f sqr;

inline constexpr auto greater = rhs_ < lhs_;

struct cmp3way_f : operation<cmp3way_f>
{
    static constexpr auto formula = cast<int>(greater) - cast<int>(less);
};
inline constexpr cmp3way_f cmp3way;

inline constexpr auto none_of = !any_of;
```

using this syntax, we can write shorter lambda code, we remove the boiler plate code and replace it with expression code.

_add 2 to each elements_

```cpp
//before
ranges::transform(src, dst, [](auto x) { return x + 2; });
//after
ranges::transform(src, dst, x_ + 2);
```

_find single element larger than 2_

```cpp
//before
ranges::find_if(src, [](auto x) { return x > 2; });
//after
ranges::find_if(src, x_ > 2);
```

_transform all even elements to string_

```cpp
//before
vi
    | views::filter([](int i) { return i % 2 == 0; })
    | views::transform([](int i) { return std::to_string(i); });

//after
vi
    | views::filter(x_ % 2 == 0)
    | views::transform(to_string);
```

in the C++ Standard Library we have algorithm with the `_if` suffix, because they are a common use-case.

> - `std::count_if`
> - `std::find_if`
> - `std::find_if_not`
> - `std::find_first_of`
> - `std::remove_if`
> - `std::remove_copy_if`
> - `std::replace_if`
> - `std::replace_copy_if`
> - `ranges::find_if(range, predicate);`
> - `ranges::find(range, value);`

we can change them to use a single form.

| Usage                              | Current                               | Proposed (expression)                   |
| ---------------------------------- | ------------------------------------- | --------------------------------------- |
| by predicate                       | `find_if(range, pred)`                | `find(range, pred)`                     |
| by value                           | `find(range, value)`                  | `find(range, x_ == value)`              |
| not matching predicate             | `find_if_not(range, pred)`            | `find(range, !pred)`                    |
| first match in both                | `find_first_of(range1, range2)`       | `find(range1, is_in(x_, range2))`       |
| first match in both with predicate | `find_first_of(range1, range2, pred)` | `find(range1, is_in(x_, pred(range2)))` |

this can also be used to implement _upper_bound_ and _lower_bound_.

### Policy Based Design

the motivating example for the talk was comparing vectors of 3d points. we want to check the equality between them with a threshold.

```cpp
//normal form
std::vector<glm::vec3> expected, actual;
assert(expected.size() == actual.size());
for (size_t i = 0; i < expected.size(); ++i) {
    assert(std::abs(expected[i].x - actual[i].x) < 1e-6f);
    assert(std::abs(expected[i].y - actual[i].y) < 1e-6f);
    assert(std::abs(expected[i].z - actual[i].z) < 1e-6f);
}

//ranges form
assert(ranges::equal(expected, actual, compare_vectors{1e-6f}));
```

the proposal is to have operations affected by polices.

```cpp
std::vector<glm::vec3> expected, actual;
assert ((equal | absolute_error{1e-6f})(expected, actual));

//alias
auto my_equal = equal | absolute_error{1e-6f};
assert(my_equal(expected, actual));

```

the form is `operation | policy1 |policy2|....`, each policy effects the policy. so in our example

```cpp
template <class T>
struct absolute_error {
    struct is_policy;
    T threshold;
};


template <class T>
auto apply_policy(scalar::equal_f, absolute_error<T> policy) {
    return abs(sub) <= policy.threshold;
}
template <class T>
auto apply_policy(scalar::less_f, absolute_error<T> policy) {
    return policy.threshold < rhs_ - lhs_;
}
```

while there is a compiler flag for "fast math", it means that we can't have some functions use the fast math and some not, policies allow us to decide per case.

while c++ has execution policies, he doesn't like the form of the function. also the Generator function in `<random>` header and the allocators. we can have policies to effect how the allocator works.

for example, the _std::vector_ has a growth factor, but it's not up for the programmer to set, it's defined in the compiler. we could create a growth policy and pass it.

in the _std::unordered_set_ class:

```cpp
template <
    class Key,
    class Hash = std::hash<Key>,
    class KeyEqual = std::equal_to<Key>,
    class Allocator = std::allocator<Key>
    >
class unordered_set;
```

the _Hash_ and _KeyEqual_ classes must be consistent, but that is not enforced by the interface.

> Problems:
>
> - There are too many template parameters.
> - Custom allocator specification requires specifying defaults for all other parameters.
> - It’s easy to forget to specify correct KeyEqual unless it’s the default.
>
> Solution:
>
> - Introduce hashing policy that consistently affects hash and equality.
> - Compose hashing and allocator policy into a single Policy parameter.

Treap example - a data structure with many customization points that could make use of policy.

> Treap is a data structure used in Competitive Programming:
> https://en.wikipedia.org/wiki/Treap
>
> Treap can represent an array that supports:
>
> - slicing and concatenation
> - reverse on a subarray
> - query on a subarray (sum of elements)
> - modification on a subarray (increment)
>
> All in O(log N).
>
> Customizable operations:
>
> - allocation
> - random numbers generation
> - combination of queried values
> - reverse of queried value
> - modification of queried value
> - identity modification value
> - combination of modifying values
> - shift of modifying value
> - reverse of modifying value

### Generic Policy Based Mechanisms

customzing advanced algorithms, steps in creating the expression pipeline. applying the policy and the operation expresion with an expression tree.

templates for operation instead of types?

</details>

##

[Main](README.md)
