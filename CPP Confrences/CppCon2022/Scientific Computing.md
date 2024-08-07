<!--
ignore these words in spell check for this file
// cSpell:ignore Kokkos asynchrony München Frankfürt Würzburg Kriegsmarine Bomba submdspan Amdahl vertexlist edgelist vertice Philox Rpass noalias parallelizable Cuda bina binb tinb Scholes
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Scientific Computing

### C++ Performance Portablity - A Decade of Lessons Learned - Christian Trott

<details>
<summary>
Lessons learned from using the Kokkos framework for super computers.
</summary>

[C++ Performance Portablity - A Decade of Lessons Learned](https://youtu.be/jNGGKFkt4lA), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/ChristianTrott-ADecadeOfPerformancePortability.pptx), [Kokkos github](https://github.com/kokkos).

> Goals of the talk:
>
> - Demonstrate issues of heterogeneous hardware and performance portability from a C++ software design perspective.
> - Show how the C++ Performance Portability Programming Model Kokkos addresses these concerns in its design.
> - Connect it all to the design of C++23 mdspan and (hopefully) C++26 stdBLAS.
> - Summarize lessons learned beyond C++ design decisions to achieve practical performance portability.

#### The HPC Hardware and Software Landscape

the super computers of the world have different CPUs and different GPUs.

this means that for each combination, a the different hardware architecture determines how the code should be built for them.

**Kokkos **

> A C++ Programming Model for Performance Portability.

it replaces the usage of CUDA or OpenMP, HIP at the source code, and it's then mapped to the correct framework. it's an eco-system.

#### Performance Portability Challenges And How to Tackle Them

**Device and Host functions**, host and device compilations, which have different instruction sets, which brings a load of problems.

**Dispatch to devices and the challenge of value semantics**, trying to access memory in the CPU from the GPU, so it's pointless to use value semantics and references. this means that all the data structures with value semantic aren't usable, so there goes the C++ Standard Library.

**ParUnseq - The need to restrict guarantees to users** - non-indenes branches, not all accelerators have the concept of threads.

the primary data structure in Kokkos is the view, a reference counted multi dimensional array.

**Data accessibility and how to control where to execute** - execution spaces and memory spaces. the memory space can tell us where to dispatch the execution to.

**The curse of cached vs coalesced data access**, C layout vs fortran layout, row layout vs column layout. different between CPU and GPU, the optimal storage order depends on the architecture.

**MemoryTraits: Customization point for special access hardware**. special memory access paths. we can try and map the usage to the hardware, like read only caches, etc...\
even if not all cases exist for all architecture, he framework should try and get it right.

**Hierarchial hardware and parallelism** - GPU architecture, number of "threads" varies.

**Policies and Algorithms (patterns)**, fundamental algorithms combined with execution policies.

**Resource handles and asynchronicity** execution policies and resource handles. asynchronicity is baked into most frameworks.

Kokkos Abstractions:

> | Phase   | Data structures                                      | Parallel execution                                             |
> | ------- | ---------------------------------------------------- | -------------------------------------------------------------- |
> | "Where" | **Memory spaces** (HBM, DDR, Non-volatile, scratch)  | **Execution Spaces** (CPU, GPU, Executor mechanism)            |
> |         | **Memory Layouts** (Row/Column-major tiles, strided) | **Execution Patterns** (parallel_for/reduce/ scan, task-spawn) |
> | "How"   | **Memory Traits** (streaming, atomic, restrict)      | **Execution Policies** (range, team, task-graph)               |

#### Building Blocks in the (future) ISO C++ Standard

mdspan - multidimensional non-owning array wrapper. it has layout types, accessors as customization points.

- typesafe memorySpace access
- Atomic Access
- std::linalg - linear algebra

#### Final Thoughts

#### Lessons learned:

> 1. Single source compilation for heterogeneous architecture requires some form of “host/device” function attributes.
> 2. Data structures accessed in compute kernels need to have reference semantics.
> 3. Parallel un-sequenced loops are the fundamental building block for performance portable code.
> 4. Execution and Memory Resources need to know about each other.
> 5. Data Layout Mappings need to be configurable and integral part of the API design.
> 6. More and more special hardware capabilities exist – for full performance these need to be leveraged.
> 7. Hardware is hierarchical and algorithms often don’t expose enough parallelism on a single level: need to support nested parallelism.
> 8. Make Simple things Simple – And complex things easier.
> 9. Build helpful debugging and profiling capabilities into your APIs.

</details>

### HPX - A C++ Library for Parallelism and Concurrency - Hartmut Kaiser

<details>
<summary>
An overview of HPX
</summary>

[HPX - A C++ Library for Parallelism and Concurrency](https://youtu.be/npufmMlGOoM), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/HPX-A-C-Standard-Library-for-Parallelism-and-Concurrency-CppCon-2022-1.pdf)

> HPX – An Asynchronous Many-task
> Runtime System

differences in performance between compute-bound and memory-bound parallel tasks. also difference performance patterns depending on the number of cores.

HPX: a threading implementation, more efficient than naive threads (jthread, std::thread, pthread), with several functional layers, conforming to the standard, and offers some extensions. allows for distributes execution.

this talk will focus on parallel loop and algorithms

> - Simple iterative algorithms
>   - One pass over the input sequence.
>   - for_each, copy, fill, generate, reverse, etc.
> - Iterative algorithms ‘with a twist’
>   - One pass over the input sequence.
>   - Parallel execution requires additional operation after first pass, most of
>     the time this is a reduction step
>   - min_element, all_of, find, count, equal, etc.
> - Scan based algorithms
>   - At least three algorithmic steps.
>   - inclusive_scan, exclusive_scan, etc.
> - Auxillary algorithms
>   - Sorting, heap operations, set operations, rotate

parallelization can work on CPU (threads and cores), and on GPUs.

#### Parallelize Loops

execution policies:

> Convey guarantees/requirements imposed by loop body
>
> - seq: execute in-order (sequenced) on current thread
> - unseq: allow out-of-order execution (un-sequenced) on current thread - vectorization
> - par: allow parallel execution on different threads
> - par_unseq: allow parallel out-of-order (vectorized) execution on different threads

in the future the standard might include _std::simd_, which will require explicit vectorization. but that's still in the experimental stage.

HPX adds more policies, explicit parallelized vectorization and execution.

the first example uses std::future to launch threads and parallelize on them, it cuts the input data into chunks.

#### Background

Amdahl's Law (Strong scaling), the speed up of parallelization is capped by the number of processors and how much of the code can be parallelized

$
S = \frac 1{(1-p) + \frac {P}{N}}
$

SLOW - problems with parallelization

> - Starvation
>   - Insufficient concurrent work to maintain high utilization of resources.
> - Latencies
>   - Time-distance delay of remote resource
>     access and services
> - Overheads
>   - Work for management of parallel actions and resources on critical path which are not necessary in sequential variant.
> - Waiting for Contention resolution.
>   - Delays due to lack of availability of
>     oversubscribed shared resources.

there is a U-shaped curve of gains from spliting chunks. the overhead can be greater than the gains, but if the there are too many chunks (with smaller data), the performance will go down. this goes together with the number of cores.

#### Executors

> Executors abstract different task launching infrastructures
>
> - Synchronization using futures
>   - HPX historically uses futures as main means of coordinating.
> - Synchronization using sender/receivers (C++26?)
> - C++ standardization focusses on developing an infrastructure for anything related to asynchrony and parallelism.
> - P2300: std::execution (senders & receivers)
> - Computational basis for asynchronous programming
> - Current discussions focus on integrating parallel algorithms.

some examples: execution policies, attaching to an executor, parallel task execution (asynchronous and futures), senders & receivers. eager and lazy executions.

**Explicit vectorization**:
simd - single instruction, multiple data.

**Linear algebra**, a proposal for standardization, might also work with execution policies. `std::linalg::scale`

</details>

### Graph Algorithms and Data Structures in C++20 - Phil Ratzloff & Andrew Lumsdaine

<details>
<summary>
A proposal for adding a graph container to the standard library.
</summary>

[Graph Algorithms and Data Structures in C++20](https://youtu.be/jCnBFjkVuN0), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/C-Phil-Ratzloff-CppCon-2022.pdf), [graph-v2 github repository](https://github.com/stdgraph/graph-v2).

creating a graph library and integrating it into the C++ Standard Library (towards std::graph). since 2001 there is boost:graph, and since 2017 North West Graph library (NWgraph).

#### What Is a Graph?

> A graph _G = (V, E)_ is a set of _vertices V_, points in a space,and _edges E_, links between these vertices.\
> _Edges may or may not be oriented_, that is, directed or undirected, respectively. Moreover, _edges may be weighted_, that is, assigned a value."

the vertices and the graph itself may also be assigned a value as well. a graph can be the distance between the cities.

the raw form of the data doesn't work well, so we convert it to an Adjacency List, a table with the locations, and each location item has the Edge List as inner edges in a linked list.\
_(note: the elements in the linked list are separate destinations, not "routes")_

> | From      | To        | Distance |
> | --------- | --------- | -------- |
> | Frankfürt | Mannheim  | 85       |
> | Frankfürt | Würzburg  | 217      |
> | Frankfürt | Kassel    | 173      |
> | Mannheim  | Karlsruhe | 80       |
> | Karlsruhe | Augsburg  | 250      |
> | Augsburg  | München   | 84       |
> | Würzburg  | Erfurt    | 186      |
> | Würzburg  | Nürnberg  | 103      |
> | Nürnberg  | Stuttgart | 183      |
> | Nürnberg  | München   | 167      |
> | Kassel    | München   | 502      |

Adjacency + Edges List

> | id  | location (vertices) | inner ranges (edges)     |
> | --- | ------------------- | ------------------------ |
> | 0   | Frankfürt           | [1,85], [4,217], [6,173] |
> | 1   | Mannheim            | [2,80]                   |
> | 2   | Karlsruhe           | [3,250]                  |
> | 3   | Augsburg            | [8,84]                   |
> | 4   | Würzburg            | [5,103], [7,186]         |
> | 5   | Nürnberg            | [8,502], [9,183]         |
> | 6   | Kassel              | [8,85]                   |
> | 7   | Erfurt              | &empty;                  |
> | 8   | München             | &empty;                  |
> | 9   | Stuttgart           | &empty;                  |

```cpp
struct route {int target_id; double distance;};
using edges_type = std::list<route>;
struct vertex {edges_type edges; string name;};
using vertices_type = std:vector<vertex>;
```

we want this to be bi-directional graph, so we duplicate the edges.

> | id  | location (vertices) | inner ranges (edges)      |
> | --- | ------------------- | ------------------------- |
> | 0   | Frankfürt           | [1,85], [4,217], [6,173]  |
> | 1   | Mannheim            | [2,80], [0,85]            |
> | 2   | Karlsruhe           | [3,250], [1,80]           |
> | 3   | Augsburg            | [8,84], [,250]            |
> | 4   | Würzburg            | [5,103], [7,186], [0,217] |
> | 5   | Nürnberg            | [8,502], [9,183], [4,103] |
> | 6   | Kassel              | [8,85], [0,173]           |
> | 7   | Erfurt              | [4,186]                   |
> | 8   | München             | [3,84], [5,167], [6,502]  |
> | 9   | Stuttgart           | [5,183]                   |

but this format doesn't always work, especially if you have many properties and connections. an edge must belong in one container. so there are trade offs in performance and usability. there are all sorts of ways to define graphs.

other kinds of graphs:

- Bipartite and n-partite graphs – partitioned graphs
- HyperGraphs – more than one vertices on an edge

> Challenges
>
> - Enable high performance algorithms
> - What algorithms to include initially?
> - How to represent a container as a range-of-ranges?
> - How to represent STL separation of container, iterator and algorithm?
> - “Bring your own graph”
> - How are user-defined values defined?
> - How to use modern C++ to make it easy and fun?
> - Where are the boundaries?

Naming Conventions

| Template Parameter | Variable Name            | Description                      |
| ------------------ | ------------------------ | -------------------------------- |
| G                  | g                        | Graph object                     |
| V                  | u, v, x, y               | Vertex (reference)               |
| VId                | uid, vid, xid, yid, seed | Vertex Id                        |
| VR                 | ur, vr                   | Vertex Range                     |
| VI                 | ui, vi                   | Vertex Iterator                  |
| VV                 | val                      | Vertex value (user-defined type) |
| VVF                | vvf                      | Vertex Value Function            |
| E                  | uv, vw                   | Edge (reference)                 |
| ER                 | er Edge                  | Range                            |
| EI                 | uvi, vwi                 | Edge Iterator                    |
| EV                 | val                      | Edge value                       |
| EVF                | evf                      | Edge Value Function              |

#### Example

we start with defining the data (doubled for the route back).

```cpp
// city data (vertices)
using city_id_type = int32_t;
using city_name_type = string;
std::vector<city_name_type> city_names = {"Frankfürt", "Mannheim", "Karlsruhe", "Augsburg", "Würzburg",
"Nürnberg", "Kassel", "Erfurt", "München", "Stuttgart"};
// edge data (EdgeList)
using route_data = copyable_edge_t<city_id_type, double>; // {source_id, target_id, value}
std::vector<route_data> routes_doubled = {
{0, 1, 85.0}, {0, 4, 217.0}, {0, 6, 173.0}, //
{1, 0, 85.0}, {1, 2, 80.0}, //
{2, 1, 80.0}, {2, 3, 250.0}, //
{3, 2, 250.0}, {3, 8, 84.0}, //
{4, 0, 217.0}, {4, 5, 103.0}, {4, 7, 186.0}, //
{5, 4, 103.0}, {5, 8, 167.0}, {5, 9, 183.0}, //
{6, 0, 173.0}, {6, 8, 502.0}, //
{7, 4, 186.0}, //
{8, 3, 84.0}, {8, 5, 167.0}, {8, 6, 502.0}, //
{9, 5, 183.0},
};
```

the rr_adaptor stands for "ranges range"

```cpp
struct route { // edge
  city_id_type target_id = 0;
  double distance = 0.0; // km
};
using AdjList = std::vector<std::list<route>>; // range of ranges
using G = rr_adaptor<AdjList, city_names_type>; // graph
G g(city_names, routes_doubled);
// Useful demo values_city_id_type frankfurt_id = 0;
vertex_reference_t<G> frankfurt = *find_vertex(g, frankfurt_id);
```

traversal example (creating the graph in text from)

```cpp
cout << "Traverse the vertices & outgoing edges" << endl;
for (auto&& [uid, u] : vertexList(g)) { // [id,vertex&]
  cout << city_id(g, uid) << endl; // city name [id]
  for (auto&& [vid, uv] : incidence(g, uid)) { // [target_id,edge&]
    cout << " --> " << city_id(g, vid) << endl;  // "--> "target city" [target_id]
  }
}
```

Dijkstra's Shortest path algorithm defintion, takes a graph, starting location, a weighting function, and two out parameters.

```cpp
void dijkstra_clrs(
G&& g, // graph
vertex_id_t<G> seed, // starting vertex_id
Distance& distance, // out: distance[uid] of uid from seed
Predecessor& predecessor, // out: predecessor[uid] of uid in shortest path
WF weight = [](edge_reference_t<G> uv) { return 1; } // weight function (non-negative)
)
```

shortest path (segments): the weight is always 1 for each possible segment, so starting from Frankfürt, we can list all the routes to other locations and how many segments are required.

```cpp
auto weight_1 = [](edge_reference_t<G> uv) -> int
{ return 1; };
std::vector<int> distance(size(vertices(g))); // output parameter
std::vector<vertex_id_t<G>> predecessor(size(vertices(g))); // output parameter
dijkstra_clrs(g, frankfurt_id, distance, predecessor, weight_1);
cout << "Shortest distance (segments) from "
<< city(g, frankfurt_id) << endl;
for (vertex_id_t<G> uid = 0; uid < size(vertices(g)); ++uid) {
  if (distance[uid] > 0){
    cout << " --> " << city_id(g, uid) << " - " << distance[uid] << " segments" << endl;
  }
}
```

shortest path (distance in kilometers), this time the weight function returns the edge value (distance). so the output will be the shortest distances from the starting location.

```cpp
auto weight = [&g](edge_reference_t<G> uv)
{ return edge_value(g, uv); }; // { return 1; };

std::vector<double> distance(size(vertices(g))); // output parameter
std::vector<vertex_id_t<G>> predecessor(size(vertices(g))); // output parameter
dijkstra_clrs(g, frankfurt_id, distance, predecessor, weight);
cout << "Shortest distance (km) from " << city(g, frankfurt_id) << endl;
for (vertex_id_t<G> uid = 0; uid < size(vertices(g)); ++uid){
  if (distance[uid] > 0) {
    cout << " --> " << city_id(g, uid) << " - " << distance[uid] << "km" << endl;
  }
}
```

finding the farthest city and the shortest path to it.

```cpp
// Find farthest city
vertex_id_t<G> farthest_id = frankfurt_id;
double farthest_dist = 0.0;

for (vertex_id_t<G> uid = 0; uid < size(vertices(g)); ++uid) {
  if (distance[uid] > farthest_dist) {
    farthest_dist = distance[uid];
    farthest_id = uid;
  }
}

cout << "The farthest city from " << city(g, frankfurt_id)
<< " is " << city(g, farthest_id) << " at " << distance[farthest_id] << "km" << endl;

cout << "The shortest path from " << city(g, farthest_id)
<< " to " << city(g, frankfurt_id) << " is: " << endl
<< " ";

// Output path for farthest distance
for (vertex_id_t<G> uid = farthest_id; uid != frankfurt_id; uid = predecessor[uid]) {
  if (uid != farthest_id) {
    cout << " -- ";
    cout << city_id(g, uid);
  }
}
cout << " -- " << city_id(g, frankfurt_id) << endl;
```

#### Algorithms

now we go over the algorithm, it's templated, it has a _requires_ clause, it takes an adjacency list, the edge weight function is a concept that takes an edge reference and returns an integral type.\
a vertex range concept, a target edge and source edge concepts and the adjacency list concept.

some algorithms are confirmed to be part of the proposal, and some might wait.

#### Views

planned views (ways of iterating over the graph). there is a overload set that allows for getting the reference to the iterated object, the value and other optional stuff.

- vertexlist and vertex_view - vertex id (overloads: vertex reference, value)
- incidence and edge_view - target_id (overloads: edge reference, source_id, edge value)
- neighbors and neighbor_view - neighboring vertices (overloads: target reference, source_id, target value)
- edgelist - source id, target id, edge reference
- depth_first_search - either by vertice or edge,
- breadth_first_search - same as depth_first, but different ordering
- topological_sort - can't visit a location without visiting all of the sources for it before

#### `csr_graph` Graph Container

the proposed graph container

> Graph Containers: Unique Among Containers
>
> - Range of ranges
> - All functions are free functions (c.f. begin, end, size, empty) (no member functions)
> - All functions are customization points
> - User-defined values are optional on edge, vertex and graph
>
> `csr_graph` Graph Container:
>
> - Compressed Sparse Row (Matrix)
> - High performance
> - Compact memory use
> - Static structure – can’t change after construction
> - Values can change
> - Values are stored separately from structure

```cpp
template <class EV = void, // edge value type
          class VV = void, // vertex value type
          class GV = void, // graph value type
          integral VId = uint32_t, // vertex id type
          class Alloc = allocator<uint32_t>> // for internal containers
class csr_graph;

using G = std::graph::csr_graph<double, std::string_view, std::string>;
```

we can construct a graph from edges only, (edge projection function), or from edges values and vertices values.

Interface:

- functions
  - graph and vertex functions
  - edges functions
  - source edge functions
- types
  - graph and vertex types - iterator,reference, values
  - edge types - iterator,reference, values
- traits
  - unordered_edge
  - ordered_edge
  - adjacency_matrix
- concepts

#### Other Graph Containers

> Integrating External Graphs
>
> - csr_partite_graph?
> - rr_adaptor
> - dynamic_graph

some function would have to be overloaded, some overloads would be optional. partitioned graphs (holding different sub graphs with different vertex and edge types). dynamic graphs would allow us to have different inner containers.

</details>

### Breaking Enigma With the Power of Modern C++ - Mathieu Ropert

<details>
<summary>
Trying to break the Enigma code by using C++ and brute force techniques.
</summary>

[Breaking Enigma With the Power of Modern C++](https://youtu.be/zx3wX-fAv_o), [slides](https://docs.google.com/presentation/d/16UAV1ZZ_igk122BzqjUN-2GgOCnK6-D2K-7g31LDh68/edit#slide=id.g152a159cfc1_0_0)

The project to break the enigma cipher was called "Ultra" - what would it take to replicate it today with modern c++?

#### The Enigma

starting with explaining how the enigma machine worked:

- Rotors - simple substituion
- Chained Rotors - each key stroke changes the far-most rotor, and each complete cycle of a rotor moves an inner rotor.
- no printer, uses light to indicate the output.

complexity:

- 3 rotors, at any order => 6 options
- 26 *26 *26 initial rotor positions => 17,756 options
- increased complexity options
  - change when first rotation happens for each rotor
  - extra letter pairs substituion wires
  - add possible rotors options
  - use another rotor (Kriegsmarine M4 model)

so the basic model had 105,456 possible initial positions, and the advanced model has much, much more initial positions.

$
672*(26^4)*(26^3)*150,738,274,9372,50 \approx 8*(10^26)
$

the key is distributed physically in advanced, and each message starts with a 4 letter code, which uses the common settings, the rest of the message is encrypted with the rotors set to that unique message code.

however, the german had the key letters repeated twice (to avoid messing up the message), which would be used by the "Bomba" machine to try and find the correct code. once that changed, there were other ways to find patterns in the messages. the allied forces were able to read German ciphers for most of the war. the Germans knew it was theoretically possible, but they believed it was too costly to be done in practice.

#### Cracking the Code With Modern computers

in the 80's a new set of messages were discovered in sunken submarine, so we had the data but not the keys. in 2012 most of them were cracked by the Enigma@Home project which used distributed work to test out keys.

we focus on one message from the set (P1030681), it is 372 characters long.

```cpp
input = m_plugBoard[input - 'A'];

input = m_rotor[3].m_wiring[input -'A' + offsets[3] + 26];
input = m_rotor[2].m_wiring[input -'A' + offsets[2] - offsets[3] + 26];
input = m_rotor[1].m_wiring[input -'A' + offsets[1] - offsets[2] + 26];
input = m_rotor[0].m_wiring[input -'A' + offsets[0] - offsets[1] + 26];

input = m_reflector.wiring[input - 'A' - offsets[0] + 26];

input = m_rotor[0].m_reversed_wiring[input -'A' + offsets[0] + 26];
input = m_rotor[1].m_reversed_wiring[input -'A' + offsets[1] - offsets[0] + 26];
input = m_rotor[1].m_reversed_wiring[input -'A' + offsets[2] - offsets[1] + 26];
input = m_rotor[0].m_reversed_wiring[input -'A' + offsets[3] - offsets[2] + 26];

input = enigma::rotors [static_cast<int>(enigma::rotor_index::ETW)].m_wiring[input -'A' - offsets[3] + 26];

input = m_plugboard[input - 'A']
```

rotor implementation

- 26 character string
- generate the forward and reverse wiring by repeating the action 3 times.
- `output = wiring[(input+offset) % 26]`
- but module operations are costly, it creates a lot of assembly code.
- instead, we pay with memory and duplicate the data before and after, so we don't care if we go above or below the target.
- `output = wiring[input + offset + 26]`

the first attempt will cost us about 52 hours, this is after simplifying the problem space. we next try using multiple cores.

```cpp
std::array<char, 26> values;
std::iota(std::begin(values), std::end(values,0));
std::for_each(std::execution::par_unseq,
  std::begin(values),
  std::end(values),
  [&](char right_ring_setting){/*...*/})'
```

this gets us down to about 3.5 hours on a 16 core machine.

the third attempt relies on knowing that with a short message, the inner rotors aren't likely to turn much, so we can assume the setting is zero, and check for partial match (1/4 correct), and then fine tune the ring setting positions afterwards. this get us down to 6 minutes.

But we cheated earlier because we ignored the plugboard setup (letter swapping), which had 150 trillion possible combinations. so brute force won't help us anymore.

1. 10 cables to switch letter pairs mean that 6 letters are correct.
2. but if we use an empty plugboard, we can count the top letter, and if they match the expected message more than a threshold, we can fine tune afterwards
3. the problem space is split again

the final touch is to ignore the ring settings entirely, this means that a part of the message will be correct, and this fragment will always have a similar length. and this will give us another way to determine that we are on the correct way. this brings us down to $26*26*26$ passes before fine tuning, which is trivial work.

this relies on the fact that the enigma was a machine, and didn't use a true encryption key like modern computers use. getting one bit wrong in sha256 key creates a complete garbage output, but with the enigma, missing the configuration by one position will still produce a not random message.

according to him the heuristic works for short messages, could be better if it used more cryptographic techniques, but a true 60 bits cipher would still be too hard for one machine.

</details>

### MDSPAN - A Deep Dive Spanning C++, Kokkos & SYCL - Nevin Liber

<details>
<summary>
The road of development leading up to C++23 std::mdspan.
</summary>

[MDSPAN - A Deep Dive Spanning C++, Kokkos & SYCL](https://youtu.be/lvylnaqB96U), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/MDSPAN-A-Deep-Dive-Spanning-C-Kokkos-SYCL.pdf)

mdspan - a non-owning multidimensional array view for c++ 23.

```cpp
template <class ElementType,
  class Extent,
  class LayoutPolicy = layout_right,
  class AccessorPolicy = default_accessor<ElementType>>
struct mdspan{
  template <class... OtherIndexTypes>
  explicit constexpr mdspan(data_handle_type p, OtherIndexTypes... exts);
  // ...
  template <class... OtherIndexTypes>
  constexpr reference operator[](OtherIndexTypes... indices) const;
};
```

- ElementType - the elements in the arrays
- Extent - the dimensions of the multidimensional array.
  - `template<class IndexType, size_t... Es> class extents`
  - can be static (known at compile time) or std::dynamic_extent.
- LayoutPolicy
  - layout_right = row major order
  - layout_left = column major order
  - layout_stride
  - user defined
  - LayoutPolicy::mapping
- AccessorPolicy
  - customize the data_handle_type (pointer) and refences behavior.
  - decorations
  - remote memory
  - compressed memory
  - atomic access with `std::atomic_ref`

to construct it with use the date_handle_type (pointer) and specify the extents (dimensions).

the idea started in 2014 `array_view`, that was based on C++AMP, it used only static dimensions. but it wasn't enough. while the committee wanted it, but it got rejected at the technical specification stage of **Arrays TS** which eventually fell apart at 2016.

it was then suggested for Library Fundamentals, and over the time more papers started coming out pointing out issues for performance, compatiblity with vectoring, etc...

at the same time, there ws `kokkos::view`, which was a multi-dimensional array of zero or more dimensions, it had some special notations, and was sometimes owning and sometimes not, sometime reference counting and sometimes not.

back in 2015, there were proposal for `shared_array` and `weak_array` (owning and non-owning), and more changes for `array)view`. the next version removed the multidimensional part and became `std::span`.

over the years, more changes, name changes to `array_ref`. more committee, name changes again to `mdspan`. and even the use of the comma operator was deprecated in square brackets.

and then 2020.

starting to be explicit about _trivially copyable_, copy constructor and copy assignment operator might run code, which would be weird if the source and destination aren't at the same device. comma in square brackets now has new meaning. CTAD - class template argument deduction from c++17 helps us when we have many template parameters.

then 2022, many revisions to get it into c++23 features set, debating about the the name the first parameter, and others.

the future will have _mdarray_ - an owning multi-dimensional array type, there are still some things to be solved. _submdspan_ - a sub mdspan that uses a 'slice' of the data.

</details>

### Fast, High-Quality Pseudo-Random Numbers for Non-Cryptographers in C++ - Roth Michaels

<details>
<summary>
Using the Pseudo Random Numbers in audio plug-ins.
</summary>

[Fast, High-Quality Pseudo-Random Numbers for Non-Cryptographers in C++](https://youtu.be/I5UY3yb0128), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Fast-High-Quality-Pseudo-Random-Numbers-CPPCon2022-Roth-Michaels.pdf).

what is the best way to seed random numbers? the standard supports std::mt19937, but is it really the best thing to use? should we focus on "better" random generators (he says that we should, as our code might end up being used in something important).

when we use random number generators are are concerned with performance (how fast we get the numbers, how they mesh with the program design) and quality (how random they are). In the field of sound editing/mixing, they use psychoAcoustic analysis.

Types of random numbers:

- truly random
- pseudo random - "deterministic algorithm that approximates randomness"
- quasi random

we focus of pseudo Random numbers.

random number sources:

- Coin toss
- Fair die roll
- `/dev/random`, `/dev/urandom`
- Atmospheric noise
- Radioactive decay
- Lava lamps

Use-cases for random numbers:

- Cryptography
- Slot machines
- Monte-Carlo and other simulation
- Machine Learning (initial weight, input to generative networks)
- Video Games
- Pinball
- VFX/Graphics
- Audio

for each use case, there are some different concerns, resource usage, portability, how much randomness is really required...

a six sided die, a basic (and bad) example.

```cpp
uint8_t d6(){
  return std::rand() % 6u + 1u;
}
```

`std::rand` is not thread-safe. and the modulo distribution doesn't give us truly uniform distribution.

better example, using a random device to get a seed, creating a generator. (note: some of the random generators have state). for unit testing, we will have a way to expose the seed and initialize it

```cpp
uint8_t d6(){
 static std::mt19937 g{std::random_device{}()};
 static std::uniform_int_distribution<uint8_t> d(1,6);
 return d(g);
}

// and as a class
class D6 {
  public:
  D6() : m_gen{std::random_device{}()} {}
  uint8_t Roll() {
    return m_dist(m_gen);
  }

  private:
    std::mt19937 m_gen;
    std::uniform_int_distribution<uint8_t> m_dist{1u, 6u};
};
```

(**Mersenne Prime**)

how to seed the generator correctly? there are some numbers that can't be produced as the first result of simple seeds.

seeding strategies:

- Random device
- `std::seed_seq`
- Time
- Sequential
- Jumps

34d party generators libraries, most of which have options for 32bit and 64bit seeds,

- Xoshiro, Xoroshiro
- PCG
- Random123/Philox (proposed for the standard)

#### Demo

(listening to some audio)

#### Benchmarks

Xoshiro is really fast, even more than simple, less quality generators.
starting with the "scratch" generator (simulate a scratch on a vinyl record), we want to see that the "scratches" don't appear again and again at the same locations. also white-noise generators, dust generators, etc...\
Checking Unit tests and end-to-end profiling. there's actually not much difference at the end.

</details>

### GPU Performance Portability Using Standard C++ with SYCL - Hugh Delaney & Rod Burns

<details>
<summary>
GPU abstraction and portability without compromising on performance.
</summary>

[GPU Performance Portability Using Standard C++ with SYCL](https://youtu.be/8Cs_uI-O51s)

this talk is for C++ programmers who write for GPUs and want portability, without losing on performance.

> "SYCL is a C++ interface enabling heterogenous acceleration"

pure modern C++, no weird stuff going around, part of the open standard, used to run single source code on AMD, Nvidia, Intel GPUs and more. performance is comparable to native APIs.

it uses exception, lambdas, templated, CTAD. exceptions aren't usually supported by devices, but exceptions can be propagated upwards. single source code means that the same code runs on the main device (CPU) and the periphery (GPU).
Kernel code is compiled for any platform passed to the compiler, and the device code is bundled within the final binary (this does make the binary larger).

trade off between specialization and abstraction, writing optimized code for one feature or one device inhibits using the same code at other scenarios.

</details>

### LLVM Optimization Remarks - Ofek Shilon

<details>
<summary>
Looking into optimization remarks and diagnostics.
</summary>

[LLVM Optimization Remarks](https://youtu.be/qmEsx4MbKoc), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/optRemarks_cppcon.pptx).

Optimization Remarks, or Optimization Diagnostics, are logs about optimizations which were attempted but failed.

#### Getting Optimization Remarks

the `-Rpass` flag dumps the optimization remarks into the output stream, which is very hard to read.then there was the _llvm-opt-report_ tool, next _opt-viewer_ tool which presents in html. this is part of the llvm master branch. to get the data, we add the `-fsave-optimization-record` flag to the build switch, which would create the _\*.opt.yaml_ files, which are then fed to the _opy-viewer_ tool.

```shell
opt-viewer.py --output-dir <html folder> --source-dir <repo> <yaml files folder>
```

opt-viewer can show "Hotness" scores (if compiled with PGO) to help focus on the important points, it can also show inlining context.

_OptView2_, which is more advanced, has smaller memory footprint, and tries to narrow down the results to provide actionable items for developers (rather than just for compiler authors). it is expected to be integrated into LLVM in the coming months.

it is also possible to add support for optimizations in compiler explorer.

#### Understanding and Using the Optimization Remarks

inlining: unavailable defintion (not in header), implementation too long.

"clobbered by store"

possible **aliasing**

```cpp
void foo(int* a, const int& b){
  for (int i =0; i<10; i++) {
    a[i] += b;
  }
}
```

in this case, we are worried that b is an alias of an element inside a, so the compiler must get the value of b each time. we could get around this by using the \_\_\_restrict\_\_\_ keyword.

```cpp
void foo(int* __restrict__ a, const int& b){
  for (int i =0; i<10; i++) {
    a[i] += b;
  }
}
```

a different thing would happen if we changed the type of the array. now the compiler generates the better code right away, this is because of **strict aliasing**. which is a permission given to the compiler to assume that object of different types will never refer to the same location. however, some compilers still don't support this.

```cpp
void foo(long* a, const int& b){
  for (int i =0; i<16; i++) {
    a[i] += b;
  }
}
```

"clobbered by call"

```cpp
void someFunc(const int&);
int otherFunc();
void f(int i, int* res){
  someFunc(i);
  i++;
  res[0] = otherFunc();
  i++;
  res[1] = otherFunc();
  i++;
  res[2] = otherFunc();
}
```

in this example, i's address is stored in the `someFunc()`, where it might be existing or used in a different function, so we can't eliminate the variable, and we can't simply increment it at one go.
we also would again want the value of `otherFunc()` to be commited to a register, but without knowing the function body, there is no guarantee it doesn't have side effects and that it always returns the same value (I.E, it could be a counter).

we can get around this by using the _pure_ attribute to mark the function as not modifying global state. but this is cheating, because if it doesn't have side effects and doesn't return anything, then it could be eliminated entirely. there is also the _const_ attribute, which is a stronger restriction, that says the function doesn't even read global memory.

```cpp
void someFunc(const int&) __attribute__((pure));
int otherFunc() __attribute__((const));
void f(int i, int* res){
  someFunc(i);
  i++;
  res[0] = otherFunc();
  i++;
  res[1] = otherFunc();
  i++;
  res[2] = otherFunc();
}
```

there is also the _noescape_ attribute of the function argument, which implies that address of the argument does bot escape the function call.

```cpp
void someFunc(const int& __attribute__((noescape)));
```

if instead of using _i_ as the argument to the internal function, we chose to use _+i_, then the argument is a temporary object, so there is no address to escape from the internal function.

sometimes the offending function that inhibits the optimization is from the standard library, such as the output operator `<<`.

```cpp
#include <fstream>
void otherFunc();

void f(int i){
  std::ofstream fs("myFile");
  fs << &i;
  i++;
  otherFunc();
  i++;
  otherFunc();
  i++;
  otherFunc();
}
```

"Failed to move load with loop invariant address"

we would like to "hoist" the result of m_cond outside the loop, and only check it once. but the `f` could still modify the member variable.

```cpp
class C {
  bool m_cond;
  void method1();
};

void f();

void C::method1(){
  for (int i=0; i<5; ++i>) {
    if (m_cond) {
      f();
    }
  }
}
```

(an example of changing a member variable), `const` doesn't mean much to the compiler, and it can go over it.

```cpp
struct C {
  int m_i;
  int * m_p = &m_i;
  void constMethod() const{
    ++(*m_p); // m_i is modified.
  }
};
```

| Symptom                              | Probable Cause    | Action                                                                      |
| ------------------------------------ | ----------------- | --------------------------------------------------------------------------- |
| Inlining Failure                     |                   | add header, force inlining, increase threshold                              |
| "Clobbered by store"                 | Aliasing          | restrict, force type difference                                             |
| "Clobbered by call"                  | escape            | attributes(pure, const, noescape), typically before the remark site         |
| "Failed to move load loop invariant" | Escape            | all of the above + copy to local                                            |
| \*                                   | don't understand? | Reduce to bare minimum in compiler explorer, might be a compiler limitation |

#### Beyond Classical Clang Toolchain

LTO builds - link time optimization. the optimizations are performed at the linker state, so using the regular flag creates bogus optimization remarks and does not create the real files. we can use `llvm-lto -lto-pass-remarks-output=<yaml output path> -j=10 -O=3 <obj files list>` instead, but it isn't a well maintained tool.

gcc accepts the flag, but works in a different way, and is mostly dead.

Decorations across compilers: all compilers struggle with aliasing.
| clang | gcc | icc | msvc |
| --------------------------- | ------- | ------- | ----------------------------------------- |
| `__restrict` | &check; | &check; | `__restrict` (pertains to locals), `__declspec(restrict)`(decorates a function return value) |
| `__attribute__((pure))` | &check; | &empty; | &empty; |
| `__attribute__((const))` | &check; | &check; | `__declspec(noalias)` |
| `__attribute__((noescape))` | &empty; | &empty; | &empty; |

rust works really well to combat aliasing (borrow-checker), it's really hard to generate aliasing code in rust.

```rust
pub fn foo2(a:&mut[i32; 10], b:&i32){
  for i in 0..10{
    a[i] += *b;
  }
}
```

carbon might help with escape analysis. if we don't decorate the function with _var_, then we say the variable doesn't have an address and so the address can't escape.

```carbon
fn someFunc(i:i32); // no prefixing with 'var'
fn whatEver()->i32;

fn f(var i:i32, var res: [i32; 3]){
  someFunc(i);
  i = i+1;
  res[0] = whatEver();
  i = i+1;
  res[1] = whatEver();
  i = i+1;
  res[2] = whatEver();
}
```

Recommendations:

- concentrate on known bottlenecks
- invest when you:
  - work at sub-millisecond scale
  - very tight loops

</details>

### GPU Accelerated Computing & Optimizations on Cross-Vendor Graphics Cards with Vulkan & Kompute - Alejandro Saucedo

<details>
<summary>
Introduction to Vulkan and Kompute for parallel processing
</summary>

[GPU Accelerated Computing & Optimizations on Cross-Vendor Graphics Cards with Vulkan & Kompute](https://youtu.be/RT-g1LtiYYU), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/CppCon-2022-GPU-Computing-Made-Simple-with-the-C-Vulkan-SDK-the-C-Kompute-Framework-AMD-Qualcomm-NVIDIA-Friends.pdf)

#### Parallel Processing

> Why Parallel Processing?
>
> - Functions can often be reduced to highly parallelizable stages (Matrix Mult, ML Layers, etc)
> - Micro-batching allows for further parallelization of multiple inputs (eg. cost instead of loss)
> - Breaking up fractions of each ensemble component across tightly coupled hardware (eg. multi-GPU)

phyiscal differences between CPU and GPU, different compute models, memory model - copying from host to device and back.

there are now many vendor with many models of graphics card, and tools that make using all those options standardized (openGL, openCL, Cuda, etc...). \
Vulkan is an opensource SDK by khronos. it provides low level rich access to components at the GPU level, using a C-style API as a core interface, and is highly compatible across platforms, mobile devices and suppliers.

Architecture:

- Application
- Instance (vulkan can have multiple instances of application)
- Phyiscal Device (the underlying GPU card hardware)
- Logical Device (abstraction of physical devices with set properties)
- Queue - Command buffers (work gets submitted into the queue and is executued in the GPU)

graphics and compute pipelines, using shaders (code that runs in the GPU to perform any relevant processing). descriptor sets define types of native data types. this talk deals with compute pipelines only.

Compute pipeline:

- Shade Module
  - Compute State
- Pipeline Layout
  - Descriptor Set Layouts
  - Push Constants

> Different steps of a vulkan program
>
> - Setup Phyiscal and logical device
> - Load and compile Shader code
> - Create Compute pipeline
> - Copy data to host visible memory
> - Record Commands
> - Copy Data to GPU only memory
> - Bind pipeline (+ shader)
> - Bind descriptor sets (data)
> - Dispatch command
> - Copy data from GPU only memory
> - Copy output to host & print data

#### Kompute Framework

> Enter Kompute: The General Purpose Vulkan Computing Framework.
>
> - Dozens instead of thousands of lines of code required.
> - Augments Vulkan interface instead of abstracting it.
> - BYOV: Bring-your-own-Vulkan design to play nice with existing Vulkan applications.
> - Non-Vulkan name convention to disambiguate components.

part of the linux foundation, in the incubation stage.

Kompute Components

- Kompute Manager - (top level manages device and queue)
- Kompute Sequence (manages and execute operations as batch)
- Kompute Operation (abstraction of a instruction with tensor and shader)
  - Kompute Tensor (core data unit)
  - Kompute Algorithm (compute pipeline, shader modules and descriptor set binding)

```cpp
// 1. Create Kompute Manager
kp::Manager mgr;

// 2. Create and initialise Kompute Tensors through manager
auto tensorInA = mgr.tensor({ 2., 2., 2. });
auto tensorInB = mgr.tensorT<float>({ 1., 2., 3. });
auto tensorOut = mgr.tensorT<float>({ 0., 0., 0. }); //output
auto params = { tensorInA, tensorInB, tensorOut };

static std::vector<uint32_t> shader = compileShader(R"(
 #version 450

 layout (local_size_x = 1) in;

 // The input tensors bind index is relative
 layout(binding = 0) buffer bina { float tina[]; };
 layout(binding = 1) buffer binb { float tinb[]; };
 layout(binding = 2) buffer bout { float tout[]; };

 void main() {
 uint index = gl_GlobalInvocationID.x;
 tout[index] = tina[index] * tinb[index];
 }
)"); // GPU Code (shader) in glsl language

// 3. Run operation synchronously
auto algo = mgr.algorithm(params, shader);

// 4. Copy Tensor and execute algorithm
mgr.sequence()
 ->record<kp::OpTensorSyncDevice>(params); // copy from host to device
 ->record<kp::OpAlgoDispatch>(algo); // dispatch algorithm
 ->record<kp::OpTensorSyncLocal>(params) // copy from device to host
 ->eval();

// Prints the output which is Output: { 2, 4, 6 }
for (const float& elem : tensorOut->data())
 std::cout << elem << " ";
```

> Deeper Optimizations
>
> - run a single command/operation in sequence with manager
> - ruse multiple sequences in same tensors with pre-recorded cmds
> - asynchronous execution of sequences
> - concurrent execution of sequences across GPU queues

hardware parallel, multiple queues in concurrent execution

```cpp
// Kompute Manager with custom settings
uint32_t deviceId(1);
std::vector<uint32_t> queues({ 0, 2 });

kp::Manager mgr(deviceId, queues);

// Create parameters to use for each computation
std::vector<std::shared_ptr<kp::Tensor>> paramsA = { ... };
std::vector<std::shared_ptr<kp::Tensor>> paramsB = { ... };

// Create seq on relative index - name sequence
auto sq1 = mgr.sequence(0);
auto sq2 = mgr.sequence(1);

// Create seq on relative index
auto algo1 = mgr.algorithm(paramsA, shader);
auto algo2 = mgr.algorithm(paramsB, shader);

sq1->evalAsync<kp::OpaAlgoDispatch>(algo1); // no waiting
sq2->evalAsync<kp::OpaAlgoDispatch>(algo2);

sq1->evalAwait(); // now waiting
sq2->evalAwait();
```

</details>

### Fast C++ by using SIMD Types with Generic Lambdas and Filters - Andrew Drakeford

<details>
<summary>
DR-cubed - A wrapper library for SIMD operations
</summary>

[Fast C++ by using SIMD Types with Generic Lambdas and Filters](https://youtu.be/sQvlPHuE9KY), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/11/mainPresent_0.2d.pdf). [github](https://github.com/andyD123/Dr3)

SIMD - Single instruction, Multiple Data.

intrinsics, special compiler intrinsic for data and operation that know about registers and represent SIMD data and instructions. such as `_m256d` (data type that holds 4 doubles) and `_mm256_ad_pd(_m256d a, _m256d b)` (element-wise adding two of those types). to get more readable code, we can wrap those types into classes, such such _boost simd_ and _std::simd_ and the VCL (vector class library).

DR<sup>3</sup> (DR cubed) is another level of abstraction.

- Large Vector Type
- Lambda utilities
- Filters & Views

> Vector VecXX -
>
> - Memory managed vector type.
> - Supports math functions and operations.
> - Contiguous, aligned and padded.
> - Can change the scalar type and instruction set.
> - Substitutable for scalar type so we drop into existing code to make it vectorized.

[implementing **Black Scholes** model](https://en.wikipedia.org/wiki/Black%E2%80%93Scholes_model)

$$
\begin{align*}
& C(S_t,t)=N(d_1)S_t-N(d_2)Ke^{-r(T-t)} \\\
& d_1=\frac{1}{\sigma\sqrt{T-t}}[\ln(\frac{S_t}{K})+(r+\frac{\sigma^2}{2})+(T-_t)] \\\
&d_2=d_1-\sigma\sqrt{T-t}
\end{align*}
$$

(video)

> The problem: low intensity operations define the interface.

and the answer is generic lambdas, but this tends to have an effect on performance. which is where DR-cubed comes.

- Transform
- Branching
- Filters And Views

#### Performance and Simplicity

- transform - memcpy
- reduce - max element
- transformReduce - sum of squared using inner product

Composing - joining lambda together

**Filters And Views**

the lambda output is a boolean determining if the value is pipelined to the next lambda.

**Branching**

- select
- transformSelect
- filterTransform

we can either filter the data and run the true transformation on the elements which passed the condition (and the false transformation on those who don't), or run both lambda on all cases and merge the results using a mask - this is branchless.

experiment with different branching strategies on a real problem: Inverse cumulative normal distribution: $\Phi^{-1}$

- complex single pass
- sparse update
- filter-to-views
- transform-filter, transform-write
</details>

### C++ Algorithmic Complexity, Data Locality, Parallelism, Compiler Optimizations, & Some Concurrency - Avi Lachmish

<details>
<summary>
more thoughts about cache friendly code.
</summary>

[C++ Algorithmic Complexity, Data Locality, Parallelism, Compiler Optimizations, & Some Concurrency](https://youtu.be/0iXRRCnurvo), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/10/efficiency-is-usually-the-name-of-the-game-final.pptx)

cache friendly matrix traversal, column-wise or row-wise. row by row is much faster.

caches hold less memory than RAM, but are much faster to access.

types of caches:

- Data (D-cache, D$)
- Instructions (I-cache I$)
- Translation lookaside buffer (TLB)

there are usually layers of caches, starting with L1 (fastest), they become slower as the size increases. usually each core has separate L1 and L2 caches, and they all share the L3 cache. if the data isn't in either of the caches, then the main memory is accessed.

the specs on the demo machine are
| level | cycles | size | ---
| ----------- | ------ | ---- | ---|
| L1 | 2-4 | 32KB D-cache, 32KB I-cache | per core
| L2 | ~10 | 256KB (data and instructions) | per core
| L3 | ~4o | 8MB (data and instructions) | shared between cores
| Main memory | ~100 |virtual memory | shared between cores

when we use the cache, we always work with cache lines, even if we take one bit from it,a cache line is usually 64 bytes long. A cache can be fully associative - every line from the main memory can fit any location in the cache. i can also be partially associative, so each line of main memory can only fit some locations in cache. it's a tradeoff between flexibility and lookup speed.

with our example from before, when we use row-wise traversal, he first element is a cache miss, but the later elements are already in the cache until the entire cacheline is consumed. with a column-wise traversal, we have much more cache misses.

There is also the "pre-fetcher", which tries to speculate on how we access data and bring the next line from the main memory.

#### Cache misses

> - **Cold Miss** - The first time a cache block is accessed.
> - **Capacity Miss** - The accessed memory was recently evicted due to cache being full and the eviction would have occurred even if the cache was fully associative.
> - **Conflict Miss** - Too many cache lines in the same set in the cache. The cache line would not have been evicted with a fully associative cache.
> - **Sharing Miss** - Other processors are accessing the same data on the cache line (parallelism).
>   - True Sharing Miss: 2 processors are accessing the _same data_ on the cache line
>   - False Sharing Miss: 2 processors are accessing _different data_ on the cache line

a writing thread invalidates the cacheline for all other threads, and forces them to fetch the line again from the main memory.

##### Conflict Miss

assuming

> - _w_ = Word width = 64
> - _M_ = L1 Cache size = 32K (32768)
> - _B_ = Line (Block) size = 64 Bytes
> - _K_ = 8-way associativity = 8
>
> therefore:
>
> - tag = $w - \lg(M/K)$ = $64 -\lg(32768/8)$ = 52
> - set = $\lg(m/(K*B))$ = $\lg(32768/(8*64))$ = 6
> - offset = $\lg(B)$ = $\lg(64)$ = 6

depending on how we do the traversal, we might change the tag, set and offset, which could require cache eviction.

##### False Sharing Misses

both cores access different pieces of data, but it's on the same cache line. so if one process writes, then the entire cacheline is invalidated and the other process must fetch the line again from the main memory.

Pre-fetcher

- sequential (taking each element one by one)
- stride (repeating steps, each time taking the element at the same distance from the previous one)
- random access

pointers are horrible for fetching data. it always means that there will be another memory access. it's better to replace virtual inheritance with <cpp>std::variant</cpp> if possible.

another problem can be with large data, if we have an array of big objects, and we only look at one field, then we effectively load all the data for nothing (like stride access)

```cpp
std::find_if(data.begin(), data.end(), [&](const MyData& d){ return d.id == id;});
```

it would be better to replace the array of structs with a structs of arrays.

```cpp
std::find_if(data.ids.begin(), data.ids.end(), [&](const auto& dataId){ return dataId == id;});
```

we should choose cache-friendly containers.

code is also cached, same caching behavior applies to instructions as it applies to data. so if we can re-arrange our branches to be together (green flow vs error flow) to avoid memory operations.

#### Design Paradigms

Object oriented Design: when we have virtual inheritance, we suffer from having to call the data through the pointer (cache miss), and we also call different code for each type, so we evict the code from the instruction cache. if the data was sorted according to type, then at least the same functions would have been in the cache.

> - Cache-unfriendly indirection due to <cpp>std::unique_ptr</cpp>.
> - Wasted cache line space.
> - Function pointers in hot loop.
> - Invisible entities are processed as well.

with a Data oriented Design, we could change the code to store relevant data together, and avoid accessing through pointers.

> - Group data based on how it is actually going to be used.
> - Implement the common task, not the special case.
> - Eliminate booleans and object states by making them implicit.

unfortunately, writing Data oriented design is clunky and has poor readability.

</details>

### Take Advantage of All the MIPS - SYCL & C++ - Wong, Delaney, Keryell, Liber, Chlanda

<details>
<summary>
SYCL as a single language for multiple targets.
</summary>

[Take Advantage of All the MIPS - SYCL & C++](https://youtu.be/ZxNBber1GOs).

we want a single language that takes advantages of all the speed in the machine, not switching between languages depending on the device: CPU, GPU, vector unit.\
our language must support an open interface, able to interact across different architectures and have the programming models which are stable and will continue to work even as things change.

SYCL is one such language. it has also experience massive growth over the past years. there is still room to improve and evolve.

There are a few SCYL Implementations: DPC by Intel, ComputeCPP by codePlay, and hipSYCL by the Heidelberg university. depending on which machine architecture is used, then a different implementation (or even a custom one) is needed. building for a different target just requires changing the compiler flag.

Cooperation between SYCL and OpenAPI, bringing the language to HPC and super computers across the world.

#### SYCL Hello World

running SYCL on compiler explorer. single code source for device and host code.

create data, provide it into a buffer, using a queue to wrap over the device and command groups.

this generated code can be profiled and analyzed, if we have existing backend code (such of for NVIDIA and CUDA) then we simply use them. as an example, we change the code to mess up with cache locality, and we see the performance degradation. also looking at how inlining gets better speed.

(SYCL LLVM Implementation, using backend specific objects from top level)

<cpp>mdspan</cpp> - a non-owning multidimensional array view for C++23. a bit of refresher about class template argument deduction. mdspan will also get into SYCL eventually.

</details>

##

[Main](README.md)
