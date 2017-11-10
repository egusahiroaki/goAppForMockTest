package main

import "github.com/goSamples/mysqlSample/service"
import "fmt"

func main() {
	s := service.NewUserService()
	point := s.GetTotalPointByID(1)
	fmt.Printf("point: %v\n", point)
}
