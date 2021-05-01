package main

import (
	"embed"
	"fmt"
)

//go:embed adir
var spa embed.FS

func main() {
	fmt.Print(spa)
}
