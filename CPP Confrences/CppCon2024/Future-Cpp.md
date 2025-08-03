<!--
// cSpell:ignore nsdm Unruh ftrivial Berne Heisenbugs ification reflexpr unreflexpr
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Future C++

<summary>
7 Talks about future features.
</summary>

- [x] Contracts for C++ - Timur Doumler
- [x] Peering forward — C++'s next decade - Herb Sutter
- [x] Perspectives on Contracts - Lisa Lippincott
- [ ] Safety and Security Panel - Michael Wong, Andreas Weis, Gabriel Dos Reis, Herb Sutter, Lisa Lippincott, Timur Doumler
- [x] Template-less Meta-programming - Kris Jusiak
- [ ] This is C++ - Jon Kalb
- [x] Gazing Beyond Reflection for C++26 - Daveed Vandevoorde

---

### Gazing Beyond Reflection for C++26 - Daveed Vandevoorde

<details>
<summary>
What can we expect to get from reflection.
</summary>

[Gazing Beyond Reflection for C++26](https://youtu.be/wpjiowJW2ks?si=UbzpocPx26ok8Bfs)

Talking about future suggestions for reflection, even beyond the basic things planned for C++26.

#### Basic Reflection (P2996)

In a simple definition, reflection allows us to look at our code as data, this will require some new language support features.

The <cpp>^^</cpp> reflection operator, which we apply to a type and get the <cpp>std::meta::info</cpp> type (computation domain). in the new <cpp>meta</cpp> header.

```cpp
#include <meta>
auto r = ^^Vec; // r is of type std::meta::info
// ...is_template(r)...
r2 = substitute(r, {^^int}); // reflection for specialization
r3 = doTheThing(members_of(r2));
```

next we have ways to turn back the reflected data into code. we will use the <cpp>[:_reflected_:]</cpp> splicing operator for that.

```cpp
[:^^int:] i = 42; // i is int;
typename[:r:] y; // y is type of what we got from ^^r
template[:q:]<int> z; // z is template
[:n:]::X x = [:v:]; // x is X from namespace ^^n, and the value is ^^v;
```

another example, we can take this to compiler explorer.

```cpp
#include <experimental/meta>
#include <iostream>

struct Entry{
  int key:24; // bit fields
  int flags:8;
  double value;
} e = {100, 0x0d, 42.0};

int main()
{
  constexpr std::meta::info r = ^^Entry;
  std::cout << identifier_of(r) << '\n';
  constexpr std::meta::info dm = nonstatic_data_members_of(r); // vector of meta info
  std::cout << identifier_of(dm[2]) << '\n';
  std::cout << e.[:dm[2]:] << '\n'; // like calling e.value
}
```

we can also define new classes using reflection (non static data members only).

```cpp
struct I; // incomplete type

int main()
{
  constexpr info r = define_class(^^I, {
    data_member_spec(^^int, {.name="index"}),
    data_member_spec(^^bool, {.name="flag", .width=1})
  });
  I x = { .index = 42, .flag = true};
  static_assert(std::is_same_v<[:r:], I>);
}
```

and an example of creating a tuple with pack expansion

```cpp
template<typename... Ts> struct Tuple {
  struct storage;
  [:
    define_class(^^storage, {
      data_member_spec(^^Ts, {.no_unique_address=true})...}
      );
  :] data
  // constructors
  Tuple(): data{} {}
  Tuple(Ts const& ...ts): data{ts...} {}

  // getting element from non static data members
  static consteval std::meta::info nth_nsdm(std::size_t n)
  {
    return nonstatic_data_members_of(^^storage)[n];
  }

  template<std::size_t I, typename... Ts>
  constexpr auto get(Tuple<Ts...> &t) noexcept -> std::tuple_element_t<I, Tuple<Ts...>>&
  {
    return t.data.[:t.nth_nsdm(I):];
  }
  // ... other functions
};

// defining the tuple element type
template<std::size_t I, typename... Ts>
struct std::tuple_element<I, Tuple<Ts...>>
{
  using type = [: std::array{^^Ts...}[I] :];
}
```

#### Code Injection with Token Sequences (P3294)

Beyond the basic reflection, looking for inspiration from other languages, such as Swift Macros, to see how we can better express complex types. we replace semantic generation with code injection.

- String injection - e.g. D, CppFront
- Token injection - e.g. original templates, rust Macros
- Grammatical code injection - e.g. modern templates (with the <cpp>typename</cpp> in the template)

they decide to go with Token injection.

```cpp
#include <experimental/meta>
#include <iostream>

constval std::meta::info make_output_stmt()
{
  return ^^{std::cout << "Hello, World";};
}

int main()
{
  queue_injection(make_output_stmt());
}

template<bool B, typename T = void> struct enable_if {
  consteval { // consteval block, experimental feature that is syntactic sugar for static_assert
    if (B) queue_injection(^^{using type = T;});
  }
}
enable_if<true, int*>::type p = nullptr;
enable_if<false, int>::type i = 42; // ERROR
```

interpolation, using the `\` to mark interpolation, which can become code. the `\id` turns text into identifiers, and `\tokens` creates composition of tokens.

```cpp
consteval auto make_field(info type, string_view name, int val)
{
  return ^^{[:\(type):] \id(name) = \(val*2); };
}
consteval auto make_function(info type, string_view name, info body)
{
  return ^^{[:\(type):] \id(name) = \tokens(body) };
}

struct S {
  consteval {
    queue_injection(make_field(^^int, "x", 21));
    queue_injection(make_func(^^int, "f", ^^{ {retrun 42; }}));
  }
};

int main()
{
  return S{}.x != S{}.f();
}
```

Automatic Type Erasure - a facade pattern for interface. (incomplete code)

```cpp
consteval info param_tokens(vector<info> params, string_view prefix = "")
{
  std::meta::list_builder result(^^{,});
  for (int k = 0; info p : params) {
    p = type_of(p);
    if (prefix.size() != 0)
    {
      result += ^^{ typename[:\(p):] };
    }
    else // no prefix
    {
      result += ^^{ typename[:\(p):] \id(prefix, k++)};
    }
  }
  return result;
}

consteval void inject_Vtable(info interface)
{
  std::meta::list_builder vtable_members(^^{}); // empty list, no separator character
  for (info mem: members_of(interface))
  {
    if (is_function(mem) && !is_special_member(mem) && !is_static_member(mem)) // only member functions
    {
      info r = return_type_of(mem);
      auto name = identifier_of(mem);
      std::meta::list_builder parms(^^{,}); // list with the comma as a separator.
      params += ^^{ void* };
      params += param_tokens(parameters_of(mem));

      vtable_members += ^^{
        [:\(r):] (*\id(name))(\tokens(params));
      }; // add to the list
    }
  }

  queue_injection(^^{
    struct VTable{
      \tokens(vtable_members);
    } const *vtable;
  });
}

template<typename Interface> class Dyn
{
  void* data;
  consteval {
    inject_Vtable(^^Interface);
  };
  consteval {
    inject_vtable_for(^^Interface);
  };

  public:
  consteval {
    inject_interface(^^Interface);
  };
consteval {
    inject_erasing_ctor();
  };
  Dyn(Dyn&) = default;
  Dyn(Dyn const &) = default;
}
```

and using it to create a type erased thing.

```cpp
struct Interface {
  void draw(std::ostream&) const;
};

int main()
{
  struct Hello {
    void draw(std::ostream &os) const {
      os << "Hello\n";
    }
  } hello;
  struct Number {
    int i;
    void draw(std::ostream &os) const {
      os << "Number{"<< i << "}\n";
    }
  } one{1}, two{2};

  std::vector<Dyn<Interface>> v = {one, hello, two};
  for (auto &dyn : v) {
    dyn.draw(std::cout);
  }
}
```

#### Annotations for Reflection (P3394)

Also called "custom attributes".

for example, reflection based hash function of data members, using a custom attribute exclude some members as needed.

```cpp
enum class HashNotes {ignore};
template <typename T>
unsigned long hash (T const obj&)
{
  unsigned long result = 17;
  // magic to get thing into templated constants
  expand[:nonstatic_data_members_of(^^T):] >>
  [&]<info dm> {
      if (annotation_of_type<HashNotes>(dm) != HashNotes::ignore)
      {
        // .. do meta stuff with obj.[:dm:]
      }
  };

  return result;
}

struct Ultra {
  float data[3];
  Cache cache [[=HashNotes::ignore]];
}
```

another option could be to describe argument for a command line program.

we want something like this in the end.

```cpp
struct Args: clap::Clap {
  [[=Help("name to greet")]]
  [[=Short, =Long]]
  std::string name;

  [[=Help("number of times to repeat")]]
  [[=Long("repeat")]]
  int count;
};

int main(int argc, char** argv)
{
  Args args;
  args.parse(argc. argv);
  for (int i = 0; i< args.count; ++i)
  {
    std::cout << "Hello" << args.name << "!\n";
  }
}
```

(more code I won't copy), using C++23 explicit `this` for the parse function.

</details>

### Peering Forward - C++'s Next Decade - Herb Sutter - CppCon 2024

<details>
<summary>
Compile time expressions and reflection and how they can help us.
</summary>

[Peering Forward - C++'s Next Decade](https://youtu.be/FNi1-x4pojs?si=cVbQUucdV9T3-Aw5), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Peering_Forward_Cpps_Next_Decade.pdf)

will C++26 be the most influential release since C++11? there are four major things that are planned.

- <cpp>std::execution</cpp> (concurrency and parallelism) (voted in)
- Type and memory safety improvements(voted in)
- Reflection + code generation ('injection') (not voted in yet)
- Contracts (not voted in yet)

#### Reflection

> the program can **see** itself and **generate** itself... hence, **metaprograms**.\
> static -> zero runtime overhead (not runtime reflection)

this is part of the arc towards bringing more C++ code into compile time.

there is a famous 1994 code by _Erwin Unruh_ that doesn't compile. the errors it raises show the prime numbers.

```cpp
template <int i> struct D
{
  D(void*);
  operator int();
};

template <int p, int i>
struct is_prime
{
  enum { prim = (p%i) && is_prime<(i>2 ? p : 0), i-1>::prim };
};

template <int i>
struct Prime_print
{
  Prime_print<i-1> a;
  enum { prim = is_prime<i, i-1>::prim };
  void f() { D<i> d = prim; }
};

struct is_prime<0,0>
{
  enum {prim=1};
};

struct is_prime<0,1>
{
  enum {prim=1};
};

struct Prime_print<2>
{
  enum {prim=1};
  void f() { D<2> d = prim; }
};

main ()
{
  Prime_print<10> a;
}
```

C++11 introduced the <cpp>constexpr</cpp> keyword. and since then, more and more code become supported in compile time.

> - C++11 - single return statement
> - C++14 - more statements, local variable, conditions, member function
> - C++17 - lambdas, compile time destructors, <cpp>if constexpr</cpp>, <cpp>std::static_assert</cpp>
> - C++20 - memory allocation, virtual function, try catch, some data structures.
> - C++23 - more math, allocators, non-literal parameters

and also similar movement to have code run in the GPU. we started with Shaders, but as time evolved, we could do more things there.

and just like compile time code and running code on GPUs, reflection will start small and evolve over time.

an example of parsing command line parameters using reflection.

#### Safety

improving memory safety, both in terms of security (defending against malicious attackers) and safety (defending against unintended harm).

we want to achieve parity with other languages in terms of program safety, this comes down to four categories.

- type
- bounds
- initialization
- lifetime safety

progress is being made, and there will be an effort to enforce protection through profiles and move the protection from external tools to the compilers.\
we will move from "watch out" to "opt out" model. keeping all the abilities of the language, but making them only available by explicit choice. there is a suggested concept for **safety profiles** - a set of safety rules that are enforced in compile time and grantee that weaknesses are absent.

C++26 introduces "erroneous" behavior => "Well-defined as being just wrong". the first application of the tool is to categorize reading uninitialized local variables as erroneous behavior (rather than undefined behavior). A C++26 compiler will be required to overwrite the value with a known, defined value. no manual code-changes, just re-compiling and the code becomes safer.

example of leaking secrets by reading un-initialzed local variables. in most programs today, we will read the existing data, but moving forward, the program will read and return some pre-defined value, this will prevent some data leakages.

```cpp
auto f1() {
  char a[] = {'s', 'e', 'c', 'r', 'e', 't' };
}


auto f2() {
  char a[6]; // or std::array<char,6>
  print(a); // today this likely prints "secret"
}

int main() {
  f1();
  f2();
}
```

we can even experiment with it today by passing some flags to the compiler `-ftrivial-auto-var-init=<pattern>` (GCC, Clang) or `/RTR1` (MSVC). the value won't be zero, for reasons (see slide). we could use the <cpp>[[indeterminate]]</cpp> attribute to opt-out of using this feature.

code in compile time (<cpp>constexpr</cpp>, <cpp>consteval</cpp>) already rejects undefined behavior, so we can already write safe code. if we can do it in compile time, then we can do it everywhere

#### Simplicity

we can combine reflection and safety to write simpler code. most features add complexity, but some good features actually make other code simpler, and reduce overall complexity. ranges and futures are examples of this, adding them to the standard reduced the page count. we take code patterns and elevate them to language definations. the code doesn't only express the "how", but also the "what". we have intent built-in into the code.

an example of using reflection to define interfaces and the syntactic sugar.

```cpp
consteval void interface(std::meta::info proto)
{
  std::string_view name = identifier_of(proto);
  queue_injection(^^{
    class \id(name) {
    public:
      \tokens(make_interface_functions(proto))
      virtual ~\id(name)() = default;
      \id(name)() = default;
      \id(name)(\id(name) const&) = delete;
      void operator=(\id(name) const&) = delete;
    };
  });
}

consteval auto make_interface_functions(info proto) -> info
{
  info ret = ^^{};
  for (info mem : members_of(proto)) {
    if (is_nonspecial_member_function(mem)) {
      ret = ^^{
        \tokens(ret)
        virtual [:\(return_type_of(mem)):]
          \id(identifier_of(mem)) (\tokens(parameter_list_of(mem))) = 0;
      };
    }
    else if (is_variable(mem)) {
    // --- reporting compile time errors not yet implemented ---
    // print_error( "interfaces may not contain data members" );
    }
    // etc. for other kinds of interface constraint checks
  }
  return ret;
}

class(interface) Widget {
  int f();
  void f(std::string);
};
```

this will also allow for compile time regular expressions with performance comparable to the best regex engine. we might even get better compile performance and speed, despite the additional build steps.\
with reflection and generation, we could improve engines and language extensions like QT and microsofts' COM-IDL, reducing the amount of external tools and separate sets of knowledge.

small talk with _Andrei Alexandrescu_ discussing reflection to instrument types (wrapping around them with counts for each method) and using reflection to define domain specific languages like mathematical expressions, SQL statements and even C++ code.\
Things that are hard today will become easier, things which are impossible will become possible. for this future to come, all information in the source code must be reflect-able (e.g. the default accessability of structs and classes), and all code must generate-able, and all generated code must be visible (for debugging and visualization).

</details>

### Contracts for C++ - Timur Doumler

<details>
<summary>
The P2900 proposal for contracts, and what are contracts used for.
</summary>

[Contracts for C++](https://youtu.be/8niXcszTjis?si=JKyRi4sr3xd4uAOU), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Contracts_For_Cpp.pdf), [event](https://cppcon2024.sched.com/event/1gZfV/contracts-for-c++)

The C++26 contracts proposal, adding function contract assertions.

```cpp
Widget getWidget (index i)
  pre (i > 0) // precondition
  post (w: w.index() == i); // postcondition
  {
    auto *db = getDatabase();
    contract_assert (db != nullptr); // assortment statements
    return db->retrieveWidget;
  }
```

this is part of the SG21 contracts working group in the c++ standard.

the story of contracts starts in 1986 with the book "Design by Contract" and the Eiffel programming languages, in 2001 D came out with native contract support, and in 2012 Ada gained contracts.\
C++ first contract proposal was in 2004 ("D-Like Contracts"), then another proposal in 2013 ("BDE Contracts") which also failed, a merged proposal in 2015("attribute-like") and then in 2019 the SG21 workgroup was formed, and in 2023 they suggested the "Contracts MVP" approach. which is currently suggested to C++26.

#### What Are Contracts?

> Design By Contracts: An approach for software design:
>
> - Define formal, precise and verifiable interface specifications for software components, extending their ordinary definition with:
>   - precondition - Condition on passed-in arguments and/or program state when a function is called. Obligation of caller (client)
>   - postcondition - Condition on return value and/or program state when a function returns. Obligation of callee (implementer)
> - invariants (class invariants, loop invariants...)
> - called "Contracts" in accordance with a conceptual metaphor with the conditions and obligations of business contract

this talk is only about pre and post condition, today we have contracts as plain-language description (in the comments), but they aren't enforced. for example, the pre-condition for the square brackets operator is to use indexes larger than zero, and if the contract isn't followed there will be undefined behavior.

#### What Are Contracts For?

different usecases for contracts, as reported by the community

> - static analysis
> - verification
> - formal proofs
> - optimisation
> - safety
> - diagnose bugs
> - security
> - correctness
> - tooling support
> - annotations
> - debugging
> - expressivity

a single proposal can't do all of those things, but the P2900 proposal focuses on a subset of the needs.

> - Timur Doumler - P2900 Contracts enhance a C++ program with configurable checks of its correctness, thereby helping to diagnose and fix bugs, across API boundaries.
> - Lisa Lippincott - P2900 Contracts allow the programmer to express expectations about program state, and optionally verify that those expectations are met.
> – Joshua Berne - P2900 Contracts allow the programmer to specify states that are considered incorrect at certain points in a C++ program, particularly when calling and returning from functions, and then manage how such defects can be detected and mitigated, in a portable and scalable fashion, during program evaluation.

```cpp
// without contracts

class Widget<T> {
public:
  // Returns a reference to the element at position `index`.
  // The behaviour is undefined unless `index < size()`.
  T& operator[] (size_t index) {
    return _data[index]; // potentially UB here :(
  }

  T& at (size_t index) {
    if (index >= size())
      throw std::logic_error("Index out of bounds!");
    return _data[index];
  }
  
  T& safe_get_element (size_t index) {
    if (index >= size())
      std::terminate();
    return _data[index];
  }

  std::size_t size() const;

private:
  T* _data;
};

// with contracts
class WidgetX<T> {
public:
  // Returns a reference to the element at position `index`.
  // The behaviour is undefined unless `index < size()`.
  T& operator[] (size_t index)
    pre (index < size())
  {
    return _data[index]; 
  }

  std::size_t size() const;

private:
  T* _data;

};
```

with the contract, we express through code what are the conditions we operate on. then we divide the responsibilities between different roles.

> - Caller / client
>   - reads contract assertions
>   - ensures preconditions are met
>   - expects postconditions to be met
> - Callee / implementer
>   - writes contract assertions
>   - expects preconditions to be met
>   - ensures postconditions are met
> - Build engineer / owner of main()
>   - enables & configures contract checks
>   - decides what happens in case of contract violation

since the contracts are part of the code, we can operate on them programmatically, however, they aren't part of the 'regular' code flow, so we can change the behavior regarding them without changing the code itself.\
this is similar to <cpp>assert</cpp>, which we already have and we can enable/disable when compiling, and we can define the behavior when it fails. but the classic assertion isn't flexible enough.\
P2900 contracts are a superior replacement for assertions and macros. they are proper C++ code without the weird behavior of parsing macros, without the ODR that can happen when linking and without the possible compilation pitfalls. Macros can't go on declarations, only in the definations. contracts belong to the function declaration (but the can be repeated in the definition, if they are exactly the same as those in the declaration). post conditions can refer to the return value without messing up elisions and RVO.

contracts are evaluated at different points.

> - Precondition assertions (<cpp>pre</cpp>):
>   - after the initialisation of function parameters,
>   - before the evaluation of the function body
> - Postcondition assertions (<cpp>post</cpp>):
>   - after the result object value has been initialised and local automatic variables have been destroyed
>   - but prior to the destruction of function parameters
> - Assertion statements (<cpp>contract_assert</cpp>):
>   - when the statement is executed

contracts (pre and post) can refer to private variable, the name lookup rules behave as-if the statement is in the function body scope. when declareing a post condition, the variable it refer to must be const.

```cpp
int clamp(int v, int min, int max)
  pre (min <= max);
  post (r: val < min ? r == min : r == val)
  post (r: val > max ? r == max : r == val)
{
  min = max = value = 0;
  return 0; // won't compile
}

int clamp(const int val, const int min, const int max)
  pre (min <= max)
  post (r: val < min ? r == min : r == val) // parameters must be const
  post (r: val > max ? r == max : r == val); // on all declarations!
  {
    // code goes here
  }
```

With the classic assert statements, we always terminate the program. with macros, the behavior changes for each component. When a contract assertion fails, a contract-violation handler is called, the default behavior is to log a message to stdout and terminate, but we can customize it to our needs. just like the <cpp>operator new()</cpp>, we can change the behavior at link time (also called replaceable function or 'weak symbol').

```cpp
void ::handle_contract_violation(const std::contracts::contract_violation& violation)
{
  LOG(std::format("Contract violated at: {}\n", violation.location()));
}
```

we can log the error, throw an exception, trigger a breakpoint if a debugger is present, or optionally invoke the default behavior with <cpp>std::contracts::invoke_default_contract_violation_handler(violation);
</cpp>.

> Lakos Rule - Do not put <cpp>noexcept</cpp> on a function with preconditions, even if it never throws when called correctly!\
> C++ Standard follows the Lakos Rule, e.g. <cpp>std::vector::operator[]</cpp>

the standard contract library API is defined in the <cpp>std::contracts</cpp> namespace, under the <cpp>contracts</cpp> header. this header is only used to write the user-defined contract violation handler, not for writing contracts.

```cpp
namespace std::contracts {
  class contract_violation {
  // No user-accessible constructor, not copyable/movable/assignable
  public:
    std::source_location location() const noexcept;
    const char* comment() const noexcept;
    detection_mode detection_mode() const noexcept;
    evaluation_semantic semantic() const noexcept;
    assertion_kind kind() const noexcept;
  };
  void invoke_default_contract_violation_handler(const contract_violation&);

  enum class detection_mode : int {
    predicate_false = 1,
    evaluation_exception = 2,
    // implementation-defined additional values allowed, must be >= 1000
  };

  enum class assertion_kind : int {
    pre = 1,
    post = 2,
    assert = 3,
    // implementation-defined additional values allowed, must be >= 1000
  };
}
```

#### Evaluation Semantics

the evaluation policy controls how we interact with the predicates, if we are sure the program is correct, we can ignore the checks and conserve cycles, and we can choose how to behave if the predicate is not fulfilled.

> A contract assertion can be evaluated with one of the following evaluation semantics:
>
> - `ignore`: do not check the predicate (but still parse it)
> - `observe`: check the predicate, if the check fails call the contract-violation handler, when handler returns continue
> - `enforce`: check the predicate, if the check fails call the contract-violation handler, when handler returns terminate the program
> - `quick_enforce`: check the predicate, if the check fails immediately
terminate the program
> - `assume`: do not check the predicate and optimize on the assumption that it is true (= if it is false, the behaviour is undefined) - not in P2900

the proposal doesn't specify when the choice is made, it can be made at compile time with a compiler flag, and load time, at link time or even chosen at runtime (e.g. if a debugger is present). we can also choose how much of the contracts to check, we can choose to evaluate only a subset of them (pre or post), choose them at random, at regular intervals, and other stuff.\
This means that we can't rely on side effects in the contract, which was something Macro occasionally did.

```cpp
#include <cassert>
#ifndef NDEBUG
  unsigned nIter = 0;
#endif
while (keepIterating()) {
  assert(++nIter < maxIter);
  // ...
}

// can't (yet) have code conditional on
// whether contract checks are enabled
unsigned nIter = 0;
while (keepIterating()) {
  contract_assert(++nIter < maxIter); // bad, can't change state!
  // ... ^^^^^^^ may break!
}
```

> The Contracts Prime Directive:\
>The presence or evaluation of a contract assertion in a program should not alter the correctness of that program
>
> - Statically enforced:
>   - Adding a contract assertion can't affect Concepts / overload resolution / noexcept operator / SFINAE / if constexpr ...
>   - Adding a contract assertion that is not checked can't cause runtime overhead
> - Responsibility of the user:
>   - Don't use predicates with side effects that can alter the correctness of the program, the result of another contract assertion, or a subsequent check of the same contract assertion
> - Benefit: You can't get "Heisenbugs" - (e.g. bugs appearing/disappearing when you enable/disable a contract check)

#### New In P2900 Revision 8: Virtual Function Support!

a recent addison to the proposal, allowing predicates on virtual functions, separating callee and caller facing contracts. it combines static checks on the caller side depending on the interface used.

```cpp
struct UnaryFunction {
  virtual Value compute(ArgList args)
  pre (args.size() == 1);
};

struct BinaryFunction {
  virtual Value compute(ArgList args)
  pre (args.size() == 2);
};

struct VariadicFunction: UnaryFunction, BinaryFunction {
  Value compute(ArgList args) override
  /* no preconditions */;
};

int main() {
  VariadicFunction varFunc;
  test(varFunc);
}

void test(VariadicFunction& varFunc) {
  varFunc.compute({1}); // OK
  varFunc.compute({2, 3}); // OK
  varFunc.compute({4, 5, 6}); // OK
}

void test(UnaryFunction& unaryFunc) {
  unaryFunc.compute({1}); // OK
  unaryFunc.compute({2, 3}); // violation
  unaryFunc.compute({4, 5, 6}); // violation
}

void test(BinaryFunction& binFunc) {
  binFunc.compute({1}); // violation
  binFunc.compute({2, 3}); // OK
  binFunc.compute({4, 5, 6}); // violation
}
```

#### Open Question For P2900

> - make pre / post work on coroutines?
> - make pre / post work on function pointers
> - Keep `const`-ification?
> - undefined behavior in contract predicates

for function pointers, it's not clear where the information lives, it it's in the code then it interacts with templates, overloading and name mangling. if it's in the value, then it has runtimes overheads and wouldn't work with <cpp>typedef</cpp>, and also AST won't work for non-trivial stuff and across translation units.\
for contracts, parameters and local variables used in them are implicitly `const`, so there is a question what do with contracts when the stuff isn't marked as such. there is a problem of what to do with undefined behavior in contracts, which can still happen (overflow).\
There are suggestions about how to extend the contracts in the future, such as comparing 'old values' from previous calls.

> P2900 Contracts vs. safety & security
>
> - Contract assertions can significantly improve correctness & safety of code,
>   - but you have to actually add them to your code!
> - Contract assertions can detect if a contract was violated,
>   - but not prove that no contract was violated! (only a subset of the plain-language contract can be expressed in code)
> - Contract assertions check for correctness during evaluation of the program
>   - includes constant evaluation
> - Contract assertions do not change the semantics of the language
>   - No new language constraints
>   - No new language guarantees (e.g. guaranteed memory safety)

</details>

### Template-less Meta-programming - Kris Jusiak

<details>
<summary>
Value based template metaprogramming.
</summary>

[Template-less Meta-programming](https://youtu.be/yriNqhv-oM0?si=IdLDgNmfdB2gZsNl), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Template-Less_Meta-Programming.pdf), [event](https://cppcon2024.sched.com/event/1gZeh/template-less-meta-programming).

Metaprogramming is generating code by code, templates are one way of doing metaprogramming, which was chosen by C++. this talk will be about meta-programming that isn't templates-based.\

```cpp
template<class... Ts>
template<class T>
constexpr variant<Ts...>::variant(T&& t)
  : index{find_index<T, Ts...>} // Powered by TMP
  , // ...
{ }

template<size_t I, class... Ts>
constexpr auto get(tuple<Types...>&&) noexcept ->
  typename tuple_element<I, tuple<Ts...>>::type&&; // Powered by TMP

template<class TFirst, class... TRest>
array(TFirst, TRest...) -> array<
  typename Enforce_same<TFirst, TRest...>::type, // Powered by TMP
  1 + sizeof...(TRest)
  >;
```

The standard template library (STL) doesn't have a metaprogramming library. there was a proposal by Peter Dimov for a library.\ we could use MP to pack structs into a tighter format, and then get better performance, we could also use it for domain specific languages and state-machine.

> A brief history of template metaprogramming
>
> - C++
>   - Type-based TMP (<cpp>boost.mpl</cpp> -> <cpp>boost.mp11</cpp>)
>   - Heterogeneous-based TMP (<cpp>boost.fusion</cpp> -> <cpp>boost.hana</cpp>)
>   - Value-based TMP (<cpp>mp</cpp>, P2996 proposal)
>   - Circle-lang (meta model)
> - Zig-lang (comptime)

an example of value-based TMP - finding the index of a specific type in a list of types,

```cpp
template<class T, class... Ts>
constexpr auto find_index() -> std::size_t; // TODO
static_assert(0u == find_index<int, int, float, short>()); 
static_assert(1u == find_index<int, float, int, short>());
static_assert(2u == find_index<int, float, short, int>());
static_assert(3u == find_index<void, float, short, int>());
```

we will want to somehow iterate over the types, check each of them and return the index, we need some magic to happen here using typeIds. a typeId must be unique for the type.

```cpp
template<class T, class... Ts>
constexpr auto find_index() -> std::size_t {
  std::array ts{meta<Ts>...};
  for (auto i = 0u; i < ts.size(); ++i) {
    if (ts[i] == meta<T>) {
      return i;
    }
  }
  return ts.size();
}
```

we can refactor this into something more modern, use ranges and <cpp>constexpr</cpp>, or use stl algorithms internally, but does removing raw loops help making compilation faster? or does it make it slower?

we can do another example for returning a type, rather than the index. until c++26 the standard did it with recursive templates.

more examples, one example of getting types that can be used for constructing a different type.

```cpp
struct bar { };
struct foo {
  foo(int) { }
  foo(bar) { }
};
static_assert(std::is_same_v<
  std::variant<int, bar>, // because of foo(int) and foo(bar)
  variant_for_t<foo, const int&, int&&, std::string_view, bar, void>>
);

template<class T>
constexpr auto variant_for(const std::ranges::range auto& ts) -> std::ranges::range auto {
  std::vector<info> r;
  for (auto t : ts) {
    if (invoke<std::is_constructible, T>(t)) {
      r.push_back(invoke<std::remove_cvref>(t));
    }
  }
  std::ranges::sort(r);
  r.erase(std::ranges::unique(r), r.end());
  return r;
}

static_assert(std::vector{meta<int>} == variant_for<foo>(meta<int&>));
```

or with ranges

```cpp
template<class T>
constexpr auto variant_for(const std::ranges::range auto& ts) -> std::ranges::range auto {
  auto&& r = ts
    | std::views::filter(is_constructible<T>)
    | std::views::transform(remove_cvref)
    | std::ranges::to<std::vector>();
  std::ranges::sort(r);
  r.erase(std::ranges::unique(r), r.end());
  return r;
}

template<class T> constexpr auto is_constructible = [](auto t) {
  return invoke<std::is_constructible, T>(t);
}
```

#### Adding Reflection To The Mix

the same behavior of value based metaprogramming works with the upcoming reflection feature. we get the meta type for free instead of creating stuff ourselves.

```cpp
//^^T // reflection operator (reflexpr)
static_assert(^^int == ^^int);
static_assert(^^int != ^^void);
static_assert(typeid(^^int) == typeid(^^void));

//[: ... :] // splicer operator (unreflexpr)
typename [: ^^int :] i = 42; // int i = 42;
static_assert(typeid([: ^^int :]) == typeid(int));
```

back to a previous example

```cpp
constexpr auto find_index(auto t, const std::ranges::range auto& ts) -> std::size_t {
  if (const auto found = std::ranges::find(ts, t); found) {
    return std::distance(v.begin(), found);
  }
  return ts.size();
}
static_assert(
  0u == find_index(^^int, std::array{^^int, ^^float, ^^short}) and
  1u == find_index(^^float, std::array{^^int, ^^float, ^^short}) and
  2u == find_index(^^short, std::array{^^int, ^^float, ^^short}) and
  3u == find_index(^^void, std::array{^^int, ^^float, ^^short})
);
```

#### Benchmarking

[comparison between the different approaches](https://qlibs.github.io/mp/)

> - Circle-lang meta model is the fastest to compile all around
> - Type-based Metaprogramming with template aliases/builtins (<cpp>boost.mp11</cpp>) is much faster to compile than recursive template instantiations (<cpp>std::tuple</cpp>)
> - Value-based Metaprogramming with is significantly slower to compile than STL with raw primitives!
> - Value-based Metaprogramming has a lot of potential (<cpp>std::simd</cpp>, <cpp>std::execution</cpp>) but <cpp>constexpr</cpp> evaluation has to be JITTED instead of INTERPRETED

</details>

### Perspectives on Contracts - Lisa Lippincott

<details>
<summary>
Getting consensus on contracts.
</summary>

[Perspectives on Contracts](https://youtu.be/yhhSW-FSWkE?si=WBw6eDEQ9MA5779J), [event](https://cppcon2024.sched.com/event/1h1r4/perspectives-on-contracts).

Error detection requires redundancy, assertions are redundant by nature. we can choose between error detection and performance with the <cpp>NDEBUG</cpp> flag. In C, we could have made this choice separately for each translation unit, but in C++, since we have template instatantiaction and code the can be seen across them and ODR violations.\
contracts are a way to solve this problem in C++, we have a standard definition of what assertion statements do, no matter what compiler settings are used. this definition is very vague.

> Different points of view:
>
> - contracts make no difference to the program - a successful assertion should **do** nothing significant.
> - contracts make a significant difference to the program - a assertion should **say** something significant.

Pre and Post conditions exist between the caller and callee. this is the "contract", the points of agreements, this is how we decide if the bug is at the client (caller) or the library (callee). logically, they happen between the two translation units, but since they are code, it can't be possible. so actually, the conditions are more complicated than that.

condition-type     | translation unit | direction
-------------------|------------------|----------
caller-facing pre  | caller           | backward
callee-facing pre  | callee           | forward
callee-facing post | callee           | backward
caller-facing post | caller           | forward

There is a difference in view point. backwards facing is the detective, we suspect a bug exists, we want to track it to the root and verify the code we wrote is meeting expectations. in contrast, the forward facing point of view is that of the engineer (safety first), it protects from bugs before they happen, and limit the effect of the bugs to avoid disasters.

An assertion might be executed more than once, and because of that, the assertion statement (contract) shouldn't do anything significant. there is also the issue of virtual functions and function pointers. the caller only sees the baseclass pre and post conditions, while the callee sees the over-riding conditions. those might not be the same. if the contracts don't match up, there is a problem of substituion, a gap in the inheritance, not following Liskov SOLID principles. (something about V-tables layout).\
There are situations where we make de-virtual calls, in those cases the compiler knows which derived class is being called, and it could check the overrides pre and post condition even from the caller side. this would lead us to executing the assertions an additional time. (something about old code with suspect assertions).

Can we test an assertion doesn't do anything "significant"?  for now this is mostly a metric that exists in our mental model. we could do some calculations on a known "significant" thing, run the 'suspect' assertions and then check the known significant thing again. in practice, this means we can re-check assertions again and again, as long as nothing other than assertions happened between them.
</details>
