# Day 6
I decided to write Day 6's solution in Go

## The solution
My solution uses some for loops to read the answers in and combine them into an array of arrays or an array of strings depending on part1/2

`time` when using `go run`

```
$ time go run main.go
part1Output
part2Output

real    0m0.505s
user    0m0.156s
sys     0m0.453s
```

`time` when using compiled version

```
$ go build

$ time ./day6
part1Output
part2Output

real    0m0.007s
user    0m0.000s
sys     0m0.000s```