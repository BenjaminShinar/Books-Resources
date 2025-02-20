<!--
// cSpell:ignore Sawler gotemplate fatih randomMillseconds taskkill
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

<details>
<summary>
Starting the final Project.
</summary>

A fake application that sells subscriptions, with an API and backend.

### Setting Up A Simple Web Application

<details>
<summary>
The Project Outline and getting imports
</summary>

we start with setting the project as a web application. we set it up in a new folder "final-project" and create the structure.

```sh
mkdir final-project
cd final-project
go mod init final-project
mkdir cmd
mkdir cmd/web
touch cmd/web/main.go
```

we start with the main file of the package - "main.go", we define constants with <golang>const</golang> and write the outline of the program.


- database connection
- <golang>channels</golang>
- <golang>waitGroups</golang>
- application config
- listen to web connections
- manage sessions (user logins)
- set up mail

```go
package main

const webPort = "80"

func main() {
	// connect to the database

	// create sessions

	// create channels

	// create waitGroup

	// set up the application config

	// set up mail

	// listen for web connections
}
```

for the database, we will use PostgresSQL. for session management we will use the package by Alex Edwards with redis store and the go-chi routing package. we can download the packages with `go get`.


```sh
go get github.com/jackc/pgconn
go get github.com/jackc/pgx/v4
go get github.com/alexedwards/scs/v2
go get github.com/alexedwards/scs/redisstore
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
```

</details>

### Setting Up Our Docker Development Environment

<details>
<summary>
Adding containers using docker compose
</summary>

since we're using packages for Postgres and Redis, we can run them as docker containers. we add a docker compose file to our project. mailhog is a dummy mail server.

```yaml
version: '3'

services:
  # start Postgres, and ensure that data is stored to a mounted volume
  postgres:
    image: 'postgres:14.2'
    ports:
      - "5432:5432"
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: concurrency
    volumes:
      - ./db-data/postgres/:/var/lib/postgresql/data/

  #  start Redis, and ensure that data is stored to a mounted volume
  redis:
    image: 'redis:alpine'
    ports:
      - "6379:6379"
    restart: always
    volumes:
      - ./db-data/redis/:/data

  #  start mailhog
  mailhog:
    image: 'mailhog/mailhog:latest'
    ports:
      - "1025:1025"
      - "8025:8025"
    restart: always
```

we also create some folder to store the data.

```sh
mkdir db-data
mkdir db-data/postgres
mkdir db-data/redis
```

now we can run the `docker-compose up -d` to spin up the containers. this will also populate the folders with data. we can also use a database client like BeeKeeper to view the data.
</details>

### Adding Postgres

<details>
<summary>
Connecting to the database.
</summary>

now we connect to the database from our go application. we first make sure we can connect via our client. once we make sure our database is up, we start adding functions to connect to the database, we set 10 retries to connect to the database.

```go
func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
```

we also need to import the packages, we use `_` to import without using in the code.

</details>

### Setting Up A Makefile

<details>
<summary>
Making our life easier using makefile.
</summary>

we can use a makefile (different ones for linux and windows). we store them in the project root folder.

> - `@` suppresses the normal 'echo' of the command that is executed.
> - `-` means ignore the exit status of the command that is executed (normally, a non-zero exit status would stop that part of the build).
> - `+` means 'execute this command under make -n' (or `make -t` or `make -q`) when commands are not normally executed.

```makefile
BINARY_NAME=my_app
PG_PASS=
DSN="host=localhost port=5432 user=postgres password=${PG_PASS}  dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5"
REDIS="127.0.0.1:6379"

## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/web
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env DSN=${DSN} REDIS=${REDIS} ./${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...
```

