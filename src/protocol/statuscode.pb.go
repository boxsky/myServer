// Code generated by protoc-gen-go.
// source: statuscode.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type StatusCode int32

const (
	StatusCode_Success                  StatusCode = 0
	StatusCode_GetUserDataFail          StatusCode = 10001
	StatusCode_GetLessUserDataFail      StatusCode = 10002
	StatusCode_FeedBackFail             StatusCode = 10003
	StatusCode_GetGloryScoresFail       StatusCode = 10005
	StatusCode_GetGloryScoresDetailFail StatusCode = 10006
	StatusCode_GetRollSubstitleFail     StatusCode = 10007
	StatusCode_LackOfRoomCard           StatusCode = 10008
	// 代理平台状态码
	StatusCode_PlatformFail                  StatusCode = 1
	StatusCode_PlatformUserLess              StatusCode = 2
	StatusCode_PlatformUserNotLogin          StatusCode = 3
	StatusCode_PlatformInvalidSign           StatusCode = 4
	StatusCode_PlatformInvalidParameter      StatusCode = 5
	StatusCode_PlatformInvalidActivity       StatusCode = 6
	StatusCode_PlatformLackDCoin             StatusCode = 7
	StatusCode_PlatformInvalidAppId          StatusCode = 8
	StatusCode_PlatformInvalidAppKey         StatusCode = 9
	StatusCode_PlatformRequestIllegal        StatusCode = 10
	StatusCode_PlatformNotPrize              StatusCode = 11
	StatusCode_PlatformAlreadyPrize          StatusCode = 12
	StatusCode_PlatformGivePrize             StatusCode = 13
	StatusCode_PlatformFinishActivity        StatusCode = 14
	StatusCode_PlatformInvalidBattle         StatusCode = 15
	StatusCode_PlatformLimitSignUp           StatusCode = 16
	StatusCode_PlatformShare                 StatusCode = 17
	StatusCode_PlatformLoginFail             StatusCode = 100
	StatusCode_PlatformInvalid               StatusCode = 101
	StatusCode_PlatformLoginVeoverdue        StatusCode = 102
	StatusCode_PlatformInvalidToken          StatusCode = 103
	StatusCode_PlatformPaymentFail           StatusCode = 200
	StatusCode_PlatformInvalidTradeno        StatusCode = 201
	StatusCode_PlatformNotPayTradeno         StatusCode = 202
	StatusCode_PlatformMismathchTradenoMoney StatusCode = 203
	StatusCode_PlatformInvalidProduct        StatusCode = 204
	StatusCode_PlatformInviteCodeFail        StatusCode = 300
	StatusCode_PlatformInvalidGame           StatusCode = 301
	StatusCode_PlatformInvalidInviteCode     StatusCode = 303
	StatusCode_PlatformBeinputInviteCode     StatusCode = 304
	StatusCode_PlatformNotInputInviteCode    StatusCode = 305
	// 房间错误码
	StatusCode_OPERATE_IELLEGAL          StatusCode = 20000
	StatusCode_GAME_STARTED_CANNOT_LEAVE StatusCode = 20001
	StatusCode_USER_DATA_ERROR           StatusCode = 20004
	StatusCode_NOT_IN_ROOM               StatusCode = 20007
	// 创建房间
	StatusCode_CREAT_ROOM_FAIL                     StatusCode = 20009
	StatusCode_CREAT_ROOM_FAIL_IN_ROOM             StatusCode = 20096
	StatusCode_CREAT_ROOM_FAIL_DIAMOND_NOT_ENOUGH  StatusCode = 20097
	StatusCode_CREAT_ROOM_FAIL_ROOMCARD_NOT_ENOUGH StatusCode = 20098
	StatusCode_STOP_GAME_SERVERING                 StatusCode = 20099
	// 进入房间
	StatusCode_ROOM_FULL               StatusCode = 20011
	StatusCode_ROOM_NOT_EXIST          StatusCode = 20012
	StatusCode_USER_GOLD_NOT_ENOUGH    StatusCode = 20013
	StatusCode_USER_DIAMOND_NOT_ENOUGH StatusCode = 20014
	StatusCode_GET_ROOM_LIST_FAIL      StatusCode = 20050
)

