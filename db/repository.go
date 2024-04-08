package db

// import (
// 	"context"

// 	"gorm.io/gorm"
// )

// // https://github.com/Ompluscator/gorm-generics/blob/master/repository.go

// type GormModel[E any] interface {
// 	ToEntity() E
// 	FromEntity(entity E) interface{}
// }

// func NewCrud[M GormModel[E], E any](db *gorm.DB) *CrudBase[M, E] {
// 	return &CrudBase[M, E]{
// 		db: db,
// 	}
// }

// type CrudBase1[M GormModel[E], E any] struct {
// 	db *gorm.DB
// }

// func (c *CrudBase1[M, E]) Create(ctx context.Context, entity *E) error {
// 	var start M

// 	model := start.FromEntity(*entity).(M)

// 	err := c.db.WithContext(ctx).Create(&model).Error

// 	if err != nil {
// 		return err
// 	}
// 	*entity = model.ToEntity()
// 	return nil
// }

// func ()