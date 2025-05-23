<!--
ignore these words in spell check for this file
// cSpell:ignore Alinsky
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Education & Coaching

### Reviewing Beginners' C++ Code - Patrice Roy

<details>
<summary>
Giving feedback in teaching environment and in code reviews.
</summary>

[Reviewing Beginners' C++ Code](https://youtu.be/9dMvkiw_-IQ), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-2022-Patrice-Roy-Reviewing-Beginners-Code.pdf)

how to get and give feedback on our code. it's hard to absorb too much information at once, receiving criticism can feel like being attacked.

#### What?

we need to know what we are searching for in a feedback session, we can't have a complete list, but we can prepare some points.

- what is important to me as someone who provides feedback?
- what points would most benefit the receiver of the feedback?
- what would be most beneficial for the company?
- what are the essential parts of the code that must be addressed?
- what do I know about the person who writes the code? what are his strengths and blind spots?
- what are the constraints of the system?
- what are the company guidelines that must be followed?

#### How?

how to give feedback in a constructive manner. how to receive feedback?. it's important to choose what to focus on. feedback in written form is different from verbal feedback. writing can seem dry and without empathy.

having a list of points which we expect to see and address.

- comments - missing, misleading
- constants, magic numbers
- plain dangerous code
- language - spelling errors
- indentations
- modularization - classes, functions
- things that just won't work
- things which aren't part of the task
- naming issues
- types

feedback must be given in a way that the receiver will understand.

#### Some Ideas

expose the receiver to good code, suggest other parts of the code base to learn from. share articles, videos, books and cool tricks that can enrich the receiver.\
Look at how other people provide feedback, and how they teach. try exchanging roles with others, ask for their opinion. some people prefer smaller groups and one on one talks.

**Don't give someone a task they can't complete**

</details>

### Back to Basics: The C++ Core Guidelines - Rainer Grimm

<details>
<summary>
A different approach to understanding the core guidelines.
</summary>

[Back to Basics: The C++ Core Guidelines](https://youtu.be/UONLB7wBVSc), [slides](https://www.modernescpp.org/wp-content/uploads/2022/09/TheCCoreGuidelines.pdf).

> Best Practices for the Usage of C++: \
> Why do we need guidelines?
>
> - C++ is a complex language in a complex domain.
> - A new C++ standard is published every three years.
> - C++ is used in safety-critical systems.

reflecting on the code habits and on old code.

- MISRA - Motor Industry Software Reliability Association, also used in avionic and medical domain (c++03)
- AUTUSAR - automotive domain (c++14)
- C++ Core guidelines - community driven

the guidelines are about 350 rules and hundreds of pages long, just like other seminal books, each rule has a complex structure:

- the rule
- reference number
- reasons
- examples
- alternatives
- exceptions
- enforcements
- see also
- notes
- discussions

it's a reference books, but doesn't fit casual reading.

a different way to describe the guidelines is to divide them into higher order sections.

#### Philosophy

MetaRules for the concrete rules

> - Express intent and ideas directly in code.
> - Write in ISO Standard C++ and use support libraries and supporting tools.
> - A program should be statically type safe. When this is not possible, catch run time errors early.
> - Don't waste resources such as space or time.
> - Encapsulate messy constructs behind a stable interface.

#### Interfaces

> Interfaces should
>
> - be explicit
> - be strongly typed
> - have a low number of arguments
> - separate similar arguments

avoid misusing, help guid users to correct behavior.

#### Functions

> Distinguish between in, in/out, and out parameter

- in - `func(x)` (cheap), `func(const x&)` (costly)
- in & retain ("copy") - same
- in & move from - `func(x&&)`
- in/out - `func(x)`
- out - `x func()`, `func(x&)` (when expensive to move)

Ownership semantic of function parameters.

| Example                 | Ownership Semantic                |
| ----------------------- | --------------------------------- |
| `func(x)` - value       | independent owner of the resource |
| `func(x*)` - pointer    | borrowed the resource             |
| `func(x&)` - reference  | borrowed the resource             |
| `func(std::unique_ptr)` | independent owner of the resource |
| `func(std::shared_ptr)` | shared owner of the resource      |

#### Classes

> Class hierarchies organize related classes into hierarchical structures.
>
> class versus struct
>
> - Use a class if it has an invariant
> - Establish the invariant in a constructor

Concrete Types

> A concrete type (value type) is not part of a type hierarchy. It can be created on the stack. A concrete type should be regular.

- default constructor - `X()`
- copy constructor - `X(const &X)`
- copy assignment `operator=(const X&)`
- move constructor - `X(X&&)`
- move assignments `operator=(X&&)`
- destructor - `~X()`
- swap operator - `swap(X&,X&)`
- equality operator - `operator==(const X&)`

> The Big Six
>
> - The compiler can generate them
> - You can request a special member function via `default`
> - You can delete a automatically generated function via `delete`
> - Define all of them or none of them (rule of six or rule of zero)
> - Define them consistently
> - There are strong dependencies between the big six

Constructor

> Don’t define a default constructor that only initializes data members; use member initialization instead.

```cpp
struct Widget {
  Widget() = default;
  Widget(int w): width(w) {}
private:
  int width = 640;
};
```

single element constructors should be marked as `explicit` to avoid implicit conversation.

```cpp
class MyClass{
public:
  explicit MyClass(A){} // converting constructor
  explicit operator B(){} // converting operator
};
```

example of integer promotion from the boolean conversions.

> Destructors
>
> - Define a destructor if a class needs an explicit action at object destruction
> - A base class destructor should either be **public and virtual**, or **protected and non-virtual**
> - public and virtual - You can destroy instances of derived classes through a base class pointer or reference
> - protected and non-virtual - You cannot destroy instances of derived classes through a base class pointer or reference
> - Destructors should not fail make them `noexcept`

#### Enumerations

> Enumerations are used to define sets of integer values and also a type for such sets of values.
>
> - Use enumerations to represent sets of related named constants
> - Prefer enum classes over "plain" enums
> - Specify enumerator values only when necessary

strongly typed enums are better.

#### Resource Managements

> RAII stands for Resource Acquisition Is Initialization.
>
> - Key idea:
> - Create a local guard object for your resource.
> - The constructor of the guard acquires the resource and the destructor of the guard releases the resource.
> - The C++ run time manages the lifetime of the guard and, therefore, of the resource.
>
> Implementations
>
> - Containers of the STL
> - Smart pointers
> - Locks
> - std::jthread

#### Expressions and Statements

- good names are the most important for good software.
- avoid mixing signed and unsigned integers in arithmetics (int turns into unsigned and goes ballistic)

```cpp
#include <iostream>
int main() {
  int x = -3;
  unsigned int y = 7;
  std::cout << x - y << '\n'; // 4294967286
  std::cout << x + y << '\n'; // 4
  std::cout << x * y << '\n'; // 4294967275
  std::cout << x / y << '\n'; // 613566756
}
```

#### Performance

> Wrong optimization
>
> - “premature optimization is the root of all evil” (Donald Knuth)
>
> Rule for optimization
>
> - Measure with real-world data
> - Version-ize your performance test

having versions of the performance test for each compiler/platform. be able to show the decision process that lead to choosing.

> Importance of measuring
>
> - Which part of the program is the bottleneck?
> - How fast is good enough for the user?
> - How fast could the program potentially be?

use move sematics if possible, `constexpr`, give the compiler additional hints (`noexcept`, `final`).
local and simple code can be optimized easily and the compiler can help us.

an example is sorting with a predicate, a local lambda can be optimized by the compiler in cases where a free function cannot.

#### Concurrency

> Threads
>
> - Prefer `std::jthread` to `std::thread`
> - Don’t detach a thread
> - Pass small amounts of data between threads by value
> - To share ownership between unrelated threads use `std::shared_ptr`

don't reference local variables from threads (they can change or go out of scope), use tools such as sanitizers and code analysis. `gcc -sanitize=thread -g`

#### Error Handling

> Error handling consists of
>
> - Detect the error
> - Transmit information about an error to some handler code
> - Preserve the valid state of a program
> - Avoid resource leaks

#### Constants and Immutability

> By default, make objects immutable
>
> - Cannot be a victim of a data race
> - Guarantee that they are initialized in a thread-safe way
> - Distinguish between physical and logical constness of an object
> - Casting away `const` from an original `const` object is undefined behavior if you modify it

the `mutable` keyword on member variable mean that it can be changed inside `const` member functions.

#### Templates

> Use
>
> - Use templates to express algorithms that apply to many argument types
>
> Interfaces
>
> - Use function objects (lambdas) to pass operations to algorithms.
> - Let the compiler deduce the template arguments.
> - Template arguments should be at least SemiRegular or Regular.

regular = supports big six operations + equals + swap.

#### The Standard Library

> Prefer `std::array` and `std::vector` to a C-array

if the size is know at compile time and small, use `std::array`, if not known or size is large, use `std::vector`.

#### Guidelines Support Library

a header only library that allows us to validate that our code conforms to the guidelines.

</details>

### Rules for Radical Cpp Engineers - Improve Your C++ Code, Team, & Organization - David Sankel

<details>
<summary>
Lessons from political activism on how to get things to happen.
</summary>

[Rules for Radical Cpp Engineers - Improve Your C++ Code, Team, & Organization](https://youtu.be/ady2mUIQpt4)

we want things like:

- Good Engineering Practices
- Coding Standards
- Design Reviews
- Cleanup
- Real Code Reviews
- Planning
- Education
- Development Infrastructure

but how do we make them happen? how do we change our "culture" to have them. so we turn to **Political Activism** to find inspiration and learn from them how to create changes. one of the ,_Saul Alinsky_, wrote a book called "Rules for Radicals: A Pragmatic Primer for Realistic Radicals", we also look at the works of _Ella Jo Baker_, which focused on grassroots movements.

In a sprint, there can be:

- Issue of the week - something that is highly prioritized by the management, and introduced out of schedule.
- "Save the company" - a critical bug fix or production issue that must be addressed.
- Yearly major initiative - the big planned change of the year.
- Finishing previous year initiative - actually finishing the thing we said we did last year.
- "Fantasy Zone" - stuff that we want to have, but won't happen.

if a software engineer has a suggestion, it might be pushed into the "fantasy zone" backlog, and will never be implemented. this is similar to what Alinsky called _"pointless sure-loser confrontation"_.

Alinsky's principals of communication:

- Communicate within the experience of your audience
- Work the system
- Be patient and pragmatic

> "Men don't like to step abruptly out f the security of familiar experience; they need a bridge to cross from their own experience to a new way." \
> ~ Saul Alinsky

(example of communicating between a software engineer and the manager using those principals)

#### The Rules

1. Speak their language
2. Get organized
   1. utilize existing structures
   2. build when structures are needed
   3. embrace democratic process (straw polls to determine general attitude towards something)
3. Be pragmatic
   1. be realistic, make "demands" that can be met.
4. Evolve the message
   1. have a coherent group goal that encompasses different interests together
   2. a team that does one thing can and should do something next
5. Build leadership
   1. > "Strong people don't need strong leaders" ~ Ella Jo Baker
   2. demonstrate a win early
   3. leaders should be able to leave

---

These rules apply to a team, a department, a company and to the entire c++ community. The C++ language needs to change, and it can change.

</details>

##

[Main](README.md)
