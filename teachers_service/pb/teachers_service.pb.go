// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: teachers_service.proto

package pb

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

type GetTeachersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTeachersRequest) Reset() {
	*x = GetTeachersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teachers_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTeachersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTeachersRequest) ProtoMessage() {}

func (x *GetTeachersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teachers_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTeachersRequest.ProtoReflect.Descriptor instead.
func (*GetTeachersRequest) Descriptor() ([]byte, []int) {
	return file_teachers_service_proto_rawDescGZIP(), []int{0}
}

type TeacherList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Teacher `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TeacherList) Reset() {
	*x = TeacherList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teachers_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherList) ProtoMessage() {}

func (x *TeacherList) ProtoReflect() protoreflect.Message {
	mi := &file_teachers_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherList.ProtoReflect.Descriptor instead.
func (*TeacherList) Descriptor() ([]byte, []int) {
	return file_teachers_service_proto_rawDescGZIP(), []int{1}
}

func (x *TeacherList) GetList() []*Teacher {
	if x != nil {
		return x.List
	}
	return nil
}

type Teacher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Subject string `protobuf:"bytes,4,opt,name=memory,proto3" json:"memory,omitempty"`
	Salary  int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Teacher) Reset() {
	*x = Teacher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teachers_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Teacher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Teacher) ProtoMessage() {}

func (x *Teacher) ProtoReflect() protoreflect.Message {
	mi := &file_teachers_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Teacher.ProtoReflect.Descriptor instead.
func (*Teacher) Descriptor() ([]byte, []int) {
	return file_teachers_service_proto_rawDescGZIP(), []int{2}
}

func (x *Teacher) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Teacher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Teacher) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Teacher) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *Teacher) GetSalary() int64 {
	if x != nil {
		return x.Salary
	}
	return 0
}

type TeacherId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TeacherId) Reset() {
	*x = TeacherId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teachers_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeacherId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeacherId) ProtoMessage() {}

func (x *TeacherId) ProtoReflect() protoreflect.Message {
	mi := &file_teachers_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeacherId.ProtoReflect.Descriptor instead.
func (*TeacherId) Descriptor() ([]byte, []int) {
	return file_teachers_service_proto_rawDescGZIP(), []int{3}
}

func (x *TeacherId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_teachers_service_proto protoreflect.FileDescriptor

var file_teachers_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x08, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03,
	0x63, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x1c, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa4, 0x02, 0x0a, 0x10, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x31, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x10,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x1a,
	0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x72, 0x22,
	0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teachers_service_proto_rawDescOnce sync.Once
	file_teachers_service_proto_rawDescData = file_teachers_service_proto_rawDesc
)

func file_teachers_service_proto_rawDescGZIP() []byte {
	file_teachers_service_proto_rawDescOnce.Do(func() {
		file_teachers_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teachers_service_proto_rawDescData)
	})
	return file_teachers_service_proto_rawDescData
}

var file_teachers_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teachers_service_proto_goTypes = []interface{}{
	(*GetTeachersRequest)(nil), // 0: auth.GetTeachersRequest
	(*TeacherList)(nil),        // 1: auth.TeacherList
	(*Teacher)(nil),            // 2: auth.Teacher
	(*TeacherId)(nil),          // 3: auth.TeacherId
}
var file_teachers_service_proto_depIdxs = []int32{
	2, // 0: auth.TeacherList.list:type_name -> auth.Teacher
	3, // 1: GetTeacher:input_type -> auth.TeacherId
	0, // 2: GetTeachers:input_type -> auth.GetTeachersRequest
	2, // 3: CreateTeacher:input_type -> auth.Teacher
	3, // 4: DeleteTeacher:input_type -> auth.TeacherId
	2, // 5: UpdateTeacher:input_type -> auth.Teacher
	2, // 6: GetTeacher:output_type -> auth.Teacher
	1, // 7: GetTeachers:output_type -> auth.TeacherList
	2, // 8: CreateTeacher:output_type -> auth.Teacher
	2, // 9: DeleteTeacher:output_type -> auth.Teacher
	2, // 10: UpdateTeacher:output_type -> auth.Teacher
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_teachers_service_proto_init() }
func file_teachers_service_proto_init() {
	if File_teachers_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teachers_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTeachersRequest); i {
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
		file_teachers_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherList); i {
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
		file_teachers_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Teacher); i {
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
		file_teachers_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeacherId); i {
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
			RawDescriptor: file_teachers_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teachers_service_proto_goTypes,
		DependencyIndexes: file_teachers_service_proto_depIdxs,
		MessageInfos:      file_teachers_service_proto_msgTypes,
	}.Build()
	File_teachers_service_proto = out.File
	file_teachers_service_proto_rawDesc = nil
	file_teachers_service_proto_goTypes = nil
	file_teachers_service_proto_depIdxs = nil
}