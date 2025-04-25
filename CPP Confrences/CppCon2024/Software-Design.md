<!--
// cSpell:ignore relocatability Björn Fahller Blarg multiplicator soooo maat
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Software Design

<summary>
17 Talks about Software Design.
</summary>

- [x] 10 Problems Large Companies Have with Managing C++ Dependencies and How to Solve Them - Augustin Popa
- [x] Adventures with Legacy Codebases: Tales of Incremental Improvement - Roth Michaels
- [ ] C++ Under the Hood: Internal Class Mechanisms - Chris Ryan
- [ ] Dependency Injection in C++ : A Practical Guide - Peter Muldoon
- [ ] Design Patterns - The Most Common Misconceptions (2 of N) - Klaus Iglberger
- [ ] Embracing an Adversarial Mindset for C++ Security - Amanda Rousseau
- [ ] Hiding your Implementation Details is Not So Simple - Amir Kirsh
- [ ] High-performance Cross-platform Architecture: C++ 20 Innovations - Noah Stein
- [ ] How Meta Made Debugging Async Code Easier with Coroutines and Senders - Ian Petersen, Jessica Wong
- [ ] Modern C++ Error Handling - Phil Nash
- [ ] Monadic Operations in Modern C++: A Practical Approach - Vitaly Fanaskov
- [ ] Newer Isn't Always Better, Investigating Legacy Design Trends and Their Modern Replacements - Katherine Rocha
- [ ] Ranges++: Are Output Range Adaptors the Next Iteration of C++ Ranges? - Daisy Hollman
- [ ] Reflection Is Not Contemplation - Andrei Alexandrescu
- [x] Relocation: Blazing Fast Save And Restore, Then More! - Eduardo Madrid
- [x] Reusable code, reusable data structures - Sebastian Theophil
- [x] The Most Important Design Guideline is Testability - Jody Hagins

---

### 10 Problems Large Companies Have Managing C++ Dependencies and How to Solve Them - Augustin Popa

<details>
<summary>
Problems with large dependencies and how to solve them.
</summary>

