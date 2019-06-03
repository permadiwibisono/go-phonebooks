package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Load environment variables...")
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
}