and for windows
```makefile
PG_PASS=
DSN=host=localhost port=5432 user=postgres password=${PG_PASS} dbname=concurrency sslmode=disable timezone=UTC connect_timeout=5
BINARY_NAME=my_app.exe
REDIS="127.0.0.1:6379"

## build: builds all binaries
build:
	@go build -o ${BINARY_NAME} ./cmd/web
	@echo back end built!

run: build
	@echo Starting...
	set "DSN=${DSN}"
	set "REDIS=${REDIS}"
	start /min cmd /c ${BINARY_NAME} &
	@echo back end started!

clean:
	@echo Cleaning...
	@DEL ${BINARY_NAME}
	@go clean
	@echo Cleaned!

start: run

stop:
	@echo "Stopping..."
	@taskkill /IM ${BINARY_NAME} /F
	@echo Stopped back end

restart: stop start

test:
	@echo "Testing..."
	go test -v ./...
```

we can run `make start` to run the application and see that we connect to the postgres application.

</details>

### Adding Sessions & Redis

<details>
<summary>
Connecting To Redis for session management.
</summary>

we next want to connect to Redis and set up the user sessions. we create function to create a session manager, and one to connect to Redis. we also set up an inline function using environment variables, which we can define in the makefile.\
when we create sessions, we set the storage to redis and define some nice defaults.

```go

func initSession() *scs.SessionManager {
	// set up session
	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}

	return redisPool
}
```
</details>

### Setting Up The Application Config

<details>
<summary>
adding a configuration type.
</summary>

we create a new file "config.go" in the main package. this file contains the application configuration, starting with the session manager, database manager, loggers and a waitGroup pointer.\
we will add more to it as we go.

```go
package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session  *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
}
```

and we update our main file to populate the configuration file with the data we already have. we also setup the loggers, pointing them to the standard out with a prefix and additional data (time and source code location).

```go
func main() {
	// connect to the database
	db := initDB()

	// create sessions
	session := initSession()

	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// create channels

	// create waitGroup
	wg := sync.WaitGroup{}

	// set up the application config
	app := Config{
		Session:  session,
		DB:       db,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
	}

	// set up mail

	// listen for web connections
}
```
</details>

### Setting Up A Route & Handler For The Home Page, And Starting The Web Server

<details>
<summary>
Setting up the server and handlers.
</summary>

our application still doesn't work. we have unused variables and we are missing handlers and we aren't listen to wee connections.

in the "cmd/web" folder, we add a "handlers.go" file, and a "routes" file

```sh
touch cmd/web/handlers.go
touch cmd/web/routes.go
```

in the handlers file, we set up a receiver function on the config type, it will take a response writer and a request.

```go
package main

import "net/http"

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	
}
```

in the routes file, we set up a <golang>mux</golang> router from the <golang>go-chi</golang> package, which will work the middleware and set routes for the paths. the root path will return the homepage.

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)

	// define application routes
	mux.Get("/", app.HomePage)

	return mux
}
```

finally, we add a function to "main.go"  to listen and serve web pages. it will listen on the port 80, and we use the routes function as the server handler.

```go
func (app *Config) serve() {
	// start http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.InfoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
```

at this point, we should be able to run the application, but we don't render a response.
</details>

### Setting Up Templates And Building A Render Function

<details>
<summary>
render a template http response.
</summary>

now we want to have our application serve web pages for us. we have some templates in the source code, which we add to a new folder.

```sh
mkdir cmd/web/templates
```

they are gohtml files which use go-syntax to evaluate into html pages. the "base.layout.html" defines the html structure, with other files handling more stuff like header, footer, navigation, alerts. and allowing for different content

- `template` - call a sub-template
  - `include` - is helm extension
  - `block` - like template, but with defaults
- `define` - sub-template (acts like a function)

we add a new file "render.go" which will handle general rendering of template files.\
inside this file we have the templates path as a variable (so we could change it for testing), and a general struct for passing data into templates. this includes general data and user specific data.\
We start with the render function, taking a response writer, the requests, the name of the template, and the template data struct. we always need to have some templates, no matter which page we're rendering. this files we be parsed and render using the <golang>http/template</golang> package.

```go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var pathToTemplates = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	// User *data.User
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplates),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))

	// put the defaults and the specific template in one range.
	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	// create it empty
	if td == nil {
		td = &TemplateData{}
	}

	// parse the templates into a single object.
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// execute (populate) with data.
	if err := tmpl.Execute(w, app.AddDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
fill up notifications and user specific data.
*/
func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	if app.IsAuthenticated(r) {
		td.Authenticated = true
		// TODO - get more user information
	}
	td.Now = time.Now()

	return td
}