var StatusCode_name = map[int32]string{
	0:     "Success",
	10001: "GetUserDataFail",
	10002: "GetLessUserDataFail",
	10003: "FeedBackFail",
	10005: "GetGloryScoresFail",
	10006: "GetGloryScoresDetailFail",
	10007: "GetRollSubstitleFail",
	10008: "LackOfRoomCard",
	1:     "PlatformFail",
	2:     "PlatformUserLess",
	3:     "PlatformUserNotLogin",
	4:     "PlatformInvalidSign",
	5:     "PlatformInvalidParameter",
	6:     "PlatformInvalidActivity",
	7:     "PlatformLackDCoin",
	8:     "PlatformInvalidAppId",
	9:     "PlatformInvalidAppKey",
	10:    "PlatformRequestIllegal",
	11:    "PlatformNotPrize",
	12:    "PlatformAlreadyPrize",
	13:    "PlatformGivePrize",
	14:    "PlatformFinishActivity",
	15:    "PlatformInvalidBattle",
	16:    "PlatformLimitSignUp",
	17:    "PlatformShare",
	100:   "PlatformLoginFail",
	101:   "PlatformInvalid",
	102:   "PlatformLoginVeoverdue",
	103:   "PlatformInvalidToken",
	200:   "PlatformPaymentFail",
	201:   "PlatformInvalidTradeno",
	202:   "PlatformNotPayTradeno",
	203:   "PlatformMismathchTradenoMoney",
	204:   "PlatformInvalidProduct",
	300:   "PlatformInviteCodeFail",
	301:   "PlatformInvalidGame",
	303:   "PlatformInvalidInviteCode",
	304:   "PlatformBeinputInviteCode",
	305:   "PlatformNotInputInviteCode",
	20000: "OPERATE_IELLEGAL",
	20001: "GAME_STARTED_CANNOT_LEAVE",
	20004: "USER_DATA_ERROR",
	20007: "NOT_IN_ROOM",
	20009: "CREAT_ROOM_FAIL",
	20096: "CREAT_ROOM_FAIL_IN_ROOM",
	20097: "CREAT_ROOM_FAIL_DIAMOND_NOT_ENOUGH",
	20098: "CREAT_ROOM_FAIL_ROOMCARD_NOT_ENOUGH",
	20099: "STOP_GAME_SERVERING",
	20011: "ROOM_FULL",
	20012: "ROOM_NOT_EXIST",
	20013: "USER_GOLD_NOT_ENOUGH",
	20014: "USER_DIAMOND_NOT_ENOUGH",
	20050: "GET_ROOM_LIST_FAIL",
}
var StatusCode_value = map[string]int32{
	"Success":                             0,
	"GetUserDataFail":                     10001,
	"GetLessUserDataFail":                 10002,
	"FeedBackFail":                        10003,
	"GetGloryScoresFail":                  10005,
	"GetGloryScoresDetailFail":            10006,
	"GetRollSubstitleFail":                10007,
	"LackOfRoomCard":                      10008,
	"PlatformFail":                        1,
	"PlatformUserLess":                    2,
	"PlatformUserNotLogin":                3,
	"PlatformInvalidSign":                 4,
	"PlatformInvalidParameter":            5,
	"PlatformInvalidActivity":             6,
	"PlatformLackDCoin":                   7,
	"PlatformInvalidAppId":                8,
	"PlatformInvalidAppKey":               9,
	"PlatformRequestIllegal":              10,
	"PlatformNotPrize":                    11,
	"PlatformAlreadyPrize":                12,
	"PlatformGivePrize":                   13,
	"PlatformFinishActivity":              14,
	"PlatformInvalidBattle":               15,
	"PlatformLimitSignUp":                 16,
	"PlatformShare":                       17,
	"PlatformLoginFail":                   100,
	"PlatformInvalid":                     101,
	"PlatformLoginVeoverdue":              102,
	"PlatformInvalidToken":                103,
	"PlatformPaymentFail":                 200,
	"PlatformInvalidTradeno":              201,
	"PlatformNotPayTradeno":               202,
	"PlatformMismathchTradenoMoney":       203,
	"PlatformInvalidProduct":              204,
	"PlatformInviteCodeFail":              300,
	"PlatformInvalidGame":                 301,
	"PlatformInvalidInviteCode":           303,
	"PlatformBeinputInviteCode":           304,
	"PlatformNotInputInviteCode":          305,
	"OPERATE_IELLEGAL":                    20000,
	"GAME_STARTED_CANNOT_LEAVE":           20001,
	"USER_DATA_ERROR":                     20004,
	"NOT_IN_ROOM":                         20007,
	"CREAT_ROOM_FAIL":                     20009,
	"CREAT_ROOM_FAIL_IN_ROOM":             20096,
	"CREAT_ROOM_FAIL_DIAMOND_NOT_ENOUGH":  20097,
	"CREAT_ROOM_FAIL_ROOMCARD_NOT_ENOUGH": 20098,
	"STOP_GAME_SERVERING":                 20099,
	"ROOM_FULL":                           20011,
	"ROOM_NOT_EXIST":                      20012,
	"USER_GOLD_NOT_ENOUGH":                20013,
	"USER_DIAMOND_NOT_ENOUGH":             20014,
	"GET_ROOM_LIST_FAIL":                  20050,
}

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}
func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (x *StatusCode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(StatusCode_value, data, "StatusCode")
	if err != nil {
		return err
	}
	*x = StatusCode(value)
	return nil
}

func init() {
	proto.RegisterEnum("protocol.StatusCode", StatusCode_name, StatusCode_value)
}
