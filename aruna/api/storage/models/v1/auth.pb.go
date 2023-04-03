// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: aruna/api/storage/models/v1/auth.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Permission int32

const (
	Permission_PERMISSION_UNSPECIFIED Permission = 0
	Permission_PERMISSION_NONE        Permission = 1 // No permissions granted, used for users that are in the
	// project but have no default permissions
	Permission_PERMISSION_READ   Permission = 2 // Read only
	Permission_PERMISSION_APPEND Permission = 3 // Append objects to the collection cannot modify existing objects
	Permission_PERMISSION_MODIFY Permission = 4 // Can Read/Append/Modify objects in the collection
	// that owns the object / Create new collections
	Permission_PERMISSION_ADMIN Permission = 5 // Can modify the collections itself and permanently
)

// Enum value maps for Permission.
var (
	Permission_name = map[int32]string{
		0: "PERMISSION_UNSPECIFIED",
		1: "PERMISSION_NONE",
		2: "PERMISSION_READ",
		3: "PERMISSION_APPEND",
		4: "PERMISSION_MODIFY",
		5: "PERMISSION_ADMIN",
	}
	Permission_value = map[string]int32{
		"PERMISSION_UNSPECIFIED": 0,
		"PERMISSION_NONE":        1,
		"PERMISSION_READ":        2,
		"PERMISSION_APPEND":      3,
		"PERMISSION_MODIFY":      4,
		"PERMISSION_ADMIN":       5,
	}
)

func (x Permission) Enum() *Permission {
	p := new(Permission)
	*p = x
	return p
}

func (x Permission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission) Descriptor() protoreflect.EnumDescriptor {
	return file_aruna_api_storage_models_v1_auth_proto_enumTypes[0].Descriptor()
}

func (Permission) Type() protoreflect.EnumType {
	return &file_aruna_api_storage_models_v1_auth_proto_enumTypes[0]
}

func (x Permission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission.Descriptor instead.
func (Permission) EnumDescriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{0}
}

type PermType int32

const (
	PermType_PERM_TYPE_UNSPECIFIED PermType = 0
	PermType_PERM_TYPE_USER        PermType = 1 // Regular OAuth users
	PermType_PERM_TYPE_ANONYMOUS   PermType = 2 // Anonymous users without an OAuth token
	PermType_PERM_TYPE_TOKEN       PermType = 3 // Access token on behalf of a user
)

// Enum value maps for PermType.
var (
	PermType_name = map[int32]string{
		0: "PERM_TYPE_UNSPECIFIED",
		1: "PERM_TYPE_USER",
		2: "PERM_TYPE_ANONYMOUS",
		3: "PERM_TYPE_TOKEN",
	}
	PermType_value = map[string]int32{
		"PERM_TYPE_UNSPECIFIED": 0,
		"PERM_TYPE_USER":        1,
		"PERM_TYPE_ANONYMOUS":   2,
		"PERM_TYPE_TOKEN":       3,
	}
)

func (x PermType) Enum() *PermType {
	p := new(PermType)
	*p = x
	return p
}

func (x PermType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermType) Descriptor() protoreflect.EnumDescriptor {
	return file_aruna_api_storage_models_v1_auth_proto_enumTypes[1].Descriptor()
}

func (PermType) Type() protoreflect.EnumType {
	return &file_aruna_api_storage_models_v1_auth_proto_enumTypes[1]
}

func (x PermType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermType.Descriptor instead.
func (PermType) EnumDescriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{1}
}

type TokenType int32

const (
	TokenType_TOKEN_TYPE_UNSPECIFIED TokenType = 0
	TokenType_TOKEN_TYPE_PERSONAL    TokenType = 1
	TokenType_TOKEN_TYPE_SCOPED      TokenType = 2
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "TOKEN_TYPE_UNSPECIFIED",
		1: "TOKEN_TYPE_PERSONAL",
		2: "TOKEN_TYPE_SCOPED",
	}
	TokenType_value = map[string]int32{
		"TOKEN_TYPE_UNSPECIFIED": 0,
		"TOKEN_TYPE_PERSONAL":    1,
		"TOKEN_TYPE_SCOPED":      2,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_aruna_api_storage_models_v1_auth_proto_enumTypes[2].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_aruna_api_storage_models_v1_auth_proto_enumTypes[2]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{2}
}

// A Project is a list of collections with associated users
// This is used to manage access to multiple collections at the same time
// Each Collection can only be in one Project at a time
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UserPermissions []*ProjectPermission `protobuf:"bytes,3,rep,name=user_permissions,json=userPermissions,proto3" json:"user_permissions,omitempty"`
	CollectionIds   []string             `protobuf:"bytes,4,rep,name=collection_ids,json=collectionIds,proto3" json:"collection_ids,omitempty"`
	Description     string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetUserPermissions() []*ProjectPermission {
	if x != nil {
		return x.UserPermissions
	}
	return nil
}

