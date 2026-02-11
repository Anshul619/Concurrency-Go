# Context in Go
- [Context](https://pkg.go.dev/context) provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutines.

# Benefits

| Benefit                                                                                                                                   |
|-------------------------------------------------------------------------------------------------------------------------------------------|
| Propagate deadlines or timeouts across API boundaries and between goroutines.                                                             |
| Signal cancellation (e.g. when a client disconnects, or when a timeout is hit).                                                           |
| Carry request-scoped values (like auth tokens, request IDs) across API calls without polluting function signatures with extra parameters. |

# Real-World Scenarios

| Title                    | Description                                                                                                                                                                                                                                                                                                                |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Context in Microservices | In a microservices architecture, each service often relies on various external dependencies and communicates with other services. <br/>- Context can be used to propagate important information, such as **authentication tokens**, **request metadata**, or **tracing identifiers**, throughout the service interactions. |
| Context in Test Suites   | When writing test suites, context can be utilized to manage test timeouts, control test-specific configurations, and enable graceful termination of tests.                                                                                                                                                                 |
| Context in Web Servers   | Web servers handle multiple concurrent requests, and context helps manage the lifecycle of each request.                                                                                                                                                                                                                   |

# Best Practices

| Title                                  | Description                                                                                                                                                                                                                    |
|----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Pass Context Explicitly                | Always pass the context as an explicit argument to functions or goroutines instead of using global variables. <br/>- This makes it easier to manage the context’s lifecycle and prevents potential data races.                 |
| Prefer Cancel Over Timeout             | Use `context.WithCancel()` for cancellation when possible, as it allows you to explicitly trigger cancellation when needed. <br/>- `context.WithTimeout()` is more suitable when you need an automatic cancellation mechanism. |
| Use `context.TODO()`                   | If you are unsure which context to use in a particular scenario, consider using `context.TODO()`. However, make sure to replace it with the appropriate context later.                                                         |
| Keep Context Size Small                | Avoid storing large or unnecessary data in the context. Only include the data required for the specific operation.                                                                                                             |
| Be Mindful of Goroutine Leaks          | Always ensure that goroutines associated with a context are properly closed or terminated to avoid goroutine leaks.                                                                                                            |
| :x: Avoid using `context.Background()` | Instead of using `context.Background()` directly, create a specific context using `context.WithCancel()` or `context.WithTimeout()` to manage its lifecycle and avoid resource leaks.                                          |
| :x: Avoid Chaining Contexts            | Chaining contexts can lead to confusion and make it challenging to manage the context hierarchy. <br/>- Instead, propagate a single context throughout the application.                                                        |

# Constructs

| Title                                                                                    | Quick Note                        | Description                                                            |
|------------------------------------------------------------------------------------------|-----------------------------------|------------------------------------------------------------------------|
| ctx, cancel := context.Background()                                                      | New Context                       | Create new context                                                     |
| ctx, cancel := context.WithCancel(context.Background())                                  | Manual cancellation               | Create child context which would be cancelled using `cancel()` call    |
| ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)                  | Auto-cancel after duration        | Create child context with "duration based" timeout                     |
| ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second)) | Auto-cancel at a specific time    | Create child context with "time based" timeout                         |
| ctx, cancel := context.WithValue(context.Background(), "UserID", 123)                    | Attaches key/value to the context | Create child context with custom values                                |
| <-ctx.Done()                                                                             |                                   | Blocking channel until context is done (completed/cancelled/timed-out) |

# Context interface methods

````
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
````

# References
- [The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)
- [Debugging High Latency Due to Context Leaks](https://engineering.grab.com/debugging-high-latency-market-store)