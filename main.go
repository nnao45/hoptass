package main

import (
	"fmt"

	"github.com/nao4arale/hopass/hopass"
)

func main() {
	fmt.Println(hopass.NewPassword(8))
}
