package main

import (
	"fmt"

	"github.com/example/modak/bootstrapper/servicebuilder"
)

func main() {
	service := servicebuilder.GetUserService()
	result := service.GetUserById(1)

	fmt.Println(result.UserName)
}
