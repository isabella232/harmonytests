// Code generated by protoc-gen-gogo.
// source: thetest.proto
// DO NOT EDIT!

package gogoprototest

import proto "code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import code_google_com_p_gogoprotobuf_test_custom "code.google.com/p/gogoprotobuf/test/custom"

import fmt "fmt"
import strings "strings"
import reflect "reflect"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SimpleMessage struct {
	Id               code_google_com_p_gogoprotobuf_test_custom.Uuid `protobuf:"bytes,1,opt,customtype=code.google.com/p/gogoprotobuf/test/custom.Uuid" json:"Id"`
	SimpleName       string                                          `protobuf:"bytes,2,opt" json:"SimpleName"`
	Time             *int64                                          `protobuf:"varint,3,opt" json:"Time,omitempty"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *SimpleMessage) Reset()      { *m = SimpleMessage{} }
func (*SimpleMessage) ProtoMessage() {}

type NidOptNative struct {
	Field1           float64 `protobuf:"fixed64,1,opt" json:"Field1"`
	Field2           float32 `protobuf:"fixed32,2,opt" json:"Field2"`
	Field3           int32   `protobuf:"varint,3,opt" json:"Field3"`
	Field4           int64   `protobuf:"varint,4,opt" json:"Field4"`
	Field5           uint32  `protobuf:"varint,5,opt" json:"Field5"`
	Field6           uint64  `protobuf:"varint,6,opt" json:"Field6"`
	Field7           int32   `protobuf:"zigzag32,7,opt" json:"Field7"`
	Field8           int64   `protobuf:"zigzag64,8,opt" json:"Field8"`
	Field9           uint32  `protobuf:"fixed32,9,opt" json:"Field9"`
	Field10          int32   `protobuf:"fixed32,10,opt" json:"Field10"`
	Field11          uint64  `protobuf:"fixed64,11,opt" json:"Field11"`
	Field12          int64   `protobuf:"fixed64,12,opt" json:"Field12"`
	Field13          bool    `protobuf:"varint,13,opt" json:"Field13"`
	Field14          string  `protobuf:"bytes,14,opt" json:"Field14"`
	Field15          []byte  `protobuf:"bytes,15,opt" json:"Field15"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NidOptNative) Reset()      { *m = NidOptNative{} }
func (*NidOptNative) ProtoMessage() {}

type NinOptNative struct {
	Field1           *float64 `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32 `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *int32   `protobuf:"varint,3,opt" json:"Field3,omitempty"`
	Field4           *int64   `protobuf:"varint,4,opt" json:"Field4,omitempty"`
	Field5           *uint32  `protobuf:"varint,5,opt" json:"Field5,omitempty"`
	Field6           *uint64  `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32   `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *int64   `protobuf:"zigzag64,8,opt" json:"Field8,omitempty"`
	Field9           *uint32  `protobuf:"fixed32,9,opt" json:"Field9,omitempty"`
	Field10          *int32   `protobuf:"fixed32,10,opt" json:"Field10,omitempty"`
	Field11          *uint64  `protobuf:"fixed64,11,opt" json:"Field11,omitempty"`
	Field12          *int64   `protobuf:"fixed64,12,opt" json:"Field12,omitempty"`
	Field13          *bool    `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string  `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte   `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *NinOptNative) Reset()      { *m = NinOptNative{} }
func (*NinOptNative) ProtoMessage() {}

type NidRepNative struct {
	Field1           []float64 `protobuf:"fixed64,1,rep" json:"Field1"`
	Field2           []float32 `protobuf:"fixed32,2,rep" json:"Field2"`
	Field3           []int32   `protobuf:"varint,3,rep" json:"Field3"`
	Field4           []int64   `protobuf:"varint,4,rep" json:"Field4"`
	Field5           []uint32  `protobuf:"varint,5,rep" json:"Field5"`
	Field6           []uint64  `protobuf:"varint,6,rep" json:"Field6"`
	Field7           []int32   `protobuf:"zigzag32,7,rep" json:"Field7"`
	Field8           []int64   `protobuf:"zigzag64,8,rep" json:"Field8"`
	Field9           []uint32  `protobuf:"fixed32,9,rep" json:"Field9"`
	Field10          []int32   `protobuf:"fixed32,10,rep" json:"Field10"`
	Field11          []uint64  `protobuf:"fixed64,11,rep" json:"Field11"`
	Field12          []int64   `protobuf:"fixed64,12,rep" json:"Field12"`
	Field13          []bool    `protobuf:"varint,13,rep" json:"Field13"`
	Field14          []string  `protobuf:"bytes,14,rep" json:"Field14"`
	Field15          [][]byte  `protobuf:"bytes,15,rep" json:"Field15"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NidRepNative) Reset()      { *m = NidRepNative{} }
