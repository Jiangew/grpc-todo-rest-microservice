// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo-service.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	todo-service.proto

It has these top-level messages:
	ToDo
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	ReadAllRequest
	ReadAllResponse
*/
package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Task we have to do
type ToDo struct {
	// Unique integer identifier of the todo task
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Title of the task
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// Detail description of the todo task
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Date and time to remind the todo task
	Reminder *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=reminder" json:"reminder,omitempty"`
}

func (m *ToDo) Reset()                    { *m = ToDo{} }
func (m *ToDo) String() string            { return proto.CompactTextString(m) }
func (*ToDo) ProtoMessage()               {}
func (*ToDo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ToDo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ToDo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ToDo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ToDo) GetReminder() *google_protobuf.Timestamp {
	if m != nil {
		return m.Reminder
	}
	return nil
}

// Request data to create new todo task
type CreateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Task entity to add
	ToDo *ToDo `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains data of created todo task
type CreateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// ID of created task
	Id int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Request data to read todo task
type ReadRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Unique integer identifier of the todo task
	Id int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains todo task data specified in by ID request
type ReadResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Task entity read by ID
	ToDo *ToDo `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Request data to update todo task
type UpdateRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Task entity to update
	ToDo *ToDo `protobuf:"bytes,2,opt,name=toDo" json:"toDo,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetToDo() *ToDo {
	if m != nil {
		return m.ToDo
	}
	return nil
}

// Contains status of update operation
type UpdateResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Contains number of entities have beed updated
	// Equals 1 in case of succesfull update
	Updated int64 `protobuf:"varint,2,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

// Request data to delete todo task
type DeleteRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Unique integer identifier of the todo task to delete
	Id int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains status of delete operation
type DeleteResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// Contains number of entities have beed deleted
	// Equals 1 in case of succesfull delete
	Deleted int64 `protobuf:"varint,2,opt,name=deleted" json:"deleted,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

// Request data to read all todo task
type ReadAllRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
}

func (m *ReadAllRequest) Reset()                    { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()               {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

// Contains list of all todo tasks
type ReadAllResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api" json:"api,omitempty"`
	// List of all todo tasks
	ToDos []*ToDo `protobuf:"bytes,2,rep,name=toDos" json:"toDos,omitempty"`
}

