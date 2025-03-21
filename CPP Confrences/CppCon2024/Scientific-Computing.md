<!--
// cSpell:ignore Vectorizing kmph Electronvolt Kathir Farghani Alfraganus
-->

<link rel="stylesheet" type="text/css" href="../../markdown-style.css">

## Scientific Computing

<summary>
10 Talks about Scientific Computing C++.
</summary>

- [ ] - A New Dragon In The Den: Fast Conversion From Floating-Point Numbers - Cassio Neri
- [x] - Application Of C++ In Computational Cancer Modeling - Ruibo Zhang
- [ ] - Bridging The Gap: Writing Portable Programs For Cpu And Gpu - Thomas Mejstrik
- [ ] - Data Is All You Need For Fusion - Manya Bansal
- [ ] - High-Performance Numerical Integration In The Age Of C++26 - Vincent Reverdy
- [ ] - High-Performance, Parallel Computer Algebra In C++ - David Tran
- [x] - Improving Our Safety With A Quantities And Units Library - Mateusz Pusz
- [ ] - Linear Algebra With The Eigen C++ Library - Daniel Hanson
- [ ] - To Int Or To Uint, This Is The Question - Alex Dathskovsky
- [ ] - Vectorizing A Cfd Code With `Std::Simd` Supplemented By (Almost) Transparent Loading And Storing - Olaf Krzikalla

---

### Improving Our Safety With A Quantities And Units Library - Mateusz Pusz

<details>
<summary>
P2981: Improving our safety with a physical quantities and units library. summary of how the library behaves and improves the code safety and quality.
</summary>

