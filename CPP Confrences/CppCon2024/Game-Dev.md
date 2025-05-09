<!--
// cSpell:ignore Bataille Royale Tetris pgetinker
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Game Dev

<summary>
7 Talks about game development.
</summary>

- [ ] A Simple Rollback System in C++: The Secret Behind Online Multiplayer Games - Elias Farhan
- [x] Blazing Trails: Building the World's Fastest GameBoy Emulator in Modern C++ - Tom Tesch
- [ ] Cross-Platform Determinism Out of the Box - Sherry Ignatchenko
- [x] Many ways to kill an Orc (or a Hero) - Patrice Roy
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
to solve this, we introduce a _free link_ into the container - when we remove objects, we store the index at the free list rather than remove the content. if we want to add an object, we first check the free list for available locations, and we only allocate memory if there's nothing in this list. (the list stores evictions). this way we don't have fragmentation.

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

however, we can't send everything to the GPU. and more than that, not all elements are even seen by the player. we should only render objects that are in the players field of view. this is called _Frustum Culling_.\
the naive implementation is to wrap the objects with a bounding volume and use math to determine if the objects intersects with the frustum. but since we have thousands of objects, we can't linearly search all elements each time, and we need a solution that scales.\
instead, we could use _QuadTrees_ and _OcTrees_, which are trees that subdivide the game space in a recursive way. we focus on _QuadTrees_ which are used for 2d games, and are simpler to understand. the level is the root node, and it subdivides the space into 4 quadrants. each node can either be a game object itself or an additional sub division. an object can be part of more than one quadrant.\
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

### Many ways to kill an Orc (or a Hero) - Patrice Roy

<details>
<summary>
Different ways to implement the same thing: either with inheritance or with variants.
</summary>

