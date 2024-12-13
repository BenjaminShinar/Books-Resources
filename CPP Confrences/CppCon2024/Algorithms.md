<!--
// cSpell:ignore hashers Adler32 inplace
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Algorithms

<summary>
9 Talks about algorithms.
</summary>

- [x] C++26 Preview - Jeff Garland
- [ ] Composing Ancient Mathematical Knowledge Into Powerful Bit-fiddling Techniques - Jamie Pond
- [ ] Designing a Slimmer Vector of Variants - Christopher Fretz
- [ ] Interesting Upcoming Features from Low latency, Parallelism and Concurrency from Kona 2023, Tokyo 2024, and St. Louis 2024 - Paul E. McKenney, Maged Michael, Michael Wong
- [x] So You Think You Can Hash - Victor Ciura
- [ ] Taming the C++ Filter View - Nicolai Josuttis
- [ ] The Beman Project: Bringing Standard Libraries to the Next Level - David Sankel
- [x] When Lock-Free Still Isn't Enough: An Introduction to Wait-Free Programming and Concurrency Techniques - Daniel Anderson
- [ ] Work Contracts – Rethinking Task Based Concurrency and Parallelism for Low Latency C++ - Michael Maniscalco

---

### So You Think You Can Hash - Victor Ciura

<details>
<summary>
Creating a hashing framework.
</summary>

