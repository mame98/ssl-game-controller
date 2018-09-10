// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_controller_auto_ref.proto

package refproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AutoRefToControllerRequest_State int32

const (
	// unknown or not set
	AutoRefToControllerRequest_UNKNOWN AutoRefToControllerRequest_State = 0
	// the game can continue (all conditions met: Ball is placed and bots are located at legal positions)
	AutoRefToControllerRequest_READY_TO_CONTINUE AutoRefToControllerRequest_State = 1
	// waiting for the ball to be placed
	AutoRefToControllerRequest_WAIT_FOR_PLACEMENT AutoRefToControllerRequest_State = 2
	// waiting for the bots to move to valid positions
	AutoRefToControllerRequest_WAIT_FOR_VALID_POSITIONS AutoRefToControllerRequest_State = 3
)

var AutoRefToControllerRequest_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "READY_TO_CONTINUE",
	2: "WAIT_FOR_PLACEMENT",
	3: "WAIT_FOR_VALID_POSITIONS",
}
var AutoRefToControllerRequest_State_value = map[string]int32{
	"UNKNOWN":                  0,
	"READY_TO_CONTINUE":        1,
	"WAIT_FOR_PLACEMENT":       2,
	"WAIT_FOR_VALID_POSITIONS": 3,
}

func (x AutoRefToControllerRequest_State) Enum() *AutoRefToControllerRequest_State {
	p := new(AutoRefToControllerRequest_State)
	*p = x
	return p
}
func (x AutoRefToControllerRequest_State) String() string {
	return proto.EnumName(AutoRefToControllerRequest_State_name, int32(x))
}
func (x *AutoRefToControllerRequest_State) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AutoRefToControllerRequest_State_value, data, "AutoRefToControllerRequest_State")
	if err != nil {
		return err
	}
	*x = AutoRefToControllerRequest_State(value)
	return nil
}
func (AutoRefToControllerRequest_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{1, 0}
}

