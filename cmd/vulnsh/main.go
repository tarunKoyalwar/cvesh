package main

import (
	"fmt"

	"github.com/tarunKoyalwar/vulnsh/client"
)

const (
	version = "v1.0.0"
)

func main() {
	fmt.Println("Hello, World! from vulnsh ->", version)
	client.NewClient()
	fmt.Println("exiting vulnsh")
}
