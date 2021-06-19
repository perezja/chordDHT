// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dht

import (
	context "context"
	common "github.com/gedilabs/services/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DHTClient is the client API for DHT service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DHTClient interface {
	// Ask node n to find id's successor.
	// This protobuf input message should only have an id.
	FindSuccessor(ctx context.Context, in *Key, opts ...grpc.CallOption) (*common.Node, error)
	// return closest finger preceding id.
	ClosestPrecedingFinger(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Chord, error)
	// Getters/Setters for remote variable lookup and assignment
	Predecessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.Node, error)
	Successor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.Node, error)
	UpdatePredecessor(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Status, error)
	// Update the given index 'i' in a nodes finger table with `Chord.Node` if they
	// precede the current finger node occupying that position.
	UpdateFingerTable(ctx context.Context, in *Finger, opts ...grpc.CallOption) (*common.Status, error)
	// Check if node is predecessor and update accordingly
	Notify(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Status, error)
}

type dHTClient struct {
	cc grpc.ClientConnInterface
}

func NewDHTClient(cc grpc.ClientConnInterface) DHTClient {
	return &dHTClient{cc}
}

func (c *dHTClient) FindSuccessor(ctx context.Context, in *Key, opts ...grpc.CallOption) (*common.Node, error) {
	out := new(common.Node)
	err := c.cc.Invoke(ctx, "/dht.DHT/FindSuccessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) ClosestPrecedingFinger(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Chord, error) {
	out := new(Chord)
	err := c.cc.Invoke(ctx, "/dht.DHT/ClosestPrecedingFinger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) Predecessor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.Node, error) {
	out := new(common.Node)
	err := c.cc.Invoke(ctx, "/dht.DHT/Predecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) Successor(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*common.Node, error) {
	out := new(common.Node)
	err := c.cc.Invoke(ctx, "/dht.DHT/Successor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) UpdatePredecessor(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Status, error) {
	out := new(common.Status)
	err := c.cc.Invoke(ctx, "/dht.DHT/UpdatePredecessor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) UpdateFingerTable(ctx context.Context, in *Finger, opts ...grpc.CallOption) (*common.Status, error) {
	out := new(common.Status)
	err := c.cc.Invoke(ctx, "/dht.DHT/UpdateFingerTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dHTClient) Notify(ctx context.Context, in *common.Node, opts ...grpc.CallOption) (*common.Status, error) {
	out := new(common.Status)
	err := c.cc.Invoke(ctx, "/dht.DHT/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DHTServer is the server API for DHT service.
// All implementations must embed UnimplementedDHTServer
// for forward compatibility
type DHTServer interface {
	// Ask node n to find id's successor.
	// This protobuf input message should only have an id.
	FindSuccessor(context.Context, *Key) (*common.Node, error)
	// return closest finger preceding id.
	ClosestPrecedingFinger(context.Context, *Key) (*Chord, error)
	// Getters/Setters for remote variable lookup and assignment
	Predecessor(context.Context, *Empty) (*common.Node, error)
	Successor(context.Context, *Empty) (*common.Node, error)
	UpdatePredecessor(context.Context, *common.Node) (*common.Status, error)
	// Update the given index 'i' in a nodes finger table with `Chord.Node` if they
	// precede the current finger node occupying that position.
	UpdateFingerTable(context.Context, *Finger) (*common.Status, error)
	// Check if node is predecessor and update accordingly
	Notify(context.Context, *common.Node) (*common.Status, error)
	mustEmbedUnimplementedDHTServer()
}

// UnimplementedDHTServer must be embedded to have forward compatible implementations.
type UnimplementedDHTServer struct {
}

func (UnimplementedDHTServer) FindSuccessor(context.Context, *Key) (*common.Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindSuccessor not implemented")
}
func (UnimplementedDHTServer) ClosestPrecedingFinger(context.Context, *Key) (*Chord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClosestPrecedingFinger not implemented")
}
func (UnimplementedDHTServer) Predecessor(context.Context, *Empty) (*common.Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predecessor not implemented")
}
func (UnimplementedDHTServer) Successor(context.Context, *Empty) (*common.Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Successor not implemented")
}
func (UnimplementedDHTServer) UpdatePredecessor(context.Context, *common.Node) (*common.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePredecessor not implemented")
}
func (UnimplementedDHTServer) UpdateFingerTable(context.Context, *Finger) (*common.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFingerTable not implemented")
}
func (UnimplementedDHTServer) Notify(context.Context, *common.Node) (*common.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedDHTServer) mustEmbedUnimplementedDHTServer() {}

// UnsafeDHTServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DHTServer will
// result in compilation errors.
type UnsafeDHTServer interface {
	mustEmbedUnimplementedDHTServer()
}

func RegisterDHTServer(s grpc.ServiceRegistrar, srv DHTServer) {
	s.RegisterService(&_DHT_serviceDesc, srv)
}

func _DHT_FindSuccessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).FindSuccessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/FindSuccessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).FindSuccessor(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_ClosestPrecedingFinger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).ClosestPrecedingFinger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/ClosestPrecedingFinger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).ClosestPrecedingFinger(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_Predecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Predecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/Predecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Predecessor(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_Successor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Successor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/Successor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Successor(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_UpdatePredecessor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).UpdatePredecessor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/UpdatePredecessor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).UpdatePredecessor(ctx, req.(*common.Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_UpdateFingerTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Finger)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).UpdateFingerTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/UpdateFingerTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).UpdateFingerTable(ctx, req.(*Finger))
	}
	return interceptor(ctx, in, info, handler)
}

func _DHT_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DHTServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dht.DHT/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DHTServer).Notify(ctx, req.(*common.Node))
	}
	return interceptor(ctx, in, info, handler)
}

var _DHT_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dht.DHT",
	HandlerType: (*DHTServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindSuccessor",
			Handler:    _DHT_FindSuccessor_Handler,
		},
		{
			MethodName: "ClosestPrecedingFinger",
			Handler:    _DHT_ClosestPrecedingFinger_Handler,
		},
		{
			MethodName: "Predecessor",
			Handler:    _DHT_Predecessor_Handler,
		},
		{
			MethodName: "Successor",
			Handler:    _DHT_Successor_Handler,
		},
		{
			MethodName: "UpdatePredecessor",
			Handler:    _DHT_UpdatePredecessor_Handler,
		},
		{
			MethodName: "UpdateFingerTable",
			Handler:    _DHT_UpdateFingerTable_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _DHT_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dht/dht.proto",
}
