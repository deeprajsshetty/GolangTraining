## Design Guidelines

You must develop a design philosophy that establishes a set of guidelines. This is more important than developing a set of rules or patterns you apply blindly. Guidelines help to formulate, drive and validate decisions. You can't begin to make the best decisions without understanding the impact of your decisions. Every decision you make, every line of code you write comes with trade-offs.

* Philosophy
    * [Prepare Your Mind](https://github.com/deeprajsshetty/GolangTraining/tree/master/001-Design%20Guidelines#prepare-your-mind)
    * [Legacy Software](https://github.com/ardanlabs/gotraining/tree/master/topics/go#legacy-software)

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