<!--
ignore these words in spell check for this file
// cSpell:ignore unpromoted Voxel Scanlines Luma
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

[main](README.md)

## Math & Numbers

### Principia Mathematica - The Foundations of Arithmetic in C++ - Lisa Lippincott

<details>
<summary>
defining mathematics in c++ terms, creating operators, integer promotion and conversions.
</summary>

[Principia Mathematica - The Foundations of Arithmetic in C++](https://youtu.be/0TDBna3PWgY), [slides](https://github.com/CppCon/CppCon2022/blob/main/Presentations/Principia-Mathematica.pdf).

the original "Principia Mathematica" builds up mathematics from the basic units, all in a formal and structured way.

the foundations of (arithmetic in C++) vs (the foundations of arithmetic) in C++)

```
result_type function_name (parameter list)
interface
{
  //pre conditions
  implementation;
  //post conditions
}
```

Basic principals:

- Stability - over certain periods, an object state, and therefore its value, remains stable.
- Substitutability - At certain times, two different objects of the same type have interchangeable values.
- Repeatability - an operation may be repeated by a sufficiently similar operation, producing similar results.

usages:

- right of stability - we pass the right of stability, assignment operators, anything that changes. when we create a new object (function result), it also passes the right of stability.
- immunity from instability - the operation assures us that there will be no change to the object.
- claim substitutable - the result of the "function" refers to the original object passed to it. (assignment, comparisons)
- discernable inputs and outputs

an assignment operator "a=b"

- claim_right a;
- claim_immunity b;
- discern b;
- implementation;
- claim_substitutable(&a, &result);
- claim_right result;
- discern result;
- claim_substitutable(result, b);

integer types have a width (range) and signed/unsigned.

```cpp
class integer_kind {
  constexpr bool is_signed() const;
  constexpr bit_size_t width() const;
};

inline constexpr bool operator <(integer_kind a,integer_kind b) {
  return a.is_signed() <= b.is_signed() && a.width < b.width();
}

template<class To, class From>
requires (integer_kind_of<To> >= integer_kind_of<From>)
To convert(const From&);
```

all values of five (int, long, unsigned int, signed long long) theoretically map to the same "five". but converting can also be narrowed, and there is convert_modular, which only considers lower bits.

now we can extended our assignment model, rather than just using int,we can have convert modular and convert narrowing. for operators that combine assignment and modifying, and we build the interface module again. then we have to work on the postfix and prefix.

we don't like to do math on some of the integer types, so we are promoting them into int or unsigned int. we can express that in a template using _std::conditional_t_, but there are some issues with bit fields and _std::underlying_type_. we can use other methods to find the promoted type. we write an interface for unpromoted integers.

(a whole lot of make-believe functions to convert integers from one kind to other). a lot of narrowing, widening, and whatnot.

we make an exception of the bit shifting operators, and we look at special cases of dividing by zero, or modulo zero, and when the result of a division is of a different type. (we can't divide the minimal_negative_number by minus-one, because there isn't a positive corrsponding number). tons of undefined behavior stuff.

(creating more new types and templates), then doing something to create the plus operator. (everything is weird)

if we have written all this code, we can now prove theorems, like how addition_is_commentative (a+b == b+a). we say that with these pre-conditions and the implementation, then the postconditions hold.

we end up with executable code that is both code and formal proof.

</details>

### What Is an Image? - Cpp Computer Graphics Tutorial, (GPU, GUI, 2D Graphics and Pixels Explained) - Will Rosecrans

<details>
<summary>
Brief overview of Images and related pitfalls.
</summary>

[What Is an Image? - Cpp Computer Graphics Tutorial, (GPU, GUI, 2D Graphics and Pixels Explained)](https://youtu.be/zi57OkPwzbk), [slides](https://docs.google.com/presentation/d/1zF716ULMyQGVJYmyIujDn-WxcmVLsb3Njzs_ZL5fOwM/edit).

a lot of how we work with images was defined years ago to support early monitors, even if things no longer make sense with modern hardware.

luminescence, colors (red, green, blue)

> A digital image is a contiguous 2D array of little squares called pixels with 3x8-bit RGB colors, stored in the the memory of a computer.\
> Pixels are accessed by iterating a pointer through the image in order, left to right, top to bottom.

```cpp
class Pixel {
  uint8_t r, g, b;
};

class Image {
  uint8_t *pixels;
  int size;
};
```

some history stuff:

- vector graphics - image doesn't exist in memory, only instruction to draw it.
- The early macintosh had black and white, 1-bit pixels, no color.
- The nintendo used tiled graphics, no image in memory, can't draw arbitrary image, used palette colors (pixels didn't have colors directly).
- The amiga had different video modes in the same screen.

