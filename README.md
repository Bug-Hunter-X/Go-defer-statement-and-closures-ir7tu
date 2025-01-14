# Go defer statement and closures

This example demonstrates a common issue in Go involving the `defer` statement and closures.  The output might be unexpected for those unfamiliar with how closures work in Go.

The `defer` statement schedules the execution of `fmt.Println(i)` to occur after the function completes. However, because the function closes over the variable `i`, its value is not fixed at the time the defer is called, but rather evaluated at the time the deferred function is run. This leads to the printed value being 20 instead of the initial value of 10.