# A Tour of Go: Exercise Solutions

This repository contains my solutions to the exercises from [A Tour of Go](https://go.dev/tour).

Every solution comes with a video walkthrough, available in this [YouTube course](https://www.youtube.com/playlist?list=PLeVAytXWNTIA).

## Exercises

| #   | Exercise                                                     | Video                                                | Solution                                     |
| --- | ------------------------------------------------------------ | ---------------------------------------------------- | -------------------------------------------- |
| 01  | [Loops and Functions](https://go.dev/tour/flowcontrol/8)     | [Watch](https://www.youtube.com/watch?v=taRekUoZD8I) | [Code](./01-loops-and-functions/main.go)     |
| 02  | [Slices](https://go.dev/tour/moretypes/18)                   | [Watch](https://www.youtube.com/watch?v=vE_Vz7vuaOU) | [Code](./02-slices/main.go)                  |
| 03  | [Maps](https://go.dev/tour/moretypes/23)                     | [Watch](https://www.youtube.com/watch?v=YXYL0pmbjBU) | [Code](./03-maps/main.go)                    |
| 04  | [Fibonacci Closure](https://go.dev/tour/moretypes/26)        | [Watch](https://www.youtube.com/watch?v=doBVd00Qj1g) | [Code](./04-fibonacci-closure/main.go)       |
| 05  | [Stringers](https://go.dev/tour/methods/18)                  | [Watch](https://www.youtube.com/watch?v=NM-qJ2DgqKI) | [Code](./05-stringers/main.go)               |
| 06  | [Errors](https://go.dev/tour/methods/20)                     | [Watch](https://www.youtube.com/watch?v=m1O8LtW-LuQ) | [Code](./06-errors/main.go)                  |
| 07  | [Readers](https://go.dev/tour/methods/22)                    | [Watch](https://www.youtube.com/watch?v=YHtBbU4f-6A) | [Code](./07-readers/main.go)                 |
| 08  | [rot13Reader](https://go.dev/tour/methods/23)                | [Watch](https://www.youtube.com/watch?v=ak6IilRFRQE) | [Code](./08-rot13reader/main.go)             |
| 09  | [Images](https://go.dev/tour/methods/25)                     | [Watch](https://www.youtube.com/watch?v=P3hUrm1-HeI) | [Code](./09-images/main.go)                  |
| 10  | [Equivalent Binary Trees](https://go.dev/tour/concurrency/7) | [Watch](https://www.youtube.com/watch?v=Yl6ix6lhtFs) | [Code](./10-equivalent-binary-trees/main.go) |
| 11  | [Web Crawler](https://go.dev/tour/concurrency/10)            | [Watch](https://www.youtube.com/watch?v=2yUyMtXKWlk) | [Code](./11-web-crawler/main.go)             |

## Running the solutions

Each solution is runnable as a standalone Go program. For example:

```sh
go run ./01-loops-and-functions
```

> [!NOTE]
> `./02-slices` and `./09-images` print base64-encoded image data, which the [Tour's playground](https://go.dev/tour) renders as a picture.
