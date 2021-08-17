# More Effective C++: 35 New Ways to Improve Your Programs and Designs

## Introduction

If in question, name pointers (int main) with “p” as first letter and references with “r”, also remember lhs and rhs (left-hand and right-hand sides) as common naming conventions.
We can leak stuff other than memory, file handlers, mutexes, semaphores..

## Basics

### Item 1: Distinguish between pointers and references.

Pointers and references are similar and different.
No such thing as a null reference, a reference must be initialized when declared.
(except in this case…):

Because reference can’t be null, we don’t need to check their existence before using, so we have one less if statement (ptr == NULL).
Pointers can be reassigned to point to other objects (or no object at all), references always refer to the object they were initialized with. So if we know we always have something to refer to, and we know we don’t want to refer to anything else, then that’s a situation to use references.
Another situation is when implementing some features (mostly operators), such as []. We use [] to return an object that we want to change/modify (x[2]=6), so we return a reference.

### Item 2: Prefer C++-style casts.

Sometimes casts are necessary. But it doesn't mean we can’t use them correctly and responsibly. The C-style cast is far too crude and wide. It can do everything, but we want some more specialized forms,.
Also, because of their simple form: “int y = (int)x’;” they can be difficult to find and identify, even with tools like grep,
C++ casts are broken down to specialized forms, with indictive names:
static_cast<T>(t), const_cast<T>(t), dynamic_cast<T>(t) and reinterpret_cast<T>(t).
Static_cast is the general purpose cast, it can cast from different types according to basic casting rules. It can’t cast between classes and primitives and between types that don’t have a casting operation. It also can’t cast away constness (we have const_cast<T> for this).
Const_cast is used to remove constiness or volatileness of an object, that’s all it does. It can’t change the type.

dynamic_cast<T> is used for safe casts down/across inheritance hierarchy. We use it to cast pointers or references of base class objects into pointers/references of derived classes.
It also allows us to determine if the cast was successful, as it returns a null pointer if we can’t downcast a pointer, or an exception if we can’t downcast a reference. It can’t change the constiness, it won’t work on objects without virtual functions, and it’s silly to try to cast between classes that aren’t in the same inheritance chain.

The reinterpret_cast<T> is mostly platform dependent, it’s mostly not portable and is very dangerous. We mostly use it to cast between pointer types.

### Item 3: Never treat arrays polymorphically.

When we have OOP and inheritance, we can do stuff polymorphically, treat any derived class as the base class and stuff will be fine.
The problem is with arrays. They don’t play nice with polymorphism.
Because of the way arrays are implemented (continues blocks of memory), and the way pointer arithmetic works, passing a derived class array as a base class array is a recipe for trouble. The size is determined by the pointer, so we get very weird behavior and we try to accesses memory as if it was something else (this is similar to a horrible use or reinterpret_cast<T>).
This also happens if we try to delete an array of derived objects through a base class pointer,

### Item 4: Avoid gratuitous default constructors.

Some classes don’t have reasonable default (empty) values, this means they shouldn’t have a default constructor. however:
In classes without a default (no parameters) constructors, there are three problematic contexts:

Creation of arrays
When we create an array, we initialize the objects of the array, if there’s no default ctor, we can’t initialize them! If we’re working on the stack, we have a workaround, as we can provide argument in the point of defining the array:

But this won’t work on heap arrays. We can overcome this by making an array of pointers (which don’t need a ctor), and then initializing each of them separately.

This is problematic because we need to individually delete each item, and we have an extra dereference and we need extra space for the pointers.
We can have a overloaded ‘placement’ new operator that does this for us (allocates memory as needed, and then initialize), but we would need a matching ‘delete’ operator and in general this is a hard thing to do and many things can go wrong.

Use with template based containers
Many Template based container classes require us to have a default ctor, this wasn’t nescarly by design (some templates don’t require it), but it’s enough to be a problem.
Virtual base classes
If we have a virtual base class (multiple inheritance) with no default constructor, any derived class (no matter how far removed!) must explicitly call it. This breaks the idea of inheritance.

Some people argue that therefore we should provide a default ctor even when it doesn’t make sense how to. This can lead to problems, as now any function needs to verify that the object calling it isn’t null default object, which makes our program much more complicated and causes performance issues.

## Operators

<details>
<summary>
When and how to overload operators, and more importantly, when not.
</summary>

### Item 5: Be wary of user-defined conversion functions.

C++ silently converts not only in promotions (char to short, short to int, int to float, float to double…, int to long). But can also do narrowing without any problems, even from doubles to char.
This is the problem with single argument constructors and implicit type conversion operators.
Even if we define c conversion operator from our class to another, it doesn’t only apply explicit casting, it also means the compiler can use our casting operator to pass our class to the incorrect functions!

We can avoid this by removing casting operators and going we defined functions with the same functionality,
Class a
{
Double operator();//casting operator.
Double asDouble();//actual function.
};
We might need to call more functions this way, but we make sure we aren’t calling our function by some mistake. This is why std::string has a an explicit call to c_str() and not a casting operator.
The other side of the problem is single argument constructors. They also act as implicit casting operations. So all of the sudden of code becomes less secure and the compiler doesn’t protect us as much as before.
We avoid this problem by simply making our constructors explicit. That is, we instruct the compiler not to call them as casting operators. If we would still want to use them as casting operators, we should simply call a static_cast<T> function.

Older compilers don’t have the explicit keyword, so there are some other ways to avoid the problem. We can take advantage of the fact the compilers allows only one User defined conversion to occur in a call. So rather than take a base type as an argument, we take a nested class that itself has a single argument ctor.

In a constructor statement, we pass an int, which is silently converted to a arraySize object, which our class accepts.
However, if we have a function that takes our class, if we pass an int to the function call, the compiler can at most turn it into an object of ArraySize. This is the most it’s willing to do.

This idea is called Proxy classes. We will see them later

### Item 6: Distinguish between prefix and postfix forms of increment and decrement operators.

X++,++X have different behaviors, but are the same operator. This distinction is done via operator overloading and a silent int argument.

The two forms also differ with their return types, the prefix returns a reference, while the postix returns a const object.
This const protects us from behavior such as i++++; we don’t want to able to call it.
The posifix version requires us to create a new object to return, this is bad for efficiency. We should use the prefix version unless we specifically need the postfix version.
Another point to consider is that both operators should do the same, so we have duplicate code.
Wait a few minutes, we don’t have duplicate code, someone updated the class, and only updated one version of the operator.
To avoid this situation, we should implement the postfix operator in terms of the prefix.

Now we have less things to worry about. This is better.

### Item 7: Never overload &&, ||, or ,.

&& and ||
&& and || are evaluated in order (short-circuit), so if we have a definite answer, we don’t continue evaluating. This is both for efficiency reasons and allows us to assume that if we reach the 2nd condition, the first statement must have finished (if we know it’s not null, we can de-reference it).

