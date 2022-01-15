# advanced-database-fp
![](https://github.com/tachunwu/advanced-database-fp/blob/master/doc/architecture.png)

## Introduction
This project is my Advanced Database Final Project when I study at NTHU. In this system, I'm using two different databases to handle different workloads. Furthermore, to perform requests properly, I also use the streaming system to decompose the system into distributed.

## Infrasturcture
I choose these software as my infrastructure...
* Container Orchestration - kubernetes: Google level container orchestration also battle proof
* Main OLTP - Cockroachdb: Similar concept as Google Spanner
* Graph Database - Dgraph: GraphQL native database for Geo-Search and FullText-Search
* Streaming - NATS: CNCF's streaming system written in Go which is my most familiar language  

## Services
There are 4 demo services in this repo. If you want to try this concept make sure you install all the infrasturctures.
* cmd/client: Demo client request by using gRPC client
* cmd/stream: Frontend of all storage service, just a simple publish layer.
* cmd/txn:  Handle transaction workload, also use work steal concept.
* cmd/graph: Handle graph workload, also use work steal concept.

## Author
Tachunwu with LSD