func (*NidRepNative) ProtoMessage() {}

type NinRepNative struct {
	Field1           []float64 `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32 `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []int32   `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4           []int64   `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field5           []uint32  `protobuf:"varint,5,rep" json:"Field5,omitempty"`
	Field6           []uint64  `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32   `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []int64   `protobuf:"zigzag64,8,rep" json:"Field8,omitempty"`
	Field9           []uint32  `protobuf:"fixed32,9,rep" json:"Field9,omitempty"`
	Field10          []int32   `protobuf:"fixed32,10,rep" json:"Field10,omitempty"`
	Field11          []uint64  `protobuf:"fixed64,11,rep" json:"Field11,omitempty"`
	Field12          []int64   `protobuf:"fixed64,12,rep" json:"Field12,omitempty"`
	Field13          []bool    `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string  `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte  `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *NinRepNative) Reset()      { *m = NinRepNative{} }
func (*NinRepNative) ProtoMessage() {}

type NidRepPackedNative struct {
	Field3           []int32 `protobuf:"varint,3,rep,packed" json:"Field3"`
	Field4           []int64 `protobuf:"varint,4,rep,packed" json:"Field4"`
	Field13          []bool  `protobuf:"varint,13,rep,packed" json:"Field13"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NidRepPackedNative) Reset()      { *m = NidRepPackedNative{} }
func (*NidRepPackedNative) ProtoMessage() {}

type NinRepPackedNative struct {
	Field3           []int32 `protobuf:"varint,3,rep" json:"Field3,omitempty"`
	Field4           []int64 `protobuf:"varint,4,rep" json:"Field4,omitempty"`
	Field13          []bool  `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *NinRepPackedNative) Reset()      { *m = NinRepPackedNative{} }
func (*NinRepPackedNative) ProtoMessage() {}

type NidOptStruct struct {
	Field1           float64      `protobuf:"fixed64,1,opt" json:"Field1"`
	Field2           float32      `protobuf:"fixed32,2,opt" json:"Field2"`
	Field3           NidOptNative `protobuf:"bytes,3,opt" json:"Field3"`
	Field4           NinOptNative `protobuf:"bytes,4,opt" json:"Field4"`
	Field6           uint64       `protobuf:"varint,6,opt" json:"Field6"`
	Field7           int32        `protobuf:"zigzag32,7,opt" json:"Field7"`
	Field8           NidOptNative `protobuf:"bytes,8,opt" json:"Field8"`
	Field13          bool         `protobuf:"varint,13,opt" json:"Field13"`
	Field14          string       `protobuf:"bytes,14,opt" json:"Field14"`
	Field15          []byte       `protobuf:"bytes,15,opt" json:"Field15"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NidOptStruct) Reset()      { *m = NidOptStruct{} }
func (*NidOptStruct) ProtoMessage() {}

