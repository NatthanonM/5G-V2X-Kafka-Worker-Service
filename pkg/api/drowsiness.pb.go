// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.0
// source: drowsiness.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DrowsinessData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string               `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	CarId        string               `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Time         *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	ResponseTime float64              `protobuf:"fixed64,4,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	WorkingHour  float64              `protobuf:"fixed64,5,opt,name=working_hour,json=workingHour,proto3" json:"working_hour,omitempty"`
	Latitude     float64              `protobuf:"fixed64,6,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude    float64              `protobuf:"fixed64,7,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *DrowsinessData) Reset() {
	*x = DrowsinessData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drowsiness_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DrowsinessData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DrowsinessData) ProtoMessage() {}

func (x *DrowsinessData) ProtoReflect() protoreflect.Message {
	mi := &file_drowsiness_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DrowsinessData.ProtoReflect.Descriptor instead.
func (*DrowsinessData) Descriptor() ([]byte, []int) {
	return file_drowsiness_proto_rawDescGZIP(), []int{0}
}

func (x *DrowsinessData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DrowsinessData) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *DrowsinessData) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *DrowsinessData) GetResponseTime() float64 {
	if x != nil {
		return x.ResponseTime
	}
	return 0
}

func (x *DrowsinessData) GetWorkingHour() float64 {
	if x != nil {
		return x.WorkingHour
	}
	return 0
}

func (x *DrowsinessData) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *DrowsinessData) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type CreateDrowsinessDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DrowsinessId string `protobuf:"bytes,1,opt,name=drowsiness_id,json=drowsinessId,proto3" json:"drowsiness_id,omitempty"`
}

func (x *CreateDrowsinessDataResponse) Reset() {
	*x = CreateDrowsinessDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drowsiness_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDrowsinessDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDrowsinessDataResponse) ProtoMessage() {}

func (x *CreateDrowsinessDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_drowsiness_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDrowsinessDataResponse.ProtoReflect.Descriptor instead.
func (*CreateDrowsinessDataResponse) Descriptor() ([]byte, []int) {
	return file_drowsiness_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDrowsinessDataResponse) GetDrowsinessId() string {
	if x != nil {
		return x.DrowsinessId
	}
	return ""
}

var File_drowsiness_proto protoreflect.FileDescriptor

var file_drowsiness_proto_rawDesc = []byte{
	0x0a, 0x10, 0x64, 0x72, 0x6f, 0x77, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x0e, 0x44,
	0x72, 0x6f, 0x77, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x43, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x6f, 0x77,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x72, 0x6f, 0x77, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x72, 0x6f, 0x77, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x32, 0x69, 0x0a, 0x11, 0x44, 0x72, 0x6f, 0x77, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x6f, 0x77, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x72, 0x6f,
	0x77, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x6f, 0x77, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_drowsiness_proto_rawDescOnce sync.Once
	file_drowsiness_proto_rawDescData = file_drowsiness_proto_rawDesc
)

func file_drowsiness_proto_rawDescGZIP() []byte {
	file_drowsiness_proto_rawDescOnce.Do(func() {
		file_drowsiness_proto_rawDescData = protoimpl.X.CompressGZIP(file_drowsiness_proto_rawDescData)
	})
	return file_drowsiness_proto_rawDescData
}

var file_drowsiness_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_drowsiness_proto_goTypes = []interface{}{
	(*DrowsinessData)(nil),               // 0: proto.DrowsinessData
	(*CreateDrowsinessDataResponse)(nil), // 1: proto.CreateDrowsinessDataResponse
	(*timestamp.Timestamp)(nil),          // 2: google.protobuf.Timestamp
}
var file_drowsiness_proto_depIdxs = []int32{
	2, // 0: proto.DrowsinessData.time:type_name -> google.protobuf.Timestamp
	0, // 1: proto.DrowsinessService.CreateDrowsinessData:input_type -> proto.DrowsinessData
	1, // 2: proto.DrowsinessService.CreateDrowsinessData:output_type -> proto.CreateDrowsinessDataResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_drowsiness_proto_init() }
func file_drowsiness_proto_init() {
	if File_drowsiness_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drowsiness_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DrowsinessData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drowsiness_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDrowsinessDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drowsiness_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_drowsiness_proto_goTypes,
		DependencyIndexes: file_drowsiness_proto_depIdxs,
		MessageInfos:      file_drowsiness_proto_msgTypes,
	}.Build()
	File_drowsiness_proto = out.File
	file_drowsiness_proto_rawDesc = nil
	file_drowsiness_proto_goTypes = nil
	file_drowsiness_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DrowsinessServiceClient is the client API for DrowsinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DrowsinessServiceClient interface {
	CreateDrowsinessData(ctx context.Context, in *DrowsinessData, opts ...grpc.CallOption) (*CreateDrowsinessDataResponse, error)
}

type drowsinessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDrowsinessServiceClient(cc grpc.ClientConnInterface) DrowsinessServiceClient {
	return &drowsinessServiceClient{cc}
}

func (c *drowsinessServiceClient) CreateDrowsinessData(ctx context.Context, in *DrowsinessData, opts ...grpc.CallOption) (*CreateDrowsinessDataResponse, error) {
	out := new(CreateDrowsinessDataResponse)
	err := c.cc.Invoke(ctx, "/proto.DrowsinessService/CreateDrowsinessData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DrowsinessServiceServer is the server API for DrowsinessService service.
type DrowsinessServiceServer interface {
	CreateDrowsinessData(context.Context, *DrowsinessData) (*CreateDrowsinessDataResponse, error)
}

// UnimplementedDrowsinessServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDrowsinessServiceServer struct {
}

func (*UnimplementedDrowsinessServiceServer) CreateDrowsinessData(context.Context, *DrowsinessData) (*CreateDrowsinessDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDrowsinessData not implemented")
}

func RegisterDrowsinessServiceServer(s *grpc.Server, srv DrowsinessServiceServer) {
	s.RegisterService(&_DrowsinessService_serviceDesc, srv)
}

func _DrowsinessService_CreateDrowsinessData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrowsinessData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DrowsinessServiceServer).CreateDrowsinessData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DrowsinessService/CreateDrowsinessData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DrowsinessServiceServer).CreateDrowsinessData(ctx, req.(*DrowsinessData))
	}
	return interceptor(ctx, in, info, handler)
}

var _DrowsinessService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DrowsinessService",
	HandlerType: (*DrowsinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDrowsinessData",
			Handler:    _DrowsinessService_CreateDrowsinessData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "drowsiness.proto",
}
