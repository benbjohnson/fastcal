fastcal
=======

fast epoch -> calendar date conversion in Go

Running the benchmark
====================

go test -test.bench=.*

CPU Profiling
=============

build the test binary:
go test -c

run benchmark, generating profile:
./fastcal.test -test.bench=.* -test.cpuprofile=prof.out 

open pprof
go tool pprof fastcal.test prof.out
top10


