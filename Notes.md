# Notes

## Remarks

1. Property Based Testing

Once we're more familiar with Go's syntax I will introduce a technique called "Property Based Testing",
which would stop annoying developers and help you find bugs.

1. A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

1. These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.

1. Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

1. The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples could be very large data or perhaps things you intend only to have one instance of (like database connection pools).

1. Errors more details: <https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully>

1. Learned more about errors
    1. How to create errors that are constants
    1. Writing error wrappers

## Progression (Breakpoints)

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#writing-tests>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#switch>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration#practice-exercises>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor-1>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection#refactor-5>

1. <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context>