func (x *Project) GetCollectionIds() []string {
	if x != nil {
		return x.CollectionIds
	}
	return nil
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ProjectOverview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CollectionIds []string `protobuf:"bytes,4,rep,name=collection_ids,json=collectionIds,proto3" json:"collection_ids,omitempty"`
	UserIds       []string `protobuf:"bytes,5,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *ProjectOverview) Reset() {
	*x = ProjectOverview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectOverview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectOverview) ProtoMessage() {}

func (x *ProjectOverview) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectOverview.ProtoReflect.Descriptor instead.
func (*ProjectOverview) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ProjectOverview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProjectOverview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProjectOverview) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProjectOverview) GetCollectionIds() []string {
	if x != nil {
		return x.CollectionIds
	}
	return nil
}

func (x *ProjectOverview) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal Aruna UserID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Oidc subject ID
	ExternalId string `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// (optional) User display_name
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Is the user activated
	Active bool `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	// Is the user admin ?
	IsAdmin bool `protobuf:"varint,5,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	// Is service account
	IsServiceAccount bool `protobuf:"varint,6,opt,name=is_service_account,json=isServiceAccount,proto3" json:"is_service_account,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *User) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *User) GetIsServiceAccount() bool {
	if x != nil {
		return x.IsServiceAccount
	}
	return false
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TokenType    TokenType              `protobuf:"varint,4,opt,name=token_type,json=tokenType,proto3,enum=aruna.api.storage.models.v1.TokenType" json:"token_type,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	CollectionId string                 `protobuf:"bytes,7,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	ProjectId    string                 `protobuf:"bytes,8,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Permission   Permission             `protobuf:"varint,9,opt,name=permission,proto3,enum=aruna.api.storage.models.v1.Permission" json:"permission,omitempty"`
	IsSession    bool                   `protobuf:"varint,10,opt,name=is_session,json=isSession,proto3" json:"is_session,omitempty"`
	UsedAt       *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=used_at,json=usedAt,proto3" json:"used_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Token) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Token) GetTokenType() TokenType {
	if x != nil {
		return x.TokenType
	}
	return TokenType_TOKEN_TYPE_UNSPECIFIED
}

func (x *Token) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Token) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Token) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Token) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Token) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_PERMISSION_UNSPECIFIED
}

func (x *Token) GetIsSession() bool {
	if x != nil {
		return x.IsSession
	}
	return false
}

func (x *Token) GetUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UsedAt
	}
	return nil
}

type ProjectPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProjectId      string     `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Permission     Permission `protobuf:"varint,3,opt,name=permission,proto3,enum=aruna.api.storage.models.v1.Permission" json:"permission,omitempty"`
	ServiceAccount bool       `protobuf:"varint,4,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
}

func (x *ProjectPermission) Reset() {
	*x = ProjectPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectPermission) ProtoMessage() {}

func (x *ProjectPermission) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectPermission.ProtoReflect.Descriptor instead.
func (*ProjectPermission) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *ProjectPermission) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProjectPermission) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectPermission) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_PERMISSION_UNSPECIFIED
}

func (x *ProjectPermission) GetServiceAccount() bool {
	if x != nil {
		return x.ServiceAccount
	}
	return false
}

type ProjectPermissionDisplayName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string     `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProjectId   string     `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Permission  Permission `protobuf:"varint,3,opt,name=permission,proto3,enum=aruna.api.storage.models.v1.Permission" json:"permission,omitempty"`
	DisplayName string     `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *ProjectPermissionDisplayName) Reset() {
	*x = ProjectPermissionDisplayName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectPermissionDisplayName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectPermissionDisplayName) ProtoMessage() {}

func (x *ProjectPermissionDisplayName) ProtoReflect() protoreflect.Message {
	mi := &file_aruna_api_storage_models_v1_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProjectPermissionDisplayName.ProtoReflect.Descriptor instead.
func (*ProjectPermissionDisplayName) Descriptor() ([]byte, []int) {
	return file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ProjectPermissionDisplayName) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ProjectPermissionDisplayName) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *ProjectPermissionDisplayName) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_PERMISSION_UNSPECIFIED
}

func (x *ProjectPermissionDisplayName) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_aruna_api_storage_models_v1_auth_proto protoreflect.FileDescriptor

