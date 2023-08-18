<!--
// cSpell:ignore objdump
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

# Core C++ 2023

<!-- <details> -->
<summary>
Israel C++ Convention.
</summary>

[playlist](https://www.youtube.com/playlist?list=PLn4wYlDYx4bs0p9S6aFvKaASoCLFVwt_E), [Schedule](https://corecpp.org/schedule/).

## David Sankel :: The Semicolon is a Lie

<details>
<summary>
Historical tour of how programming evolves and moves away from directly translating source code into machine code.
</summary>

[The Semicolon is a Lie](https://youtu.be/ICf_6L1kJcE)

programming history, David's history and how they intersect.

> 1. Computers are fast
> 2. Programming is an illusion

**TI 99/4A computer**

Texas instruments old "computers" came with manuals of the code to type into it and then you could run the "games".

| Metric           | TI 99/4a  | 386 sx      | Pentium       | Z600                                     |
| ---------------- | --------- | ----------- | ------------- | ---------------------------------------- |
| Ram              | 16 Kb     | 4 Mb        | 16 Mb         | 24 Gb                                    |
| Registers        | 16 bit    | 32 bit      | 32 bit        | 64 bit                                   |
| CPU              | 3 Mhz     | -           | 133Mgz        | 2.64Ghz                                  |
| Memory           | -         | 40Mb        | 500Mb         | 1Tb                                      |
| Speed Over human | 3,829,787 | 101,333,333 | 2,537,000,000 | 72,090,000,000 (without multi-threading) |

> **semicolon** - A mark (`;`) of punctuation, indicating a greater degree of separation than the comma.

**1947** - Assembly language, by _Kathleen Booth_ and her husband _Andrew Booth_, created for the A.R.C machine, in preparation that the same instructions could carry over to more modern machines as they become available. **1951** - abstraction from a machine that creates instructions to a language that creates a the instructions. later we got _Grace Hopper_ and the A-0 system (arithmetic Language version 0), which laid the foundation for the first compiler. next we meet _John Backus_ and **Fortran**, which introduced the optimizing compiler, and since then we no longer directly translate source code into machine code, we have something that changes it. _Frances E.Allen_ introduces graph theory in the sixties.\
_Dennis Richie_ and _Ken Thompson_ creating the C language to work on the Unix operating system, later on _Bjarne Stroustrup_ which borrowed from **Simula** and BCPL and created **C++**.

When Pentium 4 were released, the "NetBurst" architecture, instructions execute at the same time, using something called "shadow registers". so it's not only that the compiler modifies the source code into something else, the CPU also modifies the assembly code.

</details>

## Bryce Adelstein Lelbach :: AI-Assisted Software Engineering

<details>
<summary>
Using AI to change how we write code.
</summary>

[AI-Assisted Software Engineering](https://youtu.be/9P0PN29VrfY), [C-Why github](https://github.com/plasma-umass/cwhy).

Large Language Model (LLM), using text context (widnow of text), which can be dropped after a while. so we need to know how to interact with the model.

Neural networks are the building blocks of AI, not the end-all type. we are moving towards more advanced model, such as neural network agents, information retrieval and execution environments.

### What Can We Do With AI

Creation and Analysis. Creation is the process of answering open-questions, creating new code, there is no "right" answer. Analysis is the next stage, reviewing code, fixing errors, re-factoring. these are questions that have an answer, this is a task that is easier for AI to do.

one example is the **C-why** tool which explains why code compilation fails. it takes a diagnostics output and tries to understand it.

> - Classification - what tools are involved?
> - Source Identification - what files or resources do we need to access?

the cycle is:
1. Analyze what we have.
2. Determine what else we need.
3. Collect what we need.

these are series of tasks, so it's suitable for AI tools. data extraction can be text only, but it's better to have code entities (such as function + comments).

We would want the tool to be integrated with the automation CI-CD pipeline, and to run when it fails. it will analyze the diagnostic check, gather the data, and create a suggestion for a fix and re-run tests on the patch.

</details>

## Noam Weiss :: Virtual Templated Methods

<details>
<summary>
Virtual Templated Function don't actual exist, but we can get around it.
</summary>

[Virtual Templated Methods](https://youtu.be/Z-WzYbTm8k0)

> the use case:
> - I want to decouple my debugging/logging by using dependency injection.
> - but I also want to support types that I don't know in advance.
> - Templated virtual functions would have been great.

our other options would be:
1. using inheritance instead of templates.
2. Break into two functions:
   1. use template to reduce to a common type.
   2. use virtual function on the common type.
3. Combine both approaches:
   1. use template to create the derived class.
   2. use the virtual function with the base class.
4. Reflection (if we implement it ourselves).

</details>

## Tomer Vromen :: Linker & Loader: The Build Process After-Party

<details>
<summary>
Some Stuff About the Linker and the Linking Process.
</summary>

[Linker & Loader: The Build Process After-Party](https://youtu.be/xc23weUlZ0A)

Linkage errors aren't the same as compiling errors, the compiler turns the source code into a machine code (object file), the linker takes all the object files and system libraries and creates the executable.

The Linker's Responsibilities are:
- Layout Code
- Layout Data
- Resolve Symbols

The gcc `-c` flag makes it so only the first step of compilation is performed, and it outputs an object file, we can then call `objdump -d` and look at the code disassembly (`objdump -t` will show the symbols).\
This includes the mangled names and the machine code instructions (we can pass it through `| c++fill` to get de-mangled names for better readability).\
the assembly code has function calls in assembly, the instuction is call, and the op-code is zero. this is because the compiler doesn't know where the actual code is, and it needs the linker to fill it in. we could also add the `--reloc` flag to the object dump and see how the code expects the re-locations should work.\
we can run the disassembler on the executable file and see how it looks after linking. now the addresses are filled out with actual locations.\
if we want to link with a library (static archive) we pass the library with `-L` path argument, we can use the `-###` flag to tell gcc to print what it actually will run, and it will show the entire command it would use, including linking the standard libraries. the linking order is sometimes important.

C++ bring some complexity to the table, function overloading and templates create name mangling. there are also C function that we need to define as <cpp>extern</cpp>, and there are inline class method defintions, which relate to ODR. there is a special memory location for <cpp>thread_local</cpp> data. we can even use unicode identifiers (🦆).

> Common errors:
>- "undefined reference"
>  - missing library object file in linkage command
>  - missing <cpp>extern "C"</cpp>
>  - wrong linkage order (it's the opposite of how `# include` works)
>  - "abi::cxx11" or "__cxx11" - libstdc++ dual ABI mismatch
>  - missing destructor, in virtual classes, must be defined, even if <cpp>~Class() =0;</cpp>
>- "multiple definitons" - probably a function defined in the header
>- "linker out of memory" - are you creating to many types?

### LTO - Link Time Optimization
all modern compilers support LTO, it requires a special flag in both compilation and linkage (so the object file keeps some information), and it's not always worth doing it.

### Share Libraries
instead of packaging the same common libraries, we can have one shared version of it in the memory and use it for all programs, but it can lead to "DLL hell". we can also have dynamic linking loader, or use <cpp>dlopen</cpp>. there is also the issue of **Wrapping/Hijacking**, we can tell the linker to call a wrapper object instead of calling the function directly,and then we can use the wrapper to redirect the calls.

> Caveats:
> - Hijacking inside a library doesn't always work
> - Hijacking non-function symbols is not officially supported
> - Hijacking class methods is complicated

we do this by passing two flags `--Wl` which instructs gcc to pass a command to the linker, and `--wrap=<mangled name>` which replaces the symbol with a symbol that is defined with the same name. (the demo didn't work so great).
</details>

##

</details>
