# Effective STL: 50 Specific Ways to Improve Your Use of the Standard

## Introduction

## Chapter 1: Containers

Most of the world uses the STL for the containers, but it has so much more: iterators, algorithms and function objects (functors). But containers reign above in popularity.
But even the STL containers aren’t perfect, and must be used with case.

### Item 1: Choose your containers with care.

Stl has many types of containers:
standard sequence containers (vector, string, deque, list)
standard associative containers (set, multiset,map, multimap)
Non standard sequence containers (slist - singly linked list, rope - heavy duty ‘string’ class)
Non standard associative containers (hash_set,hash_multiet, hash_map, hash_multimap)
vector<char> as a replacement for string
Vector as a replacement to the associative containers
And some containers that aren’t in the stl standard
non-stl containers: arrays, bitset, valarr, stack, queue, priority queue.

We should know our options before choosing which fits us best, different considerations imply different containers to use. Vector, list and deque all offer the same functionalities, but have different complexities for the operations, vector is the default one to use, but if we know we will work mostly with the head and tail of the data, then we should choose a deque, and if we know that we are going to do middle insertions, then we should choose the list.
But there are more stuff to consider.

Another way to categorize stuff:
Contiguous memory containers (array based)- vector, strine, deque, also rope.
elements are stored together, and if we insert or remove, we need to shift the entire memory.
Node based containers: linked list (list, slist,maps, sets, hash variations, trees...).
Each element is stored separately, and contains the direction of the next element.

Questions to ask ourselves:

Question
Answer and implication
“Do you need to be able to insert a new element at an arbitrary position in the container?”
Yes: Sequence containers
“Do you care how elements are ordered in the container?”
No: Sequence containers
“Must the container be part of standard C++?”
Yes: Avoid hashed variants, slist and rop
“What category of iterators do you require?”
Random Access: vector, deque string, maybe rope
BiDirectional: avoid slist, avoid some hashed
“Is it important to avoid movement of existing container elements when insertions or erasures take place?“
Yes: avoid Contiguous memory containers.
“Does the data in the container need to be layout-compatible with C?”
Yes: you’re stuck with vectors. Tough luck.
“Is lookup speed a critical consideration?”
Yes: check hash\_ variations. Then sorted vectors and the associative containers.
“Do you mind if the underlying container uses reference counting?”
Yes: avoid string, rope, consider vector<char>
“Do you need transactional semantics for insertions and erasures?”
Yes: use node based containers, if multi element insertions is required then use list.
“Do you need to minimize iterator, pointer, and reference invalidation?”
Yes: use node-based memory
“Do you care if using swap on containers invalidates iterators, pointers, or references?”
Yes: avoid string.
“Would it be helpful to have a sequence container with random access iterators where pointers and references to the data are not invalidated as long as nothing is erased and insertions take place only at the ends of the container?”
Yes: special case to use deque!
And many more...
Read the book!

### Item 2: Beware the illusion of container-independent code.

Wel like to think that we can write code that will work with all containers. We can’t.
Iterators behave differently in different containers, associative and sequence containers aren’t the same.some containers allow invaliding iterators during a for_each run, some don’t, some allow for capacity, some don’t.
Some have compatibility with c, some don’t.\
We can’t get all the strenghts, and we should strive to.

If we really are afraid of choosing a container class and we want our code to work in case we change the container class, we should encapsulate our class with a wrapper class (typedef, encapsulating, whatever).

We can use typedef and composition, if we use composition, we have a inner class that can only be accessed by our class, so all changes are contained in one place.

We can’t write container independent code.
We can provide an interface that allows our clients to do so.

### Item 3: Make copying cheap and correct for objects in containers.

Containers hold copies of elements, not the elements themselves.
STL loves copying, and it uses the copy constructor and athe copy assignment operator, the compiler generates those for whatever class that needs them.
We want these operations to be as efficient as possible and we want to avoid slicing.
If we have a container of base class objects and we try to insert a derived class, we will get slicing behavior (which we don’t want!)
We should never have a container of base class and add a derived class.
We can have a container of pointers, which is both immune to slicing and cheap in terms of copying, although this might run into some problems with STL stuff. We can have a smart pointer container.

The stl tries to avoid unnecessary copying and object creation as much as it can.
Widget w[100];
Creates an array of 100 widgets. All constructed.
We can do better
vector<widget> w;
w.reserve(100);
We created a vector with zero active widgets, but enough space to hold 100 of them. Now the compiler could use this empty space to directly construct them if we need.

### Item 4: Call empty instead of checking size() against zero.

size() can be N linear time, while empty() can be reduced to O(1). it’s worth the pay of writing an extra function.
An explanation about why list can’t have constant time size (answer: because of splice).

### Item 5: Prefer range member functions to their single-element counterparts.

(a reminder about the assign member function, which can do stuff that operator= can’t).

Range operations allow us to avoid loops, everything is handled internally.
We can also use the copy version

(start where, finish where, put where)
Or the range version of insert:
v1.insert(v1.end(), v2.begin() + v2.size() /2,v2.end());

We can use copy, but if we use iterators, we should use range version members instead.

Less code to write
Easier code to read

They are also more efficient:

Three problems with explicit looping
Function calls. Maybe it can be inlined, maybe not.
Moving the position each time,
Memory allocation

Range member functions know the amount of items in advance, and they can make room avoid repeating calls. This is most prevalent is sequence containers (and contiguous memory ones), but can also pop up in associative containers.

Types of members to note. I.e. are great candidates for range member functions:
Range construction(inputIterator begin, inputIterator end).
Range insertion(iterator position,inputIterator begin, inputIterator end).
Range erasure(inputIterator begin, inputIterator end).
Range assignment(inputIterator begin, inputIterator end).

Remember that push_back() and push_front() probably also use inset, so if we have a loop with them, we should change it.
Note that erase behaves differently between associative and sequence containers.

