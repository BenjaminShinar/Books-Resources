<!--
ignore these words in spell check for this file
// cSpell:ignore nand elif shmat
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

### Intro

ISA - instruction set architecture, connects between the software and the hardware.\
CPU - central processing unit.\
CISC - complex instruction set. many instructions that do fairly complex actions load+action+store,etc..
RISC - reduced instruction set. less instructions, pipelined into one another, less power consumption

designs:

- X86(CISC)
- ARM(RISC)
- MIPS(modified RISC for embedded processors)

and something entirely different once quantum computers get going.

### Instruction Set architecture

the ISA is what connects the software and the hardware,

> - User programs
> - High level languages
> - Compiler
> - Assembly language
> - Instruction set architecture
> - Computer hardware

The cpu has three components:

- control unit (CU)
- Arithmetic and logic unit (ALU)
- Registers (immediate access store)

the control unit is the overseer of the cpu, monitoring input and output. the alu does all the processing, even chaning single pixels. the registers are used for storing data for immediate access. everthing is done via connections (wires, conducters, semi-conducters) that transmit energy as binary state (On for 1, Off for 0).

the control unit has an inner clock, which it uses to send a electronic signal to the other components to signall them to run at the same time.

the alu does the calculations both the arithmetic and the logical operations.

data for immediate access is stored in the registers, a cpu can have a different number of registers, with different sizes.

other than the cpu, the is also RAM (random access memory) and IO (input output) components. the ram is used for short term actions, it's volatile, so it loses all information when power is closed.

signals are transferred over 3 lanes called busses, a _bus_ is a job specific high-speed wire, usually grouped together in bundles to deliver data (in serial or parallel form).

- control bus, cpu --> ram, cpu --> IO componenets. (unidirectional). control busses carry control signals and clock signals.
- address bus, cpu --> ram, cpu --> IO componenets (unidirectional). address buses carry specific address data.
- data bus, cpu <--> ram, cpu <--> IO componenets. (bidirectional). data buses carry data, all sorts of data.

hard disk (hard drives) are long term storage, they are none-volatile and they retain their state even without power.

Machine instruction are specific, pre-determined packages of data that the hardware knows how to handle, RISC use machine code that is all the same length, while CISC instructions have varying lengths. this machine instructions are how we tell the computer what do to, if we send the wrong instruction (or something that doesn't exists), we get garbage.

#### OPCODE

the instruction length might be different for CISC and RISC, but they share some commonalities.\
the first few bits are the OPCode (operation code), which is the way of telling the processor what type of instruction is being received.

> opcode instructions
>
> | Name     | OPCODE | Formal Defintion                | Description                                                             |
> | -------- | ------ | ------------------------------- | ----------------------------------------------------------------------- |
> | ADD      | 000001 | rs_reg <- op_reg_1 + op_reg_2   | Loads two numbers from registers and saves result into another register |
> | SUBTRACT | 000010 |                                 |                                                                         |
> | MULTIPLY | 000011 |                                 |                                                                         |
> | DIVIDE   | 000100 |                                 |                                                                         |
> | LOAD     | 000101 | rs_reg <- mem\[op_reg_1_addr]   | Loads a number from a memory address location into a register           |
> | STORE    | 000110 | mem\[op_reg_1_addr] <- op_reg_2 | Copies data in a register to a memory address for long-term storage     |

after the OPCODE the remaining bits are usually called 'operands', they can be addresses, literal value or either pieces of data.

#### Instruction Formatting

we know that the first part of the instruction is the opcode. and the resets is the opeards, memory locations and addition functionality for the processor.

CISC code is long, because the goal was to reduce the total number of instructions that were fed to the hardware, even if each instruction took longer to process.\
RISC code is short and broken up,there are more tasks to complete, but each of them was shorter, this could then be used to pipe instructions in sequences to achieve the same results.

[MIPS](https://www.mips.com/products/architectures/mips32-2/) - micro pressor with interlocked pipeline stages uses a fixed 32 bit instruction length.

#### MIPS instructions

the mips isa is broken into three types of instructions:

- R-type (_register_) - for arithmetic and logical operations
- I-type (_immediate_) - for data transfer and operations with constants
- J-type (_jump_) - for flow control, like loops, branches

the mips isa also requires the cpu to hold 32 registers, each holding 32-bit data. data is stored either in one of those registers or is encoded on 16 bit data an _immediate_ (or constant) that doesn't need to take space in a register. mips is used in embedded systems, as low energy consumption is important for small systems.

R-type instructions are formatted as such:

- op (6 bits) - opcode
- rs (5 bits) - "first source register"
- rt (5 bits) - "second source register
- rd (5 bits) - "destination register"
- shmat (5 bits) - "bit shift amount"
- func (6 bits) - "extra bits for additional functions"

R-type are the most common type of instructions in MIPS. all R-types have opcode of 000000, so the processor look at the func part to determine which procss to execute. the register bits can range from 0 to 31 ((5^2)-1 = 31), which indicate which of the 32 registers are used. 'rt' and 'rs' are the operands, and 'rd' is the destination. the 'shmat' is the bit shift amount, the 'func' is the function to perform.

> "000000 00101 10010 00110 00000 100000"\
> op=0 rs=5 rt=18 rd=6 shmat=0 func=16 (add)\
> "take whats in registers 5 and 18, add together and store in register 6"

register zero is a protected register.

### Ultra Super Calculation Computer (project)

a project to do an ISA with five functionalities: add,substracy, multiply, divide, display history.\
also needed

> - Read and split up our incoming data
> - Store a binary number to a register
> - Access what is stored in the register
> - Allocate some registers for a 'history' of our calculations
> - Store/Load from the history when needed

see python code in file 'ultrasupercalculationcomputer'

## Assembly

## Cache

## Instruction Parallelism

## Data-level Parallelism
