package response

type ListT[T any] struct {
	Count   int `json:"count"`
	Results []T `json:"results"`
}

type BaseResponse[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type BaseListResponse[T any] struct {
	Code    int      `json:"code"`
	Data    ListT[T] `json:"data,omitempty"`
	Message string   `json:"message"`
}

func SetResponse[T any](code int, data T, message string) BaseResponse[T] {
	return BaseResponse[T]{
		Code:    code,
		Data:    data,
		Message: message,
	}
}

func SetListResponse[T any](code int, data ListT[T], message string) BaseListResponse[T] {
	return BaseListResponse[T]{
		Code:    code,
		Data:    data,
		Message: message,
	}
}
