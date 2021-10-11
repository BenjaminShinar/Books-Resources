<!--
ignore these words in spell check for this file
// cSpell:ignore nlex heapify Kruskal Treap inorder
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

<details>
<summary>
The Union-Find, uses group of elements with a root for each group. Can get good performance by amortizing.
</summary>

> "Union Find is a data structure that keeps track of elements which are split into one or more **disjoint sets**. It has two primary operations: **find** and **union**."

- find - given an element, find which group it belongs to
- union - merge two group together

example with magnets:

- we label all the numbers and merge them together based on attraction (proximity).
- each round we unify magnets into groups, or groups with other group.
- eventually we get one group.
- (this looks like an HLM thing)

Union Find (**UF**) are used in Kruskal's algorithm, grid percolation, network connectivity, finding the least common ancestor in trees, image processing.

the complexity for construction is O(n) uses &alpha;(n), which is amortized constant time for most operations.

### Kruskal Algorithm

Kruskal's minimum spanning tree algorithm.

> "Given a graph G = (V,E), we want to find a **Minimum Spanning Tree** in the graph (it may not be unique). A minimum spanning tree is a subset of the edges which connect all vertices in the graph with the minimal total edge cost."

each edge/link (connection between vetices/nodes) has a cost, and we want to touch all the vertices with the minimal cost total.

steps:

1. sort edges by ascending edge weight.
2. walk through the sorted edges and look at the nodes that are connected, if they are already unified (belong in the same group), keep going, otherwise, unify them and include the edge.
3. the algorithm terminates when every edge has been processed or all vertices have been unified into one group.

### Union and Find operations.

> "To begin using Union Find, first construct a bijection (a mapping) between your object and the intgers in the range [0,n)"

note: this step is not necessary, but it will allow us to construct and array-based union-find.

we don't have a specific order to how to map object to numbers, it just needs to be one-to-one relation, we should store it a lookup table. \
we then create an array, where each element is the mapping index of an object.

example, 12 objects, mapped into A-L, for now, we don't use path compression

| Object     | E     | F     | I     | D     | C   | A     | J     | L     | G     | K     | B     | H     |
| ---------- | ----- | ----- | ----- | ----- | --- | ----- | ----- | ----- | ----- | ----- | ----- | ----- |
| Lookup     | 0     | 1     | 2     | 3     | 4   | 5     | 6     | 7     | 8     | 9     | 10    | 11    |
| Unions - 1 | 0     | **0** | **4** | **4** | 4   | **6** | 6     | **0** | 8     | **4** | **6** | 11    |
| Unions - 2 | **4** | 0     | 4     | 4     | 4   | 6     | **4** | 0     | **0** | 4     | 6     | **8** |

for some index i in the array,the parent is going to be what's inside i in the array.
to unify, we look at the root node, and we change the parents.

- Union(C,K) root nodes are 4,9, so one of them (C) becomes the parent.
- Union (F,E) root nodes are 0,1, so one of them (F) becomes the parent.
- Union (A,J) root nodes are 5,6, J becomes the parent
- Union (A,B) root nodes are 6,10, so we merge the smaller group, and J is the parent.
- Union (C,D) root nodes are 4,3, C is the bigger group, now parent
- Union (D,I) root nodes are 4,2, C is the bigger group, now parent
- Union (L,F) root nodes are 6,0, E is the parent
- Union (C,A) root nodes are 4,6, both are groups, C is larger, J's parent is C, but A,B still point to J,
- Union (A,B) root nodes are 4,4 no need to do anything
- Union (H,G) root nodes are 11,8, choose one as parent
- Union (H,F) root nodes are 8,0, E group is larger.
- Union (H,B) root nodes are 8,4, C group is larger

and now we are done.

> - Find: "To _find_ which componenet a particular element belongs to find the root of that componenet by following the parent nodes until a self loop is reached (a node who's parent is itself)."
> - Union: "To _unify_ two elements find which are the root nodes of each component and if the root nodes are different make one of the root nodes be the parent of the other."
>   we don't "un-union' elements, we only change the top root nodes, not all the children.
>   the number of componenets is the number of the roots remaning, the number of root nodes never increase.

so far, we don't support the nice amortized &alpha;(n) complexity.\
checking if H and B belong to the same group requires five hops, and this can get much worse. depending on how we grouped them together.
H->G->E->C \
B->J->C

### Path Compression

this is what makes the Union-Find have &alpha;(n) complexity.

everything point to the root node, at each search we change the root node, we should always have at most two levels.

| Object   | A   | B   | C   | D   | E   | F   | G   | H   | I   | J   |
| -------- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Lookup   | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   |
| Unions-1 | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   |

- Union(A,B)
- Union(C,D)
- Union(E,F)
- Union(G,H)
- Union(I,J)
- Union(J,G)
- Union(H,F)
- Union(A,C)
- Union(D,E)
- Union(G,B)
- Union(I,J)

regular mode:

> A->B->C->D->E \
> F->E \
> G->H->E \
> J->I->H-->>E

checking if J and A are in the same group will take a lot of jumps.

with path compression:

> C->D->E \
> A->D->E \
> B->E \
> G->E \
> J->E \
> I->E \
> F->E \
> H->E

the check is much shorter, because we reduced the levels between each node and the parent. it also becomes stable. at each iteration we compress the path, so subsequent checks become faster.

### Implementation

java source code, we track the parents of each element (root), we have an array of sizes for each component. we start with each elements having itself as a root and size of 1. when we call find, we first find the root regularly, and then we update the path from it to the parent and making all elements point to the root as their root element.\
the root nodes contain the sizes, once something has been merged, the 'size' of it becomes irreverent, and we would never look it up again. we use intgers rather than objects, it's easier to do so, we can simply keep the elements in a map and use the index number to find it in our arrays.

```java
public int find (int p){
  int root = p;
  while (root != id[root])
  {
    root = id[root]; // find root
  }
  while (p! =root)
  {
    int next = id[p]; //exchange and update
    id[p] = root;
    p= next;
  }
  return root;
}

public void unify(int p, int q){
  int root1 =find(p); //path compression happens here
  int root2 =find(q); //path compression happens here
  if (root1 == root2) return; //same group already
  if (sz[root1] < sz[root2]) //which root is larger
  {
    sz[root2]+= sz[root1]; // update sizes
    id[root1] = root2; //make one the root of the other
  }
  else
  {
    sz[root1]+= sz[root2];
    id[root2] = root1;
  }

  --numComponents; // decrease the number of components, because we merged two groups
}
```

</details>

## Binary Search Trees

<details>
<summary>
A tree with at most two child nodes, where the elements in the left sub tree are all smaller than the root node, and the elements in the right subtree are all larger than it.
</summary>

> "A tree is an undirected graph which satisfies any of the following defintions:
>
> - An acyclic connected graph (no cycles).
> - A connected graph with N nodes and N-1 edges.
> - A graph in which any two vertices are connected by _exactly_ one path"

