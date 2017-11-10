package service

import "github.com/goSamples/mysqlSample/model"

// UserService is
type UserService struct {
}

// NewUserService is
func NewUserService() *UserService {
	return &UserService{}
}

// GetTotalPointByID is
func (ups UserService) GetTotalPointByID(id int) int {
	var r model.UserPointRepository
	r = model.NewUserPointRepository()
	result := r.GetTotalPointByUserID(id)
	return result
}
