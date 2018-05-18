// Code generated by protoc-gen-go.
// source: cmd.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type CMD int32

const (
	CMD_HEARTBEAT CMD = 1
	// 网管和内部服务器消息协议
	CMD_GATEWAY_REGIST_TO_INTERNAL    CMD = 100
	CMD_SERVICEREGIST                 CMD = 101
	CMD_GATEWAYREGIST                 CMD = 105
	CMD_USERLINEOFFNOTIFY             CMD = 111
	CMD_GATEWAYMIGRATENOTIFY          CMD = 121
	CMD_UPDATEGAMEUSERSTATUS          CMD = 122
	CMD_CMD_GATEWAY_MAX               CMD = 999
	CMD_AMPQ_MIN                      CMD = 1501
	CMD_AMPQTOSINGLEUSER              CMD = 1502
	CMD_AMPQRECV                      CMD = 1503
	CMD_AMPQ_REGIST                   CMD = 1504
	CMD_AMPQ_UNREGIST                 CMD = 1505
	CMD_AMPQ_APP_SERVER_REGIST        CMD = 1506
	CMD_AMPQ_APP_SERVER_REGIST_RESULT CMD = 1512
	CMD_AMPQ_APP_SERVER_REGIST_NOTIFY CMD = 1507
	CMD_AMPQ_ROOM_INFO_NOTIFY         CMD = 1508
	CMD_AMPQ_FAIL                     CMD = 1509
	CMD_AMPQ_MAX                      CMD = 1510
	CMD_DETECT_CURRENT_ROOM           CMD = 1511
	CMD_GET_ROOM_PLAYERS              CMD = 1513
	CMD_AMPQ_APP_UPDATE_SERVER_INFO   CMD = 1514
	CMD_AMPQ_UPDATE_GAME_CONFIG       CMD = 1515
	CMD_GET__PLATFORM_PLAYER_COUNT    CMD = 1516
	CMD_KICK_PLAYER_OFFLINE           CMD = 1517
	CMD_UPDATE_REVENUE_INFO           CMD = 1518
	CMD_UPDATE_ROOM_PLAYER_INFO       CMD = 1519
	CMD_CHANGE_GAME_SERVER_STATUS     CMD = 1520
	CMD_REF_GAME_CONFIG               CMD = 1521
	CMD_UPDATE_SYS_PRIZE_POOL         CMD = 1522
	CMD_GET_App_ROOMLIST              CMD = 1523
	CMD_NOTICE_CHARGE_TO_GS           CMD = 1524
	CMD_CMD_CONSOLE_MAX               CMD = 1999
)

var CMD_name = map[int32]string{
	1:    "HEARTBEAT",
	100:  "GATEWAY_REGIST_TO_INTERNAL",
	101:  "SERVICEREGIST",
	105:  "GATEWAYREGIST",
	111:  "USERLINEOFFNOTIFY",
	121:  "GATEWAYMIGRATENOTIFY",
	122:  "UPDATEGAMEUSERSTATUS",
	999:  "CMD_GATEWAY_MAX",
	1501: "AMPQ_MIN",
	1502: "AMPQTOSINGLEUSER",
	1503: "AMPQRECV",
	1504: "AMPQ_REGIST",
	1505: "AMPQ_UNREGIST",
	1506: "AMPQ_APP_SERVER_REGIST",
	1512: "AMPQ_APP_SERVER_REGIST_RESULT",
	1507: "AMPQ_APP_SERVER_REGIST_NOTIFY",
	1508: "AMPQ_ROOM_INFO_NOTIFY",
	1509: "AMPQ_FAIL",
	1510: "AMPQ_MAX",
	1511: "DETECT_CURRENT_ROOM",
	1513: "GET_ROOM_PLAYERS",
	1514: "AMPQ_APP_UPDATE_SERVER_INFO",
	1515: "AMPQ_UPDATE_GAME_CONFIG",
	1516: "GET__PLATFORM_PLAYER_COUNT",
	1517: "KICK_PLAYER_OFFLINE",
	1518: "UPDATE_REVENUE_INFO",
	1519: "UPDATE_ROOM_PLAYER_INFO",
	1520: "CHANGE_GAME_SERVER_STATUS",
	1521: "REF_GAME_CONFIG",
	1522: "UPDATE_SYS_PRIZE_POOL",
	1523: "GET_App_ROOMLIST",
	1524: "NOTICE_CHARGE_TO_GS",
	1999: "CMD_CONSOLE_MAX",
}
var CMD_value = map[string]int32{
	"HEARTBEAT":                     1,
	"GATEWAY_REGIST_TO_INTERNAL":    100,
	"SERVICEREGIST":                 101,
	"GATEWAYREGIST":                 105,
	"USERLINEOFFNOTIFY":             111,
	"GATEWAYMIGRATENOTIFY":          121,
	"UPDATEGAMEUSERSTATUS":          122,
	"CMD_GATEWAY_MAX":               999,
	"AMPQ_MIN":                      1501,
	"AMPQTOSINGLEUSER":              1502,
	"AMPQRECV":                      1503,
	"AMPQ_REGIST":                   1504,
	"AMPQ_UNREGIST":                 1505,
	"AMPQ_APP_SERVER_REGIST":        1506,
	"AMPQ_APP_SERVER_REGIST_RESULT": 1512,
	"AMPQ_APP_SERVER_REGIST_NOTIFY": 1507,
	"AMPQ_ROOM_INFO_NOTIFY":         1508,
	"AMPQ_FAIL":                     1509,
	"AMPQ_MAX":                      1510,
	"DETECT_CURRENT_ROOM":           1511,
	"GET_ROOM_PLAYERS":              1513,
	"AMPQ_APP_UPDATE_SERVER_INFO":   1514,
	"AMPQ_UPDATE_GAME_CONFIG":       1515,
	"GET__PLATFORM_PLAYER_COUNT":    1516,
	"KICK_PLAYER_OFFLINE":           1517,
	"UPDATE_REVENUE_INFO":           1518,
	"UPDATE_ROOM_PLAYER_INFO":       1519,
	"CHANGE_GAME_SERVER_STATUS":     1520,
	"REF_GAME_CONFIG":               1521,
	"UPDATE_SYS_PRIZE_POOL":         1522,
	"GET_App_ROOMLIST":              1523,
	"NOTICE_CHARGE_TO_GS":           1524,
	"CMD_CONSOLE_MAX":               1999,
}

func (x CMD) Enum() *CMD {
	p := new(CMD)
	*p = x
	return p
}
func (x CMD) String() string {
	return proto.EnumName(CMD_name, int32(x))
}
func (x *CMD) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(CMD_value, data, "CMD")
	if err != nil {
		return err
	}
	*x = CMD(value)
	return nil
}

func init() {
	proto.RegisterEnum("protocol.CMD", CMD_name, CMD_value)
}