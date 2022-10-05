## Concurrency

Concurrency allows a program to process multiple tasks at the same time (parallel processing where each task is assigned to a separate processor) or what appears to be at the same time where tasks are multiplexed so progress is made on all tasks over time. If the multiplexing is very fast, it appears that the concurrent processes are running at the same time but are run in overlapping periods of time.

In most languages that support concurrent processing, the thread construct is used to support concurrency. There is memory overhead associated with a thread, so the number of threads that can be spawned at the same time is limited.

## Goroutine

In developing the Go language, Google introduced a lightweight process called a **goroutine** that requires less memory overhead than a thread.
Goroutines are functions that run concurrent with other functions. When a regular function is invoked.

**WaitGroup**  
* In 001-SimpleGoroutineExample ***(Code REf - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/001-SimpleGoroutineExample)***, we have not used WaitGroup so the goroutine running concurrently with main completes it work before the regularFunction and before the main goroutine exits. To handle this sceanrio,
Go provides a mechanism for allowing multiple goroutines to all complete their work before main exits while killing off unfinished goroutines. Go introduce the sync package and the WaitGroup.
* 006-Concurrency -> 002-WaitGroup explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/002-WaitGroup)***.

## The Channel

We often want to be able to synchronize the sequence of goroutines and have them communicate with each other. We introduce the powerful construct of the **channel** to accomplish this.
***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/001-SimpleChannelExample)***

**Select Statement**
* 006-Concurrency -> 002-Channel -> 002-ChannelUsingSelect explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/002-ChannelUsingSelect)***.

**Use a quit Channel to Avoid Using WaitGroup**
* 006-Concurrency -> 002-Channel -> 003-QuitChannel explained the details with Example ***(Code REf - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/003-QuitChannel)***.

## Channel Direction
Channel direction can be added to a goroutine signature. An arrow pointing to the chan from the right, as shown in the signatures to pingGenerator and pongGenerator, requires the goroutine to assign to the channel (a generator). An arrow to the left of chan and pointing to the channel variable requires the goroutine to only consume values in the channel.
* 006-Concurrency -> 002-Channel -> 004-ChannelDirection explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/004-ChannelDirection)***.

**Race Condition**
* A pervasive problem using concurrency is race condition. This problem occurs when two or more goroutines modify the same shared data.
* 006-Concurrency -> 002-Channel -> 005-RaceCondition explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/005-RaceCondition)***.

**Mutex**
* We can correct the race-condition problem by using a mutex.
* Program execution is slowed down using the mutex, but the program is protected from the race condition
* 006-Concurrency -> 002-Channel -> 006-Mutex explained the details with Example ***(Code Ref - https://github.com/deeprajsshetty/GolangTraining/tree/master/006-Concurrency/003-Channel/006-Mutex)***.

## Benchmarking Concurrent Applications
The goal in using concurrency is to speed up program execution. There is overhead in deploying goroutines, so sometimes, using concurrency is counterproductive. Because debugging concurrent code is challenging and dealing with deadlocks and race conditions is sometimes tricky, one needs to be careful when crafting concurrent solutions to a problem. Testing a concurrent application and comparing its performance with a nonconcurrent solution is useful.

**Benchmarking Construct and sum 100 million floating-point numbers**
* Using Concurrency Vs Sequencing Process: 006-Concurrency -> 004-BenchmarkingConcurrentApplications -> 001-ConcurrentVsSequentialProcess
* We can speed up above algorithm by summing half the numbers in each of two goroutines: 006-Concurrency -> 004-BenchmarkingConcurrentApplications -> 002-ConcurrencyUsingDivideAndConquer

## Examples

**Playing Chess Using Goroutines**
* 006-Concurrency -> 005-Examples -> 001-ChessUsingGoroutines

**Fibonacci Numbers Using Goroutines**
* This example output a sequence of Fibonacci numbers using a goroutine.
* 006-Concurrency -> 005-Examples -> 002-FibonacciNoUsingGoroutines