// AutoRefRegistration is the first message that a client must send to the controller to identify itself
type AutoRefRegistration struct {
	// identifier is a unique name of the client
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature            *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AutoRefRegistration) Reset()         { *m = AutoRefRegistration{} }
func (m *AutoRefRegistration) String() string { return proto.CompactTextString(m) }
func (*AutoRefRegistration) ProtoMessage()    {}
func (*AutoRefRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{0}
}
func (m *AutoRefRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefRegistration.Unmarshal(m, b)
}
func (m *AutoRefRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefRegistration.Marshal(b, m, deterministic)
}
func (dst *AutoRefRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefRegistration.Merge(dst, src)
}
func (m *AutoRefRegistration) XXX_Size() int {
	return xxx_messageInfo_AutoRefRegistration.Size(m)
}
func (m *AutoRefRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefRegistration proto.InternalMessageInfo

func (m *AutoRefRegistration) GetIdentifier() string {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return ""
}

func (m *AutoRefRegistration) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// AutoRefToControllerRequest is the wrapper message for all subsequent requests from the client to the controller
type AutoRefToControllerRequest struct {
	// game_event is an optional event that the autoRef detected during the game
	GameEvent *GameEvent `protobuf:"bytes,1,opt,name=game_event,json=gameEvent" json:"game_event,omitempty"`
	// auto_ref_message is an optional message that describes the current state or situation of the game/autoRef
	AutoRefMessage *AutoRefMessage `protobuf:"bytes,2,opt,name=auto_ref_message,json=autoRefMessage" json:"auto_ref_message,omitempty"`
	// the current state of the autoRef
	State *AutoRefToControllerRequest_State `protobuf:"varint,3,opt,name=state,enum=AutoRefToControllerRequest_State" json:"state,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature            *Signature `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AutoRefToControllerRequest) Reset()         { *m = AutoRefToControllerRequest{} }
func (m *AutoRefToControllerRequest) String() string { return proto.CompactTextString(m) }
func (*AutoRefToControllerRequest) ProtoMessage()    {}
func (*AutoRefToControllerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{1}
}
func (m *AutoRefToControllerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefToControllerRequest.Unmarshal(m, b)
}
func (m *AutoRefToControllerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefToControllerRequest.Marshal(b, m, deterministic)
}
func (dst *AutoRefToControllerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefToControllerRequest.Merge(dst, src)
}
func (m *AutoRefToControllerRequest) XXX_Size() int {
	return xxx_messageInfo_AutoRefToControllerRequest.Size(m)
}
func (m *AutoRefToControllerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefToControllerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefToControllerRequest proto.InternalMessageInfo

func (m *AutoRefToControllerRequest) GetGameEvent() *GameEvent {
	if m != nil {
		return m.GameEvent
	}
	return nil
}

func (m *AutoRefToControllerRequest) GetAutoRefMessage() *AutoRefMessage {
	if m != nil {
		return m.AutoRefMessage
	}
	return nil
}

func (m *AutoRefToControllerRequest) GetState() AutoRefToControllerRequest_State {
	if m != nil && m.State != nil {
		return *m.State
	}
	return AutoRefToControllerRequest_UNKNOWN
}

func (m *AutoRefToControllerRequest) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

// a message from autoRef, describing the current state or situation
type AutoRefMessage struct {
	// Types that are valid to be assigned to Message:
	//	*AutoRefMessage_Custom
	//	*AutoRefMessage_BallPlaced_
	//	*AutoRefMessage_WaitForBots_
	Message              isAutoRefMessage_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *AutoRefMessage) Reset()         { *m = AutoRefMessage{} }
func (m *AutoRefMessage) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage) ProtoMessage()    {}
func (*AutoRefMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2}
}
func (m *AutoRefMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage.Unmarshal(m, b)
}
func (m *AutoRefMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage.Marshal(b, m, deterministic)
}
func (dst *AutoRefMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage.Merge(dst, src)
}
func (m *AutoRefMessage) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage.Size(m)
}
func (m *AutoRefMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage proto.InternalMessageInfo

type isAutoRefMessage_Message interface {
	isAutoRefMessage_Message()
}

type AutoRefMessage_Custom struct {
	Custom string `protobuf:"bytes,1,opt,name=custom,oneof"`
}

type AutoRefMessage_BallPlaced_ struct {
	BallPlaced *AutoRefMessage_BallPlaced `protobuf:"bytes,2,opt,name=ball_placed,json=ballPlaced,oneof"`
}

type AutoRefMessage_WaitForBots_ struct {
	WaitForBots *AutoRefMessage_WaitForBots `protobuf:"bytes,3,opt,name=wait_for_bots,json=waitForBots,oneof"`
}

func (*AutoRefMessage_Custom) isAutoRefMessage_Message() {}

func (*AutoRefMessage_BallPlaced_) isAutoRefMessage_Message() {}

func (*AutoRefMessage_WaitForBots_) isAutoRefMessage_Message() {}

func (m *AutoRefMessage) GetMessage() isAutoRefMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *AutoRefMessage) GetCustom() string {
	if x, ok := m.GetMessage().(*AutoRefMessage_Custom); ok {
		return x.Custom
	}
	return ""
}

func (m *AutoRefMessage) GetBallPlaced() *AutoRefMessage_BallPlaced {
	if x, ok := m.GetMessage().(*AutoRefMessage_BallPlaced_); ok {
		return x.BallPlaced
	}
	return nil
}

func (m *AutoRefMessage) GetWaitForBots() *AutoRefMessage_WaitForBots {
	if x, ok := m.GetMessage().(*AutoRefMessage_WaitForBots_); ok {
		return x.WaitForBots
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AutoRefMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AutoRefMessage_OneofMarshaler, _AutoRefMessage_OneofUnmarshaler, _AutoRefMessage_OneofSizer, []interface{}{
		(*AutoRefMessage_Custom)(nil),
		(*AutoRefMessage_BallPlaced_)(nil),
		(*AutoRefMessage_WaitForBots_)(nil),
	}
}

func _AutoRefMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AutoRefMessage)
	// message
	switch x := m.Message.(type) {
	case *AutoRefMessage_Custom:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Custom)
	case *AutoRefMessage_BallPlaced_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BallPlaced); err != nil {
			return err
		}
	case *AutoRefMessage_WaitForBots_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WaitForBots); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AutoRefMessage.Message has unexpected type %T", x)
	}
	return nil
}

func _AutoRefMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AutoRefMessage)
	switch tag {
	case 1: // message.custom
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Message = &AutoRefMessage_Custom{x}
		return true, err
	case 2: // message.ball_placed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutoRefMessage_BallPlaced)
		err := b.DecodeMessage(msg)
		m.Message = &AutoRefMessage_BallPlaced_{msg}
		return true, err
	case 3: // message.wait_for_bots
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AutoRefMessage_WaitForBots)
		err := b.DecodeMessage(msg)
		m.Message = &AutoRefMessage_WaitForBots_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AutoRefMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AutoRefMessage)
	// message
	switch x := m.Message.(type) {
	case *AutoRefMessage_Custom:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Custom)))
		n += len(x.Custom)
	case *AutoRefMessage_BallPlaced_:
		s := proto.Size(x.BallPlaced)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AutoRefMessage_WaitForBots_:
		s := proto.Size(x.WaitForBots)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// the result of the ball placement
