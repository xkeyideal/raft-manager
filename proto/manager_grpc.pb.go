// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: manager.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	RaftManager_AddNonvoter_FullMethodName                = "/RaftManager/AddNonvoter"
	RaftManager_AddVoter_FullMethodName                   = "/RaftManager/AddVoter"
	RaftManager_AppliedIndex_FullMethodName               = "/RaftManager/AppliedIndex"
	RaftManager_ApplyLog_FullMethodName                   = "/RaftManager/ApplyLog"
	RaftManager_Barrier_FullMethodName                    = "/RaftManager/Barrier"
	RaftManager_DemoteVoter_FullMethodName                = "/RaftManager/DemoteVoter"
	RaftManager_GetConfiguration_FullMethodName           = "/RaftManager/GetConfiguration"
	RaftManager_LastContact_FullMethodName                = "/RaftManager/LastContact"
	RaftManager_LastIndex_FullMethodName                  = "/RaftManager/LastIndex"
	RaftManager_Leader_FullMethodName                     = "/RaftManager/Leader"
	RaftManager_LeadershipTransfer_FullMethodName         = "/RaftManager/LeadershipTransfer"
	RaftManager_LeadershipTransferToServer_FullMethodName = "/RaftManager/LeadershipTransferToServer"
	RaftManager_RemoveServer_FullMethodName               = "/RaftManager/RemoveServer"
	RaftManager_Shutdown_FullMethodName                   = "/RaftManager/Shutdown"
	RaftManager_Snapshot_FullMethodName                   = "/RaftManager/Snapshot"
	RaftManager_State_FullMethodName                      = "/RaftManager/State"
	RaftManager_Stats_FullMethodName                      = "/RaftManager/Stats"
	RaftManager_VerifyLeader_FullMethodName               = "/RaftManager/VerifyLeader"
	RaftManager_Await_FullMethodName                      = "/RaftManager/Await"
	RaftManager_Forget_FullMethodName                     = "/RaftManager/Forget"
)

// RaftManagerClient is the client API for RaftManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftManagerClient interface {
	AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*Future, error)
	AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*Future, error)
	AppliedIndex(ctx context.Context, in *AppliedIndexRequest, opts ...grpc.CallOption) (*AppliedIndexResponse, error)
	ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*Future, error)
	Barrier(ctx context.Context, in *BarrierRequest, opts ...grpc.CallOption) (*Future, error)
	DemoteVoter(ctx context.Context, in *DemoteVoterRequest, opts ...grpc.CallOption) (*Future, error)
	GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error)
	LastContact(ctx context.Context, in *LastContactRequest, opts ...grpc.CallOption) (*LastContactResponse, error)
	LastIndex(ctx context.Context, in *LastIndexRequest, opts ...grpc.CallOption) (*LastIndexResponse, error)
	Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error)
	LeadershipTransfer(ctx context.Context, in *LeadershipTransferRequest, opts ...grpc.CallOption) (*Future, error)
	LeadershipTransferToServer(ctx context.Context, in *LeadershipTransferToServerRequest, opts ...grpc.CallOption) (*Future, error)
	RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*Future, error)
	Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*Future, error)
	Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*Future, error)
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error)
	VerifyLeader(ctx context.Context, in *VerifyLeaderRequest, opts ...grpc.CallOption) (*Future, error)
	Await(ctx context.Context, in *Future, opts ...grpc.CallOption) (*AwaitResponse, error)
	Forget(ctx context.Context, in *Future, opts ...grpc.CallOption) (*ForgetResponse, error)
}

type raftManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftManagerClient(cc grpc.ClientConnInterface) RaftManagerClient {
	return &raftManagerClient{cc}
}

