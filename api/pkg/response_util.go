package pkg

import (
	"api/constant"
	utils "api/utils/response"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) utils.Response[T] {
	return BuildResponse_(responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](message string, data T) utils.Response[T] {
	return utils.Response[T]{
		Message: message,
		Data:    data,
	}
}

func BuildResponse__(message string) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
	}
}
