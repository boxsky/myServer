// Code generated by protoc-gen-go.
// source: login_pay.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// 输入邀请码
type CInviteCode struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	Uid              *uint32 `protobuf:"varint,2,opt" json:"Uid,omitempty"`
	InviteCode       *uint32 `protobuf:"varint,3,opt" json:"InviteCode,omitempty"`
	Sign             *string `protobuf:"bytes,4,opt" json:"Sign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CInviteCode) Reset()         { *m = CInviteCode{} }
func (m *CInviteCode) String() string { return proto.CompactTextString(m) }
func (*CInviteCode) ProtoMessage()    {}

func (m *CInviteCode) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CInviteCode) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *CInviteCode) GetInviteCode() uint32 {
	if m != nil && m.InviteCode != nil {
		return *m.InviteCode
	}
	return 0
}

func (m *CInviteCode) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

// 输入邀请码响应
type SInviteCode struct {
	State            *uint32 `protobuf:"varint,1,opt" json:"State,omitempty"`
	Uid              *uint32 `protobuf:"varint,2,opt" json:"Uid,omitempty"`
	Player           *uint32 `protobuf:"varint,3,opt" json:"Player,omitempty"`
	Inviter          *uint32 `protobuf:"varint,4,opt" json:"Inviter,omitempty"`
	Sign             *string `protobuf:"bytes,5,opt" json:"Sign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SInviteCode) Reset()         { *m = SInviteCode{} }
func (m *SInviteCode) String() string { return proto.CompactTextString(m) }
func (*SInviteCode) ProtoMessage()    {}

func (m *SInviteCode) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SInviteCode) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *SInviteCode) GetPlayer() uint32 {
	if m != nil && m.Player != nil {
		return *m.Player
	}
	return 0
}

func (m *SInviteCode) GetInviter() uint32 {
	if m != nil && m.Inviter != nil {
		return *m.Inviter
	}
	return 0
}

func (m *SInviteCode) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

// 查询邀请码
type CInviteCodeQuery struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	Uid              *uint32 `protobuf:"varint,2,opt" json:"Uid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CInviteCodeQuery) Reset()         { *m = CInviteCodeQuery{} }
func (m *CInviteCodeQuery) String() string { return proto.CompactTextString(m) }
func (*CInviteCodeQuery) ProtoMessage()    {}

func (m *CInviteCodeQuery) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CInviteCodeQuery) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

// 查询邀请码响应
type SInviteCodeQuery struct {
	State            *uint32 `protobuf:"varint,1,opt" json:"State,omitempty"`
	IsInputted       *bool   `protobuf:"varint,2,opt" json:"IsInputted,omitempty"`
	Role             *uint32 `protobuf:"varint,3,opt" json:"Role,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SInviteCodeQuery) Reset()         { *m = SInviteCodeQuery{} }
func (m *SInviteCodeQuery) String() string { return proto.CompactTextString(m) }
func (*SInviteCodeQuery) ProtoMessage()    {}

func (m *SInviteCodeQuery) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SInviteCodeQuery) GetIsInputted() bool {
	if m != nil && m.IsInputted != nil {
		return *m.IsInputted
	}
	return false
}

func (m *SInviteCodeQuery) GetRole() uint32 {
	if m != nil && m.Role != nil {
		return *m.Role
	}
	return 0
}

// 商品列表
type CProducts struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CProducts) Reset()         { *m = CProducts{} }
func (m *CProducts) String() string { return proto.CompactTextString(m) }
func (*CProducts) ProtoMessage()    {}

func (m *CProducts) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

// 商品列表响应
type SProducts struct {
	State            *uint32 `protobuf:"varint,1,opt" json:"State,omitempty"`
	List             *string `protobuf:"bytes,2,opt" json:"List,omitempty"`
	Info             *string `protobuf:"bytes,3,opt" json:"Info,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SProducts) Reset()         { *m = SProducts{} }
func (m *SProducts) String() string { return proto.CompactTextString(m) }
func (*SProducts) ProtoMessage()    {}

func (m *SProducts) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SProducts) GetList() string {
	if m != nil && m.List != nil {
		return *m.List
	}
	return ""
}

func (m *SProducts) GetInfo() string {
	if m != nil && m.Info != nil {
		return *m.Info
	}
	return ""
}

// 下单
type CPay struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	Uid              *uint32 `protobuf:"varint,2,opt" json:"Uid,omitempty"`
	Productid        *uint32 `protobuf:"varint,3,opt" json:"Productid,omitempty"`
	PF               *uint32 `protobuf:"varint,4,opt" json:"PF,omitempty"`
	Sign             *string `protobuf:"bytes,5,opt" json:"Sign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CPay) Reset()         { *m = CPay{} }
func (m *CPay) String() string { return proto.CompactTextString(m) }
func (*CPay) ProtoMessage()    {}

func (m *CPay) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CPay) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *CPay) GetProductid() uint32 {
	if m != nil && m.Productid != nil {
		return *m.Productid
	}
	return 0
}

func (m *CPay) GetPF() uint32 {
	if m != nil && m.PF != nil {
		return *m.PF
	}
	return 0
}

func (m *CPay) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

