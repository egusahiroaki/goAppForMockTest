package main

import (
	"fmt"

	"github.com/goSamples/mysqlSample/model"
)

func main() {
	r := model.NewUserPointRepository()
	result := r.GetTotalPointByUserID(1)
	fmt.Println(result)
}
