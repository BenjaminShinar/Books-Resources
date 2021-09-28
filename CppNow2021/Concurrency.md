<!--
// cSpell:ignore simd Steagall intrinsics cstdio immintrin loadu mmask storeu permutexvar permutex2var mmsetr maskz fmadd
 -->

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
    #define KEWB_FORECE_INLINE inline __attribute__((__always_inline__))
#else
    #define __OPTIMIZE__
    #include <immintrin.h>
    #undef __OPTIMIZE__
    #define KEWB_FORECE_INLINE inline
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
KEWB_FORECE_INLINE rf_512 load_value(float v)
{
    return _mm512_set1_ps(v);
}

KEWB_FORECE_INLINE ri_512 load_value(int32_t i)
{
    return _mm512_set1_epi32(i);
}

KEWB_FORECE_INLINE rf_512 load_from(float const * ptr_float)
{
    return _mm512_loadu_ps(ptr_float);
}

KEWB_FORECE_INLINE ri_512 load_from(float const * ptr_int)
{
    return _mm512_loadu_epi32i(ptr_int);
}

KEWB_FORECE_INLINE rf_512 masked_load_from(float const * ptr_float,rf_512 fill, msk_512 mask)
{
    return _mm512_mask_loadu_ps(fill,(__mmask16) mask,ptr_float);
}

KEWB_FORECE_INLINE rf_512 masked_load_from(float const * ptr_float,float fill, msk_512 mask)
{
    return _mm512_mask_loadu_ps(_mm512_set1_ps(fill),(__mmask16) mask,ptr_float);
}

KEWB_FORECE_INLINE void store_to(float * ptr_destination,rf_512 r)
{
    _mm512_storeu_ps(ptr_destination,r)
}

KEWB_FORECE_INLINE void store_to(float * ptr_destination,rf_512 r,msk_512 mask)
{
    _mm512_mask_storeu_ps(ptr_destination,(__mmask16)mask,r)
}

template <unsigned A = 0,....,unsigned P =0>
KEWB_FORECE_INLINE constexpr uint32_t make_bit_mask()
{
    //.. to much code for me to write, maybe I could use a folding expression here...
}

KEWB_FORECE_INLINE rf_512 blend(rf_512 a,rf_512 b,msk_512 mask)
{
    return _mm512_mask_blend_ps((__mmask16)mask,a,b);
}

KEWB_FORECE_INLINE rf_512 permute(rf_512 r,ri_512 perm)
{
    return _mm512_permutexvar_ps(perm,r);
}

KEWB_FORECE_INLINE rf_512 masked_permute(rf_512 a,rf_512 b,ri_512 perm,msk_512 mask)
{
    return _mm512_mask_permutexvar_ps(a,(__mmask16)mask,prem,b);
}

template <unsigned A,....,unsigned P>
KEWB_FORECE_INLINE constexpr ri_512 make_perm_mask()
{
    //static assert
    retrun _mmsetr_epi32(A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P)
}

template<int R>
KEWB_FORECE_INLINE rf_512 rotate(rf_512 r)
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
KEWB_FORECE_INLINE rf_512 rotate_down(rf_512 r)
{
    static_assert(R >= 0)
    return rotate<-R>(r);
}

template<int R>
KEWB_FORECE_INLINE rf_512 rotate_up(rf_512 r)
{
    static_assert(R >= 0)
    return rotate<R>(r);
}

template<int S>
KEWB_FORECE_INLINE rf_512 shift_down(rf_512 r)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_down<S>(r),load_value(0.0f), shift_down_blend_mask<S>());
}

template<int S>
KEWB_FORECE_INLINE rf_512 shift_up(rf_512 r)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_up<S>(r),load_value(0.0f), shift_up_blend_mask<S>());
}

template<int S>
KEWB_FORECE_INLINE rf_512 shift_down_with_carry(rf_512 a,ref_512 b)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_down<S>(a),rotate_down<S>(b), shift_down_blend_mask<S>());
}

template<int S>
KEWB_FORECE_INLINE rf_512 shift_up_with_carry(rf_512 a,ref_512 b)
{
    static_assert(S >= 0 && S<=16)
    return blend(rotate_up<S>(a),rotate_up<S>(b), shift_up_blend_mask<S>());
}

template<int S>
KEWB_FORECE_INLINE void in_place_shift_down_with_carry(rf_512 &a,ref_512 &b)
{
    static_assert(S >= 0 && S<=16)
    constexpr msk_512 z_mask = (0xFFFFu >> (unsigned)S);
    constexpr msk_512 b_mask = ~z_mask & 0xFFFFU;
    ri_512 perm = make_shift_permutations<S,b_mask> ()
    a = _mm512_permutex2var_ps(a, perm,b);
    b = _mm512_maskz_permutex2var_ps((__mmask16)z_mask,b,perm,b)
}

KEWB_FORECE_INLINE rf_512 add(rf_512 a,ref_512 b)
{
    return _mm512_add_ps(a,b);
}

KEWB_FORECE_INLINE rf_512 sub(rf_512 a,ref_512 b)
{
    return _mm512_sub_ps(a,b);
}

KEWB_FORECE_INLINE rf_512 minimum(rf_512 a,ref_512 b)
{
    return _mm512_min_ps(a,b);
}
KEWB_FORECE_INLINE rf_512 maximum(rf_512 a,ref_512 b)
{
    return _mm512_max_ps(a,b);
}
```

now lets build some functions that use those building blocks

### Intra-register Sorting with Sorting networks.

- _compare_with_exchange_ - usefull for sorting, we can sort pairs of positions.

```cpp
KEWB_FORECE_INLINE rf_512 compare_with_exchange(rf_512 vals, ri_512 perm, msk_512 mask)
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
KEWB_FORECE_INLINE rf_512 sort_two_lanes_of_8(rf_512 vals)
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