// 下单
type SPay struct {
	State            *uint32 `protobuf:"varint,1,opt" json:"State,omitempty"`
	Tradeno          *string `protobuf:"bytes,2,opt" json:"Tradeno,omitempty"`
	Partnerid        *string `protobuf:"bytes,3,opt" json:"Partnerid,omitempty"`
	Prepayid         *string `protobuf:"bytes,4,opt" json:"Prepayid,omitempty"`
	Noncestr         *string `protobuf:"bytes,5,opt" json:"Noncestr,omitempty"`
	Timestamp        *string `protobuf:"bytes,6,opt" json:"Timestamp,omitempty"`
	Sign             *string `protobuf:"bytes,7,opt" json:"Sign,omitempty"`
	Sign2            *string `protobuf:"bytes,8,opt" json:"Sign2,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SPay) Reset()         { *m = SPay{} }
func (m *SPay) String() string { return proto.CompactTextString(m) }
func (*SPay) ProtoMessage()    {}

func (m *SPay) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SPay) GetTradeno() string {
	if m != nil && m.Tradeno != nil {
		return *m.Tradeno
	}
	return ""
}

func (m *SPay) GetPartnerid() string {
	if m != nil && m.Partnerid != nil {
		return *m.Partnerid
	}
	return ""
}

func (m *SPay) GetPrepayid() string {
	if m != nil && m.Prepayid != nil {
		return *m.Prepayid
	}
	return ""
}

func (m *SPay) GetNoncestr() string {
	if m != nil && m.Noncestr != nil {
		return *m.Noncestr
	}
	return ""
}

func (m *SPay) GetTimestamp() string {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return ""
}

func (m *SPay) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

func (m *SPay) GetSign2() string {
	if m != nil && m.Sign2 != nil {
		return *m.Sign2
	}
	return ""
}

// 支付查询
type COrderQuery struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	Uid              *uint32 `protobuf:"varint,2,opt" json:"Uid,omitempty"`
	Tradeno          *string `protobuf:"bytes,3,opt" json:"Tradeno,omitempty"`
	PF               *uint32 `protobuf:"varint,4,opt" json:"PF,omitempty"`
	Sign             *string `protobuf:"bytes,5,opt" json:"Sign,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *COrderQuery) Reset()         { *m = COrderQuery{} }
func (m *COrderQuery) String() string { return proto.CompactTextString(m) }
func (*COrderQuery) ProtoMessage()    {}

func (m *COrderQuery) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *COrderQuery) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *COrderQuery) GetTradeno() string {
	if m != nil && m.Tradeno != nil {
		return *m.Tradeno
	}
	return ""
}

func (m *COrderQuery) GetPF() uint32 {
	if m != nil && m.PF != nil {
		return *m.PF
	}
	return 0
}

func (m *COrderQuery) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

// 支付查询响应
type SOrderQuery struct {
	State            *uint32  `protobuf:"varint,1,opt" json:"State,omitempty"`
	Appid            *uint32  `protobuf:"varint,2,opt" json:"Appid,omitempty"`
	Uid              *uint32  `protobuf:"varint,3,opt" json:"Uid,omitempty"`
	Paymoney         *float64 `protobuf:"fixed64,4,opt" json:"Paymoney,omitempty"`
	Tradeno          *string  `protobuf:"bytes,5,opt" json:"Tradeno,omitempty"`
	Transactionid    *string  `protobuf:"bytes,6,opt" json:"Transactionid,omitempty"`
	Paytime          *string  `protobuf:"bytes,7,opt" json:"Paytime,omitempty"`
	Num              *uint32  `protobuf:"varint,8,opt" json:"Num,omitempty"`
	Sign             *string  `protobuf:"bytes,9,opt" json:"Sign,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SOrderQuery) Reset()         { *m = SOrderQuery{} }
func (m *SOrderQuery) String() string { return proto.CompactTextString(m) }
func (*SOrderQuery) ProtoMessage()    {}

func (m *SOrderQuery) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SOrderQuery) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *SOrderQuery) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *SOrderQuery) GetPaymoney() float64 {
	if m != nil && m.Paymoney != nil {
		return *m.Paymoney
	}
	return 0
}

func (m *SOrderQuery) GetTradeno() string {
	if m != nil && m.Tradeno != nil {
		return *m.Tradeno
	}
	return ""
}

func (m *SOrderQuery) GetTransactionid() string {
	if m != nil && m.Transactionid != nil {
		return *m.Transactionid
	}
	return ""
}

func (m *SOrderQuery) GetPaytime() string {
	if m != nil && m.Paytime != nil {
		return *m.Paytime
	}
	return ""
}

func (m *SOrderQuery) GetNum() uint32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *SOrderQuery) GetSign() string {
	if m != nil && m.Sign != nil {
		return *m.Sign
	}
	return ""
}

// 房卡变动推送
type SRoomCardChange struct {
	Error            *uint32 `protobuf:"varint,1,opt" json:"Error,omitempty"`
	Value            *int32  `protobuf:"varint,2,opt" json:"Value,omitempty"`
	Left             *uint32 `protobuf:"varint,3,opt" json:"Left,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SRoomCardChange) Reset()         { *m = SRoomCardChange{} }
func (m *SRoomCardChange) String() string { return proto.CompactTextString(m) }
func (*SRoomCardChange) ProtoMessage()    {}

func (m *SRoomCardChange) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func (m *SRoomCardChange) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *SRoomCardChange) GetLeft() uint32 {
	if m != nil && m.Left != nil {
		return *m.Left
	}
	return 0
}

func init() {
}
