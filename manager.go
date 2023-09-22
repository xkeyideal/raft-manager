package raftmanager

import (
	"context"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	pb "github.com/xkeyideal/raft-manager/proto"
	"google.golang.org/grpc"
)

type manager struct {
	pb.UnimplementedRaftManagerServer
	r *raft.Raft
}

func Get(r *raft.Raft) pb.RaftManagerServer {
	return &manager{
		r: r,
	}
}

func Register(s *grpc.Server, r *raft.Raft) {
	pb.RegisterRaftManagerServer(s, Get(r))
}

func timeout(ctx context.Context) time.Duration {
	if dl, ok := ctx.Deadline(); ok {
		return dl.Sub(time.Now())
	}
	return 0
}

var (
	mtx        sync.Mutex
	operations = map[string]*future{}
)

type future struct {
	f   raft.Future
	mtx sync.Mutex
}

func toFuture(f raft.Future) (*pb.Future, error) {
	token := fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d", rand.Uint64()))))
	mtx.Lock()
	operations[token] = &future{f: f}
	mtx.Unlock()
	return &pb.Future{
		OperationToken: token,
	}, nil
}

func (m *manager) Await(ctx context.Context, req *pb.Future) (*pb.AwaitResponse, error) {
	mtx.Lock()
	f, ok := operations[req.GetOperationToken()]
	mtx.Unlock()
	if !ok {
		return nil, fmt.Errorf("token %q unknown", req.GetOperationToken())
	}
	f.mtx.Lock()
	err := f.f.Error()
	f.mtx.Unlock()
	if err != nil {
		return &pb.AwaitResponse{
			Error: err.Error(),
		}, nil
	}
	r := &pb.AwaitResponse{}
	if ifx, ok := f.f.(raft.IndexFuture); ok {
		r.Index = ifx.Index()
	}
	return r, nil
}

func (m *manager) Forget(ctx context.Context, req *pb.Future) (*pb.ForgetResponse, error) {
	mtx.Lock()
	delete(operations, req.GetOperationToken())
	mtx.Unlock()
	return &pb.ForgetResponse{}, nil
}

// AddNonvoter will add the given server to the cluster but won't assign it a
// vote. The server will receive log entries, but it won't participate in
// elections or log entry commitment. If the server is already in the cluster,
// this updates the server's address. This must be run on the leader or it will
// fail. For prevIndex and timeout, see AddVoter.
func (m *manager) AddNonvoter(ctx context.Context, req *pb.AddNonvoterRequest) (*pb.Future, error) {
	return toFuture(m.r.AddNonvoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), timeout(ctx)))
}

// AddVoter will add the given server to the cluster as a staging server. If the
// server is already in the cluster as a voter, this updates the server's address.
// This must be run on the leader or it will fail. The leader will promote the
// staging server to a voter once that server is ready. If nonzero, prevIndex is
// the index of the only configuration upon which this change may be applied; if
// another configuration entry has been added in the meantime, this request will
// fail. If nonzero, timeout is how long this server should wait before the
// configuration change log entry is appended.
func (m *manager) AddVoter(ctx context.Context, req *pb.AddVoterRequest) (*pb.Future, error) {
	return toFuture(m.r.AddVoter(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress()), req.GetPreviousIndex(), timeout(ctx)))
}

// AppliedIndex returns the last index applied to the FSM. This is generally
// lagging behind the last index, especially for indexes that are persisted but
// have not yet been considered committed by the leader. NOTE - this reflects
// the last index that was sent to the application's FSM over the apply channel
// but DOES NOT mean that the application's FSM has yet consumed it and applied
// it to its internal state. Thus, the application's state may lag behind this
// index.
func (m *manager) AppliedIndex(ctx context.Context, req *pb.AppliedIndexRequest) (*pb.AppliedIndexResponse, error) {
	return &pb.AppliedIndexResponse{
		Index: m.r.AppliedIndex(),
	}, nil
}

// ApplyLog performs Apply but takes in a Log directly. The only values
// currently taken from the submitted Log are Data and Extensions.
func (m *manager) ApplyLog(ctx context.Context, req *pb.ApplyLogRequest) (*pb.Future, error) {
	return toFuture(m.r.ApplyLog(raft.Log{Data: req.GetData(), Extensions: req.GetExtensions()}, timeout(ctx)))
}

// Barrier is used to issue a command that blocks until all preceding
// operations have been applied to the FSM. It can be used to ensure the
// FSM reflects all queued writes. An optional timeout can be provided to
// limit the amount of time we wait for the command to be started. This
// must be run on the leader, or it will fail.
func (m *manager) Barrier(ctx context.Context, req *pb.BarrierRequest) (*pb.Future, error) {
	return toFuture(m.r.Barrier(timeout(ctx)))
}

// DemoteVoter will take away a server's vote, if it has one. If present, the
// server will continue to receive log entries, but it won't participate in
// elections or log entry commitment. If the server is not in the cluster, this
// does nothing. This must be run on the leader or it will fail. For prevIndex
// and timeout, see AddVoter.
func (m *manager) DemoteVoter(ctx context.Context, req *pb.DemoteVoterRequest) (*pb.Future, error) {
	return toFuture(m.r.DemoteVoter(raft.ServerID(req.GetId()), req.GetPreviousIndex(), timeout(ctx)))
}