[10 Problems Large Companies Have Managing C++ Dependencies and How to Solve Them](https://youtu.be/kOW74IUH7IA?si=9R-_FSdv3w1QKKTI), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/10_Problems_Large_Companies_Have_with_Managing_Cpp_Dependencies_and_How_to_Solve_Them.pdf).

package libraries and dependencies management is still a top pain point for C++ developers.

- Conan
- vcpkg
- NuGet

> Preview – 10 conclusions
>
> 1. Support building from source
> 2. Build a verified binary cache
> 3. Version using baselines
> 4. Build open-source with a package manager
> 5. Cache build assets internally
> 6. Monitor, prevent, and respond to vulnerabilities
> 7. Centralize common tasks
> 8. Produce SBOMs
> 9. Global, reproducible builds
> 10. Break large migrations down into smaller milestones

#### Problem 1: ABI incompatible C++ binaries

binaries aren't portable, they depend on how they were compiled. so it leads to a very complex dependency graph, different compilers, target OS, feature flags, build configurations, etc...

the consumer needs to get the correct binary and use it, but a solution for this is to build the dependant packages from source, some package managers support this, and we can also use the CI build chain to build the dependencies. some companies moved to Monorepo architecture, where everything is built together.

this also allows for editing and tweaking the source code for the dependencies.

#### Problem 2: Build times are too long when building from source

Of course, building packages takes time, it scales poorly, as each package has it's own decencies.\
we can use a caching strategy with a shared cache for binaries with a unique hash identifier based on versions and ABI.

#### Problem 3: Version conflicts – The Diamond problem

each decency has dependencies of it's own, and if two packages consumer the same third package, they might get a version conflict.

a solution could be to move from package/libraries into "baselines", a baseline contains a set of packages with compatible versions between them.

#### Problem 4: Building open-source dependencies is hard

in-house development is more work to write, but easier to integrate, maintain and protects against legal concerns and some vulnerabilities.

a package manager can handle some of the problems with building the packages, and bridge the gap.

#### Problem 5: Organization restricts access to online downloads

organizations don't want to allow open-source code without a process, and there's always a risk that the open-source library goes down.\
we can get around this by maintain a cache (proxy) of 3rd party packages source code, which can also be moved to a protected environment that won't allow internet connection.

#### Problem 6: Security vulnerabilities in open-source code

Open source packages might have known vulnerabilities. the organization needs a strategy to monitor, review and respond to them.

#### Problem 7: Duplicated engineering cost to maintain dependencies

managing dependencies gets harder with scale, there are problems with communication, and there are times where similar work is being done in different teams.

#### Problem 8: Difficult to track or report on all dependencies

audit dependencies, build the dependency graph, producing SBOMs - software bill of materials.

#### Problem 9: Build toolchain variations across the organization

reproducible builds, tests and deployments. building in containers, establishing company-wide policies and interior tools storage.

#### Problem 10: Moving to new solution is complex or too time consuming

takes time, effort, causes un-expected problems.

we can set milestones - small changes that represent a gain in value, even if there aren't any more changes to the process. we identify the action and what benefit it provides us.

---

> Summary – 10 conclusions
>
> 1. Build C++ dependencies from source
> 2. Establish binary caching where possible
> 3. Organize dependencies + versions into baselines (fixed points in time)
> 4. Use an open-source package manager to save time / effort
> 5. Create an asset cache for sources needed to build dependencies
> 6. Develop a vulnerability monitoring, prevention, and response strategy and associated tools / workflows
> 7. Centralize common dependency management tasks, enforce consistency across organization at scale
> 8. Organize dependencies into coherent packages and start producing SBOMs
> 9. Establish a global toolchain policy and build in containers if possible
> 10. Break large migrations down into smaller milestones with a win at each step.

</details>

### Reusable Code, Reusable Data Structures - Sebastian Theophil

<details>
<summary>
(not really sure about this talk)
</summary>

[Reusable Code, Reusable Data Structures](https://youtu.be/5zkDeiyF5Rc?si=tmsS-hGBMTMCey1r), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Reusable_Code_Reusable_Data_Structures.pdf)

> DRY -- don't repeat yourself -- is an important software engineering principle.\
> Repetitive code is error-prone code. Inevitably, and sooner rather than later, we will forget to change one of these repetitive code locations. C++ offers many different tools to share code and data and I often see novice and intermediate programmers struggle with choosing the best one in each situation.\
> We have template functions, template classes, std::variant, virtual classes and std::any. We have some common associated programming patterns like CRTP, templated base classes, template functions taking function arguments.\
> All of these have their uses and in my talk, I want to develop some intuitions on when to use which.

what should we use and what should we not use.\
starting with an example of centering elements in html.

for code to be reused in an efficient manner, we need to generalize the problem, understand which algorithm can help us, etc...

when we have a shared algorithm and different data, we use generic functions and concepts.

looking at some patterns: output iterators, predicate, and at the command design pattern.

> GENERIC FUNCTIONS
>
> - Need clear requirements
> - Expressed in well-known concepts
> - Reduce dependencies
> - Flexible and customizable through customization points

Generic classes, data is being templated, either in the base class or with the curiously recurring template pattern (CRTP).

mixin classes, adding functionality to un-related classes.

other options are inheritance (virtual functions) and <cpp>std::variant</cpp>, but should they really be used for code sharing? runtime polymorphism ties things together, it creates coupling and has lasting implication.

- heap allocation
- lifetime issues
- ownership

external polymorphism, global overloads, <cpp>std::variant</cpp> and <cpp>std::visit</cpp>.

</details>

### C++ Relocation - How to Achieve Blazing Fast Save and Restore and More! - Eduardo Madrid

<details>
<summary>
Moving and Storing Data for better performance.
</summary>

[C++ Relocation - How to Achieve Blazing Fast Save and Restore and More!](https://youtu.be/LnGrrfBMotA?si=RQ6R3iWcMSCTZzTK), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Relocation.pdf), [event](https://cppcon2024.sched.com/event/1gZeM/relocation-blazing-fast-save-and-restore-then-more). [Data Orientation For The Win! - Eduardo Madrid - CppCon 2021](https://www.youtube.com/watch?v=QbffGSgsCcQ)

The opposite of pointe chasing.

> What's what we mean to improve?
>
> - Our systems need objects to refer to each other, and we need runtime polymorphism
> - Nothing bad per se, just that it craters performance
> - What if we are willing to do a lot of effort to improve performance lot?
> - One option is to apply **Data Orientation** techniques in general.
> - Today, we give extra attention to "relocatability"

Relocatability is how we apply data orientation techniques to make our data easy to be moved around, so it will reside in memory when it has the best performance.\
(relocatability is not about position independent code for dynamically linked libraries).

Writing code the "normal" way means we don't focus on mapping data in memory.

- storing data in classes or structs with heterogenous members.
- using pointers for objects to refer to one another.
- using virtual inheritance and overrides for runtime polymorphism
- indirections, allocations, lifetime, sharing ownership (<cpp>std::unique_ptr</cpp>, <cpp>std::shared_ptr</cpp>)

#### Data Orientation

The non-trivial way to write code takes data to memory mapping into account.

> Columnar Representation (Scattering):\
> Transforming collections of structs of members of heterogeneous typesinto one structure with arrays of homogeneous data types (going from what in databases is called a row-representation to a columnar representation) also called "scattering" or achieving structures of arrays.

That's a long way for saying "struct of arrays" over "array of structs", or moving from row-wise representations to column-wise.\
With *row-wise* representations, we have structs containing different kinds of members, pointers (which can be null), data that is differentiating the state of the object and we have containers with multiple elements of the class.\
with *column-wise*, we combine the data and the container together, we move to homogenous collection of data, all fields of the same type are stored continuously. to access a single 'item', we use a handle (Facade/Proxy design pattern).\
After scattering, we are using "entity-component systems", we no longer have objects, instead we have entities, and members are replaced with compnotents. entities and components are bound through an index in a global structure.

```cpp
// row-wise
struct Item {
  int id;
  double price;
  Specials *specials; // pointer, can be null
  Category *category;
  bool is_homemade; // messes with the data alignment
};
std::array<Item, N> items;

// column-wise
template <std::size_t N>
struct State {
  std::array<int, N> ids;
  std::array<double, N> prices;
  EfficientMap<std::size_t, Specials> specials;
  std::array<CategoryHandle, N> categories;
  EfficientSet<std::size_t> homemade;
};

auto state = State<N>;
class ItemHandle {
  std::size index;
public:
  int id() { return state.ids[index]; }
  double price() { return state.prices[index]; }
  Specials *specials() { return state.specials[id()]; }
  CategoryHandle category() { return state.categories[id()]; }
  bool &is_homemade() { return state.homemade.contains(id()); }
};
```

Allocating data homogeneously allows for better performance (SIMD, CUDA), reduces allocations, it dis-incentives communication between state change. also, moving away from using pointers allows more efficient memory dump. data is given different addresses by the allocator, so we can't restore the application correctly (the pointers are no longer valid). if we have this kind of a layout, the relations between entities are preserved through the index.

#### Allocators And Relocatability

Allocators don't care about the efficient allocation of the data, so we must be able to control and move the data, and we need the ability to maintain the allocation across the lifetime of the program, as objects are created and destroyed.

we start by re-introducing some indirection - using stable indices.

> Recapping Björn Fahller ["Cache Friendly Data + Functional + Ranges = &heartsuit;"](https://www.youtube.com/watch?v=XJzs4kC9d-Y) @C++ On Sea 2024:
>
> - He calls stable indices "stable ids", and the technique is to use an extra indirection that pays for itself in terms of performance.
> - Furthermore, you can have both the forward mapping and the backward mapping, depending on your needs.
> - This indirection is what provides the freedom to relocate!
> - Björn Fahller presented benchmarks that dove into details of wall-clock performance, cache misses, IMO his results are representative of what would be real world use.

Relocatability makes saving the state of the application trivial - just dump the data as is. and restoring it is just as easy. if the data is ever "shaken" and becomes not efficient, we can re-allocate them to make them efficient again. "shaken" data can be fragmentation (empty holes in the array), or having old and new entities stored together (if we re-used the old memory to store new data).\
Memory performance relies on time and space locality, we want to have objects that are accessed together stored together, relocating objects allows us to maintain and improve this.

One issue is using *Value Mangers*, using <cpp>std::optional</cpp>, which currently doesn't take allocators, but can take a <cpp>std::vector</cpp> type with an custom allocator. this has something to do with type erasure (<cpp>std::any</cpp>, <cpp>std::function</cpp>). the language doesn't currently have support to access internals (???).

</details>

### The Most Important C++ Design Guideline is Testability - Jody Hagins

<details>
<summary>
A different view on testability
</summary>

[The Most Important C++ Design Guideline is Testability](https://youtu.be/kgE8v5M1Eoo?si=5w4dnwmZakqDf0Yv), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/The_Most_Important_Design_Guideline_is_Testability.pdf), [event](https://cppcon2024.sched.com/event/1gZfl/the-most-important-design-guideline-is-testability).

> Scott Meyers has famously proclaimed that the most important general design guideline is to make interfaces easy to use correctly and hard to use incorrectly.  I don't dispute that this is one of the most important design guidelines.\
> However, in my close to 40 years of fighting in the C++ trenches, I'd argue that testability is by far the more important design guideline, and antecedent to both ease of use and performance (a particular C++ penchant).\
> In this talk, we will discuss what testability means, and why it is so important.  We will briefly discuss some popular testing techniques, but most of our time will be spent looking into items of testability that are rarely discussed, but are extremely important in practice.

Design interfaces for testability, have an holistic thought about systems and how they can be tested.

> Why is Testing Important? We are human!
>
> - Making Mistakes: Every human is really, really good at making mistakes
> - Being annoyed when OTHER people make mistakes

Testability isn't just Unit Tests and TDD (Test Driven Development).

#### Hidden State

side tangent an example about having a hidden state and how it complicates testing. the best way is to having composability, but there's an old pattern to provide access to all private members.

doing some tricks with template specializations... this shouldn't be done. but it's possible.

```cpp
template <typename TagT, auto...>
struct Sudo;

template <typename TagT>
struct Sudo<TagT>
{
  inline static typename TagT::type value{};
};

template <typename TagT, typename TagT::type Member>
struct Sudo<TagT, Member>
{
  inline static const decltype(Member) value = Sudo<TagT>::value = Member;
};

template <typename TagT, typename ObjectT>
decltype(auto)
sudo(ObjectT && object)
{
  assert(Sudo<TagT>::value);
  return std::invoke(  Sudo<TagT>::value,  std::forward<ObjectT>(object));
}

template <typename ClassT, typename MemberT, typename = void>
struct SudoTag
{
  using type = MemberT ClassT::*;
};

// usage
namespace foo::bar { class Blarg { int x = 42; }; }

using Blarg_x = SudoTag<foo::bar::Blarg, int>;
template struct Sudo<Blarg_x, &foo::bar::Blarg::x>;

int main()
{
  auto blarg = foo::bar::Blarg{};
  std::cout << sudo<Blarg_x>(blarg) << '\n'; // print blarg.x
}

// with macro - less boilerplate
  #define SUDO(NAMESPACE, TYPE, CLASS, MEMTYPE, MEMBER) \
  namespace NAMESPACE { \
  struct TYPE \
  : SudoTag<CLASS, MEMTYPE> \
  { }; \
  } \
  template struct Sudo<NAMESPACE::TYPE, &CLASS::MEMBER>
// Then, you can use it like this...
BP_SUDO(foo::bar, Blarg_x, foo::bar::Blarg, int, x);
```

#### Back to Testability

Donald Knuth - offered bounties for bugs. but didn't think highly of unit tests and other testing methodologies. has some quotes that go against what we usually think as "best practice".

- property based testing
- fuzzing

Avoid naked types - everything is a special type of it's own. never confuse the parameters and the order of the arguments. always have strong types.

#### Case Study - Knight Capital

> everything is an API.

a company that had an issue in 2012 which lost the company 10 million dollars per minute for 45 minutes. they did some stock trading and things went bad. they had unused code still in the system, and they wanted to reuse a flag that was for the old system for the new one. the code was changed as part of refactoring but wasn't tested, eventually, the new version of the code was deployed to all but one machine. when the application went live, the one computer run the old code, which was called because of the repurposed flag, and then old code did the wrong thing, and because it wasn't tested properly, it behaved in a way that isn't part of normal procedure.

this was a chain of errors, if any one of them didn't happen, this would be a non-event.

> - Remove old code and add new code at the same time
> - Reuse an existing "value" for completely different functionality
> - Stopped using Power Peg in 2003, dead code still around
> - Refactored code in 2005 ; moved dead code ; test still pass
> - Manual deployment ; no review ; no tests
> - Different components had different view of same thing
> - email for important log messages
> - Rollback made matters worse

(how to test each of these points)

</details>

### Design Patterns - The Most Common Misconceptions (2 of N) - Klaus Iglberger

<details>
<summary>
CRTP, variant and design patterns.
</summary>

[Design Patterns - The Most Common Misconceptions (2 of N)](https://youtu.be/pmdwAf6hCWg?si=Ihpx8evhi69KwtMK), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Design_Patterns.pdf), [event](https://cppcon2024.sched.com/event/1gZfn/design-patterns-the-most-common-misconceptions-2-of-n).

in the previous talk in "Meeting C++ 2023", the topics were:

- Builder
- Factory Method
- Bridge
- Design Pattern

this talk will be about design patterns and virtual patterns

#### CRTP - Curiously Recursive Template Pattern

the base class is templated on the derived class. it allows us to sometimes remove virtual function calls.

```cpp
template<typename Derived>
class Animal
{
private:
  Animal() = default;
  ~Animal() = default;
  friend Derived;
public:
  void make_sound() const {
    static_cast<Derived const&>(*this).make_sound()
  }
};

class Sheep: public<Sheep>
{
  // not defining the destructor
   void make_sound() const { std::cout << "baa"; }
};
```

there are two limitations: there is no shared base class, each derived class has a different base class, so we can't have a collection of objects. but since everything is a template, this forces everything that comes into touch with it becomes a template itself. this moves our code from the source files to the headers.

in C++23, we can replace CRTP with explicit object parameter ("deducing this")

```cpp
struct NumericalFunctions
{
  void scale(this auto&& self, double multiplicator)
  {
    self.setValue( self.getValue() * multiplicator );
}
};

struct Sensitivity : public NumericalFunctions
{
  double getValue() const { return value; }
  void setValue(double v) { value = v; }
  double value;
};

int main()
{
  Sensitivity s{1.2} ;
  s.scale(2.0);
  std::println(std::cout, "s.getValue() = {}", s.getValue());
}
```

however, we can't replace everything, we can't deduce the derived class from the base class. we only deduce the static type, not the dynamic one. CRTP has two behaviors: adding functionality, and creating static interfaces. we can replace CRTP with explicit object parameters when we add functionality, but not when we used CRTP for static interfaces.\
Since the term is ambagious and refers to two things, we need terms to differentiate them.

> CRTP for Static Interfaces
>
> - provides a base class for a related set/family of types.
> - defines a common interface.
> - is used via the base class interface.
> - introduces an abstraction and is a design pattern.
> - should be called **Static Interface**
>
> Adding Functionality via CRTP
>
> - provides implementation details for the derived class.
> - does not define a common interface.
> - is not used via the base class interface.
> - does not introduce an abstraction, hence is no design pattern.
> - should be called **Mixin**.

c++20 concepts might seem similar to static interfaces, but concepts are "duck typing", it is much more permissive, static interfaces are "opt in". we could get around this by adding a requirement of <cpp>std::derived_from\<T,AnimalTag\></cpp> to the concept to restrict the set of types.

```cpp
template< typename T >
concept Animal = requires(T animal) {
  animal.make_sound();
};

template<Animal T>
void print(T const& animal) {
  animal.make_sound();
}
class Sheep {
public:
  void make_sound() const { std::cout << "baa"; }
};
int main() {
  Sheep sheep;
  print(sheep);
}
```

#### std::variant

The usual example of drawing shapes, starting with the classic OOP solution and adding a draw strategy object to avoid coupling with an implementation. we want to focus on separating code that changes often (low level implementation) from code that doesn't change that much.\
This draws from the Gang of Four approach of using inheritance and virtual functions, which is something the C++ community tries to move away from.

we can implement this with <cpp>std::variant</cpp>

```cpp
class Circle
{
public:
 explicit Circle( double rad )
 : radius{ rad }
 , // ... Remaining data members
 {}
 double getRadius() const noexcept;
 // ... getCenter(), getRotation(), ...
private:
 double radius;
 // ... Remaining data members
};
class Square
{
public:
 explicit Square( double s )
 : side{ s }
 , // ... Remaining data members
 {}
 double getSide() const noexcept;
 // ... getCenter(), getRotation(), ...
private:
 double side;
 // ... Remaining data members
}
using Shape = std::variant<Circle, Square>;
using Shapes = std::vector<Shape>;

class ShapesFactory
{
public:
Shapes create( std::string_view filename )
{
  Shapes shapes{};
  std::string shape{};
  std::ifstream shape_file{ filename };
  while( shape_file >> shape ) {
    if( shape == "circle" ) {
      double radius;
      shape_file >> radius;
      shapes.emplace_back( Circle{radius} );
    }
    else if( shape == "square" ) {
      double side;
      shape_file >> side;
      shapes.emplace_back( Square{side} );
    }
    else {
      break;
    }
  }
    return shapes;
  }
};

using Factory = std::variant<ShapesFactory>;

using Drawer = std::variant<OpenGLDrawer>;
void drawAllShapes(Shapes const& shapes, Drawer drawer) {
  for(auto const& shape : shapes){
    std::visit([](auto d, auto s){ d(s); }, drawer, shape);
  }
}

// more code
```

> This solution is soooo much better:
>
> - No inheritance, but a functional approach
> - No (smart) pointers, but values
> - Proper management of graphics code
> - Automatic, elegant life-time management
> - Less code to write
> - Soooo much simpler
> - Better performance

(some benchmarks)

however, there is a problem. our design mixes level of implementation and responsibilities. rather than having two levels of architecture, we only have one thing. but maybe templates can help us? this might turn us into a templated nightmare...

<cpp>std::variant</cpp> isn't a replacement to virtual functions. variants support a fixed set of types and an open set of operations, while virtual functions have an open set of types and a closed set of operations. a variant is a visitor design pattern.
</details>

### Adventures with Legacy Codebases: Tales of Incremental Improvement - Roth Michaels

<details>
<summary>
Incrementally improving legacy code base.
</summary>

[Adventures with Legacy Codebases: Tales of Incremental Improvement](https://youtu.be/lN-dd-0PjRg?si=oytVnTjGIEqPG-5-), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Adventures_With_Legacy_Codebases.pdf), [event](https://cppcon2024.sched.com/event/1gZft/adventures-with-legacy-codebases-tales-of-incremental-improvement).

> What is legacy code?
>
> - No tests
> - Lot's of code
> - Very old
> - Authors may be gone
> - Many C++ standards
> - New/old styles
> - New/old paradigms
> - Bad decisions from the past that once made sense
> - Possibly rewritten by a less skilled engineer

[code-maat tool](https://github.com/adamtornhill/code-maat) to analyze code change from version control.

not adding tests for code that won't change. focus on user/business value.

linting with style guides (clang format), should be enforced by machine, not by human feedback. this messes with `git blame` in the codebase. using githooks to apply formatting on incoming changes and verifying it as part of the ci pipeline.\
using static analysis: address, undefined behavior and thread-safety analyzers. also make this part of the pipeline, especially important when code should work on both AMD and ARM processors.\
testing pyramid, moving up from unit test to integration test and up to End-to-End tests.

#### Should we change legacy APIs?

transitioning from custom home-made objects to external libraries to the standard library, like <cpp>std::optional</cpp>. we should avoid breaking changes.

"Live at Head" - all internal libraries are always built from scratch, no internal versions, problems are detected immediately, not months afterwards when the new version is pulled.

handling dependencies, using `clang-tidy` custom rules, `clang-tidy-diff` to only report on changed lines.

#### Sharing Legacy Code

Desire to standardize, reducing duplicate code, share unique code, have a single strategy for behavior (one installer, one UI library). share code you are proud of and think is good and useful.

</details>