However, if we overload them (which we can), then we aren’t dealing with a short_circuit semantics anymore, but with a regular function
In the past it was like this:
If (expressio1 && expression 2)
But now it's this:
If (expression1.operator&&(expression2))
Or this:
if(operator&&(expression1, expression2))

And when a function call is made, all parameters must be evaluated, and the order is not determined.
This means we lose both efficiency (as both expressions have to be evaluated), and security (the 2nd expression will be evaluated no matter the result of the first one), and we simply messed up our code. This is bad.
The comma operator
We usually don’t consider the comma as an operator, but it is.
We see it mostly in loop statements with more than one index.

The comma operator also has some special rules for behavior (different than the comma in a function call). The comma operator has a left associativity, and it functions as a ‘set/break/sequence evaluation point).
All side effects happen when the comma is encountered, and before going to what’s in the right side of it. That what an sequence point means. However, if we overload it we are making it a normal function (whether we overload it as member function or not),we can no longer guarantee the expected behavior, as the compiler is free to choose what and how to evaluate. So we shouldn’t overload it.
A quick refresh: we can overload these

And these we can’t

We can’t overload the global new and global delete, but we can overload them as operators.
We should only overload what makes sense to overload, not everything under the sun.

### Item 8: Understand the different meanings of new and delete.

The difference between the ‘new operator’ and ‘operator new’. Yes. those are two different things.
The ‘new operator’ is the regular one it’s built in into the language like sizeof(), it’s global, and we can’t overload it. It allocates memory and calls a constructor, this is what it does, and we can’t change it.

We can change how the memory is allocated. And this is done by ‘operator new’ (GRRR!)
So ‘new operator’ calls ‘operator new’ isn’t that simply simple.
The ‘operator new’ is (usually) declared as such:
Void \* operator new(size_t size). This is something we can overload, and have more arguments, the first parameter must be size_t. All that ‘operator new’ does is to allocate memory, the ‘new operator’ takes the memory and calls constructors on it to make it into an object.
Note: the compiler disallows us to call a constructor on an existing object, but the compiler can, the best we can do is call ‘placement new’.

Placement new
#include <new>
Placement new is a version of ‘operator new’. It can take memory and construct an object on it. It is sometimes used in shared memory or IO.
Placement new is kinda shady looking:

It takes a size t parameter (which is unused! But also unnames, so the compiler doesn’t complain), as all ‘operator new’ s do, and a location that it returns.
Now that the ‘operator new’ has done it’s deal, the ‘new operator’ can call the constructors and finish the job.
Creating an object on the heap: ‘new operator’.
Customizing memory allocation: ‘operator new’. And then the new operator will use it.
Creating an object on existing memory- placement new.

</details>
## Chapter Deletion and memory deallocation
Each dynamic allocation must be matched by an opposite allocation.
If we deal with raw / uninitialized memory we should bypass the ‘new operator’ and ‘delete operator’ and use ‘operator new’ and ‘operator delete’ to handle it.

If we used placement new. We must have placement delete. The new operator and delete operator assume heap memory, but there is no such guarantee from placement overloads. We need to know where the memory came from and release it from there, while our placement ‘delete’ should only call destructors.

Arrays
Array new (new[]) is also calling the ‘new operator’, but rather than regularly allocating memory (with operator new), it’s using the operator new[] version (which we can overload as we wish). Isn’t this all lovely?
The new operator behaves differently for arrays as it calls the constructor many times: once for each object (and the destructor does the same).

My conclusion:
New operator calls operator new and the constructors.
Placement new calls operator new that simply returns the memory, on which the constructors are called.
New operator can also call operator[], which is modified to return the correct amount of memory (and some extra bits), and then calls the constructors in order.

New operator always calls calls operator new and constructors (or operator new[], or placement new…).

Exceptions
Once exceptions entered the stage, it became harder to write code, there are more opportunities for dangling pointers and resource leak, code needs to be written differently to handle exceptions. Exception safe code is a product of design, not by chance.
We could try to ignore exceptions and stick with error codes and have setjmp and longjmp, but they are probably more trouble than they are worth.

### Item 9: Use destructors to prevent resource leaks.

We shouldn’t use pointers to manipulate local resources.
We can’t have ‘delete’ (or new) in anything that isn’t managed. If a delete appears in the code and an exception is thrown somewhere, then we won’t call that and we will leak resources (remember, resources can also be files, not just memory!).
We move our resources to live inside a local class, something that always has an destructor called on it. This is RAII, and more exactly, smart pointers. the smart_ptr (either auto_ptr in the past, or unique_ptr/shared_ptr today) knows to take the control of the resources when it’s created, and relinquish it when it leaves the scope (and has the destructor called on it). We use those smart pointers instead of raw pointes.
(not in the case of arrays in auto_ptr).
This way of thinking (RAII) is suitable to anything that needs to be released, not just pointers, if we need a resource, we create a class that acquires it at initialization, and releases it when destructors are called. The rest is bookkeeping.
The only case to be worried is if an exception is thrown while constructing the object!.

### Item 10: Prevent resource leaks in constructors.

Constructors are called only after an object was fully created (fully). If it hasn’t the destructor won’t be called. This means that if an exception is thrown, a resource might not be released.

C++ doesn't call destructors on half created objects.
This means we must design our constructors to deal with this situation catch any exception, clean up, and rethrowing the exception (because it might inside a different object itself).

This brings us to a clean up function that is called borth inside the constructor and the destructor,
But if our members are const, we need to initialize them in the MIL, which makes use the ternary operator (?:) and then we function inside the MIL that does everything (including catching exceptions).
Another option is to treat each member as a managed object, so when the stack unwinds they will have the destructor called on them like any other local object. Use unique_ptr and sharted_ptr (the text says auto_ptr, but NO).

### Item 11: Prevent exceptions from leaving destructors.

Destructors can be called in two situations.
If the object is normally destroyed when it goes out of scope or has ‘delete’ called on it.
The second case is when an exception occurs and a stack unwinding occurs,
This means that an exception might be active during a call to the destructor, and if we have two exceptions at the same time, the terminate function is called, and then not even local objects get destroyed.
We usually don’t want this.
One options is to put an empty catch(...) block in our destructor, so any exceptions are caught. This means that only one object fails to properly be destroyed, but any objects that use it are destroyed properly.

### Item 12: Understand how throwing an exception differs from passing a parameter or calling a virtual function.

Exceptions aren’t the same as normal variables. The throw isn’t the same as calling a function, etc.

We can pass function parameters and exceptions by value, by reference and by pointer, but what happens inside isn’t the same.
When we call a function, the calling convention dictates that the caller does the work, and that the control will eventually return to the code that called this function.
This isn’t true with throw statements, the code will not return control to the throw site.

Consider this case:

