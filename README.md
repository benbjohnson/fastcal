fastcal
=======

fast epoch -> calendar date conversion in Go

Running the benchmark
====================

go test -test.bench=.*

CPU Profiling
=============

Build the test binary:

```sh
$ go test -c
```

Run benchmark, generating profile:

```sh
$ ./fastcal.test -test.bench=.* -test.cpuprofile=prof.out 
```

Open pprof:

```sh
$ go tool pprof fastcal.test prof.out
> top10
```


