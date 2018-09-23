// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: config.proto

/*
Package v1 is a generated protocol buffer package.

It is generated from these files:
	config.proto
	resolver_map.proto
	schema.proto

It has these top-level messages:
	Config
	ResolverMap
	TypeResolver
	Resolver
	GlooResolver
	Function
	MultiFunction
	WeightedFunction
	TemplateResolver
	NodeJSResolver
	Schema
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Config is a top-level config object. It is used internally by Sqoop as a container for the entire set of config objects.
type Config struct {
	Schemas      []*Schema      `protobuf:"bytes,3,rep,name=schemas" json:"schemas,omitempty"`
	ResolverMaps []*ResolverMap `protobuf:"bytes,4,rep,name=resolver_maps,json=resolverMaps" json:"resolver_maps,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorConfig, []int{0} }

func (m *Config) GetSchemas() []*Schema {
	if m != nil {
		return m.Schemas
	}
	return nil
}

func (m *Config) GetResolverMaps() []*ResolverMap {
	if m != nil {
		return m.ResolverMaps
	}
	return nil
}

func init() {
	proto.RegisterType((*Config)(nil), "sqoop.api.v1.Config")
}
func (this *Config) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Config)
	if !ok {
		that2, ok := that.(Config)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Schemas) != len(that1.Schemas) {
		return false
	}
	for i := range this.Schemas {
		if !this.Schemas[i].Equal(that1.Schemas[i]) {
			return false
		}
	}
	if len(this.ResolverMaps) != len(that1.ResolverMaps) {
		return false
	}
	for i := range this.ResolverMaps {
		if !this.ResolverMaps[i].Equal(that1.ResolverMaps[i]) {
			return false
		}
	}
	return true
}

func init() { proto.RegisterFile("config.proto", fileDescriptorConfig) }

var fileDescriptorConfig = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0xcc, 0xc9, 0xcf, 0xd7, 0x4b,
	0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0x94, 0xe2, 0x29, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x84, 0x48, 0x49,
	0x09, 0x15, 0xa5, 0x16, 0xe7, 0xe7, 0x94, 0xa5, 0x16, 0xc5, 0xe7, 0x26, 0x16, 0x40, 0xc5, 0x44,
	0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0x54, 0xc6, 0xc5, 0xe6, 0x0c,
	0x36, 0x54, 0x48, 0x97, 0x8b, 0x1d, 0x62, 0x46, 0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0xb0, 0x1e, 0x92, 0x05, 0x7a, 0xc1, 0x60, 0xb9, 0x20, 0x98, 0x1a, 0x21, 0x5b, 0x2e, 0x5e, 0x64,
	0x4b, 0x8a, 0x25, 0x58, 0xc0, 0x9a, 0x24, 0x50, 0x34, 0x05, 0x41, 0x55, 0xf8, 0x26, 0x16, 0x04,
	0xf1, 0x14, 0x21, 0x38, 0xc5, 0x4e, 0x7a, 0x2b, 0x1e, 0xc9, 0x31, 0x46, 0x69, 0xa4, 0x67, 0x96,
	0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x17, 0xe7, 0xe7, 0xe4, 0xeb, 0x66, 0xe6, 0xeb,
	0x83, 0xf4, 0xeb, 0x17, 0x64, 0xa7, 0xeb, 0x27, 0x16, 0x64, 0xea, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0xeb, 0x97, 0x19, 0x26, 0xb1, 0x81, 0x9d, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xd6,
	0x87, 0x78, 0x03, 0x01, 0x00, 0x00,
}
