<!--
// cSpell:ignore hashers Adler32
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Algorithms

<summary>
9 Talks about algorithms.
</summary>

- [ ] C++26 Preview - Jeff Garland
- [ ] Composing Ancient Mathematical Knowledge Into Powerful Bit-fiddling Techniques - Jamie Pond
- [ ] Designing a Slimmer Vector of Variants - Christopher Fretz
- [ ] Interesting Upcoming Features from Low latency, Parallelism and Concurrency from Kona 2023, Tokyo 2024, and St. Louis 2024 - Paul E. McKenney, Maged Michael, Michael Wong
- [ ] So You Think You Can Hash - Victor Ciura
- [ ] Taming the C++ Filter View - Nicolai Josuttis
- [ ] The Beman Project: Bringing Standard Libraries to the Next Level - David Sankel
- [ ] When Lock-Free Still Isn't Enough: An Introduction to Wait-Free Programming and Concurrency Techniques - Daniel Anderson
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
