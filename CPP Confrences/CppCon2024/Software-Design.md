<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Software Design

<summary>
17 Talks about Software Design.
</summary>

- [ ] 10 Problems Large Companies Have with Managing C++ Dependencies and How to Solve Them - Augustin Popa
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
- [ ] Relocation: Blazing Fast Save And Restore, Then More! - Eduardo Madrid
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
