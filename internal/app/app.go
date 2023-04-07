// Package app configures and runs application.
package app

import (
	"fmt"
	"github.com/kr/pretty"
	"hw-5/config"
	"time"
)

// Run creates objects via constructors.
func Run(cfg config.Config) {

	for {
		fmt.Printf("%# v", pretty.Formatter(cfg))
		time.Sleep(time.Second)
	}
}
