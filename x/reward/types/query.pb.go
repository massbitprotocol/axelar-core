// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/reward/v1beta1/query.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// InflationRateRequest represents a message that queries the Axelar specific
// inflation RPC method.
type InflationRateRequest struct {
}

func (m *InflationRateRequest) Reset()         { *m = InflationRateRequest{} }
func (m *InflationRateRequest) String() string { return proto.CompactTextString(m) }
func (*InflationRateRequest) ProtoMessage()    {}
func (*InflationRateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea20e5bdb695fbb5, []int{0}
}
func (m *InflationRateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationRateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationRateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationRateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationRateRequest.Merge(m, src)
}
func (m *InflationRateRequest) XXX_Size() int {
	return m.Size()
}
func (m *InflationRateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationRateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InflationRateRequest proto.InternalMessageInfo

type InflationRateResponse struct {
	InflationRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation_rate,json=inflationRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate"`
}

func (m *InflationRateResponse) Reset()         { *m = InflationRateResponse{} }
func (m *InflationRateResponse) String() string { return proto.CompactTextString(m) }
func (*InflationRateResponse) ProtoMessage()    {}
func (*InflationRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea20e5bdb695fbb5, []int{1}
}
func (m *InflationRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationRateResponse.Merge(m, src)
}
func (m *InflationRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *InflationRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InflationRateResponse proto.InternalMessageInfo

// ParamsRequest represents a message that queries the params
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea20e5bdb695fbb5, []int{2}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

type ParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea20e5bdb695fbb5, []int{3}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InflationRateRequest)(nil), "axelar.reward.v1beta1.InflationRateRequest")
	proto.RegisterType((*InflationRateResponse)(nil), "axelar.reward.v1beta1.InflationRateResponse")
	proto.RegisterType((*ParamsRequest)(nil), "axelar.reward.v1beta1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "axelar.reward.v1beta1.ParamsResponse")
}

func init() { proto.RegisterFile("axelar/reward/v1beta1/query.proto", fileDescriptor_ea20e5bdb695fbb5) }

var fileDescriptor_ea20e5bdb695fbb5 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xac, 0x48, 0xcd,
	0x49, 0x2c, 0xd2, 0x2f, 0x4a, 0x2d, 0x4f, 0x2c, 0x4a, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49,
	0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x85,
	0x28, 0xd1, 0x83, 0x28, 0xd1, 0x83, 0x2a, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0,
	0x07, 0xb1, 0x20, 0x8a, 0xa5, 0x94, 0xb0, 0x9b, 0x57, 0x90, 0x58, 0x94, 0x98, 0x5b, 0x0c, 0x51,
	0xa3, 0x24, 0xc6, 0x25, 0xe2, 0x99, 0x97, 0x96, 0x93, 0x58, 0x92, 0x99, 0x9f, 0x17, 0x94, 0x58,
	0x92, 0x1a, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0xa2, 0x94, 0xc7, 0x25, 0x8a, 0x26, 0x5e, 0x5c,
	0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x14, 0xca, 0xc5, 0x97, 0x09, 0x93, 0x88, 0x2f, 0x4a, 0x2c, 0x49,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x71, 0xd2, 0x3b, 0x71, 0x4f, 0x9e, 0xe1, 0xd6, 0x3d, 0x79,
	0xb5, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xe4, 0xfc, 0xe2, 0xdc,
	0xfc, 0x62, 0x28, 0xa5, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x5a, 0xac, 0xe7, 0x92,
	0x9a, 0x1c, 0xc4, 0x9b, 0x89, 0x6c, 0xbc, 0x12, 0x3f, 0x17, 0x6f, 0x00, 0xd8, 0x5d, 0x30, 0x07,
	0xf8, 0x72, 0xf1, 0xc1, 0x04, 0xa0, 0x36, 0x5b, 0x73, 0xb1, 0x41, 0x9c, 0x0e, 0xb6, 0x91, 0xdb,
	0x48, 0x56, 0x0f, 0x6b, 0x60, 0xe8, 0x41, 0xb4, 0x39, 0xb1, 0x80, 0x1c, 0x14, 0x04, 0xd5, 0xe2,
	0x14, 0x78, 0xe2, 0xa1, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0x19, 0x23, 0x39, 0x1a, 0x62, 0x68, 0x5e, 0x6a, 0x49, 0x79, 0x7e, 0x51, 0x36, 0x94, 0xa7, 0x9b,
	0x9c, 0x5f, 0x94, 0xaa, 0x5f, 0x01, 0x0b, 0x49, 0xb0, 0x2f, 0x92, 0xd8, 0xc0, 0x21, 0x68, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xca, 0x0e, 0x12, 0xb7, 0x01, 0x00, 0x00,
}

func (m *InflationRateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationRateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationRateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *InflationRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InflationRate.Size()
		i -= size
		if _, err := m.InflationRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationRateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *InflationRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InflationRate.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationRateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InflationRateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationRateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InflationRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InflationRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
