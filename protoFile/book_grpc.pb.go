// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: book.proto

package protoFile

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

// BookManagementServiceClient is the client API for BookManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookManagementServiceClient interface {
	//Create Book
	CreateBook(ctx context.Context, in *ReqCreateBook, opts ...grpc.CallOption) (*ResCreatebook, error)
	//Get All Books
	GetAllBooks(ctx context.Context, in *ReqGetAllBooks, opts ...grpc.CallOption) (BookManagementService_GetAllBooksClient, error)
	//Search Book
	SearchBook(ctx context.Context, in *ReqSearchBook, opts ...grpc.CallOption) (BookManagementService_SearchBookClient, error)
	//Update Book
	UpdateBook(ctx context.Context, in *ReqUpdateBook, opts ...grpc.CallOption) (*ResUpdateBook, error)
	//Delete book
	DeleteBook(ctx context.Context, in *ReqDeleteBook, opts ...grpc.CallOption) (*ResDeleteBook, error)
}

type bookManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookManagementServiceClient(cc grpc.ClientConnInterface) BookManagementServiceClient {
	return &bookManagementServiceClient{cc}
}

func (c *bookManagementServiceClient) CreateBook(ctx context.Context, in *ReqCreateBook, opts ...grpc.CallOption) (*ResCreatebook, error) {
	out := new(ResCreatebook)
	err := c.cc.Invoke(ctx, "/protFile.BookManagementService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookManagementServiceClient) GetAllBooks(ctx context.Context, in *ReqGetAllBooks, opts ...grpc.CallOption) (BookManagementService_GetAllBooksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookManagementService_ServiceDesc.Streams[0], "/protFile.BookManagementService/GetAllBooks", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookManagementServiceGetAllBooksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookManagementService_GetAllBooksClient interface {
	Recv() (*ResGetAllBooks, error)
	grpc.ClientStream
}

type bookManagementServiceGetAllBooksClient struct {
	grpc.ClientStream
}

func (x *bookManagementServiceGetAllBooksClient) Recv() (*ResGetAllBooks, error) {
	m := new(ResGetAllBooks)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookManagementServiceClient) SearchBook(ctx context.Context, in *ReqSearchBook, opts ...grpc.CallOption) (BookManagementService_SearchBookClient, error) {
	stream, err := c.cc.NewStream(ctx, &BookManagementService_ServiceDesc.Streams[1], "/protFile.BookManagementService/SearchBook", opts...)
	if err != nil {
		return nil, err
	}
	x := &bookManagementServiceSearchBookClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BookManagementService_SearchBookClient interface {
	Recv() (*ResSearchBook, error)
	grpc.ClientStream
}

type bookManagementServiceSearchBookClient struct {
	grpc.ClientStream
}

func (x *bookManagementServiceSearchBookClient) Recv() (*ResSearchBook, error) {
	m := new(ResSearchBook)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bookManagementServiceClient) UpdateBook(ctx context.Context, in *ReqUpdateBook, opts ...grpc.CallOption) (*ResUpdateBook, error) {
	out := new(ResUpdateBook)
	err := c.cc.Invoke(ctx, "/protFile.BookManagementService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookManagementServiceClient) DeleteBook(ctx context.Context, in *ReqDeleteBook, opts ...grpc.CallOption) (*ResDeleteBook, error) {
	out := new(ResDeleteBook)
	err := c.cc.Invoke(ctx, "/protFile.BookManagementService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookManagementServiceServer is the server API for BookManagementService service.
// All implementations must embed UnimplementedBookManagementServiceServer
// for forward compatibility
type BookManagementServiceServer interface {
	//Create Book
	CreateBook(context.Context, *ReqCreateBook) (*ResCreatebook, error)
	//Get All Books
	GetAllBooks(*ReqGetAllBooks, BookManagementService_GetAllBooksServer) error
	//Search Book
	SearchBook(*ReqSearchBook, BookManagementService_SearchBookServer) error
	//Update Book
	UpdateBook(context.Context, *ReqUpdateBook) (*ResUpdateBook, error)
	//Delete book
	DeleteBook(context.Context, *ReqDeleteBook) (*ResDeleteBook, error)
	mustEmbedUnimplementedBookManagementServiceServer()
}

// UnimplementedBookManagementServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookManagementServiceServer struct {
}

func (UnimplementedBookManagementServiceServer) CreateBook(context.Context, *ReqCreateBook) (*ResCreatebook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookManagementServiceServer) GetAllBooks(*ReqGetAllBooks, BookManagementService_GetAllBooksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllBooks not implemented")
}
func (UnimplementedBookManagementServiceServer) SearchBook(*ReqSearchBook, BookManagementService_SearchBookServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchBook not implemented")
}
func (UnimplementedBookManagementServiceServer) UpdateBook(context.Context, *ReqUpdateBook) (*ResUpdateBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookManagementServiceServer) DeleteBook(context.Context, *ReqDeleteBook) (*ResDeleteBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookManagementServiceServer) mustEmbedUnimplementedBookManagementServiceServer() {}

// UnsafeBookManagementServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookManagementServiceServer will
// result in compilation errors.
type UnsafeBookManagementServiceServer interface {
	mustEmbedUnimplementedBookManagementServiceServer()
}

func RegisterBookManagementServiceServer(s grpc.ServiceRegistrar, srv BookManagementServiceServer) {
	s.RegisterService(&BookManagementService_ServiceDesc, srv)
}

func _BookManagementService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqCreateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protFile.BookManagementService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServiceServer).CreateBook(ctx, req.(*ReqCreateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookManagementService_GetAllBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqGetAllBooks)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookManagementServiceServer).GetAllBooks(m, &bookManagementServiceGetAllBooksServer{stream})
}

type BookManagementService_GetAllBooksServer interface {
	Send(*ResGetAllBooks) error
	grpc.ServerStream
}

type bookManagementServiceGetAllBooksServer struct {
	grpc.ServerStream
}

func (x *bookManagementServiceGetAllBooksServer) Send(m *ResGetAllBooks) error {
	return x.ServerStream.SendMsg(m)
}

func _BookManagementService_SearchBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReqSearchBook)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookManagementServiceServer).SearchBook(m, &bookManagementServiceSearchBookServer{stream})
}

type BookManagementService_SearchBookServer interface {
	Send(*ResSearchBook) error
	grpc.ServerStream
}

type bookManagementServiceSearchBookServer struct {
	grpc.ServerStream
}

func (x *bookManagementServiceSearchBookServer) Send(m *ResSearchBook) error {
	return x.ServerStream.SendMsg(m)
}

func _BookManagementService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqUpdateBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protFile.BookManagementService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServiceServer).UpdateBook(ctx, req.(*ReqUpdateBook))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookManagementService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDeleteBook)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookManagementServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protFile.BookManagementService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookManagementServiceServer).DeleteBook(ctx, req.(*ReqDeleteBook))
	}
	return interceptor(ctx, in, info, handler)
}

// BookManagementService_ServiceDesc is the grpc.ServiceDesc for BookManagementService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookManagementService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protFile.BookManagementService",
	HandlerType: (*BookManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBook",
			Handler:    _BookManagementService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookManagementService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookManagementService_DeleteBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllBooks",
			Handler:       _BookManagementService_GetAllBooks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchBook",
			Handler:       _BookManagementService_SearchBook_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "book.proto",
}
