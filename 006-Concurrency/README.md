## Concurrency

Concurrency allows a program to process multiple tasks at the same time (parallel processing where each task is assigned to a separate processor) or what appears to be at the same time where tasks are multiplexed so progress is made on all tasks over time. If the multiplexing is very fast, it appears that the concurrent processes are running at the same time but are run in overlapping periods of time.

In most languages that support concurrent processing, the thread construct is used to support concurrency. There is memory overhead associated with a thread, so the number of threads that can be spawned at the same time is limited.

## Goroutine

In developing the Go language, Google introduced a lightweight process called a **goroutine** that requires less memory overhead than a thread.
Goroutines are functions that run concurrent with other functions. When a regular function is invoked.

**WaitGroup**  
* In 001-SimpleGoroutineExample, we have not used WaitGroup so the goroutine running concurrently with main completes it work before the regularFunction and before the main goroutine exits. To handle this sceanrio,
Go provides a mechanism for allowing multiple goroutines to all complete their work before main exits while killing off unfinished goroutines. Go introduce the sync package and the WaitGroup.
* 006-Concurrency -> 002-WaitGroup explained the details with Example.