We must make a copy of localWidget somewhere so it can be copied and caught, otherwise the stack will unwind and we will lose all the data in it. C++ specifies that any object thrown as exception is copied, this is true even if there is no danger of it being destroyed.
This means that a thrown object is copied, and cannot be modified, and is also slower than returning a parameter.
The type that’s being copied is the static type, not the dynamic type.
The copy distinction also means that the following two variations of code are different.

The first catches a copy, and then throws the original.
The second catches a copy, and then throws a new exception, which means another copy is created (copy of copy!). If the exception was of some derived class, the new exception is now of the Widget class, and information can be lost.
This is why we should use ‘throw’ without any arguments. So we won’t change the type. And we don’t need an additional copying costs.

An exception thrown is always a temporary object,(created with the copy constructor), and it usually isn’t allowed to catch a temporary object with a non-const reference. But exceptionsa are… the exception.

In all cases above we make a copy of the exception (with the copy constructor), but in the first case we make two copies (temporary objects, and then into w), so catching by value seems silly.
Catching pointer is simply returning a pointer object, it follows the same rules as returning by value, but the object is a pointer, we must not return a pointer to a local object. As destructors are called when the stack unwinds.

Another difference between functions and throwing exceptions is the way the value is return up the stack or propengated.
In functions implicit conversations (promotions) our allowed. Not in matching exceptions and catch blocks.
double x = sqrt(55); //implicit promotion from 55 to double.
Catch (double d){} // no implicit promotion!
Catch (int i){}

This catch clause wont work!

It’s designed to catch double value, and we throw int. Type conversions are restricted, inherice based conversions can occur, and untyped pointers conversions.
A catch block for a base class will catch and publicly inherited derived class, just like java.

This logic expends to pointers to these classes,and also applies to values, although we shouldn’t throw pointers and we should always catch by reference.

If we have a catch statement for a void pointer (void _), then it will catch any pointer exception (actually const void _ will be better).

A final difference is that catch blocks are evaluated in order of appearance (in cases where more than one applies), this is different from determining functions with function overloading, where there is some race and score that determines which functions is called. Essentially, this is ‘best fit’ in functions vs ‘first fit’ in exceptions.

### Item 13: Catch exceptions by reference.

Just like before, we can catch by value, by pointer and by reference.
We should always catch by reference.

In theory, catching by pointer should be the most efficient way, as copying pointers shouldn’t involve copying the entire object. It seems right, but it works best with static/global objects, otherwise we must ensure that our object exists when the stack unwinds, and now we are back with memory management which we hate.
Even if we create an exception with ‘new’, this is even worse, as now the user in the catching block must call delete on our objects, and he has no idea how the memory was allocated.

And passing by pointers is also against the language convention, the four standard exceptions
Bad_alloc - when ‘operator new’ fails.
Bad_cast - when dynamic_cast<T>() fails
Bad_typeid - when we try to dereference a null pointer in some typeid way
Bad_exception - something else.

Are all objects, not pointers to objects, so we match catch them as such (by value or by reference).
Catch by value is nice, but costs us two copies of the object, and we risk slicing problems if we use the copy constructor of a base class on derived object,
Catch by reference is safe for the standard exceptions, doesn’t cause unnecessary copying and doesn’t risk slicing. This is what we should always use (together with throw, and not throw()). We should also catch a const reference, just to be doubly safe.

### Item 14: Use exception specifications judiciously.

(is this even true today?)
Avoid this.
Never trust anything.
Even templates.

All sorts of ways to try.

### Item 15: Understand the costs of exception handling.

Exceptions at runtime are costly. We pay the price even if we never use them. The c++ standard made the choice to support exceptions, and even if try to compile without exception support (some compilers allow this), people who use our code might still need exceptions.

Try-catch blocks have an extra cost, even if they catch block is never used, all this cost is both in code size, file size and performance speed.

The worst cost is when an exception is actually thrown. They shouldn’t be thrown, and our program should have alternatives to deal with most (if not all cases) that can lead to an exception and we want to avoid this. The cost of throwing an exception vs a normal function return is about 3 orders of magnitude larger.

This is of course situation dependent, and exceptions have their benefits, but if we can write code that minimizes the use of exceptions, we should do so.

Efficiency
Efficiency is good, and important, in all software.but c++ really cares about it.

### Item 16: Remember the 80-20 rule.

As all things in life, 20% of the code is responsible to 80% of the (cost, work,problem, performance issues… whatever).
There is no magic way to know which part it is other than testing. And the profiler can also impose issues that are specific to it,

### Item 17: Consider using lazy evaluation.

Wait with computations until you’re sure they are needed. If something isn’t required, we would rather not compute it at all.
Reference counting
This way of thinking is very prominent with copying, as many times we can share a single copy, until we need to modify one of those copies, at which situation we don’t have any choice and we must do real work.
Distinguishing reads from writes
The [] operator can be ‘read from’ or ‘write to’. If we only read, we don’t need to change anything.
We use proxy classes in this case to deter the choice and the initialization to the last moment.

Lazy fetching
When we have objects in remote databases, we might not need the whole object right away, and maybe we simply need some small amount of data from it.
So we create a ‘shell’ object that holds a connection, but does the actual fetching only when a field is requested, we combine mutable data members for this. Before accessing data we check if we already had accessed it, and if not, we initiate a request.

Lazy expression evaluation
Do we really need everything we compute? Not really, let’s wait.
APL -- ???

## Summary

### Item 18: Amortize the cost of expected computations.

A different approach from lazy evaluation. In this case, we are over-eager to to do stuff, and we perform calculation before we are even asked.
If we know an operation is going to be requested very frequently, we can calculate it right away (or when it changes), and then return that value.
(memoization, mapping, cache, look up tables, they all depend on this way of thinking).

Goes together with localizability of reference. This is similar to dynamic vector (STL::vector<T>) that does automatic resizing of itself.
We reduce our computational time by paying with space.

### Item 19: Understand the origin of temporary objects.

Unnamed variables that don’t appear in the source code, created either with implicit type conversions when calling a function or when a function returns an objects.
There are creation and destruction costs.
Implicit conversions occur when we pass by value or when we pass a reference to a const.
They don’t happen with reference to non consts.

Operator+ returns a const object.
operator+= doesn’t. Isn’t this nice? Sometimes we can abuse this.
(in our head x = x+y is the same as x+=y. But for cpp this is different).

We can sometimes skip the costs of temporary objects, by careful planning and using the correct functions

### Item 20: Facilitate the return value optimization.

We don’t like return by value because of construction costs.
Sometimes we can’t avoid this, sometimes we can.
Avoid returning pointers (can’t be from heap, if on heap, we need to delete them).
Sometimes we can return references, but sometimes we can’t.

What we can do is cheat the system by creating the new object as a temporary object (create it in the return statement),

This can allow compilers to do optimization magic, like eliding constructors and such.

### Item 21: Overload to avoid implicit type conversions.

For functions/operators that create new objects, we would like to avoid implicit conversations.
If we have a type X that can do + with an int, we would want this int to be an int, not have a X(int) constructor be called just for this case.
We write extra functions (and avoid implicit casting)

