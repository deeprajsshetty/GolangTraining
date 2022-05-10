## Design Guidelines

You must develop a design philosophy that establishes a set of guidelines. This is more important than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't begin to make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs.

* Philosophy
    * [Prepare Your Mind](https://github.com/deeprajsshetty/GolangTraining/tree/master/001-Design%20Guidelines#prepare-your-mind)
    * [Legacy Software](https://github.com/deeprajsshetty/GolangTraining/tree/master/001-Design%20Guidelines#legacy-software)

---

### Prepare Your Mind

**Somewhere Along The Line**  
* We became impressed with programs that contain large amounts of code.
* We strived to create large abstractions in our code base.
* We forgot that the hardware is the platform.
* We lost the understanding that every decision comes with a cost.

**These Days Are Gone**  
* We can throw more hardware at the problem.
* We can throw more developers at the problem.

**Open Your Mind**  
* Technology changes quickly but people's minds change slowly.
* Easy to adopt new technology but hard to adopt new ways of thinking.

**Interesting Questions - What do they mean to you?**  
* Is it a good program?
* Is it an efficient program?
* Is it correct?
* Was it done on time?
* What did it cost?

**Aspire To**  
* Be a champion for quality, efficiency and simplicity.
* Have a point of view.
* Value introspection and self-review.

---

### Reading Code

Go is a language that focuses on code being readable as a first principle.

**Quotes**

_“If most computer people lack understanding and knowledge, then what they will select will also be lacking.” - Alan Kay_

_"The software business is one of the few places we teach people to write before we teach them to read." - Tom Love (inventor of Objective C)_

_"Code is read many more times than it is written." - Dave Cheney_

_"Programming is, among other things, a kind of writing. One way to learn writing is to write, but in all other forms of writing, one also reads. We read examples both good and bad to facilitate learning. But how many programmers learn to write programs by reading programs?" - Gerald M. Weinberg_

_"Skill develops when we produce, not consume." - Katrina Owen_

---

### Legacy Software

Do you care about the legacy you are leaving behind?

**Quotes**

_"There are two kinds of software projects: those that fail, and those that turn into legacy horrors." - Peter Weinberger (inventor of AWK)_

_"Legacy software is an unappreciated but serious problem. Legacy code may be the downfall of our civilization." - Chuck Moore (inventor of Forth)_

_"Few programmers of any experience would contradict the assertion that most programs are modified in their lifetime. Why then do we rarely find a program that contains any evidence of having been written with an eye to subsequent modification." - Gerald M. Weinberg_

_"We think awful code is written by awful devs. But in reality, it's written by reasonable devs in awful circumstances." - Sarah Mei_

_"There are many reasons why programs are built the way they are, although we may fail to recognize the multiplicity of reasons because we usually look at code from the outside rather than by reading it. When we do read code, we find that some of it gets written because of machine limitations, some because of language limitations, some because of programmer limitations, some because of historical accidents, and some because of specifications—both essential and inessential." - Gerald M. Weinberg_

---

### Mental Models

You must constantly make sure your mental model of the code you are writing and maintaining is clear. When you can't remember where a piece of logic is or you can't remember how something works, you’re losing your mental model of the code. This is a clear indication that you need to refactor the code. Focus time on structuring code that provides the best mental model possible and during code reviews validate your mental models are still intact.

How much code do you think you can maintain in your head? I believe asking a single developer to maintain a mental model of more than one ream of copy paper (~10k lines of code) is asking a lot. If you do the math, it takes a team of 100 people to work on a code base that hits a million lines of code. That’s 100 people that need to be coordinated, grouped, tracked and in a constant feedback loop of communication.

**Quotes**

_"Let's imagine a project that's going to end up with a million lines of code or more. The probability of those projects being successful in the United States these days is very low - well under 50%. That's debatable." - Tom Love (inventor of Objective C)_

_"100k lines of code fit inside a box of paper." - Tom Love (inventor of Objective C)_

_"One of our many problems with thinking is “cognitive load”: the number of things we can pay attention to at once. The cliche is 7±2, but for many things it is even less. We make progress by making those few things be more powerful." - Alan Kay_

_"The hardest bugs are those where your mental model of the situation is just wrong, so you can't see the problem at all." - Brian Kernighan_

_"Everyone knows that debugging is twice as hard as writing a program in the first place. So if you're as clever as you can be when you write it, how will you ever debug it?" - Brian Kernighan_

_"Debuggers don't remove bugs. They only show them in slow motion." - Unknown_

_"Fixing bugs is just a side effect. Debuggers are for exploration." - @Deech (Twitter)_

---
