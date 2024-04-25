package main

import (
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

	fmt.Println(w.P2PKH())
	fmt.Println(w.Bench32())
}