[Many ways to kill an Orc (or a Hero)](https://youtu.be/g-SnU4xRHMQ?si=daAZx6KOrnsEgh-5), [slides](<https://github.com/CppCon/CppCon2024/blob/main/Presentations/Many_Ways_to_Kill_an_Orc_(or_a_Hero).pdf>), [event](https://cppcon2024.sched.com/event/1gZg5/many-ways-to-kill-an-orc-or-a-hero).

starting with something simple, we have a base class for characters, and then sub classes for types of characters (orc and hero), and they hit one another.

```cpp
// the easy case
import std; // or #include <random>, <string>, <string_view>, <print>
class Character {
    std::string name_;
    int life_;
public:
    constexpr Character(std::string_view name, int life) : name_{ name }, life_{life } {}
    constexpr std::string name() const { return name_; }
    constexpr bool alive() const { return life() > 0; }
    constexpr bool dead() const { return !alive(); }
    constexpr int life() const { return life_; }
    constexpr void suffer(int damage) { life_ -= damage; }
};

class Orc;
class Hero : Character {
    int strength_;
public:
    constexpr Hero(std::string_view name, int life, int strength)
    : Character { name, life }, strength_{ strength } {}
    constexpr int strength() const { return strength_; }
    using Character::name, Character::alive, Character::dead,
    Character::life, Character::suffer;
    void hit(Orc&);
};
class Orc : Character {
    int strength_;
public:
    constexpr Orc(std::string_view name, int life, int strength)
    : Character { name, life }, strength_{ strength } {}
 constexpr int strength() const { return strength_; }
    using Character::name, Character::alive, Character::dead,
    Character::life, Character::suffer;
    void hit(Hero&);
};
void Hero::hit(Orc &orc) {
    orc.suffer(strength());
}
void Orc::hit(Hero &hero) {
    hero.suffer(strength());
}

int main() {
    std::mt19937 prng { std::random_device{}() };
    std::uniform_int_distribution dice{ 1, 100 };
    Orc orc{ "URG", 100, dice(prng) / 10 };
    Hero hero{ "William", 100, dice(prng) / 10 };
    while(orc.alive() && hero.alive())
    {
        if(dice(prng) % 2 == 0)
            hero.hit(orc);
        else
            orc.hit(hero);
    }
    std::print("{} won", orc.alive() ? orc.name() : hero.name());
}
```

we'll look at different ways to implements the same basic idea. in our simple implementation we used private inheritance, and exposed base class properties and methods with the <cpp>using</cpp> expression. we didn't use pointers, so we didn't have virtual functions. we can change our design to use inheritance, and that effects how the other classes behave and interact.\
for example, we might want to restrict which characters can hit and be hit by others, and we don't want to simply make patches every time.

> What do we want?
>
> - We want heroes to be able to hit monsters (Orcs in particular)
> - We want monsters to be able to hit heroes (otherwise it's unfair)
> - To make things fun, let's suppose heroes can wear armor, which potentially gives them an edge
>
> We have implementation considerations:
>
> - Let's make it so all characters will have a _name_, and _life points_
> - The same rules will apply for all characters regarding the ideas of being dead, or alive
> - There can be variations in the way characters name themselves
> - Open questions: should all characters be able to suffer? To heal? etc...
>
> We have interface concerns
>
> - Some characters are bellicose and can hit others
> - The way damage is dealt can vary with type, probably with individual objects too
> - The way damage is received can also vary with type, probably with individual objects (e.g.: when someone wears armor)

#### Inheritance Attempt

our first attempt to apply this design will have the following:

- Character: has name, life, state (dead, alive) and other properties
- Damageable: can take damage
- Bellicose (violent?): can inflict damage
- Hero: can wear armor
- Monster: speaks simply
- Armor: grants protection
- Orc: smelly and proud of it

this is a combination of inheritance (interfaces) and composition, but it might lead us to diamond inheritance.

> This solution uses well-known but intrusive techniques such as
> public inheritance:
>
> - It's not necessarily bad, but it introduces coupling in our code
> - We've learned over time that when alternatives exist that involve less coupling, they tend to help us perform code maintenance

inheritance makes objects bigger, it doesn't play well with <cpp>constexpr</cpp> functions, and complicates the shared base class.

#### Using Templates Attempt

we might want something less intrusive, we replace the interfaces of "Damageable" with a template, so if an character has the public method of taking damage, then we generate a template function to hit it. otherwise, we would get a compilation error when trying to hit a character that isn't damageable.

```cpp
// ...
class Hero : public Character {
    int strength_;
    std::unique_ptr<Armor> armor;
    void equip(std::unique_ptr<Armor> p) { armor = std::move(p); }
public:
    template <class Damageable> // note: public
    void hit(Damageable &other) {
        other.suffer(strength());
    }
    Hero(std::string_view name, int life, int strength) : Character { name, life }, strength_{ strength } { }
    Hero(std::string_view name, int life, int strength, std::unique_ptr<Armor> armor) : Hero{ name, life, strength } {
        equip(std::move(armor));
    }
    int strength() const { return strength_; }
    void suffer(int damage) {
        Character::suffer(armor? static_cast<int>(damage * armor->protection()): damage);
    }
};
```

there are upsides and downsides, we make the objects smaller and function calls are direct, but we can no longer have a vector of "damageable" objects, and when we try to attack a non-damageable characters, the error message will be horrible.

#### Clarifying Intent

maybe our code doesn't convey the intent properly without having a named type? that's what <cpp>concepts</cpp> are for! we define a concept that dictates a character can suffer, and then we have a concept for Bellicose characters. there are some problems with the intuitive syntax which forces us to have the test workaround.

```cpp
template <class T>
concept Damageable = requires(T &a) {
    a.suffer(std::declval<int>());
};
void bellicose_test(auto && d) {
    struct X { void suffer(int) {}; } x; d.hit(x);
}
template <class T> concept Bellicose = requires(T &a) {
    bellicose_test(a);
};

void attack(Bellicose auto &from, Damageable auto &to) {
    from.hit(to);
}
```

our code is now clearer, easier to understand and should have better error messages now. but still not way to have a vector of damageable characters.

#### Grouping Objects

we can't use the intuitive way of grouping objects without a shared base class, but we can use <cpp>std::variant</cpp>!

```cpp
// other Damageable types
struct Furniture {
 void suffer(int) { /* breaks easily */ }
};
struct Bystander : Character {
    using Character::Character;
};

template <Damageable ... Ts>
void attack(Bellicose auto &from, std::variant<Ts...> &to) {
    std::visit([&from](Damageable auto && to) { from.hit(to); }, to);
}
template <Damageable ... Ts>
std::vector<std::variant<Ts...>> make_victims(Ts &&... args) {
    return { std::forward<Ts>(args)... };
}

int main() {
    std::mt19937 prng { std::random_device{}() };
    std::uniform_int_distribution dice{ 1, 100 };
    auto victims = make_victims(
        Orc{ "URG", 100, dice(prng) / 10, Smell{ 0.7 } },
        Furniture{},
        Bystander{ "Fred", 20 }
    );
    Hero hero{ "William", 100, dice(prng) / 10 };
    // William swings his halberd!
    for(auto && p : victims)
        attack(hero, p);
}
```

there is a problem about uniqueness in the variant, but we can ignore it for now. but we are able to group the objects together now.

#### Comparing The Approaches

we can compare our approaches according to class size and execution speed. the basic approach with virtual base class had 48 bytes for the character class, 72 for the hero, 64 for the monster, and 72 for orcs. the size changes depending on the compiler,

| Metric    | Virtual Base Class | Variant |
| --------- | ------------------ | ------- |
| Character | 42                 | 48      |
| Hero      | 72                 | 56      |
| Monster   | 64                 | 48      |
| Orc       | 72                 | 56      |

using variants is also faster, but there are more considerations, such as compile time, portability and maintenance.

#### Creating A Bataille Royale

a team battle between a group of heroes and a group of orcs. we hide some stuff in a private implementation, and add some complexity with weapons and spells at the hero side. we also add some randomness.\
we need some changes to make sure hero can only attack monsters and vice-versa, and we need some driver code to choose which character attacks and which is attacked.
</details>

### Blazing Trails: Building the World's Fastest GameBoy Emulator in Modern C++ - Tom Tesch

<details>
<summary>
Showing some stuff about the GameBoy console and emulation.
</summary>

[Blazing Trails: Building the World's Fastest GameBoy Emulator in Modern C++](https://youtu.be/4lliFwe5_yg?si=iV6ueM2WeBo71gtC), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Blazing_Trails.pdf), [event](https://cppcon2024.sched.com/event/1gZgB/blazing-trails-building-the-worlds-fastest-gameboy-emulator-in-modern-c)

> Emulators require a strong grasp of hardware, low-level operations, and CPU architecture.Building one helps you understand how hardware like memory, registers, and buses work, and how they interact with machine code. This is valuable for a C++ programmer, as C++ often deals with system-level programming.

it's reasonably "easy" to write an emulator, and there are testing rom, so this is a project that's a relatively doable, there is enough documentation for the chip design and known layout of the memory mapping, there is a fully mapped instruction set, which includes information about T-cycles (tick cycle, machine cycle) and M-cycles(memory cycle). there are two pages of OP codes (instructions).

rather than implementing everything, we can choose a toy example: Tetris. this was the launch title and came bundled with every game. it only require M-cycle precision, uses 284 of the 501 opcodes, is small in size (doesn't require memory banking).

using [pgetinker](https://pgetinker.com) to run the game online, like compiler explorer. taking the sample game from [olcPixelGameEngine](https://github.com/OneLoneCoder/olcPixelGameEngine).

(more stuff)

</details>
