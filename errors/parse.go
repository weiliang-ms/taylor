package errors

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 入参：error实例、http状态码，解析返回错误对象
func ParseHttpError(err error, statusCode int) error {
	// 更多http状态码
	// https://cloud.tencent.com/developer/chapter/13553
	switch statusCode {
	// 404对应错误
	case http.StatusNotFound:
		err = NotFound(err)
	// 400对应错误
	case http.StatusBadRequest:
		err = InvalidParameter(err)
	// 409对应错误
	// 冲突最有可能发生以回应PUT请求。
	// 例如，在上传比服务器上已存在的文件更早的文件时，您可能会收到409响应，从而导致版本控制冲突。
	case http.StatusConflict:
		err = Conflict(err)
	// 401未经授权错误
	case http.StatusUnauthorized:
		err = Unauthorized(err)
	// 503服务器尚未准备好处理请求
	// 常见原因是服务器因维护或重载而停机
	case http.StatusServiceUnavailable:
		err = Unavailable(err)
	// 403
	// 客户端错误状态响应代码指示服务器理解请求但拒绝授权。
	case http.StatusForbidden:
		err = Forbidden(err)
	// 304
	// 客户端重定向响应代码指示不需要重新传输请求的资源。
	case http.StatusNotModified:
		err = NotModified(err)
	// 501
	// 服务器错误响应代码指示请求方法不受服务器支持并且无法处理
	case http.StatusNotImplemented:
		err = NotImplemented(err)
	// 500
	// 服务器错误响应代码指示服务器遇到阻止它履行请求的意外情况。
	case http.StatusInternalServerError:
		if !IsSystem(err) && !IsUnknown(err) && !IsDataLoss(err) && !IsDeadline(err) && !IsCancelled(err) {
			err = System(err)
		}

		// 不匹配以上状态码情况
	default:
		// 记录日志
		logrus.WithFields(logrus.Fields{
			"module":      "api",
			"status_code": fmt.Sprintf("%d", statusCode),
		}).Debugf("FIXME: Got an status-code for which error does not match any expected type!!!: %d", statusCode)

		// 判断状态码值，进行分类
		switch {
		// 200 <= statusCode < 400，属于客户端错误
		case statusCode >= 200 && statusCode < 400:
			// it's a client error
		// 400 <= statusCode < 500，属于请求非法参数错误
		case statusCode >= 400 && statusCode < 500:
			err = InvalidParameter(err)
		// 500 <= statusCode < 600 属于系统错误（服务端）
		case statusCode >= 500 && statusCode < 600:
			err = System(err)
		// 未知错误
		default:
			err = Unknown(err)
		}
	}
	return err
}

func ParseFileError(err error) {
	if err != nil {
		err.Error()
	}
}
