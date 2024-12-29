<!--
// cSpell:ignore nsdm Unruh ftrivial
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Future C++

<summary>
7 Talks about future features.
</summary>

- [ ] Contracts for C++ - Timur Doumler
- [x] Peering forward — C++'s next decade - Herb Sutter
- [ ] Perspectives on Contracts - Lisa Lippincott
- [ ] Safety and Security Panel - Michael Wong, Andreas Weis, Gabriel Dos Reis, Herb Sutter, Lisa Lippincott, Timur Doumler
- [ ] Template-less Meta-programming - Kris Jusiak
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

In a simple defintion, reflection allows us to look at our code as data, this will require some new language support features.

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

we can combine reflection and safety to write simpler code. most features add complexity, but some good features actually make other code simpler, and reduce overall complexity. ranges and futures are examples of this, adding them to the standard reduced the page count. we take code patterns and elevate them to language defintions. the code doesn't only express the "how", but also the "what". we have intent built-in into the code.

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