func (m *ReadAllResponse) Reset()                    { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()               {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetToDos() []*ToDo {
	if m != nil {
		return m.ToDos
	}
	return nil
}

func init() {
	proto.RegisterType((*ToDo)(nil), "v1.ToDo")
	proto.RegisterType((*CreateRequest)(nil), "v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "v1.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "v1.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "v1.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "v1.DeleteResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "v1.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "v1.ReadAllResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ToDoService service

type ToDoServiceClient interface {
	// Read all todo tasks
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	// Create new todo task
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read todo task
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update todo task
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete todo task
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type toDoServiceClient struct {
	cc *grpc.ClientConn
}

func NewToDoServiceClient(cc *grpc.ClientConn) ToDoServiceClient {
	return &toDoServiceClient{cc}
}

func (c *toDoServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/ReadAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *toDoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/v1.ToDoService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ToDoService service

type ToDoServiceServer interface {
	// Read all todo tasks
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	// Create new todo task
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read todo task
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update todo task
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete todo task
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterToDoServiceServer(s *grpc.Server, srv ToDoServiceServer) {
	s.RegisterService(&_ToDoService_serviceDesc, srv)
}

func _ToDoService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ToDoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToDoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.ToDoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToDoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToDoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.ToDoService",
	HandlerType: (*ToDoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadAll",
			Handler:    _ToDoService_ReadAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ToDoService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ToDoService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ToDoService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ToDoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo-service.proto",
}

func init() { proto.RegisterFile("todo-service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x53, 0xd3, 0x40,
	0x14, 0x9f, 0xa4, 0xa5, 0xc0, 0x2b, 0x2d, 0xb8, 0xe0, 0xd8, 0xa9, 0x8c, 0x66, 0x72, 0x62, 0x3a,
	0x36, 0xa1, 0x91, 0xe1, 0x50, 0x19, 0x01, 0xe9, 0x38, 0x9e, 0x23, 0x5e, 0xbc, 0x85, 0xe4, 0x99,
	0x2e, 0x93, 0x66, 0x63, 0x76, 0x53, 0x74, 0x1c, 0x3c, 0x78, 0xf0, 0xec, 0xe8, 0xcd, 0xaf, 0xe5,
	0x57, 0xe0, 0xe0, 0xc7, 0x70, 0xb2, 0x9b, 0x14, 0x02, 0xb4, 0x17, 0x4f, 0xed, 0xbe, 0xfd, 0xfd,
	0x7b, 0xbb, 0x2f, 0x0b, 0x44, 0xb0, 0x80, 0xf5, 0x39, 0xa6, 0x53, 0xea, 0xa3, 0x95, 0xa4, 0x4c,
	0x30, 0xa2, 0x4f, 0x07, 0xdd, 0xa7, 0x21, 0x63, 0x61, 0x84, 0xb6, 0xac, 0x9c, 0x65, 0x1f, 0x6c,
	0x41, 0x27, 0xc8, 0x85, 0x37, 0x49, 0x14, 0xa8, 0xbb, 0x5d, 0x00, 0xbc, 0x84, 0xda, 0x5e, 0x1c,
	0x33, 0xe1, 0x09, 0xca, 0x62, 0x5e, 0xec, 0x3e, 0x93, 0x3f, 0x7e, 0x3f, 0xc4, 0xb8, 0xcf, 0x2f,
	0xbc, 0x30, 0xc4, 0xd4, 0x66, 0x89, 0x44, 0xdc, 0x45, 0x9b, 0xdf, 0x35, 0xa8, 0x9f, 0xb2, 0x11,
	0x23, 0x6d, 0xd0, 0x69, 0xd0, 0xd1, 0x0c, 0x6d, 0xa7, 0xe6, 0xea, 0x34, 0x20, 0x5b, 0xb0, 0x24,
	0xa8, 0x88, 0xb0, 0xa3, 0x1b, 0xda, 0xce, 0xaa, 0xab, 0x16, 0xc4, 0x80, 0x66, 0x80, 0xdc, 0x4f,
	0xa9, 0x14, 0xec, 0xd4, 0xe4, 0xde, 0xcd, 0x12, 0xd9, 0x87, 0x95, 0x14, 0x27, 0x34, 0x0e, 0x30,
	0xed, 0xd4, 0x0d, 0x6d, 0xa7, 0xe9, 0x74, 0x2d, 0x95, 0xd7, 0x2a, 0x1b, 0xb2, 0x4e, 0xcb, 0x86,
	0xdc, 0x19, 0xd6, 0x3c, 0x84, 0xd6, 0x49, 0x8a, 0x9e, 0x40, 0x17, 0x3f, 0x66, 0xc8, 0x05, 0xd9,
	0x80, 0x9a, 0x97, 0x50, 0x99, 0x68, 0xd5, 0xcd, 0xff, 0x92, 0x6d, 0xa8, 0x0b, 0x36, 0x62, 0x32,
	0x51, 0xd3, 0x59, 0xb1, 0xa6, 0x03, 0x2b, 0x8f, 0xee, 0xca, 0xaa, 0xe9, 0x40, 0xbb, 0x14, 0xe0,
	0x09, 0x8b, 0x39, 0xde, 0xa3, 0xa0, 0x9a, 0xd4, 0xcb, 0x26, 0x4d, 0x1b, 0x9a, 0x2e, 0x7a, 0xc1,
	0x7c, 0xcb, 0xdb, 0x84, 0x97, 0xb0, 0xa6, 0x08, 0x73, 0x2d, 0x16, 0x87, 0x3c, 0x84, 0xd6, 0xbb,
	0x24, 0xf8, 0x8f, 0x2e, 0x0f, 0xa0, 0x5d, 0x0a, 0xcc, 0x8d, 0xd0, 0x81, 0xe5, 0x4c, 0x62, 0xca,
	0xe4, 0xe5, 0xd2, 0x1c, 0x40, 0x6b, 0x84, 0x11, 0x2e, 0xb2, 0xbf, 0xdd, 0xf1, 0x01, 0xb4, 0x4b,
	0xca, 0x22, 0xc3, 0x40, 0x62, 0x66, 0x86, 0xc5, 0xd2, 0x34, 0xa1, 0x9d, 0x9f, 0xd7, 0x71, 0x14,
	0xcd, 0x75, 0x34, 0x4f, 0x60, 0x7d, 0x86, 0x99, 0x6b, 0xf1, 0x04, 0x96, 0xf2, 0xfe, 0x79, 0x47,
	0x37, 0x6a, 0x95, 0x63, 0x51, 0x65, 0xe7, 0x47, 0x0d, 0x9a, 0xf9, 0xfa, 0xad, 0xfa, 0x9c, 0xc8,
	0x1b, 0x58, 0x2e, 0x44, 0x09, 0xc9, 0xb1, 0xd5, 0x14, 0xdd, 0xcd, 0x4a, 0x4d, 0xb9, 0x9a, 0x5b,
	0xdf, 0xfe, 0x5c, 0xfd, 0xd2, 0xdb, 0x64, 0xcd, 0x9e, 0x0e, 0xec, 0xfc, 0xe3, 0xb4, 0xbd, 0x28,
	0x22, 0x23, 0x68, 0xa8, 0xb9, 0x22, 0x0f, 0x72, 0x52, 0x65, 0x48, 0xbb, 0xe4, 0x66, 0xa9, 0x90,
	0xd9, 0x94, 0x32, 0x2d, 0x73, 0xa5, 0x94, 0x19, 0x6a, 0x3d, 0x72, 0x04, 0xf5, 0xdc, 0x8e, 0xac,
	0x97, 0xc6, 0xa5, 0xc2, 0xc6, 0x75, 0xa1, 0xe0, 0x3f, 0x94, 0xfc, 0x75, 0xd2, 0x9a, 0xc5, 0xf8,
	0x42, 0x83, 0x4b, 0x12, 0x42, 0x43, 0xdd, 0xbc, 0xca, 0x51, 0x19, 0x23, 0x95, 0xa3, 0x3a, 0x18,
	0xe6, 0xbe, 0xd4, 0xd9, 0xed, 0x92, 0x6b, 0x9d, 0xfc, 0xac, 0x2c, 0x1a, 0x5c, 0x0e, 0xb5, 0xde,
	0xfb, 0x47, 0xce, 0xfd, 0x1b, 0xe4, 0x35, 0x34, 0xd4, 0x8d, 0x2b, 0xa3, 0xca, 0xc0, 0x28, 0xa3,
	0xea, 0x40, 0x94, 0x81, 0x7b, 0xd5, 0xc0, 0xaf, 0xfe, 0x6a, 0x3f, 0x8f, 0xaf, 0x34, 0xf2, 0x15,
	0xd6, 0xf2, 0x8b, 0x31, 0x8a, 0x87, 0xce, 0xa4, 0xf0, 0x38, 0x4c, 0x13, 0xbf, 0x2f, 0x5f, 0xbf,
	0x14, 0xb9, 0xe8, 0x4f, 0xa8, 0x9f, 0xb2, 0x62, 0x9b, 0xec, 0x8f, 0x85, 0x48, 0xf8, 0xd0, 0xb6,
	0x43, 0x2a, 0xc6, 0xd9, 0x99, 0xe5, 0xb3, 0x89, 0x7d, 0x4e, 0xbd, 0x38, 0xc4, 0x0b, 0x7b, 0x01,
	0xaf, 0x4b, 0x0a, 0xd0, 0x91, 0x17, 0xd1, 0xcf, 0x59, 0x9c, 0xf3, 0x9c, 0xda, 0xc0, 0xda, 0xed,
	0x69, 0x9a, 0xb3, 0xe1, 0x25, 0x49, 0x44, 0x7d, 0xf9, 0xf0, 0xd9, 0xe7, 0x9c, 0xc5, 0xc3, 0x3b,
	0x15, 0xf7, 0x05, 0xd4, 0xf6, 0x76, 0xf7, 0xc8, 0x1e, 0xf4, 0x5c, 0x14, 0x59, 0x1a, 0x63, 0x60,
	0x5c, 0x8c, 0x31, 0x36, 0xc4, 0x18, 0x8d, 0x14, 0x39, 0xcb, 0x52, 0x1f, 0x8d, 0x80, 0x21, 0x37,
	0x62, 0x26, 0x0c, 0xfc, 0x44, 0xb9, 0xb0, 0x48, 0x03, 0xea, 0xbf, 0x75, 0x6d, 0xf9, 0xac, 0x21,
	0x9f, 0xb6, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x12, 0x9c, 0xc9, 0xd3, 0x05, 0x00,
	0x00,
}
