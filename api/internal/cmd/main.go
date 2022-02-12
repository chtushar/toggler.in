package cmd

import (
	"fmt"

	"toggler.in/internal/configs"
)

func Execute()  {
	cfg := configs.Get()

	fmt.Printf("%+v\n", cfg)
}