package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/goSamples/mysqlSample/model/datasource"
)

// // UserPointRepository is
// type UserPointRepository interface {
// 	GetTotalPointByUserID(userID int) int
// }

// UserPointRepositoryImpl is
type UserPointRepositoryImpl struct {
	db datasource.DataBase
}

// NewUserPointRepositoryImpl is
func NewUserPointRepositoryImpl(db datasource.DataBase) UserPointRepositoryImpl {
	return UserPointRepositoryImpl{db: db}
}

// GetTotalPointByUserID is
// [todo] return should error
func (u *UserPointRepositoryImpl) GetTotalPointByUserID(userID int) int {
	return u.db.SumPointByUserID(userID)
}
