<!--
ignore these words in spell check for this file
// cSpell:ignore Neumann Eigen metabench
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Templates & MetaProgramming

### Back to Basics: Templates in C++ - Nicolai Josuttis

<details>
<summary>
Basic Templates in C++.
</summary>

[Back to Basics: Templates in C++](https://youtu.be/HqsEHG0QJXU), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CTemplates_cppcon_220918.pdf).

C++ started as "C with classes", with the intent to being in object oriented programming into the language.

templates are generic code for arbitrary types/values. code then gets instantiated to specific types at compile time. We have function templates, functions that operate on any type.

#### Function Templates

```cpp
template <typename T>
T myMax(T a, T a)
{
  return b < a ? a: b
}
```

templates become real code when they are used. this can lead to code bloat. we can be explicit about the template parameter

```cpp
s = myMax<std::string>("hello", "world");
```

we can also have generic iterating, in a range for loop.

```cpp
template <typename T>
void print(const T& coll)
{
  for (const auto& elem: coll){
    std::cout << elem << '\n';
  }
}
```

template must be **defined** (not only declared) in header files. there's no need to _inline_. things might change when modules become commonplace.

in c++20, _auto_ was added for ordinary function parameters, it's an abbreviated form of writing a templated function.

```cpp
void print(const auto& coll)
{
  for (const auto& elem: coll){
    std::cout << elem << '\n';
  }
}
```

a template has implicit requirements, these are the actions taken on the parameters in the body of the function, some types can't be copied into the function.\
if we don't provide an explicit template parameter, we can get errors at complication. templates errors tend to look monstrous and messy.

concepts were added to the language at c++20, they allow explicit definiton of the requirements for the template parameters.

```cpp
template <typename T>
requires std::copyable<T> && SupportsLessThan<T>
T myMax(T a, T a)
{
  return b < a ? a: b
}
```

templates can take more than one template parameters.

```cpp
template <typename T1, typename T2>
void print (const T&1 v1, const T&2 v2){
  std::cout << v1 << "&"<< v2 << '\n';
}
```

if we have a return type, it might not be clear what happens when there is more than have template type parameter. we can just _auto_ to deduce the compiler type. (it was previously done with some `decltype`)

```cpp
template <typename T1, template T2>
auto myMax(T1 a, T2 a)
{
  return b < a ? a: b
}
```

#### Class Templates

we can write template classes, this is done for containers, so we can have a vector of integers, of double and of strings. when we use class templates, we usually need to specify the type at the variable defintion.\
type in used class defintions must support the parts of the template which are **used** in the code. so if we don't use some member function, the instantiation of it won't be created, so there's a possible bug hidden there.

Class template argument deduction - when the constructor can deduce the type, we don't need to explicitly define the type.

```cpp
std::vector<int>{1,2,3}; // explicit class parameter
std::vector{1,2,3}; // argument deduction
```

most containers have several constructors, so there might be several matches, and there are rules for which matches takes priority over others. using curly braces `{}` can cause weird behavior with class template argument deduction.

std::array is c templated aggregate type. it's a wrapper of C code.

```cpp
template <typename T, size_t SZ>
struct array {
  //...
};
```

the template parameters don't have to be types.

Supported types:

- Types for constant integral values (int, long, enum, ...)
- std::nullptr_t (the type of nullptr)
- Pointers to globally visible objects/functions/members
- Lvalue references to objects or functions
- Floating-point types (float, double, ...) (since c++20)
- Data structures with public members (since c++20)
- Lambdas (since c++20)

Not supported are:

- String literals (directly)
- Classes
- Since C++20 supported are:

#### Variadic templates

One parameter representing multiple arguments. we need a base case of no parameters

```cpp

void print()
{
  // empty case
}
template <typename T, typename... Types>
void print(T arg1, Types... args)
{
  std::cout <<arg1 <<'\n';
  print(args...);
}
```

all the code in the instantiated function must be valid, so if we wish to skip the base empty case, we need a compile time check, not a runtime check, so in c++17 we can use `if constexpr`.

```cpp
template <typename T, typename... Types>
void print(T arg1, Types... args)
{
  std::cout <<arg1 <<'\n';
  if constexpr(sizeof... args) > 0 {
    print(args...); // no need to have the base case function
  }
}
```

concepts can be used together with auto, which can help the compiler reduce ambiguity when determining function overload resolution.

```cpp
template <typename Coll>
concept HasPushBack = requires(Coll c, Coll::value_type v){
  c.push_back(v);
}

void add(HasPushBack auto& coll, const auto& val){
  coll.push_back(val);
}
void add( auto& coll, const auto& val){
  coll.insert(val);
}
```

we could also use this with compile time expressions.

```cpp
void add( auto& coll, const auto& val){
  if constexpr(requires { coll.push_back(val)})
  {
    coll.push_back(val);
  }
  else {
    coll.insert(val);
  }
}
```

</details>

### Help! My Codebase has 5 JSON Libraries - How Generic Programming Rescued Me - Christopher McArthur

<details>
<summary>
Using Type Traits
</summary>

[Help! My Codebase has 5 JSON Libraries - How Generic Programming Rescued Me](https://youtu.be/Oq4NW5idmiI), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/CppCon-2022-How-Generic-Programming-came-to-the-rescue.pdf)

> Focus of the talk - “**implementing traits with functions**”.\
> Explanation of template MetaProgramming implementation to abstraction JSON libraries.
>
> - Detecting if a function or method are implement for a type
> - Checking if an ADL implementation exists
> - Compile time requirements and SFINAE

once they added a package manager, it was easy to get more packages, and different libraries which do the same things are added to the stack.

> “Why don’t you template out the logic and metaProgram a traits implementation?”

```cpp
template<typename json_traits>
class basic_claim {
  static_assert(details::is_valid_traits<json_traits>::value,
  "traits must satisfy requirements");

  static_assert(
  details::is_valid_json_types<typename json_traits::value_type,
  typename json_traits::string_type,
  typename json_traits::integer_type,
  typename json_traits::object_type,
  typename json_traits::array_type>::value,
  "must satisfy json container requirements");
}
```

JWT -JSON Web Token, some predefined key. we need to both **create** and **verify** tokens.

the values can only be a limited set of types.

#### Verifying Claims

> Q: How can we check a type implements a function?

combine: using `std::experimental::is_detected`, which employs SFINAE to detect a named entity, but we need to add `std::is_function`, and eventually, we want to create a `is_signature` trait.

#### Creating Tokens

decltype doesn't work easily with overload functions, we need to resolve the function at compile time. more assertion errors over member functions.

variance with methods (array `.at` vs `[]`).

> Review
>
> - Check static function signatures with `is_detected`, `is_function`, and
>   `is_same`
> - We can resolve overloaded functions with the help of `declval`
> - To overcome `declval`’s lack of substitution we can add template helpers to return `true_type` of `false_type` is it does not resolve.
> - More indirection is usually the answer with SFINAE

</details>

### High Speed Query Execution with Accelerators and C++ - Alex Dathskovsky

<details>
<summary>
High speed Queries using CPU and GPU, in the way of compositing schemas and operators.
</summary>

[High Speed Query Execution with Accelerators and C++](https://youtu.be/V_5p5rukBlQ), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/High-Speed-query-execution-with-accelerators-and-CPP.pdf)

APU - Analytic Processing Unit.\
MFC - Multi Flow Core.

**Why do we need Accelerator for Analytics?**

big data gets bigger, so queries need to process more data, take more time, and become more energy consuming and more expensive. if we can speed up the query, we can get better results

#### Architectural Talk :

**History of computer architecture**

Von Neumann architecture, arithmetical and logic units, memory, input, output, registers, control units.

**RISC** - reduced instruction set computing, fixed length instruction, single operation instructions.\
**CISC** - complex instruction set computing - variable length instructions, instructions can be more complex.

then cpus got pipelines, we can run some instructions in a row, (some may stall), then we got branch prediction, if we got the prediction wrong then we lose cycle. another optimization was running instructions out of order, some instructions can be re-ordered to avoid stalling.\
Processing more data by each instruction, SIMD (single instruction, multiple data) and vector machines.

GPU can also help, they help with non-dependant (parallel) data, SIMT (single instruction multiple threads), good for graphics and games. it's still a von neumann machine, and it needs a CPU as a host. the CPU and GPU communicate through memory.

> Dataflow Architectures
>
> - No Program Counter
> - Execution occurs when data is ready
> - Can be non-deterministic

a program is a graph, instructions are nodes, data is tokens. it's a theoretical concept, which isn't used directly, but inspired MFC as a hybrid solution.

**How does APU(MFC) fit in**

APU - Analytic Processing Unit.\
MFC - Multi Flow Core.

> MFC
>
> - Performs very complex operations in one instruction
> - Handle large amounts of data
> - Vector Instructions (threads)
> - Support divergence between flow

each instruction is 64 bit (like RISC), but there are complex instruction (like CISC)

> MFC - Program
>
> - Each Basic Block of the Program is a DFG (Handling Multiple Data)
> - At The end of Each DFG there's a jump to a new DFG
> - All flows start at DFG1
> - Program Ends when all flows are at Terminator DFG

DFG - Diagonal Flow Graph.

DFG is a basic block is a compiled thing without branches, it's a deterministic set, between each DFG blocks there is a jump, which is a way to move from one block to another. this is how we do conditions and loops.

CGRA (Coarse-Grained Reconfigurable Arrays).

> MFC - CGRA
>
> - Mesh of nodes
> - Nodes connected with static switches
> - Every Node has its own operation
> - Tokens Correspond to flow
> - Tokens pass Through execution Nodes and end in TERM
> - TERM node determines next DFG for each flow
> - 3 different state storages:
>   - LVC
>   - Memory
>   - Streaming Buffer

instead of registers they use LVC, for passing data between DFGS.

LSU - Load Store Unit

#### How to Code It?

> DOGE - Dataflow Oriented Graph Execution
>
> - A low-level interface to the APU
> - Similar to CUDA
> - Gives the user to ability to access low level instruction and device built-ins
> - Clang frontend and LLVM backend

a toy example of code, which looks like a CUDA. it still uses macros.

```cpp
KERNEL_DEFINE(ex1_single_dfg, long* d_out0){
  int tid = get_flow_id();
  d_out0[tid]= 14 * tid +11;
  return;
}
```

and actual host code. still work in progress

```cpp
long* d_out0 = apu::memory_mgr_DeviceMemMgr::GetInstance().Allocate(NUM_OF_ELEMS * sizeof(long));
long* h_out = new long[NUM_OF_ELEMS]; // output array
for (int = 0; i < NUM_OF_ELEMS; i++){
  h_out0[i]= EMPTY_KEY_64;
}
apu::memory_mgr_DeviceMemMgr::GetInstance().CopyToDevice(h_out0, d_out0, NUM_OF_ELEMS * sizeof(long));

// Do the thing
KERNEL_LAUNCH(ex1_single_dfg, BLOCK_SIZE, d_out0);

// Copy the GPU memory back to the CPU here
apu::memory_mgr_DeviceMemMgr::GetInstance().CopyFromDevice(h_out0, d_out0, NUM_OF_ELEMS * sizeof(long));
fmt::print("h_out[{}] = {}\n",1, h_out[1]);
```

the DFG is an assembly code block,

```MIPS
# DFG_ID,1,entry:
  TCREATE1 = TCREATE
  REG0 = GETFID.64D
  REG1 = IMM.U $14
  REG2 = MADD.64D.U REG0, REGG1 $11 #multiple and add
  REG5 = ADDS.64D.U REG0, $3, &11 #add and shift
  REG6 = STP.PANY.L.64M.U REG5, REG2, TCREATE1
  REG7 = THALT REG6 #halt, like return
# End-Function
```

and because this DFG is small, multiple instances of this can be replicated on a single CGRA, allowing for more parallelism.

**Process explanation and code examples**

but this is hard to write, so there is a spark plugin called PIDGIN to create the code. the plug in will create DOGE code from spark sql queries.

Code example for simple query:\
a simple table of ID (numeric) and Plant (string), and a simple query (`SELECT * FROM TBL WHERE planet=='Mars';`). spark creates a plan, which PIDGIN uses to create a c++ code (using MetaProgramming and templates) with many types, and eventually a KERNEL_DEFINE function (with composition).

```cpp
KERNEL_DEFINE(kernel0 size_t num_rows, uint8_t* arg0,uint8_t* arg1, uint8_t* arg2) {
  CompositionType20 kernel0_plan(arg0, arg1, arg2);
  auto thread_id = git_flow_id();
  auto stride_size = get_flow_block();
  for (size_t row_idx = thread_id; row_idx <num_rows; row_idx +=stride_size) {
    kernel0_plan.Process(row_idx);
  }
}
```

Code example for complex query and scale:\
the same principles can be used for a more complex SQL query, this time using more kernel functions, each kernel is a simple code that creates a table from the input, and then feeds it forward to the next kernel.

</details>

### Taking Static Type-Safety to the Next Level - Physical Units for Matrices - Daniel Withopf

<details>
<summary>
A library for working with type safe matrices and doing type safe operations in terms of matrices
</summary>

[Taking Static Type-Safety to the Next Level - Physical Units for Matrices](https://youtu.be/aF3samjRzD4), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Taking-static-type-safety-to-the-next-level-physical-units-for-matrices-Daniel-Withopf-CppCon-2022.pdf), [CppCon 2020 lecture: A physical units library for the next C++](https://www.youtube.com/watch?v=7dExYGSOJzo).,[github](https://github.com/mpusz/units).

> "If it compiles, it works!"

Physical units in vectors and matrices.

> Design principle
>
> - Anticipate problematic usages and prevent them at compile-time
> - Built-in static analyzer
> - Enable users to make sweeping changes that work once they compile again
> - Make user code as expressive as possible

making phyiscal units strong types, different types for vectors and matrixes, combing linear algebra with physical units.

```cpp
si::length<si::metre> u = 3 * m; // units library
fs_vector<double, 3> v = { 1, 2, 3 };  // linear algebra library
```

combining the two libraries.

```cpp
fs_vector<si::length<si::metre>, 3> v = { 1 * m, 2 * m, 3 * m };
fs_vector<si::length<si::metre>, 3> u = { 3 * m, 2 * m, 1 * m };
std::cout << "v + u = " << v + u << "\n"; // [4m, 4m, 4m]
```

but we have a vector to hold different strong types,

if we tried using an std::variant, we would get an explosion of types which we would need to code against
4 by 4 matrix of possible types.

> Can you tell what this line of code does?

```cpp
vector_a[2] = vector_b[3];
```

> - What do the vector entries describe?
> - Is this an out-of-bounds access?
>
> 2nd try:

```cpp
vector_a[VELOCITY_X] = vector_b[POSITION_X];
```

> - Have the right index constants for the vector type been used?
> - Is assigning a position to a velocity really intended?
>
> 3rd try:

```cpp
vector_a[VELOCITY_X] = vector_b[VELOCITY_X];
```

> - In which coordinate frame are vector_a and vector_b?

#### type safe matrices

so what we would like in our ideal linear algebra library:

- check against out-of-bounds access (compile time)
- enforced expressive names for entries
- support for non-uniform phyiscal units
- physical units check for all matrix operations (compile time)
- coordinate frame annotations

we can started with _Named index structs_

the naming conventions is:

1. phyiscal quantity: Distance, Velocity, Acceleration.
2. Axis: X, Y, Z
3. Coordinate frame

```cpp
struct DX_C : CartesianIdxType<VehicleFrontAxleCoords,
tsm::CartesianXAis, si::Metre>
{};

struct DY_SENSOR_C : CartesianIdxType<SensorCoords,
tsm::CartesianYAxis, si::Metre>
{};

struct VehicleFrontAxleCoords : public CoordinateSystem<si::Metre> {
using Moving = std::false_type; // is the frame moving wrt to the “fixed” earth frame?
};
```

the most commonly used types:

```cpp
template<class Scalar, class RowIdxList, class ColIdxList, class MatrixTag>
class TypeSafeMatrix {
... // methods
private:
Eigen::Matrix<Scalar, SizeOf<RowIdxList>::value, SizeOf<ColIdxList>::value> m_matrix;
};

// vector is a matrix with column indexes
template<class Scalar, class RowIdxList, class MatrixTag>
using TypeSafeVector = TypeSafeMatrix<Scalar,
RowIdxList, tsm::TypeList<tsm::NoIdxType>,
MatrixTag>;

// position
template<class Scalar>
using PosVec3InVehicleFrame<Scalar> =
tsm::TypeSafeVector<Scalar, tsm::TypeList<DX_C, DY_C, DZ_C>, tsm::VectorTag>;
```

we have different ways of populating the matrix, some are more type safe and and some are less. there all sorts of complex ways to create matrices. the index types detect when we try to assign incompatible types and dimensions.

**Matrices taxonomy**:
jacobian matrices, type safe matrix, row and column exponents

> - Matrices types
>   - Covariance
>   - Jacobian
>   - Information
> - Vector types
>   - Position
>   - Displacement
>   - Information
> - Vector collection types types
>   - Position
>   - Displacement
>   - Information

trying all sorts of operations on different kinds of matrices.

> "unit safety is not enough, we need index-type safety"

matrix multiplication, going over the code for the plus operator, assertions (library uses c+11, so SFINAE and not c++20 concepts)

#### The Big Picture

looking at the problem dimension space and seeing what the library solves as opposed to just using units as value types.

(explaining common c++ bugs when using the library)

</details>

### C++ for Enterprise Applications - Vincent Lextrait

<details>
<summary>
describing a enterprise database object with C++
</summary>

[C++ for Enterprise Applications](https://youtu.be/4jrhF1pwJ_o),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/C-for-Enterprise-Applications.pdf)

C++ is used in most of backend processes, but only a few (the largest) enterprises use C++, while small and medium companies avoid c++. this is a paradox.\
C++ doesn't have good community toolkits, so people believe c++ should only be used for low level code, and that the language is lacking in abstractions.

enterprise applications need to work with large data (from databases), need fast response times and be reliable. those are constrained environments.\
enterprises work with loads of data types and relationships, and use less algorithms. this leads to **Data Integrity** issues. we need better abstractions, but so far, all attempts have failed.

mission statement

> - std::unique_ptr and std::shared_ptr as is are insufficient to describe relationships
> - Boost is very useful but still falls short
> - No Application Server available
> - SQL makes our C++ code fragile: it relies upon a cross-cut of the type system
> - Zero-cost abstractions pave the way to layered abstractions (aka Hyper-Automation/Hyper-abstraction), let's use them
> - Let's learn from the past and go back to basics

attemps at diagram based engineering, build software from the specifications.

- flow charts
- uml
- low-code/no-code

we can start working from the other end, rather than start with specifications, maybe we can start the abstraction from the lower level - _Data Modeling_.

> "[...] ontology encompasses a representation,
> formal naming, and definition of the **categories**, **properties**, and **relations** between the concepts, data, and entities that substantiate one, many, or all domains of discourse.\
> More simply, an ontology is a way of showing the properties of a subject area and how they are related, by defining a set of concepts and categories that represent the subject."

hyper abstraction: C++ code that is concise and specific enough to be compiled into a working program. every layer of abstraction moves us away from the raw code, and gives us better productivity, reduces bloat code, makes testing less required. we want a solution that makes expanding the code easier, and doesn't create a complexity wall that makes changes impossible.

They created an application web-server, and out of the box abstractions, including documents that persists in databases.

The document model: a root object, which can either own other objects, or be linked to other objects, (not necessarily other roots) which can be in the same collection or other collections.

#### Examples

**Referential Integrity Rules**

> - An object can own another one
> - An object can be owned at most by one other object
> - Root types cannot be owned (they have a UUID and are referenced in dictionaries)
> - A document is the set of objects owned transitively by a root object
> - An object can own another one only if the first is in a document
> - An object can possess a link to another only if the second is in a document
> - An object can be the target of as many links as needed
> - Links cycles are authorized
> - The system automatically maintains these rules when an ownership is broken

```cpp
struct characteristics;
struct category;
struct product: public root<>
{
  HX2A_ROOT(product, "prod", 1, root); // MACRO magic
  product(reserved_t): root(reserved), _chars(*this), _category(*this) {}

  own<characteristics, "ch"> _chars;
  strong_link<category, "ct"> _category;
};
```

**Objects automatic destruction as a “free” byproduct**

better than garbage collection, better than std::shared_pointer. real time guarantees.

**Replacing SQL with C++ functions**

> - We already have a language: C++
> - Object-oriented (can resolve SQL's cross cut maintenance headache issue)
> - Most powerful MetaProgramming model in existence
> - Compiles into lightweight binary code
> - In the "domain of discourse", types and owns/links form the syntax, let's add semantics

```cpp
class contract: public root<> {
// …
  slot<time_t, "s"> _start;
  slot<uint32_t, "d"> _duration;
  time_t end() {
  return _start + _duration;
  }
  attribute<time_t, &contract::end, "e"> _end;
};
```

</details>

### From C++ Templates to C++ Concepts - MetaProgramming: an Amazing Journey - Alex Dathskovsky

<details>
<summary>
how the language evolved since c++11 to support "concepts"
</summary>

[From C++ Templates to C++ Concepts - MetaProgramming: an Amazing Journey](https://youtu.be/_doRiQS4GS8), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/From-Templates-to-Concepts.pdf)

most people have bad experiences with templates, they don't like to write them, and hate to debug them.

#### Basic Template Rules

> **FUNCTION TEMPLATES**
>
> - Function templates can be fully specialized or overloaded.
> - Partial specialization is not allowed
> - but we can overload

```cpp
template<typename T> // base template
void print(T t){fmt::print("T: {}\n",t);};

template<> // fully specialized
void print(int t){fmt::print("T:Int=={}\n",t);};

template<typename T> // partial specialization is not allowed
void print<T*>(T* p){fmt::print("T*: {}\n",*p);};

template<typename T> // overload is allowed
void print(T* p){fmt::print("T*: {}\n",*p);};

int main(){
  print("1");
  print(1);
  int i = 1;
  print(&i);
}
```

overloading and specializations are tricky.

```cpp
template<typename T>
void print(T){fmt::print("Generic");};
template<typename T>
void print(T*){fmt::print("Overload");};
template<>
void print(double*){fmt::print("Specialization");};

int main(){
  double d=1.5;
  print(&d); // "Specialization"
}
```

but if we change the order, we get a difference result. an overload resolution considers only the base template.

```cpp
template<typename T>
void print(T){fmt::print("Generic");};
template<>
void print(double*){fmt::print("Specialization");};
template<typename T>
void print(T*){fmt::print("Overload");};

int main(){
  double d=1.5;
  print(&d); // "Overload"
}
```

therefore, we should prefer to use overloading and not specializations.

> Class Templates
>
> - support partial and complete specialization

in c++17 there are deduction guidelines

```cpp
template<typename T>
struct more{};
template<typename T>
struct more<T*>{};
template<>
struct more<int>{};

int main(){
  auto m1 = more<double>();
  auto m2 = more<double*>();
  auto m3 = more<int>();
}
```

#### Traits

in c++11 we got the standard type trait library

> Example of useful traits:
>
> - `is_pointer<T>`
> - `is_abstract<T>`
> - `is_assignable<T>`
> - `is_convertible<T, U>`
> - `is_same<T, U>`

```cpp
template<typename T>
void print(T const & t){
  fmt::print("{}",t);
}
int main(){
  int i{1};
  print(&i); // error can't fmt::print pointer (except void pointer)
}
```

we can use traits to get around this, we create specialize class that only works for pointers, and the function checks tht type trait and based on it chooses the correct specialization.

```cpp
template <typename T,bool>
struct printHelper{
  static void print(T const & t){fmt::print("{}",t);};
};

template <typename T>
struct printHelper<T,true>{
  static void print(T const & t){fmt::print("{}",*t);};
};

template<typename T>
void print(T const & t){
 printHelper<T,std::is_pointer<T>::value>::print(t);
}
```

**in this example we ignore the case of a void pointer (which we can't dereference)**

in c++14 some traits got alias for the inner type with the "\_t" suffix, and c++17 some traits got the value with the "\_v" suffix. these are aliasing to make things easier

```cpp
template <typename T>
using add_pointer_t = typename add_pointer<T>::type;
template <typename T>
constexpr bool is_pointer_v = typename is_pointer<T>::value;
```

another way to make thing cleaner is to use tag dispatch
("std::true_type" and "std::false_type").

```cpp

template <typename T>
void print(std::false_type, T const & t){fmt::print("{}",t);};

template <typename T>
void print(std::true_type, T const & t){fmt::print("{}",*t);};

template<typename T>
void print(T const & t){
 printHelper(typename std::is_pointer<T>::type{},t);
}
```

moving forward, in c++17 we can get even more simplified by using `if constexpr` to decide the branch in compile time.

```cpp
template <typename T>
void print(T const & t){
  if constexpr (std::is_pointer_v<T>)
  {
    fmt::print("{}",*t);
  }else{
    fmt::print("{}",t);
  }
}
```

in c++20, we can use `auto` for a simple concept, which makes things similar to tag-dispatch, but much easier.

```cpp
void print(auto & t){
  fmt::print("{}",t);
}

void print(auto* t){
  fmt::print("{}",*t);
}
```

however, we lose the const part, so we will create a pointer concept (later),

```cpp
void print(const auto & t){
  fmt::print("{}",t);
}

void print(const pointer auto& t){
  fmt::print("{}",*t);
}
```

#### SFINEA

lets start by writing a concept that detects a container. all the STL containers have a `::iterator` type, so lets try that.

```cpp
template <typename T>
struct is_container
{
  static const bool value = ???;
};
```

we need to learn about SFINEA

> SFINAE - SUBSTITUTION FAILURE IS NOT AN ERROR
>
> Special rule for function template overload resolution:\
> If an overload candidate would cause a compilation
> error during type substitution, it is silently removed from the overload set.

function with ellipses (...) for variadic argument are always inferior in overload resolution. so we try using that in a naive implementation, and we need `std::enable_if`, which is how we force SFINAE and make the compiler choose the correct overload.

```cpp
template <typename T>
void print(const auto& t, std::enable_if_t<is_container_v<T>,void*> =nullptr){
  for (auto && e: t){
    fmt::print("{}",e);
  }
}
```

the next advancement was variadic templates, being able to pass any number of template parameters, and we combine this with pack and unfolding

```cpp
template <typename... T>
struct are_all_integral:
  public std::conjunction<std::is_integral<T>...>{};

template <typename... T>
vod check(T... vals){
  static_assert(are_all_integral<T...>::value, "All vals must be integral");
}
```

> VOID_T
>
> - An extremely simple alias template that helps verify well-formed-ness.
> - Can be used for arbitrary member/trait detection
> - `void_t<T>` is well formed void only if T is well-formed, just like `enable_if<b, T>::type`

#### Concepts

now we start creating a concept for "is*container", we use `declval`, which is only good for compile time expressions. we create some struct that must be well formed for SFINAE, we use some built-in tools. but until we use actual \_concepts*, it's a very hard process, and the error messages are long and horrible.

in c++20. we got actual concepts. they make code cleaner, readable, and easy to maintain. we declare it with the `concept` keyword and `requires` clause. there are different forms of writing this.

</details>

### The Dark Corner of STL in Cpp: MinMax Algorithms - Simon Toth

<details>
<summary>
Some edge cases and problems with the min max algorithms.
</summary>

[The Dark Corner of STL in Cpp: MinMax Algorithms](https://youtu.be/jBeTvNgW25M), [slides](https://github.com/HappyCerberus/cppcon22-talk).

Free book about standard C++ algorithms available on github.

why are the min/max algorithms so hard? aren't they simple?

```cpp
auto min = std::min(1,2); // 1
auto max = std::max(1,2); // 2
auto clamped = std::clamp(0,1,2); //1, value, min, max
auto minmax = std::minmax(1,2);
```

but if we look at the templates, we see that minmax returns a pair.

```cpp
template<class T>
const T& min(const T& a, const T& b);

template<class T>
const T& max(const T& a, const T& b);

template<class T>
const T& clamp(const T& v,const T& lo, const T& hi);

template<class T>
std::pair<const T&, const T&> min(const T& a, const T& b);
```

so if we write the code we get references to temporary elements, auto type deduction doesn't deduce reference type.

```cpp
std::pair<const int&, const int&> minMax =std::minmax(1,2);
const int& min = std::min(1,2); // min is a dangling reference

auto [x,y] = std::minmax(1,2); // still dangling
std::pair<int,int> a = std::minmax(1,2); // this is ok.
```

to find this behavior, we need to run an address sanitizer.

there are some variants of the min max algorithms, the c++20 range versions behave the same.

c++14 has variant that take an initializer list, which return by value and not by reference.

```cpp
auto x = std::min({1,2});
// ok, decltype(x) => int

auto pair = std::minmax({1,2});
// ok, decltype(pair) => std::pair<int, int>

const int &z = std::max({1,2});
// ok, lifetime extension
```

but there are problem, because it's impossible to move from an initializer list. so it fails for move only types, and can incur heavy costs for multiple copies.

```cpp
auto x = std::min({MoveOnly{}, MoveOnly{}});
// wouldn't compile

ExpensiveToCopy a,b;
auto y = std::min({a,b});
// 3 copies

auto z = std::min(ExpensiveToCopy{},ExpensiveToCopy{});
// 1 copy since c++17, copy-initialization from prvalue
```

next is the problem of **const correctness**. we have this code

```cpp
MyType a,b;
if (b<a){
  b.do_something();
} else {
  a.do_something();
}
```

but we want it to be more simple, like this:

```cpp
MyType a,b;
std::min(a,b).do_something();
```

this will only work if the method we call is const, because the return value of the algorithm is a const reference. we could use `const_cast<>`, but that isn't very readable, and we might have undefined behavior if the method mutates state.

we would like to fix this:

1. remove the need for `const_cast`
2. remove the potential for dangling reference
3. avoid excessive copies

we can fix this by adding more overloads for the algorithms. then we get things better by using `auto` templates. there's something about the ternary operator here, so we need to call `std::common_reference_t` if we don't use it. we add _requires_ clause.

</details>

### What Can Compiler Benchmarks Reveal About Meta-Programming Implementation Strategies - Vincent Reverdy

<details>
<summary>
Creating a benchmark for compile time meta programming, using meta programming
</summary>

[What Can Compiler Benchmarks Reveal About Meta-Programming Implementation Strategies](https://youtu.be/9bSG1aHXU60), [github repository](https://github.com/vreverdy/metabench).

type based and function based programming - the example of adding tuples together. in large projects, we would need to decide which implementation is better (focusing on compile-time)

```cpp
// type based
template <class, class>
struct tuple_sum;

template <class... Ts, class Us>
struct tuple_sum<std::tuple<Ts...>, std::tuple<Us...>>{
  using type = std::tuple<Ts..., Us...>;
};

// function based
template <class... Ts, class Us>
std::tuple<Ts..., Us...> tuple_sum(std::tuple<Ts...>, std::tuple<Us...>){
  return {};
}

template <class Tuple1, class Tuple1>
using tuple_sum_t = decltype(tuple_sum(Tuple1{} ,Tuple2{}));
```

however, benchmarking is usually done for runtime, we don't have frameworks for benchmarking the work done in compile-time.

- DSL: Domain specific languages
- EDSL: embedded domain specific languages

the goal is a standalone meta-library to benchmark meta-programs and compilers.

```cpp
// Preamble
#include <tuple>
#include <type_traits>

// build a tuple from a repeated sequence of types: deceleration
template<class Sequence, class... Args>
struct type_sequence_maker;

// build a tuple from a repeated sequence of types: defintion
template<std::size_t... I, class... Args>
struct type_sequence_maker<std::index_sequence<I...>, Args...>{
  using type = std::tuple<
    std::tuple_element_t<I % sizeof...(Args), std::tuple<Args...>>...;>
};

// build a tuple from a repeated sequence of types
template<std::size_t N, class... Args>
using make_type_sequence = typename type_sequence_maker<
std::make_index_sequence<N>, Args...>::type;

// print types for debugging
template<class... Types, class... Args>
void metaPrint(Args...){
  std::cout<< __PRETTY_FUNCTION__ << std::endl;
}

// example of usage
int main(int argc, char* argv[]){
  using type = make_type_sequence<7, bool, char, int, double>;
  type value = {};
  metaPrint(x); // or metaPrint<type>()
  return 0;
}
```

in the example above, we will see this output:

```
void metaPrint(Args...)
[with Types={}, Args={std::tuple<bool, char, int, double, bool, char, int>}]
```

and we can increase N from 7 to something large and see when the compiler crashes because of the template depth...

so we can't use this approach directly to make an arbitrarily hard task for the compiler. and the types aren't unique, they are just repeating

we can move from type sequences to type trees:

```cpp
std::tuple<
  std::tuple<
    std::tuple<Args1...>,
    std::tuple<Args2...>
  >,
  std::tuple<
    std::tuple<Args3...>,
    std::tuple<Args4...>
  >,
  std::tuple<
    std::tuple<Args5...>,
    std::tuple<Args6...>
  >,
  // ... more of them
>;
```

this removes the problem of the parameter count, and we can have all the leaves with the same depth and length.
to get the unique types we use a permutation strategy:

> - Start from a sequence of types.
> - Generate all permutations of the sequence.
> - Execute the meta-function to benchmark on all the permutations.
> - Remove the timing taken by the identity meta-function.

#### The Pack Interface

we first need the type sequence itself, it should be something similar to a parameter pack, so we need to follow a given defintion.

> A type _P_ is a pack if:
>
> - `pack_size<P>::value` exists and returns the size of the pack.
> - `pack_element<I,P>::type` exists and returns the type at the I-th index in the pack.
> - `pack_rebind<P, T0,T1,...,Tn-1>::type` exists and corresponds to a pack of types `(T0,T1,...,Tn-1)`.

```cpp
// `pack_size`
// The interface to get the size of a pack type
template<class T>
struct pack_size;

// Returns the size of pack type
template<class T>
inline constexpr std::size_t pack_size_v = pack_size<T>::value;

// `pack_element`
// The interface to get the element of the pack type at the specified index
template<std::size_t I, class T>
struct pack_element;

// Returns the element of the pack type at the specified index
template<std::size_t I, class T>
using pack_element_t = pack_element<I, T>::type;

// `pack_rebind`
// The interface to rebind a pack type to another set of template parameters
template<class T, class... Args>
struct pack_rebind;

// Rebinds a pack type to another set of template parameters
template<class T, class... Args>
using pack_rebind_t = typename pack_rebind<T, Args...>::type;


// type traits combinations
// The base class to check if a type is a pack: the type is not a pack
template<class T, class = void>
struct is_pack: std::false_type{};

// Checks if a type is a pack through pack_size, pack_element, and pack_rebind
template<class T>
struct is_pack<T, std::enable_if_t<
  pack_size_v<pack_rebind_t<T,void>> == 1 &&
  std::is_void_v<pack_element_t<0, pack_rebind<T, void>>>>
>: std::true_type {};

// Returns true when the provided type is a pack, and false otherwise
template<class T>
inline constexpr bool is_pack_v = is_pack<T>::value;
```

```cpp
// Simplified defintion
template<class T>
struct is_pack<T, pack_element_t<
  pack_size_v<pack_rebind_t<T, void>> -1,
  pack_rebind_t<T, void>
>>: std::true_type {};
```

this could be done in SFINAE

```cpp
template<class T, class... Types>
struct if_pack;

template<class T>
struct if_pack<T>: std::enable_if<is_pack_v<T>> {};

template<class T, class True>
struct if_pack<T,True>: std::enable_if<is_pack_v<T>, True> {};

template<class T, class True, class False>
struct if_pack<T,True, False>: std::conditional<is_pack_v<T>, True, False> {};

template<class T, class... Types>
using if_pack_t = typename if_pack<T, Types...>::type;
```

or with c++20 concepts

```cpp
template<class T>
concept pack_type =
  pack_size_v<pack_rebind_t<T,void>> == 1 &&
  std::is_void_v<pack_element_t<0, pack_rebind<T, void>>>;
```

implementing the pack_index and indexer, and we want to adapt the tuple to the pack type:

```cpp
template<class... Args>
struct pack_size<std::tuple<Args...>> :
  std::integral_constant<std::size_t, sizeof...(args)> {};

template<std::size_t I, class... Args>
struct pack_element<I, std::tuple<Args...>> {
  using type = std::tuple_element_t<I, std::tuple<Args...>>;
};

template<class... Types, class... Args>
struct pack_rebind<std::tuple<Types...>, Args...> {
  using type = std::tuple<Args...>
};
```

we next build the functions for elements access, as well as recursive access, and getting the front and last elements of a given pack (even nested packs).
we also have truncation, resizing, appending, swapping as pack operations. and with these operations we can implement the permutation of the types.

Heap's algorithm (the problem is that the array is passed by reference)

```cpp
template<class F, class T, std::size_t N>
void permute(F&& f, std::array<T, N>& array. std::size_t k =N) {
  if (k == 1) {
    std::invoke(std::forward<F>(f), array);
  } else {
    permute(std::forward<F>(f), array, k-1);
    for (std::size_t =0; i < k-1; ++i) {
      if (k % 2== 0) {
        std::swap(array[i], array[k-i]);
      } else {
        std::swap(array[0], array[k-i]);
      }
      permute(std::forward<F>(f), array, k-1);
    }
  }
}
```

#### Meta Functions

C++ is type generic, but not _kind generic_.

```cpp
// these can exist together
template<class T> struct entity;
template<class T> struct entity<std::vector<T>>;
// but these cannot exist together
template<class T> struct entity;
template<template <class...> class T> struct entity;
```

for these we need wrappers to remove and create layers.

uniform meta-function invocations - what our framework should support

> - Traits
> - Templated call operator
> - Function objects

uniform meta-function invocation.

#### Computation and Benchmarks

generate and transform type permutations. adding specializations...

the user code is something like this.

```cpp
#include <metabench.h>

struct A{};
struct B{};
struct C{};
struct D{};

int main(int argc, char* argv[]) {
  return metabench::compute<std::tuple<A,B,C,D>,std::add_const, 3>();
}
```

and we can now benchmark the two implementations of adding tuples together.
we can also compare if there is a better implementation of std::tuple (without recursive inheritance).

</details>

### Nth Pack Element in C++ - A Case Study - Kris Jusiak

<details>
<summary>
Implementing `get` on variadic pack arguments (`...`)
</summary>

[Nth Pack Element in C++ - A Case Study](https://youtu.be/MLmDm1XFhEM), [slides](https://krzysztof-jusiak.github.io/talks/nth_pack_element/#/)

#### Motivation / Goal

use cases:

- performance
- reflection
- serialization

we want something like `std::get<N>` for tuples, but that will work on pack elements.

```cpp
std::get<int>(std::tuple{1,2,3}); // error: ambiguous, all elements are integers
std::get<0>(std::tuple{1,2,3}); // ok
std::get<1>(std::tuple{1,2,3}); // ok
std::get<2>(std::tuple{1,2,3}); // ok
std::get<3>(std::tuple{1,2,3}); // error

const auto t= std::tuple{1,2,3}; // std::tuple<int, int, int>
auto print = []<auto... Ns>(auto t, auto fn, std::index_sequence<Ns...>){
  return fn(std::get<Ns>(t)...); // nth_pack_element...
};

print(t, [](auto... ts) {(std::print("{} ",t), ...);} ,
std::make_index_sequence<std::tuple_size_v<decltype(t)>>{}
);
```

it's easy to get the first element, we simply expand the pack and use the first element.

```cpp
constexpr auto first_pack_element = [](auto first, auto... args) {return first;}(args); // immediately invoked function
static_assert(1 == first_pack_element(1,2,3));
```

getting the last element is also simple, relying on the fact the comma operator returns the last element by default.

```cpp
constexpr auto last_pack_element = (...,args);

static_assert(3 == last_pack_element(1,2,3));
static_assert(3 == (1,2,3)); // this works!
```

if we want both the first and the last elements together, we combine the two.

```cpp
constexpr auto [first, last] = [](auto first, auto... ts) {
  return std::tuple{first, (ts, ...)};
}(1,2,3)'
```

there is a proposal to bring those options into the standard and have it work with structured bindings (similar to javascript?)

```cpp
constexpr auto [first, ...ts] = std::tuple{1,2,3};
constexpr auto [...ts, last] = std::tuple{1,2,3};
constexpr auto [first, ...ts, last] = std::tuple{1,2,3};
```

### The `nth` pack element

so, the goal is to have something similar to `std::get<N>` of tuples, but working for template pack elements.

```cpp
template <auto N>
[[nodiscard]] constexpr auto get (auto t){

  return nth_pack_element<N>(t);
}
```

> `nth_pack_element` - returns element at **N** position **from** a variadic pack...
>
> Goal - Objectives
>
> - flexibility - supports different types
> - readability - easy to follow? how much "magic" is involved?
> - performance - is there an overhead / compilation time

three possible ways to implement this:

1. separate into elements before, nth element, elements after.
2. create a type with the type id as index and matches with the elements
3. language feature/extension
   1. _std::array_ (same type only)
   2. recursion
   3. arg expansion
   4. concept expansion
   5. boost.mp11 (meta programming)
   6. boost.preprocessor
   7. `\_\_type_pack_element` (clang)
   8. **@circle**
   9. Generalized pack deceleration and usage proposal (P1858)

> [IIFE](https://godbolt.org/z/aM58T1), proposal 1061 - "structured binding can introduce a pack"

```cpp
template<auto N>
const auto nth_pack_element([[magic]] auto..., auto nth, auto...){
  return nth;
}

template<auto N>
const auto nth_pack_element(auto... args){
  auto t = [[magic]] make_type_id{args...};
  return []<class T>(type_id<T, N>& t) {return *t}(t); // IIFE
}
```

| Option                                                             | Flexibility (works with multiple types)                               | Readability                                                                                         | Overhead (naive prediction)                                            |
| ------------------------------------------------------------------ | --------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| [Std.array (same types)](https://godbolt.org/z/1G7Gj16nx)          | no - only support the same type of elements                           | simple to follow                                                                                    | language feature (probably fast)                                       |
| [Recursion.1](https://godbolt.org/z/5WqMc5bjo)                     | works with different types                                            | mostly, if you understand pack stuff                                                                | slow - recursion is slow                                               |
| [Recursion.2](https://godbolt.org/z/vGh1nTnaq)                     | works with different types                                            | mostly, if you understand template stuff                                                            | slow - recursion is slow, but a bit better than the first option       |
| [Arg expansion](https://godbolt.org/z/a8vsqM848)                   | yes                                                                   | bad readability - "dark magic" involved                                                             | faster than recursion                                                  |
| [Concept expansion](https://godbolt.org/z/WPhe5MerW)               | yes                                                                   | bad readability - "dark magic" with concepts                                                        | faster than recursion                                                  |
| [boost.mp11](https://godbolt.org/z/zEM1eGqxK)                      | works with different types                                            | low readability - bringing in a new library                                                         | boost.mp11 is well optimized for compilation times                     |
| [boost.preprocessor](https://godbolt.org/z/ervdav89b)              | works with different types - but has limits on the number of elements | low readability - preprocessor macros, text replacement                                             | compilation overhead for generating all the functions for each element |
| [any](https://godbolt.org/z/vravKn7h3)                             | works with different types                                            | low readability - magic of struct inheritance                                                       | no recursion, but has some extra types to create                       |
| [\_\_type_pack_element (clang)](https://godbolt.org/z/jdM6v47P1)   | clang only                                                            | compiler inartistic                                                                                 | probably fast                                                          |
| [@Circle](https://godbolt.org/z/KGhP5aedx)                         | works with different types                                            | very readable, but requires using a template meta-programming framework which isn't in the standard | probably fast                                                          |
| [Generalized pack deceleration and usage](https://wg21.link/P1858) | yes                                                                   | yes (if accepted into c++23/26)                                                                     | language feature                                                       |

why isn't "get" a member of the tuple? it was avoided to not have (`./*template*/`) argument everywhere inside templates.

we also want our solution to work with tuples in the same way it does with pack arguments. something about a "map like tuple interface" which takes integer constant with type suffix literals. @circle also suggest having tuple as a language feature (and not a container). using c++20 ranges on tuples to get elements.

#### Benchmarking

use cases:

- many tuples with less than 20 elements
- a few big tuples with more than 100/1000 elements
- a lot of tuples of all cases

of course, the recursion solutions are really bad with high number of element in the tuple, the preprocessor solution also has problems at high numbers, but the argument expansion and concept expansion usually win. using language features (such as @circle meta programming) makes things much better, so if the generalized pack proposal is accepted, then it will win.

comparison to other languages shows the clang and gcc don't scale well.

</details>

### C++ Class Template Argument Deduction - History, Uses, & Enabling it for Classes - Marshall Clow

<details>
<summary>
Understanding the Class Template Argument Deduction, explicit and implicit deduction guides.
</summary>

[C++ Class Template Argument Deduction - History, Uses, & Enabling it for Classes](https://youtu.be/EPfPMW-rOtc)

CTAD - Class Template Argument Deduction. letting the compiler figure things out.

#### Brief History of CTAD

- Default Template Argument in c++03
- Template Argument deduction in c++03
- The <cpp>auto</cpp> keyword in c++11
- Class Template Argument Deduction in c++17
- CTAD for Alias templates in c++20
- CTAD for aggregates in c++20
- CTAD for inherited constructors in c++23

in many cases, the compiler can figure out what we want, and that is the correct thing, so if we don't write it manually, we won't get it wrong.

here is a code without help from the compiler, it's obviously a disaster

```cpp
template<typename Iterator>
bool findA(Iterator first, Iterator last){
  std::vector<typename std::iterator_traits<Iterator>::value_type, std::allocator<typename std::iterator_traits<Iterator>::value_type> > v(first, last);

  std::allocator<typename std::vector<typename std::iterator_traits<Iterator>::value_type, std::allocator<typename std::iterator_traits<Iterator>::value_type> >::iterator it = std::find<std::vector<typename std::iterator_traits<Iterator>::value_type, std::allocator<typename std::iterator_traits<Iterator>::value_type> >::iterator, int>(v.begin(), v.end(),3);

  return it == v.end();
}
```

in c++98, <cpp>std::vector</cpp> has a value type and an allocator. now that the allocator has a default value, we can simplify the code a little

```cpp
template<typename Iterator>
bool findB(Iterator first, Iterator last){
  std::vector<typename std::iterator_traits<Iterator>::value_type> v(first, last);

  std::allocator<typename std::vector<typename std::iterator_traits<Iterator>::value_type>::iterator it = std::find<std::vector<typename std::iterator_traits<Iterator>::value_type>::iterator, int>(v.begin(), v.end(), 3);

  return it == v.end();
}
```

when we call the templated function, we usually don't have to specify the template parameters. the compiler deduces them from the types of the parameters passed to the function.

```cpp
template<typename Iterator>
bool findC(Iterator first, Iterator last){
  std::vector<typename std::iterator_traits<Iterator>::value_type> v(first, last);

  std::allocator<typename std::vector<typename std::iterator_traits<Iterator>::value_type>::iterator it = std::find(v.begin(), v.end(), 3);

  return it == v.end();
}
```

in c++11, we got the <cpp>auto</cpp> keyword, which lets the compiler figure out the type.

```cpp
template<typename Iterator>
bool findD(Iterator first, Iterator last){
  std::vector<typename std::iterator_traits<Iterator>::value_type> v(first, last);

  auto it = std::find(v.begin(), v.end(), 3);
  return it == v.end();
}
```

and in c++17, we got the class template argument deduction, so the compiler can figure out the type.

```cpp
template<typename Iterator>
bool findE(Iterator first, Iterator last){
  auto v(first, last);
  auto it = std::find(v.begin(), v.end(), 3);
  return it == v.end();
}
```

and now we see that this function is bugged, it returns true when the element does not exist, and false when it does.

#### What does CTAD Do?

if we instantiate a template class without any template arguments, the compiler fills in the argument. it helps in removing some of `make_xxx` functions (<cpp>std::make_pair</cpp>, <cpp>std::make_tuple</cpp>).

```cpp
auto p1 = std::make_pair(1, 23L);
auto p2 = std::pair{1, 23L};

auto t1 = std::make_tuple("abc"sv, 12, nullptr);
auto t2 = std::tuple{"abc"sv, 12, nullptr};
```

CTAD works by deduction guide, a new feature from C++17, one guide is the implicitly created deduction guide which the compiler generates, and one is the explicit guide that the programmer writes.

```
template <template arguments>
class-name(argument-type-list) -> class-name<template-argument-list>;
```

an actual example, a deduction guide for a copy constructor.

```cpp
template <typename T, typename Alloc>
vector(vector<T, Alloc>) -> vector<T, Alloc>;
```

**Vector example**

we start by looking at all the constructors, and we map them into a deduction guide.

| Constructor call                       | Deduction guide                      |
| -------------------------------------- | ------------------------------------ |
| `vector()`                             | no guide                             |
| `vector(Alloc)`                        | no guide                             |
| `vector(size_t, Alloc)`                | no guide                             |
| `vector(const vector & v)`             | `-> decltype(v)`                     |
| `vector(const vector & v, Alloc)`      | `-> decltype(v)`                     |
| `vector(vector & v)`                   | `-> decltype(v)`                     |
| `vector(vector & v, Alloc)`            | `-> decltype(v)`                     |
| `vector(size_type, Value_type)`        | `-> vector<Value_type>`              |
| `vector(size_type, Value_type, Alloc)` | `-> vector<Value_type, Alloc>`       |
| `vector(iter, iter)`                   | `-> vector<iter::value_type>`        |
| `vector(iter, iter, Alloc)`            | `-> vector<iter::value_type, Alloc>` |
| `vector(initializer_list<T>)`          | `-> vector<T>`                       |
| `vector(initializer<T>, Alloc)`        | `-> vector<T, Alloc>`                |

only the constructors using iterators are explicitly written.

```cpp
template<class Iterator,
  class value_type = typename iterator_traits<Iterator>::value_type,
  class Alloc = allocator<value_type>
  >
vector (Iterator, Iterator, Alloc = Alloc()) -> vector<value_type, Alloc>;
```

because unlike initializer lists, the compiler doesn't know about iterators, so the deduction guide must be explicitly written.

#### Writing Our Own Deduction Guide

If we want write deduction guides to our own classes, we start by writing the tests (stuff which we should compile), and most of them will be handled by the implicit deduction guides. if possible, we should prevent wrong deductions from being materialized. but if want deduction guides for containers, we have to answer two questions:

- what is an iterator?
- what is an allocator?

there are many pitfalls, so we need compiler tests to make sure our types are what we expect them to be. some classes don't have deduction guide on purpose.

```cpp
int* one = new int;
int* two = new int[5];

std::share_ptr p1(one;)
std::share_ptr p2(two;)
```

in this code, both are pointer to int, but the delete function should be different. so it was preferred to disallow deduction guide than to get it wrong.

#### Post C++17

- CTAD for Alias templates(c++20)
- CTAD for aggregates(c++20)
- CTAD for inherited constructors(c++23)

```cpp
template<typename T>
myVec = using std::vector<T>;
myVec v(first, last);

template<typename T>
struct Pt(T x; T y;);
Pt p(4L, 3L);
```

CTAD helps reduce boilerplate code, it's not terribly hard to write or test. Failure to deduce is a valid option, the worst case is that the user has to specify the type. miss-deducing and ambiguity are much worse.

</details>

### Binary Object Serialization with Data Structure Traversal & Reconstruction in Cpp - Chris Ryan

<details>
<summary>
Another Serialization library.
</summary>

[Binary Object Serialization with Data Structure Traversal & Reconstruction in Cpp](https://youtu.be/rt-c7igYkFw),

serialization library for c++14/17, taking blobs of data, writing them as a string and then read them back in.

#### Serialization

it's easy to serialize simple hierarchical data, but it's hard to serial stuff with inside reference, as we can get stuck in loops.

objects can value, pointers, arrays and references. there are problems with serializing references, and it's also dangerous to work with const data, as it can violate rules.

#### Implementation

the simple, intuitive approach to serializing requires defining a serializing method. the archive class works for each type of object (value, pointer, arrays). we define functions by using SFINAE and we employ <cpp>std::enable_if_t<T,void></cpp>

```cpp
void Save(const MyClass& myClass) {
  FileSource file("filename.ext");
  Archive ar(file, save);
  ar << myClass;
}

void Load(MyClass& myClass) {
  FileSource file("filename.ext");
  Archive ar(file, load);
  ar >> myClass;
}

void MyClass::Serialize(Archive& arr) {
  if (ar.IsStoring()) {
    arr << m_data1;
    arr << m_data2;
  }
  else {
    arr >> m_date1;
    arr >> m_date2;
  }
}
```

going over the code, looking at a pointer data serialization, base classes and virtual inheritance with CRTP, object id and keeping track of repeated objects being loaded, using TypeInfo and mapping a hash to it.

#### Extra Features and C++20 Concepts

we can optimize for different integer values and save some space for small numbers, working for complex nested data types.

if we want to take this to C++20, we can switch the constexpr <cpp>std::enable_if</cpp> SFINAE and use concepts.

demo of using a client/server connection over sockets and sending data between them.

</details>

##

[Main](README.md)
