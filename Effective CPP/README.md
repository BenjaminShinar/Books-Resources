# Effective C++

Effective C++: 55 Specific Ways to Improve Your Programs and Designs

## Introduction

## Chapter 1: Accustoming Yourself to C++

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

## Chapter 2: Constructors, Destructors, and Assignment Operators

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

## Chapter 3: Resource Management

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

Chapter 4: Designs and Declarations

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

Chapter 6: Inheritance and Object-Oriented Design
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

Chapter 7: Templates and Generic Programming

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

Chapter 8: Customizing new and delete
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

Chapter 9: Miscellany

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

More Effective C++: 35 New Ways to Improve Your Programs and Designs
Introduction
If in question, name pointers (int main) with “p” as first letter and references with “r”, also remember lhs and rhs (left-hand and right-hand sides) as common naming conventions.
We can leak stuff other than memory, file handlers, mutexes, semaphores..

Basics

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

Operators
When and how to overload operators, and more importantly, when not.

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

Deletion and memory deallocation
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

Summary

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

Summary

### Item 35: Familiarize yourself with the language standard.

The first C++ standard was released in 1990. We refer to it as ARM (annotated c++ reference manual). Today we use the ISO/ANSI standard.
(this chapter isn’t that interesting… it seems outdated)
To note: almost everything is a template, even std::string is actually a templated class of basic_string<char>
And that is before we mention the allocator.
The stl - standard template library: containers, algorithms, iterators, and more.

Effective STL: 50 Specific Ways to Improve Your Use of the Standard
Introduction

Chapter 1: Containers
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

Chapter 2: vector and string
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

Chapter 3: Associative Containers

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

Chapter 4: Iterators

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

Chapter 5: Algorithms
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

Chapter 6: Functors, Functor Classes, Functions, etc
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
Chapter 7: Programming with the STL
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

Appendix A: Locales and Case-InsensitiveString Comparisons
Locales are horrible.
Appendix B: Remarks on Microsoft’s STL Platforms
Something about early versions of visual studio (6 and below).
