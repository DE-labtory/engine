// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Envelope
	Empty
	Message
	Block
	Transaction
	PeerTable
	Peer
	ConsensusMessage
	View
	ElectionMessage
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Envelope struct {
	// marshalled Message
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// signed Message
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// sender's public key
	Pubkey []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Envelope) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Envelope) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Envelope) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Message struct {
	Channel []byte `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	// Types that are valid to be assigned to Content:
	//	*Message_Block
	//	*Message_Transaction
	//	*Message_PeerTable
	//	*Message_ConsensusMessage
	//	*Message_ElectionMessage
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Block struct {
	Block *Block `protobuf:"bytes,2,opt,name=block,oneof"`
}
type Message_Transaction struct {
	Transaction *Transaction `protobuf:"bytes,3,opt,name=transaction,oneof"`
}
type Message_PeerTable struct {
	PeerTable *PeerTable `protobuf:"bytes,4,opt,name=peerTable,oneof"`
}
type Message_ConsensusMessage struct {
	ConsensusMessage *ConsensusMessage `protobuf:"bytes,5,opt,name=consensusMessage,oneof"`
}
type Message_ElectionMessage struct {
	ElectionMessage *ElectionMessage `protobuf:"bytes,6,opt,name=electionMessage,oneof"`
}

func (*Message_Block) isMessage_Content()            {}
func (*Message_Transaction) isMessage_Content()      {}
func (*Message_PeerTable) isMessage_Content()        {}
func (*Message_ConsensusMessage) isMessage_Content() {}
func (*Message_ElectionMessage) isMessage_Content()  {}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Message) GetChannel() []byte {
	if m != nil {
		return m.Channel
	}
	return nil
}

func (m *Message) GetBlock() *Block {
	if x, ok := m.GetContent().(*Message_Block); ok {
		return x.Block
	}
	return nil
}

func (m *Message) GetTransaction() *Transaction {
	if x, ok := m.GetContent().(*Message_Transaction); ok {
		return x.Transaction
	}
	return nil
}

func (m *Message) GetPeerTable() *PeerTable {
	if x, ok := m.GetContent().(*Message_PeerTable); ok {
		return x.PeerTable
	}
	return nil
}

func (m *Message) GetConsensusMessage() *ConsensusMessage {
	if x, ok := m.GetContent().(*Message_ConsensusMessage); ok {
		return x.ConsensusMessage
	}
	return nil
}

func (m *Message) GetElectionMessage() *ElectionMessage {
	if x, ok := m.GetContent().(*Message_ElectionMessage); ok {
		return x.ElectionMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_Block)(nil),
		(*Message_Transaction)(nil),
		(*Message_PeerTable)(nil),
		(*Message_ConsensusMessage)(nil),
		(*Message_ElectionMessage)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Block:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Block); err != nil {
			return err
		}
	case *Message_Transaction:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transaction); err != nil {
			return err
		}
	case *Message_PeerTable:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PeerTable); err != nil {
			return err
		}
	case *Message_ConsensusMessage:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ConsensusMessage); err != nil {
			return err
		}
	case *Message_ElectionMessage:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ElectionMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Content has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 2: // content.block
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Block)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Block{msg}
		return true, err
	case 3: // content.transaction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transaction)
		err := b.DecodeMessage(msg)
		m.Content = &Message_Transaction{msg}
		return true, err
	case 4: // content.peerTable
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PeerTable)
		err := b.DecodeMessage(msg)
		m.Content = &Message_PeerTable{msg}
		return true, err
	case 5: // content.consensusMessage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ConsensusMessage)
		err := b.DecodeMessage(msg)
		m.Content = &Message_ConsensusMessage{msg}
		return true, err
	case 6: // content.electionMessage
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ElectionMessage)
		err := b.DecodeMessage(msg)
		m.Content = &Message_ElectionMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// content
	switch x := m.Content.(type) {
	case *Message_Block:
		s := proto.Size(x.Block)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Transaction:
		s := proto.Size(x.Transaction)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PeerTable:
		s := proto.Size(x.PeerTable)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ConsensusMessage:
		s := proto.Size(x.ConsensusMessage)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_ElectionMessage:
		s := proto.Size(x.ElectionMessage)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Block struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Transaction struct {
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type PeerTable struct {
	MyID    string           `protobuf:"bytes,1,opt,name=MyID" json:"MyID,omitempty"`
	PeerMap map[string]*Peer `protobuf:"bytes,2,rep,name=PeerMap" json:"PeerMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PeerTable) Reset()                    { *m = PeerTable{} }
func (m *PeerTable) String() string            { return proto.CompactTextString(m) }
func (*PeerTable) ProtoMessage()               {}
func (*PeerTable) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PeerTable) GetMyID() string {
	if m != nil {
		return m.MyID
	}
	return ""
}

