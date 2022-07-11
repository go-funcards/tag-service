// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/tag.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId   string `protobuf:"bytes,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	OwnerId string `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	BoardId string `protobuf:"bytes,3,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Color   string `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *CreateTagRequest) Reset() {
	*x = CreateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTagRequest) ProtoMessage() {}

func (x *CreateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTagRequest.ProtoReflect.Descriptor instead.
func (*CreateTagRequest) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTagRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *CreateTagRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *CreateTagRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *CreateTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTagRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type UpdateTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId   string `protobuf:"bytes,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	BoardId string `protobuf:"bytes,2,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Color   string `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *UpdateTagRequest) Reset() {
	*x = UpdateTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTagRequest) ProtoMessage() {}

func (x *UpdateTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTagRequest.ProtoReflect.Descriptor instead.
func (*UpdateTagRequest) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateTagRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *UpdateTagRequest) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *UpdateTagRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTagRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type DeleteTagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId string `protobuf:"bytes,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
}

func (x *DeleteTagRequest) Reset() {
	*x = DeleteTagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTagRequest) ProtoMessage() {}

func (x *DeleteTagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTagRequest.ProtoReflect.Descriptor instead.
func (*DeleteTagRequest) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTagRequest) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type TagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageIndex uint64   `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	PageSize  uint32   `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	TagIds    []string `protobuf:"bytes,3,rep,name=tag_ids,json=tagIds,proto3" json:"tag_ids,omitempty"`
	OwnerIds  []string `protobuf:"bytes,4,rep,name=owner_ids,json=ownerIds,proto3" json:"owner_ids,omitempty"`
	BoardIds  []string `protobuf:"bytes,5,rep,name=board_ids,json=boardIds,proto3" json:"board_ids,omitempty"`
}

func (x *TagsRequest) Reset() {
	*x = TagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsRequest) ProtoMessage() {}

func (x *TagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsRequest.ProtoReflect.Descriptor instead.
func (*TagsRequest) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{3}
}

func (x *TagsRequest) GetPageIndex() uint64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *TagsRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *TagsRequest) GetTagIds() []string {
	if x != nil {
		return x.TagIds
	}
	return nil
}

func (x *TagsRequest) GetOwnerIds() []string {
	if x != nil {
		return x.OwnerIds
	}
	return nil
}

func (x *TagsRequest) GetBoardIds() []string {
	if x != nil {
		return x.BoardIds
	}
	return nil
}

type TagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total uint64              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Tags  []*TagsResponse_Tag `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TagsResponse) Reset() {
	*x = TagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsResponse) ProtoMessage() {}

func (x *TagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsResponse.ProtoReflect.Descriptor instead.
func (*TagsResponse) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{4}
}

func (x *TagsResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *TagsResponse) GetTags() []*TagsResponse_Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

type TagsResponse_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TagId     string                 `protobuf:"bytes,1,opt,name=tag_id,json=tagId,proto3" json:"tag_id,omitempty"`
	OwnerId   string                 `protobuf:"bytes,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	BoardId   string                 `protobuf:"bytes,3,opt,name=board_id,json=boardId,proto3" json:"board_id,omitempty"`
	Name      string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Color     string                 `protobuf:"bytes,5,opt,name=color,proto3" json:"color,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TagsResponse_Tag) Reset() {
	*x = TagsResponse_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagsResponse_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagsResponse_Tag) ProtoMessage() {}

func (x *TagsResponse_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_v1_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagsResponse_Tag.ProtoReflect.Descriptor instead.
func (*TagsResponse_Tag) Descriptor() ([]byte, []int) {
	return file_v1_tag_proto_rawDescGZIP(), []int{4, 0}
}

func (x *TagsResponse_Tag) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

func (x *TagsResponse_Tag) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *TagsResponse_Tag) GetBoardId() string {
	if x != nil {
		return x.BoardId
	}
	return ""
}

func (x *TagsResponse_Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TagsResponse_Tag) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *TagsResponse_Tag) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_v1_tag_proto protoreflect.FileDescriptor

var file_v1_tag_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x6e, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x22, 0x29, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67, 0x49, 0x64, 0x22, 0x9c, 0x01,
	0x0a, 0x0b, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x67,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x67, 0x49,
	0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x73, 0x22, 0x8e, 0x02, 0x0a,
	0x0c, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x1a, 0xb7, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x15, 0x0a, 0x06, 0x74,
	0x61, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x82, 0x02,
	0x0a, 0x03, 0x54, 0x61, 0x67, 0x12, 0x3f, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x44, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x42,
	0x08, 0x54, 0x61, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x76, 0x31, 0xaa, 0x02, 0x13, 0x46, 0x75, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x73, 0x4f, 0x72, 0x67,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_tag_proto_rawDescOnce sync.Once
	file_v1_tag_proto_rawDescData = file_v1_tag_proto_rawDesc
)

func file_v1_tag_proto_rawDescGZIP() []byte {
	file_v1_tag_proto_rawDescOnce.Do(func() {
		file_v1_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_tag_proto_rawDescData)
	})
	return file_v1_tag_proto_rawDescData
}

var file_v1_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_tag_proto_goTypes = []interface{}{
	(*CreateTagRequest)(nil),      // 0: proto.v1.CreateTagRequest
	(*UpdateTagRequest)(nil),      // 1: proto.v1.UpdateTagRequest
	(*DeleteTagRequest)(nil),      // 2: proto.v1.DeleteTagRequest
	(*TagsRequest)(nil),           // 3: proto.v1.TagsRequest
	(*TagsResponse)(nil),          // 4: proto.v1.TagsResponse
	(*TagsResponse_Tag)(nil),      // 5: proto.v1.TagsResponse.Tag
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
}
var file_v1_tag_proto_depIdxs = []int32{
	5, // 0: proto.v1.TagsResponse.tags:type_name -> proto.v1.TagsResponse.Tag
	6, // 1: proto.v1.TagsResponse.Tag.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: proto.v1.Tag.CreateTag:input_type -> proto.v1.CreateTagRequest
	1, // 3: proto.v1.Tag.UpdateTag:input_type -> proto.v1.UpdateTagRequest
	2, // 4: proto.v1.Tag.DeleteTag:input_type -> proto.v1.DeleteTagRequest
	3, // 5: proto.v1.Tag.GetTags:input_type -> proto.v1.TagsRequest
	7, // 6: proto.v1.Tag.CreateTag:output_type -> google.protobuf.Empty
	7, // 7: proto.v1.Tag.UpdateTag:output_type -> google.protobuf.Empty
	7, // 8: proto.v1.Tag.DeleteTag:output_type -> google.protobuf.Empty
	4, // 9: proto.v1.Tag.GetTags:output_type -> proto.v1.TagsResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_tag_proto_init() }
func file_v1_tag_proto_init() {
	if File_v1_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTagRequest); i {
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
		file_v1_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTagRequest); i {
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
		file_v1_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTagRequest); i {
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
		file_v1_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagsRequest); i {
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
		file_v1_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagsResponse); i {
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
		file_v1_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagsResponse_Tag); i {
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
			RawDescriptor: file_v1_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_tag_proto_goTypes,
		DependencyIndexes: file_v1_tag_proto_depIdxs,
		MessageInfos:      file_v1_tag_proto_msgTypes,
	}.Build()
	File_v1_tag_proto = out.File
	file_v1_tag_proto_rawDesc = nil
	file_v1_tag_proto_goTypes = nil
	file_v1_tag_proto_depIdxs = nil
}