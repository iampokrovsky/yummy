package main

import (
	"fmt"
	"github.com/kr/pretty"
	"hw-5/config"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%# v", pretty.Formatter(cfg))
}
