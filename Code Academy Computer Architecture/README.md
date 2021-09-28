<!--
ignore these words in spell check for this file
// cSpell:ignore
-->

# Computer Architecture

[code academy](https://www.codecademy.com/learn/computer-architecture)

## Introduction

### How computers work

- input
- processing
- memory
- output

primary and secondary memory.

### Binary numbering system

binary numbers are just numbers represented as binary base.
binary data is data in binary form, like machine code, boolean expressions, hardware states, networking and file storage.

binary data is basically power on (1) and power off (0), it's all about currents and states.\
eight bits are a _byte_.\
two bytes (18 bits) are a _word_.\
4 bits are a _nibble_.

storage uses bytes (kilobytes, megabytes), while networking uses bits(kilobit, megabit), so we need to be careful with the units.

binary numbers: using base 2 (rather than base 10 for decimal numbers)\
MSB: most significant bit.(the left most)\
LSM: least significant bit. (the right most)

when we count in binary, we use the power of 2, if we have n digits, the highest number we can represent is $2^{n}-1$.
odd numbers end with the lsb as 1 (on) and even number have the lsb as 0(off).

```python
answer1 = (2**13)-1
answer2a = 31
answer2b = 2**15-1
#0b101111
num = int('01110010011', 2)
answer3a = 0 #num msb
answer3b=1 #num lsb
```

converting from binary to decimal

```python
decimal_conversion1 = int('100110',2)
decimal_conversion2 = int('1111011110011',2)
print(decimal_conversion1,decimal_conversion2)
```

we can convert from decimal to binart by dividing by 2 and taking the reminder (module), the reminder is then put on the binary representation, until the number itself is zero.
example with 27

> - 27 % 2 = 1, 27 / 2 =13 0b1
> - 13 % 2 = 1, 13 / 2 = 6 0b11
> - 6 % 2 = 0, 6 / 2 = 3 0b011
> - 3 % 2 = 1, 3 / 2 = 1 0b1011
> - 1 % 2 = 1, 1 / 2 = 0 0b11011
>
> 27 in binary form is 0b11011

adding binary numbers together is all about carrying the 1 bits upwards

> 1+0=1\
> 1+1=10\
> 1+1+1=11

substracting binary numbers

> 0b11010 -0b11 = 0b10111

multiplication is taking each bit of the small number, and it with each bit of the large number, and shift the result by the bit position, then add everything together (horrible explainnation)

division

## Instruction Set

## Assembly

## Cache

## Instruction Parallelism

## Data-level Parallelism
