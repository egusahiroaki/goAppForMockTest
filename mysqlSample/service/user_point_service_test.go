package service

import (
	"fmt"
	"testing"

	"github.com/goSamples/mysqlSample/model/datasource"
	"github.com/goSamples/mysqlSample/service"
)

type DBMock struct{}

func (db DBMock) SumPointByUserID(_ int) int {
	return 10
}

func setMockDB(db datasource.DB) {
	fmt.Printf("[setMockDB]\n")
	fmt.Printf("db: %v\n", db)
	dbSet = func() datasource.DB { return db }
	// fmt.Printf("dbSet: %v\n", dbSet().SumPointByUserID(1))
}

func TestUserServiceGetTotalPointByID(t *testing.T) {
	//	dbmock := DBMock{}
	setMockDB(DBMock{})
	// fmt.Printf("hogehoge dbSet: %v\n", dbSet().SumPointByUserID(1))
	s := service.NewUserService()
	point := s.GetTotalPointByID(1)
	if point == 10 {
		t.Log("right answer")
	}
	if point != 10 {
		t.Fatal("get total point by user id fail")
	}

}
