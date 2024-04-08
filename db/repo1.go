package db

// import (
// 	"reflect"

// 	"gorm.io/gorm"
// )

// type CreateSchema[T any] interface{}
// type UpdateSchema[T any] interface{}

// type CrudBase[ModelT any] struct {
// 	model reflect.Type
// 	db *gorm.DB
// }

// func NewCrud[ModelT any](db *gorm.DB, modelType reflect.Type) *CrudBase[ModelT] {
// 	return &CrudBase[ModelT]{
// 		model: reflect.TypeOf(new(ModelT)).Elem(),
// 		db: db,
// 	}
// }

// func ()