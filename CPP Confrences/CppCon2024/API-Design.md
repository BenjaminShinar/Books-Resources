<!--
// cSpell:ignore beman
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## API Design

<summary>
6 lectures about API design.
</summary>

- [x] Creating a Sender/Receiver HTTP Server - Dietmar Kühl
- [ ] Deciphering C++ Coroutines - Mastering Asynchronous Control Flow - Andreas Weis
- [ ] Irksome C++ - Walter E Brown
- [ ] Security Beyond Memory Safety - Using Modern C++ to Avoid Vulnerabilities by Design - Max Hoffmann
- [x] Hidden Overhead of a Function API - Oleksandr Bacherikov
- [x] Reflection based libraries to look forward to - Saksham Sharma

---

### C++ Reflection Based Libraries to Look Forward To - Saksham Sharma

<details>
<summary>
What can reflection do for us?
</summary>

[C++ Reflection Based Libraries to Look Forward To](https://youtu.be/7I40gHiLpiE?si=1uc9RCMOCNRnEqND), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Reflection_Based_Libraries_to_Look_Forward_To.pdf)

#### What is reflection

reflection is code that has information about other code, and operate over the information.

```cpp
class MyClass {
  int a;
  int b;
};

for (auto member_info : gimme_class_member<MyClass>()) {
  std::cout << "member - " << member_info.name() << '\n';
}
```

this is similar, but not the same as template meta-programming. templates were designed to write generic code, but over time, the usage shifted to features that check code at compile time. those new features are what reflection wants to achieve.

#### Reflection in other languages (Go, Python, Java)

> Python: At runtime your code can
>
> - Access class layout
> - Modify your class to instrument function calls.
> - Change what it means to access a field on an object.
> - Add or remove methods or attributes from any object.

for example, we modify the `copy` function of the class if it exists.

```python
def modify_cls(cls):
  if not hasattr(cls, "copy"):
    return cls
  orig_copy = cls.copy

  def _wrapped_copy(obj):
    print("Calling wrapped copy")
    attrs = obj.__dict__.keys()
    print("Attributes: " + " ".join(attrs))
    result = orig_copy(obj)
    return result
  cls.copy = _wrapped_copy

class MyClass:
  def __init__(self, x):
    self.x = x
  def copy(self):
    return MyClass(self.x)

modify_cls(MyClass)
MyClass(2).copy()
```

in [golang](https://go.dev/blog/laws-of-reflection)

> - Golang is a compiled but duck-typed language.
>   - Well, structurally typed, but close enough
> - Runtime reflection similar to python.
> - No special compile time constructs
> - Provides a package reflect to get "reflection values".

```golang
type T struct {
  A string
  B int
}

t := T{"CppCon!", 24}
s := reflect.ValueOf(&t).Elem()
typeOfT := s.Type()
for i := 0; i < s.NumField(); i++ {
  f := s.Field(i)
  fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
}
```

and java is a bit similar to python, because the type system is visible at runtime.

``` java
// Surprising, lookup types with string!
Class cls = Class.forName("method1");
Method methodList[] = cls.getDeclaredMethods();
for (int i = 0; i < methodList.length; i++) {
  Method m = methodList[i];
  System.out.println("name = " + m.getName());
  System.out.println("decl class = " + m.getDeclaringClass());
}

Class pvec[] = m.getParameterTypes();
for (int j = 0; j < pvec.length; j++)
  System.out.println("param #" + j + " " + pvec[j]);
```

anything that uses strings to access methods is error prone, and anything that runs in runtime is costly.

#### Reflection in C++ as per P2996

C++ reflection is planned to run on compile time, and is already well defined and has two working implementations.

it uses a new unary reflection operator `^` to lift objects into the land of reflection(it might end up as `^^`). we also have a new type <cpp>std::meta::info</cpp> which is the result of the reflection operator, it doesn't have any methods. the splice operator `[: r :]` takes a reflection value and splices it back into regular code.

```cpp
struct MyStruct {
  static int a;
  static int b;
};

constexpr auto elem = ^MyStruct::a;
std::cout << [:elem:] << '\n';
```

we could use reflection to create classes programmatically with <cpp>std::meta::define_class</cpp>. all reelection code is compile time, using <cpp>consteval</cpp> to ensure they are never called during runtime.

for example, we could use reflection to generate Enum to String, command line parsing, transforming an array of struct into struct of arrays.

#### Reflection libraries!

> - Reflection is a really powerful language feature - With great power comes great responsibility
> - Easier to write general-purpose / boilerplate-reducing libraries
> - Solve multiple pain-points through a single feature - The hallmark of a useful language feature

we don't want end users to write reflection, and we don't think beginners should concern themselves with it either, so we will create libraries that use reflection internally.

we can create simplified implementation

- automatic Python bindings
- ABI hashing ( <cpp>boost::abi_hash</cpp>?)
- A duck-typed <cpp>std::any</cpp> (<cpp>boost::virtual_any</cpp>?)

running python and C++ at the same process, and having python manipulate the C++ code. so we need to expose the C++ objects to the python code. today there's a lot of boilerplate code t expose the code, providing the names, types and if it's a value type or a reference type, etc...\
our goal is to remove that long code and replace it with a single function that provides the same functionality.

this is a simple example, it still lacks customization, and handling edge_cases.

```cpp
template <typename T> object make_python_type() 
{
  std::string cls_name{meta::identifier_of (^T)};
  auto type_obj = class_<T>(cls_name.c_str(), no_init);
  [:expand(meta::members_of (^T)):] >> [&]<auto e> 
  {
    if constexpr(!meta::is_public(e))
    {
      return;
    }

    std::string name{meta::identifier_of(e)};
    if constexpr(meta::is_nonstatic_data_member(e))
    {
      type_obj.def_readwrite(name.c_str(), &[:e:]);
    }

    if constexpr(meta::is_function(e) && !meta::is_constructor(e) && !meta::is_destructor(e)) 
    {
      using return t = typename return type<decltype(&[:e:])>::type;
      if constexpr(!std::is_reference_v<return_t>) {
        type_obj.def(name.c_str(), &[:e:]);
      }
    }
  }:
  return type_obj;
}
```

an alternative option is user-defined properties that we just tag objects with an attribute that marks them as exposed.

the next topic is ABI hashing, which takes the type memory layout and hashes them, this saves us sending the entire schema in the header of each message, and it gives us a way to identify different versions of the schema (if some server is still using the old schema). there are other options, each with different problems and limitations.

> - Is a decent test of the capabilities of reflection (P2996).
> - Requires recursively computing the hash of types. Avoid cycles!
> - Requires a compile time hashing function.
> - Requires full visibility into the class' data layout - Sounds scary actually, private members!

recursive code that does hashing for each base class and member of the type.

the next topic was a python <cpp>std::any</cpp>, duck-typing using reflection to support message passing. we will create a new type called "virtual_any" which is a virtual interface. then we can access attributes based on the name (string value), regardless of what class is it. it will still use RTTI (run-time type information).\
but maybe we could store the hashing of the types we used into the virtual any and get away from RTTI, or use some linker magic.

#### Alternatives ways to achieve "reflection"

for now, we use stuff that is similar to reflection, like manually annotating code (python binding). code generation tools like protobuf and Apache Avro, or rely on AI and LLM for code completion.
</details>

### Hidden Overhead of a Function API - Oleksandr Bacherikov

<details>
<summary>
Costs that happen when we call a function. and how can we bring them down.
</summary>

[Hidden Overhead of a Function API](https://youtu.be/PCP3ckEqYK8?si=8MGrFo0PcAnjvodq),[event](https://cppcon2024.sched.com/event/1gZeD/hidden-overhead-of-a-function-api), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Hidden_Overhead_of_a_Function_API.pdf), [Not Leaving Performance On The Jump Table - Eduardo Madrid - CppCon 2020](https://youtu.be/e8SyxB3_mnw?si=OKRSWNxcz1j5zpRX).

#### Introduction

> Tony Van Eerd: "people are not writing enough functions"
>
> When people finally start writing more functions, we'd prefer to get only the well designed ones!\
> When talking about performance, we typically think about the function logic. We'll see that a well designed function API can have an even larger impact.
>
> How will we compare performance?
>
> - Benchmarks at this low level are not too reliable, and also don't represent performance in large projects well.
> - Dynamic instruction count is more reliable on modern CPUs.
> - We'll use simple examples, so that we can just compare the number of instructions generated by a compiler.

according to a research by BOLT, about 30% of work is getting data into the instruction pipeline, so there's room for an improvement boost.

there is the option to inline function code, the effect on performance is mixed, sometimes it helps, sometimes it doesn't.

when we write code, it's then generated into system calls, which means either windows ABI, or system V system calls and other platform specific calls.

a good place to start is by looking at the C++ Core guidelines.

#### Return value

> F.20: For “out” output values, prefer return values to output parameters.\
> Reason: A return value is self-documenting, whereas a & could be either in-out or out-only and is liable to be misused.

```cpp
#include <memory>

std::unique_ptr<int> value_ptr() {
  return nullptr;
}

void output_ptr(std::unique_ptr<int>& dst) {
  dst = nullptr;
}
```

even at this simple code above, we can see performance difference. this becomes more evident by checking the calling code.

```cpp
#include <memory>

// avoid inline by forward declaration
std::unique_ptr<int> value_ptr();
void output_ptr(std::unique_ptr<int>& dst);

int value_ptr_call() {
  auto ptr = value_ptr();
  return *ptr;
}

int output_ptr_call() {
  std::unique_ptr<int> ptr;
  output_ptr(ptr);
  return *ptr;
}
```

the next guideline is about object initialization.

> ES.20: Always initialize an object.\
> Reason: Avoid used-before-set errors and their associated undefined behavior. Avoid problems with comprehension of complex initialization. Simplify refactoring.

there were options to have deferred construction of the parameter, but it's still cumbersome.

> F.26: Use a unique_ptr\<T> to transfer ownership where a pointer is needed.\
> Reason: Using unique_ptr is the cheapest way to pass a pointer safely.

actually, even the simplest <cpp>std::unique_ptr</cpp> has some costs over raw pointers,

> F.26: Use a unique_ptr\<T> to transfer ownership where a pointer is needed.\
> Reason: Using unique_ptr is the cheapest way to pass a pointer safely.

we can create a wrapper over an integer value to see additional overhead, this comes from from it being a non-trivial return type. we need to make the wrapper trivial, but it's still not enough, we need to remove the destructor, and it helps a bit. it turns out that x86 architecture can only return fundamental types in the registers, regardless of the size.

> C.20: If you can avoid defining default operations, do.\
> Reason: It's the simplest and gives the cleanest semantics.(Note This is known as "the rule of zero".)

even if we look at popular libraries, such as <cpp>std::chrono</cpp>, this wasn't done because it would effect performance in other cases.

<cpp>std::pair</cpp> is trivially destructible since C++17, if the elements themselves are., <cpp>std::tuple</cpp> is never trivially move constructable. this means they might have performance costs.

RVO - return value optimization (copy elision), this is part of the standard since C++17. there is a problem that effects containers where copy constructor is used instead of the move constructor.

there are valid cases for output parameters, like in the <cpp>std::ranges</cpp> library.

#### Parameter passing

pass by value is usually better than passing by reference, provided the object is small enough to fit into a register. there is also an effect on calling opaque functions inside a function, a reference can be changed from another function, even if it's not directly passed to it.\
perfect forwarding is still a reference, so <cpp>std::forward</cpp> isn't passed inside registers.

besides built-in types, there are some other standard types the standard says we should pass by value, such as <cpp>std::span</cpp>, <cpp>std::span_p</cpp> and <cpp>std::mdspan</cpp>. however, they aren't free for all platforms. this is, again, because of some architecture specifications about what can be passed in registers.

adding empty parameters can also have affect(this can happen with tag dispatch).

#### Multiple parameters

chaining function calls can behave differently depending on the parameter order, this is because the order of assigning parameters is fixed, so we might have swaps.

the guidelines also say that we shouldn't pass an array as a single parameter, and we should prefer using non-owing <cpp>std::span</cpp>. this can have an affect on performance.

prefer functions with a smaller number of arguments.

> Most important guidelines to avoid function call overhead
>
> - Return by value
> - Pass “trivial” types by value, others by reference
> - Follow the Rule of 0 (or at least support trivial copy)
> - Make APIs consistent
> - Understand abstractions cost on your target platform

</details>

### How to Use the Sender/Receiver Framework in C++ to Create a Simple HTTP Server - Dietmar Kühl

<details>
<summary>
live coding session of creating an HTTP server.
</summary>

[How to Use the Sender/Receiver Framework in C++ to Create a Simple HTTP Server](https://youtu.be/Nnwanj5Ocrw?si=cERC9Qcd_wPab3Zx), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Creating_a_Sender_Receiver_HTTP_Server.pdf), [event](https://cppcon2024.sched.com/event/1gZeX/creating-a-senderreceiver-http-server), [additional github repository](https://github.com/bemanproject/net29).

> Objectives:
>
> - Create a basic HTTP server.
> - Allow a single-threaded server handling multiple clients.
> - Use the sender/receiver asynchronous framework.
> - Use a minimalistic sender/receiver networking interface
>
> Basic Design:
>
> - `main()` runs an event loop for network and timer events.
> - It uses an <cpp>async_scope</cpp> for outstanding work.
> - Initial work consist of accepting incoming client connections.
> - Each client processes requests until an error is received

Using an implementation from the *Beman Project* which follows the standard specification.

starting from the empty example. it contains the header files and some name space aliases. we will build up from it.

```cpp
#include <beman/net29/net.hpp>
#include <beman/execution26/execution.hpp>
#include "demo_algorithm.hpp"
#include "demo_error.hpp"
#include "demo_scope.hpp"
#include "demo_task.hpp"
#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#include <string_view>
#include <unordered_map>

namespace ex  = beman::execution26;
namespace net = beman::net29;
using namespace std::chrono_literals;

// ----------------------------------------------------------------------------

std::unordered_map<std::string, std::string> files{
  {"/", "examples/data/index.html"},
  {"/favicon.ico", "examples/data/favicon.ico"},
  {"/logo.png", "examples/data/logo.png"},
};
```

we start by accepting connections, we work backwards, we need a stream, so we need an acceptor, for the acceptor we need an endpoint and context, and to run the server we need an asynchronous execution via a coroutine with a scope.

(live coding session).

```cpp
auto process(auto& stream, auto const& request) -> demo::task<>
{
  std::string method, url, version;
  std::string body;
  std::ostringstream out;
  if (std::istringstream(request) >> method >> url >> version && files.contains(url))
  {
    out << std::ifstream(files[url]).rdbuf();
    body = out.str();
    out.str({});
  }

  out << "HTTP/1.1 "<< (body.empty() ? "404 not found" : "200 found")  << "\r\n"
  << "Content-Length: " << body.size() << "\r\n"
  << "\r\n"
  << body;
  auto response {out.str()};
  co_await net::async_send(stream, net::buffer(response));
}

auto timeout(auto scheduler, auto duration, auto& sender)
{
  return demon::when_any(
    std::move(sender()),
    net::resume_after(scheduler, duration) 
    | demo::into_error([]() { return std::error_code(demo::timeout, demo::category());})
    ) | demo::into_expected();
}

auto make_client_handler(auto scheduler, auto stream) -> demo::task<>
{
  char buffer[16];
  std::string request;
  while (true)
  {
    try {
      if (auto n = co_await timeout(scheduler, 2s, net::async_receive(stream, net::buffer(buffer))))
      {
        std::string_view data(buffer, n.value());
        std::cout << "received data=" << data << '\n';
        request += data;
        if (request.find("\r\n\r\n") != request.npos) // end of data
        {
          co_await process(stream, request);
        }
      }
      else 
      {
        std::cout << "ERROR (VIA expected): " << std::get<0>(n.error()).message() << '\n';
        break; // break while look
      }
    }
    catch (std::variant<std::error_code> const& ex) {
      std::cout << "ERROR: " << std::get<0>(ex).message() << '\n';
      break; // exit  while loop;
    }
  }
  co_retrun;
}

auto main() -> int
{
  net::io_context context;
  net::ip::tcp:endpoint endpoint(net::ip::address_v4::any(), 12345);
  net::ip::tcp:acceptor acceptor(context, endpoint);
  demo::scope scope;
  scope.spawn(std::invoke([](auto scheduler, auto& scope, auto& acceptor)-> demo::task<> {
    while (true)
    {
      auto[stream, address] = co_await net::async_accept(acceptor);
      std::cout << "received client:" << address << '\n';
      scope.spawn(make_client_handler(scheduler, std::move(stream)));
    }
  }, context.get_scheduler(), scope, acceptor)); // execute inside a scope

  context.run();
}
```

from another terminal, we can curl to our server with `curl http://localhost::12345` or keep an open connection with `telnet localhost 12345` to see timeouts.

</details>

### Deciphering C++ Coroutines - Mastering Asynchronous Control Flow - Andreas Weis

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Deciphering C++ Coroutines - Mastering Asynchronous Control Flow](https://youtu.be/qfKFfQSxvA8?si=XC3uguw1axRK3txD), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Deciphering_Cpp_Coroutines.pdf), [event-1](https://cppcon2024.sched.com/event/1gZeS/deciphering-c-coroutines-mastering-asynchronous-control-flow),[event-2](https://cppcon2024.sched.com/event/1lQEu/deciphering-coroutines-recap-and-prerequisites).

expanding on a talk from CppCon 2022 -[Deciphering C++ Coroutines - A Diagrammatic Coroutine Cheat Sheet](https://youtu.be/J7fYddslH0Q).


</details>
