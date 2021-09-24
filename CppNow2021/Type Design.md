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
