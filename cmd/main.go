package main

import (
	"fmt"
	jwtconfig "go_jwt/config/jwt_config"
)

func main() {

	response, _ := jwtconfig.Load()

	fmt.Println(response)
}
