// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google_duration.proto

package google

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	math "math"
	reflect "reflect"
	sync "sync"
	time "time"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Duration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos   int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

// New constructs a new Duration from the provided time.Duration.
func New(d time.Duration) *Duration {
	nanos := d.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return &Duration{Seconds: int64(secs), Nanos: int32(nanos)}
}

// AsDuration converts x to a time.Duration,
// returning the closest duration value in the event of overflow.
func (x *Duration) AsDuration() time.Duration {
	secs := x.GetSeconds()
	nanos := x.GetNanos()
	d := time.Duration(secs) * time.Second
	overflow := d/time.Second != time.Duration(secs)
	d += time.Duration(nanos) * time.Nanosecond
	overflow = overflow || (secs < 0 && nanos < 0 && d > 0)
	overflow = overflow || (secs > 0 && nanos > 0 && d < 0)
	if overflow {
		switch {
		case secs < 0:
			return time.Duration(math.MinInt64)
		case secs > 0:
			return time.Duration(math.MaxInt64)
		}
	}
	return d
}

// IsValid reports whether the duration is valid.
// It is equivalent to CheckValid == nil.
func (x *Duration) IsValid() bool {
	return x.check() == 0
}

// CheckValid returns an error if the duration is invalid.
// In particular, it checks whether the value is within the range of
// -10000 years to +10000 years inclusive.
// An error is reported for a nil Duration.
func (x *Duration) CheckValid() error {
	switch x.check() {
	case invalidNil:
		return protoimpl.X.NewError("invalid nil Duration")
	case invalidUnderflow:
		return protoimpl.X.NewError("duration (%v) exceeds -10000 years", x)
	case invalidOverflow:
		return protoimpl.X.NewError("duration (%v) exceeds +10000 years", x)
	case invalidNanosRange:
		return protoimpl.X.NewError("duration (%v) has out-of-range nanos", x)
	case invalidNanosSign:
		return protoimpl.X.NewError("duration (%v) has seconds and nanos with different signs", x)
	default:
		return nil
	}
}

const (
	_ = iota
	invalidNil
	invalidUnderflow
	invalidOverflow
	invalidNanosRange
	invalidNanosSign
)

func (x *Duration) check() uint {
	const absDuration = 315576000000 // 10000yr * 365.25day/yr * 24hr/day * 60min/hr * 60sec/min
	secs := x.GetSeconds()
	nanos := x.GetNanos()
	switch {
	case x == nil:
		return invalidNil
	case secs < -absDuration:
		return invalidUnderflow
	case secs > +absDuration:
		return invalidOverflow
	case nanos <= -1e9 || nanos >= +1e9:
		return invalidNanosRange
	case (secs > 0 && nanos < 0) || (secs < 0 && nanos > 0):
		return invalidNanosSign
	default:
		return 0
	}
}

func (x *Duration) Reset() {
	*x = Duration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_duration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_google_duration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_google_duration_proto_rawDescGZIP(), []int{0}
}

func (x *Duration) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Duration) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_google_duration_proto protoreflect.FileDescriptor

var file_google_duration_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x3a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x42, 0x1b, 0x5a, 0x19, 0x67, 0x6f, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_duration_proto_rawDescOnce sync.Once
	file_google_duration_proto_rawDescData = file_google_duration_proto_rawDesc
)

func file_google_duration_proto_rawDescGZIP() []byte {
	file_google_duration_proto_rawDescOnce.Do(func() {
		file_google_duration_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_duration_proto_rawDescData)
	})
	return file_google_duration_proto_rawDescData
}

var file_google_duration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_duration_proto_goTypes = []interface{}{
	(*Duration)(nil), // 0: google.protobuf.Duration
}
var file_google_duration_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_duration_proto_init() }
func file_google_duration_proto_init() {
	if File_google_duration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_duration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Duration); i {
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
			RawDescriptor: file_google_duration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_duration_proto_goTypes,
		DependencyIndexes: file_google_duration_proto_depIdxs,
		MessageInfos:      file_google_duration_proto_msgTypes,
	}.Build()
	File_google_duration_proto = out.File
	file_google_duration_proto_rawDesc = nil
	file_google_duration_proto_goTypes = nil
	file_google_duration_proto_depIdxs = nil
}
