<!--
// cSpell:ignore
-->

# Core C++ 2023

<!-- <details> -->
<summary>
//TODO: add Summary
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

| Metric           | TI 99/4a  | 386 sx      | Pentium       | Z600
| ---------------- | --------- | ----------- | -------|------ |
| Ram              | 16 Kb     | 4 Mb        | 16 Mb         | 24 Gb
| Registers        | 16 bit    | 32 bit      | 32 bit        | 64 bit
| CPU              | 3 Mhz     | -           | 133Mgz        | 2.64Ghz
| Memory           | -         | 40Mb        | 500Mb         | 1Tb
| Speed Over human | 3,829,787 | 101,333,333 | 2,537,000,000 | 72,090,000,000 (without multi-threading)


> **semicolon** - A mark (`;`) of punctuation, indicating a greater degree of separation than the comma.

**1947** - Assembly language, by *Kathleen Booth* and her husband *Andrew Booth*, created for the A.R.C machine, in preparation that the same instructions could carry over to more modern machines as they become available. **1951** - abstraction from a machine that creates instructions to a language that creates a the instructions. later we got *Grace Hopper* and the A-0 system (arithmetic Language version 0), which laid the foundation for the first compiler. next we meet *John Backus* and **Fortran**, which introduced the optimizing compiler, and since then we no longer directly translate source code into machine code, we have something that changes it. *Frances E.Allen* introduces graph theory in the sixties.\
*Dennis Richie* and *Ken Thompson* creating the C language to work on the Unix operating system, later on *Bjarne Stroustrup* which borrowed from **Simula** and BCPL and created **C++**.

When Pentium 4 were released, the "NetBurst" architecture, instructions execute at the same time, using something called "shadow registers". so it's not only that the compiler modifies the source code into something else, the CPU also modifies the assembly code.
</details>

</details>
