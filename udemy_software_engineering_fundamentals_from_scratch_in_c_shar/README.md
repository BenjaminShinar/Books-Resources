<!--
// cSpell:ignore Kruskal
 -->

# Software Engineering Fundamentals From Scratch in C Sharp

<details>
<summary>
Surface level course about data structures, sorting and some basic problems.
</summary>

udemy course [Software Engineering Fundamentals From Scratch in C Sharp](https://www.udemy.com/course/software-engineering-fundamentals-from-scratch-in-c-sharp)

## Sorting Algorithms

<details>
<summary>
Sorting an array of integers.
</summary>

### Selection Sort

> 1. Divide an array into sorted and unsorted parts
> 2. Find smallest number in unsorted part
> 3. Swap it with first number in unsorted part
> 4. Add the first number from unsorted part to the end of sorted part.

we have a "wall", which divides between the sorted and unsorted parts. the sort is completed when the wall reaches the end of the array. if the first number after the wall is the smallest, there is no need to swap. The wall isn't an actual element, it's a marker to the position which we use.

```
[W,16,5,11,8,1,2,20]
[1,W,5,11,8,16,2,20]
[1,2,W,11,8,16,5,20]
[1,2,5,W,8,16,11,20]
[1,2,5,8,W,16,11,20]
[1,2,5,8,11,W,16,20]
[1,2,5,8,11,16,W,20]
[1,2,5,8,11,16,20,W]
```

code:

```csharp
public static void SelectionSort(int[] input)
{
  for (int i=0; i< input.Length -1; i++)
  {
    int indexOfSmallest =i;
    for (int j=i; j < input.Length -1; j++)
    {
      if (input[j] < input[indexOfSmallest])
      {
        indexOfSmallest=j;
      }
    }
    //swap
    int temp = input[i];
    input[i] = input[indexOfSmallest];
    input[indexOfSmallest] = temp;
  }
}
```

### Bubble Sort

Let the biggest number "bubble" to the end of an array.

1. Compare two numbers next to each other
2. If the first is one is smaller, we move to the next two numbers.
3. If the first is one is larger, we swap the first number with the second then move on.

```
[18,8,1,22,12,9,4]
[8,1,18,12,9,4,22]
[1,8,12,9,4,18,22]
[1,8,9,4,12,18,22]
[1,8,4,9,12,18,22]
[1,4,8,9,12,18,22]
```

code:

```csharp
public static void BubbleSort(int[] input)
{
  for (int i=0; i< input.Length -1; i++)
  {
    for (int j=0; j+1 < input.Length -i ; j++)
    {
      if (input[j] > input[j+1])
      {
        //swap
        int temp = input[j];
        input[j] = input[j+1];
        input[j+1] = temp;
      }
    }
  }
}
```

we can optimize to exit the code early if there was no swap, and if there wasn't any swap, then the array is sorted.

### Recursion

A function which calls itself, until some condition is met, and then it ends.

```csharp
public static int Factorial(int n)
{
  if (n <=1)
  {
    return 1;
  }
  else
  {
    return n * factorial(n - 1);
  }
}
```

### Merge Sort

> 1. Recursive sotring Algorithm
> 1. Split the array into two parts.
> 1. Recursively split until each array has only one element.
> 1. Merge these parts.

```
[18,8,22,4,11,1,9,28]
[18,8,22,4],[11,1,9,28]
[18,8],[22,4],[11,1],[9,28]
[18],[8],[22],[4],[11],[1],[9],[28]
[8,18],[4,22],[1,11],[9,28]
[4,8,18,22],[1,9,11,28]
[1,4,8,9,11,18,22,28]
```

the merge function combines two sorted arrays, it starts at the beginning of each array, and takes the smaller number between the two, and moves to the next element in that array.

```csharp
public static void MergeSort(int[] input, int from, int to)
{
  if (from < to)
  {
    int middle = (from + to)/2;
    MergeSort(input, from, middle);
    MergeSort(input, middle+1, to);
    Merge(input,from, middle, to);
  }
}
public static void Merge(int[] input, int from, int middle, int to)
{
  int LengthLeft = middle-from+1;
  int LengthRight = to-middle;
  int left[]= new int[LengthLeft +1];
  int right[]= new int[LengthRight +1];

  for (int i=0; i< LengthLeft; i++)
  {
    left[i]=input[from+i];
  }

  for (int i=0; i< LengthRight; i++)
  {
    right[i]=input[middle+i+1];
  }

  left[LengthLeft]= int.MaxValue;
  right[LengthRight]= int.MaxValue;

  int leftPointer=0;
  int rightPointer=0;
  for (int i = from; i <=to; i++)
  {
    if (left[leftPointer] > right[rightPointer])
    {
      input[i]=right[rightPointer];
      rightPointer++;
    }
    else
    {
      input[i]=left[leftPointer];
      leftPointer++;
    }
  }
}
```

### Quick Sort

> 1. Choose "pivot"
> 2. Move Smaller numbers on it's left side, larger on it's right side.
> 3. Recursively call quick sort on the left and right side.

```
[2,5,20,15,1,11,8P]
[2,5,1P,8,20,11,15P]
[1,5,2P,8,11,15P,20]
[1,2,5,8,11,15,20]
```

code:

```csharp
public static void QuickSort(int[] input, int from, int to)
{
  if (from < to)
  {
    int indexOfPivot=Partition(input,from,to);
    QuickSort(input, from, indexOfPivot-1);
    QuickSort(input, indexOfPivot+1, to);
  }
}

public static void Partition(int[] input, int from, int to)
{
  int pivot = input[to];
  int wall = from;
  for (int i = from; i< to; i++)
  {
    if (input[i] <= pivot)
    {
        //swap
        int temp = input[wall];
        input[wall] = input[i];
        input[i] = temp;
        wall++;
    }
  }
  input[to]=input[wall];
  input[wall]=pivot;
  return wall;
}
```

we can use a different algorithm to choose the pivot number, like the median value, we always want the pivot to be at the last element in the section which we are sorting.

### Comparing

we use the big O notation to denote time or space complexity. we care about the largest exponent, and we look at the worst case.

| Metric           | Selection Sort | Bubble Sort | Merge Sort | Quick Sort        |
| ---------------- | -------------- | ----------- | ---------- | ----------------- |
| Time Complexity  | O(n^2)         | O(n^2)      | O(n log n) | O(n^2) worst case |
| Space Complexity | O(n)           | O(n)        | O (2^n)    | O(n)              |
| In-place         | Yes            | Yes         | No         | Yes               |
| Stable           | No             | Yes         | Yes        | No                |

</details>

## Data Structures

<details>
<summary>
examples of basic data structures.
</summary>

### What is Data Structure

a way to store and modify data, so that it could accessed efficiently. Arrays and lists are such strcutrues, but they have limitations, such as serching for a value, which requires checking each element one by one.

### Binary Search Tree

[tree visualization tool](https://www.cs.usfca.edu/~galles/visualization/BST.html)

A tree is structure Nodes, a node has a value and it can have children nodes. in a binary tree, a node can have up to two children (zero, one, or two), each node can only have one parent, except for the **Root** node, which only has children nodes and not parent node. nodes without children are called **Leafs**. in a binary search tree, all the nodes under the left children tree have values lower than the parent node. and all the nodes in the right subtree have a value larger than that of the parent.

so to **search** a value, we compare the searched value to the root, and based on the result we continue down the correct subtree. to **insert** an element, we perform the same process, and we go down the correct path, if the value doesn't exist, we add a node as either a left or right side subtree. when we wish to **delete** a node, there are 3 possible options:

- deleting a leaf - find the leaf value, disconnect the parent node from it.
- delete a node with one child - make the child node connect to the parent node instead of the removed node.
- delete a node with two children - we need to replace the removed node with either the highest value at the left sub tree, or the lowest value at the right subtree. we usually replace the values and then delete new "leaf" (it will have at most one child node).

for all cases of deletion, we need to update the parent node and replace the reference it holds.

```cs
class BinaryTreeNode
{
  public int value{get};
  public BinaryTreeNode leftNode;
  public BinaryTreeNode rightNode;

  public BinaryTreeNode (int nodeValue,BinaryTreeNode left, BinaryTreeNode right)
  {
    value = nodeValue;
    leftNode=left;
    rightNode=right;
  }
}

class BinarySearchTree
{
  BinaryTreeNode root = null;

  public void Insert(int value)
  {
    if (root == null)
    {
      root = new BinaryTreeNode(value, null, null);
      return;
    }
    BinaryTreeNode parent = null;
    BinaryTreeNode current = root;

    while (current != null)
    {
      parent = current;
      if (current.value < value)
      {
        current = current.rightNode;
      }
      else if (current.value > value)
      {
        current = current.leftNode;
      }
      else
      {
        return; // exists already
      }
    }

    //parent can't be null
    if (parrent.value < value)
    {
      parent.leftNode = newBinaryTreeNode(value);
    }
    else
    {
      parent.rightNode = newBinaryTreeNode(value);
    }
  }

  public bool Search(int value)
  {
    BinaryTreeNode current = root;

    while (current != null)
    {
      if (current.value < value)
      {
        current = current.rightNode;
      }
      else if (current.value > value)
      {
        current = current.leftNode;
      }
      else
      {
        return true;
      }
    }
    return false;
  }

  public bool delete(int value)
  {
    if (root == null)
    {
      return false
    }
    BinaryTreeNode parent = null;
    BinaryTreeNode current = root;

    while (current.value != value)
    {
      parent = current;
      if (current.value < value)
      {
        current = current.rightNode;
      }
      else if (current.value > value)
      {
        current = current.leftNode;
      }
      else
      {
        //value doesn't exists in tree
        return false;
      }
    }

    //case 1: no children (leaf)

    if ((current.leftNode== null) && (current.rightNode ==null))
    {
      if (current == root)
      {
        root = null;
      }
      else if (parent.value < value)
      {
        parent.rightNode = null;
      }
      else
      {
        parent.leftNode = null;
      }
      return true;
    }

    //case 2: one childrent

    if (current.leftNode== null)
    {
      if (current == root)
      {
        root= current.rightNode;
      }
      else if (parent.value < value)
      {
        parent.rightNode = current.rightNode;
      }
      else
      {
        parent.leftNode = current.rightNode;
      }
      return true;
    }
    else if (current.rightNode == null)
    {
      if (current == root)
      {
        root= current.leftNode;
      }
      else if (parent.value < value)
      {
        parent.rightNode = current.leftNode;
      }
      else
      {
        parent.leftNode = current.leftNode;
      }
      return true;
    }

    // case 3: node has two childrent.
     if ((current.leftNode!= null) && (current.rightNode !=null))
     {
      BinaryTreeNode successor = getBiggestNodeFromLeftSubtree(current); // another function
      successor.leftNode=current.leftNode;
      successor.rightNode=current.rightNode;

      if (current==root)
      {
        root=successor;
      }
      else if(parent.value< successor.value)
      {
        parent.rightNode=successor;
      }
      else
      {
        parent.leftNode=successor;
      }
      return true;
     }

     return false;
  }

  BinaryTreeNode getBiggestNodeFromLeftSubtree(BinaryTreeNode start)
  {
    BinaryTreeNode parent = start.leftNode;
    BinaryTreeNode rightChild = parent.rightNode;

    if (rightChild == null)
    {
      start.leftNode = parent.leftNode;
      return parent;
    }

    while (parent.rightNode != null)
    {
      parent = rightChild;
      rightChild= rightChild.rightNode;
    }

    parent.rightNode = rightChild.leftNode;
    return rightChild;
  }
}
```

### AVL tree

The AVL tree is binary search tree that balances itself. so it won't ever have a unbalanced subtrees. a balance tree is measured by comparing the height of each subtree, if the height absolute value is larger than 1, then the tree is imbalanced. height is defined as the number of level in each subtree. if the AVL tree is not balanced, then we use rotations to correct it.

- LL : 1 (parent) - 2 (right child of parent) - 3 (right child of 2)
- RR : 3 (parent) - 2 (left child of parent) - 1 (left child of 2)
- LR : 1 (parent) - 3 (right child parent) - 2 (left child of 3)
- RL : 3 (parent) - 1 (left child parent) - 2 (right child of 1)

in all cases, we want to balance the tree

- 2 (parent) - 1 (left child of parent) - 3 (right child of parent)

**insertion** is the same as regular tree, but after each insertion, we calculate the balance and perfrom rotations. **deletion** is similar, after deletion we perform rotations as needed. **search** doesn't change the tree, so there is no need to make rotations.

```cs
class AvlNode
{
  public int value;
  AvlNode parent;
  AvlNode leftNode;
  AvlNode rightNode;
}

class AvlSearchTree
{
  AvlNode root;
  void Insert (value)
  {
    if (root == null)
    {
      root = new AvlNode(value);
      return;
    }

    AvlNode current = root;
    while (true)
    {
      if (current.value <value )
      {
        if (current.rightNode != null)
        {
          current = current.rightNode;
        }
        else
        {
          current.rightNode = new AvlNode(value)
          {parent=current};
          break;
        }
      }
      else if(current.value > value)
      {
        if (current.leftNode != null)
        {
          current = current.leftNod;
        }
        else
        {
          current.leftNod = new AvlNode(value)
          {parent=current};
          break;
        }
      }
      else
      {
        //exists
        return;
      }
    }
    rebalance(current);
  }


  bool Delete (int value)
  {
    AvlNode current = root;
    while (current != null)
    {
      if (current.value < value)
      {
        current = current.rightNode;
      }
      else if (current.value> value)
      {
        current = current.leftNode;
      }
      else
      {
        deleteNode(current);
        return true;
      }
    }
    return false;
  }

  public void deleteNode(AvlNode node)
  {
    // case 1: node is leaf
    if ((node.leftNode == null) && (node.rightNode == null))
    {
      if (node != root)
      {
        AvlNode parent = node.parent;
        if (parent.value < node.value)
        {
          parent.rightNode = null;
        }
        else
        {
          parent.leftNode = null;
        }
        rebalance(parent);
      }
      else
      {
        root = null;
      }
    }
    // case 2: node has on child;
    else if (node.leftNode == null)
    {
      if (node != root)
      {
        AvlNode parent = node.parent;
        if (parent.value < node.value)
        {
          parent.rightNode = node.rightNode;
        }
        else
        {
          parent.leftNode = node.rightNode;
        }
      }
      else
      {
        node.rightNode.parent=null;
        root = node.rightNode;
      }
      rebalace(node.rightNode);
      return;
    }
    else if (node.rightNode == null)
    {
      if (node != root)
      {
        AvlNode parent = node.parent;
        if (parent.value < node.value)
        {
          parent.rightNode = node.leftNode;
        }
        else
        {
          parent.leftNode = node.leftNode;
        }
      }
      else
      {
        node.leftNode.parent=null;
        root = node.leftNode;
      }
      rebalace(node.leftNode);
      return;
    }
    // case3: node has both children
    if ((node.leftNode != null) && (node.rightNode != null))
    {
      AvlNode successor = getBiggestNodeFromLeftSubtree(current); // another function
      successor.leftNode = current.leftNode;
      successor.rightNode = current.rightNode;
      AvlNode parentNode = node.parent;
      if (current!=root)
      {
        if(parent.value< successor.value)
        {
          parent.rightNode=successor;
        }
        else
        {
          parent.leftNode=successor;
        }
      }
      else
      {
        successor.parent=null;
        root=successor;
      }
      rebalance(successor);
    }
  }

  void rebalance(AvlNode start)
  {
    int balance = node.getBalance();
    if (balance == -2)
    {
      if (node.leftNode.getBalance() == 1)
      {
        node = LR(node);
      }
      else
      {
        node = RR(node)
      }
    }
    else if (balance == 2)
    {
      if (node.rightNode.getBalance() == -1)
      {
        node = RL(node);
      }
      else
      {
        node = LL(node)
      }
    }

    if (node != root)
    {
      rebalance(node.parent);
    }
  }

  void LL(AvlNode node)
  {
    AvlNode right = node.rightNode;
    AvlNode parent = node.parent;
    node.rightNode= right.leftNode;
    right.leftNode=node;
    if (node!= root)
    {
      if (parent.value < right.value)
      {
        parent.rightNode = right;
      }
      else
      {
        parent.leftNode = right;
      }
    }
    else
    {
      right.parent =null;
      root=right;
    }
    return right;
  }

void RR(AvlNode node)
  {
    AvlNode left = node.leftNode;
    AvlNode parent = node.parent;
    node.leftNode= right.rightNode;
    left.rightNode=node;
    if (node!= root)
    {
      if (parent.value < left.value)
      {
        parent.rightNode = left;
      }
      else
      {
        parent.leftNode = left;
      }
    }
    else
    {
      left.parent =null;
      root=left;
    }
    return left;
  }

  public void RL(AvlNode node)
  {
    RR(node.rightNode);
    return LL(node);
  }
  public void LR(AvlNode node)
  {
    LL(node.leftNode);
    retrun RR(node);
  }
  public int getBalacne()
  {
    var leftBalance = leftNode?.getHeight() ?? 0;
    var rightBalance = rightNode?.getHeight() ?? 0;

    return leftBalance - rightBalance;
  }

  public int getHeight()
  {
    if ((leftNode == null) &&(rightNode == null))
    {
      return 1;
    }
    else if (leftNode == null)
    {
      return rightNode.getHeight() + 1;
    }
    else if(rightNode == null)
    {
      return leftNode.getHeight() + 1;
    }
    else
    {
      int rightHeight = rightNode.getHeight();
      int leftHeight = leftNode.getHeight();
      return max(rightHeight,leftHeight)+1;
    }
  }
}
```

### Linked List

A linked list is a collection of connected nodes, it starts with a head node. adding a node at the head is always a constant time operation. whenever we add a value, we create a new node, point it to the current head, and set this node as Head.

searching is done by going over all nodes, inserting and deleting nodes doesn't require moving them in memory, because they are only connected by references.

```cs
class Node
{
  public int value;
  public Node next=null;
  public Node(int nodeValue, Node nextNode)
  {
    value=nodeValue;
    next=nextNode;
  }
}

class LinkedList
{
  public Node head = null;

  public void insert(int value)
  {
    head = new Node(value, head);
  }
  public bool search(int value)
  {
    Node current = head;

    while(current != null)
    {
      if (current.value == value)
      {
        return true;
      }
      current = current.next;
    }
    return false;
  }
  public bool delete(int value)
  {

    if (head == null)
    {
      return false;
    }
    if (head.value == value)
    {
      head = head.next;
      return true;
    }

    Node previous = head;
    Node current = head.next;

    while (current != null)
    {
      if (current.value == value)
      {
        previous.next= current.next;
        return true;
      }
      else
      {
        previous=current;
        current=current.next;
      }
    }
    return false;
  }
}
```

### Trie

Trie is a tree-like structure, but unlike the binary search tree, it uses more than two children nodes. the common use-case is when we want auto-completion.

each trieNode contains a map of children nodes, the value of the node is all the directions taken to reach it from the root node. in this example, the Trie is for dictionaries, so we also record whether this node is a valid word or not. the number of children nodes is the number of letters in the alphabet.

when we add a word, we follow the path of the letters in the word, and if the node doesn't exist we create it. the power of the data structure comes into play when we have multiple elements with similar prefixes. searching is done by following each part of the word one by one. Deleting requires checking whether there are empty nodes to delete in the path.

```cs
class TrieNode
{
  public Dictionary<char, TrieNode> table = new Dictionary<char,TrieNode>();
  public bool isWord=false;

  public bool hasRecord(char c){
    return table.containsKey(c);
  }
  public bool isEmpty()
  {
    return table.Count == 0;
  }
  public TrieNode followChar(char c)
  {
    if (hasRecord(c))
    {
      return table[c];
    }
    return null;
  }
  public void addRecord(char c, TrieNode n)
  {
    table.add(c,n);
  }
  public void deleteRecord(char c)
  {
    table.remove(c);
  }
}

public class Trie
{
  TrieNode root = new TrieNode();
  public void insert(string value)
  {
    char[] input = value.toCharArray();
    TrieNode node = root;
    foreach (char c in input)
    {
      if (!node.hasRecord(c))
      {
        node.addRecord(c, new TrieNode());
      }
      node  = node.followChar(c);
    }
  }

  public book search(string value)
  {
    char[] input = value.toCharArray();
    TrieNode node = root;
    foreach (char c in input)
    {
      if (!node.hasRecord(c))
      {
        return false;
      }
      node  = node.followChar(c);
    }
    return node.isWord;
  }

    public bool delete(string value)
  {
    char[] input = value.toCharArray();
    TrieNode node = root;
    TrieNode[] path = new TrieNode[input.Length];
    int pathLength = 0;
    for (int i=0; i< input.Length; i++)
    {
      if (!node.hasRecord(input[i]))
      {
        return false;
      }
      node = node.followChar(input[i]);
      path[pathLength] =n;
      pathLength++;
    }
    node.isWord=false;

    // deleting unused nodes

    int inputIndex = input.Length -1;
    for (int i = pathLength -2; i >= 0 ;i--)
    {
      if (node.isEmpty() && !node.isWord)
      {
        node=path[i];
        node.deleteRecord(input[inputIndex]);
        inputIndex--;
      }
      else
      {
        break;
      }
    }
    return true;
  }
}
```

### Hash Table

Matching Keys with index position, the size is set by us, and the hash function takes an input and returns an index. we use a design that combines an array with linkedList. so we don't care about having multiple values with the same resulting index position.

our has function can be something as simple as using two primes numbers, or by using some bit operators.

```
hash = 7;
for (character in input)
  hash = hash*31 + character
return hash
```

and we just use the module operator to get the correct index.

```cs
class HashTable
{
  private LinkedList[] data;
  public hashTable(int size)
  {
    for (int i = 0; i< size; i++)
    {
      data[i]= new LinkedList();
    }
  }

  public void insert(string value)
  {
    int index = Math.abs(hashFunction(value)) % data.Length;
    data[index].insert(value);
  }
  public bool search(string value)
  {
    int index = Math.abs(hashFunction(value)) % data.Length;
    return data[index].search(value);
  }
  public bool delete(string value)
  {
    int index = Math.abs(hashFunction(value)) % data.Length;
    return data[index].delete(value);
  }

  public int hashFunction(string value)
  {
    int hash = 7;
    foreach (char c in value.toCharArray())
    {
      hash = hash * 31 + c;
    }
    return hash
  }
}
```

</details>

## Problem Solving Techniques

<details>
<summary>
Basic problem solving algorithms and approaches.
</summary>

### Divide and Conquer

An algorithm pattern which breaks down a big problem into smaller sub problems, and does this over and over until the problem become simple, then the results are combined together. **merge sort** is an example of such algorithm, we divided the data into smaller parts, until each part was sorted (size of one), and then on the way back, we combined the solutions together to get a sorted range.

### Dynamic Programming

Dynamic programming is a method of solving problems that relies on storing the results of smaller sub problems, and then reusing those results again and again. this comes into play with the fibonacci sequence, rather than recurse each time until we reach the base condition for each branch, we calculate the results only once.

### Greedy Approach

in this pattern, we focus on the current case, ignoring all the other future steps. this makes calculations simpler, but won't always bring us to the optimal solution. the travelling salesmen is an example of a situation where this approach fails.

### Backtracking

Backtracking uses brute force to go over all the options, and find a solution, or going over all the possible solutions and finding out that there are none.\
The **N Queens problem** can be solved with backtracking. the goal is to set N chess queens on a N \* N sized chess board without having any queen threating another one. for each row, we start by putting a queen and determing if it's a valid location, if it is, we move to the next row, if not, we try another option in the row, and if we don't have any options, we move back to the previous row.

</details>

## Complex Problems

<details>
<summary>
Basic Problems
</summary>

### Knapsack problem

The **Knapsack** (or **Thieves bag**) problem is an optimization problem, we items of different weights and values, and we want to fit them all into a bad with limited capacity.

| Item | Weight | Value |
| ---- | ------ | ----- |
| A    | 1      | 2     |
| B    | 2      | 2     |
| C    | 3      | 5     |
| D    | 4      | 6     |

with the bag size being 5.

we use dynamic programming, determing what is the optimal composition of items for each bag size until the given size.

we start with empty first row and column. then we calculate for each bag size and item the optimal worth.
the pseudo code is:

```
for i in [0:n] do:
  m[i,0] = 0
for i in [0:b] do:
  m[0,i] = 0

for i in [0:n] do:
  for c in [0:b] do:
    if c < w(i) then:
      m[i,c] = m[i-1,c]
    else
      m[i,c] = max(m[i-1,c], m[i-1,c-w(i)] + p(i))
```

| value | weight | 0   | 1   | 2   | 3   | 4   | 5   |
| ----- | ------ | --- | --- | --- | --- | --- | --- |
| 0     | 0      | 0   | 0   | 0   | 0   | 0   | 0   |
| 2     | 1      | 0   | 2   | 2   | 2   | 2   | 2   |
| 3     | 2      | 0   | 2   | 2   | 4   | 4   | 4   |
| 5     | 3      | 0   | 2   | 2   | 5   | 7   | 7   |
| 6     | 4      | 0   | 2   | 2   | 5   | 7   | 8   |

to determine which items we take , we start with optimal solution, and check if the bag has enough space for the item, and if taking it is beneficial by comparing to the cell above it, if it's larger, we add it to the bag and decrease the free space

```
s = {} // empty set
c = b // free size
for i = [n:0] do:
  if (c - w(i) >= 0 and (m[i-1,c] < m[i-1,c-w(i)+p(i)] )) do:
    s <- {i} // insert item to sack
    c = c-w(i) // reduce empty space

return s //best worth elements
```

### Kruskal's algorithm

solving a spanning tree (a tree where each node is connected to more than one other node), by passing the through all nodes, while avoid cycles.

we use the greedy approach, trying each edge and avoid cycles. for N nodes, the minimal number is N-1 edges. we use disjoined sets, and we create components. it changes the root of each edge.

example:

| source | destination | edge cost |
| ------ | ----------- | --------- |
| A      | B           | 4         |
| A      | C           | 3         |
| B      | C           | 1         |
| B      | D           | 8         |
| B      | E           | 8         |
| C      | D           | 6         |
| D      | E           | 2         |

steps, starting with the smallert

- {B:C} (1)
- {B:C} {D:E} (1,2)
- {C:{A:B}} {D:E} (1,2,3)
- {C:{{A:B},{D:E}}} (1,2,3,6)

</details>

</details>
