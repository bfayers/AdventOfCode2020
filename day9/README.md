# Day 9
Continuing the recent theme, this is written in Go.

## The solution

```
$ go build

$ time ./day9
Incorrect value: part1Output
Weakness: part2Output

real    0m12.410s
user    0m12.313s
sys     0m0.016s
```

Since this was pretty slow I changed my solution to instead of converting strings to integers on the fly to do it at file load time resulting in this time:

```
$ go build

$ time ./day9
Incorrect value: part1Output
Weakness: part2Output

real    0m0.140s
user    0m0.141s
sys     0m0.031s
```

Much better!