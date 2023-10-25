
install pprof
```shell
go install github.com/google/pprof@latest
```

benchmark
```shell
go run .
go test .

# run benchmark
go test -bench='.'
# run test and bench
go test -run='^$' -bench='.'
# select benchmarks to run
go test -run='^$' -bench='BenchmarkCalculateSquaresMain'
go test -run='^$' -bench='BenchmarkCalculateSquares'
# select number of times to run a bench
go test -run='^$' -bench='BenchmarkCalculateSquares' -count='2'
# select duration of bench
go test -run='^$' -bench='BenchmarkCalculateSquares' -count='2' -benchtime='3s'
# generate profile files
go test -run='^$' -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'
```

run pprof
```shell
go tool pprof cpu.prof
```

pprof commands
```
top20
granularity=lines
sort=cum
sort=flat
granularity=functions
web
gif
exit
```

```
nodecount=20
focus=Pow
focus=
show=expensive
png
```

```
granularity=lines
top
granularity=functions
peek processSliceParallel
tree
```

```
tree
ignore=Wait
hide=Wait
hide=
granularity=functions
top
```

```
hide=runtime
focus=math
top
web
```

```
http://localhost:6060/debug/pprof/
go tool pprof http://localhost:6060/debug/pprof/profile
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/profile?seconds=10

go tool pprof -http=:8080 http://localhost:6060/debug/pprof/allocs
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/block
go tool pprof -http=:8080 http://localhost:6060/debug/pprof/mutex
```