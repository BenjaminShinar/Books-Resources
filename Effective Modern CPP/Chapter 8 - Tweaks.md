## Chapter 8 - Tweaks

<summary>
The general rules have exceptions to them.
</summary>

There is a general case advice for most situations, but sometimes there is a reason to do things differently.

### Item 41: Consider Pass by Value for Copyable Parameters That are Cheap to Move and Always Copied

<details>
<summary>
There can be situations where passing by value and always creating copies can make the code better, but mostly not.
</summary>

some function parameters are designed to be copied (copy and move are the same for this issue),in general, we should copy rvalue arguments, and move lvalue arguments. in this example, we have two functions, based on the type of the argument. depending on the compilers, we might not be able to inline to functions, which would mean real code bloat, not just source code duplication.

```cpp
class Widget {
    public:
    void addName(const std::string & newName)
    {
        names.push_back(newName);
    }
    void addName(std::string&& newName)
    {
        names.push_back(std::move(newName));
    }

    private:
    std::vector<std::string> names;
};
```

on thine we can do is use perfect forwarding, the removes the duplication in the source code, and even opens up some more possibilities to pass to the function, but the downside is that we might have even more functions, less inlining, and this code must be written in the header file as part of the class declaration. and if the user somehow passes an argument of an incompatible class, we can get the monster sized compiler warnings.

```cpp
class Widget {
    public:
    template <typename T>
    void addName(T&& newName) //universal reference
    {
        names.push_back(std::forward<T>(newName));
    }
    //...
};
```

if we want to make this function clear, short, not intimidating, and only one, we might need to drop a core principle of C++. it's time to pass by value again. we are allowed to call _std::move_ on the parameter, because we know it was created right now for this function,as a copy of the argument, and is completely independent from it. and this is the final use of the parameter. so if someone down the line want to move from it, be our guest.

```cpp
class Widget {
    public:
    void addName(std::string newName) //copy by value
    {
        names.push_back(std::move(newName)); //make it eligible for moving
    }
    //...
};
```

the common reason to avoid pass-by-value was to avoid copies. anytime we performed this operation, we made a call to the copy constructor. the trick in c++11 is that we now also have a move constructor, so things can be cheaper. even if we are creating a new object, maybe it's not that expensive anymore.

```cpp
Widget w;
std::string name("bart");
w.addName(name); // 1. called with lvalue. one call to copy constructor
w.addName(name + "Jeanne"); //2. addName is called with a rvalue, so we create the string locally, but pass it to the function as an rvalue with the move constructor.
```

this is nice, but there are some problems lurking.

if we think back to the three version of the function we considered before(const rvalue, lvalue, templated for universal reference). the first two are 'by-reference-approaches'.

Overloading: one copy constructor, one move. no cost for bounding.
Using A universal reference: again, no cost for bounding. one copy constructor, one move.
passing by value: a new std::string must be created, and then we always perform a move. so one copy constructor plus a move for lvalue , or two calls to _std::move_ for rvalues.

seems to be less good, so why does the item suggest something else?

> "Consider pass by value for copyable parameters that are cheap to move and always copied."

1. **Consider** is a key word, the benefits of reducing code bloat don't always offset the costs of the copies and moves.
2. "Consider pass by value for **copyable parameters**..." - parameters that aren't copyable only have move semantics, they aren't part of this item. if the parameters is move only, there's no need for two functions in the first place. in this example we have a _std::unique_ptr_ (a move only type) as a data member.

   ```cpp
   class Widget {

   public:
       void setPtr (std::unique_ptr<std::string> && ptr)
       {
           p = std::move(ptr);
       }
       void setPtrValue (std::unique<std::string> ptr)
       {
           p = std::move(ptr);
       }
       private:
       std::unique_ptr<std::string> p;
   };
   Widget w;
   w.setPtr(std::make_unique<std::string>("Modern C++"));
   w.setPtrValue(std::make_unique<std::string>("pass By value?"));
   ```

   the cost to call the classic (rvalue) function is one move (from the rvalue created by make_unique to the function itself) before reaching the function itself. the cost for the pass by value function is two moves.one to create the object, and another to move from it into the move constructor. that's twice as many operations again.

