package main

import (
	"alpha-core/internal/config"
	"fmt"
)

func main() {
	configure := config.Load()

	fmt.Print(configure)
}