If we write functions that know to deal with int directly, we avoid calling constructors.
This won’t work on the case of X = int +int, because we can’t override the default int + operator, but we still save some costs.
As always, this should be weighted against maintainability constraints, if we don’t have a bottleneck in this section, it’s not much use to spend time adding more and more functions.

### Item 22: Consider using op= instead of stand-alone op.

x+=y and x =x+y. Are in no way necessarily connected, as far as cpp is concerned.
We saw this before when talking about return values.
In many cases we should implement the op+ in terms of op+=, and so on. This saves us some object creation and improves maintenance, and allows decoupling and non friend functions.

This uses template functions, using the += operator rather than +, and because we aren’t naming our object, it should be a better candidate for return value optimization.
It just looks so ugly. But that’s performance.

### Item 23: Consider alternative libraries.

“Library design is an exercise in compromise. The ideal library is small, fast, powerful, flexible, extensible, intuitive, universally available, well supported, free of use restrictions, and bug-free. It is also nonexistent.”
Choosing a library is a trade off. In the case of stdio vs iostream (<stdio.h>,<iostream>), stdio is faster, while iostream is type safe and extensible. The choice is based on the requirements. If one section of the code is causing serious performance issues, maybe a different library can do better.

### Item 24: Understand the costs of virtual functions, multiple inheritance, virtual base classes, and RTTI.

Virtual functions, dynamic typing, vtable.
Where is the vtable code loaded? Some compilers use weak symbols and allow the linker to deal with them, and some compilers choose a different strategy.

We should avoid declaring virtual functions as inline, because of some reasons (most compilers ignore the inline completely for virtual).

The vtble pointer sits somewhere in the object. it costs us an extra block of memory per object, and extra work for each constructor. even if memory isn’t a problem, this can reduce caching and increase paging.

Calling a virtual function requires:
Get the the pointer to the virtual table. This is done by offset.
Get the corresponding function in the table, this is fast.
Call the function through the pointer, one dereference.

We get extra costs because we can’t effectively inline virtual calls. As they aren’t resolved until runtime.(there is a difference between invoking virtual functions by objects against by pointers or by reference).

This is relatively simple with single inheritance, but multiple inheritance is awful. They might have more than one vptr table, and we might need some complicated offset calculations to determine the correct vtable to call. And then the virtual base class comes and complicates things even further with the diamond inheritance shit messing everything up.

Runtime type identification - RTTI
How can we tell what the type of an object is in run time?
With it’s type_info object (for classes) which we get by calling the type_id operator!
(this is guaranteed to work if we have at least one virtual function), in some cases the type_id is even implemented as a virtual function

Techniques
The good ol’ bag of tricks.
Sometimes also called techniques, idiom, and if we are feeling really fancy, patterns

### Item 25: Virtualizing constructors and non-member functions.

In general, no such thing as a virtual constructor, but we still have them because they are very useful. A constructor method that can create different concrete types based on the input and the request.
Also the virtual copy constructor, which is sometimes called copyself, cloneself, or just clone.

The virtual version simply calls the real version, this is possible now that a derived class virtual function can return types that aren’t the same as the base class (as long as they are derived from that class).
In this example:

Actual code:
“
newsLetter::newsLetter(const newsLetter &rhs)
{
For (list<NLComponent *>::const_iterator it = rhs.compnenets.begin(); it!=rhs.compenents.end();++it)
{
components.push_back((*it)->clone());
}
}”
This is a copy constructor.
We loop over each elements in other newsletterm and then we call the clone method, the clone method knows which copy constructor to call, and we don’t need to care about they derived types and which constructor to call.

Making non-member functions act virtual
As we mentioned before, constructors can’t really be virtual, but we it’s useful to have them as virtual, so we create virtual functions that call the real ones.
It can also be useful to have non member functions that depend on the type of the class, however, this isn’t easy:
The <<operator isn’t a member function, it takes the ostream& as left hand side argument, so if we overload it (and make it virtual) then stuff gets weird and unintuitive for the user.
Like before, we solve our problems by having a function call another, this time, a concrete function calls this the virtual one.

Print is a virtual member function, but it’s called by the regular <<operator function!

### Item 26: Limiting the number of objects of a class.

We sometimes want to limit the number of objects in a class, maybe we have a limited amount of connections to a network, maybe our buffer is limited, etc…

Allowing zero or one objects
If we don’t want our constructor to be called, we simply declare it private!.
Let’s say we have a printer class, and only one printer in the office. It would make no sense to have several printer objects connected to one real printer, and we would like it to be shared.
Now, we stick this printer inside a function (static variable inside a static function), and if someone wants to use the printer, he simply calls the static function, and gets the one printer.
The reason to use a static variable inside a function and not as a static class member is because static variables inside functions are initialized when we first call them, so if never call the function, we never pay the cost. Class members are always initialized (but we don’t always know when exactly, which can lead to bugs).

Note: we should NOT make this function an inline function. That defeats the purpose.

We can have a counter in the constructor that throws an exception, which is simpler to read and less confusing, but this also has problems, if we have a derived classes then they all share the base class behavior, and if we have class that contain members of our limited class.this can lead to many bugs, but private constructors protect us. The object stays the same, and once it’s created it doesn’t have any problems, but the creation of this object is limited.
This also helps us with limit derivations,

Allowing objects to come and go
If we limit our access to an object, we need to know if we wish the destroy it and re-create it.
(maybe it has some shared resources that are needed for something else).
We can get this result by combining the counting strategy with the pseudo constructors. We also have a wrapper around the destructor, that decrements the count.

An object counting base class
If we have many classes that need to be limited, writing the same code again and again becomes tedious. For this we have templated classes. They act as base classes, and they define the count, the constructors and the destructors. We use private inheritance to signal that this inheritance is for implementation, not for ‘is-a’ relationship.
If we do want to expose members of the base class, we can simply employ a ‘using’ declaration.

If we don’t want to hardwire the number of objects, we simply leave it for our user.

### Item 27: Requiring or prohibiting heap-based objects.

Sometimes (Mostly in embedded systems) we want to make sure that stuff can’t be on the heap, or maybe we want it to be only on the heap.

Requiring heap-based objects
If we don’t want our users to be able to create something, we can simply limit access to the constructor by making it private.we can also make the destructor private (which prevents the compiler from calling the dtor, and prevents the operator delete from calling the dtor).now the user must use new and a custom destruction function.

If we want our object to be viable for inheritance and composition, we can either make it’s destructor protected or require classes to hold pointers to the object instead.

Determining whether an object is on the heap
All kinds of games that can fail.
Also making the program hardcoded and trying to check the addresses to make sure they are on the heap can fail (static variable aren’t on the heap or the stack!).

This is horrible, and we usually have this only because we want to be able call ‘delete this’ at some point.