type AutoRefMessage_BallPlaced struct {
	// the time taken for placing the ball
	TimeTaken *float32 `protobuf:"fixed32,1,req,name=time_taken,json=timeTaken" json:"time_taken,omitempty"`
	// the distance between placement location and actual ball position
	Precision *float32 `protobuf:"fixed32,2,req,name=precision" json:"precision,omitempty"`
	// the distance between the initial ball location and the placement position
	Distance             *float32 `protobuf:"fixed32,3,req,name=distance" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoRefMessage_BallPlaced) Reset()         { *m = AutoRefMessage_BallPlaced{} }
func (m *AutoRefMessage_BallPlaced) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage_BallPlaced) ProtoMessage()    {}
func (*AutoRefMessage_BallPlaced) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2, 0}
}
func (m *AutoRefMessage_BallPlaced) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage_BallPlaced.Unmarshal(m, b)
}
func (m *AutoRefMessage_BallPlaced) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage_BallPlaced.Marshal(b, m, deterministic)
}
func (dst *AutoRefMessage_BallPlaced) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage_BallPlaced.Merge(dst, src)
}
func (m *AutoRefMessage_BallPlaced) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage_BallPlaced.Size(m)
}
func (m *AutoRefMessage_BallPlaced) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage_BallPlaced.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage_BallPlaced proto.InternalMessageInfo

func (m *AutoRefMessage_BallPlaced) GetTimeTaken() float32 {
	if m != nil && m.TimeTaken != nil {
		return *m.TimeTaken
	}
	return 0
}

func (m *AutoRefMessage_BallPlaced) GetPrecision() float32 {
	if m != nil && m.Precision != nil {
		return *m.Precision
	}
	return 0
}

func (m *AutoRefMessage_BallPlaced) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

// the bots that is waited for
type AutoRefMessage_WaitForBots struct {
	// the bots that are waited for
	Violators            []*AutoRefMessage_WaitForBots_Violator `protobuf:"bytes,1,rep,name=violators" json:"violators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *AutoRefMessage_WaitForBots) Reset()         { *m = AutoRefMessage_WaitForBots{} }
func (m *AutoRefMessage_WaitForBots) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage_WaitForBots) ProtoMessage()    {}
func (*AutoRefMessage_WaitForBots) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2, 1}
}
func (m *AutoRefMessage_WaitForBots) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Unmarshal(m, b)
}
func (m *AutoRefMessage_WaitForBots) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Marshal(b, m, deterministic)
}
func (dst *AutoRefMessage_WaitForBots) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage_WaitForBots.Merge(dst, src)
}
func (m *AutoRefMessage_WaitForBots) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage_WaitForBots.Size(m)
}
func (m *AutoRefMessage_WaitForBots) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage_WaitForBots.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage_WaitForBots proto.InternalMessageInfo

func (m *AutoRefMessage_WaitForBots) GetViolators() []*AutoRefMessage_WaitForBots_Violator {
	if m != nil {
		return m.Violators
	}
	return nil
}

