<!--
ignore these words in spell check for this file
// cSpell:ignore Emscripten emcc debuggable protobuf msgpack revent
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Networking & Web

### WebAssembly: Taking Your C++ and Going Places - Nipun Jindal & Pranay Kumar

<details>
<summary>
Basics of using C++ and WebAssembly. live demo.
</summary>

[WebAssembly: Taking Your C++ and Going Places](https://youtu.be/ZS6OUzDFrE0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Draft-CppCon2022-WebAssembly.pdf), [github repo](https://github.com/nipunjindal/cpp-con-wasm)

Agenda:

- basics of WebAssembly
- Tooling and setup
- Porting a C++ project
- Debugging WebAssembly

WebAssembly is not a language you write in, it's something you compile into. supported by most browsers (although chrome has the best support), and is faster than regular JavaScript.

Advantages:

- efficient and fast
- open and debuggable
- pluggable
- safe and innovative

some notable project have been ported into WebAssembly already: Doom3, unity engine, Vim, SQLite.

- WAT - WebAssembly Text format
- WAST - a superset of WAT/

we can use _Emscripten_ (emcc) as the toolchain, it's a wrapper of LLVM, so it's just like clang in some ways. there are some more flags which are needed, such as the memory size, allowing threading, allowing fetching, etc...

```sh
emcc -v # get version
emcc test/hello_world.c -o hello.js # compile
emcc test/hello_world.c -o hello.html # compile with a html file as well
```

when we compile, we get a wasm file, a js file (which is a wrapper over the wasm) to consume with a JavaScript glue layer such as Node.js, and optionally, an html file to try out directly in the browser.

the code is unoptimized by default, but we can added the optimization flags, the first compilation takes longer, but the next times are cached.

Emcc has it's own implementation of system libraries, WASI - WebAssembly System Interface. there are also some pre-ported popular libraries such as boost.
WASM doesn't work directly with the DOM, so it connects with the JavaScript to manipulate the DOM html.

#### Live Demo

we can run the examples in docker, we can export the C++ function and make them usable in JavaScript. and we can write JavaScript code inside the c++ file and get the data from it. we can call the functions from the browser console and it will show on the the html.

if we want, we can expose our C++ objects outside and have them interact with JavaScript, the lifetime of objects isn't managed by js,so we need to call `delete` ourselves. templates can also be used, shared pointers under the "smart_ptr" moniker, unique pointer are natively supported.

#### Porting a C++ Project

Fetching data from the web (pass the `-sFetch` flag), access to a storage file system (in a sandbox version), we can define how the virtual file system is simulated. we can use JavaScript exceptions or WebAssembly exceptions, depending of the flag we pass. something about the event loop.

we can use the browser persistent storage (IDBFS), or create fetch queries from inside the C++ code. we can even use threading, and in the example we make sure the logger always uses a certain thread.

Debugging in WebAssembly uses the DWARF file format, we need to pass the debug flag (`-g`) with the correct level of maintaining debug information. we also need the correct extension in the browser development tools.

of course, there are pitfalls when using WebAssembly, the file system, the memory limits, exception handling, and having the main thread blocked,

</details>

### A Faster Serialization Library Based on Compile-time Reflection and C++ 20 - Yu Qi

<details>
<summary>
A serialization library based on compile-time reflection.
</summary>

[A Faster Serialization Library Based on Compile-time Reflection and C++ 20](https://youtu.be/myhB8ZlwOlE),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/A-Faster-Serialization-Library-Based-on-Compile-time-Reflection-and-C-20-Yu-Qi-CppCon-2022.pdf).

#### Introduction of _struct_pack_

_struct_pack_ is a serialization library, based on c++20 compile time reflection. other libraries of protobuf, msgpack or the boost.serialization library.

struct_pack is a non intrusive way of serializing, tha doesn't require specifying the a special function.

#### How to Make Serialization Faster

the first way to get faster serialization is by dropping the type information, we don't pack or unpack the field information, only the values. there is no need to do that because we can employ compile time reflection.

```cpp
struct dummy {
  int id;
  float number;
  std::string str;
};

void serialize(dummy t) {
  // reflect all fields of dummy into items…
  visit_members(t, [](auto &&...items){
  // serialize each item.
  serialize_many(items...);
  });
}
```

using concepts, aggregate types, member count, etc...

a possible problem is deserialization - how to check the type for de-serialization, they get around it by encoding all the fields into a single hash code. the deserializing code can re-create the hash code because of compile time reflection.

#### Compile-time Type Hash

steps:

- mapping each member type to a unique type id
- generate a compile time string from the those types.
- create an MD5 hash from the string

```cpp
template <typename T>
void serialize(T &&t) {
  // generate hash code of t.

  constexpr uint32_t types_code =
  get_types_code<decltype(get_types(std::forward<T>(t)))>();
  // serialize the hash code for deserialization type checking

  std::memcpy(data_ + pos_, &types_code, sizeof(uint32_t));
  pos_ += sizeof(uint32_t);
  // serialize t
  serialize_one(t);
}
```

we can get better performance if the object is trivially copyable, and even more for continues container types.

#### Backward Compatibility

if an object changed (fields were added), we need to keep the serialization backward compatible. this is done with a special type `compatible<T>` that tells the serializer to omit them when creating the hash. this means those fields must go at the end of the object.

#### Benchmark

benchmarking shows that the new library is much faster than traditional ways of serializing and deserializing objects.

</details>

### Bringing a Mobile C++ Codebase to the Web - Li Feng

<details>
<summary>
TypeScript, webAssembly and C++.
</summary>

[Bringing a Mobile C++ Codebase to the Web](https://youtu.be/-K72c5_IvIw), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon2022.pptx), [github](https://github.com/snapchat/djinni).

Djinni is a project that bridges between C++ code and other programming languages. it started with objective C, originally developed by dropbox, and now open sourced by snapchat.

Djinni's newest feature is webAssembly support. this allows them to use the same code in mobile and desktop, so there is no additional learning required to add another platform. they use external binding tools to link C++ and typescript together. c++ containers are mapped to typescript native objects, asynchronous objects are matched with C++ future object.

(demo of compiling code, bundling, etc..)

explaining how data passes the boundary through webAssembly, memory management, data marshalling, exception handling (bi-directional).

</details>

### Structured Networking in C++ - Dietmar Kühl

<details>
<summary>
Live demo of sender/receiver of asynchronous operations.
</summary>

[Structured Networking in C++](https://youtu.be/XaNajUp-sGY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Structured-Networking-CppCon.pdf), [github repository](https://github.com/dietmarkuehl/kuhllib).

integrate networking with concurrency, have the same approach for both cases.

> Structured Concurrency
>
> 1. decompose work into senders (representing work).
> 2. combine the work representation with receiver.
> 3. start the resulting entity.

Receiver: Destination for Results

- `get_env(receiver) -> env`
- `get_stop_token(env)`
- `get_allocator(env)`
- `set_value(receiver, results…)` - green flow
- `set_error(receiver, error)` - error flow
- `set_stopped(receiver)` - cancellation

Sender: Description of Work

- factory schedules
- adapters - (`s | then(f)`)
- consumer - `start detached(s)`
- `connect(sender, receiver) -> operation_state`

Operation State: Ready to Execute Task

- `start(operation_state)` &rarr; call one of the completions

#### Coding example

taking an synchronous echo-server and converting it to asynchronous one. the code is in the github under "src/toy" folder.

we have a task implemented as a coroutine. we start with a blocking server that can only handle one client.

this is the blocking code:

```cpp
for (int i{}; i!=2; ++i){
  sockaddr_storage clnt{};
  socklen_t len{sizeof(clnt)};
  desc c{accept(server.fd, (sockaddr*)&clnt,&len)};

  char buff[4];
  while (size_t n = read(c.fd, buff. sizeof(buff))) {
    for (size_t o{}, w(1); o!=n && 0 < w; o += w) {
      w = write(c.fd, buff + o, n-o);
    }
  }
}
```

so now want it to not block, so we change the _accept_ function to a coroutine and <cpp>co_await</cpp> it. we implement is as an lambda and pass the "server" object. we need to implement this "async_accept", so we make it a struct and start filling it up. we pass things by value for this demo, rather than deal with forwarding.

(lots of live coding that i gave up on doing myself)

</details>

##

[Main](README.md)
