// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/fault/v2/fault.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
import v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v2"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FaultAbort struct {
	// An integer between 0-100 indicating the percentage of requests/operations/connections
	// that will be aborted with the error code provided.
	Percent uint32 `protobuf:"varint,1,opt,name=percent,proto3" json:"percent,omitempty"`
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType            isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_fault_51fe2e769f9ff448, []int{0}
}
func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(dst, src)
}
func (m *FaultAbort) XXX_Size() int {
	return m.Size()
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultAbort_OneofMarshaler, _FaultAbort_OneofUnmarshaler, _FaultAbort_OneofSizer, []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

func _FaultAbort_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		_ = b.EncodeVarint(2<<3 | proto.WireVarint)
		_ = b.EncodeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		return fmt.Errorf("FaultAbort.ErrorType has unexpected type %T", x)
	}
	return nil
}

func _FaultAbort_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultAbort)
	switch tag {
	case 2: // error_type.http_status
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.ErrorType = &FaultAbort_HttpStatus{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _FaultAbort_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultAbort)
	// error_type
	switch x := m.ErrorType.(type) {
	case *FaultAbort_HttpStatus:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.HttpStatus))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object. At least *abort* or *delay* must be specified.
	Delay *v2.FaultDelay `protobuf:"bytes,1,opt,name=delay" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percent
	// <envoy_api_field_config.filter.http.fault.v2.FaultAbort.percent>` field. The filter will
	// check the request's headers against all the specified headers in the filter
	// config. A match will happen if all the headers in the config are present in
	// the request with the same values (or based on presence if the *value* field
	// is not in the config).
	Headers []*route.HeaderMatcher `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes      []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes" json:"downstream_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_fault_51fe2e769f9ff448, []int{1}
}
func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(dst, src)
}
func (m *HTTPFault) XXX_Size() int {
	return m.Size()
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v2.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.config.filter.http.fault.v2.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.config.filter.http.fault.v2.HTTPFault")
}
func (m *FaultAbort) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FaultAbort) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Percent != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Percent))
	}
	if m.ErrorType != nil {
		nn1, err := m.ErrorType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *FaultAbort_HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x10
	i++
	i = encodeVarintFault(dAtA, i, uint64(m.HttpStatus))
	return i, nil
}
func (m *HTTPFault) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPFault) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Delay != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Delay.Size()))
		n2, err := m.Delay.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Abort != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintFault(dAtA, i, uint64(m.Abort.Size()))
		n3, err := m.Abort.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.UpstreamCluster) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFault(dAtA, i, uint64(len(m.UpstreamCluster)))
		i += copy(dAtA[i:], m.UpstreamCluster)
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x22
			i++
			i = encodeVarintFault(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFault(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FaultAbort) Size() (n int) {
	var l int
	_ = l
	if m.Percent != 0 {
		n += 1 + sovFault(uint64(m.Percent))
	}
	if m.ErrorType != nil {
		n += m.ErrorType.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FaultAbort_HttpStatus) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovFault(uint64(m.HttpStatus))
	return n
}
func (m *HTTPFault) Size() (n int) {
	var l int
	_ = l
	if m.Delay != nil {
		l = m.Delay.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	if m.Abort != nil {
		l = m.Abort.Size()
		n += 1 + l + sovFault(uint64(l))
	}
	l = len(m.UpstreamCluster)
	if l > 0 {
		n += 1 + l + sovFault(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if len(m.DownstreamNodes) > 0 {
		for _, s := range m.DownstreamNodes {
			l = len(s)
			n += 1 + l + sovFault(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFault(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozFault(x uint64) (n int) {
	return sovFault(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FaultAbort) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FaultAbort: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FaultAbort: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percent", wireType)
			}
			m.Percent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Percent |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpStatus", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ErrorType = &FaultAbort_HttpStatus{v}
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HTTPFault) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFault
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPFault: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPFault: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Delay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Delay == nil {
				m.Delay = &v2.FaultDelay{}
			}
			if err := m.Delay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Abort == nil {
				m.Abort = &FaultAbort{}
			}
			if err := m.Abort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpstreamCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UpstreamCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &route.HeaderMatcher{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DownstreamNodes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFault
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthFault
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DownstreamNodes = append(m.DownstreamNodes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFault(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFault
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFault(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFault
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
					return 0, ErrIntOverflowFault
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFault
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthFault
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFault
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFault(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFault = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFault   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/fault/v2/fault.proto", fileDescriptor_fault_51fe2e769f9ff448)
}

var fileDescriptor_fault_51fe2e769f9ff448 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x8b, 0xdb, 0x30,
	0x1c, 0xc5, 0x2b, 0x3b, 0x4e, 0x1a, 0x99, 0xd0, 0xa0, 0x0e, 0x35, 0x19, 0x8c, 0x93, 0x2e, 0x6e,
	0x20, 0x72, 0x71, 0xc7, 0x42, 0xa1, 0x49, 0x29, 0x59, 0x5a, 0x8a, 0x9b, 0xe9, 0x96, 0xa0, 0xd8,
	0xca, 0xc5, 0xe0, 0xb3, 0x8c, 0x2c, 0xfb, 0x2e, 0x9f, 0xe7, 0xbe, 0xc4, 0x71, 0x53, 0xc6, 0x8c,
	0x37, 0xde, 0x78, 0x64, 0xcb, 0xb7, 0x38, 0x24, 0xc5, 0x64, 0x09, 0xdc, 0x62, 0xc4, 0xff, 0xff,
	0x7b, 0xef, 0x49, 0xcf, 0x70, 0x42, 0xf3, 0x9a, 0x6d, 0x83, 0x98, 0xe5, 0xeb, 0xf4, 0x3a, 0x58,
	0xa7, 0x99, 0xa0, 0x3c, 0xd8, 0x08, 0x51, 0x04, 0x6b, 0x52, 0x65, 0x22, 0xa8, 0x43, 0x7d, 0xc0,
	0x05, 0x67, 0x82, 0xa1, 0xa1, 0xc2, 0xb1, 0xc6, 0xb1, 0xc6, 0xb1, 0xc4, 0xb1, 0xa6, 0xea, 0x70,
	0xe0, 0x5f, 0x72, 0xbc, 0x64, 0x36, 0x70, 0x35, 0x49, 0x8a, 0x54, 0x6e, 0x38, 0xab, 0x04, 0xd5,
	0xdf, 0xd3, 0xfe, 0x53, 0x4d, 0xb2, 0x34, 0x21, 0x82, 0x06, 0xcd, 0x41, 0x2f, 0x46, 0x77, 0x10,
	0xfe, 0x96, 0x3e, 0x3f, 0x57, 0x8c, 0x0b, 0xf4, 0x19, 0x76, 0x0a, 0xca, 0x63, 0x9a, 0x0b, 0x07,
	0x78, 0xc0, 0xef, 0x4d, 0xbb, 0x8f, 0xc7, 0x9d, 0xd9, 0x1a, 0x1b, 0x4e, 0x12, 0x35, 0x1b, 0xf4,
	0x15, 0xda, 0xf2, 0x9a, 0xcb, 0x52, 0x10, 0x51, 0x95, 0x8e, 0xa1, 0xc0, 0x9e, 0x04, 0xdf, 0x8f,
	0xdb, 0xfd, 0xe7, 0x96, 0xbf, 0x07, 0xf3, 0x77, 0x11, 0x94, 0xcc, 0x7f, 0x85, 0x4c, 0x3f, 0x42,
	0x48, 0x39, 0x67, 0x7c, 0x29, 0xb6, 0x05, 0x45, 0xd6, 0xc3, 0x71, 0x67, 0x82, 0xd1, 0xbd, 0x01,
	0xbb, 0xf3, 0xc5, 0xe2, 0x9f, 0x8a, 0x47, 0x3f, 0xa0, 0x95, 0xd0, 0x8c, 0x6c, 0x55, 0xae, 0x1d,
	0xfa, 0xf8, 0x52, 0x3b, 0x4d, 0x31, 0x58, 0x69, 0x7e, 0x49, 0x3e, 0xd2, 0x32, 0x34, 0x83, 0x16,
	0x91, 0x4f, 0x50, 0xd7, 0xb1, 0xc3, 0x09, 0x7e, 0xb3, 0x5d, 0x7c, 0x7e, 0x77, 0xa4, 0xb5, 0xe8,
	0x0b, 0xec, 0x57, 0x45, 0x29, 0x38, 0x25, 0x37, 0xcb, 0x38, 0xab, 0x4a, 0x41, 0xb9, 0x63, 0x7a,
	0xc0, 0xef, 0x46, 0x1f, 0x9a, 0xf9, 0x4c, 0x8f, 0xd1, 0x77, 0xd8, 0xd9, 0x50, 0x92, 0x50, 0x5e,
	0x3a, 0x2d, 0xcf, 0xf4, 0xed, 0x70, 0x78, 0x4a, 0x24, 0x45, 0x2a, 0xcd, 0x75, 0xf9, 0x73, 0x85,
	0xfc, 0x21, 0x22, 0xde, 0x50, 0x1e, 0x35, 0x0a, 0x99, 0x93, 0xb0, 0xdb, 0xfc, 0x94, 0x94, 0xb3,
	0x84, 0x96, 0x8e, 0xe5, 0x99, 0x32, 0xe7, 0x3c, 0xff, 0x2b, 0xc7, 0xd3, 0xfe, 0xfe, 0xe0, 0x82,
	0xa7, 0x83, 0x0b, 0x5e, 0x0e, 0x2e, 0xb8, 0x32, 0xea, 0x70, 0xd5, 0x56, 0x3f, 0xee, 0xdb, 0x6b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xb1, 0x65, 0xae, 0x6f, 0x02, 0x00, 0x00,
}