package model

import (
	"testing"

	"github.com/goSamples/mysqlSample/model"
)

type DBMock struct{}

func (db DBMock) SumPointByUserID(_ int) int {
	return 10
}

func TestUserRepositoryGetTotalPointByUserID(t *testing.T) {
	dbmock := DBMock{}
	r := model.NewUserPointRepositoryImpl(dbmock)
	result := r.GetTotalPointByUserID(1)
	if result == 10 {
		t.Log("right answer")
	}
	if result != 10 {
		t.Fatal("get total point by user id fail")
	}
}