func (c *raftManagerClient) AddNonvoter(ctx context.Context, in *AddNonvoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_AddNonvoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) AddVoter(ctx context.Context, in *AddVoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_AddVoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) AppliedIndex(ctx context.Context, in *AppliedIndexRequest, opts ...grpc.CallOption) (*AppliedIndexResponse, error) {
	out := new(AppliedIndexResponse)
	err := c.cc.Invoke(ctx, RaftManager_AppliedIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) ApplyLog(ctx context.Context, in *ApplyLogRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_ApplyLog_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Barrier(ctx context.Context, in *BarrierRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_Barrier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) DemoteVoter(ctx context.Context, in *DemoteVoterRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_DemoteVoter_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) GetConfiguration(ctx context.Context, in *GetConfigurationRequest, opts ...grpc.CallOption) (*GetConfigurationResponse, error) {
	out := new(GetConfigurationResponse)
	err := c.cc.Invoke(ctx, RaftManager_GetConfiguration_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) LastContact(ctx context.Context, in *LastContactRequest, opts ...grpc.CallOption) (*LastContactResponse, error) {
	out := new(LastContactResponse)
	err := c.cc.Invoke(ctx, RaftManager_LastContact_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) LastIndex(ctx context.Context, in *LastIndexRequest, opts ...grpc.CallOption) (*LastIndexResponse, error) {
	out := new(LastIndexResponse)
	err := c.cc.Invoke(ctx, RaftManager_LastIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Leader(ctx context.Context, in *LeaderRequest, opts ...grpc.CallOption) (*LeaderResponse, error) {
	out := new(LeaderResponse)
	err := c.cc.Invoke(ctx, RaftManager_Leader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) LeadershipTransfer(ctx context.Context, in *LeadershipTransferRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_LeadershipTransfer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) LeadershipTransferToServer(ctx context.Context, in *LeadershipTransferToServerRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_LeadershipTransferToServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) RemoveServer(ctx context.Context, in *RemoveServerRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_RemoveServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Shutdown(ctx context.Context, in *ShutdownRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_Shutdown_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Snapshot(ctx context.Context, in *SnapshotRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_Snapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, RaftManager_State_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Stats(ctx context.Context, in *StatsRequest, opts ...grpc.CallOption) (*StatsResponse, error) {
	out := new(StatsResponse)
	err := c.cc.Invoke(ctx, RaftManager_Stats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) VerifyLeader(ctx context.Context, in *VerifyLeaderRequest, opts ...grpc.CallOption) (*Future, error) {
	out := new(Future)
	err := c.cc.Invoke(ctx, RaftManager_VerifyLeader_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Await(ctx context.Context, in *Future, opts ...grpc.CallOption) (*AwaitResponse, error) {
	out := new(AwaitResponse)
	err := c.cc.Invoke(ctx, RaftManager_Await_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftManagerClient) Forget(ctx context.Context, in *Future, opts ...grpc.CallOption) (*ForgetResponse, error) {
	out := new(ForgetResponse)
	err := c.cc.Invoke(ctx, RaftManager_Forget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftManagerServer is the server API for RaftManager service.
// All implementations must embed UnimplementedRaftManagerServer
// for forward compatibility
type RaftManagerServer interface {
	AddNonvoter(context.Context, *AddNonvoterRequest) (*Future, error)
	AddVoter(context.Context, *AddVoterRequest) (*Future, error)
	AppliedIndex(context.Context, *AppliedIndexRequest) (*AppliedIndexResponse, error)
	ApplyLog(context.Context, *ApplyLogRequest) (*Future, error)
	Barrier(context.Context, *BarrierRequest) (*Future, error)
	DemoteVoter(context.Context, *DemoteVoterRequest) (*Future, error)
	GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error)
	LastContact(context.Context, *LastContactRequest) (*LastContactResponse, error)
	LastIndex(context.Context, *LastIndexRequest) (*LastIndexResponse, error)
	Leader(context.Context, *LeaderRequest) (*LeaderResponse, error)
	LeadershipTransfer(context.Context, *LeadershipTransferRequest) (*Future, error)
	LeadershipTransferToServer(context.Context, *LeadershipTransferToServerRequest) (*Future, error)
	RemoveServer(context.Context, *RemoveServerRequest) (*Future, error)
	Shutdown(context.Context, *ShutdownRequest) (*Future, error)
	Snapshot(context.Context, *SnapshotRequest) (*Future, error)
	State(context.Context, *StateRequest) (*StateResponse, error)
	Stats(context.Context, *StatsRequest) (*StatsResponse, error)
	VerifyLeader(context.Context, *VerifyLeaderRequest) (*Future, error)
	Await(context.Context, *Future) (*AwaitResponse, error)
	Forget(context.Context, *Future) (*ForgetResponse, error)
	mustEmbedUnimplementedRaftManagerServer()
}

// UnimplementedRaftManagerServer must be embedded to have forward compatible implementations.
type UnimplementedRaftManagerServer struct {
}

func (UnimplementedRaftManagerServer) AddNonvoter(context.Context, *AddNonvoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNonvoter not implemented")
}
func (UnimplementedRaftManagerServer) AddVoter(context.Context, *AddVoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVoter not implemented")
}
func (UnimplementedRaftManagerServer) AppliedIndex(context.Context, *AppliedIndexRequest) (*AppliedIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedIndex not implemented")
}
func (UnimplementedRaftManagerServer) ApplyLog(context.Context, *ApplyLogRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyLog not implemented")
}
func (UnimplementedRaftManagerServer) Barrier(context.Context, *BarrierRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Barrier not implemented")
}
func (UnimplementedRaftManagerServer) DemoteVoter(context.Context, *DemoteVoterRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DemoteVoter not implemented")
}
func (UnimplementedRaftManagerServer) GetConfiguration(context.Context, *GetConfigurationRequest) (*GetConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (UnimplementedRaftManagerServer) LastContact(context.Context, *LastContactRequest) (*LastContactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastContact not implemented")
}
func (UnimplementedRaftManagerServer) LastIndex(context.Context, *LastIndexRequest) (*LastIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastIndex not implemented")
}
func (UnimplementedRaftManagerServer) Leader(context.Context, *LeaderRequest) (*LeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Leader not implemented")
}
func (UnimplementedRaftManagerServer) LeadershipTransfer(context.Context, *LeadershipTransferRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeadershipTransfer not implemented")
}
func (UnimplementedRaftManagerServer) LeadershipTransferToServer(context.Context, *LeadershipTransferToServerRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeadershipTransferToServer not implemented")
}
func (UnimplementedRaftManagerServer) RemoveServer(context.Context, *RemoveServerRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveServer not implemented")
}
func (UnimplementedRaftManagerServer) Shutdown(context.Context, *ShutdownRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Shutdown not implemented")
}
func (UnimplementedRaftManagerServer) Snapshot(context.Context, *SnapshotRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snapshot not implemented")
}
func (UnimplementedRaftManagerServer) State(context.Context, *StateRequest) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (UnimplementedRaftManagerServer) Stats(context.Context, *StatsRequest) (*StatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stats not implemented")
}
func (UnimplementedRaftManagerServer) VerifyLeader(context.Context, *VerifyLeaderRequest) (*Future, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyLeader not implemented")
}
func (UnimplementedRaftManagerServer) Await(context.Context, *Future) (*AwaitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Await not implemented")
}
func (UnimplementedRaftManagerServer) Forget(context.Context, *Future) (*ForgetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Forget not implemented")
}
func (UnimplementedRaftManagerServer) mustEmbedUnimplementedRaftManagerServer() {}

// UnsafeRaftManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftManagerServer will
// result in compilation errors.
type UnsafeRaftManagerServer interface {
	mustEmbedUnimplementedRaftManagerServer()
}

func RegisterRaftManagerServer(s grpc.ServiceRegistrar, srv RaftManagerServer) {
	s.RegisterService(&RaftManager_ServiceDesc, srv)
}

func _RaftManager_AddNonvoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNonvoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).AddNonvoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_AddNonvoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).AddNonvoter(ctx, req.(*AddNonvoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_AddVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).AddVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_AddVoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).AddVoter(ctx, req.(*AddVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_AppliedIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppliedIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).AppliedIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_AppliedIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).AppliedIndex(ctx, req.(*AppliedIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_ApplyLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).ApplyLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_ApplyLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).ApplyLog(ctx, req.(*ApplyLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Barrier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarrierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Barrier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Barrier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Barrier(ctx, req.(*BarrierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_DemoteVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DemoteVoterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).DemoteVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_DemoteVoter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).DemoteVoter(ctx, req.(*DemoteVoterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_GetConfiguration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).GetConfiguration(ctx, req.(*GetConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_LastContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).LastContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_LastContact_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).LastContact(ctx, req.(*LastContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_LastIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).LastIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_LastIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).LastIndex(ctx, req.(*LastIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Leader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Leader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Leader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Leader(ctx, req.(*LeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_LeadershipTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadershipTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).LeadershipTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_LeadershipTransfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).LeadershipTransfer(ctx, req.(*LeadershipTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_LeadershipTransferToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeadershipTransferToServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).LeadershipTransferToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_LeadershipTransferToServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).LeadershipTransferToServer(ctx, req.(*LeadershipTransferToServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_RemoveServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).RemoveServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_RemoveServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).RemoveServer(ctx, req.(*RemoveServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShutdownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Shutdown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Shutdown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Shutdown(ctx, req.(*ShutdownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Snapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Snapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Snapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Snapshot(ctx, req.(*SnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_State_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Stats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Stats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Stats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Stats(ctx, req.(*StatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_VerifyLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyLeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).VerifyLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_VerifyLeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).VerifyLeader(ctx, req.(*VerifyLeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Await_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Future)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Await(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Await_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Await(ctx, req.(*Future))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftManager_Forget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Future)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftManagerServer).Forget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RaftManager_Forget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftManagerServer).Forget(ctx, req.(*Future))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftManager_ServiceDesc is the grpc.ServiceDesc for RaftManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RaftManager",
	HandlerType: (*RaftManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNonvoter",
			Handler:    _RaftManager_AddNonvoter_Handler,
		},
		{
			MethodName: "AddVoter",
			Handler:    _RaftManager_AddVoter_Handler,
		},
		{
			MethodName: "AppliedIndex",
			Handler:    _RaftManager_AppliedIndex_Handler,
		},
		{
			MethodName: "ApplyLog",
			Handler:    _RaftManager_ApplyLog_Handler,
		},
		{
			MethodName: "Barrier",
			Handler:    _RaftManager_Barrier_Handler,
		},
		{
			MethodName: "DemoteVoter",
			Handler:    _RaftManager_DemoteVoter_Handler,
		},
		{
			MethodName: "GetConfiguration",
			Handler:    _RaftManager_GetConfiguration_Handler,
		},
		{
			MethodName: "LastContact",
			Handler:    _RaftManager_LastContact_Handler,
		},
		{
			MethodName: "LastIndex",
			Handler:    _RaftManager_LastIndex_Handler,
		},
		{
			MethodName: "Leader",
			Handler:    _RaftManager_Leader_Handler,
		},
		{
			MethodName: "LeadershipTransfer",
			Handler:    _RaftManager_LeadershipTransfer_Handler,
		},
		{
			MethodName: "LeadershipTransferToServer",
			Handler:    _RaftManager_LeadershipTransferToServer_Handler,
		},
		{
			MethodName: "RemoveServer",
			Handler:    _RaftManager_RemoveServer_Handler,
		},
		{
			MethodName: "Shutdown",
			Handler:    _RaftManager_Shutdown_Handler,
		},
		{
			MethodName: "Snapshot",
			Handler:    _RaftManager_Snapshot_Handler,
		},
		{
			MethodName: "State",
			Handler:    _RaftManager_State_Handler,
		},
		{
			MethodName: "Stats",
			Handler:    _RaftManager_Stats_Handler,
		},
		{
			MethodName: "VerifyLeader",
			Handler:    _RaftManager_VerifyLeader_Handler,
		},
		{
			MethodName: "Await",
			Handler:    _RaftManager_Await_Handler,
		},
		{
			MethodName: "Forget",
			Handler:    _RaftManager_Forget_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "manager.proto",
}