3. "...that are **cheap to move**..." if moving is a single operation, the cost of an extra move might not be so large, the benefits from the reduced code bloat and extra inlining might be worth it.
4. "...and **always copied**." we should only consider this approach if we are always copying the value, if the value is checked and validated before the copy, it might not be great.
   ```cpp
   class Widget{
       public:
       void addName(std::string newName) //pass by value
       {
           if ((name.length()>= minLen) && (name.length()<= maxLen))
           {
               names.push_back(std::move(newName));
           }
       }
       private:
       //... minLen, manLen, names vector
   }
   ```
   in this case, even if don't end up pushing newName into the vector, we still paid for making a copy of it, this wouldn't happen had we used references

even if everything holds true (an unconditional copy, cheap to move,copyable type, and we really considered the options), it still might not be a good idea.

#### Assignment vs Construction

a function can copy a a parameter either via _construction_ (using the copy or move constructors) or by _assignment_ (copy or move assignment operator). if the function uses construction, then we saw all that there is to see, but for assignment, we might have some extra things to consider.

```cpp
class Password{
public:
    explicit Password(std::string pwd) //pass by value and construction
    :text(std::move(pwd))
    {
    }
    void changeTo(std::string newPwd)
    {
        text = std::move(newPwd); //assignment
    }
private:
    std::string text; //let's pretend we didn't store a password as plain text
};

std::string initPwd("lets say this is a real password!");
Password p(initPwd);
std::string newPassword = "Another totally real password!";
p.changeTo(newPassword);
```

the call to changeValue requires to create a copy of the std::string (copy constructor for lvalue), and before the assignment, we clear the old text value (assume we've checked they aren't the same), and now we use the move constructor to store the new password, and de-allocate the memory for it. So two operations in total (allocation for the copy and de-allocation for the assignment).

had we used the overloading approached, we could probably avoid the costs, because the old password already has enough memory allocated. we don't need to pay for allocation when creating parameter (we take by reference), and if the memory we have is sufficient, we don't need to request anything more

```cpp
class Password{
public:
    //...
    void changeTo(const std::string & newPwd) //the overload for lvalue
    {
        text = newPwd; //we can re-use the text.capacity.
    }
private:
    //..
};

std::string initPwd("lets say this is a real password!");
Password p(initPwd);
std::string newPassword = "Another totally real password!";
p.changeTo(newPassword);
```

of course, if we had tried to request a really long new password, there would be no difference in costs. at both cases we would require to allocate extra memory and release the older memory. but when using pass-by-value, this can't be avoided. The costs of assignment based copying depends on the data stored in the participating objects, some library classes (std::string, std::vector) know how to optimize when possible.
this usually comes into play when using lvalue arguments, for rvalue arguments, we can avoid true copies most of the time and use move semantics.

this makes the computation of which code is suitable for 'pass-by-value' even more complicated.

in general, for software that is build for speed, we shouldn't use 'pass-by-value' at all. these games and assumptions about 'one extra move won't hurt in the long run if we get less code bloat' don't hold, even more if there is a chain of functions that employ this idea, one extra operation at each level can lead to several extra operations for each call, and maybe several copies that could have been avoided if we simply accepted that we might have a slightly larger binary file.

#### Slicing

unrelated to performance, passing by value is suspectable to _'the slicing problem'_, this happens when we pass a derived object by value to a function that expects a base class object. this is a c++98 issue, and part of why we moved away from pass-by-value to pass-by-reference

```cpp
class Widget{};
class SpecialWidget : public Widget{};
void processWidget(Widget w); //pass by value, base class
SpecialWidget sw;
processWidget(sw); //oops! pass by value! slicing! this will be seen as Widget, not a SpecialWidget
```

#### Things to Remember

> - For copyable, cheap-to-move parameters that are always copied, pass by value
>   may be nearly as efficient as pass by reference, it’s easier to implement, and it
>   can generate less object code.
> - Copying parameters via construction may be significantly more expensive
>   than copying them via assignment.
> - Pass by value is subject to the slicing problem, so it’s typically inappropriate
>   for base class parameter types.