[So You Think You Can Hash](https://youtu.be/lNR_AWs0q9w?si=D6eY4ngakXcwwrrH), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/So_You_Think_You_Can_Hash.pdf).

we use hashing for efficient data retrieval and storage.

> A hashing "framework" for:
>
> - easy experimenting and benchmarking with different hash algorithms
> - easy swapping of hashing algorithms (later on)
> - hashing complex aggregated user-defined types
> - enabling easy comparisons of hashing techniques

the people who write hash algorithms are mathematicians, they focus on the uniformity of the hashing. but the users are developers, who want to consume algorithms, without being concerned about the details.

associative containers, map input into a container based on the hash value, each slot (bucket) stores a set of records.

the hashing function should be:

- deterministic
- uniformity (as much as possible)
- defined range (constrained)
- non-invertible - hard or impossible to reconstruct the data from the hash - not always required

the questions we ask ourselves

> - How should one combine hash codes from your data members to create a "good" hash function?
> - How does one know if you have a good hash function?
> - If somehow you knew you had a bad hash aggregate function, how would you change it for a type built out of several data members (that are not primitive types)?
> - How to separate concerns: hash algorithms from the aggregation of the digest (combine) and from the collection type itself (HashMap, BTreeMap, etc)?

in the following code example, we want to match all records of a customer, but there is no clear unique identifier for the customer. we would need to somehow create a unique key for the customer.

```cpp
class Customer
{
  std::string firstName;
  std::string lastName;
  int         age;
};
std::unordered_map<Customer, Records> customer_records;
```

there is a built in hashing function, <cpp>std::hash</cpp>, with specializations for common data types.

in our example, we can hash each of the members, but then we need somehow combine them together while maintaining the properties of a good hash key.

```cpp
class Customer
{
  std::string firstName;
  std::string lastName;
  int         age;

  std::size_t hash_code() const
  {
    std::size_t k1 = std::hash<std::string>{}(firstName);
    std::size_t k2 = std::hash<std::string>{}(lastName);
    std::size_t k3 = std::hash<int>{}(age);
    return hash_combine(k1, k2, k3);
  }
};
```

so what's really behind the C++ Standard Library hash function and how do we combine them?

this is one such example, that modifies the value of the input seed.

```cpp
template <class T>
inline void hash_combine(std::size_t & seed, const T & v)
{
  std::hash<T> hasher;
  seed ^= hasher(v) + 0x9e3779b9 + (seed<<6) + (seed>>2);
}
```

however, this only works with the standard hash, and depends on the first seed - so if the seed is zero, we get worse behavior.

so we have actually two steps:

1. hashing data member
2. combining the hashes together

as for the hashing function, there are some common algorithms

- FNV-1a
- SipHash
- Spooky
- Murmur
- CityHash

in truth, the algorithm in the standard library is usually FNV-1a (acronym of Fowl-Noll-Vo), which was designed for fast hash-table and checksum usage (not cryptographically secure).

it has two magic constants, the offset basis and the prime.

```cpp
std::size_t fnv1a(void const * key, std::size_t len)
{
  std::size_t h = 14695981039346656037u;
  unsigned char const * p = static_cast<unsigned char const*>(key);
  unsigned char const * const e = p + len;
  for (; p < e; ++p)
  {
    h = (h ^ *p) * 1099511628211u;
  }
  return h;
}
```

#### Externalizing Hashing

looking at this and other hashing algorithms, there is a common anatomy.

> Anatomy of a Hash Function
>
> 1. Initialize internal state
> 2. Consume bytes into internal state
> 3. Finalize internal state to result type (usually size_t)

if our initialization part is costly, we won't want to do it multiple times. we would also want to expose and externalize those steps, so we could modify them independently.

> What we need to do is to repackage the algorithm, in a generic way (to work with all types of hashers), to make the 3 stages above separately accessible:
>
> 1. Init / construction of the hasher
> 2. Write overloads for primitive/std types (append to the hash)
> 3. Finalize function -> size_t
>
> This technique ensures that:
>
> - we no longer need to have a combine step
> - we're using the same hash algorithm for the entire data structure (no special "glue" for intermediate hash codes)

so, we make the previous function into a stateful object.

```cpp
class fnv1a
{
  std::size_t h = 14695981039346656037u; //initialize internal state

public:
  // consume bytes into internal state
  void operator()(void const * key, std::size_t len) noexcept
  {
    unsigned char const * p = static_cast<unsigned char const*>(key);
    unsigned char const * const e = p + len;
    for (; p < e; ++p)
    {
      h = (h ^ *p) * 1099511628211u;
    }
  }

  explicit operator size_t() noexcept // finalize internal state to size_t
  {
    return h;
  }
};
```

we can update our code for the customer class to include this object, so now we don't have special "glue" for the hash combine step.

```cpp
class Customer
{
  std::string firstName;
  std::string lastName;
  int         age;

  std::size_t hash_code() const
  {
    fnv1a hasher;
    hasher(firstName.data(), firstName.size());
    hasher(lastName.data(), lastName.size());
    hasher(&age, sizeof(age));
    return static_cast<std::size_t>(hasher);
  }
};
```

but what if our class is nested inside another class? that annoying hash_combine comes back!

```cpp
class Sale
{
  Customer customer;
  Product product;
  Date     date;

public:
  std::size_t hash_code() const
  {
    std::size_t h1 = customer.hash_code();
    std::size_t h2 = product.hash_code();
    std::size_t h3 = date.hash_code();

    return hash_combine(h1, h2, h3);
  }
};
```

we would want to have just a single hasher, so we need to modify our hash_code function again. so we make the into arguments, let's call the new function "hash_append".

```cpp
class Customer
{
  std::string firstName;
  std::string lastName;
  int         age;

public:
  friend void hash_append(fnv1a & hasher, const Customer & c)
  {
    hasher(c.firstName.data(), c.firstName.size());
    hasher(c.lastName.data(), c.lastName.size());
    hasher(&c.age, sizeof(c.age));
  }
};

class Sale
{
  Customer customer;
  Product product;
  Date     date;

public:
  friend void hash_append(fnv1a & hasher, const Sale & s)
  {
    hash_append(hasher, s.customer);
    hash_append(hasher, s.product);
    hash_append(hasher, s.date);
  }
};
```

we will also need to create specialized overloads of hash_append for primitive types and std defined types.

we can also make this a template function and pass any hasher.

```cpp
template<class HashAlgorithm>
friend void hash_append(HashAlgorithm & hasher, const Customer & c)
{
  hash_append(hasher, c.firstName);
  hash_append(hasher, c.lastName);
  hash_append(hasher, c.age);
}

template <class HashAlgorithm>
void hash_append(HashAlgorithm & hasher, int i)
{
  hasher(&i, sizeof(i));
}

template <class HashAlgorithm, class T>
void hash_append(HashAlgorithm & hasher, T * p)
{
  hasher(&p, sizeof(p));
}
```

the recipe for hashing so far goes by the following logic.

> Even a complicated class is ultimately made up of scalars, located in discontiguous memory.\
> hash_append() appends each byte to the HashAlgorithm state by "recursing down" into the aggregated data structure to find the scalars.
>
> Steps:
>
> 1. Every type has a hash_append() overload
> 1. Each overload will either call hash_append() on its bases and members, or it will
> 1. send bytes of its memory representation to the HashAlgorithm (scalars)
> 1. No type is aware of the concrete HashAlgorithm implementation

There are still questions about special types, such as <cpp>std::optional</cpp> and <cpp>std::variant</cpp>.

to use this thing in our containers, we add another wrapping layer and pass the new type as the template parameter.

```cpp
template <class HashAlgorithm>
struct GenericHash
{
  using result_type = typename hashAlgorithm::result_type;
  template <class T>
  result_type operator()(const T & t) const noexcept
  {
    HashAlgorithm hasher;
    hash_append(hasher, t);
    return static_cast<result_type>(hasher);
  }
};

std::unordered_set<Customer, GenericHash<fnv1a>> my_set;
```

#### Hashing in Rust

looking at how the Rust language does hashing, and what can we learn from it.\
Rust uses traits, and there's a trait for hashing which requires taking a Hasher.

```rust
// Required method
fn hash<H>(&self, state: &mut H)
  where H: Hasher;

// implementation for Customer type
impl Hash for Customer {
  fn hash<H: Hasher>(&self, state: &mut H) {
    self.first_name.hash(state);
    self.last_name.hash(state);
    self.age.hash(state);
    self.premium.hash(state);
  }
}
```

also using a macro to apply hashing on all properties.

```rust
#[derive(Hash)]
struct Customer {
  first_name: String,
  last_name: String,
  age: i32,
  premium: bool,
}
```

also option to ensure equality and hashing don't deviate from another using `#[derive(PartialEq, Eq,Hash)]`

the trait is implemented for almost all types:

```rust
impl Hash for str {
  #[inline]
  fn hash<H: Hasher>(&self, state: &mut H) {
    state.write_str(self);
  }
}
impl Hash for String {
  #[inline]
  fn hash<H: Hasher>(&self, hasher: &mut H) {
    (**self).hash(hasher) // falls back on the &str impl
  }
}
```

the hasher object has it's own protocol (interface) with required methods.

```rust
let mut hasher = DefaultHasher::new();
hasher.write_u32(1989);
hasher.write_u8(11);
hasher.write_i64(1729);
hasher.write_str("Foo");
println!("Hash is {:x}", hasher.finish());
```

there are some predefine hashers available, and there's a builder to create instances of them.

- RandomState
- DefaultHasher
- SipHasher
- Adler32

</details>

### Introduction to Wait-free Algorithms in C++ Programming - Daniel Anderson

<details>
<summary>
Wait free algorithms.
</summary>

[Introduction to Wait-free Algorithms in C++ Programming](https://youtu.be/kPh8pod0-gk?si=QJad5eqaT7_6x0SM), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/When_Lock-Free_Still_Isn't_Enough.pdf)

lock-free design patterns and wait-free designs.

<cpp>std::atomic</cpp>, <cpp>compare_exchange</cpp>.

we use an example of a _sticky counter_ (it gets stuck at zero) to understand the issue.

```cpp
struct Counter {
  // it the counter is greater than zero, add one and return true
  // otherwise do nothing and return false
  bool increment_if_not_zero();

  // decrement the counter
  // if the counter is now zero, return true
  // otherwise return false
  // precondition: the counter is not zero
  bool decrement();

  // return the current value of the count
  uint64_t read();
};
```

This is used for <cpp>std::weak_ptr\<T>::Lock</cpp> in the C++ Standard Library.

we start with a naive implementation

```cpp
struct Counter {
  bool increment_if_not_zero() {
    if (counter > 0) {
      counter++;
      return true;
    }
    return false;
  }

  bool decrement() {
    return (--counter == 0);
  }
  uint64_t read() { return counter; }
  uint64_t counter{1};
};
```

of course, this isn't thread safe, so we add a <cpp>std::lock_guard</cpp>.

```cpp
struct Counter {
  bool increment_if_not_zero() {
    std::lock_guard g_{m};
    if (counter > 0) {
      counter++;
      return true;
    }
    return false;
  }

  bool decrement() {
    std::lock_guard g_{m};
    return (--counter == 0);
  }

  std::mutex m;
  uint64_t counter {1};
};
```

the above implementations is correct, as it is thread-safe by eliminating concurrency, but it's not efficient, because threads have to wait for the lock to be freed.

> Progress guarantees are a way to theoretically categorize concurrent algorithms:
>
> - **Blocking**: No guarantee
> - **Obstruction free** (progress in isolation): A single thread executed in isolation will complete the operation in a bounded number of steps.
>   - Obstruction-free algorithms are immune to deadlock.
> - **Lock free** (at least one thread makes progress): At any given time, at least one thread is making progress on its operation.
>   - Guarantees system-wide throughput. Some operations are always completing, but individual operations are never guaranteed to ever complete.
> - **Wait free** (all threads make progress): Every operation completes in a bounded number of steps regardless of other concurrent operations.
>   - Guaranteed bounded completion time for every individual operation.

so let's look at the lock-free implementation

```cpp
struct Counter {
  bool increment_if_not_zero() {
    auto current = counter.load();
    while (current > 0 && !counter.compare_exchange_weak(current, current + 1)) { } // empty loop, the value is updated each time.
    return current > 0;
  }

  bool decrement() {
    return counter.fetch_sub(1) == 1;
  }

  uint64_t read() { return counter.load(); }
  std::atomic<uint64_t> counter{1};
};

// the underlying algorithm is something like this
compare_exchange(expected&, desired) {
  if (current_value == expected) {
    current_value = desired;
    return true;
  } else {
    expected = current_value;
    return false;
  }
}
```

we use the <cpp>std::atomic\<T>::compare_exchange_weak</cpp> in a loop. we first atomically read the value with the <cpp>.load()</cpp> method, and then we perform the comparison and increment at the same time, and if our value is not the most updated, we update it and try again. if at any time the value is zero, we exit the loop. for decrementing, we can count on the precondition that the counter isn't zero, and call `.fetch_sub(1) == 1` - we decrement by one and check if the value before us was 1, if it was, then it's now zero and the counter is locked at zero.

> The "CAS loop"
>
> - The so-called "CAS loop" (compare-and-swap loop) is the bread and butter of lock-free algorithms and data structures
>
>   - Read the current state of the data structure
>   - Compute the new desired state from the current state
>   - Commit the change only if no one else has already changed it (compare-exchange)
>   - If someone else changed it, try again
>
> - Progress is lock free because if an operation fails to make progress (the compare-exchange returns false) it can only be because a different operation made progress.
> - Progress is not wait free because a particular operation can fail the CAS loop forever because of competing operations succeeding.

#### Towards a Wait-free Algorithm

> A wait-free algorithm can not contain an unbounded CAS loop
>
> - This does not mean you can not use compare-exchange, just not in an unbounded loop!
> - Most wait-free algorithms will make use of atomic read-modify-write operations:
>   - `compare_exchange_weak/strong`(expected, desired): Atomically replaces the current value with desired if current equals expected, otherwise loads the current value
>   - `fetch_add(x)` / `fetch_sub(x)`: Atomically add/subtract x from the given variable and return the original value
>   - `exchange(desired)`: Stores the value desired and returns the old value

our problem is that threads compete with one another to make progress, and are blocking one another. for a lock-free design, we would want our threads to work together and collaborate. the threads need to be able to detect that others are in progress, which will require some re-design. for our example, we also need a way for threads to signal that they have set the counter to zero or are about to.

our first idea is to use some bits as a flags, one flag marks the counter as being zero - regardless of whats really in it. we hid the flag as the topmost bit, so adding to the counter doesn't change it. we also "linearize" the operations, saying that the "order" they happened was different than reality.

(three different iterations)

we need whoever sets the flag to get the correct response from the decrement operation, so it could perform the clean up. only one thread can take get that response.

```cpp
struct Counter {
  static constexpr uint64_t is_zero = 1ull << 63; // flag bit to indicate the value is zero
  static constexpr uint64_t helped = 1ull << 62; // flag bit to indicate the value is zero, but it wasn't set in a decrement operation
  bool increment_if_not_zero() {
    return (counter.fetch_add(1) & is_zero) == 0; // if someone set the flag, return zero
  }

  bool decrement() {
    if (counter.fetch_sub(1) == 1) { // if the atomic was 1 before we decrement it
      uint64_t e = 0;
      if (counter.compare_exchange_strong(e, is_zero))
      {
        return true; // we managed to push the is_zero value into the counter
      }
      else if ((e & helped) && (counter.exchange(is_zero) & helped))
      {
        // the helping bit was set already by a read operation, we remove the helped flag and and put the is_zero flag
        return true;
      }
    }
    return false;
  }

  uint64_t read() {
    auto val = counter.load();
    if (val == 0 && counter.compare_exchange_strong(val, is_zero | helped)) // set both bits
    {
      return 0; // helping!
    }
    return (val & is_zero) ? 0 : val;
  }
  std::atomic<uint64_t> counter{1};
};
```

#### Summary

benchmarking shows better latency for wait-free counters when there are more threads, but it depends on the workloads (reads vs writes), the more writes there are, lock-free algorithms become better. for workloads that focus on reads, wait-free is usually faster.

> Progress guarantees
>
> - Useful theoretical classification of concurrent algorithms that can inform algorithm design
> - Lock-free algorithms guarantee that one thread is making progress, while wait-free algorithms guarantee that every thread is making progress
>
> Wait-free algorithm design
>
> - The bread-and-butter technique is helping. Operations help concurrent operations rather than waiting for them (blocking) or compete with them (lock-free)
>
> Performance Implications
>
> - Never guess about performance
> - But do hypothesize about performance by analyzing an algorithm’s progress guarantees, and use these progress guarantees to guide the design of your algorithm

</details>

### C++26 Preview - The Smaller Features - Jeff Garland

<details>
<summary>
The smaller features we are expecting to get in C++26.
</summary>

[C++26 Preview - The Smaller Features](https://youtu.be/xmqkRcAslw8?si=YBLLVHaTw0zePma_), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Cpp_26_Preview.pdf)

Since 2011, there was a decision to follow "the train model" for C++ releases: features which are ready become part of the standard, and features which aren't ready will be shipped in the next release. the standard is shipped on time, no matter which features aren't in it.

there are large features, like reflection, concurrency, and contracts (probably won't be shipped), but also smaller changes, which are less flashy, but will probably be used more in the day-to-day work of many developers.

topics outline:

- Language & Library
  - debugging
  - structured bindings
- Language
  - Templates
  - Misc
  - ~~Contracts~~
  - ~~Reflection~~
- Library
  - string processing
  - format additions
  - containers
  - ranges
  - utilities
  - general math support
  - constexpr all the things
  - ~~concurrency~~
  - ~~simd~~
  - ~~linear algebra and mdspan~~

#### Language and Library

adding user generated <cpp>static_assert</cpp> messages. using compile-time formating, build-time diagnostics.\
adding a reason for function `= delete` declaration, explaining why a method was removed from a overload set, instead of having a comment. also allows "deleting" free function, not just member functions.\
Making the `assert()` legacy macro user friendly, allowing for a custom message errors.

a new library header <cpp>debugging</cpp> that enables special behavior for debugging mode.

```cpp
#include <debugging>

int main()
{
  std::breakpoint_if_debugging(); // stop if in debugger
}
```

structured binding as a condition, first evaluated and then destructred.

```cpp
//before (needs P2497 update to to_chars)
if (auto result = std::to_chars(p, last, 42)) {
  // okay, use char pointer
  auto [ptr, _] = result;
} else {
  // handle errors
  auto [_, ec] = result;
}

//after:
if (auto [to, ec] = std::to_chars(p, last, 42)) {
  auto s = std::string(p, to);
  // ...
}
```

unnamed placeholder variable, using the `_` symbol like many other languages, can also be used in structured binding. it can't be interacted with.

```cpp
std::lock_guard namingIsHard(mutex); // before
std::lock_guard _(mutex); // after
// Structured binding
[[maybe_unused]] auto [x, y, iDontCare] = f(); // before
auto [x, y, _] = f(); // after
```

adding attributes for structured bindings internally to one member, instead of applying it to all the destructred variables. (goes on the right side, rather than the expected left hand side).
Also a new library feature to get the parts of a complex number into a tuple and therefore a structured binding.

#### Language

adding some characters to the basic character set like C23 does.\
static storage for braced initlazyres, reducing copies at runtime.

pack indexing for templating, using square brackets with indexes inside variadic templates at compile time.

```cpp
// syntax is name-of-a-pack ... [constant-expression]
template <typename... T>
constexpr auto first_plus_last(T... values) -> T...[0] {
 return T...[0](values...[0] + values...[sizeof...(values)-1]);
}
int main() {
  //first_plus_last(); // ill formed
  static_assert(first_plus_last(1, 2, 10) == 11);
}
```

passing a concept or variable template to a template-template as parameters, more sophisticated behavior with type system.\
Variadic friend, removing boiler platecode.

#### Library

string processing changes, such interfacing <cpp>stringstream</cpp> from a string-view, calling <cpp>subview</cpp> on a string, concatenating string and string view together with the `+` operator. adding a string-view interface for <cpp>std::bitset</cpp>.\
changes to <cloud>std::to_string</cloud> to better handle floating point values. make it behave similar to <cpp>std::format</cpp>.\
Testing for success of <cpp>\<charconv></cpp> functions (like <cpp>std::to_chars</cpp>), this is required for the change mentioned above for using strcurted bindings inside conditions.

<cpp>std::format</cpp> gets more type-checking at compile time for the arguments.

| expression                            | result                                      |
| ------------------------------------- | ------------------------------------------- |
| `format("{:d}", "I am not a number")` | compile error (invalid specifier for strings) |
| `format("{:7^\*}", "hello")`          | compile error (should be \*^7)                |
| `format("{:>10}", "hello")`           | ok                                          |
| `format("{0:>{1}}", "hello", 10)`     | ok                                          |
| `format("{0:>{2}}", "hello", 10)`     | compile error (argument 2 is out of bounds)   |
| `format("{:>{}}", "hello", "10")`     | runtime error <– wait why runtime?            |

moving away <cpp>std::vformat</cpp> and using <cpp>std::runtime_format</cpp> with a more consistent api. allowing formatting on pointer types and file-system paths. calling <cpp>std::println()</cpp> without any parameters for an empty line.\
range support for <cpp>std::optional</cpp>, used for pipelines. support for optional on reference types, use monadic function <cpp>and_then</cpp>, <cpp>or_else</cpp>.\
adding the <cpp>std::inplace_vector</cpp>, a vector with a known size, doesn't allocate. has some operations which won't throw.\
<cpp>std::span</cpp> and initializer lists, better conversion from types, adding the <cpp>std::span::at()</cpp> interface.

</details>