### Item 6: Be alert for C++’s most vexing parse.

What we thought to be a call to a constructor is actually a function declaration.
Whaaaaaaaaaaaaaaaaaaaaaaaaaaaaat?
Let's work from a different angle

C++ expression
Human words
int f (double d);
A function f, takes double d, returns int
int f (double (d));
Same, we ignore the parentheses
Int f (double);
Same, no parameter name, we are allowed to do this (remember ++ and --)?
Int g(double (\*pf)());
g takes a pointer to a function that returns a double, the function parameter is pf;
Int g(double pf());
Same, but we don’t use pointer syntax, we can do this in C and C++
Ing g(double ());
Same,we haven't named our parameter, but it’s still a function pointer
list<int> data(istream_iterator<int>(dataFile), istream_iterator<int>())
A function called data, which returns list<int>, and takes a parameter names datafile of type Istream_iterator, the second parameter is an unnamed pointer to a function that takes nothing and return istram_iterator<int>.

Isn’t this some goddamn magic?

This is how it should look

This is like saying the first parameter is an expression by itself.

Not all compilers allow this.
We might need to actually name them outside, like some simpleton.

### Item 7: When using containers of newed pointers, remember to delete the pointers before the container is destroyed.

Stl Containers know how do take care of their members and call the destructor, but if the members are pointers, they don’t know to call delete, this is up to the programmer.

We can stick the delete behavior inside a ‘for loop’, but it isn’t exception-safe.
We can use an stl::for_each loop, but this requires a function object that does the deletion, and this objejs needs to know what it’s going to delete.
We can move the declaration into the template, and have the compiler instatinate accordingly. But this still isn’t exception safe. For this we need smart pointers. We can’t have containers that automatically delete their pointers, no matter what. We need to think and decide.

### Item 8: Never create containers of auto_ptrs.

COAP - containers of auto_ptr.
Stl containers shouldn’t allow auto pointers (auto_ptr<>) as arguments, but not all compilers catch this.
We might want to try to create such containers, because we don’t want to have manually check the pointes and destroy them.

When we copy an auto_ptr, we change the value. That’s the behavior of the auto_ptr idea.
This is very dangerous, like in sort algorithms, once the auto_ptr is copied, the original becomes NULL, and we might lose the data when it goes out of scope.

(advanced smart pointers are probably ok).

### Item 9: Choose carefully among erasing options.

There are different ways to remove elements from containers, some times we should use erase, sometimes remove…
Contiguous memory containers: vector<>, deque<>, string:

For lists we can use the member function .remove()

For associative containers (sets,map, multi_set,multi_map) we can’t call .remove, and we should call .erase().

Sequence containers:
c.erase(remove_of(c.begin(),c.end(), bad_value),c.end());
lists:
c.remove_if(bad_value);
Associative container:
Easy way:
remove_copy_if(c.begin(), c.end(),inserter(goodValues, goodValues.end()), bad_value);
c.swap(goodValues)
We copied only the good values,and then we swapped our contents with the other container.
The better way, using a loop to replicate remove_if. We need to be careful of invalidating our iterators

### Item 10: Be aware of allocator conventions and restrictions.

Allocators are weird.
There are stuff allocators can’t really do.
Allocators of the same type must be equivalent - this is why we can have list splicing.
All sorts of stuff i don’t really get.

### Item 11: Understand the legitimate uses of custom allocators.

I don’t get it.

### Item 12: Have realistic expectations about the thread safety of STL containers.

We would want containers that support multiple readers to the same containers, and containers that support multiple writers to different containers of the same class.

Old c style is to get a mutex (getMutexFor, releaseMutexFor)
And OOP RAII style is to have a mutex object, which is destroyed/released at the end of the local scope and has exception safety.
STL doesn’t offer thread safety. We need to create it.

## Chapter 2: vector and string

The two most common STL containers are vector<> and string.they replace the array and the char\* of base c.

### Item 13: Prefer vector and string to dynamically allocated arrays.

When we use a dynamic allocation, we take some responsibilities on ourselves:
We got resources from the system, so they must be released (call delete) to avoid leaking.
The correct form of delete must be used, new-delete, or new[]-delete[], and the even worse case of placement new.
We must call delete once, and only once.
Failure to comply to one of the responsibilities will cause memory leak, memory corruption or some other sort of problems. We would prefer not to shoulder this burden if we can avoid them.

In any case we have a dynamic allocation (we write new[]), we should consider using vector or string instead (string when we are using characters, vector<T> otherwise).
Both classes manage their own memory, take care of resizing when it’s needed, and destroy the elements the hold without our help.
They also are sequence containers, which means we can use all our cool stl algorithms and functions, like end(), begin(), size(), and get nice iterator and nice functionality for string. We could achieve the same result with plain old C arrays, but the new code is cleaner.

We might suffer some problems if we previously had a string implementation that used reference counting and now we are copying around entire books for each function call, but this is also part of the string implementation in most cases.
Sometimes we want to avoid reference counting (this can be more costly than actually copying it when we are in a multi threaded world), in these cases, we can:
Find a way to shut down reference counting (via some preprocessor macro).
Use a different string implementation.
Use vector<char> which can’t be reference counted - this means giving up on some fancy string behavior, but that’s life for you.

### Item 14: Use reserve to avoid unnecessary reallocations.

The vector and string classes are contingent memory containers, and grow dynamically in size. Whenever a container gets more memory, it needs to copy all of it’s elements around, invalidate iterators and other stuff that is expensive and harms performance.
We can use the reserve function to preallocate memory and minimize the amount of copying the container does.
Functions:
size() - how many elements are actually in the container
capacity() - how many elements can the container currently store without resizing.
resize() - forces the container to change the amount of elements, if the argument is lower than the current size, some elements will be destroyed, if we ask for more than the capacity, the container will reallocate itself.
reserve() - request more memory without creating more elements.

