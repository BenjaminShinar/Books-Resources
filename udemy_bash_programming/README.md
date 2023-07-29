<!--
// cSpell:ignore Ermin Kreponic Aldin Omerdic dslconnection besename
 -->

# BASH Programming Course: Master the Linux Command Line!

udemy course [BASH Programming Course: Master the Linux Command Line!
](https://www.udemy.com/course/bash-programming/) by *Ermin Kreponic* and *Aldin Omerdic*. 

>Go from beginner to advanced with the Linux command line in this BASH programming course!
 
Bash and the command line are useful when working on servers, which don't have a gui interface. and the terminal commands have more control.

[online bash](https://www.tutorialspoint.com/execute_bash_online.php)

## Section 2: Setting Up the Environment

<details>
<summary>
Setting up a Linux virtual machine on a Windows Computer.
</summary>

open the windows system data and download the virtual box application for windows (matching the host machine), and then navigate to [getFedora](www.GetFedora.org) to grab the image file. it can also be done with vmware.

**note: only use one hypervisor, uninstall one before installing the other.**

follow the wizard and create a vm machine. power it on, then it will start installing. make sure the vm is connected to the internet.
on. set up a user and password. eventually things will complete and you could log-in into the machine.

some stuff to get full screen machine.

</details>

## Section 3: Unique Characters

<details>
<summary>
Special Characters in Bash
</summary>

Working with Bash
we can change the host name displayed as the prompt, we first need to be the root user.

```sh
su # change to root
hostnamectl set-hostname --static "CustomHostName"
```

we can update dependencies and install the vim editor.
```sh
dnf update
dnf install vim
```

we exit the root user with `exit`.

### Terminal Customization and Hash Character (`#`)

<details>
<summary>
Code comment, characters count, numeric bases.
</summary>

we can customize the terminal to look how we want it to.

when we want to clear the screen, we can type `clear`.

there are some unique characters which are used in bash. they have special built-in meaning.

we start with the hash symbol `#`, it is used to denote comments, and it has a special use to tell the machine how to interpret the text in the file. most of the script should start with `#!/bin/bash`

```sh
#!/bin/bash

echo "hi there?" # comment
# echo "nothing here" because it's a commented line!
```
we want to run the script, but we can't and we get a "permission denied" error,

```sh
./first.sh
ls -l
chmod +x first.sh # add execution permission
./first.sh
```
we clear the file and fill it again, this time seeing new meanings of the hash symbol
```sh
#!/bin/bash
echo "lala  #not a comment" 
echo another # yes comment
name=tea # set variables
echo "the $name contains ${#chars}"
echo $(( 2#111 )) # seven
echo $(( 10#111 )) # 111
echo $(( 16#111 )) # 273 = (16*16) + 16 +1
```

</details>

### SemiColon (`;`)

<details>
<summary>
Line separator.
</summary>

the semi colon works as a line separator, which helps when writing conditional statements. when we see the error response, it will count the semi-colons as new lines.
```sh
#!/bin/bash
echo "first" echo "second"
echo "line"; echo "break";

var=10
if [ "$var" -gt 0 ]; then echo "yes"; else echo "no"; fi;
```

</details>

### Dot (`.`)
<details>
<summary>
current directory
</summary>
the dot symbol indicates the current directory, the two dots is the previous directory

```sh
ls . # this directory
ls .. # previous directory
```

</details>

### Double quotes (`"`), Single Quotes (`'`) 

<details>
<summary>
partial and full quoting
</summary>

double quotes is partial quoting, it preserves most of the special charterers in a string.
```sh
colors="red black white"
echo $colors

for color in $colors
do
echo $color
done
```
if we wrap the variable in double quotes, it shows as a single line. it prevents the variable from being split into the components.
```sh
for color in "$colors"
do
echo $color
done
```

the single quote symbol is for full quoting, it preserves every special symbol and doesn't allow replacing sub strings.
```sh
for color in '$colors'
do
echo $color
done
```
</details>

### Comma(`,`), Double Comma (`,,`)

<details>
<summary>
Background Operations, lowercase conversions.
</summary>

when the single comma is user, only the last is returned, the value of y is 5, but the value of x is still initialized.
```sh
let "y=((x=20,10/2))"
echo $y
```
the double comma is used in lower case letter conversions.

```sh
var=DslConnection
echo ${var,} #dslConnection
echo ${var,,} #dslconnection
```
</details>


### Backslash (`\`), Slash (`/`) 

<details>
<summary>
Escape character
</summary>

the backslash is used as an escape charters, it tells the interpreter to render the characters as it is. it also helps to search for patterns which appear lke flags.
```sh
echo "\"double quote marks\""
ls --help | grep \"-U\"
```

the forward slash is a filesystem separator, the path of files and folders. it is also the division operator

```sh
let val=500/2
echo $val
```
</details>

### Back Quote(\`) 

<details>
<summary>
output capturing
</summary>

the back quote is used for capturing results of calls.
```sh
let val=500/2
num2=`echo $val`
echo $val
```
</details>

### Double Colon (`:`)

<details>
<summary>
Null command
</summary>

the double colon acts as the null command (do nothing).
```sh
var=20
if [ "$var" -lt 15]
then:
else 
echo $var
fi
```
it can also be used to empty file content, it's a "nothing value"
```sh
touch abc.txt
cat "some text" > abc.txt
: > abc.txt # now it's empty
```

</details>


### Exclamation Mark(`!`)

<details>
<summary>
Negation
</summary>
the examination mark is used to negate the result of some calculation

```sh
var=10
if [ "$var" != 0 ]
then
echo "not zero"
else
echo "zero
fi
```

</details>

### Asterisk (`*`)

<details>
<summary>
Wildcard
</summary>

the asterisk character is used as a wildcard, it can match any character.

```sh
touch test1 test2 test3 test4
ls test*
rm test*
```
it is also the multiplication operator, or for exportation (power)

```sh
let var=100*10
let var2=100**3

echo "$var $var2"
```
</details>

### Question Mark (`?`)

<details>
<summary>
The test operator (ternary operator)
</summary>

the question mark symbol is used a part of expressions, together with the  double column `:` for the else statement

```sh
var1=10
echo $(( var2 = var1<20?1:2))

```

</details>

### Parenthesis (`()`), Curly Braces(`{}`)

<details>
<summary>
Sub shell commands, array initialization, Substitution.
</summary>

grouping commands, sub-shell with internal variables

```sh
var=5
(var=10;) # sub shell
echo $var # still 5
```
also used for array initializations
```sh
color=(red white brown purple)
```

the curly braces is used for formatting and substitution, each of them gets the suffix and the prefix. the surrounding command applies to all of them. it also works for quick loops

```sh
echo \${test1,test2,test3}\$
cat {testFile01,testFile02} > testFile00

echo {0..9}
```
they are also used for anonymous functions that are visible to the outer scope.


```sh
var1=01
var2=02
{
var1=11
var2=12
}
```


</details>

### OR (`||` ) AND (`&&`)

<details>
<summary>
Logical operators
</summary>

the double pipe `||` acts as the **OR** operator
```sh
var=1
if [ "$var" -gt 0] || [ "$var" -eq -5]
then
echo "THEN PART - one or both are true"
else
echo "ELSE PART - both conditions are false"
fi
```
if we require both conditions to be true, then we use the logical **AND** operator of `&&`.

</details>

### Dash (`-`)

<details>
<summary>
parameters
</summary>
the dash operator can be be used as flag in the shell. it's a symbol for parameters

```sh
ls --help
```

also used for directing data from and into standard input/output.

(it's also a minus sign)

</details>

###  Tilde (`~`)

<details>
<summary>
home directory
</summary>

the `~` symbol takes us to the home directory

```sh
cd ~ # home directory
echo ~+ # current directory
echo ~- # previous directory
```
</details>

### Caret (`^`)

<details>
<summary>
upper case conversions
</summary>

```sh
word=tEst
echo ${word^} # TEst
echo ${word^^} # TEST
```

</details>

</details>

## Section 4: Variables and Parameters

<details>
<summary>
Variables and Parameters
</summary>

Variables are labels to represent values. variables in bash don't have types,.

### Variables Assignment

<details>
<summary>
Setting the value of variables
</summary>

the variable name and value aren't the same. we use the dollar sign `$` to get value from the name. we can remove the value with `unset`. there are some situations when we don't need the dollar sign to get the value

```sh
var=10
echo var # var
echo $var # 10
unset var
echo $var 
(( var2=var>5?1:0 ))
echo $var2
```

we can assign variables in several ways: directly, reading from input stream, or as a loop variable.
```sh
v1=1
echo "type some value:"
read v2
echo v2

for v in 1 2 3
do
echo "vale of v is $v"
done
```
</details>

### Properties of variables

<details>
<summary>
Whitespace in variables
</summary>

quoting variables preserves the whitespace

```sh
var="T r al la ll    "
echo $var # trimming variables
echo "$var" # keeps the whitespace
```

if we declare variables without quotes, we must escape the white spaces
```sh
v1="test1 test2 test3"
v2=test1\ test2\ test3
```
we can set values to *null*, and set multiple values in the same line.
```sh
v3= #empty, null
v4=11 v5=12 v6=13
echo "$v4 $v5 $v6"
```

when a variable is uninitialized, it is treated as a zero in arithmetic operators.

```sh
var1=
let "var1 +=10"
echo $var1
```
the back quote allows us to assign variables from return values. or wrap the command in parentheses and take the value with the dollar sign.
```sh
hi=`echo test`
echo $hi #test
v2=$(ls -la)
echo $v2
```
null values aren't zero, they have a different error when diving by them as compared to division by zero.

</details>

### Mathematical Operations and Substitutions

<details>
<summary>
Math operation and substation's
</summary>

```sh
num=1100
let "num=-10"
echo $num # 1090
var=${num/10/B} # substation
echo $var #B90
``` 
the integer value of a string is zero
```sh
var=A0
let "var += 10"
echo "$var" #10

var=hey1100
echo "$var"
num=${var/hey/20}
echo $num
```
</details>

### Variables Scope

<details>
<summary>
Local variables and Environment variables
</summary>

Variables can be scoped to the shell session or be environment variables, which are loaded when the shell starts.

the `export` command creates environment variables for spawned shells to inherit

```sh
local="not inherited"
export env="inherited"
```


</details>

### Positional Parameters

<details>
<summary>
Parameters from the command line
</summary>

positional parameters are passed from the calling command, we refer to them with the dollar sign and the positional number. the first nine parameters are accessed directly with the dollar sign, while the 10th and after require the `{}` wraps. the `-n` command checks if the argument was passed.

the name of the script is the zeroth parameter (`$0`), we can get the number of additional argument with `$#`

```sh
#!/bin/bash

MIN=10

if [ -n "$1" ]
then
echo "first one is $1"
fi

if [ -n "$2" ]
then
echo "second is $2"
fi

if [ -n "$3" ]
then
echo "third one is $3"
fi

if [ -n "${10}" ]
then
echo "tenth one is ${10}"
fi

echo "name of the script "$0"
echo "List of argument "$*""
echo "number of argument $#"
```
</details>

</details>

## Section 5: Return Values

<details>
<summary>
Getting the status of running a command
</summary>

A script can execute successfully or fail. the success return value is zero. 

### Getting the Return Value
<details>
<summary>
Finding out if the command was successful
</summary>

We can get the exit status with `echo $?`. if there is no return value from the command, then the exit status is the last return value or exit status inside them.

```sh
ls
echo $? # 0
echo "some test"
echo $? # 0
someCommandThatIsNotReal # error
echo $? # 127 
```
</details>

### Setting the Exit Status

<details>
<summary>
Manually deciding the exit status
</summary>

we can define the exit status of the script we create, this is done with the `exit` function.

the `-r` checks if the file exists and is readable
```sh
#!/bin/bash
NO_OF_ARGS=2
ER_BADARGS=85
ER_UNREADBLE=86

if [ $# -ne "$NO_OF_ARGS" ]
then
  echo "Usage `besename $0` testFile1 testFile2"
  exit $ER_BADARGS
fi

if [[ ! -r "$1" || ! -r "$2" ]]
then
  echo "one of the files doesn't exist or isn't readable"
  exit $ER_UNREADBLE
fi

cmp $1 $2 &> /dev/null # redirect output into trash
if [ $? -eq 0 ]
then
  echo "files are the same"
else
  echo "files are different"
fi

exit 0
```
</details>

</details>

## Section 6: Conditional Statements

<details>
<summary>
Conditions and Testing statements.
</summary>

### Nested Conditional Statements

<details>
<summary>
Nesting commands
</summary>

```sh
num=1
if [ "$num" -gt 0 ]
then
  if [ "$num" -lt 5 ]
  then
    if [ "$num" -gt 3 ]
    then 
      echo "gt 0, Lt 5, gt 4"
    fi
  fi
elif [ "$num" -eq 0]
  then
  echo "EQ 0"
else
  echo "no idea"
fi
```

we can use a special command `-e` to check if a file exists

```sh
var=/home/folder/file.txt

if [[ -e $var ]]
then
echo "file exists"
else
echo "file doesn't exist"
fi
```

</details>

### Double Parentheses

<details>
<summary>
evaluating mathematical expressions
</summary>

the double parentheses evaluates an arithmetic expression, if the result is not zero, then the exit status is zero (the normal exit code), however, if the expression evaluates to zero (or false), then the exit status is 1 (abnormal exit).

```sh
(( 2 > 1 )) # true
echo "exit status is $?" # 0
(( 2 < 1 )) # false
echo "exit status is $?" # 1
```
</details>

### File Operators

<details>
<summary>
conditions on files
</summary>

we have some operators to run on file, such as testing if it exists, if it's a regular file, if it has a size larger than zero, and if the user can read/execute or write to it.

</details>

### More conditionals

<details>
<summary>
Looking at more complicated script
</summary>

double square brackets are usually safer to use, we must use them when we have `&&` or `||` inside the condition.

we return to the previous script and finish it up.
</details>

</details>

## Section 7: Variables (continued)

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

### 
</details>

## Section 8: Loops
## Section 9: Internal Commands
## Section 10: Regular Expressions
## Section 11: Functions
## Section 12: Arrays
## Section 13: Lists
## Section 14: Debugging
## TakeAways

<!-- <details> -->
<summary>
Things worth remembering
</summary>

`#!/bin/bash` - at the top of the script

special characters:
- `#` - hash mark
  - comment
  - part of the `#!/bin/bash` definition
  - count characters in text
  - count argument parameters (`echo "$#"`)
  - define numeric base for numbers
- `;` - semi-colon
  - line separator
  - important for conditionals
- `.` - dot
  - current directory (`ls .`)
  - previous directory (`ls ..`)
  - range operator (`echo {0..9}`)
- `"` - double quotes
  - partial quoting
- `'` - single quote
  - full quoting - not evaluating the text inside
- `,` - comma
  - background operations
  - convert single character to lowercase (`echo ${var,}`)
  - convert all to lowercase (`echo ${var,,}`)
- `\` - backslash
  - escape character
- `/` - slash
  - filesystem separator
  - division operator
  - substitution (`${var/a/b}`)
- \` - back quote (tick mark)
  - take value from call return value
- `:` - double colon
  - null command, the nothing value
  - the *else* part of a ternary operator
- `!` - exclamation mark
  - negation
- `*` - asterisk 
  - wildcard
  - multiplication operator
  - exponent operator (power)
- `?` - question mark
  - ternary test operator (`x<5?1:2`)
  - the return status of the previous command (`echo $?`)
- `()` - parentheses
  - sub shell commands
  - array initializations
- `{}` - curly braces
  - formatting / substitution
  - visible anonymous functions
- `|` - pipe
- `||` - double pipe
  - logical **OR**
- `&` - ampersand
  - redirection
- `&&` - double ampersand
  - logical **AND**
- `-` - dash
  - flag
  - parameter
  - minus operator
- `%` - modules operator
  - reminder
- `~` - tilde
  - home directory
  - current directory (`~+`)
  - previous directory (`~+`)
- `^` - caret
  - convert single character to uppercase (`echo ${var,}`)
  - convert all to uppercase (`echo ${var,,}`)


conditionals:
- `-n $1` - variable exists
- `-z $1` - variable doesn't exits (like negation of `-n`)
- `-e "$file"` - file exists
- `-r "$file"` - file exists and is readable by the user running the test
- `-x "$file"` - file exists and is executable by the user running the test
- `-w "$file"` - file exists and is writable by the user running the test
- `-f "$file"` - file is regular file (not folder or device)
- `-s "$file"` - file is not zero size
  
</details>




