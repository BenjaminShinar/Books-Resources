<!--
ignore these words in spell check for this file
// cSpell:ignore Eerd adjacens plecto Bilger
-->
[Main](README.md)
Lightning Talks


### Complecting Made Easy - Tony Van Eerd

[Complecting Made Easy](https://youtu.be/jSIMzirLJKE)

'simple made easy' is hard.

the word "easy" comes from the latin "adjacens", which means "close at hand", things that are easy are "close to us". 

> complect
> - "plecto" (latin) = fold/ braid
> - simple - one fold
> - complex - many folds
> - Complecting - interweaving

complecting is easy, make code intertwined:(bad code!)

- add global variable
  - one could change the global variables
- shared pointer (raw pointer as well)
  - "a shared pointer is as good as a global variable"
- a class that has access to both of them. (MODEL)
- base classes
  - static variable in the base class!
- grab code that you know and copy it!
- members functions are easier than free function,
  - "classes are made of Velcro"
  - naming files is hard... (utils is a bad name)

## Dashboards to the Rescue - Matthias Bilger

[Dashboards to the Rescue](https://youtu.be/eOMqO0OKsCw)

everybody loves cats and dashboards. nice looking, no knowledge required, easy access to KPI (key performance index)

if we remove the assert statements, we can make the test pass! we can also count the assert statements, or add pointless tests to ramp up the KPI of line coverage or code coverage.

> - "easy measurable does not mean useful"
> - making a metric a target is bad
> - dashboards will not save us.

## Universal Function Call Syntax in C++20 - Devon Richards
[Universal Function Call Syntax in C++20](https://youtu.be/uT1ZJHM8DkE)

calling a free function and a member function with the same syntax, like `std::size` on containers, which uses the member function `size` from the container.

but it also lets us write function from left to right, like the pipe operator for ranges. he has a macro that allows it in compiler explorer, and the performance hit is really small, the macro does a hack to allow templates somehow.

##

[Main](README.md)