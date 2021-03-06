// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobufexample

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

// SquareCalculatorClient is the client API for SquareCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SquareCalculatorClient interface {
	CalcSquare(ctx context.Context, in *CalcSquareRequest, opts ...grpc.CallOption) (*CalcSquareResponse, error)
}

type squareCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewSquareCalculatorClient(cc grpc.ClientConnInterface) SquareCalculatorClient {
	return &squareCalculatorClient{cc}
}

func (c *squareCalculatorClient) CalcSquare(ctx context.Context, in *CalcSquareRequest, opts ...grpc.CallOption) (*CalcSquareResponse, error) {
	out := new(CalcSquareResponse)
	err := c.cc.Invoke(ctx, "/squarecalculator.SquareCalculator/CalcSquare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareCalculatorServer is the server API for SquareCalculator service.
// All implementations must embed UnimplementedSquareCalculatorServer
// for forward compatibility
type SquareCalculatorServer interface {
	CalcSquare(context.Context, *CalcSquareRequest) (*CalcSquareResponse, error)
	mustEmbedUnimplementedSquareCalculatorServer()
}

// UnimplementedSquareCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedSquareCalculatorServer struct {
}

func (UnimplementedSquareCalculatorServer) CalcSquare(context.Context, *CalcSquareRequest) (*CalcSquareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalcSquare not implemented")
}
func (UnimplementedSquareCalculatorServer) mustEmbedUnimplementedSquareCalculatorServer() {}

// UnsafeSquareCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SquareCalculatorServer will
// result in compilation errors.
type UnsafeSquareCalculatorServer interface {
	mustEmbedUnimplementedSquareCalculatorServer()
}

func RegisterSquareCalculatorServer(s grpc.ServiceRegistrar, srv SquareCalculatorServer) {
	s.RegisterService(&SquareCalculator_ServiceDesc, srv)
}

func _SquareCalculator_CalcSquare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcSquareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareCalculatorServer).CalcSquare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squarecalculator.SquareCalculator/CalcSquare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareCalculatorServer).CalcSquare(ctx, req.(*CalcSquareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SquareCalculator_ServiceDesc is the grpc.ServiceDesc for SquareCalculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SquareCalculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "squarecalculator.SquareCalculator",
	HandlerType: (*SquareCalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CalcSquare",
			Handler:    _SquareCalculator_CalcSquare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
