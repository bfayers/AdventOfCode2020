# Day 5
I decided to write Day 5's solution in Go

## The solution
My solution uses some for loops to split the array of plane rows and columns down to the correct place, it then uses some for loops to find the missing seat id for part 2

`time` when using `go run`

```
$ time go run main.go
part1Output
part2Output

real    0m0.532s
user    0m0.109s
sys     0m0.484s
```

`time` when using compiled version

```
$ go build

$ time ./day5
part1Output
part2Output

real    0m0.006s
user    0m0.000s
sys     0m0.000s```