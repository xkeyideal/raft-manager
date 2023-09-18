# raft-manager

RaftAdmin is a gRPC service to invoke methods on https://godoc.org/github.com/hashicorp/raft#Raft. It only works with [Hashicorp's Raft implementation](https://github.com/hashicorp/raft).

## Usage

```go
r, err := raft.NewRaft(...)
s := grpc.NewServer()
raftmanager.Register(s, r)
```

Adding the call to `raftmanager.Register` will register a new gRPC service on your existing server that exposes a bunch of methods so they can be called remotely.

For example, I use this to add servers (voters) after initial bootstrap.

## Invocations

```shell
$ raftmanager
Usage: raftmanager <host:port> <command> <args...>
Commands: add_nonvoter, add_voter, applied_index, apply_log, await, barrier, demote_voter, forget, get_configuration, last_contact, last_index, leader, leadership_transfer, leadership_transfer_to_server, remove_server, shutdown, snapshot, state, stats, verify_leader

$ raftmanager 127.0.0.1:50051 add_voter serverb 127.0.0.1:50052 0
Invoking AddVoter(id: "serverb" address: "127.0.0.1:50052")
Response: operation_token:  "4a86d2efa417af281ac540bfede8fcb735e0b224"
Invoking Await(operation_token: "4a86d2efa417af281ac540bfede8fcb735e0b224")
Response: index:  3
```

AddVoter starts a new raft operation and returns once it is enqueued. It returns an operation_token with which you can call Await. Nearly all errors are detected by Await and returns as AwaitResponse.error.

Last, call Forget to make the server forget the operation token and free up the memory.

## Inspired

[raftadmin](https://github.com/Jille/raftadmin)

The [Jille/raft-grpc-leader-rpc](https://github.com/Jille/raft-grpc-leader-rpc) use grpc healthv1 by setting `HealthCheckResponse_SERVING` to send RPCs to the leader.

```go
func setServingStatus(hs *health.Server, services []string, isLeader bool) {
	v := grpc_health_v1.HealthCheckResponse_NOT_SERVING
	if isLeader {
		v = grpc_health_v1.HealthCheckResponse_SERVING
	}
	for _, srv := range services {
		hs.SetServingStatus(srv, v)
	}
	hs.SetServingStatus("quis.RaftLeader", v)
}
```

## License
raft-example is under the BSD 2-Clause License. See the LICENSE file for details.