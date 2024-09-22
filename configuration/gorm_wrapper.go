package configuration

import (
	"context"

	"gorm.io/gorm"
)

type GormWrapper interface {
	WithContext(ctx context.Context) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Preload(rel string, args ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Updates(values interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
}

type GormDB struct {
	*gorm.DB
}

func (g *GormDB) WithContext(ctx context.Context) *gorm.DB {
	return g.DB.WithContext(ctx)
}
