<!--
ignore these words in spell check for this file
// cSpell:ignore Gábor Horváth Andrzej Krzemieński cassert
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
