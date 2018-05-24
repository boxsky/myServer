// Code generated by protoc-gen-go.
// source: login_broadcast.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

// 公告
type CNotice struct {
	Appid            *uint32 `protobuf:"varint,1,opt" json:"Appid,omitempty"`
	Typeid           *uint32 `protobuf:"varint,2,opt" json:"Typeid,omitempty"`
	Uid              *uint32 `protobuf:"varint,3,opt" json:"Uid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CNotice) Reset()         { *m = CNotice{} }
func (m *CNotice) String() string { return proto.CompactTextString(m) }
func (*CNotice) ProtoMessage()    {}

func (m *CNotice) GetAppid() uint32 {
	if m != nil && m.Appid != nil {
		return *m.Appid
	}
	return 0
}

func (m *CNotice) GetTypeid() uint32 {
	if m != nil && m.Typeid != nil {
		return *m.Typeid
	}
	return 0
}

func (m *CNotice) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

// 公告响应
type SNotice struct {
	State            *uint32 `protobuf:"varint,1,opt" json:"State,omitempty"`
	Data             *string `protobuf:"bytes,2,opt" json:"Data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SNotice) Reset()         { *m = SNotice{} }
func (m *SNotice) String() string { return proto.CompactTextString(m) }
func (*SNotice) ProtoMessage()    {}

func (m *SNotice) GetState() uint32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

func (m *SNotice) GetData() string {
	if m != nil && m.Data != nil {
		return *m.Data
	}
	return ""
}

// 获取跑马灯信息,每次切换到大厅请求一次数据
type CRollingSubtitle struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CRollingSubtitle) Reset()         { *m = CRollingSubtitle{} }
func (m *CRollingSubtitle) String() string { return proto.CompactTextString(m) }
func (*CRollingSubtitle) ProtoMessage()    {}

type SRollingInfo struct {
	ID               *int32  `protobuf:"varint,1,opt" json:"ID,omitempty"`
	Content          *string `protobuf:"bytes,2,opt" json:"Content,omitempty"`
	ValidendTime     *string `protobuf:"bytes,3,opt" json:"ValidendTime,omitempty"`
	Frequency        *int32  `protobuf:"varint,4,opt" json:"Frequency,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SRollingInfo) Reset()         { *m = SRollingInfo{} }
func (m *SRollingInfo) String() string { return proto.CompactTextString(m) }
func (*SRollingInfo) ProtoMessage()    {}

func (m *SRollingInfo) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *SRollingInfo) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func (m *SRollingInfo) GetValidendTime() string {
	if m != nil && m.ValidendTime != nil {
		return *m.ValidendTime
	}
	return ""
}

func (m *SRollingInfo) GetFrequency() int32 {
	if m != nil && m.Frequency != nil {
		return *m.Frequency
	}
	return 0
}

// 返回跑马灯信息
type SRollingSubtitle struct {
	Error            *uint32         `protobuf:"varint,1,opt" json:"Error,omitempty"`
	RollingInfos     []*SRollingInfo `protobuf:"bytes,2,rep" json:"RollingInfos,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SRollingSubtitle) Reset()         { *m = SRollingSubtitle{} }
func (m *SRollingSubtitle) String() string { return proto.CompactTextString(m) }
func (*SRollingSubtitle) ProtoMessage()    {}

func (m *SRollingSubtitle) GetError() uint32 {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return 0
}

func (m *SRollingSubtitle) GetRollingInfos() []*SRollingInfo {
	if m != nil {
		return m.RollingInfos
	}
	return nil
}

// 推送消息
type SPushMessage struct {
	Uid              *int32  `protobuf:"varint,1,opt" json:"Uid,omitempty"`
	Type             *uint32 `protobuf:"varint,2,opt" json:"Type,omitempty"`
	ID               *int32  `protobuf:"varint,3,opt" json:"ID,omitempty"`
	Title            *string `protobuf:"bytes,4,opt" json:"Title,omitempty"`
	Gold             *int64  `protobuf:"varint,5,opt" json:"Gold,omitempty"`
	CarryGold        *int64  `protobuf:"varint,6,opt" json:"CarryGold,omitempty"`
	SafeGold         *int64  `protobuf:"varint,7,opt" json:"SafeGold,omitempty"`
	RoomCard         *int32  `protobuf:"varint,8,opt" json:"RoomCard,omitempty"`
	ValidendTime     *string `protobuf:"bytes,9,opt" json:"ValidendTime,omitempty"`
	Frequency        *int32  `protobuf:"varint,10,opt" json:"Frequency,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SPushMessage) Reset()         { *m = SPushMessage{} }
func (m *SPushMessage) String() string { return proto.CompactTextString(m) }
func (*SPushMessage) ProtoMessage()    {}

func (m *SPushMessage) GetUid() int32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *SPushMessage) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SPushMessage) GetID() int32 {
	if m != nil && m.ID != nil {
		return *m.ID
	}
	return 0
}

func (m *SPushMessage) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *SPushMessage) GetGold() int64 {
	if m != nil && m.Gold != nil {
		return *m.Gold
	}
	return 0
}

func (m *SPushMessage) GetCarryGold() int64 {
	if m != nil && m.CarryGold != nil {
		return *m.CarryGold
	}
	return 0
}

func (m *SPushMessage) GetSafeGold() int64 {
	if m != nil && m.SafeGold != nil {
		return *m.SafeGold
	}
	return 0
}

func (m *SPushMessage) GetRoomCard() int32 {
	if m != nil && m.RoomCard != nil {
		return *m.RoomCard
	}
	return 0
}

func (m *SPushMessage) GetValidendTime() string {
	if m != nil && m.ValidendTime != nil {
		return *m.ValidendTime
	}
	return ""
}

func (m *SPushMessage) GetFrequency() int32 {
	if m != nil && m.Frequency != nil {
		return *m.Frequency
	}
	return 0
}

func init() {
}