</details>

### Item 42: Consider Emplacement Instead of Insertion

<details>
<summary>
Emplacement (constructing new elements on the container) tends to be faster than insertion, but not always. there are some pitfalls where emplacement can cause problems: resource managements, explicit constructors.
</summary>

when we have a container, we should sometimes use emplacement rather than insertion.

consider the following example, we have a vector of strings, and we try to add another element to it. in general, the insertion function (like _.push_back()_) are overloaded for both lvalue and rvalue.

```cpp
template <class T, Class Allocator<T>>
class vector {
    public:
    //...
    void push_back(const T& x); //lvalue
    void push_back(T && x); //rvalue
};

std::vector<std::string> vs
vs.push_back("Stacy");
```

so in our call, the compiler creates a temporary value for from the string literal when calling the function.

```cpp
vs.push_back(std::string("Stacy"));
```

but if we drill deeper, we have an extra constructor and destructor calls which we didn't notice.

1. we create a temporary std::string value from the string literal with the approximate **constructor**, this objects is an unnamed rvalue, so we'll call it 'temp'.
2. we pass this 'temp' object into the vector's _.push_back()_ function as an rvalue, and now we push it inside the vector with a **move constructor**.
3. being moved from, and temporary rvalue object, we call the **destructor** for 'temp'.

we are paying for an extra constructor and destructor call that we don't need, we could avoid this by using the emplacement functions! of course, this involves perfect forwarding, we can use _.emplace_back()_ or _.emplace_front()_ for every container that supports _.push_back()_ and _.push_front()_. only _std::forward_list_ has _.emplace_after()_ (matching _.insert_after()_) and _std::array_ doesn't have emplacements at all.

another gain from the emplacement functions is that we can use all sorts of arguments, not just complete objects! we just need to avoid braced uniform initializations

```cpp
vs.emplace_back("Stacy");
vs.emplace_back(50,'x'); // call the std::string constructor that makes 50 'x' characters
```

insertion functions take complete objects, while emplacement functions construct the objects in place and take the arguments needs for the constructor calls. this doesn't matter much with lvalue arguments, but makes a difference for rvalue objects.

```cpp
std::string s = "Queen of Disco";
vs.push_back(s); //copy constructor called.
vs.emplace_back(s); //also, copy constructor called.
```

#### Why Not Always Use Emplacement?

there are case where insertion operations are better than emplacements, this is implementation dependent, and also depend on the containers,the arguments, the safety guarantees and even the location where the insertion/emplacement takes place. benchmarking is the only way to know for sure.

but if we want some heuristic, we can says that if the following conditions are satisfied, emplacement is probably the faster approach. alternatively, each condition represents a situation where insertion has an advantage over emplacement.

> - **The value being added is constructed into the container, not assigned**

in the opening example, we had a value being constructed, but if we change it a bit, the emplacement operation loses it's edge.

```cpp
std::vector<std::string> vs;
vs.emplace(vs.begin(),"Stacy");
```

most implementation won't construct the object immediately on the memory vs\[0], they would rather create a temporary object and than move from it. which means the primary advantage is gone.

```cpp
vs[0]= std::move(std::string("Stacy"));
```

containers can be either construction based or assignment based, most containers are node-based and use constructions, with notable exceptions of _std::vector_, _std::deque_ and _std::string_. we can use _.emplace_back()_ to force a construction on most non-node containers, and _.emplace_front()_ for the _std::deque_ container.

> - **The argument type(s) being passed differ from the type held by the container**

as before, the main advantage of emplacement is not using temporary objects, but if the arguments is already of the correct type, there won't be a needless object creation.

> - **The Container is unlikely to reject the new value as a duplicate**

if the container permits duplicates, emplacement is fine, but if we require unique values, the implementation might create a temporary object to compare with the values and then move from it if needed. thus nullifying the main edge over insertion once again.

if we return to our example, we can see why it likely that emplacement will be better than insertion:

1. we construct the value, as we are using _emplace_back()_.
2. the argument type is `const char *`, rather than _std::string_
3. there is no duplicate checking in the container.

