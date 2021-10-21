<!--
ignore these words in spell check for this file
// cSpell:ignore ostringstream ptrdiff_t Filipp Downey Inlines fmodules Andrzej Krzemieński nodiscard Richárd Szalay Levehnstein Loew varargs Dimov arity Alloc cvref deducer
 -->

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
