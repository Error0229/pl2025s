= Homework 7 Answers

Using `GOMAXPROCS=1` the sequential program finished in about 2.2s. With a worker
pool equal to `runtime.NumCPU()` and dispatching rows over a channel the runtime
dropped to roughly 0.9s – a bit over a 2× speed‑up on this machine. A variant that
queued every pixel individually took around 0.7s but used far more goroutines
and did not scale much better.

The best results came from using a worker count near the number of CPUs. More
goroutines gave no significant improvement and sometimes slowed the program due
to scheduling overhead.