comparing our defintion to those earlier ways of showing 'images' makes the defintion clearer.

there is a paper from 1995 named: "A Pixel Is _Not_ A Little Square, A Pixel Is _Not_ A Little Square, A Pixel Is _Not_ A Little Square! (And a Voxel is _Not_ a Little Cube)".

alpha channel (opacity/transparency) - RBGA.

```cpp
struct Pixel {
  uint8_t r, g, b;
};

struct Pixel2 {
  uint8_t r, g, b, a;
};
```

but because of alignment stuff, using the larger struct actually performs better, since the data is better aligned in memory lines. Different libraries such as _Qt QImage_ and _Open Image IO_, prefer using Scanlines or tiles, and the aren't always compatible with one another, there can be buffering and padding problems (beware of buffer overflow).

#### Colors

Colorspaces, Gamut and Primaries, linear RGB and non linear sRGB. transfer functions for adding "colors" together. a lot of variability between libraries. mathematically necessary but impossible "Imaginary images". old televisions had narrower range of black and white (from 16 to 235, rather than 0 to 255), and even today we use more bandwidth for brightness than we do for colors. the UV plane - store the colors in half the resolution of the brightness.

```cpp
class YUVImage {
  uint8_t *Luma;
  int width, height;
  uint8_t *U, *V;
  int uv_width() {
    return width/2;
  }
};
```

this means iterating over pixels is non-trivial.

#### GPU

graphical processing unit. many different formats, which the GPU might support only some of them, so the software needs to adapt to the hardware, and we want good cache locality. which is hard because images are stored in continuous memory, which means the pixel directly above any given pixel isn't at the same memory line.

GPUs sometimes have a 16bit floating point data type, but it's not defined in c++ as a native type. A lot of things are only known at runtime, and codecs change all the time.\
Tiling - storing pixels "out-of-order", improves locality. MIP maps - storing scaled down layers (copies) of an image.