One way is to have a ‘global’ table of allocated memory addresses, and then check against it before calling delete, it’s frightening idea, which might not always work because other programs might want to control the operator new and a operator delete, and we can’t know whether we are the base class or derived,

….
[CONTINUE THIS]

Prohibiting Heap-Based Objects
We might want to disallow creation of objects on the heap. Again, three cases to consider: direct instantiation, instatination as a base class and instantiation inside other objects.

To stop direct instantiation, we simply need to stop the ability of the new operator, which we can’t because the new operator is global. But it calls the ‘operator new’ of the class, which we can control. So if we make it private (along with the delete call, for good measure), then the user can’t call new and directly create classes. We should also do the same with array new (operator new[]) to stop it.
As a bonus fact,having operator new private often prevents other objects with this class from being instantiated. This isn’t some magic. Simply that operator new and operator delete are inherited, so if the derived class doesn’t redefine them, it won’t be able to call new and create itself. This is not a real defence measure, just a nice bonus that we get for free.

In the end. We can’t do it. What a let down.

### Item 28: Smart pointers.

(this is long! Definitely important!)
Smart pointers are object that are designed to look and feel and behave like raw pointers, but also have greater functionality like resource management, and repetitive coding task automation.
When we use smart pointers, we gain control over:
Construction and destruction - what happens when the pointer is created and when it goes out of scope.
Copying and assignment - what happens when we assign or copy the pointer, do we have deep or shallow copy? We choose.
Dereferencing - we can use this for lazy behavior. Yeah!.

Smart pointer are generated from templates, and they must be strongly typed.
They usually have a ctor that takes a real pointer as an argument, a copy constructor, an assignment operator overload (the = operator), destructor, and overloading of the dereference operator (\*) and the arrow operator (->). The real pointer is stored as a private member (PIMPI idiom - pointer to implementation).

The copy ctor and the assignment operator a public and have const arguments because they don’t modify the original (unless we use unique ptrs and move semantics).

An example of using smart pointers:
An application that handles both local objects and remote objects (stored elsewhere, either in cloud or on the disk or some other process, i guess). We want the application to be able to handle both types of objects the same, without knowing which type of object is exactly used.
In the example in the book we have a smart pointer class that can take either a local pointer or a database id as constructor arguments, and a class that logs changes.

Smart pointers are supposed to be simple to use. We should be able to treat them as regular pointers.

Construction, Assignment, and Destruction of Smart Pointers

Construction is usually straightforward, get an object you want to point to, and store the address as pointer.
The copy constructor, assignment operator and the destructors are a bit more complicated because of ‘ownership’ issues. If a smart pointer owns the object, then it should be in control of destroying it. But in theses cases, we can’t be sure that the current one is really the owner of the object. This is problem with the (now deprecated) auto_ptr class. It worked only when it was the only auto_ptr pointing to that object,but when there were two or more, it caused problems, there were some ways to circumvent this, but in the end the chosen solution was to transfer the ownership of the pointer. This ensured that only one auto_ptr could point to the object at all times.
This solved some problems, but caused others: we should note that the copy constructor and the assignment operator no longer took ‘const’ objects, and they actually changed the arguments passed to them. An unsuspecting programmer could use the auto_ptr and lose the ownership along the way and he would have no way to know this until either checking or crushing the program.this also meant that passing the auto_ptr by value (i.e, calling the copy constructor) would usually yield bad results, as the ownership was lost during the function call (we could ask for it back, but that also seems troubling). We could still pass by reference, which doesn’t involve the copy constructor.

The destructor is actually also not that complicated, it simply should check if it owns the object, and if it doesn, he should delete it. An auto_ptr always owns something (or is null), so calling delete is always the correct behavior (delete on null is fine).

Implementing the Dereferencing Operators
Smart pointer are so easy to use because they mimic the behavior of raw pointers directly, we want them to be transparent as possible, so dereferencing a smart pointer should be exactly as dereferencing a pointer.
The following is a simple implementation:

In this example there isn’t any lazy evaluation or remote fetching, so we directly dereference the raw pointer member and return a reference (not the actual object),this reference protects us from calling the copy constructor and from the possibility of slicing and allows us to keep using virtual functions.
The above implementation doesn’t defend against dereferencing a null pointer. We get undefined behavior. Just like real pointers! Isn’t this nice? We can protect the user against this behavior, but that’s entirely up to us.

The arrow operator (->) is similar.
Here is an example:

We can have some other behaviors. Like the smart pointer class of reference counting (std::shared_ptr?), but this behavior is the basic idea.
Testing Smart Pointers for Nullness
One problem is that we want the user to be able to know if his pointer is valid. In raw pointers we do this by checking against null. But in the case of smart pointers, we aren’t so lucky.

Ofcourse, we can provide an ‘isnull’ member function, but then our smart pointers won’t behave like a regular pointer anymore. What we should do is have an conversion function to void*. This conversion should return zero if our pointer is null, and non zero if not (not necessarily the address, we still want to hide our pointer).
This allows checking for null,but can lead into some other problems, because implicit casting means that any place that takes a void* can take our class as well.
So the following code will compile:

And based on how we defined the conversion operator, might even be true.
A different approach is to convert to const void \* or bool, to reduce the chance of running into implicit conversion cases, the problems are still there, but most code doesn’t use them as much. A different approach is to overload the not operator (operator! ()const). This eliminates the problem of implicit conversions, and provides a decent way to check for nullnes.
The operator! Will return true only if the pointer is null.
So checking for null is now:
“if (!ptn){...}
else {...}”

Which is similar a calling
“If (ptr != null){...}
Else{..}”
On regular pointers. Both are allowed.
We can still find weird nooks for programmers to get smart with us

But nobody should write like that.

Converting Smart Pointers to Dumb Pointers
We might have existing legacy code that works with regular pointers, and we want to add smart pointers to the mix without rewriting our entire code.
We can create conversion function to our templated classes that returns the pointers when asked. This actually solves our earlier problem of testing against null. Our smart pointer just hands out the private member whenever someone asks for it. Isn’t it great?
This solves some of the issue, but probably complicates others. What good is our smart pointer if we allow someone else to mess around with it’s inner parts? This goes double for reference counting smart pointers.
Another problem is that this kind of conversion is a user defined conversion, and the compiler allows just one of those in a conversion chain. And this also opens up a case for a compiler bug.

This shouldn’t have compiled at all. But it does. We shouldn’t be able to call delete on non pointer objects, but pt can be converted to such, so the compiler allows this. Now we will have double deletion (onec at this delete, and once when the pt object goes out of scope). We should provide conversion to pointer only when we must.

Smart Pointers and Inheritance-Based Type Conversions
Smart pointer types don’t know about inheritance.
A smart pointer of a derived class isn’t a derived class from the smart pointer class for the base class.
(even if B is a derived class from Base class A, smart_ptr<B> isn’t derived from smart_ptr<A>)
They are different, unrelated classes. Nothing connects them. So although a function with a signature of a base class pointer can accept derived class pointers, a similar function with smart pointer argument can’t.

