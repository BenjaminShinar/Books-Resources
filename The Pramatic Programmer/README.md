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
CODING
(continue later)

### Reversibility

### Tracer Bullets

### Prototypes and Post-it Notes

### Domain Languages

### Estimating

</details>

## Take Away Points
