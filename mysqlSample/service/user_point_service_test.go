package service

import (
	"fmt"
	"testing"

	"github.com/goSamples/mysqlSample/model/datasource"
	"github.com/goSamples/mysqlSample/service"
)

type DBMock struct {
}

func (db *DBMock) SumPointByUserID(_ int) int {
	return 10
}

// NewMySQL is
func NewDBMock() datasource.DataBase {
	return &DBMock{}
}

func setMockDB(db datasource.DataBase) {
	fmt.Printf("[setMockDB] db: %v\n", db)
	dbSet = func() datasource.DataBase { return db }
	// fmt.Printf("dbSet: %v\n", dbSet().SumPointByUserID(1))
}

// mockサーバをうまく使えない。MySQLを実行してしまう
func TestUserServiceGetTotalPointByID(t *testing.T) {

	fmt.Println("[Test Start] TestUserServiceGetTotalPointByID")
	// db := NewDBMock()
	// p, ok := db.(datasource.DataBase)
	// fmt.Printf("p: %v\n", p)
	// fmt.Printf("ok: %v\n", ok)

	// setMockDB(db)

	dbSet = NewDBMock

	s := service.NewUserService()
	point := s.GetTotalPointByID(1)
	if point == 10 {
		t.Log("right answer")
	}
	if point != 10 {
		t.Fatal("get total point by user id fail")
	}

}
