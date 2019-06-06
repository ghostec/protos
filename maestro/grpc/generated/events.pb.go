// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package eventforwarder

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

type RoomStatus_RoomStatusType int32

const (
	RoomStatus_ready       RoomStatus_RoomStatusType = 0
	RoomStatus_occupied    RoomStatus_RoomStatusType = 1
	RoomStatus_terminating RoomStatus_RoomStatusType = 2
	RoomStatus_terminated  RoomStatus_RoomStatusType = 3
)

var RoomStatus_RoomStatusType_name = map[int32]string{
	0: "ready",
	1: "occupied",
	2: "terminating",
	3: "terminated",
}
var RoomStatus_RoomStatusType_value = map[string]int32{
	"ready":       0,
	"occupied":    1,
	"terminating": 2,
	"terminated":  3,
}

func (x RoomStatus_RoomStatusType) String() string {
	return proto.EnumName(RoomStatus_RoomStatusType_name, int32(x))
}
func (RoomStatus_RoomStatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{1, 0}
}

type PlayerEvent_PlayerEventType int32

const (
	PlayerEvent_PLAYER_JOINED PlayerEvent_PlayerEventType = 0
	PlayerEvent_PLAYER_LEFT   PlayerEvent_PlayerEventType = 1
)

var PlayerEvent_PlayerEventType_name = map[int32]string{
	0: "PLAYER_JOINED",
	1: "PLAYER_LEFT",
}
var PlayerEvent_PlayerEventType_value = map[string]int32{
	"PLAYER_JOINED": 0,
	"PLAYER_LEFT":   1,
}

func (x PlayerEvent_PlayerEventType) String() string {
	return proto.EnumName(PlayerEvent_PlayerEventType_name, int32(x))
}
func (PlayerEvent_PlayerEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{3, 0}
}

type RoomEvent struct {
	Room                 *Room             `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	EventType            string            `protobuf:"bytes,2,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RoomEvent) Reset()         { *m = RoomEvent{} }
func (m *RoomEvent) String() string { return proto.CompactTextString(m) }
func (*RoomEvent) ProtoMessage()    {}
func (*RoomEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{0}
}
func (m *RoomEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomEvent.Unmarshal(m, b)
}
func (m *RoomEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomEvent.Marshal(b, m, deterministic)
}
func (dst *RoomEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomEvent.Merge(dst, src)
}
func (m *RoomEvent) XXX_Size() int {
	return xxx_messageInfo_RoomEvent.Size(m)
}
func (m *RoomEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RoomEvent proto.InternalMessageInfo

func (m *RoomEvent) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *RoomEvent) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *RoomEvent) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type RoomStatus struct {
	Room                 *Room                     `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	StatusType           RoomStatus_RoomStatusType `protobuf:"varint,2,opt,name=statusType,proto3,enum=eventforwarder.RoomStatus_RoomStatusType" json:"statusType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RoomStatus) Reset()         { *m = RoomStatus{} }
func (m *RoomStatus) String() string { return proto.CompactTextString(m) }
func (*RoomStatus) ProtoMessage()    {}
func (*RoomStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{1}
}
func (m *RoomStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomStatus.Unmarshal(m, b)
}
func (m *RoomStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomStatus.Marshal(b, m, deterministic)
}
func (dst *RoomStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomStatus.Merge(dst, src)
}
func (m *RoomStatus) XXX_Size() int {
	return xxx_messageInfo_RoomStatus.Size(m)
}
func (m *RoomStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RoomStatus proto.InternalMessageInfo

func (m *RoomStatus) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *RoomStatus) GetStatusType() RoomStatus_RoomStatusType {
	if m != nil {
		return m.StatusType
	}
	return RoomStatus_ready
}

type Room struct {
	Game                 string            `protobuf:"bytes,1,opt,name=game,proto3" json:"game,omitempty"`
	RoomId               string            `protobuf:"bytes,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Host                 string            `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32             `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RoomType             string            `protobuf:"bytes,7,opt,name=roomType,proto3" json:"roomType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{2}
}
func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (dst *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(dst, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetGame() string {
	if m != nil {
		return m.Game
	}
	return ""
}

func (m *Room) GetRoomId() string {
	if m != nil {
		return m.RoomId
	}
	return ""
}

func (m *Room) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Room) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Room) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Room) GetRoomType() string {
	if m != nil {
		return m.RoomType
	}
	return ""
}

