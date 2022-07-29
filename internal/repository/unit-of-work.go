package repository

import "gorm.io/gorm"

type UnitOfWork interface {
	Begin() (*gorm.DB, error)
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error
}

type GormUnitOfWork struct {
	db *gorm.DB
}

// NewGormUnitOfWork will create a new gorm unit of work
func NewGormUnitOfWork(db *gorm.DB) UnitOfWork {
	return &GormUnitOfWork{db: db}
}

func (u *GormUnitOfWork) Begin() (*gorm.DB, error) {
	tx := u.db.Begin()
	return tx, tx.Error
}

func (u *GormUnitOfWork) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (u *GormUnitOfWork) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}
