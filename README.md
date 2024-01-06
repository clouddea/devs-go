# Introduction

This is a `Parallel DEVS` co-simulation engine implementation in Go language, which supports running simulation tasks in single node or distributed enviroment like Kubernetes.

# Environment

for development:
+ OS: windows 11 / Linux x64
+ DevTools: Goland 1.19.4 or later

Other environments may be ok, but have not been tested and verified.

# Quick Start

Clone this repo and set up GO environment appropriately.  
Simulation tasks can be run in a stand-alone or a distributed manner.
You can desing and build your own DEVS models by referring to the examples in source code.
## stand-alone
Change Directory into this project and type this command:
```shell
go run test/testGenr.go
```
## distributed
Change Directory into this project and type this command:
```shell
go run test/distributed/node1.go
go run test/distributed/node2.go
go run test/distributed/node3.go
go run test/distributed/testDistributed.go

```
You will see outputs like this:

node1.go:
```text
served on localhost:8081 ...
generator1 out count0, state is active
generator1: delta int
generator1 out count1, state is hot active
generator1: delta int
generator1 out count2, state is active
generator1: delta int
generator1 out count3, state is hot active
generator1: delta int
```
node2.go:
```text
served on localhost:8082 ...
transmitter1 out{[{transmitter1 out ##NoTarget ##NoTargetPort data 0}]}, state is active
transmitter1 out{[{transmitter1 out ##NoTarget ##NoTargetPort data 1}]}, state is active
transmitter1 out{[{transmitter1 out ##NoTarget ##NoTargetPort data 2}]}, state is active
transmitter1 out{[{transmitter1 out ##NoTarget ##NoTargetPort data 3}]}, state is active
```

node3.go
```text
served on localhost:8083 ...
processor1 out job data 0, state is active
processor1 out job data 1, state is active
processor1 out job data 2, state is active
processor1 out job data 3, state is active
```

testDistributed.go
```text
served on localhost:8080 ...
time advance: 10
time advance: 10
time advance: 14
time advance: 15
time advance: 15
time advance: 19
time advance: 25
time advance: 25
time advance: 29
time advance: 30
time advance: 30
time advance: 34

```