The key to avoid reallocations is to have the correct size from the start, which we can get by calling reserve once, or if we must avoid reallocations, only allow insertions when the size() is less than the capacity.

### Item 15: Be aware of variations in string implementations.

Special note: what does ‘sizeof(string)’ give us? Is it the same as the char\* size, what about member variables? Does it change between different implementations of string? YES.

Four different implementations of string from different libraries:

So the takeaway is that we need to be careful with the assumptions we use with the string class. We really can’t know what kind of implementation we are using.

### Item 16: Know how to pass vector and string data to legacy APIs.

Old C legacy api requires passing arrays and char* pointers, so we can’t ignore them completely and focus only on the sweet-sweet use of vectors and strings.
To pass the data of our string to code that requires char*, we use the .c_str() function.
To get the address of the first element in a vector, we use v&[0], this is the same as getting the address of it.
This doesn’t work if we have an empty vector, in which case v[0] wouldn’t yield a decent reference.
We can’t use &[0] for strings, as string aren’t necessarily contingent in memory, usually they are. But implementations can do something else. If we really need a contingent memory representation, we can use a vector<char> class (hoo ra! We found its’ use!).

Some other details and shticks

### Item 17: Use “the swap trick” to trim excess capacity.

We might have had a vector with many elements before, but later, the vector holds much less elements, and we would want to reduce the memory requirement (change the capacity).
Here the ‘shrink to fit’ trick comes into play.

This all happens in an expression scope. The vector we create is a temporary and unnamed vector.
This vector uses a constructor that takes two iterators from the original vector, and once it’s initialized, we swap the contents of the two vectors, now the new vector has the large capacity, and the old one has the reduced capacity. At the end of the expression the new vector is destroyed,
We can also do the same trick with strings

This usually works, but it depends on how the library implements the code.
We can also do the same idea to erase everything and minimize the capacity:

### Item 18: Avoid using vector<bool>.

(this is why we use bitset).
vector<bool> isn’t really.
It’s not and STL container, and it doesn’t hold bools.
The reason it’s not an stl container is that it doesn’t properly support the operator[] to return a reference to a boolean element.
We want our STL containers code to support
T \*p = &c[0];
(a pointer to the first element of the container).
This doesn’t happen with vector<bool>.
vector<bool> doesn’t contain booleans, that would be a waste of 7 bits per element. Instead, we store a representation of the bits in some other type of data.
When we ask for c[0] in a vector<bool> we don’t get an actual reference to the data, we get a proxy class. And this is why the above code isn’t legal.

vector<bool> is still in the standard because in earlier times they wanted it to allow specialized dynamic bitset, it didn’t work, and now it’s an ‘old shame’ we carry around.

If we know the size of the container at compile time, we can use bitset<>, which offers memory efficiency, but not dynamic resizing and doesn’t allow the pointer behavior.
If we really want a dynamic container of boolean values, we can use deque<bool>, which doesn’t save us the memory costs, but behaves as we expect it to behave. And is a valid dynamic container.

## Chapter 3: Associative Containers

### Item 19: Understand the difference between equality and equivalence.

Associative containers use equivalence, not equality.
Equality is based on the == operator, equivalence uses the < operation.
(unlike Java, we can redefine the == operator, so we have no need to .isEqual() or .Equals() methods).
Equivalence is the relative position in a sorted array, (Comparable interface in java). So w1 and w2 are equivalent if :
!(w1<w2) && !(w2<w1)
If w1 isn’t smaller than w2, and w2 isn’t smaller than w1, then they are equivalent.
In the general case of associative containers, the ordering is done with the some predicate accessible via the key_comp() function of the container.

The return type of key_comp() is a function, which then takes x,y as arguments.

When equivalence and equality aren’t identical, we have a problem.
If we try to insert two equivalent elements into the same set container, only one will exist.
However, if we try to .find() them, we will be able to find only one of them.
So we can’t insert B because it’s equivalent to A, but we can’t search for B and find A, because they aren’t equal. How annoying!
The result depends on which type of ‘find’ we use.

Also the bit about sorted containers, but that comes later.

### Item 20: Specify comparison types for associative containers of pointers.

If we store pointers, we should give the container a comparison function to work with, otherwise it’ll use pointer addresses.
(we shouldn’t use loops, but we should use range functions).

The set actual arguments are the type of elements, the comparison predicate type, and an allocator. This ### Item is about the comparison predicate type.
So if we want our container of pointers to work, we should pass it a functor. Like this one:

Which is obviously a great candidate for a template class,
And now for the printing part, we can use an old loop and dereference the iterator twice (onec to get the pointer, and another time to get the string), we could write a print function that takes an iterator, of write a generic dereference() function, whatever.
The point is that we want a comparison type in our associative sets.
Notice that we need a comparison type, so an function pointer won’t be enough here.

### Item 21: Always have comparison functions return false for equal values.

When we write a comparison type (a functor!), we need to make sure that if ws1==ws2, then ws1 < ws is false.
Why? Because of how key_comp() works.
!(w1 < w2) && !(w2 < w1).
For equivalent elements, we want both sides to be false, and negated to true, as true && true is true.
However, if w1<w2 would return true for equivalent elements…
!(w1 <= w2) && !(w2 <= w1).
Now both sides initially evaluate to true, then false && false is false, and we end up with undefined behavior.
This is very dangerous if we try to reuse old code, maybe we want descending order instead of ascending order, so we copy our old factor and simply negate the result? This careless act made our container corrupt. As the negation of < isn’t >, it’ >=. Tough life.

This hold true even if our containers allow duplicates, like multiset or multimap.
The key issue is that associative containers define ‘strict weak ordering’, which dictates that the result of two equal value needs to be false.

### Item 22: Avoid in-place key modification in set and multiset.

