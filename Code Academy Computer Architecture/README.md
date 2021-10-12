<!--
ignore these words in spell check for this file
// cSpell:ignore nand elif shmat addiu mult mflo fifo_indicies
-->

# Computer Architecture

[code academy](https://www.codecademy.com/learn/computer-architecture)

(I wouldn't trust the python code to run outside of code academy platform)

## Introduction

<details>
<summary>
TODO
</summary>

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

</details>

## Instruction Set

<details>
<summary>
TODO
</summary>

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

</details>

## Assembly

<details>
<summary>
TODO
</summary>

The code we write in c, python, js or any other language is the source code, but this isn't what the machine can execute. the instructions to the cpu need to simplified, and that's the assembly language.

### Assembly language

Assembly language is directly translated instructions, it's human readable, but still the same.

most of the assembly language was abstracted away from most programming languages, it's hidden behind layers of high-level programming. however, it is still used in embedded programming, direct hardware testing and software optimizations.

embedded systems have low memory and storage, so assembly allows us to manually optimize and control each task separately to ensure that we don't overtax the hardware.

assembly languages vary by the ISA used, MIPS, X86 (CISC) and ARM(RISC) and by vendor.

#### Compilation

when we write source code, we need to eventually turn it into machine instructions, this is achieved by following the four stages of software:

1. preprocessing - removing comments, expanding macro, etc...
2. compiling - turning into assembly code
3. assembling - turning the assembly code into machine instructions.
4. linking - filling in the blanks, like locations, addresses and additional libraries.

#### Assembly code format

assembly are strongly related to the machine instructions (nearly identical):\
assembly begins with an opcode, but rather than six bits, it's one of the predefined words.

this assembly code multiplies the data in register 3 by that of the data in register 2.

```MIPS
MULT $3,$2
```

is the same as the machine code

```
00000000111001100000000000011000
000000  00111 00110 00000 00000 011000
op 0 rs 7 rt 6 rd 0 shmat 0 func 24
```

the $ symbol in MIPS means direct register addresses.as before we have opcode and operands.

#### Arithmetic Operations

most of the stuff the cpu does is arithmetics, but different arithmetic operations depend on how the numbers are stored at hardware, registers, cache and other types of storage have fixed binary lengths, so there needs to be somewhere to store 'overflow' from operations, we can have operations that act on two values from registers or one from register and one constant (and immediate).\
for example ADD takes three arguments, two register to take the value from and one to store it, ADDI (add immediate) takes one register of data, a register to store the result, and the immediate constant (the order is different!)

```mips
ADD $4,$5, $6
ADDI $4,$6, 7
```

other common arithmetics operations are SUB,SUBI, MULT,MULTII,DIV,DIVI.

#### Memory Access Operations

we can control where information is stored, we can store it in a register for immediate use or push it back into a different memory storage location. the commands are LW and SW: load word, store word. a "word" is a fixed size data, usually 32 bit.

load what's inside register 3 (indirect accesses) into register 1.\
add the constant 15 to what's inside register 1 and store the result inside register 1.\
store what inside register 1 into where register 3 points to ((indirect accesses)
XOR with self is basically resting the register back to zero

```mips
LW $1,($3)
ADDI $1,$1, 15
SW $1, ($3)
XOR $1,$1,$1
```

in assembly coding, there are no variables, and the programmer must keep track of everything.

#### Control Flow Operations

we get some branching and stuff for control, we can also jump directly to memory locations

- BEQ (branch on equal ==)
- BNE (branch on not equal != )
- BGTZ, BGEZ (branch on greater than zero, branch on greater or equal to zero >0, >= 0)
- BLTZ, BLEZ (branch on less than zero, branch on less or equal to zero <0, <= 0)
- J (jump to location)

#### Memory Addressing, Direct and Indirect

the parentheses in some of the code aren't just for show, they can mean direct and indirect reference. we can use the registers to store the address of other pieces of memory, and then read from that.
direct access takes the data in the register, indirect access (with parentheses) uses that value to read from a different memory locations

in our example, register 5 has the value `0b1101000111` (839 decimal), and somewhere we have memory with that adderess that contains `0b10001110001112` (4551 decimal).

```mips
ADD $5,$5,$6
```

now registers six has the result of adding 839 with 839.

however

```mips
LW $4,($5)
ADD $5,$4,$6
```

we first load the indirect value from the address in register 5 so now register 4 stores (4551 decimal), and then we add them together and store the result in register 6 (4551+839 =5390)

#### Translation between Assembly and Binary

there is nearly a one-to-one relation between assembly code and machine code.

trying to understand this code

```mips
  square:
     addiu $sp,$sp,8
     sw $fp,4($sp)
     move $fp,$sp
     sw $4,8($fp)
     lw $3,8($fp)
     lw $2,8($fp)
     nop
     mult $3,$2
     mflo $2
     move $sp,$fp
     lw $fp,4($sp)
     addiu $sp,$sp,8
     j $31
     nop
```

- ADDIU - add immediate unsigned word
- MOVE- move function pointer
- MFLO - move from lower register

my guess

1. add 8 to stack pointer
2. store what inside function into offset 4 from stack pointer
3. move fp to the stack pointer
4. load into register 3 the value which is offest 8 from stack pointer
5. load into register 2 the value which is offest 8 from stack pointer
6. do nothing
7. multiply register $3 and $2
8. move value into register 2
9. move stack pointer to function pointer
10. load into function pointer offset 4 from stack pointer
11. add 8 to stack pointer
12. jump to register 31 value
13. do nothing

(sure as hell Im not trying to write this myself)

| opcode | rt    | rs    | rd    | shmat | func   | assembly                 |
| ------ | ----- | ----- | ----- | ----- | ------ | ------------------------ |
| 001001 | 11101 | 11101 | 00000 | 00000 | 001000 | `addiu $sp,$sp,8 `       |
| 101011 | 11101 | 11110 | 00000 | 00000 | 000100 | `sw $fp,4($sp) `         |
| 010001 | 00000 | 00000 | 11101 | 11110 | 000110 | `move $fp,$sp`           |
| 101011 | 11110 | 00100 | 00000 | 00000 | 001000 | `sw $4,8($fp) `          |
| 100011 | 11110 | 00011 | 00000 | 00000 | 001000 | `lw $3,8($fp) `          |
| 100011 | 11110 | 00010 | 00000 | 00000 | 001000 | `lw $2,8($fp) `          |
| 000000 | 00000 | 00000 | 00000 | 00000 | 000000 | `nop `                   |
| 000000 | 00011 | 00010 | 00000 | 00000 | 011000 | `mult $3,$2 `            |
| 000000 | 00000 | 00000 | 00010 | 00000 | 010010 | `mflo $2 `               |
| 010001 | 00000 | 00000 | 11110 | 11101 | 000110 | `move $sp,$fp `          |
| 100011 | 11101 | 11110 | 00000 | 00000 | 000100 | `lw $fp,4($sp) `         |
| 001001 | 11101 | 11101 | 00000 | 00000 | 001000 | `add : addiu $sp,$sp,8 ` |
| 000010 | 00000 | 00000 | 00000 | 00000 | 011111 | `j $31 `                 |
| 000000 | 00000 | 00000 | 00000 | 00000 | 000000 | `nop `                   |

### Assembly language problem Set

some question with assembly that are ridiculous to think of

00000000000000000000000000101001 \* 00000000000000000000000111100111 == \
(41\*423) \
00000000000000000000000111100111 + \
00000000000000000000111100111000 + \
00000000000000000011110011100000 == \
################################ \
00000000000000000100110111111111 == \
19967

</details>

## Cache

<details>
<summary>
TODO
</summary>

### Introduction to Cache

Cache memory, minimizing the delay when accessing the memory.

DDR: Double Data Rate Synchronous Dynamic Random Access

#### Memory Hierarchy

performance speed (processing data) grew faster tha memory speed (accessing the data), so getting data became a bottle neck. the performance of memory operations decreases (get slower) the larger the memory is. the registers store only a few bytes each, and are very fast, while the hdd can store terrabytes of data, and is very slow. between them we have the _cache_ and the _RAM_.

```python
from isa import ISA
from memory import Memory, MainMemory

if __name__ == "__main__":
  cache_arch = ISA()
  # Write your code below
  cache_arch.set_memory(MainMemory())


  # Architecture runs the instructions
  cache_arch.read_instructions("ex2_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

#### Cache Memory

The cache holds more memory than the processor (registers), and less memory than the main memory (RAM). the speed is also between them.
the cache is composed out of blocks, each blocks stores a copy of the memory from the main data, and is paired with a 'tag' which is the address of the data is the main memory, so the same address are used for both the cache and the main data.
together, each pair of tag and data are called 'entries'.

```python
from isa import ISA
from memory import Memory, MainMemory

class Cache(Memory):
  def __init__(self):
    # Write your code below
    super().__init__(name="Cache", access_time=0.5)
    self.data = [
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""}
    ]



  # Returns the cache total execution time
  def get_exec_time(self):
    return self.exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  # Change the code below
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex3_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} seconds")
```

#### Cache Hit

when we request data from the cache, we check if it exists, if it does, it's called a 'Cache hit', this is good.

```python
from isa import ISA
from memory import Memory, MainMemory

class Cache(Memory):
  def __init__(self):
    super().__init__(name="Cache", access_time=0.5)

    self.data = [
      {"tag": 0, "data": "M"},
      {"tag": 1, "data": "i"},
      {"tag": 2, "data": "s"},
      {"tag": 3, "data": "p"},
    ]


  def read(self, address):
    super().read()
    data = None
    # Write your code below
    entry = self.get_entry(address)
    if entry is not None:
      data=entry["data"]
    return data

  # Returns entry on cache hit
  # Returns None on cache miss
  def get_entry(self, address):
    for entry in self.data:
      if entry["tag"] == address:
          print(f"HIT: ", end="")
          return entry

    print(f"MISS", end="")
    return None

  def get_exec_time(self):
    return self.exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex4_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

#### Cache Miss

when we request data that wasn't found in the cache, it's called a 'Cache miss', which means that we should access the main memory instead, the data is retrived from the main memory and stored in the cache, and is returned from there. we would ideally like as few cache misses as we can.

```python
from isa import ISA
from memory import Memory, MainMemory

class Cache(Memory):
  def __init__(self):
    super().__init__(name="Cache", access_time=0.5)
    self.main_memory = MainMemory()

    self.data = [
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
    ]

  def read(self, address):
    super().read()
    data = None
    entry = self.get_entry(address)
    if entry is not None:
      data = entry["data"]
    # Write your code below
    else:
      data = self.main_memory.read(address)
      self.add_entry(address,data)


    return data
  # Adds data to an empty entry
  def add_entry(self, address, data):
    for entry in self.data:
      if entry["tag"] == None:
        entry["tag"] = address
        entry["data"] = data
        return

  # Returns entry on cache hit
  # Returns None on cache miss
  def get_entry(self, address):
    for entry in self.data:
      if entry["tag"] == address:
          print(f"HIT: ", end="")
          return entry

    print(f"MISS", end="")
    return None

  def get_exec_time(self):
    exec_time = self.exec_time + self.main_memory.get_exec_time()
    return exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex5_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

#### Replament Policy

when we get a cache miss, we would want to add the entry to the cache, but what do we do when the cache is full? we need to decide which entries get overwritten. this is decided by the 'replacement policy' of the cache. this is part of the architecture design, simple policies can be:

- FIFO - first in, first out. keep an index of the last entry used, and rollover back to the start when finished.
- LRU - Least recently used. keep track of the entry which wasn't used the most time and replace it, the cost of calculating this might not be worth the benefits.
- Random Replacement. choose at random.

the goal is to increase the number of cache hits and limit cache misses.

```python
from isa import ISA
from memory import Memory, MainMemory
from random import randint

class Cache(Memory):
  def __init__(self):
    super().__init__(name="Cache", access_time=0.5)
    self.main_memory = MainMemory()
    self.fifo_index = 0

    self.data = [
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
    ]

  def read(self, address):
    super().read()
    data = None
    entry = self.get_entry(address)
    if entry is not None:
      data = entry["data"]
    else:
      data = self.main_memory.read(address)
      # Change the code below
      self.replace_entry(address, data)

    return data

  def replace_entry(self, address, data):
    # Change the code below
    index = self.fifo_policy()
    self.data[index] = {"tag": address, "data": data}

  def random_policy(self):
    return randint(0, len(self.data)-1)

  def fifo_policy(self):
    index = self.fifo_index
    self.fifo_index += 1
    if self.fifo_index == len(self.data):
      self.fifo_index = 0

    return index

  # Adds data in an empty entry
  def add_entry(self, address, data):
    for entry in self.data:
      if entry["tag"] == None:
        entry["tag"] = address
        entry["data"] = data
        return

  # Returns entry on cache hit
  # Returns None on cache miss
  def get_entry(self, address):
    for entry in self.data:
      if entry["tag"] == address:
          print(f"HIT: ", end="")
          return entry

    print(f"MISS", end="")
    return None

  def get_exec_time(self):
    exec_time = self.exec_time + self.main_memory.get_exec_time()
    return exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex6_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

#### Associativity

what id data from the main memory is placed in a specific location inside the cache?

- Fully Associative. any location on the main memory can go anywhere.
- Direct Mapped. a main memory location can only appear in a specific location, this overrides the replacement policy.
- N-Way Set Associative. ~~a main memory location is limited to a specific set of blocks in which it can appear. the 'n' determins how many possible locations are there. a 2 way set associative mapping means that a single main memory address can only appear in one of two locations in the cache, if it's not in one of those two, then it's a cache miss. replacement policy is required, but can only replace the data in those places.~~

~~if we have a cache with 32 locations, then Fully Associative can be described as 1-Way set associative, Direct-Mapped is 32-way set associative.~~

```python
from isa import ISA
from memory import Memory, MainMemory
from random import randint

class Cache(Memory):
  def __init__(self):
    super().__init__(name="Cache", access_time=0.5)
    self.main_memory = MainMemory()
    # Change the value below
    self.sets = 4 # Set to 1, 2 or 4
    self.fifo_indices = [0, None, None, None]

    # Sets self.fifo_indicies based on
    # the number of sets in the cache
    if self.sets == 2:
      self.fifo_indices = [0, 2, None, None]
    elif self.sets == 4:
      self.fifo_indices = [0, 1, 2, 3]

    self.data = [
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
    ]

  def read(self, address):
    super().read()
    data = None
    entry = self.get_entry(address)
    if entry is not None:
      data = entry["data"]
    else:
      data = self.main_memory.read(address)
      self.replace_entry(address, data)

    return data

  def replace_entry(self, address, data):
    index = 0
    # Write your code below
    set_number = address % self.sets
    index = self.fifo_policy(set_number)


    self.data[index] = {"tag": address, "data": data}

  def random_policy(self, set_number):
    if self.sets == 1:
      return randint(0, len(self.data)-1)
    elif self.sets == 2:
      return randint(set_number*2, set_number*2+1)

    return set_number

  def fifo_policy(self, set_number):
    index = self.fifo_indices[set_number]
    self.fifo_indices[set_number] += 1
    if self.fifo_indices[set_number] == len(self.data)/self.sets+(set_number*int(len(self.data)/self.sets)):
      self.fifo_indices[set_number] = set_number*int(len(self.data)/self.sets)

    return index

  # Returns entry on cache hit
  # Returns None on cache miss
  def get_entry(self, address):
    for entry in self.data:
      if entry["tag"] == address:
          print(f"HIT: ", end="")
          return entry

    print(f"MISS", end="")
    return None

  def get_exec_time(self):
    exec_time = self.exec_time + self.main_memory.get_exec_time()
    return exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex7_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

#### Writing Policy

eventually, we would want to write the data in the cache to the memory (so we could retrive it, or even write to the long term memory).

the decsion when to send data from the cache to the main memory is handled by the 'Write Policy'.

- Write-through. when data is written to the cache, it's also written to the main memory, easy to implement, but costly, as we require the slow process of writing to the main memory each time we change the cache memory.
- Write-back. the data is written to the main memory only when the entry is overwritten. so right before we lose the data in the cache, it's stored in the main memory.

```python
from isa import ISA
from memory import Memory, MainMemory
from random import randint

class Cache(Memory):
  def __init__(self):
    super().__init__(name="Cache", access_time=0.5)
    self.main_memory = MainMemory()
    self.fifo_indices = [0, 0, 0, 0]
    self.sets = 1 # Set to 1, 2 or 4
    self.fifo_indices = [0, None, None, None]

    if self.sets == 2:
      self.fifo_indices = [0, 2, None, None]
    elif self.sets == 4:
      self.fifo_indices = [0, 1, 2, 3]

    self.data = [
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
      {"tag": None, "data": ""},
    ]

  def write(self, address, data):
    super().write()
    entry = self.get_entry(address)
    # Write your code below
    if entry is not None:
      entry["data"]=data
    else:
      self.replace_entry(address,data)
    self.main_memory.write(address,data)


  def read(self, address):
    super().read()
    data = None
    entry = self.get_entry(address)
    if entry is not None:
      data = entry["data"]
    else:
      data = self.main_memory.read(address)
      self.replace_entry(address, data)

    return data

  def replace_entry(self, address, data):
    index = 0
    set_number = address % self.sets
    index = self.fifo_policy(set_number)
    self.data[index] = {"tag": address, "data": data}

  def random_policy(self, set_number):
    if self.sets == 1:
      return randint(0, len(self.data)-1)
    elif self.sets == 2:
      return randint(set_number*2, set_number*2+1)

    return set_number

  def fifo_policy(self, set_number):
    index = self.fifo_indices[set_number]
    self.fifo_indices[set_number] += 1
    if self.fifo_indices[set_number] == len(self.data)/self.sets+(set_number*int(len(self.data)/self.sets)):
      self.fifo_indices[set_number] = set_number*int(len(self.data)/self.sets)

    return self.fifo_indices[set_number]

  # Returns entry on cache hit
  # Returns None on cache miss
  def get_entry(self, address):
    for entry in self.data:
      if entry["tag"] == address and entry["data"] is not "":
          print(f"HIT: ", end="")
          return entry

    print(f"MISS", end="")
    return None

  def get_exec_time(self):
    exec_time = self.exec_time + self.main_memory.get_exec_time()
    return exec_time

if __name__ == "__main__":
  cache_arch = ISA()
  cache_arch.set_memory(Cache())

  # Architecture runs the instructions
  cache_arch.read_instructions("ex8_instructions")

  # This outputs the memory data and code execution time
  exec_time = cache_arch.get_exec_time()
  if exec_time > 0:
    print(f"OUTPUT STRING: {cache_arch.output}")
    print(f"EXECUTION TIME: {exec_time:.2f} nanoseconds")
```

### Cache Problem Set

#### Problem 1-A

four blocks cache, main memory of 16, reading the following data (address)
\[8,3,4,12,10,7,3,2,6,3,1,7,8,6], fifo replacement policy

1. how many cache misses?
2. how many cache hits?
3. what is the final state of the cache?

- 8 Miss \[8]
- 3 Miss \[8,3]
- 4 Miss \[8,3,4]
- 12 Miss \[8,3,4,12]
- 10 Miss \[**10**,3,4,12]
- 7 Miss \[10,**7**,4,12]
- 3 Miss \[10,7,**3**,12]
- 2 Miss \[10,7,3,**2**]
- 6 Miss \[**6**,7,3,2]
- 3 Hit \[6,7,3,2]
- 1 Miss \[6,**1**,3,2]
- 7 Miss \[6,1,**7**,2]
- 8 Miss \[6,1,7,**8**]
- 6 Hit \[6,1,7,8]

#### Problem 1-B

four blocks cache, main memory of 16, reading the following data (address)
\[8,3,4,12,10,7,3,2,6,3,1,7,8,6], fifo replacement policy, now with 2 set associativity, which alternates each time,

1. how many cache misses?
2. how many cache hits?
3. what is the final state of the cache?

(replace the sets, even on the left, odd on the right)

- 8 Miss \[, ,~ , ]
- 3 Miss \[3, ,~ 8,]
- 4 Miss \[3, ,~8,4]
- 12 Miss \[3, ,~**12**,4]
- 10 Miss \[3, ,~12,**10**]
- 7 Miss\[3,7~12,**10**]
- 3 Hit\[3,7 ~12,**10** ]
- 2 Miss\[3,7 ~**2**,10 ]
- 6 Miss \[3,7 ~2,**6** ]
- 3 Hit\[ 3,7~2,**6** ]
- 1 Miss\[**1**,7 ~2,**6** ]
- 7 Hit \[**1**,7 ~ 2,**6**]
- 8 Miss\[**1**,7 ~ **8**,6 ]
- 6 Hit\[**1**,7 ~ **8**,6]

#### Problem 1-C

four blocks cache, main memory of 16, reading the following data (address)
\[8,3,4,12,10,7,3,2,6,3,1,7,8,6], fifo replacement policy, direct mapped, which is like 4 way set associative (???)

1. how many cache misses?
2. how many cache hits?
3. what is the final state of the cache?

- 8 -> 0 miss
- 3 -> 3 miss
- 4 -> 0 miss
- 12 -> 0 miss
- 10 -> 2 miss
- 7 -> 3 miss
- 3 -> 3 miss
- 2 -> 2 miss
- 6 -> 2 miss
- 3 -> 3 hit
- 1 -> 1 hit
- 7 -> 3 miss
- 8 -> 0 miss
- 6 -> 2 hit

3 hits, 11 misses, \[8,1,6,7]

#### Problem 2-A

four blocks cache, fully associate, FIFO replacement policy, access to cache cost 0.5 ns and access to main memory is 30 ns. data is \[1,4,6,2,1,4,7,4,1,4,7,0].
write policy is write-through,

1. what is the total access time of cache writes
2. what is the total access time of the main memory writes
3. what is the total access time of both cache and main memory writes.

- 1 \[1,,,] 0.5
- 4 \[1,4,,] 0.5
- 6 \[1,4,6,] 0.5
- 2 \[1,4,6,2] 0.5
- 1 \[1,4,6,2] 0.5
- 4 \[1,4,6,2] 0.5
- 7 \[**7**,4,6,2] 0.5, 30
- 4 \[**7**,4,6,2] 0.5
- 1 \[7,**1**,6,2] 0.5, 30
- 4 \[7,1,**4**,2] 0.5, 30
- 7 \[7,1,**4**,6] 0.5
- 0 \[7,1,4,**0**] 0.5, 30

#### Problem 2-B

four blocks cache, fully associate, FIFO replacement policy, access to cache cost 0.5 ns and access to main memory is 30 ns. data is \[1,4,6,2,1,4,7,4,1,4,7,0].
write policy is write-back,

1. what is the total access time of cache writes
2. what is the total access time of the main memory writes
3. what is the total access time of both cache and main memory writes.

12 writes, each to the cache (0.5\*12 = 6) and the main memory (30\*12=360), total is 366.

</details>

## Instruction Parallelism

<details>
<summary>
TODO
</summary>

### The Instruction Cycle

a set of operations the cpu must to to execute a single instruction.also called the fetch-excute cycle. depending on the cpu, but usually consists of

- fetch
- decode
- execute
- memory access
- registry write-back

#### Fetching

using the **program counter register** (PC) and the **instruction register** (IR), the PC stores the memory address of the next instruction to run, when the fetch cycle starts, this value is copied into the IR for decoding.

#### Decoding

now the Control Unit decodes (deciphers) what the instruction in the IR is and what should it send to to which components, such as as the ALU or other hardware. a single instruction is turned into a series of control signals.

#### Executing

now the control signlas are sent to the correct part of the ALU for processing.

#### Memory Access

Sometimes we need to retrive data in order to perform an instruction, if we used immediate, then there's no need, but when we get data from registers, we do need to perform memory access (even go to the cache).

#### Registry Write Back

if our instruction requires to store the data in memory (not for use in immediate calculations), we need to perform a write to the registers.

#### Deli Example

the example they give is buying something in a deli.

- an instruction is a note with our order
- fetching is when someone looks at our order
- decoding is when he understands what we wanted
- executing is when the deli works on our order
- memory access is when the deli needs to open up a new jar of mustard for our order
- write back is when the deli writes our name as a favorite guest or something.

### Instruction Pipelining

rather then perform a single instructions at a time, the cpu can actually process multiple instructions at the same time. this is done by the hardware, pipelinning is the connecting tissue between hardware and software.

#### Linear Instructions

example of a laundrymat taking orders.

if we have to wait for each stage of the fetch-excute cycle to complete,we might have dead time waiting, once we decoded the instructions, there no reason to keep it in the IR register, so maybe we can start fetching the next instruction?

#### Pipelinning

in a none-digital world, like our laundrymat, we don't have to wait for all the clothes to be folded from one order before we start the next order (putting the clothes in the washing machine), the same is true for instructions. ideally, we would want to process them in parallel.

The cost of this pipelining is on the hardware, this means more operations, some the cpu is running more (and is hotter) and more complex (more expansive).

### Hazards of Instruction Pipelining

Pipelining is useful, but can also have problems. we might skip processing instructions in a cycle, we might have a 'pipeline flush' that causes us to lose all of the instructions currently in the pipeline.

#### Strutural hazards

limitations of the hardware itself, like when we need to access the RAM rather than the cache, which brings the speed down. the ALU can do only one instruction at a time, and some instructions are more demanding (division), so this too can create hazards on the pipeline.

#### Data hazards

a data hazard occurs when an instruction is dependent on another instruction still in the pipeline. if we have to finish the previous instruction before we can start the current one.

#### Control hazards

control hazards happen at branches (if,loops, virtual tables). if we don't know which operation will be next, we can't start processing it, the cpu takes a guess, but is sometimes wrong, and it has to do a pipeline flush and start again with the right instruction.

#### Reducing Hazards in Pipline

there is no one perfect way to remove all risks, we can try to find and limit them.

for data hazards, we can reduce memory read/write back chaining results from one instruction to another. we can reorder the instructions to reduce the risk of direct dependency, this is done by the processor.\
a last method is for the processor to create 'bubble', opeartions that take time to buffer between instructions that are dependent on one another.

for control hazards, processors can stall (wait until they know which instruction to run), or they can try predict which branch to take. for loops, they can simply unroll the loop into sequential commands, which are faster.

structural hazards can be mitigated by the design of the processor, like getting a better cache strategy.

### Superscaler Architecture

a strategy to run parallel process by having several execution context units.

#### What is a Superscaler

a design that tries to make things more parallel by sending instructions to different execution units at the same time, each execution unit (such as the ALU) is inside a single cpu, so if we come across instructions that can be run in parallel, we can direct them to a specific unit. we can also have units dedicated for some tasks (intger ALU and floating point ALU). in modern computer, other than very low level embedded devices, we use superscaler CPUs.

#### How it is different from pipelining

pipelining parallels instructions by separating the stages of the fetch-execute cycle, superscaler has multiple instructions in the same stage of the cycle, running inside different (specialized) execution units.
they can be used together.

#### How is it different from multicore processor

multi core processors are at a higher level the superscaler, a multicore system can have several cores, each running a cpu with pipelining and superscalering.

#### Hazards that come with Superscaler

nothing is free, we can get a poor assignment of instructions to execution units, which would be sub-optimal, we can get registry conflicts as well. for control hazards, we can sometimes process both branches and discard the unused results, trading heat for speed. superscaling makes data hazards more dangerous and complex, we have to be sure that the order is maintained. which might mean that we can't use a free execution context because we must wait for an instruction to be finished.

#### limitations

sometimes, the cost of trying to predict hazards and problems is more than giving up and doing the simple thing, the cost of checking dependencies and unrolling loops might be more than taking the hit and running instructions in sequence

</details>

## Data-level Parallelism

<!-- <details> -->
<summary>
TODO
</summary>

### Data-level Parallelism

if a pipeline is a instruction level parallism approach for processing multiple instructions simultaneously, then data-level parallism represents another step for increase throughput of processing.

- Defining three DLP approaches.
- Connecting how DLP approaches influence each other.
- Exploring the hardware implementations of each application.

#### SIMD

Single instruction, multiple data.

analog to walking four dogs at once in the same route, rather than walking them one by one. we can do the same instruction on multiple data at once.
this is used for:

- Vector processing
- SIMD extentsions
- Graphical Processing Units (GPUs)

#### Vector Processing

the old way of doing stuff is scalar processing (as opposed to vector processing).

using a single instruction on multiple points of data at once. turning loops into a single instruction.

this increase performance.

- Less instruction overhead: fewer fetch and decode stages, we pay a slight overhead for the initial instruction, but then multiple data elements can go through the pipeline without additional instructions.
- Memory access Bandwidth: we can also pipeline the memory access, we take the set of the data from the memory at once, rather than piece by piece.
- Pipelining: we can leverage the pipelining of the data for better performance because we don't require a write back.

#### Vector Architecture

there is a difference in the architecture between a scalar architecture and a vector.

- Vector registers: holding large amounts of data from the same type. can be much larger than normal registers.
- Internal looping: rather then fetch and decode the same instruction over and over for each element, we can store it inside the vector register.
- Lanes: processing vector data happens across _lanes_, each lane can do all the operations a scalar processor can do.

vector registers are more complex than scalar ones, but they provide better performance.

#### SIMD Extentsions

processor can carry both the scalar operations and the new vector operations, these additions are called SIMD extensions.

in the x86 instruction set, SIMD extensions began with _Streaming SIMD Extensions (SSE)_, which added 8 registers (xmm0 to xmm7), each with 128 bits (16 bytes, 4 floating points values).
Over the years more registers were added, such as zmm registers which hold 512 bits (64 bytes, 16 floating points,8 doubles... etc)

x86 instructions:

```MIPS
# add instruction
ADD r1, r2

# add packed floating-point
ADDPS xmm1, xmm2
VADDPS zmm1, zmm2, zmm3
```

ADDPS adds together the first and second operands, and stores the result back into the first one. VADDPS is an updated version, it supports up to 512 bits of data, it adds the results of adding the second and third operands and stores them inside the first operand.

#### GPU

GPU: Graphical Processing Unit

today gpus perform processing not only for graphics, also for many other things, like digital singaling, machine learning and cryptocurrency. some GPUs come together with the CPU, and some are external.\
Since GPUS were designed to handle graphics, they are built differently, they use simpler functionalities and favor volume over complexicy, they were build to handle large amounts of data. their internal clock may lag behind that of the cpu, and they aren't as efficient as cpu when dealing with branches and predictions.

**SIMT - Single instruction, Multiple Thread**

the SIMT architecture is designed to address the issue of high workloads of simple instructions over large data. the processing is done in bulk.

#### Summary

> "Great work finishing the topics on Data-Level Parallelism. In this lesson you covered:"
>
> - SIMD architectures and their benefits in data-heavy applications.
> - Vector processors and their early influence on SIMD architecture.
> - Vector architecture elements such as vector registers, lanes, and internal looping.
> - SIMD extensions and how vector processors influenced their implementation of commercial processing.
> - GPUs and how they take a different approach to SIMD architecture.

### Data-Level Parallelism Problem Set

</details>
