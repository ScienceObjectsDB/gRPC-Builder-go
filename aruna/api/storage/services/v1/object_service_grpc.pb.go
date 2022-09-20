// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: aruna/api/storage/services/v1/object_service.proto

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

// ObjectServiceClient is the client API for ObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectServiceClient interface {
	// This initializes a new object
	// Initializing an object will put it in a staging area.
	// Staged objects will get a separate staging id and need to be finished
	// before they can be used.
	InitializeNewObject(ctx context.Context, in *InitializeNewObjectRequest, opts ...grpc.CallOption) (*InitializeNewObjectResponse, error)
	// This method will return a (multi-part) url that can be used to upload a
	// file to S3. Part is a optional query parameter that can be used to upload a
	// part of the file / multipart upload.
	GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error)
	// This method will return a url that can be used to download a file from S3.
	GetDownloadURL(ctx context.Context, in *GetDownloadURLRequest, opts ...grpc.CallOption) (*GetDownloadURLResponse, error)
	// This method can be used to get download urls for multiple objects.
	// The order of the returned urls will be the same as the order of the object
	// ids in the request.
	GetDownloadLinksBatch(ctx context.Context, in *GetDownloadLinksBatchRequest, opts ...grpc.CallOption) (*GetDownloadLinksBatchResponse, error)
	// Creates a stream of objects and presigned links based on the provided query
	// This can be used retrieve a large number of Objects as a stream that would
	// otherwise cause issues with the connection
	CreateDownloadLinksStream(ctx context.Context, in *CreateDownloadLinksStreamRequest, opts ...grpc.CallOption) (ObjectService_CreateDownloadLinksStreamClient, error)
	// This method completes the staging of an object.
	FinishObjectStaging(ctx context.Context, in *FinishObjectStagingRequest, opts ...grpc.CallOption) (*FinishObjectStagingResponse, error)
	// Objects are immutable!
	// Updating an object will create a new revision for the object
	// This method will put the new revision in a staging area.
	// Staged objects will get a separate staging id and need to be finished
	// before they can be used.
	UpdateObject(ctx context.Context, in *UpdateObjectRequest, opts ...grpc.CallOption) (*UpdateObjectResponse, error)
	CreateObjectReference(ctx context.Context, in *CreateObjectReferenceRequest, opts ...grpc.CallOption) (*CreateObjectReferenceResponse, error)
	// This method clones an object and creates a copy in the same collection.
	// This copy has a new id and revision and will not receive any updates from
	// the original object.
	CloneObject(ctx context.Context, in *CloneObjectRequest, opts ...grpc.CallOption) (*CloneObjectResponse, error)
	// Deletes the object with the complete revision history.
	// This should be avoided if possible.
	// This method allows the owner to cascade the deletion of all objects that
	// were cloned from this object.
	// -> GDPR compliant procedure.
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	// GetObjectByID gets a specific Object by ID that is associated to the
	// current collection By default only the latest revision of an object will be
	// returned Specify a revision_number to select an older revision With the
	// optional with_url boolean a download link can automatically be requested
	GetObjectByID(ctx context.Context, in *GetObjectByIDRequest, opts ...grpc.CallOption) (*GetObjectByIDResponse, error)
	// GetObjects returns a (paginated) list of objects in a specific collection
	// By default only the latest revisions of all objects will be shown
	// This behaviour can be changed with the include_history flag
	// With the optional with_url boolean a download link can automatically be
	// requested for each Object This request contains a LabelOrIDQuery message,
	// this is either a list of request ObjectIDs or a query filtered by Labels
	GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error)
	// GetObjectRevisions returns the full list of revisions of a specified object
	// With the optional with_url boolean a download link can automatically be
	// requested for each Object This is by default a paginated request
	GetObjectRevisions(ctx context.Context, in *GetObjectRevisionsRequest, opts ...grpc.CallOption) (*GetObjectRevisionsResponse, error)
	// GetLatestObjectRevision returns the latest revision of a specific object
	// The returned `latest` object will have a different id if the current
	// object is not the latest revision
	GetLatestObjectRevision(ctx context.Context, in *GetLatestObjectRevisionRequest, opts ...grpc.CallOption) (*GetLatestObjectRevisionResponse, error)
	// GetObjectEndpoints returns a list of endpoints
	// One endpoint will be the "default" endpoint
	GetObjectEndpoints(ctx context.Context, in *GetObjectEndpointsRequest, opts ...grpc.CallOption) (*GetObjectEndpointsResponse, error)
	// AddLabelToObject is a specific request to add a new label
	// to an existing object, in contrast to UpdateObject
	// this will not create a new object in the staging area
	// Instead it will directly add the specified label(s) to the object
	AddLabelToObject(ctx context.Context, in *AddLabelToObjectRequest, opts ...grpc.CallOption) (*AddLabelToObjectResponse, error)
	// SetHooksOfObject is a specific request to update the complete list
	// of hooks for a specific object. This will not update the object
	// and create a new id, instead it will overwrite all hooks of the existing
	// object.
	SetHooksOfObject(ctx context.Context, in *SetHooksOfObjectRequest, opts ...grpc.CallOption) (*SetHooksOfObjectResponse, error)
	// Get a list of references for this object (optional) including all revisions
	GetReferences(ctx context.Context, in *GetReferencesRequest, opts ...grpc.CallOption) (*GetReferencesResponse, error)
}

type objectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectServiceClient(cc grpc.ClientConnInterface) ObjectServiceClient {
	return &objectServiceClient{cc}
}

func (c *objectServiceClient) InitializeNewObject(ctx context.Context, in *InitializeNewObjectRequest, opts ...grpc.CallOption) (*InitializeNewObjectResponse, error) {
	out := new(InitializeNewObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/InitializeNewObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetUploadURL(ctx context.Context, in *GetUploadURLRequest, opts ...grpc.CallOption) (*GetUploadURLResponse, error) {
	out := new(GetUploadURLResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetUploadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetDownloadURL(ctx context.Context, in *GetDownloadURLRequest, opts ...grpc.CallOption) (*GetDownloadURLResponse, error) {
	out := new(GetDownloadURLResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetDownloadURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetDownloadLinksBatch(ctx context.Context, in *GetDownloadLinksBatchRequest, opts ...grpc.CallOption) (*GetDownloadLinksBatchResponse, error) {
	out := new(GetDownloadLinksBatchResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetDownloadLinksBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) CreateDownloadLinksStream(ctx context.Context, in *CreateDownloadLinksStreamRequest, opts ...grpc.CallOption) (ObjectService_CreateDownloadLinksStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ObjectService_ServiceDesc.Streams[0], "/aruna.api.storage.services.v1.ObjectService/CreateDownloadLinksStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectServiceCreateDownloadLinksStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectService_CreateDownloadLinksStreamClient interface {
	Recv() (*CreateDownloadLinksStreamResponse, error)
	grpc.ClientStream
}

type objectServiceCreateDownloadLinksStreamClient struct {
	grpc.ClientStream
}

func (x *objectServiceCreateDownloadLinksStreamClient) Recv() (*CreateDownloadLinksStreamResponse, error) {
	m := new(CreateDownloadLinksStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectServiceClient) FinishObjectStaging(ctx context.Context, in *FinishObjectStagingRequest, opts ...grpc.CallOption) (*FinishObjectStagingResponse, error) {
	out := new(FinishObjectStagingResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/FinishObjectStaging", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) UpdateObject(ctx context.Context, in *UpdateObjectRequest, opts ...grpc.CallOption) (*UpdateObjectResponse, error) {
	out := new(UpdateObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/UpdateObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) CreateObjectReference(ctx context.Context, in *CreateObjectReferenceRequest, opts ...grpc.CallOption) (*CreateObjectReferenceResponse, error) {
	out := new(CreateObjectReferenceResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/CreateObjectReference", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) CloneObject(ctx context.Context, in *CloneObjectRequest, opts ...grpc.CallOption) (*CloneObjectResponse, error) {
	out := new(CloneObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/CloneObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetObjectByID(ctx context.Context, in *GetObjectByIDRequest, opts ...grpc.CallOption) (*GetObjectByIDResponse, error) {
	out := new(GetObjectByIDResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetObjectByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetObjects(ctx context.Context, in *GetObjectsRequest, opts ...grpc.CallOption) (*GetObjectsResponse, error) {
	out := new(GetObjectsResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetObjectRevisions(ctx context.Context, in *GetObjectRevisionsRequest, opts ...grpc.CallOption) (*GetObjectRevisionsResponse, error) {
	out := new(GetObjectRevisionsResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetObjectRevisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetLatestObjectRevision(ctx context.Context, in *GetLatestObjectRevisionRequest, opts ...grpc.CallOption) (*GetLatestObjectRevisionResponse, error) {
	out := new(GetLatestObjectRevisionResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetLatestObjectRevision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetObjectEndpoints(ctx context.Context, in *GetObjectEndpointsRequest, opts ...grpc.CallOption) (*GetObjectEndpointsResponse, error) {
	out := new(GetObjectEndpointsResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetObjectEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) AddLabelToObject(ctx context.Context, in *AddLabelToObjectRequest, opts ...grpc.CallOption) (*AddLabelToObjectResponse, error) {
	out := new(AddLabelToObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/AddLabelToObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) SetHooksOfObject(ctx context.Context, in *SetHooksOfObjectRequest, opts ...grpc.CallOption) (*SetHooksOfObjectResponse, error) {
	out := new(SetHooksOfObjectResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/SetHooksOfObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) GetReferences(ctx context.Context, in *GetReferencesRequest, opts ...grpc.CallOption) (*GetReferencesResponse, error) {
	out := new(GetReferencesResponse)
	err := c.cc.Invoke(ctx, "/aruna.api.storage.services.v1.ObjectService/GetReferences", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServiceServer is the server API for ObjectService service.
// All implementations should embed UnimplementedObjectServiceServer
// for forward compatibility
type ObjectServiceServer interface {
	// This initializes a new object
	// Initializing an object will put it in a staging area.
	// Staged objects will get a separate staging id and need to be finished
	// before they can be used.
	InitializeNewObject(context.Context, *InitializeNewObjectRequest) (*InitializeNewObjectResponse, error)
	// This method will return a (multi-part) url that can be used to upload a
	// file to S3. Part is a optional query parameter that can be used to upload a
	// part of the file / multipart upload.
	GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error)
	// This method will return a url that can be used to download a file from S3.
	GetDownloadURL(context.Context, *GetDownloadURLRequest) (*GetDownloadURLResponse, error)
	// This method can be used to get download urls for multiple objects.
	// The order of the returned urls will be the same as the order of the object
	// ids in the request.
	GetDownloadLinksBatch(context.Context, *GetDownloadLinksBatchRequest) (*GetDownloadLinksBatchResponse, error)
	// Creates a stream of objects and presigned links based on the provided query
	// This can be used retrieve a large number of Objects as a stream that would
	// otherwise cause issues with the connection
	CreateDownloadLinksStream(*CreateDownloadLinksStreamRequest, ObjectService_CreateDownloadLinksStreamServer) error
	// This method completes the staging of an object.
	FinishObjectStaging(context.Context, *FinishObjectStagingRequest) (*FinishObjectStagingResponse, error)
	// Objects are immutable!
	// Updating an object will create a new revision for the object
	// This method will put the new revision in a staging area.
	// Staged objects will get a separate staging id and need to be finished
	// before they can be used.
	UpdateObject(context.Context, *UpdateObjectRequest) (*UpdateObjectResponse, error)
	CreateObjectReference(context.Context, *CreateObjectReferenceRequest) (*CreateObjectReferenceResponse, error)
	// This method clones an object and creates a copy in the same collection.
	// This copy has a new id and revision and will not receive any updates from
	// the original object.
	CloneObject(context.Context, *CloneObjectRequest) (*CloneObjectResponse, error)
	// Deletes the object with the complete revision history.
	// This should be avoided if possible.
	// This method allows the owner to cascade the deletion of all objects that
	// were cloned from this object.
	// -> GDPR compliant procedure.
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	// GetObjectByID gets a specific Object by ID that is associated to the
	// current collection By default only the latest revision of an object will be
	// returned Specify a revision_number to select an older revision With the
	// optional with_url boolean a download link can automatically be requested
	GetObjectByID(context.Context, *GetObjectByIDRequest) (*GetObjectByIDResponse, error)
	// GetObjects returns a (paginated) list of objects in a specific collection
	// By default only the latest revisions of all objects will be shown
	// This behaviour can be changed with the include_history flag
	// With the optional with_url boolean a download link can automatically be
	// requested for each Object This request contains a LabelOrIDQuery message,
	// this is either a list of request ObjectIDs or a query filtered by Labels
	GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error)
	// GetObjectRevisions returns the full list of revisions of a specified object
	// With the optional with_url boolean a download link can automatically be
	// requested for each Object This is by default a paginated request
	GetObjectRevisions(context.Context, *GetObjectRevisionsRequest) (*GetObjectRevisionsResponse, error)
	// GetLatestObjectRevision returns the latest revision of a specific object
	// The returned `latest` object will have a different id if the current
	// object is not the latest revision
	GetLatestObjectRevision(context.Context, *GetLatestObjectRevisionRequest) (*GetLatestObjectRevisionResponse, error)
	// GetObjectEndpoints returns a list of endpoints
	// One endpoint will be the "default" endpoint
	GetObjectEndpoints(context.Context, *GetObjectEndpointsRequest) (*GetObjectEndpointsResponse, error)
	// AddLabelToObject is a specific request to add a new label
	// to an existing object, in contrast to UpdateObject
	// this will not create a new object in the staging area
	// Instead it will directly add the specified label(s) to the object
	AddLabelToObject(context.Context, *AddLabelToObjectRequest) (*AddLabelToObjectResponse, error)
	// SetHooksOfObject is a specific request to update the complete list
	// of hooks for a specific object. This will not update the object
	// and create a new id, instead it will overwrite all hooks of the existing
	// object.
	SetHooksOfObject(context.Context, *SetHooksOfObjectRequest) (*SetHooksOfObjectResponse, error)
	// Get a list of references for this object (optional) including all revisions
	GetReferences(context.Context, *GetReferencesRequest) (*GetReferencesResponse, error)
}

// UnimplementedObjectServiceServer should be embedded to have forward compatible implementations.
type UnimplementedObjectServiceServer struct {
}

func (UnimplementedObjectServiceServer) InitializeNewObject(context.Context, *InitializeNewObjectRequest) (*InitializeNewObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitializeNewObject not implemented")
}
func (UnimplementedObjectServiceServer) GetUploadURL(context.Context, *GetUploadURLRequest) (*GetUploadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadURL not implemented")
}
func (UnimplementedObjectServiceServer) GetDownloadURL(context.Context, *GetDownloadURLRequest) (*GetDownloadURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadURL not implemented")
}
func (UnimplementedObjectServiceServer) GetDownloadLinksBatch(context.Context, *GetDownloadLinksBatchRequest) (*GetDownloadLinksBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadLinksBatch not implemented")
}
func (UnimplementedObjectServiceServer) CreateDownloadLinksStream(*CreateDownloadLinksStreamRequest, ObjectService_CreateDownloadLinksStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateDownloadLinksStream not implemented")
}
func (UnimplementedObjectServiceServer) FinishObjectStaging(context.Context, *FinishObjectStagingRequest) (*FinishObjectStagingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishObjectStaging not implemented")
}
func (UnimplementedObjectServiceServer) UpdateObject(context.Context, *UpdateObjectRequest) (*UpdateObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObject not implemented")
}
func (UnimplementedObjectServiceServer) CreateObjectReference(context.Context, *CreateObjectReferenceRequest) (*CreateObjectReferenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObjectReference not implemented")
}
func (UnimplementedObjectServiceServer) CloneObject(context.Context, *CloneObjectRequest) (*CloneObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneObject not implemented")
}
func (UnimplementedObjectServiceServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedObjectServiceServer) GetObjectByID(context.Context, *GetObjectByIDRequest) (*GetObjectByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectByID not implemented")
}
func (UnimplementedObjectServiceServer) GetObjects(context.Context, *GetObjectsRequest) (*GetObjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (UnimplementedObjectServiceServer) GetObjectRevisions(context.Context, *GetObjectRevisionsRequest) (*GetObjectRevisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectRevisions not implemented")
}
func (UnimplementedObjectServiceServer) GetLatestObjectRevision(context.Context, *GetLatestObjectRevisionRequest) (*GetLatestObjectRevisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestObjectRevision not implemented")
}
func (UnimplementedObjectServiceServer) GetObjectEndpoints(context.Context, *GetObjectEndpointsRequest) (*GetObjectEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectEndpoints not implemented")
}
func (UnimplementedObjectServiceServer) AddLabelToObject(context.Context, *AddLabelToObjectRequest) (*AddLabelToObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLabelToObject not implemented")
}
func (UnimplementedObjectServiceServer) SetHooksOfObject(context.Context, *SetHooksOfObjectRequest) (*SetHooksOfObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHooksOfObject not implemented")
}
func (UnimplementedObjectServiceServer) GetReferences(context.Context, *GetReferencesRequest) (*GetReferencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferences not implemented")
}

// UnsafeObjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectServiceServer will
// result in compilation errors.
type UnsafeObjectServiceServer interface {
	mustEmbedUnimplementedObjectServiceServer()
}

func RegisterObjectServiceServer(s grpc.ServiceRegistrar, srv ObjectServiceServer) {
	s.RegisterService(&ObjectService_ServiceDesc, srv)
}

func _ObjectService_InitializeNewObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeNewObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).InitializeNewObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/InitializeNewObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).InitializeNewObject(ctx, req.(*InitializeNewObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetUploadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetUploadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetUploadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetUploadURL(ctx, req.(*GetUploadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetDownloadURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetDownloadURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetDownloadURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetDownloadURL(ctx, req.(*GetDownloadURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetDownloadLinksBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadLinksBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetDownloadLinksBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetDownloadLinksBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetDownloadLinksBatch(ctx, req.(*GetDownloadLinksBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_CreateDownloadLinksStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateDownloadLinksStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectServiceServer).CreateDownloadLinksStream(m, &objectServiceCreateDownloadLinksStreamServer{stream})
}

type ObjectService_CreateDownloadLinksStreamServer interface {
	Send(*CreateDownloadLinksStreamResponse) error
	grpc.ServerStream
}

type objectServiceCreateDownloadLinksStreamServer struct {
	grpc.ServerStream
}

func (x *objectServiceCreateDownloadLinksStreamServer) Send(m *CreateDownloadLinksStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectService_FinishObjectStaging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishObjectStagingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).FinishObjectStaging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/FinishObjectStaging",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).FinishObjectStaging(ctx, req.(*FinishObjectStagingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_UpdateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).UpdateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/UpdateObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).UpdateObject(ctx, req.(*UpdateObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_CreateObjectReference_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectReferenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).CreateObjectReference(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/CreateObjectReference",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).CreateObjectReference(ctx, req.(*CreateObjectReferenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_CloneObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).CloneObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/CloneObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).CloneObject(ctx, req.(*CloneObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetObjectByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetObjectByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetObjectByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetObjectByID(ctx, req.(*GetObjectByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetObjects(ctx, req.(*GetObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetObjectRevisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRevisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetObjectRevisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetObjectRevisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetObjectRevisions(ctx, req.(*GetObjectRevisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetLatestObjectRevision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestObjectRevisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetLatestObjectRevision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetLatestObjectRevision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetLatestObjectRevision(ctx, req.(*GetLatestObjectRevisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetObjectEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetObjectEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetObjectEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetObjectEndpoints(ctx, req.(*GetObjectEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_AddLabelToObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLabelToObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).AddLabelToObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/AddLabelToObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).AddLabelToObject(ctx, req.(*AddLabelToObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_SetHooksOfObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetHooksOfObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).SetHooksOfObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/SetHooksOfObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).SetHooksOfObject(ctx, req.(*SetHooksOfObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_GetReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReferencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).GetReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aruna.api.storage.services.v1.ObjectService/GetReferences",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).GetReferences(ctx, req.(*GetReferencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectService_ServiceDesc is the grpc.ServiceDesc for ObjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aruna.api.storage.services.v1.ObjectService",
	HandlerType: (*ObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitializeNewObject",
			Handler:    _ObjectService_InitializeNewObject_Handler,
		},
		{
			MethodName: "GetUploadURL",
			Handler:    _ObjectService_GetUploadURL_Handler,
		},
		{
			MethodName: "GetDownloadURL",
			Handler:    _ObjectService_GetDownloadURL_Handler,
		},
		{
			MethodName: "GetDownloadLinksBatch",
			Handler:    _ObjectService_GetDownloadLinksBatch_Handler,
		},
		{
			MethodName: "FinishObjectStaging",
			Handler:    _ObjectService_FinishObjectStaging_Handler,
		},
		{
			MethodName: "UpdateObject",
			Handler:    _ObjectService_UpdateObject_Handler,
		},
		{
			MethodName: "CreateObjectReference",
			Handler:    _ObjectService_CreateObjectReference_Handler,
		},
		{
			MethodName: "CloneObject",
			Handler:    _ObjectService_CloneObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _ObjectService_DeleteObject_Handler,
		},
		{
			MethodName: "GetObjectByID",
			Handler:    _ObjectService_GetObjectByID_Handler,
		},
		{
			MethodName: "GetObjects",
			Handler:    _ObjectService_GetObjects_Handler,
		},
		{
			MethodName: "GetObjectRevisions",
			Handler:    _ObjectService_GetObjectRevisions_Handler,
		},
		{
			MethodName: "GetLatestObjectRevision",
			Handler:    _ObjectService_GetLatestObjectRevision_Handler,
		},
		{
			MethodName: "GetObjectEndpoints",
			Handler:    _ObjectService_GetObjectEndpoints_Handler,
		},
		{
			MethodName: "AddLabelToObject",
			Handler:    _ObjectService_AddLabelToObject_Handler,
		},
		{
			MethodName: "SetHooksOfObject",
			Handler:    _ObjectService_SetHooksOfObject_Handler,
		},
		{
			MethodName: "GetReferences",
			Handler:    _ObjectService_GetReferences_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateDownloadLinksStream",
			Handler:       _ObjectService_CreateDownloadLinksStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aruna/api/storage/services/v1/object_service.proto",
}