/*
check if session contains user id
*/
func (app *Config) IsAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}

```

and finally, at the "handlers.go" file, we can serve the homepage template.

```go
func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}
```

we still don't have session information, so that's the next step.
</details>

### Adding Session Middleware

<details>
<summary>
Adding the session middleware to the router.
</summary>

we add a new file "middleware.go", like all middleware,it takes an http handler, modifies it and returns.

```go
package main

import "net/http"

func (app *Config) SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
```

we add it to the router we created in "routes.go", simply by telling the router object to use it.

```go

func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)

	return mux
}
```

at this point, we should see an actual web page when we navigate to the web application.
</details>

### Setting Up Additional Stub Handlers And Routes

<details>
<summary>
Add empty methods for handlers and routes
</summary>

we add some stub handlers to handle other pages. this is done in the "handlers.go" and the "routes.go" files.

stubs:

```go
func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
```

registering handlers on the routes:

```go
func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate-account", app.ActivateAccount)

	return mux
}
```

</details>

### Implementing Graceful Shutdown

<details>
<summary>
Listen to Os.Events and wait for operations to finish.
</summary>

when our application stops, we want to wait for actions to complete, such as sending emails, creating invoices, etc..\
we run goroutine that listens on the interrupt system calls <golang>Os.Signal</golang> as a channel, and blocks while waiting for the signals on the channel. when it receive the signals we defined, it will wake up, perform some actions and the quit the program. in our case, we will delay shutdown until the waitGroup is empty.

```go
func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// block until waitGroup is empty
	app.Wait.Wait()

	app.InfoLog.Println("closing channels and shutting down application...")
}
```
</details>

### Populating The Database

<details>
<summary>
Filling in some dummy data.
</summary>

from the course resource files, we can take the "db.sql" file. it contains sql commands to create our postgres database tables and insert a dummy user. 

- `CREATE TABLE`
- `ALTER TABLE`
- `CREATE SEQUENCE`
- `INSERT INTO`

</details>

### Adding A Data Package And Database Models

<details>
<summary>
Interfacing with the data in the database.
</summary>

we add some source files to a new folder "data", this is a package that is consumed by the application.

- "user.go" - defines the user data object as a go struct with CRUD operations
- "plan.go" - plan go struct, sql commands, utilities
- "models.go" - exposes the tables from the database connection

we update the Config struct with the new models type in "config.go.

```go
type Config struct {
	Session  *scs.SessionManager
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Wait     *sync.WaitGroup
	Models   data.Models
}
```

in the "main.go" file, we instantiate the struct with the call 
```go
func main() {
	// code before
	// set up the application config
	app := Config{
		Session:  session,
		DB:       db,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
		Models:   data.New(db),
	}
	// code after
}
```
</details>


### Implementing The Login/Logout Functions

<details>
<summary>
implement logic about log-in.
</summary>

back in the "handlers.go" file, we want to implement to logic of user log-in and log-out.\
at the start of each operation, we renew the session token (using the context), we don't care about the result.\
For log-in, our request is the http form with email and password. we search for the user in the database based on the email, and then check the user object if the password matches. if we get errors, we can put them into the session object. if our login data matches, we push the user data into the session.\
before we store the data, we need to register this type with the session by calling `gob.Register(data.User{})`

```go
func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	// parse form post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword{
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// okay, so log user in
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)

	app.Session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

the logout is simpler, we destroy the user session to remove all stored data from the context, renew the session and redirect the user.

```go
func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	// clean up session
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```

</details>


</details>

## Section 7: Sending Email Concurrently

<details>
<summary>
Sending emails from our application.
</summary>

when a user registers, we want to do some stuff, and this will run in the background, so here our concurrency comes into place. this will be done through sending emails and using channels.\
we will also add logic to wait for emails to finish sending before shutting down the app.

