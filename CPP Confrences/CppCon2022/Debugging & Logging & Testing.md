<!--
ignore these words in spell check for this file
// cSpell:ignore Bazel intellej whatis Semvar Deployers Automatable Kitware readelf objdump dwarfdump gsplit debuginfod jork uzzer fuzzer fsanitize
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Debugging & Logging & Testing

### Compilation Speedup Using C++ Modules: A Case Study - Chuanqi Xu

<details>
<summary>
A case study of compilation times when using modules.
</summary>

[Compilation Speedup Using C++ Modules: A Case Study](https://youtu.be/0f5N1JKo4D4)

Introduction

- modules
- complexity theory
- compilation model
- inline functions and templates

Modules can be named modules or header units, named modules can compiled into object files, but header units cannot. the focus of this talk is on name modules. Modules use `import std` instead of `#include`. a module consists of several module units.

With header files, the code in the header can be compiled multiple times across different object files. with modules, the code is compile into a prebuilt module interface, and is then imported. theoretically, if we have _n_ header used in _m_ source files, the compile complexity is _O(n\*m)_. with _n_ modules and _m_ source files, the complexity becomes _O(n+m)_. this is of course a simplification, which ignored compiler optimizations, templates and inlining.

#### Case Study

the case study is the **async_simple** project, which wraps two libraries (std, asio) into named modules, and contains consumers of those modules.\
We measure the compilation time under different optimization flags.

the results show improvement for nearly all files without optimizations, but the turning optimizations on removes the savings in some cases and reduces them in the other cases. the speedup depends on the pattern and the usage for each file.

#### Possible Improvement

- compiler side
- library side
- long term

combining function reduction with LTO (link time optimization). elision-ing unreachable declarations in the global module fragment.

</details>

### Back to Basics: Debugging in C++ - Mike Shah

<details>
<summary>
introduction to debugging and using GBD
</summary>

[Back to Basics: Debugging in C++](https://youtu.be/YzIBwqWC6EM), [slides](http://www.mshah.io/conf/22/CPPCON2022_BacktoBasicsDebugging.pdf), [github](https://github.com/MikeShah/Talks/tree/main/2022_cppcon_debugging).

debugging constitutes about half the time and cost of programmers. knowing how to debug programs allows the developer to get out of sticky situation.

defining bugs - logic, correctness, performance, non-deterministic behaviors.\
Bugs can hide in the codebase for a long time.

debugging and testing are related. test for the presence of bug, remove the bug and test again.

#### Debugging Strategies

even if the code compiles, it doesn't mean it's correct.

**scan and look** - look at the code and hope something pops out.

in some cases, the compiler can help if we give it the appropiate flags, but we don't always have the source code, and we don't always want compile everything.

**printf debugging** - spreading console logs in the source code. it's a useful strategy, but it means we change the source code and pollute it, and it requires modifying the code and re-compiling it, and it doesn't scale well.

**delta debugging** - searching for where the bug can be.

**printf debugging with language support** - using the preprocessor and debug flags to conditionally print debug messages.

#### Introduction to GDB

an interactive debugger should be the first choice. every IDE has integration with a debugger.\
debuggers are attaching themselves into a process, and if we compile the program with debug symbols, then we have a better experience.

1. add the `-g` flag when compiling.\
2. run the debugger with with the `--tui` flag (text user interface mode). (`--silent` can help)

gdb interactive

- run - run program
- start - temporary breakpoint at main.
- next - next line (or number of commands)
- continue - execute until breakpoint
- list - show the code around us
- step - step into the current function
- break #line number - add breakpoint to line or function or whatever
  - break #line number if condition - conditional breakpoint.
- layout - change layout in tui mode
- print - print variable
- info breakpoints - view breakpoints
- refresh - refresh tui window
- locals - show local variables
- backtrace - show call stack
  - up, down - navigate frames
- watch - add variable to watch list - we will be notified when it changes
- whatis - get type of a variable
- ptype - print the type definition of a variable
- focus - move focus to somewhere in the code
- set variable = value - modify the value of a variable
- target record-full - allow for going back in time

we can attach to an already running process.

`gdb attach #processId`

</details>

### C++ Dependencies Don’t Have To Be Painful - Why You Should Use a Package Manager - Augustin Popa

<details>
<summary>
Detailing the advantages of using a Package manager
</summary>

[C++ Dependencies Don’t Have To Be Painful - Why You Should Use a Package Manager](https://youtu.be/rrcngYMAJ-w), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/C-Dependencies-Dont-Have-To-Be-Painful-Why-You-Should-Use-a-Package-Manager.pdf)

surveys from previous years about frustration points in c++ developments point out that the number one major pain point over the years is about managing libraries.\
another question in those surveys finds that package managers for c++ aren't very popular, and that many developers still manage them manually.

ABI compatibility between libraries and consuming project

> A common C++ problem, uncommon in other programming languages

ways to break ABI and your builds

- change the compiler
- change the compiler version
- change the target os
- change the target architecture
- add optional features to a library

the "diamond problem" is when a project depends on two libraries, which in turn depend on a third library. we can only have one version of each library, so the two libraries must depend on the same version of the shared library.

package managers were designed to solve some (or all) of these problems. for vcpkg, updating version of a library means that they test each dependant library.

many people avoid touching dependencies, but it's important to keep dependencies up to date. it allows moving to new tools, new standards, and avoid security vulnerabilities.

package managers have additional benefits for productivity, testability, and performance.

cases when a package manager should be used:

> 1. When your project has more than 1-2 dependencies, or you have dependencies of dependencies
> 2. When you have open-source dependencies
> 3. When your project has no dependencies, but you want to implement something that is already available in the public domain
> 4. When you are thinking about making your library header-only because it will make it more portable
> 5. If you are concerned about maintenance time or security

modules won't fix all the issues, they will address some issues, but not those of maintaining abi stability, solving the diamond problem, and they will still require the open source libraries to be migrated into modules.

#### Package managers

- System package managers: apt, yum, rpp, chocolate, etc..
- C++ package managers: vcpkg, conan, hunter
- Other language package managers: npm, nuget, pip...

there is a time and case for each of these, but if the project is primarily c++ and is complex enough, then a c++ package manger should be used.

the package manager should be able to help with creating a reproducible development enviornment. including the SDK, compilers, tools etc...

principles

> - Be able to build dependencies from source (when necessary)
> - Keep your dependencies up to date
> - Cross-platform should be first class experience
> - Make your build environment reproducible
> - Do download prebuilt binaries, if they are verified
> - Simplify workflow for authoring and publishing dependencies
> - Take advantage of existing opensource solutions
> - Enforce ABI requirements across all packages, not one at a time
> - Use more than one package manager if it improves your productivity

</details>

### import CMake, CMake and C++20 Modules - Bill Hoffman

<details>
<summary>
How CMake is handling Modules.
</summary>

[import CMake, CMake and C++20 Modules](https://youtu.be/5X803cXe02Y), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-2022.pdf)

introducing the connection between the Kitware company and CMake. Cmake was born from a need to create a library.

if Boost aims to give c++ a set of useful libraries, then CMake aims to give C++ compile portability, be able to run the same tool and compile for many platforms. CMake is now shipped with Visual Studio.

> CMake adapts to new technologies so developers don't have to.

when there's a new tool, like a compiler version, package manager,or something then projects which are set up for CMake get that tool.

- cmake-gui - visual tool
- ccmake - interactive cli tool
- cmake - non interactive tool

Cmake became "target-centric", no difference between external and internal target.

| option               | meaning                                                          |
| -------------------- | ---------------------------------------------------------------- |
| PRIVATE              | Only the given target will use it                                |
| INTERFACE            | Only consuming targets use it                                    |
| PUBLIC               | PRIVATE + INTERFACE                                              |
| $<BUILD_INTERFACE>   | Used by consumers from this project or using the build directory |
| $<INSTALL_INTERFACE> | Used by consumers after this target has been installed           |

presets allow for common configuration, both shared between developers and private. and there is packaging wizard (CPack), which creates installers.

Testing: `add_test` for normal test, `ctest` is the command to run the tests. `gtest_discover_tests`, processor affinities. the CDash test dashboard.

#### C++ Modules

in c++20, modules mean that the build order matters. we start working with **Build module interfaces** (BMI). Cmake actually got a head start supporting modules because it supported Fortran. there was already a basis for dynamically creating a build graph of dependencies.

```fortran
module math
contains
  function add(a,b)
    real :: add
    real, intent(in) :: a
    real, intent(in) :: b
    add = a + b
  end function
end module
```

```fortran
program main
  use math
  print *,'sum is', add(1.,2.)
end program
```

cmake needs support from the standard to have a format for dependencies,

FILE_SET - a new feature in CMake, adding files to a fileSet, helps with header-only libraries and modules.

> Named Module Types
>
> - "module unit" is a translation unit that contains a module-declaration. A "named module" is the collection of module units with the same module-name.
> - A "module interface unit" is a module unit whose module-declarations tarts with export-keyword; any other module unit is a "module implementation unit".
> - A "module partition" is a module unit whose module-declaration contains module-partition.

everything is still in testing stages, but there are experimental versions which we can try.

</details>

### Going Beyond Build Distribution - Using Incredibuild to Accelerate Static Code Analysis and Builds - Jonathan "Beau" Peck

<details>
<summary>
Demonstrating the power the static analyzers
</summary>

[Going Beyond Build Distribution - Using Incredibuild to Accelerate Static Code Analysis and Builds](https://youtu.be/M7zMl2WOp6g), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-Sep2022-GoingBeyondBuildDistribution-Final-Upload.pptx)

"effort over time", getting stuff done faster, a tool that reduces compilation time and allows to make better use of the developer time.

we want good tools, minimum effort and maximum results, simple to use with little configuration, flexible integration with different tools and pipelines.

- solve a problem - go fast, accelerate development cycles
- simple to use - single binary, comsume existing infrastructure
- flexible integration - seamless, agnostic consumption

Development Acceleration Platform.

it began with the SETI (search for extra terrestrial intelligence) project. the root is workload distribution, or "Process Virtualization". one machine initiates the build, and the IncrediBuild platform intercepts this call and manages the build across builder/helper machines, including synchronizing files, linking, etc... there is not enviornment duplication, we don't have to copy all the source code and configuration each time.

three components

- Initiator machine
- Coordinator
- Helpers

it's all the same binary, just different configuration, the coordinator is service, and the other machines are agents.

#### The Demos

IncrediBuild has native integration with visual studio with a plugin. we can configure what is used locally and what is directed towards IncrediBuild.\
The coordinator shows us the state of the machines, including the initiators and the helpers. this is centrally managed. so we can configure groups, priorities, logs, etc..\
When a build is running, we get a visualization of the distributed processes, and we can save the build data for later review and monitoring.

for integration with CI\CD, there is IncrediBuild cloud, which can work with spot Instances (which are cheaper than "OnDemand" instances), even if one machine fails (on the cloud or locally), the work is directed to a different machine and nothing is lost.

another integration option is with GitHub actions (or jenkins). this done via the IncrediBuild CLI (rather than a plugin). github actions has two separate runners (self-hosted and github-hosted), but the flow is similar with both.

this is all done with a yaml.

</details>

### Case For a Standardized Package Description Format for External C++ Libraries - Luis Caro Campos

<details>
<summary>
build 3rd party libraries
</summary>

[Case For a Standardized Package Description Format for External C++ Libraries](https://youtu.be/wJYxehofwwc)
, [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/The-case-for-a-standardized-package-description-format-publish.pdf)

consuming packages in C++. according to the recent survey, most companies still integrate the external code into the build process. even though vcpkg and conan are getting more popularity. and even though managing libraries and build times are major pain points for developers.

when we have to consume a 3rd party library and we build it, then we need to create build scripts for it. there are all sorts of ways to do this.\
because of all these problems, header only libraries are really popular, consuming them just requires add an "include" library, and things work.

dependencies can lead to versioning conflicts, so we need to be careful of that.

> Consuming Libraries:\
> How we refer to libraries depends on the abstractions provided by the build system we are using (CMake, Bazel, Make files, Visual Studio or Xcode Projects…)
>
> - The “modern” way is based on “usage requirements” -
> - But in some cases we still see build scripts that propagate "flags" explicitly

"config" files packaged with libraries. this could make things easier, but there are still issues. some libraries implement behavior like gnus' libTools or pkg-config.

> Recap
>
> - We have our project, our sources, our build scripts
> - We want our build scripts to not be concerned about:
>   - External library source files
>   - How those files are built
> - But we do want to be able to consume libraries in our code (compile+link)
> - We want to be agnostic as to how/where the library was built
> - The current approaches all have limitations

libraries: static linker,dynamic linker, module interface unit, a build system needs to communicate with the library package. is there a difference between a library and a package?

could we have something that works for developers, and is still agnostic to the build system? decoupling what we use for the project build and what the 3rd party are using.

considering existing proposals, what they aim to solve and what they don't. interoperability between package managers.

</details>

### C++ Coding with Neovim - Prateek Raman

<details>
<summary>
Intro to using Neovim
</summary>

[C++ Coding with Neovim](https://youtu.be/nzRnWUjGJl8), [slides](https://vvnraman.github.io/cppcon-2022-cpp-neovim/).

#### Overview

LSP - language server protocol, a development tool can use SLP clients to communicate with language servers (such as clang) to get code suggestion (like intellej)

#### Context

vim, inherited from vi. modal editing: insert mode and 'normal mode'. pressing <kbd>i, o/O, a/A</kbd> to move into insertion mode, and <kbd>escape</kbd> to move into normal mode. and <kbd>:</kbd>(colon) to move from normal mode to command mode.

- `:write`
- `:quit`
- or `:wq` in short

neovim is a fork of vim with async support, so that plugins don't have to run on the UI thread. it has lua scripting as first class citizen, rather than just vimScript (which is still supported).

#### Command Line Enviornment

multiple terminals using Tmux.

using **vcpkg**

```sh
git submodule add https:://github.com/microsoft/vcpkg.git
git submodule update
cd vcpkg
./bootstrap-vcpkg.sh -disableMetrics
```

using **CMake**, combining with ninja build system. (_tasks.py_, `invoke`)

to make neovim connect to clang LSP server, we just need the _compile_commands.json_ file in our root folder.

vim has customizable keybindings, which users are encouraged to use and store in a "namespace".

#### C++ Workflows with Neovim

(live coding session)

plugins, file explorer, syntax highlighting, real time compiler feedback from the LSP server.

#### Neovim Setup

same setup works no matter the language. this is the difference from using an IDE (which is usually based around a single language). install from scratch or use an opinionated distribution.

</details>

### -memory-safe C++ - Jim Radigan

<details>
<summary>
Overview of some attack vector through memory-unsafe code and the address sanitizer.
</summary>

[-memory-safe C++](https://youtu.be/ml4t-6bg9-M),[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CPPCON-2022-JRADIGAN-BACKUP.pptx).

back in 2018 the root cause of many CVEs was memory safety, and this still true for many CVEs today.

this talk wil show why memory safety is important, and introduce a new tool to tackle ths issue.

common attack vector use memory safety errors

- ROP - return oriented programming
- DOP - data oriented programming
- BOP - block oriented programming
- DDM - direct memory manipulation

showing the machine instructions code for `mem[val2] = val1;`.

```mips
mov eax, [esp]
mov eb, [esp+8]
mov [ebx], eax
```

or by moving the stack pointer ("gadget") - we modify where the stack will be.

```mips
pop eax
ret
pop ebx
ret
mov [ebx], eax
```

return oriented programming is turing complete, and it can even execute system calls.

list of memory safety-issues - address sanitizer can find them with zero false positives. microsoft has many safety technologies:

- source modification
- static sanitizers
- dynamic sanitizers
- secure code generation

CFG - control flow guard - disallow indirect functions to move into targets which weren't created bt the compiler. Control Flow Integrity.

when ROP attacks got harder, the hackers move to data oriented attacks, rather than violating the control flow, the attacking code modifies the data used.\
(examples of CVE and demo of using static and dynamic analysis).

address sanitizer - recompile with compiler flag, adds "hooks" to memory allocation functions and controls the dynamic memory allocation. we link the program with the memory sanitizer, so we need the program to actually cover the code itself properly (code coverage).

changes to the Address sanitizer:

- interoperability in windows
- continue on-error
- compatiblity with DLLs

(another demo - GCC compiler)

future stuff to memory safety.

</details>

### What's New in Conan 2.0 C/C++ Package Manager - Diego Rodriguez-Losada

<details>
<summary>
new changes in Conan 2.0 package manager, based on feedback from the c++ community.
</summary>

[What's New in Conan 2.0 C/C++ Package Manager](https://youtu.be/NM-xp3tob2Q), [slides](https://github.com/memsharded/whatsnew2/blob/main/Whats_new_Conan_2_0.pdf)

increase in popularity of Conan, more downloads, more PRs, more activity on slack and more support tickets.

five lessons:

1. Learning to fly
2. Cats are a good default (or not?)
3. Building a dam
4. Repeating yourself
5. Dying of a thousand bites

#### Learning to fly

conan file recipes, dependencies, cmake stuff integration `set_property(TARGET math::math ...)`

```sh
git clone ... game & cd game
conan install .
cmake ...
```

but in 2.0, there are optional argument to express dynamic linking (not just static library), propagation of linkage requirements.

#### Cats are a good default (or not?)

each package has a unique id based on the binaries, there is something about using the Semvar, and about when the major version is updated. embedding and not embedding package code into the binary.

`conan graph build-order --requires=<> --build=missing`

this controls if we need to rebuild libraries or just re-link them.

#### Building a dam

Deployers: deployment scripts, not just for programmer who build locally, also for publishing. so in that case, we can have all the dependencies packaged into the release for the final release (or for a CICD pipeline).

in conan 1.0, this was part of the recipe, and in 2.0 it's an independent thing.

#### Repeating yourself

having a reproducible build process, which is completely frozen in time. creating a lockfile
`conan install . --lockfile-out=game.lock` and using the lockfile `conan install . --lockfile=game.lock`.

#### Dying of a thousand bites

framework on top of the conan commands.

```py
@conan_command()
def hello(conan_api, parse, *args)
  """
  Simple hello command
  """
  msg = "Hello world!"
  conanOutput().info(msg)
  return msg
```

extending the conan python API, so creating custom scripts is simpler. we can install those commands directly into the conan cli.

#### Conclusions

major changes to conan in 2.0, graph, package_id, deployers, lock files, custom command and the enhanced python API.

</details>

### GitHub Features Every C++ Developer Should Know - Michael Price

<details>
<summary>
Some cool things about Github.
</summary>

[GitHub Features Every C++ Developer Should Know](https://youtu.be/Mo8MeVzzdE8), [slides ](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/GitHub-Features-Every-C-Developer-Should-Know.pdf)

C and C++ are still in the top ten used languages on github,

ISO C++ 20220 Survey results:

- Visual Studio and Visual Studio Code are the most popular editors
- Pain Points:
  - managing library dependencies
  - setting up a CI
  - creating a new dev environment
- About half of the users use Cloud CI/CD
- About half of the users don't use the cloud at all

#### Automation

GitHub Action Components:

**Workflows** are the top level for jobs, they are processes the run one or more jobs, defined by yaml files in the repository, can be triggered manually, on schedule or by an event.\
**Events** trigger a workflow, there are events for branches, PR, issues, projects, discussions, releases, forks, watches/stars and so on.\
**Job** are step that are executed on a _runner_, they can be parallel or serialized.\
**Actions** are re-usable step components: Docker, JavaScript(TypeScript) or composite, they can be published in github marketplace.\
**Runners** are execution environments for jobs, they can be github hosted VM or self hosted.

Runners are the cost of github action, both in money and time. the pay is for computation and storage. each tier of github offers a different amount of free runners credit.

#### Developer Environments

vscode.dev, github.dev:\
vscode.dev runs in the browser, while github.dev is dedicated to integration with github. they are limited to what we can run in the browser.

Github CodeSpaces are on-demand, container-based cloud development environments, they have persistent state across work sessions, and they offer some portability between VMs. we can also pre-build containers from github actions to save time when creating CodeSpaces (it can take a few minutes otherwise).

#### Collaboration

> Github &larr; Git &larr; Linux &larr; Unix + C

git was designed to support linux development, it was based on _pull_ model. but github uses the _push_ model when using **forks**.

> Issues
>
> - Easy to file
> - Decomposable
> - Labels for organization
> - Automatable
>
> Projects
>
> - Long-term planning
> - Flexibility
>
> Pull Requests
>
> - Code Review
> - Gated merges
> - Access control
> - Automatable

#### Demo

> Adding continuous integration and fixing bugs in a cross platform C++ project without cloning the repository locally.

creating an action:\
<kbd>Actions</kbd>, configuring a default cmake action yaml file. we can <kbd>view runs</kbd> and trigger the workflow with <kbd>run workflow</kbd>. next we add a workflow for code analysis.

workflow fields:

- `on` - when is it triggered
- `jobs.jon.runs-on` - which enviornment this run on.
- `jobs.job.strategy.matrix` - setting up a multi dimensional array, combining values in a cartesian product.
- `jobs.job.steps.*.run` - the action in the step
- `jobs.job.steps.*.uses` - use a different action from another workflow

we can see in the demo how the code scanning shows issues in the code. and we can use the CodeSpaces environment and the github extension to view the issues directly.

cherry picking commits `git cherry-pick <branch>`

#### Visibility and Openness

github community, github public road-map.

</details>

### New in Visual Studio Code! Clang-Tidy, makefile, CMake, GitHub & More - Marian Luparu, Sinem Akinci

<details>
<summary>
Some demo of features in Visual Studio Code.
</summary>

[New in Visual Studio Code! Clang-Tidy, makefile, CMake, GitHub & More](https://youtu.be/iTaOCVzOenM) , [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/VSCode-session-CppCon2022-upload.pdf)

Demo 1 - starting from scratch on macOs, getting a working project from a repository. using CmakePresets, a configuration json file, we set a tool chain using vcpkg. we create a vcpkg manifest file and list the dependencies.

- sticky headers - showing the name of the function in the editor at all time, allowing to quickly jump to the top of long functions.
- doxygen comments integration and auto generation
- inlay hints - showing parameter hints (names), or the type of `auto` variables.
- clang tidy integration - code analysis, formatting.
- debug configuration templates.

Demo 2 - starting from scratch on a windows machine and WSL. windows requires installing C++ compiler (mingw).\

- conditional dependencies based on platform.
- automatic connection to wsl from vsCode.
- liveShare - shared collaboration environment (including co-authorship in commit messages)
- debugging - break on value change, extra customizations.

</details>

### Personal Log - Where No Init Has Gone Before in C++ - Andrei Zissu

<details>
<summary>
Hiding the log messages in production code.
</summary>

[Personal Log - Where No Init Has Gone Before in C++](https://youtu.be/0a3wjaeP6eQ), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Personal-Log_-Where-No-Init-Has-Gone-Before.pdf), [github demo](https://github.com/cppal/hashed_logger)

Talk starts with this code:

```cpp
#include <iostream>

void f() {
  DO_ON_INIT(std::cout << "Let's see if I can print my line number: " <<
  __LINE__ << '\n');
}

int main(){
  return 0;
}
```

the line is printed, even though we don't call the `f()` in the main function, so we will now work towards creating the macro.

#### Implanting `DO_ON_INIT`

Removing log strings from shipped binary data. we don't the logs the client easily sees to be clear, as we don't want to tell hackers where we do sensitive stuff, such as password checking. but we do want the logs to be retrievable.

we could try hashing, there are many functions we can find online.

```cpp
template typename<String>
constexpr std::size_t hash_string(const String& value, std::size_t seed){
  std::size_t d = (0x811c9dc5 ^ seed) * static_cast<std::size_t>(0x01000193)
  for (const auto & c: value)
    d = (d^ static_cast<std::size_t(c))  * static_cast<std::size_t>(0x01000193);
  return d >> 8;
}
```

but constexpr isn't enough, it "may" run in compile time, but we need to force it to do so.

```cpp
#define FORCE_CONST_EVAL(expr) std::integral_constant<decltype(expr), (expr)>::value
#define LOG(MSG) std::cout << FORCE_CONST_EVAL(hash_str(MSG))<< '\n'
```

but how do we get them back? we want to read our logs strings! hash functions are one way only. we need a decoder. we want to collect all the logged strings, without invoking them, before anything else happens.

doing things automatically in C++ usually means either a global objects or static data members. our first attempt is creating a static class object.

```cpp
#include <iostream>

void f() {
  struct InitExec{
    struct Impl {
      std::cout << "Let's see if I can print my line number: " <<  __LINE__ << '\n';
    };

    static Impl impl; // ERROR!  can't have static data members in local classes.
  };
}

int main(){
  return 0;
}
```

this fails, because we can't have static data member in locally defined classes. the next attempt is to make it a template class.

```cpp
#include <iostream>

template<size_t N>
struct InitExec{
  struct Impl {
    std::cout << "Let's see if I can print my line number: " <<  N << '\n';
  };

  static Impl impl;
};


void f() {
  InitExec<__LINE__> reg;
}

int main(){
  return 0;
}
```

this builds, but doesn't execute the code. it's optimized out. unused template instantiation are removed. we would need to force a side effect to happen so that it will be kept. this starts with a linker error, but we can get it to pass.

```cpp
template<size_t N>
typename InitExec<N>::Impl InitExec<N>::impl; // make linker happy!

void f() {
  std::cout << & InitExec<__LINE__>::impl;
}
```

this is actually executed. so let's stop printing to cout;

```cpp
void f() {
  (void)&InitExec<__LINE__>::impl; // discard result
}
```

we managed to pass a small problem, we can print the line number, but we want to print any arbitrary string data, so we need a template parameter, just like N. the first attempt of using a lambda doesn't compile.

```cpp
#include <iostream>

template<typename P>
struct InitExec{
  struct Impl{
    Impl() {P();}
  };

  static Impl impl;
};


void f() {
  auto p = [](){std::cout << "Let's see if I can print my line number: " <<  N << '\n';};
  (void)&InitExec<p>::impl; // discard result
}

int main(){
  return 0;
}
```

but if we change somethings, we can overcome the problem. we just need the lambda to have a concrete type that is suitable to be used at non type template parameter (NTTP).

```cpp
#include <iostream>
using void_fn_t = void(*)(); //

template<void_fn_t P>
struct InitExec{
  struct Impl{
    Impl() {P();}
  };

  static Impl impl;
};

template<void_fn_t P>
typename InitExec<P>::Impl InitExec<P>::impl;

void f() {
  constexpr void_fn_t p = [](){std::cout << "Let's see if I can print my line number: " <<  N << '\n';};
  (void)&InitExec<p>::impl; // discard result
}

int main(){
  return 0;
}
```

this works! now we wrap it into something usable

```cpp
#define DO_ON_INIT(...) \
{ \
constexpr void_fn_t fn_on_init = [](){__VA_ARGS__;}; \
(void) &InitExec<fn_on_init>::impl; \
}
```

and we got to the first example!

doing this in C++20 is even easier. we can use lambda as regular type parameter.

```cpp
template<class T>
struct S {
  S(T) { (void)x; }
  static inline int x = T{}();
};

#define DO_ON_INIT(...) S([]{__VA_ARGS__; return 0;})
```

#### The Decoder Tool

the production code doesn't have the original strings, but the source code does. we need to map the string hashes to the original strings. we have compile time flag, one state does the encoding, and one state creates the dictionary mapping.

```cpp
#ifdef BUILD_FOR_ENCODING
  #define LOG(MSG) std::cout << HASH(MSG) << '\n';
#else
  #define LOG(MSG) DO_ON_INIT(register_message(MSG));
#endif
```

registering a message

```cpp
static std::map<size_t, const char*> msg_reg;
auto& get_reg() {
  static std::map<size_t, const char*> msg_reg; // lazy access
  return msg_reg;
}

void register_message(const char* msg) {
  auto reg& = get_reg();
  const auto key = hash_str(msg);
  assert((reg.find(key)== reg.end() || (reg.at(key) == msg)) // make sure there are no collisions.
  || (std::strcmp(reg.at(key), msg) == 0)); // avoid string pooling
  reg.emplace(key, msg);
}

const char* GetLogMessage(size_t msg_hash) {
  return get_reg().at(msg_hash);
}
```

#### Demo & Compiler Differences

**Demo**: see code in github. we need to include the production code files (so we can't include stuff from the "main" function)

this code wouldn't work as it is in gcc or clang. or at least, it would require more hack or using a special attribute. there is a problem with static instantiation order.

</details>

### "It's A Bug Hunt" - Armor Plate Your Unit Tests in Cpp - Dave Steffen

<details>
<summary>
Designing test cases, understanding accuracy and precision.
</summary>

["It's A Bug Hunt" - Armor Plate Your Unit Tests in Cpp](https://youtu.be/P8qYIerTYA0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/its-a-bug-hunt-fix1.pdf)

> Aspects of Unit Testing
>
> - Why to write unit tests
> - Process for writing unit tests
> - Psychology of writing unit tests
> - Writing good unit test code
> - Making unit tests good tests

we draw from "Popper Falsifiability criterion", it's hard to prove that our code is correct, but we can try and fail to prove it's wrong.

> a good unit test is:
>
> - **Repeatable** and **Replicable** - you see the same answer, and your colleagues gets the same answer as you do.
> - **Accurate** - measurement are close to a specific value - "equipment is correct"
> - **Precise** - measurement are close to each other - "equipment is reliable"

accuracy for unit tests

> - True Positive: Test fails and there is a bug.
> - True Negative: Test passes because there are no bugs.
> - False Positive: Test fails but there is no bug.
> - False Negative: Test passes but there are bugs.

$Accuracy = \frac{P_T + N_T}{P_T + N_T + P_F + N_F}$

> "Tests should fail because the code under test fails, and
> for no other reason" ~ Titus Winters

Because unit testing is binary (pass/fail), we don't have some fancy math equation for precision. but we can define a precise test as giving useful data.

#### Accuracy

**Completeness**

look for bugs everywhere, but we can't have infinite number of tests and all inputs. so we need to select our test cases to maximize accuracy by maximizing the chances of finding a bug per test case.

Equivalence Class (EC)- a set of inputs that will follow the same code path. so we just need to test once case from each partition. an example is the `abs` function to determine absolute value. the possible ranges of input is [INT_MIN, INT_MAX]. we can't test them all, but we should partition them.

```cpp
int abs(int x) {return x<0 ? -x :x;}
```

the possible output is in the range of [0, INT_MAX]. there are supposedly two Equivelent classes, one for positive numbers and one for negative. but because they are ranges, we should check the boundaries.

- places where behavior changes
- places where easy mistake live (&lt; vs &le;)
- places next to these.

so if we choose our cases according to that, we decide on testing [INT_MAX,1,0,-1,INT_MIN]. and then we find the bug. there are more negative numbers in the integer domain then positive numbers. this is undefined behavior.

**Wide and Narrow contracts**:
we re-write the function to have a defined behavior, and now we can see that there are three ECs, and the contract we are testing is wider.

```cpp
int abs(int x){
  if (x === std::numeric_limits<int>::min())
    throw std::domain_error("Can't take abs of INT_MIN");
  return x<0 ? -x :x;
}
```

we next move to write test cases to handle "interesting" cases, based on our domain knowledge, zero is an interesting value, empty strings, null. we can also add test cases for values which are already in a EC, just to assure whoever reads the code that we are handling those cases correctly.

An interesting question is whether equivalence partitioning should be considered "black box" (observable behavior) or "white box" (implementation details) testing. white box testing is highly coupled with implementation, so it's usually not recommended, but without knowing the details, we can miss those ECs.

example of sorting function unit tests:

- zero elements
- one element
- two elements (some times)
- many element
- all the same values
- already sorted
- reverse

functions can also have multiple input parameters, so we should also partition for them, but ideally, we should't have to test each combination.

**Correctness**

> Correctness: does the test correctly identify:
>
> - correct output as correct
> - incorrect output as incorrect
> - Correct Error handling
>
> Correctness is part of Accuracy and also part of maintainability

testing loggers, testing floating point precision, testing standard containers.

**Validity**

beware of using circular logic, don't confuse regression/acceptance test with proving correctness.

> The result of your test matches reality of code
>
> - Test completely (Equivalence Partitioning)
> - Test Correctly (use no more information than you have, use no less than you need)
> - Test Validity: no circular logic; write falsifiable code and falsifiable tests

#### Precision

> Unit test results provide a maximum amount of useful information,
> How fast can we move from:\
> "we know there's a problem"\
> to\
> "we know what and where the problem is"

**Clarity**

use a good unit test framework. passing tests produces little or no outputs, failing test indicate what, where, how what was the expected and actual values.

**Organization**

what are the chances the test fails, what are the chances the test fails and everything else still works? if we have test that validates the standard library, it's pointless. the test won't fail, and if something in the standard library does change as to make it fail, a whole lot of stuff will also fail at the same time. even if we use a custom implementation, the test should be for the the team who writes the implementation, and not inside our test classes. if the problem isn't in our code, then the test for it should also not be in our code.

two approaches to using external tools, we can use the real thing or use a mock.

> - Your class or test uses a Thing.
>   - If Thing fails, its tests fail and your tests fail (less precision)
>   - Tests use real Things (more accurate)
> - Use a MockThing in your tests.
>   - If Thing fails, its tests fail and nothing else (more precision).
>   - Tests use non-real MockThings(less accurate)

using mocks reduces the blast radios, if something fails, it's because of the code tested, not because of external code. on the other hand, every difference between the TestDouble and the real things is a gap for bugs to come through and hide.

humans are not good at handling contradictory input, and we get worse when we are tired, angry, stressed. so out test should be clear.

</details>

### Observability Tools C++: Beyond GDB and printf - Tools to Understand the Behavior of Your Program - Ivica Bogosavljevic

<details>
<summary>
Some tools for observability of code metrics.
</summary>

[Observability Tools C++: Beyond GDB and printf - Tools to Understand the Behavior of Your Program](https://youtu.be/C9vmS5xV23A), [github](https://github.com/ibogosavljevic/johnysswlab/tree/master/talks/tools)

most people use debuggers or print statements when they interact with programs, but we can use observability tools as well. especially for unfamiliar code.

- hot and cold functions
- timeline
- memory usage
- hardware efficiency
- interactions with the operating system

#### FlameGraphs

FlameGraphs tools, such as speedScope. its a visualizer that's running in the browser and uses the sampling data generated by "pref".

1. Time Order - stack trace, calls over time. which frame was called by who and when.
2. Left Heavy - accumulative view of each function, how much time was spent at each function fame,
3. Sandwich - functions sorted by execution time.

#### Heap Profilers

heap profilers tools, such as heapTrack, tack memory consumption by monitoring allocation calls, keeping track of how many calls are made, what sizes, etc...

it also detects memory leaks, and can tell who made the calls and how long the memory was used.

#### Coverage Tools

coverage tools, such as _l-cov_, which tracks how many times each line is executed. using coverage tools requires compiling the code with coverage flags enabled. this does mean we will have slower execution due to all the added code.

this can be used to understand branching and see which branch was used more.

#### Kernel Call Traces

tools which track system calls, such as strace. we can see calls related to files, memory allocations, mutexes, and so on.

#### Gard Counters

hardware counters and event counters, the _perf_ tool. we see time spent in userspace and system space, branches hits and misses, instructions and cpu cycles, context switches, cpu migrations and page faults.

(example of matrix operations with column-wise and row-wise access)

</details>

### Generating Parsers in C++ with Maphoon - Hans de Nivelle

<details>
<summary>
A tool to create tokenizers and parser.
</summary>

[Generating Parsers in C++ with Maphoon(part-1)](https://youtu.be/Ebp5UPi9ny0), [part-2](https://youtu.be/T_uNzSP-9ik), [slides-1](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/slides01.pdf), [slides-2](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/slides01.pdf)

Maphoon is a programming language for the implementation of logic - verification of mathematical proofs.

> Q: "What is Parsing?"\
> A:
> In computer science, nearly everything can be represented as a tree. We are given some input as sequence of characters. The task of the parser is to extract the tree structure. The resulting tree is called **abstract syntax tree**.

the expression can be turned into an AST:

```cpp
if (c >= 'A' && c <= 'Z')
  c += 'a' - 'A'; // If c is upper case, make it lower case.
```

#### Tokenizing

first step is tokenizing - turning the stream of characters from the source code into atomic "tokens".

> Tokens typically have the following forms:
>
> - Numbers (integers or floating point)
> - Strings (can be quite complicated, when there are escape codes)
> - Reserved words, operators
> - Comments
> - Indentation changes (In Python, these result in tokens)

classifying tokens:

- regular expressions &rarr;
- non deterministic finite automata &rarr;
- deterministic finite automata &rarr;
- minimal deterministic finite automata &rarr;

there are existing parsers, but they don't fit the requirements. sometimes we need to write a function recognizes specific tokens when it can't be done automatically

if we have "while" in the code, when we tokenize it, it can be a reserved work or an identifier (until we see the whole word and the ending, we won't know).

flat automata, state, transistion functions, classifiers.

#### Parsing

the second step is parsing.

> Parsing:\
> We have cut the input in bite-sized pieces, and we need to build a tree from them.

context-free grammar

- non terminal symbols V
- terminal symbols &Sigma;
- set of rules in the form of $\alpha \to w$, with $\alpha \in V$ and $w$ a finite word over $V \cup \Sigma$
- a start symbol.

we have an attribute grammar with a formal defintion.

#### Calculator Example

defining the rewrite rules for a simple calculator, we need to solve the ambiguity with preferences (priority, associativity).

top-down parsing: shift, reduce, read. bottom-up parsing: viable and non-viable states (decision making).

> **Maphoon** reads the grammar and the action code.
>
> - It creates two files symbol.h and symbol.cpp containing the symbol definition.
> - It also creates two files parser.h and parser.cpp containing a
> - runnable parser that correctly applies the action code when a rule is reduced.
> - Every class that has correct life cycle operations (constructor, assignment, destructor) can be used as attribute.
> - It is even better when attributes are movable.

(prolog example) - defining operators at runtime.

</details>

### tipi.build A New C++ Package Manager - Damien Buhl

<details>
<summary>
A cloud based build package manager.
</summary>

[tipi.build A New C++ Package Manager](https://youtu.be/cxNDmugjlFk), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/DamienBuhl-CppCon2022.pdf)

"reinventing the wheel" by creating a package manager.

reasons to have package managers

- code reuse & distribution
- industrializing updates
- abstracting build complexity

avoiding ABI-mismatches with precompiled libraries. build system interoperability.

> 1. ABI compatibility cannot be computed
> 2. Packages know _too little_ about their content
> 3. Build systems rely on developers repeating themselves
> 4. code tells enough about dependencies
> 5. introducing "fast build from sources in the cloud with a global build cache"

all the "build" information is expressed through the code - namespaces, #include statements, features used... the build system replicates the information and requires someone to keep both in sync.

so the solution is to scan the code, match any dependencies and generate the build script and use cache packages from the cloud. build remotely, have both the safety of building from source and the speed of using pre-built libraries.

incremental builds, storing snapshots through revisions, supports local and remote builds/caches. requires access to the source code.

</details>

### Reproducible Developer Environments in C++ - Michael Price

<details>
<summary>
Microsoft's solutions for developer environments: DevBox and github CodeSpaces.
</summary>

[Reproducible Developer Environments in C++](https://youtu.be/9GKGp3zeOOc), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Reproducible-Developer-Environments.pdf)

qualities of a developer environemnt

- Affinity - which projects is it meant for?
- Platform - based on which platform? windows, unix,cloud provider, web browser?
- Resources - computation, storage, maybe specialized processing units and embedded devices.
- Configuration - customizing the environment to the needs
- Interactivity - not the same as a ci-cd environment, suitable for ad-hoc changes and experimentation

setting up a development environment is one of the pain points that comes up in developers surveys. right after using external libraries and setting a ci-cd chain. so all three are pretty similar to one another.

problems:

> - Procurement
> - Misconfiguration
> - Fragility
> - Genericity
> - Permanence
> - Inconsistency

#### Evaluating Possible Solutions

we have some criteria to evaluate solutions

1. Reproducible - same for all users
1. Purpose Built - fits the needs of the project
1. Managed - controlled by IT for security and follows regulations
1. Isolated - doesn't effect other, limited blast radius
1. Available - requires network?
1. Customizable - leave room for developers differences
1. Documented - extensive, up to date?
1. Versioned - projects change, and we need to be able to return to previous points in time
1. Procurable - how fast to get
1. Responsive - how does it feel to interact with

no solution will be perfect, but we can compare them and rank them according to our organization priorities.

#### Microsoft DevBox

> Provide developers with self-service access to high-performance, cloud-based workstations
> pre-configured and ready-to-code for specific projects

built on top of Azure windows VM, allows for any number of machines, any number of configurations, IT can set up policies, security settings, etc.. the developers manage the configuration, and then log-in to use with a separate DevBox-portal and RDP (remote desktop).

we pay for storage all the time, but pay for compute only for what we use.

currently in public preview.

#### Github CodeSpaces

> - On-demand, container-based, cloud development environments
> - Persistent state across work sessions
> - Limited portability between different VM SKUs
> - Customizable environments with dev containers
> - Can be prebuilt from GitHub Actions

built on top of docker, highly documented. linux based, uses visual studio code tools.\
no graphics, but we can have it direct gui to our local machine through ssh.

</details>

### The Surprising Complexity of Formatting Ranges in Cpp - Barry Revzin

<details>
<summary>
Understanding how the 'fmt' library might handle formatting ranges.
</summary>

[The Surprising Complexity of Formatting Ranges in Cpp](https://youtu.be/EQELdyecZlU), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/The-Surprising-Complexity-of-Formatting-Ranges.pdf)

we started with `printf`, where the specification format determines the type. the problem is that we can't extended the format, so if we have a new class, there is no way to decide how to format it.

```cpp
std::printf("The price of %x is %d\n",48879, 1234);
```

then came _iostreams_, which made thing composable and has pieces. the problem is that some manipulators are sticky and some are just for one "thing". but it does support custom formatting.

```cpp
std::cout << "The price of"
  << std::hex
  << 48879
  << " is "
  << 1234
  <<'\n';
```

`{fmt}` is the next step, it's similar to to `printf`, but it does allow to reference the same value at multiple replacement fields, without relying on the order.

```cpp
std::print("The price of {:x} is {}\n", 48879, 1234);
std::print("The price of {0:#x} is {0}\n", 48879);
```

- replacement field - the entire thing (the curly braces)
- format specifier - things after the colon (":") that determine the format
- arg-id - number before the colon that references the argument postion

we can extend the formatter.

#### Parsing `{fmt}`

```cpp
template<class charT>
class basic_format_parse_context {
public:
  using char_type = charT;
  using const_iterator = basic_string_view<charT>::const_iterator;
  using iterator = const_iterator;
  constexpr const_iterator begin() const noexcept;
  constexpr const_iterator end() const noexcept;
  constexpr void advance_to(const_iterator it);
  constexpr size_t next_arg_id();
  constexpr void check_arg_id(size_t id);
};
```

the format string can also be complicated.
the `*` symbol is fill, `^` alignment, and we take the width from the argument (10). or using arbitrary characters which fit the type.

```cpp
std::print("{:*^{}}\n"," hi", 10);
// print ****hi****
std::print("{:%Y-%m-%d %H:%M}\n", std::chrono::system_clock::now());
// print the time
```

the context begins at the `:`, but the ending depends on the parser. we can get an empty context, or something else.

(example of writing a formatter)

getting the formatting from the caller, creating generic formatter for `std::optional`.

#### Formatting ranges

the default is square brackets and commas, which works for most cases, but not for strings. and we would want sometimes to describe relationship between elements (such as key value pairs).

we need a format spec for ranges. we get this by adding colons, each colon is a different nested thing. so for `std::vector<std::vector<char>>` we can have `std::print("{:n:*^18:#x}\n",v);`

- `:n` - top level - vector of vectors, no square brackets
- `:*^18` - middle level - vector of characters, center aligned with a width of 18 and filled with "\*" symbols
- `:#x` - deepest level - char objects. hex formatting with leading "0x".

the next problem we have is with `filter` function and non-const iterable. this was a real problem with the ranges library in C++20.

top level specifiers, we can't write more than once and we can't read more than once.

(looking at how the C++ Standard Library does it)

delimeter

we could have something like this, with `join`

```cpp
int main() {
vector<uint8_t> mac = {0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff};

print("{}\n", mac); // [170, 187, 204, 221, 238, 255]
print("{::02x}\n", mac); // [aa, bb, cc, dd, ee, ff]
print("{:n:02x}\n", mac); // aa, bb, cc, dd, ee, ff
print("{:02x}\n", join(mac, ":")); // aa:bb:cc:dd:ee:ff
}
```

#### Formatting tuples

let's say we have a pair, we need to specify a format for each element, then we need to find the boundary between them.

this is a mess with no good solution.

but it was adopted into C++23 'fmt' library implementation, and it's still not done and doesn't support everything.

</details>

### How to Use C++ Dependency Injection to Write Maintainable Software - Francesco Zoffoli

<details>
<summary>
Overview of dependency injection and how to achieve it.
</summary>

[How to Use C++ Dependency Injection to Write Maintainable Software](https://youtu.be/l6Y9PqyK1Mc), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/dependency-injection-1.pptx)

#### Dependency Injection

- observable behavior
- intended interface
- usage

users expect interfaces, and program against it. we pass the dependencies as parameters. this allows us to reason about the components and code in a local manner. without having global view of the system.

the way we use the dependencies imposes constraints on the user and one the using components.
life time management

> - Never instantiate your own dependency
>   - dependencies should be passed as parameters
> - Don’t force the method of lifetime management of the dependency
>   - take the dependency as a reference
>   - no smart pointers

#### Types of Dependencies

> - Data - Dependency on some information
> - Behaviors - Dependency on some action

##### Data

Data should always be read-only, we don't want it to be changed by some other component, as that will mean an implicit dependency. data should be explicit (structured types, schematized), and preferably a value-type.\
The data should have a shape interface, rather than concrete type. there is a tradeoff.

| shape              | example                      | performance | freedom | notes                      |
| ------------------ | ---------------------------- | ----------- | ------- | -------------------------- |
| concrete type      | <cpp>std::vector</cpp>       | 5/5         | 1/5     | limited                    |
| view object        | <cpp>std::span</cpp>         | 5/5         | 2/5     | only contigious containers |
| polymorphic object | <cpp>ranges::any_range</cpp> | 1/5         | 5/5     | any sequence               |

the other option is to use templates, when static polymorphism is possible. there is also problem with composition, a span can substitute for a vector, but a span of spans cannot substitute for a vector of vectors.

making data constant. using const reference, and not changing the data through it.

for data that might change during the runtime:

- should still be read only, it can change only from the outside.
- <cpp>std::atomic_ref\<const T></cpp> - copies for each access
- <cpp>const folly::Synchronized\<T>&</cpp> - lock for access (risk contention)
- <cpp>folly::Observer\<T></cpp> - no lock, exposes snapshots (decouples contention, higher costs)

the advantage of having this approach (over fetching the data directly) is that it makes testing easier, and it's much simpler to change the source of the data.

##### Behavior

> - Stateful effect
> - Get data on-demand
> - Perform pure computation

can be single action or a bundle of action

**Runtime Polymorphism** - <cpp>std::function</cpp>, <cpp>std::function_ref</cpp>

```cpp
using ConfigLoader = function<Config(ConfigPath)>;
void doStuff(ConfigLoader loader);
doStuff([](ConfigPath path) { … });
```

**CompileTime polymorphism** - <cpp>concept: std::invocable</cpp>, <cpp>std::invoke_result</cpp>.

```cpp
template<class T>
concept ConfigLoader =
  invocable<T, ConfigPath> &&
  same_as<Config, invoke_result_t<T, ConfigPath>>;

void doStuff(ConfigLoader auto&& loader);
doStuff([](ConfigPath path) { … });
```

for a bundle of actions, in runtime we will have an abstract base class with virtual functions. we need to avoid concrete base classes, so no data members or side effects in the constructor.

```cpp
class ISender {
  virtual ~ISender() = default;

  virtual Receipt sendWithPriority(Package) = 0;
  virtual Receipt sendCheapest(Package) = 0;
};

```

for CompileTime, we against use concepts, and now we don't have to to inherit from an abstract base class.

```cpp
template<class Sender>
concept ISender = requires (Sender& s, Package p) {
  { s.sendWithPriority(p) } -> same_as<Receipt>;
  { s.sendCheapest(p) } -> same_as<Receipt>;
};

class MockSender {
  MOCK_METHOD(Receipt, sendWithPriority, (Package), ());
  MOCK_METHOD(Receipt, sendCheapest, (Package), ());
};
static_assert(ISender<MockSender>);
```

#### Passing Dependencies

we can pass the dependencies to either the constructor or the function, but we would want to avoid setter methods.

> - They enable cycles
> - Confuse lifetimes
> - Force to handle not set
> - Force to handle change at runtime

when we pass through the constructor, we should avoid coupling the dependencies with the component.

- problems with <cpp>std::unique_ptr</cpp>
- using <cpp>std::shared_ptr</cpp> forces us to share lifetime management.

when we pass through function arguments.

> Use Cases
>
> - Functions
> - Caller controls dependency instance
> - Limited lifetime or changes across invocations

we would like to avoid piping the dependency internally, as that could require us to have concrete type instantiated.

##### Factories

> Factory: construct a dependency from parameters
>
> - Factories are dependencies too
> - Just return instances of dependencies

the return type must be abstract (except for static polymorphism), so our options are smart pointers (unique or shared), which have the risk of dangling references.

#### Abstracting Components

we got ourself into a mess, so we want something more simple to use, without forcing the user to know so much about the wiring.

we can use a Holder struct, which takes the dependencies and performs the wiring, only exposing the intended interface.

</details>

### Cross-Building Strategies in the Age of C++ Package Managers - Luis Caro Campos

<details>
<summary>
Using Conan for cross platform building.
</summary>

[Cross-Building Strategies in the Age of C++ Package Managers](https://youtu.be/6fzf9WNWB30), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Cross-building-strategies-in-the-age-of-C-package-managers-publish.pdf)

there many ARM devices (smartphones, boards, apple chips), but most development is done on x86_64 machines.

we could build our code on the target device, but not all devices have build tools, and the compile time can be horrible at weaker machines. it's easier to have a cross-building toolchain as part of the continues integration process.

most simple cases "just work" for different compilers and different architecture targets. but in other cases
cross compiling can be a mess.

managing dependencies is always a top pain point for c++ developers.

> GNU convention:
>
> - Build - The machine we are building on
> - **Host** - The machine we are building for
> - Target - In the context of building a compiler, the machine that compiler produces code for - (not the topic of this talk)
>
> Different tools use different conventions:
>
> - Conan: same as GNU convention (build and host)
> - vcpkg:
>   - **"Host"** - the machine we are building on
>   - "Target" - the machine we are building for
> - CMake:
>   - `CMAKE_HOST_XXX`: variables relevant to the machine CMake is running on.

the term **Host** refers to different things.

the build system needs to be aware of the cross-building scenario during the entire run. we need to make sure it knows about the building machine and the target.

sometimes our dependencies aren't just libraries, they can be executables for the build machine, or the target machine, so we need to know and build correctly. some libraries are coupled with executables, such as ProtoBuf and Qt tools. sometimes the executable needs to interact with both types of libraries.

</details>

### Nobody Can Program Correctly - Lessons From 20 Years of Debugging C++ Code - Sebastian Theophil

<details>
<summary>
General workflow for finding and fixing bugs
</summary>

[Nobody Can Program Correctly - Lessons From 20 Years of Debugging C++ Code](https://youtu.be/2uk2Z6lSams)

general debugging strategies, not focusing on tools.

> What is a bug?
>
> - your program has a **specification**: _Implicit_ (usually) or _explicit_ (rarely).
> - Your program has a **bug** if its behavior does not conform to the specification.
> - A **bug report** describes the _symptom_ of a bug, not the bug itself.

bugs can be crashes (with or without a coredump), an error message, a line in the log, or any observable behavior. performance can also be a bug.

#### Bug Report

some of them are hard to notice, especially if they happen on the clients machine. we would want to increase the number of the "bug reports", we can notice bugs before they reach the client by having them happen during compilation.

- type checking
- `static_assert`
- `constexpr` evaluations.

we should have testing as part of the development stage, with unit tests and other forms of automated testing.\
For runtime, we should employ strict error checking, check all API return values and report them, asserting pre and post conditions, both for common "error" states and for unexpected behavior.

**Learning From a Single Occurrence**

a rare error that happens once is an opportunity to catch it before it becomes a common case. even if we can't fix it, we can be ready for the next time, with better logging, stricter invariants enforcements.

example of using a synchronization method that works well between processes, but not between threads (file lock).

#### Reproduction

levels of reproducible bugs - does it happen in debug build as well? does it happen on all machine? can we find it in an interactive debugger session?

some tools that help with detecting hard-to-reproduce issues:

- address sanitizer
- thread sanitizer
- undefine behavior sanitizer

when a bug only appears on some machines, it might mean some other tools is interfering, or that the specific machine is very busy or very strong.

if everything fails, and we can't fix the issue, we can still write code to report the problem, write better error checking, try getting more examples of failures.

example of a problem between 32bit and 64 bit processes.

#### Identify the Problem

the symptom isn't the bug itself, it's just how it shows itself. using interactive debuggers, tracing function calls (which might make timing dependant bug disappear). advanced debuggers allow tracing breaking points, and those debuggers also have APIs so commands can added programmatically. data visualizers.

if we suspect the problem is with the interaction with the OS, there are logging tools which monitor those calls:

- Windows: ProcessMonitor
- MacOS: dTrace
- Linux: sTrace

if we have a new bug, which we know is new and didn't happen before, we can look for the breaking change and find when the bug was created, we should also understand what was the purpose of the changes, so we won't introduce the same problem the original change was meant to solve.

one debugging technique is to improve the code, adding asserts, modernize legacy code if possible (smart pointers, RAII).

there are also reverse debugging tools, which allow to step backwards in time during debugging. this requires being familiar with the debugger, if we don't have the source code itself, we can still set breakpoint into function calls and look at the assembly code.

the most important thing is to question your own assumption.

#### Classify the Bug

small bugs: the code doesn't do what you meant. this usually means a small fix, and some cleanup to find similar possible issues that might be hiding in the codebase.

- uninitialized data
- memory management data
- stack corruption
- data corruptor through missing locks

large bugs: your mental model was wrong. this might require a larger change, rethinking about the problem.

- not understanding the specifications
- wrong use of internal and external api
- using OS facilities that don't work like you thought

is dereferencing a null pointer a small bug that we don't check the pointer, or is it a large problem that we allow null pointers into the system? in some cases it's hard to decide.

#### Fix The Bug

the general practice is to commit the smallest possible fix - solving the problem with the least chance of introducing new bugs. but that might not solve the root issue, it may only hide one instance of the problem, and ad-hoc fixes tend to accumulate over time, reduce code quality and make code harder to understand.

- ship a fix quickly
- attempt a a through fix in development branch

understand why the bug happened at the first place, was something missing from the helper libraries? is there a bad practice that people still use? can we add something to the codebase to prevent this happening again? maybe we can introduce an abstraction and use it across the whole codebase.

document the fix, make sure the fix is connected to the original bug issue so it won't be deleted by accident, add tests to find and maintain the fix.

</details>

### Linux Debuginfo Formats - DWARF, ELF, dwo, dwp - What are They All? - Greg Law

<details>
<summary>
looking at debug information with various tools
</summary>

[Linux Debuginfo Formats - DWARF, ELF, dwo, dwp - What are They All?](https://youtu.be/l3h7F9za_pc), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/greg_law_cppcon_2022.odp)

the `-g` flag, inserting debugging info into the executable binary. it doesn't go into the executed process, but it effects the size of the binary on disk.

```sh
gcc -g3 hello.c
readelf -WS a.out # view binary sections
readelf -m  a.out # read notes
objdump -j .interp -s  # specific section
gcc -static -g3 hello.c # no .interp section
```

(audience question about addresses during runtime compared to what appears on the binary)

plt - has something to do with dynamicly linked objects (such as the standard library), "got" is global offset table, bss is zero initialzed values. if we touch data in the ro (read only) section and try to change it, we get a segmentation fault.

_addr2line_ takes the debug information and translates the stack pointer to the source file. this is done even outside of gdb.

```sh
gcc -g hello.c
gdb a.out
(gdb) p $pc # print program counter in gdb
addr2line -j .text a.out
```

_dwarfdump_ can also dump some information, such as type information and addresses.

```sh
dwarfdump -a.out
```

adding debug information increases the link time. _split dwarf_ flag generates a debug information file, so it decouples the information and reduce how much of it goes into the executable.

```sh
gcc -g -gsplit-dwarf hello.c
```

there are other utilities we can use, such as "eu-readelf" or "eu-objdump" - elf utils (not related to europe), those are lighter versions because they removed support for other formats.

- ELF - executable and linkage file
- DWARF - debugging with attributes (something)
- BFD - binary file description (unofficially "big fucking deal")

there is a tool that serves debug info over http (like microsoft symbol server), developed by redhat.

```sh
sudo apt install debuginfod
debuginfod -F . # serve debug info
DEBUGINFOD_URLS=localhost:8002 gdb a.out # get debug info
```

</details>

### New in Visual Studio 2022 - Conformance, Performance, Important Features - Marian Luparu & Sy Brand

<details>
<summary>
New stuff in Visual Studio
</summary>

[New in Visual Studio 2022 - Conformance, Performance, Important Features](https://youtu.be/vdblR5Ty9f8)

#### Conformance

visual studio runs natively on Arm64 applications, with native compiler and vcpkg support.

Now has the complete set of c++20 features, some c++23 features already in preview mode. adding c++20 integrations to all layers of visual studio (code navigation, linters, intelliSense, debugging). module support and experimental mode wih CMake.

(demo) - text adventure game, using vcpkg- manifest mode (declarative style), cmake configuration presets with a toolchain pointing to vcpkg. asset files with the "ixx" extension. showing the usage of <cpp>std::expected</cpp>, visual studio in now aware of "refactoring", so if we change a function signature, we can still find who was using it previously and go to fix them directly from the navigation "go-to" tool-tip and not just through the "errors" pane. showing the c++23 <cpp>deducing this</cpp> feature.

```cpp
template <class Self>
constexpr const _Ty& value(this Self&& self){
  if (_Has_value) {
    return std::forward<Self>(self).Value;
  }

  _Throw_bad_expected_access_lv();
}
```

(this will also help with recursive lambdas, CRTP, etc...)

#### Performance

some examples of how source code is better compiled into more efficient assembly code.

in the example below, the condition doesn't change during the loop, so we have an unnecessary checking of it, which results in extra work in branch predictions.

```cpp
void loop_if_un_switching(int cnd) {
  for (int i = 0; i < 1024; ++i) {
    if (cnd)
      foo(i);
    bar(i);
  }
}
```

we would want to "hoist" the condition out of the loop, so it would be something like this, but we don't want to write code like that, we want the compiler to know how to do this.

```cpp
void loop_if_un_switching(int cnd) {
  if (cnd) {
    for (int i = 0; i < 1024; ++i) {
      foo(i);
      bar(i);
    }
  } else {
    for (int i = 0; i < 1024; ++i) {
      bar(i);
    }
  }
}
```

another example with chained min/max calls. improvements on backwards vectorization. better optimization on byte swap code.

#### Features of Importance

- inline hints for parameters and <cpp>auto</cpp>
- "go to defintion/deceleration" after signature change
- context aware autocomplete for enum types

speed improvement for search and navigation scenarios, and also for other stuff (indexing, opening the IDE)

- extra support for unreal engine (more upcoming)
- static analyzers and improved checkers
- cmake integration as a first-class project system, basic module support
- targeting linux distros, including WSL-2, SSH interactive terminal.
- MacOs support - remote debugging

adding msvc execution and static analysis support on compiler explorer.

- over 1900 open source package on vcpkg
- toolchains for embedded projects, hardware breakpoints, monitoring.

(demo) - enum intelliSense, option to stage a single line, comparing between branches directly, fuzzing on cmake with the fuzzing sanitizer.

```cmake
add_executable(parser_fuzzer "fuzzer_entry.cpp" "parser.cpp" "parser.ixx")
target_compile_options(parse_fuzzer PRIVATE /std:c++later /fsanitize=address /fsanitize=fuzzer /Zi)
target_link_options(parse_fuzzer PRIVATE /INCREMENTAL:NO)

```

which will call this code in "fuzzer_entry.cpp" file

```cpp
extern "C" int LLVMFuzzerTestOneInput(const char *date, long long size) {
  jork::parse_command(std::string(data, size));
  return 0;
}
```

Experimental upcoming diagnostics feature, generating SARIF (diagnositics) files and exploring them in the IDE and navigating the source code based on them with an extension.

</details>

##

[Main](README.md)
