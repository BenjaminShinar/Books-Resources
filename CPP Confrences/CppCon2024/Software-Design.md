<!--
// cSpell:ignore relocatability Björn Fahller
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Software Design

<summary>
17 Talks about Software Design.
</summary>

- [x] 10 Problems Large Companies Have with Managing C++ Dependencies and How to Solve Them - Augustin Popa
- [ ] Adventures with Legacy Codebases: Tales of Incremental Improvement - Roth Michaels
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
- [ ] The Most Important Design Guideline is Testability - Jody Hagins

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