```go
func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	// create a user

	// send an activation email

	// subscribe the user to an account
}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url 

	// generate an invoice

	// send an email with attachments

	// send an email with the invoice attached
}
```

### Getting Started With The Mailer Code

<details>
<summary>
Sending emails from go.
</summary>

we will implement the mail sending logic by using <golang>goroutines</golang>, in a new file called "mailer.go" inside "cmd/web" folder.

we will add some packages using `go get`
- mailing package
- mail styling package

```sh
go get github.com/vanng822/go-premailer/premailer
go get github.com/xhit/go-simple-mail/v2
```

we start with synchronous mailing. first thing are the types, the mail sender object and the message object. the mail object has a <golang>waitGroup</golang> and some <golang>channels</golang> - messages, errors and "done".\
The message struct doesn't use concurrency, but will have data of type <golang>any</golang>.

```go
type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
	Wait        *sync.WaitGroup
	MailerChan  chan Message
	ErrorChan   chan error
	DoneChan    chan bool
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	Template    string
}
```

we next want a function to send the messages themselves. we can have a template for special kinds of mails, if the message doesn't specify one of the fields, we will use the values from the mail object.\
our function can handle both simple plain text messages or html messages, each using one of the <golang>gohtml</golang> files from the templates folder (they will not use the chain of templates parsing like the web pages). *this will be covered in the next lecture*.

Sending the mails will be done with the *server* object from the <golang>go-simple-mail</golang> package. An email message can have different kinds of encryptions, depending on what the external server supports. there are some other parameters to define on the mail server object and on our email message object. the email has the body in plaintext and alternative body which it the html formatted email. we can also add attachments.

If we fail in one of our utility functions, we will send the error into the error channel.

```go
func (m *Mail) sendMail(msg Message, errorChan chan error) {
	if msg.Template == "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}

	msg.DataMap = data

	// build html mail
	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		errorChan <- err
	}

	// build plain text mail
	plainMessage, err := m.buildPlainTextMessage(msg)
	if err != nil {
		errorChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		errorChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)

	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, x := range msg.Attachments {
			email.AddAttachment(x)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		errorChan <- err
	}
}

func (m *Mail) getEncryption(e string) mail.Encryption {
	switch e {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
```
#### Building Html And Plain Text Messages

we finish the message body input and formatting, and add some css. we use the same template behavior as we did with the web pages. for the html message we also inline the css style using the premailer package.

```go
func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./cmd/web/templates/%s.html.gohtml", msg.Template)

	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) buildPlainTextMessage(msg Message) (string, error) {
	templateToRender := fmt.Sprintf("./cmd/web/templates/%s.plain.gohtml", msg.Template)

	t, err := template.New("email-plain").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err = t.ExecuteTemplate(&tpl, "body", msg.DataMap); err != nil {
		return "", err
	}

	plainMessage := tpl.String()

	return plainMessage, nil
}

func (m *Mail) inlineCSS(s string) (string, error) {
	options := premailer.Options{
		RemoveClasses: false,
		CssToAttributes: false,
		KeepBangImportant: true,
	}

	prem, err := premailer.NewPremailerFromString(s, &options)
	if err != nil {
		return "", err
	}

	html, err := prem.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}
```
</details>

### Sending A Message (Synchronously)

<details>
<summary>
Testing that mails are sent.
</summary>

we add a test route that will send an email synchronously. this is just temporary for testing and will be removed shortly. we define an inline function with the mail server and the message, and having it send the mail.

```go
func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate-account", app.ActivateAccount)

	mux.Get("/test-email", func(w http.ResponseWriter, r *http.Request) {
		m := Mail{
			Domain: "localhost",
			Host: "localhost",
			Port: 1025,
			Encryption: "none",
			FromAddress: "info@mycompany.com",
			FromName: "info",
			ErrorChan: make(chan error),

		}

		msg := Message{
			To: "me@here.com",
			Subject: "Test email",
			Data: "Hello, world.",
		}

		m.sendMail(msg, make(chan error))
	})

	return mux
}
```
we can test the functionality by navigating to the new route and checking the mailhog web interface to see the send message.

