<!--
// cSpell:ignore Regehr
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Embedded

<summary>
9 Talks about embedded C++.
</summary>

- [ ] Balancing Efficiency and Flexibility: Cost of Abstractions in Embedded Systems - Marcell Juhasz
- [ ] Boosting Software Efficiency: A Case Study of 100% Performance Improvement in an Embedded C++ System - Gili Kamma
- [ ] C++ Exceptions for Smaller Firmware - Khalil Estell
- [ ] Leveraging C++20/23 Features for Low Level Interactions - Jeffrey Erickson
- [ ] Message Handling with Boolean Algebra - Ben Deane
- [ ] Multi Producer, Multi Consumer, Lock Free, Atomic Queue - User API and Implementation - Erez Strauss
- [ ] Sender Patterns to Wrangle Concurrency in Embedded Devices - Michael Caisse
- [ ] SuperCharge Your IPC Programs With C++20 and CCI Pattern - Arian Ajdari
- [x] What Volatile Means (and Doesn't Mean) - Ben Saks

---

### What Volatile Means (and Doesn't Mean) in C++ Programming - Ben Saks

<details>
<summary>
The Volatile Keyword, what it doesn, when and to not use it.
</summary>

[What Volatile Means (and Doesn't Mean) in C++ Programming](https://youtu.be/GeblxEQIPFM?si=V2vUx6gdiGcdEAGX), [slides](<https://github.com/CppCon/CppCon2024/blob/main/Presentations/What_Volatile_Means_(and_Doesn't_Mean).pdf>)

> The volatile qualifier is a vital tool for preventing compilers from performing certain harmful optimizations. Unfortunately, many C++ programmers aren't clear on exactly what protections volatile provides.\
> As such, many programmers apply the volatile qualifier incorrectly. A misapplied volatile might:
>
> - prevent optimizations unnecessarily, or worse
> - fail to provide the expected protection, leading to subtle run-time bugs

the classic use of the keyword is when using device register, which we use to communicate with external hardware. we usually map the addresses of the external device into addresses in the memory space we use. there are all kinds of bits that tell us the status of the device.

> For example, the E7T has two serial ports (UARTs), each with the same layout:
>
> | Offset    | Register | Description                 |
> | --------- | -------- | --------------------------- |
> | 0x00 (0)  | ULCON    | line control                |
> | 0x04 (4)  | UCON     | control                     |
> | 0x08 (8)  | USTAT    | status                      |
> | 0x0C (12) | UTXBUF   | transmit buffer             |
> | 0x10 (16) | URXBUF   | receive buffer              |
> | 0x14 (20) | UBRDIV   | baud rate divisor (control) |

here is an example of why we need the <cpp>volatile</cpp> keyword.

```cpp
std::uint32_t &USTAT0 =
*reinterpret_cast<special_register *>(0x03FFD008); // status
std::uint32_t &UTXBUF0 =
*reinterpret_cast<special_register *>(0x03FFD00C); // transmit
while ((USTAT0 & TBE) == 0) {} // busy wait

UTXBUF0 = c;
```

just based on the code, the compiler might optimize away checking the actual memory. but those addresses aren't contained just in the program, they are connected to hardware and have side effects which we need.\
it's possible to disable optimizations (either entirely or for a block of code), but we usually want those optimizations. so the <cpp>volatile</cpp> keywords indicates to the compiler that the value of the memory can change, even if the program doesn't change it, and that the compiler should not omit any reads or writes to it.

<cpp>volatile</cpp> is one of the "cv-qualifiers", together with the <cpp>const</cpp> keyword. it goes together with declaration specifiers (types, pointer, reference, array, function, static, extern, inline, typedef).\
With pointers, the volatile can be applied to the value (the pointer points to a volatile address) or the pointer itself (the address the pointer points to might change).

compilers also have an assumption that volatile variables are connected, and therefore it can't re-order them, and code related to them can be re-ordered only with respect to them. this is actually why it isn't a reliable tool for communication between threads.

```cpp
bool volatile buffer_ready;
char buffer[BUF_SIZE];

void buffer_init() {
  for (int i = 0; i < BUF_SIZE; ++i) {
  buffer[i] = 0;
  }
  buffer_ready = true; // looks like a reliable signal, but...
}
```

in this case, a compiler might decide to first set the volatile variable before filling the buffer, since there is no connection between the buffer object and the flag. if we set the buffer as volatile as well, then there won't be re-ordering, but it will have other costs. if we want multi-threading, there are dedicated tools for that:

- mutexes
- semaphores
- condition variables

volatile also doesn't mean atomic.  we aren't protected from a data race.

```cpp
double volatile v = 0.0;
void modify_value() {
  v = 8.67; // #1
  v = 53.09; // #2
}
```

if operations on double aren't atomic in our platform, then access to object from other threads might result in a value that isn't 8.67 or 53.09.

> Of particular note, increment, decrement, and compound assignment expressions such as these are not guaranteed to be atomic for volatile objects:

```cpp
double volatile v = 0.0;
v++;
--v;
v += 3;
v <<= 2;
```

Compilers aren't always great with volatile variables, back in 2008 (Eric, Eide and Regehr), some tested compilers produced code that accessed variables in ways they shouldn't have. this is especially bad, since hardware devices are often used in safety critical scenarios.\
This might be an on-going problem, since compilers want to optimize and this is directly against what volatile variables are supposed to protect from. we can turn off optimizations for the compiled code if we think the compiler isn't behaving correctly.

```cpp
void [[gnu::optimize("O0")]] void foo() f() { /*...*/ }

// or
#pragmas:
#pragma GCC push_options
#pragma GCC optimize ("O0")
void f() { /*...*/ }
#pragma GCC pop_options
```

there are other workarounds, such as marking functions to not be inlined.

> Takeaways
>
> - volatile tells the compiler that accessing an object may have side effects that mustn't be optimized away.
> - The compiler must keep accesses to volatile objects in order, but may reorder accesses to non-volatile objects around them.
> - Use synchronization tools (e.g., mutexes and semaphores) rather than volatile objects to manage inter-thread communication.
> - Accesses to volatile objects are not guaranteed to be atomic
> - If you find that your compiler is mishandling volatile, try these remedies:
>   - Disable optimizations for that code.
>   - Use a different version of the compiler.
>   - Use Eide and Regehr's workaround.
>
> - If you do use Eide and Regehr's workaround, make sure that the functions aren't inlined

actually, in C++20, many operations around volatile variables got deprecated (and some were un-deprecated later).

</details>