func (m *PeerTable) GetPeerMap() map[string]*Peer {
	if m != nil {
		return m.PeerMap
	}
	return nil
}

type Peer struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ipAddress" json:"ipAddress,omitempty"`
	Port      string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
	PeerID    string `protobuf:"bytes,3,opt,name=peerID" json:"peerID,omitempty"`
	HeartBeat int32  `protobuf:"varint,4,opt,name=heartBeat" json:"heartBeat,omitempty"`
	PubKey    []byte `protobuf:"bytes,5,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Peer) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *Peer) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Peer) GetPeerID() string {
	if m != nil {
		return m.PeerID
	}
	return ""
}

func (m *Peer) GetHeartBeat() int32 {
	if m != nil {
		return m.HeartBeat
	}
	return 0
}

func (m *Peer) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type ConsensusMessage struct {
	ConsensusID string `protobuf:"bytes,1,opt,name=ConsensusID" json:"ConsensusID,omitempty"`
	View        *View  `protobuf:"bytes,2,opt,name=View" json:"View,omitempty"`
	SequenceID  int64  `protobuf:"varint,3,opt,name=SequenceID" json:"SequenceID,omitempty"`
	Block       *Block `protobuf:"bytes,4,opt,name=Block" json:"Block,omitempty"`
	SenderID    string `protobuf:"bytes,5,opt,name=SenderID" json:"SenderID,omitempty"`
	MsgType     int32  `protobuf:"varint,6,opt,name=MsgType" json:"MsgType,omitempty"`
}

func (m *ConsensusMessage) Reset()                    { *m = ConsensusMessage{} }
func (m *ConsensusMessage) String() string            { return proto.CompactTextString(m) }
func (*ConsensusMessage) ProtoMessage()               {}
func (*ConsensusMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ConsensusMessage) GetConsensusID() string {
	if m != nil {
		return m.ConsensusID
	}
	return ""
}

func (m *ConsensusMessage) GetView() *View {
	if m != nil {
		return m.View
	}
	return nil
}

func (m *ConsensusMessage) GetSequenceID() int64 {
	if m != nil {
		return m.SequenceID
	}
	return 0
}

func (m *ConsensusMessage) GetBlock() *Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ConsensusMessage) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *ConsensusMessage) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

type View struct {
	ViewID   string   `protobuf:"bytes,1,opt,name=ViewID" json:"ViewID,omitempty"`
	LeaderID string   `protobuf:"bytes,2,opt,name=LeaderID" json:"LeaderID,omitempty"`
	PeerID   []string `protobuf:"bytes,3,rep,name=PeerID" json:"PeerID,omitempty"`
}

func (m *View) Reset()                    { *m = View{} }
func (m *View) String() string            { return proto.CompactTextString(m) }
func (*View) ProtoMessage()               {}
func (*View) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *View) GetViewID() string {
	if m != nil {
		return m.ViewID
	}
	return ""
}

func (m *View) GetLeaderID() string {
	if m != nil {
		return m.LeaderID
	}
	return ""
}

func (m *View) GetPeerID() []string {
	if m != nil {
		return m.PeerID
	}
	return nil
}

type ElectionMessage struct {
	LastBlockHash string   `protobuf:"bytes,1,opt,name=LastBlockHash" json:"LastBlockHash,omitempty"`
	SenderID      string   `protobuf:"bytes,2,opt,name=SenderID" json:"SenderID,omitempty"`
	MsgType       int32    `protobuf:"varint,3,opt,name=MsgType" json:"MsgType,omitempty"`
	Term          int64    `protobuf:"varint,4,opt,name=Term" json:"Term,omitempty"`
	PeerIDs       []string `protobuf:"bytes,5,rep,name=PeerIDs" json:"PeerIDs,omitempty"`
}

func (m *ElectionMessage) Reset()                    { *m = ElectionMessage{} }
func (m *ElectionMessage) String() string            { return proto.CompactTextString(m) }
func (*ElectionMessage) ProtoMessage()               {}
func (*ElectionMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ElectionMessage) GetLastBlockHash() string {
	if m != nil {
		return m.LastBlockHash
	}
	return ""
}

func (m *ElectionMessage) GetSenderID() string {
	if m != nil {
		return m.SenderID
	}
	return ""
}

func (m *ElectionMessage) GetMsgType() int32 {
	if m != nil {
		return m.MsgType
	}
	return 0
}

func (m *ElectionMessage) GetTerm() int64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *ElectionMessage) GetPeerIDs() []string {
	if m != nil {
		return m.PeerIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "message.Envelope")
	proto.RegisterType((*Empty)(nil), "message.Empty")
	proto.RegisterType((*Message)(nil), "message.Message")
	proto.RegisterType((*Block)(nil), "message.Block")
	proto.RegisterType((*Transaction)(nil), "message.Transaction")
	proto.RegisterType((*PeerTable)(nil), "message.PeerTable")
	proto.RegisterType((*Peer)(nil), "message.Peer")
	proto.RegisterType((*ConsensusMessage)(nil), "message.ConsensusMessage")
	proto.RegisterType((*View)(nil), "message.View")
	proto.RegisterType((*ElectionMessage)(nil), "message.ElectionMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageService service

type MessageServiceClient interface {
	Stream(ctx context.Context, opts ...grpc.CallOption) (MessageService_StreamClient, error)
	// Ping is used to probe a remote peer's aliveness
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Stream(ctx context.Context, opts ...grpc.CallOption) (MessageService_StreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MessageService_serviceDesc.Streams[0], c.cc, "/message.MessageService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceStreamClient{stream}
	return x, nil
}

type MessageService_StreamClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type messageServiceStreamClient struct {
	grpc.ClientStream
}

func (x *messageServiceStreamClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceStreamClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/message.MessageService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageService service

type MessageServiceServer interface {
	Stream(MessageService_StreamServer) error
	// Ping is used to probe a remote peer's aliveness
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Stream(&messageServiceStreamServer{stream})
}

type MessageService_StreamServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type messageServiceStreamServer struct {
	grpc.ServerStream
}

func (x *messageServiceStreamServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceStreamServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _MessageService_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _MessageService_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}

// Client API for PeerService service

type PeerServiceClient interface {
	GetPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Peer, error)
}

type peerServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerServiceClient(cc *grpc.ClientConn) PeerServiceClient {
	return &peerServiceClient{cc}
}

func (c *peerServiceClient) GetPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Peer, error) {
	out := new(Peer)
	err := grpc.Invoke(ctx, "/message.PeerService/GetPeer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeerService service

type PeerServiceServer interface {
	GetPeer(context.Context, *Empty) (*Peer, error)
}

func RegisterPeerServiceServer(s *grpc.Server, srv PeerServiceServer) {
	s.RegisterService(&_PeerService_serviceDesc, srv)
}

func _PeerService_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServiceServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.PeerService/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServiceServer).GetPeer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.PeerService",
	HandlerType: (*PeerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPeer",
			Handler:    _PeerService_GetPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// Client API for TestConsensusService service

type TestConsensusServiceClient interface {
	StartConsensus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type testConsensusServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestConsensusServiceClient(cc *grpc.ClientConn) TestConsensusServiceClient {
	return &testConsensusServiceClient{cc}
}

func (c *testConsensusServiceClient) StartConsensus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/message.TestConsensusService/StartConsensus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestConsensusService service

type TestConsensusServiceServer interface {
	StartConsensus(context.Context, *Empty) (*Empty, error)
}

func RegisterTestConsensusServiceServer(s *grpc.Server, srv TestConsensusServiceServer) {
	s.RegisterService(&_TestConsensusService_serviceDesc, srv)
}

func _TestConsensusService_StartConsensus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestConsensusServiceServer).StartConsensus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.TestConsensusService/StartConsensus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestConsensusServiceServer).StartConsensus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestConsensusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.TestConsensusService",
	HandlerType: (*TestConsensusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartConsensus",
			Handler:    _TestConsensusService_StartConsensus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xb6, 0x93, 0x38, 0x69, 0x5e, 0x9a, 0xb6, 0x8c, 0x2a, 0x64, 0x02, 0x82, 0x62, 0x2a, 0x14,
	0xb1, 0xa8, 0x90, 0x61, 0xd1, 0xb2, 0xa3, 0x34, 0x6a, 0x4a, 0x1b, 0xa9, 0x9a, 0x44, 0x2c, 0xd8,
	0x4d, 0x9c, 0xa7, 0x34, 0xaa, 0x33, 0x36, 0x9e, 0x49, 0x91, 0x2f, 0xc0, 0x25, 0x58, 0x73, 0xa2,
	0x5e, 0x08, 0xcd, 0x78, 0x6c, 0x27, 0x8e, 0x10, 0x2b, 0xcf, 0xf7, 0xfe, 0x7f, 0xbe, 0x67, 0xe8,
	0x2e, 0x51, 0x08, 0x36, 0xc7, 0x93, 0x38, 0x89, 0x64, 0x44, 0x5a, 0x06, 0x7a, 0xdf, 0x61, 0x67,
	0xc0, 0x1f, 0x30, 0x8c, 0x62, 0x24, 0x2e, 0xb4, 0x62, 0x96, 0x86, 0x11, 0x9b, 0xb9, 0xf6, 0x91,
	0xdd, 0xdf, 0xa5, 0x39, 0x24, 0x2f, 0xa0, 0x2d, 0x16, 0x73, 0xce, 0xe4, 0x2a, 0x41, 0xb7, 0xa6,
	0x75, 0xa5, 0x80, 0x3c, 0x85, 0x66, 0xbc, 0x9a, 0xde, 0x63, 0xea, 0xd6, 0xb5, 0xca, 0x20, 0xaf,
	0x05, 0xce, 0x60, 0x19, 0xcb, 0xd4, 0x7b, 0xac, 0x41, 0x6b, 0x94, 0x25, 0x54, 0x49, 0x82, 0x3b,
	0xc6, 0x39, 0x86, 0x79, 0x12, 0x03, 0xc9, 0x5b, 0x70, 0xa6, 0x61, 0x14, 0xdc, 0xeb, 0x04, 0x1d,
	0x7f, 0xef, 0x24, 0x2f, 0xf9, 0x5c, 0x49, 0x87, 0x16, 0xcd, 0xd4, 0xe4, 0x14, 0x3a, 0x32, 0x61,
	0x5c, 0xb0, 0x40, 0x2e, 0x22, 0xae, 0x73, 0x76, 0xfc, 0xc3, 0xc2, 0x7a, 0x52, 0xea, 0x86, 0x16,
	0x5d, 0x37, 0x25, 0x3e, 0xb4, 0x63, 0xc4, 0x64, 0xc2, 0xa6, 0x21, 0xba, 0x0d, 0xed, 0x47, 0x0a,
	0xbf, 0xdb, 0x5c, 0x33, 0xb4, 0x68, 0x69, 0x46, 0x2e, 0xe1, 0x20, 0x88, 0xb8, 0x40, 0x2e, 0x56,
	0xc2, 0xf4, 0xe0, 0x3a, 0xda, 0xf5, 0x59, 0xe1, 0xfa, 0xa5, 0x62, 0x30, 0xb4, 0xe8, 0x96, 0x13,
	0xb9, 0x80, 0x7d, 0x0c, 0x51, 0x17, 0x92, 0xc7, 0x69, 0xea, 0x38, 0x6e, 0x11, 0x67, 0xb0, 0xa9,
	0x1f, 0x5a, 0xb4, 0xea, 0x72, 0xde, 0x86, 0x56, 0x10, 0x71, 0x89, 0x5c, 0x7a, 0xcf, 0xc1, 0xd1,
	0x93, 0x21, 0x04, 0x1a, 0x33, 0x26, 0x99, 0x99, 0xa7, 0x7e, 0x7b, 0x5d, 0xe8, 0xac, 0x0d, 0xc2,
	0xfb, 0x63, 0x43, 0xbb, 0x68, 0x50, 0x39, 0x8c, 0xd2, 0xab, 0x0b, 0xed, 0xd0, 0xa6, 0xfa, 0x4d,
	0xce, 0xa0, 0xa5, 0x0c, 0x46, 0x2c, 0x76, 0x6b, 0x47, 0xf5, 0x7e, 0xc7, 0x7f, 0xb5, 0x3d, 0x99,
	0x13, 0x63, 0x31, 0xe0, 0x32, 0x49, 0x69, 0x6e, 0xdf, 0xbb, 0x82, 0xdd, 0x75, 0x05, 0x39, 0x80,
	0xba, 0x22, 0x43, 0x16, 0x5d, 0x3d, 0xc9, 0x1b, 0x70, 0x1e, 0x58, 0xb8, 0x42, 0xb3, 0xda, 0xee,
	0x46, 0x68, 0x9a, 0xe9, 0x3e, 0xd5, 0x4e, 0x6d, 0xef, 0x97, 0x0d, 0x0d, 0x25, 0x53, 0x8c, 0x5b,
	0xc4, 0x9f, 0x67, 0xb3, 0x04, 0x85, 0x30, 0x91, 0x4a, 0x81, 0x6a, 0x20, 0x8e, 0x12, 0xa9, 0xc3,
	0xb5, 0xa9, 0x7e, 0x6b, 0x16, 0x22, 0x26, 0x57, 0x17, 0x9a, 0x11, 0x6d, 0x6a, 0x90, 0x8a, 0x74,
	0x87, 0x2c, 0x91, 0xe7, 0xc8, 0xa4, 0x5e, 0xba, 0x43, 0x4b, 0x81, 0xe1, 0xee, 0x35, 0xa6, 0x7a,
	0xa9, 0x19, 0x77, 0xaf, 0x31, 0xf5, 0x1e, 0x6d, 0x38, 0xa8, 0xae, 0x95, 0x1c, 0x41, 0xa7, 0x90,
	0x15, 0xe3, 0x5b, 0x17, 0x91, 0xd7, 0xd0, 0xf8, 0xb6, 0xc0, 0x9f, 0x5b, 0x7d, 0x2a, 0x21, 0xd5,
	0x2a, 0xf2, 0x12, 0x60, 0x8c, 0x3f, 0x56, 0xc8, 0x03, 0x34, 0xb5, 0xd6, 0xe9, 0x9a, 0x84, 0x1c,
	0x9b, 0xb5, 0x1a, 0x82, 0x56, 0xce, 0x80, 0x9a, 0x9d, 0xf7, 0x60, 0x67, 0x8c, 0x7c, 0xa6, 0xfb,
	0x75, 0x74, 0x1d, 0x05, 0x56, 0x27, 0x36, 0x12, 0xf3, 0x49, 0x1a, 0x67, 0x0c, 0x73, 0x68, 0x0e,
	0x3d, 0x9a, 0x95, 0xa7, 0xba, 0x56, 0xdf, 0xa2, 0x07, 0x83, 0x54, 0xd4, 0x1b, 0x64, 0x59, 0xd4,
	0x6c, 0xb6, 0x05, 0x56, 0x3e, 0xb7, 0xf9, 0x7c, 0xeb, 0xca, 0x27, 0x43, 0xde, 0x6f, 0x1b, 0xf6,
	0x2b, 0xc4, 0x25, 0xc7, 0xd0, 0xbd, 0x61, 0x42, 0x66, 0x87, 0xcb, 0xc4, 0x9d, 0x49, 0xb3, 0x29,
	0xdc, 0xe8, 0xa1, 0xf6, 0xef, 0x1e, 0xea, 0x1b, 0x3d, 0xa8, 0xdd, 0x4f, 0x30, 0x59, 0xea, 0xf1,
	0xd4, 0xa9, 0x7e, 0x2b, 0xeb, 0xac, 0x1a, 0xe1, 0x3a, 0xba, 0xb8, 0x1c, 0xfa, 0x31, 0xec, 0x99,
	0xa2, 0xc6, 0x98, 0x3c, 0x2c, 0x02, 0x24, 0x1f, 0xa1, 0x39, 0x96, 0x09, 0xb2, 0x25, 0x79, 0x52,
	0x1e, 0x9e, 0xf9, 0x05, 0xf6, 0xb6, 0x45, 0x9e, 0xd5, 0xb7, 0xdf, 0xdb, 0xa4, 0x0f, 0x8d, 0xdb,
	0x05, 0x9f, 0x93, 0x72, 0x1d, 0xfa, 0xd7, 0xd6, 0xab, 0x60, 0xcf, 0xf2, 0xcf, 0xa0, 0xa3, 0x92,
	0xe7, 0xe9, 0xde, 0x41, 0xeb, 0x12, 0xa5, 0xe6, 0x74, 0xd5, 0x77, 0xf3, 0x0c, 0x3c, 0xcb, 0xff,
	0x0a, 0x87, 0x13, 0x14, 0xb2, 0x20, 0x54, 0x1e, 0xc3, 0x87, 0xbd, 0xb1, 0x64, 0x49, 0xa9, 0xf8,
	0x7f, 0x19, 0xd3, 0xa6, 0xfe, 0xd1, 0x7f, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xe5, 0x2d,
	0x86, 0xf9, 0x05, 0x00, 0x00,
}