// GetConfiguration returns the latest configuration. This may not yet be
// committed. The main loop can access this directly.
func (m *manager) GetConfiguration(ctx context.Context, req *pb.GetConfigurationRequest) (*pb.GetConfigurationResponse, error) {
	f := m.r.GetConfiguration()
	if err := f.Error(); err != nil {
		return nil, err
	}
	resp := &pb.GetConfigurationResponse{}
	for _, s := range f.Configuration().Servers {
		cs := &pb.GetConfigurationResponse_Server{
			Id:      string(s.ID),
			Address: string(s.Address),
		}
		switch s.Suffrage {
		case raft.Voter, raft.Nonvoter, raft.Staging:
			cs.Suffrage = s.Suffrage.String()
		default:
			return nil, fmt.Errorf("unknown server suffrage %v for server %q", s.Suffrage, s.ID)
		}
		resp.Servers = append(resp.Servers, cs)
	}
	return resp, nil
}

// LastContact returns the time of last contact by a leader.
// This only makes sense if we are currently a follower.
func (m *manager) LastContact(ctx context.Context, req *pb.LastContactRequest) (*pb.LastContactResponse, error) {
	t := m.r.LastContact()
	return &pb.LastContactResponse{
		UnixNano:   t.UnixNano(),
		TimeFormat: t.String(),
	}, nil
}

// LastIndex returns the last index in stable storage,
// either from the last log or from the last snapshot.
func (m *manager) LastIndex(ctx context.Context, req *pb.LastIndexRequest) (*pb.LastIndexResponse, error) {
	return &pb.LastIndexResponse{
		Index: m.r.LastIndex(),
	}, nil
}

// LeaderWithID is used to return the current leader address and ID of the cluster.
// It may return empty strings if there is no current leader
// or the leader is unknown.
func (m *manager) Leader(ctx context.Context, req *pb.LeaderRequest) (*pb.LeaderResponse, error) {
	addr, id := m.r.LeaderWithID()
	return &pb.LeaderResponse{
		Address:  string(addr),
		ServerId: string(id),
	}, nil
}

// LeadershipTransfer will transfer leadership to a server in the cluster.
// This can only be called from the leader, or it will fail. The leader will
// stop accepting client requests, make sure the target server is up to date
// and starts the transfer with a TimeoutNow message. This message has the same
// effect as if the election timeout on the target server fires. Since
// it is unlikely that another server is starting an election, it is very
// likely that the target server is able to win the election.  Note that raft
// protocol version 3 is not sufficient to use LeadershipTransfer. A recent
// version of that library has to be used that includes this feature.  Using
// transfer leadership is safe however in a cluster where not every node has
// the latest version. If a follower cannot be promoted, it will fail
// gracefully.
func (m *manager) LeadershipTransfer(ctx context.Context, req *pb.LeadershipTransferRequest) (*pb.Future, error) {
	return toFuture(m.r.LeadershipTransfer())
}

// LeadershipTransferToServer does the same as LeadershipTransfer but takes a
// server in the arguments in case a leadership should be transitioned to a
// specific server in the cluster.  Note that raft protocol version 3 is not
// sufficient to use LeadershipTransfer. A recent version of that library has
// to be used that includes this feature. Using transfer leadership is safe
// however in a cluster where not every node has the latest version. If a
// follower cannot be promoted, it will fail gracefully.
func (m *manager) LeadershipTransferToServer(ctx context.Context, req *pb.LeadershipTransferToServerRequest) (*pb.Future, error) {
	return toFuture(m.r.LeadershipTransferToServer(raft.ServerID(req.GetId()), raft.ServerAddress(req.GetAddress())))
}

// RemoveServer will remove the given server from the cluster. If the current
// leader is being removed, it will cause a new election to occur. This must be
// run on the leader or it will fail. For prevIndex and timeout, see AddVoter.
func (m *manager) RemoveServer(ctx context.Context, req *pb.RemoveServerRequest) (*pb.Future, error) {
	return toFuture(m.r.RemoveServer(raft.ServerID(req.GetId()), req.GetPreviousIndex(), timeout(ctx)))
}

// Shutdown is used to stop the Raft background routines.
// This is not a graceful operation. Provides a future that
// can be used to block until all background routines have exited.
func (m *manager) Shutdown(ctx context.Context, req *pb.ShutdownRequest) (*pb.Future, error) {
	return toFuture(m.r.Shutdown())
}

// Snapshot is used to manually force Raft to take a snapshot. Returns a future
// that can be used to block until complete, and that contains a function that
// can be used to open the snapshot.
func (m *manager) Snapshot(ctx context.Context, req *pb.SnapshotRequest) (*pb.Future, error) {
	return toFuture(m.r.Snapshot())
}

// State returns the state of this raft peer.
func (m *manager) State(ctx context.Context, req *pb.StateRequest) (*pb.StateResponse, error) {
	switch s := m.r.State(); s {
	case raft.Follower, raft.Candidate, raft.Leader, raft.Shutdown:
		return &pb.StateResponse{State: s.String()}, nil
	default:
		return nil, fmt.Errorf("unknown raft state %v", s)
	}
}

func (m *manager) Stats(ctx context.Context, req *pb.StatsRequest) (*pb.StatsResponse, error) {
	ret := &pb.StatsResponse{}
	ret.Stats = map[string]string{}
	for k, v := range m.r.Stats() {
		ret.Stats[k] = v
	}
	return ret, nil
}

// VerifyLeader is used to ensure this peer is still the leader. It may be used
// to prevent returning stale data from the FSM after the peer has lost
// leadership.
func (m *manager) VerifyLeader(ctx context.Context, req *pb.VerifyLeaderRequest) (*pb.Future, error) {
	return toFuture(m.r.VerifyLeader())
}