video support in the GPU, avoiding round trips between the CPU and GPU, computer shaders (which aren't c++).

#### C++ Language Support

we don't know what _std::image_ will be, but it won't please everyone, and we won't get everything at first.

so our updated defintion of an image is something like this:

> " ~~A~~ (arbitrarily many) ~~contiguous 2D~~ (N-Dimensional) array(s) of ~~little squares called~~ pixels with ~~3x8 bit~~ (Nx??? bit int or float) (possibly) RGB colors (whatever that means) (or arbitrary data), stored in a computer (or a GPU or something's) memory.\
> Access patterns may be complex, indirect, remote, interpolated, planar, tiled, …"

</details>

### Quantifying Dinosaur Pee - Expressing Probabilities as Floating-Point Values in C++ - John Lakos

<details>
<summary>
The problems of using floating point types as probabilities
</summary>

[Quantifying Dinosaur Pee - Expressing Probabilities as Floating-Point Values in C++](https://youtu.be/emKOZldM22w), [slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/Dino-Pee.220916-FINAL3_submitted_p4.pdf)

boundary condition in testing.

#### The Birthday Problem

> The birthday problem:
> What is the probability that two or more people in a room of _n_ people have the same birthday?\
> (or the revised form)\
> What is the _minimum number_ of people we would need in a room for the probability that "at least two of them have the same birthday" is greater than 50%?

three kinds of boundaries:

1. defined by the interface
2. create by the implementation
3. imposed by the platform

we make some assumptions to make the problem easier (uniformity, only day matters, no leap years).

the birthday function provides the probability value in the range of Zero to One as a function of the number of people in the room. we need to define the function.

`double sameBirthday(int numPeople)`

for zero or one people $p=0.0$, for 366 or more people $p=1.0$, these are the boundaries by the logical defintion. we create assertion on them. checking `INT_MAX` is platform imposition.

if we do a naive check, we find the assertion starts failing at 184 people, it returns a probability that is indistinguishable from 1.0.

```cpp
double sameBirthday(int numPeople) {
  assert(0 <= numPeople>); // pre condition
  if (365 < numPeople>) {
    return 1.0; // early exit
  }
  double probability = 0.0;
  for (int i = 1; i < numPeople; ++i) {
    probability += (1.0 - probability) * (i/365.0);
  }
  return probability;
}
```

the case of 184 people is another platform boundary, because of us using a double to represent "almost one".

the first bit is a sign-bit, then 11 bits of exponents, and 52 bits of the mantissa. the closest we can get to 1 is $1-2^{-52}$. but it's much less precises than $2^{-1074}$ which is the closest positive number we can get to zero.

| number      | math        | sign bit | exponents   | mantissa                                             |
| ----------- | ----------- | -------- | ----------- | ---------------------------------------------------- |
| zero        | $0.0$       | 0        | 00000000000 | 0000000000000000000000000000000000000000000000000000 |
| almost zero | $2^{-1074}$ | 0        | 00000000000 | 0000000000000000000000000000000000000000000000000001 |
| almost one  | $1-2^{-52}$ | 0        | 01111111110 | 1111111111111111111111111111111111111111111111111111 |
| one         | $1.0$       | 0        | 01111111111 | 0000000000000000000000000000000000000000000000000000 |

can we fix this? we need to rework our interface and contract. lets reframe the question as "uniqueBirthday", the probability of no two people sharing the same birthday.

$P(N) = \frac{365}{365}\times\frac{364}{365}\times\frac{363}{365}\times... = \frac{365!}{(365-N)! \times 365^{N}}$

then we use _Stirling's approximation_ to see the range of the highest option (365 people) and it's only about $\cong 1.45*10^{-157}$, which is inside the resolution we have.

```cpp
double uniqueBirthday(int numPeople) {
  assert(0 <= numPeople>); // pre condition
  if (365 < numPeople>) {
    return 0.0; // early exit
  }
  double probability = 0.0;
  for (int i = 364; i < (366 - numPeople); --i) {
    probability *= i/365.0;
  }
  return probability;
}
```

now we pass our assertions!

we did some quick test at the interface boundaries, we found platform boundaries, and we had to change the contract.

#### Quantifying Dinosaur Pee

> How much of the world's water is dinosaur pee?

lets start with assumptions to make things simple, and now we have four questions

> 1. What proportion of the Earth's water is dinosaur
>    pee today?
> 2. What is the probability that ALL molecules of
>    water on earth were once inside a dinosaur?
> 3. Suppose the expected amount of untainted
>    water turns out to be on the order of a single
>    molecule; what would then be the probability
>    that ALL the molecules were tainted?
> 4. If I have an 8-ounce (236.5 ml) glass of water
>    today, what is the probability that it has no
>    dinosaur pee in it?

or in simpler form:

> 1. What fraction of all water is pee?
> 2. What’s the chance every molecule is pee?
> 3. What would _ans. #2_ be if _ans. #1_ were "all but 1 molecule"?
> 4. What’s the chance my cup of water is pure?

(some math I don't understand, showing another way of doing the match with combinatorics), a geometric sum formula
$\sum\limits_{n=0}^{x-1}{\frac{100P}{Q}\times(\frac{Q-P}{Q})^n}$

eventually we get a probability larger than 1, which again means that we run into a rounding issue and that we should be framing the question as probability close to zero.

#### Floating-Point Subtraction Algorithm

- express the two number in IEEE754 format (float)
- shifting the mantissa
- subtract
- normalization
- rounding
- might need normalization again

we lose all of our bits. different accuracy of floating points types.

</details>

### Using `std::chrono` Calendar Dates for Finance in Cpp - Daniel Hanson

<details>
<summary>
The Year_Month_Day class and it's uses in finance.
</summary>

[Using `std::chrono` Calendar Dates for Finance in Cpp](https://youtu.be/iVnZGqAvEEg),
[slides](https://cppcon.b-cdn.net/wp-content/uploads/2022/09/StdChronoDates_CppCon2022_CORRECTED.pdf), [sample code](https://github.com/QuantDevHacks/CppCon-2022-C-20-Dates-in-Finance), [low level date algorithms](https://howardhinnant.github.io/date_algorithms.html).

new c++20 features, not just concepts, modules and ranges. there were also changes to the dates classes. Dates are critical for finance (bond/fixed income trading).

fixed income products involve payment on a regular basis:

- Monthly
- Quarterly
- Semiannual
- Annual

we see them as

- Coupon-paying bonds
- Mortgage and car loans
- Annuities
- Interest rate swaps (fixed/float rate payments)
- Futures and options on bonds and swaps

example:

> $1000 Face Value
>
> - 5% annual coupon paid **semiannually** over 30 years.
> - Regular coupon payment = $(.05)(1000)/2$ = $25
> - Face value returned on final coupon payment date.

the bonds has contractual dates, but when the date falls on a weekend or a holiday, there are rolled over to the next business day. we denote each date with $c_n$ with the first payment and the last payment falling on irregular dates, and they might have different values.\
time calculations are important, and they are different ways of doing them.

1. Actual/365 - number of days between two dates, divided by 365.
2. 30/360 - assume months with 30 days and years with 360 days.need to adjust for days of month (31 days in a month)\
   $$
   DayCountFactor = \frac{360*(Y_2-Y_1) + 30*(M_2-M_1)+(D_2-D_1)}{360}
   $$

a bond is traded in the market, and it's price is the settlement date, which can be any date between the issuing and maturity. the value is calculated based on discounts, which depends on market expectations on interest rates.

#### `std::chrono::year_month_day`

the <cpp>std::chrono::year_month_day</cpp> is a prominent class of the chrono library.

```cpp
import <chrono>; // modules
#include <chrono>; // headers
```

there are several ways tos construct this object.

- with the <cpp>std::chrono</cpp> objet of year, month, day
- with the <cpp>std::chrono</cpp> hardcoded months
- using the forward slash operator as as separator
- when the first parameters is known (<cpp>std::chrono::year</cpp>), then we can use intergers instead.
- many more

a _ymd_ date can be measured as an epoch - number of days since January 1,1970, this is a serial data. (excel start at 1900). dates prior to that are negative numbers.

```cpp
ymd = std::chrono::year{ 2002 } / std::chrono::November / 14;
int days_since_epoch_count =
std::chrono::sys_days(ymd).time_since_epoch().count();
```

we can get date differences by getting the days with <cpp>sys_days</cpp> and subtracting them, which returns a <cpp>duration</cpp> object.

it's possible to create invalid dates. we can check if a data is valid by calling the `.ok()` method. we can check if the date is inside a leap year with `.is_leap()`.\
however, there is no member function to get the last day of the month. so instead, we can construct it.

```cpp
year_month_day_last eom{std::chrono::year{2009} / std::chrono::April / std::chrono::last};
auto last_day = static_cast<unsigned>(eom.day());
```

then we can assign it back to the original value

#### Date Algorithms

we can get some date algorithms which are compatible with <cpp>std::chrono</cpp> and provide ways to interact with them.

```cpp
// User-defined last_day_of_the_month
unsigned last_day_of_the_month(const std::chrono::year_month_day& ymd)
{
  unsigned m = static_cast<unsigned>(ymd.month());
  std::array<unsigned, 12> normal_end_dates{ 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31 };
  return (m != 2 || !ymd.year().is_leap() ? normal_end_dates[m - 1] : 29);
}
```

we can use the addition arithmetic operator to add chrono units (only years and months) to a date, we can't use integers directly. adding can still cause invalid dates to be created. this can happen with irregular months and leap years. to add days we need to convert into a serial date and back.

#### Date Class Wrapper

wrapping the complexity into a class that handles the work for us. most of what we need is provided by the<cpp>std::chrono</cpp> object, we are mostly concerned about rolling over the date when it falls on a weekend.

```cpp
ChronoDate& ChronoDate::weekend_roll() {
  date::weekday wd{ sys_days(date_) };
  month orig_mth{ date_.month() };
  unsigned wdn{ wd.iso_encoding() }; // Mon = 1, ..., Sat = 6, Sun = 7
  if (wdn > 5)
  {
    date_ = sys_days(date_) + days(8 - wdn);
  }
  // If advance to next month, roll back; also handle roll to January
  if (orig_mth < date_.month() || (orig_mth == December && date_.month() == January))
  {
    date_ = sys_days(date_) - days(3);
  }

  reset_serial_date_();
  return *this;
}
```

#### Summary

> The inclusion of dates in C++20
>
> - Is great to have for computational finance
> - Especially fixed income/derivatives trading
> - Possible to have invalid dates
>
> Wrap year_month_day in a user-defined class
>
> - yyyy/mm/dd representation
> - Serial date representation
> - Is leap year, date valid, number of days in month
> - Accessors for year, month, day
> - Number of days between two dates
> - Add years, months, days
> - Business day rules for weekends
> - More intuitive interface
> - Handles invalid date cases
>
> We now have a user-defined date class available to use in
>
> - Day count classes
> - Yield curve classes and term structure models
> - Bond pricing
> - Interest rate derivatives pricing models

</details>

##

[Main](README.md)
