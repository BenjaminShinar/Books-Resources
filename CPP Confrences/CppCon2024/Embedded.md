<!--
// cSpell:ignore Regehr dscp
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
- [x] Message Handling with Boolean Algebra - Ben Deane
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

volatile also doesn't mean atomic. we aren't protected from a data race.

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
>
>   - Disable optimizations for that code.
>   - Use a different version of the compiler.
>   - Use Eide and Regehr's workaround.
>
> - If you do use Eide and Regehr's workaround, make sure that the functions aren't inlined

actually, in C++20, many operations around volatile variables got deprecated (and some were un-deprecated later).

</details>

### Message Handling with Boolean Algebra - Ben Deane

<details>
<summary>
boolean algebra and how to compose expressions.
</summary>

[Message Handling with Boolean Algebra](https://youtu.be/-q9Omkhl62I?si=xyYrPRUi02glqTDA), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Message_Handling_with_Boolean_Algebra.pdf), [additional code](https://github.com/intel/compile-time-init-build).

> "The Unreasonable Effectiveness of Boolean Algebra in Software Design,
> Showing the Particular Application of a Message Handling Library,
> with an Excursion into the Roots of Programming"
>
> The workings of a message-handling library:
>
> - How messages and the fields in them are specified.
> - Efficiently identifying (with matchers) a message coming off the wire.
> - The role of Boolean algebra in composing matchers.
> - Understanding Boolean implication and using it to simplify matchers.

starting with a single request:

> Factor arbitrary matcher expressions into "sum-of-products" expressions

#### Fields and messages: how they are structured and specified

a field has a name (human readable), a type, and one or more locators that define the layout:

- word index (in 32-bit words)
- most significant bit (inclusive)
- least significant bit

using the IPv4 header as example, we can either extract from a field on insert, (get and set), the field is both the spec (name and type) and the locator, the same spec can be located at different messages in different locations.\
we can either indicate the location as offset from the current word index, or use absolute terms (offset from start of the message). fields can be split across locations, like having an eight bit field spread across two 4-bit locations.

```cpp
using version = field<"version", std::uint8_t>::located<at{0_dw, 3_msb, 0_lsb}>;
using ihl = field<"ihl", std::uint8_t>::located<at{0_dw, 7_msb, 4_lsb}>;  // internet header length
using dscp = field<"dscp", std::uint8_t>::located<at{0_dw, 13_msb, 8_lsb}>;  // differentiated services code point
using ecn = field<"ecn", std::uint8_t>::located<at{0_dw, 15_msb, 14_lsb}>; // explicit congestion notification

template <range R>
constexpr static auto extract(R &&r) -> value_type;

template <range R>
constexpr static auto insert(R &&r, value_type const &value) -> void;

using ipv4_header_layout = message<"ipv4", version, ihl, dscp, ecn, /*... more stuff*/>;
owning<ipv4_header_layout> ipv4_header; // layout + right-sized array
auto version = ipv4_header.get("version"_field);
ipv4_header.set("ihl"_field = 5);
```

#### Matching on fields in a message

> To determine the correct message type for data arrived off the wire, we use a matcher.\
> A matcher is a predicate on a message.

```cpp
template <typename T>
concept matcher = requires(T const &t, message const &msg)
{ // a prototypical message
  { t(msg) } -> std::convertible_to<bool>;
};

// simple matcher
template <typename Field, auto Value> struct equal_to_t
{
  constexpr auto operator()(auto const &msg) const -> bool
  {
    return Value == msg.template.get<Field>();
  }
};

// generic
template <typename Op, typename Field, auto Value>
struct relational_matcher_t
{
  constexpr auto operator()(auto const &msg) const -> bool
  {
    return Op{}(Value, msg.template.get<Field>());
  }
};

// convenient aliases
template <typename Field, auto Value>
using equal_to_t = relational_matcher_t<std::equal_to<>, Field, Value>;

template <typename Field, auto Value>
using less_than_t = relational_matcher_t<std::less<>, Field, Value>;
```

since matchers return a boolean, we can treat them like boolean values themselves, and start composing them together, so we can create matchers out matchers (and, or, not).

and now we can start using this library. we define fields, messages, callbacks, etc...\
we have an expression tree, but we want to make sure the library doesn't perform extra work. it should be able to simplify expressions to reduce evaluations.

the following composite matchers could be reduced

> - less_than<5> or greater_equal<5>
> - less_than<5> and less_than<7>
> - lot less_than<5> and greater_equal<5>

#### Simplifying matchers

there are some mathematical rules, but how can we let the library know about them?

we start by defining matchers for true and false: always and never. we can do some compile time stuff to replace matchers and collapse part of the tree. we define a compile time simplification function and make it customization point. we define matchers for conjunction ("and"), disjunction ("or") and negation ("not"), and define some identities and teach our library to simplify them.\

#### Programming as boolean algebra

when we define a function prototype (with a return type), we actually define the boolean implication (A=>B). this is called **Curry-Howard Isomorphism**. something about proposition.

we can take this knowledge and bring it into our library. implication is an operator, so it has a truth table, it's actually the same as "not A or B" (`!A or B`)

|  A  |  B  | A and B | A or B | A implies B |
| :-: | :-: | :-----: | :----: | :---------: |
|  0  |  0  |    0    |   0    |      1      |
|  0  |  1  |    0    |   1    |      1      |
|  1  |  0  |    0    |   1    |      0      |
|  1  |  1  |    1    |   1    |      1      |

this table hides simplifications, A=>B is always true when B is true, and A=>B is the value of B when A is false.
knowing the state of implication also tells us about the other columns. if A=>B is true, then "A and B" is the value of A, and "A or B" is the value of B.

#### Using implication to simplify

we have the expression tree of "less_than<5> and less_than<7>", which we as humans know is $X < 5 ⇒ X < 7$ (if x is smaller than 5, it's also smaller than 7), if "A=>B" is true, then "A and B" is A. so we need to bring this into the code.

```cpp
template <matcher X, matcher Y>
constexpr auto implies(X const &, Y const &) -> bool 
{
  return std::same_as<X, Y>;
}

constexpr auto implies(matcher auto &&, always_t) -> bool 
{
  return true;
}

constexpr auto implies(never_t, matcher auto &&) -> bool 
{
  return true;
}

constexpr auto implies(never_t, always_t) -> bool 
{ // ambiguity breaker!
  return true;
}

// some overloads
template <matcher M, matcher L, matcher R>
constexpr auto implies(and_t<L, R> const &a, M const &m) -> bool
{
  return implies(a.lhs, m) or implies(a.rhs, m);
}

template <matcher M, matcher L, matcher R>
constexpr auto implies(M const &m, or_t<L, R> const &o) -> bool
{
  return implies(m, o.lhs) or implies(m, o.rhs);
}

// generic matchs to simplify the same operations
template <typename Op, typename Field, auto X, auto Y>
constexpr auto implies(relational_matcher_t<Op, Field, X>, relational_matcher_t<Op, Field, Y>) 
{
  return X == Y or Op{}(X, Y);
}

// explicitly stating implication,
template <typename Field, auto X, auto Y>
constexpr auto implies(less_equal<Field, X> const &,
less_than<Field, Y> const &) -> bool 
{
  return X < Y;
}
```

we take this "implies" function and bring it into the simplification functions.

#### Epilogue

disjunctive normal forms: an expression with only "or", "and" and single terms. negation only applies to single terms, and the conjunctions and disjunctions don't contain additional expressions inside them.\
we can transform any expression into a disjunctive normal form using DeMorgans' laws and distributive laws.

going back to the "sum_of_products" topic.
</details>
