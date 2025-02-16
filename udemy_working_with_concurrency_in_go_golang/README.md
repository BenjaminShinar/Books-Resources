<!--
// cSpell:ignore Sawler gotemplate fatih randomMillseconds
-->

<link rel="stylesheet" type="text/css" href="../markdown-style.css">

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

1. primitives from go <golang>sync</golang> package
2. locks and mutexes
3. waitGroups
4. channels

we shouldn't use shared memory and complicated synchronization primitives, instead, we should just pass data in channels.\
concurrent programming is hard and error-prone, so if we don't use it, we shouldn't have it. and if we must use it, it should be kept to the minimum.

the course will show the basic types in the sync package: <golang>mutexes</golang> (semaphores) and <golang>waitGroups</golang>. then apply them in three classic computer science problems:

- [Producer/Consumer](https://en.wikipedia.org/wiki/Producer-consumer_problem)
- [Dining Philosophers](https://en.wikipedia.org/wiki/Dining_philosophers_problem)
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
actually the <golang>main</golang> itself is a goroutine. goroutines aren't normal processor threads, instead, they are specialized lightweight threads. they are managed by the go scheduler.\
we add a new function "printSomething" that prints whatever is passed to it. to make a function call concurrent, we prefix it with the <golang>go</golang> keyword, then it runs in it's own thread. but if the program concludes before the goroutine completes, then we never see the output.

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

now lets show why the <golang>sleep</golang> solution is a bad idea, we create a slice (range) of strings called words, we loop over it and call the goroutine. they are all printed, but not it the original order.\
In the real order, we don't know how long an operation will be, so how can we choose how long to wait for? if the list of words was thousands of words, maybe we wouldn't see them all if we just waited a single second. this line of action gets us nowhere, so we introduce <golang>waitGroups</golang> as an alternative.

we create the variable wg of type <golang>sync.WaitGroup</golang>, and we add entries the size of the words slice and after the loop use the <golang>wg.Wait()</golang> operator. we need to modify function to take the workGroup as a pointer, and decrease the value with a deferred command. **WaitGroups shouldn't be copied or modified**.\
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

lets look at the testfile.\
we capture the standard output stream from the operating system with <golang>os.Pipe()</golang> and <golang>os.StdOut</golang>. once we finished with the waitGroup, we can close the pipe and read the data from the stream.

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

<details>
<summary>
Understanding race conditions, locks, channels and the producer/consumer problem.
</summary>

this section covers other issues in concurrent programming: race conditions, locking with <golang>mutex</golang> and <golang>channels</golang>.

<golang>mutex</golang> allows to lock resources that are used by two or more goroutines, and we need to control access to it, and to prevent them from changing the data at the same time. <golang>channels</golang> share data between goroutines (either uniDirectional or biDirectional).

### Race Conditions: An Example

<details>
<summary>
Showing how a race condition can happen.
</summary>

we start by creating a go program with a race condition. we still use <golang>waitGroups</golang>, so we will wait for the operations to finish, but the two updates can still happen in the same time!

```go
package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!")
	go updateMessage("Hello, cosmos!")
	wg.Wait()

	fmt.Println(msg)

}
```

we can fire the program with the `go run` command, and after running them a few times, we see different results. we can also run `go run -race` and get a warning that we have a race condition.

#### Adding `sync.Mutex` to Our Code

the fixed code adds a <golang>sync.Mutex</golang>, a lock that only one thread can hold.  a mutex should never be copied, and only passed as a pointer.

```go
package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock() // take mutex
	msg = s
	m.Unlock() // release mutex
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmos!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
```

we still have indeterminate order, but no race condition. we can confirm with `go run -race`.

#### Testing For Race Conditions

we can write tests to check our earlier code, we can add the `-race` flag to the `go test` command.

```go
package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Goodbye, cruel world!")
	wg.Wait()

	if msg != "Goodbye, cruel world!" {
		t.Error("incorrect value in msg")
	}
}
```

</details>

### A More Complex Example

<details>
<summary>
A more complex example of modifying data with locks.
</summary>

a more complex command, using both <golang>waitGroups</golang> and <golang>mutex</golang> and showing how race conditions can corrupt our data. our program will be an income calculator, we have custom <golang>struct</golang> with two fields. and we will calculate our income for each week and build up a yearly total. we also use the goroutine as an inlined expression, rather than define it outside.\
if we only use waitGroups, each goroutine will read the value and modify it, without knowing that other has modified between reading and writing. we can check for data races and we get a warning.

```go
package main

import (
	"fmt"
	"sync"
)


var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()
				
				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}
```

#### Writing a Test for our Weekly Income Project

we can a test to see the correct amount is being tallied. we know that $(500 + 100 + 50 + 10)* 12 = 34320$, so we test for that number. and again, we need to capture the output.

```go
package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if ! strings.Contains(output, "$34320.00") {
		t.Error("wrong balance returned")
	}
}
```

</details>

### Producer/Consumer - Using Channels For The First Time

<details>
<summary>
The Producer Consumer project about a Pizzeria.
</summary>

the producer-consumer problem (by *Dijkstra*) deals with the issue of having concurrent "writers" who create or produce data and "readers" that need to read or consume the data. the readers must wait until there's data to read.

we start with an outline of the program. our example will describe a pizzeria, with the producers being the pizzeria making pizzas, and the consumers being customers who eat the pizzas. in our example, sometimes we will fail in creating a pizza.\
The important thing is that we use the <golang>channels</golang> with <golang>chan</golang> keyword. we can have a channel of channels!

```go
package main

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	// seed the random number generator

	// print out a message

	// create a producer

	// run the producer in the background

	// create and run consumer

	// print out the ending message
}
```

#### Getting Started with the Producer - The Pizzeria Function

Now we focus on creating the producer function. we initialize our random number generator with a seed from the <golang>math/rand</golang> package, and we add a module from github to print colored text.

the producer has two fields, a channel of pizza order and a channel of channels of errors. we create the using the <golang>make</golang> function to create the inner channels. we pass the producer to a function that we run as a goroutine.\
**when we are done with a channel, we must close it.**, so we create a function that on the producer type called *Close* that we can use. in the next section, we will use the <golang>select</golang> operator to make decision based on the data in the channel.

```go
package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		// try to make a pizza
		// decision
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer

	// print out the ending message
}

```

#### Making a Pizza: the `makePizza` Function

now we create a function that creates a pizza, which returns a pointer a pizzaOrder object. we add some delay to make things easier to see during program execution. we decide that for some cases, the pizza creation failed, and we have different reasons for it to fail.\
if we don't need to create a new pizza, we return a pizza order object without increasing the number.

```go
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds....\n", pizzaNumber, delay)
		// delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <=2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message: msg,
			success: success,
		}

		return &p

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		// try to make a pizza
		// decision
	}
}
```

#### Finishing Up the Producer Code

now we want to listen to the channels we created.

```go
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we receive a quit notification
	// try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
				// we tried to make a pizza (something was sent to the data channel)
				case pizzaMaker.data <- *currentPizza:
				case quitChan:= <- pizzaMaker.quit:
					// close channels
					close(pizzaMaker.data)
					close(quitChan)
					return // exit goroutine
			}
		}
	}
}
```

the program still doesn't work as we wanted, since we don't wait for a the goroutine, and we haven't created consumers.

#### Creating and Running the Consumer: Ordering a Pizza

we head back to the `main()` function. we want to create a consumer.
we loop over the channel of pizza orders, and if we're done with the orders, we close the channel.

```go
func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	// print out the ending message
}

```
#### Finishing up Our Producer/Consumer Project

we can also add a finishing message to make sure all our producers finished. we use the `switch` statement.

```go
	// ... the code from above
	// print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day....")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day!")
	default:
		color.Green("It was a great day!")
	}
```

</details>

</details>

## Section 4: A Classic Problem: The Dining Philosophers

<details>
<summary>
Resource deadlock.
</summary>

Another classic problem, also by *Dijkstra*, meant to show resource deadlock. this time we'll go back to using the <golang>sync</golang> package.

### Getting Started With The Problem

<details>
<summary>
introducing the problem.
</summary>

> The Dining Philosophers problem is well known in computer science circles. Five philosophers, numbered from 0 through 4, live in a house where the table is laid for them.\
> each philosopher has their own place at the table. Their only difficulty – besides those of philosophy – is that the dish served is a very difficult kind of spaghetti which has to be eaten with two forks. There are two forks next to each plate, so that presents no difficulty.\
> As a consequence, however, this means that no two neighbors may be eating simultaneously, since there are five philosophers and five forks.

we start with the problem statement, and create an outline for the program. we create data structures for our objects. a philosopher needs to have two specific forks in order to eat. we define some variables:

- how many times does each philosopher needs to eat before he's done.
- a delay between each time a philosopher eats
- the time it takes for the philosopher to eat
- some delay


```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Philosopher is a struct which stores information about a philosopher.
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// philosophers is list of all philosophers.
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// Define a few variables.
var hunger = 3                  // how many times a philosopher eats
var eatTime = 1 * time.Second   // how long it takes to eatTime
var thinkTime = 3 * time.Second // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")

}
```

the important functions are `dine()` (the driver function) and `diningProblem()` goroutine. we have a waitGroup to count the number of hungry philosophers (when it's down to zero, we're done), a waitGroup to wait until all philosophers are seated at the table, and a map of mutexes for each fork.\


```go
func dine() {
	// wg is the WaitGroup that keeps track of how many philosophers are still at the table. When it reaches zero, everyone is finished eating and has left. We add 5 (the number of philosophers) to this wait group.
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// We want everyone to be seated before they start eating, so create a WaitGroup for that, and set it to 5.
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks. Forks are assigned using the fields leftFork and rightFork in the Philosopher type. Each fork, then, can be found using the index (an integer), and each fork has a unique mutex.
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// Start the meal by iterating through our slice of Philosophers.
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}
```
each philosophers is represented by a goroutine, which when completed, means the philosopher finished eating.

```go
// diningProblem is the function fired off as a goroutine for each of our philosophers. It takes one philosopher, our WaitGroup to determine when everyone is done, a map containing the mutexes for every fork on the table, and a WaitGroup used to pause execution of every instance of this goroutine until everyone is seated at the table.
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
}
```

</details>

### Implementing the `diningProblem` Logic

<details>
<summary>
Doing the actual work
</summary>

now we want the philosophers to actually eat. we first have them seated (which means all goroutines have started). and now each philosopher can try grabbing the forks (there is a specific order for this to happen to avoid a logical deadlock).


```go
func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	
	// Decrement the seated WaitGroup by one.
	seated.Done()

	// Wait until everyone is seated.
	seated.Wait()

	// Have this philosopher eatTime and thinkTime "hunger" times (3).
	for i := hunger; i > 0; i-- {
		// Get a lock on the left and right forks. We have to choose the lower numbered fork first in order to avoid a logical race condition, which is not detected by the -race flag in tests; 
		// if we don't do this, we have the potential for a deadlock, since two philosophers will wait endlessly for the same fork.
		// Note that the goroutine will block (pause) until it gets a lock on both the right and left forks.
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}
		
		// By the time we get to this line, the philosopher has a lock (mutex) on both forks.
		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		// The philosopher starts to think, but does not drop the forks yet.
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		// Unlock the mutexes for both forks.
		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	// The philosopher has finished eating, so print out a message.
	fmt.Println(philosopher.name, "is satisfied.")
	fmt.Println(philosopher.name, "left the table.")
}
```

</details>

### Challenge 2: Printing Out The Order in Which the Meal is Finished

<details>
<summary>
Checking What We learned
</summary>

the challenge is to take the existing code, and modify it so that we keep a record of the order in which they finished eating, so we need to lock something.

my code is in "\challenges\challenge-2\main.go".

1. I created a slice of strings, and checked for each iteration if this was the final one, and if so, pushed the name into the array, however, this can lead to a race condition, if two philosophers need finished at the same time and aren't sharing a fork.
2. instead, I created a new mutex, which is taken after the hunger loop, and appends the name of the philosopher to it.

this is the same as the solution provided in the course.
</details>

### Writing a Test for Our Program

<details>
<summary>
adding some tests.
</summary>

we add a test file that tests the `dine` function. we test it with zero delay values, and with varying delays as well.

</details>

</details>

## Section 5: Channels, and Another Classic: The Sleeping Barber Problem

<details>
<summary>
More experiments with channels
</summary>

More focus on <golang>channels</golang>, which are a means to communicate with and from goroutines. they can buffered or unbuffered. a channel must be closed after use. the channel only accepts one kind of data.\
later, we will implement the sleeping barbershop problem, and even go further with it!

### Introduction to Channels

<details>
<summary>
Simple Channels program
</summary>

a simple application to demonstrate the use of channels. we create a ping-pong application, which will communicate with two channels, and print one response for each input from the user, until the user quits the problem by supplying the "Q" input.\
The user writes a string, and the program writes it back in upper-case.

```go
package main

import (
	"fmt"
	"strings"
)

// shout has two parameters: a receive only chan ping, and a send only chan pong.
// Note the use of <- in function signature. It simply takes whatever string it gets from the ping channel, 
// converts it to uppercase and appends a few exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		// read from the ping channel. Note that the GoRoutine waits here -- it blocks until something is received on this channel.
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels. Ping is what we send to, and pong is what comes back.
	ping := make(chan string)
	pong := make(chan string)

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			// jump out of for loop
			break
		}

		// send userInput to "ping" channel
		ping <- userInput

		// wait for a response from the pong channel. 
		// Again, program blocks (pauses) until it receives something from that channel.
		response := <-pong

		// print the response to the console.
		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")

	// close the channels
	close(ping)
	close(pong)
}
```

notice that at the end of the program we close both channels. we do this to avoid resource leaks.

we can define our channels as being "receive only" or "send only". we simply put the `<-` before or after the <golang>chan</golang> keyword. in this case "ping" is a receive-only channel, and pong is "send-only".\

```go
func shout(ping <-chan string, pong chan<- string) {
	// ....
}
```

reading from a channel has an optional parameter, a boolean that tells us if the channel is healthy. if the value is false, then the channel was already closed.

</details>

### The `select` Statement

<details>
<summary>
The select statement with channels
</summary>

we have two function (which we will use a goroutines) that will write to a channel, the main function will listen to the channels with a <golang>select</golang> statement. we can read from the same channel in multiple statements,  and either of them can be chosen. we can have a default case to avoid a deadlock, but it can lead to an busy-waiting and wasting CPU.

```go
package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for {
		select {
		// because we have multiple cases listening to the same channels, random ones are selected
		case s1 := <-channel1:
			fmt.Println("Case one:", s1)
		case s2 := <-channel1:
			fmt.Println("Case two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case three:", s3)
		case s4 := <-channel2:
			fmt.Println("Case four:", s4)
		// default:
			// avoiding deadlock
		}
	}
}
```

</details>

### Buffered Channels

<details>
<summary>
channels that hold more than a single piece of data.
</summary>

a channel can contain more than one value. this makes the channel able to hold more values: `ch := make(chan int, 10)`. while there is space in the buffer, we don't need to wait for someone to read from the channel before writing into it.

```go
package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		// the first 10 times through this loop, things go quickly; after that, things slow down.
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done!")
	close(ch)
}
```
</details>

### The Sleeping Barber Project

<summary>
Another problem by Dijkstra.
</summary>
<details>


> - A barber goes to work in a barbershop with a waiting room with a fixed number of seats.
> - If no one is in the waiting room, the barber goes to sleep.
> - When a client shows up, if there are no seats available, they leave.
> - If there is a seat available, and the barber is sleeping, the client wakes the barber up and gets a haircut.
> - If the barber is busy, the client takes a seat and wait their turn.
> - Once the shop closes, no more clients are allowed in, but the barber has to stay until everyone who is waiting gets a haircut.

the point of this program in this course is that we don't need to use primitives. we can use channels instead.\
As always, we first define the outline for the application.

```go
package main

// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem which illustrates the complexities that arise when there are multiple operating system processes.
// Here, we have a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is open, and clients arriving at (roughly) regular intervals.
// When a barber has nothing to do, he or she checks the waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to sleep until a new client arrives.


// variables

func main() {
	// seed our random number generator

	// print welcome message

	// create channels if we need any

	// create the barbershop

	// add barbers

	// start the barbershop as a goroutine

	// add clients

	// block until the barbershop is closed
}
```

#### Defining some Variables, the Barbershop, and Getting Started with the Code

we define some data structures, such as BarberShop (the course has it in a different file), and we assign values to our variables:

- seating capacity
- customer arrival rate
- duration to cut each haircut

we start with seeding the random number generator and we write some colored text.\
we want some channels, one for clients - buffered to allow more than one. and a channel that marks that we are done for the day.

```go
import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// different file
type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

// variables
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity: seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan: clientChan,
		BarbersDoneChan: doneChan,
		Open: true,
	}

	color.Green("The shop is open for the day!")

	// add barbers

	// start the barbershop as a goroutine

	// add clients

	// block until the barbershop is closed
}
```

next time we'll define a barber as a goroutine.

#### Adding a Barber

now we define the `addBarber` function, with the BarberShop pointer as the receiver. it itself launches a goroutine, but the calling code doesn't launch this such.\
there is a loop with a break condition, and we listen on a channel with the second parameter (which will be false once the channel is closed). when the shop is closed and there are no customers in the channel, we decrease the number of barbers in the shop. if there are no customers, the barber goes to sleep, and will sleep until there is a customer or the shop closes.

```go
func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// if there are no clients, the barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, so send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}
```
#### Starting the Barbershop as a Goroutine

In the main function, we add a goroutine that will wake up a after a specified time and close the shop. we use the <golang>time.After()</golang> function for this. we have a channel saying the shop is in the process of closing, and a channel saying the shop is closed.

```go
func main() {
	// seed our random number generator
	rand.Seed(time.Now().UnixNano())

	// print welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if we need any
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create the barbershop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")

	// add barbers
	shop.addBarber("Frank")

	// start the barbershop as a goroutine
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// add clients

	// block until the barbershop is closed

	time.Sleep(5 * time.Second)
}
```

more interesting, we add a function on the barbershop to close the shop and wait for all the remaining barbers in the shop to be done. then we close the channel and print a message.

```go
func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("---------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
}
```

#### Sending Clients to the Shop

time to add clients, people who would go to the shop and get a haircut. in the main function, we add a goroutine. if the shop is closing or closed, the goroutine finishes (returns), if not, for each random period of time (defined by the arrival rate), we call a new function to add a client. we also replace our sleep delay at the end of the program with a read from the barbershop closed channel.

```go
func main() {
	//..

	// add clients
	i := 1

	go func() {
		for {
			// get a random number with average arrival rate
			randomMillseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
}
```

if there is a client, we also need to check the waiting room capacity. if the selector matches the channel, the client is added to waiting queue, if the channel is full, then we match the default case is matched and we don't add the client. 

```go
func (shop *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves!", client)
	}
}
```

#### Trying Things out

we can run the program and see the output, the color really help us. we can change the seating capacity or add more barbers to work in the shop. 

</details>

</details>

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
  - `-race` - see race condition
- `go fmt`
- `go install`
- `go get`
- `go test`
  - `-race` - check for race condition
- `go mod`
  - `init` - start a new mod file
  - `tidy`
  - `vendor`

### The Sync Package

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>

[documentation](https://pkg.go.dev/sync)

<golang>synchronization stuff</golang>

- pass waitGroup variables by reference (pointer), not by copy.
- if the waitGroup value goes below 0, we get an error.
</details>


</details>