a tree has a root node. any node can be a root.
a child is a node extending from another node. A parent node is the inverse of this. the root note has no parent (although it's sometimes useful to assign to sat that it's its' own parent, like file systems tree). a leaf node is a node with no children. subtree is a tree entirely containted inside a another tree. a single node is a subtree.

a binarty tree is a tree in which every node has at most two child nodes.

> "A **binary search tree** is a binary tree that satisfies the **BST invariant**: left subtree has smaller elements and right subtree has larger elements than of that of parent node."

duplicate values are sometimes allowed, but we usually don't need them.

usages

- Binary search tree (BSTs)
  - Implementation of some Map and Set ADTs
  - Red black Tree
  - AVL Tree
  - Splay Tree
  - etc...
- Implementain a binary heap
- Syntax tree (for the compiler and calculator)
- Treap - a probabilistic DS (a randomized BST)

the complexity of a binary search tree operations is O(log(n)) in the average case, but in the worst case (degenerate tree) we can have worse performance (linear complexicy).

[tree visualization tool](https://www.cs.usfca.edu/~galles/visualization/BST.html)

### Insertion

elements must be comparable, so that they could be ordered inside the tree.
we compare a the value of the inserted element to the current node, if larger, we move to the right node, if smaller, the left node,if the child node it that direction doesn't exist, we create a new node. for equality cases we can allow duplicates or not.

on random data, we get O(log(n)) complexity, but for sorted data we get linear complexity.

### Removal

removing elements is more complicated, but we can look at is a two steps process.

> 1.  Find the element we wish to remove (if it exists)\
>     when we search, there are four options.
>
>     1.  we hit a null node, which means the value does not exist within our tree.
>     2.  Compeator value equal to 0, element found
>     3.  Compeator value less than 0, the value can be the left subtree
>     4.  Compeator value less greater 0, the value can be the right subtree
>
> 2.  Replace the node we want to remove with it's successor (if any) to maintain the BST invariant.\
>     to remove a node, there are again four cases:
>     1. Node to remove is a Leaf node (no children). just remove it,don't forget to fix it from the parent (remove the reference).
>     2. node to remove has left subtree, but not right subtree. the successor of the node we remove will become the root of the subtree (don't forget to fix the parent link).
>     3. node to remove has right subtree, but not left subtree. same as the earlier case.
>     4. node to remove has both subtrees. the successor can "either the **largest value** in the _left subtree_, or the the **smallest value** in the _right subtree_".

> "the **largest value** in the _left subtree_ satisfies the BST invariant because it is larger then anything else in the left subtree, but smaller than anything in the right subtree, because it was found in the left subtree."

same logic applies for the **smallest value** in the _right subtree_.

we go once to one side, and then continue all the way with the other side.

we copy the value from the successor and overwrite the value we wanted to remove, now we perform the remove operation on the 'duplicated' successor element. the successor node will never have both children nodes, so we don't have much complications to worry about.

### Traversal

three types of traversals that are defined recursively
-PreOrder - current node is processed _before_ the children node
-InOrder - current node is processed _between_ children node
-PostOrder - current node is processed _after_ the children node
-Level Order Traversal - process layers by layer. uses breadth first search with a queue.

```
preorder(node):
  if node == null: return
  print(node.value)
  preorder(node.left)
  postorder(node.right)

inorder(node):
if node == null: return
  inorder(node.left)
  print(node.value)
  inorder(node.right)

postorder(node)
  if node == null: return
  preorder(node.left)
  postorder(node.right)
  print(node.value)
```

when we do **inorder** traversal and print, it means we print the elements in ascending order.

a level order travesal uses a queue, we do a breadth first search, we push the root to the queue, and then we start cycling until the queue is empty, we take the element from the queue, and push the children of that node into the queue.

```
let q = Queue
q.push(root)
while (!q.empty):
  let n = q.pop()
  process(n)
  q.push(n.left)
  q.push(n.right)

```

### Implementation

java source code. the data must extend the Comparable interface, some iterator stuff, as usual.

</details>

## Hash Tables

<details>
<summary>
Key-Value mapping. collision resolution schemes (separate chainning, open addressing)
</summary>

> "A Hash table is a data structure that provides a mapping form the key to the values using a technique called _hasing_."

key values pair, kays must be unique (values don't have to be), a key is mapped to a value. we can use Hash table to track frequencies.

we can construct a mapping between any type of keys and value, as long as the key is **hashable**.

### Hash Functions

> "A hash function _H(x)_ is a function that maps a key 'x' to a whole number in a fixed range."

the result of the hash funcion doesn't have to be unique,we can have hash functions for any type of key, not just numeric values(strings and objects). hash function can be simple or complicated

> "if _H(x) == H(y)_ then objects x,y **might be equal**, but if _H(x) != H(y)_ then x,y are **certainly not equal**."

this property means we can use hash functions to speed up comparison of objects, if we have their hash value, we can compare the hash value to see if they match or don't match, if they don't match, we don't need to compare the objects themselves, only if they match, we need to check the complete contents. for files, this is part of the checksum.

a hash function must be **deterministic**, _H(x)_ should always produce the same hash value, and never something else. we would like the function to be uniform, so that we minimize the number of _hash collision_, a hash collision is when two elements map to the same hash value.

> Q: "What makes a key of type T _hashable?"_
>
> A: "Since we are going to use hash functions in the implementaton of our hash tale we need our hash functions to be _deterministic_. To Enforce this behavior,we demand the the _keys used in our has table are immutable data types_. Hence, if a key of type T is immutable, and we have a hash function _H(k)_ defind for keys k of type T then we say a key of type T is hashable."

we would like a very fast insertion, lookup and removal time the data we are placing within our hash table. we can get this is O(1) complexcity by using a hash function as a way to index into as hash table. the O(1) applies only if we really do have a good unifrom hash function.

### Collision Resolutions

when two values are hashed to the same location, we get a hash collision. there are many techniques two solve this issue, we will focus on two of them

> - "**Separate chainging** deals with hash collisions by maintaining a data structure (usually a linked list) to hold all the different values which hashed to a particular value."
> - "**Open addressing** deals with hash collisions by finding another place within the hash table for the object to go by offsetting it from the position to which it hashed to."

### Separate Chaining

we store both the key and the value in the linked list block. when we insert and detect a collision, we push the new value into the contained data structure. for searching (finding), we use the hash function to detect where (at which bin/list) the element should be, and then we limit our search.\
if our hashtable becomes too big (with too many elements in each bin), and we want to maintain the constant time behavior, we can use a larger hash table (have more bins that keys can match into) with shorter chains, so we re-calculate the hash value for each element and insert it into the new hash table.\
we don't necessarily have to use a linked lists, we can use a tree (binary, self balancing), an array, or some hybrid combinantion (like Java HashMap, switching from linked list to tree), these might be more memory intensive then simple linked lists and harder to implement.

source code in java, we might want to cache the hash code (calculate once and store) so we don't repeat this calculation (especially for keys with linear time complexity to calculate the hash). we have a deafult capacity (number of bins) and some logic to decide when to increase the the capacity. our keys are created from a separate function that uses the hash value (which is an unconstrained number) and normalizes it into the range of the capacity. a method to resize the table.

### Open Addresssing

key &rarr; hash value &rarr; index.

in open addressing, they key-value pairs are stored in the table (the array) itself, and not in a separeate data structure. this means we really care about the the size of the hash table and how many elements are currently in it.

the load factor is denoted as &alpha;, and is the ration between the number of element in the table and the size (capacity).
$$ Load\ factor = \frac{Items\ in \ Table}{Size\ of\ Table}$$

once we have too many elements in the table, it becomes diffult to find empty places to add elements, and matching elements can also become less efficient.

> "The O(1) constant time behavior attributed to hash tables assumes the load factor(&alpha;) is kept below a certain fixed value. this means that once &alpha; > _threshold_ we need to grow the table size (ideally exponentially, e.g double)."

a graph from wikipedia the shows that once we reach a certain threshold of the load factor, the performance of the linear probing methods (such as open addressing) becomes much worse than that of separate chainning.

when we insert an element into an open addressing table, we compute the original hash value for the key and the required position, if the position is empty, great, we add the element there. if it's occupied, we use a probing sequence _P(x)_ and offset the current position, we keep doing this until we find an onoccupied spot.

there are an infinite number of probing sequences, such as:

- Linear probing $P(x) = ax+b$. a,b are constants.
- Quadratic probing $P(x) = ax^2+bx +c$. a,b,c are constants.
- Double Hashing $P(k,x) = X*H_2(k)$. H2 is c secondary hash function on the key.
- Pseudo random $P(k,x) = x*RNG(H(k),x))$. RNG is a random number generator function seeded with H(k).

pseudo random will still be determinsic, as we always seed it using the hash value of the key.

Chaos with cycles:\
a problem is that most probing sequences tend to produce a cycle shorter than the table size itself, this means that even if the load factor is below 1 (there are still open locations), we won't reach them and be stuck in a loop. we usually restrict the domain of the probing to those who produce cycles of size n.

for table of size N, H(k) is a hashing function and P(k,x) is a probing function. the following is the general insertion method.

```
x:=1
keyHash :=H(k)
index := keyHash

while table[index] != null:
  index = (keyHash + P(k,x)) Mod N
  x=x+1

insert (k,v) at table [index]
```

#### Linear Probing

we offset the index by using a probing function.

Linear probing is a method which probes according to a linear formula: $P(x) = ax +b$. where a and b are constants and a is not zero. the constant b is obsolete. the problem is that this probing won't necessarily produce a fully cycle of size n.\
assume p(x) = 3x, H(k) =4, N=9. we get an infinite loop of probing only three locations.

> 1. H(k) + P(0) Mod N = 4
> 2. H(k) + P(1) Mod N = 7
> 3. H(k) + P(2) Mod N = 1
> 4. H(k) + P(3) Mod N = 4
> 5. H(k) + P(4) Mod N = 7
> 6. H(k) + P(5) Mod N = 1
> 7. H(k) + P(6) Mod N = 4
> 8. H(k) + P(7) Mod N = 7
> 9. H(k) + P(8) Mod N = 1
>
> ...

we need to determine:

> Q: "which value(s) of consant _a_ in P(x) = ax produce a full cycle module N?"\
> A: "This happens when _a_ and N are **relatively prime**. two number are relatively prime if the **greatest common denominator (GCD)** is equal to one. Hence, when GCD(_a_,N) = 1 the probing function P(x) will be able to generate a complete cycle and we will ways be able fo find an empty bucket!."

relatively prime numbers examples:

| N   | _a_ | GCD | relatively prime? |
| --- | --- | --- | ----------------- |
| 9   | 6   | 3   | No                |
| 9   | 9   | 9   | No                |
| 10  | 6   | 2   | No                |
| 11  | 6   | 1   | Yes               |
| 11  | 22  | 11  | No                |
| 15  | 14  | 1   | Yes               |
| 15  | 12  | 1   | Yes               |

a common choice is to use P(x) =1\*x, because the gcd will always be 1. when we resize the table the probing function shouldn't change (so the GCD will still be 1).

#### Quadratic Probing

this time, our probing function P(x) is quadratic formula $P(x) = ax^2 +bx +c$. where a,b,c are constants, a is not zero, and c is obsolete. we still need to be worried about infinite loops.

assume p(x) = 2x^2 +2 , H(k) =4, N=9. we get an infinite loop of probing only two locations.

> 1. H(k) + P(0) Mod N = 4
> 2. H(k) + P(1) Mod N = 7
> 3. H(k) + P(3) Mod N = 4
> 4. H(k) + P(4) Mod N = 7
> 5. H(k) + P(6) Mod N = 4
> 6. H(k) + P(7) Mod N = 7
>
> ...

> so what kinds of probing function do always work?
>
> 1. let P(x) = x^2, keep the table size a prime number greater than 3 and keep the load factor (&alpha;) below 0.5
> 2. let P(x) = (x^2 + x)/2 and keep the table size a power of 2.
> 3. let P(x) = (-1^x)\*x^2 and keep the table size a prime N when N is congruent to 3 mod 4;

(e.g N = 23, a prime where 23 mod 4 = 3)

we will focus on the second form, $P(x) = (x^2 +x)/2$. where table size is power of 2.

assume N=8 (power of two), load factor = 0.4. so threshold before resize is 3.

| Operation       | Hash          | Probing result          | index | note                            |
| --------------- | ------------- | ----------------------- | ----- | ------------------------------- |
| insert(k1,v1)   | H(k1) = 6     | p(0)=0                  | 6     |                                 |
| insert(k2,v2)   | H(k2) = 5     | p(0)=0                  | 5     |                                 |
| insert(k3,v3)   | H(k3) = 5     | p(1)=1,P(2)=3           | 0     | collisions, resize from 8 to 16 |
| re-inset(k3,v3) | H(k3) = 5     | P(0)=0                  | 5     |                                 |
| re-inset(k2,v2) | H(k2) = 5     | p(1)=1                  | 9     | collision                       |
| re-inset(k1,v1) | H(k1) = 6     | p(1)=1                  | 7     | collision                       |
| insert(k4,v4)   | H(k4) = 35410 | P(0)=0                  | 2     |                                 |
| update(k3,v5)   | H(k3) = 5     | P(1)=0                  | 5     | update value, k3 already exists |
| insert(k6,v6)   | H(k6) = -6413 | P(0)                    | 3     |                                 |
| insert(k7,v7)   | H(k7) = 2     | P(1)=1, p(2)=3, P(3) =6 | 8     | collisions                      |

#### Double Hashing

This time we use a double hashing probing function, $P(k,x) = x*H_2(k)$ where H2(k) is a second hash function.
h2(k) must has the same types of keys as h1(k),

> Note: double hashing reduces to linear probing (except that the constant is unknown until linear run time)

we still can run into the same problem where we can run into infinite cycles.
to fix this we can pick a table size which is a prime factor and choose a delta value &delta;
$\delta =H_2(K)\ mod\ N$. if delta is zero, we are guranteed to be stuck in a cycle. so we limit the delta between 1 inclusive and N (exclusive), so if N is prime, the gcd will always be 1.
this ensures that we won't get in a cycle.

we need an additional hash function that works with the same type, we want a systematic way to produce a hash function.
there are many well known hash function (a pool called _universal hash functions_), which operate on fundamental data types.

example with double hashing, P(x)= x\*H2(K), table size n =7, factor &alpha; = 0.75, so we resize at 5 elements/

| Operation        | Hash                      | delta &delta;          | Result                                   | index | note                                        |
| ---------------- | ------------------------- | ---------------------- | ---------------------------------------- | ----- | ------------------------------------------- |
| insert(k1,v1)    | H1(k1) = 67 , H2(k1) = 34 | 34 mod 7 = 6           | 67 + 0\*6 mod 7                          | 4     |                                             |
| insert(k2,v2)    | H1(k2) = 2, H2(k2) = -79  | -79 mod 7 = 5          | 2 + 0\*5 mod 7                           | 2     |                                             |
| insert(k3,v3)    | H1(k3) = 2, H2(k3) =10    | 10 mod 7 =3            | 2 +0\*3 mod 7, 2+1\*3 mod 7              | 5     | hash collision                              |
| insert(k4,v4)    | H1(k4) = 2, H2(k4) = 7    | 7 mod 7 =0 -> **1**    | 2 +0 \*1 mod 7, 2+1 \*1 mod 7            | 3     | limit delta to 1, hash collision            |
| update(k3,v5)    | H1(k3) =2 , H2(k3) = 10   | 10 mod 7 =3            |                                          | 5     | update                                      |
| insert(k6,v6)    | H1(k6) =3 , H2(k6) =23    | 23 mod 7 =2            | 3+0\*2 mod 7, 3+1\*2 mod 7, 3+2\*2 mod 7 | 1     | collisions, need to resize, 7\*2 = 14 ~= 17 |
| re-insert(k6,v16 | H1(k6) = 3 , H2(k6) = 23  | 23 mode 17 = 6         | 3 + 0\*6 mod 17                          | 3     |                                             |
| re-insert(k2,v2) | H1(k2) = 2, H2(k2) = -79  | -79 mod 17 = 11        | 2 + 0\*11 mod 7                          | 2     |                                             |
| re-insert(k4,v4) | H1(k4) = 2, H2(k4) = 7    | 7 mod 17 =7            | 2 +0 \*7 mod 17, 2+1 \* 7 mod 17         | 9     | hash collision                              |
| re-insert(k1,v1) | H1(k1) = 67, H2(k1) =34   | 34 mod 17 = 0 -> **1** | 67 +0\*\1 mod 17                         | 16    |                                             |
| re-insert(k3,v5) | H1(k3) = 2, H2(k3) =10    | 10 mod 17 =10          | 2 +0\*10 mod 17, 2+1\*10 mod 17          | 12    | hash collision                              |
| insert(k7,v7)    | H1(k7) = 15, H2(k7) =3    | 3 mod 17 =3            | 15+0\*3 mod 17                           | 15    |                                             |

#### Removing Elements

when we remove naively, we will run into problem, we replace a remove element with a null, but that means we stop searching.
one method to handle this is to use a specific marker called 'a tombstone', which indicates that this index once contained an element, so we need to keep searching.
tombstone count as filled elements, so once we increase the size we get rid of them as part of the re-insertion process.
Additionally, we can replace a tombstone when we insert elements. when we search for values and go over tombstone, we can replace the first tombstone with the value we found.

### Implementation

java, generic, implements Iterable, has default capacity and load factors, threshold.
Tombstone element, a method to get a number which is a power of 2. quadratic probing function, normalizing indexes (make positive, mod by capacity).
resizing the table, trying to replace the tombstones we encounter, we will never run out of space because we keep increasing the size.

</details>

## Fenwick Tree/Binary Indexed Tree

<details>
<summary>
A tree like structure that stores computed values in an hierarchical order and allows range based operations on cached values.
</summary>
Binary indexed tree.

- Range Query
- Point updates
- Construction

motivation:

> "given an array of integer value, compute the range sum betwen index \[i,j) (include i, exclude j)"

we can do this for each range, start from the buttom and search upwards. but this gets redundant, we end up calculating the same values many times.
instead we can define P to ba array containing all the **prefix sums** of the array

> Array A: [5,-3,6,1,0,-4,11,6,2,7]\
> Array P: [0,5,2,8,9,9,5,16,22,24,31]

so when we want a sum of a range, we can compute the differences.
The problem is that when we want to update the original data, we need to calculate all the prefix sums again.
this means that we get O(1) for static unchanging arrays, but O(n) for updates.

> "A Fenwick Tree (also called a binary indexed tree) is a data structure that supports _sum ranges queries_
> as well as setting values in a static array and getting the value of the prefix sum up some index efficiently."

the construction is linear time, point update and range operations (sum, update) are log linear.

each cll is resposbile not only for itself, but also for other cells.

> "Unlike a regular array, in a fenwick tee, a specific cell is responsible for other cells as well.
> the position of the **least significant bit** (LSB) determins the range of responsability the cell has to the cells below himself."

fenwick trees are one based.

- for index 12, the binary is 1100, lsb is position 3, so this cells is responsible for 2^(3-1) = 4 cells.
- for index 11, the binary is 1001, lsb is position 1, so this cells is responsible for 2^(1-1) = 1 cells (itself).
- for index 10, the binary is 1010, lsb is position 2, so this cells is responsible for 2^(2-1) = 2 cells.

cells with odd number indexing (one-based) have lsb at position 1, so they are only responsible for themselves

to do a range query, we compute the prefix sum up to c certain index, which let's us perform the range sum queries. we start at some index and cascade down until we reach zero.

sum of up to 7 = A\[7] + A\[6]+ A\[4]. A\[7] is responsible only for itself,A\[6] is responsible for both itself and another cell, A\[4] is responsible for itself and three more cells, so we reach zero.
sum of up to 11 = A\[11] (self), + A\[10] (self + one), A\[8] (self + seven), which get us to zero.
sum of up to 4 = A\[4] (self + 3), we got to zero.

lets try an interval sum between 11 and 15:\
sum of 15 = A\[15] (1) + A\[14] (2) + A\[12] (4) ->A\[8] (8)\
sum of 11 (exclusive) = A\[10] (2) + A\[8] (8)

in the worst case,we are querying a cell that has a binary representation of all ones (number of the form 2^n-1), like 7,15,31... .
so, the worst case is for us to do two queries that cost log2(n) operations each.

> Range Query Algorithm:
> "To do a range query from [i,j] (both inclusive), a Fenwick tree of size N:

```
function prefixSum(i):
  sum:= 0
  while i!=0:
    sum = sum +tree[i]
    i = i - LSB(i)
  return sum

function rangeQuery(i,j):
  return prefixSum(j) - prefixSum(i-1)
```

### Tree Point Update

the reverse of the query sum

"we started at a value, and then continuously removed the lsb until we reach zero"

> - 13 = 0b1101 ,0b1101-0b0001 = 0b1100 = 12
> - 12 = 0b1100 ,0b1100-0b0100 = 0b1000 = 8
> - 8 = 0b1000 ,0b1000-0b1000 = 0b0000 = 0
> - 0 = 0b0000

a point update is adding the lsb, and cascading up until we land out of the array bounds. at every index we reach we need to update it with the new value

> - 9 = 0b1001, 0b1001 + 0b0001 = 0b1010 = 10
> - 10 = 0b1010, 0b1010 + 0b0010 = 0b1100 = 12
> - 12 = 0b1100, 0b1100 + 0b0100 = 0b10000 = 16
> - 16 = 0b10000, 0b10000 +0b10000 =0b100000 = 32

if we add x to postion 6, we need to update three cells in the fenwick tree. we add x to postions, 6,8,16
6 = 0b0110 -> 8 =0b1000 -> 16 = 0b10000

> To update the cell at index i in a fenwick Tree of size N:

```
function add(i,x):
  while i<N:
    tree[i] = tree[i]+x
    i = i + LSB(i)
```

### Fenwick Tree Construction

construction is O(n).

> Naive construction:\
> "Let A be an array of values, for each element in A at index i do a point update on the fenwick tree with a value of A\[i], there are n elements and each point update takes O(log(n)) for a total of *O(n*log(n))\*, can we do better?"

```
A=[a1,a2,,...,aN] //array
for (i=1 ;i <= N,++i )
{
  tree.add(i,a[i]);
}
```

> Linear construction:\
> Input values we wish to turn into a legitmate fenwick tree
>
> "Add the value in the current cell to the immediate cell that is responsible for us, this resembles what we did for th point update, but only once cell at a time.
> this is make the 'cascading' in range queries possible by propagatin the value in each cell throughout the tree."
>
> let i be the current index, the immediate cell above us is at position j given by: j:= i+LSB(i)

```
A=[a1,a2,,...,aN] //array
for (i=1 ;i <= N,++i )
{
  tree[i] +=A[i];
  tree[i+LSB(i)] += A[i]
}
```

or

```
A=[a1,a2,,...,aN] //array
for (i=1 ;i <= N,++i )
{
  tree[i] = A[i]; //initial construction
}

for (i=1 ;i <= N,++i )
{
  j = i + LSB(i)
  if j <= N:
    tree[j] += tree[i] ; update parents
}
```

we update two cells at each time, so the construction is linear time. we ignore parent which are out of bounds

```
#make sure vales is 1-based
function construct(values):
  N:= length(values)

  #clone the values
  tree = deepCopy(values)
  for i= 1,2,3,...,N:
    j:=i + LSB(i)
    if J<= N:
      tree[j] = tree[j]+tree[i]
  return tree
```

### Implementation

Java source code. take care about the indexing, we never use index zero. we need things to be one-based.
cascading up, cascading down, method set

```java
private int lsb(int i)
{
  return i & i-1;
  //return Integer.LowestOneBit(i); // java built in
}

public long prefixSum(int i)
{
  long sum = 0L;
  while (i !0)
  {
    sum += tree[i];
    i &= ~lsb(i); // same as i -= lsb(i); equivalent forms
  }
  return sum;
}
public long sum(int, int j)
{
  //make sure i <j and the are in the range
  return prefixSum(j)- prefixSum(i-1);
}
//add k value to index i (one base)
public add(int i,long k)
{
  while (i< tree.length)
  {
    tree[i]+=k;
    i+=lsb(i);
  }
}
// set index i to be equal to k, one based
public void set(int i, long k)
{
  long value = sum(i,i); //check what inside/ like tree[i]
  add(i,k-value); // add the difference, and let it propegate upwards
}
```

</details>

## AVL Balanced Binary Search Tree

<details>
<summary>
Balanced Binary Search Trees. 
</summary>

> "A **Balanced Binary Search Tree (BBST)** is a _self-balancing_ binary search tree.
> This type of tree will adjust itself in order to maintain a low (logarithmic) height allowing for faster operation such as insertions and deletions."

binary tree don't perform well with sorted data, but balanced search tree keep themselves balanced and maintain the form which allows them the same complexicy of O(long(n)) without degradation

> "The secret ingredient to most BBST algorithms is the clever usage of a _tree invariant_ and _tree rotations_.\
> A tree invariant is a property/rule you impose on your tree that it must meet after each operation.
> To ensure that the invariant is always satisfied a series of tree rotations are normally applied."

<img src='https://g.gravizo.com/svg?
 digraph Tree {
   A1 -> B1;
   A1 -> C1;
   B1 -> D1;
   B1 -> E1;
   ;
   B2 -> D2;
   B2 -> A2;
   A2 -> E2;
   A2 -> C2;
 }
'/>

we know that in the left tree, D < B < E < A < C, and this holds true in the right tree as well. for every node is the left sub tree, all the value are smaller than the parent and the element in the right subTree

we don't care about which is the root, as long as the binary tree invariant remains. we are free to shuffle and transform the tree.

we start with tree where P->[A,] A->[B,C], B->[D,E], where p may or may not exist.

```
function rightRotate(A):
  B:= A.left
  A.left = B.right
  B.right = A
  return B
```

we end up with a tree B->[D,A],A->[E,C].
if node A has a parent node P, then we end up with it still pointing to node A, so we would have two nodes pointing to it, which is bad. we need to update this link as well.
this is usually down with the return value, so it's sometimes easier to store the parent references in each node, which makes the rotations more complicated (six pointers instead of three).

```
function rightRotate(A):
  P:=A.parent
  B:=A.left
  A.left = B.right
  if B.right != null:
    B.right.parent = A
  B.right = A
  A.parent = B
  B.parent= P

  #update parent down link

  if P != null:
    if P.left == A:
      P.left =B
    else:
      P.right = B
  return B
```

### AVL Inserting Elements

> "An **AVL tree** is one of many types of balacned binary search trees which allow for logarithmic O(log(n)) insertions, deleteion and search operations.
>
> In fact, it was the first type of BBST to be discovered, soon after, many other types of BBST's statered to emerge
> including the 2-3 tree the AA tree, the scapegoat tree, and it's main rival, the red-black tree."

the AVL tree is governed by the **AVL Tree Invariant**:\

> The property which keeps an ABL tree balanced is called the **Balanced Factor (BF)**.\
> BF(node) = H(node.right) - H(node.left)\
> Where H(x) is the height of node x, recall that H(x) is calculated as the _number of edges_ between x and the furthest leaf.\
> The invariant in the AVL which forces it remain balanced is the requirement that the balance factor to always be either -1,0 or +1

the difference between the heights must be [-1,0,+1], in other words, one subtree can't have more than one layyer deeper than it's sibling.

> Node information to store:\
>
> - The actual value we are storing in the node (must be comparable, so we can insert it).
> - A value storing this node's _balance factor_.
> - The _height_ of this node in the tree.
> - Pointers to the left/right child nodes (subtrees).

when the algorithm runs, we need to update those values.
if the balanced factor isn't {-1,0,+1}, then it must be either {-2,+2}, which we can adjust using tree rotations.

there are four cases: two simple cases and tow complex case that resolve into the simple cases. they are all symmetrical.\
In this example, A =5,B= 4,C=3

| case        | initial state    | action                                                                       | action states             |
| ----------- | ---------------- | ---------------------------------------------------------------------------- | ------------------------- |
| left left   | A->[B,], B->[C,] | right notation                                                               | B->[C,A]                  |
| left right  | A->[B,], B->[,C] | left rotation on the right child and, then a right rotation (left left case) | A->[B,],B->[C,]. B->[C,A] |
| right right | C->[,B], B->[,A] | left rotation                                                                | B->[C,A]                  |
| right left  | C->[,A], A->[B,] | right rotation on the left child, and then left rotation(right right case)   | C->[,B],B->[,A]. B->[C,A] |

public facing method, returning true when successfully inserted, and false when it already exists.

```
function insert(value):
  if value == null:
    return false
  # Only insert Unique values
  if !contains(root,value):
    root = insert(root,value)
    nodeCount = nodeCount +1
    return true
  # Value already exists in tree
    return false

# Private function
function insert(node,value):
  cmp:= compate(value,node.value)

  if (cmp<0):
    node.left = insert(node.left, value)
  else:
    node.right = insert(node.right,value)

  # Update balance factor and height values
  update(node)

  # Rebalance tree.
  return balance(node)

function update(node):
  lh:= -1;
  rh:= -1
  if node.left != null: lh = node.left.height
  if node.right != null: rh = node.right.height

  # Update this Node height
  node.height = 1+ max(lh,rh)
  # Update balance factor
  node.bf = rh-lh

function balance(node):
  #left Heavy subtree
  if node.bf == -2:
    if node.left.bf <= 0:
      return leftLeftCase(node)
    else:
      return leftRightCase(node)

  #right heavy subtree
  else if node.bf== 2:
    if node.right.bf >= 0:
      return rightRightCase(node)
    else:
      return rightLeftCase(node)

  # balance factor is -1,0,+1
  # no need to balance
  return node;


# cases

function leftLeftCase(node):
  return rightRotation(node)

function leftRightCase(node):
  node.left = leftRotation(node.left)
  return leftLeftCase(node)

function rightRightCase(node):
  return leftRotation(node)

function leftLeftCase(node):
  node.right = rightRotation(node.right)
  return rightRightCase(node)
```

we also need out update function inside the rotations to augment the height when we do rotations.

### AVL Removing Elements

removing elements is similar to removing elements from a regular binary tree.

reminder:

> removal phases:
>
> 1. _Find_ the element we wish to remove.
> 2. _Replace_ the node with it's successor (if it has any) to maintain the BST invariant.
>
> the four cases of removing:
>
> 1. removing leaf.
> 2. removing a node with only left subtree.
> 3. removing a node with only right subtree.
> 4. removing a node with both left and right subtrees.

for an AVL tree, we need to ensure the AVL invaraint, just like when we insert nodes,we update the height and balance factor, and then re-balance them.

### Implementation

Java source code. each node contains the balance factor and the height,doesn't allow duplicates or nulls.
updating and rebalancing the tree. we get the balance factor from the heights, and then we balance according to it and then update again in the rotations (order matters, child before parent).
in the remove methods we choose to replace with the successor node from the subtree which is higher (hoping this will make balancing easier).
when we remove nodes we update and rebalance. this is done in a recursive way.

</details>

## Indexed Priority Queue

<details>
<summary>
A heap that provides faster operations by maintaining a lookup between the key and the index of the element in an array. mapping exist from the key to the position and the value, and from the position to the key index.
</summary>

building on top of priority queues.

> "An indexed Priority Queue is a traditional priority queue varian which on top of the regular PQ operations also supports _quick updates and deletions_ of _key-value pairs_."

looking up and dynamically changing values inside the queue.
we assign index values to all the keys, we form bi-directional mapping.

A typicall priority queue is implemented as heaps under the hood (which are internally arrays), so we map between the keys (which can be of any type) and the index (a number between 0 and N) with an hash table for the mapping.

> Supported opperations of the IPQ ADT Interface:\
> assume k is the key, so the index is **ki=map\[k]**
>
> - delete(ki)
> - valueOf(ki)
> - contains(ki)
> - peekMinKeyIndex()
> - pollMinKeyIndex()
> - peekMinValue()
> - pollMinValue
> - insert(ki, value)
> - update(ki, value)
> - decreaseKey(ki, value)
> - increaseKey(ki, value)

there are several ways to implement this ADT, but we will use a binary heap, which gets some logarithmic operations as opposed to linear time.

recall that in a binary heap, for index i, children are 2i+1, 2i+2 (zero based). insertion is done at the end of the heap and bubbling up and down to maintain the heap invariant. removal is done by swapping and then bubbling.
key-> key Index -> value (index in heap?)

we maintain an N array of values a **position map** for the lookup. "the index of the node in the heap for a given key index (ki)". and an **inverse map** for the reverse lookup. "getting a key given the node position".

1. map from key to key index.
2. values are orgainzed in an array (vals) which the keyIndex maps to.
3. the positional array is mapping from the key index to the heap index.
4. an inverse lookup table, Inverse Map, from the node position to the key index.
5. ~~the heap itself~~

> to find a node:
>
> - get the key index from the mapping of the keys to indexing.
> - use that key index in the positional mapping to get the index of the node in the heap.
> - (?verify that the value in the values array matches the value in the heap?)
>
>   to find a key from the node:
>
> - use the node index as a key on the inverse map

bidirectional hash table: key to key index

| name (key) | key index (ki) |
| ---------- | -------------- |
| Anna       | 0              |
| Bella      | 1              |
| Carly      | 2              |
| ...        | ...            |
| Laura      | 11             |

| index          | 0   | 1   | 2   | ... | 10  | 11  | 12  | 13  | 14  | notes                           |
| -------------- | --- | --- | --- | --- | --- | --- | --- | --- | --- | ------------------------------- |
| Values         | 3   | 15  | 11  | ... | 16  | 14  | -1  | -1  | -1  | ki -> value                     |
| Positional Map | 2   | 9   | 11  | ... | 10  | 5   | -1  | -1  | -1  | ki -> heap index                |
| Inverse Map    | 7   | 6   | 0   | ... | 10  | 2   | -1  | -1  | -1  | heap index -> key mapping index |

### Inserting Elements

> insertion:
>
> 1. add key to the bidirectional hash table
> 2. add node at the end of the heap, and update the values array, the Positional map and inverse map accordingly.
> 3. if heap invariant isn't satisfied perform bubbling, at each bubbling stage, use the inverse map to find the key index of the swapped node, and swap both the values in the positional mapping and the inverse mapping.
> 4. keep doing so until the heap invariant is satisfied.

```
# Insert a value into the min indexed binary heap.
# The key index must no already be in the heap.
# The value may not be null.

function insert(ki, value):
  values[ki] = value
  #'sz' is the current size of the heap
  pm[ki] = sz
  im[sz] = ki
  swim(sz)
  za = sz+1

#swims up node i (zero based) until the heap invariant is satisfied
function swim(i):
   for(p = i-1/2; i> 0 and less(i,p); ):
    swap(i,p)
    i=p
    p=i-1/2;

#swap inverse map and positional mapping, not the values
function swap(i,j):
  #update positional mapping
  pm[im[j]] =i
  pm[im[i]] =j

  #swap inverse mapping
  temp:= im[i]
  im[i] = im[j]
  im[j] = temp

function less(i,j):
  #compare values with the inverse mapping
  return values[im[i]] < values[im[j]]

```

### Polling and Removing Elemens

_Polling_ is still O(log(n)) in an IPQ, but _removing_ is improved from O(n) in the tradition PQ to O(log(n)) since the node position lookup are O(1) but repositioning is O(log(n)).

> Polling:
>
> 1.  Exchange the root node with the buttom node.
> 2.  Swap the index values in the positional map and the inverse map.
> 3.  Remove the node from the heap and return the key-value pair.
> 4.  Mark the position of the remove value as null in the values array.
> 5.  Restore the heap invariant by moving the swapped node up or down as needed (update arrays as needed).
>
> To Remove a node by a key:
>
> 1. use the bidirectional map to find the key index.
> 2. use the positional map and the key index to find the node position.
> 3. swap that node with the last node, (update positional mapping and inverse mapping)
> 4. remove the node from the heap and return the key-value pair.
> 5. continue as before (sink up or down as needed)

```
# Deletes the node with the key index ki in the heap.
# The key index ki must exists and be present in the heap.

function remove(ki):
  i := pm[ki]
  # sz is the size of the heap
  swap(i, sz)
  sz = sz-1
  sink(i)
  swim(i)
  value[ki] = null
  pm[ki] = -1
  im[sz] = -1

# Sinks the node at index i by swapping itself withe the smallest of the right or left child nodes
function swim(i):
    while(true):
      left = 2*i+1
      right = 2*i+2
      # Default left
      smallest:= left
      if right < sz and less(right,left):
        smallest = right

      # Exit condition, no children or no need to swap anymore
      if left >= sz or less(i, smallest):
        break

      swap(smallest,i)
      i = smallest
```

### Updates

like removals, updates also take O(log(n)) time, because lookup is O(1) with the positional mapping, but adjusting the heap after the change is O(log(n)).

> Update value by key
>
> 1. find key index with the key.
> 2. update the values array
> 3. maintain the heap invariant by sinking or swimming

```
# Update the value of a key in a binary heap.
# The key index ki must exists and be present in the heap.

function update(ki, value):
  i=pm[ki]
  values=[ki] = value
  sink(i)
  swim(i)
```

### Decrease and Increase Key

> "In many applications (e.g Dijkstra's and Prims algorithm) it is often useful to only update a given key to make its value either always smaller (or larger). in the even that worse value is given the value in the IPQ should not be updated.\
> In such situations it is useful to define a more restrictive form of update operation we call _increaseKey(ki,v)_ and _decreaseKey(ki,v)_."

the update only happens if the new value is towards one particular direction as compared to the existing value

```
# assume ki and newValue are valid inputs, and we are dealing with a min indexed binary heap

function decreaseKey(ki,newValue):
  if less(newValue, values[ki]):
    values[ki] =newValue
    swim(pm[ki])

function increase(ki,newValue):
  if less(values[ki],newValue):
    values[ki] =newValue
    sink(pm[ki])
```

### Implementation

java source code. size, capacity, values, positional mapping, inverse mapping.\
also has child-parent arrays for quick lookups, but that's not required.\ positional mapping and inverted map are initiliazed to -1. operations use the KeyIndex, rather than the key itself.

</details>

## Sparse Tables

<details>
<summary>
a table used to pre-compute results of associative function on a static range of inputs, which can later be used for range queries. gets even better performance with overlap friendly functions, costly in memory.
</summary>

> "Sparese tables are all about doing efficient **range queries** on _static arrays_. Range queries come in a variety of flavors, but the most common types are _min_, _max_, _sum_ and _gcd_ range queries.

| index  | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  |
| ------ | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Values | 4   | 2   | 3   | 7   | 1   | 5   | 3   | 3   | 9   | 6   | 7   | -1  | 4   |

for example, we can ask what is the maximum value between \[4,12], the gcd between \[6,9] or the sum of values between \[2,11].

the data is immutable, it won't change.

> Intuition for sparse tables
>
> - Every positive integer can be easily represented as a sum of powers of 2 given it's binary representation.
>   - 19 = 0b10011 = 2^4 (16) + 2^1 (2) + 2^0(1) = 19
> - Similarly, any interval \[l,r] can be broken into smaller intervals of powers of two.
>   - \[5,17] = \[5, 5+2^3) U \[13, 13+2^2) U \[17, 17+2^0) = \[5, 13) U \[13,17) U \[17,18)

the length of the intervals is powers of two.

### Range Combination Function

> "A sparse table can help us answer all these questions efficiently, for **associative functions**, a sparse table can answer range queries in _O(log(n))_."\
> a function f(x,y) is associative if:\
> f(a,f(b,c)) = f(f(a,b),c) for all a,b,c\
> Operations such as addition and multiplication are associative, but substraction and exponantation are not.

- $1+(2+3) = 6 = (1+2)+3$
- $2*(3*4) = 24 = (2*3)*4$
- $1-(2-3) = 2 \ne (1-2)-3 = -4$

but we can do better.

> "When the range query combination function is **"overlap friendly**, then range queries on a sparse table can answered in O(1)."\
> being overlap friendly mean a function yields the same answer regardless of whether it is combining ranges which overlap or those that do not:\
> f(f(a,b),f(b,c)) = f(a,f(b,c)) for all valid a,b,c\

assume array \[4,2,3,-6,6,7,1,-2,8,9,8,1,1,3,-6,2], where f(x,y) = x+y.\
r1\[0,6] = 16, r2\[7,9] = 7, r3\[10,16] = 18

addition is not overlap friendly:

$
f(f(r1,r2),f(r2,r3)) = f(f(16,7),f(7,18)) = (16+7) + (7+18) =48 \\
\ne \\
f(r1,f(r2,r3)) = f(16,f(7,18)) = 16 +7 +18 = 41
$

| function f(a,b) | associative | overlap friendly? | notes          |
| --------------- | ----------- | ----------------- | -------------- |
| 1\*b            |             | yes               |                |
| a\*\*b          | yes         | no                | multiplication |
| min(a,b)        | yes         | yes               |                |
| max(a,b)        | yes         | yes               |                |
| a\*\*b          | no          | no                | substraction   |
| (a\*b)/a, a!=0  |             | yes               | same as 1\*b   |
| gcd(a,b)        | yes         | yes               |                |

the idea behind a sparse table is to precompute the answer for all interval of size 2^x to efficiently answer range queries.

N = size of the input, 2^P is the largest power of 2 that fits in the length of the values array.\
for N = 13:
P = floor(log2(N)) = floor(log2(13)) =3

| index  | 0   | 1   | 2   | 3   | 4   | 5   | 6   | 7   | 8   | 9   | 10  | 11  | 12  |
| ------ | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Values | 4   | 2   | 3   | 7   | 1   | 5   | 3   | 3   | 9   | 6   | 7   | -1  | 4   |

the memory is stored in memory, we need N^N memory to store them.\
we initialize a table with P+1 rows,and N columns, each cell (i,j) represents the answer for the range \[j,j+2^i) in the original array.\
thr first row is the same as the input values. the rest of the rows are filled according to the function.

|         | 0   | 1   | 2                                              | 3   | 4   | 5                      | 6       | 7          | 8       | 9       | 10      | 11      | 12      |
| ------- | --- | --- | ---------------------------------------------- | --- | --- | ---------------------- | ------- | ---------- | ------- | ------- | ------- | ------- | ------- |
| $0,2^0$ | 4   | 2   | 3                                              | 7   | 1   | 5                      | 3       | 3          | 9       | 6       | 7       | -1      | 4       |
| $1,2^1$ |     |     |                                                |     |     |                        |         | $f(a7,a8)$ |         |         |         |         | &empty; |
| $2,2^2$ |     |     |                                                |     |     | $f(f(a5,a6),f(a7,a8))$ |         |            |         |         | &empty; | &empty; | &empty; |
| $3,2^3$ |     |     | $f(f(f(a2,a3),f(a4,a5)),f(f(a6,a7),f(a8,a9)))$ |     |     |                        | &empty; | &empty;    | &empty; | &empty; | &empty; | &empty; | &empty; |

for example:

- cell(1,7) represents the answer for range \[7,9), so for a min sparse table it will be the value of min(3,9) -> 3. for a sum sparse table query, the value will be sum(3,9) -> 12.
- cell(2,5) represents the answer for range \[5,9), so for a min sparse table it will be the value of min(5,3,9,9) -> 3. for a sum sparse table query, the value will be sum(5,3,3,9) -> 20.
- cell(3,2) represents the answer for the range \[2,10), so for min value it's min(3,7,1,5,3,3,3,9,6)-> 1, and fo sum it's sum(3,7,1,5,3,3,3,9,6) -> 40.
- cell(2,10) is the range \[10,14), which is out of bounds. we can simply ignore it, we don't care about partial ranges. let's mark this cell as &empty;, and all the other partial cells as well.

now for an example with a range combinnation value. we need a range combination function, so let f(x,y) = min(x,y)

we can combine use the results of the above cells because of the associativity, each range can be split into two intervals of even length. cell (i,j) can be split into ranges \[j,j+2^i) and \[(j+2^i) -1, j+2^(i-1)), which corespond to cells(i-1,j) and (i-1, j+2^(i-1)).

| min range | 0   | 1   | 2          | 3   | 4          | 5   | 6          | 7       | 8          | 9       | 10      | 11      | 12      |
| --------- | --- | --- | ---------- | --- | ---------- | --- | ---------- | ------- | ---------- | ------- | ------- | ------- | ------- |
| $0,2^0$   | 4   | 2   | 3          | 7   | 1          | 5   | 3          | 3       | 9          | 6       | 7       | -1      | 4       |
| $1,2^1$   |     |     | $f(3,7)=3$ |     | $f(1,5)=1$ |     | $f(3,3)=3$ |         | $f(9,6)=6$ |         |         |         | &empty; |
| $2,2^2$   |     |     | $f(3,1)=1$ |     |            |     | $f(3,6)=3$ |         |            |         | &empty; | &empty; | &empty; |
| $3,2^3$   |     |     | $f(1,3)=1$ |     |            |     | &empty;    | &empty; | &empty;    | &empty; | &empty; | &empty; | &empty; |

we can continue to fill the table by combining values from the previous rows.

suppose we want to know the minimal value between \[1,11]? the table already includes the answer for intervals of lengths which are power of 2.

> "Let k be the largest power of two that fits in the length of the range between \[l,r].\
> Knowing k we can easily do a lookup in the table to find the minimum in between the ranges \[l, l+k] (left interval) and \[r-k+1,r] (right interval). the right and the left intervals may overlap, but this doesn't matter (given the overlap friendly property), so long as the the entire range is covered."

in our example,p =3, k=2^3=8, we can compute the result of ranges \[1,1+8=9) and \[11-8+1=4,11) and get a correct answer. we don't care about the overlap of the ranges.\
another example, minimal of \[2,7], $p=floor(\log_2(7-2+1))$, so this time p is 2, k=4.\
the ranges are \[2,2+4=6) and \[7-4+1=4,7), so we look at cells(2,2) and 2(,4).\
final example \[3,5], p = $floor(\log_2(5-3+1))$ = 1, k=2, ranges are \[3,3+2=5) and \[5-2+1=4,5) so cells (1,3) and (1,4).

### Associative Function Queries (non overlap friendly)

> "Some functions such as multiplication and summation are associative, but are _not overlap friendly_. A sparse table can still handle these types of queries, but in _O(log(n))_ rather than O(1). The main issue with non-overlap friendly function with our previous approach is that overlapping intervals would yield the wrong answer.\
> The alternative approach to performing a range query is to do a **cascading query** on the sparse table by breaking the range \[l,r] into smaller ranger of size 2^x which **do not overlap**."

for example, the range \[2,15] can be split into three intervals of lengths 8,4 and 2.\
\[2,15] = \[2,2+2^3) U \[10,10+2^2) U \[14,14+2^1) \
= \[2,10) U \[10,14) U \[14,16)

in this example we use our sparse table to calculate products, f(x,y) = x\*y.

| product | 0   | 1   | 2   | 3   | 4       | 5       | 6       |
| ------- | --- | --- | --- | --- | ------- | ------- | ------- |
| $0,2^0$ | 1   | 2   | -3  | 2   | 4       | -1      | 5       |
| $1,2^1$ | 2   | -6  | -6  | 8   | -4      | -5      | &empty; |
| $2,2^2$ | -12 | -48 | 24  | -40 | &empty; | &empty; | &empty; |

for the product of range \[0,6], we break it apart p = $floor(\log_2(6-0+1))=2$, k =4.\
range \[0,6] -> \[0,0+2^2) U \[4,4+2^1) U \[6,6+2^0)

$
cell(2,0)*cell(1,4)*cell(0,6) \\
= -12*-4*5 \\
= 240
$

example for range \[1,5] -> \[1,1+2^2) U \[5,5+2^0)
$
cell(2,1)*cell(0,5) \\
= -48 * -1 \\
= 48
$

sometimes we want to get the index of the value, not just the value itself, this only makes sense for some operations (min, max), and not other(sum, gcd)

```
# The number of elements in the input array;
N =...

# Short for power,the larger 2^p that fits in N
P =... # calculated as flor(log_2(N))

# A quick lookup table for floor(log2(i)), 1<= i <= N
log2[] =... # size N+1, index 0 unused

# the sparse table contining integer values
dp[][] = .... #P+1 rows, N columns

# index table, associated with the values in the sparse table.
# only useful to query the index of the min(or max) element in the range [l,r] and not the value
# doesn't make sense for most other range query types like gcd or sum
it[][] = .... #P+1 rows, N columns

function BuildMinSparseTable(values):
  N = length(values)
  P = floor(log(n)/log(2))

  # Quick lookup table
  log2=[0,0,...,0,0] #size N+1
  for (i:=2;i <= N;i++):
    log2[i] = log2[i/2]+1

  #fill the first row

  for (i:=0;i < N;i++):
    dp[0][i] = values[i]
    it[0][i] = i

  # fill other rows
  for (p:=1;p <=P;p++):
    for (i:=0; i + (1<<p) <= N; i++):
      left := dp[p-1][i]
      right := dp[p-1][i+(1<<(p-1))]
      dp[p][i]= min(left,right)

      # Save/Propgate the index of smallest elements
      if left <= right:
        it[p][i] = it[p-1][i]
      else:
        it[p][i] = it[p-1][i+(1<<(p-1))]

# Query the smallest element in the range[l,r], O(1)
function MinQuery(l,r):
  len:= r-l-1
  p:= log2[len]
  left := dp[p-1][l]
  right := dp[p-1][r+(1<<(p-1))]
  return min(left, right)

# Query the smallest element in the range[l,r] by doing a cascading min query, O(log(n))
function CascadingMinQuery(l,r):
  min_val := +∞ # initialize with max value
  for (p:=log2[r-l+1]; l <=r; p =log2[r-l+1]):
    min_val = min(min_val,dp[p][l]
    l+= (1<<p)

  return min_val

# Return the index of the minimum element in the range [l,r] in the input array
# Will return the left most index if there are multiple elements of the same minimal value
function MinIndexQuery(l,r):
  len:= r-l-1
  p:= log2[len]
  left := dp[p-1][l]
  right := dp[p-1][r-(1<<p)+1]
  if left <= right:
    return it[p][l]
  else:
    return it[p][r-(1<<p) +1]

```

### Implementation

java source code, min range query table. storing the number of values = n, the maximum power of 2 (floor(log_2(n))) = p, a quick lookup table for the log_2 values for each length, storing the results in the two dimensional array dp (dynamic programming, data points), also storing the index in the two dimensional array it (index table).\
the input array can only contain immutable values. we use the left shift operator to get the interval sizes for each 'row' (value of p). method for O(1) look up, and method for finding the index of the minimal element in the range.

</details>

## Complexity Table

<details>
<summary>
Complexity for data structure operations.
</summary>

| Data structure                  | Access        | Search | Insertion        | Appending            | Deletion                                   |
| ------------------------------- | ------------- | ------ | ---------------- | -------------------- | ------------------------------------------ |
| Static Array                    | O(1)          | O(n)   | N/A              | N/A                  | N/A                                        |
| Dynamic Array                   | O(1)          | O(n)   | O(n)             | O(n)                 | O(n)                                       |
| Singly Linked List              | N/A           | O(n)   | at head O(1)     | at tail O(1)         | at head O(1), at tail O(n), in middle O(n) |
| Doubly Linked List              | N/A           | O(n)   | at head O(1)     | at tail O(1)         | at head O(1), at tail O(1), in middle O(n) |
| Stack                           | peek top O(1) | O(N)   | N/A, only push   | push at top O(1)     | pop top O(1)                               |
| Queue                           | front O(1)    | O(N)   | N/A,only enqueue | enqueue at back O(1) | dequeue front O(1) in middle O(N)          |
| Priority Queue with Binary Heap | Peeking O(1)  | N/A    |                  | Adding O(log(n))     | Polling O(log(n))                          |

### Union Find / Disjoint Sets

&alpha; stands for _Amortized constat time_.

| Operation          | Complexicy |
| ------------------ | ---------- |
| Construction       | O(n)       |
| Union              | &alpha;(n) |
| Find               | &alpha;(n) |
| Get component size | &alpha;(n) |
| Check if connected | &alpha;(n) |
| Count components   | O(1)       |

### Binary Search Tree

| Operation | Average   | Worst |
| --------- | --------- | ----- |
| Insert    | O(log(n)) | O(n)  |
| Delete    | O(log(n)) | O(n)  |
| Remove    | O(log(n)) | O(n)  |
| Search    | O(log(n)) | O(n)  |

### Hash Table

the constant time behavior attributed to hash tables is only true if there is a good **uniform** hash function.

| Operation | Average | Worst |
| --------- | ------- | ----- |
| Insert    | O(1)    | O(n)  |
| Remove    | O(1)    | O(n)  |
| Search    | O(1)    | O(n)  |

### Fenwick Tree

| Operation      | Complexicy |
| -------------- | ---------- |
| Construction   | O(n)       |
| Point update   | O(log(n))  |
| Range sum      | O(log(n))  |
| Range update   | O(log(n))  |
| Adding index   | N/A        |
| Removing index | N/A        |

### Balanced Binary Search Tree (AVL)

| Operation | Average   | Worst     |
| --------- | --------- | --------- |
| Insert    | O(log(n)) | O(log(n)) |
| Delete    | O(log(n)) | O(log(n)) |
| Remove    | O(log(n)) | O(log(n)) |
| Search    | O(log(n)) | O(log(n)) |

### Index Priority Queue as a binary Heap

| Operation              | Complexity |
| ---------------------- | ---------- |
| delete(ki)             | O(log(n))  |
| valueOf(ki)            | O(1)       |
| contains(ki)           | O(1)       |
| peekMinKeyIndex()      | O(1)       |
| pollMinKeyIndex()      | O(log(n))  |
| peekMinValue()         | O(1)       |
| pollMinValue           | O(log(n))  |
| insert(ki, value)      | O(log(n))  |
| update(ki, value)      | O(log(n))  |
| decreaseKey(ki, value) | O(log(n))  |
| increaseKey(ki, value) | O(log(n))  |

### Sparse Tables

| Function type                        | Complexity |
| ------------------------------------ | ---------- |
| Associative but not overlap friendly | O(log(n))  |
| Associative and overlap friendly     | O(1)       |

</details>
