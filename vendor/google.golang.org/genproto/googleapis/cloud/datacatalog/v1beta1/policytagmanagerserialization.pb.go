// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/policytagmanagerserialization.proto

package datacatalog

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/iam/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Message capturing a taxonomy and its policy tag hierarchy as a nested proto.
// Used for taxonomy import/export and mutation.
type SerializedTaxonomy struct {
	// Required. Display name of the taxonomy. Max 200 bytes when encoded in
	// UTF-8.
	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description of the serialized taxonomy. The length of the
	// description is limited to 2000 bytes when encoded in UTF-8. If not set,
	// defaults to an empty description.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Top level policy tags associated with the taxonomy if any.
	PolicyTags           []*SerializedPolicyTag `protobuf:"bytes,3,rep,name=policy_tags,json=policyTags,proto3" json:"policy_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SerializedTaxonomy) Reset()         { *m = SerializedTaxonomy{} }
func (m *SerializedTaxonomy) String() string { return proto.CompactTextString(m) }
func (*SerializedTaxonomy) ProtoMessage()    {}
func (*SerializedTaxonomy) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{0}
}

func (m *SerializedTaxonomy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializedTaxonomy.Unmarshal(m, b)
}
func (m *SerializedTaxonomy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializedTaxonomy.Marshal(b, m, deterministic)
}
func (m *SerializedTaxonomy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedTaxonomy.Merge(m, src)
}
func (m *SerializedTaxonomy) XXX_Size() int {
	return xxx_messageInfo_SerializedTaxonomy.Size(m)
}
func (m *SerializedTaxonomy) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedTaxonomy.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedTaxonomy proto.InternalMessageInfo

func (m *SerializedTaxonomy) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SerializedTaxonomy) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SerializedTaxonomy) GetPolicyTags() []*SerializedPolicyTag {
	if m != nil {
		return m.PolicyTags
	}
	return nil
}

// Message representing one policy tag when exported as a nested proto.
type SerializedPolicyTag struct {
	// Required. Display name of the policy tag. Max 200 bytes when encoded in
	// UTF-8.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Description of the serialized policy tag. The length of the
	// description is limited to 2000 bytes when encoded in UTF-8. If not set,
	// defaults to an empty description.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Children of the policy tag if any.
	ChildPolicyTags      []*SerializedPolicyTag `protobuf:"bytes,4,rep,name=child_policy_tags,json=childPolicyTags,proto3" json:"child_policy_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SerializedPolicyTag) Reset()         { *m = SerializedPolicyTag{} }
func (m *SerializedPolicyTag) String() string { return proto.CompactTextString(m) }
func (*SerializedPolicyTag) ProtoMessage()    {}
func (*SerializedPolicyTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{1}
}

func (m *SerializedPolicyTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SerializedPolicyTag.Unmarshal(m, b)
}
func (m *SerializedPolicyTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SerializedPolicyTag.Marshal(b, m, deterministic)
}
func (m *SerializedPolicyTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SerializedPolicyTag.Merge(m, src)
}
func (m *SerializedPolicyTag) XXX_Size() int {
	return xxx_messageInfo_SerializedPolicyTag.Size(m)
}
func (m *SerializedPolicyTag) XXX_DiscardUnknown() {
	xxx_messageInfo_SerializedPolicyTag.DiscardUnknown(m)
}

var xxx_messageInfo_SerializedPolicyTag proto.InternalMessageInfo

func (m *SerializedPolicyTag) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *SerializedPolicyTag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SerializedPolicyTag) GetChildPolicyTags() []*SerializedPolicyTag {
	if m != nil {
		return m.ChildPolicyTags
	}
	return nil
}

// Request message for
// [ImportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ImportTaxonomies].
type ImportTaxonomiesRequest struct {
	// Required. Resource name of project that the newly created taxonomies will
	// belong to.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Source taxonomies to be imported in a tree structure.
	//
	// Types that are valid to be assigned to Source:
	//	*ImportTaxonomiesRequest_InlineSource
	Source               isImportTaxonomiesRequest_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ImportTaxonomiesRequest) Reset()         { *m = ImportTaxonomiesRequest{} }
