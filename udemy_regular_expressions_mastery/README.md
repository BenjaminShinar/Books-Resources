<!--
// cSpell:ignore Sujith POSIX fooa fcdplb abcx
 -->

# Regular Expressions (Regex) Mastery

 
udemy course [Regular Expressions (Regex) for Java,Linux,JavaScript,Python or other languages, with 30 illustrated exercises/examples](https://www.udemy.com/course/regular-expressions-mastery/) by *Sujith George*. 


## Building A Foundation
<details>
<summary>
Introduction, what are regex and where are they used.
</summary>

>"Regular expressions are a way to search for patterns within data sets."

finding and extracting patterns from datasets, for cases where <kbd>CTRL + F</kbd> isn't enough.

regex isn't a programming language.

we can find regex in most (if not all) programming languages,and in linux command line.

> "A regular expression, regex or regexp is a sequence of characters that define a search pattern"

###  Use Cases For Regular Expressions

lets look at some common usecases for regular expressions

in the google sign up page, it requires us to create a password:
> "Use 8 or more characters with a mix of letters, numbers and characters"

our password is validated using regex. 


(this looks like email validation):
```regex
^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$
```


even some text editors support regex, like notepad++.

here is a pattern that is used to find data from a csv.
```regex
.*Product(1|2).*N(J|Y),United\sStates
```

### Deep Dive - First Example

we allways start with an input file, a text file with strings on each line. we want our pattern to match some, but not all of them.

we want a pattern to match any string starting with "foo", with zero or more reptetion of the letter "a", and then "bar"

`*` - zero or more occurrences of the preceding token (character).

### A Generic Solution to Any Regex Problem

Steps to solve regex:
> - Understand the requirements
>   - what needs to be included
>   - what needs to be excluded
> - Identify patterns in the inclusion or exclusion list
> - Represent the patterns using regular expressions
> - use any regex engine to try the pattern

### Hands-on with Linux Grep Regex Engine and Java

all programming languages have a regex engine, but those engine are all different, however, almost all of them are POSIX compliant.

we start with linux `grep` command.

```sh
cd input_files
cat regex01.txt
grep 'fooa*bar' regex01.txt
```

```ps
Select-String -Path .\regex01.txt -Pattern "fooa*bar"
```

java example


```java
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Regex {

	public static void main(String args[]) {
		// The regex pattern
		final String REGEX_PATTERN = "fooa*bar";
		final String inputFileName = "regex01.txt";
		// Create a Pattern object
		Pattern r = Pattern.compile(REGEX_PATTERN);

		// Read the input file line by line
		try (BufferedReader bufferedReader = new BufferedReader(
				new InputStreamReader(Regex.class.getClassLoader().getResourceAsStream(inputFileName)))) {
			String line;
			while ((line = bufferedReader.readLine()) != null) {

				// Now create matcher object.
				Matcher m = r.matcher(line);

				// Apply the regex pattern to each line
				// If pattern matches, output the current line.
				if (m.find()) {
					System.out.println(line);
				}

			}
		} catch (IOException e) {
			e.printStackTrace();
		}

	}

}

```
### Quiz 1: Building a Foundation

> - what is the short name for regular expressions?
> - where is regex used?
> - what does the pattern "a*" stand for?


</details>

## The Basic Set
<details>
<summary>
Basic symbols for regex patterns: Quantifiers, WildCard, Anchors, Character class range.
</summary>

the Posix standard divides regex expression into a basic and extended sets.

### The Wildcard Symbol

in the file "regex02.txt", we want to match all the elements which start with "foo", and with "bar", and have a single letter between those two. so we use the `.` wild symbol to match everything. in matches any character (or white space).

`foo.bar`

```sh
Select-String -Path .\regex02.txt -Pattern "foo.bar"
```

### Wildcard Asterisk Combo

in the file "regex03.txt", we want to match all elements which start with "foo" and end with "bar", and between those words anything can exists, nothing, any letter or any letters which don't need to reapeat.

so we combine the wildcard symbol and the zero-or-more symbol. the combination of `.*` can match anything (and nothing), so it's very common to see.

```sh
Select-String -Path .\regex03.txt -Pattern "foo.*bar"
```

### Representing Whitespaces

in the file "regex04.txt", we want to match "foo", zero or more whitespaces, and then "bar". we can represent white spaces as either a space ` `, or with the special escape character `\s`, which can also match tabs (horizontal and vertical).

so we combine the whitespace symbol with the zero-or-more star quantifier `foo\s*bar`.

```sh
Select-String -Path .\regex04.txt -Pattern "foo\s*bar"
```
### Character Classes

in the file "regex05.txt" we want to select some characters, but not other. so we introduce the `[]` character class, which matches one of the elements inside it.

`[fcl]oo` will match either f,c,l, followed by "oo", this matches just a single character, it doesn't match all of them together or one following the other.

in file "regex06.txt", we match more symbols, but we see how it becomes hard to manage.
```sh
Select-String -Path .\regex05.txt -Pattern "[fcl]oo"
Select-String -Path .\regex06.txt -Pattern "[fcdplb]oo"
```
next, for file "regex07.txt", we don't want to match the "m" or h character, but we want to match any other characters. for this we use the `^` symbol inside the `[]` to match anything except the characters. this is the negation symbol,also called *caret* or *exponenet* symbol.

this also teaches us that in regex, there can be more than one solution.
```sh
Select-String -Path .\regex07.txt -Pattern "[^mh]oo"
```

### Character Classes With Ranges

If we have characters that follow on another alphabetically,we can use the range variation inside the character class `[a-f]` is like `[abcdef]`, the characters have to be in order. this is done with the ASCI values, so it's case sensitive (as all regex is).

```sh
Select-String -Path .\regex08.txt -Pattern "[j-m]oo"
```

we can combine ranges with regular characeter class. we simply add the individual letter before or after the range. `[a-cx]` is like `[abcx]`, so we can use the shorter range form, and add what's missing.

```sh
Select-String -Path .\regex09.txt -Pattern "[j-mz]oo"
```

as we said before, regex is case sensitive, so if we want to use upper case letters, we need to write them so as such, either individual characters or ranges.

```sh
Select-String -Path .\regex10.txt -Pattern "[j-mzJ-M]oo"
```

### Escaping With Backslash

example file "regex11.txt", they all have one or more "x", a single dot, and then one or more "y", so we want to match the dot somehow, even if we previously used it as a wildcard.

to use the dot character directly, we need to escape it from the normal usage as a wild card, this is done by preceding it with a backslash (this is called "escaping").

we can escape the following symbols: "^$*,[()\" 

```sh
Select-String -Path .\regex11.txt -Pattern "x+\.y+"
```

next, we want to match some symbols, but not others.
we don't need to escape the period wild card inside character class. if we had a hyphen `-`, we would need to escape it.

```sh
Select-String -Path .\regex12.txt -Pattern "x[.#:]y"
```

if we have a caret, we need to escape it, as it is used for negation.
```sh
Select-String -Path .\regex13.txt -Pattern "x[\^#:]y"
```

if we have a baskslash we wish to match, we have to escape it with a backslash as well `\\`. a literal backslash should always be escaped.
```sh
Select-String -Path .\regex14.txt -Pattern "x[\^#\\]y"
```
### Anchors

anchors control where the pattern appear in the entry, we can use the caret `^` to match the begining of the string.

```sh
Select-String -Path .\regex15.txt -Pattern "^foo"
```

the dollar symbol `$` matches the end of the string
```sh
Select-String -Path .\regex16.txt -Pattern "bar$"
```

if we want to match exactly on the entire entry (and not inside it), we can combine the anchors together.
```sh
Select-String -Path .\regex17.txt -Pattern "^foo$"
```
### Quiz 2: Regex: The Basic Set
> - match both "gray" and "grey".
> - match a two digits even number.
> - match a three digits that is divisible by 5.

</details>

## The Extended Set
<details>
<summary>
Advanced symbols in patterns. Repeaters and Binary Pipes.
</summary>

NOTE: we might need to pass the `-E` flag to the engine to specify we are using the extended regex set.

### Curly Braces Repeater

we start by first matching any three digits number, so we first use the character class range three times. we can instead use the **repeater**, the curly braces `{}` which takes a number, this is number of expected occurrences
```sh
Select-String -Path .\regex18.txt -Pattern "^[0-9][0-9][0-9]$"
Select-String -Path .\regex18.txt -Pattern "^[0-9]{3}$"
```

we can also specify a range of allowed repeats, this is done by providing the minimal and maximal number `{min,max}`. (both inclusive)
```sh
Select-String -Path .\regex19.txt -Pattern "^[a-z]{4-6}$"
```

we want to match a number of repetitions of a two character patters, so we use the parenthesis to group together a pattern into an entity, and use the curly braces with a minimal value and a comma, but without a maximal value, so we specify a minimal number of allowed repeats. `{min,}`.
```sh
Select-String -Path .\regex20.txt -Pattern "^(ha){4,}$"
```

~~we can also limit to a maximal number of repetitions, without a minimal value.~~
**this doesn't work in all engines!**

```sh
Select-String -Path .\regex21.txt -Pattern "^(ha){,2}$" # didn't work for me
```
### The Plus Repeater

we can match one-or-more occurrences of a pattern by using the plus `+` quantifier, which is like `a{1,}` (one or more "a") or `aa*` (a, followed by zero-or-more "a").
```sh
Select-String -Path .\regex22.txt -Pattern "^fooa+bar$" 
```

### The Question Mark Binary

we can also require zero-or-one occurrences, using the question mark symbol `?`. it's like `a{0,1}`.

```sh
Select-String -Path .\regex23.txt -Pattern "^https?://website$" 
```
### Making Choices With Pipe

if we want to match one pattern (which isn't a single character), we can use the pipe `|` operator to separate options. we put them inside parenthesis to specify this a single entity. we can also repeat this operator for more than one option.

```sh
Select-String -Path .\regex24.txt -Pattern "^(log|ply)wood$" 
Select-String -Path .\regex24.txt -Pattern "^(log|ply|red)wood$" 
```
### Quiz 3: Regex: The Extended Set
> - match "colour" and "color"
> - match "ascending" and "descending"
> - match one or more "a"


</details>

## Find and Replace with Capture Groups
<details>
<summary>
Capture Groups
</summary>

### The Monitor Resolutions Problem
### The First Name Last Name Problem
### The Clock Time Problem
### The Phone Number Problem
### The Date Problem
### Another Phone Number Problem
### Quiz 4: Regex: Find and Replace with Capture Groups

</details>

## Takeaways
<details>
<summary>
Things worth remembering
</summary>

in powershell
```ps
Select-String -Path C:\temp\*.log -Pattern "pat"
```

email validation - `^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$`

some symbols need to be escaped in different situations.
- inside character class (square brackets):
  - Escape the caret `^` negation symbol.
  - Escape the hyphen `-` range symbol.
  - Escape the backslash `\` escape character.
  - **no need** to escape the `.` wild card.

character class: negation caret, range hyphen

- special symbols
- quantifiers
- anchors
- character classes

</details>