</details>

### Getting Started Sending A Message (Asynchronously)

<details>
<summary>
sending mails in the background.
</summary>

now that we know our application can send mails, it's time to have them send in the background.\
we define a function on the application config object and use the <golang>select</golang> statement to listen on channels. we add the mailer to this object so we could listen on the three channels:

- mails to send
- errors from sending mails
- mailing is finished

```go
func (app *Config) listenForMail() {
	for {
		select {
		case msg := <- app.Mailer.MailerChan:
			go app.Mailer.sendMail(msg, app.Mailer.ErrorChan)
		case err := <- app.Mailer.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.Mailer.DoneChan:
			return
		}
	}
}
```

back in our main function, we create the mailer and listen for the mailing in a <golang>goroutine</golang>. we need to create the mailer object and the channels in it. we will allow the channel to hold up to 100 messages before blocking. we will re-use the same <golang>waitGroup</golang>.

```go
func main() {
	// code
	// set up the application config
	app := Config{
		Session:  session,
		DB:       db,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		Wait:     &wg,
		Models:   data.New(db),
	}

	// set up mail
	app.Mailer = app.createMail()
	go app.listenForMail()

	// code
}

func (app *Config) createMail() Mail {
	// create channels
	errorChan := make(chan error)
	mailerChan := make(chan Message, 100)
	mailerDoneChan := make(chan bool)

	m := Mail{
		Domain: "localhost",
		Host: "localhost",
		Port: 1025,
		Encryption: "none",
		FromName: "Info",
		FromAddress: "info@mycompany.com",
		Wait: app.Wait,
		ErrorChan: errorChan,
		MailerChan: mailerChan,
		DoneChan: mailerDoneChan,
	}

	return m
}
```

now we also need to decrement the wait group in the send mail function, using `defer m.Wait.Done()` so our application could finish.

#### Writing A Helper Function To Send Email Easily

we want the <golang>waitGroup</golang> to control how many tasks are in process, so we add a helper file with some utility wrappers. like incrementing the wait group as needed and pushing the message into the channel.

```go
func (app *Config) sendEmail(msg Message) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <-msg
}
```
</details>

### Sending An Email On Incorrect Login

<details>
<summary>
sending a mail when login fails
</summary>

now that we have the capability to send emails, we can use it as part of our application flow.

in the "handlers.go" file, we can send an email if the password is not correct. we create the message with the user email and use the helper function to send the email.

```go
func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	// more code

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword{
		msg := Message{
			To: email,
			Subject: "Failed log in attempt",
			Data: "Invalid login attempt!",
		}

		app.sendEmail(msg)
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	// more code
}
```
</details>

### Adding Cleanup Tasks To The `shutdown()` Function

<details>
<summary>
Adding to the cleanup task.
</summary>

back in the main.go file, we had a function that captured the system interrupts and performed cleanup functions. we need to send a message to the "doneChan" to stop listening for emails, and we need to close all of the channels.

```go
func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// block until waitGroup is empty
	app.Wait.Wait()

	app.Mailer.DoneChan <- true

	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
}
```

</details>

</details>

## Section 8: Registering a User and Displaying Plans

<details>
<summary>
Moving forward with the application.
</summary>

we continue with the application, focusing on the users and plans. the application will send activation emails, which will be signed.

### Adding Mail Templates And Url Signer Code

<details>
<summary>
Mail templates
</summary>

when a user creates an account in our application, we will create the account in the backend and send him an activation link in the mail.
for this, we will create go templates for the emails (plaintext and html).

the next thing to do is to protect the url, we will go the <golang>bwmarrin/go-alone</golang> package for that, which we get with `go get`. we will go back to this later.

