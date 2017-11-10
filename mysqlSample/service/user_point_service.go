package service

import (
	"github.com/goSamples/mysqlSample/model"
	"github.com/goSamples/mysqlSample/model/datasource"
)

// UserService is
type UserService struct {
}

// NewUserService is
func NewUserService() *UserService {
	return &UserService{}
}

// GetTotalPointByID is
func (ups *UserService) GetTotalPointByID(id int) int {
	mysql := datasource.NewMySQL()
	r := model.NewUserPointRepositoryImpl(mysql)
	result := r.GetTotalPointByUserID(id)
	return result
}
