<!--
// cSpell:ignore relocatability Björn Fahller Blarg multiplicator soooo maat Monostate runge_kutta4 TPOIASI Hyrum
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Software Design

<summary>
17 Talks about Software Design.
</summary>

- [x] 10 Problems Large Companies Have with Managing C++ Dependencies and How to Solve Them - Augustin Popa
- [x] Adventures with Legacy Codebases: Tales of Incremental Improvement - Roth Michaels
- [x] C++ Under the Hood: Internal Class Mechanisms - Chris Ryan
- [x] Dependency Injection in C++ : A Practical Guide - Peter Muldoon
- [x] Design Patterns - The Most Common Misconceptions (2 of N) - Klaus Iglberger
- [ ] Embracing an Adversarial Mindset for C++ Security - Amanda Rousseau
- [x] Hiding your Implementation Details is Not So Simple - Amir Kirsh
- [ ] High-performance Cross-platform Architecture: C++ 20 Innovations - Noah Stein
- [ ] How Meta Made Debugging Async Code Easier with Coroutines and Senders - Ian Petersen, Jessica Wong
- [x] Modern C++ Error Handling - Phil Nash
- [x] Monadic Operations in Modern C++: A Practical Approach - Vitaly Fanaskov
- [x] Newer Isn't Always Better, Investigating Legacy Design Trends and Their Modern Replacements - Katherine Rocha
- [x] Ranges++: Are Output Range Adaptors the Next Iteration of C++ Ranges? - Daisy Hollman
- [x] Reflection Is Not Contemplation - Andrei Alexandrescu
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

### Newer Isn't Always Better, Investigating Legacy Design Trends and Their Modern Replacements - Katherine Rocha

<details>
<summary>
Comparing Old and Modern Software Patterns.
</summary>

