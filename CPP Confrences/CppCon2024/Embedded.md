<!--
// cSpell:ignore Regehr dscp OSPEEDR MODER ULCON UTXBUF URXBUF UBRDIV strb
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Embedded

<summary>
9 Talks about embedded C++.
</summary>

- [x] Balancing Efficiency and Flexibility: Cost of Abstractions in Embedded Systems - Marcell Juhasz
- [ ] Boosting Software Efficiency: A Case Study of 100% Performance Improvement in an Embedded C++ System - Gili Kamma
- [ ] C++ Exceptions for Smaller Firmware - Khalil Estell
- [x] Leveraging C++20/23 Features for Low Level Interactions - Jeffrey Erickson
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

### Balancing Efficiency and Flexibility: Cost of Abstractions in Embedded Systems - Marcell Juhasz

<details>
<summary>
Effects of design decisions on performance of embedded code.
</summary>

[Balancing Efficiency and Flexibility: Cost of Abstractions in Embedded Systems](https://youtu.be/7gz98K_hCEM?si=P_f1maqyB9ejvoy4), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Balancing_Efficiency_and_Flexibility.pdf), [event](https://cppcon2024.sched.com/event/1gZeI/balancing-efficiency-and-flexibility-cost-of-abstractions-in-embedded-systems).

HAL - Hardware Abstraction Layer

> - Encapsulation
> - Inheritance
> - Polymorphism
> - Template Metaprogramming
> - Concepts
> - Constant Expressions
> - Immediate Functions
> - Parameter Pack and Fold Expressions
> - <cpp>constexpr If </cpp> Statements

measuring binary size, runtime performance (analytical and empirical)

> Absolute minimal embedded project:
>
> - Main function:
>   - Consists of a single, empty infinite loop
> - Startup script:
>   - Defines the vector table
>   - Sets stack pointer
>   - Copies data section from Flash to RAM
>   - Initializes uninitialized global and static variables to zero
>   - Calls static constructors
>   - Calls `main()`
> - Linker Script: specifies the memory

this is how it traditionally looks:

```cpp
typedef struct {
  uint32_t pin;
  GPIO_Modes mode;
} GPIO_InitStruct;

void GPIO_Init(GPIO_InitStruct* conf) {
  uint32_t temp;
  /* check the values */
  if (!IS_GPIO_PIN(conf->pin)) { return; }
  if (!IS_GPIO_MODE(conf->mode)) { return; }
  /* configure the GPIO based on the settings */
  if (conf->mode == GPIO_MODE_OUTPUT) {
    temp = OSPEEDR;
    temp &= ~(OSPEEDR_MASK << (conf->pin * 2u));
    temp |= (GPIO_SPEED_FREQ_LOW << (conf->pin * 2u));
    OSPEEDR = temp;
    /* ... */
  }
  /* ... */
}

int main (void) {
  GPIO_InitStruct conf = { 0 };
  conf.pin = GPIO_PIN_6;
  conf.mode = GPIO_MODE_INPUT;
  GPIO_Init(&conf) ;
  while (1) { }
}
```

#### Layers Of Abstractions

bitmask calculations for every register

##### Encapsulation

We start with encapsulation, the generic class of _CRegister_, and the conataining classes for specific register operations (_CModeRegister_, _COutputTypeRegister_, _COutPutSpeedRegister_).

```cpp
class CRegister {
private:
  const std::uint32_t m_address;
public:
  CRegister(std::uint32_t address) : m_address(address) { }
  void set(std::uint32_t val) const {
    *(reinterpret_cast<volatile std::uint32_t *>(m_address)) = val;
  }
  /* ... */
};

class CModeRegister {
private:
  const CRegister m_register{0x48000000};
  inline std::uint32_t calculate_value(std::uint32_t pin, GPIO_Modes mode) {
    return (mode & GPIO_MODE) << (pin * 2);
  }
  inline std::uint32_t calculate_bitmask(std::uint32_t pin) {
    return MODER_MASK << (pin * 2);
  }
public:
  inline void set_mode(std::uint32_t pin, GPIO_Modes mode) {
    m_register.set(calculate_value(pin, mode), calculate_bitmask(pin));
  }
}
```

Encapsulation also helps with testability, we can inject a different implementation of _CRegister_ to find out what's really happening.\
Encapsulation doesn't have effect on the binary, so we get the same binary size and same behavior, and we gain in readability and simplicity. we can also use some static members since the structs don't have state, but it causes an increase in binary size, since the static objects must be initialized somewhere. we can use templated classes with the address as the template parameter, which brings us back to the original binary size.

##### Inheritance

Next we look at inheritance, the generic class acts as base class, and the specific classes extend it.

```cpp
class COutputSpeedRegister : public CRegister<0x48000008> {
private:
  static inline std::uint32_t calculate_value(std::uint32_t pin, GPIO_Output_Speeds speed) {
    return speed << (pin * 2);
  }
  static inline std::uint32_t calculate_bitmask(std::uint32_t pin) {
    return OSPEEDR_MASK << (pin * 2);
  }
public:
  static inline void set_speed(std::uint32_t pin, GPIO_Output_Speeds speed) {
    set(calculate_value(pin, speed), calculate_bitmask(pin));
  }
};
```

because we don't have anything virtual here, we don't have any effect on runtime performance, and the binary is still the same.

##### Polymorphism

for polymorphism, we need our microcontrolher to communicate with a Peripheral, we could have the implementation coupled into the HAL, which would mean that we only work with a specific peripheral, but with polymorphism, we add an abstraction layer so we are decoupled. this is dynamic polymorphism, and now we use virtual inheritance.

```cpp
class IPin {
public:
  virtual void set() = 0;
  virtual void reset() = 0;
};

class CLed {
private:
  bool m_state{false};
  IPin* m_pin{nullptr}; // pimpl
public:
  CLed() = delete;
  CLed(IPin* pin) : m_pin(pin) {
    m_pin->reset();
  }
  /* ... */
};

class CPin : public IPin {
private:
  std::uint8_t m_pin{0};
public:
  CPin() = delete;
  CPin(std::uint8_t pin) : m_pin(pin) {}
  void set() override {
    CBitSetResetRegister::set_pin(m_pin);
  }
  void reset() override {
    CBitSetResetRegister::reset_pin(m_pin);
  }
};
```

if we include the RTTI (RunTime Type Information) in the binary, it increases the binary size substantially. it stores the <cpp>\_type_info</cpp> section into the binary, since it is necessary for <cpp>typeid</cpp> and <cpp>dynamic_cast</cpp>.\
if we manage to de-virtualize the calls, then we can decrease back the binary size to nearly the original size.

we can look at the assembly without de-virtualization - when we have a virtual table.

| CPin vtable             | address? | mangled name        |
| ----------------------- | -------- | ------------------- |
| 0x08000258 : 0x080001b1 | 080001b0 | <\_ZN4CPin3setEv>   |
| 0x0800025c : 0x08000199 | 08000198 | <\_ZN4CPin5resetEv> |

```x86asm
;no de-virtualization
<main>:
push {r0, r1, r2, lr}
ldr r3, [pc, #28] ; (80001e4 <main+0x20>) ; load virtual table
mov r2, sp
str r3, [sp, #0]
movs r3, #6
strb r3, [r2, #4]
ldr r3, [sp, #0]
mov r0, sp
ldr r3, [r3, #0]
blx r3
ldr r3, [sp, #0]
mov r0, sp
ldr r3, [r3, #4]
blx r3
b.n 80001d0 <main+0xc>
Nop ; (mov r8, r8)
.word 0x08000258
```

for static polymorphism (with de-virtualization), we can use <cpp>concepts</cpp> instead of virtual tables.

```cpp
template <typename T>
concept IPin = requires(T pin) {
  { pin.set() } -> std::same_as<void>;
  { pin.reset() } -> std::same_as<void>;
};

class CPin {
private:
  std::uint8_t m_pin{0};
public:
  CPin () = delete;
  CPin (std::uint8_t pin) : m_pin(pin) {}
  void set() const {
    CBitSetResetRegister::set_pin(m_pin);
  }
  void reset() const {
    CBitSetResetRegister::reset_pin(m_pin);
  }
};

static_assert(IPin<CPin>); // optional

template <IPin TPin>
class CLed {
private:
  bool m_state{false};
  TPin* m_pin{nullptr};
public:
  CLed() = delete;
  CLed(TPin* pin) : m_pin(pin) {
    m_pin->reset();
  }
  /* ... */
};
```

this compiles into the same binary as the original.

#### Architecture Matters

we usually take the architecture abstraction as a given, and write our c++ code as a wrapper over it.\
for our example, we have a seven segments display device, where each segment corresponds to a bit (an output data register). we have an operation to reset it - binary AND `&` with the correct bitmask (known at advance), and we set it with a binary OR `|` and some other bitmask.

```cpp
ODR &= ~0b1111_1111; // reset to empty
ODR |= 0b0101_1100; // set bits to the pattern of '5'
```

if our architecture only supports pin operations, we would have to iterate over each pin to set it, so we want our underlying hardware abstraction layer to operate on a collection of bits, and not just one-by-one.

example of inefficient code:

```cpp
typedef struct {
  uint32_t pin;
  GPIO_Modes mode;
} GPIO_InitStruct;

void GPIO_Init(GPIO_InitStruct* conf) {
  uint32_t temp;
  /* check the values */
  if (!IS_GPIO_PIN(conf->pin)) { return; }
  if (!IS_GPIO_MODE(conf->mode)) { return; } // enum
  /* configure the GPIO based on the settings */
  if (conf->mode == GPIO_MODE_OUTPUT) { // branching
    temp = OSPEEDR;
    temp &= ~(OSPEEDR_MASK << (conf->pin * 2u)); // runtime bitmask calculations
    temp |= (GPIO_SPEED_FREQ_LOW << (conf->pin * 2u));
    OSPEEDR = temp;
    /* ... */
  }
  /* ... */
}
```

> - Enums are basically integers with no value restrictions
> - Runtime branching based on input parameters
> - Runtime bitmask calculations
> - Function calls

we want to get rid of these inefficiencies, so we bring in some advance c++ to move the calculations into compile time.

```cpp
enum class modes : std::uint32_t {
  input = 0b00,
  output = 0b01,
  alt_func = 0b10,
  analog = 0b11
};
enum class ports : std::uint8_t {
  port_f,
  port_d,
  port_c,
  port_b,
  port_a
};

template <modes mode>
concept is_valid_mode = (
  (mode == modes::input) ||
  (mode == modes::output) ||
  (mode == modes::alt_func) ||
  (mode == modes::analog)
);

template <pins pin>
concept is_valid_pin = (
  is_valid_low_pin<pin> || is_valid_high_pin<pin>
);

template <pins... pin>
concept are_valid_pins = (is_valid_pin<pin> && ...);

template <modes mode, pins... pin>
requires (are_valid_pins<pin ...> && is_valid_mode<mode>)
consteval std::uint32_t moder_value () {
  return (... | (static_cast<std::uint32_t>(mode) << (static_cast<std::uint32_t>(pin) * 2)));
}

template <pins... pin>
requires (are_valid_pins<pin ...>)
consteval std::uint32_t moder_bitmask () {
  return (... | (GPIO_MODER_MODER0 << (static_cast<std::uint32_t>(pin) * 2)));
}

template <GpioInitConfig conf, pins... pin>
requires (is_valid_gpio_config<conf> && are_valid_pins<pin...>)
static inline void configure_pins() {
  static_assert((sizeof...(pin) > 0), "No pins provided.");
  if constexpr ((conf.mode == modes::output) || (conf.mode == modes::alt_func)) {
  /* ... */
  }
  if constexpr (conf.mode != modes::analog) {
  /* ... */
  }
  if constexpr (conf.mode == modes::alt_func) {
  /* ... */
  }
  /* ... */
}
```

with this re-write, we employ C++ code to get better performance.

> - Enumerations
>   - Not plain integers anymore
>   - No implicit conversion between enum values and integral types
>     - Explicit static_cast required
>     - Harder to misuse accidentally
> - Template parameters
>   - Compile time parameter passing Concepts
> - Concepts
>   - Set of requirements on template arguments
>   - Compile time parameter validation
> - Immediate Functions (<cpp>consteval</cpp>)
>   - Executed at compile time
>     - Useful for e.g., bitmask calculations
> - Compile-time Branching (<cpp>constexpr if</cpp>)
>   - Eliminates:
>     - Unused code
>     - Comparisons
>     - Jump instructions

after measuring, the improved code has a smaller binary size and executes less instructions overall. this is a win for us.

</details>

### Leveraging C++20/23 Features for Low Level Interactions - Jeffrey Erickson

<details>
<summary>
Using C++ for better HW safety.
</summary>

[Leveraging C++20/23 Features for Low Level Interactions](https://youtu.be/rfkSHxSoQVE?si=hTx0bTH2wp7G82U7), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Leveraging_Cpp20_23_Features_for_Low_Level_Interactions.pdf), [event](https://cppcon2024.sched.com/event/1gZfB/leveraging-c++2023-features-for-low-level-interactions).

> In a bare-metal environment, we're going to demonstrate effective use of C++:
>
> - How did we end up with C language HW interactions?
> - What are best practices for using C from C++?
> - How can we use C++ to make HW access cleaner, safer, and more testable?

C++ has some reputation for being bloated for bare-metal, this isn't always true. C++ has many advantages over C. we can't get rid of all the existing C code (there is a lot of it), but we can write C compatible code using the c++ best practices.

moving between C and C++ boundaries, the <cpp>std::unique_ptr</cpp> and the `.get()` method, (c++23 has <cpp>std::out_ptr</cpp> and <cpp>std::inout_ptr</cpp> as syntactic sugar helpers). the memory layout can be different between C and C++ objects, the compiler can re-order the layout, so there are some guidelines to follow when passing data. we can test the data with <cpp>std::is_trivial</cpp> to make sure it supports static initialization, and <cpp>std::is_standard_layout</cpp> to ensure the memory layout is the same in C and C++.

breaking free of the kernel, if we need C because the kernel needs it, maybe we an bypass the kernel, and therefore remove the need for C code.\
using the _pimpl_ idiom to wrap a pointer to the register inside a class - a pointer to a register struct.

```cpp
struct my_uart_regs {
// ...
  volatile uint32_t BAUD_RATE_REG;
// ...
};

class my_uart {
...
private:
  std::unique_ptr<my_uart_regs> p_regs;
};
```

we can pass a pointer to the real registry address, or use a mocking object. we need to define the register set. for hardware access, we need to use the <cpp>volatile</cpp> key word, since the hardware can change the data values, so we won't allow the compiler to take shortcuts.

```cpp
typedef struct __attribute__((packed)) __attribute__((aligned(4))){
  volatile uint32_t reg1; // Offset 0
  volatile uint32_t reg2; // Offset 4
  volatile uint32_t pad[4];
  volatile uint32_t reg3; // Offset 20
// ...
} regs_t;
```

this is one of the few cases where we need to use <cpp>reinterpret_cast</cpp>, we probably can't use <cpp>std::make_unique</cpp>, but anyway, we need a custom delete, since we would deallocate data that wasn't allocated, and doesn't have the metadata we expect. we need to defend from the delete, so we define a custom deleter ourselves.

```cpp

struct my_uart_regs __attribute__((packed)) __attribute__((align(4))){
// ...
  volatile uint32_t BAUD_RATE_REG;
// ...
};

struct uart_regs_no_deleter {
  operator()(my_uart_regs* ptr) {
  // do nothing
  }
};

class my_uart {
  //...
public:
  void set_baud_rate(const uint32_t& rate) {
    p_regs -> BAUD_RATE_REG = rate;
  }

  uint32_t get_baud_rate() const {
    return p_regs -> BAUD_RATE_REG;
  }
private:
  std::unique_ptr<my_uart_regs, my_uart_no_deleter>-> p_regs;
};
```

now we got to the registers themselves, and we need operate on them, we could use C binary logic operators directly, but C++ could give us better safety with strong typing, and use <cpp>constexpr</cpp> to move code from runtime to compile time.

```cpp
// c-style
uint32_t value = 0;
value |= 0x1;
value &= ~(0x1);
```

Macro are dangerous because they aren't type-checked, and the speed gains only come from inlining the code. we could define the masks as <cpp>constexpr</cpp> and it works great with floating numbers.

```cpp
// old macro
#define REPLACE_BITS(x, mask, bits) ((x & (~(mask))) | (bits & mask))
// c++ 

// templated function
template<std::size_t N>
std::bitset<N> replace_bits(std::bitset<N> x, std::bitset<N> mask, std::bitset<N> bits){
  return ((x & (~(mask))) | (bits & mask));
}
```

> Wrapping up
>
> - Developers do a lot in C to make it 'safer'
> - But C++ has advantages:
>   - Strong typing and a more thorough type system makes code safer
>   - Lifetime management is important and C++ takes many of those questions off the table
> - Still need static analysis, but the compiler does more for you

</details>
