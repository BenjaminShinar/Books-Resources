<!--
ignore these words in spell check for this file
// cSpell:ignore nand elif
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

multiplication is taking each bit of the small number, binary and (&) it with each bit of the large number, and shift the result by the bit position, then add everything together (horrible explainnation).

division
doing long division

(can't figure this one out)

### Logic gates: Voltage and bits

- nand - not and
- and
- or
- xor
- not

truth tables

| NAND | a     | b     | output |
| ---- | ----- | ----- | ------ |
|      | true  | true  | false  |
|      | true  | false | true   |
|      | false | true  | true   |
|      | false | false | true   |

| NOT | a     | b   | output |
| --- | ----- | --- | ------ |
|     | true  |     | false  |
|     | false |     | true   |

| AND | a     | b     | output |
| --- | ----- | ----- | ------ |
|     | true  | true  | true   |
|     | true  | false | false  |
|     | false | true  | false  |
|     | false | false | false  |

| OR  | a     | b     | output |
| --- | ----- | ----- | ------ |
|     | true  | true  | true   |
|     | true  | false | true   |
|     | false | true  | true   |
|     | false | false | false  |

| XOR | a     | b     | output |
| --- | ----- | ----- | ------ |
|     | true  | true  | false  |
|     | true  | false | true   |
|     | false | true  | true   |
|     | false | false | false  |

we can pass the same value to a nand gate and it will act as not gate\

not(a) == nand(a,a)\
and(a,b) == not(nand(a,b))\
or(a,b) == nand(not(a),not(b))\
xor = and (nand(a,b),or(a,b))\

### Creating a Circuit Adder

ALU - arithemetic logic unit
the adder part, is two half adders, half adders take two input, and return a sum bit and a carry bit

| Half Adder | a   | b   | Output -> Sum bit | Output -> Carry bit |
| ---------- | --- | --- | ----------------- | ------------------- |
|            | 1   | 1   | 0                 | 1                   |
|            | 1   | 0   | 1                 | 0                   |
|            | 0   | 1   | 1                 | 0                   |
|            | 0   | 0   | 0                 | 0                   |

| Full Adder | a   | b   | Carry-in bit | Output -> Sum bit | Output-> Carry-out bit |
| ---------- | --- | --- | ------------ | ----------------- | ---------------------- |
|            | 1   | 1   | 1            | 1                 | 1                      |
|            | 1   | 1   | 0            | 0                 | 1                      |
|            | 1   | 0   | 1            | 0                 | 1                      |
|            | 1   | 0   | 0            | 1                 | 0                      |
|            | 0   | 1   | 1            | 0                 | 1                      |
|            | 0   | 1   | 0            | 1                 | 0                      |
|            | 0   | 0   | 1            | 1                 | 0                      |
|            | 0   | 0   | 0            | 0                 | 0                      |

trying to make an ALU

```python
from nand import NAND_gate
from not_gate import NOT_gate
from and_gate import AND_gate
from or_gate import OR_gate
from xor_gate import XOR_gate

def half_adder(a,b):
  s = XOR_gate(a,b)
  c = AND_gate(a,b)
  return (s,c)

print(half_adder(1,1),"half adder expected(0,1)")
print(half_adder(1,0),"half adder expected(1,0)")
print(half_adder(0,1),"half adder expected(1,0)")
print(half_adder(0,0),"half adder expected(0,0)")

def full_adder(a,b,c):
  x,y =half_adder(a,b)
  s,c2=half_adder(x,c)
  c_out= OR_gate(y,c2)
  return(s,c_out)

print(full_adder(1,1,1),"full adder expected(1,1)")
print(full_adder(1,1,0),"full adder expected(0,1)")
print(full_adder(1,0,1),"full adder expected(0,1)")
print(full_adder(1,0,0),"full adder expected(1,0)")
print(full_adder(0,1,1),"full adder expected(0,1)")
print(full_adder(0,1,0),"full adder expected(1,0)")
print(full_adder(0,0,1),"full adder expected(1,0)")
print(full_adder(0,0,0),"full adder expected(0,0)")

def ALU(a,b,c,opcode):
  if (opcode ==0):
    return half_adder(a,b)
  elif(opcode ==1):
    return full_adder(a,b,c)

print(ALU(0,0,1,0),"ALU expected(0,0)")
print(ALU(0,0,1,1),"ALU expected(1,0)")
print(ALU(1,1,1,0),"ALU expected(0,1)")
print(ALU(1,1,1,1),"ALU expected(1,1)")
```

## Instruction Set

## Assembly

## Cache

## Instruction Parallelism

## Data-level Parallelism