```go
package main

import (
	"fmt"
	"github.com/bwmarrin/go-alone"
	"strings"
	"time"
)

var secret // take from env somehow

var secretKey []byte

// NewURLSigner creates a new signer
func NewURLSigner() {
	secretKey = []byte(secret)
}

// GenerateTokenFromString generates a signed token
func GenerateTokenFromString(data string) string {
	var urlToSign string

	s := goalone.New(secretKey, goalone.Timestamp)
	if strings.Contains(data, "?") {
		urlToSign = fmt.Sprintf("%s&hash=", data)
	} else {
		urlToSign = fmt.Sprintf("%s?hash=", data)
	}

	tokenBytes := s.Sign([]byte(urlToSign))
	token := string(tokenBytes)

	return token
}

// VerifyToken verifies a signed token
func VerifyToken(token string) bool {
	s := goalone.New(secretKey, goalone.Timestamp)
	_, err := s.Unsign([]byte(token))

	if err != nil {
		// signature is not valid. Token was tampered with, forged, or maybe it's
		// not even a token at all! Either way, it's not safe to use it.
		return false
	}
	// valid hash
	return true

}

// Expired checks to see if a token has expired
func Expired(token string, minutesUntilExpire int) bool {
	s := goalone.New(secretKey, goalone.Timestamp)
	ts := s.Parse([]byte(token))

	// time.Duration(seconds)*time.Second
	return time.Since(ts.Timestamp) > time.Duration(minutesUntilExpire)*time.Minute
}
```

</details>

### Starting On The Handler To Create A User

<details>
<summary>
Sending the activation email.
</summary>

in the "handlers.go" file, we have a function that's called when someone tries to register and send the form. we take the form and parse it, if this was a really application, we would also validate the data. we create a User object from the data and try to insert it into the database, if we succeed in that, we can send an activation email. the activation email will have the generated token as the activation link.\
we will also update the notification in the session info and redirect the user to the log-in page.

```go
unc (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// TODO - validate data

	// create a user
	u := data.User{
		Email: r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName: r.Form.Get("last-name"),
		Password: r.Form.Get("password"),
		Active: 0,
		IsAdmin: 0,
	}

	_, err = u.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to create user.")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// send an activation email
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedURL := GenerateTokenFromString(url)
	app.InfoLog.Println(signedURL)

	msg := Message{
		To: u.Email,
		Subject: "Activate your account",
		Template: "confirmation-email",
		Data: template.HTML(signedURL),
	}

	app.sendEmail(msg)

	app.Session.Put(r.Context(), "flash", "Confirmation email sent. Check your email.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```
</details>

### Activating A User

<details>
<summary>
Using the activation link.
</summary>

after we sent the email, the user will click on it and come to a new route in the application.

```go
func (app *Config) routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount) // new page

	return mux
}
```

so we can modify the handler to handle the request. we first verify the url that it was created from a url and was not tempered with. then we get the user from the database, and update it to be active. we can end things by creating a notification and redirecting the user to log-in.

```go

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url 
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost%s", url)
	okay := VerifyToken(testURL)

	if !okay {
		app.Session.Put(r.Context(), "error", "Invalid token.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// activate account
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	err = u.Update()
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to update user.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Account activated. You can now log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	// generate an invoice

	// send an email with attachments

	// send an email with the invoice attached

	// subscribe the user to an account
}
```
</details>

### Giving User Data To Our Templates

<details>
<summary>
Adding the user to templates
</summary>

if our user is authenticated, we want to pass the user data to our templateData, so we could use it in our templates when we serve web pages. so back in the "render.go" file, we can update the method to add the user to the template.

```go
func (app *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	if app.IsAuthenticated(r) {
		td.Authenticated = true
		user, ok := app.Session.Get(r.Context(), "user").(data.User)
		if !ok {
			app.ErrorLog.Println("can't get user from session")
		} else {
			td.User = &user
		}
	}
	td.Now = time.Now()

	return td
}
```

</details>

### Displaying The Subscription Plans Page

<details>
<summary>
Showing a webpage with the available plans.
</summary>

we have three plans in the database: bronze, silver and gold. we want the user to choose a plan to subscribe to. so we need a new handler. this page is only available for logged-in users, so we protect against that - if a user that isn't logged in tries to access us, we flash a warning and redirect them. then we take all the plans from the database and render a template with the plans in the templateData.