func (m *ImportTaxonomiesRequest) String() string { return proto.CompactTextString(m) }
func (*ImportTaxonomiesRequest) ProtoMessage()    {}
func (*ImportTaxonomiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{2}
}

func (m *ImportTaxonomiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportTaxonomiesRequest.Unmarshal(m, b)
}
func (m *ImportTaxonomiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportTaxonomiesRequest.Marshal(b, m, deterministic)
}
func (m *ImportTaxonomiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportTaxonomiesRequest.Merge(m, src)
}
func (m *ImportTaxonomiesRequest) XXX_Size() int {
	return xxx_messageInfo_ImportTaxonomiesRequest.Size(m)
}
func (m *ImportTaxonomiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportTaxonomiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImportTaxonomiesRequest proto.InternalMessageInfo

func (m *ImportTaxonomiesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

type isImportTaxonomiesRequest_Source interface {
	isImportTaxonomiesRequest_Source()
}

type ImportTaxonomiesRequest_InlineSource struct {
	InlineSource *InlineSource `protobuf:"bytes,2,opt,name=inline_source,json=inlineSource,proto3,oneof"`
}

func (*ImportTaxonomiesRequest_InlineSource) isImportTaxonomiesRequest_Source() {}

func (m *ImportTaxonomiesRequest) GetSource() isImportTaxonomiesRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *ImportTaxonomiesRequest) GetInlineSource() *InlineSource {
	if x, ok := m.GetSource().(*ImportTaxonomiesRequest_InlineSource); ok {
		return x.InlineSource
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ImportTaxonomiesRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ImportTaxonomiesRequest_InlineSource)(nil),
	}
}

