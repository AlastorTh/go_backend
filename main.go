package main

import (
	"fmt"

	"github.com/AlastorTh/go_backend/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println(conf)
}
