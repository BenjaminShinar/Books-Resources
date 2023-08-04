<!--
// cSpell:ignore Ermin Kreponic Aldin Omerdic dslconnection besename cdspell Docunents
 -->

# BASH Programming Course: Master the Linux Command Line!

udemy course [BASH Programming Course: Master the Linux Command Line!
](https://www.udemy.com/course/bash-programming/) by _Ermin Kreponic_ and _Aldin Omerdic_.

> Go from beginner to advanced with the Linux command line in this BASH programming course!

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

the question mark symbol is used a part of expressions, together with the double column `:` for the else statement

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

### Tilde (`~`)

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

we can set values to _null_, and set multiple values in the same line.

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

<details>
<summary>
Diving deeper
</summary>

### Built-In Variables

<details>
<summary>
Variables that we get from the shell
</summary>

```sh
echo $BASH # shell path
echo $$ # process id
echo $BASH_VERSION # bash version

for n in {0..5}
do
  echo "BASH_VERSINFO[$n] = ${BASH_VERSINFO[$n]}"
done

echo $PATH # source path
echo $UID # user id
echo $EUID # user effective id (for assumed roles)
echo $GROUPS # Groups the user belongs to (array)
echo $HOME # home directory
echo $HOSTNAME # hosting machine name
echo $HOSTTYPE # hosting machine type
echo $MACHTYPE # hosting machine type (more data)
echo $OSTYPE # hosting machine type (more data)

colors1="red-brown-orange"
colors2="red+brown+orange"
IFS=-
echo $colors1 # red brown orange
echo $colors2 # red+brown+orange
IFS=+
echo $colors1 # red-brown-orange
echo $colors2 # red brown orange

echo $PWD # working directory
echo $OLDPWD # previous working directory
cat someFile.txt | sort | wc -l
echo ${PIPESTATUS[*]} # the status of all commands in the last pipe

echo "some Question"
read # store in $REPLY
echo "The answer is $REPLY"

echo "some other Question"
read testVar # store in $testVar
echo "The answer is $testVar"

echo $SECONDS
sleep 12 # sleep for number of seconds
```

script only

```sh
echo $BASH_ENV
echo $CDPATH
echo $EDITOR
echo $FUNCNAME
echo $LINENO # line number
```

</details>

### Setting Variables Properties

<details>
<summary>
Crating variables with special behavior
</summary>

`declare` and setType

- `-r` - read only (unChangeable)
- `-i` - integer (can be used in arithmetic commands without `let`)
- `-a` - array
- `-f` - function
- `-x` - can be exported

```sh
declare -r varName=5 # read only variable

echo "\$varName = $varName"

declare -i intVar = 7
echo $
intVar = 10/5 # evaluated as number
echo $
intVar = blue # evaluated as an integer
echo "\$intVar = $intVar" # 0

declare -a ar=(10 27 39 42 54)
for i in {0..4}
do
  echo "${ar[$i]}"
done

declare -f someFunction

someFunction()
{
  echo "calling from some function"
}

someFunction
```

</details>

### Random Number Generation

<details>
<summary>
Getting a Random Number
</summary>

the `$RANDOM` function generates an integer between 0 and 32767 (max short: $2^{15}-1$)

```sh
#!/bin/bash
MAX=10
i=1

while [ "$i" -le $MAX ]
do
  n = $RANDOM
  echo $n
  let "i += 1"
done
```

</details>

</details>

## Section 8: Loops

<details>
<summary>
Repeating Actions for a number of iterations and control flow actions.
</summary>

### For Loop

<details>
<summary>
Looping for a pre-determined number of times
</summary>

nested _for-loops_

```sh
for in in 1 2 3 4 5
do
  echo "outer Loop $i"
  for j in 1 2 3
    echo "inner loop $j"
  done
done
```

</details>

### While Loop

<details>
<summary>
Looping while a condition is true.
</summary>

first test a condition, if it's true, perform the loop body. we can perform actions before testing condition

```sh
a = "not set"
prev = $a

while
  echo "Previous variable = $prev"
  echo
  prev=$a
  [ "$a" != end ]
do
  echo "input env to exit or anything else to go on"
  read a
  echo "variable a = $a"
done
```

</details>

### Until Loop

<details>
<summary>
Loop until the condition is true.
</summary>

the inversion of a `while` loop.

```sh
until [ "$n" = end ]
do
    echo "input env to exit or anything else to go on"
    read n
    echo "variable n = $n"
done
```

</details>

### `Break` and `Continue`

<details>
<summary>
Control flow statements inside loops.
</summary>

- `break` - break outside the loop.
- `continue` - go to the next loop iteration without executing the rest of the code.

```sh
n=0
LIMIT=20
while [ "$n" -le "$LIMIT" ]
do
  let "n =+1"
  if [ "$n" -eq 3 ] || [ "$n" -eq 9 ]
  then
    continue
  fi
  echo $n
  if [ "$n" 17 ]
  then
    break
  fi
done
```

if we have nested loops, we can add the level of nesting to the `break` and continue commands

```sh
for i in {0..5}
do
  for j in {0..5}
    do
    for k in {0..5}
    do
      echo "$i, $j, $k"
      if [ "$i" = 2 ] && [ "$j" = 4 ]
      then
        continue 2
      fi
      if [ "$i" = 3 ] && [ "$j" = 4 ]
      then
        break 3
      fi
    done
  done
done
```
</details>

### Case Construct

<details>
<summary>
choosing behavior based on state
</summary>

a different way to formulate conditional statements

requires double semi-colons (`;;`) at the end of each condition. we stick a closing parentheses (`)`) to end the case condition.

```sh
read a
case "$a" in
[[:upper:]] ) echo "$a is a uppercase letter";;
[[:lower:]] ) echo "$a is a lowercase letter";;
[0-9] ) echo "$a is a digit";;
* ) echo "$a is a special character";;
esac
```

</details>

### Select Construct

<details>
<summary>
Structured statement
</summary>

this is a structured while loop that displays the options. it uses the `$PS3` variable to display the prompt each time
```sh
PS3='Pick a color' # prompt
select color in "brown" "grey" "black" "red" "green"
do
  echo "you selected $color"
  break
done
```
</details>

</details>

## Section 9: Internal Commands

<details>
<summary>
commonly used shell commands.
</summary>

more commands, such as I|O commands, like `echo`, `read` and `printf`. other commands work on variables, such as `let` and `eval`. `set` and `unset` change the properties of shell variables.

### `printf`

<details>
<summary>
Print formatted string
</summary>

printing a formatted string
```sh
declare -r PI=3.145926
printf "Second Decimal is %1.2f\n" $PI
printf "Second Decimal is %1.4\n" $PI
```

</details>

### `read`

<details>
<summary>
Reading data into a variables
</summary>

variables on the read command.
- `-s` - silent, don't echo input
- `-n <number>` - limit to a number of characters
- `-p <prompt>` - text prompt

```sh
up=$'\x1b[A'
down=$'\x1b[B'
left=$'\x1b[C'
right=$'\x1b[D'

read -s -n3 -p "Press an arrow key " arrow
case "$arrow" in
  $up) echo "you pressed up";;
  $down) echo "you pressed down";;
  $left) echo "you pressed left";;
  $right) echo "you pressed right";;
esac
```

we can also read from a file, we direct the file to become the input.

```sh
echo "Read"
while read line
do
  echo $line
done <wood.txt
```
</details>

### `eval`

<details>
<summary>
combine commands and execute them
</summary>

```sh
if [ ! -z $1 ] # if argument 1 exists
then
  process="ps -e | grep $1" # string
fi
eval $process # execute as the string as a command
```

</details>

### `set` and `unset`

<details>
<summary>
Change the values of shell options a variables.
</summary>

 we use dash `-` to set them and plus `+` to remove them (contrary to common sense). running `set` alone displays all the option.

for example, we can choose to store or not store  commands in the shell history.
 ```sh
history
set +o history # stop recording command in history"
history
echo "not shown"
echo "missing from history"
echo "how did this happen"
history
set -o history # start recording history again
echo "after bringing the history command back"
history
```

we can also set the shell variables, instead of passing positional parameters. this is done with `--` and the parameters.
```sh
echo "before setting"
echo "\$1 = $1"
echo "\$2 = $2"
set `echo "tree four"` # now the $1 and $2 have values
echo "after setting"
echo "\$1 = $1"
echo "\$2 = $2"
```

```sh
var="1 2 3"
echo $var
echo $#
set -- $var # set positional variables to to the values of bar
i=1
while [ $i -le $# ]
do 
  echo $i
done
echo $#
set -- # unset positional  variables
echo $#
echo "\$1 = $1"
echo "\$2 = $2"
echo "\$3 = $3"
```
we can unset (delete) declared variables with `unset`, we simply pass the name, not the value.

```sh
var="a b c"
echo $var
unset var
echo $var
```

</details>

### `getopts`

<details>
<summary>
Parsing options that are passed to the script
</summary>


```sh
while getopts :dm option
do
  case $option in
    d) d_option=1;; # set value
    m) m_option=1;; # set value
    *) echo "usage: -dm";;
  esac
done

day=`date | awk '{print $1 " " $2}'`
month=`date | awk '{print $3}'`
if [[ ! -z $d_option ]]
then
  echo "Date is $day"
fi

if [[ ! -z $m_option ]]
then
  echo "Month is $month"
fi

shift $(($OPTIND -1)) # get next position parameter
```

we then call this script with the options (flags) and see the behavior.
</details>


### `shopt`

<details>
<summary>
Show, set and unset shell option
</summary>

set and unset shell options. `shopt` displays all the options, we can set them with `-s` and unset with `-u`.

```sh
ls Docunents # wrong text
shopt -s cdspell # set auto correct
ls Docunents # wrong text
shopt -u cdspell # unset auto correct
```

</details>

### `type`

<details>
<summary>
Getting the type of the command
</summary>

check if a command exists, what is it's type and is it an alias to some other command.
```sh
type cd
type git
type ls
type lala
```
</details>

###  `jobs`, `disown`, `fg`, `wait` and `kill`

<details>
<summary>
running jobs in the background
</summary>

we can see the running jobs in the background, and remove them with `disown`.
```sh
sleep 15 & # run in background
jobs
# wait
jobs

sleep 60 & # run in background
jobs
disown
jobs
```

`fg` brings a background command to the foreground, after we pushed them to the background with `&`.
`wait` hangs the shell until all background commands finish, `times` writes the cpu times.\
the `kill` commands terminates a process, we can pass the process id to determine which task to kill. we get the process id with `ps` and `ps aux`. if we have command that shares the name with a built in shell command, then we can ignore the shadowing with the `command` prefix, that makes sure to call the shell commands.

```sh
ls()
{
  echo "fake ls command!"
}

ls # call fake command
command ls # calls shell command
```
</details>

</details>


## Section 10: Regular Expressions

<details>
<summary>
Searching and manipulating text and strings.
</summary>

characters can be literal or have special meaning.

### Grep

<details>
<summary>
Basic Regular expressions
</summary>

```sh
E_NOPATTERN=71
DICT=/usr/share/dict/linux.words

if [ -z "$1" ] # no parameter
then
  echo ""
  echo "Usage:"
  echo "`basename $0` \"pattern,\""
  echo "where \"pattern\" is in the form"
  echo "ooo..oo.o..."
  echo
  exit $E_NOPATTERN
fi

grep ^"$1"$ "$DICT"
```

the `^` matches the beginning of the line, while `$` matches the end of the line.

this script wants a pattern like "wo.d", which will match the known characters and using the `.` wild card.

</details>

### Extended Regular Expressions

<details>
<summary>
using the sed command
</summary>

this example iterates over the input parameters and tries to delete all the lines before the first empty line, and then deletes the empty lines as well

```sh
if [ $# -eq 0 ]
then
  echo "needs arguments"
  exit $E_BADARGS
else
  for i # loop over positional arguments
  do
    sed -e '1,/^$/d' -e '^%/d' $i
  done
fi
```

</details>

### Globing

<details>
<summary>
A different way to match files
</summary>

bash doesn't support regular expressions natively (some commands do). instead, it has support for **glob**.

- `?` match one character
- `*` match any number of character
- `[ab]` - match `a` or `b` characters
- `[b-t]` - match a range of characters
- `{w*, *oo*}` - match either of the conditions


the `echo` command can also do some globing (file name expansion). we can the options in the shell command options `shopt`.

</details>

</details>

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

when using echo, add `-e` flag to escape special characters such at new lines and vertical tabs
special characters.
```sh
echo "aaa \t bbb \n ccc"
echo -e "aaa \t bbb \n ccc"
echo $((1+3)) # numeric operation
```

reading variables from input:\
`read -s -n 3 -p "enter your selection\n" num`
- `-s` - silent, don't echo input
- `-n <number>` - limit to a number of characters
- `-p <prompt>` - text prompt
### Special Characters Meanings
special characters
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
  - the _else_ part of a ternary operator
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
  - run command in the background
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

### Conditions
conditionals:

- `-n $1` - variable exists
- `-z $1` - variable doesn't exits (like negation of `-n`)
- `-e "$file"` - file exists
- `-r "$file"` - file exists and is readable by the user running the test
- `-x "$file"` - file exists and is executable by the user running the test
- `-w "$file"` - file exists and is writable by the user running the test
- `-f "$file"` - file is regular file (not folder or device)
- `-s "$file"` - file is not zero size

case statement
```sh
read a
case "$a" in
[[:upper:]] ) echo "$a is a uppercase letter";;
[[:lower:]] ) echo "$a is a lowercase letter";;
[0-9] ) echo "$a is a digit";;
* ) echo "$a is a special character";;
esac
```

`select` is a like a *while* loop that prompts and can choose a value from the list with a number
```sh
PS3='Pick a color' # prompt
select color in "brown" "grey" "black" "red" "green"
do
  echo "you selected $color"
  break
done
```

### Built-in variables

variables that are pre-build

- `HOME` - home directory
- `HOSTNAME` - host machine name
- `HOSTTYPE` - host machine type
- `MACHTYPE` - machine type (more data)
- `OS` - operating system
- `OSTYPE` - operating system type
- `PWD` - previous working directory
- `OLDPWD` - previous working directory
- `BASH` - the path to Bash
- `BASH_VERSION` - bash version (simple)
- `BASH_VERSINFO` - bash version in array form with more data
- `$$` - process id
- `PATH` - source path
- `UID` - user id
- `EUID` - user effective id
- `GROUPS` - an array with the ids of the groups the user belongs to
- `IFS` - string field separator
- `PIPESTATUS` - status code for the last pipelined commands, use `echo ${PIPESTATUS[*]}` to see all
- `REPLY` - the default placeholder variables to store replies from `read`
- `SECONDS` - number of seconds since shell was started (`echo $(($SECONDS/60))`) for minutes
- `RANDOM` - generates a pseudo-random number between $[0, 32767]$
- `PS3` - prompt string 3, works with `select` statements
- `OPTIND` - together with `getopts`
- variables that don't always have values
  - `BASH_ENV` - (script only)
  - `CDPATH` - (script only)
  - `EDITOR` - (script only)
  - `FUNCNAME` - the name of the surrounding function or empty when called from the top level.
  - `LINENO` - the line number (script only).

### Declaring Variables

declaring variables flags

- `-r` - read only
- `-i` - as integer (used in numeric calculations without `let`, assignments to it will be evaluated as numeric). there are no floating point types
- `-a` - array
- `-f` - function
- `-x` - exported
</details>
