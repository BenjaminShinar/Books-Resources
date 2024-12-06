<!--
// cSpell:ignore nsdm
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Future C++

<summary>
7 Talks about future features.
</summary>

- [ ] Contracts for C++ - Timur Doumler
- [ ] Peering forward — C++'s next decade - Herb Sutter
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