type NinOptStruct struct {
	Field1           *float64      `protobuf:"fixed64,1,opt" json:"Field1,omitempty"`
	Field2           *float32      `protobuf:"fixed32,2,opt" json:"Field2,omitempty"`
	Field3           *NidOptNative `protobuf:"bytes,3,opt" json:"Field3,omitempty"`
	Field4           *NinOptNative `protobuf:"bytes,4,opt" json:"Field4,omitempty"`
	Field6           *uint64       `protobuf:"varint,6,opt" json:"Field6,omitempty"`
	Field7           *int32        `protobuf:"zigzag32,7,opt" json:"Field7,omitempty"`
	Field8           *NidOptNative `protobuf:"bytes,8,opt" json:"Field8,omitempty"`
	Field13          *bool         `protobuf:"varint,13,opt" json:"Field13,omitempty"`
	Field14          *string       `protobuf:"bytes,14,opt" json:"Field14,omitempty"`
	Field15          []byte        `protobuf:"bytes,15,opt" json:"Field15,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinOptStruct) Reset()      { *m = NinOptStruct{} }
func (*NinOptStruct) ProtoMessage() {}

type NidRepStruct struct {
	Field1           []float64      `protobuf:"fixed64,1,rep" json:"Field1"`
	Field2           []float32      `protobuf:"fixed32,2,rep" json:"Field2"`
	Field3           []NidOptNative `protobuf:"bytes,3,rep" json:"Field3"`
	Field4           []NinOptNative `protobuf:"bytes,4,rep" json:"Field4"`
	Field6           []uint64       `protobuf:"varint,6,rep" json:"Field6"`
	Field7           []int32        `protobuf:"zigzag32,7,rep" json:"Field7"`
	Field8           []NidOptNative `protobuf:"bytes,8,rep" json:"Field8"`
	Field13          []bool         `protobuf:"varint,13,rep" json:"Field13"`
	Field14          []string       `protobuf:"bytes,14,rep" json:"Field14"`
	Field15          [][]byte       `protobuf:"bytes,15,rep" json:"Field15"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *NidRepStruct) Reset()      { *m = NidRepStruct{} }
func (*NidRepStruct) ProtoMessage() {}

type NinRepStruct struct {
	Field1           []float64       `protobuf:"fixed64,1,rep" json:"Field1,omitempty"`
	Field2           []float32       `protobuf:"fixed32,2,rep" json:"Field2,omitempty"`
	Field3           []*NidOptNative `protobuf:"bytes,3,rep" json:"Field3,omitempty"`
	Field4           []*NinOptNative `protobuf:"bytes,4,rep" json:"Field4,omitempty"`
	Field6           []uint64        `protobuf:"varint,6,rep" json:"Field6,omitempty"`
	Field7           []int32         `protobuf:"zigzag32,7,rep" json:"Field7,omitempty"`
	Field8           []*NidOptNative `protobuf:"bytes,8,rep" json:"Field8,omitempty"`
	Field13          []bool          `protobuf:"varint,13,rep" json:"Field13,omitempty"`
	Field14          []string        `protobuf:"bytes,14,rep" json:"Field14,omitempty"`
	Field15          [][]byte        `protobuf:"bytes,15,rep" json:"Field15,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinRepStruct) Reset()      { *m = NinRepStruct{} }
func (*NinRepStruct) ProtoMessage() {}

type NidEmbeddedStruct struct {
	NidOptNative     `protobuf:"bytes,1,opt,embedded=Field1" json:"Field1"`
	Field200         NidOptNative `protobuf:"bytes,200,opt" json:"Field200"`
	Field210         bool         `protobuf:"varint,210,opt" json:"Field210"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *NidEmbeddedStruct) Reset()      { *m = NidEmbeddedStruct{} }
func (*NidEmbeddedStruct) ProtoMessage() {}

type NinEmbeddedStruct struct {
	*NidOptNative    `protobuf:"bytes,1,opt,embedded=Field1" json:"Field1,omitempty"`
	Field200         *NidOptNative `protobuf:"bytes,200,opt" json:"Field200,omitempty"`
	Field210         *bool         `protobuf:"varint,210,opt" json:"Field210,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *NinEmbeddedStruct) Reset()      { *m = NinEmbeddedStruct{} }
func (*NinEmbeddedStruct) ProtoMessage() {}

type NidNestedStruct struct {
	Field1           NidOptStruct   `protobuf:"bytes,1,opt" json:"Field1"`
	Field2           []NidRepStruct `protobuf:"bytes,2,rep" json:"Field2"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *NidNestedStruct) Reset()      { *m = NidNestedStruct{} }