Associative containers are ordered containers. The ordering isn’t important to the user, but they are ordered nonetheless (they are implemented with a binary tree, after all).
This means that if we change the key, we are corrupting the order and the data structure.
Maps and multimaps are ‘immune’ to this problem because map<K,V> stores pairs of <const K,V>.
In a set the values aren’t const. The only thing that matters is that the order is kept.
So if have a set of elements with three data members, and only the first member is used to hold the order (is part of the compare type), then nothing should prevent us from changing the rest of the members, the c++ standardization committee allows us to change anything, and the other two members are simply riding along with the first member.
(should this also be true for maps? The book says maybe, but the standards committee made its decision, and we need to stick with it).

This non-const means we can change our data and still compile, as long as we don’t change the actual part that determines the order (the ‘key’ part), we are fine.

We can have some implementations that protect our entire set/multiset from being changed. This isn’t part of the standard, and there’s no way to know if the library we use supports this or not.

If we want to be sure we are portable, we should always use const_cast<>() in this situation.
And we cast to a non-const reference of the type, by the way. This is important. If we don’t, we end up with this code:

This code actually calls the copy constructor and creates a temporary object, which is then used in the member function. Casts to other objects use constructors (or conversions operators), so we need to cast to a reference.

Now with maps and multimaps, which are defined as having a pairs<const K, V>.
We can’t use const_cast<>() to remove to constiness of the key element, so our always works and always safe way to do so it by creating a copy of the pair (without the const part), modify the copy, remove the original pair, and reinsert the copy.
Always works, always safe.

### Item 23: Consider replacing associative containers with sorted vectors.

When we want quick lookups, we use sets and maps. But if we really want fast lookups, we need the non standard hashed containers. But even in the realms of quick access, the sorted vector might be better than the map containers.
The reason? Mape containers are generalized and use binary tree structure, it does well with insertions, deletions, and look ups. But if we know that our programs will do a lot of lookup and very few deletions or insertions after the initial stage, we can get better performance with a sorted vector. The vector class has lower overhead per element,and it has increased memory locality (the contingent memory part). Even the binary_search algorithm works better in this case.
We do pay extra costs in this situation when we delete or insert an element, as the entire vector needs to be resorted (rather than just balanced), and sometimes we need to do an entire copy of the vector if the size reaches the capacity. However, if we know our program does mostly lookups, and hardly any insertions/deletions, we might be getting a sweet deal here.

If we change a map to a sorted vector, we need a bit more work, as we need to sort the element according to the keys, and each element is pair<K,V> (the vector itself is vector<pair<K,V>>, not const). We also need an additional lookup versions, one for taking a pair, and one for taking Key element and a pair, and one for a pair and a key element.

### Item 24: Choose carefully between map::operator[] and map::insert when efficiency is important.

The [] operator is easy to use, it allows us to get a value from a map, update the value or even create a new entry! It’s that simple.
Unfortunately, it also means the container does a lot of behind-the-scenes work to allow this to happen.
The [] operator returns a reference, and if the key doesn’t exists, it creates an empty pair (without adding it to the map yet), and then recreates the actual value element and assigns it to the entry, which means the older (default) value is destroyed.

This whole mess of a work can be reduced with a straightforward call to insert and a temporary object
Typdedef map<int, widget> IntWidgetMap;
IntWidgetMap m;
m[1]=1.5;
m.insert(IntWidgetMap::value_type(2,1.5));

The later form is much more efficient.
However, for update, the case is reversed. Operator [] is more efficient, while insert() creates a new pair (destroys the earlier pair), and then modifies the new pair, again.

Here is an example that does both:

The book has an implementations, it uses lower and upper bound and the hint form of insert.

### Item 25: Familiarize yourself with the nonstandard hashed containers.

