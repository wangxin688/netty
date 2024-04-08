package types

import (
	"time"

)

type Order string


const (
	ASC  Order = "ascend"
	DESC Order = "descend"
)

type ValidationError struct {
	Field string
	Value any
}

type QueryParams struct {
	Order   Order  `form:"order" binding:"omitempty,oneof=ascend descend"`
	Orderby string `form:"order_by" binding:"omitempty"`
	Limit   uint   `form:"limit,default=10" binding:"omitempty,gte=0"`
	Offset  uint   `form:"offset,default=0" binding:"omitempty,gte=0"`
	Id      uint    `form:"id" binding:"omitempty,gt=0"`
}

type AuditTimeQuery struct {
	CreatedAt__gte time.Time `form:"created_at__gte" binding:"omitempty"`
	CreatedAt__lte time.Time `form:"created_at__lte" binding:"omitempty"`
	UpdatedAt__gte time.Time `form:"updated_at__gte" binding:"omitempty"`
	UpdatedAt__lte time.Time `form:"updated_at__lte" binding:"omitempty"`
}

type AuditTime struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

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