func (*NidNestedStruct) ProtoMessage() {}

type NinNestedStruct struct {
	Field1           *NidOptStruct   `protobuf:"bytes,1,opt" json:"Field1,omitempty"`
	Field2           []*NidRepStruct `protobuf:"bytes,2,rep" json:"Field2,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *NinNestedStruct) Reset()      { *m = NinNestedStruct{} }
func (*NinNestedStruct) ProtoMessage() {}

type NidOptUuidAsBytes struct {
	Id               code_google_com_p_gogoprotobuf_test_custom.Uuid `protobuf:"bytes,1,opt,customtype=code.google.com/p/gogoprotobuf/test/custom.Uuid" json:"Id"`
	XXX_unrecognized []byte                                          `json:"-"`
}

func (m *NidOptUuidAsBytes) Reset()      { *m = NidOptUuidAsBytes{} }
func (*NidOptUuidAsBytes) ProtoMessage() {}

type NinOptUuidAsBytes struct {
	Id               *code_google_com_p_gogoprotobuf_test_custom.Uuid `protobuf:"bytes,1,opt,customtype=code.google.com/p/gogoprotobuf/test/custom.Uuid" json:"Id,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *NinOptUuidAsBytes) Reset()      { *m = NinOptUuidAsBytes{} }
func (*NinOptUuidAsBytes) ProtoMessage() {}

type NidRepUuidAsBytes struct {
	Id               []code_google_com_p_gogoprotobuf_test_custom.Uuid `protobuf:"bytes,1,rep,customtype=code.google.com/p/gogoprotobuf/test/custom.Uuid" json:"Id"`
	XXX_unrecognized []byte                                            `json:"-"`
}

func (m *NidRepUuidAsBytes) Reset()      { *m = NidRepUuidAsBytes{} }
func (*NidRepUuidAsBytes) ProtoMessage() {}

type NinRepUuidAsBytes struct {
	Id               []code_google_com_p_gogoprotobuf_test_custom.Uuid `protobuf:"bytes,1,rep,customtype=code.google.com/p/gogoprotobuf/test/custom.Uuid" json:"Id,omitempty"`
	XXX_unrecognized []byte                                            `json:"-"`
}

func (m *NinRepUuidAsBytes) Reset()      { *m = NinRepUuidAsBytes{} }
func (*NinRepUuidAsBytes) ProtoMessage() {}

type NidOptUint128AsBytes struct {
	Value            code_google_com_p_gogoprotobuf_test_custom.Uint128 `protobuf:"bytes,1,opt,customtype=code.google.com/p/gogoprotobuf/test/custom.Uint128" json:"Value"`
	XXX_unrecognized []byte                                             `json:"-"`
}

func (m *NidOptUint128AsBytes) Reset()      { *m = NidOptUint128AsBytes{} }
func (*NidOptUint128AsBytes) ProtoMessage() {}

type NinOptUint128AsBytes struct {
	Value            *code_google_com_p_gogoprotobuf_test_custom.Uint128 `protobuf:"bytes,1,opt,customtype=code.google.com/p/gogoprotobuf/test/custom.Uint128" json:"Value,omitempty"`
	XXX_unrecognized []byte                                              `json:"-"`
}

func (m *NinOptUint128AsBytes) Reset()      { *m = NinOptUint128AsBytes{} }
func (*NinOptUint128AsBytes) ProtoMessage() {}

type NidRepUint128AsBytes struct {
	Value            []code_google_com_p_gogoprotobuf_test_custom.Uint128 `protobuf:"bytes,1,rep,customtype=code.google.com/p/gogoprotobuf/test/custom.Uint128" json:"Value"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *NidRepUint128AsBytes) Reset()      { *m = NidRepUint128AsBytes{} }
func (*NidRepUint128AsBytes) ProtoMessage() {}

type NinRepUint128AsBytes struct {
	Value            []code_google_com_p_gogoprotobuf_test_custom.Uint128 `protobuf:"bytes,1,rep,customtype=code.google.com/p/gogoprotobuf/test/custom.Uint128" json:"Value,omitempty"`
	XXX_unrecognized []byte                                               `json:"-"`
}

func (m *NinRepUint128AsBytes) Reset()      { *m = NinRepUint128AsBytes{} }
func (*NinRepUint128AsBytes) ProtoMessage() {}

func init() {
}
func (this *SimpleMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SimpleMessage{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`SimpleName:` + fmt.Sprintf("%v", this.SimpleName) + `,`,
		`Time:` + valueToStringThetest(this.Time) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidOptNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidOptNative{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + fmt.Sprintf("%v", this.Field3) + `,`,
		`Field4:` + fmt.Sprintf("%v", this.Field4) + `,`,
		`Field5:` + fmt.Sprintf("%v", this.Field5) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + fmt.Sprintf("%v", this.Field8) + `,`,
		`Field9:` + fmt.Sprintf("%v", this.Field9) + `,`,
		`Field10:` + fmt.Sprintf("%v", this.Field10) + `,`,
		`Field11:` + fmt.Sprintf("%v", this.Field11) + `,`,
		`Field12:` + fmt.Sprintf("%v", this.Field12) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinOptNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinOptNative{`,
		`Field1:` + valueToStringThetest(this.Field1) + `,`,
		`Field2:` + valueToStringThetest(this.Field2) + `,`,
		`Field3:` + valueToStringThetest(this.Field3) + `,`,
		`Field4:` + valueToStringThetest(this.Field4) + `,`,
		`Field5:` + valueToStringThetest(this.Field5) + `,`,
		`Field6:` + valueToStringThetest(this.Field6) + `,`,
		`Field7:` + valueToStringThetest(this.Field7) + `,`,
		`Field8:` + valueToStringThetest(this.Field8) + `,`,
		`Field9:` + valueToStringThetest(this.Field9) + `,`,
		`Field10:` + valueToStringThetest(this.Field10) + `,`,
		`Field11:` + valueToStringThetest(this.Field11) + `,`,
		`Field12:` + valueToStringThetest(this.Field12) + `,`,
		`Field13:` + valueToStringThetest(this.Field13) + `,`,
		`Field14:` + valueToStringThetest(this.Field14) + `,`,
		`Field15:` + valueToStringThetest(this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidRepNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidRepNative{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + fmt.Sprintf("%v", this.Field3) + `,`,
		`Field4:` + fmt.Sprintf("%v", this.Field4) + `,`,
		`Field5:` + fmt.Sprintf("%v", this.Field5) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + fmt.Sprintf("%v", this.Field8) + `,`,
		`Field9:` + fmt.Sprintf("%v", this.Field9) + `,`,
		`Field10:` + fmt.Sprintf("%v", this.Field10) + `,`,
		`Field11:` + fmt.Sprintf("%v", this.Field11) + `,`,
		`Field12:` + fmt.Sprintf("%v", this.Field12) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinRepNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinRepNative{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + fmt.Sprintf("%v", this.Field3) + `,`,
		`Field4:` + fmt.Sprintf("%v", this.Field4) + `,`,
		`Field5:` + fmt.Sprintf("%v", this.Field5) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + fmt.Sprintf("%v", this.Field8) + `,`,
		`Field9:` + fmt.Sprintf("%v", this.Field9) + `,`,
		`Field10:` + fmt.Sprintf("%v", this.Field10) + `,`,
		`Field11:` + fmt.Sprintf("%v", this.Field11) + `,`,
		`Field12:` + fmt.Sprintf("%v", this.Field12) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidRepPackedNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidRepPackedNative{`,
		`Field3:` + fmt.Sprintf("%v", this.Field3) + `,`,
		`Field4:` + fmt.Sprintf("%v", this.Field4) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinRepPackedNative) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinRepPackedNative{`,
		`Field3:` + fmt.Sprintf("%v", this.Field3) + `,`,
		`Field4:` + fmt.Sprintf("%v", this.Field4) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidOptStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidOptStruct{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + strings.Replace(strings.Replace(this.Field3.String(), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field4:` + strings.Replace(strings.Replace(this.Field4.String(), "NinOptNative", "NinOptNative", 1), `&`, ``, 1) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + strings.Replace(strings.Replace(this.Field8.String(), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinOptStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinOptStruct{`,
		`Field1:` + valueToStringThetest(this.Field1) + `,`,
		`Field2:` + valueToStringThetest(this.Field2) + `,`,
		`Field3:` + strings.Replace(fmt.Sprintf("%v", this.Field3), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field4:` + strings.Replace(fmt.Sprintf("%v", this.Field4), "NinOptNative", "NinOptNative", 1) + `,`,
		`Field6:` + valueToStringThetest(this.Field6) + `,`,
		`Field7:` + valueToStringThetest(this.Field7) + `,`,
		`Field8:` + strings.Replace(fmt.Sprintf("%v", this.Field8), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field13:` + valueToStringThetest(this.Field13) + `,`,
		`Field14:` + valueToStringThetest(this.Field14) + `,`,
		`Field15:` + valueToStringThetest(this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidRepStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidRepStruct{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Field3), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field4:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Field4), "NinOptNative", "NinOptNative", 1), `&`, ``, 1) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Field8), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinRepStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinRepStruct{`,
		`Field1:` + fmt.Sprintf("%v", this.Field1) + `,`,
		`Field2:` + fmt.Sprintf("%v", this.Field2) + `,`,
		`Field3:` + strings.Replace(fmt.Sprintf("%v", this.Field3), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field4:` + strings.Replace(fmt.Sprintf("%v", this.Field4), "NinOptNative", "NinOptNative", 1) + `,`,
		`Field6:` + fmt.Sprintf("%v", this.Field6) + `,`,
		`Field7:` + fmt.Sprintf("%v", this.Field7) + `,`,
		`Field8:` + strings.Replace(fmt.Sprintf("%v", this.Field8), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field13:` + fmt.Sprintf("%v", this.Field13) + `,`,
		`Field14:` + fmt.Sprintf("%v", this.Field14) + `,`,
		`Field15:` + fmt.Sprintf("%v", this.Field15) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidEmbeddedStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidEmbeddedStruct{`,
		`NidOptNative:` + strings.Replace(strings.Replace(this.NidOptNative.String(), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field200:` + strings.Replace(strings.Replace(this.Field200.String(), "NidOptNative", "NidOptNative", 1), `&`, ``, 1) + `,`,
		`Field210:` + fmt.Sprintf("%v", this.Field210) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinEmbeddedStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinEmbeddedStruct{`,
		`NidOptNative:` + strings.Replace(fmt.Sprintf("%v", this.NidOptNative), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field200:` + strings.Replace(fmt.Sprintf("%v", this.Field200), "NidOptNative", "NidOptNative", 1) + `,`,
		`Field210:` + valueToStringThetest(this.Field210) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidNestedStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidNestedStruct{`,
		`Field1:` + strings.Replace(strings.Replace(this.Field1.String(), "NidOptStruct", "NidOptStruct", 1), `&`, ``, 1) + `,`,
		`Field2:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Field2), "NidRepStruct", "NidRepStruct", 1), `&`, ``, 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinNestedStruct) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinNestedStruct{`,
		`Field1:` + strings.Replace(fmt.Sprintf("%v", this.Field1), "NidOptStruct", "NidOptStruct", 1) + `,`,
		`Field2:` + strings.Replace(fmt.Sprintf("%v", this.Field2), "NidRepStruct", "NidRepStruct", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidOptUuidAsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidOptUuidAsBytes{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinOptUuidAsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinOptUuidAsBytes{`,
		`Id:` + valueToStringThetest(this.Id) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidRepUuidAsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidRepUuidAsBytes{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinRepUuidAsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinRepUuidAsBytes{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidOptUint128AsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidOptUint128AsBytes{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinOptUint128AsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinOptUint128AsBytes{`,
		`Value:` + valueToStringThetest(this.Value) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NidRepUint128AsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NidRepUint128AsBytes{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NinRepUint128AsBytes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NinRepUint128AsBytes{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringThetest(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
