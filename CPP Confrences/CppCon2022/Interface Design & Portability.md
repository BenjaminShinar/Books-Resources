<!--
ignore these words in spell check for this file
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Interface Design & Portability

### Purging Undefined Behavior & Intel Assumptions in a Legacy C++ Codebase - Roth Michaels

<details>
<summary>
Examples of Undefined behavior in a large codebase.
</summary>

[Purging Undefined Behavior & Intel Assumptions in a Legacy C++ Codebase](https://youtu.be/vEtGtphI3lc), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Purging-Undefined-Behavior-and-Intel-Assumptions-in-a-Legacy-Codebase-CPPCon2022-Roth-Michaels.pdf)

the tools for undefined behavior are getting better each year.

#### Undefined Behavior (UB) Overview

> Behavior of the C++ Abstract Machine
>
> - Defined Behavior: Deterministic behavior specified by the standard.
> - Implementation-Defined Behavior: Platforms can determine behavior, must document.
> - Unspecified Behavior: Non-deterministic behavior; suggested by standard, documentation not required.
> - Undefined Behavior: Standard makes no guarantees, compilers can assume this never happens, anything is allowed.

functions can be safe for all input, safe for some input, or unsafe for every possible input, all sorts of examples of poorly written functions that create undefined behavior.

some code works on intel machines, but then it needed to work on other kinds of hardware. in large code bases with a long legacy, there can be many undefined behavior hiding there.

#### Interesting Bugs Caused by Invoking UB

an example of undefined behavior, using a temporary value. the fix was to move the object declaration outside and use it as an lvalue.

```cpp
void DrawCircleOutline(float x, float y,
 float radius) {
  auto e = ellipse{x, y, radius, radius};
  auto s = stroke<ellipse>{e};
  auto p = transform<decltype(s)>{
  s, getTransform() // fix by moving this outside
  };
  m_rasterizer->addPath(p);
}

template <class V, class T>
  class transform {
  public:
  transform(V& source, const T& tr)
  : m_source(&source)
  , m_transform(&tr) {}
  private:
  V* m_source;
  const T* m_transform;
};
```

another undefined behavior because of a windows update. there is uninitialized member without a default value. it got bad when they saved the data with compression. the code to calculate the size giving different results. the fix was to replace the line with `uint32_t snapshotColor{};` to make it default initialized.

```cpp
struct SnapshotData {
  std::vector<float> frequencies{20.f, 20000.f};
  std::vector<Float> dBAmps{2.f, -200.f};
  uint32_t snapshotColor; // the fix was to add {} to default initialize
  unsigned opacity{191};
  bool visible{false};
  bool enabled{false};
};

std::vector<std::byte> compress(std::byte*, std::size_t size);
std::vector<std::byte> saveState() {
  const auto size = [] {
    auto s = SnapshotData{};
    compress(reinterpret_cast<std::byte*>(&s), sizeof(s));
  }(); //IIFE
  auto state = std::vector<std::byte>(size);
  auto s = SnapshotData{};
  auto data = compress(reinterpret_cast<std::byte*>(&s), sizeof(s));
  std::copy(data.begin(), data.end(), state.begin());
}
```

another bug about alignment, which worked correctly, until it didn't. it's not worth understanding how UB worked before, just fix it.

#### Culture and Tooling Changes to Fight UB

start taking UB seriously, teaching about it, and stop writing it. have to CICD warn about UB. add the correct warnings, especially for new and changed code. use sanitizers, static analyzers.\
we can write our our clang rules matcher to get warnings or to replace bad code.\
make a plan and work towards having the entire test process sanitized.

</details>

### Managing External API's in Enterprise Systems - Pete Muldoon

<details>
<summary>
Detailing a process of moving between API versions over time.
</summary>

[Managing External API's in Enterprise Systems](https://youtu.be/cDfX1AqNYcE), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/ManagingApis3c.pdf)

Using apis, not just writing them. based on engineering decisions which were made, and how the system operates in the real world. the main focus is about handling change in a live production system.

Enterprise systems use many API's

- database
- cache
- IPC (inter process communication)
- Functionality

many talks focus on the producer side of API's, rather than how to consume them.

producer side:

- general use / flexible
- many available modes of operations
- testability
- ergonomics

consumer:

- specific use / outcomes
- constrained modes of operation
- code efficiency / bundling

so the producer and the consumer have different goals, so the consumer usually has a wrapper abstraction over the general api.

we start with a general api and we constrain it to be using move-only responses, and only be asynchronous, so we use a std::future object and a unique_ptr.

showing mock testing, introducing virtual functions, using dependency injection, showing the green and ren path.

but eventually there will be requirement changes. we would to decommission and phase out old api's.

> Prioritize Migration Path:\
> Requirements:
>
> - Localize meaningful changes
> - Keep global usage/calling semantics unchanged
> - Minimize throw-away work
> - Final decommissioning simple

so we go another level of wrapper, a service proxy. and then some other changes arise.

> API USE:\
> channel good practices via wrapper with constraints
>
> - move only data
> - future/promise semantics
> - layer of abstraction

a large scale change of requirements without large change of code, reducing risk, reducing life-time complications, mitigating production risk.

</details>

### Back to Basics: C++ API Design - Jason Turner

<details>
<summary>
Ways to make our APIs hard to use wrong.
</summary>

[Back to Basics: C++ API Design](https://youtu.be/zL-vn_pGGgY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-2022-Jason-Turner-API-Design-Back-to-Basics.pdf).

> "Make your API hard to use WRONG"

we start with `bool empty() const;` from the c++98 std::vector class.\
Is it verb? a check? can we drop the return value? rewrite into `[[nodiscard]] bool is_empty() const;` is there error handling? (what happened if it was moved from - must be in a valid space). lets add `noexcept`.

> "The two hardest problems in computer science are
>
> 1. Cache invalidation
> 2. Naming
> 3. Off-by-One errors
> 4. Scope Creep
> 5. Bounds checking"

C++23 will allow for `[[nodiscard]]` for the call operator of a lambda. we can also have this for a type, or for constructors. we can have a reason msg as part of it. should be used extensively, any non-mutating getter/accessor/const function should have it. this is one thing that could be checked and enforced by static analysis tools.

`noexcept` is a promise to the user and compiler that a function must not throw an exception, and if it does, then `terminate` must be called.

moving to the next example, a factory function.
`Widget *make_widget (int widget_type);`.\
the first problem is the unbounded integer parameter, we don't know about the return value, can we ignore it? do we need to `free` or `delete` it? we should rewrite it into `[[nodiscard]] std::unique_ptr<Widget> make_widget(WidgetType type);`. it's still possible to use this wrong.

```cpp
enum class WidgetType{
  Slider = 0,
  Button = 1
};
// wrong use
auto widget = make_widget(static_cast<WidgetType>(-42));
```

we still don't have an error handling, is it possible to fail?

the lesson is to **Never return a raw pointer**. even use a wrapper type (alias) to document the behavior (owning pointer, none-owning pointer). in general, we should have a consistent error handling strategy in the library, and we should avoid making the user check errors with _out-of-band error reporting_ (things like `get_last_error()` and `errno`) - they don't work well with concurrent code. and we want to make our error impossible to ignore. we don't want to be in an unknown state. avoid `std::optional<>` to indicate error condition, `std::expected<>` (c++23) is better.

`FILE *fopen(const char *pathname, const char* mode);`

error handling, possible to drop the return value and leak the resource, what is the format for <kbd>mode</kbd>?. we could start with a wrapper to implement RAII on the unique pointer, and make the string arguments harder to get wrong. we should be careful of having easily swappable parameters. clang-tidy has a check for it. but as long as we have implicit conversion, we can easily write bad code. *const char \*\*, *std::string*, *std::string_view*, *std::path\* are all implicitly constructable.

```cpp
using FilePtr = std::unique_ptr<FILE, decltype([](File *f){fclose f;})>;
// option 1
[[nodiscard]] FilePtr fopen (const char *pathname, const char *mode);
// option 2
[[nodiscard]] FilePtr fopen (const std::filesystem::path &path, std::string_view mode);

// still wrong usage - easy to write
auto file = fopen("rw+","my/file");
```

we can avoid the implicit conversions but creating an explicitly deleted 'catch-all' template.

```cpp
void fopen(const auto &,const auto &) = delete;
```

> - any function can be `=delete`.
> - if you `=delete` a template, it will become the match for any non-exact parameters, and prevent implicit conversions.

returning to the widget example. if we know at compile time all the possible derived classes, we can create a template.

```cpp
template<typename WidgetType>
[[nodiscard]] WidgetType make_widget() requires(std::is_base_of_v<Widget,WidgetType>);
```

if pass a pointer to a function, we should check it for `nullptr`, or take a reference instead.

its always a good idea to run fuzzing tests on the APIs, finding all the edge cases which we didn't consider.

#### Summary

> - Try to use your API incorrectly.
> - Use better naming.
> - Use `[[nodiscard]]` (with reasons) liberally.
> - Never return a raw pointer.
> - Use `noexcept` to help indicate the type of error handling.
> - Provide consistent, impossible to ignore, in-band error handling.
> - Use stronger types and avoid default conversions.
> - _(Sparingly) delete problematic overloads / prevent conversions._
> - Avoid passing pointers (only to be used for single/optional objects).
> - Avoid passing smart pointers.
> - Limit your API as much as possible.
> - Fuzz your API.

</details>

##

[Main](README.md)
