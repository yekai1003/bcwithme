package main

import (
	"fmt"
	"log"

	"github.com/tyler-smith/go-bip39"
)

func create_mnemonic() {

	//1. entropy
	b, err := bip39.NewEntropy(128)
	if err != nil {
		log.Panic("failed to NewEntropy", err)
	}
	//2. mnemonic
	mne, err := bip39.NewMnemonic(b)
	if err != nil {
		log.Panic("failed to NewMnemonic", err)
	}
	fmt.Println(mne)

}
func main() {
	create_mnemonic()
}
