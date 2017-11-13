package service

import (
	"fmt"
	"reflect"

	"github.com/goSamples/mysqlSample/model"
	"github.com/goSamples/mysqlSample/model/datasource"
)

var dbSet = datasource.NewMySQL

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
	// p, ok := db.(datasource.DataBase)
	// fmt.Printf("p: %v\n", &p)
	// fmt.Printf("ok: %v\n", ok)
	// fmt.Printf("aaaaaaaaaa: %v\n", dbSet().SumPointByUserID(1))
	fmt.Printf("[GetTotalPointByID] db type: %v\n", reflect.TypeOf(db))
	r := model.NewUserPointRepositoryImpl(db)
	result := r.GetTotalPointByUserID(id)
	return result
}
