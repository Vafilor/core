// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.1
// source: workspace_template.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateWorkspaceTemplateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace         string             `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	WorkspaceTemplate *WorkspaceTemplate `protobuf:"bytes,2,opt,name=workspaceTemplate,proto3" json:"workspaceTemplate,omitempty"`
}

func (x *CreateWorkspaceTemplateRequest) Reset() {
	*x = CreateWorkspaceTemplateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWorkspaceTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkspaceTemplateRequest) ProtoMessage() {}

func (x *CreateWorkspaceTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWorkspaceTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateWorkspaceTemplateRequest) Descriptor() ([]byte, []int) {
	return file_workspace_template_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWorkspaceTemplateRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *CreateWorkspaceTemplateRequest) GetWorkspaceTemplate() *WorkspaceTemplate {
	if x != nil {
		return x.WorkspaceTemplate
	}
	return nil
}

type WorkspaceTemplate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid              string            `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name             string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version          int64             `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Manifest         string            `protobuf:"bytes,4,opt,name=manifest,proto3" json:"manifest,omitempty"`
	CreatedAt        string            `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	WorkflowTemplate *WorkflowTemplate `protobuf:"bytes,6,opt,name=workflowTemplate,proto3" json:"workflowTemplate,omitempty"`
}

func (x *WorkspaceTemplate) Reset() {
	*x = WorkspaceTemplate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceTemplate) ProtoMessage() {}

func (x *WorkspaceTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceTemplate.ProtoReflect.Descriptor instead.
func (*WorkspaceTemplate) Descriptor() ([]byte, []int) {
	return file_workspace_template_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceTemplate) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *WorkspaceTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkspaceTemplate) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *WorkspaceTemplate) GetManifest() string {
	if x != nil {
		return x.Manifest
	}
	return ""
}

func (x *WorkspaceTemplate) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *WorkspaceTemplate) GetWorkflowTemplate() *WorkflowTemplate {
	if x != nil {
		return x.WorkflowTemplate
	}
	return nil
}

var File_workspace_template_proto protoreflect.FileDescriptor

var file_workspace_template_proto_rawDesc = []byte{
	0x0a, 0x18, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x84, 0x01, 0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x44, 0x0a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x52, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x11, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0xbd, 0x01, 0x0a, 0x18, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x48, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x42, 0x22, 0x2d, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x11, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_workspace_template_proto_rawDescOnce sync.Once
	file_workspace_template_proto_rawDescData = file_workspace_template_proto_rawDesc
)

func file_workspace_template_proto_rawDescGZIP() []byte {
	file_workspace_template_proto_rawDescOnce.Do(func() {
		file_workspace_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_template_proto_rawDescData)
	})
	return file_workspace_template_proto_rawDescData
}

var file_workspace_template_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_workspace_template_proto_goTypes = []interface{}{
	(*CreateWorkspaceTemplateRequest)(nil), // 0: api.CreateWorkspaceTemplateRequest
	(*WorkspaceTemplate)(nil),              // 1: api.WorkspaceTemplate
	(*WorkflowTemplate)(nil),               // 2: api.WorkflowTemplate
}
var file_workspace_template_proto_depIdxs = []int32{
	1, // 0: api.CreateWorkspaceTemplateRequest.workspaceTemplate:type_name -> api.WorkspaceTemplate
	2, // 1: api.WorkspaceTemplate.workflowTemplate:type_name -> api.WorkflowTemplate
	0, // 2: api.WorkspaceTemplateService.CreateWorkspaceTemplate:input_type -> api.CreateWorkspaceTemplateRequest
	1, // 3: api.WorkspaceTemplateService.CreateWorkspaceTemplate:output_type -> api.WorkspaceTemplate
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_workspace_template_proto_init() }
func file_workspace_template_proto_init() {
	if File_workspace_template_proto != nil {
		return
	}
	file_workflow_template_proto_init()
	file_label_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_workspace_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWorkspaceTemplateRequest); i {
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
		file_workspace_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceTemplate); i {
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
			RawDescriptor: file_workspace_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workspace_template_proto_goTypes,
		DependencyIndexes: file_workspace_template_proto_depIdxs,
		MessageInfos:      file_workspace_template_proto_msgTypes,
	}.Build()
	File_workspace_template_proto = out.File
	file_workspace_template_proto_rawDesc = nil
	file_workspace_template_proto_goTypes = nil
	file_workspace_template_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkspaceTemplateServiceClient is the client API for WorkspaceTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkspaceTemplateServiceClient interface {
	// Creates a Workflow
	CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error)
}

type workspaceTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkspaceTemplateServiceClient(cc grpc.ClientConnInterface) WorkspaceTemplateServiceClient {
	return &workspaceTemplateServiceClient{cc}
}

func (c *workspaceTemplateServiceClient) CreateWorkspaceTemplate(ctx context.Context, in *CreateWorkspaceTemplateRequest, opts ...grpc.CallOption) (*WorkspaceTemplate, error) {
	out := new(WorkspaceTemplate)
	err := c.cc.Invoke(ctx, "/api.WorkspaceTemplateService/CreateWorkspaceTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkspaceTemplateServiceServer is the server API for WorkspaceTemplateService service.
type WorkspaceTemplateServiceServer interface {
	// Creates a Workflow
	CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error)
}

// UnimplementedWorkspaceTemplateServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkspaceTemplateServiceServer struct {
}

func (*UnimplementedWorkspaceTemplateServiceServer) CreateWorkspaceTemplate(context.Context, *CreateWorkspaceTemplateRequest) (*WorkspaceTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWorkspaceTemplate not implemented")
}

func RegisterWorkspaceTemplateServiceServer(s *grpc.Server, srv WorkspaceTemplateServiceServer) {
	s.RegisterService(&_WorkspaceTemplateService_serviceDesc, srv)
}

func _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkspaceTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.WorkspaceTemplateService/CreateWorkspaceTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkspaceTemplateServiceServer).CreateWorkspaceTemplate(ctx, req.(*CreateWorkspaceTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkspaceTemplateService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.WorkspaceTemplateService",
	HandlerType: (*WorkspaceTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkspaceTemplate",
			Handler:    _WorkspaceTemplateService_CreateWorkspaceTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workspace_template.proto",
}
