package model

import "errors"

// 自定义一些错误常量
// 服务器发送给客户端的相关错误
var (
	ERROR_USER_NOTEXISTS   = errors.New("用户不存在")
	ERROR_USER_EXISTS      = errors.New("用户ID已经存在")
	ERROR_USER_PWD         = errors.New("密码不正确")

	ERROR_REDIS_DO_FAILED  = errors.New("redis操作失败")
	ERROR_JSON_MARSHAL     = errors.New("json解析出错")
	ERROR_UNKNOWN          = errors.New("服务器未知错误")
)

// 服务器本身会遇到和处理的错误
var (
	ERROR_CLIENT_CLOSE_CONNECTION         = errors.New("客户端断开了与服务器的连接")
	ERROR_CLIENT_CONNECT_TIMEOUT          = errors.New("客户端连接超时")
	ERROR_READ_DATA_FROM_CLIENT_TIMEOUT   = errors.New("读取客户端的消息超时")
	ERROR_WRITE_DATA_TO_CLIENT_TIMEOUT    = errors.New("向客户端发送消息超时")
)

// 客户端本身会遇到和处理的错误
var (
	ERROR_SERVER_CLOSE_CONNECTION = errors.New("服务器主动断开了与客户端的连接")
)

