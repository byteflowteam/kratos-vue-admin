package constant

import (
	"encoding/json"
	"errors"
)

const NormalErrPrefix = "NORMAL_ERROR@"

// key = errCode, string = errMsg
type ErrInfo struct {
	ErrCode int32
	ErrMsg  string
}

var (
	OK        = ErrInfo{0, ""}
	ErrServer = ErrInfo{500, "server error"}

	ErrParams = ErrInfo{ErrCode: 400, ErrMsg: "接口参数错误"}

	ErrParseToken               = ErrInfo{700, "parse token failed"}
	ErrTokenExpired             = ErrInfo{701, "token is timed out, please log in again"}
	ErrTokenInvalid             = ErrInfo{702, "token is invalid"}
	ErrTokenMalformed           = ErrInfo{703, "that's not even a token"}
	ErrTokenNotValidYet         = ErrInfo{704, "token has been invalidated"}
	ErrTokenUnknown             = ErrInfo{705, "that's not even a token"}
	ErrTokenKicked              = ErrInfo{706, "user has been kicked"}
	ErrTokenDifferentPlatformID = ErrInfo{707, "different platformID"}
	ErrTokenDifferentUserID     = ErrInfo{708, "different userID"}

	ErrAccess                = ErrInfo{ErrCode: 801, ErrMsg: "no permission"}
	ErrDB                    = ErrInfo{ErrCode: 802, ErrMsg: "db failed"}
	ErrArgs                  = ErrInfo{ErrCode: 803, ErrMsg: "args failed"}
	ErrStatus                = ErrInfo{ErrCode: 804, ErrMsg: "status is abnormal"}
	ErrCallback              = ErrInfo{ErrCode: 809, ErrMsg: "callback failed"}
	ErrSendLimit             = ErrInfo{ErrCode: 810, ErrMsg: "send msg limit, to many request, try again later"}
	ErrMessageHasReadDisable = ErrInfo{ErrCode: 811, ErrMsg: "message has read disable"}
	ErrInternal              = ErrInfo{ErrCode: 812, ErrMsg: "internal error"}
	ErrWsConnNotExist        = ErrInfo{ErrCode: 813, ErrMsg: "ws conn not exist"}

	ErrAccountHasRegistered  = ErrInfo{ErrCode: 10002, ErrMsg: "account has been registered"}
	ErrAccountHasNotRegister = ErrInfo{ErrCode: 10003, ErrMsg: "account have not register"}
	ErrPasswordError         = ErrInfo{ErrCode: 10004, ErrMsg: "password error"}
	ErrUpdateError           = ErrInfo{ErrCode: 10005, ErrMsg: "update error"}
	ErrRepeatedError         = ErrInfo{ErrCode: 10005, ErrMsg: "repeated error"}
)

func (e ErrInfo) Error() string {
	return e.ErrMsg
}

func (e *ErrInfo) Code() int32 {
	return e.ErrCode
}

func (e ErrInfo) ToString() error {
	errByte, _ := json.Marshal(e)
	errStr := NormalErrPrefix + string(errByte)
	return errors.New(errStr)
}