var file_aruna_api_storage_models_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x99, 0x01, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc9, 0x03, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x73, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x75, 0x73, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xbd, 0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x47,
	0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xc2, 0x01, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x47, 0x0a, 0x0a, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e,
	0x61, 0x72, 0x75, 0x6e, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x96, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10,
	0x03, 0x12, 0x15, 0x0a, 0x11, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x45, 0x52, 0x4d,
	0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x05, 0x2a, 0x67,
	0x0a, 0x08, 0x50, 0x65, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x45,
	0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x45, 0x52, 0x4d, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x50, 0x45, 0x52,
	0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x4f, 0x4e, 0x59, 0x4d, 0x4f, 0x55, 0x53,
	0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x45, 0x52, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x03, 0x2a, 0x57, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x17, 0x0a, 0x13, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50,
	0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41,
	0x72, 0x75, 0x6e, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x72, 0x75, 0x6e, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aruna_api_storage_models_v1_auth_proto_rawDescOnce sync.Once
	file_aruna_api_storage_models_v1_auth_proto_rawDescData = file_aruna_api_storage_models_v1_auth_proto_rawDesc
)

func file_aruna_api_storage_models_v1_auth_proto_rawDescGZIP() []byte {
	file_aruna_api_storage_models_v1_auth_proto_rawDescOnce.Do(func() {
		file_aruna_api_storage_models_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_aruna_api_storage_models_v1_auth_proto_rawDescData)
	})
	return file_aruna_api_storage_models_v1_auth_proto_rawDescData
}

var file_aruna_api_storage_models_v1_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_aruna_api_storage_models_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_aruna_api_storage_models_v1_auth_proto_goTypes = []interface{}{
	(Permission)(0),                      // 0: aruna.api.storage.models.v1.Permission
	(PermType)(0),                        // 1: aruna.api.storage.models.v1.PermType
	(TokenType)(0),                       // 2: aruna.api.storage.models.v1.TokenType
	(*Project)(nil),                      // 3: aruna.api.storage.models.v1.Project
	(*ProjectOverview)(nil),              // 4: aruna.api.storage.models.v1.ProjectOverview
	(*User)(nil),                         // 5: aruna.api.storage.models.v1.User
	(*Token)(nil),                        // 6: aruna.api.storage.models.v1.Token
	(*ProjectPermission)(nil),            // 7: aruna.api.storage.models.v1.ProjectPermission
	(*ProjectPermissionDisplayName)(nil), // 8: aruna.api.storage.models.v1.ProjectPermissionDisplayName
	(*timestamppb.Timestamp)(nil),        // 9: google.protobuf.Timestamp
}
var file_aruna_api_storage_models_v1_auth_proto_depIdxs = []int32{
	7, // 0: aruna.api.storage.models.v1.Project.user_permissions:type_name -> aruna.api.storage.models.v1.ProjectPermission
	2, // 1: aruna.api.storage.models.v1.Token.token_type:type_name -> aruna.api.storage.models.v1.TokenType
	9, // 2: aruna.api.storage.models.v1.Token.created_at:type_name -> google.protobuf.Timestamp
	9, // 3: aruna.api.storage.models.v1.Token.expires_at:type_name -> google.protobuf.Timestamp
	0, // 4: aruna.api.storage.models.v1.Token.permission:type_name -> aruna.api.storage.models.v1.Permission
	9, // 5: aruna.api.storage.models.v1.Token.used_at:type_name -> google.protobuf.Timestamp
	0, // 6: aruna.api.storage.models.v1.ProjectPermission.permission:type_name -> aruna.api.storage.models.v1.Permission
	0, // 7: aruna.api.storage.models.v1.ProjectPermissionDisplayName.permission:type_name -> aruna.api.storage.models.v1.Permission
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_aruna_api_storage_models_v1_auth_proto_init() }
func file_aruna_api_storage_models_v1_auth_proto_init() {
	if File_aruna_api_storage_models_v1_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectOverview); i {
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
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectPermission); i {
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
		file_aruna_api_storage_models_v1_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectPermissionDisplayName); i {
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
			RawDescriptor: file_aruna_api_storage_models_v1_auth_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aruna_api_storage_models_v1_auth_proto_goTypes,
		DependencyIndexes: file_aruna_api_storage_models_v1_auth_proto_depIdxs,
		EnumInfos:         file_aruna_api_storage_models_v1_auth_proto_enumTypes,
		MessageInfos:      file_aruna_api_storage_models_v1_auth_proto_msgTypes,
	}.Build()
	File_aruna_api_storage_models_v1_auth_proto = out.File
	file_aruna_api_storage_models_v1_auth_proto_rawDesc = nil
	file_aruna_api_storage_models_v1_auth_proto_goTypes = nil
	file_aruna_api_storage_models_v1_auth_proto_depIdxs = nil
}
