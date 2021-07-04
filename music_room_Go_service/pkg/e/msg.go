package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	NOT_FOUND:                      "没有找到房间",
	ERROR_EXIST_CODE:               "已存在该房间号",
	ERROR_NOT_EXIST_CODE:           "该房间号不存在",
	ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	NOT_HOST_OF_ROOM:               "你不是房间的主人",
	NOT_ROOMCODE_OF_ROOM:           "没有根据房间号找到房间",
	NOT_CONTENT:                    "没有内容",
	FORBIDDEN_PAUSE_OR_PLAY_SONG:   "没有权限（播放和暂停）",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
