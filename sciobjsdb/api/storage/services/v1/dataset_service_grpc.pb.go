// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// DatasetServiceClient is the client API for DatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatasetServiceClient interface {
	// CreateNewDataset Creates a new dataset and associates it with a dataset
	CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error)
	// Dataset Returns a specific dataset
	GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error)
	// Lists Versions of a dataset
	GetDatasetVersions(ctx context.Context, in *GetDatasetVersionsRequest, opts ...grpc.CallOption) (*GetDatasetVersionsResponse, error)
	// Lists all object groups of a dataset
	GetDatasetObjectGroups(ctx context.Context, in *GetDatasetObjectGroupsRequest, opts ...grpc.CallOption) (*GetDatasetObjectGroupsResponse, error)
	// Returns a signed link that can be used to download all objects from the
	// specified request The link is signed using hmac and the resulting data can
	// be shared without exposing any secrets
	GetObjectGroupsStreamLink(ctx context.Context, in *GetObjectGroupsStreamLinkRequest, opts ...grpc.CallOption) (*GetObjectGroupsStreamLinkResponse, error)
	// Updates a field of a dataset
	UpdateDatasetField(ctx context.Context, in *UpdateDatasetFieldRequest, opts ...grpc.CallOption) (*UpdateDatasetFieldResponse, error)
	// DeleteDataset Delete a dataset
	DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*DeleteDatasetResponse, error)
	// Returns all object groups that were created within a specific date range
	// The date range is not the date when the data was created in the system but
	// byte the externally date that indicates the actual creation of the data
	// rather than the date the data was ingested into the system
	GetObjectGroupsInDateRange(ctx context.Context, in *GetObjectGroupsInDateRangeRequest, opts ...grpc.CallOption) (*GetObjectGroupsInDateRangeResponse, error)
	// ReleaseDatasetVersion Release a new dataset version
	ReleaseDatasetVersion(ctx context.Context, in *ReleaseDatasetVersionRequest, opts ...grpc.CallOption) (*ReleaseDatasetVersionResponse, error)
	GetDatasetVersion(ctx context.Context, in *GetDatasetVersionRequest, opts ...grpc.CallOption) (*GetDatasetVersionResponse, error)
	GetDatasetVersionObjectGroups(ctx context.Context, in *GetDatasetVersionObjectGroupsRequest, opts ...grpc.CallOption) (*GetDatasetVersionObjectGroupsResponse, error)
	DeleteDatasetVersion(ctx context.Context, in *DeleteDatasetVersionRequest, opts ...grpc.CallOption) (*DeleteDatasetVersionResponse, error)
}

type datasetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetServiceClient(cc grpc.ClientConnInterface) DatasetServiceClient {
	return &datasetServiceClient{cc}
}

