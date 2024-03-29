<!--
ignore these words in spell check for this file
// cSpell:ignore saxpy
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Resource Management

### Back to Basics: C++ Move Semantics - Andreas Fertig

<details>
<summary>
Basic overview of move semantics
</summary>

[Back to Basics: C++ Move Semantics](https://youtu.be/knEaMpytRMA), [slides](https://andreasfertig.com/talks/dl/afertig-2022-cppcon-back-to-basics--move-semantics.pdf)

an analogy of moving an apartment, buying everything is "copy", while moving the furniture into the new apartment is "move".

move operations work great with dynamic memory, it's about swapping pointers, if we don't have pointers to data on the heap, it's not likely that we will see a significant improvement.

overload set:

```cpp
void Fun(std::vector<int>& byRef)
{
  std::cout << "byRef\n";
}

void Fun(const std::vector<int>& byConstRef)
{
  std::cout << "byConstRef\n";
}

void Fun(std::vector<int>&& byRvalueRef)
{
  std::cout << "byMoveRef\n";
}

void Use()
{
  std::vector v{2, 3, 4};
  const std::vector cv{5, 6, 7};
  Fun(v); //A We pass a lvalue
  Fun(cv); //B We pass a const lvalue
  Fun({3, 5, 6}); //C We pass a temporary
}

```

object lifetime extension, temporary object overloads.

_std::move_ and _std::static_cast<T&&>_ are equivelent. when we std::move we tell the compiler we won't use this object anymore.

> - std::move is only a cast. The overloads can move if they exist.
> - Temporary objects are picked up by default using move operations.
> - Only if we have an object we no longer want to use we can say std::move.
> - Move semantics is nothing else than an additional overload that is allowed and expected to steal data from a source object.

lvalue and rvalue categories, lvalue object have a memory address,we can give them a name. this defintion expanded into: lvalue, xvalue (expiring value), prvalue (pure rvalue), glvalue(generalized lvalue,lvalue or xvalue).

an example of real code is the std::string class, which utilizes move semantics to get better performance. we use the utility function of _std::exchange_.

a moved from object is an valid, but unknown state, C++ uses non-destructive moves. the object must be destructible and assignable.

> dealing with moved-from objects
>
> - Simple Rule: Never touch a moved­from object.
> - You know what you’re doing rule: You can reuse a moved­from object once you brought the
>   object back in a valid and known state. For all data types, assigning a new value to the moved from object is a safe operation.

an example of push_back and `noexcept` for the strong guarantee, using copy constructor until we explicitly mark the operations as memory safe.

perfect forwarding, _std::forward_ inside a template with && value. _std::forward_ behaves properly on both lvalue and rvalue objects.

we shouldn't move the return value, the compiler will optimize this into a copy-elision, which is better.

a table about what operations we get from the compiler when we write an object. we can also mark function as reference qualifiers, that means overloads for temporary objects.

</details>

### Back to Basics: C++ Smart Pointers - David Olsen

<details>
<summary>
Going over the basics of smart pointers
</summary>

[Back to Basics: C++ Smart Pointers](https://youtu.be/YokY6HzLkXs), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Olsen-Smart-Pointers-CppCon22.pdf)

example of bug

```cpp
void science(double* data, int N) {
  double* temp = new double[N*2];
  do_setup(data, temp, N);
  if (not needed(data, temp, N))
    return; // early return, memory leak
  calculate(data, temp, N);
  delete[] temp;
}

void science(double* x, int N) {
  double* y = new double[N];
  double* z = new double[N]; // if this throws, then we leak y
  calculate(x, y, z, N);
  delete[] z;
  delete[] y;
}

float* science(float* x, float* y, int N) {
  float* z = new float[N];
  saxpy(2.5, x, y, z, N);
  delete[] x; //are we allowed to delete
  delete[] y; // what if they were allocated with malloc?
  return z;  // caller is expected to delete
}
```

raw pointers have too many uses, they are too powerful:

can point to a single objects or array

- new, delete or new[], delete[]
- single object pointers can't use ++p,--p or p[n]
- array pointers can use ++p,--p or p[n]

owning vs non-owning

- owner must free the memory when done
- non-owner must never free the memory

nullable vs non-nullable

- some pointers can never be null
- it would be nice to get this from the type system

the type system doesn't help: `T*` can be used for all combinations of those characteristics.

smart pointers can help us with this. they behave like pointers, the point to an object and can be dereferenced. they also contain additional behavior: such as releasing memory, enforcing restrictions, and performing extra checks.

there are uses for raw pointers: such as non-owning pointers to single objects, (for arrays using span, the non-owning wrapper).

_std::unique_ptr_

- owns memory
- assumes it is the only owner
- automatically destroy the object and releases the memory
- move-only type
- the required template parameter is the element type.

```cpp
void calculate_more(HelperType&);
ResultType do_work(InputType inputs) {
  std::unique_ptr<HelperType> owner{new HelperType (inputs)};
  owner->calculate();
  calculate_more(*owner);
  return owner->important_result();
}
```

it can also be a member of a class, and the class destructor will take care of calling the destructor of the unique pointer. it's a move only type which cannot be copied, only moved. unique ownership cannot be copied.

unique pointers member functions don't throw exceptions, it doesn't allocate, it takes ownership of existing memory. we should use `std::make_unique` to create the unique pointer, but it can't be used everywhere. we can use unique pointers for array types, which would make the destructor use the `delete[]` operator and expose the indexing operator.

when we transfer ownership, we pass and return the unique pointer by value.

```cpp
std::unique_ptr<float[]> science(
  std::unique_ptr<float[]> x,
  std::unique_ptr<float[]> y,
  int N) {
  auto z = std::make_unique<float[]>(N);
  saxpy(2.5, x.get(), y.get(), z.get(), N);
  return z;
}
```

we need to be careful to never pass the same pointer to more than one unique pointer objects, this can create a double free crush. another problem is that when we delete a pointer, any other variables which point to the data are dangling and can cause undefined behavior.\
standard containers know how to handle move only types.

_std::shared_ptr_ has a shared ownership model (most often done with reference counting), and when the last owner is deleted, then the data itself is deleted. shared pointers can be copied. all shared pointers are treated equally.

shared pointers point both to the object and the control block, the control block counts how many objects point to the data. creating a shared pointer can throw exceptions, because of the control block which might be allocated. there is no way to release the ownership from any single shared pointer.\
there is a `std::make_shared` utility function, which has better memory footprint. if we want to share memory, then we need to create the first shared pointer from the memory, and the rest of them from the initial shared pointer, not from the raw pointer. if we don't, then we have a problem of double free.

thread safety:

> - Updating the same control block from different threads is thread safe (decrementing and incrementing the reference count is usually atomic)

```cpp
auto a = std::make_shared<int>(42);
std::thread t([](std::shared_ptr<int> b) {
  std::shared_ptr<int> c = b;
  work(*c);
  }, a);
{
  std::shared_ptr<int> d = a;
  a.reset((int*)nullptr);
  more_work(*d);
}
t.join();
```

> - Updating the managed object from different threads is not thread safe

there is no synchronization of the managed object.

```cpp
auto a = std::make_shared<int>(42);
std::thread t([](std::shared_ptr<int> b) {
  std::shared_ptr<int> c = b;
  *c = 100;
}, a);

{
  std::shared_ptr<int> d = a;
  a.reset((int*)nullptr);
  *d = 200;
}
t.join();
```

> Updating the same shared_ptr object from different threads is not thread safe

one threads the object while the other modifies it.

```cpp
auto a = std::make_shared<int>(42);
std::thread t([&]() {
  work(*a); // a is captured
});
a = std::make_shared<int>(100);
t.join();
```

the shared pointer added supports for arrays at c++17, and make_shared array support since c++20.
it is easier to switch from std::unique_ptr to std::shared_ptr

_std::weak_ptr_ is a non-owning reference to a shared_ptr managed object. it knows the lifetime of the object, to use it we lock the object and create a shared pointer.

```cpp
std::weak_ptr<int> w;
{
  auto s = std::make_shared<int>(42);
  w = s;
  std::shared_ptr<int> t = w.lock(); // create a shared pointer
  if (t) // check if data exists
    printf("%d\n", *t);
}
std::shared_ptr<int> u = w.lock(); // reference count is down to zero!
if (!u) printf("empty\n");
```

they are used for cacheing, to avoid dangling reference (instead of using raw pointers) and to avoid circular references.

for smart pointers, we can define a custom deleter, which is one of the template parameters, we define how the pointer deletes the resources (no way to do so with make_unique). for shared_ptr, the deleter is part of the constructor argument, the deleter is type erased, because the shared pointer allocates memory of its own.

</details>

##

[Main](README.md)
