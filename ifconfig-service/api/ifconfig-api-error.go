package api

type ApiError interface {
    Info() ApiErrorInfo
}

type ApiErrorInfo struct {
    Message string
    Code int
}

func NewError(message string, code int) ApiError {
    return &ApiErrorInfo{Message: message, Code: code}
}

func (e *ApiErrorInfo) Info() ApiErrorInfo {
    return *e
}
