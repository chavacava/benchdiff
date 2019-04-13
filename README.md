# benchcmp
benchcmp displays performance changes between benchmarks

This is a fork of [`golang/tools/cmd/benchcmp`](https://github.com/golang/tools/tree/master/cmd/benchcmp) 
It adds the following functionalities:

* Returns error code != 0 if there are deltas between the benchmarks
* Allows to set tolerance of deltas
