<!--
// cSpell:ignore Kruskal
 -->

# Software Engineering Fundamentals From Scratch in C Sharp

<!-- <details> -->
<summary>

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

## Problem Solving Techniques

## Complex Problems

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### 0-1 Knapsack problem

### Kruskal's algorithm

</details>

</details>
