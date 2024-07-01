// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: manager_event.proto

package manager_service

import (
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

type EventID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EventID) Reset() {
	*x = EventID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventID) ProtoMessage() {}

func (x *EventID) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventID.ProtoReflect.Descriptor instead.
func (*EventID) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StudentName string `protobuf:"bytes,1,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
	Topic       string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Day         string `protobuf:"bytes,3,opt,name=day,proto3" json:"day,omitempty"`
	Branch      string `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	StartTime   string `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventRequest) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *CreateEventRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *CreateEventRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *CreateEventRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *CreateEventRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentName string `protobuf:"bytes,2,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
	Topic       string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Day         string `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`
	Branch      string `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	StartTime   string `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Event) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *Event) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Event) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *Event) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *Event) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

type UpdateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentName string `protobuf:"bytes,2,opt,name=student_name,json=studentName,proto3" json:"student_name,omitempty"`
	Topic       string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Day         string `protobuf:"bytes,4,opt,name=day,proto3" json:"day,omitempty"`
	Branch      string `protobuf:"bytes,5,opt,name=branch,proto3" json:"branch,omitempty"`
	StartTime   string `protobuf:"bytes,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
}

func (x *UpdateEventRequest) Reset() {
	*x = UpdateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEventRequest) ProtoMessage() {}

func (x *UpdateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEventRequest.ProtoReflect.Descriptor instead.
func (*UpdateEventRequest) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEventRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateEventRequest) GetStudentName() string {
	if x != nil {
		return x.StudentName
	}
	return ""
}

func (x *UpdateEventRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *UpdateEventRequest) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

func (x *UpdateEventRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *UpdateEventRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

type GetAllEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   uint64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetAllEventsRequest) Reset() {
	*x = GetAllEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllEventsRequest) ProtoMessage() {}

func (x *GetAllEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllEventsRequest.ProtoReflect.Descriptor instead.
func (*GetAllEventsRequest) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllEventsRequest) GetPage() uint64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllEventsRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllEventsRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetAllEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Events []*Event `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *GetAllEventsResponse) Reset() {
	*x = GetAllEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllEventsResponse) ProtoMessage() {}

func (x *GetAllEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_manager_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllEventsResponse.ProtoReflect.Descriptor instead.
func (*GetAllEventsResponse) Descriptor() ([]byte, []int) {
	return file_manager_event_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllEventsResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetAllEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

var File_manager_event_proto protoreflect.FileDescriptor

var file_manager_event_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x07, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a,
	0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x57, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x99, 0x03, 0x0a, 0x13, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c,
	0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x5d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x24, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x18, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_manager_event_proto_rawDescOnce sync.Once
	file_manager_event_proto_rawDescData = file_manager_event_proto_rawDesc
)

func file_manager_event_proto_rawDescGZIP() []byte {
	file_manager_event_proto_rawDescOnce.Do(func() {
		file_manager_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_event_proto_rawDescData)
	})
	return file_manager_event_proto_rawDescData
}

var file_manager_event_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_manager_event_proto_goTypes = []interface{}{
	(*EventID)(nil),              // 0: manager_service.EventID
	(*CreateEventRequest)(nil),   // 1: manager_service.CreateEventRequest
	(*Event)(nil),                // 2: manager_service.Event
	(*UpdateEventRequest)(nil),   // 3: manager_service.UpdateEventRequest
	(*GetAllEventsRequest)(nil),  // 4: manager_service.GetAllEventsRequest
	(*GetAllEventsResponse)(nil), // 5: manager_service.GetAllEventsResponse
}
var file_manager_event_proto_depIdxs = []int32{
	2, // 0: manager_service.GetAllEventsResponse.events:type_name -> manager_service.Event
	1, // 1: manager_service.ManagerEventService.CreateEvent:input_type -> manager_service.CreateEventRequest
	0, // 2: manager_service.ManagerEventService.GetEventByID:input_type -> manager_service.EventID
	4, // 3: manager_service.ManagerEventService.GetAllEvents:input_type -> manager_service.GetAllEventsRequest
	3, // 4: manager_service.ManagerEventService.UpdateEvent:input_type -> manager_service.UpdateEventRequest
	0, // 5: manager_service.ManagerEventService.DeleteEvent:input_type -> manager_service.EventID
	2, // 6: manager_service.ManagerEventService.CreateEvent:output_type -> manager_service.Event
	2, // 7: manager_service.ManagerEventService.GetEventByID:output_type -> manager_service.Event
	5, // 8: manager_service.ManagerEventService.GetAllEvents:output_type -> manager_service.GetAllEventsResponse
	2, // 9: manager_service.ManagerEventService.UpdateEvent:output_type -> manager_service.Event
	0, // 10: manager_service.ManagerEventService.DeleteEvent:output_type -> manager_service.EventID
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manager_event_proto_init() }
func file_manager_event_proto_init() {
	if File_manager_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventID); i {
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
		file_manager_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_manager_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_manager_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEventRequest); i {
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
		file_manager_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllEventsRequest); i {
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
		file_manager_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllEventsResponse); i {
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
			RawDescriptor: file_manager_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_manager_event_proto_goTypes,
		DependencyIndexes: file_manager_event_proto_depIdxs,
		MessageInfos:      file_manager_event_proto_msgTypes,
	}.Build()
	File_manager_event_proto = out.File
	file_manager_event_proto_rawDesc = nil
	file_manager_event_proto_goTypes = nil
	file_manager_event_proto_depIdxs = nil
}
