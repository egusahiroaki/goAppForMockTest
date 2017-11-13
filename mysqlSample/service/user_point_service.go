package service

import (
	"fmt"
	"reflect"

	"github.com/goSamples/mysqlSample/model"
	"github.com/goSamples/mysqlSample/model/datasource"
)

var dbSet = datasource.SetDB

// UserService is
type UserService struct {
}

// NewUserService is
func NewUserService() *UserService {
	return &UserService{}
}

// GetTotalPointByID is
func (ups *UserService) GetTotalPointByID(id int) int {
	db := dbSet()
	// fmt.Printf("aaaaaaaaaa: %v\n", dbSet().SumPointByUserID(1))
	fmt.Printf("db type: %v\n", reflect.TypeOf(db))
	r := model.NewUserPointRepositoryImpl(db)
	result := r.GetTotalPointByUserID(id)
	return result
}