There is a way to take care of this, we give each smart pointer a conversion operator to the base class.
Only that this is horrible. This means we need to manually define each smart pointer class (and say goodbye to the template!), we also need to define a conversion operator for each of the base classes, because the compiler won’t do more than one user defined conversion for each conversion chain, so if my class is twice derived, i need a conversion function for both classes. This is horrible.

Luckily for us, if we ask nice enough, the compiler will generate this for us, if we ask the right way:

This is again our smart pointer class template, but we have another templated function inside it. Isn’t it nice?
This template function is the overload of conversion operator from our class to another class. It simply calls the constructor of the other smart_ptr class with our pointer.
We get type safety because smart_ptr<Base>(derived \*ptr) is legal code.
The compiler will only generate conversion functions that are legal and are requested from him.

This method works not only for inheritance, it works for any legal conversation.
One last problem:
If we have two function with the same name, and our smart pointer can be converted to either of them (in the example, the direct base class, and the base class above it), we get ambiguity. Unlike regular inheritance, there isn’t ‘more likely’ conversion that says we should try the immediate base type smart pointer before going further up. All conversion functions are equally likely for us.
Two more problems: in time of writing the book, the support fort templated member functions is still lacking in some compilers (probably not today, though).and this code is hardly readable for a novice programmer, and being too clever can be dangerously.
But about our problem? We can’t solve it yet. Smart pointers are smart, but not pointers. If we get to this situation, we should probably use a cast and hope that the cast isn’t neglected if the inheritance chain is changed in some matter that creates a better suited function match.
Smart Pointers and const
Pointers can be both const themselves and point to a const object. Smart pointers aren’t like that. We can only have const smart pointers, not the objects they are point to.
We can have different classes for smart_pointer<T> and smart_pointer<const T>, which are again different classes in the eyes of the compiler. We can have a templated conversion function between smart_pointer <T> and smart_pointer<const T>, this will be a one way conversion.
And const classes have some behaviors that you can’t do with them, while non const classes do support them, which is suspiciously like public inheritance.

Our base class is the const class template, which defines behavior that both classes can use.
The pointer is now stored in a union class with both a normal pointer and a const pointer (it’s a union, so we use the same amount of storage), the non const smart pointer template is derived from the const version, and provides access to the non const pointer.
Function members now use either the const pointer or the non const pointer. This responsibility is for the user to implement.

Evaluation
With all the problems, should we really use smart pointers?
Yes. we should.

### Item 29: Reference counting.

A technique that allows multiple object with the same value to share one actual representation of the value, it helps us circumvent the issue of ownership and transferring it between objects, we switch to the idea of bookkeeping counts instead, and this acts as a simple form of garbage collection. It also helps us save space and get better speed, if the value is the same, why should we have more than one instance of it?

Imagine this:
String a,b,c,d,e;
a=b=c=d=e=”Hello”;

We want all our variables to be the same and have their value as ‘hello’.

When one object goes out of scope (and has the destructor called), we would want some way to know if other objects are holding the same value or not. If there are none, we should destroy the ‘hello’ string and remove it from memory (prevent resource leaks). For this we employ a special object that all of the other object refer to (share between them), this object holds the count and the actual object.

The diagram above describes the basic idea, the actual object is held by a middleman, who also knows how many ‘real’ objects hold the same value. Each of the ‘real objects’ interacts only with the middleman, not with the actual object.
Implementing Reference Counting
A basic trick, we create a struct (a class with default public access) inside a different class private members. This way it’s completely transparent to the member functions, but unreachable for the outside user. Here is the simple design.

We define the stringValue as nested class of the String class, and we store a pointer to it. Because this class offers the same functionality as the real string class, we should keep the name. But we should really push it inside some namespace.
And here is a base implementation of the string Value nested class:

The actual job of incrementing and decrementing the counter is done by the outer class, which can access all the parts of the struct.
The regular constructor allocates the memory and copies the string.
The string copy constructor simply copies the reference and increments the counter.
The string destructor decreases the reference count, and deletes the inner class only if the value of the reference count is zero.
The assignment operator decrements it’s own value, deletes if needed. And then takes the one from the other objects, and increments it (don’t forget to self check!).
Copy-on-Write
With the string class, one of the allowable operators is [], which can change the value of the string, so we would need to first copy the string value and then change it.
We also need some way to tell apart between reading and writing
Char c= str[2]; //read, no problem
Str[2] =c; //write, bulk of work needed.

C sharp has indexers and can tell when something is used for ‘reading’ or for ‘writing’ (setter and getters), c++ isn’t so lucky, and we need to do this ourselves.
We can either use proxy classes (next item), or treat all non const [] as write operations, so each time we have a [] operator, we assume this is a write operation, and we copy the ‘real string’ and the current object will now be the sole owner of a string.
This is copy-on-write and lazy evaluation combined.

Pointers, References, and Copy-on-Write
Here is a problem we haven’t dealt with.

We said the p modifies string s1, then how come it modifies s2 as well?
The problem is that we allowed unconditional access to the inner string, so the user can do whatever he wants, and we can’t stop him!

Three ways to deal with this.
Ignore it and pretend it doesn’t exist.
Ignore it, but write in the documentation that the user shouldn’t do it, and that it can cause problems.
Handle the problem by marking objects as safe or unsafe to share.

The 3rd option means that once a [] operator was called on an object, we assume that it was compromised and we will no longer allow any other objects to share this string. Any new string (copy constructor, assignment operator) will copy the data and have a new counting object (which would be safe to share). This means that any objects that shared the data before the [] operator was called are suspect to changes, but new objects deriving from them will be safe.
A Reference-Counting Base Class
Rather than re-writing this code for each type of object we with to share, we can write a base class and have other classes inherit it.
We mark this class as base class by making the destructor virtual, and we make the class abstract by making the function=0;
(this isn’t a pure interface class, as it hold data members).

We push all the logic of the counting to the base class, the derived class (which is nested somewhere else) holds the data itself and is the one who adds and remove references, and the function to remove a reference can actually delete the entire object! (“delete this;”).

Is it weird that nested class inherits from a class that isn’t related to the outer class, yes.
But that’s life.
Automating Reference Count Manipulations
We still have some problems with our code, mainly that the ‘string’ class member functions have to do some of the work with calling the reference counting object. We don’t want this. We can handle most of the problems by using a version of a smart_pointer. Not all of them, but most.
We wrap our class inside a Reference counting smart templated pointer, this takes care both of the reference counting bits and the operator overloading for \* and ->.
The problem is that for this to work we need to assume that the class we inhibit has a deep copy copy constructor (which makes a completely independent copy of the object), and we have the problem with derived class.
Even though our Reference counting templated smart pointer can hold class T pointers and any of the derived subclasses, the copy constructor will still assume it has a pointer of type T, and will create a copy of this class, not of the base class (hello slicing our old friend!). We might avoid the issue by having a ‘virtual constructor function’ - a clone function.
The rest is easier, the destructor calls the remove-reference of the counting object, which is in charge of deleting itself if needed.
Putting it All Together
Here is the diagram