[Improving Our Safety With A Quantities And Units Library](https://youtu.be/pPSdmrmMdjY?si=XqQIA63B6O9eDPJn), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Improving_Our_Safety_With_a_Quantities_and_Units_Library.pdf), [event](https://cppcon2024.sched.com/event/1gZed/improving-our-safety-with-a-quantities-and-units-library), [P2981 proposal](https://wg21.link/p2981), [mp-units repository](https://github.com/mpusz/mp-units).

C++ safety, communicating units between processes, making errors can lead to disasters. it's important to get this right.

> Affected Industries
>
> - Aerospace
> - Autonomous cars
> - Embedded industries
> - Manufacturing
> - Maritime industry
> - Freight transport
> - Military
> - Astronomy
> - 3D design
> - Robotics
> - Audio
> - Medical devices
> - National laboratories
> - Scientific institutions and universities
> - All kinds of navigation and charting
> - GUI frameworks
> - Finance (including HFT)

It's not enough to just invest in training, it's still up to human skill, and many of the engineers writing C++ safety critical code aren't professional programmers, they are domain experts.

the goal of the library is to generate compile time errors - easy to understand, debug, and fix.

#### Typical Production Issues

examples of potential problems from real code bases.

1. The proliferation of `double` - same type every where
2. The proliferation of magic numbers - values that only make sense to domain experts.
3. The proliferation of conversion macros - using macros, redefining the same name again and again in different ways.
4. Lack of consistency - APIs that are easy to err with.

#### MP-Units & Standardization

The C++20/23 <cpp>mp-units</cpp> library, already available in github, conan package manager and compiler explorer.

> Goals:
>
> - Compile-time safety:
>   - correct handling of physical quantities, units, and numerical values
> - Performance:
>   - as fast or even faster than working with fundamental types
>   - no runtime overhead
>   - no space size overhead
> - Great user experience:
>   - optimized for readable compilation errors and great debugging experience
>   - easy to use and flexible
> - Scope:
>   - any unit's magnitude (huge, small, floating-point)
>   - systems of quantities
>   - systems of units
>   - the affine space
>   - highly adjustable text-output formatting
>   - scalar, vector, and tensor quantities
>   - natural units systems

#### A Taste Of Quantities And Units Library

moving from a manual implementation to using the library. using types with defined behaviors.

```cpp
// before
constexpr auto M_PER_KM = 1000.;
constexpr auto CM_PER_MI = 2.54 * 12. * 5280;
constexpr auto M_PER_MI = CM_PER_MI / 100.;
constexpr auto S_PER_H = 3600.;
constexpr auto MPS_PER_KMPH = M_PER_KM / S_PER_H;
constexpr auto MPS_PER_MPH = M_PER_MI / S_PER_H;

const double distance_m = 30.;
const double speed_mph = 25.;
const double speed_mps = speed_mph * MPS_PER_MPH;
const double time_to_goal_s = distance_m / speed_mps;
std::println("TTG: {:.6} s", time_to_goal_s);

// after
const quantity distance = 30. * m;
const quantity speed = 25. * mi / h;
const quantity time_to_goal = (distance / speed).in(s);
std::println("TTG: {::N[.6]}", time_to_goal_s);
```

if we try to multiply the distance and speed instead of dividing them, we would get an error. with the library, rather than denoting the unit (seconds, meters, kilometers per hour) in the names of the parameters, they are encoded into the type, which eliminates bugs.

```cpp
// before
double time_to_goal_s(double distance_m, double speed_kmph)
{
  return distance_m / (speed_kmph * MPS_PER_KMPH);
}

// after
quantity<s> time_to_goal(quantity<m> distance quantity<km/h> speed)
{
  return distance / speed;
}
```

if we pass the wrong variable (distance in km), then we get an error. we also cant mix up the argument order. it's much harder to ship bad code, since we get compile time checks through the type system.\
if we look at the compiled assembly code, it's basically the same.\
if we to have the input or output at different units, we can use generic programming and <cpp>concepts</cpp>.

```cpp
QuantityOf<isq::time> auto time_to_goal (QuantityOf<isq::length> auto distance, QuantityOf<isq::speed> auto speed)
{
  return distance / speed;
}

const quantity distance_to_turn = 400. * ft;
const quantity car_speed = 40. * mi / h;
const quantity ttg = time_to_goal(distance_to_turn, car_speed);
std::println("Turn right after {::N[.1]}", ttg.in(s));
```

we can add together values (minutes and seconds, distances) without manually scaling them.

#### Safety Features

safe unit conversions. magnitudes are known at compile time.\
<cpp>std::chrono</cpp> is still missing some units that are defined in the standard, which are either too large or too small to define with 64 bits. such as "electronvolt" ($1 _eV = 1.602176634 \times 10^{-19}J$) or Dalton ($1 Da = 1.660539040(20) \times 10^{-27} Kg$) or some units that require conversions with irrational numbers.\
we can define our own prefixes, and apply them for multiple units (kilogram, kilometer, etc...) since they can collide with either namespaces, they need to be opt-in into. there are also definitions to non-standard units (yards, miles, etc...).

Preventing truncation of data.

> Conversion of a quantity with the integral representation type to one with a unit of a lower resolution is truncating.

by default, we define things as integral, and we don't allow  truncating them down.  if we want quantities with fractions, we must define them as floating points, at our own risk.

```cpp
quantity q1 = 5 * m; 
std::cout << q1.in(km) << '\n'; // Compile-time error
quantity<si::kilo<si::metre>, int> q2 = q1; // Compile-time error

quantity q1f = 5. * m; // source quantity uses 'double' as a representation type
std::cout << q1f.in(km) << '\n';
quantity<si::kilo<si::metre>> q2f = q1f;
```

#### Tracing Columbus Route To The Bahamas

the story of Columbus and repressing it in code, going over the things he knew and what sort of units he used. Columbus used roman units, but relied on calculation made with persian units, leading to differences in what the length of a mile is.

```cpp
// length of degree of latitude estimation by medieval Persian geographer
// Abu al Abbas Ahmad ibn Muhammad ibn Kathir al-Farghani (a.k.a. Alfraganus)
// (degree of longitude at the equator should be roughly equivalent)
template<UnitOf<isq::length> auto Mile>
struct estimated_degree final : named_unit<"deg", mag_ratio<5667, 100> * Mile> {};

// roman units
inline constexpr struct roman_foot final : named_unit<"ft_r", mag<296> * si::milli<si::metre>> {} roman_foot;
inline constexpr struct roman_pace final : named_unit<"pace_r", mag<5> * roman_foot> {} roman_pace;
inline constexpr struct roman_mile final : named_unit<"mi_r", mag<1000> * roman_pace> {} roman_mile;

// used in Persia
// extended the Roman mile to fit an astronomical approximation of 1 minute of an arc of latitude
inline constexpr struct arabic_mile final : named_unit<"mi_a", mag<2163> * si::metre> {} arabic_mile;

// 1 minute of arc along the Earth's equator
inline constexpr struct geographical_mile final : named_unit<"mi_g", mag_ratio<18'553, 10> * si::metre> {} geographical_mile;

inline constexpr auto Columbus_degree = estimated_degree<roman_mile>{};
inline constexpr auto Alfraganus_degree = estimated_degree<arabic_mile>{};
inline constexpr struct equator_degree final : named_unit<"deg", mag<60> * geographical_mile> {} equator_degree;

template<Quantity Q1, Quantity Q2>
  requires std::invocable<std::minus<>, Q1, Q2>
quantity<percent> error(const Q1& approximate, const Q2& exact)
{
  return abs(approximate - exact) / exact;
}

std::cout << "Roman mile: " << (1. * roman_mile).in(si::metre) << "\n";
std::cout << "Arabic mile: " << (1. * arabic_mile).in(si::metre) << "\n";
std::cout << "Mile error: " << error(1. * roman_mile, 1. * arabic_mile) << "\n";

const quantity Columbus_equator_length = 360. * Columbus_degree;
const quantity Alfraganus_equator_length = 360. * Alfraganus_degree;
const quantity equator_length = 360. * equator_degree;

std::cout << "Columbus equator length: " << Columbus_equator_length.in(nmi) << "\n";
std::cout << "Alfraganus equator length: " << Alfraganus_equator_length.in(nmi) << "\n";
std::cout << "Equator length: " << equator_length.in(nmi) << "\n";
std::cout << "Equator error: " << error(Columbus_equator_length, equator_length) << "\n";

const quantity Columbus_distance = 68. * Columbus_degree;
const quantity Tenerife_Bahamas_distance = 5'982. * km;
const quantity Tenerife_Japan_distance = 10'600. * nmi;

std::cout << "Columbus distance: " << Columbus_distance.in(nmi) << "\n";
std::cout << "Tenerife-Japan distance: " << Tenerife_Japan_distance.in(nmi) << "\n";
std::cout << "Distance error: " << error(Columbus_distance, Tenerife_Japan_distance) << "\n";
std::cout << "Tenerife-Bahamas distance: " << Tenerife_Bahamas_distance.in(nmi) << "\n";
```

> Thanks to the usage of quantities and units library a developer has to focus only on a program logic and does not have to carefully verify every unit conversion and quantity arithmetics.

#### More Issues

> Implementing a physical quantities and units library is much
harder than it may initially appear.

explicit constructors, everywhere, always provide the unit and the value. interacting with legacy code that still use primitives. requiring more than one dimension for a quantity. length is one thing, but height, width, distance and wavelengths aren't the same thing, even if they are all measured with the same units. our type system must be able to tell them apart and prevent confusion. this is achieved by defining <cpp>quantity_spec</cpp> following the ISO defintions. we can be as safe as we wish, depending on how exact we want to be.

type quantities

> `res = 1 * Hz + 1 * Bq + 1 * Bd;`
>
> - Hz (hertz) - unit of frequency
> - Bq (becquerel) - unit of activity
> - Bd (baud) - unit of modulation rate

running the calculation in different languages:

```cpp
// boost
using namespace boost::units::si;
std::cout << 1 * hertz + 1 * becquerel << '\n'; // 2 Hz
std::cout << 1 * becquerel + 1 * hertz << '\n'; // 2 Hz

// other units library
using namespace units::literals;
std::cout << 1_Hz + 1_Bq << '\n'; // 2 s^-1
```

with python

```python
print(1 * ureg.hertz + 1 * ureg.becquerel + 1 * ureg.baud) # 3.0 hertz
print(1 * ureg.becquerel + 1 * ureg.hertz + 1 * ureg.baud) # 3.0 becquerel
```

and with java - we get a compilation error.

```java
System.out.println(Quantities.getQuantity(1, Units.HERTZ)
  .add(Quantities.getQuantity(1, Units.BECQUEREL)));
```

even though they all a qualities of similar thing (dimension $T^{-1}$), they aren't comparable since they don't belong to the same domain, and shouldn't be mixed. in the <cpp>mp-units</cpp> library, their is a hierarchy tree for quantities that belong to the same kind: the <cpp>kind_of\<QS></cpp> modifier. so even though hertz and becquerel are both the same dimension, they aren't te same kind (frequency vs activity).

```cpp
static_assert(get_kind(isq::width) == get_kind(isq::height));
static_assert(get_kind(isq::width) == kind_of<isq::length>);
static_assert(implicitly_convertible(kind_of<isq::length>, isq::width));

namespace mp_units::si {
  // base quantities
  inline constexpr struct second final : named_unit<"s", kind_of<isq::time>> {} second;
  inline constexpr struct metre final : named_unit<"m", kind_of<isq::length>> {} metre;
  inline constexpr struct gram final : named_unit<"g", kind_of<isq::mass>> {} gram;
  inline constexpr auto kilogram = kilo<gram>;
  inline constexpr struct ampere final : named_unit<"A", kind_of<isq::electric_current>> {} ampere;
  inline constexpr struct kelvin final : named_unit<"K", kind_of<isq::thermodynamic_temperature>> {} kelvin;
  inline constexpr struct mole final : named_unit<"mol", kind_of<isq::amount_of_substance>> {} mole;
  inline constexpr struct candela final : named_unit<"cd", kind_of<isq::luminous_intensity>> {} candela;

  // derived quantities
  inline constexpr struct radian final : named_unit<"rad", metre / metre, kind_of<isq::angular_measure>> {} radian;
  inline constexpr struct steradian final : named_unit<"sr", square(metre) / square(metre), kind_of<isq::solid_angular_measure>> {} steradian;
  inline constexpr struct hertz final : named_unit<"Hz", inverse(second), kind_of<isq::frequency>> {} hertz;
  inline constexpr struct becquerel final : named_unit<"Bq", inverse(second), kind_of<isq::activity>> {} becquerel;
  inline constexpr struct newton final : named_unit<"N", kilogram * metre / square(second)> {} newton;
  inline constexpr struct pascal final : named_unit<"Pa", newton / square(metre)> {} pascal;
  inline constexpr struct joule final : named_unit<"J", newton * metre> {} joule;
  inline constexpr struct watt final : named_unit<"W", joule / second> {} watt;
  inline constexpr struct coulomb final : named_unit<"C", ampere * second> {} coulomb;
}
```

so we can get the same compile time error as we saw in the java code, we can't do operations on types where it doesn't make sense.\
the library also has affine spaces: a point (position) and displacement vector (difference between two points). the affine space limits the allowed operations.
</details>

### Application Of C++ In Computational Cancer Modeling - Ruibo Zhang

<details>
<summary>
Using C++ to simulate cancer.
</summary>

[Application Of C++ In Computational Cancer Modeling](https://youtu.be/_SDySGM_gJ8?si=shtZkbvwMDKjpTyZ), [slides](https://github.com/CppCon/CppCon2024/blob/main/Presentations/Application_Of_Computational_Cancer_Modeling.pdf), [event](https://cppcon2024.sched.com/event/1gZiB/application-of-c-in-computational-cancer-modeling).

> Main Topic: use C++ to simulate the process of cancer initiation
>
> - The mathematical model and simulation study
>   - Generate a single tumor (A single step of evolution)
>   - Generate multiple tumors (Tasked Based Concurrency)
>   - Obtain statistical properties of the tumors (Parallel STL algorithms)
> - Eigen (Array Class)
>   - <cpp>Eigen</cpp> is a C++ template library for linear algebra: matrices, vectors, numerical solvers, and related algorithms.
> - Modern C++:
>   - <cpp>random</cpp>: Pseudo-random number generation
>   - <cpp>future</cpp>: Task-Based Concurrency
>   - <cpp>numeric</cpp>: Parallel versions of certain STL algorithms

defining cancer, uncontrolled division of abnormal ells, we want a mathematical model to understand the evolution of cancer and predict the widnow of opportunity for screening. we define our model as having cells of different types, a cell can either alter it's type or divide into two cells of the same type. this constitutes a markov chain. the event happens on a random schedule - mutation rate and growth rate.

```cpp
#include <random>
std::mt19937_64 rnd_generator;
std::exponential_distribution<> exp{rate};
double time = exp(rnd_generator);
```

the inputs to our model are the starting population, the rates and the possible changes, the output is the disribution of cells at different timepoints. we use a dynamic two-dimension array from the <cpp>eigen</cpp> library, and we have a transition matrix between cell states.

more code examples, doing matrix stuff, column-wise operations and so on. then doing thing in parallel using <cpp>std::future</cpp> and launching the simulation in another thread.

</details>
