// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.12.2
// source: infra/grpc/task.proto

package grpc_task

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

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Summary   string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	OwnerId   int32  `protobuf:"varint,3,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	When      string `protobuf:"bytes,5,opt,name=when,proto3" json:"when,omitempty"`
	FirstName string `protobuf:"bytes,6,opt,name=firstName,json=first_name,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,7,opt,name=lastName,json=last_name,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_grpc_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_infra_grpc_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_infra_grpc_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *Task) GetOwnerId() int32 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Task) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

func (x *Task) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Task) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Task) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type TaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Summary        string `protobuf:"bytes,2,opt,name=summary,proto3" json:"summary,omitempty"`
	OwnerTaskId    int32  `protobuf:"varint,3,opt,name=ownerTaskId,proto3" json:"ownerTaskId,omitempty"`
	UserLoggedId   int32  `protobuf:"varint,4,opt,name=userLoggedId,proto3" json:"userLoggedId,omitempty"`
	UserLoggedRole string `protobuf:"bytes,5,opt,name=userLoggedRole,proto3" json:"userLoggedRole,omitempty"`
	Status         string `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	When           string `protobuf:"bytes,7,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *TaskRequest) Reset() {
	*x = TaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_grpc_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRequest) ProtoMessage() {}

func (x *TaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_grpc_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRequest.ProtoReflect.Descriptor instead.
func (*TaskRequest) Descriptor() ([]byte, []int) {
	return file_infra_grpc_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TaskRequest) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *TaskRequest) GetOwnerTaskId() int32 {
	if x != nil {
		return x.OwnerTaskId
	}
	return 0
}

func (x *TaskRequest) GetUserLoggedId() int32 {
	if x != nil {
		return x.UserLoggedId
	}
	return 0
}

func (x *TaskRequest) GetUserLoggedRole() string {
	if x != nil {
		return x.UserLoggedRole
	}
	return ""
}

func (x *TaskRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskRequest) GetWhen() string {
	if x != nil {
		return x.When
	}
	return ""
}

type TasksListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLoggedId   int32  `protobuf:"varint,1,opt,name=userLoggedId,proto3" json:"userLoggedId,omitempty"`
	UserLoggedRole string `protobuf:"bytes,2,opt,name=userLoggedRole,proto3" json:"userLoggedRole,omitempty"`
	OwnerTaskId    int32  `protobuf:"varint,3,opt,name=ownerTaskId,proto3" json:"ownerTaskId,omitempty"`
	Limit          int32  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *TasksListRequest) Reset() {
	*x = TasksListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_grpc_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TasksListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TasksListRequest) ProtoMessage() {}

func (x *TasksListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infra_grpc_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TasksListRequest.ProtoReflect.Descriptor instead.
func (*TasksListRequest) Descriptor() ([]byte, []int) {
	return file_infra_grpc_task_proto_rawDescGZIP(), []int{2}
}

func (x *TasksListRequest) GetUserLoggedId() int32 {
	if x != nil {
		return x.UserLoggedId
	}
	return 0
}

func (x *TasksListRequest) GetUserLoggedRole() string {
	if x != nil {
		return x.UserLoggedRole
	}
	return ""
}

func (x *TasksListRequest) GetOwnerTaskId() int32 {
	if x != nil {
		return x.OwnerTaskId
	}
	return 0
}

func (x *TasksListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type TaskList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *TaskList) Reset() {
	*x = TaskList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infra_grpc_task_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskList) ProtoMessage() {}

func (x *TaskList) ProtoReflect() protoreflect.Message {
	mi := &file_infra_grpc_task_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskList.ProtoReflect.Descriptor instead.
func (*TaskList) Descriptor() ([]byte, []int) {
	return file_infra_grpc_task_proto_rawDescGZIP(), []int{3}
}

func (x *TaskList) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_infra_grpc_task_proto protoreflect.FileDescriptor

var file_infra_grpc_task_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x22, 0xc8, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xd1, 0x01,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65,
	0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x77, 0x68, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x68, 0x65,
	0x6e, 0x22, 0x96, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x67, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x75, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x67, 0x65, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x31, 0x0a, 0x08, 0x54, 0x61,
	0x73, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0x95, 0x03,
	0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a,
	0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infra_grpc_task_proto_rawDescOnce sync.Once
	file_infra_grpc_task_proto_rawDescData = file_infra_grpc_task_proto_rawDesc
)

func file_infra_grpc_task_proto_rawDescGZIP() []byte {
	file_infra_grpc_task_proto_rawDescOnce.Do(func() {
		file_infra_grpc_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_infra_grpc_task_proto_rawDescData)
	})
	return file_infra_grpc_task_proto_rawDescData
}

var file_infra_grpc_task_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infra_grpc_task_proto_goTypes = []interface{}{
	(*Task)(nil),             // 0: grpc_task.Task
	(*TaskRequest)(nil),      // 1: grpc_task.TaskRequest
	(*TasksListRequest)(nil), // 2: grpc_task.TasksListRequest
	(*TaskList)(nil),         // 3: grpc_task.TaskList
}
var file_infra_grpc_task_proto_depIdxs = []int32{
	0, // 0: grpc_task.TaskList.tasks:type_name -> grpc_task.Task
	1, // 1: grpc_task.TaskService.CreateTaskRequest:input_type -> grpc_task.TaskRequest
	2, // 2: grpc_task.TaskService.ListTasksRequest:input_type -> grpc_task.TasksListRequest
	1, // 3: grpc_task.TaskService.FindOneTaskRequest:input_type -> grpc_task.TaskRequest
	1, // 4: grpc_task.TaskService.UpdateTaskRequest:input_type -> grpc_task.TaskRequest
	1, // 5: grpc_task.TaskService.DeleteTaskRequest:input_type -> grpc_task.TaskRequest
	1, // 6: grpc_task.TaskService.CloseTaskRequest:input_type -> grpc_task.TaskRequest
	0, // 7: grpc_task.TaskService.CreateTaskRequest:output_type -> grpc_task.Task
	3, // 8: grpc_task.TaskService.ListTasksRequest:output_type -> grpc_task.TaskList
	0, // 9: grpc_task.TaskService.FindOneTaskRequest:output_type -> grpc_task.Task
	0, // 10: grpc_task.TaskService.UpdateTaskRequest:output_type -> grpc_task.Task
	0, // 11: grpc_task.TaskService.DeleteTaskRequest:output_type -> grpc_task.Task
	0, // 12: grpc_task.TaskService.CloseTaskRequest:output_type -> grpc_task.Task
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_infra_grpc_task_proto_init() }
func file_infra_grpc_task_proto_init() {
	if File_infra_grpc_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infra_grpc_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_infra_grpc_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRequest); i {
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
		file_infra_grpc_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TasksListRequest); i {
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
		file_infra_grpc_task_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskList); i {
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
			RawDescriptor: file_infra_grpc_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infra_grpc_task_proto_goTypes,
		DependencyIndexes: file_infra_grpc_task_proto_depIdxs,
		MessageInfos:      file_infra_grpc_task_proto_msgTypes,
	}.Build()
	File_infra_grpc_task_proto = out.File
	file_infra_grpc_task_proto_rawDesc = nil
	file_infra_grpc_task_proto_goTypes = nil
	file_infra_grpc_task_proto_depIdxs = nil
}
