package main

import "github.com/goSamples/mysqlSample/service"
import "fmt"

func main() {
	// r := model.NewUserPointRepository()
	// result := r.GetTotalPointByUserID(1)
	// fmt.Println(result)
	s := service.NewUserService()
	point := s.GetTotalPointByID(1)
	fmt.Printf("point: %v\n", point)
}
