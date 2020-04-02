package main

import (
	"fmt"

	"github.com/yekai1003/bcwithme/wallet/hdwallet"
)

func main() {
	w, err := hdwallet.NewWallet("./keystore")
	if err != nil {
		fmt.Println("Failed to NewWallet", err)
		return
	}
	w.StoreKey("123")
}
