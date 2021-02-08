package main

import (
	"fmt"
	"newsapp/config"
)

func main() {
	conf := config.GetConfig()
	fmt.Println((conf))
}
