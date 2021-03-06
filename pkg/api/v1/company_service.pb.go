// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: company_service.proto

package pb

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

type AirlineCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company_ID  int64                  `protobuf:"varint,1,opt,name=company_ID,json=companyID,proto3" json:"company_ID,omitempty"`
	CompanyName string                 `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	IataCode    string                 `protobuf:"bytes,3,opt,name=iata_code,json=iataCode,proto3" json:"iata_code,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` //    string mainAirport = 4;
}

func (x *AirlineCompany) Reset() {
	*x = AirlineCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirlineCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirlineCompany) ProtoMessage() {}

func (x *AirlineCompany) ProtoReflect() protoreflect.Message {
	mi := &file_company_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirlineCompany.ProtoReflect.Descriptor instead.
func (*AirlineCompany) Descriptor() ([]byte, []int) {
	return file_company_service_proto_rawDescGZIP(), []int{0}
}

func (x *AirlineCompany) GetCompany_ID() int64 {
	if x != nil {
		return x.Company_ID
	}
	return 0
}

func (x *AirlineCompany) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *AirlineCompany) GetIataCode() string {
	if x != nil {
		return x.IataCode
	}
	return ""
}

func (x *AirlineCompany) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// API versioning: it is my best practice to specify version explicitly.
//string api = 1;
type CreateAirlineCompanyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Company_ID  int64                  `protobuf:"varint,1,opt,name=company_ID,json=companyID,proto3" json:"company_ID,omitempty"`
	CompanyName string                 `protobuf:"bytes,2,opt,name=company_name,json=companyName,proto3" json:"company_name,omitempty"`
	IataCode    string                 `protobuf:"bytes,3,opt,name=iata_code,json=iataCode,proto3" json:"iata_code,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` //    string mainAirport = 4;
}

func (x *CreateAirlineCompanyParams) Reset() {
	*x = CreateAirlineCompanyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAirlineCompanyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAirlineCompanyParams) ProtoMessage() {}

func (x *CreateAirlineCompanyParams) ProtoReflect() protoreflect.Message {
	mi := &file_company_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAirlineCompanyParams.ProtoReflect.Descriptor instead.
func (*CreateAirlineCompanyParams) Descriptor() ([]byte, []int) {
	return file_company_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAirlineCompanyParams) GetCompany_ID() int64 {
	if x != nil {
		return x.Company_ID
	}
	return 0
}

func (x *CreateAirlineCompanyParams) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CreateAirlineCompanyParams) GetIataCode() string {
	if x != nil {
		return x.IataCode
	}
	return ""
}

func (x *CreateAirlineCompanyParams) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetAirlineCompanyParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompanyID int64 `protobuf:"varint,1,opt,name=companyID,proto3" json:"companyID,omitempty"`
}

func (x *GetAirlineCompanyParams) Reset() {
	*x = GetAirlineCompanyParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAirlineCompanyParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAirlineCompanyParams) ProtoMessage() {}

func (x *GetAirlineCompanyParams) ProtoReflect() protoreflect.Message {
	mi := &file_company_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAirlineCompanyParams.ProtoReflect.Descriptor instead.
func (*GetAirlineCompanyParams) Descriptor() ([]byte, []int) {
	return file_company_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetAirlineCompanyParams) GetCompanyID() int64 {
	if x != nil {
		return x.CompanyID
	}
	return 0
}

type ListAirlineCompanyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*AirlineCompany `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *ListAirlineCompanyResponse) Reset() {
	*x = ListAirlineCompanyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAirlineCompanyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAirlineCompanyResponse) ProtoMessage() {}

func (x *ListAirlineCompanyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_company_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAirlineCompanyResponse.ProtoReflect.Descriptor instead.
func (*ListAirlineCompanyResponse) Descriptor() ([]byte, []int) {
	return file_company_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAirlineCompanyResponse) GetValue() []*AirlineCompany {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_company_service_proto protoreflect.FileDescriptor

var file_company_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x61, 0x77, 0x6f, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x41, 0x69, 0x72, 0x6c,
	0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xb6, 0x01, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x49, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x61, 0x74, 0x61, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x37, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x49, 0x44, 0x22, 0x51, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x69,
	0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x61, 0x77, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x7c, 0x0a, 0x16, 0x41, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x69, 0x72,
	0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x29, 0x2e, 0x63, 0x61,
	0x77, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x77, 0x6f, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x69, 0x72, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_service_proto_rawDescOnce sync.Once
	file_company_service_proto_rawDescData = file_company_service_proto_rawDesc
)

func file_company_service_proto_rawDescGZIP() []byte {
	file_company_service_proto_rawDescOnce.Do(func() {
		file_company_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_service_proto_rawDescData)
	})
	return file_company_service_proto_rawDescData
}

var file_company_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_company_service_proto_goTypes = []interface{}{
	(*AirlineCompany)(nil),             // 0: cawo.v1alpha1.AirlineCompany
	(*CreateAirlineCompanyParams)(nil), // 1: cawo.v1alpha1.CreateAirlineCompanyParams
	(*GetAirlineCompanyParams)(nil),    // 2: cawo.v1alpha1.GetAirlineCompanyParams
	(*ListAirlineCompanyResponse)(nil), // 3: cawo.v1alpha1.ListAirlineCompanyResponse
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_company_service_proto_depIdxs = []int32{
	4, // 0: cawo.v1alpha1.AirlineCompany.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: cawo.v1alpha1.CreateAirlineCompanyParams.created_at:type_name -> google.protobuf.Timestamp
	0, // 2: cawo.v1alpha1.ListAirlineCompanyResponse.value:type_name -> cawo.v1alpha1.AirlineCompany
	1, // 3: cawo.v1alpha1.AirlineCompanyServices.CreateAirlineCompany:input_type -> cawo.v1alpha1.CreateAirlineCompanyParams
	0, // 4: cawo.v1alpha1.AirlineCompanyServices.CreateAirlineCompany:output_type -> cawo.v1alpha1.AirlineCompany
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_company_service_proto_init() }
func file_company_service_proto_init() {
	if File_company_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_company_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirlineCompany); i {
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
		file_company_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAirlineCompanyParams); i {
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
		file_company_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAirlineCompanyParams); i {
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
		file_company_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAirlineCompanyResponse); i {
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
			RawDescriptor: file_company_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_company_service_proto_goTypes,
		DependencyIndexes: file_company_service_proto_depIdxs,
		MessageInfos:      file_company_service_proto_msgTypes,
	}.Build()
	File_company_service_proto = out.File
	file_company_service_proto_rawDesc = nil
	file_company_service_proto_goTypes = nil
	file_company_service_proto_depIdxs = nil
}