func (c *datasetServiceClient) CreateDataset(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error) {
	out := new(CreateDatasetResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/CreateDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDataset(ctx context.Context, in *GetDatasetRequest, opts ...grpc.CallOption) (*GetDatasetResponse, error) {
	out := new(GetDatasetResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDatasetVersions(ctx context.Context, in *GetDatasetVersionsRequest, opts ...grpc.CallOption) (*GetDatasetVersionsResponse, error) {
	out := new(GetDatasetVersionsResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDatasetObjectGroups(ctx context.Context, in *GetDatasetObjectGroupsRequest, opts ...grpc.CallOption) (*GetDatasetObjectGroupsResponse, error) {
	out := new(GetDatasetObjectGroupsResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetObjectGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetObjectGroupsStreamLink(ctx context.Context, in *GetObjectGroupsStreamLinkRequest, opts ...grpc.CallOption) (*GetObjectGroupsStreamLinkResponse, error) {
	out := new(GetObjectGroupsStreamLinkResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetObjectGroupsStreamLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) UpdateDatasetField(ctx context.Context, in *UpdateDatasetFieldRequest, opts ...grpc.CallOption) (*UpdateDatasetFieldResponse, error) {
	out := new(UpdateDatasetFieldResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/UpdateDatasetField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) DeleteDataset(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*DeleteDatasetResponse, error) {
	out := new(DeleteDatasetResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/DeleteDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetObjectGroupsInDateRange(ctx context.Context, in *GetObjectGroupsInDateRangeRequest, opts ...grpc.CallOption) (*GetObjectGroupsInDateRangeResponse, error) {
	out := new(GetObjectGroupsInDateRangeResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetObjectGroupsInDateRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) ReleaseDatasetVersion(ctx context.Context, in *ReleaseDatasetVersionRequest, opts ...grpc.CallOption) (*ReleaseDatasetVersionResponse, error) {
	out := new(ReleaseDatasetVersionResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/ReleaseDatasetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDatasetVersion(ctx context.Context, in *GetDatasetVersionRequest, opts ...grpc.CallOption) (*GetDatasetVersionResponse, error) {
	out := new(GetDatasetVersionResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDatasetVersionObjectGroups(ctx context.Context, in *GetDatasetVersionObjectGroupsRequest, opts ...grpc.CallOption) (*GetDatasetVersionObjectGroupsResponse, error) {
	out := new(GetDatasetVersionObjectGroupsResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersionObjectGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) DeleteDatasetVersion(ctx context.Context, in *DeleteDatasetVersionRequest, opts ...grpc.CallOption) (*DeleteDatasetVersionResponse, error) {
	out := new(DeleteDatasetVersionResponse)
	err := c.cc.Invoke(ctx, "/sciobjsdb.api.storage.services.v1.DatasetService/DeleteDatasetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetServiceServer is the server API for DatasetService service.
// All implementations should embed UnimplementedDatasetServiceServer
// for forward compatibility
type DatasetServiceServer interface {
	// CreateNewDataset Creates a new dataset and associates it with a dataset
	CreateDataset(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error)
	// Dataset Returns a specific dataset
	GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, error)
	// Lists Versions of a dataset
	GetDatasetVersions(context.Context, *GetDatasetVersionsRequest) (*GetDatasetVersionsResponse, error)
	// Lists all object groups of a dataset
	GetDatasetObjectGroups(context.Context, *GetDatasetObjectGroupsRequest) (*GetDatasetObjectGroupsResponse, error)
	// Returns a signed link that can be used to download all objects from the
	// specified request The link is signed using hmac and the resulting data can
	// be shared without exposing any secrets
	GetObjectGroupsStreamLink(context.Context, *GetObjectGroupsStreamLinkRequest) (*GetObjectGroupsStreamLinkResponse, error)
	// Updates a field of a dataset
	UpdateDatasetField(context.Context, *UpdateDatasetFieldRequest) (*UpdateDatasetFieldResponse, error)
	// DeleteDataset Delete a dataset
	DeleteDataset(context.Context, *DeleteDatasetRequest) (*DeleteDatasetResponse, error)
	// Returns all object groups that were created within a specific date range
	// The date range is not the date when the data was created in the system but
	// byte the externally date that indicates the actual creation of the data
	// rather than the date the data was ingested into the system
	GetObjectGroupsInDateRange(context.Context, *GetObjectGroupsInDateRangeRequest) (*GetObjectGroupsInDateRangeResponse, error)
	// ReleaseDatasetVersion Release a new dataset version
	ReleaseDatasetVersion(context.Context, *ReleaseDatasetVersionRequest) (*ReleaseDatasetVersionResponse, error)
	GetDatasetVersion(context.Context, *GetDatasetVersionRequest) (*GetDatasetVersionResponse, error)
	GetDatasetVersionObjectGroups(context.Context, *GetDatasetVersionObjectGroupsRequest) (*GetDatasetVersionObjectGroupsResponse, error)
	DeleteDatasetVersion(context.Context, *DeleteDatasetVersionRequest) (*DeleteDatasetVersionResponse, error)
}

// UnimplementedDatasetServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDatasetServiceServer struct {
}

func (UnimplementedDatasetServiceServer) CreateDataset(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDataset not implemented")
}
func (UnimplementedDatasetServiceServer) GetDataset(context.Context, *GetDatasetRequest) (*GetDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDataset not implemented")
}
func (UnimplementedDatasetServiceServer) GetDatasetVersions(context.Context, *GetDatasetVersionsRequest) (*GetDatasetVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetVersions not implemented")
}
func (UnimplementedDatasetServiceServer) GetDatasetObjectGroups(context.Context, *GetDatasetObjectGroupsRequest) (*GetDatasetObjectGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetObjectGroups not implemented")
}
func (UnimplementedDatasetServiceServer) GetObjectGroupsStreamLink(context.Context, *GetObjectGroupsStreamLinkRequest) (*GetObjectGroupsStreamLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupsStreamLink not implemented")
}
func (UnimplementedDatasetServiceServer) UpdateDatasetField(context.Context, *UpdateDatasetFieldRequest) (*UpdateDatasetFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDatasetField not implemented")
}
func (UnimplementedDatasetServiceServer) DeleteDataset(context.Context, *DeleteDatasetRequest) (*DeleteDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataset not implemented")
}
func (UnimplementedDatasetServiceServer) GetObjectGroupsInDateRange(context.Context, *GetObjectGroupsInDateRangeRequest) (*GetObjectGroupsInDateRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectGroupsInDateRange not implemented")
}
func (UnimplementedDatasetServiceServer) ReleaseDatasetVersion(context.Context, *ReleaseDatasetVersionRequest) (*ReleaseDatasetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseDatasetVersion not implemented")
}
func (UnimplementedDatasetServiceServer) GetDatasetVersion(context.Context, *GetDatasetVersionRequest) (*GetDatasetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetVersion not implemented")
}
func (UnimplementedDatasetServiceServer) GetDatasetVersionObjectGroups(context.Context, *GetDatasetVersionObjectGroupsRequest) (*GetDatasetVersionObjectGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDatasetVersionObjectGroups not implemented")
}
func (UnimplementedDatasetServiceServer) DeleteDatasetVersion(context.Context, *DeleteDatasetVersionRequest) (*DeleteDatasetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatasetVersion not implemented")
}

// UnsafeDatasetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetServiceServer will
// result in compilation errors.
type UnsafeDatasetServiceServer interface {
	mustEmbedUnimplementedDatasetServiceServer()
}

func RegisterDatasetServiceServer(s grpc.ServiceRegistrar, srv DatasetServiceServer) {
	s.RegisterService(&DatasetService_ServiceDesc, srv)
}

func _DatasetService_CreateDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).CreateDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/CreateDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).CreateDataset(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDataset(ctx, req.(*GetDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDatasetVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDatasetVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDatasetVersions(ctx, req.(*GetDatasetVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDatasetObjectGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetObjectGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDatasetObjectGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetObjectGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDatasetObjectGroups(ctx, req.(*GetDatasetObjectGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetObjectGroupsStreamLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupsStreamLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetObjectGroupsStreamLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetObjectGroupsStreamLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetObjectGroupsStreamLink(ctx, req.(*GetObjectGroupsStreamLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_UpdateDatasetField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatasetFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).UpdateDatasetField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/UpdateDatasetField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).UpdateDatasetField(ctx, req.(*UpdateDatasetFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_DeleteDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).DeleteDataset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/DeleteDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).DeleteDataset(ctx, req.(*DeleteDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetObjectGroupsInDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectGroupsInDateRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetObjectGroupsInDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetObjectGroupsInDateRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetObjectGroupsInDateRange(ctx, req.(*GetObjectGroupsInDateRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_ReleaseDatasetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseDatasetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).ReleaseDatasetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/ReleaseDatasetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).ReleaseDatasetVersion(ctx, req.(*ReleaseDatasetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDatasetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDatasetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDatasetVersion(ctx, req.(*GetDatasetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDatasetVersionObjectGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetVersionObjectGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDatasetVersionObjectGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/GetDatasetVersionObjectGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDatasetVersionObjectGroups(ctx, req.(*GetDatasetVersionObjectGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_DeleteDatasetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).DeleteDatasetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sciobjsdb.api.storage.services.v1.DatasetService/DeleteDatasetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).DeleteDatasetVersion(ctx, req.(*DeleteDatasetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetService_ServiceDesc is the grpc.ServiceDesc for DatasetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sciobjsdb.api.storage.services.v1.DatasetService",
	HandlerType: (*DatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDataset",
			Handler:    _DatasetService_CreateDataset_Handler,
		},
		{
			MethodName: "GetDataset",
			Handler:    _DatasetService_GetDataset_Handler,
		},
		{
			MethodName: "GetDatasetVersions",
			Handler:    _DatasetService_GetDatasetVersions_Handler,
		},
		{
			MethodName: "GetDatasetObjectGroups",
			Handler:    _DatasetService_GetDatasetObjectGroups_Handler,
		},
		{
			MethodName: "GetObjectGroupsStreamLink",
			Handler:    _DatasetService_GetObjectGroupsStreamLink_Handler,
		},
		{
			MethodName: "UpdateDatasetField",
			Handler:    _DatasetService_UpdateDatasetField_Handler,
		},
		{
			MethodName: "DeleteDataset",
			Handler:    _DatasetService_DeleteDataset_Handler,
		},
		{
			MethodName: "GetObjectGroupsInDateRange",
			Handler:    _DatasetService_GetObjectGroupsInDateRange_Handler,
		},
		{
			MethodName: "ReleaseDatasetVersion",
			Handler:    _DatasetService_ReleaseDatasetVersion_Handler,
		},
		{
			MethodName: "GetDatasetVersion",
			Handler:    _DatasetService_GetDatasetVersion_Handler,
		},
		{
			MethodName: "GetDatasetVersionObjectGroups",
			Handler:    _DatasetService_GetDatasetVersionObjectGroups_Handler,
		},
		{
			MethodName: "DeleteDatasetVersion",
			Handler:    _DatasetService_DeleteDatasetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sciobjsdb/api/storage/services/v1/dataset_service.proto",
}
