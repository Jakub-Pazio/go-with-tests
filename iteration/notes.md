# Iteration chapter

New (v1.22) for loop looks great

## Benchmark results

```
Go-with-tests/iteration on  main [!] via 🐹 v1.22.0 took 3s
❯ go test -bench=.
goos: linux
goarch: amd64
pkg: go-with-tests/iteration
cpu: Intel(R) Core(TM) i5-10300H CPU @ 2.50GHz
BenchmarkBigRepeat-8              632067              1847 ns/op
BenchmarkBigSlowRepeat-8           12615             94952 ns/op
BenchmarkSmallRepeat-8          51601208                22.16 ns/op
BenchmarkSmallSlowRepeat-8      41133814                28.78 ns/op
PASS
ok      go-with-tests/iteration 5.734s
```

Conclusion is simple: 
- even for small number of concatenation it is better to use string builder
- when doing lots of string concatenations string builder is no-brainer
- Go benchmark tooling is dead simple, and very powerful