```go
func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	if !app.Session.Exists(r.Context(), "userID") {
		app.Session.Put(r.Context(), "warning", "You must log in to see this page!")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}

	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}
```

the template has a table, and it loops over the plans property in the template data to populate it. we also have a button for each plan, which will select the plan. our button will trigger a js script, using the sweetAlert npm package as inlined source.\
we will fire up a dialog for the user to click, if the user confirms, we will re-direct the user to a new route.

```go
{{template "base" .}}

{{define "content" }}
    <div class="container">
        <div class="row">
            <div class="col-md-8 offset-md-2">
                <h1 class="mt-5">Plans</h1>
                <hr>
                <table class="table table-compact table-striped">
                    <thead>
                        <tr>
                            <th>Plan</th>
                            <th class="text-center">Price</th>
                            <th class="text-center">Select</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range index .Data "plans"}}
                            <tr>
                                <td>{{.PlanName}}</td>
                                <td class="text-center">{{.PlanAmountFormatted}}/month</td>
                                <td class="text-center">
									<a class="btn btn-primary btn-sm" href="#!" onclick="selectPlan({{.ID}}, '{{.PlanName}}')">Select</a>
                                </td>
                            </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>

        </div>
    </div>
{{end}}

{{define "js"}}
    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11.4.14/dist/sweetalert2.all.min.js"></script>
    <script>
        function selectPlan(x, plan) {
            Swal.fire({
                title: 'Subscribe',
                html: 'Are you sure you want to subscribe to the ' + plan + '?',
                showCancelButton: true,
                confirmButtonText: 'Subscribe',
            }).then((result) => {
                if (result.isConfirmed) {
                    {{/* window.location.href = '/subscribe?id=' + x; */}}
                }
            })
        }
    </script>
{{end}}
```

</details>

### Adding A Route And Trying Things Out For The Plans Page

<details>
<summary>
choosing a plan.
</summary>

we add the "plans" route in the navigation bar template, and make it only visible for logged in users. we will also modify the plans template to display which plan the user is subscribed to.\
we can modify our database manually to add a plan to a user, so we could see the change in the ui.

#### Writing A Stub Handler For Choosing A Plan

we add the "subscribe" route and a new handler. for now, it won't do anything, but we can write down the outline of what we would like to do in the next section.

```go
func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	// get the id of the plan that is chosen

	// get the plan from the database

	// get the user from the session

	// generate an invoice

	// send an email with the invoice attached

	// generate a manual

	// send an email with the manual attached

	// subscribe the user to an account

	// redirect
}
```

</details>

</details>

## Section 9: Adding Concurrency to Choosing a Plan

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>

## Section 10: Testing

<!-- <details> -->
<summary>
//TODO: add Summary
</summary>


</details>


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
- `go mod` - modules(dependencies) 
  - go.mod and go.sum files
  - `init` - start a new mod file
  - `tidy`
  - `vendor`
- `go work` - workspace setup
  - go.work file

### The Sync Package

<details>
<summary>
Synchronization Primitives: mutexes, waitGroups and channels.
</summary>

[documentation](https://pkg.go.dev/sync)

<golang>Synchronization stuff</golang>

- pass <golang>waitGroup</golang> variables by reference (pointer), not by copy.
- decrement wait groups with deferred `wait.Done()`.
- if the waitGroup value goes below 0, we get an error.
- directional channels (as function parameters):
  - `func foo(ping <-chan string, pong chan<- string)`
  - parameter-name, chan, type
  - ping is a receiving/input channel - we "read" from it. can't close the channel from here.
  - pong is a sending/output channel - we "write" to it.

</details>

Other Packages

- [go-chi/chi](https://github.com/go-chi/chi) - routing
- [alexedwards/scs](https://github.com/alexedwards/scs) - session management
- [jackc/pgconn](https://github.com/jackc/pgconn) - postgres connection
- [vanng822/go-premailer](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in golang
- [xhit/go-simple-mail](https://github.com/xhit/go-simple-mail) - mailing service
- [bwmarrin/go-alone](https://github.com/bwmarrin/go-alone) - MAC signatures


</details>
