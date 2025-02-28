<!--
// cSpell:ignore
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Game Dev

<summary>
7 Talks about game development.
</summary>

- [ ] A Simple Rollback System in C++: The Secret Behind Online Multiplayer Games - Elias Farhan
- [ ] Blazing Trails: Building the World's Fastest GameBoy Emulator in Modern C++ - Tom Tesch
- [ ] Cross-Platform Determinism Out of the Box - Sherry Ignatchenko
- [ ] Many ways to kill an Orc (or a Hero) - Patrice Roy
- [ ] Techniques to Optimize Multithreaded Data Building During Game Development - Dominik Grabiec
- [ ] Using Modern C++ to Build XOffsetDataStructure: A Zero-Encoding and Zero-Decoding High-Performance Serialization Library in the Game Industry - Fanchen Su
- [x] Data Structures That Make Video Games Go Round - Al-Afiq Yeong

---

### Cross-Platform Floating-Point Determinism Out of the Box - Sherry Ignatchenko

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[Cross-Platform Floating-Point Determinism Out of the Box](https://youtu.be/7MatbTHGG6Q?si=nZg15xz28m8fDCCe), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Cross-Platform_Floating-Point_Determinism_Out_of_the_Box.pdf), [event](https://cppcon2024.sched.com/event/1gZgC/cross-platform-determinism-out-of-the-box)

> For multiplayer games, deterministic simulations have been a Holy Grail for a long while. Indeed, if we have perfectly deterministic simulations - we can simply pass all the inputs to all the clients, and rely on each client to produce the same results, reducing network traffic by orders of magnitude. While especially important for RTS games, all kinds of multiplayer games would benefit from it.\
> However, while determinism was achieved in practice for single-platform, it is known to be a next-to-impossible to achieve for cross-platform clients. We will discuss the (well-known) reasons for it first - and will proceed into discussing our approach to the solution (with our open-source lib actually providing some implementations).\
> This talk is important for multiplayer game devs - and for anybody who is interested in deterministic calculations.

</details>

### C++ Data Structures That Make Video Games Go Round - Al-Afiq Yeong

<details>
<summary>
Some special considerations for data structures in game design.
</summary>

[C++ Data Structures That Make Video Games Go Round](https://youtu.be/cGB3wT0U5Ao?si=yxQV9ICXe52XNdVc), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Data_Structures_That_Make_Video_Games_Go_Round.pdf), [event](https://cppcon2024.sched.com/event/1gZgx/data-structures-that-make-video-games-go-round)

> Modern video games are complex beasts that contains multiple systems interacting with one another storing, transferring and processing large sets of data in real time. While some data structures from the standard library such as the std::vector gets you by 90% of the time you need to store and process data, there will be the occasional 10% that requires a unique take.\
> This presentation aims to discuss the unique data structures that are commonly used in video games / game engines that caters to the occasional 10%. We will go over several systems outlining their requirements, constraints and present custom data structures that gets the job done.

games are complicated, and modern games are much more complicated than those of the past. multi threading, 3d graphics, using game engines and with development cycles measured in years.

games still have a common structure:

1. Game Initialization
2. I/O
3. GamePlay
4. Physics
5. Rendering
6. Audio

the first step is to initialize the game, setting up systems that will be used for the game.

> Initialization
>
> - Environment Variables.
> - Error Handler.
> - Job system.
> - Resource Registry.
> - Reads boot data.
> - Listens to kernel message pump.

#### Resource Registry

The resource registry stores, manages,manipulates and exposes game assets. these assets are textures, meshes, materials, animation, sounds and more.The data is taken from a resource bundle/pack that is associated with the game stage ("level").\
the naive data structure would be an unordered map, but it is not cache friendly, especially when the load factor is high. instead, most game companies use an "Open Addressing Hash Map".

> - Simplest implementation only require a contiguous block of memory.
> - Elements are more densely packed.
> - Better cache performance.
> - Collisions are resolved via probing. probing methods include:
>   - Linear Probing.
>   - Quadratic Probing.
>   - Double Hashing.
>   - Robin Hood Hashing.

we focus on robin hood hashing.

PSL - probe sequence length - how far away an object is from the place where would want to put it originally.

the algorithm performs swaps when inserting elements, ensuring that over time, the psl of all elements remains the same as other elements, and there isn't much variance.

> Core algorithm:
>
> 1. Every bucket contains metadata about it's probe sequence length, PSL and is defaulted to some value.
> 2. On insert, elements start with a PSL value of 0.
> 3. PSL of the inserted element is compared to the PSL of the bucket.
> 4. If bucket's PSL value is the default, simply insert the element and store it's PSL.
> 5. If the bucket's PSL value is less than the element's, swap.
> 6. Probe to the next bucket and increment the PSL value of the swapped element.
> 7. Repeat steps 5 and 6 until an empty bucket( bucket with default PSL value) is found.

the next option to consider is <cpp>std::vector</cpp>, since it has constant time for random access and for inserting and removing at the end. but it has the problems of memory reallocation, which messes up pointers and allocators. this doesn't fly with the resource registry, which uses the pointer addresses.\
we could try using <cpp>std::array</cpp>, and instead of copying data around, we simply allocate another chunk for the new data and append this chunk to the pool of arrays. <cpp>std::deque</cpp> also doesn't help us if we want to store the pointers to those chunks.\
to solve this, we introduce a *free link* into the container - when we remove objects, we store the index at the free list rather than remove the content. if we want to add an object, we first check the free list for available locations, and we only allocate memory if there's nothing in this list. (the list stores evictions). this way we don't have fragmentation.

the index that we store is a data structure with the page index (which pool of memory is used) and the offset inside the page, we could encapsulate this as an iterator. we also need a way to prevent access to destroyed objects, we call this flag "version", so for each possible location, we have a metadata information about which generation it is. if the indexing objects has a different version than the current version, then the object was removed.

this is actually the suggested <cpp>plf::colony</cpp> data struture.

> <cpp>plf::colony</cpp> Authored by Matthew Bentley. The same idea with some differences:
>
> - Newly allocated chunks is always twice as large.
> - Uses a skipfield to skip iteration on empty elements.
> - Freelist is embedded into freed slots in chunks.
> - No versioning for each slots in a chunk.
>
> Has all the benefits of the previously discussed container:
>
> - Pointer / iterator stability.
> - Memory block reuse.
> - O(1) on insert amortized.
> - O(1) on erase amortized (regardless of location).

skipfield allow iterating over empty blocks without branching, and work for both directions.

if this data structure will get into the standard libray, it will be called <cpp>std::hive</cpp>.

#### Simulation

> Imagine our hypothetical game is 3D with a vast world.
>
> - Multiple systems ticking,
> - Entity Component Systems (ECS)
> - Artificial Intelligence
> - Animation
> - Physics
> - GamePlay specific
>
> We've simulated everything and all that's left is to send them to the renderer.

however, we can't send everything to the GPU. and more than that, not all elements are even seen by the player. we should only render objects that are in the players field of view. this is called *Frustum Culling*.\
the naive implementation is to wrap the objects with a bounding volume and use math to determine if the objects intersects with the frustum. but since we have thousands of objects, we can't linearly search all elements each time, and we need a solution that scales.\
instead, we could use *QuadTrees* and *OcTrees*, which are trees that subdivide the game space in a recursive way. we focus on *QuadTrees* which are used for 2d games, and are simpler to understand. the level is the root node, and it subdivides the space into 4 quadrants. each node can either be a game object itself or an additional sub division. an object can be part of more than one quadrant.\
to determine which objects should appear, we traverse the tree and do some math stuff. this gives us $O(log_N)$ complexity for the search. but it isn't cache friendly, and has problems with SIMD and when running on the GPU.

The render now has a list of items to process, which is a whole lot of graphic works. each frame is a chain of "passes", which depend on one another. so we need another data structure for this.

> Core idea- Represent the frame as a Directed Acyclic Graph (DAG).
>
> - Each pass is represented as a vertex.
> - Vertices that are dependent on the output of another are connected by a directed edge.
> - A vertex cannot depend on itself (acyclic).
> - Once the graph is built, topological sort to get the order of execution.

DAG are also in build systems, package management chains, and other cases.

> A DAG needs to be topologically sorted, into a linear list. Vertices that has more incoming edges needs to be lower in the list.
>
> 1. Depth-first search, $O(N)$.
> 2. Kahn's algorithm, $O(N)$,
>
> Kahn's algorithm is preferred. DFS can't detect cycles. So how do we sort with Kahn's algorithm?

for sorting, we need two object: "queue" and "result buffer". we first add vertices that don't have incoming edges into the queue. since the queue isn't empty, we pop an object and move it into the result buffer, and we remove the outgoing edges from that objects, which means we have a smaller graph now. we can add the vertices that no longer have incoming edges into the queue. we pop the next element from the queue and repeat the process of shrinking the graph. at the end of the algorithm, the result buffer will have the passes which don't depend on any other pass in the start, and next we will have the passes which depend on the earlier passes, without any duplications or cycles.
</details>