```cpp
vs.emplace_back("Stacy");
```

#### Resource Managements

if we want to add a new _std::shared_ptr_ to the list that has a custom deleter, we would prefer to create it with the utility function of _std::make_shared_. but since we can't add the custom deleter, we must use the naked new operator.

```cpp
std::list<std::shared_ptr<Widget>> widget_ptrs;
void killWidget(Widget *pWidget); //custom deleter
widget_ptrs.push_back(std::shared_ptr<Widget>(new Widget, killWidget));
//widget_ptrs.push_back({new Widget, killWidget}); //same
```

1. in either type of call, we would create a temporary _std::shared_ptr_ to hold the result of the 'new' operator.
2. _.push_back()_ takes a temporary rvalue by reference, but it fails to allocate memory for it.
3. the exception is propagated upwards, and the temporary _std::shared_ptr_ is destroyed, along with the memory it requested. no memory leaks were added.

could the temporary object creation be avoided with using _emplace_back()_?

```cpp
std::list<std::shared_ptr<Widget>> widget_ptrs;
void killWidget(Widget *pWidget); //custom deleter
widget_ptrs.emplace_back(new Widget, killWidget);
```

1. the _new Widget_ is perfect forwarded, the container tries to allocate a memory for the node where it would construct the _std::shared_ptr_ object, but fails.
2. as the exception is propagated upwards, no one has custody of the results of the call to the _new Widget_, so the naked pointer is lost and we have a memory leak.

the same problem could happen with _std::unique_ptr_ with a custom deleter. the RAII classes are only useful if they take control over the resources immediately, any time wasted between the creation of the resource and the acquisition is putting the concept at risk. using perfect forwarding means we hold construction until we are sure we can construct the object at the optimal place, but this means we increase the window of time where resource are un-managed.

the correct way to write the above code is by first creating the _std::shared_ptr_ and then moving it into the container. but this means we don't have an advantage for emplacement over insertion.

```cpp
std::shared_ptr<Widget> spw (new Widget, killWidget);
widget_ptrs.push_back(std::move(spw));
widget_ptrs.emplace_back(std::move(spw));
```

#### Explicit Constructors

another case that emplacement operations have issues is for explicit constructors. if we have a container of _std::regex_ expressions, it's possible call an emplacement function with _nullptr_ as an argument, but not insertion.

```cpp
std::vector<std::regex> regex_expressions;
//std::regex r = nullptr; //this doesn't work
//regex_expressions.push_back(nullptr);//this doesn't work
regex_expressions.emplace_back(nullptr); //so why does this work? with no warnings!?
```

the reason is that _std::regex_ has a constructor for characters strings, but it's only an explicit constructor.

```cpp
std::regex upperCaseWord("[A-Z]+"); //one or more characters from A to Z.
//std::regex r = nullptr; //this doesn't work
//regex_expressions.push_back(nullptr);//this doesn't work
std::regex r(nullptr); //this compiles, explicit conversion.
```

in both cases that don't compile, we are requesting an implicit conversion from the `const char*` pointer, which is not allowed. but in the emplacement function, we passing arguments to a constructor, so the explicit constructor is called.

the fact that the code compiles doesn't mean it runs, if we're lucky, it wll crash at runtime each time, if we're unlucky, we might need to hunt for this bug for many hours.

as a tangent. the two forms of initializations are different, the first form is called 'copy initialization',and the second form is 'direct initialization'. copy initialization isn't allowed to call explicit constructors, only implicit.

```cpp
std::regex r1 = nullptr; //error! doesn't compile!
std::regex r2(nullptr); //compiles!
```

the emplacement functions use direct initializations, so they call explicit constructors

#### Things to Remember

> - In principle, emplacement functions should sometimes be more efficient than their insertion counterparts, and they should never be less efficient.
> - In practice, they’re most likely to be faster when
>   1. the value being added is constructed into the container, not assigned.
>   2. the argument type(s) passed differ from the type held by the container.
>   3. the container won’t reject the value being added due to it being a duplicate.
> - Emplacement functions may perform type conversions that would be rejected
>   by insertion functions.

</details>
