package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/nem0z/LightStash/wallet"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	w, err := wallet.New()
	handle(err)

	handle(w.Export("tmp/test.pem"))

	newWallet, err := wallet.Load("tmp/test.pem")
	handle(err)

	fmt.Println(bytes.Equal(w.Public, newWallet.Public))
	fmt.Println(w.P2PKH(true))
}
