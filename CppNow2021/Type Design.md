<!--
ignore these words in spell check for this file
// cSpell:ignore fmodules Levehnstein cvref deducer psssint ultr defences moda unsynchronized mdspan Eigen edsl kiwaku
 -->

[Main](README.md)

Type Design

## How to: Colony - Matthew Bentley

<details>
<summary>
Explaining how the current plf::colony container works. it might become part of the standard library.
</summary>

[How to: Colony](https://youtu.be/V6ZVUBhls38)

this talk is about a data structure called **Colony** that is int the process of being added to the standard.

[PlfLib - Some Header-only libraries](https://plflib.org/)

the main thing about it is that it maintains pointers/iterators/reference integrity.

> use scenarios
>
> - you have a lot of unordered data and you're erasing/inserting on the fly.
> - you have multiple collections of interrelated data.
> - preserving the pointer/iterator validity of non-erased elements is important

started from the design of a game, entities that have bidirectional references and many interrelationships, with a lot of inserting and creating entities which link to one another.

an existing solution is a 'bucket array' and entries can be active or inactive to determine if they're processed or not.

### Core Aspects

> Three Core Aspects
>
> 1. A collection of element memory blocks + metadata, to prevent reallocation during insertions (as opposed to a single memory block).
> 2. A method of skipping erased elements in O(1) time during iterations (as opposed to reallocating subsequent elements during erasure, or doing individual allocation of elements)
> 3. An erased-element location recording mechanism, to enable the re-use of memory from erased element in subsequent insertions, which in turn increases cache locality and reduces the number of block allocations/de-allocations.

Avoid reallocations to avoid invalidating iterators/pointers/references. Reuse memory locations that have been erased, keep the data together for cache locality.

A collection of element memory blocks + metadata

> - Can do linked list of blocks, or vector of pointers to blocks.
> - Blocks have growth factor, so can not do vector of blocks.
> - Minimum/Maximum block capacities can be user-defined.
> - Can house block metadata separately, or together in a struct.
> - Metadata includes skip-field of other erasure-skipping mechanism, and any data related to erased-location recording.
> - Necessary metadata:
>   - size - to remove blocks once empty.
>   - capacity - to ascertain end-of-block.

we remove empty blocks to maintain iteration at O(1). but we can retain them as reserved blocks for later use.

considerations about which block to retain.
A method of skipping erased elements in O(1)

> Definitions:
>
> - LCJC:'low complexity jump-counting'.
> - HCJC:'high complexity jump-counting'.
> - Block: a colony's element memory block.
> - Skipfield: an array of intgers of bits used to skip over certain object in an accompanying data structure during iteration. Separate from the elements.
> - Skipblock: a run of skipped nodes within a skipfield.
>   - Start node: the first node in any skipblock.
>   - End node: the last node in any skipblock.
>   - Middle node: any non start/end node in skipblock.

booleans SkipFields are good because they simple, and might be usefull for multi-threaded environments (atomics, etc). they are prolemeatic because they aren't constant time (not O(1)), cause branching and latency.

low and high complexity jump counting:
time-complexity of algorithms differs for modification of fields, but both have O(1) for iteration.
HCJC allows recording of middle nodes, LCJC doesn't allow.

> acronym "Theyaton" - Traversing Homogenus Elements Yielding Amortised Time O(n).

> boolean skipfield example
> 0 0 0 1 1 1 1 1 1 0 0 0
> equivalent HCJC
> 0 0 0 6 2 3 4 5 6 0 0 0
> the 6's are the start and end nodes. the other numbers are Middle nodes.Start and End record the length of runs of erased elements. Middle node record left distance to first non erased element.

```
// skipping
++i;
i +=S[i];
```

> equivalent LCJC. multiple forms.
> 0 0 0 6 2 6 7 3 6 0 0 0
> 0 0 0 6 0 0 0 0 6 0 0 0
> 0 0 0 6 0 5 2 1 6 0 0 0

The middle nodes are ignored and have no meaning. The start and end nodes record the skip length.
good for recording and re-using skipblocks rather than individual skipped nodes. lower time complexity O(1) in all operations, where as HJCJ can have undefined time complexity, fewer instructions overall.

```
// skipping
do {
    ++i;
} while(S[i] == 1);
```

we can have per-memory-block skip fields or global skipfields. global skipfields can create problems of reallocation (invalidating iterators and causing latency). the bit-width of the skipfield must be able to describe the memory block, so it must be about the same size (capacity -1),

possible ideas with skipfields:

- using two 8 skipblocks instead of 16 bit skipblocks. forces some computations.
- using a boolean bitfield and storing the skipdata instrad of the elements. forces more memory reads.

colony once used a stack of pointers, which was problematic because it meant creating memory allocation during erasure, which is not up to the standarts.

**"Free List"**

> Linked list of erased elements (typically singly-linked) using erased element memory space reinterpert_cast'd via pointers as linked list nodes.
> requires over-alignment of the type to the width of a free list node.
> per-block (not global) free-lists reduce bit depth.
> a global free-list must use pointers (not indexes),also causes O(N) when erasure.

effect of removing a block requires something with the skip-block. something about doubly-linked free list.

summary of the current implementation, the container structure and the iterator structure.

example with a blackboard.

extra operations:

- advance/next/prev/distance/range-erase optimization.
- range/fill/initializer_list insert/assign optimization.
- splicing.
- sorting.

</details>

## Variations on Variants - Roi Barkan

<details>
<summary>
Variants, Unions, proposals for new kind of variants, standard-layout considerations.
</summary>

[Variations on Variants](https://youtu.be/YBXRiPKa_bc), [slides](https://docs.google.com/presentation/d/1W0QBblWpJ-AXsdo70kvtxg5G2RZLiqb-mewQKySTf6s/preview?slide=id.p)

> - What are variants
> - Variant vs unions
> - Intrusive variants
> - Streams of variants
> - Variants for de-virtualization

### Variants Introduction

std::variants, a typesafe union

> - CppReference.com: "the class template std::variants represent a type-safe union",
> - boost.org: "the variants class template is a safe, generic, stack-bases discriminated union container"
> - plain english: "a union that knows (holds) its type".

we can use std::get\<type> accessor to get the memberm or use std::visit.\
sum type, as opposed to product types
memory layout, the std::variant has some extra memory requirements for the tag.

> usages
>
> - State machines
> - Value Semantics for Dynamic Types
>   - commands
> - Success / Fail
>   - expected\<T>
> - Exist / void
>   - optional\<T>
> - Runtime Dispatch (polymorphism)
> - Pattern Matching

(roi looks at some old lectures from different people)

std::optional is a specialzed form of std::variant. \
runtime dispatch can allow us to avoid virtual function call.\
pattern matching is another way for dynamic, trying to do something similar to a switch case, perhaps using std::visit.

extra reading:\
[std::visit is everything wrong with modern C++](https://bitbashing.io/std-visit.html)

pattern matching vs concepts/contracts:\
concepts - compile time sfinae.\
contracts - run time assertions (not existing yet, but soon)\
pattern matching - run time advanced switch case (not existing yet, low priority), might have 'inspect' keyword with some new mechanics, combining values, lambda, dynamic values, different types, etc...

_inspect() in a nutshell_
switch + std::visit() + [structured, bindings]

example of using pattern matching for balancing a red/black tree.

### Variants vs Unions

the tag in the variant is private. only the constructor and assignemt operations can change the flag. the compiler knows in advance all the possible alternative forms.

this is a buggy code:\
missing break statements (fall through), missing cases and no default, calling the wrong function.

```cpp
union IdentityCard
{
    IDNumber nationalID;
    PassportNumber passport;
    UUID factoryCertificate;
};
enum IDType {CITIZEN,TOURIST,ROBOT};
void checkID(IdentityCard id, IDType type)
{
    switch(type)
    {
        case TOURIST: checkPassport(id.passport);
        case CITIZEN: checkPassport(id.passport);
    }
}
```

if we used variants, it would look like this. the compiler must account for all the cases, it's not just a warning if one is missing. we aren't handling the tag.

```cpp
using IdentityCard = std::variant<IDNumber,PassportNumber, UUID>;
void checkID(IdentityCard id)
{
    std::visit([](auto& x){x.check();},id);
}
```

the example for the C style code was a strawman, real union code looks more like a variant, with a struct that holds the tag (explicit header). or as a union with structs that each have the same header part (implicit).
the explicit type looks more organized and less crowded, but the implicit type allows direct access to the header, if we need data from it and we got a pointer/reference to the inner member data.

```cpp
struct IdentityCardExplicit
{
    IDType type;
    union value {
        IDNumber nationalID;
        PassportNumber passport;
    } value;
};

union IdentityCardImplicit
{
    struct Header {
        IDType id;
        };
    struct Citizen{
        IDType id;
        IDNumber nationalID;
        };
    struct Tourist{
        IDType id;
        PassportNumber passport;
        };
}
```

the header field can contain more than the tag itself, and we use a macro to define them.\
_(I don't like this, it feels like some semi-colons are missing)_

```cpp
#define HEADER_FIELDS \
IDType type; \
Date expiration;

struct Header { HEADER_FIELDS};
struct Citizen{
    HEADER_FIELDS
    IDNumber nationalID;
    };
struct Tourist{
    HEADER_FIELDS
    PassportNumber passport;
    };

union IdentityCardImplicit
{
    Header header;
    Citizen citizen;
    Tourist tourist;
}
```

> Keeping the C Layout
>
> - Header with type is common.
>   - Network protocols- TCP/IP, Finance, etc...
>   - File formats - ELF, etc...
>   - Serialization - [Cap'n Proto](https://capnproto.org/), [Apache Avro](https://avro.apache.org/)
> - C layout is important.
>   - Compatibility with existing code.
> - Goal - Be safer than C, keep the layout.
>   - Sacrifice some safety.

some protocols use a specific layout that we can't change. we must remain compatible we legacy code.\
the c++ standard says that tier are cases when unions are good, it makes standard layour classes possible.

> - "Standard-Layout classes are usefull for communicating with code written in other programming languages"
> - various constraints [StandardLayoutType](https://en.cppreference.com/w/cpp/named_req/StandardLayoutType)
>   - no virtual functions or virtual base classes
>   - single access control - all (non static member are the same, no mixing between public/protected/private.
>   - all non-statics in the same class - all none static members live in the same class (this one or a base class).
>   - and more.
> - [std::is_standard_layout](https://en.cppreference.com/w/cpp/types/is_standard_layout)
>   ```cpp
>   static_assert(std::is_standard_layout<Citizen>::value,"not standard layout"); //c++11
>   static_assert(std::is_standard_layout_v<Citizen>); //c++11
>   ```
> - Layout-compatibility allows accessing members without knowing the type
>   - "In a standard-layout union with an active member of struct type T1, it is permitted to read a non-static data member m of another union member of struct T2 provided m is part of the common initial sequence of T1 and T2"
>   - [std::is_corresponding_member](https://en.cppreference.com/w/cpp/types/is_corresponding_member)
>   ```cpp
>   static_assert(std::is_corresponding_member(&Header::type, &Citizen::type)); //c++20
>   ```

### Intrusive variants

the a variant, but making the tag visible, and therefore layout compatible. a suggestion, using the [offsetof macro](https://en.cppreference.com/w/cpp/types/offsetof) to either create something similar to the union of struct (implicit headers)

```cpp
using ID_intrusive = intrusive_variant<
IDType, offsetoff(IDHeader, type), //where is the tag
intrusive_variant_tag_type<IDType::CITIZEN,Citizen>, //add options per tag
intrusive_variant_tag_type<IDType::TOURIST,Tourist>>; //

//will become something like this
union IdentityCard
{
    struct IDHeader {
        const IDType id;
        };
    struct Citizen{
        const IDType id;
        IDNumber nationalID;
        };
    struct Tourist{
        const IDType id;
        PassportNumber passport;
        };
};
```

a different approach will be using static constexpr tags.

```cpp
using ID_intrusive = intrusive_variant<
IDType, offsetoff(IDHeader, type), //where is the tag
intrusive_variant_type<IDType::Citizen>, //add types
intrusive_variant_type<IDType::Tourist>); //

//will become something like this
union IdentityCard
{
    struct IDHeader {
        const IDType id;
        };
    struct Citizen{
        const IDType id;
        static constexpr IDType TAG = IDType::CITIZEN;
        IDNumber id;
        } citizen;
    struct Tourist{
        const IDType id;
        static constexpr IDType TAG = IDType::TOURIST;
        PassportNumber passport;
        } tourist;
};
```

> - the user dictates the type and location of the tag.
> - visit() will still be O(1).
>   - potentially larger lookup table.
> - Customization Point for tag deduction.

we can have Different Approaches to get the Tag, even if we don't know what the type is, standard layout can have all private members, or the tag might be calculated somehow

> - Offset of the field in the object
>   - getTag<IDType>(hdr, offsetof(Hdr, tag));
>   - getTag<IDType>(hdr, std::integral_constant<size_t, offsetof(Hdr, tag)>());
> - Pointer to the field
>   - getTag<IDType>(hdr, &Hdr::tag);
> - Call a member function
> - getTag<IDType>(hdr, &Hdr::getTag);
>   - Useful when the tag is private.
>   - Useful when the tag is calculated.
> - Call a free-function / lambda.
>   - getTag<IDType>(hdr, [](const Hdr& h) { return h.tag; });
>   - Less intrusive

possible implementations, using c++20 concepts (either with reinterpret_cast or with std::invoke).

c++20 can give us extra type safety, with possible validations and without explicitly stating the offset.

```cpp
using ID_intrusive = decltype(decl_safe_intrusive_variant(
    &Citizen::type, IDType::CITIZEN,
    &Tourist::type, IDType::TOURIST));
```

intrusive variant will have a safe visit() function, but it still has place for bugs and still requires boilerplate code. class hierarchies could help us, and we will need some help from the utilities functions, but adding base classes with data breaks the standard-layout specification.

(clip from sean Parent)

trying to show a example of 'variant_of_base',base class with data and then a variant that knows to use the derived classes with better safety, but still not standard conforming.

### Streams of variants

sending arrays of variants,comparing std::variant, intrusive variant, and how it's done it the real world (send tag, then struct, so only use the ram we need, trying to minimize waste).

helper with arrays of variants, like a forward iterator, a container that like a special queue for the variants. jumping between elements.

Summary so far:

> - variants are different than unions.
> - real-world unions already have tags (and headers).
> - Intrusive_variant - C++ safety with high C compatibility.
> - Variant_of_base - add classes to your code.
>   - Not standard-layout, undefined behavior.
>   - Perhaps we should widen the rules, add \[\[standard_layout]] attribute?
> - condensed_variant - real world streams of binary data.

### Variants for de-virtualization

virtual dispatch- polymorphism.
de-virtualization - trying to get virtual function to run just as fast as non virtual functions. a talk from 2013 about creating our own virtual table because virtual functions are pretty wastefull, especially for small hireachy.
the problem with virtual functions is **Branch MisPredictions**, but compilers and processors get better with branch prediction over the years, the compiler only sees the code, but the processor sees the data so it can re-arrange the date by itself.
compilers can use PGO (profilers guided optimizations) or LTO (link time optimization), or the \[\[likely]] attribute. this allows the compiler to create code that is better suited for performance.
we want to inline, rearrange and inspect functions code, and virtual functions aren't great for that.

we might be able to use variant and std::visit() to get better de-virtualization. if we have a different implementaition of visit (without a jump table), we could get a much better performace.

</details>

## Simplest Strong Typing Instead of Language Proposal ( P0109 ) - Peter Sommerlad

<details>
<summary>
Strong types.
</summary>

[Simplest Strong Typing instead of Language Proposal](https://youtu.be/ABkxMSbejZI), [slides](https://github.com/PeterSommerlad/talks_public/raw/master/C%2B%2Bnow/2021/SimplerStrongTypes-handout.pdf), [Peter Sommerlad's Simple Strong Typing](https://github.com/PeterSommerlad/PSsst)

"Type-Errors are less likely than many other types of errors. This is why strong types are worthwhile."

strong types capture errors in compile time, so in C++ the problem don't reach the field. built-in types have dangerous implicit conversions.

motivations:

> - Order of argument bug prevention.
> - Communicate and check semantics of values.
> - Limit operations to useful subset.

### Order of Arguments

function that takes arguments of the same type,

we should use VOP - value oriented programming.

Whole Value Pattern

[CHECKS pattern language](http://c2.com/ppr/checks.html), PeterSo

> "Because bits, strings and numbers can be used to represent almost anything, any one in isolation means almost nothing".
>
> - Instead of using built-in types for domain values, provide value types wrapping them.
> - Define values types and use these as parameters
> - Provide only useful operations and functions
> - Include formatters for input and outpit
> - Do not use string or numeric representations of the same information.

strong typing is already available in all sorts of frameworks. all kinds of naming issues.\
c++17 aggregate initialization allow to avoid all sorts of issues.\
c++20 helps us with the spaceship operator<=> and constraints.

limit the useful operations:

prevent accidental expression errors.\
for example, a strong type of distance (underlying integer) can be multiplied by a factor or we can added two distances together, but we can't meaningfully multiple distance by distance and get a distance back (we can get area).

only useful operations\
this mess of code is legal:

```cpp
ASSERT_EQUAL(42.0,7. *((((true << 3) *3) % 7 )<< true)); //
```

c++ has too many operations allowed for built-in types, and they can be called via **integral promotion** and **implicit arithemetic conversions**. when we use the built-in types, we can unintentionally call on them, this is basically the issue of comparing apples to oranges. we can add apples to apples (two apples plus one apples is three apples), but adding apples and oranges doesn't have a meaning (two apples plus one orange is still two apples and an orange), nor does multiplying appels by appels.

```cpp
int appels = 2;
int oranges = 5;
auto r1 = appels * 2; //legal, makes sense. twice as many appels
auto r2 = appels +5; //legal, makes sense. five more apples
auto r3 = appels * appels; //legal, but doesn't make sense, what is apples times apples
auto r4 = appels + oranges;//legal, but doesn't make sense, adding apples to oranges
```

the P0109 paper\
Function aliases + Extended inheritance = Opaque typedefs (a.k.a strong types)

some goals that didn't make it into PSsst,

simple example of P0109, a simple opaque int type and and opaque energy strong type.

```cpp
using opaqueTypeInt = public int {
    opaqueTypeInt operator+(opaqueTypeInt o) {return opaqueTypeInt{+int{0}};}
};
opaqueTypeInt o1 = 16;
auto o2 = +o1; //o2 is of type opaqueTypeInt

using energy = protected double
{
    energy operator+ (energy, energy) = default;
    energy& operator*= (energy, double) = default;
    energy operator* (energy, energy) = delete;
    energy operator* (energy, double) = default;
    energy operator* (double, energy) = default;
};

energy e{1.23}; //ok, explicit
double d{e}; //ok, explicit
d=e; //error! protected disallows implicit type adjustments here
e=e+e; //ok, sum has type energy
e=e*e; // error, call to a deleted function
e *= 2.71828; //okay
```

visibility defines convertibility.
limiting operator defintions.
we can have opaque types recursive chain, sort of like inheritance.

peter has comments on the paper, liked it at start, but had some reservations,

### Simpest Strong Type

the simplest thing to use is struct with value.

```cpp
struct literGas{
    double value;
}
struct kmDriven
{
    double value;
}

double consumption(literGas liter, kmDriven km)
{
    return liter.value/(km.value/100);
}
```

this helps, but because of aggregate initialization, we can run into issues when a struct is wrongly initiated.

```cpp
void demonstrateStrongTypeProblem()
{
    literGas consumed{40};
    kmDriven distance{500};
    ASSERT_EQUAL(8.0,consumption(consume, distance)); //great.
    ASSERT_EQUAL(8.0,consumption({500}, {40})); //oops! braced initialization! wrong again!
}
```

but if we want to return another strong type (different struct), we need more and more boilerplate code. we write a lot of code for comparison, and then we might need more operators for i/o...

lets try to eliminate duplication. we can use generics

function template

```cpp
template <typename T>
std::ostream& operator <<(std::ostream & out, const T & val)
{
    return out << val.value;
}
```

> issues
>
> - Must be in the namespace of value classes for ADL
> - May be selected by too many classes (concept might help)
> - Assumes specific public member variables.

we can use CRTP - curiously recurring template parameters pattern.

```cpp
template <typename derived>
struct base{
//...
};
struct X : base<X>{};
struct Y : base<Y>{};
```

used for mix-ins, but we can't constrain derived because it's incomplete. c++17 allows derived class to remain an aggregate.

usage: friend functions are instantiated when used, in c++17 we can use structured binding to take the single public member function()

```cpp
template <typename OutType>
struct Out
{
    friend std::ostream& operator<<(std::ostream & out, const OutType &r)
    {
        return out<<r.value; //must be called value
    }
};

struct Literper100km: Out<literper100km> //crtp patten
{
    double value;
};

template <typename OutTypeSingleValue>
struct Out17
{
    friend std::ostream& operator<<(std::ostream & out, const OutTypeSingleValue &r)
    {
        auto const & [v]=r;  //structured binding name doesn't matter, as long as it's the only public member value and has defined output operator.
        return out<<v;
    }
};
```

Structured bindind

```cpp
auto [x,y] = f_returningStruct();
```

> - Decompose aggregates on the fly.
>   - _struct_, _std::tuple_, _std::pair_.
> - Number of names in the bracket [] = Number of data members / array elements.
> - Usually _auto_ or _const auto &_.
>   - Might get lifetime extentsion
> - _auto &_ only if function return lvalue reference.
> - not possible yet - parameter pack
>   ```cpp
>   auto [x...] =f();
>   ```

example for a crtp pattern with structured bindings and type parameters. requires naming convetions for prefix and suffix.

```cpp
template <typename OutPreSuffix>
struct OutStructured
{
    friend std::ostream& operator<<(std::ostream & out, const OutPreSuffix &r)
    {
        if constexpr(detail_::has_prefix<OutPreSuffix>{})
        {
            out << OutPreSuffix::prefix;
        }
        const auto &[v]=r;
        out <<v;
        if constexpr(detail_::has_suffix<OutPreSuffix>{})
        {
            out << OutPreSuffix::suffix;
        }
        return out;
    }
};
struct literPer100km:OutStructured<literPer100km>
{
    double value;
    constexpr static inline auto prefix ="consumption ";
    constexpr static inline auto suffix =" 1/100km";
};

void demo_output_crtp()
{
    literPer100km consumed{{},9.5};//ugly see later
    std::ostringstream out{};
    out << consumed;
    ASSERT_EQUAL("consumption 9.5 1/100km",out.str());
}
```

detection idiom

```cpp
template <typename T,typename=void>
struct has_prefix
    : std::false_type{};

template <typename T>
struct has_prefix<T,std::void_t<decltype(T::prefix)>> // actual check
    : std::true_type{};

// same for suffix

//actually does
decltype(std::declval<std::ostream&>() << T::prefix>>)
```

in c++20 we have concepts,

```cpp
template <typename T>
concept has_suffix = requires (T t){T::suffix;};
//probably needs also to check for << operator of suffix
```

we had a boiler plate comparison in out struct before.\
in cpp+17 we do with Mix-in comparison. Equality should never compare different strong types.

```cpp
template <typename T>
struct Eq{
    friend constexpr bool operator==(const T & lhs, const T & rhs) noexcept
    {
        auto const &[lhs_value] = lhs;
        auto const &[rhs_value] = rhs;
        return (lhs_value ==  rhs_value);
    }
    friend constexpr bool operator!=(const T & lhs, const T & rhs) noexcept
    {
        return !(lhs==rhs);
    };
}
```

we can also do order comparisons.in c++20 we can't default defind the spaceship operator for type T.

```cpp
template <typename T>
struct Order:Eq<T>{
    friend constexpr bool operator<(const T & lhs, const T & rhs) noexcept
    {
        auto const &[lhs_value] = lhs;
        auto const &[rhs_value] = rhs;
        return (lhs_value <  rhs_value);
    }
    friend constexpr bool operator>(const T & lhs, const T & rhs) noexcept
    {
        return  rhs<lhs;
    };
    friend constexpr bool operator<=(const T & lhs, const T & rhs) noexcept
    {
        return !(rhs<lhs);
    };
}
```

for arithmetics, example of adding template of friend functions, this works for some binary operations, but not all.

```cpp
template <typename R>
struct Add<R>{
    friend constexpr R& operator<(R & lhs, const R & rhs) noexcept
    {
        auto &[lhs_value] = lhs; //lvalue captured binding
        auto const &[rhs_value] = rhs;
        lhs_value += rhs_value;
        return lhs;
    }
    friend constexpr R operator+(R & lhs, const R & rhs) noexcept
    {
        return lhs+=rhs;
    };
}
```

> "We can not contrain the mix-in class parameter"\
> "When instantiated, argument type is still incomplete `struct energy : add<energy>`"

crtp example for scalar multiplication, uses the type and a SCALAR type parameter, we need some tricks.

```cpp
template <typename R, typename Scalar>
struct ScalarMultiImpl<R>{
    friend constexpr R& operator*=*(R & lhs, const Scalar & rhs) noexcept
    {
        auto &[lhs_value] = lhs; //lvalue captured binding
        lhs_value *= rhs;
        return lhs;
    }
    friend constexpr R operator*(R & lhs, const Scalar & rhs) noexcept
    {
        return lhs*=rhs;
    };
        friend constexpr R operator*(const Scalar & lhs, R & rhs) noexcept
    {
        return rhs*=lhs;
    };
}
```

and lets see it in action. we still have too many {} curly braces in our code, this is because of all of our base classes that we pushed

```cpp
struct kmDriven:Out<kmDriven>,ScalarMultiImpl<kmDriven,double>
{
    double km;
};
lieterPer100km consumption (lieterGas l, kmDriven km)
{
    return {{},{},l.liter/(km*0.01).km};
}

void demonstrante()
{
    lieterGas l{{},40};
    kmDriven km {{},{},500};
    ASSERT_EQUAL(lieterPer100km({},{},8.0),consumption(l,km));
}
```

do stuff with bit operator, only work for unsigned, we use _static_assert_. for shift operator,we have the type and the shifting number type, which might be a built-in type or a strong type. we can added another assert to check if we don't shift too many bits.

### Making Mix-ins Work for strong types

how do we remove the curly braces for all the base class of the CRTP? what about useful operator combinations?

combining mix-ins Bases. using template parameter template packs.

> - takes a strong type "derived class" T
> - takes a list of mix-in bases class templates
> - instantiates all bases with T

```cpp

template <typename T, template <typename ...> class ...CRTP>
struct ops: CRTP<T>...{};
template <typename T>
using Additive =ops<T,TPlus,TMinus, Abs,Add,Sub, Inc,Dec>; // all sorts of structs we created earlier

//usage
struct liter :ops<liter,Additive, Order, Out> // also stuff, works recursive
{
    double l{};
};
```

we can define an explist constructor to remove the curly braces. this will prevent implicitly conversions on return type

```cpp
struct literGs: ops<literGas, Additive, Order, Out>
{
    constexpr explicit litergas(double lit):l(lit){

    }
    double l{}
}
```

other versions, put data member in the first base class object and the rest of the base classes get elided.

```cpp
template <typename V,typename TAG>
struct holder {
    static_assert(std::is_object_v<V>), "no reference or incomplete types allowed!");
    using value_type= V;
    V value{};
};
struct literGas : holder<double,literGas>, ops<literGas, Additive, Order, Out>{}; //no need to define constructor or type
```

we can get another level of indirection, put the data member first, and then combine with ops. variable template template pack.

```cpp
template <typename V, typenam Tag, template<typename...>class OPS>
struct strong : detail_::holder<V,Tag>,ops<Tag,OPS...>
{};

//usage
struct literGas: strong<double,literGas,Additive, Order, Out>
{};
```

skipped slide for "Different Inits for return",slide "Trait for determinging init possibility"

uas a mapping of mathematical functions for strong types. like absolute value of rounding operations. relies ob macros.

but to scalar multiplication, module operations (different base classes and SFINAE), remember that `std::is_integral_v(bool)` is true. some how preventing repearing scalar types. template aliases with template members, so there is more trickery involved.

we have different versions of this library for different c++ standards.

### Linear spaces

we use same type for a _vector space_ as well as for _affine space_. its a mixture of domains.

| Vector Space    | Affine Space      |
| --------------- | ----------------- |
| displacement    | position          |
| no origin       | definitive origin |
| ptrdiff_t       | size_t            |
| difference_type | size_type         |

the chrono library does this right. the vector space is represented by 'duration' and the affine space is 'time_point'.

time_point - time_point -> duration
time_point + duration -> time_point
time_point + time_point -> error 🐄💩

we can affine spaces that are related, like celsius and kelvin, which are the same with different origin points.

### Summary

> - P0109 was a good attempt, but failed.
> - A library solution allows simpler strong typing.
> - Naming the domain type allows for nicer mangling.

units library doesn't solve the same problems that strong types do
slides continue with more stuff..

</details>

## Simplest Safe Integers - Peter Sommerlad

<details>
<summary>
Implementing a safe integer type.
</summary>

[Simplest Safe Integers - Peter Sommerlad](https://youtu.be/Z0X_TFCcTXA),[slides](https://github.com/PeterSommerlad/talks_public/raw/master/C%2B%2Bnow/2021/SimplestSafeIntegers.pdf), [github code](https://github.com/PeterSommerlad/PSsimplesafeint).

starting with an example of undefined integer behavior
[is_modulo](https://en.cppreference.com/w/cpp/types/numeric_limits/is_modulo)

```cpp
std::cout << 65535 * 32768 << '\n';
// print 2147450880
std::cout << 65536 * 32768 << '\n';
// print ?
std::cout << std::numeric_limits<int>::is_modulo << '\n';
// print false
```

if we use unsigned integers

```cpp
std::cout << 65535u * 32768u << '\n';
// print 2147450880
std::cout << 65536u * 32768u << '\n';
// print 2147483648
std::cout << 65536u * 32768u *2u << '\n';
// print ?
std::cout << std::numeric_limits<unsigned>::is_modulo << '\n';
// print true
```

we have some domains where undefined behavior is critical, so we want to avoid it. there is the **MISRA-C++** standard with guidelines to avoid UB, one of the recommendations is to avoid impelemnation defiend types,to use bit operations only for with unsigend types, and to have a tool that enforces them.

goals:\
no mixing of integers with characters types (char, char8_t, char16_t char32_t), no promoting and no changing of sign, no signed integer overflow UB, no silent conversions (except for widening of the same sign).\
another goal is to have the exact same code generation.

Enums of integer replacements:\
limitations:

> - cannot have implicit conversions
> - needs name conversion function for checks
> - division by zero, bit shifting (undefined behavior)
>   - throws exception (macro controlled)
>   - asserts (no exceptions)
>   - returns 0 (NDEBUG)

`enum` for wrapping integers:

> - enum class types (`std::byte` does it as well)
> - operator overloading
> - concepts constrained operators
> - `consteval` user-defined-literal operators

### Defining the class

so we can have non-promoting integers:

```cpp
// unsigned
enum class ui8 : std::uint8_t { tag_to_prevent_mixing_other_enums };
enum class ui16: std::uint16_t{ tag_to_prevent_mixing_other_enums };
enum class ui32: std::uint32_t{ tag_to_prevent_mixing_other_enums };
enum class ui64: std::uint64_t{ tag_to_prevent_mixing_other_enums };
// signed
enum class si8 : std::int8_t { tag_to_prevent_mixing_other_enums };
enum class si16: std::int16_t{ tag_to_prevent_mixing_other_enums };
enum class si32: std::int32_t{ tag_to_prevent_mixing_other_enums };
enum class si64: std::int64_t{ tag_to_prevent_mixing_other_enums };
```

user defined literals (UDL), creating a user defined literal suffix, which checks at compile time.

```cpp
inline namespace literals {

// none-standard user defined literal
// used in compile time
consteval ui16 operator""_ui16(unsigned long long val)
{
    if (val <= std::numeric_limits<std::underlying_type_t<ui16>>::max())
    {
        return ui16(val);
    }
    else
    {
        throw "integral constant too large"; // trigger compile-time error
    }
}
// etc...
```

usage, detecting errors, checking equality

```cpp
using namespace psssint::literals; // required for UDLs
void ui16intExists()
{
    using psssint::ui16;
    auto large=0xff00_ui16;
    //0x10000_ui16; // compile error
    //ui16{0xfffff}; // narrowing detection
    ASSERT_EQUAL(ui16{0xff00u},large);
}
```

traits and concepts: `is_` for traits, `a_` and `an_` for concepts. we have a concept to detect that our type is an enum, and then check that we can't convert to the underlying type (scoped enum), and then we define that as another concept. we can use traits even without c++20.

```cpp
template<typename T>
using plain = std::remove_cvref_t<T>; // just a shorthand

template<typename T>
concept an_enum = std::is_enum_v<plain<T>>;

// from C++23
template<an_enum T>
constexpr bool is_scoped_enum_v = !std::is_convertible_v<T,std::underlying_type_t<T>>;

template<typename T>
concept a_scoped_enum = is_scoped_enum_v<T>;
```

the "Detection Idiom" with concepts, we define a trait `is_safeint_v` that is false, and specialize it on the earlier concept of `a_scoped_enum` with a and a _requires_ clause to check for our special tag. then we create another concept that defines a safe integer type which is a scoped enum with this tag.

```cpp
template<typename T>
constexpr bool is_safeint_v = false;

template<a_scoped_enum E>
constexpr bool is_safeint_v<E> = requires
{
     E{} == E::tag_to_prevent_mixing_other_enums;
};

template<typename E>
concept a_safeint = is_safeint_v<E>;
```

testing detection idiom

```cpp
namespace _testing {
using namespace psssint;
//positive tests
static_assert(is_safeint_v<ui8>);
// more asserts on the types we created

// negative tests
enum class enum4test{};
static_assert(!is_safeint_v<enum4test>); // enum is not a safe integer,
static_assert(!is_safeint_v<std::byte>); // nor is a std::byte
static_assert(!detail_::is_safeint_v<int>); // nor is int
```

we started with an enum class that inherits from a primitive type and has a special enum value `tag_to_prevent_mixing_other_enums`, then we created a concept `a_scoped_enum` that this enum cannot be converted back to the underlying type, and then another concept `a_safeint` that a scoped enum must have this predefined tag.

### Promotions and Conversions

meta programming for promotion. we allow integer promotion, but not sign change. we have a way to do this.

```cpp
template<typename E> // ULT = underlying type
using ULT = std::conditional_t<std::is_enum_v<plain<E>>,
 std::underlying_type_t<plain<E>>, plain<E>>; //

template<typename E>
using promoted_t = // will promote keeping signedness
 std::conditional_t<
 (sizeof(ULT<E>) < sizeof(int)) // test if promotion is needed
 , std::conditional_t<std::is_unsigned_v<ULT<E>> , unsigned , int > // return the same sign
 , ULT<E> //no promotion
 >;

template<a_safeint E>
constexpr auto to_int(E val) noexcept
{
  return static_cast<promoted_t<E>>(val); // promote keeping signedness
}
```

testing sign promotions, are thing the same type and sign?

```cpp
static_assert(std::is_same_v<unsigned,decltype(to_int(1_ui8)+1)>);
static_assert(std::is_same_v<unsigned,decltype(to_int(2_ui16)+1)>);
static_assert(std::is_same_v<int,decltype(to_int(1_si8))>);
static_assert(std::is_same_v<int,decltype(to_int(2_si16))>)
```

this is bad, a sigend char is actually `int8_t`. but long is weird. the promotion flow should be:

> char < short <int < long < long long

```cpp
template<std::integral T>
constexpr bool is_integer_v =
    std::is_same_v<uint8_t , T> ||
    std::is_same_v<uint16_t, T> ||
    std::is_same_v<uint32_t, T> ||
    std::is_same_v<uint64_t, T> ||
    std::is_same_v<int8_t , T> ||
    std::is_same_v<int16_t , T> ||
    std::is_same_v<int32_t , T> ||
    std::is_same_v<int64_t , T>;

template<typename T>
concept an_integer = is_integer_v<T>;
```

so lets find something better, we check for the integral type, remove char and bool, check the sign to be the same and the range to be the same as that of int.

```cpp
template<typename INT, typename TESTED>
constexpr bool is_compatible_integer_v =
    std::is_same_v<TESTED,INT> || // either the same type
    (
        std::is_integral_v<TESTED> && //
        !std::is_same_v<bool,TESTED> &&// exclude bool
        !is_chartype_v<TESTED> && // exclude characters
        (std::is_unsigned_v<INT> == std::is_unsigned_v<TESTED>) &&
        std::numeric_limits<TESTED>::max() == std::numeric_limits<INT>::max()
    );

// concept
template<typename INT, typename TESTED>
constexpr bool is_similar_v=is_compatible_integer_v<INT,TESTED>;
```

character types are not intgers, the MISRA-C++ excludes them. so lets check againts them (also wide character )

```cpp
template<typename CHAR>
constexpr bool is_chartype_v =
    std::is_same_v<char, CHAR>
 || std::is_same_v<wchar_t, CHAR>
#ifdef __cpp_char8_t // feature test to check if type exists
 || std::is_same_v<char8_t, CHAR> //c++20 addition
#endif
 || std::is_same_v<char16_t,CHAR>
 || std::is_same_v<char32_t,CHAR> ;
```

so now our check looks like this

```cpp
template<typename TESTED>
constexpr bool
is_known_integer_v = is_similar_v<std::uint8_t, TESTED>
 || is_similar_v<std::uint16_t, TESTED>
 || is_similar_v<std::uint32_t, TESTED>
 || is_similar_v<std::uint64_t, TESTED>
 || is_similar_v<std::int8_t, TESTED>
 || is_similar_v<std::int16_t, TESTED>
 || is_similar_v<std::int32_t, TESTED>
 || is_similar_v<std::int64_t, TESTED>;
// deliberately not std::integral, because of bool and characters!

template<typename T>
concept an_integer = detail_::is_known_integer_v<T>;
```

and the conversion is

```cpp
template<an_integer T>
constexpr auto from_int(T val)
{
    using detail_::is_similar_v;
    using std::conditional_t;
    struct cannot_convert_integer{}; //type for unconvertible values
    using result_t = //this is basically a huge if-else statements that returns the correct safe type.
        conditional_t<is_similar_v<std::uint8_t,T>, ui8,
        conditional_t<is_similar_v<std::uint16_t,T>, ui16,
        conditional_t<is_similar_v<std::uint32_t,T>, ui32,
        conditional_t<is_similar_v<std::uint64_t,T>, ui64,
        conditional_t<is_similar_v<std::int8_t,T>, si8,
        conditional_t<is_similar_v<std::int16_t,T>, si16,
        conditional_t<is_similar_v<std::int32_t,T>, si32,
        conditional_t<is_similar_v<std::int64_t,T>, si64,
        cannot_convert_integer>>>>>>>>;
        return static_cast<result_t>(val);
}
```

and a way to directly convert from an integer type to an safe integer type. we check the value is larger than zero for unsigned safe types, and that it fits inside the range of the requested type, and for signed types, we check the min and max ranges.

```cpp
template<a_safeint TO, an_integer FROM>
constexpr auto from_int_to(FROM val)
{
    using result_t = TO;
    using ultr = std::underlying_type_t<result_t>;
    if constexpr(std::is_unsigned_v<ultr>)
    {
        if ( val >= FROM{} && val <= std::numeric_limits<ultr>::max())
        {
            return static_cast<result_t>(val);
        }
        else
        {
            throw "integral out of range";
        }
    }
    else
    {
        if (val <= std::numeric_limits<ultr>::max() &&
        val >= std::numeric_limits<ultr>::min())
        {
            return static_cast<result_t>(val);
        }
        else
        {
            throw "integral out of range";
        }
    }
}
```

testing this explicit conversion, and checking the compile time behavior (that we can convert from the default value) by making sure an expression is a valid behavior. there is still an issue with the equal sign after `>>`.
`>> =` is ok, `>>=` is still _shift equal_.

```cpp
static_assert(1_ui8 == from_int(uint8_t(1)));
static_assert(42_si8 == from_int_to<si8>(42));
//static_assert(32_ui8 == from_int(' ')); // does not compile
//static_assert(1_ui8 == from_int_to<ui8>(true)); // does not compile

void checkedFromInt()
{
    using namespace psssint;
    ASSERT_THROWS(from_int_to<ui8>(2400u), char const *); // this should throw, we 2400u doesn't fit in eight bits
}

template<typename FROM, typename=void>
constexpr bool from_int_compiles=false;

template<typename FROM>
constexpr bool from_int_compiles<FROM, std::void_t<decltype(psssint::from_int(FROM{}))>> = true;
//...

static_assert(from_int_compiles<long>); // detect long bug
static_assert(! from_int_compiles<bool>); // can't convert from bool (false)
static_assert(! from_int_compiles<char>); // can't convert from char (0)
```

### Operators

output operator `<<`. concept `a_safeint` prevents using it for other types, we can use either form to define the operator, this is a matter of style preference.

```cpp
//form 1
template<a_safeint E>
std::ostream& operator<<(std::ostream &out, E value){
    out << to_int(value);
    return out;
}

//form 2
std::ostream& operator<<(std::ostream &out, a_safeint auto value)
{
    out << to_int(value);
     return out;
} // concept short-hand notation
```

arithmetic operator example, for addition, we combine to safe integers, which must be of the same sign. we allow for promotion (to the wider type), and before the operator itself we convert to unsigned, so we can get wrong results (wrapping), but not overflow

```cpp
template<a_safeint E, a_safeint F>
constexpr auto operator+(E l, F r) noexcept
requires same_sign<E,F>
{
    using result_t=std::conditional_t<sizeof(E)>=sizeof(F),E,F>;
    return static_cast<result_t>(
        static_cast<ULT<result_t>>(
            to_uint(l)
            + // use unsigned op to prevent signed overflow, but wrap.
            to_uint(r))
        );
}
```

an example of code generation...

```cpp
template<typename INT>
struct operations {
operations(std::initializer_list<INT> seedvalues):values{seedvalues}{};
    std::vector<INT> values;
    INT sum() const
    {
        return std::accumulate(begin(values),end(values),INT());
    }
};
std::initializer_list<int8_t> i8_seed{1,1,2,3,5,8,13,21,34,55,89};

std::initializer_list<psssint::si8>
si8_seed{1_si8,1_si8,2_si8,3_si8,5_si8,8_si8,13_si8,21_si8,34_si8,55_si8,89_si8};

auto sum(operations<int8_t> const &ops)
{
    return ops.sum();
}

auto sum(operations<si8> const &ops)
{
    return ops.sum();
}
```

### Supressing All undefined behavior

some operations might fail, division by zero, modulo by zero, shifting too many bytes. in some cases, we can throw exceptions, but some domains disallow exceptions. a known _wrong result_ is better than undefined behavior.

this macro throws on debug, and returns a know wrong result (default value) in regular mode. if we can detect the error at compile time, we can fail the exception. if we are in release mode and the condition fails, we return the default value.

```cpp

#ifdef NDEBUG /* case 3 - release mode*/
    #define ps_assert(default_value, cond) \
    if (std::is_constant_evaluated()) {\
    if (not (cond)) throw(#cond); /* compile error */\
    } else {\
    if (not (cond) ) return (default_value);/* last resort avoid UB */\
    }
    #define NOEXCEPT_WITH_THROWING_ASSERTS noexcept
#else /* debug mode */
    #ifdef PS_ASSERT_THROWS /* case 1 - we allow throwing exceptions*/
        #define ps_assert(default_value, cond) ((cond)?true: throw(#cond))
        #define NOEXCEPT_WITH_THROWING_ASSERTS noexcept(false)
    #else /* case 2 crush the program with asset */
        #include <cassert>
        #define ps_assert(default_value, cond) assert(cond)
        #define NOEXCEPT_WITH_THROWING_ASSERTS noexcept
    #endif
#endif
```

and now actually using this, return zero (deafult value of type) when dividing by zero. we check if we can divide by zero.

```cpp
template<a_safeint E, a_safeint F>
constexpr auto operator/(E l, F r) NOEXCEPT_WITH_THROWING_ASSERTS // <--
requires same_signedness<E,F>
{
    using result_t=std::conditional_t<sizeof(E)>=sizeof(F),E,F>;
    ps_assert(result_t{}, r != F{} && " division by zero"); // <--

    return static_cast<result_t>(
        static_cast<detail_::ULT<result_t>>(
            to_uint(l)
            / // use unsigned op to prevent signed overflow, but wrap.
            to_uint(r)
        )
    );
}
```

and now testing the assertion handling, again with macros, and traits, and templates... we turn all the values into void
(not writing the code)

now, we can go back to the conversion function and make it better, we get the same behavior (throw, crush, return some value) as before.

```cpp
template<a_safeint TO, an_integer FROM>
constexpr auto from_int_to(FROM val) NOEXCEPT_WITH_THROWING_ASSERTS
{
    using result_t = TO;
    using ultr = std::underlying_type_t<result_t>;
    if constexpr(std::is_unsigned_v<ultr>)
    {
        ps_assert( result_t{}, (val >= FROM{} && val <= std::numeric_limits<ultr>::max()));
        return static_cast<result_t>(val);
    }
    else
    {
        ps_assert( result_t{}, (val <= std::numeric_limits<ultr>::max() &&
        val >= std::numeric_limits<ultr>::min()));
        return static_cast<result_t>(val);
    }
}
```

there is a bug with gcc compiler that gives us a false alarm so we might need a pragma to disable this warning for this code block.

### Extra stuff

we would also like our types to have limits, so we specialize over the `numeric_limits` struct,

```cpp
template<psssint::a_safeint type> // constrained
struct numeric_limits<type>
{
    using ult = psssint::detail_::ULT<type>;
    static constexpr bool is_specialized = true;
    static constexpr type min() noexcept
    { return type{numeric_limits<ult>::min()}; }
    static constexpr type max() noexcept
    { return type{numeric_limits<ult>::max()}; }

    // ...
    static constexpr int digits = numeric_limits<ult>::digits;
    static constexpr int digits10 = numeric_limits<ult>::digits10;

    //...
    static constexpr bool is_modulo = true; // even for signed
    static constexpr bool traps = false; // never signals
};
```

theres also a c++17 branch without concepts, uses `std::enable_if` and SFINAE and not concepts, and also something else for `std::numeric_limits<>`.

the alternatives are _boost/safe_numerics_ or _CNL: Compositional Numeric Library_ (also fixed points arithemtic).

Q&A

> - _why not pass `modulo` to the base numeric limit?_. because he wanted signed numerics to wrap around.
> - _matching to a standard behavior of numeric types_
> - _using compiler flags to force wrapping_

</details>

## Techniques for Overloading any_invocable - Filipp Gelman

<details>
<summary>
A template type that's like std::function, but with overloads
</summary>

[Techniques for Overloading any_invocable](https://youtu.be/JnXpGA7SYHQ), [Slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/tfoai.pdf)

will probably a template heavy lecture

outline

### What is overloaded _any_invocable_

_any_invocable_ is a type erasing container for function objects, similar to _std::function_, but it has move only semantics, and support overloads of operator().

an example of an async interface

- using a function pointer, if error code is zero, resp is the response recivied, if it's not zero (error), then respone is NULL. the data is opaque.
- using std::function, encapsulates the type, but the function object is copied, and if we have a non copyable resource in the lambda, it fails.
- using _any_invocable_, not requires copyability

```cpp
int sendAsyncFunctionPointer(Request const & request,
             void (*on_response)(int error_code,Response const resp*, void* data),
             void* data);

int sendAsyncSTDFunction(Request const & request,
            std::function<void(int error_code, Response const resp*)> on_response);

int sendAsyncAnyInvocable(Request const & request,
            any_invocable<void(int error_code, Response const resp*)> on_response);
int main()
{
    Request request;
    auto callback = [resource= std::make_unique<Resource>()](int error_code, Response const * resp){/*...code*/};
    sendAsyncSTDFunction(request,std::move(callback)); //fails! tries to verify that copy operations exist, and fails.
}
```

but we still need to handle both cases (request success or fail) in the same function, won't it be easier to have separate this logic? like having two overloads, one for error_code (when not zero) and one for the response (but the can still share state a private data, so not entirely separate). so now we can use a reference to the respone rather than a pointer, because we will never call it with null data.

```cpp
int sendAsyncAnyInvocable(Request const & request,
            any_invocable<void(int error_code), void(Response const &resp)> on_response);

struct Callback{
    void operator()(int error_code)
    {
        //handle error
    }
    void operator()(Response const & resp)
    {
        //process response
    }
};
any_invocable<void(int),void(Response const &)> f = Callback{};
```

this also allows us to use _std::overload_ (not part of the standart yet) to create a single object from many callable objects and produces a composed objects without type erasing.

### Building _any_invocable_

a simple _any_invocable_, it has a converting constructor and call operator.

```cpp
template <typename> class any_invocable;

template <typename RET, typename... ARGS>
class any_invocable<RET(ARGS ...)>
{
    //...
public:
    //special member functions

    template </**/>
    any_invocable(T object);
    RET operator()(ARGS ...);
};
```

one way to build is to use virtual inheritance, unique ptr, forwarding.

```cpp
template <typename RET, typename... ARGS>
class any_invocable<RET(ARGS ...)>
{
    struct base{
        virtual RET operator()(ARGS&& ...) = 0;
        virtual ~base() = default;
    };

    template<typename T>
    struct derived : base
    {
        T object;

        template <typename... CARGS> //arguments for creation of the object
        derived (CARGS&& ... cArgs) : object(std::forward<CARGS>(cArgs)...){} //constructor, forward constructor

        RET operator()(ARGS&& ... args) override
        {
            return std::invoke(object, std::forward<ARGS>(args)...); //calling the callable function of object
        }
    };
    std::unique_ptr<base> ptr_; //always points to a derived class
public:
    //...
    any_invocable() noexcept= default;
    any_invocable(any_invocable&& ) noexcept= default; //move constructor
    any_invocable& operator=(any_invocable&&) noexcept= default; //move assignment
    ~any_invocable() default; //destructor

    //converting constructor
    template <typename T>
    any_invocable(T&& object):ptr_(std::make_unique<derived<std::decay_t<T>>>)(std::forward<T>(object)){}
    //operator()
    RET operator()(ARGS ...)
    {
        retrun (*ptr)(std::forward<ARGS>(args)...);
    }
};
```

it's important to constrain the constructor. so we we should concepts

this example fails without constrains

```cpp
void test(std::string);
void test(any_invocable);
void call_test()
{
    test("Hi"); //not a std::string, a string literal, fails because it can also be any_invocable.
}
```

but once we use concepts, the converting function no longer applies for string literals, because it's not invocable.

```cpp
template <typename RET, typename... ARGS>
class any_invocable<RET(ARGS ...)>
{
    //all the stuff from before
    public:
    //converting constructor with constraints
    template <typename T>
    any_invocable(T&& object)
        requires std::is_invocable_r_v<RET,std::decay_t<T>&, ARGS&&...>
    :ptr_(std::make_unique<derived<std::decay_t<T>>>)(std::forward<T>(object)){}
}
```

### Adding more overloads

we want more than one overload, not just for type of arguments, but for a whole bunch of signatures, if we had refelection this would be easy.

```cpp
struct base{
    virtual RET1 operator()(ARGS1&&...) =0;
    virtual RET2 operator()(ARGS2&&...) =0;
    virtual RET3 operator()(ARGS3&&...) =0;
    //...
};
```

we can still build it, just in a different way, polymorphism has many forms. rather than relay on the built-in virtual table, we can recreate it using templates.

```cpp
struct vtable {
    void (&destroy)(base&);
    RET1 (&invoke1)(base&, ARGS1...);
    RET2 (&invoke2)(base&, ARGS2...);
    RET3 (&invoke3)(base&, ARGS3...);
};
```

we start by decomposing the vtable into table entries, each entry is a template, including the destructor.

there is a bug here:

```cpp
template <typename> struct vtable_entry;

template <typename RET, typename ... ARGS>
struct vtable_entry<RET(ARGS...)>
{
    using fun_t = RET (&)(BASE&, ARGS&& ...);
    fun_t invoke;
};
struct vtable_dtor{
    using fun_t = void (&)(Base&) noexcept;
    fun_t destroy;
};

//struct vtable: vtable_dtor, vtable_entry<FNS...> //the bug is here?
struct vtable: vtable_dtor, vtable_entry<FNS>... //un bugged?
{
    constexpr explicit vtable(vtable_dtor::fun_t dtor,
        typename vtable_entry<FNS>::fun_t ... invoke) noexcept :
        vtable_dtor(dtor), vtable_entry<FNS>{invoke}...
        {}
};
```

now, lets look again at _any_invocable_, now the base isn't using the built virtual inheritance, it directly stores the pointer reference. for every function in the vtable, the derived class has to implement a static method,

```cpp
template <typename... FNS>
class any_invocable
{
    // template struct vtable_entry
    // struct vtable_dtor
    // struct vtable

    struct base {
        vtable const & vtable_;
    };

    template <typename T>
    struct derived : base {
        T object;

        static void destroy(base&) noexcept
        {
            delete &static_cast<derived&>(base);
        }
        template <typename RET, typename... ARGS>
        static RET invoke(struct base& base, ARGS... args) //shouldn't it be &&?
        {
            return std::invoke(static_cast<derived&>(base).object, std::forward<ARGS>(args)...);
        }
        static inline constexpr struct vtable const vtable{
            derived::destory,
            static_cast<typename vtable_entry<FNS>::fun_t>(derived::invoke)...
        };

        //constructor
        template <typename ... CARGS>
        derived(CARGS&& ..cArgs): base{vtable},object(std::forward<CARGS>(cArgs)...){}
    };

    struct base* ptr_; //no longe a unique ptr
    public:
    //... constructors can no longer be defaulted
    any_invocable() noexcept: ptr_(nullptr){}
    any_invocable(any_invocable&& other ) noexcept : ptr_(std::exchange(other.ptr_,nullptr){}; //move constructor
    any_invocable& operator=(any_invocable&& other) noexcept
    {
        any_invocable(std::move(other)).swap(*this); //move swap idiom
        return *this;
    }
    ~any_invocable(){
        if(ptr_)
        {
            ptr_->vtable_->destory(*ptr); //
        }
    } //destructor

    //converting constructor
    template <typename T>
    any_invocable(T&& object):ptr_(std::make_unique<derived<std::decay_t<T>>>)(std::forward<T>(object)){}
    //operator()
};
```

we still need all the operators, we would want each of them to be something like this, and all of them inside the any_invocable

```cpp
RET1 operator()(ARGS1...args)
{
    //get the vtable
    vtable const & vt = _ptr->vtable;
    //get the function pointer
    RET1 (& invokeOverload)(base&, ARGS1...) = vt.vtable_entry<RET1(ARGS1...)>::invoke;
    //call with base
    return invokeOverload(*ptr_, std::forward<ARGS1>(args)...);
}
```

we can make an interface from this

```cpp
template <typename RET, typename... ARGS>
struct invocable_interface<RET(ARGS...)>
{
    RET operator()(ARGS... args)
    {
        return ptr_->vtable::vtable_entry<RET(ARGS,,,)>::invoke(*ptr_, std::forward<ARGS>(args)...);
    }
};
```

but how can the interface access the pointer that is inside any_invocable? we can use the curiously recursive template pattern!
it can downcast itself to get the pointer.

```cpp
template <typename,typename>
struct invocable_interface;

template <typename RET, typename... ARGS,typename ...FNS>
struct invocable_interface<RET(ARGS...),and_invocable<FNS...>>
{
    RET operator()(ARGS... args)
    {
        any_invocable<FNS...>& self = static_cast<any_invocable<FNS...>&>(*this);
        return self.ptr_->vtable::vtable_entry<RET(ARGS,,,)>::invoke(*(self.ptr_), std::forward<ARGS>(args)...);
    }
};
```

so now _any_invocable_ needs to inherit from this CRTP interface. and we also need to expose the ptr\_ as a friend, and we pull up the operator with the variadic using template alias.

also we need to check if it's invocable, it can't be a concept, because reasons. so we have some tempate specialization before that.

```cpp
template <typename... FNS>
class any_invocable: invocable_interface<FNS,any_invocable<FNS...>>...
{
    //...
    template<typename, typename>
    friend struct invocable_interface;
public:
    //...
    using invocable_interface<FNS, any_invocable<FNS...>>::operator()...;

    template<typename, typename>
    inline constexpr bool is_invocable_v = false;

    //template specialization
    template<typename T, typename RET, typename... ARGS>
    inline constexpr bool is_invocable_v<T,RET(ARGS...)> =
        std::is_invocable_r_v<RET,T, ARGS&& ...>;

        //converting constructor with constraints
    template <typename T>
    any_invocable(T&& object)
        requires (std::is_invocable_v<std::decay_t<T>&, FNS> &&...)
    :ptr_(new derived<std::decay_t<T>>>)(std::forward<T>(object)){}
};
```

we can still do better! we have a statically allocated vtable, it always goes on the heap, maybe it can be allocated on the stack for small object optimizations? we can have buffer somewhere with all of those functions, lambda, objects and other stuff. this buffer can be on the stack or on the heap.

(not writing this, some code of the buffer and changes in the other classes to accommodate).

trivially relocatable proposal (P1144). if it was added to the standard, stuff would be easier.

### Disassembly

examples taken from compiler explorer. looking at the assembly code

- comparing std::function and any_invocable.
- constructors costs, depends on the triviality of the overloads and their size.

### Benchmarks

comapring std::function with any_invocable of different overloads. compile-time and object file size and runtime. compilation time and file-size were linear with the number overloads, but runtime was constant(and faster than std::function,maybe).

### Alternative Implementions

implementations could be different,

> Function storage
>
> - Point to static struct with references to functions
> - In-place pointers to non-member functions
> - In-place pointers to member functions
>
>   Argument forwarding
>
> - forwarding references
> - forwarding values

different function storage could mean that some operations are easier and encapsulated into a storage class.
all sorts of trade-offs.

</details>

## Writing a C++20 Module - Steve Downey

<details>
<summary>
Modules in C++20, what exists, what can we do, what are the problems?
</summary>

[Writing a C++20 Module](https://youtu.be/AO4piAqV9mg),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/modulating-a-component-slides.pdf)

creating a c++20 module interface, implementing a simple data structure (functional tree),exporting types, inline code, hiding implementation, making sure that necessary un-exported defintions are still reachable.

### C++20 Modules Overview

Modules are hygiene. Modules are not packages, they don't solve the problem of packaging, they add more to the problem.

> terminology
>
> - Module Unit - a TU that contains a module declaration.
> - Named Module - the collection of Module Units module name.
> - Module Interface Unit (**MIU**) - a module unit that **exports**.
> - Module Implementation Unit - a module unit that **does not** export.
> - Primary Module Interface Unit (**PIMU**) - there will exactly one MIU that is not a partition.
> - Module Partition - Part of the module, MIU partitions must be exported by the PMIU.

example (from the IS). each translation unit should be a different file

```cpp
//Translation unit #1: PMI
export module A; //export module A
export import :Foo; // import and then export
export int baz(); // export declaration of function.

//Translation unit #2: Partion A:Foo
export module A:Foo //export module A:Foo
import :Internals; // import internals
export int foo() {return 2*(bar() +1);} //export declertion and defintion

//Translation unit #3: Partition A:Internals
module A:Internals; //declare self as module A:Internals
int bar(); // function declaration

//Translation unit #4: an implementation unit
module A; //declare self as module A
import :Internals; //import Internals
int bar(){return baz()-10;} //definition
int baz(){return 30;} //defintion
```

we have control over what we want to see, we can export declerations, defintions, we can forward exports (import and export).

> The models is Retrofiting existing tech
>
> - The standard is complicated because it is trying not to describe and implementation.
> - a module interface TU produces an object file and BMI.
> - a module TU is a TU and produces an object file.
> - the consumer of a module reads the BMI.
> - The program links the library or objects from the module.

> "**Export**s make names from the module available to the consumers"

we can still get the un-exported types, we can use decltyp or such.

> "**Import**s make names from the module visible in the current TU"

this code makes module M's exported names visible to the importer (consumer) of the current module.

```cpp
export import M; //import and re-export
```

Private module fragments (**PMF**) \
similar to java style single file modules.\
when we write this in the Primary Module Interface Unit (PMUI), the names and definitions thereafter are not reachable from the importers.

```cpp
//reachable
module: private
//unreachable
```

> "**Instation context **- how we figure out what declarations are in play for ADL and which are reachable"

reachability isn't the same as name availability, things can be reachable even when not exported.

> "Whether a declarations is exported has no bearing on whether it is reachable."
>
> **A tranlation unit U is reachable from P**
>
> - if the unit P is in has an interface dependency on U.
> - if the unit P is in imports U.
> - other unspecified reasons we should not depend on.
>
> **A declartion is reachable from P**
>
> - if it appears before P in the same TU.
> - if it's not discarded, is in a unit reachable from P, not in a PMF (private module fragment).
>
> "The things we export make more things reachable"\
> The consumer can use what we export, we don't have to export everything.

in this example, we can use Y, which is internally X, because X is reachable, however, we can't directly initialize X, because it's not visible (not exported).

```cpp
//translation unit #1:
export module A; //declare export module
struct X{};
export using Y= X; //exporting alias

//translation unit #2:
module B;
import A;
Y y; //ok,Y is exported, defintion of X is reachable
X x; //error, X is not visible to unqualified lookup.
```

**Reachability is ABI.**

export is making the names available for the programmer writing source code.

### Moduleing example - before Changing.

next, we will modulate a component [fringeTree](https://github.com/steve-downey/fringetree), which was designed as an example of the [same fringe problem](http://nsl.com/papers/samefringe.htm). this is an intentionally poor implementations of a tree that is non-mutable and produces a new tree at updates(functional programming). the real implementation is called 'fingertree', which is more complicated. fringe trees store all the data in the leaves, internal branches don't store data by themselves.\
this uses _std::variant_, _std::shared_ptr_ and visitors. the templated nature of these classes mean they must be exposed in the header file, even thought the interface doesn't require them.

> terminology
>
> - Tag - "a monodial type describing the tree"
> - Nodes of the tree can be:
>   - Branch - points to left and right tree
>   - Leaf - holds data and tag
>   - Empty - nill value, not nulls
>   - Tree - a variant of \<Empty, Leaf, Branch>

```cpp
//branch
template <typename Tag, typename Value>
class Branch
{
    Tag tag_;
    std::shared_ptr<Tree<Tag,Value>> left_;
    std::shared_ptr<Tree<Tag,Value>> right_;
};
// leaf
template <typename Tag, typename Value>
class Leaf
{
    Tag tag_;
    Value v_;
};

// empty
template <typename Tag, typename Value>
class Empty
{
    public:
        empty(){};
        auto tag() const -> Tag { return {};}
};

// Tree
template <typename Tag, typename Value>
class Tree
{
    private:
        std::variant<Empty_, Leaf_, Branch_> data_; //the class with _ are aliases
    public:
        Tree(Empty_ const & empty): data_(empty){};
        Tree(Leaf_ const & leaf): data_(leaf){};
        Tree(Branch_ const & branch): data_(branch){};
        //...
        template <typename Callable>
        auto visit(Callable && c) const
        {
            return std::visit(c, data_);
        }
};
```

operations on trees produces trees that share nodes with the original. Tree exposes factory functions that return _std::shared_ptr\<Tree>_ constructing empty, leaf or branch ("smart constructor" idiom).
tag of a branch is the result of adding the tags of it's left and right children. ags must be monodial (see slide)

expose function objects as interface, such as Depth (a callable struct that knows to visit Branch,Leaf, Empty), flattenToVector (another visitor), printer(external visitor object).

example. see slides for photos of results.

```cpp
auto t0= Tree::branch(
    Tree::branch(
        Tree::leaf(1),Tree:leaf(2)
        ),
    Tree:leaf(3)
    );

auto t1 = prepend(0,t0);
auto t2 = append(4,t1);

std::cout << "digraph G {\n";
printer _p(std::cout);
t0->visit(p);
t1->visit(p);
t1->visit(p);
std::cout << "{\n";
```

### Consideratios for a Module

Export has fine-grained control (compared to header files), we choose everything or just particular names. we should export the things the client needs to name. we shouldn't export implementation details and infrastructure (at least not initially). as a rule of thumb, if it's part of the tests, it's probably a good idea to expose it. when we write unit tests we usually probe the state of the objects, which is probably something the user will want to do.

Exporting code for inlining

> "if you want to export code as part of your interface you must explicitly inline. Functions defined in the class declaration are not implicitly inline in a module. Inlines cant not refer to anything with internal linkage."

this is unlike header files where the class defintion functions are implicitly inlined if we write them in the header file (think templates).

Organization is Not Exposed to Customers

> - "you can use partitions, the PMF, Module implementation units, and all of it looks the same to the customers."
>   - changes can be done without worries.
> - "Re-exporting a name from a different module might be visible. Names attached to modules, and that may be part of the name."
>   - name conflicts, depends on strong vs weak models of name ownership from modules. currently different for gcc and windows compiler.

### Hello World Module

includes before the modules are not exported.

```cpp
module;
#include <iostream>
#include <string_view>

export module smd.hello;

//export namespace and all it's content, names don't have to match
export namespace hello {
    void hello(std::string_view name)
    {
        std::cout << "Hello, "<< name << '\n';
    }
}
```

usage:

```cpp
import smd.hello; //import module
int main()
{
    hello:hello("steve"); //use function from namespace hello
}
```

makefile, requires g++11.\
_(couldn't get this to run on my machine)_

```makefile
hello_main: hello_main.o hello.o
	g++11 -o hello_main hello_main.o hello.o

hello_main.o: hello_main.cpp gcm.cache/smd.hello.gcm
	g++11 -fPIC -fmodules-ts -x c++ -o hello_main.o -c hello_main.cpp

hello.o: hello.cpp
	g++11 -fPIC -fmodules-ts -x c++ -o hello.o -c hello.cpp

gcm.cache/smd.hello.gcm: hello.o
	@test -f $@ || rm-f hello.o
	@test -f $@ || $(MAKE) hello.o

clean:
	rm hello.o hello_main.o gcm.cache/smd.hello.gcm

clean-gcm:
	rm gcm.cache/smd.hello.gcm

test:
	.hello_main
```

gcm is the binary artifact of modules which gcc produces.

### Coding

primary module interface.

> "every name that that clients consume is exported through the primary module interface. those may be rexported from module partitions or from other modules."

```cpp
module;
//global module fragment;

#include <non_module.h>
export module foo;
export import :part; //exports foo:part. a module partition
import std; // maybe we can do better someday
import bar; //not exported, reachable

export namespace foo{
    //everything inside is exported
    int theAnswer();
}
```

> modules compose\
> "as long as there is a strict dependency directed acyclic graph (**DAG**) between the more fine grained modules. the dot (.) is a a convetion. It has has no hierarchical meaning to the compiler."

```cpp
export module foo;
export import foo.bar;
export import foo.baz;
export import foo.qua;
```

module implementation units. almost the same as regular translation unit, except they have access to module linkage names.

```cpp
module foo;
int foo::theAnswer(){ //foo is the name space, not the module
    return 42;
}
```

Module partitions can be used to decompose large modules. no one outside the module can import a partition, only from inside the module. Module partitions have access to all of the names and defintions from the module interface.

```cpp
export module foo:part;
export int qua_foo(int);
```

Private fragment

> "A special partition that can appear in a primary module interface. they allow unexposed and and unreachable defintions to be included in the PMI."

example the standard

```cpp
export module A;
export inline void fn_e(); //error, exported inline function not defined before private module fragment

inline void fn_m(); //ok, module-linkage inline function
static void vn_s(); //old style 'static', only visible here

export struct X;
export void g(X *x)
{
    fn_s(); //ok, call to static function in the same translation unit
    fn_m(); //ok, call to module linkage inline function
}
export X *factory(); //ok

module :private; //special name
struct X{}; //defintion not reachable from importers of A
X *factory()
{
    return new X();
}
void fn_e(){};
void fn_m(){};
void fn_s(){};
```

export needs to be applied at the first point of declaration. with private module fragments, we can conceal the type, it's an opaque pointer,we can't get the type with `decltype`. this option mixes interface and defintions. probably not the best idea, but usefull to allow single file modules (like header only libraries). will probably cause a complete re-build process.

### Building Modules

things won't work smoothly, module names are not the file name which means parsing c++ rather than the usual file-based dependency lookup for current configurations.\
we must built the module in the DAG order (directed acyclic graph), and hope that linking will fail rather than simply being wrong (modules imported before being built).\
We will probably have a period where we use something like 'MakeDeps' from the old times,which run before the build and updated all the dependencies. the makefile example did that by hand.

we don't have a solution yet for packaging modules, stuff are still too dependant on compiler internals.

### Modulating Fringetree

it worked for him the day before, depending on the compiler.

the header file moved into a interface file (_.ixx_)

- smd is a module space, arbitrary name.
- the fringetree namepace is not exposed.
- the node types are reachable, but not visible

```cpp
//fringetree.ixx
module;
#include <memory>
#include <variant>
#include <vector>

export module smd.fringetree;
namespace fringetree{

//branch
template <typename Tag, typename Value>
class Branch
{
    //same as before
};

// leaf
template <typename Tag, typename Value>
class Leaf
{
    //same as before
};

// empty
template <typename Tag, typename Value>
class Empty
{
    //same as before
};

// Tree
export // make the Tree template available
template <typename Tag, typename Value>
class Tree
{
    //same as before
};
}
```

usage

```cpp
using namespace fringetree;
using Tree = tree<int, int>; //alias
auto t0 = Tree::branch(Tree::branch(Tree::leaf(1),Tree:leaf(2))
                        ,Tree:leaf(3));
```

inline only happens explicitly. the definition of 'branch' is not inlined in client code, trade off with exposing implementation vs optimization oppertunities

```cpp
// Tree, still inside the namespace from before
export // make the Tree template available
template <typename Tag, typename Value>
class Tree
{
    //same as before
    public:
    bool isEmpty() {return std::holds_alternative<Empty_>(data_);} //not inlined
};
```

exporting function objects

```cpp
constexpr inline struct breadth
{
    template <typename T, typename V>
    auto operator()(Empty<T,V> const &) const -> T
    {
        return 0;
    }
        template <typename T, typename V>
    auto operator()(Leaf<T,V> const &) const -> T
    {
        return 1;
    }
        template <typename T, typename V>
    auto operator()(Branch<T,V> const &) const -> T
    {
        return b.left()->visit(*this) + b.right()->visit(*this);
    }
} breath_;
//lambda object that uses the breath_ object. exported.
export constexpr auto breadth = [](auto tree){return tree->visit(breadth_);};
```

</details>

## Preconditions, Postconditions, Invariants: How They Help Write Robust Programs - Andrzej Krzemieński

<details>
<summary>
Contracts as opposed to interfaces, how can we detect them, what can we do to enforce them.
</summary>

[Preconditions, Postconditions, Invariants: How They Help Write Robust Programs](https://youtu.be/4Qyu8uBrRUs), [slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/andrzej_preconditions.pdf)

### Function Contract

we start with a need, then we write the interface, and eventually we write the code that will satisty the need.

lets say we have a collection of aircrafts and we want to sort them according to their weight.

```cpp
double Aircraft::weight() const; //weight in kilograms
std::vector<Aircraft*> aircrafts;
bool is_lighter(const Aircraft *a,const Aircraft b*);
```

abstraction gaps: the difference between what we mean and what we write.

- we mean aircrafts, but we actually have memory address (pointers).
- we mean weight in KG, but we have type _double_.
- pointers are abstractions, we have memory address, but we assume they aren't null pointers, and they refer to an actual valid object.

another need, determine if an integral value is in a close range.

```cpp
int lower = config.get("LOWER_BOUND");
int upper = config.get("UPPER_BOUND");
bool is_in_range(int val, int lo, int hi);
bool f(int a, int b, int c);
```

the two functions are the same for the type system, but the first has some sort of assumption that is part of the abstraction. we can call it in a way that fits the type system, but not the abstraction.

> _Function Contract_ - all you need to know to call the function correctly
>
> - How the inputs and outputs are interpetted.
> - Limits (domain).
> - What must happen before or after the call.
> - Effects.

parts of the contracts can be managed by the language, such as the number and the types of the arguments, and their immutability. other parts cannot be managed: exception safety guarantees, disallowed values.

we can define _interface_ as either the complete function contract, or as the parts of it which can be expressed in the language (which is the conventional defintion).

in the 'setlocale' case, locale is an optional string, it **may be** null. in the 'strlen' case, s is a required string, it **must not** be null. we can't get the complete picture just from reading the function signature (the interface)

```cpp
char * setlocale(int category,const char * locale);
size_t strlen(const char* s);
```

### Strong and Weak Contracts

```cpp
bool is_in_range(int val, int lo,int hi)
{
    // checks if val is in closed range [lo,hi]
    // expects lo <= hi
}
is_in_range(3,2,1); // valid language expression, but not the intended use
```

what do we do? how do we cope with the possibility of passing unreasonable values? we can weaken the contract.

```cpp
bool is_in_range(int val, int lo,int hi)
{
    // if lo <= hi
    //      checks if val is in closed range[lo,hi]
    // otherwise return false;
}
```

> Weak contracts have drawbacks:
>
> - Weaker abstractios.
> - Increased complexity.
> - No use.
> - False sense of safety.
> - Missed opportunity to detect bugs.

we no longer talk about ranges, we moved to lower level abstractios, this encourages bugs. we add additional if statement, which means the code is more complicated, and we need more testing to be sure.

'no use' means that we code for a situation that shouldn't happen.

in the aircraft example, we add a check for null, which is an additional responsability, and more importantly, we ignore the serious question of why would someone pass a null pointer to this function anyway?

```cpp
bool is_lighter(const Aircraft *a,const Aircraft b*);
// if a or b is null, retrun false
// otherwise return whether *a has lower weight than *b
```

false sense of security: what if there are other bad values, like 0xffff`ffff or 0x12345678? we only protected against one possible bad value.

the weaker contract hides the bug in this example, we mistakenly messed up the order of the arguments. if every value is "good", we don't detect "bad" input.

```cpp
bool validate(int val)
{
    int lower = config.get("LOWER_BOUND");
    int upper = config.get("UPPER_BOUND");
    retrun is_in_range(lower, upper, val); //wrong order of arguments
}
```

in contract, strong contract are the opposite.

> Strong Contacts:
>
> - Simple abstraction.
> - Discourage bugs.
> - Can help detect bugs.

we shouldn't weaken contracts, we need to look for language features that enforce strong contracts.

### Checking Contract Violation

```cpp
bool is_lighter(const Aircraft *a, const Aircraft *b);
//expects a and b point to objects of type Aircraft
```

pointer dereference. the need is to access the object under a given address. it can't be null, and it must be a valid Aircraft objects.

```cpp
bool is_lighter(const Aircraft *a, const Aircraft *b)
{
    if (!a || !b) std::abort(); //injected by UB-sanitizer
    return a->weight() < b-> weight();
}
```

if the code is entirely available, a static analyzer can see both defintion and usage, it can warn about bugs. inlining helps.

```cpp
inline bool is_lighter(const Aircraft *a, const Aircraft *b)
{
    return a->weight() < b-> weight();
}

Aircraft *p = nullptr;
Aircraft *q = getAircraft();
Aircraft *r = getAircraft();
return is_lighter(p,r); // static analyzer warning.
```

std::string has in it's contract that it cannot receive a null pointer to the constructor. a clever compiler (or a UB sanitizer) could detect this.

```cpp
const char *p=nullptr;
const char *q="config.cfg";

if (!p) std::abort(); //can be injected by UB-sanitizer
std::string file_name(p);
```

the library can also enforce the contract with asserts.

```cpp
basic_string(CharT* __str) : /*... */
{
    assert(__str != nullptr); // contract enforcement in std implementation
    //...
}
```

we would like similar behavior in our code. _Contract check_ - derived from a function contract. It can determine that a program has a bug.

when we write the implementation of the code, we forget about the world outside. and we implicitly prefer bugs over UB. we assume we can avoid UB entirely.

#### Preconditions

```cpp
bool is_in_range(int val, int lo, int hi)
{
    return lo <= val && val <= hi;
}

bool exceeds_weight(const Aircraft 8 a, double limit_kg)
{
    if (!a)
    {
        REACT(); //try to avoid UB of dereferencing a null pointer.
    }
    return a->weight() > limit_kg;
}
```

if our functions are visible to the static analyzer, it might be able to detect the UB bug, rather than simply hide it.

```cpp
bool exceeds_weight(const Aircraft 8 a, double limit_kg)
{
    if (!a)
    {
        return true; // "safe deafult"
    }
    return a->weight() > limit_kg;
}
```

this also assumes the user doesn't care for the difference betwen exceeding the weight limit and being unable to compute the result (and getting the default true return value).\
Imagine the same function in a different context:

```cpp
bool enough_fuel = exceeds_weight(&a,required_fuel);
if (!enough_fuel)
{
    report_danger();
}
```

the user did something stupid and used the function for the wrong intent. the default behavior is now very wrong and potentially dangerous.

we can throw an exception, which won't necessarily be detected by a static analyzer, and postones the bug detection to runtime. it assumes that after throwing the exception, the program is in a valid, not bugged, state.

```cpp
bool is_in_range(int val, int lo, int hi)
{
    if(lo > hi) throw Bug{}; //skip me and the caller
    return lo <= val && val <= hi;
}

bool exceeds_weight(const Aircraft 8 a, double limit_kg)
{
    if (!a)
    {
        throw Bug{}; //skip me and he caller
    }
    return a->weight() > limit_kg;
}
exceeds_weight(nullptr,120`000); //no help from static analyzer
```

in this example, the arguments are in the correct order,
but

> "Upon UB your code no longer coressponds with _the binary_"\
> "You cannot draw any conclusions from the code"

```cpp
bool validate(int val)
{
    int lower = 0;
    int upper = std::max(100, config.get("LIMIT"));
    return is_in_range(val,lower,upper);
}
```

> Preconditions violation can be caused by prior undefined behavior:
>
> - Dangling pointers.
> - Bad usage of _memset_.
> - Data races.

#### Undefined Behavior

example:

```cpp
bool is_small (int * p)
{
    return *p <10;
}

is_small(nullptr);
```

what does this mean in the language? pointer dereference means accessing memory under address. nullptr means no address. so we have a contract violation,no guarneed behavior, a bug.

but why is there no guaranteed behavior? why aren't all UB defined?\
the reason is that there is no use case, there is no viable case for a program to meaningfully dereference a null pointer. this would be sanctioning bugs.

(_there is quote somewhere about how users use every observable part of the interface, not just the intended ones. somewhere there is a user relining on our bug to make their programs work_)

the outcome of the code is either getting some junk code, or (more likely), a program shutdown.

if we update our code, we trade 'bugs' for crashes, which might be ok sometimes, but we still have to worry about it being used in situations where crashes are dangerous.

```cpp
bool is_in_range(int val, int lo, int hi)
{
    if(lo > hi) std::abort(); //fail on this case,
    return lo <= val && val <= hi;
}

bool in_danger(Aircraft const &ac)
{
    return is_in_range(danger_zoner.lower(),
                        danger_zone.upper(),
                        ac.stress()); //order of arguments
}
```

the bug is outside the function (wrong order), and the implementation might change in the future. the positive sides of calling abort are that we protect against further damage, we usually get a core-dump for fixing, and it keeps our program stable.

```cpp
bool is_in_range(int val, int lo, int hi)
{
    if(lo > hi) std::abort(); //fail on this case,
    _stats[hi-lo] +=1; //for logging, it critical that hi > lo here, otherwise we get a negative index.
    return lo <= val && val <= hi;
}
```

> Is crashing a good idea? depends on the context:
>
> - Weight & balance calculator: user can go to manual.
> - Word processor: better to waste 1hr of work that 1 day of work.
> - Drone: better to restart than do random actions.
> - Financial server: better go down than make bad decisions.
> - Assisting troops: better go down than give false sense of security.

the primary goad of the applications it to give good results, it's better to have no program (crashing) than give misleading results. not crashing is a secondary requirement.

giving a hint to the static analyzer. explicitly invoking UB.\
a bug is a violation of the function contract.\
UB is a violation of the language contract.

```cpp
inline bool is_in_range(int val, int lo, int hi)
{
    if(lo > hi)
    {
        *((int*)0)=0; //explicit UB,
    }

    return lo <= val && val <= hi;
}

is_in_range(0,100,50); //static analyzer warring, null pointer dereference
```

we can do this conditionally with a macro.

```cpp
#ifdef STATIC_ANALYSIS
#define TRAP() {*((int*)0)=;} //explicit UB, for testing ca
#else
#define TRAP() std::abort() //abort, for production
#endif
inline bool is_in_range(int val, int lo, int hi)
{
    if(lo > hi) TRAP();

    return lo <= val && val <= hi;
}

is_in_range(0,100,50); //static analyzer warring, null pointer dereference
```

this is very close to what assert does. it declares the bug criterion, and its' effect is depending on the configuration. assert statements is about communicating intent, what we know about the behavior. we know what the situation should be, and what is considers a bug.

```cpp
bool is_in_range(int val, int lo, int hi)
{
    assert(lo <= hi);
    //Expects(lo<= hi)l //GSL.assert
    return lo <= val && val <= hi;
}
```

GSL.assert: `Expects()` is for preconditions, `Ensures()` is for postconditions
we can still get bugs,

```cpp
bool check(int lo, int hi, int val)
{
    return is_in_range(lo,hi,val); //wrong order, but check passes
}
return check(1,2,3);
```

contract checks can determine that we have a bug, but they cannot determine we don't have a bug.

> - we **don't mean** that `a != nullptr && b != nullptr` is good.
> - we **mean** that `a == nullptr || b == nullptr` is bad.

```cpp
bool is_lighter(const Aircraft *a, const Aircraft *b)
{
    //asserts, expects, etc...
    return a->weight() < b-> weight();
}
```

a better approach is to wrap the two integers into a class, this will protect us from some problems with the order of the arguments. we got the type system involved.

```cpp
bool is_in_range(int val, Range r)
{
    //expects r.lo() <= r.hi()
}

bool validate(int val)
{
    int lower = config.get("LOWER_BOUND");
    int upper = config.get("UPPER_BOUND");
    retrun is_in_range(val,Range{lower,upper});
}
```

### Constrained Types for Enforcing Contracts

#### Invariant

what we gained was both a stricter check of the arguments passed, and now we made the properties of lo and hi into class invariants. we expose the constraint as a class invariant.

```cpp
class Range {
    int _lo, _hi;
    public:
    Range(int l, int h)
    int lo() const {return _lo;}
    int hi() const {return _hi;}
    //invariant lo() <= hi()
};
class Range2 {
    int _lo, _hi;
    public:
    Range2(int l, int h)
    int lo() const {assert(invariant()); return _lo;}
    int hi() const {assert(invariant()); return _hi;}
    bool invaraint() const {return lo() <= hi();}
};

bool is_in_range(int val, Range r)
{
    //expects r.invarint(); // this is redundant
}
```

we can check the invariant inside the calls to the objects, or before using it (but that is redundant). an invariant on function parameters is an implied precondition.

it's still possible to pass the values in the wrong order;

```cpp
if (is_in_range(1,Range(highLimit,lowLimit))) //oops, wrong order
```

lets try to enforce it

```cpp
class Range {
    int _lo, _hi;
    public:
    Range(int l, int h) : //precondition l <= h
    _lo((assert(l<=h),l), //use the comma operator
    _hi(h)
    {}
    int lo() const {return _lo;}
    int hi() const { return _hi;}
    bool invaraint() const {return lo() <= hi();}
};
```

> "An invariant is a _conditional_ guarantee. It depends on the preconditions of member functions."

we still have bugs, we only catch them when we create the range, but we narrowed the scope. if we could propigate this type across the program, we would get reduce the bugs earlier.

if only one function can create ranges, we know where the bugs can originate. the bugs cannot happen from passing the object, only from constructing it.

```cpp
auto [a,b]= bounds(); //a >b
is_in_range(x,Range{a,b}); // bug
is_in_range(y,Range{a,b}); // bug
is_in_range(z,Range{a,b}); // bugs
Range r= boundsR();
is_in_range(x,r);
is_in_range(y,r);
is_in_range(z,r);
```

not everything can be expressed as a contract

not every precondition can be turned into a type. we can't precondition the indexing of the vector `[]` operator, because the size can change.

```cpp
class Kilograms {
double _value; // weight difference can be negative
public:
explicit Kilograms(double val) : _value{val} {} //we don't want random constructor calls
static Kilograms from_double(double val) { return Kilograms{val};
};
Kilograms Aircraft::weight() const;
```

our classes reflect the invariant, and how the values are interpretted.

```cpp
size_t length(const char* str)
// expects: str != nullptr
// expects: "str is null terminated" // inexpressible,
;
class C_string {
const char * _str;
public:
// invariant: _str != nullptr && "str is nul terminated"
};

size_t length(C_string str);
C_string str = get_name()
return length(str);
```

but if we eventually do need the `const char *` type for other libraries, we would need a converting constructor. we cant make this runtime check, but maybe a static analyzer could detect bug.

```cpp
class C_string {
const char * _str;
public:
explicit(false) C_string (const char * s)
    //expects: s!= nullptr && "s is nul terminated"
// invariant: _str != nullptr && "str is nul terminated"
};
```

a fantasy, a function that changes it's behavior on static analyzer runs, it's a no-op in runtime, but matches preconditions and postconditions of function.

#### Post Conditions

> A postcondition is expected to hold if:
>
> - Function does not report failure (e.g., by throwing exception)
> - Function’s preconditions are satisfied

### Language Support

what can add to the language to support this?
so far we saw comments, but they are for humans or some dedicated tools. we saw assertions, which help prevent damage, help testing, but don't detect bugs. post-condition asserts require changing implementation (making code longer, possibly messing with RVO).\
a possible solution would be 'contart Annotations', which would appear in function declarations, and are checked by the compiler.

```cpp
int better(int a, int b)
[[pre: a >= 0]] // fantasy contract annotation
[[pre: b >= 0]] // fantasy contract annotation
[[post r: r>=0]] // fantasy contract annotation
;
```

> Standardized contract anntations:
>
> - Communicate (parts of) our contracts to different tools.
> - Provide same bug-detection experience as with language contracts.
> - Static Analyzer can detect bugs without seeing function bodies (no need to inline).
> - The IDE could hint the user when using the function (an correctly map the names of the argument and parameters)
> - The UB sanitizer could inject runtime checks based on the precondition.

however, this will make the function declaration even more cumbersome. this entire block of code is just one function declaraton ('clone').

```cpp
class window : public widget {
[[deprecated("use simpler decl")]] [[nodiscard]] const widget & clone
(const point & x, const point& y) const & noexcept override
[[pre: are_rectangle(x, y)]]
[[pre: ordered([](const auto& a) -> auto && { return a.h; }, x, y) ]]
[[post x: *this == x]];
};
```

### Bug Source vs Bug Symptom

fixing bugs isn't the same as concealing the source of those bugs.

```cpp
Aircraft * get_aircraft(std::string_view name)
[[post a: a!= nullptr]]; // function declaration with contract annotation
//...
Aircraft *x = nullptr, *y = nullptr;
try
{
    x = get_aircraft("x");
    y = get_aircraft("y");
}
catch(...) {
    // TODO: handle it
}
if (!x)
    return is_lighter(x, y); // <-- analyzer warning: y might be null
```

the analyzer only mentions y, not x. it doesn't understand the bug.

> Analyzer detects only a symptom of a bug.
>
> - Analyzer does not know where the bug is.
> - Programmer must look for it.

we might fix the symptom (the warning), without fixing the cause. the situation is worse, we don't see the warning anymore, but the bug didn't go away.

```cpp
//...
if (!x || !y)
    return is_lighter(x, y); // warning silenced, bug is still present
```

the real bug was that we had a premature try-catch block, we resumed normal operations too quickly.

```cpp
Aircraft *x = nullptr, *y = nullptr;
try
{
    x = get_aircraft("x");
    y = get_aircraft("y");
    return is_lighter(x, y);
}
catch(...)  // <-- real bug: the premature try catch statement
{
    // TODO: handle it
}
```

and if we don't know what to do with the exception, why are we even catching it? (other than to pinpoint it's location?).

```cpp
Aircraft *x = get_aircraft("x");
Aircraft *y = get_aircraft("t");
return is_lighter(x,y);
```

a warning is a start of an instigation process.

### Summary

> Function contract
>
> - Inexpressible, human-to-human
> - Try to express parts of it in the language
> - Classes for providing interpretation of values
> - Contract annotations for expressing disallowed values
>
> Contract annotations – not only about runtime checks
>
> - Provide same tool experience as the language contract
> - Static analysis
> - IDE hints
> - Human understanding

</details>

## Weak Interfaces --> Weak Defences: The Bane of Implicit Conversion in Function Calls - Richárd Szalay

<details>
<summary>
Finding situations where we might have incorrect calls to functions because of how conversions work, we should write defensive design to make misusing hard.
</summary>

[Weak Interfaces --> Weak Defences: The Bane of Implicit Conversion in Function Calls](https://youtu.be/-UW4tA5r2QE), [Slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Richard-Szalay_Weak-Interfaces-Weak-Defences_SLIDES.pdf)

preventing misuse of funcions

### Argument Swapping

- "Name-based analysis"
- Reactiveness --> Proactiveness
- Existing guidelines

one problem is calling the arguments in the wrong order, or calling them with arguments that don't have the right meaning. adjacency increases the chance of mistake, the more parameters the function takes, the more likely it is for the user to make a mistake.

```cpp
int f (std::string host_name, int port, std::string message);
int f2 (std::string host_name, std::string message,int port);
std::string author = "richard";
std::string greeting = "hello, world";
f(author,8080,greeting); //correct order, but author is not a host name
f2(greeting,author,8080); //correct order, but author is not a host name
```

one idea of detecting swapped arguments is to use "name-based analysis", which somehow tries to detect swapped arguments, algorithms exist, but tools aren't as common.

Ways to detect swaps: trying to figure out if one argument is a better match to a different parameter than the one it was sent to.

- String equality, suffix/prefix coverage, pattern containment
- Edit distances (e.g. Levehnstein-distance, ...)
- Morpheme extraction
- Typo analysis (e.g. key distance)
  The problem is that this approach requires naming, not only does this require the user to properly name stuff, it also has problem with literals and expressions as arguments. and some API signature don't provide parameter names.

the problem is that this approach is reactive, it helps after the fact, it requires the code to be deployed and analyzed afterwards, and it has bad accuracy. we would want something that helps us before the code is deployed, that catches potential errors in compile time, and possibly takes advnatage of the type system.

the c++ core guidelines contain a section about avoiding unrelated adjacent unrelated parameters of the same type.

we are in a static, strongly typed language, we can use this for our advantage.

> issues
>
> - typedef / using
> - const T & troubles
> - const T == T?

example, will this be caught?

```cpp
struct Complex
{
    double Re,Im;
    //...
};
void h (int Scalar,Complex comp);
void test()
{
    int S = 8;
    Complex C = Complex{.5f,-.25f}; // = (1/2 - 1/4i)

    h(C,S); // bug
}
```

the answer is that this depends, if the complex number can be converted (like with conversion to double which is then narrowed), and if the complex number can be created from a integer (like having a constructor that takes a double and uses 0 for the imaginary part), then the compiler will accept this call. this is a bug.

```cpp
struct Complex {
    double R, I;
    Complex(double real): R(real),I(0.0){}
    operator double() const {return R;}
    //...
}
void h (int Scalar,Complex comp);
void test()
{
    h(Complex{.5f,-.25f},8); // same call as below
    h(0,Complex{8.0}); // same call as above
}
```

### Implicit Conversions

there is a sequence of conversions, not all stages must be used, but if they are, the must be followed in order.

> An **implicit conversion sequence** T1 -> T2 exists and defined as:
>
> 1. _Maybe_ **standard conversion sequence**:
>    1. Maybe decay (lvalue→rvalue, array/function→pointer)
>    2. Maybe numeric adjustment
>       - integral ↔ integral, enum → integral
>       - floating ↔ integral
>       - derived → base class, T\* → **void\***
>       - **null**-constants → pointer (?!)
>       - **Anything you can imagine → bool (!!)**
>    3. Maybe function pointer adjustment ("lose _noexcept_")
>    4. Maybe qualifier adjustment ("gain _const volatile_")
> 2. Maybe user-defined conversion (one function call!)
> 3. Maybe standard conversion sequence (same as above)
>
> ... if the path taken uniquely exists.

it might be possible to remove some of those conversions in the language, but that might break user code everywhere (even if the standard library will be ok), but some conversions can't be removed.

but there exists a better way, we can use the type system. we should provide good types and good interface, in other words. **strong(er) type(s)**.

### Type Based Guards

#### Mixable Adjacent Parameter Ranges

```cpp
void p (int i, int j,, double d, Complex c, std::string s);
```

we go through our functions parameters and determine which of them might be mixed with one another.

> - int mixed with int (same type)
> - int mixed with double (standard conversion)
> - int mixed with Complex (standard conversion + user conversion)
> - double mixed with Complex (user conversion)
> - complex can't be mixed with with std::string

there is a proposal for a clang-tidy check, which performs this mixable parameters analysis, and tells us which parameters are easily mixable and prone to errors

lots of graphs of analysis of popular packages. where is the danger coming from?

some ways to reduce the number of false positive, we get far too many warning.
there are some funcions that we can't really fix, like max, or binary operations on the same type, and functions that concatenate strings. they have to take the same type of parameters, and the result of swapping them won't 'really' effect the output of the program.

> "avoid adjacent **unrelated** parameters of same type" \
> "... can be invoked .. either order ... different meaning"

what does "unrelated" mean?

can we deterimine which parameters want to be used together? can we check if the order matters? this is a way to remove what we think is a false positive.

when we get those warnings, we find that we might have some strong types waiting to be found, or some functions that should have different arguments, etc...

### Summary

> - Built-in types are bad, implicit conversions are dangerous.
> - Help users not to make mistakes, before the fact.
> - Use tools to find the low hanging fruits.

</details>

## hop: A Language to Design Function-Overload-Sets - Tobias Loew

<details>
<summary>
A library to better match overload in templated functions.
Limiting types and using multiple arguments while retaining the regular overload resolution rules.
</summary>

[hop: A Language to Design Function-Overload-Sets](https://youtu.be/XykLWWBHXGk),[github](https://github.com/tobias-loew/hop)

### Homogenous Variadic Functions

Homogenous Variadic Functions (HFV) are functions such as

```cpp
void logger("this","is","a","log","message");
double max(0.0,d1,d2,d3,d4,d5);
Array<double,1> vec(1000);
Array<double,2> matrix(6,100);
Array<double,3> cube(2,3,5);
```

> "Homogenous Variadic Functions are functions that take an arbitrary (non-zero) number of arguments, all of the **same type**."

blitz++ is a library that started with c++98, it was used for linear algebra calculations, pioneered expression templates.\
it has an custom Array class template, with many, many constructors and `operator()` overloads, as well as several other Homogenous Variadic Functions with reapeating code.

[fluent-c++: How to Define a Variadic Number of Arguments of the Same Type](https://www.fluentcpp.com/2019/01/25/variadic-number-function-parameters-type/) a multipart blog post about different approches.

in c++98, we would write overloads manually. it was explicit, easy to understand, beginner friendly (no templates or macros), but this is the example of bad code writing and violating DRY.

```cpp
R foo(A a1);
R foo(A a1, A a2);
R foo(A a1, A a2, A a3);
R foo(A a1, A a2, A a3,A a4);
R foo(A a1, A a2, A a3,A a4,A a5);
```

c++98 also had **Boost.Preprocessor**, which is the exact opposite, it avoids repetitions and easily scalable, but it's hard to write, harder to read, even harder to debug. and the number of arguments is still fixed.

```cpp
#define NUMBER_OF_ARGUMENT_OF_F 10
#define GENERATE_OVERLOAD_OF_F(Z,N,_) \
    R foo(BOOST_PP_ENUM_PARAMS(N,A a)) {/*...code...*/}

BOOST_PP_REPEAT_FROM_TO(1,BOOST_PP_INC(NUMBER_OF_ARGUMENTS_OF_F), GENERATE_OVERLOAD_OF_F,_)
```

in c++11 we got _std::initializer_list_, which works for any number of arguments, but requires additional braces when called, is inflexible for mutability (need to use reference wrapper of some sort), and can even be called with empty list, and causes weird behavior for overload resolution

```cpp
R foo(std::initializer_list<A> as)
{
    /*...*/
}
```

but we also got parameter packs, which could work for any number of arguments, but on the downside, matched any input, and hade potential for hard to define errors in compilation and runtime, and could lead to logical errors.

```cpp
template <typename T,typename... Ts>
R foo (T,t,TS.... ts)
{
    /*...*/
}
```

this could be expanded to incorporate SFINEA, below is a c++14 example, where we could limit the number of argument to be non zero, and matches only valid input

```cpp
template <typename ... Ts,
    std::enable_of<(sizeof...(Ts) > 0) && //at least one argument
    (std::is_convertible_v(Ts,A) && ...) // all arguments are convertible to A
    >*= nullptr>
R foo(Ts&&... ts)
{
    /*....*/
}
```

c++20 simplifies the syntax, replacing SFINEA with requires.

```cpp
template <typename ... Ts>
requires
    (sizeof...(Ts) > 0) && //at least one argument
    (std::is_convertible_v(Ts,A) && ...) // all arguments are convertible to A
R foo(Ts&&... ts)
{
    /*....*/
}
```

future versions of c++ might make this even easier

P1219R2 suggested HFV parameters become supported by native language, so we could write template parameter packs directly (without template parameter types), which would be simpler and finally easy to understand, but the syntax would conflict with varargs syntax. apparently the comma before the trailing `...` is optional, so varargs `R foo(int,...)` is equivalent to `R foo(int...)`. so this suggestion was scrapped and is consider dead today.

```cpp
R foo(A a, A...as){/*....*/}
```

| version        | usage                                    | pros                                                        | cons                                                                |
| -------------- | ---------------------------------------- | ----------------------------------------------------------- | ------------------------------------------------------------------- |
| c++98 vanila   | manually writing overloads               | easy to understand                                          | violation of DRY                                                    |
| c++98 Boost    | Boost.Preprocessor macro                 | write once                                                  | hard to use                                                         |
| c++11          | std::initializer_list                    | works for any number of arguments                           | additional braces {}, mutability issues, overload resolution issues |
| c++11          | parameter pack                           | simple                                                      | matches any input                                                   |
| c++11          | parameter pack + SFINAE                  | matches any non-zero number of arguments, only valid inputs | none                                                                |
| c++20          | parameter pack + _requires_              | simpler syntax than SFINEA                                  | none                                                                |
| P1219R2 (dead) | Homogenous variadic functions parameters | native language support for HVF                             | syntax conflict with varargs                                        |

the current best practices for HVF are to use parameter packs with SFINAE or with the _requires_ clause (if possible). but those ignore overload resolutions and implicit converstions.

in the following example, there are two visible overloads, one for int types and one for float, there are also thee calls with different arguments, so which is called? does it even compile?

```cpp
template <typename ... Ts>
requires
    (sizeof...(Ts) > 0) && //at least one argument
    (std::is_convertible_v(Ts,int) && ...) // all arguments are convertible to int
R foo(Ts&&... ts)
{
    /*....*/
}

template <typename ... Ts>
requires
    (sizeof...(Ts) > 0) && //at least one argument
    (std::is_convertible_v(Ts,float) && ...) // all arguments are convertible to float
R foo(Ts&&... ts)
{
    /*....*/
}

foo(1,2,3); // (a)
foo(0.5f, -2.4f); //(b)
foo(1.5f,3); // (c)
```

in actuality, all calls are ambiguous and the program doesn't compile. [compiler explorer](https://godbolt.org/z/YbbY5ah7d). all arguments are convertible to eachother (int to float, implicitly), both overloads use forwarding reference `&&` which gives a perfect match, and according the SFINAE condition,only tests if overload is viable, it is not part of overload resolution. both overloads are equivalent.

Can this be solved by c++20 concepts?

> \[over.match.best.general]: desired overload must be _more constrained_
>
> - requires creating concepts like `Matches_T1_BetterThan_T2`.
> - this means rebuilding overload-resolution with concepts.
> - maybe possible for two types, probably messy for 3 or more.
> - technical issue: fold expression over constraints is an atomic constraint.

so for now, parameter pack + SFINAE (either clean or with the _requires_ clause) is the best we have, but it doesn't play well with overload resolution. if more than one overload is viable, we get ambiguity and errors.

can the solution be letting overloaded function know about one another?

> merging overloaded HVFs: step-by-step guide
>
> 1. provide interface to specify the overloaded type.
> 2. perform overloaded resolution
>
>    - generate all possible overloads
>    - use built-in overload-resolution to resolve call
>
> 3. report select type to the user
>
> - this guide uses C++17 for simplicity.
> - this guid uses Peter Dimov's **Boost.MP11**

the interface to the overloaded types uses template parameter pack, and is encapsulated in Boost.MP11 `mp_list`.

arity means the number of arguments.

```cpp
// list containing types for overloaded-resolution
using overloaded_types = mp_list<int, float>;

// helper alias template
template <typename... Args>
using enabler = decltype(enable<overloaded_types,Args... >());

// "overloaded" HVF for int and float
template <typename... Args, enabler<Args...>* = nullptr>
void foo(Args&& ... args) {/*...*/}
```

to perform the overload resolution, we generate all possible overloads (which is a possbile infinite number, but actually just for those viable on the call site), we generate the test-function for they type and run the built-in overload resolution on all of them.

library for overloading HVFs (26 LOCS), complete code.

```cpp
template <size_t Index, class Params>
struct single_function;

template <size_t Index, class... Params>
struct single_function<Index, mp_list<Params ...>>
{
    constexpr std::integral_constant<size_t,Index> test (Params...) const;
};

template <size_t arity, class Indices, class... Types>
struct _overloads;

template <size_t arity, size_t... Indices, class... Types>
struct _overloads<arity, std::index_sequence<Indices...>,Types...>
    : single_function<Indices, mp_repeat_c<mp_list<Types>,arity>>...
{
    using single_function<Indices, mp_repeat_c<mp_list<Types>,arity>>::test...; // c++17 required
}

// required to create compile time index-sequence
template <class Types, size_t arity>
struct overloads;

template <class .. Types, size_t arity>
struct overloads<mp_list<Types...>,arity>
    : _overload<arity,std::index_sequence_for<Types...>, Types...>
{};

template <class Types, typename... Args>
constexpr decltype(overloads<Types, sizeof...(Args)>{}.test(std::declval<Args>()...)) enable();
```

that's too much to understand at once, so we look at it piece by piece, first, the enable function. it is invoked by the caller with type list and actual arguments. the return type is the result of the overload resolution calling `test`, which will be an `std::integral_constant<size_t,Index>` indicating the best match.

```cpp
template <class Types, typename... Args>
constexpr
decltype(overloads<Types, sizeof...(Args)>{}.test(std::declval<Args>()...)) enable();
```

we now move to the part where we are creating the 'viable' type-lists. we get all the types, and create the list of the types with desired arity. instantiate `single_function` with such a list for each 'overloaded' type using pack expansion.
the `using single_function`... line brings all the base class functions into the scope of the struct and make them visible.

```cpp
template <size_t arity, class Indices, class... Types>
struct _overloads;

template <size_t arity, size_t... Indices, class... Types>
struct _overloads<arity, std::index_sequence<Indices...>,Types...>
    : single_function<Indices, mp_repeat_c<mp_list<Types>,arity>>...
{
    using single_function<Indices, mp_repeat_c<mp_list<Types>,arity>>::test...; // c++17 required
}
```

the 'overloads' (without the underscore) are used for unpacking the mp_list and getting the types and index_sequence.

```cpp
// required to create compile time index-sequence
template <class Types, size_t arity>
struct overloads;

template <class .. Types, size_t arity>
struct overloads<mp_list<Types...>,arity>
    : _overload<arity,std::index_sequence_for<Types...>, Types...>
{};
```

the core of the library generates the test function, the function generated is the overload of each option.

> - instanitated for each overloaded type
> - the `sizeof...(Params)` is the arity of the call.
> - defines `test` with desired argument type and arity.
> - this is also used in the STL implementions to genereed the converting constructors in _std::variant_ for type-alternatives.

```cpp
template <size_t Index, class Params>
struct single_function;

template <size_t Index, class... Params>
struct single_function<Index, mp_list<Params ...>>
{
    constexpr std::integral_constant<size_t,Index> test (Params...) const;
};
```

the final part is to report the selected type or fail in compilation. we have a function that decides in compile time which 'overload' is used.

```cpp
// list containing types for overload-resolution
using overload_types = mp_list<int,float>'

// helper alias template (simplifies repetitions)
template <typename... Args>
using enabler= decltype(enable<overloaded_types, Args...>());

// "overloaded" function for int and float
template <typename... Args,enabler<Args...>* = nullptr>
void foo(Args&&... args)
{
    constexpr auto selected_type_index = enabler<Args...>::value; //index of the select overload
    if constexpr (selected_type_index == 0)
    {
        // int overload invoked
        std::cout<< "overload:(int,...") <<'\n';
    }
    else if constexpr (selected_type_index == 1)
    {
        // float overload invoked
        std::cout<< "overload:(float,...") <<'\n';
    }
}
```

[example in compiler explorer](https://godbolt.org/z/94nGrbaxd)

this time the call with the none-matching types is rejected. also the empty call is rejected. this is what we would like to see. if we use a double instead of float, we get an error, because conversion from double to float and from double to int are the same priority in overload resolution, but it does work if we overload on double and send floats.

### Hop - Function-Parameter Building Blocks

> **"From HVF to function-parameter building blocks"**
> a solution to expand the previous idea and make things better.

> required features for the library
>
> - create parameter-list with type-generators.
> - create overload-resultion condition from overload-set.
> - within the function implementations
>   - identify active overload.
>   - aceessing /forwarding arguments (especially defaulted parmeters).
>
> hop: type-generators
>
> - any cv-qualified C++ type
> - repetition
> - sequencing
> - alternatives
> - trailing/not-trailing default parameters.
> - forwarding references with/without condition
> - template argument deduction (per argument, global, mixed)
> - adapt existing functions
> - tag types/overloads for easier access

Template for **repetition**: match min to max times.\
base case and some syntactic sugar specializations for it.

```cpp
// base case repetition
template <class T, size_t _min, size_t _max = infinite>
struct repeat;

template <class T>
using pack = repeat<T,0,infinite>;  // parameter pack

template <class T>
using non_empty_pack = repeat<T,1,infinite>;  // non-empty parameter pack

template <class T>
using optional = repeat<T,0,1>;  // optional parameter, with or without default value

template <class T, size_t _times>
using n_times = repeat<T,_times,_times>;  // exactly n-times

using eps = repeat<char,0,0>;  // no parameter(epsilon)
```

**sequenceing**: matching type list one after the other

```cpp
template <typename... Ts>
struct seq;
```

**alternatives**: matching exactly one of the specified type lists

```cpp
template <typename... Ts>
struct alt;
```

**defaulted parameters**: consume argument, if it fits the type, use it, otherwise inject default created argument.

```cpp
// c++ style default parameter, only at the end of the parameter list (trailing)
template <class T, class _Init = impl::default_create<T>>
struct cpp_defaulted_param;

// general default parameter, can appear anywhere in the parameter list (non-trailing)
template <class T, class _Init = impl::default_create<T>>
struct general_defaulted_param;
```

**forwarding references**, either unrestriced or conditional (with SFINAE). quoted function are described in mp11 documentation.

```cpp
// unrestricted perfect forwarding
struct fwd;

// forward reference guarded by meta-function
template<template<class> class _If>
struct fwd_if;

// forward reference guarded by quoted meta-function
template<class _If>
struct fwd_if_q;
```

**template argument deduction**:

- global: all instances deduce the same set of types.
- local: deduced instances are independent of each other.
- mixed: specify which template arguments are matched globally.
  - like sequence of _std::vector/<T,Alloc>_ where T is the same but differnet Allocator.
  - sequence of _std::map_ with the same key-type.

```cpp
// global
template<template<class...> class Patterns>
struct deduce;

// local
template<template<class...> class Patterns>
struct deduce_local;

// mixed
template<GlobalDeductionBindings,template<class...> class Pattern>
struct deduce_mixed;
```

**adapting existing functions**: either simple functions or those with overload-sets, template, or defaulted parameters.

```cpp
template<auto function>
struct adapt;


template<class Adapter>
struct adapted;
// adapt overload-set 'qux'
struct adapt_qux
{
    template <class... Ts>
    static decltype(qux(std::declval<Ts>()...)) forward(Ts&&... ts)
    {
        return qux(std::forward<Ts>(ts)...);
    }
};
```

other stuff:

- tagging types to refer to type by name and not position.
  - tagging overloads
  - tagging type. acessing arguments from a pack
- SFINAE condition on teh whole overload set

The hop grammer for function parameters

```
CppType ::= any(cv-qualified) C++ type

Type ::= CppType | tagged_by<tag,Type>

ArgumentList ::= Argument | ArgumentList, Argument

Argument ::= Type
| repeat<Argument,min,max>
| seq<ArgumentList>
| alt<ArgumentList>
| cpp_defaulted_param<Type,init>
| general_defaulted_param<Type,init>
| fwd
| fwd_if<condition>
| deduce{_local|_mixed}<Pattern>
```

examples for the hop-grammer:

```cpp
// at least one int
non_empty_pack<int>;

// list of double, followed by optional options_t
seq<pack<double>,cpp_defaulted_param<options_t>>;

// optional options_t, followed by a list of double
seq<general_defaulted_param<options_t>,pack<double>>;

// name-value pairs
pack<seq<string, alt<bool, int,double,string>>>;

// local type deduction, a list of vector of arbitrary (different) value types
template <typename T>
using vector_alias = vector<T>const&;
pack<deduce_local<vector_alias>>;

// mixed global/local type deduction: a pack of maps, all having the same key-type
{
    // T1 is bound with globabl_deduction_binding
    template <typename T1,typename T2>
    using map_alias = map<T1,T2>const&;

    // map index in pattern to tag-type
    // all deduced types with that index have to be the same
    using bindings = mp_list<global_deduction_binding<0,tag_map_key_type>>;
    pack<deduce_mixed<bindings,map_alias>>;
}

// only types fitting into a pointer (word size)
template <typename T>
struct is_small : mp_bool<sizeof(remove_cvref_t<T>) <= sizeof(void*)> {};
pack<fwd_if<is_small>>;


// matching all the above as an overload set
using overloads = ol_list<>
    ol<non_empty_pack<int>>,
    ol<pack<double>,cpp_defaulted_param<options_t>>,
    ol<general_defaulted_param<options_t>,pack<double>>,
    ol<pack<seq<string, alt<bool, int,double,string>>>>,
    ol<pack<deduce_local<vector_alias>>>,
    ol<pack<deduce_mixed<bindings,map_alias>>>,
    ol<pack<fwd_if<is_small>>>
    >;
```

of course there can still be ambiguity.

if we want to expand the grammar, we might want to generate type-list for an overload.

for HFVs, we had "repeat the type _arity_ times".

```cpp
mp_repeat_c<mp_list<Types,arity>>;
```

but for hop, we have a combinatorial problem.

> "generatea all parameter-lists that accept _arity_ arguments"

we can do this by recursively travesing the overload, and we have to ensure progress to avoid infinit recursion.

> 1. for each sub-expression of the overload:
>    - analyse the min/max length of the producible type-list.
> 2. recursively traverse the overload:
>    - build types lists, `TL1,TL2,...,TLn` for all sub expressions.
>    - genereate cartesian product `TL1 x ... x TLn` and concatenate tuples.
>    - prohibit infinite recursion by using the min/max info to limit to types-list that can match the call's arity.
>    - avoid repeating type-lists.

generating test function from the type list:

for HFVs, we had to unpack mp_list and expand it as arguments of test

```cpp
template <size_t Index,class... Params>
struct single_function<Index, mp_list<Params...>>
{
    constexpr std::integral_constant<size_t, Index> test(Params...) const;
};
```

> for the hop, we have additional tasks
>
> - templated arguments:
>   - forwarding references, local/global conditions, adapters.
> - templated arguments deduction.
> - removing the tag-types.
>   in the

- the `mp_invoke_q<_If,Args>` checks global conditions of the overload against actual parameter types.
- `decoction_helper` - global template argument deduction
- `unpack`
  - returns parameter types/ geneters types for forwarded parameters.
  - check local condition
  - local template argument deduction
  - removes tag-types

```cpp
template<typename ... Args,
std::enable_if_t<mp_invoke_q<_If,Args...>::value &&
deduction_helper<mp_list<Args...>, mp_list<TypeList ...>>::value,
int* > = nullptr>
constexpr result_t test(typename unpack<typeList,Args>::type...) const;
```

the test function - template argument deduction

in an ideal world it would be be validated against a pack of types,

```cpp
template <template<class...> class... Patterns> //deduction patterns
struct deducer_t
{
    template <template<class...> class... Patterns, class T> //deduction types
    static mp_list<std::true_type,mp_list<T....>> test(Patterns<T...>...);
    static mp_list<std::false_type,mp_list<>> test();
};
```

but in the real world, we have some problems doing that, this code is invalid in C++. `Patterns` might be an alias template, which is not allowed with a pack.

```cpp
//invalid code
template <template<class...> class... Patterns, class T> //deduction types
static mp_list<std::true_type,mp_list<T....>> test(Patterns<T...>...);
```

we overcome that with **Booset.Preprocessor** to generate type-deduction test functions.

some stuff still remain.

### Summary

we can see some demos on compiler explorer

- [int and floats](https://godbolt.org/z/esqc6GMsY)
- [reference-types](https://godbolt.org/z/359Y34b9P)
- [tagging overloads](https://godbolt.org/z/8oT5T7Wb6)
- [default arguments](https://godbolt.org/z/aWPMdb6jG)
- [one entry point to rule them all?](https://godbolt.org/z/)
- [template argument deduction](https://godbolt.org/z/19ddvTfvW)

some statistics about the library, about ~2000 lines of code.
what's missing.

</details>

## Frictionless Allocators - Alisdair Meredith

<details>
<summary>
A suggestion to add language support allocators, the issue, the problems, and how it might look
</summary>

[Frictionless Allocators](https://youtu.be/1PWFQVvaOl4)

### Why Allocators

What is an allocator?

> A service that grants exclusive use of a region of memory to clients.\
> Nice clients will return that region to the service when no longer needed.

Why do we Want allocators? is the `operator new` not good enough? the answer is mostly performance, but also instrumentation and special memory regions.

using allocators get us about 3-5 speed up compared to `new`, and even order of magnitude in extreme cases.

A Faster Allocator:

> - A general purpose allocator tries to minimize _contention_.
>   - a better/replacement for `operator new`
> - Avoid synchronization if we can guarantee all access from a single thread.
> - Simplified bookkeeping if we never reclaim memory.
>   - A _monotonic_ allocator simply advances a pointer through a buffer on each allocation.

Performance benefits:

> Better memory locality improves runtime **after** allocation.
>
> - Keeping memory in L1/L2 cache has an enormous impact on runtime performance.
> - Altough the CPU is trying to manage the cache to make this happen anyway, a local memory pool goes a long way to help.
> - Memory pools minimize the effects of _diffusion_ on a single task.
> - Memory pools on the stack reduce fragmentation of long running processes.
> - No synchronization needed if allocation is confined to a single thread.

Two common strategies to improve locality:

- The first is **allocation on the thread stack**, typically from a pre-sized memory buffer.
- The second way is by **managing a pool of memory** to avoid system calls to the _memory manager_. this is the common implementation of the `operator new`, but having a specific pool for each data structure is better for locality.

custom operators also have additional functionalies besides the main utility of handing out memory.

- Debugging
- Logging
- Profiling
- Test drivers

> **Special Memory**\
> Special memory may be hardware specific, or have some other property, e.g. shared memory. it ofter requires a _handle_ with more info than a native c++ pointer. for example, `boost::interprocess` for shared memory containers.\
> Some architectures provide different access to differnet regions of memory. such as VRAM on video cards.

(introducing his colleague, Emery Berger)

> goals: Accelerating Programs Via Custom Allocators:
>
> - **Demonstrate** value of custom allocation. empirically measure opportunities, performance impact and space impact.
> - Help programmers
>   - **Automatically** identify opportuninites for custom allocation in legacy code.
>   - Provide tools to ensure **efficiency** for programers using custom allocation.

initial results with the **BufferedSequentialAllocator** custom allocator, using three benchmarks,

- _stresstest_ - sort, filter and change data based on random numbers.
- _moda_moda_: single-stage and multi-stage. execl worksheet computations, join views and perfrom computations.

about 25% improvemnt.

### What do allocators look like in C++20?

In c++11, we got [allocator traits](https://en.cppreference.com/w/cpp/memory/allocator_traits). some `using` defintions, we pass the allocator to some other facility. we have customization points at the _construct_ and _destroy_. there was also a probelm of propageting allocators across containers.

```cpp
template <class Alloc>
struct allocator_traits {
	using allocator_type = Alloc;
	using value_type = typename Alloc::value_type;
	using pointer = /*see documentation*/;
	using const_pointer = /*see documentation*/;
	using void_pointer = /*see documentation*/;
	using const_void_pointer = /*see documentation*/;
	using difference_type = /*see documentation*/;
	using size_type = /*see documentation*/;
	//...
	template <class T>
	using rebind_alloc = /*see documentation*/;

	template <class T>
	using rebind_traits = allocator_traits<rebind_alloc<T>>;
	//...
	static pointer allocate(Alloc& a, size_type n);
	static pointer allocate(Alloc& a, size_type n, const_void_pointer hint);
	static void deallocate(Alloc& a, pointer p, size_type n);

	template <class T, class ... Args>
	static void construct(Alloc& a, T* p, Args&&... args);

	template <class T>
	static void destroy(Alloc& a, T* p);

	static size_type max_size(const Alloc& a);
	//...
	using propagate_on_container_copy_assignment = /*see documentation*/;
	using propagate_on_container_move_assignment = /*see documentation*/;
	using propagate_on_container_swap = /*see documentation*/;
	static Alloc select_on_container_copyconstruction(const Alloc& rhs);
};
```

in c++17 things were improved

> - `is_always_equal` - propagated tag.
> - Improved exception specifications on containers.
> - `std::pmr` - the polymorphic memory resources namespace and the containers. supplying allocator (actually, memory resource) at runtime.
> - constraint: fancy pointers must be _contiguous_ iterators.

```cpp
template <class T>
struct allocator {
	using value_type = T;
	using size_type = size_t;
	using difference_type = ptrdiff_t;
	using propagate_on_container_move_assignment = true_type;
	using is_always_equal = true_type;

	allocator() noexcept;
	allocator(const allocator&) noexcept;

	template<class U>
	allocator(const allocator<U>&) noexcept;
	~allocator();
	allocator & operator=(consst allocator&) = default;

	T* allocate(size_t n);
	void deallocate(T* p, size_t n);
};
```

in c++20, `constexpr` support was added to allocators. now we can use custom allocators inside constexpr blocks, even though we can't get them out of the compile time blocks into the runtime.

```cpp
template <class T>
struct allocator {
	using value_type = T;
	using size_type = size_t;
	using difference_type = ptrdiff_t;
	using propagate_on_container_move_assignment = true_type;
	using is_always_equal = true_type;

	constexpr allocator() noexcept;
	constexpr allocator(const allocator&) noexcept;

	template<class U>
	constexpr allocator(const allocator<U>&) noexcept;
	constexpr ~allocator();
	constexpr allocator & operator=(consst allocator&) = default;

	[[nodiscard]] constexpr T* allocate(size_t n);
	constexpr void deallocate(T* p, size_t n);
};
```

**How does `pmr` work?**

> - Resource derived from _std::pmr::memory_resoure_.
>   - Override the pure abstract members.
> - Clients store pointer to memory resoure ("like a vtable")
> - Resource pointer never propagates
>   - Scopre resource object with a longer lifetime than customers.
>   - Data strucuters "guarantee" all elements using the same resource.
> - Alias templates for all standard containers to use tis new scheme.

**memory resource** - public interface for the user, protected interface for derived classes.

```cpp
class memory_resource{
	// For exposition only
	static constexpr size_t max_algin = alignof(max_align_t);
	public:
		virtual ~memory_resource();
		void* allocate(size_t bytes, size_t alignment = max_align);
		void deallocate(void p*,size_t bytes, size_t alignment = max_align);

		bool is_equal(const memory_resource& other) const noexcept;

	protected:
		virtual void* do_allocate(size_t bytes, size_t alignment) = 0;
		virtual void do_deallocate(void p*,size_t bytes, size_t alignment) = 0;

		virtual bool do_is_equal(const memory_resource& other) const noexcept = 0;
};
```

standard resources:

```cpp
memory_resource* new_delete_resource() noexcept; // delegates to the new/delete
memory_resource* null_memory_resource() noexcept; // always throws bad_alloc
memory_resource* get_default_resource() noexcept; // customizable

class monotonic_buffer_resource; //fixed size memory, never released
class synchronized_pool_resource; //static memory pools for different sized memory
class unsynchronized_pool_resource; // when we don't have threeads
```

> **Idiom and usage of pmr**
>
> - Memory resources are objects, typically scoped to a function, with a lifetime longer than their clients below them.
> - Default, object, adn global "allocators"
>   - System-wide default used for all objects unless otherwise specified.
>   - Use the default for function-scope objects unless otherwise specified.
>   - Use object allocator (supplied at construction) only (and always) for data that is part of the objects data structure (that persists beyone the function call).
>   - Use another "global allocator" for any object with static of thread-local sotrage duration, as it may outlive the the default resource after _main_.
> - _no support for fancy pointers_

quick example, potential bug because of RVO, so we force a creation of a copy.

```cpp
pmr::string make_string(const char *s)
{
	pmr::monotonic_resource res;
	pmr::string x(s,&res);
	//return x; //potential RVO - which means leaking
	return {x}; //force creation of a temporary, using the default resource.
}
```

> **Scoped Allocator Model**
>
> - Simple idea: every element in the data structure uses the same allocator(memory-resource).
> - Class design: every ember of the object graph (all bases and members) uses the same allocator.
> - Invariant: all dynamically allocated _persistent_ data uses the same allocator.
> - Key Benfit: underpins performance when looking to avoid diffusion and fragmentation.
> - Important benefit: easy to guarantee allocator/resource has a longer lifetime that it's clients.
> - Implication: **allocators can never propagate**, or else any swap or assignment could invalidate the whole system.

structure of std::pmr::allocator, [documentation](https://en.cppreference.com/w/cpp/memory/polymorphic_allocator). no assignment operator (deleted), c++20 added some behavior.

> **pmr limitations**
>
> - Solves the vocabulary problem, but only if used consistently (i.e., the vocabulary problem! it's still there).
> - No support for fancy pointer / special memory regions.
> - Storing an extra pointer in every object, repeatedly though the whole data structure.
> - Cost of dynamic dispatch.

### What causes friction?

one cause of friction is types that don't have support for the scoped allocator model, and the users didn't provide appropiate constructors.

> - c-style array - `pmr::string data[42]`
> - std::array can't have an allocator. `std::array<pmr::string,42> more_data`
> - lambda objects - capturing and closures.
> - strucutred bindings.
> - deafult member initlizars - which allocator is used?
> - this problem is recursive,such as `std::vector<std::array<string,10>>` - can't provide the allocator on the way down the chain.

Allocator Propagation

> - Allocator is bound at construction.
> - Should the allocator be rebound on assignment?
>   - Assignment copies data.
>   - Allocator is orthogonal, specific to each container object.
> - Traits give contorl of the Propagation strategy (deafult is to never propagate)
> - Propagation is complex
>   - too many fine-grainted traits dimensions to reasonably support.
>   - what does it mean to propagate on _swap_, but not on _move-assignment_? or vice-versa?
>   - the trait for copy construction is actually a function call?

the syntantic overhead is high, writing code through _allocator_traits_ looks like expert level code. the amount of constructors is usually doubled to allow for optional allocator, and this prevents us from having multiple defaulted arguments, because we must have one for the allocator.\
as an example, c++11 _std::unordered_map_ had 8 constructors (and that was when we thought it was allocator aware), but c++17 has 15, with most delegating to the original eight.\
there is also a complexity of argument order, should the allocator be the final argument, or should we pass the allocator arguments and the allocator at the start to avoid template issues.

There is a **copy constructor issue**, copy ellison might lead to a situation where we return an object with a reference to a memory resource that is about to leave scope. this was improved in c++17, and compiler warnings might help in the future, but behavior will be unpredictable.

### What should we do about it?

> **The ideal Model**
>
> - No allocator spam in the interface.
> - A single data structure uses the same allocator throughout.
>   - a container and all its elements.
>   - e.g. a graph, its nodes, their contents, etc...
> - If a type manages dynamic memory, it _always_ supports an allocator.
> - "Allocator aware" types are known to the type system.
>   - it can query if a type is allocator aware.
>   - it can query which allocator an object uses.

this will require us to have a model of allocator awareness: explicitly aware types will be classes marked as such. and implicitly aware types are those which derive from other allocator aware classes or have data members which are allocator aware. awareness is viral from both base class and data members.

> **Allocator Aware Properties**
>
> - The allocator aware class will use the supplied allocators to acquire all memory for the data structure persistent needs.
> - There is a consistent (customizable) API to query which allocator an object is using.
> - The allocator for an object will not change during its lifetime. **i.e. alloctoars do not propagate**.

We need Querying to allow operations that require allocators to be the same, like move and swap. if we want to build an object outside the object (to avoid copies), we will need to create it with the correct allocator.

we don't want to add allocator overloads to every constructor, maybe a new syntax will be needed.

```cpp
multipool_resource res;
set<string>x {"hello","world"} using res; // suggestion, pass the allocator this way
```

### What is our experience?

here is an example, we need to specify the behavior of the move constructor to take a differnet allocator

```cpp
class Object {
		std::pmr::string d_name;

	public:
	using allocator_type = std::pmr::polymorphic_allocator<>;

	explicit Object (allocator_type a = {}): d_name("<UNKNOWN>",a){}
	Object (const Object& rhs,allocator_type a = {}): d_name(rhs.d_name,a){}
	Object(Object&&) = deafult;
	Object (Object&& rhs,allocator_type a): d_name(std::move(rhs.d_name),a){}

	//apply rule of 6
	~Object() = deafult;
	Object& operator=(const Object& rhs) = default;
	Object& operator=(Object&& rhs) = default;
};
```

if the new syntax was implemented, it would look like this: some things could be simplified,

```cpp
class Object {
		std::pmr2::string d_name;

	public:
	//using allocator_type = std::pmr::polymorphic_allocator<>;

	Object (): d_name("<UNKNOWN>"){} //no longer explicit
	Object (const Object& rhs) = default;
	Object(Object&&) = deafult;
	//Object (Object&& rhs,allocator_type a): d_name(std::move(rhs.d_name),a){}

	//apply rule of 6
	~Object() = deafult;
	Object& operator=(const Object& rhs) = default;
	Object& operator=(Object&& rhs) = default;
};
```

and if we give the _d_name_ argument a default initialization, all our constructor are defaulted, so we go down to the rule of zero and we no longer have to write them

```cpp
class Object {
	std::pmr2::string d_name = <"UNKOWN">;
};
pmr::multipool_resource res;
Object X{"Hello world"} using res;
```

> Implementing Awareness
>
> - Stash an allocator pointer as construction, much like vtable pointer.
>   - does not vary through construting a hierarchy.
> - Customization API to give precise control of storage if needed.
>   - e.g. _optional_ object needs to stash allocator when empty, but can re-use to storage for the missing object.
> - If awareness is _implicit_ , access the allocator through the entity granting awareness.
>   - Do not pay to store excess copies of the pointer.
>   - Leaf nodes of data strucutres will always need a pointer.

in most use cases, the awareness will be implicit, so it won't impact common users and will have low implementation costs. implicit awarensscan be supported by c++11 containers, using an allocator-aware allocator object (the allocator template parameter). and implicit awareness implies the need for a new language feature.

> Allocator injection
>
> - Inject an allocator at object creation time, in addition to constructor arguments.
>   - needs language support with an extension syntax.
>   - `new` operator one obvious customization, but need local object support too.
> - An implicit extra argument for every constructor.
>   - No constructor spam with allocator overloads.
>   - process-wide deafult is provided if not supplied by the caller.
>   - note: the move constructor is special
> - implicitly propgagate that injection through member initlizars for all bases and members.
>   - but **not** into constructor body.

Sean Baxter has his own c++20 compiler, where he works on Circle as a post c++ language, and he managed to get a running prototype working quickly. it seems to work fine so far, even if it doesn't have implicit generation of _allocator_type_.

```cpp
// allocator-specifier in initializers and postfix-expressions

#include <list>
#include "logger.hxx"

int main(){
	logging_resource_t logger("logger");
	// using-allocator in a braced initializer for a declaration creates a PMR list when the allocator expresion dervies memory_resource.
	pmr::list<int> my_list{1,2,3} using logger;

	// using-allocator in a braced initializer on an expresion.
	auto my_list2 = pmr::list<int>{4.4,5.5,6.6} using logger;
}
```

this will also cascade into other cases, such a factory functions, like _std::make_shared_ and _std::to_string_. we will either need to pass the allocator as an argument, or add the same syntax support. but we also need to avoid redundant allocations with the the default allocator inside the factory.

and what about the move constructor? can we have a support for a move constructor that uses diffent allocator?

```cpp
template <class T>
struct NamedValue{
	std::string name;
	T value;
	NamedValue(NamedValue &&) = deafult;
	NamedValue(NamedValue &&) using Alloc = deafult;
};
```

> Option 1:\
> Member initializers call extended-move-with-allocator for each base and member. The classes that handle memory allocation directly(i.e. vector, rahter than classes with a vector member) will implement the custom logic to test the allocator, and move or make copies as needed.\
> this supports move-only types as members, as long as they manage any extended-move logic themsevels, which is trivial by deafult for _non_-allocator aware types.
>
> Option 2:\
> Test whether the allocators are compatible. if so delegate to the regular move constructor, if not compatible, delegate to the extended-copy-constructor.\
> if the extended-copy-constructor is not available, the move is ill-formed, and therefore deleted at compile level.\
> The ideal is to provide custom overloads to handle non-default cases.

the correct course is probably option 2.

Internal Pointers example:

```cpp
struct highlight{
	std2::vector<int> d_values;
	int* d_focus; // invarient: points to an element in d_values

	highlight(const highlight &) = delete; //copy constructor
	highlight(highlight &&) = deafult; //move constructor
};
```

copying _d_focus_ would violate the invarient so the copy constructor was delete. an alternative way was to implement a look-up to fix and make the pointer point to the correct coresponding element. in this example, only option 2 is safe.

Dispatching Extended Move

```cpp
struct highlight{
	std2::vector<int> d_values;
	int* d_focus; // invarient: points to an element in d_values

	highlight(const highlight &):d_values(other.d_values), d_focus(nullptr)
	{
		//maybe find a new address for d_focus in d_values.
	} // copy constructor
	highlight(highlight &&) = deafult; //move constructor
	highlight(highlight &&) [[?]] // magical using overload
	: hightlight(copy_or_move(other)) {} // delegating constructor

	private:
		//factory function
		static hightlight copy_or_move(highlight& other){
			if (allocator_of(other) == allocator_of(copy_or_move)){
				return std::move(other); // move
			}
			else {
				return other; //copy
			}
		}
};
```

### What next?

> Basic Feature Set
>
> - Pass allocators to _factory_ function and initializers thorough extra _using_ argument.
> - A special allocator type that imbues enclosing classes as allocator aware
>   - Fundamental type to avoid specifying customization interface.
>   - Acts like a reference to a _pmr::memory_resource_
> - all constructors all allocator aware types are implicitly allocator aware.
>   - move constructor is special and split in two.
> - _allocator_of_ implicit hidden friend function.
> - implicit factory functions
>   - support a _using_ when return type is allocator aware (what will this mean when calling from template?)
>   - implicitly supply allocator to return value
>   - when guaranteed (N)RVO applies, supply allocator to variable declerations.

</details>

## Taking Template Classes Further with Opaque Types & Generic NTTPs - Joel Falcou & Vincent Reverdy

<details>
<summary>
The "kiwaku" library implementing Opaque Types, keyword parameters and Non Type Template Parameters.
</summary>

[Taking Template Classes Further with Opaque Types & Generic NTTPs](https://youtu.be/sFSMzLVpe90), [slides](https://github.com/jfalcou/presentations/blob/conference/cpp-now/taking_template_further.pdf), [kiwaku github](https://github.com/jfalcou/kiwaku), [RABERU (keyword parameters) github](https://github.com/jfalcou/raberu).

NTTP: Non Type Template Parameters

simulations are replacing experimenting.

arrays and matrixes (N-dimension arrays). an nD-array must bbe fast,easy to use and expressive. this means playing well with modern hardware (SIMD,GPU), having proper abstraction, and be intuitive for people with proper domain knowledge (numeric savvy users).

existing solutions:

- View/Container
  - std::vector, std::array
  - std::span,
  - std::mdspan (future proposal)
  - boost.QVMM
  - std::valarray
- Expression Templates
  - Blitz++
  - Eigen
  - NT<sup>2</sup>
  - Armadillo
  - Blaze

however, the solutions aren't adequate, there some concerns

- lazy evaluation
- nD array handliing
- customization protocols
- hardware support

this lecture focuses on nD-array handling and customization.

when we write a library, we want the compiler to make use of compile time optimizations, so we need to get high level informaiton. on the other hand, we want to avoid implementation leaks. ("anything that is exposed in the api, someone will write code that depends on")

examples:

- -1 as a dynamic size tag for _std::span_/_std::mdspan_.
- _Eigen::Matrix<typename Scalar, int Rows, int Cols>_
- passing allocator as type+value instead of pure type.

> | Problem                      | Specification                                                                                                 | Kiwaku Solution                        |
> | ---------------------------- | ------------------------------------------------------------------------------------------------------------- | -------------------------------------- |
> | Runtime Componenets handling | Full runtime should be handled at runtime. No need for type-based specification                               | **Opaque types**                       |
> | Optimizations specifications | Array and view behavior must be trivial to setup. Users should have access to an intuitive option passing API | **Keyword Parameters**                 |
> | Compile-time/Runtime Barrier | Compile-time options must have a rich semantic                                                                | **NTTP - Non Type Template Parametes** |

### Opaque Types

we start with the constructors part:

```cpp
// Dynamic array using the default allocator
kwk::array<float, kwk::_2d>a1(kwk::of_shape(200,200));

// Dynamic array using some other allocator
kwk::array<float, kwk::_2d>a2(kwk::of_shape(10,50), some_allocator{});

// Allocator and data are copied to a1
a1=a2;
```

> Challenges:
>
> - How can we get rid of passing the allocator type as a template parameters?
> - Can we ensure proper copy and move semantics?

some defitions:

> A type is **opaque** if you can't see through it.
>
> - the contents of its implementations are not accessible directly.
> - often implemented using _type-erasure_.
> - if users cant look at one type's internals, there are less oppertunity for abstraction leaks.
>
> State of the Art:
>
> - based on **Sean Parent**'s talk on Polymorphism.
> - Use polymorphism as an implementation detail instead of as a first class property.
> - Provides a full _Regular Type_ interface on top of the polymorphic behavior.
> - Does not require intrusive adaption from user code.

common example ares _FILE\*_ (pointer to file), _std::any_, _std::any_.

For nd-Arrays, the arrays are often large (or even very large), but the allocation is often outside of the critical part, and there a few resizing and growth operations.\
In the kiwaku library, allocators deal with block of _void\*_, this is based on _Alexandescu's Allocator_. The size of the allocation is known, and the allocators can be chained or selected based on arbitrary policies. The allocator must be simple, no CRTP, no polymorphic base class. allocator must be **SemiRegularType**, and are copied along with their tables.

Basic block of memory and simple _malloc_ allocator. there isn't a _virtual_ interface or any complex CRTP-like defintion.

```cpp
struct block
{
  explicit operator bool () const {return length != 0;}

  friend bool operator!=(block const & lhs, block const & rhs) noexcept;
  friend bool operator==(block const & lhs, block const & rhs) noexcept
  {
    return lhs.data == rgs.data && lhs.length == rhs.length;
  }


  void reset() noexcept { *this= block{};}
  void swap (block& other){/*...*/}
  void* data = nullptr; // pointer to the allocated bock of memory
  std::ptrdiff_t length = 0; // Size in bytes of he allocated block of memory
};

struct heap_allocator
{
  [[nodiscard]] block allocate(std::ptrdiff_t n) noexcept
  {
    return n!=0? block{malloc(n), n} : block{nullptr,n};
  }

  void deallocate (block & b) noexcept { if (b.data) free(b.date);}
  void swap (heap_allocators &){}
};
```

> Allocator Design in Kiwaku:\
> the _any_allocator_:
>
> - Users Parent-style polymorphic object design.
> - Distinct from _std::pmr::polymorphic_allocator_ (no _memory_resource_).
> - Provides an associated concept: _kwk::concepts::allocator_.
>
> The Allocator Trifecta:
>
> 1. A virtual API object.
> 1. A template adapter implementing said API.
> 1. A _SemiRegularType_ wrapper.

the allocator concept expands the _std::semiregular_ and _std::swappable_ concepts, and requires _allocate_ and _deallocate_ functions.

```cpp
template <typename A>
concept allocator =
std::semiregular<A> &&
std::swappable<A> &&
requires(A a, block& b, std::ptrdiff_t n)
{
  {a.allocate(n)} -> std::same_as<block>;
  {a.deallocate(b)};
};
```

the _any_allocator_ is the only polymorphic piece of the design, and is an internat type to _kwk::any_allocator_. there is an template adaptor that implements the api.

```cpp
struct api_t
{
  virtual ~api_t(){} //base virtual destructor

  virtual block allocate(std::size_t) =0; // Actual allocator interface
  virtual void deallocate(block &) =0; // Actual deallocator interface

  virtual std::unique_ptr<api_t> clone() const = 0; // Helper for polymorphic copy
};

template<concepts::allocator T> struct model_t final: api_t
{
  model_t = default;
  model_t(const T& t): object(t){}
  model_t(T&& t): object(std::move(t)){}

  block allocate(std::size_t n) override { return object.allocate(n); }
  void allocate(block& b) override { object.deallocate(b);}
  std::unique_ptr<api_t> clone() const override { return std::make_unique<model_t>(object);}

  private:
  T object;
};
```

there is also the _any_allocator_ class, which is a _SemiRegularType_ Wrapper.

```cpp
class any_allocator
{

public:
  any_allocator (any_allocator const & a): data(a.data->clone()){}

  // ... All other obvious special members

  template<typename T> any_allocator(T&& t): data(make_model(std::forward<T>(T))){}

  void swap(any_allocator& other) noexcept {data.swap(other.data);}
  [[nodiscard]] block allocate(std::size_t n) { return data->allocate(n);}
  void deallocate(block & b) { data->deallocate(b);}

private:
  struct api_t {
    /*...*/
  };
  template<concepts::allocator T>struct model_f final: api_t
  {
    /*...*/
  };

  std::unique_ptr<api_t> data;
};
```

in a benchmark test, concrete allocation and opaque allocations are similar.

the advantages is that we managed to remove allocator from the template part, we have a less rigid template API, and this might get benefits from pre-compilation (maybe LTO - link type optimization).

### Key Word Parameters

kiwaku containter constrctors

```cpp
using namespace kwk::literals;

// Dynamic array using the default allocator
kwk::array<float, kwk::_2D> a1(kwk::of_shape(200,200));

// Dynamic array using some other allocator modeling concepts:allocator
kwk::array<float, kwk::_2D> a2("allocator"_kw = some_allocator{}, "shape"_kw = kwk::of_shape(20,20));
a1 = a2;
```

Keyword parameters are a syntax to pass argument to function based on the the parameter name. this is already part of python, C# and other languages. in C++ there are chalannges of name mangeling, which names count and others (see N4172 paper).

in the kiwaku library, this will help to simplify the API. this has be done be using library for keyword parameters (RABERU), keywords aer designed locally as _constexpr_, retriving data from them is done using a lambda as a container, and the keywords can be restricted based on concepts.

defining a keyword, using _rbr::keyword_ as keyword builder,

```cpp
namespace kwk::keyword
{
  // The rbr::keyword inline variable generate a new keyword_type
  inline constexpr auto shape = rbr::keyword<struct shape_option>;

  // The _kw UDL generate a new keyword from the list of character of the string
  inline constexpr auto allocator = "allocator"_kw;

  // Equivelent without UDL
  inline constexpr auto allocator = rbr::keyword<id_<'a','l','l','o','c','a','t','o','r'>>;
};
```

the _keyword_ has a generic assignment operator, returning _linked_value_ constructed from the keyword, which is initilazed with a lambda capturing the value of the parameter. this lambda accepts the keyword and returns the value.\

all keyword/value pairs are gather in a _overload_-like structure, where every _operator()_ of each pair is put back into the interface, so fetching a value is done by calling the overload with the required keyword.

```cpp
some_function(sharep=extent[4][6]);
```

binding a value to a keyword, different behavior for lvalue and rvalue values:

```cpp
template <typename T>
template <typename V>
constexpr auto keyword_type<T>::operator=(V &&v) const noexcept
{
  using type = keyword_type<T>;
  if constexpr(std::is_lvalue_reference_v<V>)
  {
    return linked_value(*this,[&v](type const &)->decltype(auto){return v;});
  }
  else
  {
    return linked_value(*this,[w=std::move(v)](type const &) -> V const & {return w;});
  }
}
```

retriving a value from the keyword is done by aggregating lambda,

```cpp
struct unknown_key{template<typename ... T> unknown_key(T&&....){}};

template<typename ... Ts> struct aggregator: Ts...
{
  constexpr aggregator(Ts && ...t) noexcept TS:(RBR_FWD(t)) ... {}
  using TS::operator()...;
  template <typename K> constexpr auto operator()(keyword_type<K> const &) const noexcept
  {
    // If not found before, return the unkown_key value
    return unknown_key{};
  }
};
```

for convience, there is a _settings_ helper which takes care of type deduction, validation, default values, keyword detection, and so on. in addition, there is _keyword_parameter_ and _match_ that allow proper constraints on functions with keyword parameters, and enable non trivial _requires_ clauses.

```cpp
template <typename P0, typename P1>
auto replicate(P0 p0, P1 p1)
{
  using namespace rbr::literals;
  auto const params = rbr::settings(p0,p1);
  return std::string(params["replication"_kw],params["letter"_kw]);
}

std::cout << replicate("replication"_kw= 9, "letter"_kw='Z') << '\n';
```

or using template Params and have default values, the default value can also be a function

```cpp
template <typename ... Params>
auto replicate(Params ... ps)
{
  using namespace rbr::literals;
  auto const params = rbr::settings(ps);
  return std::string(
    params["replication"_kw | 5], // default 5
    params["letter"_kw | '*'] // default '*'
    );
}

std::cout << replicate("letter"_kw='Z') << '\n';
```

we can also provide concepts and require clauses to make sure the keywords match a patterns (like saying that at least one should have a non default value).

this allows to isolate common use cases from power users concerns without having multiple apis. It future proofs and makes the API less likely to change, while keeping compile costs low with concepts and _if constexpr_.

### Generic NTTP - Non Type Template Parameters

non-type template parameters are:

- an integral type
- an enumeration type
- a pointer type
- a pointer a member type
- _std::nullptr_t_
- a lvalue reference type
- a floating point type (added in c++20)
- literal class type (with some restrictions) (added in c++20)

we can now combine expression templated with NTTPS to create EDSL: Embedded Domain Specific Language. capture arbitrary constexpr expressions as NTTP and process them to generate a proper implementation.

```cpp
template <auto Value> class_type;
template <auto Expression> edsl_compiler;
```

we can use this to define array shapes. while supporting runtime, compile time an combinations of.

> Context:
>
> - Arrays gather data in a n-dimensional grid.
> - The number of effective dimension is supposedly _known at compile time_.
> - The number of elements along each dimension may vary.
> - The number of elements along a given axis may be known at compile time.
> - The inital ordering of those sizes is _domain specific_ and _arbitrary_.

this is used in machine learning.

currrently _std::mdspan_ faces a difficult situation: it can support dynamic types at the cost of very verbose template.

```cpp
using my_span= std::mdspan<double, extents<3,9,7>>;
using other_span= std::mdspan<double, extents<3,std::dynamic_extent,7>>;
```

the prefferable option is to have compile time shape, using unified interface for both static and dynamic.

> the main idea is to design an _extent_ type that only cares about runtime size storage, this type is usable as NTTP. and then helpes are provided to smooth the defintion.

they define a _shaper_ struct that builds the array in a incremental way. there is _shape_ struct that defines the storage - it stores the dynamic dimensions, and has ways to minimize the size when dealing with static (compiler time known) dimensions.

compraison of output code shows that there is no overhead. some compiler explorer demonstration.

</details>

##

[Main](README.md)