type PlayerEvent struct {
	PlayerId             string                      `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Room                 *Room                       `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	EventType            PlayerEvent_PlayerEventType `protobuf:"varint,3,opt,name=eventType,proto3,enum=eventforwarder.PlayerEvent_PlayerEventType" json:"eventType,omitempty"`
	Metadata             map[string]string           `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PlayerEvent) Reset()         { *m = PlayerEvent{} }
func (m *PlayerEvent) String() string { return proto.CompactTextString(m) }
func (*PlayerEvent) ProtoMessage()    {}
func (*PlayerEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{3}
}
func (m *PlayerEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerEvent.Unmarshal(m, b)
}
func (m *PlayerEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerEvent.Marshal(b, m, deterministic)
}
func (dst *PlayerEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerEvent.Merge(dst, src)
}
func (m *PlayerEvent) XXX_Size() int {
	return xxx_messageInfo_PlayerEvent.Size(m)
}
func (m *PlayerEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerEvent proto.InternalMessageInfo

func (m *PlayerEvent) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *PlayerEvent) GetRoom() *Room {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *PlayerEvent) GetEventType() PlayerEvent_PlayerEventType {
	if m != nil {
		return m.EventType
	}
	return PlayerEvent_PLAYER_JOINED
}

func (m *PlayerEvent) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type RoomInfo struct {
	RoomType               string            `protobuf:"bytes,1,opt,name=roomType,proto3" json:"roomType,omitempty"`
	Game                   string            `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
	NumberOfTeams          int32             `protobuf:"varint,3,opt,name=numberOfTeams,proto3" json:"numberOfTeams,omitempty"`
	PlayersPerTeam         int32             `protobuf:"varint,4,opt,name=playersPerTeam,proto3" json:"playersPerTeam,omitempty"`
	MinimumNumberOfPlayers int32             `protobuf:"varint,5,opt,name=minimumNumberOfPlayers,proto3" json:"minimumNumberOfPlayers,omitempty"`
	MatchmakingScript      string            `protobuf:"bytes,6,opt,name=matchmakingScript,proto3" json:"matchmakingScript,omitempty"`
	WebhookUrl             string            `protobuf:"bytes,7,opt,name=webhookUrl,proto3" json:"webhookUrl,omitempty"`
	Metadata               map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tags                   map[string]string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{4}
}
func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (dst *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(dst, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetRoomType() string {
	if m != nil {
		return m.RoomType
	}
	return ""
}

func (m *RoomInfo) GetGame() string {
	if m != nil {
		return m.Game
	}
	return ""
}

func (m *RoomInfo) GetNumberOfTeams() int32 {
	if m != nil {
		return m.NumberOfTeams
	}
	return 0
}

func (m *RoomInfo) GetPlayersPerTeam() int32 {
	if m != nil {
		return m.PlayersPerTeam
	}
	return 0
}

func (m *RoomInfo) GetMinimumNumberOfPlayers() int32 {
	if m != nil {
		return m.MinimumNumberOfPlayers
	}
	return 0
}

func (m *RoomInfo) GetMatchmakingScript() string {
	if m != nil {
		return m.MatchmakingScript
	}
	return ""
}

func (m *RoomInfo) GetWebhookUrl() string {
	if m != nil {
		return m.WebhookUrl
	}
	return ""
}

func (m *RoomInfo) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *RoomInfo) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_events_3c71931a4e8022f3, []int{5}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RoomEvent)(nil), "eventforwarder.RoomEvent")
	proto.RegisterMapType((map[string]string)(nil), "eventforwarder.RoomEvent.MetadataEntry")
	proto.RegisterType((*RoomStatus)(nil), "eventforwarder.RoomStatus")
	proto.RegisterType((*Room)(nil), "eventforwarder.Room")
	proto.RegisterMapType((map[string]string)(nil), "eventforwarder.Room.MetadataEntry")
	proto.RegisterType((*PlayerEvent)(nil), "eventforwarder.PlayerEvent")
	proto.RegisterMapType((map[string]string)(nil), "eventforwarder.PlayerEvent.MetadataEntry")
	proto.RegisterType((*RoomInfo)(nil), "eventforwarder.RoomInfo")
	proto.RegisterMapType((map[string]string)(nil), "eventforwarder.RoomInfo.MetadataEntry")
	proto.RegisterMapType((map[string]string)(nil), "eventforwarder.RoomInfo.TagsEntry")
	proto.RegisterType((*Response)(nil), "eventforwarder.Response")
	proto.RegisterEnum("eventforwarder.RoomStatus_RoomStatusType", RoomStatus_RoomStatusType_name, RoomStatus_RoomStatusType_value)
	proto.RegisterEnum("eventforwarder.PlayerEvent_PlayerEventType", PlayerEvent_PlayerEventType_name, PlayerEvent_PlayerEventType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPCForwarderClient is the client API for GRPCForwarder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCForwarderClient interface {
	SendRoomInfo(ctx context.Context, in *RoomInfo, opts ...grpc.CallOption) (*Response, error)
	SendRoomStatus(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error)
	SendRoomPing(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error)
	SendRoomResync(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error)
	SendRoomEvent(ctx context.Context, in *RoomEvent, opts ...grpc.CallOption) (*Response, error)
	SendPlayerEvent(ctx context.Context, in *PlayerEvent, opts ...grpc.CallOption) (*Response, error)
}

type gRPCForwarderClient struct {
	cc *grpc.ClientConn
}

func NewGRPCForwarderClient(cc *grpc.ClientConn) GRPCForwarderClient {
	return &gRPCForwarderClient{cc}
}

func (c *gRPCForwarderClient) SendRoomInfo(ctx context.Context, in *RoomInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendRoomInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCForwarderClient) SendRoomStatus(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendRoomStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCForwarderClient) SendRoomPing(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendRoomPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCForwarderClient) SendRoomResync(ctx context.Context, in *RoomStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendRoomResync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCForwarderClient) SendRoomEvent(ctx context.Context, in *RoomEvent, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendRoomEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCForwarderClient) SendPlayerEvent(ctx context.Context, in *PlayerEvent, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/eventforwarder.GRPCForwarder/SendPlayerEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCForwarderServer is the server API for GRPCForwarder service.
type GRPCForwarderServer interface {
	SendRoomInfo(context.Context, *RoomInfo) (*Response, error)
	SendRoomStatus(context.Context, *RoomStatus) (*Response, error)
	SendRoomPing(context.Context, *RoomStatus) (*Response, error)
	SendRoomResync(context.Context, *RoomStatus) (*Response, error)
	SendRoomEvent(context.Context, *RoomEvent) (*Response, error)
	SendPlayerEvent(context.Context, *PlayerEvent) (*Response, error)
}

func RegisterGRPCForwarderServer(s *grpc.Server, srv GRPCForwarderServer) {
	s.RegisterService(&_GRPCForwarder_serviceDesc, srv)
}

func _GRPCForwarder_SendRoomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendRoomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendRoomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendRoomInfo(ctx, req.(*RoomInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCForwarder_SendRoomStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendRoomStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendRoomStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendRoomStatus(ctx, req.(*RoomStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCForwarder_SendRoomPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendRoomPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendRoomPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendRoomPing(ctx, req.(*RoomStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCForwarder_SendRoomResync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendRoomResync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendRoomResync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendRoomResync(ctx, req.(*RoomStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCForwarder_SendRoomEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendRoomEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendRoomEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendRoomEvent(ctx, req.(*RoomEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCForwarder_SendPlayerEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCForwarderServer).SendPlayerEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventforwarder.GRPCForwarder/SendPlayerEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCForwarderServer).SendPlayerEvent(ctx, req.(*PlayerEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPCForwarder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventforwarder.GRPCForwarder",
	HandlerType: (*GRPCForwarderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRoomInfo",
			Handler:    _GRPCForwarder_SendRoomInfo_Handler,
		},
		{
			MethodName: "SendRoomStatus",
			Handler:    _GRPCForwarder_SendRoomStatus_Handler,
		},
		{
			MethodName: "SendRoomPing",
			Handler:    _GRPCForwarder_SendRoomPing_Handler,
		},
		{
			MethodName: "SendRoomResync",
			Handler:    _GRPCForwarder_SendRoomResync_Handler,
		},
		{
			MethodName: "SendRoomEvent",
			Handler:    _GRPCForwarder_SendRoomEvent_Handler,
		},
		{
			MethodName: "SendPlayerEvent",
			Handler:    _GRPCForwarder_SendPlayerEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events.proto",
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_events_3c71931a4e8022f3) }

var fileDescriptor_events_3c71931a4e8022f3 = []byte{
	// 722 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x4f, 0xdb, 0x4c,
	0x10, 0xc6, 0xf9, 0x22, 0x99, 0x90, 0x60, 0x56, 0x08, 0xf9, 0xcd, 0xfb, 0xea, 0x55, 0x64, 0x55,
	0x34, 0xa8, 0x55, 0x0e, 0xa9, 0x4a, 0x51, 0x2b, 0x55, 0x6a, 0x21, 0xb4, 0x41, 0x14, 0xa2, 0x25,
	0x3d, 0xf4, 0x54, 0x2d, 0xf1, 0x62, 0x2c, 0xb0, 0xd7, 0xda, 0xdd, 0x80, 0x72, 0xeb, 0x9f, 0xeb,
	0xb1, 0xc7, 0x1e, 0xfa, 0x3b, 0xfa, 0x07, 0xaa, 0x5d, 0x7f, 0xe0, 0x40, 0x1c, 0x95, 0xa2, 0xde,
	0x66, 0x66, 0x67, 0x9e, 0x9d, 0x99, 0xe7, 0xf1, 0x26, 0xb0, 0x42, 0xaf, 0x68, 0x20, 0x45, 0x37,
	0xe4, 0x4c, 0x32, 0xd4, 0xd4, 0xde, 0x19, 0xe3, 0xd7, 0x84, 0x3b, 0x94, 0xdb, 0xdf, 0x0d, 0xa8,
	0x61, 0xc6, 0xfc, 0xbe, 0x0a, 0xa3, 0x0e, 0x94, 0x38, 0x63, 0xbe, 0x65, 0xb4, 0x8d, 0x4e, 0xbd,
	0xb7, 0xde, 0x9d, 0x4d, 0xee, 0xaa, 0x44, 0xac, 0x33, 0xd0, 0x7f, 0x50, 0xd3, 0x87, 0xa3, 0x69,
	0x48, 0xad, 0x42, 0xdb, 0xe8, 0xd4, 0xf0, 0x4d, 0x00, 0xed, 0x42, 0xd5, 0xa7, 0x92, 0x38, 0x44,
	0x12, 0xab, 0xd8, 0x2e, 0x76, 0xea, 0xbd, 0xc7, 0xf3, 0xb0, 0xf4, 0xa5, 0xdd, 0x0f, 0x71, 0x66,
	0x3f, 0x90, 0x7c, 0x8a, 0xd3, 0xc2, 0xd6, 0x2b, 0x68, 0xcc, 0x1c, 0x21, 0x13, 0x8a, 0x17, 0x74,
	0xaa, 0x9b, 0xab, 0x61, 0x65, 0xa2, 0x75, 0x28, 0x5f, 0x91, 0xcb, 0x49, 0xd2, 0x41, 0xe4, 0xbc,
	0x2c, 0xec, 0x18, 0xf6, 0x37, 0x03, 0x40, 0x5d, 0x71, 0x22, 0x89, 0x9c, 0x88, 0x7b, 0x0c, 0x36,
	0x00, 0x10, 0xba, 0x26, 0x9d, 0xac, 0xd9, 0xdb, 0x9a, 0x97, 0x1f, 0x21, 0x67, 0x4c, 0x55, 0x80,
	0x33, 0xc5, 0xf6, 0x01, 0x34, 0x67, 0x4f, 0x51, 0x0d, 0xca, 0x9c, 0x12, 0x67, 0x6a, 0x2e, 0xa1,
	0x15, 0xa8, 0xb2, 0xf1, 0x78, 0x12, 0x7a, 0xd4, 0x31, 0x0d, 0xb4, 0x0a, 0x75, 0x49, 0xb9, 0xef,
	0x05, 0x44, 0x7a, 0x81, 0x6b, 0x16, 0x50, 0x13, 0x20, 0x09, 0x50, 0xc7, 0x2c, 0xda, 0x3f, 0x0d,
	0x28, 0x29, 0x30, 0x84, 0xa0, 0xe4, 0x12, 0x9f, 0xc6, 0x5b, 0xd0, 0x36, 0xda, 0x80, 0x8a, 0xea,
	0x7d, 0xe0, 0xc4, 0x7b, 0x88, 0x3d, 0x95, 0x7b, 0xce, 0x84, 0xb4, 0x4a, 0x51, 0xae, 0xb2, 0x55,
	0x2c, 0x64, 0x5c, 0x5a, 0xe5, 0xb6, 0xd1, 0x29, 0x63, 0x6d, 0xa3, 0xd7, 0x19, 0xba, 0x2a, 0x9a,
	0x2e, 0x7b, 0xde, 0xc4, 0x79, 0x4c, 0xa1, 0x16, 0x54, 0xd5, 0x8d, 0x7a, 0x63, 0xcb, 0xfa, 0xae,
	0xd4, 0x7f, 0x18, 0x8b, 0x3f, 0x0a, 0x50, 0x1f, 0x5e, 0x92, 0x29, 0xe5, 0x91, 0x3e, 0x5b, 0x50,
	0x0d, 0xb5, 0x3b, 0x70, 0x62, 0x80, 0xd4, 0x4f, 0x29, 0x2e, 0xfc, 0x06, 0xc5, 0x19, 0xed, 0x16,
	0x35, 0xc3, 0x4f, 0x6e, 0xa7, 0x67, 0x6e, 0xcd, 0xda, 0x9a, 0xe3, 0x8c, 0xd0, 0xfb, 0x99, 0xcd,
	0x95, 0xf4, 0xe6, 0xb6, 0x16, 0x21, 0xfd, 0x15, 0xa9, 0x3f, 0x87, 0xd5, 0x5b, 0x1d, 0xa2, 0x35,
	0x68, 0x0c, 0x0f, 0xdf, 0x7c, 0xea, 0xe3, 0xcf, 0x07, 0xc7, 0x83, 0xa3, 0xfe, 0x9e, 0xb9, 0xa4,
	0x14, 0x16, 0x87, 0x0e, 0xfb, 0xfb, 0x23, 0xd3, 0xb0, 0xbf, 0x94, 0xa0, 0xaa, 0x96, 0x32, 0x08,
	0xce, 0xd8, 0x0c, 0x83, 0xc6, 0x2c, 0x83, 0xa9, 0xe2, 0x0a, 0x19, 0xc5, 0x3d, 0x82, 0x46, 0x30,
	0xf1, 0x4f, 0x29, 0x3f, 0x3e, 0x1b, 0x51, 0xe2, 0x0b, 0xbd, 0xc6, 0x32, 0x9e, 0x0d, 0xa2, 0x4d,
	0x68, 0x46, 0xf4, 0x88, 0x21, 0xe5, 0x2a, 0xa4, 0x95, 0x58, 0xc6, 0xb7, 0xa2, 0x68, 0x1b, 0x36,
	0x7c, 0x2f, 0xf0, 0xfc, 0x89, 0x7f, 0x14, 0xd7, 0x47, 0x03, 0x89, 0x58, 0xa5, 0x39, 0xa7, 0xe8,
	0x29, 0xac, 0xf9, 0x44, 0x8e, 0xcf, 0x7d, 0x72, 0xe1, 0x05, 0xee, 0xc9, 0x98, 0x7b, 0xa1, 0xb4,
	0x2a, 0xba, 0xcd, 0xbb, 0x07, 0xe8, 0x7f, 0x80, 0x6b, 0x7a, 0x7a, 0xce, 0xd8, 0xc5, 0x47, 0x7e,
	0x19, 0xeb, 0x34, 0x13, 0x41, 0x6f, 0x33, 0x5c, 0x56, 0x35, 0x97, 0x9b, 0xf3, 0x44, 0xa4, 0xf6,
	0x95, 0xfb, 0x25, 0x6c, 0x43, 0x49, 0x12, 0x57, 0x58, 0xb5, 0xfc, 0xaf, 0x48, 0xd7, 0x8f, 0x88,
	0x2b, 0xa2, 0x5a, 0x9d, 0xff, 0x20, 0x01, 0xb4, 0x5e, 0x40, 0x2d, 0xc5, 0xbb, 0x97, 0x72, 0x76,
	0xa0, 0x8a, 0xa9, 0x08, 0x59, 0x20, 0x34, 0xcb, 0x63, 0xe6, 0x44, 0xec, 0x97, 0xb1, 0xb6, 0x91,
	0x05, 0xcb, 0x3e, 0x15, 0x82, 0xb8, 0x49, 0x6d, 0xe2, 0xf6, 0xbe, 0x16, 0xa1, 0xf1, 0x0e, 0x0f,
	0x77, 0xf7, 0x93, 0xd1, 0xd0, 0x1e, 0xac, 0x9c, 0xd0, 0xc0, 0x49, 0x15, 0x65, 0xe5, 0xcd, 0xde,
	0xba, 0x7b, 0x12, 0xf7, 0x60, 0x2f, 0xa1, 0xf7, 0xd0, 0x4c, 0x50, 0xe2, 0x97, 0xbb, 0x95, 0xff,
	0xf6, 0x2e, 0x44, 0xda, 0xbf, 0xe9, 0x67, 0xe8, 0x05, 0xee, 0x1f, 0xe3, 0x64, 0x3a, 0xc2, 0x54,
	0x4c, 0x83, 0xf1, 0x03, 0x3a, 0x6a, 0x24, 0x48, 0xd1, 0x6b, 0xf6, 0x4f, 0xee, 0x6f, 0xe2, 0x42,
	0x9c, 0x03, 0x58, 0x55, 0x38, 0xd9, 0x77, 0xf1, 0xdf, 0x05, 0x8f, 0xce, 0x22, 0xac, 0xd3, 0x8a,
	0xfe, 0x57, 0xf0, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0xd6, 0xff, 0x11, 0x25, 0x08,
	0x00, 0x00,
}
