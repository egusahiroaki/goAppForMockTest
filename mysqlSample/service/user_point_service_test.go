package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goSamples/mysqlSample/model/datasource"

	"github.com/goSamples/mysqlSample/service"
)

type DBMock struct{}

func (db *DBMock) SumPointByUserID(_ int) int {
	return 10
}

func setMockDB(db datasource.DataBase) {
	fmt.Printf("reflect.TypeOf(db): %v\n", reflect.TypeOf(db))
	p, ok := db.(datasource.DataBase)
	fmt.Printf("p: %v, ok: %v\n", p, ok)
	fmt.Printf("[setMockDB] before\n")
	dbSet = func() datasource.DataBase {
		fmt.Printf("[setMockDB] hogehogedb: %v\n", db)
		return db
	}
	dbSet = func() datasource.DataBase { return &DBMock{} }
	fmt.Printf("dbSet: %v\n", dbSet())
	fmt.Printf("[setMockDB] after\n")
}

// mockサーバをうまく使えない。MySQLを実行してしまう
func TestUserServiceGetTotalPointByID(t *testing.T) {

	fmt.Println("[Test Start] TestUserServiceGetTotalPointByID")
	hoge := &DBMock{}
	// p, ok := db.(datasource.DataBase)
	// fmt.Printf("p: %v\n", p)
	// fmt.Printf("ok: %v\n", ok)

	setMockDB(hoge)
	dbSet = func() datasource.DataBase {
		fmt.Printf("[setMockDB] hogehogedb: %v\n", hoge)
		return hoge
	}

	fmt.Println("After SetMockDB")
	s := service.NewUserService()
	point := s.GetTotalPointByID(10)
	if point == 10 {
		t.Log("right answer")
	}
	if point != 10 {
		t.Fatal("get total point by user id fail")
	}

}
