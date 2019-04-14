// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*

The benchdiff command displays performance changes between benchmarks.

benchdiff parses the output of two 'go test' benchmark runs,
correlates the results per benchmark, and displays the deltas.

To measure the performance impact of a change, use 'go test'
to run benchmarks before and after the change:

	go test -run=NONE -bench=. ./... > old.txt
	# make changes
	go test -run=NONE -bench=. ./... > new.txt

Then feed the benchmark results to benchdiff:

	benchdiff old.txt new.txt

benchdiff will summarize and display the performance changes,
in a format like this:

	$ benchdiff old.txt new.txt
	benchmark           old ns/op     new ns/op     delta
	BenchmarkConcat     523           68.6          -86.88%

	benchmark           old allocs     new allocs     delta
	BenchmarkConcat     3              1              -66.67%

	benchmark           old bytes     new bytes     delta
	BenchmarkConcat     80            48            -40.00%

*/
package main // import "github.com/chavacava/benchdiff"
