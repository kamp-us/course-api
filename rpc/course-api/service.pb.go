// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: rpc/course-api/service.proto

package course_api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Course Start
type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_course_api_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_course_api_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_rpc_course_api_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCourseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type CreateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CategoryIds []string `protobuf:"bytes,4,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
}

func (x *CreateCourseRequest) Reset() {
	*x = CreateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_course_api_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseRequest) ProtoMessage() {}

func (x *CreateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_course_api_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseRequest.ProtoReflect.Descriptor instead.
func (*CreateCourseRequest) Descriptor() ([]byte, []int) {
	return file_rpc_course_api_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCourseRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateCourseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCourseRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateCourseRequest) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

type UpdateCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string                    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        *wrapperspb.StringValue   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description *wrapperspb.StringValue   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CategoryIds []*wrapperspb.StringValue `protobuf:"bytes,4,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
}

func (x *UpdateCourseRequest) Reset() {
	*x = UpdateCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_course_api_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCourseRequest) ProtoMessage() {}

func (x *UpdateCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_course_api_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCourseRequest.ProtoReflect.Descriptor instead.
func (*UpdateCourseRequest) Descriptor() ([]byte, []int) {
	return file_rpc_course_api_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCourseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *UpdateCourseRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *UpdateCourseRequest) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *UpdateCourseRequest) GetCategoryIds() []*wrapperspb.StringValue {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

type DeleteCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCourseRequest) Reset() {
	*x = DeleteCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_course_api_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCourseRequest) ProtoMessage() {}

func (x *DeleteCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_course_api_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCourseRequest.ProtoReflect.Descriptor instead.
func (*DeleteCourseRequest) Descriptor() ([]byte, []int) {
	return file_rpc_course_api_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCourseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name        string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CategoryIds []string `protobuf:"bytes,5,rep,name=category_ids,json=categoryIds,proto3" json:"category_ids,omitempty"`
	Slug        string   `protobuf:"bytes,6,opt,name=slug,proto3" json:"slug,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_course_api_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_course_api_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_rpc_course_api_service_proto_rawDescGZIP(), []int{4}
}

func (x *Course) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Course) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetCategoryIds() []string {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *Course) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

var File_rpc_course_api_service_proto protoreflect.FileDescriptor

var file_rpc_course_api_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x22, 0xd8, 0x01, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x9e, 0x01,
	0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75, 0x67, 0x32, 0xc7,
	0x02, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x41, 0x50, 0x49, 0x12, 0x49, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x22, 0x2e, 0x6b, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x6b, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x25, 0x2e, 0x6b, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_rpc_course_api_service_proto_rawDescOnce sync.Once
	file_rpc_course_api_service_proto_rawDescData = file_rpc_course_api_service_proto_rawDesc
)

func file_rpc_course_api_service_proto_rawDescGZIP() []byte {
	file_rpc_course_api_service_proto_rawDescOnce.Do(func() {
		file_rpc_course_api_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_course_api_service_proto_rawDescData)
	})
	return file_rpc_course_api_service_proto_rawDescData
}

var file_rpc_course_api_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_rpc_course_api_service_proto_goTypes = []interface{}{
	(*GetCourseRequest)(nil),       // 0: kampus.courseapi.GetCourseRequest
	(*CreateCourseRequest)(nil),    // 1: kampus.courseapi.CreateCourseRequest
	(*UpdateCourseRequest)(nil),    // 2: kampus.courseapi.UpdateCourseRequest
	(*DeleteCourseRequest)(nil),    // 3: kampus.courseapi.DeleteCourseRequest
	(*Course)(nil),                 // 4: kampus.courseapi.Course
	(*wrapperspb.StringValue)(nil), // 5: google.protobuf.StringValue
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_rpc_course_api_service_proto_depIdxs = []int32{
	5, // 0: kampus.courseapi.UpdateCourseRequest.name:type_name -> google.protobuf.StringValue
	5, // 1: kampus.courseapi.UpdateCourseRequest.description:type_name -> google.protobuf.StringValue
	5, // 2: kampus.courseapi.UpdateCourseRequest.category_ids:type_name -> google.protobuf.StringValue
	0, // 3: kampus.courseapi.CourseAPI.GetCourse:input_type -> kampus.courseapi.GetCourseRequest
	1, // 4: kampus.courseapi.CourseAPI.CreateCourse:input_type -> kampus.courseapi.CreateCourseRequest
	2, // 5: kampus.courseapi.CourseAPI.UpdateCourse:input_type -> kampus.courseapi.UpdateCourseRequest
	3, // 6: kampus.courseapi.CourseAPI.DeleteCourse:input_type -> kampus.courseapi.DeleteCourseRequest
	4, // 7: kampus.courseapi.CourseAPI.GetCourse:output_type -> kampus.courseapi.Course
	4, // 8: kampus.courseapi.CourseAPI.CreateCourse:output_type -> kampus.courseapi.Course
	4, // 9: kampus.courseapi.CourseAPI.UpdateCourse:output_type -> kampus.courseapi.Course
	6, // 10: kampus.courseapi.CourseAPI.DeleteCourse:output_type -> google.protobuf.Empty
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_course_api_service_proto_init() }
func file_rpc_course_api_service_proto_init() {
	if File_rpc_course_api_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_course_api_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseRequest); i {
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
		file_rpc_course_api_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseRequest); i {
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
		file_rpc_course_api_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCourseRequest); i {
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
		file_rpc_course_api_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCourseRequest); i {
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
		file_rpc_course_api_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
			RawDescriptor: file_rpc_course_api_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_course_api_service_proto_goTypes,
		DependencyIndexes: file_rpc_course_api_service_proto_depIdxs,
		MessageInfos:      file_rpc_course_api_service_proto_msgTypes,
	}.Build()
	File_rpc_course_api_service_proto = out.File
	file_rpc_course_api_service_proto_rawDesc = nil
	file_rpc_course_api_service_proto_goTypes = nil
	file_rpc_course_api_service_proto_depIdxs = nil
}
