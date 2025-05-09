<!--
// cSpell:ignore Spanny URDF
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Robotics and AI

<summary>
6 lectures about Robotics and AI.
</summary>

- [x] Building Safe And Reliable Surgical Robotics With C++ - Milad Khaledyan
- [ ] Making Hard Tests Easy: A Case Study From The Motion Planning Domain - Chip Hogg
- [ ] Implementing Particle Filters With Ranges - Nahuel Espinosa
- [ ] Leveraging C++ For Efficient Motion Planning: RRT Algorithm For Robotic Arms - Aditi Pawaskar
- [ ] Spanny 2: Rise Of `std::mdspan` - Griswald Brooks
- [ ] Code Generation From Unified Robot Description Format (URDF) For Accelerated Robotics - Paul Gesel

---

### Building Safe And Reliable Surgical Robotics With C++ - Milad Khaledyan

<details>
<summary>
Safe C++ code for medical devices.
</summary>

[Building Safe And Reliable Surgical Robotics With C++](https://youtu.be/Lnr75tbeYyA?si=N0Ur-nPzV_KBpJyz), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Building_Safe_and_Reliable_Surgical_Robotics_using_Cpp.pdf), [event](https://cppcon2024.sched.com/event/1gZgI/building-safe-and-reliable-surgical-robotics-with-c).

the usual stuff about safety, memory safety, common CVEs, white house recommendation to move away from memory-unsafe langages, etc...

Medical device failure analysis: what causes them to fail (not considering security), showing that a software is the top cause of device failures. there are regulatory standards for software in medical devices, reports, guidelines, document and more. using SBOM (software bill of materials), supply chain transparency, CBOM (cybersecurity bill of materials), SOUP (software of unknown provenance) - documentacting all third part libraries (including the compiler!) and verifying them. but this isn't enough for medical devices. this is correct for every industry.

> Strive to Achieve Correctness, Safety and Security
>
> - Functional Correctness: Meeting specifications and requirements.
> - Functional Safety: Zero unintended behavior; Medical device operates correctly in response to inputs, including in failure scenarios (Fail-safe Design), to prevent harm or hazards to patient.
> - Security: Protection of systems, networks, and data from unauthorized access, attacks, damage, or theft.

#### Medical Use Case

Robotically Assisted Surgical Platform, if it fails, it could potentially hurt the patient.

> 1. Culture
> 2. Architecture
> 3. Processes
> 4. Tooling
> 5. Explore(?)

Risk-Driven Architecture, separating different subsystems according to safety concerns, does a failure in a particular subsystem risk the patient? what kind of segregation is needed (physical/logical), which subsystems are time sensitive and require real-time responsiveness?

compiler hardening - always using the correct flags. treating warnings as errors, preventing conversions between types and signs, add protections against deleting null pointers, protecting stack and buffer overflows with runtime flags. using sanitizers (address, threads, leaks, undefined behavior, realtime).

having a CI-CD pipeline, continuous testing, not introducing new bugs, not adding new vulnerabilities, guardrails against code debt, testing on multiple hardware, running regression and integration tests.

#### Writing Safe Code

Avoiding dynamic memory, first we need to identify it. it's not just calling <cpp>malloc</cpp> and <cpp>new</cpp>. anything that has an uknown size or unknown type (due to type erasure), even logging operations can allocate heap memory. but it's best to use tools and instrumentations. Not forcing abstraction when we don't need it.
Avoid Exceptions, using result types instead of throwing exceptions.  avoid blocking calls (file and network I/O), it's better to acquire resources before the safety critical path, but even that can be dangerous.

</details>
