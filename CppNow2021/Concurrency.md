<!--
// cSpell:ignore simd Steagall intrinsics cstdio immintrin loadu mmask storeu permutexvar permutex2var mmsetr maskz fmadd Giannis Gonidelis asynchrony KEWB unseq Nikunj Exascale randomizer kokkos hpx lcos Harel lconcore luceto cuda cudaflow sycl syclflow saxpy
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

## Converting a State Machine to a C++ 20 Coroutine - Steve Downey

<details>
<summary>
making a state machine using c++20 Coroutines.
</summary>

[Converting a State Machine to a C++ 20 Coroutine](https://youtu.be/Z8jHi9Cs6Ug), [slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/04/convert-state-machine-coroutine-slides-1.pdf)

> C++ 20 coroutines can naturally express in linear code components that are today written as state machines that wait on async operations.\
> This talk walks through using the low-level machinery and customization points in c++20 to convert a state machien, which waits at the end of steps for async service operations to complete, into a single coroutine that `co_awaits` those operations.

### Basics

C++20 Co Routines: Inaccurate summary

like a lambda, excepts:

> - the lambda is the return type
> - they control when they suspend
> - no stacks, threads or fibers

Stackfull vs stackless?

> stackless
>
> - they execute on the regular stack
> - the architectural model is very different from from fibers or threads.
> - Coroutine == Resumable Statefull Function

if it has `co_await`, it's a coroutine. there are some versions of a co_awaits:

- co_await
- co_yield
- co_return

a coroutine body.

```cpp
{
   promise-type promise {promise-constructor-arguments};
   try{
       co_await promise.initial_suspend();
       // function-body
   } catch(...){
       if (!initial-await-resume-called)
       {
            throw;
       }
       promise.unhandled_exception();

   }
   final-suspend:
    co_await promise.final_suspend();
}
```

> **terms defined:**
>
> - promise-type: determined by coroutine_traits<>, but usually a typedef in the return type.
> - promise-constructor-arguments: there parameters if there's a valid overload from promise-type that takes them, otherwise empty.
> - function-body: the body of the coroutine function
> - initial-await-resume-called: was the await_resume of the initial suspend called? did we start?
> - final-suspend: target for `co_return` which calls either `return_value` or `return_void` first then executes `goto final-suspend`.
>
> **awaitables**
>
> - bool await_read(): proceed or suspend, false is suspend.
> - await_suspend: callied if `await_read` is (contextually) false
>   - `void await_suspend(coroutine_handle<> h):` call `await_suspend` and suspend.
>   - `bool await_suspend(coroutine_handle<> h)`: call `await_suspend` and resume if false.
>   - `std::coroutine_handle<Z> await_suspend(coroutine_handle<> h)`: call `resume` on return.
> - T await*resume(): call when resume, T is the results of co_await. \_Awaitable* interface is programmer facing.

minimal example

```cpp
template <typename T>
struct awaitable: public std::suspend_always{
    //constexpr bool await_read() const noexcept{return false;} //from std::suspend always
    costexpr void await_suspend(coroutine_handle<> H) const noexcept {h.resume();}
    costexpr T await_resume() const noexcept {return T{};}
}
```

> **promises**
>
> - ReturnType::promise_type: typedef for the promise.
> - get_return_object(): the return type of the coroutine.
>   - `return_value()` : return value or...
>   - `return_void()`
>   - `return_`
> - initial_suspend(): initial suspend before body
> - final_suspend(): final suspend after body
> - unhandled_exception(): called if an exception escapes Promises and coroutine return types are library writer facing.
>
> GCC's Implementation is almost exactly a lambda\
> Theres an insane of an unnameable type that is tied to the particular coroutine frame, the type has a bit of astate that indicates where the `jmp` to upon entry goes to. The coroutine function allocates one of these, and ties it to the return type via the promise.

minimal example:

```cpp
#include <coroutine>

struct MinimalCoRo{
    struct promise_type{
        MinimalCoRo get_return_object(){
            return {.h_: std::coroutine_handle<promise_type>::from_promise(*this)};
        }
        std::suspend_always inital_suspend() noexcept{return {};}
        std::suspend_always final_suspend() noexcept{return {};}
        void unhandled_exception(){}
    };
std::coroutine_handle<promise_type> h_;
};

void before();
void after();

MinimalCoRo func()
{
    before();
    co_await std::suspend_always{};
    after();
}
```

### A Bit of Therory

> UML State Diagrams\
> Describes a "finite automaton", standardized as part of the Unified modeling language back in the last century.

- activation/deactivation
- sub-states
- orthogonal regions, effect nearly everything (think <kbd>CapsLock</kbd> and <kbd>NumLock</kbd> keys)

> UML based of Harel state charts:\
> A generalization of state machine diagrams more usable for human being, allows for grouping states with the same parameters together as the substate charts. Allows for history, returning to a state with the substate active when the the superstate left. A full formal model.

UMLs are models, but it doesn't necessarily translate 1-to-1 to code. tools can generate code based on a model, but it won't be the best (easiest, maintaible) way to express it. the state chart is a documention tool in some cases.

> **The Core Coroutine Transformation is to a State Machine**\
> C++20 coroutines are resumable functions. a coroutine is transformed into:
>
> - a handle to the frame holding the stack variables.
> - an indicator of where to resume
> - an instance comprising this particular execution.

The State is maintained in the coroutine frame. the coroutine frame is equivalent to the member variables of an object.\
`co_await` points are the states: the coroutine is waiting for input (to resume).\
_Resumptions_ are transitions firing. When a transition fires the coroutine can decide how to proceed to the next state.

State machines aren't just ways to implement regex, there are large state machines and (mostly) small state machines. for large state machine management tools are needed.\
writing down the state machine model helps clarify the transitions.

most state machines are simple, and have different paths:

- **Golden path** - things go well.
- **Error path** - things go badly in expected ways:
  - bad input
  - file not found
- **Failure path** - things go badly in unexpected ways:
  - "2+2 ==5"

The 7&pm; rule (seven plus-minus two), this is about the size of a state machine we can mentally model, anything larger requires extrating substates or using management tools to maintain.

> "Generality might mean `goto`"

some times statemachines have states that can be reached from any other state, and states might need to go forward of backwards. this is ok, because we don't leave the scope of the machine / coroutine.

> **Suspension and Decision**: guarded transitions just _if tests_ after a suspension point.

in the diagrams, these are labels next to a transition. decision from input on where to transition.

there are not standard library solutions and coroutine types defined as of c++20. there might be some added in the c++23 release. this isn't something new, in earlier versions of c++ it was expected of users to write their own containers and iterators types. it's okay for users and library writers to write and handcraft types. this is part of how the standardization process works. the community understands what is needed, what works and what is important. coroutines can be implemented by the users, and any additions to the standard won't break them.

### Simple Multistep Async Operations

basic example,(not actual production code)

```cpp
class CreateUser
{
    CreateUser (std::string id); //constructor
};
```

Lookup user or create

```cpp
Result CreateUser::findUser(){
    db::getUser(id,[](std::unique_ptr<User> user){
        userCallback(user);
        });
    return CONTINUE;
}

void CreateUser::userCallback(std::unique_ptr<User> usr)
{
    user_ = std::move(user);
    resume_();
}
```

Validate request wih "Compliance"

```cpp
Result CreateUser::findUser()
{
    compliance::checkOK(user_, [] (bool isOK){
        complianceCallback(isOk);
        });
    return CONTINUE;
}

void CreateUser::complianceCallback(bool isOK)
{
    isOK_ = isOK;
    resume_();
}
```

Broadcast operation

```cpp
Result CreateUser::broadcastNewUser()
{
    if (isOK_){
        queueBroadcast(*user);
    }
    return CONTINUE;
}
```

Return status for request

```cpp
Result CreateUser::endTransaction()
{
    return DONE;
}
```

```cpp
class CreateUser
{
    Result CreateUser::findUser();
    Result CreateUser::okToCreate();
    Result CreateUser::broadcastNewUser();
    Result CreateUser::endTransaction();

    void CreateUser::userCallback(std::unique_ptr<User> user);
    void CreateUser::complianceCallback(bool isOK);

};
```

> "Natural Non-Async Code is the Inverse of Coroutine Transform": if this were all synchonous it would just be a sequence of calls.

but we don't want to simple wait for responses and block operations. we don't want to tie up the thread.

> "While Not Done" : externally this is driven checking if the object said it was done, and if not, scheduling the next operation.

### Async Callbacks and Threads

```cpp
void (* callback)(void * context, void * response, void * error); //function type declration
void install (callback cb, void * context);
```

> Typical generic C-ish callback interface
>
> - you give the framework the context to give back to you
> - it gives you the response you were waiting for
> - alternatively or additionally it tells you about any errors.
>
> C++ callback is often a type-erased callable, like `std::function<>`, binding `this` and other parameters.

the context is the _this_ pointer or the coroutine frame. we (or the framework) cast it back to the known type.

a frequent source of errors is with threads, we might run into deadlocks and issues with locks. IO stalling. we might need to make use of threadpools and reschedule operations back to them, and the problems compound.

**converting a callback to an awaitable**

```cpp
void api_with_callback(std::string p, std::function<void(int result)> callback);
auto api_with_callback_awaitable(const std::string* parameter)
{
    struct awaiter :
    {
        std::sting parameter_;
        int result_;
        awaiter(const std::string& parameter): parameter_(parameter){}
        bool await_ready(){return false;} // suspend always
        void await_suspend(std::coroutine_handle<> handle)
        {
            api_with_callback(parameter_,[this,handle](int result){
                result_=result;
                handle.resume();
            });
        }
        int await_resume { return result_; }
    }

    return awaiter(parameter);
}
```

rescheduling on the thread pool

```cpp
// for exposition only
void thread_pool::await_suspend(coroutine_handle<> handle)
{
    schedule(job([](){ handle.resume(); }));
}
```

> **Coroutines are Not Async** : theres no magic that makes them asynchronous.\
> **Coroutines are Deterministic**: transfer of control from the coroutine is deterministic, either outward to te owner or to a particular coroutine. resumption of a coroutine is direct.

direct - not to a thread or fiber, just back to the normal stack. the frame is stored on the heap

> **Suspension is Not Async**: nothing happens to a suspend coroutine, there are no threads.\
> **Transfer of Control is Sync**: suspension hands control on the same thread. Resumption happens on the same thread as the resumer.\
> **Async is External to the Coroutine**: Async can be built with coroutines, but it's external to the coroutine mechanism itself. Sync can be built from Async, the other way around is far more difficult.

if we want async operations, we need to build them, there is nothing inherently asynchronous about coroutines. this is because coroutines are stackless. suspending a coroutine doesn't end the scope. if we have a lock, then it's not unlocked when we suspend the function.

```cpp
task<Excpected<std::unique_ptr<User>,bool>> createUSer(std::string id)
{
    std::unique_ptr<User> user = co_await db::getUser(id);
    co_await threadpool_;
    bool isOK = co_await compliance::checkOk(user);
    co_await threadpool_;
    if (isOK)
    {
        queueBroadcast(*user);
    }

    co_return {std::move(user), isOK};
}
```

in the code above:

- logic is clearer
- writing new async state machines is easier

the code is linear. `co_await` the threadpool is to request a reschedule for ourselves.

### Questions

- _canceling or timeout a coroutine which is async_
- _tla + modeling_ (TLA: Temporal Logic of Actions)
- _compiler stuff_
- _changing behavior based on internal state_
- _when does the work happen_ - it happens in compile time
- _allocation costs and efficiency_
- _`co_awaiting` a list of tasks_
- _benchmarks and scaling_
- _writing unit tests for coroutines_

</details>

## Designing Concurrent C++ Applications - Lucian Radu Teodorescu

<details>
<summary>
High level abstraction without using low-level primitives.
</summary>

[Designing Concurrent C++ Applications](https://youtu.be/nGqE48_p6s4),[slides](http://lucteo.ro/content/pres/C++Now2021-Designing-Concurrent-C++-Applications-pres.pdf), [github code examples](https://github.com/lucteo/cppnow2021-examples), [No Locks Manifesto](http://nolocks.org/).

a graph showing what people find frustrating, with concurrency safety issues being on top.

the talk will try to build a high level concurrency framework, mostly without locks, that will be high performant.

### Threads Considered Harmful

a talk from earlier in the year [Threads Considered Harmful](https://youtu.be/_T1XjxXNSCs).

threads in this context mean raw threads + synchronization (locks), the problems are: Performance, Understandability, Thread-Safety, and Composability. it's very likely to make a mistake and get it wrong and cause a problem with on of them.

we want a general method, without locks, without safety issues (as much as possbile), with good performance, and have it being composobale and decomposable.\
this will be done by using tasks(independent units of work), those tasks have all the dependencies explicitly stated. a unit of work is a series of instructions.

there are article in "overload" journal showing the theoritcal results:

> - all concurrent algorithms
> - safety insured
> - no need fo locks
> - high efficiency for greedy algorithm
> - high speedups
> - easy composition and decomposition

this doesn't include GPU, SIMD and c++20 coroutines.

### Conccurent Design by Example

concurrency without using locks.

we start with an example, we use the _concore_ library, but we can use other libraries as well, we care about the design, not the implementation.

```cpp
#include <concore/spawn.hpp>

int main() {
    // Create a task and executes it
    // The task can run in the same thread, or a different thread
    concore::spawn_and_wait([] {
        printf("Hello, concurrent world!\n");
    });

    return 0;
}
```

a bit more serious example, creating task to run concurrently.

```cpp
#include <concore/spawn.hpp>

#include "../common/utils.hpp"

void print_message_task(const char* msg) {
    CONCORE_PROFILING_SCOPE();
    CONCORE_PROFILING_SET_TEXT(msg);

    printf(" %s", msg);

    sleep_for(100ms);
}

int main() {
    profiling_sleep profiling_helper;
    CONCORE_PROFILING_FUNCTION();

    // Create a task group, so that we keep track of the running tasks
    auto grp = concore::task_group::create();

    // Create 9 tasks to be run concurrently
    concore::spawn([=] { print_message_task("How"); }, grp);
    concore::spawn([=] { print_message_task("did"); }, grp);
    concore::spawn([=] { print_message_task("the"); }, grp);
    concore::spawn([=] { print_message_task("multi-threaded"); }, grp);
    concore::spawn([=] { print_message_task("chicken"); }, grp);
    concore::spawn([=] { print_message_task("cross"); }, grp);
    concore::spawn([=] { print_message_task("the"); }, grp);
    concore::spawn([=] { print_message_task("road"); }, grp);
    concore::spawn([=] { print_message_task("?"); }, grp);

    // Ensure that all the tasks are completed
    // This performs a BUSY WAIT -- it takes tasks and executes them
    concore::wait(grp);

    printf("\n");
    return 0;
}
```

- Tracy profiler
- Spawning tasks & waiting for them
- Task system

we can rebuild the code above with the profing option enabled
`clang++ -std=c++17 -DTRACY_ENABLE=1 -I/Users/luceto/work/other/tracy -stdlib=libc++ -lconcore -lconcore_profling -o out/02_fork 02_fork.cpp`. and now wee see the timeline of the threads.

the number of threads created is equal to the number of cores in the macine, and then a thread can be reused.

example 03.1: using a callback with tasks, in this example the task is executed on the same thread. example 03.2: using tasks vs using mutexes. example 03.3 uses a chain of async operations, makes use of templates.

example 4 is about joining tasks (waiting for them to finish), and we have an option using a task group to set the order of execution.

example 5 is _fork-join_, divide and conquer approach. we split the task into smaller chunks, each time creating a new task, either as new thread or the existing, and then we wait for the parts to finish.

example 6 is _concurrent for_, which splits the work across threads, similar to `std::for_each(std::execution::par, int_iter{0}, int_iter{20},work)`.

example 7 is _concurrent_reduce_, which tries to create a single value from multiple value.

example 8 is _concurrent_scan_, each input produces an output, but each output requires knowledge of the previous inputs. in this case, we use a prefix-sun.

example 9 is about _task graphs_, a series of tasks which depend on another in a known way, the number of threads used is determined by the depencies between the tasks.

example 10 is _pipeline_, we can set the order and the concurrency model, so some tasks need to be called in a certain order, and some can be run together with others.

example 11.2 is _serializers_, in this example we have a running window average, which for some reason cannot be used in a concurrent method, so the serializer is an executor that ensures the safety. the tasks can be run in different threads, but never concurrently. this is a way to avoid using mutexes.\
example 11.2 is a _read-write serializer_, which replaces the read-write problem, so we no longer use mutexes. example 11.3 sets a limit on the number of parallel operations,so it replaces the semaphore.

this concludes the first part, we now see that we can use tasks as high-level concurrency abstractions, and we have no need for low-level primitives.

### C++23 Executors

all the examples comply with the proposed executors of c++23.

- executors
- senders and receivers
- senders algorithm

example 1 shows how the abstractions work within the executor framework. executors are really simple.

example 2.1 _senders and receivers_ demonstrate a connection between a sender and receiver, the scheduler from the thread pool creates the operation state. we skip example 2.2. in example 2.3 we show custom sender and receivers.

example 3 is _sender algorithms_, as proposed in c++23.

### Performance Topics

> Targeting throughput. latency is also a concern, but not the main one.

in a global pool of workers threads, we usually one thread per core, but if we know our tasks have large wait time, we can have more threads.\
the important thing is to have more tasks than cores, we want to always have something running and getting work done. keep threads busy.

there is a small overhead for the library, so the tasks should be big enough to make the process worth it.

example 1 _cpu_intensive_, we try to keep the cpus busy, if there aren't doing work, we're wasting time and losing progress.

example 2 _limit threads_. no example 3. example 4 shows the difference in speedup depending on the number of threads, the best performance is twice the number of cores (because of hyper-threading).

example 5 shows how serializers compare to mutex. we see the times it takes and how mutexes prevent us from using all of our cpu. we skip example 6.

### Building New Concurrency Abstractions

> Extensibility is the key

the standart won't ship with all we the need, we will have to create our own implementations for the first period of time.

we have an example of composition and decomposition, the same as the earlier pipeline example. we can change the steps without effecting the pipeline and the concerns, we mix concurrent abstractions together.

in example 2 we mix serializers, in example 3 we have a partial priority serializer, example 4 is matrix processing, example 5 is data streams, which reacts to a source in real-time.

### Conclusions

Concurrency without locks is possible. it's not complicated to write or to extend. the low level primitives exists in the framework level, not the user code. we even get good performance.

</details>

## Taskflow: A Heterogeneous Task Graph Programming System with Control Flow - Tsung-Wei Huang

<details>
<summary>
A lightweight framework to run tasks in parallel processing and create execution graphs.
</summary>

[Taskflow: A Heterogeneous Task Graph Programming System with Control Flow](https://youtu.be/4XhH0XN0zQQ),[slides](https://cppnow.digital-medium.co.uk/wp-content/uploads/2021/05/talk.pdf), [TaskFlow github](https://taskflow.github.io), [Proflier/Visualizer](https://taskflow.github.io/rfprof)

Agenda:

- Express your parallisem in the right way
- Parallelize your applications using Taskflow
- Understand our scheduing algorithm
- Boost performance in real applications
- Make C++ amenable to heterogenous parallelism

parallel computing makes computation faster by several orders of magnitude, modern computers CPU have several cores. however, they are challenges to parallel computing

- Debuging
- Dependency constraints
- Concurrency control
- Scheduing efficiencies
- Task and data race
- Dynamic load balancing

> How can we make it easier for C++ developers to quickly write parallel and heterogenous programs with _high performance and scalability_ and _simultaneous high productivity_?\
> **Taskflow** offers a solution

a simple example, we runt to run four tasks, A,B,C,D, where A must run first, D must run last, and B and C can run in either order or in parallel.

```cpp
#include <taskflow/taskflow.hpp> //Taskflow is header only

int main()
{
  tf::Taskflow taskflow;
  tf::Executor executor;

  auto [A,B,C,D] = taskflow.emplace(
    [] () { std::cout << "TaskA\n"},
    [] () { std::cout << "TaskB\n"},
    [] () { std::cout << "TaskC\n"},
    [] () { std::cout << "TaskD\n"}
  );

  A.precede(B,C); // A runs before B and C
  D.succeed(B,C); // D runs after B and C
  executor.run(taskflow).wait() // submit the taskflow o the executor
  return 0;
}
```

to run the example we need to clone the repository and include the library. it's a header-only library, and doesn't require linking.

```sh
git clone htttps://github.com/taskflow/taskflow.git #clone once
g++ -std=c++17 simple.cpp -I taskflow/taskflow -O2 -pthread -o simple
./simple
```

also a built-in profiler/visualizer

```sh
TF_ENABLE_PROFILER=simple.json ./simple
cat simple.json
#paste the profiling json data to https://taskflow.github.io/tfprof/
```

### Express your Parallisem in the Right Way

the motivation was - VLSI-CAD tools, designing chips.

(a horbile looking graph)

> "how can we write efficient C++ parallel programs for this monster computation task graph with **millios of CPU-GPU dependant tasks along with algoithmic control flow**?"

the existing tools were not sufficent:

1. complex task dependencies.
   > existing tools are good at loop parallism, but weak in expressing heterogenous task graphs at this large scale.
2. complex control flow - dynamic control flow(combinatorial optimization, analtical methods)
   > existing tools are _directed acyclic graph_ (DAG)-based and do not anticipate cycles or conditional dependencies, lacking _end-to-end_ parallelism.

an example of an **iterative optimizer**

> 4 computation tasks with dynamic control flow:
>
> 1. starts with _init_ task
> 2. enters the _optimizer_ task (e.g. GPU math solver)
> 3. check if the optimization converged: if not, loop back, if yes, continue to the final stage
> 4. output the results
>
> how can we easily descrive this workload of dynamic control flow using existing tools to acehive end-to-end parallelism?

designing parallel programing is not trivial. we need a way (infrastructure, framework) to express the dependencies of heterogenous task in a parallel, multi-computation enviornment.

### Parallelize your Applications using Taskflow

task flow has five task types:

1. static tasks - basic task on a callable object
2. dynamic tasks - dynamic parallelism task
3. cudaFlow/syclFlow task - gpu task
4. condition task - control flow of the parallelism
5. module task - composable tasks

#### Static Tasks

basic code example in OpenMP library:

```cpp
#include <omp.h> //OpenMP is a langauge extension to describe parallelism using compiler directives
int main()
{
  #omp parallel num_thread(std::thread::hardware_concurrency())
  {
    int A_B ,A_C, B_D, C_D;

    #pragma omp task depend(out: A_B,A_C) //task depedency clauses
    {
      std::cout<<"TaskA\n";
    }
    #pragma omp task depend(in: A_B,out: B_D) //task depedency clauses
    {
      std::cout<<"TaskB\n";
    }
    #pragma omp task depend(in: A_C,out: C_D) //task depedency clauses
    {
      std::cout<<"TaskC\n";
    }
    #pragma omp task depend(int: B_D,C_D) //task depedency clauses
    {
      std::cout<<"TaskD\n";
    }
  }
  return 0;
}
```

> in openMP, task clauses are static and explicit, programmers are responsible for the proper order of writing task consistent with sequential execution.

the order of the code must be the same as the order of execution.

example of TTB flow graph library by Intel,

```cpp
#include <tbb.h> // Intels's TBB is a general-purpose programming library in C++
int main()
{
  using namespace tbb;
  using namespace tbb::flow;
  int n = task_scheduler init::default_num_threads();
  task task_scheduler init(n);
  graph g; // dependency flow graph

  // declaring tasks as a continue_node
  continue_node<continue_msg>A(g,[](const continu_msg &) {
    std::cout<<"Task A\n";
  });

  continue_node<continue_msg>B(g,[](const continu_msg &) {
    std::cout<<"Task BA\n";
  });

  continue_node<continue_msg>C(g,[](const continu_msg &) {
    std::cout<<"Task C\n";
  });

  continue_node<continue_msg>D(g,[](const continu_msg &) {
    std::cout<<"Task D\n";
  });

  make_edge(A,B);
  make_edge(A,C);
  make_edge(B,D);
  make_edge(C,D);
  A.try_put(continue_msg());
  g.wait_for_all(;)
}
```

> TBB has excellent performance in generic parallel computing. its drawback is mostly in the _ease-of-use_ standpoint (simplicity, expressivity and programmability).

#### Dynamic Taskwing - subflow

**TaskFlow**, a task that is compsed of inner tasking, its a new task dependecy graph that is spawned in exectuiton of another task.

```cpp
tf::task A = tf.emplace([](){}).name("A");
tf::task C = tf.emplace([](){}).name("C");
tf::task D = tf.emplace([](){}).name("D");


// create a subflow graph (dynamic Tasking)
tf::Task B = tf.emplace([](tf::Subflow& subflow)
{
  tf::task B1 = subflow.emplace([](){}).name("B1");
  tf::task B2 = subflow.emplace([](){}).name("B2");
  tf::task B3 = subflow.emplace([](){}).name("B3");

  B1.preceded(B3);
  B2.preceded(B3);
}).name("B");

A.precede(B); //B runs after A
A.precede(C); //C runs after A
B.precede(D); //D runs after B
C.precede(D); //D runs after C
```

dynamic subflow can be nested of recursive, like fibonacci number computation, each task spawns a new subflow.

#### Heterogenous Tasking (cudaFlow) - offloading a task to GPU

> Single percision AX + Y ("SAXPY")
>
> - get x and y vectors on CPU (`allocate_x`, `allocate_y`)
> - copy x and y to GPU (`g2d_x`, `h2d_y`)
> - run saxpy kernel on x and y (`saxpy kernel`)
> - copy x and y back to CPU (`d2h_x`,`d2h_y`)

we

```cpp
const unsigned N = 1<<20;
std::vector<float> hx(N,1.0f),hy(N,2.0f);
float *dx{nullptr}, *dy{nullptr};
auto allocate_x = taskflow.emplace([&](){cudaMalloc(&dx,4*N);});
auto allocate_y = taskflow.emplace([&](){cudaMalloc(&dy,4*N);});

auto cudaflow = taskflow.emplace([&](tf::cudaFlow& tf)
{
  auto h2d_x = cd.copy(dx, hx.data(),N); // CPU-GPU data Transfer
  auto h2d_y = cd.copy(dy, hy.data(),N);

  auto d2h_x = cf.copy(hx.data(),dx,N); // GPU-CPU data Transfer
  auto d2h_y = cf.copy(hy.data(),dy,N);

  auto kernel= cf.kernel((N+255)/256,256,0,saxpy,N, dx,dy);
  kernel.succeed(h2d_x,h2d_y).precede(d2h_x,d2h_y);
});

cudaflow.succeed(allocate_x, allocate_y);
executor.run(taskflow).wait();
```

> Key motivations
>
> - Our closure enables stateful interface
>   - user capture data in reference to marshal data exchange between CPU and GPU tasks.
> - Our closure hides implementation details judiciously
>   - we use cudaGraph (since cuda 10) due to it's excellent performance, much faster than streams in large graphs.
> - Our closure extendes to new accelerato types
>   - syclFlow, openclFlow, coralFlow, tpuFlow, fpgaFlow, etc...)
>
> "We do not simplify kernel programming but **focus on CPU-GPU tasking that affects the performance to a large extent!** (same for data abstraction)

(something about keeping things visible to make performance better)

also **syclFlow**, logic as C++ code, associated with a SYCL queue.

```cpp
auto syclflow = taskflow.emplace([&](tf::syclFlow& sf)
{
  auto h2d_x = cd.copy(dx, hx.data(),N); // CPU-GPU data Transfer
  auto h2d_y = cd.copy(dy, hy.data(),N);

  auto d2h_x = cf.copy(hx.data(),dx,N); // GPU-CPU data Transfer
  auto d2h_y = cf.copy(hy.data(),dy,N);

  auto kernel = cf.parallel_for(sycl::range<1>(N),[=](sycl::id<1> id){
    dx[id]=2.0f * dx[id]+dy[id];
  });

  kernel.succeed(h2d_x,h2d_y).precede(d2h_x,d2h_y);
},queue); // create a syclFlow form a SYCL queue on a SYCL device
```

#### Conditional Tsking

> condition tasks intgrate control flow into a task graph to form end-to-end parallelism.

when a task return zero, it's a conditional task.

simple `if-else` task

```cpp
auto init = taskflow.emplace([&]() {initialize_data_structure();}).name("init");
auto optimizer = taskflow.emplace([&]() {matrix_solver();}).name("optimizer");
auto converged = taskflow.emplace([&]() {return converged() ? 1 : 0;}).name("converged"); // conditional task
auto output = taskflow.emplace([&]() {std::cout<<"Done!\n";}).name("output");

init.precede(optimizer);
optimizer.precede(converged);
converged.precede(optimizer, output); // return 0 to the optimizer again
```

`while/for loop` task. some our static task and some a conditional tasks.

```cpp
tf::Taskflow taskflow;
int i;
auto [init,cond, body, back, done] = taskflow.emplace(
  [&]() {std::cout<< "i=0"; i=0;},
  [&]() {std::cout<< "while i<5\n"; return i<5 ? 0 : 1;},
  [&]() {std::cout<< "i++=" << i++ <<'\n';},
  [&]() {std::cout<< "back\n"; return 0;},
  [&]() {std::cout<< "done\n";}
);
init.precede(cond);
cond.precede(body, done);
body.precede(back); // increment i
back.precede(cond); // not actually needed in this case, depending on application
```

non deterministic loops. flip 5 coins until you get 5 heads.

```cpp
auto A = taskflow.emplace([&](){});

auto B = taskflow.emplace([&](){ return rand()%2; });
auto C = taskflow.emplace([&](){ return rand()%2; });
auto D = taskflow.emplace([&](){ return rand()%2; });
auto E = taskflow.emplace([&](){ return rand()%2; });
auto F = taskflow.emplace([&](){ return rand()%2; });
auto G = taskflow.emplace([&](){});

A.precede(B).name("init");
// on zero it goes to the next task, on 1 it goes back to task B
B.precede(C,B).name("flip-coin-1");
C.precede(D,B).name("flip-coin-2");
D.precede(E,B).name("flip-coin-3");
E.precede(F,B).name("flip-coin-4");
F.precede(G,B).name("flip-coin-5");
G.name("end");
```

switch case, task as a switch statement, the return value specifies which task is used.

```cpp
auto [source, swichCond, case1,case2,case3,target] = taskflow.emplace(
[](){ std::cout << "source\n"; },
[](){ std::cout << "switch\n"; return rand()%3; },
[](){ std::cout << "case1\n"; return 0; },
[](){ std::cout << "case2\n"; return 0; },
[](){ std::cout << "case3\n"; return 0; },
[](){ std::cout << "target\n" ;}
);
source.preced(swichCond);
swichCond.precede(case1,case2,case3);
target.succeed(case1,case2,case3);
```

> "Existing frameworks on expressing conditional tasking or dynamic control flow suffer from _exponential growth_ of code complexcity."

#### Composable tasking

allowing compsability, so it's easier to optimize them.

```cpp
tf::Taskflow f1,f2;

auto [f1A, f1B]= f1.emplace(
[]() { std::cout << "Task f1A\n";},
[]() { std::cout << "Task f1B\n";}
);


auto [f2A, f2B,f2C]= f2.emplace(
[]() { std::cout << "Task f2A\n";},
[]() { std::cout << "Task f2B\n";},
[]() { std::cout << "Task f2C\n";}
);
auto f1_module_task = f2.composed_of(f1);
f1_module_task.succeed(f2A,f2B).precede(f2C);
```

> in Taskflow library, everthing is unified
>
> - Use `emplace` to create a task
> - Use `precede` to add a task dependency
> - No need to learned different sets of API
> - You can create a really complex graph
>   - Subflow(ConditionTask(cudaFlow))
>   - ConditionTask(StaticTask(cudaFlow))
>   - Composition(Subflow(ConditionTask))
>   - Subflow(ConditionTask(cudaFlow))
>   - ...
> - Scheduler performs end-to-end optimization
>   - Runtime, energy efficiency and throughput

**K means Clustering Example**

> - one cudaFlow for host-to-device data transfers.
> - one cudaFlow for finding the _k_ centroids.
> - one condition task to model iteration.
> - one cudaFlow for device-to-host data transfers.

### Understand our Scheduing Algorithm

The executor manages a set of threads to run taskflows:

- all execution methos are _non-blocking_
- all execution methos are _thread-safe_

submit taskflow to executor.

```cpp
tf::Taskflow taskflow1,taskflow2, taskflow3;
tf::Executor executor;

// create tasks and dependencies

auto future1 = executor.run(taskflow1);
auto future2 = executor.run_2(taskflow2,1000);
auto future3 = executor.run_until(taskflow3,[1=0](){return i++>5; });

executor.async([]() {std::cout << "async task\n"; });
executor.wait_for_all(); // wait fo all the above task to finish
```

> Task level scheduling: decide how tasks are enqueued under control flow.
>
> - ensure a feasible path to carry out control flow
> - avoids task race under cyclic and conditional execution
> - maximizes the capability of conditional tasking
>
> Worker level scheduling: decide how task are executed by which workers.
>
> - adopts work stealing to dynamically balance load
> - adapts workers to available task parallelism
> - maximes performance, energy and throughput

strong dependency and weak dependency. a weak dependency is "dependency" on a conditional task. a scheduler must start with a task without any depencides,

it's easy to make mistakes with condition tasks. like starting with a conditional task, or having race conditions when combining strong a weak dependencies.

work stealing allows threads to improve performance thorugh dynamic load balancing.

### Boost Performance in Real Applications

real life use cases:

one example is optimizing cell locations on a chip (VLSI placement).dynamic control flow, cudaFlow tasks, conditional cycle and static tasks.\
comparing against other libraries shows that taskFlow performs better in many case, and never worse.

another example is a machine learning, computing a neural network.

takeaway:

> "Different model give different implementations. the parallel code/algorithm may run fast, ye the parallel computing infrastructure to soupport that algorithm may dominate the entrire performance."

> Taskflow enables end-to-end expresion of CPU-GPU dependant tasks along with algorithmic control flow.

### Make C++ Amenable to Heterogenous Parallelism

parallelism is never standalone, it's a tool that we use to apply to our programs, no one tool could express all parallelisms.

> - C++ parallelism is primitive (but in a good shape)
>   - _std::thread_ is powerfull but very low level
>   - _std::async_ leaves off handling task dependencies
>   - No easy way to describe control flow in parallelism
>     - C++17 parallel STL count on bulk synchronous parallelism
> - C++ has no standard way to offload tasks to accelerators

</details>

##

[Main](README.md)
