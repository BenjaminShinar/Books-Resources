<!--
// cSpell:ignore Udemy Sawler gotemplate
-->

# Working with Concurrency in Go (Golang)

Udemy course [Working with Concurrency in Go (Golang)](https://www.udemy.com/course/working-with-concurrency-in-go-golang) by *Trevor Sawler*.

advanced class about go and concurrency.

## Section 1: Introduction

<details>
<summary>
Getting started
</summary>

> Don't communicate by sharing memory, share memory by communicating.

it's easy to run stuff in the background, we simply add `go` and launch a goroutines. but if we wish to communicate between those functions (goroutines), we can use:

1. primitives from go **sync** package
2. locks and mutexes
3. waitGroups
4. channels

we shouldn't use shared memory and complicated synchronization primitives, instead, we should just pass data in channels.\
concurrent programming is hard and error-prone, so if we don't use it, we shouldn't have it. and if we must use it, it should be kept to the minimum.

the course will show the basic types in the sync package: mutexes (semaphores) and wait groups. then apply them in three classic computer science problems:

- Producer/Consumer
- [Dining Philosopher](https://en.wikipedia.org/wiki/Dining_philosophers_problem)
- [Sleeping Barber](https://en.wikipedia.org/wiki/Sleeping_barber_problem)


then we'll also build a project for ourself, a subscription service that sends emails, generates PDF files and we'll have testing for it.

we install go, visual studio code,the go extension for vsCode (we install all the suggested tools), the gotemplate-syntax extension and make.

</details>

## Section 2: Goroutines, the `go` Keyword and WaitGroups

<details>
<summary>
Starting with go waitGroups.
</summary>

This section will focus on goroutines, how to use them, what are the problems with them, and how to solve the problems.
Goroutines are functions that run in the background (concurrently with other code). they are simple to use, but can create problems.

### Creating GoRoutines and Identifying a Problem

<details>
<summary>
demonstrating a problem with goroutines.
</summary>

we start with the basic main file. and copy the sample code into it.\
actually the `main` itself is a goroutine. goroutines aren't normal processor threads, instead, they are specialized lightweight threads. they are managed by the go scheduler.\
we add a new function "printSomething" that prints whatever is passed to it. to make a function call concurrent, we prefix it with the `go` keyword, then it runs in it's own thread. but if the program concludes before the goroutine completes, then we never see the output.

we can fix this in several ways, and we'll start with the **worst one** - this is by delaying the main thread execution using `time.sleep(1 * time.Second)` to waste time.


```go
package main

import (
	"fmt"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	// if you run the program with this line uncommented, and the lines 20 commented,
	// everything works as expected
	printSomething("This is the first thing to be printed!")

	// but if you comment out line 15 and uncomment the one below this comment,
	// running the program will (probably) just print out the final message,
	// since the program terminates before the goroutine started by this
	// command does not have time to finish.
	//go printSomething("This is the first thing to be printed!")

	// in order to give the goroutine from line 20 time to finish, we could
	// wait for second by uncommenting the line below, but this is hardly
	// a good solution.
	//time.Sleep(1 * time.Second)

	printSomething("This is the second thing to be printed!")
}
```

</details>

### WaitGroups to the Rescue

<details>
<summary>
Using waitGroups.
</summary>

now lets show why the `sleep` solution is a bad idea, we create a slice (range) of strings called words, we loop over it and call the goroutine. they are all printed, but not it the original order.\
In the real order, we don't know how long an operation will be, so how can we choose how long to wait for? if the list of words was thousands of words, maybe we wouldn't see them all if we just waited a single second. this line of action gets us nowhere, so we introduce `waitGroups` as an alternative.

we create the variable wg of type `sync.WaitGroup`, and we add entries the size of the words slice and after the loop use the `wg.Wait()` operator. we need to modify function to take the workGroup as a pointer, and decrease the value with a deferred command. **WaitGroups shouldn't be copied or modified**.\
if the waitGroup is at zero, then we get an error. they can't be decreased below that value.

```go
package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the second thing to be printed!", &wg)
}
```

</details>

### Writing Tests with WaitGroups

<details>
<summary>
Testing Goroutines
</summary>

if we have too many entires in the waitGroup and all our goroutines have completed, we won't hang forever. instead, we get a deadlock fatal error - "all goroutines are asleep".


lets look at the testfile.

we capture the standard output stream from the operating system with `os.Pipe()` and `os.StdOut`. once we finished with the waitGroup, we can close the pipe and read the data from the stream.

```go
package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("epsilon", &wg)

	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "epsilon"){
		t.Errorf("Expected to find epsilon, but it is not there")
	}
}
```

</details>

### Challenge 1: Working With WaitGroup

<details>
<summary>
Checking What We learned
</summary>

now we have a challenge, we need to modify the code so that it uses goroutines and prints at the correct order. we also need to add tests!

this it the original code!

```go
package main

import (
	"fmt"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	updateMessage("Hello, universe!")
	printMessage()

	updateMessage("Hello, cosmos!")
	printMessage()

	updateMessage("Hello, world!")

	printMessage()
}
```

my code is in "\challenges\challenge-1\main.go".

```sh
cd challenges\challenge-1
go run .
go test .
```

#### Solution to Challenge

the solution used package level variables for the waiting group. I hate this.

</details>

</details>

## Section 3: Race Conditions, Mutexes and an Introduction to Channels

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


### Race Conditions: an example

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Adding sync.Mutex to our code

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Testing for race conditions

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### A more complex example

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Writing a test for our weekly income project

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Producer/Consumer - Using Channels for the first time

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Getting started with the Producer - the pizzeria function

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Making a pizza: the makePizza function

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Finishing up the Producer code

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Creating and running the consumer: ordering a pizza

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

### Finishing up our Producer/Consumer project

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>



</details>

## Section 4: A Classic Problem: The Dining Philosophers
## Section 5: Channels, and Another Classic: The Sleeping Barber Problem
## Section 6: Final Project - Building a Subscription Service
## Section 7: Sending Email Concurrently
## Section 8: Registering a User and Displaying Plans
## Section 9: Adding Concurrency to Choosing a Plan
## Section 10: Testing

## Takeaways

<details>
<summary>
Stuff worth remembering.
</summary>

- `go version`
- `go build`
- `go run`
- `go fmt`
- `go install`
- `go get`
- `go test`

### The Sync Package

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[documentation](https://pkg.go.dev/sync)

if the waitGroup value goes below 0, we get an error.
</details>


</details>

