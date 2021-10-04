<!--
ignore these words in spell check for this file
// cSpell:ignore nlex heapify
-->

# Easy to Advanced Data Structures

[Easy to Advanced Data Structures](https://www.udemy.com/course/introduction-to-data-structures/) \
[Source files](https://github.com/williamfiset/Algorithms)

## Introduction

<details>
<summary>
what is a data structure? How do we compare them? What is the Big O notation
</summary>

> - A data structure is a way of organizing data so that it can be used effectively
> - they are essential ingratiates in creating fast and powerfull algorithms
> - they help to manage and organize data
> - they make code cleaner and easier to understand

**ADT** - abstract data type:\
an abstraction of a data structure that only provides the interface that a concrete data structure must uphold. it defines how the data structure should behave and what methods it has, but not how.

| Abstraction (**ADT**) | Implementation (**DS**)                                       |
| --------------------- | ------------------------------------------------------------- |
| List                  | Dynamic array, linked List                                    |
| Queue                 | Linked list based queue, Array based queue, stack based queue |
| Map                   | Tree map, Hash map, Hash table                                |
| Vehicle               | Golf cart, Bicycle, smart car                                 |

#### Understanding Space-Time Complexity

when we compare data structures as programmers, we ask ourselves

> - "how mush **time** does this algorithm need to finish?"
> - "how much **space** does this algorithm need to finish?"

the standard notations we use is the BIG O notations.\
there is also big theta notation, but it's less important

> "Big-O Notation gives an upper bound of the complexity in the **worst** case, helping to quantify performance as the input size becomes **arbitrarily large**"

for example, in the linked list, the worst case for search is when the searched element is in the end of the list, so the complexity is linear. same for space. we only care about arbitrary large input, the rest is meaningless in theory, so factors, constants and such are ignored.

these are the common complexcities, n being the input size.

| Complexicy        | Notation     | Notes         |
| ----------------- | ------------ | ------------- |
| Constant          | O(1)         |               |
| Logarithmic       | O(log(n))    | binary search |
| Linear            | O(n)         |               |
| Liner-logarithmic | O(n log (n)) | merge sort    |
| Quadric           | O(n^2)       |               |
| Cubic             | O(n^3)       |               |
| Exponential       | O(b^n), b>1  |               |
| Factorial         | O(n!)        |               |

> "Let _f_ be a function that describes the running time of a particular algorithm for an input size _n_"

we only care about the highest exponenet size, we can ignore constants, either for additions or multiplications, in theory, it doesn't matter, because the higher order power (term) will be dominant.

$
f(n) = 7log(n)^3 + 15n^2 +2n^3 +8 \\
O(f(n))= O(n^3)
$

examples of constants time running algorithms, linear time, quadric time. logarithmic time (binary search),

- finding all subsets of a set - O(n^2)
- finding all permutations of a string -O(n!)
- sorting with mergesort - O(n log(n))
- Iterating over all the cells in a matrix with n rows and m columns -O(n\*m)

</details>

## Static and Dynamic Arrays

<details>
<summary>
Continues blocks of data storing elements which we can accesse by elements.
</summary>

The array is probably the most used data structure.

> "A static array is a fixed length container containing n elements indexable from the range \[0,n-1]"

indexable mean that each slot/index (element) can be referenced with a number. static arrays are contentious slice of memory,all the elements are adjacent.

> - Used for storing and accessing sequentail data
> - Temporarily storing objects
> - Used by IO routines as buffer
> - Lookup tables and inverse lookup tables
> - Can be used to return multiple values from a function
> - Used in dynamic programming to cache answer to subproblems

access time is constant, because they are indexable. searching is linear time (we might need to traverse everyting), we can't insert /append or delete from static arrays. in dynamic arrayes, insertion and deletion is linear (we might need to shift everything), appending is constant time (assuming their is space available).

elements in array are postion based, and we use zero base indexing. elements can be iterated over, we usually use the square brackets notation "\[index]" to denote indexing. using a negative index or an index n or larger, we should get an out of bounds exception.

dynamic arrays are also indexed, but they also offer mechanics for growing and shrinking in size, we can add elements (to the end) or insert them(at any postion), and even remove them (requires reordering the array, or even shrinking it).

wc can implement a dynamic array with an underlying static array, when we need to add more elements, we create a new static array and copy all elements.

java source code implementation. the only intersting part is the implementation of the iterator.

</details>

## Linked Lists

<details>
<summary>
A sequential list of nodes that hold data that point to other nodes that also contain data.
</summary>

Each node contains data and a reference to the next node, the last node has a reference to null. used for separate chaining in hash tables and when implementing graphs.

terminology:

- head - the first node in a linked list
- tail - the last node in a linked list
- pointer - also called reference, tells us who the next node is
- node - a structure containing both the data and the pointer to the next node

A doubly linked list holds an additional pointer to the previous node, it makes traversal easier (going backwards is possible), but requires more space. doubly linked list also keep the tail exposed, so we can traverse it backwards, and removing nodes can be done in a constant time.

inserting - travel until we find where we want to insert the next node, stich it together.

```
inserted = new NODE(value)
inserted->next = current->next
current->next = inserted;
```

doubly linked list is similar, but with a lot more stiching to be done (we can't forget the 'previous' pointer)

```
inserted = new NODE(value)
inserted->next = current->next
inserted->prev = current
current->next = inserted;
inserted->next->prev = inserted
```

removing elements from a singly linked list, we can either use two pointers, with one lagging behind, so when the first pointer matches the node to remove, we can stich the nodes together, we can also use a trick that we use one travesal pointer, and when we find the element to remove, we swap the contents with the content of the next node and then we can safely remove the next on.

```
currentNode;
nodeBefore;
nodeBefore->next = curretNode->next;
delete currentNode;
```

removing a node from a doubly linked list is easier, once we find the node, we simply stich the previous and the next nodes together.

```
currentNode->prev->next = currentNode->next;
currentNode->next->prev = currentNode->prev;
delete currentNode;
```

in singly linked list, we can't elements from the tail easily, we need to reach it each time. in a doubly linked list we can always get the previous element so we can fix the tail.

java source code implementation. uses a NODE\<T> class. size is stored (not calculated). edge cases are removing when there is only one element. removing nodes by index (possible, just usually not exposed), another iterator implementation

</details>

## Stacks

<details>
<summary>
Last in, First out data structure.
</summary>

LIFO - last in, first out. push, pop, and peek.

> "one ended linear data structure which models a real world stack by having two primary operations, namely **push** and **pop**."

stacks are used in text editors, to undo operations, to keep track of matching brackets, used in programming to model recursion, using Depth First Search (DFS) on a graph.

getting the size is O(1), searching is O(N), because we need to pop all elements, and then push them back.

Example of using a stack to match brackets:\
every left bracket we find, we push to the stack, for a right brackets, we check if the stack not empty, and we check the top of the stack, if the top of the stack is the same type of the incoming bracket, we pop and continue. if they don't match then there is a problem, no need to check anymore. in the end we check that the stack is empty at the end of the operation.

Tower of Hanoi:\
moving elemens from stacks with constraints.

Stacks are usually implemented with arrays, linked list or double linked list. we have a head, and each time we add an element, we add it before the head, and it becomes the new head. popping an element is removing the head and setting the head to what is wasp pointing to. if we use an array then we keep track of the index of the last added element.

java source code implementation. uses a doubly linked list, nothing intresting.

</details>

## Queues

<details>
<summary>
First in, First out data structure. Push to the tail, pop from the head.
</summary>

FIFO - First in, first out.

> "Linear data structure which models a real world queues by having two primary operations, namely **enqueue** and **dequeue**."

we can enqueue (push element to the back), or dequeue (remove from front). terminology can be inconsistent, "enqueuing = adding = offering - pushing to the back", while "dequeue = polling = pop front = removing".

we always have a front and a back back end.

queues are used to model Breadth First Search (BFS) graph traversal, to keep track of a limited number of elements, to manage requests in order.

### Breadth First Search (BFS)

we have a graph, for every element, we (enqueue) push all the connections of the node into the queue, and then search from the front element (dequeue it) and repeat this until the queue is empty. in this pseudo code we modify the elements.

```pseudo
Let Q be a Queue
Q.enqueue(starting_model)
starting_node.visited = true
While Q is not Empty Do:
    node = Q.dequeue()
    For neighbor in neighbors(node):
        If neighbor as not been visited:
            neighbot.visited = true;
            Q.enqueue(neighbor)
```

### Implementations

We can implement queues with arrays (circular buffer?), or with linked lists. for a singly linked list, we add(enqueue) at the tail (back), and we pop (dequeue) from the head, each time we move the head forward.

java source code implementation. uses a doubly linked list, nothing intresting.

</details>

## Priority Queues

<details>
<summary>
Priority queue, Binary heap (a complete binary tree in a continues memory).
</summary>

Priority queues and heaps. Min and Max priority queues, binary heaps (sinking and swimming, sift down and sift up, bubble up and bubble down).

> "An ADT that operates similarly to a regular queue, except that **each elements has a certain priority**. The priority of the elements in the priority queue determines the order in which elements are removed from it"
>
> Priority queues can only hold elements that are comparable (support ordering, usually the 'less than' operator).

example with numbers, taking the smallest number.

we implement this ADT with a **heap**.

> - "A heap is a **tree** based DS that satisfies the **heap invariant** (the heap property): If A is a parent node of B then A is ordered with respect to B for all nodes A, B in the heap."
> - this means that the value of the parent node is always at the same relation to all of it's child nodes.
> - in a max heap, the parent node is always larger than than the child nodes, and in a min heap, always smaller. there is no defined relation between sibling nodes.

the heap helps us implement the priority queue. heaps aren't necessarily binary, they must be trees (can't contain cycles).

Priority queues usages:

- certain implementations of Dijkstras's shortest path algorithm.
- anytime we dynamically use 'get next best' or 'next worst'.
- Huffman encoding (lossless data compression).
- Best First Search algorithms such as **A\***.
- Minimum Spanning Tree algorithms/

complexity:

- construction is O(n) - from linear array, basis for heap sort.
- polling is O(log(n)) - take the root, might require reordering the heap to maintain the heap invariant.
- peeking is O(1) - without removing.
- adding is O(log(n)) - might need to reshuffle the heap.
- naive removing O(n)
- advanced removing (with hash table) O(log(n))
- naive contains O(n)
- advance contains (with has table) O(1)

as we can see, we can use a hash table to optmize some methods. this will be covered later.

### Turning Min Priority queue into a max Priority queue

most libraries provide just one of these two, either a max or a min priority queue. if we want the other, we need to make it ourselves. one way to hack this is to provide the comparable function and negate it. the other way is to negate the numbers before inserting them and when removing them. this works for signed numbers, not so much for classes or unsigned numbers.\
for strings, suppose we have compartoar _lex_,and it' negation _nlex_, s1 and s2 are strings, to get nlex we simply negate the value of lex (multiplying by -1).

### Inserting elements to binary heap

we use binary heaps for priority queues because it usually gives the best time complexity (better than linked list). There are many types of heaps(binary, fibonacci, binomial, pairing...).

> - "A binary is a binary tree that supports the heap in variant. in a binary tree, every node has exactly two children".
> - even leafs have exactly two children, its just that those children are null.
> - "A complete binary tree is a tree in which at every level, except possibly the last, is completely filled and all the nodes are as far left as possbile"

a canonical way of representing the complete tree haps is with an array. this gives us fast operations, as long as we maintain the structure of the complete binary tree.

- level 0 : index 0 (1)
- level 1 : index 1,2 (2)
- level 2 : index 3,4,5,6 (4)
- level 3 : index 7,8,9,10,11,12,13,14 (8)

> "let i be the parent node index:
>
> - left child index: 2i +1
> - right child index: 2i + 2
>
> (assuming zero based)"

- level 0 : 0 -> \[1,2]
- level 1 : 1 -> \[3,4], 2 -> \[5,6]
- level 2 : index 3 -> \[7,8], 4 -> \[9,10], 5 -> \[11,12], 6 -> \[13,14]

when we add nodes, we should manitain the heap invariants, we always add the new element at the lowest, first empty position, and from there we start bubbling up if needed. if the element is larger than parent, swap with parent, contniue to do so until we no longer violate the heap invariant.

### Removing elements from binary heap

removing the root is called polling, we don't need to search for the index, it's always the top element at index 0. to remove the root, we swap it with the last element index, remove the last element (which contained the previous root). and now that we are violating the heap invariant, we start bubbling down. we look at the children and swap with the smallest (prefring the left node), continue doing so.\
if we want to remove an element which isn't the root, we first search for it in the tree (linear search), we swap it with the last node, and do a bubble up again.\
we always work with swapping the last element and then bubbling up and down.

- polling is O(log(n)) - we know where the root is, and we do one operations per tree level.
- removing is O(n), we first search for the element, and then we perform the bubbling operations from that point.
- there is actually a better way to remove element.
  > "The inefficiency of the removal algorithm comes from the fact that we have to perform a linear search to find out where the an element is indexed at. What if instead we did a lookup using a _Hashtable_ to find out where a node is indexed at?\
  > A _hashtable_ provides constant time lookup and update for mapping from a key (the node value) to a value (the index)"

every value is mapped to the index, we can map the value to several indices with a set or tree set of indexes. at each bubbling operations we swap the values in both trees (data tree and index tree).

### Implementations

java implementations, using a comparable interface, heap size (last added index), heap capacity (which can grow), the '_heapify_' process (complexity of O(n)), we also have a map for the indices, which we use for checking if a heap contains an element, and for removing elements. swim is bubble up, sink is bubble down. swapping requires additional overhead for swapping the index map. when we remove from the middle of the heap we first check sinking, ad afterwards we check swimming.

</details>

## Union Find/Disjoint Set

## Binary Search Trees

## Hash Tables

## Fenwick Tree/Binary Indexed Tree

## AVL Tree

## Indexed Priority Queue

## Sparse Tables

## Complexity Table

| Data structure                  | Access        | Search | Insertion        | Appending            | Deletion                                   |
| ------------------------------- | ------------- | ------ | ---------------- | -------------------- | ------------------------------------------ |
| Static Array                    | O(1)          | O(n)   | N/A              | N/A                  | N/A                                        |
| Dynamic Array                   | O(1)          | O(n)   | O(n)             | O(n)                 | O(n)                                       |
| Singly Linked List              | N/A           | O(n)   | at head O(1)     | at tail O(1)         | at head O(1), at tail O(n), in middle O(n) |
| Doubly Linked List              | N/A           | O(n)   | at head O(1)     | at tail O(1)         | at head O(1), at tail O(1), in middle O(n) |
| Stack                           | peek top O(1) | O(N)   | N/A, only push   | push at top O(1)     | pop top O(1)                               |
| Queue                           | front O(1)    | O(N)   | N/A,only enqueue | enqueue at back O(1) | dequeue front O(1) in middle O(N)          |
| Priority Queue with Binary Heap | Peeking O(1)  | N/A    |                  | Adding O(log(n))     | Polling O(log(n))                          |
