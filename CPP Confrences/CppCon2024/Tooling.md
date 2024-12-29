<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Tooling

<summary>
15 Talks about tooling and supports for C++.
</summary>

- [ ] Beyond Compilation Databases to Support C++ Modules: Build Databases - Ben Boeckel
- [ ] Building Cppcheck - What We Learned from 17 Years of Development - Daniel Marjamäki
- [ ] C++/Rust Interop: Using Bridges in Practice - Tyler Weaver
- [ ] Common Package Specification (CPS) in practice: A full round trip implementation in Conan C++ package manager - Diego Rodriguez-Losada Gonzalez
- [ ] Compile-time Validation - Alon Wolf
- [ ] Implementing Reflection using the new C++20 Tooling Opportunity: Modules - Maiko Steeman
- [ ] import CMake; // Mastering C++ Modules - Bill Hoffman
- [ ] LLVM's Realtime Safety Revolution: Tools for Modern Mission Critical Systems - Christopher Apple, David Trevelyan
- [ ] Mix Assertion, Logging, Unit Testing and Fuzzing: Build Safer Modern C++ Application - Xiaofan Sun
- [ ] Secrets of C++ Scripting Bindings: Bridging Compile Time and Run Time - Jason Turner
- [ ] Shared Libraries and Where To Find Them - Luis Caro Campos
- [ ] What's eating my RAM? - Jianfei Pan
- [ ] What's new for Visual Studio Code: Performance, GitHub Copilot, and CMake Enhancements - Alexandra Kemper, Sinem Akinci
- [x] What's New in Visual Studio for C++ Developers - Michael Price & Mryam Girmay
- [ ] Why is my build so slow? Compilation Profiling and Visualization - Samuel Privett

---

### What's New in Visual Studio for C++ Developers - Michael Price & Mryam Girmay

<details>
<summary>
What's new and what's coming up in Visual Studio IDE.
</summary>

[What's New in Visual Studio for C++ Developers](https://youtu.be/Ulq3yUANeCA?si=voZfhAjwwzOx_544), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/What's_New_in_Visual_Studio_For_Cpp_Developers.pdf)

The yearly talk by the Visual Studio team in Microsoft

#### Productivity

Github Copilot will be bundled with visual studio, mainly as a chatbot and editor suggestion. requires license from github. combines context from the project and opened files. we can also add context to the project with the ".github/copilot-instructions.md" file.\
We can ask copilot to improve memory layout for our classes, or to ask it to reduce the build time (with build insight).

> Build Insights - Analyze and optimize your build
>
> - Detailed analytics about your C++ builds
> - Integrated into Visual Studio
> - Visualize your include tree
> - Identify "expensive" included files
> - Find inlined functions that bloat your binaries

#### Game Development

direct support for unreal engine 5 projects, better integration and a dedicated toolbar.

#### MSVC Toolchain

improved support for C++23 features, work towards C++26 features. improvements to the address sanitizer, also integrated with copilot.\
better integration with the vcpkg package manager.

#### Debugging, Cross-Platform & Source Control

> - CMake Debugger in Visual Studio
> - Remote File Explorer for Linux
> - Target View Improvements
> - Automatically Install WSL from Visual Studio
> - Debug Linux Console Apps in Integrated Terminal

Better source control integration with popular repository hosting platforms. more copilot stuff to create commit messages.\
Better experience for connecting to remote server systems. also running tests on remote machines and modify files over there.

</details>
