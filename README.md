# benchcmp
benchcmp displays performance changes between benchmarks

This is a fork of [`golang/tools/cmd/benchcmp`](https://github.com/golang/tools/tree/master/cmd/benchcmp) 
It adds the following functionalities:

* Returns error code != 0 if there are deltas between the benchmarks
* Allows to set tolerance of deltas

## Usage

```
usage: ./benchcomp old.txt new.txt

  -best
        compare best times from old and new
  -changed
        show only benchmarks that have changed
  -mag
        sort benchmarks by magnitude of change
  -tallocop float
        tolerance for deltas of allocs/op
  -tbop float
        tolerance for deltas of bytes/op
  -tmbs float
        tolerance for deltas of Mb/s
  -tnsop float
        tolerance for deltas of ns/op

Each input file should be from:
        go test -run=NONE -bench=. > [old,new].txt

Benchcmp compares old and new for each benchmark.

If -test.benchmem=true is added to the "go test" command
benchcmp will also compare memory allocations.
```

## Examples
```
$ benchcomp ./fixtures/example1.new ./fixtures/example1.old
benchmark                   old ns/op     new ns/op     delta
BenchmarkConcatString-4     143           148           +3.50%
benchcmp: +3.50% ns/op delta between benchmarks
$ echo $?
1
```
Set a tolerance of 5% for deltas of ns/op

```
$ benchcomp -tnsop 5  ./fixtures/example1.new ./fixtures/example1.old
benchmark                    old ns/op     new ns/op     delta
BenchmarkConcatString-4      143           148           +3.50%
BenchmarkConcatBuffer-4      8.91          8.78          -1.46%
BenchmarkConcatBuilder-4     2.81          2.82          +0.36%

benchmark                    old allocs     new allocs     delta
BenchmarkConcatString-4      0              0              +0.00%
BenchmarkConcatBuffer-4      0              0              +0.00%
BenchmarkConcatBuilder-4     0              0              +0.00%

benchmark                    old bytes     new bytes     delta
BenchmarkConcatString-4      530           530           +0.00%
BenchmarkConcatBuffer-4      2             2             +0.00%
BenchmarkConcatBuilder-4     2             2             +0.00%
$ echo $?
0
```
