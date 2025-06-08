= Homework 7 Answers

== Question 1
Benchmarks were run with `go test -bench .` to measure runtime. The original
sequential implementation finished in about 2.5s. Setting `GOMAXPROCS=1`
for the row renderer shortened this slightly to roughly 2.2s. A worker pool
equal to `runtime.NumCPU()` (8 on my machine) lowered the runtime to around
0.9s. Dispatching every pixel as a separate job took about 0.7s but required
many more goroutines and offered little additional speed‑up.

== Question 2
The best results came from using a worker count near the number of CPUs. More
goroutines gave no significant improvement and sometimes slowed the program due
to scheduling overhead.