// Inline source used for taxonomies import.
type InlineSource struct {
	// Required. Taxonomies to be imported.
	Taxonomies           []*SerializedTaxonomy `protobuf:"bytes,1,rep,name=taxonomies,proto3" json:"taxonomies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *InlineSource) Reset()         { *m = InlineSource{} }
func (m *InlineSource) String() string { return proto.CompactTextString(m) }
func (*InlineSource) ProtoMessage()    {}
func (*InlineSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{3}
}

func (m *InlineSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InlineSource.Unmarshal(m, b)
}
func (m *InlineSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InlineSource.Marshal(b, m, deterministic)
}
func (m *InlineSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InlineSource.Merge(m, src)
}
func (m *InlineSource) XXX_Size() int {
	return xxx_messageInfo_InlineSource.Size(m)
}
func (m *InlineSource) XXX_DiscardUnknown() {
	xxx_messageInfo_InlineSource.DiscardUnknown(m)
}

var xxx_messageInfo_InlineSource proto.InternalMessageInfo

func (m *InlineSource) GetTaxonomies() []*SerializedTaxonomy {
	if m != nil {
		return m.Taxonomies
	}
	return nil
}

// Response message for
// [ImportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ImportTaxonomies].
type ImportTaxonomiesResponse struct {
	// Taxonomies that were imported.
	Taxonomies           []*Taxonomy `protobuf:"bytes,1,rep,name=taxonomies,proto3" json:"taxonomies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ImportTaxonomiesResponse) Reset()         { *m = ImportTaxonomiesResponse{} }
func (m *ImportTaxonomiesResponse) String() string { return proto.CompactTextString(m) }
func (*ImportTaxonomiesResponse) ProtoMessage()    {}
func (*ImportTaxonomiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{4}
}

func (m *ImportTaxonomiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImportTaxonomiesResponse.Unmarshal(m, b)
}
func (m *ImportTaxonomiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImportTaxonomiesResponse.Marshal(b, m, deterministic)
}
func (m *ImportTaxonomiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImportTaxonomiesResponse.Merge(m, src)
}
func (m *ImportTaxonomiesResponse) XXX_Size() int {
	return xxx_messageInfo_ImportTaxonomiesResponse.Size(m)
}
func (m *ImportTaxonomiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImportTaxonomiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImportTaxonomiesResponse proto.InternalMessageInfo

func (m *ImportTaxonomiesResponse) GetTaxonomies() []*Taxonomy {
	if m != nil {
		return m.Taxonomies
	}
	return nil
}

// Request message for
// [ExportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ExportTaxonomies].
type ExportTaxonomiesRequest struct {
	// Required. Resource name of the project that taxonomies to be exported
	// will share.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. Resource names of the taxonomies to be exported.
	Taxonomies []string `protobuf:"bytes,2,rep,name=taxonomies,proto3" json:"taxonomies,omitempty"`
	// Required. Taxonomies export destination.
	//
	// Types that are valid to be assigned to Destination:
	//	*ExportTaxonomiesRequest_SerializedTaxonomies
	Destination          isExportTaxonomiesRequest_Destination `protobuf_oneof:"destination"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ExportTaxonomiesRequest) Reset()         { *m = ExportTaxonomiesRequest{} }
func (m *ExportTaxonomiesRequest) String() string { return proto.CompactTextString(m) }
func (*ExportTaxonomiesRequest) ProtoMessage()    {}
func (*ExportTaxonomiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{5}
}

func (m *ExportTaxonomiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportTaxonomiesRequest.Unmarshal(m, b)
}
func (m *ExportTaxonomiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportTaxonomiesRequest.Marshal(b, m, deterministic)
}
func (m *ExportTaxonomiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTaxonomiesRequest.Merge(m, src)
}
func (m *ExportTaxonomiesRequest) XXX_Size() int {
	return xxx_messageInfo_ExportTaxonomiesRequest.Size(m)
}
func (m *ExportTaxonomiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTaxonomiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTaxonomiesRequest proto.InternalMessageInfo

func (m *ExportTaxonomiesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ExportTaxonomiesRequest) GetTaxonomies() []string {
	if m != nil {
		return m.Taxonomies
	}
	return nil
}

type isExportTaxonomiesRequest_Destination interface {
	isExportTaxonomiesRequest_Destination()
}

type ExportTaxonomiesRequest_SerializedTaxonomies struct {
	SerializedTaxonomies bool `protobuf:"varint,3,opt,name=serialized_taxonomies,json=serializedTaxonomies,proto3,oneof"`
}

func (*ExportTaxonomiesRequest_SerializedTaxonomies) isExportTaxonomiesRequest_Destination() {}

func (m *ExportTaxonomiesRequest) GetDestination() isExportTaxonomiesRequest_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *ExportTaxonomiesRequest) GetSerializedTaxonomies() bool {
	if x, ok := m.GetDestination().(*ExportTaxonomiesRequest_SerializedTaxonomies); ok {
		return x.SerializedTaxonomies
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExportTaxonomiesRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExportTaxonomiesRequest_SerializedTaxonomies)(nil),
	}
}

// Response message for
// [ExportTaxonomies][google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization.ExportTaxonomies].
type ExportTaxonomiesResponse struct {
	// List of taxonomies and policy tags in a tree structure.
	Taxonomies           []*SerializedTaxonomy `protobuf:"bytes,1,rep,name=taxonomies,proto3" json:"taxonomies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExportTaxonomiesResponse) Reset()         { *m = ExportTaxonomiesResponse{} }
func (m *ExportTaxonomiesResponse) String() string { return proto.CompactTextString(m) }
func (*ExportTaxonomiesResponse) ProtoMessage()    {}
func (*ExportTaxonomiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bcaf8b94fa1fe913, []int{6}
}

func (m *ExportTaxonomiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExportTaxonomiesResponse.Unmarshal(m, b)
}
func (m *ExportTaxonomiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExportTaxonomiesResponse.Marshal(b, m, deterministic)
}
func (m *ExportTaxonomiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExportTaxonomiesResponse.Merge(m, src)
}
func (m *ExportTaxonomiesResponse) XXX_Size() int {
	return xxx_messageInfo_ExportTaxonomiesResponse.Size(m)
}
func (m *ExportTaxonomiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExportTaxonomiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExportTaxonomiesResponse proto.InternalMessageInfo

func (m *ExportTaxonomiesResponse) GetTaxonomies() []*SerializedTaxonomy {
	if m != nil {
		return m.Taxonomies
	}
	return nil
}

func init() {
	proto.RegisterType((*SerializedTaxonomy)(nil), "google.cloud.datacatalog.v1beta1.SerializedTaxonomy")
	proto.RegisterType((*SerializedPolicyTag)(nil), "google.cloud.datacatalog.v1beta1.SerializedPolicyTag")
	proto.RegisterType((*ImportTaxonomiesRequest)(nil), "google.cloud.datacatalog.v1beta1.ImportTaxonomiesRequest")
	proto.RegisterType((*InlineSource)(nil), "google.cloud.datacatalog.v1beta1.InlineSource")
	proto.RegisterType((*ImportTaxonomiesResponse)(nil), "google.cloud.datacatalog.v1beta1.ImportTaxonomiesResponse")
	proto.RegisterType((*ExportTaxonomiesRequest)(nil), "google.cloud.datacatalog.v1beta1.ExportTaxonomiesRequest")
	proto.RegisterType((*ExportTaxonomiesResponse)(nil), "google.cloud.datacatalog.v1beta1.ExportTaxonomiesResponse")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/policytagmanagerserialization.proto", fileDescriptor_bcaf8b94fa1fe913)
}

var fileDescriptor_bcaf8b94fa1fe913 = []byte{
	// 734 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdf, 0x6a, 0x13, 0x4f,
	0x14, 0xee, 0x26, 0xa5, 0xf4, 0x37, 0x69, 0xf9, 0xd5, 0x51, 0x69, 0x0c, 0x8a, 0x61, 0xfd, 0x43,
	0x89, 0xb8, 0x43, 0xab, 0x45, 0x8c, 0x8a, 0x24, 0x5a, 0x68, 0x2d, 0x96, 0x90, 0x56, 0x41, 0x6f,
	0xc2, 0x74, 0x33, 0xdd, 0x8c, 0xec, 0xce, 0x8c, 0x3b, 0x93, 0xfe, 0x51, 0xbc, 0xf1, 0x11, 0xf4,
	0xd6, 0xe7, 0xf0, 0x56, 0xf0, 0x52, 0x10, 0x44, 0x5f, 0xa0, 0x17, 0xbe, 0x80, 0x77, 0xe2, 0x95,
	0x64, 0x76, 0x92, 0x8c, 0x49, 0x63, 0xd2, 0xaa, 0x77, 0xc9, 0x9e, 0x73, 0xbe, 0xf3, 0x7d, 0xdf,
	0x39, 0x33, 0x03, 0xee, 0x06, 0x9c, 0x07, 0x21, 0x41, 0x7e, 0xc8, 0x9b, 0x75, 0x54, 0xc7, 0x0a,
	0xfb, 0x58, 0xe1, 0x90, 0x07, 0x68, 0x7b, 0x7e, 0x93, 0x28, 0x3c, 0x8f, 0x04, 0x0f, 0xa9, 0xbf,
	0xa7, 0x70, 0x10, 0x61, 0x86, 0x03, 0x12, 0x4b, 0x12, 0x53, 0x1c, 0xd2, 0x67, 0x58, 0x51, 0xce,
	0x3c, 0x11, 0x73, 0xc5, 0x61, 0x3e, 0x41, 0xf1, 0x34, 0x8a, 0x67, 0xa1, 0x78, 0x06, 0x25, 0x77,
	0xda, 0xf4, 0xc1, 0x82, 0x22, 0xcc, 0x18, 0x57, 0xba, 0x5c, 0x26, 0xf5, 0xb9, 0x59, 0x2b, 0xea,
	0x87, 0x94, 0x30, 0x65, 0x02, 0x67, 0xad, 0xc0, 0x16, 0x25, 0x61, 0xbd, 0xb6, 0x49, 0x1a, 0x78,
	0x9b, 0xf2, 0xd8, 0x24, 0x9c, 0xb2, 0x12, 0x62, 0x22, 0x79, 0x33, 0xf6, 0x89, 0x09, 0x5d, 0x3b,
	0xb4, 0x34, 0x53, 0x98, 0x33, 0x85, 0x14, 0x47, 0x68, 0xbb, 0x9d, 0x95, 0xc4, 0xdc, 0xb7, 0x0e,
	0x80, 0xeb, 0xc6, 0x01, 0x52, 0xdf, 0xc0, 0xbb, 0x9c, 0xf1, 0x68, 0x0f, 0x5e, 0x04, 0x53, 0x75,
	0x2a, 0x45, 0x88, 0xf7, 0x6a, 0x0c, 0x47, 0x24, 0xeb, 0xe4, 0x9d, 0xb9, 0xff, 0xca, 0xe9, 0xfd,
	0x52, 0xaa, 0x9a, 0x31, 0x81, 0x35, 0x1c, 0x11, 0x98, 0x07, 0x99, 0x3a, 0x91, 0x7e, 0x4c, 0x45,
	0x4b, 0x7e, 0x36, 0xd5, 0x4a, 0xab, 0xda, 0x9f, 0xe0, 0x43, 0x90, 0x49, 0x1a, 0xd6, 0x14, 0x0e,
	0x64, 0x36, 0x9d, 0x4f, 0xcf, 0x65, 0x16, 0x16, 0xbd, 0x61, 0x06, 0x7b, 0x5d, 0x52, 0x15, 0x5d,
	0xbe, 0x81, 0x83, 0x2a, 0x10, 0xed, 0x9f, 0xd2, 0x7d, 0xef, 0x80, 0xe3, 0x07, 0xe4, 0xf4, 0x31,
	0x4f, 0x8d, 0xc6, 0x3c, 0xdd, 0xcf, 0x1c, 0x83, 0x63, 0x7e, 0x83, 0x86, 0xf5, 0x9a, 0xcd, 0x7f,
	0xfc, 0x4f, 0xf8, 0xff, 0xaf, 0xf1, 0x2a, 0x5d, 0x11, 0xef, 0x1c, 0x30, 0xbb, 0x12, 0x09, 0x1e,
	0x2b, 0xe3, 0x3c, 0x25, 0xb2, 0x4a, 0x9e, 0x36, 0x89, 0x54, 0xf0, 0x0e, 0x98, 0x10, 0x38, 0x26,
	0x4c, 0x19, 0xf3, 0x2f, 0xed, 0x97, 0x52, 0x3f, 0x4a, 0x17, 0xe0, 0x39, 0xbb, 0x5b, 0x42, 0x03,
	0x0b, 0x2a, 0x3d, 0x9f, 0x47, 0xa8, 0x3d, 0xbf, 0xaa, 0x29, 0x85, 0x0f, 0xc0, 0x34, 0x65, 0x21,
	0x65, 0xa4, 0x96, 0xac, 0x92, 0xb6, 0x23, 0xb3, 0xe0, 0x0d, 0xe7, 0xbf, 0xa2, 0xcb, 0xd6, 0x75,
	0xd5, 0xf2, 0x58, 0x75, 0x8a, 0x5a, 0xff, 0xcb, 0x93, 0x60, 0x22, 0xc1, 0x73, 0x29, 0x98, 0xb2,
	0x33, 0xe1, 0x23, 0x00, 0x54, 0x47, 0x4a, 0xd6, 0xd1, 0x6e, 0x5d, 0x3d, 0x8c, 0x5b, 0x6d, 0x09,
	0xc9, 0xc8, 0x2c, 0x30, 0x77, 0x0b, 0x64, 0xfb, 0xbd, 0x92, 0x82, 0x33, 0x49, 0xe0, 0xbd, 0x03,
	0xda, 0x16, 0x86, 0xb7, 0xed, 0xf8, 0x65, 0xf7, 0xf9, 0xe6, 0x80, 0xd9, 0xa5, 0xdd, 0x7f, 0x38,
	0x94, 0xd5, 0x5f, 0xc8, 0xa6, 0xf2, 0xe9, 0x2e, 0x10, 0x18, 0x09, 0xc8, 0x2a, 0x87, 0x8b, 0xe0,
	0xa4, 0xec, 0x98, 0x57, 0xb3, 0x70, 0x5b, 0x1b, 0x3d, 0xb9, 0x3c, 0x56, 0x3d, 0x21, 0x7b, 0xbd,
	0xa5, 0x44, 0x96, 0xa7, 0xf5, 0xfa, 0x2b, 0xca, 0xf4, 0xbd, 0xe5, 0x0a, 0x90, 0xed, 0x97, 0x6c,
	0xbc, 0xdd, 0xf8, 0x5b, 0x23, 0xb5, 0x79, 0x2f, 0xbc, 0x1a, 0x07, 0x67, 0x3a, 0x27, 0xe1, 0x7e,
	0x72, 0x5f, 0xad, 0xdb, 0x57, 0x31, 0xfc, 0xe4, 0x80, 0x99, 0xde, 0x81, 0xc3, 0xeb, 0x23, 0x6c,
	0xee, 0xc1, 0x07, 0x2a, 0x57, 0x3c, 0x4a, 0x69, 0xe2, 0x81, 0xbb, 0xf4, 0xf2, 0xcb, 0xd7, 0xd7,
	0xa9, 0xdb, 0x6e, 0xb1, 0x73, 0xd7, 0x3e, 0x4f, 0x86, 0x79, 0x4b, 0xc4, 0xfc, 0x09, 0xf1, 0x95,
	0x44, 0x05, 0x14, 0x72, 0x3f, 0x79, 0x05, 0x50, 0xe1, 0x05, 0xea, 0x4a, 0x2d, 0x52, 0x8d, 0x5a,
	0x74, 0x0a, 0xf0, 0xa3, 0x03, 0x66, 0x7a, 0x7d, 0x1e, 0x45, 0xd2, 0x80, 0x75, 0x1c, 0x45, 0xd2,
	0xa0, 0xb1, 0xba, 0x65, 0x2d, 0xe9, 0x26, 0x3c, 0x92, 0x24, 0xa2, 0x51, 0x73, 0x6b, 0x1f, 0x4a,
	0xb9, 0xc1, 0x0b, 0xfb, 0xb9, 0xe4, 0x35, 0x94, 0x12, 0xb2, 0x88, 0xd0, 0xce, 0xce, 0x4e, 0xef,
	0x36, 0xe3, 0xa6, 0x6a, 0x24, 0x6f, 0xd9, 0x65, 0x11, 0x62, 0xb5, 0xc5, 0xe3, 0xa8, 0xfc, 0xc6,
	0x01, 0xe7, 0x7d, 0x1e, 0x0d, 0x55, 0x55, 0x76, 0x7f, 0xbb, 0x3a, 0x95, 0xd6, 0xd3, 0x56, 0x71,
	0x1e, 0xaf, 0x1a, 0x9c, 0x80, 0x87, 0x98, 0x05, 0x1e, 0x8f, 0x03, 0x14, 0x10, 0xa6, 0x1f, 0x3e,
	0xd4, 0x65, 0x33, 0xf8, 0x41, 0xbd, 0x61, 0x7d, 0xfb, 0xee, 0x38, 0x9b, 0x13, 0xba, 0xf4, 0xca,
	0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x68, 0xaf, 0x0a, 0x65, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolicyTagManagerSerializationClient is the client API for PolicyTagManagerSerialization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolicyTagManagerSerializationClient interface {
	// Imports all taxonomies and their policy tags to a project as new
	// taxonomies.
	//
	// This method provides a bulk taxonomy / policy tag creation using nested
	// proto structure.
	ImportTaxonomies(ctx context.Context, in *ImportTaxonomiesRequest, opts ...grpc.CallOption) (*ImportTaxonomiesResponse, error)
	// Exports all taxonomies and their policy tags in a project.
	//
	// This method generates SerializedTaxonomy protos with nested policy tags
	// that can be used as an input for future ImportTaxonomies calls.
	ExportTaxonomies(ctx context.Context, in *ExportTaxonomiesRequest, opts ...grpc.CallOption) (*ExportTaxonomiesResponse, error)
}

type policyTagManagerSerializationClient struct {
	cc *grpc.ClientConn
}

func NewPolicyTagManagerSerializationClient(cc *grpc.ClientConn) PolicyTagManagerSerializationClient {
	return &policyTagManagerSerializationClient{cc}
}

func (c *policyTagManagerSerializationClient) ImportTaxonomies(ctx context.Context, in *ImportTaxonomiesRequest, opts ...grpc.CallOption) (*ImportTaxonomiesResponse, error) {
	out := new(ImportTaxonomiesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization/ImportTaxonomies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *policyTagManagerSerializationClient) ExportTaxonomies(ctx context.Context, in *ExportTaxonomiesRequest, opts ...grpc.CallOption) (*ExportTaxonomiesResponse, error) {
	out := new(ExportTaxonomiesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization/ExportTaxonomies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolicyTagManagerSerializationServer is the server API for PolicyTagManagerSerialization service.
type PolicyTagManagerSerializationServer interface {
	// Imports all taxonomies and their policy tags to a project as new
	// taxonomies.
	//
	// This method provides a bulk taxonomy / policy tag creation using nested
	// proto structure.
	ImportTaxonomies(context.Context, *ImportTaxonomiesRequest) (*ImportTaxonomiesResponse, error)
	// Exports all taxonomies and their policy tags in a project.
	//
	// This method generates SerializedTaxonomy protos with nested policy tags
	// that can be used as an input for future ImportTaxonomies calls.
	ExportTaxonomies(context.Context, *ExportTaxonomiesRequest) (*ExportTaxonomiesResponse, error)
}

// UnimplementedPolicyTagManagerSerializationServer can be embedded to have forward compatible implementations.
type UnimplementedPolicyTagManagerSerializationServer struct {
}

func (*UnimplementedPolicyTagManagerSerializationServer) ImportTaxonomies(ctx context.Context, req *ImportTaxonomiesRequest) (*ImportTaxonomiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportTaxonomies not implemented")
}
func (*UnimplementedPolicyTagManagerSerializationServer) ExportTaxonomies(ctx context.Context, req *ExportTaxonomiesRequest) (*ExportTaxonomiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportTaxonomies not implemented")
}

func RegisterPolicyTagManagerSerializationServer(s *grpc.Server, srv PolicyTagManagerSerializationServer) {
	s.RegisterService(&_PolicyTagManagerSerialization_serviceDesc, srv)
}

func _PolicyTagManagerSerialization_ImportTaxonomies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportTaxonomiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyTagManagerSerializationServer).ImportTaxonomies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization/ImportTaxonomies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyTagManagerSerializationServer).ImportTaxonomies(ctx, req.(*ImportTaxonomiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolicyTagManagerSerialization_ExportTaxonomies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportTaxonomiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolicyTagManagerSerializationServer).ExportTaxonomies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization/ExportTaxonomies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolicyTagManagerSerializationServer).ExportTaxonomies(ctx, req.(*ExportTaxonomiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PolicyTagManagerSerialization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.datacatalog.v1beta1.PolicyTagManagerSerialization",
	HandlerType: (*PolicyTagManagerSerializationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportTaxonomies",
			Handler:    _PolicyTagManagerSerialization_ImportTaxonomies_Handler,
		},
		{
			MethodName: "ExportTaxonomies",
			Handler:    _PolicyTagManagerSerialization_ExportTaxonomies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/datacatalog/v1beta1/policytagmanagerserialization.proto",
}
