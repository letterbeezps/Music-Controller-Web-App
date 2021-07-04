package e

// 定义错误编码
const (
	SUCCESS                        = 200
	ERROR                          = 500
	INVALID_PARAMS                 = 400
	NOT_FOUND                      = 404
	ERROR_EXIST_CODE               = 10001
	ERROR_NOT_EXIST_CODE           = 10002
	ERROR_NOT_EXIST_ARTICLE        = 10003
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	NOT_HOST_OF_ROOM               = 20005
	NOT_ROOMCODE_OF_ROOM           = 20006
	NOT_CONTENT                    = 20007
	FORBIDDEN_PAUSE_OR_PLAY_SONG   = 20008
)
