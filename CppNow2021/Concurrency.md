<!--
// cSpell:ignore simd Steagall intrinsics cstdio immintrin loadu mmask storeu permutexvar permutex2var mmsetr maskz fmadd Giannis Gonidelis asynchrony KEWB unseq Nikunj Exascale randomizer kokkos hpx lcos
 -->

[Main](README.md)

## Adventures in SIMD-Thinking - Bob Steagall

<details>
<summary>
Creating some SIMD function to do cool stuff.
</summary>

[Adventures in SIMD-Thinking](https://youtu.be/1FPobiebZLE)

> SIMD - Single instruction, multiple data

(getting high performance from running the same instruction on a register that contains more than one data point)

> Agenda
>
> - Create some usefull basis function using some SIMD (AVX-512) intrinsics.
> - Try some SIMD-style thinking to tackle a few interesting problems.
>   - Intra-register sorting.
>   - Fast linear median-of-seven filter.
>   - Fast small-kernel convolution.
> - No heavy code, but lots of pictures
>   - Thinking "vertically"

SSE/AVX registers

- SSE 2(~2000)/3(~2004)/4(~2008)
  - 8 registers, which is 128 bits/ 16 bytes / 4 floats(or int32_t)
- AVX 2 (~2013)
  - 16 regisers (256 bits)
  - allows permuting of 32-bit elements across the two 128 lanes
  - gather primitives
- AVX 512 (~2017)
  - 32 registers (512 bits)
  - allows permutting across all 128-bit lanes.
  - gather, scatter and compressed store primitives.
  - one /two/four sockets versions

getting started with some boiler plate code and functions

```cpp
#include <cstdio>
#include <cstdint>
#include <type_traits>
#ifdef __OPTIMIZE__
    #include <immintrin.h>
    #define KEWB_FORCE_INLINE inline __attribute__((__always_inline__))
#else
    #define __OPTIMIZE__
    #include <immintrin.h>
    #undef __OPTIMIZE__
    #define KEWB_FORCE_INLINE inline
#endif

namespace simd {
    using rf_512 = __m512; //float register type
    using ri_512 = __m512i; // int register type
    using msk_512 = uint32_t; //mask
    //..
}
```

### Basic Functions

registers can be treated as groups of values from the same type, and we do the same operation on all of them, masks allow us to choose which registers we change and which not.\
operations are done elementwise.\
we need to consider the order of lsb and msb.\
intrinsics can't be constexpr.

functions have different implementations for float and intgers, but are functionally the same.

- _load_value_ (float and integers) - to fill register with value
- _load_from_ - to fill register with a value from a pointer
- _masked_load_from_ - to load from memory with a mask and register value or a single value overload.
  - a mask means we either keep the value as it is or load from memory.
- _store_to_ - unaligned store in ptr destination
- _masked_store_to_ - store with mask
- _make_bit_mask_ - a template that creates bit masks from.
- _blend_ - combine two registers based on a mask (take from either register a or register b)
- _permute_ - reorder positions of the register based on the the values inside the intgers register
- _masked_permute_ - conditionally choose from a or a permuted version of b.
  - similar to blend with a, permute(b)
  - if mask is off, use a, if on, use the permuted value from b.
- _make_perm_map_ - a template that creates a permutation mask.
- _rotate_ - create a permutation mask and reorder
  - _rotate_down_
  - -rotate*up*
- _shift_down_, _shift_up_ - perform a blend of the rotated values with a register
- _shift_down_with_carry_,_shift_up_with_carry_ - blend two register, from a position, the rotation point partitions from which register ro take the value.
  - like taking a window from two registers, take parts of one register and some parts of another.
- _in_place_shift_down_with_carry_ - change the registers with the contents from the rotate.
- _add_,_sub_ - arithmetics (a+b,a-b)
- _fused_multiply_add_ - multiply two registers and then add a third register ((a\*b) + c)
  - usefull on convulsion algorithms, like a running total (sum product)
- _minimum_,_maximum_ - register with min/max values of the two register

```cpp
KEWB_FORCE_INLINE rf_512 load_value(float v)
{
    return _mm512_set1_ps(v);
}

KEWB_FORCE_INLINE ri_512 load_value(int32_t i)
{
    return _mm512_set1_epi32(i);
}

KEWB_FORCE_INLINE rf_512 load_from(float const * ptr_float)
{
    return _mm512_loadu_ps(ptr_float);
}

KEWB_FORCE_INLINE ri_512 load_from(float const * ptr_int)
{
    return _mm512_loadu_epi32i(ptr_int);
}

KEWB_FORCE_INLINE rf_512 masked_load_from(float const * ptr_float,rf_512 fill, msk_512 mask)
{
    return _mm512_mask_loadu_ps(fill,(__mmask16) mask,ptr_float);
}

KEWB_FORCE_INLINE rf_512 masked_load_from(float const * ptr_float,float fill, msk_512 mask)
{
    return _mm512_mask_loadu_ps(_mm512_set1_ps(fill),(__mmask16) mask,ptr_float);
}

KEWB_FORCE_INLINE void store_to(float * ptr_destination,rf_512 r)
{
    _mm512_storeu_ps(ptr_destination,r)
}

KEWB_FORCE_INLINE void store_to(float * ptr_destination,rf_512 r,msk_512 mask)
{
    _mm512_mask_storeu_ps(ptr_destination,(__mmask16)mask,r)
}

template <unsigned A = 0,....,unsigned P =0>
KEWB_FORCE_INLINE constexpr uint32_t make_bit_mask()
{
    //.. to much code for me to write, maybe I could use a folding expression here...
}

KEWB_FORCE_INLINE rf_512 blend(rf_512 a,rf_512 b,msk_512 mask)
{
    return _mm512_mask_blend_ps((__mmask16)mask,a,b);
}

KEWB_FORCE_INLINE rf_512 permute(rf_512 r,ri_512 perm)
{
    return _mm512_permutexvar_ps(perm,r);
}

KEWB_FORCE_INLINE rf_512 masked_permute(rf_512 a,rf_512 b,ri_512 perm,msk_512 mask)
{
    return _mm512_mask_permutexvar_ps(a,(__mmask16)mask,prem,b);
}

template <unsigned A,....,unsigned P>
KEWB_FORCE_INLINE constexpr ri_512 make_perm_mask()
{
    //static assert
    retrun _mmsetr_epi32(A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P)
}

template<int R>
KEWB_FORCE_INLINE rf_512 rotate(rf_512 r)
{
    if constexpr((R%16)==0)
    {
        return r;
    }
    else
    {
        constexpr int S = (R>0) ? (16 -(R & 16)) : -R;
        constexpr int A = (S+0) R % 16;
        constexpr int B = (S+1) R % 16;
        //...
        constexpr int O = (S+14) R % 16;
        constexpr int P = (S+15) R % 16;

        return _mm512_permutexvar_ps(_mmsetr_epi32(A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P),r);
    }
}

template<int R>
KEWB_FORCE_INLINE rf_512 rotate_down(rf_512 r)
{
    static_assert(R >= 0)
    return rotate<-R>(r);
}

template<int R>
KEWB_FORCE_INLINE rf_512 rotate_up(rf_512 r)
{
    static_assert(R >= 0)
    return rotate<R>(r);
}

template<int S>
KEWB_FORCE_INLINE rf_512 shift_down(rf_512 r)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_down<S>(r),load_value(0.0f), shift_down_blend_mask<S>());
}

template<int S>
KEWB_FORCE_INLINE rf_512 shift_up(rf_512 r)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_up<S>(r),load_value(0.0f), shift_up_blend_mask<S>());
}

template<int S>
KEWB_FORCE_INLINE rf_512 shift_down_with_carry(rf_512 a,ref_512 b)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_down<S>(a),rotate_down<S>(b), shift_down_blend_mask<S>());
}

template<int S>
KEWB_FORCE_INLINE rf_512 shift_up_with_carry(rf_512 a,ref_512 b)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_up<S>(a),rotate_up<S>(b), shift_up_blend_mask<S>());
}

template<int S>
KEWB_FORCE_INLINE void in_place_shift_down_with_carry(rf_512 &a,ref_512 &b)
{
    static_assert(S >= 0 && S<=16)
    constexpr msk_512 z_mask = (0xFFFFu >> (unsigned)S);
    constexpr msk_512 b_mask = ~z_mask & 0xFFFFu;
    ri_512 perm = make_shift_permutations<S,b_mask> ()
    a = _mm512_permutex2var_ps(a, perm,b);
    b = _mm512_maskz_permutex2var_ps((__mmask16)z_mask,b,perm,b)
}

KEWB_FORCE_INLINE rf_512 add(rf_512 a,ref_512 b)
{
    return _mm512_add_ps(a,b);
}

KEWB_FORCE_INLINE rf_512 sub(rf_512 a,ref_512 b)
{
    return _mm512_sub_ps(a,b);
}

KEWB_FORCE_INLINE rf_512 minimum(rf_512 a,ref_512 b)
{
    return _mm512_min_ps(a,b);
}
KEWB_FORCE_INLINE rf_512 maximum(rf_512 a,ref_512 b)
{
    return _mm512_max_ps(a,b);
}
```

now lets build some functions that use those building blocks

### Intra-register Sorting with Sorting networks.

- _compare_with_exchange_ - usefull for sorting, we can sort pairs of positions.

```cpp
KEWB_FORCE_INLINE rf_512 compare_with_exchange(rf_512 vals, ri_512 perm, msk_512 mask)
{
    rf_512 exch =permute(vals,perm); //create a permuted register.
    rf_512 v_min = minimum(vals,exch); // create register of minimums
    rf_512 v_max = maximum(vals,exch); // create register of maximums
    return blend(v_min,v_max,mask); // combine those register by mask.
}
```

> A sorting network (SN) is an abstract device build from:
>
> - A fixed number of "wires" which carry "values"
> - "comparators" which connect pairs of wires and swap the values on the wires if they are not in the desired order.

example:
![wikipedia](https://upload.wikimedia.org/wikipedia/commons/thumb/9/9b/SimpleSortingNetworkFullOperation.svg/650px-SimpleSortingNetworkFullOperation.svg.png)

1. start with unsorted data \[3,2,4,1]
2. first point tests and swaps between the first and third element, but since 3< 4, we don't swap \[3,2,4,1]
3. next, we compare_and exchange second and fourth elements, 2 > 1 so we swap \[3,1,4,2]
4. next, we can do two operations at the same time first and second, third and fourth. 3 > 1 (swap), 4>2 (swap) \[1,3,2,4]
5. and now we compare again, the second and third elements 3 >2 (swap) \[1,2,3,4]
6. our data is now sorted

there are Sorting networks listed for different sizes (number of wires), the less switching points, the better, the optimal networks were proven up to size 12.

we can use this sorting network to sort our registers efficiently.

(this really reminds me of algorithms to get number of bits with set bit masks)

```cpp
KEWB_FORCE_INLINE rf_512 sort_two_lanes_of_8(rf_512 vals)
{
    const ri_512 perm_0 = make_perm_mam<1,0,3,2,5,4,7,6,9,8,11,10,13,12,15,14>();
    constexpr mask_512 mask_0 = make_bit_mast<0,1,0,1,0,1,0,1,0,1,0,1,0,1,0,1>();

    const ri_512 perm_1 = make_perm_mam<3,2,1,0,7,6,5,4,11,10,9,8,15,14,13,12>();
    constexpr mask_512 mask_1 = make_bit_mast<0,0,1,1,0,0,1,1,0,0,1,1,0,0,1,1,>();
    //... repeat this few more times
    vals = compare_with_exchange(vals, perm0, mask0);
    vals = compare_with_exchange(vals, perm1, mask1);
    vals = compare_with_exchange(vals, perm2, mask2);
    vals = compare_with_exchange(vals, perm3, mask3);
    vals = compare_with_exchange(vals, perm4, mask4);
    vals = compare_with_exchange(vals, perm5, mask5);
    return vals;
}
```

he goes over an example of this and show how things get swapped. there will always be the same amount of calls, no branching.

### Fast Medain Filter

if we can sort into two lanes of eight, why not two lanes of seven? if we have 7 elements, the median is the fourth element.

median filters are good

> - Preserving edge features in a singal.
> - Preserving large discontinueties.
> - Eliminating outliers without blur.
> - De-noising.

function avx_median_of_7()
creating a windows of seven values, we run over the data, calculate median of seven, store them in an accumulator.

(some code that I'm not writing)

some benchmarking results. comparing _std::nth_element_, _std::sort_ and the _avx_median_of_7_ (what he built), for sorted values and random values. the simd function works faster, and it's working at linear time.

### Small Kernel Convolution

[Convolution wikipedia](https://en.wikipedia.org/wiki/Convolution).\
convolution, signal S, kernel K, output S*K is the confultion.
"every point of result s*k is equal to S at that point weighted by every point of K"
(something about centering)

real world applications

> - Signal and image processing
> - Probability and processing
> - Computer vision
> - Differential equations

example singal with six data points, kernel with three points, we get a result of size six. we center the kernel (the median value) on each of the signal points, and we start reducing the relevent signal points using the kernel as weights.

$
S\ Signal = s0,s1,s2...s6 \\
K\ Kernel = k0,k1,k2\\
R\ Result = r0,r1,r2...r6\\
r0 = s0*0 + s0k1 + s1k2\\
r1 = s0k0 + s1k1 + s2k2\\
r2 = s1k0 + s2k1 + s3k2\\
r3 = s2k0 + s3k1 + s4k2\\
r4 = s3k0 + s4k1 + s5k2\\
r5 = s4k0 + s5k1 + s6k2 \\
r6 = s5k0 + s6k1 + 0*k2 \\
$

we have windows in the size of the kernel, and we do a sum product on the element-wise multiplication. there is a connection between convolution and correlation. this is fitting for an simd algorithm _avx_convolve_.

(more code that i'm not writing).

using the _fused_multiply_add_ function from before. another sliding window algorithm.

benchmarking again, checking against [Intel MKL Math Kernel Library](https://en.wikipedia.org/wiki/Math_Kernel_Library). we get a nice speed up.

</details>

## Parallelism on Ranges: Should We? - Giannis Gonidelis

<details>
<summary>
Combining Parallel execution algorithms with ranges and views.
</summary>

[Parallelism on Ranges: Should We?](https://youtu.be/gA4HaQOlmSY),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/Parallelism-on-Ranges.pptx)

[HPX](https://github.com/STEllAR-GROUP/hpx) - concurrency and parrallism.

### Algorithms and Ranges

the stl came into life in 1998, with algorithms, containers and iterators. in c++17 parallelism algorithm were included in the stl, and the execution policies were introducted into the world. but we still didn't have:

> - Composability: Coding multiple sequencies is still inconvenient.
> - Performant Composability: immediate effect of lack of Composability.

[range-v3](https://github.com/ericniebler/range-v3) is a library that provides Composability. this makes code more readable, and has the potential to make it much faster.

> A range is:
>
> - an abstraction of "a sequence of items"
> - something iterable
>
> A range is actually:
>
> - a begin iterator & sentinel pair, where sentinel:
>   - an end iterator of the same type as begin iterator
>   - a value
>   - a distance from the begin iterator

in a `c_string` the begin iterator is the start of the chars, and the sentinel is the null-terminator. it can also be the address of the null terminator, or the distance from the start.

we no longer need to pass around the begin and end iterator

```c++
std::vector<int> v{1,2,3,4};
std::find(std::begin(v),std::end(v),3);
// ranges
ranges::find(v,3);
ranges::find(begin(v), sentinel<int>{4},3);
```

for composability, in this example we want to filter squared values which are odd (keep only even squared elements). with stl algorithms, we need to pass around the iterators, and we have temporary values. ranges don't require all that.

```cpp
std::vector<int> vi {1,2,3,4,5};
std::transform(std::begin(vi),std::end(vi),std::begin(vi),[](int i){return i*i;});
auto res = std::remove_if(std::begin(vi),std::end(vi),[](int i){return i%2 ==1;});

//ranges
auto rng = vi |
ranges::view::transform([](int i){return i*i;}) |
ranges::view::remove_if([](int i){return i%2==1;});
std::cout<< rng <<'\n';
```

views are lazy ranges algorithms that evaluate on demand, we only calculate it when we call it. range adaptors take a range and return a view. we employ the pipe operator, just like unix.

in c++20, ranges v3 are partial standardized, but unfortunately, we don't have execution policies with them.

### HPX

HPX, a standard conforming library for concurrency and parallism. it follows the same api as the stanard library. but it does it better. is's also a general purpose library, works for local development and distributed systems.\
provides parallelism and asynchrony, with stl parallel algorithms and "futures" that go past what other libraries provide.

- Reallocate work on the fly, avoid static scheduling.
- Always keep your threads busy, don't let them idle.
- dynamic scheduling of tasks, removing barriers.

uses the standard execution policies:

- sequential execution (`seq`)
- parallel execution (`par`)
- vector execution (`unseq`)
- parallel vector execution (`par_unseq`)
- asynchronous executuion (`par(task)`)
  - this is something we didn't have until now.

more control to the user over the parallelization.

we no longer block the execution, and the execution waits until we need the future.

```cpp
future<int> f1 =async(&fun);

// or

future<void> f2= for_each(par(task), std::begin(v),std::end(v), /* some lambda*/);


f2.get();
//or
f2.then(
    /* do next thing*/
)
```

hpx algorithm support

```cpp
hpx::reduce(par,std::cbegin(v),std::cend(v),/*some lambda*/);
//async
hpx::reduce(par(task),std::cbegin(v),std::cend(v),/*some lambda*/);
//ranges overloads
hpx::ranges::reduce(v,/*some lambda*/);
hpx::ranges::reduce(std::begin(v),sentinel,/*some lambda*/);
```

### Parallel Ranges

combining ranges and execution policies,

base form

```cpp
hpx::for_each(par, v.begin(), v.end(),/*lambad*/)
```

range form

```cpp
namespace hpx {
    namespace ranges{
        result_type for_each(ExPolicy policy, Rng rng, F f)
        {
            return for_each(policy, hpx::util::begin(rng),hpx::util::end(rng),f);
        }
    }
}
```

stage 1.5, iterator and sentinel

```cpp
namespace hpx {
    namespace ranges{
        result_type for_each(ExPolicy policy, Iter iter, Sentinel sent, F f)
        {
            auto new_end_iter = //do something with sentinel to get the end iterator with ranges::next, ranges::advance, ranges::distance... etc
            return base_impl::for_each(policy, iter,new_end_iter,f);
        }
    }
}
```

but the final goal is to use ranges and views,

```cpp

std::vector<int> vi {1,2,3,4,5};
auto rng = vi |
ranges::views::transfrom([](int i){return i*i;}) |
ranges::views::remove_if([](int i){return i % 2 ==1;});
```

options

> 1. provide combined implementations for each combination of operators (combinatorial explosion)
> 2. use fork-join strategy (also rejected)
> 3. fusion (this was chosen)

views are lazily evaluated, so we fuse together the stages.

some operation combinations are harder to parallelize like this than others:

> hard:
>
> - transform | remove_if
> - adjacent_remove_if | reverse
>
> easy:
>
> - transform | reverse
> - accumulate | transform

this depends on how the iterator types is exposed, and when we have temporaries, container resizing and predcates about more then one element things are more difficult.

```cpp
std::vector<int> vi(10'000'000);
std::iota(std::begin(vi),std::end(vi),1);

auto rng = vi |
ranges::views::transform([](int i){return i*i;}) |
ranges::views::reverse;


hpx::ranges::for_each(hpx::execution::par, rng,[](auto i){return i;});
```

hpx stages:

> - c++20 conformance
> - parallelize when single range argument input
> - parallelize when iterator-sentinel input
> - parallelize when input is composed from a chain of views

### Results

some things don't get performance boost from parallelization, and some do.

### Future Work

should we parallelize ranges?\
sometimes, yes. there are good and bad cases, we should take advantage of inherent fusion.

</details>

## Executors: The Art of Generating Composable APIs - Nikunj Gupta

<details>
<summary>
Executors, and composable operations
</summary>

[Executors: The Art of Generating Composable APIs](https://youtu.be/8rRTKWdfAOU),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/CNow-2021.pptx)

HPX - task based parallelism model, standard confirming with similar syntax. supports parallel, distributed and heterogenous applications, has light-weight threads. similar syntax for local and remote operations.

### Resilience

Exascale computing - 10^18 operations.\
SDC - silent data corruptions, not detected. usually have low probability for happening in a single processor, but will happen for thousends of them. do we even care about them?

### HPX Implementation

> assumptions:
>
> - No global variables for state changes.
> - use built-in constructs (channels)
> - Task do not change the input data parameters.
>   Task boundary is an ideal position to add resilience

example: task 1 computes a result and feeds it to task 2. but if there was a silent error, we can check the value (add resiliency, credability) before passing on the data. we use _Task Replay_ and _Task Replicate_.

async replay: do task A, if there is an exception, replay the task, if not, continue. this is done recursively.

async replicate: do task A some times.

```cpp
template <typename F, typename... Ts>
auto async_replay(std::size_t n, F&& f,TS&&... ts)
{
    using result_t = typename std::invoke_result<F,Ts..>::type;
    return detail::async_replay_helper<result_t>(n, std::forward<F>(f),std::forward<Ts>(ts)...);
}

template <typename Result, typename F, typename ... Ts>
hpx::future<Result> async_replay_helper(std::size_t n, F&& f,TS&&... ts)
{
    hpx::future<Result> f_ = hpx::async(f,ts...);
    return f_.then(hpx::launch::sync,
    [n, f=std::forward<F>(f),...ts = std::forward<Ts>(ts)](hpx::future<Result>&& f_)
        {
            if (f_.has_exception())
            {
                //get handle to exception
                auto ex = rethrow_on_abort_replay(f_);
                if (n!=0)
                {
                    return async_replay_helper(n-1,std::forward<F>(f),std:forward<Ts>(ts)...);
                }
                std::rethrow_exception(ex);
            }
            return hpx::make_ready_future(f_.get());
        }
    );
}
template <typename F, typename... Ts>
auto async_replicate(std::size_t n, F&& f,TS&&... ts)
{
    using result_t = typename std::invoke_result<F,Ts..>::type;

    std::vector<hpx::future<result_t>> results;
    results.reserve(n);

    for (std::size_t i =0; i!=n; ++i)
    {
        results.emplace_back(hpx::async(f,ts...));
    }
    return hpx::dataflow(
        hpx::launch::sync,
        [n](std::vector<hpx::future<result_t>>&& results) mutable {
            std::exception_ptr ex;
            for (auto && f: std::move(results))
            {
                if (!f.has_exception())
                {
                    return hpx::make_ready_future(f.get());
                }
                else
                {
                    ex =rethrow_on_abort_replicate();
                }
            }
               std::rethrow_exception(ex);
        },std::move(results));
}
```

### Implementation Variations

Algorithm based fault tolerance, based on validation function.

we can use the async replicate function to validate, as we have more than one valid result:

- intoduce consensus through vote functions
- introduce results validation through predicates
- introduce consensus on valid results from predicates.

**distributed software resilience**:\
 we need entities that are serializable, we can't send function pointers over network because of how the address randomizer works.

```cpp
template <typename Result, typename Pred, typename F, typename...Ts>
auto async_replay_helper(std::size_t n, Pred&& pred, F&& f, Ts&&... ts)
{
//..
//.. within lambda after `if(f.has_exception())`

auto && res = f.get();
if (!HPX_INVOKE(pred, res)&& n != 0)
{
// validation failed
// try again, with n-1;

return async_replay_helper(n-1, std::forward<Pred>(pred), std::forward<F>(f), std::forward<TS>(ts)...);
}
return hpx::make_ready_future(std::move(res));
}
```

now we have some results, and we want to reach a consensus

```cpp
template <typename Result, typename Vote, typename F, typename...Ts>
auto async_replicate_vote(std::size_t n, Vote&& vote, F&& f, Ts&&... ts)
{
//..
//.. within gpx::dataflow (vote is forward captured in the lambda

std::vector<hpx::future<Result>> exceptionless_results;
exceptionless_results.reserve(n);

std::exception_ptr ex;

for (auto&& f:std::move(results))
{
if (!f.has_exception())
{
exceptionless_results.emplace_back(f.get());
}
else
{
ex= rethrow_on_abort_replicate();
}
}

if (exceptionless_results.empty()
{
std::rethrow_exception(ex);
}

// where did valid results come from?
return hpx::make_ready_future(HPX_INVOKE(std::forward<Vote>(vote), std::move(valid_results));
}
```

the same scenario, but on different machine (distributed), we send the command over the network and then other machine does the action.

```cpp
template <typename Result, typename Vote, typename Action, typename...Ts>
auto async_replicate_vote(std::vector<hpx::id_type> ids, Vote&& vote, Action&& action, Ts&&... ts)
{
using result_t = typename std::invoke_result<Action, hpx::id_type, Ts..>::type;
std::vector<hpx::future<result_t>> results;
results.reserve(ids.size());

for (std::size_t i = 0; i != ids.size(); ++i)
{
    results.emplace_back(gpx::async(action,ids.at(i),ts..));
}
//..
}
```

the performace cost is based on how many futures are accessed, so there a small performance cost for replay+validate, but a high cost for replicate+validate.

some benchmarking.

### The Need For Executors

> if overheads are low, why not use it everywhere?

```cpp
auto f1 = hpx::async(my_func, args...);
//can be converted into
auto f2 = hpx::async_replay(n,my_func, args...);

auto f3= my_algorithm(args...);
//can be converted into
auto f4 = hpx::async_replay(n, my_algorithm, args);;

hpx::for_each(hpx::execution::par, my_range.begin(), my_range.end(), my_func);
//doesn't convery nicely
```

> "Executors are modular components for creating execution"\
> (P0443,2016)

executors work on an executing resource and provide abstraction over it.

```cpp
template<InputRange Ir, OutputRange Or>
auto some_algorithm(Ir&& ir, Or&& or)
{
//some work
}

//executor unaware algorithm
template<Executor Ex,InputRange Ir, OutputRange Or>
auto some_algorithm(Ex ex,Ir&& ir, Or&& or)
{
ex.execute(/* some work*/);
}

//executor aware algorithm
template<Executor Ex,InputRange Ir, OutputRange Or>
auto executor_aware_algorithm(Ex ex,Ir&& ir, Or&& or)
{
return algorithm(ex, std::forward<Ir>(ir), std::forward<Or>(or));
}
```

now we can have clean and composable API

```cpp
auto f1 = hpx::async(my_func, args...);
//can be converted into executor
auto f2 = hpx::async(ex,my_func, args...);

auto f3= my_algorithm(args...);
//can be converted into executor
auto f4 = my_algorithm(ex,args...);

hpx::for_each(hpx::execution::par, my_range.begin(), my_range.end(), my_func);
//can be converted into executor!
hpx::for_each(hpx::execution::par.on(ex), my_range.begin(), my_range.end(), my_func);
```

hpx executors (based on P0443R4):

member function:

- post - fire and forget
- sync_excute - blocking , like std::invoke
- async_excute - non blocking, like std::async(func, args...)
- bulk_async_excute - async_excute, but in bulk
- then_execute - support `.then()`
- bulk_then_execute - bulk version `.then()`

an executor can have one or more of those function. we want compile time performance, so we create customization points objects. we have executor categories

- is_one_way_executor - no channels to return results
- is_two_way_executor - has return results
- is_bulk_two_way_executor - for bulk operations.

### example

```cpp
hpx::async(ex, func, args...);
// calls
template<typename Executor>
struct async_dispatch<Executor, typename std::enable_if<traits::is_one_way_executor<Executor>>::value || traits::is_two_wat_executor<Executor>::value>::type>;

async_execute(std::forward<Executor(exec), std::forward<F>(f), std::forward<Ts>(ts)...);

exec.async_execute(std::forward<F>(f), std::forward<Ts>(ts)...);
```

now we go back to the resilience replay executor and add a way to handle two way execution

```cpp
template<typename BaseExecutor, typename Validate>
class replay_executor
{
private:
BaseExecutor & exec_;
std::size_t replay_count_;
Validate validator_;

public:

template<typename F>
explicit replay_executor(BaseExecutor& exec, std::size_t n, F&& f)
: exec_(exec), replay_count_(n), validator_(std::forward<F>(f))
{}

template<typename F, typename...Ts>
auto async_execute(F&& f, Ts&&... ts)const
{
return async_replay_validate(exec_, replat_count_, validator_, std::forward<F>(f), std::forward<Ts>(ts)...);
}
//...
};
```

and for the bulk two way executor, we add to the above class

```cpp
template <typename F, typename S, typename..Ts>
auto bulk_async_execute(F&& f, S const& shape, Ts&&... ts) const
{
using namespace hpx::parallel::execution;
std::size_t size = hpx::util::size(shape);
using result_type= typename detail::bulk_function_result<F,S,Ts...>::type;
using future_type= typename executor_future<BaseExecutor, result_type>::type;

std::vector<future_type> results;
results.resize(size);

hpx::lcos::local::latch l(size+1);

spawn_hierarchical(results,l, 0,size, num_task, f, hpx::util::begin(shape), ts...);
l.count_down_and_wait();
return results;
}
// this should be somewhere in teh spawn_hierarchical function
results[base+i] = async_execute(func, *it, ts...);
```

and the driver code itself

```cpp

hpx::execution::parallel_executor base_exec;
auto exec = hpx::resillency::experimental::make_replay_executor(base_exec,3);

auto f= hpx::async(exec, fuc, args...);
some_algorithm(exec, args...);
hpx::for_each(hpx::execution::par.on(exec), my_range.begin(), my_range.end(), my_func);
```

virtually no effort for the user, easy to add. it also produces clean and readable code as compared to replicate and replay, the executors are composbile!

> - Resilience executors are base-executor unaware.
> - Resilience executors are algorithm unaware.
> - Resilience executors are runtime unaware.

```cpp
hpx::kokkos::default_host_executor exec_;
auto exec = hpx::kokkos::resiliency::make_replay_executor(exec_, n, validate);
auto f = hpx::async(exec, func, args...);
```

</details>

[Main](README.md)
