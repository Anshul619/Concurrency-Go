# GoRoutines
- [A Goroutine](https://go.dev/tour/concurrency/1) is a function or method which executes independently and simultaneously in connection with any other Goroutines present in your program. ([asynchronous function execution](https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1))
- Goroutines, can be very cheap (light weighted thread): they have little overhead beyond the memory for the stack, which is just a few kilobytes.
- All the Goroutines are working under the main Goroutines if the main Goroutine terminated, then all the goroutine present in the program also terminated. Goroutine always works in the background.
- When a new Goroutine executed, the Goroutine call return immediately.
- The control does not wait for Goroutine to complete their execution just like normal function they always move forward to the next line after the Goroutine call and ignores the value returned by the Goroutine.
- Goroutines can communicate using the [channel](Channels/Readme.md) and these [channels](Channels/Readme.md) are specially designed to prevent race conditions when accessing shared memory using Goroutines.