type AutoRefMessage_WaitForBots_Violator struct {
	// the id of the violator
	BotId *BotId `protobuf:"bytes,1,req,name=bot_id,json=botId" json:"bot_id,omitempty"`
	// the distance to the next valid position
	Distance             *float32 `protobuf:"fixed32,2,req,name=distance" json:"distance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoRefMessage_WaitForBots_Violator) Reset()         { *m = AutoRefMessage_WaitForBots_Violator{} }
func (m *AutoRefMessage_WaitForBots_Violator) String() string { return proto.CompactTextString(m) }
func (*AutoRefMessage_WaitForBots_Violator) ProtoMessage()    {}
func (*AutoRefMessage_WaitForBots_Violator) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fbb3ba3bab9727c, []int{2, 1, 0}
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Unmarshal(m, b)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Marshal(b, m, deterministic)
}
func (dst *AutoRefMessage_WaitForBots_Violator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Merge(dst, src)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_Size() int {
	return xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.Size(m)
}
func (m *AutoRefMessage_WaitForBots_Violator) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoRefMessage_WaitForBots_Violator.DiscardUnknown(m)
}

var xxx_messageInfo_AutoRefMessage_WaitForBots_Violator proto.InternalMessageInfo

func (m *AutoRefMessage_WaitForBots_Violator) GetBotId() *BotId {
	if m != nil {
		return m.BotId
	}
	return nil
}

func (m *AutoRefMessage_WaitForBots_Violator) GetDistance() float32 {
	if m != nil && m.Distance != nil {
		return *m.Distance
	}
	return 0
}

func init() {
	proto.RegisterType((*AutoRefRegistration)(nil), "AutoRefRegistration")
	proto.RegisterType((*AutoRefToControllerRequest)(nil), "AutoRefToControllerRequest")
	proto.RegisterType((*AutoRefMessage)(nil), "AutoRefMessage")
	proto.RegisterType((*AutoRefMessage_BallPlaced)(nil), "AutoRefMessage.BallPlaced")
	proto.RegisterType((*AutoRefMessage_WaitForBots)(nil), "AutoRefMessage.WaitForBots")
	proto.RegisterType((*AutoRefMessage_WaitForBots_Violator)(nil), "AutoRefMessage.WaitForBots.Violator")
	proto.RegisterEnum("AutoRefToControllerRequest_State", AutoRefToControllerRequest_State_name, AutoRefToControllerRequest_State_value)
}

func init() { proto.RegisterFile("ssl_game_controller_auto_ref.proto", fileDescriptor_2fbb3ba3bab9727c) }

var fileDescriptor_2fbb3ba3bab9727c = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x63, 0x87, 0xfe, 0xf1, 0x58, 0x94, 0xb0, 0x08, 0x64, 0x4c, 0x8b, 0x42, 0xc4, 0x21,
	0x5c, 0x2c, 0xf0, 0x05, 0xf5, 0xc0, 0xc1, 0x69, 0x5d, 0x62, 0xd1, 0xda, 0xd1, 0xc6, 0x6d, 0xc4,
	0x69, 0xb5, 0x71, 0x36, 0xd1, 0x0a, 0xdb, 0x1b, 0xbc, 0x9b, 0xf6, 0x55, 0x78, 0x15, 0xde, 0x80,
	0xc7, 0x42, 0x76, 0x1c, 0xa7, 0xa9, 0xa0, 0x27, 0x7b, 0x7e, 0xf3, 0xed, 0xcc, 0x37, 0xb3, 0x0b,
	0x3d, 0x29, 0x53, 0xb2, 0xa0, 0x19, 0x23, 0x89, 0xc8, 0x55, 0x21, 0xd2, 0x94, 0x15, 0x84, 0xae,
	0x94, 0x20, 0x05, 0x9b, 0x3b, 0xcb, 0x42, 0x28, 0x61, 0x77, 0xff, 0xa5, 0x49, 0x44, 0x96, 0x89,
	0xbc, 0x56, 0xbc, 0x6e, 0x14, 0xec, 0x96, 0xe5, 0x8a, 0xb8, 0x1f, 0x3f, 0x9d, 0xae, 0x53, 0x3d,
	0x02, 0x2f, 0xbc, 0x95, 0x12, 0x98, 0xcd, 0x31, 0x5b, 0x70, 0xa9, 0x0a, 0xaa, 0xb8, 0xc8, 0xd1,
	0x5b, 0x00, 0x3e, 0x63, 0xb9, 0xe2, 0x73, 0xce, 0x0a, 0x4b, 0xeb, 0xea, 0x7d, 0x03, 0xdf, 0x23,
	0xa8, 0x0f, 0x86, 0xe4, 0x8b, 0x9c, 0xaa, 0x55, 0xc1, 0x2c, 0xbd, 0xab, 0xf5, 0x4d, 0x17, 0x9c,
	0xf1, 0x86, 0xe0, 0x6d, 0xb2, 0xf7, 0x47, 0x07, 0xbb, 0xee, 0x10, 0x8b, 0xb3, 0xc6, 0x20, 0x66,
	0x3f, 0x57, 0x4c, 0x2a, 0xf4, 0x01, 0x60, 0x6b, 0xcc, 0xd2, 0xea, 0x4a, 0x5f, 0x69, 0xc6, 0xfc,
	0x92, 0x60, 0x63, 0xb1, 0xf9, 0x45, 0xa7, 0xd0, 0xd9, 0x4c, 0x4e, 0x32, 0x26, 0x25, 0x5d, 0x6c,
	0x5a, 0x3f, 0x73, 0xea, 0x0e, 0x57, 0x6b, 0x8c, 0x8f, 0xe8, 0x4e, 0x8c, 0x3e, 0xc3, 0x9e, 0x54,
	0x54, 0x31, 0xab, 0xdd, 0xd5, 0xfa, 0x47, 0xee, 0x3b, 0xe7, 0xff, 0x8e, 0x9c, 0x71, 0x29, 0xc4,
	0x6b, 0xfd, 0xee, 0x9c, 0x4f, 0x1e, 0x9b, 0x93, 0xc2, 0x5e, 0x75, 0x12, 0x99, 0x70, 0x70, 0x1d,
	0x7e, 0x0b, 0xa3, 0x49, 0xd8, 0x69, 0xa1, 0x97, 0xf0, 0x1c, 0xfb, 0xde, 0xf9, 0x77, 0x12, 0x47,
	0xe4, 0x2c, 0x0a, 0xe3, 0x20, 0xbc, 0xf6, 0x3b, 0x1a, 0x7a, 0x05, 0x68, 0xe2, 0x05, 0x31, 0xb9,
	0x88, 0x30, 0x19, 0x5d, 0x7a, 0x67, 0xfe, 0x95, 0x1f, 0xc6, 0x1d, 0x1d, 0x1d, 0x83, 0xd5, 0xf0,
	0x1b, 0xef, 0x32, 0x38, 0x27, 0xa3, 0x68, 0x1c, 0xc4, 0x41, 0x14, 0x8e, 0x3b, 0xed, 0xde, 0xef,
	0x36, 0x1c, 0xed, 0x0e, 0x8a, 0x2c, 0xd8, 0x4f, 0x56, 0x52, 0x89, 0xac, 0x5a, 0x9d, 0x31, 0x6c,
	0xe1, 0x3a, 0x46, 0x5f, 0xc0, 0x9c, 0xd2, 0x34, 0x25, 0xcb, 0x94, 0x26, 0x6c, 0x56, 0x2f, 0xca,
	0x7e, 0xb0, 0x28, 0x67, 0x40, 0xd3, 0x74, 0x54, 0x29, 0x86, 0x2d, 0x0c, 0xd3, 0x26, 0x42, 0x1e,
	0x3c, 0xbd, 0xa3, 0x5c, 0x91, 0xb9, 0x28, 0xc8, 0x54, 0x28, 0x59, 0x6d, 0xce, 0x74, 0xdf, 0x3c,
	0x2c, 0x30, 0xa1, 0x5c, 0x5d, 0x88, 0x62, 0x20, 0x94, 0x1c, 0xb6, 0xb0, 0x79, 0xb7, 0x0d, 0x6d,
	0x06, 0xb0, 0x2d, 0x8f, 0x4e, 0x00, 0x14, 0xcf, 0x18, 0x51, 0xf4, 0x07, 0xcb, 0xab, 0x17, 0xa5,
	0x63, 0xa3, 0x24, 0x71, 0x09, 0xd0, 0x31, 0x18, 0xcb, 0x82, 0x25, 0x5c, 0x72, 0x91, 0x5b, 0xfa,
	0x3a, 0xdb, 0x00, 0x64, 0xc3, 0xe1, 0x8c, 0x4b, 0x45, 0xf3, 0xa4, 0xbc, 0xc2, 0x32, 0xd9, 0xc4,
	0xf6, 0x2f, 0x0d, 0xcc, 0x7b, 0x2e, 0xd0, 0x00, 0x8c, 0x5b, 0x2e, 0x52, 0xaa, 0x44, 0x21, 0x2d,
	0xad, 0xdb, 0xee, 0x9b, 0xee, 0xfb, 0x47, 0x5c, 0x3b, 0x37, 0xb5, 0x18, 0x6f, 0x8f, 0xd9, 0x3e,
	0x1c, 0x6e, 0x30, 0x3a, 0x81, 0xfd, 0xa9, 0x50, 0x84, 0xcf, 0x2a, 0xd3, 0xa6, 0xbb, 0xef, 0x0c,
	0x84, 0x0a, 0x66, 0x78, 0x6f, 0x5a, 0x7e, 0x76, 0xac, 0xe9, 0xbb, 0xd6, 0x06, 0x06, 0x1c, 0xd4,
	0x0f, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x94, 0xf3, 0x1d, 0xc9, 0x03, 0x00, 0x00,
}
