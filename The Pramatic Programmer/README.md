# The Pragmatic Programmer

Andrew Hunt, David Thomas

[PragmaticProgrammer](http://www.pragmaticprogrammer.com)

a programming book by programmers for programmers.

pragmatic programmers are

> - Early adopters / fast adopters
> - Inquisitive
> - Critical thinkers
> - Realists
> - Jack of all trades

here are some tips

> - care about your craft
> - Think! about your work

Individualism in large teams.\
Daily work, it's a continues process.

## Chapter 1 - A Pragmatic Philosophy

<details>
<summary>
//todo?
</summary>

### The Cat Ate My Source Code

Take responsability. if you're in charge, be in charge.

### Software Entropy

Don't live with "broken windows" - fix issues when they are small, don't let problem persist, if you can't fix them, remove them, make sure that stuff are being handled and not neglected.

### Stone Soup and Boiled Frogs

be a catalyst for a change. lead and others will follow. start small and let other help. avoid 'start-up fatigue' (won't a better name be 'start up anxiety'?).

### Good Enough Software

Be aware and be communicative about the trade-offs you're doing. different projects have different standards for what counts as 'quality' or 'good enough', not every piece of software runs pacemakers or spaceships. sometimes it's better to ship a program with some bugs today then delay it and release a 'perfect' (which means better, we'll never run out of bugs) version next year. and eventually, we have to stop working on a software, there is a point where fixes and additions are just harming us.

### Your Knowledge Portfolio

knowledge and experience are the best guides we have, but they decay, things change, best practices come and go, and changes in environment might mean that what was good last year is no longer recommended.

to keep our knowledge portfolio relevenat and valuable, we can treat in like our financial investments.

1. invsitage, regularly.
2. diversify.
3. balance between high risk - high reward investments and conservative ones.
4. buy low, sell high
5. periodically review and rebalanced the portfolio.

some goals that we can follow to make sure we are doing those things correctly. we should set goals and time limits, and more importantly, keep them.

- learning new languages
- reading technical books
- reading none-technical books - keep in touch with what none-programmers humans are up to
- take classes
- participate in local user group
- experiment with different environments
- stay current with professional magazine and knowledge sources
- get wired and seek out information that hasn't been put into the standardized knowledge bases yet (newsgroups, mailing lists)

seek out opportunities for learning. if you find something you don't know (even if it's not needed for work), try, seek out advice, build up the network.

think critically about what you read and hear, don't fall for the hype.

### Communicate

A good idea is nothing without communication. write memos, proposals, status reports, new ideas and approches.

know what you want to say. start with an outline, be sure that you cover all the points you want to get across.
know your audience. be aware of the other sides needs and intrests.
choose the settings, when and where to make suggestions.
have a style of presentation appropriate to the audience. don't forget about presentation, good looking documents are more likely to be read.
get your audience involved, listen to suggestions, incorporate input.
listen to others, encourage questions, promote dialog.
get back to people, keep the connection alive, respond, keep the other side informed.

</details>

## Chapter 2 - A Pragmatic Approach

<!-- <details> -->
<summary>
//todo
</summary>

more grounded section about pragmatism.

### The Evils of Duplication

code maintenance doesn't begin at release, it starts whenever you start writing.\
the dry principle

> "Every piece of knowledge must have a single, unambiguous, authoritative representation within a system"

duplication can arise from several sources

> - Imposed duplication
> - Inadvertent duplication
> - Impatient duplication
> - Interdeveloper duplication

#### Imposed duplication

sometimes it seems that we can't avoid duplication, comments are a form of duplication (the code is already there, and the comments aren't updated, and go out of sync). multiple platforms with different languages. client/server code.\
we can try to make our code build from the same source, such as a schema or metadata, so class defintions are kept up-to-date and don't require manual work. comments that explain the 'how' are duplication, they should explain the 'why', unless there is a special reason. we can have duplication from the language, like header files and defintions in cpp,

#### Inadvertent duplication

issues in the design that cause duplication. maybe the same value exists in several classes and when it changes we need to update all of them, maybe a property of class is dependent on other properties but isn't generated from them (rectangle with height, width and area as members). we might decide to store this value for performance, but we need to make this decision knowingly

#### Impatient duplication

taking the easy way, duplicating code because it's easier, it's just a small thing...
but remember, "short cuts make for long delays".

#### Interdeveloper duplication

many teams, all doing the same thing in differnet ways. teams that end up developing the same functionality across subsystems,
we need to search for duplicated code, try and see if we can grab code from other rather than re-write it ourselves, etc..

### Orthogonality

in geometry, lines are orthogonal if the meet at the right angle (90 degrees). like the X and Y axis. in vector terms, if we move along one line, it doesn't change how our position is projected unto the other line. in coding, the term has come to mean independence or decoupling, in a well designed system, changes to one shouldn't effect other. if we change the database code, we don't need to change the interface code, and vice-versa.

non orthogonal systems are such that change in one property can require changes in other. non-orthogonal systems are more complicated to control, harder to change, and there is no such thing as 'local fix' on an orthogonal system.
orthogonal systems are self contained, independent,and have a single purpose (also called cohesion). components are isolated from one another, and communication between subsystem is clearly defined by an external interface.
this means we can do localized changes, components are small, simple, isolated, and easy to test. they can also be reused across the system. the risk is also contained, the system is more flexile and isn't fragile, and it probably isn't tied to a specific vendor or external platform.

there's also orthogonality in teams, if each person has a domain and a clear goal, they can work on it alone, without stepping on other developers toes. we should separate infrastructure from application. in design language, we call orthogonality 'modularity','component based' or 'layered', but it's all the same thing, we have distinct areas who only connect to one another through abstraction.

a change in one module should only affect that module.
using external libraries and toolkits can effect orthogonality, the way the library behaves might impose restrictions on you.

ways to conserve orthogonality when coding:\

- keep code decoupled, don't reveals anything unnecessary to other modules, don't rely on other modules implementation (use only the public interface), and if we want to change a state of the object, get the object itself to do it.
- avoid global data, don't tie yourself to data that others can read or change. it's better to explicitly pass the context than rely on global data. be careful of singletons, they tend to cause unnecessary linkage.
- avoid similar functions, if there are too many function that look similar, do something about it.extract, dry, template, compose, whatever it takes. be critical of your code.

orthogonality makes testing easier. it allows smaller unit tests to be run in isolation. but tests can also be orthogonal to the code, how related is the testing project to the actual project? does it need to import everything? why? does it rely on inner data from the code? how hard is it to run the tests?

orthogonality is also part of documentation, separate appearance and content.

if the DRY principle is about reducing duplication in the code, then orthogonality is about reducing interdepency. the two goals work together.

#### challenges and questions

open questions:

- command line stuff are more flexible, of course, we can pipe into them, out of them, set them anywhere, log the results, run them automatically...
- multiple inheritance is bad, is adds more complications, interfaces lessen the problem, as they don't introduce extra data, a interface is a public api, inheritance is an additional set of problems someone forces you to carry.

questions

1. Split2 looks more orthogonal. why do we need to call a void function without any parameters?
2. no idea what this means
3. OOP probably, just because they do encapsulation better (at least it was once true)

### Reversibility

the world isn't stable, nothing is set in stone, things will change. constants, requirements, vendors. as time goes by, more decisions are made that limit the available options for us as programmers. many critical decisions are easily reversible.

many of the suggestions from the earlier parts of the chapter support reversability, if our code is decoupled, orthogonal, and not duplicated, this can mitigate the costs of reversing decisions and changing course. one module that wraps around the database means that changing database vendors only effects that module,having clearly separated layers in a client-server model can help us deploy a stand alone version of the program if we need.

nothing is set in stone, all decisions are only temporary final. and there is no such thing as a final decision.

it's not just code that needs to be flexible, we should also strive for flexible architecture, deployment and vendor integration. we can use stuff like COBRA (is it used today???) to bring together components from different languages. we need to be ready for changes, and if we can't completely lock off each module in it's independent world, we should have a way to automatically find those weak points.

### Tracer Bullets

in ammo, tracer bullets tell us where the bullets go, they are bright and visible, so it helps us aim and we can see if we're are getting off target.

when we have a new project, one response to uncertainty is to requests specifications, more and more. a different response it to start moving and correct course, just like shooting tracer bullets. build code in a way that gets you closer to the target, even if the final target isn't yet unknown, check and validate the things you can do, and flesh it out gradually. don't get rid of those intermediate steps, they are the road blocks that any future change should follow.

advantages:

- users get to see something working early: there is always something new to present, progress is visible, which makes everyone involved more motivated.
- developers build a structure to work in: there is a basis of code to start from, changes can be made, code can be expanded.
- you have an integration platform: if we connect the system end-to-end early, we can see that it's being built correctly, no surprises of incompatible libraries, missing modules or other stuff at the last minute. integration is a continues progress, and problems are discovered early.
- you have something to demonstrate: higher ups like to see results
- you have a feel of progress: you can tell how much of the project was completed and how much more is left.

of course, our initial guesses aren't perfect, and we might be missing the target. that's the point. find those problems early, get feedback, discover performance blocks, nick the problems in the bud before they get bigger and the deadline is close.

Tracer code isn't prototyping. prototyping isn't part of the final product, it's a demonstartion of the expected final version product. tracer code is exactly the opposite, it's how things will work in practice, it's simple and it's operational, if the world ends tomorrow, this is the version we shipped. prototype is disposable code, tracer code is the skelton of the code.

### Prototypes and Post-it Notes

### Domain Languages

### Estimating

</details>

## Take Away Points
