## Jonathan Bodner Intro to ~80% of Go

[Tutorial](https://youtu.be/3zcKjKySMKE)

[Go Modules](https://youtu.be/3zcKjKySMKE?t=1909)
[What's missing?](https://youtu.be/3zcKjKySMKE?t=2100)
[The `error` Interface](https://youtu.be/3zcKjKySMKE?t=2236)

- [Sentinal Errors](https://youtu.be/3zcKjKySMKE?t=2479)
- [When to use errors.Is() vs. errors?](https://youtu.be/3zcKjKySMKE?t=2904)
  - errors.Is() - when looking for specific instance and/or values
  - errors.As() - when looking for specific type
- Use 3rd party error wrappers to generate stack traces for errors
  - [Dave Chaney's](https://github.com/pkg/errors) - Use `fmt.Printf(%+v)` for verbose output

[Routines vs. Channels](https://youtu.be/3zcKjKySMKE?t=3603)
[Generics](https://learning-go-book.dev/chapter15_learningGo.pdf)
