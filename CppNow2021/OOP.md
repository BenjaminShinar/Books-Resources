<!--
ignore these words in spell check for this file
// cSpell:ignore ostringstream affordance O’Dwyer Stringize
 -->

OOP

## Polymorphism À La Carte - Eduardo Madrid & Phil Nas

<details>
<summary>
Thinking about polymorphism, can we move away from virtual dispatch and use duck typing?
</summary>

[Polymorphism À La Carte](https://youtu.be/tn8vlo14FT0), [slides](https://levelofindirection.com/refs/polymorphism.html)

> [À La Carte](https://www.merriam-webster.com/dictionary/%C3%A0%20la%20carte) - "according to a menu or list that prices items separately"

continuation of the 'OO considered harmful' and 'not leaving performance on the jump table' from cpp-con 2020 and 'type erasing the pains of runtime polymorphism" from c++ London

different ways of doing OO, virtual inheritance- method dispatch (C++) vs message passing (smalltalk, objective C). virtual dispatch at dynamic runtime is just another passage from the virtual table (vtable)of the calling object.

message passing is passing a message to an object, the object uses the selector + args to call the correct method, but there are some added benefits:

some demonstaration of objective c, apparently catch can also test objective c code, it has null-object pattern baked in, if we call anything on null, we get the null value, zero or nill. we can have 'fully typed' decleration and dynamically typed ("id") classes. we can then do runtime reflections and graft methods onto the class so we can get 'duck typing'. we decide that we want our class to fail either at compile time (the class doesn't contain the method) or at runtime (doesn't contian the method and nobody added this method to it).

> Runtime polymorphism
>
> - "We want _substitution_ as in Liskov's Substition Principle, this implies _subtyping_. Subtyping is not the same thing as _subclassing_
> - something that "quacks", if a dog can learn to quack, great. we don't care about all the formal relations.

using referential semantics cause performance costs and all sorts of other pain points.

> Beyond Inheritance + Virtual
>
> - Sean Parent's [Inheritance is the Base Class of Evil](https://youtu.be/bIhUE5uUFOA) lecture from 2013.
> - Referential semantics
>   - Indirections
>   - Allocation
>   - Incentive to share state
>   - Non-local reasoning
> - Intrusiveness
> - also: the subclassing relation is too strong:
>   - Supports only monophyletic relations, but para- and polyphyletic relation do exist as well
> - There are many ways to do message passing

monophyletic, paraphyletic, polyphyletic relations - terms from taxonomy

type erasure _Affordance_

[The Space of Design Choices for std::function](https://quuxplusone.github.io/blog/2019/03/27/design-space-for-std-function/) by Arthur O’Dwyer. runtime polymorphism using std::function and callable objects, picking and choosing how we want stuff to be designed, storage, ownership, const and non-const version.

_zoo::AnyContainer_ allows tools for affordance.

### Duck Typing

python has duck typing.

- if it waddles like a duck, swims like a duck and quacks like a duck, we shouldn't care that it isn't really a duck, we can use it as if it was simply a duck.
- and example of "Stringize" generic duck typing function, leaverage `operator<<` and `to_string`, having a bridge for the runtime for overload resolution and template specialization, doesn't require changing types.
- an ad-hoc approach.
- a _zoo::AnyContainer_ with some affordance can serve as a basis for another continaers with more affordance.
  - Typical example: A container that is move-only (does not require copyability) is extended with the affordance of copyability.
  - This extensibility applies to containers that themselves are extended for _zoo::AnyContainer_, such as _zoo::Function_

support composability.

```cpp
using NormalPolicy =
    zoo::Policy<
    void *[2], //size of two pointers
    zoo::Destroy, zoo::Move, zoo::RTTI>; //normal destructibility, move operations, RTTI behavior
using TypeErasureProvider = zoo::AnyContainer<NormalPolicy>;
using OrderConsumer = zoo::Function<TypeErasureProvider, CallSignature>; //template adapter

//adding copying capabilities.
using CopyableOC = zoo::AnyContainer<
        zoo::DerivedVTablePolicy<OrderConsumer, zoo::Copy>>;
```

[example](https://godbolt.org/z/xKsq3f) in compiler explorer

Faking and mocking using duck-typing in objective c. We simply declare a class with the required operations, we don't need the whole chain of interfaces or inheritances,exchaning and injection methods from singletons ("method swizzling").

### Message Passign in Rust and Swift

all sorts of important people in those langauges teams,

Protocol oriented programming in swift. abstractions, models, etc... (swift example of protocols, extensions) \
Also a Rust example with traits.

</details>
