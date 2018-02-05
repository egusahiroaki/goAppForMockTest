package datasource

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// DataBase is
type DataBase interface {
	SumPointByUserID(userID int) int
}

// // DatabaseFactory is to convert MySQL type to DataBase Interface
// type DatabaseFactory func() DataBase

// MySQL is
type mySQL struct {
	db *sql.DB
}

// SetDB is
func SetDB() DataBase {
	fmt.Println("[db.go] SetDB called")
	return NewMySQL()
}

// NewMySQL is
func NewMySQL() DataBase {
	db, err := sql.Open("mysql", "root:@/user_point")
	if err != nil {
		panic(err.Error())
	}

	return &mySQL{db: db}
}

// SumPointByUserID is
// [todo] return should error
func (ur *mySQL) SumPointByUserID(userID int) int {
	defer ur.db.Close()
	// rows, err := db.Query("select sum(point) from user_point where user_id =1;")
	query := "select sum(point) from user_point where user_id =" + strconv.Itoa(userID)
	fmt.Printf("query: %v\n", query)
	rows, err := ur.db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var value string
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	result, _ := strconv.Atoi(value)
	return result
}