[Newer Isn't Always Better, Investigating Legacy Design Trends and Their Modern Replacements](https://youtu.be/ffz4oTMGh5E?si=9XLcGF9nVeLUJoiy), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Newer_Isn%E2%80%99t_Always_Better.pdf), [event](https://cppcon2024.sched.com/event/1gZhF/newer-isnt-always-better-investigating-legacy-design-trends-and-their-modern-replacements).

Two camps of people: those who look at existing patterns and thing we should keep them, and those who insist on shaking up how things work. we need to find a middle way of modernizing software without sacrificing what works. we look at trends (software patterns), both old and new, and investigate if they hold up.

#### Examples: Old Pattern vs New Pattern

first example: comparing index based loops to newer ranged loops.

```cpp
std::vector<std::string> vec {"sun", "earth", "moon", "jupiter"};

// old style
for (auto i = 0uz; i < vec.size(); ++i)
{
  std::cout << vec[i] << "\n";
  std::cout << vec.at(i) << "\n";
}

// new style
for (const auto& object : vec)
{
  std::cout << object << "\n";
}
```

> Original Trend – Index For Loop:
>
> - Provides an index that can be used to access an element
> - Index can be used for in loop side effects/calculations
> - Doesn't require a group of items
> - More dangerous access operations
>
> New Trend – Range Based For Loop:
>
> - More data oriented
> - More readable

next example is Global Interfaces vs Global State.\
the original trend is using a singleton to hold state, with the new trend of "Monostate" and dependency injection (dependency inversion principle). we can have global interfaces (logging, I/O, resource management, graphics) or global data (initial parameters, state parameters).

> Original Trend - Singleton:
>
> - Hold one copy of global data/interface and allow others access
> - Usually accessed through a getInstance() or Instance() function
> - Easily accessed
> - Identifiable
> - Hard to test
> - Quintessentially Overused
>
> New Trend – Monostate:
>
> - Make every object in the class static
> - Multiple objects all with the same value
> - Easy to transition to multiple objects
> - May not work well to replace interface singletons
>
> New Trend – Dependency Injection:
>
> - Not a global object
> - Injects the dependency into each of the using objects

```cpp
template <typename T>
class Singleton {
public:
  static T& getInstance()
  {
    static T instance;
    return instance;
  }
private:
  Singleton();
};

class Plotting {
public:
  void plot(double x, double y)
  {
  // ...
  }
};

using PlottingSingleton = Singleton<Plotting>;

// monostate

class Plotting {
public:
  void plot(double x, double y)
  {
  // ...
  }
private:
  static std::queue plottingQueue;
};

// dependency injection

class Plotting {
public:
  void plot(double x, double y)
  {
  // ...
  }
};

class Gps {
public:
  Gps(Plotting& plotter);
  void setPlotter(Plotting& plotter);
  void getPositionVelocityAcceleration(Plotting& plotter);
private:
  Plotting& plotter;
};
```

a C++ specific example is SFINAE vs c++20 <cpp>concepts</cpp>. both constrain templated code (function, classes) and break during compilation time rather than runtime.

> Original Trend – Substitution Failure is Not an Error (SFINAE):
>
> - Substitution Failure Is Not An Error
> - Constraints on templates
> - Known for difficult to read errors
> - Difficult to constrain
>
> New Trend – Concepts:
>
> - Compile Time constraints
> - Named set of requirements
> - Improved compiler errors
> - Easier to create custom constraints for

```cpp
// SFINAE
#include <boost/type_traits/has_operator.hpp>
template <typename Time,
  typename OutputType,
  typename = std::enable_if_t<std::is_arithmetic_v<Time>>,
  typename = std::enable_if_t<std::is_arithmetic_v<OutputType> ||
  (
    boost::has_multiplies<OutputType, Time>::value && 
    boost::has_plus<OutputType>::value
  )>,
  bool = true>
inline constexpr OutputType runge_kutta4_sfinae(std::function<OutputType(Time, OutputType)> fun,
Time time,
OutputType y0,
Time timestep)
{
  auto k1 = fun(time, y0);
  auto k2 = fun(time + timestep * 0.5, y0 + k1 * timestep * 0.5);
  auto k3 = fun(time + timestep * 0.5, y0 + k2 * timestep * 0.5);
  auto k4 = fun(time + timestep, y0 + k3 * timestep);
  return (y0 + (k1 + 2 * k2 + 2 * k3 + k4) * timestep / 6);
}

// concepts
template<typename T>
concept arithmetic = std::integral<T> || std::floating_point<T>;

template<class T, typename Num>
concept add_multiply = requires(T t, Num num) {
  t * num;
  t + t;
};

template <arithmetic Time, typename OutputType>
requires (add_multiply<OutputType, Time>)
inline constexpr OutputType runge_kutta4_concepts(std::function<OutputType(Time, OutputType)> fun,
Time time,
OutputType y0,
Time timestep)
{
  auto k1 = fun(time, y0);
  auto k2 = fun(time + timestep * 0.5, y0 + k1 * timestep * 0.5);
  auto k3 = fun(time + timestep * 0.5, y0 + k2 * timestep * 0.5);
  auto k4 = fun(time + timestep, y0 + k3 * timestep);
  return (y0 + (k1 + 2 * k2 + 2 * k3 + k4) * timestep / 6);
}
```

concepts help simplify the requirements and make the compiler errors easier to read.

Polymorphism also has different patterns, the old pattern being virtual functions, and the new trend being CRTP and the C++23 explicit object parameter ("Deducing This"). polymorphism is one of the key concepts of object oriented programming, we define a general interface (abstraction) and allow concrete implementations for different classes.

> Original Trend – Virtual Functions:
>
> - Run-Time Polymorphism
> - Quintessential Object Oriented Method
> - Overused
>
> New Trend – Curiously Recurring Template Pattern (CRTP):
>
> - Compile Time Polymorphism
> - Force a Downcast from the Parent to Access Child Elements
> - Explicit Cast
>
> New Trend – Explicit Object Parameter/Deducing This:
>
> - C++23 Feature
> - Simplifies Compile Time Polymorphism

```cpp
// virtual functions
struct NetworkConnection
{
  virtual void initializeConfig() = 0; // Pure Virtual
  void init()
  {
    initializeConfig();
    // ...
  };
};

struct Tcp : public NetworkConnection
{
  void initializeConfig() override
  {
    // ...
  }
};

struct Udp : public NetworkConnection
{
  void initializeConfig() override
  {
    // ...
  }
};

// CRTP

template <class Derived>
class NetworkConnection
{
public:
  void init()
  {
    (static_cast<Derived*>(this))->initializeConfig();
    // ...
  };
};

class Tcp : public NetworkConnection<Tcp>
{
public:
  void initializeConfig()
  {
    // ...
  }
};

class Udp : public NetworkConnection<Udp>
{
public:
  void initializeConfig()
  {
    // ...
  }
};

// Explicit Object Parameter

struct NetworkConnection
{
public:
  void init(this auto&& self)
  {
    self.initializeConfig();
    // ...
  };
};

class Tcp : public NetworkConnection
{
public:
  void initializeConfig()
  {
    // ...
  }
};

class Udp : public NetworkConnection
{
public:
  void initializeConfig()
  {
    // ...
  }
};
```

examples of multi-level inheritance over at compiler explorer

- [virtual function](https://godbolt.org/z/T51xE5qbK)
- [CRTP](https://godbolt.org/z/s3ed4Yorv)
- [Implicit Object Parameter](https://godbolt.org/z/ccsoaf3ec)

in the CRTP example, it's not trivial to get the same behavior as the virtual function example, we would need to add some templating magic to overcome it.

#### Conclusion

> Other Potential Evaluations:
>
> - Union vs Variant
> - Enum vs Enum Class
> - Raw Pointers vs Reference vs Smart Pointers
> - Raw Iterators vs Standard Algorithms
> - C-Style Casts vs Fancy Casts (static, dynamic, reinterpret, const casts)
> - Allocators vs <cpp>PMR</cpp>
> - <cpp>printf</cpp> vs <cpp>std::cout</cpp> vs <cpp>libfmt</cpp>
> - Object Oriented Programming vs Functional Programming vs Data-Oriented Design

we see that in some cases, newer patterns aren't strictly better than old patterns, we need to evaluate alternatives for each use case.
</details>

### Reflection Is Not Contemplation - Andrei Alexandrescu

<details>
<summary>
Reflection and it's relation to identity.
</summary>

[Reflection Is Not Contemplation](https://youtu.be/H3IdVM4xoCU?si=S-RHuAIiHrERG_4B), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Reflection_Is_Not_Contemplation.pdf), [event](https://cppcon2024.sched.com/event/1gZhJ/reflection-is-not-contemplation).

> - Static reflection without code generation is incomplete
> - The "reading" part of reflection generally agreed upon
> - The "generation" part of reflection suffered of neglect
>   - P2996 very gingerly sneaks in a foot in the door (define_class)
>   - P3294 finally blows the door off its hinges
> - The two facets of reflection are equally important
> - Where do AI tools fit within this craze?
>
> The Reflection Circularity Problem - Without generation, we're chasing our tails

reflection basics:

> - Recall `^^x` reflects on `x`, `[:y:]` un-reflects (splices) `y`; `[:^^x:]` is `x`
> - Large consensus on introspection query: contemplation is great
> - Fear and loathing about code generation
>   - Expansion of existing introspection objects deemed acceptable
> - Consequence:
>   - Severely limited: can't do 3D with 2D abilities
>   - No consensus on "how much" generation is just enough
>   - No clarity on necessary primitives

OOP and templates were movements towards reflection, with OOP allowing us to call function that haven't been written yet (abstraction), and templates allow us to design with types that weren't written yet.

> Reflection: "I can customize types that will not have been meant for customization".

we want to use reflection to create classes, in the following example, we would want to define a class with data member, who have getters and setters.

```cpp
// reflection
struct Book {
  consteval { 
    property("author", ^^std::string);
    property("title", ^^std::string); 
  }
};

// Equivalent hand-written code
struct Book {
  std::string m_author;   
  const std::string& get_author() const { return m_author; }
  void set_author(std::string s) { m_author = std::move(s); }
  std::string m_title;
  const std::string& get_title() const { return m_title; }
  void set_title(std::string s) { m_title = std::move(s); }
};
```

we create identifiers from the string, and added some stuff them ("m_", "get_" and "_set"), it's crucial that translating strings into identifiers will be fluid, easy, and accurate. identifiers might bet introduced and then used in the generated code.

(splicing is the act of generated code)

> Comparison of Splicing Models
>
> - **The Spec API**: function calls create a "spec" of a type that is spliced
>   - Complex; problematic; death of a thousand cuts
> - **The CodeReckons API**: OOP interface for AST building
>   - Verbose, loquacious, garrulous, prolix, long-winded, circumlocutory, and unceasingly inclined toward linguistic superfluity.
> - **String Injection**: offer a primitive that splices compile time strings into code
>   - Horribly unstructured. Also… can't use macros?!?
> - **Fragments**: C++ fragments of stylized code
>   - Thorough early checking makes complexity explode

#### Token Sequences - P3294

> "There's a representation of C++ code that is well-defined, complete, and easy to understand: C++ code."\
> ~ Daveed Vandevoorde

Creating C++ code from C++ code.

> - At this point complexity is a huge liability
> - Adding yet another sublanguage to C++ deemed undesirable
> - String-based generation best; let's eliminate its disadvantages
> - Strings opaque/unwieldy? Use token sequences instead of strings
>   - Cost: one added literal kind
> - Injection risks and dangers? Restrict string expansion
>   - Carefully controlled escapes, interpolation-style

```cpp
constexpr auto t1 = ^^{ a + /* hi! */ b }; // three tokens - comment is removed
static_assert(std::is_same_v<decltype(t1), std::meta::info>);
constexpr auto t2 = ^^{ a += ( }; // partial expression, 3 tokens again
constexpr auto t3 = ^^{ abc { def }; // Error, unpaired brace
```

a new kind of literal - <cpp>token sequences</cpp> (or <cpp>token strings</cpp>)

> - Escapes inside a token sequence:
>   - `\( expr )` evaluates expr eagerly during creation of the token sequence
>   - `\id( e1, e2, ... )` concatenates strings and integrals, creates an identifier
>   - `\tokens( expr )` expands another token sequence
> - Inside any <cpp>consteval</cpp> function:
>   - `queue_injection( tokens_expr )` injects a token sequence into the current declaration context

code comparisons:

```cpp
// property Using The Spec API
template <class T> using getter_type = auto() const -> T const&;
template <class T> using setter_type = auto(T const&) -> void;
consteval auto property(string_view name, meta::info type) -> void {
  auto member = inject(
    data_member_spec{.name=std::format("m_{}", name), .type=type}
    );
  inject(
    function_member_spec{
      .name=std::format("get_{}", name),
      .signature=substitute(^getter_type, {^type}),
      .body=defer(member, ^[]<std::meta::info M>(auto const& self) -> auto const& {
          return self.[:M:];
        }
      )
    }
  );
  inject(
    function_member_spec{
      .name=std::format("set_{}", name),
      .signature=substitute(^setter_type, {^type}),
      .body=defer(member, ^[]<std::meta::info M>(auto& self, typename [:type_of(M):] const& x) -> void {
          self.[:M:] = x;
        }
      )
    }
  );
}

// property Using The CodeReckons API
object_type(mp, make_const(decl_of(b)));
return_type(mp, make_lvalue_reference(make_const(type)));
append_method(
  b,
  identifier{("get_" + name).c_str()},
  mp,
  [member_name](method_builder& b){
    append_return(b,
      make_field_expr(make_deref_expr(make_this_expr(b)), member_name)
    );
  }
);
method_prototype mp1;
append_parameter(
  mp1,
  "x",
  make_lvalue_reference(make_const(type))
);
object_type(mp1, decl_of(b));
return_type(mp1, ^void);
append_method(
  b,
  identifier{("set_" + name).c_str()},
  mp1,
  [member_name](method_builder& b){
    append_expr(
      b,
      make_operator_expr(
        operator_kind::assign,
        make_field_expr(make_deref_expr(make_this_expr(b)), member_name),
        make_decl_ref_expr(parameters(decl_of(b))[1])
      )
    );
  }
 );
}

// property Using Token Sequences

consteval void property(std::meta::info type, std::string_view name) {
  queue_injection(^^{
    [:\(type):] \id("m_", name); // define the private member
    [:\(type):] const& \id("get_", name)() const {
      return \id("m_", name);
    } // define the getter
    void \id("set_", name)([:\(type):] x) {
      \id("m_", name) = std::move(x);
    } // define the setter
  });
}
```

#### Implementing Identity

> "You can observe a lot by just implementing identity"\
> ~ Yogi Berra

Identity defines fungibility, comparision of "same object" vs "a copy", it also affects aliasing, self-assignment, self-move, double deletion and other stuff. an identity function can be inserted at any sub-expression without changing the meaning. it is a fundamental concept of computer science, this is the reason javascript has both the `==` and the `===` operators.

there are iterations of C++ identity functions in C++, which need to deal with rvalue, lvalue, references and other weird edge-cases. <cpp>decltype</cpp>, <cpp>decltype(auto)</cpp> and <cpp>std::forward</cpp> are heavily used.

```cpp
// naive implementation
template <typename T>
T&& identity(T&& x) { return x; }

const int& a = identity(42); // dangling
int&& b = identity(42); // dangling
auto&& c = identity(42); // dangling

// better - two templates
template <class T>
T& identity(T& x) { return x; } // returns reference
template <class T>
T identity(T&& x) { return std::move(x); } // returns a value

// even better - universal reference
template <class T>
T identity(T&& x) { return T(std::forward<T>(x));} // might call a constructor, or a noop cast

// C++14 style
template <class T>
decltype(auto) identity(T&& x) {return T(std::forward<T>(x));}

template <class A, class B>
decltype(auto) min(A&& a, B&& b) {
  static_assert( /* ... no dangerous comparisons ...*/ );
  return b < a ? B(std::forward<B>(b)) : A(std::forward<A>(a));
}
```

[implementing `min` with identity example](godbolt.org/z/bGPnjM76e)

#### Identity And Reflection

Reflection is the process of deconstructing and reconstructing an entity back into an identical one. once we have this ability, everything becomes easier.

> - Herb keynote's interface example?
>   - Deconstruct type, clone inserting `virtual` and `=0` for each member function
>   - Complain if you find data members or other suspect items
> - polymorphic_base:
>   - Deconstruct type, reassemble making sure no copying allowed
>   - Ensure the destructor is public/virtual or protected/nonvirtual
> - ordered:
>   - Deconstruct type, reassemble and add <cpp>operator<=></cpp>
> - The point is (re)assembly from small, replaceable pieces

it's fairly easy to clone a class, but a harder challenge is cloning a class template. the reflection must preserve the order, introspect function templates, signature constraints (<cpp>noexcept</cpp> clauses, attributes), inner classes (<cpp>iterator</cpp>), and template function signatures (<cpp>std::erase</cpp>), and there are probably many other pitfalls we aren't aware off.

#### Can AI Help?

can generative AI do this? could we simply ask AI to write the code we want instead of reflection? AI models are the stronger than ever, but are also the dumber than they'll ever be in the future. it can write code faster than any human, but not as fast as compilation. AI can also have different results every time, rather than determistic output which is what we need.
</details>

### Modern C++ Error Handling - Phil Nash

<details>
<summary>
Errors in C++, past, present and what might be in the future.
</summary>

[Modern C++ Error Handling](https://youtu.be/n1sJtsjbkKo?si=lqV65wNXHvYFwLD_), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Modern_Cpp_Error_Handling.pdf), [event](https://cppcon2024.sched.com/event/1gZgs/modern-c++-error-handling).

#### Evolution of Error Handling

errors can be disappointments, things that we know are possible, we know they are coming.\
we start with a simple numbers parser. we don't have error handling.

```cpp
int parse_int(std::string_view number) {
  acc = 0;
  for(char c : number) {
    if(c < '0' || c > '9') // TODO: +, -, digit separators?
      return acc;
    acc *= 10;
    acc += c-'0';
  }
  return acc;
}
std::println("{}", parse_int("42")); // 42
std::println("{}", parse_int("42x")); // 42
std::println("{}", parse_int("x42")); // 0
```

let's test if the number contains non-numeric characters

```cpp
[[nodiscard]] bool is_int(std::string_view number) {
  for(char c : number) {
    if(c < '0' || c > '9')
      return false;
  }
  return true;
}
void test(std::string_view str) {
  if(is_int(str))
    std::println("{} is an int", str);
  else
    std::println("{} is not an int", str);
}

test("42"); // 42 is an int
test("42x"); // 42x is not an int
test("x42"); // x42 is not an int
```

we don't know if trailing characters are allowed, so let's externalize that information and let the user decide.

```cpp
enum class ParseStatus {
 Numeric,
 StartsNumerically,
 NonNumeric
};

[[nodiscard]] ParseStatus is_int(std::string_view number) {
  bool first_digit = true;
  for(char c : number) {
    if(c < '0' || c > '9') {
      return first_digit ? ParseStatus::NonNumeric : ParseStatus::StartsNumerically;
    }
  first_digit = false;
  }
  return ParseStatus::Numeric;
}

void test(std::string_view str) {
  switch(is_int(str)) {
    case ParseStatus::Numeric:
      std::println("{} is an int", str);
    break;
    case ParseStatus::StartsNumerically:
      std::println("{} starts with an int", str);
    break;
    case ParseStatus::NonNumeric:
      std::println("{} is not an int", str);
    break;
  }
}

test("42"); // 42 is an int
test("42x"); // 42x starts with an int
test("x42"); // x42 is not an int
```

we can use an out parameter for the result and the return value for the status. or set a global flag for errors (<cpp>errno</cpp>, like classic C). a C++ approach might be to throw exception, it's not popular, but they are good for performance in the happy path, and they push the error handling back onto the caller. but it's really falling out of style. they have affect on performance in the error path, and are non-determistic.

C++ 17 added the <cpp>std::optional</cpp> class, which can signal success and errors, but doesn't distinguish between different errors.

```cpp
std::optional<int> parse_int(std::string_view number) {
  int out = 0;
  for(char c : number) {
    if(c < '0' || c > '9')
      return {};
    out *= 10;
    out += c-'0';
  }
  return out;
}
void test(std::string_view str) {
  if(auto oi = parse_int(str))
    std::println("{}", *oi);
  else
    std::println("Error");
}
```

C++23 has <cpp>std::expected</cpp>, which can either return the value, or an unexpected type, so we can contain the error inside it.

```cpp
std::expected<int, std::runtime_error>
parse_int(std::string_view number) {
  int out = 0;
  bool first_digit = true;
  for(char c : number) {
    if(c < '0' || c > '9') {
      return first_digit ?
        std::unexpected(std::runtime_error("passed string is not a number")) :
        std::unexpected(std::runtime_error("passed string has non-numeric digits"));
    }
    out *= 10;
    out += c-'0';
    first_digit = false;
  }
  return out;
}
 
void test(std::string_view str) {
  if(auto ei = parse_int(str))
    std::println("{}",*ei);
  else
    std::println("Error: {}", ei.error().what());
}
```

we have a problem with composability, we might have nested uses of <cpp>std::expected</cpp> which we need to convert between. we can overcome it with c++23 monadic operations: <cpp>.transform()</cpp>, <cpp>.and_then()</cpp>, <cpp>.or_else()</cpp> and <cpp>.transform_error()</cpp> (only for <cpp>std::expected</cpp>).

```cpp
void test(std::string_view str) {
  auto ef = parse_int(str)
  .transform([](int i) { return i-1; } )
  .and_then([](int i) -> std::expected<float, std::runtime_error> { 
      if(i != 0)
        return 1.f / i;
      else
        return std::unexpected(std::runtime_error("Divide by zero"));
    })
  .transform_error([](auto const& ex) { 
      return std::string(ex.what()); 
    });
  if(ef)
    std::println("{}",*ef);
  else
    std::println("Error: {}", ef.error());
}
```

#### Avoid Disappointments

another example, getting the name of the month based on the number.

```cpp
std::string month_names[] = {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"};

std::string month_as_string(int month) {
  return month_names[month-1];
}

std::println("Month: {}", month_as_string(1)); // JAN
std::println("Month: {}", month_as_string(0)); // <nothing>
std::println("Month: {}", month_as_string(10000)); // garbage
```

lets fix this to constrain the values and throw an exception.

```cpp
std::string month_as_string(int month) {
  if(month < 1 || month > 12)
    throw std::out_of_range("month must be between 1 and 12");
  return month_names[month-1];
}

std::string month_as_string(int month) {
  assert(month >= 1 && month <= 12); // static assert
  return month_names[month-1];
}
```

we could have two versions of this, one for values we've previously checked and we know are valid, and one for un-vetted arguments. we could do all kinds of stuff, like templates and <cpp>constexpr</cpp> with enums.

we don't want all that, we want our code and types to be *correct by construction*, the class can't be instantiated with bad values.

```cpp
struct unchecked{}; 

class Month {
  int month;
public:
  static bool is_valid(int month) {
    return month >= 1 && month <= 12;
  }
  Month(int month): month(month) {
    if(!is_valid(month))
      throw std::out_of_range("month must be between 1 and 12");
  }
  Month(int month, unchecked): month(month) {
    assert(is_valid(month));
  }
  int value() const { return month; }

};

std::string month_as_string(Month month) {
  return month_names[month.value()-1]; // can be member function
}
```

we still don't have contracts in C++, but it's an upcoming feature. we crate a macro to replicate the behavior.

```cpp
#ifndef NDEBUG
#define pre(expr) do{ if(!(expr)){ violation_handler({#expr}); } } while(false)
#else
#define pre(expr) do{} while(false)
#endif

bool is_valid_month(int month) {
  return month >= 1 && month <= 12;
}
void test(int month) {
  std::println("{}", month_as_string(month));
}

std::string month_as_string(int month) {
  pre(is_valid_month(month));
  return month_names[month-1];
}
```

a violation handler is a function pointer that we can assign to.

```cpp
struct ViolationInfo{
  std::string_view expr;
  std::source_location loc = std::source_location::current();
  std::string make_message() const {
    return std::format("{}:{}:{}: precondition failed: ({}): in function: {}",
    loc.file_name(), loc.line(), loc.column(), expr, loc.function_name() );
  }
};

void default_violation_handler(ViolationInfo const& info) {
  std::print("{}", info.make_message());
  std::abort();
}
auto violation_handler = &default_violation_handler;
```

we can replace the default behavior to a "throwing" violation handler, there is some problem with marking function as <cpp>noexcept</cpp>, since the violation handler might throw.\
there is a principal called "Lakos Rule" which calls for conservative use of <cpp>noexcept</cpp> in the standard library. it distinguishes between wide contracts (no precondidtos) and narrow contracts (have precondition), and functions with narrow preconditions should be marked <cpp>noexcept</cpp>.

(some stuff about testing).
</details>

### Monadic Operations in Modern C++: A Practical Approach - Vitaly Fanaskov

<details>
<summary>
Using Monadic Operations
</summary>

[Monadic Operations in Modern C++: A Practical Approach](https://youtu.be/Ely9_5M7sCo?si=NxPmgOeDnlktlKXW), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Monadic_Operations_in_Modern_Cpp.pdf), [event](https://cppcon2024.sched.com/event/1gZgZ/monadic-operations-in-modern-c++-a-practical-approach).

> Agenda
>
> - Briefly about expected and optional
> - Common use cases of expected
> - Monadic operations in software development
> - Tips and tricks

<cpp>std::optional</cpp> is a container that either has or doesn't have a value, it can be used as a return value or a parameter, and in the upcoming C++26 standard library, it would be <cpp>view</cpp> that either contains one value (if the optional has a value) or none (if it doesn't), this will allow us to use range operations with it. <cpp>std::expected</cpp> can either contain a value of the expected type, or a value from the unexpected type (<cpp>std::unexpected</cpp>).

an example of using <cpp>std::expected</cpp>, with one possible pitfall.

```cpp
void loadWidget()
{
  if (const auto widgetBox = getNewWidget(); widgetBox.has_value()) {
    const auto widget = widgetBox.value();
    // Do something with the widget ...
  } else {
    const auto error = widgetBox.error();
    // Handle the error ...
 }
}

void loadWidgetV2()
{
  const auto widgetBox = getNewWidget();
  if (widgetBox.has_value()) {
    // Do something with the widget ...
  } else {
    const auto error = widgetBox.error();
    log("Cannot get a new widget {}: {}.", widgetBox.value(), error); // error! calling .value() 
  }
}
```

monadic operations - functional programming style

- <cpp>and_then</cpp>
- <cpp>or_else</cpp>
- <cpp>transform</cpp>
- <cpp>transform_error</cpp>

```cpp
getWidget()
  .and_then(
  [](const auto &widget) -> std::expected<Widget, WidgetError> {
    // Do something with the widget ...
    return widget;
  });

getWidget()
  .transform([](const auto &widget)->ID {return widget.id();});

getWidget()
  .and_then(/*...*/) // do something
  .or_else([](const auto& error) {log(error);}); // do this on the error

getWidget()
  .and_then(/*...*/) // do something
  .transform_error(&fmt::to_string<WidgetError>); // call this on the error,
```

moving to monadic operations begins with defining boundaries, what can be changed, usually this is the class or library boundary. this depends on who consumes the code and how much we can change the API.

we can use <cpp>std::bind_front</cpp> and <cpp>std::bind_back</cpp> to create closures over functions to pass them as monadic operations. (suggestion to move away from OOP to functional programming).
</details>

### Ranges++: Are Output Range Adaptors the Next Iteration of C++ Ranges? - Daisy Hollman

<details>
<summary>
Understanding the problem with range iterator and suggesting improvements.
</summary>

[Ranges++: Are Output Range Adaptors the Next Iteration of C++ Ranges?](https://youtu.be/NAwn5WqNXJw?si=JK89fj9RizxRymgI), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Ranges.html), [event](https://cppcon2024.sched.com/event/1gZgc/ranges++-are-output-range-adaptors-the-next-iteration-of-c++-ranges), [daisy-chain repository](https://github.com/dhollman/daisychains).

Major flaw in the design of ranges.

#### The Problem of ranges

> TPOIASI - "The Terrible Problem Of Incrementing A Smart Iterator"

```cpp
println("{}",
  std::vector{1, 2, 3, 4, 5}
  | transform([](int n) { return n * 2; })
  | filter([](int n) { return n % 4 == 0; })
); 
// [4, 8]
```

if we do debug printing and use lambdas,  we can see the calls. for elements which pass the filter stage, we see another call to the previous step in the pipeline.

```cpp
auto t = [](int n) {
  println("transform({})", n);
  return n * 2;
};
auto f = [](int n) {
  println("filter({})", n);
  return n % 4 == 0;
};
println("{}", 
  vector{1, 2, 3, 4, 5}
  | transform(t) | filter(f));

/*
transform(1)
filter(2)
transform(2)
filter(4)
transform(2) - why this?
transform(3) 
filter(6)
transform(4)
filter(8)
transform(4) - why this?
transform(5)
filter(10)
[4, 8]
*/
```

we can look at the standard library implementation and try to get an understanding of the problem, there is a lot of weird code in <cpp>std::transform_view</cpp>. eventually, we drill down enough and find the <cpp>operator++</cpp> of the iterator, which simply increments the underlying iterator. there is also something about the dereference operator.

> Is <cpp>filter_view</cpp> broken?
>
> Anything that dereferences the underlying iterator somewhere other than <cpp>operator*()</cpp> is going to have this problem. such as:
>
> - <cpp>drop_while</cpp> - in <cpp>drop_while_view::begin()</cpp>
> - <cpp>take_while</cpp> - in <cpp>sentinel_t<take_while_view>::operator==()</cpp>
> - <cpp>chunk_by</cpp>
> - <cpp>slide</cpp>
> - <cpp>set_difference</cpp>
> - <cpp>is_permutation</cpp>
> - <cpp>std::sort</cpp>

we invoke the dereference behavior both when we iterate and when we dereference, so we have more calls than expected.

```cpp
println("{}", vector{1, 2, 3, 4, 5} 
  | slide(3));
/*
[[1,2,3], [2,3,4], [3,4,5]]
*/

auto f = [](int x) { return x * 2; };
println("{}", vector{1, 2, 3, 4, 5} 
  | transform(f))

/*
[[2,4,6], [4,6,8], [6,8,10]]
*/
```

lets do more debug printing:

```cpp
auto f = [](int x) { 
  if(x == 3) println("f(3)");
  return x * 2;
};
println("{}", vector{1, 2, 3, 4, 5} 
  | transform(f) 
  | slide(3));
/*
f(3)
f(3)
f(3)
[[2,4,6], [4,6,8], [6,8,10]]
*/
```

> If `f` is "expensive", the author of this code probably doesn't understand that `f` isn't evaluated until it's needed ("evaluated lazily").\
> We can argue about whether or not that's a reason to disallow ranges in large-scale software engineering context.\
> We can argue about whether or not that's a reason to disallow ranges in large-scale software engineering context.
>
> Is `transform(f) | filter(g)` the same problem?

```cpp
int count_f_2 = 0;
auto f = [&](int n) { count_f_2 += (n == 2); return n * 2; };
auto g = [](int n) { return n % 4 == 0; };
println("{}", vector{1, 2, 3, 4, 5} 
  | transform(f)
  | filter(g));
println("f(2) was called {} times", count_f_2);

// [4, 8]
// f(2) was called 2 times
```

we need to understand lazy and eager evaluations.

> Question: when is the result of f(2) "needed"?\
> It depends on how lazy you are!\
> The latest possible time when you need the result of f(2) is right before you have to print it.\
> The "surprise" here comes from eagerly evaluating <cpp>iterator_t<filter_view>::operator++()</cpp> before we know when (or if) <cpp>iterator_t<filter_view>::operator*()</cpp> is going to be called!
>
> - This is the most important point in this presentation.
> - This is a different "flavor" of surprise than not understanding laziness generally.
> - What if we wait until we know what we're going to do with the result of `f(2)` before we call it?

#### Pull vs. Push Programming Models

trying to solve the issue.

pull and push models aren't the same as eager and lazy models.

```cpp
auto f = [](int x) { return x * 2; }
auto g = [](int y) { return y % 4 == 0; }
auto v = iota(1, 6) 
  | transform(f) 
  | filter(g)
  | to<vector>();
println("{}", v);
// [4, 8]
```

we can model this as each view wrapping its' input, wrapping from right-to-left. but alternatively, what if we wrapped over the output, and treated it as wrapping from left-to-right?

```cpp
// Pseudo-code hand-wavy you-get-the-idea
// input wrapping
auto view1 = transform(iota(1, 6), f);
auto view2 = filter(view1, g);
auto view3 = to<vector>(view2);

// output wrapping
auto link3 = filter(to<vector>(), g);
auto link2 = transform(link3, f);
auto link1 = iota(link2, 1, 6);
```

the pull model is right-to-left - wrapping the input. the push model is left-to-right - wrapping the output.

```cpp
auto v = evaluate(
  iota(1, 6) 
  >>= transform([](int x) { return x * 2; })
  >>= filter([](int y) { return y % 4 == 0; })
  >>= to<vector<int>>());
std::println("{}", v);
```

<cpp>operator>>=()</cpp> has different associativity than the pipe operator. the right associativity means that we first evaluate the right side, just as we want for the push model!

```cpp
w | x | y | z;  // ((w | x) | y) | z
a >>= b >>= c >>= d;  // a >>= (b >>= (c >>= d))
```

we want something like a pipe, but using the `>>=` behavior with the right-to-left associativity.\
we will try to implement this ourselves. we implement the filter, the transformer, the `iota` generator, the creation of the vector, the evaluator

```cpp
auto obj1 = filter(f);
auto obj2 = obj1{transform(g)};

template <class Pred>
struct filter {
  Pred pred;
  template <class NextLink>
  struct adaptor {
    Pred pred;
    NextLink next;
    template <class T>
    auto push_value(T value) {
      if (pred(value)) {
        next.push_value(value);
      }
    }
  };
};

template <class Fn>
struct transform {
  Fn fn;
  template <class NextLink>
  struct adaptor {
    Fn fn;
    NextLink next;
    template <class T>
    auto push_value(T value) {
      next.push_value(fn(value));
    }
  };
};
 
struct iota {
  int ibegin, iend;
  template <class NextLink>
  struct adaptor {
    int ibegin, iend;
    NextLink next;
    auto generate() {
      for (int i = ibegin; i < iend; ++i) {
        next.push_value(i);
      }
    }
  };
};

template <class Container>
struct to {
  Container output;
  template <class T>
  auto push_value(T value) {
    output.push_back(value);
  }
  template <class Self>
  decltype(auto) get_output(this Self&& self) { 
    return std::forward_like<Self>(self.output);
  }
};
```

we need to define the operator overload and create the evaluator. we use a mixin (not CRTP) for an intrusive inheritance.

```cpp
template <class Chain>
auto evaluate(Chain chain) {
  chain.generate();
  return std::move(chain).get_output();
}

struct link_base {
  template <class Self, class NextLink>
  auto operator>>=(this Self&& self, NextLink next) {
    return std::forward<Self>(self).adapt(std::move(next));
  }
};
 
template <class Fn>
struct filter : link_base {
  Fn fn;
  explicit filter(Fn f) : fn(f) {}
  template <class NextLink>
  auto adapt(NextLink next) {
    return adaptor<NextLink>{fn, std::move(next)};
  }
  /* ... */
};

struct output_passthrough {
  template <class Self>
  decltype(auto) get_output(this Self&& self) { 
    return std::forward_like<Self>(self.next).get_output();
  }
};
```

we are wrapping right-to-left, from outermost to inner most.

[compiler explorer](https://cppcon.godbolt.org/z/hY63aEfKo) example with all the code,

we can also do the same for <cpp>take_while</cpp>, we need to update the convention for breaking out of a link in a chain.

we still have some problems, reverse doesn't work with output wrappers, neither does sorting, and there are use cases which are awkward to create (slide window). but there are also use-cases which don't work well with the pull-model, like flattening a one-to-many transformation. the two models can exist alongside one-another.
</details>

### C++ Under the Hood: Internal Class Mechanisms - Chris Ryan

<details>
<summary>
Learning a bit about how C++ classes work.
</summary>

[C++ Under the Hood: Internal Class Mechanisms](https://youtu.be/gWinNE5rd6Q?si=ZyXX7nt5Cof56LZc), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Cpp_Under_The_Hood.pdf), [event](https://cppcon2024.sched.com/event/1gZfx/c++-under-the-hood-internal-class-mechanisms).

> We are going to talk about internal C++ class mechanisms:
>
> - C++ Onion, inheritance and polymorphic mechanisms.
> - Member Data Pointers
> - Member Function Pointers
>   - Stack Frame / Base Pointer mechanics
>   - Construction/Destruction order
>   - Running code before and after main()

#### The Virtual Table

inheritance vs composition (aggregation). a person **has a** name, but a person **is not a** name. an employee **is a** person, but an employee **doesn't have a** name.

the inheritance hierarchy, the hidden pointer to the virtual table in the base class.

```cpp
struct A {
  v_table* __v_table_ptr; // hidden
  int a{};
  std::string str1{};
  A() {Bar(); /*... */}
  virtual A~() {Bar(); /*... */}
  virtual void Foo() { std::cout << "A::Foo()"}
  void Bar() { Foo();}
};

struct B : A {
  int b{};
  std::string str2{};
  B() {Bar(); /*... */}
  virtual B~() {Bar(); /*... */}
  virtual void Foo() { std::cout << "B::Foo()"}
};

struct C : B {
  int c{};
  std::string str3{};
  C() {Bar(); /*... */}
  virtual C~() {Bar(); /*... */}
  virtual void Foo() { std::cout << "C::Foo()"}
};

static A::'v_table'[]= {A::~(), A::Foo()}
static B::'v_table'[]= {B::~(), B::Foo()}
static C::'v_table'[]= {C::~(), C::Foo()}
```

when a constructor is called, it initializes the pointer to the class v_table, each constructor and destructor set the pointer according to the class, so the value of the pointer can change. the base class is effectively a data member.\
constructing is done inside-out (expanding out), while destructing is outside-in (digging in).

#### Pointer Types

> - C Pointer types
>   - `int *` Pointer to Data
>   - `int(*)(int)` Pointer to Function
> - Member Pointer types
>   - `int Foo::*` Pointer to Member Data
>   - `int(Foo::*)(int)` Pointer to Member Function

[compiler explorer](https://godbolt.org/z/rKzdTsejz) example.

pointer to function declaration: `int (*foo) (int, int)`, the return type of the target function on the left, the parameter types of the target function on the right, and the declaration (pointer type and the name) in the middle. it's a lot easier to define in outside using `typedef` or a `using` alias. C is lenient with the placement of the `*` symbol, but C++ is stricter.\
a pointer to a member function has some weird stuff when taking the address of a virtual function. there is internal secret `Thunk` function, which is returned when taking the address of a polymorphic function, the function calculates the appropiate address based on the call site, and does the re-direction.

> Fundamental Theorem of Software Engineering:\
> Any problem can be solved by adding a layer of indirection.

`Null` is usually zero in C, but that's not the standard, the C standard says the Null should evaluate to zero, and unassigned pointers should have the value of null. C++ has <cpp>std::nullptr</cpp>, which is more clearly defined, it's still not required to be zero. our null is actually -1 in memory, because of some reasons relating to offset in memory.

we can create a sorting function that takes a data member pointer, so it's a generic thing which can sort based on a key chosen programmatically. we can also expand this to sort using multiple keys with templated pack expansion.

[compiler explorer](https://godbolt.org/z/e9r9fbvvz) example.

#### Object Memory Footprint.

the virtual table is always at offset zero, it's always nice the destructor is the first method in the virtual table array. things get complicated with multiple inheritance (even without the diamond inheritance problem).

#### Stack Frames

the startup routine calls the main function, it pushes the stack pointer and the base pointer accordingly. something similar happens for every function call (which opens a new frame). also something with the instruction pointer.

The RTTI (real time type information) is also part of the virtual table.
</details>

### How to Hide C++ Implementation Details - Amir Kirsh

<details>
<summary>
Ideas for hiding how our code is implemented.
</summary>

[How to Hide C++ Implementation Details](https://youtu.be/G5tXjYzfg9A?si=bMXXQyCMo4ymlDEa), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Hiding_your_Implementation_Details_is_Not_So_Simple.pdf), [event](https://cppcon2024.sched.com/event/1gZgW/hiding-your-implementation-details-is-not-so-simple).

> - Encapsulation
>   - Protect Object Integrity - Expose the necessary interfaces only.
>   - Easier to Use - Users see only what they need to.
>   - Improves Maintainability - Internal changes don't affect external code.
>   - Easier to Debug - Data modifications happens internally, in specific places.
> - Decoupling
>   - Reduces Dependencies - Components depend only what we actually expose.
>   - Facilitates Changes - Changes in one part of the system don't necessarily affect others.
>   - Enhances Reusability - Components become more generic, thus can be better reused.
>   - Improves Testability - Easier to test components in isolation.

we agree that hiding implementation details is important (access modifiers), but it's not always easy. there is always time pressure and other stuff with higher pritoties. some times there are even real technical constraints, which we need to work with. we don't want to always work with interfaces to avoid virtual calls, and sometimes exposing the returned type helps with performance.

```cpp
std::pair<First, Second> p;
auto k = p.first;
auto v = p.second;

class Bar {
  // ... more code
public:
  const std::map<string, int>& getProperties() const;
};

class Foo {
public:
  // we prefer the ctor below to be private but it doesn't compile :(
  Foo(int value) : value(value) {}
  static unique_ptr<Foo> create(int value) {
    return std::make_unique<Foo>(value); // this needs a public constructor
  }
};
auto foo = Foo::create(7);
```

> our goal: Investigate code examples:
>
> - See how we expose implementation details
> - Understand why it's not the best design, per case
> - Propose a better design, hiding the implementation details

for example, <cpp>std::pair</cpp> having it private members exposed prevents us from having a lazy evaluated second member.

```cpp
// assume function is constrained on PAIR having fields first and second
template<typename PAIR>
std::ostream& operator<<(std::ostream& out, const PAIR& pair) {
  return out << pair.first << ' ' << pair.second;
}
// ok
auto stdpair = std::make_pair("bye", 42);
std::cout << stdpair << std::endl;
// can we allow our pair to lazily evaluate its second?
auto squarepair = SquarePair(7);
std::cout << squarepair << std::endl;
```

also, structs, in general, have public data members by default. if we use private data members and expose "getters", then we can treat those getters as constraints for templates, and have different kinds of objects with the same public interface. (duck typing).

> Takeaway
>
> - Keep data members (and static members) in the private.
> - Always.
> - To begin with.

#### API (including internal API of all kinds!)

> Hyrum's Law:\
> "Developers come to depend on all observable traits and behaviors of an interface, even if they are not defined in the contract"\
> With a sufficient number of users of an API, it does not matter what you promise in the contract: all observable behaviors of your system will be depended on by somebody.

applications tend to use a small part of the API, and the unused parts of APIs are more bug prone. having too many options (including overloads) will lead to the wrong option being selected at some point. smaller APIs are more testable.

> Be Lean and Mean
>
> 1. Keep your API short and concise. (add as needed, and only for there is a real need).
> 2. Limit your API to what should actually be used (arguments, return value etc.) - do not expose too much of your design!
> 3. Different usages may get different *view* (or copy) of the same data, exposing only what is relevant.

context-specific, things might be required to be exposed for one use-case, but be hidden for other usages.  

> - If they don't need it, don't give it.
> - If they get it, they'll use it.
> - If they use it, they'll abuse it.
> - Trust no one

language tools for passing arguments:

- Value semantics
- Interfaces
- Pimpl idiom
- Wrapper classes as a view, wrapping the original
- C++20 concepts

> Exercise - "Const Map, Mutable Vals"\
> Initialize a map, then pass it in a way such that values can be modified but the map itself cannot be modified.

reminder about SFINAE and c++20 <cpp>concepts</cpp>, the `requires` clause and `requires` expression.

```cpp
// note: function may modify values but should not alert the map itself!
void foo(ConstDictMutableVals auto& d) {
  d[1] = 3;
}
int main() {
  std::map<int, int> myMap = {{1, 0}, {2, 0}};
  foo(myMap);
  for (const auto& [k, v] : myMap) {
    std::cout << k << ": " << v << std::endl;
  }
}
```

the above code isn't enough, since if the key doesn't exists, it creates it. let's check the key exists first.

```cpp
void foo(ConstDictMutableVals auto& d) {
  auto itr = d.find(1);
  if(itr != d.end()) itr->second = 3;
}

template <typename T>
concept ConstDictMutableVals = requires(T t, typename T::key_type key) {
  typename T::iterator;
  typename T::key_type;
  typename T::mapped_type;
  { t.find(key) } -> std::same_as<typename T::iterator>;
  t.find(key) != t.end();
};
```

this isn't enough, it only tells us what we can do, but it doesn't prevent us from using other stuff. we might be able to test it by testing with a class that only provides what's defined in the concept, and not anything else.\
we can also apply the wrapper approach.

```cpp
template<Dictionary Dict>
void foo(ConstDictMutableValues<Dict> d) {
  d.assign(1, 3);
}

template<Dictionary Dict>
class ConstDictMutableValues {
  Dict& d_;
public:
  explicit ConstDictMutableValues(Dict& d) : d_(d) {}
  bool assign(const Key& key, Value value) {
    auto itr = d_.find(key);
    if(itr == d_.end()) return false;
    itr->second = std::move(value);
    return true;
  }
};
```

our wrapper class takes the reference and only exposes what we need to use in the method.

#### Hiding Smart Pointers and Hierarchies

this code has too many pointers a allocation with raw pointers.

```cpp
// we want the following code:
Expression* e = new Sum (
  new Exp(new Number(3), new Number(2)),
  new Number(-1)
);
cout << *e << " = " << e->eval() << endl;
delete e;
// to print something like:
// ((3 ^ 2) + (-1)) = 8
```

moving to smart pointers, it still looks bad because we are requiring the caller to pass the <cpp>std::unique_ptr</cpp>.

```cpp
auto e = std::make_unique<Sum>(
  std::make_unique<Exp>(
    std::make_unique<Number>(3),
    std::make_unique<Number>(2)
  ),
  std::make_unique<Number>(-1)
);
cout << *e << " = " << e->eval() << endl;
```

we could hide the smart pointer inside the constructor of the class, it's easier to read for the user, and we can change how we implement the behavior in the future.

```cpp
class BinaryExpression: public Expression {
  std::unique_ptr<Expression> e1, e2;
public:
template<typename Expression1, typename Expression2>
BinaryExpression(Expression1 e1, Expression2 e2)
: e1(std::make_unique<Expression1>(std::move(e1))),
e2(std::make_unique<Expression2>(std::move(e2))) {}
// ...
};
```

there are also design patterns to hid implementation details:

> - State Pattern - Employee is an Employee, regardless of his employment type - a state
> - Strategy Pattern - A PathFinder is a PathFinder, regardless if using BFS or DFS - a strategy
> - Factory Method - The user shall not be bothered with the exact object type instantiated

exposing less information on return values:

```cpp
// what do you say about this one?
vector<int> foo() {
  return vector {1, 2, 4};
}

// is this one better?
auto foo() {
  return vector {1, 2, 4};
}

// and what about this one?
random_access_range_of<int> auto foo2() {
  return vector {1, 2, 4};
}
```

#### Summary and Additional Stuff

leaky abstraction - such as <cpp>std::vector\<bool></cpp> which returns a proxy and basically shouldn't be used. hiding away helper functions inside inner classes. preferring protected member functions over protected data members. keeping the data members private as much as possible.

testing private behavior is usually a sign that something is wrong. it's ok to get the state, but it's a problem if we have a set 'setter'.

the private token idiom also called "Access Token Pattern", "Passkey Pattern", "Authorization Token Pattern", "Key Token Idiom". this is how we solve the original problem of having <cpp>std::make_unique</cpp> and still prevent others from creating objects of the class.

```cpp
class Foo {
  int value;
  class PrivateToken {};
public:
  Foo(int value, PrivateToken) : value(value) {}
  static unique_ptr<Foo> create(int value) {
    return std::make_unique<Foo>(value, PrivateToken{});
  }
  int getValue() const { return value; }
};

auto foo = Foo::create(7);
// we can never create a private token outside the class, only the class and it's friends can can the private token.
```

</details>

### Refactoring C++ Code for Unit testing with Dependency Injection - Peter Muldoon

<details>
<summary>
Bringing in dependency injection to real code-bases.
</summary>

[Refactoring C++ Code for Unit testing with Dependency Injection](https://youtu.be/as5Z45G59Ws?si=ECxmQ7wf571cc4oY), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Dependency_Injection_in_Cpp.pdf), [event](https://cppcon2024.sched.com/event/1gZfe/dependency-injection-in-c++-a-practical-guide).

> Dependency Injection
>
> 1. Decreases coupling between functionality blocks
> 2. Used to make a class/function independent of its underlying dependencies
> 3. Improves the reusability – and so testability - of code.
> 4. Better long term maintainability of code.

Testing without dependency injection is more similar to integration testing, it requires bringing up a function full enviornment (more or less), configurations, databases, components, etc...\
It's harder to specify the test-cases, it's longer to run the entire thing which makes development slower.

Dependency injection relies on interfaces, rather than use real components like a real database we rely on mock components which behave like them, but without side effects.

#### Dependency Injection Options

> Link-time Dependency Injection - Uses Link-time switching of functionality
>
> - Allows limited Testing
> - No code changes/contamination in actual production application required
> - The code using the dependent functionality has no say in which implementation is being executed
>   - Externally Injected during compilation via `LIBPATH` or `#IFDEF`

there are drawbacks, if each test case requires a different build, we end up with a very complicate build file. this also leads to many undefined behavior and ODR violations.

> Dependency Injection via inheritance- Create a base class interface or extend from an existing Class
> 
> - Can handle lots of methods
> - Rich interface
> - Well understood mechanism
> - Virtual functions + override
> - Easier to add to older codebases

this works well with old code bases from the era of OOP, so it's easy to substitute them. we can also use interfaces instead of inheritance, so the test class doesn't interfere with the real class.\
however, this can lead to messy interfaces which must be in the real interface, and there's the problem of adding extra indirection calls (not a real performance effect, to be honest).

> Dependency Injection via templates - Create a Class that satisfies the calls made on the class by the function
> 
> - Can handle lots of methods being mocked
>   - Only need to define the methods actually used
> - Compile time so no runtime virtual calls overhead
> - Can use concepts(C++20) to define an "interface"

this requires a lot of code change, heavy use of templates, and increases compilation time.

> Dependency Injection via type erasure - Call any thing satisfying a function signature – via <cpp>std::function</cpp>/<cpp>std::move_only_function</cpp>/<cpp>std::invoke</cpp> or something similar
>
> - Invokable on any callable target
> - Versatile

this has similar overhead to runtime virtual calls, and we can only handle one method being substituted.

> Null Valued Objects - A stub with no functionality - only satisfying the type requirements
> 
> - Disables a part(s) of the system not under test
> - Supplies the correct type but no actual implementation logic
>   - Supplied arguments discarded
>   - Returns fixed values

##### Dependency Injection Types

when and how we inject the mocking substitute.

- Setter Dependency Injection - might leave the class in unusable state, makes the whole thing dangerous.
- Method Injection - changes the function signature to take the method
- construction injection - capture the dependency in the constructor. changes the signature of the class creation.

#### Real World Dependency Injection

the concept is something like this:

> Control all Dependencies in a system
> 
> - Identify functional blocks
>   - Allow Injection of flexible functionality
>   - Capture inputs, control outputs
> - Where to insert Dependencies
>   - Drop all "invariant" dependencies into constructors
>   - Drop all other dependencies into function method calls
> - Just add a virtual hop for called functions
>   - Add virtual to function declarations
>   - Use function forwarding of calls
>   - Ignores bad interfaces

in real codebases, there are problems, code gathers complexity and noise over time.

> Dependency Injection roadblocks
>
> - Objects full creation hidden inside functions/classes
>   - No handle to inject new functionality
>   - Default class constructors initialized via Singletons/Globals
> - Reaching through multiple objects
>   - Long chains of mock classes needed as boilerplate
>   - Breaks the principle of least knowledge
> - Disentangling getting information from setting state
>   - Dig out the pure functions
> - Having too many dependencies in a class / functional block
>   - Impractical to pass large number of Dependencies in constructor / method
> - Classes (hierarchies) packed with huge chunks of functionality
>   - God Classes doing too many things
>   - Many dependencies, too numerous to inject
> - Functionality splintered and spread throughout the codebase
>   - Fragmented throughout the inheritance chain
>   - Duplicated throughout the codebase
>   - Blended into general utility classes
> - Lack of Data structure
>   - Ungrouped data

there are ways to refactor the code to be more suited to unit-testing, we can slowly extract pieces of code into internal classes and interfaces, we can separate stuff and start using the new things to test one another with dependency injection. breaking apart methods with too many parameters. these are all good practices regardless of dependency injection.

Dependency injection for widely used APIs, using default parameters (keep API, break ABI), having the old function forward to a new function (keep API and ABI).\
Using lazy initialization and passing stuff as injection, using dependencies providers. also a problem with templated member functions (they can't be virtual).

</details>
