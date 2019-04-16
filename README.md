[![Build Status](https://travis-ci.org/chavacava/benchdiff.svg?branch=master)](https://travis-ci.org/chavacava/benchdiff)

# benchdiff

`benchdiff` displays performance changes between benchmarks

This is a fork and **drop-in replacement** of [`golang/tools/cmd/benchcmp`](https://github.com/golang/tools/tree/master/cmd/benchcmp) 

It adds the following functionalities:

* Can be configured to return error code != 0 if there are positive deltas (performance regressions) between the benchmarks
* Allows to set tolerance of deltas

## Installation

```
go get github.com/chavacava/benchdiff
```

Requirements:

* GO >= 1.11 installed

### Building from sources

1. clone the repo: `git clone https://github.com/chavacava/benchdiff.git`
2. set `GO111MODULE=on`
3. `make build` will generate an executable under `./bin`

## Usage

```
usage: ./benchdiff old.txt new.txt

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

Replacement of `golang/tools/cmd/benchcmp`

```
$ benchdiff -best ./fixtures/strconcat.old ./fixtures/strconcat.new
benchmark                    old ns/op     new ns/op     delta
BenchmarkConcatString-4      148           143           -3.38%
BenchmarkConcatBuffer-4      8.78          8.91          +1.48%
BenchmarkConcatBuilder-4     2.82          2.81          -0.35%

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

Fail on positive deltas (performance regressions)

```
$ benchdiff -errdelta ./fixtures/strconcat.old ./fixtures/strconcat.new
benchmark                   old ns/op     new ns/op     delta
BenchmarkConcatString-4     148           143           -3.38%
BenchmarkConcatBuffer-4     8.78          8.91          +1.48%
benchdiff: +1.48% ns/op delta between benchmarks
$ echo $?
1
```
Set a tolerance of 2% for deltas of ns/op

```
$ benchdiff -errdelta -tnsop 2  ./fixtures/strconcat.old ./fixtures/strconcat.new
benchmark                    old ns/op     new ns/op     delta
BenchmarkConcatString-4      148           143           -3.38%
BenchmarkConcatBuffer-4      8.78          8.91          +1.48%
BenchmarkConcatBuilder-4     2.82          2.81          -0.35%

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