STL didn’t have at first hash variations of containers (sets and maps), but since c++ 11 it does, and they are called unorderd_set,unordered_map and the multi_set and multi_map variations.
(the ### Item was written before, so who knows).

Unordered containers require a hashing function (function type), as well as the comparator object (like before) and the allocator.
Also some other stuff about the pre-official suggestions for the hashed container and how they maintain the hit ratio and the number of bins.

## Chapter 4: Iterators

### Item 26: Prefer iterator to const_iterator, reverse_iterator, and const_reverse_iterator.

Each container offers the following: iterator, const_itertator, reverse iterator, and const_revere_iterator.

Non const - can modify
Const - can’t modify
Start to end (front to back) traversal
iterator<T>
const_iterator<T>
End to start (back to front)
reverse_iterator<T>
const_reverse_iterator<T>

Pretty self explanatory.
And these are the signatures for common insert and erase operations in vector<T>

Note that the all ask for a normal iterator, which apparently has privileges that other operatore don’t have.
Here is a existing conversions diagram of the iterator types:

Any non const iterator can become a const_iterator.
Any forward traversing iterator can become a backward traversing iterator.
Any reverse iterator can become a regular iterator with the .base() function.

iterator
const_iterator
reveres_iterator
const_reverse_iterator
iterator

- yes
  yes
  yes(two step)
  const_iterator
  no
- no
  yes
  reveres_iterator
  Yes (base()
  Yes (two steps)
- yes()
  const_reverse_iterator
  no
  Yes (base()
  no
-

It’s always easier to get to whatever iterator you want if you start from the regular iterator.

The == operator is usually defined as a global function, not a member function.
This allows both arguments to go through implicit casting and conversions, but if it’s not, we can see annoying bugs that require use to switch the order of the arguments for the code to compile. Even if the operator were implemented correctly, we should still minimize mixing of iterator types. And we should use const_iterator only when we must, and avoid it at other times.

### Item 27: Use distance and advance to convert a container’s const_iterators to iterators.

Sometimes we have a const_iterator, but we really need a regular iterator. We can’t really use casting, because iterator<const T> and iterator<T> are different classes, and const_cast doesn’t work. Static_cast won’t work either, again, the classes are different.
(actually, they might compile in some cases, like vector<T> or string, but not always, and not for reverse operators, this is playing games with the compiler, and the compiler always wins).
However, there is a safe and portable way to get a regular iterator from a const_iterator, as long as we have both the iterator and the container.
The trick is to use the advance function and the distance function:
Get the distance of the const_iterator from the start position, and then advance a regular iterator that many steps. Seems pretty simple!
This is the code we would want to write:

But we can’t, the problem is with the .distance() function.
It’s actually a template function, as all the function in the std are…

The argument are two inputIterators of the same type, but we pass it two iterators of different types, and now it doesn’t know how to solve the ambiguity of which function to call.
The solution is to tell the compiler to use the version of the const_iter, and that will mean the non_const iterator will be converted to a const_iterator, and now we will get the correct behavior:

The efficiency of the technique is limited by the container and iterators.
If we are using a contingent container with random access (vector, string, deque), it’s constant time, if not, linear time. And if this is a problem, then maybe it’s time to rethink our design.

### Item 28: Understand how to use a reverse_iterator’s base iterator.

If we call base() on a reverse iterator, we are supposed to get the regular iterator, seems simple, but what does it really mean?

Note that the find() function takes the rbegin() and rend() as start and end arguments,
So, we got a reverse iterator is pointing to the 3rd element, but calling base() puts it on the 4th element? What gives?

Let's take a moment to think about insert. When we call insert, we want to put the next element ‘before’ the current one. So if we use insert and .end(), we want the new element to be ‘before’ the out of bounds element, obviously. And if we use insert and .begin(), we want our new element to the first one, element, not to come after ,begin(). Seems about right.
Insert doesn’t take reverse iterators as argument (nor does erase).but if it did, we would want the new element to be positioned ‘before’ the current one, but in the reversed order, so calling insert on a reverse iterator should produce behavior like ‘insert after’ on a regular iterator.
The base() function was designed to support the insert function, so if we have a reverse iterator, and we call insert on the base() outcome, we get the wanted behavior!
For erase, things are a bit different, but after all the problems, the correct behavior is to first increment the reverse iterator, then call base() and then delete.
So in general, be very careful with reversed iterator.

### Item 29: Consider istreambuf_iterators for character-by character input.

Some examples of copying from a file using istram_iterators<char>, the first example doesn’t work because it fails to copy whitespace characters.
The 2nd example disable the skipping of whitespace in the input file (.unsetf(ios::skipws)), but it’s awfully slow.
The correct behavior is to use the istreambuf_iterator<char> iterator insead.

The corresponding output iterator is ostreambuf_iterator. Both are always a better choice when dealing with unformatted character by character input/output.

## Chapter 5: Algorithms

There are many, many, many algorithms: the book says over 100 of them, as opposed to about 15 containers. So while the containers are fairly easy to remember, the algorithms aren’t, and it’s a shame, as they do great things.

### Item 30: Make sure destination ranges are big enough.

Stl containers can usually resize themselves. This works when inserting objects into them, but if we aren’t telling the code we are expanding the container, some algorithms can mess up our containers.

In this example, we try to transform our elements into something else (basically, a mapper function?), but transform doesn’t know it’s supposed to create new elements, it assumes that it has the correct place to assign the transformed values.
The code will try to do this:
_(results+results.size+0) = transmogirfy(_(values+0);)
_(results+results.size+1) = transmogirfy(_(values+1);)

The first line already produces a bug, the end iterator points outside of the range, we can’t safely dereference it! Even it was possible, the the data after the end element is definitely not part of our container memory, this is a bug.
If we would have wanted to add elements, we should have said so:

Now we have a destination iterator that knows how to add elements, and this will compile if the container offers a push_back() operation, the same with front_inserter and push_front. And if the code doesn’t compile that great, because we avoided a really nasty runtime bug.
In a general note, back_insterte is prefered, because it maintains the order from the source container. If we want to maintain the order and still have the elements at the start, we can combine the front inserter and the backwards traversing iterators:

For the general case we can use the custom inserter(container, position) and put the new elements wherever we want. It doesn’t matter, as long as we requested that new elements be inserted, and not overwrite the existing ones.

If our container is a contingent memory container (vector/ string, deque), then perhaps we could have used the range assign or range insert to reduce reserving behavior and copying around stuff. Or at least reserved the needed space before hand and avoided reallocations when using the inserter.

All of these examples are just to make one point: if we have an algorithm that changes elements in a container, we need to be sure we have elements to change.

### Item 31: Know your sorting options.

Qsort is great, and sort is even better, but sometimes we don’t need all that power.
If we want to sort only some elements, then a partial sort would be enough,
If we don’t care about ordering of those N elements, just that they are collectively the top elements, we can use the n_th element algorithm (which is like the pivot algorithm, i guess? We pivot all our elements around the nth element.
Some sorting algorithms are stable, and some are not. Sort, nth_element and partial sort are unstable algorithms, while std::stable_sort is as stable version of sort
The nth_element can be used to get the median or some percentile.
Sometimes we don’t want a sort, we want a filter, some way to get only the elements that satisfy some condition,this is the idea of partitioning the data, it puts all the elements in the beginning of the container. The return value is the new end element iterator - the first element that doesn’t match the criteria.
The sort, stable sort, partial sort and nth_element algorithms only work on random access containers. Sets and maps are always ordered, but to sort a list we can do some tricks.list itself offers a sort member function, but for the other sorting algorithms we need some other stuff, like copy-sort-copy back, or copy the iterators into a vector and sort them.

In terms of performance, the ranking is:
Partition
Stable_partition
Nth_element
Partial_sort
Sort
Stable_sort

### Item 32: Follow remove-like algorithms by erase if you really want to remove something.

The remove algorithm is weird, it doesn’t really eliminate elements from a container. It can’t. It doesn’t know the container, so how could it?

It also doesn’t change the number of elements in the container.
What remove really does is copy elements around, it writes the ‘good’ elements instead of bad elements, and returns an iterator to where the ‘new’ end of the container is. It does compacting,

If we really want to remove elements, we should call erase after calling remove. Erase can remove element, because it’s a member function of the container.

We call erase with two iterators, and it erases everything between them, so the end iterator is the .end() argument, and the start iterator is the result of our ‘std::remove()’ algorithm. That’s what we actually wanted.

This behavior should be ingrained into our minds, and in fact, stl::list<> has a member function .remove() that does it. But that’s the only container that offers this behavior.

This same behavior is shared with the std::remove_if and std::unique algorithms. So if we call them, we should enclose them in an erase function call.

### Item 33: Be wary of remove-like algorithms on containers of pointers.

We need to be very careful if we have a container of pointers.
If we use ‘remove’ behavior on pointers, we overwrite the pointers, and this might lead to resource leaks.

In the example above, we call remove if on the container, which arranges the pointers so that instead of pointing to B and C, they now point to D and E, and the return value is the 4th element pointer iterator., however, the 4th element doesn’t point to B or C, it still points to D!
If we call ‘erase’ (as we should),we simply ‘resize’ the container, the data is already lost.
C++ simply destroyed pointers, it doesn’t call ‘delete’ on them (if it would, we would have many problems), so we have to manually delete and nullify them ourselves, before we can remove the the iterators from the containers.
So instead of ‘remove_if’, we need to write a custom function that ‘deletes_if’, and then set the pointer to some value that indicates it’s no longer valid (nullptr, or something else), because once we deleted the data, we can’t check for the condition anymore!
We can use smart_pointers (shared_ptr) as elements of the container instead of raw pointers, but it must support implicit conversion to raw pointer.

### Item 34: Note which algorithms expect sorted ranges.

Not all algorithms work with all types of containers and iterators, for example: remove algorithms require forward iterators and the ability to assign through those iterators, so they can’t use input iterators, or be applied to maps, multi_maps (and usually sets and multi_sets also aren’t possible). Many also, many sorting algorithms require random access iterators, so they don’t fit with the list container (luckily, list has member function sort).

Even if our containers can work with the algorithms, some of the algorithms require a sorted range, and will cause undefined behavior if they are called on an unsorted range:

Note that ‘unique’ has a ‘remove-like’ behavior,which means it pushes the ‘duplicates’ to the end of the list and returns an iterator to the first duplicate value (the iterator right after the final unique value). So we need call to ‘erase’ afterwards.

The algorithms should be given the same comparison object that was used to sort the values, otherwise they will think they have an unsorted range.

Other than ‘unique’ and ‘unique_copy’, the algorithms above use equivalence (comparison), rather than equality

### Item 35: Implement simple case-insensitive string comparisons via mismatch or lexicographical_compare.

First of all, be careful when working with languages other than english, as strcmp wasn’t designed for that.

We have a function that casts into unsigned char, and then calls classic C to lower, and then then we compare the two characters.

To make this into a string function we first need to compare sizes (we can’t count on the null terminator in std::string). And according to that we decide on the order, and then we use the ‘mismatch’ algorithm to go over all the elements , and then we check.

So we have a basic ‘char vs char’ comparison, which casts both elements to unsigned char, calls ‘tolower’ and then returns the result.
We have a function that goes over all the elements in the string with a ‘mismatch’ algorithm: we negate the result of the comparison function with ‘not2’, because we want the first mismatch (result of zero is good in our comparison, but predicates think it’s bad), and then we check to see if we reached the end of of both strings and return the result of the last comparison (which we call again), in order to avoid going out of the container, we need to make as many comparisons as there are in the shotred containers, so we need to determine the order before passing the call to the strcmp function.

Another version is to call an stl algorithm called std::lexicographical_compare and pass it a boolean predicate (this means we get yes/no answers, rather then 1/0/-1),

Note that we pass the ending iterators of both ranges,and we pass our function to compare between them.

Actually, if we know our strings are normal english, and they don’t have any embedded null values, the easiest way is also the simplest way:

### Item 36: Understand the proper implementation of copy_if.

There are many algorithms with ‘copy’ in the name, but none of them are ‘copy_if’,

This is weird, as it means we have to to this on our own,

First, an ‘not-quite-right’ example:

This implementation uses ‘remove_copy_if’, it takes a predicate and negates it.
Stl doesn’t allow “copy anything that is true’, but it allows “copy everything except when this is not true”.
The problem is that stl want as function objects, and this code requires an ‘adaptable’ function object, which we can’t always remember to pass.
So the following code is what we actually need, although it’s really ugly:

The return value is the ‘end iterator’ of the destination: better readability is:
{
while(begin != end)
{
If (p(*begin))
{
*destBegin = \*begin;
++destBegin;
}
++begin;
}
return destBegin;
}

The book even suggests we stick this code into our stl related library (our local extensions of the std namespace?) and use it.

### Item 37: Use accumulate or for_each to summarize ranges.

We sometimes have a range of elements, and we want to reduce it into a single element, it can be the size, the sum, the maximum value (or the minimal value), etc…
All of these are called summary operations, and we should use an accumulator to get this result.
The accumulator doesn't live inside #include <algorithm>”, it lives in “#include <numeric>”, along with three other numeric algorithms,: inner_product, adjacent_difference, partial_sum,

The accumulate algorithm has two forms: one for taking a pair of iterators () and an initial value (make sure the initial value fits the return value, so 0.0 and no 0 if we want a double result).
Because the algorithm only takes input iterators, we can use it on input streams iterators, like istream_iterator and istreambuf_iteratror.

The default behavior of accumulate is to sum, but we can also have a custom acumelator.
This will return the total number of characters in a container of strings.

The only interesting thing is the string::size_type type, which isn’t necessarily, but in all relevant cases, is identical to size_t.

For an averageing acumelator we should someohow remeber the number of all points, which means a functor object, and passing it some default value, and hopefully deal with the case of no fitting elements..
We can also use std::for_each, which is a bit more relaxed than the accumulate function in terms of the the std standard requires (for accumulate, we aren’t technically supposed to change the functor). However, for_each doesn’t convey the same direct meaning as accumulate in conveying our objective.

## Chapter 6: Functors, Functor Classes, Functions, etc

Function objects (functors) are a crucial part of the stl, and we need to know how to properly use them.

### Item 38: Design functor classes for pass-by-value.

Functors are designed to behave like function pointers, which means to be passed by value, ouch. The end user can try playing with this, but he usually won’t. The pass by value means that our objects are copied each time, so we need them to be small (so copying will be cheap), and contain no virtual functions (our old friend of slicing).
However, both expectations are unrealistic, functors are designed to hold state, so they are big as they need to be. And also, c++ uses inheritance, so it’s foolish to discard all the advantages of base classes, virtual functions etc…

There is still a way to make this work.
We make our functor object hold just a pointer to the real functor (the PIMPI idiom - pointer to implementation, the ‘bridge’ design pattern), we we pass the small object and have it call the real object, now we only need to have our function object support copying by value in the proper way, maybe with reference counting to the real function object, and we’re done.

### Item 39: Make predicates pure functions.

Predicate: a function that return boolean, or something convertible to boolean,
Pure function: a function that depends only on the arguments, without side effects or changes to the arguments.

The stl expects the predicates to be pure functions, and it works by copying by value.
This means that anytime we use a predicate, we use a copy of it, some the ‘state’ of the functor might be different.
To avoid the problems detailed in the chapter, we should make our () operator const, so we could never change our members from within the operator, but it’s not enough, because const operators can’t change the ‘this’ membersm but they can change static members, mutable members, whatever.

This is on us, not them.

### Item 40: Make functor classes adaptable.

We saw this before, we thes strcmp example, we want our functor to be ‘adaptable’,

This works:

This doesn’t:

(Not1, not2 are helper function that create a functor object from other function objects, the 1 and 2 mean unary and binary, we also have bind1st, bind2nd).

This is what we want:

We can’t call not1 directly on the pointer, we first must call ptr_fun on the pointer. What actually happens is that we simply create some typedefs the not1 helper function needs, this typedefs are what makes a function object ‘adaptable’, and eligible to be used in some sort of situations.
The typedefs are:
arguement_type, first_arguement_type,second_argument_type, result_type.
And also some other subsets of the names, and this becomes a hassle all of a sudden. The correct way to provide the typedef is to leave it to someone who already did this, a base class. Std::unary_function and std::binary_function are templates, so we can use them to create base classes and then inherit from them.

All sorts of things i don’t understand.

### Item 41: Understand the reasons for ptr_fun, mem_fun, and mem_fun_ref.

Those three annoying function are there to bridge between the three types of function calls into one common form, the three forms are:
f(x); normal function call.
x.f() - member function call from an object or a reference.
p->f() - member function call from pointer.

Now, if i have a function (like for_each) that takes function pointers, it needs to know what type of function to call, normal, member, pointer? The default behavior is to take normal function, soIt needs some way to bridge it.
This is where mem_fun and mem_fun_ref come in. the take member function of a pointer function and make a regular function out of it. They are also called function object adaptors.
(they also provide some typedefs that other stl functions need).

But what about ptr_fun?
It also adds the typedefs, we usually don’t need it, but sometimes we do.

### Item 42: Make sure less<T> means operator<.

less<T> is actually a template for a functor class we need to make sure it does what we think it does. There is no reason why someone other program couldn’t specialize the less<> template for our class in some place in the code and make it do something else than expected. For that reason shouldn’t override less for ‘sorting’ behavior, unless this truly is the natural sort (and in this case, we should write the < operator in the class definition), if we want a special case of a functor for comparison reasons, we should simply create it, and not write over the sacred less<T> template. Anywhere we pass less<T>, we can pass the new, specific and specialized functor.

Less should always call <. We shouldn’t change it, even if we are technically allowed to. This is how you get bugs.

## Chapter 7: Programming with the STL

This chapters is about how to use STL effectively. The Do’s and don’ts

### Item 43: Prefer algorithm calls to hand-written loops.

We should use algorithm instead of loops. Even though both are the same.
Algorithms are usually more efficient, more maintainable (easier to read) and are more correct (it’s harder to err with them).

The looping mechanisms in the stl for_each are probably more efficient, even more than inlining the function call.
Traversal in st loops is also usually specialized because we can make use of how the container stores element internally, and traverse better. Other algorithms also use better specific implementations that those a common c++ programmer can be expected to write for a simple loop.
In terms of correctness, using stl algorithms usually allows us to avoid common mistakes, like ordering when inserting elements, or invalidating iterators.

Another plus of using algorithms is that they have common names. Even if there are over 70 algorithms and over 100 templates in the STL, that’s still a finite amount with descriptive names.

Any programmer can see a algorithm from the stl and figure out the intention: sort, accumulate, remove, copy, search … and if he doesn’t know the algorithm, he can google the name.
With regular loops, we need to look at the code to understand the meaning, if we’re lucky, it’s a simple code, or well documented, but maybe it’s not. And then we need to closely inspect the code and understand what’s supposed to be happening, and decide if the bug we are looking for can originate from here, and then verify.

Ofcourse, sometime it’s easier to use a for loop than an algorithm, like this case, where we search for the 1st object that is between two points:

The first version is simple, loop over the vector, and for each element check if it’s larger than x and smaller than y. Once we found it, break.
Then we simply check to see if i is equal to .end(). (note that we define i outside of the loop, because we need to use it afterwards).
For the algorithm version, we use find_if, and then have a horrible compose functional that takes two functors (greater<int>, less<int>) and binds x and y to them.
This is a mess of a way to create a predicate, but that’s what the find_if requires. (maybe we could use a lambda here). We could pass the definition of the functor outside, but the number of lines would be incredible, and we won’t be able to use many template tricks, so it seems that for one-time use, the old-fashion loop really is better.

In a general note, code written with stl algorithm has less loops, which means more abstraction. So less for, while and do, and more for_each, assinn, insert and range functions.

### Item 44: Prefer member functions to algorithms with the same names.

If someone wrote a member function, it probably is optimized for the class (well, it should be), or at least, tested. This is why we should trust them and use the member functions when they are available.
Imagine searching algorithms. The member search (find) function knows how the container is used, so it can use all sorts of tricks to get better results, the search algorithm doesn’t know about the container, it only knows about the iterators, in this case, it can’t optimize the search behavior. Even if we use a binary search algorithm and pass the container, we need to know the implementation of the comparison function.
We also run into the issue of equality and equivalence. So in this case we might get different results if we use member function s and algorithms. This is especially prevalent with associative containers.

### Item 45: Distinguish among count, find, binary_search, lower_bound, upper_bound, and equal_range.

We want fast and simple. Always.
What algorithm we use is based on what kinds iterators we have, is the container sorted?
If it’s not sorted we have count and find, we can also use count as ‘containers’, because count returns zero or a number, which can convert to true and false, we can use find as contained by checking it against the end iterator.
For sorted ranges, we can use search algorithm with logarithmic times (now using equivalence, not equality).
We can use lower_bound as search algorithm that answers the question of ‘does this exists? If yes, give me, if not, where would it go’. The problem is the return value, as we need to use a equivalence check, not an equality check.
To overcome this problem we should use eqaul_range (that should be called equivalence range, actually), which returns a pair of iterators. And now we check that pairs’ first and second elements, if they are the same, then the element doesn’t exist. We can also use the distance between first and second as a count for the range between.

We can combine lower_bound and erase to remove all the elements that are lower than the criterion, or upper_bound and erase to remove all the elements that are lower and equal to the criterion.

Difference for associate containers (map,set, multi_map, multi_set). They usually offer member function for searching, (and usually for upper,lower, equal range).

### Item 46: Consider function objects instead of functions as algorithm parameters.

The downside of abstraction is that we trade generalization with speciality, the more general we are, the less efficient we usually are. However, it’s different in the case of function pointers vs function object. Surprise!

It has something to do with inlining. When we pass a pointer, we tell the compiler to use this pointer and invoke through it. A pointer to function is that different than any other pointer, and who’s to say that the pointer can’t change it’s value during runtime? Better to be safe and always invoke the function, even if it’s inline.

Function objects, on the other hand, a much more defined. The compiler knows about them, so it can inline them and optimize them. We suddenly get an abstraction bonus!
Also, some compilers still have bugs with function pointers,

### Item 47: Avoid producing write-only code.

Write only code is code that only you and god can read during the time of writing, and only god after a week.
This means sandwiching functions together, obscure code, etc…

This code, for example:

So, we got an v.erase wrapping something,
This something is a remove_if, but as the the first argument, not the common use - so we can deduce that we want to remove all the elements starting from some element.
Inside we have a reverse iterator search, note the .base(), .rbegin() and .rend(), and some binded function.
This code takes several minutes to read,all in all, the intentations was:
Here is the same code, but with non inductive names and without indentation:

Isn't that lovely?
Even with the correct names, we have so much to unpack: erase_remove idiom.
Using bind2nd, reverse iterators and base(), and nested find within a remove.
Even a good comment won’t save this code. It will help. But that’s about it.

The problem is that the code is the natural result of how we phrase the problem,
We were asked to remove elements, so we know we need an erase remove combo. And we want to look at elements after the last elements the fulfils some condition, so that leads to reverse iterator use.. And then we end up with the same code as before. ‘Easy’ to write, hard to read.
The ‘better’ readable version of this code will still be similar, but it’ll probably be divided a bit and use named variables (worse performance), and we will still need a big comment to explain what’s going on and how to work with it (and a warning not no separate the lines),but it will be more manageable.

### Item 48: Always #include the proper headers.

Don’t rely on header files having the right include statements, even if they do in one implementation.We need to have the correct #include statements, in each file that we use them,
That’s life.
A refresher:
Containers: each containers is declared in an #include file with the same name. The exceptions are <map> and <set>, which include the regular and multi version(map, multi_map, set, multi_set). Also don’t forget unordered versions.
Algorithms go in <algorithm>, except the accumulate, inner_product, adjact_difference and partial_sm, which come from <numeric>.
Any iterator other than the base one are in the <iterator>, this includes istream_iterator and istrembuf_iterator.
Standard function wrappers come from #include <functional>, this means less<T>, not1, bind.. Etc.
If we forget one of them,we might be fucked when we try to port them, or when someone changes their STL implementation.

### Item 49: Learn to decipher STL-related compiler diagnostics.

STL can create cryptic compiler errors.
Each implementation can be different, this is partly because of templates, which can have a frighteningly large real name in diagnostic messages.
No magic solutions, learn how to read through and focus on what lines might really mean something. You’ll eventually develop an eye for it.

### Item 50: Familiarize yourself with STL-related web sites.

Sgi stl - was retired.
Stlport - still there
Boost - stiil there.

Have some resourecs about STL, implemenation examples and some experimental containers and algorithms. Debuggers for portings, nice libraries, etc...

## Extras

### Appendix A: Locales and Case-InsensitiveString Comparisons

Locales are horrible.

### Appendix B: Remarks on Microsoft’s STL Platforms

Something about early versions of visual studio (6 and below).
