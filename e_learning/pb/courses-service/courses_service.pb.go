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

type GetCoursesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCoursesRequest) Reset() {
	*x = GetCoursesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courses_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCoursesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCoursesRequest) ProtoMessage() {}

func (x *GetCoursesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_courses_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCoursesRequest.ProtoReflect.Descriptor instead.
func (*GetCoursesRequest) Descriptor() ([]byte, []int) {
	return file_courses_service_proto_rawDescGZIP(), []int{0}
}

type CourseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Course `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *CourseList) Reset() {
	*x = CourseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courses_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseList) ProtoMessage() {}

func (x *CourseList) ProtoReflect() protoreflect.Message {
	mi := &file_courses_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseList.ProtoReflect.Descriptor instead.
func (*CourseList) Descriptor() ([]byte, []int) {
	return file_courses_service_proto_rawDescGZIP(), []int{1}
}

func (x *CourseList) GetList() []*Course {
	if x != nil {
		return x.List
	}
	return nil
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Credits int64  `protobuf:"varint,4,opt,name=credits,proto3" json:"credits,omitempty"`
	Price   int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courses_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_courses_service_proto_msgTypes[2]
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
	return file_courses_service_proto_rawDescGZIP(), []int{2}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}


func (x *Course) GetCredits() int64 {
	if x != nil {
		return x.Credits
	}
	return 0
}

func (x *Course) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CourseId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CourseId) Reset() {
	*x = CourseId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courses_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseId) ProtoMessage() {}

func (x *CourseId) ProtoReflect() protoreflect.Message {
	mi := &file_courses_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseId.ProtoReflect.Descriptor instead.
func (*CourseId) Descriptor() ([]byte, []int) {
	return file_courses_service_proto_rawDescGZIP(), []int{3}
}

func (x *CourseId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_courses_service_proto protoreflect.FileDescriptor

var file_courses_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x12, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x2c, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x6d,
	0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x19, 0x0a,
	0x07, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf4, 0x01, 0x0a, 0x0d, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x00, 0x42,
	0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_courses_service_proto_rawDescOnce sync.Once
	file_courses_service_proto_rawDescData = file_courses_service_proto_rawDesc
)

func file_courses_service_proto_rawDescGZIP() []byte {
	file_courses_service_proto_rawDescOnce.Do(func() {
		file_courses_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_courses_service_proto_rawDescData)
	})
	return file_courses_service_proto_rawDescData
}

var file_courses_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_courses_service_proto_goTypes = []interface{}{
	(*GetCoursesRequest)(nil), // 0: auth.GetCoursesRequest
	(*CourseList)(nil),         // 1: auth.CourseList
	(*Course)(nil),            // 2: auth.Course
	(*CourseId)(nil),          // 3: auth.CourseId
}
var file_courses_service_proto_depIdxs = []int32{
	2, // 0: auth.CourseList.list:type_name -> auth.Course
	3, // 1: auth.CoursesService.GetCourse:input_type -> auth.CourseId
	0, // 2: auth.CoursesService.GetCourses:input_type -> auth.GetCoursesRequest
	2, // 3: auth.CoursesService.CreateCourse:input_type -> auth.Course
	3, // 4: auth.CoursesService.DeleteCourse:input_type -> auth.CourseId
	2, // 5: auth.CoursesService.UpdateCourse:input_type -> auth.Course
	2, // 6: auth.CoursesService.GetCourse:output_type -> auth.Course
	1, // 7: auth.CoursesService.GetCourses:output_type -> auth.CourseList
	2, // 8: auth.CoursesService.CreateCourse:output_type -> auth.Course
	2, // 9: auth.CoursesService.DeleteCourse:output_type -> auth.Course
	2, // 10: auth.CoursesService.UpdateCourse:output_type -> auth.Course
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_courses_service_proto_init() }
func file_courses_service_proto_init() {
	if File_courses_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courses_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCoursesRequest); i {
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
		file_courses_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseList); i {
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
		file_courses_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_courses_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseId); i {
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
			RawDescriptor: file_courses_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courses_service_proto_goTypes,
		DependencyIndexes: file_courses_service_proto_depIdxs,
		MessageInfos:      file_courses_service_proto_msgTypes,
	}.Build()
	File_courses_service_proto = out.File
	file_courses_service_proto_rawDesc = nil
	file_courses_service_proto_goTypes = nil
	file_courses_service_proto_depIdxs = nil
}
