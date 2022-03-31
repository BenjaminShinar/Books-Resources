<!--
// cSpell:ignore Sujith POSIX fooa
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
Basic symbols for regex patterns.
</summary>


### The Wildcard Symbol
### Wildcard Asterisk Combo
### Representing Whitespaces
### Character Classes
### Character Classes With Ranges
### Escaping With Backslash
### Anchors
### Quiz 2: Regex: The Basic Set


</details>

## The Extended Set
<details>
<summary>
Advanced symbols in patterns.
</summary>
</details>

## Find and Replace with Capture Groups
<details>
<summary>
Capture Groups
</summary>
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

</details>

