<!--
ignore these words in spell check for this file
// cSpell:ignore ostringstream affordance O’Dwyer Stringize Eerd binarize iset Hyurum
 -->

[Main](README.md)

OOP

## Keynote: SOLID, Revisited - Tony Van Eerd

<details>
<summary>
Looking at the SOLID principles again, mostly on the SRP, avoiding massive classes with too many member functions.
</summary>

[SOLID, Revisited](https://youtu.be/glYq-dvgby4)

history of the "solid" acronym, started in the 2000's. divide between Object oriented programming and "Value oriented programming".

### Single Responsability

it's easy to tell that we got it wrong, but not as easy to know when we got it correctly. the mega class with member variables that are grouped, but still can't be separated because they still effect each other and use eachother, big classes make it easy to cross between 'sections' and break the separation. this is called **complecting** (from which the word 'complex' derives from)

> - "A component should have only one reason to change"
> - KYSS - keep your stuff separate.

holding data vs using data. having an entity class that keeps everything together, array of structs vs structs of arrays. it's probably better to pass data in the function call then to grab it inside the function somehow.

be suspicious of functions with no parameters and no return type. they are probably hiding a huge mess inside.

#### Classes

the example used is the image struct:

starting with this:

```cpp
struct Image
{
  unsigned char * pixels;
  Format format; // RGB, RGBA, etc..
  int width;
  int height;

  Image();
  Image(int width, int height, Format format);
  ~Image();
  Image(string filename);
  void save(string filename) const;
  //... copy operations, move operations

  void invert();
  void monchrome();
  void binarize(int threshold);
};
```

> "classes are made of velcro"

we add features onto classes, this isn't the best approach, we shouldn't have so many member function, it might be better to have free functions that manipulate, store/load (serialize).

> "Any image class with load/save is bad and should feel bad" \
> \-Tony Van Eerd

so, after extracting away the save/load methods and maniplators

```cpp
struct Image
{
  unsigned char * pixels;
  Format format; // RGB, RGBA, etc..
  int width;
  int height;

  Image();
  Image(int width, int height, Format format);
  ~Image();
  //... copy operations, move operations
};
  Image loadImage(string filename);
  void saveImage(Image const & img,string filename);
  //...

  void invertImage(Image & img,);
  void monchromeImage(Image & img,);
  void binarizeImage(Image & img,int threshold);
```

a class should hold what it needs, it should hold it's invariants, invariants are the responsabilities of the class. the class should only take care of those, it should not provide 'specific' behavior for some cause.

after the changes, the Image struct has three responsibilities

- pixels allocating/deallocating.
- pixels match the width and the height.
- pixels match the format.

the next step is to remove the resposability of deallocating data by using a smart pointer. this also takes care of the copy operations, move operations which we no longer need to specify. a class should strive to have the 'rule of zero'.

```cpp
struct Image
{
  std::unique_ptr<unsigned char> pixels;
  Format format; // RGB, RGBA, etc..
  int width;
  int height;

  Image();
  Image(int width, int height, Format format);
};
```

the next step is to redefine our responsibilities into "correctly allocate Pixels and Maintain that invariant".

we still have the large number of mutating functions that we use to manipulate the image, so we can't make everything private, so we must provide the minimal functionalities for that.

> "min functions to write the rest"

```cpp
class Image
{
  std::unique_ptr<unsigned char> pixels;
  Format format; // RGB, RGBA, etc..
  int width;
  int height;
  public:
  Image();
  Image(int width, int height, Format format);
  Pixel getPixel(int x, int y); //but what about the format?
  void setPixel(int x, int y, Pixel p);
};
```

this won't work because of the format, we can use a proxy class for pixel that know how to convert to each format, which will be very slow (same operation for every pixel). We can get around that with some templating.

```cpp
class Image
{
  std::unique_ptr<unsigned char> pixels;
  Format format; // RGB, RGBA, etc..
  int width;
  int height;
  public:
  Image();
  Image(int width, int height, Format format);
  Format getFormat() const;
  template<Format> Pixel<Format> * begin();
  template<Format> Pixel<Format> * end();
};
```

[How Non-Member Functions Improve Encapsulation, Scott Meyers (2000)](https://www.aristeia.com/Papers/CUJ_Feb_2000.pdf).

small set of functions, extrenal functions that use them

> $Types + Functions = Programming$

the important things about the class are the Name and the border. if you got the name right, the border (what belongs in the class and what doesn't) will be easy to define.

#### Functions

it's not just classes that should follow the SRP, functions should do so as well. if a function has more than one step, it's probably more than one responsability.

this is the example function (unnamed in the lecture)

```cpp
int someFunction()
{
  //...
  // step 1, turn 2D camera pts into 3D points
  std::vector<Point3D> cam3d[2];
  std::vector<Point2D> proj2d[2];

  auto &camData = getCamData();
  auto &projData = getProjData();

  for (int ptIdx = 0; ptIdx < camData.size(); ptIdx++)
  {
    auto camPoint = camData[ptIdx];
    Point3D p1;
    Point3D p2;
    m_camera->mapPointTo3D(campoint,p1,p2,calibrationType);
    auto iset = sg->intersect(p1,p2);
    for (auto & intersection : iset)
    {
      //...
    }
  }

  //...
  // step 2, solve via mapper
  //...
  // step 3,...
  //...
  // validate fit
}
```

we can start by turning step 1 into it's own function. and actually, make everythin into it's own function. now the function still does the same 4 step, but it's only responsible for calling them, each sub function is responsible for one step. the responsability of the wrapping function is to delegate the calls.

```cpp
int someFunction()
{
  //...
  // step 1, turn 2D camera pts into 3D points
  std::vector<Point3D> cam3d[2];
  std::vector<Point2D> proj2d[2];

  projectTo3D(cam3d,proj2d,getCamData(),getProjData());

  //...
  // step 2, solve via mapper
  Step2Function(cam3d,mapper);
  // step 3,...
  Step3Function();
  // validate fit
  ValidatingFunction(res);
}
```

another thing to search for are functions with the word 'And', using this regex `"[a-z]And[A-Z].*\("`,sometimes we find some basic types are missing (x,y is Point, width and height is Resultion, min and max is Bounds), and sometimes we find bundled function that are more efficient to call together (even if it's not required), sometimes we bundeles stuff together to avoid creating temporary objects. we might be able to divide those into functions that chain into one another.

"efficiency" version

```
for (item : container)
{
  get();
  or_maybe();
  calculate();
  res =something();

  hand();
  answer();
  to();
  something_else(res);
}
```

"KYSS" version (keep your stuff separate)

```
bunch = getAll();
setAll(bunch)
```

separated but still efficient

```
for (item: container)
{
  res = produce();
  consume(res);
}
```

which is basically _std::transform_.

> "Write the functions you wish to see in the world"

imagine the code the way you wish it was possible, and then create what's missing. "write more function".

another kind of "And" functions are convienet, "Do A and then Do B". which are sometimes forcing us to go through all the steps of A when we just wanted B, or vice versa.

we should also keep an eye for "Or" functions.\
some functions that take in flags simply use the flag to call additional behavior, this defintally could be done outside the function call.\
other functions are just an if-else block, which means that there are two separate function hidding. same for giant switch cases. we should extract away the inner function, the top level decides what to do, the low level controls 'how'.

### Open-Closed

> "A software artifact should be open for extension, but closed for modification"\
> \- Bertand Meyer

managing change, a base class shouldn't need to change because you want to add a derived class. sometimes we need to change the base class to give it the option to call the new features of the derived class, which means changing all the subclasses, this could be done by giving default behavior or by using reflection somehow.

> "An API is defined by the code that calls it, not by the code that implements it"

the **SOLID** is about managing change, not for the code when you first write it, it's for the code that you add and modify.

we should differentiate between intrinsic and extrinsic properties of our classes. rectangles have width and height, not x and y. Circles have radius, but not a x,y center location. the intrinsic property of a shape is the relation between it's edges (and the angles), not the locations on some plane. those are extrinsic properties. they don't belong inside the class.

what does OOP mean?\
Different ideas for different people, is it about hierarchy and inheritance? is it about encapsulation? is it about identity and how the object is passed and returned from method?

### liskov Substitution

> Subtype Requirement:\
> "Let $\phi(x)$ be a property provable about object x of type T, then $\phi(y)$ should be true for objects y of type S where S is a subtype of T"
>
> defining a "is-a" relation.\

the subtype should behave the same as the derived class, and can substitute for one.

however...

> Hyurum's law:\
> "With a sufficient number of users of an API, it does not matter what you promise in the contract:
> all observable behaviors of you system will be depended on by somebody"

```cpp
//not sure about the syntax
template <typename S>
concept Sequence = requires (S s)
{
  {s.begin()} -> Iterable;
  {s.end()} -> Iterable;
};
```

concepts and contracts.

the liskov substation principle can be a litmus test to determine whether a conversion should be implicit (there is an 'is-a' relation) or explicit (no, just some common traits).

### Interface Segregation

> - "No client should be forced to depend on methods it doesn't use"
> - "Don't give them what they don't need"
> - "Don't make complecting easy"

interfaces should be small, minimal, just what it needs, not exposing everything. and if we can just pass small concrete types, even better.

> "Code top-down on the way down, and bottom-up on the way back up"

when we write, we first pass in what we have, and then when we're done, we can see what we used and redefine our code to only pass those parts.

### Dependency Inversion

> "The most flexible systems are those in which source code dependencies refer only to abstractions, not to concretions"

how does this work with the earlier example of using 'small concrete types'?

it's okay to depend on vocabulary types,the important part is that we don't depend on stuff that can be changed or added.

</details>

## Polymorphism À La Carte - Eduardo Madrid & Phil Nas

<details>
<summary>
Thinking about polymorphism, can we move away from virtual dispatch and use duck typing?
</summary>

[Polymorphism À La Carte](https://youtu.be/tn8vlo14FT0), [slides](https://levelofindirection.com/refs/polymorphism.html)

> [À La Carte](https://www.merriam-webster.com/dictionary/%C3%A0%20la%20carte) - "according to a menu or list that prices items separately"

continuation of the 'OO considered harmful' and 'not leaving performance on the jump table' from cpp-con 2020 and 'type erasing the pains of runtime polymorphism" from c++ London

different ways of doing OO, virtual inheritance- method dispatch (C++) vs message passing (smalltalk, objective C). virtual dispatch at dynamic runtime is just another passage from the virtual table (vtable)of the calling object.

message passing is passing a message to an object, the object uses the selector + args to call the correct method, but there are some added benefits:

some demonstaration of objective c, apparently catch can also test objective c code, it has null-object pattern baked in, if we call anything on null, we get the null value, zero or nill. we can have 'fully typed' decleration and dynamically typed ("id") classes. we can then do runtime reflections and graft methods onto the class so we can get 'duck typing'. we decide that we want our class to fail either at compile time (the class doesn't contain the method) or at runtime (doesn't contian the method and nobody added this method to it).

> Runtime polymorphism
>
> - "We want _substitution_ as in Liskov's Substition Principle, this implies _subtyping_. Subtyping is not the same thing as _subclassing_
> - something that "quacks", if a dog can learn to quack, great. we don't care about all the formal relations.

using referential semantics cause performance costs and all sorts of other pain points.

> Beyond Inheritance + Virtual
>
> - Sean Parent's [Inheritance is the Base Class of Evil](https://youtu.be/bIhUE5uUFOA) lecture from 2013.
> - Referential semantics
>   - Indirections
>   - Allocation
>   - Incentive to share state
>   - Non-local reasoning
> - Intrusiveness
> - also: the subclassing relation is too strong:
>   - Supports only monophyletic relations, but para- and polyphyletic relation do exist as well
> - There are many ways to do message passing

monophyletic, paraphyletic, polyphyletic relations - terms from taxonomy

type erasure _Affordance_

[The Space of Design Choices for std::function](https://quuxplusone.github.io/blog/2019/03/27/design-space-for-std-function/) by Arthur O’Dwyer. runtime polymorphism using std::function and callable objects, picking and choosing how we want stuff to be designed, storage, ownership, const and non-const version.

_zoo::AnyContainer_ allows tools for affordance.

### Duck Typing

python has duck typing.

- if it waddles like a duck, swims like a duck and quacks like a duck, we shouldn't care that it isn't really a duck, we can use it as if it was simply a duck.
- and example of "Stringize" generic duck typing function, leaverage `operator<<` and `to_string`, having a bridge for the runtime for overload resolution and template specialization, doesn't require changing types.
- an ad-hoc approach.
- a _zoo::AnyContainer_ with some affordance can serve as a basis for another continaers with more affordance.
  - Typical example: A container that is move-only (does not require copyability) is extended with the affordance of copyability.
  - This extensibility applies to containers that themselves are extended for _zoo::AnyContainer_, such as _zoo::Function_

support composability.

```cpp
using NormalPolicy =
    zoo::Policy<
    void *[2], //size of two pointers
    zoo::Destroy, zoo::Move, zoo::RTTI>; //normal destructibility, move operations, RTTI behavior
using TypeErasureProvider = zoo::AnyContainer<NormalPolicy>;
using OrderConsumer = zoo::Function<TypeErasureProvider, CallSignature>; //template adapter

//adding copying capabilities.
using CopyableOC = zoo::AnyContainer<
        zoo::DerivedVTablePolicy<OrderConsumer, zoo::Copy>>;
```

[example](https://godbolt.org/z/xKsq3f) in compiler explorer

Faking and mocking using duck-typing in objective c. We simply declare a class with the required operations, we don't need the whole chain of interfaces or inheritances,exchaning and injection methods from singletons ("method swizzling").

### Message Passign in Rust and Swift

all sorts of important people in those langauges teams,

Protocol oriented programming in swift. abstractions, models, etc... (swift example of protocols, extensions) \
Also a Rust example with traits.

</details>

[Main](README.md)
