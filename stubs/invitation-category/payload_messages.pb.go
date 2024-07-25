// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: invitation-category/payload_messages.proto

package invitation_category

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

type InvitationCategoryPaging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,proto3" json:"current_page,omitempty"`
	TotalPage   int32 `protobuf:"varint,2,opt,name=total_page,proto3" json:"total_page,omitempty"`
	Count       int32 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *InvitationCategoryPaging) Reset() {
	*x = InvitationCategoryPaging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationCategoryPaging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationCategoryPaging) ProtoMessage() {}

func (x *InvitationCategoryPaging) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationCategoryPaging.ProtoReflect.Descriptor instead.
func (*InvitationCategoryPaging) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{0}
}

func (x *InvitationCategoryPaging) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *InvitationCategoryPaging) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *InvitationCategoryPaging) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type InvitationCategoryBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                    `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    *anypb.Any                `protobuf:"bytes,2,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Paging  *InvitationCategoryPaging `protobuf:"bytes,3,opt,name=paging,proto3,oneof" json:"paging,omitempty"`
}

func (x *InvitationCategoryBaseResponse) Reset() {
	*x = InvitationCategoryBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvitationCategoryBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvitationCategoryBaseResponse) ProtoMessage() {}

func (x *InvitationCategoryBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvitationCategoryBaseResponse.ProtoReflect.Descriptor instead.
func (*InvitationCategoryBaseResponse) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{1}
}

func (x *InvitationCategoryBaseResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *InvitationCategoryBaseResponse) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InvitationCategoryBaseResponse) GetPaging() *InvitationCategoryPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type CreateInvitationCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateInvitationCategoryRequest) Reset() {
	*x = CreateInvitationCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvitationCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvitationCategoryRequest) ProtoMessage() {}

func (x *CreateInvitationCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateInvitationCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateInvitationCategoryRequest) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CreateInvitationCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAllInvitationCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search *string `protobuf:"bytes,1,opt,name=search,proto3,oneof" json:"search,omitempty"`
	Page   *int32  `protobuf:"varint,2,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Limit  *int32  `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
}

func (x *GetAllInvitationCategoryRequest) Reset() {
	*x = GetAllInvitationCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllInvitationCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllInvitationCategoryRequest) ProtoMessage() {}

func (x *GetAllInvitationCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllInvitationCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetAllInvitationCategoryRequest) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllInvitationCategoryRequest) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

func (x *GetAllInvitationCategoryRequest) GetPage() int32 {
	if x != nil && x.Page != nil {
		return *x.Page
	}
	return 0
}

func (x *GetAllInvitationCategoryRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type GetAllInvitationCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string                               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*GetInvitationCategoryByIDResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Paging  *InvitationCategoryPaging            `protobuf:"bytes,3,opt,name=paging,proto3" json:"paging,omitempty"`
}

func (x *GetAllInvitationCategoryResponse) Reset() {
	*x = GetAllInvitationCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllInvitationCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllInvitationCategoryResponse) ProtoMessage() {}

func (x *GetAllInvitationCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllInvitationCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllInvitationCategoryResponse) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllInvitationCategoryResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllInvitationCategoryResponse) GetData() []*GetInvitationCategoryByIDResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetAllInvitationCategoryResponse) GetPaging() *InvitationCategoryPaging {
	if x != nil {
		return x.Paging
	}
	return nil
}

type GetInvitationCategoryByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInvitationCategoryByIDRequest) Reset() {
	*x = GetInvitationCategoryByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitationCategoryByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitationCategoryByIDRequest) ProtoMessage() {}

func (x *GetInvitationCategoryByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitationCategoryByIDRequest.ProtoReflect.Descriptor instead.
func (*GetInvitationCategoryByIDRequest) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{5}
}

func (x *GetInvitationCategoryByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetInvitationCategoryByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,proto3" json:"deleted_at,omitempty"`
}

func (x *GetInvitationCategoryByIDResponse) Reset() {
	*x = GetInvitationCategoryByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInvitationCategoryByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInvitationCategoryByIDResponse) ProtoMessage() {}

func (x *GetInvitationCategoryByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInvitationCategoryByIDResponse.ProtoReflect.Descriptor instead.
func (*GetInvitationCategoryByIDResponse) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{6}
}

func (x *GetInvitationCategoryByIDResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetInvitationCategoryByIDResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetInvitationCategoryByIDResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetInvitationCategoryByIDResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *GetInvitationCategoryByIDResponse) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type UpdateInvitationCategoryByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateInvitationCategoryByIDRequest) Reset() {
	*x = UpdateInvitationCategoryByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInvitationCategoryByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInvitationCategoryByIDRequest) ProtoMessage() {}

func (x *UpdateInvitationCategoryByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInvitationCategoryByIDRequest.ProtoReflect.Descriptor instead.
func (*UpdateInvitationCategoryByIDRequest) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateInvitationCategoryByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateInvitationCategoryByIDRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteInvitationCategoryByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInvitationCategoryByIDRequest) Reset() {
	*x = DeleteInvitationCategoryByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invitation_category_payload_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInvitationCategoryByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInvitationCategoryByIDRequest) ProtoMessage() {}

func (x *DeleteInvitationCategoryByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invitation_category_payload_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInvitationCategoryByIDRequest.ProtoReflect.Descriptor instead.
func (*DeleteInvitationCategoryByIDRequest) Descriptor() ([]byte, []int) {
	return file_invitation_category_payload_messages_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteInvitationCategoryByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_invitation_category_payload_messages_proto protoreflect.FileDescriptor

var file_invitation_category_payload_messages_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x74, 0x0a, 0x18, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x1e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x67, 0x48, 0x01, 0x52, 0x06, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x22, 0x35, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xb3, 0x01,
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x06, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x67, 0x22, 0x32, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xfb, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3a, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x49, 0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x35, 0x0a, 0x23, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x67, 0x69, 0x73, 0x61, 0x74, 0x61, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x75, 0x62,
	0x73, 0x2f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invitation_category_payload_messages_proto_rawDescOnce sync.Once
	file_invitation_category_payload_messages_proto_rawDescData = file_invitation_category_payload_messages_proto_rawDesc
)

func file_invitation_category_payload_messages_proto_rawDescGZIP() []byte {
	file_invitation_category_payload_messages_proto_rawDescOnce.Do(func() {
		file_invitation_category_payload_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_invitation_category_payload_messages_proto_rawDescData)
	})
	return file_invitation_category_payload_messages_proto_rawDescData
}

var file_invitation_category_payload_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_invitation_category_payload_messages_proto_goTypes = []any{
	(*InvitationCategoryPaging)(nil),            // 0: proto.InvitationCategoryPaging
	(*InvitationCategoryBaseResponse)(nil),      // 1: proto.InvitationCategoryBaseResponse
	(*CreateInvitationCategoryRequest)(nil),     // 2: proto.CreateInvitationCategoryRequest
	(*GetAllInvitationCategoryRequest)(nil),     // 3: proto.GetAllInvitationCategoryRequest
	(*GetAllInvitationCategoryResponse)(nil),    // 4: proto.GetAllInvitationCategoryResponse
	(*GetInvitationCategoryByIDRequest)(nil),    // 5: proto.GetInvitationCategoryByIDRequest
	(*GetInvitationCategoryByIDResponse)(nil),   // 6: proto.GetInvitationCategoryByIDResponse
	(*UpdateInvitationCategoryByIDRequest)(nil), // 7: proto.UpdateInvitationCategoryByIDRequest
	(*DeleteInvitationCategoryByIDRequest)(nil), // 8: proto.DeleteInvitationCategoryByIDRequest
	(*anypb.Any)(nil),                           // 9: google.protobuf.Any
	(*timestamppb.Timestamp)(nil),               // 10: google.protobuf.Timestamp
}
var file_invitation_category_payload_messages_proto_depIdxs = []int32{
	9,  // 0: proto.InvitationCategoryBaseResponse.data:type_name -> google.protobuf.Any
	0,  // 1: proto.InvitationCategoryBaseResponse.paging:type_name -> proto.InvitationCategoryPaging
	6,  // 2: proto.GetAllInvitationCategoryResponse.data:type_name -> proto.GetInvitationCategoryByIDResponse
	0,  // 3: proto.GetAllInvitationCategoryResponse.paging:type_name -> proto.InvitationCategoryPaging
	10, // 4: proto.GetInvitationCategoryByIDResponse.created_at:type_name -> google.protobuf.Timestamp
	10, // 5: proto.GetInvitationCategoryByIDResponse.updated_at:type_name -> google.protobuf.Timestamp
	10, // 6: proto.GetInvitationCategoryByIDResponse.deleted_at:type_name -> google.protobuf.Timestamp
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_invitation_category_payload_messages_proto_init() }
func file_invitation_category_payload_messages_proto_init() {
	if File_invitation_category_payload_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invitation_category_payload_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InvitationCategoryPaging); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InvitationCategoryBaseResponse); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateInvitationCategoryRequest); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllInvitationCategoryRequest); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllInvitationCategoryResponse); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetInvitationCategoryByIDRequest); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetInvitationCategoryByIDResponse); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateInvitationCategoryByIDRequest); i {
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
		file_invitation_category_payload_messages_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteInvitationCategoryByIDRequest); i {
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
	file_invitation_category_payload_messages_proto_msgTypes[1].OneofWrappers = []any{}
	file_invitation_category_payload_messages_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_invitation_category_payload_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_invitation_category_payload_messages_proto_goTypes,
		DependencyIndexes: file_invitation_category_payload_messages_proto_depIdxs,
		MessageInfos:      file_invitation_category_payload_messages_proto_msgTypes,
	}.Build()
	File_invitation_category_payload_messages_proto = out.File
	file_invitation_category_payload_messages_proto_rawDesc = nil
	file_invitation_category_payload_messages_proto_goTypes = nil
	file_invitation_category_payload_messages_proto_depIdxs = nil
}