The user sees a String object, which has a member of class RCPTR (smart pointer)who points to a string Value objects, which inherits from the Reference counting object.
We no longer have any need to write copy constructor or an assignment operator for the string class. The automatic generation will call the corresponding ones on the smart_ptr RCPtr object,
The string class defines the base constructor and any functionally that the user needs to use.
The RCObject base class handles the reference counting behavior of increasing and decreasing stuff, with the StringValue derived class handling the actual string data.
The smart pointer class modulates the calls to the stringValue class through RAII principles.

Note that we haven’t changed anything with the public interface of ‘string’, this is still the same behavior, only that the implementation has changed.
Adding Reference Counting to Existing Classes
If we can’t change the class we use, we can create a wrapping class that does the same work as before, we simply need the smart pointer object to know about the actual object and perform the calls.

Evaluation
Reference counting is a more complicated class than simple use, and it improves performance only in cases when we have many objects with the same values

Some problem with circular references that isn’t solved by the simple case.
Something about memory on the heap

### Item 30: Proxy classes.

Cpp doesn’t really support multi-dimensional arrays, we can get by, but stuff fails, both at compile time and even if we try with new. Isn’t this weird?
Implementing Two-Dimensional Arrays
One way to make this work is create templated class for dimensional Arrays

Now we got a specialized class that does the work in terms of creation, that’s nice.
But if we want to use the classes, we still have a problem, we want to be able to properly use the [] operator,
array2D<int> arr(10,20);
Arr[5][0] = 2;
Why shouldn’t we be able to do so?
We can’t create new operators, so operator[][] is out of the question.
We also can’t overload operator[] and make it accept a different amount of parameters, so having it accept two arguments is also out of the question.
We can override the () operator and have it take whatever we want, so we can change our code to
arr(5,2)= 6;
But this syntax is alien, we want our 2d array to feel like an array, not like some weird hybrid thing that is a function call that gets a value.
Turns out that we can actually get the behavior we wanted from the start, with [][]. It’s just that we need a proxy class.
The expression arr[5][2] is actually (arr[5])[2]. the element in the 2 index position of the index 5 array.
So what we do is create another class, the Array1D class, which knows itself how to handle a subsequent [] call.
The end user doesn’t know that there exists an Array1D class,and he should never really hold one or declare it. The proxy class is a surrogate that we create to bridge some problems.
Distinguishing Reads from Writes via operator[]
Other the Ndimensional arrays, we also use proxy classes for other cases.
Before the explicit keyword on constructors, a proxy classes were used to prevent unwanted implicit type conversions
Class A
{
A(B b);
Class B
{
B(int);
};
}
In this example, A can be created only with an argument of class B, and class B can created with an int. This means we can create A with an int argument (which constructs type B), but if we expect a class A argument, we can’t pass an int, because that would involve two user-defined type conversions (from int to B to A), which is not allowed.

The more common use is to control the [] operator.
This allows us another layer of protection for our reference counting class, and it involves the concept of lvalue and rvalue.
Rather than assume any call to operator [] is a write and therefore unsafe, we can delay the decision, instead of returning an immediate reference to the object, we have a proxy class and we wait to see how it’s being used.
If it’s used as rvalue (right hand side, not left hand), we are safe.
If it’s used as an lvalue (left hand side) we can take precautions and do copy on write. Great.

When we are used as an rvalue, the implicit conversion to char takes place, great,and if we are used as a lvalue, the operator= comes into play, with the behavior it defines.
The behavior it defines is to modify itself and change the string it contains, here we finally copy the string, when we are absolutely sure that we can’t keep using the same reference counted objects as before.
Limitations
This technique is nice, but not perfect. Proxy classes aren’t the same as our classes, so there are still places this falls.
If we assume that [] returns a reference, we can take the address of it and get a pointer,
If [] returns a proxy class, this is no longer true.
We can overload the address operator for the proxy class and make it return the real address. It’s simple for const proxy classes, but for non const versions, we are back at the ‘treat all [] as write operations’ situation. We separate our class from the rest of the shared pool and mark it as unsafe. Now any changes the user makes are contained in this instance.
Another problem is the case of templates, back to the Dimensional array, if we don’t define all the operators of +=,-=,\*=... they simply won’t exist. This stands for whatever function we failed to overload in the proxy object.

Another problem is that proxy object fail when we pass them to functions that take references to non-const objects.also in cases when we have an implicit cast from the outer object to some other class. This is the reverse of the earlier use of proxy classes, we can convert from the proxy to a real class, but can’t convert any further.
Evaluation
Some advantages, some disadvantages.

### Item 31: Making functions virtual with respect to more than one object.

Imagine having a function that takes two arguments which can be derived types, and it’s execution depends on their dynamic types.
This is a mess. We need the function to be virtual on both it’s arguments, not just one.
CLOS (Common LISP Object System) offers more control on how calls to overloaded methods are resolved. But in c++ we don’t have this direct ability. Instead we use something that is called ‘double dispatching’, based on the idea that a virtual function is actually a ‘message dispatch’. So to get two a function to work with two ‘virtual’ functions, we need to work out a system ourselves.
Using Virtual Functions and RTTI
Virtual function implement something called ‘single dispatch’, which is then done by the compiler.
To create our double dispatch, we start with a regular virtual function in the base class.
The most basic way is the cruel world of breakin encapsulation, using RTTI and having a chain of if-else statement (if switch case, still the same).

In the code above, get the the dynamic type (type_info class) with typeid and we start comparing it against known possible cases. Because we are inside member function, we know our type (spaceship), and we figure out how to process the other options.
The problem is that each new class derived from game object must implement this extensive list, and each list must include each new class. This is a maintenance hell and also a complete break of encapsulation.
This type of code is actually the reason why we have virtual functions in the first place (and part of what made C code unmaintainable), we want to compiler to warn us when we call non existing functions, not runtime errors and throwing exceptions!.

Using Virtual Functions Only
Before going full ‘double dispatch’ there is a way to attack the problem with nothing but virtual functions.

The virtual function overload actually contains just this:

This calls the virtual function of the other object with the calling object as the parameter.
The beauty is that now we know both types of the objects.
The ‘this’ argument is resolved because we are inside the class scope, and because we use a virtual function of the other object, we will get the correct types each time.

