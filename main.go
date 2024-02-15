package main

import (
	"fmt"

	"github.com/aghilish/cryptit/encrypt"
	"github.com/aghilish/cryptit/decrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("Hello"))
	fmt.Println((decrypt.Nimbus(encrypt.Nimbus("Hello"))))
}
