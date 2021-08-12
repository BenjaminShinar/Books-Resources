## Chapter 2 - auto

<summary>
The auto keyword, when to use it, when to avoid using it, what are the pitfalls.
</summary>

The auto keyword saves typing and cascading changes of types, but it can also cause problems. we know how auto deducts type, but if it's not what we wanted, we need to know why and how to deal with this.

### Item 5: Prefer auto to Explicit Type Declarations

<details>
<summary>
Using auto saves time, avoids problems and increases performance.
</summary>

In the past we could write *int x;* and we would be fine. except if it's not initialized, so now this depends on the context. is it zero, is it undefined, is it garbage? and sometimes we would need to get the value type from the iterator. and we might want a variable whose type is a clojure (a lambda).
```cpp
int x; // the value of int changes on the context where it's declared.
template<typename It>
void dwim(It b, It e) // dwim = "do wha I mean"
{
    while (b != e)
    {
        typename std::iterator_traits<It>::value_type currentValue = *b;
        // ...
    }
}
```
using auto is a way to avoid some of these issues. and the compiler can deduce types known only yo the compiler.
```cpp
int x; // the value of int changes on the context where it's declared.
auto x1; // error! initializer required!
auto x2 =0; // now, this is good.
template<typename It>
void dwim(It b, It e) // dwim = "do wha I mean"
{
    while (b != e)
    {
        auto currentValue = *b; // no need to write the whole thing
        // ...
    }
}

// c++11.
auto derefUPLess = []
    (const std::unique_ptr<Widget> & p1,const std::unique_ptr<Widget> & p2)
    {return *p1<*p2;}; // who knows what's the type of the lambda? do we care?

// c++14. now we don't need to define the parameters!
auto derefLess =[]
    (const auto & p1,const auto & p2)
    {return *p1 < *p2;};
```

std::function is a template in the standard library the generalizes the concept of function pointer, but it can refer to any callable object (anything that can be invoked with the operator()).

anything we this signature:  
*bool (const std::unique_ptr\<Widget\>&, const std::unique_ptr\<Widget\> &);*  
can be captured by a std::function from this:
*std::function\<bool(const std::unique_ptr\<Widget\>&, const std::unique_ptr\<Widget\> &)\> func;*

```cpp
std::function<bool(const std::unique_ptr<Widget>&, const std::unique_ptr<Widget> &)> derefUPLess(const std::unique_ptr<Widget> & p1,const std::unique_ptr<Widget> & p2) {return *p1< *p2;}; // without auto
```
std::function isn't the same as auto, auto has the same clojure type and the same size. while the std::function object has a fixed memory on the stack, and may request more memory from the heap to store the clojure. Additionally, invoking functions through std::function is usually slower, due to inlining rules. we can have a similar argument about std::bind.

here is another advantage of auto, avoiding "type shortcuts".
```cpp
std::vector<int> v;
unsigned sz = v.size(); //what is unsigned
```
the type of 'unsigned' can change depending on the platform, it's different between 32 and 64 bit systems. however

```cpp
std::vector<int> v;
auto sz = v.size(); //sz type is the correct type, std::vector<int>::size_type, whatever it may actual be
```
another use of auto. can we spot the problem?
```cpp
std::unordered_map<std::string, int> m;
for (const std::pair<std::string, int>&p:m)
{
    // ... do something with p
}
```
the problem is that type of the key in of std:::unordered_map is const, so the iterable type should be *const std::pair\<const std::pair\<const std::string, int\>\>*. so in the example above, the compiler will find a way to convert the real type to the requested type and create a temporary object. all these problems would disappear if we simply used auto.
```cpp
std::unordered_map<std::string, int> m;
for (const auto &p:m)
{
    // ... do something with p
}
```

in both examples, we declared a type and paid the price of implicit conversations. however, there are some downsides to using auto. one of the bigger issues is readability.

#### Things to Remember

> * auto variables must be initialized, are generally immune to type mismatches
that can lead to portability or efficiency problems, can ease the process of
refactoring, and typically require less typing than variables with explicitly
specified types.
> * auto-typed variables are subject to the pitfalls described in Items 2 and 6.

</details>

### Item 6: Use the Explicitly Typed Initializer Idiom when auto Deduces Undesired Types

<details>
<summary>
auto might fail us, but we shouldn't abandon it just yet
</summary>

auto usually works, except when it doesn't, take this code. we return a temporary vector of bool, take the value from index 5 and store it in a variable.
```cpp
std::vector<bool> features(const Widget& w);
Widget w;
bool highPriority = features(w)[5];
processWidget(w, highPriority);
```
but what if we used auto?
```cpp
std::vector<bool> features(const Widget& w);
Widget w;
auto highPriority = features(w)[5];
processWidget(w, highPriority);
```
now we get undefined behavior. highPriority is no longer bool, it's whatever type the operator[] of std::vector\<bool> returns, which happens to be a nested type std::vector\<bool>::reference. the return type is something that acts like bool, and is supposed to be able to fit wherever a bool is required. however, in this case, the object is a reference to a none existent container, which means we are dealing with a dangling pointer.  
this situation happens with vector of boos, because we can't return a single bit. (note: for this reason [vector of bool](https://en.cppreference.com/w/cpp/container/vector_bool) doesn't follow the rules of other sequence containers). we can see similar issues with std::bitset and the operator[].  
the std::vector\<bool>::reference is a *proxy* class, like many other class (smart pointers are proxies for pointers), only that this class is designed to be invisible to the user. these kinds of classes are called *expression templates*, and are supposed to be converted into the actual class when they become lvalue. 
``` cpp
class Matrix;
Matrix m1{},m2{};
Matrix m3 = m1 +m2;
```
auto doesn't play well with those invisible proxy classes, and the programmer is responsible to detecting these cases. those proxy classes are usually declared in the header files.  

#### Explicitly Typed Initializer Idiom

<details>
<summary>
We can suggest what type auto should be by casting.
</summary>

this problem doesn't mean we should abandon auto to avoid the potential problem, instead, we should employ the *explicitly typed initializer idiom*. this means, still using auto, but casting the initialization process into the correct type.
``` cpp
auto highPriority = static_cast<bool>(features(w)[5]); //force the reference into bool.
```
this also serves the process of emphasizing the intent (and should come with a comment explaining it), even when there is no proxy class involved. we might want to save space and use float instead of double (and we don't care about loss of precision), or maybe we want to get an index in some percentage of a container.
```cpp
double foo(); //function returns a double.
float f1 = foo(); // implicit conversion and loss of precision, are we doing this in purpose or did we just forget to change f1? the safe thing to do is to make it auto.
auto f2 = static_cast<float>(foo()); // we want f2 to be a float, and we make our intention clear, 
std::vector<int> v;
auto i1 = 0.5 * v.size(); // we want the middle element, but i1 is double!
int i2 = 0.5 * v.size(); // now i2 is integer, but this seems like a mistake.
auto i3 = static_cast<int>(0.5 * v.size()); // i3 is an int, and somebody thought it was very important that i3 be an int.
```
</details>

#### Things to Remember

> * “Invisible” proxy types can cause auto to deduce the “wrong” type for an initializing expression.
> * The explicitly typed initializer idiom forces auto to deduce the type you want it to have.

</details>