The advantage of this approach is that we can drop the horrible if-else list, as types are resolved via the virtual functions, we just need to define the correct function behaviors. And if we make them pure virtual methods of the base class (define the signature and = 0) then we can even get compiler errors if we forgot to override it in some class.
The problems with this approach is that we still break encapsulation (to some extent), as each class needs to be aware of all classes in the inheritance chain.we also might have problems with modifying the base class. In many times this means re-compiling and re-testing everything, which might not be possible or practical. It also completely violates the open-close OOP principle.
Emulating Virtual Function Tables
The best way to solve this issue is to redesign our program and avoid needing this behavior. Failing that, we take our chances with the next approach: we emulate the way vtables behave and isolate all our RTTI and encapsulation breaking code in one place.
We now have different functions for each derived class, and a function for the base class.
For convenience, we create a lookup function that returns a function pointer to a member function,

And this is the definition of the collide function:

Inside our lookup function, we create a static map<string, function> object - a map (associative container) that takes a string as key (because the result of typeid is a the dynamic name as a string. usually.) and returns a member function pointer.

Together with collide function from above, we have a function that calls the lookup,if we got a valid result, we run it. If we got no match the collide function throws an exception.
Initializing Emulated Virtual Function Tables
The function map needs to be initialized only once, so the obvious way is probably a static function that creates and initializes it. But if do it directly we might have a resource leak or copying costs.

The tricks is to use smart pointers.

But there is still some problems with types and compiling.

We can use reinterpert_cast<>() to force the compiler to accept our code, but lying to compilers is usually a bad idea, and can result in crazy behavior when we having multiple inheritance or virtual base classes. Reinterpert_cast works when assume the memory layout is ‘normal’, but if the layout is different, the compiler has no way to defend against, here is an example:

The address of the entire object is the same address of the vtable pointer (vptr), but the base class we want to cast as has a different layout, and the vtable is at some other address. In all likely cases, we will get garbage at run time.

The solution doesn’t involve lying to the compiler at anyway. It also explains why we had different function names for each object (and not overloadings). We change the member functions to take the base class instead (and not the derived class), and now the mapping will work without problems.
Inside each function, we use dynamic_cast<>() to cast the base object reference to the correct form, and we can now can proceed as usual.
Using Non-Member Collision-Processing Functions
The solution above solved the problem of type identification, but we still break encapsulation by having a new function for each new derived class, and we still need to recompile many times.
The trick is ofcourse to have non member functions handle the business. This moves the responsibility away from classes themselves, and keeps all the relevant code in one place.

We just define a new function for each combination (both left hand side and right hand side, but one should call the other) of types, with the parameters still being the base class types.
The example uses unnamed namespaces, which makes the code visible only inside the current translation unit. Only the function that the other classes call is visible to them.
We have some small changes the relate to the fact we are using non member functions (our typedef just got even uglier), we use make_pair() extensively to avoid declaring the type of pairing<string, string> each time.

Inheritance and Emulated Virtual Function Tables
A final problem occurs with levels of inheritance.
Imagine this:

We would want both types of spaceships to use the same function when interacting with a different class, but with the current design, this can’t happen.
The name of each class (from the typeid) is different, and pays no respect to inheritance.

In this case, we give up on the approach and return to double dispatch with virtual functions, which knew how to deal with inheritance. Sometimes that’s how life is.
(isn’t there a way to get the base class and try again? Like with a loop or something?)
Initializing Emulated Virtual Function Tables (Reprise)
There is another option we can employ. Rather than have our mapping static, we can allow dynamic changes to it, this way we can change the behavior of the program during runtime.
We do this by having functions that can change the map, and having classes that use those functions. And we can probably make it work with functors as well.

Miscellany

### Item 32: Program in the future tense.

Write code that is ready to be changed.
If the program has constraints, express them in code, not just the documentation (if a class should never be derived, make it so. And if it should have only one instance, make sure it can’t be created more than once.
Be careful and consider which functions should be virtual and which not.
Anything that someone can do, will be done. So hide away stuff we don’t expect to be used in the private members, so the compiler won’t generate default copy constructor or assignment operators.
Strive for portability, encapsulate, use unnamed namespace to disallow calling of functions, be careful with virtual base classes, avoid RTTI and if-else lists.
Remember that we need virtual destructors when we play with class inheritance and pointers.

### Item 33: Make non-leaf classes abstract.

The problem with mixed types assignments…
Class animal.
Class chicken:animal.
Class lizard: animal.

Chicken ck;
Lizard liz;
Animal *p1 = &ck;
Animal *p2 = &lz;

\*p1=p2;

This seems wrong, but yet, if the assignment operator is virtual, we can do this.
As the definition for it is to accept any kind of animal,
This is the horrid world of runtime type errors.
We can use dynamic_cast<>, but that’s a whole mess.
We can make the = private, but that’s a mistake as well, because assignment operators are used by the derived classes well.

We can avoid some of the problems by making the base class abstract.it cannot be initialized, and no object can call directly the operators, so we get some safety from mixed assignment and partial types.

The destructor is a great contender to be a pure virtual function, but we actually still need to provide an implementation, because destructors are called by the derived destructors.

Having abstract base classes is usually a good idea, as it helps us understand what is really our shared properties, and what is not.

### Item 34: Understand how to combine C++ and C in the same program.

Mixing c++ and c has the same issues that mixing c code from different vendors does. We need to know that the architecture and the calling conventions are compatible.
name mangling
If we combine c and c++ cope, we need to tell our compiler not to mangle function that will be with c or taken from a c library.
To do so, we use the extern “C” directive at the function declaration.
There is no ‘extern “Pascal”’ or anything like that, the extern is a general statement that this function is to be used as if it were written in c.
It also has other meanings, but firstly, it means that there will be no name mangling.
It also applies to c++ code we want to put in a library for other programs to use.
We can declare everything as extern C with the curly braces.
And the #ifdef \_\_cplusplus trick that makes the code work with both c and c++

There is no such thing as standard name mangelig.
initialization of statics
The main function is the entry point of our code. But the compiler does a lot of stuff before hand, and afterwards. This is where the static initialization, global objects and other stuff happen.
This behavior depends on our code being compiled in c++, if it isn’t then we can’t be sure this initialization and destruction happens.

dynamic memory allocation
New goes with delete.
New[] with delete[]
Malloc with free.
Don’t confuse them.
This happens in functions that return memory that the receiver needs to release, like strdup.
Strdup isn’t standardized, so who knows how the memory is supposed to be released?
Nobody.
So RAII when you can, and don’t use stuff that aren’t in the standard.
data structure compatibility
C doesn’t know about c++. C++ does know about C, so if we want to talk between the languages, c++ should only use stuff c understands, that means primitive, pointers, structs, regular functions…
But not virtual functions.
Virtual functions change the memory layout and add a vptr that C can’t understand.
Also inheritance can make a mess of the memory layout, so we can’t have it.

## Summary

### Item 35: Familiarize yourself with the language standard.

The first C++ standard was released in 1990. We refer to it as ARM (annotated c++ reference manual). Today we use the ISO/ANSI standard.
(this chapter isn’t that interesting… it seems outdated)
To note: almost everything is a template, even std::string is actually a templated class of basic_string<char>
And that is before we mention the allocator.
The stl - standard template library: containers, algorithms, iterators, and more.
