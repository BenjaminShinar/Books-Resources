# Effective C++

Effective C++: 55 Specific Ways to Improve Your Programs and Designs

## Introduction

## Chapter 1: Accustoming Yourself to C++

<details>

### Item 1: View C++ as a federation of languages.

C++ isn’t one language, it’s four languages together, it’s C, object oriented C++, template C++ and the STL. each part has its’ own conventions, styles and pitfalls. They can work together, but we must know how to combine them.

For example. For C basic stuff, we love passing by value (for built-in, C-like types), but in OOP c++, we prefer to pass references, and in templated c++, we have to pass references, but we when go to the STL, we have pointers again,and we go back to passing by value.

### Item 2: Prefer consts, enums, and inlines to #defines.

C++ gives us more power to the compiler, and we should use it rather than the precompiler, when possible.
Rather than having defined constants (#define ASPECT 1.653) we should use cons symbols,
Const is stronger in cpp than in C,
If we have a pointer, we need to use two const specifiers.
And if we use a string, we should use the std::string type.
Make it const.

Making the const part of the code means we can limit it’s visibility and scope. We should define it as a static member of the class.
Class X {
Static const int NumTurns = 5;
};
Older versions of the compiler wouldn’t allow that behavior, and require us to define the value outside of the definitions.

A different way is to declare an enum member, which can also be encapsulated.
Enums cannot be referenced or addresseed, which is more similar to #define.

We should also avoid using #define to create macros.
The following example is one possible fuckup:

It seems innocent, but side effects
f(++a > b ? ++a,b); we do ++ twice.
f(++a > b+10 ? ++a,b); we do ++ once.

Here we should use a templated function, which gets const references, and the ++ will happen just once.

We still have #include, #ifndef,#define,#endif.
But we should reduce our use of the preprocessor.

### Item 3: Use const whenever possible.

Const means that the object shouldn’t be modified. And it asks the compiler to help us enforce this rule
We have many ways of using ‘const’:
If the const is to the left of the asterix, it’s a pointer to a const value.
If the const is to the right of the asterix, it’s a const pointer to a value.
If on both, const pointer to a const value, basically immutable.

In terms of the type, the side of the const doesn’t matter. Use the spiral rule to determine.

Iterators in the STL are modeled on pointers, so something about that.
In a function, const can be the return value, parameters, or the function itself (as a member function).
Sometimes we want to return a const objects, this protects us from some weird behaviors.

Const member functions, if it’s not a const, const objects can’t call it.
Constiness of a member function is part of the signature, so if we want to override it, we should make sure to keep the constiness.

If we didn’t have two versions we would either open ourselves to attacks on our objects (a[0] = 0;) or reduce the ability to use the object.

Bitwise and logical constness
Bitwise - no bits are changed, easy for the compiler to see.
Logical - a philosophy: a const member function can change the bits, but only in ways the client can’t tell.

We might want to change some private members to provide better performance, but constiness stops us!
For that we have the ‘mutable’ keyword. We set it for data members (private!) and it means we can modify these members even inside const functions.

Avoiding duplication
We might have const and non const functions with almost identical behavior, which we would want to avoid writing twice.
One solution is to move what we can to private functions.
Another solution is to cast consintess away.
This is the one place we can should use cast.
What we do is take the non const version, cast it as const (static_cast<>()) and the const_cast<non const>() the return value.
After all, if we are in a non-const function, it means the user has a non const object, so we aren’t doing anything wrong.

### Item 4: Make sure that objects are initialized before they’re used.

Uninitialized values are dangerous. The rules for when they are given default values are too complicated to care about.
For objects, we should always initialize everything. And do it it in the member Initialization list, not in the body of the ctor.this is usually safer and faster.
We should explicitly call the default ctor for the inner members, even if they aren’t given values.

For const and references members, we must use MIL,
We can sometimes use a private function to take care of the assignments, especially if we have many constructors, and we don’t want to clatter the MIL with repetitions, this is only when we can safely use the assignment behavior.

There is a set order of initialization.
First the base class, tehn the members according to the ABEntry list (class declaration?)

The order of initialization of non-local objects defined in different translation units
Static objects, alive from their creation time until the end of the program.
Static inside functions are local static.
Static inside classes, and outside of classes are global (non local).

Translation unit: a compilation unit, a cpp file and it’s #includes.

Some problems i don’t understand. The answer is to call a function that has a local static object, and returns that object. This is the singleton pattern.
This also saves us some performance cost in creating the non-local objects.

</details>

## Chapter 2: Constructors, Destructors, and Assignment Operators

<details>

### Item 5: Know what functions C++ silently writes and calls.

The compiler creates a default constructor, a copy constructor, a copy assignment operator (=) and a destructor,
They are created only if needed (used), but they almost always creep inside.

If we have a member that is a reference, c++ won’t create a copy assignment operator.
If we have const members, the compiler also won’t agree.
Copy assignment operators aren’t inherited from the base class.

### Item 6: Explicitly disallow the use of compiler-generated functions you do not want.

If we don’t want to allow copying (or assignment) of the object, we must explicitly disallow it.
If we don’t declare them, the compiler create them for us, and if we do create them then, the compiler won’t protect us from using them!
The trick to avoid this problem is to make the functions ‘private’. And thus we prevent the compiler from auto-generating them, and we also get it to avoid compiling code that tries to call on them!
We should also just declare them, and leave them without definition! This will get us linktime errors if some private function tries to call them, or a friend member / function tries.
If we would want to get a compile time error, we could push this invalid functionality into a base class, we make the allowed operations protected (inherited) and the actions we want to disallow are private. This means our derived class won’t be able to call on them,

And the code won’t even compile!

### Item 7: Declare destructors virtual in polymorphic base classes.

Make destructors virtual if there is at least one virtual function.
If there aren’t any virtual functions, the class probably isn’t meant to be a base class.

### Item 8: Prevent exceptions from leaving destructors.

We should never have two exceptions active at the same time, c++ might terminate or have other undefined behavior.

If we have an exception (that we can’t handle) inside a destructor, we have two options, either abort everything with an explicit call to .abort(), or catch it, log it, and hope that the exception is contained, borth aren’t great options.

A suggestion is to move the responsibility away from the destructor and allow the user to call on this function, so he could see the exceptions. Our destructor will also call on these functions, but it’s a last resort, not the desired behavior.

### Item 9: Never call virtual functions during construction or destruction.

The vptr is of the most derived class, so we can never know if we call a virtual function, whether or not it has what it needs. If it’s inside a ctor, the data it’s using hasn’t been initialized, or if it’s a dtor, it might have been deleted already!

It might not be easy to find this, especially if the call to the virtual is actually inside a private function (which we created to avoid code duplication).

### Item 10: Have assignment operators return a reference to \*this.

Assignments are right associative,
We can chain assignments together.
Cpp allows us to do this:
Int x,y,z;
x=y=y=z=15;
Which makes them all 15. This is bad code writing, but legal.
To keep the behavior, we we need to make our =operator return a value, this value is the \*this, the result of the assignment.
This is the convention for all operators, we should deviate from in only if we have a good reason to.

### Item 11: Handle assignment to self in operator=.

We need to make sure that x=x is a legal and working behavior in our code.

An example of a class with a private member. The = operator means we need to get rid of the left side (we are replacing it’s contents, so we make sure to clean up the memory), and then replace it with a copy of the right hand side object.
This is unsafe:

If rhs and lhs are the same, then we got ourselves in problems.
We can do identity check, but this isn’t always enough.
There is also the thing about exception safety.
So one option is to hold a reference to the old object, make sure the new one is copied successfully, and only then delete the old one. This handles both the exception safety and self assignment.
Another way to do this is by ‘copy and swap’. We define first create a copy of the desired object (rhs), swap it with the current one, and then release it. This means we make use of the copy constructor.
(or if we pass something by value, we already get a free copy).

### Item 12: Copy all parts of an object.

The compiler won’t warn us about partial copying. So if we add a member, we need to add it to the copy ctor/ copy assignment as well.
This is also a problem with derived classes, as we will need to explicitly copy the base class members. Otherwise we get partially copies with the old base or the default values instead.

In general, we shouldn’t have the copy constructor or the assignment operator call one another, if we have too much code duplication we can use a private member function, and make sure it’s safe.

</details>

## Chapter 3: Resource Management

<details>
<summary>

</summary>

If we take something from the system, we must return it.
Memory on the heap, file descriptors, locks, sockets and all sorts of other resources.

### Item 13: Use objects to manage resources.

We need to make sure that we have resources in stuff that gets its destructor called automatically, not only through delete.
One way to so is with the standard library auto_ptr (smart pointer), which knows to call the destructor on what it’s pointing to.
(the pointer is on the stack, so it gets released, and it’s destructor calls delete on what’s it’s pointing to).
Resources are acquired and immediately turned over to resource-managing objects.
Resource managing objects use their destructors to ensure that resources are released.

Because auto_ptr deletes the object from the heap when auto_ptr goes out of scope, we can’t have multiple copies of auto_ptr, so auto_ptr has some weird behavior with the copy ctor and the copy assignment operators,

This means that they aren’t always the best way of doing dynamic allocations, and STL doesn’t allow the weird copy behavior.
An alternative is the reference-counting-smart-pointer (RCSP), which keeps track of how many pointers refer to each object and delete the object when the last on goes out of scope (like the garbage collector in java), and it works for most cases.

A special note is that both auto_ptr and shared_ptr call delete, and not delete [] on their members, so this can lead to problems!
The boost package has some implementations of auto_ptr and shared_ptr for arrays. If we need them and for some reason std::string and std::vector<> aren’t enough for us.

Here’s how i imagine auto_ptr look:
template<typename T>
Class auto_ptr<T>
{
public:
auto_ptr<T> (T t){m_ptr=tl;}
~auto_ptr<T>() {delete m_ptr}
Private:
T\* m_ptr;
}

Also probably has some fancy ways to derefences it so it behaves just like T.

### Item 14: Think carefully about copying behavior in resource management classes

We should avoid copying managerial objects. Always. This is the same problem that led to creating the shared_ptr, but it’s not always enough.

We can either:
Prohibit copying all together
If it doesn’t make sense for an object to be copied, simply don’t allow it. Use the suggestions in ### Item 6 to stop the compiler from generating copy constructors and copy assignment (base class, private members functions).
Reference-count the underlying resource
Copy, but increment a counter and manage the behavior, then we can add 2nd optional parameter to the class that specifies the ‘deletion’ behavior.

Copy the underlying resource
Sometimes we really want a copy, and we just want it to be managed as well. In theses cases, we also copy the underlying resource and create and new ‘managerial’ object, this behavior is ‘deep copy’.
The new object is independent from the old object, and simply happens to have the same data.
Transfer ownership of the underlying resource
This is the behavior of auto_ptr, only one resource exists at all time that can control the object, the rest can’t.

### Item 15: Provide access to raw resources in resource-managing classes.

We can’t always use simple managed resource, even if we really want to. Many API’s require the raw pointer, not a managed one.
We can either explicitly give the pointer with .get() command, or have an conversion function (operator RawName() const {return raw_pointer)), and overload the \* (dereference) and -> operators.
This can lead to other problems,

### Item 16: Use the same form in corresponding uses of new and delete.

The problem with delete vs delete[].
How many destructors are called?
The memory structure for an arrays is different from that of a single object.
If we use delete[] on a single item, we can fuck things up. Even an array of 1 is different than a single object in terms of the memory layout.

Therefore, we should always match [new [] and delete[] .
This can get weird in tyepdef.
We might typedef an array[], and then we have no problems calling new on it, but when we call delete we mess up, as we should have been using delete[].

### Item 17: Store newed objects in smart pointers in standalone statements.

There are some weird cases in which we might get a memory leak if we try to initialize a smart_pointer as part of the expression and another part of the expression causes an exception before we pass control of the object to the manager

Assume that we successfully create the new widget, and then the compiler decides to run the priority() function and not call to initialize the shared_ptr.
If the priority function throws an exception, then the call to create the shared_pointer won’t run, and no object will take responsibility on the widget, and it won’t be released!

</details>
## Chapter 4: Designs and Declarations

<details>

### Item 18: Make interfaces easy to use correctly and hard to use incorrectly.

We should make our interfaces (prototypes) easy to use, and be aware of what errors the users might do, and be ready to protect against them,

If we expect arguments from the same fundamental type, but with different meanings, our users might confuse the order and provide unreasonable values.
We can avoid this by having specialized structs for arguments,

We have class Data, with three int members,
So instead of a constructor
Date(int day, int month, int year)
And the user calls
Date(03,07,1987)
We can have a constructor
Date(Day day, Month month, Year year)
And the user calls
date(Day(03), Month(07), Year(1987));
This acts as a reminder to the user how to order the arguments.
We can also use enums, for integer like types, or have static predefined sets as static class members.

We should restrict possible behaviors of our classes and try to make them as similar to predefined ones, consistency is important, and we should strive for that.

An example of a factory class that returns a pointer to a dynamically allocated object.
This requires the user to either delete it after he uses it, or store the reference in a smart pointer.
We can make the users’ like easier by returning a smart pointer object directly, so the user must store it in one, and we prevent memory leaks.

An example with a deleter.
Rather than expect the user to use the correct ‘deletion’ function, we bind the function to the shared_ptr so it comes pre-built with the means to destroy it.
This prevent some other problems.
We can look at ‘boost’ to see how they implement the tr1::shared_ptr class.
It’s bigger, slower, uses dynamic memory, but all those runtime costs are nothing compared to the gains in preventing client errors.

### Item 19: Treat class design as type design.

If we write cpp as an object oriented programming language, we need to think about how our classes are defined. Clases will behave like built-in types, and we should give them proper consideration.
Important questions for us to consider
How should objects be created and destroyed?
Constructors, destructors, new, delete, new[], delete[].

How should object initialization differ from assignment?
Don’t forget the copy constructor and the assignment operator.
What does ‘pass by value’ mean for the object?
The copy constructor again.
Are there restriction on legal values for our type? What are they?
What do we do with illegal values? Who checks for them?
How does the type fit in terms of inheritance?
Can we inherit from a different class? Do we expect to be a base class? Should our destructor be virtual?
What kind of conversions are allowed for the type?
Can we be converted to other types? Explicitly or implicitly? Don’t forget the constructors again!
What operators and functions make sense for the type?
What behavior do we want for our type, with whom should it interact?
What standard functions and operators should be disallowed?
What behavior do we want to cross off? The usual suspect are the copy constructor, assignment operator, sometimes constructors and other functions should be private.
Who should be able to access members of the type?
Who is the audience for the class, how does he interact with the class?
Member functions, friend classes and functions, nested classes?
What is the ‘undeclared interface’ of the type?
What about mutexes, exceptions, dynamic memory, performance?
How general is the type?
Maybe we actually need a templated class?
Is a new type really needed?
Maybe we can use one of the existing types and classes instead? Maybe a new function or two will be enough?

### Item 20: Prefer pass-by-reference-to-const to pass-by-value.

Argument are passed by value, this means copies of the data. For fundamental types there is no difference, but for classes, this involves the copy constructor and can be a costly operation. Especially if we have inheritance involved.
This might mean that we pass an object of tens or hundreds of bytes (copying them all each time, and then destroying them afterwards), just to perform a simple action.
In many cases, this can be avoided by passing a reference (and making it const), now we simply pass a single reference and we get the same behavior.
This also protects us from slicing behavior if we pass a derived class to something expects an base object.
Because pass by value involves the copy constructor, the ‘new’ object is constructed using the base constructor, and uses the base virtual table.
Passing it by reference to const protects us from this behavior.

Passing by reference doesn’t apply for fundamental types and for STL objects, they are usually passed by value, because of the way they are implemented.

### Item 21: Don’t try to return a reference when you must return an object.

Passing by reference is usually good.
Returning by reference is dangerous.

A reference must be to an existing object, object on the stack don’t always exist.
Some bad ideas to avoid all the constructors calls and somehow return reference.

Don’t.
Create a new object, pay the price. Let the compiler deal with micro optimizations.
Only return a reference if you received this reference and your chaining objects.

### Item 22: Declare data members private.

Everything should be private, unless there’s a good reason not.
Avoiding getters/setters isn’t a good reason.

Encapsulation allows us to choose behavior based on needs and constraints, we are never tied to one behavior that can’t be changed. Anything public is something that must remain public.

Protected access aren’t much better than public, we simply don’t know how many derived classes exist, and we can’t be sure.
Anything that isn’t a toy app (or a nested class) should be private.

### Item 23: Prefer non-member non-friend functions to member functions.

Encapsulation means flexibility, we want to reduce the number of accessing ways into our members. If a function only uses public methods, it has no need for private access to the inner workings of the object.
It shouldn’t be a friend function, it should be inside the same namespace.
This keeps the encapsulation and keeps the related functions around, but not breaching (convenience functions, quality of life).
These functions can be declared across different source files, all belonging to the same namespace. We simply take what we need from each header file and this extends the namespace.

### Item 24: Declare non-member functions when type conversions should apply to all parameters.

Type conversion can happen only if the argument is in the parameter list, not if it’s implicit like the ‘this’.
So instead of having
Class A
{
Const A operator+(A other){}//member
}
Const A operator+(A rhs,int lhs){}//non member
Const A operator+(int rhs,A lhs){}//non member

(three functions!)
We can have a an implicit int constructor, and one function
Const A operator+(A rhs,A lhs){}
If either of our arguments are an int (or any type that can be used to create A), then implicit casting will take care of this.

---

If we have an implicit conversion from int (constructor) to Type A, then we can could do
A a1;
Int x;
A aa = a1 +x;

Which will be
A aa = a1 + A(x);

But not the other way around
A aa = x +a1; // won’t work.
But if the operator wasn't a member function, we could do this!

The operator should only be a friend if it needs the inner workings, not by default.
This means we must have implicit casting constructors, so we will have two constructor calls.
It might be bad for performance, but it’s usually worth the cost.

### Item 25: Consider support for a non-throwing swap.

Swap is great for situation of self assignment and for exception-safe programming.

The swap template makes use of the = operator overload and the copy constructor.

(swap takes b:a reference to T, a: another reference to B, and return nothing).
This implementation involves three operations (even four!) that might be heavy.
A copy constructor to create temp.
A assignment operator to copy a=b.
A assignment operator to copy b=temp.
And then a destructor for temp.

In some cases, this behavior might be costly, and unnecessary.
The pimpi idiom (pointer to implementation).
If our objects really hold just pointers to data, why do we need to create all the data (deep copy) each time? And also to call delete before each assignment?
If we could know that our class just holds pointers, we could swap the pointers and be done with this. No need for special copying, constructing, etc…

We can create a specialized template for this class,

We usually can’t extend namespace, but total template specialization is allowed, so we could write something like:

This fails to compile not because we extend the std namespace, but because we are trying to access a private member.
We can declare this swap function as a friend to to our class,
Class {
Public:
Friend void std::swap<Widget>(Widget &A, Widget &B);
}
But the convention is not to do so.
The convention is to specialize the swap as member function, and have an namespace std extension in our class.

We specialize the template version to call swap, and that’s all we need.

Swapping templated classes
C++ allows for partial specialization of classes, but not for partial function templates.

We need to fiddle around with the code to find a solution,
We usually can overload function templates, but not for std namespace.

The better solution is to still have a non member swap that calls a member swap. But we don’t declare this function to be an overload of std::swap. it's a template by itself in the class source files.

The name lookup rules will find this version of swap before the std::swap, and this will be called. This is true

Another special thing to remember:

This means we ‘import’ the std::swap function into the scope, and now we call scope.
If there is a swap function for the objects, it will be used.
If there is a overloaded std::swap function, it will be used
In the worst case, the regular std::swap will be used.

We should not qualify this call, as we want to allow the linker to find the best version to use.

Steps:
Write a public swap member function, that doesn’t throw exceptions.
Have a non member swap in the same namespace,
If we have a a class (not a class template), add a std::swap overload that calls the swap member function.

If we’re calling swap, make sure to have using std::swap to make sure we can find the best version.

We don’t want the ‘member’ swap to throw exceptions, ever.
If we can’t assure this, we don’t have a fitting candidate for a swap function.

Chapter 5: Implementations
Some problem that occur while implementing our code.

### Item 26: Postpone variable definitions as long as possible.

We should wait with defining and declaring variables until we are sure they will be used.
This is especially important if we might throw an exception somewhere in our code before this variable will be used (and then we payed for both the constructor and the destructor). We would prefer to initialize our variable with the needed values and not default construct it and then change it’s values.

This not only improves the code performance, it also makes it easier to read, as the variable is introduced in the context of the variables it takes.

The way to go with variables inside of loops depends on usage.

But unless we know for sure that the assignment is very cheap and this section is performance critical, we should use approach B, and construct a new object each time.

### Item 27: Minimize casting.

A reminder of casting styles
(T) var; // c-style
T(var); //cast expression
And the c++ style conversions, or new-style.
static_cast<T>(var); //force implicit conversions, that aren’t necessarily promotions, general casting behavior.
dynamic_cast<T>(var); //allows for safe downcasting, heavy runtime costs.
reinterpret_cast<T>(var); //low level cast, can do pretty much anything, not safe.
const_cast<T>(var); //the only c++ style cast to remove const from a value.

We can still use the old style casts (c-style, cast expressions), but it’s better not to. The modern style is easier to identify, easier to interpret the meaning and works better with compilers.

In the text he says that he uses old style casting when he wants to call the explicit constructor.

Unlike common misconception, casting isn’t just telling the compiler to treat one type as another, in many cases we actually have run time costs.
Like casting from int to a double, while simple, still means that we need to create a new data, because int’s are 4 bytes and doubles are 8 byes and are totally different. Even int and unsigned int are different and require work. Not a lot of work, but this isn’t free at runtime.

(a cost free cast is reinterpert_cast<T>(), which is a whole bag of worms)

Another example:

We think that this is free behavior, as the base part of the derived class should be at the start of the object. But it isn’t, we still need some runtime work to get the correct address. Remember multiple inheritance?
Even in single inheritance, this can happen. The layout of an object is decided by the compiler.
Here is another problematic code:

We want to call the base class behavior,so we cast ourselves to it and call the function.
This isn’t right.
The function that we call isn’t operating on the current object, but on a ‘casted copy’ of it. We when casted ‘_this_ to the base class, we called a (copy) constructor on the base part of \*this, and then called the member function of that temporary copy.
What we should do it to explicitly call the correct base function, this is similar to ‘super.func()’ in java, if we would have only one base class and no multiple inheritance (and the ‘super’ keyword) in c++.

This suggestion stands double for dynamic_cast<T>().
dynamic_cast<T> is quite slow by itself, it’s designed to fit with both multiple inheritance and dynamic linking, so it shouldn’t be used in performance critical sections.
We use dynamic casting when we have a base class pointer that we suspect to actually be a derived class.
We can avoid this by using containers that store pointers to derived class directly (actually, use smart pointers instead of regular ones), this will work if we store only one type inside the container.
An alternative is to use virtual functions, and provide them in the base class, and just have them do nothing it that case.
(this does bloat the base class)
Neither approach is perfect, but both are preferable to dynamic casting.
We definitely want to avoid cascading dynamic casting, i.e something with many ‘ifs’ and cases for each derived class. It’s both slow (many calls to dynamic casting) and hard to maintain (if we add or remove a derived class). In this case, we should probably use a virtual function.

As a general rule, we should avoid all types of casting, and dynamic_cast<T> the most. If it’s not possible we should minimize the use of casting and hide it away,

### Item 28: Avoid returning “handles” to object internals.

Defensive programming.
Don’t return a modifiable handler (reference) to a private member. This can break encapsulation. We can’t allow the user to change our internal data.
While changing return types to const can work, it still has the problem of dangling handlers.
The handle might outlive the object that created it, and it will no longer be valid.

### Item 29: Strive for exception-safe code.

When using locks, have them wrapped inside a lock object, so that if an exception occurs, the destructor is called and the lock is unlocked and removed. Also, less code.

Function should strive to exception safe i one of three ways.
Basic guarantee - the data is always at a valid state, all the member function can be called and the object can operate as usual, even if an exception occured.
Strong guarantee - if an exception occurred during a call, then the data is in the same state as if the call hasn’t been made. Like atomic transactions, either it succeeded completely, or failed and nothing changed.
Nothrow guarantee- the call doesn’t throw exceptions (we usually can’t say this).

Our code must be one of those, preferably the stronger.

The copy and swap strategy, closely related to the pimpi idiom. We hold all the data we might change in a different object, and once we are sure the new object is functional, we swap it in place of the new one.
We need to be careful when dealing with non local data, we can’t assure that if we reset our changes, we won’t we be changing that data in the progress.

### Item 30: Understand the ins and outs of inlining.

Inline functions are great, because we can avoid the overhead of calling a function.
However, no free lunches, inline function cause code bloating, more paging, and other stuff.
Unless the function itself is smaller than that of the function call.
Inlining is a request, not a command, we can’t be sure that the compiler does this, we simply suggest. Virtual functions are never inlined.
We can make inline function either by providing the definition inside the class declaration (java style!), or by marking it as ‘inline’ and defining it in the h file.friend functions can also be inlined.

Even though templates and inline functions are similar in that regard, they aren’t the same.
If we try to take a pointer to an inline function,we might end up with two versions,
We shouldn’t inline constructor and destructors.
Debuggers also tend to have problems with inline functions,
We should remember that inlinning is an optimization, not a requirement. It should be decided after the code runs properly.

### Item 31: Minimize compilation dependencies between files.

Header files ‘private’ members are implementation details, but they are included in every file, so we can’t easily change them (even though we declared them private exactly for this reason!)
The problem is that to compile code properly, we must know the size of the object before hand, so we must have room for the private members.
If we can avoid having direct definitions inside our class objects, we can reduce cross file dependency. (this is what have done in C, we provided an interface in the .h file and all the implementation in a separate file).
This is part of the PIMPI idiom.

A bit about having declarations in .h files and definitions in cpp files. We return handles and not actual classes.
Interface classes (abstract classes without any members), factory method (or virtual constructors)

</details>

## Chapter 6: Inheritance and Object-Oriented Design

<details>

Oop in c++ is different than other languages (like java) and it has some special behaviors of its own. Remember, c++ can be used to write oop code, it isn’t oop by itself.

### Item 32: Make sure public inheritance models “is-a.

Public inheritance means a ‘is a’ relationship.
If we write a derived class from a public inheritance, it means we say that our derived object is a base object, and could be used wherever the base class could be used. Our derived object is a specialized form of the base class.
This is similar to java, where derived classes cannot change the access modifiers of the base class.
Every Derived object is also a Base object, but not vice versa.

(the example of the bird base class and the non flying birds like the penguin or emo).

We can either make a class of flying birds that penguins don’t derive from, or make penguins throw an exception when trying to fly.
In general, we prefer compiler errors to linker errors, and linker errors to runtime errors (and runtime errors to undefined / unexpected behavior).

The example of squares and rectangles, is a square a rectangle?
No. the behavior of only increasing the width of of a shape is possible for a rectangle, but not for squares, so a square isn’t a candidate for an ‘is-a’ relationship with a s rectangle. We can’t take a function designed to work with rectangles, give it squares and expect the same results.

The other relationship models are “has-a” and “is-implemented-in-terms-of”, sometimes it’s better to use them.

### Item 33: Avoid hiding inherited names.

We should be careful with our naming, and avoid having the same name in different scopes, and doubly so when working with inheritance.
The name searching scopes works the same no matter the type, or class of the variable.
If we use the same names, we hide the earlier functions, this is very problematic with function overloading. If we hide one version of overloaded function, we’ve basically hidden all versions (even if we hadn’t declared them!)/

We can counter this behavior and make function visible with the using statement.

This will bring the base mf1 and mf3 function into light, with the new function declaration overriding only the functions with the same parameters.

We sometimes don’t want to inherit everything from the base class, this is against the idea of public inheritance, but it makes sense if we use private inheritance, and we only want to inherit some function (not all the overloads of it).
We do this by forwarding our function, or having it call the fully qualified name of the base class function ({base::foo();}),

There is also a problem with inherited names in templated classes (### Item 43, specialized classes might be in use and derived classes might inherit from them and they won’t have the same behavior as other derived classes)

### Item 34: Differentiate between inheritance of interface and inheritance of implementation.

Inheritance might mean inheriting definitions and API (interface), and might mean inheriting the implementation as well. Sometimes we want both, sometimes we don’t.

Virtual, pure virtual and regular member functions.
Pure virtual functions must be declared by any inheriting class (they also make this class to be an abstract class, which is a bonus). If we only have pure virtual member functions, we actually have an interface, and not a concrete base class.
(we can have an implementation for a pure virtual function, we first make it = 0; and then define it. It can be accessed only with the fully qualified name).

Regular virtual functions provide both an interface and an implementation, but we allow the derived class to change (override) how it implements it.

This is usually good practice for OOP, but it hold some dangers, and we might with to make this inheritance more explicit. This is done by making the public function pure virtual, and providing a protected ‘default’ function (non virtual) that the derived class can inline call. Any new derived class must explicitly call on this function to use it.
If we don’t like having separate definition functions and implementation functions, we can have the implementation be defined as the pure virtual function, which means that the derived classes must still explicitly state what function they plan to use, and fully qualify that they with to use the base default version. This design has less functions in the namespace, but we lose the ability to hide the default implementation under the ‘protected’ access modifier.

Non virtual functions should be the same across all classes, not override, not redeclared, not redefined, not anything, we want this variant to be called, not a specialized version (and if some one decide to declare it again, he’s doing a mistake that shouldn’t be done).

We should avoid creating a base class without any virtual functions (especially the destructor, which should always be virtual), or making all of them virtual without a good reason.

Calling virtual functions has a cost, but it most cases, this isn’t what’s slowing the program.

### Item 35: Consider alternatives to virtual functions.

Virtual functions are great, but there are alternatives. Let’s get to know them.
The Template Method Pattern via the Non-Virtual Interface Idiom
Have a non virtual function call a private virtual function. This is called non-virtual interface (NVI) idiom. Or a wrapper behavior. We seperate the non changeable parts (the shared, set in stone) behavior that the base class defines from the smaller portion that needs specializing.
We also separate that ‘how stuff is done’ into the virtual function, but the ‘when stuff is done’ is still controlled by the base class. This enforces order and structure.
A special note is that derived class override function that they can’t access (private functions), which is odd, but legal. We can also make the function protected and expect the derived class to call the base class virtual function themselves.

The Strategy Pattern via Function Pointers
Rather than define each derived class a virtual function overload, we can simply pass each object (or class) a function pointer and have it call that function to do the required work. This means stronger decoupling, and allows different instances of each class to have different behaviors. This design patterns is sometimes called strategy. the downside is that we distanced the calculation from the object,and it now has to use public access to the object or be given specialized access into it and make the encapsulation weaker.
The Strategy Pattern via tr1::function
Rather than a simple (rigid) function pointer, we can send something that behaves the same, it can be a function pointer, a member function, a functor (function object) that returns not just the type, but anything that can be converted to that type?
We replace the function pointer with a tr1::function object (something that is/has a callable entity).

This declaration is typedefed with a reasonable typename, and now rather than just sending functions, we can send much more.
We can still send functions, but we can also send objects and class members.

We use the std::tr1::bind(...) command to force a constant object into a member function, so it now has a ‘this’ member and can be called from anywhere!

The “Classic” Strategy Pattern
A final option is to have the function itself be a class, and specialized behaviors be sub classes of it. This is composition or dependency injection, whatever. It’s a conventional oop approach that doesn’t involve any c++ features, but is recognizable and understandable.

### Item 36: Never redefine an inherited non-virtual function.

Don’t redefine non-virtual functions. Just don’t.
It makes a mess of things, non-virtual functions are statically bound, not dynamically, if we use a Base pointer to hold a derived object, all non-virtual actions will be of the base class. This means that the type of non-virtual function depends on the pointer type, not on the object. This is obviously bad and not what we wanted.
If something needs to be overridden, it’s a virtual function. If not, it’s a non-virtual function that shouldn’t be redfiend.
As easy as that, and twice as important for destructor. Always virtual.

### Item 37: Never redefine a functions’ inherited default parameter value.

The problem: virtual functions are dynamically bounded (late binding) but default parameter values are statically bound (early binding).
While the function that we call is dynamically bounded in runtime by the object type, the default parameter is statically bound at compile time based on the pointer type. This means that when there is no given value we set the default value based on the pointer type, not the object.
So if our base and derived classes have different default parameters, the behavior will change based on the pointer type,
Derived d;
Derived _ dp = &d;
Base _ bp = &d;
dp->do();
bp->do();

The default parameter will be different in the two cases.
We can avoid this situation by using the Non virtual interface idiom (NVI) and have a non virtual function call the virtual function with the default value, making it ‘safe’ to use again.

### Item 38: Model “has-a” or “is-implemented-in-terms-of” through composition.

When the composition is part of the domain logic, we use a ‘has-a’ relationship (a person has name, address, kids.. Etc, but he isn’t a name, address, etc), when the composition is part of the application logic, it’s ‘is-implemented-in-terms-of’ relationship (We implement a thread pool in terms of a queue, but a thread pool isn’t a queue).
In this case, public inheritance isn't the way to go.

### Item 39: Use private inheritance judiciously.

Public inheritance means ‘is-a’ relationship. But what does private inheritance mean?
It means that there isn’t any implicit conversion between the derived class to the base class.
Private inheritance goes with ‘is-implemented-in-terms-of’ relationship, we take advantage of existing classes, without saying we adhere to the interface itself.

We prefer composition over private inheritance, and we should only use private inheritance when we must (because of protected members or in some edge cases). We might want to have base class with some behavior, but not allow derived classes to change it, so private inheritance comes in handy.
This also helps us in some terms of decoupling .
And this edge case about empty classes (no data, no virtual functions, no virtual base classes) , empty base optimization (EBO), only for single inheritance,

### Item 40: Use multiple inheritance judiciously.

There are people who don’t like multiple inheritance.
When we have multiple inheritance, we can inherit the same name from both base classes and have ambiguity.
We can have two instances of the common base class and then we need the diamond inheritance design with virtual inheritance and virtual base classes.
The general rule is to avoid using unless we must.
If we must, avoid putting data inside them (interface classes, like java and c#)

</details>

## Chapter 7: Templates and Generic Programming

<details>
<summary>

</summary>

### Item 41: Understand implicit interfaces and compile-time polymorphism.

OOP programming uses explicit interfaces, and runtime polymorphism - virtual functions, vtables, and declaring types.
In the case of templating, we give less importance to those, and we focus on implicit interfaces and compile-time polymorphism.

The implicit interface is based on the actions taken in the template. If we asked for .size(), it means only types that use have .size() are acceptable type parameters.
Each line in the template is a constraint for the type, it’s an implicit interface.

### Item 42: Understand the two meanings of typename.

Template <class T> and template <typename T> are usually the same.
But not always, and typename should be prefered, because it has some stuff that only it can do.

In a template, we can refer to two kinds of names.

Names in the template that depend on the template parameter are dependent names.
Can also be nested dependant names. In the example, iter is a nested dependent type name. (suppose we pass it a std::vector<int>, then it’s type is std::vector<int>::const_iterator).
The other variable, value, is not dependent, it’s just an int.

We think Iter is going to be a pointer, but this isn’t always the case, maybe we pass to the template a type that happens to have a static member called const_itertaor (not a nested class), another weird thing:

We think we are declaring a variable x from the nested class.
But maybe c::const_iterator is a static member and x is global variable? Then our code will actually be a multiplication expression!

To make sure we parse C as a typename, we add the typename qualifier before it.
This tells the compiler that C::const_iterator must be a type, and not anything else

We should use typename anytime we refer to a nested dependent typename in a template.
This also holds for template signatures.
template<typename C>
Void f(const C& container, typename C::iterator iter){;}
Now we have templated function that takes a type and a nested type,

One exception is that “typename must not precede nested dependent type names in a list of base classes or as a base class identifier in a member initialization list”.

Another thing is that we can qualify one part of the type, even if it’s templated.
Template <typename IterT>
Void Work(IterT iter)
{
typename std::iterator_traits<iterT>::value_type temp(\*iter);
}

Let's unpack:
Templated function that takes a type called IterT.
We want to make a copy, temp, based on the contents of iter (\*iter).
Temp is going to be some type, this type is the value type of the templated class iterator_traits from the STL, so it’s going to be a nested class of the the templated class (from the iterator_traits) class. So we precede it with typename.

There is a convention to create a local typedef definition:

Now we have the horrible line once, and we get a typedef that is easy to read.

### Item 43: Know how to access names in templatized base classes.

We can create templated classes. But there is a problem with creating a derived templated class.
Base template:

Derived template:

The template doesn’t know that base class with ‘sendClear()’ function exists. It can’t know about the base class until it’s instancized, and it can’t instancize it without compiling!

There is a better example using specialized templates.

Instead of writing “SendClear(info).
Three ways to fix the issue:
this->SendClear(info);
Using MsgSender<Company>::SendClear(info);
MsgSender<Company>::SendClear(info);

### Item 44: Factor parameter-independent code out of templates.

When we have templates, we might lead to bloated binary files, even if the code is lean.
If we have a template, we should refractor away all of the independent code. We should treat our one templated function as several functions, and try to identify any expressions that don’t depend on the parameter type, and move it outside.

One suggestion is to create a derived class that uses the base class, but hold the member function, and then do some calls to base function with a parameter.

I need to revisit this ### Item in the future

### Item 45: Use member function templates to accept “all compatible types”.

Smart pointers (what is used as the STL iterators),
Creating a template with a “generalized copy constructor”.

Ths class smartPtr has templated constructor that accepts any other templated smartPtr element. We also avoid making this constructor explicit, so we could use it a casting operator.

Then there’s some part about inheritance and which kinds of conversions we allow.

I need to revisit this ### Item in the future

### Item 46: Define non-member functions inside templates when type conversions are desired.

### Item 47: Use traits classes for information about types.

### Item 48: Be aware of template metaprogramming.

</details>

## Chapter 8: Customizing new and delete

<details>
<summary>

</summary>

Other programming languages offer automatic garbage collection (memory management), but c++ allows us to control the memory allocation to get better performance.
We use new and delete (new[], delete[] for arrays) to get memory, free it and call constructors and destructors. We also have some concerns for multi-threaded environments.

### Item 49: Understand the behavior of the new-handler.

When we can’t get enough memory from the system, old compilers returned a null pointer.
Today, the behavior is different, and if there wasn’t enough memory, the program should call a new_handler function.
This function is a client specific function, which we can set with the set_new_handler function from the standard library (similar to set_terminate and set_unexcpected).
This is a void (void) function that doesn’t throw anything.
If the handler function also can’t find memory, it’s called again repeatedly, which is troublesome, so the handler function should do one of the following
Make more memory available - somehow, it’s suggested that we preemptively save a block of memory for the function to use before starting the program
Install a different new_handler - call set_new_handler again, change something about itself. Whatever.
Remove the handler function, so the default exception is thrown.
Throw exception itself
Not return - call exit or abort.

We can’t have different built-in handlers for different classes, but we can implement the behavior.
We simply overwrite the new operator and the ‘set_new_handler) functions for our class, our new ‘new’ operator will set the class static handler as the handler function (through some resource with RAII behavior), and then it will call the default new operator (::new , the global qualifier) to do memory allocations. If the memory allocation fails, the new handler is called. At any case, the destructor of the local resource inside the ‘new’ operator will return the previous handler function to be the default one.

There is also a suggestion making this into a templated class and having other classes inherit from it so they also get the same managed allocations behavior, without any extra work.
There is some weird parts about having a templated class that never uses it’s typename parameter, but this is done for the sake of having different static members for each class. It’s called ‘curiously recurring template pattern (CRTP)’.

Because it’s c++, we can use this base class as part of our multiple inheritance, we just need to be careful (as we always need to be when using multiple inheritance).
The old behavior of having a null pointer return when allocation fails is still supported with an alternative form of the ‘new’ operator, the new(std::nothrow) overload.

### Item 50: Understand when it makes sense to replace new and delete.

Three common reasons to replace the new and delete operator.
To detect usage errors
We can use our new and delete operators to create safe usage of functions, and to add extra protection from double freeing. We can add a marker to out memory (start and finish) and validate that it’s entacts and that it’s where it supposed to be before deleting, and if we see that something in our marker is not as we set it to be, we can tell right there that there is going to be a memory problem. And in this case we can act before the delete and gather information about it to provide better feedback and error reports (think valgrind).
To improve efficiency
The default new and delete operators are designed to fit everything, so they have to consider many situations, and therefore they take the ‘middle of the road’ strategy. If we know how our objects are to be used in our program, we can handle the requests better and get better performance.
To collect usage statistics
Similar to the first reason, but not just for errors, we can use the operators to gather information about lifetime of our objects, when they are allocated, how much memory we use at any given time… etc.

To write an our version of ‘new’ we need to handle the memory allocation ourselves (hello malloc our old friend! We can also use mmap to be extra special) and we take extra memory to write signatures(remember 0xDEADBEEF? That’s one nice signature, isn’t it?) and whatnot. We then return the memory block with the actual data that user requested.

There are all sorts of other conventions that we should follow that appear later, but we should also consider alignment. Malloc might give us an aligned memory, but we might give our user a non aligned address (based on the offset of our signature data)
Even if our compiler knows how to create a program that isn’t aligned, we can still experience bad performance because of it. This is one of the dangers with manual memory allocation, some compilers have built in alternatives for memory management, or might we use open source allocator libraries (such as Pool library from Boost), this can get better performance, but we might sacrifice portability, TR1 can also help.

And now, here are more possible reasons
To increase the speed of allocation and deallocation
Especially in single thread programs, we can abuse non thread safe allocation (like the fixed size memory pools in Boost::Pool) to get better speed performance.
To reduce space overhead of default memory management
If our memory needs are small, we might be wasting a lot of space with overhead data, and a custom new operator can reduce it.
To compensate for suboptimal alignment in the default allocator
Maybe our default allocator doesn’t work in 8 byte alignment?
To cluster related objects near one another
If we know some data structures or classes tend to go together, we might get better performance if we store them in the same ‘pages’ of memory. If we store them next to one another, we can work faster.
To obtain unconventional behaviors
An example is a case where memory management is done via c code api, and we want to maintain the c++ style of our program, so we hide all the memory details in the new and delete operators to reduce the complexity of the code. We might also want to have better security by manually overwriting deallocated bytes with zero to hide sensitive information.

### Item 51: Adhere to convention when writing new and delete.

As mentioned before, there are some conventions to follow when writing the ‘new’ operator.

We need to consider the case of no available memory.
We need an infinite loop, an handler function, and some way to throw an exception (or abort the program) if we can’t allocate the memory.
There is also a problem of requesting zero bytes, as we must return a legitimate point.

We have an infinite loop of trying to allocate memory, if we succeed, we return a pointer to the memory, if we fail, we do the trick with our class specific handler function (which might be a problem in multithreaded environments) and try again, eventually we will either succeed in something, or we will run out of new_handlers and we will use the one that calls abort() and terminate early.
Ofcourse, if our object is a derived class, it also needs the correct size to allocate for the derived class, we can make sure that the request size in bytes is the same as the size of the class whose new operator we are using, so we won’t end up using the base new operator for a derived class. Of course.

(this check also includes the request for zero sized memory, because even interfaces have non zero size).

We also need to handle the ‘new[]’ operator, (called ‘array new’ in speech), which again has problems with sizes of derived classes (the wonderful problem of object slicing?)

In terms of delete, we need to make sure we handle deletion of null pointers (a simply check and return), and we need the same behavior with derived classes in terms of size and with arrays, and not to forget we must have virtual destructors!.

### Item 52: Write placement delete if you write placement new.

If memory allocation succeeds by the constructor throws an exception, we need to use the matching delete operator.
A ‘new’ operator with parameters (besides the mandatory size parameter) is called a placement version of the operator, such as the new operator with size and void pointer to decide where to construct the object. This version is part of the standard library, inside #include<new>.
The runtime environment calls new, and if it fails, it calls delete to remove deallocate memory. The runtime looks for a delete operator with the same number and type of arguments, but if it doesn’t find any, it doesn’t call any delete if there was an exception.
Therefore, we must have a delete operator with the same arguments as the new operator, and they must work together...

Also, if we declare a different version of new, we might hide other versions of ‘new’ from the user. There are three default versions of new/delete.

If we define any of them, it hides the other versions. We should also declare the corresponding delete operators.
Again, the suggestion in the book is to create a base class that has all the special forms declared and that way they will always be available.

</details>

## Chapter 9: Miscellany

<details>
<summary>

</summary>

### Item 53: Pay attention to compiler warnings.

Read the damn message. Don’t ignore warnings.
Remember that warnings are compiler implementation, so be careful.

### Item 54: Familiarize yourself with the standard library, including TR1.

TR1 stands for technical report 1.
TR libraries usually contain what’s expected to be part of the next c++ release.
Before TR1, the c++ standard contains:
The standard template library (STL), containers, iterators, algorithms, function objects, etc…
Iostreams - better control for input and output
Support for internationalization - more than one active locales, the wide_char and wide string classes
Support for numeric processing: complex numbers template, array of pure values (valarray)
Exception hierarchy - base class of std::exception, with derived class of logic_error and runtime_error.
C89 standard library.
The book details new stuff that TR1 should contain,
Smart pointer
Shared_ptr, unique_ptr, weak_ptr. Class designed around the move semantic.
Tr1::function
Register any callable object, not just function pointers, also function objects and class members.
Tr1::bind
Binding objects as ‘this’ or scope elements.
Hash tables
Emphasis on the fact that the containers aren’t ordered in any way.
Regular expressions
regex.
Tuples
Generalization of the ‘pair’ idea,.
Tr1::array
An array that behaves like an STL array, no dynamic allocation.
Tr1::mem_fn
Changing from member pointer to function.
Tr1::reference_wrapper
Having references behave more like objects, throw them all inside a class.
Random number generation
Better random numbers than what C gave us.
Mathematical special function
Some nice math stuff
C99 compatibility
Features that are in c99.
Type traits
Compile time information about classes, and other nice things.
Tr1::result_of
A way to deduce return types of function calls, templating.

### Item 55: Familiarize yourself with Boost.

Boost is an open source library of c++ stuff.
A lot of stuff in TR1 is based on boost, and boost has many other stuff which is nice, including lambda

</details>
